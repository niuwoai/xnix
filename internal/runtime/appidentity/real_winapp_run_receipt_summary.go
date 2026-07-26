package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	RealWinAppRunReceiptSummarySchemaVersion = "xnix.runtime.real_winapp_run_receipt_summary.v1"
	RealWinAppRunReceiptSummaryRequestType   = "real-winapp-run-receipt-summary-preview"
)

type RealWinAppRunReceiptSummaryRequest struct {
	RemoteSmokeReportPath string
}

type RealWinAppRunReceiptSummary struct {
	Version                                       string `json:"version"`
	SchemaVersion                                 string `json:"schema_version"`
	RequestType                                   string `json:"request_type"`
	Source                                        string `json:"source"`
	RuntimeMethod                                 string `json:"runtime_method"`
	ReadMethod                                    string `json:"read_method"`
	ReceiptType                                   string `json:"receipt_type"`
	ReportConsumed                                bool   `json:"report_consumed"`
	ReportPathExposed                             bool   `json:"report_path_exposed"`
	RemoteHostExposed                             bool   `json:"remote_host_exposed"`
	AppID                                         string `json:"app_id"`
	DisplayName                                   string `json:"display_name"`
	AppVersion                                    string `json:"app_version"`
	GUIAppName                                    string `json:"gui_app_name"`
	ExecutionHostClass                            string `json:"execution_host_class"`
	BackendClass                                  string `json:"backend_class"`
	LaunchMode                                    string `json:"launch_mode"`
	RunPassed                                     bool   `json:"run_passed"`
	RealExecutionObserved                         bool   `json:"real_execution_observed"`
	WindowObserved                                bool   `json:"window_observed"`
	WindowMatch                                   string `json:"window_match,omitempty"`
	WindowMatchObserved                           bool   `json:"window_match_observed"`
	FileOpenVerified                              bool   `json:"file_open_verified"`
	FileArgumentCount                             int    `json:"file_argument_count"`
	FileArgumentCopiedCount                       int    `json:"file_argument_copied_count"`
	FileArgumentsPassed                           bool   `json:"file_arguments_passed"`
	FileArgumentWinePathTranslated                bool   `json:"file_argument_winepath_translated"`
	FileArgumentWinePathTranslatedCount           int    `json:"file_argument_winepath_translated_count"`
	RawFileArgumentPathExposed                    bool   `json:"raw_file_argument_path_exposed"`
	OwnerControlledLaunchVerified                 bool   `json:"owner_controlled_launch_verified"`
	OwnerManagedLauncherInvoked                   bool   `json:"owner_managed_launcher_invoked"`
	OwnerFileOpenEntrypointInvoked                bool   `json:"owner_file_open_entrypoint_invoked"`
	OwnerDelegatedSmokePassed                     bool   `json:"owner_delegated_smoke_passed"`
	RuntimeEvidenceConsumed                       bool   `json:"runtime_evidence_consumed"`
	RuntimeEvidenceOwnerFileOpenVerified          bool   `json:"runtime_evidence_owner_file_open_verified"`
	RuntimeEvidenceOwnerFileOpenEntrypointInvoked bool   `json:"runtime_evidence_owner_file_open_entrypoint_invoked"`
	KDEPageEvidenceConsumed                       bool   `json:"kde_page_evidence_consumed"`
	KDEActionEvidenceConsumed                     bool   `json:"kde_action_evidence_consumed"`
	KDEActionOwnerFileOpenVerified                bool   `json:"kde_action_owner_file_open_verified"`
	KDEActionOwnerFileOpenEntrypointInvoked       bool   `json:"kde_action_owner_file_open_entrypoint_invoked"`
	HostRootModified                              bool   `json:"host_root_modified"`
	PrivilegedContainerRequired                   bool   `json:"privileged_container_required"`
	HostNetworkingRequired                        bool   `json:"host_networking_required"`
	DockerSocketMounted                           bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                        bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                         bool   `json:"backend_details_exposed"`
	RawLauncherOutputExposed                      bool   `json:"raw_launcher_output_exposed"`
	ReceiptReady                                  bool   `json:"receipt_ready"`
	DesktopSafeSummary                            string `json:"desktop_safe_summary"`
}

func PreviewRealWinAppRunReceiptSummary(request RealWinAppRunReceiptSummaryRequest) (RealWinAppRunReceiptSummary, error) {
	path := strings.TrimSpace(request.RemoteSmokeReportPath)
	if path == "" {
		return RealWinAppRunReceiptSummary{}, errors.New("real Windows app run receipt summary requires --remote-smoke-report")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return RealWinAppRunReceiptSummary{}, fmt.Errorf("read real Windows app remote smoke report: %w", err)
	}
	return PreviewRealWinAppRunReceiptSummaryJSON(content)
}

func PreviewRealWinAppRunReceiptSummaryJSON(content []byte) (RealWinAppRunReceiptSummary, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return RealWinAppRunReceiptSummary{}, fmt.Errorf("parse real Windows app remote smoke report: %w", err)
	}
	if remoteString(report, "request_type") != "remote-wine-guest-gui-smoke" {
		return RealWinAppRunReceiptSummary{}, errors.New("real Windows app run receipt summary requires a remote Wine guest GUI smoke report")
	}
	if remoteString(report, "smoke_status") != "passed" && remoteString(report, "status") != "passed" {
		return RealWinAppRunReceiptSummary{}, errors.New("real Windows app run receipt summary requires a passed remote smoke report")
	}
	appID := firstRemoteString(report, "known_app_id", "runtime_evidence_app_id", "kde_page_app_id")
	if appID == "" {
		return RealWinAppRunReceiptSummary{}, errors.New("real Windows app run receipt summary requires safe app identity")
	}
	displayName := firstRemoteString(report, "known_app_name", "runtime_evidence_display_name", "kde_page_display_name")
	appVersion := firstRemoteString(report, "known_app_version", "runtime_evidence_app_version")
	if displayName == "" {
		displayName = appID
	}

	windowObserved := remoteBool(report, "x_window_observed") || remoteBool(report, "runtime_evidence_window_observed")
	fileOpenVerified := remoteBool(report, "runtime_evidence_owner_file_open_verified") &&
		remoteBool(report, "kde_action_owner_file_open_verified") &&
		remoteBool(report, "owner_delegated_file_arguments_passed") &&
		remoteBool(report, "owner_delegated_file_argument_winepath_translated") &&
		!remoteBool(report, "owner_delegated_raw_file_argument_path_exposed")
	ownerControlledLaunchVerified := remoteBool(report, "owner_controlled_launch_requested") &&
		remoteBool(report, "owner_evidence_handoff_ready") &&
		remoteBool(report, "owner_managed_launcher_invoked") &&
		remoteBool(report, "owner_delegated_smoke_passed")
	kdePageEvidenceConsumed := remoteBool(report, "kde_page_output_written") &&
		remoteInt(report, "kde_page_known_app_gui_evidence_count") > 0
	kdeActionEvidenceConsumed := remoteBool(report, "kde_action_output_written") &&
		remoteBool(report, "kde_action_owner_file_open_verified")
	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required")
	receiptReady := windowObserved &&
		fileOpenVerified &&
		ownerControlledLaunchVerified &&
		remoteBool(report, "runtime_evidence_report_consumed") &&
		kdePageEvidenceConsumed &&
		kdeActionEvidenceConsumed &&
		!unsafeGateOpen
	if !receiptReady {
		return RealWinAppRunReceiptSummary{}, errors.New("real Windows app run receipt summary requires observed execution, owner file-open evidence, KDE evidence, and closed safety gates")
	}

	summary := RealWinAppRunReceiptSummary{
		Version:                              remoteString(report, "version"),
		SchemaVersion:                        RealWinAppRunReceiptSummarySchemaVersion,
		RequestType:                          RealWinAppRunReceiptSummaryRequestType,
		Source:                               "remote-wine-guest-gui-smoke+runtime-real-run-receipt-summary",
		RuntimeMethod:                        "PreviewRealWinAppRunReceiptSummary",
		ReadMethod:                           "GetRealWinAppRunReceiptSummary",
		ReceiptType:                          "real-windows-app-run-receipt-summary",
		ReportConsumed:                       true,
		ReportPathExposed:                    false,
		RemoteHostExposed:                    false,
		AppID:                                appID,
		DisplayName:                          displayName,
		AppVersion:                           appVersion,
		GUIAppName:                           "known-gui-app",
		ExecutionHostClass:                   "q4-remote-validation-host",
		BackendClass:                         "managed-guest-gui",
		LaunchMode:                           remoteString(report, "launch_mode"),
		RunPassed:                            true,
		RealExecutionObserved:                true,
		WindowObserved:                       windowObserved,
		WindowMatch:                          remoteString(report, "window_match"),
		WindowMatchObserved:                  remoteBool(report, "window_match_observed"),
		FileOpenVerified:                     fileOpenVerified,
		FileArgumentCount:                    remoteInt(report, "file_argument_count"),
		FileArgumentCopiedCount:              remoteInt(report, "file_argument_copied_count"),
		FileArgumentsPassed:                  remoteBool(report, "file_arguments_passed"),
		FileArgumentWinePathTranslated:       remoteBool(report, "file_argument_winepath_translated"),
		FileArgumentWinePathTranslatedCount:  remoteInt(report, "file_argument_winepath_translated_count"),
		RawFileArgumentPathExposed:           remoteBool(report, "raw_file_argument_path_exposed") || remoteBool(report, "owner_delegated_raw_file_argument_path_exposed"),
		OwnerControlledLaunchVerified:        ownerControlledLaunchVerified,
		OwnerManagedLauncherInvoked:          remoteBool(report, "owner_managed_launcher_invoked"),
		OwnerFileOpenEntrypointInvoked:       remoteBool(report, "owner_file_open_entrypoint_invoked"),
		OwnerDelegatedSmokePassed:            remoteBool(report, "owner_delegated_smoke_passed"),
		RuntimeEvidenceConsumed:              remoteBool(report, "runtime_evidence_report_consumed"),
		RuntimeEvidenceOwnerFileOpenVerified: remoteBool(report, "runtime_evidence_owner_file_open_verified"),
		RuntimeEvidenceOwnerFileOpenEntrypointInvoked: remoteBool(report, "runtime_evidence_owner_file_open_entrypoint_invoked"),
		KDEPageEvidenceConsumed:                       kdePageEvidenceConsumed,
		KDEActionEvidenceConsumed:                     kdeActionEvidenceConsumed,
		KDEActionOwnerFileOpenVerified:                remoteBool(report, "kde_action_owner_file_open_verified"),
		KDEActionOwnerFileOpenEntrypointInvoked:       remoteBool(report, "kde_action_owner_file_open_entrypoint_invoked"),
		HostRootModified:                              false,
		PrivilegedContainerRequired:                   false,
		HostNetworkingRequired:                        false,
		DockerSocketMounted:                           false,
		BroadHostMountRequired:                        false,
		BackendDetailsExposed:                         false,
		RawLauncherOutputExposed:                      false,
		ReceiptReady:                                  true,
		DesktopSafeSummary:                            displayName + " has a q4-verified Runtime-owner file-open GUI run with observed window and closed unsafe host gates.",
	}
	if summary.DisplayName == "" {
		summary.DisplayName = summary.AppID
	}
	if summary.AppVersion == "" {
		summary.AppVersion = summary.Version
	}
	if err := validateNoBackendTerms(summary, "real Windows app run receipt summary"); err != nil {
		return RealWinAppRunReceiptSummary{}, err
	}
	return summary, nil
}

func KnownAppSmokeEvidenceFromRealWinAppRunReceiptSummary(payload []byte) (KnownAppSmokeEvidenceSummary, error) {
	var summary RealWinAppRunReceiptSummary
	if err := json.Unmarshal(payload, &summary); err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("parse real Windows app run receipt summary: %w", err)
	}
	switch {
	case summary.SchemaVersion != RealWinAppRunReceiptSummarySchemaVersion:
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("real Windows app run receipt summary has unsupported schema %q", summary.SchemaVersion)
	case summary.RequestType != RealWinAppRunReceiptSummaryRequestType:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary has invalid request type")
	case summary.ReceiptType != "real-windows-app-run-receipt-summary":
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary has invalid receipt type")
	case !summary.ReportConsumed || summary.ReportPathExposed || summary.RemoteHostExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary must consume the report without exposing report location details")
	case !summary.ReceiptReady || !summary.RunPassed || !summary.RealExecutionObserved || !summary.WindowObserved:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires a ready passed observed GUI run")
	case !summary.FileOpenVerified || !summary.FileArgumentsPassed || !summary.FileArgumentWinePathTranslated:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires verified file-open handoff evidence")
	case summary.FileArgumentCount <= 0 || summary.FileArgumentCopiedCount < summary.FileArgumentCount || summary.FileArgumentWinePathTranslatedCount < summary.FileArgumentCount:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires copied and translated file arguments")
	case summary.RawFileArgumentPathExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary must not expose raw file argument paths")
	case !summary.OwnerControlledLaunchVerified || !summary.OwnerManagedLauncherInvoked || !summary.OwnerFileOpenEntrypointInvoked || !summary.OwnerDelegatedSmokePassed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires Runtime-owner file-open launch evidence")
	case !summary.RuntimeEvidenceConsumed || !summary.RuntimeEvidenceOwnerFileOpenVerified || !summary.RuntimeEvidenceOwnerFileOpenEntrypointInvoked:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires consumed Runtime evidence")
	case !summary.KDEPageEvidenceConsumed || !summary.KDEActionEvidenceConsumed || !summary.KDEActionOwnerFileOpenVerified || !summary.KDEActionOwnerFileOpenEntrypointInvoked:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires consumed KDE page and action evidence")
	case summary.HostRootModified || summary.PrivilegedContainerRequired || summary.HostNetworkingRequired || summary.DockerSocketMounted || summary.BroadHostMountRequired:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires closed host and container safety gates")
	case summary.BackendDetailsExposed || summary.RawLauncherOutputExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary must not expose backend or launcher details")
	case strings.TrimSpace(summary.WindowMatch) == "" || !singleLine(summary.WindowMatch) || !summary.WindowMatchObserved:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app run receipt summary requires safe observed window-match evidence")
	}

	evidence, err := normalizeKnownAppSmokeEvidenceItem(KnownAppSmokeEvidenceSummary{
		AppID:                                        summary.AppID,
		DisplayName:                                  summary.DisplayName,
		AppVersion:                                   summary.AppVersion,
		EvidenceSource:                               GUISmokeEvidenceSourceWineGuest,
		SmokeStatus:                                  "passed",
		XWindowObserved:                              true,
		WindowObserved:                               true,
		ExecutionEvidenceRecorded:                    true,
		StagedLauncherVerified:                       true,
		OwnerControlledRuntimeLaunchVerified:         true,
		OwnerFileOpenVerified:                        true,
		OwnerFileOpenEntrypointInvoked:               true,
		OwnerDelegatedFileArgumentCount:              summary.FileArgumentCount,
		OwnerDelegatedFileArgumentCopiedCount:        summary.FileArgumentCopiedCount,
		OwnerDelegatedFileArgumentsPassed:            true,
		OwnerDelegatedFileArgumentWinepathTranslated: true,
		OwnerDelegatedFileArgumentWinepathTranslatedCount: summary.FileArgumentWinePathTranslatedCount,
		OwnerDelegatedRawFileArgumentPathExposed:          false,
		OwnerDelegatedWindowMatch:                         summary.WindowMatch,
		OwnerDelegatedWindowMatchObserved:                 true,
		RuntimeDispatchVerified:                           true,
	})
	if err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("consume real Windows app run receipt summary: %w", err)
	}
	return evidence, nil
}

func remoteString(payload map[string]any, key string) string {
	value, _ := payload[key].(string)
	return strings.TrimSpace(value)
}

func firstRemoteString(payload map[string]any, keys ...string) string {
	for _, key := range keys {
		if value := remoteString(payload, key); value != "" {
			return value
		}
	}
	return ""
}

func remoteBool(payload map[string]any, key string) bool {
	value, _ := payload[key].(bool)
	return value
}

func remoteInt(payload map[string]any, key string) int {
	switch value := payload[key].(type) {
	case float64:
		return int(value)
	case int:
		return value
	default:
		return 0
	}
}
