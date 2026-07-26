package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRecordKnownAppVerifiedCatalogLaunchHandoffConsumesAcceptance(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := filepath.Join(t.TempDir(), "acceptance.json")
	if err := os.WriteFile(acceptancePath, knownAppVerifiedCatalogLaunchHandoffAcceptanceFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile acceptance returned error: %v", err)
	}
	record, err := RecordKnownAppVerifiedCatalogLaunchHandoff(KnownAppVerifiedCatalogLaunchHandoffRequest{
		StateRoot:      stateRoot,
		AcceptancePath: acceptancePath,
		RecordedAtUTC:  time.Date(2026, 7, 27, 3, 4, 5, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("RecordKnownAppVerifiedCatalogLaunchHandoff returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppVerifiedCatalogLaunchHandoffSchemaVersion ||
		record.RequestType != KnownAppVerifiedCatalogLaunchHandoffRequestType ||
		record.Source != KnownAppVerifiedCatalogRunAcceptanceRequestType+"+runtime-owner-launch-handoff" ||
		record.RuntimeMethod != "RecordKnownAppVerifiedCatalogLaunchHandoff" ||
		record.ReadMethod != "GetKnownAppVerifiedCatalogLaunchHandoff" ||
		record.AppID != "7zr" ||
		record.DisplayName != "7-Zip standalone console executable" ||
		record.AppVersion != "26.02" ||
		record.AcceptanceRequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType ||
		record.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" ||
		!record.AcceptanceConsumed ||
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
		t.Fatalf("unexpected launch handoff record evidence: %#v", record)
	}
	if record.HandoffID == "" ||
		record.HandoffState != "persisted-for-runtime-owner-materialization" ||
		!strings.HasPrefix(record.HandoffRelativePath, knownAppVerifiedCatalogLaunchHandoffRecordDir+"/") ||
		!strings.HasSuffix(record.HandoffRelativePath, ".json") ||
		record.HandoffSHA256 == "" ||
		!record.HandoffPayloadDigestVerified {
		t.Fatalf("unexpected persisted handoff fields: %#v", record)
	}
	if record.DesktopCallableRoute != "kde-dbus-verified-catalog-launch-handoff" ||
		record.DesktopDBusMethod != "org.xnix.Compatibility1.RequestRuntimeOwnedLaunch" ||
		!record.DesktopTriggerReady ||
		!record.DesktopEvidenceHandleForwarded ||
		!record.OwnerServiceCallReady ||
		record.OwnerServiceBoundary != "go-runtime-owner-in-process-service" ||
		record.OwnerServiceMethod != "RequestRuntimeOwnedLaunch" ||
		record.OwnerServiceCallType != "desktop-action-dispatch" ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(record.OwnerServiceCallArgs, []string{"RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", record.HandoffRelativePath}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(record.OwnerServiceCLIArgs, []string{"--service-call", "RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", record.HandoffRelativePath}) ||
		!record.OwnerMaterializationRequired ||
		!record.RuntimeOwnerServiceSuppliesInputs ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		!record.KDEForwardsOnlyEvidenceHandle ||
		record.DesktopReceiptFieldsReconstructed ||
		record.DesktopKDEStateRootAccess {
		t.Fatalf("unexpected desktop owner route fields: %#v", record)
	}
	if record.StateRootPathExposed ||
		record.AcceptancePathExposed ||
		record.HandoffPathExposed ||
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
		record.BackendProcessStarted {
		t.Fatalf("launch handoff opened unsafe gates: %#v", record)
	}
	handoffPath := filepath.Join(stateRoot, filepath.FromSlash(record.HandoffRelativePath))
	content, err := os.ReadFile(handoffPath)
	if err != nil {
		t.Fatalf("handoff payload must be written: %v", err)
	}
	if sha256Hex(string(content)) != record.HandoffSHA256 {
		t.Fatalf("handoff digest mismatch")
	}
	if strings.Contains(string(content), stateRoot) || strings.Contains(string(content), acceptancePath) {
		t.Fatalf("handoff payload exposed local paths: %s", string(content))
	}
}

func TestRecordKnownAppVerifiedCatalogLaunchHandoffRejectsUnsafeAcceptance(t *testing.T) {
	stateRoot := t.TempDir()
	var acceptance KnownAppVerifiedCatalogRunAcceptancePreview
	if err := json.Unmarshal(knownAppVerifiedCatalogLaunchHandoffAcceptanceFixture(t), &acceptance); err != nil {
		t.Fatalf("Unmarshal acceptance returned error: %v", err)
	}
	acceptance.AcceptanceReady = false
	content, err := json.MarshalIndent(acceptance, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent acceptance returned error: %v", err)
	}
	acceptancePath := filepath.Join(t.TempDir(), "unsafe-acceptance.json")
	if err := os.WriteFile(acceptancePath, append(content, '\n'), 0o600); err != nil {
		t.Fatalf("WriteFile acceptance returned error: %v", err)
	}
	_, err = RecordKnownAppVerifiedCatalogLaunchHandoff(KnownAppVerifiedCatalogLaunchHandoffRequest{
		StateRoot:      stateRoot,
		AcceptancePath: acceptancePath,
	})
	if err == nil || !strings.Contains(err.Error(), "requires matched accepted q4 evidence") {
		t.Fatalf("unsafe acceptance must be rejected, got: %v", err)
	}
}

func knownAppVerifiedCatalogLaunchHandoffAcceptanceFixture(t *testing.T) []byte {
	t.Helper()
	acceptance := KnownAppVerifiedCatalogRunAcceptancePreview{
		Version:                              currentProjectVersion(t),
		SchemaVersion:                        KnownAppVerifiedCatalogRunAcceptanceSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogRunAcceptanceRequestType,
		Source:                               "known-app-verified-catalog-run-plan+known-existing-windows-app-acceptance",
		RuntimeMethod:                        "PreviewKnownVerifiedApplicationRunAcceptance",
		ReadMethod:                           "GetKnownVerifiedApplicationRunAcceptance",
		AcceptanceType:                       "verified-catalog-app-q4-real-run-acceptance",
		RunPlanConsumed:                      true,
		RunReportConsumed:                    true,
		RunPlanPathExposed:                   false,
		RunReportPathExposed:                 false,
		RemoteHostExposed:                    false,
		RawPathExposed:                       false,
		RawOutputExposed:                     false,
		RuntimeArgvExposed:                   false,
		RunnerPathExposed:                    false,
		RequestedAppID:                       "7zr",
		AppID:                                "7zr",
		DisplayName:                          "7-Zip standalone console executable",
		AppVersion:                           "26.02",
		VerificationState:                    "verified-real-q4-matrix-run",
		CompatibilityState:                   "real-qemu-wine-verified",
		DesktopCatalogState:                  "visible-review-only",
		RunPlanMatched:                       true,
		ExistingWindowsApp:                   true,
		KnownPortableCatalogBacked:           true,
		LaunchAttempted:                      true,
		ChecksumVerified:                     true,
		MarkerObserved:                       true,
		RuntimeStartedIsolatedGuest:          true,
		IsolatedGuestExecutionObserved:       true,
		CompatibilityEngineExecutionObserved: true,
		LoopbackOnlyNetworking:               true,
		SerialLogPersisted:                   true,
		OutputRedacted:                       true,
		Q4ExecutionObserved:                  true,
		HostCompilationAvoided:               true,
		NetworkRequired:                      false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		DockerExecuted:                       false,
		ColimaExecuted:                       false,
		NetworkChecksRun:                     false,
		PackageManagerInvoked:                false,
		AcceptanceReady:                      true,
		DesktopSafeSummary:                   "7-Zip standalone console executable matched the verified catalog run plan and completed the q4 known Windows app acceptance lane.",
	}
	content, err := json.MarshalIndent(acceptance, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent acceptance returned error: %v", err)
	}
	return append(content, '\n')
}
