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

func TestBackendSelectionPreviewCommandRendersProfiles(t *testing.T) {
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
	err := run([]string{"backend-selection-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.backend_selection.v1" ||
		payload["request_type"] != "backend-selection-preview" ||
		payload["plan_type"] != "compatibility-backend-selection-plan" ||
		payload["source"] != "compatibility-center" ||
		payload["runtime_method"] != "GetBackendSelectionPlan" {
		t.Fatalf("unexpected backend selection schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["display_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected backend selection identity: %#v", payload)
	}
	if payload["recommended_profile_id"] != "local-compatibility" ||
		payload["candidate_count"] != float64(2) ||
		payload["ready_candidate_count"] != float64(0) ||
		payload["blocked_candidate_count"] != float64(2) {
		t.Fatalf("unexpected backend selection recommendation: %#v", payload)
	}
	profiles := payload["candidate_profiles"].([]any)
	first := profiles[0].(map[string]any)
	second := profiles[1].(map[string]any)
	if first["id"] != "local-compatibility" ||
		first["selection_state"] != "recommended" ||
		first["recommended"] != true ||
		first["ready"] != false ||
		first["blocked"] != true ||
		first["selection_committed"] != false ||
		first["environment_created"] != false ||
		first["backend_process_started"] != false ||
		first["backend_details_exposed"] != false {
		t.Fatalf("unexpected local compatibility profile: %#v", first)
	}
	if second["id"] != "isolated-compatibility" || second["recommended"] != false {
		t.Fatalf("unexpected isolated compatibility profile: %#v", second)
	}
	reviews := payload["required_reviews"].([]any)
	if len(reviews) != 5 ||
		reviews[0] != "backend-capability-review" ||
		reviews[1] != "backend-binding-review" ||
		reviews[2] != "application-state-root-review" ||
		reviews[3] != "portal-policy-review" ||
		reviews[4] != "snapshot-baseline-review" {
		t.Fatalf("unexpected backend selection reviews: %#v", reviews)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false || payload["user_visible"] != true ||
		payload["selection_committed"] != false ||
		payload["selection_change_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["capability_activation_enabled"] != false ||
		payload["environment_created"] != false ||
		payload["request_object_created"] != false ||
		payload["state_root_created"] != false ||
		payload["snapshot_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected backend selection safety flags: %#v", payload)
	}
}

func TestExecutionReadinessPreviewCommandKeepsLaunchGated(t *testing.T) {
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
	err := run([]string{"execution-readiness-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.launch_readiness.v1" ||
		payload["request_type"] != "execution-readiness-preview" ||
		payload["readiness_type"] != "compatibility-execution-readiness" ||
		payload["source"] != "compatibility-center" ||
		payload["runtime_method"] != "GetExecutionReadiness" {
		t.Fatalf("unexpected execution readiness schema: %#v", payload)
	}
	application := payload["application"].(map[string]any)
	if application["id"] != "org.example.ledger" ||
		application["name"] != "Example Ledger" ||
		application["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution readiness application: %#v", application)
	}
	profile := payload["compatibility_profile"].(map[string]any)
	if profile["id"] != "local-compatibility" ||
		profile["kind"] != "local" ||
		profile["ready"] != false ||
		profile["launch_enabled"] != false ||
		profile["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution readiness profile: %#v", profile)
	}
	if payload["execution_state"] != "blocked" ||
		payload["overall_status"] != "not-ready" ||
		payload["gate_count"] != float64(5) ||
		payload["required_gate_count"] != float64(2) ||
		payload["pending_gate_count"] != float64(1) ||
		payload["blocked_gate_count"] != float64(1) {
		t.Fatalf("unexpected execution readiness state: %#v", payload)
	}
	gates := payload["gates"].([]any)
	if len(gates) != 5 ||
		gates[0].(map[string]any)["id"] != "recipe-validation" ||
		gates[1].(map[string]any)["id"] != "portal-policy-review" ||
		gates[2].(map[string]any)["id"] != "snapshot-baseline" ||
		gates[3].(map[string]any)["id"] != "backend-binding" ||
		gates[4].(map[string]any)["id"] != "runtime-launch-write-gate" {
		t.Fatalf("unexpected execution readiness gates: %#v", gates)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["portal_policy_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["user_action_required"] != true ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution readiness safety flags: %#v", payload)
	}
}

func TestLaunchIntentPreviewCommandCapturesDesktopLaunch(t *testing.T) {
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
	err := run([]string{"launch-intent-preview", "--registry", registryPath, "--app", "org.example.ledger", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.launch_intent.v1" ||
		payload["request_type"] != "launch-intent-preview" ||
		payload["intent_type"] != "runtime-launch-intent" ||
		payload["source"] != "desktop-launcher" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetLaunchIntent" {
		t.Fatalf("unexpected launch intent schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected launch intent identity: %#v", payload)
	}
	profile := payload["compatibility_profile"].(map[string]any)
	if profile["id"] != "local-compatibility" ||
		profile["ready"] != false ||
		profile["launch_enabled"] != false ||
		profile["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch intent profile: %#v", profile)
	}
	runPlan := payload["run_plan"].(map[string]any)
	if runPlan["plan_type"] != "compatibility-run" ||
		runPlan["strategy"] != "automatic-managed" ||
		runPlan["backend_details_exposed"] != false ||
		runPlan["backend_ready"] != false ||
		runPlan["portal_policy_required"] != true ||
		runPlan["snapshot_before_risky_change"] != true ||
		runPlan["runtime_write_gate_required"] != true ||
		runPlan["execution_request_created"] != false {
		t.Fatalf("unexpected launch intent run plan: %#v", runPlan)
	}
	if payload["execution_state"] != "blocked" ||
		payload["overall_status"] != "not-ready" ||
		payload["write_gate_decision"] != "blocked-until-production-backend" ||
		payload["denial_error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected launch intent gate state: %#v", payload)
	}
	if payload["portal_required"] != true ||
		payload["file_count"] != float64(1) ||
		payload["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" {
		t.Fatalf("unexpected launch intent files: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["standard_desktop_entry"] != true ||
		payload["launch_uses_runtime"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch intent safety flags: %#v", payload)
	}
}

func TestExecutionRequestPreviewCommandKeepsRequestUncreated(t *testing.T) {
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
	err := run([]string{"execution-request-preview", "--registry", registryPath, "--app", "org.example.ledger", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.request_intake.v1" ||
		payload["request_type"] != "execution-request-preview" ||
		payload["intent_type"] != "runtime-launch-intent" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "runtime-launch-intent" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionRequestPreview" {
		t.Fatalf("unexpected execution request schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution request identity: %#v", payload)
	}
	launchIntent := payload["launch_intent"].(map[string]any)
	if launchIntent["source"] != "desktop-launcher" ||
		launchIntent["intent_type"] != "runtime-launch-intent" ||
		launchIntent["runtime_method"] != "Launch" ||
		launchIntent["read_method"] != "GetLaunchIntent" ||
		launchIntent["launch_allowed"] != false ||
		launchIntent["launch_enabled"] != false ||
		launchIntent["request_object_created"] != false ||
		launchIntent["execution_started"] != false ||
		launchIntent["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch intent summary: %#v", launchIntent)
	}
	profile := payload["compatibility_profile"].(map[string]any)
	if profile["id"] != "local-compatibility" ||
		profile["kind"] != "local" ||
		profile["ready"] != false ||
		profile["launch_enabled"] != false ||
		profile["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution request profile: %#v", profile)
	}
	gateSummary := payload["gate_summary"].(map[string]any)
	if gateSummary["gate_count"] != float64(5) ||
		gateSummary["required_gate_count"] != float64(2) ||
		gateSummary["pending_gate_count"] != float64(1) ||
		gateSummary["blocked_gate_count"] != float64(1) {
		t.Fatalf("unexpected execution request gate summary: %#v", gateSummary)
	}
	if payload["execution_state"] != "blocked" ||
		payload["overall_status"] != "not-ready" ||
		payload["write_gate_decision"] != "blocked-until-production-backend" ||
		payload["denial_error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected execution request gate state: %#v", payload)
	}
	if payload["portal_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["file_count"] != float64(1) ||
		payload["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" ||
		launchIntent["portal_required"] != true ||
		launchIntent["file_count"] != float64(1) ||
		launchIntent["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" {
		t.Fatalf("unexpected execution request file routing: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["execution_request_persisted"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution request safety flags: %#v", payload)
	}
}

func TestExecutionReviewPreviewCommandRendersCompatibilityCenterCard(t *testing.T) {
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
	err := run([]string{"execution-review-preview", "--registry", registryPath, "--app", "org.example.ledger", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.request_review.v1" ||
		payload["request_type"] != "execution-review-preview" ||
		payload["review_type"] != "compatibility-center-launch-review" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-request-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionReviewPreview" {
		t.Fatalf("unexpected execution review schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution review identity: %#v", payload)
	}
	request := payload["execution_request"].(map[string]any)
	if request["schema_version"] != "xnix.runtime.request_intake.v1" ||
		request["request_type"] != "execution-request-preview" ||
		request["request_state"] != "blocked" ||
		request["source"] != "runtime-launch-intent" ||
		request["read_method"] != "GetExecutionRequestPreview" ||
		request["portal_required"] != true ||
		request["snapshot_required"] != true ||
		request["file_count"] != float64(1) ||
		request["launch_intent_captured"] != true ||
		request["execution_request_created"] != false ||
		request["execution_request_persisted"] != false ||
		request["execution_started"] != false ||
		request["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution request summary: %#v", request)
	}
	card := payload["review_card"].(map[string]any)
	if card["id"] != "org.example.ledger:launch-review" ||
		card["card_type"] != "compatibility-center-launch-review" ||
		card["status"] != "blocked" ||
		card["severity"] != "requires-runtime-gates" ||
		card["user_review_required"] != true ||
		card["runtime_approval_needed"] != true ||
		card["primary_action"] != "Open Compatibility Center" {
		t.Fatalf("unexpected execution review card: %#v", card)
	}
	queue := payload["action_queue"].(map[string]any)
	if queue["queue_type"] != "compatibility-center-request-review-queue" ||
		queue["action_count"] != float64(1) ||
		queue["pending_action_count"] != float64(1) ||
		queue["user_review_required_count"] != float64(1) ||
		queue["execution_enabled"] != false ||
		queue["queue_persisted"] != false ||
		queue["review_receipt_recorded"] != false ||
		queue["actions"].([]any)[0] != "review-launch-request" {
		t.Fatalf("unexpected execution review queue: %#v", queue)
	}
	profile := payload["compatibility_profile"].(map[string]any)
	if profile["id"] != "local-compatibility" ||
		profile["kind"] != "local" ||
		profile["ready"] != false ||
		profile["launch_enabled"] != false ||
		profile["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution review profile: %#v", profile)
	}
	gateSummary := payload["gate_summary"].(map[string]any)
	if gateSummary["gate_count"] != float64(5) ||
		gateSummary["required_gate_count"] != float64(2) ||
		gateSummary["pending_gate_count"] != float64(1) ||
		gateSummary["blocked_gate_count"] != float64(1) {
		t.Fatalf("unexpected execution review gate summary: %#v", gateSummary)
	}
	if payload["portal_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["file_count"] != float64(1) ||
		payload["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" {
		t.Fatalf("unexpected execution review file routing: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["action_queue_candidate"] != true ||
		payload["review_receipt_required"] != true ||
		payload["review_receipt_recorded"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["execution_request_persisted"] != false ||
		payload["action_queue_persisted"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution review safety flags: %#v", payload)
	}
}

func TestExecutionDecisionPreviewCommandCapturesDecisionWithoutApproval(t *testing.T) {
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
	err := run([]string{"execution-decision-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.request_decision.v1" ||
		payload["request_type"] != "execution-decision-preview" ||
		payload["decision_type"] != "compatibility-center-launch-decision" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-review-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionDecisionPreview" {
		t.Fatalf("unexpected execution decision schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution decision identity: %#v", payload)
	}
	review := payload["execution_review"].(map[string]any)
	if review["schema_version"] != "xnix.runtime.request_review.v1" ||
		review["request_type"] != "execution-review-preview" ||
		review["review_type"] != "compatibility-center-launch-review" ||
		review["source"] != "execution-request-preview" ||
		review["read_method"] != "GetExecutionReviewPreview" ||
		review["action_queue_candidate"] != true ||
		review["review_receipt_required"] != true ||
		review["review_receipt_recorded"] != false ||
		review["action_queue_persisted"] != false ||
		review["execution_started"] != false ||
		review["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution review summary: %#v", review)
	}
	decision := payload["decision"].(map[string]any)
	if decision["decision"] != "approved" ||
		decision["decision_accepted"] != true ||
		decision["user_intent_captured"] != true ||
		decision["decision_recorded"] != false ||
		decision["runtime_approval_granted"] != false ||
		decision["execution_allowed"] != false ||
		decision["review_receipt_created"] != false ||
		decision["queue_state_changed"] != false ||
		decision["permission_grant_created"] != false ||
		decision["desktop_notification_intent"] != "show-launch-approval-pending" {
		t.Fatalf("unexpected decision payload: %#v", decision)
	}
	queue := payload["action_queue"].(map[string]any)
	if queue["execution_enabled"] != false ||
		queue["queue_persisted"] != false ||
		queue["review_receipt_recorded"] != false {
		t.Fatalf("unexpected action queue: %#v", queue)
	}
	if payload["portal_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["file_count"] != float64(1) ||
		payload["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" {
		t.Fatalf("unexpected execution decision file routing: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["action_queue_candidate"] != true ||
		payload["user_decision_captured"] != true ||
		payload["review_receipt_required"] != true ||
		payload["review_receipt_recorded"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_request_created"] != false ||
		payload["execution_request_persisted"] != false ||
		payload["action_queue_persisted"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution decision safety flags: %#v", payload)
	}
}

func TestExecutionPreflightPreviewCommandBlocksLaunchUntilGatesPass(t *testing.T) {
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
	err := run([]string{"execution-preflight-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.launch_preflight.v1" ||
		payload["request_type"] != "execution-preflight-preview" ||
		payload["preflight_type"] != "compatibility-launch-preflight" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-decision-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionPreflightPreview" {
		t.Fatalf("unexpected execution preflight schema: %#v", payload)
	}
	decision := payload["execution_decision"].(map[string]any)
	if decision["schema_version"] != "xnix.runtime.request_decision.v1" ||
		decision["request_type"] != "execution-decision-preview" ||
		decision["decision"] != "approved" ||
		decision["decision_accepted"] != true ||
		decision["user_intent_captured"] != true ||
		decision["decision_recorded"] != false ||
		decision["runtime_approval_granted"] != false ||
		decision["execution_allowed"] != false ||
		decision["review_receipt_created"] != false ||
		decision["queue_state_changed"] != false {
		t.Fatalf("unexpected execution decision summary: %#v", decision)
	}
	if payload["check_count"] != float64(5) ||
		payload["passed_check_count"] != float64(1) ||
		payload["required_check_count"] != float64(2) ||
		payload["pending_check_count"] != float64(1) ||
		payload["blocked_check_count"] != float64(1) {
		t.Fatalf("unexpected execution preflight counts: %#v", payload)
	}
	checks := payload["preflight_checks"].([]any)
	if len(checks) != 5 ||
		checks[0].(map[string]any)["id"] != "user-decision" ||
		checks[0].(map[string]any)["status"] != "pass" ||
		checks[1].(map[string]any)["id"] != "portal-policy-review" ||
		checks[1].(map[string]any)["status"] != "required" ||
		checks[2].(map[string]any)["id"] != "snapshot-baseline" ||
		checks[3].(map[string]any)["id"] != "backend-binding" ||
		checks[4].(map[string]any)["id"] != "runtime-launch-write-gate" ||
		checks[4].(map[string]any)["status"] != "blocked" {
		t.Fatalf("unexpected execution preflight checks: %#v", checks)
	}
	if payload["portal_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["file_count"] != float64(1) ||
		payload["file_uris"].([]any)[0] != "file:///home/test/Documents/book.abc" {
		t.Fatalf("unexpected execution preflight file routing: %#v", payload)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["preflight_complete"] != false ||
		payload["preflight_passed"] != false ||
		payload["portal_preflight_ready"] != false ||
		payload["snapshot_preflight_ready"] != false ||
		payload["backend_preflight_ready"] != false ||
		payload["write_gate_open"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_binding_ready"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution preflight safety flags: %#v", payload)
	}
}

func TestExecutionResourceGrantPreviewCommandKeepsPermissionsUngrantable(t *testing.T) {
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
	err := run([]string{"execution-resource-grant-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.resource_grant.v1" ||
		payload["request_type"] != "execution-resource-grant-preview" ||
		payload["grant_type"] != "compatibility-launch-resource-grant" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-preflight-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionResourceGrantPreview" {
		t.Fatalf("unexpected execution resource grant schema: %#v", payload)
	}
	preflight := payload["execution_preflight"].(map[string]any)
	if preflight["schema_version"] != "xnix.runtime.launch_preflight.v1" ||
		preflight["request_type"] != "execution-preflight-preview" ||
		preflight["user_decision_allows_launch"] != true ||
		preflight["preflight_passed"] != false ||
		preflight["portal_required"] != true ||
		preflight["write_gate_open"] != false ||
		preflight["runtime_launch_approval"] != false ||
		preflight["permission_granted"] != false {
		t.Fatalf("unexpected preflight summary: %#v", preflight)
	}
	if payload["grant_count"] != float64(7) ||
		payload["allow_count"] != float64(1) ||
		payload["ask_count"] != float64(5) ||
		payload["deny_count"] != float64(1) ||
		payload["portal_required_count"] != float64(6) ||
		payload["pending_review_count"] != float64(5) {
		t.Fatalf("unexpected grant counts: %#v", payload)
	}
	grants := payload["resource_grants"].([]any)
	if len(grants) != 7 {
		t.Fatalf("unexpected resource grants: %#v", grants)
	}
	grantByID := map[string]map[string]any{}
	for _, rawGrant := range grants {
		grant := rawGrant.(map[string]any)
		grantByID[grant["id"].(string)] = grant
		if grant["request_object_created"] != false ||
			grant["permission_granted"] != false ||
			grant["direct_access_allowed"] != false ||
			grant["backend_details_exposed"] != false {
			t.Fatalf("grant enabled unsafe access: %#v", grant)
		}
	}
	if grantByID["documents"]["grant_state"] != "requires-review" ||
		grantByID["downloads"]["grant_state"] != "requires-review" ||
		grantByID["camera"]["grant_state"] != "planned-deny" ||
		grantByID["network"]["grant_state"] != "planned-allow" ||
		grantByID["network"]["portal_required"] != false {
		t.Fatalf("unexpected resource grant states: %#v", grantByID)
	}
	bridge := payload["bridge_summary"].(map[string]any)
	if bridge["request_type"] != "desktop-resource-bridge-preview" ||
		bridge["plan_type"] != "desktop-resource-bridge-plan" ||
		bridge["bridge_state"] != "planned" ||
		bridge["resource_count"] != float64(5) ||
		bridge["portal_mediated"] != true ||
		bridge["resource_bridges_enabled"] != false ||
		bridge["request_objects_created"] != false ||
		bridge["direct_host_file_access"] != false {
		t.Fatalf("unexpected bridge summary: %#v", bridge)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["preflight_passed"] != false ||
		payload["portal_review_required"] != true ||
		payload["snapshot_required"] != true ||
		payload["grant_plan_created"] != true ||
		payload["grant_objects_created"] != false ||
		payload["request_objects_created"] != false ||
		payload["permission_granted"] != false ||
		payload["resource_bridges_enabled"] != false ||
		payload["settings_persisted"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_permission_changed"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution resource grant safety flags: %#v", payload)
	}
}

func TestExecutionTransactionPreviewCommandBlocksCommitAndLaunch(t *testing.T) {
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
	err := run([]string{"execution-transaction-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.launch_transaction.v1" ||
		payload["request_type"] != "execution-transaction-preview" ||
		payload["transaction_type"] != "compatibility-launch-transaction" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-resource-grant-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionTransactionPreview" {
		t.Fatalf("unexpected execution transaction schema: %#v", payload)
	}
	resourceGrant := payload["resource_grant"].(map[string]any)
	if resourceGrant["request_type"] != "execution-resource-grant-preview" ||
		resourceGrant["grant_type"] != "compatibility-launch-resource-grant" ||
		resourceGrant["grant_count"] != float64(7) ||
		resourceGrant["portal_required_count"] != float64(6) ||
		resourceGrant["pending_review_count"] != float64(5) ||
		resourceGrant["grant_plan_created"] != true ||
		resourceGrant["grant_objects_created"] != false ||
		resourceGrant["permission_granted"] != false ||
		resourceGrant["resource_bridges_enabled"] != false {
		t.Fatalf("unexpected transaction resource grant: %#v", resourceGrant)
	}
	readiness := payload["readiness"].(map[string]any)
	if readiness["request_type"] != "execution-readiness-preview" ||
		readiness["readiness_type"] != "compatibility-execution-readiness" ||
		readiness["execution_state"] != "blocked" ||
		readiness["overall_status"] != "not-ready" ||
		readiness["gate_count"] != float64(5) ||
		readiness["launch_allowed"] != false ||
		readiness["launch_enabled"] != false ||
		readiness["backend_binding_ready"] != false {
		t.Fatalf("unexpected transaction readiness: %#v", readiness)
	}
	backendBinding := payload["backend_binding"].(map[string]any)
	if backendBinding["recommended_profile_id"] != "local-compatibility" ||
		backendBinding["candidate_count"] != float64(2) ||
		backendBinding["ready_candidate_count"] != float64(0) ||
		backendBinding["blocked_candidate_count"] != float64(2) ||
		backendBinding["binding_required"] != true ||
		backendBinding["binding_committed"] != false ||
		backendBinding["environment_created"] != false ||
		backendBinding["backend_launch_enabled"] != false {
		t.Fatalf("unexpected backend binding: %#v", backendBinding)
	}
	snapshot := payload["snapshot_baseline"].(map[string]any)
	if snapshot["required"] != true ||
		snapshot["state"] != "required" ||
		snapshot["baseline_created"] != false ||
		snapshot["restore_point_available"] != false ||
		snapshot["user_documents_included"] != false ||
		snapshot["host_system_included"] != false {
		t.Fatalf("unexpected snapshot baseline: %#v", snapshot)
	}
	writeGate := payload["write_gate"].(map[string]any)
	if writeGate["method_name"] != "Launch" ||
		writeGate["gate_decision"] != "blocked-until-production-backend" ||
		writeGate["denial_error_name"] != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		writeGate["write_method_enabled"] != false ||
		writeGate["dispatch_enabled"] != false ||
		writeGate["request_object_created"] != false {
		t.Fatalf("unexpected write gate: %#v", writeGate)
	}
	if payload["step_count"] != float64(7) ||
		payload["passed_step_count"] != float64(2) ||
		payload["required_step_count"] != float64(2) ||
		payload["pending_step_count"] != float64(2) ||
		payload["blocked_step_count"] != float64(1) {
		t.Fatalf("unexpected transaction step counts: %#v", payload)
	}
	steps := payload["transaction_steps"].([]any)
	if len(steps) != 7 ||
		steps[0].(map[string]any)["id"] != "identity-validation" ||
		steps[1].(map[string]any)["id"] != "user-decision" ||
		steps[1].(map[string]any)["status"] != "pass" ||
		steps[2].(map[string]any)["id"] != "portal-resource-review" ||
		steps[3].(map[string]any)["id"] != "resource-grant-objects" ||
		steps[4].(map[string]any)["id"] != "snapshot-baseline" ||
		steps[5].(map[string]any)["id"] != "backend-binding" ||
		steps[6].(map[string]any)["id"] != "runtime-launch-write-gate" ||
		steps[6].(map[string]any)["status"] != "blocked" {
		t.Fatalf("unexpected transaction steps: %#v", steps)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["transaction_plan_created"] != true ||
		payload["transaction_committed"] != false ||
		payload["readiness_passed"] != false ||
		payload["resource_grants_committed"] != false ||
		payload["portal_approval_recorded"] != false ||
		payload["snapshot_baseline_created"] != false ||
		payload["backend_binding_committed"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_permission_changed"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution transaction safety flags: %#v", payload)
	}
}

func TestExecutionSessionPreviewCommandPlansDesktopIdentityWithoutStarting(t *testing.T) {
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
	err := run([]string{"execution-session-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.session_identity.v1" ||
		payload["request_type"] != "execution-session-preview" ||
		payload["session_type"] != "compatibility-execution-session" ||
		payload["request_state"] != "blocked" ||
		payload["source"] != "execution-transaction-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionSessionPreview" {
		t.Fatalf("unexpected execution session schema: %#v", payload)
	}
	transaction := payload["transaction"].(map[string]any)
	if transaction["request_type"] != "execution-transaction-preview" ||
		transaction["transaction_type"] != "compatibility-launch-transaction" ||
		transaction["step_count"] != float64(7) ||
		transaction["blocked_step_count"] != float64(1) ||
		transaction["transaction_committed"] != false ||
		transaction["runtime_launch_approval"] != false ||
		transaction["launch_allowed"] != false ||
		transaction["execution_started"] != false {
		t.Fatalf("unexpected session transaction summary: %#v", transaction)
	}
	window := payload["window_identity"].(map[string]any)
	if window["schema_version"] != "xnix.runtime.window_identity.v1" ||
		window["window_kind"] != "compatibility-application" ||
		window["class_group"] != "xnix-compatibility" ||
		window["resource_name"] != "org.example.ledger" ||
		window["launcher_url"] != "applications:xnix-org.example.ledger.desktop" ||
		window["title_hint"] != "Example Ledger" ||
		window["window_observed"] != false ||
		window["window_registration"] != "planned" ||
		window["backend_details_exposed"] != false {
		t.Fatalf("unexpected window identity summary: %#v", window)
	}
	taskManager := payload["task_manager"].(map[string]any)
	if taskManager["grouping_key"] != "org.example.ledger" ||
		taskManager["pinning_allowed"] != true ||
		taskManager["restore_allowed"] != true ||
		taskManager["skip_taskbar"] != false ||
		taskManager["show_in_switcher"] != true ||
		taskManager["prefer_existing_window"] != true ||
		taskManager["entry_planned"] != true ||
		taskManager["entry_active"] != false {
		t.Fatalf("unexpected task manager summary: %#v", taskManager)
	}
	kwin := payload["kwin"].(map[string]any)
	if kwin["script_role"] != "identity-and-layout" ||
		kwin["resource_name"] != "org.example.ledger" ||
		kwin["class_group"] != "xnix-compatibility" ||
		kwin["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		kwin["task_manager_grouping_key"] != "org.example.ledger" ||
		kwin["launcher_url"] != "applications:xnix-org.example.ledger.desktop" ||
		kwin["placement"] != "normal-window" ||
		kwin["window_manager_policy_only"] != true ||
		kwin["rule_planned"] != true ||
		kwin["rule_applied"] != false {
		t.Fatalf("unexpected KWin summary: %#v", kwin)
	}
	tray := payload["tray"].(map[string]any)
	if tray["status_type"] != "tray-status-preview" ||
		tray["registered_application_count"] != float64(1) ||
		tray["active_application_count"] != float64(0) ||
		tray["attention_required_count"] != float64(0) ||
		tray["compatibility_state"] != "ready" ||
		tray["tray_bridge_state"] != "planned" ||
		tray["entry_planned"] != true ||
		tray["live_bridge_enabled"] != false ||
		tray["bridge_configuration_persisted"] != false ||
		tray["backend_details_exposed"] != false {
		t.Fatalf("unexpected tray summary: %#v", tray)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["session_plan_created"] != true ||
		payload["session_created"] != false ||
		payload["session_registered"] != false ||
		payload["window_observed"] != false ||
		payload["task_manager_entry_planned"] != true ||
		payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_planned"] != true ||
		payload["kwin_rule_applied"] != false ||
		payload["tray_entry_planned"] != true ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["transaction_committed"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution session safety flags: %#v", payload)
	}
}

func TestExecutionSessionStatusPreviewCommandKeepsLiveStateReadOnly(t *testing.T) {
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
	err := run([]string{"execution-session-status-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.session_status.v1" ||
		payload["request_type"] != "execution-session-status-preview" ||
		payload["status_type"] != "compatibility-session-status" ||
		payload["session_type"] != "compatibility-execution-session" ||
		payload["request_state"] != "blocked" ||
		payload["session_state"] != "planned-blocked" ||
		payload["source"] != "execution-session-preview" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetExecutionSessionStatusPreview" {
		t.Fatalf("unexpected execution session status schema: %#v", payload)
	}
	session := payload["session"].(map[string]any)
	if session["session_state"] != "planned-blocked" ||
		session["transaction_state"] != "blocked" ||
		session["transaction_step_count"] != float64(7) ||
		session["blocked_transaction_steps"] != float64(1) ||
		session["window_registration"] != "planned" ||
		session["desktop_surface_state"] != "planned" ||
		session["compatibility_center_state"] != "waiting-for-runtime-gates" ||
		session["task_manager_state"] != "planned" ||
		session["tray_state"] != "planned" ||
		session["session_plan_created"] != true ||
		session["session_created"] != false ||
		session["session_registered"] != false ||
		session["session_active"] != false ||
		session["live_state_observed"] != false ||
		session["status_persisted"] != false ||
		session["runtime_launch_approval"] != false ||
		session["launch_allowed"] != false ||
		session["execution_started"] != false ||
		session["backend_process_started"] != false ||
		session["backend_details_exposed"] != false {
		t.Fatalf("unexpected session status summary: %#v", session)
	}
	desktopSurface := payload["desktop_surface"].(map[string]any)
	if desktopSurface["window_kind"] != "compatibility-application" ||
		desktopSurface["class_group"] != "xnix-compatibility" ||
		desktopSurface["resource_name"] != "org.example.ledger" ||
		desktopSurface["launcher_url"] != "applications:xnix-org.example.ledger.desktop" ||
		desktopSurface["window_state"] != "not-observed" ||
		desktopSurface["task_manager_grouping_key"] != "org.example.ledger" ||
		desktopSurface["task_manager_state"] != "planned" ||
		desktopSurface["kwin_state"] != "planned" ||
		desktopSurface["tray_state"] != "planned" ||
		desktopSurface["compatibility_center_state"] != "waiting-for-runtime-gates" ||
		desktopSurface["window_observed"] != false ||
		desktopSurface["task_manager_entry_planned"] != true ||
		desktopSurface["task_manager_entry_active"] != false ||
		desktopSurface["kwin_rule_planned"] != true ||
		desktopSurface["kwin_rule_applied"] != false ||
		desktopSurface["tray_entry_planned"] != true ||
		desktopSurface["live_tray_bridge_enabled"] != false {
		t.Fatalf("unexpected desktop surface status: %#v", desktopSurface)
	}
	userState := payload["user_visible_state"].(map[string]any)
	if userState["primary_label"] != "Example Ledger" ||
		userState["secondary_label"] != "Waiting for Runtime gates" ||
		userState["taskbar_badge"] != "Planned" ||
		userState["tray_label"] != "Ready for review" ||
		userState["compatibility_center_status"] != "Runtime gates required" ||
		userState["next_user_action"] != "Open Compatibility Center" ||
		userState["user_facing_mode"] != "Automatic" ||
		userState["user_facing_access"] != "Review required" ||
		userState["backend_details_exposed"] != false {
		t.Fatalf("unexpected user-visible status: %#v", userState)
	}
	gates := payload["gates"].([]any)
	if len(gates) != 5 ||
		payload["gate_count"] != float64(5) ||
		payload["passed_gate_count"] != float64(2) ||
		payload["required_gate_count"] != float64(1) ||
		payload["pending_gate_count"] != float64(2) ||
		payload["blocked_gate_count"] != float64(1) {
		t.Fatalf("unexpected gate counts: %#v", payload)
	}
	runtimeGate := gates[2].(map[string]any)
	if runtimeGate["id"] != "runtime-launch-approval" ||
		runtimeGate["status"] != "blocked" ||
		runtimeGate["required"] != true ||
		runtimeGate["blocks_live_session"] != true ||
		runtimeGate["runtime_gate_required"] != true {
		t.Fatalf("unexpected runtime launch gate: %#v", runtimeGate)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["launch_intent_captured"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["status_read_model_created"] != true ||
		payload["session_plan_created"] != true ||
		payload["session_created"] != false ||
		payload["session_registered"] != false ||
		payload["session_active"] != false ||
		payload["live_state_observed"] != false ||
		payload["status_persisted"] != false ||
		payload["window_observed"] != false ||
		payload["task_manager_entry_planned"] != true ||
		payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_planned"] != true ||
		payload["kwin_rule_applied"] != false ||
		payload["tray_entry_planned"] != true ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["transaction_committed"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected execution session status safety flags: %#v", payload)
	}
}

func TestKDEEntryPointsPreviewCommandCoversFirstReleaseSurface(t *testing.T) {
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
	err := run([]string{"kde-entrypoints-preview", "--registry", registryPath, "--app", "org.example.ledger", "--decision", "approved", "file:///home/test/Documents/book.abc"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.kde_entrypoints.v1" ||
		payload["request_type"] != "kde-entrypoints-preview" ||
		payload["surface_type"] != "kde-first-release-entrypoints" ||
		payload["source"] != "execution-session-status-preview" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "Launch" ||
		payload["read_method"] != "GetKDEEntryPointsPreview" {
		t.Fatalf("unexpected KDE entrypoints schema: %#v", payload)
	}
	if payload["application_id"] != "org.example.ledger" ||
		payload["application_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" ||
		payload["file_count"] != float64(1) {
		t.Fatalf("unexpected KDE entrypoints identity: %#v", payload)
	}
	sessionStatus := payload["session_status"].(map[string]any)
	if sessionStatus["request_type"] != "execution-session-status-preview" ||
		sessionStatus["session_state"] != "planned-blocked" ||
		sessionStatus["gate_count"] != float64(5) ||
		sessionStatus["blocked_gate_count"] != float64(1) ||
		sessionStatus["desktop_surface_state"] != "planned" ||
		sessionStatus["user_visible_state"] != "Runtime gates required" ||
		sessionStatus["runtime_launch_approval"] != false ||
		sessionStatus["launch_allowed"] != false ||
		sessionStatus["execution_started"] != false {
		t.Fatalf("unexpected session status summary: %#v", sessionStatus)
	}
	if payload["entry_point_count"] != float64(7) ||
		payload["visible_entry_point_count"] != float64(7) ||
		payload["planned_entry_point_count"] != float64(7) ||
		payload["active_entry_point_count"] != float64(0) ||
		payload["portal_entry_point_count"] != float64(1) ||
		payload["runtime_gate_entry_point_count"] != float64(7) {
		t.Fatalf("unexpected entrypoint counts: %#v", payload)
	}
	ids := payload["entry_point_ids"].([]any)
	expectedIDs := []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "settings"}
	for index, expected := range expectedIDs {
		if ids[index] != expected {
			t.Fatalf("unexpected entrypoint ids: %#v", ids)
		}
	}
	entryPoints := payload["entry_points"].([]any)
	for _, rawEntryPoint := range entryPoints {
		entryPoint := rawEntryPoint.(map[string]any)
		if entryPoint["visible"] != true ||
			entryPoint["planned"] != true ||
			entryPoint["active"] != false ||
			entryPoint["requires_runtime_gate"] != true ||
			entryPoint["blocked_by_runtime_gate"] != true ||
			entryPoint["writes_host"] != false ||
			entryPoint["starts_backend"] != false ||
			entryPoint["backend_details_exposed"] != false {
			t.Fatalf("unexpected entrypoint safety flags: %#v", entryPoint)
		}
	}
	fileManager := entryPoints[2].(map[string]any)
	if fileManager["id"] != "file-manager" ||
		fileManager["kde_component"] != "Dolphin" ||
		fileManager["runtime_source"] != "file-open-preview" ||
		fileManager["requires_portal"] != true {
		t.Fatalf("unexpected file-manager entrypoint: %#v", fileManager)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["official_desktop_only"] != true ||
		payload["stable_desktop_contract"] != true ||
		payload["normal_application_surface"] != true ||
		payload["compatibility_center_card"] != true ||
		payload["safe_for_ai_diagnostics"] != true ||
		payload["desktop_entry_launch_visible"] != true ||
		payload["user_decision_captured"] != true ||
		payload["user_decision_allows_launch"] != true ||
		payload["entry_point_plan_created"] != true ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["settings_persisted"] != false ||
		payload["notifications_sent"] != false ||
		payload["task_manager_entry_active"] != false ||
		payload["kwin_rule_applied"] != false ||
		payload["live_tray_bridge_enabled"] != false ||
		payload["runtime_launch_approval"] != false ||
		payload["launch_allowed"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["request_objects_created"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE entrypoints safety flags: %#v", payload)
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

func TestModeSwitchPreviewCommandRendersUserFacingModes(t *testing.T) {
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
	err := run([]string{"mode-switch-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "prefer-performance"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.mode_switch.v1" ||
		payload["request_type"] != "mode-switch-preview" ||
		payload["plan_type"] != "compatibility-mode-switch-plan" ||
		payload["source"] != "unified-settings" ||
		payload["runtime_method"] != "GetCompatibilityModeSwitchPlan" {
		t.Fatalf("unexpected mode switch schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["display_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected mode switch identity: %#v", payload)
	}
	if payload["current_mode"] != "automatic" ||
		payload["requested_mode"] != "prefer-performance" ||
		payload["mode_state"] != "planned" ||
		payload["mode_count"] != float64(4) {
		t.Fatalf("unexpected mode switch request: %#v", payload)
	}
	modes := payload["modes"].([]any)
	if modes[0].(map[string]any)["id"] != "automatic" ||
		modes[1].(map[string]any)["id"] != "prefer-performance" ||
		modes[2].(map[string]any)["id"] != "prefer-compatibility" ||
		modes[3].(map[string]any)["id"] != "isolated-execution" {
		t.Fatalf("unexpected mode order: %#v", modes)
	}
	if modes[0].(map[string]any)["selected"] != true ||
		modes[1].(map[string]any)["requested"] != true {
		t.Fatalf("unexpected selected/requested mode markers: %#v", modes)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false || payload["user_visible"] != true ||
		payload["valid_mode"] != true || payload["requires_user_confirmation"] != true ||
		payload["portal_review_required"] != true || payload["snapshot_required"] != true ||
		payload["settings_persistence_enabled"] != false ||
		payload["backend_reconfiguration_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected mode switch safety flags: %#v", payload)
	}

	output.Reset()
	err = run([]string{"mode-switch-preview", "--registry", registryPath, "--app", "org.example.ledger", "--mode", "unsupported-mode"}, &output)
	if err == nil {
		t.Fatalf("mode-switch-preview accepted an unsupported mode")
	}
}

func TestSettingsChangePreviewCommandRendersReviewPlan(t *testing.T) {
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
	err := run([]string{"settings-change-preview", "--registry", registryPath, "--app", "org.example.ledger", "--section", "resource-access", "--field", "documents", "--value", "allow"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.settings_change.v1" ||
		payload["request_type"] != "settings-change-preview" ||
		payload["plan_type"] != "settings-change-plan" ||
		payload["source"] != "unified-settings" ||
		payload["runtime_method"] != "GetCompatibilitySettingsChangePlan" {
		t.Fatalf("unexpected settings change schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["display_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected settings change identity: %#v", payload)
	}
	if payload["section_id"] != "resource-access" ||
		payload["field_id"] != "documents" ||
		payload["requested_value"] != "allow" ||
		payload["change_state"] != "planned" {
		t.Fatalf("unexpected settings change request: %#v", payload)
	}
	if payload["user_confirmation_required"] != true ||
		payload["portal_policy_review_required"] != true ||
		payload["snapshot_recommended"] != false ||
		payload["runtime_restart_required"] != false {
		t.Fatalf("unexpected review requirements: %#v", payload)
	}
	affected := payload["affected_policy"].(map[string]any)
	if affected["section"] != "resource-access" ||
		affected["field"] != "documents" ||
		affected["value"] != "allow" {
		t.Fatalf("unexpected affected policy: %#v", affected)
	}
	steps := payload["steps"].([]any)
	if len(steps) != 5 ||
		steps[1].(map[string]any)["status"] != "required" ||
		steps[2].(map[string]any)["status"] != "required" ||
		steps[4].(map[string]any)["status"] != "pending" {
		t.Fatalf("unexpected settings change steps: %#v", steps)
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true || payload["apply_enabled"] != false ||
		payload["settings_persisted"] != false ||
		payload["settings_persistence_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected settings change safety flags: %#v", payload)
	}

	output.Reset()
	err = run([]string{"settings-change-preview", "--registry", registryPath, "--app", "org.example.ledger", "--section", "run-mode", "--field", "mode", "--value", "performance"}, &output)
	if err != nil {
		t.Fatalf("run-mode run returned error: %v", err)
	}
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal run-mode returned error: %v", err)
	}
	if payload["snapshot_recommended"] != true ||
		payload["portal_policy_review_required"] != false {
		t.Fatalf("unexpected run-mode requirements: %#v", payload)
	}
}

func TestReviewFlowPreviewCommandRendersConnectedReviewPlan(t *testing.T) {
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
	err := run([]string{"review-flow-preview", "--registry", registryPath, "--app", "org.example.ledger", "--section", "resource-access", "--field", "documents", "--value", "ask", "--operation", "file-open"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.review_flow.v1" ||
		payload["request_type"] != "review-flow-preview" ||
		payload["plan_type"] != "compatibility-review-flow-plan" ||
		payload["source"] != "compatibility-center-review" ||
		payload["runtime_method"] != "GetCompatibilityReviewFlowPlan" {
		t.Fatalf("unexpected review flow schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["display_name"] != "Example Ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected review flow identity: %#v", payload)
	}
	if payload["section_id"] != "resource-access" ||
		payload["field_id"] != "documents" ||
		payload["requested_value"] != "ask" ||
		payload["operation"] != "file-open" ||
		payload["review_state"] != "planned" {
		t.Fatalf("unexpected review flow request: %#v", payload)
	}
	if payload["step_count"] != float64(5) ||
		payload["required_review_count"] != float64(3) ||
		payload["blocked_step_count"] != float64(1) ||
		payload["pending_step_count"] != float64(1) {
		t.Fatalf("unexpected review flow counts: %#v", payload)
	}
	steps := payload["steps"].([]any)
	if steps[0].(map[string]any)["id"] != "settings-change-review" ||
		steps[1].(map[string]any)["id"] != "permission-review" ||
		steps[2].(map[string]any)["id"] != "portal-request-review" ||
		steps[3].(map[string]any)["status"] != "blocked" ||
		steps[4].(map[string]any)["status"] != "pending" {
		t.Fatalf("unexpected review flow steps: %#v", steps)
	}
	settingsChange := payload["settings_change_plan"].(map[string]any)
	if settingsChange["plan_type"] != "settings-change-plan" ||
		settingsChange["apply_enabled"] != false ||
		settingsChange["settings_persisted"] != false {
		t.Fatalf("unexpected settings change summary: %#v", settingsChange)
	}
	permissionReview := payload["permission_review_plan"].(map[string]any)
	if permissionReview["plan_type"] != "compatibility-permission-review-plan" ||
		permissionReview["permission_count"] != float64(7) ||
		permissionReview["permissions_granted"] != false {
		t.Fatalf("unexpected permission review summary: %#v", permissionReview)
	}
	portalRequest := payload["portal_request_plan"].(map[string]any)
	if portalRequest["request_type"] != "portal-request-preview" ||
		portalRequest["operation"] != "file-open" ||
		portalRequest["decision"] != "ask" ||
		portalRequest["request_object_created"] != false ||
		portalRequest["permission_granted"] != false {
		t.Fatalf("unexpected Portal request summary: %#v", portalRequest)
	}
	writeGate := payload["runtime_write_gate"].(map[string]any)
	if writeGate["gate_type"] != "runtime-write-gate" ||
		writeGate["method_name"] != "Launch" ||
		writeGate["write_method_enabled"] != false ||
		writeGate["dispatch_enabled"] != false {
		t.Fatalf("unexpected write gate summary: %#v", writeGate)
	}
	receipt := payload["review_receipt"].(map[string]any)
	if receipt["receipt_type"] != "compatibility-center-action-review-receipt" ||
		receipt["decision_recorded"] != false ||
		receipt["execution_enabled"] != false ||
		receipt["resource_grant_created"] != false {
		t.Fatalf("unexpected review receipt summary: %#v", receipt)
	}
	if payload["runtime_owned"] != true || payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false || payload["user_visible"] != true ||
		payload["user_confirmation_required"] != true ||
		payload["portal_policy_review_required"] != true ||
		payload["apply_enabled"] != false ||
		payload["request_object_created"] != false ||
		payload["permission_granted"] != false ||
		payload["settings_persisted"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected review flow safety flags: %#v", payload)
	}
}

func TestPermissionReviewPreviewCommandRendersKDEPermissionReview(t *testing.T) {
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
	err := run([]string{"permission-review-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.permission_review.v1" ||
		payload["request_type"] != "permission-review-preview" ||
		payload["plan_type"] != "compatibility-permission-review-plan" {
		t.Fatalf("unexpected permission review schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["source"] != "unified-settings" ||
		payload["runtime_method"] != "GetCompatibilityPermissionReviewPlan" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected permission review identity: %#v", payload)
	}
	if payload["permission_count"] != float64(7) ||
		payload["allow_count"] != float64(1) ||
		payload["ask_count"] != float64(5) ||
		payload["deny_count"] != float64(1) {
		t.Fatalf("unexpected permission counts: %#v", payload)
	}
	permissions := payload["permissions"].([]any)
	if permissions[0].(map[string]any)["id"] != "documents" ||
		permissions[2].(map[string]any)["id"] != "camera" ||
		permissions[2].(map[string]any)["decision"] != "deny" ||
		permissions[3].(map[string]any)["id"] != "network" ||
		permissions[3].(map[string]any)["decision"] != "allow" {
		t.Fatalf("unexpected permissions: %#v", permissions)
	}
	for _, value := range permissions {
		permission := value.(map[string]any)
		if permission["change_pending"] != false ||
			permission["request_object_created"] != false ||
			permission["permission_granted"] != false ||
			permission["direct_access_allowed"] != false ||
			permission["backend_details_exposed"] != false {
			t.Fatalf("permission gate unexpectedly open: %#v", permission)
		}
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["user_review_required"] != true ||
		payload["portal_review_required"] != true ||
		payload["permission_changes_applied"] != false ||
		payload["request_objects_created"] != false ||
		payload["permissions_granted"] != false ||
		payload["settings_persisted"] != false ||
		payload["settings_persistence_enabled"] != false ||
		payload["host_permission_changed"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected permission review safety flags: %#v", payload)
	}
}

func TestDesktopResourceBridgePreviewCommandRendersKDEBridgePlan(t *testing.T) {
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
	err := run([]string{"desktop-resource-bridge-preview", "--registry", registryPath, "--app", "org.example.ledger"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.desktop_resource_bridge.v1" ||
		payload["request_type"] != "desktop-resource-bridge-preview" ||
		payload["plan_type"] != "desktop-resource-bridge-plan" {
		t.Fatalf("unexpected desktop resource bridge schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["source"] != "runtime-resource-boundary" ||
		payload["runtime_method"] != "GetDesktopResourceBridgePlan" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop resource bridge identity: %#v", payload)
	}
	if payload["bridge_state"] != "planned" ||
		payload["resource_count"] != float64(5) ||
		payload["portal_mediated"] != true ||
		payload["file_bridge_planned"] != true ||
		payload["uri_bridge_planned"] != true ||
		payload["print_bridge_planned"] != true ||
		payload["clipboard_bridge_planned"] != true ||
		payload["screenshot_bridge_planned"] != true {
		t.Fatalf("unexpected bridge plan state: %#v", payload)
	}
	resources := payload["resources"].([]any)
	if resources[0].(map[string]any)["id"] != "file-open" ||
		resources[1].(map[string]any)["id"] != "uri-open" ||
		resources[2].(map[string]any)["id"] != "print" ||
		resources[3].(map[string]any)["id"] != "clipboard" ||
		resources[4].(map[string]any)["id"] != "screenshot" {
		t.Fatalf("unexpected bridge resources: %#v", resources)
	}
	for _, value := range resources {
		resource := value.(map[string]any)
		if resource["runtime_method"] != "GetPortalRequestPlan" ||
			resource["state"] != "planned" ||
			resource["portal_required"] != true ||
			resource["user_approval_required"] != true ||
			resource["bridge_enabled"] != false ||
			resource["request_created"] != false ||
			resource["direct_backend_access_allowed"] != false ||
			resource["backend_details_exposed"] != false {
			t.Fatalf("resource bridge gate unexpectedly open: %#v", resource)
		}
	}
	if payload["runtime_owned"] != true || payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["bridges_enabled"] != false ||
		payload["requests_created"] != false ||
		payload["backend_process_started"] != false ||
		payload["direct_host_file_access"] != false ||
		payload["direct_clipboard_access"] != false ||
		payload["direct_print_access"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected desktop resource bridge safety flags: %#v", payload)
	}
}

func TestPortalRequestPreviewCommandRendersPortalRequest(t *testing.T) {
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
	err := run([]string{"portal-request-preview", "--registry", registryPath, "--app", "org.example.ledger", "--operation", "file-open", "--reason", "Open a selected document."}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.portal_request.v1" ||
		payload["request_type"] != "portal-request-preview" ||
		payload["source"] != "runtime-portal-request-plan" {
		t.Fatalf("unexpected Portal request schema: %#v", payload)
	}
	if payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "GetPortalRequestPlan" ||
		payload["application_id"] != "org.example.ledger" ||
		payload["desktop_file"] != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected Portal request identity: %#v", payload)
	}
	if payload["operation"] != "file-open" ||
		payload["reason"] != "Open a selected document." ||
		payload["decision"] != "ask" ||
		payload["request_allowed"] != true {
		t.Fatalf("unexpected Portal request decision: %#v", payload)
	}
	portal := payload["portal"].(map[string]any)
	if portal["destination"] != "org.freedesktop.portal.Desktop" ||
		portal["interface"] != "org.freedesktop.portal.FileChooser" ||
		portal["method"] != "OpenFile" ||
		portal["object_path"] != "/org/freedesktop/portal/desktop" ||
		portal["dbus_api"] != "XDG Desktop Portal" {
		t.Fatalf("unexpected Portal endpoint: %#v", portal)
	}
	request := payload["request"].(map[string]any)
	if request["object_path_required"] != true ||
		request["request_object_created"] != false ||
		request["handle_token"] != "xnix_org_example_ledger_file_open" ||
		request["user_mediation_required"] != true ||
		request["runtime_policy_owner"] != true ||
		request["desktop_shell_policy_owner"] != false {
		t.Fatalf("unexpected Portal request object: %#v", request)
	}
	resources := request["resources"].([]any)
	if resources[0] != "documents" || resources[1] != "downloads" || resources[2] != "selected-files" {
		t.Fatalf("unexpected Portal resources: %#v", resources)
	}
	completion := payload["completion"].(map[string]any)
	if completion["signal"] != "Response" ||
		completion["response_field"] != "response" ||
		completion["success_code"] != float64(0) ||
		completion["cancelled_code"] != float64(1) ||
		completion["denied_code"] != float64(2) ||
		completion["result_owner"] != "Runtime" {
		t.Fatalf("unexpected Portal completion: %#v", completion)
	}
	if payload["denied"] != nil {
		t.Fatalf("allowed Portal request included denial guidance: %#v", payload["denied"])
	}
	safety := payload["safety"].(map[string]any)
	if safety["direct_access_allowed"] != false ||
		safety["portal_required"] != true ||
		safety["permission_granted"] != false ||
		safety["host_permission_changed"] != false ||
		safety["host_root_modified"] != false ||
		safety["backend_details_exposed"] != false {
		t.Fatalf("unexpected Portal safety flags: %#v", safety)
	}

	output.Reset()
	err = run([]string{"portal-request-preview", "--registry", registryPath, "--app", "org.example.ledger", "--operation", "camera"}, &output)
	if err != nil {
		t.Fatalf("camera run returned error: %v", err)
	}
	payload = map[string]any{}
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("camera Unmarshal returned error: %v", err)
	}
	if payload["decision"] != "deny" || payload["request_allowed"] != false {
		t.Fatalf("unexpected camera Portal decision: %#v", payload)
	}
	denied := payload["denied"].(map[string]any)
	if denied["next_action"] != "open-compatibility-settings" ||
		denied["notification_event"] != "approval-required" {
		t.Fatalf("unexpected denied guidance: %#v", denied)
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
