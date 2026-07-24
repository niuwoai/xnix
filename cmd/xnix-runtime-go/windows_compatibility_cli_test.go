package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestWindowsCompatibilityWorkstreamsPreviewCommand(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"windows-compatibility-workstreams-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_compatibility_workstreams.v1" ||
		payload["request_type"] != "windows-compatibility-workstreams-preview" ||
		payload["plan_type"] != "kde-first-windows-compatibility-workstreams" ||
		payload["source"] != "go-runtime-product-workstream-model" ||
		payload["read_method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		payload["official_desktop"] != "KDE Plasma" ||
		payload["product_target"] != "best Linux desktop for existing Windows applications" {
		t.Fatalf("unexpected workstream payload schema: %#v", payload)
	}
	if payload["entry_point_count"] != float64(7) ||
		payload["workstream_count"] != float64(11) ||
		payload["first_wave_count"] != float64(4) {
		t.Fatalf("unexpected workstream counts: %#v", payload)
	}
	entryIDs := payload["entry_point_ids"].([]any)
	if strings.Join(anyStrings(entryIDs), ",") != "start-menu,task-manager,file-manager,system-tray,notification-center,ai-compatibility-center,unified-settings" {
		t.Fatalf("unexpected entrypoint ids: %#v", entryIDs)
	}
	workstreamIDs := payload["workstream_ids"].([]any)
	if strings.Join(anyStrings(workstreamIDs), ",") != "CW1,CW2,CW3,CW4,CW5,CW6,CW7,CW8,CW9,CW10,CW11" {
		t.Fatalf("unexpected workstream ids: %#v", workstreamIDs)
	}
	firstWaveIDs := payload["first_wave_ids"].([]any)
	if strings.Join(anyStrings(firstWaveIDs), ",") != "CW1,CW2,CW3,CW10" {
		t.Fatalf("unexpected first-wave ids: %#v", firstWaveIDs)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["c_core_allowed"] != true ||
		payload["ruby_core_logic_allowed"] != false ||
		payload["ruby_test_harness"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_presentation_only"] != true ||
		payload["deep_desktop_fork_required"] != false ||
		payload["gnome_first_release_supported"] != false ||
		payload["xfce_first_release_supported"] != false ||
		payload["production_dbus_ownership_enabled"] != false ||
		payload["runtime_write_methods_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["portal_transport_calls_enabled"] != false ||
		payload["network_fetch_enabled"] != false ||
		payload["host_package_manager_enabled"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["compatibility_storage_path_exposed"] != false ||
		payload["backend_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected safety flags: %#v", payload)
	}
	if payload["protected_implementation_package_doc"] != "docs/claude-code-implementation-packages.md" {
		t.Fatalf("unexpected protected implementation package doc: %#v", payload["protected_implementation_package_doc"])
	}
	if !strings.Contains(payload["desktop_safe_summary"].(string), "CW1, CW2, CW3, and CW10") {
		t.Fatalf("desktop summary does not name the first wave: %#v", payload["desktop_safe_summary"])
	}
	assertWindowsCompatibilityCLISafe(t, output.String())
}

func TestWindowsCompatibilityWorkstreamsPreviewCommandRejectsArguments(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-compatibility-workstreams-preview", "--registry", "ignored"}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept arguments") {
		t.Fatalf("expected argument rejection, got %v", err)
	}
}

func TestWindowsAppRunSmokeCommandUsesRuntimeRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--runner-bottle", "private-bottle-name",
		"--runner-arg", "--shim-mode",
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_smoke.v1" ||
		payload["request_type"] != "windows-app-run-smoke" ||
		payload["status"] != "passed" ||
		payload["executable_name"] != "hello.exe" ||
		payload["executable_format"] != "pe-mz" ||
		payload["windows_executable_signature_observed"] != true ||
		payload["executable_architecture"] != "x86_64" ||
		payload["executable_architecture_supported"] != true ||
		payload["wine_architecture"] != "win64" ||
		payload["wine_prefix_mode"] != "architecture-scoped" ||
		payload["wine_prefix_prepared"] != true ||
		payload["runner_available"] != true ||
		payload["success_mode"] != "marker" ||
		payload["working_directory_mode"] != "executable-directory" ||
		payload["application_workspace_mode"] != "direct-executable" ||
		payload["application_staged"] != false ||
		payload["runner_argument_count"] != float64(3) ||
		payload["compatibility_layer"] != "windows-compatibility-layer" ||
		payload["wine_bootstrap_attempted"] != false ||
		payload["wine_bootstrap_succeeded"] != false ||
		payload["wine_bootstrap_skipped"] != false ||
		payload["wine_bootstrap_exit_code"] != float64(-1) ||
		payload["marker_observed"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false {
		t.Fatalf("unexpected smoke payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), runnerPath) ||
		strings.Contains(output.String(), "private-bottle-name") ||
		strings.Contains(output.String(), "--shim-mode") {
		t.Fatalf("smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppSmokeProfileRenderCommand(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"windows-app-smoke-profile-render",
		"--exe", "path/to/app.exe",
		"--working-dir", "path/to/app",
		"--state-root", ".local/xnix/winapp-smoke/profile-state",
		"--runner", "path/to/wine",
		"--runner-bottle", "operator-bottle",
		"--runner-arg", "--shim",
		"--arg", "--open",
		"--timeout", "45s",
		"--expected-marker", "APP_OK",
		"--success-mode", "exit-code",
		"--redact-output",
		"--skip-bootstrap",
		"--stage-app-dir",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_smoke_profile.v1" ||
		payload["executable_path"] != "path/to/app.exe" ||
		payload["working_directory"] != "path/to/app" ||
		payload["state_root"] != ".local/xnix/winapp-smoke/profile-state" ||
		payload["runner_path"] != "path/to/wine" ||
		payload["runner_bottle"] != "operator-bottle" ||
		payload["timeout"] != "45s" ||
		payload["expected_marker"] != "APP_OK" ||
		payload["success_mode"] != "exit-code" ||
		payload["redact_output"] != true ||
		payload["skip_bootstrap"] != true ||
		payload["stage_app_dir"] != true {
		t.Fatalf("unexpected profile render payload: %#v", payload)
	}
	runnerArgs := payload["runner_arguments"].([]any)
	appArgs := payload["arguments"].([]any)
	if len(runnerArgs) != 1 || runnerArgs[0] != "--shim" ||
		len(appArgs) != 1 || appArgs[0] != "--open" {
		t.Fatalf("unexpected profile render args: %#v", payload)
	}
}

func TestWindowsAppLauncherBundleRecordCommand(t *testing.T) {
	tempDir := t.TempDir()
	stateRoot := filepath.Join(tempDir, "state")
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeCLIJSON(t, profilePath, map[string]any{
		"schema_version":  "xnix.runtime.windows_app_smoke_profile.v1",
		"executable_path": filepath.Join(tempDir, "app", "hello.exe"),
		"state_root":      stateRoot,
		"runner_path":     filepath.Join(tempDir, "private-wine"),
		"runner_bottle":   "private-bottle",
		"runner_arguments": []string{
			"--private-runner-arg",
		},
		"skip_bootstrap": true,
		"timeout":        "30s",
		"stage_app_dir":  true,
	})

	var output bytes.Buffer
	err := run([]string{
		"windows-app-launcher-bundle-record",
		"--profile", profilePath,
		"--app-id", "org.xnix.realapp",
		"--name", "Real Windows App",
		"--runtime-bin", "go",
		"--runtime-arg", "run",
		"--runtime-arg", "./cmd/xnix-runtime-go",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_launcher_bundle.v1" ||
		payload["request_type"] != "windows-app-launcher-bundle-record" ||
		payload["status"] != "passed" ||
		payload["application_id"] != "org.xnix.realapp" ||
		payload["display_name"] != "Real Windows App" ||
		payload["desktop_file_name"] != "org.xnix.realapp.desktop" ||
		payload["launcher_script_name"] != "org.xnix.realapp.sh" ||
		payload["launcher_mode"] != "execute" ||
		payload["launcher_command"] != "windows-app-run-smoke" ||
		payload["runtime_argument_count"] != float64(2) ||
		payload["files_written"] != true ||
		payload["launcher_script_written"] != true ||
		payload["desktop_entry_written"] != true ||
		payload["receipt_written"] != true ||
		payload["desktop_entry_exec_uses_managed_launcher"] != true ||
		payload["raw_profile_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_runner_path_exposed"] != false ||
		payload["raw_runtime_argv_exposed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected launcher bundle payload: %#v", payload)
	}
	if strings.Contains(output.String(), profilePath) ||
		strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), filepath.Join(tempDir, "private-wine")) {
		t.Fatalf("launcher bundle output leaked private paths: %s", output.String())
	}
	desktopPath := filepath.Join(stateRoot, "launcher-bundle", "applications", "org.xnix.realapp.desktop")
	desktopText, err := os.ReadFile(desktopPath)
	if err != nil {
		t.Fatalf("ReadFile desktop entry returned error: %v", err)
	}
	if !strings.Contains(string(desktopText), "[Desktop Entry]") ||
		!strings.Contains(string(desktopText), "Exec=") ||
		!strings.Contains(string(desktopText), "X-Xnix-RuntimeOwned=true") {
		t.Fatalf("unexpected desktop entry: %s", string(desktopText))
	}

	var preflightOutput bytes.Buffer
	err = run([]string{
		"windows-app-launcher-bundle-record",
		"--profile", profilePath,
		"--app-id", "org.xnix.realapp.preflight",
		"--launcher-mode", "preflight",
	}, &preflightOutput)
	if err != nil {
		t.Fatalf("preflight run returned error: %v", err)
	}
	var preflightPayload map[string]any
	if err := json.Unmarshal(preflightOutput.Bytes(), &preflightPayload); err != nil {
		t.Fatalf("Unmarshal preflight returned error: %v", err)
	}
	if preflightPayload["status"] != "passed" ||
		preflightPayload["launcher_mode"] != "preflight" ||
		preflightPayload["launcher_command"] != "windows-app-smoke-profile-preflight" ||
		preflightPayload["raw_profile_path_exposed"] != false ||
		preflightPayload["raw_runtime_argv_exposed"] != false {
		t.Fatalf("unexpected preflight launcher payload: %#v", preflightPayload)
	}
	preflightLauncherPath := filepath.Join(stateRoot, "launcher-bundle", "launchers", "org.xnix.realapp.preflight.sh")
	preflightLauncherText, err := os.ReadFile(preflightLauncherPath)
	if err != nil {
		t.Fatalf("ReadFile preflight launcher returned error: %v", err)
	}
	if !strings.Contains(string(preflightLauncherText), "windows-app-smoke-profile-preflight --profile") ||
		strings.Contains(string(preflightLauncherText), "windows-app-run-smoke") {
		t.Fatalf("unexpected preflight launcher script: %s", string(preflightLauncherText))
	}

	var launchOutput bytes.Buffer
	err = run([]string{
		"windows-app-launcher-bundle-record",
		"--profile", profilePath,
		"--app-id", "org.xnix.realapp.launch",
		"--launcher-mode", "launch",
	}, &launchOutput)
	if err != nil {
		t.Fatalf("launch run returned error: %v", err)
	}
	var launchPayload map[string]any
	if err := json.Unmarshal(launchOutput.Bytes(), &launchPayload); err != nil {
		t.Fatalf("Unmarshal launch returned error: %v", err)
	}
	if launchPayload["status"] != "passed" ||
		launchPayload["launcher_mode"] != "launch" ||
		launchPayload["launcher_command"] != "windows-app-launch-profile" ||
		launchPayload["raw_profile_path_exposed"] != false ||
		launchPayload["raw_runtime_argv_exposed"] != false {
		t.Fatalf("unexpected launch profile launcher payload: %#v", launchPayload)
	}
	launchLauncherPath := filepath.Join(stateRoot, "launcher-bundle", "launchers", "org.xnix.realapp.launch.sh")
	launchLauncherText, err := os.ReadFile(launchLauncherPath)
	if err != nil {
		t.Fatalf("ReadFile launch launcher returned error: %v", err)
	}
	if !strings.Contains(string(launchLauncherText), "windows-app-launch-profile --profile") ||
		strings.Contains(string(launchLauncherText), "windows-app-run-smoke") {
		t.Fatalf("unexpected launch profile launcher script: %s", string(launchLauncherText))
	}
}

func TestWindowsAppLaunchProfileCommandBlocksBeforeExecutionWhenRunnerUnavailable(t *testing.T) {
	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeCLIJSON(t, profilePath, map[string]any{
		"schema_version":  "xnix.runtime.windows_app_smoke_profile.v1",
		"executable_path": exePath,
		"state_root":      stateRoot,
		"runner_path":     filepath.Join(tempDir, "missing-runner"),
		"timeout":         "5s",
		"stage_app_dir":   true,
	})

	var output bytes.Buffer
	err := run([]string{
		"windows-app-launch-profile",
		"--profile", profilePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_launch_profile.v1" ||
		payload["request_type"] != "windows-app-launch-profile" ||
		payload["status"] != "blocked" ||
		payload["preflight_status"] != "blocked" ||
		payload["launch_attempted"] != false ||
		payload["runtime_payload"] != nil ||
		payload["executable_name"] != "hello.exe" ||
		payload["executable_format"] != "pe-mz" ||
		payload["windows_executable_signature_observed"] != true ||
		payload["executable_architecture"] != "x86_64" ||
		payload["executable_architecture_supported"] != true ||
		payload["wine_architecture"] != "win64" ||
		payload["application_workspace_mode"] != "staged-application-directory" ||
		payload["runner_available"] != false ||
		payload["raw_output_redacted"] != true ||
		payload["raw_profile_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["wine_executed"] != false {
		t.Fatalf("unexpected launch profile payload: %#v", payload)
	}
	if strings.Contains(output.String(), profilePath) ||
		strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), exePath) {
		t.Fatalf("launch profile output leaked private paths: %s", output.String())
	}
	if _, err := os.Stat(stateRoot); err == nil {
		t.Fatalf("blocked launch profile command must not create state root before runner readiness")
	}
}

func TestWindowsAppRunSmokeCommandCanStageApplicationDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	appDir := filepath.Join(tempDir, "app")
	if err := os.Mkdir(appDir, 0o700); err != nil {
		t.Fatalf("Mkdir app dir returned error: %v", err)
	}
	exePath := filepath.Join(appDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	sidecarPath := filepath.Join(appDir, "sidecar.dat")
	if err := os.WriteFile(sidecarPath, []byte("sidecar"), 0o600); err != nil {
		t.Fatalf("WriteFile sidecar returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -f \"$(dirname \"$1\")/sidecar.dat\" || exit 77\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(appDir, "state"),
		"--runner", runnerPath,
		"--stage-app-dir",
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["application_workspace_mode"] != "staged-application-directory" ||
		payload["application_staged"] != true ||
		payload["application_staged_file_count"] != float64(2) ||
		payload["working_directory_mode"] != "staged-application-workspace" ||
		payload["marker_observed"] != true ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected staged application payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), sidecarPath) ||
		strings.Contains(output.String(), runnerPath) {
		t.Fatalf("staged smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanSkipBootstrap(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "wine")
	runnerBody := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	bootstrapMarker := filepath.Join(tempDir, "bootstrap.marker")
	winebootBody := "#!/bin/sh\n" +
		"printf should-not-bootstrap > '" + strings.ReplaceAll(bootstrapMarker, "'", "'\\''") + "'\n" +
		"exit 99\n"
	if err := os.WriteFile(filepath.Join(tempDir, "wineboot"), []byte(winebootBody), 0o700); err != nil {
		t.Fatalf("WriteFile wineboot returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--skip-bootstrap",
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["wine_bootstrap_attempted"] != false ||
		payload["wine_bootstrap_succeeded"] != false ||
		payload["wine_bootstrap_skipped"] != true ||
		payload["wine_bootstrap_exit_code"] != float64(-1) ||
		payload["marker_observed"] != true {
		t.Fatalf("unexpected skip-bootstrap payload: %#v", payload)
	}
	if _, err := os.Stat(bootstrapMarker); !os.IsNotExist(err) {
		t.Fatalf("wineboot marker must not exist when bootstrap is skipped: %v", err)
	}
}

func TestWindowsAppRunSmokeCommandRejectsNonPEExecutable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, []byte("plain text"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nprintf 'should-not-run\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
	}, &output)
	if err == nil || !strings.Contains(err.Error(), "Windows PE file") {
		t.Fatalf("expected PE validation error, got err=%v output=%s", err, output.String())
	}
	if output.Len() != 0 {
		t.Fatalf("non-PE executable should be rejected before JSON smoke output, got %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandReportsUnsupportedArchitecture(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "arm-app.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0xaa64), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nprintf 'should-not-run\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "failed" ||
		payload["executable_name"] != "arm-app.exe" ||
		payload["executable_format"] != "pe-mz" ||
		payload["windows_executable_signature_observed"] != true ||
		payload["executable_architecture"] != "arm64" ||
		payload["executable_architecture_supported"] != false ||
		payload["wine_architecture"] != "unknown" ||
		payload["wine_prefix_mode"] != "unknown" ||
		payload["wine_prefix_prepared"] != false ||
		payload["failure_reason"] != "Windows executable architecture is not supported" ||
		payload["runner_available"] != false ||
		payload["isolated_state_root"] != false {
		t.Fatalf("unexpected unsupported architecture payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), runnerPath) {
		t.Fatalf("unsupported architecture output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanUseOperatorWorkingDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	workingDir := filepath.Join(tempDir, "runtime-cwd")
	if err := os.Mkdir(workingDir, 0o700); err != nil {
		t.Fatalf("Mkdir working dir returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workingDir, "sidecar.ini"), []byte("sidecar"), 0o600); err != nil {
		t.Fatalf("WriteFile sidecar returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -f sidecar.ini || exit 74\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--working-dir", workingDir,
		"--runner", runnerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["working_directory_mode"] != "operator-supplied" ||
		payload["marker_observed"] != true {
		t.Fatalf("unexpected working directory payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), runnerPath) ||
		strings.Contains(output.String(), workingDir) {
		t.Fatalf("working directory smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanUseProfile(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "profile-app.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	workingDir := filepath.Join(tempDir, "profile-cwd")
	if err := os.Mkdir(workingDir, 0o700); err != nil {
		t.Fatalf("Mkdir working dir returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workingDir, "profile-sidecar.ini"), []byte("sidecar"), 0o600); err != nil {
		t.Fatalf("WriteFile sidecar returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test \"$1\" = \"--shim-mode\" || exit 73\n" +
		"test \"$3\" = \"--profile-arg\" || exit 74\n" +
		"test -f profile-sidecar.ini || exit 75\n" +
		"printf 'CUSTOM_PROFILE_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "profile.json")
	profile := map[string]any{
		"schema_version":    "xnix.runtime.windows_app_smoke_profile.v1",
		"executable_path":   exePath,
		"state_root":        filepath.Join(tempDir, "state"),
		"working_directory": workingDir,
		"runner_path":       runnerPath,
		"runner_arguments":  []string{"--shim-mode"},
		"arguments":         []string{"--profile-arg"},
		"timeout":           "5s",
		"expected_marker":   "CUSTOM_PROFILE_OK",
		"success_mode":      "marker",
	}
	profileData, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("Marshal profile returned error: %v", err)
	}
	if err := os.WriteFile(profilePath, profileData, 0o600); err != nil {
		t.Fatalf("WriteFile profile returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"windows-app-run-smoke",
		"--profile", profilePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["executable_name"] != "profile-app.exe" ||
		payload["expected_marker"] != "CUSTOM_PROFILE_OK" ||
		payload["working_directory_mode"] != "operator-supplied" ||
		payload["runner_argument_count"] != float64(1) ||
		payload["marker_observed"] != true {
		t.Fatalf("unexpected profile smoke payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), runnerPath) ||
		strings.Contains(output.String(), workingDir) ||
		strings.Contains(output.String(), profilePath) ||
		strings.Contains(output.String(), "--shim-mode") {
		t.Fatalf("profile smoke output leaked host paths or runner args: %s", output.String())
	}
}

func TestWindowsAppSmokeProfilePreflightCommandReportsReady(t *testing.T) {
	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "profile-app.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	workingDir := filepath.Join(tempDir, "profile-cwd")
	if err := os.Mkdir(workingDir, 0o700); err != nil {
		t.Fatalf("Mkdir working dir returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nexit 99\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "profile.json")
	profileData, err := json.Marshal(map[string]any{
		"schema_version":    "xnix.runtime.windows_app_smoke_profile.v1",
		"executable_path":   exePath,
		"state_root":        filepath.Join(tempDir, "state"),
		"working_directory": workingDir,
		"runner_path":       runnerPath,
		"runner_arguments":  []string{"--private-shim"},
		"arguments":         []string{"--profile-arg"},
		"timeout":           "5s",
		"expected_marker":   "CUSTOM_PROFILE_OK",
		"success_mode":      "marker",
	})
	if err != nil {
		t.Fatalf("Marshal profile returned error: %v", err)
	}
	if err := os.WriteFile(profilePath, profileData, 0o600); err != nil {
		t.Fatalf("WriteFile profile returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"windows-app-smoke-profile-preflight",
		"--profile", profilePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_smoke_profile_preflight.v1" ||
		payload["request_type"] != "windows-app-smoke-profile-preflight" ||
		payload["status"] != "ready" ||
		payload["profile_supplied"] != true ||
		payload["executable_name"] != "profile-app.exe" ||
		payload["executable_format"] != "pe-mz" ||
		payload["windows_executable_signature_observed"] != true ||
		payload["executable_architecture"] != "x86_64" ||
		payload["executable_architecture_supported"] != true ||
		payload["wine_architecture"] != "win64" ||
		payload["wine_prefix_mode"] != "architecture-scoped" ||
		payload["wine_prefix_prepared"] != false ||
		payload["working_directory_mode"] != "operator-supplied" ||
		payload["runner_available"] != true ||
		payload["runner_argument_count"] != float64(1) ||
		payload["app_argument_count"] != float64(1) ||
		payload["wine_executed"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected profile preflight payload: %#v", payload)
	}
	if strings.Contains(output.String(), profilePath) ||
		strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), workingDir) ||
		strings.Contains(output.String(), runnerPath) ||
		strings.Contains(output.String(), "--private-shim") {
		t.Fatalf("profile preflight output leaked private values: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanUseExitCodeSuccessMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nprintf 'GUI app exited cleanly\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--success-mode", "exit-code",
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["success_mode"] != "exit-code" ||
		payload["marker_observed"] != false ||
		payload["failure_reason"] != nil {
		t.Fatalf("unexpected exit-code success payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) || strings.Contains(output.String(), runnerPath) {
		t.Fatalf("exit-code smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanUseStartupWindowSuccessMode(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nsleep 1\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--success-mode", "startup-window",
		"--timeout", "50ms",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["success_mode"] != "startup-window" ||
		payload["startup_window_observed"] != true ||
		payload["marker_observed"] != false ||
		payload["failure_reason"] != nil {
		t.Fatalf("unexpected startup-window payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) || strings.Contains(output.String(), runnerPath) {
		t.Fatalf("startup-window smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppRunnerDiagnosticsCommandUsesExplicitRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nprintf 'ready\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-runner-diagnostics",
		"--runner", runnerPath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_runner_diagnostics.v1" ||
		payload["request_type"] != "windows-app-runner-diagnostics" ||
		payload["status"] != "passed" ||
		payload["runner_available"] != true ||
		payload["explicit_runner_supplied"] != true ||
		payload["selected_runner_name"] != "fake-runner" ||
		payload["raw_path_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["package_manager_invoked"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["colima_executed"] != false {
		t.Fatalf("unexpected runner diagnostics payload: %#v", payload)
	}
	commandHints, ok := payload["runner_command_hints"].([]any)
	if !ok || len(commandHints) != 3 || !strings.Contains(fmt.Sprint(commandHints[0]), "path/to/app.exe") {
		t.Fatalf("unexpected runner command hints: %#v", payload["runner_command_hints"])
	}
	if strings.Contains(output.String(), runnerPath) || strings.Contains(output.String(), tempDir) {
		t.Fatalf("runner diagnostics leaked raw host paths: %s", output.String())
	}
}

func TestWindowsAppRunnerDiagnosticsCommandUsesEnvRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nprintf 'ready\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	t.Setenv("XNIX_WINDOWS_RUNNER", runnerPath)

	var output bytes.Buffer
	err := run([]string{"windows-app-runner-diagnostics"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["runner_available"] != true ||
		payload["explicit_runner_supplied"] != false ||
		payload["env_runner_configured"] != true ||
		payload["selected_runner_name"] != "fake-runner" ||
		payload["raw_path_exposed"] != false {
		t.Fatalf("unexpected env runner diagnostics payload: %#v", payload)
	}
	if strings.Contains(output.String(), runnerPath) || strings.Contains(output.String(), tempDir) {
		t.Fatalf("env runner diagnostics leaked raw host paths: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandCanRedactRawOutput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\nraw-host-path=/private/tmp/secret\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--timeout", "5s",
		"--redact-output",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["marker_observed"] != true ||
		payload["stdout"] != "" ||
		payload["stderr"] != "" ||
		payload["raw_output_included"] != false ||
		payload["raw_output_redacted"] != true ||
		payload["stdout_bytes"] == float64(0) ||
		payload["stdout_line_count"] != float64(2) ||
		!strings.Contains(payload["kde_safe_output_summary"].(string), "expected smoke marker observed") {
		t.Fatalf("unexpected redacted smoke payload: %#v", payload)
	}
	if strings.Contains(output.String(), "/private/tmp") || strings.Contains(output.String(), "secret") || strings.Contains(output.String(), runnerPath) || strings.Contains(output.String(), exePath) {
		t.Fatalf("redacted smoke output leaked raw values: %s", output.String())
	}
}

func TestWindowsAppRunSmokeCommandUsesCustomExpectedMarker(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "custom.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf 'CUSTOM_APP_OK\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", runnerPath,
		"--timeout", "5s",
		"--expected-marker", "CUSTOM_APP_OK",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["marker_observed"] != true ||
		payload["expected_marker"] != "CUSTOM_APP_OK" ||
		payload["executable_name"] != "custom.exe" {
		t.Fatalf("unexpected custom marker payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) || strings.Contains(output.String(), runnerPath) {
		t.Fatalf("custom marker smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppContainerRunSmokeCommandUsesRestrictedRuntimeRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell docker fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 0; fi\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\n'\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"windows-app-container-run-smoke",
		"--exe", exePath,
		"--state-root", filepath.Join(tempDir, "state"),
		"--image", "local/wine-smoke:test",
		"--platform", "linux/amd64",
		"--docker", dockerPath,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_container_smoke.v1" ||
		payload["request_type"] != "windows-app-container-run-smoke" ||
		payload["status"] != "passed" ||
		payload["executable_name"] != "hello.exe" ||
		payload["container_image"] != "local/wine-smoke:test" ||
		payload["container_platform"] != "linux/amd64" ||
		payload["pull_policy"] != "never" ||
		payload["network_mode"] != "none" ||
		payload["runner_available"] != true ||
		payload["image_available"] != true ||
		payload["compatibility_layer"] != "containerized-windows-compatibility-layer" ||
		payload["marker_observed"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["host_mount_count"] != float64(1) ||
		payload["container_state_mode"] != "tmpfs" ||
		payload["wine_bootstrap_required"] != true ||
		payload["wine_bootstrap_timed_out"] != false ||
		payload["wine_bootstrap_exit_code"] != float64(-1) {
		t.Fatalf("unexpected container smoke payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) || strings.Contains(output.String(), dockerPath) {
		t.Fatalf("container smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppGuestWineSmokeCommandUsesLoopbackGuestRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	exePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	logPath := filepath.Join(tempDir, "guest.log")
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *' wine '*'hello.exe'*) printf 'KNOWN_PORTABLE_APP_OK\\n'; exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	scpPath := filepath.Join(tempDir, "fake-scp")
	scpBody := "#!/bin/sh\n" +
		"printf 'scp %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"exit 0\n"
	if err := os.WriteFile(scpPath, []byte(scpBody), 0o700); err != nil {
		t.Fatalf("WriteFile scp returned error: %v", err)
	}
	keyPath := filepath.Join(tempDir, "id_ed25519")

	var output bytes.Buffer
	err := run([]string{
		"windows-app-guest-wine-smoke",
		"--exe", exePath,
		"--host", "127.0.0.1",
		"--port", "2222",
		"--user", "root",
		"--key", keyPath,
		"--remote-dir", "/tmp/xnix-winapp-smoke",
		"--ssh", sshPath,
		"--scp", scpPath,
		"--timeout", "5s",
		"--expected-marker", "KNOWN_PORTABLE_APP_OK",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_guest_wine_smoke.v1" ||
		payload["request_type"] != "windows-app-guest-wine-smoke" ||
		payload["status"] != "passed" ||
		payload["executable_name"] != "hello.exe" ||
		payload["guest_transport"] != "loopback-ssh" ||
		payload["guest_reachable"] != true ||
		payload["wine_available"] != true ||
		payload["executable_copied"] != true ||
		payload["expected_marker"] != "KNOWN_PORTABLE_APP_OK" ||
		payload["marker_observed"] != true ||
		payload["loopback_only_networking"] != true ||
		payload["qemu_required"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false {
		t.Fatalf("unexpected guest smoke payload: %#v", payload)
	}
	if strings.Contains(output.String(), exePath) ||
		strings.Contains(output.String(), sshPath) ||
		strings.Contains(output.String(), scpPath) ||
		strings.Contains(output.String(), keyPath) {
		t.Fatalf("guest smoke output leaked host paths: %s", output.String())
	}
}

func TestWindowsAppGuestWineGUISmokeCommandObservesWindow(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "guest-gui.log")
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *'wineboot --init'*) printf 'boot initialized\\n' >&2; exit 0 ;;\n" +
		"  *'wine '*'winemine.exe'*) exit 0 ;;\n" +
		"  *'cat '*'stderr.txt'*) printf ''; exit 0 ;;\n" +
		"  *'wineserver -k'*) exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	xwininfoPath := filepath.Join(tempDir, "fake-xwininfo")
	xwininfoBody := "#!/bin/sh\n" +
		"printf 'xwininfo display=%s\\n' \"$DISPLAY\" >> '" + logPath + "'\n" +
		"printf 'xwininfo: Window id: 0x3a7 (the root window)\\n'\n" +
		"printf '  0x200001 \"WineMine\": ()  320x240+0+0  +0+0\\n'\n"
	if err := os.WriteFile(xwininfoPath, []byte(xwininfoBody), 0o700); err != nil {
		t.Fatalf("WriteFile xwininfo returned error: %v", err)
	}
	keyPath := filepath.Join(tempDir, "id_ed25519")
	guestDisplay := strings.Join([]string{"10", "0", "2", "2"}, ".") + ":100"

	var output bytes.Buffer
	err := run([]string{
		"windows-app-guest-wine-gui-smoke",
		"--gui-app", "/usr/lib/wine/i386-windows/winemine.exe",
		"--host", "127.0.0.1",
		"--port", "2222",
		"--user", "root",
		"--key", keyPath,
		"--remote-dir", "/tmp/xnix-wine-guest-gui-smoke",
		"--ssh", sshPath,
		"--xwininfo", xwininfoPath,
		"--guest-display", guestDisplay,
		"--host-display", ":100",
		"--timeout", "5s",
		"--wait", "1ms",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_app_guest_wine_gui_smoke.v1" ||
		payload["request_type"] != "windows-app-guest-wine-gui-smoke" ||
		payload["status"] != "passed" ||
		payload["gui_app_name"] != "winemine.exe" ||
		payload["backend"] != "qemu-guest-wine-x11" ||
		payload["wineboot_invoked"] != true ||
		payload["launch_attempted"] != true ||
		payload["xwininfo_invoked"] != true ||
		payload["x_window_observed"] != true ||
		payload["x_window_child_count"] != float64(1) ||
		payload["loopback_ssh_forwarding_only"] != true ||
		payload["qemu_required"] != true ||
		payload["xvfb_required"] != true ||
		payload["qemu_user_network_restrict_disabled_for_display"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_guest_gui_app_path_exposed"] != false ||
		payload["raw_command_exposed"] != false {
		t.Fatalf("unexpected guest GUI smoke payload: %#v", payload)
	}
	for _, forbidden := range []string{sshPath, xwininfoPath, keyPath, "/usr/lib/wine/i386-windows/winemine.exe"} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("guest GUI smoke output leaked raw path %q: %s", forbidden, output.String())
		}
	}
}

func TestWindowsAppGuestWineGUISmokeCommandCopiesLocalExecutable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "guest-gui-copy.log")
	exePath := filepath.Join(tempDir, "hello-gui.exe")
	if err := os.WriteFile(exePath, minimalPEFixture(0x014c), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	sshPath := filepath.Join(tempDir, "fake-ssh")
	sshBody := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *'wineboot --init'*) printf 'boot initialized\\n' >&2; exit 0 ;;\n" +
		"  *'wine '*'hello-gui.exe'*) exit 0 ;;\n" +
		"  *'cat '*'stderr.txt'*) printf ''; exit 0 ;;\n" +
		"  *'wineserver -k'*) exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(sshPath, []byte(sshBody), 0o700); err != nil {
		t.Fatalf("WriteFile ssh returned error: %v", err)
	}
	scpPath := filepath.Join(tempDir, "fake-scp")
	scpBody := "#!/bin/sh\n" +
		"printf 'scp %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"exit 0\n"
	if err := os.WriteFile(scpPath, []byte(scpBody), 0o700); err != nil {
		t.Fatalf("WriteFile scp returned error: %v", err)
	}
	xwininfoPath := filepath.Join(tempDir, "fake-xwininfo")
	xwininfoBody := "#!/bin/sh\n" +
		"printf 'xwininfo display=%s\\n' \"$DISPLAY\" >> '" + logPath + "'\n" +
		"printf 'xwininfo: Window id: 0x3a7 (the root window)\\n'\n" +
		"printf '  0x200001 \"Xnix Hello GUI\": ()  320x240+0+0  +0+0\\n'\n"
	if err := os.WriteFile(xwininfoPath, []byte(xwininfoBody), 0o700); err != nil {
		t.Fatalf("WriteFile xwininfo returned error: %v", err)
	}
	keyPath := filepath.Join(tempDir, "id_ed25519")

	var output bytes.Buffer
	err := run([]string{
		"windows-app-guest-wine-gui-smoke",
		"--executable", exePath,
		"--host", "127.0.0.1",
		"--port", "2222",
		"--user", "root",
		"--key", keyPath,
		"--remote-dir", "/tmp/xnix-wine-guest-gui-smoke",
		"--ssh", sshPath,
		"--scp", scpPath,
		"--xwininfo", xwininfoPath,
		"--guest-display", "10.0.2.2:100",
		"--host-display", ":100",
		"--timeout", "5s",
		"--wait", "1ms",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "passed" ||
		payload["gui_app_name"] != "hello-gui.exe" ||
		payload["executable_copied"] != true ||
		payload["launch_attempted"] != true ||
		payload["x_window_observed"] != true ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_guest_gui_app_path_exposed"] != false ||
		payload["raw_command_exposed"] != false {
		t.Fatalf("unexpected copied guest GUI smoke payload: %#v", payload)
	}
	for _, forbidden := range []string{exePath, sshPath, scpPath, xwininfoPath, keyPath, "/tmp/xnix-wine-guest-gui-smoke/hello-gui.exe"} {
		if strings.Contains(output.String(), forbidden) {
			t.Fatalf("guest GUI smoke output leaked raw path %q: %s", forbidden, output.String())
		}
	}
	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "scp ") ||
		!strings.Contains(log, exePath) ||
		!strings.Contains(log, "root@127.0.0.1:/tmp/xnix-wine-guest-gui-smoke/hello-gui.exe") ||
		!strings.Contains(log, "wine '/tmp/xnix-wine-guest-gui-smoke/hello-gui.exe'") {
		t.Fatalf("guest GUI smoke command did not copy and launch expected executable: %s", log)
	}
}

func TestWindowsKnownAppGuestWineSmokeCommandSkipsUntilArtifactIsFetched(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-guest-wine-smoke",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_guest_wine_smoke.v1" ||
		payload["request_type"] != "windows-known-app-guest-wine-smoke" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["app_version"] != "26.02" ||
		payload["architecture"] != "windows-x86" ||
		payload["executable_name"] != "7zr.exe" ||
		payload["expected_marker"] != "7-Zip" ||
		payload["checksum_verified"] != false ||
		payload["loopback_only_networking"] != true ||
		payload["qemu_required"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false {
		t.Fatalf("unexpected known app guest payload: %#v", payload)
	}
	if strings.Contains(output.String(), tempDir) {
		t.Fatalf("known app guest output leaked host paths: %s", output.String())
	}
}

func TestWindowsKnownAppManagedLaunchPreviewCommandKeepsDesktopSurfaceRedacted(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-managed-launch-preview",
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_managed_launch.v1" ||
		payload["request_type"] != "windows-known-app-managed-launch-preview" ||
		payload["status"] != "needs-artifact" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["app_version"] != "26.02" ||
		payload["architecture"] != "windows-x86" ||
		payload["launch_surface_id"] != "known-app-7zr" ||
		payload["desktop_action_id"] != "launch-known-app-7zr" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_enabled"] != false ||
		payload["preparation_required"] != true ||
		payload["managed_launch_surface"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["real_app_smoke_gate_required"] != true ||
		payload["real_app_smoke_gate"] != "managed-known-app-guest-smoke" ||
		payload["loopback_only_networking"] != true ||
		payload["guest_runtime_required"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected managed launch payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if strings.Join(anyStrings(argv), " ") != "xnix-compat-launch --app 7zr" {
		t.Fatalf("unexpected managed launcher argv: %#v", argv)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppKDELauncherPreviewCommandConsumesManagedLaunchSurface(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-kde-launcher-preview",
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_kde_launcher.v1" ||
		payload["request_type"] != "windows-known-app-kde-launcher-preview" ||
		payload["source"] != "windows-known-app-managed-launch-preview" ||
		payload["status"] != "visible-needs-preparation" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["entry_point_id"] != "launcher" ||
		payload["kde_component"] != "Plasma application launcher" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["desktop_file"] != "xnix-known-app-7zr.desktop" ||
		payload["icon"] != "xnix-known-app-7zr" ||
		payload["launch_surface_id"] != "known-app-7zr" ||
		payload["desktop_action_id"] != "launch-known-app-7zr" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_visible"] != true ||
		payload["launch_enabled"] != false ||
		payload["preparation_required"] != true ||
		payload["managed_launch_surface"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["desktop_entry_preview_created"] != true ||
		payload["desktop_files_written"] != false ||
		payload["mimeapps_written"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected KDE launcher payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if strings.Join(anyStrings(argv), " ") != "xnix-compat-launch --app 7zr" {
		t.Fatalf("unexpected KDE launcher argv: %#v", argv)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppLaunchRequestPreviewCommandCreatesRuntimeOwnedRequest(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-launch-request-preview",
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_launch_request.v1" ||
		payload["request_type"] != "windows-known-app-launch-request-preview" ||
		payload["source"] != "windows-known-app-kde-launcher-preview" ||
		payload["status"] != "request-blocked" ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["runtime_method"] != "LaunchKnownWindowsApp" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["entry_point_id"] != "launcher" ||
		payload["desktop_file"] != "xnix-known-app-7zr.desktop" ||
		payload["launch_surface_id"] != "known-app-7zr" ||
		payload["desktop_action_id"] != "launch-known-app-7zr" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_visible"] != true ||
		payload["launch_allowed"] != false ||
		payload["launch_request_created"] != true ||
		payload["dispatch_ready"] != false ||
		payload["preparation_required"] != true ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["dry_run"] != true ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch request payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if strings.Join(anyStrings(argv), " ") != "xnix-compat-launch --app 7zr" {
		t.Fatalf("unexpected launch request argv: %#v", argv)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppDispatchPreviewCommandMapsLaunchRequestToManagedLane(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-dispatch-preview",
		"--app", "7zr",
		"--cache-root", tempDir,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_dispatch.v1" ||
		payload["request_type"] != "windows-known-app-dispatch-preview" ||
		payload["source"] != "windows-known-app-launch-request-preview" ||
		payload["status"] != "dispatch-blocked" ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["dispatch_id"] != "known-app-dispatch-7zr" ||
		payload["runtime_method"] != "DispatchKnownWindowsApp" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["entry_point_id"] != "launcher" ||
		payload["desktop_file"] != "xnix-known-app-7zr.desktop" ||
		payload["launch_surface_id"] != "known-app-7zr" ||
		payload["desktop_action_id"] != "launch-known-app-7zr" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["runner_lane"] != "known-app-guest-smoke" ||
		payload["runner_request_type"] != "managed-known-app-guest-smoke" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_request_created"] != true ||
		payload["dispatch_preview_created"] != true ||
		payload["dispatch_allowed"] != false ||
		payload["dispatch_ready"] != false ||
		payload["preparation_required"] != true ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["dry_run"] != true ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected dispatch payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if strings.Join(anyStrings(argv), " ") != "xnix-compat-launch --app 7zr" {
		t.Fatalf("unexpected dispatch argv: %#v", argv)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppDispatchSmokeCommandBlocksBeforeArtifactPreparation(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-dispatch-smoke",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--guest-boundary", "managed-known-app-guest-smoke",
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_dispatch_smoke.v1" ||
		payload["request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["source"] != "windows-known-app-dispatch-preview" ||
		payload["status"] != "dispatch-blocked" ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["dispatch_id"] != "known-app-dispatch-7zr" ||
		payload["runtime_method"] != "DispatchKnownWindowsApp" ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["runner_lane"] != "known-app-guest-smoke" ||
		payload["guest_boundary"] != "managed-known-app-guest-smoke" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_request_created"] != true ||
		payload["dispatch_preview_created"] != true ||
		payload["dispatch_ready"] != false ||
		payload["dispatch_allowed"] != false ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["managed_guest_runner_invoked"] != false ||
		payload["managed_guest_reachable"] != false ||
		payload["managed_guest_runtime_ready"] != false ||
		payload["managed_artifact_copied"] != false ||
		payload["marker_observed"] != false ||
		payload["smoke_passed"] != false ||
		payload["exit_code"] != float64(-1) ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected dispatch smoke payload: %#v", payload)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppLaunchBridgePreviewCommandMaterializesManagedRequest(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-launch-bridge-preview",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--launcher-arg=xnix-compat-launch",
		"--launcher-arg=--app",
		"--launcher-arg=7zr",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_launch_bridge.v1" ||
		payload["request_type"] != "windows-known-app-launch-bridge-preview" ||
		payload["source"] != "windows-known-app-dispatch-preview" ||
		payload["status"] != "bridge-blocked" ||
		payload["desktop"] != "KDE Plasma" ||
		payload["entry_point_id"] != "launcher" ||
		payload["desktop_file"] != "xnix-known-app-7zr.desktop" ||
		payload["launch_surface_id"] != "known-app-7zr" ||
		payload["desktop_action_id"] != "launch-known-app-7zr" ||
		payload["managed_launcher"] != "xnix-compat-launch --app 7zr" ||
		payload["launcher_argv_accepted"] != true ||
		payload["request_id"] != "known-app-launch-request-7zr" ||
		payload["dispatch_id"] != "known-app-dispatch-7zr" ||
		payload["runtime_method"] != "BridgeKnownLauncherToDispatchSmoke" ||
		payload["dispatch_smoke_request_type"] != "windows-known-app-dispatch-smoke" ||
		payload["dispatch_smoke_request_materialized"] != false ||
		payload["app_id"] != "7zr" ||
		payload["display_name"] != "7-Zip standalone console executable" ||
		payload["dispatch_gate"] != "managed-known-app-guest-smoke" ||
		payload["runner_lane"] != "known-app-guest-smoke" ||
		payload["guest_boundary_required"] != true ||
		payload["guest_boundary_supplied"] != false ||
		payload["smoke_harness_required"] != true ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["launch_request_created"] != true ||
		payload["dispatch_preview_created"] != true ||
		payload["bridge_preview_created"] != true ||
		payload["dispatch_ready"] != false ||
		payload["preparation_required"] != true ||
		payload["runtime_owned_request"] != true ||
		payload["runtime_owned_launch"] != true ||
		payload["runtime_owned_dispatch"] != true ||
		payload["runtime_owned_bridge"] != true ||
		payload["kde_presentation_only"] != true ||
		payload["dry_run"] != true ||
		payload["dispatch_started"] != false ||
		payload["execution_started"] != false ||
		payload["backend_process_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected launch bridge payload: %#v", payload)
	}
	argv := payload["managed_launcher_argv"].([]any)
	if strings.Join(anyStrings(argv), " ") != "xnix-compat-launch --app 7zr" {
		t.Fatalf("unexpected launch bridge argv: %#v", argv)
	}
	assertKnownManagedLaunchCLISafe(t, output.String(), tempDir)
}

func TestWindowsKnownAppFetchCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-fetch", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppGuestWineSmokeCommandRecognizesBusyBoxW32(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-guest-wine-smoke",
		"--app", "busybox-w32",
		"--cache-root", tempDir,
		"--timeout", "5s",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_guest_wine_smoke.v1" ||
		payload["request_type"] != "windows-known-app-guest-wine-smoke" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "busybox-w32" ||
		payload["display_name"] != "BusyBox-w32 standalone console executable" ||
		payload["app_version"] != "current-2026-07-24" ||
		payload["architecture"] != "windows-x86" ||
		payload["executable_name"] != "busybox.exe" ||
		payload["expected_sha256"] != "7bfee530965315665044e6e01db58125f2763c8a39c2e72ba1a6beb6923e0e1f" ||
		payload["expected_marker"] != "BusyBox" ||
		payload["checksum_verified"] != false ||
		payload["loopback_only_networking"] != true ||
		payload["qemu_required"] != true ||
		payload["host_root_modified"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["raw_host_path_exposed"] != false {
		t.Fatalf("unexpected BusyBox-w32 guest payload: %#v", payload)
	}
	if strings.Contains(output.String(), tempDir) {
		t.Fatalf("known app guest output leaked host paths: %s", output.String())
	}
}

func TestWindowsKnownAppLaunchProfileMaterializeCommandSkipsMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-launch-profile-materialize",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runtime-bin", "go",
		"--runtime-arg", "run",
		"--runtime-arg", "./cmd/xnix-runtime-go",
		"--runner", "/private/runner",
		"--runner-bottle", "private-bottle",
		"--runner-arg", "--private-runner-arg",
		"--skip-bootstrap",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_launch_profile_materialize.v1" ||
		payload["request_type"] != "windows-known-app-launch-profile-materialize" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["cache_status"] != "missing" ||
		payload["artifact_verified"] != false ||
		payload["profile_written"] != false ||
		payload["launcher_bundle_written"] != false ||
		payload["launcher_mode"] != "launch" ||
		payload["launcher_command"] != "windows-app-launch-profile" ||
		payload["runner_configured"] != true ||
		payload["runner_bottle_configured"] != true ||
		payload["runner_argument_count"] != float64(3) ||
		payload["skip_bootstrap"] != true ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_profile_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_runtime_argv_exposed"] != false ||
		payload["skip_reason"] != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected known launch profile materialize payload: %#v", payload)
	}
	assertCLIOutputOmitsValues(t, output.String(), "/private/runner", "private-bottle", "--private-runner-arg")
}

func TestWindowsKnownAppPrepareLaunchProfileCommandSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-prepare-launch-profile",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runtime-bin", "go",
		"--runtime-arg", "run",
		"--runtime-arg", "./cmd/xnix-runtime-go",
		"--runner", "/private/runner",
		"--runner-bottle", "private-bottle",
		"--runner-arg", "--private-runner-arg",
		"--skip-bootstrap",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_prepare_launch_profile.v1" ||
		payload["request_type"] != "windows-known-app-prepare-launch-profile" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["allow_download"] != false ||
		payload["network_required"] != false ||
		payload["fetch_status"] != "skipped" ||
		payload["fetch_cache_status"] != "missing" ||
		payload["downloaded"] != false ||
		payload["checksum_verified"] != false ||
		payload["materialize_status"] != "not-run" ||
		payload["profile_written"] != false ||
		payload["launcher_bundle_written"] != false ||
		payload["runner_configured"] != true ||
		payload["runner_bottle_configured"] != true ||
		payload["runner_argument_count"] != float64(3) ||
		payload["skip_bootstrap"] != true ||
		payload["host_root_modified"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["raw_profile_path_exposed"] != false ||
		payload["raw_state_root_path_exposed"] != false ||
		payload["raw_runtime_argv_exposed"] != false ||
		payload["skip_reason"] != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected known prepare launch profile payload: %#v", payload)
	}
	fetchPayload := payload["fetch_payload"].(map[string]any)
	if fetchPayload["network_required"] != false {
		t.Fatalf("offline prepare fetch payload must not require network: %#v", fetchPayload)
	}
	assertCLIOutputOmitsValues(t, output.String(), "/private/runner", "private-bottle", "--private-runner-arg")
}

func TestWindowsKnownAppPrepareAndLaunchProfileCommandSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-prepare-and-launch-profile",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runtime-bin", "go",
		"--runtime-arg", "run",
		"--runtime-arg", "./cmd/xnix-runtime-go",
		"--runner", "/private/runner",
		"--runner-bottle", "private-bottle",
		"--runner-arg", "--private-runner-arg",
		"--skip-bootstrap",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_prepare_and_launch_profile.v1" ||
		payload["request_type"] != "windows-known-app-prepare-and-launch-profile" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["allow_download"] != false ||
		payload["prepare_status"] != "skipped" ||
		payload["launch_status"] != "not-run" ||
		payload["profile_written"] != false ||
		payload["launcher_bundle_written"] != false ||
		payload["launch_attempted"] != false ||
		payload["runner_configured"] != true ||
		payload["runner_bottle_configured"] != true ||
		payload["runner_argument_count"] != float64(3) ||
		payload["runner_available"] != false ||
		payload["skip_bootstrap"] != true ||
		payload["wine_executed"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["network_checks_run"] != false ||
		payload["package_manager_invoked"] != false ||
		payload["raw_runner_path_exposed"] != false ||
		payload["skip_reason"] != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected known prepare-and-launch payload: %#v", payload)
	}
	if _, ok := payload["launch_payload"]; ok {
		t.Fatalf("missing artifact prepare-and-launch must not emit launch payload: %#v", payload)
	}
	assertCLIOutputOmitsValues(t, output.String(), "/private/runner", "private-bottle", "--private-runner-arg")
}

func TestWindowsKnownAppRunCommandLocalBackendSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-run",
		"--backend", "local",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--state-root", filepath.Join(tempDir, "state"),
		"--runner", "/private/runner",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_run.v1" ||
		payload["request_type"] != "windows-known-app-run" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["backend"] != "local" ||
		payload["backend_ready"] != false ||
		payload["launch_attempted"] != false ||
		payload["runner_available"] != false ||
		payload["checksum_verified"] != false ||
		payload["profile_written"] != false ||
		payload["launcher_bundle_written"] != false ||
		payload["loopback_only_networking"] != false ||
		payload["qemu_required"] != false ||
		payload["wine_executed"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["skip_reason"] != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected known app run local payload: %#v", payload)
	}
	if _, ok := payload["local_payload"]; !ok {
		t.Fatalf("local known app run must include local payload: %#v", payload)
	}
	if _, ok := payload["guest_payload"]; ok {
		t.Fatalf("local known app run must not include guest payload: %#v", payload)
	}
	assertCLIOutputOmitsValues(t, output.String(), "/private/runner")
}

func TestWindowsKnownAppRunCommandGuestWineBackendSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-run",
		"--backend", "guest-wine",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--timeout", "5s",
		"--redact-output",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_run.v1" ||
		payload["request_type"] != "windows-known-app-run" ||
		payload["status"] != "skipped" ||
		payload["app_id"] != "7zr" ||
		payload["backend"] != "guest-wine" ||
		payload["backend_ready"] != false ||
		payload["launch_attempted"] != false ||
		payload["runner_available"] != false ||
		payload["checksum_verified"] != false ||
		payload["loopback_only_networking"] != true ||
		payload["qemu_required"] != true ||
		payload["raw_output_redacted"] != true ||
		payload["wine_executed"] != false ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["skip_reason"] != "known Windows app artifact unavailable or checksum mismatch" {
		t.Fatalf("unexpected known app run guest payload: %#v", payload)
	}
	if _, ok := payload["local_payload"]; ok {
		t.Fatalf("guest known app run must not include local payload: %#v", payload)
	}
	if _, ok := payload["guest_payload"]; !ok {
		t.Fatalf("guest known app run must include guest payload: %#v", payload)
	}
	if strings.Contains(output.String(), tempDir) {
		t.Fatalf("known app run guest output leaked host paths: %s", output.String())
	}
}

func TestWindowsKnownAppRunCommandWritesReportOutput(t *testing.T) {
	tempDir := t.TempDir()
	reportPath := filepath.Join(tempDir, "evidence", "known-run.json")

	var output bytes.Buffer
	err := run([]string{
		"windows-known-app-run",
		"--backend", "guest-wine",
		"--app", "7zr",
		"--cache-root", tempDir,
		"--timeout", "5s",
		"--redact-output",
		"--report-output", reportPath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	reportBytes, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("ReadFile report returned error: %v", err)
	}
	if string(reportBytes) != output.String() {
		t.Fatalf("report output must match stdout:\nstdout=%s\nreport=%s", output.String(), string(reportBytes))
	}

	var payload map[string]any
	if err := json.Unmarshal(reportBytes, &payload); err != nil {
		t.Fatalf("Unmarshal report returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.known_windows_app_run.v1" ||
		payload["request_type"] != "windows-known-app-run" ||
		payload["status"] != "skipped" ||
		payload["backend"] != "guest-wine" ||
		payload["raw_output_redacted"] != true ||
		payload["skip_reason"] != "known Windows app artifact unavailable or checksum mismatch" {
		t.Fatalf("unexpected report payload: %#v", payload)
	}
	if strings.Contains(output.String(), reportPath) {
		t.Fatalf("known app run output leaked report path: %s", output.String())
	}
}

func TestWindowsKnownAppRunCommandRejectsUnsupportedBackend(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-run", "--backend", "missing-backend"}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["status"] != "failed" ||
		payload["backend"] != "missing-backend" ||
		payload["failure_reason"] != "unsupported known app run backend" ||
		payload["docker_executed"] != false ||
		payload["qemu_executed"] != false ||
		payload["wine_executed"] != false {
		t.Fatalf("unexpected unsupported backend payload: %#v", payload)
	}
}

func TestWindowsKnownAppManagedLaunchPreviewCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-managed-launch-preview", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppPrepareLaunchProfileCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-prepare-launch-profile", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppLaunchProfileMaterializeCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-launch-profile-materialize", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppPrepareAndLaunchProfileCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-prepare-and-launch-profile", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppRunCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-run", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppKDELauncherPreviewCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-kde-launcher-preview", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppLaunchRequestPreviewCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-launch-request-preview", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppDispatchPreviewCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-dispatch-preview", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppDispatchSmokeCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-dispatch-smoke", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func TestWindowsKnownAppLaunchBridgePreviewCommandRejectsUnknownApp(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-known-app-launch-bridge-preview", "--app", "missing-app"}, &output)
	if err == nil || !strings.Contains(err.Error(), "unknown known Windows app") {
		t.Fatalf("expected unknown app rejection, got %v", err)
	}
}

func anyStrings(values []any) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.(string))
	}
	return result
}

func writeCLIJSON(t *testing.T, path string, payload map[string]any) {
	t.Helper()
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal payload returned error: %v", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile JSON returned error: %v", err)
	}
}

func minimalPEFixture(machine uint16) []byte {
	data := make([]byte, 0x88)
	data[0] = 'M'
	data[1] = 'Z'
	binary.LittleEndian.PutUint32(data[0x3c:0x40], 0x80)
	copy(data[0x80:0x84], []byte{'P', 'E', 0, 0})
	binary.LittleEndian.PutUint16(data[0x84:0x86], machine)
	return data
}

func assertWindowsCompatibilityCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("Windows compatibility CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func assertKnownManagedLaunchCLISafe(t *testing.T, text string, hostPath string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"wine", "qemu", strings.ToLower(hostPath)} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("known managed launch CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}

func assertCLIOutputOmitsValues(t *testing.T, text string, values ...string) {
	t.Helper()
	for _, value := range values {
		if strings.TrimSpace(value) != "" && strings.Contains(text, value) {
			t.Fatalf("CLI output exposed forbidden value %q: %s", value, text)
		}
	}
}
