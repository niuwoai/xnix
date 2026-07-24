#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
DEFAULT_REPORT_JSON = PROJECT_ROOT.join("output", "full-smoke-report.json")
DEFAULT_REPORT_MARKDOWN = PROJECT_ROOT.join("output", "full-smoke-report.md")
DEFAULT_EXPECTED_RELEASE = "0.2.640"
SCHEMA_VERSION = "xnix.runtime.full_checkpoint_promotion_packet.v1"
REPORT_TYPE = "full-checkpoint-promotion-packet"
FULL_SMOKE_SCHEMA = "xnix.full_smoke_report.v1"
FULL_SMOKE_REPORT_TYPE = "full-build-qemu-smoke-report"
SAFE_RETRY_COMMAND = "ruby scripts/full_smoke.rb"

NO_EXECUTION_FLAGS = {
  "full_smoke_executed_by_packet" => false,
  "docker_executed_by_packet" => false,
  "qemu_executed_by_packet" => false,
  "wine_executed_by_packet" => false,
  "colima_executed_by_packet" => false,
  "dbus_called_by_packet" => false,
  "desktop_launch_enabled" => false,
  "backend_launch_enabled" => false,
  "network_required_by_packet" => false,
  "privileged_container_required_by_packet" => false,
  "host_root_modified" => false
}.freeze

def parse_options
  options = {
    format: "json",
    expected_release: DEFAULT_EXPECTED_RELEASE,
    current_version: nil,
    report_json: DEFAULT_REPORT_JSON,
    report_markdown: DEFAULT_REPORT_MARKDOWN
  }

  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/full_checkpoint_promotion_packet.rb [--format json|markdown] [options]"
    parser.on("--format FORMAT", "Output format: json or markdown") { |value| options[:format] = value }
    parser.on("--expected-release VERSION", "Expected formal release version") { |value| options[:expected_release] = value }
    parser.on("--current-version VERSION", "Override repository version for tests") { |value| options[:current_version] = value }
    parser.on("--report-json PATH", "Existing full-smoke JSON report path") { |value| options[:report_json] = Pathname.new(value) }
    parser.on("--report-markdown PATH", "Existing full-smoke Markdown report path") { |value| options[:report_markdown] = Pathname.new(value) }
  end.parse!

  unless %w[json markdown].include?(options[:format])
    abort "full_checkpoint_promotion_packet supports --format json or --format markdown"
  end

  options
end

def current_version(options)
  return options[:current_version].to_s.strip if options[:current_version]

  PROJECT_ROOT.join("VERSION").read.strip
end

def promotable_version?(version, expected_release)
  version == expected_release || version.match?(/\A#{Regexp.escape(expected_release)}-rc\d+\z/)
end

def release_candidate?(version, expected_release)
  version.match?(/\A#{Regexp.escape(expected_release)}-rc\d+\z/)
end

def display_path(path, missing_label)
  expanded = path.expand_path
  root = PROJECT_ROOT.to_s
  expanded_s = expanded.to_s
  return expanded.relative_path_from(PROJECT_ROOT).to_s if expanded_s == root || expanded_s.start_with?("#{root}/")

  path.basename.to_s.empty? ? missing_label : path.basename.to_s
rescue StandardError
  missing_label
end

def sanitize_text(value)
  value.to_s
       .gsub(%r{/Users/[^\s:'"]+}, "[redacted-host-path]")
       .gsub(%r{/private/(?:tmp|var)/[^\s:'"]+}, "[redacted-host-path]")
       .gsub(%r{/var/folders/[^\s:'"]+}, "[redacted-host-path]")
       .gsub(/(?:token|secret|password|credential)=\S+/i, "[redacted-secret]")
       .lines
       .first(24)
       .join
       .strip
end

def load_full_smoke_report(path)
  return [nil, "missing-report", "No full-smoke JSON report exists."] unless path.exist?

  parsed = JSON.parse(path.read)
  [parsed, nil, nil]
rescue JSON::ParserError => e
  [nil, "malformed-report", sanitize_text(e.message)]
end

def valid_full_smoke_report?(report)
  report.is_a?(Hash) &&
    report["schema_version"] == FULL_SMOKE_SCHEMA &&
    report["report_type"] == FULL_SMOKE_REPORT_TYPE
end

def full_smoke_state(report, report_error)
  return report_error if report_error
  return "malformed-report" unless valid_full_smoke_report?(report)
  return "passed" if report["formal_release_ready"] == true && report["full_smoke_failed"] == false

  failure_class = report["failure_class"].to_s
  return "failed" if report["full_smoke_failed"] == true || (!failure_class.empty? && failure_class != "none")

  "blocked-incomplete-report"
end

def promotion_decision(full_smoke_state, version_promotable)
  return "promote" if full_smoke_state == "passed" && version_promotable
  return "blocked-version-mismatch" if full_smoke_state == "passed" && !version_promotable

  case full_smoke_state
  when "missing-report"
    "blocked-missing-full-smoke-report"
  when "malformed-report"
    "blocked-malformed-full-smoke-report"
  when "failed"
    "blocked-full-smoke-failed"
  when "blocked-incomplete-report"
    "blocked-incomplete-full-smoke-report"
  else
    "blocked-full-smoke-evidence"
  end
end

def observed_evidence(report)
  steps = Array(report && report["steps"])
  {
    "restricted_container_evidence_observed" => steps.include?("build-tools") || steps.include?("build"),
    "qemu_evidence_observed" => report && (report["qemu_booted"] == true || steps.include?("boot-system")),
    "wine_evidence_observed" => report && (report["wine_guest_built"] == true || steps.include?("winapp-guest-wine-smoke")),
    "known_windows_app_smoke_observed" => report && (report["known_app_smoke_passed"] == true || steps.include?("staged-launcher-dispatch-smoke")),
    "fixture_windows_app_smoke_observed" => report && (report["fixture_app_smoke_passed"] == true || steps.include?("winapp-guest-wine-smoke")),
    "kde_action_smoke_observed" => report && (report["kde_action_smoke_passed"] == true || steps.include?("kde-controlled-launch-action-dbus-fixture-smoke"))
  }.transform_values { |value| value == true }
end

def packet_for(options)
  version = current_version(options)
  expected_release = options[:expected_release]
  report_json = options[:report_json]
  report_markdown = options[:report_markdown]
  report, report_error, report_error_detail = load_full_smoke_report(report_json)
  state = full_smoke_state(report, report_error)
  version_promotable = promotable_version?(version, expected_release)
  decision = promotion_decision(state, version_promotable)
  promotion_allowed = decision == "promote"

  failure_summary = if report_error_detail
                      report_error_detail
                    elsif report
                      sanitize_text(report["failure_summary"])
                    else
                      ""
                    end

  {
    "schema_version" => SCHEMA_VERSION,
    "report_type" => REPORT_TYPE,
    "current_version" => version,
    "expected_release_version" => expected_release,
    "release_candidate" => release_candidate?(version, expected_release),
    "version_promotable" => version_promotable,
    "full_smoke_report_json_path" => display_path(report_json, "full-smoke-report.json"),
    "full_smoke_report_markdown_path" => display_path(report_markdown, "full-smoke-report.md"),
    "full_smoke_report_present" => report_json.exist?,
    "full_smoke_markdown_present" => report_markdown.exist?,
    "full_smoke_report_valid" => valid_full_smoke_report?(report),
    "full_smoke_state" => state,
    "full_smoke_failed" => report ? report["full_smoke_failed"] == true : nil,
    "failure_class" => report ? sanitize_text(report["failure_class"].to_s.empty? ? "none" : report["failure_class"]) : state,
    "failed_step" => report ? sanitize_text(report["failed_step"]) : nil,
    "operator_action_required" => report ? report["operator_action_required"] == true : state == "missing-report",
    "project_defect_possible" => report ? report["project_defect_possible"] == true : false,
    "safe_retry_command" => report ? sanitize_text(report["safe_retry_command"] || SAFE_RETRY_COMMAND) : SAFE_RETRY_COMMAND,
    "failure_summary" => failure_summary,
    "formal_release_ready" => promotion_allowed,
    "promotion_allowed" => promotion_allowed,
    "promotion_decision" => decision,
    "operator_required_command" => promotion_allowed ? nil : SAFE_RETRY_COMMAND,
    "next_operator_action" => promotion_allowed ? "Promote the release candidate with the human-owned release workflow." : "Run the human-authorized full smoke and review the generated report before promotion.",
    "desktop_safe_summary" => promotion_allowed ? "Existing full-smoke evidence allows v#{expected_release} promotion review." : "v#{expected_release} promotion is blocked until existing full-smoke PASS evidence is present and valid."
  }.merge(observed_evidence(report)).merge(NO_EXECUTION_FLAGS)
end

def markdown_for(packet)
  lines = [
    "# Full Checkpoint Promotion Packet",
    "",
    "- Schema: #{packet.fetch("schema_version")}",
    "- Current version: #{packet.fetch("current_version")}",
    "- Expected release: #{packet.fetch("expected_release_version")}",
    "- Full-smoke JSON report: #{packet.fetch("full_smoke_report_json_path")}",
    "- Full-smoke report present: #{packet.fetch("full_smoke_report_present")}",
    "- Full-smoke report valid: #{packet.fetch("full_smoke_report_valid")}",
    "- Full-smoke state: #{packet.fetch("full_smoke_state")}",
    "- Failure class: #{packet.fetch("failure_class")}",
    "- Failed step: #{packet.fetch("failed_step") || "none"}",
    "- Safe retry command: `#{packet.fetch("safe_retry_command")}`",
    "- Promotion decision: #{packet.fetch("promotion_decision")}",
    "- Promotion allowed: #{packet.fetch("promotion_allowed")}",
    "- Formal release ready: #{packet.fetch("formal_release_ready")}",
    "",
    "## Observed Existing Evidence",
    "",
    "- Restricted container evidence observed: #{packet.fetch("restricted_container_evidence_observed")}",
    "- QEMU evidence observed: #{packet.fetch("qemu_evidence_observed")}",
    "- Wine evidence observed: #{packet.fetch("wine_evidence_observed")}",
    "- Known Windows app smoke observed: #{packet.fetch("known_windows_app_smoke_observed")}",
    "- Fixture Windows app smoke observed: #{packet.fetch("fixture_windows_app_smoke_observed")}",
    "- KDE action smoke observed: #{packet.fetch("kde_action_smoke_observed")}",
    "",
    "## Operations Not Executed By This Packet",
    "",
    "- Full smoke executed: #{packet.fetch("full_smoke_executed_by_packet")}",
    "- Docker executed: #{packet.fetch("docker_executed_by_packet")}",
    "- QEMU executed: #{packet.fetch("qemu_executed_by_packet")}",
    "- Wine executed: #{packet.fetch("wine_executed_by_packet")}",
    "- Colima executed: #{packet.fetch("colima_executed_by_packet")}",
    "- D-Bus called: #{packet.fetch("dbus_called_by_packet")}",
    "- Backend launch enabled: #{packet.fetch("backend_launch_enabled")}",
    "- Host root modified: #{packet.fetch("host_root_modified")}",
    "",
    "## Next Operator Action",
    "",
    packet.fetch("next_operator_action"),
    ""
  ]
  lines.join("\n")
end

if $PROGRAM_NAME == __FILE__
  options = parse_options
  packet = packet_for(options)
  if options[:format] == "json"
    puts JSON.pretty_generate(packet)
  else
    puts markdown_for(packet)
  end
end
