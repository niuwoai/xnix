#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require "shellwords"
require_relative "../lib/xnix/qemu"
require_relative "../lib/xnix/ssh_test_key"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapp-guest-wine-smoke")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
GO_TMP_ROOT = GO_CACHE_ROOT.join("tmp")
KNOWN_APP_CACHE_ROOT = PROJECT_ROOT.join(".cache", "xnix", "known-winapps")
SERIAL_LOG_PATH = WORK_ROOT.join("qemu-serial.log")
APP_ID = ENV.fetch("XNIX_KNOWN_WINAPP_ID", "7zr")
APP_ARGS = Shellwords.split(ENV.fetch("XNIX_KNOWN_WINAPP_ARGS", ""))
def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

unless File.file?(Xnix::SshTestKey::PRIVATE_KEY_PATH)
  puts "SKIP: known Windows app QEMU guest Wine smoke (SSH test key unavailable; run prepare-ssh-test-key and start-build-ssh-wine-guest first)"
  exit 0
end

qemu = Xnix::Qemu.wine_guest

unless File.file?(qemu.kernel_image)
  puts "SKIP: known Windows app QEMU guest Wine smoke (Wine guest kernel unavailable; run configure-wine-guest and start-build-ssh-wine-guest first)"
  exit 0
end

FileUtils.mkdir_p(WORK_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))
FileUtils.mkdir_p(GO_TMP_ROOT)
FileUtils.rm_f(SERIAL_LOG_PATH)

smoke_stdout, smoke_stderr, smoke_status = run_command(
  {
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s,
    "GOTMPDIR" => GO_TMP_ROOT.to_s
  },
  "go", "run", "./cmd/xnix-runtime-go", "windows-known-app-run",
  "--backend", "guest-wine",
  "--start-qemu",
  "--app", APP_ID,
  "--cache-root", KNOWN_APP_CACHE_ROOT.to_s,
  "--key", Xnix::SshTestKey::PRIVATE_KEY_PATH,
  "--qemu-binary", qemu.binary,
  "--qemu-kernel", qemu.kernel_image,
  "--qemu-memory", qemu.memory,
  "--qemu-cpus", qemu.cpu_count,
  "--qemu-cpu", qemu.cpu_model,
  "--qemu-serial-log", SERIAL_LOG_PATH.to_s,
  "--qemu-boot-timeout", ENV.fetch("XNIX_KNOWN_WINAPP_QEMU_BOOT_TIMEOUT", "180s"),
  "--timeout", ENV.fetch("XNIX_KNOWN_WINAPP_GUEST_TIMEOUT", "90s"),
  *APP_ARGS.flat_map { |argument| ["--arg", argument] }
)

unless smoke_status.zero?
  warn "QEMU serial log: #{SERIAL_LOG_PATH}"
  warn smoke_stdout unless smoke_stdout.empty?
  warn smoke_stderr unless smoke_stderr.empty?
  warn "FAIL: known Windows app QEMU guest Wine smoke command failed"
  exit 1
end

payload = JSON.parse(smoke_stdout)
case payload.fetch("status")
when "passed"
  abort "known Windows app QEMU guest Wine smoke used unexpected backend" unless payload.fetch("backend") == "guest-wine"
  abort "known Windows app QEMU guest Wine smoke did not use Go QEMU start mode" unless payload.fetch("guest_start_mode") == "go-qemu"
  abort "known Windows app QEMU guest Wine smoke did not start QEMU from Go" unless payload.fetch("guest_started") == true

  puts "PASS: known Windows app QEMU guest Wine smoke (#{payload.fetch("app_id")} #{payload.fetch("app_version")})"
  exit 0
when "skipped"
  puts "SKIP: known Windows app QEMU guest Wine smoke (#{payload.fetch("skip_reason")})"
  exit 0
else
  warn smoke_stdout
  warn smoke_stderr unless smoke_stderr.empty?
  warn "FAIL: known Windows app QEMU guest Wine smoke"
  exit 1
end
