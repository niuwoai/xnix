package winapp

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	KnownFetchSchemaVersion         = "xnix.runtime.known_windows_app_fetch.v1"
	KnownFetchRequestType           = "windows-known-app-fetch"
	KnownGuestSchemaVersion         = "xnix.runtime.known_windows_app_guest_wine_smoke.v1"
	KnownGuestRequestType           = "windows-known-app-guest-wine-smoke"
	KnownManagedLaunchSchemaVersion = "xnix.runtime.known_windows_app_managed_launch.v1"
	KnownManagedLaunchRequestType   = "windows-known-app-managed-launch-preview"
	KnownKDELauncherSchemaVersion   = "xnix.runtime.known_windows_app_kde_launcher.v1"
	KnownKDELauncherRequestType     = "windows-known-app-kde-launcher-preview"
	DefaultKnownAppID               = "7zr"
	DefaultKnownAppCacheRoot        = ".cache/xnix/known-winapps"
	DefaultKnownAppFetchTimeout     = 60 * time.Second
	DefaultKnownAppGuestTimeout     = 90 * time.Second
	defaultKnownAppDownloadLimit    = 32 << 20
)

type KnownPortableApp struct {
	ID             string
	DisplayName    string
	Version        string
	Architecture   string
	ExecutableName string
	SourcePageURL  string
	DownloadURL    string
	SHA256         string
	ExpectedMarker string
	Arguments      []string
}

type KnownFetchRequest struct {
	AppID         string
	CacheRoot     string
	AllowDownload bool
	HTTPClient    *http.Client
	Timeout       time.Duration
}

type KnownFetchResult struct {
	SchemaVersion          string `json:"schema_version"`
	RequestType            string `json:"request_type"`
	Status                 string `json:"status"`
	AppID                  string `json:"app_id"`
	DisplayName            string `json:"display_name"`
	AppVersion             string `json:"app_version"`
	Architecture           string `json:"architecture"`
	ExecutableName         string `json:"executable_name"`
	SourcePageURL          string `json:"source_page_url"`
	DownloadURL            string `json:"download_url"`
	ExpectedSHA256         string `json:"expected_sha256"`
	ActualSHA256           string `json:"actual_sha256,omitempty"`
	CacheRelativePath      string `json:"cache_relative_path"`
	CacheStatus            string `json:"cache_status"`
	Downloaded             bool   `json:"downloaded"`
	ChecksumVerified       bool   `json:"checksum_verified"`
	NetworkRequired        bool   `json:"network_required"`
	HostRootModified       bool   `json:"host_root_modified"`
	PrivilegedRequired     bool   `json:"privileged_required"`
	HostNetworkingRequired bool   `json:"host_networking_required"`
	DockerSocketMounted    bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired bool   `json:"broad_host_mount_required"`
	RawHostPathExposed     bool   `json:"raw_host_path_exposed"`
	SkipReason             string `json:"skip_reason,omitempty"`
	FailureReason          string `json:"failure_reason,omitempty"`
}

type KnownGuestRequest struct {
	AppID     string
	CacheRoot string
	Arguments []string
	Host      string
	Port      string
	User      string
	KeyPath   string
	RemoteDir string
	SSHPath   string
	SCPPath   string
	Timeout   time.Duration
}

type KnownGuestResult struct {
	SchemaVersion               string      `json:"schema_version"`
	RequestType                 string      `json:"request_type"`
	Status                      string      `json:"status"`
	AppID                       string      `json:"app_id"`
	DisplayName                 string      `json:"display_name"`
	AppVersion                  string      `json:"app_version"`
	Architecture                string      `json:"architecture"`
	ExecutableName              string      `json:"executable_name"`
	ExpectedSHA256              string      `json:"expected_sha256"`
	ActualSHA256                string      `json:"actual_sha256,omitempty"`
	ChecksumVerified            bool        `json:"checksum_verified"`
	ExpectedMarker              string      `json:"expected_marker"`
	Guest                       GuestResult `json:"guest"`
	LoopbackOnlyNetworking      bool        `json:"loopback_only_networking"`
	QEMURequired                bool        `json:"qemu_required"`
	HostRootModified            bool        `json:"host_root_modified"`
	PrivilegedContainerRequired bool        `json:"privileged_container_required"`
	HostNetworkingRequired      bool        `json:"host_networking_required"`
	DockerSocketMounted         bool        `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool        `json:"broad_host_mount_required"`
	RawHostPathExposed          bool        `json:"raw_host_path_exposed"`
	SkipReason                  string      `json:"skip_reason,omitempty"`
	FailureReason               string      `json:"failure_reason,omitempty"`
}

type KnownManagedLaunchRequest struct {
	AppID     string
	CacheRoot string
}

type KnownManagedLaunchResult struct {
	SchemaVersion               string   `json:"schema_version"`
	RequestType                 string   `json:"request_type"`
	Status                      string   `json:"status"`
	AppID                       string   `json:"app_id"`
	DisplayName                 string   `json:"display_name"`
	AppVersion                  string   `json:"app_version"`
	Architecture                string   `json:"architecture"`
	LaunchSurfaceID             string   `json:"launch_surface_id"`
	DesktopActionID             string   `json:"desktop_action_id"`
	DesktopActionLabel          string   `json:"desktop_action_label"`
	ManagedLauncher             string   `json:"managed_launcher"`
	ManagedLauncherArgv         []string `json:"managed_launcher_argv"`
	CacheStatus                 string   `json:"cache_status"`
	ArtifactVerified            bool     `json:"artifact_verified"`
	LaunchEnabled               bool     `json:"launch_enabled"`
	PreparationRequired         bool     `json:"preparation_required"`
	ManagedLaunchSurface        bool     `json:"managed_launch_surface"`
	RuntimeOwnedLaunch          bool     `json:"runtime_owned_launch"`
	KDEPresentationOnly         bool     `json:"kde_presentation_only"`
	RealAppSmokeGateRequired    bool     `json:"real_app_smoke_gate_required"`
	RealAppSmokeGate            string   `json:"real_app_smoke_gate"`
	LoopbackOnlyNetworking      bool     `json:"loopback_only_networking"`
	GuestRuntimeRequired        bool     `json:"guest_runtime_required"`
	HostRootModified            bool     `json:"host_root_modified"`
	PrivilegedContainerRequired bool     `json:"privileged_container_required"`
	HostNetworkingRequired      bool     `json:"host_networking_required"`
	DockerSocketMounted         bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool     `json:"broad_host_mount_required"`
	RawHostPathExposed          bool     `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed           bool     `json:"raw_command_exposed"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	DesktopSafeSummary          string   `json:"desktop_safe_summary"`
	BlockedReason               string   `json:"blocked_reason,omitempty"`
}

type KnownKDELauncherRequest struct {
	AppID     string
	CacheRoot string
}

type KnownKDELauncherResult struct {
	SchemaVersion              string   `json:"schema_version"`
	RequestType                string   `json:"request_type"`
	Source                     string   `json:"source"`
	Status                     string   `json:"status"`
	Desktop                    string   `json:"desktop"`
	EntryPointID               string   `json:"entry_point_id"`
	KDEComponent               string   `json:"kde_component"`
	AppID                      string   `json:"app_id"`
	DisplayName                string   `json:"display_name"`
	AppVersion                 string   `json:"app_version"`
	Architecture               string   `json:"architecture"`
	DesktopFile                string   `json:"desktop_file"`
	Icon                       string   `json:"icon"`
	LaunchSurfaceID            string   `json:"launch_surface_id"`
	DesktopActionID            string   `json:"desktop_action_id"`
	DesktopActionLabel         string   `json:"desktop_action_label"`
	ManagedLauncher            string   `json:"managed_launcher"`
	ManagedLauncherArgv        []string `json:"managed_launcher_argv"`
	CacheStatus                string   `json:"cache_status"`
	ArtifactVerified           bool     `json:"artifact_verified"`
	LaunchVisible              bool     `json:"launch_visible"`
	LaunchEnabled              bool     `json:"launch_enabled"`
	PreparationRequired        bool     `json:"preparation_required"`
	ManagedLaunchSurface       bool     `json:"managed_launch_surface"`
	RuntimeOwnedLaunch         bool     `json:"runtime_owned_launch"`
	KDEPresentationOnly        bool     `json:"kde_presentation_only"`
	DesktopEntryPreviewCreated bool     `json:"desktop_entry_preview_created"`
	DesktopFilesWritten        bool     `json:"desktop_files_written"`
	MIMEAppsWritten            bool     `json:"mimeapps_written"`
	BackendProcessStarted      bool     `json:"backend_process_started"`
	HostRootModified           bool     `json:"host_root_modified"`
	HostNetworkingRequired     bool     `json:"host_networking_required"`
	DockerSocketMounted        bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired     bool     `json:"broad_host_mount_required"`
	RawHostPathExposed         bool     `json:"raw_host_path_exposed"`
	RawExecutablePathExposed   bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed          bool     `json:"raw_command_exposed"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	DesktopSafeSummary         string   `json:"desktop_safe_summary"`
	BlockedReason              string   `json:"blocked_reason,omitempty"`
}

var knownPortableCatalog = []KnownPortableApp{
	{
		ID:             "7zr",
		DisplayName:    "7-Zip standalone console executable",
		Version:        "26.02",
		Architecture:   "windows-x86",
		ExecutableName: "7zr.exe",
		SourcePageURL:  "https://www.7-zip.org/download.html",
		DownloadURL:    "https://github.com/ip7z/7zip/releases/download/26.02/7zr.exe",
		SHA256:         "56b8cc9f4971cef253644fafe54063ed7fdca551d4dee0f8c6baa81b855acd72",
		ExpectedMarker: "7-Zip",
	},
}

func KnownPortableCatalog() []KnownPortableApp {
	result := make([]KnownPortableApp, len(knownPortableCatalog))
	copy(result, knownPortableCatalog)
	return result
}

func LookupKnownPortableApp(appID string) (KnownPortableApp, error) {
	id := strings.TrimSpace(appID)
	if id == "" {
		id = DefaultKnownAppID
	}
	for _, app := range KnownPortableCatalog() {
		if app.ID == id {
			return app, nil
		}
	}
	return KnownPortableApp{}, fmt.Errorf("unknown known Windows app: %s", id)
}

func FetchKnownPortableApp(ctx context.Context, request KnownFetchRequest) (KnownFetchResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownFetchResult{}, err
	}
	result := baseKnownFetchResult(app)

	cachePath, err := knownAppCachePath(request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	result.CacheRelativePath = knownAppCacheRelativePath(app)

	if actual, ok, err := verifyKnownAppFile(cachePath, app); err != nil {
		return result, err
	} else if ok {
		result.Status = PassedStatus
		result.CacheStatus = "verified"
		result.ActualSHA256 = actual
		result.ChecksumVerified = true
		return result, nil
	}

	if !request.AllowDownload {
		result.Status = SkippedStatus
		result.CacheStatus = "missing"
		result.SkipReason = "known Windows app artifact unavailable"
		return result, nil
	}

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = DefaultKnownAppFetchTimeout
	}
	fetchCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := os.MkdirAll(filepath.Dir(cachePath), 0o700); err != nil {
		return result, fmt.Errorf("create known app cache directory: %w", err)
	}
	tempFile, err := os.CreateTemp(filepath.Dir(cachePath), "."+app.ID+"-*.download")
	if err != nil {
		return result, fmt.Errorf("create known app temporary file: %w", err)
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	actual, err := downloadKnownPortableApp(fetchCtx, request.HTTPClient, app.DownloadURL, tempFile)
	closeErr := tempFile.Close()
	if err != nil {
		result.Status = FailedStatus
		result.CacheStatus = "download-failed"
		result.FailureReason = err.Error()
		return result, nil
	}
	if closeErr != nil {
		return result, fmt.Errorf("close known app temporary file: %w", closeErr)
	}
	result.ActualSHA256 = actual
	result.Downloaded = true

	if actual != strings.ToLower(app.SHA256) {
		result.Status = FailedStatus
		result.CacheStatus = "checksum-mismatch"
		result.FailureReason = "known Windows app checksum mismatch"
		return result, nil
	}
	if err := os.Rename(tempPath, cachePath); err != nil {
		return result, fmt.Errorf("store known app artifact: %w", err)
	}

	result.Status = PassedStatus
	result.CacheStatus = "verified"
	result.ChecksumVerified = true
	return result, nil
}

func RunKnownPortableGuestSmoke(ctx context.Context, request KnownGuestRequest) (KnownGuestResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownGuestResult{}, err
	}
	result := baseKnownGuestResult(app)

	cachePath, err := knownAppCachePath(request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	actual, ok, err := verifyKnownAppFile(cachePath, app)
	if err != nil {
		return result, err
	}
	result.ActualSHA256 = actual
	if !ok {
		result.Status = SkippedStatus
		result.SkipReason = "known Windows app artifact unavailable or checksum mismatch"
		return result, nil
	}
	result.ChecksumVerified = true

	timeout := request.Timeout
	if timeout <= 0 {
		timeout = DefaultKnownAppGuestTimeout
	}
	arguments := append([]string{}, app.Arguments...)
	arguments = append(arguments, request.Arguments...)
	guest, err := RunGuestSmoke(ctx, GuestRequest{
		ExecutablePath: cachePath,
		Arguments:      arguments,
		Host:           request.Host,
		Port:           request.Port,
		User:           request.User,
		KeyPath:        request.KeyPath,
		RemoteDir:      request.RemoteDir,
		SSHPath:        request.SSHPath,
		SCPPath:        request.SCPPath,
		Timeout:        timeout,
		ExpectedMarker: app.ExpectedMarker,
	})
	if err != nil {
		return result, err
	}

	result.Guest = guest
	result.Status = guest.Status
	result.LoopbackOnlyNetworking = guest.LoopbackOnlyNetworking
	result.QEMURequired = guest.QEMURequired
	result.HostRootModified = guest.HostRootModified
	result.PrivilegedContainerRequired = guest.PrivilegedContainerRequired
	result.HostNetworkingRequired = guest.HostNetworkingRequired
	result.DockerSocketMounted = guest.DockerSocketMounted
	result.BroadHostMountRequired = guest.BroadHostMountRequired
	result.RawHostPathExposed = guest.RawHostPathExposed
	result.SkipReason = guest.SkipReason
	result.FailureReason = guest.FailureReason
	return result, nil
}

func PreviewKnownPortableManagedLaunch(request KnownManagedLaunchRequest) (KnownManagedLaunchResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownManagedLaunchResult{}, err
	}
	result := baseKnownManagedLaunchResult(app)

	cachePath, err := knownAppCachePath(request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	actual, ok, err := verifyKnownAppFile(cachePath, app)
	if err != nil {
		return result, err
	}
	if ok {
		result.Status = "ready"
		result.CacheStatus = "verified"
		result.ArtifactVerified = true
		result.LaunchEnabled = true
		result.PreparationRequired = false
		result.DesktopSafeSummary = fmt.Sprintf("%s is ready to launch through the managed compatibility runtime.", app.DisplayName)
		result.BlockedReason = ""
		return result, nil
	}
	if actual != "" {
		result.CacheStatus = "checksum-mismatch"
		result.BlockedReason = "managed application artifact failed checksum verification"
	} else {
		result.CacheStatus = "missing"
		result.BlockedReason = "managed application artifact must be fetched before launch"
	}
	return result, nil
}

func PreviewKnownPortableKDELauncher(request KnownKDELauncherRequest) (KnownKDELauncherResult, error) {
	launchSurface, err := PreviewKnownPortableManagedLaunch(KnownManagedLaunchRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownKDELauncherResult{}, err
	}
	result := baseKnownKDELauncherResult(launchSurface)
	if launchSurface.LaunchEnabled {
		result.Status = "visible-ready"
		result.DesktopSafeSummary = launchSurface.DisplayName + " is visible in the KDE launcher and ready for Runtime launch."
	} else {
		result.Status = "visible-needs-preparation"
		result.DesktopSafeSummary = launchSurface.DisplayName + " is visible in the KDE launcher but needs managed artifact preparation before launch."
		result.BlockedReason = launchSurface.BlockedReason
	}
	return result, nil
}

func baseKnownFetchResult(app KnownPortableApp) KnownFetchResult {
	return KnownFetchResult{
		SchemaVersion:          KnownFetchSchemaVersion,
		RequestType:            KnownFetchRequestType,
		Status:                 FailedStatus,
		AppID:                  app.ID,
		DisplayName:            app.DisplayName,
		AppVersion:             app.Version,
		Architecture:           app.Architecture,
		ExecutableName:         app.ExecutableName,
		SourcePageURL:          app.SourcePageURL,
		DownloadURL:            app.DownloadURL,
		ExpectedSHA256:         strings.ToLower(app.SHA256),
		CacheRelativePath:      knownAppCacheRelativePath(app),
		CacheStatus:            "unknown",
		NetworkRequired:        true,
		HostRootModified:       false,
		PrivilegedRequired:     false,
		HostNetworkingRequired: false,
		DockerSocketMounted:    false,
		BroadHostMountRequired: false,
		RawHostPathExposed:     false,
	}
}

func baseKnownKDELauncherResult(launchSurface KnownManagedLaunchResult) KnownKDELauncherResult {
	return KnownKDELauncherResult{
		SchemaVersion:              KnownKDELauncherSchemaVersion,
		RequestType:                KnownKDELauncherRequestType,
		Source:                     KnownManagedLaunchRequestType,
		Status:                     "visible-needs-preparation",
		Desktop:                    "KDE Plasma",
		EntryPointID:               "launcher",
		KDEComponent:               "Plasma application launcher",
		AppID:                      launchSurface.AppID,
		DisplayName:                launchSurface.DisplayName,
		AppVersion:                 launchSurface.AppVersion,
		Architecture:               launchSurface.Architecture,
		DesktopFile:                "xnix-known-app-" + launchSurface.AppID + ".desktop",
		Icon:                       "xnix-known-app-" + launchSurface.AppID,
		LaunchSurfaceID:            launchSurface.LaunchSurfaceID,
		DesktopActionID:            launchSurface.DesktopActionID,
		DesktopActionLabel:         launchSurface.DesktopActionLabel,
		ManagedLauncher:            launchSurface.ManagedLauncher,
		ManagedLauncherArgv:        append([]string{}, launchSurface.ManagedLauncherArgv...),
		CacheStatus:                launchSurface.CacheStatus,
		ArtifactVerified:           launchSurface.ArtifactVerified,
		LaunchVisible:              true,
		LaunchEnabled:              launchSurface.LaunchEnabled,
		PreparationRequired:        launchSurface.PreparationRequired,
		ManagedLaunchSurface:       launchSurface.ManagedLaunchSurface,
		RuntimeOwnedLaunch:         launchSurface.RuntimeOwnedLaunch,
		KDEPresentationOnly:        launchSurface.KDEPresentationOnly,
		DesktopEntryPreviewCreated: true,
		DesktopFilesWritten:        false,
		MIMEAppsWritten:            false,
		BackendProcessStarted:      false,
		HostRootModified:           launchSurface.HostRootModified,
		HostNetworkingRequired:     launchSurface.HostNetworkingRequired,
		DockerSocketMounted:        launchSurface.DockerSocketMounted,
		BroadHostMountRequired:     launchSurface.BroadHostMountRequired,
		RawHostPathExposed:         launchSurface.RawHostPathExposed,
		RawExecutablePathExposed:   launchSurface.RawExecutablePathExposed,
		RawCommandExposed:          launchSurface.RawCommandExposed,
		BackendDetailsExposed:      launchSurface.BackendDetailsExposed,
		DesktopSafeSummary:         launchSurface.DesktopSafeSummary,
		BlockedReason:              launchSurface.BlockedReason,
	}
}

func baseKnownManagedLaunchResult(app KnownPortableApp) KnownManagedLaunchResult {
	return KnownManagedLaunchResult{
		SchemaVersion:               KnownManagedLaunchSchemaVersion,
		RequestType:                 KnownManagedLaunchRequestType,
		Status:                      "needs-artifact",
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		LaunchSurfaceID:             "known-app-" + app.ID,
		DesktopActionID:             "launch-known-app-" + app.ID,
		DesktopActionLabel:          "Launch " + app.DisplayName + " in managed compatibility runtime",
		ManagedLauncher:             "xnix-compat-launch --app " + app.ID,
		ManagedLauncherArgv:         []string{"xnix-compat-launch", "--app", app.ID},
		CacheStatus:                 "unknown",
		ArtifactVerified:            false,
		LaunchEnabled:               false,
		PreparationRequired:         true,
		ManagedLaunchSurface:        true,
		RuntimeOwnedLaunch:          true,
		KDEPresentationOnly:         true,
		RealAppSmokeGateRequired:    true,
		RealAppSmokeGate:            "managed-known-app-guest-smoke",
		LoopbackOnlyNetworking:      true,
		GuestRuntimeRequired:        true,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
		RawExecutablePathExposed:    false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          app.DisplayName + " needs managed artifact preparation before launch.",
		BlockedReason:               "managed application artifact must be fetched before launch",
	}
}

func baseKnownGuestResult(app KnownPortableApp) KnownGuestResult {
	return KnownGuestResult{
		SchemaVersion:               KnownGuestSchemaVersion,
		RequestType:                 KnownGuestRequestType,
		Status:                      FailedStatus,
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		ExpectedSHA256:              strings.ToLower(app.SHA256),
		ExpectedMarker:              app.ExpectedMarker,
		LoopbackOnlyNetworking:      true,
		QEMURequired:                true,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
	}
}

func knownAppCachePath(cacheRoot string, app KnownPortableApp) (string, error) {
	root := strings.TrimSpace(cacheRoot)
	if root == "" {
		root = DefaultKnownAppCacheRoot
	}
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve known app cache root: %w", err)
	}
	return filepath.Join(absoluteRoot, app.ID, app.ExecutableName), nil
}

func knownAppCacheRelativePath(app KnownPortableApp) string {
	return filepath.ToSlash(filepath.Join(app.ID, app.ExecutableName))
}

func verifyKnownAppFile(path string, app KnownPortableApp) (string, bool, error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("inspect known app artifact: %w", err)
	}
	if info.IsDir() {
		return "", false, errors.New("known app artifact path must be a file")
	}
	actual, err := sha256File(path)
	if err != nil {
		return "", false, err
	}
	return actual, actual == strings.ToLower(app.SHA256), nil
}

func sha256File(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open known app artifact: %w", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("hash known app artifact: %w", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func downloadKnownPortableApp(ctx context.Context, client *http.Client, url string, writer io.Writer) (string, error) {
	if client == nil {
		client = http.DefaultClient
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("create known app download request: %w", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("download known app: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download known app returned HTTP %d", response.StatusCode)
	}

	hash := sha256.New()
	limited := io.LimitReader(response.Body, defaultKnownAppDownloadLimit+1)
	written, err := io.Copy(io.MultiWriter(writer, hash), limited)
	if err != nil {
		return "", fmt.Errorf("write known app download: %w", err)
	}
	if written > defaultKnownAppDownloadLimit {
		return "", errors.New("known app download exceeded size limit")
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
