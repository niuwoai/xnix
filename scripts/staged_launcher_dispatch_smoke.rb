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
    {},
    STAGED_LAUNCHER.to_s,
    "--app", APP_ID,
    "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
    "--guest-boundary", GUEST_BOUNDARY,
    "--state-root", AUTHORIZATION_STATE_ROOT.to_s,
    "--receipt-id", receipt_preview.fetch("receipt_id"),
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

  payload = JSON.parse(launcher_stdout)
  assert(payload["request_type"] == "windows-known-app-dispatch-smoke", "staged launcher must enter dispatch smoke with the guest boundary")
  assert(payload["guest_boundary"] == GUEST_BOUNDARY, "staged launcher dispatch must preserve the guest boundary")
  assert(payload["runtime_owned_dispatch"] == true, "staged launcher dispatch must be Runtime-owned")
  assert(payload["host_root_modified"] == false, "staged launcher dispatch must not mutate the host root")
  assert(payload["docker_socket_mounted"] == false, "staged launcher dispatch must not mount the Docker socket")
  assert(payload["broad_host_mount_required"] == false, "staged launcher dispatch must not require broad host mounts")
  assert_no_forbidden(launcher_stdout, [PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "staged launcher dispatch output")

  case payload.fetch("status")
  when "passed"
    center_preview_args = [
      "go", "run", "./cmd/xnix-runtime-go",
      "compatibility-center-preview",
      "--registry", "runtime/recipes/registry.json",
      "--known-app-smoke-app", payload.fetch("app_id"),
      "--known-app-smoke-name", payload.fetch("display_name"),
      "--known-app-smoke-version", payload.fetch("app_version"),
      "--known-app-smoke-source", "staged-launcher-dispatch-smoke",
      "--known-app-smoke-status", "passed",
      "--known-app-launch-authorization-receipt-state", "recorded",
      "--known-app-launch-authorization-receipt-id", receipt_preview.fetch("receipt_id"),
      "--known-app-launch-gate-state", launch_gate.fetch("launch_gate_state")
    ]
    center_preview_args << "--known-app-smoke-marker-observed" if payload["marker_observed"]
    center_preview_args << "--known-app-smoke-checksum-verified" if payload["artifact_verified"]
    center_preview_args << "--known-app-launch-gate-consumed" if controlled_dispatch["receipt_accepted"]
    center_preview_args << "--known-app-launch-gate-receipt-accepted" if controlled_dispatch["receipt_accepted"]
    center_preview_args << "--known-app-launch-gate-guest-boundary-accepted" if controlled_dispatch["guest_boundary_accepted"]
    center_preview_args << "--known-app-controlled-dispatch-ready" if controlled_dispatch["controlled_dispatch_ready"]
    center_preview_args.concat(["--known-app-launch-gate-blocked-reason", launch_gate["launch_gate_blocked_reason"]]) if launch_gate["launch_gate_blocked_reason"]
    center_preview, center_preview_stdout = run_json(go_env, *center_preview_args)
    assert(center_preview["known_app_smoke_evidence_count"] == 1, "Compatibility Center must receive known app smoke evidence")
    assert(center_preview["known_app_smoke_passed_count"] == 1, "Compatibility Center must count passed known app smoke evidence")
    assert(center_preview["known_app_staged_launcher_passed_count"] == 1, "Compatibility Center must count staged launcher smoke evidence")
    assert(center_preview["known_app_launch_authorization_required_count"] == 1, "Compatibility Center must count launch authorization requirements")
    assert(center_preview["known_app_launch_authorization_recorded_count"] == 1, "Compatibility Center must count recorded launch authorization receipts")
    assert(center_preview["known_app_launch_gate_consumed_count"] == 1, "Compatibility Center must count launch gate consumption")
    assert(center_preview["known_app_controlled_dispatch_ready_count"] == 1, "Compatibility Center must count controlled dispatch readiness")
    center_evidence = center_preview.fetch("known_app_smoke_evidence").first
    assert(center_evidence["evidence_source"] == "staged-launcher-dispatch-smoke", "Compatibility Center evidence must identify the staged launcher source")
    assert(center_evidence["center_card_state"] == "validated-launch-gate-consumed", "Compatibility Center evidence must expose launch-gate-consumed card state")
    assert(center_evidence["launch_authorization_state"] == "recorded", "Compatibility Center evidence must expose recorded launch authorization state")
    assert(center_evidence["primary_action_id"] == "review-controlled-dispatch", "Compatibility Center evidence must expose controlled dispatch review as the primary action")
    assert(center_evidence["primary_action_kind"] == "launch-gate-review", "Compatibility Center evidence must expose a launch gate review action")
    assert(center_evidence["primary_action_enabled"] == true, "Compatibility Center evidence must allow the safe launch gate review action")
    assert(center_evidence["direct_launch_enabled"] == false, "Compatibility Center evidence must not enable direct launch")
    assert(center_evidence["launch_authorization_receipt_state"] == "recorded", "Compatibility Center evidence must expose recorded receipt state")
    assert(center_evidence["launch_authorization_receipt_id"] == receipt_preview.fetch("receipt_id"), "Compatibility Center evidence must expose the opaque receipt id")
    assert(center_evidence["launch_gate_state"] == "controlled-dispatch-ready", "Compatibility Center evidence must expose controlled dispatch readiness")
    assert(center_evidence["launch_gate_consumed"] == true, "Compatibility Center evidence must expose launch gate consumption")
    assert(center_evidence["launch_gate_receipt_accepted"] == true, "Compatibility Center evidence must expose receipt acceptance")
    assert(center_evidence["launch_gate_guest_boundary_accepted"] == true, "Compatibility Center evidence must expose guest boundary acceptance")
    assert(center_evidence["controlled_dispatch_ready"] == true, "Compatibility Center evidence must expose controlled dispatch readiness")
    assert(center_evidence["staged_launcher_verified"] == true, "Compatibility Center evidence must verify the staged launcher path")
    assert(center_evidence["runtime_dispatch_verified"] == true, "Compatibility Center evidence must verify Runtime dispatch")
    assert(center_evidence["launch_authorization_required"] == true, "Compatibility Center evidence must keep launch authorization required")
    assert(center_evidence["desktop_launch_enabled"] == false, "Compatibility Center evidence must not enable desktop launch")
    assert(center_evidence["backend_launch_enabled"] == false, "Compatibility Center evidence must not enable backend launch")
    assert(center_evidence["host_root_modified"] == false, "Compatibility Center evidence must not mutate the host root")
    assert_no_forbidden(center_preview_stdout, [PROJECT_ROOT.to_s, "wine ", "wine/", ".wine", "qemu-system", "program files"], "Compatibility Center preview output")
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
