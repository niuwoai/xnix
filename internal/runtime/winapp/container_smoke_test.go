package winapp

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestRunContainerSmokeUsesRestrictedDockerRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell docker fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	logPath := filepath.Join(tempDir, "docker.log")
	dockerPath := writeFakeDocker(t, tempDir, logPath)

	result, err := RunContainerSmoke(context.Background(), ContainerRequest{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		Image:          "local/wine-smoke:test",
		DockerPath:     dockerPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunContainerSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		!result.RunnerAvailable ||
		!result.ImageAvailable ||
		!result.MarkerObserved ||
		result.ExitCode != 0 ||
		result.ExecutableName != "hello.exe" ||
		result.ContainerImage != "local/wine-smoke:test" ||
		result.ContainerPlatform != DefaultWinePlatform ||
		result.ContainerStateMode != "tmpfs" ||
		result.PullPolicy != "never" ||
		result.NetworkMode != "none" ||
		!result.WineBootstrapRequired ||
		result.WineBootstrapTimedOut ||
		result.WineBootstrapExitCode != -1 ||
		result.CompatibilityLayer != "containerized-windows-compatibility-layer" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if !result.IsolatedStateRoot ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.HostMountCount != 1 {
		t.Fatalf("unexpected safety flags: %#v", result)
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	for _, token := range []string{
		"image inspect local/wine-smoke:test",
		"run --rm --platform linux/amd64 --pull never --network none",
		"--cpus 2",
		"--memory 2g",
		"--security-opt no-new-privileges",
		"--cap-drop ALL",
		"--tmpfs /state:rw,nosuid,nodev,size=768m",
		"--env WINEPREFIX=/state/wineprefix",
		"--env WINEARCH=win64",
		"--env WINEDEBUG=-all",
		"--env WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d",
		"--env XNIX_WINE_BOOTSTRAP_TIMEOUT_SECONDS=300",
		"--env HOME=/state/home",
		"local/wine-smoke:test sh -lc",
		"wineboot --init",
		"exec wine \"$@\" xnix-wine-smoke /work/hello.exe",
	} {
		if !strings.Contains(log, token) {
			t.Fatalf("fake docker log missing %q: %s", token, log)
		}
	}
}

func TestRunContainerSmokeSkipsWhenImageUnavailable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell docker fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	dockerPath := filepath.Join(tempDir, "fake-docker")
	body := "#!/bin/sh\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 1; fi\n" +
		"exit 1\n"
	if err := os.WriteFile(dockerPath, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}

	result, err := RunContainerSmoke(context.Background(), ContainerRequest{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		Image:          "missing/wine-smoke:test",
		DockerPath:     dockerPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunContainerSmoke returned error: %v", err)
	}
	if result.Status != SkippedStatus ||
		!result.RunnerAvailable ||
		result.ImageAvailable ||
		!strings.Contains(result.SkipReason, "image unavailable") {
		t.Fatalf("unexpected skip result: %#v", result)
	}
}

func TestRunContainerSmokeReportsWineBootstrapTimeout(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell docker fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	dockerPath := filepath.Join(tempDir, "fake-docker")
	body := "#!/bin/sh\n" +
		"if test \"$1 $2\" = 'image inspect'; then printf 'linux/amd64\\n'; exit 0; fi\n" +
		"printf 'XNIX_WINE_BOOTSTRAP_EXIT:124\\n' >&2\n" +
		"exit 124\n"
	if err := os.WriteFile(dockerPath, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}

	result, err := RunContainerSmoke(context.Background(), ContainerRequest{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		Image:          "local/wine-smoke:test",
		DockerPath:     dockerPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunContainerSmoke returned error: %v", err)
	}
	if result.Status != FailedStatus ||
		result.FailureReason != "wine bootstrap timed out" ||
		!result.WineBootstrapTimedOut ||
		result.WineBootstrapExitCode != 124 ||
		result.MarkerObserved ||
		result.ExitCode != 124 {
		t.Fatalf("unexpected timeout result: %#v", result)
	}
}

func TestRunContainerXGUISmokeObservesWindowWithRestrictedDockerRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell docker fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, "docker-x-gui.log")
	dockerPath := writeFakeXGUIDocker(t, tempDir, logPath)

	result, err := RunContainerXGUISmoke(context.Background(), ContainerXGUIRequest{
		ApplicationName: "notepad.exe",
		WindowMatch:     "notepad.exe",
		Image:           "local/wine-x-gui:test",
		Platform:        "linux/amd64",
		DockerPath:      dockerPath,
		Timeout:         5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunContainerXGUISmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.SchemaVersion != ContainerXGUISchemaVersion ||
		result.RequestType != ContainerXGUIRequestType ||
		result.ApplicationName != "notepad.exe" ||
		result.WindowMatch != "notepad.exe" ||
		result.ContainerImage != "local/wine-x-gui:test" ||
		result.ContainerPlatform != "linux/amd64" ||
		result.PullPolicy != "never" ||
		result.NetworkMode != "none" ||
		result.DesktopDisplay != "Xvfb" ||
		!result.RunnerAvailable ||
		!result.ImageAvailable ||
		!result.XServerStarted ||
		!result.WineBootstrapAttempted ||
		!result.XWindowObserved ||
		result.WindowEvidenceSummary == "" ||
		result.ExitCode != 0 {
		t.Fatalf("unexpected X GUI result: %#v", result)
	}
	if result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.HostMountCount != 0 {
		t.Fatalf("unexpected X GUI safety flags: %#v", result)
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	for _, token := range []string{
		"image inspect local/wine-x-gui:test",
		"run --rm --pull never --network none",
		"--cpus 2",
		"--memory 3g",
		"--pids-limit 512",
		"--security-opt no-new-privileges",
		"--cap-drop ALL",
		"--tmpfs /state:rw,nosuid,nodev,size=1g",
		"--env WINEPREFIX=/state/wineprefix",
		"--env DISPLAY=:99",
		"--env XNIX_GUI_APP=notepad.exe",
		"--env XNIX_WINDOW_MATCH=notepad.exe",
		"--platform linux/amd64",
		"local/wine-x-gui:test sh -lc",
		"Xvfb \"$DISPLAY\"",
		"xwininfo -root -tree",
		"wine \"$XNIX_GUI_APP\"",
	} {
		if !strings.Contains(log, token) {
			t.Fatalf("fake X GUI docker log missing %q: %s", token, log)
		}
	}
	if strings.Contains(log, "docker.sock") || strings.Contains(log, "--privileged") || strings.Contains(log, "--network host") || strings.Contains(log, "--volume") {
		t.Fatalf("X GUI docker log contains unsafe host access: %s", log)
	}
}

func writeFakeDocker(t *testing.T, tempDir string, logPath string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-docker")
	body := "#!/bin/sh\n" +
		"printf '%s\\n' \"$*\" >> '" + logPath + "'\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 0; fi\n" +
		"case \"$*\" in\n" +
		"  *'run --rm --platform linux/amd64 --pull never --network none'*'--cpus 2'*'--memory 2g'*'--security-opt no-new-privileges'*'--cap-drop ALL'*'--tmpfs /state:rw,nosuid,nodev,size=768m'*'--env WINEPREFIX=/state/wineprefix'*'--env WINEARCH=win64'*'local/wine-smoke:test sh -lc'*'wineboot --init'*'exec wine \"$@\"'*'xnix-wine-smoke /work/hello.exe'*) printf 'XNIX_WINAPP_SMOKE_OK\\n'; exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	return path
}

func writeFakeXGUIDocker(t *testing.T, tempDir string, logPath string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-x-gui-docker")
	body := "#!/bin/sh\n" +
		"printf '%s\\n' \"$*\" >> '" + logPath + "'\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 0; fi\n" +
		"if test \"$1\" = 'run'; then printf 'XNIX_X_GUI_XSERVER_STARTED=true\\nXNIX_X_GUI_WINE_BOOTSTRAP_ATTEMPTED=true\\n0x600001 \"Untitled - Notepad\": (\"notepad.exe\" \"notepad.exe\") 721x519+4+23 +4+23\\nXNIX_X_GUI_WINDOW_OBSERVED=true\\n'; exit 0; fi\n" +
		"exit 2\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	return path
}
