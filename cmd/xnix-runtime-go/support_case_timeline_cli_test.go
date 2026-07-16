package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSupportCaseTimelinePreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	fixturePath, stateRoot := writeDiagnosticRecordFixture(t)
	if err := run([]string{"diagnostic-run-record", "--state-root", stateRoot, "--app", app, "--run-id", "timeline-001", "--fixture", fixturePath}, &bytes.Buffer{}); err != nil {
		t.Fatalf("record diagnostic run: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"support-case-timeline-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", stateRoot,
		"--runtime-root", projectRootForRuntimeServiceBindingCommandTest(t),
	}, &output); err != nil {
		t.Fatalf("support case timeline returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.support_case_timeline.v1" ||
		payload["request_type"] != "support-case-timeline-preview" ||
		payload["runtime_method"] != "GetSupportCaseTimeline" ||
		payload["read_method"] != "GetSupportCaseTimelinePreview" ||
		payload["ticket_created"] != false ||
		payload["bundle_exported"] != false ||
		payload["ai_provider_called"] != false ||
		payload["repair_executed"] != false ||
		payload["action_executed"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["state_root_path_exposed"] != false {
		t.Fatalf("unexpected support case timeline payload: %#v", payload)
	}
	if payload["event_count"].(float64) == 0 {
		t.Fatalf("timeline must contain events: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), registryPath) || strings.Contains(output.String(), fixturePath) {
		t.Fatalf("timeline output exposed local paths: %s", output.String())
	}
	assertSupportCaseTimelineCLISafe(t, output.String())
}

func TestSupportCaseTimelinePreviewCLIMissingStateRootDoesNotCreate(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	missingRoot := filepath.Join(t.TempDir(), "missing-state")
	var output bytes.Buffer
	if err := run([]string{
		"support-case-timeline-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", missingRoot,
		"--runtime-root", projectRootForRuntimeServiceBindingCommandTest(t),
	}, &output); err != nil {
		t.Fatalf("support case timeline returned error: %v", err)
	}
	if _, err := os.Stat(missingRoot); !os.IsNotExist(err) {
		t.Fatalf("read-only timeline should not create missing state root: %v", err)
	}
	assertSupportCaseTimelineCLISafe(t, output.String())
}

func TestSupportCaseTimelinePreviewCLIMalformedRecordIsBlocked(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stateRoot := t.TempDir()
	runs := filepath.Join(stateRoot, "diagnostics-ledger", "runs")
	if err := os.MkdirAll(runs, 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(filepath.Join(runs, "bad-receipt.json"), []byte("{not-json"), 0o600); err != nil {
		t.Fatalf("WriteFile malformed receipt: %v", err)
	}
	var output bytes.Buffer
	if err := run([]string{
		"support-case-timeline-preview",
		"--registry", registryPath,
		"--app", app,
		"--state-root", stateRoot,
		"--runtime-root", projectRootForRuntimeServiceBindingCommandTest(t),
	}, &output); err != nil {
		t.Fatalf("support case timeline returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["malformed_history"] != true {
		t.Fatalf("timeline must flag malformed history: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["malformed_history_events"] != float64(1) {
		t.Fatalf("timeline must include malformed history event: %#v", counts)
	}
	assertSupportCaseTimelineCLISafe(t, output.String())
}

func TestSupportCaseTimelinePreviewCLIRequiresRecipeSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	cases := [][]string{
		{"support-case-timeline-preview"},
		{"support-case-timeline-preview", "--registry", registryPath},
		{"support-case-timeline-preview", "--registry", registryPath, "--app", app, "--recipe", registryPath},
		{"support-case-timeline-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range cases {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected support-case-timeline-preview to reject args: %#v", args)
		}
	}
}

func assertSupportCaseTimelineCLISafe(t *testing.T, text string) {
	t.Helper()
	lower := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "file://"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("support case timeline CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
