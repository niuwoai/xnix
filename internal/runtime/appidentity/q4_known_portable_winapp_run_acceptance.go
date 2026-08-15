package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Q4KnownPortableWinAppRunAcceptanceSchemaVersion = "xnix.runtime.q4_known_portable_winapp_run_acceptance.v1"
	Q4KnownPortableWinAppRunAcceptanceRequestType   = "q4-known-portable-winapp-run-acceptance-preview"
)

type Q4KnownPortableWinAppRunAcceptanceRequest struct {
	RunReportPath string
}

type Q4KnownPortableWinAppRunAcceptance struct {
	Version                              string `json:"version"`
	SchemaVersion                        string `json:"schema_version"`
	RequestType                          string `json:"request_type"`
	Source                               string `json:"source"`
	RuntimeMethod                        string `json:"runtime_method"`
	ReadMethod                           string `json:"read_method"`
	AcceptanceType                       string `json:"acceptance_type"`
	RunReportConsumed                    bool   `json:"run_report_consumed"`
	RunReportPathExposed                 bool   `json:"run_report_path_exposed"`
	RemoteHostExposed                    bool   `json:"remote_host_exposed"`
	DelegatedCommandExposed              bool   `json:"delegated_command_exposed"`
	RawPathExposed                       bool   `json:"raw_path_exposed"`
	AppID                                string `json:"app_id"`
	DisplayName                          string `json:"display_name"`
	KnownCatalogApp                      bool   `json:"known_catalog_app"`
	PortableDirectoryExternalApp         bool   `json:"portable_directory_external_app"`
	SingleFileExternalApp                bool   `json:"single_file_external_app"`
	OfficialArchiveChecksumVerified      bool   `json:"official_archive_checksum_verified"`
	OfficialExecutableChecksumVerified   bool   `json:"official_executable_checksum_verified"`
	KnownPortableBundleImported          bool   `json:"known_portable_bundle_imported"`
	KnownPortableBundleStageLaunchStatus string `json:"known_portable_bundle_stage_launch_status"`
	KnownPortableBundleAcceptanceReady   bool   `json:"known_portable_bundle_acceptance_ready"`
	StagedExternalWinAppAcceptanceReady  bool   `json:"staged_external_winapp_acceptance_ready"`
	ExternalFileOpenRequested            bool   `json:"external_file_open_requested"`
	ExternalFileBridgeReady              bool   `json:"external_file_bridge_ready"`
	WindowsProcessFileArgumentObserved   bool   `json:"windows_process_file_argument_window_observed"`
	WindowObserved                       bool   `json:"window_observed"`
	XWindowObserved                      bool   `json:"x_window_observed"`
	RuntimeAcceptedChainVerified         bool   `json:"runtime_accepted_chain_verified"`
	AcceptedApplicationDetailState       string `json:"accepted_application_detail_state"`
	KDEAcceptedPageState                 string `json:"kde_accepted_page_state"`
	OperatorRunReady                     bool   `json:"operator_run_ready"`
	DelegatedArtifactFetchCount          int    `json:"delegated_artifact_fetch_count"`
	Q4DownloadRequired                   bool   `json:"q4_download_required"`
	Q4ExtractRequired                    bool   `json:"q4_extract_required"`
	Q4CompileRequired                    bool   `json:"q4_compile_required"`
	Q4ExecutionRequired                  bool   `json:"q4_execution_required"`
	HostCompilationAvoided               bool   `json:"host_compilation_avoided"`
	HostDownloadAvoided                  bool   `json:"host_download_avoided"`
	HostRootModified                     bool   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool   `json:"privileged_container_required"`
	HostNetworkingRequired               bool   `json:"host_networking_required"`
	DockerSocketMounted                  bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                bool   `json:"backend_details_exposed"`
	AcceptanceReady                      bool   `json:"acceptance_ready"`
	DesktopSafeSummary                   string `json:"desktop_safe_summary"`
}

func PreviewQ4KnownPortableWinAppRunAcceptance(request Q4KnownPortableWinAppRunAcceptanceRequest) (Q4KnownPortableWinAppRunAcceptance, error) {
	path := strings.TrimSpace(request.RunReportPath)
	if path == "" {
		return Q4KnownPortableWinAppRunAcceptance{}, errors.New("q4 known portable Windows app run acceptance requires --q4-known-portable-winapp-run")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Q4KnownPortableWinAppRunAcceptance{}, fmt.Errorf("read q4 known portable Windows app run report: %w", err)
	}
	return PreviewQ4KnownPortableWinAppRunAcceptanceJSON(content)
}

func PreviewQ4KnownPortableWinAppRunAcceptanceJSON(content []byte) (Q4KnownPortableWinAppRunAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return Q4KnownPortableWinAppRunAcceptance{}, fmt.Errorf("parse q4 known portable Windows app run report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.scripts.q4_known_portable_winapp_run.v1" ||
		remoteString(report, "request_type") != "q4-known-portable-winapp-run" {
		return Q4KnownPortableWinAppRunAcceptance{}, errors.New("q4 known portable Windows app run acceptance requires q4 known portable Windows app run output")
	}
	if remoteString(report, "status") != "passed" || !remoteBool(report, "execute") {
		return Q4KnownPortableWinAppRunAcceptance{}, errors.New("q4 known portable Windows app run acceptance requires an executed passed run")
	}
	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required")
	appID := remoteString(report, "app_id")
	displayName := remoteString(report, "display_name")
	bundleReady := appID == "org.xnix.external.notepadplusplus" &&
		displayName == "Notepad++ Portable" &&
		remoteBool(report, "known_catalog_app") &&
		remoteBool(report, "portable_directory_external_app") &&
		!remoteBool(report, "single_file_external_app") &&
		remoteBool(report, "official_download_required") &&
		remoteBool(report, "pinned_checksum_required") &&
		remoteBool(report, "official_archive_checksum_verified") &&
		remoteBool(report, "known_portable_bundle_imported") &&
		remoteString(report, "known_portable_bundle_stage_launch_status") == "passed" &&
		remoteString(report, "known_portable_bundle_acceptance_request_type") == Q4KnownPortableBundleWinAppAcceptanceRequestType &&
		remoteBool(report, "known_portable_bundle_acceptance_ready") &&
		remoteBool(report, "runtime_accepted_chain_verified") &&
		remoteString(report, "accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
		remoteString(report, "kde_accepted_page_state") == "runtime-accepted-real-app-run" &&
		remoteBool(report, "operator_run_ready") &&
		remoteInt(report, "delegated_artifact_fetch_count") >= 13 &&
		remoteBool(report, "q4_download_required") &&
		remoteBool(report, "q4_extract_required") &&
		remoteBool(report, "q4_compile_required") &&
		remoteBool(report, "q4_execution_required") &&
		remoteBool(report, "host_compilation_avoided") &&
		remoteBool(report, "host_download_avoided") &&
		!unsafeGateOpen
	singleExecutableReady := appID == "org.xnix.external.putty" &&
		displayName == "PuTTY" &&
		remoteBool(report, "known_catalog_app") &&
		!remoteBool(report, "portable_directory_external_app") &&
		remoteBool(report, "single_file_external_app") &&
		remoteBool(report, "official_download_required") &&
		remoteBool(report, "pinned_checksum_required") &&
		remoteBool(report, "official_executable_checksum_verified") &&
		!remoteBool(report, "known_portable_bundle_imported") &&
		remoteString(report, "staged_external_winapp_acceptance_request_type") == Q4StagedExternalWinAppAcceptanceRequestType &&
		remoteBool(report, "staged_external_winapp_acceptance_ready") &&
		!remoteBool(report, "external_file_open_requested") &&
		!remoteBool(report, "external_file_bridge_ready") &&
		!remoteBool(report, "windows_process_file_argument_window_observed") &&
		remoteBool(report, "window_observed") &&
		remoteBool(report, "x_window_observed") &&
		remoteBool(report, "runtime_accepted_chain_verified") &&
		remoteString(report, "accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
		remoteString(report, "kde_accepted_page_state") == "runtime-accepted-real-app-run" &&
		remoteBool(report, "operator_run_ready") &&
		remoteInt(report, "delegated_artifact_fetch_count") >= 13 &&
		remoteBool(report, "q4_download_required") &&
		!remoteBool(report, "q4_extract_required") &&
		remoteBool(report, "q4_compile_required") &&
		remoteBool(report, "q4_execution_required") &&
		remoteBool(report, "host_compilation_avoided") &&
		remoteBool(report, "host_download_avoided") &&
		!unsafeGateOpen
	if !bundleReady && !singleExecutableReady {
		return Q4KnownPortableWinAppRunAcceptance{}, errors.New("q4 known portable Windows app run acceptance requires a passed catalog-backed bundle or single-executable operator run, Go-owned acceptance, accepted Runtime/KDE state, fetched artifacts, and closed safety gates")
	}
	acceptance := Q4KnownPortableWinAppRunAcceptance{
		Version:                              remoteString(report, "version"),
		SchemaVersion:                        Q4KnownPortableWinAppRunAcceptanceSchemaVersion,
		RequestType:                          Q4KnownPortableWinAppRunAcceptanceRequestType,
		Source:                               "q4-known-portable-winapp-run+go-runtime-operator-acceptance",
		RuntimeMethod:                        "PreviewQ4KnownPortableWinAppRunAcceptance",
		ReadMethod:                           "GetQ4KnownPortableWinAppRunAcceptance",
		AcceptanceType:                       "q4-known-portable-real-winapp-operator-run-acceptance",
		RunReportConsumed:                    true,
		RunReportPathExposed:                 false,
		RemoteHostExposed:                    false,
		DelegatedCommandExposed:              false,
		RawPathExposed:                       false,
		AppID:                                appID,
		DisplayName:                          displayName,
		KnownCatalogApp:                      true,
		PortableDirectoryExternalApp:         bundleReady,
		SingleFileExternalApp:                singleExecutableReady,
		OfficialArchiveChecksumVerified:      bundleReady,
		OfficialExecutableChecksumVerified:   singleExecutableReady,
		KnownPortableBundleImported:          bundleReady,
		KnownPortableBundleStageLaunchStatus: remoteString(report, "known_portable_bundle_stage_launch_status"),
		KnownPortableBundleAcceptanceReady:   bundleReady,
		StagedExternalWinAppAcceptanceReady:  singleExecutableReady,
		ExternalFileOpenRequested:            remoteBool(report, "external_file_open_requested"),
		ExternalFileBridgeReady:              remoteBool(report, "external_file_bridge_ready"),
		WindowsProcessFileArgumentObserved:   remoteBool(report, "windows_process_file_argument_window_observed"),
		WindowObserved:                       remoteBool(report, "window_observed"),
		XWindowObserved:                      remoteBool(report, "x_window_observed"),
		RuntimeAcceptedChainVerified:         true,
		AcceptedApplicationDetailState:       "runtime-accepted-real-app-run",
		KDEAcceptedPageState:                 "runtime-accepted-real-app-run",
		OperatorRunReady:                     true,
		DelegatedArtifactFetchCount:          remoteInt(report, "delegated_artifact_fetch_count"),
		Q4DownloadRequired:                   true,
		Q4ExtractRequired:                    bundleReady,
		Q4CompileRequired:                    true,
		Q4ExecutionRequired:                  true,
		HostCompilationAvoided:               true,
		HostDownloadAvoided:                  true,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		BackendDetailsExposed:                false,
		AcceptanceReady:                      true,
		DesktopSafeSummary:                   displayName + " completed the q4 known portable Windows app operator run with Runtime-owned acceptance and KDE accepted-state evidence.",
	}
	if err := validateNoBackendTerms(acceptance, "q4 known portable Windows app run acceptance"); err != nil {
		return Q4KnownPortableWinAppRunAcceptance{}, err
	}
	return acceptance, nil
}
