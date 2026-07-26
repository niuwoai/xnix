package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	RealWinAppRunAcceptanceSchemaVersion = "xnix.runtime.real_winapp_run_acceptance.v1"
	RealWinAppRunAcceptanceRequestType   = "real-winapp-run-acceptance-preview"
)

type RealWinAppRunAcceptanceRequest struct {
	RemoteExecuteResultPath string
}

type RealWinAppRunAcceptance struct {
	Version                                       string `json:"version"`
	SchemaVersion                                 string `json:"schema_version"`
	RequestType                                   string `json:"request_type"`
	Source                                        string `json:"source"`
	RuntimeMethod                                 string `json:"runtime_method"`
	ReadMethod                                    string `json:"read_method"`
	AcceptanceType                                string `json:"acceptance_type"`
	ExecuteResultConsumed                         bool   `json:"execute_result_consumed"`
	ExecuteResultPathExposed                      bool   `json:"execute_result_path_exposed"`
	ReceiptSummaryPathExposed                     bool   `json:"receipt_summary_path_exposed"`
	CenterProjectionPathExposed                   bool   `json:"center_projection_path_exposed"`
	KDEPageProjectionPathExposed                  bool   `json:"kde_page_projection_path_exposed"`
	RemoteHostExposed                             bool   `json:"remote_host_exposed"`
	AppID                                         string `json:"app_id"`
	DisplayName                                   string `json:"display_name"`
	AppVersion                                    string `json:"app_version"`
	LaunchMode                                    string `json:"launch_mode"`
	RunPassed                                     bool   `json:"run_passed"`
	RealExecutionObserved                         bool   `json:"real_execution_observed"`
	WindowObserved                                bool   `json:"window_observed"`
	WindowMatchObserved                           bool   `json:"window_match_observed"`
	OwnerControlledFileOpenVerified               bool   `json:"owner_controlled_file_open_verified"`
	OwnerFileOpenEntrypointInvoked                bool   `json:"owner_file_open_entrypoint_invoked"`
	RuntimeEvidenceConsumed                       bool   `json:"runtime_evidence_consumed"`
	ReceiptSummaryGenerated                       bool   `json:"receipt_summary_generated"`
	ReceiptSummaryReady                           bool   `json:"receipt_summary_ready"`
	ReceiptSummaryFileOpenVerified                bool   `json:"receipt_summary_file_open_verified"`
	CompatibilityCenterProjectionConsumed         bool   `json:"compatibility_center_projection_consumed"`
	CompatibilityCenterKnownAppSmokeEvidenceCount int    `json:"compatibility_center_known_app_smoke_evidence_count"`
	KDEPageProjectionConsumed                     bool   `json:"kde_page_projection_consumed"`
	KDEPageKnownAppGUIEvidenceCount               int    `json:"kde_page_known_app_gui_evidence_count"`
	KDEPageOwnerFileOpenVerifiedCount             int    `json:"kde_page_owner_file_open_verified_count"`
	KDEPageOwnerFileOpenEntrypointCount           int    `json:"kde_page_owner_file_open_entrypoint_count"`
	HostRootModified                              bool   `json:"host_root_modified"`
	PrivilegedContainerRequired                   bool   `json:"privileged_container_required"`
	HostNetworkingRequired                        bool   `json:"host_networking_required"`
	DockerSocketMounted                           bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                        bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                         bool   `json:"backend_details_exposed"`
	RawLauncherOutputExposed                      bool   `json:"raw_launcher_output_exposed"`
	RawWindowEvidenceExposed                      bool   `json:"raw_window_evidence_exposed"`
	ExecutableNameExposed                         bool   `json:"executable_name_exposed"`
	AcceptanceReady                               bool   `json:"acceptance_ready"`
	DesktopSafeSummary                            string `json:"desktop_safe_summary"`
}

func PreviewRealWinAppRunAcceptance(request RealWinAppRunAcceptanceRequest) (RealWinAppRunAcceptance, error) {
	path := strings.TrimSpace(request.RemoteExecuteResultPath)
	if path == "" {
		return RealWinAppRunAcceptance{}, errors.New("real Windows app run acceptance requires --remote-execute-result")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return RealWinAppRunAcceptance{}, fmt.Errorf("read real Windows app remote execute-result: %w", err)
	}
	return PreviewRealWinAppRunAcceptanceJSON(content)
}

func PreviewRealWinAppRunAcceptanceJSON(content []byte) (RealWinAppRunAcceptance, error) {
	var result map[string]any
	if err := json.Unmarshal(content, &result); err != nil {
		return RealWinAppRunAcceptance{}, fmt.Errorf("parse real Windows app remote execute-result: %w", err)
	}
	if remoteString(result, "schema_version") != "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1" ||
		remoteString(result, "request_type") != "remote-wine-guest-gui-smoke" {
		return RealWinAppRunAcceptance{}, errors.New("real Windows app run acceptance requires a remote GUI smoke execute-result")
	}
	if remoteString(result, "status") != "passed" || remoteString(result, "smoke_status") != "passed" {
		return RealWinAppRunAcceptance{}, errors.New("real Windows app run acceptance requires a passed remote GUI smoke")
	}
	appID := firstRemoteString(result, "known_app_id", "runtime_evidence_app_id", "kde_page_app_id")
	if appID == "" {
		return RealWinAppRunAcceptance{}, errors.New("real Windows app run acceptance requires safe app identity")
	}
	displayName := firstRemoteString(result, "known_app_name", "runtime_evidence_display_name", "kde_page_display_name")
	appVersion := firstRemoteString(result, "known_app_version", "runtime_evidence_app_version")
	if displayName == "" {
		displayName = appID
	}

	unsafeGateOpen := remoteBool(result, "host_root_modified") ||
		remoteBool(result, "privileged_container_required") ||
		remoteBool(result, "host_networking_required") ||
		remoteBool(result, "docker_socket_mounted") ||
		remoteBool(result, "broad_host_mount_required")
	realExecutionObserved := remoteBool(result, "x_window_observed") && remoteBool(result, "runtime_evidence_window_observed")
	ownerFileOpenVerified := remoteBool(result, "owner_controlled_launch_requested") &&
		remoteBool(result, "file_open_entrypoint_requested") &&
		remoteBool(result, "owner_file_open_entrypoint_invoked") &&
		remoteBool(result, "runtime_evidence_owner_file_open_verified") &&
		remoteBool(result, "kde_action_owner_file_open_verified") &&
		remoteBool(result, "owner_delegated_file_arguments_passed") &&
		remoteBool(result, "owner_delegated_file_argument_winepath_translated") &&
		!remoteBool(result, "owner_delegated_raw_file_argument_path_exposed")
	receiptSummaryGenerated := remoteBool(result, "execute_result_output_written") &&
		remoteBool(result, "real_run_receipt_summary_preview_planned") &&
		remoteBool(result, "real_run_receipt_summary_output_written")
	receiptSummaryReady := receiptSummaryGenerated &&
		remoteBool(result, "real_run_receipt_summary_ready") &&
		remoteBool(result, "real_run_receipt_summary_file_open_verified")
	centerConsumed := remoteBool(result, "real_run_receipt_summary_center_output_written") &&
		remoteInt(result, "real_run_receipt_summary_center_known_app_smoke_evidence_count") > 0
	kdePageConsumed := remoteBool(result, "real_run_receipt_summary_kde_page_output_written") &&
		remoteInt(result, "real_run_receipt_summary_kde_page_known_app_gui_evidence_count") > 0 &&
		remoteInt(result, "real_run_receipt_summary_kde_page_owner_file_open_verified_count") > 0 &&
		remoteInt(result, "real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count") > 0
	acceptanceReady := realExecutionObserved &&
		remoteBool(result, "window_match_observed") &&
		ownerFileOpenVerified &&
		remoteBool(result, "runtime_evidence_report_consumed") &&
		receiptSummaryReady &&
		centerConsumed &&
		kdePageConsumed &&
		!unsafeGateOpen
	if !acceptanceReady {
		return RealWinAppRunAcceptance{}, errors.New("real Windows app run acceptance requires observed execution, owner file-open entrypoint evidence, ready receipt summary, desktop projections, and closed safety gates")
	}

	acceptance := RealWinAppRunAcceptance{
		Version:                               remoteString(result, "version"),
		SchemaVersion:                         RealWinAppRunAcceptanceSchemaVersion,
		RequestType:                           RealWinAppRunAcceptanceRequestType,
		Source:                                "remote-gui-smoke+real-run-receipt-summary+desktop-projection",
		RuntimeMethod:                         "PreviewRealWinAppRunAcceptance",
		ReadMethod:                            "GetRealWinAppRunAcceptance",
		AcceptanceType:                        "real-windows-app-run-acceptance",
		ExecuteResultConsumed:                 true,
		ExecuteResultPathExposed:              false,
		ReceiptSummaryPathExposed:             false,
		CenterProjectionPathExposed:           false,
		KDEPageProjectionPathExposed:          false,
		RemoteHostExposed:                     false,
		AppID:                                 appID,
		DisplayName:                           displayName,
		AppVersion:                            appVersion,
		LaunchMode:                            remoteString(result, "launch_mode"),
		RunPassed:                             true,
		RealExecutionObserved:                 true,
		WindowObserved:                        true,
		WindowMatchObserved:                   true,
		OwnerControlledFileOpenVerified:       true,
		OwnerFileOpenEntrypointInvoked:        true,
		RuntimeEvidenceConsumed:               true,
		ReceiptSummaryGenerated:               true,
		ReceiptSummaryReady:                   true,
		ReceiptSummaryFileOpenVerified:        true,
		CompatibilityCenterProjectionConsumed: true,
		CompatibilityCenterKnownAppSmokeEvidenceCount: remoteInt(result, "real_run_receipt_summary_center_known_app_smoke_evidence_count"),
		KDEPageProjectionConsumed:                     true,
		KDEPageKnownAppGUIEvidenceCount:               remoteInt(result, "real_run_receipt_summary_kde_page_known_app_gui_evidence_count"),
		KDEPageOwnerFileOpenVerifiedCount:             remoteInt(result, "real_run_receipt_summary_kde_page_owner_file_open_verified_count"),
		KDEPageOwnerFileOpenEntrypointCount:           remoteInt(result, "real_run_receipt_summary_kde_page_owner_file_open_entrypoint_count"),
		HostRootModified:                              false,
		PrivilegedContainerRequired:                   false,
		HostNetworkingRequired:                        false,
		DockerSocketMounted:                           false,
		BroadHostMountRequired:                        false,
		BackendDetailsExposed:                         false,
		RawLauncherOutputExposed:                      false,
		RawWindowEvidenceExposed:                      false,
		ExecutableNameExposed:                         false,
		AcceptanceReady:                               true,
		DesktopSafeSummary:                            displayName + " completed a Runtime-owner file-open GUI run and produced desktop compatibility cards.",
	}
	if acceptance.DisplayName == "" {
		acceptance.DisplayName = acceptance.AppID
	}
	if acceptance.AppVersion == "" {
		acceptance.AppVersion = acceptance.Version
	}
	if err := validateNoBackendTerms(acceptance, "real Windows app run acceptance"); err != nil {
		return RealWinAppRunAcceptance{}, err
	}
	return acceptance, nil
}
