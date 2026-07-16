package execution

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRecordSessionPersistsBlockedSessionFromTransaction(t *testing.T) {
	stateRoot := t.TempDir()
	ledger, err := NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	tx := readyLedgerTransaction(t)
	transactionRecord, err := ledger.Record(tx)
	if err != nil {
		t.Fatalf("Record transaction: %v", err)
	}

	session, err := ledger.RecordSession(tx.RequestID)
	if err != nil {
		t.Fatalf("RecordSession: %v", err)
	}
	if session.SchemaVersion != "xnix.runtime.execution_session_record.v1" ||
		session.RecordType != "execution-session-status-record" ||
		session.Source != "go-runtime-state-root-execution-session" ||
		session.RequestID != tx.RequestID ||
		session.ApplicationID != app ||
		session.RelativePath != "execution-ledger/sessions/"+tx.RequestID+".json" ||
		session.TransactionRelativePath != transactionRecord.RelativePath ||
		session.TransactionState != StateBlocked ||
		session.SessionState != "blocked" ||
		session.TaskManagerState != "blocked" ||
		session.TrayState != "blocked" ||
		session.KWinState != "blocked" ||
		session.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		session.SHA256 == "" {
		t.Fatalf("unexpected session record: %#v", session)
	}
	if !session.RuntimeOwned ||
		!session.GoRuntimeBacked ||
		session.KDEPolicyOwner ||
		session.StateRootPathExposed ||
		!session.StatusPersisted ||
		session.SessionCreated ||
		session.SessionRegistered ||
		session.SessionActive ||
		session.LiveStateObserved ||
		session.WindowObserved ||
		session.TaskManagerEntryActive ||
		session.KWinRuleApplied ||
		session.LiveTrayBridgeEnabled ||
		session.LaunchAllowed ||
		session.LaunchEnabled ||
		session.ExecutionStarted ||
		session.BackendProcessStarted ||
		session.PermissionGranted ||
		session.HostRootModified ||
		session.NetworkRequired ||
		session.PrivilegedContainerRequired ||
		session.BackendDetailsExposed {
		t.Fatalf("unexpected session safety flags: %#v", session)
	}
	if filepath.IsAbs(session.RelativePath) || filepath.IsAbs(session.TransactionRelativePath) {
		t.Fatalf("session record must expose relative paths only: %#v", session)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(session.RelativePath))); err != nil {
		t.Fatalf("session record was not written under state root: %v", err)
	}
}

func TestRecordSessionCarriesPortalReceiptEvidence(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	tx := Transaction{
		RequestID:     "txn-session-portal",
		ApplicationID: app,
		State:         StateBlocked,
		PortalPermissionReceipts: []PortalPermissionReceipt{
			{
				Operation:         "file-open",
				RelativePath:      "portal-requests/xnix_org_example_file_open_1.json",
				RequestState:      "completed",
				PermissionState:   "granted",
				PermissionGranted: true,
			},
		},
	}
	if _, err := ledger.Record(tx); err != nil {
		t.Fatalf("Record transaction: %v", err)
	}
	session, err := ledger.RecordSession(tx.RequestID)
	if err != nil {
		t.Fatalf("RecordSession: %v", err)
	}
	if len(session.PortalPermissionReceiptPaths) != 1 ||
		session.PortalPermissionReceiptPaths[0] != "portal-requests/xnix_org_example_file_open_1.json" ||
		len(session.PortalPermissionReceiptStates) != 1 ||
		session.PortalPermissionReceiptStates[0] != "file-open:granted/completed" ||
		session.PermissionGranted ||
		session.LaunchEnabled {
		t.Fatalf("unexpected session Portal evidence: %#v", session)
	}
}

func TestRecordSessionRejectsMissingTransaction(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	if _, err := ledger.RecordSession("missing"); err == nil {
		t.Fatalf("missing transaction must be rejected")
	}
	for _, requestID := range []string{"../escape", "bad/id", "bad id", ".."} {
		if _, err := ledger.RecordSession(requestID); err == nil {
			t.Fatalf("unsafe request id %q must be rejected", requestID)
		}
	}
}
