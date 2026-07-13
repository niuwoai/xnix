#!/usr/bin/env ruby
# frozen_string_literal: true

require "pathname"
require_relative "../lib/xnix/compatibility/dbus_runtime_client"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

Status = Struct.new(:success?)

class FakeCapture
  attr_reader :commands

  def initialize
    @commands = []
  end

  def call(*command)
    @commands << command
    method = command.fetch(command.index("--method") + 1)

    case method
    when "org.xnix.Compatibility1.ListApplications"
      [
        "([{'id': <'org.xnix.sample.notepad'>, 'name': <'Sample Notepad'>, 'icon': <'accessories-text-editor'>, 'mode': <'automatic'>, 'supported_extensions': <['.txt', '.log']>}],)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetDiagnostics"
      [
        "({'application_id': <'org.xnix.sample.notepad'>, 'status': <'known'>, 'runtime_mode': <'automatic'>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetEngineCatalog"
      [
        "({'catalog_type': <'compatibility-engine'>, 'runtime_policy_owner': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRunPlan"
      [
        "({'plan_type': <'compatibility-run'>, 'strategy': <'automatic-managed'>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetApplicationStateRoot"
      [
        "({'root_type': <'compatibility-application-state-root'>, 'allocation_state': <'planned'>, 'snapshot_eligible': <true>, 'user_documents_included': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityPackageSource"
      [
        "({'source_type': <'compatibility-package-source'>, 'source_selection_state': <'planned'>, 'package_source_ready': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityAcquisitionPreflight"
      [
        "({'preflight_type': <'compatibility-acquisition-preflight'>, 'preflight_state': <'planned'>, 'acquisition_ready': <false>, 'download_enabled': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityArtifactManifest"
      [
        "({'manifest_type': <'compatibility-artifact-manifest'>, 'manifest_state': <'planned'>, 'manifest_ready': <false>, 'signature_verified': <false>, 'download_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetCompatibilityInstallPlan"
      [
        "({'plan_type': <'compatibility-install-plan'>, 'install_state': <'planned'>, 'install_ready': <false>, 'desktop_activation_ready': <false>, 'download_enabled': <false>, 'install_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetBackendBinding"
      [
        "({'binding_type': <'compatibility-backend-binding'>, 'selected_strategy': <'automatic-managed'>, 'managed_binding_ready': <false>, 'launch_enabled': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRepairPlan"
      [
        "({'plan_type': <'compatibility-repair'>, 'issue': <'engine-binding-pending'>, 'snapshot_required': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestPlan"
      [
        "({'plan_type': <'compatibility-test'>, 'test_type': <'preflight'>, 'runtime_owned': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetTestResult"
      [
        "({'result_type': <'compatibility-test-result'>, 'test_type': <'preflight'>, 'overall_status': <'pending'>, 'safe_for_ai_diagnostics': <true>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticInput"
      [
        "({'input_type': <'ai-diagnostic-input'>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIDiagnosticRecommendation"
      [
        "({'recommendation_type': <'ai-diagnostic-recommendation'>, 'safe_for_ai_diagnostics': <true>, 'ai_provider_called': <false>, 'network_required': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetAIRepairApprovalGate"
      [
        "({'gate_type': <'ai-repair-approval-gate'>, 'gate_decision': <'blocked-until-approval'>, 'auto_execution_allowed': <false>, 'repair_executed': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetSnapshotPlan"
      [
        "({'plan_type': <'compatibility-snapshot'>, 'reason': <'before-repair'>, 'enabled_by_default': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetPortalAccessPolicy"
      [
        "({'policy_type': <'portal-access'>, 'operation': <'file-open'>, 'portal_required': <true>},)\n",
        "",
        Status.new(true)
      ]
    when "org.xnix.Compatibility1.GetRuntimeServiceBinding"
      [
        "({'binding_type': <'runtime-service-binding'>, 'activation_binding_ready': <true>, 'live_dbus_owner_ready': <false>, 'backend_details_exposed': <false>},)\n",
        "",
        Status.new(true)
      ]
    else
      ["", "unexpected method", Status.new(false)]
    end
  end
end

capture = FakeCapture.new
client = Xnix::Compatibility::DBusRuntimeClient.new(capture: capture)

source = client.source_metadata
assert(source["kind"] == "runtime-dbus-session", "D-Bus client must declare a session bus source")
assert(source["bus_name"] == "org.xnix.Compatibility1", "D-Bus client must use the Runtime bus name")

applications = client.list_applications
assert(applications.length == 1, "D-Bus client must parse the application list")
assert(applications.first["id"] == "org.xnix.sample.notepad", "D-Bus client must parse string fields")
assert(applications.first["mode"] == "automatic", "D-Bus client must parse application mode")
assert(applications.first["supported_extensions"] == [".txt", ".log"], "D-Bus client must parse string array fields")

diagnostics = client.diagnostics("org.xnix.sample.notepad")
assert(diagnostics["status"] == "known", "D-Bus client must parse diagnostics status")

engine_catalog = client.engine_catalog
assert(engine_catalog["catalog_type"] == "compatibility-engine", "D-Bus client must parse engine catalog")
assert(engine_catalog["runtime_policy_owner"], "D-Bus client must parse boolean true values")
assert(!engine_catalog["backend_details_exposed"], "D-Bus client must parse boolean false values")

run_plan = client.run_plan("org.xnix.sample.notepad")
assert(run_plan["plan_type"] == "compatibility-run", "D-Bus client must parse run plans")
assert(!run_plan["backend_details_exposed"], "D-Bus client must parse run plan booleans")

state_root = client.state_root("org.xnix.sample.notepad")
assert(state_root["root_type"] == "compatibility-application-state-root", "D-Bus client must parse application state roots")
assert(state_root["snapshot_eligible"], "D-Bus client must parse state root booleans")
assert(!state_root["user_documents_included"], "D-Bus client must parse state root exclusions")

package_source = client.package_source("org.xnix.sample.notepad")
assert(package_source["source_type"] == "compatibility-package-source", "D-Bus client must parse compatibility package sources")
assert(package_source["source_selection_state"] == "planned", "D-Bus client must parse package source state")
assert(!package_source["package_source_ready"], "D-Bus client must parse package source readiness")
assert(!package_source["install_enabled"], "D-Bus client must parse package install status")

acquisition_preflight = client.acquisition_preflight("org.xnix.sample.notepad")
assert(acquisition_preflight["preflight_type"] == "compatibility-acquisition-preflight", "D-Bus client must parse compatibility acquisition preflight")
assert(acquisition_preflight["preflight_state"] == "planned", "D-Bus client must parse acquisition preflight state")
assert(!acquisition_preflight["acquisition_ready"], "D-Bus client must parse acquisition readiness")
assert(!acquisition_preflight["download_enabled"], "D-Bus client must parse download status")

artifact_manifest = client.artifact_manifest("org.xnix.sample.notepad")
assert(artifact_manifest["manifest_type"] == "compatibility-artifact-manifest", "D-Bus client must parse compatibility artifact manifests")
assert(artifact_manifest["manifest_state"] == "planned", "D-Bus client must parse artifact manifest state")
assert(!artifact_manifest["manifest_ready"], "D-Bus client must parse artifact manifest readiness")
assert(!artifact_manifest["signature_verified"], "D-Bus client must parse artifact manifest signature status")
assert(!artifact_manifest["download_enabled"], "D-Bus client must parse artifact download status")

install_plan = client.install_plan("org.xnix.sample.notepad")
assert(install_plan["plan_type"] == "compatibility-install-plan", "D-Bus client must parse compatibility install plans")
assert(install_plan["install_state"] == "planned", "D-Bus client must parse install plan state")
assert(!install_plan["install_ready"], "D-Bus client must parse install readiness")
assert(!install_plan["desktop_activation_ready"], "D-Bus client must parse desktop activation readiness")
assert(!install_plan["download_enabled"], "D-Bus client must parse install plan download status")
assert(!install_plan["install_enabled"], "D-Bus client must parse install plan install status")

backend_binding = client.backend_binding("org.xnix.sample.notepad")
assert(backend_binding["binding_type"] == "compatibility-backend-binding", "D-Bus client must parse backend bindings")
assert(!backend_binding["managed_binding_ready"], "D-Bus client must parse backend binding readiness")
assert(!backend_binding["launch_enabled"], "D-Bus client must parse backend launch status")

repair_plan = client.repair_plan("org.xnix.sample.notepad", "engine-binding-pending")
assert(repair_plan["plan_type"] == "compatibility-repair", "D-Bus client must parse repair plans")
assert(repair_plan["snapshot_required"], "D-Bus client must parse repair plan booleans")

test_plan = client.test_plan("org.xnix.sample.notepad")
assert(test_plan["plan_type"] == "compatibility-test", "D-Bus client must parse test plans")
assert(test_plan["runtime_owned"], "D-Bus client must parse test plan booleans")
assert(!test_plan["backend_details_exposed"], "D-Bus client must parse test plan false booleans")

test_result = client.test_result("org.xnix.sample.notepad")
assert(test_result["result_type"] == "compatibility-test-result", "D-Bus client must parse test results")
assert(test_result["safe_for_ai_diagnostics"], "D-Bus client must parse test result booleans")
assert(!test_result["backend_details_exposed"], "D-Bus client must parse test result false booleans")

ai_input = client.ai_diagnostic_input("org.xnix.sample.notepad")
assert(ai_input["input_type"] == "ai-diagnostic-input", "D-Bus client must parse AI diagnostic inputs")
assert(ai_input["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic input booleans")
assert(!ai_input["ai_provider_called"], "D-Bus client must parse AI provider false booleans")
assert(!ai_input["network_required"], "D-Bus client must parse network false booleans")

ai_recommendation = client.ai_diagnostic_recommendation("org.xnix.sample.notepad")
assert(ai_recommendation["recommendation_type"] == "ai-diagnostic-recommendation", "D-Bus client must parse AI diagnostic recommendations")
assert(ai_recommendation["safe_for_ai_diagnostics"], "D-Bus client must parse AI diagnostic recommendation booleans")
assert(!ai_recommendation["ai_provider_called"], "D-Bus client must parse recommendation AI provider false booleans")
assert(!ai_recommendation["network_required"], "D-Bus client must parse recommendation network false booleans")

ai_gate = client.ai_repair_approval_gate("org.xnix.sample.notepad")
assert(ai_gate["gate_type"] == "ai-repair-approval-gate", "D-Bus client must parse AI repair approval gates")
assert(ai_gate["gate_decision"] == "blocked-until-approval", "D-Bus client must parse AI repair approval gate decisions")
assert(!ai_gate["auto_execution_allowed"], "D-Bus client must parse AI repair gate execution booleans")
assert(!ai_gate["repair_executed"], "D-Bus client must parse AI repair execution false booleans")

snapshot_plan = client.snapshot_plan("org.xnix.sample.notepad", "before-repair")
assert(snapshot_plan["plan_type"] == "compatibility-snapshot", "D-Bus client must parse snapshot plans")
assert(snapshot_plan["enabled_by_default"], "D-Bus client must parse snapshot plan booleans")

portal_policy = client.portal_access_policy("org.xnix.sample.notepad", "file-open")
assert(portal_policy["policy_type"] == "portal-access", "D-Bus client must parse Portal access policies")
assert(portal_policy["portal_required"], "D-Bus client must parse Portal policy booleans")

service_binding = client.runtime_service_binding
assert(service_binding["binding_type"] == "runtime-service-binding", "D-Bus client must parse Runtime service binding")
assert(service_binding["activation_binding_ready"], "D-Bus client must parse Runtime service binding readiness")
assert(!service_binding["live_dbus_owner_ready"], "D-Bus client must parse live D-Bus owner readiness")

assert(
  capture.commands.all? { |command| command.include?("--session") },
  "D-Bus client must use the session bus for KDE-facing reads"
)

puts "PASS: compatibility runtime D-Bus client unit tests"
