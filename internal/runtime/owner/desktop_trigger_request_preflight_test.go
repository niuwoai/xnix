package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewDesktopTriggerRequestPreflightReadyForOperatorRequest(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerRequestPreflight(DesktopTriggerRequestPreflightRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerRequestPreflight returned error: %v", err)
	}
	if preview.SchemaVersion != DesktopTriggerRequestPreflightSchemaVersion ||
		preview.RequestType != DesktopTriggerRequestPreflightRequestType ||
		preview.PreflightState != "ready-for-operator-request" ||
		preview.MaterializationState != "ready-for-human-authorized-service-call" ||
		preview.DryRunReviewState != "accepted-review" ||
		preview.OwnerTriggerState != "ready" ||
		preview.FullCheckpointState != "ready" ||
		preview.RuntimeStatusEvidenceState != "ready" ||
		preview.EvidenceHandleKind != "evidence-relative-path" ||
		preview.EvidenceHandle != record.EvidenceRelativePath ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceDigestVerified ||
		!preview.ExpectedDigestMatched ||
		preview.DesktopCallableRoute != "kde-dbus-runtime-status-action" ||
		preview.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		preview.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.OwnerServiceBoundary != "go-runtime-owner-in-process-service" ||
		preview.OwnerServiceMethod != "ShowRuntimeControlledLaunch" ||
		preview.OwnerServiceCallType != "desktop-action-dispatch" ||
		!preview.OwnerServiceCallShapeVerified ||
		!preview.OwnerServiceCallReady ||
		!preview.OperatorRequestReady ||
		!preview.FormalPromotionRequired ||
		!preview.FormalPromotionObserved ||
		preview.FormalReleaseReady ||
		!preview.HumanAuthorizationRequired ||
		!preview.RuntimeOwnerServiceSuppliesInputs ||
		!preview.DesktopEvidenceHandleForwarded ||
		!preview.RuntimeOwned ||
		!preview.RuntimeOwnedDispatch ||
		!preview.GoRuntimeBacked ||
		!preview.KDEPresentationOnly ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		preview.KDEReceivesMaterializedOwnerArgs ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.StateRootPathExposed ||
		preview.CacheRootPathExposed ||
		preview.LauncherPathExposed ||
		preview.RawExecutablePathExposed ||
		preview.RawCommandExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.RequestObjectWritten ||
		preview.PermissionGrantCreated ||
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
		preview.ProductionAuthorizationAccepted {
		t.Fatalf("unexpected desktop-trigger request preflight: %#v", preview)
	}
	assertDesktopTriggerRequestPreflightRedacted(t, preview, stateRoot)
}

func TestPreviewDesktopTriggerRequestPreflightBlocksMissingPromotion(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerRequestPreflight(DesktopTriggerRequestPreflightRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerRequestPreflight returned error: %v", err)
	}
	if preview.PreflightState != "blocked-missing-promotion" ||
		preview.MaterializationState != "blocked-missing-full-checkpoint" ||
		preview.FullCheckpointState != "needs-full-checkpoint" ||
		preview.OperatorRequestReady ||
		preview.OwnerServiceCallReady ||
		preview.OwnerServiceCallShapeVerified ||
		preview.FormalPromotionObserved ||
		preview.FormalReleaseReady ||
		preview.ServiceCallDispatched ||
		preview.DBusCalled ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified {
		t.Fatalf("unexpected missing-promotion preflight: %#v", preview)
	}
	assertDesktopTriggerRequestPreflightRedacted(t, preview, stateRoot)
}

func TestPreviewDesktopTriggerRequestPreflightBlocksUnsafeEnvelope(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerRequestPreflight(DesktopTriggerRequestPreflightRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
		KDEStateRoot:           "/private/runtime",
		KDEBackendCommand:      "private engine command",
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerRequestPreflight returned error: %v", err)
	}
	if preview.PreflightState != "blocked-unsafe-envelope" ||
		preview.MaterializationState != "blocked-unsafe-envelope" ||
		preview.OwnerServiceCallReady ||
		preview.OwnerServiceCallShapeVerified ||
		preview.OperatorRequestReady ||
		preview.ServiceCallDispatched ||
		preview.HostRootModified {
		t.Fatalf("unexpected unsafe-envelope preflight: %#v", preview)
	}
	assertDesktopTriggerRequestPreflightRedacted(t, preview, "/private/runtime")
	assertDesktopTriggerRequestPreflightRedacted(t, preview, "private engine command")
}

func TestPreviewDesktopTriggerRequestPreflightClassifiesStaleAndMalformedEvidence(t *testing.T) {
	malformed, err := PreviewDesktopTriggerRequestPreflight(DesktopTriggerRequestPreflightRequest{
		StateRoot:              t.TempDir(),
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   "../escape.json",
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("malformed preflight should be classified, got error: %v", err)
	}
	if malformed.PreflightState != "malformed" ||
		malformed.OwnerServiceCallReady ||
		malformed.OwnerServiceCallShapeVerified ||
		malformed.OperatorRequestReady ||
		malformed.ServiceCallDispatched ||
		malformed.HostRootModified {
		t.Fatalf("unexpected malformed preflight: %#v", malformed)
	}

	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	stale, err := PreviewDesktopTriggerRequestPreflight(DesktopTriggerRequestPreflightRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: strings.Repeat("0", 64),
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("stale preflight should be classified, got error: %v", err)
	}
	if stale.PreflightState != "blocked-stale-evidence" ||
		stale.ExpectedDigestMatched ||
		stale.OwnerServiceCallReady ||
		stale.OwnerServiceCallShapeVerified ||
		stale.OperatorRequestReady ||
		stale.ServiceCallDispatched ||
		stale.HostRootModified {
		t.Fatalf("unexpected stale preflight: %#v", stale)
	}
	assertDesktopTriggerRequestPreflightRedacted(t, stale, stateRoot)
}

func assertDesktopTriggerRequestPreflightRedacted(t *testing.T, preview DesktopTriggerRequestPreflightPreview, forbidden string) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	for _, unsafe := range []string{
		forbidden,
		"owner_service_call_args",
		"owner_service_cli_args",
		"XNIX_RUNTIME_OWNER_",
		" --service-call ",
		" --state-root ",
		" --cache-root ",
		" --launcher ",
		".exe",
		"SECRET_TOKEN",
		"USER=",
	} {
		if unsafe != "" && strings.Contains(string(encoded), unsafe) {
			t.Fatalf("desktop-trigger request preflight exposed unsafe detail %q: %s", unsafe, string(encoded))
		}
	}
}
