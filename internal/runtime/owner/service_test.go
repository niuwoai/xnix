package owner

import (
	"encoding/json"
	"strings"
	"testing"
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
