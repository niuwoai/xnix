package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

func seedCrashHangRecord(t *testing.T, root, appID, runID, testType, signalID string, outcome diagnostics.Outcome) {
	t.Helper()
	store, err := diagnostics.NewRunRecordStore(root)
	if err != nil {
		t.Fatalf("open diagnostic run record store: %v", err)
	}
	if _, err := store.Record(diagnostics.RunRecordRequest{
		ApplicationID: appID,
		RunID:         runID,
		Fixture: diagnostics.Fixture{
			TestType: testType,
			Signals:  []diagnostics.Signal{{ID: signalID, Category: "compatibility", Outcome: outcome, Summary: "Fixture signal for tests."}},
		},
	}); err != nil {
		t.Fatalf("record diagnostic run %q: %v", runID, err)
	}
}

func TestCrashHangSignalSummaryPreviewCLI(t *testing.T) {
	root := t.TempDir()
	seedCrashHangRecord(t, root, "org.example.editor", "run-001", "smoke", "app-crash-on-start", diagnostics.OutcomeFail)
	seedCrashHangRecord(t, root, "org.example.editor", "run-002", "smoke", "app-crash-on-start", diagnostics.OutcomeFail)
	seedCrashHangRecord(t, root, "org.example.editor", "run-003", "preflight", "portal-approval-required", diagnostics.OutcomeBlocked)

	var output bytes.Buffer
	if err := run([]string{
		"crash-hang-signal-summary-preview",
		"--state-root", root,
		"--app", "org.example.editor",
	}, &output); err != nil {
		t.Fatalf("run crash-hang signal summary preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.crash_hang_signal_summary.v1" ||
		payload["request_type"] != "crash-hang-signal-summary-preview" ||
		payload["overall_state"] != "needs-review" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	for _, key := range []string{"private_log_read", "file_content_read", "ai_provider_call_enabled", "repair_executed", "backend_process_started", "state_root_path_exposed", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("CLI output must not expose state-root path: %s", output.String())
	}
}

func TestCrashHangSignalSummaryPreviewCLIMalformedRecordIsBlocked(t *testing.T) {
	root := t.TempDir()
	seedCrashHangRecord(t, root, "org.example.editor", "run-001", "smoke", "app-crash-on-start", diagnostics.OutcomeFail)
	malformed := filepath.Join(root, "diagnostics-ledger", "runs", "run-broken.json")
	if err := os.WriteFile(malformed, []byte("not-json"), 0o600); err != nil {
		t.Fatalf("write malformed record: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"crash-hang-signal-summary-preview", "--state-root", root}, &output); err != nil {
		t.Fatalf("run malformed-history preview: %v", err)
	}
	if !strings.Contains(output.String(), "blocked-malformed-history") ||
		!strings.Contains(output.String(), "run-broken") {
		t.Fatalf("malformed history should be surfaced without failing: %s", output.String())
	}
}

func TestCrashHangSignalSummaryPreviewCLIMissingRootIsReadOnly(t *testing.T) {
	missingRoot := filepath.Join(t.TempDir(), "missing")
	var output bytes.Buffer
	if err := run([]string{"crash-hang-signal-summary-preview", "--state-root", missingRoot}, &output); err != nil {
		t.Fatalf("run missing-root preview: %v", err)
	}
	if _, err := os.Stat(missingRoot); !os.IsNotExist(err) {
		t.Fatalf("missing-root preview must not create state root, stat err=%v", err)
	}
	if !strings.Contains(output.String(), "no-history") {
		t.Fatalf("missing-root output should report no history: %s", output.String())
	}
}

func TestCrashHangSignalSummaryPreviewCLIRequiresFlags(t *testing.T) {
	root := t.TempDir()
	tests := [][]string{
		{"crash-hang-signal-summary-preview"},
		{"crash-hang-signal-summary-preview", "--state-root", root, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
