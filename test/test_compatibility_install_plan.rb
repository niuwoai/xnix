#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_install_plan"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
plan = Xnix::Compatibility::CompatibilityInstallPlan.new(recipe: recipe).to_h
phase_ids = plan.fetch("phases").map { |phase| phase.fetch("id") }

assert(plan["version"] == "0.2.64", "compatibility install plan must expose the current version")
assert(plan["plan_type"] == "compatibility-install-plan", "compatibility install plan must identify plan type")
assert(plan["application"]["id"] == "org.xnix.sample.notepad", "compatibility install plan must preserve the application id")
assert(plan["runtime_owned"], "Runtime must own compatibility install plans")
assert(!plan["kde_policy_owner"], "KDE must not own compatibility install plans")
assert(plan["environment"] == "development", "compatibility install plan must default to development gate evaluation")
assert(plan["selected_strategy"] == "automatic-managed", "compatibility install plan must align with artifact strategy")
assert(plan["install_state"] == "planned", "compatibility install plan must not claim install execution")
assert(!plan["install_ready"], "compatibility install plan must not claim install readiness yet")
assert(!plan["desktop_activation_ready"], "compatibility install plan must not claim desktop activation readiness yet")
assert(!plan["download_enabled"], "compatibility install plan must not enable downloads")
assert(!plan["install_enabled"], "compatibility install plan must not enable installation")
assert(!plan["network_request_created"], "compatibility install plan must not create network requests")
assert(!plan["artifacts_downloaded"], "compatibility install plan must not download artifacts")
assert(!plan["host_root_modified"], "compatibility install plan must not mutate the host root")
assert(!plan["privileged_container_required"], "compatibility install plan must not require privileged containers")
assert(!plan["desktop_shell_command_exposed"], "compatibility install plan must not expose commands to KDE")
assert(!plan["readiness"]["artifact_manifest_ready"], "compatibility install plan must wait for artifact manifest readiness")
assert(!plan["readiness"]["artifact_signature_verified"], "compatibility install plan must wait for artifact signatures")
assert(!plan["readiness"]["acquisition_ready"], "compatibility install plan must wait for acquisition readiness")
assert(!plan["readiness"]["package_source_ready"], "compatibility install plan must wait for package source readiness")
assert(!plan["readiness"]["state_root_allocated"], "compatibility install plan must wait for state allocation")
assert(plan["readiness"]["recipe_install_allowed"], "development compatibility install plan must preserve digest-verified staging")
assert(phase_ids == %w[resolve-artifact-manifest verify-artifact-digests prepare-package-source allocate-application-state stage-desktop-integration enable-launch-binding], "compatibility install plan must expose expected phases")
assert(plan["blocked_actions"].include?("mutate the host root during install planning"), "compatibility install plan must block host root mutation")
assert(!plan["backend_details_exposed"], "compatibility install plan must hide backend details")

json = JSON.pretty_generate(plan)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility install plan must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "compatibility install plan must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-install-plan").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility install plan CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == plan, "compatibility install plan CLI must emit the install plan model")

production_plan = Xnix::Compatibility::CompatibilityInstallPlan.new(recipe: recipe, environment: "production").to_h
assert(!production_plan["readiness"]["recipe_install_allowed"], "production compatibility install plan must block development-only recipe registries")
assert(production_plan["readiness"]["recipe_install_decision"] == "block", "production compatibility install plan must expose blocked decision")

puts "PASS: compatibility install plan unit tests"
