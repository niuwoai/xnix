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
      "--known-app-smoke-status", "passed"
    ]
    center_preview_args << "--known-app-smoke-marker-observed" if payload["marker_observed"]
    center_preview_args << "--known-app-smoke-checksum-verified" if payload["artifact_verified"]
    center_preview, center_preview_stdout = run_json(go_env, *center_preview_args)
    assert(center_preview["known_app_smoke_evidence_count"] == 1, "Compatibility Center must receive known app smoke evidence")
    assert(center_preview["known_app_smoke_passed_count"] == 1, "Compatibility Center must count passed known app smoke evidence")
    assert(center_preview["known_app_staged_launcher_passed_count"] == 1, "Compatibility Center must count staged launcher smoke evidence")
    assert(center_preview["known_app_launch_authorization_required_count"] == 1, "Compatibility Center must count launch authorization requirements")
    center_evidence = center_preview.fetch("known_app_smoke_evidence").first
    assert(center_evidence["evidence_source"] == "staged-launcher-dispatch-smoke", "Compatibility Center evidence must identify the staged launcher source")
    assert(center_evidence["center_card_state"] == "validated-launch-authorization-required", "Compatibility Center evidence must expose validated authorization-required card state")
    assert(center_evidence["launch_authorization_state"] == "review-required", "Compatibility Center evidence must expose launch authorization review state")
    assert(center_evidence["primary_action_id"] == "review-launch-authorization", "Compatibility Center evidence must expose launch authorization review as the primary action")
    assert(center_evidence["primary_action_kind"] == "authorization-review", "Compatibility Center evidence must expose an authorization review action")
    assert(center_evidence["primary_action_enabled"] == true, "Compatibility Center evidence must allow the safe authorization review action")
    assert(center_evidence["direct_launch_enabled"] == false, "Compatibility Center evidence must not enable direct launch")
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
