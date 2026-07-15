#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/application_state_root"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
state_root = Xnix::Compatibility::ApplicationStateRoot.new(recipe: recipe).to_h
scope_ids = state_root.fetch("managed_scopes").map { |item| item.fetch("id") }

assert(state_root["version"] == "0.2.172", "application state root must expose the current version")
assert(state_root["root_type"] == "compatibility-application-state-root", "application state root must identify the root type")
assert(state_root["application"]["id"] == "org.xnix.sample.notepad", "application state root must preserve the application id")
assert(state_root["runtime_owned"], "Runtime must own application state roots")
assert(!state_root["kde_policy_owner"], "KDE must not own application state roots")
assert(state_root["state_namespace"] == "org.xnix.sample.notepad", "application state root must expose a stable namespace")
assert(state_root["storage_scope"] == "per-application", "application state root must be scoped per application")
assert(state_root["allocation_state"] == "planned", "application state root must not claim allocation yet")
assert(!state_root["directories_created"], "application state root must not create directories during planning")
assert(!state_root["host_root_modified"], "application state root must not mutate the host root")
assert(!state_root["user_documents_included"], "application state root must exclude user documents")
assert(state_root["portal_required_for_user_files"], "application state root must require Portal grants for user files")
assert(state_root["snapshot_eligible"], "application state root must be snapshot eligible")
assert(state_root["restore_requires_confirmation"], "application state root restore must require confirmation")
assert(scope_ids == %w[application-data runtime-metadata diagnostic-cache desktop-activation-receipts], "application state root must expose managed scopes")
assert(state_root["blocked_actions"].include?("expose host storage paths to KDE"), "application state root must block host path exposure")
assert(!state_root["backend_details_exposed"], "application state root must hide backend details")

json = JSON.pretty_generate(state_root)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "application state root must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "application state root must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-state-root").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "application state root CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == state_root, "application state root CLI must emit the state root model")

puts "PASS: application state root unit tests"
