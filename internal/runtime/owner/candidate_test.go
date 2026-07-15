package owner

import "testing"

func TestNewCandidateExposesSafeGoOwnerCandidate(t *testing.T) {
	candidate, err := NewCandidate(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewCandidate returned error: %v", err)
	}

	if candidate.Version != "0.2.233" ||
		candidate.SchemaVersion != "xnix.runtime.owner_candidate.v1" ||
		candidate.RequestType != "runtime-owner-candidate" ||
		candidate.OwnerType != "go-runtime-owner-candidate" ||
		candidate.Source != "runtime-service-binding-preview+runtime-owner-route-manifest-preview+runtime-write-gate-preview" {
		t.Fatalf("unexpected candidate schema: %#v", candidate)
	}
	if candidate.BusName != "org.xnix.Compatibility1" ||
		candidate.ObjectPath != "/org/xnix/Compatibility1" ||
		candidate.Interface != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected D-Bus identity: %#v", candidate)
	}
	if !candidate.ServiceActivationReady ||
		!candidate.ReadOnlyRouteTableReady ||
		!candidate.ReadOnlyServeReady ||
		candidate.RouteCount != 57 ||
		candidate.GoRouteCount != 57 ||
		candidate.CCoreRouteCount != 0 ||
		candidate.RubyLegacyRouteCount != 0 ||
		len(candidate.Routes) != 57 {
		t.Fatalf("unexpected route readiness: %#v", candidate)
	}
	if candidate.WriteMethodCount != 4 || len(candidate.WriteMethods) != 4 {
		t.Fatalf("unexpected write method count: %#v", candidate.WriteMethods)
	}
	for _, write := range candidate.WriteMethods {
		if write.ErrorName != writeMethodDisabledError ||
			write.DispatchEnabled ||
			write.RequestCreated {
			t.Fatalf("write method is not safely disabled: %#v", write)
		}
	}
	if !candidate.RuntimeOwned ||
		!candidate.GoRuntimeBacked ||
		candidate.KDEPolicyOwner ||
		candidate.KDEMayClaimRuntimeOwnership ||
		!candidate.SmokeOwnerMode ||
		candidate.ProductionOwnerMode ||
		candidate.EventLoopStarted ||
		candidate.SessionBusClaimed ||
		candidate.ProductionBusClaimed ||
		candidate.SystemServiceStarted ||
		candidate.NetworkRequired ||
		candidate.HostRootModified ||
		candidate.PrivilegedContainerRequired ||
		candidate.BackendDetailsExposed ||
		candidate.WriteMethodsEnabled {
		t.Fatalf("unexpected safety flags: %#v", candidate)
	}
	if got, want := candidate.CheckIDs, []string{"service-activation", "read-only-route-table", "write-method-gate", "smoke-owner-mode", "production-bus-claim", "host-safety-boundary"}; !sameStrings(got, want) {
		t.Fatalf("CheckIDs = %#v, want %#v", got, want)
	}
	if candidate.Counts.Total != 6 ||
		candidate.Counts.Passed != 5 ||
		candidate.Counts.Pending != 1 ||
		candidate.Counts.Blocked != 0 {
		t.Fatalf("unexpected candidate check counts: %#v", candidate.Counts)
	}
	if len(candidate.BlockedActions) != 6 ||
		candidate.BlockedActions[0] != "claim production D-Bus name from owner candidate" ||
		len(candidate.NextRequirements) != 4 ||
		candidate.NextRequirements[0] != "Bind this candidate to a restricted session-bus smoke." {
		t.Fatalf("unexpected candidate guidance: actions=%#v next=%#v", candidate.BlockedActions, candidate.NextRequirements)
	}
	if candidate.DesktopSafeSummary != "Runtime owner candidate can serve as a restricted smoke target, but production D-Bus ownership remains gated." {
		t.Fatalf("unexpected summary: %q", candidate.DesktopSafeSummary)
	}
}

func TestDisabledWriteResponseIsDeterministic(t *testing.T) {
	response, err := DisabledWriteResponse("Launch")
	if err != nil {
		t.Fatalf("DisabledWriteResponse returned error: %v", err)
	}
	if response.Method != "Launch" ||
		response.ErrorName != writeMethodDisabledError ||
		response.DispatchEnabled ||
		response.RequestCreated {
		t.Fatalf("unexpected disabled write response: %#v", response)
	}
}

func projectRoot(t *testing.T) string {
	t.Helper()
	return "../../.."
}

func sameStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
