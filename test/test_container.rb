#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/container"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath.to_s
VERSION = "0.1.3"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

container = Xnix::Container.new(project_root: PROJECT_ROOT, version: VERSION)
build_command = container.build_command
offline_command = container.offline_run_command(["ruby", "scripts/verify_layout.rb"])

assert(build_command.first(2) == ["docker", "build"], "build command must invoke docker build")
assert(build_command.include?("--memory"), "build command must set a memory limit")
assert(build_command.include?(Xnix::Container::BUILD_MEMORY_LIMIT), "build command must use the configured memory limit")
assert(build_command.include?("--cpus"), "build command must set a CPU limit")
assert(build_command.include?(Xnix::Container::CPU_LIMIT), "build command must use the configured CPU limit")
assert(build_command.include?(container.image_tag), "build command must use the versioned image tag")

assert(offline_command.include?("--cap-drop"), "offline container must drop capabilities")
assert(offline_command.include?("ALL"), "offline container must drop all capabilities")
assert(offline_command.include?("--security-opt"), "offline container must set no-new-privileges")
assert(offline_command.include?("no-new-privileges"), "offline container must forbid privilege escalation")
assert(offline_command.include?("--network"), "offline container must configure networking")
assert(offline_command.fetch(offline_command.index("--network") + 1) == "none", "offline container must have no network")
assert(offline_command.include?("--read-only"), "offline container root filesystem must be read-only")
assert(offline_command.include?("--pids-limit"), "offline container must set a process limit")
assert(offline_command.include?(Xnix::Container::PROCESS_LIMIT), "offline container must use the configured process limit")
assert(!offline_command.include?("--privileged"), "offline container must not be privileged")
assert(!offline_command.include?("--network=host"), "offline container must not use host networking")
assert(!offline_command.any? { |argument| argument.include?("docker.sock") }, "offline container must not mount the Docker socket")

workspace_mount = offline_command.fetch(offline_command.index("--mount") + 1)
expected_mount = "type=bind,source=#{PROJECT_ROOT},target=/workspace,rw"
assert(workspace_mount == expected_mount, "offline container must mount only the project workspace")

puts "PASS: constrained container command unit tests"
