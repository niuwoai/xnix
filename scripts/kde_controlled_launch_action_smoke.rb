#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
SMOKE_NAME = "KDE controlled launch action smoke"
PASS_MARKER = "PASS: KDE controlled launch action smoke"
SKIP_MARKER = "SKIP: KDE controlled launch action smoke"
APP_ID = "7zr"
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "kde-controlled-launch-action-smoke", RUN_ID)
STATE_ROOT = WORK_ROOT.join("state")
CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
EXECUTE_ENV = "XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def command_available?(name)
  ENV.fetch("PATH", "").split(File::PATH_SEPARATOR).any? do |directory|
    path = File.join(directory, name)
    File.file?(path) && File.executable?(path)
  end
end

def runtime_go_command
  return ["xnix-runtime-go"] if command_available?("xnix-runtime-go")
  return ["go", "run", "./cmd/xnix-runtime-go"] if command_available?("go")

  nil
end

def go_env
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
    "GOTMPDIR" => GO_TMP_ROOT.to_s
  }
end

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def run_json(env, *argv)
  stdout, stderr, status = run_command(env, *argv)
  assert(status.zero?, "#{argv.join(" ")} must exit successfully: #{stderr}\n#{stdout}")
  [JSON.parse(stdout), stdout]
rescue JSON::ParserError => e
  warn stdout
  warn "FAIL: #{argv.join(" ")} must emit JSON: #{e.message}"
  exit 1
end

def assert_no_forbidden(text, forbidden_terms, label)
  downcased = text.downcase
  forbidden_terms.each do |term|
    next if term.to_s.empty?

    assert(!downcased.include?(term.to_s.downcase), "#{label} must not expose #{term}")
  end
end

def assert_false_payload(payload, fields, label)
  fields.each do |field|
    assert(payload[field] == false, "#{label} must keep #{field}=false")
  end
end

FileUtils.mkdir_p(STATE_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)

runtime_command = runtime_go_command
unless runtime_command
  puts "#{SKIP_MARKER} (xnix-runtime-go and go are unavailable)"
  exit 0
end

fixture, fixture_stdout = run_json(
  go_env,
  *runtime_command,
  "known-app-runtime-status-launch-owner-fixture-record",
  "--app", APP_ID,
  "--state-root", STATE_ROOT.to_s,
  "--cache-root", CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)

unless fixture["fixture_ready"] == true
  reason = fixture.fetch("skip_reason", "known Windows app artifact unavailable")
  puts "#{SKIP_MARKER} (#{reason})"
  exit 0
end

evidence_relative_path = fixture.fetch("evidence_relative_path")
assert(evidence_relative_path.start_with?("runtime/kde-runtime-status-launch-evidence/"), "fixture must expose only a relative evidence path")
assert(fixture["desktop_trigger_ready"] == true, "fixture must expose a desktop trigger")
assert(fixture["desktop_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "fixture must target the controlled-launch D-Bus method")
assert(fixture["kde_forwards_only_evidence_handle"] == true, "fixture must keep KDE evidence-only")
assert_false_payload(
  fixture,
  %w[state_root_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required privileged_container_required network_required],
  "fixture"
)
assert_no_forbidden(fixture_stdout, [STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "fixture output")

plan, plan_stdout = run_json(
  go_env,
  *runtime_command,
  "kde-controlled-launch-session-bus-smoke-plan-preview",
  "--state-root", STATE_ROOT.to_s,
  "--evidence-relative-path", evidence_relative_path
)

assert(plan["request_type"] == "kde-controlled-launch-session-bus-smoke-plan-preview", "plan request type must match")
assert(plan["kde_action_id"] == "xnix.runtime-status.controlled-launch", "plan must target the KDE controlled-launch action")
assert(plan["kde_action_preview_request_type"] == "kde-controlled-launch-action-preview", "plan must consume the KDE action preview")
assert(plan["public_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "plan must target the controlled-launch D-Bus method")
assert(plan["evidence_relative_path"] == evidence_relative_path, "plan must preserve the fixture evidence path")
assert(plan["evidence_sha256"] == fixture.fetch("evidence_sha256"), "plan must preserve the fixture evidence digest")
assert(plan["evidence_handoff_consumed"] == true, "plan must consume the evidence handoff")
assert(plan["evidence_digest_verified"] == true, "plan must verify the evidence digest")
assert(plan["kde_forwarded_arguments"] == [evidence_relative_path], "plan must forward only the evidence handle")
assert(plan["kde_forwarded_argument_kind"] == "evidence-relative-path", "plan must label the forwarded evidence argument")
assert(plan["kde_forwards_only_evidence_handle"] == true, "plan must keep KDE evidence-only")
assert(plan["restricted_session_bus_plan_ready"] == true, "plan must be ready for the restricted session-bus smoke")
assert(plan["private_session_bus_required"] == true, "plan must require a private session bus")
assert(plan["dbus_session_bus_address_required"] == true, "plan must require DBUS_SESSION_BUS_ADDRESS")
assert(plan["outer_private_session_bus_command"] == ["dbus-run-session", "--", "ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"], "plan must expose the private session-bus runner")
assert(plan["runtime_status_owner_session_smoke_command"] == ["ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"], "plan must expose the Runtime-status owner session smoke command")
assert(plan["inner_staged_launcher_dispatch_command"] == ["ruby", "scripts/staged_launcher_dispatch_smoke.rb"], "plan must expose the staged launcher dispatch smoke command")
assert(plan["container_smoke_command"] == ["ruby", "scripts/container.rb", "runtime-status-owner-service-session-bus-smoke"], "plan must expose the restricted container smoke command")
assert(plan["expected_pass_marker"] == "PASS: Runtime-status owner service session-bus smoke", "plan must expose the pass marker")
assert(plan["expected_skip_marker"] == "SKIP: Runtime-status owner service session-bus smoke", "plan must expose the skip marker")
assert(plan["runtime_owned"] == true, "plan must keep Runtime ownership")
assert(plan["go_runtime_backed"] == true, "plan must be Go backed")
assert_false_payload(
  plan,
  %w[kde_policy_owner owner_service_args_exposed_to_kde desktop_kde_state_root_access desktop_receipt_fields_reconstructed state_root_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required privileged_container_required host_network_required execution_started backend_process_started smoke_executed_by_preview],
  "plan"
)
assert_no_forbidden(plan_stdout, [STATE_ROOT.to_s, "owner_service_call_args", "wine ", "wine/", ".wine", "qemu-system", "program files"], "plan output")

unless ENV.fetch(EXECUTE_ENV, "") == "1"
  puts "#{SKIP_MARKER} (validated Go-owned plan; set #{EXECUTE_ENV}=1 to execute restricted session-bus smoke)"
  exit 0
end

command = plan.fetch("runtime_status_owner_session_smoke_command")
stdout, stderr, status = run_command({}, *command)
print stdout
warn stderr unless stderr.empty?
assert(status.zero?, "#{command.join(" ")} must pass or skip cleanly")
assert_no_forbidden(stdout, [STATE_ROOT.to_s, "docker.sock", "--privileged", "--network host", "type=bind", "program files", ".wine", "qemu-system"], "KDE action smoke stdout")
assert_no_forbidden(stderr, [STATE_ROOT.to_s, "docker.sock", "--privileged", "--network host", "type=bind", "program files", ".wine", "qemu-system"], "KDE action smoke stderr")

case stdout
when /#{Regexp.escape(plan.fetch("expected_pass_marker"))}/
  puts PASS_MARKER
when /#{Regexp.escape(plan.fetch("expected_skip_marker"))}/
  reason = stdout.lines.find { |line| line.include?(plan.fetch("expected_skip_marker")) }.to_s.sub(plan.fetch("expected_skip_marker"), "").strip
  puts "#{SKIP_MARKER} #{reason}".rstrip
else
  warn stdout
  warn stderr unless stderr.empty?
  warn "FAIL: #{SMOKE_NAME} must report a plan-defined pass or skip marker"
  exit 1
end
