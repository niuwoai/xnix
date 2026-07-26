package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKnownAppVerifiedCatalogLaunchHandoffRecordCommandConsumesAcceptance(t *testing.T) {
	stateRoot := t.TempDir()
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")
	var output bytes.Buffer
	err := run([]string{
		"known-app-verified-catalog-launch-handoff-record",
		"--state-root", stateRoot,
		"--known-app-verified-catalog-run-acceptance", acceptancePath,
	}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-launch-handoff-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("handoff output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_launch_handoff.v1" ||
		payload["request_type"] != "known-app-verified-catalog-launch-handoff-record" ||
		payload["app_id"] != "7zr" ||
		payload["acceptance_request_type"] != "known-app-verified-catalog-run-acceptance-preview" ||
		payload["acceptance_consumed"] != true ||
		payload["acceptance_ready"] != true ||
		payload["run_plan_matched"] != true ||
		payload["desktop_trigger_ready"] != true ||
		payload["desktop_dbus_method"] != "org.xnix.Compatibility1.RequestRuntimeOwnedLaunch" ||
		payload["owner_service_call_ready"] != true ||
		payload["owner_service_method"] != "RequestRuntimeOwnedLaunch" ||
		payload["owner_materialization_required"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false {
		t.Fatalf("unexpected launch handoff payload: %#v", payload)
	}
	relativePath, ok := payload["handoff_relative_path"].(string)
	if !ok || !strings.HasPrefix(relativePath, "runtime/known-app-verified-catalog-launch-handoffs/") || !strings.HasSuffix(relativePath, ".json") {
		t.Fatalf("unexpected handoff relative path: %#v", payload)
	}
	ownerServiceArgs, ok := payload["owner_service_call_args"].([]any)
	if !ok || len(ownerServiceArgs) != 3 ||
		ownerServiceArgs[0] != "RequestRuntimeOwnedLaunch" ||
		ownerServiceArgs[1] != "verified-catalog-launch-handoff" ||
		ownerServiceArgs[2] != relativePath {
		t.Fatalf("unexpected owner service args: %#v", payload)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(relativePath))); err != nil {
		t.Fatalf("handoff payload must be written under state root: %v", err)
	}
	if strings.Contains(output.String(), stateRoot) || strings.Contains(output.String(), acceptancePath) {
		t.Fatalf("handoff output exposed local paths: %s", output.String())
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(strings.ToLower(output.String()), forbidden) {
			t.Fatalf("handoff output exposed backend term %q: %s", forbidden, output.String())
		}
	}
}

func TestKnownAppVerifiedCatalogLaunchHandoffRecordCommandRejectsMissingInputs(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-launch-handoff-record"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
	stateRoot := t.TempDir()
	err = run([]string{"known-app-verified-catalog-launch-handoff-record", "--state-root", stateRoot}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --known-app-verified-catalog-run-acceptance") {
		t.Fatalf("missing acceptance must be rejected, got: %v", err)
	}
}
