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

wine_command = Xnix::Qemu.wine_guest.boot_command(ssh: true)
wine_network = wine_command.fetch(wine_command.index("-netdev") + 1)
assert(wine_command.first == "qemu-system-i386", "Wine guest QEMU command must use i386 system emulation")
assert(wine_command.fetch(wine_command.index("-kernel") + 1) == Xnix::Qemu::WINE_KERNEL_IMAGE, "Wine guest QEMU command must boot the Wine kernel")
assert(wine_command.fetch(wine_command.index("-m") + 1) == Xnix::Qemu::WINE_MEMORY, "Wine guest QEMU command must use the Wine memory limit")
assert(wine_command.fetch(wine_command.index("-smp") + 1) == Xnix::Qemu::WINE_CPU_COUNT, "Wine guest QEMU command must use the Wine CPU limit")
assert(wine_network.include?("hostfwd=tcp:127.0.0.1:2222-:22"), "Wine guest SSH forwarding must bind only the loopback address")
assert(!wine_network.include?("0.0.0.0"), "Wine guest SSH forwarding must never bind all interfaces")

puts "PASS: QEMU command unit tests"
