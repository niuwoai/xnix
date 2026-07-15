#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_acquisition_preflight"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
preflight = Xnix::Compatibility::CompatibilityAcquisitionPreflight.new(recipe: recipe).to_h
check_ids = preflight.fetch("checks").map { |item| item.fetch("id") }

assert(preflight["version"] == "0.2.206", "compatibility acquisition preflight must expose the current version")
assert(preflight["preflight_type"] == "compatibility-acquisition-preflight", "compatibility acquisition preflight must identify preflight type")
assert(preflight["application"]["id"] == "org.xnix.sample.notepad", "compatibility acquisition preflight must preserve the application id")
assert(preflight["runtime_owned"], "Runtime must own compatibility acquisition preflight")
assert(!preflight["kde_policy_owner"], "KDE must not own compatibility acquisition preflight")
assert(preflight["selected_strategy"] == "automatic-managed", "compatibility acquisition preflight must align with package source strategy")
assert(preflight["preflight_state"] == "planned", "compatibility acquisition preflight must not claim execution yet")
assert(!preflight["acquisition_ready"], "compatibility acquisition preflight must not claim readiness yet")
assert(!preflight["download_enabled"], "compatibility acquisition preflight must not enable downloads")
assert(!preflight["install_enabled"], "compatibility acquisition preflight must not enable installation")
assert(!preflight["network_required_for_planning"], "compatibility acquisition preflight planning must not require network access")
assert(!preflight["network_request_created"], "compatibility acquisition preflight must not create network requests")
assert(!preflight["artifacts_downloaded"], "compatibility acquisition preflight must not download artifacts")
assert(!preflight["host_root_modified"], "compatibility acquisition preflight must not mutate the host root")
assert(!preflight["privileged_container_required"], "compatibility acquisition preflight must not require privileged containers")
assert(!preflight["desktop_shell_command_exposed"], "compatibility acquisition preflight must not expose commands to KDE")
assert(!preflight["package_source_ready"], "compatibility acquisition preflight must inherit pending package source readiness")
assert(check_ids == %w[package-source-ready signed-artifact-manifest runtime-cache-space network-policy-review rollback-marker], "compatibility acquisition preflight must expose expected checks")
assert(preflight["blocked_actions"].include?("invoke network access from KDE"), "compatibility acquisition preflight must block KDE network access")
assert(!preflight["backend_details_exposed"], "compatibility acquisition preflight must hide backend details")

json = JSON.pretty_generate(preflight)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility acquisition preflight must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "compatibility acquisition preflight must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-acquisition-preflight").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility acquisition preflight CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == preflight, "compatibility acquisition preflight CLI must emit the preflight model")

puts "PASS: compatibility acquisition preflight unit tests"
