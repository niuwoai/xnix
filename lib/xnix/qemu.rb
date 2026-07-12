# frozen_string_literal: true

module Xnix
  class Qemu
    MEMORY = "512M"
    CPU_COUNT = "1"
    KERNEL_IMAGE = "/workspace/.cache/xnix-output/images/bzImage"

    def boot_command
      [
        "qemu-system-x86_64",
        "-machine", "q35,accel=tcg",
        "-cpu", "max",
        "-m", MEMORY,
        "-smp", CPU_COUNT,
        "-nographic",
        "-serial", "mon:stdio",
        "-no-reboot",
        "-kernel", KERNEL_IMAGE,
        "-append", "console=ttyS0,115200 panic=-1",
        "-netdev", "user,id=net0,restrict=on",
        "-device", "e1000,netdev=net0"
      ]
    end
  end
end
