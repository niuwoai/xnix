package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestServiceCallServesReadDispatchInProcess(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("GetRuntimeWriteGate", []string{"Launch"})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}

	if call.Version != "0.2.244" ||
		call.SchemaVersion != "xnix.runtime.owner_service_call.v1" ||
		call.RequestType != "runtime-owner-service-call" ||
		call.ServiceType != "go-runtime-owner-in-process-service" ||
		call.Method != "GetRuntimeWriteGate" ||
		call.CallType != "read-dispatch" ||
		call.BusName != "org.xnix.Compatibility1" ||
		call.RouteCount != 61 ||
		call.GoRouteCount != 61 ||
		call.WriteMethodCount != 4 ||
		!call.ReadOnlyServeReady ||
		!call.ReadOnlyDispatch ||
		call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		!call.InProcessServiceReady {
		t.Fatalf("unexpected read service call: %#v", call)
	}
	if call.RuntimeOwned != true ||
		call.GoRuntimeBacked != true ||
		call.KDEPolicyOwner != false ||
		call.KDEMayClaimRuntimeOwnership != false ||
		call.EventLoopStarted != false ||
		call.SessionBusClaimed != false ||
		call.ProductionBusClaimed != false ||
		call.SystemServiceStarted != false ||
		call.NetworkRequired != false ||
		call.HostRootModified != false ||
		call.PrivilegedContainerRequired != false ||
		call.BackendDetailsExposed != false {
		t.Fatalf("unexpected read service safety flags: %#v", call)
	}

	var dispatch map[string]any
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if dispatch["schema_version"] != "xnix.runtime.owner_read_dispatch.v1" ||
		dispatch["method"] != "GetRuntimeWriteGate" ||
		dispatch["read_only_dispatch"] != true {
		t.Fatalf("unexpected nested dispatch payload: %#v", dispatch)
	}
}

func TestServiceCallDeniesWriteMethods(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}

	call, err := service.Call("Launch", nil)
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}

	if call.CallType != "write-denial" ||
		call.Method != "Launch" ||
		call.ReadOnlyDispatch ||
		!call.WriteMethod ||
		call.WriteMethodsEnabled ||
		!call.DispatchReady ||
		call.ErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		call.EventLoopStarted ||
		call.SessionBusClaimed ||
		call.ProductionBusClaimed ||
		call.HostRootModified {
		t.Fatalf("unexpected write service call: %#v", call)
	}

	var denial map[string]any
	if err := json.Unmarshal(call.Payload, &denial); err != nil {
		t.Fatalf("payload unmarshal returned error: %v", err)
	}
	if denial["method"] != "Launch" ||
		denial["dispatch_enabled"] != false ||
		denial["request_created"] != false {
		t.Fatalf("unexpected nested write denial payload: %#v", denial)
	}
}

func TestServiceCallRejectsUnknownMethods(t *testing.T) {
	service, err := NewService(projectRoot(t), ModeSmokeOwner)
	if err != nil {
		t.Fatalf("NewService returned error: %v", err)
	}
	_, err = service.Call("GetUnknownThing", nil)
	if err == nil || !strings.Contains(err.Error(), "unsupported Runtime owner service method") {
		t.Fatalf("Call must reject unknown methods, got %v", err)
	}
}
