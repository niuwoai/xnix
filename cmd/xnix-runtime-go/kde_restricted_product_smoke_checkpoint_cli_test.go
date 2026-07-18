package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestKDERestrictedProductSmokeCheckpointRecordCommandIsReviewOnly(t *testing.T) {
	stateRoot := t.TempDir()
	repositoryRoot := projectRootForRuntimeServiceBindingCommandTest(t)
	var output bytes.Buffer
	err := run([]string{"kde-restricted-product-smoke-checkpoint-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", stateRoot, "--repo-root", repositoryRoot, "--mode", "test-only", "--authorize", "authorize-restricted-test-preparation"}, &output)
	if err != nil {
		t.Fatalf("kde-restricted-product-smoke-checkpoint-record returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("decode command output: %v", err)
	}
	productImage := payload["product_image"].(map[string]any)
	if payload["checkpoint_ready"] != true || payload["ready_for_train_gate"] != true || payload["human_authorization_required"] != true || payload["core_receipt_count"] != float64(10) || productImage["product_image_metadata_ready"] != true || productImage["ready_for_authorized_smoke"] != true {
		t.Fatalf("unexpected restricted product smoke checkpoint: %s", output.String())
	}
	if productImage["production_runtime_ready"] != false {
		t.Fatalf("checkpoint must keep production Runtime activation disabled: %s", output.String())
	}
	for _, sensitive := range []string{stateRoot, repositoryRoot, "image/kinoite/manifest.json"} {
		if strings.Contains(output.String(), sensitive) {
			t.Fatalf("checkpoint output must not expose input path %q: %s", sensitive, output.String())
		}
	}
	for _, key := range []string{"execution_authorized", "docker_executed", "qemu_executed", "product_smoke_executed", "serial_log_persisted", "release_ready", "launch_preflight_passed", "launch_authorized", "execution_approved", "process_start_authorized", "command_materialized", "executable_path_resolved", "backend_selected_for_launch", "backend_launch_enabled", "backend_process_started", "production_bus_ownership", "network_required", "docker_socket_mounted", "host_network_enabled", "broad_host_mount_enabled", "privileged_container_required", "host_root_modified", "state_root_path_exposed", "repository_root_path_exposed", "manifest_source_path_exposed", "raw_command_exposed", "backend_details_exposed"} {
		if payload[key] != false {
			t.Fatalf("unsafe gate %s must remain false: %s", key, output.String())
		}
	}
}

func TestKDERestrictedProductSmokeCheckpointRecordCommandRequiresAuthorization(t *testing.T) {
	args := []string{"kde-restricted-product-smoke-checkpoint-record", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad", "--state-root", t.TempDir(), "--repo-root", projectRootForRuntimeServiceBindingCommandTest(t), "--mode", "test-only"}
	if err := run(args, &bytes.Buffer{}); err == nil {
		t.Fatal("expected exact restricted preparation authorization to be required")
	}
}
