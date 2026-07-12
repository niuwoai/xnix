# frozen_string_literal: true

module Xnix
  class Buildroot
    SOURCE_DIRECTORY = "/workspace/.cache/buildroot/buildroot-2025.02.15"
    OUTPUT_DIRECTORY = "/workspace/.cache/xnix-output"
    EXTERNAL_DIRECTORY = "/workspace/buildroot"
    DEFCONFIG = "xnix_x86_64_defconfig"

    def configure_command
      make_command("BR2_EXTERNAL=#{EXTERNAL_DIRECTORY}", DEFCONFIG)
    end

    def build_command
      make_command
    end

    private

    def make_command(*arguments)
      ["make", "-C", SOURCE_DIRECTORY, "O=#{OUTPUT_DIRECTORY}", *arguments]
    end
  end
end
