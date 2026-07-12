#!/usr/bin/env ruby
# frozen_string_literal: true

require "open3"
require "fileutils"
require_relative "../lib/xnix/qemu"
require_relative "../lib/xnix/ssh_probe"
require_relative "../lib/xnix/ssh_test_key"

BOOT_TIMEOUT_SECONDS = 45
RETRY_INTERVAL_SECONDS = 1

def stop_qemu(wait_thread)
  return unless wait_thread.alive?

  Process.kill("TERM", wait_thread.pid)
  wait_thread.join(5)
  Process.kill("KILL", wait_thread.pid) if wait_thread.alive?
rescue Errno::ESRCH
  nil
end

stdin, output, wait_thread = Open3.popen2e(*Xnix::Qemu.new.boot_command(ssh: true))
stdin.close
reader = Thread.new { output.read }
deadline = Process.clock_gettime(Process::CLOCK_MONOTONIC) + BOOT_TIMEOUT_SECONDS
probe = Xnix::SshProbe.new

begin
  loop do
    if system(*probe.command, out: File::NULL, err: File::NULL)
      puts "PASS: QEMU loopback SSH smoke test"
      exit 0
    end

    abort "QEMU exited before SSH became ready" unless wait_thread.alive?
    abort "QEMU SSH smoke test timed out" if Process.clock_gettime(Process::CLOCK_MONOTONIC) >= deadline

    sleep RETRY_INTERVAL_SECONDS
  end
ensure
  stop_qemu(wait_thread)
  output.close unless output.closed?
  reader.join
  FileUtils.rm_rf(Xnix::SshTestKey::DIRECTORY)
end
