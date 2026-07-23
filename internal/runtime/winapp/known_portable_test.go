package winapp

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
