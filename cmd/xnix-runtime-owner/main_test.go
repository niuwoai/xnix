package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func TestRuntimeOwnerCommandRendersSmokeCandidate(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_candidate.v1" ||
		payload["request_type"] != "runtime-owner-candidate" ||
		payload["owner_type"] != "go-runtime-owner-candidate" ||
		payload["mode"] != "smoke-owner" {
		t.Fatalf("unexpected owner candidate schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["read_only_serve_ready"] != true ||
		payload["route_count"] != float64(61) ||
		payload["go_route_count"] != float64(61) ||
		payload["c_core_route_count"] != float64(0) ||
		payload["ruby_legacy_route_count"] != float64(0) ||
		payload["write_method_count"] != float64(4) {
		t.Fatalf("unexpected owner candidate readiness: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["smoke_owner_mode"] != true ||
		payload["production_owner_mode"] != false ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected owner candidate safety flags: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetBackendAdapterProfileAudit",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetBackendAdapterProfileAudit" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected redacted adapter profile service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "backend-adapter-redacted-profile-audit-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" ||
		nested["request_type"] != "backend-adapter-redacted-profile-audit-preview" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["full_contract_fixture_local"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected redacted adapter profile nested payload: %#v", nested)
	}
	if strings.Contains(strings.ToLower(output.String()), "wine") ||
		strings.Contains(strings.ToLower(output.String()), "proton") ||
		strings.Contains(strings.ToLower(output.String()), "windows-vm") {
		t.Fatalf("redacted adapter profile owner command exposed internal adapter terms: %s", output.String())
	}
}

func TestRuntimeOwnerCommandRendersShowRuntimeControlledLaunchServiceCall(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeOwnerCommandRuntimeStatusLaunchExecutionFixture(t, stateRoot)
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	evidenceRecord := recordOwnerCommandRuntimeStatusLaunchEvidenceFixture(t, stateRoot)
	fakeLauncher, fakeArgsPath := writeOwnerCommandFakeRuntimeStatusManagedLauncher(t)
	t.Setenv("XNIX_RUNTIME_OWNER_STATE_ROOT", stateRoot)
	t.Setenv("XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT", t.TempDir())
	t.Setenv("XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER", fakeLauncher)
	t.Setenv("XNIX_RUNTIME_OWNER_TIMEOUT", "5s")
	t.Setenv("XNIX_RUNTIME_OWNER_GUEST_TIMEOUT", "1s")
	t.Setenv("XNIX_RUNTIME_OWNER_GUI_EXECUTABLE", "/tmp/xnix-owner-controlled-input/owner-messagebox.exe")

	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--mode", "smoke-owner",
		"--service-call", "ShowRuntimeControlledLaunch",
		"evidence-relative-path", evidenceRecord.EvidenceRelativePath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "ShowRuntimeControlledLaunch" ||
		payload["call_type"] != "desktop-action-dispatch" ||
		payload["read_only_dispatch"] != false ||
		payload["write_method"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["dispatch_ready"] != true ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime controlled launch owner command response: %#v", payload)
	}
	action := payload["payload"].(map[string]any)
	if action["owner_request_type"] != "runtime-owner-show-runtime-controlled-launch" ||
		action["owner_runtime_method"] != "ShowRuntimeControlledLaunch" ||
		action["desktop_callable_action_id"] != appidentity.KnownAppKDERuntimeStatusLaunchAction ||
		action["desktop_callable_route"] != "kde-dbus-runtime-status-action" ||
		action["desktop_evidence_handle_forwarded"] != true ||
		action["desktop_receipt_fields_reconstructed"] != false ||
		action["desktop_kde_state_root_access"] != false ||
		action["runtime_owner_service_action_dispatch"] != true ||
		action["runtime_owner_service_supplies_owner_inputs"] != true ||
		action["kde_forwards_only_evidence_handle"] != true ||
		action["evidence_relative_path"] != evidenceRecord.EvidenceRelativePath ||
		action["launch_authorization_receipt_id"] != launchReceiptID ||
		action["session_gated_review_receipt_id"] != reviewReceiptID ||
		action["controlled_execution_session_id"] != sessionID ||
		action["managed_launcher_invoked"] != true ||
		action["delegated_status"] != "passed" ||
		action["delegated_session_gated_review_receipt_id"] != reviewReceiptID ||
		action["delegated_controlled_execution_session_id"] != sessionID ||
		action["delegated_host_root_modified"] != false ||
		action["delegated_docker_socket_mounted"] != false ||
		action["delegated_broad_host_mount_required"] != false ||
		action["delegated_raw_command_exposed"] != false ||
		action["delegated_backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime controlled launch owner command action: %#v", action)
	}
	projection := action["compatibility_center_known_app_evidence"].(map[string]any)
	if projection["status"] != "passed" ||
		projection["controlled_execution_session_id"] != sessionID ||
		projection["compatibility_center_projection_ready"] != true ||
		projection["kde_center_projection_ready"] != true ||
		projection["state_root_path_exposed"] != false ||
		projection["managed_launcher_path_exposed"] != false ||
		projection["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime controlled launch projection: %#v", projection)
	}
	argsData, err := os.ReadFile(fakeArgsPath)
	if err != nil {
		t.Fatalf("ReadFile fake launcher args returned error: %v", err)
	}
	argsText := string(argsData)
	for _, token := range []string{"--state-root\n" + stateRoot, "--receipt-id\n" + launchReceiptID, "--review-receipt-id\n" + reviewReceiptID, "--session-id\n" + sessionID, "--timeout\n1s", "--executable\n/tmp/xnix-owner-controlled-input/owner-messagebox.exe"} {
		if !strings.Contains(argsText, token) {
			t.Fatalf("fake launcher args missing %q: %s", token, argsText)
		}
	}
	text := strings.ToLower(output.String())
	for _, forbidden := range []string{strings.ToLower(stateRoot), strings.ToLower(fakeLauncher), ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Runtime controlled launch owner command exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerSmokeCoverage(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--redacted-adapter-profile-owner-smoke-coverage",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "backend-adapter-redacted-profile-owner-smoke-coverage-preview" ||
		payload["coverage_type"] != "redacted-adapter-profile-owner-smoke-coverage" ||
		payload["runtime_method"] != "GetBackendAdapterProfileAudit" ||
		payload["owner_route_request_type"] != "backend-adapter-redacted-profile-audit-preview" ||
		payload["smoke_batch_record_count"] != float64(75) ||
		payload["smoke_read_dispatch_record_count"] != float64(71) ||
		payload["smoke_write_denial_record_count"] != float64(4) ||
		payload["coverage_record_found"] != true ||
		payload["service_call_dispatch_ready"] != true ||
		payload["dispatch_route_source"] != "go-owner-local-preview" ||
		payload["dispatch_go_command"] != "backend-adapter-redacted-profile-audit-preview" ||
		payload["route_decision"] != "redacted-profile-route-ready" ||
		payload["profile_count"] != float64(3) ||
		payload["internal_adapter_ids_redacted"] != true ||
		payload["internal_profile_paths_redacted"] != true ||
		payload["full_contract_fixture_local"] != true ||
		payload["smoke_coverage_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false ||
		payload["adapter_invocation_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected redacted adapter profile owner smoke coverage payload: %#v", payload)
	}
	if strings.Contains(strings.ToLower(output.String()), "wine") ||
		strings.Contains(strings.ToLower(output.String()), "proton") ||
		strings.Contains(strings.ToLower(output.String()), "windows-vm") {
		t.Fatalf("redacted adapter profile smoke coverage exposed internal adapter terms: %s", output.String())
	}
}

func TestRuntimeOwnerCommandRendersRestrictedOwnerSmokeReceiptLookupReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetRestrictedOwnerSmokeReceiptLookupPreview", "restricted-owner-smoke-receipt-id",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetRestrictedOwnerSmokeReceiptLookupPreview" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected restricted owner smoke lookup service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "restricted-owner-smoke-receipt-lookup-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" ||
		nested["request_type"] != "restricted-owner-smoke-receipt-lookup-preview" ||
		nested["opaque_receipt_id"] != "restricted-owner-smoke-receipt-id" ||
		nested["owner_managed_lookup"] != true ||
		nested["caller_state_root_required"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected restricted owner smoke lookup nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersMaterializationReceiptLookupReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetKDETestLaunchMaterializationReceiptLookupPreview", "kde-test-launch-materialization-receipt-id",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetKDETestLaunchMaterializationReceiptLookupPreview" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected materialization receipt lookup service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "kde-test-launch-materialization-receipt-lookup-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" ||
		nested["request_type"] != "kde-test-launch-materialization-receipt-lookup-preview" ||
		nested["opaque_materialization_receipt_id"] != "kde-test-launch-materialization-receipt-id" ||
		nested["owner_managed_opaque_receipt_lookup_ready"] != true ||
		nested["requires_caller_state_root"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected materialization receipt lookup nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersMaterializationFanOutReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetKDETestLaunchMaterializationFanOut", "kde-test-launch-materialization-receipt-id",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetKDETestLaunchMaterializationFanOut" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected materialization fan-out owner route service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" ||
		nested["request_type"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		nested["opaque_materialization_receipt_id"] != "kde-test-launch-materialization-receipt-id" ||
		nested["owner_managed_opaque_receipt_lookup_ready"] != true ||
		nested["requires_caller_state_root"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected materialization fan-out owner route nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersMaterializationFanOutOwnerSmokeCoverage(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--materialization-fanout-owner-smoke-coverage",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "kde-test-launch-materialization-fanout-owner-smoke-coverage-preview" ||
		payload["coverage_type"] != "materialization-fanout-owner-smoke-coverage" ||
		payload["runtime_method"] != "GetKDETestLaunchMaterializationFanOut" ||
		payload["owner_route_request_type"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		payload["smoke_batch_record_count"] != float64(75) ||
		payload["smoke_read_dispatch_record_count"] != float64(71) ||
		payload["smoke_write_denial_record_count"] != float64(4) ||
		payload["coverage_record_found"] != true ||
		payload["service_call_dispatch_ready"] != true ||
		payload["dispatch_route_source"] != "go-owner-local-preview" ||
		payload["dispatch_go_command"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		payload["opaque_materialization_receipt_id"] != "kde-test-launch-materialization-receipt-id" ||
		payload["receipt_lookup_state"] != "missing-receipt" ||
		payload["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		payload["smoke_coverage_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected materialization fan-out owner smoke coverage payload: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersRestrictedOwnerSmokeReceiptFanOutReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetRestrictedOwnerSmokeReceiptFanOut", "restricted-owner-smoke-receipt-id",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected restricted owner smoke fan-out service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" ||
		nested["request_type"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		nested["opaque_receipt_id"] != "restricted-owner-smoke-receipt-id" ||
		nested["owner_managed_lookup"] != true ||
		nested["caller_state_root_required"] != false ||
		nested["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected restricted owner smoke fan-out nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersRestrictedSmokeFanOutOwnerSmokeCoverage(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--restricted-smoke-fanout-owner-smoke-coverage",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview" ||
		payload["coverage_type"] != "restricted-owner-smoke-fanout-owner-smoke-coverage" ||
		payload["runtime_method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		payload["owner_route_request_type"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		payload["smoke_batch_record_count"] != float64(75) ||
		payload["smoke_read_dispatch_record_count"] != float64(71) ||
		payload["smoke_write_denial_record_count"] != float64(4) ||
		payload["coverage_record_found"] != true ||
		payload["service_call_dispatch_ready"] != true ||
		payload["dispatch_route_source"] != "go-owner-local-preview" ||
		payload["dispatch_go_command"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		payload["opaque_receipt_id"] != "restricted-owner-smoke-receipt-id" ||
		payload["receipt_lookup_state"] != "missing-receipt" ||
		payload["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		payload["smoke_coverage_ready"] != true ||
		payload["production_dbus_exposure_ready"] != false ||
		payload["support_bundle_exported"] != false ||
		payload["support_case_created"] != false ||
		payload["notification_sent"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected restricted smoke fan-out owner smoke coverage payload: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersDisabledWrite(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--deny-write", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "Launch" ||
		payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		payload["dispatch_enabled"] != false ||
		payload["request_created"] != false {
		t.Fatalf("unexpected disabled write response: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--dispatch-read", "GetRuntimeWriteGate", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		payload["request_type"] != "runtime-owner-read-dispatch" ||
		payload["dispatch_type"] != "go-owner-read-dispatch" ||
		payload["method"] != "GetRuntimeWriteGate" ||
		payload["go_command"] != "runtime-write-gate-preview" ||
		payload["read_only_dispatch"] != true ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "runtime-write-gate-preview" ||
		nested["method_name"] != "Launch" ||
		nested["dispatch_enabled"] != false {
		t.Fatalf("unexpected read dispatch payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersWindowsCompatibilityOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--dispatch-read", "GetWindowsCompatibilityWorkstreamsPreview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		payload["method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		payload["route_source"] != "go-owner-local-preview" ||
		payload["go_command"] != "windows-compatibility-workstreams-preview" ||
		payload["route_status"] != "owner-local-preview-ready" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Windows compatibility read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "windows-compatibility-workstreams-preview" ||
		nested["read_method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		nested["official_desktop"] != "KDE Plasma" ||
		nested["runtime_owned"] != true ||
		nested["go_runtime_backed"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected Windows compatibility nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersNotificationDigestOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--dispatch-read", "GetKDENotificationDigestPreview",
		"org.xnix.sample.notepad",
		"permission-attention:permission-review",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetKDENotificationDigestPreview" ||
		payload["route_source"] != "go-owner-local-preview" ||
		payload["go_command"] != "kde-notification-digest-preview" ||
		payload["route_status"] != "owner-local-preview-ready" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected notification digest read dispatch response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["request_type"] != "kde-notification-digest-preview" ||
		nested["read_method"] != "GetKDENotificationDigestPreview" ||
		nested["notifications_sent"] != false ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected notification digest nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersOfflineKDEIdentityOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetKDEOfflineApplicationIdentityPreview",
		"org.xnix.sample.notepad",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetKDEOfflineApplicationIdentityPreview" ||
		payload["call_type"] != "read-dispatch" || payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false || payload["production_bus_claimed"] != false ||
		payload["network_required"] != false || payload["host_root_modified"] != false {
		t.Fatalf("unexpected offline KDE identity service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "kde-offline-application-identity-preview" ||
		nested["surface_count"] != float64(9) ||
		nested["cross_surface_identity_consistent"] != true ||
		nested["settings_persisted"] != false ||
		nested["compatibility_center_persisted"] != false ||
		nested["launch_enabled"] != false || nested["host_root_modified"] != false {
		t.Fatalf("unexpected offline KDE identity nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersSignedRecipeOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetSignedRecipeVerificationPreview",
		"org.xnix.sample.notepad",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetSignedRecipeVerificationPreview" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected signed recipe service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "signed-recipe-verifier-preview" ||
		nested["verification_state"] != "production-signature-required" ||
		nested["recipe_digest_verified"] != true ||
		nested["signature_verified"] != false ||
		nested["signature_material_exposed"] != false {
		t.Fatalf("unexpected signed recipe nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersRestrictedSmokeOwnerLocalReadDispatch(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{
		"--root", "../..",
		"--service-call", "GetRestrictedProductSmokePacketPreview",
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["method"] != "GetRestrictedProductSmokePacketPreview" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected restricted smoke service response: %#v", payload)
	}
	dispatch := payload["payload"].(map[string]any)
	nested := dispatch["payload"].(map[string]any)
	if dispatch["go_command"] != "restricted-product-smoke-packet-preview" ||
		nested["packet_prepared"] != true ||
		nested["human_authorization_required"] != true ||
		nested["execution_authorized"] != false ||
		nested["docker_executed"] != false ||
		nested["qemu_executed"] != false ||
		nested["release_ready"] != false {
		t.Fatalf("unexpected restricted smoke nested payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersServiceCall(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner", "--service-call", "GetRuntimeWriteGate", "Launch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_service_call.v1" ||
		payload["request_type"] != "runtime-owner-service-call" ||
		payload["service_type"] != "go-runtime-owner-in-process-service" ||
		payload["method"] != "GetRuntimeWriteGate" ||
		payload["call_type"] != "read-dispatch" ||
		payload["read_only_dispatch"] != true ||
		payload["write_method"] != false ||
		payload["dispatch_ready"] != true ||
		payload["in_process_service_ready"] != true ||
		payload["event_loop_started"] != false ||
		payload["session_bus_claimed"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected service call response: %#v", payload)
	}
	nested := payload["payload"].(map[string]any)
	if nested["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		nested["method"] != "GetRuntimeWriteGate" {
		t.Fatalf("unexpected nested service call payload: %#v", nested)
	}
}

func TestRuntimeOwnerCommandRendersLifecycleJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--mode", "smoke-owner", "--lifecycle-log"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 4 {
		t.Fatalf("lifecycle line count = %d, want 4: %q", len(lines), output.String())
	}
	wantTypes := []string{"startup", "route-table", "readiness", "shutdown"}
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_lifecycle_event.v1" ||
			payload["request_type"] != "runtime-owner-lifecycle-event" ||
			payload["event_type"] != wantTypes[index] ||
			payload["sequence"] != float64(index+1) ||
			payload["mode"] != "smoke-owner" ||
			payload["route_count"] != float64(61) ||
			payload["go_route_count"] != float64(61) ||
			payload["write_method_count"] != float64(4) {
			t.Fatalf("unexpected lifecycle event %d: %#v", index, payload)
		}
		if payload["event_loop_started"] != false ||
			payload["session_bus_claimed"] != false ||
			payload["production_bus_claimed"] != false ||
			payload["write_methods_enabled"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected lifecycle safety flags at %d: %#v", index, payload)
		}
	}
}

func TestRuntimeOwnerCommandRendersSmokeBatchJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--smoke-batch"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 75 {
		t.Fatalf("smoke batch line count = %d, want 75", len(lines))
	}
	readCount := 0
	writeCount := 0
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_smoke_batch.v1" ||
			payload["request_type"] != "runtime-owner-smoke-batch-record" ||
			payload["batch_type"] != "restricted-session-owner-call-batch" ||
			payload["sequence"] != float64(index+1) ||
			payload["read_dispatch_method_count"] != float64(71) ||
			payload["write_method_count"] != float64(4) ||
			payload["runtime_owned"] != true ||
			payload["go_runtime_backed"] != true ||
			payload["event_loop_started"] != false ||
			payload["session_bus_claimed"] != false ||
			payload["production_bus_claimed"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected smoke batch record %d: %#v", index, payload)
		}
		switch payload["record_type"] {
		case "read-dispatch":
			readCount++
			if payload["read_only_dispatch"] != true ||
				payload["write_method"] != false ||
				payload["dispatch_ready"] != true {
				t.Fatalf("unexpected read smoke batch record %d: %#v", index, payload)
			}
			nested := payload["payload"].(map[string]any)
			if nested["request_type"] != "runtime-owner-service-call" ||
				nested["call_type"] != "read-dispatch" ||
				nested["read_only_dispatch"] != true {
				t.Fatalf("unexpected read smoke batch payload %d: %#v", index, nested)
			}
		case "write-denial":
			writeCount++
			if payload["read_only_dispatch"] != false ||
				payload["write_method"] != true ||
				payload["dispatch_ready"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				t.Fatalf("unexpected write smoke batch record %d: %#v", index, payload)
			}
			nested := payload["payload"].(map[string]any)
			if nested["request_type"] != "runtime-owner-service-call" ||
				nested["call_type"] != "write-denial" ||
				nested["write_method"] != true {
				t.Fatalf("unexpected write smoke batch payload %d: %#v", index, nested)
			}
		default:
			t.Fatalf("unexpected smoke batch record type at %d: %#v", index, payload)
		}
	}
	if readCount != 71 || writeCount != 4 {
		t.Fatalf("smoke batch counts read=%d write=%d, want 71/4", readCount, writeCount)
	}
}

func TestRuntimeOwnerCommandRendersSessionBusSmokeJSONL(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--session-bus-smoke"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(output.String()), "\n")
	if len(lines) != 80 {
		t.Fatalf("session bus smoke line count = %d, want 80", len(lines))
	}
	readCount := 0
	writeCount := 0
	unsupportedCount := 0
	for index, line := range lines {
		var payload map[string]any
		if err := json.Unmarshal([]byte(line), &payload); err != nil {
			t.Fatalf("Unmarshal line %d returned error: %v", index, err)
		}
		if payload["version"] != currentProjectVersion(t) ||
			payload["schema_version"] != "xnix.runtime.owner_session_bus_smoke.v1" ||
			payload["request_type"] != "runtime-owner-session-bus-smoke-step" ||
			payload["transcript_type"] != "restricted-private-session-bus-owner-smoke" ||
			payload["sequence"] != float64(index+1) ||
			payload["read_dispatch_method_count"] != float64(71) ||
			payload["write_method_count"] != float64(4) ||
			payload["runtime_owned"] != true ||
			payload["go_runtime_backed"] != true ||
			payload["kde_policy_owner"] != false ||
			payload["private_session_bus"] != true ||
			payload["event_loop_started"] != true ||
			payload["session_bus_claimed"] != true ||
			payload["production_bus_claimed"] != false ||
			payload["system_service_started"] != false ||
			payload["write_methods_enabled"] != false ||
			payload["host_root_modified"] != false ||
			payload["backend_details_exposed"] != false {
			t.Fatalf("unexpected session bus smoke payload at %d: %#v", index, payload)
		}
		switch payload["step_type"] {
		case "read-dispatch":
			readCount++
			if payload["read_only_dispatch"] != true || payload["write_method"] != false {
				t.Fatalf("unexpected read step at %d: %#v", index, payload)
			}
		case "write-denial":
			writeCount++
			if payload["read_only_dispatch"] != false ||
				payload["write_method"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				t.Fatalf("unexpected write step at %d: %#v", index, payload)
			}
		case "reject-unsupported-read":
			unsupportedCount++
			if payload["unsupported_read"] != true ||
				payload["error_name"] != "org.xnix.Compatibility1.Error.UnsupportedMethod" {
				t.Fatalf("unexpected unsupported-read step at %d: %#v", index, payload)
			}
		}
	}
	if readCount != 71 || writeCount != 4 || unsupportedCount != 1 {
		t.Fatalf("session bus smoke counts read=%d write=%d unsupported=%d, want 71/4/1", readCount, writeCount, unsupportedCount)
	}
}

func TestRuntimeOwnerCommandRejectsMixedOperations(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--deny-write", "Launch", "--lifecycle-log"}, &output)
	if err == nil || !strings.Contains(err.Error(), "accepts only one") {
		t.Fatalf("run must reject mixed lifecycle/write operations, got %v", err)
	}
}

func TestRuntimeOwnerCommandRendersRouteCheckpoint(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--route-checkpoint"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.owner_route_checkpoint.v1" ||
		payload["formal_read_route_count"] != float64(61) ||
		payload["go_formal_read_route_count"] != float64(61) ||
		payload["owner_read_method_count"] != float64(71) ||
		payload["owner_local_read_method_count"] != float64(10) ||
		payload["smoke_read_record_count"] != float64(71) ||
		payload["smoke_write_denial_count"] != float64(4) ||
		payload["route_band_ready"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected route checkpoint payload: %#v", payload)
	}
}

func TestRuntimeOwnerCommandRendersKDEOfflineIdentityCheckpoint(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"--root", "../..", "--kde-identity-checkpoint"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.kde_offline_identity_checkpoint.v1" ||
		payload["application_id"] != "org.xnix.sample.notepad" ||
		payload["surface_count"] != float64(9) || payload["expected_surface_count"] != float64(9) ||
		payload["owner_read_method_count"] != float64(71) || payload["owner_local_read_method_count"] != float64(10) ||
		payload["offline_identity_ready"] != true || payload["production_signature_ready"] != false ||
		payload["production_bus_claimed"] != false || payload["write_methods_enabled"] != false ||
		payload["launch_enabled"] != false || payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected offline KDE identity checkpoint payload: %#v", payload)
	}
}

func writeOwnerCommandRuntimeStatusLaunchExecutionFixture(t *testing.T, stateRoot string) string {
	t.Helper()
	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")
	ledger, err := execution.NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	if _, err := ledger.Record(execution.Transaction{
		RequestID:      sessionID,
		ApplicationID:  "7zr",
		Profile:        "known-app-managed-guest",
		State:          execution.StateBlocked,
		ReviewDecision: execution.DecisionApproved,
		Gates: []execution.Gate{
			{ID: "controlled-execution-session-handoff", Status: execution.GatePass, Reason: "Runtime execution session handoff ready"},
			{ID: "runtime-write-gate", Status: execution.GateBlocked, Reason: "dispatch runner must consume the recorded session before execution"},
		},
		BlockedReasons:   []string{"runtime-write-gate: dispatch runner must consume the recorded session before execution"},
		LaunchAllowed:    false,
		LaunchEnabled:    false,
		BackendStarted:   false,
		HostRootModified: false,
		NetworkRequired:  false,
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
	if _, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: appidentity.KnownAppLaunchAuthorizationReceiptAction,
	}); err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	if _, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  appidentity.KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	}); err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	return sessionID
}

func recordOwnerCommandRuntimeStatusLaunchEvidenceFixture(t *testing.T, stateRoot string) appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()
	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := appidentity.ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip Console",
		AppVersion:                             "26.02",
		RequestType:                            "windows-known-app-dispatch-smoke",
		Status:                                 "passed",
		GuestBoundary:                          "managed-known-app-guest-smoke",
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       true,
		MarkerObserved:                         true,
		SessionGatedControlledDispatchConsumed: true,
		SessionGatedControlledDispatchState:    "created-after-session-gated-review",
		SessionGatedReviewReceiptID:            reviewReceiptID,
		LaunchAuthorizationReceiptID:           launchReceiptID,
		LaunchAuthorizationReceiptState:        "recorded",
		LaunchGateState:                        "controlled-dispatch-ready",
		LaunchGateConsumed:                     true,
		LaunchGateReceiptAccepted:              true,
		LaunchGateGuestBoundaryAccepted:        true,
		ControlledDispatchReady:                true,
		ControlledExecutionSessionConsumed:     true,
		ControlledExecutionSessionID:           sessionID,
		ControlledSessionDigestVerified:        true,
		ControlledSessionRelativePath:          "execution-ledger/sessions/" + sessionID + ".json",
		RuntimeOwnerConsumableSession:          true,
		KDEReadModelConsumableSession:          true,
	})
	if err != nil {
		t.Fatalf("ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence returned error: %v", err)
	}
	record, err := appidentity.RecordKnownAppKDERuntimeStatusLaunchEvidence(appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	return record
}

func writeOwnerCommandFakeRuntimeStatusManagedLauncher(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()
	launcherPath := filepath.Join(dir, "fake-xnix-compat-launch")
	argsPath := filepath.Join(dir, "launcher-args.txt")
	t.Setenv("XNIX_OWNER_COMMAND_FAKE_LAUNCHER_ARGS_FILE", argsPath)
	script := "#!/bin/sh\nprintf '%s\n' \"$@\" > \"$XNIX_OWNER_COMMAND_FAKE_LAUNCHER_ARGS_FILE\"\nprintf '%s\n' '{\"request_type\":\"windows-known-app-dispatch-smoke\",\"status\":\"passed\",\"guest_boundary\":\"managed-known-app-guest-smoke\",\"runtime_owned_dispatch\":true,\"artifact_verified\":true,\"marker_observed\":true,\"smoke_passed\":true,\"execution_started\":true,\"backend_process_started\":false,\"session_gated_controlled_dispatch_consumed\":true,\"session_gated_controlled_dispatch_state\":\"created-after-session-gated-review\",\"session_gated_review_receipt_id\":\"known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02\",\"launch_authorization_receipt_id\":\"known-app-launch-authorization-7zr-26.02\",\"controlled_execution_session_consumed\":true,\"controlled_execution_session_id\":\"known-app-controlled-execution-session-7zr-26.02\",\"controlled_session_digest_verified\":true,\"controlled_session_relative_path\":\"execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json\",\"runtime_owner_consumable_session\":true,\"kde_read_model_consumable_session\":true,\"controlled_session_live_state_observed\":false,\"controlled_session_registered\":false,\"controlled_session_window_observed\":false,\"controlled_session_host_root_modified\":false,\"controlled_session_backend_process_start\":false,\"host_root_modified\":false,\"docker_socket_mounted\":false,\"broad_host_mount_required\":false,\"raw_command_exposed\":false,\"backend_details_exposed\":false}'\n"
	if err := os.WriteFile(launcherPath, []byte(script), 0o755); err != nil {
		t.Fatalf("WriteFile fake launcher returned error: %v", err)
	}
	return launcherPath, argsPath
}
