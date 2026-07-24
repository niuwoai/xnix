#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "tempfile"
require "tmpdir"

PROJECT_ROOT = File.expand_path("..", __dir__)
SCRIPT = File.join(PROJECT_ROOT, "scripts", "full_checkpoint_promotion_packet.rb")

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def run_packet(*args)
  stdout, stderr, status = Open3.capture3("ruby", SCRIPT, *args, chdir: PROJECT_ROOT)
  assert(status.success?, "packet command failed: #{stderr}")
  JSON.parse(stdout)
end

def run_packet_markdown(*args)
  stdout, stderr, status = Open3.capture3("ruby", SCRIPT, "--format", "markdown", *args, chdir: PROJECT_ROOT)
  assert(status.success?, "packet markdown command failed: #{stderr}")
  stdout
end

def with_temp_report(report)
  Dir.mktmpdir("xnix-promotion-packet") do |dir|
    json_path = File.join(dir, "full-smoke-report.json")
    markdown_path = File.join(dir, "full-smoke-report.md")
    File.write(json_path, JSON.pretty_generate(report))
    File.write(markdown_path, "# Existing report\n")
    yield json_path, markdown_path
  end
end

def passing_report(version: "0.2.640-rc11")
  {
    "version" => version,
    "schema_version" => "xnix.full_smoke_report.v1",
    "report_type" => "full-build-qemu-smoke-report",
    "steps" => %w[
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
    "qemu_booted" => true,
    "wine_guest_built" => true,
    "known_app_smoke_passed" => true,
    "fixture_app_smoke_passed" => true,
    "kde_action_smoke_passed" => true,
    "failure_class" => "none",
    "failed_step" => nil,
    "operator_action_required" => false,
    "project_defect_possible" => false,
    "safe_retry_command" => "ruby scripts/full_smoke.rb",
    "failure_summary" => "No full-smoke failure was recorded.",
    "full_smoke_failed" => false,
    "formal_release_ready" => true
  }
end

missing_report = run_packet(
  "--report-json", "tmp/nonexistent-full-smoke-report.json",
  "--report-markdown", "tmp/nonexistent-full-smoke-report.md",
  "--current-version", "0.2.640-rc11"
)
assert(missing_report.fetch("schema_version") == "xnix.runtime.full_checkpoint_promotion_packet.v1", "packet must expose schema")
assert(missing_report.fetch("report_type") == "full-checkpoint-promotion-packet", "packet must identify report type")
assert(missing_report.fetch("full_smoke_state") == "missing-report", "missing report must be explicit")
assert(missing_report.fetch("promotion_decision") == "blocked-missing-full-smoke-report", "missing report must block promotion")
assert(!missing_report.fetch("promotion_allowed"), "missing report must not allow promotion")
assert(!missing_report.fetch("formal_release_ready"), "missing report must not claim release readiness")
assert(missing_report.fetch("operator_required_command") == "ruby scripts/full_smoke.rb", "missing report must name safe operator command")
assert(!missing_report.fetch("full_smoke_executed_by_packet"), "packet must not run full smoke")
assert(!missing_report.fetch("docker_executed_by_packet"), "packet must not run Docker")
assert(!missing_report.fetch("qemu_executed_by_packet"), "packet must not run QEMU")
assert(!missing_report.fetch("wine_executed_by_packet"), "packet must not run Wine")
assert(!missing_report.fetch("host_root_modified"), "packet must not mutate host root")

Dir.mktmpdir("xnix-promotion-packet-malformed") do |dir|
  json_path = File.join(dir, "full-smoke-report.json")
  File.write(json_path, "{not valid json")
  malformed = run_packet("--report-json", json_path, "--current-version", "0.2.640-rc11")
  assert(malformed.fetch("full_smoke_state") == "malformed-report", "malformed report must be explicit")
  assert(malformed.fetch("promotion_decision") == "blocked-malformed-full-smoke-report", "malformed report must block promotion")
  assert(!malformed.fetch("full_smoke_report_valid"), "malformed report must not be valid")
end

failed_report = passing_report.merge(
  "failure_class" => "kde-action-fixture-failure",
  "failed_step" => "kde-controlled-launch-action-dbus-fixture-smoke",
  "operator_action_required" => false,
  "project_defect_possible" => true,
  "safe_retry_command" => "ruby scripts/full_smoke.rb",
  "failure_summary" => "missing PASS under /Users/rocky/Sites/xnix/output/serial.log token=secret-value",
  "full_smoke_failed" => true,
  "formal_release_ready" => false,
  "kde_action_smoke_passed" => false
)
with_temp_report(failed_report) do |json_path, markdown_path|
  failed = run_packet(
    "--report-json", json_path,
    "--report-markdown", markdown_path,
    "--current-version", "0.2.640-rc11"
  )
  assert(failed.fetch("full_smoke_state") == "failed", "failed report must stay failed")
  assert(failed.fetch("promotion_decision") == "blocked-full-smoke-failed", "failed report must block promotion")
  assert(failed.fetch("failure_class") == "kde-action-fixture-failure", "failed report must preserve failure class")
  assert(failed.fetch("failed_step") == "kde-controlled-launch-action-dbus-fixture-smoke", "failed report must preserve failed step")
  assert(failed.fetch("project_defect_possible"), "failed report must preserve project defect flag")
  assert(failed.fetch("safe_retry_command") == "ruby scripts/full_smoke.rb", "failed report must preserve safe retry command")
  assert(!failed.fetch("failure_summary").include?("/Users/rocky"), "packet must redact host paths")
  assert(!failed.fetch("failure_summary").include?("token=secret-value"), "packet must redact sensitive values")
end

incomplete_report = passing_report.merge(
  "failure_class" => nil,
  "full_smoke_failed" => false,
  "formal_release_ready" => false,
  "kde_action_smoke_passed" => false
)
with_temp_report(incomplete_report) do |json_path, markdown_path|
  incomplete = run_packet(
    "--report-json", json_path,
    "--report-markdown", markdown_path,
    "--current-version", "0.2.640-rc11"
  )
  assert(incomplete.fetch("full_smoke_state") == "blocked-incomplete-report", "incomplete report must not be classified as failed")
  assert(incomplete.fetch("promotion_decision") == "blocked-incomplete-full-smoke-report", "incomplete report must block promotion")
  assert(incomplete.fetch("failure_class") == "none", "missing failure class must be rendered as none")
end

with_temp_report(passing_report) do |json_path, markdown_path|
  passed = run_packet(
    "--report-json", json_path,
    "--report-markdown", markdown_path,
    "--current-version", "0.2.640-rc11"
  )
  assert(passed.fetch("full_smoke_state") == "passed", "passing report must be recognized")
  assert(passed.fetch("promotion_decision") == "promote", "matching passing report must allow promotion")
  assert(passed.fetch("promotion_allowed"), "matching passing report must allow promotion")
  assert(passed.fetch("formal_release_ready"), "matching passing report must claim formal readiness")
  assert(passed.fetch("restricted_container_evidence_observed"), "passing report must observe restricted container evidence")
  assert(passed.fetch("qemu_evidence_observed"), "passing report must observe QEMU evidence")
  assert(passed.fetch("wine_evidence_observed"), "passing report must observe Wine evidence")
  assert(passed.fetch("known_windows_app_smoke_observed"), "passing report must observe known app evidence")
  assert(passed.fetch("fixture_windows_app_smoke_observed"), "passing report must observe fixture app evidence")
  assert(passed.fetch("kde_action_smoke_observed"), "passing report must observe KDE action evidence")
  assert(!passed.fetch("backend_launch_enabled"), "packet must keep backend launch disabled")
end

with_temp_report(passing_report) do |json_path, markdown_path|
  mismatch = run_packet(
    "--report-json", json_path,
    "--report-markdown", markdown_path,
    "--current-version", "0.2.641-rc1"
  )
  assert(mismatch.fetch("full_smoke_state") == "passed", "version mismatch still sees the report pass")
  assert(mismatch.fetch("promotion_decision") == "blocked-version-mismatch", "version mismatch must block promotion")
  assert(!mismatch.fetch("promotion_allowed"), "version mismatch must not allow promotion")
  assert(!mismatch.fetch("formal_release_ready"), "version mismatch must not claim formal readiness")
end

with_temp_report(passing_report) do |json_path, markdown_path|
  markdown = run_packet_markdown(
    "--report-json", json_path,
    "--report-markdown", markdown_path,
    "--current-version", "0.2.640-rc11"
  )
  assert(markdown.include?("# Full Checkpoint Promotion Packet"), "markdown must include title")
  assert(markdown.include?("- Promotion decision: promote"), "markdown must include decision")
  assert(markdown.include?("- Docker executed: false"), "markdown must include skipped Docker operation")
  assert(markdown.include?("- QEMU executed: false"), "markdown must include skipped QEMU operation")
  assert(markdown.include?("- Wine executed: false"), "markdown must include skipped Wine operation")
end

puts "PASS: full checkpoint promotion packet tests"
