# frozen_string_literal: true

module Xnix
  class FullSmokeReport
    SCHEMA_VERSION = "xnix.full_smoke_report.v1"
    REPORT_TYPE = "full-build-qemu-smoke-report"
    NO_FAILURE = {
      "failure_class" => "none",
      "failed_step" => nil,
      "operator_action_required" => false,
      "project_defect_possible" => false,
      "safe_retry_command" => "ruby scripts/full_smoke.rb",
      "failure_summary" => "No full-smoke failure was recorded.",
      "original_error_text_retained" => false
    }.freeze

    def self.classify_failure(step:, error_text:)
      text = error_text.to_s
      failure_class = if docker_hub_eof?(text)
                        "docker-hub-eof"
                      elsif docker_daemon_unavailable?(text)
                        "docker-daemon-unavailable"
                      elsif missing_local_base_image?(text)
                        "missing-local-base-image"
                      elsif step == "boot-system" || text.include?("Boot markers missing") || text.include?("QEMU boot process failed")
                        "qemu-serial-timeout"
                      elsif step == "staged-launcher-dispatch-smoke"
                        "known-windows-app-smoke-failure"
                      elsif step == "winapp-guest-wine-smoke"
                        "fixture-windows-app-smoke-failure"
                      elsif step == "kde-controlled-launch-action-dbus-fixture-smoke"
                        "kde-action-fixture-failure"
                      elsif %w[build-system build-ssh-wine-guest configure-system configure-wine-guest download-system download-wine-guest].include?(step)
                        "buildroot-build-failure"
                      else
                        "full-smoke-step-failure"
                      end

      {
        "failure_class" => failure_class,
        "failed_step" => step,
        "operator_action_required" => operator_action_required?(failure_class),
        "project_defect_possible" => project_defect_possible?(failure_class),
        "safe_retry_command" => "ruby scripts/full_smoke.rb",
        "failure_summary" => sanitize_error_text(text),
        "original_error_text_retained" => !text.empty?
      }
    end

    def self.docker_hub_eof?(text)
      text.include?("EOF") && (
        text.include?("registry-1.docker.io") ||
        text.include?("auth.docker.io") ||
        text.include?("production.cloudflare.docker.com") ||
        text.include?("production.cloudfront.docker.com") ||
        text.include?("debian:bookworm-slim")
      )
    end

    def self.docker_daemon_unavailable?(text)
      text.include?("Cannot connect to the Docker daemon") ||
        text.include?("docker.sock") ||
        (text.include?("permission denied") && text.include?("docker"))
    end

    def self.missing_local_base_image?(text)
      text.include?("No such image") ||
        text.include?("pull access denied") ||
        text.include?("not found: manifest unknown") ||
        text.include?("not found")
    end

    def self.operator_action_required?(failure_class)
      %w[
        docker-hub-eof
        docker-daemon-unavailable
        missing-local-base-image
      ].include?(failure_class)
    end

    def self.project_defect_possible?(failure_class)
      !operator_action_required?(failure_class)
    end

    def self.sanitize_error_text(text)
      sanitized = text.to_s
                      .gsub(%r{/Users/[^\s:'"]+}, "[redacted-host-path]")
                      .gsub(%r{/private/(?:tmp|var)/[^\s:'"]+}, "[redacted-host-path]")
                      .gsub(%r{/var/folders/[^\s:'"]+}, "[redacted-host-path]")
      sanitized.lines.first(24).join.strip
    end

    def initialize(version:, steps:, serial_log_path:, serial_log:, failure: nil)
      @version = version
      @steps = steps
      @serial_log_path = serial_log_path
      @serial_log = serial_log
      @failure = failure || NO_FAILURE
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
        "wine_guest_built" => wine_guest_built?,
        "known_app_smoke_included" => known_app_smoke_included?,
        "known_app_smoke_passed" => smoke_passed?("staged-launcher-dispatch-smoke"),
        "fixture_app_smoke_included" => fixture_app_smoke_included?,
        "fixture_app_smoke_passed" => smoke_passed?("winapp-guest-wine-smoke"),
        "kde_action_smoke_included" => kde_action_smoke_included?,
        "kde_action_smoke_passed" => smoke_passed?("kde-controlled-launch-action-dbus-fixture-smoke"),
        "failure_class" => @failure.fetch("failure_class"),
        "failed_step" => @failure.fetch("failed_step"),
        "operator_action_required" => @failure.fetch("operator_action_required"),
        "project_defect_possible" => @failure.fetch("project_defect_possible"),
        "safe_retry_command" => @failure.fetch("safe_retry_command"),
        "failure_summary" => @failure.fetch("failure_summary"),
        "original_error_text_retained" => @failure.fetch("original_error_text_retained"),
        "full_smoke_failed" => failed?,
        "formal_release_ready" => formal_release_ready?,
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
        "- Wine guest built: #{data.fetch("wine_guest_built")}",
        "- Known Windows app smoke included: #{data.fetch("known_app_smoke_included")}",
        "- Known Windows app smoke passed: #{data.fetch("known_app_smoke_passed")}",
        "- Fixture Windows app smoke included: #{data.fetch("fixture_app_smoke_included")}",
        "- Fixture Windows app smoke passed: #{data.fetch("fixture_app_smoke_passed")}",
        "- KDE action smoke included: #{data.fetch("kde_action_smoke_included")}",
        "- KDE action smoke passed: #{data.fetch("kde_action_smoke_passed")}",
        "- Failure class: #{data.fetch("failure_class")}",
        "- Failed step: #{data.fetch("failed_step") || "none"}",
        "- Operator action required: #{data.fetch("operator_action_required")}",
        "- Project defect possible: #{data.fetch("project_defect_possible")}",
        "- Safe retry command: `#{data.fetch("safe_retry_command")}`",
        "- Formal release ready: #{data.fetch("formal_release_ready")}",
        "",
        "## Steps",
        ""
      ]
      data.fetch("steps").each_with_index do |step, index|
        lines << "#{index + 1}. #{step}"
      end
      lines << ""
      lines << "## Failure Summary"
      lines << ""
      lines << data.fetch("failure_summary")
      lines << ""
      lines << data.fetch("desktop_safe_summary")
      lines << ""
      lines.join("\n")
    end

    private

    def wine_guest_built?
      @steps.include?("build-ssh-wine-guest")
    end

    def known_app_smoke_included?
      @steps.include?("staged-launcher-dispatch-smoke")
    end

    def fixture_app_smoke_included?
      @steps.include?("winapp-guest-wine-smoke")
    end

    def kde_action_smoke_included?
      @steps.include?("kde-controlled-launch-action-dbus-fixture-smoke")
    end

    def failed?
      @failure.fetch("failure_class") != "none"
    end

    def failed_step
      @failure.fetch("failed_step")
    end

    def smoke_passed?(step)
      @steps.include?(step) && failed_step != step
    end

    def formal_release_ready?
      !failed? && @serial_log.booted? && known_app_smoke_included? && fixture_app_smoke_included? && kde_action_smoke_included?
    end

    def desktop_safe_summary
      if failed?
        return "Full smoke did not pass; the failure was classified as #{@failure.fetch("failure_class")} at #{@failure.fetch("failed_step")}."
      end

      if formal_release_ready?
        return "Full build, constrained QEMU serial smoke, Wine guest, real Windows app smokes, and KDE controlled-launch action smoke reached the expected markers."
      end

      return "Full build and constrained QEMU serial smoke reached the expected boot markers." if @serial_log.booted?

      "Full build or QEMU serial smoke did not prove all boot markers."
    end
  end
end
