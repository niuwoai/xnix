package appidentity

import (
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
