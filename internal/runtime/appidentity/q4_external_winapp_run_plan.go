package appidentity

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

const (
	Q4ExternalWinAppRunPlanSchemaVersion = "xnix.runtime.q4_external_winapp_run_plan.v1"
	Q4ExternalWinAppRunPlanRequestType   = "q4-external-winapp-run-plan-preview"
)

type Q4ExternalWinAppRunPlanRequest struct {
	Version             string
	ProjectRoot         string
	LocalExecutable     string
	AppID               string
	DisplayName         string
	WindowMatch         string
	RemoteHost          string
	RemoteMaterialsRoot string
	Output              string
	MarkdownOutput      string
	Execute             bool
}

type Q4ExternalWinAppRunPlanPreview struct {
	Version                                 string   `json:"version"`
	SchemaVersion                           string   `json:"schema_version"`
	RequestType                             string   `json:"request_type"`
	Source                                  string   `json:"source"`
	RuntimeMethod                           string   `json:"runtime_method"`
	ReadMethod                              string   `json:"read_method"`
	AppID                                   string   `json:"app_id"`
	DisplayName                             string   `json:"display_name"`
	WindowMatchConfigured                   bool     `json:"window_match_configured"`
	WindowMatchRequired                     bool     `json:"window_match_required"`
	WindowMatch                             string   `json:"window_match"`
	LocalExecutableConfigured               bool     `json:"local_executable_configured"`
	LocalExecutablePathClass                string   `json:"local_executable_path_class"`
	LocalExecutablePathExposedToDesktop     bool     `json:"local_executable_path_exposed_to_desktop"`
	LocalExecutableMZHeaderRequired         bool     `json:"local_executable_mz_header_required"`
	LocalExecutableSHA256Required           bool     `json:"local_executable_sha256_required"`
	RemoteHostConfigured                    bool     `json:"remote_host_configured"`
	RemoteHostExposedToDesktop              bool     `json:"remote_host_exposed_to_desktop"`
	RemoteMaterialsRootClass                string   `json:"remote_materials_root_class"`
	RemoteExecutableGeneratedFromDigest     bool     `json:"remote_executable_generated_from_digest"`
	RemoteExecutablePathExposedToDesktop    bool     `json:"remote_executable_path_exposed_to_desktop"`
	RemoteUploadRequired                    bool     `json:"remote_upload_required"`
	RemoteExecutionRequired                 bool     `json:"remote_execution_required"`
	Q4MaterializationRequired               bool     `json:"q4_materialization_required"`
	Q4BuildPreferred                        bool     `json:"q4_build_preferred"`
	Q4ExecutionRequired                     bool     `json:"q4_execution_required"`
	HostCompilationRequired                 bool     `json:"host_compilation_required"`
	HostCompilationAvoided                  bool     `json:"host_compilation_avoided"`
	LocalHostRole                           string   `json:"local_host_role"`
	DelegatedScript                         string   `json:"delegated_script"`
	DelegatedRequestType                    string   `json:"delegated_request_type"`
	DelegatedAcceptanceRequestType          string   `json:"delegated_acceptance_request_type"`
	DelegatedCommand                        []string `json:"delegated_command"`
	DelegatedCommandExposedToOperator       bool     `json:"delegated_command_exposed_to_operator"`
	DelegatedCommandExposedToDesktop        bool     `json:"delegated_command_exposed_to_desktop"`
	StagedExternalFixtureRequired           bool     `json:"staged_external_fixture_required"`
	RealFileOpenEvidenceRequired            bool     `json:"real_file_open_evidence_required"`
	WindowsProcessWindowObservationRequired bool     `json:"windows_process_window_observation_required"`
	AcceptedApplicationDetailStateRequired  string   `json:"accepted_application_detail_state_required"`
	RuntimeOwned                            bool     `json:"runtime_owned"`
	GoRuntimeBacked                         bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool     `json:"kde_policy_owner"`
	FullSmokeRequired                       bool     `json:"full_smoke_required"`
	HostRootModified                        bool     `json:"host_root_modified"`
	PrivilegedContainerRequired             bool     `json:"privileged_container_required"`
	HostNetworkingRequired                  bool     `json:"host_networking_required"`
	DockerSocketMounted                     bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired                  bool     `json:"broad_host_mount_required"`
	PlanReady                               bool     `json:"plan_ready"`
	BlockedActions                          []string `json:"blocked_actions"`
	DesktopSafeSummary                      string   `json:"desktop_safe_summary"`
}

func PreviewQ4ExternalWinAppRunPlan(request Q4ExternalWinAppRunPlanRequest) (Q4ExternalWinAppRunPlanPreview, error) {
	version := strings.TrimSpace(request.Version)
	if version == "" {
		return Q4ExternalWinAppRunPlanPreview{}, errors.New("q4 external Windows app run plan requires version metadata")
	}
	localExecutable := strings.TrimSpace(request.LocalExecutable)
	if localExecutable == "" {
		return Q4ExternalWinAppRunPlanPreview{}, errors.New("q4 external Windows app run plan requires --executable")
	}
	appID, err := requiredSingleLine("app id", request.AppID)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}
	displayName, err := requiredSingleLine("display name", request.DisplayName)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}
	windowMatch, err := requiredSingleLine("window match", request.WindowMatch)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}
	remoteHost, err := requiredSingleLine("remote host", request.RemoteHost)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}
	remoteRoot := strings.TrimSpace(request.RemoteMaterialsRoot)
	remoteRootClass, err := q4RemoteMaterialsRootClass(remoteRoot)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}
	localPathClass, err := q4ExternalRunLocalPathClass(strings.TrimSpace(request.ProjectRoot), localExecutable)
	if err != nil {
		return Q4ExternalWinAppRunPlanPreview{}, err
	}

	command := []string{
		"ruby",
		"scripts/q4_external_winapp_run.rb",
		"--executable", localExecutable,
		"--app-id", appID,
		"--display-name", displayName,
		"--window-match", windowMatch,
		"--remote", remoteHost,
		"--remote-materials-root", remoteRoot,
	}
	if output := strings.TrimSpace(request.Output); output != "" {
		command = append(command, "--output", output)
	}
	if markdownOutput := strings.TrimSpace(request.MarkdownOutput); markdownOutput != "" {
		command = append(command, "--markdown-output", markdownOutput)
	}
	if request.Execute {
		command = append(command, "--execute")
	}

	return Q4ExternalWinAppRunPlanPreview{
		Version:                                 version,
		SchemaVersion:                           Q4ExternalWinAppRunPlanSchemaVersion,
		RequestType:                             Q4ExternalWinAppRunPlanRequestType,
		Source:                                  "runtime-q4-external-winapp-run-plan+upload-staged-smoke",
		RuntimeMethod:                           "PlanQ4ExternalWinAppRun",
		ReadMethod:                              "GetQ4ExternalWinAppRunPlanPreview",
		AppID:                                   appID,
		DisplayName:                             displayName,
		WindowMatchConfigured:                   true,
		WindowMatchRequired:                     true,
		WindowMatch:                             windowMatch,
		LocalExecutableConfigured:               true,
		LocalExecutablePathClass:                localPathClass,
		LocalExecutablePathExposedToDesktop:     false,
		LocalExecutableMZHeaderRequired:         true,
		LocalExecutableSHA256Required:           true,
		RemoteHostConfigured:                    true,
		RemoteHostExposedToDesktop:              false,
		RemoteMaterialsRootClass:                remoteRootClass,
		RemoteExecutableGeneratedFromDigest:     true,
		RemoteExecutablePathExposedToDesktop:    false,
		RemoteUploadRequired:                    true,
		RemoteExecutionRequired:                 true,
		Q4MaterializationRequired:               true,
		Q4BuildPreferred:                        true,
		Q4ExecutionRequired:                     true,
		HostCompilationRequired:                 false,
		HostCompilationAvoided:                  true,
		LocalHostRole:                           "scoped-upload-and-operator-plan-only",
		DelegatedScript:                         "scripts/q4_external_winapp_run.rb",
		DelegatedRequestType:                    "q4-external-winapp-run",
		DelegatedAcceptanceRequestType:          Q4StagedExternalWinAppAcceptanceRequestType,
		DelegatedCommand:                        command,
		DelegatedCommandExposedToOperator:       true,
		DelegatedCommandExposedToDesktop:        false,
		StagedExternalFixtureRequired:           true,
		RealFileOpenEvidenceRequired:            true,
		WindowsProcessWindowObservationRequired: true,
		AcceptedApplicationDetailStateRequired:  "runtime-accepted-real-app-run",
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		FullSmokeRequired:                       false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		HostNetworkingRequired:                  false,
		DockerSocketMounted:                     false,
		BroadHostMountRequired:                  false,
		PlanReady:                               true,
		BlockedActions: []string{
			"Do not compile or launch the Windows app on the local host.",
			"Do not accept executables outside the checkout or /tmp/xnix-* local staging roots.",
			"Do not expose executable paths, remote hosts, or backend details to KDE-facing summaries.",
			"Do not run privileged containers, host networking, Docker socket mounts, or broad host mounts.",
		},
		DesktopSafeSummary: "External Windows app execution is planned for q4 through the Runtime-owned staged desktop lane; local host work is limited to scoped upload orchestration.",
	}, nil
}

func requiredSingleLine(label string, value string) (string, error) {
	clean := strings.TrimSpace(value)
	if clean == "" {
		return "", fmt.Errorf("%s must be non-empty", label)
	}
	if strings.ContainsAny(clean, "\r\n") {
		return "", fmt.Errorf("%s must be single-line", label)
	}
	return clean, nil
}

func q4ExternalRunLocalPathClass(projectRoot string, executable string) (string, error) {
	if strings.HasPrefix(executable, "/tmp/xnix-") {
		return "scoped-temporary", nil
	}
	if filepath.IsAbs(executable) {
		cleanExecutable := filepath.Clean(executable)
		if projectRoot != "" {
			cleanRoot, err := filepath.Abs(projectRoot)
			if err == nil && (cleanExecutable == cleanRoot || strings.HasPrefix(cleanExecutable, cleanRoot+string(filepath.Separator))) {
				return "checkout", nil
			}
		}
		return "", errors.New("local executable must stay under this checkout or /tmp/xnix-*")
	}
	if strings.HasPrefix(executable, ".."+string(filepath.Separator)) || executable == ".." {
		return "", errors.New("relative executable must stay inside this checkout")
	}
	return "checkout-relative", nil
}

func q4RemoteMaterialsRootClass(remoteRoot string) (string, error) {
	clean := filepath.Clean(remoteRoot)
	switch {
	case strings.HasPrefix(clean, "/home/xnix-"):
		return "q4-home-scoped", nil
	case strings.HasPrefix(clean, "/tmp/xnix-"):
		return "q4-temporary-scoped", nil
	default:
		return "", errors.New("remote materials root must stay under /home/xnix-* or /tmp/xnix-* on q4")
	}
}
