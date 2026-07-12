#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/sshd"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

assert(Xnix::Sshd.new("Starting sshd: OK\n").started?, "successful sshd init output must be accepted")
assert(!Xnix::Sshd.new("Starting sshd: FAIL\n").started?, "failed sshd init output must be rejected")

project_root = Pathname.new(__dir__).join("..").realpath
config = project_root.join("buildroot/board/xnix/rootfs-overlay/etc/ssh/sshd_config").read
assert(config.include?("PasswordAuthentication no"), "sshd must disable password authentication")
assert(config.include?("PermitRootLogin prohibit-password"), "sshd must forbid root password login")

puts "PASS: sshd service unit tests"
