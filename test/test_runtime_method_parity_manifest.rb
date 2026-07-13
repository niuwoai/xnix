#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require_relative "../lib/xnix/compatibility/runtime_method_parity_manifest"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
manifest = Xnix::Compatibility::RuntimeMethodParityManifest.new.to_h
check_ids = manifest.fetch("parity_checks").map { |item| item.fetch("id") }

assert(manifest["version"] == "0.2.84", "Runtime method parity manifest must expose the current version")
assert(manifest["manifest_type"] == "runtime-method-parity-manifest", "Runtime method parity manifest must identify the manifest type")
assert(manifest["runtime_owned"], "Runtime must own method parity")
assert(!manifest["kde_policy_owner"], "KDE must not own method parity")
assert(manifest["bus_name"] == "org.xnix.Compatibility1", "Runtime method parity manifest must expose the stable bus name")
assert(manifest["object_path"] == "/org/xnix/Compatibility1", "Runtime method parity manifest must expose the stable object path")
assert(manifest["interface"] == "org.xnix.Compatibility1", "Runtime method parity manifest must expose the stable interface")
assert(manifest["method_count"] == 34, "Runtime method parity manifest must count read-only methods")
assert(manifest["read_only_methods"].include?("GetDesktopActivationManifest"), "Runtime method parity manifest must include desktop activation manifest reads")
assert(manifest["read_only_methods"].include?("GetDesktopEntryPlan"), "Runtime method parity manifest must include desktop entry plan reads")
assert(manifest["read_only_methods"].include?("GetTaskManagerIdentityPlan"), "Runtime method parity manifest must include task manager identity plan reads")
assert(manifest["read_only_methods"].include?("GetFileAssociationPlan"), "Runtime method parity manifest must include file association plan reads")
assert(manifest["read_only_methods"].include?("GetNotificationPlan"), "Runtime method parity manifest must include notification plan reads")
assert(manifest["read_only_methods"].include?("GetPortalRequestPlan"), "Runtime method parity manifest must include Portal request plan reads")
assert(manifest["read_only_methods"].include?("GetRuntimeMethodParityManifest"), "Runtime method parity manifest must include itself")
assert(manifest["read_only_methods"].include?("GetRuntimeWriteGate"), "Runtime method parity manifest must include write gate queries")
assert(manifest["read_only_methods"].include?("GetRuntimeOwnerSmokePlan"), "Runtime method parity manifest must include owner smoke planning")
assert(check_ids == %w[dbus-contract runtime-dispatch dbus-client smoke-adapter session-smoke], "Runtime method parity manifest must check expected sources")
assert(manifest["parity_checks"].all? { |item| item["status"] == "pass" }, "Runtime method parity checks must pass")
assert(manifest["parity_checks"].all? { |item| item["missing_methods"].empty? }, "Runtime method parity checks must not miss methods")
assert(manifest["counts"] == { "total" => 5, "passed" => 5, "blocked" => 0, "pending" => 0 }, "Runtime method parity manifest must count checks")
assert(manifest["read_only_method_parity_ready"], "Runtime method parity manifest must report ready parity")
assert(manifest["write_methods"] == %w[InstallRecipe Launch CreateSnapshot RestoreSnapshot], "Runtime method parity manifest must list gated write methods")
assert(!manifest["write_methods_supported"], "Runtime method parity manifest must not claim write support")
assert(!manifest["write_method_dispatch_enabled"], "Runtime method parity manifest must not enable write dispatch")
assert(!manifest["network_required"], "Runtime method parity manifest must not require network access")
assert(!manifest["host_root_modified"], "Runtime method parity manifest must not mutate the host root")
assert(!manifest["privileged_container_required"], "Runtime method parity manifest must not require privileged containers")
assert(!manifest["backend_details_exposed"], "Runtime method parity manifest must hide backend details")

json = JSON.pretty_generate(manifest)
assert(!json.match?(/wine|prefix|\.wine|proton|virtual machine/i), "Runtime method parity manifest must not expose backend implementation terms")

stdout, stderr, status = Open3.capture3("ruby", project_root.join("bin/xnix-runtime-method-parity-manifest").to_s)
assert(status.success?, "Runtime method parity manifest CLI must exit successfully: #{stderr}")
assert(JSON.parse(stdout) == manifest, "Runtime method parity manifest CLI must emit the manifest model")

puts "PASS: Runtime method parity manifest unit tests"
