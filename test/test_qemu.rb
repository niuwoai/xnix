#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/qemu"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

command = Xnix::Qemu.new.boot_command
assert(command.first == "qemu-system-x86_64", "QEMU command must use x86_64 system emulation")
assert(command.include?("q35,accel=tcg"), "QEMU command must use TCG software emulation")
assert(command.fetch(command.index("-m") + 1) == Xnix::Qemu::MEMORY, "QEMU command must limit memory")
assert(command.fetch(command.index("-smp") + 1) == Xnix::Qemu::CPU_COUNT, "QEMU command must limit CPUs")
assert(command.include?("-nographic"), "QEMU command must use the serial console")
assert(command.include?("-no-reboot"), "QEMU command must not reboot indefinitely")
assert(command.include?("user,id=net0,restrict=on"), "QEMU network must remain restricted")
assert(!command.any? { |argument| argument.include?("hostfwd") }, "initial boot command must not expose host ports")

ssh_command = Xnix::Qemu.new.boot_command(ssh: true)
ssh_network = ssh_command.fetch(ssh_command.index("-netdev") + 1)
assert(ssh_network.include?("hostfwd=tcp:127.0.0.1:2222-:22"), "SSH forwarding must bind only the loopback address")
assert(!ssh_network.include?("0.0.0.0"), "SSH forwarding must never bind all interfaces")

puts "PASS: QEMU command unit tests"
