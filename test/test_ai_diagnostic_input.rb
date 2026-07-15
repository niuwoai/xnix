#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/ai_diagnostic_input"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
input = Xnix::Compatibility::AIDiagnosticInput.new(recipe: recipe).to_h
section_ids = input.fetch("context_sections").map { |section| section.fetch("id") }
signal_ids = input.fetch("diagnostic_signals").map { |signal| signal.fetch("id") }

assert(input["version"] == "0.2.175", "AI diagnostic input must expose the current version")
assert(input["input_type"] == "ai-diagnostic-input", "AI diagnostic input must identify the input type")
assert(input["application"]["id"] == "org.xnix.sample.notepad", "AI diagnostic input must preserve the application id")
assert(input["runtime_owned"], "Runtime must own AI diagnostic input")
assert(input["runtime_method"] == "GetAIDiagnosticInput", "AI diagnostic input must expose the Runtime method")
assert(input["c_runtime_backed"], "AI diagnostic input must be backed by the C Runtime contract")
assert(!input["kde_policy_owner"], "KDE must not own AI diagnostic policy")
assert(!input["ai_provider_called"], "AI diagnostic input must not call an AI provider")
assert(!input["network_required"], "AI diagnostic input must not require network access")
assert(input["safe_for_ai_diagnostics"], "AI diagnostic input must be marked safe for AI diagnostics")
assert(section_ids == %w[recipe run-plan test-result repair-plan], "AI diagnostic input must include the expected context sections")
assert(signal_ids.include?("pending-runtime-launch-binding"), "AI diagnostic input must expose pending launch binding signal")
assert(signal_ids.include?("pending-test-work"), "AI diagnostic input must expose pending test signal")
assert(signal_ids.include?("snapshot-before-risky-change"), "AI diagnostic input must expose snapshot signal")
assert(input["privacy_boundaries"]["user_documents_included"] == false, "AI diagnostic input must exclude user documents")
assert(input["privacy_boundaries"]["host_paths_included"] == false, "AI diagnostic input must exclude host paths")
assert(input["privacy_boundaries"]["raw_backend_logs_included"] == false, "AI diagnostic input must exclude raw backend logs")
assert(input["privacy_boundaries"]["secrets_included"] == false, "AI diagnostic input must exclude secrets")
assert(input["privacy_boundaries"]["network_calls_allowed"] == false, "AI diagnostic input must forbid network calls")
assert(input["allowed_ai_tasks"].include?("summarize compatibility status"), "AI diagnostic input must list allowed AI tasks")
assert(input["blocked_ai_tasks"].include?("read user documents"), "AI diagnostic input must list blocked AI tasks")
assert(!input["backend_details_exposed"], "AI diagnostic input must hide backend details")

smoke_input = Xnix::Compatibility::AIDiagnosticInput.new(recipe: recipe, test_type: "smoke").to_h
test_section = smoke_input.fetch("context_sections").find { |section| section.fetch("id") == "test-result" }
assert(test_section.fetch("facts").fetch("test_type") == "smoke", "AI diagnostic input must preserve explicit test type")

json = JSON.pretty_generate(input)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "AI diagnostic input must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-diagnostic-input").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "AI diagnostic input CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == input, "AI diagnostic input CLI must emit the input")

_stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-ai-diagnostic-input").to_s,
  "--app",
  "org.xnix.sample.notepad",
  "--test-type",
  "unknown"
)
assert(!status.success?, "AI diagnostic input CLI must reject unknown test types")
assert(stderr.include?("invalid argument"), "AI diagnostic input CLI must explain invalid test types")

puts "PASS: AI diagnostic input unit tests"
