package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func artifactDigest(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func writeArtifactStageFixture(t *testing.T) (string, string, string) {
	t.Helper()
	root := t.TempDir()
	fixtureRoot := filepath.Join(root, "fixtures")
	cacheRoot := filepath.Join(root, "cache")
	if err := os.MkdirAll(fixtureRoot, 0o700); err != nil {
		t.Fatalf("MkdirAll fixture root: %v", err)
	}
	launchDigest := artifactDigest("launch-bytes")
	payloadDigest := artifactDigest("payload-bytes")
	for digest, content := range map[string]string{launchDigest: "launch-bytes", payloadDigest: "payload-bytes"} {
		if err := os.WriteFile(filepath.Join(fixtureRoot, digest), []byte(content), 0o600); err != nil {
			t.Fatalf("WriteFile fixture: %v", err)
		}
	}
	manifestPath := filepath.Join(root, "manifest.json")
	manifest := `{
  "application_id": "org.example.ledger",
  "groups": [
    {"id": "runtime-launch-metadata", "refs": [{"id": "launch.json", "kind": "metadata", "sha256": "` + launchDigest + `", "size": 12, "required": true}]},
    {"id": "local-execution-artifacts", "refs": [{"id": "payload.bin", "kind": "execution-artifacts", "sha256": "` + payloadDigest + `", "size": 13, "required": false}]}
  ]
}`
	if err := os.WriteFile(manifestPath, []byte(manifest), 0o600); err != nil {
		t.Fatalf("WriteFile manifest: %v", err)
	}
	return manifestPath, fixtureRoot, cacheRoot
}

func TestArtifactStageRecordCommandStagesLocalFixtureArtifacts(t *testing.T) {
	manifestPath, fixtureRoot, cacheRoot := writeArtifactStageFixture(t)
	var output bytes.Buffer
	err := run([]string{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot, "--cache-root", cacheRoot}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.artifact_stage_receipt.v1" ||
		payload["record_type"] != "compatibility-artifact-stage-receipt" ||
		payload["source"] != "go-runtime-local-fixture-artifact-staging" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["relative_path"] != "artifact-ledger/receipts/org.example.ledger.json" ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["cache_root_path_exposed"] != false ||
		payload["fixture_root_path_exposed"] != false ||
		payload["network_required"] != false ||
		payload["network_fetch_enabled"] != false ||
		payload["package_manager_invoked"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected artifact stage payload: %#v", payload)
	}
	if strings.Contains(output.String(), cacheRoot) || strings.Contains(output.String(), fixtureRoot) {
		t.Fatalf("artifact stage output must not expose roots: %s", output.String())
	}
	plan := payload["plan"].(map[string]any)
	if plan["dry_run"] != false ||
		plan["network_enabled"] != false ||
		plan["host_root_mutated"] != false ||
		len(plan["staged_keys"].([]any)) != 2 {
		t.Fatalf("unexpected artifact stage plan: %#v", plan)
	}
	receiptPath := filepath.Join(cacheRoot, "artifact-ledger", "receipts", "org.example.ledger.json")
	if _, err := os.Stat(receiptPath); err != nil {
		t.Fatalf("artifact receipt was not written under cache root: %v", err)
	}
}

func TestArtifactStageRecordCommandRequiresRoots(t *testing.T) {
	manifestPath, fixtureRoot, cacheRoot := writeArtifactStageFixture(t)
	for _, args := range [][]string{
		{"artifact-stage-record", "--fixture-root", fixtureRoot, "--cache-root", cacheRoot},
		{"artifact-stage-record", "--manifest", manifestPath, "--cache-root", cacheRoot},
		{"artifact-stage-record", "--manifest", manifestPath, "--fixture-root", fixtureRoot},
	} {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected artifact-stage-record args to fail: %#v", args)
		}
	}
}
