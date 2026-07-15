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

func TestKDECenterPagePreviewCommandRendersApplicationPage(t *testing.T) {
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
	err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/example.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_center_page.v1" ||
		payload["request_type"] != "kde-center-page-preview" ||
		payload["page_type"] != "compatibility-center-application-page" ||
		payload["source"] != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+settings-preview" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetKDECenterPage" ||
		payload["read_method"] != "GetKDECenterPagePreview" {
		t.Fatalf("unexpected KDE center page schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected KDE center page identity: %#v", payload)
	}
	header := payload["header"].(map[string]any)
	if header["title"] != "Example Ledger" ||
		header["badge"] != "Review ready" ||
		header["badge_tone"] != "warning" ||
		header["primary_action_target"] != "compatibility-center-gates" ||
		header["primary_action_enabled"] != true ||
		header["backend_details_exposed"] != false {
		t.Fatalf("unexpected center page header: %#v", header)
	}
	summary := payload["application_summary"].(map[string]any)
	if summary["application_id"] != "org.example.ledger" ||
		summary["compatibility_state"] != "registered" ||
		summary["runtime_mode"] != "Automatic" ||
		summary["action_execution_enabled"] != false ||
		summary["repair_execution_enabled"] != false ||
		summary["backend_launch_enabled"] != false ||
		summary["settings_persistence_enabled"] != false ||
		summary["host_root_modified"] != false ||
		summary["backend_details_exposed"] != false {
		t.Fatalf("unexpected application summary: %#v", summary)
	}
	deck := payload["action_deck"].(map[string]any)
	if deck["request_type"] != "kde-action-card-deck-preview" ||
		deck["card_count"] != float64(7) ||
		deck["waiting_card_count"] != float64(7) ||
		deck["navigation_action_count"] != float64(29) ||
		deck["disabled_action_count"] != float64(21) ||
		deck["primary_card_id"] != "org.example.ledger:review-launcher-action:card" ||
		deck["action_queue_created"] != true ||
		deck["deck_preview_created"] != true ||
		deck["deck_persisted"] != false ||
		deck["cards_persisted"] != false ||
		deck["card_actions_enabled"] != false ||
		deck["runtime_launch_approval"] != false ||
		deck["launch_allowed"] != false ||
		deck["execution_started"] != false ||
		deck["backend_details_exposed"] != false {
		t.Fatalf("unexpected action deck: %#v", deck)
	}
	settings := payload["settings_snapshot"].(map[string]any)
	if settings["request_type"] != "settings-preview" ||
		settings["settings_state"] != "planned" ||
		settings["section_count"] != float64(5) ||
		settings["settings_persisted"] != false ||
		settings["settings_persistence_enabled"] != false ||
		settings["host_root_modified"] != false ||
		settings["backend_details_exposed"] != false {
		t.Fatalf("unexpected settings snapshot: %#v", settings)
	}
	navigation := payload["navigation"].([]any)
	backend := payload["backend_selection_snapshot"].(map[string]any)
	if backend["request_type"] != "backend-selection-preview" ||
		backend["runtime_method"] != "GetBackendSelectionPlan" ||
		backend["recommended_profile_id"] != "local-compatibility" ||
		backend["candidate_count"] != float64(2) ||
		backend["selection_committed"] != false ||
		backend["selection_change_enabled"] != false ||
		backend["backend_launch_enabled"] != false ||
		backend["environment_created"] != false ||
		backend["host_root_modified"] != false ||
		backend["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend selection snapshot: %#v", backend)
	}
	activation := payload["activation_status_snapshot"].(map[string]any)
	if activation["request_type"] != "desktop-activation-status-preview" ||
		activation["runtime_method"] != "GetDesktopActivationStatus" ||
		activation["renderer"] != "xnix-runtime-go desktop-activation-status-preview" ||
		activation["activation_state"] != "ready-for-runtime-commit" ||
		activation["commit_enabled"] != false ||
		activation["launch_enabled"] != false ||
		activation["host_root_modified"] != false ||
		activation["backend_details_exposed"] != false {
		t.Fatalf("unexpected activation status snapshot: %#v", activation)
	}
	execution := payload["execution_readiness_snapshot"].(map[string]any)
	if execution["request_type"] != "execution-readiness-preview" ||
		execution["runtime_method"] != "GetExecutionReadiness" ||
		execution["execution_state"] != "blocked" ||
		execution["overall_status"] != "not-ready" ||
		execution["gate_count"] != float64(5) ||
		execution["desktop_entry_launch_visible"] != true ||
		execution["launch_allowed"] != false ||
		execution["launch_enabled"] != false ||
		execution["execution_request_created"] != false ||
		execution["backend_binding_ready"] != false ||
		execution["host_root_modified"] != false ||
		execution["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution readiness snapshot: %#v", execution)
	}
	launchIntent := payload["launch_intent_snapshot"].(map[string]any)
	if launchIntent["request_type"] != "launch-intent-preview" ||
		launchIntent["intent_type"] != "runtime-launch-intent" ||
		launchIntent["source"] != "desktop-launcher" ||
		launchIntent["runtime_method"] != "Launch" ||
		launchIntent["read_method"] != "GetLaunchIntent" ||
		launchIntent["file_count"] != float64(1) ||
		launchIntent["portal_required"] != true ||
		launchIntent["snapshot_required"] != true ||
		launchIntent["standard_desktop_entry"] != true ||
		launchIntent["launch_uses_runtime"] != true ||
		launchIntent["launch_allowed"] != false ||
		launchIntent["launch_enabled"] != false ||
		launchIntent["execution_request_created"] != false ||
		launchIntent["execution_started"] != false ||
		launchIntent["backend_binding_ready"] != false ||
		launchIntent["request_object_created"] != false ||
		launchIntent["permission_granted"] != false ||
		launchIntent["host_root_modified"] != false ||
		launchIntent["network_required"] != false ||
		launchIntent["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch intent snapshot: %#v", launchIntent)
	}
	window := payload["window_identity_snapshot"].(map[string]any)
	if window["schema_version"] != "xnix.runtime.window_identity.v1" ||
		window["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		window["launcher_url"] != "applications:xnix-org.example.ledger.desktop" ||
		window["window_kind"] != "compatibility-application" ||
		window["class_group"] != "xnix-compatibility" ||
		window["resource_name"] != "org.example.ledger" ||
		window["title_hint"] != "Example Ledger" ||
		window["task_manager_grouping_key"] != "org.example.ledger" ||
		window["task_manager_pinning_allowed"] != true ||
		window["task_manager_restore_allowed"] != true ||
		window["task_manager_skip_taskbar"] != false ||
		window["task_manager_show_in_switcher"] != true ||
		window["prefer_existing_window"] != true ||
		window["kwin_script_role"] != "identity-and-layout" ||
		window["kwin_placement"] != "normal-window" ||
		window["window_manager_policy_only"] != true ||
		window["runtime_owns_backend_policy"] != true ||
		window["task_manager_entry_active"] != false ||
		window["kwin_rule_applied"] != false ||
		window["host_root_modified"] != false ||
		window["backend_details_exposed"] != false {
		t.Fatalf("unexpected window identity snapshot: %#v", window)
	}
	files := payload["file_association_snapshot"].(map[string]any)
	mimeTypes := files["mime_types"].([]any)
	if files["plan_type"] != "file-association-plan" ||
		files["association_type"] != "desktop-file-association" ||
		files["runtime_method"] != "GetFileAssociationPlan" ||
		files["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		files["mimeapps_path"] != "usr/share/applications/mimeapps.list" ||
		files["mime_type_count"] != float64(1) ||
		len(mimeTypes) != 1 ||
		mimeTypes[0] != "application/x-xnix-abc" ||
		files["file_open_command"] != "xnix-compat-open" ||
		files["file_open_argument"] != "%U" ||
		files["standard_mimeapps_list"] != true ||
		files["staged_root_only"] != true ||
		files["overwrite_existing_mimeapps"] != false ||
		files["portal_required_for_file_open"] != true ||
		files["file_association_ready"] != true ||
		files["file_open_preview_available"] != true ||
		files["direct_host_file_access"] != false ||
		files["request_object_created"] != false ||
		files["permission_granted"] != false ||
		files["files_written"] != false ||
		files["mimeapps_written"] != false ||
		files["host_root_modified"] != false ||
		files["backend_details_exposed"] != false {
		t.Fatalf("unexpected file association snapshot: %#v", files)
	}
	tray := payload["tray_status_snapshot"].(map[string]any)
	actions := tray["actions"].([]any)
	if tray["status_type"] != "tray-status-preview" ||
		tray["runtime_method"] != "GetTrayStatus" ||
		tray["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		tray["registered_application_count"] != float64(1) ||
		tray["active_application_count"] != float64(0) ||
		tray["attention_required_count"] != float64(0) ||
		tray["compatibility_state"] != "ready" ||
		tray["compatibility_label"] != "Ready" ||
		tray["tray_bridge_state"] != "planned" ||
		tray["tray_bridge_label"] != "Tray bridge is planned" ||
		tray["bridged_tray_application_count"] != float64(0) ||
		tray["action_count"] != float64(2) ||
		len(actions) != 2 ||
		actions[0] != "open-compatibility-center" ||
		actions[1] != "open-settings" ||
		tray["user_visible"] != true ||
		tray["tray_status_ready"] != true ||
		tray["live_backend_bridge_enabled"] != false ||
		tray["bridge_configuration_persisted"] != false ||
		tray["host_root_modified"] != false ||
		tray["backend_details_exposed"] != false {
		t.Fatalf("unexpected tray status snapshot: %#v", tray)
	}
	notification := payload["notification_snapshot"].(map[string]any)
	notificationActions := notification["actions"].([]any)
	if notification["request_type"] != "desktop-notification-preview" ||
		notification["runtime_method"] != "GetNotificationPlan" ||
		notification["source"] != "runtime-event" ||
		notification["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		notification["event_type"] != "approval-required" ||
		notification["notification_id"] != "org.example.ledger.approval-required" ||
		notification["urgency"] != "critical" ||
		notification["category"] != "compatibility.approval" ||
		notification["title"] != "Example Ledger needs approval" ||
		notification["action_count"] != float64(2) ||
		len(notificationActions) != 2 ||
		notificationActions[0] != "open-compatibility-center" ||
		notificationActions[1] != "review-request" ||
		notification["requires_user_review"] != true ||
		notification["user_visible"] != true ||
		notification["notification_ready"] != true ||
		notification["action_execution_enabled"] != false ||
		notification["repair_execution_enabled"] != false ||
		notification["settings_persistence_enabled"] != false ||
		notification["notifications_sent"] != false ||
		notification["host_root_modified"] != false ||
		notification["backend_details_exposed"] != false {
		t.Fatalf("unexpected notification snapshot: %#v", notification)
	}
	if payload["navigation_count"] != float64(12) ||
		payload["primary_navigation_target"] != "compatibility-center-gates" ||
		navigation[0].(map[string]any)["id"] != "overview" ||
		navigation[1].(map[string]any)["id"] != "backend" ||
		navigation[2].(map[string]any)["id"] != "activation" ||
		navigation[3].(map[string]any)["id"] != "execution" ||
		navigation[4].(map[string]any)["id"] != "launch" ||
		navigation[5].(map[string]any)["id"] != "window" ||
		navigation[6].(map[string]any)["id"] != "files" ||
		navigation[7].(map[string]any)["id"] != "tray" ||
		navigation[8].(map[string]any)["id"] != "notifications" ||
		navigation[9].(map[string]any)["id"] != "actions" ||
		navigation[10].(map[string]any)["id"] != "settings" ||
		navigation[11].(map[string]any)["id"] != "diagnostics" {
		t.Fatalf("unexpected navigation: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["official_desktop_only"] != true ||
		payload["user_visible"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["page_preview_created"] != true ||
		payload["page_persisted"] != false ||
		payload["settings_persisted"] != false ||
		payload["settings_persistence_enabled"] != false ||
		payload["notifications_sent"] != false ||
		payload["resource_grant_created"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_grant_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected center page safety flags: %#v", payload)
	}

	output.Reset()
	err = run([]string{"kde-center-page-section-detail-preview", "--registry", registryPath, "--app", "org.example.ledger", "--section", "diagnostics", "--decision", "approved", "file:///home/test/Documents/example.abc"}, &output)
	if err != nil {
		t.Fatalf("diagnostics section detail returned error: %v", err)
	}
	var detail map[string]any
	if err := json.Unmarshal(output.Bytes(), &detail); err != nil {
		t.Fatalf("Unmarshal diagnostics detail returned error: %v", err)
	}
	route := detail["diagnostic_history_route"].(map[string]any)
	if detail["section_id"] != "diagnostics" ||
		detail["section_runtime_method"] != "GetDiagnostics" ||
		detail["section_read_model"] != "diagnostic-history-preview" ||
		route["request_type"] != "diagnostic-history-route" ||
		route["runtime_method"] != "GetDiagnostics" ||
		route["read_method"] != "GetDiagnosticHistoryPreview" ||
		route["read_model"] != "diagnostic-history-preview" ||
		route["cli_command"] != "diagnostic-history-preview" ||
		route["application_id"] != "org.example.ledger" ||
		route["state_root_required"] != true ||
		route["state_root_path_exposed"] != false ||
		route["history_preview_created"] != false ||
		route["ai_provider_call_enabled"] != false ||
		route["file_content_read"] != false ||
		route["file_paths_exposed"] != false ||
		route["request_object_created"] != false ||
		route["permission_granted"] != false ||
		route["launch_enabled"] != false ||
		route["execution_started"] != false ||
		route["repair_execution_enabled"] != false ||
		route["host_root_modified"] != false ||
		route["backend_details_exposed"] != false {
		t.Fatalf("unexpected diagnostics history route: %#v", detail)
	}
}
