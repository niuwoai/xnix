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

	if dispatch.Version != "0.2.229" ||
		dispatch.SchemaVersion != "xnix.runtime.owner_read_dispatch.v1" ||
		dispatch.RequestType != "runtime-owner-read-dispatch" ||
		dispatch.DispatchType != "go-owner-read-dispatch" ||
		dispatch.Source != "go-runtime-owner-candidate+in-process-read-dispatch" ||
		dispatch.Method != "GetRuntimeOwnerReadiness" ||
		dispatch.RouteSource != "go-owner-local-preview" ||
		dispatch.GoCommand != "runtime-owner-readiness-preview" ||
		dispatch.RouteStatus != "owner-local-preview-ready" ||
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
	}
	if !sameStrings(methods, want) {
		t.Fatalf("SupportedReadDispatchMethods = %#v, want %#v", methods, want)
	}
}

func sampleReadDispatchArgs(method string) []string {
	const appID = "org.xnix.sample.notepad"
	switch method {
	case "ListApplications", "GetEngineCatalog", "GetKDEIntegrationStatus", "GetKDEShellIntegrationPlan",
		"GetTrayStatus", "GetBackendCapabilityMatrix", "GetRuntimeServiceBinding", "GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerSmokePlan", "GetRuntimeMethodParityManifest", "GetRuntimeOwnerProcess",
		"GetRuntimeOwnerRouteManifest", "GetRuntimeOwnerRecipeTrust", "GetRuntimeOwnerReadiness":
		return nil
	case "GetDesktopActivationTransactionPreview", "GetDesktopActivationStatus":
		return []string{appID, "development"}
	case "GetNotificationPlan":
		return []string{appID, "install-failed"}
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
