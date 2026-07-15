#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
MAINLINE_DOCUMENT = "docs/claude-code-mainline-implementation-plan.md"
MAINLINE_FIRST_WAVE = [
  {
    mainline_package: "M1",
    suggested_branch: "codex/runtime-owner-read-service",
    reason: "Runtime-owned read paths should replace preview-only dispatch before broader product ownership advances.",
    minimal_mergeable_outcome: "A constrained session-bus owner serves the Go owner read dispatch table for every D-Bus read-only Runtime method and fails every write method closed."
  },
  {
    mainline_package: "M2",
    suggested_branch: "codex/recipe-artifact-trust-pipeline",
    reason: "Install, environment, activation, and execution work need trusted local inputs first.",
    minimal_mergeable_outcome: "Local recipes and fixture artifacts verify digests, stage under a controlled root, and produce blocked receipts for invalid inputs."
  },
  {
    mainline_package: "M3",
    suggested_branch: "codex/environment-lifecycle-state",
    reason: "Execution readiness needs durable environment state before launch paths become meaningful.",
    minimal_mergeable_outcome: "A state-root lifecycle store records missing, planned, staged, ready, repair-required, and blocked states."
  },
  {
    mainline_package: "M8",
    suggested_branch: "codex/implementation-evidence-harness",
    reason: "Contract-heavy work needs an evidence harness that keeps empty domains visible.",
    minimal_mergeable_outcome: "JSON and Markdown reports classify domain evidence and fail on orphan contract-only surfaces."
  }
].freeze
WRITE_METHODS = %w[
  InstallRecipe
  Launch
  CreateSnapshot
  RestoreSnapshot
].freeze

STATUS_ORDER = {
  "missing" => 0,
  "contract-only" => 1,
  "fixture-implemented" => 2,
  "state-root-implemented" => 3,
  "smoke-owned" => 4,
  "production-gated" => 5
}.freeze

DOMAIN_DEFINITIONS = [
  {
    id: "runtime-owner-service",
    name: "Runtime owner service",
    package: "P1",
    mainline_package: "M1",
    contract_files: %w[
      runtime/dbus/org.xnix.Compatibility1.xml
      internal/runtime/appidentity/runtime_owner_readiness.go
      internal/runtime/appidentity/runtime_owner_route_manifest.go
    ],
    fixture_files: %w[
      cmd/xnix-runtime-owner/main.go
      internal/runtime/owner/candidate.go
      internal/runtime/owner/dispatch.go
      internal/runtime/owner/dispatch_test.go
      internal/runtime/owner/lifecycle.go
      internal/runtime/owner/lifecycle_test.go
      internal/runtime/owner/smoke_batch.go
      internal/runtime/owner/smoke_batch_test.go
    ],
    state_files: [],
    smoke_files: %w[
      runtime/dbus/xnix_compatd_smoke.c
      runtime/dbus/xnix_compatd_runtime_models.inc
      scripts/runtime_owner_candidate_smoke.rb
      scripts/dbus_session_smoke.rb
      test/test_runtime_owner_candidate_smoke_script.rb
    ],
    gate_tokens: {
      "cmd/xnix-runtime-owner/main.go" => %w[smoke-owner deny-write],
      "internal/runtime/owner/lifecycle.go" => %w[xnix.runtime.owner_lifecycle_event.v1 runtime-owner-lifecycle-event preview-complete],
      "internal/runtime/owner/smoke_batch.go" => %w[xnix.runtime.owner_smoke_batch.v1 runtime-owner-smoke-batch-record restricted-session-owner-call-batch],
      "runtime/dbus/xnix_compatd_smoke.c" => %w[go_owner_write_gate_dispatch],
      "runtime/dbus/xnix_compatd_runtime_models.inc" => %w[go_owner_dispatch_available go_owner_dispatch_json go-runtime-owner-dispatch+c-smoke-bridge],
      "internal/runtime/appidentity/runtime_owner_readiness.go" => %w[ProductionOwnerEnabled ProductionBusClaimed WriteMethodsEnabled]
    },
    summary: "The Go owner candidate has full D-Bus read dispatch coverage, lifecycle JSONL, smoke-batch read/write evidence, and selected C D-Bus bridge evidence for Go owner payloads; production ownership remains gated."
  },
  {
    id: "recipe-artifact-trust-pipeline",
    name: "Recipe and artifact trust pipeline",
    package: "P2",
    mainline_package: "M2",
    contract_files: %w[
      runtime/recipes/registry.json
      internal/runtime/appidentity/runtime_owner_recipe_trust.go
      internal/runtime/appidentity/artifact_manifest.go
      internal/runtime/appidentity/acquisition_preflight.go
      internal/runtime/appidentity/install_plan.go
    ],
    fixture_files: %w[
      internal/runtime/recipe/store.go
      internal/runtime/recipe/trust.go
      internal/runtime/artifact/artifact.go
      internal/runtime/artifact/acquire.go
      internal/runtime/artifact/cache.go
      internal/runtime/artifact/artifact_test.go
      cmd/xnix-runtime-go/artifact_stage_commands.go
      cmd/xnix-runtime-go/artifact_stage_cli_test.go
    ],
    state_files: %w[
      internal/runtime/artifact/stage.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/artifact/acquire.go" => %w[ErrNetworkDisabled NewNetworkSource],
      "internal/runtime/artifact/stage.go" => %w[xnix.runtime.artifact_stage_receipt.v1 go-runtime-local-fixture-artifact-staging artifact-ledger cache_root_path_exposed fixture_root_path_exposed package_manager_invoked],
      "cmd/xnix-runtime-go/artifact_stage_commands.go" => %w[artifact-stage-record cache-root fixture-root],
      "internal/runtime/appidentity/runtime_owner_recipe_trust.go" => %w[ProductionRecipeTrustReady SignedRecipeValidation]
    },
    summary: "Local recipe and artifact staging persists receipts under a controlled cache root; production signing and network fetch remain gated."
  },
  {
    id: "environment-lifecycle-state",
    name: "Environment lifecycle state",
    package: "P3",
    mainline_package: "M3",
    contract_files: %w[
      internal/runtime/appidentity/backend_environment.go
      internal/runtime/appidentity/backend_binding.go
      internal/runtime/appidentity/backend_lifecycle.go
      internal/runtime/appidentity/execution_readiness.go
    ],
    fixture_files: %w[
      internal/runtime/environment/environment.go
      internal/runtime/environment/lifecycle.go
      internal/runtime/environment/lifecycle_test.go
    ],
    state_files: %w[
      internal/runtime/appidentity/state_root.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/environment/lifecycle.go" => %w[repair-required StateBlocked],
      "internal/runtime/appidentity/execution_readiness.go" => %w[LaunchAllowed BackendDetailsExposed]
    },
    summary: "Runtime environment lifecycle logic exists with state-root integration; real backend starts remain disabled."
  },
  {
    id: "portal-snapshot-control-plane",
    name: "Portal and snapshot control plane",
    package: "P4",
    mainline_package: "M4",
    contract_files: %w[
      internal/runtime/appidentity/portal_access_policy.go
      internal/runtime/appidentity/snapshot_plan.go
    ],
    fixture_files: %w[
      internal/runtime/portal/broker.go
      internal/runtime/portal/request.go
      internal/runtime/portal/broker_test.go
      internal/runtime/portal/request_test.go
    ],
    state_files: %w[
      internal/runtime/snapshot/store.go
      internal/runtime/snapshot/store_test.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/portal/broker.go" => %w[FakeBroker denied cancelled failed],
      "internal/runtime/snapshot/store.go" => %w[Rollback Verify]
    },
    summary: "Fake Portal records and content-addressed snapshots exist under controlled roots; real Portal and system snapshots remain gated."
  },
  {
    id: "kde-activation-shell-materialization",
    name: "KDE activation and shell materialization",
    package: "P5",
    mainline_package: "M5",
    contract_files: %w[
      internal/runtime/appidentity/desktop_activation_bundle.go
      internal/runtime/appidentity/desktop_activation_status.go
      kde/dolphin/servicemenus/xnix-open-with-compatibility.desktop
      kde/plasmoids/org.xnix.compatibilitycenter/metadata.json
    ],
    fixture_files: %w[
      internal/runtime/activation/stage.go
      internal/runtime/activation/stage_test.go
      scripts/install_runtime_activation.rb
      scripts/runtime_activation_smoke.rb
    ],
    state_files: %w[
      internal/runtime/activation/stage.go
    ],
    smoke_files: %w[
      scripts/runtime_activation_smoke.rb
      test/test_runtime_activation_smoke_script.rb
    ],
    gate_tokens: {
      "internal/runtime/activation/stage.go" => ["refusing to stage", "host_root_modified", "RollbackReceiptWritten"],
      "internal/runtime/appidentity/desktop_activation_transaction.go" => %w[TransactionCommitted WriteGate]
    },
    summary: "KDE activation can stage desktop files, MIME data, service menus, manifests, and receipts under explicit roots."
  },
  {
    id: "execution-transaction-ledger",
    name: "Execution transaction ledger",
    package: "P6",
    mainline_package: "M6",
    contract_files: %w[
      internal/runtime/appidentity/launch_intent.go
      internal/runtime/appidentity/execution_request.go
      internal/runtime/appidentity/execution_preflight.go
      internal/runtime/appidentity/execution_transaction.go
      internal/runtime/appidentity/execution_session_status.go
    ],
    fixture_files: %w[
      internal/runtime/execution/execution.go
      internal/runtime/execution/execution_test.go
      internal/runtime/execution/ledger_test.go
      cmd/xnix-runtime-go/execution_ledger_commands.go
      cmd/xnix-runtime-go/execution_ledger_cli_test.go
    ],
    state_files: %w[
      internal/runtime/execution/ledger.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/execution/execution.go" => %w[blocked runtimeWriteGate],
      "internal/runtime/execution/ledger.go" => %w[execution-transaction-ledger-record go-runtime-state-root-execution-ledger state_root_path_exposed],
      "cmd/xnix-runtime-go/execution_ledger_commands.go" => %w[execution-ledger-record state-root],
      "internal/runtime/appidentity/runtime_write_gate.go" => %w[Launch WriteMethodDisabled]
    },
    summary: "Execution transaction records persist under a controlled state root; real Launch remains disabled."
  },
  {
    id: "diagnostics-repair-ai-boundary",
    name: "Diagnostics, repair, and AI boundary",
    package: "P7",
    mainline_package: "M7",
    contract_files: %w[
      internal/runtime/appidentity/test_plan.go
      internal/runtime/appidentity/test_result.go
      internal/runtime/appidentity/repair_plan.go
      internal/runtime/appidentity/ai_diagnostics.go
    ],
    fixture_files: %w[
      internal/runtime/diagnostics/runner.go
      internal/runtime/diagnostics/ai.go
      internal/runtime/diagnostics/diagnostics_test.go
      internal/runtime/diagnostics/ai_test.go
      cmd/xnix-runtime-go/diagnostic_record_commands.go
      cmd/xnix-runtime-go/diagnostic_record_cli_test.go
      internal/runtime/appidentity/diagnostic_history_test.go
    ],
    state_files: %w[
      internal/runtime/diagnostics/record.go
      internal/runtime/diagnostics/history.go
      internal/runtime/appidentity/diagnostic_history.go
    ],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/diagnostics/ai.go" => %w[DisabledProvider FakeProvider],
      "internal/runtime/diagnostics/record.go" => %w[xnix.runtime.diagnostic_run_record.v1 diagnostic-run-record go-runtime-state-root-diagnostic-run-record diagnostics-ledger state_root_path_exposed fixture_path_exposed ai_provider_called real_ai_provider_enabled repair_executed],
      "internal/runtime/diagnostics/history.go" => %w[xnix.runtime.diagnostic_run_history.v1 diagnostic-run-history go-runtime-state-root-diagnostic-run-history state_root_path_exposed ai_provider_called repair_executed],
      "internal/runtime/appidentity/diagnostic_history.go" => %w[xnix.runtime.diagnostic_history_preview.v1 diagnostic-history-preview go-runtime-state-root-diagnostic-run-history+kde-read-model compatibility_center_card safe_for_ai_diagnostics ai_provider_call_enabled repair_execution_enabled backend_details_exposed],
      "cmd/xnix-runtime-go/diagnostic_record_commands.go" => %w[diagnostic-run-record diagnostic-run-history diagnostic-history-preview state-root fixture run-id],
      "internal/runtime/appidentity/ai_diagnostics.go" => %w[AIProviderCalled AutoExecutionAllowed]
    },
    summary: "Fixture diagnostics can persist state-root run records, expose KDE-safe history summaries, and project diagnostic history into a Compatibility Center read model; real provider calls, backend launch, and auto-repair remain disabled."
  },
  {
    id: "atomic-kde-image-qemu-acceptance",
    name: "Atomic KDE image and QEMU acceptance",
    package: "P8",
    mainline_package: "M9",
    contract_files: %w[
      image/kinoite/manifest.json
      image/kinoite/Containerfile
      docs/kde-image-pipeline.md
    ],
    fixture_files: %w[
      internal/runtime/image/manifest.go
      internal/runtime/image/smoke.go
      internal/runtime/image/smoke_test.go
      scripts/build_kde_image.rb
      scripts/boot_kde_image.rb
    ],
    state_files: [],
    smoke_files: %w[
      scripts/boot_kde_image.rb
      test/test_kde_image.rb
      test/test_qemu.rb
    ],
    gate_tokens: {
      "image/kinoite/Containerfile" => %w[org.xnix.image org.xnix.flagship],
      "internal/runtime/image/smoke.go" => %w[SerialMarkers RuntimeReady]
    },
    summary: "KDE image validation and QEMU smoke scaffolding exists; full product image proof remains a gated milestone."
  },
  {
    id: "developer-verification-harness",
    name: "Developer verification harness",
    package: "P9",
    mainline_package: "M8",
    contract_files: %w[
      scripts/verify_layout.rb
      scripts/runtime_contract_drift_report.rb
      scripts/kde_first_presence_smoke.rb
    ],
    fixture_files: %w[
      scripts/implementation_evidence_report.rb
      test/test_implementation_evidence_report.rb
    ],
    state_files: [],
    smoke_files: [],
    gate_tokens: {
      "scripts/implementation_evidence_report.rb" => %w[contract-only fixture-implemented state-root-implemented smoke-owned production-gated]
    },
    summary: "Local evidence reporting classifies implementation depth and flags orphan Runtime read methods."
  }
].freeze

def parse_options
  options = {
    root: PROJECT_ROOT,
    format: "json"
  }

  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/implementation_evidence_report.rb [--root PATH] [--format json|markdown]"
    parser.on("--root PATH", "Project root to inspect") do |value|
      options[:root] = Pathname.new(value).expand_path
    end
    parser.on("--format FORMAT", "Output format: json or markdown") do |value|
      options[:format] = value
    end
  end.parse!

  unless %w[json markdown].include?(options[:format])
    abort "implementation_evidence_report supports --format json or --format markdown"
  end

  options
end

def read_project_file(root, relative_path)
  root.join(relative_path).read
rescue Errno::ENOENT
  ""
end

def file_present?(root, relative_path)
  root.join(relative_path).file?
end

def present_files(root, relative_paths)
  relative_paths.select { |path| file_present?(root, path) }
end

def missing_files(root, relative_paths)
  relative_paths.reject { |path| file_present?(root, path) }
end

def token_results(root, token_map)
  token_map.flat_map do |relative_path, tokens|
    source = read_project_file(root, relative_path)
    tokens.map do |token|
      {
        "file" => relative_path,
        "token" => token,
        "present" => source.include?(token)
      }
    end
  end
end

def choose_status(contract_missing:, fixture_missing:, state_missing:, smoke_missing:, gate_missing:, has_state_files:, has_smoke_files:)
  return "missing" unless contract_missing.empty?
  return "contract-only" unless fixture_missing.empty?

  status = "fixture-implemented"
  status = "state-root-implemented" if has_state_files && state_missing.empty?
  status = "smoke-owned" if has_smoke_files && smoke_missing.empty?
  status
end

def domain_report(root, definition)
  contract_missing = missing_files(root, definition.fetch(:contract_files))
  fixture_missing = missing_files(root, definition.fetch(:fixture_files))
  state_files = definition.fetch(:state_files)
  smoke_files = definition.fetch(:smoke_files)
  state_missing = missing_files(root, state_files)
  smoke_missing = missing_files(root, smoke_files)
  tokens = token_results(root, definition.fetch(:gate_tokens))
  gate_missing = tokens.reject { |item| item.fetch("present") }

  status = choose_status(
    contract_missing: contract_missing,
    fixture_missing: fixture_missing,
    state_missing: state_missing,
    smoke_missing: smoke_missing,
    gate_missing: gate_missing,
    has_state_files: !state_files.empty?,
    has_smoke_files: !smoke_files.empty?
  )

  {
    "id" => definition.fetch(:id),
    "name" => definition.fetch(:name),
    "package" => definition.fetch(:package),
    "mainline_package" => definition.fetch(:mainline_package),
    "mainline_document" => MAINLINE_DOCUMENT,
    "status" => status,
    "rank" => STATUS_ORDER.fetch(status),
    "contract_files_present" => present_files(root, definition.fetch(:contract_files)),
    "contract_files_missing" => contract_missing,
    "fixture_files_present" => present_files(root, definition.fetch(:fixture_files)),
    "fixture_files_missing" => fixture_missing,
    "state_files_present" => present_files(root, state_files),
    "state_files_missing" => state_missing,
    "smoke_files_present" => present_files(root, smoke_files),
    "smoke_files_missing" => smoke_missing,
    "gate_tokens" => tokens,
    "gate_tokens_missing" => gate_missing,
    "production_gate_evidence" => gate_missing.empty?,
    "host_root_modified" => false,
    "network_required" => false,
    "privileged_container_required" => false,
    "backend_launch_enabled" => false,
    "summary" => definition.fetch(:summary)
  }
end

def dbus_contract_methods(root)
  read_project_file(root, "runtime/dbus/org.xnix.Compatibility1.xml").scan(/<method name="([^"]+)"/).flatten.uniq
end

def go_string_map_keys(source, variable_name)
  body = source[/var #{Regexp.escape(variable_name)} = map\[string\]string\{(.*?)\n\}/m, 1].to_s
  body.scan(/"([^"]+)"\s*:/).flatten
end

def owner_route_methods(root)
  source = read_project_file(root, "internal/runtime/appidentity/runtime_owner_route_manifest.go")
  go_string_map_keys(source, "runtimeOwnerRouteGoCommands")
end

def counts_for(domains)
  statuses = STATUS_ORDER.keys.to_h { |status| [status, 0] }
  domains.each { |domain| statuses[domain.fetch("status")] += 1 }

  {
    "total" => domains.length,
    "missing" => statuses.fetch("missing"),
    "contract_only" => statuses.fetch("contract-only"),
    "fixture_implemented" => statuses.fetch("fixture-implemented"),
    "state_root_implemented" => statuses.fetch("state-root-implemented"),
    "smoke_owned" => statuses.fetch("smoke-owned"),
    "production_gated" => statuses.fetch("production-gated"),
    "production_gate_evidence" => domains.count { |domain| domain.fetch("production_gate_evidence") }
  }
end

def mainline_first_wave(domains)
  domains_by_mainline = domains.to_h { |domain| [domain.fetch("mainline_package"), domain] }

  MAINLINE_FIRST_WAVE.filter_map do |entry|
    domain = domains_by_mainline[entry.fetch(:mainline_package)]
    next unless domain

    {
      "mainline_package" => entry.fetch(:mainline_package),
      "package" => domain.fetch("package"),
      "domain_id" => domain.fetch("id"),
      "domain_name" => domain.fetch("name"),
      "status" => domain.fetch("status"),
      "rank" => domain.fetch("rank"),
      "suggested_branch" => entry.fetch(:suggested_branch),
      "reason" => entry.fetch(:reason),
      "minimal_mergeable_outcome" => entry.fetch(:minimal_mergeable_outcome),
      "host_root_modified" => false,
      "network_required" => false,
      "privileged_container_required" => false,
      "backend_launch_enabled" => false
    }
  end
end

def build_report(root)
  version = read_project_file(root, "VERSION").strip
  domains = DOMAIN_DEFINITIONS.map { |definition| domain_report(root, definition) }
  first_wave = mainline_first_wave(domains)
  read_methods = dbus_contract_methods(root) - WRITE_METHODS
  route_methods = owner_route_methods(root)
  orphan_read_methods = read_methods - route_methods
  counts = counts_for(domains)
  highest_rank = domains.map { |domain| domain.fetch("rank") }.max || 0

  {
    "version" => version,
    "schema_version" => "xnix.runtime.implementation_evidence_report.v1",
    "report_type" => "implementation-evidence-report",
    "source" => "filesystem+runtime-contract+owner-route-manifest+empty-domain-packages+mainline-implementation-plan",
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "mainline_document" => MAINLINE_DOCUMENT,
    "mainline_plan_present" => file_present?(root, MAINLINE_DOCUMENT),
    "mainline_package_count" => domains.map { |domain| domain.fetch("mainline_package") }.uniq.length,
    "mainline_first_wave" => first_wave,
    "next_dispatch_packages" => first_wave,
    "next_dispatch_summary" => "First-wave mainline dispatch recommends M1, M2, M3, and M8 before execution or image acceptance.",
    "domain_status_order" => STATUS_ORDER.keys,
    "domains" => domains,
    "counts" => counts,
    "read_only_method_count" => read_methods.length,
    "orphan_read_methods" => orphan_read_methods,
    "orphan_preview_methods_detected" => !orphan_read_methods.empty?,
    "write_methods" => WRITE_METHODS,
    "write_methods_supported" => false,
    "write_method_dispatch_enabled" => false,
    "network_required" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "backend_launch_enabled" => false,
    "production_ready" => false,
    "highest_evidence_rank" => highest_rank,
    "highest_evidence_status" => STATUS_ORDER.key(highest_rank),
    "desktop_safe_summary" => "Implementation evidence spans #{counts.fetch("total")} domains; production readiness remains gated."
  }
end

def render_markdown(report)
  lines = [
    "# Implementation Evidence Report",
    "",
    "- Version: #{report.fetch("version")}",
    "- Schema: #{report.fetch("schema_version")}",
    "- Domains: #{report.fetch("counts").fetch("total")}",
    "- Highest evidence status: #{report.fetch("highest_evidence_status")}",
    "- Mainline document: #{report.fetch("mainline_document")}",
    "- Mainline plan present: #{report.fetch("mainline_plan_present")}",
    "- Next dispatch: #{report.fetch("next_dispatch_summary")}",
    "- Orphan read methods detected: #{report.fetch("orphan_preview_methods_detected")}",
    "- Production ready: #{report.fetch("production_ready")}",
    "- Host root modified: #{report.fetch("host_root_modified")}",
    "",
    "| Mainline | Package | Domain | Status | Missing fixtures | Missing gates |",
    "| --- | --- | --- | --- | --- | --- |"
  ]

  report.fetch("domains").each do |domain|
    missing_fixtures = domain.fetch("fixture_files_missing").empty? ? "-" : domain.fetch("fixture_files_missing").join(", ")
    missing_gates = domain.fetch("gate_tokens_missing").empty? ? "-" : domain.fetch("gate_tokens_missing").map { |item| "#{item.fetch("file")}:#{item.fetch("token")}" }.join(", ")
    lines << "| #{domain.fetch("mainline_package")} | #{domain.fetch("package")} | #{domain.fetch("name")} | #{domain.fetch("status")} | #{missing_fixtures} | #{missing_gates} |"
  end

  lines << ""
  lines << "## First-Wave Dispatch"
  lines << ""
  lines << "| Mainline | Branch | Status | Minimal mergeable outcome |"
  lines << "| --- | --- | --- | --- |"
  report.fetch("mainline_first_wave").each do |entry|
    lines << "| #{entry.fetch("mainline_package")} | `#{entry.fetch("suggested_branch")}` | #{entry.fetch("status")} | #{entry.fetch("minimal_mergeable_outcome")} |"
  end

  lines << ""
  lines << report.fetch("desktop_safe_summary")
  lines.join("\n")
end

options = parse_options
report = build_report(options.fetch(:root))

case options.fetch(:format)
when "json"
  puts JSON.pretty_generate(report)
when "markdown"
  puts render_markdown(report)
end

exit(report.fetch("orphan_preview_methods_detected") ? 1 : 0)
