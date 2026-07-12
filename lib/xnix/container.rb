# frozen_string_literal: true

module Xnix
  class Container
    BUILD_MEMORY_LIMIT = "800m"
    CPU_LIMIT = "1.0"
    PROCESS_LIMIT = "256"
    TEMPORARY_FILESYSTEM_SIZE = "64m"
    IMAGE_NAME = "xnix-builder"

    def initialize(project_root:, version:)
      @project_root = project_root
      @version = version
    end

    def image_tag
      "#{IMAGE_NAME}:#{@version}"
    end

    def build_command
      [
        "docker", "build",
        "--memory", BUILD_MEMORY_LIMIT,
        "--cpus", CPU_LIMIT,
        "--tag", image_tag,
        "--file", File.join(@project_root, "Dockerfile"),
        @project_root
      ]
    end

    def offline_run_command(command)
      [
        "docker", "run", "--rm", "--init",
        "--memory", BUILD_MEMORY_LIMIT,
        "--cpus", CPU_LIMIT,
        "--pids-limit", PROCESS_LIMIT,
        "--cap-drop", "ALL",
        "--security-opt", "no-new-privileges",
        "--network", "none",
        "--read-only",
        "--tmpfs", "/tmp:rw,noexec,nosuid,size=#{TEMPORARY_FILESYSTEM_SIZE}",
        "--mount", workspace_mount,
        image_tag,
        *command
      ]
    end

    private

    def workspace_mount
      "type=bind,source=#{@project_root},target=/workspace,rw"
    end
  end
end
