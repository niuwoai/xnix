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

func TestRunSmokeUsesIsolatedStateRootAndObservesMarker(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		!result.RunnerAvailable ||
		!result.MarkerObserved ||
		result.ExitCode != 0 ||
		result.ExecutableName != "hello.exe" ||
		result.CompatibilityLayer != "windows-compatibility-layer" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if !result.IsolatedStateRoot ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired {
		t.Fatalf("unexpected safety flags: %#v", result)
	}
}

func TestRunSmokeSkipsWhenRunnerUnavailable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     filepath.Join(tempDir, "missing-runner"),
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != SkippedStatus ||
		result.RunnerAvailable ||
		!strings.Contains(result.SkipReason, "runner unavailable") {
		t.Fatalf("unexpected skip result: %#v", result)
	}
}

func TestRunSmokeRejectsNonWindowsExecutable(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "hello")
	if err := os.WriteFile(path, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	_, err := RunSmoke(context.Background(), Request{
		ExecutablePath: path,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     filepath.Join(tempDir, "missing-runner"),
	})
	if err == nil || !strings.Contains(err.Error(), "Windows .exe") {
		t.Fatalf("expected Windows executable validation error, got %v", err)
	}
}

func writeFakeRunner(t *testing.T, tempDir string, exitCode int, stdout string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	path := filepath.Join(tempDir, "fake-runner")
	body := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf '%s' '" + strings.ReplaceAll(stdout, "'", "'\\''") + "'\n" +
		"exit " + string(rune('0'+exitCode)) + "\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	return path
}
