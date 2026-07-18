package owner

import "testing"

func TestRouteCheckpointClosesReadRouteBandWithWritesDisabled(t *testing.T) {
	checkpoint, err := NewRouteCheckpoint(projectRoot(t))
	if err != nil {
		t.Fatalf("NewRouteCheckpoint returned error: %v", err)
	}
	if checkpoint.Version != currentProjectVersion(t) ||
		checkpoint.SchemaVersion != "xnix.runtime.owner_route_checkpoint.v1" ||
		checkpoint.RequestType != "runtime-owner-route-checkpoint" ||
		checkpoint.CheckpointType != "go-owner-read-route-band-checkpoint" ||
		checkpoint.FormalReadRouteCount != 61 ||
		checkpoint.GoFormalReadRouteCount != 61 ||
		checkpoint.OwnerReadMethodCount != 68 ||
		checkpoint.OwnerLocalReadMethodCount != 7 ||
		checkpoint.SmokeReadRecordCount != 68 ||
		checkpoint.SmokeWriteDenialCount != 4 ||
		!checkpoint.MethodParityReady ||
		!checkpoint.FormalRouteCoverageReady ||
		!checkpoint.OwnerLocalRouteCoverageReady ||
		!checkpoint.SmokeBatchCoverageReady ||
		!checkpoint.DeterministicWriteDenialsReady ||
		!checkpoint.RouteBandReady {
		t.Fatalf("unexpected route checkpoint: %+v", checkpoint)
	}
	if len(checkpoint.Checks) != 5 {
		t.Fatalf("route checkpoint checks = %d, want 5", len(checkpoint.Checks))
	}
	for _, check := range checkpoint.Checks {
		if check.Status != "pass" {
			t.Fatalf("route checkpoint check did not pass: %+v", check)
		}
	}
	if !checkpoint.RuntimeOwned || !checkpoint.GoRuntimeBacked || checkpoint.KDEPolicyOwner ||
		checkpoint.ProductionBusClaimed || checkpoint.SystemServiceStarted || checkpoint.WriteMethodsEnabled ||
		checkpoint.NetworkRequired || checkpoint.HostRootModified || checkpoint.PrivilegedContainerRequired ||
		checkpoint.BackendDetailsExposed {
		t.Fatalf("route checkpoint enabled an unsafe capability: %+v", checkpoint)
	}
	if err := checkpoint.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}
