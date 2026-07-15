#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath
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
    ],
    state_files: [],
    smoke_files: %w[
      scripts/runtime_owner_candidate_smoke.rb
      scripts/dbus_session_smoke.rb
      test/test_runtime_owner_candidate_smoke_script.rb
    ],
    gate_tokens: {
      "cmd/xnix-runtime-owner/main.go" => %w[smoke-owner deny-write],
      "internal/runtime/appidentity/runtime_owner_readiness.go" => %w[ProductionOwnerEnabled ProductionBusClaimed WriteMethodsEnabled]
    },
    summary: "The Go owner candidate has read dispatch and constrained smoke evidence; production ownership remains gated."
  },
  {
    id: "recipe-artifact-trust-pipeline",
    name: "Recipe and artifact trust pipeline",
    package: "P2",
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
    ],
    state_files: [],
    smoke_files: [],
    gate_tokens: {
      "internal/runtime/diagnostics/ai.go" => %w[DisabledProvider FakeProvider],
      "internal/runtime/appidentity/ai_diagnostics.go" => %w[AIProviderCalled AutoExecutionAllowed]
    },
    summary: "Fixture diagnostics and fake AI provider boundaries exist; real provider calls and auto-repair remain disabled."
  },
  {
    id: "atomic-kde-image-qemu-acceptance",
    name: "Atomic KDE image and QEMU acceptance",
    package: "P8",
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

def build_report(root)
  version = read_project_file(root, "VERSION").strip
  domains = DOMAIN_DEFINITIONS.map { |definition| domain_report(root, definition) }
  read_methods = dbus_contract_methods(root) - WRITE_METHODS
  route_methods = owner_route_methods(root)
  orphan_read_methods = read_methods - route_methods
  counts = counts_for(domains)
  highest_rank = domains.map { |domain| domain.fetch("rank") }.max || 0

  {
    "version" => version,
    "schema_version" => "xnix.runtime.implementation_evidence_report.v1",
    "report_type" => "implementation-evidence-report",
    "source" => "filesystem+runtime-contract+owner-route-manifest+empty-domain-packages",
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
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
    "- Orphan read methods detected: #{report.fetch("orphan_preview_methods_detected")}",
    "- Production ready: #{report.fetch("production_ready")}",
    "- Host root modified: #{report.fetch("host_root_modified")}",
    "",
    "| Package | Domain | Status | Missing fixtures | Missing gates |",
    "| --- | --- | --- | --- | --- |"
  ]

  report.fetch("domains").each do |domain|
    missing_fixtures = domain.fetch("fixture_files_missing").empty? ? "-" : domain.fetch("fixture_files_missing").join(", ")
    missing_gates = domain.fetch("gate_tokens_missing").empty? ? "-" : domain.fetch("gate_tokens_missing").map { |item| "#{item.fetch("file")}:#{item.fetch("token")}" }.join(", ")
    lines << "| #{domain.fetch("package")} | #{domain.fetch("name")} | #{domain.fetch("status")} | #{missing_fixtures} | #{missing_gates} |"
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
