#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/sshd"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

assert(Xnix::Sshd.new("Starting sshd: OK\n").started?, "successful sshd init output must be accepted")
assert(!Xnix::Sshd.new("Starting sshd: FAIL\n").started?, "failed sshd init output must be rejected")

puts "PASS: sshd service unit tests"
