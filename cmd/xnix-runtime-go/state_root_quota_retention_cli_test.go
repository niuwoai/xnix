package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStateRootQuotaRetentionPreviewCLI(t *testing.T) {
	root := t.TempDir()
	writeCLIStateRootFixture(t, root, "diagnostics-ledger/runs/run-001.json", `{"estimated_bytes":256}`)
	writeCLIStateRootFixture(t, root, "portal-requests/handle-001.json", `{"estimated_bytes":128}`)

	var output bytes.Buffer
	if err := run([]string{
		"state-root-quota-retention-preview",
		"--app", "org.example.storage",
		"--state-root", root,
		"--quota-bytes", "128",
	}, &output); err != nil {
		t.Fatalf("run state-root quota retention preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.state_root_quota_retention.v1" ||
		payload["request_type"] != "state-root-quota-retention-preview" ||
		payload["source"] != "go-runtime-state-root-quota-retention-preview" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	if payload["over_quota"] != true ||
		payload["file_deletion_enabled"] != false ||
		payload["directories_created"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("preview should be over-quota dry-run without path exposure: %+v", payload)
	}
	if strings.Contains(output.String(), root) {
		t.Fatalf("CLI output must not expose state-root path: %s", output.String())
	}
}

func TestStateRootQuotaRetentionPreviewCLIMissingRootIsReadOnly(t *testing.T) {
	missingRoot := filepath.Join(t.TempDir(), "missing")
	var output bytes.Buffer
	if err := run([]string{
		"state-root-quota-retention-preview",
		"--app", "org.example.storage",
		"--state-root", missingRoot,
	}, &output); err != nil {
		t.Fatalf("run missing-root preview: %v", err)
	}
	if _, err := os.Stat(missingRoot); !os.IsNotExist(err) {
		t.Fatalf("missing-root preview must not create state root, stat err=%v", err)
	}
	if !strings.Contains(output.String(), "retain-missing-state-root") {
		t.Fatalf("missing-root output should include retention reason: %s", output.String())
	}
	if strings.Contains(output.String(), missingRoot) {
		t.Fatalf("missing-root output must not expose path: %s", output.String())
	}
}

func TestStateRootQuotaRetentionPreviewCLIRequiresFlags(t *testing.T) {
	root := t.TempDir()
	tests := [][]string{
		{"state-root-quota-retention-preview", "--state-root", root},
		{"state-root-quota-retention-preview", "--app", "org.example.storage"},
		{"state-root-quota-retention-preview", "--app", "org.example.storage", "--state-root", root, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}

func writeCLIStateRootFixture(t *testing.T, root string, relativePath string, data string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("create fixture directory: %v", err)
	}
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
}
