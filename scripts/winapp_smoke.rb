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
  exe: nil,
  runner: nil,
  docker: nil,
  image: ENV.fetch("XNIX_WINE_IMAGE", "xnix-wine-smoke:local"),
  platform: ENV.fetch("XNIX_WINE_PLATFORM", "linux/amd64"),
  state_root: STATE_ROOT.to_s,
  timeout: "30s",
  bootstrap_timeout: "300s",
  expected_marker: MARKER,
  runner_args: [],
  app_args: []
}

OptionParser.new do |parser|
  parser.banner = "Usage: winapp_smoke.rb [--format text|json|markdown] [--backend local|container] [--exe PATH] [--runner PATH] [--runner-arg VALUE] [--arg VALUE]"
  parser.on("--format FORMAT", "Output format: text, json, or markdown") { |value| options[:format] = value }
  parser.on("--backend BACKEND", "Execution backend: local or container") { |value| options[:backend] = value }
  parser.on("--redact-output", "Request redacted Runtime smoke output") { options[:redact_output] = true }
  parser.on("--exe PATH", "Existing Windows executable path; defaults to the built fixture") { |value| options[:exe] = value }
  parser.on("--runner PATH", "Explicit compatibility runner path") { |value| options[:runner] = value }
  parser.on("--docker PATH", "Explicit Docker runner path for container backend") { |value| options[:docker] = value }
  parser.on("--image IMAGE", "Local Wine container image for container backend") { |value| options[:image] = value }
  parser.on("--platform PLATFORM", "Container platform for container backend") { |value| options[:platform] = value }
  parser.on("--state-root PATH", "Isolated Runtime state root") { |value| options[:state_root] = value }
  parser.on("--timeout DURATION", "Execution timeout") { |value| options[:timeout] = value }
  parser.on("--bootstrap-timeout DURATION", "Wine prefix bootstrap timeout") { |value| options[:bootstrap_timeout] = value }
  parser.on("--expected-marker MARKER", "Expected stdout marker") { |value| options[:expected_marker] = value }
  parser.on("--runner-arg VALUE", "Argument passed to the compatibility runner before the executable path") { |value| options[:runner_args] << value }
  parser.on("--arg VALUE", "Argument passed to the Windows executable") { |value| options[:app_args] << value }
end.parse!

unless %w[text json markdown].include?(options[:format])
  warn "FAIL: unsupported output format #{options[:format]}"
  exit 1
end
unless %w[local container].include?(options[:backend])
  warn "FAIL: unsupported backend #{options[:backend]}"
  exit 1
end

options[:redact_output] = options[:format] != "text" if options[:redact_output].nil?

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def base_report(format, redact_output, expected_marker, executable_source, backend, image, platform)
  {
    "version" => PROJECT_ROOT.join("VERSION").read.strip,
    "schema_version" => SCHEMA_VERSION,
    "report_type" => "winapp-smoke",
    "format" => format,
    "backend" => backend,
    "redacted_output_requested" => redact_output,
    "executable_source" => executable_source,
    "user_executable_supplied" => executable_source == "user-supplied",
    "fixture_built" => false,
    "runner_diagnostics_invoked" => false,
    "smoke_invoked" => false,
    "status" => "failed",
    "marker" => expected_marker,
    "runner_available" => false,
    "runner_argument_count" => 0,
    "runner_diagnostics_status" => "not-run",
    "env_runner_configured" => false,
    "runner_candidate_count" => 0,
    "runner_diagnostics_next_action" => "",
    "runner_command_hints" => [],
    "wine_bootstrap_attempted" => false,
    "wine_bootstrap_succeeded" => false,
    "wine_bootstrap_exit_code" => -1,
    "marker_observed" => false,
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
    puts "- Runner diagnostics invoked: #{report.fetch("runner_diagnostics_invoked")}"
    puts "- Runner diagnostics status: #{report.fetch("runner_diagnostics_status")}"
    puts "- Runner candidate count: #{report.fetch("runner_candidate_count")}"
    puts "- Env runner configured: #{report.fetch("env_runner_configured")}"
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
    puts "- Wine bootstrap exit code: #{report.fetch("wine_bootstrap_exit_code")}"
    puts "- Marker observed: #{report.fetch("marker_observed")}"
    puts "- Raw output redacted: #{report.fetch("raw_output_redacted")}"
    puts "- KDE-safe output summary: #{report.fetch("kde_safe_output_summary")}"
    puts "- Host root modified: #{report.fetch("host_root_modified")}"
    puts "- Docker executed: #{report.fetch("docker_executed")}"
    puts "- QEMU executed: #{report.fetch("qemu_executed")}"
    puts "- Wine executed by script: #{report.fetch("wine_executed_by_script")}"
    puts "- Network checks run: #{report.fetch("network_checks_run")}"
    puts "- Package manager invoked: #{report.fetch("package_manager_invoked")}"
    puts "- Runner diagnostics next action: #{report.fetch("runner_diagnostics_next_action")}" unless report.fetch("runner_diagnostics_next_action").empty?
    puts "- Failure reason: #{report.fetch("failure_reason")}" unless report.fetch("failure_reason").empty?
    puts "- Skip reason: #{report.fetch("skip_reason")}" unless report.fetch("skip_reason").empty?
  end
end

def finish(report, exit_code)
  emit_report(report) unless report.fetch("format") == "text"
  exit exit_code
end

executable_source = options[:exe].to_s.strip.empty? ? "fixture" : "user-supplied"
report = base_report(
  options.fetch(:format),
  options.fetch(:redact_output),
  options.fetch(:expected_marker),
  executable_source,
  options.fetch(:backend),
  options.fetch(:image),
  options.fetch(:platform)
)

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
  "--expected-marker", options.fetch(:expected_marker)
]
smoke_command.concat(["--runner", options.fetch(:runner)]) unless options[:runner].to_s.strip.empty?
options.fetch(:runner_args).each { |value| smoke_command.concat(["--runner-arg", value]) }
options.fetch(:app_args).each { |value| smoke_command.concat(["--arg", value]) }
smoke_command << "--redact-output" if options.fetch(:redact_output)

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
report["runner_available"] = payload.fetch("runner_available", false)
report["runner_argument_count"] = payload.fetch("runner_argument_count", 0)
report["wine_bootstrap_attempted"] = payload.fetch("wine_bootstrap_attempted", false)
report["wine_bootstrap_succeeded"] = payload.fetch("wine_bootstrap_succeeded", false)
report["wine_bootstrap_exit_code"] = payload.fetch("wine_bootstrap_exit_code", -1)
report["marker_observed"] = payload.fetch("marker_observed", false)
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
