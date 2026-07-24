#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "fileutils"
require "open3"
require "pathname"
require "tmpdir"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/winapp_smoke.rb")

Dir.mktmpdir("xnix-winapp-smoke-test") do |dir|
  temp_root = Pathname.new(dir)
  fake_go = temp_root.join("go")
  fake_go.write(<<~'RUBY')
    #!/usr/bin/env ruby
    # frozen_string_literal: true

    require "json"

    args = ARGV.dup
    File.open(ENV.fetch("XNIX_FAKE_GO_LOG"), "a") { |file| file.puts(args.join("\u0001")) } if ENV["XNIX_FAKE_GO_LOG"]
    if args.first == "build"
      output_path = args[args.index("-o") + 1]
      File.write(output_path, "fixture exe")
      exit 0
    end

    if args[0] == "run" && args.include?("windows-app-smoke-profile-preflight")
      profile_path = args[args.index("--profile") + 1]
      profile = JSON.parse(File.read(profile_path))
      payload = {
        "schema_version" => "xnix.runtime.windows_app_smoke_profile_preflight.v1",
        "request_type" => "windows-app-smoke-profile-preflight",
        "status" => "ready",
        "profile_supplied" => true,
        "executable_name" => "custom.exe",
        "executable_exists" => true,
        "executable_format" => "pe-mz",
        "windows_executable_signature_observed" => true,
        "executable_architecture" => "x86_64",
        "executable_architecture_supported" => true,
        "wine_architecture" => "win64",
        "wine_prefix_mode" => "architecture-scoped",
        "wine_prefix_prepared" => false,
        "application_workspace_mode" => profile.fetch("stage_app_dir", false) ? "staged-application-directory" : "direct-executable",
        "stage_app_dir" => profile.fetch("stage_app_dir", false),
        "working_directory_mode" => "operator-supplied",
        "working_directory_valid" => true,
        "state_root_configured" => true,
        "runner_available" => true,
        "runner_argument_count" => 1,
        "app_argument_count" => 1,
        "success_mode" => "marker",
        "skip_bootstrap" => profile.fetch("skip_bootstrap", false),
        "timeout_configured" => true,
        "expected_marker_configured" => true,
        "runner_diagnostics_payload" => {
          "schema_version" => "xnix.runtime.windows_app_runner_diagnostics.v1",
          "request_type" => "windows-app-runner-diagnostics",
          "status" => "passed",
          "runner_available" => true
        },
        "next_action" => "Run profile smoke.",
        "raw_profile_path_exposed" => false,
        "raw_executable_path_exposed" => false,
        "raw_runner_path_exposed" => false,
        "raw_working_directory_exposed" => false,
        "raw_runner_arguments_exposed" => false,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false,
        "docker_executed" => false,
        "qemu_executed" => false,
        "wine_executed" => false,
        "colima_executed" => false,
        "network_checks_run" => false,
        "package_manager_invoked" => false
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("windows-app-runner-diagnostics")
      explicit = args.include?("--runner")
      payload = {
        "schema_version" => "xnix.runtime.windows_app_runner_diagnostics.v1",
        "request_type" => "windows-app-runner-diagnostics",
        "status" => "passed",
        "runner_available" => true,
        "explicit_runner_supplied" => explicit,
        "env_runner_configured" => false,
        "candidate_count" => 1,
        "candidates" => [{
          "id" => explicit ? "explicit-runner" : "path-wine",
          "source" => explicit ? "operator-supplied-runner" : "path-command",
          "available" => true,
          "selected" => true,
          "reason" => "available"
        }],
        "selected_runner_name" => "fake-runner",
        "next_action" => "Run windows-app-run-smoke with the selected runner.",
        "runner_command_hints" => [
          "ruby scripts/winapp_smoke.rb --exe path/to/app.exe --format json",
          "ruby scripts/winapp_smoke.rb --exe path/to/app.exe --runner path/to/wine --format json"
        ],
        "raw_path_exposed" => false,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false,
        "docker_executed" => false,
        "qemu_executed" => false,
        "colima_executed" => false,
        "network_checks_run" => false,
        "package_manager_invoked" => false
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("windows-app-smoke-profile-render")
      marker = args.include?("--expected-marker") ? args[args.index("--expected-marker") + 1] : "XNIX_WINAPP_SMOKE_OK"
      payload = {
        "schema_version" => "xnix.runtime.windows_app_smoke_profile.v1",
        "executable_path" => args[args.index("--exe") + 1],
        "working_directory" => args.include?("--working-dir") ? args[args.index("--working-dir") + 1] : "",
        "runner_path" => args.include?("--runner") ? args[args.index("--runner") + 1] : "",
        "runner_bottle" => args.include?("--runner-bottle") ? args[args.index("--runner-bottle") + 1] : "",
        "runner_arguments" => args.each_with_index.filter_map { |value, index| args[index + 1] if value == "--runner-arg" },
        "arguments" => args.each_with_index.filter_map { |value, index| args[index + 1] if value == "--arg" },
        "state_root" => args[args.index("--state-root") + 1],
        "timeout" => args[args.index("--timeout") + 1],
        "expected_marker" => marker,
        "success_mode" => args[args.index("--success-mode") + 1],
        "redact_output" => args.include?("--redact-output"),
        "skip_bootstrap" => args.include?("--skip-bootstrap"),
        "stage_app_dir" => args.include?("--stage-app-dir")
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("windows-app-run-smoke")
      redacted = args.include?("--redact-output")
      marker = args.include?("--expected-marker") ? args[args.index("--expected-marker") + 1] : "XNIX_WINAPP_SMOKE_OK"
      success_mode = args.include?("--success-mode") ? args[args.index("--success-mode") + 1] : "marker"
      exe_path = args[args.index("--exe") + 1]
      runner_arg_count = args.each_with_index.count { |value, index| value == "--runner-arg" && index + 1 < args.length }
      runner_arg_count += 2 if args.include?("--runner-bottle")
      marker_observed = success_mode == "marker"
      startup_window_observed = success_mode == "startup-window"
      working_directory_mode = args.include?("--working-dir") ? "operator-supplied" : "executable-directory"
      bootstrap_skipped = args.include?("--skip-bootstrap")
      app_staged = args.include?("--stage-app-dir")
      payload = {
        "schema_version" => "xnix.runtime.windows_app_smoke.v1",
        "request_type" => "windows-app-run-smoke",
        "status" => "passed",
        "executable_name" => File.basename(exe_path),
        "executable_format" => "pe-mz",
        "windows_executable_signature_observed" => true,
        "executable_architecture" => "x86_64",
        "executable_architecture_supported" => true,
        "wine_architecture" => "win64",
        "wine_prefix_mode" => "architecture-scoped",
        "wine_prefix_prepared" => true,
        "runner_available" => true,
        "runner_argument_count" => runner_arg_count,
        "compatibility_layer" => "windows-compatibility-layer",
        "wine_bootstrap_attempted" => false,
        "wine_bootstrap_succeeded" => false,
        "wine_bootstrap_skipped" => bootstrap_skipped,
        "wine_bootstrap_exit_code" => -1,
        "expected_marker" => marker,
        "success_mode" => success_mode,
        "marker_observed" => marker_observed,
        "startup_window_observed" => startup_window_observed,
        "working_directory_mode" => app_staged ? "staged-application-workspace" : working_directory_mode,
        "application_workspace_mode" => app_staged ? "staged-application-directory" : "direct-executable",
        "application_staged" => app_staged,
        "application_staged_file_count" => app_staged ? 2 : 0,
        "application_staged_bytes" => app_staged ? 2048 : 0,
        "exit_code" => 0,
        "duration_millis" => 1,
        "stdout" => redacted ? "" : (marker_observed ? "#{marker}\\n" : "GUI app exited cleanly\\n"),
        "stderr" => "",
        "stdout_bytes" => 39,
        "stderr_bytes" => 0,
        "stdout_line_count" => 2,
        "stderr_line_count" => 0,
        "raw_output_included" => !redacted,
        "raw_output_redacted" => redacted,
        "kde_safe_output_summary" => "expected smoke marker observed; stdout_bytes=39 stderr_bytes=0 stdout_lines=2 stderr_lines=0",
        "isolated_state_root" => true,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    if args[0] == "run" && args.include?("windows-app-container-run-smoke")
      exe_path = args[args.index("--exe") + 1]
      image = args[args.index("--image") + 1]
      platform = args[args.index("--platform") + 1]
      payload = {
        "schema_version" => "xnix.runtime.windows_app_container_smoke.v1",
        "request_type" => "windows-app-container-run-smoke",
        "status" => "passed",
        "executable_name" => File.basename(exe_path),
        "container_image" => image,
        "container_platform" => platform,
        "container_state_mode" => "tmpfs",
        "pull_policy" => "never",
        "network_mode" => "none",
        "wine_bootstrap_required" => true,
        "wine_bootstrap_timed_out" => false,
        "wine_bootstrap_exit_code" => -1,
        "runner_available" => true,
        "image_available" => true,
        "compatibility_layer" => "containerized-windows-compatibility-layer",
        "expected_marker" => "XNIX_WINAPP_SMOKE_OK",
        "marker_observed" => true,
        "exit_code" => 0,
        "duration_millis" => 2,
        "stdout" => "XNIX_WINAPP_SMOKE_OK\\n",
        "stderr" => "",
        "isolated_state_root" => true,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "host_networking_required" => false,
        "docker_socket_mounted" => false,
        "broad_host_mount_required" => false,
        "host_mount_count" => 1
      }
      puts JSON.pretty_generate(payload)
      exit 0
    end

    warn "unexpected fake go invocation: #{args.join(" ")}"
    exit 2
  RUBY
  fake_go.chmod(0o700)
  fake_go_log = temp_root.join("fake-go.log")

  env = {
    "PATH" => "#{temp_root}:#{ENV.fetch("PATH")}",
    "XNIX_FAKE_GO_LOG" => fake_go_log.to_s
  }

  stale_fixture_output = project_root.join(".local/xnix/winapp-smoke/hello.exe")
  FileUtils.mkdir_p(stale_fixture_output.dirname)
  stale_fixture_output.write("stale exe")

  json_stdout, json_stderr, json_status = Open3.capture3(env, "ruby", script.to_s, "--format", "json")
  assert(json_status.success?, "winapp smoke JSON report must succeed: #{json_stderr}")
  report = JSON.parse(json_stdout)
  assert(report.fetch("schema_version") == "xnix.runtime.winapp_smoke_report.v1", "JSON report must expose schema")
  assert(report.fetch("report_type") == "winapp-smoke", "JSON report must expose report type")
  assert(!report.fetch("profile_preflight_invoked"), "JSON fixture report must not run profile preflight")
  assert(!report.fetch("profile_written"), "JSON fixture report must not write a profile by default")
  assert(report.fetch("status") == "passed", "JSON report must preserve passed smoke state")
  assert(report.fetch("success_mode") == "marker", "JSON report must default to marker success mode")
  assert(report.fetch("working_directory_mode") == "executable-directory", "JSON report must default to executable directory working mode")
  assert(report.fetch("application_workspace_mode") == "direct-executable", "JSON report must default to direct executable workspace mode")
  assert(!report.fetch("application_staged"), "JSON report must not stage applications by default")
  assert(report.fetch("backend") == "local", "JSON report must default to the local backend")
  assert(report.fetch("fixture_built"), "JSON report must record fixture build")
  assert(report.fetch("executable_format") == "pe-mz", "JSON report must preserve executable format evidence")
  assert(report.fetch("windows_executable_signature_observed"), "JSON report must preserve Windows signature evidence")
  assert(report.fetch("executable_architecture") == "x86_64", "JSON report must preserve executable architecture")
  assert(report.fetch("executable_architecture_supported"), "JSON report must preserve executable architecture support")
  assert(report.fetch("wine_architecture") == "win64", "JSON report must preserve Wine architecture")
  assert(report.fetch("wine_prefix_mode") == "architecture-scoped", "JSON report must preserve Wine prefix mode")
  assert(report.fetch("wine_prefix_prepared"), "JSON report must preserve Wine prefix preparation state")
  assert(report.fetch("runner_diagnostics_invoked"), "JSON report must record runner diagnostics invocation")
  assert(report.fetch("runner_diagnostics_status") == "passed", "JSON report must preserve diagnostics status")
  assert(report.fetch("runner_candidate_count") == 1, "JSON report must preserve runner candidate count")
  assert(report.fetch("runner_command_hints").any? { |hint| hint.include?("path/to/app.exe") }, "JSON report must expose safe runner command hints")
  assert(report.fetch("runner_diagnostics_payload").fetch("request_type") == "windows-app-runner-diagnostics", "JSON report must embed diagnostics payload")
  assert(report.fetch("smoke_invoked"), "JSON report must record smoke invocation")
  assert(!report.fetch("wine_bootstrap_attempted"), "JSON report must preserve bootstrap attempted state")
  assert(!report.fetch("wine_bootstrap_skipped"), "JSON report must preserve bootstrap skipped state")
  assert(report.fetch("redacted_output_requested"), "JSON report must request redacted output by default")
  assert(report.fetch("raw_output_redacted"), "JSON report must preserve redacted output state")
  assert(!report.fetch("raw_output_included"), "JSON report must not include raw output by default")
  assert(report.fetch("marker_observed"), "JSON report must preserve marker observation")
  assert(!report.fetch("startup_window_observed"), "JSON report must preserve startup window state")
  assert(report.fetch("runtime_payload").fetch("stdout") == "", "JSON runtime payload must omit raw stdout")
  assert(!json_stdout.include?("raw-host-path"), "JSON report must not leak raw runner output")

  markdown_stdout, markdown_stderr, markdown_status = Open3.capture3(env, "ruby", script.to_s, "--format", "markdown")
  assert(markdown_status.success?, "winapp smoke Markdown report must succeed: #{markdown_stderr}")
  assert(markdown_stdout.include?("# Windows App Smoke Report"), "Markdown report must include title")
  assert(markdown_stdout.include?("Backend: local"), "Markdown report must include backend")
  assert(markdown_stdout.include?("Status: passed"), "Markdown report must include status")
  assert(markdown_stdout.include?("Executable format: pe-mz"), "Markdown report must expose executable format")
  assert(markdown_stdout.include?("Windows executable signature observed: true"), "Markdown report must expose executable signature evidence")
  assert(markdown_stdout.include?("Executable architecture: x86_64"), "Markdown report must expose executable architecture")
  assert(markdown_stdout.include?("Executable architecture supported: true"), "Markdown report must expose executable architecture support")
  assert(markdown_stdout.include?("Wine architecture: win64"), "Markdown report must expose Wine architecture")
  assert(markdown_stdout.include?("Wine prefix mode: architecture-scoped"), "Markdown report must expose Wine prefix mode")
  assert(markdown_stdout.include?("Wine prefix prepared: true"), "Markdown report must expose Wine prefix preparation state")
  assert(markdown_stdout.include?("Runner diagnostics invoked: true"), "Markdown report must expose runner diagnostics invocation")
  assert(markdown_stdout.include?("Runner candidate count: 1"), "Markdown report must expose runner candidate count")
  assert(markdown_stdout.include?("Profile preflight invoked: false"), "Markdown report must expose profile preflight invocation")
  assert(markdown_stdout.include?("Success mode: marker"), "Markdown report must expose success mode")
  assert(markdown_stdout.include?("Working directory mode: executable-directory"), "Markdown report must expose working directory mode")
  assert(markdown_stdout.include?("Application workspace mode: direct-executable"), "Markdown report must expose application workspace mode")
  assert(markdown_stdout.include?("Application staged: false"), "Markdown report must expose application staging state")
  assert(markdown_stdout.include?("Runner argument count: 0"), "Markdown report must expose runner argument count")
  assert(markdown_stdout.include?("Runner command hints:"), "Markdown report must expose runner command hints")
  assert(markdown_stdout.include?("ruby scripts/winapp_smoke.rb --exe path/to/app.exe"), "Markdown report must include safe smoke command hint")
  assert(markdown_stdout.include?("Wine bootstrap attempted: false"), "Markdown report must expose bootstrap attempted state")
  assert(markdown_stdout.include?("Wine bootstrap skipped: false"), "Markdown report must expose bootstrap skipped state")
  assert(markdown_stdout.include?("Startup window observed: false"), "Markdown report must expose startup window state")
  assert(markdown_stdout.include?("Raw output redacted: true"), "Markdown report must expose redaction")
  assert(markdown_stdout.include?("Wine executed by script: false"), "Markdown report must keep Wine execution-by-script false")

  custom_exe = temp_root.join("custom.exe")
  custom_exe.write("custom fixture")
  custom_runner = temp_root.join("custom-runner")
  custom_runner.write("runner")
  custom_working_dir = temp_root.join("custom-working-dir")
  custom_working_dir.mkdir
  custom_stdout, custom_stderr, custom_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--exe", custom_exe.to_s,
    "--runner", custom_runner.to_s,
    "--working-dir", custom_working_dir.to_s,
    "--runner-bottle", "private-bottle-name",
    "--runner-arg", "--shim-mode",
    "--expected-marker", "CUSTOM_APP_OK",
    "--skip-bootstrap",
    "--stage-app-dir",
    "--arg", "--custom-flag"
  )
  assert(custom_status.success?, "winapp smoke custom executable JSON report must succeed: #{custom_stderr}")
  custom_report = JSON.parse(custom_stdout)
  assert(custom_report.fetch("executable_source") == "user-supplied", "custom report must identify user-supplied source")
  assert(custom_report.fetch("user_executable_supplied"), "custom report must mark user executable supplied")
  assert(!custom_report.fetch("fixture_built"), "custom report must skip fixture build")
  assert(custom_report.fetch("runner_diagnostics_payload").fetch("explicit_runner_supplied"), "custom report must pass runner into diagnostics")
  assert(custom_report.fetch("executable_format") == "pe-mz", "custom report must preserve executable format evidence")
  assert(custom_report.fetch("windows_executable_signature_observed"), "custom report must preserve Windows signature evidence")
  assert(custom_report.fetch("executable_architecture") == "x86_64", "custom report must preserve executable architecture")
  assert(custom_report.fetch("executable_architecture_supported"), "custom report must preserve executable architecture support")
  assert(custom_report.fetch("wine_architecture") == "win64", "custom report must preserve Wine architecture")
  assert(custom_report.fetch("wine_prefix_mode") == "architecture-scoped", "custom report must preserve Wine prefix mode")
  assert(custom_report.fetch("wine_prefix_prepared"), "custom report must preserve Wine prefix preparation state")
  assert(custom_report.fetch("marker") == "CUSTOM_APP_OK", "custom report must preserve custom marker")
  assert(custom_report.fetch("success_mode") == "marker", "custom report must preserve default success mode")
  assert(custom_report.fetch("working_directory_mode") == "staged-application-workspace", "custom report must preserve staged working directory mode")
  assert(custom_report.fetch("application_workspace_mode") == "staged-application-directory", "custom report must preserve staged workspace mode")
  assert(custom_report.fetch("application_staged"), "custom report must preserve application staging")
  assert(custom_report.fetch("application_staged_file_count") == 2, "custom report must preserve staged file count")
  assert(custom_report.fetch("runtime_payload").fetch("expected_marker") == "CUSTOM_APP_OK", "custom report must pass custom marker to Runtime")
  assert(custom_report.fetch("runtime_payload").fetch("executable_name") == "custom.exe", "custom report must preserve safe executable basename")
  assert(custom_report.fetch("wine_bootstrap_skipped"), "custom report must preserve bootstrap skip")
  assert(custom_report.fetch("runtime_payload").fetch("wine_bootstrap_skipped"), "custom report must pass bootstrap skip to Runtime")
  assert(custom_report.fetch("runner_argument_count") == 3, "custom report must preserve runner argument count")
  assert(!custom_stdout.include?(custom_exe.to_s), "custom report must not leak executable path")
  assert(!custom_stdout.include?(custom_runner.to_s), "custom report must not leak runner path")
  assert(!custom_stdout.include?(custom_working_dir.to_s), "custom report must not leak working directory path")
  assert(!custom_stdout.include?("private-bottle-name"), "custom report must not leak runner arguments")
  assert(!custom_stdout.include?("--shim-mode"), "custom report must not leak raw runner arguments")
  assert(!custom_report.fetch("runner_command_hints").join("\n").include?(custom_runner.to_s), "custom command hints must not leak runner path")

  custom_invocations = fake_go_log.read.lines.map { |line| line.split("\u0001").map(&:chomp) }
  last_invocation = custom_invocations.last
  assert(!last_invocation.include?("build"), "custom executable mode must not build the fixture")
  assert(last_invocation.include?("--runner"), "custom executable mode must forward explicit runner")
  assert(last_invocation.include?("--working-dir"), "custom executable mode must forward working directory")
  assert(last_invocation.include?("--runner-bottle"), "custom executable mode must forward runner bottle")
  assert(last_invocation.include?("--runner-arg"), "custom executable mode must forward runner arguments")
  assert(last_invocation.include?("--skip-bootstrap"), "custom executable mode must forward bootstrap skip")
  assert(last_invocation.include?("--stage-app-dir"), "custom executable mode must forward application staging")
  assert(last_invocation.include?("--custom-flag"), "custom executable mode must forward app arguments")

  generated_profile = temp_root.join("generated.profile.json")
  write_profile_stdout, write_profile_stderr, write_profile_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--exe", custom_exe.to_s,
    "--runner", custom_runner.to_s,
    "--working-dir", custom_working_dir.to_s,
    "--runner-bottle", "private-bottle-name",
    "--runner-arg", "--shim-mode",
    "--expected-marker", "CUSTOM_APP_OK",
    "--skip-bootstrap",
    "--stage-app-dir",
    "--arg", "--custom-flag",
    "--write-profile", generated_profile.to_s
  )
  assert(write_profile_status.success?, "winapp smoke profile write report must succeed: #{write_profile_stderr}")
  write_profile_report = JSON.parse(write_profile_stdout)
  assert(write_profile_report.fetch("profile_write_invoked"), "profile write report must invoke profile rendering")
  assert(write_profile_report.fetch("profile_written"), "profile write report must record profile write")
  rendered_profile = JSON.parse(generated_profile.read)
  assert(rendered_profile.fetch("schema_version") == "xnix.runtime.windows_app_smoke_profile.v1", "written profile must preserve schema")
  assert(rendered_profile.fetch("expected_marker") == "CUSTOM_APP_OK", "written profile must preserve marker")
  assert(rendered_profile.fetch("skip_bootstrap"), "written profile must preserve bootstrap skip")
  assert(rendered_profile.fetch("stage_app_dir"), "written profile must preserve application staging")
  assert(rendered_profile.fetch("runner_arguments") == ["--shim-mode"], "written profile must preserve runner args")
  assert(rendered_profile.fetch("arguments") == ["--custom-flag"], "written profile must preserve app args")
  assert(!write_profile_stdout.include?(generated_profile.to_s), "profile write report must not leak written profile path")

  profile_path = temp_root.join("winapp-profile.json")
  profile_working_dir = temp_root.join("profile-working-dir")
  profile_working_dir.mkdir
  File.write(
    profile_path,
    JSON.pretty_generate(
      {
        "schema_version" => "xnix.runtime.windows_app_smoke_profile.v1",
        "executable_path" => custom_exe.to_s,
        "working_directory" => profile_working_dir.to_s,
        "runner_path" => custom_runner.to_s,
        "runner_arguments" => ["--profile-shim"],
        "arguments" => ["--profile-flag"],
        "state_root" => temp_root.join("profile-state").to_s,
        "expected_marker" => "PROFILE_APP_OK",
        "success_mode" => "marker",
        "skip_bootstrap" => true,
        "stage_app_dir" => true,
        "timeout" => "7s"
      }
    )
  )
  profile_stdout, profile_stderr, profile_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--profile", profile_path.to_s
  )
  assert(profile_status.success?, "winapp smoke profile JSON report must succeed: #{profile_stderr}")
  profile_report = JSON.parse(profile_stdout)
  assert(profile_report.fetch("profile_supplied"), "profile report must mark profile supplied")
  assert(profile_report.fetch("profile_preflight_invoked"), "profile report must run preflight before smoke")
  assert(profile_report.fetch("profile_preflight_status") == "ready", "profile report must preserve preflight readiness")
  assert(profile_report.fetch("profile_preflight_payload").fetch("request_type") == "windows-app-smoke-profile-preflight", "profile report must embed preflight payload")
  assert(profile_report.fetch("profile_preflight_payload").fetch("executable_format") == "pe-mz", "profile report must preserve preflight executable format")
  assert(profile_report.fetch("profile_preflight_payload").fetch("windows_executable_signature_observed"), "profile report must preserve Windows executable signature evidence")
  assert(profile_report.fetch("profile_preflight_payload").fetch("executable_architecture") == "x86_64", "profile report must preserve preflight executable architecture")
  assert(profile_report.fetch("profile_preflight_payload").fetch("executable_architecture_supported"), "profile report must preserve preflight executable architecture support")
  assert(profile_report.fetch("profile_preflight_payload").fetch("wine_architecture") == "win64", "profile report must preserve preflight Wine architecture")
  assert(profile_report.fetch("profile_preflight_payload").fetch("wine_prefix_mode") == "architecture-scoped", "profile report must preserve preflight Wine prefix mode")
  assert(!profile_report.fetch("profile_preflight_payload").fetch("wine_prefix_prepared"), "profile report must preserve preflight Wine prefix preparation state")
  assert(profile_report.fetch("profile_preflight_payload").fetch("skip_bootstrap"), "profile report must preserve preflight bootstrap skip")
  assert(profile_report.fetch("profile_preflight_payload").fetch("stage_app_dir"), "profile report must preserve preflight application staging")
  assert(profile_report.fetch("executable_format") == "pe-mz", "profile report must preserve report-level executable format")
  assert(profile_report.fetch("windows_executable_signature_observed"), "profile report must preserve report-level Windows signature evidence")
  assert(profile_report.fetch("executable_architecture") == "x86_64", "profile report must preserve report-level executable architecture")
  assert(profile_report.fetch("executable_architecture_supported"), "profile report must preserve report-level executable architecture support")
  assert(profile_report.fetch("wine_architecture") == "win64", "profile report must preserve report-level Wine architecture")
  assert(profile_report.fetch("wine_prefix_mode") == "architecture-scoped", "profile report must preserve report-level Wine prefix mode")
  assert(profile_report.fetch("wine_prefix_prepared"), "profile report must preserve report-level Wine prefix preparation state")
  assert(profile_report.fetch("executable_source") == "profile", "profile report must identify profile executable source")
  assert(profile_report.fetch("user_executable_supplied"), "profile report must treat profile executables as user supplied")
  assert(!profile_report.fetch("fixture_built"), "profile report must skip fixture build")
  assert(profile_report.fetch("marker") == "PROFILE_APP_OK", "profile report must preserve profile marker")
  assert(profile_report.fetch("working_directory_mode") == "staged-application-workspace", "profile report must preserve profile staged working directory mode")
  assert(profile_report.fetch("application_workspace_mode") == "staged-application-directory", "profile report must preserve profile staged workspace mode")
  assert(profile_report.fetch("application_staged"), "profile report must preserve profile application staging")
  assert(profile_report.fetch("wine_bootstrap_skipped"), "profile report must preserve profile bootstrap skip")
  assert(profile_report.fetch("runner_argument_count") == 1, "profile report must preserve profile runner argument count")
  assert(!profile_stdout.include?(profile_path.to_s), "profile report must not leak profile path")
  assert(!profile_stdout.include?(temp_root.join("profile-state").to_s), "profile report must not leak profile state root")
  assert(!profile_stdout.include?(profile_working_dir.to_s), "profile report must not leak profile working directory")
  assert(!profile_stdout.include?("--profile-shim"), "profile report must not leak profile runner arguments")
  profile_invocation = fake_go_log.read.lines.map { |line| line.split("\u0001").map(&:chomp) }.last
  assert(profile_invocation.include?("--runner"), "profile mode must forward profile runner")
  assert(profile_invocation.include?("--working-dir"), "profile mode must forward profile working directory")
  assert(profile_invocation.include?("--skip-bootstrap"), "profile mode must forward profile bootstrap skip")
  assert(profile_invocation.include?("--stage-app-dir"), "profile mode must forward profile application staging")
  assert(profile_invocation.include?("--profile-flag"), "profile mode must forward profile app arguments")

  preflight_stdout, preflight_stderr, preflight_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--profile", profile_path.to_s,
    "--preflight-only"
  )
  assert(preflight_status.success?, "winapp smoke preflight-only report must succeed: #{preflight_stderr}")
  preflight_report = JSON.parse(preflight_stdout)
  assert(preflight_report.fetch("status") == "ready", "preflight-only report must expose ready status")
  assert(preflight_report.fetch("preflight_only"), "preflight-only report must mark preflight-only mode")
  assert(preflight_report.fetch("profile_preflight_invoked"), "preflight-only report must invoke profile preflight")
  assert(!preflight_report.fetch("smoke_invoked"), "preflight-only report must not invoke smoke")
  assert(preflight_report.fetch("runtime_payload").nil?, "preflight-only report must not embed runtime smoke payload")
  assert(!preflight_stdout.include?(profile_path.to_s), "preflight-only report must not leak profile path")

  exit_code_stdout, exit_code_stderr, exit_code_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--exe", custom_exe.to_s,
    "--runner", custom_runner.to_s,
    "--success-mode", "exit-code"
  )
  assert(exit_code_status.success?, "winapp smoke exit-code success mode report must succeed: #{exit_code_stderr}")
  exit_code_report = JSON.parse(exit_code_stdout)
  assert(exit_code_report.fetch("status") == "passed", "exit-code report must pass on zero runner exit")
  assert(exit_code_report.fetch("success_mode") == "exit-code", "exit-code report must preserve success mode")
  assert(!exit_code_report.fetch("marker_observed"), "exit-code report must not require marker observation")
  assert(!exit_code_stdout.include?(custom_exe.to_s), "exit-code report must not leak executable path")

  startup_stdout, startup_stderr, startup_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--exe", custom_exe.to_s,
    "--runner", custom_runner.to_s,
    "--success-mode", "startup-window"
  )
  assert(startup_status.success?, "winapp smoke startup-window success mode report must succeed: #{startup_stderr}")
  startup_report = JSON.parse(startup_stdout)
  assert(startup_report.fetch("status") == "passed", "startup-window report must pass on startup window observation")
  assert(startup_report.fetch("success_mode") == "startup-window", "startup-window report must preserve success mode")
  assert(startup_report.fetch("startup_window_observed"), "startup-window report must preserve startup observation")
  assert(!startup_report.fetch("marker_observed"), "startup-window report must not require marker observation")
  assert(!startup_stdout.include?(custom_exe.to_s), "startup-window report must not leak executable path")

  container_stdout, container_stderr, container_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--backend", "container",
    "--image", "local/wine-smoke:test",
    "--platform", "linux/amd64"
  )
  assert(container_status.success?, "winapp smoke container backend JSON report must succeed: #{container_stderr}")
  container_report = JSON.parse(container_stdout)
  assert(container_report.fetch("backend") == "container", "container report must record container backend")
  assert(container_report.fetch("container_smoke_invoked"), "container report must record container smoke invocation")
  assert(container_report.fetch("container_image") == "local/wine-smoke:test", "container report must preserve image name")
  assert(container_report.fetch("container_image_available"), "container report must preserve image availability")
  assert(container_report.fetch("docker_executed"), "container report must record Docker execution evidence")
  assert(container_report.fetch("wine_bootstrap_attempted"), "container report must expose bootstrap attempt")
  assert(container_report.fetch("wine_bootstrap_succeeded"), "container report must expose bootstrap success")
  assert(container_report.fetch("container_payload").fetch("request_type") == "windows-app-container-run-smoke", "container report must embed container payload")
  assert(!container_stdout.include?(stale_fixture_output.to_s), "container report must not leak fixture executable host path")
end
