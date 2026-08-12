package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Q4StagedExternalWinAppAcceptanceSchemaVersion = "xnix.runtime.q4_staged_external_winapp_acceptance.v1"
	Q4StagedExternalWinAppAcceptanceRequestType   = "q4-staged-external-winapp-acceptance-preview"
)

type Q4StagedExternalWinAppAcceptanceRequest struct {
	SmokeReportPath string
}

type Q4StagedExternalWinAppAcceptance struct {
	Version                                                       string `json:"version"`
	SchemaVersion                                                 string `json:"schema_version"`
	RequestType                                                   string `json:"request_type"`
	Source                                                        string `json:"source"`
	RuntimeMethod                                                 string `json:"runtime_method"`
	ReadMethod                                                    string `json:"read_method"`
	AcceptanceType                                                string `json:"acceptance_type"`
	SmokeReportConsumed                                           bool   `json:"smoke_report_consumed"`
	SmokeReportPathExposed                                        bool   `json:"smoke_report_path_exposed"`
	OutputPathExposed                                             bool   `json:"output_path_exposed"`
	DelegatedCommandExposed                                       bool   `json:"delegated_command_exposed"`
	RemoteHostExposed                                             bool   `json:"remote_host_exposed"`
	AppID                                                         string `json:"app_id"`
	DisplayName                                                   string `json:"display_name"`
	Q4CompileRequired                                             bool   `json:"q4_compile_required"`
	HostCompilationAvoided                                        bool   `json:"host_compilation_avoided"`
	RemoteBuildCompleted                                          bool   `json:"remote_build_completed"`
	DelegatedSmokePassed                                          bool   `json:"delegated_smoke_passed"`
	DesktopExecUsesExternalAppHandle                              bool   `json:"desktop_exec_uses_external_app_handle"`
	DesktopExecInvocationExact                                    bool   `json:"desktop_exec_invocation_exact"`
	ExternalAppDesktopHandleReady                                 bool   `json:"external_app_desktop_handle_ready"`
	ExternalAppHandleConsumed                                     bool   `json:"external_app_handle_consumed"`
	ImportedArtifactDigestVerified                                bool   `json:"imported_artifact_digest_verified"`
	ExternalFileBridgeReady                                       bool   `json:"external_file_bridge_ready"`
	ExternalFileBridgeArgumentsPassed                             bool   `json:"external_file_bridge_arguments_passed"`
	ExternalFileBridgeWinepathTranslated                          bool   `json:"external_file_bridge_winepath_translated"`
	WindowsProcessFileArgumentWindowObserved                      bool   `json:"windows_process_file_argument_window_observed"`
	WindowObserved                                                bool   `json:"window_observed"`
	XWindowObserved                                               bool   `json:"x_window_observed"`
	OneShotRuntimeLaunchExecuted                                  bool   `json:"one_shot_runtime_launch_executed"`
	OneShotWindowObserved                                         bool   `json:"one_shot_window_observed"`
	RuntimePacketExternalAppRunRecordConsumed                     bool   `json:"runtime_packet_external_app_run_record_consumed"`
	KDEPageCardExternalAppRunRecordConsumed                       bool   `json:"kde_page_card_external_app_run_record_consumed"`
	CompatibilityEvidenceBundleGenerated                          bool   `json:"compatibility_evidence_bundle_generated"`
	CompatibilityEvidenceBundleRealWindowsAppRunVerified          bool   `json:"compatibility_evidence_bundle_real_windows_app_run_verified"`
	ApplicationDetailGenerated                                    bool   `json:"application_detail_generated"`
	ApplicationDetailRealWindowsAppRunVerified                    bool   `json:"application_detail_real_windows_app_run_verified"`
	ApplicationDetailFileOpenVerified                             bool   `json:"application_detail_file_open_verified"`
	KDEPageFromApplicationDetailGenerated                         bool   `json:"kde_page_from_application_detail_generated"`
	KDEPageFromApplicationDetailConsumed                          bool   `json:"kde_page_from_application_detail_consumed"`
	KDEPageFromApplicationDetailHeaderBadge                       string `json:"kde_page_from_application_detail_header_badge"`
	KDEPageFromApplicationDetailHeaderBadgeTone                   string `json:"kde_page_from_application_detail_header_badge_tone"`
	KDEPageFromApplicationDetailSummaryCompatibilityState         string `json:"kde_page_from_application_detail_summary_compatibility_state"`
	KDEPageFromApplicationDetailSummaryDiagnosticsState           string `json:"kde_page_from_application_detail_summary_diagnostics_state"`
	KDEPageFromApplicationDetailSummaryBackendLaunchEnabled       bool   `json:"kde_page_from_application_detail_summary_backend_launch_enabled"`
	KDEPageFromApplicationDetailSummaryActionExecutionEnabled     bool   `json:"kde_page_from_application_detail_summary_action_execution_enabled"`
	KDEPageFromApplicationDetailSummarySettingsPersistenceEnabled bool   `json:"kde_page_from_application_detail_summary_settings_persistence_enabled"`
	ArtifactFetchCount                                            int    `json:"artifact_fetch_count"`
	HostRootModified                                              bool   `json:"host_root_modified"`
	PrivilegedContainerRequired                                   bool   `json:"privileged_container_required"`
	HostNetworkingRequired                                        bool   `json:"host_networking_required"`
	DockerSocketMounted                                           bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                                        bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                                         bool   `json:"backend_details_exposed"`
	RawPathExposed                                                bool   `json:"raw_path_exposed"`
	AcceptanceReady                                               bool   `json:"acceptance_ready"`
	DesktopSafeSummary                                            string `json:"desktop_safe_summary"`
}

func PreviewQ4StagedExternalWinAppAcceptance(request Q4StagedExternalWinAppAcceptanceRequest) (Q4StagedExternalWinAppAcceptance, error) {
	path := strings.TrimSpace(request.SmokeReportPath)
	if path == "" {
		return Q4StagedExternalWinAppAcceptance{}, errors.New("q4 staged external Windows app acceptance requires --q4-staged-external-winapp-smoke")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Q4StagedExternalWinAppAcceptance{}, fmt.Errorf("read q4 staged external Windows app smoke report: %w", err)
	}
	return PreviewQ4StagedExternalWinAppAcceptanceJSON(content)
}

func PreviewQ4StagedExternalWinAppAcceptanceJSON(content []byte) (Q4StagedExternalWinAppAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return Q4StagedExternalWinAppAcceptance{}, fmt.Errorf("parse q4 staged external Windows app smoke report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.scripts.q4_staged_desktop_external_winapp_smoke.v1" ||
		remoteString(report, "request_type") != "q4-staged-desktop-external-winapp-smoke" {
		return Q4StagedExternalWinAppAcceptance{}, errors.New("q4 staged external Windows app acceptance requires q4 staged desktop external Windows app smoke output")
	}
	if remoteString(report, "status") != "passed" || !remoteBool(report, "execute") {
		return Q4StagedExternalWinAppAcceptance{}, errors.New("q4 staged external Windows app acceptance requires an executed passed smoke")
	}

	appID := remoteString(report, "app_id")
	displayName := remoteString(report, "display_name")
	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required") ||
		remoteBool(report, "external_file_bridge_mount_enabled") ||
		remoteBool(report, "raw_file_uri_arguments_exposed") ||
		remoteBool(report, "one_shot_host_root_modified") ||
		remoteBool(report, "one_shot_docker_socket_mounted") ||
		remoteBool(report, "one_shot_raw_paths_exposed") ||
		remoteBool(report, "kde_page_from_application_detail_header_backend_details_exposed") ||
		remoteBool(report, "kde_page_from_application_detail_summary_backend_details_exposed") ||
		remoteBool(report, "kde_page_from_application_detail_summary_host_root_modified")
	kdeVerifiedStateReady := remoteString(report, "kde_page_from_application_detail_header_badge") == "Verified real app run" &&
		remoteString(report, "kde_page_from_application_detail_header_badge_tone") == "success" &&
		remoteString(report, "kde_page_from_application_detail_summary_compatibility_state") == "real-app-run-verified" &&
		remoteString(report, "kde_page_from_application_detail_summary_compatibility_label") == "Verified real app run" &&
		remoteString(report, "kde_page_from_application_detail_summary_diagnostics_state") == "real-app-run-verified" &&
		!remoteBool(report, "kde_page_from_application_detail_summary_backend_launch_enabled") &&
		!remoteBool(report, "kde_page_from_application_detail_summary_action_execution_enabled") &&
		!remoteBool(report, "kde_page_from_application_detail_summary_settings_persistence_enabled")
	acceptanceReady := appID != "" &&
		displayName != "" &&
		remoteBool(report, "q4_compile_required") &&
		remoteBool(report, "host_compilation_avoided") &&
		remoteBool(report, "remote_build_completed") &&
		remoteString(report, "delegated_status") == "passed" &&
		remoteBool(report, "desktop_exec_uses_external_app_handle") &&
		remoteBool(report, "desktop_exec_invocation_exact") &&
		remoteBool(report, "external_app_desktop_handle_ready") &&
		remoteBool(report, "external_app_handle_consumed") &&
		remoteBool(report, "imported_artifact_digest_verified") &&
		remoteBool(report, "external_file_bridge_ready") &&
		remoteBool(report, "external_file_bridge_arguments_passed") &&
		remoteBool(report, "external_file_bridge_winepath_translated") &&
		remoteBool(report, "windows_process_file_argument_window_observed") &&
		remoteBool(report, "window_observed") &&
		remoteBool(report, "x_window_observed") &&
		remoteBool(report, "one_shot_runtime_launch_executed") &&
		remoteBool(report, "one_shot_window_observed") &&
		remoteBool(report, "runtime_packet_external_app_run_record_consumed") &&
		remoteBool(report, "kde_page_card_external_app_run_record_consumed") &&
		remoteBool(report, "compatibility_evidence_bundle_generated") &&
		remoteBool(report, "compatibility_evidence_bundle_real_windows_app_run_verified") &&
		remoteBool(report, "application_detail_generated") &&
		remoteBool(report, "application_detail_real_windows_app_run_verified") &&
		remoteBool(report, "application_detail_file_open_verified") &&
		remoteBool(report, "kde_page_from_application_detail_generated") &&
		remoteBool(report, "kde_page_from_application_detail_consumed") &&
		kdeVerifiedStateReady &&
		remoteInt(report, "artifact_fetch_count") >= 10 &&
		remoteString(report, "container_network_mode") == "none" &&
		remoteInt(report, "container_host_mount_count") == 0 &&
		!unsafeGateOpen
	if !acceptanceReady {
		return Q4StagedExternalWinAppAcceptance{}, errors.New("q4 staged external Windows app acceptance requires handle-only desktop launch, real file-open GUI evidence, Runtime/KDE detail consumption, q4 compilation, fetched artifacts, and closed safety gates")
	}

	acceptance := Q4StagedExternalWinAppAcceptance{
		Version:                                   remoteString(report, "version"),
		SchemaVersion:                             Q4StagedExternalWinAppAcceptanceSchemaVersion,
		RequestType:                               Q4StagedExternalWinAppAcceptanceRequestType,
		Source:                                    "q4-staged-desktop-external-winapp-smoke+go-runtime-acceptance",
		RuntimeMethod:                             "PreviewQ4StagedExternalWinAppAcceptance",
		ReadMethod:                                "GetQ4StagedExternalWinAppAcceptance",
		AcceptanceType:                            "staged-external-winapp-desktop-real-run-acceptance",
		SmokeReportConsumed:                       true,
		SmokeReportPathExposed:                    false,
		OutputPathExposed:                         false,
		DelegatedCommandExposed:                   false,
		RemoteHostExposed:                         false,
		AppID:                                     appID,
		DisplayName:                               displayName,
		Q4CompileRequired:                         true,
		HostCompilationAvoided:                    true,
		RemoteBuildCompleted:                      true,
		DelegatedSmokePassed:                      true,
		DesktopExecUsesExternalAppHandle:          true,
		DesktopExecInvocationExact:                true,
		ExternalAppDesktopHandleReady:             true,
		ExternalAppHandleConsumed:                 true,
		ImportedArtifactDigestVerified:            true,
		ExternalFileBridgeReady:                   true,
		ExternalFileBridgeArgumentsPassed:         true,
		ExternalFileBridgeWinepathTranslated:      true,
		WindowsProcessFileArgumentWindowObserved:  true,
		WindowObserved:                            true,
		XWindowObserved:                           true,
		OneShotRuntimeLaunchExecuted:              true,
		OneShotWindowObserved:                     true,
		RuntimePacketExternalAppRunRecordConsumed: true,
		KDEPageCardExternalAppRunRecordConsumed:   true,
		CompatibilityEvidenceBundleGenerated:      true,
		CompatibilityEvidenceBundleRealWindowsAppRunVerified:          true,
		ApplicationDetailGenerated:                                    true,
		ApplicationDetailRealWindowsAppRunVerified:                    true,
		ApplicationDetailFileOpenVerified:                             true,
		KDEPageFromApplicationDetailGenerated:                         true,
		KDEPageFromApplicationDetailConsumed:                          true,
		KDEPageFromApplicationDetailHeaderBadge:                       "Verified real app run",
		KDEPageFromApplicationDetailHeaderBadgeTone:                   "success",
		KDEPageFromApplicationDetailSummaryCompatibilityState:         "real-app-run-verified",
		KDEPageFromApplicationDetailSummaryDiagnosticsState:           "real-app-run-verified",
		KDEPageFromApplicationDetailSummaryBackendLaunchEnabled:       false,
		KDEPageFromApplicationDetailSummaryActionExecutionEnabled:     false,
		KDEPageFromApplicationDetailSummarySettingsPersistenceEnabled: false,
		ArtifactFetchCount:                                            remoteInt(report, "artifact_fetch_count"),
		HostRootModified:                                              false,
		PrivilegedContainerRequired:                                   false,
		HostNetworkingRequired:                                        false,
		DockerSocketMounted:                                           false,
		BroadHostMountRequired:                                        false,
		BackendDetailsExposed:                                         false,
		RawPathExposed:                                                false,
		AcceptanceReady:                                               true,
		DesktopSafeSummary:                                            "A q4 staged external Windows GUI app completed the handle-only desktop launch, file-open, Runtime detail, and KDE detail acceptance lane.",
	}
	if err := validateNoBackendTerms(acceptance, "q4 staged external Windows app acceptance"); err != nil {
		return Q4StagedExternalWinAppAcceptance{}, err
	}
	return acceptance, nil
}
