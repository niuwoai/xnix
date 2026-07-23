package owner

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func TestServiceCallServesReadDispatchInProcess(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("GetRuntimeWriteGate", []string{"Launch"})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}

	if call.Version != currentProjectVersion(t) ||
		call.SchemaVersion != "xnix.runtime.owner_service_call.v1" ||
		call.RequestType != "runtime-owner-service-call" ||
		call.ServiceType != "go-runtime-owner-in-process-service" ||
		call.Method != "GetRuntimeWriteGate" ||
		call.CallType != "read-dispatch" ||
		call.BusName != "org.xnix.Compatibility1" ||
		call.RouteCount != 61 ||
		call.GoRouteCount != 61 ||
		call.WriteMethodCount != 4 ||
		!call.ReadOnlyServeReady ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		!call.InProcessServiceReady {
		t.Fatalf("unexpected read service call: %#v", call)
	}
	if call.RuntimeOwned != true ||
		call.GoRuntimeBacked != true ||
		call.KDEPolicyOwner != false ||
		call.KDEMayClaimRuntimeOwnership != false ||
		call.EventLoopStarted != false ||
		call.SessionBusClaimed != false ||
		call.ProductionBusClaimed != false ||
		call.SystemServiceStarted != false ||
		call.NetworkRequired != false ||
		call.HostRootModified != false ||
		call.PrivilegedContainerRequired != false ||
		call.BackendDetailsExposed != false {
		t.Fatalf("unexpected read service safety flags: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		dispatch["method"] != "GetRuntimeWriteGate" ||
		dispatch["read_only_dispatch"] != true {
		t.Fatalf("unexpected nested dispatch payload: %#v", dispatch)
	}
}

func TestServiceCallServesOwnerLocalOfflineKDEIdentity(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetKDEOfflineApplicationIdentityPreview", []string{"org.xnix.sample.notepad"})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetKDEOfflineApplicationIdentityPreview" || call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch || call.WriteMethod || call.WriteMethodsEnabled || !call.DispatchReady ||
		call.ProductionBusClaimed || call.NetworkRequired || call.HostRootModified {
		t.Fatalf("unexpected offline KDE identity service call: %#v", call)
	}
	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetKDEOfflineApplicationIdentityPreview" ||
		dispatch["go_command"] != "kde-offline-application-identity-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested offline KDE identity dispatch: %#v", dispatch)
	}
}

func TestServiceCallServesOwnerLocalRedactedBackendAdapterProfileAudit(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetBackendAdapterProfileAudit", nil)
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetBackendAdapterProfileAudit" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected redacted adapter profile service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetBackendAdapterProfileAudit" ||
		dispatch["go_command"] != "backend-adapter-redacted-profile-audit-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested redacted adapter profile dispatch: %#v", dispatch)
	}
	nested := dispatch["payload"].(map[string]any)
	if nested["request_type"] != "backend-adapter-redacted-profile-audit-preview" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected nested redacted adapter profile payload: %#v", nested)
	}
}

func TestServiceCallDispatchesShowRuntimeControlledLaunch(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeOwnerRuntimeStatusLaunchExecutionFixture(t, stateRoot)
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	evidenceRecord := recordOwnerRuntimeStatusLaunchEvidenceFixture(t, stateRoot)
	fakeLauncher, fakeArgsPath := writeOwnerFakeRuntimeStatusManagedLauncher(t)
	t.Setenv("XNIX_RUNTIME_OWNER_STATE_ROOT", stateRoot)
	t.Setenv("XNIX_RUNTIME_OWNER_KNOWN_APP_CACHE_ROOT", t.TempDir())
	t.Setenv("XNIX_RUNTIME_OWNER_MANAGED_LAUNCHER", fakeLauncher)
	t.Setenv("XNIX_RUNTIME_OWNER_TIMEOUT", "5s")
	t.Setenv("XNIX_RUNTIME_OWNER_GUEST_TIMEOUT", "1s")

	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("ShowRuntimeControlledLaunch", []string{"--evidence-relative-path", evidenceRecord.EvidenceRelativePath})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}

	if call.Method != "ShowRuntimeControlledLaunch" ||
		call.CallType != "desktop-action-dispatch" ||
		call.ReadOnlyDispatch ||
		!call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected Runtime controlled launch service call: %#v", call)
	}
	var payload map[string]any
	if err := json.Unmarshal(call.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["owner_schema_version"] != "xnix.runtime.owner_show_runtime_controlled_launch.v1" ||
		payload["owner_request_type"] != "runtime-owner-show-runtime-controlled-launch" ||
		payload["owner_runtime_method"] != "ShowRuntimeControlledLaunch" ||
		payload["owner_service_boundary"] != "go-runtime-owner-in-process-service" ||
		payload["desktop_callable_action_id"] != appidentity.KnownAppKDERuntimeStatusLaunchAction ||
		payload["desktop_callable_route"] != "kde-dbus-runtime-status-action" ||
		payload["desktop_callable_runtime_method"] != "ShowRuntimeControlledLaunch" ||
		payload["desktop_callable_execution_type"] != appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		payload["desktop_evidence_handle_forwarded"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false ||
		payload["desktop_runtime_owner_adapter_used"] != true ||
		payload["desktop_state_root_supplied_by_runtime_owner"] != true ||
		payload["desktop_cache_root_supplied_by_runtime_owner"] != true ||
		payload["desktop_launcher_supplied_by_runtime_owner"] != true ||
		payload["desktop_timeout_supplied_by_runtime_owner"] != true ||
		payload["runtime_owner_service_action_dispatch"] != true ||
		payload["runtime_owner_service_call_ready"] != true ||
		payload["runtime_owner_service_supplies_owner_inputs"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["write_methods_enabled"] != false ||
		payload["dispatch_ready"] != true ||
		payload["evidence_handoff_consumed"] != true ||
		payload["evidence_digest_verified"] != true ||
		payload["evidence_relative_path"] != evidenceRecord.EvidenceRelativePath ||
		payload["launch_authorization_receipt_id"] != launchReceiptID ||
		payload["session_gated_review_receipt_id"] != reviewReceiptID ||
		payload["controlled_execution_session_id"] != sessionID ||
		payload["state_root_injected_by_runtime"] != true ||
		payload["state_root_supplied_by_runtime"] != true ||
		payload["kde_state_root_access"] != false ||
		payload["managed_launcher_invoked"] != true ||
		payload["existing_managed_launcher_invoked"] != true ||
		payload["launcher_output_json_observed"] != true ||
		payload["delegated_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["delegated_status"] != "passed" ||
		payload["delegated_session_gated_review_receipt_id"] != reviewReceiptID ||
		payload["delegated_controlled_execution_session_id"] != sessionID ||
		payload["delegated_host_root_modified"] != false ||
		payload["delegated_docker_socket_mounted"] != false ||
		payload["delegated_broad_host_mount_required"] != false ||
		payload["delegated_raw_command_exposed"] != false ||
		payload["delegated_backend_details_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Runtime controlled launch action payload: %#v", payload)
	}
	projection := payload["compatibility_center_known_app_evidence"].(map[string]any)
	if projection["projection_type"] != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		projection["request_type"] != "windows-known-app-dispatch-smoke" ||
		projection["status"] != "passed" ||
		projection["app_id"] != "7zr" ||
		projection["controlled_execution_session_id"] != sessionID ||
		projection["compatibility_center_projection_ready"] != true ||
		projection["kde_center_projection_ready"] != true ||
		projection["state_root_path_exposed"] != false ||
		projection["managed_launcher_path_exposed"] != false ||
		projection["raw_launcher_output_exposed"] != false ||
		projection["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime controlled launch projection: %#v", projection)
	}
	argsData, err := os.ReadFile(fakeArgsPath)
	if err != nil {
		t.Fatalf("ReadFile fake launcher args returned error: %v", err)
	}
	argsText := string(argsData)
	for _, token := range []string{"--state-root\n" + stateRoot, "--receipt-id\n" + launchReceiptID, "--review-receipt-id\n" + reviewReceiptID, "--session-id\n" + sessionID, "--timeout\n1s"} {
		if !strings.Contains(argsText, token) {
			t.Fatalf("fake launcher args missing %q: %s", token, argsText)
		}
	}
	text := strings.ToLower(string(call.Payload))
	for _, forbidden := range []string{strings.ToLower(stateRoot), strings.ToLower(fakeLauncher), ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Runtime controlled launch service payload exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestServiceCallServesOwnerLocalRestrictedOwnerSmokeReceiptLookup(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetRestrictedOwnerSmokeReceiptLookupPreview", []string{RestrictedOwnerSmokeOpaqueReceiptID})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetRestrictedOwnerSmokeReceiptLookupPreview" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected restricted owner smoke lookup service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetRestrictedOwnerSmokeReceiptLookupPreview" ||
		dispatch["go_command"] != "restricted-owner-smoke-receipt-lookup-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested restricted owner smoke lookup dispatch: %#v", dispatch)
	}
	nested := dispatch["payload"].(map[string]any)
	if nested["request_type"] != "restricted-owner-smoke-receipt-lookup-preview" ||
		nested["owner_managed_lookup"] != true ||
		nested["caller_state_root_required"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected nested restricted owner smoke lookup payload: %#v", nested)
	}
}

func TestServiceCallServesOwnerLocalRestrictedOwnerSmokeReceiptFanOut(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetRestrictedOwnerSmokeReceiptFanOut", []string{RestrictedOwnerSmokeOpaqueReceiptID})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected restricted owner smoke fan-out service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		dispatch["go_command"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested restricted owner smoke fan-out dispatch: %#v", dispatch)
	}
	nested := dispatch["payload"].(map[string]any)
	if nested["request_type"] != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		nested["owner_managed_lookup"] != true ||
		nested["caller_state_root_required"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected nested restricted owner smoke fan-out payload: %#v", nested)
	}
}

func TestServiceCallServesOwnerLocalKDETestLaunchMaterializationReceiptLookup(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetKDETestLaunchMaterializationReceiptLookupPreview", []string{appidentity.KDETestLaunchMaterializationOpaqueReceiptID})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetKDETestLaunchMaterializationReceiptLookupPreview" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected materialization receipt lookup service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetKDETestLaunchMaterializationReceiptLookupPreview" ||
		dispatch["go_command"] != "kde-test-launch-materialization-receipt-lookup-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested materialization receipt lookup dispatch: %#v", dispatch)
	}
	nested := dispatch["payload"].(map[string]any)
	if nested["request_type"] != "kde-test-launch-materialization-receipt-lookup-preview" ||
		nested["owner_managed_opaque_receipt_lookup_ready"] != true ||
		nested["requires_caller_state_root"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["materialization_writes_enabled"] != false ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected nested materialization receipt lookup payload: %#v", nested)
	}
}

func TestServiceCallServesOwnerLocalKDETestLaunchMaterializationFanOut(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	call, err := service.Call("GetKDETestLaunchMaterializationFanOut", []string{appidentity.KDETestLaunchMaterializationOpaqueReceiptID})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetKDETestLaunchMaterializationFanOut" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected materialization fan-out owner route service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetKDETestLaunchMaterializationFanOut" ||
		dispatch["go_command"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested materialization fan-out owner route dispatch: %#v", dispatch)
	}
	nested := dispatch["payload"].(map[string]any)
	if nested["request_type"] != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		nested["owner_managed_opaque_receipt_lookup_ready"] != true ||
		nested["requires_caller_state_root"] != false ||
		nested["receipt_lookup_state"] != "missing-receipt" ||
		nested["fan_out_result_state"] != "missing-receipt-fail-closed" ||
		nested["owner_local_route_candidate_ready"] != true ||
		nested["production_dbus_exposure_ready"] != false ||
		nested["fan_out_writes_enabled"] != false ||
		nested["backend_launch_enabled"] != false ||
		nested["host_root_modified"] != false {
		t.Fatalf("unexpected nested materialization fan-out owner route payload: %#v", nested)
	}
}

func TestServiceCallDeniesWriteMethods(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("Launch", nil)
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}

	if call.CallType != "write-denial" ||
		call.Method != "Launch" ||
		call.ReadOnlyDispatch ||
		!call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.ErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		call.EventLoopStarted ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.HostRootModified {
		t.Fatalf("unexpected write service call: %#v", call)
	}

	var denial map[string]any
	if err := json.Unmarshal(call.Payload, &denial); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if denial["method"] != "Launch" ||
		denial["dispatch_enabled"] != false ||
		denial["request_created"] != false {
		t.Fatalf("unexpected nested write denial payload: %#v", denial)
	}
}

func TestServiceCallServesOwnerLocalNotificationDigest(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("GetKDENotificationDigestPreview", []string{
		"org.xnix.sample.notepad",
		"blocked-action:execution-blocked",
	})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetKDENotificationDigestPreview" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.HostRootModified {
		t.Fatalf("unexpected notification digest service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetKDENotificationDigestPreview" ||
		dispatch["go_command"] != "kde-notification-digest-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested notification digest dispatch: %#v", dispatch)
	}
}

func TestServiceCallServesOwnerLocalSignedRecipeVerification(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("GetSignedRecipeVerificationPreview", []string{"org.xnix.sample.notepad"})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetSignedRecipeVerificationPreview" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected signed recipe service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetSignedRecipeVerificationPreview" ||
		dispatch["go_command"] != "signed-recipe-verifier-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested signed recipe dispatch: %#v", dispatch)
	}
}

func TestServiceCallServesOwnerLocalRestrictedSmokePacket(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("GetRestrictedProductSmokePacketPreview", nil)
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if call.Method != "GetRestrictedProductSmokePacketPreview" ||
		call.CallType != "read-dispatch" ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.NetworkRequired ||
		call.HostRootModified ||
		call.BackendDetailsExposed {
		t.Fatalf("unexpected restricted smoke service call: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["method"] != "GetRestrictedProductSmokePacketPreview" ||
		dispatch["go_command"] != "restricted-product-smoke-packet-preview" ||
		dispatch["route_source"] != "go-owner-local-preview" {
		t.Fatalf("unexpected nested restricted smoke dispatch: %#v", dispatch)
	}
}

func TestServiceCallRejectsUnknownMethods(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	_, err = service.Call("GetUnknownThing", nil)
	if err == nil || !strings.Contains(err.Error(), "unsupported Runtime owner service method") {
		t.Fatalf("Call must reject unknown methods, got %v", err)
	}
}

func writeOwnerRuntimeStatusLaunchExecutionFixture(t *testing.T, stateRoot string) string {
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

func recordOwnerRuntimeStatusLaunchEvidenceFixture(t *testing.T, stateRoot string) appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord {
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

func writeOwnerFakeRuntimeStatusManagedLauncher(t *testing.T) (string, string) {
	t.Helper()
	dir := t.TempDir()
	launcherPath := filepath.Join(dir, "fake-xnix-compat-launch")
	argsPath := filepath.Join(dir, "launcher-args.txt")
	t.Setenv("XNIX_OWNER_FAKE_LAUNCHER_ARGS_FILE", argsPath)
	script := "#!/bin/sh\nprintf '%s\n' \"$@\" > \"$XNIX_OWNER_FAKE_LAUNCHER_ARGS_FILE\"\nprintf '%s\n' '{\"request_type\":\"windows-known-app-dispatch-smoke\",\"status\":\"passed\",\"guest_boundary\":\"managed-known-app-guest-smoke\",\"runtime_owned_dispatch\":true,\"artifact_verified\":true,\"marker_observed\":true,\"smoke_passed\":true,\"execution_started\":true,\"backend_process_started\":false,\"session_gated_controlled_dispatch_consumed\":true,\"session_gated_controlled_dispatch_state\":\"created-after-session-gated-review\",\"session_gated_review_receipt_id\":\"known-app-session-gated-launch-review-7zr-26.02-known-app-controlled-execution-session-7zr-26.02\",\"launch_authorization_receipt_id\":\"known-app-launch-authorization-7zr-26.02\",\"controlled_execution_session_consumed\":true,\"controlled_execution_session_id\":\"known-app-controlled-execution-session-7zr-26.02\",\"controlled_session_digest_verified\":true,\"controlled_session_relative_path\":\"execution-ledger/sessions/known-app-controlled-execution-session-7zr-26.02.json\",\"runtime_owner_consumable_session\":true,\"kde_read_model_consumable_session\":true,\"controlled_session_live_state_observed\":false,\"controlled_session_registered\":false,\"controlled_session_window_observed\":false,\"controlled_session_host_root_modified\":false,\"controlled_session_backend_process_start\":false,\"host_root_modified\":false,\"docker_socket_mounted\":false,\"broad_host_mount_required\":false,\"raw_command_exposed\":false,\"backend_details_exposed\":false}'\n"
	if err := os.WriteFile(launcherPath, []byte(script), 0o755); err != nil {
		t.Fatalf("WriteFile fake launcher returned error: %v", err)
	}
	return launcherPath, argsPath
}
