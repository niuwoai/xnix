package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreviewManagedLauncherAcceptanceReportNeedsFullCheckpoint(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")

	preview, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewManagedLauncherAcceptanceReport returned error: %v", err)
	}
	if preview.SchemaVersion != ManagedLauncherAcceptanceReportSchemaVersion ||
		preview.RequestType != ManagedLauncherAcceptanceReportRequestType ||
		preview.AcceptanceState != "needs-full-checkpoint" ||
		preview.KnownAppIdentityState != "ready" ||
		preview.ArtifactDigestState != "ready" ||
		preview.LaunchAuthorizationState != "ready" ||
		preview.SessionGatedReviewState != "ready" ||
		preview.ControlledExecutionSessionState != "ready" ||
		preview.ManagedLauncherRequestState != "ready" ||
		preview.GuestSmokeBoundaryState != "ready" ||
		preview.RuntimeStatusEvidenceState != "ready" ||
		preview.KDESafeProjectionState != "ready" ||
		preview.FullCheckpointState != "needs-full-checkpoint" ||
		preview.DesktopTriggerReadinessState != "needs-full-checkpoint" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceDigestVerified ||
		!preview.ExpectedDigestMatched ||
		!preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.ManagedLauncherArgvReady ||
		preview.KnownAppAcceptanceReady ||
		preview.KnownAppSmokeStatus != "passed" ||
		!preview.KnownAppArtifactVerified ||
		!preview.KnownAppMarkerObserved ||
		preview.GuestBoundary != "managed-known-app-guest-smoke" ||
		preview.LaunchAuthorizationReceiptID == "" ||
		preview.SessionGatedReviewReceiptID == "" ||
		preview.ControlledExecutionSessionID == "" ||
		!preview.RuntimeOwned ||
		!preview.RuntimeOwnedDispatch ||
		!preview.GoRuntimeBacked ||
		!preview.KDEPresentationOnly ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
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
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.FormalReleaseReady ||
		len(preview.Sections) != 11 {
		t.Fatalf("unexpected managed launcher acceptance report: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) ||
		strings.Contains(string(encoded), "owner_service_call_args") ||
		strings.Contains(string(encoded), "owner_service_cli_args") ||
		strings.Contains(string(encoded), " --service-call ") ||
		strings.Contains(string(encoded), ".exe") ||
		strings.Contains(string(encoded), "SECRET_TOKEN") ||
		strings.Contains(string(encoded), "USER=") {
		t.Fatalf("acceptance report exposed unsafe details: %s", string(encoded))
	}
}

func TestPreviewManagedLauncherAcceptanceReportAcceptedAfterFullCheckpoint(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")

	preview, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:              stateRoot,
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewManagedLauncherAcceptanceReport returned error: %v", err)
	}
	if preview.AcceptanceState != "accepted" ||
		preview.FullCheckpointState != "ready" ||
		preview.DesktopTriggerReadinessState != "ready" ||
		!preview.KnownAppAcceptanceReady ||
		!preview.FormalReleaseReady ||
		preview.NextHumanAuthorizedSmoke != "desktop-triggered-staged-launch-smoke" ||
		strings.Join(preview.NextHumanAuthorizedSmokeCommand, " ") != "ruby scripts/staged_launcher_dispatch_smoke.rb" {
		t.Fatalf("unexpected accepted managed launcher report: %#v", preview)
	}
}

func TestPreviewManagedLauncherAcceptanceReportClassifiesMissingMalformedAndStaleEvidence(t *testing.T) {
	missing, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot: t.TempDir(),
	})
	if err != nil {
		t.Fatalf("missing evidence report should be a classified packet, got error: %v", err)
	}
	if missing.AcceptanceState != "missing-evidence" || missing.RuntimeStatusEvidenceState != "missing-evidence" || missing.FormalReleaseReady {
		t.Fatalf("unexpected missing evidence packet: %#v", missing)
	}

	stateRoot := t.TempDir()
	record := recordDesktopTriggerReadinessEvidence(t, stateRoot, true, true, "passed")
	evidencePath := filepath.Join(stateRoot, filepath.FromSlash(record.EvidenceRelativePath))
	if err := os.WriteFile(evidencePath, []byte("{"), 0o600); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
	malformed, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("malformed evidence report should be a classified packet, got error: %v", err)
	}
	if malformed.AcceptanceState != "malformed" || malformed.RuntimeStatusEvidenceState != "malformed" || malformed.FormalReleaseReady {
		t.Fatalf("unexpected malformed evidence packet: %#v", malformed)
	}

	staleRoot := t.TempDir()
	staleRecord := recordDesktopTriggerReadinessEvidence(t, staleRoot, true, true, "passed")
	stale, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:              staleRoot,
		EvidenceRelativePath:   staleRecord.EvidenceRelativePath,
		ExpectedEvidenceSHA256: strings.Repeat("0", 64),
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("stale evidence report should be a classified packet, got error: %v", err)
	}
	if stale.AcceptanceState != "stale-evidence" || stale.ExpectedDigestMatched || stale.FormalReleaseReady {
		t.Fatalf("unexpected stale evidence packet: %#v", stale)
	}
}

func TestPreviewManagedLauncherAcceptanceReportBlocksMissingArtifactReviewAndGuestSmoke(t *testing.T) {
	missingArtifactRoot := t.TempDir()
	missingArtifactRecord := recordDesktopTriggerReadinessEvidence(t, missingArtifactRoot, false, true, "passed")
	missingArtifact, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:              missingArtifactRoot,
		EvidenceRelativePath:   missingArtifactRecord.EvidenceRelativePath,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("missing artifact report should be a classified packet, got error: %v", err)
	}
	if missingArtifact.AcceptanceState != "missing-evidence" ||
		missingArtifact.ArtifactDigestState != "missing-evidence" ||
		missingArtifact.FormalReleaseReady {
		t.Fatalf("unexpected missing artifact packet: %#v", missingArtifact)
	}

	missingReviewRoot := t.TempDir()
	missingReviewRecord := recordManagedLauncherAcceptanceEvidence(t, missingReviewRoot, true, true, "passed", false)
	missingReview, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:              missingReviewRoot,
		EvidenceRelativePath:   missingReviewRecord.EvidenceRelativePath,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("missing review report should be a classified packet, got error: %v", err)
	}
	if missingReview.AcceptanceState != "blocked" ||
		missingReview.SessionGatedReviewState != "blocked" ||
		missingReview.ManagedLauncherRequestState != "blocked" ||
		missingReview.FormalReleaseReady {
		t.Fatalf("unexpected missing review packet: %#v", missingReview)
	}

	failedGuestRoot := t.TempDir()
	failedGuestRecord := recordDesktopTriggerReadinessEvidence(t, failedGuestRoot, true, false, "failed")
	failedGuest, err := PreviewManagedLauncherAcceptanceReport(ManagedLauncherAcceptanceReportRequest{
		StateRoot:              failedGuestRoot,
		EvidenceRelativePath:   failedGuestRecord.EvidenceRelativePath,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("failed guest report should be a classified packet, got error: %v", err)
	}
	if failedGuest.AcceptanceState != "blocked" ||
		failedGuest.GuestSmokeBoundaryState != "blocked" ||
		failedGuest.FormalReleaseReady {
		t.Fatalf("unexpected failed guest packet: %#v", failedGuest)
	}
}

func recordManagedLauncherAcceptanceEvidence(t *testing.T, stateRoot string, artifactVerified bool, markerObserved bool, status string, includePostReview bool) KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	postReviewState := ""
	if includePostReview {
		postReviewState = "created-after-session-gated-review"
	}
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
		SessionGatedControlledDispatchConsumed: includePostReview,
		SessionGatedControlledDispatchState:    postReviewState,
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
