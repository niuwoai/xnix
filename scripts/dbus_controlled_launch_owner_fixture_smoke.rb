#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
ORIGINAL_ARGV = ARGV.dup
DEFAULT_BUS_NAME = "org.xnix.Compatibility1"
DEFAULT_OBJECT_PATH = "/org/xnix/Compatibility1"
DEFAULT_INTERFACE = "org.xnix.Compatibility1"
DEFAULT_PUBLIC_METHOD = "org.xnix.Compatibility1.ShowRuntimeControlledLaunch"
DEFAULT_DESKTOP_ENTRY_FILE = PROJECT_ROOT.join("kde/actions/xnix-runtime-status-controlled-launch.desktop").to_s
DEFAULT_APP_ID = "7zr"
DEFAULT_GUEST_BOUNDARY = "managed-known-app-guest-smoke"
SMOKE_NAME = "D-Bus controlled launch owner fixture smoke"
PASS_MARKER = "PASS: #{SMOKE_NAME}"
SKIP_MARKER = "SKIP: #{SMOKE_NAME}"
ALLOW_LOCAL_GO_COMPILE_ENV = "XNIX_ALLOW_LOCAL_GO_COMPILE"

options = {
  app_id: DEFAULT_APP_ID,
  guest_boundary: DEFAULT_GUEST_BOUNDARY,
  state_root: "",
  cache_root: "",
  gui_smoke_evidence_file: "",
  runtime_command: "",
  desktop_entry_file: DEFAULT_DESKTOP_ENTRY_FILE
}

OptionParser.new do |parser|
  parser.banner = "Usage: ruby scripts/dbus_controlled_launch_owner_fixture_smoke.rb [--app APP_ID] [--gui-smoke-evidence-file PATH]"
  parser.on("--app APP_ID", "Known Windows application id for the Runtime owner fixture.") { |value| options[:app_id] = value }
  parser.on("--guest-boundary NAME", "Runtime guest boundary, default: #{DEFAULT_GUEST_BOUNDARY}.") { |value| options[:guest_boundary] = value }
  parser.on("--state-root PATH", "Existing or creatable Runtime owner state root for the fixture run.") { |value| options[:state_root] = value }
  parser.on("--cache-root PATH", "Known Windows app cache root for the fixture run.") { |value| options[:cache_root] = value }
  parser.on("--gui-smoke-evidence-file PATH", "Optional GUI smoke or verified-catalog app-execution evidence file.") { |value| options[:gui_smoke_evidence_file] = value }
  parser.on("--runtime-command PATH", "Use an existing xnix-runtime-go binary instead of PATH discovery.") { |value| options[:runtime_command] = value }
  parser.on("--desktop-entry-file PATH", "KDE desktop action metadata used as the D-Bus invocation source.") { |value| options[:desktop_entry_file] = value }
end.parse!

abort "D-Bus controlled launch owner fixture smoke does not accept positional arguments" unless ARGV.empty?

def load_desktop_action_metadata(path)
  clean = Pathname.new(path).expand_path(PROJECT_ROOT).cleanpath
  abort "desktop entry file must stay under this checkout or /workspace" unless clean.to_s.start_with?(PROJECT_ROOT.to_s) || clean.to_s.start_with?("/workspace/")
  abort "desktop entry file must exist" unless clean.file?

  fields = {}
  clean.read.each_line do |line|
    next unless line.start_with?("X-Xnix-")

    key, value = line.strip.split("=", 2)
    fields[key] = value.to_s
  end
  expected = {
    "X-Xnix-KDE-Action-ID" => "xnix.runtime-status.controlled-launch",
    "X-Xnix-Runtime-Preview" => "xnix-runtime-go kde-controlled-launch-action-preview",
    "X-Xnix-Restricted-Smoke-Plan" => "xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview",
    "X-Xnix-DBus-Service" => DEFAULT_BUS_NAME,
    "X-Xnix-DBus-Object-Path" => DEFAULT_OBJECT_PATH,
    "X-Xnix-DBus-Method" => DEFAULT_PUBLIC_METHOD,
    "X-Xnix-Forwarded-Argument" => "evidence-relative-path",
    "X-Xnix-Forwards-Only-Evidence-Handle" => "true",
    "X-Xnix-KDE-Policy-Owner" => "false",
    "X-Xnix-Owner-Service-Args-Exposed-To-KDE" => "false",
    "X-Xnix-State-Root-Access" => "false",
    "X-Xnix-Receipt-Reconstruction" => "false",
    "X-Xnix-Backend-Launch-Enabled" => "false",
    "X-Xnix-Execution-Started" => "false",
    "X-Xnix-Host-Root-Modified" => "false",
    "X-Xnix-Docker-Socket-Mounted" => "false",
    "X-Xnix-Privileged-Container-Required" => "false",
    "X-Xnix-Host-Network-Required" => "false"
  }
  expected.each do |key, value|
    abort "desktop action metadata mismatch for #{key}" unless fields.fetch(key, "") == value
  end
  {
    "action_id" => fields.fetch("X-Xnix-KDE-Action-ID"),
    "bus_name" => fields.fetch("X-Xnix-DBus-Service"),
    "object_path" => fields.fetch("X-Xnix-DBus-Object-Path"),
    "interface" => fields.fetch("X-Xnix-DBus-Method").split(".")[0...-1].join("."),
    "method" => fields.fetch("X-Xnix-DBus-Method"),
    "forwarded_argument" => fields.fetch("X-Xnix-Forwarded-Argument")
  }
end

APP_ID = options.fetch(:app_id).to_s.strip
GUEST_BOUNDARY = options.fetch(:guest_boundary).to_s.strip
DESKTOP_ACTION = load_desktop_action_metadata(options.fetch(:desktop_entry_file))
BUS_NAME = DESKTOP_ACTION.fetch("bus_name")
OBJECT_PATH = DESKTOP_ACTION.fetch("object_path")
INTERFACE = DESKTOP_ACTION.fetch("interface")
PUBLIC_METHOD = DESKTOP_ACTION.fetch("method")
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
STATE_ROOT = options.fetch(:state_root).to_s.strip.empty? ? RUN_ROOT.join("state") : Pathname.new(options.fetch(:state_root)).expand_path(PROJECT_ROOT).cleanpath
CACHE_ROOT = options.fetch(:cache_root).to_s.strip.empty? ? PROJECT_ROOT.join(".cache", "xnix", "known-winapps") : Pathname.new(options.fetch(:cache_root)).expand_path(PROJECT_ROOT).cleanpath
GUI_SMOKE_EVIDENCE_FILE = options.fetch(:gui_smoke_evidence_file).to_s.strip
RUNTIME_COMMAND_OVERRIDE = options.fetch(:runtime_command).to_s.strip
FAKE_LAUNCHER = RUN_ROOT.join("fake-xnix-compat-launch")
FAKE_LAUNCHER_ARGS = RUN_ROOT.join("fake-launcher-args.txt")

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
  return [RUNTIME_COMMAND_OVERRIDE] unless RUNTIME_COMMAND_OVERRIDE.empty?
  return ["xnix-runtime-go"] if command_available?("xnix-runtime-go")
  return ["go", "run", "./cmd/xnix-runtime-go"] if ENV.fetch(ALLOW_LOCAL_GO_COMPILE_ENV, "") == "1" && command_available?("go")

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

def delegated_payload(fixture)
  {
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
    "session_gated_review_receipt_id" => fixture.fetch("session_gated_review_receipt_id"),
    "launch_authorization_receipt_id" => fixture.fetch("launch_authorization_receipt_id"),
    "controlled_execution_session_consumed" => true,
    "controlled_execution_session_id" => fixture.fetch("controlled_execution_session_id"),
    "controlled_session_digest_verified" => true,
    "controlled_session_relative_path" => fixture.fetch("controlled_session_relative_path"),
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
  }
end

def write_fake_launcher(payload)
  source = [
    "#!/usr/bin/env ruby",
    "# frozen_string_literal: true",
    "require \"json\"",
    "args_path = ENV.fetch(\"XNIX_DBUS_CONTROLLED_LAUNCH_ARGS_FILE\")",
    "File.write(args_path, ARGV.join(\"\\n\") + \"\\n\")",
    "puts JSON.generate(#{payload.inspect})"
  ].join("\n")
  File.write(FAKE_LAUNCHER, source)
  File.chmod(0o755, FAKE_LAUNCHER)
end

def prepare_fixture(runtime_command)
  env = {}
  args = [
    *runtime_command,
    "known-app-runtime-status-launch-owner-fixture-record",
    "--app", APP_ID,
    "--state-root", STATE_ROOT.to_s,
    "--cache-root", CACHE_ROOT.to_s,
    "--guest-boundary", GUEST_BOUNDARY
  ]
  args += ["--gui-smoke-evidence-file", GUI_SMOKE_EVIDENCE_FILE] unless GUI_SMOKE_EVIDENCE_FILE.empty?
  fixture, fixture_stdout = run_json(
    env,
    *args
  )
  unless fixture["fixture_ready"] == true
    reason = fixture.fetch("skip_reason", "known Windows app artifact unavailable")
    puts "#{SKIP_MARKER} (#{reason})"
    exit 0
  end
  assert(fixture["launch_authorization_receipt_id"].to_s.start_with?("known-app-launch-authorization-"), "launch receipt fixture id must be stable")
  assert(fixture["controlled_execution_session_id"].to_s.start_with?("known-app-controlled-execution-session-"), "controlled session fixture id must be stable")
  assert(fixture["session_gated_review_receipt_id"].to_s.start_with?("known-app-session-gated-launch-review-"), "review receipt fixture id must be stable")
  assert(fixture["evidence_relative_path"].to_s.start_with?("runtime/kde-runtime-status-launch-evidence/"), "Runtime-status evidence fixture must expose only a relative path")
  assert(fixture["runtime_owned"] == true, "Runtime-status evidence fixture must be Runtime-owned")
  assert(fixture["go_runtime_backed"] == true, "Runtime-status evidence fixture must be Go backed")
  assert(fixture["desktop_trigger_ready"] == true, "Runtime-status evidence fixture must expose a desktop trigger")
  assert(fixture["desktop_callable_route"] == "kde-dbus-runtime-status-action", "Runtime-status evidence fixture must target the KDE D-Bus Runtime-status route")
  assert(fixture["desktop_callable_runtime_method"] == "ShowRuntimeControlledLaunch", "Runtime-status evidence fixture must target ShowRuntimeControlledLaunch")
  assert(fixture["desktop_callable_execution_type"] == "known-app-kde-runtime-status-launch-execution", "Runtime-status evidence fixture must target the Runtime launch execution")
  assert(fixture["desktop_dbus_method"] == PUBLIC_METHOD, "Runtime-status evidence fixture must expose the public D-Bus method")
  assert(fixture["owner_service_call_ready"] == true, "Runtime-status evidence fixture must expose a ready owner service call")
  assert(fixture["owner_service_boundary"] == "go-runtime-owner-in-process-service", "Runtime-status evidence fixture must route through the Go owner service")
  assert(fixture["owner_service_method"] == "ShowRuntimeControlledLaunch", "Runtime-status evidence fixture must preserve the owner service method")
  assert(fixture["owner_service_call_type"] == "desktop-action-dispatch", "Runtime-status evidence fixture must classify the owner call as a desktop action dispatch")
  assert(fixture["owner_service_call_args"] == ["ShowRuntimeControlledLaunch", "evidence-relative-path", fixture.fetch("evidence_relative_path")], "Runtime-status evidence fixture must expose evidence-only owner service args")
  assert(fixture["runtime_owner_service_supplies_inputs"] == true, "Runtime-status evidence fixture must keep owner inputs supplied by Runtime")
  assert(fixture["desktop_evidence_handle_forwarded"] == true, "Runtime-status evidence fixture must forward only the evidence handle")
  assert(fixture["kde_forwards_only_evidence_handle"] == true, "Runtime-status evidence fixture must keep KDE evidence-only")
  assert(fixture["desktop_receipt_fields_reconstructed"] == false, "Runtime-status evidence fixture must not reconstruct receipt fields in KDE")
  assert(fixture["desktop_kde_state_root_access"] == false, "Runtime-status evidence fixture must not grant KDE state-root access")
  assert_no_forbidden(fixture_stdout, [STATE_ROOT.to_s, GUI_SMOKE_EVIDENCE_FILE, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime-status owner fixture output")
  fixture
end

def service_call_materialization_from_runtime(runtime_command, fixture)
  materialization, materialization_stdout = run_json(
    {},
    *runtime_command,
    "desktop-trigger-service-call-materialization-preview",
    "--state-root", STATE_ROOT.to_s,
    "--desktop-entry-file", PROJECT_ROOT.join("kde/actions/xnix-runtime-status-controlled-launch.desktop").to_s,
    "--human-authorized-smoke",
    "--evidence-relative-path", fixture.fetch("evidence_relative_path")
  )
  assert(materialization["schema_version"] == "xnix.runtime.desktop_trigger_service_call_materialization.v1", "Runtime-status service call materialization must expose the Go-owned schema")
  assert(materialization["request_type"] == "desktop-trigger-service-call-materialization-preview", "Runtime-status service call materialization must use the Go-owned materialization type")
  assert(materialization["runtime_method"] == "PreviewDesktopTriggerServiceCallMaterialization", "Runtime-status service call materialization must use the Go preview")
  assert(materialization["read_method"] == "GetDesktopTriggerServiceCallMaterialization", "Runtime-status service call materialization must expose the read method")
  assert(materialization["materialization_state"] == "ready-for-human-authorized-service-call", "Runtime-status service call materialization must be ready for human-authorized fixture smoke")
  assert(materialization["dry_run_review_state"] == "blocked-missing-full-checkpoint", "Runtime-status service call materialization must not claim full checkpoint promotion")
  assert(materialization["owner_trigger_state"] == "ready", "Runtime-status service call materialization must consume the owner trigger internally")
  assert(materialization["full_checkpoint_state"] == "needs-full-checkpoint", "Runtime-status service call materialization must keep formal full checkpoint pending")
  assert(materialization["human_authorized_smoke"] == true, "Runtime-status service call materialization must require explicit human smoke authorization")
  assert(materialization["full_checkpoint_promotion_claimed"] == false, "Runtime-status service call materialization must not claim checkpoint promotion")
  assert(materialization["formal_release_ready"] == false, "Runtime-status service call materialization must not claim release readiness")
  assert(materialization["evidence_relative_path"] == fixture.fetch("evidence_relative_path"), "Runtime-status service call materialization must consume the fixture evidence path")
  assert(materialization["evidence_sha256"] == fixture.fetch("evidence_sha256"), "Runtime-status service call materialization must preserve the fixture evidence digest")
  assert(materialization["evidence_digest_verified"] == true, "Runtime-status service call materialization must verify the evidence digest")
  assert(materialization["desktop_callable_route"] == "kde-dbus-runtime-status-action", "Runtime-status service call materialization must target the KDE D-Bus Runtime-status route")
  assert(materialization["desktop_callable_runtime_method"] == "ShowRuntimeControlledLaunch", "Runtime-status service call materialization must target ShowRuntimeControlledLaunch")
  assert(materialization["desktop_callable_execution_type"] == "known-app-kde-runtime-status-launch-execution", "Runtime-status service call materialization must target the Runtime launch execution")
  assert(materialization["desktop_dbus_method"] == PUBLIC_METHOD, "Runtime-status service call materialization must expose the public D-Bus method")
  assert(materialization["owner_service_call_ready"] == true, "Runtime-status service call materialization must expose a ready owner service call")
  assert(materialization["owner_service_boundary"] == "go-runtime-owner-in-process-service", "Runtime-status service call materialization must route through the Go owner service")
  assert(materialization["owner_service_method"] == "ShowRuntimeControlledLaunch", "Runtime-status service call materialization must preserve the owner service method")
  assert(materialization["owner_service_call_type"] == "desktop-action-dispatch", "Runtime-status service call materialization must classify the owner call as a desktop action dispatch")
  assert(materialization["owner_service_call_args"] == fixture.fetch("owner_service_call_args"), "Runtime-status service call materialization must preserve fixture owner service args")
  assert(materialization["owner_service_cli_args"] == ["--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", fixture.fetch("evidence_relative_path")], "Runtime-status service call materialization must expose evidence-only owner service CLI args")
  assert(materialization["runtime_owner_service_supplies_inputs"] == true, "Runtime-status service call materialization must keep owner inputs supplied by Runtime")
  assert(materialization["desktop_evidence_handle_forwarded"] == true, "Runtime-status service call materialization must forward only the evidence handle")
  assert(materialization["kde_forwards_only_evidence_handle"] == true, "Runtime-status service call materialization must keep KDE evidence-only")
  assert(materialization["kde_receives_materialized_owner_args"] == false, "Runtime-status service call materialization must not hand owner args to KDE")
  assert(materialization["desktop_receipt_fields_reconstructed"] == false, "Runtime-status service call materialization must not reconstruct receipt fields in KDE")
  assert(materialization["desktop_kde_state_root_access"] == false, "Runtime-status service call materialization must not grant KDE state-root access")
  assert(materialization["state_root_path_exposed"] == false, "Runtime-status service call materialization must not expose the state root")
  assert(materialization["raw_launcher_output_exposed"] == false, "Runtime-status service call materialization must not expose raw launcher output")
  assert(materialization["backend_details_exposed"] == false, "Runtime-status service call materialization must not expose backend details")
  assert(materialization["service_call_dispatched"] == false, "Runtime-status service call materialization preview must not dispatch the service call")
  assert(materialization["dbus_called"] == false, "Runtime-status service call materialization preview must not call D-Bus")
  assert(materialization["host_root_modified"] == false, "Runtime-status service call materialization must not mutate the host root")
  assert_no_forbidden(materialization_stdout, [STATE_ROOT.to_s, PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime-status service call materialization output")
  args = materialization.fetch("owner_service_call_args")
  {
    "method" => materialization.fetch("desktop_dbus_method"),
    "owner_service_method" => args.fetch(0),
    "handoff_kind" => args.fetch(1),
    "evidence_relative_path" => args.fetch(2)
  }
end

def assert_service_call_materialization_response(stdout, evidence_record, desktop_trigger)
  payload = variant_string_field(stdout, "go_service_call_materialization_json")
  materialization = JSON.parse(payload)
  assert(materialization["schema_version"] == "xnix.runtime.desktop_trigger_service_call_materialization.v1", "D-Bus materialization must expose the Go materialization schema")
  assert(materialization["request_type"] == "desktop-trigger-service-call-materialization-preview", "D-Bus materialization must expose the Go materialization request type")
  assert(materialization["runtime_method"] == "PreviewDesktopTriggerServiceCallMaterialization", "D-Bus materialization must preserve the Go preview runtime method")
  assert(materialization["read_method"] == "GetDesktopTriggerServiceCallMaterialization", "D-Bus materialization must preserve the Go preview read method")
  assert(materialization["materialization_state"] == "ready-for-human-authorized-service-call", "D-Bus materialization must be ready for human-authorized fixture smoke")
  assert(materialization["dry_run_review_state"] == "blocked-missing-full-checkpoint", "D-Bus materialization must keep full checkpoint pending")
  assert(materialization["owner_trigger_state"] == "ready", "D-Bus materialization must consume the owner trigger internally")
  assert(materialization["human_authorized_smoke"] == true, "D-Bus materialization must require human smoke authorization")
  assert(materialization["full_checkpoint_promotion_claimed"] == false, "D-Bus materialization must not claim checkpoint promotion")
  assert(materialization["formal_release_ready"] == false, "D-Bus materialization must not claim release readiness")
  assert(materialization["desktop_dbus_method"] == desktop_trigger.fetch("method"), "D-Bus materialization must preserve the public D-Bus method")
  assert(materialization["owner_service_call_args"] == [desktop_trigger.fetch("owner_service_method"), desktop_trigger.fetch("handoff_kind"), desktop_trigger.fetch("evidence_relative_path")], "D-Bus materialization must provide the owner service call args consumed by the C adapter")
  assert(materialization["evidence_relative_path"] == evidence_record.fetch("evidence_relative_path"), "D-Bus materialization must preserve the relative evidence path")
  assert(materialization["evidence_sha256"] == evidence_record.fetch("evidence_sha256"), "D-Bus materialization must preserve the evidence digest")
  assert(materialization["evidence_digest_verified"] == true, "D-Bus materialization must verify the handoff digest")
  assert(materialization["runtime_owner_service_supplies_inputs"] == true, "D-Bus materialization must keep owner inputs supplied by Runtime")
  assert(materialization["kde_forwards_only_evidence_handle"] == true, "D-Bus materialization must keep KDE evidence-only")
  assert(materialization["kde_receives_materialized_owner_args"] == false, "D-Bus materialization must not hand owner args to KDE")
  assert(materialization["desktop_receipt_fields_reconstructed"] == false, "D-Bus materialization must keep receipt reconstruction disabled")
  assert(materialization["desktop_kde_state_root_access"] == false, "D-Bus materialization must keep KDE state-root access disabled")
  assert(materialization["state_root_path_exposed"] == false, "D-Bus materialization must not expose state root paths")
  assert(materialization["raw_launcher_output_exposed"] == false, "D-Bus materialization must not expose raw launcher output")
  assert(materialization["backend_details_exposed"] == false, "D-Bus materialization must not expose backend details")
  assert(materialization["service_call_dispatched"] == false, "D-Bus materialization preview must not dispatch the owner service itself")
  assert(materialization["host_root_modified"] == false, "D-Bus materialization must not mutate the host root")
end

def assert_owner_service_call(stdout, evidence_record, desktop_trigger)
  payload = variant_string_field(stdout, "go_owner_service_call_json")
  call = JSON.parse(payload)
  assert(call["schema_version"] == "xnix.runtime.owner_service_call.v1", "D-Bus owner call must expose the owner service-call schema")
  assert(call["request_type"] == "runtime-owner-service-call", "D-Bus owner call must expose the owner service-call request type")
  assert(call["service_type"] == "go-runtime-owner-in-process-service", "D-Bus owner call must route through the Go owner service")
  assert(call["method"] == desktop_trigger.fetch("owner_service_method"), "D-Bus owner call must preserve the Go-owned owner service method")
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
    "delegated_session_gated_review_receipt_id" => evidence_record.fetch("session_gated_review_receipt_id"),
    "delegated_launch_authorization_receipt_id" => evidence_record.fetch("launch_authorization_receipt_id"),
    "delegated_controlled_execution_session_id" => evidence_record.fetch("controlled_execution_session_id")
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
  puts "#{SKIP_MARKER} (xnix-runtime-go is unavailable; local Go compilation is disabled by default, set #{ALLOW_LOCAL_GO_COMPILE_ENV}=1 only for an explicit local override or build on q4 with scripts/remote_go_build.rb --execute)"
  exit 0
end
unless command_available?("dbus-run-session") && command_available?("gdbus") && command_available?("xnix-dbus-smoke")
  puts "#{SKIP_MARKER} (D-Bus smoke tools unavailable)"
  exit 0
end

unless ENV["DBUS_SESSION_BUS_ADDRESS"]
  stdout, stderr, status = Open3.capture3("dbus-run-session", "--", "ruby", __FILE__, *ORIGINAL_ARGV, chdir: PROJECT_ROOT.to_s)
  print stdout
  warn stderr unless stderr.empty?
  exit status.exitstatus
end

FileUtils.mkdir_p(RUN_ROOT)
FileUtils.mkdir_p(STATE_ROOT)
evidence_record = prepare_fixture(runtime_command)
write_fake_launcher(delegated_payload(evidence_record))
desktop_trigger = service_call_materialization_from_runtime(runtime_command, evidence_record)

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
    "--method", DESKTOP_ACTION.fetch("method"),
    desktop_trigger.fetch("evidence_relative_path")
  )
  assert(status.success?, "runtime smoke adapter must answer ShowRuntimeControlledLaunch: #{stderr}")
  assert(stdout.include?("runtime-controlled-launch-dbus-action"), "D-Bus response must expose controlled launch action evidence")
  assert(stdout.include?("desktop-action-dispatch"), "D-Bus response must classify the call as a desktop action dispatch")
  assert(stdout.include?("kde-dbus-runtime-status-action"), "D-Bus response must preserve the KDE Runtime-status route")
  assert(DESKTOP_ACTION.fetch("action_id") == "xnix.runtime-status.controlled-launch", "D-Bus invocation must be sourced from the KDE controlled-launch action metadata")
  assert(DESKTOP_ACTION.fetch("forwarded_argument") == "evidence-relative-path", "D-Bus invocation must forward only the KDE evidence handle")
  assert(stdout.include?("go_service_call_materialization_available"), "D-Bus response must expose Go service-call materialization availability")
  assert(stdout.include?("go_service_call_materialization_json"), "D-Bus response must expose Go service-call materialization metadata")
  assert(stdout.include?("go_owner_service_call_available"), "D-Bus response must expose owner service-call availability")
  assert(stdout.include?("kde_forwards_only_evidence_handle"), "D-Bus response must prove KDE forwards only the evidence handle")
  assert(stdout.include?("desktop_receipt_fields_reconstructed"), "D-Bus response must keep receipt reconstruction gates observable")
  assert(stdout.include?("desktop_kde_state_root_access"), "D-Bus response must keep KDE state-root gates observable")
  assert_service_call_materialization_response(stdout, evidence_record, desktop_trigger)
  assert_owner_service_call(stdout, evidence_record, desktop_trigger)
  assert_no_forbidden(stdout, [STATE_ROOT.to_s, FAKE_LAUNCHER.to_s, ".exe", "wine ", "wine/", ".wine", "qemu-system", "program files", "docker.sock", "--privileged", "--network host", "type=bind"], "D-Bus controlled launch response")

  launcher_args = FAKE_LAUNCHER_ARGS.read
  assert(launcher_args.include?("--state-root\n#{STATE_ROOT}"), "fake managed launcher must receive Runtime-injected state root")
  assert(launcher_args.include?("--receipt-id\n#{evidence_record.fetch("launch_authorization_receipt_id")}"), "fake managed launcher must receive the launch authorization receipt")
  assert(launcher_args.include?("--review-receipt-id\n#{evidence_record.fetch("session_gated_review_receipt_id")}"), "fake managed launcher must receive the session-gated review receipt")
  assert(launcher_args.include?("--session-id\n#{evidence_record.fetch("controlled_execution_session_id")}"), "fake managed launcher must receive the controlled execution session")
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
