package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExecutionLedgerRecordCommandPersistsUnderStateRoot(t *testing.T) {
	stateRoot := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"execution-ledger-record",
		"--state-root", stateRoot,
		"--app", "org.example.ledger",
		"--decision", "approved",
		"--recipe-trust", "production",
		"--environment", "ready",
		"--snapshot-baseline=true",
		"--portal-required", "file-open,print",
		"--portal-granted", "file-open,print",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.execution_ledger.v1" ||
		payload["record_type"] != "execution-transaction-ledger-record" ||
		payload["source"] != "go-runtime-state-root-execution-ledger" ||
		payload["request_id"] != "xnix-exec-org-example-ledger-1" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_started"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution ledger payload: %#v", payload)
	}
	relativePath := payload["relative_path"].(string)
	if relativePath != "execution-ledger/transactions/xnix-exec-org-example-ledger-1.json" {
		t.Fatalf("unexpected ledger relative path: %q", relativePath)
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("ledger output must not expose the state root path: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(relativePath))); err != nil {
		t.Fatalf("ledger command did not write inside state root: %v", err)
	}
	transaction := payload["transaction"].(map[string]any)
	if transaction["state"] != "blocked" ||
		transaction["launch_allowed"] != false ||
		transaction["launch_enabled"] != false ||
		transaction["backend_started"] != false ||
		transaction["host_root_modified"] != false {
		t.Fatalf("unexpected transaction safety flags: %#v", transaction)
	}
}

func TestExecutionLedgerRecordCommandConsumesPortalReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	handleToken := "xnix_org_example_ledger_file_open_1"

	for _, args := range [][]string{
		{"portal-request-record", "--state-root", stateRoot, "--action", "create", "--app", "org.example.ledger", "--operation", "file-open"},
		{"portal-request-record", "--state-root", stateRoot, "--action", "resolve", "--handle-token", handleToken, "--outcome", "granted"},
		{"portal-request-record", "--state-root", stateRoot, "--action", "complete", "--handle-token", handleToken},
	} {
		var stepOutput bytes.Buffer
		if err := run(args, &stepOutput); err != nil {
			t.Fatalf("portal setup failed for %#v: %v", args, err)
		}
	}

	var output bytes.Buffer
	err := run([]string{
		"execution-ledger-record",
		"--state-root", stateRoot,
		"--app", "org.example.ledger",
		"--decision", "approved",
		"--recipe-trust", "production",
		"--environment", "ready",
		"--snapshot-baseline=true",
		"--portal-required", "file-open",
		"--portal-granted", "",
		"--portal-request", handleToken,
	}, &output)
	if err != nil {
		t.Fatalf("execution-ledger-record returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["portal_permission_receipt_count"] != float64(1) ||
		payload["portal_permission_receipt_consumed"] != true ||
		payload["permission_granted"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_started"] != false {
		t.Fatalf("unexpected Portal receipt ledger payload: %#v", payload)
	}
	receiptPaths := payload["portal_permission_receipt_relative_paths"].([]any)
	if len(receiptPaths) != 1 || receiptPaths[0] != "portal-requests/"+handleToken+".json" {
		t.Fatalf("unexpected Portal receipt path evidence: %#v", receiptPaths)
	}
	transaction := payload["transaction"].(map[string]any)
	receipts := transaction["portal_permission_receipts"].([]any)
	if len(receipts) != 1 {
		t.Fatalf("transaction must include one sanitized receipt: %#v", transaction)
	}
	receipt := receipts[0].(map[string]any)
	if receipt["permission_state"] != "granted" ||
		receipt["request_state"] != "completed" ||
		receipt["execution_approved"] != false ||
		receipt["real_portal_call_enabled"] != false ||
		receipt["state_root_path_exposed"] != false ||
		receipt["host_permission_changed"] != false {
		t.Fatalf("unexpected sanitized receipt: %#v", receipt)
	}
	gates := transaction["gates"].([]any)
	if !jsonGateHasStatus(gates, "portal-permission", "pass") ||
		!jsonGateHasStatus(gates, "runtime-write-gate", "blocked") {
		t.Fatalf("receipt must satisfy only the Portal gate while write gate blocks: %#v", gates)
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("execution ledger output must not expose state root: %s", output.String())
	}
}

func TestExecutionLedgerRecordCommandBlocksDeniedPortalReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	handleToken := "xnix_org_example_ledger_file_open_1"

	for _, args := range [][]string{
		{"portal-request-record", "--state-root", stateRoot, "--action", "create", "--app", "org.example.ledger", "--operation", "file-open"},
		{"portal-request-record", "--state-root", stateRoot, "--action", "resolve", "--handle-token", handleToken, "--outcome", "denied"},
	} {
		var stepOutput bytes.Buffer
		if err := run(args, &stepOutput); err != nil {
			t.Fatalf("portal setup failed for %#v: %v", args, err)
		}
	}

	var output bytes.Buffer
	err := run([]string{
		"execution-ledger-record",
		"--state-root", stateRoot,
		"--app", "org.example.ledger",
		"--portal-required", "file-open",
		"--portal-granted", "",
		"--portal-request", handleToken,
	}, &output)
	if err != nil {
		t.Fatalf("execution-ledger-record returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	transaction := payload["transaction"].(map[string]any)
	gates := transaction["gates"].([]any)
	if !jsonGateHasStatus(gates, "portal-permission", "blocked") {
		t.Fatalf("denied Portal receipt must block the Portal gate: %#v", gates)
	}
	if transaction["launch_enabled"] != false || transaction["backend_started"] != false {
		t.Fatalf("denied receipt must not enable launch: %#v", transaction)
	}
}

func TestExecutionSessionRecordCommandPersistsStatusFromLedger(t *testing.T) {
	stateRoot := t.TempDir()
	requestID := "xnix-exec-org-example-ledger-1"

	var ledgerOutput bytes.Buffer
	err := run([]string{
		"execution-ledger-record",
		"--state-root", stateRoot,
		"--app", "org.example.ledger",
		"--decision", "approved",
		"--recipe-trust", "production",
		"--environment", "ready",
		"--snapshot-baseline=true",
		"--portal-required", "file-open",
		"--portal-granted", "file-open",
	}, &ledgerOutput)
	if err != nil {
		t.Fatalf("execution-ledger-record returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"execution-session-record",
		"--state-root", stateRoot,
		"--request-id", requestID,
	}, &output)
	if err != nil {
		t.Fatalf("execution-session-record returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.execution_session_record.v1" ||
		payload["record_type"] != "execution-session-status-record" ||
		payload["source"] != "go-runtime-state-root-execution-session" ||
		payload["request_id"] != requestID ||
		payload["application_id"] != "org.example.ledger" ||
		payload["session_state"] != "blocked" ||
		payload["task_manager_state"] != "blocked" ||
		payload["tray_state"] != "blocked" ||
		payload["kwin_state"] != "blocked" ||
		payload["compatibility_center_state"] != "waiting-for-runtime-gates" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["status_persisted"] != true ||
		payload["session_active"] != false ||
		payload["live_state_observed"] != false ||
		payload["window_observed"] != false ||
		payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_applied"] != false ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution session record payload: %#v", payload)
	}
	if payload["relative_path"] != "execution-ledger/sessions/"+requestID+".json" ||
		payload["transaction_relative_path"] != "execution-ledger/transactions/"+requestID+".json" {
		t.Fatalf("unexpected execution session paths: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("execution session output must not expose state root: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "execution-ledger", "sessions", requestID+".json")); err != nil {
		t.Fatalf("execution session record was not written under state root: %v", err)
	}
}

func TestExecutionSessionRecordCommandRejectsMissingTransaction(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"execution-session-record", "--state-root", t.TempDir(), "--request-id", "missing"}, &output)
	if err == nil {
		t.Fatalf("execution-session-record must reject missing transactions")
	}
}

func TestExecutionLedgerRecordCommandRequiresStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"execution-ledger-record", "--app", "org.example.ledger"}, &output)
	if err == nil {
		t.Fatalf("execution-ledger-record must require --state-root")
	}
}

func TestExecutionLedgerRecordCommandRejectsInvalidInputs(t *testing.T) {
	stateRoot := t.TempDir()
	for _, args := range [][]string{
		{"execution-ledger-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--decision", "pending"},
		{"execution-ledger-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--recipe-trust", "unknown"},
		{"execution-ledger-record", "--state-root", stateRoot, "--app", "org.example.ledger", "--environment", "unknown"},
	} {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected invalid args to fail: %#v", args)
		}
	}
}

func jsonGateHasStatus(gates []any, id string, status string) bool {
	for _, gateValue := range gates {
		gate := gateValue.(map[string]any)
		if gate["id"] == id && gate["status"] == status {
			return true
		}
	}
	return false
}
