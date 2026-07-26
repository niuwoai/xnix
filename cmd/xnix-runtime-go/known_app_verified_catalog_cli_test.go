package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestKnownAppVerifiedCatalogPreviewCommandConsumesMatrixEvidence(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_app_verified_catalog.v1" ||
		payload["request_type"] != "known-app-verified-catalog-preview" ||
		payload["source"] != "known-app-matrix-evidence+runtime-verified-catalog" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["runtime_method"] != "ListKnownVerifiedApplications" ||
		payload["read_method"] != "ListKnownVerifiedApplicationsPreview" ||
		payload["verification_source"] != "known-app-matrix-evidence-preview" ||
		payload["matrix_evidence_consumed"] != true ||
		payload["matrix_status"] != "passed" ||
		payload["application_count"] != float64(2) ||
		payload["verified_application_count"] != float64(2) ||
		payload["qemu_executed_count"] != float64(2) ||
		payload["wine_executed_count"] != float64(2) ||
		payload["checksum_verified_count"] != float64(2) ||
		payload["marker_observed_count"] != float64(2) ||
		payload["raw_output_redacted_count"] != float64(2) {
		t.Fatalf("unexpected verified catalog payload: %#v", payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["user_visible"] != true ||
		payload["standard_desktop_entries"] != true ||
		payload["review_only"] != true ||
		payload["operator_review_required"] != true ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["desktop_files_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["remote_path_exposed"] != false {
		t.Fatalf("unexpected verified catalog safety flags: %#v", payload)
	}
	if strings.Contains(output.String(), matrixEvidencePath) || strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("verified catalog exposed evidence paths: %s", output.String())
	}
	applications := payload["applications"].([]any)
	busybox := applications[1].(map[string]any)
	launchCommand := busybox["launch_request_command"].([]any)
	if busybox["app_id"] != "busybox-w32" ||
		busybox["verification_state"] != "verified-real-q4-matrix-run" ||
		busybox["desktop_catalog_state"] != "visible-review-only" ||
		busybox["launcher_surface"] != "xnix-compat-launch" ||
		launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "busybox-w32" ||
		busybox["direct_launch_enabled"] != false ||
		busybox["operator_review_required"] != true ||
		busybox["qemu_executed"] != true ||
		busybox["wine_executed"] != true ||
		busybox["checksum_verified"] != true ||
		busybox["marker_observed"] != true ||
		busybox["raw_output_redacted"] != true ||
		busybox["serial_log_evidence"] != true ||
		busybox["backend_launch_enabled"] != false ||
		busybox["backend_details_exposed"] != false ||
		busybox["raw_output_exposed"] != false ||
		busybox["remote_path_exposed"] != false ||
		busybox["host_root_modified"] != false {
		t.Fatalf("unexpected BusyBox verified catalog entry: %#v", busybox)
	}
}

func TestKnownAppVerifiedCatalogPreviewCommandConsumesGUIEvidencePacket(t *testing.T) {
	tempDir := t.TempDir()
	matrixEvidencePath := filepath.Join(tempDir, "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}
	guiPacketPath := filepath.Join(tempDir, "gui-evidence-packet.json")
	if err := os.WriteFile(guiPacketPath, []byte(knownAppVerifiedCatalogCLIGUIEvidencePacketFixture()), 0o600); err != nil {
		t.Fatalf("WriteFile GUI packet returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{
		"known-app-verified-catalog-preview",
		"--matrix-evidence", matrixEvidencePath,
		"--gui-evidence-packet", guiPacketPath,
	}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["source"] != "known-app-matrix-evidence+real-winapp-gui-evidence-packet+runtime-verified-catalog" ||
		payload["matrix_evidence_consumed"] != true ||
		payload["gui_evidence_packet_consumed"] != true ||
		payload["application_count"] != float64(3) ||
		payload["verified_application_count"] != float64(3) ||
		payload["gui_verified_application_count"] != float64(1) ||
		payload["gui_window_observed_count"] != float64(1) ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["remote_path_exposed"] != false {
		t.Fatalf("unexpected GUI catalog payload: %#v", payload)
	}
	applications := payload["applications"].([]any)
	messagebox := applications[2].(map[string]any)
	if messagebox["app_id"] != "org.xnix.apps.messagebox" ||
		messagebox["verification_state"] != "verified-real-gui-q4-run" ||
		messagebox["evidence_source"] != "wine-guest-gui-smoke" ||
		messagebox["gui_evidence"] != true ||
		messagebox["window_observed"] != true ||
		messagebox["file_open_verified"] != true ||
		messagebox["direct_launch_enabled"] != false ||
		messagebox["backend_launch_enabled"] != false ||
		messagebox["backend_details_exposed"] != false ||
		messagebox["raw_output_exposed"] != false ||
		messagebox["remote_path_exposed"] != false ||
		messagebox["host_root_modified"] != false {
		t.Fatalf("unexpected GUI application entry: %#v", messagebox)
	}
	if strings.Contains(output.String(), matrixEvidencePath) ||
		strings.Contains(output.String(), guiPacketPath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("GUI verified catalog exposed evidence paths: %s", output.String())
	}
}

func TestKnownAppVerifiedCatalogPreviewCommandRequiresMatrixEvidence(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --matrix-evidence") {
		t.Fatalf("expected missing matrix evidence error, got %v", err)
	}
}

func TestKnownAppVerifiedCatalogPreviewCommandRejectsUnsafeMatrixEvidence(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	var evidence appidentity.KnownAppMatrixEvidencePreview
	if err := json.Unmarshal(knownAppVerifiedCatalogCLIFixture(t), &evidence); err != nil {
		t.Fatalf("Unmarshal fixture returned error: %v", err)
	}
	evidence.RawOutputExposed = true
	content, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal fixture returned error: %v", err)
	}
	if err := os.WriteFile(matrixEvidencePath, content, 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &output)
	if err == nil || !strings.Contains(err.Error(), "redacted matrix evidence") {
		t.Fatalf("expected unsafe matrix evidence error, got %v", err)
	}
}

func TestCompatibilityCenterPreviewCommandConsumesKnownAppVerifiedCatalog(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-verified-catalog", catalogPath}, &output)
	if err != nil {
		t.Fatalf("compatibility-center-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_verified_catalog_consumed"] != true ||
		payload["known_app_verified_catalog_application_count"] != float64(2) ||
		payload["known_app_verified_catalog_review_only"] != true ||
		payload["action_execution_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected center verified catalog summary: %#v", payload)
	}
	appIDs := payload["known_app_verified_catalog_application_ids"].([]any)
	if len(appIDs) != 2 || appIDs[0] != "7zr" || appIDs[1] != "busybox-w32" {
		t.Fatalf("unexpected verified catalog application ids: %#v", appIDs)
	}
	applications := payload["known_app_verified_catalog_applications"].([]any)
	if len(applications) != 2 {
		t.Fatalf("expected two verified catalog applications, got %#v", applications)
	}
	busybox := applications[1].(map[string]any)
	launchCommand := busybox["launch_request_command"].([]any)
	if busybox["app_id"] != "busybox-w32" ||
		busybox["verification_state"] != "verified-real-q4-matrix-run" ||
		busybox["desktop_catalog_state"] != "visible-review-only" ||
		busybox["launcher_surface"] != "xnix-compat-launch" ||
		launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "busybox-w32" ||
		busybox["direct_launch_enabled"] != false ||
		busybox["operator_review_required"] != true ||
		busybox["qemu_executed"] != true ||
		busybox["wine_executed"] != true ||
		busybox["checksum_verified"] != true ||
		busybox["marker_observed"] != true ||
		busybox["raw_output_redacted"] != true ||
		busybox["serial_log_evidence"] != true ||
		busybox["backend_launch_enabled"] != false ||
		busybox["backend_details_exposed"] != false ||
		busybox["raw_output_exposed"] != false ||
		busybox["remote_path_exposed"] != false ||
		busybox["host_root_modified"] != false {
		t.Fatalf("unexpected BusyBox center catalog entry: %#v", busybox)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), matrixEvidencePath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("Compatibility Center exposed verified catalog paths: %s", output.String())
	}
}

func TestCompatibilityCenterPreviewCommandConsumesKnownAppVerifiedCatalogGUIEvidence(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	catalogPath := writeKnownAppVerifiedCatalogCLIFile(t, true)

	var output bytes.Buffer
	err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-verified-catalog", catalogPath}, &output)
	if err != nil {
		t.Fatalf("compatibility-center-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_verified_catalog_consumed"] != true ||
		payload["known_app_verified_catalog_application_count"] != float64(3) ||
		payload["known_app_verified_catalog_gui_application_count"] != float64(1) ||
		payload["known_app_verified_catalog_review_only"] != true ||
		payload["action_execution_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected center GUI catalog summary: %#v", payload)
	}
	guiIDs := payload["known_app_verified_catalog_gui_application_ids"].([]any)
	if len(guiIDs) != 1 || guiIDs[0] != "org.xnix.apps.messagebox" {
		t.Fatalf("unexpected GUI verified catalog application ids: %#v", guiIDs)
	}
	guiApps := payload["known_app_verified_catalog_gui_applications"].([]any)
	if len(guiApps) != 1 {
		t.Fatalf("expected one GUI verified catalog application, got %#v", guiApps)
	}
	messagebox := guiApps[0].(map[string]any)
	if messagebox["app_id"] != "org.xnix.apps.messagebox" ||
		messagebox["verification_state"] != "verified-real-gui-q4-run" ||
		messagebox["evidence_source"] != "wine-guest-gui-smoke" ||
		messagebox["gui_evidence"] != true ||
		messagebox["window_observed"] != true ||
		messagebox["file_open_verified"] != true ||
		messagebox["direct_launch_enabled"] != false ||
		messagebox["backend_launch_enabled"] != false ||
		messagebox["backend_details_exposed"] != false ||
		messagebox["raw_output_exposed"] != false ||
		messagebox["remote_path_exposed"] != false ||
		messagebox["host_root_modified"] != false {
		t.Fatalf("unexpected center GUI catalog application: %#v", messagebox)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") ||
		strings.Contains(strings.ToLower(output.String()), "wine/") ||
		strings.Contains(strings.ToLower(output.String()), ".wine") {
		t.Fatalf("center GUI catalog output exposed unsafe details: %s", output.String())
	}
}

func TestCompatibilityCenterPreviewCommandRejectsUnsafeKnownAppVerifiedCatalog(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	var catalog appidentity.KnownAppVerifiedCatalogPreview
	if err := json.Unmarshal(catalogOutput.Bytes(), &catalog); err != nil {
		t.Fatalf("Unmarshal verified catalog returned error: %v", err)
	}
	catalog.LaunchEnabled = true
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal verified catalog returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, content, 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-verified-catalog", catalogPath}, &output)
	if err == nil || !strings.Contains(err.Error(), "review-only") {
		t.Fatalf("expected unsafe verified catalog error, got %v", err)
	}
}

func TestKDECenterPagePreviewCommandConsumesKnownAppVerifiedCatalog(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-verified-catalog", catalogPath}, &output)
	if err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !strings.Contains(payload["source"].(string), "known-app-verified-catalog") ||
		payload["known_app_verified_catalog_consumed"] != true ||
		payload["known_app_verified_catalog_application_count"] != float64(2) ||
		payload["known_app_verified_catalog_review_only"] != true ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE verified catalog summary: %#v", payload)
	}
	cards := payload["known_app_verified_catalog_cards"].([]any)
	if len(cards) != 2 {
		t.Fatalf("expected two KDE verified catalog cards, got %#v", cards)
	}
	busybox := cards[1].(map[string]any)
	launchCommand := busybox["launch_request_command"].([]any)
	if busybox["app_id"] != "busybox-w32" ||
		busybox["verification_state"] != "verified-real-q4-matrix-run" ||
		busybox["desktop_catalog_state"] != "visible-review-only" ||
		busybox["launcher_surface"] != "xnix-compat-launch" ||
		launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "busybox-w32" ||
		busybox["operator_review_required"] != true ||
		busybox["qemu_executed"] != true ||
		busybox["wine_executed"] != true ||
		busybox["direct_launch_enabled"] != false ||
		busybox["backend_launch_enabled"] != false ||
		busybox["backend_details_exposed"] != false ||
		busybox["remote_path_exposed"] != false ||
		busybox["host_root_modified"] != false {
		t.Fatalf("unexpected BusyBox KDE catalog card: %#v", busybox)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), matrixEvidencePath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("KDE Center page exposed verified catalog paths: %s", output.String())
	}
}

func TestKDECenterPagePreviewCommandConsumesKnownAppVerifiedCatalogGUIEvidence(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	catalogPath := writeKnownAppVerifiedCatalogCLIFile(t, true)

	var output bytes.Buffer
	err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-verified-catalog", catalogPath}, &output)
	if err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !strings.Contains(payload["source"].(string), "known-app-verified-catalog-gui") ||
		payload["known_app_verified_catalog_consumed"] != true ||
		payload["known_app_verified_catalog_application_count"] != float64(3) ||
		payload["known_app_verified_catalog_gui_application_count"] != float64(1) ||
		payload["known_app_verified_catalog_review_only"] != true ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE GUI catalog summary: %#v", payload)
	}
	guiCards := payload["known_app_verified_catalog_gui_cards"].([]any)
	if len(guiCards) != 1 {
		t.Fatalf("expected one KDE GUI verified catalog card, got %#v", guiCards)
	}
	messagebox := guiCards[0].(map[string]any)
	launchCommand := messagebox["launch_request_command"].([]any)
	if messagebox["app_id"] != "org.xnix.apps.messagebox" ||
		messagebox["verification_state"] != "verified-real-gui-q4-run" ||
		messagebox["evidence_source"] != "wine-guest-gui-smoke" ||
		messagebox["gui_evidence"] != true ||
		messagebox["window_observed"] != true ||
		messagebox["file_open_verified"] != true ||
		launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "org.xnix.apps.messagebox" ||
		messagebox["operator_review_required"] != true ||
		messagebox["direct_launch_enabled"] != false ||
		messagebox["backend_launch_enabled"] != false ||
		messagebox["backend_details_exposed"] != false ||
		messagebox["remote_path_exposed"] != false ||
		messagebox["host_root_modified"] != false {
		t.Fatalf("unexpected KDE GUI catalog card: %#v", messagebox)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") ||
		strings.Contains(strings.ToLower(output.String()), "wine/") ||
		strings.Contains(strings.ToLower(output.String()), ".wine") {
		t.Fatalf("KDE GUI catalog output exposed unsafe details: %s", output.String())
	}
}

func TestKDECenterPagePreviewCommandRejectsUnsafeKnownAppVerifiedCatalog(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	var catalog appidentity.KnownAppVerifiedCatalogPreview
	if err := json.Unmarshal(catalogOutput.Bytes(), &catalog); err != nil {
		t.Fatalf("Unmarshal verified catalog returned error: %v", err)
	}
	catalog.Applications[0].DirectLaunchEnabled = true
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal verified catalog returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, content, 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-verified-catalog", catalogPath}, &output)
	if err == nil || !strings.Contains(err.Error(), "safe for review") {
		t.Fatalf("expected unsafe KDE verified catalog error, got %v", err)
	}
}

func TestKnownAppVerifiedCatalogRunPlanPreviewCommandSelectsQ4RunnableApp(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", "busybox-w32"}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-run-plan-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_run_plan.v1" ||
		payload["request_type"] != "known-app-verified-catalog-run-plan-preview" ||
		payload["source"] != "known-app-verified-catalog+q4-run-plan" ||
		payload["runtime_method"] != "PlanKnownVerifiedApplicationRun" ||
		payload["read_method"] != "GetKnownVerifiedApplicationRunPlan" ||
		payload["verified_catalog_consumed"] != true ||
		payload["requested_app_id"] != "busybox-w32" ||
		payload["app_id"] != "busybox-w32" ||
		payload["verification_state"] != "verified-real-q4-matrix-run" ||
		payload["desktop_catalog_state"] != "visible-review-only" {
		t.Fatalf("unexpected run plan payload: %#v", payload)
	}
	launchCommand := payload["launch_request_command"].([]any)
	remoteSmokeCommand := payload["remote_smoke_command"].([]any)
	if launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "busybox-w32" ||
		remoteSmokeCommand[0] != "ruby" ||
		remoteSmokeCommand[1] != "scripts/remote_known_winapp_guest_wine_smoke.rb" ||
		remoteSmokeCommand[2] != "--execute" ||
		remoteSmokeCommand[3] != "--app" ||
		remoteSmokeCommand[4] != "busybox-w32" ||
		payload["remote_smoke_request_type"] != "remote-known-winapp-guest-wine-smoke" {
		t.Fatalf("unexpected run plan commands: %#v", payload)
	}
	if payload["q4_execution_required"] != true ||
		payload["q4_execution_planned"] != true ||
		payload["q4_execution_started"] != false ||
		payload["review_only"] != true ||
		payload["operator_review_required"] != true ||
		payload["direct_launch_enabled"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["desktop_files_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["remote_path_exposed"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["host_compilation_required"] != false ||
		payload["host_compilation_avoided"] != true ||
		payload["targeted_remote_verification_ready"] != true {
		t.Fatalf("unexpected run plan safety flags: %#v", payload)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), matrixEvidencePath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("run plan exposed local or q4 evidence paths: %s", output.String())
	}
}

func TestKnownAppVerifiedCatalogRunPlanPreviewCommandSelectsGUIRunnableApp(t *testing.T) {
	catalogPath := writeKnownAppVerifiedCatalogCLIFile(t, true)

	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", "org.xnix.apps.messagebox"}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-run-plan-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_run_plan.v1" ||
		payload["request_type"] != "known-app-verified-catalog-run-plan-preview" ||
		payload["source"] != "known-app-verified-catalog+q4-run-plan" ||
		payload["runtime_method"] != "PlanKnownVerifiedApplicationRun" ||
		payload["read_method"] != "GetKnownVerifiedApplicationRunPlan" ||
		payload["verified_catalog_consumed"] != true ||
		payload["requested_app_id"] != "org.xnix.apps.messagebox" ||
		payload["app_id"] != "org.xnix.apps.messagebox" ||
		payload["verification_state"] != "verified-real-gui-q4-run" ||
		payload["compatibility_state"] != "owner-controlled-gui-qemu-wine-verified" ||
		payload["desktop_catalog_state"] != "visible-review-only" {
		t.Fatalf("unexpected GUI run plan payload: %#v", payload)
	}
	launchCommand := payload["launch_request_command"].([]any)
	remoteSmokeCommand := payload["remote_smoke_command"].([]any)
	forwardedArguments := payload["desktop_forwarded_arguments"].([]any)
	if launchCommand[0] != "xnix-compat-launch" ||
		launchCommand[1] != "--app" ||
		launchCommand[2] != "org.xnix.apps.messagebox" ||
		remoteSmokeCommand[0] != "ruby" ||
		remoteSmokeCommand[1] != "scripts/q4_messagebox_smoke.rb" ||
		remoteSmokeCommand[2] != "--execute" ||
		remoteSmokeCommand[3] != "--owner-file-open" ||
		payload["remote_smoke_request_type"] != "q4-messagebox-smoke" ||
		forwardedArguments[0] != "org.xnix.apps.messagebox" {
		t.Fatalf("unexpected GUI run plan commands: %#v", payload)
	}
	if payload["runtime_owned_action_ready"] != true ||
		payload["desktop_callable_action_id"] != "review-known-app-gui-evidence" ||
		payload["desktop_callable_route"] != "runtime-owner://known-app-verified-catalog/run-plan" ||
		payload["desktop_callable_runtime_method"] != "PlanKnownVerifiedApplicationRun" ||
		payload["desktop_callable_execution_type"] != "review-only-q4-gui-smoke" ||
		payload["desktop_forwards_only_app_id"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false ||
		payload["desktop_owner_inputs_exposed"] != false ||
		payload["gui_evidence_required"] != true ||
		payload["gui_evidence_consumed"] != true ||
		payload["window_observation_required"] != true ||
		payload["owner_file_open_required"] != true ||
		payload["owner_file_open_verified"] != true {
		t.Fatalf("unexpected GUI run plan action handoff fields: %#v", payload)
	}
	if payload["q4_execution_required"] != true ||
		payload["q4_execution_planned"] != true ||
		payload["q4_execution_started"] != false ||
		payload["review_only"] != true ||
		payload["operator_review_required"] != true ||
		payload["direct_launch_enabled"] != false ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["desktop_files_written"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["remote_path_exposed"] != false ||
		payload["host_compilation_required"] != false ||
		payload["host_compilation_avoided"] != true ||
		payload["targeted_remote_verification_ready"] != true {
		t.Fatalf("unexpected GUI run plan safety flags: %#v", payload)
	}
	if strings.Contains(output.String(), catalogPath) ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("GUI run plan exposed local or q4 evidence paths: %s", output.String())
	}
}

func TestKnownAppVerifiedCatalogRunPlanPreviewCommandRejectsUnknownApp(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "missing-app") {
		t.Fatalf("expected unknown app error, got %v", err)
	}
}

func TestKnownAppVerifiedCatalogRunAcceptancePreviewCommandConsumesMatchedRun(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}

	var runPlanOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", "7zr"}, &runPlanOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-run-plan-preview returned error: %v", err)
	}
	runPlanPath := filepath.Join(t.TempDir(), "known-app-verified-catalog-run-plan.json")
	if err := os.WriteFile(runPlanPath, runPlanOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile run plan returned error: %v", err)
	}
	runReportPath := filepath.Join(t.TempDir(), "known-winapp-run.json")
	if err := os.WriteFile(runReportPath, []byte(knownExistingWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile run report returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-run-acceptance-preview", "--run-plan", runPlanPath, "--known-winapp-run", runReportPath}, &output)
	if err != nil {
		t.Fatalf("known-app-verified-catalog-run-acceptance-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_app_verified_catalog_run_acceptance.v1" ||
		payload["request_type"] != "known-app-verified-catalog-run-acceptance-preview" ||
		payload["source"] != "known-app-verified-catalog-run-plan+known-existing-windows-app-acceptance" ||
		payload["runtime_method"] != "PreviewKnownVerifiedApplicationRunAcceptance" ||
		payload["read_method"] != "GetKnownVerifiedApplicationRunAcceptance" ||
		payload["acceptance_type"] != "verified-catalog-app-q4-real-run-acceptance" ||
		payload["run_plan_consumed"] != true ||
		payload["run_report_consumed"] != true ||
		payload["requested_app_id"] != "7zr" ||
		payload["app_id"] != "7zr" ||
		payload["run_plan_matched"] != true ||
		payload["acceptance_ready"] != true {
		t.Fatalf("unexpected run acceptance payload: %#v", payload)
	}
	if payload["q4_execution_observed"] != true ||
		payload["host_compilation_avoided"] != true ||
		payload["checksum_verified"] != true ||
		payload["marker_observed"] != true ||
		payload["runtime_started_isolated_guest"] != true ||
		payload["compatibility_engine_execution_observed"] != true ||
		payload["output_redacted"] != true ||
		payload["run_plan_path_exposed"] != false ||
		payload["run_report_path_exposed"] != false ||
		payload["remote_host_exposed"] != false ||
		payload["raw_path_exposed"] != false ||
		payload["raw_output_exposed"] != false ||
		payload["runtime_argv_exposed"] != false ||
		payload["runner_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected run acceptance safety flags: %#v", payload)
	}
	if strings.Contains(output.String(), runPlanPath) ||
		strings.Contains(output.String(), runReportPath) ||
		strings.Contains(output.String(), "root@q4") ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("run acceptance exposed paths or remote host: %s", output.String())
	}
}

func TestKnownAppVerifiedCatalogRunAcceptancePreviewCommandRejectsMismatchedRun(t *testing.T) {
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}

	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}
	var runPlanOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", "busybox-w32"}, &runPlanOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-run-plan-preview returned error: %v", err)
	}
	runPlanPath := filepath.Join(t.TempDir(), "known-app-verified-catalog-run-plan.json")
	if err := os.WriteFile(runPlanPath, runPlanOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile run plan returned error: %v", err)
	}
	runReportPath := filepath.Join(t.TempDir(), "known-winapp-run.json")
	if err := os.WriteFile(runReportPath, []byte(knownExistingWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile run report returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{"known-app-verified-catalog-run-acceptance-preview", "--run-plan", runPlanPath, "--known-winapp-run", runReportPath}, &output)
	if err == nil || !strings.Contains(err.Error(), "app ids to match") {
		t.Fatalf("expected mismatched app error, got %v", err)
	}
}

func TestCompatibilityCenterPreviewCommandConsumesKnownAppVerifiedCatalogRunAcceptance(t *testing.T) {
	registryPath, _ := writeTestRepairGroupRegistry(t)
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")

	var output bytes.Buffer
	err := run([]string{"compatibility-center-preview", "--registry", registryPath, "--known-app-verified-catalog-run-acceptance", acceptancePath}, &output)
	if err != nil {
		t.Fatalf("compatibility-center-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["known_app_smoke_evidence_count"] != float64(1) ||
		payload["known_app_smoke_passed_count"] != float64(1) ||
		payload["known_app_launch_authorization_required_count"] != float64(1) ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected compatibility center accepted run payload: %#v", payload)
	}
	evidenceItems := payload["known_app_smoke_evidence"].([]any)
	first := evidenceItems[0].(map[string]any)
	if first["app_id"] != "7zr" ||
		first["evidence_kind"] != "known-application-verified-catalog-run-acceptance" ||
		first["evidence_source"] != "verified-catalog-app-q4-real-run-acceptance" ||
		first["compatibility_state"] != "verified-catalog-real-q4-run-accepted" ||
		first["center_card_state"] != "validated-verified-catalog-real-runtime-run" ||
		first["primary_action_id"] != "review-known-app-verified-catalog-run-acceptance" ||
		first["primary_action_kind"] != "review" ||
		first["marker_observed"] != true ||
		first["checksum_verified"] != true ||
		first["execution_evidence_recorded"] != true ||
		first["runtime_dispatch_verified"] != true ||
		first["desktop_launch_enabled"] != false ||
		first["backend_launch_enabled"] != false ||
		first["backend_details_exposed"] != false ||
		first["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected accepted run evidence item: %#v", first)
	}
	if strings.Contains(output.String(), acceptancePath) ||
		strings.Contains(output.String(), "root@q4") ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("compatibility center accepted run exposed unsafe details: %s", output.String())
	}
}

func TestKDECenterPagePreviewCommandConsumesKnownAppVerifiedCatalogRunAcceptance(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	acceptancePath := writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t, "7zr")

	var output bytes.Buffer
	err := run([]string{"kde-center-page-preview", "--registry", registryPath, "--app", app, "--decision", "approved", "--known-app-verified-catalog-run-acceptance", acceptancePath}, &output)
	if err != nil {
		t.Fatalf("kde-center-page-preview returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if !strings.Contains(payload["source"].(string), "known-app-verified-catalog-run-acceptance") ||
		payload["known_app_matrix_evidence_count"] != float64(1) ||
		payload["launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE center accepted run payload: %#v", payload)
	}
	cards := payload["known_app_matrix_evidence_cards"].([]any)
	card := cards[0].(map[string]any)
	if card["app_id"] != "7zr" ||
		card["evidence_kind"] != "known-application-verified-catalog-run-acceptance" ||
		card["evidence_source"] != "verified-catalog-app-q4-real-run-acceptance" ||
		card["compatibility_state"] != "verified-catalog-real-q4-run-accepted" ||
		card["center_card_state"] != "validated-verified-catalog-real-runtime-run" ||
		card["primary_action_id"] != "review-known-app-verified-catalog-run-acceptance" ||
		card["primary_action_kind"] != "review" ||
		card["marker_observed"] != true ||
		card["checksum_verified"] != true ||
		card["execution_evidence_recorded"] != true ||
		card["runtime_dispatch_verified"] != true ||
		card["desktop_launch_enabled"] != false ||
		card["backend_launch_enabled"] != false ||
		card["host_root_modified"] != false ||
		card["backend_details_exposed"] != false ||
		card["raw_artifact_path_exposed"] != false {
		t.Fatalf("unexpected KDE accepted run card: %#v", card)
	}
	if strings.Contains(output.String(), acceptancePath) ||
		strings.Contains(output.String(), "root@q4") ||
		strings.Contains(output.String(), "/home/xnix-run-materials") {
		t.Fatalf("KDE center accepted run exposed unsafe details: %s", output.String())
	}
}

func writeKnownAppVerifiedCatalogRunAcceptanceCLIFile(t *testing.T, appID string) string {
	t.Helper()
	matrixEvidencePath := filepath.Join(t.TempDir(), "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}
	var catalogOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(t.TempDir(), "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}
	var runPlanOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-run-plan-preview", "--verified-catalog", catalogPath, "--app", appID}, &runPlanOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-run-plan-preview returned error: %v", err)
	}
	runPlanPath := filepath.Join(t.TempDir(), "known-app-verified-catalog-run-plan.json")
	if err := os.WriteFile(runPlanPath, runPlanOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile run plan returned error: %v", err)
	}
	runReportPath := filepath.Join(t.TempDir(), "known-winapp-run.json")
	if err := os.WriteFile(runReportPath, []byte(knownExistingWinAppAcceptanceCLIFixture(currentProjectVersion(t))), 0o600); err != nil {
		t.Fatalf("WriteFile run report returned error: %v", err)
	}
	var acceptanceOutput bytes.Buffer
	if err := run([]string{"known-app-verified-catalog-run-acceptance-preview", "--run-plan", runPlanPath, "--known-winapp-run", runReportPath}, &acceptanceOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-run-acceptance-preview returned error: %v", err)
	}
	acceptancePath := filepath.Join(t.TempDir(), "known-app-verified-catalog-run-acceptance.json")
	if err := os.WriteFile(acceptancePath, acceptanceOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile run acceptance returned error: %v", err)
	}
	return acceptancePath
}

func writeKnownAppVerifiedCatalogCLIFile(t *testing.T, includeGUIEvidence bool) string {
	t.Helper()
	tempDir := t.TempDir()
	matrixEvidencePath := filepath.Join(tempDir, "known-app-matrix-evidence.json")
	if err := os.WriteFile(matrixEvidencePath, knownAppVerifiedCatalogCLIFixture(t), 0o600); err != nil {
		t.Fatalf("WriteFile matrix evidence returned error: %v", err)
	}
	args := []string{"known-app-verified-catalog-preview", "--matrix-evidence", matrixEvidencePath}
	if includeGUIEvidence {
		guiPacketPath := filepath.Join(tempDir, "gui-evidence-packet.json")
		if err := os.WriteFile(guiPacketPath, []byte(knownAppVerifiedCatalogCLIGUIEvidencePacketFixture()), 0o600); err != nil {
			t.Fatalf("WriteFile GUI packet returned error: %v", err)
		}
		args = append(args, "--gui-evidence-packet", guiPacketPath)
	}
	var catalogOutput bytes.Buffer
	if err := run(args, &catalogOutput); err != nil {
		t.Fatalf("known-app-verified-catalog-preview returned error: %v", err)
	}
	catalogPath := filepath.Join(tempDir, "known-app-verified-catalog.json")
	if err := os.WriteFile(catalogPath, catalogOutput.Bytes(), 0o600); err != nil {
		t.Fatalf("WriteFile verified catalog returned error: %v", err)
	}
	return catalogPath
}

func knownAppVerifiedCatalogCLIFixture(t *testing.T) []byte {
	t.Helper()
	apps := []appidentity.KnownAppMatrixEvidenceApp{
		knownAppVerifiedCatalogCLIApp("7zr", "7-Zip standalone console executable", "26.02"),
		knownAppVerifiedCatalogCLIApp("busybox-w32", "BusyBox-w32 standalone console executable", "current-2026-07-24"),
	}
	evidence := appidentity.KnownAppMatrixEvidencePreview{
		SchemaVersion:                      appidentity.KnownAppMatrixEvidencePreviewSchemaVersion,
		RequestType:                        appidentity.KnownAppMatrixEvidencePreviewRequestType,
		Source:                             "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
		RuntimeMethod:                      "PreviewKnownAppMatrixEvidence",
		ReadMethod:                         "GetKnownAppMatrixEvidence",
		MatrixStatus:                       "passed",
		MatrixReportConsumed:               true,
		MatrixReportPathExposed:            false,
		MatrixReportOutputWritten:          true,
		AppCount:                           2,
		PassedCount:                        2,
		FailedCount:                        0,
		EvidenceCount:                      2,
		PassedEvidenceCount:                2,
		FailedEvidenceCount:                0,
		QEMUExecutedCount:                  2,
		WineExecutedCount:                  2,
		MarkerObservedCount:                2,
		ChecksumVerifiedCount:              2,
		RawOutputRedactedCount:             2,
		SerialLogEvidenceCount:             2,
		GuestStartedCount:                  2,
		GuestPortAutoCount:                 2,
		CompatibilityCenterProjectionReady: true,
		KDECenterProjectionReady:           true,
		Apps:                               apps,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		DesktopLaunchEnabled:               false,
		BackendLaunchEnabled:               false,
		ActionExecutionEnabled:             false,
		BackendDetailsExposed:              false,
		RawOutputExposed:                   false,
		RemotePathExposed:                  false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		HostNetworkingRequired:             false,
		DockerSocketMounted:                false,
		BroadHostMountRequired:             false,
	}
	content, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal fixture returned error: %v", err)
	}
	return content
}

func knownAppVerifiedCatalogCLIApp(appID string, displayName string, appVersion string) appidentity.KnownAppMatrixEvidenceApp {
	return appidentity.KnownAppMatrixEvidenceApp{
		AppID:                 appID,
		DisplayName:           displayName,
		AppVersion:            appVersion,
		SmokeStatus:           "passed",
		CompatibilityState:    "real-qemu-wine-verified",
		MarkerObserved:        true,
		ChecksumVerified:      true,
		QEMUExecuted:          true,
		WineExecuted:          true,
		GuestStarted:          true,
		GuestPortAuto:         true,
		RawOutputRedacted:     true,
		SerialLogEvidence:     true,
		ReportEvidence:        true,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		DesktopLaunchEnabled:  false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
		RawOutputExposed:      false,
		RemotePathExposed:     false,
		HostRootModified:      false,
	}
}

func knownAppVerifiedCatalogCLIGUIEvidencePacketFixture() string {
	return `{
  "version": "0.2.640-test",
  "schema_version": "xnix.runtime.real_winapp_gui_evidence_packet.v1",
  "request_type": "real-winapp-gui-evidence-packet-preview",
  "packet_type": "real-windows-app-gui-evidence",
  "source": "wine-guest-gui-smoke+runtime-evidence-consumer+real-winapp-desktop-packet",
  "runtime_method": "PreviewRealWinAppGUIEvidencePacket",
  "read_method": "GetRealWinAppGUIEvidencePacket",
  "report_status": "passed",
  "report_consumed": true,
  "report_path_exposed": false,
  "app_id": "org.xnix.apps.messagebox",
  "display_name": "Xnix MessageBox",
  "app_version": "0.2.640-test",
  "gui_app_name": "xnix-messagebox-smoke.exe",
  "evidence_source": "wine-guest-gui-smoke",
  "compatibility_state": "owner-controlled-gui-qemu-wine-verified",
  "center_card_state": "validated-owner-controlled-gui-runtime-run",
  "known_app_gui_evidence_count": 1,
  "known_app_gui_evidence_verified_count": 1,
  "known_app_smoke_evidence": {
    "app_id": "org.xnix.apps.messagebox",
    "display_name": "Xnix MessageBox",
    "app_version": "0.2.640-test",
    "evidence_kind": "known-application-gui-smoke",
    "evidence_source": "wine-guest-gui-smoke",
    "smoke_status": "passed",
    "x_window_observed": true,
    "window_observed": true,
    "compatibility_state": "owner-controlled-gui-qemu-wine-verified",
    "center_card_state": "validated-owner-controlled-gui-runtime-run",
    "owner_file_open_verified": true,
    "owner_file_open_entrypoint_invoked": true,
    "owner_delegated_file_argument_count": 1,
    "owner_delegated_file_argument_copied_count": 1,
    "owner_delegated_file_arguments_passed": true,
    "owner_delegated_file_argument_winepath_translated": true,
    "owner_delegated_file_argument_winepath_translated_count": 1,
    "owner_delegated_raw_file_argument_path_exposed": false,
    "owner_delegated_window_match_observed": true,
    "marker_observed": true,
    "execution_evidence_recorded": true,
    "runtime_dispatch_verified": true,
    "launch_authorization_required": true,
    "desktop_launch_enabled": false,
    "runtime_owned": true,
    "kde_policy_owner": false,
    "action_execution_enabled": false,
    "backend_launch_enabled": false,
    "host_root_modified": false,
    "backend_details_exposed": false,
    "raw_artifact_path_exposed": false
  },
  "wineboot_invoked": true,
  "x_window_observed": true,
  "window_observed": true,
  "x_window_child_count": 1,
  "compatibility_center_projection_ready": true,
  "kde_center_projection_ready": true,
  "container_runtime_used": false,
  "container_host_mount_count": 0,
  "runtime_owned": true,
  "go_runtime_backed": true,
  "kde_policy_owner": false,
  "desktop_launch_enabled": false,
  "backend_launch_enabled": false,
  "action_execution_enabled": false,
  "backend_details_exposed": false,
  "raw_output_exposed": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "desktop_safe_summary": "Xnix MessageBox produced owner-controlled GUI evidence."
}`
}
