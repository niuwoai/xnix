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
		result.PullPolicy != "never" ||
		result.NetworkMode != "none" ||
		result.CompatibilityLayer != "containerized-windows-compatibility-layer" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if !result.IsolatedStateRoot ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.HostMountCount != 2 {
		t.Fatalf("unexpected safety flags: %#v", result)
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	for _, token := range []string{
		"image inspect local/wine-smoke:test",
		"run --rm --pull never --network none",
		"--security-opt no-new-privileges",
		"--cap-drop ALL",
		"--env WINEPREFIX=/state/wineprefix",
		"--env HOME=/state/home",
		"local/wine-smoke:test wine /work/hello.exe",
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

func writeFakeDocker(t *testing.T, tempDir string, logPath string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-docker")
	body := "#!/bin/sh\n" +
		"printf '%s\\n' \"$*\" >> '" + logPath + "'\n" +
		"if test \"$1 $2\" = 'image inspect'; then exit 0; fi\n" +
		"case \"$*\" in\n" +
		"  *'run --rm --pull never --network none'*'--security-opt no-new-privileges'*'--cap-drop ALL'*'--env WINEPREFIX=/state/wineprefix'*'local/wine-smoke:test wine /work/hello.exe'*) printf 'XNIX_WINAPP_SMOKE_OK\\n'; exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile docker returned error: %v", err)
	}
	return path
}
