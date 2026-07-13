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
command = ["ruby", project_root.join("bin/xnix-compatd").to_s]

stdout, stderr, status = Open3.capture3(*command, "probe")
assert(status.success?, "runtime daemon probe must exit successfully: #{stderr}")
probe = JSON.parse(stdout)
assert(probe["version"] == "0.2.43", "runtime daemon probe must report the current version")
assert(probe["bus_name"] == "org.xnix.Compatibility1", "runtime daemon probe must keep the stable bus name")
assert(probe["capabilities"]["recipe_store"], "runtime daemon probe must expose recipe store capability")
assert(probe["capabilities"]["registry_backed_recipe_store"], "runtime daemon probe must expose registry-backed recipe loading")
assert(probe["capabilities"]["application_state_roots"], "runtime daemon probe must expose application state roots")
assert(probe["capabilities"]["compatibility_engine_catalog"], "runtime daemon probe must expose engine catalog capability")
assert(probe["capabilities"]["compatibility_package_sources"], "runtime daemon probe must expose package source capability")
assert(probe["capabilities"]["compatibility_backend_binding"], "runtime daemon probe must expose backend binding capability")
assert(probe["capabilities"]["compatibility_run_planning"], "runtime daemon probe must expose run planning capability")
assert(probe["capabilities"]["compatibility_repair_planning"], "runtime daemon probe must expose repair planning capability")
assert(probe["capabilities"]["compatibility_test_planning"], "runtime daemon probe must expose test planning capability")
assert(probe["capabilities"]["compatibility_test_results"], "runtime daemon probe must expose test result capability")
assert(probe["capabilities"]["ai_diagnostic_inputs"], "runtime daemon probe must expose AI diagnostic input capability")
assert(probe["capabilities"]["ai_diagnostic_recommendations"], "runtime daemon probe must expose AI diagnostic recommendation capability")
assert(probe["capabilities"]["ai_repair_approval_gates"], "runtime daemon probe must expose AI repair approval gate capability")
assert(probe["capabilities"]["runtime_service_binding"], "runtime daemon probe must expose Runtime service binding capability")
assert(probe["capabilities"]["dbus_method_dispatch"], "runtime daemon probe must expose method dispatch capability")
assert(!probe["capabilities"]["dbus_binding"], "runtime daemon must not claim a D-Bus binding before it exists")
assert(probe["recipe_trust"]["registry_backed"], "runtime daemon probe must report registry-backed recipe loading")
assert(probe["recipe_trust"]["digest_verified"], "runtime daemon probe must report digest-verified recipes")
assert(!probe["recipe_trust"]["signed_recipe_validation"], "runtime daemon probe must not claim production recipe signatures yet")

stdout, stderr, status = Open3.capture3(*command, "list-applications")
assert(status.success?, "runtime daemon list must exit successfully: #{stderr}")
applications = JSON.parse(stdout)
assert(applications.length == 1, "runtime daemon must list the bundled sample application")
assert(applications.first["id"] == "org.xnix.sample.notepad", "runtime daemon must expose sample recipe id")
assert(applications.first["mime_types"].include?("application/x-xnix-txt"), "runtime daemon must expose MIME types")

stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon diagnostics must exit successfully: #{stderr}")
diagnostics = JSON.parse(stdout)
assert(diagnostics["status"] == "known", "runtime daemon diagnostics must know bundled recipes")
assert(diagnostics["checks"].any? { |check| check["status"] == "pending" }, "runtime daemon must report backend work as pending")
assert(diagnostics["test_plan"]["plan_type"] == "compatibility-test", "runtime daemon diagnostics must include compatibility test planning")
assert(diagnostics["test_plan"]["pending_step_count"] == 3, "runtime daemon diagnostics must count pending test steps")
assert(diagnostics["test_result"]["result_type"] == "compatibility-test-result", "runtime daemon diagnostics must include compatibility test results")
assert(diagnostics["test_result"]["overall_status"] == "pending", "runtime daemon diagnostics must expose pending test status")
assert(diagnostics["test_result"]["counts"]["pending"] == 3, "runtime daemon diagnostics must count pending test results")
assert(diagnostics["package_source"]["source_type"] == "compatibility-package-source", "runtime daemon diagnostics must include package sources")
assert(diagnostics["package_source"]["source_selection_state"] == "planned", "runtime daemon diagnostics must expose package source state")
assert(!diagnostics["package_source"]["package_source_ready"], "runtime daemon diagnostics must not claim package source readiness")
assert(!diagnostics["package_source"]["install_enabled"], "runtime daemon diagnostics must not enable package installation")
assert(diagnostics["state_root"]["root_type"] == "compatibility-application-state-root", "runtime daemon diagnostics must include application state roots")
assert(diagnostics["state_root"]["allocation_state"] == "planned", "runtime daemon diagnostics must expose planned state roots")
assert(!diagnostics["state_root"]["user_documents_included"], "runtime daemon diagnostics must exclude user documents from state roots")
assert(diagnostics["backend_binding"]["binding_type"] == "compatibility-backend-binding", "runtime daemon diagnostics must include backend binding")
assert(!diagnostics["backend_binding"]["managed_binding_ready"], "runtime daemon diagnostics must not claim backend binding readiness")
assert(!diagnostics["backend_binding"]["launch_enabled"], "runtime daemon diagnostics must not enable backend launch yet")
assert(diagnostics["ai_diagnostic_input"]["input_type"] == "ai-diagnostic-input", "runtime daemon diagnostics must include AI diagnostic input")
assert(diagnostics["ai_diagnostic_input"]["safe_for_ai_diagnostics"], "runtime daemon diagnostics must mark AI diagnostic input safe")
assert(!diagnostics["ai_diagnostic_input"]["ai_provider_called"], "runtime daemon diagnostics must not claim an AI provider call")
assert(diagnostics["ai_diagnostic_recommendation"]["recommendation_type"] == "ai-diagnostic-recommendation", "runtime daemon diagnostics must include AI diagnostic recommendations")
assert(diagnostics["ai_diagnostic_recommendation"]["recommendation_count"] == 3, "runtime daemon diagnostics must count AI diagnostic recommendations")
assert(!diagnostics["ai_diagnostic_recommendation"]["ai_provider_called"], "runtime daemon diagnostics must not claim recommendation AI provider calls")
assert(diagnostics["ai_repair_approval_gate"]["gate_type"] == "ai-repair-approval-gate", "runtime daemon diagnostics must include AI repair approval gates")
assert(diagnostics["ai_repair_approval_gate"]["gate_decision"] == "blocked-until-approval", "runtime daemon diagnostics must block repair execution until approval")
assert(!diagnostics["ai_repair_approval_gate"]["auto_execution_allowed"], "runtime daemon diagnostics must not allow automatic AI repair execution")
assert(diagnostics["runtime_service_binding"]["binding_type"] == "runtime-service-binding", "runtime daemon diagnostics must include Runtime service binding status")
assert(diagnostics["runtime_service_binding"]["activation_binding_ready"], "runtime daemon diagnostics must expose activation binding readiness")
assert(!diagnostics["runtime_service_binding"]["live_dbus_owner_ready"], "runtime daemon diagnostics must not claim live D-Bus ownership")
assert(diagnostics["repair_plan"]["plan_type"] == "compatibility-repair", "runtime daemon diagnostics must include repair planning")
assert(diagnostics["repair_plan"]["issue"] == "engine-binding-pending", "runtime daemon diagnostics must identify the pending repair issue")
assert(diagnostics["repair_plan"]["snapshot_required"], "runtime daemon diagnostics repair plan must require snapshots")
assert(diagnostics["repair_plan"]["snapshot_plan"]["plan_type"] == "compatibility-snapshot", "runtime daemon diagnostics must include snapshot planning")
assert(diagnostics["repair_plan"]["snapshot_plan"]["restore_available"], "runtime daemon diagnostics snapshot plan must expose restore availability")

diagnostics_json = JSON.pretty_generate(diagnostics)
assert(!diagnostics_json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "runtime diagnostics must not expose backend implementation terms")

_stdout, stderr, status = Open3.capture3(*command, "diagnostics", "org.xnix.missing")
assert(!status.success?, "runtime daemon must reject unknown applications")
assert(stderr.include?("unknown application"), "runtime daemon must explain unknown applications")

stdout, stderr, status = Open3.capture3(*command, "test-plan", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon test-plan must exit successfully: #{stderr}")
test_plan = JSON.parse(stdout)
assert(test_plan["plan_type"] == "compatibility-test", "runtime daemon test-plan command must emit a test plan")
assert(test_plan["steps"].any? { |step| step["id"] == "portal-preflight" }, "runtime daemon test-plan must include Portal preflight")

stdout, stderr, status = Open3.capture3(*command, "test-result", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon test-result must exit successfully: #{stderr}")
test_result = JSON.parse(stdout)
assert(test_result["result_type"] == "compatibility-test-result", "runtime daemon test-result command must emit a test result")
assert(test_result["overall_status"] == "pending", "runtime daemon test-result must expose pending status")

stdout, stderr, status = Open3.capture3(*command, "backend-binding", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon backend-binding must exit successfully: #{stderr}")
backend_binding = JSON.parse(stdout)
assert(backend_binding["binding_type"] == "compatibility-backend-binding", "runtime daemon backend-binding command must emit a backend binding")
assert(!backend_binding["managed_binding_ready"], "runtime daemon backend-binding must not claim readiness")

stdout, stderr, status = Open3.capture3(*command, "state-root", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon state-root must exit successfully: #{stderr}")
state_root = JSON.parse(stdout)
assert(state_root["root_type"] == "compatibility-application-state-root", "runtime daemon state-root command must emit a state root")
assert(state_root["allocation_state"] == "planned", "runtime daemon state-root must not claim allocation")

stdout, stderr, status = Open3.capture3(*command, "package-source", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon package-source must exit successfully: #{stderr}")
package_source = JSON.parse(stdout)
assert(package_source["source_type"] == "compatibility-package-source", "runtime daemon package-source command must emit a package source")
assert(!package_source["install_enabled"], "runtime daemon package-source must not enable installation")

stdout, stderr, status = Open3.capture3(*command, "ai-diagnostic-input", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon ai-diagnostic-input must exit successfully: #{stderr}")
ai_input = JSON.parse(stdout)
assert(ai_input["input_type"] == "ai-diagnostic-input", "runtime daemon ai-diagnostic-input command must emit an AI diagnostic input")
assert(ai_input["safe_for_ai_diagnostics"], "runtime daemon ai-diagnostic-input must be safe for AI diagnostics")

stdout, stderr, status = Open3.capture3(*command, "ai-diagnostic-recommendation", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon ai-diagnostic-recommendation must exit successfully: #{stderr}")
ai_recommendation = JSON.parse(stdout)
assert(ai_recommendation["recommendation_type"] == "ai-diagnostic-recommendation", "runtime daemon ai-diagnostic-recommendation command must emit recommendations")
assert(ai_recommendation["recommendations"].length == 3, "runtime daemon ai-diagnostic-recommendation must expose recommendations")

stdout, stderr, status = Open3.capture3(*command, "ai-repair-approval-gate", "org.xnix.sample.notepad")
assert(status.success?, "runtime daemon ai-repair-approval-gate must exit successfully: #{stderr}")
ai_gate = JSON.parse(stdout)
assert(ai_gate["gate_type"] == "ai-repair-approval-gate", "runtime daemon ai-repair-approval-gate command must emit a gate")
assert(ai_gate["gate_decision"] == "blocked-until-approval", "runtime daemon ai-repair-approval-gate must block until approval")

stdout, stderr, status = Open3.capture3(*command, "service-binding")
assert(status.success?, "runtime daemon service-binding must exit successfully: #{stderr}")
service_binding = JSON.parse(stdout)
assert(service_binding["binding_type"] == "runtime-service-binding", "runtime daemon service-binding command must emit service binding")
assert(service_binding["activation_binding_ready"], "runtime daemon service-binding must report activation readiness")

puts "PASS: compatibility runtime daemon unit tests"
