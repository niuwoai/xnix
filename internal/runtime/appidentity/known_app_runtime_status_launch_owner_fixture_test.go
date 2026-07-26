package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

func TestRecordKnownAppRuntimeStatusLaunchOwnerFixtureConsumesVerifiedCatalogAppExecutionEvidenceForMessageBox(t *testing.T) {
	version := currentProjectVersion(t)
	guiPacket := strings.ReplaceAll(knownAppVerifiedCatalogGUIEvidencePacketFixture(), "0.2.640-test", version)
	catalog, err := PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t), []byte(guiPacket))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogWithGUIEvidenceJSON returned error: %v", err)
	}
	catalogContent, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogContent, 0o600); err != nil {
		t.Fatalf("WriteFile catalog returned error: %v", err)
	}
	reportPath := filepath.Join(t.TempDir(), "q4-messagebox-smoke.json")
	if err := os.WriteFile(reportPath, []byte(knownAppVerifiedCatalogMessageBoxRunReportFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile q4 MessageBox report returned error: %v", err)
	}

	execution, err := RunKnownAppVerifiedCatalogAppExecution(KnownAppVerifiedCatalogAppExecutionRequest{
		VerifiedCatalogPath: catalogPath,
		AppID:               "org.xnix.apps.messagebox",
		SmokeReportPath:     reportPath,
	})
	if err != nil {
		t.Fatalf("RunKnownAppVerifiedCatalogAppExecution returned error: %v", err)
	}
	executionContent, err := json.Marshal(execution)
	if err != nil {
		t.Fatalf("Marshal app execution returned error: %v", err)
	}
	executionPath := filepath.Join(t.TempDir(), "known-app-verified-catalog-app-execution-messagebox.json")
	if err := os.WriteFile(executionPath, executionContent, 0o600); err != nil {
		t.Fatalf("WriteFile app execution returned error: %v", err)
	}

	stateRoot := t.TempDir()
	record, err := RecordKnownAppRuntimeStatusLaunchOwnerFixture(KnownAppRuntimeStatusLaunchOwnerFixtureRequest{
		AppID:                "org.xnix.apps.messagebox",
		StateRoot:            stateRoot,
		CacheRoot:            t.TempDir(),
		GUISmokeEvidencePath: executionPath,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppRuntimeStatusLaunchOwnerFixture returned error: %v", err)
	}
	if !record.FixtureReady ||
		record.FixtureState != "ready" ||
		record.AppID != "org.xnix.apps.messagebox" ||
		record.DisplayName != "Xnix MessageBox" ||
		record.Source != "gui-smoke-evidence-preview+runtime-status-evidence" ||
		record.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence" ||
		!record.CompatibilityCenterProjectionReady ||
		!record.KDECenterProjectionReady ||
		!record.KnownAppSmokeEvidenceReady ||
		!record.DesktopTriggerReady ||
		!record.OwnerServiceCallReady ||
		!record.DesktopEvidenceHandleForwarded ||
		record.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		record.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		!record.RuntimeOwnerServiceSuppliesInputs ||
		!record.KDEForwardsOnlyEvidenceHandle {
		t.Fatalf("unexpected app-execution-backed owner fixture: %#v", record)
	}
	if len(record.OwnerServiceCallArgs) != 3 ||
		record.OwnerServiceCallArgs[0] != "ShowRuntimeControlledLaunch" ||
		record.OwnerServiceCallArgs[1] != "evidence-relative-path" ||
		record.OwnerServiceCallArgs[2] != record.EvidenceRelativePath {
		t.Fatalf("owner fixture must expose only the evidence handle to KDE: %#v", record.OwnerServiceCallArgs)
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
		t.Fatalf("app-execution-backed fixture opened unsafe gates: %#v", record)
	}

	plan, err := PreviewKDEControlledLaunchSessionBusSmokePlan(KDEControlledLaunchSessionBusSmokePlanRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchSessionBusSmokePlan returned error: %v", err)
	}
	if !plan.RestrictedSessionBusPlanReady ||
		!plan.DBusControlledLaunchFixturePlanReady ||
		!plan.PrivateSessionBusRequired ||
		!plan.KDEForwardsOnlyEvidenceHandle ||
		plan.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		plan.EvidenceRelativePath != record.EvidenceRelativePath ||
		plan.StateRootPathExposed ||
		plan.BackendDetailsExposed ||
		plan.HostRootModified ||
		plan.ExecutionStarted ||
		plan.BackendProcessStarted {
		t.Fatalf("unexpected session-bus plan from app-execution-backed fixture: %#v", plan)
	}
}
