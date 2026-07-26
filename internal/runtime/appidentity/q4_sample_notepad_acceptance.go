package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Q4SampleNotepadAcceptanceSchemaVersion = "xnix.runtime.q4_sample_notepad_acceptance.v1"
	Q4SampleNotepadAcceptanceRequestType   = "q4-sample-notepad-acceptance-preview"
)

type Q4SampleNotepadAcceptanceRequest struct {
	SmokeReportPath string
}

type Q4SampleNotepadAcceptance struct {
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
	LaunchMode                                    string `json:"launch_mode"`
	FileOpenEntrypointRequested                   bool   `json:"file_open_entrypoint_requested"`
	HostCompilationAvoided                        bool   `json:"host_compilation_avoided"`
	DelegatedExecuteResultConsumed                bool   `json:"delegated_execute_result_consumed"`
	RealRunReceiptSummaryReady                    bool   `json:"real_run_receipt_summary_ready"`
	RealRunReceiptSummaryFileOpenVerified         bool   `json:"real_run_receipt_summary_file_open_verified"`
	RealRunAcceptanceOutputWritten                bool   `json:"real_run_acceptance_output_written"`
	RealRunAcceptanceReady                        bool   `json:"real_run_acceptance_ready"`
	RealRunAcceptanceCenterProjectionConsumed     bool   `json:"real_run_acceptance_center_projection_consumed"`
	RealRunAcceptanceKDEPageProjectionConsumed    bool   `json:"real_run_acceptance_kde_page_projection_consumed"`
	WindowMatchObserved                           bool   `json:"window_match_observed"`
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
	ExecutableNameExposed                         bool   `json:"executable_name_exposed"`
	AcceptanceReady                               bool   `json:"acceptance_ready"`
	DesktopSafeSummary                            string `json:"desktop_safe_summary"`
}

func PreviewQ4SampleNotepadAcceptance(request Q4SampleNotepadAcceptanceRequest) (Q4SampleNotepadAcceptance, error) {
	path := strings.TrimSpace(request.SmokeReportPath)
	if path == "" {
		return Q4SampleNotepadAcceptance{}, errors.New("q4 Sample Notepad acceptance requires --q4-sample-notepad-smoke")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Q4SampleNotepadAcceptance{}, fmt.Errorf("read q4 Sample Notepad smoke report: %w", err)
	}
	return PreviewQ4SampleNotepadAcceptanceJSON(content)
}

func PreviewQ4SampleNotepadAcceptanceJSON(content []byte) (Q4SampleNotepadAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return Q4SampleNotepadAcceptance{}, fmt.Errorf("parse q4 Sample Notepad smoke report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.scripts.q4_sample_notepad_smoke.v1" ||
		remoteString(report, "request_type") != "q4-sample-notepad-smoke" {
		return Q4SampleNotepadAcceptance{}, errors.New("q4 Sample Notepad acceptance requires q4_sample_notepad_smoke output")
	}
	if remoteString(report, "status") != "passed" || !remoteBool(report, "execute") {
		return Q4SampleNotepadAcceptance{}, errors.New("q4 Sample Notepad acceptance requires an executed passed smoke")
	}

	unsafeGateOpen := remoteBool(report, "host_root_modified") ||
		remoteBool(report, "privileged_container_required") ||
		remoteBool(report, "host_networking_required") ||
		remoteBool(report, "docker_socket_mounted") ||
		remoteBool(report, "broad_host_mount_required")
	delegatedConsumed := remoteString(report, "delegated_execute_result_schema") == "xnix.scripts.remote_wine_guest_gui_smoke.execute_result.v1" &&
		remoteBool(report, "delegated_execute_result_output_written")
	acceptanceReady := remoteString(report, "app_id") == "org.xnix.sample.notepad" &&
		remoteString(report, "launch_mode") == "owner-controlled-launch" &&
		remoteBool(report, "file_open_entrypoint_requested") &&
		remoteBool(report, "host_compilation_avoided") &&
		delegatedConsumed &&
		remoteBool(report, "real_run_receipt_summary_ready") &&
		remoteBool(report, "real_run_receipt_summary_file_open_verified") &&
		remoteBool(report, "real_run_acceptance_output_written") &&
		remoteBool(report, "real_run_acceptance_ready") &&
		remoteBool(report, "real_run_acceptance_center_projection_consumed") &&
		remoteBool(report, "real_run_acceptance_kde_page_projection_consumed") &&
		remoteBool(report, "window_match_observed") &&
		remoteBool(report, "owner_file_open_entrypoint_invoked") &&
		remoteBool(report, "runtime_evidence_owner_file_open_entrypoint_invoked") &&
		remoteInt(report, "kde_page_known_app_owner_file_open_entrypoint_count") > 0 &&
		!unsafeGateOpen
	if !acceptanceReady {
		return Q4SampleNotepadAcceptance{}, errors.New("q4 Sample Notepad acceptance requires passed app identity, owner file-open entrypoint, Go real-run acceptance, window evidence, and closed safety gates")
	}

	acceptance := Q4SampleNotepadAcceptance{
		Version:                               remoteString(report, "version"),
		SchemaVersion:                         Q4SampleNotepadAcceptanceSchemaVersion,
		RequestType:                           Q4SampleNotepadAcceptanceRequestType,
		Source:                                "q4-sample-notepad-smoke+go-real-run-acceptance",
		RuntimeMethod:                         "PreviewQ4SampleNotepadAcceptance",
		ReadMethod:                            "GetQ4SampleNotepadAcceptance",
		AcceptanceType:                        "q4-sample-notepad-real-windows-app-acceptance",
		SmokeReportConsumed:                   true,
		SmokeReportPathExposed:                false,
		OutputPathExposed:                     false,
		DelegatedCommandExposed:               false,
		RemoteHostExposed:                     false,
		AppID:                                 "org.xnix.sample.notepad",
		DisplayName:                           "Sample Notepad",
		LaunchMode:                            "owner-controlled-launch",
		FileOpenEntrypointRequested:           true,
		HostCompilationAvoided:                true,
		DelegatedExecuteResultConsumed:        true,
		RealRunReceiptSummaryReady:            true,
		RealRunReceiptSummaryFileOpenVerified: true,
		RealRunAcceptanceOutputWritten:        true,
		RealRunAcceptanceReady:                true,
		RealRunAcceptanceCenterProjectionConsumed:     true,
		RealRunAcceptanceKDEPageProjectionConsumed:    true,
		WindowMatchObserved:                           true,
		OwnerFileOpenEntrypointInvoked:                true,
		RuntimeEvidenceOwnerFileOpenEntrypointInvoked: true,
		KDEPageOwnerFileOpenEntrypointCount:           remoteInt(report, "kde_page_known_app_owner_file_open_entrypoint_count"),
		HostRootModified:                              false,
		PrivilegedContainerRequired:                   false,
		HostNetworkingRequired:                        false,
		DockerSocketMounted:                           false,
		BroadHostMountRequired:                        false,
		BackendDetailsExposed:                         false,
		RawWindowEvidenceExposed:                      false,
		ExecutableNameExposed:                         false,
		AcceptanceReady:                               true,
		DesktopSafeSummary:                            "Sample Notepad completed the q4 Runtime-owner file-open GUI acceptance lane.",
	}
	if err := validateNoBackendTerms(acceptance, "q4 Sample Notepad acceptance"); err != nil {
		return Q4SampleNotepadAcceptance{}, err
	}
	return acceptance, nil
}
