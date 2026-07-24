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
  assert(report.fetch("fixture_built"), "JSON report must record fixture build")
  assert(report.fetch("smoke_invoked"), "JSON report must record smoke invocation")
  assert(report.fetch("redacted_output_requested"), "JSON report must request redacted output by default")
  assert(report.fetch("raw_output_redacted"), "JSON report must preserve redacted output state")
  assert(!report.fetch("raw_output_included"), "JSON report must not include raw output by default")
  assert(report.fetch("marker_observed"), "JSON report must preserve marker observation")
  assert(report.fetch("runtime_payload").fetch("stdout") == "", "JSON runtime payload must omit raw stdout")
  assert(!json_stdout.include?("raw-host-path"), "JSON report must not leak raw runner output")

  markdown_stdout, markdown_stderr, markdown_status = Open3.capture3(env, "ruby", script.to_s, "--format", "markdown")
  assert(markdown_status.success?, "winapp smoke Markdown report must succeed: #{markdown_stderr}")
  assert(markdown_stdout.include?("# Windows App Smoke Report"), "Markdown report must include title")
  assert(markdown_stdout.include?("Status: passed"), "Markdown report must include status")
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
end
