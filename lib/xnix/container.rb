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
    KDE_CONTROLLED_LAUNCH_ACTION_DBUS_FIXTURE_ENV = "XNIX_KDE_CONTROLLED_LAUNCH_ACTION_SMOKE_EXECUTE_DBUS_FIXTURE"
    DOCKER_ENV = "XNIX_DOCKER_BIN"
    TOOLS_BASE_IMAGE_ENV = "XNIX_TOOLS_BASE_IMAGE"
    MEMORY_LIMIT_ENV = "XNIX_CONTAINER_MEMORY_LIMIT"
    CPU_LIMIT_ENV = "XNIX_CONTAINER_CPU_LIMIT"
    KNOWN_WINAPP_FETCH_TIMEOUT_ENV = "XNIX_KNOWN_WINAPP_FETCH_TIMEOUT"
    DEFAULT_DOCKER_BIN = "docker"

    def self.docker_bin
      value = ENV.fetch(DOCKER_ENV, DEFAULT_DOCKER_BIN).to_s
      return DEFAULT_DOCKER_BIN if value.empty?

      value
    end

    def self.tools_base_image
      value = ENV.fetch(TOOLS_BASE_IMAGE_ENV, "").to_s.strip
      return nil if value.empty?

      value
    end

    def self.memory_limit
      value = ENV.fetch(MEMORY_LIMIT_ENV, BUILD_MEMORY_LIMIT).to_s.strip
      return BUILD_MEMORY_LIMIT if value.empty?

      value
    end

    def self.cpu_limit
      value = ENV.fetch(CPU_LIMIT_ENV, CPU_LIMIT).to_s.strip
      return CPU_LIMIT if value.empty?

      value
    end

    def self.known_winapp_fetch_timeout
      value = ENV.fetch(KNOWN_WINAPP_FETCH_TIMEOUT_ENV, "").to_s.strip
      return nil if value.empty?

      value
    end

    def initialize(project_root:, version:, docker_bin: self.class.docker_bin, tools_base_image: self.class.tools_base_image, memory_limit: self.class.memory_limit, cpu_limit: self.class.cpu_limit, known_winapp_fetch_timeout: self.class.known_winapp_fetch_timeout)
      @project_root = project_root
      @version = version
      @docker_bin = docker_bin
      @tools_base_image = tools_base_image
      @memory_limit = memory_limit
      @cpu_limit = cpu_limit
      @known_winapp_fetch_timeout = known_winapp_fetch_timeout
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
        *tools_base_image_build_arg,
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
        *tools_base_image_build_arg,
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
      networked_cache_run_command(
        ["ruby", "scripts/known_winapp_fetch.rb"],
        extra_env: known_winapp_fetch_env
      )
    end

    def known_winapp_guest_wine_smoke_command
      tools_cache_run_command(["ruby", "scripts/known_winapp_guest_wine_smoke.rb"])
    end

    def winapp_container_x_gui_smoke_command
      tools_cache_run_command(["ruby", "scripts/winapp_smoke.rb", "--backend", "container-x-gui"])
    end

    def staged_launcher_dispatch_smoke_command
      runtime_command(
        network: "none",
        extra_mounts: [source_cache_mount],
        command: ["ruby", "scripts/staged_launcher_dispatch_smoke.rb"],
        image: image_tag
      )
    end

    def staged_external_winapp_desktop_smoke_command
      ["ruby", "scripts/staged_desktop_external_winapp_smoke.rb", "--docker", @docker_bin]
    end

    def prepare_wine_smoke_image_command
      ["ruby", "scripts/build_wine_smoke_image.rb"]
    end

    def runtime_status_owner_service_session_bus_smoke_command
      tools_cache_run_command(["ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"])
    end

    def kde_controlled_launch_action_smoke_command
      tools_cache_run_command(["ruby", "scripts/kde_controlled_launch_action_smoke.rb"])
    end

    def desktop_trigger_request_preflight_smoke_command
      tools_cache_run_command(["ruby", "scripts/desktop_trigger_request_preflight_smoke.rb"])
    end

    def kde_controlled_launch_action_dbus_fixture_smoke_command
      runtime_command(
        network: "none",
        extra_mounts: [source_cache_mount],
        extra_tmpfs: [dbus_controlled_launch_scratch_tmpfs],
        extra_env: { KDE_CONTROLLED_LAUNCH_ACTION_DBUS_FIXTURE_ENV => "1" },
        command: ["ruby", "scripts/kde_controlled_launch_action_smoke.rb"],
        image: image_tag
      )
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

    def networked_cache_run_command(command, extra_env: {})
      runtime_command(network: "bridge", extra_mounts: [source_cache_mount], extra_env: extra_env, command: command, image: tools_image_tag)
    end

    def observed_cache_run_command(name:, command:)
      runtime_command(network: "none", extra_mounts: [source_cache_mount], command: command, remove: false, name: name, detach: true, image: tools_image_tag)
    end

    private

    def tools_base_image_build_arg
      @tools_base_image.nil? ? [] : ["--build-arg", "XNIX_TOOLS_BASE_IMAGE=#{@tools_base_image}"]
    end

    def known_winapp_fetch_env
      return {} if @known_winapp_fetch_timeout.nil? || @known_winapp_fetch_timeout.empty?

      { KNOWN_WINAPP_FETCH_TIMEOUT_ENV => @known_winapp_fetch_timeout }
    end

    def runtime_command(network:, extra_mounts:, command:, remove: true, name: nil, detach: false, image: image_tag, extra_tmpfs: [], extra_env: {})
      [
        @docker_bin, "run", *(remove ? ["--rm"] : []), *(detach ? ["--detach"] : []), *(name.nil? ? [] : ["--name", name]), "--init",
        "--memory", @memory_limit,
        "--cpus", @cpu_limit,
        "--pids-limit", PROCESS_LIMIT,
        "--cap-drop", "ALL",
        "--security-opt", "no-new-privileges",
        "--network", network,
        "--read-only",
        "--tmpfs", "/tmp:rw,noexec,nosuid,size=#{TEMPORARY_FILESYSTEM_SIZE}",
        *extra_tmpfs.flat_map { |tmpfs| ["--tmpfs", tmpfs] },
        *extra_env.flat_map { |key, value| ["--env", "#{key}=#{value}"] },
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
