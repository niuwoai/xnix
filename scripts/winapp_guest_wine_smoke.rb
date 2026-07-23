#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/qemu"
require_relative "../lib/xnix/ssh_probe"
require_relative "../lib/xnix/ssh_test_key"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
WORK_ROOT = PROJECT_ROOT.join(".cache", "xnix", "winapp-guest-wine-smoke")
GO_CACHE_ROOT = PROJECT_ROOT.join(".cache", "go")
APP_ROOT = WORK_ROOT.join("app")
EXE_PATH = APP_ROOT.join("hello.exe")
MARKER = "XNIX_WINAPP_SMOKE_OK"
BOOT_TIMEOUT_SECONDS = 60
RETRY_INTERVAL_SECONDS = 1
FIXTURE_GOARCH = "386"

def run_command(env, *argv)
  stdout, stderr, status = Open3.capture3(env, *argv, chdir: PROJECT_ROOT.to_s)
  [stdout, stderr, status.exitstatus]
end

def stop_qemu(wait_thread)
  return unless wait_thread.alive?

  Process.kill("TERM", wait_thread.pid)
  wait_thread.join(5)
  Process.kill("KILL", wait_thread.pid) if wait_thread.alive?
rescue Errno::ESRCH
  nil
end

unless File.file?(Xnix::SshTestKey::PRIVATE_KEY_PATH)
  puts "SKIP: QEMU guest Wine smoke (SSH test key unavailable; run prepare-ssh-test-key and start-build-ssh-wine-guest first)"
  exit 0
end

qemu = Xnix::Qemu.wine_guest

unless File.file?(qemu.kernel_image)
  puts "SKIP: QEMU guest Wine smoke (Wine guest kernel unavailable; run configure-wine-guest and start-build-ssh-wine-guest first)"
  exit 0
end

FileUtils.mkdir_p(APP_ROOT)
FileUtils.mkdir_p(GO_CACHE_ROOT.join("build"))
FileUtils.mkdir_p(GO_CACHE_ROOT.join("mod"))

build_stdout, build_stderr, build_status = run_command(
  {
    "GOOS" => "windows",
    "GOARCH" => FIXTURE_GOARCH,
    "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
    "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
  },
  "go", "build", "-o", EXE_PATH.to_s, "./test/fixtures/winapp/hello"
)

unless build_status.zero?
  warn build_stdout unless build_stdout.empty?
  warn build_stderr unless build_stderr.empty?
  warn "FAIL: Windows app fixture build failed"
  exit 1
end

stdin, output, wait_thread = Open3.popen2e(*qemu.boot_command(ssh: true))
stdin.close
reader = Thread.new { output.read }
probe = Xnix::SshProbe.new
deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + BOOT_TIMEOUT_SECONDS

begin
  loop do
    break if system(*probe.command, out: File::NULL, err: File::NULL)

    abort "QEMU exited before SSH became ready" unless wait_thread.alive?
    abort "QEMU guest Wine smoke timed out waiting for SSH" if Process.clock_gettime(Process::CLOCK_MONOTONIC) >= deadline

    sleep RETRY_INTERVAL_SECONDS
  end

  smoke_stdout, smoke_stderr, smoke_status = run_command(
    {
      "GOCACHE" => GO_CACHE_ROOT.join("build").to_s,
      "GOMODCACHE" => GO_CACHE_ROOT.join("mod").to_s
    },
    "go", "run", "./cmd/xnix-runtime-go", "windows-app-guest-wine-smoke",
    "--exe", EXE_PATH.to_s,
    "--key", Xnix::SshTestKey::PRIVATE_KEY_PATH,
    "--timeout", "90s"
  )

  unless smoke_status.zero?
    warn smoke_stdout unless smoke_stdout.empty?
    warn smoke_stderr unless smoke_stderr.empty?
    warn "FAIL: QEMU guest Wine smoke command failed"
    exit 1
  end

  payload = JSON.parse(smoke_stdout)
  case payload.fetch("status")
  when "passed"
    if payload["marker_observed"] && payload["stdout"].include?(MARKER)
      puts "PASS: QEMU guest real Windows app Wine smoke"
      exit 0
    end
    warn "FAIL: QEMU guest Windows app smoke marker missing"
    exit 1
  when "skipped"
    puts "SKIP: QEMU guest real Windows app Wine smoke (#{payload.fetch("skip_reason")})"
    exit 0
  else
    warn smoke_stdout
    warn smoke_stderr unless smoke_stderr.empty?
    warn "FAIL: QEMU guest real Windows app Wine smoke"
    exit 1
  end
ensure
  stop_qemu(wait_thread)
  output.close unless output.closed?
  reader.join
end
