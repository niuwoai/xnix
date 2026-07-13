#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_package_source"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
source = Xnix::Compatibility::CompatibilityPackageSource.new(recipe: recipe).to_h
channel_ids = source.fetch("source_channels").map { |item| item.fetch("id") }
preflight_ids = source.fetch("required_preflight").map { |item| item.fetch("id") }

assert(source["version"] == "0.2.88", "compatibility package source must expose the current version")
assert(source["source_type"] == "compatibility-package-source", "compatibility package source must identify source type")
assert(source["application"]["id"] == "org.xnix.sample.notepad", "compatibility package source must preserve the application id")
assert(source["runtime_owned"], "Runtime must own compatibility package source selection")
assert(!source["kde_policy_owner"], "KDE must not own compatibility package source selection")
assert(source["selected_strategy"] == "automatic-managed", "compatibility package source must align with the run strategy")
assert(source["source_selection_state"] == "planned", "compatibility package source must not claim source selection yet")
assert(!source["package_source_ready"], "compatibility package source must not claim readiness yet")
assert(!source["install_enabled"], "compatibility package source must not enable installation yet")
assert(!source["network_required_for_planning"], "compatibility package source planning must not require network access")
assert(!source["host_root_modified"], "compatibility package source must not mutate the host root")
assert(!source["privileged_container_required"], "compatibility package source must not require privileged containers")
assert(!source["desktop_shell_command_exposed"], "compatibility package source must not expose commands to KDE")
assert(source["source_policy"]["signed_source_required"], "compatibility package source must require signed sources")
assert(!source["source_policy"]["direct_desktop_install_allowed"], "compatibility package source must block direct desktop installation")
assert(!source["source_policy"]["host_package_manager_invoked"], "compatibility package source must not invoke a host package manager")
assert(channel_ids == %w[os-managed-compatibility-packages runtime-managed-toolcache isolated-environment-template-catalog], "compatibility package source must expose expected channels")
assert(preflight_ids == %w[signed-source-verification source-policy-review runtime-cache-quota offline-fallback], "compatibility package source must expose expected preflight")
assert(source["blocked_actions"].include?("expose package manager commands to KDE"), "compatibility package source must block package command exposure")
assert(!source["backend_details_exposed"], "compatibility package source must hide backend details")

json = JSON.pretty_generate(source)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility package source must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "compatibility package source must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-package-source").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility package source CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == source, "compatibility package source CLI must emit the package source model")

puts "PASS: compatibility package source unit tests"
