package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/snapshot"
)

func seedSnapshot(t *testing.T, stateRoot, dataFile, content, id, reason string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(stateRoot, dataFile), []byte(content), 0o600); err != nil {
		t.Fatalf("write snapshot data: %v", err)
	}
	store, err := snapshot.New(stateRoot)
	if err != nil {
		t.Fatalf("open snapshot store: %v", err)
	}
	if _, err := store.Create(id, reason); err != nil {
		t.Fatalf("create snapshot %q: %v", id, err)
	}
}

func TestSnapshotRestoreCandidatesPreviewCLI(t *testing.T) {
	stateRoot := t.TempDir()
	seedSnapshot(t, stateRoot, "state.dat", "v1", "snap-001", "before-engine-change")
	seedSnapshot(t, stateRoot, "state.dat", "v2", "snap-002", "manual")

	var output bytes.Buffer
	if err := run([]string{"snapshot-restore-candidates-preview", "--app", "org.example.ledger", "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run snapshot restore candidates preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.snapshot_restore_candidates.v1" ||
		payload["overall_state"] != "candidates-available" ||
		payload["candidate_count"].(float64) != 2 {
		t.Fatalf("expected two ranked candidates: %+v", payload)
	}
	for _, key := range []string{"restore_executed", "snapshot_deletion_enabled", "file_content_read", "session_terminated", "backend_launch_enabled", "state_root_path_exposed", "host_root_modified"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("CLI output must not expose the state-root path: %s", output.String())
	}
}

func TestSnapshotRestoreCandidatesPreviewCLIEmptyStateRoot(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	if err := run([]string{"snapshot-restore-candidates-preview", "--app", "org.example.ledger", "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run empty preview: %v", err)
	}
	if !strings.Contains(output.String(), "no-candidates") {
		t.Fatalf("empty store should report no candidates: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, ".xnix-snapshots")); !os.IsNotExist(err) {
		t.Fatalf("read-only preview must not create the snapshot store: %v", err)
	}
}

func TestSnapshotRestoreCandidatesPreviewCLIRequiresFlags(t *testing.T) {
	stateRoot := t.TempDir()
	tests := [][]string{
		{"snapshot-restore-candidates-preview"},
		{"snapshot-restore-candidates-preview", "--app", "org.example.ledger"},
		{"snapshot-restore-candidates-preview", "--state-root", stateRoot},
		{"snapshot-restore-candidates-preview", "--app", "org.example.ledger", "--state-root", stateRoot, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
