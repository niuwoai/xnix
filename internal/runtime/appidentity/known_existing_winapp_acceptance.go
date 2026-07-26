package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	KnownExistingWinAppAcceptanceSchemaVersion = "xnix.runtime.known_existing_winapp_acceptance.v1"
	KnownExistingWinAppAcceptanceRequestType   = "known-existing-winapp-acceptance-preview"
)

type KnownExistingWinAppAcceptanceRequest struct {
	RunReportPath string
}

type KnownExistingWinAppAcceptance struct {
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
	GuestEndpointExposed                 bool   `json:"guest_endpoint_exposed"`
	RawPathExposed                       bool   `json:"raw_path_exposed"`
	RawOutputExposed                     bool   `json:"raw_output_exposed"`
	RuntimeArgvExposed                   bool   `json:"runtime_argv_exposed"`
	RunnerPathExposed                    bool   `json:"runner_path_exposed"`
	AppID                                string `json:"app_id"`
	DisplayName                          string `json:"display_name"`
	AppVersion                           string `json:"app_version"`
	Architecture                         string `json:"architecture"`
	ExistingWindowsApp                   bool   `json:"existing_windows_app"`
	KnownPortableCatalogBacked           bool   `json:"known_portable_catalog_backed"`
	BackendClass                         string `json:"backend_class"`
	LaunchAttempted                      bool   `json:"launch_attempted"`
	RunnerAvailable                      bool   `json:"runner_available"`
	ChecksumVerified                     bool   `json:"checksum_verified"`
	ArtifactCopied                       bool   `json:"artifact_copied"`
	MarkerObserved                       bool   `json:"marker_observed"`
	GuestReachable                       bool   `json:"guest_reachable"`
	RuntimeStartedIsolatedGuest          bool   `json:"runtime_started_isolated_guest"`
	IsolatedGuestExecutionObserved       bool   `json:"isolated_guest_execution_observed"`
	CompatibilityEngineExecutionObserved bool   `json:"compatibility_engine_execution_observed"`
	LoopbackOnlyNetworking               bool   `json:"loopback_only_networking"`
	SerialLogPersisted                   bool   `json:"serial_log_persisted"`
	OutputRedacted                       bool   `json:"output_redacted"`
	StdoutBytes                          int    `json:"stdout_bytes"`
	StderrBytes                          int    `json:"stderr_bytes"`
	StdoutLineCount                      int    `json:"stdout_line_count"`
	StderrLineCount                      int    `json:"stderr_line_count"`
	NetworkRequired                      bool   `json:"network_required"`
	HostRootModified                     bool   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool   `json:"privileged_container_required"`
	HostNetworkingRequired               bool   `json:"host_networking_required"`
	DockerSocketMounted                  bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool   `json:"broad_host_mount_required"`
	DockerExecuted                       bool   `json:"docker_executed"`
	ColimaExecuted                       bool   `json:"colima_executed"`
	NetworkChecksRun                     bool   `json:"network_checks_run"`
	PackageManagerInvoked                bool   `json:"package_manager_invoked"`
	AcceptanceReady                      bool   `json:"acceptance_ready"`
	DesktopSafeSummary                   string `json:"desktop_safe_summary"`
}

func PreviewKnownExistingWinAppAcceptance(request KnownExistingWinAppAcceptanceRequest) (KnownExistingWinAppAcceptance, error) {
	path := strings.TrimSpace(request.RunReportPath)
	if path == "" {
		return KnownExistingWinAppAcceptance{}, errors.New("known existing Windows app acceptance requires --known-winapp-run")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return KnownExistingWinAppAcceptance{}, fmt.Errorf("read known existing Windows app run report: %w", err)
	}
	return PreviewKnownExistingWinAppAcceptanceJSON(content)
}

func PreviewKnownExistingWinAppAcceptanceJSON(content []byte) (KnownExistingWinAppAcceptance, error) {
	var report map[string]any
	if err := json.Unmarshal(content, &report); err != nil {
		return KnownExistingWinAppAcceptance{}, fmt.Errorf("parse known existing Windows app run report: %w", err)
	}
	if remoteString(report, "schema_version") != "xnix.runtime.known_windows_app_run.v1" ||
		remoteString(report, "request_type") != "windows-known-app-run" {
		return KnownExistingWinAppAcceptance{}, errors.New("known existing Windows app acceptance requires windows-known-app-run output")
	}
	if remoteString(report, "status") != "passed" {
		return KnownExistingWinAppAcceptance{}, errors.New("known existing Windows app acceptance requires a passed run report")
	}

	guestPayload := remoteMap(report, "guest_payload")
	guestReport := remoteMap(guestPayload, "guest")
	appID := remoteString(report, "app_id")
	displayName := remoteString(report, "display_name")
	appVersion := remoteString(report, "app_version")
	architecture := remoteString(report, "architecture")
	safeKnownExistingApp := knownExistingWindowsAppID(appID)
	topLevelSafetyClosed := remoteBool(report, "loopback_only_networking") &&
		!remoteBool(report, "network_required") &&
		!remoteBool(report, "host_root_modified") &&
		!remoteBool(report, "privileged_container_required") &&
		!remoteBool(report, "host_networking_required") &&
		!remoteBool(report, "docker_socket_mounted") &&
		!remoteBool(report, "broad_host_mount_required") &&
		!remoteBool(report, "docker_executed") &&
		!remoteBool(report, "colima_executed") &&
		!remoteBool(report, "network_checks_run") &&
		!remoteBool(report, "package_manager_invoked") &&
		!remoteBool(report, "raw_host_path_exposed") &&
		!remoteBool(report, "raw_executable_path_exposed") &&
		!remoteBool(report, "raw_profile_path_exposed") &&
		!remoteBool(report, "raw_state_root_path_exposed") &&
		!remoteBool(report, "raw_runtime_argv_exposed") &&
		!remoteBool(report, "raw_runner_path_exposed") &&
		!remoteBool(report, "raw_qemu_path_exposed")
	guestEvidenceReady := remoteString(guestPayload, "schema_version") == "xnix.runtime.known_windows_app_guest_wine_smoke.v1" &&
		remoteString(guestPayload, "request_type") == "windows-known-app-guest-wine-smoke" &&
		remoteString(guestPayload, "status") == "passed" &&
		remoteBool(guestPayload, "checksum_verified") &&
		remoteString(guestReport, "schema_version") == "xnix.runtime.windows_app_guest_wine_smoke.v1" &&
		remoteString(guestReport, "status") == "passed" &&
		remoteBool(guestReport, "guest_reachable") &&
		remoteBool(guestReport, "wine_available") &&
		remoteBool(guestReport, "executable_copied") &&
		remoteBool(guestReport, "marker_observed") &&
		remoteInt(guestReport, "exit_code") == 0 &&
		remoteBool(guestReport, "raw_output_redacted") &&
		!remoteBool(guestReport, "raw_output_included") &&
		remoteBool(guestReport, "loopback_only_networking") &&
		!remoteBool(guestReport, "host_root_modified") &&
		!remoteBool(guestReport, "privileged_container_required") &&
		!remoteBool(guestReport, "host_networking_required") &&
		!remoteBool(guestReport, "docker_socket_mounted") &&
		!remoteBool(guestReport, "broad_host_mount_required") &&
		!remoteBool(guestReport, "raw_host_path_exposed")
	acceptanceReady := safeKnownExistingApp &&
		displayName != "" &&
		appVersion != "" &&
		architecture != "" &&
		remoteBool(report, "backend_ready") &&
		remoteBool(report, "launch_attempted") &&
		remoteBool(report, "runner_available") &&
		remoteBool(report, "checksum_verified") &&
		remoteBool(report, "guest_start_attempted") &&
		remoteBool(report, "guest_started") &&
		remoteBool(report, "qemu_serial_log_written") &&
		remoteBool(report, "executable_copied") &&
		remoteBool(report, "marker_observed") &&
		remoteBool(report, "raw_output_redacted") &&
		guestEvidenceReady &&
		topLevelSafetyClosed
	if !acceptanceReady {
		return KnownExistingWinAppAcceptance{}, errors.New("known existing Windows app acceptance requires executed, checksum-verified, marker-observed known portable app evidence with isolated guest execution and closed safety gates")
	}

	acceptance := KnownExistingWinAppAcceptance{
		Version:                              remoteString(report, "version"),
		SchemaVersion:                        KnownExistingWinAppAcceptanceSchemaVersion,
		RequestType:                          KnownExistingWinAppAcceptanceRequestType,
		Source:                               "known-existing-windows-app-run+go-runtime-acceptance",
		RuntimeMethod:                        "PreviewKnownExistingWinAppAcceptance",
		ReadMethod:                           "GetKnownExistingWinAppAcceptance",
		AcceptanceType:                       "known-existing-windows-app-real-run-acceptance",
		RunReportConsumed:                    true,
		RunReportPathExposed:                 false,
		RemoteHostExposed:                    false,
		GuestEndpointExposed:                 false,
		RawPathExposed:                       false,
		RawOutputExposed:                     false,
		RuntimeArgvExposed:                   false,
		RunnerPathExposed:                    false,
		AppID:                                appID,
		DisplayName:                          displayName,
		AppVersion:                           appVersion,
		Architecture:                         architecture,
		ExistingWindowsApp:                   true,
		KnownPortableCatalogBacked:           true,
		BackendClass:                         "managed-isolated-compatibility",
		LaunchAttempted:                      true,
		RunnerAvailable:                      true,
		ChecksumVerified:                     true,
		ArtifactCopied:                       true,
		MarkerObserved:                       true,
		GuestReachable:                       true,
		RuntimeStartedIsolatedGuest:          true,
		IsolatedGuestExecutionObserved:       true,
		CompatibilityEngineExecutionObserved: true,
		LoopbackOnlyNetworking:               true,
		SerialLogPersisted:                   true,
		OutputRedacted:                       true,
		StdoutBytes:                          remoteInt(guestReport, "stdout_bytes"),
		StderrBytes:                          remoteInt(guestReport, "stderr_bytes"),
		StdoutLineCount:                      remoteInt(guestReport, "stdout_line_count"),
		StderrLineCount:                      remoteInt(guestReport, "stderr_line_count"),
		NetworkRequired:                      false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		DockerExecuted:                       false,
		ColimaExecuted:                       false,
		NetworkChecksRun:                     false,
		PackageManagerInvoked:                false,
		AcceptanceReady:                      true,
		DesktopSafeSummary:                   "A known third-party Windows app completed the isolated Runtime-owned acceptance lane.",
	}
	if err := validateNoBackendTerms(acceptance, "known existing Windows app acceptance"); err != nil {
		return KnownExistingWinAppAcceptance{}, err
	}
	return acceptance, nil
}

func knownExistingWindowsAppID(appID string) bool {
	switch appID {
	case "7zr", "busybox-w32":
		return true
	default:
		return false
	}
}

func remoteMap(values map[string]any, key string) map[string]any {
	raw, ok := values[key]
	if !ok {
		return map[string]any{}
	}
	nested, ok := raw.(map[string]any)
	if !ok {
		return map[string]any{}
	}
	return nested
}
