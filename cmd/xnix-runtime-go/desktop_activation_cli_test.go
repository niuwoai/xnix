package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestDesktopIdentityPlanCommandLoadsRegistryApplication(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"desktop-identity-plan", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["application_id"] != "org.example.ledger" {
		t.Fatalf("application_id = %#v", payload["application_id"])
	}
	if payload["recipe_source"] != "registry" || payload["registry_name"] != "test-registry" || payload["recipe_digest_verified"] != true {
		t.Fatalf("unexpected registry provenance: %#v", payload)
	}
}

func TestDesktopEntryPreviewCommandRendersManagedLauncher(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"desktop-entry-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	entry := output.String()
	required := []string{
		"[Desktop Entry]\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"MimeType=application/x-xnix-abc;\n",
	}
	for _, fragment := range required {
		if !bytes.Contains(output.Bytes(), []byte(fragment)) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
}

func TestDesktopActivationBundlePreviewCommandAggregatesKDEMaterials(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".log"]}`)
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

	var output bytes.Buffer
	err := run([]string{"desktop-activation-bundle-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_bundle.v1" ||
		payload["request_type"] != "desktop-activation-bundle-preview" ||
		payload["plan_type"] != "normal-linux-application-activation" ||
		payload["runtime_method"] != "GetDesktopActivationBundlePreview" {
		t.Fatalf("unexpected desktop activation bundle schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["launcher_url"] != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop activation bundle identity: %#v", payload)
	}
	if payload["material_count"] != float64(9) {
		t.Fatalf("unexpected material count: %#v", payload)
	}
	materialIDs := payload["material_ids"].([]any)
	expectedIDs := []string{"launcher", "file-association", "desktop-icon", "task-manager", "kwin-window-rule", "system-tray", "notification-center", "unified-settings", "compatibility-center"}
	for index, expected := range expectedIDs {
		if materialIDs[index] != expected {
			t.Fatalf("material_ids[%d] = %#v, want %q", index, materialIDs[index], expected)
		}
	}
	if !bytes.Contains(output.Bytes(), []byte("Exec=xnix-compat-launch --app org.example.ledger %U")) ||
		!bytes.Contains(output.Bytes(), []byte("application/x-xnix-abc=xnix-org.example.ledger.desktop")) ||
		!bytes.Contains(output.Bytes(), []byte("application/x-xnix-log=xnix-org.example.ledger.desktop")) {
		t.Fatalf("activation bundle is missing launcher or MIME material:\n%s", output.String())
	}
	windowIdentity := payload["window_identity"].(map[string]any)
	if windowIdentity["schema_version"] != "xnix.runtime.window_identity.v1" ||
		windowIdentity["launcher_url"] != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected window identity summary: %#v", windowIdentity)
	}
	notification := payload["notification"].(map[string]any)
	if notification["request_type"] != "desktop-notification-preview" ||
		notification["event_type"] != "approval-required" {
		t.Fatalf("unexpected notification summary: %#v", notification)
	}
	settings := payload["settings"].(map[string]any)
	if settings["request_type"] != "settings-preview" ||
		settings["settings_persisted"] != false {
		t.Fatalf("unexpected settings summary: %#v", settings)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false || payload["normal_application_surface"] != true ||
		payload["standard_desktop_entry"] != true || payload["file_association_ready"] != true ||
		payload["task_manager_identity_ready"] != true || payload["kwin_identity_ready"] != true ||
		payload["tray_status_ready"] != true || payload["notification_ready"] != true ||
		payload["settings_ready"] != true || payload["compatibility_center_ready"] != true ||
		payload["desktop_files_written"] != false || payload["mimeapps_written"] != false ||
		payload["settings_persisted"] != false || payload["notifications_sent"] != false ||
		payload["task_manager_entry_active"] != false || payload["kwin_rule_applied"] != false ||
		payload["live_tray_bridge_enabled"] != false || payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false || payload["execution_started"] != false ||
		payload["host_root_modified"] != false || payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false || payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected desktop activation bundle safety flags: %#v", payload)
	}
}

func TestDesktopActivationPreflightPreviewCommandRendersInstallGate(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".log"]}`)
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

	var output bytes.Buffer
	err := run([]string{"desktop-activation-preflight-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_preflight.v1" ||
		payload["request_type"] != "desktop-activation-preflight-preview" ||
		payload["preflight_type"] != "normal-linux-application-activation-preflight" ||
		payload["runtime_method"] != "GetDesktopActivationPreflight" {
		t.Fatalf("unexpected activation preflight schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["install_mode"] != "development" ||
		payload["preflight_decision"] != "development-staging-ready" {
		t.Fatalf("unexpected activation preflight identity or decision: %#v", payload)
	}
	bundle := payload["bundle"].(map[string]any)
	if bundle["request_type"] != "desktop-activation-bundle-preview" ||
		bundle["material_count"] != float64(9) ||
		bundle["desktop_files_written"] != false ||
		bundle["host_root_modified"] != false {
		t.Fatalf("unexpected activation bundle summary: %#v", bundle)
	}
	trust := payload["recipe_trust"].(map[string]any)
	if trust["source"] != "registry" ||
		trust["registry_name"] != "test-registry" ||
		trust["digest_verified"] != true ||
		trust["signature_status"] != "development-only" ||
		trust["trust_decision"] != "development-only" ||
		trust["production_trusted"] != false ||
		trust["development_only"] != true {
		t.Fatalf("unexpected recipe trust: %#v", trust)
	}
	installGate := payload["install_gate"].(map[string]any)
	if installGate["gate_type"] != "recipe-install" ||
		installGate["mode"] != "development" ||
		installGate["decision"] != "allow" {
		t.Fatalf("unexpected install gate: %#v", installGate)
	}
	binding := payload["backend_binding"].(map[string]any)
	if binding["request_type"] != "backend-binding-preview" ||
		binding["binding_state"] != "planned-blocked" ||
		binding["binding_committed"] != false ||
		binding["binding_persisted"] != false ||
		binding["launch_enabled"] != false {
		t.Fatalf("unexpected backend binding summary: %#v", binding)
	}
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{"recipe-digest", "recipe-signature", "recipe-install-gate", "activation-materials", "backend-binding", "staging-root", "host-root-write-gate"}
	for index, expected := range expectedIDs {
		if checkIDs[index] != expected {
			t.Fatalf("check_ids[%d] = %#v, want %q", index, checkIDs[index], expected)
		}
	}
	if payload["check_count"] != float64(7) ||
		payload["passed_check_count"] != float64(4) ||
		payload["pending_check_count"] != float64(2) ||
		payload["blocked_check_count"] != float64(1) {
		t.Fatalf("unexpected check counts: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["normal_application_surface"] != true ||
		payload["desktop_activation_ready"] != true ||
		payload["development_staging_eligible"] != true ||
		payload["production_activation_eligible"] != false ||
		payload["installer_may_proceed"] != true ||
		payload["staging_root_required"] != true ||
		payload["host_root_allowed"] != false ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["manifest_written"] != false ||
		payload["receipt_written"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false ||
		payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected activation preflight safety flags: %#v", payload)
	}
}

func TestDesktopActivationStagingPreviewCommandRendersPlannedFiles(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".log"]}`)
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

	var output bytes.Buffer
	err := run([]string{"desktop-activation-staging-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_staging.v1" ||
		payload["request_type"] != "desktop-activation-staging-preview" ||
		payload["staging_type"] != "kde-activation-staging-plan" ||
		payload["runtime_method"] != "GetDesktopActivationStaging" {
		t.Fatalf("unexpected staging schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["install_mode"] != "development" ||
		payload["preflight_decision"] != "development-staging-ready" ||
		payload["staging_state"] != "staging-plan-ready" {
		t.Fatalf("unexpected staging identity or decision: %#v", payload)
	}
	preflight := payload["preflight"].(map[string]any)
	if preflight["request_type"] != "desktop-activation-preflight-preview" ||
		preflight["install_gate_decision"] != "allow" ||
		preflight["installer_may_proceed"] != true ||
		preflight["host_root_allowed"] != false {
		t.Fatalf("unexpected preflight summary: %#v", preflight)
	}
	fileIDs := payload["planned_file_ids"].([]any)
	expectedIDs := []string{"desktop-entry", "dolphin-service-menu", "mimeapps-list", "desktop-integration-manifest", "desktop-activation-receipt"}
	for index, expected := range expectedIDs {
		if fileIDs[index] != expected {
			t.Fatalf("planned_file_ids[%d] = %#v, want %q", index, fileIDs[index], expected)
		}
	}
	if payload["planned_file_count"] != float64(5) ||
		payload["receipt_file_id"] != "desktop-activation-receipt" ||
		payload["activated_entry_point_count"] != float64(7) {
		t.Fatalf("unexpected staging counts: %#v", payload)
	}
	files := payload["planned_files"].([]any)
	first := files[0].(map[string]any)
	if first["relative_path"] != "usr/share/applications/xnix-org.example.ledger.desktop" ||
		first["mode"] != "0644" ||
		len(first["sha256"].(string)) != 64 ||
		first["planned_for_staging"] != true ||
		first["written"] != false ||
		first["host_root_modified"] != false {
		t.Fatalf("unexpected first staged file: %#v", first)
	}
	receipt := files[4].(map[string]any)
	if receipt["kind"] != "desktop-activation-receipt" ||
		receipt["entry_point"] != "rollback" ||
		receipt["relative_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("unexpected receipt file: %#v", receipt)
	}
	command := payload["installer_command_preview"].([]any)
	if command[0] != "xnix-install-desktop-integration" ||
		command[2] != "runtime-go" ||
		command[4] != "runtime-go" {
		t.Fatalf("unexpected installer command preview: %#v", command)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["development_staging_eligible"] != true ||
		payload["production_activation_eligible"] != false ||
		payload["installer_may_proceed"] != true ||
		payload["staging_plan_ready"] != true ||
		payload["staging_root_required"] != true ||
		payload["staging_root_path_exposed"] != false ||
		payload["host_root_allowed"] != false ||
		payload["file_writes_performed"] != false ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["manifest_written"] != false ||
		payload["receipt_written"] != false ||
		payload["rollback_receipt_planned"] != true ||
		payload["rollback_receipt_written"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false ||
		payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected staging safety flags: %#v", payload)
	}
}

func TestDesktopActivationTransactionPreviewCommandRendersCommitAndRollbackPlan(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".log"]}`)
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

	var output bytes.Buffer
	err := run([]string{"desktop-activation-transaction-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_transaction.v1" ||
		payload["request_type"] != "desktop-activation-transaction-preview" ||
		payload["transaction_type"] != "kde-desktop-activation-transaction" ||
		payload["read_method"] != "GetDesktopActivationTransactionPreview" ||
		payload["write_method"] != "ActivateDesktopIntegration" {
		t.Fatalf("unexpected transaction schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["install_mode"] != "development" ||
		payload["preflight_decision"] != "development-staging-ready" ||
		payload["transaction_state"] != "transaction-ready" {
		t.Fatalf("unexpected transaction identity or state: %#v", payload)
	}
	staging := payload["staging"].(map[string]any)
	if staging["request_type"] != "desktop-activation-staging-preview" ||
		staging["staging_state"] != "staging-plan-ready" ||
		staging["planned_file_count"] != float64(5) ||
		staging["receipt_file_id"] != "desktop-activation-receipt" ||
		staging["installer_may_proceed"] != true ||
		staging["host_root_allowed"] != false {
		t.Fatalf("unexpected staging summary: %#v", staging)
	}
	writeGate := payload["write_gate"].(map[string]any)
	if writeGate["method_name"] != "ActivateDesktopIntegration" ||
		writeGate["gate_decision"] != "disabled-for-preview" ||
		writeGate["write_method_enabled"] != false ||
		writeGate["dispatch_enabled"] != false {
		t.Fatalf("unexpected transaction write gate: %#v", writeGate)
	}
	stepIDs := payload["transaction_step_ids"].([]any)
	expectedSteps := []string{"validate-preflight", "prepare-staging-root", "verify-staged-file-digests", "install-desktop-entry", "install-dolphin-service-menu", "merge-mimeapps-associations", "install-desktop-integration-manifest", "write-rollback-receipt", "refresh-kde-service-cache"}
	for index, expected := range expectedSteps {
		if stepIDs[index] != expected {
			t.Fatalf("transaction_step_ids[%d] = %#v, want %q", index, stepIDs[index], expected)
		}
	}
	rollbackIDs := payload["rollback_step_ids"].([]any)
	expectedRollback := []string{"load-activation-receipt", "verify-installed-file-digests", "remove-desktop-entry", "remove-dolphin-service-menu", "restore-mimeapps-associations", "remove-desktop-integration-manifest", "mark-receipt-rolled-back"}
	for index, expected := range expectedRollback {
		if rollbackIDs[index] != expected {
			t.Fatalf("rollback_step_ids[%d] = %#v, want %q", index, rollbackIDs[index], expected)
		}
	}
	if payload["transaction_step_count"] != float64(len(expectedSteps)) ||
		payload["ready_step_count"] != float64(len(expectedSteps)) ||
		payload["blocked_step_count"] != float64(0) ||
		payload["rollback_step_count"] != float64(len(expectedRollback)) {
		t.Fatalf("unexpected transaction counts: %#v", payload)
	}
	activationCommand := payload["activation_command_preview"].([]any)
	if activationCommand[0] != "xnix-activate-desktop-integration" ||
		activationCommand[2] != "runtime-go" ||
		activationCommand[4] != "staging-root" {
		t.Fatalf("unexpected activation command preview: %#v", activationCommand)
	}
	rollbackCommand := payload["rollback_command_preview"].([]any)
	if rollbackCommand[0] != "xnix-rollback-desktop-integration" ||
		rollbackCommand[2] != "runtime-go" {
		t.Fatalf("unexpected rollback command preview: %#v", rollbackCommand)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["installer_may_proceed"] != true ||
		payload["transaction_plan_created"] != true ||
		payload["transaction_ready"] != true ||
		payload["transaction_committed"] != false ||
		payload["staged_file_digests_required"] != true ||
		payload["staged_file_digests_verified"] != false ||
		payload["staging_root_path_exposed"] != false ||
		payload["target_root_path_exposed"] != false ||
		payload["host_root_allowed"] != false ||
		payload["file_writes_performed"] != false ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["receipt_written"] != false ||
		payload["rollback_receipt_required"] != true ||
		payload["rollback_receipt_planned"] != true ||
		payload["rollback_available"] != false ||
		payload["kde_service_cache_refreshed"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false ||
		payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected transaction safety flags: %#v", payload)
	}
}

func TestCompatibilityCenterPreviewCommandRendersRegistrySummary(t *testing.T) {
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

	var output bytes.Buffer
	err := run([]string{"compatibility-center-preview", "--registry", registryPath}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.compatibility_center.v1" ||
		payload["summary_type"] != "compatibility-center-preview" ||
		payload["desktop"] != "KDE Plasma" {
		t.Fatalf("unexpected center schema: %#v", payload)
	}
	if payload["application_count"] != float64(1) || payload["known_issue_count"] != float64(0) ||
		payload["repair_record_count"] != float64(0) || payload["pending_review_count"] != float64(0) {
		t.Fatalf("unexpected center counts: %#v", payload)
	}
	source := payload["source"].(map[string]any)
	if source["kind"] != "runtime-go-registry" || source["registry_name"] != "test-registry" ||
		source["recipe_digest_verified"] != true || source["recipe_signature_status"] != "development-only" {
		t.Fatalf("unexpected source: %#v", source)
	}
	applications := payload["applications"].([]any)
	application := applications[0].(map[string]any)
	if application["application_id"] != "org.example.ledger" ||
		application["display_name"] != "Example Ledger" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		application["compatibility_state"] != "registered" ||
		application["diagnostics_state"] != "not-run" ||
		application["runtime_mode"] != "Automatic" {
		t.Fatalf("unexpected application: %#v", application)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["action_execution_enabled"] != false ||
		payload["repair_execution_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["settings_persistence_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected center safety flags: %#v", payload)
	}
}
