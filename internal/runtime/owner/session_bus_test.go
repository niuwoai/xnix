package owner

import (
	"encoding/json"
	"testing"
)

func TestSessionBusSmokeTranscriptWrapsFullOwnerCallBatch(t *testing.T) {
	steps, err := NewSessionBusSmokeTranscript(projectRoot(t))
	if err != nil {
		t.Fatalf("NewSessionBusSmokeTranscript returned error: %v", err)
	}
	if len(steps) != 73 {
		t.Fatalf("session bus smoke step count = %d, want 73", len(steps))
	}
	readCount := 0
	writeCount := 0
	unsupportedCount := 0
	for index, step := range steps {
		if step.Version != currentProjectVersion(t) ||
			step.SchemaVersion != "xnix.runtime.owner_session_bus_smoke.v1" ||
			step.RequestType != "runtime-owner-session-bus-smoke-step" ||
			step.TranscriptType != "restricted-private-session-bus-owner-smoke" ||
			step.Sequence != index+1 ||
			step.BusName != "org.xnix.Compatibility1" ||
			step.ReadDispatchMethodCount != 64 ||
			step.WriteMethodCount != 4 ||
			!step.RuntimeOwned ||
			!step.GoRuntimeBacked ||
			step.KDEPolicyOwner ||
			step.KDEMayClaimRuntimeOwnership ||
			!step.PrivateSessionBus ||
			!step.EventLoopStarted ||
			!step.SessionBusClaimed ||
			step.ProductionBusClaimed ||
			step.SystemServiceStarted ||
			step.WriteMethodsEnabled ||
			step.NetworkRequired ||
			step.HostRootModified ||
			step.PrivilegedContainerRequired ||
			step.BackendDetailsExposed {
			t.Fatalf("unexpected session bus smoke step %d: %#v", index, step)
		}
		switch step.StepType {
		case "startup", "claim-private-session-bus", "route-table-ready", "shutdown":
		case "read-dispatch":
			readCount++
			if !step.ReadOnlyDispatch || step.WriteMethod || !step.RouteReady || !step.DispatchReady || len(step.Payload) == 0 {
				t.Fatalf("unexpected read session bus step %d: %#v", index, step)
			}
			var nested map[string]any
			if err := json.Unmarshal(step.Payload, &nested); err != nil {
				t.Fatalf("nested read payload %d unmarshal returned error: %v", index, err)
			}
			if nested["request_type"] != "runtime-owner-smoke-batch-record" ||
				nested["record_type"] != "read-dispatch" {
				t.Fatalf("unexpected nested read payload %d: %#v", index, nested)
			}
		case "write-denial":
			writeCount++
			if step.ReadOnlyDispatch || !step.WriteMethod || !step.DispatchReady ||
				step.ErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
				len(step.Payload) == 0 {
				t.Fatalf("unexpected write session bus step %d: %#v", index, step)
			}
		case "reject-unsupported-read":
			unsupportedCount++
			if !step.UnsupportedRead ||
				step.DispatchReady != true ||
				step.ErrorName != "org.xnix.Compatibility1.Error.UnsupportedMethod" ||
				len(step.Payload) == 0 {
				t.Fatalf("unexpected unsupported-read session bus step %d: %#v", index, step)
			}
		default:
			t.Fatalf("unexpected session bus step type %q at %d", step.StepType, index)
		}
	}
	if readCount != 64 || writeCount != 4 || unsupportedCount != 1 {
		t.Fatalf("session bus smoke counts read=%d write=%d unsupported=%d, want 64/4/1", readCount, writeCount, unsupportedCount)
	}
}
