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
script = project_root.join("scripts/runtime_contract_drift_report.rb")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json")
assert(status.success?, "Runtime contract drift report JSON must exit successfully: #{stderr}")
report = JSON.parse(stdout)

assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "Runtime contract drift report must expose the current version")
assert(report.fetch("schema_version") == "xnix.runtime.contract_drift_report.v1", "Runtime contract drift report must expose the schema")
assert(report.fetch("report_type") == "runtime-contract-drift-report", "Runtime contract drift report must identify its report type")
assert(report.fetch("runtime_owned"), "Runtime contract drift report must keep Runtime ownership explicit")
assert(report.fetch("go_runtime_backed"), "Runtime contract drift report must be Go Runtime backed")
assert(!report.fetch("kde_policy_owner"), "Runtime contract drift report must not make KDE the policy owner")
assert(report.fetch("read_only_method_count") == 61, "Runtime contract drift report must count read-only methods")
assert(report.fetch("read_only_methods").include?("GetRuntimeMethodParityManifest"), "Runtime contract drift report must include method parity")
assert(report.fetch("read_only_methods").include?("GetRuntimeWriteGate"), "Runtime contract drift report must include write gate reads")
assert(report.fetch("owner_read_dispatch_method_count") >= 61, "Runtime contract drift report must count owner read dispatch methods")
assert(report.fetch("owner_smoke_batch_record_count") >= 61, "Runtime contract drift report must count owner smoke batch records")
assert(report.fetch("owner_session_bus_smoke_step_count") >= 66, "Runtime contract drift report must count owner session-bus smoke steps")
assert(report.fetch("owner_local_methods").include?("GetRuntimeOwnerReadiness"), "Runtime contract drift report must include owner-local readiness")
assert(report.fetch("owner_local_methods").include?("GetKDENotificationDigestPreview"), "Runtime contract drift report must include the owner-local KDE notification digest")
assert(report.fetch("owner_local_methods").include?("GetSignedRecipeVerificationPreview"), "Runtime contract drift report must include owner-local signed recipe verification")
assert(report.fetch("owner_local_methods").include?("GetRestrictedProductSmokePacketPreview"), "Runtime contract drift report must include the owner-local restricted smoke packet")
assert(report.fetch("owner_local_methods").include?("GetBackendAdapterProfileAudit"), "Runtime contract drift report must include the owner-local redacted adapter profile audit")
assert(report.fetch("owner_local_methods").include?("GetRestrictedOwnerSmokeReceiptLookupPreview"), "Runtime contract drift report must include the owner-local restricted owner smoke receipt lookup")
assert(report.fetch("owner_local_methods").include?("GetRestrictedOwnerSmokeReceiptFanOut"), "Runtime contract drift report must include the owner-local restricted owner smoke receipt fan-out")
assert(report.fetch("owner_local_methods").include?("GetKDETestLaunchMaterializationReceiptLookupPreview"), "Runtime contract drift report must include the owner-local KDE materialization receipt lookup")
assert(report.fetch("owner_local_methods").include?("GetKDETestLaunchMaterializationFanOut"), "Runtime contract drift report must include the owner-local KDE materialization fan-out route")
assert(report.fetch("write_methods") == %w[InstallRecipe Launch CreateSnapshot RestoreSnapshot], "Runtime contract drift report must list gated write methods")
assert(report.fetch("desktop_action_methods") == %w[ShowRuntimeControlledLaunch], "Runtime contract drift report must list controlled desktop action methods")
assert(!report.fetch("write_methods_supported"), "Runtime contract drift report must not support write methods")
assert(!report.fetch("write_method_dispatch_enabled"), "Runtime contract drift report must not enable write dispatch")
assert(!report.fetch("drift_detected"), "Runtime contract drift report must not detect drift")
assert(!report.fetch("network_required"), "Runtime contract drift report must not require network")
assert(!report.fetch("host_root_modified"), "Runtime contract drift report must not mutate the host root")
assert(!report.fetch("privileged_container_required"), "Runtime contract drift report must not require privileged containers")
assert(!report.fetch("backend_details_exposed"), "Runtime contract drift report must not expose backend details")
assert(report.fetch("counts") == { "total" => 16, "passed" => 16, "failed" => 0 }, "Runtime contract drift report must count checks")

check_ids = report.fetch("checks").map { |check| check.fetch("id") }
expected_checks = %w[
  parity-read-methods
  owner-route-go-methods
  dbus-client-method-map
  owner-read-dispatch-local-methods
  owner-read-dispatch-all-read-methods
  owner-read-dispatch-contract-subset
  owner-smoke-batch-source
  owner-session-bus-smoke-source
  go-owner-smoke-bridge
  session-smoke-service-call-envelope
  desktop-action-methods
  runtime-dispatch
  dbus-client-definitions
  smoke-adapter
  session-smoke
  go-cli-route-commands
]
assert(check_ids == expected_checks, "Runtime contract drift report must expose stable checks")
assert(report.fetch("checks").all? { |check| check.fetch("status") == "pass" }, "Runtime contract drift report checks must pass")
assert(report.fetch("checks").all? { |check| check.fetch("missing_methods").empty? }, "Runtime contract drift report checks must not miss methods")
assert(report.fetch("checks").all? { |check| check.fetch("extra_methods").empty? }, "Runtime contract drift report checks must not report extra methods")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "markdown")
assert(status.success?, "Runtime contract drift report markdown must exit successfully: #{stderr}")
assert(stdout.include?("# Runtime Contract Drift Report"), "Runtime contract drift report markdown must include a title")
assert(stdout.include?("| parity-read-methods | pass |"), "Runtime contract drift report markdown must include check rows")
assert(stdout.include?("Runtime read-only contract drift checks pass."), "Runtime contract drift report markdown must include a safe summary")

puts "PASS: Runtime contract drift report tests"
