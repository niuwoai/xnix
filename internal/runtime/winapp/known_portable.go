package winapp

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	KnownFetchSchemaVersion            = "xnix.runtime.known_windows_app_fetch.v1"
	KnownFetchRequestType              = "windows-known-app-fetch"
	KnownGuestSchemaVersion            = "xnix.runtime.known_windows_app_guest_wine_smoke.v1"
	KnownGuestRequestType              = "windows-known-app-guest-wine-smoke"
	KnownManagedLaunchSchemaVersion    = "xnix.runtime.known_windows_app_managed_launch.v1"
	KnownManagedLaunchRequestType      = "windows-known-app-managed-launch-preview"
	KnownKDELauncherSchemaVersion      = "xnix.runtime.known_windows_app_kde_launcher.v1"
	KnownKDELauncherRequestType        = "windows-known-app-kde-launcher-preview"
	KnownLaunchRequestSchemaVersion    = "xnix.runtime.known_windows_app_launch_request.v1"
	KnownLaunchRequestType             = "windows-known-app-launch-request-preview"
	KnownDispatchSchemaVersion         = "xnix.runtime.known_windows_app_dispatch.v1"
	KnownDispatchRequestType           = "windows-known-app-dispatch-preview"
	KnownDispatchSmokeSchemaVersion    = "xnix.runtime.known_windows_app_dispatch_smoke.v1"
	KnownDispatchSmokeRequestType      = "windows-known-app-dispatch-smoke"
	KnownLaunchBridgeSchemaVersion     = "xnix.runtime.known_windows_app_launch_bridge.v1"
	KnownLaunchBridgeRequestType       = "windows-known-app-launch-bridge-preview"
	KnownLaunchProfileSchemaVersion    = "xnix.runtime.known_windows_app_launch_profile_materialize.v1"
	KnownLaunchProfileRequestType      = "windows-known-app-launch-profile-materialize"
	KnownPrepareLaunchSchemaVersion    = "xnix.runtime.known_windows_app_prepare_launch_profile.v1"
	KnownPrepareLaunchRequestType      = "windows-known-app-prepare-launch-profile"
	KnownPrepareAndLaunchSchemaVersion = "xnix.runtime.known_windows_app_prepare_and_launch_profile.v1"
	KnownPrepareAndLaunchRequestType   = "windows-known-app-prepare-and-launch-profile"
	KnownRunSchemaVersion              = "xnix.runtime.known_windows_app_run.v1"
	KnownRunRequestType                = "windows-known-app-run"
	KnownDispatchGuestBoundary         = "managed-known-app-guest-smoke"
	DefaultKnownAppID                  = "7zr"
	DefaultKnownAppCacheRoot           = ".cache/xnix/known-winapps"
	DefaultKnownAppFetchTimeout        = 60 * time.Second
	DefaultKnownAppGuestTimeout        = 90 * time.Second
	KnownAppDownloadUserAgent          = "Xnix-Compatibility-Runtime/known-app-fetch"
	defaultKnownAppDownloadLimit       = 32 << 20
)

const (
	KnownRunBackendLocal     = "local"
	KnownRunBackendGuestWine = "guest-wine"
)

type KnownPortableApp struct {
	ID              string
	DisplayName     string
	Version         string
	Architecture    string
	ExecutableName  string
	SourcePageURL   string
	DownloadURL     string
	SHA256          string
	ExpectedMarker  string
	Arguments       []string
	GuestBuiltinGUI bool
	GuestGUIAppPath string
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
	AppID        string
	CacheRoot    string
	Arguments    []string
	Host         string
	Port         string
	User         string
	KeyPath      string
	RemoteDir    string
	SSHPath      string
	SCPPath      string
	Timeout      time.Duration
	RedactOutput bool
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

type KnownLaunchRequestRequest struct {
	AppID     string
	CacheRoot string
}

type KnownLaunchRequestResult struct {
	SchemaVersion            string   `json:"schema_version"`
	RequestType              string   `json:"request_type"`
	Source                   string   `json:"source"`
	Status                   string   `json:"status"`
	RequestID                string   `json:"request_id"`
	RuntimeMethod            string   `json:"runtime_method"`
	AppID                    string   `json:"app_id"`
	DisplayName              string   `json:"display_name"`
	AppVersion               string   `json:"app_version"`
	Architecture             string   `json:"architecture"`
	Desktop                  string   `json:"desktop"`
	EntryPointID             string   `json:"entry_point_id"`
	DesktopFile              string   `json:"desktop_file"`
	LaunchSurfaceID          string   `json:"launch_surface_id"`
	DesktopActionID          string   `json:"desktop_action_id"`
	ManagedLauncher          string   `json:"managed_launcher"`
	ManagedLauncherArgv      []string `json:"managed_launcher_argv"`
	DispatchGate             string   `json:"dispatch_gate"`
	CacheStatus              string   `json:"cache_status"`
	ArtifactVerified         bool     `json:"artifact_verified"`
	LaunchVisible            bool     `json:"launch_visible"`
	LaunchAllowed            bool     `json:"launch_allowed"`
	LaunchRequestCreated     bool     `json:"launch_request_created"`
	DispatchReady            bool     `json:"dispatch_ready"`
	PreparationRequired      bool     `json:"preparation_required"`
	RuntimeOwnedRequest      bool     `json:"runtime_owned_request"`
	RuntimeOwnedLaunch       bool     `json:"runtime_owned_launch"`
	KDEPresentationOnly      bool     `json:"kde_presentation_only"`
	DryRun                   bool     `json:"dry_run"`
	ExecutionStarted         bool     `json:"execution_started"`
	BackendProcessStarted    bool     `json:"backend_process_started"`
	HostRootModified         bool     `json:"host_root_modified"`
	HostNetworkingRequired   bool     `json:"host_networking_required"`
	DockerSocketMounted      bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired   bool     `json:"broad_host_mount_required"`
	RawHostPathExposed       bool     `json:"raw_host_path_exposed"`
	RawExecutablePathExposed bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed        bool     `json:"raw_command_exposed"`
	BackendDetailsExposed    bool     `json:"backend_details_exposed"`
	DesktopSafeSummary       string   `json:"desktop_safe_summary"`
	BlockedReason            string   `json:"blocked_reason,omitempty"`
}

type KnownDispatchRequest struct {
	AppID     string
	CacheRoot string
}

type KnownDispatchResult struct {
	SchemaVersion            string   `json:"schema_version"`
	RequestType              string   `json:"request_type"`
	Source                   string   `json:"source"`
	Status                   string   `json:"status"`
	RequestID                string   `json:"request_id"`
	DispatchID               string   `json:"dispatch_id"`
	RuntimeMethod            string   `json:"runtime_method"`
	AppID                    string   `json:"app_id"`
	DisplayName              string   `json:"display_name"`
	AppVersion               string   `json:"app_version"`
	Architecture             string   `json:"architecture"`
	Desktop                  string   `json:"desktop"`
	EntryPointID             string   `json:"entry_point_id"`
	DesktopFile              string   `json:"desktop_file"`
	LaunchSurfaceID          string   `json:"launch_surface_id"`
	DesktopActionID          string   `json:"desktop_action_id"`
	ManagedLauncher          string   `json:"managed_launcher"`
	ManagedLauncherArgv      []string `json:"managed_launcher_argv"`
	DispatchGate             string   `json:"dispatch_gate"`
	RunnerLane               string   `json:"runner_lane"`
	RunnerRequestType        string   `json:"runner_request_type"`
	CacheStatus              string   `json:"cache_status"`
	ArtifactVerified         bool     `json:"artifact_verified"`
	LaunchRequestCreated     bool     `json:"launch_request_created"`
	DispatchPreviewCreated   bool     `json:"dispatch_preview_created"`
	DispatchAllowed          bool     `json:"dispatch_allowed"`
	DispatchReady            bool     `json:"dispatch_ready"`
	PreparationRequired      bool     `json:"preparation_required"`
	RuntimeOwnedRequest      bool     `json:"runtime_owned_request"`
	RuntimeOwnedLaunch       bool     `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch     bool     `json:"runtime_owned_dispatch"`
	KDEPresentationOnly      bool     `json:"kde_presentation_only"`
	DryRun                   bool     `json:"dry_run"`
	DispatchStarted          bool     `json:"dispatch_started"`
	ExecutionStarted         bool     `json:"execution_started"`
	BackendProcessStarted    bool     `json:"backend_process_started"`
	HostRootModified         bool     `json:"host_root_modified"`
	HostNetworkingRequired   bool     `json:"host_networking_required"`
	DockerSocketMounted      bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired   bool     `json:"broad_host_mount_required"`
	RawHostPathExposed       bool     `json:"raw_host_path_exposed"`
	RawExecutablePathExposed bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed        bool     `json:"raw_command_exposed"`
	BackendDetailsExposed    bool     `json:"backend_details_exposed"`
	DesktopSafeSummary       string   `json:"desktop_safe_summary"`
	BlockedReason            string   `json:"blocked_reason,omitempty"`
}

type KnownDispatchSmokeRequest struct {
	AppID          string
	CacheRoot      string
	Arguments      []string
	GuestBoundary  string
	Host           string
	Port           string
	User           string
	KeyPath        string
	RemoteDir      string
	SSHPath        string
	SCPPath        string
	XWinInfoPath   string
	ExecutablePath string
	GUIAppPath     string
	GuestDisplay   string
	HostDisplay    string
	Timeout        time.Duration
	Wait           time.Duration
}

type KnownDispatchSmokeResult struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	Source                      string `json:"source"`
	Status                      string `json:"status"`
	RequestID                   string `json:"request_id"`
	DispatchID                  string `json:"dispatch_id"`
	RuntimeMethod               string `json:"runtime_method"`
	AppID                       string `json:"app_id"`
	DisplayName                 string `json:"display_name"`
	AppVersion                  string `json:"app_version"`
	Architecture                string `json:"architecture"`
	DispatchGate                string `json:"dispatch_gate"`
	RunnerLane                  string `json:"runner_lane"`
	GuestBoundary               string `json:"guest_boundary"`
	CacheStatus                 string `json:"cache_status"`
	ArtifactVerified            bool   `json:"artifact_verified"`
	LaunchRequestCreated        bool   `json:"launch_request_created"`
	DispatchPreviewCreated      bool   `json:"dispatch_preview_created"`
	DispatchReady               bool   `json:"dispatch_ready"`
	DispatchAllowed             bool   `json:"dispatch_allowed"`
	DispatchStarted             bool   `json:"dispatch_started"`
	ExecutionStarted            bool   `json:"execution_started"`
	BackendProcessStarted       bool   `json:"backend_process_started"`
	ManagedGuestRunnerInvoked   bool   `json:"managed_guest_runner_invoked"`
	ManagedGuestReachable       bool   `json:"managed_guest_reachable"`
	ManagedGuestRuntimeReady    bool   `json:"managed_guest_runtime_ready"`
	ManagedArtifactCopied       bool   `json:"managed_artifact_copied"`
	MarkerObserved              bool   `json:"marker_observed"`
	SmokePassed                 bool   `json:"smoke_passed"`
	ExitCode                    int    `json:"exit_code"`
	DurationMillis              int64  `json:"duration_millis"`
	RuntimeOwnedRequest         bool   `json:"runtime_owned_request"`
	RuntimeOwnedLaunch          bool   `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch        bool   `json:"runtime_owned_dispatch"`
	KDEPresentationOnly         bool   `json:"kde_presentation_only"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	HostNetworkingRequired      bool   `json:"host_networking_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	RawHostPathExposed          bool   `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool   `json:"raw_executable_path_exposed"`
	RawCommandExposed           bool   `json:"raw_command_exposed"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
	SkipReason                  string `json:"skip_reason,omitempty"`
	FailureReason               string `json:"failure_reason,omitempty"`
	BlockedReason               string `json:"blocked_reason,omitempty"`
}

type KnownLaunchBridgeRequest struct {
	AppID               string
	CacheRoot           string
	ManagedLauncherArgv []string
}

type KnownLaunchProfileMaterializeRequest struct {
	AppID            string
	CacheRoot        string
	StateRoot        string
	ProfileOutput    string
	ApplicationID    string
	DisplayName      string
	RuntimeBinary    string
	RuntimeArguments []string
	RunnerPath       string
	RunnerBottle     string
	RunnerArguments  []string
	SkipBootstrap    bool
}

type KnownLaunchProfileMaterializeResult struct {
	SchemaVersion               string                `json:"schema_version"`
	RequestType                 string                `json:"request_type"`
	Status                      string                `json:"status"`
	AppID                       string                `json:"app_id"`
	DisplayName                 string                `json:"display_name"`
	AppVersion                  string                `json:"app_version"`
	Architecture                string                `json:"architecture"`
	ExecutableName              string                `json:"executable_name"`
	ExpectedSHA256              string                `json:"expected_sha256"`
	ActualSHA256                string                `json:"actual_sha256,omitempty"`
	CacheStatus                 string                `json:"cache_status"`
	ArtifactVerified            bool                  `json:"artifact_verified"`
	ProfileWritten              bool                  `json:"profile_written"`
	ProfileFileName             string                `json:"profile_file_name"`
	LauncherBundleWritten       bool                  `json:"launcher_bundle_written"`
	LauncherMode                string                `json:"launcher_mode"`
	LauncherCommand             string                `json:"launcher_command"`
	LauncherBundlePayload       *LauncherBundleRecord `json:"launcher_bundle_payload,omitempty"`
	RunnerConfigured            bool                  `json:"runner_configured"`
	RunnerBottleConfigured      bool                  `json:"runner_bottle_configured"`
	RunnerArgumentCount         int                   `json:"runner_argument_count"`
	SkipBootstrap               bool                  `json:"skip_bootstrap"`
	ExpectedMarker              string                `json:"expected_marker"`
	SuccessMode                 string                `json:"success_mode"`
	ApplicationWorkspaceMode    string                `json:"application_workspace_mode"`
	RawHostPathExposed          bool                  `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool                  `json:"raw_executable_path_exposed"`
	RawProfilePathExposed       bool                  `json:"raw_profile_path_exposed"`
	RawStateRootPathExposed     bool                  `json:"raw_state_root_path_exposed"`
	RawRuntimeArgvExposed       bool                  `json:"raw_runtime_argv_exposed"`
	NetworkRequired             bool                  `json:"network_required"`
	HostRootModified            bool                  `json:"host_root_modified"`
	PrivilegedContainerRequired bool                  `json:"privileged_container_required"`
	HostNetworkingRequired      bool                  `json:"host_networking_required"`
	DockerSocketMounted         bool                  `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                  `json:"broad_host_mount_required"`
	SkipReason                  string                `json:"skip_reason,omitempty"`
	FailureReason               string                `json:"failure_reason,omitempty"`
}

// Raw profile fields runner_path, runner_bottle, runner_arguments, and skip_bootstrap are stored only in the operator-local smoke profile.
// Product-facing known-app materialization reports expose only runner_configured, runner_bottle_configured, runner_argument_count, and skip_bootstrap.
type KnownPrepareLaunchProfileRequest struct {
	AppID            string
	CacheRoot        string
	StateRoot        string
	ProfileOutput    string
	ApplicationID    string
	DisplayName      string
	RuntimeBinary    string
	RuntimeArguments []string
	RunnerPath       string
	RunnerBottle     string
	RunnerArguments  []string
	SkipBootstrap    bool
	AllowDownload    bool
	HTTPClient       *http.Client
	Timeout          time.Duration
}

type KnownPrepareLaunchProfileResult struct {
	SchemaVersion               string                               `json:"schema_version"`
	RequestType                 string                               `json:"request_type"`
	Status                      string                               `json:"status"`
	AppID                       string                               `json:"app_id"`
	DisplayName                 string                               `json:"display_name"`
	AppVersion                  string                               `json:"app_version"`
	Architecture                string                               `json:"architecture"`
	ExecutableName              string                               `json:"executable_name"`
	AllowDownload               bool                                 `json:"allow_download"`
	FetchStatus                 string                               `json:"fetch_status"`
	FetchCacheStatus            string                               `json:"fetch_cache_status"`
	Downloaded                  bool                                 `json:"downloaded"`
	ChecksumVerified            bool                                 `json:"checksum_verified"`
	MaterializeStatus           string                               `json:"materialize_status"`
	ProfileWritten              bool                                 `json:"profile_written"`
	LauncherBundleWritten       bool                                 `json:"launcher_bundle_written"`
	FetchPayload                KnownFetchResult                     `json:"fetch_payload"`
	MaterializePayload          *KnownLaunchProfileMaterializeResult `json:"materialize_payload,omitempty"`
	RunnerConfigured            bool                                 `json:"runner_configured"`
	RunnerBottleConfigured      bool                                 `json:"runner_bottle_configured"`
	RunnerArgumentCount         int                                  `json:"runner_argument_count"`
	SkipBootstrap               bool                                 `json:"skip_bootstrap"`
	NetworkRequired             bool                                 `json:"network_required"`
	HostRootModified            bool                                 `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                 `json:"privileged_container_required"`
	HostNetworkingRequired      bool                                 `json:"host_networking_required"`
	DockerSocketMounted         bool                                 `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                                 `json:"broad_host_mount_required"`
	RawHostPathExposed          bool                                 `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool                                 `json:"raw_executable_path_exposed"`
	RawProfilePathExposed       bool                                 `json:"raw_profile_path_exposed"`
	RawStateRootPathExposed     bool                                 `json:"raw_state_root_path_exposed"`
	RawRuntimeArgvExposed       bool                                 `json:"raw_runtime_argv_exposed"`
	SkipReason                  string                               `json:"skip_reason,omitempty"`
	FailureReason               string                               `json:"failure_reason,omitempty"`
}

type KnownPrepareAndLaunchProfileRequest struct {
	AppID            string
	CacheRoot        string
	StateRoot        string
	ProfileOutput    string
	ApplicationID    string
	DisplayName      string
	RuntimeBinary    string
	RuntimeArguments []string
	RunnerPath       string
	RunnerBottle     string
	RunnerArguments  []string
	SkipBootstrap    bool
	AllowDownload    bool
	HTTPClient       *http.Client
	Timeout          time.Duration
}

type KnownPrepareAndLaunchProfileResult struct {
	SchemaVersion               string                          `json:"schema_version"`
	RequestType                 string                          `json:"request_type"`
	Status                      string                          `json:"status"`
	AppID                       string                          `json:"app_id"`
	DisplayName                 string                          `json:"display_name"`
	AppVersion                  string                          `json:"app_version"`
	Architecture                string                          `json:"architecture"`
	ExecutableName              string                          `json:"executable_name"`
	AllowDownload               bool                            `json:"allow_download"`
	PrepareStatus               string                          `json:"prepare_status"`
	LaunchStatus                string                          `json:"launch_status"`
	ProfileWritten              bool                            `json:"profile_written"`
	LauncherBundleWritten       bool                            `json:"launcher_bundle_written"`
	LaunchAttempted             bool                            `json:"launch_attempted"`
	RunnerConfigured            bool                            `json:"runner_configured"`
	RunnerBottleConfigured      bool                            `json:"runner_bottle_configured"`
	RunnerArgumentCount         int                             `json:"runner_argument_count"`
	RunnerAvailable             bool                            `json:"runner_available"`
	SkipBootstrap               bool                            `json:"skip_bootstrap"`
	ExecutableFormat            string                          `json:"executable_format"`
	WindowsExecutableSignature  bool                            `json:"windows_executable_signature_observed"`
	ExecutableArchitecture      string                          `json:"executable_architecture"`
	ExecutableArchitectureReady bool                            `json:"executable_architecture_supported"`
	WineArchitecture            string                          `json:"wine_architecture"`
	ApplicationWorkspaceMode    string                          `json:"application_workspace_mode"`
	RawOutputRedacted           bool                            `json:"raw_output_redacted"`
	PreparePayload              KnownPrepareLaunchProfileResult `json:"prepare_payload"`
	LaunchPayload               *LaunchProfileResult            `json:"launch_payload,omitempty"`
	NetworkRequired             bool                            `json:"network_required"`
	HostRootModified            bool                            `json:"host_root_modified"`
	PrivilegedContainerRequired bool                            `json:"privileged_container_required"`
	HostNetworkingRequired      bool                            `json:"host_networking_required"`
	DockerSocketMounted         bool                            `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                            `json:"broad_host_mount_required"`
	DockerExecuted              bool                            `json:"docker_executed"`
	QEMUExecuted                bool                            `json:"qemu_executed"`
	WineExecuted                bool                            `json:"wine_executed"`
	ColimaExecuted              bool                            `json:"colima_executed"`
	NetworkChecksRun            bool                            `json:"network_checks_run"`
	PackageManagerInvoked       bool                            `json:"package_manager_invoked"`
	RawHostPathExposed          bool                            `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool                            `json:"raw_executable_path_exposed"`
	RawProfilePathExposed       bool                            `json:"raw_profile_path_exposed"`
	RawStateRootPathExposed     bool                            `json:"raw_state_root_path_exposed"`
	RawRuntimeArgvExposed       bool                            `json:"raw_runtime_argv_exposed"`
	RawRunnerPathExposed        bool                            `json:"raw_runner_path_exposed"`
	NextAction                  string                          `json:"next_action"`
	SkipReason                  string                          `json:"skip_reason,omitempty"`
	FailureReason               string                          `json:"failure_reason,omitempty"`
}

type KnownRunRequest struct {
	AppID            string
	Backend          string
	CacheRoot        string
	StateRoot        string
	ProfileOutput    string
	ApplicationID    string
	DisplayName      string
	RuntimeBinary    string
	RuntimeArguments []string
	RunnerPath       string
	RunnerBottle     string
	RunnerArguments  []string
	SkipBootstrap    bool
	AllowDownload    bool
	Arguments        []string
	Host             string
	Port             string
	User             string
	KeyPath          string
	RemoteDir        string
	SSHPath          string
	SCPPath          string
	Timeout          time.Duration
	RedactOutput     bool
	StartQEMU        bool
	QEMUBinary       string
	QEMUKernelImage  string
	QEMUMemory       string
	QEMUCPUCount     string
	QEMUCPUModel     string
	QEMUBootTimeout  time.Duration
	QEMUSerialLog    string
	HTTPClient       *http.Client
}

type KnownRunResult struct {
	SchemaVersion               string                              `json:"schema_version"`
	RequestType                 string                              `json:"request_type"`
	Status                      string                              `json:"status"`
	AppID                       string                              `json:"app_id"`
	DisplayName                 string                              `json:"display_name"`
	AppVersion                  string                              `json:"app_version"`
	Architecture                string                              `json:"architecture"`
	ExecutableName              string                              `json:"executable_name"`
	Backend                     string                              `json:"backend"`
	BackendReady                bool                                `json:"backend_ready"`
	LaunchAttempted             bool                                `json:"launch_attempted"`
	RunnerAvailable             bool                                `json:"runner_available"`
	ChecksumVerified            bool                                `json:"checksum_verified"`
	ProfileWritten              bool                                `json:"profile_written"`
	LauncherBundleWritten       bool                                `json:"launcher_bundle_written"`
	GuestStartAttempted         bool                                `json:"guest_start_attempted"`
	GuestStarted                bool                                `json:"guest_started"`
	GuestStartMode              string                              `json:"guest_start_mode"`
	GuestHost                   string                              `json:"guest_host,omitempty"`
	GuestPort                   string                              `json:"guest_port,omitempty"`
	GuestPortAuto               bool                                `json:"guest_port_auto"`
	QEMUSerialLogWritten        bool                                `json:"qemu_serial_log_written"`
	ExecutableCopied            bool                                `json:"executable_copied"`
	MarkerObserved              bool                                `json:"marker_observed"`
	ExecutableFormat            string                              `json:"executable_format"`
	ExecutableArchitecture      string                              `json:"executable_architecture"`
	WineArchitecture            string                              `json:"wine_architecture"`
	ApplicationWorkspaceMode    string                              `json:"application_workspace_mode"`
	RawOutputRedacted           bool                                `json:"raw_output_redacted"`
	LocalPayload                *KnownPrepareAndLaunchProfileResult `json:"local_payload,omitempty"`
	GuestPayload                *KnownGuestResult                   `json:"guest_payload,omitempty"`
	LoopbackOnlyNetworking      bool                                `json:"loopback_only_networking"`
	QEMURequired                bool                                `json:"qemu_required"`
	NetworkRequired             bool                                `json:"network_required"`
	HostRootModified            bool                                `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                `json:"privileged_container_required"`
	HostNetworkingRequired      bool                                `json:"host_networking_required"`
	DockerSocketMounted         bool                                `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                                `json:"broad_host_mount_required"`
	DockerExecuted              bool                                `json:"docker_executed"`
	QEMUExecuted                bool                                `json:"qemu_executed"`
	WineExecuted                bool                                `json:"wine_executed"`
	ColimaExecuted              bool                                `json:"colima_executed"`
	NetworkChecksRun            bool                                `json:"network_checks_run"`
	PackageManagerInvoked       bool                                `json:"package_manager_invoked"`
	RawHostPathExposed          bool                                `json:"raw_host_path_exposed"`
	RawExecutablePathExposed    bool                                `json:"raw_executable_path_exposed"`
	RawProfilePathExposed       bool                                `json:"raw_profile_path_exposed"`
	RawStateRootPathExposed     bool                                `json:"raw_state_root_path_exposed"`
	RawRuntimeArgvExposed       bool                                `json:"raw_runtime_argv_exposed"`
	RawRunnerPathExposed        bool                                `json:"raw_runner_path_exposed"`
	RawQEMUPathExposed          bool                                `json:"raw_qemu_path_exposed"`
	NextAction                  string                              `json:"next_action"`
	SkipReason                  string                              `json:"skip_reason,omitempty"`
	FailureReason               string                              `json:"failure_reason,omitempty"`
}

type KnownLaunchBridgeResult struct {
	SchemaVersion                    string   `json:"schema_version"`
	RequestType                      string   `json:"request_type"`
	Source                           string   `json:"source"`
	Status                           string   `json:"status"`
	Desktop                          string   `json:"desktop"`
	EntryPointID                     string   `json:"entry_point_id"`
	DesktopFile                      string   `json:"desktop_file"`
	LaunchSurfaceID                  string   `json:"launch_surface_id"`
	DesktopActionID                  string   `json:"desktop_action_id"`
	ManagedLauncher                  string   `json:"managed_launcher"`
	ManagedLauncherArgv              []string `json:"managed_launcher_argv"`
	LauncherArgvAccepted             bool     `json:"launcher_argv_accepted"`
	RequestID                        string   `json:"request_id"`
	DispatchID                       string   `json:"dispatch_id"`
	RuntimeMethod                    string   `json:"runtime_method"`
	DispatchSmokeRequestType         string   `json:"dispatch_smoke_request_type"`
	DispatchSmokeRequestMaterialized bool     `json:"dispatch_smoke_request_materialized"`
	AppID                            string   `json:"app_id"`
	DisplayName                      string   `json:"display_name"`
	AppVersion                       string   `json:"app_version"`
	Architecture                     string   `json:"architecture"`
	DispatchGate                     string   `json:"dispatch_gate"`
	RunnerLane                       string   `json:"runner_lane"`
	GuestBoundaryRequired            bool     `json:"guest_boundary_required"`
	GuestBoundarySupplied            bool     `json:"guest_boundary_supplied"`
	SmokeHarnessRequired             bool     `json:"smoke_harness_required"`
	CacheStatus                      string   `json:"cache_status"`
	ArtifactVerified                 bool     `json:"artifact_verified"`
	LaunchRequestCreated             bool     `json:"launch_request_created"`
	DispatchPreviewCreated           bool     `json:"dispatch_preview_created"`
	BridgePreviewCreated             bool     `json:"bridge_preview_created"`
	DispatchReady                    bool     `json:"dispatch_ready"`
	PreparationRequired              bool     `json:"preparation_required"`
	RuntimeOwnedRequest              bool     `json:"runtime_owned_request"`
	RuntimeOwnedLaunch               bool     `json:"runtime_owned_launch"`
	RuntimeOwnedDispatch             bool     `json:"runtime_owned_dispatch"`
	RuntimeOwnedBridge               bool     `json:"runtime_owned_bridge"`
	KDEPresentationOnly              bool     `json:"kde_presentation_only"`
	DryRun                           bool     `json:"dry_run"`
	DispatchStarted                  bool     `json:"dispatch_started"`
	ExecutionStarted                 bool     `json:"execution_started"`
	BackendProcessStarted            bool     `json:"backend_process_started"`
	HostRootModified                 bool     `json:"host_root_modified"`
	HostNetworkingRequired           bool     `json:"host_networking_required"`
	DockerSocketMounted              bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool     `json:"broad_host_mount_required"`
	RawHostPathExposed               bool     `json:"raw_host_path_exposed"`
	RawExecutablePathExposed         bool     `json:"raw_executable_path_exposed"`
	RawCommandExposed                bool     `json:"raw_command_exposed"`
	BackendDetailsExposed            bool     `json:"backend_details_exposed"`
	DesktopSafeSummary               string   `json:"desktop_safe_summary"`
	BlockedReason                    string   `json:"blocked_reason,omitempty"`
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
	{
		ID:             "busybox-w32",
		DisplayName:    "BusyBox-w32 standalone console executable",
		Version:        "current-2026-07-24",
		Architecture:   "windows-x86",
		ExecutableName: "busybox.exe",
		SourcePageURL:  "https://frippery.org/busybox/",
		DownloadURL:    "https://frippery.org/files/busybox/busybox.exe",
		SHA256:         "7bfee530965315665044e6e01db58125f2763c8a39c2e72ba1a6beb6923e0e1f",
		ExpectedMarker: "BusyBox",
		Arguments:      []string{"--help"},
	},
	{
		ID:              "org.xnix.apps.mines",
		DisplayName:     "Mines",
		Version:         "0.2.640-rc107",
		Architecture:    "windows-x86-gui",
		ExecutableName:  "winemine.exe",
		SourcePageURL:   "runtime-managed-guest-gui-fixture",
		ExpectedMarker:  "Mines",
		GuestBuiltinGUI: true,
		GuestGUIAppPath: DefaultGuestGUIApp,
	},
	{
		ID:              "org.xnix.apps.messagebox",
		DisplayName:     "Xnix MessageBox",
		Version:         "0.2.640-rc107",
		Architecture:    "windows-x86-gui",
		ExecutableName:  "xnix-messagebox-smoke.exe",
		SourcePageURL:   "runtime-managed-external-gui-fixture",
		ExpectedMarker:  "MessageBox",
		GuestBuiltinGUI: true,
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
	if app.GuestBuiltinGUI {
		result.Status = PassedStatus
		result.CacheStatus = "guest-builtin-gui"
		result.SkipReason = ""
		result.Downloaded = false
		result.ChecksumVerified = false
		result.NetworkRequired = false
		result.HostRootModified = false
		result.HostNetworkingRequired = false
		return result, nil
	}

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
	if app.GuestBuiltinGUI {
		result.Status = SkippedStatus
		result.SkipReason = "known GUI application requires the Runtime guest GUI smoke lane"
		return result, nil
	}

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
		RedactOutput:   request.RedactOutput,
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
	if app.GuestBuiltinGUI {
		result.Status = "ready"
		result.CacheStatus = "guest-builtin-gui"
		result.ArtifactVerified = true
		result.LaunchEnabled = true
		result.PreparationRequired = false
		result.DesktopSafeSummary = app.DisplayName + " is ready for Runtime-managed GUI launch using recorded guest GUI evidence instead of a host-cached download artifact."
		result.BlockedReason = ""
		return result, nil
	}

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

func PreviewKnownPortableLaunchRequest(request KnownLaunchRequestRequest) (KnownLaunchRequestResult, error) {
	launcher, err := PreviewKnownPortableKDELauncher(KnownKDELauncherRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownLaunchRequestResult{}, err
	}
	result := baseKnownLaunchRequestResult(launcher)
	if launcher.LaunchEnabled {
		result.Status = "request-ready"
		result.LaunchAllowed = true
		result.DispatchReady = true
		result.PreparationRequired = false
		result.BlockedReason = ""
		result.DesktopSafeSummary = launcher.DisplayName + " has a Runtime-owned launch request ready for managed dispatch."
	} else {
		result.Status = "request-blocked"
		result.BlockedReason = launcher.BlockedReason
		result.DesktopSafeSummary = launcher.DisplayName + " has a Runtime-owned launch request, but managed artifact preparation is required before dispatch."
	}
	return result, nil
}

func PreviewKnownPortableDispatch(request KnownDispatchRequest) (KnownDispatchResult, error) {
	launchRequest, err := PreviewKnownPortableLaunchRequest(KnownLaunchRequestRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownDispatchResult{}, err
	}
	result := baseKnownDispatchResult(launchRequest)
	if launchRequest.DispatchReady {
		result.Status = "dispatch-ready"
		result.DispatchAllowed = true
		result.DispatchReady = true
		result.PreparationRequired = false
		result.BlockedReason = ""
		result.DesktopSafeSummary = launchRequest.DisplayName + " is ready for Runtime-owned managed guest dispatch."
	} else {
		result.Status = "dispatch-blocked"
		result.BlockedReason = launchRequest.BlockedReason
		result.DesktopSafeSummary = launchRequest.DisplayName + " has a Runtime-owned dispatch preview, but managed artifact preparation is required before dispatch."
	}
	return result, nil
}

func RunKnownPortableDispatchSmoke(ctx context.Context, request KnownDispatchSmokeRequest) (KnownDispatchSmokeResult, error) {
	dispatchPreview, err := PreviewKnownPortableDispatch(KnownDispatchRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownDispatchSmokeResult{}, err
	}
	result := baseKnownDispatchSmokeResult(dispatchPreview, request.GuestBoundary)
	if !dispatchPreview.DispatchReady {
		result.Status = "dispatch-blocked"
		result.BlockedReason = dispatchPreview.BlockedReason
		result.SkipReason = dispatchPreview.BlockedReason
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " dispatch is blocked until managed artifact preparation completes."
		return result, nil
	}
	if strings.TrimSpace(request.GuestBoundary) != KnownDispatchGuestBoundary {
		result.Status = "dispatch-blocked"
		result.BlockedReason = "controlled managed guest boundary must be supplied by the smoke harness"
		result.SkipReason = result.BlockedReason
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " dispatch is ready, but the smoke harness did not supply the controlled guest boundary."
		return result, nil
	}

	guest, err := RunKnownPortableGuestSmoke(ctx, KnownGuestRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
		Arguments: append([]string{}, request.Arguments...),
		Host:      request.Host,
		Port:      request.Port,
		User:      request.User,
		KeyPath:   request.KeyPath,
		RemoteDir: request.RemoteDir,
		SSHPath:   request.SSHPath,
		SCPPath:   request.SCPPath,
		Timeout:   request.Timeout,
	})
	if err != nil {
		return result, err
	}

	result.Status = guest.Status
	result.DispatchAllowed = true
	result.DispatchStarted = true
	result.ExecutionStarted = guest.Guest.ExecutableCopied && guest.Guest.WineAvailable
	result.ManagedGuestRunnerInvoked = true
	result.ManagedGuestReachable = guest.Guest.GuestReachable
	result.ManagedGuestRuntimeReady = guest.Guest.WineAvailable
	result.ManagedArtifactCopied = guest.Guest.ExecutableCopied
	result.MarkerObserved = guest.Guest.MarkerObserved
	result.SmokePassed = guest.Status == PassedStatus
	result.ExitCode = guest.Guest.ExitCode
	result.DurationMillis = guest.Guest.DurationMillis
	result.PrivilegedContainerRequired = guest.PrivilegedContainerRequired
	result.HostRootModified = guest.HostRootModified
	result.HostNetworkingRequired = guest.HostNetworkingRequired
	result.DockerSocketMounted = guest.DockerSocketMounted
	result.BroadHostMountRequired = guest.BroadHostMountRequired
	result.RawHostPathExposed = guest.RawHostPathExposed
	result.SkipReason = guest.SkipReason
	result.FailureReason = guest.FailureReason
	if result.SmokePassed {
		result.DesktopSafeSummary = guest.DisplayName + " passed the gated Runtime managed guest dispatch smoke."
	} else {
		result.DesktopSafeSummary = guest.DisplayName + " did not pass the gated Runtime managed guest dispatch smoke."
	}
	return result, nil
}

func RunKnownPortableGuestGUIDispatchSmoke(ctx context.Context, request KnownDispatchSmokeRequest) (KnownDispatchSmokeResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownDispatchSmokeResult{}, err
	}
	dispatchPreview, err := PreviewKnownPortableDispatch(KnownDispatchRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownDispatchSmokeResult{}, err
	}
	result := baseKnownDispatchSmokeResult(dispatchPreview, request.GuestBoundary)
	result.CacheStatus = "guest-builtin-gui"
	result.ArtifactVerified = false
	result.MarkerObserved = false
	if !app.GuestBuiltinGUI {
		result.Status = "dispatch-blocked"
		result.BlockedReason = "known Windows app is not a Runtime-managed guest GUI application"
		result.SkipReason = result.BlockedReason
		return result, nil
	}
	if !dispatchPreview.DispatchReady {
		result.Status = "dispatch-blocked"
		result.BlockedReason = dispatchPreview.BlockedReason
		result.SkipReason = dispatchPreview.BlockedReason
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " GUI dispatch is blocked until Runtime preparation completes."
		return result, nil
	}
	if strings.TrimSpace(request.GuestBoundary) != KnownDispatchGuestBoundary {
		result.Status = "dispatch-blocked"
		result.BlockedReason = "controlled managed guest boundary must be supplied by the GUI smoke harness"
		result.SkipReason = result.BlockedReason
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " GUI dispatch is ready, but the smoke harness did not supply the controlled guest boundary."
		return result, nil
	}

	gui, err := RunGuestGUISmoke(ctx, GuestGUIRequest{
		ExecutablePath:  request.ExecutablePath,
		GUIAppPath:      knownPortableGuestGUIAppPath(request.GUIAppPath),
		KnownAppID:      app.ID,
		KnownAppName:    app.DisplayName,
		KnownAppVersion: app.Version,
		Host:            request.Host,
		Port:            request.Port,
		User:            request.User,
		KeyPath:         request.KeyPath,
		RemoteDir:       request.RemoteDir,
		SSHPath:         request.SSHPath,
		SCPPath:         request.SCPPath,
		XWinInfoPath:    request.XWinInfoPath,
		GuestDisplay:    request.GuestDisplay,
		HostDisplay:     request.HostDisplay,
		Timeout:         request.Timeout,
		Wait:            request.Wait,
	})
	if err != nil {
		return result, err
	}

	result.Status = gui.Status
	result.DispatchAllowed = true
	result.DispatchStarted = true
	result.ExecutionStarted = gui.LaunchAttempted
	result.ManagedGuestRunnerInvoked = true
	result.ManagedGuestReachable = gui.GuestReachable
	result.ManagedGuestRuntimeReady = gui.WineAvailable && gui.GuestX11DriverAvailable
	result.ManagedArtifactCopied = gui.ExecutableCopied
	result.SmokePassed = gui.Status == PassedStatus
	result.DurationMillis = gui.DurationMillis
	result.PrivilegedContainerRequired = gui.PrivilegedContainerRequired
	result.HostRootModified = gui.HostRootModified
	result.HostNetworkingRequired = gui.HostNetworkingRequired
	result.DockerSocketMounted = gui.DockerSocketMounted
	result.BroadHostMountRequired = gui.BroadHostMountRequired
	result.RawHostPathExposed = gui.RawHostPathExposed
	result.RawExecutablePathExposed = false
	result.RawCommandExposed = gui.RawCommandExposed
	result.SkipReason = gui.SkipReason
	result.FailureReason = gui.FailureReason
	if result.SmokePassed {
		result.DesktopSafeSummary = app.DisplayName + " passed the gated Runtime managed guest GUI smoke."
	} else {
		result.DesktopSafeSummary = app.DisplayName + " did not pass the gated Runtime managed guest GUI smoke."
	}
	return result, nil
}

func knownPortableGuestGUIAppPath(guiAppPath string) string {
	clean := strings.TrimSpace(guiAppPath)
	if clean == "" {
		return DefaultGuestGUIApp
	}
	return clean
}

func PreviewKnownPortableLaunchBridge(request KnownLaunchBridgeRequest) (KnownLaunchBridgeResult, error) {
	dispatchPreview, err := PreviewKnownPortableDispatch(KnownDispatchRequest{
		AppID:     request.AppID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownLaunchBridgeResult{}, err
	}
	result := baseKnownLaunchBridgeResult(dispatchPreview)
	if !managedLauncherArgvMatches(request.ManagedLauncherArgv, dispatchPreview.ManagedLauncherArgv) {
		result.Status = "bridge-blocked"
		result.BlockedReason = "managed launcher argv does not match the Runtime-owned launch surface"
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " launch bridge rejected a launcher request that did not match the Runtime-owned launch surface."
		return result, nil
	}

	result.LauncherArgvAccepted = true
	if dispatchPreview.DispatchReady {
		result.Status = "bridge-ready"
		result.DispatchSmokeRequestMaterialized = true
		result.PreparationRequired = false
		result.BlockedReason = ""
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " launch bridge materialized a gated dispatch smoke request."
	} else {
		result.Status = "bridge-blocked"
		result.BlockedReason = dispatchPreview.BlockedReason
		result.DesktopSafeSummary = dispatchPreview.DisplayName + " launch bridge accepted the managed launcher, but artifact preparation is required before dispatch smoke materialization."
	}
	return result, nil
}

func MaterializeKnownPortableLaunchProfile(request KnownLaunchProfileMaterializeRequest) (KnownLaunchProfileMaterializeResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownLaunchProfileMaterializeResult{}, err
	}
	result := baseKnownLaunchProfileMaterializeResult(app)
	result.RunnerConfigured = strings.TrimSpace(request.RunnerPath) != ""
	result.RunnerBottleConfigured = strings.TrimSpace(request.RunnerBottle) != ""
	result.RunnerArgumentCount = len(runnerInvocationArguments(Request{
		RunnerBottle:    request.RunnerBottle,
		RunnerArguments: append([]string{}, request.RunnerArguments...),
	}))
	result.SkipBootstrap = request.SkipBootstrap
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
		if actual != "" {
			result.CacheStatus = "checksum-mismatch"
			result.SkipReason = "known Windows app artifact checksum mismatch"
		} else {
			result.CacheStatus = "missing"
			result.SkipReason = "known Windows app artifact unavailable"
		}
		return result, nil
	}
	result.CacheStatus = "verified"
	result.ArtifactVerified = true

	stateRoot, err := knownLaunchProfileStateRoot(request.StateRoot, request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	profilePath, err := knownLaunchProfileOutputPath(request.ProfileOutput, stateRoot, app)
	if err != nil {
		return result, err
	}
	profile := SmokeProfileFromRequest(Request{
		ExecutablePath:  cachePath,
		Arguments:       append([]string{}, app.Arguments...),
		RunnerArguments: append([]string{}, request.RunnerArguments...),
		RunnerBottle:    request.RunnerBottle,
		StateRoot:       stateRoot,
		RunnerPath:      request.RunnerPath,
		Timeout:         30 * time.Second,
		ExpectedMarker:  app.ExpectedMarker,
		SuccessMode:     SuccessModeMarker,
		RedactOutput:    true,
		SkipBootstrap:   request.SkipBootstrap,
		StageAppDir:     true,
	})
	if err := os.MkdirAll(filepath.Dir(profilePath), 0o700); err != nil {
		return result, fmt.Errorf("create known app profile directory: %w", err)
	}
	profileData, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return result, fmt.Errorf("encode known app launch profile: %w", err)
	}
	if err := os.WriteFile(profilePath, append(profileData, '\n'), 0o600); err != nil {
		return result, fmt.Errorf("write known app launch profile: %w", err)
	}
	result.ProfileWritten = true
	result.ProfileFileName = filepath.Base(profilePath)

	applicationID := strings.TrimSpace(request.ApplicationID)
	if applicationID == "" {
		applicationID = "org.xnix.known." + app.ID
	}
	displayName := strings.TrimSpace(request.DisplayName)
	if displayName == "" {
		displayName = app.DisplayName
	}
	launcherRecord, err := RecordLauncherBundle(LauncherBundleRequest{
		ProfilePath:      profilePath,
		ApplicationID:    applicationID,
		DisplayName:      displayName,
		RuntimeBinary:    request.RuntimeBinary,
		RuntimeArguments: append([]string{}, request.RuntimeArguments...),
		LauncherMode:     LauncherModeLaunch,
	})
	if err != nil {
		return result, err
	}
	result.LauncherBundlePayload = &launcherRecord
	result.LauncherBundleWritten = launcherRecord.Status == PassedStatus && launcherRecord.FilesWritten
	result.LauncherMode = launcherRecord.LauncherMode
	result.LauncherCommand = launcherRecord.LauncherCommand
	if !result.LauncherBundleWritten {
		result.Status = FailedStatus
		result.FailureReason = launcherRecord.FailureReason
		return result, nil
	}
	result.Status = PassedStatus
	return result, nil
}

func PrepareKnownPortableLaunchProfile(ctx context.Context, request KnownPrepareLaunchProfileRequest) (KnownPrepareLaunchProfileResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownPrepareLaunchProfileResult{}, err
	}
	result := baseKnownPrepareLaunchProfileResult(app, request.AllowDownload)
	result.RunnerConfigured = strings.TrimSpace(request.RunnerPath) != ""
	result.RunnerBottleConfigured = strings.TrimSpace(request.RunnerBottle) != ""
	result.RunnerArgumentCount = len(runnerInvocationArguments(Request{
		RunnerBottle:    request.RunnerBottle,
		RunnerArguments: append([]string{}, request.RunnerArguments...),
	}))
	result.SkipBootstrap = request.SkipBootstrap
	fetch, err := FetchKnownPortableApp(ctx, KnownFetchRequest{
		AppID:         app.ID,
		CacheRoot:     request.CacheRoot,
		AllowDownload: request.AllowDownload,
		HTTPClient:    request.HTTPClient,
		Timeout:       request.Timeout,
	})
	if err != nil {
		return result, err
	}
	if !request.AllowDownload {
		fetch.NetworkRequired = false
	}
	result.FetchPayload = fetch
	result.FetchStatus = fetch.Status
	result.FetchCacheStatus = fetch.CacheStatus
	result.Downloaded = fetch.Downloaded
	result.ChecksumVerified = fetch.ChecksumVerified
	result.NetworkRequired = request.AllowDownload
	result.HostRootModified = fetch.HostRootModified
	result.HostNetworkingRequired = fetch.HostNetworkingRequired
	result.DockerSocketMounted = fetch.DockerSocketMounted
	result.BroadHostMountRequired = fetch.BroadHostMountRequired
	result.RawHostPathExposed = fetch.RawHostPathExposed
	if fetch.Status != PassedStatus || !fetch.ChecksumVerified {
		result.Status = fetch.Status
		if result.Status == FailedStatus {
			result.FailureReason = fetch.FailureReason
		} else {
			result.SkipReason = fetch.SkipReason
		}
		return result, nil
	}

	materialized, err := MaterializeKnownPortableLaunchProfile(KnownLaunchProfileMaterializeRequest{
		AppID:            app.ID,
		CacheRoot:        request.CacheRoot,
		StateRoot:        request.StateRoot,
		ProfileOutput:    request.ProfileOutput,
		ApplicationID:    request.ApplicationID,
		DisplayName:      request.DisplayName,
		RuntimeBinary:    request.RuntimeBinary,
		RuntimeArguments: append([]string{}, request.RuntimeArguments...),
		RunnerPath:       request.RunnerPath,
		RunnerBottle:     request.RunnerBottle,
		RunnerArguments:  append([]string{}, request.RunnerArguments...),
		SkipBootstrap:    request.SkipBootstrap,
	})
	if err != nil {
		return result, err
	}
	result.MaterializePayload = &materialized
	result.MaterializeStatus = materialized.Status
	result.ProfileWritten = materialized.ProfileWritten
	result.LauncherBundleWritten = materialized.LauncherBundleWritten
	result.RunnerConfigured = materialized.RunnerConfigured
	result.RunnerBottleConfigured = materialized.RunnerBottleConfigured
	result.RunnerArgumentCount = materialized.RunnerArgumentCount
	result.SkipBootstrap = materialized.SkipBootstrap
	result.RawExecutablePathExposed = materialized.RawExecutablePathExposed
	result.RawProfilePathExposed = materialized.RawProfilePathExposed
	result.RawStateRootPathExposed = materialized.RawStateRootPathExposed
	result.RawRuntimeArgvExposed = materialized.RawRuntimeArgvExposed
	result.PrivilegedContainerRequired = materialized.PrivilegedContainerRequired
	result.HostRootModified = result.HostRootModified || materialized.HostRootModified
	result.HostNetworkingRequired = result.HostNetworkingRequired || materialized.HostNetworkingRequired
	result.DockerSocketMounted = result.DockerSocketMounted || materialized.DockerSocketMounted
	result.BroadHostMountRequired = result.BroadHostMountRequired || materialized.BroadHostMountRequired
	result.Status = materialized.Status
	result.SkipReason = materialized.SkipReason
	result.FailureReason = materialized.FailureReason
	return result, nil
}

func PrepareAndLaunchKnownPortableProfile(ctx context.Context, request KnownPrepareAndLaunchProfileRequest) (KnownPrepareAndLaunchProfileResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownPrepareAndLaunchProfileResult{}, err
	}
	result := baseKnownPrepareAndLaunchProfileResult(app, request)
	prepare, err := PrepareKnownPortableLaunchProfile(ctx, KnownPrepareLaunchProfileRequest{
		AppID:            app.ID,
		CacheRoot:        request.CacheRoot,
		StateRoot:        request.StateRoot,
		ProfileOutput:    request.ProfileOutput,
		ApplicationID:    request.ApplicationID,
		DisplayName:      request.DisplayName,
		RuntimeBinary:    request.RuntimeBinary,
		RuntimeArguments: append([]string{}, request.RuntimeArguments...),
		RunnerPath:       request.RunnerPath,
		RunnerBottle:     request.RunnerBottle,
		RunnerArguments:  append([]string{}, request.RunnerArguments...),
		SkipBootstrap:    request.SkipBootstrap,
		AllowDownload:    request.AllowDownload,
		HTTPClient:       request.HTTPClient,
		Timeout:          request.Timeout,
	})
	if err != nil {
		return result, err
	}
	result.PreparePayload = prepare
	result.PrepareStatus = prepare.Status
	result.ProfileWritten = prepare.ProfileWritten
	result.LauncherBundleWritten = prepare.LauncherBundleWritten
	result.RunnerConfigured = prepare.RunnerConfigured
	result.RunnerBottleConfigured = prepare.RunnerBottleConfigured
	result.RunnerArgumentCount = prepare.RunnerArgumentCount
	result.SkipBootstrap = prepare.SkipBootstrap
	result.NetworkRequired = prepare.NetworkRequired
	result.HostRootModified = prepare.HostRootModified
	result.PrivilegedContainerRequired = prepare.PrivilegedContainerRequired
	result.HostNetworkingRequired = prepare.HostNetworkingRequired
	result.DockerSocketMounted = prepare.DockerSocketMounted
	result.BroadHostMountRequired = prepare.BroadHostMountRequired
	result.RawHostPathExposed = prepare.RawHostPathExposed
	result.RawExecutablePathExposed = prepare.RawExecutablePathExposed
	result.RawProfilePathExposed = prepare.RawProfilePathExposed
	result.RawStateRootPathExposed = prepare.RawStateRootPathExposed
	result.RawRuntimeArgvExposed = prepare.RawRuntimeArgvExposed
	if prepare.Status != PassedStatus || !prepare.ProfileWritten {
		result.Status = prepare.Status
		result.SkipReason = prepare.SkipReason
		result.FailureReason = prepare.FailureReason
		return result, nil
	}

	stateRoot, err := knownLaunchProfileStateRoot(request.StateRoot, request.CacheRoot, app)
	if err != nil {
		return result, err
	}
	profilePath, err := knownLaunchProfileOutputPath(request.ProfileOutput, stateRoot, app)
	if err != nil {
		return result, err
	}
	launch, err := LaunchProfile(ctx, LaunchProfileRequest{ProfilePath: profilePath})
	if err != nil {
		return result, err
	}
	result.LaunchPayload = &launch
	result.LaunchStatus = launch.Status
	result.Status = launch.Status
	result.LaunchAttempted = launch.LaunchAttempted
	result.RunnerAvailable = launch.RunnerAvailable
	result.ExecutableFormat = launch.ExecutableFormat
	result.WindowsExecutableSignature = launch.WindowsExecutableSignature
	result.ExecutableArchitecture = launch.ExecutableArchitecture
	result.ExecutableArchitectureReady = launch.ExecutableArchitectureReady
	result.WineArchitecture = launch.WineArchitecture
	result.ApplicationWorkspaceMode = launch.ApplicationWorkspaceMode
	result.RawOutputRedacted = launch.RawOutputRedacted
	result.NextAction = launch.NextAction
	result.SkipReason = launch.SkipReason
	result.FailureReason = launch.FailureReason
	result.HostRootModified = result.HostRootModified || launch.HostRootModified
	result.PrivilegedContainerRequired = result.PrivilegedContainerRequired || launch.PrivilegedContainerRequired
	result.HostNetworkingRequired = result.HostNetworkingRequired || launch.HostNetworkingRequired
	result.DockerSocketMounted = result.DockerSocketMounted || launch.DockerSocketMounted
	result.BroadHostMountRequired = result.BroadHostMountRequired || launch.BroadHostMountRequired
	result.DockerExecuted = launch.DockerExecuted
	result.QEMUExecuted = launch.QEMUExecuted
	result.WineExecuted = launch.WineExecuted
	result.ColimaExecuted = launch.ColimaExecuted
	result.NetworkChecksRun = launch.NetworkChecksRun
	result.PackageManagerInvoked = launch.PackageManagerInvoked
	result.RawProfilePathExposed = result.RawProfilePathExposed || launch.RawProfilePathExposed
	result.RawExecutablePathExposed = result.RawExecutablePathExposed || launch.RawExecutablePathExposed
	result.RawRunnerPathExposed = launch.RawRunnerPathExposed
	return result, nil
}

func RunKnownPortableApp(ctx context.Context, request KnownRunRequest) (KnownRunResult, error) {
	app, err := LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownRunResult{}, err
	}
	backend := knownRunBackend(request.Backend)
	result := baseKnownRunResult(app, backend)
	result.RawOutputRedacted = request.RedactOutput
	switch backend {
	case KnownRunBackendLocal:
		local, err := PrepareAndLaunchKnownPortableProfile(ctx, KnownPrepareAndLaunchProfileRequest{
			AppID:            app.ID,
			CacheRoot:        request.CacheRoot,
			StateRoot:        request.StateRoot,
			ProfileOutput:    request.ProfileOutput,
			ApplicationID:    request.ApplicationID,
			DisplayName:      request.DisplayName,
			RuntimeBinary:    request.RuntimeBinary,
			RuntimeArguments: append([]string{}, request.RuntimeArguments...),
			RunnerPath:       request.RunnerPath,
			RunnerBottle:     request.RunnerBottle,
			RunnerArguments:  append([]string{}, request.RunnerArguments...),
			SkipBootstrap:    request.SkipBootstrap,
			AllowDownload:    request.AllowDownload,
			HTTPClient:       request.HTTPClient,
			Timeout:          request.Timeout,
		})
		if err != nil {
			return result, err
		}
		result.LocalPayload = &local
		result.Status = local.Status
		result.BackendReady = local.RunnerAvailable
		result.LaunchAttempted = local.LaunchAttempted
		result.RunnerAvailable = local.RunnerAvailable
		result.ChecksumVerified = local.PreparePayload.ChecksumVerified
		result.ProfileWritten = local.ProfileWritten
		result.LauncherBundleWritten = local.LauncherBundleWritten
		result.ExecutableFormat = local.ExecutableFormat
		result.ExecutableArchitecture = local.ExecutableArchitecture
		result.WineArchitecture = local.WineArchitecture
		result.ApplicationWorkspaceMode = local.ApplicationWorkspaceMode
		result.RawOutputRedacted = local.RawOutputRedacted
		result.LoopbackOnlyNetworking = false
		result.QEMURequired = false
		result.NetworkRequired = local.NetworkRequired
		result.HostRootModified = local.HostRootModified
		result.PrivilegedContainerRequired = local.PrivilegedContainerRequired
		result.HostNetworkingRequired = local.HostNetworkingRequired
		result.DockerSocketMounted = local.DockerSocketMounted
		result.BroadHostMountRequired = local.BroadHostMountRequired
		result.DockerExecuted = local.DockerExecuted
		result.QEMUExecuted = local.QEMUExecuted
		result.WineExecuted = local.WineExecuted
		result.ColimaExecuted = local.ColimaExecuted
		result.NetworkChecksRun = local.NetworkChecksRun
		result.PackageManagerInvoked = local.PackageManagerInvoked
		result.RawHostPathExposed = local.RawHostPathExposed
		result.RawExecutablePathExposed = local.RawExecutablePathExposed
		result.RawProfilePathExposed = local.RawProfilePathExposed
		result.RawStateRootPathExposed = local.RawStateRootPathExposed
		result.RawRuntimeArgvExposed = local.RawRuntimeArgvExposed
		result.RawRunnerPathExposed = local.RawRunnerPathExposed
		result.NextAction = local.NextAction
		result.SkipReason = local.SkipReason
		result.FailureReason = local.FailureReason
	case KnownRunBackendGuestWine:
		result.GuestStartMode = "external"
		var qemuGuest *QEMUStartedGuest
		defer func() {
			if qemuGuest != nil {
				qemuGuest.Stop()
			}
		}()
		if request.StartQEMU {
			result.GuestStartMode = "go-qemu"
			guestPreflight, ok, err := preflightKnownGuestArtifact(request.CacheRoot, app)
			if err != nil {
				return result, err
			}
			if !ok {
				result.GuestPayload = &guestPreflight
				result.Status = guestPreflight.Status
				result.ChecksumVerified = false
				result.SkipReason = guestPreflight.SkipReason
				return result, nil
			}
			result.ChecksumVerified = true
		}
		if request.StartQEMU {
			result.GuestStartAttempted = true
			result.GuestStartMode = "go-qemu"
			guest, err := StartQEMUStartedGuest(ctx, QEMUStartedGuestRequest{
				Binary:        request.QEMUBinary,
				KernelImage:   request.QEMUKernelImage,
				Memory:        request.QEMUMemory,
				CPUCount:      request.QEMUCPUCount,
				CPUModel:      request.QEMUCPUModel,
				Host:          request.Host,
				Port:          request.Port,
				User:          request.User,
				KeyPath:       request.KeyPath,
				SSHPath:       request.SSHPath,
				BootTimeout:   request.QEMUBootTimeout,
				SerialLogPath: request.QEMUSerialLog,
			})
			if err != nil {
				result.Status = SkippedStatus
				result.SkipReason = safeQEMUGuestStartSkipReason(err)
				result.QEMUSerialLogWritten = strings.TrimSpace(request.QEMUSerialLog) != ""
				return result, nil
			}
			result.GuestStarted = true
			result.GuestHost = guest.Host
			result.GuestPort = guest.Port
			result.GuestPortAuto = guest.AutoPort
			result.QEMUExecuted = true
			qemuGuest = guest
		}
		guest, err := RunKnownPortableGuestSmoke(ctx, KnownGuestRequest{
			AppID:        app.ID,
			CacheRoot:    request.CacheRoot,
			Arguments:    append([]string{}, request.Arguments...),
			Host:         stringDefault(result.GuestHost, request.Host),
			Port:         stringDefault(result.GuestPort, request.Port),
			User:         request.User,
			KeyPath:      request.KeyPath,
			RemoteDir:    request.RemoteDir,
			SSHPath:      request.SSHPath,
			SCPPath:      request.SCPPath,
			Timeout:      request.Timeout,
			RedactOutput: request.RedactOutput,
		})
		if err != nil {
			return result, err
		}
		result.GuestPayload = &guest
		result.Status = guest.Status
		result.BackendReady = guest.Guest.GuestReachable && guest.Guest.WineAvailable
		result.LaunchAttempted = guest.Guest.ExecutableCopied
		result.RunnerAvailable = guest.Guest.WineAvailable
		result.ChecksumVerified = guest.ChecksumVerified
		result.ExecutableCopied = guest.Guest.ExecutableCopied
		result.MarkerObserved = guest.Guest.MarkerObserved
		result.RawOutputRedacted = request.RedactOutput || guest.Guest.RawOutputRedacted
		result.LoopbackOnlyNetworking = guest.LoopbackOnlyNetworking
		result.QEMURequired = guest.QEMURequired
		result.NetworkRequired = false
		result.HostRootModified = guest.HostRootModified
		result.PrivilegedContainerRequired = guest.PrivilegedContainerRequired
		result.HostNetworkingRequired = guest.HostNetworkingRequired
		result.DockerSocketMounted = guest.DockerSocketMounted
		result.BroadHostMountRequired = guest.BroadHostMountRequired
		result.DockerExecuted = false
		result.QEMUExecuted = result.QEMUExecuted || request.StartQEMU
		result.WineExecuted = guest.Guest.WineAvailable
		result.ColimaExecuted = false
		result.NetworkChecksRun = false
		result.PackageManagerInvoked = false
		result.RawHostPathExposed = guest.RawHostPathExposed
		result.SkipReason = guest.SkipReason
		result.FailureReason = guest.FailureReason
		if qemuGuest != nil {
			result.QEMUSerialLogWritten = qemuGuest.Stop()
			qemuGuest = nil
		}
	default:
		result.Status = FailedStatus
		result.FailureReason = "unsupported known app run backend"
	}
	return result, nil
}

func preflightKnownGuestArtifact(cacheRoot string, app KnownPortableApp) (KnownGuestResult, bool, error) {
	result := baseKnownGuestResult(app)
	cachePath, err := knownAppCachePath(cacheRoot, app)
	if err != nil {
		return result, false, err
	}
	actual, ok, err := verifyKnownAppFile(cachePath, app)
	if err != nil {
		return result, false, err
	}
	result.ActualSHA256 = actual
	if !ok {
		result.Status = SkippedStatus
		result.SkipReason = "known Windows app artifact unavailable or checksum mismatch"
		return result, false, nil
	}
	result.Status = PassedStatus
	result.ChecksumVerified = true
	return result, true, nil
}

func safeQEMUGuestStartSkipReason(err error) string {
	text := err.Error()
	switch {
	case strings.Contains(text, "qemu guest runner unavailable"):
		return "qemu guest runner unavailable"
	case strings.Contains(text, "qemu guest kernel unavailable"):
		return "qemu guest kernel unavailable"
	case strings.Contains(text, "guest ssh transport unavailable"):
		return "guest ssh transport unavailable"
	case strings.Contains(text, "qemu guest loopback host required"):
		return "qemu guest loopback host required"
	case strings.Contains(text, "qemu guest auto port unavailable"):
		return "qemu guest auto port unavailable"
	case strings.Contains(text, "timed out waiting for ssh"):
		return "qemu guest timed out waiting for ssh"
	case strings.Contains(text, "exited before ssh became ready"):
		return "qemu guest exited before ssh became ready"
	default:
		return "qemu guest start failed"
	}
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

func baseKnownLaunchProfileMaterializeResult(app KnownPortableApp) KnownLaunchProfileMaterializeResult {
	return KnownLaunchProfileMaterializeResult{
		SchemaVersion:               KnownLaunchProfileSchemaVersion,
		RequestType:                 KnownLaunchProfileRequestType,
		Status:                      SkippedStatus,
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		ExpectedSHA256:              strings.ToLower(app.SHA256),
		CacheStatus:                 "unknown",
		LauncherMode:                LauncherModeLaunch,
		LauncherCommand:             LaunchProfileRequestType,
		ExpectedMarker:              app.ExpectedMarker,
		SuccessMode:                 SuccessModeMarker,
		ApplicationWorkspaceMode:    ApplicationWorkspaceModeStaged,
		RawHostPathExposed:          false,
		RawExecutablePathExposed:    false,
		RawProfilePathExposed:       false,
		RawStateRootPathExposed:     false,
		RawRuntimeArgvExposed:       false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
	}
}

func baseKnownPrepareLaunchProfileResult(app KnownPortableApp, allowDownload bool) KnownPrepareLaunchProfileResult {
	return KnownPrepareLaunchProfileResult{
		SchemaVersion:               KnownPrepareLaunchSchemaVersion,
		RequestType:                 KnownPrepareLaunchRequestType,
		Status:                      SkippedStatus,
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		AllowDownload:               allowDownload,
		FetchStatus:                 SkippedStatus,
		FetchCacheStatus:            "unknown",
		MaterializeStatus:           "not-run",
		NetworkRequired:             allowDownload,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawHostPathExposed:          false,
		RawExecutablePathExposed:    false,
		RawProfilePathExposed:       false,
		RawStateRootPathExposed:     false,
		RawRuntimeArgvExposed:       false,
	}
}

func baseKnownPrepareAndLaunchProfileResult(app KnownPortableApp, request KnownPrepareAndLaunchProfileRequest) KnownPrepareAndLaunchProfileResult {
	return KnownPrepareAndLaunchProfileResult{
		SchemaVersion:               KnownPrepareAndLaunchSchemaVersion,
		RequestType:                 KnownPrepareAndLaunchRequestType,
		Status:                      SkippedStatus,
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		AllowDownload:               request.AllowDownload,
		PrepareStatus:               "not-run",
		LaunchStatus:                "not-run",
		RunnerConfigured:            strings.TrimSpace(request.RunnerPath) != "",
		RunnerBottleConfigured:      strings.TrimSpace(request.RunnerBottle) != "",
		RunnerArgumentCount:         len(runnerInvocationArguments(Request{RunnerBottle: request.RunnerBottle, RunnerArguments: append([]string{}, request.RunnerArguments...)})),
		SkipBootstrap:               request.SkipBootstrap,
		ExecutableFormat:            "unknown",
		ExecutableArchitecture:      "unknown",
		WineArchitecture:            "unknown",
		ApplicationWorkspaceMode:    ApplicationWorkspaceModeDirect,
		RawOutputRedacted:           true,
		NetworkRequired:             request.AllowDownload,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		DockerExecuted:              false,
		QEMUExecuted:                false,
		WineExecuted:                false,
		ColimaExecuted:              false,
		NetworkChecksRun:            false,
		PackageManagerInvoked:       false,
		RawHostPathExposed:          false,
		RawExecutablePathExposed:    false,
		RawProfilePathExposed:       false,
		RawStateRootPathExposed:     false,
		RawRuntimeArgvExposed:       false,
		RawRunnerPathExposed:        false,
	}
}

func baseKnownRunResult(app KnownPortableApp, backend string) KnownRunResult {
	return KnownRunResult{
		SchemaVersion:               KnownRunSchemaVersion,
		RequestType:                 KnownRunRequestType,
		Status:                      SkippedStatus,
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		Architecture:                app.Architecture,
		ExecutableName:              app.ExecutableName,
		Backend:                     backend,
		GuestStartMode:              "none",
		ExecutableFormat:            "unknown",
		ExecutableArchitecture:      "unknown",
		WineArchitecture:            "unknown",
		ApplicationWorkspaceMode:    ApplicationWorkspaceModeDirect,
		RawOutputRedacted:           true,
		LoopbackOnlyNetworking:      backend == KnownRunBackendGuestWine,
		QEMURequired:                backend == KnownRunBackendGuestWine,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		DockerExecuted:              false,
		QEMUExecuted:                false,
		WineExecuted:                false,
		ColimaExecuted:              false,
		NetworkChecksRun:            false,
		PackageManagerInvoked:       false,
		RawHostPathExposed:          false,
		RawExecutablePathExposed:    false,
		RawProfilePathExposed:       false,
		RawStateRootPathExposed:     false,
		RawRuntimeArgvExposed:       false,
		RawRunnerPathExposed:        false,
		RawQEMUPathExposed:          false,
	}
}

func knownRunBackend(value string) string {
	switch strings.TrimSpace(value) {
	case "", KnownRunBackendLocal:
		return KnownRunBackendLocal
	case KnownRunBackendGuestWine:
		return KnownRunBackendGuestWine
	default:
		return strings.TrimSpace(value)
	}
}

func baseKnownLaunchBridgeResult(dispatchPreview KnownDispatchResult) KnownLaunchBridgeResult {
	return KnownLaunchBridgeResult{
		SchemaVersion:                    KnownLaunchBridgeSchemaVersion,
		RequestType:                      KnownLaunchBridgeRequestType,
		Source:                           KnownDispatchRequestType,
		Status:                           "bridge-blocked",
		Desktop:                          dispatchPreview.Desktop,
		EntryPointID:                     dispatchPreview.EntryPointID,
		DesktopFile:                      dispatchPreview.DesktopFile,
		LaunchSurfaceID:                  dispatchPreview.LaunchSurfaceID,
		DesktopActionID:                  dispatchPreview.DesktopActionID,
		ManagedLauncher:                  dispatchPreview.ManagedLauncher,
		ManagedLauncherArgv:              append([]string{}, dispatchPreview.ManagedLauncherArgv...),
		LauncherArgvAccepted:             false,
		RequestID:                        dispatchPreview.RequestID,
		DispatchID:                       dispatchPreview.DispatchID,
		RuntimeMethod:                    "BridgeKnownLauncherToDispatchSmoke",
		DispatchSmokeRequestType:         KnownDispatchSmokeRequestType,
		DispatchSmokeRequestMaterialized: false,
		AppID:                            dispatchPreview.AppID,
		DisplayName:                      dispatchPreview.DisplayName,
		AppVersion:                       dispatchPreview.AppVersion,
		Architecture:                     dispatchPreview.Architecture,
		DispatchGate:                     dispatchPreview.DispatchGate,
		RunnerLane:                       dispatchPreview.RunnerLane,
		GuestBoundaryRequired:            true,
		GuestBoundarySupplied:            false,
		SmokeHarnessRequired:             true,
		CacheStatus:                      dispatchPreview.CacheStatus,
		ArtifactVerified:                 dispatchPreview.ArtifactVerified,
		LaunchRequestCreated:             dispatchPreview.LaunchRequestCreated,
		DispatchPreviewCreated:           dispatchPreview.DispatchPreviewCreated,
		BridgePreviewCreated:             true,
		DispatchReady:                    dispatchPreview.DispatchReady,
		PreparationRequired:              dispatchPreview.PreparationRequired,
		RuntimeOwnedRequest:              dispatchPreview.RuntimeOwnedRequest,
		RuntimeOwnedLaunch:               dispatchPreview.RuntimeOwnedLaunch,
		RuntimeOwnedDispatch:             dispatchPreview.RuntimeOwnedDispatch,
		RuntimeOwnedBridge:               true,
		KDEPresentationOnly:              dispatchPreview.KDEPresentationOnly,
		DryRun:                           true,
		DispatchStarted:                  false,
		ExecutionStarted:                 false,
		BackendProcessStarted:            false,
		HostRootModified:                 dispatchPreview.HostRootModified,
		HostNetworkingRequired:           dispatchPreview.HostNetworkingRequired,
		DockerSocketMounted:              dispatchPreview.DockerSocketMounted,
		BroadHostMountRequired:           dispatchPreview.BroadHostMountRequired,
		RawHostPathExposed:               dispatchPreview.RawHostPathExposed,
		RawExecutablePathExposed:         dispatchPreview.RawExecutablePathExposed,
		RawCommandExposed:                dispatchPreview.RawCommandExposed,
		BackendDetailsExposed:            dispatchPreview.BackendDetailsExposed,
		DesktopSafeSummary:               dispatchPreview.DesktopSafeSummary,
		BlockedReason:                    dispatchPreview.BlockedReason,
	}
}

func baseKnownDispatchSmokeResult(dispatchPreview KnownDispatchResult, guestBoundary string) KnownDispatchSmokeResult {
	return KnownDispatchSmokeResult{
		SchemaVersion:               KnownDispatchSmokeSchemaVersion,
		RequestType:                 KnownDispatchSmokeRequestType,
		Source:                      KnownDispatchRequestType,
		Status:                      "dispatch-blocked",
		RequestID:                   dispatchPreview.RequestID,
		DispatchID:                  dispatchPreview.DispatchID,
		RuntimeMethod:               dispatchPreview.RuntimeMethod,
		AppID:                       dispatchPreview.AppID,
		DisplayName:                 dispatchPreview.DisplayName,
		AppVersion:                  dispatchPreview.AppVersion,
		Architecture:                dispatchPreview.Architecture,
		DispatchGate:                dispatchPreview.DispatchGate,
		RunnerLane:                  dispatchPreview.RunnerLane,
		GuestBoundary:               guestBoundary,
		CacheStatus:                 dispatchPreview.CacheStatus,
		ArtifactVerified:            dispatchPreview.ArtifactVerified,
		LaunchRequestCreated:        dispatchPreview.LaunchRequestCreated,
		DispatchPreviewCreated:      dispatchPreview.DispatchPreviewCreated,
		DispatchReady:               dispatchPreview.DispatchReady,
		DispatchAllowed:             false,
		DispatchStarted:             false,
		ExecutionStarted:            false,
		ManagedGuestRunnerInvoked:   false,
		ManagedGuestReachable:       false,
		ManagedGuestRuntimeReady:    false,
		ManagedArtifactCopied:       false,
		MarkerObserved:              false,
		SmokePassed:                 false,
		ExitCode:                    -1,
		DurationMillis:              0,
		RuntimeOwnedRequest:         dispatchPreview.RuntimeOwnedRequest,
		RuntimeOwnedLaunch:          dispatchPreview.RuntimeOwnedLaunch,
		RuntimeOwnedDispatch:        dispatchPreview.RuntimeOwnedDispatch,
		KDEPresentationOnly:         dispatchPreview.KDEPresentationOnly,
		HostRootModified:            dispatchPreview.HostRootModified,
		PrivilegedContainerRequired: false,
		HostNetworkingRequired:      dispatchPreview.HostNetworkingRequired,
		DockerSocketMounted:         dispatchPreview.DockerSocketMounted,
		BroadHostMountRequired:      dispatchPreview.BroadHostMountRequired,
		RawHostPathExposed:          dispatchPreview.RawHostPathExposed,
		RawExecutablePathExposed:    dispatchPreview.RawExecutablePathExposed,
		RawCommandExposed:           dispatchPreview.RawCommandExposed,
		BackendDetailsExposed:       dispatchPreview.BackendDetailsExposed,
		DesktopSafeSummary:          dispatchPreview.DesktopSafeSummary,
		BlockedReason:               dispatchPreview.BlockedReason,
	}
}

func baseKnownDispatchResult(launchRequest KnownLaunchRequestResult) KnownDispatchResult {
	return KnownDispatchResult{
		SchemaVersion:            KnownDispatchSchemaVersion,
		RequestType:              KnownDispatchRequestType,
		Source:                   KnownLaunchRequestType,
		Status:                   "dispatch-blocked",
		RequestID:                launchRequest.RequestID,
		DispatchID:               "known-app-dispatch-" + launchRequest.AppID,
		RuntimeMethod:            "DispatchKnownWindowsApp",
		AppID:                    launchRequest.AppID,
		DisplayName:              launchRequest.DisplayName,
		AppVersion:               launchRequest.AppVersion,
		Architecture:             launchRequest.Architecture,
		Desktop:                  launchRequest.Desktop,
		EntryPointID:             launchRequest.EntryPointID,
		DesktopFile:              launchRequest.DesktopFile,
		LaunchSurfaceID:          launchRequest.LaunchSurfaceID,
		DesktopActionID:          launchRequest.DesktopActionID,
		ManagedLauncher:          launchRequest.ManagedLauncher,
		ManagedLauncherArgv:      append([]string{}, launchRequest.ManagedLauncherArgv...),
		DispatchGate:             launchRequest.DispatchGate,
		RunnerLane:               "known-app-guest-smoke",
		RunnerRequestType:        "managed-known-app-guest-smoke",
		CacheStatus:              launchRequest.CacheStatus,
		ArtifactVerified:         launchRequest.ArtifactVerified,
		LaunchRequestCreated:     launchRequest.LaunchRequestCreated,
		DispatchPreviewCreated:   true,
		DispatchAllowed:          false,
		DispatchReady:            false,
		PreparationRequired:      launchRequest.PreparationRequired,
		RuntimeOwnedRequest:      launchRequest.RuntimeOwnedRequest,
		RuntimeOwnedLaunch:       launchRequest.RuntimeOwnedLaunch,
		RuntimeOwnedDispatch:     true,
		KDEPresentationOnly:      launchRequest.KDEPresentationOnly,
		DryRun:                   true,
		DispatchStarted:          false,
		ExecutionStarted:         false,
		BackendProcessStarted:    false,
		HostRootModified:         launchRequest.HostRootModified,
		HostNetworkingRequired:   launchRequest.HostNetworkingRequired,
		DockerSocketMounted:      launchRequest.DockerSocketMounted,
		BroadHostMountRequired:   launchRequest.BroadHostMountRequired,
		RawHostPathExposed:       launchRequest.RawHostPathExposed,
		RawExecutablePathExposed: launchRequest.RawExecutablePathExposed,
		RawCommandExposed:        launchRequest.RawCommandExposed,
		BackendDetailsExposed:    launchRequest.BackendDetailsExposed,
		DesktopSafeSummary:       launchRequest.DesktopSafeSummary,
		BlockedReason:            launchRequest.BlockedReason,
	}
}

func baseKnownLaunchRequestResult(launcher KnownKDELauncherResult) KnownLaunchRequestResult {
	return KnownLaunchRequestResult{
		SchemaVersion:            KnownLaunchRequestSchemaVersion,
		RequestType:              KnownLaunchRequestType,
		Source:                   KnownKDELauncherRequestType,
		Status:                   "request-blocked",
		RequestID:                "known-app-launch-request-" + launcher.AppID,
		RuntimeMethod:            "LaunchKnownWindowsApp",
		AppID:                    launcher.AppID,
		DisplayName:              launcher.DisplayName,
		AppVersion:               launcher.AppVersion,
		Architecture:             launcher.Architecture,
		Desktop:                  launcher.Desktop,
		EntryPointID:             launcher.EntryPointID,
		DesktopFile:              launcher.DesktopFile,
		LaunchSurfaceID:          launcher.LaunchSurfaceID,
		DesktopActionID:          launcher.DesktopActionID,
		ManagedLauncher:          launcher.ManagedLauncher,
		ManagedLauncherArgv:      append([]string{}, launcher.ManagedLauncherArgv...),
		DispatchGate:             "managed-known-app-guest-smoke",
		CacheStatus:              launcher.CacheStatus,
		ArtifactVerified:         launcher.ArtifactVerified,
		LaunchVisible:            launcher.LaunchVisible,
		LaunchAllowed:            false,
		LaunchRequestCreated:     true,
		DispatchReady:            false,
		PreparationRequired:      launcher.PreparationRequired,
		RuntimeOwnedRequest:      true,
		RuntimeOwnedLaunch:       launcher.RuntimeOwnedLaunch,
		KDEPresentationOnly:      launcher.KDEPresentationOnly,
		DryRun:                   true,
		ExecutionStarted:         false,
		BackendProcessStarted:    false,
		HostRootModified:         launcher.HostRootModified,
		HostNetworkingRequired:   launcher.HostNetworkingRequired,
		DockerSocketMounted:      launcher.DockerSocketMounted,
		BroadHostMountRequired:   launcher.BroadHostMountRequired,
		RawHostPathExposed:       launcher.RawHostPathExposed,
		RawExecutablePathExposed: launcher.RawExecutablePathExposed,
		RawCommandExposed:        launcher.RawCommandExposed,
		BackendDetailsExposed:    launcher.BackendDetailsExposed,
		DesktopSafeSummary:       launcher.DesktopSafeSummary,
		BlockedReason:            launcher.BlockedReason,
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

func knownLaunchProfileStateRoot(stateRoot string, cacheRoot string, app KnownPortableApp) (string, error) {
	root := strings.TrimSpace(stateRoot)
	if root == "" {
		cache := strings.TrimSpace(cacheRoot)
		if cache == "" {
			cache = DefaultKnownAppCacheRoot
		}
		root = filepath.Join(cache, app.ID, "runtime-state")
	}
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve known app launch profile state root: %w", err)
	}
	return absoluteRoot, nil
}

func knownLaunchProfileOutputPath(profileOutput string, stateRoot string, app KnownPortableApp) (string, error) {
	path := strings.TrimSpace(profileOutput)
	if path == "" {
		path = filepath.Join(stateRoot, "profiles", app.ID+".windows-app-smoke-profile.json")
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve known app launch profile output: %w", err)
	}
	return absolutePath, nil
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

func managedLauncherArgvMatches(provided []string, expected []string) bool {
	if len(provided) == 0 {
		return true
	}
	if len(provided) != len(expected) {
		return false
	}
	for index, value := range provided {
		if value != expected[index] {
			return false
		}
	}
	return true
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
	request.Header.Set("User-Agent", KnownAppDownloadUserAgent)
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
