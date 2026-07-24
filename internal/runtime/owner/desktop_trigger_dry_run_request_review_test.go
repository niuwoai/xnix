package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewDesktopTriggerDryRunRequestReviewAcceptsSafeInputs(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerDryRunRequestReview returned error: %v", err)
	}
	if preview.SchemaVersion != DesktopTriggerDryRunRequestReviewSchemaVersion ||
		preview.RequestType != DesktopTriggerDryRunRequestReviewRequestType ||
		preview.ReviewState != "accepted-review" ||
		preview.EnvelopeGuardState != "accepted-for-review" ||
		preview.ActionSurfaceState != "safe" ||
		preview.ManagedLauncherAcceptanceState != "accepted" ||
		preview.FullCheckpointState != "ready" ||
		preview.RuntimeStatusEvidenceState != "ready" ||
		preview.EvidenceHandleKind != "evidence-relative-path" ||
		preview.EvidenceHandle != record.EvidenceRelativePath ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceDigestVerified ||
		!preview.ExpectedDigestMatched ||
		preview.RouteID != launchEnvelopeExpectedRoute ||
		preview.MethodID != launchEnvelopeExpectedMethod ||
		preview.ActionID != launchEnvelopeExpectedActionID ||
		preview.CallerRole != launchEnvelopeExpectedCallerRole ||
		preview.KDEActionID != launchEnvelopeExpectedActionID ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.ForwardedArgumentKind != "evidence-relative-path" ||
		!preview.EnvelopeAcceptedForReview ||
		!preview.ActionSurfaceSafeForHumanSmoke ||
		!preview.ManagedLauncherAccepted ||
		!preview.DryRunReviewOnly ||
		preview.RequestObjectWritten ||
		preview.PermissionGrantCreated ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		!preview.KDEPresentationOnly ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		preview.StateRootPathExposed ||
		preview.CacheRootPathExposed ||
		preview.LauncherPathExposed ||
		preview.RawExecutablePathExposed ||
		preview.RawCommandExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.ServiceCallDispatchEnabled ||
		preview.ServiceCallDispatched ||
		preview.DBusCalled ||
		preview.DBusOwnershipEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RuntimeStateWritten ||
		preview.KDEConfigurationWritten ||
		preview.SmokeExecutedByPreview ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.ProductionAuthorizationAccepted ||
		preview.NextHumanAuthorizedSmoke != "desktop-triggered-staged-launch-smoke" {
		t.Fatalf("unexpected accepted dry-run review: %#v", preview)
	}
	assertDesktopTriggerDryRunRedacted(t, preview, stateRoot)
}

func TestPreviewDesktopTriggerDryRunRequestReviewBlocksMissingFullCheckpoint(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerDryRunRequestReview returned error: %v", err)
	}
	if preview.ReviewState != "blocked-missing-full-checkpoint" ||
		preview.FullCheckpointState != "needs-full-checkpoint" ||
		preview.EnvelopeGuardState != "accepted-for-review" ||
		preview.ActionSurfaceState != "safe" ||
		preview.ManagedLauncherAcceptanceState != "needs-full-checkpoint" ||
		preview.ServiceCallDispatchEnabled ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified {
		t.Fatalf("unexpected missing full checkpoint review: %#v", preview)
	}
}

func TestPreviewDesktopTriggerDryRunRequestReviewBlocksUnsafeEnvelope(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
		KDEStateRoot:           "/private/runtime",
		KDEBackendCommand:      "private engine command",
		KDEReceiptID:           "receipt-1",
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerDryRunRequestReview returned error: %v", err)
	}
	if preview.ReviewState != "blocked-unsafe-envelope" ||
		preview.EnvelopeGuardState != "blocked-owner-args" ||
		preview.ActionSurfaceState != "safe" ||
		preview.ServiceCallDispatched ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified {
		t.Fatalf("unexpected unsafe envelope review: %#v", preview)
	}
	assertDesktopTriggerDryRunRedacted(t, preview, "/private/runtime")
	assertDesktopTriggerDryRunRedacted(t, preview, "private engine command")
}

func TestPreviewDesktopTriggerDryRunRequestReviewBlocksUnsafeActionSurface(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	unsafeAction := strings.Replace(safeDesktopTriggerDryRunActionMetadata(), "X-Xnix-State-Root-Access=false", "X-Xnix-State-Root-Access=true", 1)

	preview, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    unsafeAction,
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerDryRunRequestReview returned error: %v", err)
	}
	if preview.ReviewState != "blocked-unsafe-action" ||
		preview.EnvelopeGuardState != "accepted-for-review" ||
		preview.ActionSurfaceState != "unsafe-owner-args" ||
		preview.ActionSurfaceSafeForHumanSmoke ||
		preview.ServiceCallDispatched ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified {
		t.Fatalf("unexpected unsafe action review: %#v", preview)
	}
}

func TestPreviewDesktopTriggerDryRunRequestReviewClassifiesMalformedAndStaleEvidence(t *testing.T) {
	malformed, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              t.TempDir(),
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   "../escape.json",
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("malformed review should be a classified packet, got error: %v", err)
	}
	if malformed.ReviewState != "malformed" || malformed.ServiceCallDispatchEnabled || malformed.HostRootModified {
		t.Fatalf("unexpected malformed review: %#v", malformed)
	}

	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	stale, err := PreviewDesktopTriggerDryRunRequestReview(DesktopTriggerDryRunRequestReviewRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: strings.Repeat("0", 64),
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("stale review should be a classified packet, got error: %v", err)
	}
	if stale.ReviewState != "stale-evidence" || stale.ExpectedDigestMatched || stale.ServiceCallDispatched || stale.HostRootModified {
		t.Fatalf("unexpected stale review: %#v", stale)
	}
}

func assertDesktopTriggerDryRunRedacted(t *testing.T, preview DesktopTriggerDryRunRequestReviewPreview, forbidden string) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	for _, unsafe := range []string{forbidden, "owner_service_call_args", "owner_service_cli_args", " --service-call ", ".exe", "SECRET_TOKEN", "USER="} {
		if unsafe != "" && strings.Contains(string(encoded), unsafe) {
			t.Fatalf("dry-run review exposed unsafe detail %q: %s", unsafe, string(encoded))
		}
	}
}

func safeDesktopTriggerDryRunActionMetadata() string {
	return `[Desktop Entry]
Type=Service
Name=Xnix Runtime controlled launch
Comment=Forward a Runtime-status evidence handle to the Xnix Runtime D-Bus controlled-launch action.
Icon=media-playback-start
X-Xnix-KDE-Action-ID=xnix.runtime-status.controlled-launch
X-Xnix-Runtime-Preview=xnix-runtime-go kde-controlled-launch-action-preview
X-Xnix-Restricted-Smoke-Plan=xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview
X-Xnix-DBus-Service=org.xnix.Compatibility1
X-Xnix-DBus-Object-Path=/org/xnix/Compatibility1
X-Xnix-DBus-Method=org.xnix.Compatibility1.ShowRuntimeControlledLaunch
X-Xnix-Forwarded-Argument=evidence-relative-path
X-Xnix-Forwards-Only-Evidence-Handle=true
X-Xnix-KDE-Policy-Owner=false
X-Xnix-Owner-Service-Args-Exposed-To-KDE=false
X-Xnix-State-Root-Access=false
X-Xnix-Receipt-Reconstruction=false
X-Xnix-Backend-Launch-Enabled=false
X-Xnix-Execution-Started=false
X-Xnix-Host-Root-Modified=false
X-Xnix-Docker-Socket-Mounted=false
X-Xnix-Privileged-Container-Required=false
X-Xnix-Host-Network-Required=false
`
}
