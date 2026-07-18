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
  file = Tempfile.new("xnix-release-evidence-index")
  file.write(JSON.pretty_generate(payload))
  file.flush
  file
end

def write_text_fixture(text)
  file = Tempfile.new("xnix-release-evidence-index")
  file.write(text)
  file.flush
  file
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/release_evidence_index.rb")

implementation = write_json_fixture(
  "domains" => [
    {
      "id" => "runtime-owner-service",
      "status" => "smoke-owned",
      "summary" => "Go owner read boundary exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "recipe-artifact-trust-pipeline",
      "status" => "state-root-implemented",
      "summary" => "Recipe and artifact evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "environment-lifecycle-state",
      "status" => "fixture-implemented",
      "summary" => "Backend lifecycle fixture evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "portal-snapshot-control-plane",
      "status" => "contract-only",
      "summary" => "Portal and snapshot contracts exist.",
      "contract_files_missing" => [],
      "fixture_files_missing" => ["internal/runtime/portal/ledger.go"]
    },
    {
      "id" => "kde-activation-shell-materialization",
      "status" => "smoke-owned",
      "summary" => "KDE seven entry point smoke exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "execution-transaction-ledger",
      "status" => "state-root-implemented",
      "summary" => "Execution ledger exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "diagnostics-repair-ai-boundary",
      "status" => "missing",
      "summary" => "Diagnostics evidence is missing.",
      "contract_files_missing" => ["internal/runtime/appidentity/ai_diagnostics.go"],
      "fixture_files_missing" => []
    },
    {
      "id" => "developer-verification-harness",
      "status" => "fixture-implemented",
      "summary" => "Report fixture evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "atomic-kde-image-qemu-acceptance",
      "status" => "smoke-owned",
      "summary" => "Image smoke scaffolding exists but needs explicit authorization.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    }
  ]
)
contract_drift = write_json_fixture("drift_detected" => false)
mainline = write_json_fixture(
  "protected_claude_file_modified" => true,
  "unclassified_file_count" => 2
)
kde_smoke = write_json_fixture("entrypoint_count" => 7)
malformed = nil

begin
  stdout, stderr, status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path
  )
  assert(status.success?, "release evidence index JSON must exit successfully: #{stderr}")
  report = JSON.parse(stdout)

  assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "release evidence index must expose the current version")
  assert(report.fetch("schema_version") == "xnix.runtime.release_evidence_index.v1", "release evidence index must expose the schema")
  assert(report.fetch("report_type") == "release-evidence-index", "release evidence index must identify its report type")
  assert(report.fetch("runtime_owned"), "release evidence index must keep Runtime ownership explicit")
  assert(!report.fetch("go_runtime_backed"), "release evidence index must identify itself as a Ruby report, not Go Runtime business logic")
  assert(report.fetch("ruby_report_only"), "release evidence index must be marked Ruby report only")
  assert(!report.fetch("kde_policy_owner"), "release evidence index must not make KDE the policy owner")
  assert(report.fetch("offline_default"), "release evidence index must run in offline default mode")
  assert(!report.fetch("docker_executed"), "release evidence index must not run Docker")
  assert(!report.fetch("qemu_executed"), "release evidence index must not run QEMU")
  assert(!report.fetch("network_checks_run"), "release evidence index must not run network checks")
  assert(!report.fetch("package_manager_invoked"), "release evidence index must not invoke package managers")
  assert(!report.fetch("backend_launch_enabled"), "release evidence index must not launch backends")
  assert(!report.fetch("host_root_modified"), "release evidence index must not mutate the host root")
  assert(!report.fetch("automatic_staging_enabled"), "release evidence index must not stage automatically")
  assert(!report.fetch("automatic_release_tagging_enabled"), "release evidence index must not tag releases automatically")
  assert(!report.fetch("release_ready"), "release evidence index must keep release readiness gated")
  assert(report.fetch("human_authorization_required"), "release evidence index must flag human authorization requirements")

  counts = report.fetch("counts")
  assert(counts.fetch("implemented").positive?, "release evidence index must include implemented evidence")
  assert(counts.fetch("fixture_only").positive?, "release evidence index must include fixture-only evidence")
  assert(counts.fetch("contract_only").positive?, "release evidence index must include contract-only evidence")
  assert(counts.fetch("blocked").positive?, "release evidence index must include blocked evidence")
  assert(counts.fetch("skipped").positive?, "release evidence index must include skipped evidence")
  assert(counts.fetch("human_authorized").zero?, "completed product smoke must not retain human-authorized evidence classification")

  claims = report.fetch("claims").to_h { |claim| [claim.fetch("id"), claim] }
  assert(claims.fetch("runtime-owner-read-boundary").fetch("evidence_level") == "implemented", "runtime owner claim must be implemented")
  assert(claims.fetch("runtime-state-backend-lifecycle").fetch("evidence_level") == "fixture-only", "backend lifecycle claim must cover fixture-only evidence")
  assert(claims.fetch("portal-snapshot-safety").fetch("evidence_level") == "contract-only", "Portal snapshot claim must cover contract-only evidence")
  assert(claims.fetch("diagnostics-repair-ai-boundary").fetch("evidence_level") == "blocked", "diagnostics claim must cover blocked evidence")
  assert(claims.fetch("protected-claude-file").fetch("evidence_level") == "blocked", "protected file claim must block when modified")
  assert(claims.fetch("protected-claude-file").fetch("blockers").include?("protected-file"), "protected file claim must name protected-file blocker")
  assert(claims.fetch("unclassified-files").fetch("evidence_level") == "blocked", "unclassified file claim must block when files are unclassified")
  assert(claims.fetch("unclassified-files").fetch("blockers").include?("unclassified-file"), "unclassified file claim must name unclassified-file blocker")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("evidence_level") == "implemented", "product image claim must consume the authorized q4 evidence")
  assert(!claims.fetch("product-image-qemu-acceptance").fetch("human_authorization_required"), "completed product image evidence must not require retroactive authorization")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("blockers").empty?, "completed product image evidence must have no evidence blockers")
  assert(claims.fetch("restricted-heavy-smoke-skipped").fetch("evidence_level") == "skipped", "heavy smoke claim must be skipped by default")
  assert(claims.fetch("kde-first-presence").fetch("evidence_level") == "implemented", "KDE presence claim must be implemented when seven entry points are present")
  assert(claims.fetch("contract-drift").fetch("evidence_level") == "implemented", "contract drift claim must be implemented when no drift is reported")
  assert(claims.values.all? { |claim| claim.fetch("unsafe_gates").values.none? }, "every claim must keep unsafe gates disabled")

  markdown, markdown_stderr, markdown_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "markdown",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path
  )
  assert(markdown_status.success?, "release evidence index Markdown must exit successfully: #{markdown_stderr}")
  assert(markdown.include?("# Release Evidence Index"), "Markdown report must include a title")
  assert(markdown.include?("runtime-owner-read-boundary"), "Markdown report must include runtime owner claim")
  assert(markdown.include?("product-image-qemu-acceptance"), "Markdown report must include product image claim")

  malformed = write_text_fixture("{not-json")
  malformed_stdout, malformed_stderr, malformed_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    malformed.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path
  )
  assert(malformed_status.success?, "malformed report case must still emit a release index: #{malformed_stderr}")
  malformed_report = JSON.parse(malformed_stdout)
  assert(malformed_report.fetch("malformed_report_detected"), "release evidence index must flag malformed reports")
  malformed_claims = malformed_report.fetch("claims").to_h { |claim| [claim.fetch("id"), claim] }
  assert(malformed_claims.fetch("report-integrity").fetch("evidence_level") == "blocked", "report integrity claim must block malformed reports")
ensure
  [implementation, contract_drift, mainline, kde_smoke, malformed].compact.each do |file|
    file.close
    file.unlink
  end
end
