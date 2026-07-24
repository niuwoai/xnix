#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
WORK_ROOT = PROJECT_ROOT.join(".local", "xnix", "winapp-smoke")
GO_CACHE_ROOT = PROJECT_ROOT.join(".gocache")
EXE_PATH = WORK_ROOT.join("hello.exe")
STATE_ROOT = WORK_ROOT.join("state")
MARKER = "XNIX_WINAPP_SMOKE_OK"
SCHEMA_VERSION = "xnix.runtime.winapp_smoke_report.v1"

options = {
  format: "text",
  redact_output: nil,
  backend: "local",
  profile: nil,
  write_profile: nil,
  preflight_only: false,
  exe: nil,
  runner: nil,
  docker: nil,
  image: ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local"),
  platform: ENV.fetch("XNIX_WINE_PLATFORM", "linux/amd64"),
  state_root: STATE_ROOT.to_s,
  working_dir: nil,
  timeout: "30s",
  bootstrap_timeout: "300s",
  expected_marker: MARKER,
  success_mode: "marker",
  skip_bootstrap: false,
  stage_app_dir: false,
  runner_bottle: nil,
  runner_args: [],
  app_args: []
}

OptionParser.new do |parser|
  parser.banner = "Usage: winapp_smoke.rb [--format text|json|markdown] [--backend local|container] [--profile PATH] [--write-profile PATH] [--preflight-only] [--exe PATH] [--runner PATH] [--runner-bottle NAME] [--runner-arg VALUE] [--success-mode marker|exit-code|startup-window] [--skip-bootstrap] [--stage-app-dir] [--arg VALUE]"
  parser.on("--format FORMAT", "Output format: text, json, or markdown") { |value| options[:format] = value }
  parser.on("--backend BACKEND", "Execution backend: local or container") { |value| options[:backend] = value }
  parser.on("--redact-output", "Request redacted Runtime smoke output") { options[:redact_output] = true }
  parser.on("--profile PATH", "Windows app smoke profile JSON path") { |value| options[:profile] = value }
  parser.on("--write-profile PATH", "Write a reusable Windows app smoke profile for the resolved settings") { |value| options[:write_profile] = value }
  parser.on("--preflight-only", "Validate a profile and emit readiness without launching the Windows app") { options[:preflight_only] = true }
  parser.on("--exe PATH", "Existing Windows executable path; defaults to the built fixture") { |value| options[:exe] = value }
  parser.on("--runner PATH", "Explicit compatibility runner path") { |value| options[:runner] = value }
  parser.on("--docker PATH", "Explicit Docker runner path for container backend") { |value| options[:docker] = value }
  parser.on("--image IMAGE", "Local Wine container image for container backend") { |value| options[:image] = value }
  parser.on("--platform PLATFORM", "Container platform for container backend") { |value| options[:platform] = value }
  parser.on("--state-root PATH", "Isolated Runtime state root") { |value| options[:state_root] = value }
  parser.on("--working-dir PATH", "Working directory for the compatibility runner; defaults to the executable directory") { |value| options[:working_dir] = value }
  parser.on("--timeout DURATION", "Execution timeout") { |value| options[:timeout] = value }
  parser.on("--bootstrap-timeout DURATION", "Wine prefix bootstrap timeout") { |value| options[:bootstrap_timeout] = value }
  parser.on("--expected-marker MARKER", "Expected stdout marker") { |value| options[:expected_marker] = value }
  parser.on("--success-mode MODE", "Success mode: marker, exit-code, or startup-window") { |value| options[:success_mode] = value }
  parser.on("--skip-bootstrap", "Skip Wine prefix bootstrap before app execution") { options[:skip_bootstrap] = true }
  parser.on("--stage-app-dir", "Stage the application directory into the isolated Runtime state root before app execution") { options[:stage_app_dir] = true }
  parser.on("--runner-bottle NAME", "Compatibility runner bottle name passed before the executable path") { |value| options[:runner_bottle] = value }
  parser.on("--runner-arg VALUE", "Argument passed to the compatibility runner before the executable path") { |value| options[:runner_args] << value }
  parser.on("--arg VALUE", "Argument passed to the Windows executable") { |value| options[:app_args] << value }
end.parse!

unless options[:profile].to_s.strip.empty?
  begin
    profile = JSON.parse(File.read(options.fetch(:profile)))
    unless profile.fetch("schema_version", "") == "xnix.runtime.windows_app_smoke_profile.v1"
      warn "FAIL: unsupported Windows app smoke profile schema"
      exit 1
    end
    options[:exe] = profile["executable_path"] if options[:exe].to_s.strip.empty? && !profile["executable_path"].to_s.strip.empty?
    options[:state_root] = profile["state_root"] if options[:state_root] == STATE_ROOT.to_s && !profile["state_root"].to_s.strip.empty?
    options[:working_dir] = profile["working_directory"] if options[:working_dir].to_s.strip.empty? && !profile["working_directory"].to_s.strip.empty?
    options[:runner] = profile["runner_path"] if options[:runner].to_s.strip.empty? && !profile["runner_path"].to_s.strip.empty?
    options[:runner_bottle] = profile["runner_bottle"] if options[:runner_bottle].to_s.strip.empty? && !profile["runner_bottle"].to_s.strip.empty?
    options[:timeout] = profile["timeout"] if options[:timeout] == "30s" && !profile["timeout"].to_s.strip.empty?
    options[:expected_marker] = profile["expected_marker"] if options[:expected_marker] == MARKER && !profile["expected_marker"].to_s.strip.empty?
    options[:success_mode] = profile["success_mode"] if options[:success_mode] == "marker" && !profile["success_mode"].to_s.strip.empty?
    options[:redact_output] = profile["redact_output"] unless options.key?(:redact_output) && !options[:redact_output].nil?
    options[:skip_bootstrap] = profile["skip_bootstrap"] if !options[:skip_bootstrap] && profile.key?("skip_bootstrap")
    options[:stage_app_dir] = profile["stage_app_dir"] if !options[:stage_app_dir] && profile.key?("stage_app_dir")
    options[:runner_args] = Array(profile["runner_arguments"]) + options.fetch(:runner_args)
    options[:app_args] = Array(profile["arguments"]) + options.fetch(:app_args)
  rescue JSON::ParserError, KeyError, Errno::ENOENT
    warn "FAIL: Windows app smoke profile invalid"
    exit 1
  end
end

unless %w[text json markdown].include?(options[:format])
  warn "FAIL: unsupported output format #{options[:format]}"
  exit 1
end
unless %w[local container].include?(options[:backend])
  warn "FAIL: unsupported backend #{options[:backend]}"
  exit 1
end
unless %w[marker exit-code startup-window].include?(options[:success_mode])
  warn "FAIL: unsupported success mode #{options[:success_mode]}"
  exit 1
end
if options.fetch(:preflight_only) && options[:profile].to_s.strip.empty?
  warn "FAIL: --preflight-only requires --profile"
  exit 1
end

options[:redact_output] = options[:format] != "text" if options[:redact_output].nil?

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def base_report(format, redact_output, expected_marker, success_mode, executable_source, profile_supplied, backend, image, platform)
  {
    "version" => PROJECT_ROOT.join("VERSION").read.strip,
    "schema_version" => SCHEMA_VERSION,
    "report_type" => "winapp-smoke",
    "format" => format,
    "backend" => backend,
    "redacted_output_requested" => redact_output,
    "executable_source" => executable_source,
    "executable_format" => "unknown",
    "windows_executable_signature_observed" => false,
    "executable_architecture" => "unknown",
    "executable_architecture_supported" => false,
    "wine_architecture" => "unknown",
    "wine_prefix_mode" => "unknown",
    "wine_prefix_prepared" => false,
    "profile_supplied" => profile_supplied,
    "profile_write_invoked" => false,
    "profile_written" => false,
    "profile_preflight_invoked" => false,
    "profile_preflight_status" => "not-run",
    "profile_preflight_next_action" => "",
    "profile_preflight_payload" => nil,
    "preflight_only" => false,
    "user_executable_supplied" => executable_source != "fixture",
    "fixture_built" => false,
    "runner_diagnostics_invoked" => false,
    "smoke_invoked" => false,
    "status" => "failed",
    "marker" => expected_marker,
    "success_mode" => success_mode,
    "working_directory_mode" => "executable-directory",
    "application_workspace_mode" => "direct-executable",
    "application_staged" => false,
    "application_staged_file_count" => 0,
    "application_staged_bytes" => 0,
    "runner_available" => false,
    "runner_argument_count" => 0,
    "runner_diagnostics_status" => "not-run",
    "env_runner_configured" => false,
    "runner_candidate_count" => 0,
    "runner_diagnostics_next_action" => "",
    "runner_command_hints" => [],
    "wine_bootstrap_attempted" => false,
    "wine_bootstrap_succeeded" => false,
    "wine_bootstrap_skipped" => false,
    "wine_bootstrap_exit_code" => -1,
    "marker_observed" => false,
    "startup_window_observed" => false,
    "raw_output_included" => false,
    "raw_output_redacted" => redact_output,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "host_networking_required" => false,
    "docker_socket_mounted" => false,
    "broad_host_mount_required" => false,
    "docker_executed" => false,
    "qemu_executed" => false,
    "wine_executed_by_script" => false,
    "colima_executed" => false,
    "network_checks_run" => false,
    "package_manager_invoked" => false,
    "kde_safe_output_summary" => "smoke has not run",
    "failure_reason" => "",
    "skip_reason" => "",
    "runner_diagnostics_payload" => nil,
    "container_smoke_invoked" => false,
    "container_image" => image,
    "container_platform" => platform,
    "container_image_available" => false,
    "container_payload" => nil,
    "runtime_payload" => nil
  }
end

def emit_report(report)
  case report.fetch("format")
  when "json"
    puts JSON.pretty_generate(report)
  when "markdown"
    puts "# Windows App Smoke Report"
    puts
    puts "- Version: #{report.fetch("version")}"
    puts "- Backend: #{report.fetch("backend")}"
    puts "- Status: #{report.fetch("status")}"
    puts "- Fixture built: #{report.fetch("fixture_built")}"
    puts "- Executable format: #{report.fetch("executable_format")}"
    puts "- Windows executable signature observed: #{report.fetch("windows_executable_signature_observed")}"
    puts "- Executable architecture: #{report.fetch("executable_architecture")}"
    puts "- Executable architecture supported: #{report.fetch("executable_architecture_supported")}"
    puts "- Wine architecture: #{report.fetch("wine_architecture")}"
    puts "- Wine prefix mode: #{report.fetch("wine_prefix_mode")}"
    puts "- Wine prefix prepared: #{report.fetch("wine_prefix_prepared")}"
    puts "- Runner diagnostics invoked: #{report.fetch("runner_diagnostics_invoked")}"
    puts "- Runner diagnostics status: #{report.fetch("runner_diagnostics_status")}"
    puts "- Runner candidate count: #{report.fetch("runner_candidate_count")}"
    puts "- Env runner configured: #{report.fetch("env_runner_configured")}"
    puts "- Profile preflight invoked: #{report.fetch("profile_preflight_invoked")}"
    puts "- Profile preflight status: #{report.fetch("profile_preflight_status")}"
    puts "- Profile written: #{report.fetch("profile_written")}"
    puts "- Success mode: #{report.fetch("success_mode")}"
    puts "- Working directory mode: #{report.fetch("working_directory_mode")}"
    puts "- Application workspace mode: #{report.fetch("application_workspace_mode")}"
    puts "- Application staged: #{report.fetch("application_staged")}"
    puts "- Application staged file count: #{report.fetch("application_staged_file_count")}"
    puts "- Runner argument count: #{report.fetch("runner_argument_count")}"
    unless report.fetch("runner_command_hints").empty?
      puts "- Runner command hints:"
      report.fetch("runner_command_hints").each { |hint| puts "  - #{hint}" }
    end
    puts "- Smoke invoked: #{report.fetch("smoke_invoked")}"
    puts "- Container smoke invoked: #{report.fetch("container_smoke_invoked")}"
    puts "- Container image available: #{report.fetch("container_image_available")}"
    puts "- Runner available: #{report.fetch("runner_available")}"
    puts "- Wine bootstrap attempted: #{report.fetch("wine_bootstrap_attempted")}"
    puts "- Wine bootstrap succeeded: #{report.fetch("wine_bootstrap_succeeded")}"
    puts "- Wine bootstrap skipped: #{report.fetch("wine_bootstrap_skipped")}"
    puts "- Wine bootstrap exit code: #{report.fetch("wine_bootstrap_exit_code")}"
    puts "- Marker observed: #{report.fetch("marker_observed")}"
    puts "- Startup window observed: #{report.fetch("startup_window_observed")}"
    puts "- Raw output redacted: #{report.fetch("raw_output_redacted")}"
    puts "- KDE-safe output summary: #{report.fetch("kde_safe_output_summary")}"
    puts "- Host root modified: #{report.fetch("host_root_modified")}"
    puts "- Docker executed: #{report.fetch("docker_executed")}"
    puts "- QEMU executed: #{report.fetch("qemu_executed")}"
    puts "- Wine executed by script: #{report.fetch("wine_executed_by_script")}"
    puts "- Network checks run: #{report.fetch("network_checks_run")}"
    puts "- Package manager invoked: #{report.fetch("package_manager_invoked")}"
    puts "- Profile preflight next action: #{report.fetch("profile_preflight_next_action")}" unless report.fetch("profile_preflight_next_action").empty?
    puts "- Runner diagnostics next action: #{report.fetch("runner_diagnostics_next_action")}" unless report.fetch("runner_diagnostics_next_action").empty?
    puts "- Failure reason: #{report.fetch("failure_reason")}" unless report.fetch("failure_reason").empty?
    puts "- Skip reason: #{report.fetch("skip_reason")}" unless report.fetch("skip_reason").empty?
  end
end

def finish(report, exit_code)
  emit_report(report) unless report.fetch("format") == "text"
  exit exit_code
end

profile_supplied = !options[:profile].to_s.strip.empty?
executable_source = if options[:exe].to_s.strip.empty?
                      "fixture"
                    elsif profile_supplied
                      "profile"
                    else
                      "user-supplied"
                    end
report = base_report(
  options.fetch(:format),
  options.fetch(:redact_output),
  options.fetch(:expected_marker),
  options.fetch(:success_mode),
  executable_source,
  profile_supplied,
  options.fetch(:backend),
  options.fetch(:image),
  options.fetch(:platform)
)
report["preflight_only"] = options.fetch(:preflight_only)

FileUtils.mkdir_p(WORK_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))

selected_exe_path = options[:exe]
if executable_source == "fixture"
  FileUtils.rm_f(EXE_PATH)
  build_stdout, build_stderr, build_status = run_command(
    {
      "GOOS" => "windows",
      "GOARCH" => "amd64",
      "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
      "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
    },
    "go", "build", "-o", EXE_PATH.to_s, "./test/fixtures/winapp/hello"
  )

  unless build_status.zero?
    report["failure_reason"] = "Windows app fixture build failed"
    if options.fetch(:format) == "text"
      warn build_stdout unless build_stdout.empty?
      warn build_stderr unless build_stderr.empty?
      warn "FAIL: Windows app fixture build failed"
    end
    finish(report, 1)
  end
  report["fixture_built"] = true
  selected_exe_path = EXE_PATH.to_s
end

go_env = {
  "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
  "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
}

unless options[:write_profile].to_s.strip.empty?
  render_command = [
    "go", "run", "./cmd/xnix-runtime-go", "windows-app-smoke-profile-render",
    "--exe", selected_exe_path.to_s,
    "--state-root", options.fetch(:state_root),
    "--timeout", options.fetch(:timeout),
    "--expected-marker", options.fetch(:expected_marker),
    "--success-mode", options.fetch(:success_mode)
  ]
  render_command.concat(["--profile", options.fetch(:profile)]) unless options[:profile].to_s.strip.empty?
  render_command.concat(["--runner", options.fetch(:runner)]) unless options[:runner].to_s.strip.empty?
  render_command.concat(["--working-dir", options.fetch(:working_dir)]) unless options[:working_dir].to_s.strip.empty?
  render_command.concat(["--runner-bottle", options.fetch(:runner_bottle)]) unless options[:runner_bottle].to_s.strip.empty?
  options.fetch(:runner_args).each { |value| render_command.concat(["--runner-arg", value]) }
  options.fetch(:app_args).each { |value| render_command.concat(["--arg", value]) }
  render_command << "--redact-output" if options.fetch(:redact_output)
  render_command << "--skip-bootstrap" if options.fetch(:skip_bootstrap)
  render_command << "--stage-app-dir" if options.fetch(:stage_app_dir)

  profile_stdout, profile_stderr, profile_status = run_command(go_env, *render_command)
  report["profile_write_invoked"] = true
  unless profile_status.zero?
    report["failure_reason"] = "Windows app smoke profile render command failed"
    if options.fetch(:format) == "text"
      warn profile_stdout unless profile_stdout.empty?
      warn profile_stderr unless profile_stderr.empty?
      warn "FAIL: Windows app smoke profile render command failed"
    end
    finish(report, 1)
  end
  write_profile_path = Pathname.new(options.fetch(:write_profile))
  write_profile_path = PROJECT_ROOT.join(write_profile_path) unless write_profile_path.absolute?
  FileUtils.mkdir_p(write_profile_path.dirname)
  File.write(write_profile_path, profile_stdout)
  report["profile_written"] = true
end

if profile_supplied
  preflight_command = [
    "go", "run", "./cmd/xnix-runtime-go", "windows-app-smoke-profile-preflight",
    "--profile", options.fetch(:profile)
  ]
  preflight_stdout, preflight_stderr, preflight_status = run_command(go_env, *preflight_command)
  report["profile_preflight_invoked"] = true

  unless preflight_status.zero?
    report["failure_reason"] = "Windows app profile preflight command failed"
    if options.fetch(:format) == "text"
      warn preflight_stdout unless preflight_stdout.empty?
      warn preflight_stderr unless preflight_stderr.empty?
      warn "FAIL: Windows app profile preflight command failed"
    end
    finish(report, 1)
  end

  preflight_payload = JSON.parse(preflight_stdout)
  report["profile_preflight_payload"] = preflight_payload
  report["profile_preflight_status"] = preflight_payload.fetch("status")
  report["profile_preflight_next_action"] = preflight_payload.fetch("next_action", "")
  report["executable_format"] = preflight_payload.fetch("executable_format", "unknown")
  report["windows_executable_signature_observed"] = preflight_payload.fetch("windows_executable_signature_observed", false)
  report["executable_architecture"] = preflight_payload.fetch("executable_architecture", "unknown")
  report["executable_architecture_supported"] = preflight_payload.fetch("executable_architecture_supported", false)
  report["wine_architecture"] = preflight_payload.fetch("wine_architecture", "unknown")
  report["wine_prefix_mode"] = preflight_payload.fetch("wine_prefix_mode", "unknown")
  report["wine_prefix_prepared"] = preflight_payload.fetch("wine_prefix_prepared", false)
  report["runner_available"] = preflight_payload.fetch("runner_available", false)
  report["runner_argument_count"] = preflight_payload.fetch("runner_argument_count", 0)
  report["wine_bootstrap_skipped"] = preflight_payload.fetch("skip_bootstrap", false)
  report["application_workspace_mode"] = preflight_payload.fetch("application_workspace_mode", "direct-executable")
  report["application_staged"] = false
  report["working_directory_mode"] = preflight_payload.fetch("working_directory_mode", "executable-directory")
  report["host_root_modified"] = preflight_payload.fetch("host_root_modified", false)
  report["privileged_container_required"] = preflight_payload.fetch("privileged_container_required", false)
  report["host_networking_required"] = preflight_payload.fetch("host_networking_required", false)
  report["docker_socket_mounted"] = preflight_payload.fetch("docker_socket_mounted", false)
  report["broad_host_mount_required"] = preflight_payload.fetch("broad_host_mount_required", false)
  report["docker_executed"] = preflight_payload.fetch("docker_executed", false)
  report["qemu_executed"] = preflight_payload.fetch("qemu_executed", false)
  report["wine_executed_by_script"] = preflight_payload.fetch("wine_executed", false)
  report["colima_executed"] = preflight_payload.fetch("colima_executed", false)
  report["network_checks_run"] = preflight_payload.fetch("network_checks_run", false)
  report["package_manager_invoked"] = preflight_payload.fetch("package_manager_invoked", false)

  if preflight_payload.fetch("status") == "ready"
    report["status"] = "ready"
    finish(report, 0) if options.fetch(:preflight_only)
  elsif preflight_payload.fetch("skip_reason", "") != ""
    report["status"] = "skipped"
    report["skip_reason"] = preflight_payload.fetch("skip_reason")
    puts "SKIP: Windows app profile preflight (#{report.fetch("skip_reason")})" if options.fetch(:format) == "text"
    finish(report, 0)
  else
    report["status"] = "blocked"
    report["failure_reason"] = preflight_payload.fetch("failure_reason", "Windows app profile preflight blocked")
    warn "FAIL: Windows app profile preflight blocked" if options.fetch(:format) == "text"
    finish(report, 1)
  end
end

if options.fetch(:backend) == "container"
  container_command = [
    "go", "run", "./cmd/xnix-runtime-go", "windows-app-container-run-smoke",
    "--exe", selected_exe_path.to_s,
    "--state-root", options.fetch(:state_root),
    "--image", options.fetch(:image),
    "--platform", options.fetch(:platform),
    "--timeout", options.fetch(:timeout),
    "--bootstrap-timeout", options.fetch(:bootstrap_timeout)
  ]
  container_command.concat(["--docker", options.fetch(:docker)]) unless options[:docker].to_s.strip.empty?
  options.fetch(:app_args).each { |value| container_command.concat(["--arg", value]) }

  container_stdout, container_stderr, container_status = run_command(go_env, *container_command)
  report["container_smoke_invoked"] = true
  report["smoke_invoked"] = true

  unless container_status.zero?
    report["failure_reason"] = "Windows app container smoke command failed"
    if options.fetch(:format) == "text"
      warn container_stdout unless container_stdout.empty?
      warn container_stderr unless container_stderr.empty?
      warn "FAIL: Windows app container smoke command failed"
    end
    finish(report, 1)
  end

  payload = JSON.parse(container_stdout)
  report["container_payload"] = payload
  report["runtime_payload"] = payload
  report["status"] = payload.fetch("status")
  report["runner_available"] = payload.fetch("runner_available", false)
  report["container_image"] = payload.fetch("container_image", options.fetch(:image))
  report["container_platform"] = payload.fetch("container_platform", options.fetch(:platform))
  report["container_image_available"] = payload.fetch("image_available", false)
  report["wine_bootstrap_attempted"] = payload.fetch("wine_bootstrap_required", false) && payload.fetch("runner_available", false) && payload.fetch("image_available", false)
  report["wine_bootstrap_succeeded"] = payload.fetch("status") == "passed" && payload.fetch("wine_bootstrap_exit_code", -1) == -1
  report["wine_bootstrap_exit_code"] = payload.fetch("wine_bootstrap_exit_code", -1)
  report["marker_observed"] = payload.fetch("marker_observed", false)
  report["host_root_modified"] = payload.fetch("host_root_modified", false)
  report["privileged_container_required"] = payload.fetch("privileged_container_required", false)
  report["host_networking_required"] = payload.fetch("host_networking_required", false)
  report["docker_socket_mounted"] = payload.fetch("docker_socket_mounted", false)
  report["broad_host_mount_required"] = payload.fetch("broad_host_mount_required", false)
  report["docker_executed"] = payload.fetch("runner_available", false)
  report["failure_reason"] = payload.fetch("failure_reason", "")
  report["skip_reason"] = payload.fetch("skip_reason", "")

  case payload.fetch("status")
  when "passed"
    if payload["marker_observed"] && payload["stdout"].include?(options.fetch(:expected_marker))
      puts "PASS: containerized real Windows app smoke" if options.fetch(:format) == "text"
      finish(report, 0)
    end
    report["failure_reason"] = "Containerized Windows app smoke marker missing"
    warn "FAIL: containerized Windows app smoke marker missing" if options.fetch(:format) == "text"
    finish(report, 1)
  when "skipped"
    puts "SKIP: containerized real Windows app smoke (#{payload.fetch("skip_reason")})" if options.fetch(:format) == "text"
    finish(report, 0)
  else
    if options.fetch(:format) == "text"
      warn container_stdout
      warn container_stderr unless container_stderr.empty?
      warn "FAIL: containerized real Windows app smoke"
    end
    finish(report, 1)
  end
end

diagnostics_command = ["go", "run", "./cmd/xnix-runtime-go", "windows-app-runner-diagnostics"]
diagnostics_command.concat(["--runner", options.fetch(:runner)]) unless options[:runner].to_s.strip.empty?
diagnostics_stdout, diagnostics_stderr, diagnostics_status = run_command(go_env, *diagnostics_command)
report["runner_diagnostics_invoked"] = true

unless diagnostics_status.zero?
  report["failure_reason"] = "Windows app runner diagnostics command failed"
  if options.fetch(:format) == "text"
    warn diagnostics_stdout unless diagnostics_stdout.empty?
    warn diagnostics_stderr unless diagnostics_stderr.empty?
    warn "FAIL: Windows app runner diagnostics command failed"
  end
  finish(report, 1)
end

diagnostics_payload = JSON.parse(diagnostics_stdout)
report["runner_diagnostics_payload"] = diagnostics_payload
report["runner_diagnostics_status"] = diagnostics_payload.fetch("status")
report["env_runner_configured"] = diagnostics_payload.fetch("env_runner_configured", false)
report["runner_candidate_count"] = diagnostics_payload.fetch("candidate_count", 0)
report["runner_diagnostics_next_action"] = diagnostics_payload.fetch("next_action", "")
report["runner_command_hints"] = diagnostics_payload.fetch("runner_command_hints", [])
report["runner_available"] = diagnostics_payload.fetch("runner_available", false)

smoke_command = [
  "go", "run", "./cmd/xnix-runtime-go", "windows-app-run-smoke",
  "--exe", selected_exe_path.to_s,
  "--state-root", options.fetch(:state_root),
  "--timeout", options.fetch(:timeout),
  "--expected-marker", options.fetch(:expected_marker),
  "--success-mode", options.fetch(:success_mode)
]
smoke_command.concat(["--runner", options.fetch(:runner)]) unless options[:runner].to_s.strip.empty?
smoke_command.concat(["--working-dir", options.fetch(:working_dir)]) unless options[:working_dir].to_s.strip.empty?
smoke_command.concat(["--runner-bottle", options.fetch(:runner_bottle)]) unless options[:runner_bottle].to_s.strip.empty?
options.fetch(:runner_args).each { |value| smoke_command.concat(["--runner-arg", value]) }
options.fetch(:app_args).each { |value| smoke_command.concat(["--arg", value]) }
smoke_command << "--redact-output" if options.fetch(:redact_output)
smoke_command << "--skip-bootstrap" if options.fetch(:skip_bootstrap)
smoke_command << "--stage-app-dir" if options.fetch(:stage_app_dir)

smoke_stdout, smoke_stderr, smoke_status = run_command(go_env, *smoke_command)
report["smoke_invoked"] = true

unless smoke_status.zero?
  report["failure_reason"] = "Windows app runtime smoke command failed"
  if options.fetch(:format) == "text"
    warn smoke_stdout unless smoke_stdout.empty?
    warn smoke_stderr unless smoke_stderr.empty?
    warn "FAIL: Windows app runtime smoke command failed"
  end
  finish(report, 1)
end

payload = JSON.parse(smoke_stdout)
report["runtime_payload"] = payload
report["status"] = payload.fetch("status")
report["executable_format"] = payload.fetch("executable_format", "unknown")
report["windows_executable_signature_observed"] = payload.fetch("windows_executable_signature_observed", false)
report["executable_architecture"] = payload.fetch("executable_architecture", "unknown")
report["executable_architecture_supported"] = payload.fetch("executable_architecture_supported", false)
report["wine_architecture"] = payload.fetch("wine_architecture", "unknown")
report["wine_prefix_mode"] = payload.fetch("wine_prefix_mode", "unknown")
report["wine_prefix_prepared"] = payload.fetch("wine_prefix_prepared", false)
report["runner_available"] = payload.fetch("runner_available", false)
report["success_mode"] = payload.fetch("success_mode", options.fetch(:success_mode))
report["working_directory_mode"] = payload.fetch("working_directory_mode", "executable-directory")
report["application_workspace_mode"] = payload.fetch("application_workspace_mode", "direct-executable")
report["application_staged"] = payload.fetch("application_staged", false)
report["application_staged_file_count"] = payload.fetch("application_staged_file_count", 0)
report["application_staged_bytes"] = payload.fetch("application_staged_bytes", 0)
report["runner_argument_count"] = payload.fetch("runner_argument_count", 0)
report["wine_bootstrap_attempted"] = payload.fetch("wine_bootstrap_attempted", false)
report["wine_bootstrap_succeeded"] = payload.fetch("wine_bootstrap_succeeded", false)
report["wine_bootstrap_skipped"] = payload.fetch("wine_bootstrap_skipped", false)
report["wine_bootstrap_exit_code"] = payload.fetch("wine_bootstrap_exit_code", -1)
report["marker_observed"] = payload.fetch("marker_observed", false)
report["startup_window_observed"] = payload.fetch("startup_window_observed", false)
report["raw_output_included"] = payload.fetch("raw_output_included", false)
report["raw_output_redacted"] = payload.fetch("raw_output_redacted", options.fetch(:redact_output))
report["host_root_modified"] = payload.fetch("host_root_modified", false)
report["privileged_container_required"] = payload.fetch("privileged_container_required", false)
report["host_networking_required"] = payload.fetch("host_networking_required", false)
report["docker_socket_mounted"] = payload.fetch("docker_socket_mounted", false)
report["broad_host_mount_required"] = payload.fetch("broad_host_mount_required", false)
report["kde_safe_output_summary"] = payload.fetch("kde_safe_output_summary", "")
report["failure_reason"] = payload.fetch("failure_reason", "")
report["skip_reason"] = payload.fetch("skip_reason", "")

case payload.fetch("status")
when "passed"
  if %w[exit-code startup-window].include?(report.fetch("success_mode"))
    puts "PASS: real Windows app smoke" if options.fetch(:format) == "text"
    finish(report, 0)
  end
  if payload["marker_observed"] && payload["stdout"].include?(options.fetch(:expected_marker))
    puts "PASS: real Windows app smoke"
    exit 0
  end
  if payload["marker_observed"] && options.fetch(:redact_output)
    finish(report, 0)
  end
  report["failure_reason"] = "Windows app smoke marker missing"
  warn "FAIL: Windows app smoke marker missing" if options.fetch(:format) == "text"
  finish(report, 1)
when "skipped"
  puts "SKIP: real Windows app smoke (windows compatibility runner unavailable)" if options.fetch(:format) == "text"
  finish(report, 0)
else
  if options.fetch(:format) == "text"
    warn smoke_stdout
    warn smoke_stderr unless smoke_stderr.empty?
    warn "FAIL: real Windows app smoke"
  end
  finish(report, 1)
end
