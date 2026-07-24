package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreviewDesktopTriggerStagedInvocationReadinessNeedsFullCheckpoint(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")

	preview, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerStagedInvocationReadiness returned error: %v", err)
	}
	if preview.SchemaVersion != DesktopTriggerStagedInvocationReadinessSchemaVersion ||
		preview.RequestType != DesktopTriggerStagedInvocationReadinessRequestType ||
		preview.ReadinessState != "needs-full-checkpoint" ||
		preview.ReleaseGateState != "needs-full-checkpoint" ||
		preview.KDEActionState != "ready" ||
		preview.PublicDBusRouteState != "ready" ||
		preview.OwnerServiceTriggerState != "ready" ||
		preview.RuntimeStatusEvidenceState != "ready" ||
		preview.ManagedLauncherRequestState != "ready" ||
		preview.KnownAppArtifactState != "ready" ||
		preview.GuestSmokeBoundaryState != "ready" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceDigestVerified ||
		!preview.ExpectedDigestMatched ||
		preview.KDEActionID != "xnix.runtime-status.controlled-launch" ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		!preview.ManagedLauncherArgvReady ||
		preview.KnownAppSmokeStatus != "passed" ||
		!preview.KnownAppArtifactVerified ||
		!preview.KnownAppMarkerObserved ||
		preview.GuestBoundary != "managed-known-app-guest-smoke" ||
		!preview.RuntimeOwned ||
		!preview.RuntimeOwnedDispatch ||
		!preview.GoRuntimeBacked ||
		!preview.KDEPresentationOnly ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		preview.OwnerServiceArgsExposedToKDE ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.StateRootPathExposed ||
		preview.EvidencePathExposed ||
		preview.ManagedLauncherPathExposed ||
		preview.RawArtifactPathExposed ||
		preview.RawCommandExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RuntimeStateWritten ||
		preview.KDEConfigurationWritten ||
		preview.DBusCalled ||
		preview.SmokeExecutedByPreview ||
		preview.NetworkFetchRequired ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.FormalReleaseReady ||
		len(preview.Sections) != 8 {
		t.Fatalf("unexpected desktop-trigger readiness preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) ||
		strings.Contains(string(encoded), "owner_service_call_args") ||
		strings.Contains(string(encoded), " --service-call ") ||
		strings.Contains(string(encoded), ".exe") {
		t.Fatalf("readiness preview exposed unsafe details: %s", string(encoded))
	}
}

func TestPreviewDesktopTriggerStagedInvocationReadinessReadyAfterFullCheckpoint(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")

	preview, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:              stateRoot,
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerStagedInvocationReadiness returned error: %v", err)
	}
	if preview.ReadinessState != "ready" ||
		preview.ReleaseGateState != "ready" ||
		!preview.FormalReleaseReady ||
		preview.NextHumanAuthorizedSmoke != "desktop-triggered-staged-launch-smoke" ||
		strings.Join(preview.NextHumanAuthorizedSmokeCommand, " ") != "ruby scripts/staged_launcher_dispatch_smoke.rb" {
		t.Fatalf("unexpected ready desktop-trigger preview: %#v", preview)
	}
}

func TestPreviewDesktopTriggerStagedInvocationReadinessClassifiesMissingAndMalformedEvidence(t *testing.T) {
	missing, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot: t.TempDir(),
	})
	if err != nil {
		t.Fatalf("missing evidence preview should be a classified packet, got error: %v", err)
	}
	if missing.ReadinessState != "missing-evidence" || missing.RuntimeStatusEvidenceState != "missing-evidence" || missing.FormalReleaseReady {
		t.Fatalf("unexpected missing evidence packet: %#v", missing)
	}

	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")
	evidencePath := filepath.Join(stateRoot, filepath.FromSlash(record.EvidenceRelativePath))
	if err := os.WriteFile(evidencePath, []byte("{"), 0o600); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
	malformed, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("malformed evidence preview should be a classified packet, got error: %v", err)
	}
	if malformed.ReadinessState != "malformed" || malformed.RuntimeStatusEvidenceState != "malformed" || malformed.FormalReleaseReady {
		t.Fatalf("unexpected malformed evidence packet: %#v", malformed)
	}
}

func TestPreviewDesktopTriggerStagedInvocationReadinessClassifiesStaleDigestAndMissingArtifact(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")
	stale, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:              stateRoot,
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: strings.Repeat("0", 64),
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("stale digest preview should be a classified packet, got error: %v", err)
	}
	if stale.ReadinessState != "stale-evidence" || stale.ExpectedDigestMatched || stale.FormalReleaseReady {
		t.Fatalf("unexpected stale digest packet: %#v", stale)
	}

	missingArtifactRoot := t.TempDir()
	missingArtifactRecord := recordDesktopTriggerReadinessEvidence(t, missingArtifactRoot, false, true, "passed")
	missingArtifact, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:              missingArtifactRoot,
		EvidenceRelativePath:   missingArtifactRecord.EvidenceRelativePath,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("missing artifact preview should be a classified packet, got error: %v", err)
	}
	if missingArtifact.ReadinessState != "missing-evidence" ||
		missingArtifact.KnownAppArtifactState != "missing-evidence" ||
		missingArtifact.FormalReleaseReady {
		t.Fatalf("unexpected missing artifact packet: %#v", missingArtifact)
	}
}

func TestPreviewDesktopTriggerStagedInvocationReadinessBlocksFailedGuestSmoke(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, false, "failed")
	preview, err := PreviewDesktopTriggerStagedInvocationReadiness(DesktopTriggerStagedInvocationReadinessRequest{
		StateRoot:              stateRoot,
		EvidenceRelativePath:   record.EvidenceRelativePath,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerStagedInvocationReadiness returned error: %v", err)
	}
	if preview.ReadinessState != "blocked" || preview.GuestSmokeBoundaryState != "blocked" || preview.FormalReleaseReady {
		t.Fatalf("unexpected failed guest smoke packet: %#v", preview)
	}
}

func recordDesktopTriggerReadinessEvidence(t *testing.T, stateRoot string, artifactVerified bool, markerObserved bool, status string) KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
		AppVersion:                             "26.02",
		RequestType:                            "windows-known-app-dispatch-smoke",
		Status:                                 status,
		GuestBoundary:                          "managed-known-app-guest-smoke",
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       artifactVerified,
		MarkerObserved:                         markerObserved,
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
	return record
}
