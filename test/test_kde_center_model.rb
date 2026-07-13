#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/kde_center_model"
require_relative "../lib/xnix/compatibility/recipe_store"
require_relative "../lib/xnix/compatibility/runtime_daemon"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
runtime = Xnix::Compatibility::RuntimeDaemon.new(
  recipe_store: Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes"))
)
model = Xnix::Compatibility::KdeCenterModel.new(runtime: runtime).to_h

assert(model["version"] == "0.2.40", "KDE center model must expose the current version")
assert(model["source"]["kind"] == "runtime-local-read-model", "KDE center model must describe the local fallback read model")
assert(model["source"]["bus_name"] == "org.xnix.Compatibility1", "KDE center model must keep the Runtime bus boundary visible")
assert(model["summary"]["application_count"] == 1, "KDE center model must summarize bundled applications")
assert(model["summary"]["known_application_count"] == 1, "KDE center model must summarize known applications")
assert(model["summary"]["pending_action_count"] == 1, "KDE center model must summarize pending compatibility work")
assert(model["summary"]["pending_test_step_count"] == 3, "KDE center model must summarize pending compatibility test work")
assert(model["summary"]["pending_test_result_count"] == 1, "KDE center model must summarize pending test results")
assert(model["summary"]["ai_diagnostic_ready_count"] == 1, "KDE center model must summarize AI diagnostic readiness")
assert(model["summary"]["ai_recommendation_ready_count"] == 1, "KDE center model must summarize AI recommendation readiness")
assert(model["summary"]["blocked_ai_repair_gate_count"] == 1, "KDE center model must summarize blocked AI repair gates")
assert(model["summary"]["runtime_service_binding_ready_count"] == 1, "KDE center model must summarize Runtime service binding readiness")

application = model.fetch("applications").first
assert(application["id"] == "org.xnix.sample.notepad", "KDE center model must include the sample application")
assert(application["mode_label"] == "Automatic", "KDE center model must present user-facing mode labels")
assert(application["compatibility_label"] == "Known", "KDE center model must present user-facing status labels")
assert(application["summary"] == "1 compatibility task pending", "KDE center model must present a concise task summary")
assert(application["repair"]["issue"] == "engine-binding-pending", "KDE center model must expose repair issue summaries")
assert(application["repair"]["snapshot_required"], "KDE center model must expose repair snapshot requirements")
assert(application["repair"]["snapshot"]["plan_type"] == "compatibility-snapshot", "KDE center model must expose snapshot summaries")
assert(application["repair"]["notification_event"] == "approval-required", "KDE center model must map repair plans to notifications")
assert(application["test_plan"]["plan_type"] == "compatibility-test", "KDE center model must expose test plan summaries")
assert(application["test_plan"]["pending_step_count"] == 3, "KDE center model must expose pending test steps")
assert(application["test_result"]["result_type"] == "compatibility-test-result", "KDE center model must expose test result summaries")
assert(application["test_result"]["overall_status"] == "pending", "KDE center model must expose test result status")
assert(application["ai_diagnostic_input"]["input_type"] == "ai-diagnostic-input", "KDE center model must expose AI diagnostic input summaries")
assert(application["ai_diagnostic_input"]["safe_for_ai_diagnostics"], "KDE center model must expose AI diagnostic safety")
assert(!application["ai_diagnostic_input"]["ai_provider_called"], "KDE center model must not claim AI provider calls")
assert(application["ai_diagnostic_recommendation"]["recommendation_type"] == "ai-diagnostic-recommendation", "KDE center model must expose AI diagnostic recommendation summaries")
assert(application["ai_diagnostic_recommendation"]["recommendation_count"] == 3, "KDE center model must expose AI recommendation counts")
assert(!application["ai_diagnostic_recommendation"]["ai_provider_called"], "KDE center model must not claim recommendation AI provider calls")
assert(application["ai_repair_approval_gate"]["gate_type"] == "ai-repair-approval-gate", "KDE center model must expose AI repair approval gates")
assert(application["ai_repair_approval_gate"]["gate_decision"] == "blocked-until-approval", "KDE center model must expose AI repair gate decisions")
assert(!application["ai_repair_approval_gate"]["auto_execution_allowed"], "KDE center model must not allow automatic AI repair execution")
assert(application["runtime_service_binding"]["binding_type"] == "runtime-service-binding", "KDE center model must expose Runtime service binding")
assert(application["runtime_service_binding"]["activation_binding_ready"], "KDE center model must expose activation binding readiness")
assert(!application["runtime_service_binding"]["live_dbus_owner_ready"], "KDE center model must not claim live D-Bus ownership")
assert(application["supported_extensions"].include?(".txt"), "KDE center model must include supported file extensions")

json = JSON.pretty_generate(model)
assert(!json.match?(/prefix|\.wine|proton|virtual machine/i), "KDE center model must not expose backend storage or implementation terms")

MinimalRuntime = Struct.new(:application) do
  def list_applications
    [application]
  end

  def diagnostics(_application_id)
    {
      "application_id" => application.fetch("id"),
      "status" => "known",
      "runtime_mode" => "automatic"
    }
  end

  def source_metadata
    {
      "kind" => "runtime-dbus-session",
      "bus_name" => "org.xnix.Compatibility1",
      "object_path" => "/org/xnix/Compatibility1",
      "interface" => "org.xnix.Compatibility1"
    }
  end
end

minimal_model = Xnix::Compatibility::KdeCenterModel.new(runtime: MinimalRuntime.new(runtime.list_applications.first)).to_h
assert(minimal_model["summary"]["pending_test_step_count"].zero?, "KDE center model must tolerate missing D-Bus test plan summaries")
assert(minimal_model["summary"]["runtime_service_binding_ready_count"].zero?, "KDE center model must tolerate missing D-Bus service binding summaries")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-kde-center-model").to_s, "--source", "local")
assert(status.success?, "KDE center model CLI must exit successfully: #{stderr}")
cli_model = JSON.parse(stdout)
assert(cli_model == model, "KDE center model CLI must emit the same model")

puts "PASS: KDE compatibility center model unit tests"
