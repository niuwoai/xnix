package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecordKnownAppControlledExecutionSessionBlocksUntilHandoffReady(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	record, err := RecordKnownAppControlledExecutionSession(KnownAppControlledExecutionSessionRecordRequest{
		AppID:         "7zr",
		StateRoot:     stateRoot,
		ReceiptID:     receipt.ReceiptID,
		CacheRoot:     t.TempDir(),
		GuestBoundary: "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppControlledExecutionSession returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppControlledExecutionSessionRecordSchemaVersion ||
		record.RequestType != KnownAppControlledExecutionSessionRecordRequestType ||
		record.Source != KnownAppControlledExecutionSessionRequestType+"+execution-ledger+execution-session-record" ||
		record.AppID != "7zr" ||
		record.ReceiptAccepted != true ||
		record.GuestBoundaryAccepted != true ||
		record.ControlledDispatchRequestCreated ||
		record.ExecutionSessionHandoffCreated ||
		record.SessionHandoffReady ||
		record.RecordState != "blocked" ||
		record.LedgerRecordWritten ||
		record.SessionRecordWritten ||
		record.TransactionRelativePath != "" ||
		record.SessionRelativePath != "" ||
		record.SessionSHA256 != "" ||
		record.SessionState != "not-recorded" ||
		record.CompatibilityCenterState != "not-recorded" ||
		record.RuntimeOwnedExecutionSession ||
		record.StateRootPathExposed ||
		record.TransactionPathExposed ||
		record.SessionPathExposed ||
		record.DispatchStarted ||
		record.ExecutionStarted ||
		record.BackendProcessStarted ||
		record.HostRootModified {
		t.Fatalf("unexpected blocked controlled execution session record: %#v", record)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "execution-ledger")); !os.IsNotExist(err) {
		t.Fatalf("blocked controlled execution session record must not create execution ledger: %v", err)
	}
	encoded, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session record exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("controlled execution session record exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
