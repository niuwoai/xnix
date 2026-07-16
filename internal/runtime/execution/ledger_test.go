package execution

import (
	"os"
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/recipe"
)

func readyLedgerTransaction(t *testing.T) Transaction {
	t.Helper()
	inputs := Inputs{
		Trust:                   recipe.TrustState{DigestVerified: true, ProductionTrusted: true},
		Environment:             readyEnvironment(t),
		SnapshotBaselinePresent: true,
		PortalRequiredOps:       []string{"file-open"},
		PortalGrantedOps:        []string{"file-open"},
	}
	p := NewPipeline()
	tx, err := p.Create(app, inputs)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	tx, err = tx.Review(DecisionApproved)
	if err != nil {
		t.Fatalf("Review: %v", err)
	}
	return tx.Preflight(inputs)
}

func TestLedgerRecordsAndLoadsTransactionUnderStateRoot(t *testing.T) {
	stateRoot := t.TempDir()
	ledger, err := NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	tx := readyLedgerTransaction(t)
	record, err := ledger.Record(tx)
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.execution_ledger.v1" ||
		record.RecordType != "execution-transaction-ledger-record" ||
		record.Source != "go-runtime-state-root-execution-ledger" ||
		record.RequestID != tx.RequestID ||
		record.ApplicationID != app ||
		record.RelativePath != "execution-ledger/transactions/"+tx.RequestID+".json" ||
		record.Transaction.State != StateBlocked ||
		record.SHA256 == "" {
		t.Fatalf("unexpected ledger record: %#v", record)
	}
	if !record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		record.StateRootPathExposed ||
		record.LaunchAllowed ||
		record.LaunchEnabled ||
		record.BackendStarted ||
		record.PermissionGranted ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.BackendDetailsExposed {
		t.Fatalf("unexpected ledger safety flags: %#v", record)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "execution-ledger", "transactions", tx.RequestID+".json")); err != nil {
		t.Fatalf("ledger record was not written under state root: %v", err)
	}

	loaded, err := ledger.Load(tx.RequestID)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.RequestID != tx.RequestID || loaded.RelativePath != record.RelativePath || loaded.Transaction.State != StateBlocked {
		t.Fatalf("loaded wrong ledger record: %#v", loaded)
	}
	if filepath.IsAbs(loaded.RelativePath) {
		t.Fatalf("ledger must expose relative paths only: %#v", loaded)
	}
}

func TestLedgerListIsSorted(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	txA := Transaction{RequestID: "txn-b", ApplicationID: app, State: StateBlocked}
	txB := Transaction{RequestID: "txn-a", ApplicationID: app, State: StateBlocked}
	if _, err := ledger.Record(txA); err != nil {
		t.Fatalf("Record txA: %v", err)
	}
	if _, err := ledger.Record(txB); err != nil {
		t.Fatalf("Record txB: %v", err)
	}
	records, err := ledger.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(records) != 2 || records[0].RequestID != "txn-a" || records[1].RequestID != "txn-b" {
		t.Fatalf("ledger list must be sorted: %#v", records)
	}
}

func TestLedgerGuardsUnsafeRootsAndIDs(t *testing.T) {
	if _, err := NewLedger(""); err == nil {
		t.Fatalf("empty state root must be rejected")
	}
	if _, err := NewLedger(string(os.PathSeparator)); err == nil {
		t.Fatalf("filesystem root must be rejected")
	}
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	for _, requestID := range []string{"../escape", "bad/id", "bad id", ".."} {
		if _, err := ledger.Record(Transaction{RequestID: requestID, ApplicationID: app}); err == nil {
			t.Fatalf("unsafe request id %q must be rejected", requestID)
		}
	}
	if _, err := ledger.Record(Transaction{RequestID: "txn-empty-app"}); err == nil {
		t.Fatalf("empty application id must be rejected")
	}
}

func TestLedgerRecordsPendingEnvironmentWithoutEnablingLaunch(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	tx := Transaction{
		RequestID:      "txn-pending-env",
		ApplicationID:  app,
		Profile:        string(environment.ProfileLocal),
		State:          StateBlocked,
		ReviewDecision: DecisionApproved,
		LaunchAllowed:  false,
		LaunchEnabled:  false,
		BackendStarted: false,
	}
	record, err := ledger.Record(tx)
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	if record.LaunchAllowed || record.LaunchEnabled || record.BackendStarted || record.HostRootModified {
		t.Fatalf("ledger record must not enable side effects: %#v", record)
	}
}

func TestLedgerRecordsPortalPermissionReceiptEvidence(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	tx := Transaction{
		RequestID:     "txn-portal-receipt",
		ApplicationID: app,
		State:         StateBlocked,
		PortalPermissionReceipts: []PortalPermissionReceipt{
			{
				HandleToken:       "xnix_org_example_file_open_1",
				Operation:         "file-open",
				RelativePath:      "portal-requests/xnix_org_example_file_open_1.json",
				RequestState:      "completed",
				PermissionState:   "granted",
				PermissionGranted: true,
			},
		},
	}
	record, err := ledger.Record(tx)
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	if record.PortalPermissionReceiptCount != 1 ||
		!record.PortalPermissionReceiptConsumed ||
		len(record.PortalPermissionReceiptPaths) != 1 ||
		record.PortalPermissionReceiptPaths[0] != "portal-requests/xnix_org_example_file_open_1.json" ||
		len(record.PortalPermissionReceiptStates) != 1 ||
		record.PortalPermissionReceiptStates[0] != "file-open:granted/completed" ||
		record.PermissionGranted ||
		record.LaunchEnabled ||
		record.BackendStarted {
		t.Fatalf("unexpected portal receipt ledger evidence: %#v", record)
	}
	if filepath.IsAbs(record.PortalPermissionReceiptPaths[0]) {
		t.Fatalf("portal receipt path must remain relative: %#v", record.PortalPermissionReceiptPaths)
	}
}
