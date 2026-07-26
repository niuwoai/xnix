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
  file = Tempfile.new("xnix-merge-readiness-packet")
  file.write(JSON.pretty_generate(payload))
  file.flush
  file
end

def write_text_fixture(text)
  file = Tempfile.new("xnix-merge-readiness-packet")
  file.write(text)
  file.flush
  file
end

def known_app_matrix_payload(version)
  app_defaults = {
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
  {
    "version" => version,
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
      app_defaults.merge(
        "app_id" => "7zr",
        "display_name" => "7-Zip standalone console executable",
        "app_version" => "26.02",
        "executable_name" => "7zr.exe"
      ),
      app_defaults.merge(
        "app_id" => "busybox-w32",
        "display_name" => "BusyBox-w32 standalone console executable",
        "app_version" => "current-2026-07-24",
        "executable_name" => "busybox.exe"
      )
    ]
  }
end

project_root = Pathname.new(__dir__).join("..").realpath
script = project_root.join("scripts/merge_readiness_packet.rb")
version = File.read(project_root.join("VERSION")).strip

layout = write_text_fixture("PASS: Xnix fixture scaffold is consistent\n")
implementation = write_json_fixture(
  "domains" => [
    { "id" => "developer-verification-harness", "status" => "fixture-implemented" }
  ]
)
contract_drift = write_json_fixture("drift_detected" => false)
kde_smoke = write_json_fixture("entrypoint_count" => 7)
mainline = write_json_fixture(
  "changed_file_count" => 3,
  "unclassified_file_count" => 0,
  "protected_claude_file_modified" => false,
  "risky_file_count" => 0,
  "code_change_count" => 2,
  "lanes" => [
    {
      "id" => "cw10-evidence-drift-harness",
      "workstream" => "CW10 / A8",
      "title" => "Evidence and drift harness",
      "file_count" => 2,
      "statuses" => [" M"],
      "required_verification" => [
        "ruby scripts/merge_readiness_packet.rb --format json",
        "ruby scripts/verify_layout.rb"
      ],
      "safety_guards" => ["read-only-report"]
    },
    {
      "id" => "unclassified",
      "workstream" => "unknown",
      "title" => "Unclassified",
      "file_count" => 0,
      "statuses" => [],
      "required_verification" => [],
      "safety_guards" => ["manual-classification-required"]
    }
  ]
)
release_evidence = write_json_fixture("claims" => [], "claim_count" => 0)
completed_release_evidence = write_json_fixture(
  "claims" => [
    { "id" => "product-image-qemu-acceptance", "evidence_level" => "implemented", "blockers" => [] }
  ],
  "claim_count" => 1
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
fixture_matrix = write_json_fixture("row_count" => 7, "counts" => { "blocked" => 1 })
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
  "service_call_dispatched" => false,
  "dbus_called" => false,
  "backend_launch_enabled" => false,
  "host_root_modified" => false
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
  "service_call_dispatched" => false,
  "dbus_called" => false,
  "backend_launch_enabled" => false,
  "host_root_modified" => false
)
q4_sample_notepad_smoke = write_json_fixture(
  "schema_version" => "xnix.scripts.q4_sample_notepad_smoke.v1",
  "request_type" => "q4-sample-notepad-smoke",
  "version" => File.read(project_root.join("VERSION")).strip,
  "status" => "passed",
  "app_id" => "org.xnix.sample.notepad",
  "launch_mode" => "owner-controlled-launch",
  "file_open_entrypoint_requested" => true,
  "real_run_acceptance_required" => true,
  "go_owned_q4_sample_notepad_acceptance_schema" => "xnix.runtime.q4_sample_notepad_acceptance.v1",
  "go_owned_q4_sample_notepad_acceptance_request_type" => "q4-sample-notepad-acceptance-preview",
  "go_owned_q4_sample_notepad_acceptance_ready" => true,
  "go_owned_q4_sample_notepad_acceptance_consumed" => true,
  "go_owned_q4_sample_notepad_acceptance_path_exposed" => false,
  "go_owned_q4_sample_notepad_acceptance_remote_host_exposed" => false,
  "go_owned_q4_sample_notepad_acceptance_delegated_command_exposed" => false,
  "real_run_acceptance_ready" => true,
  "real_run_acceptance_center_projection_consumed" => true,
  "real_run_acceptance_kde_page_projection_consumed" => true,
  "real_run_receipt_summary_ready" => true,
  "real_run_receipt_summary_file_open_verified" => true,
  "window_match_observed" => true,
  "owner_file_open_entrypoint_invoked" => true,
  "runtime_evidence_owner_file_open_entrypoint_invoked" => true,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
)
q4_messagebox_smoke = write_json_fixture(
  "schema_version" => "xnix.scripts.q4_messagebox_smoke.v1",
  "request_type" => "q4-messagebox-smoke",
  "version" => File.read(project_root.join("VERSION")).strip,
  "status" => "passed",
  "app_id" => "org.xnix.apps.messagebox",
  "display_name" => "Xnix MessageBox",
  "window_match" => "Xnix document opened by Windows app",
  "launch_mode" => "owner-controlled-launch",
  "file_open_entrypoint_requested" => true,
  "real_run_acceptance_required" => true,
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
  "schema_version" => "xnix.runtime.known_existing_winapp_acceptance.v1",
  "request_type" => "known-existing-winapp-acceptance-preview",
  "version" => File.read(project_root.join("VERSION")).strip,
  "acceptance_type" => "known-existing-windows-app-real-run-acceptance",
  "acceptance_ready" => true,
  "app_id" => "7zr",
  "display_name" => "7-Zip standalone console executable",
  "app_version" => "26.02",
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
  "schema_version" => "xnix.runtime.known_existing_winapp_acceptance.v1",
  "request_type" => "known-existing-winapp-acceptance-preview",
  "version" => File.read(project_root.join("VERSION")).strip,
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
known_app_matrix_evidence = write_json_fixture(known_app_matrix_payload(version))
failed_known_app_matrix_payload = known_app_matrix_payload(version)
failed_known_app_matrix_payload["app_count"] = 1
failed_known_app_matrix_payload["passed_count"] = 1
failed_known_app_matrix_payload["evidence_count"] = 1
failed_known_app_matrix_payload["passed_evidence_count"] = 1
failed_known_app_matrix_payload["qemu_executed_count"] = 1
failed_known_app_matrix_payload["wine_executed_count"] = 1
failed_known_app_matrix_payload["marker_observed_count"] = 1
failed_known_app_matrix_payload["checksum_verified_count"] = 1
failed_known_app_matrix_payload["raw_output_redacted_count"] = 1
failed_known_app_matrix_payload["serial_log_evidence_count"] = 1
failed_known_app_matrix_payload["compatibility_center_projection_ready"] = false
failed_known_app_matrix_payload["kde_center_projection_ready"] = false
failed_known_app_matrix_payload["apps"] = failed_known_app_matrix_payload.fetch("apps").first(1)
failed_known_app_matrix_evidence = write_json_fixture(failed_known_app_matrix_payload)
failed_q4_messagebox_smoke = write_json_fixture(
  "schema_version" => "xnix.scripts.q4_messagebox_smoke.v1",
  "request_type" => "q4-messagebox-smoke",
  "version" => File.read(project_root.join("VERSION")).strip,
  "status" => "passed",
  "app_id" => "org.xnix.apps.messagebox",
  "display_name" => "Xnix MessageBox",
  "window_match" => "Xnix Windows GUI Smoke",
  "launch_mode" => "owner-controlled-launch",
  "file_open_entrypoint_requested" => true,
  "real_run_acceptance_required" => true,
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
failed_q4_sample_notepad_smoke = write_json_fixture(
  "schema_version" => "xnix.scripts.q4_sample_notepad_smoke.v1",
  "request_type" => "q4-sample-notepad-smoke",
  "version" => File.read(project_root.join("VERSION")).strip,
  "status" => "blocked",
  "app_id" => "org.xnix.sample.notepad",
  "launch_mode" => "owner-controlled-launch",
  "file_open_entrypoint_requested" => true,
  "real_run_acceptance_required" => true,
  "go_owned_q4_sample_notepad_acceptance_schema" => "xnix.runtime.q4_sample_notepad_acceptance.v1",
  "go_owned_q4_sample_notepad_acceptance_request_type" => "q4-sample-notepad-acceptance-preview",
  "go_owned_q4_sample_notepad_acceptance_ready" => false,
  "go_owned_q4_sample_notepad_acceptance_consumed" => false,
  "go_owned_q4_sample_notepad_acceptance_path_exposed" => false,
  "go_owned_q4_sample_notepad_acceptance_remote_host_exposed" => false,
  "go_owned_q4_sample_notepad_acceptance_delegated_command_exposed" => false,
  "real_run_acceptance_ready" => false,
  "real_run_acceptance_center_projection_consumed" => false,
  "real_run_acceptance_kde_page_projection_consumed" => false,
  "real_run_receipt_summary_ready" => false,
  "real_run_receipt_summary_file_open_verified" => false,
  "window_match_observed" => false,
  "owner_file_open_entrypoint_invoked" => false,
  "runtime_evidence_owner_file_open_entrypoint_invoked" => false,
  "host_compilation_avoided" => true,
  "host_root_modified" => false,
  "privileged_container_required" => false,
  "host_networking_required" => false,
  "docker_socket_mounted" => false,
  "broad_host_mount_required" => false
)
malformed = write_text_fixture("{not-json")
protected_mainline = write_json_fixture(
  "changed_file_count" => 2,
  "unclassified_file_count" => 1,
  "protected_claude_file_modified" => true,
  "risky_file_count" => 0,
  "code_change_count" => 1,
  "lanes" => []
)
unsafe_kde = write_json_fixture("entrypoint_count" => 7, "docker_executed" => true)

base_args = [
  "--layout-report", layout.path,
  "--implementation-report", implementation.path,
  "--contract-drift-report", contract_drift.path,
  "--kde-smoke-report", kde_smoke.path,
  "--mainline-review", mainline.path,
  "--release-evidence-index", release_evidence.path,
  "--full-checkpoint-promotion", blocked_promotion.path,
  "--desktop-trigger-request-preflight-smoke", preflight_smoke.path,
  "--q4-sample-notepad-smoke", q4_sample_notepad_smoke.path,
  "--q4-messagebox-smoke", q4_messagebox_smoke.path,
  "--known-existing-winapp-acceptance", known_existing_winapp_acceptance.path,
  "--known-app-matrix-evidence", known_app_matrix_evidence.path,
  "--fixture-matrix-report", fixture_matrix.path
]

begin
  stdout, stderr, status = Open3.capture3("ruby", script.to_s, "--format", "json", *base_args)
  assert(status.success?, "merge readiness packet JSON must exit successfully: #{stderr}")
  packet = JSON.parse(stdout)

  assert(packet.fetch("version") == File.read(project_root.join("VERSION")).strip, "packet must expose current version")
  assert(packet.fetch("schema_version") == "xnix.runtime.merge_readiness_packet.v1", "packet must expose schema")
  assert(packet.fetch("report_type") == "merge-readiness-packet", "packet must identify report type")
  assert(packet.fetch("runtime_owned"), "packet must keep Runtime ownership explicit")
  assert(packet.fetch("ruby_report_only"), "packet must identify itself as Ruby report-only tooling")
  assert(!packet.fetch("kde_policy_owner"), "packet must not make KDE the policy owner")
  assert(packet.fetch("offline_only"), "packet must default to offline-only")
  assert(packet.fetch("merge_ready"), "packet must be merge-ready when only restricted smoke remains")
  assert(!packet.fetch("release_ready"), "packet must keep release readiness gated")
  assert(packet.fetch("release_blocking_reasons").include?("restricted-docker-or-qemu-smoke-requires-human-authorization"), "packet must keep restricted Docker/QEMU as release blocker")
  assert(packet.fetch("release_blocking_reasons").include?("full-checkpoint-promotion-not-allowed"), "packet must include the full checkpoint promotion blocker")
  assert(packet.fetch("release_blocking_reasons").include?("production-runtime-and-windows-execution-remain-disabled"), "packet must keep production execution as a product release blocker")
  assert(packet.fetch("merge_blocking_reasons").empty?, "packet must not merge-block the clean fixture case")
  promotion_status = packet.fetch("full_checkpoint_promotion_status")
  assert(!promotion_status.fetch("promotion_allowed"), "packet must expose denied promotion status")
  assert(promotion_status.fetch("promotion_decision") == "blocked-incomplete-full-smoke-report", "packet must expose promotion decision")
  assert(promotion_status.fetch("operator_required_command") == "ruby scripts/full_smoke.rb", "packet must expose promotion operator command")
  preflight_status = packet.fetch("desktop_trigger_request_preflight_smoke_status")
  assert(preflight_status.fetch("evidence_supplied"), "packet must report supplied desktop-trigger request preflight smoke evidence")
  assert(preflight_status.fetch("smoke_passed"), "packet must report passing desktop-trigger request preflight smoke evidence")
  assert(preflight_status.fetch("status") == "passed", "packet must expose preflight smoke status")
  q4_sample_status = packet.fetch("q4_sample_notepad_smoke_status")
  assert(q4_sample_status.fetch("evidence_supplied"), "packet must report supplied q4 Sample Notepad acceptance evidence")
  assert(q4_sample_status.fetch("smoke_passed"), "packet must report passing q4 Sample Notepad acceptance evidence")
  assert(q4_sample_status.fetch("status") == "passed", "packet must expose q4 Sample Notepad acceptance status")
  assert(!packet.fetch("release_blocking_reasons").include?("q4-sample-notepad-acceptance-smoke-not-passed"), "passing q4 Sample Notepad evidence must close its release blocker")
  q4_messagebox_status = packet.fetch("q4_messagebox_smoke_status")
  assert(q4_messagebox_status.fetch("evidence_supplied"), "packet must report supplied q4 MessageBox document evidence")
  assert(q4_messagebox_status.fetch("smoke_passed"), "packet must report passing q4 MessageBox document evidence")
  assert(q4_messagebox_status.fetch("status") == "passed", "packet must expose q4 MessageBox document status")
  assert(!packet.fetch("release_blocking_reasons").include?("q4-messagebox-document-content-smoke-not-passed"), "passing q4 MessageBox document evidence must close its release blocker")
  known_existing_status = packet.fetch("known_existing_winapp_acceptance_status")
  assert(known_existing_status.fetch("evidence_supplied"), "packet must report supplied known existing Windows app acceptance evidence")
  assert(known_existing_status.fetch("acceptance_passed"), "packet must report passing known existing Windows app acceptance evidence")
  assert(known_existing_status.fetch("status") == "passed", "packet must expose known existing Windows app acceptance status")
  assert(!packet.fetch("release_blocking_reasons").include?("known-existing-winapp-acceptance-not-passed"), "passing known existing Windows app evidence must close its release blocker")
  known_app_matrix_status = packet.fetch("known_app_matrix_evidence_status")
  assert(known_app_matrix_status.fetch("evidence_supplied"), "packet must report supplied known app matrix evidence")
  assert(known_app_matrix_status.fetch("matrix_passed"), "packet must report passing known app matrix evidence")
  assert(known_app_matrix_status.fetch("status") == "passed", "packet must expose known app matrix status")
  assert(known_app_matrix_status.fetch("required_app_ids").sort == %w[7zr busybox-w32].sort, "packet must expose required matrix app ids")
  assert(!packet.fetch("release_blocking_reasons").include?("known-app-matrix-evidence-not-passed"), "passing known app matrix evidence must close its release blocker")

  tool_statuses = packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(tool_statuses.values.all? { |tool| tool.fetch("status") == "pass" }, "all fixture-backed tools must pass")
  assert(tool_statuses.fetch("layout").fetch("command") == "fixture:layout", "fixture command must not expose fixture path")
  assert(tool_statuses.fetch("full_checkpoint_promotion").fetch("command") == "fixture:full_checkpoint_promotion", "promotion fixture command must not expose fixture path")
  assert(tool_statuses.fetch("desktop_trigger_request_preflight_smoke").fetch("command") == "fixture:desktop_trigger_request_preflight_smoke", "preflight smoke fixture command must not expose fixture path")
  assert(tool_statuses.fetch("q4_sample_notepad_smoke").fetch("command") == "fixture:q4_sample_notepad_smoke", "q4 Sample Notepad fixture command must not expose fixture path")
  assert(tool_statuses.fetch("q4_messagebox_smoke").fetch("command") == "fixture:q4_messagebox_smoke", "q4 MessageBox fixture command must not expose fixture path")
  assert(tool_statuses.fetch("known_existing_winapp_acceptance").fetch("command") == "fixture:known_existing_winapp_acceptance", "known existing Windows app fixture command must not expose fixture path")
  assert(tool_statuses.fetch("known_app_matrix_evidence").fetch("command") == "fixture:known_app_matrix_evidence", "known app matrix fixture command must not expose fixture path")
  assert(packet.fetch("changed_file_counts").fetch("total") == 3, "packet must include changed file counts")
  assert(packet.fetch("lane_classification").any? { |lane| lane.fetch("id") == "cw10-evidence-drift-harness" }, "packet must include lane classification")
  assert(packet.fetch("required_follow_up_commands").include?("ruby scripts/verify_layout.rb"), "packet must include follow-up commands")
  unsafe_status = packet.fetch("unsafe_operation_status")
  assert(!unsafe_status.fetch("docker_executed"), "packet itself must not run Docker")
  assert(!unsafe_status.fetch("qemu_executed"), "packet itself must not run QEMU")
  assert(!unsafe_status.fetch("network_checks_run"), "packet itself must not run network checks")
  assert(!unsafe_status.fetch("package_manager_invoked"), "packet itself must not invoke package managers")
  assert(!unsafe_status.fetch("host_root_modified"), "packet itself must not mutate host root")
  assert(unsafe_status.fetch("detected_unsafe_findings").empty?, "clean packet must not report unsafe findings")

  markdown, markdown_stderr, markdown_status = Open3.capture3("ruby", script.to_s, "--format", "markdown", *base_args)
  assert(markdown_status.success?, "merge readiness packet Markdown must exit successfully: #{markdown_stderr}")
  assert(markdown.include?("# Merge Readiness Packet"), "Markdown must include title")
  assert(markdown.include?("cw10-evidence-drift-harness"), "Markdown must include lane classification")
  assert(markdown.include?("restricted-docker-or-qemu-smoke-requires-human-authorization"), "Markdown must include release blocker")
  assert(markdown.include?("Full checkpoint promotion: blocked-incomplete-full-smoke-report"), "Markdown must include promotion decision")
  assert(markdown.include?("Desktop-trigger request preflight smoke: passed"), "Markdown must include preflight smoke status")
  assert(markdown.include?("q4 Sample Notepad acceptance smoke: passed"), "Markdown must include q4 Sample Notepad acceptance status")
  assert(markdown.include?("q4 MessageBox document smoke: passed"), "Markdown must include q4 MessageBox document status")
  assert(markdown.include?("Known existing Windows app acceptance: passed"), "Markdown must include known existing Windows app acceptance status")
  assert(markdown.include?("Known Windows app matrix evidence: passed"), "Markdown must include known app matrix status")

  no_preflight_args = base_args.each_slice(2).reject { |option, _path| option == "--desktop-trigger-request-preflight-smoke" }.flatten
  no_preflight_stdout, no_preflight_stderr, no_preflight_status = Open3.capture3("ruby", script.to_s, "--format", "json", *no_preflight_args)
  assert(no_preflight_status.success?, "missing preflight smoke evidence case must still emit packet: #{no_preflight_stderr}")
  no_preflight_packet = JSON.parse(no_preflight_stdout)
  assert(no_preflight_packet.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status") == "not-supplied", "missing preflight smoke evidence must be non-blocking")
  assert(!no_preflight_packet.fetch("release_blocking_reasons").include?("desktop-trigger-request-preflight-smoke-not-passed"), "missing preflight smoke evidence must not block release by itself")

  no_q4_args = base_args.each_slice(2).reject { |option, _path| option == "--q4-sample-notepad-smoke" }.flatten
  no_q4_stdout, no_q4_stderr, no_q4_status = Open3.capture3("ruby", script.to_s, "--format", "json", *no_q4_args)
  assert(no_q4_status.success?, "missing q4 Sample Notepad evidence case must still emit packet: #{no_q4_stderr}")
  no_q4_packet = JSON.parse(no_q4_stdout)
  assert(no_q4_packet.fetch("q4_sample_notepad_smoke_status").fetch("status") == "not-supplied", "missing q4 Sample Notepad evidence must be visible")
  assert(no_q4_packet.fetch("release_blocking_reasons").include?("q4-sample-notepad-acceptance-smoke-not-passed"), "missing q4 Sample Notepad evidence must block release")
  assert(!no_q4_packet.fetch("merge_blocking_reasons").include?("q4-sample-notepad-acceptance-smoke-not-passed"), "missing q4 Sample Notepad evidence must not block merge")

  failed_q4_args = base_args.dup
  failed_q4_args[failed_q4_args.index("--q4-sample-notepad-smoke") + 1] = failed_q4_sample_notepad_smoke.path
  failed_q4_stdout, failed_q4_stderr, failed_q4_status = Open3.capture3("ruby", script.to_s, "--format", "json", *failed_q4_args)
  assert(failed_q4_status.success?, "failed q4 Sample Notepad evidence case must still emit packet: #{failed_q4_stderr}")
  failed_q4_packet = JSON.parse(failed_q4_stdout)
  assert(failed_q4_packet.fetch("q4_sample_notepad_smoke_status").fetch("status") == "blocked", "failed q4 Sample Notepad evidence must be visible")
  assert(failed_q4_packet.fetch("release_blocking_reasons").include?("q4-sample-notepad-acceptance-smoke-not-passed"), "failed q4 Sample Notepad evidence must add a release blocker")
  assert(!failed_q4_packet.fetch("merge_blocking_reasons").include?("q4-sample-notepad-acceptance-smoke-not-passed"), "failed q4 Sample Notepad evidence must not block merge")

  no_messagebox_args = base_args.each_slice(2).reject { |option, _path| option == "--q4-messagebox-smoke" }.flatten
  no_messagebox_stdout, no_messagebox_stderr, no_messagebox_status = Open3.capture3("ruby", script.to_s, "--format", "json", *no_messagebox_args)
  assert(no_messagebox_status.success?, "missing q4 MessageBox evidence case must still emit packet: #{no_messagebox_stderr}")
  no_messagebox_packet = JSON.parse(no_messagebox_stdout)
  assert(no_messagebox_packet.fetch("q4_messagebox_smoke_status").fetch("status") == "not-supplied", "missing q4 MessageBox evidence must be visible")
  assert(no_messagebox_packet.fetch("release_blocking_reasons").include?("q4-messagebox-document-content-smoke-not-passed"), "missing q4 MessageBox evidence must block release")
  assert(!no_messagebox_packet.fetch("merge_blocking_reasons").include?("q4-messagebox-document-content-smoke-not-passed"), "missing q4 MessageBox evidence must not block merge")

  failed_messagebox_args = base_args.dup
  failed_messagebox_args[failed_messagebox_args.index("--q4-messagebox-smoke") + 1] = failed_q4_messagebox_smoke.path
  failed_messagebox_stdout, failed_messagebox_stderr, failed_messagebox_status = Open3.capture3("ruby", script.to_s, "--format", "json", *failed_messagebox_args)
  assert(failed_messagebox_status.success?, "failed q4 MessageBox evidence case must still emit packet: #{failed_messagebox_stderr}")
  failed_messagebox_packet = JSON.parse(failed_messagebox_stdout)
  assert(failed_messagebox_packet.fetch("q4_messagebox_smoke_status").fetch("status") == "blocked", "failed q4 MessageBox evidence must be visible")
  assert(failed_messagebox_packet.fetch("release_blocking_reasons").include?("q4-messagebox-document-content-smoke-not-passed"), "failed q4 MessageBox evidence must add a release blocker")
  assert(!failed_messagebox_packet.fetch("merge_blocking_reasons").include?("q4-messagebox-document-content-smoke-not-passed"), "failed q4 MessageBox evidence must not block merge")

  no_known_existing_args = base_args.each_slice(2).reject { |option, _path| option == "--known-existing-winapp-acceptance" }.flatten
  no_known_existing_stdout, no_known_existing_stderr, no_known_existing_status = Open3.capture3("ruby", script.to_s, "--format", "json", *no_known_existing_args)
  assert(no_known_existing_status.success?, "missing known existing Windows app evidence case must still emit packet: #{no_known_existing_stderr}")
  no_known_existing_packet = JSON.parse(no_known_existing_stdout)
  assert(no_known_existing_packet.fetch("known_existing_winapp_acceptance_status").fetch("status") == "not-supplied", "missing known existing Windows app evidence must be visible")
  assert(no_known_existing_packet.fetch("release_blocking_reasons").include?("known-existing-winapp-acceptance-not-passed"), "missing known existing Windows app evidence must block release")
  assert(!no_known_existing_packet.fetch("merge_blocking_reasons").include?("known-existing-winapp-acceptance-not-passed"), "missing known existing Windows app evidence must not block merge")

  failed_known_existing_args = base_args.dup
  failed_known_existing_args[failed_known_existing_args.index("--known-existing-winapp-acceptance") + 1] = failed_known_existing_winapp_acceptance.path
  failed_known_existing_stdout, failed_known_existing_stderr, failed_known_existing_status = Open3.capture3("ruby", script.to_s, "--format", "json", *failed_known_existing_args)
  assert(failed_known_existing_status.success?, "failed known existing Windows app evidence case must still emit packet: #{failed_known_existing_stderr}")
  failed_known_existing_packet = JSON.parse(failed_known_existing_stdout)
  assert(failed_known_existing_packet.fetch("known_existing_winapp_acceptance_status").fetch("status") == "blocked", "failed known existing Windows app evidence must be visible")
  assert(failed_known_existing_packet.fetch("release_blocking_reasons").include?("known-existing-winapp-acceptance-not-passed"), "failed known existing Windows app evidence must add a release blocker")
  assert(!failed_known_existing_packet.fetch("merge_blocking_reasons").include?("known-existing-winapp-acceptance-not-passed"), "failed known existing Windows app evidence must not block merge")

  no_matrix_args = base_args.each_slice(2).reject { |option, _path| option == "--known-app-matrix-evidence" }.flatten
  no_matrix_stdout, no_matrix_stderr, no_matrix_status = Open3.capture3("ruby", script.to_s, "--format", "json", *no_matrix_args)
  assert(no_matrix_status.success?, "missing known app matrix evidence case must still emit packet: #{no_matrix_stderr}")
  no_matrix_packet = JSON.parse(no_matrix_stdout)
  assert(no_matrix_packet.fetch("known_app_matrix_evidence_status").fetch("status") == "not-supplied", "missing known app matrix evidence must be visible")
  assert(no_matrix_packet.fetch("release_blocking_reasons").include?("known-app-matrix-evidence-not-passed"), "missing known app matrix evidence must block release")
  assert(!no_matrix_packet.fetch("merge_blocking_reasons").include?("known-app-matrix-evidence-not-passed"), "missing known app matrix evidence must not block merge")

  failed_matrix_args = base_args.dup
  failed_matrix_args[failed_matrix_args.index("--known-app-matrix-evidence") + 1] = failed_known_app_matrix_evidence.path
  failed_matrix_stdout, failed_matrix_stderr, failed_matrix_status = Open3.capture3("ruby", script.to_s, "--format", "json", *failed_matrix_args)
  assert(failed_matrix_status.success?, "failed known app matrix evidence case must still emit packet: #{failed_matrix_stderr}")
  failed_matrix_packet = JSON.parse(failed_matrix_stdout)
  assert(failed_matrix_packet.fetch("known_app_matrix_evidence_status").fetch("status") == "blocked", "failed known app matrix evidence must be visible")
  assert(failed_matrix_packet.fetch("release_blocking_reasons").include?("known-app-matrix-evidence-not-passed"), "failed known app matrix evidence must add a release blocker")
  assert(!failed_matrix_packet.fetch("merge_blocking_reasons").include?("known-app-matrix-evidence-not-passed"), "failed known app matrix evidence must not block merge")

  failed_preflight_args = base_args.dup
  failed_preflight_args[failed_preflight_args.index("--desktop-trigger-request-preflight-smoke") + 1] = failed_preflight_smoke.path
  failed_preflight_stdout, failed_preflight_stderr, failed_preflight_status = Open3.capture3("ruby", script.to_s, "--format", "json", *failed_preflight_args)
  assert(failed_preflight_status.success?, "failed preflight smoke evidence case must still emit packet: #{failed_preflight_stderr}")
  failed_preflight_packet = JSON.parse(failed_preflight_stdout)
  assert(failed_preflight_packet.fetch("desktop_trigger_request_preflight_smoke_status").fetch("status") == "blocked", "failed preflight smoke evidence must be visible")
  assert(failed_preflight_packet.fetch("release_blocking_reasons").include?("desktop-trigger-request-preflight-smoke-not-passed"), "failed preflight smoke evidence must add a release-only blocker")
  assert(!failed_preflight_packet.fetch("merge_blocking_reasons").include?("desktop-trigger-request-preflight-smoke-not-passed"), "failed preflight smoke evidence must not block merge")

  malformed_preflight_args = base_args.dup
  malformed_preflight_args[malformed_preflight_args.index("--desktop-trigger-request-preflight-smoke") + 1] = malformed.path
  malformed_preflight_stdout, malformed_preflight_stderr, malformed_preflight_status = Open3.capture3("ruby", script.to_s, "--format", "json", *malformed_preflight_args)
  assert(malformed_preflight_status.success?, "malformed preflight smoke evidence case must still emit packet: #{malformed_preflight_stderr}")
  malformed_preflight_packet = JSON.parse(malformed_preflight_stdout)
  assert(malformed_preflight_packet.fetch("release_blocking_reasons").include?("desktop-trigger-request-preflight-smoke-not-passed"), "malformed preflight smoke evidence must add a release-only blocker")
  assert(malformed_preflight_packet.fetch("merge_blocking_reasons").empty?, "malformed optional preflight smoke evidence must not block merge")

  completed_args = base_args.dup
  completed_args[completed_args.index("--release-evidence-index") + 1] = completed_release_evidence.path
  completed_args[completed_args.index("--full-checkpoint-promotion") + 1] = completed_promotion.path
  completed_stdout, completed_stderr, completed_status = Open3.capture3("ruby", script.to_s, "--format", "json", *completed_args)
  assert(completed_status.success?, "completed product smoke case must emit a packet: #{completed_stderr}")
  completed_packet = JSON.parse(completed_stdout)
  assert(!completed_packet.fetch("release_blocking_reasons").include?("restricted-docker-or-qemu-smoke-requires-human-authorization"), "persisted authorized smoke evidence must close the QEMU authorization blocker")
  assert(!completed_packet.fetch("release_blocking_reasons").include?("full-checkpoint-promotion-not-allowed"), "passing promotion packet must close the promotion blocker")
  assert(completed_packet.fetch("release_blocking_reasons").include?("production-runtime-and-windows-execution-remain-disabled"), "product release must remain gated by production execution")
  assert(completed_packet.fetch("merge_ready"), "production execution gates must not block the completed stabilization train merge")

  skipped_stdout, skipped_stderr, skipped_status = Open3.capture3(
    "ruby", script.to_s, "--format", "json", "--skip-tool", "implementation", *base_args
  )
  assert(skipped_status.success?, "skipped tool case must still emit packet: #{skipped_stderr}")
  skipped_packet = JSON.parse(skipped_stdout)
  skipped_tools = skipped_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(skipped_tools.fetch("implementation").fetch("status") == "skipped", "packet must record skipped tools")
  assert(skipped_packet.fetch("merge_blocking_reasons").include?("implementation:skipped"), "required skipped tool must block merge")

  missing_stdout, missing_stderr, missing_status = Open3.capture3(
    "ruby",
    script.to_s,
    "--format",
    "json",
    "--tool-command",
    "implementation=definitely_missing_xnix_packet_tool",
    *base_args.reject.with_index { |_value, index| [2, 3].include?(index) }
  )
  assert(missing_status.success?, "missing command case must still emit packet: #{missing_stderr}")
  missing_packet = JSON.parse(missing_stdout)
  missing_tools = missing_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(missing_tools.fetch("implementation").fetch("status") == "missing-command", "packet must detect missing commands")
  assert(missing_packet.fetch("merge_blocking_reasons").include?("implementation:missing-command"), "missing required command must block merge")

  malformed_args = base_args.dup
  malformed_args[malformed_args.index("--implementation-report") + 1] = malformed.path
  malformed_stdout, malformed_stderr, malformed_status = Open3.capture3("ruby", script.to_s, "--format", "json", *malformed_args)
  assert(malformed_status.success?, "malformed JSON case must still emit packet: #{malformed_stderr}")
  malformed_packet = JSON.parse(malformed_stdout)
  malformed_tools = malformed_packet.fetch("tool_statuses").to_h { |tool| [tool.fetch("id"), tool] }
  assert(malformed_tools.fetch("implementation").fetch("status") == "malformed-json", "packet must detect malformed JSON")
  assert(malformed_packet.fetch("merge_blocking_reasons").include?("implementation:malformed-json"), "malformed required JSON must block merge")

  protected_args = base_args.dup
  protected_args[protected_args.index("--mainline-review") + 1] = protected_mainline.path
  protected_stdout, protected_stderr, protected_status = Open3.capture3("ruby", script.to_s, "--format", "json", *protected_args)
  assert(protected_status.success?, "protected file case must still emit packet: #{protected_stderr}")
  protected_packet = JSON.parse(protected_stdout)
  assert(protected_packet.fetch("protected_file_status").fetch("status") == "blocked", "protected file change must block")
  assert(protected_packet.fetch("merge_blocking_reasons").include?("protected-claude-file-modified"), "protected file blocker must be explicit")
  assert(protected_packet.fetch("merge_blocking_reasons").include?("unclassified-files-present"), "unclassified file blocker must be explicit")

  unsafe_args = base_args.dup
  unsafe_args[unsafe_args.index("--kde-smoke-report") + 1] = unsafe_kde.path
  unsafe_stdout, unsafe_stderr, unsafe_status_result = Open3.capture3("ruby", script.to_s, "--format", "json", *unsafe_args)
  assert(unsafe_status_result.success?, "unsafe flag case must still emit packet: #{unsafe_stderr}")
  unsafe_packet = JSON.parse(unsafe_stdout)
  findings = unsafe_packet.fetch("unsafe_operation_status").fetch("detected_unsafe_findings")
  assert(findings.any? { |finding| finding.fetch("key") == "docker_executed" }, "packet must report unsafe-operation findings")
  assert(unsafe_packet.fetch("merge_blocking_reasons").include?("unsafe-operation-detected"), "unsafe operation must block merge")
ensure
  [
    layout,
    implementation,
    contract_drift,
    kde_smoke,
    mainline,
    release_evidence,
    completed_release_evidence,
    blocked_promotion,
    completed_promotion,
    fixture_matrix,
    preflight_smoke,
    failed_preflight_smoke,
    q4_sample_notepad_smoke,
    q4_messagebox_smoke,
    known_existing_winapp_acceptance,
    known_app_matrix_evidence,
    failed_q4_sample_notepad_smoke,
    failed_q4_messagebox_smoke,
    failed_known_existing_winapp_acceptance,
    failed_known_app_matrix_evidence,
    malformed,
    protected_mainline,
    unsafe_kde
  ].compact.each do |file|
    file.close
    file.unlink
  end
end
