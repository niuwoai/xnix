#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "securerandom"
require "shellwords"
require_relative "../lib/xnix/qemu"
require_relative "../lib/xnix/ssh_probe"
require_relative "../lib/xnix/ssh_test_key"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "staged-launcher-dispatch-smoke")
RUN_ID = "#{Time.now.utc.strftime("%Y%m%d%H%M%S")}-#{Process.pid}-#{SecureRandom.hex(4)}"
RUN_ROOT = WORK_ROOT.join(RUN_ID)
BUILD_DIR = RUN_ROOT.join("build")
STAGE_ROOT = RUN_ROOT.join("stage")
AUTHORIZATION_STATE_ROOT = RUN_ROOT.join("authorization-state")
RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH = RUN_ROOT.join("runtime-status-launch-evidence.json")
KNOWN_APP_CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
LAUNCHER_BIN = BUILD_DIR.join("xnix-compat-launch")
STAGED_LAUNCHER = STAGE_ROOT.join("usr/local/bin/xnix-compat-launch")
SERIAL_LOG_PATH = WORK_ROOT.join("qemu-serial.log")
APP_ID = ENV.fetch("XNIX_KNOWN_WINAPP_ID", "7zr")
APP_ARGS = Shellwords.split(ENV.fetch("XNIX_KNOWN_WINAPP_ARGS", ""))
BOOT_TIMEOUT_SECONDS = Integer(ENV.fetch("XNIX_KNOWN_WINAPP_BOOT_TIMEOUT", "180"))
RETRY_INTERVAL_SECONDS = 1
GUEST_BOUNDARY = "managed-known-app-guest-smoke"
SMOKE_NAME = "staged managed launcher dispatch smoke"
GO_BUILD_STEP = "go build"

def go_env
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
    "GOTMPDIR" => GO_TMP_ROOT.to_s
  }
end

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
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

def stop_qemu(wait_thread)
  return unless wait_thread.alive?

  Process.kill("TERM", wait_thread.pid)
  wait_thread.join(5)
  Process.kill("KILL", wait_thread.pid) if wait_thread.alive?
rescue Errno::ESRCH
  nil
end

FileUtils.mkdir_p(WORK_ROOT)
FileUtils.mkdir_p(BUILD_DIR)
FileUtils.mkdir_p(AUTHORIZATION_STATE_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)

dispatch_preview, dispatch_preview_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go", "windows-known-app-dispatch-preview",
  "--app", APP_ID,
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s
)

assert(dispatch_preview["request_type"] == "windows-known-app-dispatch-preview", "dispatch preview request type must match")
assert(dispatch_preview["runtime_owned_dispatch"] == true, "dispatch preview must be Runtime-owned")
assert(dispatch_preview["dispatch_gate"] == GUEST_BOUNDARY, "dispatch preview must preserve the managed guest boundary")
assert(dispatch_preview["runner_lane"] == "known-app-guest-smoke", "dispatch preview must target the known app guest smoke lane")
assert(dispatch_preview["dispatch_started"] == false, "dispatch preview must not start dispatch")
assert(dispatch_preview["execution_started"] == false, "dispatch preview must not start execution")
assert(dispatch_preview["host_root_modified"] == false, "dispatch preview must not mutate the host root")
assert_no_forbidden(dispatch_preview_stdout, [PROJECT_ROOT.to_s, ".exe", "wine ", "wine/", ".wine", "qemu-system"], "dispatch preview output")

unless dispatch_preview["dispatch_ready"] == true && dispatch_preview["artifact_verified"] == true
  cache_status = dispatch_preview.fetch("cache_status", "unknown")
  if cache_status == "checksum-mismatch"
    warn dispatch_preview_stdout
    warn "FAIL: #{SMOKE_NAME} requires a checksum-verified known Windows app artifact"
    exit 1
  end

  puts "SKIP: #{SMOKE_NAME} (known Windows app artifact unavailable; run `ruby scripts/container.rb fetch-known-winapp` first)"
  exit 0
end

unless File.file?(Xnix::SshTestKey::PRIVATE_KEY_PATH)
  puts "SKIP: #{SMOKE_NAME} (SSH test key unavailable; run prepare-ssh-test-key and start-build-ssh-wine-guest first)"
  exit 0
end

qemu = Xnix::Qemu.wine_guest

unless File.file?(qemu.kernel_image)
  puts "SKIP: #{SMOKE_NAME} (Wine guest kernel unavailable; run configure-wine-guest and start-build-ssh-wine-guest first)"
  exit 0
end

build_stdout, build_stderr, build_status = run_command(go_env, "go", "build", "-o", LAUNCHER_BIN.to_s, "./cmd/xnix-compat-launch")
assert(build_status.zero?, "#{GO_BUILD_STEP} must exit successfully: #{build_stderr}\n#{build_stdout}")

stage, stage_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "desktop-activation-stage",
  "--registry", "runtime/recipes/registry.json",
  "--app", "org.xnix.sample.notepad",
  "--mode", "development",
  "--staging-root", STAGE_ROOT.to_s,
  "--managed-launcher-bin", LAUNCHER_BIN.to_s
)
assert(stage["written_file_count"] == 7, "desktop activation stage must copy the managed launcher")
assert(stage.fetch("written_file_ids").include?("managed-launcher-executable"), "desktop activation stage must report managed-launcher-executable")
assert(stage["launch_enabled"] == false, "desktop activation stage must not enable desktop launch")
assert(stage["execution_started"] == false, "desktop activation stage must not start execution")
assert(stage["host_root_modified"] == false, "desktop activation stage must not mutate the host root")
assert_no_forbidden(stage_stdout, [PROJECT_ROOT.to_s, STAGE_ROOT.to_s, LAUNCHER_BIN.to_s], "desktop activation stage output")
assert(STAGED_LAUNCHER.file?, "staged launcher executable must exist")
assert((STAGED_LAUNCHER.stat.mode & 0o777) == 0o755, "staged launcher executable must use mode 0755")

receipt_preview, receipt_preview_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-launch-authorization-receipt-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--authorize", "review-launch-authorization"
)
assert(receipt_preview["request_type"] == "known-app-launch-authorization-receipt-preview", "launch authorization receipt request type must match")
assert(receipt_preview["receipt_written"] == true, "launch authorization receipt must be written")
assert(receipt_preview["launch_authorization_recorded"] == true, "launch authorization receipt must be recorded")
assert(receipt_preview["receipt_path_exposed"] == false, "launch authorization receipt output must not expose receipt paths")
assert(receipt_preview["state_root_path_exposed"] == false, "launch authorization receipt output must not expose the state root path")
assert(receipt_preview["direct_launch_enabled"] == false, "launch authorization receipt must not enable direct launch")
assert(receipt_preview["desktop_launch_enabled"] == false, "launch authorization receipt must not enable desktop launch")
assert(receipt_preview["backend_launch_enabled"] == false, "launch authorization receipt must not enable backend launch")
assert(receipt_preview["backend_process_started"] == false, "launch authorization receipt must not start backend processes")
assert(receipt_preview["host_root_modified"] == false, "launch authorization receipt must not mutate the host root")
assert_no_forbidden(receipt_preview_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "launch authorization receipt output")

launch_gate, launch_gate_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-launch-gate-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--receipt-id", receipt_preview.fetch("receipt_id"),
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(launch_gate["request_type"] == "known-app-launch-gate-preview", "launch gate request type must match")
assert(launch_gate["receipt_lookup_state"] == "accepted-receipt", "launch gate must accept the Runtime authorization receipt")
assert(launch_gate["receipt_accepted"] == true, "launch gate must mark the receipt accepted")
assert(launch_gate["guest_boundary_accepted"] == true, "launch gate must accept the controlled guest boundary")
assert(launch_gate["launch_gate_state"] == "controlled-dispatch-ready", "launch gate must reach controlled dispatch readiness before staged dispatch")
assert(launch_gate["controlled_dispatch_ready"] == true, "launch gate must report controlled dispatch readiness")
assert(launch_gate["dispatch_request_materialized"] == true, "launch gate must materialize the controlled dispatch request")
assert(launch_gate["direct_launch_enabled"] == false, "launch gate must not enable direct launch")
assert(launch_gate["desktop_launch_enabled"] == false, "launch gate must not enable desktop launch")
assert(launch_gate["backend_launch_enabled"] == false, "launch gate must not enable backend launch")
assert(launch_gate["backend_process_started"] == false, "launch gate must not start backend processes")
assert(launch_gate["execution_started"] == false, "launch gate preview must not start execution")
assert(launch_gate["host_root_modified"] == false, "launch gate must not mutate the host root")
assert(launch_gate["receipt_path_exposed"] == false, "launch gate output must not expose receipt paths")
assert(launch_gate["state_root_path_exposed"] == false, "launch gate output must not expose the state root path")
assert_no_forbidden(launch_gate_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "launch gate output")

controlled_dispatch, controlled_dispatch_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-controlled-dispatch-request-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--receipt-id", receipt_preview.fetch("receipt_id"),
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(controlled_dispatch["request_type"] == "known-app-controlled-dispatch-request-preview", "controlled dispatch request type must match")
assert(controlled_dispatch["receipt_accepted"] == true, "controlled dispatch request must require accepted receipt evidence")
assert(controlled_dispatch["guest_boundary_accepted"] == true, "controlled dispatch request must require the accepted guest boundary")
assert(controlled_dispatch["launch_gate_state"] == "controlled-dispatch-ready", "controlled dispatch request must consume a ready launch gate")
assert(controlled_dispatch["controlled_dispatch_ready"] == true, "controlled dispatch request must require controlled dispatch readiness")
assert(controlled_dispatch["controlled_dispatch_request_created"] == true, "controlled dispatch request must be materialized before staged dispatch")
assert(controlled_dispatch["controlled_dispatch_request_state"] == "created", "controlled dispatch request must report created state")
assert(controlled_dispatch["runtime_owned_dispatch_request"] == true, "controlled dispatch request must be Runtime-owned")
assert(controlled_dispatch["dispatch_request_type"] == "windows-known-app-dispatch-preview", "controlled dispatch request must preserve the dispatch preview type")
assert(controlled_dispatch["dispatch_smoke_request_type"] == "windows-known-app-dispatch-smoke", "controlled dispatch request must preserve the dispatch smoke request type")
assert(controlled_dispatch["dispatch_gate"] == GUEST_BOUNDARY, "controlled dispatch request must preserve the managed guest boundary")
assert(controlled_dispatch["artifact_verified"] == true, "controlled dispatch request must require verified managed artifact evidence")
assert(controlled_dispatch["dispatch_ready"] == true, "controlled dispatch request must require dispatch readiness")
assert(controlled_dispatch["dispatch_allowed"] == false, "controlled dispatch request preview must not allow dispatch directly")
assert(controlled_dispatch["dispatch_started"] == false, "controlled dispatch request preview must not start dispatch")
assert(controlled_dispatch["execution_started"] == false, "controlled dispatch request preview must not start execution")
assert(controlled_dispatch["direct_launch_enabled"] == false, "controlled dispatch request must not enable direct launch")
assert(controlled_dispatch["desktop_launch_enabled"] == false, "controlled dispatch request must not enable desktop launch")
assert(controlled_dispatch["backend_launch_enabled"] == false, "controlled dispatch request must not enable backend launch")
assert(controlled_dispatch["backend_process_started"] == false, "controlled dispatch request must not start backend processes")
assert(controlled_dispatch["host_root_modified"] == false, "controlled dispatch request must not mutate the host root")
assert(controlled_dispatch["receipt_path_exposed"] == false, "controlled dispatch request output must not expose receipt paths")
assert(controlled_dispatch["state_root_path_exposed"] == false, "controlled dispatch request output must not expose the state root path")
assert_no_forbidden(controlled_dispatch_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "controlled dispatch request output")

controlled_session, controlled_session_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-controlled-execution-session-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--receipt-id", receipt_preview.fetch("receipt_id"),
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(controlled_session["request_type"] == "known-app-controlled-execution-session-preview", "controlled execution session request type must match")
assert(controlled_session["receipt_accepted"] == true, "controlled execution session must require accepted receipt evidence")
assert(controlled_session["guest_boundary_accepted"] == true, "controlled execution session must require accepted guest boundary")
assert(controlled_session["controlled_dispatch_request_created"] == true, "controlled execution session must require a controlled dispatch request")
assert(controlled_session["runtime_owned_dispatch_request"] == true, "controlled execution session must preserve Runtime-owned dispatch evidence")
assert(controlled_session["runtime_owned_execution_session"] == true, "controlled execution session must be Runtime-owned")
assert(controlled_session["execution_session_request_type"] == "known-app-runtime-execution-session-handoff", "controlled execution session must expose the handoff request type")
assert(controlled_session["execution_session_state"] == "handoff-created", "controlled execution session must create the handoff before staged dispatch")
assert(controlled_session["execution_session_handoff_created"] == true, "controlled execution session must report handoff creation")
assert(controlled_session["execution_session_portable"] == true, "controlled execution session must be portable")
assert(controlled_session["session_handoff_ready"] == true, "controlled execution session must be ready for the Runtime-managed runner")
assert(controlled_session["session_registered"] == false, "controlled execution session preview must not register live sessions")
assert(controlled_session["window_observed"] == false, "controlled execution session preview must not observe windows")
assert(controlled_session["dispatch_allowed"] == false, "controlled execution session preview must not allow dispatch directly")
assert(controlled_session["dispatch_started"] == false, "controlled execution session preview must not start dispatch")
assert(controlled_session["execution_started"] == false, "controlled execution session preview must not start execution")
assert(controlled_session["direct_launch_enabled"] == false, "controlled execution session must not enable direct launch")
assert(controlled_session["desktop_launch_enabled"] == false, "controlled execution session must not enable desktop launch")
assert(controlled_session["backend_launch_enabled"] == false, "controlled execution session must not enable backend launch")
assert(controlled_session["backend_process_started"] == false, "controlled execution session must not start backend processes")
assert(controlled_session["host_root_modified"] == false, "controlled execution session must not mutate the host root")
assert(controlled_session["receipt_path_exposed"] == false, "controlled execution session output must not expose receipt paths")
assert(controlled_session["state_root_path_exposed"] == false, "controlled execution session output must not expose the state root path")
assert_no_forbidden(controlled_session_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "controlled execution session output")

controlled_session_record, controlled_session_record_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-controlled-execution-session-record",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--receipt-id", receipt_preview.fetch("receipt_id"),
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(controlled_session_record["request_type"] == "known-app-controlled-execution-session-record", "controlled execution session record request type must match")
assert(controlled_session_record["record_state"] == "persisted", "controlled execution session record must persist before staged dispatch")
assert(controlled_session_record["ledger_record_written"] == true, "controlled execution session record must write the ledger transaction")
assert(controlled_session_record["session_record_written"] == true, "controlled execution session record must write the session record")
assert(controlled_session_record["runtime_owned_execution_session"] == true, "controlled execution session record must be Runtime-owned")
assert(controlled_session_record["execution_session_handoff_created"] == true, "controlled execution session record must require handoff creation")
assert(controlled_session_record["session_handoff_ready"] == true, "controlled execution session record must require handoff readiness")
assert(controlled_session_record["transaction_relative_path"].start_with?("execution-ledger/transactions/"), "controlled execution session record must expose only relative transaction evidence")
assert(controlled_session_record["session_relative_path"].start_with?("execution-ledger/sessions/"), "controlled execution session record must expose only relative session evidence")
assert(controlled_session_record["session_sha256"].to_s.length == 64, "controlled execution session record must expose digest evidence")
assert(controlled_session_record["session_registered"] == false, "controlled execution session record must not register live sessions")
assert(controlled_session_record["window_observed"] == false, "controlled execution session record must not observe windows")
assert(controlled_session_record["dispatch_started"] == false, "controlled execution session record must not start dispatch")
assert(controlled_session_record["execution_started"] == false, "controlled execution session record must not start execution")
assert(controlled_session_record["backend_process_started"] == false, "controlled execution session record must not start backend processes")
assert(controlled_session_record["host_root_modified"] == false, "controlled execution session record must not mutate the host root")
assert(controlled_session_record["state_root_path_exposed"] == false, "controlled execution session record output must not expose the state root path")
assert(controlled_session_record["transaction_path_exposed"] == false, "controlled execution session record output must not expose transaction paths")
assert(controlled_session_record["session_path_exposed"] == false, "controlled execution session record output must not expose session paths")
assert_no_forbidden(controlled_session_record_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "controlled execution session record output")

controlled_session_consume, controlled_session_consume_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-controlled-execution-session-consume-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--session-id", controlled_session_record.fetch("execution_session_id")
)
assert(controlled_session_consume["request_type"] == "known-app-controlled-execution-session-consume-preview", "controlled execution session consumption request type must match")
assert(controlled_session_consume["execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "controlled execution session consumption must read the recorded session id")
assert(controlled_session_consume["record_consumed"] == true, "controlled execution session consumption must consume the record")
assert(controlled_session_consume["ledger_record_consumed"] == true, "controlled execution session consumption must consume the ledger record")
assert(controlled_session_consume["session_record_consumed"] == true, "controlled execution session consumption must consume the session record")
assert(controlled_session_consume["session_digest_verified"] == true, "controlled execution session consumption must verify the session digest")
assert(controlled_session_consume["session_relative_path"] == controlled_session_record.fetch("session_relative_path"), "controlled execution session consumption must preserve relative session evidence")
assert(controlled_session_consume["session_sha256"] == controlled_session_record.fetch("session_sha256"), "controlled execution session consumption must preserve digest evidence")
assert(controlled_session_consume["fan_out_request_type"] == "execution-session-fanout-evidence", "controlled execution session consumption must expose fan-out evidence")
assert(controlled_session_consume["surface_count"] == 4, "controlled execution session consumption must fan out to four KDE read-model consumers")
assert(controlled_session_consume["runtime_owner_consumable"] == true, "controlled execution session consumption must be Runtime-owner consumable")
assert(controlled_session_consume["kde_read_model_consumable"] == true, "controlled execution session consumption must be KDE read-model consumable")
assert(controlled_session_consume["safe_for_kde"] == true, "controlled execution session consumption must be safe for KDE")
assert(controlled_session_consume["task_manager_state"] == "blocked", "controlled execution session consumption must expose task-manager state")
assert(controlled_session_consume["kwin_state"] == "blocked", "controlled execution session consumption must expose KWin state")
assert(controlled_session_consume["tray_state"] == "blocked", "controlled execution session consumption must expose tray state")
assert(controlled_session_consume["compatibility_center_state"] == "waiting-for-runtime-gates", "controlled execution session consumption must expose Compatibility Center state")
assert(controlled_session_consume["session_registered"] == false, "controlled execution session consumption must not register live sessions")
assert(controlled_session_consume["window_observed"] == false, "controlled execution session consumption must not observe windows")
assert(controlled_session_consume["task_manager_entry_active"] == false, "controlled execution session consumption must not activate task-manager entries")
assert(controlled_session_consume["kwin_rule_applied"] == false, "controlled execution session consumption must not apply KWin rules")
assert(controlled_session_consume["live_tray_bridge_enabled"] == false, "controlled execution session consumption must not enable live tray bridges")
assert(controlled_session_consume["dispatch_started"] == false, "controlled execution session consumption must not start dispatch")
assert(controlled_session_consume["execution_started"] == false, "controlled execution session consumption must not start execution")
assert(controlled_session_consume["direct_launch_enabled"] == false, "controlled execution session consumption must not enable direct launch")
assert(controlled_session_consume["desktop_launch_enabled"] == false, "controlled execution session consumption must not enable desktop launch")
assert(controlled_session_consume["backend_process_started"] == false, "controlled execution session consumption must not start backend processes")
assert(controlled_session_consume["host_root_modified"] == false, "controlled execution session consumption must not mutate the host root")
assert(controlled_session_consume["state_root_path_exposed"] == false, "controlled execution session consumption output must not expose the state root path")
assert(controlled_session_consume["transaction_path_exposed"] == false, "controlled execution session consumption output must not expose transaction paths")
assert(controlled_session_consume["session_path_exposed"] == false, "controlled execution session consumption output must not expose session paths")
assert_no_forbidden(controlled_session_consume_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "controlled execution session consumption output")

launch_review, launch_review_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-session-gated-launch-review-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--session-id", controlled_session_record.fetch("execution_session_id"),
  "--action", "review-session-gated-dispatch",
  "--decision", "approved"
)
assert(launch_review["request_type"] == "known-app-session-gated-launch-review-preview", "Runtime launch review route must be session-gated")
assert(launch_review["source"] == "known-app-controlled-execution-session-consume-preview+kde-center-page-session-gate-card", "Runtime launch review route must consume the controlled session evidence")
assert(launch_review["action_id"] == "review-session-gated-dispatch", "Runtime launch review route must preserve the KDE card action")
assert(launch_review["review_route_created"] == true, "Runtime launch review route must be available")
assert(launch_review["read_before_write_required"] == true, "Runtime launch review route must stay read-before-write")
assert(launch_review["runtime_receipt_required"] == true, "Runtime launch review route must require a later Runtime receipt")
assert(launch_review["execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "Runtime launch review route must preserve the opaque session id")
assert(launch_review["session_record_consumed"] == true, "Runtime launch review route must consume the session record")
assert(launch_review["session_digest_verified"] == true, "Runtime launch review route must verify the session digest")
assert(launch_review["session_relative_path"] == controlled_session_record.fetch("session_relative_path"), "Runtime launch review route must expose only relative session evidence")
assert(launch_review["runtime_owner_consumable"] == true, "Runtime launch review route must be Runtime-owner consumable")
assert(launch_review["kde_read_model_consumable"] == true, "Runtime launch review route must be KDE read-model consumable")
assert(launch_review["runtime_launch_approval"] == false, "Runtime launch review route must not grant launch approval")
assert(launch_review["desktop_launch_enabled"] == false, "Runtime launch review route must not enable desktop launch")
assert(launch_review["backend_launch_enabled"] == false, "Runtime launch review route must not enable backend launch")
assert(launch_review["execution_started"] == false, "Runtime launch review route must not start execution")
assert(launch_review["review_receipt_recorded"] == false, "Runtime launch review route must not record a receipt during preview")
assert(launch_review["host_root_modified"] == false, "Runtime launch review route must not mutate the host root")
assert_no_forbidden(launch_review_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime launch review route output")

launch_review_receipt, launch_review_receipt_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-session-gated-launch-review-receipt-record",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--session-id", controlled_session_record.fetch("execution_session_id"),
  "--action", "review-session-gated-dispatch",
  "--decision", "approved"
)
assert(launch_review_receipt["request_type"] == "known-app-session-gated-launch-review-receipt-record", "Runtime launch review receipt request type must match")
assert(launch_review_receipt["source"] == "known-app-session-gated-launch-review-preview+runtime-review-receipt-store", "Runtime launch review receipt must consume the preview route")
assert(launch_review_receipt["runtime_method"] == "RecordKnownAppSessionGatedLaunchReviewReceipt", "Runtime launch review receipt must be written by the Go Runtime")
assert(launch_review_receipt["read_method"] == "GetKnownAppSessionGatedLaunchReviewReceipt", "Runtime launch review receipt must expose a stable read method")
assert(launch_review_receipt["action_id"] == "review-session-gated-dispatch", "Runtime launch review receipt must preserve the KDE card action")
assert(launch_review_receipt["decision"] == "approved", "Runtime launch review receipt must preserve the explicit decision")
assert(launch_review_receipt["read_before_write_consumed"] == true, "Runtime launch review receipt must consume read-before-write evidence")
assert(launch_review_receipt["review_route_consumed"] == true, "Runtime launch review receipt must consume the review route")
assert(launch_review_receipt["runtime_receipt_required"] == true, "Runtime launch review receipt must still require a later launch gate")
assert(launch_review_receipt["execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "Runtime launch review receipt must preserve the opaque session id")
assert(launch_review_receipt["session_record_consumed"] == true, "Runtime launch review receipt must consume the session record")
assert(launch_review_receipt["session_digest_verified"] == true, "Runtime launch review receipt must verify the session digest")
assert(launch_review_receipt["session_relative_path"] == controlled_session_record.fetch("session_relative_path"), "Runtime launch review receipt must expose only relative session evidence")
assert(launch_review_receipt["receipt_id"].is_a?(String) && !launch_review_receipt["receipt_id"].empty?, "Runtime launch review receipt must expose an opaque receipt id")
assert(launch_review_receipt["receipt_relative_path"].is_a?(String) && launch_review_receipt["receipt_relative_path"].start_with?("runtime/session-gated-launch-review-receipts/"), "Runtime launch review receipt must expose only relative receipt evidence")
assert(launch_review_receipt["receipt_sha256"].is_a?(String) && launch_review_receipt["receipt_sha256"].length == 64, "Runtime launch review receipt must expose a receipt digest")
assert(launch_review_receipt["review_receipt_recorded"] == true, "Runtime launch review receipt must be recorded")
assert(launch_review_receipt["runtime_owner_consumable"] == true, "Runtime launch review receipt must be Runtime-owner consumable")
assert(launch_review_receipt["kde_read_model_consumable"] == true, "Runtime launch review receipt must be KDE read-model consumable")
assert(launch_review_receipt["runtime_launch_approval"] == false, "Runtime launch review receipt must not grant launch approval")
assert(launch_review_receipt["launch_allowed"] == false, "Runtime launch review receipt must not allow launch")
assert(launch_review_receipt["desktop_launch_enabled"] == false, "Runtime launch review receipt must not enable desktop launch")
assert(launch_review_receipt["backend_launch_enabled"] == false, "Runtime launch review receipt must not enable backend launch")
assert(launch_review_receipt["execution_started"] == false, "Runtime launch review receipt must not start execution")
assert(launch_review_receipt["request_objects_created"] == false, "Runtime launch review receipt must not create request objects")
assert(launch_review_receipt["permission_grant_created"] == false, "Runtime launch review receipt must not create permission grants")
assert(launch_review_receipt["receipt_path_exposed"] == false, "Runtime launch review receipt must not expose receipt paths")
assert(launch_review_receipt["state_root_path_exposed"] == false, "Runtime launch review receipt must not expose the state root")
assert(launch_review_receipt["host_root_modified"] == false, "Runtime launch review receipt must not mutate the host root")
assert_no_forbidden(launch_review_receipt_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime launch review receipt output")

launch_review_gate, launch_review_gate_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-session-gated-launch-review-gate-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--session-id", controlled_session_record.fetch("execution_session_id"),
  "--receipt-id", launch_review_receipt.fetch("receipt_id")
)
assert(launch_review_gate["request_type"] == "known-app-session-gated-launch-review-gate-preview", "Runtime launch review gate request type must match")
assert(launch_review_gate["source"] == "known-app-session-gated-launch-review-receipt-record+runtime-review-receipt-gate", "Runtime launch review gate must consume the review receipt")
assert(launch_review_gate["runtime_method"] == "PreviewKnownAppSessionGatedLaunchReviewGate", "Runtime launch review gate must be owned by the Go Runtime")
assert(launch_review_gate["read_method"] == "GetKnownAppSessionGatedLaunchReviewGate", "Runtime launch review gate must expose a stable read method")
assert(launch_review_gate["execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "Runtime launch review gate must preserve the opaque session id")
assert(launch_review_gate["session_record_consumed"] == true, "Runtime launch review gate must re-consume the session record")
assert(launch_review_gate["session_digest_verified"] == true, "Runtime launch review gate must verify the session digest")
assert(launch_review_gate["read_before_write_revalidated"] == true, "Runtime launch review gate must revalidate read-before-write evidence")
assert(launch_review_gate["receipt_id"] == launch_review_receipt.fetch("receipt_id"), "Runtime launch review gate must consume the selected receipt id")
assert(launch_review_gate["receipt_relative_path"] == launch_review_receipt.fetch("receipt_relative_path"), "Runtime launch review gate must expose only relative receipt evidence")
assert(launch_review_gate["receipt_sha256"] == launch_review_receipt.fetch("receipt_sha256"), "Runtime launch review gate must verify the receipt digest")
assert(launch_review_gate["receipt_lookup_state"] == "accepted-receipt", "Runtime launch review gate must accept the review receipt")
assert(launch_review_gate["receipt_decision"] == "approved", "Runtime launch review gate must accept only approved receipts")
assert(launch_review_gate["review_receipt_consumed"] == true, "Runtime launch review gate must consume the review receipt")
assert(launch_review_gate["review_receipt_accepted"] == true, "Runtime launch review gate must accept the review receipt")
assert(launch_review_gate["review_gate_ready"] == true, "Runtime launch review gate must be ready")
assert(launch_review_gate["controlled_dispatch_gate_ready"] == true, "Runtime launch review gate must prepare the controlled dispatch gate")
assert(launch_review_gate["dispatch_state_advance_ready"] == true, "Runtime launch review gate must allow the next dispatch state to advance")
assert(launch_review_gate["dispatch_state_advanced"] == false, "Runtime launch review gate must not advance dispatch state itself")
assert(launch_review_gate["runtime_launch_approval"] == false, "Runtime launch review gate must not grant direct launch approval")
assert(launch_review_gate["launch_allowed"] == false, "Runtime launch review gate must not allow direct launch")
assert(launch_review_gate["desktop_launch_enabled"] == false, "Runtime launch review gate must not enable desktop launch")
assert(launch_review_gate["backend_launch_enabled"] == false, "Runtime launch review gate must not enable backend launch")
assert(launch_review_gate["execution_started"] == false, "Runtime launch review gate must not start execution")
assert(launch_review_gate["controlled_dispatch_request_created"] == false, "Runtime launch review gate must not create dispatch requests")
assert(launch_review_gate["permission_grant_created"] == false, "Runtime launch review gate must not create permission grants")
assert(launch_review_gate["receipt_path_exposed"] == false, "Runtime launch review gate must not expose receipt paths")
assert(launch_review_gate["state_root_path_exposed"] == false, "Runtime launch review gate must not expose the state root")
assert(launch_review_gate["host_root_modified"] == false, "Runtime launch review gate must not mutate the host root")
assert_no_forbidden(launch_review_gate_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime launch review gate output")

session_gated_dispatch, session_gated_dispatch_stdout = run_json(
  go_env,
  "go", "run", "./cmd/xnix-runtime-go",
  "known-app-session-gated-controlled-dispatch-request-preview",
  "--app", APP_ID,
  "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
  "--session-id", controlled_session_record.fetch("execution_session_id"),
  "--review-receipt-id", launch_review_receipt.fetch("receipt_id"),
  "--launch-receipt-id", receipt_preview.fetch("receipt_id"),
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--guest-boundary", GUEST_BOUNDARY
)
assert(session_gated_dispatch["request_type"] == "known-app-session-gated-controlled-dispatch-request-preview", "post-review controlled dispatch request type must match")
assert(session_gated_dispatch["source"] == "known-app-session-gated-launch-review-gate-preview+known-app-controlled-dispatch-request-preview", "post-review controlled dispatch must consume the session-gated review gate")
assert(session_gated_dispatch["runtime_method"] == "PreviewKnownAppSessionGatedControlledDispatchRequest", "post-review controlled dispatch must be owned by the Go Runtime")
assert(session_gated_dispatch["read_method"] == "GetKnownAppSessionGatedControlledDispatchRequest", "post-review controlled dispatch must expose a stable read method")
assert(session_gated_dispatch["execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "post-review controlled dispatch must preserve the opaque session id")
assert(session_gated_dispatch["session_record_consumed"] == true, "post-review controlled dispatch must re-consume the session record")
assert(session_gated_dispatch["session_digest_verified"] == true, "post-review controlled dispatch must verify the session digest")
assert(session_gated_dispatch["session_relative_path"] == controlled_session_record.fetch("session_relative_path"), "post-review controlled dispatch must expose only relative session evidence")
assert(session_gated_dispatch["review_receipt_id"] == launch_review_receipt.fetch("receipt_id"), "post-review controlled dispatch must consume the selected review receipt")
assert(session_gated_dispatch["review_receipt_relative_path"] == launch_review_receipt.fetch("receipt_relative_path"), "post-review controlled dispatch must expose only relative review receipt evidence")
assert(session_gated_dispatch["review_receipt_sha256"] == launch_review_receipt.fetch("receipt_sha256"), "post-review controlled dispatch must verify the review receipt digest")
assert(session_gated_dispatch["review_receipt_consumed"] == true, "post-review controlled dispatch must consume the review receipt")
assert(session_gated_dispatch["review_receipt_accepted"] == true, "post-review controlled dispatch must accept the review receipt")
assert(session_gated_dispatch["review_gate_ready"] == true, "post-review controlled dispatch must require the review gate")
assert(session_gated_dispatch["dispatch_state_advance_ready"] == true, "post-review controlled dispatch must require review-gated dispatch advance readiness")
assert(session_gated_dispatch["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "post-review controlled dispatch must preserve the launch authorization receipt id")
assert(session_gated_dispatch["launch_gate_state"] == "controlled-dispatch-ready", "post-review controlled dispatch must consume a ready launch gate")
assert(session_gated_dispatch["launch_gate_receipt_accepted"] == true, "post-review controlled dispatch must accept the launch authorization receipt")
assert(session_gated_dispatch["launch_gate_guest_boundary_accepted"] == true, "post-review controlled dispatch must accept the managed guest boundary")
assert(session_gated_dispatch["controlled_dispatch_gate_ready"] == true, "post-review controlled dispatch must require controlled dispatch readiness")
assert(session_gated_dispatch["controlled_dispatch_request_created"] == true, "post-review controlled dispatch must create the controlled dispatch request state")
assert(session_gated_dispatch["controlled_dispatch_request_state"] == "created-after-session-gated-review", "post-review controlled dispatch must report session-gated created state")
assert(session_gated_dispatch["runtime_owned_dispatch_request"] == true, "post-review controlled dispatch request must be Runtime-owned")
assert(session_gated_dispatch["dispatch_request_type"] == "windows-known-app-dispatch-preview", "post-review controlled dispatch must preserve the dispatch preview type")
assert(session_gated_dispatch["dispatch_smoke_request_type"] == "windows-known-app-dispatch-smoke", "post-review controlled dispatch must preserve the dispatch smoke request type")
assert(session_gated_dispatch["dispatch_gate"] == GUEST_BOUNDARY, "post-review controlled dispatch must preserve the managed guest boundary")
assert(session_gated_dispatch["artifact_verified"] == true, "post-review controlled dispatch must require verified managed artifact evidence")
assert(session_gated_dispatch["dispatch_ready"] == true, "post-review controlled dispatch must require dispatch readiness")
assert(session_gated_dispatch["request_objects_created"] == true, "post-review controlled dispatch must create only controlled request state")
assert(session_gated_dispatch["dispatch_allowed"] == false, "post-review controlled dispatch preview must not allow dispatch directly")
assert(session_gated_dispatch["dispatch_started"] == false, "post-review controlled dispatch preview must not start dispatch")
assert(session_gated_dispatch["execution_started"] == false, "post-review controlled dispatch preview must not start execution")
assert(session_gated_dispatch["direct_launch_enabled"] == false, "post-review controlled dispatch must not enable direct launch")
assert(session_gated_dispatch["desktop_launch_enabled"] == false, "post-review controlled dispatch must not enable desktop launch")
assert(session_gated_dispatch["backend_launch_enabled"] == false, "post-review controlled dispatch must not enable backend launch")
assert(session_gated_dispatch["backend_process_started"] == false, "post-review controlled dispatch must not start backend processes")
assert(session_gated_dispatch["permission_grant_created"] == false, "post-review controlled dispatch must not create permission grants")
assert(session_gated_dispatch["review_receipt_path_exposed"] == false, "post-review controlled dispatch must not expose review receipt paths")
assert(session_gated_dispatch["receipt_path_exposed"] == false, "post-review controlled dispatch must not expose launch receipt paths")
assert(session_gated_dispatch["state_root_path_exposed"] == false, "post-review controlled dispatch must not expose the state root")
assert(session_gated_dispatch["session_path_exposed"] == false, "post-review controlled dispatch must not expose session paths")
assert(session_gated_dispatch["host_root_modified"] == false, "post-review controlled dispatch must not mutate the host root")
assert_no_forbidden(session_gated_dispatch_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "post-review controlled dispatch output")

FileUtils.rm_f(SERIAL_LOG_PATH)
stdin, output, wait_thread = Open3.popen2e(*qemu.boot_command(ssh: true))
stdin.close
serial_log = +""
reader = Thread.new { serial_log = output.read }
probe = Xnix::SshProbe.new
deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + BOOT_TIMEOUT_SECONDS

begin
  loop do
    break if system(*probe.command, out: File::NULL, err: File::NULL)

    abort "QEMU exited before SSH became ready" unless wait_thread.alive?
    abort "#{SMOKE_NAME} timed out waiting for SSH" if Process.clock_gettime(Process::CLOCK_MONOTONIC) >= deadline

    sleep RETRY_INTERVAL_SECONDS
  end

  launcher_stdout, launcher_stderr, launcher_status = run_command(
    go_env,
    "go", "run", "./cmd/xnix-runtime-go",
    "known-app-kde-runtime-status-launch-execution",
    "--app", APP_ID,
    "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--launch-authorization-receipt-id", receipt_preview.fetch("receipt_id"),
    "--session-gated-review-receipt-id", launch_review_receipt.fetch("receipt_id"),
    "--session-id", controlled_session_record.fetch("execution_session_id"),
    "--center-card-state", "validated-post-review-dispatch",
    "--primary-action-id", "show-runtime-controlled-launch",
    "--post-review-dispatch-state", "created-after-session-gated-review",
    "--launcher", STAGED_LAUNCHER.to_s,
    "--key", Xnix::SshTestKey::PRIVATE_KEY_PATH,
    "--timeout", ENV.fetch("XNIX_KNOWN_WINAPP_GUEST_TIMEOUT", "90s"),
    *APP_ARGS.flat_map { |argument| ["--arg", argument] }
  )

  unless launcher_status.zero?
    warn "QEMU serial log: #{SERIAL_LOG_PATH}"
    warn launcher_stdout unless launcher_stdout.empty?
    warn launcher_stderr unless launcher_stderr.empty?
    warn "FAIL: #{SMOKE_NAME} command failed"
    exit 1
  end

  runtime_execution = JSON.parse(launcher_stdout)
  assert(runtime_execution["request_type"] == "known-app-kde-runtime-status-launch-execution", "Runtime status launch execution must use the Runtime-owned entrypoint")
  assert(runtime_execution["runtime_method"] == "PrepareKnownAppKDERuntimeStatusLaunchExecution", "Runtime status launch execution must prepare through Go Runtime")
  assert(runtime_execution["execution_method"] == "RunKnownAppKDERuntimeStatusLaunchExecution", "Runtime status launch execution must own the execution bridge")
  assert(runtime_execution["request_preview_type"] == "known-app-kde-runtime-status-launch-request-preview", "Runtime status launch execution must consume the KDE request preview")
  assert(runtime_execution["state_root_injected_by_runtime"] == true, "Runtime status launch execution must inject state-root internally")
  assert(runtime_execution["state_root_supplied_by_runtime"] == true, "Runtime status launch execution must keep state-root supplied by Runtime")
  assert(runtime_execution["kde_state_root_access"] == false, "Runtime status launch execution must not give KDE state-root access")
  assert(runtime_execution["launch_receipt_revalidated"] == true, "Runtime status launch execution must revalidate launch receipt")
  assert(runtime_execution["review_receipt_revalidated"] == true, "Runtime status launch execution must revalidate review receipt")
  assert(runtime_execution["controlled_session_revalidated"] == true, "Runtime status launch execution must revalidate controlled session")
  assert(runtime_execution["guest_boundary_revalidated"] == true, "Runtime status launch execution must revalidate the managed guest boundary")
  assert(runtime_execution["managed_launcher_invoked"] == true, "Runtime status launch execution must invoke the managed launcher")
  assert(runtime_execution["existing_managed_launcher_invoked"] == true, "Runtime status launch execution must invoke the existing staged launcher")
  assert(runtime_execution["launcher_output_json_observed"] == true, "Runtime status launch execution must observe delegated launcher JSON")
  assert(runtime_execution["raw_launcher_output_exposed"] == false, "Runtime status launch execution must not expose raw launcher output")
  assert(runtime_execution["delegated_request_type"] == "windows-known-app-dispatch-smoke", "Runtime status launch execution must delegate to the dispatch smoke path")
  assert(runtime_execution["delegated_guest_boundary"] == GUEST_BOUNDARY, "Runtime status launch execution must preserve delegated guest boundary")
  assert(runtime_execution["delegated_runtime_owned_dispatch"] == true, "Runtime status launch execution must preserve Runtime-owned delegated dispatch")
  assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, STAGED_LAUNCHER.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime status launch execution output")

  payload = runtime_execution.fetch("compatibility_center_known_app_evidence")
  assert(payload["projection_type"] == "known-app-kde-runtime-status-launch-delegated-evidence", "Runtime status launch execution must provide a Center-ready delegated evidence projection")
  assert(payload["compatibility_center_projection_ready"] == true, "Runtime status launch execution projection must be ready for Compatibility Center")
  assert(payload["kde_center_projection_ready"] == true, "Runtime status launch execution projection must be ready for KDE Center")
  assert(payload["request_type"] == "windows-known-app-dispatch-smoke", "staged launcher must enter dispatch smoke with the guest boundary")
  assert(payload["guest_boundary"] == GUEST_BOUNDARY, "staged launcher dispatch must preserve the guest boundary")
  assert(payload["runtime_owned_dispatch"] == true, "staged launcher dispatch must be Runtime-owned")
  assert(payload["session_gated_controlled_dispatch_consumed"] == true, "staged launcher dispatch must consume the post-review controlled dispatch state")
  assert(payload["session_gated_controlled_dispatch_state"] == "created-after-session-gated-review", "staged launcher dispatch must preserve the post-review controlled dispatch state")
  assert(payload["session_gated_review_receipt_id"] == launch_review_receipt.fetch("receipt_id"), "staged launcher dispatch must preserve the opaque review receipt id")
  assert(payload["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "staged launcher dispatch must preserve the opaque launch authorization receipt id")
  assert(payload["controlled_execution_session_consumed"] == true, "staged launcher dispatch must consume the controlled execution session before dispatch")
  assert(payload["controlled_execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "staged launcher dispatch must preserve the controlled execution session id")
  assert(payload["controlled_session_digest_verified"] == true, "staged launcher dispatch must verify the controlled execution session digest")
  assert(payload["controlled_session_relative_path"] == controlled_session_record.fetch("session_relative_path"), "staged launcher dispatch must preserve relative controlled session evidence")
  assert(payload["runtime_owner_consumable_session"] == true, "staged launcher dispatch must require Runtime-owner consumable session evidence")
  assert(payload["kde_read_model_consumable_session"] == true, "staged launcher dispatch must require KDE read-model consumable session evidence")
  assert(payload["controlled_session_live_state_observed"] == false, "staged launcher dispatch must not observe live session state before dispatch")
  assert(payload["controlled_session_registered"] == false, "staged launcher dispatch must not register a live session before dispatch")
  assert(payload["controlled_session_window_observed"] == false, "staged launcher dispatch must not observe windows before dispatch")
  assert(payload["controlled_session_host_root_modified"] == false, "staged launcher dispatch session gate must not mutate the host root")
  assert(payload["controlled_session_backend_process_start"] == false, "staged launcher dispatch session gate must not start backend processes")
  assert(payload["host_root_modified"] == false, "staged launcher dispatch must not mutate the host root")
  assert(payload["docker_socket_mounted"] == false, "staged launcher dispatch must not mount the Docker socket")
  assert(payload["broad_host_mount_required"] == false, "staged launcher dispatch must not require broad host mounts")
  assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "staged launcher dispatch output")
  File.write(RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH, JSON.pretty_generate(payload))
  evidence_record, evidence_record_stdout = run_json(
    go_env,
    "go", "run", "./cmd/xnix-runtime-go",
    "known-app-kde-runtime-status-launch-evidence-record",
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--evidence-file", RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s
  )
  assert(evidence_record["request_type"] == "known-app-kde-runtime-status-launch-evidence-record", "Runtime status launch evidence record must use the Go-owned record type")
  assert(evidence_record["runtime_method"] == "RecordKnownAppKDERuntimeStatusLaunchEvidence", "Runtime status launch evidence record must be written by Go Runtime")
  assert(evidence_record["read_method"] == "GetKnownAppKDERuntimeStatusLaunchEvidence", "Runtime status launch evidence record must expose the KDE-safe read method")
  assert(evidence_record["evidence_state"] == "persisted", "Runtime status launch evidence record must persist the handoff")
  assert(evidence_record["evidence_relative_path"].to_s.start_with?("runtime/kde-runtime-status-launch-evidence/"), "Runtime status launch evidence record must expose only relative evidence path")
  assert(evidence_record["evidence_sha256"].to_s.match?(/\A[0-9a-f]{64}\z/), "Runtime status launch evidence record must expose evidence digest")
  assert(evidence_record["projection_type"] == "known-app-kde-runtime-status-launch-delegated-evidence", "Runtime status launch evidence record must preserve projection type")
  assert(evidence_record["compatibility_center_projection_ready"] == true, "Runtime status launch evidence record must preserve Compatibility Center readiness")
  assert(evidence_record["kde_center_projection_ready"] == true, "Runtime status launch evidence record must preserve KDE Center readiness")
  assert(evidence_record["known_app_smoke_evidence_ready"] == true, "Runtime status launch evidence record must be convertible to known app smoke evidence")
  assert(evidence_record["runtime_owned"] == true, "Runtime status launch evidence record must be Runtime-owned")
  assert(evidence_record["runtime_owned_dispatch"] == true, "Runtime status launch evidence record must preserve Runtime-owned dispatch")
  assert(evidence_record["go_runtime_backed"] == true, "Runtime status launch evidence record must be Go Runtime backed")
  assert(evidence_record["kde_policy_owner"] == false, "Runtime status launch evidence record must not make KDE the policy owner")
  assert(evidence_record["state_root_path_exposed"] == false, "Runtime status launch evidence record must not expose state-root paths")
  assert(evidence_record["evidence_path_exposed"] == false, "Runtime status launch evidence record must not expose evidence absolute paths")
  assert(evidence_record["managed_launcher_path_exposed"] == false, "Runtime status launch evidence record must not expose launcher paths")
  assert(evidence_record["raw_launcher_output_exposed"] == false, "Runtime status launch evidence record must not expose launcher output")
  assert(evidence_record["backend_details_exposed"] == false, "Runtime status launch evidence record must not expose backend details")
  assert(evidence_record["host_root_modified"] == false, "Runtime status launch evidence record must not mutate host root")
  assert(evidence_record["docker_socket_mounted"] == false, "Runtime status launch evidence record must not mount Docker socket")
  assert(evidence_record["broad_host_mount_required"] == false, "Runtime status launch evidence record must not require broad host mounts")
  assert(evidence_record["desktop_launch_enabled"] == false, "Runtime status launch evidence record must not enable desktop launch")
  assert(evidence_record["backend_launch_enabled"] == false, "Runtime status launch evidence record must not enable backend launch")
  assert(evidence_record["execution_started"] == false, "Runtime status launch evidence record must not start execution")
  assert(evidence_record["backend_process_started"] == false, "Runtime status launch evidence record must not start backend processes")
  assert_no_forbidden(evidence_record_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime status launch evidence record output")
  evidence_preview, evidence_preview_stdout = run_json(
    go_env,
    "go", "run", "./cmd/xnix-runtime-go",
    "known-app-kde-runtime-status-launch-evidence-preview",
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--evidence-relative-path", evidence_record.fetch("evidence_relative_path")
  )
  assert(evidence_preview["request_type"] == "known-app-kde-runtime-status-launch-evidence-preview", "Runtime status launch evidence preview must use the Go-owned preview type")
  assert(evidence_preview["runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchEvidence", "Runtime status launch evidence preview must be read by Go Runtime")
  assert(evidence_preview["read_method"] == "GetKnownAppKDERuntimeStatusLaunchEvidence", "Runtime status launch evidence preview must expose the KDE-safe read method")
  assert(evidence_preview["evidence_read_state"] == "consumed", "Runtime status launch evidence preview must consume the persisted handoff")
  assert(evidence_preview["evidence_handoff_consumed"] == true, "Runtime status launch evidence preview must mark the handoff consumed")
  assert(evidence_preview["evidence_relative_path"] == evidence_record.fetch("evidence_relative_path"), "Runtime status launch evidence preview must preserve the safe relative handoff path")
  assert(evidence_preview["evidence_sha256"] == evidence_record.fetch("evidence_sha256"), "Runtime status launch evidence preview must preserve the persisted evidence digest")
  assert(evidence_preview["evidence_digest_verified"] == true, "Runtime status launch evidence preview must verify the persisted evidence digest")
  assert(evidence_preview["known_app_smoke_evidence_ready"] == true, "Runtime status launch evidence preview must expose known app smoke evidence")
  evidence_preview_smoke = evidence_preview.fetch("known_app_smoke_evidence")
  assert(evidence_preview_smoke["center_card_state"] == "validated-post-review-dispatch", "Runtime status launch evidence preview must expose the post-review dispatch card state")
  assert(evidence_preview_smoke["primary_action_id"] == "show-runtime-controlled-launch", "Runtime status launch evidence preview must preserve the Runtime status action")
  assert(evidence_preview_smoke["primary_action_kind"] == "runtime-status", "Runtime status launch evidence preview must preserve the Runtime status action kind")
  assert(evidence_preview_smoke["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "Runtime status launch evidence preview must preserve the opaque launch authorization receipt id")
  assert(evidence_preview_smoke["session_gated_review_receipt_id"] == launch_review_receipt.fetch("receipt_id"), "Runtime status launch evidence preview must preserve the opaque session-gated review receipt id")
  assert(evidence_preview_smoke["controlled_execution_session_id"] == controlled_session_record.fetch("execution_session_id"), "Runtime status launch evidence preview must preserve the controlled execution session id")
  assert(evidence_preview["runtime_owned"] == true, "Runtime status launch evidence preview must stay Runtime-owned")
  assert(evidence_preview["runtime_owned_dispatch"] == true, "Runtime status launch evidence preview must preserve Runtime-owned dispatch")
  assert(evidence_preview["go_runtime_backed"] == true, "Runtime status launch evidence preview must be Go Runtime backed")
  assert(evidence_preview["kde_policy_owner"] == false, "Runtime status launch evidence preview must not make KDE the policy owner")
  assert(evidence_preview["state_root_path_exposed"] == false, "Runtime status launch evidence preview must not expose state-root paths")
  assert(evidence_preview["evidence_path_exposed"] == false, "Runtime status launch evidence preview must not expose evidence absolute paths")
  assert(evidence_preview["managed_launcher_path_exposed"] == false, "Runtime status launch evidence preview must not expose launcher paths")
  assert(evidence_preview["raw_launcher_output_exposed"] == false, "Runtime status launch evidence preview must not expose launcher output")
  assert(evidence_preview["backend_details_exposed"] == false, "Runtime status launch evidence preview must not expose backend details")
  assert(evidence_preview["host_root_modified"] == false, "Runtime status launch evidence preview must not mutate host root")
  assert(evidence_preview["docker_socket_mounted"] == false, "Runtime status launch evidence preview must not mount Docker socket")
  assert(evidence_preview["broad_host_mount_required"] == false, "Runtime status launch evidence preview must not require broad host mounts")
  assert(evidence_preview["desktop_launch_enabled"] == false, "Runtime status launch evidence preview must not enable desktop launch")
  assert(evidence_preview["backend_launch_enabled"] == false, "Runtime status launch evidence preview must not enable backend launch")
  assert(evidence_preview["execution_started"] == false, "Runtime status launch evidence preview must not start execution")
  assert(evidence_preview["backend_process_started"] == false, "Runtime status launch evidence preview must not start backend processes")
  assert_no_forbidden(evidence_preview_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime status launch evidence preview output")
  action_trigger, action_trigger_stdout = run_json(
    go_env,
    "go", "run", "./cmd/xnix-runtime-go",
    "known-app-kde-runtime-status-launch-action-trigger-preview",
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--evidence-relative-path", evidence_record.fetch("evidence_relative_path")
  )
  assert(action_trigger["request_type"] == "known-app-kde-runtime-status-launch-action-trigger-preview", "Runtime status launch action trigger must use the Go-owned trigger type")
  assert(action_trigger["runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger", "Runtime status launch action trigger must be owned by Go Runtime")
  assert(action_trigger["read_method"] == "GetKnownAppKDERuntimeStatusLaunchActionTrigger", "Runtime status launch action trigger must expose the KDE-safe trigger read method")
  assert(action_trigger["action_id"] == "show-runtime-controlled-launch", "Runtime status launch action trigger must consume the KDE Runtime status action")
  assert(action_trigger["action_kind"] == "runtime-status", "Runtime status launch action trigger must preserve Runtime status action kind")
  assert(action_trigger["trigger_state"] == "runtime-launch-request-assembled", "Runtime status launch action trigger must assemble the Runtime launch request")
  assert(action_trigger["evidence_read_state"] == "consumed", "Runtime status launch action trigger must consume the handoff evidence")
  assert(action_trigger["evidence_handoff_consumed"] == true, "Runtime status launch action trigger must mark handoff evidence consumed")
  assert(action_trigger["evidence_relative_path"] == evidence_record.fetch("evidence_relative_path"), "Runtime status launch action trigger must preserve the safe relative handoff path")
  assert(action_trigger["evidence_sha256"] == evidence_record.fetch("evidence_sha256"), "Runtime status launch action trigger must preserve evidence digest")
  assert(action_trigger["evidence_digest_verified"] == true, "Runtime status launch action trigger must verify evidence digest")
  assert(action_trigger["launch_request_type"] == "known-app-kde-runtime-status-launch-request-preview", "Runtime status launch action trigger must create a Runtime launch request preview")
  assert(action_trigger["launch_request_runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchRequest", "Runtime status launch action trigger must reuse the Go launch request preview")
  assert(action_trigger["launch_request_created"] == true, "Runtime status launch action trigger must report launch request creation")
  assert(action_trigger["managed_launcher_argv_ready"] == true, "Runtime status launch action trigger must assemble managed launcher argv")
  assert(action_trigger["managed_launcher_argv"] == [
    "xnix-compat-launch",
    "--app", payload.fetch("app_id"),
    "--guest-boundary", "managed-known-app-guest-smoke",
    "--receipt-id", receipt_preview.fetch("receipt_id"),
    "--review-receipt-id", payload.fetch("session_gated_review_receipt_id"),
    "--session-id", payload.fetch("controlled_execution_session_id")
  ], "Runtime status launch action trigger must assemble managed launcher argv from handoff evidence")
  action_trigger_request = action_trigger.fetch("launch_request")
  assert(action_trigger_request["request_type"] == "known-app-kde-runtime-status-launch-request-preview", "Runtime status launch action trigger must embed the launch request preview")
  assert(action_trigger_request["managed_launcher_argv"] == action_trigger.fetch("managed_launcher_argv"), "Runtime status launch action trigger embedded request must match managed launcher argv")
  assert(action_trigger["runtime_owned_trigger"] == true, "Runtime status launch action trigger must be Runtime-owned")
  assert(action_trigger["runtime_owned_request"] == true, "Runtime status launch action trigger must create a Runtime-owned request")
  assert(action_trigger["runtime_owned_launch"] == true, "Runtime status launch action trigger must keep launch Runtime-owned")
  assert(action_trigger["runtime_owned_dispatch"] == true, "Runtime status launch action trigger must keep dispatch Runtime-owned")
  assert(action_trigger["kde_presentation_only"] == true, "Runtime status launch action trigger must keep KDE presentation-only")
  assert(action_trigger["kde_action_forwarded"] == true, "Runtime status launch action trigger must model KDE action forwarding")
  assert(action_trigger["state_root_required"] == true, "Runtime status launch action trigger must keep state-root required")
  assert(action_trigger["state_root_supplied_by_runtime"] == true, "Runtime status launch action trigger must keep state-root supplied by Runtime")
  assert(action_trigger["kde_state_root_access"] == false, "Runtime status launch action trigger must not give KDE state-root access")
  assert(action_trigger["direct_launch_enabled"] == false, "Runtime status launch action trigger must not enable direct launch")
  assert(action_trigger["desktop_launch_enabled"] == false, "Runtime status launch action trigger must not enable desktop launch")
  assert(action_trigger["backend_launch_enabled"] == false, "Runtime status launch action trigger must not enable backend launch")
  assert(action_trigger["execution_started"] == false, "Runtime status launch action trigger must not start execution")
  assert(action_trigger["backend_process_started"] == false, "Runtime status launch action trigger must not start backend processes")
  assert(action_trigger["request_objects_created"] == false, "Runtime status launch action trigger must not write request objects")
  assert(action_trigger["permission_grant_created"] == false, "Runtime status launch action trigger must not create permission grants")
  assert(action_trigger["state_root_path_exposed"] == false, "Runtime status launch action trigger must not expose state-root paths")
  assert(action_trigger["evidence_path_exposed"] == false, "Runtime status launch action trigger must not expose evidence absolute paths")
  assert(action_trigger["receipt_path_exposed"] == false, "Runtime status launch action trigger must not expose receipt paths")
  assert(action_trigger["session_path_exposed"] == false, "Runtime status launch action trigger must not expose session paths")
  assert(action_trigger["raw_artifact_path_exposed"] == false, "Runtime status launch action trigger must not expose raw artifact paths")
  assert(action_trigger["raw_command_exposed"] == false, "Runtime status launch action trigger must not expose raw commands")
  assert(action_trigger["raw_launcher_output_exposed"] == false, "Runtime status launch action trigger must not expose launcher output")
  assert(action_trigger["backend_details_exposed"] == false, "Runtime status launch action trigger must not expose backend details")
  assert(action_trigger["host_root_modified"] == false, "Runtime status launch action trigger must not mutate host root")
  assert(action_trigger["docker_socket_mounted"] == false, "Runtime status launch action trigger must not mount Docker socket")
  assert(action_trigger["broad_host_mount_required"] == false, "Runtime status launch action trigger must not require broad host mounts")
  assert_no_forbidden(action_trigger_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime status launch action trigger output")
  service_call_materialization, service_call_materialization_stdout = run_json(
    go_env,
    "go", "run", "./cmd/xnix-runtime-go",
    "desktop-trigger-service-call-materialization-preview",
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--desktop-entry-file", PROJECT_ROOT.join("kde/actions/xnix-runtime-status-controlled-launch.desktop").to_s,
    "--expected-evidence-sha256", evidence_record.fetch("evidence_sha256"),
    "--human-authorized-smoke",
    "--evidence-relative-path", evidence_record.fetch("evidence_relative_path")
  )
  assert(service_call_materialization["request_type"] == "desktop-trigger-service-call-materialization-preview", "Runtime status launch service call materialization must use the Go-owned materialization type")
  assert(service_call_materialization["runtime_method"] == "PreviewDesktopTriggerServiceCallMaterialization", "Runtime status launch service call materialization must be owned by Go Runtime")
  assert(service_call_materialization["read_method"] == "GetDesktopTriggerServiceCallMaterialization", "Runtime status launch service call materialization must expose a KDE-safe read method")
  assert(service_call_materialization["materialization_state"] == "ready-for-human-authorized-service-call", "Runtime status launch service call materialization must be ready for a human-authorized smoke candidate")
  assert(service_call_materialization["dry_run_review_state"] == "blocked-missing-full-checkpoint", "Runtime status launch service call materialization must not claim full checkpoint promotion during candidate smoke")
  assert(service_call_materialization["owner_trigger_state"] == "ready", "Runtime status launch service call materialization must consume the owner trigger internally")
  assert(service_call_materialization["full_checkpoint_state"] == "needs-full-checkpoint", "Runtime status launch service call materialization must keep the formal full checkpoint pending")
  assert(service_call_materialization["runtime_status_evidence_state"] == "ready", "Runtime status launch service call materialization must consume verified Runtime-status evidence")
  assert(service_call_materialization["human_authorized_smoke"] == true, "Runtime status launch service call materialization must require explicit human smoke authorization")
  assert(service_call_materialization["full_checkpoint_promotion_claimed"] == false, "Runtime status launch service call materialization must not claim checkpoint promotion")
  assert(service_call_materialization["formal_release_ready"] == false, "Runtime status launch service call materialization must not claim release readiness")
  assert(service_call_materialization["desktop_callable_route"] == "kde-dbus-runtime-status-action", "Runtime status launch service call materialization must target the KDE D-Bus Runtime-status route")
  assert(service_call_materialization["desktop_callable_runtime_method"] == "ShowRuntimeControlledLaunch", "Runtime status launch service call materialization must target the owner Runtime method")
  assert(service_call_materialization["desktop_callable_execution_type"] == "known-app-kde-runtime-status-launch-execution", "Runtime status launch service call materialization must target the Runtime launch execution")
  assert(service_call_materialization["desktop_dbus_method"] == "org.xnix.Compatibility1.ShowRuntimeControlledLaunch", "Runtime status launch service call materialization must expose the public D-Bus method")
  assert(service_call_materialization["owner_service_call_ready"] == true, "Runtime status launch service call materialization must expose a ready owner service call")
  assert(service_call_materialization["owner_service_call_type"] == "desktop-action-dispatch", "Runtime status launch service call materialization must classify the owner call as desktop action dispatch")
  assert(service_call_materialization["owner_service_call_args"] == ["ShowRuntimeControlledLaunch", "evidence-relative-path", evidence_record.fetch("evidence_relative_path")], "Runtime status launch service call materialization must expose evidence-only owner service args")
  assert(service_call_materialization["owner_service_cli_args"] == ["--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", evidence_record.fetch("evidence_relative_path")], "Runtime status launch service call materialization must expose evidence-only owner service CLI args")
  assert(service_call_materialization["runtime_owner_service_supplies_inputs"] == true, "Runtime status launch service call materialization must keep owner-only inputs supplied by Runtime")
  assert(service_call_materialization["kde_forwards_only_evidence_handle"] == true, "Runtime status launch service call materialization must keep KDE evidence-only")
  assert(service_call_materialization["kde_receives_materialized_owner_args"] == false, "Runtime status launch service call materialization must not hand owner args to KDE")
  assert(service_call_materialization["desktop_receipt_fields_reconstructed"] == false, "Runtime status launch service call materialization must not reconstruct receipt fields")
  assert(service_call_materialization["desktop_kde_state_root_access"] == false, "Runtime status launch service call materialization must not give KDE state-root access")
  assert(service_call_materialization["state_root_path_exposed"] == false, "Runtime status launch service call materialization must not expose state-root paths")
  assert(service_call_materialization["cache_root_path_exposed"] == false, "Runtime status launch service call materialization must not expose cache-root paths")
  assert(service_call_materialization["launcher_path_exposed"] == false, "Runtime status launch service call materialization must not expose launcher paths")
  assert(service_call_materialization["raw_launcher_output_exposed"] == false, "Runtime status launch service call materialization must not expose launcher output")
  assert(service_call_materialization["backend_details_exposed"] == false, "Runtime status launch service call materialization must not expose backend details")
  assert(service_call_materialization["service_call_dispatch_enabled"] == false, "Runtime status launch service call materialization preview must not dispatch the owner service call")
  assert(service_call_materialization["service_call_dispatched"] == false, "Runtime status launch service call materialization preview must not dispatch the owner service call")
  assert(service_call_materialization["dbus_called"] == false, "Runtime status launch service call materialization preview must not call D-Bus")
  assert(service_call_materialization["desktop_launch_enabled"] == false, "Runtime status launch service call materialization preview must not enable desktop launch")
  assert(service_call_materialization["backend_launch_enabled"] == false, "Runtime status launch service call materialization preview must not enable backend launch")
  assert(service_call_materialization["execution_started"] == false, "Runtime status launch service call materialization preview must not start execution")
  assert(service_call_materialization["runtime_state_written"] == false, "Runtime status launch service call materialization preview must not write Runtime state")
  assert(service_call_materialization["host_root_modified"] == false, "Runtime status launch service call materialization must not mutate host root")
  assert_no_forbidden(service_call_materialization_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime status launch service call materialization output")
  runtime_owner_env = go_env.merge(
    "XNIX_RUNTIME_OWNER_STATE_ROOT" => AUTHORIZATION_STATE_ROOT.to_s,
    "XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT" => KNOWN_APP_CACHE_ROOT.to_s,
    "XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER" => STAGED_LAUNCHER.to_s,
    "XNIX_RUNTIME_OWNER_GUEST_KEY" => Xnix::SshTestKey::PRIVATE_KEY_PATH,
    "XNIX_RUNTIME_OWNER_GUEST_TIMEOUT" => ENV.fetch("XNIX_KNOWN_WINAPP_GUEST_TIMEOUT", "90s"),
    "XNIX_RUNTIME_OWNER_TIMEOUT" => ENV.fetch("XNIX_RUNTIME_OWNER_TIMEOUT", "5m")
  )
  trigger_launcher_stdout, trigger_launcher_stderr, trigger_launcher_status = run_command(
    runtime_owner_env,
    "go", "run", "./cmd/xnix-runtime-owner",
    "--root", ".",
    "--mode", "smoke-owner",
    *service_call_materialization.fetch("owner_service_cli_args")
  )
  unless trigger_launcher_status.zero?
    warn "QEMU serial log: #{SERIAL_LOG_PATH}"
    warn trigger_launcher_stdout unless trigger_launcher_stdout.empty?
    warn trigger_launcher_stderr unless trigger_launcher_stderr.empty?
    warn "FAIL: #{SMOKE_NAME} trigger-fed command failed"
    exit 1
  end
  trigger_owner_service_call = JSON.parse(trigger_launcher_stdout)
  assert(trigger_owner_service_call["request_type"] == "runtime-owner-service-call", "trigger-fed Runtime owner launch must use the owner service call envelope")
  assert(trigger_owner_service_call["service_type"] == "go-runtime-owner-in-process-service", "trigger-fed Runtime owner launch must use the Go owner service boundary")
  assert(trigger_owner_service_call["method"] == "ShowRuntimeControlledLaunch", "trigger-fed Runtime owner launch must call ShowRuntimeControlledLaunch")
  assert(trigger_owner_service_call["call_type"] == "desktop-action-dispatch", "trigger-fed Runtime owner launch must expose desktop action dispatch")
  assert(trigger_owner_service_call["read_only_dispatch"] == false, "trigger-fed Runtime owner launch must not pretend to be a read-only dispatch")
  assert(trigger_owner_service_call["write_method"] == true, "trigger-fed Runtime owner launch must be guarded as a write-like desktop action")
  assert(trigger_owner_service_call["write_methods_enabled"] == false, "trigger-fed Runtime owner launch must keep generic write methods disabled")
  assert(trigger_owner_service_call["dispatch_ready"] == true, "trigger-fed Runtime owner launch must report dispatch readiness after delegated pass")
  assert(trigger_owner_service_call["runtime_owned"] == true, "trigger-fed Runtime owner launch must stay Runtime-owned")
  assert(trigger_owner_service_call["go_runtime_backed"] == true, "trigger-fed Runtime owner launch must stay Go Runtime backed")
  assert(trigger_owner_service_call["kde_policy_owner"] == false, "trigger-fed Runtime owner launch must not make KDE the policy owner")
  assert(trigger_owner_service_call["production_bus_claimed"] == false, "trigger-fed Runtime owner launch must not claim the production bus")
  assert(trigger_owner_service_call["network_required"] == false, "trigger-fed Runtime owner launch service envelope must not require host networking")
  assert(trigger_owner_service_call["host_root_modified"] == false, "trigger-fed Runtime owner launch service envelope must not mutate host root")
  assert(trigger_owner_service_call["backend_details_exposed"] == false, "trigger-fed Runtime owner launch service envelope must not expose backend details")
  trigger_runtime_execution = trigger_owner_service_call.fetch("payload")
  assert(trigger_runtime_execution["owner_request_type"] == "runtime-owner-show-runtime-controlled-launch", "trigger-fed Runtime owner launch must return the owner action result")
  assert(trigger_runtime_execution["owner_runtime_method"] == "ShowRuntimeControlledLaunch", "trigger-fed Runtime owner launch must preserve the owner Runtime method")
  assert(trigger_runtime_execution["owner_service_boundary"] == "go-runtime-owner-in-process-service", "trigger-fed Runtime owner launch action must stay inside the owner service boundary")
  assert(trigger_runtime_execution["runtime_owner_service_action_dispatch"] == true, "trigger-fed Runtime owner launch must report service action dispatch")
  assert(trigger_runtime_execution["runtime_owner_service_call_ready"] == true, "trigger-fed Runtime owner launch must report service-call readiness")
  assert(trigger_runtime_execution["runtime_owner_service_supplies_owner_inputs"] == true, "trigger-fed Runtime owner launch must prove owner-only inputs came from the service")
  assert(trigger_runtime_execution["kde_forwards_only_evidence_handle"] == true, "trigger-fed Runtime owner launch must prove KDE forwards only evidence handles")
  assert(trigger_runtime_execution["desktop_callable_action_id"] == "show-runtime-controlled-launch", "trigger-fed Runtime status launch execution must be callable through the KDE desktop action")
  assert(trigger_runtime_execution["desktop_callable_route"] == "kde-dbus-runtime-status-action", "trigger-fed Runtime status launch execution must expose the D-Bus-callable owner route")
  assert(trigger_runtime_execution["desktop_callable_runtime_method"] == "ShowRuntimeControlledLaunch", "trigger-fed Runtime status launch execution must route through the owner service method")
  assert(trigger_runtime_execution["desktop_callable_execution_type"] == "known-app-kde-runtime-status-launch-execution", "trigger-fed Runtime status launch execution must preserve the Runtime execution type")
  assert(trigger_runtime_execution["desktop_evidence_handle_forwarded"] == true, "trigger-fed Runtime status launch execution must forward only the evidence handoff")
  assert(trigger_runtime_execution["desktop_receipt_fields_reconstructed"] == false, "trigger-fed Runtime status launch execution must not reconstruct receipt or session fields in KDE")
  assert(trigger_runtime_execution["desktop_kde_state_root_access"] == false, "trigger-fed Runtime status launch execution must not grant KDE state-root access")
  assert(trigger_runtime_execution["desktop_runtime_owner_adapter_used"] == true, "trigger-fed Runtime status launch execution must use the Runtime-owner adapter")
  assert(trigger_runtime_execution["desktop_state_root_supplied_by_runtime_owner"] == true, "trigger-fed Runtime status launch execution must receive state-root from the Runtime owner")
  assert(trigger_runtime_execution["desktop_cache_root_supplied_by_runtime_owner"] == true, "trigger-fed Runtime status launch execution must receive cache root from the Runtime owner")
  assert(trigger_runtime_execution["desktop_launcher_supplied_by_runtime_owner"] == true, "trigger-fed Runtime status launch execution must receive launcher path from the Runtime owner")
  assert(trigger_runtime_execution["desktop_timeout_supplied_by_runtime_owner"] == true, "trigger-fed Runtime status launch execution must receive timeout settings from the Runtime owner")
  assert(trigger_runtime_execution["request_type"] == "known-app-kde-runtime-status-launch-execution", "trigger-fed Runtime status launch execution must use the Runtime-owned entrypoint")
  assert(trigger_runtime_execution["action_trigger_type"] == "known-app-kde-runtime-status-launch-action-trigger-preview", "trigger-fed Runtime status launch execution must consume the action trigger")
  assert(trigger_runtime_execution["action_trigger_runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger", "trigger-fed Runtime status launch execution must consume the Go Runtime action trigger")
  assert(trigger_runtime_execution["action_trigger_read_method"] == "GetKnownAppKDERuntimeStatusLaunchActionTrigger", "trigger-fed Runtime status launch execution must expose the action trigger read method")
  assert(trigger_runtime_execution["action_trigger_state"] == "runtime-launch-request-assembled", "trigger-fed Runtime status launch execution must receive an assembled launch request")
  assert(trigger_runtime_execution["evidence_handoff_consumed"] == true, "trigger-fed Runtime status launch execution must consume handoff evidence")
  assert(trigger_runtime_execution["evidence_digest_verified"] == true, "trigger-fed Runtime status launch execution must verify handoff digest")
  assert(trigger_runtime_execution["evidence_relative_path"] == evidence_record.fetch("evidence_relative_path"), "trigger-fed Runtime status launch execution must preserve relative handoff evidence path")
  assert(trigger_runtime_execution["evidence_sha256"] == evidence_record.fetch("evidence_sha256"), "trigger-fed Runtime status launch execution must preserve handoff digest")
  assert(trigger_runtime_execution["state_root_injected_by_runtime"] == true, "trigger-fed Runtime status launch execution must inject state-root internally")
  assert(trigger_runtime_execution["state_root_supplied_by_runtime"] == true, "trigger-fed Runtime status launch execution must keep state-root supplied by Runtime")
  assert(trigger_runtime_execution["kde_state_root_access"] == false, "trigger-fed Runtime status launch execution must not give KDE state-root access")
  assert(trigger_runtime_execution["managed_launcher_invoked"] == true, "trigger-fed Runtime status launch execution must invoke the managed launcher")
  assert(trigger_runtime_execution["existing_managed_launcher_invoked"] == true, "trigger-fed Runtime status launch execution must invoke the existing staged launcher")
  assert(trigger_runtime_execution["launcher_output_json_observed"] == true, "trigger-fed Runtime status launch execution must observe delegated launcher JSON")
  assert(trigger_runtime_execution["delegated_status"] == "passed", "trigger-fed Runtime status launch execution must pass delegated launcher evidence")
  assert(trigger_runtime_execution["delegated_session_gated_review_receipt_id"] == payload.fetch("session_gated_review_receipt_id"), "trigger-fed Runtime status launch execution must preserve delegated review receipt id")
  assert(trigger_runtime_execution["delegated_controlled_execution_session_id"] == payload.fetch("controlled_execution_session_id"), "trigger-fed Runtime status launch execution must preserve delegated controlled session id")
  assert(trigger_runtime_execution["raw_launcher_output_exposed"] == false, "trigger-fed Runtime status launch execution must not expose raw launcher output")
  assert(trigger_runtime_execution["state_root_path_exposed"] == false, "trigger-fed Runtime status launch execution must not expose state-root paths")
  assert(trigger_runtime_execution["managed_launcher_path_exposed"] == false, "trigger-fed Runtime status launch execution must not expose launcher paths")
  assert(trigger_runtime_execution["backend_details_exposed"] == false, "trigger-fed Runtime status launch execution must not expose backend details")
  assert(trigger_runtime_execution["host_root_modified"] == false, "trigger-fed Runtime status launch execution must not mutate host root")
  assert_no_forbidden(trigger_launcher_stdout, [PROJECT_ROOT.to_s, AUTHORIZATION_STATE_ROOT.to_s, STAGED_LAUNCHER.to_s, RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "trigger-fed Runtime status launch execution output")

  case payload.fetch("status")
  when "passed"
    center_preview_args = [
      "go", "run", "./cmd/xnix-runtime-go",
      "compatibility-center-preview",
      "--registry", "runtime/recipes/registry.json",
      "--known-app-evidence-file", RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s
    ]
    center_preview, center_preview_stdout = run_json(go_env, *center_preview_args)
    assert(center_preview["known_app_smoke_evidence_count"] == 1, "Compatibility Center must receive known app smoke evidence")
    assert(center_preview["known_app_smoke_passed_count"] == 1, "Compatibility Center must count passed known app smoke evidence")
    assert(center_preview["known_app_staged_launcher_passed_count"] == 1, "Compatibility Center must count staged launcher smoke evidence")
    assert(center_preview["known_app_launch_authorization_required_count"] == 1, "Compatibility Center must count launch authorization requirements")
    assert(center_preview["known_app_launch_authorization_recorded_count"] == 1, "Compatibility Center must count recorded launch authorization receipts")
    assert(center_preview["known_app_launch_gate_consumed_count"] == 1, "Compatibility Center must count launch gate consumption")
    assert(center_preview["known_app_controlled_dispatch_ready_count"] == 1, "Compatibility Center must count controlled dispatch readiness")
    assert(center_preview["known_app_launcher_session_gate_consumed_count"] == 1, "Compatibility Center must count launcher-side session gate consumption")
    assert(center_preview["known_app_post_review_dispatch_consumed_count"] == 1, "Compatibility Center must count post-review dispatch consumption")
    center_evidence = center_preview.fetch("known_app_smoke_evidence").first
    assert(center_evidence["evidence_source"] == "staged-launcher-dispatch-smoke", "Compatibility Center evidence must identify the staged launcher source")
    assert(center_evidence["center_card_state"] == "validated-post-review-dispatch", "Compatibility Center evidence must expose post-review dispatch card state")
    assert(center_evidence["launch_authorization_state"] == "recorded", "Compatibility Center evidence must expose recorded launch authorization state")
    assert(center_evidence["primary_action_id"] == "show-runtime-controlled-launch", "Compatibility Center evidence must expose Runtime-controlled launch status as the primary action")
    assert(center_evidence["primary_action_kind"] == "runtime-status", "Compatibility Center evidence must expose a Runtime status action")
    assert(center_evidence["primary_action_enabled"] == true, "Compatibility Center evidence must allow the safe Runtime status action")
    assert(center_evidence["direct_launch_enabled"] == false, "Compatibility Center evidence must not enable direct launch")
    assert(center_evidence["launch_authorization_receipt_state"] == "recorded", "Compatibility Center evidence must expose recorded receipt state")
    assert(center_evidence["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "Compatibility Center evidence must expose the opaque receipt id")
    assert(center_evidence["launch_gate_state"] == "controlled-dispatch-ready", "Compatibility Center evidence must expose controlled dispatch readiness")
    assert(center_evidence["launch_gate_consumed"] == true, "Compatibility Center evidence must expose launch gate consumption")
    assert(center_evidence["launch_gate_receipt_accepted"] == true, "Compatibility Center evidence must expose receipt acceptance")
    assert(center_evidence["launch_gate_guest_boundary_accepted"] == true, "Compatibility Center evidence must expose guest boundary acceptance")
    assert(center_evidence["controlled_dispatch_ready"] == true, "Compatibility Center evidence must expose controlled dispatch readiness")
    assert(center_evidence["controlled_execution_session_id"] == payload.fetch("controlled_execution_session_id"), "Compatibility Center evidence must expose the opaque controlled execution session id")
    assert(center_evidence["launcher_session_gate_consumed"] == true, "Compatibility Center evidence must expose launcher-side session gate consumption")
    assert(center_evidence["launcher_session_digest_verified"] == true, "Compatibility Center evidence must expose launcher-side session digest verification")
    assert(center_evidence["launcher_session_relative_path"] == payload.fetch("controlled_session_relative_path"), "Compatibility Center evidence must expose relative launcher session evidence")
    assert(center_evidence["launcher_session_runtime_owner_consumable"] == true, "Compatibility Center evidence must expose Runtime-owner session consumption readiness")
    assert(center_evidence["launcher_session_kde_read_model_consumable"] == true, "Compatibility Center evidence must expose KDE read-model session consumption readiness")
    assert(center_evidence["post_review_dispatch_consumed"] == true, "Compatibility Center evidence must expose post-review dispatch consumption")
    assert(center_evidence["post_review_dispatch_state"] == "created-after-session-gated-review", "Compatibility Center evidence must expose post-review dispatch state")
    assert(center_evidence["session_gated_review_receipt_id"] == payload.fetch("session_gated_review_receipt_id"), "Compatibility Center evidence must expose the opaque session-gated review receipt id")
    assert(center_evidence["staged_launcher_verified"] == true, "Compatibility Center evidence must verify the staged launcher path")
    assert(center_evidence["runtime_dispatch_verified"] == true, "Compatibility Center evidence must verify Runtime dispatch")
    assert(center_evidence["launch_authorization_required"] == true, "Compatibility Center evidence must keep launch authorization required")
    assert(center_evidence["desktop_launch_enabled"] == false, "Compatibility Center evidence must not enable desktop launch")
    assert(center_evidence["backend_launch_enabled"] == false, "Compatibility Center evidence must not enable backend launch")
    assert(center_evidence["host_root_modified"] == false, "Compatibility Center evidence must not mutate the host root")
    assert_no_forbidden(center_preview_stdout, [PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Compatibility Center preview output")
    kde_page_args = [
      "go", "run", "./cmd/xnix-runtime-go",
      "kde-center-page-preview",
      "--registry", "runtime/recipes/registry.json",
      "--app", "org.xnix.sample.notepad",
      "--decision", "approved",
      "--known-app-evidence-file", RUNTIME_STATUS_LAUNCH_EVIDENCE_PATH.to_s
    ]
    kde_page, kde_page_stdout = run_json(go_env, *kde_page_args)
    assert(kde_page["known_app_session_gate_evidence_count"] == 1, "KDE Center page must receive known app session gate evidence")
    assert(kde_page["known_app_launcher_session_gate_consumed_count"] == 1, "KDE Center page must count launcher-side session gate consumption")
    assert(kde_page["known_app_post_review_dispatch_consumed_count"] == 1, "KDE Center page must count post-review dispatch consumption")
    kde_page_card = kde_page.fetch("known_app_session_gate_cards").first
    assert(kde_page_card["center_card_state"] == "validated-post-review-dispatch", "KDE Center page card must expose post-review dispatch state")
    assert(kde_page_card["controlled_execution_session_id"] == payload.fetch("controlled_execution_session_id"), "KDE Center page card must expose the opaque controlled execution session id")
    assert(kde_page_card["launcher_session_gate_consumed"] == true, "KDE Center page card must expose launcher session gate consumption")
    assert(kde_page_card["launcher_session_digest_verified"] == true, "KDE Center page card must expose launcher session digest verification")
    assert(kde_page_card["launcher_session_relative_path"] == payload.fetch("controlled_session_relative_path"), "KDE Center page card must expose relative launcher session evidence")
    assert(kde_page_card["runtime_owner_consumable_session"] == true, "KDE Center page card must expose Runtime-owner session consumption readiness")
    assert(kde_page_card["kde_read_model_consumable_session"] == true, "KDE Center page card must expose KDE read-model session consumption readiness")
    assert(kde_page_card["post_review_dispatch_consumed"] == true, "KDE Center page card must expose post-review dispatch consumption")
    assert(kde_page_card["post_review_dispatch_state"] == "created-after-session-gated-review", "KDE Center page card must expose post-review dispatch state")
    assert(kde_page_card["session_gated_review_receipt_id"] == payload.fetch("session_gated_review_receipt_id"), "KDE Center page card must expose the opaque session-gated review receipt id")
    assert(kde_page_card["primary_action_id"] == "show-runtime-controlled-launch", "KDE Center page card must expose Runtime-controlled launch status as the primary action")
    assert(kde_page_card["primary_action_kind"] == "runtime-status", "KDE Center page card must expose a Runtime status action")
    assert(kde_page_card["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "KDE Center page card must expose the opaque launch authorization receipt id")
    assert(kde_page_card["runtime_status_launch_request_type"] == "known-app-kde-runtime-status-launch-request-preview", "KDE Center page card must expose the Runtime-status launch request type")
    assert(kde_page_card["runtime_status_launch_runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchRequest", "KDE Center page card must expose the Runtime-status launch runtime method")
    assert(kde_page_card["runtime_status_launch_read_method"] == "GetKnownAppKDERuntimeStatusLaunchRequest", "KDE Center page card must expose the Runtime-status launch read method")
    assert(kde_page_card["runtime_status_launch_required_id_count"] == 3, "KDE Center page card must require three opaque ids for Runtime launch request assembly")
    assert(kde_page_card["runtime_status_launch_collected_id_count"] == 3, "KDE Center page card must collect three opaque ids for Runtime launch request assembly")
    assert(kde_page_card["runtime_status_launch_request_ready"] == true, "KDE Center page card must mark the Runtime-status launch request ready")
    assert(kde_page_card["runtime_status_launch_state_root_required"] == true, "KDE Center page card must keep state-root required for Runtime execution")
    assert(kde_page_card["runtime_status_launch_state_root_owned_by_runtime"] == true, "KDE Center page card must keep state-root owned by Runtime")
    assert(kde_page_card["runtime_status_launch_managed_launcher_argv"] == [
      "xnix-compat-launch",
      "--app", payload.fetch("app_id"),
      "--guest-boundary", "managed-known-app-guest-smoke",
      "--receipt-id", receipt_preview.fetch("receipt_id"),
      "--review-receipt-id", payload.fetch("session_gated_review_receipt_id"),
      "--session-id", payload.fetch("controlled_execution_session_id")
    ], "KDE Center page card must collect managed launcher argv without state-root")
    assert(kde_page_card["desktop_launch_enabled"] == false, "KDE Center page card must not enable desktop launch")
    assert(kde_page_card["backend_launch_enabled"] == false, "KDE Center page card must not enable backend launch")
    assert(kde_page_card["host_root_modified"] == false, "KDE Center page card must not mutate the host root")
    assert(kde_page["launch_enabled"] == false, "KDE Center page must remain launch-gated")
    assert(kde_page["execution_started"] == false, "KDE Center page must not start execution")
    assert_no_forbidden(kde_page_stdout, [PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "KDE Center page preview output")
    runtime_status_launch, runtime_status_launch_stdout = run_json(go_env,
      "go", "run", "./cmd/xnix-runtime-go",
      "known-app-kde-runtime-status-launch-request-preview",
      "--app", payload.fetch("app_id"),
      "--launch-authorization-receipt-id", kde_page_card.fetch("launch_authorization_receipt_id"),
      "--session-gated-review-receipt-id", kde_page_card.fetch("session_gated_review_receipt_id"),
      "--session-id", kde_page_card.fetch("controlled_execution_session_id"),
      "--center-card-state", kde_page_card.fetch("center_card_state"),
      "--primary-action-id", kde_page_card.fetch("primary_action_id"),
      "--post-review-dispatch-state", kde_page_card.fetch("post_review_dispatch_state")
    )
    assert(runtime_status_launch["request_type"] == "known-app-kde-runtime-status-launch-request-preview", "Runtime-status launch request preview must use the Go Runtime request type")
    assert(runtime_status_launch["runtime_method"] == "PreviewKnownAppKDERuntimeStatusLaunchRequest", "Runtime-status launch request preview must be owned by the Go Runtime")
    assert(runtime_status_launch["managed_launcher_argv_ready"] == true, "Runtime-status launch request preview must assemble managed launcher argv")
    assert(runtime_status_launch["required_opaque_id_count"] == 3, "Runtime-status launch request preview must require three opaque ids")
    assert(runtime_status_launch["collected_opaque_id_count"] == 3, "Runtime-status launch request preview must collect three opaque ids")
    assert(runtime_status_launch["managed_launcher_argv"] == kde_page_card.fetch("runtime_status_launch_managed_launcher_argv"), "Runtime-status launch request preview must match the KDE-collected managed launcher argv")
    assert(runtime_status_launch["state_root_required"] == true, "Runtime-status launch request preview must keep state-root required")
    assert(runtime_status_launch["state_root_supplied_by_runtime"] == true, "Runtime-status launch request preview must keep state-root supplied by Runtime")
    assert(runtime_status_launch["kde_state_root_access"] == false, "Runtime-status launch request preview must not give KDE state-root access")
    assert(runtime_status_launch["direct_launch_enabled"] == false, "Runtime-status launch request preview must not enable direct launch")
    assert(runtime_status_launch["desktop_launch_enabled"] == false, "Runtime-status launch request preview must not enable desktop launch")
    assert(runtime_status_launch["backend_launch_enabled"] == false, "Runtime-status launch request preview must not enable backend launch")
    assert(runtime_status_launch["execution_started"] == false, "Runtime-status launch request preview must not start execution")
    assert(runtime_status_launch["request_objects_created"] == false, "Runtime-status launch request preview must not write request objects")
    assert(runtime_status_launch["permission_grant_created"] == false, "Runtime-status launch request preview must not create permission grants")
    assert(runtime_status_launch["host_root_modified"] == false, "Runtime-status launch request preview must not mutate the host root")
    assert_no_forbidden(runtime_status_launch_stdout, [PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Runtime-status launch request preview output")
    puts "PASS: #{SMOKE_NAME} (#{payload.fetch("app_id")} #{payload.fetch("app_version")})"
    exit 0
  when "skipped"
    puts "SKIP: #{SMOKE_NAME} (#{payload.fetch("skip_reason")})"
    exit 0
  else
    warn launcher_stdout
    warn launcher_stderr unless launcher_stderr.empty?
    warn "FAIL: #{SMOKE_NAME}"
    exit 1
  end
ensure
  stop_qemu(wait_thread)
  output.close unless output.closed?
  reader.join
  SERIAL_LOG_PATH.write(serial_log) unless serial_log.empty?
end
