#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
BUS_NAME = "org.xnix.Compatibility1"
OBJECT_PATH = "/org/xnix/Compatibility1"
INTERFACE = "org.xnix.Compatibility1"
PUBLIC_METHOD = "org.xnix.Compatibility1.ShowRuntimeControlledLaunch"
APP_ID = "7zr"
APP_VERSION = "26.02"
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
SESSION_ID = "known-app-controlled-execution-session-7zr-26.02"
LAUNCH_RECEIPT_ID = "known-app-launch-authorization-7zr-26.02"
REVIEW_RECEIPT_ID = "known-app-session-gated-launch-review-7zr-26.02-#{SESSION_ID}"
SMOKE_NAME = "D-Bus controlled launch owner fixture smoke"
PASS_MARKER = "PASS: #{SMOKE_NAME}"
SKIP_MARKER = "SKIP: #{SMOKE_NAME}"
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
CONTAINER_SCRATCH_ROOT = "/workspace/.xnix-dbus-controlled-launch-scratch"
DEFAULT_WORK_ROOT = if Dir.exist?(CONTAINER_SCRATCH_ROOT)
                      CONTAINER_SCRATCH_ROOT
                    elsif Dir.exist?("/dev/shm")
                      "/dev/shm/xnix-dbus-controlled-launch-owner-fixture-smoke"
                    else
                      PROJECT_ROOT.join(".cache", "xnix", "dbus-controlled-launch-owner-fixture-smoke").to_s
                    end
WORK_ROOT = Pathname.new(ENV.fetch("XNIX_DBUS_CONTROLLED_LAUNCH_WORK_ROOT", DEFAULT_WORK_ROOT))
RUN_ROOT = WORK_ROOT.join(RUN_ID)
STATE_ROOT = RUN_ROOT.join("state")
CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
FAKE_LAUNCHER = RUN_ROOT.join("fake-xnix-compat-launch")
FAKE_LAUNCHER_ARGS = RUN_ROOT.join("fake-launcher-args.txt")

DELEGATED_PAYLOAD = {
  "request_type" => "windows-known-app-dispatch-smoke",
  "status" => "passed",
  "guest_boundary" => GUEST_BOUNDARY,
  "runtime_owned_dispatch" => true,
  "artifact_verified" => true,
  "marker_observed" => true,
  "smoke_passed" => true,
  "execution_started" => true,
  "backend_process_started" => false,
  "session_gated_controlled_dispatch_consumed" => true,
  "session_gated_controlled_dispatch_state" => "created-after-session-gated-review",
  "session_gated_review_receipt_id" => REVIEW_RECEIPT_ID,
  "launch_authorization_receipt_id" => LAUNCH_RECEIPT_ID,
  "controlled_execution_session_consumed" => true,
  "controlled_execution_session_id" => SESSION_ID,
  "controlled_session_digest_verified" => true,
  "controlled_session_relative_path" => "execution-ledger/sessions/#{SESSION_ID}.json",
  "runtime_owner_consumable_session" => true,
  "kde_read_model_consumable_session" => true,
  "controlled_session_live_state_observed" => false,
  "controlled_session_registered" => false,
  "controlled_session_window_observed" => false,
  "controlled_session_host_root_modified" => false,
  "controlled_session_backend_process_start" => false,
  "host_root_modified" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "raw_command_exposed" => false,
  "backend_details_exposed" => false
}.freeze

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

def variant_string_field(stdout, key)
  marker = "'#{key}': <'"
  start = stdout.index(marker)
  assert(start, "D-Bus response must include #{key}")
  value_start = start + marker.length
  value_end = stdout.index("'>", value_start)
  assert(value_end, "D-Bus response must terminate #{key}")
  stdout[value_start...value_end].gsub("\\n", "\n")
end

def assert_no_forbidden(text, forbidden_terms, label)
  downcased = text.downcase
  forbidden_terms.each do |term|
    next if term.to_s.empty?

    assert(!downcased.include?(term.to_s.downcase), "#{label} must not expose #{term}")
  end
end

def write_fake_launcher
  source = [
    "#!/usr/bin/env ruby",
    "# frozen_string_literal: true",
    "require \"json\"",
    "args_path = ENV.fetch(\"XNIX_DBUS_CONTROLLED_LAUNCH_ARGS_FILE\")",
    "File.write(args_path, ARGV.join(\"\\n\") + \"\\n\")",
    "puts JSON.generate(#{DELEGATED_PAYLOAD.inspect})"
  ].join("\n")
  File.write(FAKE_LAUNCHER, source)
  File.chmod(0o755, FAKE_LAUNCHER)
end

def prepare_fixture(runtime_command)
  env = {}
  fixture, fixture_stdout = run_json(
    env,
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
  assert(fixture["launch_authorization_receipt_id"] == LAUNCH_RECEIPT_ID, "launch receipt fixture id must be stable")
  assert(fixture["controlled_execution_session_id"] == SESSION_ID, "controlled session fixture id must be stable")
  assert(fixture["session_gated_review_receipt_id"] == REVIEW_RECEIPT_ID, "review receipt fixture id must be stable")
  assert(fixture["evidence_relative_path"].to_s.start_with?("runtime/kde-runtime-status-launch-evidence/"), "Runtime-status evidence fixture must expose only a relative path")
  assert(fixture["runtime_owned"] == true, "Runtime-status evidence fixture must be Runtime-owned")
  assert(fixture["go_runtime_backed"] == true, "Runtime-status evidence fixture must be Go backed")
  assert(fixture["kde_forwards_only_evidence_handle"] == true, "Runtime-status evidence fixture must keep KDE evidence-only")
  assert(fixture["desktop_receipt_fields_reconstructed"] == false, "Runtime-status evidence fixture must not reconstruct receipt fields in KDE")
  assert(fixture["desktop_kde_state_root_access"] == false, "Runtime-status evidence fixture must not grant KDE state-root access")
  assert_no_forbidden(fixture_stdout, [STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime-status owner fixture output")
  fixture
end

def assert_owner_service_call(stdout, evidence_record)
  payload = variant_string_field(stdout, "go_owner_service_call_json")
  call = JSON.parse(payload)
  assert(call["schema_version"] == "xnix.runtime.owner_service_call.v1", "D-Bus owner call must expose the owner service-call schema")
  assert(call["request_type"] == "runtime-owner-service-call", "D-Bus owner call must expose the owner service-call request type")
  assert(call["service_type"] == "go-runtime-owner-in-process-service", "D-Bus owner call must route through the Go owner service")
  assert(call["method"] == "ShowRuntimeControlledLaunch", "D-Bus owner call must preserve ShowRuntimeControlledLaunch")
  assert(call["call_type"] == "desktop-action-dispatch", "D-Bus owner call must be a desktop action dispatch")
  assert(call["read_only_dispatch"] == false, "D-Bus owner call must not pretend the launch action is read-only")
  assert(call["write_method"] == true, "D-Bus owner call must keep launch action guarded as write-like")
  assert(call["write_methods_enabled"] == false, "D-Bus owner call must keep generic write methods disabled")
  assert(call["dispatch_ready"] == true, "D-Bus owner call must report dispatch readiness")
  assert(call["kde_policy_owner"] == false, "D-Bus owner call must not make KDE the policy owner")
  assert(call["host_root_modified"] == false, "D-Bus owner call must not mutate the host root")
  assert(call["backend_details_exposed"] == false, "D-Bus owner call must not expose backend details")

  action = call.fetch("payload")
  {
    "owner_request_type" => "runtime-owner-show-runtime-controlled-launch",
    "owner_runtime_method" => "ShowRuntimeControlledLaunch",
    "desktop_callable_runtime_method" => "ShowRuntimeControlledLaunch",
    "desktop_callable_route" => "kde-dbus-runtime-status-action",
    "request_type" => "known-app-kde-runtime-status-launch-execution",
    "action_trigger_type" => "known-app-kde-runtime-status-launch-action-trigger-preview",
    "delegated_status" => "passed",
    "delegated_guest_boundary" => GUEST_BOUNDARY,
    "delegated_session_gated_review_receipt_id" => REVIEW_RECEIPT_ID,
    "delegated_launch_authorization_receipt_id" => LAUNCH_RECEIPT_ID,
    "delegated_controlled_execution_session_id" => SESSION_ID
  }.each do |key, expected|
    assert(action[key] == expected, "D-Bus owner action payload must preserve #{key}")
  end

  %w[
    runtime_owner_service_action_dispatch
    runtime_owner_service_call_ready
    runtime_owner_service_supplies_owner_inputs
    kde_forwards_only_evidence_handle
    desktop_evidence_handle_forwarded
    desktop_runtime_owner_adapter_used
    desktop_state_root_supplied_by_runtime_owner
    desktop_cache_root_supplied_by_runtime_owner
    desktop_launcher_supplied_by_runtime_owner
    desktop_timeout_supplied_by_runtime_owner
    evidence_handoff_consumed
    evidence_digest_verified
    managed_launcher_invoked
    existing_managed_launcher_invoked
    launcher_output_json_observed
    delegated_runtime_owned_dispatch
    delegated_artifact_verified
    delegated_marker_observed
    delegated_smoke_passed
    delegated_execution_started
    delegated_session_gated_controlled_dispatch_consumed
    delegated_controlled_execution_session_consumed
    delegated_controlled_session_digest_verified
  ].each do |field|
    assert(action[field] == true, "D-Bus owner action payload must set #{field}=true")
  end

  %w[
    desktop_receipt_fields_reconstructed
    desktop_kde_state_root_access
    kde_state_root_access
    direct_launch_enabled
    desktop_launch_enabled
    backend_launch_enabled
    backend_process_started
    raw_launcher_output_exposed
    backend_details_exposed
    host_root_modified
    docker_socket_mounted
    broad_host_mount_required
    delegated_backend_process_started
    delegated_host_root_modified
    delegated_docker_socket_mounted
    delegated_broad_host_mount_required
    delegated_raw_command_exposed
    delegated_backend_details_exposed
  ].each do |field|
    assert(action[field] == false, "D-Bus owner action payload must set #{field}=false")
  end

  assert(action["evidence_relative_path"] == evidence_record.fetch("evidence_relative_path"), "D-Bus owner action payload must preserve the relative evidence path")
  assert(action["evidence_sha256"] == evidence_record.fetch("evidence_sha256"), "D-Bus owner action payload must preserve the evidence digest")
end

runtime_command = runtime_go_command
unless runtime_command
  puts "#{SKIP_MARKER} (xnix-runtime-go and go are unavailable)"
  exit 0
end
unless command_available?("dbus-run-session") && command_available?("gdbus") && command_available?("xnix-dbus-smoke")
  puts "#{SKIP_MARKER} (D-Bus smoke tools unavailable)"
  exit 0
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__, chdir: PROJECT_ROOT.to_s)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

FileUtils.mkdir_p(RUN_ROOT)
FileUtils.mkdir_p(STATE_ROOT)
write_fake_launcher
evidence_record = prepare_fixture(runtime_command)

server_log = RUN_ROOT.join("xnix-dbus-smoke.log")
server_env = {
  "XNIX_RUNTIME_OWNER_STATE_ROOT" => STATE_ROOT.to_s,
  "XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT" => CACHE_ROOT.to_s,
  "XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER" => FAKE_LAUNCHER.to_s,
  "XNIX_RUNTIME_OWNER_GUEST_TIMEOUT" => "1s",
  "XNIX_RUNTIME_OWNER_TIMEOUT" => "5s",
  "XNIX_DBUS_CONTROLLED_LAUNCH_ARGS_FILE" => FAKE_LAUNCHER_ARGS.to_s
}
server_pid = spawn(server_env, "xnix-dbus-smoke", out: server_log.to_s, err: [:child, :out])

begin
  _stdout, stderr, status = Open3.capture3("gdbus", "wait", "--session", "--timeout", "5", BUS_NAME)
  assert(status.success?, "runtime smoke adapter must own #{BUS_NAME}: #{stderr}")

  stdout, stderr, status = Open3.capture3(
    "gdbus", "call",
    "--session",
    "--dest", BUS_NAME,
    "--object-path", OBJECT_PATH,
    "--method", PUBLIC_METHOD,
    evidence_record.fetch("evidence_relative_path")
  )
  assert(status.success?, "runtime smoke adapter must answer ShowRuntimeControlledLaunch: #{stderr}")
  assert(stdout.include?("runtime-controlled-launch-dbus-action"), "D-Bus response must expose controlled launch action evidence")
  assert(stdout.include?("desktop-action-dispatch"), "D-Bus response must classify the call as a desktop action dispatch")
  assert(stdout.include?("kde-dbus-runtime-status-action"), "D-Bus response must preserve the KDE Runtime-status route")
  assert(stdout.include?("go_owner_service_call_available"), "D-Bus response must expose owner service-call availability")
  assert(stdout.include?("kde_forwards_only_evidence_handle"), "D-Bus response must prove KDE forwards only the evidence handle")
  assert(stdout.include?("desktop_receipt_fields_reconstructed"), "D-Bus response must keep receipt reconstruction gates observable")
  assert(stdout.include?("desktop_kde_state_root_access"), "D-Bus response must keep KDE state-root gates observable")
  assert_owner_service_call(stdout, evidence_record)
  assert_no_forbidden(stdout, [STATE_ROOT.to_s, FAKE_LAUNCHER.to_s, ".exe", "wine ", "wine/", ".wine", "qemu-system", "program files", "docker.sock", "--privileged", "--network host", "type=bind"], "D-Bus controlled launch response")

  launcher_args = FAKE_LAUNCHER_ARGS.read
  assert(launcher_args.include?("--state-root\n#{STATE_ROOT}"), "fake managed launcher must receive Runtime-injected state root")
  assert(launcher_args.include?("--receipt-id\n#{LAUNCH_RECEIPT_ID}"), "fake managed launcher must receive the launch authorization receipt")
  assert(launcher_args.include?("--review-receipt-id\n#{REVIEW_RECEIPT_ID}"), "fake managed launcher must receive the session-gated review receipt")
  assert(launcher_args.include?("--session-id\n#{SESSION_ID}"), "fake managed launcher must receive the controlled execution session")
  assert(launcher_args.include?("--timeout\n1s"), "fake managed launcher must receive the owner-supplied guest timeout")

  puts PASS_MARKER
ensure
  begin
    Process.kill("TERM", server_pid)
    Process.wait(server_pid)
  rescue Errno::ESRCH, Errno::ECHILD
    nil
  end
end
