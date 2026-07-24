#!/usr/bin/env ruby
# frozen_string_literal: true

require_relative "../lib/xnix/full_smoke_report"
require_relative "../lib/xnix/serial_log"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

serial_log = Xnix::SerialLog.new("Welcome to Buildroot\nxnix login: ")
version = File.read(File.expand_path("../VERSION", __dir__), encoding: "UTF-8").strip
report = Xnix::FullSmokeReport.new(
  version: version,
  steps: %w[
    build-tools
    build
    fetch-sources
    configure-system
    download-system
    build-system
    prepare-ssh-test-key
    configure-wine-guest
    download-wine-guest
    build-ssh-wine-guest
    fetch-known-winapp
    boot-system
    staged-launcher-dispatch-smoke
    winapp-guest-wine-smoke
    kde-controlled-launch-action-dbus-fixture-smoke
  ],
  serial_log_path: "output/serial.log",
  serial_log: serial_log
)

data = report.to_h
assert(data.fetch("version") == version, "full smoke report must expose the canonical version")
assert(data.fetch("schema_version") == "xnix.full_smoke_report.v1", "full smoke report must expose its schema")
assert(data.fetch("report_type") == "full-build-qemu-smoke-report", "full smoke report must identify its type")
assert(data.fetch("step_count") == 15, "full smoke report must count smoke steps")
assert(data.fetch("steps").last == "kde-controlled-launch-action-dbus-fixture-smoke", "full smoke report must include the KDE action smoke step")
assert(data.fetch("serial_log_path") == "output/serial.log", "full smoke report must expose the serial log path")
assert(data.fetch("serial_log_persisted"), "full smoke report must record serial log persistence")
assert(data.fetch("qemu_booted"), "full smoke report must record successful QEMU boot markers")
assert(data.fetch("missing_boot_markers").empty?, "full smoke report must expose missing boot markers")
assert(!data.fetch("host_root_modified"), "full smoke report must not claim host root mutation")
assert(!data.fetch("privileged_container_required"), "full smoke report must not require privileged containers")
assert(!data.fetch("docker_socket_mounted"), "full smoke report must not mount the Docker socket")
assert(!data.fetch("host_network_enabled"), "full smoke report must keep host networking disabled")
assert(data.fetch("qemu_network_restricted"), "full smoke report must keep QEMU networking restricted")
assert(data.fetch("network_required_for_source_download"), "full smoke report must identify source download network use")
assert(data.fetch("wine_guest_built"), "full smoke report must record Wine guest build coverage")
assert(data.fetch("known_app_smoke_included"), "full smoke report must record known Windows app smoke coverage")
assert(data.fetch("known_app_smoke_passed"), "full smoke report must record known Windows app smoke pass evidence")
assert(data.fetch("fixture_app_smoke_included"), "full smoke report must record fixture Windows app smoke coverage")
assert(data.fetch("fixture_app_smoke_passed"), "full smoke report must record fixture Windows app smoke pass evidence")
assert(data.fetch("kde_action_smoke_included"), "full smoke report must record KDE action smoke coverage")
assert(data.fetch("kde_action_smoke_passed"), "full smoke report must record KDE action smoke pass evidence")
assert(data.fetch("failure_class") == "none", "full smoke report must classify successful runs as no failure")
assert(data.fetch("failed_step").nil?, "full smoke report must not report a failed step on success")
assert(!data.fetch("operator_action_required"), "full smoke report success must not require operator action")
assert(!data.fetch("project_defect_possible"), "full smoke report success must not claim a project defect")
assert(!data.fetch("full_smoke_failed"), "full smoke report success must not be marked failed")
assert(data.fetch("formal_release_ready"), "full smoke report success must be release ready")

markdown = report.to_markdown
assert(markdown.include?("# Full Build and QEMU Smoke Report"), "full smoke report markdown must include a title")
assert(markdown.include?("- QEMU booted: true"), "full smoke report markdown must include boot status")
assert(markdown.include?("- Known Windows app smoke passed: true"), "full smoke report markdown must include known app smoke status")
assert(markdown.include?("12. boot-system"), "full smoke report markdown must list the boot step")
assert(markdown.include?("14. winapp-guest-wine-smoke"), "full smoke report markdown must list the fixture app smoke step")
assert(markdown.include?("15. kde-controlled-launch-action-dbus-fixture-smoke"), "full smoke report markdown must list the KDE action smoke step")
assert(markdown.include?("- Failure class: none"), "full smoke report markdown must include failure class")
assert(markdown.include?("- Formal release ready: true"), "full smoke report markdown must include release readiness")

classification_cases = [
  [
    "build-tools",
    "failed to resolve reference \"docker.io/library/debian:bookworm-slim\": Head https://registry-1.docker.io/v2/library/debian/manifests/bookworm-slim: EOF",
    "docker-hub-eof",
    true,
    false
  ],
  [
    "build-tools",
    "Cannot connect to the Docker daemon at unix:///Users/example/.colima/default/docker.sock. Is the docker daemon running?",
    "docker-daemon-unavailable",
    true,
    false
  ],
  [
    "build-tools",
    "Error response from daemon: No such image: debian:bookworm-slim",
    "missing-local-base-image",
    true,
    false
  ],
  [
    "build-system",
    "make: *** [all] Error 2",
    "buildroot-build-failure",
    false,
    true
  ],
  [
    "boot-system",
    "Boot markers missing: Welcome to Buildroot, xnix login:",
    "qemu-serial-timeout",
    false,
    true
  ],
  [
    "staged-launcher-dispatch-smoke",
    "missing PASS: staged managed launcher dispatch smoke",
    "known-windows-app-smoke-failure",
    false,
    true
  ],
  [
    "winapp-guest-wine-smoke",
    "missing PASS: QEMU guest real Windows app Wine smoke",
    "fixture-windows-app-smoke-failure",
    false,
    true
  ],
  [
    "kde-controlled-launch-action-dbus-fixture-smoke",
    "missing PASS: KDE controlled launch action smoke",
    "kde-action-fixture-failure",
    false,
    true
  ]
]

classification_cases.each do |step, error_text, expected_class, operator_action, project_defect|
  failure = Xnix::FullSmokeReport.classify_failure(step: step, error_text: error_text)
  assert(failure.fetch("failure_class") == expected_class, "classifier must identify #{expected_class}")
  assert(failure.fetch("failed_step") == step, "classifier must retain the failed step for #{expected_class}")
  assert(failure.fetch("operator_action_required") == operator_action, "classifier must set operator action for #{expected_class}")
  assert(failure.fetch("project_defect_possible") == project_defect, "classifier must set project-defect possibility for #{expected_class}")
  assert(failure.fetch("safe_retry_command") == "ruby scripts/full_smoke.rb", "classifier must expose the safe retry command for #{expected_class}")
  assert(failure.fetch("original_error_text_retained"), "classifier must retain an error-text summary for #{expected_class}")
end

redacted_failure = Xnix::FullSmokeReport.classify_failure(
  step: "build-tools",
  error_text: "permission denied while opening /Users/rocky/.colima/default/docker.sock"
)
assert(redacted_failure.fetch("failure_summary").include?("[redacted-host-path]"), "classifier must redact host paths")
assert(!redacted_failure.fetch("failure_summary").include?("/Users/rocky"), "classifier must not expose raw host paths")

failed_report = Xnix::FullSmokeReport.new(
  version: version,
  steps: %w[
    build-tools
    build
    fetch-sources
    configure-system
    download-system
    build-system
    prepare-ssh-test-key
    configure-wine-guest
    download-wine-guest
    build-ssh-wine-guest
    fetch-known-winapp
    boot-system
    staged-launcher-dispatch-smoke
    winapp-guest-wine-smoke
    kde-controlled-launch-action-dbus-fixture-smoke
  ],
  serial_log_path: "output/serial.log",
  serial_log: serial_log,
  failure: Xnix::FullSmokeReport.classify_failure(
    step: "kde-controlled-launch-action-dbus-fixture-smoke",
    error_text: "missing PASS: KDE controlled launch action smoke"
  )
)
failed_data = failed_report.to_h
assert(failed_data.fetch("failure_class") == "kde-action-fixture-failure", "failed report must expose failure class")
assert(failed_data.fetch("full_smoke_failed"), "failed report must be marked failed")
assert(!failed_data.fetch("formal_release_ready"), "failed report must not be release ready")
assert(!failed_data.fetch("kde_action_smoke_passed"), "failed report must not mark the failed KDE action step passed")
assert(failed_report.to_markdown.include?("Full smoke did not pass"), "failed report markdown must explain the failure")

puts "PASS: full smoke report unit tests"
