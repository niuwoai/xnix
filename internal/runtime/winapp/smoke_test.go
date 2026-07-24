package winapp

import (
	"context"
	"encoding/binary"
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
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
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
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "x86_64" ||
		!result.ExecutableArchitectureReady ||
		result.WineArchitecture != "win64" ||
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

func TestRunSmokeUsesManagedWineEnvironment(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
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
	if result.Status != PassedStatus {
		t.Fatalf("expected managed Wine environment smoke to pass, got %#v", result)
	}
}

func TestRunSmokeUsesWin32WineArchitectureForX86Executable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "legacy.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x014c), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunnerExpectingWineArchitecture(t, tempDir, "win32", 0, DefaultMarker+"\n")

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
		result.ExecutableArchitecture != "x86" ||
		!result.ExecutableArchitectureReady ||
		result.WineArchitecture != "win32" ||
		!result.MarkerObserved {
		t.Fatalf("unexpected win32 architecture smoke result: %#v", result)
	}
}

func TestRunSmokeDefaultsWorkingDirectoryToExecutableDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "companion.dll"), []byte("sidecar"), 0o600); err != nil {
		t.Fatalf("WriteFile sidecar returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\n" +
		"test -f companion.dll || exit 74\n" +
		"printf '" + DefaultMarker + "\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

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
		result.WorkingDirectoryMode != WorkingDirectoryModeExecutable ||
		!result.MarkerObserved {
		t.Fatalf("unexpected executable-directory working dir result: %#v", result)
	}
}

func TestRunSmokeUsesOperatorWorkingDirectoryWithoutReportingPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
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
		"printf '" + DefaultMarker + "\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath:   executablePath,
		StateRoot:        filepath.Join(tempDir, "state"),
		WorkingDirectory: workingDir,
		RunnerPath:       runnerPath,
		Timeout:          5 * time.Second,
		RedactOutput:     true,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.WorkingDirectoryMode != WorkingDirectoryModeOperator ||
		!result.MarkerObserved ||
		strings.Contains(result.KDESafeOutputSummary, workingDir) {
		t.Fatalf("unexpected operator working dir result: %#v", result)
	}
}

func TestRunSmokePassesRunnerArgumentsBeforeExecutable(t *testing.T) {
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
		"test \"$1\" = --bottle || exit 78\n" +
		"test \"$2\" = smoke-bottle || exit 77\n" +
		"test \"$3\" = '" + strings.ReplaceAll(executablePath, "'", "'\\''") + "' || exit 76\n" +
		"test \"$4\" = --app-flag || exit 75\n" +
		"printf '" + DefaultMarker + "\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath:  executablePath,
		Arguments:       []string{"--app-flag"},
		RunnerArguments: []string{"--bottle", "smoke-bottle"},
		StateRoot:       filepath.Join(tempDir, "state"),
		RunnerPath:      runnerPath,
		Timeout:         5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.RunnerArgumentCount != 2 ||
		!result.MarkerObserved {
		t.Fatalf("unexpected runner argument smoke result: %#v", result)
	}
}

func TestRunSmokeExpandsRunnerBottleBeforeRunnerArguments(t *testing.T) {
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
		"test \"$1\" = --bottle || exit 78\n" +
		"test \"$2\" = smoke-bottle || exit 77\n" +
		"test \"$3\" = --shim-mode || exit 76\n" +
		"test \"$4\" = '" + strings.ReplaceAll(executablePath, "'", "'\\''") + "' || exit 75\n" +
		"printf '" + DefaultMarker + "\\n'\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath:  executablePath,
		RunnerBottle:    "smoke-bottle",
		RunnerArguments: []string{"--shim-mode"},
		StateRoot:       filepath.Join(tempDir, "state"),
		RunnerPath:      runnerPath,
		Timeout:         5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.RunnerArgumentCount != 3 ||
		!result.MarkerObserved {
		t.Fatalf("unexpected runner bottle smoke result: %#v", result)
	}
}

func TestRunSmokeCanPassOnExitCodeWithoutMarker(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, "GUI app exited cleanly\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
		Timeout:        5 * time.Second,
		SuccessMode:    SuccessModeExitCode,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.SuccessMode != SuccessModeExitCode ||
		result.MarkerObserved ||
		result.FailureReason != "" {
		t.Fatalf("unexpected exit-code success result: %#v", result)
	}
}

func TestRunSmokeCanPassWhenProcessSurvivesStartupWindow(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(tempDir, "fake-runner")
	runnerBody := "#!/bin/sh\nsleep 1\n"
	if err := os.WriteFile(runnerPath, []byte(runnerBody), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
		Timeout:        50 * time.Millisecond,
		SuccessMode:    SuccessModeStartupWindow,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.SuccessMode != SuccessModeStartupWindow ||
		!result.StartupWindowObserved ||
		result.MarkerObserved ||
		result.FailureReason != "" {
		t.Fatalf("unexpected startup-window result: %#v", result)
	}
}

func TestRunSmokeFailsStartupWindowWhenProcessExitsEarly(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, "closed quickly\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: executablePath,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
		Timeout:        5 * time.Second,
		SuccessMode:    SuccessModeStartupWindow,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != FailedStatus ||
		result.StartupWindowObserved ||
		result.FailureReason != "process exited before startup window elapsed" {
		t.Fatalf("unexpected early-exit startup-window result: %#v", result)
	}
}

func TestRunSmokeBootstrapsWinePrefixWhenWinebootIsAvailable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeNamedFakeRunner(t, tempDir, "wine", "win64", 0, DefaultMarker+"\n")
	bootstrapMarker := filepath.Join(tempDir, "bootstrap.marker")
	winebootBody := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"test \"$WINEARCH\" = win64 || exit 88\n" +
		"printf bootstrapped > '" + strings.ReplaceAll(bootstrapMarker, "'", "'\\''") + "'\n"
	if err := os.WriteFile(filepath.Join(tempDir, "wineboot"), []byte(winebootBody), 0o700); err != nil {
		t.Fatalf("WriteFile wineboot returned error: %v", err)
	}

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
		!result.WineBootstrapAttempted ||
		!result.WineBootstrapSucceeded ||
		result.WineBootstrapExitCode != 0 ||
		!result.MarkerObserved {
		t.Fatalf("unexpected bootstrap smoke result: %#v", result)
	}
	if _, err := os.Stat(bootstrapMarker); err != nil {
		t.Fatalf("expected wineboot marker to exist: %v", err)
	}
}

func TestRunSmokePassesRunnerArgumentsToWineboot(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeNamedFakeRunner(t, tempDir, "wine", "win64", 0, DefaultMarker+"\n")
	winebootBody := "#!/bin/sh\n" +
		"test \"$1\" = --bottle || exit 78\n" +
		"test \"$2\" = smoke-bottle || exit 77\n" +
		"test \"$3\" = --init || exit 76\n"
	if err := os.WriteFile(filepath.Join(tempDir, "wineboot"), []byte(winebootBody), 0o700); err != nil {
		t.Fatalf("WriteFile wineboot returned error: %v", err)
	}

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath:  executablePath,
		RunnerArguments: []string{"--bottle", "smoke-bottle"},
		StateRoot:       filepath.Join(tempDir, "state"),
		RunnerPath:      runnerPath,
		Timeout:         5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.RunnerArgumentCount != 2 ||
		!result.WineBootstrapSucceeded ||
		!result.MarkerObserved {
		t.Fatalf("unexpected runner-argument bootstrap result: %#v", result)
	}
}

func TestRunSmokeSkipsWhenRunnerUnavailable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
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

func TestRunnerDiagnosticsReportsExplicitRunnerWithoutRawPath(t *testing.T) {
	tempDir := t.TempDir()
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")

	result := RunnerDiagnostics(runnerPath)
	if result.Status != PassedStatus ||
		!result.RunnerAvailable ||
		!result.ExplicitRunnerSupplied ||
		result.CandidateCount != 1 ||
		result.SelectedRunnerName != "fake-runner" ||
		len(result.RunnerCommandHints) != 3 ||
		result.RawPathExposed ||
		result.HostRootModified ||
		result.PackageManagerInvoked ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.ColimaExecuted {
		t.Fatalf("unexpected explicit runner diagnostics: %#v", result)
	}
	if len(result.Candidates) != 1 ||
		result.Candidates[0].ID != "explicit-runner" ||
		result.Candidates[0].Source != "operator-supplied-runner" ||
		!result.Candidates[0].Available ||
		!result.Candidates[0].Selected ||
		result.Candidates[0].Reason != "available" {
		t.Fatalf("unexpected candidate evidence: %#v", result.Candidates)
	}
	if !strings.Contains(strings.Join(result.RunnerCommandHints, "\n"), "ruby scripts/winapp_smoke.rb --exe path/to/app.exe --format json") ||
		!strings.Contains(strings.Join(result.RunnerCommandHints, "\n"), "--runner-bottle bottle-name") ||
		strings.Contains(strings.Join(result.RunnerCommandHints, "\n"), runnerPath) {
		t.Fatalf("unexpected available runner command hints: %#v", result.RunnerCommandHints)
	}
}

func TestRunSmokeUsesConfiguredRunnerEnvironmentVariable(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")
	t.Setenv(RunnerEnvVar, runnerPath)

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
		!result.MarkerObserved {
		t.Fatalf("unexpected env runner smoke result: %#v", result)
	}
}

func TestRunnerDiagnosticsReportsEnvRunnerWithoutRawPath(t *testing.T) {
	tempDir := t.TempDir()
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")
	t.Setenv(RunnerEnvVar, runnerPath)

	result := RunnerDiagnostics("")
	if result.Status != PassedStatus ||
		!result.RunnerAvailable ||
		result.ExplicitRunnerSupplied ||
		!result.EnvRunnerConfigured ||
		result.CandidateCount != 1 ||
		result.SelectedRunnerName != "fake-runner" ||
		result.RawPathExposed {
		t.Fatalf("unexpected env runner diagnostics: %#v", result)
	}
	if len(result.Candidates) != 1 ||
		result.Candidates[0].ID != "env-runner" ||
		result.Candidates[0].Source != "env-configured-runner" ||
		!result.Candidates[0].Available ||
		!result.Candidates[0].Selected {
		t.Fatalf("unexpected env candidate evidence: %#v", result.Candidates)
	}
}

func TestRunnerDiagnosticsReportsUnavailableWithoutRawPath(t *testing.T) {
	tempDir := t.TempDir()
	missingRunner := filepath.Join(tempDir, "missing-runner")

	result := RunnerDiagnostics(missingRunner)
	if result.Status != SkippedStatus ||
		result.RunnerAvailable ||
		!result.ExplicitRunnerSupplied ||
		result.CandidateCount != 1 ||
		result.SelectedRunnerName != "" ||
		len(result.RunnerCommandHints) != 4 ||
		result.RawPathExposed ||
		!strings.Contains(result.NextAction, "--runner PATH") {
		t.Fatalf("unexpected missing runner diagnostics: %#v", result)
	}
	hints := strings.Join(result.RunnerCommandHints, "\n")
	if !strings.Contains(hints, "XNIX_WINDOWS_RUNNER=path/to/wine") ||
		!strings.Contains(hints, "--runner path/to/wine") ||
		!strings.Contains(hints, "--runner-bottle bottle-name") ||
		strings.Contains(hints, missingRunner) {
		t.Fatalf("unexpected missing runner command hints: %#v", result.RunnerCommandHints)
	}
	if len(result.Candidates) != 1 ||
		result.Candidates[0].Available ||
		result.Candidates[0].Reason != "not-found" {
		t.Fatalf("unexpected missing candidate evidence: %#v", result.Candidates)
	}
}

func TestRunnerCandidatesIncludeCommonDarwinCompatibilityRunners(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS app bundle runner candidates are only available on Darwin hosts")
	}

	candidates := runnerCandidates()
	ids := make(map[string]bool, len(candidates))
	for _, candidate := range candidates {
		ids[candidate.id] = true
	}

	for _, id := range []string{
		"crossover-system-wine",
		"crossover-system-wine64",
		"crossover-user-wine",
		"crossover-user-wine64",
		"whisky-user-wine",
		"whisky-user-wine64",
	} {
		if !ids[id] {
			t.Fatalf("expected Darwin runner candidate %q in %#v", id, candidates)
		}
	}
}

func TestRunSmokeDiscoversWine64OnPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("PATH-based shell runner fixture is not portable to Windows hosts")
	}
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	writeNamedFakeRunner(t, tempDir, "wine64", "win64", 0, DefaultMarker+"\n")
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
	if err := os.WriteFile(executablePath, minimalPEFixture(0x8664), 0o600); err != nil {
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

func TestRunSmokeRejectsNonPEExecutableBeforeRunnerResolution(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "hello.exe")
	if err := os.WriteFile(path, []byte("plain text"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: path,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
	})
	if err == nil || !strings.Contains(err.Error(), "Windows PE file") {
		t.Fatalf("expected Windows PE validation error, got result=%#v err=%v", result, err)
	}
	if result.ExecutableName != "hello.exe" ||
		result.ExecutableFormat != "unknown" ||
		result.WindowsExecutableSignature ||
		result.RunnerAvailable ||
		result.IsolatedStateRoot {
		t.Fatalf("unexpected pre-run validation result: %#v", result)
	}
}

func TestRunSmokeBlocksUnsupportedArchitectureBeforeRunnerResolution(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "arm-app.exe")
	if err := os.WriteFile(path, minimalPEFixture(0xaa64), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := writeFakeRunner(t, tempDir, 0, DefaultMarker+"\n")

	result, err := RunSmoke(context.Background(), Request{
		ExecutablePath: path,
		StateRoot:      filepath.Join(tempDir, "state"),
		RunnerPath:     runnerPath,
	})
	if err != nil {
		t.Fatalf("RunSmoke returned error: %v", err)
	}
	if result.Status != FailedStatus ||
		result.ExecutableName != "arm-app.exe" ||
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "arm64" ||
		result.ExecutableArchitectureReady ||
		result.FailureReason != "Windows executable architecture is not supported" ||
		result.RunnerAvailable ||
		result.IsolatedStateRoot {
		t.Fatalf("unexpected unsupported architecture result: %#v", result)
	}
}

func writeFakeRunner(t *testing.T, tempDir string, exitCode int, stdout string) string {
	return writeNamedFakeRunner(t, tempDir, "fake-runner", "win64", exitCode, stdout)
}

func writeFakeRunnerExpectingWineArchitecture(t *testing.T, tempDir string, wineArchitecture string, exitCode int, stdout string) string {
	return writeNamedFakeRunner(t, tempDir, "fake-runner", wineArchitecture, exitCode, stdout)
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

func writeNamedFakeRunner(t *testing.T, tempDir string, name string, wineArchitecture string, exitCode int, stdout string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}
	path := filepath.Join(tempDir, name)
	body := "#!/bin/sh\n" +
		"test -n \"$WINEPREFIX\" || exit 89\n" +
		"test \"$WINEARCH\" = " + wineArchitecture + " || exit 88\n" +
		"test \"$WINEDEBUG\" = -all || exit 87\n" +
		"case \"$WINEDLLOVERRIDES\" in *winemenubuilder.exe=d*mscoree=d*mshtml=d*) ;; *) exit 86 ;; esac\n" +
		"printf '%s' '" + strings.ReplaceAll(stdout, "'", "'\\''") + "'\n" +
		"exit " + string(rune('0'+exitCode)) + "\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	return path
}
