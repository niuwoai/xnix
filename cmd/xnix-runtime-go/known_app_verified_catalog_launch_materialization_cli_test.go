package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownAppVerifiedCatalogLaunchMaterializationRecordCommandConsumesHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")
	var handoffOutput bytes.Buffer
	if err := run([]string{
		"known-app-verified-catalog-launch-handoff-record",
		"--state-root", stateRoot,
		"--known-app-verified-catalog-run-acceptance", acceptancePath,
	}, &handoffOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-launch-handoff-record returned error: %v", err)
	}
	var handoff map[string]any
	if err := json.Unmarshal(handoffOutput.Bytes(), &handoff); err != nil {
		t.Fatalf("handoff output must be JSON: %v\n%s", err, handoffOutput.String())
	}
	relativePath, ok := handoff["handoff_relative_path"].(string)
	if !ok || relativePath == "" {
		t.Fatalf("handoff output must expose a relative handoff handle: %#v", handoff)
	}
	var output bytes.Buffer
	err := run([]string{
		"known-app-verified-catalog-launch-materialization-record",
		"--state-root", stateRoot,
		"--handoff-relative-path", relativePath,
		"--cache-root", filepath.Join(t.TempDir(), "missing-cache"),
	}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-launch-materialization-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("materialization output must be JSON: %v\n%s", err, output.String())
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_launch_materialization.v1" ||
		payload["request_type"] != "known-app-verified-catalog-launch-materialization-record" ||
		payload["app_id"] != "7zr" ||
		payload["handoff_consumed"] != true ||
		payload["handoff_relative_path"] != relativePath ||
		payload["handoff_digest_verified"] != true ||
		payload["acceptance_ready"] != true ||
		payload["run_plan_matched"] != true ||
		payload["launch_authorization_receipt_recorded"] != true ||
		payload["controlled_session_record_state"] != "blocked" ||
		payload["materialization_state"] != "blocked-managed-artifact-required" ||
		payload["materialization_ready"] != false ||
		payload["runtime_owner_materialized"] != false ||
		payload["owner_materialization_required"] != true ||
		payload["runtime_owner_service_supplies_inputs"] != true ||
		payload["dispatch_runner_required"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false {
		t.Fatalf("unexpected materialization payload: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), acceptancePath) {
		t.Fatalf("materialization output exposed local paths: %s", output.String())
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(relativePath))); err != nil {
		t.Fatalf("handoff payload should remain readable under state root: %v", err)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("materialization output exposed backend term %q: %s", forbidden, output.String())
		}
	}
}

func TestKnownAppVerifiedCatalogLaunchMaterializationRecordCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-launch-materialization-record"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
	err = run([]string{"known-app-verified-catalog-launch-materialization-record", "--state-root", t.TempDir()}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --handoff-relative-path") {
		t.Fatalf("missing handoff path must be rejected, got: %v", err)
	}
}
