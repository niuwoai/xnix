package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRecordKnownAppVerifiedCatalogLaunchMaterializationConsumesHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := filepath.Join(t.TempDir(), "acceptance.json")
	if err := os.WriteFile(acceptancePath, knownAppVerifiedCatalogLaunchHandoffAcceptanceFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile acceptance returned error: %v", err)
	}
	handoff, err := RecordKnownAppVerifiedCatalogLaunchHandoff(KnownAppVerifiedCatalogLaunchHandoffRequest{
		StateRoot:      stateRoot,
		AcceptancePath: acceptancePath,
		RecordedAtUTC:  time.Date(2026, 7, 27, 3, 4, 5, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("RecordKnownAppVerifiedCatalogLaunchHandoff returned error: %v", err)
	}
	record, err := RecordKnownAppVerifiedCatalogLaunchMaterialization(KnownAppVerifiedCatalogLaunchMaterializationRequest{
		StateRoot:           stateRoot,
		HandoffRelativePath: handoff.HandoffRelativePath,
		CacheRoot:           filepath.Join(t.TempDir(), "missing-cache"),
		RecordedAtUTC:       time.Date(2026, 7, 27, 4, 5, 6, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("RecordKnownAppVerifiedCatalogLaunchMaterialization returned error: %v", err)
	}
	if record.Version != currentProjectVersion(t) ||
		record.SchemaVersion != KnownAppVerifiedCatalogLaunchMaterializationSchemaVersion ||
		record.RequestType != KnownAppVerifiedCatalogLaunchMaterializationRequestType ||
		record.Source != KnownAppVerifiedCatalogLaunchHandoffRequestType+"+runtime-owner-materialization" ||
		record.RuntimeMethod != "RecordKnownAppVerifiedCatalogLaunchMaterialization" ||
		record.ReadMethod != "GetKnownAppVerifiedCatalogLaunchMaterialization" ||
		record.AppID != "7zr" ||
		record.DisplayName != "7-Zip standalone console executable" ||
		record.AppVersion != "26.02" ||
		!record.HandoffConsumed ||
		record.HandoffRelativePath != handoff.HandoffRelativePath ||
		!record.HandoffDigestVerified ||
		record.AcceptanceRequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType ||
		record.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" ||
		!record.AcceptanceReady ||
		!record.RunPlanMatched ||
		!record.ExistingWindowsApp ||
		!record.KnownPortableCatalogBacked ||
		!record.LaunchAttempted ||
		!record.ChecksumVerified ||
		!record.MarkerObserved ||
		!record.RuntimeStartedIsolatedGuest ||
		!record.IsolatedGuestExecutionObserved ||
		!record.CompatibilityEngineExecutionObserved ||
		!record.OutputRedacted ||
		!record.Q4ExecutionObserved ||
		!record.HostCompilationAvoided {
		t.Fatalf("unexpected materialization evidence: %#v", record)
	}
	if record.LaunchAuthorizationReceiptID != KnownAppLaunchAuthorizationReceiptID("7zr", "26.02") ||
		!record.LaunchAuthorizationReceiptRecorded ||
		record.ControlledExecutionSessionID != "" ||
		record.ControlledSessionRecordState != "blocked" ||
		record.MaterializationState != "blocked-managed-artifact-required" ||
		record.MaterializationReady ||
		record.RuntimeOwnerMaterialized ||
		record.SessionGatedReviewReceiptRecorded ||
		record.DispatchRunnerRequired ||
		record.NextOwnerAction != "prepare-managed-artifact" {
		t.Fatalf("unexpected blocked materialization state: %#v", record)
	}
	if !record.OwnerMaterializationRequired ||
		!record.RuntimeOwnerServiceSuppliesInputs ||
		!record.RuntimeOwned ||
		record.RuntimeOwnedDispatch ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		!record.KDEForwardsOnlyEvidenceHandle ||
		record.DesktopReceiptFieldsReconstructed ||
		record.DesktopKDEStateRootAccess {
		t.Fatalf("unexpected ownership fields: %#v", record)
	}
	if record.StateRootPathExposed ||
		record.HandoffPathExposed ||
		record.ReceiptPathExposed ||
		record.SessionPathExposed ||
		record.ReviewReceiptPathExposed ||
		record.RemoteHostExposed ||
		record.RawOutputExposed ||
		record.RuntimeArgvExposed ||
		record.RunnerPathExposed ||
		record.BackendDetailsExposed ||
		record.HostRootModified ||
		record.PrivilegedContainerRequired ||
		record.HostNetworkingRequired ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired ||
		record.DesktopLaunchEnabled ||
		record.BackendLaunchEnabled ||
		record.ExecutionStarted ||
		record.BackendProcessStarted ||
		record.DispatchAllowed ||
		record.DispatchStarted {
		t.Fatalf("materialization opened unsafe gates: %#v", record)
	}
	if strings.Contains(record.DesktopSafeSummary, stateRoot) || strings.Contains(record.DesktopSafeSummary, acceptancePath) {
		t.Fatalf("materialization summary exposed local paths: %s", record.DesktopSafeSummary)
	}
}

func TestRecordKnownAppVerifiedCatalogLaunchMaterializationRejectsUnsafePath(t *testing.T) {
	_, err := RecordKnownAppVerifiedCatalogLaunchMaterialization(KnownAppVerifiedCatalogLaunchMaterializationRequest{
		StateRoot:           t.TempDir(),
		HandoffRelativePath: "../outside.json",
	})
	if err == nil || !strings.Contains(err.Error(), "requires a safe relative path") {
		t.Fatalf("unsafe handoff path must be rejected, got: %v", err)
	}
}

func TestRecordKnownAppVerifiedCatalogDispatchRequestWritesOwnerOnlyRunnerRequest(t *testing.T) {
	stateRoot := t.TempDir()
	cacheRoot := filepath.Join(t.TempDir(), "known-cache")
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	materialization := readyKnownAppVerifiedCatalogLaunchMaterializationFixture(sessionID)

	record, err := RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization(stateRoot, cacheRoot, "xnix-compat-launch", "", materialization, time.Date(2026, 7, 27, 5, 6, 7, 0, time.UTC))
	if err != nil {
		t.Fatalf("RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppVerifiedCatalogDispatchRequestSchemaVersion ||
		record.RequestType != KnownAppVerifiedCatalogDispatchRequestType ||
		record.Source != KnownAppVerifiedCatalogLaunchMaterializationRequestType+"+dispatch-runner-request" ||
		record.RuntimeMethod != "RecordKnownAppVerifiedCatalogDispatchRequest" ||
		record.ReadMethod != "GetKnownAppVerifiedCatalogDispatchRequest" ||
		record.AppID != "7zr" ||
		record.DisplayName != "7-Zip standalone console executable" ||
		record.AppVersion != "26.02" ||
		!record.HandoffConsumed ||
		record.HandoffRelativePath != materialization.HandoffRelativePath ||
		record.MaterializationState != "persisted-runtime-owned-launch-session" ||
		!record.MaterializationReady ||
		!record.RuntimeOwnerMaterialized ||
		record.LaunchAuthorizationReceiptID != KnownAppLaunchAuthorizationReceiptID("7zr", "26.02") ||
		record.SessionGatedReviewReceiptID != KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID) ||
		record.ControlledExecutionSessionID != sessionID ||
		record.ControlledSessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		!record.ControlledSessionDigestVerified {
		t.Fatalf("unexpected dispatch request identity evidence: %#v", record)
	}
	if record.DispatchRequestRecordState != "persisted-runtime-owned-dispatch-request" ||
		!record.DispatchRequestWritten ||
		record.DispatchRequestID == "" ||
		record.DispatchRequestRelativePath == "" ||
		record.DispatchRequestSHA256 == "" ||
		record.DispatchRunnerRequestType != "windows-known-app-dispatch-smoke" ||
		record.DispatchRunnerName != "xnix-compat-launch" ||
		record.DispatchRunnerArgumentCount != 15 ||
		!record.RuntimeOwnerDispatchInputsReady ||
		!record.RuntimeOwnedDispatch ||
		record.NextOwnerAction != "dispatch-runner-execute-request" {
		t.Fatalf("unexpected dispatch request state: %#v", record)
	}
	if !record.RuntimeOwnerServiceSuppliesInputs ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		!record.KDEForwardsOnlyEvidenceHandle ||
		record.DesktopReceiptFieldsReconstructed ||
		record.DesktopKDEStateRootAccess ||
		!record.GuestTransportRequired ||
		!record.RunnerTransportArgsDeferred {
		t.Fatalf("unexpected dispatch ownership fields: %#v", record)
	}
	if record.StateRootPathExposed ||
		record.CacheRootPathExposed ||
		record.DispatchRequestPathExposed ||
		record.RunnerPathExposed ||
		record.BackendDetailsExposed ||
		record.RawOutputExposed ||
		record.RawCommandExposed ||
		record.DispatchRunnerArgumentValuesExposed ||
		record.HostRootModified ||
		record.PrivilegedContainerRequired ||
		record.HostNetworkingRequired ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired ||
		record.DispatchAllowed ||
		record.DispatchStarted ||
		record.DesktopLaunchEnabled ||
		record.BackendLaunchEnabled ||
		record.ExecutionStarted ||
		record.BackendProcessStarted {
		t.Fatalf("dispatch request record opened unsafe gates: %#v", record)
	}
	encoded, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) || strings.Contains(string(encoded), cacheRoot) {
		t.Fatalf("dispatch request record exposed owner paths: %s", string(encoded))
	}
	requestPath := filepath.Join(stateRoot, filepath.FromSlash(record.DispatchRequestRelativePath))
	requestBytes, err := os.ReadFile(requestPath)
	if err != nil {
		t.Fatalf("dispatch request should be written under state root: %v", err)
	}
	var payload knownAppVerifiedCatalogDispatchRequestPayload
	if err := json.Unmarshal(requestBytes, &payload); err != nil {
		t.Fatalf("dispatch request payload must be JSON: %v", err)
	}
	if payload.RunnerRequestType != "windows-known-app-dispatch-smoke" ||
		payload.LaunchAuthorizationReceiptID != record.LaunchAuthorizationReceiptID ||
		payload.SessionGatedReviewReceiptID != record.SessionGatedReviewReceiptID ||
		payload.ControlledExecutionSessionID != record.ControlledExecutionSessionID ||
		payload.GuestBoundary != "managed-known-app-guest-smoke" ||
		!containsString(payload.RunnerArgv, "--state-root") ||
		!containsString(payload.RunnerArgv, stateRoot) ||
		!containsString(payload.RunnerArgv, "--cache-root") ||
		!containsString(payload.RunnerArgv, cacheRoot) {
		t.Fatalf("unexpected internal dispatch request payload: %#v", payload)
	}
}

func TestRecordKnownAppVerifiedCatalogDispatchRequestBlocksUntilMaterializationReady(t *testing.T) {
	stateRoot := t.TempDir()
	materialization := readyKnownAppVerifiedCatalogLaunchMaterializationFixture(KnownAppControlledExecutionSessionID("7zr", "26.02"))
	materialization.MaterializationReady = false
	materialization.MaterializationState = "blocked-managed-artifact-required"
	materialization.RuntimeOwnerMaterialized = false
	materialization.RuntimeOwnedDispatch = false
	materialization.SessionGatedReviewReceiptRecorded = false
	materialization.SessionGatedReviewReceiptID = ""
	materialization.ControlledExecutionSessionID = ""
	materialization.ControlledSessionRelativePath = ""
	materialization.ControlledSessionDigestVerified = false
	materialization.ControlledSessionRecordState = "blocked"
	materialization.DispatchRunnerRequired = false

	record, err := RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization(stateRoot, filepath.Join(t.TempDir(), "cache"), "xnix-compat-launch", "", materialization, time.Date(2026, 7, 27, 6, 7, 8, 0, time.UTC))
	if err != nil {
		t.Fatalf("RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization returned error: %v", err)
	}
	if record.DispatchRequestRecordState != "blocked-materialization-required" ||
		record.DispatchRequestWritten ||
		record.DispatchRequestRelativePath != "" ||
		record.RuntimeOwnerDispatchInputsReady ||
		record.RuntimeOwnedDispatch ||
		record.NextOwnerAction != "materialize-launch-session" ||
		record.BlockedReason == "" {
		t.Fatalf("unexpected blocked dispatch request record: %#v", record)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "runtime", "known-app-verified-catalog-dispatch-requests")); !os.IsNotExist(err) {
		t.Fatalf("blocked dispatch request must not create request directory: %v", err)
	}
}

func readyKnownAppVerifiedCatalogLaunchMaterializationFixture(sessionID string) KnownAppVerifiedCatalogLaunchMaterializationRecord {
	return KnownAppVerifiedCatalogLaunchMaterializationRecord{
		Version:                              currentProjectVersionForAppIdentityTest(),
		SchemaVersion:                        KnownAppVerifiedCatalogLaunchMaterializationSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogLaunchMaterializationRequestType,
		Source:                               KnownAppVerifiedCatalogLaunchHandoffRequestType + "+runtime-owner-materialization",
		RuntimeMethod:                        "RecordKnownAppVerifiedCatalogLaunchMaterialization",
		ReadMethod:                           "GetKnownAppVerifiedCatalogLaunchMaterialization",
		AppID:                                "7zr",
		DisplayName:                          "7-Zip standalone console executable",
		AppVersion:                           "26.02",
		HandoffConsumed:                      true,
		HandoffRelativePath:                  "runtime/known-app-verified-catalog-launch-handoffs/known-app-verified-catalog-launch-handoff-7zr-26.02-verified-catalog-app-q4-real-run-acceptance.json",
		HandoffDigestVerified:                true,
		AcceptanceRequestType:                KnownAppVerifiedCatalogRunAcceptanceRequestType,
		AcceptanceType:                       "verified-catalog-app-q4-real-run-acceptance",
		AcceptanceReady:                      true,
		RunPlanMatched:                       true,
		ExistingWindowsApp:                   true,
		KnownPortableCatalogBacked:           true,
		LaunchAttempted:                      true,
		ChecksumVerified:                     true,
		MarkerObserved:                       true,
		RuntimeStartedIsolatedGuest:          true,
		IsolatedGuestExecutionObserved:       true,
		CompatibilityEngineExecutionObserved: true,
		OutputRedacted:                       true,
		Q4ExecutionObserved:                  true,
		HostCompilationAvoided:               true,
		LaunchAuthorizationReceiptID:         KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		LaunchAuthorizationReceiptRecorded:   true,
		ControlledExecutionSessionID:         sessionID,
		ControlledSessionRecordState:         "persisted",
		ControlledSessionRelativePath:        "execution-ledger/sessions/" + sessionID + ".json",
		ControlledSessionDigestVerified:      true,
		SessionGatedReviewReceiptID:          KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID),
		SessionGatedReviewReceiptRecorded:    true,
		MaterializationState:                 "persisted-runtime-owned-launch-session",
		MaterializationReady:                 true,
		RuntimeOwnerMaterialized:             true,
		OwnerMaterializationRequired:         true,
		RuntimeOwnerServiceSuppliesInputs:    true,
		DispatchRunnerRequired:               true,
		DispatchAllowed:                      false,
		DispatchStarted:                      false,
		RuntimeOwned:                         true,
		RuntimeOwnedDispatch:                 true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		KDEForwardsOnlyEvidenceHandle:        true,
		DesktopReceiptFieldsReconstructed:    false,
		DesktopKDEStateRootAccess:            false,
		StateRootPathExposed:                 false,
		HandoffPathExposed:                   false,
		ReceiptPathExposed:                   false,
		SessionPathExposed:                   false,
		ReviewReceiptPathExposed:             false,
		RemoteHostExposed:                    false,
		RawOutputExposed:                     false,
		RuntimeArgvExposed:                   false,
		RunnerPathExposed:                    false,
		BackendDetailsExposed:                false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		DesktopLaunchEnabled:                 false,
		BackendLaunchEnabled:                 false,
		ExecutionStarted:                     false,
		BackendProcessStarted:                false,
		RecordedAtUTC:                        "2026-07-27T05:06:07Z",
		NextOwnerAction:                      "dispatch-runner-consume-session",
		DesktopSafeSummary:                   "7-Zip standalone console executable verified catalog launch handoff materialized Runtime-owned receipts and a controlled execution session; execution still waits for the dispatch runner.",
	}
}

func currentProjectVersionForAppIdentityTest() string {
	content, err := os.ReadFile(filepath.Join("..", "..", "..", "VERSION"))
	if err != nil {
		return "0.2.640-test"
	}
	return strings.TrimSpace(string(content))
}
