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

func TestRunSmokeDiscoversWine64OnPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("PATH-based shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	writeNamedFakeRunner(t, tempDir, "wine64", 0, DefaultMarker+"\n")
	t.Setenv("PATH", tempDir)

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		Timeout:        5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		!result.RunnerAvailable ||
		!result.MarkerObserved ||
		result.ExitCode != 0 {
		t.Fatalf("unexpected discovered runner result: %#v", result)
	}
}

func TestRunSmokeCanRedactRawOutputForDesktopConsumers(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, []byte("fixture"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\nraw-host-path=/private/tmp/secret\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
		Timeout:        5 * time.Second,
		RedactOutput:   true,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		!result.MarkerObserved ||
		result.RawOutputIncluded ||
		!result.RawOutputRedacted ||
		result.Stdout != "" ||
		result.Stderr != "" ||
		result.StdoutBytes == 0 ||
		result.StdoutLineCount != 2 ||
		!strings.Contains(result.KDESafeOutputSummary, "expected smoke marker observed") {
		t.Fatalf("unexpected redacted result: %#v", result)
	}
	if strings.Contains(result.KDESafeOutputSummary, "/private/tmp") ||
		strings.Contains(result.KDESafeOutputSummary, "secret") {
		t.Fatalf("KDE-safe output summary leaked raw runner output: %q", result.KDESafeOutputSummary)
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
	return writeNamedFakeRunner(t, tempDir, "fake-runner", exitCode, stdout)
}

func writeNamedFakeRunner(t *testing.T, tempDir string, name string, exitCode int, stdout string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	path := filepath.Join(tempDir, name)
	body := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"printf '%s' '" + strings.ReplaceAll(stdout, "'", "'\\''") + "'\n" +
		"exit " + string(rune('0'+exitCode)) + "\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	return path
}
