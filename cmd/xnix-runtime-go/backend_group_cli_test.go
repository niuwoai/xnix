package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/environment"
)

func writeBackendGroupRegistry(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}
	return registryPath, "org.example.ledger"
}

func TestBackendCapabilityMatrixPreviewCommandRendersMatrix(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-capability-matrix-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_capability_matrix.v1" ||
		payload["runtime_method"] != "GetBackendCapabilityMatrix" ||
		payload["go_runtime_backed"] != true ||
		payload["selection_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend capability matrix payload: %#v", payload)
	}
	if payload["profile_count"] != float64(2) || payload["capability_count"] != float64(7) {
		t.Fatalf("unexpected matrix counts: %#v", payload)
	}

	// The command must reject positional arguments.
	var rejected bytes.Buffer
	if err := run([]string{"backend-capability-matrix-preview", "extra"}, &rejected); err == nil {
		t.Fatalf("backend-capability-matrix-preview must reject arguments")
	}
}

func TestBackendManagerPreviewCommandRendersRuntimeManagedBackends(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-manager-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_manager.v1" ||
		payload["request_type"] != "backend-manager-preview" ||
		payload["manager_type"] != "compatibility-backend-manager" ||
		payload["source"] != "go-runtime-backend-manager" ||
		payload["go_runtime_backed"] != true ||
		payload["runtime_owned"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_visible"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_install_enabled"] != false ||
		payload["backend_download_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["vm_process_started"] != false ||
		payload["backend_details_exposed_to_kde"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected backend manager payload: %#v", payload)
	}
	if payload["backend_count"] != float64(3) ||
		payload["user_facing_profile_count"] != float64(3) ||
		payload["wine_managed"] != true ||
		payload["proton_managed"] != true ||
		payload["windows_vm_managed"] != true {
		t.Fatalf("unexpected backend manager inventory: %#v", payload)
	}
	backendIDs := payload["backend_ids"].([]any)
	if len(backendIDs) != 3 || backendIDs[0] != "wine" || backendIDs[1] != "proton" || backendIDs[2] != "windows-vm" {
		t.Fatalf("unexpected backend ids: %#v", backendIDs)
	}

	var rejected bytes.Buffer
	if err := run([]string{"backend-manager-preview", "extra"}, &rejected); err == nil {
		t.Fatalf("backend-manager-preview must reject arguments")
	}
}

func TestBackendManagerRecordCommandPersistsInventory(t *testing.T) {
	stateRoot := t.TempDir()
	var output bytes.Buffer
	if err := run([]string{"backend-manager-record", "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_manager_record.v1" ||
		payload["record_type"] != "backend-manager-inventory-record" ||
		payload["source"] != "go-runtime-state-root-backend-manager" ||
		payload["relative_path"] != "backend-manager/inventory.json" ||
		payload["state_root_path_exposed"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["vm_process_started"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected backend manager record payload: %#v", payload)
	}
	preview := payload["preview"].(map[string]any)
	if preview["backend_count"] != float64(3) ||
		preview["wine_managed"] != true ||
		preview["proton_managed"] != true ||
		preview["windows_vm_managed"] != true {
		t.Fatalf("unexpected backend manager record preview: %#v", preview)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "backend-manager", "inventory.json")); err != nil {
		t.Fatalf("backend-manager-record must persist inventory under state root: %v", err)
	}

	var missingStateRoot bytes.Buffer
	if err := run([]string{"backend-manager-record"}, &missingStateRoot); err == nil {
		t.Fatalf("backend-manager-record must require --state-root")
	}
}

func TestBackendLifecyclePreviewCommandRendersStages(t *testing.T) {
	registryPath, app := writeBackendGroupRegistry(t)

	var output bytes.Buffer
	if err := run([]string{"backend-lifecycle-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_lifecycle.v1" ||
		payload["runtime_method"] != "GetBackendLifecycle" ||
		payload["lifecycle_state"] != "blocked" ||
		payload["overall_status"] != "not-ready" ||
		payload["launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend lifecycle payload: %#v", payload)
	}
	stages := payload["stage_ids"].([]any)
	if len(stages) != 5 || stages[0] != "recipe-loaded" {
		t.Fatalf("unexpected backend lifecycle stages: %#v", stages)
	}
}

func TestBackendLifecyclePreviewCommandConsumesStateRoot(t *testing.T) {
	registryPath, app := writeBackendGroupRegistry(t)
	stateRoot := t.TempDir()
	lifecycle, err := environment.New(stateRoot)
	if err != nil {
		t.Fatalf("environment.New returned error: %v", err)
	}
	if _, err := lifecycle.Plan(app, environment.ProfileLocal); err != nil {
		t.Fatalf("Plan returned error: %v", err)
	}
	if _, err := lifecycle.Stage(app, environment.ProfileLocal); err != nil {
		t.Fatalf("Stage returned error: %v", err)
	}
	if _, err := lifecycle.SatisfyGate(app, environment.ProfileLocal, "recipe-trust"); err != nil {
		t.Fatalf("SatisfyGate recipe-trust returned error: %v", err)
	}
	if _, err := lifecycle.SatisfyGate(app, environment.ProfileLocal, "backend-binding"); err != nil {
		t.Fatalf("SatisfyGate backend-binding returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"backend-lifecycle-preview", "--registry", registryPath, "--app", app, "--state-root", stateRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["source"] != "run-plan-preview+go-runtime-backend-lifecycle+environment-state-root" ||
		payload["lifecycle_state"] != "staged" ||
		payload["overall_status"] != "not-ready" ||
		payload["state_root_backed"] != true ||
		payload["state_root_path_exposed"] != false ||
		payload["environment_profile"] != "local-compatibility" ||
		payload["state_root_ready"] != true ||
		payload["backend_binding_ready"] != true ||
		payload["portal_review_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["launch_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected state-backed backend lifecycle payload: %#v", payload)
	}
	pending := payload["pending_gates"].([]any)
	if len(pending) != 2 || pending[0] != "portal-policy-review" || pending[1] != "snapshot-baseline" {
		t.Fatalf("unexpected pending gates: %#v", pending)
	}
	stages := payload["stages"].([]any)
	stateRootStage := stages[1].(map[string]any)
	bindingStage := stages[2].(map[string]any)
	if stateRootStage["status"] != "pass" || bindingStage["status"] != "pass" {
		t.Fatalf("unexpected state-backed stages: %#v", stages)
	}
}

func TestBackendLifecycleRecordCommandPersistsStateTransitions(t *testing.T) {
	registryPath, app := writeBackendGroupRegistry(t)
	stateRoot := t.TempDir()

	var plannedOutput bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "plan"}, &plannedOutput); err != nil {
		t.Fatalf("plan run returned error: %v", err)
	}
	var planned map[string]any
	if err := json.Unmarshal(plannedOutput.Bytes(), &planned); err != nil {
		t.Fatalf("Unmarshal planned returned error: %v", err)
	}
	if planned["schema_version"] != "xnix.runtime.backend_lifecycle_record.v1" ||
		planned["record_type"] != "backend-lifecycle-state-record" ||
		planned["source"] != "go-runtime-state-root-backend-lifecycle" ||
		planned["action"] != "plan" ||
		planned["relative_path"] != "environments/org.example.ledger__local-compatibility.json" ||
		planned["state_root_path_exposed"] != false ||
		planned["launch_enabled"] != false ||
		planned["backend_process_started"] != false ||
		planned["host_root_modified"] != false ||
		planned["backend_details_exposed"] != false {
		t.Fatalf("unexpected planned lifecycle record: %#v", planned)
	}
	environmentRecord := planned["environment_record"].(map[string]any)
	if environmentRecord["state"] != "planned" ||
		environmentRecord["launch_enabled"] != false ||
		environmentRecord["backend_started"] != false ||
		environmentRecord["host_root_modified"] != false ||
		environmentRecord["backend_details_shown"] != false {
		t.Fatalf("unexpected planned environment record: %#v", environmentRecord)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "environments", "org.example.ledger__local-compatibility.json")); err != nil {
		t.Fatalf("backend-lifecycle-record must persist environment record under state root: %v", err)
	}

	var stagedOutput bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "stage"}, &stagedOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}
	var staged map[string]any
	if err := json.Unmarshal(stagedOutput.Bytes(), &staged); err != nil {
		t.Fatalf("Unmarshal staged returned error: %v", err)
	}
	stagedRecord := staged["environment_record"].(map[string]any)
	stagedPreview := staged["preview"].(map[string]any)
	if stagedRecord["state"] != "staged" ||
		stagedPreview["lifecycle_state"] != "staged" ||
		stagedPreview["state_root_backed"] != true ||
		stagedPreview["launch_enabled"] != false ||
		stagedPreview["backend_process_started"] != false {
		t.Fatalf("unexpected staged lifecycle record: %#v", staged)
	}

	for _, gate := range []string{"recipe-trust", "backend-binding", "portal-policy-review", "snapshot-baseline"} {
		var gateOutput bytes.Buffer
		if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "satisfy-gate", "--gate", gate}, &gateOutput); err != nil {
			t.Fatalf("satisfy-gate %s returned error: %v", gate, err)
		}
	}

	var readyOutput bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "mark-ready"}, &readyOutput); err != nil {
		t.Fatalf("mark-ready run returned error: %v", err)
	}
	var ready map[string]any
	if err := json.Unmarshal(readyOutput.Bytes(), &ready); err != nil {
		t.Fatalf("Unmarshal ready returned error: %v", err)
	}
	readyRecord := ready["environment_record"].(map[string]any)
	readyPreview := ready["preview"].(map[string]any)
	if readyRecord["state"] != "ready" ||
		readyPreview["overall_status"] != "ready-with-launch-disabled" ||
		readyPreview["launch_enabled"] != false ||
		readyPreview["backend_process_started"] != false ||
		readyPreview["host_root_modified"] != false ||
		readyPreview["backend_details_exposed"] != false {
		t.Fatalf("unexpected ready lifecycle record: %#v", ready)
	}
}

func TestBackendLifecycleRecordCommandRejectsIncompleteRequests(t *testing.T) {
	registryPath, app := writeBackendGroupRegistry(t)
	stateRoot := t.TempDir()

	var missingStateRoot bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--action", "plan"}, &missingStateRoot); err == nil {
		t.Fatalf("backend-lifecycle-record must require --state-root")
	}

	var missingAction bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot}, &missingAction); err == nil {
		t.Fatalf("backend-lifecycle-record must require --action")
	}

	var missingGate bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "satisfy-gate"}, &missingGate); err == nil {
		t.Fatalf("backend-lifecycle-record satisfy-gate must require --gate")
	}

	var positional bytes.Buffer
	if err := run([]string{"backend-lifecycle-record", "--registry", registryPath, "--app", app, "--state-root", stateRoot, "--action", "inspect", "extra"}, &positional); err == nil {
		t.Fatalf("backend-lifecycle-record must reject positional arguments")
	}
}
