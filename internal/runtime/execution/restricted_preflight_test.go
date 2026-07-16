package execution

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRestrictedPreflightStorePersistsFailClosedPacket(t *testing.T) {
	root := t.TempDir()
	authorizationStore, err := NewRestrictedAuthorizationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedAuthorizationStore: %v", err)
	}
	authorization, err := authorizationStore.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("record authorization: %v", err)
	}
	transaction := LedgerRecord{RequestID: "request-1", ApplicationID: "org.xnix.sample.notepad", RelativePath: "execution-ledger/transactions/request-1.json", SHA256: string(make([]byte, 64)), Transaction: Transaction{RequestID: "request-1", ApplicationID: "org.xnix.sample.notepad", State: StateBlocked, Gates: []Gate{{ID: "recipe-trust", Status: GatePending}, {ID: "runtime-write-gate", Status: GateBlocked}}}}
	store, err := NewRestrictedPreflightStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedPreflightStore: %v", err)
	}
	packet, err := store.Record(RestrictedPreflightRequest{Authorization: authorization, Transaction: transaction, SafeInputsReady: true})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	loaded, err := store.Load(packet.PacketID)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.Status != "blocked" || !sameStringSet(loaded.BlockerIDs, []string{"recipe-trust", "runtime-write-gate"}) || !loaded.ReadyForPacketAssembly || loaded.ProductImageReady || restrictedPreflightUnsafe(loaded) {
		t.Fatalf("unexpected restricted preflight packet: %+v", loaded)
	}
}

func TestRestrictedPreflightStoreRejectsTamperedPacket(t *testing.T) {
	root := t.TempDir()
	authorizationStore, err := NewRestrictedAuthorizationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedAuthorizationStore: %v", err)
	}
	authorization, err := authorizationStore.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("record authorization: %v", err)
	}
	transaction := LedgerRecord{RequestID: "request-1", ApplicationID: "org.xnix.sample.notepad", SHA256: string(make([]byte, 64)), Transaction: Transaction{RequestID: "request-1", ApplicationID: "org.xnix.sample.notepad", State: StateBlocked, Gates: []Gate{{ID: "recipe-trust", Status: GatePending}, {ID: "runtime-write-gate", Status: GateBlocked}}}}
	store, err := NewRestrictedPreflightStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedPreflightStore: %v", err)
	}
	packet, err := store.Record(RestrictedPreflightRequest{Authorization: authorization, Transaction: transaction, SafeInputsReady: true})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	path := filepath.Join(root, filepath.FromSlash(packet.RelativePath))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read packet: %v", err)
	}
	tampered := bytes.Replace(data, []byte(`"launch_preflight_passed": false`), []byte(`"launch_preflight_passed": true`), 1)
	if err := os.WriteFile(path, tampered, 0o600); err != nil {
		t.Fatalf("tamper packet: %v", err)
	}
	if _, err := store.Load(packet.PacketID); err == nil {
		t.Fatal("expected tampered restricted preflight packet to be rejected")
	}
}
