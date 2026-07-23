# frozen_string_literal: true

module Xnix
  class Qemu
    MEMORY = "512M"
    WINE_MEMORY = "1024M"
    CPU_COUNT = "1"
    WINE_CPU_COUNT = "2"
    CPU_MODEL = "max"
    WINE_CPU_MODEL = "qemu32"
    BINARY = "qemu-system-x86_64"
    WINE_BINARY = "qemu-system-i386"
    KERNEL_IMAGE = "/workspace/.cache/xnix-output/images/bzImage"
    WINE_KERNEL_IMAGE = "/workspace/.cache/xnix-wine-i386-output/images/bzImage"

    attr_reader :binary, :kernel_image, :memory, :cpu_count, :cpu_model

    def self.wine_guest
      new(
        binary: WINE_BINARY,
        kernel_image: WINE_KERNEL_IMAGE,
        memory: WINE_MEMORY,
        cpu_count: WINE_CPU_COUNT,
        cpu_model: WINE_CPU_MODEL
      )
    end

    def initialize(binary: BINARY, kernel_image: KERNEL_IMAGE, memory: MEMORY, cpu_count: CPU_COUNT, cpu_model: CPU_MODEL)
      @binary = binary
      @kernel_image = kernel_image
      @memory = memory
      @cpu_count = cpu_count
      @cpu_model = cpu_model
    end

    def boot_command(ssh: false)
      [
        binary,
        "-machine", "q35,accel=tcg",
        "-cpu", cpu_model,
        "-m", memory,
        "-smp", cpu_count,
        "-nographic",
        "-serial", "mon:stdio",
        "-no-reboot",
        "-kernel", kernel_image,
        "-append", "console=ttyS0,115200 panic=-1",
        "-netdev", network_definition(ssh),
        "-device", "e1000,netdev=net0"
      ]
    end

    private

    def network_definition(ssh)
      return "user,id=net0,restrict=on" unless ssh

      "user,id=net0,restrict=on,hostfwd=tcp:127.0.0.1:2222-:22"
    end
  end
end
