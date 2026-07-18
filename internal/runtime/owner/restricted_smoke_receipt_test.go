package owner

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRecordRestrictedOwnerSmokeReceiptConsumesPreflightAndSmokeBatch(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordRestrictedOwnerSmokeReceipt(projectRoot(t), stateRoot, RestrictedOwnerSmokeMode, RestrictedOwnerSmokeDirective)
	if err != nil {
		t.Fatalf("RecordRestrictedOwnerSmokeReceipt returned error: %v", err)
	}

	if receipt.Version != currentProjectVersion(t) ||
		receipt.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt.v1" ||
		receipt.RecordType != "restricted-owner-smoke-execution-receipt" ||
		receipt.ReceiptID != "restricted-owner-smoke-"+currentProjectVersion(t) ||
		receipt.Source != "runtime-service-activation-preflight+runtime-owner-smoke-batch" ||
		receipt.Mode != RestrictedOwnerSmokeMode ||
		receipt.Directive != RestrictedOwnerSmokeDirective ||
		receipt.RelativePath != "owner-smoke/restricted-owner-smoke-receipt.json" ||
		len(receipt.SHA256) != 64 {
		t.Fatalf("unexpected restricted owner smoke receipt schema: %#v", receipt)
	}
	if receipt.ActivationPreflight.RequestType != "runtime-service-activation-preflight-preview" ||
		receipt.ActivationPreflight.PreflightDecision != "restricted-owner-smoke-ready" ||
		!receipt.ActivationPreflight.RestrictedSmokeReady ||
		receipt.ActivationPreflight.ProductionActivationReady ||
		receipt.ActivationPreflight.BlockedCheckCount != 0 {
		t.Fatalf("unexpected restricted owner smoke preflight evidence: %#v", receipt.ActivationPreflight)
	}
	if receipt.SmokeBatch.SchemaVersion != "xnix.runtime.owner_smoke_batch.v1" ||
		receipt.SmokeBatch.RequestType != "runtime-owner-smoke-batch-record" ||
		receipt.SmokeBatch.BatchType != "restricted-session-owner-call-batch" ||
		len(receipt.SmokeBatch.SHA256) != 64 ||
		receipt.SmokeBatch.RecordCount != len(SupportedReadDispatchMethods())+4 ||
		receipt.SmokeBatch.ReadDispatchRecordCount != receipt.SmokeBatch.ReadDispatchMethodCount ||
		receipt.SmokeBatch.WriteDenialRecordCount != receipt.SmokeBatch.WriteMethodCount ||
		!receipt.SmokeBatch.AllReadDispatchReady ||
		!receipt.SmokeBatch.AllWriteDenialsReady ||
		receipt.SmokeBatch.EventLoopStarted ||
		receipt.SmokeBatch.SessionBusClaimed ||
		receipt.SmokeBatch.ProductionBusClaimed ||
		receipt.SmokeBatch.SystemServiceStarted ||
		receipt.SmokeBatch.WriteMethodsEnabled ||
		receipt.SmokeBatch.BackendDetailsExposed ||
		receipt.SmokeBatch.HostRootModified {
		t.Fatalf("unexpected restricted owner smoke batch summary: %#v", receipt.SmokeBatch)
	}
	if receipt.CheckCount != 7 ||
		receipt.PassedCheckCount != 7 ||
		!receipt.AllChecksPassed ||
		!sameStrings(receipt.CheckIDs, []string{"activation-preflight-ready", "smoke-batch-read-coverage", "write-denial-coverage", "restricted-authorization", "ownership-boundary", "unsafe-gates-closed", "receipt-boundary"}) {
		t.Fatalf("unexpected restricted owner smoke checks: %#v", receipt.Checks)
	}
	if !receipt.ReceiptPersisted ||
		!receipt.ReceiptReadBack ||
		!receipt.StateRootWritesEnabled ||
		receipt.StateRootWriteScope != "explicit-test-root-only" ||
		receipt.StateRootPathExposed ||
		!receipt.RestrictedSmokeReady ||
		!receipt.RestrictedSmokeAuthorized ||
		receipt.ProductionActivationReady ||
		receipt.ProductionOwnerEnabled ||
		receipt.SystemServiceStarted ||
		receipt.SessionBusClaimed ||
		receipt.ProductionBusClaimed ||
		receipt.WriteMethodsEnabled ||
		receipt.BackendLaunchEnabled ||
		receipt.NetworkRequired ||
		receipt.HostRootModified ||
		receipt.PrivilegedContainerRequired ||
		receipt.BackendDetailsExposed {
		t.Fatalf("unexpected restricted owner smoke safety flags: %#v", receipt)
	}
	if filepath.IsAbs(receipt.RelativePath) {
		t.Fatalf("receipt relative path must not be absolute: %s", receipt.RelativePath)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(receipt.RelativePath))); err != nil {
		t.Fatalf("restricted owner smoke receipt was not persisted: %v", err)
	}
	loaded, err := LoadRestrictedOwnerSmokeReceipt(stateRoot)
	if err != nil {
		t.Fatalf("LoadRestrictedOwnerSmokeReceipt returned error: %v", err)
	}
	if loaded.SHA256 != receipt.SHA256 || loaded.ReceiptID != receipt.ReceiptID {
		t.Fatalf("unexpected loaded receipt: %#v", loaded)
	}
}

func TestRecordRestrictedOwnerSmokeReceiptRequiresExplicitAuthorization(t *testing.T) {
	if _, err := RecordRestrictedOwnerSmokeReceipt(projectRoot(t), t.TempDir(), "", ""); err == nil {
		t.Fatalf("restricted owner smoke receipt must require explicit mode and authorization")
	}
}

func TestRecordRestrictedOwnerSmokeReceiptRejectsUnreadyPreflight(t *testing.T) {
	runtimeRoot := t.TempDir()
	if err := os.WriteFile(filepath.Join(runtimeRoot, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	if _, err := RecordRestrictedOwnerSmokeReceipt(runtimeRoot, t.TempDir(), RestrictedOwnerSmokeMode, RestrictedOwnerSmokeDirective); err == nil {
		t.Fatalf("restricted owner smoke receipt must reject unready activation preflight")
	}
}

func TestRestrictedOwnerSmokeReceiptRejectsTampering(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordRestrictedOwnerSmokeReceipt(projectRoot(t), stateRoot, RestrictedOwnerSmokeMode, RestrictedOwnerSmokeDirective)
	if err != nil {
		t.Fatalf("RecordRestrictedOwnerSmokeReceipt returned error: %v", err)
	}
	path := filepath.Join(stateRoot, filepath.FromSlash(receipt.RelativePath))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	tampered := bytes.Replace(data, []byte(`"restricted_smoke_ready": true`), []byte(`"restricted_smoke_ready": false`), 1)
	if err := os.WriteFile(path, tampered, 0o600); err != nil {
		t.Fatalf("WriteFile tampered receipt returned error: %v", err)
	}
	if _, err := LoadRestrictedOwnerSmokeReceipt(stateRoot); err == nil {
		t.Fatalf("tampered restricted owner smoke receipt must be rejected")
	}
}
