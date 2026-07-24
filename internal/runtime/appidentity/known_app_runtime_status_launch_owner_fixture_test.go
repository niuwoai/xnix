package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRecordKnownAppRuntimeStatusLaunchOwnerFixtureBlocksWithoutVerifiedArtifact(t *testing.T) {
	record, err := RecordKnownAppRuntimeStatusLaunchOwnerFixture(KnownAppRuntimeStatusLaunchOwnerFixtureRequest{
		AppID:     "7zr",
		StateRoot: t.TempDir(),
		CacheRoot: t.TempDir(),
	})
	if err != nil {
		t.Fatalf("RecordKnownAppRuntimeStatusLaunchOwnerFixture returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppRuntimeStatusLaunchOwnerFixtureSchemaVersion ||
		record.RequestType != KnownAppRuntimeStatusLaunchOwnerFixtureRequestType ||
		record.RuntimeMethod != "RecordKnownAppRuntimeStatusLaunchOwnerFixture" ||
		record.ReadMethod != "GetKnownAppRuntimeStatusLaunchOwnerFixture" {
		t.Fatalf("unexpected fixture identity: %#v", record)
	}
	if record.FixtureReady ||
		record.FixtureState != "blocked" ||
		record.SkipReason == "" ||
		record.LaunchAuthorizationReceiptID != "known-app-launch-authorization-7zr-26.02" ||
		record.LaunchAuthorizationReceiptRecorded != true {
		t.Fatalf("blocked fixture did not preserve safe receipt evidence: %#v", record)
	}
	if record.DesktopTriggerReady ||
		record.OwnerServiceCallReady ||
		record.DesktopEvidenceHandleForwarded ||
		record.RuntimeOwnerServiceSuppliesInputs ||
		len(record.OwnerServiceCallArgs) != 0 {
		t.Fatalf("blocked fixture must not expose a ready desktop trigger: %#v", record)
	}
	if record.StateRootPathExposed ||
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
		t.Fatalf("blocked fixture opened unsafe gates: %#v", record)
	}
	if !record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		!record.KDEForwardsOnlyEvidenceHandle ||
		record.DesktopReceiptFieldsReconstructed ||
		record.DesktopKDEStateRootAccess {
		t.Fatalf("blocked fixture must keep Runtime ownership and KDE evidence-only handoff: %#v", record)
	}
}

func TestRecordKnownAppRuntimeStatusLaunchOwnerFixtureConsumesGUISmokeEvidenceForMines(t *testing.T) {
	version := currentProjectVersion(t)
	guiPreview, err := PreviewGUISmokeEvidenceJSON([]byte(guiSmokeEvidenceFixture(true)), GUISmokeEvidencePreviewRequest{
		AppID:       "org.xnix.apps.mines",
		DisplayName: "Mines",
		AppVersion:  version,
	})
	if err != nil {
		t.Fatalf("PreviewGUISmokeEvidenceJSON returned error: %v", err)
	}
	evidencePayload, err := json.Marshal(guiPreview)
	if err != nil {
		t.Fatalf("Marshal GUI evidence returned error: %v", err)
	}
	evidencePath := filepath.Join(t.TempDir(), "mines-gui-evidence.json")
	if err := os.WriteFile(evidencePath, evidencePayload, 0o600); err != nil {
		t.Fatalf("WriteFile GUI evidence returned error: %v", err)
	}

	stateRoot := t.TempDir()
	record, err := RecordKnownAppRuntimeStatusLaunchOwnerFixture(KnownAppRuntimeStatusLaunchOwnerFixtureRequest{
		StateRoot:            stateRoot,
		CacheRoot:            t.TempDir(),
		GUISmokeEvidencePath: evidencePath,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppRuntimeStatusLaunchOwnerFixture returned error: %v", err)
	}
	if !record.FixtureReady ||
		record.FixtureState != "ready" ||
		record.AppID != "org.xnix.apps.mines" ||
		record.DisplayName != "Mines" ||
		record.AppVersion != version ||
		record.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		!record.DesktopTriggerReady ||
		!record.OwnerServiceCallReady ||
		!record.DesktopEvidenceHandleForwarded ||
		record.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		record.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" {
		t.Fatalf("unexpected GUI-backed owner fixture: %#v", record)
	}
	if record.StateRootPathExposed ||
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
		t.Fatalf("GUI-backed fixture opened unsafe gates: %#v", record)
	}

	trigger, err := PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppRuntimeStatusLaunchOwnerTrigger returned error: %v", err)
	}
	if !trigger.DesktopTriggerReady ||
		trigger.AppID != "org.xnix.apps.mines" ||
		trigger.OwnerServiceCallArgs[0] != "ShowRuntimeControlledLaunch" ||
		trigger.OwnerServiceCallArgs[2] != record.EvidenceRelativePath {
		t.Fatalf("unexpected GUI-backed owner trigger: %#v", trigger)
	}

	action, err := PreviewKnownAppKDERuntimeStatusLaunchActionTrigger(KnownAppKDERuntimeStatusLaunchActionTriggerRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppKDERuntimeStatusLaunchActionTrigger returned error: %v", err)
	}
	if !action.LaunchRequestCreated ||
		action.AppID != "org.xnix.apps.mines" ||
		action.LaunchRequest.AppID != "org.xnix.apps.mines" ||
		!action.ManagedLauncherArgvReady ||
		action.ExecutionStarted ||
		action.BackendLaunchEnabled {
		t.Fatalf("unexpected GUI-backed action trigger: %#v", action)
	}
}
