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
