package owner

import (
	"encoding/json"
	"testing"
)

func TestSmokeBatchRecordsCoverReadDispatchAndWriteDenials(t *testing.T) {
	records, err := NewSmokeBatchRecords(projectRoot(t))
	if err != nil {
		t.Fatalf("NewSmokeBatchRecords returned error: %v", err)
	}
	readMethods := SupportedReadDispatchMethods()
	if len(records) != len(readMethods)+4 {
		t.Fatalf("record count = %d, want %d", len(records), len(readMethods)+4)
	}
	seenReads := map[string]bool{}
	seenWrites := map[string]bool{}
	for index, record := range records {
		if record.Version != "0.2.227" ||
			record.SchemaVersion != "xnix.runtime.owner_smoke_batch.v1" ||
			record.RequestType != "runtime-owner-smoke-batch-record" ||
			record.BatchType != "restricted-session-owner-call-batch" ||
			record.Source != "go-runtime-owner-candidate+in-process-read-dispatch+write-gate" ||
			record.Sequence != index+1 ||
			record.ReadDispatchMethodCount != len(readMethods) ||
			record.WriteMethodCount != 4 ||
			!record.RuntimeOwned ||
			!record.GoRuntimeBacked ||
			record.KDEPolicyOwner ||
			record.KDEMayClaimRuntimeOwnership ||
			record.EventLoopStarted ||
			record.SessionBusClaimed ||
			record.ProductionBusClaimed ||
			record.SystemServiceStarted ||
			record.NetworkRequired ||
			record.HostRootModified ||
			record.PrivilegedContainerRequired ||
			record.BackendDetailsExposed ||
			len(record.Payload) == 0 {
			t.Fatalf("unexpected smoke batch record %d: %#v", index, record)
		}
		switch record.RecordType {
		case "read-dispatch":
			seenReads[record.Method] = true
			if !record.ReadOnlyDispatch || record.WriteMethod || !record.RouteReady || !record.DispatchReady {
				t.Fatalf("unexpected read smoke batch record %d: %#v", index, record)
			}
			var payload map[string]any
			if err := json.Unmarshal(record.Payload, &payload); err != nil {
				t.Fatalf("read payload %s unmarshal returned error: %v", record.Method, err)
			}
			if payload["request_type"] != "runtime-owner-read-dispatch" ||
				payload["method"] != record.Method ||
				payload["read_only_dispatch"] != true ||
				payload["session_bus_claimed"] != false ||
				payload["production_bus_claimed"] != false {
				t.Fatalf("unexpected read payload for %s: %#v", record.Method, payload)
			}
		case "write-denial":
			seenWrites[record.Method] = true
			if record.ReadOnlyDispatch || !record.WriteMethod || record.RouteReady || !record.DispatchReady ||
				record.ErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				t.Fatalf("unexpected write smoke batch record %d: %#v", index, record)
			}
			var payload map[string]any
			if err := json.Unmarshal(record.Payload, &payload); err != nil {
				t.Fatalf("write payload %s unmarshal returned error: %v", record.Method, err)
			}
			if payload["method"] != record.Method ||
				payload["dispatch_enabled"] != false ||
				payload["request_created"] != false {
				t.Fatalf("unexpected write payload for %s: %#v", record.Method, payload)
			}
		default:
			t.Fatalf("unexpected smoke batch record type: %s", record.RecordType)
		}
	}
	for _, method := range readMethods {
		if !seenReads[method] {
			t.Fatalf("missing read method in smoke batch: %s", method)
		}
	}
	for _, method := range []string{"InstallRecipe", "Launch", "CreateSnapshot", "RestoreSnapshot"} {
		if !seenWrites[method] {
			t.Fatalf("missing write denial in smoke batch: %s", method)
		}
	}
}
