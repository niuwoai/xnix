#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/ai_diagnostic_recommendation"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
model = Xnix::Compatibility::AIDiagnosticRecommendation.new(recipe: recipe).to_h
recommendation_ids = model.fetch("recommendations").map { |recommendation| recommendation.fetch("id") }

assert(model["version"] == "0.2.110", "AI diagnostic recommendation must expose the current version")
assert(model["recommendation_type"] == "ai-diagnostic-recommendation", "AI diagnostic recommendation must identify the recommendation type")
assert(model["input_type"] == "ai-diagnostic-input", "AI diagnostic recommendation must reference the input type")
assert(model["application"]["id"] == "org.xnix.sample.notepad", "AI diagnostic recommendation must preserve the application id")
assert(model["runtime_owned"], "Runtime must own AI diagnostic recommendations")
assert(model["runtime_method"] == "GetAIDiagnosticRecommendation", "AI diagnostic recommendation must expose the Runtime method")
assert(model["c_runtime_backed"], "AI diagnostic recommendation must be backed by the C Runtime contract")
assert(!model["kde_policy_owner"], "KDE must not own AI diagnostic recommendation policy")
assert(!model["ai_provider_called"], "AI diagnostic recommendation must not call an AI provider")
assert(!model["network_required"], "AI diagnostic recommendation must not require network access")
assert(model["safe_for_ai_diagnostics"], "AI diagnostic recommendation must be safe for AI diagnostics")
assert(!model["auto_execution_allowed"], "AI diagnostic recommendation must not allow automatic execution")
assert(recommendation_ids == %w[explain-pending-runtime-work prepare-safe-restore-point surface-test-progress], "AI diagnostic recommendation must include expected recommendations")
assert(model["recommendations"].all? { |recommendation| recommendation["user_visible"] }, "AI diagnostic recommendations must be user-visible")
assert(model["recommendations"].none? { |recommendation| recommendation["auto_execute"] }, "AI diagnostic recommendations must not auto-execute")
assert(model["approval_required_actions"].length == 1, "AI diagnostic recommendations must expose approval-required actions")
assert(model["approval_required_actions"].first["approval_surface"] == "Compatibility Center", "AI diagnostic recommendations must use the Compatibility Center approval surface")
assert(model["blocked_actions"].include?("read user documents"), "AI diagnostic recommendation must preserve blocked AI tasks")
assert(model["privacy_boundaries"]["user_documents_included"] == false, "AI diagnostic recommendation must preserve privacy boundaries")
assert(!model["backend_details_exposed"], "AI diagnostic recommendation must hide backend details")

smoke_model = Xnix::Compatibility::AIDiagnosticRecommendation.new(recipe: recipe, test_type: "smoke").to_h
assert(smoke_model["recommendation_type"] == "ai-diagnostic-recommendation", "AI diagnostic recommendation must preserve explicit test type inputs")

json = JSON.pretty_generate(model)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "AI diagnostic recommendation must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-diagnostic-recommendation").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "AI diagnostic recommendation CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == model, "AI diagnostic recommendation CLI must emit the model")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-diagnostic-recommendation").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--test-type",
  "unknown"
)
assert(!status.success?, "AI diagnostic recommendation CLI must reject unknown test types")
assert(stderr.include?("invalid argument"), "AI diagnostic recommendation CLI must explain invalid test types")

puts "PASS: AI diagnostic recommendation unit tests"
