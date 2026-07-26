package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	DesktopExternalWinAppLaunchPacketSchemaVersion = "xnix.runtime.desktop_external_winapp_launch_packet.v1"
	DesktopExternalWinAppLaunchPacketRequestType   = "desktop-external-winapp-launch-packet-preview"
)

type DesktopExternalWinAppLaunchPacket struct {
	Version                          string   `json:"version"`
	SchemaVersion                    string   `json:"schema_version"`
	RequestType                      string   `json:"request_type"`
	PacketType                       string   `json:"packet_type"`
	Status                           string   `json:"status"`
	Source                           string   `json:"source"`
	RuntimeMethod                    string   `json:"runtime_method"`
	ReadMethod                       string   `json:"read_method"`
	Desktop                          string   `json:"desktop"`
	ApplicationID                    string   `json:"application_id"`
	DisplayName                      string   `json:"display_name"`
	AppVersion                       string   `json:"app_version,omitempty"`
	ExternalAppHandle                string   `json:"external_app_handle"`
	ActivationStatusRequestType      string   `json:"activation_status_request_type"`
	ActivationReceiptBacked          bool     `json:"activation_receipt_backed"`
	ActivationReceiptSafeForKDE      bool     `json:"activation_receipt_safe_for_kde"`
	DesktopExecUsesExternalAppHandle bool     `json:"desktop_exec_uses_external_app_handle"`
	ExternalAppDesktopHandleReady    bool     `json:"external_app_desktop_handle_ready"`
	DesktopExecUsesRawImportRecord   bool     `json:"desktop_exec_uses_raw_import_record"`
	DesktopExecUsesStateRoot         bool     `json:"desktop_exec_uses_state_root"`
	RunRecordConsumed                bool     `json:"run_record_consumed"`
	RunRecordRequestType             string   `json:"run_record_request_type"`
	ExternalAppRunRecordConsumed     bool     `json:"external_app_run_record_consumed"`
	ExternalAppImportRecordConsumed  bool     `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool     `json:"external_app_handle_consumed"`
	ImportedArtifactDigestVerified   bool     `json:"imported_artifact_digest_verified"`
	RuntimeRunRequested              bool     `json:"runtime_run_requested"`
	RuntimeLaunchExecuted            bool     `json:"runtime_launch_executed"`
	ExecutionStarted                 bool     `json:"execution_started"`
	BackendProcessStarted            bool     `json:"backend_process_started"`
	WindowObserved                   bool     `json:"window_observed"`
	XWindowObserved                  bool     `json:"x_window_observed"`
	ContainerRuntimeUsed             bool     `json:"container_runtime_used"`
	ContainerNetworkMode             string   `json:"container_network_mode"`
	ContainerHostMountCount          int      `json:"container_host_mount_count"`
	RuntimeOwned                     bool     `json:"runtime_owned"`
	GoRuntimeBacked                  bool     `json:"go_runtime_backed"`
	RuntimeLaunchAuthority           bool     `json:"runtime_launch_authority"`
	KDEPolicyOwner                   bool     `json:"kde_policy_owner"`
	KDELaunchAuthority               bool     `json:"kde_launch_authority"`
	DesktopLaunchPacketReady         bool     `json:"desktop_launch_packet_ready"`
	SafeForKDE                       bool     `json:"safe_for_kde"`
	UnsafeReasonIDs                  []string `json:"unsafe_reason_ids"`
	BackendDetailsExposed            bool     `json:"backend_details_exposed"`
	RawImportRecordPathExposed       bool     `json:"raw_import_record_path_exposed"`
	RawExternalAppHandlePathExposed  bool     `json:"raw_external_app_handle_path_exposed"`
	RawStateRootPathExposed          bool     `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed         bool     `json:"raw_executable_path_exposed"`
	HostRootModified                 bool     `json:"host_root_modified"`
	PrivilegedContainerRequired      bool     `json:"privileged_container_required"`
	HostNetworkingRequired           bool     `json:"host_networking_required"`
	DockerSocketMounted              bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool     `json:"broad_host_mount_required"`
	DesktopSafeSummary               string   `json:"desktop_safe_summary"`
}

func (plan Plan) DesktopExternalWinAppLaunchPacketPreview(activationRoot string, runRecordPath string, mode string) (DesktopExternalWinAppLaunchPacket, error) {
	if strings.TrimSpace(activationRoot) == "" {
		return DesktopExternalWinAppLaunchPacket{}, errors.New("desktop external Windows app launch packet requires --activation-root")
	}
	if strings.TrimSpace(runRecordPath) == "" {
		return DesktopExternalWinAppLaunchPacket{}, errors.New("desktop external Windows app launch packet requires --run-record")
	}
	status, err := plan.DesktopActivationStatusPreviewWithReceipt(activationRoot, mode)
	if err != nil {
		return DesktopExternalWinAppLaunchPacket{}, err
	}
	receipt := status.ReceiptEvidence
	if !status.ReceiptBacked || !receipt.SafeForKDE {
		return DesktopExternalWinAppLaunchPacket{}, errors.New("desktop external Windows app launch packet requires safe receipt-backed activation status")
	}
	if strings.TrimSpace(receipt.ExternalAppHandle) != plan.ApplicationID {
		return DesktopExternalWinAppLaunchPacket{}, errors.New("desktop external Windows app launch packet receipt handle does not match the application id")
	}
	runRecord, err := loadExternalWinAppRunRecord(runRecordPath)
	if err != nil {
		return DesktopExternalWinAppLaunchPacket{}, err
	}
	if runRecord.SchemaVersion != ExternalWinAppRunSchemaVersion ||
		runRecord.RequestType != ExternalWinAppRunRequestType {
		return DesktopExternalWinAppLaunchPacket{}, fmt.Errorf("unexpected external Windows app run record schema: %s/%s", runRecord.SchemaVersion, runRecord.RequestType)
	}
	if runRecord.ApplicationID != plan.ApplicationID {
		return DesktopExternalWinAppLaunchPacket{}, fmt.Errorf("external Windows app run record application mismatch: %s", runRecord.ApplicationID)
	}
	if runRecord.DisplayName != plan.DisplayName {
		return DesktopExternalWinAppLaunchPacket{}, fmt.Errorf("external Windows app run record display name mismatch: %s", runRecord.DisplayName)
	}
	if runRecord.ExternalAppHandleConsumed && strings.TrimSpace(runRecord.ExternalAppHandle) != plan.ApplicationID {
		return DesktopExternalWinAppLaunchPacket{}, errors.New("external Windows app run record handle does not match the application id")
	}

	reasons := desktopExternalWinAppLaunchUnsafeReasons(receipt, runRecord)
	packetReady := len(reasons) == 0
	packet := DesktopExternalWinAppLaunchPacket{
		Version:                          runRecord.Version,
		SchemaVersion:                    DesktopExternalWinAppLaunchPacketSchemaVersion,
		RequestType:                      DesktopExternalWinAppLaunchPacketRequestType,
		PacketType:                       "kde-desktop-external-winapp-launch-evidence",
		Status:                           desktopExternalWinAppLaunchPacketStatus(packetReady),
		Source:                           "desktop-activation-status-preview+external-app-run-record",
		RuntimeMethod:                    "GetDesktopExternalWinAppLaunchPacket",
		ReadMethod:                       "desktop-external-winapp-launch-packet-preview",
		Desktop:                          "KDE Plasma",
		ApplicationID:                    plan.ApplicationID,
		DisplayName:                      plan.DisplayName,
		AppVersion:                       runRecord.AppVersion,
		ExternalAppHandle:                receipt.ExternalAppHandle,
		ActivationStatusRequestType:      status.RequestType,
		ActivationReceiptBacked:          status.ReceiptBacked,
		ActivationReceiptSafeForKDE:      receipt.SafeForKDE,
		DesktopExecUsesExternalAppHandle: receipt.DesktopExecUsesExternalAppHandle,
		ExternalAppDesktopHandleReady:    receipt.ExternalAppDesktopHandleReady,
		DesktopExecUsesRawImportRecord:   receipt.DesktopExecUsesRawImportRecord,
		DesktopExecUsesStateRoot:         receipt.DesktopExecUsesStateRoot,
		RunRecordConsumed:                true,
		RunRecordRequestType:             runRecord.RequestType,
		ExternalAppRunRecordConsumed:     true,
		ExternalAppImportRecordConsumed:  runRecord.ExternalAppImportRecordConsumed,
		ExternalAppHandleConsumed:        runRecord.ExternalAppHandleConsumed,
		ImportedArtifactDigestVerified:   runRecord.ImportedArtifactDigestVerified,
		RuntimeRunRequested:              runRecord.RuntimeRunRequested,
		RuntimeLaunchExecuted:            runRecord.RuntimeRunExecuted,
		ExecutionStarted:                 runRecord.ExecutionStarted,
		BackendProcessStarted:            runRecord.BackendProcessStarted,
		WindowObserved:                   runRecord.WindowObserved,
		XWindowObserved:                  runRecord.XWindowObserved,
		ContainerRuntimeUsed:             runRecord.ContainerRuntimeUsed,
		ContainerNetworkMode:             runRecord.ContainerNetworkMode,
		ContainerHostMountCount:          runRecord.ContainerHostMountCount,
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		RuntimeLaunchAuthority:           true,
		KDEPolicyOwner:                   false,
		KDELaunchAuthority:               false,
		DesktopLaunchPacketReady:         packetReady,
		SafeForKDE:                       packetReady,
		UnsafeReasonIDs:                  reasons,
		BackendDetailsExposed:            runRecord.BackendDetailsExposed,
		RawImportRecordPathExposed:       runRecord.RawImportRecordPathExposed,
		RawExternalAppHandlePathExposed:  runRecord.RawExternalAppHandlePathExposed,
		RawStateRootPathExposed:          runRecord.RawStateRootPathExposed,
		RawExecutablePathExposed:         runRecord.RawExecutablePathExposed,
		HostRootModified:                 runRecord.HostRootModified,
		PrivilegedContainerRequired:      runRecord.PrivilegedContainerRequired,
		HostNetworkingRequired:           runRecord.HostNetworkingRequired,
		DockerSocketMounted:              runRecord.DockerSocketMounted,
		BroadHostMountRequired:           runRecord.BroadHostMountRequired,
		DesktopSafeSummary:               desktopExternalWinAppLaunchSummary(packetReady),
	}
	if packet.Version == "" {
		packet.Version = plan.ApplicationVersion
	}
	if err := validateNoBackendTerms(packet, "desktop external Windows app launch packet"); err != nil {
		return DesktopExternalWinAppLaunchPacket{}, err
	}
	return packet, nil
}

func loadExternalWinAppRunRecord(path string) (ExternalWinAppRunResult, error) {
	data, err := os.ReadFile(strings.TrimSpace(path))
	if err != nil {
		return ExternalWinAppRunResult{}, fmt.Errorf("read external Windows app run record: %w", err)
	}
	var record ExternalWinAppRunResult
	if err := json.Unmarshal(data, &record); err != nil {
		return ExternalWinAppRunResult{}, fmt.Errorf("parse external Windows app run record: %w", err)
	}
	return record, nil
}

func desktopExternalWinAppLaunchUnsafeReasons(receipt DesktopActivationStatusReceiptEvidence, runRecord ExternalWinAppRunResult) []string {
	reasons := []string{}
	if !receipt.DesktopExecUsesExternalAppHandle {
		reasons = append(reasons, "desktop-exec-not-handle-routed")
	}
	if !receipt.ExternalAppDesktopHandleReady {
		reasons = append(reasons, "external-app-desktop-handle-not-ready")
	}
	if receipt.DesktopExecUsesRawImportRecord {
		reasons = append(reasons, "desktop-exec-uses-raw-import-record")
	}
	if receipt.DesktopExecUsesStateRoot {
		reasons = append(reasons, "desktop-exec-uses-state-root")
	}
	if runRecord.Status != winapp.PassedStatus {
		reasons = append(reasons, "external-app-run-not-passed")
	}
	if !runRecord.ExternalAppImportRecordConsumed {
		reasons = append(reasons, "external-app-import-record-not-consumed")
	}
	if !runRecord.ExternalAppHandleConsumed {
		reasons = append(reasons, "external-app-handle-not-consumed")
	}
	if !runRecord.ImportedArtifactDigestVerified {
		reasons = append(reasons, "imported-artifact-digest-not-verified")
	}
	if !runRecord.RuntimeRunRequested {
		reasons = append(reasons, "runtime-run-not-requested")
	}
	if !runRecord.RuntimeRunExecuted {
		reasons = append(reasons, "runtime-run-not-executed")
	}
	if !runRecord.ExecutionStarted {
		reasons = append(reasons, "execution-not-started")
	}
	if !runRecord.BackendProcessStarted {
		reasons = append(reasons, "backend-process-not-started")
	}
	if !runRecord.WindowObserved {
		reasons = append(reasons, "window-not-observed")
	}
	if !runRecord.XWindowObserved {
		reasons = append(reasons, "x-window-not-observed")
	}
	if !runRecord.ContainerRuntimeUsed {
		reasons = append(reasons, "container-runtime-not-used")
	}
	if runRecord.ContainerNetworkMode != "none" {
		reasons = append(reasons, "container-network-not-isolated")
	}
	if runRecord.ContainerHostMountCount != 0 {
		reasons = append(reasons, "container-host-mounts-present")
	}
	if runRecord.BackendDetailsExposed {
		reasons = append(reasons, "backend-details-exposed")
	}
	if runRecord.RawImportRecordPathExposed {
		reasons = append(reasons, "raw-import-record-path-exposed")
	}
	if runRecord.RawExternalAppHandlePathExposed {
		reasons = append(reasons, "raw-external-app-handle-path-exposed")
	}
	if runRecord.RawStateRootPathExposed {
		reasons = append(reasons, "raw-state-root-path-exposed")
	}
	if runRecord.RawExecutablePathExposed {
		reasons = append(reasons, "raw-executable-path-exposed")
	}
	if runRecord.HostRootModified {
		reasons = append(reasons, "host-root-modified")
	}
	if runRecord.PrivilegedContainerRequired {
		reasons = append(reasons, "privileged-container-required")
	}
	if runRecord.HostNetworkingRequired {
		reasons = append(reasons, "host-networking-required")
	}
	if runRecord.DockerSocketMounted {
		reasons = append(reasons, "docker-socket-mounted")
	}
	if runRecord.BroadHostMountRequired {
		reasons = append(reasons, "broad-host-mount-required")
	}
	return reasons
}

func desktopExternalWinAppLaunchPacketStatus(ready bool) string {
	if ready {
		return "passed"
	}
	return "blocked"
}

func desktopExternalWinAppLaunchSummary(ready bool) string {
	if ready {
		return "KDE can show that the imported external application launched through the Runtime-owned desktop handle and produced observed-window evidence without exposing raw paths or owning backend policy."
	}
	return "KDE can show that imported external application desktop launch evidence is incomplete while Runtime keeps raw paths, backend policy, and host mutation closed."
}
