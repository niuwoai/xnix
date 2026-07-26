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

def write_json_fixture(payload)
  file = Tempfile.new("xnix-release-evidence-index")
  file.write(JSON.pretty_generate(payload))
  file.flush
  file
end

def write_text_fixture(text)
  file = Tempfile.new("xnix-release-evidence-index")
  file.write(text)
  file.flush
  file
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/release_evidence_index.rb")

implementation = write_json_fixture(
  "domains" => [
    {
      "id" => "runtime-owner-service",
      "status" => "smoke-owned",
      "summary" => "Go owner read boundary exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "recipe-artifact-trust-pipeline",
      "status" => "state-root-implemented",
      "summary" => "Recipe and artifact evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "environment-lifecycle-state",
      "status" => "fixture-implemented",
      "summary" => "Backend lifecycle fixture evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "portal-snapshot-control-plane",
      "status" => "contract-only",
      "summary" => "Portal and snapshot contracts exist.",
      "contract_files_missing" => [],
      "fixture_files_missing" => ["internal/runtime/portal/ledger.go"]
    },
    {
      "id" => "kde-activation-shell-materialization",
      "status" => "smoke-owned",
      "summary" => "KDE seven entry point smoke exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "execution-transaction-ledger",
      "status" => "state-root-implemented",
      "summary" => "Execution ledger exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "diagnostics-repair-ai-boundary",
      "status" => "missing",
      "summary" => "Diagnostics evidence is missing.",
      "contract_files_missing" => ["internal/runtime/appidentity/ai_diagnostics.go"],
      "fixture_files_missing" => []
    },
    {
      "id" => "developer-verification-harness",
      "status" => "fixture-implemented",
      "summary" => "Report fixture evidence exists.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    },
    {
      "id" => "atomic-kde-image-qemu-acceptance",
      "status" => "smoke-owned",
      "summary" => "Image smoke scaffolding exists but needs explicit authorization.",
      "contract_files_missing" => [],
      "fixture_files_missing" => []
    }
  ]
)
contract_drift = write_json_fixture("drift_detected" => false)
mainline = write_json_fixture(
  "protected_claude_file_modified" => true,
  "unclassified_file_count" => 2
)
kde_smoke = write_json_fixture("entrypoint_count" => 7)
completed_promotion = write_json_fixture(
  "schema_version" => "xnix.runtime.full_checkpoint_promotion_packet.v1",
  "report_type" => "full-checkpoint-promotion-packet",
  "promotion_allowed" => true,
  "promotion_decision" => "promote",
  "full_smoke_state" => "passed",
  "formal_release_ready" => true,
  "operator_required_command" => nil,
  "full_smoke_executed_by_packet" => false,
  "docker_executed_by_packet" => false,
  "qemu_executed_by_packet" => false,
  "wine_executed_by_packet" => false,
  "host_root_modified" => false
)
blocked_promotion = write_json_fixture(
  "schema_version" => "xnix.runtime.full_checkpoint_promotion_packet.v1",
  "report_type" => "full-checkpoint-promotion-packet",
  "promotion_allowed" => false,
  "promotion_decision" => "blocked-incomplete-full-smoke-report",
  "full_smoke_state" => "blocked-incomplete-report",
  "formal_release_ready" => false,
  "operator_required_command" => "ruby scripts/full_smoke.rb",
  "full_smoke_executed_by_packet" => false,
  "docker_executed_by_packet" => false,
  "qemu_executed_by_packet" => false,
  "wine_executed_by_packet" => false,
  "host_root_modified" => false
)
preflight_smoke = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.desktop_trigger_request_preflight_smoke.v1",
  "report_type" => "desktop-trigger-request-preflight-smoke",
  "smoke_passed" => true,
  "preflight_smoke_state" => "passed",
  "blocked_preflight_state" => "blocked-missing-promotion",
  "ready_preflight_state" => "ready-for-operator-request",
  "owner_service_call_shape_verified" => true,
  "operator_request_ready" => true,
  "formal_release_ready" => false,
  "kde_receives_materialized_owner_args" => false,
  "service_call_dispatched" => false,
  "dbus_called" => false,
  "desktop_launch_enabled" => false,
  "backend_launch_enabled" => false,
  "runtime_state_written" => false,
  "kde_configuration_written" => false,
  "docker_executed" => false,
  "qemu_executed" => false,
  "wine_executed" => false,
  "colima_executed" => false,
  "network_checks_run" => false,
  "package_manager_invoked" => false,
  "host_root_modified" => false
)
q4_messagebox_smoke = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.scripts.q4_messagebox_smoke.v1",
  "request_type" => "q4-messagebox-smoke",
  "status" => "passed",
  "app_id" => "org.xnix.apps.messagebox",
  "display_name" => "Xnix MessageBox",
  "window_match" => "Xnix document opened by Windows app",
  "document_content_marker_observation_required" => true,
  "document_content_marker_observed" => true,
  "real_run_receipt_summary_ready" => true,
  "real_run_receipt_summary_file_open_verified" => true,
  "real_run_receipt_summary_document_content_marker_observed" => true,
  "real_run_acceptance_ready" => true,
  "real_run_acceptance_document_content_marker_observed" => true,
  "real_run_acceptance_center_projection_consumed" => true,
  "real_run_acceptance_kde_page_projection_consumed" => true,
  "owner_file_open_entrypoint_invoked" => true,
  "runtime_evidence_owner_file_open_entrypoint_invoked" => true,
  "go_owned_q4_winapp_acceptance_schema" => "xnix.runtime.q4_winapp_acceptance.v1",
  "go_owned_q4_winapp_acceptance_request_type" => "q4-winapp-acceptance-preview",
  "go_owned_q4_winapp_acceptance_ready" => true,
  "go_owned_q4_winapp_acceptance_consumed" => true,
  "go_owned_q4_winapp_acceptance_document_content_marker_observed" => true,
  "go_owned_q4_winapp_acceptance_path_exposed" => false,
  "go_owned_q4_winapp_acceptance_remote_host_exposed" => false,
  "go_owned_q4_winapp_acceptance_delegated_command_exposed" => false,
  "remote_executable_path_exposed" => false,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
)
known_existing_winapp_acceptance = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.known_existing_winapp_acceptance.v1",
  "request_type" => "known-existing-winapp-acceptance-preview",
  "acceptance_type" => "known-existing-windows-app-real-run-acceptance",
  "acceptance_ready" => true,
  "app_id" => "7zr",
  "display_name" => "7-Zip standalone console executable",
  "app_version" => "26.02",
  "architecture" => "windows-x86",
  "existing_windows_app" => true,
  "known_portable_catalog_backed" => true,
  "checksum_verified" => true,
  "marker_observed" => true,
  "guest_reachable" => true,
  "runtime_started_isolated_guest" => true,
  "isolated_guest_execution_observed" => true,
  "compatibility_engine_execution_observed" => true,
  "loopback_only_networking" => true,
  "serial_log_persisted" => true,
  "output_redacted" => true,
  "run_report_path_exposed" => false,
  "remote_host_exposed" => false,
  "guest_endpoint_exposed" => false,
  "raw_path_exposed" => false,
  "raw_output_exposed" => false,
  "runtime_argv_exposed" => false,
  "runner_path_exposed" => false,
  "network_required" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "docker_executed" => false,
  "colima_executed" => false,
  "network_checks_run" => false,
  "package_manager_invoked" => false
)
failed_known_existing_winapp_acceptance = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.known_existing_winapp_acceptance.v1",
  "request_type" => "known-existing-winapp-acceptance-preview",
  "acceptance_type" => "known-existing-windows-app-real-run-acceptance",
  "acceptance_ready" => false,
  "app_id" => "org.xnix.apps.messagebox",
  "existing_windows_app" => false,
  "known_portable_catalog_backed" => false,
  "checksum_verified" => true,
  "marker_observed" => false,
  "isolated_guest_execution_observed" => true,
  "compatibility_engine_execution_observed" => true,
  "output_redacted" => true,
  "run_report_path_exposed" => false,
  "remote_host_exposed" => false,
  "guest_endpoint_exposed" => false,
  "raw_path_exposed" => false,
  "raw_output_exposed" => false,
  "runtime_argv_exposed" => false,
  "runner_path_exposed" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "docker_executed" => false,
  "colima_executed" => false,
  "network_checks_run" => false,
  "package_manager_invoked" => false
)
known_app_matrix_evidence = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.known_app_matrix_evidence_preview.v1",
  "request_type" => "known-app-matrix-evidence-preview",
  "source" => "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
  "runtime_method" => "PreviewKnownAppMatrixEvidence",
  "read_method" => "GetKnownAppMatrixEvidence",
  "matrix_status" => "passed",
  "matrix_report_consumed" => true,
  "matrix_report_path_exposed" => false,
  "matrix_report_output_written" => true,
  "app_count" => 2,
  "passed_count" => 2,
  "failed_count" => 0,
  "evidence_count" => 2,
  "passed_evidence_count" => 2,
  "failed_evidence_count" => 0,
  "qemu_executed_count" => 2,
  "wine_executed_count" => 2,
  "marker_observed_count" => 2,
  "checksum_verified_count" => 2,
  "raw_output_redacted_count" => 2,
  "serial_log_evidence_count" => 2,
  "guest_started_count" => 2,
  "guest_port_auto_count" => 2,
  "compatibility_center_projection_ready" => true,
  "kde_center_projection_ready" => true,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "desktop_launch_enabled" => false,
  "backend_launch_enabled" => false,
  "action_execution_enabled" => false,
  "backend_details_exposed" => false,
  "raw_output_exposed" => false,
  "remote_path_exposed" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "apps" => [
    {
      "app_id" => "7zr",
      "display_name" => "7-Zip standalone console executable",
      "app_version" => "26.02",
      "executable_name" => "7zr.exe",
      "smoke_status" => "passed",
      "compatibility_state" => "real-qemu-wine-verified",
      "marker_observed" => true,
      "checksum_verified" => true,
      "qemu_executed" => true,
      "wine_executed" => true,
      "guest_started" => true,
      "guest_port_auto" => true,
      "raw_output_redacted" => true,
      "serial_log_evidence" => true,
      "report_evidence" => true,
      "runtime_owned" => true,
      "go_runtime_backed" => true,
      "kde_policy_owner" => false,
      "desktop_launch_enabled" => false,
      "backend_launch_enabled" => false,
      "backend_details_exposed" => false,
      "raw_output_exposed" => false,
      "remote_path_exposed" => false,
      "host_root_modified" => false
    },
    {
      "app_id" => "busybox-w32",
      "display_name" => "BusyBox-w32 standalone console executable",
      "app_version" => "current-2026-07-24",
      "executable_name" => "busybox.exe",
      "smoke_status" => "passed",
      "compatibility_state" => "real-qemu-wine-verified",
      "marker_observed" => true,
      "checksum_verified" => true,
      "qemu_executed" => true,
      "wine_executed" => true,
      "guest_started" => true,
      "guest_port_auto" => true,
      "raw_output_redacted" => true,
      "serial_log_evidence" => true,
      "report_evidence" => true,
      "runtime_owned" => true,
      "go_runtime_backed" => true,
      "kde_policy_owner" => false,
      "desktop_launch_enabled" => false,
      "backend_launch_enabled" => false,
      "backend_details_exposed" => false,
      "raw_output_exposed" => false,
      "remote_path_exposed" => false,
      "host_root_modified" => false
    }
  ]
)
failed_known_app_matrix_evidence = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.known_app_matrix_evidence_preview.v1",
  "request_type" => "known-app-matrix-evidence-preview",
  "matrix_status" => "passed",
  "matrix_report_consumed" => true,
  "matrix_report_path_exposed" => false,
  "matrix_report_output_written" => true,
  "app_count" => 1,
  "passed_count" => 1,
  "failed_count" => 0,
  "evidence_count" => 1,
  "passed_evidence_count" => 1,
  "qemu_executed_count" => 1,
  "wine_executed_count" => 1,
  "marker_observed_count" => 1,
  "checksum_verified_count" => 1,
  "raw_output_redacted_count" => 1,
  "serial_log_evidence_count" => 1,
  "compatibility_center_projection_ready" => false,
  "kde_center_projection_ready" => false,
  "runtime_owned" => true,
  "go_runtime_backed" => true,
  "kde_policy_owner" => false,
  "desktop_launch_enabled" => false,
  "backend_launch_enabled" => false,
  "action_execution_enabled" => false,
  "backend_details_exposed" => false,
  "raw_output_exposed" => false,
  "remote_path_exposed" => false,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false,
  "apps" => [
    {
      "app_id" => "7zr",
      "smoke_status" => "passed",
      "compatibility_state" => "real-qemu-wine-verified",
      "marker_observed" => true,
      "checksum_verified" => true,
      "qemu_executed" => true,
      "wine_executed" => true,
      "guest_started" => true,
      "guest_port_auto" => true,
      "raw_output_redacted" => true,
      "serial_log_evidence" => true,
      "report_evidence" => true,
      "runtime_owned" => true,
      "go_runtime_backed" => true,
      "kde_policy_owner" => false,
      "desktop_launch_enabled" => false,
      "backend_launch_enabled" => false,
      "backend_details_exposed" => false,
      "raw_output_exposed" => false,
      "remote_path_exposed" => false,
      "host_root_modified" => false
    }
  ]
)
failed_q4_messagebox_smoke = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.scripts.q4_messagebox_smoke.v1",
  "request_type" => "q4-messagebox-smoke",
  "status" => "passed",
  "app_id" => "org.xnix.apps.messagebox",
  "display_name" => "Xnix MessageBox",
  "window_match" => "Xnix Windows GUI Smoke",
  "document_content_marker_observation_required" => true,
  "document_content_marker_observed" => false,
  "real_run_receipt_summary_ready" => true,
  "real_run_receipt_summary_file_open_verified" => true,
  "real_run_receipt_summary_document_content_marker_observed" => false,
  "real_run_acceptance_ready" => false,
  "real_run_acceptance_document_content_marker_observed" => false,
  "real_run_acceptance_center_projection_consumed" => true,
  "real_run_acceptance_kde_page_projection_consumed" => true,
  "owner_file_open_entrypoint_invoked" => true,
  "runtime_evidence_owner_file_open_entrypoint_invoked" => true,
  "go_owned_q4_winapp_acceptance_schema" => "xnix.runtime.q4_winapp_acceptance.v1",
  "go_owned_q4_winapp_acceptance_request_type" => "q4-winapp-acceptance-preview",
  "go_owned_q4_winapp_acceptance_ready" => false,
  "go_owned_q4_winapp_acceptance_consumed" => true,
  "go_owned_q4_winapp_acceptance_document_content_marker_observed" => false,
  "go_owned_q4_winapp_acceptance_path_exposed" => false,
  "go_owned_q4_winapp_acceptance_remote_host_exposed" => false,
  "go_owned_q4_winapp_acceptance_delegated_command_exposed" => false,
  "remote_executable_path_exposed" => false,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
)
failed_preflight_smoke = write_json_fixture(
  "version" => File.read(project_root.join("VERSION")).strip,
  "schema_version" => "xnix.runtime.desktop_trigger_request_preflight_smoke.v1",
  "report_type" => "desktop-trigger-request-preflight-smoke",
  "smoke_passed" => false,
  "preflight_smoke_state" => "failed",
  "blocked_preflight_state" => "blocked-missing-promotion",
  "ready_preflight_state" => "blocked",
  "owner_service_call_shape_verified" => false,
  "operator_request_ready" => false,
  "formal_release_ready" => false,
  "kde_receives_materialized_owner_args" => false,
  "service_call_dispatched" => false,
  "dbus_called" => false,
  "desktop_launch_enabled" => false,
  "backend_launch_enabled" => false,
  "runtime_state_written" => false,
  "kde_configuration_written" => false,
  "docker_executed" => false,
  "qemu_executed" => false,
  "wine_executed" => false,
  "colima_executed" => false,
  "network_checks_run" => false,
  "package_manager_invoked" => false,
  "host_root_modified" => false
)
malformed_preflight_smoke = write_text_fixture("{not-json")
malformed_messagebox_smoke = write_text_fixture("{not-json")
malformed_known_existing_winapp_acceptance = write_text_fixture("{not-json")
malformed_known_app_matrix_evidence = write_text_fixture("{not-json")
malformed = nil

begin
  stdout, stderr, status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    known_existing_winapp_acceptance.path,
    "--known-app-matrix-evidence",
    known_app_matrix_evidence.path
  )
  assert(status.success?, "release evidence index JSON must exit successfully: #{stderr}")
  report = JSON.parse(stdout)

  assert(report.fetch("version") == File.read(project_root.join("VERSION")).strip, "release evidence index must expose the current version")
  assert(report.fetch("schema_version") == "xnix.runtime.release_evidence_index.v1", "release evidence index must expose the schema")
  assert(report.fetch("report_type") == "release-evidence-index", "release evidence index must identify its report type")
  preflight_status = report.fetch("desktop_trigger_request_preflight_smoke_status")
  assert(preflight_status.fetch("evidence_supplied"), "release evidence index must report supplied preflight smoke evidence")
  assert(preflight_status.fetch("smoke_passed"), "release evidence index must report passing preflight smoke evidence")
  assert(preflight_status.fetch("status") == "implemented", "release evidence index must expose implemented preflight smoke evidence")
  messagebox_status = report.fetch("q4_messagebox_smoke_status")
  assert(messagebox_status.fetch("evidence_supplied"), "release evidence index must report supplied q4 MessageBox evidence")
  assert(messagebox_status.fetch("smoke_passed"), "release evidence index must report passing q4 MessageBox evidence")
  assert(messagebox_status.fetch("status") == "implemented", "release evidence index must expose implemented q4 MessageBox evidence")
  known_existing_status = report.fetch("known_existing_winapp_acceptance_status")
  assert(known_existing_status.fetch("evidence_supplied"), "release evidence index must report supplied known existing Windows app acceptance evidence")
  assert(known_existing_status.fetch("acceptance_passed"), "release evidence index must report passing known existing Windows app acceptance evidence")
  assert(known_existing_status.fetch("status") == "implemented", "release evidence index must expose implemented known existing Windows app acceptance evidence")
  known_app_matrix_status = report.fetch("known_app_matrix_evidence_status")
  assert(known_app_matrix_status.fetch("evidence_supplied"), "release evidence index must report supplied known app matrix evidence")
  assert(known_app_matrix_status.fetch("matrix_passed"), "release evidence index must report passing known app matrix evidence")
  assert(known_app_matrix_status.fetch("status") == "implemented", "release evidence index must expose implemented known app matrix evidence")
  assert(report.fetch("runtime_owned"), "release evidence index must keep Runtime ownership explicit")
  assert(!report.fetch("go_runtime_backed"), "release evidence index must identify itself as a Ruby report, not Go Runtime business logic")
  assert(report.fetch("ruby_report_only"), "release evidence index must be marked Ruby report only")
  assert(!report.fetch("kde_policy_owner"), "release evidence index must not make KDE the policy owner")
  assert(report.fetch("offline_default"), "release evidence index must run in offline default mode")
  assert(!report.fetch("docker_executed"), "release evidence index must not run Docker")
  assert(!report.fetch("qemu_executed"), "release evidence index must not run QEMU")
  assert(!report.fetch("network_checks_run"), "release evidence index must not run network checks")
  assert(!report.fetch("package_manager_invoked"), "release evidence index must not invoke package managers")
  assert(!report.fetch("backend_launch_enabled"), "release evidence index must not launch backends")
  assert(!report.fetch("host_root_modified"), "release evidence index must not mutate the host root")
  assert(!report.fetch("automatic_staging_enabled"), "release evidence index must not stage automatically")
  assert(!report.fetch("automatic_release_tagging_enabled"), "release evidence index must not tag releases automatically")
  assert(!report.fetch("release_ready"), "release evidence index must keep release readiness gated")
  assert(report.fetch("human_authorization_required"), "release evidence index must flag human authorization requirements")

  counts = report.fetch("counts")
  assert(counts.fetch("implemented").positive?, "release evidence index must include implemented evidence")
  assert(counts.fetch("fixture_only").positive?, "release evidence index must include fixture-only evidence")
  assert(counts.fetch("contract_only").positive?, "release evidence index must include contract-only evidence")
  assert(counts.fetch("blocked").positive?, "release evidence index must include blocked evidence")
  assert(counts.fetch("skipped").positive?, "release evidence index must include skipped evidence")
  assert(counts.fetch("human_authorized").zero?, "completed product smoke must not retain human-authorized evidence classification")

  claims = report.fetch("claims").to_h { |claim| [claim.fetch("id"), claim] }
  assert(claims.fetch("runtime-owner-read-boundary").fetch("evidence_level") == "implemented", "runtime owner claim must be implemented")
  assert(claims.fetch("runtime-state-backend-lifecycle").fetch("evidence_level") == "fixture-only", "backend lifecycle claim must cover fixture-only evidence")
  assert(claims.fetch("portal-snapshot-safety").fetch("evidence_level") == "contract-only", "Portal snapshot claim must cover contract-only evidence")
  assert(claims.fetch("diagnostics-repair-ai-boundary").fetch("evidence_level") == "blocked", "diagnostics claim must cover blocked evidence")
  assert(claims.fetch("protected-claude-file").fetch("evidence_level") == "blocked", "protected file claim must block when modified")
  assert(claims.fetch("protected-claude-file").fetch("blockers").include?("protected-file"), "protected file claim must name protected-file blocker")
  assert(claims.fetch("unclassified-files").fetch("evidence_level") == "blocked", "unclassified file claim must block when files are unclassified")
  assert(claims.fetch("unclassified-files").fetch("blockers").include?("unclassified-file"), "unclassified file claim must name unclassified-file blocker")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("evidence_level") == "implemented", "product image claim must consume the authorized q4 evidence")
  assert(!claims.fetch("product-image-qemu-acceptance").fetch("human_authorization_required"), "completed product image evidence must not require retroactive authorization")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("blockers").empty?, "completed product image evidence must have no evidence blockers")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("formal_release_ready"), "product image claim must expose promotion readiness")
  assert(claims.fetch("product-image-qemu-acceptance").fetch("promotion_decision") == "promote", "product image claim must expose promotion decision")
  assert(claims.fetch("full-checkpoint-promotion").fetch("evidence_level") == "implemented", "promotion claim must be implemented when promotion packet allows release")
  assert(claims.fetch("full-checkpoint-promotion").fetch("formal_release_ready"), "promotion claim must expose formal readiness")
  assert(claims.fetch("desktop-trigger-request-preflight-smoke").fetch("evidence_level") == "implemented", "preflight smoke claim must be implemented when supplied evidence passes")
  assert(claims.fetch("desktop-trigger-request-preflight-smoke").fetch("smoke_passed"), "preflight smoke claim must expose smoke pass state")
  assert(claims.fetch("desktop-trigger-request-preflight-smoke").fetch("blocked_preflight_state") == "blocked-missing-promotion", "preflight smoke claim must preserve blocked state evidence")
  assert(claims.fetch("desktop-trigger-request-preflight-smoke").fetch("ready_preflight_state") == "ready-for-operator-request", "preflight smoke claim must preserve ready state evidence")
  assert(claims.fetch("q4-messagebox-document-smoke").fetch("evidence_level") == "implemented", "q4 MessageBox claim must be implemented when supplied evidence passes")
  assert(claims.fetch("q4-messagebox-document-smoke").fetch("smoke_passed"), "q4 MessageBox claim must expose smoke pass state")
  assert(claims.fetch("q4-messagebox-document-smoke").fetch("document_content_marker_observed"), "q4 MessageBox claim must expose document marker observation")
  assert(claims.fetch("q4-messagebox-document-smoke").fetch("go_owned_q4_winapp_acceptance_ready"), "q4 MessageBox claim must expose Go-owned q4 acceptance readiness")
  assert(claims.fetch("known-existing-winapp-acceptance").fetch("evidence_level") == "implemented", "known existing Windows app claim must be implemented when supplied evidence passes")
  assert(claims.fetch("known-existing-winapp-acceptance").fetch("acceptance_passed"), "known existing Windows app claim must expose acceptance pass state")
  assert(claims.fetch("known-existing-winapp-acceptance").fetch("app_id") == "7zr", "known existing Windows app claim must expose the accepted app id")
  assert(claims.fetch("known-existing-winapp-acceptance").fetch("isolated_guest_execution_observed"), "known existing Windows app claim must expose isolated guest execution evidence")
  assert(claims.fetch("known-existing-winapp-acceptance").fetch("compatibility_engine_execution_observed"), "known existing Windows app claim must expose compatibility engine evidence")
  assert(claims.fetch("known-app-matrix-evidence").fetch("evidence_level") == "implemented", "known app matrix claim must be implemented when supplied evidence passes")
  assert(claims.fetch("known-app-matrix-evidence").fetch("matrix_passed"), "known app matrix claim must expose matrix pass state")
  assert(claims.fetch("known-app-matrix-evidence").fetch("required_app_ids").sort == %w[7zr busybox-w32].sort, "known app matrix claim must expose required app ids")
  assert(claims.fetch("known-app-matrix-evidence").fetch("observed_required_app_ids").sort == %w[7zr busybox-w32].sort, "known app matrix claim must expose observed required app ids")
  assert(claims.fetch("known-app-matrix-evidence").fetch("qemu_executed_count") == 2, "known app matrix claim must expose QEMU execution count")
  assert(claims.fetch("known-app-matrix-evidence").fetch("wine_executed_count") == 2, "known app matrix claim must expose Wine execution count")
  assert(claims.fetch("restricted-heavy-smoke-skipped").fetch("evidence_level") == "skipped", "heavy smoke claim must be skipped by default")
  assert(claims.fetch("kde-first-presence").fetch("evidence_level") == "implemented", "KDE presence claim must be implemented when seven entry points are present")
  assert(claims.fetch("contract-drift").fetch("evidence_level") == "implemented", "contract drift claim must be implemented when no drift is reported")
  assert(claims.values.all? { |claim| claim.fetch("unsafe_gates").values.none? }, "every claim must keep unsafe gates disabled")

  markdown, markdown_stderr, markdown_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "markdown",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    known_existing_winapp_acceptance.path,
    "--known-app-matrix-evidence",
    known_app_matrix_evidence.path
  )
  assert(markdown_status.success?, "release evidence index Markdown must exit successfully: #{markdown_stderr}")
  assert(markdown.include?("# Release Evidence Index"), "Markdown report must include a title")
  assert(markdown.include?("runtime-owner-read-boundary"), "Markdown report must include runtime owner claim")
  assert(markdown.include?("product-image-qemu-acceptance"), "Markdown report must include product image claim")
  assert(markdown.include?("full-checkpoint-promotion"), "Markdown report must include promotion claim")
  assert(markdown.include?("desktop-trigger-request-preflight-smoke"), "Markdown report must include preflight smoke claim")
  assert(markdown.include?("q4-messagebox-document-smoke"), "Markdown report must include q4 MessageBox claim")
  assert(markdown.include?("known-existing-winapp-acceptance"), "Markdown report must include known existing Windows app claim")
  assert(markdown.include?("known-app-matrix-evidence"), "Markdown report must include known app matrix claim")

  missing_preflight_stdout, missing_preflight_stderr, missing_preflight_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path
  )
  assert(missing_preflight_status.success?, "missing preflight smoke evidence case must still emit a release index: #{missing_preflight_stderr}")
  missing_preflight_report = JSON.parse(missing_preflight_stdout)
  missing_preflight_claim = missing_preflight_report.fetch("claims").find { |claim| claim.fetch("id") == "desktop-trigger-request-preflight-smoke" }
  assert(missing_preflight_report.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status") == "not-supplied", "missing preflight smoke evidence must be visible")
  assert(missing_preflight_claim.fetch("evidence_level") == "skipped", "missing preflight smoke evidence must be skipped, not executed")

  missing_messagebox_stdout, missing_messagebox_stderr, missing_messagebox_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path
  )
  assert(missing_messagebox_status.success?, "missing q4 MessageBox evidence case must still emit a release index: #{missing_messagebox_stderr}")
  missing_messagebox_report = JSON.parse(missing_messagebox_stdout)
  missing_messagebox_claim = missing_messagebox_report.fetch("claims").find { |claim| claim.fetch("id") == "q4-messagebox-document-smoke" }
  assert(missing_messagebox_report.fetch("q4_messagebox_smoke_status").fetch("status") == "not-supplied", "missing q4 MessageBox evidence must be visible")
  assert(missing_messagebox_claim.fetch("evidence_level") == "skipped", "missing q4 MessageBox evidence must be skipped, not executed")

  failed_preflight_stdout, failed_preflight_stderr, failed_preflight_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    failed_preflight_smoke.path
  )
  assert(failed_preflight_status.success?, "failed preflight smoke evidence case must still emit a release index: #{failed_preflight_stderr}")
  failed_preflight_report = JSON.parse(failed_preflight_stdout)
  failed_preflight_claim = failed_preflight_report.fetch("claims").find { |claim| claim.fetch("id") == "desktop-trigger-request-preflight-smoke" }
  assert(failed_preflight_report.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status") == "blocked", "failed preflight smoke evidence must be blocked")
  assert(failed_preflight_claim.fetch("evidence_level") == "blocked", "failed preflight smoke claim must be blocked")
  assert(failed_preflight_claim.fetch("blockers").include?("desktop-trigger-request-preflight-smoke-not-passed"), "failed preflight smoke claim must name the failed smoke blocker")

  failed_messagebox_stdout, failed_messagebox_stderr, failed_messagebox_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    failed_q4_messagebox_smoke.path
  )
  assert(failed_messagebox_status.success?, "failed q4 MessageBox evidence case must still emit a release index: #{failed_messagebox_stderr}")
  failed_messagebox_report = JSON.parse(failed_messagebox_stdout)
  failed_messagebox_claim = failed_messagebox_report.fetch("claims").find { |claim| claim.fetch("id") == "q4-messagebox-document-smoke" }
  assert(failed_messagebox_report.fetch("q4_messagebox_smoke_status").fetch("status") == "blocked", "failed q4 MessageBox evidence must be blocked")
  assert(failed_messagebox_claim.fetch("evidence_level") == "blocked", "failed q4 MessageBox claim must be blocked")
  assert(failed_messagebox_claim.fetch("blockers").include?("q4-messagebox-document-marker-missing"), "failed q4 MessageBox claim must name document marker blocker")

  failed_known_existing_stdout, failed_known_existing_stderr, failed_known_existing_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    failed_known_existing_winapp_acceptance.path
  )
  assert(failed_known_existing_status.success?, "failed known existing Windows app acceptance evidence case must still emit a release index: #{failed_known_existing_stderr}")
  failed_known_existing_report = JSON.parse(failed_known_existing_stdout)
  failed_known_existing_claim = failed_known_existing_report.fetch("claims").find { |claim| claim.fetch("id") == "known-existing-winapp-acceptance" }
  assert(failed_known_existing_report.fetch("known_existing_winapp_acceptance_status").fetch("status") == "blocked", "failed known existing Windows app acceptance evidence must be blocked")
  assert(failed_known_existing_claim.fetch("evidence_level") == "blocked", "failed known existing Windows app claim must be blocked")
  assert(failed_known_existing_claim.fetch("blockers").include?("known-existing-winapp-identity-missing"), "failed known existing Windows app claim must name identity blocker")

  failed_matrix_stdout, failed_matrix_stderr, failed_matrix_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    known_existing_winapp_acceptance.path,
    "--known-app-matrix-evidence",
    failed_known_app_matrix_evidence.path
  )
  assert(failed_matrix_status.success?, "failed known app matrix evidence case must still emit a release index: #{failed_matrix_stderr}")
  failed_matrix_report = JSON.parse(failed_matrix_stdout)
  failed_matrix_claim = failed_matrix_report.fetch("claims").find { |claim| claim.fetch("id") == "known-app-matrix-evidence" }
  assert(failed_matrix_report.fetch("known_app_matrix_evidence_status").fetch("status") == "blocked", "failed known app matrix evidence must be blocked")
  assert(failed_matrix_claim.fetch("evidence_level") == "blocked", "failed known app matrix claim must be blocked")
  assert(failed_matrix_claim.fetch("blockers").include?("known-app-matrix-required-app-missing"), "failed known app matrix claim must name required app blocker")

  malformed_preflight_stdout, malformed_preflight_stderr, malformed_preflight_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    malformed_preflight_smoke.path
  )
  assert(malformed_preflight_status.success?, "malformed preflight smoke evidence case must still emit a release index: #{malformed_preflight_stderr}")
  malformed_preflight_report = JSON.parse(malformed_preflight_stdout)
  malformed_preflight_claim = malformed_preflight_report.fetch("claims").find { |claim| claim.fetch("id") == "desktop-trigger-request-preflight-smoke" }
  assert(malformed_preflight_report.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status") == "blocked", "malformed preflight smoke evidence must be blocked")
  assert(malformed_preflight_claim.fetch("evidence_level") == "blocked", "malformed preflight smoke claim must be blocked")
  assert(malformed_preflight_claim.fetch("blockers").any? { |blocker| blocker.include?("malformed-report") }, "malformed preflight smoke claim must name malformed-report blocker")

  malformed_messagebox_stdout, malformed_messagebox_stderr, malformed_messagebox_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    malformed_messagebox_smoke.path
  )
  assert(malformed_messagebox_status.success?, "malformed q4 MessageBox evidence case must still emit a release index: #{malformed_messagebox_stderr}")
  malformed_messagebox_report = JSON.parse(malformed_messagebox_stdout)
  malformed_messagebox_claim = malformed_messagebox_report.fetch("claims").find { |claim| claim.fetch("id") == "q4-messagebox-document-smoke" }
  assert(malformed_messagebox_report.fetch("q4_messagebox_smoke_status").fetch("status") == "blocked", "malformed q4 MessageBox evidence must be blocked")
  assert(malformed_messagebox_claim.fetch("evidence_level") == "blocked", "malformed q4 MessageBox claim must be blocked")
  assert(malformed_messagebox_claim.fetch("blockers").any? { |blocker| blocker.include?("malformed-report") }, "malformed q4 MessageBox claim must name malformed-report blocker")

  malformed_known_existing_stdout, malformed_known_existing_stderr, malformed_known_existing_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    malformed_known_existing_winapp_acceptance.path
  )
  assert(malformed_known_existing_status.success?, "malformed known existing Windows app evidence case must still emit a release index: #{malformed_known_existing_stderr}")
  malformed_known_existing_report = JSON.parse(malformed_known_existing_stdout)
  malformed_known_existing_claim = malformed_known_existing_report.fetch("claims").find { |claim| claim.fetch("id") == "known-existing-winapp-acceptance" }
  assert(malformed_known_existing_report.fetch("known_existing_winapp_acceptance_status").fetch("status") == "blocked", "malformed known existing Windows app evidence must be blocked")
  assert(malformed_known_existing_claim.fetch("evidence_level") == "blocked", "malformed known existing Windows app claim must be blocked")
  assert(malformed_known_existing_claim.fetch("blockers").any? { |blocker| blocker.include?("malformed-report") }, "malformed known existing Windows app claim must name malformed-report blocker")

  malformed_matrix_stdout, malformed_matrix_stderr, malformed_matrix_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path,
    "--q4-messagebox-smoke",
    q4_messagebox_smoke.path,
    "--known-existing-winapp-acceptance",
    known_existing_winapp_acceptance.path,
    "--known-app-matrix-evidence",
    malformed_known_app_matrix_evidence.path
  )
  assert(malformed_matrix_status.success?, "malformed known app matrix evidence case must still emit a release index: #{malformed_matrix_stderr}")
  malformed_matrix_report = JSON.parse(malformed_matrix_stdout)
  malformed_matrix_claim = malformed_matrix_report.fetch("claims").find { |claim| claim.fetch("id") == "known-app-matrix-evidence" }
  assert(malformed_matrix_report.fetch("known_app_matrix_evidence_status").fetch("status") == "blocked", "malformed known app matrix evidence must be blocked")
  assert(malformed_matrix_claim.fetch("evidence_level") == "blocked", "malformed known app matrix claim must be blocked")
  assert(malformed_matrix_claim.fetch("blockers").any? { |blocker| blocker.include?("malformed-report") }, "malformed known app matrix claim must name malformed-report blocker")

  blocked_stdout, blocked_stderr, blocked_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    implementation.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    blocked_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path
  )
  assert(blocked_status.success?, "blocked promotion case must still emit a release index: #{blocked_stderr}")
  blocked_report = JSON.parse(blocked_stdout)
  blocked_claims = blocked_report.fetch("claims").to_h { |claim| [claim.fetch("id"), claim] }
  blocked_product_claim = blocked_claims.fetch("product-image-qemu-acceptance")
  assert(blocked_product_claim.fetch("historical_product_smoke_evidence_passed"), "blocked promotion case must retain historical q4 evidence")
  assert(blocked_product_claim.fetch("evidence_level") == "blocked", "blocked promotion must block current product image acceptance")
  assert(blocked_product_claim.fetch("blockers").include?("full-checkpoint-promotion-not-allowed"), "blocked product image claim must name promotion blocker")
  assert(blocked_product_claim.fetch("promotion_decision") == "blocked-incomplete-full-smoke-report", "blocked product image claim must expose denied promotion decision")
  assert(!blocked_product_claim.fetch("formal_release_ready"), "blocked product image claim must not claim formal readiness")
  assert(blocked_claims.fetch("full-checkpoint-promotion").fetch("evidence_level") == "blocked", "blocked promotion claim must be blocked")
  assert(blocked_claims.fetch("full-checkpoint-promotion").fetch("blockers").include?("full-checkpoint-promotion-not-allowed"), "blocked promotion claim must name promotion blocker")

  malformed = write_text_fixture("{not-json")
  malformed_stdout, malformed_stderr, malformed_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--implementation-report",
    malformed.path,
    "--contract-drift-report",
    contract_drift.path,
    "--mainline-review",
    mainline.path,
    "--kde-smoke-report",
    kde_smoke.path,
    "--full-checkpoint-promotion",
    completed_promotion.path,
    "--desktop-trigger-request-preflight-smoke",
    preflight_smoke.path
  )
  assert(malformed_status.success?, "malformed report case must still emit a release index: #{malformed_stderr}")
  malformed_report = JSON.parse(malformed_stdout)
  assert(malformed_report.fetch("malformed_report_detected"), "release evidence index must flag malformed reports")
  malformed_claims = malformed_report.fetch("claims").to_h { |claim| [claim.fetch("id"), claim] }
  assert(malformed_claims.fetch("report-integrity").fetch("evidence_level") == "blocked", "report integrity claim must block malformed reports")
ensure
  [implementation, contract_drift, mainline, kde_smoke, completed_promotion, blocked_promotion, preflight_smoke, q4_messagebox_smoke, known_existing_winapp_acceptance, known_app_matrix_evidence, failed_preflight_smoke, failed_q4_messagebox_smoke, failed_known_existing_winapp_acceptance, failed_known_app_matrix_evidence, malformed_preflight_smoke, malformed_messagebox_smoke, malformed_known_existing_winapp_acceptance, malformed_known_app_matrix_evidence, malformed].compact.each do |file|
    file.close
    file.unlink
  end
end
