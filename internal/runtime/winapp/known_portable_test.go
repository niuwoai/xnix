package winapp

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestKnownPortableCatalogContainsPinned7ZipConsoleExecutable(t *testing.T) {
	app, err := LookupKnownPortableApp("7zr")
	if err != nil {
		t.Fatalf("LookupKnownPortableApp returned error: %v", err)
	}
	if app.DisplayName != "7-Zip standalone console executable" ||
		app.Version != "26.02" ||
		app.Architecture != "windows-x86" ||
		app.ExecutableName != "7zr.exe" ||
		app.SourcePageURL != "https://www.7-zip.org/download.html" ||
		app.DownloadURL != "https://github.com/ip7z/7zip/releases/download/26.02/7zr.exe" ||
		app.SHA256 != "56b8cc9f4971cef253644fafe54063ed7fdca551d4dee0f8c6baa81b855acd72" ||
		app.ExpectedMarker != "7-Zip" {
		t.Fatalf("unexpected 7zr catalog entry: %#v", app)
	}
}

func TestFetchKnownPortableAppDownloadsAndVerifiesPinnedArtifact(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(string(body))),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := FetchKnownPortableApp(context.Background(), KnownFetchRequest{
		AppID:         "fixture",
		CacheRoot:     t.TempDir(),
		AllowDownload: true,
		HTTPClient:    client,
		Timeout:       5 * time.Second,
	})
	if err != nil {
		t.Fatalf("FetchKnownPortableApp returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.AppID != "fixture" ||
		result.ExecutableName != "fixture.exe" ||
		result.CacheRelativePath != "fixture/fixture.exe" ||
		result.CacheStatus != "verified" ||
		!result.Downloaded ||
		!result.ChecksumVerified ||
		result.ExpectedSHA256 != hex.EncodeToString(sum[:]) ||
		result.ActualSHA256 != hex.EncodeToString(sum[:]) ||
		!result.NetworkRequired ||
		result.HostRootModified ||
		result.PrivilegedRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed {
		t.Fatalf("unexpected fetch result: %#v", result)
	}
}

func TestPreviewKnownPortableManagedLaunchBlocksUntilArtifactIsVerified(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableManagedLaunch(KnownManagedLaunchRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableManagedLaunch returned error: %v", err)
	}
	if result.SchemaVersion != KnownManagedLaunchSchemaVersion ||
		result.RequestType != KnownManagedLaunchRequestType ||
		result.Status != "needs-artifact" ||
		result.AppID != "7zr" ||
		result.DisplayName != "7-Zip standalone console executable" ||
		result.AppVersion != "26.02" ||
		result.Architecture != "windows-x86" ||
		result.LaunchSurfaceID != "known-app-7zr" ||
		result.DesktopActionID != "launch-known-app-7zr" ||
		result.ManagedLauncher != "xnix-compat-launch --app 7zr" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr" ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		result.LaunchEnabled ||
		!result.PreparationRequired ||
		!result.ManagedLaunchSurface ||
		!result.RuntimeOwnedLaunch ||
		!result.KDEPresentationOnly ||
		!result.RealAppSmokeGateRequired ||
		result.RealAppSmokeGate != "managed-known-app-guest-smoke" ||
		!result.LoopbackOnlyNetworking ||
		!result.GuestRuntimeRequired ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected managed launch preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestMaterializeKnownPortableLaunchProfileSkipsUntilArtifactIsVerified(t *testing.T) {
	tempDir := t.TempDir()

	result, err := MaterializeKnownPortableLaunchProfile(KnownLaunchProfileMaterializeRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
		StateRoot: filepath.Join(tempDir, "state"),
	})
	if err != nil {
		t.Fatalf("MaterializeKnownPortableLaunchProfile returned error: %v", err)
	}
	if result.SchemaVersion != KnownLaunchProfileSchemaVersion ||
		result.RequestType != KnownLaunchProfileRequestType ||
		result.Status != SkippedStatus ||
		result.AppID != "7zr" ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		result.ProfileWritten ||
		result.LauncherBundleWritten ||
		result.LauncherMode != LauncherModeLaunch ||
		result.LauncherCommand != LaunchProfileRequestType ||
		result.NetworkRequired ||
		result.HostRootModified ||
		result.RawExecutablePathExposed ||
		result.RawProfilePathExposed ||
		result.RawStateRootPathExposed ||
		result.RawRuntimeArgvExposed ||
		result.SkipReason != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected skipped known launch profile materialization: %#v", result)
	}
	assertKnownLaunchProfileMaterializeSafe(t, result, tempDir)
}

func TestPrepareKnownPortableLaunchProfileSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PrepareKnownPortableLaunchProfile(context.Background(), KnownPrepareLaunchProfileRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
		StateRoot: filepath.Join(tempDir, "state"),
	})
	if err != nil {
		t.Fatalf("PrepareKnownPortableLaunchProfile returned error: %v", err)
	}
	if result.SchemaVersion != KnownPrepareLaunchSchemaVersion ||
		result.RequestType != KnownPrepareLaunchRequestType ||
		result.Status != SkippedStatus ||
		result.AppID != "7zr" ||
		result.AllowDownload ||
		result.NetworkRequired ||
		result.FetchStatus != SkippedStatus ||
		result.FetchCacheStatus != "missing" ||
		result.Downloaded ||
		result.ChecksumVerified ||
		result.MaterializeStatus != "not-run" ||
		result.ProfileWritten ||
		result.LauncherBundleWritten ||
		result.HostRootModified ||
		result.RawExecutablePathExposed ||
		result.RawProfilePathExposed ||
		result.RawStateRootPathExposed ||
		result.RawRuntimeArgvExposed ||
		result.FetchPayload.NetworkRequired ||
		result.SkipReason != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected offline prepare result: %#v", result)
	}
	if _, err := os.Stat(filepath.Join(tempDir, "state")); err == nil {
		t.Fatalf("offline missing prepare must not create state root")
	}
}

func TestPrepareKnownPortableLaunchProfileDownloadsAndMaterializesWhenAllowed(t *testing.T) {
	body := minimalPEFixture(0x014c)
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	stateRoot := filepath.Join(cacheRoot, "fixture-state")
	client := &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(string(body))),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
		Arguments:      []string{"--help"},
	}})

	result, err := PrepareKnownPortableLaunchProfile(context.Background(), KnownPrepareLaunchProfileRequest{
		AppID:            "fixture",
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		RuntimeBinary:    "go",
		RuntimeArguments: []string{"run", "./cmd/xnix-runtime-go"},
		RunnerPath:       filepath.Join(cacheRoot, "private-runner"),
		RunnerBottle:     "private-bottle",
		RunnerArguments:  []string{"--private-runner-arg"},
		SkipBootstrap:    true,
		AllowDownload:    true,
		HTTPClient:       client,
		Timeout:          5 * time.Second,
	})
	if err != nil {
		t.Fatalf("PrepareKnownPortableLaunchProfile returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		!result.AllowDownload ||
		!result.NetworkRequired ||
		result.FetchStatus != PassedStatus ||
		result.FetchCacheStatus != "verified" ||
		!result.Downloaded ||
		!result.ChecksumVerified ||
		result.MaterializeStatus != PassedStatus ||
		!result.ProfileWritten ||
		!result.LauncherBundleWritten ||
		!result.RunnerConfigured ||
		!result.RunnerBottleConfigured ||
		result.RunnerArgumentCount != 3 ||
		!result.SkipBootstrap ||
		result.MaterializePayload == nil ||
		result.MaterializePayload.LauncherCommand != LaunchProfileRequestType ||
		!result.MaterializePayload.RunnerConfigured ||
		!result.MaterializePayload.RunnerBottleConfigured ||
		result.MaterializePayload.RunnerArgumentCount != 3 ||
		!result.MaterializePayload.SkipBootstrap ||
		result.RawExecutablePathExposed ||
		result.RawProfilePathExposed ||
		result.RawStateRootPathExposed ||
		result.RawRuntimeArgvExposed ||
		result.HostRootModified {
		t.Fatalf("unexpected allowed prepare result: %#v", result)
	}
	profilePath := filepath.Join(stateRoot, "profiles", "fixture.windows-app-smoke-profile.json")
	profileRequest, err := LoadSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("LoadSmokeProfile returned error: %v", err)
	}
	if profileRequest.ExpectedMarker != "FIXTURE_OK" ||
		!profileRequest.RedactOutput ||
		profileRequest.RunnerPath != filepath.Join(cacheRoot, "private-runner") ||
		profileRequest.RunnerBottle != "private-bottle" ||
		len(profileRequest.RunnerArguments) != 1 ||
		profileRequest.RunnerArguments[0] != "--private-runner-arg" ||
		!profileRequest.SkipBootstrap ||
		!profileRequest.StageAppDir ||
		len(profileRequest.Arguments) != 1 ||
		profileRequest.Arguments[0] != "--help" {
		t.Fatalf("unexpected prepared profile request: %#v", profileRequest)
	}
	assertKnownPrepareLaunchProfileSafe(t, result, cacheRoot, "private-bottle", "--private-runner-arg")
}

func TestPrepareAndLaunchKnownPortableProfileSkipsOfflineMissingArtifact(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PrepareAndLaunchKnownPortableProfile(context.Background(), KnownPrepareAndLaunchProfileRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
		StateRoot: filepath.Join(tempDir, "state"),
	})
	if err != nil {
		t.Fatalf("PrepareAndLaunchKnownPortableProfile returned error: %v", err)
	}
	if result.SchemaVersion != KnownPrepareAndLaunchSchemaVersion ||
		result.RequestType != KnownPrepareAndLaunchRequestType ||
		result.Status != SkippedStatus ||
		result.AppID != "7zr" ||
		result.AllowDownload ||
		result.PrepareStatus != SkippedStatus ||
		result.LaunchStatus != "not-run" ||
		result.ProfileWritten ||
		result.LauncherBundleWritten ||
		result.LaunchAttempted ||
		result.LaunchPayload != nil ||
		result.RunnerAvailable ||
		result.WineExecuted ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.NetworkChecksRun ||
		result.PackageManagerInvoked ||
		result.SkipReason != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected offline prepare-and-launch result: %#v", result)
	}
	if _, err := os.Stat(filepath.Join(tempDir, "state")); err == nil {
		t.Fatalf("offline missing prepare-and-launch must not create state root")
	}
}

func TestPrepareAndLaunchKnownPortableProfileRunsVerifiedArtifactWithReadyRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell runner fixture is not portable to Windows hosts")
	}

	body := minimalPEFixture(0x014c)
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(appDir, "fixture.exe"), body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	runnerPath := filepath.Join(cacheRoot, "private-runner")
	if err := os.WriteFile(runnerPath, []byte("#!/bin/sh\nprintf 'FIXTURE_OK\\nraw-host-path=/private/tmp/secret\\n'\n"), 0o700); err != nil {
		t.Fatalf("WriteFile runner returned error: %v", err)
	}
	stateRoot := filepath.Join(cacheRoot, "fixture-state")

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
		Arguments:      []string{"--help"},
	}})

	result, err := PrepareAndLaunchKnownPortableProfile(context.Background(), KnownPrepareAndLaunchProfileRequest{
		AppID:            "fixture",
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		RuntimeBinary:    "go",
		RuntimeArguments: []string{"run", "./cmd/xnix-runtime-go"},
		RunnerPath:       runnerPath,
		RunnerBottle:     "private-bottle",
		RunnerArguments:  []string{"--private-runner-arg"},
		SkipBootstrap:    true,
	})
	if err != nil {
		t.Fatalf("PrepareAndLaunchKnownPortableProfile returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.PrepareStatus != PassedStatus ||
		result.LaunchStatus != PassedStatus ||
		!result.ProfileWritten ||
		!result.LauncherBundleWritten ||
		!result.LaunchAttempted ||
		!result.RunnerConfigured ||
		!result.RunnerBottleConfigured ||
		result.RunnerArgumentCount != 3 ||
		!result.RunnerAvailable ||
		!result.SkipBootstrap ||
		result.ExecutableName != "fixture.exe" ||
		result.ExecutableFormat != "pe-mz" ||
		!result.WindowsExecutableSignature ||
		result.ExecutableArchitecture != "x86" ||
		result.WineArchitecture != "win32" ||
		result.ApplicationWorkspaceMode != ApplicationWorkspaceModeStaged ||
		!result.RawOutputRedacted ||
		result.LaunchPayload == nil ||
		result.LaunchPayload.RuntimePayload == nil ||
		!result.WineExecuted ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.NetworkChecksRun ||
		result.PackageManagerInvoked ||
		result.RawExecutablePathExposed ||
		result.RawProfilePathExposed ||
		result.RawRunnerPathExposed ||
		result.HostRootModified {
		t.Fatalf("unexpected ready prepare-and-launch result: %#v", result)
	}
	if result.LaunchPayload.RuntimePayload.Stdout != "" ||
		result.LaunchPayload.RuntimePayload.Stderr != "" ||
		!strings.Contains(result.LaunchPayload.RuntimePayload.KDESafeOutputSummary, "expected smoke marker observed") {
		t.Fatalf("unexpected redacted runtime payload: %#v", result.LaunchPayload.RuntimePayload)
	}
	assertKnownPrepareAndLaunchProfileSafe(t, result, cacheRoot, "private-bottle", "--private-runner-arg", "/private/tmp/secret")
}

func TestRunKnownPortableAppLocalBackendUsesPrepareAndLaunchPath(t *testing.T) {
	tempDir := t.TempDir()

	result, err := RunKnownPortableApp(context.Background(), KnownRunRequest{
		AppID:     "7zr",
		Backend:   KnownRunBackendLocal,
		CacheRoot: tempDir,
		StateRoot: filepath.Join(tempDir, "state"),
	})
	if err != nil {
		t.Fatalf("RunKnownPortableApp returned error: %v", err)
	}
	if result.SchemaVersion != KnownRunSchemaVersion ||
		result.RequestType != KnownRunRequestType ||
		result.Status != SkippedStatus ||
		result.AppID != "7zr" ||
		result.Backend != KnownRunBackendLocal ||
		result.BackendReady ||
		result.LaunchAttempted ||
		result.RunnerAvailable ||
		result.ChecksumVerified ||
		result.ProfileWritten ||
		result.LauncherBundleWritten ||
		result.LocalPayload == nil ||
		result.GuestPayload != nil ||
		result.LoopbackOnlyNetworking ||
		result.QEMURequired ||
		result.WineExecuted ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.NetworkChecksRun ||
		result.PackageManagerInvoked ||
		result.SkipReason != "known Windows app artifact unavailable" {
		t.Fatalf("unexpected local known app run result: %#v", result)
	}
}

func TestRunKnownPortableAppGuestWineBackendUsesVerifiedCacheAndLoopbackGuest(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(appDir, "fixture.exe"), body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	guestRoot := t.TempDir()
	logPath := filepath.Join(guestRoot, "guest.log")
	result, err := RunKnownPortableApp(context.Background(), KnownRunRequest{
		AppID:     "fixture",
		Backend:   KnownRunBackendGuestWine,
		CacheRoot: cacheRoot,
		Host:      "127.0.0.1",
		Port:      "2222",
		User:      "root",
		KeyPath:   filepath.Join(guestRoot, "id_ed25519"),
		RemoteDir: "/tmp/xnix-known-winapp-smoke",
		SSHPath:   writeFakeKnownAppGuestSSH(t, guestRoot, logPath),
		SCPPath:   writeFakeGuestSCP(t, guestRoot, logPath),
		Timeout:   5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableApp returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.Backend != KnownRunBackendGuestWine ||
		!result.BackendReady ||
		!result.LaunchAttempted ||
		!result.RunnerAvailable ||
		!result.ChecksumVerified ||
		!result.ExecutableCopied ||
		!result.MarkerObserved ||
		result.LocalPayload != nil ||
		result.GuestPayload == nil ||
		!result.LoopbackOnlyNetworking ||
		!result.QEMURequired ||
		!result.WineExecuted ||
		result.DockerExecuted ||
		result.QEMUExecuted ||
		result.NetworkChecksRun ||
		result.PackageManagerInvoked ||
		result.RawHostPathExposed ||
		result.HostRootModified {
		t.Fatalf("unexpected guest known app run result: %#v", result)
	}
}

func TestRunKnownPortableAppGuestWineBackendCanStartQEMUFromRuntime(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(appDir, "fixture.exe"), body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	guestRoot := t.TempDir()
	guestLogPath := filepath.Join(guestRoot, "guest.log")
	serialLogPath := filepath.Join(guestRoot, "qemu-serial.log")
	kernelPath := filepath.Join(guestRoot, "bzImage")
	if err := os.WriteFile(kernelPath, []byte("fake kernel"), 0o600); err != nil {
		t.Fatalf("WriteFile kernel returned error: %v", err)
	}

	result, err := RunKnownPortableApp(context.Background(), KnownRunRequest{
		AppID:           "fixture",
		Backend:         KnownRunBackendGuestWine,
		CacheRoot:       cacheRoot,
		Host:            "localhost",
		Port:            AutoGuestPort,
		User:            "root",
		KeyPath:         filepath.Join(guestRoot, "id_ed25519"),
		RemoteDir:       "/tmp/xnix-known-winapp-smoke",
		SSHPath:         writeFakeKnownAppGuestSSH(t, guestRoot, guestLogPath),
		SCPPath:         writeFakeGuestSCP(t, guestRoot, guestLogPath),
		Timeout:         5 * time.Second,
		StartQEMU:       true,
		QEMUBinary:      writeFakeQEMU(t, guestRoot),
		QEMUKernelImage: kernelPath,
		QEMUBootTimeout: 5 * time.Second,
		QEMUSerialLog:   serialLogPath,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableApp returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.Backend != KnownRunBackendGuestWine ||
		!result.GuestStartAttempted ||
		!result.GuestStarted ||
		result.GuestStartMode != "go-qemu" ||
		result.GuestHost != "127.0.0.1" ||
		result.GuestPort == "" ||
		!result.GuestPortAuto ||
		!result.QEMUSerialLogWritten ||
		!result.QEMUExecuted ||
		!result.BackendReady ||
		!result.WineExecuted ||
		!result.ChecksumVerified ||
		!result.MarkerObserved ||
		result.RawQEMUPathExposed ||
		result.RawHostPathExposed ||
		result.HostRootModified {
		t.Fatalf("unexpected Go-started QEMU known app run result: %#v", result)
	}
	serialBytes, err := os.ReadFile(serialLogPath)
	if err != nil {
		t.Fatalf("ReadFile serial log returned error: %v", err)
	}
	if !strings.Contains(string(serialBytes), "fake qemu boot") {
		t.Fatalf("serial log did not capture fake qemu output: %s", string(serialBytes))
	}
	if !strings.Contains(string(serialBytes), "hostfwd=tcp:127.0.0.1:"+result.GuestPort+"-:22") {
		t.Fatalf("serial log did not capture allocated loopback port %q: %s", result.GuestPort, string(serialBytes))
	}
	guestBytes, err := os.ReadFile(guestLogPath)
	if err != nil {
		t.Fatalf("ReadFile guest log returned error: %v", err)
	}
	if !strings.Contains(string(guestBytes), "-p "+result.GuestPort) ||
		!strings.Contains(string(guestBytes), "-P "+result.GuestPort) ||
		!strings.Contains(string(guestBytes), "root@127.0.0.1") ||
		strings.Contains(string(guestBytes), "root@localhost") {
		t.Fatalf("guest log did not use allocated port %q: %s", result.GuestPort, string(guestBytes))
	}
}

func TestRunKnownPortableAppGuestWineStartQEMUSkipsBeforeVMWhenArtifactMissing(t *testing.T) {
	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         strings.Repeat("0", 64),
		ExpectedMarker: "FIXTURE_OK",
	}})

	guestRoot := t.TempDir()
	result, err := RunKnownPortableApp(context.Background(), KnownRunRequest{
		AppID:           "fixture",
		Backend:         KnownRunBackendGuestWine,
		CacheRoot:       t.TempDir(),
		StartQEMU:       true,
		QEMUBinary:      filepath.Join(guestRoot, "must-not-run-qemu"),
		QEMUKernelImage: filepath.Join(guestRoot, "must-not-check-kernel"),
		QEMUBootTimeout: time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableApp returned error: %v", err)
	}
	if result.Status != SkippedStatus ||
		result.GuestStartMode != "go-qemu" ||
		result.GuestStartAttempted ||
		result.GuestStarted ||
		result.QEMUExecuted ||
		result.ChecksumVerified ||
		result.GuestPayload == nil ||
		result.SkipReason != "known Windows app artifact unavailable or checksum mismatch" {
		t.Fatalf("unexpected preflight skip result: %#v", result)
	}
}

func TestMaterializeKnownPortableLaunchProfileWritesProfileAndLaunchBundleForVerifiedArtifact(t *testing.T) {
	body := minimalPEFixture(0x014c)
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(cacheRoot, "fixture-state")

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
		Arguments:      []string{"--help"},
	}})

	result, err := MaterializeKnownPortableLaunchProfile(KnownLaunchProfileMaterializeRequest{
		AppID:            "fixture",
		CacheRoot:        cacheRoot,
		StateRoot:        stateRoot,
		ApplicationID:    "org.xnix.known.fixture",
		RuntimeBinary:    "go",
		RuntimeArguments: []string{"run", "./cmd/xnix-runtime-go"},
		RunnerPath:       filepath.Join(cacheRoot, "private-runner"),
		RunnerBottle:     "private-bottle",
		RunnerArguments:  []string{"--private-runner-arg"},
		SkipBootstrap:    true,
	})
	if err != nil {
		t.Fatalf("MaterializeKnownPortableLaunchProfile returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.ProfileWritten ||
		result.ProfileFileName != "fixture.windows-app-smoke-profile.json" ||
		!result.LauncherBundleWritten ||
		result.LauncherMode != LauncherModeLaunch ||
		result.LauncherCommand != LaunchProfileRequestType ||
		!result.RunnerConfigured ||
		!result.RunnerBottleConfigured ||
		result.RunnerArgumentCount != 3 ||
		!result.SkipBootstrap ||
		result.ExpectedMarker != "FIXTURE_OK" ||
		result.SuccessMode != SuccessModeMarker ||
		result.ApplicationWorkspaceMode != ApplicationWorkspaceModeStaged ||
		result.NetworkRequired ||
		result.HostRootModified ||
		result.RawExecutablePathExposed ||
		result.RawProfilePathExposed ||
		result.RawStateRootPathExposed ||
		result.RawRuntimeArgvExposed ||
		result.LauncherBundlePayload == nil ||
		result.LauncherBundlePayload.LauncherCommand != LaunchProfileRequestType ||
		result.LauncherBundlePayload.RuntimeArgumentCount != 2 {
		t.Fatalf("unexpected ready known launch profile materialization: %#v", result)
	}
	profilePath := filepath.Join(stateRoot, "profiles", "fixture.windows-app-smoke-profile.json")
	profileRequest, err := LoadSmokeProfile(profilePath)
	if err != nil {
		t.Fatalf("LoadSmokeProfile returned error: %v", err)
	}
	if profileRequest.ExecutablePath != executablePath ||
		profileRequest.StateRoot != stateRoot ||
		profileRequest.RunnerPath != filepath.Join(cacheRoot, "private-runner") ||
		profileRequest.RunnerBottle != "private-bottle" ||
		len(profileRequest.RunnerArguments) != 1 ||
		profileRequest.RunnerArguments[0] != "--private-runner-arg" ||
		profileRequest.ExpectedMarker != "FIXTURE_OK" ||
		profileRequest.SuccessMode != SuccessModeMarker ||
		!profileRequest.RedactOutput ||
		!profileRequest.SkipBootstrap ||
		!profileRequest.StageAppDir ||
		len(profileRequest.Arguments) != 1 ||
		profileRequest.Arguments[0] != "--help" {
		t.Fatalf("unexpected materialized profile request: %#v", profileRequest)
	}
	launcherPath := filepath.Join(stateRoot, "launcher-bundle", "launchers", "org.xnix.known.fixture.sh")
	launcherText, err := os.ReadFile(launcherPath)
	if err != nil {
		t.Fatalf("ReadFile launcher returned error: %v", err)
	}
	if !strings.Contains(string(launcherText), "windows-app-launch-profile --profile") ||
		strings.Contains(string(launcherText), "windows-app-run-smoke") {
		t.Fatalf("unexpected launch profile launcher: %s", string(launcherText))
	}
	assertKnownLaunchProfileMaterializeSafe(t, result, cacheRoot)
	assertKnownLaunchProfileMaterializeSafe(t, result, "private-bottle")
	assertKnownLaunchProfileMaterializeSafe(t, result, "--private-runner-arg")
}

func TestPreviewKnownPortableManagedLaunchEnablesVerifiedKnownArtifact(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := PreviewKnownPortableManagedLaunch(KnownManagedLaunchRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableManagedLaunch returned error: %v", err)
	}
	if result.Status != "ready" ||
		result.AppID != "fixture" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.LaunchEnabled ||
		result.PreparationRequired ||
		result.BlockedReason != "" ||
		result.ManagedLauncher != "xnix-compat-launch --app fixture" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app fixture" ||
		!strings.Contains(result.DesktopSafeSummary, "ready to launch") {
		t.Fatalf("unexpected ready managed launch preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestPreviewKnownPortableKDELauncherConsumesManagedLaunchSurface(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableKDELauncher(KnownKDELauncherRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableKDELauncher returned error: %v", err)
	}
	if result.SchemaVersion != KnownKDELauncherSchemaVersion ||
		result.RequestType != KnownKDELauncherRequestType ||
		result.Source != KnownManagedLaunchRequestType ||
		result.Status != "visible-needs-preparation" ||
		result.Desktop != "KDE Plasma" ||
		result.EntryPointID != "launcher" ||
		result.KDEComponent != "Plasma application launcher" ||
		result.AppID != "7zr" ||
		result.DisplayName != "7-Zip standalone console executable" ||
		result.DesktopFile != "xnix-known-app-7zr.desktop" ||
		result.Icon != "xnix-known-app-7zr" ||
		result.LaunchSurfaceID != "known-app-7zr" ||
		result.DesktopActionID != "launch-known-app-7zr" ||
		result.ManagedLauncher != "xnix-compat-launch --app 7zr" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr" ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		!result.LaunchVisible ||
		result.LaunchEnabled ||
		!result.PreparationRequired ||
		!result.ManagedLaunchSurface ||
		!result.RuntimeOwnedLaunch ||
		!result.KDEPresentationOnly ||
		!result.DesktopEntryPreviewCreated ||
		result.DesktopFilesWritten ||
		result.MIMEAppsWritten ||
		result.BackendProcessStarted ||
		result.HostRootModified ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected KDE launcher preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestPreviewKnownPortableKDELauncherEnablesVisibleLauncherForVerifiedArtifact(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := PreviewKnownPortableKDELauncher(KnownKDELauncherRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableKDELauncher returned error: %v", err)
	}
	if result.Status != "visible-ready" ||
		result.AppID != "fixture" ||
		result.DesktopFile != "xnix-known-app-fixture.desktop" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.LaunchVisible ||
		!result.LaunchEnabled ||
		result.PreparationRequired ||
		result.BlockedReason != "" ||
		!strings.Contains(result.DesktopSafeSummary, "visible in the KDE launcher and ready") {
		t.Fatalf("unexpected ready KDE launcher preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestPreviewKnownPortableLaunchRequestConsumesKDELauncher(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableLaunchRequest(KnownLaunchRequestRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableLaunchRequest returned error: %v", err)
	}
	if result.SchemaVersion != KnownLaunchRequestSchemaVersion ||
		result.RequestType != KnownLaunchRequestType ||
		result.Source != KnownKDELauncherRequestType ||
		result.Status != "request-blocked" ||
		result.RequestID != "known-app-launch-request-7zr" ||
		result.RuntimeMethod != "LaunchKnownWindowsApp" ||
		result.AppID != "7zr" ||
		result.DisplayName != "7-Zip standalone console executable" ||
		result.Desktop != "KDE Plasma" ||
		result.EntryPointID != "launcher" ||
		result.DesktopFile != "xnix-known-app-7zr.desktop" ||
		result.LaunchSurfaceID != "known-app-7zr" ||
		result.DesktopActionID != "launch-known-app-7zr" ||
		result.ManagedLauncher != "xnix-compat-launch --app 7zr" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr" ||
		result.DispatchGate != "managed-known-app-guest-smoke" ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		!result.LaunchVisible ||
		result.LaunchAllowed ||
		!result.LaunchRequestCreated ||
		result.DispatchReady ||
		!result.PreparationRequired ||
		!result.RuntimeOwnedRequest ||
		!result.RuntimeOwnedLaunch ||
		!result.KDEPresentationOnly ||
		!result.DryRun ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		result.HostRootModified ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected launch request preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestPreviewKnownPortableLaunchRequestReadiesVerifiedArtifactForDispatch(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := PreviewKnownPortableLaunchRequest(KnownLaunchRequestRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableLaunchRequest returned error: %v", err)
	}
	if result.Status != "request-ready" ||
		result.RequestID != "known-app-launch-request-fixture" ||
		result.AppID != "fixture" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.LaunchAllowed ||
		!result.LaunchRequestCreated ||
		!result.DispatchReady ||
		result.PreparationRequired ||
		result.BlockedReason != "" ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		!strings.Contains(result.DesktopSafeSummary, "Runtime-owned launch request ready") {
		t.Fatalf("unexpected ready launch request preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestPreviewKnownPortableDispatchConsumesLaunchRequest(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableDispatch(KnownDispatchRequest{
		AppID:     "7zr",
		CacheRoot: tempDir,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableDispatch returned error: %v", err)
	}
	if result.SchemaVersion != KnownDispatchSchemaVersion ||
		result.RequestType != KnownDispatchRequestType ||
		result.Source != KnownLaunchRequestType ||
		result.Status != "dispatch-blocked" ||
		result.RequestID != "known-app-launch-request-7zr" ||
		result.DispatchID != "known-app-dispatch-7zr" ||
		result.RuntimeMethod != "DispatchKnownWindowsApp" ||
		result.AppID != "7zr" ||
		result.DisplayName != "7-Zip standalone console executable" ||
		result.Desktop != "KDE Plasma" ||
		result.EntryPointID != "launcher" ||
		result.DesktopFile != "xnix-known-app-7zr.desktop" ||
		result.LaunchSurfaceID != "known-app-7zr" ||
		result.DesktopActionID != "launch-known-app-7zr" ||
		result.ManagedLauncher != "xnix-compat-launch --app 7zr" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr" ||
		result.DispatchGate != "managed-known-app-guest-smoke" ||
		result.RunnerLane != "known-app-guest-smoke" ||
		result.RunnerRequestType != "managed-known-app-guest-smoke" ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		!result.LaunchRequestCreated ||
		!result.DispatchPreviewCreated ||
		result.DispatchAllowed ||
		result.DispatchReady ||
		!result.PreparationRequired ||
		!result.RuntimeOwnedRequest ||
		!result.RuntimeOwnedLaunch ||
		!result.RuntimeOwnedDispatch ||
		!result.KDEPresentationOnly ||
		!result.DryRun ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		result.HostRootModified ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected dispatch preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestPreviewKnownPortableDispatchReadiesVerifiedArtifact(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := PreviewKnownPortableDispatch(KnownDispatchRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableDispatch returned error: %v", err)
	}
	if result.Status != "dispatch-ready" ||
		result.RequestID != "known-app-launch-request-fixture" ||
		result.DispatchID != "known-app-dispatch-fixture" ||
		result.AppID != "fixture" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.LaunchRequestCreated ||
		!result.DispatchPreviewCreated ||
		!result.DispatchAllowed ||
		!result.DispatchReady ||
		result.PreparationRequired ||
		result.BlockedReason != "" ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		!strings.Contains(result.DesktopSafeSummary, "ready for Runtime-owned managed guest dispatch") {
		t.Fatalf("unexpected ready dispatch preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestRunKnownPortableDispatchSmokeBlocksUntilArtifactVerified(t *testing.T) {
	tempDir := t.TempDir()

	result, err := RunKnownPortableDispatchSmoke(context.Background(), KnownDispatchSmokeRequest{
		AppID:         "7zr",
		CacheRoot:     tempDir,
		GuestBoundary: KnownDispatchGuestBoundary,
		Timeout:       5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableDispatchSmoke returned error: %v", err)
	}
	if result.SchemaVersion != KnownDispatchSmokeSchemaVersion ||
		result.RequestType != KnownDispatchSmokeRequestType ||
		result.Source != KnownDispatchRequestType ||
		result.Status != "dispatch-blocked" ||
		result.RequestID != "known-app-launch-request-7zr" ||
		result.DispatchID != "known-app-dispatch-7zr" ||
		result.RuntimeMethod != "DispatchKnownWindowsApp" ||
		result.AppID != "7zr" ||
		result.DispatchGate != KnownDispatchGuestBoundary ||
		result.RunnerLane != "known-app-guest-smoke" ||
		result.GuestBoundary != KnownDispatchGuestBoundary ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		!result.LaunchRequestCreated ||
		!result.DispatchPreviewCreated ||
		result.DispatchReady ||
		result.DispatchAllowed ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.ManagedGuestRunnerInvoked ||
		result.ManagedGuestReachable ||
		result.ManagedGuestRuntimeReady ||
		result.ManagedArtifactCopied ||
		result.MarkerObserved ||
		result.SmokePassed ||
		result.ExitCode != -1 ||
		!result.RuntimeOwnedRequest ||
		!result.RuntimeOwnedLaunch ||
		!result.RuntimeOwnedDispatch ||
		!result.KDEPresentationOnly ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" ||
		result.SkipReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected blocked dispatch smoke: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestRunKnownPortableDispatchSmokeRequiresControlledGuestBoundary(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := RunKnownPortableDispatchSmoke(context.Background(), KnownDispatchSmokeRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
		Timeout:   5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableDispatchSmoke returned error: %v", err)
	}
	if result.Status != "dispatch-blocked" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.DispatchReady ||
		result.DispatchAllowed ||
		result.DispatchStarted ||
		result.ManagedGuestRunnerInvoked ||
		result.BlockedReason != "controlled managed guest boundary must be supplied by the smoke harness" {
		t.Fatalf("unexpected boundary-blocked dispatch smoke: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestRunKnownPortableDispatchSmokeInvokesManagedGuestRunner(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	guestRoot := t.TempDir()
	logPath := filepath.Join(guestRoot, "guest.log")
	sshPath := writeFakeKnownAppGuestSSH(t, guestRoot, logPath)
	scpPath := writeFakeGuestSCP(t, guestRoot, logPath)

	result, err := RunKnownPortableDispatchSmoke(context.Background(), KnownDispatchSmokeRequest{
		AppID:         "fixture",
		CacheRoot:     cacheRoot,
		GuestBoundary: KnownDispatchGuestBoundary,
		Host:          "127.0.0.1",
		Port:          "2222",
		User:          "root",
		KeyPath:       filepath.Join(guestRoot, "id_ed25519"),
		RemoteDir:     "/tmp/xnix-known-winapp-smoke",
		SSHPath:       sshPath,
		SCPPath:       scpPath,
		Timeout:       5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableDispatchSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.RequestID != "known-app-launch-request-fixture" ||
		result.DispatchID != "known-app-dispatch-fixture" ||
		result.AppID != "fixture" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.DispatchReady ||
		!result.DispatchAllowed ||
		!result.DispatchStarted ||
		!result.ExecutionStarted ||
		!result.ManagedGuestRunnerInvoked ||
		!result.ManagedGuestReachable ||
		!result.ManagedGuestRuntimeReady ||
		!result.ManagedArtifactCopied ||
		!result.MarkerObserved ||
		!result.SmokePassed ||
		result.ExitCode != 0 ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		!strings.Contains(result.DesktopSafeSummary, "passed the gated Runtime managed guest dispatch smoke") {
		t.Fatalf("unexpected dispatch smoke result: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "fixture.exe root@127.0.0.1:/tmp/xnix-known-winapp-smoke/fixture.exe") ||
		!strings.Contains(log, "WINEPREFIX='/tmp/xnix-known-winapp-smoke/wineprefix' WINEDEBUG=-all wine '/tmp/xnix-known-winapp-smoke/fixture.exe'") {
		t.Fatalf("guest log missing known app execution: %s", log)
	}
}

func TestPreviewKnownPortableLaunchBridgeConsumesManagedLauncher(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableLaunchBridge(KnownLaunchBridgeRequest{
		AppID:               "7zr",
		CacheRoot:           tempDir,
		ManagedLauncherArgv: []string{"xnix-compat-launch", "--app", "7zr"},
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableLaunchBridge returned error: %v", err)
	}
	if result.SchemaVersion != KnownLaunchBridgeSchemaVersion ||
		result.RequestType != KnownLaunchBridgeRequestType ||
		result.Source != KnownDispatchRequestType ||
		result.Status != "bridge-blocked" ||
		result.Desktop != "KDE Plasma" ||
		result.EntryPointID != "launcher" ||
		result.DesktopFile != "xnix-known-app-7zr.desktop" ||
		result.LaunchSurfaceID != "known-app-7zr" ||
		result.DesktopActionID != "launch-known-app-7zr" ||
		result.ManagedLauncher != "xnix-compat-launch --app 7zr" ||
		strings.Join(result.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr" ||
		!result.LauncherArgvAccepted ||
		result.RequestID != "known-app-launch-request-7zr" ||
		result.DispatchID != "known-app-dispatch-7zr" ||
		result.RuntimeMethod != "BridgeKnownLauncherToDispatchSmoke" ||
		result.DispatchSmokeRequestType != KnownDispatchSmokeRequestType ||
		result.DispatchSmokeRequestMaterialized ||
		result.AppID != "7zr" ||
		result.DispatchGate != KnownDispatchGuestBoundary ||
		result.RunnerLane != "known-app-guest-smoke" ||
		!result.GuestBoundaryRequired ||
		result.GuestBoundarySupplied ||
		!result.SmokeHarnessRequired ||
		result.CacheStatus != "missing" ||
		result.ArtifactVerified ||
		!result.LaunchRequestCreated ||
		!result.DispatchPreviewCreated ||
		!result.BridgePreviewCreated ||
		result.DispatchReady ||
		!result.PreparationRequired ||
		!result.RuntimeOwnedRequest ||
		!result.RuntimeOwnedLaunch ||
		!result.RuntimeOwnedDispatch ||
		!result.RuntimeOwnedBridge ||
		!result.KDEPresentationOnly ||
		!result.DryRun ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		result.HostRootModified ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawExecutablePathExposed ||
		result.RawCommandExposed ||
		result.BackendDetailsExposed ||
		result.BlockedReason != "managed application artifact must be fetched before launch" {
		t.Fatalf("unexpected launch bridge preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestPreviewKnownPortableLaunchBridgeRejectsUnexpectedLauncherArgv(t *testing.T) {
	tempDir := t.TempDir()

	result, err := PreviewKnownPortableLaunchBridge(KnownLaunchBridgeRequest{
		AppID:               "7zr",
		CacheRoot:           tempDir,
		ManagedLauncherArgv: []string{"xnix-compat-launch", "--app", "not-7zr"},
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableLaunchBridge returned error: %v", err)
	}
	if result.Status != "bridge-blocked" ||
		result.LauncherArgvAccepted ||
		result.DispatchSmokeRequestMaterialized ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		result.BlockedReason != "managed launcher argv does not match the Runtime-owned launch surface" {
		t.Fatalf("unexpected rejected launch bridge preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, tempDir)
}

func TestPreviewKnownPortableLaunchBridgeMaterializesVerifiedDispatchSmokeRequest(t *testing.T) {
	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	result, err := PreviewKnownPortableLaunchBridge(KnownLaunchBridgeRequest{
		AppID:               "fixture",
		CacheRoot:           cacheRoot,
		ManagedLauncherArgv: []string{"xnix-compat-launch", "--app", "fixture"},
	})
	if err != nil {
		t.Fatalf("PreviewKnownPortableLaunchBridge returned error: %v", err)
	}
	if result.Status != "bridge-ready" ||
		result.RequestID != "known-app-launch-request-fixture" ||
		result.DispatchID != "known-app-dispatch-fixture" ||
		result.DispatchSmokeRequestType != KnownDispatchSmokeRequestType ||
		!result.DispatchSmokeRequestMaterialized ||
		result.AppID != "fixture" ||
		result.CacheStatus != "verified" ||
		!result.ArtifactVerified ||
		!result.LauncherArgvAccepted ||
		!result.DispatchReady ||
		result.PreparationRequired ||
		result.BlockedReason != "" ||
		result.DispatchStarted ||
		result.ExecutionStarted ||
		result.BackendProcessStarted ||
		!strings.Contains(result.DesktopSafeSummary, "materialized a gated dispatch smoke request") {
		t.Fatalf("unexpected ready launch bridge preview: %#v", result)
	}
	assertManagedLaunchSurfaceSafe(t, result, cacheRoot)
}

func TestRunKnownPortableGuestSmokeUsesVerifiedCacheAndLoopbackGuest(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell ssh fixture is not portable to Windows hosts")
	}

	body := []byte("fixture portable windows executable")
	sum := sha256.Sum256(body)
	cacheRoot := t.TempDir()
	appDir := filepath.Join(cacheRoot, "fixture")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	executablePath := filepath.Join(appDir, "fixture.exe")
	if err := os.WriteFile(executablePath, body, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	withKnownPortableCatalog(t, []KnownPortableApp{{
		ID:             "fixture",
		DisplayName:    "Fixture console executable",
		Version:        "1.0.0",
		Architecture:   "windows-x86",
		ExecutableName: "fixture.exe",
		SourcePageURL:  "https://example.invalid/download",
		DownloadURL:    "https://example.invalid/fixture.exe",
		SHA256:         hex.EncodeToString(sum[:]),
		ExpectedMarker: "FIXTURE_OK",
	}})

	guestRoot := t.TempDir()
	logPath := filepath.Join(guestRoot, "guest.log")
	sshPath := writeFakeKnownAppGuestSSH(t, guestRoot, logPath)
	scpPath := writeFakeGuestSCP(t, guestRoot, logPath)

	result, err := RunKnownPortableGuestSmoke(context.Background(), KnownGuestRequest{
		AppID:     "fixture",
		CacheRoot: cacheRoot,
		Host:      "127.0.0.1",
		Port:      "2222",
		User:      "root",
		KeyPath:   filepath.Join(guestRoot, "id_ed25519"),
		RemoteDir: "/tmp/xnix-known-winapp-smoke",
		SSHPath:   sshPath,
		SCPPath:   scpPath,
		Timeout:   5 * time.Second,
	})
	if err != nil {
		t.Fatalf("RunKnownPortableGuestSmoke returned error: %v", err)
	}
	if result.Status != PassedStatus ||
		result.SchemaVersion != KnownGuestSchemaVersion ||
		result.RequestType != KnownGuestRequestType ||
		result.AppID != "fixture" ||
		result.ExecutableName != "fixture.exe" ||
		result.ExpectedMarker != "FIXTURE_OK" ||
		!result.ChecksumVerified ||
		result.ActualSHA256 != hex.EncodeToString(sum[:]) ||
		result.Guest.Status != PassedStatus ||
		result.Guest.ExecutableName != "fixture.exe" ||
		result.Guest.ExpectedMarker != "FIXTURE_OK" ||
		!result.Guest.MarkerObserved ||
		!result.LoopbackOnlyNetworking ||
		!result.QEMURequired ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed {
		t.Fatalf("unexpected known guest result: %#v", result)
	}

	logBytes, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("ReadFile log returned error: %v", err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "fixture.exe root@127.0.0.1:/tmp/xnix-known-winapp-smoke/fixture.exe") ||
		!strings.Contains(log, "WINEPREFIX='/tmp/xnix-known-winapp-smoke/wineprefix' WINEDEBUG=-all wine '/tmp/xnix-known-winapp-smoke/fixture.exe'") {
		t.Fatalf("guest log missing known app execution: %s", log)
	}
}

func withKnownPortableCatalog(t *testing.T, catalog []KnownPortableApp) {
	t.Helper()
	original := knownPortableCatalog
	knownPortableCatalog = catalog
	t.Cleanup(func() {
		knownPortableCatalog = original
	})
}

func writeFakeKnownAppGuestSSH(t *testing.T, tempDir string, logPath string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-ssh")
	body := "#!/bin/sh\n" +
		"printf 'ssh %s\\n' \"$*\" >> '" + logPath + "'\n" +
		"case \"$*\" in\n" +
		"  *' true') exit 0 ;;\n" +
		"  *'command -v wine'*) exit 0 ;;\n" +
		"  *'mkdir -p'*) exit 0 ;;\n" +
		"  *' wine '*'fixture.exe'*) printf 'FIXTURE_OK\\n'; exit 0 ;;\n" +
		"esac\n" +
		"exit 2\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile fake ssh returned error: %v", err)
	}
	return path
}

func writeFakeQEMU(t *testing.T, tempDir string) string {
	t.Helper()
	path := filepath.Join(tempDir, "fake-qemu")
	body := "#!/bin/sh\n" +
		"printf 'fake qemu boot\\n'\n" +
		"printf 'fake qemu args %s\\n' \"$*\"\n" +
		"trap 'exit 0' TERM INT\n" +
		"while :; do sleep 1; done\n"
	if err := os.WriteFile(path, []byte(body), 0o700); err != nil {
		t.Fatalf("WriteFile fake qemu returned error: %v", err)
	}
	return path
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (function roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return function(request)
}

func assertManagedLaunchSurfaceSafe(t *testing.T, result any, hostPath string) {
	t.Helper()
	text := fmt.Sprintf("%#v", result)
	lower := strings.ToLower(text)
	for _, forbidden := range []string{".exe", "wine", "qemu", hostPath} {
		if strings.Contains(lower, strings.ToLower(forbidden)) {
			t.Fatalf("managed launch preview exposed forbidden term %q: %#v", forbidden, result)
		}
	}
}

func assertKnownLaunchProfileMaterializeSafe(t *testing.T, result KnownLaunchProfileMaterializeResult, hostPath string) {
	t.Helper()
	text := fmt.Sprintf("%#v", result)
	for _, forbidden := range []string{hostPath, filepath.Join(hostPath, "fixture"), filepath.Join(hostPath, "state")} {
		if strings.TrimSpace(forbidden) != "" && strings.Contains(text, forbidden) {
			t.Fatalf("known launch profile materialization exposed raw path %q: %#v", forbidden, result)
		}
	}
}

func assertKnownPrepareLaunchProfileSafe(t *testing.T, result KnownPrepareLaunchProfileResult, forbiddenValues ...string) {
	t.Helper()
	text := fmt.Sprintf("%#v", result)
	for _, forbidden := range forbiddenValues {
		if strings.TrimSpace(forbidden) != "" && strings.Contains(text, forbidden) {
			t.Fatalf("known launch profile preparation exposed forbidden value %q: %#v", forbidden, result)
		}
	}
}

func assertKnownPrepareAndLaunchProfileSafe(t *testing.T, result KnownPrepareAndLaunchProfileResult, forbiddenValues ...string) {
	t.Helper()
	text := fmt.Sprintf("%#v", result)
	for _, forbidden := range forbiddenValues {
		if strings.TrimSpace(forbidden) != "" && strings.Contains(text, forbidden) {
			t.Fatalf("known prepare-and-launch profile exposed forbidden value %q: %#v", forbidden, result)
		}
	}
}
