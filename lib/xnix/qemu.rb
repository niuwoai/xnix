# frozen_string_literal: true

module Xnix
  class Qemu
    MEMORY = "512M"
    CPU_COUNT = "1"
    KERNEL_IMAGE = "/workspace/.cache/xnix-output/images/bzImage"

    def boot_command(ssh: false)
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
