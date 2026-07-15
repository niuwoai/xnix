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
