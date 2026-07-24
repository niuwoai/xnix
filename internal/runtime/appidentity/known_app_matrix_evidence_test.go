package appidentity

import (
	"strings"
	"testing"
)

func TestPreviewKnownAppMatrixEvidenceConsumesPassedRemoteMatrix(t *testing.T) {
	preview, err := PreviewKnownAppMatrixEvidenceJSON([]byte(knownAppMatrixEvidenceFixture(true)))
	if err != nil {
		t.Fatalf("PreviewKnownAppMatrixEvidenceJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppMatrixEvidencePreviewSchemaVersion ||
		preview.RequestType != KnownAppMatrixEvidencePreviewRequestType ||
		preview.Source != "remote-known-winapp-matrix-smoke+runtime-evidence-consumer" ||
		preview.RuntimeMethod != "PreviewKnownAppMatrixEvidence" ||
		preview.ReadMethod != "GetKnownAppMatrixEvidence" ||
		preview.MatrixStatus != "passed" ||
		!preview.MatrixReportConsumed ||
		preview.MatrixReportPathExposed ||
		!preview.MatrixReportOutputWritten ||
		preview.AppCount != 2 ||
		preview.PassedCount != 2 ||
		preview.FailedCount != 0 ||
		preview.EvidenceCount != 2 ||
		preview.PassedEvidenceCount != 2 ||
		preview.FailedEvidenceCount != 0 ||
		preview.QEMUExecutedCount != 2 ||
		preview.WineExecutedCount != 2 ||
		preview.MarkerObservedCount != 2 ||
		preview.ChecksumVerifiedCount != 2 ||
		preview.RawOutputRedactedCount != 2 ||
		preview.SerialLogEvidenceCount != 2 ||
		preview.GuestStartedCount != 2 ||
		preview.GuestPortAutoCount != 2 ||
		!preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ActionExecutionEnabled ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected known app matrix preview: %#v", preview)
	}
	if len(preview.Apps) != 2 || len(preview.KnownAppSmokeEvidence) != 2 {
		t.Fatalf("unexpected app evidence count: %#v", preview)
	}
	first := preview.Apps[0]
	if first.AppID != "7zr" ||
		first.DisplayName != "7-Zip standalone console executable" ||
		first.AppVersion != "26.02" ||
		first.ExecutableName != "7zr.exe" ||
		first.CompatibilityState != "real-qemu-wine-verified" ||
		!first.MarkerObserved ||
		!first.ChecksumVerified ||
		!first.QEMUExecuted ||
		!first.WineExecuted ||
		first.BackendDetailsExposed ||
		first.RemotePathExposed ||
		first.RawOutputExposed ||
		first.HostRootModified {
		t.Fatalf("unexpected first app evidence: %#v", first)
	}
	evidence := preview.KnownAppSmokeEvidence[1]
	if evidence.AppID != "busybox-w32" ||
		evidence.DisplayName != "BusyBox-w32 standalone console executable" ||
		evidence.EvidenceKind != "known-application-matrix-smoke" ||
		evidence.EvidenceSource != "remote-known-winapp-matrix-smoke" ||
		evidence.SmokeStatus != "passed" ||
		evidence.CompatibilityState != "real-qemu-wine-verified" ||
		evidence.CenterCardState != "validated-real-runtime-run" ||
		evidence.PrimaryActionID != "review-known-app-matrix-evidence" ||
		evidence.PrimaryActionKind != "review" ||
		!evidence.PrimaryActionEnabled ||
		!evidence.MarkerObserved ||
		!evidence.ChecksumVerified ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.RuntimeDispatchVerified ||
		!evidence.LaunchAuthorizationRequired ||
		evidence.DirectLaunchEnabled ||
		evidence.DesktopLaunchEnabled ||
		evidence.ActionExecutionEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		evidence.RawArtifactPathExposed {
		t.Fatalf("unexpected known app smoke evidence: %#v", evidence)
	}
}

func TestPreviewKnownAppMatrixEvidenceRejectsPassedAppWithoutChecksum(t *testing.T) {
	_, err := PreviewKnownAppMatrixEvidenceJSON([]byte(knownAppMatrixEvidenceFixture(false)))
	if err == nil || !strings.Contains(err.Error(), "checksum verification") {
		t.Fatalf("expected checksum verification error, got %v", err)
	}
}

func knownAppMatrixEvidenceFixture(checksumVerified bool) string {
	checksumValue := "true"
	if !checksumVerified {
		checksumValue = "false"
	}
	return `{
  "schema_version": "xnix.scripts.remote_known_windows_app_matrix_smoke.v1",
  "request_type": "remote-known-winapp-matrix-smoke",
  "status": "passed",
  "execute": true,
  "matrix_report_output": "/home/xnix-run-materials/state/known-run-matrix-version-under-test.json",
  "matrix_report_output_written": true,
  "app_count": 2,
  "app_ids": ["7zr", "busybox-w32"],
  "backend": "guest-wine",
  "start_qemu": true,
  "guest_port": "auto",
  "redact_output": true,
  "apps": [
    {
      "app_id": "7zr",
      "report_output": "/home/xnix-run-materials/state/known-run-version-under-test-7zr.json",
      "serial_log_output": "/home/xnix-run-materials/state/qemu-serial-version-under-test-7zr.log",
      "status": "passed",
      "app_version": "26.02",
      "executable_name": "7zr.exe",
      "backend": "guest-wine",
      "guest_started": true,
      "guest_port_auto": true,
      "checksum_verified": true,
      "raw_output_redacted": true,
      "marker_observed": true,
      "qemu_serial_log_written": true,
      "qemu_executed": true,
      "wine_executed": true
    },
    {
      "app_id": "busybox-w32",
      "report_output": "/home/xnix-run-materials/state/known-run-version-under-test-busybox-w32.json",
      "serial_log_output": "/home/xnix-run-materials/state/qemu-serial-version-under-test-busybox-w32.log",
      "status": "passed",
      "app_version": "current-2026-07-24",
      "executable_name": "busybox.exe",
      "backend": "guest-wine",
      "guest_started": true,
      "guest_port_auto": true,
      "checksum_verified": ` + checksumValue + `,
      "raw_output_redacted": true,
      "marker_observed": true,
      "qemu_serial_log_written": true,
      "qemu_executed": true,
      "wine_executed": true
    }
  ],
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "passed_count": 2,
  "failed_count": 0
}`
}
