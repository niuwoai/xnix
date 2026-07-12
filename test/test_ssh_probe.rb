#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/ssh_probe"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

command = Xnix::SshProbe.new.command
assert(command.first == "ssh", "probe must use the SSH client")
assert(command.fetch(command.index("-p") + 1) == "2222", "probe must use the forwarded SSH port")
assert(command.include?("root@127.0.0.1"), "probe must target only loopback")
assert(command.include?("BatchMode=yes"), "probe must not prompt interactively")
assert(command.include?("ConnectTimeout=5"), "probe must fail promptly")
assert(!command.any? { |argument| argument.include?("0.0.0.0") }, "probe must not target all interfaces")

puts "PASS: SSH probe unit tests"
