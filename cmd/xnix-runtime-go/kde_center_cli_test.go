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
		payload["source"] != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+kde-action-card-deck-preview+settings-preview" ||
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
	if payload["navigation_count"] != float64(7) ||
		payload["primary_navigation_target"] != "compatibility-center-gates" ||
		navigation[0].(map[string]any)["id"] != "overview" ||
		navigation[1].(map[string]any)["id"] != "backend" ||
		navigation[2].(map[string]any)["id"] != "activation" ||
		navigation[3].(map[string]any)["id"] != "execution" ||
		navigation[4].(map[string]any)["id"] != "actions" ||
		navigation[5].(map[string]any)["id"] != "settings" ||
		navigation[6].(map[string]any)["id"] != "diagnostics" {
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
}
