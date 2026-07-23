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
