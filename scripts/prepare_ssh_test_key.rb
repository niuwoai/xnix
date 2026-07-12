#!/usr/bin/env ruby
# frozen_string_literal: true

require "fileutils"
require_relative "../lib/xnix/ssh_test_key"

private_key = Xnix::SshTestKey::PRIVATE_KEY_PATH
public_key = Xnix::SshTestKey::PUBLIC_KEY_PATH

if File.file?(private_key) && File.file?(public_key)
  puts "SSH test key already prepared in the managed cache volume"
  exit 0
end

FileUtils.mkdir_p(Xnix::SshTestKey::DIRECTORY, mode: 0o700)
FileUtils.rm_f(private_key)
FileUtils.rm_f(public_key)

success = system(
  "ssh-keygen", "-q", "-t", "ed25519", "-N", "", "-f", private_key,
  "-C", Xnix::SshTestKey::COMMENT
)
abort "SSH test key generation failed" unless success

File.chmod(0o600, private_key)
File.chmod(0o644, public_key)
puts "SSH test key prepared in the managed cache volume"
