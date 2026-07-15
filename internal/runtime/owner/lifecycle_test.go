package owner

import "testing"

func TestNewLifecycleEventsExposeSafeSmokeOwnerJSONLEvents(t *testing.T) {
	events, err := NewLifecycleEvents(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewLifecycleEvents returned error: %v", err)
	}
	if len(events) != 4 {
		t.Fatalf("event count = %d, want 4: %#v", len(events), events)
	}
	wantTypes := []string{"startup", "route-table", "readiness", "shutdown"}
	for index, event := range events {
		if event.Version != "0.2.225" ||
			event.SchemaVersion != "xnix.runtime.owner_lifecycle_event.v1" ||
			event.RequestType != "runtime-owner-lifecycle-event" ||
			event.EventType != wantTypes[index] ||
			event.Sequence != index+1 ||
			event.Mode != "smoke-owner" ||
			event.BusName != "org.xnix.Compatibility1" ||
			event.RouteTableVersion != "0.2.225" ||
			event.RouteCount != 57 ||
			event.GoRouteCount != 57 ||
			event.WriteMethodCount != 4 ||
			!event.ReadOnlyServeReady {
			t.Fatalf("unexpected lifecycle event %d: %#v", index, event)
		}
		if !event.RuntimeOwned ||
			!event.GoRuntimeBacked ||
			event.KDEPolicyOwner ||
			event.KDEMayClaimRuntimeOwnership ||
			event.EventLoopStarted ||
			event.SessionBusClaimed ||
			event.ProductionBusClaimed ||
			event.SystemServiceStarted ||
			event.WriteMethodsEnabled ||
			event.NetworkRequired ||
			event.HostRootModified ||
			event.PrivilegedContainerRequired ||
			event.BackendDetailsExposed {
			t.Fatalf("unexpected lifecycle safety flags at %d: %#v", index, event)
		}
	}
	if events[3].ShutdownReason != "preview-complete" {
		t.Fatalf("shutdown reason = %q, want preview-complete", events[3].ShutdownReason)
	}
}
