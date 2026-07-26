package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/winapp"
)

func TestPrepareKnownAppKDERuntimeStatusLaunchExecutionRevalidatesRuntimeOwnedState(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppRuntimeStatusLaunchExecutionFixture(t, stateRoot)
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	plan, err := PrepareKnownAppKDERuntimeStatusLaunchExecution(KnownAppKDERuntimeStatusLaunchExecutionRequest{
		AppID:                        "7zr",
		StateRoot:                    stateRoot,
		LaunchAuthorizationReceiptID: launchReceiptID,
		SessionGatedReviewReceiptID:  reviewReceiptID,
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err != nil {
		t.Fatalf("PrepareKnownAppKDERuntimeStatusLaunchExecution returned error: %v", err)
	}
	if plan.SchemaVersion != KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion ||
		plan.RequestType != KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		plan.RuntimeMethod != "PrepareKnownAppKDERuntimeStatusLaunchExecution" ||
		plan.ExecutionMethod != "RunKnownAppKDERuntimeStatusLaunchExecution" ||
		plan.RequestPreviewType != KnownAppKDERuntimeStatusLaunchRequestType ||
		plan.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		plan.ActionKind != "runtime-status" ||
		plan.CenterCardState != "validated-post-review-dispatch" ||
		plan.PostReviewDispatchState != "created-after-session-gated-review" ||
		plan.LaunchAuthorizationReceiptID != launchReceiptID ||
		plan.SessionGatedReviewReceiptID != reviewReceiptID ||
		plan.ControlledExecutionSessionID != sessionID ||
		!plan.LaunchReceiptRevalidated ||
		!plan.GuestBoundaryRevalidated ||
		!plan.ReviewReceiptRevalidated ||
		!plan.ControlledSessionRevalidated ||
		!plan.ControlledSessionDigestVerified ||
		!plan.RuntimeManagedLauncherArgvReady ||
		strings.Join(plan.RuntimeManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --state-root <runtime-owned-state-root> --receipt-id "+launchReceiptID+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		plan.RequiredOpaqueIDCount != 3 ||
		plan.CollectedOpaqueIDCount != 3 ||
		!plan.StateRootRequired ||
		!plan.StateRootAccepted ||
		!plan.StateRootInjectedByRuntime ||
		!plan.StateRootSuppliedByRuntime ||
		plan.KDEStateRootAccess ||
		!plan.ManagedLauncherInvocationReady ||
		!plan.ExistingManagedLauncherPathUsed ||
		!plan.DelegatesArtifactGateToLauncher ||
		!plan.RuntimeOwnedRequest ||
		!plan.RuntimeOwnedLaunch ||
		!plan.RuntimeOwnedDispatch ||
		!plan.KDEPresentationOnly ||
		!plan.KDEActionForwarded ||
		plan.DirectLaunchEnabled ||
		plan.DesktopLaunchEnabled ||
		plan.BackendLaunchEnabled ||
		plan.ExecutionStarted ||
		plan.BackendProcessStarted ||
		plan.RequestObjectsCreated ||
		plan.PermissionGrantCreated ||
		plan.StateRootPathExposed ||
		plan.ReceiptPathExposed ||
		plan.SessionPathExposed ||
		plan.ManagedLauncherPathExposed ||
		plan.RawArtifactPathExposed ||
		plan.RawCommandExposed ||
		plan.RawLauncherOutputExposed ||
		plan.BackendDetailsExposed ||
		plan.HostRootModified ||
		plan.PrivilegedContainerRequired ||
		plan.DockerSocketMounted ||
		plan.BroadHostMountRequired {
		t.Fatalf("unexpected Runtime-status launch execution plan: %#v", plan)
	}
	argv, err := KnownAppKDERuntimeStatusLaunchExecutionArgv(plan, stateRoot, "", []string{"--timeout", "1s"})
	if err != nil {
		t.Fatalf("KnownAppKDERuntimeStatusLaunchExecutionArgv returned error: %v", err)
	}
	expected := []string{
		"--app", "7zr",
		"--cache-root", winapp.DefaultKnownAppCacheRoot,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", launchReceiptID,
		"--review-receipt-id", reviewReceiptID,
		"--session-id", sessionID,
		"--timeout", "1s",
	}
	if strings.Join(argv, "\x00") != strings.Join(expected, "\x00") {
		t.Fatalf("unexpected execution argv: %#v", argv)
	}
	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{strings.ToLower(stateRoot), ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Runtime-status launch execution plan exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTriggerConsumesHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppRuntimeStatusLaunchExecutionFixture(t, stateRoot)
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
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
	record, err := RecordKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	plan, err := PrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTrigger(KnownAppKDERuntimeStatusLaunchExecutionFromActionTriggerRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTrigger returned error: %v", err)
	}
	if plan.SchemaVersion != KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion ||
		plan.RequestType != KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		plan.Source != KnownAppKDERuntimeStatusLaunchActionTriggerRequestType+"+runtime-owner-execution-entrypoint" ||
		plan.ActionTriggerType != KnownAppKDERuntimeStatusLaunchActionTriggerRequestType ||
		plan.ActionTriggerRuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		plan.ActionTriggerReadMethod != "GetKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		plan.ActionTriggerState != "runtime-launch-request-assembled" ||
		!plan.EvidenceHandoffConsumed ||
		!plan.EvidenceDigestVerified ||
		plan.EvidenceRelativePath != record.EvidenceRelativePath ||
		plan.EvidenceSHA256 != record.EvidenceSHA256 ||
		plan.RequestPreviewType != KnownAppKDERuntimeStatusLaunchRequestType ||
		plan.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		plan.LaunchAuthorizationReceiptID != launchReceiptID ||
		plan.SessionGatedReviewReceiptID != reviewReceiptID ||
		plan.ControlledExecutionSessionID != sessionID ||
		!plan.LaunchReceiptRevalidated ||
		!plan.ReviewReceiptRevalidated ||
		!plan.ControlledSessionRevalidated ||
		!plan.RuntimeManagedLauncherArgvReady ||
		strings.Join(plan.RuntimeManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --state-root <runtime-owned-state-root> --receipt-id "+launchReceiptID+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		!plan.StateRootInjectedByRuntime ||
		!plan.StateRootSuppliedByRuntime ||
		plan.KDEStateRootAccess ||
		!plan.ManagedLauncherInvocationReady ||
		!plan.RuntimeOwnedRequest ||
		!plan.RuntimeOwnedLaunch ||
		!plan.RuntimeOwnedDispatch ||
		plan.DirectLaunchEnabled ||
		plan.DesktopLaunchEnabled ||
		plan.BackendLaunchEnabled ||
		plan.ExecutionStarted ||
		plan.StateRootPathExposed ||
		plan.RawLauncherOutputExposed ||
		plan.BackendDetailsExposed ||
		plan.HostRootModified {
		t.Fatalf("unexpected trigger-fed Runtime-status launch execution plan: %#v", plan)
	}
	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) {
		t.Fatalf("trigger-fed execution plan exposed state root: %s", string(encoded))
	}
}

func TestPrepareKnownAppKDERuntimeStatusLaunchExecutionRejectsMissingReviewReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	_, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	_, err = PrepareKnownAppKDERuntimeStatusLaunchExecution(KnownAppKDERuntimeStatusLaunchExecutionRequest{
		AppID:                        "7zr",
		StateRoot:                    stateRoot,
		LaunchAuthorizationReceiptID: KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		SessionGatedReviewReceiptID:  KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID),
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err == nil || !strings.Contains(err.Error(), "accepted session-gated review receipt") {
		t.Fatalf("expected missing review receipt to be rejected, got %v", err)
	}
}

func TestProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidenceBuildsCenterPayload(t *testing.T) {
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
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
	if projection.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		projection.RuntimeMethod != "ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence" ||
		projection.RequestType != "windows-known-app-dispatch-smoke" ||
		projection.Status != "passed" ||
		projection.AppID != "7zr" ||
		projection.DisplayName != "7-Zip standalone console executable" ||
		projection.AppVersion != "26.02" ||
		projection.GuestBoundary != "managed-known-app-guest-smoke" ||
		!projection.RuntimeOwnedDispatch ||
		!projection.ArtifactVerified ||
		!projection.MarkerObserved ||
		!projection.SessionGatedControlledDispatchConsumed ||
		projection.SessionGatedControlledDispatchState != "created-after-session-gated-review" ||
		projection.SessionGatedReviewReceiptID != reviewReceiptID ||
		projection.LaunchAuthorizationReceiptID != launchReceiptID ||
		projection.LaunchAuthorizationReceiptState != "recorded" ||
		projection.LaunchGateState != "controlled-dispatch-ready" ||
		!projection.LaunchGateConsumed ||
		!projection.LaunchGateReceiptAccepted ||
		!projection.LaunchGateGuestBoundaryAccepted ||
		!projection.ControlledDispatchReady ||
		!projection.ControlledExecutionSessionConsumed ||
		projection.ControlledExecutionSessionID != sessionID ||
		!projection.ControlledSessionDigestVerified ||
		projection.ControlledSessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		!projection.RuntimeOwnerConsumableSession ||
		!projection.KDEReadModelConsumableSession ||
		projection.StateRootPathExposed ||
		projection.ManagedLauncherPathExposed ||
		projection.RawLauncherOutputExposed ||
		projection.BackendDetailsExposed ||
		projection.HostRootModified ||
		projection.DockerSocketMounted ||
		projection.BroadHostMountRequired ||
		!projection.CompatibilityCenterProjectionReady ||
		!projection.KDECenterProjectionReady {
		t.Fatalf("unexpected delegated evidence projection: %#v", projection)
	}
	encoded, err := json.Marshal(projection)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	evidence, err := KnownAppSmokeEvidenceFromKDERuntimeStatusLaunchDelegatedEvidence(projection)
	if err != nil {
		t.Fatalf("KnownAppSmokeEvidenceFromKDERuntimeStatusLaunchDelegatedEvidence returned error: %v", err)
	}
	if evidence.AppID != "7zr" ||
		evidence.EvidenceSource != "staged-launcher-dispatch-smoke" ||
		evidence.SmokeStatus != "passed" ||
		evidence.LaunchAuthorizationReceiptState != "recorded" ||
		evidence.LaunchGateState != "controlled-dispatch-ready" ||
		!evidence.LaunchGateConsumed ||
		!evidence.LaunchGateReceiptAccepted ||
		!evidence.LaunchGateGuestBoundaryAccepted ||
		!evidence.ControlledDispatchReady ||
		evidence.ControlledExecutionSessionID != sessionID ||
		!evidence.LauncherSessionGateConsumed ||
		!evidence.LauncherSessionDigestVerified ||
		evidence.LauncherSessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		!evidence.LauncherSessionRuntimeOwnerConsumable ||
		!evidence.LauncherSessionKDEReadModelConsumable ||
		!evidence.PostReviewDispatchConsumed ||
		evidence.PostReviewDispatchState != "created-after-session-gated-review" ||
		evidence.SessionGatedReviewReceiptID != reviewReceiptID ||
		!evidence.MarkerObserved ||
		!evidence.ChecksumVerified {
		t.Fatalf("unexpected Center smoke evidence conversion: %#v", evidence)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("delegated evidence projection exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidenceRedactsBackendFailureReasons(t *testing.T) {
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
		AppVersion:                             "26.02",
		RequestType:                            "windows-known-app-dispatch-smoke",
		Status:                                 "failed",
		FailureReason:                          "guest wine runner returned a non-zero exit status",
		GuestBoundary:                          "managed-known-app-guest-smoke",
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       true,
		MarkerObserved:                         false,
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
	if projection.FailureReason != "managed dispatch failed" {
		t.Fatalf("projection did not redact backend failure reason: %#v", projection)
	}
	encoded, err := json.Marshal(projection)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{".exe", "wine ", "wine/", ".wine", "qemu-system", "program files"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("projection exposed backend term %q: %s", forbidden, text)
		}
	}
}

func TestRecordKnownAppKDERuntimeStatusLaunchEvidencePersistsSafeHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
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
	record, err := RecordKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion ||
		record.RequestType != KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType ||
		record.RuntimeMethod != "RecordKnownAppKDERuntimeStatusLaunchEvidence" ||
		record.ReadMethod != "GetKnownAppKDERuntimeStatusLaunchEvidence" ||
		record.EvidenceState != "persisted" ||
		record.EvidenceRelativePath == "" ||
		filepath.IsAbs(record.EvidenceRelativePath) ||
		record.EvidenceSHA256 == "" ||
		record.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		!record.CompatibilityCenterProjectionReady ||
		!record.KDECenterProjectionReady ||
		!record.KnownAppSmokeEvidenceReady ||
		record.LaunchAuthorizationReceiptID != launchReceiptID ||
		record.SessionGatedReviewReceiptID != reviewReceiptID ||
		record.ControlledExecutionSessionID != sessionID ||
		record.ControlledSessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		!record.RuntimeOwned ||
		!record.RuntimeOwnedDispatch ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		record.StateRootPathExposed ||
		record.EvidencePathExposed ||
		record.ManagedLauncherPathExposed ||
		record.RawLauncherOutputExposed ||
		record.BackendDetailsExposed ||
		record.HostRootModified ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired ||
		record.DesktopLaunchEnabled ||
		record.BackendLaunchEnabled ||
		record.ExecutionStarted ||
		record.BackendProcessStarted {
		t.Fatalf("unexpected Runtime-status launch evidence record: %#v", record)
	}
	content, err := os.ReadFile(filepath.Join(stateRoot, filepath.FromSlash(record.EvidenceRelativePath)))
	if err != nil {
		t.Fatalf("ReadFile persisted evidence returned error: %v", err)
	}
	if sha256Hex(string(content)) != record.EvidenceSHA256 {
		t.Fatalf("record digest mismatch: %s != %s", sha256Hex(string(content)), record.EvidenceSHA256)
	}
	if strings.Contains(string(content), stateRoot) {
		t.Fatalf("persisted evidence exposed state root: %s", string(content))
	}
}

func TestPreviewKnownAppKDERuntimeStatusLaunchEvidenceConsumesSafeHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
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
	record, err := RecordKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	preview, err := PreviewKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppKDERuntimeStatusLaunchEvidencePreviewSchemaVersion ||
		preview.RequestType != KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType ||
		preview.RuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchEvidence" ||
		preview.ReadMethod != "GetKnownAppKDERuntimeStatusLaunchEvidence" ||
		preview.EvidenceReadState != "consumed" ||
		!preview.EvidenceHandoffConsumed ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceDigestVerified ||
		preview.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		!preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.KnownAppSmokeEvidenceReady ||
		preview.KnownAppSmokeEvidence.CenterCardState != "validated-post-review-dispatch" ||
		preview.KnownAppSmokeEvidence.PrimaryActionID != "show-runtime-controlled-launch" ||
		preview.KnownAppSmokeEvidence.PrimaryActionKind != "runtime-status" ||
		preview.LaunchAuthorizationReceiptID != launchReceiptID ||
		preview.SessionGatedReviewReceiptID != reviewReceiptID ||
		preview.ControlledExecutionSessionID != sessionID ||
		!preview.RuntimeOwned ||
		!preview.RuntimeOwnedDispatch ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.StateRootPathExposed ||
		preview.EvidencePathExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted {
		t.Fatalf("unexpected Runtime-status launch evidence preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) {
		t.Fatalf("preview exposed state root: %s", string(encoded))
	}
}

func writeKnownAppRuntimeStatusLaunchExecutionFixture(t *testing.T, stateRoot string) string {
	t.Helper()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	if _, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	}); err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	if _, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	}); err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	return sessionID
}
