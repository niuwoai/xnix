#!/usr/bin/env ruby
# frozen_string_literal: true

require "json"
require "open3"
require "pathname"
require "tempfile"

def assert(condition, message)
  return if condition

  warn "FAIL: #{message}"
  exit 1
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/mainline_integration_review.rb")

fixture = Tempfile.new("xnix-mainline-integration-status")
begin
  fixture.write(" M CHANGELOG.md\n")
  fixture.write(" M docs/windows-app-compatibility-implementation-brief.md\n")
  fixture.write("?? docs/mainline-integration-checkpoint.md\n")
  fixture.write("?? docs/xnix-current-mainline.md\n")
  fixture.write(" M internal/runtime/owner/session_bus.go\n")
  fixture.write("?? cmd/xnix-runtime-go/current_mainline_commands.go\n")
  fixture.write(" M internal/runtime/artifact/stage.go\n")
  fixture.write(" M internal/runtime/appidentity/backend_manager.go\n")
  fixture.write(" M internal/runtime/portal/ledger.go\n")
  fixture.write(" M internal/runtime/execution/session.go\n")
  fixture.write(" M cmd/xnix-runtime-go/main.go\n")
  fixture.write(" M test/test_ai_diagnostic_input.rb\n")
  fixture.write(" M test/test_compatibility_snapshot_plan.rb\n")
  fixture.write(" M internal/runtime/appidentity/windows_compatibility_workstreams.go\n")
  fixture.write(" M scripts/implementation_evidence_report.rb\n")
  fixture.write(" M scripts/kde_first_presence_smoke.rb\n")
  fixture.write(" M .dockerignore\n")
  fixture.write(" M lib/xnix/full_smoke_report.rb\n")
  fixture.write(" M docs/kde-image-pipeline.md\n")
  fixture.write("?? docs/release-evidence/v0.2.320-rc7-kde-product-smoke.json\n")
  fixture.write(" M docs/claude-code-implementation-packages.md\n")
  fixture.write("?? .gocache/00/cache-entry\n")
  fixture.write("?? tmp/runtime-output.json\n")
  fixture.write("?? experiments/side-quest.txt\n")
  fixture.flush

  stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json", "--status-fixture", fixture.path)
  assert(status.success?, "mainline integration review JSON must exit successfully: #{stderr}")
  report = JSON.parse(stdout)

  assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "mainline integration review must expose the current version")
  assert(report.fetch("schema_version") == "xnix.runtime.mainline_integration_review.v1", "mainline integration review must expose the schema")
  assert(report.fetch("report_type") == "mainline-integration-review", "mainline integration review must identify its report type")
  assert(report.fetch("checkpoint_document") == "docs/mainline-integration-checkpoint.md", "mainline integration review must point at the checkpoint")
  assert(report.fetch("protected_claude_file") == "docs/claude-code-implementation-packages.md", "mainline integration review must name the protected Claude file")
  assert(report.fetch("protected_claude_file_modified"), "mainline integration review must flag protected Claude file changes")
  assert(report.fetch("excluded_prefixes") == [".gocache/", "tmp/"], "mainline integration review must expose excluded prefixes")
  assert(report.fetch("excluded_file_count") == 2, "mainline integration review must exclude cache and tmp files")
  assert(report.fetch("changed_file_count") == 22, "mainline integration review must count included fixture files")
  assert(!report.fetch("safe_to_stage_all"), "mainline integration review must never mark a mixed tree safe to stage all")
  assert(!report.fetch("docker_or_qemu_required"), "mainline integration review must not require Docker or QEMU")
  assert(!report.fetch("host_root_modified"), "mainline integration review must not mutate the host root")
  assert(!report.fetch("privileged_container_required"), "mainline integration review must not require privileged containers")
  assert(!report.fetch("backend_launch_enabled"), "mainline integration review must not enable backend launch")
  assert(report.fetch("review_matrix").is_a?(Array), "mainline integration review must expose a lane review matrix")

  lanes = report.fetch("lanes").to_h { |lane| [lane.fetch("id"), lane] }
  expected_lanes = %w[
    planning-documents
    version-and-product-metadata
    cw1-runtime-owner-read-boundary
    cw2-recipe-artifact-trust
    cw3-runtime-state-backend-lifecycle
    cw5-portal-permission-safety
    cw8-execution-ledger-session-evidence
    shared-runtime-cli-plumbing
    cw6-snapshot-rollback-store
    cw7-ai-diagnostic-privacy
    cw4-kde-entrypoint-consumers
    cw10-evidence-drift-harness
    cw11-product-image-acceptance
    blocked-protected-claude-owned-file
    unclassified
  ]
  assert((expected_lanes - lanes.keys).empty?, "mainline integration review must expose all expected fixture lanes")
  assert(lanes.fetch("planning-documents").fetch("paths").include?("docs/mainline-integration-checkpoint.md"), "planning lane must include the checkpoint")
  assert(lanes.fetch("planning-documents").fetch("paths").include?("docs/xnix-current-mainline.md"), "planning lane must include the Xnix current mainline")
  assert(lanes.fetch("version-and-product-metadata").fetch("paths").include?("CHANGELOG.md"), "version lane must include CHANGELOG")
  assert(lanes.fetch("cw1-runtime-owner-read-boundary").fetch("paths").include?("internal/runtime/owner/session_bus.go"), "CW1 lane must include owner session bus work")
  assert(lanes.fetch("cw1-runtime-owner-read-boundary").fetch("paths").include?("cmd/xnix-runtime-go/current_mainline_commands.go"), "CW1 lane must include current mainline owner audit CLI work")
  assert(lanes.fetch("cw2-recipe-artifact-trust").fetch("paths").include?("internal/runtime/artifact/stage.go"), "CW2 lane must include artifact staging work")
  assert(lanes.fetch("cw3-runtime-state-backend-lifecycle").fetch("paths").include?("internal/runtime/appidentity/backend_manager.go"), "CW3 lane must include backend manager work")
  assert(lanes.fetch("cw5-portal-permission-safety").fetch("paths").include?("internal/runtime/portal/ledger.go"), "CW5 lane must include Portal ledger work")
  assert(lanes.fetch("cw8-execution-ledger-session-evidence").fetch("paths").include?("internal/runtime/execution/session.go"), "CW8 lane must include execution session work")
  assert(lanes.fetch("shared-runtime-cli-plumbing").fetch("paths").include?("cmd/xnix-runtime-go/main.go"), "shared plumbing lane must include Runtime CLI registration work")
  assert(lanes.fetch("cw6-snapshot-rollback-store").fetch("paths").include?("test/test_compatibility_snapshot_plan.rb"), "CW6 lane must include snapshot and rollback work")
  assert(lanes.fetch("cw7-ai-diagnostic-privacy").fetch("paths").include?("test/test_ai_diagnostic_input.rb"), "CW7 lane must include AI diagnostic work")
  assert(lanes.fetch("cw4-kde-entrypoint-consumers").fetch("paths").include?("scripts/kde_first_presence_smoke.rb"), "CW4 lane must include KDE-first presence smoke work")
  assert(lanes.fetch("cw10-evidence-drift-harness").fetch("paths").include?("scripts/implementation_evidence_report.rb"), "CW10 lane must include evidence report work")
  assert(lanes.fetch("cw10-evidence-drift-harness").fetch("paths").include?("internal/runtime/appidentity/windows_compatibility_workstreams.go"), "CW10 lane must include Windows compatibility workstream model work")
  cw11_paths = lanes.fetch("cw11-product-image-acceptance").fetch("paths")
  expected_cw11_paths = [".dockerignore", "docs/kde-image-pipeline.md", "docs/release-evidence/v0.2.320-rc7-kde-product-smoke.json", "lib/xnix/full_smoke_report.rb"]
  assert((expected_cw11_paths - cw11_paths).empty? && (cw11_paths - expected_cw11_paths).empty?, "CW11 lane must classify Docker context, image documentation, persisted smoke evidence, and full smoke reports")
  assert(lanes.fetch("blocked-protected-claude-owned-file").fetch("paths") == ["docs/claude-code-implementation-packages.md"], "protected lane must isolate the Claude-owned file")
  assert(lanes.fetch("unclassified").fetch("paths") == ["experiments/side-quest.txt"], "unclassified lane must isolate unmatched paths")
  assert(report.fetch("review_order").first == "planning-documents", "mainline integration review must order planning first")

  cw1_matrix = report.fetch("review_matrix").find { |entry| entry.fetch("id") == "cw1-runtime-owner-read-boundary" }
  assert(cw1_matrix.fetch("required_verification").include?("ruby scripts/runtime_contract_drift_report.rb --format json"), "CW1 matrix must require contract drift verification")
  assert(cw1_matrix.fetch("required_evidence").any? { |line| line.include?("Production D-Bus ownership remains blocked") }, "CW1 matrix must require production D-Bus ownership blocking evidence")
  assert(cw1_matrix.fetch("safety_guards").include?("runtime-writes-disabled"), "CW1 matrix must expose disabled write safety guard")

  cw10_matrix = report.fetch("review_matrix").find { |entry| entry.fetch("id") == "cw10-evidence-drift-harness" }
  assert(cw10_matrix.fetch("required_verification").include?("ruby scripts/mainline_integration_review.rb --format json"), "CW10 matrix must require the integration review report")
  assert(cw10_matrix.fetch("safety_guards").include?("never-stage-all"), "CW10 matrix must keep never-stage-all guard visible")

  markdown, markdown_stderr, markdown_status = Open3.capture3("ruby", script.to_s, "--format", "markdown", "--status-fixture", fixture.path)
  assert(markdown_status.success?, "mainline integration review Markdown must exit successfully: #{markdown_stderr}")
  assert(markdown.include?("# Mainline Integration Review"), "Markdown report must include a title")
  assert(markdown.include?("cw1-runtime-owner-read-boundary"), "Markdown report must include CW1 lane")
  assert(markdown.include?("Protected Claude file modified: true"), "Markdown report must expose protected file status")
  assert(markdown.include?("Required verification:"), "Markdown report must include required verification commands")
  assert(markdown.include?("Safety guards:"), "Markdown report must include safety guards")
ensure
  fixture.close
  fixture.unlink
end
