package appidentity

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunExternalWinAppConsumesImportRecordAndRunsContainerGUI(t *testing.T) {
	if testing.Short() {
		t.Skip("external app run fake Docker fixture uses shell")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath))
	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf '0x700001 \"External GUI\": (\"ExternalGui.exe\" \"ExternalGui.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}

	result, err := RunExternalWinApp(context.Background(), ExternalWinAppRunRequest{
		ImportRecordPath: recordPath,
		WindowMatch:      "External GUI",
		Image:            "local/wine-x-gui:test",
		Platform:         "linux/amd64",
		DockerPath:       dockerPath,
		Timeout:          5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunExternalWinApp returned error: %v", err)
	}
	if result.SchemaVersion != ExternalWinAppRunSchemaVersion ||
		result.RequestType != ExternalWinAppRunRequestType ||
		result.Status != "passed" ||
		result.ApplicationID != "org.xnix.external.gui" ||
		result.DisplayName != "External GUI" ||
		result.AppVersion != "0.2.640-test" ||
		result.ExecutableName != "ExternalGui.exe" ||
		!result.ExternalAppImportRecordConsumed ||
		!result.ImportedArtifactDigestVerified ||
		result.ImportedArtifactSHA256 != record.ArtifactSHA256 ||
		!result.RuntimeRunRequested ||
		!result.RuntimeRunExecuted ||
		!result.ExecutionStarted ||
		!result.BackendProcessStarted ||
		!result.ContainerRuntimeUsed ||
		result.ContainerNetworkMode != "none" ||
		result.ContainerHostMountCount != 0 ||
		!result.XWindowObserved ||
		!result.WindowObserved ||
		result.DesktopLaunchEnabled ||
		result.ActionExecutionEnabled ||
		result.BackendDetailsExposed ||
		result.RawImportRecordPathExposed ||
		result.RawStateRootPathExposed ||
		result.RawExecutablePathExposed ||
		result.HostRootModified ||
		result.DockerSocketMounted ||
		result.RuntimePayload.ExternalAppImportRecordConsumed != true ||
		result.RuntimePayload.ImportedArtifactSHA256 != record.ArtifactSHA256 {
		t.Fatalf("unexpected external app run result: %#v", result)
	}
	payload, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(payload), recordPath) ||
		strings.Contains(string(payload), stateRoot) ||
		strings.Contains(string(payload), executablePath) ||
		strings.Contains(string(payload), dockerPath) ||
		strings.Contains(string(payload), "docker.sock") ||
		strings.Contains(string(payload), "--privileged") ||
		strings.Contains(string(payload), "--network host") {
		t.Fatalf("external app run result exposed unsafe details: %s", string(payload))
	}
	if !strings.Contains(string(payload), `"window_observed":true`) ||
		!strings.Contains(string(payload), `"x_window_observed":true`) {
		t.Fatalf("external app run result must expose both generic and X-specific window evidence: %s", string(payload))
	}
	dockerInvocation, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("ReadFile docker log returned error: %v", err)
	}
	if !strings.Contains(string(dockerInvocation), "cp") ||
		!strings.Contains(string(dockerInvocation), "ExternalGui.exe") ||
		strings.Contains(string(dockerInvocation), "--network host") ||
		strings.Contains(string(dockerInvocation), "--privileged") ||
		strings.Contains(string(dockerInvocation), "/var/run/docker.sock") {
		t.Fatalf("unexpected fake Docker invocation: %s", string(dockerInvocation))
	}
}

func TestRunExternalWinAppConsumesHandleAndRunsContainerGUI(t *testing.T) {
	if testing.Short() {
		t.Skip("external app run fake Docker fixture uses shell")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalGui.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.gui",
		DisplayName:    "External GUI",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	dockerLog := filepath.Join(tempDir, "fake-docker.log")
	dockerPath := filepath.Join(tempDir, "fake-docker")
	dockerBody := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> \"" + dockerLog + "\"\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'create'; then printf 'fake-x-gui-container\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'cp'; then exit 0; fi\n" +
		"if test \"$1 $2\" = 'start -a'; then " +
		"printf 'XNIX_X_GUI_XSERVER_STARTED=true\\n'\n" +
		"printf 'XNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n'\n" +
		"printf '0x700001 \"External GUI\": (\"ExternalGui.exe\" \"ExternalGui.exe\") 320x160+20+20 +20+20\\n'\n" +
		"printf 'XNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"if test \"$1\" = 'rm'; then exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(dockerPath, []byte(dockerBody), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}

	result, err := RunExternalWinApp(context.Background(), ExternalWinAppRunRequest{
		StateRoot:         stateRoot,
		ExternalAppHandle: "org.xnix.external.gui",
		WindowMatch:       "External GUI",
		Image:             "local/wine-x-gui:test",
		Platform:          "linux/amd64",
		DockerPath:        dockerPath,
		Timeout:           5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunExternalWinApp returned error: %v", err)
	}
	if result.Status != "passed" ||
		result.ApplicationID != "org.xnix.external.gui" ||
		result.ExternalAppImportRecordConsumed != true ||
		result.ExternalAppHandleConsumed != true ||
		result.ExternalAppHandle != "org.xnix.external.gui" ||
		result.ImportedArtifactDigestVerified != true ||
		result.ImportedArtifactSHA256 != record.ArtifactSHA256 ||
		result.RawExternalAppHandlePathExposed != false ||
		result.RawImportRecordPathExposed != false ||
		result.RawStateRootPathExposed != false ||
		result.RawExecutablePathExposed != false ||
		result.ContainerNetworkMode != "none" ||
		result.ContainerHostMountCount != 0 ||
		result.XWindowObserved != true ||
		result.WindowObserved != true ||
		result.RuntimePayload.ExternalAppImportRecordConsumed != true {
		t.Fatalf("unexpected external app handle run result: %#v", result)
	}
	payload, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(payload), stateRoot) ||
		strings.Contains(string(payload), executablePath) ||
		strings.Contains(string(payload), dockerPath) ||
		strings.Contains(string(payload), "docker.sock") ||
		strings.Contains(string(payload), "--privileged") ||
		strings.Contains(string(payload), "--network host") {
		t.Fatalf("external app handle run result exposed unsafe details: %s", string(payload))
	}
	if _, err := RunExternalWinApp(context.Background(), ExternalWinAppRunRequest{
		ImportRecordPath:  filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath)),
		StateRoot:         stateRoot,
		ExternalAppHandle: "org.xnix.external.gui",
		Image:             "local/wine-x-gui:test",
		Platform:          "linux/amd64",
		DockerPath:        dockerPath,
		Timeout:           5 * time.Second,
	}); err == nil || !strings.Contains(err.Error(), "either --external-app-import-record or --state-root") {
		t.Fatalf("expected mutually exclusive handle/import path rejection, got %v", err)
	}
}
