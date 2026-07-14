#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/container"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath.to_s
VERSION = "0.2.145"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

container = Xnix::Container.new(project_root: PROJECT_ROOT, version: VERSION)
custom_container = Xnix::Container.new(project_root: PROJECT_ROOT, version: VERSION, docker_bin: "/tmp/xnix-docker")
build_command = container.build_command
offline_command = container.offline_run_command(["ruby", "scripts/verify_layout.rb"])
custom_build_command = custom_container.build_command
custom_offline_command = custom_container.offline_run_command(["ruby", "scripts/verify_layout.rb"])
runtime_activation_command = container.runtime_activation_smoke_command
runtime_dbus_command = container.runtime_dbus_smoke_command
kde_center_dbus_command = container.kde_center_dbus_smoke_command

assert(build_command.first(2) == ["docker", "build"], "build command must invoke docker build")
assert(custom_build_command.first(2) == ["/tmp/xnix-docker", "build"], "build command must support a custom Docker CLI")
assert(custom_offline_command.first(2) == ["/tmp/xnix-docker", "run"], "runtime command must support a custom Docker CLI")
assert(build_command.include?(container.image_tag), "build command must use the versioned image tag")
assert(build_command.include?("--pull=false"), "build command must prefer the local base image cache")
assert(!build_command.include?("--memory"), "Buildx must not receive an unsupported memory argument")
assert(!build_command.include?("--cpus"), "Buildx must not receive an unsupported CPU argument")

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

assert(!offline_command.include?("--mount"), "offline container must not mount host directories")

assert(runtime_activation_command.fetch(runtime_activation_command.index("--network") + 1) == "none", "runtime activation smoke must run without networking")
assert(runtime_activation_command.include?("--read-only"), "runtime activation smoke must keep the container root read-only")
assert(runtime_activation_command.last(2) == ["ruby", "scripts/runtime_activation_smoke.rb"], "runtime activation smoke must run the packaged activation smoke")

assert(runtime_dbus_command.fetch(runtime_dbus_command.index("--network") + 1) == "none", "runtime D-Bus smoke must run without networking")
assert(runtime_dbus_command.include?("--read-only"), "runtime D-Bus smoke must keep the container root read-only")
assert(runtime_dbus_command.last(2) == ["ruby", "scripts/dbus_session_smoke.rb"], "runtime D-Bus smoke must run the session bus smoke")

assert(kde_center_dbus_command.fetch(kde_center_dbus_command.index("--network") + 1) == "none", "KDE center D-Bus smoke must run without networking")
assert(kde_center_dbus_command.include?("--read-only"), "KDE center D-Bus smoke must keep the container root read-only")
assert(kde_center_dbus_command.last(2) == ["ruby", "scripts/kde_center_dbus_smoke.rb"], "KDE center D-Bus smoke must run the model session bus smoke")

observed = container.observed_cache_run_command(name: "xnix-full-build-test", command: ["make"])
assert(observed.include?("--detach"), "observed build must run detached")
assert(observed.include?("--name"), "observed build must have a stable name")
assert(!observed.include?("--rm"), "observed build must retain its logs after exit")

source_command = container.source_retrieval_command(["ruby", "scripts/fetch_buildroot.rb"])
assert(source_command.include?("--network"), "source retrieval must configure networking")
assert(source_command.fetch(source_command.index("--network") + 1) == "bridge", "source retrieval must use bridge networking")
assert(source_command.include?("--mount"), "source retrieval must mount its internal source cache")
source_mount = source_command.fetch(source_command.index("--mount") + 1)
expected_source_mount = "type=volume,source=#{Xnix::Container::SOURCE_CACHE_VOLUME},target=/workspace/.cache"
assert(source_mount == expected_source_mount, "source retrieval must use the managed source cache volume")
assert(!source_mount.include?("type=bind"), "source retrieval must not bind mount a host directory")

puts "PASS: constrained container command unit tests"
