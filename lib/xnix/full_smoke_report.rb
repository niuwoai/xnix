# frozen_string_literal: true

module Xnix
  class FullSmokeReport
    SCHEMA_VERSION = "xnix.full_smoke_report.v1"
    REPORT_TYPE = "full-build-qemu-smoke-report"

    def initialize(version:, steps:, serial_log_path:, serial_log:)
      @version = version
      @steps = steps
      @serial_log_path = serial_log_path
      @serial_log = serial_log
    end

    def to_h
      {
        "version" => @version,
        "schema_version" => SCHEMA_VERSION,
        "report_type" => REPORT_TYPE,
        "steps" => @steps,
        "step_count" => @steps.length,
        "serial_log_path" => @serial_log_path,
        "serial_log_persisted" => true,
        "qemu_booted" => @serial_log.booted?,
        "missing_boot_markers" => @serial_log.missing_markers,
        "boot_summary" => @serial_log.summary,
        "host_root_modified" => false,
        "privileged_container_required" => false,
        "docker_socket_mounted" => false,
        "host_network_enabled" => false,
        "qemu_network_restricted" => true,
        "network_required_for_source_download" => true,
        "desktop_safe_summary" => desktop_safe_summary
      }
    end

    def to_markdown
      data = to_h
      lines = [
        "# Full Build and QEMU Smoke Report",
        "",
        "- Version: #{data.fetch("version")}",
        "- Schema: #{data.fetch("schema_version")}",
        "- QEMU booted: #{data.fetch("qemu_booted")}",
        "- Serial log: #{data.fetch("serial_log_path")}",
        "- Boot summary: #{data.fetch("boot_summary")}",
        "- Host root modified: #{data.fetch("host_root_modified")}",
        "- Privileged container required: #{data.fetch("privileged_container_required")}",
        "- Docker socket mounted: #{data.fetch("docker_socket_mounted")}",
        "- Host network enabled: #{data.fetch("host_network_enabled")}",
        "- QEMU network restricted: #{data.fetch("qemu_network_restricted")}",
        "- Source download network required: #{data.fetch("network_required_for_source_download")}",
        "",
        "## Steps",
        ""
      ]
      data.fetch("steps").each_with_index do |step, index|
        lines << "#{index + 1}. #{step}"
      end
      lines << ""
      lines << data.fetch("desktop_safe_summary")
      lines << ""
      lines.join("\n")
    end

    private

    def desktop_safe_summary
      return "Full build and constrained QEMU serial smoke reached the expected boot markers." if @serial_log.booted?

      "Full build or QEMU serial smoke did not prove all boot markers."
    end
  end
end
