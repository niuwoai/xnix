package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewDesktopTriggerServiceCallMaterializationEmitsEvidenceOnlyServiceCall(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerServiceCallMaterialization(DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerServiceCallMaterialization returned error: %v", err)
	}
	if preview.SchemaVersion != DesktopTriggerServiceCallMaterializationSchemaVersion ||
		preview.RequestType != DesktopTriggerServiceCallMaterializationRequestType ||
		preview.MaterializationState != "ready-for-human-authorized-service-call" ||
		preview.DryRunReviewState != "accepted-review" ||
		preview.OwnerTriggerState != "ready" ||
		preview.FullCheckpointState != "ready" ||
		preview.RuntimeStatusEvidenceState != "ready" ||
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
		!preview.OwnerServiceCallReady ||
		!preview.HumanAuthorizationRequired ||
		!preview.MaterializedForHumanSmoke ||
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
		t.Fatalf("unexpected materialized service call: %#v", preview)
	}
	if !sameDesktopTriggerServiceCallMaterializationArgs(preview.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", record.EvidenceRelativePath}) ||
		!sameDesktopTriggerServiceCallMaterializationArgs(preview.OwnerServiceCLIArgs, []string{"--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", record.EvidenceRelativePath}) {
		t.Fatalf("unexpected materialized service call args: %#v", preview)
	}
	assertDesktopTriggerServiceCallMaterializationRedacted(t, preview, stateRoot)
}

func TestPreviewDesktopTriggerServiceCallMaterializationBlocksUntilDryRunAccepted(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerServiceCallMaterialization(DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerServiceCallMaterialization returned error: %v", err)
	}
	if preview.MaterializationState != "blocked-missing-full-checkpoint" ||
		preview.DryRunReviewState != "blocked-missing-full-checkpoint" ||
		preview.OwnerTriggerState != "blocked" ||
		preview.OwnerServiceCallReady ||
		preview.MaterializedForHumanSmoke ||
		len(preview.OwnerServiceCallArgs) != 0 ||
		len(preview.OwnerServiceCLIArgs) != 0 ||
		preview.ServiceCallDispatched ||
		preview.RuntimeStateWritten ||
		preview.HostRootModified {
		t.Fatalf("unexpected blocked materialization: %#v", preview)
	}
}

func TestPreviewDesktopTriggerServiceCallMaterializationBlocksUnsafeEnvelopeWithoutArgs(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewDesktopTriggerServiceCallMaterialization(DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		FullCheckpointPromoted: true,
		KDEStateRoot:           "/private/runtime",
		KDEBackendCommand:      "private engine command",
	})
	if err != nil {
		t.Fatalf("PreviewDesktopTriggerServiceCallMaterialization returned error: %v", err)
	}
	if preview.MaterializationState != "blocked-unsafe-envelope" ||
		preview.DryRunReviewState != "blocked-unsafe-envelope" ||
		preview.OwnerServiceCallReady ||
		len(preview.OwnerServiceCallArgs) != 0 ||
		len(preview.OwnerServiceCLIArgs) != 0 ||
		preview.ServiceCallDispatched ||
		preview.HostRootModified {
		t.Fatalf("unexpected unsafe-envelope materialization: %#v", preview)
	}
	assertDesktopTriggerServiceCallMaterializationRedacted(t, preview, "/private/runtime")
	assertDesktopTriggerServiceCallMaterializationRedacted(t, preview, "private engine command")
}

func TestPreviewDesktopTriggerServiceCallMaterializationClassifiesMalformedAndStaleEvidence(t *testing.T) {
	malformed, err := PreviewDesktopTriggerServiceCallMaterialization(DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              t.TempDir(),
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   "../escape.json",
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("malformed materialization should be classified, got error: %v", err)
	}
	if malformed.MaterializationState != "malformed" || len(malformed.OwnerServiceCallArgs) != 0 || malformed.ServiceCallDispatched {
		t.Fatalf("unexpected malformed materialization: %#v", malformed)
	}

	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	stale, err := PreviewDesktopTriggerServiceCallMaterialization(DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              stateRoot,
		DesktopEntryContent:    safeDesktopTriggerDryRunActionMetadata(),
		EvidenceRelativePath:   record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: strings.Repeat("0", 64),
		FullCheckpointPromoted: true,
	})
	if err != nil {
		t.Fatalf("stale materialization should be classified, got error: %v", err)
	}
	if stale.MaterializationState != "stale-evidence" ||
		stale.ExpectedDigestMatched ||
		len(stale.OwnerServiceCallArgs) != 0 ||
		stale.ServiceCallDispatched ||
		stale.HostRootModified {
		t.Fatalf("unexpected stale materialization: %#v", stale)
	}
}

func assertDesktopTriggerServiceCallMaterializationRedacted(t *testing.T, preview DesktopTriggerServiceCallMaterializationPreview, forbidden string) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	for _, unsafe := range []string{forbidden, "XNIX_RUNTIME_OWNER_", " --state-root ", " --cache-root ", " --launcher ", ".exe", "SECRET_TOKEN", "USER="} {
		if unsafe != "" && strings.Contains(string(encoded), unsafe) {
			t.Fatalf("service call materialization exposed unsafe detail %q: %s", unsafe, string(encoded))
		}
	}
}
