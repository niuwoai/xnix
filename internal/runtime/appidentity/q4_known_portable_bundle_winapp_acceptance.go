package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Q4KnownPortableBundleWinAppAcceptanceSchemaVersion = "xnix.runtime.q4_known_portable_bundle_winapp_acceptance.v1"
	Q4KnownPortableBundleWinAppAcceptanceRequestType   = "q4-known-portable-bundle-winapp-acceptance-preview"
)

type Q4KnownPortableBundleWinAppAcceptanceRequest struct {
	SmokeReportPath string
}

type Q4KnownPortableBundleWinAppAcceptance struct {
	Version                                string `json:"version"`
	SchemaVersion                          string `json:"schema_version"`
	RequestType                            string `json:"request_type"`
	Source                                 string `json:"source"`
	RuntimeMethod                          string `json:"runtime_method"`
	ReadMethod                             string `json:"read_method"`
	AcceptanceType                         string `json:"acceptance_type"`
	SmokeReportConsumed                    bool   `json:"smoke_report_consumed"`
	SmokeReportPathExposed                 bool   `json:"smoke_report_path_exposed"`
	RemoteHostExposed                      bool   `json:"remote_host_exposed"`
	DelegatedCommandExposed                bool   `json:"delegated_command_exposed"`
	RawPathExposed                         bool   `json:"raw_path_exposed"`
	AppID                                  string `json:"app_id"`
	DisplayName                            string `json:"display_name"`
	RealThirdPartyWindowsApp               bool   `json:"real_third_party_windows_app"`
	PortableDirectoryApp                   bool   `json:"portable_directory_app"`
	Q4CompileRequired                      bool   `json:"q4_compile_required"`
	HostCompilationAvoided                 bool   `json:"host_compilation_avoided"`
	HostDownloadAvoided                    bool   `json:"host_download_avoided"`
	OfficialArchiveChecksumVerified        bool   `json:"official_archive_checksum_verified"`
	KnownBundleImportVerified              bool   `json:"known_bundle_import_verified"`
	KnownBundleStageLaunchVerified         bool   `json:"known_bundle_stage_launch_verified"`
	KnownBundleGUIEvidencePacketVerified   bool   `json:"known_bundle_gui_evidence_packet_verified"`
	KnownBundleKDEPageVerified             bool   `json:"known_bundle_kde_page_verified"`
	KnownBundleCompatibilityBundleVerified bool   `json:"known_bundle_compatibility_bundle_verified"`
	KnownBundleApplicationDetailVerified   bool   `json:"known_bundle_application_detail_verified"`
	KnownBundleKDEPageFromDetailVerified   bool   `json:"known_bundle_kde_page_from_detail_verified"`
	LaunchSourceRequestType                string `json:"launch_source_request_type"`
	KnownPortableBundleStageLaunchConsumed bool   `json:"known_portable_bundle_stage_launch_consumed"`
	RuntimeAcceptedChainVerified           bool   `json:"runtime_accepted_chain_verified"`
	RuntimeAcceptedApplicationDetailState  string `json:"runtime_accepted_application_detail_state"`
	RuntimeAcceptedKDEPageState            string `json:"runtime_accepted_kde_page_state"`
	ArtifactFetchCount                     int    `json:"artifact_fetch_count"`
	HostRootModified                       bool   `json:"host_root_modified"`
	PrivilegedContainerRequired            bool   `json:"privileged_container_required"`
	HostNetworkingRequired                 bool   `json:"host_networking_required"`
	DockerSocketMounted                    bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                 bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                  bool   `json:"backend_details_exposed"`
	AcceptanceReady                        bool   `json:"acceptance_ready"`
	DesktopSafeSummary                     string `json:"desktop_safe_summary"`
}

func PreviewQ4KnownPortableBundleWinAppAcceptance(request Q4KnownPortableBundleWinAppAcceptanceRequest) (Q4KnownPortableBundleWinAppAcceptance, error) {
	path := strings.TrimSpace(request.SmokeReportPath)
	if path == "" {
		return Q4KnownPortableBundleWinAppAcceptance{}, errors.New("q4 known portable bundle Windows app acceptance requires --q4-notepadpp-portable-winapp-smoke")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Q4KnownPortableBundleWinAppAcceptance{}, fmt.Errorf("read q4 known portable bundle Windows app smoke report: %w", err)
	}
	return PreviewQ4KnownPortableBundleWinAppAcceptanceJSON(content)
}

func PreviewQ4KnownPortableBundleWinAppAcceptanceJSON(content []byte) (Q4KnownPortableBundleWinAppAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return Q4KnownPortableBundleWinAppAcceptance{}, fmt.Errorf("parse q4 known portable bundle Windows app smoke report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.scripts.q4_notepadpp_portable_winapp_smoke.v1" ||
		remoteString(report, "request_type") != "q4-notepadpp-portable-winapp-smoke" {
		return Q4KnownPortableBundleWinAppAcceptance{}, errors.New("q4 known portable bundle Windows app acceptance requires q4 Notepad++ Portable smoke output")
	}
	if remoteString(report, "status") != "passed" || !remoteBool(report, "execute") {
		return Q4KnownPortableBundleWinAppAcceptance{}, errors.New("q4 known portable bundle Windows app acceptance requires an executed passed smoke")
	}
	appID := remoteString(report, "app_id")
	displayName := remoteString(report, "display_name")
	launchSource := remoteString(report, "known_portable_bundle_application_detail_launch_source_request_type")
	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required")
	knownBundleDirectPathReady := remoteString(report, "source_kind") == "official-notepad-plus-plus-github-release" &&
		appID == "org.xnix.external.notepadplusplus" &&
		displayName == "Notepad++ Portable" &&
		remoteBool(report, "real_third_party_windows_app") &&
		!remoteBool(report, "single_file_windows_app") &&
		remoteBool(report, "portable_directory_bundle_import_required") &&
		remoteBool(report, "notepadpp_sha256_verified") &&
		remoteBool(report, "notepadpp_extracted") &&
		remoteString(report, "known_portable_bundle_import_status") == "passed" &&
		remoteBool(report, "known_portable_bundle_import_recorded") &&
		remoteBool(report, "known_portable_bundle_archive_verified") &&
		remoteBool(report, "known_portable_bundle_checksum_verified") &&
		remoteString(report, "known_portable_bundle_stage_launch_status") == "passed" &&
		remoteBool(report, "known_portable_bundle_stage_launch_record_first") &&
		remoteBool(report, "known_portable_bundle_stage_launch_existing_import_record_consumed") &&
		remoteBool(report, "known_portable_bundle_stage_launch_application_workspace_copied") &&
		remoteBool(report, "known_portable_bundle_stage_launch_windows_process_file_argument_window_observed")
	productReadModelsReady := remoteString(report, "known_portable_bundle_gui_evidence_packet_request_type") == RealWinAppGUIEvidencePacketRequestType &&
		remoteBool(report, "known_portable_bundle_gui_evidence_packet_external_app_run_record_consumed") &&
		remoteBool(report, "known_portable_bundle_gui_evidence_packet_imported_artifact_digest_verified") &&
		remoteString(report, "known_portable_bundle_kde_page_request_type") == "kde-center-page-preview" &&
		remoteString(report, "known_portable_bundle_compatibility_bundle_request_type") == ExternalWinAppCompatibilityEvidenceBundleRequestType &&
		remoteString(report, "known_portable_bundle_compatibility_bundle_launch_source_request_type") == "windows-known-app-bundle-stage-and-launch" &&
		remoteBool(report, "known_portable_bundle_compatibility_bundle_stage_launch_consumed") &&
		remoteBool(report, "known_portable_bundle_compatibility_bundle_runtime_gui_evidence_packet_verified") &&
		remoteBool(report, "known_portable_bundle_compatibility_bundle_kde_external_app_page_verified") &&
		remoteBool(report, "known_portable_bundle_compatibility_bundle_real_windows_app_run_verified") &&
		remoteString(report, "known_portable_bundle_application_detail_request_type") == ExternalWinAppApplicationDetailRequestType &&
		launchSource == "windows-known-app-bundle-stage-and-launch" &&
		remoteBool(report, "known_portable_bundle_application_detail_stage_launch_consumed") &&
		remoteBool(report, "known_portable_bundle_application_detail_real_windows_app_run_verified") &&
		remoteBool(report, "known_portable_bundle_application_detail_file_open_verified") &&
		remoteString(report, "known_portable_bundle_kde_page_from_detail_request_type") == "kde-center-page-preview" &&
		remoteString(report, "known_portable_bundle_kde_page_from_detail_launch_source_request_type") == "windows-known-app-bundle-stage-and-launch" &&
		remoteBool(report, "known_portable_bundle_kde_page_from_detail_stage_launch_consumed")
	fetchedArtifactsReady := remoteBool(report, "known_portable_bundle_gui_evidence_packet_artifact_fetched") &&
		remoteBool(report, "known_portable_bundle_kde_page_artifact_fetched") &&
		remoteBool(report, "known_portable_bundle_compatibility_bundle_artifact_fetched") &&
		remoteBool(report, "known_portable_bundle_application_detail_artifact_fetched") &&
		remoteBool(report, "known_portable_bundle_kde_page_from_detail_artifact_fetched")
	acceptedChainReady := remoteBool(report, "acceptance_ready") &&
		remoteString(report, "accepted_application_detail_state") == "runtime-accepted-real-app-run" &&
		remoteString(report, "kde_accepted_page_state") == "runtime-accepted-real-app-run"
	acceptanceReady := knownBundleDirectPathReady &&
		productReadModelsReady &&
		fetchedArtifactsReady &&
		acceptedChainReady &&
		remoteBool(report, "q4_compile_required") &&
		remoteBool(report, "q4_download_required") &&
		remoteBool(report, "q4_extract_required") &&
		remoteBool(report, "host_compilation_avoided") &&
		remoteBool(report, "host_download_avoided") &&
		remoteInt(report, "artifact_fetch_count") >= 13 &&
		!unsafeGateOpen
	if !acceptanceReady {
		return Q4KnownPortableBundleWinAppAcceptance{}, errors.New("q4 known portable bundle Windows app acceptance requires verified Notepad++ download, import, staged launch, Runtime detail, KDE detail, fetched artifacts, accepted chain, and closed safety gates")
	}
	acceptance := Q4KnownPortableBundleWinAppAcceptance{
		Version:                                remoteString(report, "version"),
		SchemaVersion:                          Q4KnownPortableBundleWinAppAcceptanceSchemaVersion,
		RequestType:                            Q4KnownPortableBundleWinAppAcceptanceRequestType,
		Source:                                 "q4-notepadpp-portable-winapp-smoke+go-runtime-known-bundle-acceptance",
		RuntimeMethod:                          "PreviewQ4KnownPortableBundleWinAppAcceptance",
		ReadMethod:                             "GetQ4KnownPortableBundleWinAppAcceptance",
		AcceptanceType:                         "q4-known-portable-bundle-real-winapp-acceptance",
		SmokeReportConsumed:                    true,
		SmokeReportPathExposed:                 false,
		RemoteHostExposed:                      false,
		DelegatedCommandExposed:                false,
		RawPathExposed:                         false,
		AppID:                                  appID,
		DisplayName:                            displayName,
		RealThirdPartyWindowsApp:               true,
		PortableDirectoryApp:                   true,
		Q4CompileRequired:                      true,
		HostCompilationAvoided:                 true,
		HostDownloadAvoided:                    true,
		OfficialArchiveChecksumVerified:        true,
		KnownBundleImportVerified:              true,
		KnownBundleStageLaunchVerified:         true,
		KnownBundleGUIEvidencePacketVerified:   true,
		KnownBundleKDEPageVerified:             true,
		KnownBundleCompatibilityBundleVerified: true,
		KnownBundleApplicationDetailVerified:   true,
		KnownBundleKDEPageFromDetailVerified:   true,
		LaunchSourceRequestType:                launchSource,
		KnownPortableBundleStageLaunchConsumed: true,
		RuntimeAcceptedChainVerified:           true,
		RuntimeAcceptedApplicationDetailState:  "runtime-accepted-real-app-run",
		RuntimeAcceptedKDEPageState:            "runtime-accepted-real-app-run",
		ArtifactFetchCount:                     remoteInt(report, "artifact_fetch_count"),
		HostRootModified:                       false,
		PrivilegedContainerRequired:            false,
		HostNetworkingRequired:                 false,
		DockerSocketMounted:                    false,
		BroadHostMountRequired:                 false,
		BackendDetailsExposed:                  false,
		AcceptanceReady:                        true,
		DesktopSafeSummary:                     "Notepad++ Portable completed the q4 known portable bundle Runtime acceptance lane through direct Go-owned launch, detail, KDE, and accepted-state evidence.",
	}
	if err := validateNoBackendTerms(acceptance, "q4 known portable bundle Windows app acceptance"); err != nil {
		return Q4KnownPortableBundleWinAppAcceptance{}, err
	}
	return acceptance, nil
}
