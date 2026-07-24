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

    if args[0] == "run" && args.include?("windows-app-run-smoke")
      redacted = args.include?("--redact-output")
      marker = args.include?("--expected-marker") ? args[args.index("--expected-marker") + 1] : "XNIX_WINAPP_SMOKE_OK"
      exe_path = args[args.index("--exe") + 1]
      payload = {
        "schema_version" => "xnix.runtime.windows_app_smoke.v1",
        "request_type" => "windows-app-run-smoke",
        "status" => "passed",
        "executable_name" => File.basename(exe_path),
        "runner_available" => true,
        "compatibility_layer" => "windows-compatibility-layer",
        "wine_bootstrap_attempted" => false,
        "wine_bootstrap_succeeded" => false,
        "wine_bootstrap_exit_code" => -1,
        "expected_marker" => marker,
        "marker_observed" => true,
        "exit_code" => 0,
        "duration_millis" => 1,
        "stdout" => redacted ? "" : "#{marker}\\n",
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
  assert(report.fetch("status") == "passed", "JSON report must preserve passed smoke state")
  assert(report.fetch("backend") == "local", "JSON report must default to the local backend")
  assert(report.fetch("fixture_built"), "JSON report must record fixture build")
  assert(report.fetch("runner_diagnostics_invoked"), "JSON report must record runner diagnostics invocation")
  assert(report.fetch("runner_diagnostics_status") == "passed", "JSON report must preserve diagnostics status")
  assert(report.fetch("runner_candidate_count") == 1, "JSON report must preserve runner candidate count")
  assert(report.fetch("runner_diagnostics_payload").fetch("request_type") == "windows-app-runner-diagnostics", "JSON report must embed diagnostics payload")
  assert(report.fetch("smoke_invoked"), "JSON report must record smoke invocation")
  assert(!report.fetch("wine_bootstrap_attempted"), "JSON report must preserve bootstrap attempted state")
  assert(report.fetch("redacted_output_requested"), "JSON report must request redacted output by default")
  assert(report.fetch("raw_output_redacted"), "JSON report must preserve redacted output state")
  assert(!report.fetch("raw_output_included"), "JSON report must not include raw output by default")
  assert(report.fetch("marker_observed"), "JSON report must preserve marker observation")
  assert(report.fetch("runtime_payload").fetch("stdout") == "", "JSON runtime payload must omit raw stdout")
  assert(!json_stdout.include?("raw-host-path"), "JSON report must not leak raw runner output")

  markdown_stdout, markdown_stderr, markdown_status = Open3.capture3(env, "ruby", script.to_s, "--format", "markdown")
  assert(markdown_status.success?, "winapp smoke Markdown report must succeed: #{markdown_stderr}")
  assert(markdown_stdout.include?("# Windows App Smoke Report"), "Markdown report must include title")
  assert(markdown_stdout.include?("Backend: local"), "Markdown report must include backend")
  assert(markdown_stdout.include?("Status: passed"), "Markdown report must include status")
  assert(markdown_stdout.include?("Runner diagnostics invoked: true"), "Markdown report must expose runner diagnostics invocation")
  assert(markdown_stdout.include?("Runner candidate count: 1"), "Markdown report must expose runner candidate count")
  assert(markdown_stdout.include?("Wine bootstrap attempted: false"), "Markdown report must expose bootstrap attempted state")
  assert(markdown_stdout.include?("Raw output redacted: true"), "Markdown report must expose redaction")
  assert(markdown_stdout.include?("Wine executed by script: false"), "Markdown report must keep Wine execution-by-script false")

  custom_exe = temp_root.join("custom.exe")
  custom_exe.write("custom fixture")
  custom_runner = temp_root.join("custom-runner")
  custom_runner.write("runner")
  custom_stdout, custom_stderr, custom_status = Open3.capture3(
    env,
    "ruby", script.to_s,
    "--format", "json",
    "--exe", custom_exe.to_s,
    "--runner", custom_runner.to_s,
    "--expected-marker", "CUSTOM_APP_OK",
    "--arg", "--custom-flag"
  )
  assert(custom_status.success?, "winapp smoke custom executable JSON report must succeed: #{custom_stderr}")
  custom_report = JSON.parse(custom_stdout)
  assert(custom_report.fetch("executable_source") == "user-supplied", "custom report must identify user-supplied source")
  assert(custom_report.fetch("user_executable_supplied"), "custom report must mark user executable supplied")
  assert(!custom_report.fetch("fixture_built"), "custom report must skip fixture build")
  assert(custom_report.fetch("runner_diagnostics_payload").fetch("explicit_runner_supplied"), "custom report must pass runner into diagnostics")
  assert(custom_report.fetch("marker") == "CUSTOM_APP_OK", "custom report must preserve custom marker")
  assert(custom_report.fetch("runtime_payload").fetch("expected_marker") == "CUSTOM_APP_OK", "custom report must pass custom marker to Runtime")
  assert(custom_report.fetch("runtime_payload").fetch("executable_name") == "custom.exe", "custom report must preserve safe executable basename")
  assert(!custom_stdout.include?(custom_exe.to_s), "custom report must not leak executable path")
  assert(!custom_stdout.include?(custom_runner.to_s), "custom report must not leak runner path")

  custom_invocations = fake_go_log.read.lines.map { |line| line.split("\u0001") }
  last_invocation = custom_invocations.last
  assert(!last_invocation.include?("build"), "custom executable mode must not build the fixture")
  assert(last_invocation.include?("--runner"), "custom executable mode must forward explicit runner")
  assert(last_invocation.include?("--custom-flag"), "custom executable mode must forward app arguments")

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
