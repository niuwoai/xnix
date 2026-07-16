#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/compatibility_artifact_manifest"
require_relative "../lib/xnix/compatibility/recipe_store"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
recipe = Xnix::Compatibility::RecipeStore.new(path: project_root.join("runtime/recipes")).find("org.xnix.sample.notepad")
manifest = Xnix::Compatibility::CompatibilityArtifactManifest.new(recipe: recipe).to_h
group_ids = manifest.fetch("artifact_groups").map { |item| item.fetch("id") }
preflight_ids = manifest.fetch("required_preflight").map { |item| item.fetch("id") }

assert(manifest["version"] == "0.2.245", "compatibility artifact manifest must expose the current version")
assert(manifest["manifest_type"] == "compatibility-artifact-manifest", "compatibility artifact manifest must identify manifest type")
assert(manifest["application"]["id"] == "org.xnix.sample.notepad", "compatibility artifact manifest must preserve the application id")
assert(manifest["runtime_owned"], "Runtime must own compatibility artifact manifests")
assert(!manifest["kde_policy_owner"], "KDE must not own compatibility artifact manifests")
assert(manifest["selected_strategy"] == "automatic-managed", "compatibility artifact manifest must align with acquisition strategy")
assert(manifest["manifest_state"] == "planned", "compatibility artifact manifest must not claim resolution yet")
assert(!manifest["manifest_ready"], "compatibility artifact manifest must not claim readiness yet")
assert(!manifest["signature_verified"], "compatibility artifact manifest must not claim signature verification yet")
assert(!manifest["acquisition_preflight_ready"], "compatibility artifact manifest must inherit pending acquisition preflight")
assert(!manifest["download_enabled"], "compatibility artifact manifest must not enable downloads")
assert(!manifest["install_enabled"], "compatibility artifact manifest must not enable installation")
assert(!manifest["network_request_created"], "compatibility artifact manifest must not create network requests")
assert(!manifest["artifacts_downloaded"], "compatibility artifact manifest must not download artifacts")
assert(!manifest["host_root_modified"], "compatibility artifact manifest must not mutate the host root")
assert(!manifest["privileged_container_required"], "compatibility artifact manifest must not require privileged containers")
assert(!manifest["desktop_shell_command_exposed"], "compatibility artifact manifest must not expose commands to KDE")
assert(group_ids == %w[runtime-launch-metadata local-execution-artifacts isolated-environment-artifacts], "compatibility artifact manifest must expose expected artifact groups")
assert(preflight_ids == %w[acquisition-preflight-ready manifest-signature-verification artifact-digest-verification cache-namespace-allocation rollback-reference], "compatibility artifact manifest must expose expected preflight")
assert(manifest["blocked_actions"].include?("expose artifact cache paths to KDE"), "compatibility artifact manifest must block cache path exposure")
assert(!manifest["backend_details_exposed"], "compatibility artifact manifest must hide backend details")

json = JSON.pretty_generate(manifest)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "compatibility artifact manifest must not expose backend implementation terms")
assert(!json.match?(%r{/Users|/home|/var|/opt|/tmp}), "compatibility artifact manifest must not expose host storage paths")

stdout, stderr, status = Open3.capture3(
  "ruby",
  project_root.join("bin/xnix-compat-artifact-manifest").to_s,
  "--app",
  "org.xnix.sample.notepad"
)
assert(status.success?, "compatibility artifact manifest CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == manifest, "compatibility artifact manifest CLI must emit the manifest model")

puts "PASS: compatibility artifact manifest unit tests"
