package appidentity

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

func TestImportKnownPortableBundleDownloadsExtractsAndRecordsImport(t *testing.T) {
	archive := fixtureKnownPortableBundleArchive(t, map[string][]byte{
		"fixture.exe":        {'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'},
		"config.xml":         []byte("<config/>"),
		"plugins/plugin.dll": []byte("plugin"),
		"locales/en-US.xml":  []byte("<locale/>"),
	})
	sum := sha256.Sum256(archive)
	app := fixtureKnownPortableBundleApp(hex.EncodeToString(sum[:]))
	tempDir := t.TempDir()
	client := &http.Client{Transport: bundleRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.String() != app.DownloadURL {
			return nil, fmt.Errorf("unexpected download URL: %s", request.URL.String())
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(archive)),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})}

	result, err := ImportKnownPortableBundle(context.Background(), KnownPortableBundleImportRequest{
		Version:       "0.2.640-test",
		AppID:         app.ID,
		CatalogApp:    app,
		CacheRoot:     filepath.Join(tempDir, "cache"),
		StateRoot:     filepath.Join(tempDir, "state"),
		AllowDownload: true,
		HTTPClient:    client,
		Timeout:       5 * time.Second,
		ReportFetch:   true,
	})
	if err != nil {
		t.Fatalf("ImportKnownPortableBundle returned error: %v", err)
	}
	if result.SchemaVersion != KnownPortableBundleImportRecordSchemaVersion ||
		result.RequestType != KnownPortableBundleImportRecordRequestType ||
		result.Status != "passed" ||
		result.AppID != app.ID ||
		result.DisplayName != app.DisplayName ||
		result.AppVersion != app.Version ||
		result.ArtifactKind != winapp.KnownPortableArtifactZipBundle ||
		result.DownloadArtifactName != "fixture-bundle.zip" ||
		result.ArtifactCacheRelativePath != "org.xnix.external.fixturebundle/fixture-bundle.zip" ||
		result.FetchStatus != "passed" ||
		result.FetchCacheStatus != "verified" ||
		!result.Downloaded ||
		!result.ChecksumVerified ||
		!result.ArchiveVerified ||
		!result.Extracted ||
		result.ExtractedFileCount != 4 ||
		result.ExtractedDirectoryCount != 0 ||
		!strings.HasPrefix(result.BundleRelativePath, "runtime/known-portable-bundles/org.xnix.external.fixturebundle/") ||
		!result.ImportRecorded ||
		result.ImportRecordRequestType != ExternalWinAppBundleImportRecordRequestType ||
		result.ExternalAppHandle != app.ID ||
		result.ApplicationID != app.ID ||
		!result.BundleManifestSHA256Present ||
		result.ImportedArtifactKind != ExternalWinAppArtifactKindPortableDirectory ||
		result.ImportRecord == nil ||
		result.Fetch == nil ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired ||
		result.RawHostPathExposed ||
		result.RawArchivePathExposed ||
		result.RawBundleRootPathExposed ||
		result.BackendDetailsExposed {
		t.Fatalf("unexpected known portable bundle import result: %#v", result)
	}
	if reasons := ValidateExternalWinAppImportRecord(*result.ImportRecord); len(reasons) != 0 {
		t.Fatalf("ValidateExternalWinAppImportRecord returned reasons: %v", reasons)
	}
	recordPath, err := ExternalWinAppImportRecordPathFromHandle(filepath.Join(tempDir, "state"), app.ID)
	if err != nil {
		t.Fatalf("ExternalWinAppImportRecordPathFromHandle returned error: %v", err)
	}
	if _, err := os.Stat(recordPath); err != nil {
		t.Fatalf("known portable bundle import record was not written: %v", err)
	}
	payload, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(payload), tempDir) {
		t.Fatalf("known portable bundle import output leaked host paths: %s", string(payload))
	}
}

func TestImportKnownPortableBundleRejectsUnsafeZipEntry(t *testing.T) {
	archive := fixtureKnownPortableBundleArchive(t, map[string][]byte{
		"fixture.exe":   {'M', 'Z', 0x90, 0x00},
		"../escape.exe": {'M', 'Z', 0x90, 0x00},
	})
	sum := sha256.Sum256(archive)
	app := fixtureKnownPortableBundleApp(hex.EncodeToString(sum[:]))
	client := &http.Client{Transport: bundleRoundTripFunc(func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(archive)),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})}

	result, err := ImportKnownPortableBundle(context.Background(), KnownPortableBundleImportRequest{
		Version:       "0.2.640-test",
		AppID:         app.ID,
		CatalogApp:    app,
		CacheRoot:     filepath.Join(t.TempDir(), "cache"),
		StateRoot:     filepath.Join(t.TempDir(), "state"),
		AllowDownload: true,
		HTTPClient:    client,
		Timeout:       5 * time.Second,
	})
	if err != nil {
		t.Fatalf("ImportKnownPortableBundle returned error: %v", err)
	}
	if result.Status != "failed" ||
		result.FailureReason == "" ||
		!strings.Contains(result.FailureReason, "escapes extraction root") ||
		result.ImportRecorded {
		t.Fatalf("unsafe zip entry was not rejected correctly: %#v", result)
	}
}

func fixtureKnownPortableBundleApp(sha256 string) winapp.KnownPortableApp {
	return winapp.KnownPortableApp{
		ID:                     "org.xnix.external.fixturebundle",
		DisplayName:            "Fixture Portable Bundle",
		Version:                "1.0.0",
		Architecture:           "windows-x86-gui",
		ExecutableName:         "fixture.exe",
		ArtifactKind:           winapp.KnownPortableArtifactZipBundle,
		DownloadArtifactName:   "fixture-bundle.zip",
		ExecutableRelativePath: "fixture.exe",
		SourcePageURL:          "https://example.invalid/fixture",
		DownloadURL:            "https://example.invalid/fixture-bundle.zip",
		SHA256:                 sha256,
		ExpectedMarker:         "fixture.txt",
		PortableBundleArchive:  true,
	}
}

func fixtureKnownPortableBundleArchive(t *testing.T, files map[string][]byte) []byte {
	t.Helper()
	var buffer bytes.Buffer
	writer := zip.NewWriter(&buffer)
	for name, content := range files {
		entry, err := writer.Create(name)
		if err != nil {
			t.Fatalf("Create zip entry returned error: %v", err)
		}
		if _, err := entry.Write(content); err != nil {
			t.Fatalf("Write zip entry returned error: %v", err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Close zip writer returned error: %v", err)
	}
	return buffer.Bytes()
}

type bundleRoundTripFunc func(*http.Request) (*http.Response, error)

func (fn bundleRoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}
