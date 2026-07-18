# frozen_string_literal: true

require_relative "kde_image"
require "shellwords"

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
      ANSI_ESCAPE = /\e(?:\[[0-?]*[ -\/]*[@-~]|\][^\a\e]*(?:\a|\e\\))/.freeze
      UNIT_NAME = /\A[A-Za-z0-9@_.:-]+\z/

      def initialize(image)
        @image = image
      end

      def expected_markers
        Array(@image.manifest.dig("boot_smoke", "expect_serial_markers"))
      end

      def active_units
        Array(@image.manifest.dig("boot_smoke", "active_units"))
      end

      def fail_marker
        @image.manifest.dig("boot_smoke", "fail_marker").to_s
      end

      def probe_command
        invalid = active_units.reject { |unit| unit.match?(UNIT_NAME) }
        raise ArgumentError, "invalid systemd unit name: #{invalid.first}" unless invalid.empty?

        checks = active_units.map { |unit| "systemctl is-active --quiet #{Shellwords.escape(unit)}" }
        pass_output = "printf 'XNIX_BOOT_%s\\n' PROBE_PASS"
        fail_output = "printf 'XNIX_BOOT_%s\\n' PROBE_FAIL"
        "#{checks.join(' && ')} && #{pass_output} || #{fail_output}"
      end

      def timeout_seconds
        @image.manifest.dig("boot_smoke", "timeout_seconds") || 180
      end

      # Given captured serial output, report whether every expected marker
      # appeared.
      def booted?(serial_contents)
        normalized = normalize_serial(serial_contents)
        expected_markers.all? { |marker| normalized.include?(marker) }
      end

      def missing_markers(serial_contents)
        normalized = normalize_serial(serial_contents)
        expected_markers.reject { |marker| normalized.include?(marker) }
      end

      def summary(serial_contents)
        return "systemd probe reported inactive unit(s)" if normalize_serial(serial_contents).include?(fail_marker)
        return "graphical login reached" if booted?(serial_contents)

        "missing markers: #{missing_markers(serial_contents).join(', ')}"
      end

      # QEMU command that boots a full disk image (qcow2/raw) with UEFI
      # firmware, headless, serial on stdio. Loopback-only networking per
      # the project's safety constraints.
      def boot_command(disk_path:, firmware_path:, memory: DEFAULT_MEMORY, cpu_count: DEFAULT_CPU_COUNT,
                       acceleration: "tcg")
        cpu = acceleration == "kvm" ? "host" : "max"
        [
          "qemu-system-x86_64",
          "-machine", "q35,accel=#{acceleration}",
          "-cpu", cpu,
          "-m", memory,
          "-smp", cpu_count,
          "-snapshot",
          "-nographic",
          "-serial", "mon:stdio",
          "-no-reboot",
          "-drive", "if=pflash,format=raw,readonly=on,file=#{firmware_path}",
          "-drive", "if=virtio,format=qcow2,file=#{disk_path}",
          "-netdev", "user,id=net0,restrict=on",
          "-device", "virtio-net-pci,netdev=net0"
        ]
      end

      private

      def normalize_serial(serial_contents)
        serial_contents.gsub(ANSI_ESCAPE, "").delete("\r")
      end
    end
  end
end
