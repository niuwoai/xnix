package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Q4WinAppAcceptanceSchemaVersion = "xnix.runtime.q4_winapp_acceptance.v1"
	Q4WinAppAcceptanceRequestType   = "q4-winapp-acceptance-preview"
)

type Q4WinAppAcceptanceRequest struct {
	SmokeReportPath string
}

type Q4WinAppAcceptance struct {
	Version                                       string `json:"version"`
	SchemaVersion                                 string `json:"schema_version"`
	RequestType                                   string `json:"request_type"`
	Source                                        string `json:"source"`
	RuntimeMethod                                 string `json:"runtime_method"`
	ReadMethod                                    string `json:"read_method"`
	AcceptanceType                                string `json:"acceptance_type"`
	SmokeReportConsumed                           bool   `json:"smoke_report_consumed"`
	SmokeReportPathExposed                        bool   `json:"smoke_report_path_exposed"`
	OutputPathExposed                             bool   `json:"output_path_exposed"`
	DelegatedCommandExposed                       bool   `json:"delegated_command_exposed"`
	RemoteHostExposed                             bool   `json:"remote_host_exposed"`
	AppID                                         string `json:"app_id"`
	DisplayName                                   string `json:"display_name"`
	KnownAppID                                    string `json:"known_app_id"`
	KnownAppSelected                              bool   `json:"known_app_selected"`
	RemoteExecutableConfigured                    bool   `json:"remote_executable_configured"`
	RemoteExecutablePathExposed                   bool   `json:"remote_executable_path_exposed"`
	RemoteFileArgumentConfigured                  bool   `json:"remote_file_argument_configured"`
	RemoteFileArgumentPathExposed                 bool   `json:"remote_file_argument_path_exposed"`
	SampleFileArgumentConfigured                  bool   `json:"sample_file_argument_configured"`
	LaunchMode                                    string `json:"launch_mode"`
	FileOpenEntrypointRequested                   bool   `json:"file_open_entrypoint_requested"`
	RealRunAcceptanceRequired                     bool   `json:"real_run_acceptance_required"`
	Q4CompileRequired                             bool   `json:"q4_compile_required"`
	HostCompilationAvoided                        bool   `json:"host_compilation_avoided"`
	DelegatedExecuteResultConsumed                bool   `json:"delegated_execute_result_consumed"`
	RemoteBuildCompleted                          bool   `json:"remote_build_completed"`
	EvidenceOutputWritten                         bool   `json:"evidence_output_written"`
	KDEPageOutputWritten                          bool   `json:"kde_page_output_written"`
	KDEActionOutputWritten                        bool   `json:"kde_action_output_written"`
	WindowObserved                                bool   `json:"window_observed"`
	WindowMatchObserved                           bool   `json:"window_match_observed"`
	RealRunReceiptSummaryReady                    bool   `json:"real_run_receipt_summary_ready"`
	RealRunReceiptSummaryFileOpenVerified         bool   `json:"real_run_receipt_summary_file_open_verified"`
	RealRunAcceptanceOutputWritten                bool   `json:"real_run_acceptance_output_written"`
	RealRunAcceptanceReady                        bool   `json:"real_run_acceptance_ready"`
	RealRunAcceptanceCenterProjectionConsumed     bool   `json:"real_run_acceptance_center_projection_consumed"`
	RealRunAcceptanceKDEPageProjectionConsumed    bool   `json:"real_run_acceptance_kde_page_projection_consumed"`
	OwnerFileOpenEntrypointInvoked                bool   `json:"owner_file_open_entrypoint_invoked"`
	RuntimeEvidenceOwnerFileOpenEntrypointInvoked bool   `json:"runtime_evidence_owner_file_open_entrypoint_invoked"`
	KDEPageOwnerFileOpenEntrypointCount           int    `json:"kde_page_owner_file_open_entrypoint_count"`
	HostRootModified                              bool   `json:"host_root_modified"`
	PrivilegedContainerRequired                   bool   `json:"privileged_container_required"`
	HostNetworkingRequired                        bool   `json:"host_networking_required"`
	DockerSocketMounted                           bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                        bool   `json:"broad_host_mount_required"`
	BackendDetailsExposed                         bool   `json:"backend_details_exposed"`
	RawWindowEvidenceExposed                      bool   `json:"raw_window_evidence_exposed"`
	ExecutablePathExposed                         bool   `json:"executable_path_exposed"`
	AcceptanceReady                               bool   `json:"acceptance_ready"`
	DesktopSafeSummary                            string `json:"desktop_safe_summary"`
}

func PreviewQ4WinAppAcceptance(request Q4WinAppAcceptanceRequest) (Q4WinAppAcceptance, error) {
	path := strings.TrimSpace(request.SmokeReportPath)
	if path == "" {
		return Q4WinAppAcceptance{}, errors.New("q4 Windows app acceptance requires --q4-winapp-smoke")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Q4WinAppAcceptance{}, fmt.Errorf("read q4 Windows app smoke report: %w", err)
	}
	return PreviewQ4WinAppAcceptanceJSON(content)
}

func PreviewQ4WinAppAcceptanceJSON(content []byte) (Q4WinAppAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return Q4WinAppAcceptance{}, fmt.Errorf("parse q4 Windows app smoke report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.scripts.q4_winapp_smoke.v1" ||
		remoteString(report, "request_type") != "q4-winapp-smoke" {
		return Q4WinAppAcceptance{}, errors.New("q4 Windows app acceptance requires q4_winapp_smoke output")
	}
	if remoteString(report, "status") != "passed" || !remoteBool(report, "execute") {
		return Q4WinAppAcceptance{}, errors.New("q4 Windows app acceptance requires an executed passed smoke")
	}

	appID := remoteString(report, "app_id")
	displayName := remoteString(report, "display_name")
	knownAppID := remoteString(report, "known_app_id")
	knownAppSelected := knownAppID != ""
	remoteExecutableConfigured := remoteBool(report, "remote_executable_configured")
	remoteFileArgumentConfigured := remoteBool(report, "remote_file_argument_configured")
	sampleFileArgumentConfigured := remoteString(report, "sample_file_argument") != ""
	realRunAcceptanceRequired := remoteBool(report, "real_run_acceptance_required")
	delegatedConsumed := remoteString(report, "delegated_execute_result_schema") == "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1" &&
		remoteBool(report, "delegated_execute_result_output_written")
	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required")
	realRunAcceptanceReady := !realRunAcceptanceRequired ||
		(remoteString(report, "launch_mode") == "owner-controlled-launch" &&
			remoteBool(report, "file_open_entrypoint_requested") &&
			remoteBool(report, "real_run_receipt_summary_ready") &&
			remoteBool(report, "real_run_receipt_summary_file_open_verified") &&
			remoteBool(report, "real_run_acceptance_output_written") &&
			remoteBool(report, "real_run_acceptance_ready") &&
			remoteBool(report, "real_run_acceptance_center_projection_consumed") &&
			remoteBool(report, "real_run_acceptance_kde_page_projection_consumed") &&
			remoteBool(report, "owner_file_open_entrypoint_invoked") &&
			remoteBool(report, "runtime_evidence_owner_file_open_entrypoint_invoked") &&
			remoteInt(report, "kde_page_known_app_owner_file_open_entrypoint_count") > 0)
	acceptanceReady := appID != "" &&
		displayName != "" &&
		knownAppSelected != remoteExecutableConfigured &&
		!remoteBool(report, "remote_executable_path_exposed") &&
		!remoteBool(report, "remote_file_argument_path_exposed") &&
		remoteBool(report, "q4_compile_required") &&
		remoteBool(report, "host_compilation_avoided") &&
		delegatedConsumed &&
		remoteBool(report, "remote_build_completed") &&
		remoteBool(report, "evidence_output_written") &&
		remoteBool(report, "kde_page_output_written") &&
		remoteBool(report, "window_observed") &&
		remoteBool(report, "window_match_observed") &&
		realRunAcceptanceReady &&
		!unsafeGateOpen
	if !acceptanceReady {
		return Q4WinAppAcceptance{}, errors.New("q4 Windows app acceptance requires app identity, delegated q4 GUI evidence, matched window observation, optional real-run acceptance, q4 host compilation, and closed safety gates")
	}

	acceptance := Q4WinAppAcceptance{
		Version:                                       remoteString(report, "version"),
		SchemaVersion:                                 Q4WinAppAcceptanceSchemaVersion,
		RequestType:                                   Q4WinAppAcceptanceRequestType,
		Source:                                        "q4-winapp-smoke+go-runtime-acceptance",
		RuntimeMethod:                                 "PreviewQ4WinAppAcceptance",
		ReadMethod:                                    "GetQ4WinAppAcceptance",
		AcceptanceType:                                "generic-q4-windows-app-real-run-acceptance",
		SmokeReportConsumed:                           true,
		SmokeReportPathExposed:                        false,
		OutputPathExposed:                             false,
		DelegatedCommandExposed:                       false,
		RemoteHostExposed:                             false,
		AppID:                                         appID,
		DisplayName:                                   displayName,
		KnownAppID:                                    knownAppID,
		KnownAppSelected:                              knownAppSelected,
		RemoteExecutableConfigured:                    remoteExecutableConfigured,
		RemoteExecutablePathExposed:                   false,
		RemoteFileArgumentConfigured:                  remoteFileArgumentConfigured,
		RemoteFileArgumentPathExposed:                 false,
		SampleFileArgumentConfigured:                  sampleFileArgumentConfigured,
		LaunchMode:                                    remoteString(report, "launch_mode"),
		FileOpenEntrypointRequested:                   remoteBool(report, "file_open_entrypoint_requested"),
		RealRunAcceptanceRequired:                     realRunAcceptanceRequired,
		Q4CompileRequired:                             true,
		HostCompilationAvoided:                        true,
		DelegatedExecuteResultConsumed:                true,
		RemoteBuildCompleted:                          true,
		EvidenceOutputWritten:                         true,
		KDEPageOutputWritten:                          true,
		KDEActionOutputWritten:                        remoteBool(report, "kde_action_output_written"),
		WindowObserved:                                true,
		WindowMatchObserved:                           true,
		RealRunReceiptSummaryReady:                    remoteBool(report, "real_run_receipt_summary_ready"),
		RealRunReceiptSummaryFileOpenVerified:         remoteBool(report, "real_run_receipt_summary_file_open_verified"),
		RealRunAcceptanceOutputWritten:                remoteBool(report, "real_run_acceptance_output_written"),
		RealRunAcceptanceReady:                        remoteBool(report, "real_run_acceptance_ready"),
		RealRunAcceptanceCenterProjectionConsumed:     remoteBool(report, "real_run_acceptance_center_projection_consumed"),
		RealRunAcceptanceKDEPageProjectionConsumed:    remoteBool(report, "real_run_acceptance_kde_page_projection_consumed"),
		OwnerFileOpenEntrypointInvoked:                remoteBool(report, "owner_file_open_entrypoint_invoked"),
		RuntimeEvidenceOwnerFileOpenEntrypointInvoked: remoteBool(report, "runtime_evidence_owner_file_open_entrypoint_invoked"),
		KDEPageOwnerFileOpenEntrypointCount:           remoteInt(report, "kde_page_known_app_owner_file_open_entrypoint_count"),
		HostRootModified:                              false,
		PrivilegedContainerRequired:                   false,
		HostNetworkingRequired:                        false,
		DockerSocketMounted:                           false,
		BroadHostMountRequired:                        false,
		BackendDetailsExposed:                         false,
		RawWindowEvidenceExposed:                      false,
		ExecutablePathExposed:                         false,
		AcceptanceReady:                               true,
		DesktopSafeSummary:                            "A q4-hosted Windows GUI app completed the generic Runtime-owned acceptance lane.",
	}
	if err := validateNoBackendTerms(acceptance, "q4 Windows app acceptance"); err != nil {
		return Q4WinAppAcceptance{}, err
	}
	return acceptance, nil
}
