package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
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

func TestOfflineApplicationFixtureMatrixPreviewCommandRejectsUnknownShape(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"offline-application-fixture-matrix-preview", "--shape", "unknown-shape"}, &output); err == nil {
		t.Fatalf("offline-application-fixture-matrix-preview accepted unknown shape")
	}
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
