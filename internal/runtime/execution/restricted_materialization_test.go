package execution

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRestrictedMaterializationStorePersistsPlanOnly(t *testing.T) {
	root := t.TempDir()
	authorization, preflight, transaction := restrictedMaterializationFixture(t, root)

	store, err := NewRestrictedMaterializationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedMaterializationStore: %v", err)
	}
	plan, err := store.Record(RestrictedMaterializationRequest{Authorization: authorization, Preflight: preflight, Transaction: transaction})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	loaded, err := store.Load(plan.PlanID)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if loaded.Status != "blocked-plan-materialized" ||
		loaded.MaterializationScope != "test-only-review-plan" ||
		!loaded.PlanMaterialized ||
		!loaded.TestOnly ||
		len(loaded.MaterializedArtifactIDs) != 4 ||
		!sameStringSet(loaded.BlockedByIDs, []string{"recipe-trust", "runtime-write-gate"}) ||
		restrictedMaterializationUnsafe(loaded) {
		t.Fatalf("unexpected restricted materialization plan: %+v", loaded)
	}
}

func TestRestrictedMaterializationStoreRejectsTamperedPlan(t *testing.T) {
	root := t.TempDir()
	authorization, preflight, transaction := restrictedMaterializationFixture(t, root)

	store, err := NewRestrictedMaterializationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedMaterializationStore: %v", err)
	}
	plan, err := store.Record(RestrictedMaterializationRequest{Authorization: authorization, Preflight: preflight, Transaction: transaction})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	path := filepath.Join(root, filepath.FromSlash(plan.RelativePath))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read plan: %v", err)
	}
	tampered := bytes.Replace(data, []byte(`"command_materialized": false`), []byte(`"command_materialized": true`), 1)
	if err := os.WriteFile(path, tampered, 0o600); err != nil {
		t.Fatalf("tamper plan: %v", err)
	}
	if _, err := store.Load(plan.PlanID); err == nil {
		t.Fatal("expected tampered restricted materialization plan to be rejected")
	}
}

func restrictedMaterializationFixture(t *testing.T, root string) (RestrictedAuthorizationReceipt, RestrictedPreflightPacket, LedgerRecord) {
	t.Helper()
	authorizationStore, err := NewRestrictedAuthorizationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedAuthorizationStore: %v", err)
	}
	authorization, err := authorizationStore.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("record authorization: %v", err)
	}
	transaction := LedgerRecord{
		RequestID:     "request-1",
		ApplicationID: "org.xnix.sample.notepad",
		RelativePath:  "execution-ledger/transactions/request-1.json",
		SHA256:        string(make([]byte, 64)),
		Transaction: Transaction{
			RequestID:     "request-1",
			ApplicationID: "org.xnix.sample.notepad",
			State:         StateBlocked,
			Gates:         []Gate{{ID: "recipe-trust", Status: GatePending}, {ID: "runtime-write-gate", Status: GateBlocked}},
		},
	}
	preflightStore, err := NewRestrictedPreflightStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedPreflightStore: %v", err)
	}
	preflight, err := preflightStore.Record(RestrictedPreflightRequest{Authorization: authorization, Transaction: transaction, SafeInputsReady: true})
	if err != nil {
		t.Fatalf("record preflight: %v", err)
	}
	return authorization, preflight, transaction
}
