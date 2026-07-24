package winapp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreflightSmokeProfileReportsReadyWithoutExecutingRunner(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "real-app.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	workingDir := filepath.Join(tempDir, "app-dir")
	if err := os.Mkdir(workingDir, 0o700); err != nil {
		t.Fatalf("Mkdir working dir returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "wine")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nexit 99\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":    SmokeProfileSchemaVersion,
		"executable_path":   executablePath,
		"working_directory": workingDir,
		"runner_path":       runnerPath,
		"runner_arguments":  []string{"--private-shim"},
		"arguments":         []string{"--open"},
		"state_root":        filepath.Join(tempDir, "state"),
		"timeout":           "15s",
		"expected_marker":   "APP_OK",
		"success_mode":      SuccessModeMarker,
	})

	result, err := PreflightSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("PreflightSmokeProfile returned error: %v", err)
	}
	if result.Status != ProfileReadyStatus ||
		result.SchemaVersion != SmokeProfilePreflightSchemaVersion ||
		result.RequestType != SmokeProfilePreflightRequestType ||
		!result.ProfileSupplied ||
		result.ExecutableName != "real-app.exe" ||
		!result.ExecutableExists ||
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "x86_64" ||
		!result.ExecutableArchitectureReady ||
		result.WorkingDirectoryMode != WorkingDirectoryModeOperator ||
		!result.WorkingDirectoryValid ||
		!result.StateRootConfigured ||
		!result.RunnerAvailable ||
		result.RunnerArgumentCount != 1 ||
		result.AppArgumentCount != 1 ||
		result.SuccessMode != SuccessModeMarker ||
		!result.TimeoutConfigured ||
		!result.ExpectedMarkerConfigured ||
		result.WineExecuted ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.HostRootModified {
		t.Fatalf("unexpected ready preflight result: %#v", result)
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal result returned error: %v", err)
	}
	output := string(encoded)
	for _, leaked := range []string{profilePath, executablePath, workingDir, runnerPath, "--private-shim"} {
		if strings.Contains(output, leaked) {
			t.Fatalf("preflight output leaked private value %q: %s", leaked, output)
		}
	}
}

func TestPreflightSmokeProfileBlocksNonWindowsExecutable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "not-windows.exe")
	if err := os.WriteFile(executablePath, []byte("plain text"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "wine")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nexit 0\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": executablePath,
		"runner_path":     runnerPath,
		"state_root":      filepath.Join(tempDir, "state"),
	})

	result, err := PreflightSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("PreflightSmokeProfile returned error: %v", err)
	}
	if result.Status != ProfileBlockedStatus ||
		result.ExecutableFormat != "unknown" ||
		result.WindowsExecutableSignature ||
		result.FailureReason != "executable is not a Windows PE file" ||
		result.WineExecuted ||
		result.HostRootModified {
		t.Fatalf("unexpected non-Windows executable preflight result: %#v", result)
	}
}

func TestPreflightSmokeProfileBlocksUnsupportedArchitecture(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "arm-app.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0xaa64), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "wine")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nexit 0\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": executablePath,
		"runner_path":     runnerPath,
		"state_root":      filepath.Join(tempDir, "state"),
	})

	result, err := PreflightSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("PreflightSmokeProfile returned error: %v", err)
	}
	if result.Status != ProfileBlockedStatus ||
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "arm64" ||
		result.ExecutableArchitectureReady ||
		result.FailureReason != "Windows executable architecture is not supported" ||
		result.RunnerAvailable ||
		result.WineExecuted ||
		result.HostRootModified {
		t.Fatalf("unexpected unsupported architecture preflight result: %#v", result)
	}
}

func TestPreflightSmokeProfileBlocksMissingExecutable(t *testing.T) {
	tempDir := t.TempDir()
	runnerPath := filepath.Join(tempDir, "wine")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nexit 0\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	profilePath := filepath.Join(tempDir, "app.profile.json")
	writeSmokeProfile(t, profilePath, map[string]any{
		"schema_version":  SmokeProfileSchemaVersion,
		"executable_path": filepath.Join(tempDir, "missing.exe"),
		"runner_path":     runnerPath,
		"state_root":      filepath.Join(tempDir, "state"),
	})

	result, err := PreflightSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("PreflightSmokeProfile returned error: %v", err)
	}
	if result.Status != ProfileBlockedStatus ||
		result.ExecutableExists ||
		result.FailureReason == "" ||
		result.NextAction == "" ||
		result.WineExecuted ||
		result.HostRootModified {
		t.Fatalf("unexpected blocked preflight result: %#v", result)
	}
}

func writeSmokeProfile(t *testing.T, path string, payload map[string]any) {
	t.Helper()
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal profile returned error: %v", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile profile returned error: %v", err)
	}
}
