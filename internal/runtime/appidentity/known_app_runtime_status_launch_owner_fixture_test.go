package appidentity

import "testing"

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
