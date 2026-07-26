package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewKnownExistingWinAppAcceptanceConsumes7zrRun(t *testing.T) {
	acceptance, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")))
	if err != nil {
		t.Fatalf("PreviewKnownExistingWinAppAcceptanceJSON returned error: %v", err)
	}
	if acceptance.SchemaVersion != KnownExistingWinAppAcceptanceSchemaVersion ||
		acceptance.RequestType != KnownExistingWinAppAcceptanceRequestType ||
		acceptance.AcceptanceType != "known-existing-windows-app-real-run-acceptance" ||
		acceptance.RunReportConsumed != true ||
		acceptance.RunReportPathExposed != false ||
		acceptance.RemoteHostExposed != false ||
		acceptance.GuestEndpointExposed != false ||
		acceptance.RawPathExposed != false ||
		acceptance.RawOutputExposed != false ||
		acceptance.RuntimeArgvExposed != false ||
		acceptance.RunnerPathExposed != false ||
		acceptance.AppID != "7zr" ||
		acceptance.DisplayName != "7-Zip standalone console executable" ||
		acceptance.AppVersion != "26.02" ||
		acceptance.Architecture != "windows-x86" ||
		acceptance.ExistingWindowsApp != true ||
		acceptance.KnownPortableCatalogBacked != true ||
		acceptance.BackendClass != "managed-isolated-compatibility" ||
		acceptance.LaunchAttempted != true ||
		acceptance.RunnerAvailable != true ||
		acceptance.ChecksumVerified != true ||
		acceptance.ArtifactCopied != true ||
		acceptance.MarkerObserved != true ||
		acceptance.GuestReachable != true ||
		acceptance.RuntimeStartedIsolatedGuest != true ||
		acceptance.IsolatedGuestExecutionObserved != true ||
		acceptance.CompatibilityEngineExecutionObserved != true ||
		acceptance.LoopbackOnlyNetworking != true ||
		acceptance.SerialLogPersisted != true ||
		acceptance.OutputRedacted != true ||
		acceptance.StdoutBytes != 2902 ||
		acceptance.StderrBytes != 170 ||
		acceptance.StdoutLineCount != 66 ||
		acceptance.StderrLineCount != 2 ||
		acceptance.NetworkRequired != false ||
		acceptance.HostRootModified != false ||
		acceptance.PrivilegedContainerRequired != false ||
		acceptance.HostNetworkingRequired != false ||
		acceptance.DockerSocketMounted != false ||
		acceptance.BroadHostMountRequired != false ||
		acceptance.DockerExecuted != false ||
		acceptance.ColimaExecuted != false ||
		acceptance.NetworkChecksRun != false ||
		acceptance.PackageManagerInvoked != false ||
		acceptance.AcceptanceReady != true {
		t.Fatalf("unexpected known existing Windows app acceptance: %#v", acceptance)
	}
	encoded, err := json.Marshal(acceptance)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	lower := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"root@q4", "/home/xnix-", "42281", "guest_port", "qemu-system", "7zr.exe", "wine ", ".wine"} {
		if strings.Contains(lower, forbidden) {
			t.Fatalf("known existing Windows app acceptance exposed forbidden term %q: %s", forbidden, lower)
		}
	}
}

func TestPreviewKnownExistingWinAppAcceptanceRejectsIncompleteEvidence(t *testing.T) {
	fixture := knownExistingWinAppAcceptanceFixture(currentProjectVersion(t), "7zr")
	internalFixture := strings.Replace(fixture, `"app_id": "7zr"`, `"app_id": "org.xnix.apps.messagebox"`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(internalFixture)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject internal fixture app ids")
	}
	missingChecksum := strings.Replace(fixture, `"checksum_verified": true`, `"checksum_verified": false`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(missingChecksum)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject missing checksum evidence")
	}
	missingMarker := strings.Replace(fixture, `"marker_observed": true`, `"marker_observed": false`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(missingMarker)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject missing marker evidence")
	}
	unredactedOutput := strings.Replace(fixture, `"raw_output_included": false`, `"raw_output_included": true`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(unredactedOutput)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject raw output inclusion")
	}
	unsafe := strings.Replace(fixture, `"docker_socket_mounted": false`, `"docker_socket_mounted": true`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(unsafe)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject unsafe host/container gates")
	}
	leakedRuntimeArgv := strings.Replace(fixture, `"raw_runtime_argv_exposed": false`, `"raw_runtime_argv_exposed": true`, 1)
	if _, err := PreviewKnownExistingWinAppAcceptanceJSON([]byte(leakedRuntimeArgv)); err == nil {
		t.Fatalf("known existing Windows app acceptance must reject raw runtime argv exposure")
	}
}

func knownExistingWinAppAcceptanceFixture(version string, appID string) string {
	fixture := strings.ReplaceAll(`{
  "schema_version": "xnix.runtime.known_windows_app_run.v1",
  "request_type": "windows-known-app-run",
  "version": "VERSION_PLACEHOLDER",
  "status": "passed",
  "app_id": "APP_ID_PLACEHOLDER",
  "display_name": "7-Zip standalone console executable",
  "app_version": "26.02",
  "architecture": "windows-x86",
  "executable_name": "7zr.exe",
  "backend": "guest-wine",
  "backend_ready": true,
  "launch_attempted": true,
  "runner_available": true,
  "checksum_verified": true,
  "guest_start_attempted": true,
  "guest_started": true,
  "guest_start_mode": "go-qemu",
  "guest_host": "127.0.0.1",
  "guest_port": 42281,
  "guest_port_auto": true,
  "qemu_serial_log_written": true,
  "executable_copied": true,
  "marker_observed": true,
  "application_workspace_mode": "direct-executable",
  "raw_output_redacted": true,
  "guest_payload": {
    "schema_version": "xnix.runtime.known_windows_app_guest_wine_smoke.v1",
    "request_type": "windows-known-app-guest-wine-smoke",
    "status": "passed",
    "checksum_verified": true,
    "expected_marker": "7-Zip",
    "guest": {
      "schema_version": "xnix.runtime.windows_app_guest_wine_smoke.v1",
      "request_type": "windows-app-guest-wine-smoke",
      "status": "passed",
      "guest_transport": "loopback-ssh",
      "guest_reachable": true,
      "wine_available": true,
      "executable_copied": true,
      "expected_marker": "7-Zip",
      "marker_observed": true,
      "exit_code": 0,
      "stdout": "",
      "stderr": "",
      "stdout_bytes": 2902,
      "stderr_bytes": 170,
      "stdout_line_count": 66,
      "stderr_line_count": 2,
      "raw_output_included": false,
      "raw_output_redacted": true,
      "loopback_only_networking": true,
      "qemu_required": true,
      "host_root_modified": false,
      "privileged_container_required": false,
      "host_networking_required": false,
      "docker_socket_mounted": false,
      "broad_host_mount_required": false,
      "raw_host_path_exposed": false
    }
  },
  "loopback_only_networking": true,
  "qemu_required": true,
  "network_required": false,
  "host_root_modified": false,
  "privileged_container_required": false,
  "host_networking_required": false,
  "docker_socket_mounted": false,
  "broad_host_mount_required": false,
  "docker_executed": false,
  "qemu_executed": true,
  "wine_executed": true,
  "colima_executed": false,
  "network_checks_run": false,
  "package_manager_invoked": false,
  "raw_host_path_exposed": false,
  "raw_executable_path_exposed": false,
  "raw_profile_path_exposed": false,
  "raw_state_root_path_exposed": false,
  "raw_runtime_argv_exposed": false,
  "raw_runner_path_exposed": false,
  "raw_qemu_path_exposed": false
}`, "VERSION_PLACEHOLDER", version)
	return strings.ReplaceAll(fixture, "APP_ID_PLACEHOLDER", appID)
}
