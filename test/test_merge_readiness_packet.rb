#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tempfile"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

def write_json_fixture(payload)
  file = Tempfile.new("xnix-merge-readiness-packet")
  file.write(JSON.pretty_generate(payload))
  file.flush
  file
end

def write_text_fixture(text)
  file = Tempfile.new("xnix-merge-readiness-packet")
  file.write(text)
  file.flush
  file
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/merge_readiness_packet.rb")

layout = write_text_fixture("PASS: Xnix fixture scaffold is consistent\n")
implementation = write_json_fixture(
  "domains" => [
    { "id" => "developer-verification-harness", "status" => "fixture-implemented" }
  ]
)
contract_drift = write_json_fixture("drift_detected" => false)
kde_smoke = write_json_fixture("entrypoint_count" => 7)
mainline = write_json_fixture(
  "changed_file_count" => 3,
  "unclassified_file_count" => 0,
  "protected_claude_file_modified" => false,
  "risky_file_count" => 0,
  "code_change_count" => 2,
  "lanes" => [
    {
      "id" => "cw10-evidence-drift-harness",
      "workstream" => "CW10 / A8",
      "title" => "Evidence and drift harness",
      "file_count" => 2,
      "statuses" => [" M"],
      "required_verification" => [
        "ruby scripts/merge_readiness_packet.rb --format json",
        "ruby scripts/verify_layout.rb"
      ],
      "safety_guards" => ["read-only-report"]
    },
    {
      "id" => "unclassified",
      "workstream" => "unknown",
      "title" => "Unclassified",
      "file_count" => 0,
      "statuses" => [],
      "required_verification" => [],
      "safety_guards" => ["manual-classification-required"]
    }
  ]
)
release_evidence = write_json_fixture("claims" => [], "claim_count" => 0)
fixture_matrix = write_json_fixture("row_count" => 7, "counts" => { "blocked" => 1 })
malformed = write_text_fixture("{not-json")
protected_mainline = write_json_fixture(
  "changed_file_count" => 2,
  "unclassified_file_count" => 1,
  "protected_claude_file_modified" => true,
  "risky_file_count" => 0,
  "code_change_count" => 1,
  "lanes" => []
)
unsafe_kde = write_json_fixture("entrypoint_count" => 7, "docker_executed" => true)

base_args = [
  "--layout-report", layout.path,
  "--implementation-report", implementation.path,
  "--contract-drift-report", contract_drift.path,
  "--kde-smoke-report", kde_smoke.path,
  "--mainline-review", mainline.path,
  "--release-evidence-index", release_evidence.path,
  "--fixture-matrix-report", fixture_matrix.path
]

begin
  stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json", *base_args)
  assert(status.success?, "merge readiness packet JSON must exit successfully: #{stderr}")
  packet = JSON.parse(stdout)

  assert(packet.fetch("version") == File.read(project_root.join("VERSION")).strip, "packet must expose current version")
  assert(packet.fetch("schema_version") == "xnix.runtime.merge_readiness_packet.v1", "packet must expose schema")
  assert(packet.fetch("report_type") == "merge-readiness-packet", "packet must identify report type")
  assert(packet.fetch("runtime_owned"), "packet must keep Runtime ownership explicit")
  assert(packet.fetch("ruby_report_only"), "packet must identify itself as Ruby report-only tooling")
  assert(!packet.fetch("kde_policy_owner"), "packet must not make KDE the policy owner")
  assert(packet.fetch("offline_only"), "packet must default to offline-only")
  assert(packet.fetch("merge_ready"), "packet must be merge-ready when only restricted smoke remains")
  assert(!packet.fetch("release_ready"), "packet must keep release readiness gated")
  assert(packet.fetch("release_blocking_reasons").include?("restricted-docker-or-qemu-smoke-requires-human-authorization"), "packet must keep restricted Docker/QEMU as release blocker")
  assert(packet.fetch("merge_blocking_reasons").empty?, "packet must not merge-block the clean fixture case")

  tool_statuses = packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(tool_statuses.values.all? { |tool| tool.fetch("status") == "pass" }, "all fixture-backed tools must pass")
  assert(tool_statuses.fetch("layout").fetch("command") == "fixture:layout", "fixture command must not expose fixture path")
  assert(packet.fetch("changed_file_counts").fetch("total") == 3, "packet must include changed file counts")
  assert(packet.fetch("lane_classification").any? { |lane| lane.fetch("id") == "cw10-evidence-drift-harness" }, "packet must include lane classification")
  assert(packet.fetch("required_follow_up_commands").include?("ruby scripts/verify_layout.rb"), "packet must include follow-up commands")
  unsafe_status = packet.fetch("unsafe_operation_status")
  assert(!unsafe_status.fetch("docker_executed"), "packet itself must not run Docker")
  assert(!unsafe_status.fetch("qemu_executed"), "packet itself must not run QEMU")
  assert(!unsafe_status.fetch("network_checks_run"), "packet itself must not run network checks")
  assert(!unsafe_status.fetch("package_manager_invoked"), "packet itself must not invoke package managers")
  assert(!unsafe_status.fetch("host_root_modified"), "packet itself must not mutate host root")
  assert(unsafe_status.fetch("detected_unsafe_findings").empty?, "clean packet must not report unsafe findings")

  markdown, markdown_stderr, markdown_status = Open3.capture3("ruby", script.to_s, "--format", "markdown", *base_args)
  assert(markdown_status.success?, "merge readiness packet Markdown must exit successfully: #{markdown_stderr}")
  assert(markdown.include?("# Merge Readiness Packet"), "Markdown must include title")
  assert(markdown.include?("cw10-evidence-drift-harness"), "Markdown must include lane classification")
  assert(markdown.include?("restricted-docker-or-qemu-smoke-requires-human-authorization"), "Markdown must include release blocker")

  skipped_stdout, skipped_stderr, skipped_status = Open3.capture3(
    "ruby", script.to_s, "--format", "json", "--skip-tool", "implementation", *base_args
  )
  assert(skipped_status.success?, "skipped tool case must still emit packet: #{skipped_stderr}")
  skipped_packet = JSON.parse(skipped_stdout)
  skipped_tools = skipped_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(skipped_tools.fetch("implementation").fetch("status") == "skipped", "packet must record skipped tools")
  assert(skipped_packet.fetch("merge_blocking_reasons").include?("implementation:skipped"), "required skipped tool must block merge")

  missing_stdout, missing_stderr, missing_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--tool-command",
    "implementation=definitely_missing_xnix_packet_tool",
    *base_args.reject.with_index { |_value, index| [2, 3].include?(index) }
  )
  assert(missing_status.success?, "missing command case must still emit packet: #{missing_stderr}")
  missing_packet = JSON.parse(missing_stdout)
  missing_tools = missing_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(missing_tools.fetch("implementation").fetch("status") == "missing-command", "packet must detect missing commands")
  assert(missing_packet.fetch("merge_blocking_reasons").include?("implementation:missing-command"), "missing required command must block merge")

  malformed_stdout, malformed_stderr, malformed_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    malformed.path,
    "--layout-report",
    layout.path,
    "--contract-drift-report",
    contract_drift.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--mainline-review",
    mainline.path
  )
  assert(malformed_status.success?, "malformed JSON case must still emit packet: #{malformed_stderr}")
  malformed_packet = JSON.parse(malformed_stdout)
  malformed_tools = malformed_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(malformed_tools.fetch("implementation").fetch("status") == "malformed-json", "packet must detect malformed JSON")
  assert(malformed_packet.fetch("merge_blocking_reasons").include?("implementation:malformed-json"), "malformed required JSON must block merge")

  protected_stdout, protected_stderr, protected_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--mainline-review",
    protected_mainline.path,
    "--layout-report",
    layout.path,
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--kde-smoke-report",
    kde_smoke.path
  )
  assert(protected_status.success?, "protected file case must still emit packet: #{protected_stderr}")
  protected_packet = JSON.parse(protected_stdout)
  assert(protected_packet.fetch("protected_file_status").fetch("status") == "blocked", "protected file change must block")
  assert(protected_packet.fetch("merge_blocking_reasons").include?("protected-claude-file-modified"), "protected file blocker must be explicit")
  assert(protected_packet.fetch("merge_blocking_reasons").include?("unclassified-files-present"), "unclassified file blocker must be explicit")

  unsafe_stdout, unsafe_stderr, unsafe_status_result = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--kde-smoke-report",
    unsafe_kde.path,
    "--layout-report",
    layout.path,
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path
  )
  assert(unsafe_status_result.success?, "unsafe flag case must still emit packet: #{unsafe_stderr}")
  unsafe_packet = JSON.parse(unsafe_stdout)
  findings = unsafe_packet.fetch("unsafe_operation_status").fetch("detected_unsafe_findings")
  assert(findings.any? { |finding| finding.fetch("key") == "docker_executed" }, "packet must report unsafe-operation findings")
  assert(unsafe_packet.fetch("merge_blocking_reasons").include?("unsafe-operation-detected"), "unsafe operation must block merge")
ensure
  [
    layout,
    implementation,
    contract_drift,
    kde_smoke,
    mainline,
    release_evidence,
    fixture_matrix,
    malformed,
    protected_mainline,
    unsafe_kde
  ].compact.each do |file|
    file.close
    file.unlink
  end
end
