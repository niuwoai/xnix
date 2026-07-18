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

func TestOfflineApplicationFixtureMatrixPreviewCommandRendersMatrix(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"offline-application-fixture-matrix-preview", "--runtime-root", "../.."}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.offline_application_fixture_matrix.v1" ||
		payload["request_type"] != "offline-application-fixture-matrix-preview" ||
		payload["matrix_type"] != "offline-cross-application-fixture-matrix" ||
		payload["runtime_method"] != "GetOfflineApplicationFixtureMatrix" ||
		payload["read_method"] != "GetOfflineApplicationFixtureMatrixPreview" ||
		payload["matrix_status"] != "covered-with-unsupported-shape" ||
		payload["ready_for_review"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["review_only"] != true ||
		payload["offline_default"] != true ||
		payload["network_fetch_enabled"] != false ||
		payload["package_manager_invoked"] != false ||
		payload["artifact_staging_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["docker_required"] != false ||
		payload["qemu_required"] != false ||
		payload["file_content_read"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected offline fixture matrix payload: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	rows := payload["rows"].([]any)
	if counts["total"] != float64(7) ||
		counts["unsupported"] != float64(1) ||
		counts["missing_fixture"] != float64(0) ||
		len(rows) != 7 {
		t.Fatalf("unexpected matrix counts: rows=%#v counts=%#v", rows, counts)
	}
	assertOfflineApplicationFixtureMatrixCLISafe(t, output.String())
}

func TestOfflineApplicationFixtureMatrixPreviewCommandFiltersShapes(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"offline-application-fixture-matrix-preview",
		"--shape", "document-editor",
		"--shape", "unsupported",
		"--runtime-root", "../..",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["matrix_status"] != "missing-fixtures" {
		t.Fatalf("filtered matrix must report missing fixtures: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(2) ||
		counts["missing_fixture"] != float64(5) ||
		counts["unsupported"] != float64(1) {
		t.Fatalf("unexpected filtered counts: %#v", counts)
	}
	rows := payload["rows"].([]any)
	unsupported := rows[1].(map[string]any)
	if unsupported["shape_id"] != "unsupported" ||
		unsupported["matrix_state"] != "blocked-unsupported" ||
		unsupported["unsupported_shape"] != true {
		t.Fatalf("unsupported row not blocked: %#v", unsupported)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewCommandConsumesArtifactReceiptRoot(t *testing.T) {
	manifestPath, fixtureRoot, cacheRoot := writeOfflineMatrixArtifactStageFixture(t)
	var stageOutput bytes.Buffer
	if err := run([]string{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot, "--cache-root", cacheRoot}, &stageOutput); err != nil {
		t.Fatalf("artifact-stage-record returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"offline-application-fixture-matrix-preview",
		"--shape", "document-editor",
		"--runtime-root", "../..",
		"--artifact-receipt-root", cacheRoot,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	rows := payload["rows"].([]any)
	row := rows[0].(map[string]any)
	receipt := row["artifact_stage_receipt"].(map[string]any)
	if row["artifact_readiness"] != "local-fixture-ready" ||
		receipt["state"] != "ready" ||
		receipt["relative_path"] != "artifact-ledger/receipts/org.xnix.fixture.document.json" ||
		receipt["required_artifacts_staged"] != true ||
		receipt["root_path_exposed"] != false ||
		strings.Contains(output.String(), cacheRoot) ||
		strings.Contains(output.String(), fixtureRoot) {
		t.Fatalf("artifact receipt root was not consumed safely: %#v", row)
	}
	assertOfflineApplicationFixtureMatrixCLISafe(t, output.String())
}

func TestOfflineApplicationFixtureMatrixPreviewCommandConsumesSnapshotStateRoot(t *testing.T) {
	snapshotRoot := writeOfflineMatrixSnapshotBaseline(t, "matrix-baseline-1")
	var output bytes.Buffer
	err := run([]string{
		"offline-application-fixture-matrix-preview",
		"--shape", "document-editor",
		"--runtime-root", "../..",
		"--snapshot-state-root", snapshotRoot,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	rows := payload["rows"].([]any)
	row := rows[0].(map[string]any)
	baseline := row["snapshot_baseline_receipt"].(map[string]any)
	if row["snapshot_readiness"] != "baseline-receipt-ready" ||
		baseline["state"] != "ready" ||
		baseline["snapshot_id"] != "matrix-baseline-1" ||
		baseline["verified"] != true ||
		baseline["state_root_path_exposed"] != false ||
		baseline["restore_executed"] != false ||
		baseline["snapshot_created"] != false ||
		baseline["snapshot_deleted"] != false ||
		strings.Contains(output.String(), snapshotRoot) {
		t.Fatalf("snapshot state root was not consumed safely: %#v", row)
	}
	assertOfflineApplicationFixtureMatrixCLISafe(t, output.String())
}

func TestOfflineApplicationFixtureMatrixPreviewCommandCoversReadyFixture(t *testing.T) {
	manifestPath, fixtureRoot, cacheRoot := writeOfflineMatrixArtifactStageFixture(t)
	var stageOutput bytes.Buffer
	if err := run([]string{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot, "--cache-root", cacheRoot}, &stageOutput); err != nil {
		t.Fatalf("artifact-stage-record returned error: %v", err)
	}
	snapshotRoot := writeOfflineMatrixSnapshotBaseline(t, "matrix-baseline-2")

	var output bytes.Buffer
	err := run([]string{
		"offline-application-fixture-matrix-preview",
		"--shape", "document-editor",
		"--runtime-root", "../..",
		"--artifact-receipt-root", cacheRoot,
		"--snapshot-state-root", snapshotRoot,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	rows := payload["rows"].([]any)
	row := rows[0].(map[string]any)
	if row["matrix_state"] != "covered-review-only" ||
		row["user_review_required"] != false ||
		jsonArrayLen(row["missing_evidence_ids"]) != 0 ||
		jsonArrayLen(row["blocked_reasons"]) != 0 {
		t.Fatalf("ready fixture row not covered: %#v", row)
	}
	assertOfflineApplicationFixtureMatrixCLISafe(t, output.String())
}

func jsonArrayLen(value any) int {
	if value == nil {
		return 0
	}
	return len(value.([]any))
}

func TestOfflineApplicationFixtureMatrixPreviewCommandRejectsUnknownShape(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"offline-application-fixture-matrix-preview", "--shape", "unknown-shape"}, &output); err == nil {
		t.Fatalf("offline-application-fixture-matrix-preview accepted unknown shape")
	}
}

func writeOfflineMatrixSnapshotBaseline(t *testing.T, snapshotID string) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "runtime-state.json"), []byte(`{"state":"ready"}`), 0o600); err != nil {
		t.Fatalf("WriteFile snapshot source: %v", err)
	}
	store, err := snapshot.New(root)
	if err != nil {
		t.Fatalf("snapshot.New returned error: %v", err)
	}
	if _, err := store.Create(snapshotID, "before-repair"); err != nil {
		t.Fatalf("snapshot Create returned error: %v", err)
	}
	return root
}

func writeOfflineMatrixArtifactStageFixture(t *testing.T) (string, string, string) {
	t.Helper()
	root := t.TempDir()
	fixtureRoot := filepath.Join(root, "fixtures")
	cacheRoot := filepath.Join(root, "cache")
	if err := os.MkdirAll(fixtureRoot, 0o700); err != nil {
		t.Fatalf("MkdirAll fixture root: %v", err)
	}
	digest := artifactDigest("fixture-document-artifact")
	if err := os.WriteFile(filepath.Join(fixtureRoot, digest), []byte("fixture-document-artifact"), 0o600); err != nil {
		t.Fatalf("WriteFile fixture: %v", err)
	}
	manifestPath := filepath.Join(root, "manifest.json")
	manifest := `{
  "application_id": "org.xnix.fixture.document",
  "groups": [
    {"id": "runtime-launch-metadata", "refs": [{"id": "launch.json", "kind": "metadata", "sha256": "` + digest + `", "size": 25, "required": true}]}
  ]
}`
	if err := os.WriteFile(manifestPath, []byte(manifest), 0o600); err != nil {
		t.Fatalf("WriteFile manifest: %v", err)
	}
	return manifestPath, fixtureRoot, cacheRoot
}

func TestOfflineApplicationFixtureMatrixPreviewCommandRejectsPositionalArgs(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"offline-application-fixture-matrix-preview", "extra"}, &output); err == nil {
		t.Fatalf("offline-application-fixture-matrix-preview accepted positional args")
	}
}

func assertOfflineApplicationFixtureMatrixCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("offline fixture matrix CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
