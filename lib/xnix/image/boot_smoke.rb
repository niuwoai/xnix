# frozen_string_literal: true

require_relative "kde_image"

module Xnix
  module Image
    # BootSmoke boots a produced KDE Plasma disk image under QEMU and
    # confirms it reaches the graphical login by scanning the serial log
    # for the markers declared in the manifest's boot_smoke section.
    #
    # The marker-scan logic is pure and unit-testable; only the QEMU
    # invocation needs a real image and a host with qemu-system-x86_64.
    class BootSmoke
      DEFAULT_MEMORY = "4G"
      DEFAULT_CPU_COUNT = "2"

      def initialize(image)
        @image = image
      end

      def expected_markers
        Array(@image.manifest.dig("boot_smoke", "expect_serial_markers"))
      end

      def timeout_seconds
        @image.manifest.dig("boot_smoke", "timeout_seconds") || 180
      end

      # Given captured serial output, report whether every expected marker
      # appeared.
      def booted?(serial_contents)
        expected_markers.all? { |marker| serial_contents.include?(marker) }
      end

      def missing_markers(serial_contents)
        expected_markers.reject { |marker| serial_contents.include?(marker) }
      end

      def summary(serial_contents)
        return "graphical login reached" if booted?(serial_contents)

        "missing markers: #{missing_markers(serial_contents).join(', ')}"
      end

      # QEMU command that boots a full disk image (qcow2/raw) with UEFI
      # firmware, headless, serial on stdio. Loopback-only networking per
      # the project's safety constraints.
      def boot_command(disk_path:, firmware_path:, memory: DEFAULT_MEMORY, cpu_count: DEFAULT_CPU_COUNT)
        [
          "qemu-system-x86_64",
          "-machine", "q35,accel=tcg",
          "-cpu", "max",
          "-m", memory,
          "-smp", cpu_count,
          "-nographic",
          "-serial", "mon:stdio",
          "-no-reboot",
          "-drive", "if=pflash,format=raw,readonly=on,file=#{firmware_path}",
          "-drive", "if=virtio,format=qcow2,file=#{disk_path}",
          "-netdev", "user,id=net0,restrict=on",
          "-device", "virtio-net-pci,netdev=net0"
        ]
      end
    end
  end
end
