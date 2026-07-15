#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_backend_binding"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
binding = Xnix::Compatibility::CompatibilityBackendBinding.new(recipe: recipe).to_h
preflight_ids = binding.fetch("required_preflight").map { |item| item.fetch("id") }

assert(binding["version"] == "0.2.156", "compatibility backend binding must expose the current version")
assert(binding["binding_type"] == "compatibility-backend-binding", "compatibility backend binding must identify the binding type")
assert(binding["application"]["id"] == "org.xnix.sample.notepad", "compatibility backend binding must preserve the application id")
assert(binding["runtime_owned"], "Runtime must own compatibility backend binding")
assert(!binding["kde_policy_owner"], "KDE must not own compatibility backend binding")
assert(binding["selected_strategy"] == "automatic-managed", "compatibility backend binding must use Runtime run-plan strategy")
assert(!binding["managed_binding_ready"], "compatibility backend binding must not claim readiness yet")
assert(!binding["launch_enabled"], "compatibility backend binding must not enable launch yet")
assert(!binding["execution_request_created"], "compatibility backend binding must not create execution requests")
assert(!binding["host_root_modified"], "compatibility backend binding must not modify the host root")
assert(!binding["network_required"], "compatibility backend binding must not require network access")
assert(!binding["privileged_container_required"], "compatibility backend binding must not require privileged containers")
assert(preflight_ids == %w[engine-package-source application-state-root portal-policy-review snapshot-baseline], "compatibility backend binding must expose required preflight")
assert(binding["blocked_actions"].include?("expose backend command to desktop shell"), "compatibility backend binding must block desktop backend command exposure")
assert(!binding["backend_details_exposed"], "compatibility backend binding must hide backend details")

json = JSON.pretty_generate(binding)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility backend binding must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-backend-binding").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility backend binding CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == binding, "compatibility backend binding CLI must emit the binding model")

puts "PASS: compatibility backend binding unit tests"
