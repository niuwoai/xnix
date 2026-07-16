#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
VERSION = PROJECT_ROOT.join("VERSION").read.strip

REPORT_COMMANDS = {
  implementation: ["ruby", "scripts/implementation_evidence_report.rb", "--format", "json"],
  contract_drift: ["ruby", "scripts/runtime_contract_drift_report.rb", "--format", "json"],
  mainline_review: ["ruby", "scripts/mainline_integration_review.rb", "--format", "json"],
  kde_smoke: ["ruby", "scripts/kde_first_presence_smoke.rb", "--format", "json"]
}.freeze

CLAIM_DEFINITIONS = [
  {
    id: "runtime-owner-read-boundary",
    title: "Runtime owner read boundary is Go-owned and write-disabled",
    domain: "runtime-owner-service",
    source_files: %w[
      cmd/xnix-runtime-owner/main.go
      internal/runtime/owner/service.go
      internal/runtime/owner/session_bus.go
      scripts/runtime_contract_drift_report.rb
    ],
    verification_commands: [
      "go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./internal/runtime/appidentity",
      "ruby scripts/runtime_contract_drift_report.rb --format json"
    ],
    next_follow_up: "Keep production D-Bus ownership blocked until read parity and service packaging evidence are reviewed."
  },
  {
    id: "recipe-artifact-trust",
    title: "Recipes and local artifacts are digest-verified before install readiness",
    domain: "recipe-artifact-trust-pipeline",
    source_files: %w[
      internal/runtime/recipe/store.go
      internal/runtime/artifact/stage.go
      internal/runtime/appidentity/install_plan.go
      internal/runtime/appidentity/application_upgrade_impact.go
    ],
    verification_commands: [
      "go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_recipe_registry.rb",
      "ruby -Ilib test/test_compatibility_install_plan.rb"
    ],
    next_follow_up: "Add more signed-recipe verifier evidence before production trust is enabled."
  },
  {
    id: "runtime-state-backend-lifecycle",
    title: "Runtime state roots track backend inventory and lifecycle without launch",
    domain: "environment-lifecycle-state",
    source_files: %w[
      internal/runtime/environment/lifecycle.go
      internal/runtime/appidentity/backend_manager.go
      internal/runtime/appidentity/backend_lifecycle.go
      internal/runtime/appidentity/state_root_quota_retention.go
      cmd/xnix-runtime-go/backend_group_cli_test.go
    ],
    verification_commands: [
      "go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby scripts/implementation_evidence_report.rb --format json"
    ],
    next_follow_up: "Use quota and retention previews as review evidence; keep real cleanup disabled until deletion-specific approval evidence exists."
  },
  {
    id: "portal-snapshot-safety",
    title: "Portal permission and snapshot evidence stays fake-mode or controlled-root only",
    domain: "portal-snapshot-control-plane",
    source_files: %w[
      internal/runtime/portal/ledger.go
      internal/runtime/snapshot/store.go
      internal/runtime/appidentity/permission_evidence_audit.go
      internal/runtime/appidentity/portal_permission_renewal.go
      internal/runtime/recipe/signature.go
      cmd/xnix-runtime-go/signed_recipe_verifier_commands.go
      internal/runtime/appidentity/snapshot_restore_candidates.go
      cmd/xnix-runtime-go/runtime_safety_cli_test.go
    ],
    verification_commands: [
      "go test ./internal/runtime/portal ./internal/runtime/snapshot ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_portal_access_policy.rb"
    ],
    next_follow_up: "Keep live Portal renewal, production recipe key configuration, and production trust disabled until reviewed with explicit user authorization."
  },
  {
    id: "kde-seven-entrypoints",
    title: "KDE Plasma is the first official shell across seven entry points",
    domain: "kde-activation-shell-materialization",
    source_files: %w[
      scripts/kde_first_presence_smoke.rb
      internal/runtime/appidentity/kde_shell_surface.go
      internal/runtime/appidentity/runtime_policy_explanation_cards.go
      internal/runtime/appidentity/settings_profile_migration.go
      internal/runtime/appidentity/kde_search_visibility.go
      internal/runtime/appidentity/kde_notification_digest.go
      cmd/xnix-runtime-go/kde_notification_digest_commands.go
      kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
    ],
    verification_commands: [
      "ruby scripts/kde_first_presence_smoke.rb --format json",
      "go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_kde_first_presence_smoke_script.rb"
    ],
    next_follow_up: "Add signed recipe verifier evidence before production recipe trust can advance."
  },
  {
    id: "execution-session-evidence",
    title: "Execution and session evidence is recorded while real launch remains disabled",
    domain: "execution-transaction-ledger",
    source_files: %w[
      internal/runtime/execution/ledger.go
      internal/runtime/execution/session.go
      internal/runtime/appidentity/application_readiness.go
      cmd/xnix-runtime-go/execution_ledger_commands.go
    ],
    verification_commands: [
      "go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_runtime_write_gate.rb"
    ],
    next_follow_up: "Keep Launch blocked until Portal, snapshot, service, and write-gate evidence all pass."
  },
  {
    id: "diagnostics-repair-ai-boundary",
    title: "Diagnostics and repair recommendations stay redacted and review-first",
    domain: "diagnostics-repair-ai-boundary",
    source_files: %w[
      internal/runtime/diagnostics/record.go
      internal/runtime/appidentity/ai_diagnostics.go
      internal/runtime/appidentity/diagnostic_history.go
      internal/runtime/appidentity/crash_hang_signal_summary.go
      internal/runtime/appidentity/support_case_timeline.go
      cmd/xnix-runtime-go/diagnostic_record_commands.go
      cmd/xnix-runtime-go/support_case_timeline_commands.go
    ],
    verification_commands: [
      "go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go",
      "ruby -Ilib test/test_ai_diagnostic_input.rb"
    ],
    next_follow_up: "Add a diagnostics playbook preview that explains support timeline findings without creating tickets, exporting bundles, calling providers, repairing, launching backends, or mutating the host."
  },
  {
    id: "evidence-drift-harness",
    title: "Release review has local evidence, drift, and lane-intake guardrails",
    domain: "developer-verification-harness",
    source_files: %w[
      scripts/implementation_evidence_report.rb
      scripts/runtime_contract_drift_report.rb
      scripts/mainline_integration_review.rb
      scripts/release_evidence_index.rb
      scripts/offline_application_fixture_matrix.rb
      scripts/merge_readiness_packet.rb
      internal/runtime/appidentity/offline_application_fixture_matrix.go
    ],
    verification_commands: [
      "ruby scripts/implementation_evidence_report.rb --format json",
      "ruby scripts/runtime_contract_drift_report.rb --format json",
      "ruby scripts/mainline_integration_review.rb --format json",
      "ruby scripts/merge_readiness_packet.rb --format json",
      "ruby -Ilib test/test_offline_application_fixture_matrix.rb"
    ],
    next_follow_up: "Keep one merge readiness packet, one release evidence index, and one offline fixture matrix per branch intake before staging."
  }
].freeze

def parse_options
  options = {
    format: "json",
    reports: {}
  }
  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/release_evidence_index.rb [--format json|markdown] [report fixtures]"
    parser.on("--format FORMAT", "Output format: json or markdown") { |value| options[:format] = value }
    parser.on("--implementation-report PATH", "Use an existing implementation evidence JSON report") { |value| options[:reports][:implementation] = value }
    parser.on("--contract-drift-report PATH", "Use an existing Runtime contract drift JSON report") { |value| options[:reports][:contract_drift] = value }
    parser.on("--mainline-review PATH", "Use an existing mainline integration review JSON report") { |value| options[:reports][:mainline_review] = value }
    parser.on("--kde-smoke-report PATH", "Use an existing KDE-first presence smoke JSON report") { |value| options[:reports][:kde_smoke] = value }
  end.parse!

  unless %w[json markdown].include?(options[:format])
    abort "release_evidence_index supports --format json or --format markdown"
  end
  options
end

def parse_report_json(text, source)
  [JSON.parse(text), nil]
rescue JSON::ParserError => e
  [nil, { source: source, error: "malformed-report", detail: e.message }]
end

def load_report(kind, fixture_path)
  if fixture_path
    return parse_report_json(Pathname.new(fixture_path).read, fixture_path)
  end

  command = REPORT_COMMANDS.fetch(kind)
  stdout, stderr, status = Open3.capture3(*command, chdir: PROJECT_ROOT.to_s)
  return parse_report_json(stdout, command.join(" ")) if status.success?

  [nil, { source: command.join(" "), error: "report-command-failed", detail: stderr.strip }]
end

def evidence_level_from_status(status)
  case status
  when "smoke-owned", "state-root-implemented", "production-gated"
    "implemented"
  when "fixture-implemented"
    "fixture-only"
  when "contract-only"
    "contract-only"
  else
    "blocked"
  end
end

def claim_state(level)
  case level
  when "implemented"
    "ready-for-review"
  when "fixture-only", "contract-only"
    "needs-more-evidence"
  when "human-authorized"
    "requires-human-authorization"
  when "skipped"
    "skipped-by-default"
  else
    "blocked"
  end
end

def domain_claim(definition, domains, report_errors)
  domain = domains[definition.fetch(:domain)]
  level = domain ? evidence_level_from_status(domain.fetch("status", "missing")) : "blocked"
  blockers = []
  blockers << "implementation evidence report is unavailable or malformed" if report_errors.any? { |err| err.fetch(:source).include?("implementation") || err.fetch(:source).include?("implementation_evidence_report") }
  blockers += domain.fetch("contract_files_missing", []) if domain
  blockers += domain.fetch("fixture_files_missing", []) if domain && level == "contract-only"
  blockers << "domain evidence is missing" unless domain

  {
    id: definition.fetch(:id),
    title: definition.fetch(:title),
    release_claim: definition.fetch(:title),
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: definition.fetch(:source_files),
    verification_commands: definition.fetch(:verification_commands),
    current_evidence: domain ? domain.fetch("summary", domain.fetch("status", "missing")) : "No domain evidence was available.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: false,
    blockers: blockers.uniq,
    next_branch_sized_follow_up: definition.fetch(:next_follow_up)
  }
end

def disabled_unsafe_gates
  {
    "docker_executed" => false,
    "qemu_executed" => false,
    "network_checks_run" => false,
    "package_manager_invoked" => false,
    "backend_launch_enabled" => false,
    "host_root_modified" => false,
    "automatic_staging_enabled" => false,
    "automatic_release_tagging_enabled" => false
  }
end

def build_report(options)
  reports = {}
  report_errors = []
  REPORT_COMMANDS.each_key do |kind|
    report, error = load_report(kind, options[:reports][kind])
    reports[kind] = report
    report_errors << error if error
  end

  implementation = reports[:implementation] || {}
  domains = implementation.fetch("domains", []).to_h { |domain| [domain.fetch("id"), domain] }
  mainline = reports[:mainline_review] || {}
  contract_drift = reports[:contract_drift] || {}
  kde_smoke = reports[:kde_smoke] || {}

  claims = CLAIM_DEFINITIONS.map { |definition| domain_claim(definition, domains, report_errors) }
  claims << report_integrity_claim(report_errors)
  claims << protected_file_claim(mainline)
  claims << unclassified_file_claim(mainline)
  claims << contract_drift_claim(contract_drift)
  claims << kde_presence_claim(kde_smoke)
  claims << product_image_claim(domains["atomic-kde-image-qemu-acceptance"])
  claims << skipped_heavy_smoke_claim

  {
    "version" => VERSION,
    "schema_version" => "xnix.runtime.release_evidence_index.v1",
    "report_type" => "release-evidence-index",
    "source" => "implementation-evidence+contract-drift+mainline-review+kde-first-presence",
    "runtime_owned" => true,
    "go_runtime_backed" => false,
    "ruby_report_only" => true,
    "kde_policy_owner" => false,
    "offline_default" => true,
    "docker_executed" => false,
    "qemu_executed" => false,
    "network_checks_run" => false,
    "package_manager_invoked" => false,
    "backend_launch_enabled" => false,
    "host_root_modified" => false,
    "automatic_staging_enabled" => false,
    "automatic_release_tagging_enabled" => false,
    "malformed_report_detected" => report_errors.any? { |error| error.fetch(:error) == "malformed-report" },
    "report_errors" => report_errors,
    "claims" => claims,
    "claim_count" => claims.length,
    "counts" => count_claims(claims),
    "human_authorization_required" => claims.any? { |claim| claim.fetch(:human_authorization_required) },
    "release_ready" => false,
    "desktop_safe_summary" => "Release-critical evidence is indexed offline; Docker, QEMU, network checks, package managers, backend launch, automatic staging, release tagging, and host-root mutation remain disabled."
  }
end

def report_integrity_claim(report_errors)
  level = report_errors.empty? ? "implemented" : "blocked"
  {
    id: "report-integrity",
    title: "Input reports are parseable and safe to consume",
    release_claim: "The release index can consume implementation, drift, mainline, and KDE smoke reports.",
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: %w[
      scripts/implementation_evidence_report.rb
      scripts/runtime_contract_drift_report.rb
      scripts/mainline_integration_review.rb
      scripts/kde_first_presence_smoke.rb
    ],
    verification_commands: REPORT_COMMANDS.values.map { |command| command.join(" ") },
    current_evidence: report_errors.empty? ? "All input reports were parseable." : "One or more input reports could not be parsed or executed.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: false,
    blockers: report_errors.map { |error| "#{error.fetch(:source)}: #{error.fetch(:error)}" },
    next_branch_sized_follow_up: "Fix malformed or failing reports before release review."
  }
end

def protected_file_claim(mainline)
  modified = mainline.fetch("protected_claude_file_modified", false)
  level = modified ? "blocked" : "implemented"
  {
    id: "protected-claude-file",
    title: "Protected Claude implementation package file is not staged by this lane",
    release_claim: "Codex release review must not silently modify docs/claude-code-implementation-packages.md.",
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: %w[scripts/mainline_integration_review.rb docs/claude-code-implementation-packages.md],
    verification_commands: ["git diff -- docs/claude-code-implementation-packages.md", "ruby scripts/mainline_integration_review.rb --format json"],
    current_evidence: modified ? "Protected Claude-owned file is modified." : "Protected Claude-owned file is not modified in the intake report.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: modified,
    blockers: modified ? ["protected-file"] : [],
    next_branch_sized_follow_up: "Ask the Claude owner or the user to reconcile protected-file changes separately."
  }
end

def unclassified_file_claim(mainline)
  count = mainline.fetch("unclassified_file_count", 0).to_i
  level = count.positive? ? "blocked" : "implemented"
  {
    id: "unclassified-files",
    title: "Changed files are classified into review lanes",
    release_claim: "Release intake should not include unclassified files.",
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: %w[scripts/mainline_integration_review.rb docs/mainline-integration-checkpoint.md],
    verification_commands: ["ruby scripts/mainline_integration_review.rb --format json"],
    current_evidence: "#{count} unclassified file(s) reported.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: count.positive?,
    blockers: count.positive? ? ["unclassified-file"] : [],
    next_branch_sized_follow_up: "Classify, remove, or split unclassified files before staging."
  }
end

def contract_drift_claim(contract_drift)
  drift = contract_drift.fetch("drift_detected", false)
  level = drift ? "blocked" : "implemented"
  {
    id: "contract-drift",
    title: "Runtime route and D-Bus contract drift is reviewed",
    release_claim: "Runtime read routes and disabled write methods stay aligned across reports and smoke adapters.",
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: %w[scripts/runtime_contract_drift_report.rb runtime/dbus/org.xnix.Compatibility1.xml],
    verification_commands: ["ruby scripts/runtime_contract_drift_report.rb --format json"],
    current_evidence: drift ? "Contract drift was detected." : "No contract drift was reported.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: false,
    blockers: drift ? ["contract-drift"] : [],
    next_branch_sized_follow_up: "Repair route parity before release review continues."
  }
end

def kde_presence_claim(kde_smoke)
  entrypoint_count = kde_smoke.fetch("entrypoint_count", 0).to_i
  level = entrypoint_count == 7 ? "implemented" : "blocked"
  {
    id: "kde-first-presence",
    title: "KDE first-release presence covers all seven official entry points",
    release_claim: "Windows-compatible apps appear as normal KDE applications across launcher, task manager, file manager, tray, notifications, Compatibility Center, and settings.",
    evidence_level: level,
    state: claim_state(level),
    evidence_source_files: %w[scripts/kde_first_presence_smoke.rb docs/kde-first-presence-smoke-spec.md],
    verification_commands: ["ruby scripts/kde_first_presence_smoke.rb --format json"],
    current_evidence: "#{entrypoint_count} KDE entry point(s) reported.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: false,
    blockers: entrypoint_count == 7 ? [] : ["missing-kde-entrypoint-evidence"],
    next_branch_sized_follow_up: "Keep KDE as the official shell while Runtime remains the policy owner."
  }
end

def product_image_claim(domain)
  {
    id: "product-image-qemu-acceptance",
    title: "Product image and QEMU acceptance remain human-authorized",
    release_claim: "Full product image proof requires explicit restricted Docker or QEMU authorization.",
    evidence_level: "human-authorized",
    state: claim_state("human-authorized"),
    evidence_source_files: %w[scripts/restricted_product_smoke_packet.rb internal/runtime/image/restricted_smoke_packet.go scripts/full_smoke.rb lib/xnix/full_smoke_report.rb buildroot/configs/xnix_x86_64_defconfig],
    verification_commands: ["ruby scripts/restricted_product_smoke_packet.rb --format json", "ruby scripts/full_smoke.rb"],
    current_evidence: domain ? domain.fetch("summary", "Restricted packet evidence exists but image execution is gated.") : "Restricted product smoke packet evidence exists; image execution remains gated.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: true,
    blockers: ["requires-explicit-docker-or-qemu-authorization"],
    next_branch_sized_follow_up: "Run restricted product smoke only after the user authorizes Docker or QEMU."
  }
end

def skipped_heavy_smoke_claim
  {
    id: "restricted-heavy-smoke-skipped",
    title: "Heavy smoke commands are skipped by default",
    release_claim: "The release evidence index does not run Docker, QEMU, package managers, backend launch, staging, or tagging.",
    evidence_level: "skipped",
    state: claim_state("skipped"),
    evidence_source_files: %w[scripts/release_evidence_index.rb docs/claude-code-seventh-wave-task-batch.md],
    verification_commands: ["ruby scripts/release_evidence_index.rb --format json"],
    current_evidence: "Heavy commands are listed as follow-up evidence and are not executed by this report.",
    unsafe_gates: disabled_unsafe_gates,
    human_authorization_required: true,
    blockers: ["offline-default-skips-heavy-smoke"],
    next_branch_sized_follow_up: "Ask for explicit restricted-smoke authorization before image acceptance."
  }
end

def count_claims(claims)
  counts = Hash.new(0)
  claims.each { |claim| counts[claim.fetch(:evidence_level)] += 1 }
  %w[implemented fixture-only contract-only blocked skipped human-authorized].to_h { |level| [level.tr("-", "_"), counts[level]] }
end

def render_markdown(report)
  lines = []
  lines << "# Release Evidence Index"
  lines << ""
  lines << "- Version: #{report.fetch("version")}"
  lines << "- Release ready: #{report.fetch("release_ready")}"
  lines << "- Human authorization required: #{report.fetch("human_authorization_required")}"
  lines << "- Docker executed: #{report.fetch("docker_executed")}"
  lines << "- QEMU executed: #{report.fetch("qemu_executed")}"
  lines << ""
  lines << "## Counts"
  lines << ""
  report.fetch("counts").each do |level, count|
    lines << "- #{level}: #{count}"
  end
  lines << ""
  lines << "## Claims"
  lines << ""
  report.fetch("claims").each do |claim|
    lines << "### #{claim_value(claim, "id")}"
    lines << ""
    lines << "- Evidence level: #{claim_value(claim, "evidence_level")}"
    lines << "- State: #{claim_value(claim, "state")}"
    lines << "- Human authorization required: #{claim_value(claim, "human_authorization_required")}"
    lines << "- Current evidence: #{claim_value(claim, "current_evidence")}"
    lines << "- Next follow-up: #{claim_value(claim, "next_branch_sized_follow_up")}"
    lines << ""
  end
  lines.join("\n")
end

def claim_value(claim, key)
  claim.fetch(key) { claim.fetch(key.to_sym) }
end

options = parse_options
report = build_report(options)
case options[:format]
when "json"
  puts JSON.pretty_generate(report)
when "markdown"
  puts render_markdown(report)
end
