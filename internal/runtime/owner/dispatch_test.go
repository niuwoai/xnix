package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDispatchReadRendersOwnerReadinessWithoutBusOwnership(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetRuntimeOwnerReadiness", nil)
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}

	if dispatch.Version != currentProjectVersion(t) ||
		dispatch.SchemaVersion != "xnix.runtime.owner_read_dispatch.v1" ||
		dispatch.RequestType != "runtime-owner-read-dispatch" ||
		dispatch.DispatchType != "go-owner-read-dispatch" ||
		dispatch.Source != "go-runtime-owner-candidate+in-process-read-dispatch" ||
		dispatch.Method != "GetRuntimeOwnerReadiness" ||
		dispatch.RouteSource != "go-runtime-cli" ||
		dispatch.GoCommand != "runtime-owner-readiness-preview" ||
		dispatch.RouteStatus != "go-preview-ready" ||
		!dispatch.RouteReady {
		t.Fatalf("unexpected dispatch schema: %#v", dispatch)
	}
	if len(dispatch.Args) != 0 {
		t.Fatalf("readiness dispatch args = %#v, want empty", dispatch.Args)
	}
	if !dispatch.ReadOnlyDispatch ||
		dispatch.WriteMethod ||
		dispatch.WriteMethodsEnabled ||
		!dispatch.RuntimeOwned ||
		!dispatch.GoRuntimeBacked ||
		dispatch.KDEPolicyOwner ||
		dispatch.KDEMayClaimRuntimeOwnership ||
		dispatch.EventLoopStarted ||
		dispatch.SessionBusClaimed ||
		dispatch.ProductionBusClaimed ||
		dispatch.SystemServiceStarted ||
		dispatch.NetworkRequired ||
		dispatch.HostRootModified ||
		dispatch.PrivilegedContainerRequired ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected dispatch safety flags: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "runtime-owner-readiness-preview" ||
		payload["readiness_type"] != "runtime-owner-readiness" ||
		payload["production_bus_claimed"] != false ||
		payload["write_methods_enabled"] != false {
		t.Fatalf("unexpected readiness payload: %#v", payload)
	}
	if len(dispatch.BlockedActions) != 6 ||
		dispatch.BlockedActions[0] != "claim a D-Bus name from read dispatch preview" ||
		len(dispatch.NextRequirements) != 4 ||
		dispatch.NextRequirements[0] != "Bind read dispatch to the restricted owner event loop." {
		t.Fatalf("unexpected guidance: actions=%#v next=%#v", dispatch.BlockedActions, dispatch.NextRequirements)
	}
}

func TestDispatchReadRendersWriteGatePayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetRuntimeWriteGate", []string{"Launch"})
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetRuntimeWriteGate" ||
		dispatch.GoCommand != "runtime-write-gate-preview" ||
		len(dispatch.Args) != 1 ||
		dispatch.Args[0] != "Launch" {
		t.Fatalf("unexpected write gate dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "runtime-write-gate-preview" ||
		payload["method_name"] != "Launch" ||
		payload["dispatch_enabled"] != false ||
		payload["request_object_created"] != false {
		t.Fatalf("unexpected write gate payload: %#v", payload)
	}
}

func TestDispatchReadRejectsWriteMethodsAndBadArity(t *testing.T) {
	if _, err := DispatchRead(projectRoot(t), "Launch", nil); err == nil || !strings.Contains(err.Error(), "write method") {
		t.Fatalf("DispatchRead must reject write methods, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetRuntimeWriteGate", nil); err == nil || !strings.Contains(err.Error(), "requires 1 argument") {
		t.Fatalf("DispatchRead must reject bad arity, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetKDENotificationDigestPreview", []string{"org.xnix.sample.notepad"}); err == nil || !strings.Contains(err.Error(), "at least one event") {
		t.Fatalf("DispatchRead must reject a missing digest event, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetKDENotificationDigestPreview", []string{"org.xnix.sample.notepad", "invalid"}); err == nil || !strings.Contains(err.Error(), "<group>:<event-id>") {
		t.Fatalf("DispatchRead must reject a malformed digest event, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetSignedRecipeVerificationPreview", nil); err == nil || !strings.Contains(err.Error(), "requires 1 argument") {
		t.Fatalf("DispatchRead must reject missing signed recipe arguments, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetRestrictedProductSmokePacketPreview", []string{"unexpected"}); err == nil || !strings.Contains(err.Error(), "requires 0 argument") {
		t.Fatalf("DispatchRead must reject restricted smoke packet arguments, got %v", err)
	}
	if _, err := DispatchRead(projectRoot(t), "GetUnknownRuntimeMethod", nil); err == nil || !strings.Contains(err.Error(), "unsupported owner read dispatch method") {
		t.Fatalf("DispatchRead must reject unsupported methods, got %v", err)
	}
}

func TestDispatchReadRendersApplicationPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetApplication", []string{"org.xnix.sample.notepad"})
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetApplication" ||
		dispatch.GoCommand != "application-preview" ||
		dispatch.RouteSource != "go-runtime-cli" ||
		dispatch.RouteStatus != "go-preview-ready" ||
		!dispatch.RouteReady ||
		dispatch.EventLoopStarted ||
		dispatch.SessionBusClaimed {
		t.Fatalf("unexpected application dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "application-preview" ||
		payload["runtime_method"] != "GetApplication" ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected application dispatch payload: %#v", payload)
	}
}

func TestDispatchReadRendersWindowsCompatibilityWorkstreamsAsOwnerLocalPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetWindowsCompatibilityWorkstreamsPreview", nil)
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetWindowsCompatibilityWorkstreamsPreview" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "windows-compatibility-workstreams-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
		!dispatch.RouteReady ||
		len(dispatch.Args) != 0 ||
		dispatch.SessionBusClaimed ||
		dispatch.ProductionBusClaimed ||
		dispatch.WriteMethodsEnabled ||
		dispatch.HostRootModified ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected Windows compatibility dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "windows-compatibility-workstreams-preview" ||
		payload["read_method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		payload["official_desktop"] != "KDE Plasma" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected Windows compatibility dispatch payload: %#v", payload)
	}
}

func TestDispatchReadRendersKDENotificationDigestAsOwnerLocalPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetKDENotificationDigestPreview", []string{
		"org.xnix.sample.notepad",
		"needs-review:approval-required",
		"readiness-change:readiness-pending",
	})
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetKDENotificationDigestPreview" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "kde-notification-digest-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
		!dispatch.RouteReady ||
		!dispatch.ReadOnlyDispatch ||
		dispatch.WriteMethodsEnabled ||
		dispatch.SessionBusClaimed ||
		dispatch.ProductionBusClaimed ||
		dispatch.HostRootModified ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected notification digest dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	counts := payload["counts"].(map[string]any)
	if payload["request_type"] != "kde-notification-digest-preview" ||
		payload["read_method"] != "GetKDENotificationDigestPreview" ||
		payload["application_id"] != "org.xnix.sample.notepad" ||
		counts["digest_entry_count"] != float64(2) ||
		payload["runtime_owned"] != true ||
		payload["notifications_sent"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected notification digest payload: %#v", payload)
	}
}

func TestDispatchReadRendersSignedRecipeVerificationAsOwnerLocalPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetSignedRecipeVerificationPreview", []string{"org.xnix.sample.notepad"})
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetSignedRecipeVerificationPreview" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "signed-recipe-verifier-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
		!dispatch.RouteReady ||
		!dispatch.ReadOnlyDispatch ||
		dispatch.WriteMethodsEnabled ||
		dispatch.SessionBusClaimed ||
		dispatch.ProductionBusClaimed ||
		dispatch.HostRootModified ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected signed recipe dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "signed-recipe-verifier-preview" ||
		payload["evidence_type"] != "owner-registry-signed-recipe-verifier-evidence" ||
		payload["application_id"] != "org.xnix.sample.notepad" ||
		payload["verification_state"] != "production-signature-required" ||
		payload["recipe_digest_verified"] != true ||
		payload["signature_verified"] != false ||
		payload["production_key_configured"] != false ||
		payload["private_key_loaded"] != false ||
		payload["recipe_path_exposed"] != false ||
		payload["public_key_path_exposed"] != false ||
		payload["signature_material_exposed"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected signed recipe payload: %#v", payload)
	}
}

func TestDispatchReadRendersRestrictedProductSmokePacketAsOwnerLocalPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetRestrictedProductSmokePacketPreview", nil)
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetRestrictedProductSmokePacketPreview" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "restricted-product-smoke-packet-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
		!dispatch.RouteReady ||
		!dispatch.ReadOnlyDispatch ||
		dispatch.WriteMethodsEnabled ||
		dispatch.SessionBusClaimed ||
		dispatch.ProductionBusClaimed ||
		dispatch.HostRootModified ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected restricted smoke dispatch metadata: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "restricted-product-smoke-packet-preview" ||
		payload["packet_prepared"] != true ||
		payload["ready_for_authorized_smoke"] != true ||
		payload["human_authorization_required"] != true ||
		payload["execution_authorized"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["product_smoke_executed"] != false ||
		payload["serial_log_persisted"] != false ||
		payload["release_ready"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected restricted smoke payload: %#v", payload)
	}
}

func TestSupportedReadDispatchMethodsRenderPayloads(t *testing.T) {
	root := projectRoot(t)
	for _, method := range SupportedReadDispatchMethods() {
		dispatch, err := DispatchRead(root, method, sampleReadDispatchArgs(method))
		if err != nil {
			t.Fatalf("DispatchRead(%s) returned error: %v", method, err)
		}
		if dispatch.Method != method ||
			!dispatch.ReadOnlyDispatch ||
			dispatch.WriteMethod ||
			dispatch.WriteMethodsEnabled ||
			dispatch.EventLoopStarted ||
			dispatch.SessionBusClaimed ||
			dispatch.ProductionBusClaimed ||
			dispatch.HostRootModified ||
			dispatch.BackendDetailsExposed ||
			len(dispatch.Payload) == 0 {
			t.Fatalf("unexpected dispatch for %s: %#v", method, dispatch)
		}
	}
}

func TestSupportedReadDispatchMethodsAreStable(t *testing.T) {
	methods := SupportedReadDispatchMethods()
	want := []string{
		"ListApplications",
		"GetApplication",
		"GetDiagnostics",
		"GetEngineCatalog",
		"GetRunPlan",
		"GetDesktopActivationManifest",
		"GetDesktopActivationTransactionPreview",
		"GetDesktopActivationStatus",
		"GetDesktopEntryPlan",
		"GetDesktopIconPlan",
		"GetTaskManagerIdentityPlan",
		"GetKDEIntegrationStatus",
		"GetKDEShellIntegrationPlan",
		"GetKDEApplicationSurfacePlan",
		"GetDesktopResourceBridgePlan",
		"GetKWinWindowRulePlan",
		"GetFileAssociationPlan",
		"GetNotificationPlan",
		"GetTrayStatus",
		"GetKRunnerQueryPlan",
		"GetPortalRequestPlan",
		"GetApplicationStateRoot",
		"GetCompatibilityPackageSource",
		"GetCompatibilityAcquisitionPreflight",
		"GetCompatibilityArtifactManifest",
		"GetCompatibilityInstallPlan",
		"GetBackendBinding",
		"GetBackendCapabilityMatrix",
		"GetBackendSelectionPlan",
		"GetBackendLifecycle",
		"GetBackendEnvironmentPlan",
		"GetRepairPlan",
		"GetTestPlan",
		"GetTestResult",
		"GetExecutionReadiness",
		"GetLaunchIntent",
		"GetAIDiagnosticInput",
		"GetAIDiagnosticRecommendation",
		"GetAIRepairApprovalGate",
		"GetSnapshotPlan",
		"GetPortalAccessPolicy",
		"GetRuntimeServiceBinding",
		"GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerSmokePlan",
		"GetRuntimeMethodParityManifest",
		"GetRuntimeWriteGate",
		"GetCompatibilitySettings",
		"GetCompatibilitySettingsChangePlan",
		"GetCompatibilityModeSwitchPlan",
		"GetCompatibilityPermissionReviewPlan",
		"GetCompatibilityReviewFlowPlan",
		"GetCompatibilityActionQueue",
		"GetCompatibilityActionReviewReceipt",
		"GetCompatibilityCenterSummary",
		"GetKDECenterPage",
		"GetKDECenterPageSections",
		"GetKDECenterPageSectionDetail",
		"GetRuntimeOwnerProcess",
		"GetRuntimeOwnerRouteManifest",
		"GetRuntimeOwnerRecipeTrust",
		"GetRuntimeOwnerReadiness",
		"GetWindowsCompatibilityWorkstreamsPreview",
		"GetKDENotificationDigestPreview",
		"GetSignedRecipeVerificationPreview",
		"GetRestrictedProductSmokePacketPreview",
		"GetKDEOfflineApplicationIdentityPreview",
	}
	if !sameStrings(methods, want) {
		t.Fatalf("SupportedReadDispatchMethods = %#v, want %#v", methods, want)
	}
}

func TestDispatchReadRendersOfflineKDEIdentityAsOwnerLocalPayload(t *testing.T) {
	dispatch, err := DispatchRead(projectRoot(t), "GetKDEOfflineApplicationIdentityPreview", []string{"org.xnix.sample.notepad"})
	if err != nil {
		t.Fatalf("DispatchRead returned error: %v", err)
	}
	if dispatch.Method != "GetKDEOfflineApplicationIdentityPreview" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "kde-offline-application-identity-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
		!dispatch.RouteReady || !dispatch.ReadOnlyDispatch || dispatch.WriteMethodsEnabled ||
		dispatch.ProductionBusClaimed || dispatch.NetworkRequired || dispatch.HostRootModified ||
		dispatch.BackendDetailsExposed {
		t.Fatalf("unexpected offline KDE identity dispatch: %#v", dispatch)
	}
	var payload map[string]any
	if err := json.Unmarshal(dispatch.Payload, &payload); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if payload["request_type"] != "kde-offline-application-identity-preview" ||
		payload["application_id"] != "org.xnix.sample.notepad" ||
		payload["surface_count"] != float64(9) ||
		payload["cross_surface_identity_consistent"] != true ||
		payload["settings_persisted"] != false ||
		payload["compatibility_center_persisted"] != false ||
		payload["launch_enabled"] != false || payload["host_root_modified"] != false {
		t.Fatalf("unexpected offline KDE identity payload: %#v", payload)
	}
	if _, err := DispatchRead(projectRoot(t), "GetKDEOfflineApplicationIdentityPreview", []string{"org.xnix.sample.notepad", "/tmp/registry.json"}); err == nil {
		t.Fatal("offline KDE identity owner route accepted a caller path")
	}
}

func sampleReadDispatchArgs(method string) []string {
	const appID = "org.xnix.sample.notepad"
	switch method {
	case "ListApplications", "GetEngineCatalog", "GetKDEIntegrationStatus", "GetKDEShellIntegrationPlan",
		"GetTrayStatus", "GetBackendCapabilityMatrix", "GetRuntimeServiceBinding", "GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerSmokePlan", "GetRuntimeMethodParityManifest", "GetRuntimeOwnerProcess",
		"GetRuntimeOwnerRouteManifest", "GetRuntimeOwnerRecipeTrust", "GetRuntimeOwnerReadiness",
		"GetWindowsCompatibilityWorkstreamsPreview", "GetRestrictedProductSmokePacketPreview":
		return nil
	case "GetDesktopActivationTransactionPreview", "GetDesktopActivationStatus":
		return []string{appID, "development"}
	case "GetNotificationPlan":
		return []string{appID, "install-failed"}
	case "GetKDENotificationDigestPreview":
		return []string{appID, "needs-review:approval-required", "readiness-change:readiness-pending"}
	case "GetSignedRecipeVerificationPreview":
		return []string{appID}
	case "GetKRunnerQueryPlan":
		return []string{"notepad"}
	case "GetPortalRequestPlan", "GetPortalAccessPolicy":
		return []string{appID, "file-open"}
	case "GetCompatibilityInstallPlan":
		return []string{appID, "development"}
	case "GetRepairPlan":
		return []string{appID, "engine-binding-pending"}
	case "GetTestPlan", "GetTestResult":
		return []string{appID, "preflight"}
	case "GetAIDiagnosticInput", "GetAIDiagnosticRecommendation", "GetAIRepairApprovalGate":
		return []string{appID, "engine-binding-pending", "preflight"}
	case "GetSnapshotPlan":
		return []string{appID, "manual"}
	case "GetRuntimeWriteGate":
		return []string{"Launch"}
	case "GetCompatibilitySettingsChangePlan":
		return []string{appID, "run-mode", "mode", "automatic"}
	case "GetCompatibilityModeSwitchPlan":
		return []string{appID, "automatic"}
	case "GetCompatibilityReviewFlowPlan":
		return []string{appID, "run-mode", "mode", "automatic", "file-open"}
	case "GetCompatibilityActionReviewReceipt":
		return []string{appID, "review-file-manager-action", "reviewed"}
	case "GetKDECenterPage":
		return []string{appID, "reviewed"}
	case "GetKDECenterPageSections":
		return []string{appID, "reviewed"}
	case "GetKDECenterPageSectionDetail":
		return []string{appID, "overview", "reviewed"}
	default:
		return []string{appID}
	}
}
