# frozen_string_literal: true

module Xnix
  class Buildroot
    SOURCE_DIRECTORY = "/workspace/.cache/buildroot/buildroot-2025.02.15"
    DEFAULT_OUTPUT_DIRECTORY = "/workspace/.cache/xnix-output"
    WINE_OUTPUT_DIRECTORY = "/workspace/.cache/xnix-wine-i386-output"
    EXTERNAL_DIRECTORY = "/workspace/buildroot"
    DEFAULT_DEFCONFIG = "xnix_x86_64_defconfig"
    WINE_DEFCONFIG = "xnix_wine_i386_defconfig"
    OUTPUT_DIRECTORY = DEFAULT_OUTPUT_DIRECTORY
    DEFCONFIG = DEFAULT_DEFCONFIG
    SSH_TEST_PUBLIC_KEY = "/workspace/.cache/xnix-ssh-test/id_ed25519.pub"

    attr_reader :profile

    def initialize(profile: :default)
      @profile = profile
    end

    def configure_command
      make_command("BR2_EXTERNAL=#{EXTERNAL_DIRECTORY}", defconfig)
    end

    def build_command(ssh_test_key: false)
      make_command(*(ssh_test_key ? ["XNIX_TEST_SSH_PUBLIC_KEY=#{SSH_TEST_PUBLIC_KEY}"] : []))
    end

    def source_command
      make_command("source")
    end

    private

    def output_directory
      case profile
      when :default
        DEFAULT_OUTPUT_DIRECTORY
      when :wine_i386
        WINE_OUTPUT_DIRECTORY
      else
        raise ArgumentError, "unknown Buildroot profile: #{profile}"
      end
    end

    def defconfig
      case profile
      when :default
        DEFAULT_DEFCONFIG
      when :wine_i386
        WINE_DEFCONFIG
      else
        raise ArgumentError, "unknown Buildroot profile: #{profile}"
      end
    end

    def make_command(*arguments)
      ["make", "-C", SOURCE_DIRECTORY, "O=#{output_directory}", *arguments]
    end
  end
end
