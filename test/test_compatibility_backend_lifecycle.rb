#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require_relative "../lib/xnix/compatibility/compatibility_backend_lifecycle"
require_relative "../lib/xnix/compatibility/recipe_store"
require_relative "../lib/xnix/compatibility/runtime_daemon"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

recipe_store = Xnix::Compatibility::RecipeStore.new(path: Xnix::Compatibility::RuntimeDaemon::DEFAULT_RECIPE_DIR)
recipe = recipe_store.find("org.xnix.sample.notepad")
lifecycle = Xnix::Compatibility::CompatibilityBackendLifecycle.new(recipe: recipe).to_h
stage_ids = lifecycle.fetch("stages").map { |stage| stage.fetch("id") }

assert(lifecycle["version"] == "0.2.142", "compatibility backend lifecycle must expose the current version")
assert(lifecycle["lifecycle_type"] == "compatibility-backend-lifecycle", "compatibility backend lifecycle must identify the lifecycle type")
assert(lifecycle["runtime_method"] == "GetBackendLifecycle", "compatibility backend lifecycle must identify the Runtime method")
assert(lifecycle["application"]["id"] == "org.xnix.sample.notepad", "compatibility backend lifecycle must preserve the application id")
assert(lifecycle["runtime_owned"], "Runtime must own backend lifecycle")
assert(lifecycle["c_runtime_backed"], "backend lifecycle must be marked as C Runtime-backed")
assert(!lifecycle["kde_policy_owner"], "KDE must not own backend lifecycle")
assert(lifecycle["lifecycle_state"] == "blocked", "backend lifecycle must remain blocked")
assert(lifecycle["overall_status"] == "not-ready", "backend lifecycle must not claim readiness")
assert(!lifecycle["backend_binding_ready"], "backend lifecycle must not claim backend binding readiness")
assert(!lifecycle["launch_enabled"], "backend lifecycle must not enable launch")
assert(!lifecycle["execution_request_created"], "backend lifecycle must not create execution requests")
assert(!lifecycle["backend_process_started"], "backend lifecycle must not start backend processes")
assert(!lifecycle["local_backend_started"], "backend lifecycle must not start local backends")
assert(!lifecycle["isolated_backend_started"], "backend lifecycle must not start isolated backends")
assert(!lifecycle["state_root_ready"], "backend lifecycle must not claim state-root readiness")
assert(lifecycle["portal_review_required"], "backend lifecycle must require Portal review")
assert(lifecycle["snapshot_required"], "backend lifecycle must require snapshot review")
assert(stage_ids == %w[recipe-loaded state-root-ready backend-binding-ready portal-and-snapshot-review runtime-launch-write-gate], "backend lifecycle must expose expected stages")
assert(lifecycle["blocked_actions"].include?("start local compatibility backend from KDE"), "backend lifecycle must block local backend starts from KDE")
assert(lifecycle["blocked_actions"].include?("start isolated compatibility backend from KDE"), "backend lifecycle must block isolated backend starts from KDE")
assert(!lifecycle["host_root_modified"], "backend lifecycle must not mutate the host root")
assert(!lifecycle["network_required"], "backend lifecycle must not require network access")
assert(!lifecycle["backend_details_exposed"], "backend lifecycle must hide backend details")

json = JSON.pretty_generate(lifecycle)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "backend lifecycle must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3(
  "ruby",
  File.expand_path("../bin/xnix-compat-backend-lifecycle", __dir__),
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility backend lifecycle CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == lifecycle, "compatibility backend lifecycle CLI must emit the lifecycle model")

puts "PASS: compatibility backend lifecycle unit tests"
