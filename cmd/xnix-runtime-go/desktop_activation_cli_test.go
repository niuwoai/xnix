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

func TestDesktopEntryPreviewCommandRendersRecipeBackedContainerGUILauncher(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"desktop-entry-preview", "--registry", "../../runtime/recipes/registry.json", "--app", "org.xnix.sample.notepad"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	entry := output.String()
	required := []string{
		"[Desktop Entry]\n",
		"Name=Sample Notepad\n",
		"Exec=xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U\n",
		"MimeType=application/x-xnix-log;application/x-xnix-txt;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(entry, fragment) {
			t.Fatalf("recipe-backed container GUI desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
	for _, forbidden := range []string{"notepad.exe", "wine ", "docker", "qemu-system", "../../runtime/recipes"} {
		if strings.Contains(strings.ToLower(entry), forbidden) {
			t.Fatalf("recipe-backed desktop entry exposes forbidden term %q: %s", forbidden, entry)
		}
	}
}

func TestDesktopEntryPreviewCommandConsumesActivationRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stageRoot := t.TempDir()

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"desktop-entry-preview", "--registry", registryPath, "--app", app, "--activation-root", stageRoot}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	entry := output.String()
	required := []string{
		"[Desktop Entry]\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"X-Xnix-ApplicationId=org.example.ledger\n",
		"X-Xnix-RuntimeOwned=true\n",
		"MimeType=application/x-xnix-abc;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(entry, fragment) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
	if strings.Contains(entry, stageRoot) {
		t.Fatalf("desktop entry exposed activation root: %s", entry)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(strings.ToLower(entry), forbidden) {
			t.Fatalf("desktop entry exposes forbidden term %q: %s", forbidden, entry)
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

func TestDesktopActivationManifestPreviewCommandRendersKDEContract(t *testing.T) {
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
	err := run([]string{"desktop-activation-manifest-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_manifest.v1" ||
		payload["request_type"] != "desktop-activation-manifest-preview" ||
		payload["manifest_type"] != "kde-desktop-activation-manifest" ||
		payload["runtime_method"] != "GetDesktopActivationManifest" ||
		payload["read_method"] != "GetDesktopActivationManifestPreview" {
		t.Fatalf("unexpected desktop activation manifest schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["launcher_url"] != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop activation manifest identity: %#v", payload)
	}
	bundle := payload["bundle"].(map[string]any)
	if bundle["request_type"] != "desktop-activation-bundle-preview" ||
		bundle["material_count"] != float64(9) ||
		bundle["desktop_files_written"] != false ||
		bundle["host_root_modified"] != false {
		t.Fatalf("unexpected bundle summary: %#v", bundle)
	}
	if payload["entry_point_count"] != float64(7) ||
		payload["activation_material_count"] != float64(9) ||
		payload["contract_section_count"] != float64(7) {
		t.Fatalf("unexpected desktop activation manifest counts: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false || payload["stable_desktop_contract"] != true ||
		payload["seven_entry_point_contract"] != true || payload["portal_mediated_file_access"] != true ||
		payload["desktop_files_written"] != false || payload["mimeapps_written"] != false ||
		payload["manifest_written"] != false || payload["settings_persisted"] != false ||
		payload["notifications_sent"] != false || payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_applied"] != false || payload["live_tray_bridge_enabled"] != false ||
		payload["launch_enabled"] != false || payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false || payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false || payload["raw_windows_executable_exposed"] != false ||
		payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected desktop activation manifest safety flags: %#v", payload)
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
	expectedIDs := []string{"desktop-entry", "dolphin-service-menu", "mimeapps-list", "desktop-integration-manifest", "managed-launcher-artifact", "desktop-activation-receipt"}
	for index, expected := range expectedIDs {
		if fileIDs[index] != expected {
			t.Fatalf("planned_file_ids[%d] = %#v, want %q", index, fileIDs[index], expected)
		}
	}
	if payload["planned_file_count"] != float64(6) ||
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
	launcherArtifact := files[4].(map[string]any)
	if launcherArtifact["kind"] != "managed-launcher-artifact" ||
		launcherArtifact["entry_point"] != "launcher" ||
		launcherArtifact["relative_path"] != "usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json" ||
		launcherArtifact["content_source"] != "cmd/xnix-compat-launch" {
		t.Fatalf("unexpected managed launcher artifact file: %#v", launcherArtifact)
	}
	receipt := files[5].(map[string]any)
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
		staging["planned_file_count"] != float64(6) ||
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
	receiptEvidence := payload["receipt_evidence"].(map[string]any)
	if receiptEvidence["evidence_type"] != "desktop-activation-transaction-receipt-evidence" ||
		receiptEvidence["evidence_state"] != "planned-runtime-gated" ||
		receiptEvidence["receipt_schema_version"] != "xnix.runtime.desktop_activation_receipt.v1" ||
		receiptEvidence["commit_receipt_type"] != "desktop-activation-receipt" ||
		receiptEvidence["rollback_receipt_type"] != "desktop-activation-rollback-receipt" ||
		receiptEvidence["receipt_relative_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		receiptEvidence["required_file_count"] != float64(6) ||
		receiptEvidence["installed_file_digest_required"] != true ||
		receiptEvidence["rollback_digest_required"] != true ||
		receiptEvidence["commit_receipt_required"] != true ||
		receiptEvidence["commit_receipt_planned"] != true ||
		receiptEvidence["commit_receipt_written"] != false ||
		receiptEvidence["rollback_receipt_required"] != true ||
		receiptEvidence["rollback_receipt_planned"] != true ||
		receiptEvidence["rollback_receipt_written"] != false ||
		receiptEvidence["digest_gate_ready"] != false ||
		receiptEvidence["commit_available"] != false ||
		receiptEvidence["rollback_available"] != false ||
		receiptEvidence["runtime_owned"] != true ||
		receiptEvidence["kde_policy_owner"] != false ||
		receiptEvidence["root_path_exposed"] != false ||
		receiptEvidence["target_root_path_exposed"] != false ||
		receiptEvidence["host_root_modified"] != false ||
		receiptEvidence["file_writes_performed"] != false ||
		receiptEvidence["backend_details_exposed"] != false {
		t.Fatalf("unexpected transaction receipt evidence: %#v", receiptEvidence)
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

func TestDesktopActivationStatusPreviewCommandRendersRegistryBackedStatus(t *testing.T) {
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
	err := run([]string{"desktop-activation-status-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_activation_status.v1" ||
		payload["request_type"] != "desktop-activation-status-preview" ||
		payload["status_type"] != "kde-desktop-activation-status" ||
		payload["source"] != "desktop-activation-transaction-preview" ||
		payload["planned_runtime_method"] != "GetDesktopActivationStatus" ||
		payload["current_renderer"] != "xnix-runtime-go desktop-activation-status-preview" ||
		payload["write_method"] != "ActivateDesktopIntegration" {
		t.Fatalf("unexpected activation status schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["install_mode"] != "development" ||
		payload["preflight_decision"] != "development-staging-ready" ||
		payload["transaction_state"] != "transaction-ready" ||
		payload["activation_state"] != "ready-for-runtime-commit" {
		t.Fatalf("unexpected activation status identity or state: %#v", payload)
	}
	transaction := payload["transaction"].(map[string]any)
	if transaction["request_type"] != "desktop-activation-transaction-preview" ||
		transaction["transaction_state"] != "transaction-ready" ||
		transaction["transaction_step_count"] != float64(9) ||
		transaction["ready_step_count"] != float64(9) ||
		transaction["blocked_step_count"] != float64(0) ||
		transaction["rollback_step_count"] != float64(7) ||
		transaction["transaction_ready"] != true ||
		transaction["transaction_committed"] != false {
		t.Fatalf("unexpected transaction summary: %#v", transaction)
	}
	staging := payload["staging"].(map[string]any)
	if staging["request_type"] != "desktop-activation-staging-preview" ||
		staging["staging_state"] != "staging-plan-ready" ||
		staging["planned_file_count"] != float64(6) ||
		staging["activated_entry_point_count"] != float64(7) ||
		staging["receipt_file_id"] != "desktop-activation-receipt" ||
		staging["staging_plan_ready"] != true ||
		staging["installer_may_proceed"] != true ||
		staging["host_root_allowed"] != false {
		t.Fatalf("unexpected staging summary: %#v", staging)
	}
	commitGate := payload["commit_gate"].(map[string]any)
	if commitGate["commit_state"] != "commit-gated" ||
		commitGate["commit_enabled"] != false ||
		commitGate["requires_digest_match"] != true ||
		commitGate["requires_rollback_receipt"] != true ||
		commitGate["requires_runtime_owner"] != true ||
		commitGate["requires_user_review"] != true {
		t.Fatalf("unexpected commit gate: %#v", commitGate)
	}
	statusSignalIDs := payload["status_signal_ids"].([]any)
	expectedSignals := []string{"transaction-plan", "staging-plan", "write-gate", "rollback-plan", "kde-surface"}
	for index, expected := range expectedSignals {
		if statusSignalIDs[index] != expected {
			t.Fatalf("status_signal_ids[%d] = %#v, want %q", index, statusSignalIDs[index], expected)
		}
	}
	blockedReasonIDs := payload["blocked_reason_ids"].([]any)
	expectedReasons := []string{"write-method-disabled", "digest-verification-pending", "rollback-receipt-not-written", "production-owner-not-active"}
	for index, expected := range expectedReasons {
		if blockedReasonIDs[index] != expected {
			t.Fatalf("blocked_reason_ids[%d] = %#v, want %q", index, blockedReasonIDs[index], expected)
		}
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["activation_ready"] != true ||
		payload["activation_committed"] != false ||
		payload["commit_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["file_writes_performed"] != false ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["kde_service_cache_refreshed"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_windows_executable_exposed"] != false ||
		payload["compatibility_storage_exposed"] != false {
		t.Fatalf("unexpected activation status safety flags: %#v", payload)
	}
}

func TestDesktopActivationStatusPreviewCommandConsumesStagedReceipt(t *testing.T) {
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
	stageRoot := filepath.Join(root, "stage")

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"desktop-activation-status-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "development", "--activation-root", stageRoot}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["source"] != "desktop-activation-transaction-preview+desktop-activation-receipt" ||
		payload["activation_state"] != "receipt-backed-runtime-gated" ||
		payload["receipt_backed"] != true ||
		payload["rollback_available"] != true ||
		payload["activation_committed"] != false ||
		payload["commit_enabled"] != false ||
		payload["launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected receipt-backed status: %#v", payload)
	}
	evidence := payload["receipt_evidence"].(map[string]any)
	if evidence["evidence_state"] != "receipt-backed" ||
		evidence["schema_version"] != "xnix.runtime.desktop_activation_receipt.v1" ||
		evidence["receipt_type"] != "desktop-activation-receipt" ||
		evidence["receipt_relative_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		evidence["application_id"] != "org.example.ledger" ||
		evidence["installed_file_count"] != float64(5) ||
		evidence["digest_gate_ready"] != true ||
		evidence["rollback_receipt_ready"] != true ||
		evidence["runtime_owned"] != true ||
		evidence["root_path_exposed"] != false ||
		evidence["host_root_modified"] != false ||
		evidence["backend_details_exposed"] != false ||
		evidence["safe_for_kde"] != true {
		t.Fatalf("unexpected receipt evidence: %#v", evidence)
	}
	statusSignalIDs := payload["status_signal_ids"].([]any)
	if statusSignalIDs[len(statusSignalIDs)-1] != "activation-receipt" {
		t.Fatalf("receipt-backed status must append activation-receipt signal: %#v", statusSignalIDs)
	}
	for _, value := range payload["blocked_reason_ids"].([]any) {
		if value == "rollback-receipt-not-written" {
			t.Fatalf("receipt-backed status must not report an unwritten rollback receipt: %#v", payload["blocked_reason_ids"])
		}
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
		payload["repair_record_count"] != float64(0) || payload["pending_review_count"] != float64(0) ||
		payload["known_app_smoke_evidence_count"] != float64(0) ||
		payload["known_app_smoke_passed_count"] != float64(0) {
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

func TestCompatibilityCenterPreviewCommandConsumesKnownAppSmokeEvidence(t *testing.T) {
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
	err := run([]string{
		"compatibility-center-preview",
		"--registry", registryPath,
		"--known-app-smoke-app", "7zr",
		"--known-app-smoke-name", "7-Zip Console",
		"--known-app-smoke-version", "26.02",
		"--known-app-smoke-source", "staged-launcher-dispatch-smoke",
		"--known-app-smoke-status", "passed",
		"--known-app-smoke-marker-observed",
		"--known-app-smoke-checksum-verified",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_smoke_evidence_count"] != float64(1) ||
		payload["known_app_smoke_passed_count"] != float64(1) ||
		payload["known_app_staged_launcher_passed_count"] != float64(1) ||
		payload["known_app_launch_authorization_required_count"] != float64(1) ||
		payload["action_execution_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected center known app evidence summary: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	evidence := evidenceItems[0].(map[string]any)
	if evidence["app_id"] != "7zr" ||
		evidence["display_name"] != "7-Zip Console" ||
		evidence["app_version"] != "26.02" ||
		evidence["evidence_kind"] != "known-application-managed-smoke" ||
		evidence["evidence_source"] != "staged-launcher-dispatch-smoke" ||
		evidence["smoke_status"] != "passed" ||
		evidence["compatibility_state"] != "validated" ||
		evidence["center_card_state"] != "validated-launch-authorization-required" ||
		evidence["launch_authorization_state"] != "review-required" ||
		evidence["primary_action_id"] != "review-launch-authorization" ||
		evidence["primary_action_label"] != "Review launch authorization" ||
		evidence["primary_action_kind"] != "authorization-review" ||
		evidence["primary_action_enabled"] != true ||
		evidence["direct_launch_enabled"] != false ||
		evidence["marker_observed"] != true ||
		evidence["checksum_verified"] != true ||
		evidence["execution_evidence_recorded"] != true ||
		evidence["staged_launcher_verified"] != true ||
		evidence["runtime_dispatch_verified"] != true ||
		evidence["launch_authorization_required"] != true ||
		evidence["desktop_launch_enabled"] != false ||
		evidence["runtime_owned"] != true ||
		evidence["kde_policy_owner"] != false ||
		evidence["action_execution_enabled"] != false ||
		evidence["backend_launch_enabled"] != false ||
		evidence["host_root_modified"] != false ||
		evidence["backend_details_exposed"] != false ||
		evidence["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected known app smoke evidence: %#v", evidence)
	}
	text := strings.ToLower(output.String())
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Compatibility Center CLI known app evidence exposes forbidden term %q: %s", forbidden, output.String())
		}
	}
}
