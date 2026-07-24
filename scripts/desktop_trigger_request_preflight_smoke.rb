#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "optparse"
require "pathname"
require "securerandom"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip
SMOKE_NAME = "desktop-trigger request preflight smoke"
PASS_MARKER = "PASS: desktop-trigger request preflight smoke"
APP_ID = "7zr"
APP_VERSION = "26.02"
DISPLAY_NAME = "7-Zip Console"
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
SESSION_ID = "known-app-controlled-execution-session-7zr-26.02"
LAUNCH_RECEIPT_ID = "known-app-launch-authorization-7zr-26.02"
REVIEW_RECEIPT_ID = "known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02"
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "desktop-trigger-request-preflight-smoke", RUN_ID)
STATE_ROOT = WORK_ROOT.join("state")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
DESKTOP_ENTRY_FILE = PROJECT_ROOT.join("kde", "actions", "xnix-runtime-status-controlled-launch.desktop")

def parse_options(argv)
  options = { format: "text" }
  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/desktop_trigger_request_preflight_smoke.rb [--format text|json]"
    parser.on("--format FORMAT", "Output format: text or json") { |value| options[:format] = value }
  end.parse!(argv)
  assert(%w[text json].include?(options.fetch(:format)), "format must be text or json")
  options
end

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

def smoke_report(blocked, ready, record)
  {
    "version" => VERSION,
    "schema_version" => "xnix.runtime.desktop_trigger_request_preflight_smoke.v1",
    "report_type" => "desktop-trigger-request-preflight-smoke",
    "smoke_passed" => true,
    "preflight_smoke_state" => "passed",
    "record_request_type" => record.fetch("request_type"),
    "record_evidence_state" => record.fetch("evidence_state"),
    "evidence_relative_path" => record.fetch("evidence_relative_path"),
    "evidence_sha256" => record.fetch("evidence_sha256"),
    "blocked_preflight_state" => blocked.fetch("preflight_state"),
    "blocked_materialization_state" => blocked.fetch("materialization_state"),
    "blocked_full_checkpoint_state" => blocked.fetch("full_checkpoint_state"),
    "ready_preflight_state" => ready.fetch("preflight_state"),
    "ready_materialization_state" => ready.fetch("materialization_state"),
    "ready_full_checkpoint_state" => ready.fetch("full_checkpoint_state"),
    "owner_service_call_shape_verified" => ready.fetch("owner_service_call_shape_verified"),
    "owner_service_call_ready" => ready.fetch("owner_service_call_ready"),
    "operator_request_ready" => ready.fetch("operator_request_ready"),
    "formal_promotion_observed_in_ready_fixture" => ready.fetch("formal_promotion_observed"),
    "formal_release_ready" => false,
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "ruby_smoke_orchestration_only" => true,
    "kde_presentation_only" => true,
    "kde_forwards_only_evidence_handle" => true,
    "kde_receives_materialized_owner_args" => false,
    "owner_call_arguments_exposed" => false,
    "owner_cli_arguments_exposed" => false,
    "request_object_written" => false,
    "permission_grant_created" => false,
    "service_call_dispatch_enabled" => false,
    "service_call_dispatched" => false,
    "dbus_called" => false,
    "desktop_launch_enabled" => false,
    "backend_launch_enabled" => false,
    "execution_started" => false,
    "runtime_state_written" => false,
    "kde_configuration_written" => false,
    "docker_executed" => false,
    "qemu_executed" => false,
    "wine_executed" => false,
    "colima_executed" => false,
    "network_required" => false,
    "network_checks_run" => false,
    "package_manager_invoked" => false,
    "privileged_container_required" => false,
    "host_root_modified" => false,
    "desktop_safe_summary" => "Desktop-trigger request preflight smoke verified blocked and promoted fixture states without dispatching service calls, calling D-Bus, launching a desktop action, starting a backend, or mutating the host."
  }
end

def runtime_status_projection
  {
    "projection_type" => "known-app-kde-runtime-status-launch-delegated-evidence",
    "runtime_method" => "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence",
    "request_type" => "windows-known-app-dispatch-smoke",
    "status" => "passed",
    "app_id" => APP_ID,
    "display_name" => DISPLAY_NAME,
    "app_version" => APP_VERSION,
    "guest_boundary" => GUEST_BOUNDARY,
    "runtime_owned_dispatch" => true,
    "artifact_verified" => true,
    "marker_observed" => true,
    "session_gated_controlled_dispatch_consumed" => true,
    "session_gated_controlled_dispatch_state" => "created-after-session-gated-review",
    "session_gated_review_receipt_id" => REVIEW_RECEIPT_ID,
    "launch_authorization_receipt_id" => LAUNCH_RECEIPT_ID,
    "launch_authorization_receipt_state" => "recorded",
    "launch_gate_state" => "controlled-dispatch-ready",
    "launch_gate_consumed" => true,
    "launch_gate_receipt_accepted" => true,
    "launch_gate_guest_boundary_accepted" => true,
    "controlled_dispatch_ready" => true,
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
    "state_root_path_exposed" => false,
    "managed_launcher_path_exposed" => false,
    "raw_launcher_output_exposed" => false,
    "backend_details_exposed" => false,
    "compatibility_center_projection_ready" => true,
    "kde_center_projection_ready" => true,
    "desktop_safe_summary" => "#{DISPLAY_NAME} delegated Runtime launcher evidence is projected for Compatibility Center and KDE Center consumption without exposing launcher output or state-root paths."
  }
end

def preflight_argv(runtime_command, record, promoted: false)
  [
    *runtime_command,
    "desktop-trigger-request-preflight-preview",
    "--state-root", STATE_ROOT.to_s,
    "--desktop-entry-file", DESKTOP_ENTRY_FILE.to_s,
    "--evidence-relative-path", record.fetch("evidence_relative_path"),
    "--expected-evidence-sha256", record.fetch("evidence_sha256"),
    *(promoted ? ["--full-checkpoint-promoted"] : [])
  ]
end

FileUtils.mkdir_p(STATE_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)

runtime_command = runtime_go_command
unless runtime_command
  warn "FAIL: #{SMOKE_NAME} requires xnix-runtime-go or go"
  exit 1
end

options = parse_options(ARGV)

record, record_stdout = run_json(
  go_env,
  *runtime_command,
  "known-app-kde-runtime-status-launch-evidence-record",
  "--state-root", STATE_ROOT.to_s,
  "--evidence-json", JSON.generate(runtime_status_projection)
)

assert(record["request_type"] == "known-app-kde-runtime-status-launch-evidence-record", "record request type must match")
assert(record["evidence_state"] == "persisted", "record must persist Runtime-status launch evidence")
assert(record["evidence_relative_path"].to_s.start_with?("runtime/kde-runtime-status-launch-evidence/"), "record must expose only a relative evidence path")
assert(record["known_app_smoke_evidence_ready"] == true, "record must produce known app smoke evidence")
assert(record["compatibility_center_projection_ready"] == true, "record must produce Compatibility Center evidence")
assert(record["kde_center_projection_ready"] == true, "record must produce KDE Center evidence")
assert_false_payload(
  record,
  %w[state_root_path_exposed evidence_path_exposed managed_launcher_path_exposed raw_launcher_output_exposed backend_details_exposed host_root_modified docker_socket_mounted broad_host_mount_required desktop_launch_enabled backend_launch_enabled execution_started backend_process_started],
  "record"
)

common_forbidden_terms = [
  PROJECT_ROOT.to_s,
  STATE_ROOT.to_s,
  DESKTOP_ENTRY_FILE.to_s,
  "owner_service_call_args",
  "owner_service_cli_args",
  " --service-call ",
  "XNIX_RUNTIME_OWNER_",
  "qemu-system",
  "wine ",
  ".wine",
  "program files",
  ".exe",
  ENV.fetch("USER", "")
]

assert_no_forbidden(record_stdout, common_forbidden_terms, "record output")

blocked, blocked_stdout = run_json(go_env, *preflight_argv(runtime_command, record, promoted: false))
assert(blocked["request_type"] == "desktop-trigger-request-preflight-preview", "blocked preflight request type must match")
assert(blocked["preflight_state"] == "blocked-missing-promotion", "preflight must block before formal promotion")
assert(blocked["materialization_state"] == "blocked-missing-full-checkpoint", "preflight must preserve the materialization blocker")
assert(blocked["full_checkpoint_state"] == "needs-full-checkpoint", "preflight must require the full checkpoint")
assert(blocked["operator_request_ready"] == false, "blocked preflight must not be operator-ready")
assert(blocked["owner_service_call_shape_verified"] == false, "blocked preflight must not verify the owner call shape")
assert(blocked["formal_promotion_required"] == true, "blocked preflight must require formal promotion")
assert(blocked["formal_promotion_observed"] == false, "blocked preflight must not observe promotion")
assert(blocked["formal_release_ready"] == false, "blocked preflight must not claim formal release readiness")
assert_false_payload(
  blocked,
  %w[owner_service_call_ready kde_receives_materialized_owner_args request_object_written permission_grant_created service_call_dispatch_enabled service_call_dispatched dbus_called desktop_launch_enabled backend_launch_enabled execution_started runtime_state_written kde_configuration_written network_required host_root_modified],
  "blocked preflight"
)
assert_no_forbidden(blocked_stdout, common_forbidden_terms, "blocked preflight output")

ready, ready_stdout = run_json(go_env, *preflight_argv(runtime_command, record, promoted: true))
assert(ready["request_type"] == "desktop-trigger-request-preflight-preview", "ready preflight request type must match")
assert(ready["preflight_state"] == "ready-for-operator-request", "promoted preflight must be ready for an operator request")
assert(ready["materialization_state"] == "ready-for-human-authorized-service-call", "promoted preflight must consume ready materialization")
assert(ready["dry_run_review_state"] == "accepted-review", "promoted preflight must consume an accepted dry-run review")
assert(ready["owner_trigger_state"] == "ready", "promoted preflight must consume a ready owner trigger")
assert(ready["full_checkpoint_state"] == "ready", "promoted preflight must observe full checkpoint readiness")
assert(ready["runtime_status_evidence_state"] == "ready", "promoted preflight must consume Runtime-status evidence")
assert(ready["evidence_digest_verified"] == true, "promoted preflight must verify the evidence digest")
assert(ready["expected_digest_matched"] == true, "promoted preflight must match the expected digest")
assert(ready["owner_service_call_shape_verified"] == true, "promoted preflight must verify the owner service call shape inside Runtime")
assert(ready["owner_service_call_ready"] == true, "promoted preflight must report owner service call readiness")
assert(ready["operator_request_ready"] == true, "promoted preflight must be ready for an operator request")
assert(ready["formal_promotion_required"] == true, "promoted preflight must keep formal promotion required")
assert(ready["formal_promotion_observed"] == true, "promoted preflight must observe explicit promotion")
assert(ready["formal_release_ready"] == false, "promoted preflight must not claim formal release readiness")
assert(ready["runtime_owner_service_supplies_inputs"] == true, "promoted preflight must keep Runtime owner supplied inputs")
assert(ready["desktop_evidence_handle_forwarded"] == true, "promoted preflight must keep desktop evidence-handle forwarding")
assert(ready["runtime_owned"] == true, "promoted preflight must stay Runtime-owned")
assert(ready["go_runtime_backed"] == true, "promoted preflight must stay Go-backed")
assert(ready["kde_presentation_only"] == true, "promoted preflight must keep KDE presentation-only")
assert(ready["kde_forwards_only_evidence_handle"] == true, "promoted preflight must keep KDE evidence-only")
assert_false_payload(
  ready,
  %w[kde_receives_materialized_owner_args request_object_written permission_grant_created service_call_dispatch_enabled service_call_dispatched dbus_called desktop_launch_enabled backend_launch_enabled execution_started runtime_state_written kde_configuration_written network_required host_root_modified state_root_path_exposed raw_launcher_output_exposed backend_details_exposed],
  "ready preflight"
)
assert_no_forbidden(ready_stdout, common_forbidden_terms, "ready preflight output")

report = smoke_report(blocked, ready, record)
assert_no_forbidden(JSON.generate(report), [PROJECT_ROOT.to_s, STATE_ROOT.to_s, DESKTOP_ENTRY_FILE.to_s, "owner_service_call_args", "owner_service_cli_args", " --service-call ", "XNIX_RUNTIME_OWNER_", "qemu-system", "wine ", ".wine", "program files", ".exe", ENV.fetch("USER", "")], "smoke report")

case options.fetch(:format)
when "json"
  puts JSON.pretty_generate(report)
else
  puts PASS_MARKER
end
