#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "optparse"
require "pathname"

WRITE_METHODS = %w[
  InstallRecipe
  Launch
  CreateSnapshot
  RestoreSnapshot
].freeze

OWNER_LOCAL_READ_METHODS = %w[
  GetRuntimeServiceBinding
  GetRuntimeLiveOwnerGate
  GetRuntimeOwnerProcess
  GetRuntimeOwnerSmokePlan
  GetRuntimeMethodParityManifest
  GetRuntimeOwnerRouteManifest
  GetRuntimeOwnerRecipeTrust
  GetRuntimeOwnerReadiness
  GetWindowsCompatibilityWorkstreamsPreview
  GetKDENotificationDigestPreview
  GetSignedRecipeVerificationPreview
  GetRestrictedProductSmokePacketPreview
  GetBackendAdapterProfileAudit
  GetRestrictedOwnerSmokeReceiptLookupPreview
  GetRestrictedOwnerSmokeReceiptFanOut
  GetRuntimeWriteGate
].freeze

PROJECT_ROOT = Pathname.new(__dir__).join("..").realpath

def parse_options
  options = {
    root: PROJECT_ROOT,
    format: "json"
  }

  OptionParser.new do |parser|
    parser.banner = "Usage: ruby scripts/runtime_contract_drift_report.rb [--root PATH] [--format json|markdown]"
    parser.on("--root PATH", "Project root to inspect") do |value|
      options[:root] = Pathname.new(value).expand_path
    end
    parser.on("--format FORMAT", "Output format: json or markdown") do |value|
      options[:format] = value
    end
  end.parse!

  unless %w[json markdown].include?(options[:format])
    abort "runtime_contract_drift_report supports --format json or --format markdown"
  end

  options
end

def read_project_file(root, relative_path)
  root.join(relative_path).read
rescue Errno::ENOENT
  ""
end

def read_sources(root, relative_paths)
  relative_paths.map { |path| read_project_file(root, path) }.join("\n")
end

def dbus_contract_methods(root)
  read_project_file(root, "runtime/dbus/org.xnix.Compatibility1.xml").scan(/<method name="([^"]+)"/).flatten.uniq
end

def go_string_slice(source, variable_name)
  body = source[/var #{Regexp.escape(variable_name)} = \[\]string\{(.*?)\n\}/m, 1].to_s
  body.scan(/"([^"]+)"/).flatten
end

def go_string_map_keys(source, variable_name)
  body = source[/var #{Regexp.escape(variable_name)} = map\[string\]string\{(.*?)\n\}/m, 1].to_s
  body.scan(/"([^"]+)"\s*:/).flatten
end

def go_string_map_values(source, variable_name)
  body = source[/var #{Regexp.escape(variable_name)} = map\[string\]string\{(.*?)\n\}/m, 1].to_s
  body.scan(/:\s*"([^"]+)"/).flatten
end

def owner_dispatch_methods(source)
  source.scan(/^\s*"([^"]+)":\s*func/).flatten.uniq
end

def compare_lists(id:, expected:, actual:, summary:)
  missing = expected - actual
  extra = actual - expected
  status = missing.empty? && extra.empty? ? "pass" : "fail"

  {
    "id" => id,
    "status" => status,
    "expected_count" => expected.length,
    "actual_count" => actual.length,
    "missing_methods" => missing,
    "extra_methods" => extra,
    "summary" => status == "pass" ? summary : "#{id} drift detected."
  }
end

def source_coverage_check(id:, expected:, source:, summary:, &covers)
  missing = expected.reject { |method| covers.call(source, method) }
  status = missing.empty? ? "pass" : "fail"

  {
    "id" => id,
    "status" => status,
    "expected_count" => expected.length,
    "actual_count" => expected.length - missing.length,
    "missing_methods" => missing,
    "extra_methods" => [],
    "summary" => status == "pass" ? summary : "#{id} is missing Runtime methods."
  }
end

def drift_counts(checks)
  {
    "total" => checks.length,
    "passed" => checks.count { |check| check.fetch("status") == "pass" },
    "failed" => checks.count { |check| check.fetch("status") == "fail" }
  }
end

def build_report(root)
  version = read_project_file(root, "VERSION").strip
  contract_methods = dbus_contract_methods(root)
  read_only_methods = contract_methods - WRITE_METHODS

  parity_source = read_project_file(root, "internal/runtime/appidentity/runtime_method_parity_manifest.go")
  route_source = read_project_file(root, "internal/runtime/appidentity/runtime_owner_route_manifest.go")
  owner_dispatch_source = read_project_file(root, "internal/runtime/owner/dispatch.go")
  owner_smoke_batch_source = read_project_file(root, "internal/runtime/owner/smoke_batch.go")
  owner_session_bus_source = read_sources(root, [
    "internal/runtime/owner/session_bus.go",
    "cmd/xnix-runtime-owner/main.go"
  ])
  go_cli_source = read_sources(root, [
    "cmd/xnix-runtime-go/main.go",
    "cmd/xnix-runtime-go/runtime_owner_commands.go"
  ])
  runtime_dispatch_source = read_project_file(root, "lib/xnix/compatibility/runtime_daemon.rb")
  dbus_client_source = read_project_file(root, "lib/xnix/compatibility/dbus_runtime_client.rb")
  smoke_adapter_source = read_sources(root, [
    "runtime/dbus/xnix_compatd_smoke.c",
    "runtime/dbus/xnix_compatd_introspection.inc",
    "runtime/dbus/xnix_compatd_kde_center.inc",
    "runtime/dbus/xnix_compatd_runtime_models.inc"
  ])
  session_smoke_source = read_project_file(root, "scripts/dbus_session_smoke.rb")

  parity_methods = go_string_slice(parity_source, "runtimeMethodParityReadOnlyMethods")
  client_method_map_keys = go_string_map_keys(parity_source, "runtimeMethodParityClientMethods")
  client_method_names = go_string_map_values(parity_source, "runtimeMethodParityClientMethods")
  route_methods = go_string_map_keys(route_source, "runtimeOwnerRouteGoCommands")
  route_commands = go_string_map_values(route_source, "runtimeOwnerRouteGoCommands")
  dispatch_methods = owner_dispatch_methods(owner_dispatch_source)
  owner_contract_methods = read_only_methods.grep(/^GetRuntime/)

  checks = [
    compare_lists(
      id: "parity-read-methods",
      expected: read_only_methods,
      actual: parity_methods,
      summary: "Go method parity manifest matches the D-Bus read-only contract."
    ),
    compare_lists(
      id: "owner-route-go-methods",
      expected: read_only_methods,
      actual: route_methods,
      summary: "Go owner route manifest maps every D-Bus read-only method."
    ),
    compare_lists(
      id: "dbus-client-method-map",
      expected: read_only_methods,
      actual: client_method_map_keys,
      summary: "D-Bus client method map covers every D-Bus read-only method."
    ),
    source_coverage_check(
      id: "owner-read-dispatch-local-methods",
      expected: OWNER_LOCAL_READ_METHODS,
      source: owner_dispatch_source,
      summary: "Owner read dispatch covers the local owner method group."
    ) { |source, method| source.include?("\"#{method}\"") },
    source_coverage_check(
      id: "owner-read-dispatch-all-read-methods",
      expected: read_only_methods,
      source: owner_dispatch_source,
      summary: "Owner read dispatch covers every D-Bus read-only method."
    ) { |source, method| source.include?("\"#{method}\"") },
    source_coverage_check(
      id: "owner-read-dispatch-contract-subset",
      expected: owner_contract_methods,
      source: owner_dispatch_source,
      summary: "Owner read dispatch covers every D-Bus owner read method."
    ) { |source, method| source.include?("\"#{method}\"") },
    source_coverage_check(
      id: "owner-smoke-batch-source",
      expected: ["SupportedReadDispatchMethods", "NewService", "service.Call", "runtime-owner-service-call", "xnix.runtime.owner_smoke_batch.v1", "restricted-session-owner-call-batch"],
      source: owner_smoke_batch_source,
      summary: "Owner smoke batch derives read/write coverage from the owner service-call boundary."
    ) { |source, token| source.include?(token) },
    source_coverage_check(
      id: "owner-session-bus-smoke-source",
      expected: [
        "NewSessionBusSmokeTranscript",
        "NewSmokeBatchRecords",
        "runtime-owner-session-bus-smoke-step",
        "xnix.runtime.owner_session_bus_smoke.v1",
        "restricted-private-session-bus-owner-smoke",
        "private-session-bus-smoke",
        "reject-unsupported-read",
        "org.xnix.Compatibility1.Error.UnsupportedMethod",
        "session-bus-smoke",
        "ProductionBusClaimed",
        "HostRootModified",
        "BackendDetailsExposed",
        "WriteMethodsEnabled"
      ],
      source: owner_session_bus_source,
      summary: "Owner private session-bus smoke wraps read dispatch, disabled writes, and unsupported-read rejection without production ownership."
    ) { |source, token| source.include?(token) },
    source_coverage_check(
      id: "go-owner-smoke-bridge",
      expected: [
        "go_owner_read_dispatch",
        "add_go_owner_dispatch_bridge_fields",
        "add_go_owner_dispatch_payload_fields",
        "add_go_owner_dispatch_bridge_fields5",
        "go_owner_dispatch_available",
        "go_owner_dispatch_schema",
        "go_owner_dispatch_json",
        "go_owner_service_call5",
        "go_owner_service_call_available",
        "go_owner_service_call_schema",
        "go_owner_service_call_request_type",
        "go_owner_service_call_json",
        "xnix.runtime.owner_service_call.v1",
        "runtime-owner-service-call",
        "go-runtime-owner-dispatch+c-smoke-bridge",
        "add_go_owner_dispatch_payload_fields(&builder,",
        "build_application(\"ListApplications\")",
        "build_application(\"GetApplication\")",
        "add_go_owner_dispatch_bridge_fields(&diagnostics, \"GetDiagnostics\"",
        "add_go_owner_dispatch_bridge_fields(&catalog, \"GetEngineCatalog\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetRunPlan\"",
        "add_go_owner_dispatch_bridge_fields(&manifest, \"GetDesktopActivationManifest\"",
        "add_go_owner_dispatch_bridge_fields2(&transaction, \"GetDesktopActivationTransactionPreview\"",
        "add_go_owner_dispatch_bridge_fields2(&status, \"GetDesktopActivationStatus\"",
        "add_go_owner_dispatch_bridge_fields(&status, \"GetKDEIntegrationStatus\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetKDEShellIntegrationPlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetKDEApplicationSurfacePlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetDesktopEntryPlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetDesktopIconPlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetTaskManagerIdentityPlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetDesktopResourceBridgePlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetKWinWindowRulePlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetFileAssociationPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetNotificationPlan\"",
        "add_go_owner_dispatch_bridge_fields(&status, \"GetTrayStatus\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetKRunnerQueryPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetPortalRequestPlan\"",
        "add_go_owner_dispatch_bridge_fields(&source, \"GetCompatibilityPackageSource\"",
        "add_go_owner_dispatch_bridge_fields(&preflight, \"GetCompatibilityAcquisitionPreflight\"",
        "add_go_owner_dispatch_bridge_fields(&manifest, \"GetCompatibilityArtifactManifest\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetCompatibilityInstallPlan\"",
        "add_go_owner_dispatch_bridge_fields(&binding, \"GetBackendBinding\"",
        "add_go_owner_dispatch_bridge_fields(&matrix, \"GetBackendCapabilityMatrix\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetBackendSelectionPlan\"",
        "add_go_owner_dispatch_bridge_fields(&lifecycle, \"GetBackendLifecycle\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetBackendEnvironmentPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetRepairPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetTestPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&result, \"GetTestResult\"",
        "add_go_owner_dispatch_bridge_fields(&readiness, \"GetExecutionReadiness\"",
        "add_go_owner_dispatch_bridge_fields(&intent, \"GetLaunchIntent\"",
        "add_go_owner_dispatch_bridge_fields3(&input, \"GetAIDiagnosticInput\"",
        "add_go_owner_dispatch_bridge_fields3(&recommendation, \"GetAIDiagnosticRecommendation\"",
        "add_go_owner_dispatch_bridge_fields3(&gate, \"GetAIRepairApprovalGate\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetSnapshotPlan\"",
        "add_go_owner_dispatch_bridge_fields2(&policy, \"GetPortalAccessPolicy\"",
        "add_go_owner_dispatch_bridge_fields(&state_root, \"GetApplicationStateRoot\"",
        "add_go_owner_dispatch_bridge_fields(&process, \"GetRuntimeOwnerProcess\"",
        "add_go_owner_dispatch_bridge_fields(&manifest, \"GetRuntimeOwnerRouteManifest\"",
        "add_go_owner_dispatch_bridge_fields(&trust, \"GetRuntimeOwnerRecipeTrust\"",
        "add_go_owner_dispatch_bridge_fields(&readiness, \"GetRuntimeOwnerReadiness\"",
        "add_go_owner_dispatch_bridge_fields(&settings, \"GetCompatibilitySettings\"",
        "add_go_owner_dispatch_bridge_fields5(&plan, \"GetCompatibilitySettingsChangePlan\"",
        "add_go_owner_dispatch_bridge_fields2(&plan, \"GetCompatibilityModeSwitchPlan\"",
        "add_go_owner_dispatch_bridge_fields(&plan, \"GetCompatibilityPermissionReviewPlan\"",
        "add_go_owner_dispatch_bridge_fields5(&plan, \"GetCompatibilityReviewFlowPlan\"",
        "add_go_owner_dispatch_bridge_fields(&queue, \"GetCompatibilityActionQueue\"",
        "add_go_owner_dispatch_bridge_fields3(&receipt, \"GetCompatibilityActionReviewReceipt\"",
        "add_go_owner_dispatch_bridge_fields(&summary, \"GetCompatibilityCenterSummary\"",
        "add_go_owner_dispatch_payload_fields2(&page, \"GetKDECenterPage\"",
        "add_go_owner_dispatch_payload_fields2(&model, \"GetKDECenterPageSections\"",
        "add_go_owner_dispatch_payload_fields3(&detail, \"GetKDECenterPageSectionDetail\""
      ],
      source: smoke_adapter_source,
      summary: "D-Bus smoke adapter exposes Go owner dispatch payload evidence for Runtime owner self-description, Runtime owner readiness reads, foundation catalog/status reads, KDE shell, desktop activation, seven-entry-point reads, second-ring KDE resource reads, install-input reads, backend lifecycle reads, execution readiness reads, AI safety reads, settings/review reads, Compatibility Center action reads, and KDE Compatibility Center page reads through the low-level C bridge."
    ) { |source, token| source.include?(token) },
    source_coverage_check(
      id: "session-smoke-service-call-envelope",
      expected: [
        "variant_string_field",
        "assert_go_owner_service_call_envelope",
        "go_owner_service_call_json",
        "xnix.runtime.owner_service_call.v1",
        "runtime-owner-service-call",
        "go-runtime-owner-in-process-service",
        '"method": "#{method_name}"',
        "\"call_type\": \"read-dispatch\"",
        "\"dispatch_ready\": true",
        "\"write_methods_enabled\": false",
        "\"kde_policy_owner\": false",
        "\"backend_details_exposed\": false",
        "\"host_root_modified\": false",
        "\"network_required\": false"
      ],
      source: session_smoke_source,
      summary: "Restricted session smoke extracts and verifies Go owner service-call envelopes for D-Bus read methods."
    ) { |source, token| source.include?(token) },
    source_coverage_check(
      id: "runtime-dispatch",
      expected: read_only_methods,
      source: runtime_dispatch_source,
      summary: "Ruby Runtime dispatch still covers the read-only contract during migration."
    ) { |source, method| source.include?("\"#{method}\"") },
    source_coverage_check(
      id: "dbus-client-definitions",
      expected: client_method_names,
      source: dbus_client_source,
      summary: "Ruby D-Bus client exposes every mapped read-only method."
    ) { |source, method| source.include?("def #{method}") },
    source_coverage_check(
      id: "smoke-adapter",
      expected: read_only_methods,
      source: smoke_adapter_source,
      summary: "Linux D-Bus smoke adapter covers every read-only method."
    ) { |source, method| source.include?(method) },
    source_coverage_check(
      id: "session-smoke",
      expected: read_only_methods,
      source: session_smoke_source,
      summary: "Restricted session smoke calls every read-only method."
    ) { |source, method| source.include?(method) },
    source_coverage_check(
      id: "go-cli-route-commands",
      expected: route_commands,
      source: go_cli_source,
      summary: "Go CLI exposes every command referenced by the owner route manifest."
    ) { |source, command| source.include?("\"#{command}\"") }
  ]
  counts = drift_counts(checks)

  {
    "version" => version,
    "schema_version" => "xnix.runtime.contract_drift_report.v1",
    "report_type" => "runtime-contract-drift-report",
    "source" => "dbus-contract+go-parity+go-owner-routes+owner-read-dispatch+owner-smoke-batch+owner-session-bus-smoke+runtime-dispatch+dbus-client+smoke-adapter+session-smoke",
    "runtime_owned" => true,
    "go_runtime_backed" => true,
    "kde_policy_owner" => false,
    "read_only_method_count" => read_only_methods.length,
    "read_only_methods" => read_only_methods,
    "owner_read_dispatch_method_count" => dispatch_methods.length,
    "owner_smoke_batch_record_count" => dispatch_methods.length + WRITE_METHODS.length,
    "owner_session_bus_smoke_step_count" => dispatch_methods.length + WRITE_METHODS.length + 5,
    "write_methods" => WRITE_METHODS,
    "write_methods_supported" => false,
    "write_method_dispatch_enabled" => false,
    "owner_local_methods" => OWNER_LOCAL_READ_METHODS,
    "checks" => checks,
    "counts" => counts,
    "drift_detected" => counts.fetch("failed").positive?,
    "network_required" => false,
    "host_root_modified" => false,
    "privileged_container_required" => false,
    "backend_details_exposed" => false,
    "desktop_safe_summary" => counts.fetch("failed").zero? ? "Runtime read-only contract drift checks pass." : "Runtime read-only contract drift checks failed."
  }
end

def render_markdown(report)
  lines = [
    "# Runtime Contract Drift Report",
    "",
    "- Version: #{report.fetch("version")}",
    "- Schema: #{report.fetch("schema_version")}",
    "- Read-only methods: #{report.fetch("read_only_method_count")}",
    "- Drift detected: #{report.fetch("drift_detected")}",
    "- Write methods supported: #{report.fetch("write_methods_supported")}",
    "- Host root modified: #{report.fetch("host_root_modified")}",
    "",
    "| Check | Status | Expected | Actual | Missing | Extra |",
    "| --- | --- | ---: | ---: | --- | --- |"
  ]

  report.fetch("checks").each do |check|
    missing = check.fetch("missing_methods").empty? ? "-" : check.fetch("missing_methods").join(", ")
    extra = check.fetch("extra_methods").empty? ? "-" : check.fetch("extra_methods").join(", ")
    lines << "| #{check.fetch("id")} | #{check.fetch("status")} | #{check.fetch("expected_count")} | #{check.fetch("actual_count")} | #{missing} | #{extra} |"
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

exit(report.fetch("drift_detected") ? 1 : 0)
