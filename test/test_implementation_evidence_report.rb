#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/implementation_evidence_report.rb")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json")
assert(status.success?, "implementation evidence report JSON must exit successfully: #{stderr}")
report = JSON.parse(stdout)

assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "implementation evidence report must expose the current version")
assert(report.fetch("schema_version") == "xnix.runtime.implementation_evidence_report.v1", "implementation evidence report must expose the schema")
assert(report.fetch("report_type") == "implementation-evidence-report", "implementation evidence report must identify its report type")
assert(report.fetch("runtime_owned"), "implementation evidence report must keep Runtime ownership explicit")
assert(report.fetch("go_runtime_backed"), "implementation evidence report must be Go Runtime backed")
assert(!report.fetch("kde_policy_owner"), "implementation evidence report must not make KDE the policy owner")
assert(report.fetch("source").include?("mainline-implementation-plan"), "implementation evidence report must include the mainline implementation plan in its source")
assert(report.fetch("source").include?("windows-compatibility-workstreams"), "implementation evidence report must include the Windows compatibility workstreams in its source")
assert(report.fetch("mainline_document") == "docs/claude-code-mainline-implementation-plan.md", "implementation evidence report must expose the mainline document path")
assert(report.fetch("mainline_plan_present"), "implementation evidence report must verify that the mainline document exists")
assert(report.fetch("windows_workstream_document") == "docs/claude-code-windows-compatibility-workstreams.md", "implementation evidence report must expose the Windows workstream document path")
assert(report.fetch("windows_workstream_board_present"), "implementation evidence report must verify that the Windows workstream document exists")
assert(report.fetch("mainline_package_count") == 9, "implementation evidence report must count mainline packages")
assert(report.fetch("mainline_first_wave").map { |entry| entry.fetch("mainline_package") } == %w[M1 M2 M3 M8], "implementation evidence report must expose the recommended first-wave mainline package order")
assert(report.fetch("windows_compatibility_first_wave").map { |entry| entry.fetch("workstream") } == %w[CW1 CW2 CW3 CW10], "implementation evidence report must expose the KDE-first Windows compatibility workstream order")
assert(report.fetch("windows_compatibility_first_wave").map { |entry| entry.fetch("mainline_package") } == %w[M1 M2 M3 M8], "implementation evidence report must map Windows workstreams to the mainline first wave")
assert(report.fetch("next_dispatch_packages") == report.fetch("mainline_first_wave"), "implementation evidence report must use the first wave as the next dispatch set")
assert(report.fetch("next_dispatch_summary").include?("M1, M2, M3, and M8"), "implementation evidence report must summarize the next dispatch set")
assert(report.fetch("windows_compatibility_next_dispatch_summary").include?("CW1, CW2, CW3, and CW10"), "implementation evidence report must summarize the Windows compatibility workstream dispatch set")
assert(report.fetch("domain_status_order") == %w[missing contract-only fixture-implemented state-root-implemented smoke-owned production-gated], "implementation evidence report must expose stable status order")
assert(report.fetch("domains").length == 9, "implementation evidence report must cover the empty-domain packages")
assert(report.fetch("counts").fetch("total") == 9, "implementation evidence report must count domains")
assert(report.fetch("read_only_method_count") == 61, "implementation evidence report must count Runtime read-only methods")
assert(report.fetch("orphan_read_methods").empty?, "implementation evidence report must not find orphan read methods")
assert(!report.fetch("orphan_preview_methods_detected"), "implementation evidence report must not detect orphan preview methods")
assert(report.fetch("write_methods") == %w[InstallRecipe Launch CreateSnapshot RestoreSnapshot], "implementation evidence report must list gated write methods")
assert(!report.fetch("write_methods_supported"), "implementation evidence report must not support write methods")
assert(!report.fetch("write_method_dispatch_enabled"), "implementation evidence report must not enable write dispatch")
assert(!report.fetch("network_required"), "implementation evidence report must not require network")
assert(!report.fetch("host_root_modified"), "implementation evidence report must not mutate the host root")
assert(!report.fetch("privileged_container_required"), "implementation evidence report must not require privileged containers")
assert(!report.fetch("backend_launch_enabled"), "implementation evidence report must not enable backend launch")
assert(!report.fetch("production_ready"), "implementation evidence report must keep production readiness gated")
assert(report.fetch("highest_evidence_status") == "smoke-owned", "implementation evidence report must recognize smoke-owned evidence as the highest current implementation depth")
assert(report.fetch("counts").fetch("production_gate_evidence") == 9, "implementation evidence report must count production gate evidence separately")

domains = report.fetch("domains").to_h { |domain| [domain.fetch("id"), domain] }
expected_domains = %w[
  runtime-owner-service
  recipe-artifact-trust-pipeline
  environment-lifecycle-state
  portal-snapshot-control-plane
  kde-activation-shell-materialization
  execution-transaction-ledger
  diagnostics-repair-ai-boundary
  atomic-kde-image-qemu-acceptance
  developer-verification-harness
]
assert(domains.keys == expected_domains, "implementation evidence report must expose stable domain ids")

assert(domains.fetch("runtime-owner-service").fetch("status") == "smoke-owned", "Runtime owner service must show smoke-owned evidence")
assert(domains.fetch("runtime-owner-service").fetch("mainline_package") == "M1", "Runtime owner service must map to M1")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/service.go"), "Runtime owner service must include in-process service boundary evidence")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/service_test.go"), "Runtime owner service must include in-process service boundary tests")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/lifecycle.go"), "Runtime owner service must include lifecycle evidence")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/smoke_batch.go"), "Runtime owner service must include smoke batch evidence")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/kde_test_launch_materialization_fanout_owner_smoke_coverage.go"), "Runtime owner service must include materialization fan-out owner smoke coverage evidence")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/restricted_smoke_receipt_fanout_owner_smoke_coverage.go"), "Runtime owner service must include restricted owner smoke fan-out smoke coverage evidence")
assert(domains.fetch("runtime-owner-service").fetch("fixture_files_present").include?("internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage.go"), "Runtime owner service must include redacted adapter profile owner smoke coverage evidence")
assert(domains.fetch("runtime-owner-service").fetch("smoke_files_present").include?("runtime/dbus/xnix_compatd_smoke.c"), "Runtime owner service must include C D-Bus bridge smoke evidence")
assert(domains.fetch("runtime-owner-service").fetch("smoke_files_present").include?("runtime/dbus/xnix_compatd_runtime_models.inc"), "Runtime owner service must include C D-Bus bridge model evidence")
assert(domains.fetch("runtime-owner-service").fetch("smoke_files_present").include?("runtime/dbus/xnix_compatd_kde_center.inc"), "Runtime owner service must include KDE Center C bridge model evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("in-process service-call boundary"), "Runtime owner service summary must mention the owner service-call boundary")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("D-Bus service-call bridge evidence"), "Runtime owner service summary must mention D-Bus service-call bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("full D-Bus read dispatch coverage"), "Runtime owner service summary must mention full read dispatch evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("service-call smoke-batch read/write evidence"), "Runtime owner service summary must mention service-call smoke batch evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("explicit materialization fan-out owner smoke coverage"), "Runtime owner service summary must mention materialization fan-out owner smoke coverage")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("explicit restricted owner smoke fan-out smoke coverage"), "Runtime owner service summary must mention restricted owner smoke fan-out smoke coverage")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("explicit redacted adapter profile owner smoke coverage"), "Runtime owner service summary must mention redacted adapter profile owner smoke coverage")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_smoke.c" &&
    entry.fetch("token") == "add_go_owner_dispatch_bridge_fields" &&
    entry.fetch("present")
}, "Runtime owner service must track generic C bridge helper evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/owner/service.go" &&
    entry.fetch("token") == "xnix.runtime.owner_service_call.v1" &&
    entry.fetch("present")
}, "Runtime owner service must track owner service-call schema evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_smoke.c" &&
    entry.fetch("token") == "go_owner_service_call_available" &&
    entry.fetch("present")
}, "Runtime owner service must track D-Bus service-call bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("second-ring KDE resource payloads"), "Runtime owner service summary must mention second-ring KDE resource bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("desktop activation payloads"), "Runtime owner service summary must mention desktop activation bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("install-input payloads"), "Runtime owner service summary must mention install input bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("backend lifecycle payloads"), "Runtime owner service summary must mention backend lifecycle bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("execution readiness payloads"), "Runtime owner service summary must mention execution readiness bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("AI safety payloads"), "Runtime owner service summary must mention AI safety bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("settings/review payloads"), "Runtime owner service summary must mention settings/review bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("Compatibility Center action payloads"), "Runtime owner service summary must mention Compatibility Center action bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("KDE Compatibility Center page payloads"), "Runtime owner service summary must mention KDE Compatibility Center page bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("foundation catalog/status payloads"), "Runtime owner service summary must mention foundation catalog/status bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("summary").include?("Runtime owner readiness payloads"), "Runtime owner service summary must mention Runtime owner readiness bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_smoke.c" &&
    entry.fetch("token") == "add_go_owner_dispatch_bridge_fields5" &&
    entry.fetch("present")
}, "Runtime owner service must track five-argument C bridge helper evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_smoke.c" &&
    entry.fetch("token") == "GetEngineCatalog" &&
    entry.fetch("present")
}, "Runtime owner service must track foundation engine catalog bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_runtime_models.inc" &&
    entry.fetch("token") == "GetRuntimeOwnerReadiness" &&
    entry.fetch("present")
}, "Runtime owner service must track owner readiness D-Bus bridge evidence")
assert(domains.fetch("runtime-owner-service").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "runtime/dbus/xnix_compatd_kde_center.inc" &&
    entry.fetch("token") == "GetKDECenterPageSectionDetail" &&
    entry.fetch("present")
}, "Runtime owner service must track KDE Center detail bridge evidence")
assert(domains.fetch("recipe-artifact-trust-pipeline").fetch("status") == "state-root-implemented", "recipe and artifact trust pipeline must show state-root implementation evidence")
assert(domains.fetch("recipe-artifact-trust-pipeline").fetch("mainline_package") == "M2", "recipe and artifact trust pipeline must map to M2")
assert(domains.fetch("environment-lifecycle-state").fetch("status") == "state-root-implemented", "environment lifecycle state must show state-root implementation evidence")
assert(domains.fetch("environment-lifecycle-state").fetch("mainline_package") == "M3", "environment lifecycle state must map to M3")
assert(domains.fetch("environment-lifecycle-state").fetch("fixture_files_present").include?("internal/runtime/appidentity/backend_manager.go"), "environment lifecycle state must include Go backend manager evidence")
assert(domains.fetch("environment-lifecycle-state").fetch("fixture_files_present").include?("internal/runtime/appidentity/backend_manager_test.go"), "environment lifecycle state must include Go backend manager tests")
assert(domains.fetch("environment-lifecycle-state").fetch("fixture_files_present").include?("internal/runtime/appidentity/backend_lifecycle_test.go"), "environment lifecycle state must include Go backend lifecycle state-root tests")
assert(domains.fetch("environment-lifecycle-state").fetch("fixture_files_present").include?("cmd/xnix-runtime-go/backend_group_cli_test.go"), "environment lifecycle state must include backend lifecycle CLI state-root tests")
assert(domains.fetch("environment-lifecycle-state").fetch("summary").include?("Wine, Proton, and Windows VM backends"), "environment lifecycle summary must mention internal Runtime-managed backend inventory")
assert(domains.fetch("environment-lifecycle-state").fetch("summary").include?("persist that inventory under an explicit state root"), "environment lifecycle summary must mention state-root backend inventory persistence")
assert(domains.fetch("environment-lifecycle-state").fetch("summary").include?("backend lifecycle records can be persisted and advanced"), "environment lifecycle summary must mention persisted lifecycle state transitions")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_manager.go" &&
    entry.fetch("token") == "windows-vm" &&
    entry.fetch("present")
}, "environment lifecycle state must track Windows VM backend manager evidence")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_manager.go" &&
    entry.fetch("token") == "BackendDetailsExposedToKDE" &&
    entry.fetch("present")
}, "environment lifecycle state must track KDE backend detail exposure safety")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_manager.go" &&
    entry.fetch("token") == "go-runtime-state-root-backend-manager" &&
    entry.fetch("present")
}, "environment lifecycle state must track backend manager state-root record evidence")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "cmd/xnix-runtime-go/main.go" &&
    entry.fetch("token") == "backend-manager-record" &&
    entry.fetch("present")
}, "environment lifecycle state must track backend manager record CLI coverage")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_lifecycle.go" &&
    entry.fetch("token") == "RecordBackendLifecycleState" &&
    entry.fetch("present")
}, "environment lifecycle state must track state-root lifecycle record implementation")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_lifecycle.go" &&
    entry.fetch("token") == "xnix.runtime.backend_lifecycle_record.v1" &&
    entry.fetch("present")
}, "environment lifecycle state must track lifecycle record schema evidence")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "cmd/xnix-runtime-go/main.go" &&
    entry.fetch("token") == "backend-lifecycle-record" &&
    entry.fetch("present")
}, "environment lifecycle state must track backend lifecycle record CLI coverage")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_lifecycle.go" &&
    entry.fetch("token") == "BackendLifecyclePreviewWithStateRoot" &&
    entry.fetch("present")
}, "environment lifecycle state must track the Go state-root lifecycle preview method")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/backend_lifecycle.go" &&
    entry.fetch("token") == "StateRootPathExposed" &&
    entry.fetch("present")
}, "environment lifecycle state must track state-root path exposure safety")
assert(domains.fetch("environment-lifecycle-state").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "cmd/xnix-runtime-go/backend_group_cli_test.go" &&
    entry.fetch("token") == "--state-root" &&
    entry.fetch("present")
}, "environment lifecycle state must track CLI state-root lifecycle coverage")
assert(domains.fetch("portal-snapshot-control-plane").fetch("status") == "state-root-implemented", "Portal and snapshot control plane must show state-root implementation evidence")
assert(domains.fetch("portal-snapshot-control-plane").fetch("mainline_package") == "M4", "Portal and snapshot control plane must map to M4")
assert(domains.fetch("portal-snapshot-control-plane").fetch("fixture_files_present").include?("internal/runtime/portal/ledger.go"), "Portal and snapshot control plane must include Portal request ledger implementation")
assert(domains.fetch("portal-snapshot-control-plane").fetch("state_files_present").include?("internal/runtime/portal/ledger.go"), "Portal and snapshot control plane must include state-root Portal ledger evidence")
assert(domains.fetch("portal-snapshot-control-plane").fetch("summary").include?("state-root permission request receipts"), "Portal and snapshot summary must mention state-root permission receipts")
assert(domains.fetch("portal-snapshot-control-plane").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/portal/ledger.go" &&
    entry.fetch("token") == "xnix.runtime.portal_request_record.v1" &&
    entry.fetch("present")
}, "Portal and snapshot control plane must track Portal request record schema evidence")
assert(domains.fetch("portal-snapshot-control-plane").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "cmd/xnix-runtime-go/main.go" &&
    entry.fetch("token") == "portal-request-record" &&
    entry.fetch("present")
}, "Portal and snapshot control plane must track Portal request record CLI coverage")
assert(domains.fetch("kde-activation-shell-materialization").fetch("status") == "smoke-owned", "KDE activation shell materialization must show smoke-owned evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("mainline_package") == "M5", "KDE activation shell materialization must map to M5")
assert(domains.fetch("kde-activation-shell-materialization").fetch("smoke_files_present").include?("scripts/kde_first_presence_smoke.rb"), "KDE activation shell materialization must include KDE-first presence smoke evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("smoke_files_present").include?("test/test_kde_first_presence_smoke_script.rb"), "KDE activation shell materialization must include KDE-first presence smoke tests")
assert(domains.fetch("kde-activation-shell-materialization").fetch("fixture_files_present").include?("internal/runtime/appidentity/desktop_safety_policy.go"), "KDE activation shell materialization must include Runtime desktop safety policy evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("fixture_files_present").include?("cmd/xnix-runtime-go/desktop_safety_policy_commands.go"), "KDE activation shell materialization must include Runtime desktop safety policy CLI evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("summary").include?("all seven KDE entry points"), "KDE activation summary must mention all seven KDE entry points")
assert(domains.fetch("kde-activation-shell-materialization").fetch("summary").include?("text, JSON, and Markdown evidence"), "KDE activation summary must mention reportable smoke evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "xnix.kde_first_presence_smoke.v1" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track KDE-first presence smoke schema evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "route_baseline" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track Runtime route baseline evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "settings_field_ids" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track user-facing settings field evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "forbidden_user_terms" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track forbidden user terminology evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/desktop_safety_policy.go" &&
    entry.fetch("token") == "xnix.runtime.desktop_safety_policy.v1" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track Runtime desktop safety policy schema evidence")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "desktop-safety-policy-preview" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track KDE smoke consumption of Runtime desktop safety policy")
assert(domains.fetch("kde-activation-shell-materialization").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "scripts/kde_first_presence_smoke.rb" &&
    entry.fetch("token") == "backend_launch_enabled" &&
    entry.fetch("present")
}, "KDE activation shell materialization must track backend launch safety evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("status") == "state-root-implemented", "execution transaction ledger must show state-root implementation evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("mainline_package") == "M6", "execution transaction ledger must map to M6")
assert(domains.fetch("execution-transaction-ledger").fetch("contract_files_present").include?("internal/runtime/appidentity/application_readiness.go"), "execution transaction ledger must include application readiness graph contract evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("fixture_files_present").include?("cmd/xnix-runtime-go/application_readiness_commands.go"), "execution transaction ledger must include application readiness CLI evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "internal/runtime/appidentity/application_readiness.go" &&
    entry.fetch("token") == "xnix.runtime.application_readiness.v1" &&
    entry.fetch("present")
}, "execution transaction ledger must track application readiness schema evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "cmd/xnix-runtime-go/application_readiness_commands.go" &&
    entry.fetch("token") == "application-readiness-preview" &&
    entry.fetch("present")
}, "execution transaction ledger must track application readiness CLI command evidence")
assert(domains.fetch("developer-verification-harness").fetch("status") == "fixture-implemented", "developer verification harness must show fixture implementation evidence")
assert(domains.fetch("developer-verification-harness").fetch("mainline_package") == "M8", "developer verification harness must map to M8")
assert(domains.fetch("atomic-kde-image-qemu-acceptance").fetch("mainline_package") == "M9", "atomic KDE image and QEMU acceptance must map to M9")
assert(domains.fetch("atomic-kde-image-qemu-acceptance").fetch("fixture_files_present").include?("lib/xnix/full_smoke_report.rb"), "atomic KDE image and QEMU acceptance must include full smoke report evidence")
assert(domains.fetch("atomic-kde-image-qemu-acceptance").fetch("smoke_files_present").include?("test/test_full_smoke_report.rb"), "atomic KDE image and QEMU acceptance must test full smoke reports")
assert(domains.fetch("atomic-kde-image-qemu-acceptance").fetch("gate_tokens").any? { |entry|
  entry.fetch("file") == "lib/xnix/full_smoke_report.rb" &&
    entry.fetch("token") == "qemu_network_restricted" &&
    entry.fetch("present")
}, "atomic KDE image and QEMU acceptance must track QEMU network restrictions")
assert(domains.values.all? { |domain| domain.fetch("mainline_document") == "docs/claude-code-mainline-implementation-plan.md" }, "all implementation domains must link to the mainline document")
assert(domains.values.all? { |domain| domain.fetch("production_gate_evidence") }, "all implementation domains must expose production gate evidence")
assert(domains.values.all? { |domain| domain.fetch("contract_files_missing").empty? }, "all implementation domains must have their contract files")
assert(domains.values.all? { |domain| domain.fetch("fixture_files_missing").empty? }, "all implementation domains must have their fixture files")
assert(domains.values.all? { |domain| domain.fetch("gate_tokens_missing").empty? }, "all implementation domains must have their gate tokens")
assert(domains.values.all? { |domain| !domain.fetch("host_root_modified") }, "domains must not mutate host root")
assert(domains.values.all? { |domain| !domain.fetch("network_required") }, "domains must not require network")
assert(domains.values.all? { |domain| !domain.fetch("privileged_container_required") }, "domains must not require privileged containers")
assert(domains.values.all? { |domain| !domain.fetch("backend_launch_enabled") }, "domains must not enable backend launch")

first_wave = report.fetch("mainline_first_wave")
assert(first_wave.fetch(0).fetch("suggested_branch") == "codex/runtime-owner-read-service", "first-wave M1 must include its suggested branch")
assert(first_wave.fetch(1).fetch("domain_id") == "recipe-artifact-trust-pipeline", "first-wave M2 must identify the recipe/artifact domain")
assert(first_wave.fetch(2).fetch("minimal_mergeable_outcome").include?("state-root lifecycle store"), "first-wave M3 must include a minimal mergeable outcome")
assert(first_wave.fetch(3).fetch("suggested_branch") == "codex/implementation-evidence-harness", "first-wave M8 must include its suggested branch")
assert(first_wave.all? { |entry| !entry.fetch("host_root_modified") }, "first-wave dispatch must not mutate host root")
assert(first_wave.all? { |entry| !entry.fetch("network_required") }, "first-wave dispatch must not require network")
assert(first_wave.all? { |entry| !entry.fetch("privileged_container_required") }, "first-wave dispatch must not require privileged containers")
assert(first_wave.all? { |entry| !entry.fetch("backend_launch_enabled") }, "first-wave dispatch must not enable backend launch")

windows_first_wave = report.fetch("windows_compatibility_first_wave")
assert(windows_first_wave.fetch(0).fetch("suggested_branch") == "codex/cw-runtime-owner-read-boundary", "Windows first-wave CW1 must include its suggested branch")
assert(windows_first_wave.fetch(1).fetch("domain_id") == "recipe-artifact-trust-pipeline", "Windows first-wave CW2 must identify the recipe/artifact domain")
assert(windows_first_wave.fetch(2).fetch("minimal_mergeable_outcome").include?("state-root lifecycle records"), "Windows first-wave CW3 must include a minimal mergeable outcome")
assert(windows_first_wave.fetch(3).fetch("workstream") == "CW10", "Windows first-wave must include CW10 as the evidence harness")
assert(windows_first_wave.all? { |entry| !entry.fetch("host_root_modified") }, "Windows first-wave dispatch must not mutate host root")
assert(windows_first_wave.all? { |entry| !entry.fetch("network_required") }, "Windows first-wave dispatch must not require network")
assert(windows_first_wave.all? { |entry| !entry.fetch("privileged_container_required") }, "Windows first-wave dispatch must not require privileged containers")
assert(windows_first_wave.all? { |entry| !entry.fetch("backend_launch_enabled") }, "Windows first-wave dispatch must not enable backend launch")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "markdown")
assert(status.success?, "implementation evidence report markdown must exit successfully: #{stderr}")
assert(stdout.include?("# Implementation Evidence Report"), "implementation evidence report markdown must include a title")
assert(stdout.include?("- Mainline document: docs/claude-code-mainline-implementation-plan.md"), "implementation evidence report markdown must include the mainline document")
assert(stdout.include?("- Windows workstream document: docs/claude-code-windows-compatibility-workstreams.md"), "implementation evidence report markdown must include the Windows workstream document")
assert(stdout.include?("- Next dispatch: First-wave mainline dispatch recommends M1, M2, M3, and M8"), "implementation evidence report markdown must include the next dispatch summary")
assert(stdout.include?("- Windows compatibility dispatch: KDE-first Windows compatibility dispatch recommends CW1, CW2, CW3, and CW10"), "implementation evidence report markdown must include the Windows compatibility dispatch summary")
assert(stdout.include?("| M1 | P1 | Runtime owner service | smoke-owned |"), "implementation evidence report markdown must include mainline domain rows")
assert(stdout.include?("| M9 | P8 | Atomic KDE image and QEMU acceptance | smoke-owned |"), "implementation evidence report markdown must show the image package as M9")
assert(stdout.include?("## First-Wave Dispatch"), "implementation evidence report markdown must include a first-wave dispatch section")
assert(stdout.include?("| M1 | `codex/runtime-owner-read-service` | smoke-owned |"), "implementation evidence report markdown must include the M1 dispatch branch")
assert(stdout.include?("| M8 | `codex/implementation-evidence-harness` | fixture-implemented |"), "implementation evidence report markdown must include the M8 dispatch branch")
assert(stdout.include?("## Windows Compatibility First-Wave Workstreams"), "implementation evidence report markdown must include the Windows compatibility first-wave section")
assert(stdout.include?("| CW1 | M1 | `codex/cw-runtime-owner-read-boundary` | smoke-owned |"), "implementation evidence report markdown must include the CW1 dispatch branch")
assert(stdout.include?("| CW10 | M8 | `codex/cw-evidence-drift-harness` | fixture-implemented |"), "implementation evidence report markdown must include the CW10 dispatch branch")
assert(stdout.include?("Implementation evidence spans 9 domains; production readiness remains gated."), "implementation evidence report markdown must include a safe summary")

puts "PASS: implementation evidence report tests"
