#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/container"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath.to_s
VERSION = "0.2.300"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

container = Xnix::Container.new(project_root: PROJECT_ROOT, version: VERSION)
custom_container = Xnix::Container.new(project_root: PROJECT_ROOT, version: VERSION, docker_bin: "/tmp/xnix-docker")
build_command = container.build_command
build_tools_command = container.build_tools_command
offline_command = container.offline_run_command(["ruby", "scripts/verify_layout.rb"])
custom_build_command = custom_container.build_command
custom_build_tools_command = custom_container.build_tools_command
custom_offline_command = custom_container.offline_run_command(["ruby", "scripts/verify_layout.rb"])
runtime_activation_command = container.runtime_activation_smoke_command
runtime_dbus_command = container.runtime_dbus_smoke_command
runtime_owner_candidate_command = container.runtime_owner_candidate_smoke_command
kde_center_dbus_command = container.kde_center_dbus_smoke_command
known_winapp_fetch_command = container.known_winapp_fetch_command
known_winapp_guest_command = container.known_winapp_guest_wine_smoke_command
staged_launcher_dispatch_command = container.staged_launcher_dispatch_smoke_command
runtime_status_owner_service_session_bus_command = container.runtime_status_owner_service_session_bus_smoke_command
kde_controlled_launch_action_smoke_command = container.kde_controlled_launch_action_smoke_command
kde_controlled_launch_action_dbus_fixture_smoke_command = container.kde_controlled_launch_action_dbus_fixture_smoke_command
dbus_controlled_launch_owner_fixture_command = container.dbus_controlled_launch_owner_fixture_smoke_command
dockerignore_entries = Pathname.new(PROJECT_ROOT).join(".dockerignore").read.lines.map(&:strip)
dockerfile = Pathname.new(PROJECT_ROOT).join("Dockerfile").read
expected_source_mount = "type=volume,source=#{Xnix::Container::SOURCE_CACHE_VOLUME},target=/workspace/.cache"

assert(dockerignore_entries.include?(".cache/"), "Docker build context must exclude the managed Buildroot cache")
assert(dockerignore_entries.include?(".gocache/"), "Docker build context must exclude the local Go build cache")
tools_target_index = dockerfile.index("AS tools")
tested_runtime_target_index = dockerfile.index("AS tested-runtime")
go_test_index = dockerfile.index("go test -timeout 90m ./...")
assert(!tools_target_index.nil?, "Dockerfile must define a lightweight tools target")
assert(!tested_runtime_target_index.nil?, "Dockerfile must define the tested Runtime target")
assert(!go_test_index.nil?, "Dockerfile must retain full Go validation in the tested Runtime target")
assert(tools_target_index < tested_runtime_target_index, "Dockerfile tools target must be available before the tested Runtime target")
assert(tested_runtime_target_index < go_test_index, "Dockerfile tools target must not run the full Go suite")

assert(build_command.first(2) == ["docker", "build"], "build command must invoke docker build")
assert(build_tools_command.first(2) == ["docker", "build"], "tools build command must invoke docker build")
assert(custom_build_command.first(2) == ["/tmp/xnix-docker", "build"], "build command must support a custom Docker CLI")
assert(custom_build_tools_command.first(2) == ["/tmp/xnix-docker", "build"], "tools build command must support a custom Docker CLI")
assert(custom_offline_command.first(2) == ["/tmp/xnix-docker", "run"], "runtime command must support a custom Docker CLI")
assert(build_command.include?(container.image_tag), "build command must use the versioned image tag")
assert(build_command.include?("--target"), "build command must select an explicit Docker target")
assert(build_command.fetch(build_command.index("--target") + 1) == "tested-runtime", "build command must retain full validation in the tested Runtime target")
assert(build_tools_command.include?(container.tools_image_tag), "tools build command must use the versioned tools image tag")
assert(build_tools_command.include?("--target"), "tools build command must select an explicit Docker target")
assert(build_tools_command.fetch(build_tools_command.index("--target") + 1) == "tools", "tools build command must avoid the full validation target")
assert(build_command.include?("--pull=false"), "build command must prefer the local base image cache")
assert(build_tools_command.include?("--pull=false"), "tools build command must prefer the local base image cache")
assert(!build_command.include?("--memory"), "Buildx must not receive an unsupported memory argument")
assert(!build_command.include?("--cpus"), "Buildx must not receive an unsupported CPU argument")
assert(!build_tools_command.include?("--memory"), "tools Buildx must not receive an unsupported memory argument")
assert(!build_tools_command.include?("--cpus"), "tools Buildx must not receive an unsupported CPU argument")

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

assert(runtime_owner_candidate_command.fetch(runtime_owner_candidate_command.index("--network") + 1) == "none", "Runtime owner candidate smoke must run without networking")
assert(runtime_owner_candidate_command.include?("--read-only"), "Runtime owner candidate smoke must keep the container root read-only")
assert(runtime_owner_candidate_command.last(2) == ["ruby", "scripts/runtime_owner_candidate_smoke.rb"], "Runtime owner candidate smoke must run the restricted session smoke")

assert(kde_center_dbus_command.fetch(kde_center_dbus_command.index("--network") + 1) == "none", "KDE center D-Bus smoke must run without networking")
assert(kde_center_dbus_command.include?("--read-only"), "KDE center D-Bus smoke must keep the container root read-only")
assert(kde_center_dbus_command.last(2) == ["ruby", "scripts/kde_center_dbus_smoke.rb"], "KDE center D-Bus smoke must run the model session bus smoke")

assert(known_winapp_fetch_command.fetch(known_winapp_fetch_command.index("--network") + 1) == "bridge", "known Windows app fetch must use explicit bridge networking")
assert(known_winapp_fetch_command.include?("--read-only"), "known Windows app fetch must keep the container root read-only")
assert(known_winapp_fetch_command.include?("--mount"), "known Windows app fetch must mount its managed cache volume")
known_fetch_mount = known_winapp_fetch_command.fetch(known_winapp_fetch_command.index("--mount") + 1)
assert(known_fetch_mount == expected_source_mount, "known Windows app fetch must use the managed source cache volume")
assert(!known_fetch_mount.include?("type=bind"), "known Windows app fetch must not bind mount a host directory")
assert(known_winapp_fetch_command.last(2) == ["ruby", "scripts/known_winapp_fetch.rb"], "known Windows app fetch must run the fetch harness")

assert(known_winapp_guest_command.fetch(known_winapp_guest_command.index("--network") + 1) == "none", "known Windows app guest smoke must run without container networking")
assert(known_winapp_guest_command.include?("--read-only"), "known Windows app guest smoke must keep the container root read-only")
assert(known_winapp_guest_command.include?("--mount"), "known Windows app guest smoke must mount its managed cache volume")
known_guest_mount = known_winapp_guest_command.fetch(known_winapp_guest_command.index("--mount") + 1)
assert(known_guest_mount == expected_source_mount, "known Windows app guest smoke must use the managed source cache volume")
assert(!known_guest_mount.include?("type=bind"), "known Windows app guest smoke must not bind mount a host directory")
assert(!known_winapp_guest_command.include?("--privileged"), "known Windows app guest smoke must not be privileged")
assert(!known_winapp_guest_command.any? { |argument| argument.include?("docker.sock") }, "known Windows app guest smoke must not mount the Docker socket")
assert(known_winapp_guest_command.last(2) == ["ruby", "scripts/known_winapp_guest_wine_smoke.rb"], "known Windows app guest smoke must run the QEMU Wine harness")

assert(staged_launcher_dispatch_command.fetch(staged_launcher_dispatch_command.index("--network") + 1) == "none", "staged launcher dispatch smoke must run without container networking")
assert(staged_launcher_dispatch_command.include?("--read-only"), "staged launcher dispatch smoke must keep the container root read-only")
assert(staged_launcher_dispatch_command.include?("--mount"), "staged launcher dispatch smoke must mount its managed cache volume")
staged_launcher_dispatch_mount = staged_launcher_dispatch_command.fetch(staged_launcher_dispatch_command.index("--mount") + 1)
assert(staged_launcher_dispatch_mount == expected_source_mount, "staged launcher dispatch smoke must use the managed source cache volume")
assert(!staged_launcher_dispatch_mount.include?("type=bind"), "staged launcher dispatch smoke must not bind mount a host directory")
assert(!staged_launcher_dispatch_command.include?("--privileged"), "staged launcher dispatch smoke must not be privileged")
assert(!staged_launcher_dispatch_command.any? { |argument| argument.include?("docker.sock") }, "staged launcher dispatch smoke must not mount the Docker socket")
assert(staged_launcher_dispatch_command.last(2) == ["ruby", "scripts/staged_launcher_dispatch_smoke.rb"], "staged launcher dispatch smoke must run through the staged launcher harness")

assert(runtime_status_owner_service_session_bus_command.fetch(runtime_status_owner_service_session_bus_command.index("--network") + 1) == "none", "Runtime-status owner service session-bus smoke must run without container networking")
assert(runtime_status_owner_service_session_bus_command.include?("--read-only"), "Runtime-status owner service session-bus smoke must keep the container root read-only")
assert(runtime_status_owner_service_session_bus_command.include?("--mount"), "Runtime-status owner service session-bus smoke must mount its managed cache volume")
runtime_status_owner_service_session_bus_mount = runtime_status_owner_service_session_bus_command.fetch(runtime_status_owner_service_session_bus_command.index("--mount") + 1)
assert(runtime_status_owner_service_session_bus_mount == expected_source_mount, "Runtime-status owner service session-bus smoke must use the managed source cache volume")
assert(!runtime_status_owner_service_session_bus_mount.include?("type=bind"), "Runtime-status owner service session-bus smoke must not bind mount a host directory")
assert(!runtime_status_owner_service_session_bus_command.include?("--privileged"), "Runtime-status owner service session-bus smoke must not be privileged")
assert(!runtime_status_owner_service_session_bus_command.any? { |argument| argument.include?("docker.sock") }, "Runtime-status owner service session-bus smoke must not mount the Docker socket")
assert(runtime_status_owner_service_session_bus_command.last(2) == ["ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"], "Runtime-status owner service session-bus smoke must run through the session-bus harness")

assert(kde_controlled_launch_action_smoke_command.fetch(kde_controlled_launch_action_smoke_command.index("--network") + 1) == "none", "KDE controlled launch action smoke must run without container networking")
assert(kde_controlled_launch_action_smoke_command.include?("--read-only"), "KDE controlled launch action smoke must keep the container root read-only")
assert(kde_controlled_launch_action_smoke_command.include?("--mount"), "KDE controlled launch action smoke must mount its managed cache volume")
kde_controlled_launch_action_smoke_mount = kde_controlled_launch_action_smoke_command.fetch(kde_controlled_launch_action_smoke_command.index("--mount") + 1)
assert(kde_controlled_launch_action_smoke_mount == expected_source_mount, "KDE controlled launch action smoke must use the managed source cache volume")
assert(!kde_controlled_launch_action_smoke_mount.include?("type=bind"), "KDE controlled launch action smoke must not bind mount a host directory")
assert(!kde_controlled_launch_action_smoke_command.include?("--privileged"), "KDE controlled launch action smoke must not be privileged")
assert(!kde_controlled_launch_action_smoke_command.any? { |argument| argument.include?("docker.sock") }, "KDE controlled launch action smoke must not mount the Docker socket")
assert(kde_controlled_launch_action_smoke_command.last(2) == ["ruby", "scripts/kde_controlled_launch_action_smoke.rb"], "KDE controlled launch action smoke must run through the KDE action harness")

assert(kde_controlled_launch_action_dbus_fixture_smoke_command.fetch(kde_controlled_launch_action_dbus_fixture_smoke_command.index("--network") + 1) == "none", "KDE controlled launch action D-Bus fixture smoke must run without container networking")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.include?("--read-only"), "KDE controlled launch action D-Bus fixture smoke must keep the container root read-only")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.include?("--mount"), "KDE controlled launch action D-Bus fixture smoke must mount its managed cache volume")
kde_controlled_launch_action_dbus_fixture_mount = kde_controlled_launch_action_dbus_fixture_smoke_command.fetch(kde_controlled_launch_action_dbus_fixture_smoke_command.index("--mount") + 1)
assert(kde_controlled_launch_action_dbus_fixture_mount == expected_source_mount, "KDE controlled launch action D-Bus fixture smoke must use the managed source cache volume")
assert(!kde_controlled_launch_action_dbus_fixture_mount.include?("type=bind"), "KDE controlled launch action D-Bus fixture smoke must not bind mount a host directory")
kde_controlled_launch_action_dbus_fixture_tmpfs = kde_controlled_launch_action_dbus_fixture_smoke_command.each_with_index.filter_map do |argument, index|
  kde_controlled_launch_action_dbus_fixture_smoke_command.fetch(index + 1) if argument == "--tmpfs"
end
assert(kde_controlled_launch_action_dbus_fixture_tmpfs.include?("/workspace/.xnix-dbus-controlled-launch-scratch:rw,exec,nosuid,size=#{Xnix::Container::CONTROLLED_LAUNCH_SCRATCH_SIZE_BYTES},mode=1777"), "KDE controlled launch action D-Bus fixture smoke must use an executable managed tmpfs scratch mount")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.include?("--env"), "KDE controlled launch action D-Bus fixture smoke must explicitly opt in to D-Bus fixture execution")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.include?("#{Xnix::Container::KDE_CONTROLLED_LAUNCH_ACTION_DBUS_FIXTURE_ENV}=1"), "KDE controlled launch action D-Bus fixture smoke must set the D-Bus fixture execution env")
assert(!kde_controlled_launch_action_dbus_fixture_smoke_command.include?("--privileged"), "KDE controlled launch action D-Bus fixture smoke must not be privileged")
assert(!kde_controlled_launch_action_dbus_fixture_smoke_command.any? { |argument| argument.include?("docker.sock") }, "KDE controlled launch action D-Bus fixture smoke must not mount the Docker socket")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.include?(container.image_tag), "KDE controlled launch action D-Bus fixture smoke must use the tested Runtime image")
assert(kde_controlled_launch_action_dbus_fixture_smoke_command.last(2) == ["ruby", "scripts/kde_controlled_launch_action_smoke.rb"], "KDE controlled launch action D-Bus fixture smoke must run through the KDE action harness")

assert(dbus_controlled_launch_owner_fixture_command.fetch(dbus_controlled_launch_owner_fixture_command.index("--network") + 1) == "none", "D-Bus controlled launch owner fixture smoke must run without container networking")
assert(dbus_controlled_launch_owner_fixture_command.include?("--read-only"), "D-Bus controlled launch owner fixture smoke must keep the container root read-only")
assert(dbus_controlled_launch_owner_fixture_command.include?("--mount"), "D-Bus controlled launch owner fixture smoke must mount its managed cache volume")
dbus_controlled_launch_owner_fixture_mount = dbus_controlled_launch_owner_fixture_command.fetch(dbus_controlled_launch_owner_fixture_command.index("--mount") + 1)
assert(dbus_controlled_launch_owner_fixture_mount == expected_source_mount, "D-Bus controlled launch owner fixture smoke must use the managed source cache volume")
assert(!dbus_controlled_launch_owner_fixture_mount.include?("type=bind"), "D-Bus controlled launch owner fixture smoke must not bind mount a host directory")
dbus_controlled_launch_owner_fixture_mounts = dbus_controlled_launch_owner_fixture_command.each_with_index.filter_map do |argument, index|
  dbus_controlled_launch_owner_fixture_command.fetch(index + 1) if argument == "--mount"
end
assert(dbus_controlled_launch_owner_fixture_mounts.all? { |mount| !mount.include?("type=bind") }, "D-Bus controlled launch owner fixture smoke mounts must not bind host directories")
dbus_controlled_launch_owner_fixture_tmpfs = dbus_controlled_launch_owner_fixture_command.each_with_index.filter_map do |argument, index|
  dbus_controlled_launch_owner_fixture_command.fetch(index + 1) if argument == "--tmpfs"
end
assert(dbus_controlled_launch_owner_fixture_tmpfs.include?("/workspace/.xnix-dbus-controlled-launch-scratch:rw,exec,nosuid,size=#{Xnix::Container::CONTROLLED_LAUNCH_SCRATCH_SIZE_BYTES},mode=1777"), "D-Bus controlled launch owner fixture smoke must use an executable managed tmpfs scratch mount")
assert(!dbus_controlled_launch_owner_fixture_command.include?("--privileged"), "D-Bus controlled launch owner fixture smoke must not be privileged")
assert(!dbus_controlled_launch_owner_fixture_command.any? { |argument| argument.include?("docker.sock") }, "D-Bus controlled launch owner fixture smoke must not mount the Docker socket")
assert(dbus_controlled_launch_owner_fixture_command.include?(container.image_tag), "D-Bus controlled launch owner fixture smoke must use the tested Runtime image")
assert(dbus_controlled_launch_owner_fixture_command.last(2) == ["ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"], "D-Bus controlled launch owner fixture smoke must run through the D-Bus owner fixture harness")

observed = container.observed_cache_run_command(name: "xnix-full-build-test", command: ["make"])
assert(observed.include?("--detach"), "observed build must run detached")
assert(observed.include?("--name"), "observed build must have a stable name")
assert(!observed.include?("--rm"), "observed build must retain its logs after exit")
assert(observed.include?(container.tools_image_tag), "observed Buildroot build must use the tools image")

source_command = container.source_retrieval_command(["ruby", "scripts/fetch_buildroot.rb"])
assert(source_command.include?("--network"), "source retrieval must configure networking")
assert(source_command.fetch(source_command.index("--network") + 1) == "bridge", "source retrieval must use bridge networking")
assert(source_command.include?("--mount"), "source retrieval must mount its internal source cache")
assert(source_command.include?(container.tools_image_tag), "source retrieval must use the tools image")
source_mount = source_command.fetch(source_command.index("--mount") + 1)
assert(source_mount == expected_source_mount, "source retrieval must use the managed source cache volume")
assert(!source_mount.include?("type=bind"), "source retrieval must not bind mount a host directory")

puts "PASS: constrained container command unit tests"
