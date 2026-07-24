package winapp

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLaunchProfileBlocksBeforeExecutionWhenRunnerUnavailable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": executablePath,
		"state_root":      stateRoot,
		"runner_path":     filepath.Join(tempDir, "missing-runner"),
		"timeout":         "5s",
		"stage_app_dir":   true,
	})

	result, err := LaunchProfile(context.Background(), LaunchProfileRequest{ProfilePath: profilePath})
	if err != nil {
		t.Fatalf("LaunchProfile returned error: %v", err)
	}
	if result.SchemaVersion != LaunchProfileSchemaVersion ||
		result.RequestType != LaunchProfileRequestType ||
		result.Status != ProfileBlockedStatus ||
		result.PreflightStatus != ProfileBlockedStatus ||
		result.LaunchAttempted ||
		result.RuntimePayload != nil ||
		result.ExecutableName != "hello.exe" ||
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "x86_64" ||
		!result.ExecutableArchitectureReady ||
		result.WineArchitecture != "win64" ||
		result.ApplicationWorkspaceMode != ApplicationWorkspaceModeStaged ||
		result.RunnerAvailable ||
		result.SkipReason != "windows compatibility runner unavailable" ||
		!result.RawOutputRedacted ||
		result.RawProfilePathExposed ||
		result.RawExecutablePathExposed ||
		result.RawRunnerPathExposed ||
		result.HostRootModified ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.WineExecuted {
		t.Fatalf("unexpected blocked launch profile result: %#v", result)
	}
	if _, err := os.Stat(stateRoot); err == nil {
		t.Fatalf("blocked launch profile must not create state root before a runner is ready")
	}
}

func TestLaunchProfileRunsReadyProfileWithRedactedOutput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"printf 'XNIX_WINAPP_SMOKE_OK\\nraw-host-path=/private/tmp/secret\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "real-app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": executablePath,
		"state_root":      filepath.Join(tempDir, "state"),
		"runner_path":     runnerPath,
		"timeout":         "5s",
	})

	result, err := LaunchProfile(context.Background(), LaunchProfileRequest{ProfilePath: profilePath})
	if err != nil {
		t.Fatalf("LaunchProfile returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.PreflightStatus != ProfileReadyStatus ||
		!result.LaunchAttempted ||
		result.RuntimePayload == nil ||
		!result.RunnerAvailable ||
		!result.WineExecuted ||
		!result.RawOutputRedacted ||
		result.RuntimePayload.Stdout != "" ||
		result.RuntimePayload.Stderr != "" ||
		!result.RuntimePayload.RawOutputRedacted ||
		!strings.Contains(result.RuntimePayload.KDESafeOutputSummary, "expected smoke marker observed") ||
		result.HostRootModified ||
		result.DockerExecuted ||
		result.QEMUExecuted {
		t.Fatalf("unexpected ready launch profile result: %#v", result)
	}
}
