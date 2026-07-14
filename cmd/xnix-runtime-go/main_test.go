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

func TestFileOpenPreviewCommandRendersPortalRequest(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc",".xls"]}`)
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
	err := run([]string{"file-open-preview", "--registry", registryPath, "file:///home/test/Documents/book.xls"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.file_open.v1" ||
		payload["request_type"] != "file-open-preview" ||
		payload["source"] != "dolphin-service-menu" ||
		payload["desktop"] != "KDE Plasma" {
		t.Fatalf("unexpected file-open schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["runtime_method"] != "Launch" {
		t.Fatalf("unexpected file-open identity: %#v", payload)
	}
	if payload["portal_required"] != true ||
		payload["portal_interface"] != "org.freedesktop.portal.FileChooser" ||
		payload["portal_method"] != "OpenFile" {
		t.Fatalf("unexpected Portal metadata: %#v", payload)
	}
	if payload["file_count"] != float64(1) ||
		payload["selected_extension"] != ".xls" ||
		payload["selection_mode"] != "extension-match" {
		t.Fatalf("unexpected file selection: %#v", payload)
	}
	action := payload["action"].(map[string]any)
	if action["type"] != "runtime-file-open" {
		t.Fatalf("unexpected action: %#v", action)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["direct_host_file_access"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected file-open safety flags: %#v", payload)
	}
}

func TestKRunnerQueryPreviewCommandSearchesRegistry(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"wine","supported_extensions":[".abc",".xls"]}`)
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
	err := run([]string{"krunner-query-preview", "--registry", registryPath, "--query", "ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.krunner_query.v1" || payload["query_type"] != "krunner-query-plan" {
		t.Fatalf("unexpected KRunner schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" || payload["entry_point"] != "krunner" || payload["query"] != "ledger" {
		t.Fatalf("unexpected KRunner identity: %#v", payload)
	}
	source := payload["source"].(map[string]any)
	if source["kind"] != "runtime-go-registry" || source["registry_name"] != "test-registry" ||
		source["recipe_digest_verified"] != true || source["recipe_signature_status"] != "development-only" {
		t.Fatalf("unexpected source: %#v", source)
	}
	matches := payload["matches"].([]any)
	if len(matches) != 1 {
		t.Fatalf("unexpected matches: %#v", matches)
	}
	match := matches[0].(map[string]any)
	if match["application_id"] != "org.example.ledger" || match["name"] != "Example Ledger" ||
		match["mode_label"] != "Managed compatibility" || match["runtime_owned_launch"] != true ||
		match["backend_details_exposed"] != false {
		t.Fatalf("unexpected match: %#v", match)
	}
	action := match["action"].(map[string]any)
	if action["type"] != "runtime-launch" || action["desktop_entry_id"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected action: %#v", action)
	}
	summary := payload["summary"].(map[string]any)
	if summary["match_count"] != float64(1) || summary["query_execution_enabled"] != false ||
		summary["backend_launch_enabled"] != false || summary["backend_details_exposed"] != false {
		t.Fatalf("unexpected summary: %#v", summary)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["host_root_modified"] != false || payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KRunner safety flags: %#v", payload)
	}
}

func TestMIMEAppsPreviewCommandRendersAssociations(t *testing.T) {
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
	err := run([]string{"mimeapps-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	mimeapps := output.String()
	required := []string{
		"[Default Applications]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop\n",
		"application/x-xnix-log=xnix-org.example.ledger.desktop\n",
		"[Added Associations]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop;\n",
		"application/x-xnix-log=xnix-org.example.ledger.desktop;\n",
	}
	for _, fragment := range required {
		if !bytes.Contains(output.Bytes(), []byte(fragment)) {
			t.Fatalf("MIME apps preview missing %q in:\n%s", fragment, mimeapps)
		}
	}
}

func TestNotificationPreviewCommandRendersKDENotification(t *testing.T) {
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
	err := run([]string{"notification-preview", "--registry", registryPath, "--app", "org.example.ledger", "--event", "approval-required"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.notification.v1" || payload["request_type"] != "desktop-notification-preview" {
		t.Fatalf("unexpected notification schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" || payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected notification identity: %#v", payload)
	}
	if payload["event_type"] != "approval-required" || payload["urgency"] != "critical" ||
		payload["category"] != "compatibility.approval" || payload["requires_user_review"] != true {
		t.Fatalf("unexpected approval notification: %#v", payload)
	}
	actions := payload["actions"].([]any)
	if actions[0] != "open-compatibility-center" || actions[1] != "review-request" {
		t.Fatalf("unexpected actions: %#v", actions)
	}
	if payload["action_execution_enabled"] != false || payload["repair_execution_enabled"] != false ||
		payload["settings_persistence_enabled"] != false || payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected notification safety flags: %#v", payload)
	}
}

func TestSettingsPreviewCommandRendersKDESettings(t *testing.T) {
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
	err := run([]string{"settings-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.settings.v1" || payload["request_type"] != "settings-preview" {
		t.Fatalf("unexpected settings schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" || payload["application_id"] != "org.example.ledger" ||
		payload["display_name"] != "Example Ledger" || payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected settings identity: %#v", payload)
	}
	if payload["section_count"] != float64(5) {
		t.Fatalf("unexpected settings section count: %#v", payload)
	}
	sections := payload["sections"].([]any)
	runMode := sections[0].(map[string]any)
	if runMode["id"] != "run-mode" {
		t.Fatalf("unexpected first settings section: %#v", runMode)
	}
	fields := runMode["fields"].([]any)
	mode := fields[0].(map[string]any)
	if mode["id"] != "mode" || mode["value"] != "automatic" {
		t.Fatalf("unexpected run mode field: %#v", mode)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true || payload["settings_persisted"] != false ||
		payload["settings_persistence_enabled"] != false || payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected settings safety flags: %#v", payload)
	}
}

func TestTrayStatusPreviewCommandRendersKDETrayStatus(t *testing.T) {
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
	err := run([]string{"tray-status-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.tray_status.v1" || payload["status_type"] != "tray-status-preview" {
		t.Fatalf("unexpected tray schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" || payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected tray identity: %#v", payload)
	}
	runtimeActivity := payload["runtime_activity"].(map[string]any)
	if runtimeActivity["registered_application_count"] != float64(1) || runtimeActivity["active_application_count"] != float64(0) {
		t.Fatalf("unexpected runtime activity: %#v", runtimeActivity)
	}
	trayBridge := payload["tray_bridge"].(map[string]any)
	if trayBridge["state"] != "planned" || trayBridge["bridged_tray_application_count"] != float64(0) {
		t.Fatalf("unexpected tray bridge: %#v", trayBridge)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["live_backend_bridge_enabled"] != false || payload["bridge_configuration_persisted"] != false ||
		payload["host_root_modified"] != false || payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected tray safety flags: %#v", payload)
	}
}

func TestWindowIdentityPreviewCommandRendersKDEWindowIdentity(t *testing.T) {
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
	err := run([]string{"window-identity-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.window_identity.v1" {
		t.Fatalf("schema_version = %#v", payload["schema_version"])
	}
	if payload["desktop_file"] != "xnix-org.example.ledger.desktop" || payload["launcher_url"] != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop identity: %#v", payload)
	}
	taskManager := payload["task_manager"].(map[string]any)
	if taskManager["grouping_key"] != "org.example.ledger" || taskManager["pinning_allowed"] != true ||
		taskManager["restore_allowed"] != true || taskManager["skip_taskbar"] != false || taskManager["show_in_switcher"] != true {
		t.Fatalf("unexpected task manager hints: %#v", taskManager)
	}
	kwin := payload["kwin"].(map[string]any)
	if kwin["script_role"] != "identity-and-layout" || kwin["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		kwin["window_manager_policy_only"] != true || kwin["runtime_owns_backend_policy"] != true {
		t.Fatalf("unexpected KWin hints: %#v", kwin)
	}
}
