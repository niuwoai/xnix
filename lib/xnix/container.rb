# frozen_string_literal: true

module Xnix
  class Container
    BUILD_MEMORY_LIMIT = "4g"
    CPU_LIMIT = "1.0"
    PROCESS_LIMIT = "256"
    TEMPORARY_FILESYSTEM_SIZE = "64m"
    IMAGE_NAME = "xnix-builder"
    TOOLS_IMAGE_NAME = "xnix-builder-tools"
    SOURCE_CACHE_VOLUME = "xnix-buildroot-cache"
    CONTROLLED_LAUNCH_SCRATCH_SIZE_BYTES = "67108864"
    DOCKER_ENV = "XNIX_DOCKER_BIN"
    DEFAULT_DOCKER_BIN = "docker"

    def self.docker_bin
      value = ENV.fetch(DOCKER_ENV, DEFAULT_DOCKER_BIN).to_s
      return DEFAULT_DOCKER_BIN if value.empty?

      value
    end

    def initialize(project_root:, version:, docker_bin: self.class.docker_bin)
      @project_root = project_root
      @version = version
      @docker_bin = docker_bin
    end

    def image_tag
      "#{IMAGE_NAME}:#{@version}"
    end

    def tools_image_tag
      "#{TOOLS_IMAGE_NAME}:#{@version}"
    end

    def build_command
      [
        @docker_bin, "build",
        "--pull=false",
        "--target", "tested-runtime",
        "--tag", image_tag,
        "--file", File.join(@project_root, "Dockerfile"),
        @project_root
      ]
    end

    def build_tools_command
      [
        @docker_bin, "build",
        "--pull=false",
        "--target", "tools",
        "--tag", tools_image_tag,
        "--file", File.join(@project_root, "Dockerfile"),
        @project_root
      ]
    end

    def offline_run_command(command)
      runtime_command(network: "none", extra_mounts: [], command: command)
    end

    def runtime_activation_smoke_command
      offline_run_command(["ruby", "scripts/runtime_activation_smoke.rb"])
    end

    def runtime_dbus_smoke_command
      offline_run_command(["ruby", "scripts/dbus_session_smoke.rb"])
    end

    def runtime_owner_candidate_smoke_command
      offline_run_command(["ruby", "scripts/runtime_owner_candidate_smoke.rb"])
    end

    def kde_center_dbus_smoke_command
      offline_run_command(["ruby", "scripts/kde_center_dbus_smoke.rb"])
    end

    def known_winapp_fetch_command
      networked_cache_run_command(["ruby", "scripts/known_winapp_fetch.rb"])
    end

    def known_winapp_guest_wine_smoke_command
      tools_cache_run_command(["ruby", "scripts/known_winapp_guest_wine_smoke.rb"])
    end

    def staged_launcher_dispatch_smoke_command
      tools_cache_run_command(["ruby", "scripts/staged_launcher_dispatch_smoke.rb"])
    end

    def runtime_status_owner_service_session_bus_smoke_command
      tools_cache_run_command(["ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"])
    end

    def dbus_controlled_launch_owner_fixture_smoke_command
      runtime_command(
        network: "none",
        extra_mounts: [source_cache_mount],
        extra_tmpfs: [dbus_controlled_launch_scratch_tmpfs],
        command: ["ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"],
        image: image_tag
      )
    end

    def source_retrieval_command(command)
      runtime_command(network: "bridge", extra_mounts: [source_cache_mount], command: command, image: tools_image_tag)
    end

    def cache_run_command(command)
      runtime_command(network: "none", extra_mounts: [source_cache_mount], command: command)
    end

    def tools_cache_run_command(command)
      runtime_command(network: "none", extra_mounts: [source_cache_mount], command: command, image: tools_image_tag)
    end

    def networked_cache_run_command(command)
      runtime_command(network: "bridge", extra_mounts: [source_cache_mount], command: command, image: tools_image_tag)
    end

    def observed_cache_run_command(name:, command:)
      runtime_command(network: "none", extra_mounts: [source_cache_mount], command: command, remove: false, name: name, detach: true, image: tools_image_tag)
    end

    private

    def runtime_command(network:, extra_mounts:, command:, remove: true, name: nil, detach: false, image: image_tag, extra_tmpfs: [])
      [
        @docker_bin, "run", *(remove ? ["--rm"] : []), *(detach ? ["--detach"] : []), *(name.nil? ? [] : ["--name", name]), "--init",
        "--memory", BUILD_MEMORY_LIMIT,
        "--cpus", CPU_LIMIT,
        "--pids-limit", PROCESS_LIMIT,
        "--cap-drop", "ALL",
        "--security-opt", "no-new-privileges",
        "--network", network,
        "--read-only",
        "--tmpfs", "/tmp:rw,noexec,nosuid,size=#{TEMPORARY_FILESYSTEM_SIZE}",
        *extra_tmpfs.flat_map { |tmpfs| ["--tmpfs", tmpfs] },
        *extra_mounts.flat_map { |mount| ["--mount", mount] },
        image,
        *command
      ]
    end

    def source_cache_mount
      "type=volume,source=#{SOURCE_CACHE_VOLUME},target=/workspace/.cache"
    end

    def dbus_controlled_launch_scratch_tmpfs
      "/workspace/.xnix-dbus-controlled-launch-scratch:rw,exec,nosuid,size=#{CONTROLLED_LAUNCH_SCRATCH_SIZE_BYTES},mode=1777"
    end
  end
end
