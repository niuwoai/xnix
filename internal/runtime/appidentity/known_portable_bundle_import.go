package appidentity

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownPortableBundleImportRecordSchemaVersion = "xnix.runtime.known_portable_bundle_import_record.v1"
	KnownPortableBundleImportRecordRequestType   = "windows-known-app-bundle-import-record"
	knownPortableBundleExtractLimit              = 256 << 20
	knownPortableBundleEntryLimit                = 4096
)

type KnownPortableBundleImportRequest struct {
	Version        string
	AppID          string
	CatalogApp     winapp.KnownPortableApp
	CacheRoot      string
	StateRoot      string
	AllowDownload  bool
	HTTPClient     *http.Client
	Timeout        time.Duration
	ReportFetch    bool
	RecordImported bool
}

type KnownPortableBundleImportResult struct {
	Version                     string                      `json:"version"`
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	Source                      string                      `json:"source"`
	RuntimeMethod               string                      `json:"runtime_method"`
	Status                      string                      `json:"status"`
	AppID                       string                      `json:"app_id"`
	DisplayName                 string                      `json:"display_name"`
	AppVersion                  string                      `json:"app_version"`
	Architecture                string                      `json:"architecture"`
	ExecutableName              string                      `json:"executable_name"`
	ExecutableRelativePath      string                      `json:"executable_relative_path"`
	ArtifactKind                string                      `json:"artifact_kind"`
	DownloadArtifactName        string                      `json:"download_artifact_name"`
	ArtifactCacheRelativePath   string                      `json:"artifact_cache_relative_path"`
	AllowDownload               bool                        `json:"allow_download"`
	FetchStatus                 string                      `json:"fetch_status,omitempty"`
	FetchCacheStatus            string                      `json:"fetch_cache_status,omitempty"`
	Downloaded                  bool                        `json:"downloaded"`
	ChecksumVerified            bool                        `json:"checksum_verified"`
	PortableBundleArchive       bool                        `json:"portable_bundle_archive"`
	ArchiveVerified             bool                        `json:"archive_verified"`
	Extracted                   bool                        `json:"extracted"`
	ExtractedFileCount          int                         `json:"extracted_file_count"`
	ExtractedDirectoryCount     int                         `json:"extracted_directory_count"`
	BundleRelativePath          string                      `json:"bundle_relative_path,omitempty"`
	ImportRecorded              bool                        `json:"import_recorded"`
	ImportRecordRequestType     string                      `json:"import_record_request_type,omitempty"`
	ExternalAppHandle           string                      `json:"external_app_handle,omitempty"`
	ApplicationID               string                      `json:"application_id,omitempty"`
	BundleManifestSHA256Present bool                        `json:"bundle_manifest_sha256_present"`
	ImportedArtifactKind        string                      `json:"imported_artifact_kind,omitempty"`
	ImportRecord                *ExternalWinAppImportRecord `json:"import_record,omitempty"`
	Fetch                       *winapp.KnownFetchResult    `json:"fetch,omitempty"`
	HostRootModified            bool                        `json:"host_root_modified"`
	NetworkRequired             bool                        `json:"network_required"`
	PrivilegedContainerRequired bool                        `json:"privileged_container_required"`
	HostNetworkingRequired      bool                        `json:"host_networking_required"`
	DockerSocketMounted         bool                        `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                        `json:"broad_host_mount_required"`
	RawHostPathExposed          bool                        `json:"raw_host_path_exposed"`
	RawArchivePathExposed       bool                        `json:"raw_archive_path_exposed"`
	RawBundleRootPathExposed    bool                        `json:"raw_bundle_root_path_exposed"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
	SkipReason                  string                      `json:"skip_reason,omitempty"`
	FailureReason               string                      `json:"failure_reason,omitempty"`
}

func ImportKnownPortableBundle(ctx context.Context, request KnownPortableBundleImportRequest) (KnownPortableBundleImportResult, error) {
	app, err := knownPortableBundleImportApp(request)
	if err != nil {
		return KnownPortableBundleImportResult{}, err
	}
	result := baseKnownPortableBundleImportResult(request, app)
	if !app.PortableBundleArchive || winapp.KnownPortableArtifactZipBundle != result.ArtifactKind {
		result.Status = "skipped"
		result.SkipReason = "known Windows app is not a portable bundle archive"
		result.DesktopSafeSummary = app.DisplayName + " is not a portable bundle archive; use the single-executable known app preparation path."
		return result, nil
	}

	fetch, err := winapp.FetchKnownPortableApp(ctx, winapp.KnownFetchRequest{
		AppID:         app.ID,
		CatalogApp:    request.CatalogApp,
		CacheRoot:     request.CacheRoot,
		AllowDownload: request.AllowDownload,
		HTTPClient:    request.HTTPClient,
		Timeout:       request.Timeout,
	})
	if err != nil {
		return result, err
	}
	result.FetchStatus = fetch.Status
	result.FetchCacheStatus = fetch.CacheStatus
	result.Downloaded = fetch.Downloaded
	result.ChecksumVerified = fetch.ChecksumVerified
	result.ArchiveVerified = fetch.ChecksumVerified
	result.NetworkRequired = fetch.NetworkRequired
	if request.ReportFetch {
		fetchCopy := fetch
		result.Fetch = &fetchCopy
	}
	if fetch.Status != "passed" || !fetch.ChecksumVerified {
		result.Status = fetch.Status
		result.SkipReason = fetch.SkipReason
		result.FailureReason = fetch.FailureReason
		if result.SkipReason == "" && result.FailureReason == "" {
			result.SkipReason = "known portable bundle archive was not verified"
		}
		result.DesktopSafeSummary = app.DisplayName + " portable bundle archive was not imported because the managed artifact is not verified."
		return result, nil
	}

	stateRoot, err := absoluteRequiredPath(request.StateRoot, "known portable bundle import requires --state-root")
	if err != nil {
		return result, err
	}
	archivePath, err := winapp.KnownPortableArtifactCachePath(request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	bundleRelativePath := knownPortableBundleExtractionRelativePath(app, fetch.ActualSHA256)
	bundlePath, err := safeStateRootPath(stateRoot, bundleRelativePath)
	if err != nil {
		return result, err
	}
	if err := os.RemoveAll(bundlePath); err != nil {
		return result, fmt.Errorf("clear known portable bundle extraction directory: %w", err)
	}
	files, directories, err := extractKnownPortableBundleArchive(archivePath, bundlePath)
	if err != nil {
		result.Status = "failed"
		result.FailureReason = err.Error()
		result.DesktopSafeSummary = app.DisplayName + " portable bundle archive could not be extracted safely."
		return result, nil
	}
	result.Extracted = true
	result.ExtractedFileCount = files
	result.ExtractedDirectoryCount = directories
	result.BundleRelativePath = bundleRelativePath

	record, err := RecordExternalWinAppBundleImport(ExternalWinAppBundleImportRequest{
		Version:                strings.TrimSpace(request.Version),
		StateRoot:              stateRoot,
		BundleRoot:             bundlePath,
		ExecutableRelativePath: app.ExecutableRelativePath,
		AppID:                  app.ID,
		DisplayName:            app.DisplayName,
		AppVersion:             app.Version,
	})
	if err != nil {
		return result, err
	}
	result.Status = "passed"
	result.ImportRecorded = true
	result.ImportRecordRequestType = record.RequestType
	result.ExternalAppHandle = record.ApplicationID
	result.ApplicationID = record.ApplicationID
	result.BundleManifestSHA256Present = record.BundleManifestSHA256 != ""
	result.ImportedArtifactKind = record.ArtifactKind
	result.ImportRecord = &record
	result.DesktopSafeSummary = app.DisplayName + " portable bundle was downloaded, verified, safely extracted, and imported into the Runtime-managed external Windows app state root; desktop launch remains gate-controlled."
	return result, nil
}

func knownPortableBundleImportApp(request KnownPortableBundleImportRequest) (winapp.KnownPortableApp, error) {
	if strings.TrimSpace(request.CatalogApp.ID) != "" {
		return request.CatalogApp, nil
	}
	return winapp.LookupKnownPortableApp(request.AppID)
}

func baseKnownPortableBundleImportResult(request KnownPortableBundleImportRequest, app winapp.KnownPortableApp) KnownPortableBundleImportResult {
	artifactKind := app.ArtifactKind
	if artifactKind == "" {
		artifactKind = winapp.KnownPortableArtifactSingleExecutable
	}
	downloadArtifactName := app.DownloadArtifactName
	if strings.TrimSpace(downloadArtifactName) == "" {
		downloadArtifactName = app.ExecutableName
	}
	executableRelativePath := strings.TrimSpace(app.ExecutableRelativePath)
	if executableRelativePath == "" {
		executableRelativePath = app.ExecutableName
	}
	return KnownPortableBundleImportResult{
		Version:                     strings.TrimSpace(request.Version),
		SchemaVersion:               KnownPortableBundleImportRecordSchemaVersion,
		RequestType:                 KnownPortableBundleImportRecordRequestType,
		Source:                      "known-portable-catalog+external-winapp-bundle-import",
		RuntimeMethod:               "ImportKnownPortableBundle",
		Status:                      "skipped",
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		ExecutableRelativePath:      executableRelativePath,
		ArtifactKind:                artifactKind,
		DownloadArtifactName:        downloadArtifactName,
		ArtifactCacheRelativePath:   winapp.KnownPortableArtifactCacheRelativePath(app),
		AllowDownload:               request.AllowDownload,
		PortableBundleArchive:       app.PortableBundleArchive,
		HostRootModified:            false,
		NetworkRequired:             request.AllowDownload,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
		RawArchivePathExposed:       false,
		RawBundleRootPathExposed:    false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          app.DisplayName + " portable bundle import is waiting for a verified managed archive.",
	}
}

func absoluteRequiredPath(value string, emptyMessage string) (string, error) {
	path := strings.TrimSpace(value)
	if path == "" {
		return "", errors.New(emptyMessage)
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve path: %w", err)
	}
	return absolute, nil
}

func knownPortableBundleExtractionRelativePath(app winapp.KnownPortableApp, artifactSHA256 string) string {
	digest := strings.TrimSpace(artifactSHA256)
	if digest == "" {
		digest = "unverified"
	}
	return filepath.ToSlash(filepath.Join("runtime", "known-portable-bundles", stateRootNamespace(app.ID), digest, "bundle"))
}

func extractKnownPortableBundleArchive(archivePath string, targetRoot string) (int, int, error) {
	reader, err := zip.OpenReader(archivePath)
	if err != nil {
		return 0, 0, fmt.Errorf("open known portable bundle archive: %w", err)
	}
	defer reader.Close()
	if len(reader.File) == 0 {
		return 0, 0, errors.New("known portable bundle archive is empty")
	}
	if len(reader.File) > knownPortableBundleEntryLimit {
		return 0, 0, errors.New("known portable bundle archive has too many entries")
	}

	var total uint64
	var files int
	var directories int
	for _, entry := range reader.File {
		relativePath, err := cleanKnownPortableZipEntryPath(entry.Name)
		if err != nil {
			return 0, 0, err
		}
		mode := entry.FileInfo().Mode()
		if mode&os.ModeSymlink != 0 {
			return 0, 0, errors.New("known portable bundle archive symlink entries are not supported")
		}
		targetPath := filepath.Join(targetRoot, filepath.FromSlash(relativePath))
		relativeTarget, err := filepath.Rel(targetRoot, targetPath)
		if err != nil {
			return 0, 0, err
		}
		if relativeTarget == ".." || strings.HasPrefix(filepath.ToSlash(relativeTarget), "../") {
			return 0, 0, errors.New("known portable bundle archive entry escaped extraction root")
		}
		if entry.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, 0o700); err != nil {
				return 0, 0, fmt.Errorf("create known portable bundle directory: %w", err)
			}
			directories++
			continue
		}
		total += entry.UncompressedSize64
		if total > knownPortableBundleExtractLimit {
			return 0, 0, errors.New("known portable bundle archive exceeds extraction size limit")
		}
		if err := os.MkdirAll(filepath.Dir(targetPath), 0o700); err != nil {
			return 0, 0, fmt.Errorf("create known portable bundle parent directory: %w", err)
		}
		if err := extractKnownPortableZipFile(entry, targetPath); err != nil {
			return 0, 0, err
		}
		files++
	}
	if files == 0 {
		return 0, 0, errors.New("known portable bundle archive contains no files")
	}
	return files, directories, nil
}

func extractKnownPortableZipFile(entry *zip.File, targetPath string) error {
	source, err := entry.Open()
	if err != nil {
		return fmt.Errorf("open known portable bundle file: %w", err)
	}
	defer source.Close()
	target, err := os.OpenFile(targetPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("create known portable bundle file: %w", err)
	}
	defer target.Close()
	if _, err := io.Copy(target, io.LimitReader(source, int64(entry.UncompressedSize64)+1)); err != nil {
		return fmt.Errorf("extract known portable bundle file: %w", err)
	}
	return nil
}

func cleanKnownPortableZipEntryPath(value string) (string, error) {
	value = filepath.ToSlash(strings.TrimSpace(value))
	if value == "" || filepath.IsAbs(value) || strings.HasPrefix(value, "/") || strings.Contains(value, "\x00") {
		return "", errors.New("known portable bundle archive entry path is unsafe")
	}
	clean := filepath.ToSlash(filepath.Clean(filepath.FromSlash(value)))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", errors.New("known portable bundle archive entry path escapes extraction root")
	}
	return clean, nil
}
