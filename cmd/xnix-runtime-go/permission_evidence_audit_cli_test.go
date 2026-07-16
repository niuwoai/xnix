package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/portal"
)

func seedPermissionAuditReceipt(t *testing.T, stateRoot, appID, operation string, outcome portal.Outcome) {
	t.Helper()
	ledger, err := portal.NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("open portal ledger: %v", err)
	}
	record, err := ledger.Create(portal.RequestSpec{ApplicationID: appID, Operation: operation, Reason: "test"})
	if err != nil {
		t.Fatalf("create portal request: %v", err)
	}
	if record.Request.Terminal() {
		return
	}
	if _, err := ledger.Resolve(record.Request.HandleToken, outcome); err != nil {
		t.Fatalf("resolve portal request: %v", err)
	}
}

func TestPermissionEvidenceAuditPreviewCLINoReceipts(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"permission-evidence-audit-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run audit preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.permission_evidence_audit.v1" ||
		payload["request_type"] != "permission-evidence-audit-preview" ||
		payload["overall_state"] != "needs-review" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	for _, key := range []string{"real_portal_call_enabled", "permission_grant_enabled", "permission_revoke_enabled", "receipt_write_enabled", "settings_persisted", "host_root_modified", "state_root_path_exposed"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
}

func TestPermissionEvidenceAuditPreviewCLIWithReceipts(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stateRoot := t.TempDir()
	seedPermissionAuditReceipt(t, stateRoot, app, "file-open", portal.OutcomeGranted)

	var output bytes.Buffer
	if err := run([]string{"permission-evidence-audit-preview", "--registry", registryPath, "--app", app, "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run audit preview: %v", err)
	}
	var payload struct {
		Rows []struct {
			ID               string `json:"id"`
			ConsistencyState string `json:"consistency_state"`
			ReceiptState     string `json:"receipt_state"`
		} `json:"rows"`
	}
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v", err)
	}
	found := false
	for _, row := range payload.Rows {
		if row.ID == "documents" {
			found = true
			if row.ConsistencyState != "consistent" || row.ReceiptState != "granted" {
				t.Fatalf("documents row should be consistent+granted: %+v", row)
			}
		}
	}
	if !found {
		t.Fatalf("documents row missing from output: %s", output.String())
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("CLI output must not expose state-root path: %s", output.String())
	}
}

func TestPermissionEvidenceAuditPreviewCLIMalformedLedgerIsSurfaced(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stateRoot := t.TempDir()
	seedPermissionAuditReceipt(t, stateRoot, app, "file-open", portal.OutcomeGranted)
	broken := filepath.Join(stateRoot, "portal-requests", "broken-1.json")
	if err := os.WriteFile(broken, []byte("not-json"), 0o600); err != nil {
		t.Fatalf("write malformed record: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"permission-evidence-audit-preview", "--registry", registryPath, "--app", app, "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run audit preview: %v", err)
	}
	if !strings.Contains(output.String(), "broken-1") || !strings.Contains(output.String(), "needs-review") {
		t.Fatalf("malformed ledger record should be surfaced without failing: %s", output.String())
	}
}

func TestPermissionEvidenceAuditPreviewCLIRequiresSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	tests := [][]string{
		{"permission-evidence-audit-preview"},
		{"permission-evidence-audit-preview", "--registry", registryPath},
		{"permission-evidence-audit-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
