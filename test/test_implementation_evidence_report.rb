#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/implementation_evidence_report.rb")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json")
assert(status.success?, "implementation evidence report JSON must exit successfully: #{stderr}")
report = JSON.parse(stdout)

assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "implementation evidence report must expose the current version")
assert(report.fetch("schema_version") == "xnix.runtime.implementation_evidence_report.v1", "implementation evidence report must expose the schema")
assert(report.fetch("report_type") == "implementation-evidence-report", "implementation evidence report must identify its report type")
assert(report.fetch("runtime_owned"), "implementation evidence report must keep Runtime ownership explicit")
assert(report.fetch("go_runtime_backed"), "implementation evidence report must be Go Runtime backed")
assert(!report.fetch("kde_policy_owner"), "implementation evidence report must not make KDE the policy owner")
assert(report.fetch("domain_status_order") == %w[missing contract-only fixture-implemented state-root-implemented smoke-owned production-gated], "implementation evidence report must expose stable status order")
assert(report.fetch("domains").length == 9, "implementation evidence report must cover the empty-domain packages")
assert(report.fetch("counts").fetch("total") == 9, "implementation evidence report must count domains")
assert(report.fetch("read_only_method_count") == 57, "implementation evidence report must count Runtime read-only methods")
assert(report.fetch("orphan_read_methods").empty?, "implementation evidence report must not find orphan read methods")
assert(!report.fetch("orphan_preview_methods_detected"), "implementation evidence report must not detect orphan preview methods")
assert(report.fetch("write_methods") == %w[InstallRecipe Launch CreateSnapshot RestoreSnapshot], "implementation evidence report must list gated write methods")
assert(!report.fetch("write_methods_supported"), "implementation evidence report must not support write methods")
assert(!report.fetch("write_method_dispatch_enabled"), "implementation evidence report must not enable write dispatch")
assert(!report.fetch("network_required"), "implementation evidence report must not require network")
assert(!report.fetch("host_root_modified"), "implementation evidence report must not mutate the host root")
assert(!report.fetch("privileged_container_required"), "implementation evidence report must not require privileged containers")
assert(!report.fetch("backend_launch_enabled"), "implementation evidence report must not enable backend launch")
assert(!report.fetch("production_ready"), "implementation evidence report must keep production readiness gated")
assert(report.fetch("highest_evidence_status") == "smoke-owned", "implementation evidence report must recognize smoke-owned evidence as the highest current implementation depth")
assert(report.fetch("counts").fetch("production_gate_evidence") == 9, "implementation evidence report must count production gate evidence separately")

domains = report.fetch("domains").to_h { |domain| [domain.fetch("id"), domain] }
expected_domains = %w[
  runtime-owner-service
  recipe-artifact-trust-pipeline
  environment-lifecycle-state
  portal-snapshot-control-plane
  kde-activation-shell-materialization
  execution-transaction-ledger
  diagnostics-repair-ai-boundary
  atomic-kde-image-qemu-acceptance
  developer-verification-harness
]
assert(domains.keys == expected_domains, "implementation evidence report must expose stable domain ids")

assert(domains.fetch("runtime-owner-service").fetch("status") == "smoke-owned", "Runtime owner service must show smoke-owned evidence")
assert(domains.fetch("recipe-artifact-trust-pipeline").fetch("status") == "state-root-implemented", "recipe and artifact trust pipeline must show state-root implementation evidence")
assert(domains.fetch("portal-snapshot-control-plane").fetch("status") == "state-root-implemented", "Portal and snapshot control plane must show state-root implementation evidence")
assert(domains.fetch("execution-transaction-ledger").fetch("status") == "state-root-implemented", "execution transaction ledger must show state-root implementation evidence")
assert(domains.fetch("developer-verification-harness").fetch("status") == "fixture-implemented", "developer verification harness must show fixture implementation evidence")
assert(domains.values.all? { |domain| domain.fetch("production_gate_evidence") }, "all implementation domains must expose production gate evidence")
assert(domains.values.all? { |domain| domain.fetch("contract_files_missing").empty? }, "all implementation domains must have their contract files")
assert(domains.values.all? { |domain| domain.fetch("fixture_files_missing").empty? }, "all implementation domains must have their fixture files")
assert(domains.values.all? { |domain| domain.fetch("gate_tokens_missing").empty? }, "all implementation domains must have their gate tokens")
assert(domains.values.all? { |domain| !domain.fetch("host_root_modified") }, "domains must not mutate host root")
assert(domains.values.all? { |domain| !domain.fetch("network_required") }, "domains must not require network")
assert(domains.values.all? { |domain| !domain.fetch("privileged_container_required") }, "domains must not require privileged containers")
assert(domains.values.all? { |domain| !domain.fetch("backend_launch_enabled") }, "domains must not enable backend launch")

stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "markdown")
assert(status.success?, "implementation evidence report markdown must exit successfully: #{stderr}")
assert(stdout.include?("# Implementation Evidence Report"), "implementation evidence report markdown must include a title")
assert(stdout.include?("| P1 | Runtime owner service | smoke-owned |"), "implementation evidence report markdown must include domain rows")
assert(stdout.include?("Implementation evidence spans 9 domains; production readiness remains gated."), "implementation evidence report markdown must include a safe summary")

puts "PASS: implementation evidence report tests"
