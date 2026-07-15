#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/ai_repair_approval_gate"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
gate = Xnix::Compatibility::AIRepairApprovalGate.new(recipe: recipe).to_h
gate_ids = gate.fetch("required_gates").map { |item| item.fetch("id") }

assert(gate["version"] == "0.2.215", "AI repair approval gate must expose the current version")
assert(gate["gate_type"] == "ai-repair-approval-gate", "AI repair approval gate must identify the gate type")
assert(gate["recommendation_type"] == "ai-diagnostic-recommendation", "AI repair approval gate must reference recommendations")
assert(gate["application"]["id"] == "org.xnix.sample.notepad", "AI repair approval gate must preserve the application id")
assert(gate["runtime_method"] == "GetAIRepairApprovalGate", "AI repair approval gate must expose the Runtime method")
assert(gate["runtime_owned"], "Runtime must own AI repair approval gates")
assert(gate["c_runtime_backed"], "AI repair approval gate must expose C Runtime backing")
assert(!gate["kde_policy_owner"], "KDE must not own AI repair approval policy")
assert(!gate["ai_provider_called"], "AI repair approval gate must not call an AI provider")
assert(!gate["network_required"], "AI repair approval gate must not require network access")
assert(gate["safe_for_ai_diagnostics"], "AI repair approval gate must be safe for AI diagnostics")
assert(!gate["repair_execution_requested"], "AI repair approval gate must not request repair execution")
assert(!gate["repair_executed"], "AI repair approval gate must not execute repairs")
assert(!gate["auto_execution_allowed"], "AI repair approval gate must block automatic execution")
assert(gate["gate_decision"] == "blocked-until-approval", "AI repair approval gate must block execution until approval")
assert(gate["approval_surface"] == "Compatibility Center", "AI repair approval gate must use the Compatibility Center")
assert(gate_ids == %w[compatibility-center-review runtime-approval-token restore-point-preflight], "AI repair approval gate must include expected gates")
assert(gate["approval_required_actions"].length == 1, "AI repair approval gate must preserve approval-required actions")
assert(gate["blocked_actions"].include?("execute repair without approval"), "AI repair approval gate must block unapproved repair execution")
assert(gate["blocked_actions"].include?("read user documents"), "AI repair approval gate must preserve inherited blocked tasks")
assert(!gate["backend_details_exposed"], "AI repair approval gate must hide backend details")

smoke_gate = Xnix::Compatibility::AIRepairApprovalGate.new(recipe: recipe, test_type: "smoke").to_h
assert(smoke_gate["gate_type"] == "ai-repair-approval-gate", "AI repair approval gate must preserve explicit test type inputs")

json = JSON.pretty_generate(gate)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "AI repair approval gate must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-repair-approval-gate").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "AI repair approval gate CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == gate, "AI repair approval gate CLI must emit the gate")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-repair-approval-gate").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--test-type",
  "unknown"
)
assert(!status.success?, "AI repair approval gate CLI must reject unknown test types")
assert(stderr.include?("invalid argument"), "AI repair approval gate CLI must explain invalid test types")

puts "PASS: AI repair approval gate unit tests"
