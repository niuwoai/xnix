package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	RealWinAppGUIEvidencePacketSchemaVersion = "xnix.runtime.real_winapp_gui_evidence_packet.v1"
	RealWinAppGUIEvidencePacketRequestType   = "real-winapp-gui-evidence-packet-preview"
)

type RealWinAppGUIEvidencePacketRequest struct {
	ReportPath  string
	AppID       string
	DisplayName string
	AppVersion  string
}

type RealWinAppGUIEvidencePacket struct {
	Version                            string                       `json:"version"`
	SchemaVersion                      string                       `json:"schema_version"`
	RequestType                        string                       `json:"request_type"`
	PacketType                         string                       `json:"packet_type"`
	Source                             string                       `json:"source"`
	RuntimeMethod                      string                       `json:"runtime_method"`
	ReadMethod                         string                       `json:"read_method"`
	ReportStatus                       string                       `json:"report_status"`
	ReportConsumed                     bool                         `json:"report_consumed"`
	ReportPathExposed                  bool                         `json:"report_path_exposed"`
	AppID                              string                       `json:"app_id"`
	DisplayName                        string                       `json:"display_name"`
	AppVersion                         string                       `json:"app_version"`
	GUIAppName                         string                       `json:"gui_app_name"`
	EvidenceSource                     string                       `json:"evidence_source"`
	RecipeBacked                       bool                         `json:"recipe_backed"`
	RecipeAppID                        string                       `json:"recipe_app_id,omitempty"`
	CompatibilityState                 string                       `json:"compatibility_state"`
	CenterCardState                    string                       `json:"center_card_state"`
	KnownAppGUIEvidenceCount           int                          `json:"known_app_gui_evidence_count"`
	KnownAppGUIEvidenceVerifiedCount   int                          `json:"known_app_gui_evidence_verified_count"`
	KnownAppSmokeEvidence              KnownAppSmokeEvidenceSummary `json:"known_app_smoke_evidence"`
	WinebootInvoked                    bool                         `json:"wineboot_invoked"`
	XWindowObserved                    bool                         `json:"x_window_observed"`
	WindowObserved                     bool                         `json:"window_observed"`
	XWindowChildCount                  int                          `json:"x_window_child_count"`
	CompatibilityCenterProjectionReady bool                         `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                         `json:"kde_center_projection_ready"`
	ContainerRuntimeUsed               bool                         `json:"container_runtime_used"`
	ContainerNetworkMode               string                       `json:"container_network_mode,omitempty"`
	ContainerHostMountCount            int                          `json:"container_host_mount_count"`
	ExecutableName                     string                       `json:"executable_name,omitempty"`
	LocalExecutableCopied              bool                         `json:"local_executable_copied"`
	ExternalAppRunRecordConsumed       bool                         `json:"external_app_run_record_consumed"`
	ExternalAppHandleConsumed          bool                         `json:"external_app_handle_consumed"`
	ExternalAppImportRecordConsumed    bool                         `json:"external_app_import_record_consumed"`
	ImportedArtifactDigestVerified     bool                         `json:"imported_artifact_digest_verified"`
	ImportedArtifactSHA256             string                       `json:"imported_artifact_sha256,omitempty"`
	RuntimeOwned                       bool                         `json:"runtime_owned"`
	GoRuntimeBacked                    bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                         `json:"kde_policy_owner"`
	DesktopLaunchEnabled               bool                         `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool                         `json:"backend_launch_enabled"`
	ActionExecutionEnabled             bool                         `json:"action_execution_enabled"`
	BackendDetailsExposed              bool                         `json:"backend_details_exposed"`
	RawOutputExposed                   bool                         `json:"raw_output_exposed"`
	HostRootModified                   bool                         `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                         `json:"privileged_container_required"`
	HostNetworkingRequired             bool                         `json:"host_networking_required"`
	DockerSocketMounted                bool                         `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool                         `json:"broad_host_mount_required"`
	DesktopSafeSummary                 string                       `json:"desktop_safe_summary"`
}

type containerXGUIRuntimePayload struct {
	SchemaVersion                            string `json:"schema_version"`
	RequestType                              string `json:"request_type"`
	Status                                   string `json:"status"`
	ApplicationID                            string `json:"application_id"`
	DisplayName                              string `json:"display_name"`
	AppVersion                               string `json:"app_version"`
	RecipeBacked                             bool   `json:"recipe_backed"`
	ExecutableName                           string `json:"executable_name"`
	LocalExecutableCopied                    bool   `json:"local_executable_copied"`
	ExternalAppRunRecordConsumed             bool   `json:"external_app_run_record_consumed"`
	ExternalAppHandleConsumed                bool   `json:"external_app_handle_consumed"`
	ExternalAppImportRecordConsumed          bool   `json:"external_app_import_record_consumed"`
	ImportedArtifactDigestVerified           bool   `json:"imported_artifact_digest_verified"`
	ImportedArtifactSHA256                   string `json:"imported_artifact_sha256"`
	ApplicationName                          string `json:"application_name"`
	WindowMatch                              string `json:"window_match"`
	ContainerImage                           string `json:"container_image"`
	ContainerPlatform                        string `json:"container_platform"`
	NetworkMode                              string `json:"network_mode"`
	XServerStarted                           bool   `json:"x_server_started"`
	WineBootstrapAttempted                   bool   `json:"wine_bootstrap_attempted"`
	ImageAvailable                           bool   `json:"image_available"`
	XWindowObserved                          bool   `json:"x_window_observed"`
	WindowEvidenceSummary                    string `json:"window_evidence_summary"`
	EvidenceSource                           string `json:"evidence_source"`
	DispatchStarted                          bool   `json:"dispatch_started"`
	ExecutionStarted                         bool   `json:"execution_started"`
	SmokePassed                              bool   `json:"smoke_passed"`
	RuntimeOwnedDispatch                     bool   `json:"runtime_owned_dispatch"`
	SessionGatedControlledDispatchConsumed   bool   `json:"session_gated_controlled_dispatch_consumed"`
	SessionGatedControlledDispatchState      string `json:"session_gated_controlled_dispatch_state"`
	SessionGatedReviewReceiptID              string `json:"session_gated_review_receipt_id"`
	LaunchAuthorizationReceiptID             string `json:"launch_authorization_receipt_id"`
	ControlledExecutionSessionConsumed       bool   `json:"controlled_execution_session_consumed"`
	ControlledExecutionSessionID             string `json:"controlled_execution_session_id"`
	ControlledSessionDigestVerified          bool   `json:"controlled_session_digest_verified"`
	ControlledSessionRelativePath            string `json:"controlled_session_relative_path"`
	RuntimeOwnerConsumableSession            bool   `json:"runtime_owner_consumable_session"`
	KDEReadModelConsumableSession            bool   `json:"kde_read_model_consumable_session"`
	ControlledSessionLiveStateObserved       bool   `json:"controlled_session_live_state_observed"`
	ControlledSessionRegistered              bool   `json:"controlled_session_registered"`
	ControlledSessionWindowObserved          bool   `json:"controlled_session_window_observed"`
	ControlledSessionHostRootModified        bool   `json:"controlled_session_host_root_modified"`
	ControlledSessionContainerProcessStarted bool   `json:"controlled_session_container_process_start"`
	RawCommandExposed                        bool   `json:"raw_command_exposed"`
	BackendDetailsExposed                    bool   `json:"backend_details_exposed"`
	HostRootModified                         bool   `json:"host_root_modified"`
	PrivilegedContainerRequired              bool   `json:"privileged_container_required"`
	HostNetworkingRequired                   bool   `json:"host_networking_required"`
	DockerSocketMounted                      bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired                   bool   `json:"broad_host_mount_required"`
	HostMountCount                           int    `json:"host_mount_count"`
}

func PreviewRealWinAppGUIEvidencePacket(request RealWinAppGUIEvidencePacketRequest) (RealWinAppGUIEvidencePacket, error) {
	path := strings.TrimSpace(request.ReportPath)
	if path == "" {
		return RealWinAppGUIEvidencePacket{}, errors.New("real Windows app GUI evidence packet requires --gui-smoke-report")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return RealWinAppGUIEvidencePacket{}, fmt.Errorf("read real Windows app GUI smoke report: %w", err)
	}
	return PreviewRealWinAppGUIEvidencePacketJSON(content, request)
}

func PreviewRealWinAppGUIEvidencePacketJSON(content []byte, request RealWinAppGUIEvidencePacketRequest) (RealWinAppGUIEvidencePacket, error) {
	var report guiSmokeReport
	if err := json.Unmarshal(content, &report); err != nil {
		return RealWinAppGUIEvidencePacket{}, fmt.Errorf("parse real Windows app GUI smoke report: %w", err)
	}
	projectionContent := content
	rawRuntimePayload := false
	var runtimePayload containerXGUIRuntimePayload
	if report.SchemaVersion == "xnix.runtime.windows_app_container_x_gui_smoke.v1" {
		payload, wrapped, err := wrapContainerXGUIRuntimePayload(content)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		rawRuntimePayload = true
		runtimePayload = payload
		report = wrapped
		projectionContent, err = json.Marshal(wrapped)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, fmt.Errorf("encode wrapped container X GUI Runtime payload: %w", err)
		}
	} else if report.SchemaVersion == ExternalWinAppRunSchemaVersion {
		payload, wrapped, err := wrapExternalWinAppRunPayload(content)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		rawRuntimePayload = true
		runtimePayload = payload
		report = wrapped
		projectionContent, err = json.Marshal(wrapped)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, fmt.Errorf("encode wrapped external Windows app run payload: %w", err)
		}
	} else if report.SchemaVersion == "xnix.runtime.external_winapp_import_stage_launch.v1" {
		runContent, err := unwrapExternalWinAppImportStageLaunchPayload(content)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		payload, wrapped, err := wrapExternalWinAppRunPayload(runContent)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		rawRuntimePayload = true
		runtimePayload = payload
		report = wrapped
		projectionContent, err = json.Marshal(wrapped)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, fmt.Errorf("encode wrapped one-shot external Windows app run payload: %w", err)
		}
	} else if report.SchemaVersion == "xnix.runtime.known_portable_bundle_stage_launch.v1" {
		stageLaunchContent, err := unwrapKnownPortableBundleStageLaunchPayload(content)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		runContent, err := unwrapExternalWinAppImportStageLaunchPayload(stageLaunchContent)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		payload, wrapped, err := wrapExternalWinAppRunPayload(runContent)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, err
		}
		rawRuntimePayload = true
		runtimePayload = payload
		report = wrapped
		projectionContent, err = json.Marshal(wrapped)
		if err != nil {
			return RealWinAppGUIEvidencePacket{}, fmt.Errorf("encode wrapped known portable bundle staged launch payload: %w", err)
		}
	}
	projectionRequest := GUISmokeEvidencePreviewRequest{
		AppID:       request.AppID,
		DisplayName: request.DisplayName,
		AppVersion:  request.AppVersion,
	}
	if rawRuntimePayload {
		if strings.TrimSpace(projectionRequest.AppID) == "" {
			projectionRequest.AppID = runtimePayload.ApplicationID
		}
		if strings.TrimSpace(projectionRequest.DisplayName) == "" {
			projectionRequest.DisplayName = runtimePayload.DisplayName
		}
		if strings.TrimSpace(projectionRequest.AppVersion) == "" {
			projectionRequest.AppVersion = runtimePayload.AppVersion
		}
	}
	projection, err := PreviewGUISmokeEvidenceJSON(projectionContent, projectionRequest)
	if err != nil {
		return RealWinAppGUIEvidencePacket{}, err
	}
	if !projection.ReportConsumed || projection.ReportStatus != "passed" {
		return RealWinAppGUIEvidencePacket{}, errors.New("real Windows app GUI evidence packet requires a consumed passed GUI report")
	}
	if !projection.CompatibilityCenterProjectionReady || !projection.KDECenterProjectionReady {
		return RealWinAppGUIEvidencePacket{}, errors.New("real Windows app GUI evidence packet requires Runtime and KDE projection readiness")
	}
	projectionPayload, err := json.Marshal(projection)
	if err != nil {
		return RealWinAppGUIEvidencePacket{}, fmt.Errorf("encode GUI smoke Runtime evidence projection: %w", err)
	}
	evidence, err := KnownAppSmokeEvidenceFromGUISmokeProjection(projectionPayload)
	if err != nil {
		return RealWinAppGUIEvidencePacket{}, err
	}
	if rawRuntimePayload {
		evidence = enhanceKnownAppEvidenceWithContainerRuntimePayload(evidence, runtimePayload)
	}

	version := strings.TrimSpace(report.Version)
	if version == "" {
		version = projection.AppVersion
	}
	verifiedCount := 0
	if evidence.SmokeStatus == "passed" && evidence.ExecutionEvidenceRecorded && evidence.RuntimeDispatchVerified {
		verifiedCount = 1
	}
	containerRuntimeUsed := evidence.EvidenceSource == GUISmokeEvidenceSourceContainerXGUI
	containerNetworkMode := ""
	containerHostMountCount := 0
	executableName := ""
	localExecutableCopied := false
	externalAppRunRecordConsumed := false
	externalAppHandleConsumed := false
	externalAppImportRecordConsumed := false
	importedArtifactDigestVerified := false
	importedArtifactSHA256 := ""
	if containerRuntimeUsed {
		containerNetworkMode = report.ContainerPayload.NetworkMode
		containerHostMountCount = report.ContainerPayload.HostMountCount
		executableName = runtimePayload.ExecutableName
		localExecutableCopied = runtimePayload.LocalExecutableCopied
		externalAppRunRecordConsumed = runtimePayload.ExternalAppRunRecordConsumed
		externalAppHandleConsumed = runtimePayload.ExternalAppHandleConsumed
		externalAppImportRecordConsumed = runtimePayload.ExternalAppImportRecordConsumed
		importedArtifactDigestVerified = runtimePayload.ImportedArtifactDigestVerified
		importedArtifactSHA256 = strings.TrimSpace(runtimePayload.ImportedArtifactSHA256)
	}

	return RealWinAppGUIEvidencePacket{
		Version:                            version,
		SchemaVersion:                      RealWinAppGUIEvidencePacketSchemaVersion,
		RequestType:                        RealWinAppGUIEvidencePacketRequestType,
		PacketType:                         "real-windows-app-gui-evidence",
		Source:                             projection.Source + "+real-winapp-desktop-packet",
		RuntimeMethod:                      "PreviewRealWinAppGUIEvidencePacket",
		ReadMethod:                         "GetRealWinAppGUIEvidencePacket",
		ReportStatus:                       projection.ReportStatus,
		ReportConsumed:                     projection.ReportConsumed,
		ReportPathExposed:                  false,
		AppID:                              projection.AppID,
		DisplayName:                        projection.DisplayName,
		AppVersion:                         projection.AppVersion,
		GUIAppName:                         projection.GUIAppName,
		EvidenceSource:                     evidence.EvidenceSource,
		RecipeBacked:                       evidence.RecipeBacked,
		RecipeAppID:                        evidence.RecipeAppID,
		CompatibilityState:                 evidence.CompatibilityState,
		CenterCardState:                    evidence.CenterCardState,
		KnownAppGUIEvidenceCount:           1,
		KnownAppGUIEvidenceVerifiedCount:   verifiedCount,
		KnownAppSmokeEvidence:              evidence,
		WinebootInvoked:                    projection.WinebootInvoked,
		XWindowObserved:                    projection.XWindowObserved,
		WindowObserved:                     projection.XWindowObserved,
		XWindowChildCount:                  projection.XWindowChildCount,
		CompatibilityCenterProjectionReady: projection.CompatibilityCenterProjectionReady,
		KDECenterProjectionReady:           projection.KDECenterProjectionReady,
		ContainerRuntimeUsed:               containerRuntimeUsed,
		ContainerNetworkMode:               containerNetworkMode,
		ContainerHostMountCount:            containerHostMountCount,
		ExecutableName:                     executableName,
		LocalExecutableCopied:              localExecutableCopied,
		ExternalAppRunRecordConsumed:       externalAppRunRecordConsumed,
		ExternalAppHandleConsumed:          externalAppHandleConsumed,
		ExternalAppImportRecordConsumed:    externalAppImportRecordConsumed,
		ImportedArtifactDigestVerified:     importedArtifactDigestVerified,
		ImportedArtifactSHA256:             importedArtifactSHA256,
		RuntimeOwned:                       projection.RuntimeOwned,
		GoRuntimeBacked:                    projection.GoRuntimeBacked,
		KDEPolicyOwner:                     projection.KDEPolicyOwner,
		DesktopLaunchEnabled:               false,
		BackendLaunchEnabled:               false,
		ActionExecutionEnabled:             false,
		BackendDetailsExposed:              false,
		RawOutputExposed:                   false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        projection.PrivilegedContainerRequired,
		HostNetworkingRequired:             projection.HostNetworkingRequired,
		DockerSocketMounted:                projection.DockerSocketMounted,
		BroadHostMountRequired:             projection.BroadHostMountRequired,
		DesktopSafeSummary:                 projection.DesktopSafeSummary,
	}, nil
}

type externalWinAppImportStageLaunchPayload struct {
	SchemaVersion                    string          `json:"schema_version"`
	RequestType                      string          `json:"request_type"`
	Status                           string          `json:"status"`
	ApplicationID                    string          `json:"application_id"`
	ExternalAppHandle                string          `json:"external_app_handle"`
	ImportRecorded                   bool            `json:"import_recorded"`
	DesktopActivationStaged          bool            `json:"desktop_activation_staged"`
	StagedLauncherInvoked            bool            `json:"staged_launcher_invoked"`
	StagedLauncherFromActivationRoot bool            `json:"staged_launcher_from_activation_root"`
	ManagedLauncherExecutableStaged  bool            `json:"managed_launcher_executable_staged"`
	DesktopExecUsesExternalAppHandle bool            `json:"desktop_exec_uses_external_app_handle"`
	ExternalAppDesktopHandleReady    bool            `json:"external_app_desktop_handle_ready"`
	DesktopLaunchPacketWritten       bool            `json:"desktop_launch_packet_written"`
	LauncherRequestType              string          `json:"launcher_request_type"`
	LauncherStatus                   string          `json:"launcher_status"`
	ExternalAppImportRecordConsumed  bool            `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool            `json:"external_app_handle_consumed"`
	ExternalFileBridgeReady          bool            `json:"external_file_bridge_ready"`
	ImportedArtifactDigestVerified   bool            `json:"imported_artifact_digest_verified"`
	RuntimeLaunchExecuted            bool            `json:"runtime_launch_executed"`
	WindowObserved                   bool            `json:"window_observed"`
	XWindowObserved                  bool            `json:"x_window_observed"`
	RuntimeOwned                     bool            `json:"runtime_owned"`
	GoRuntimeBacked                  bool            `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool            `json:"kde_policy_owner"`
	LaunchEnabled                    bool            `json:"launch_enabled"`
	BackendLaunchEnabled             bool            `json:"backend_launch_enabled"`
	HostRootModified                 bool            `json:"host_root_modified"`
	PrivilegedContainerRequired      bool            `json:"privileged_container_required"`
	HostNetworkingRequired           bool            `json:"host_networking_required"`
	DockerSocketMounted              bool            `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool            `json:"broad_host_mount_required"`
	RawImportRecordPathExposed       bool            `json:"raw_import_record_path_exposed"`
	RawStateRootPathExposed          bool            `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed         bool            `json:"raw_executable_path_exposed"`
	RawLauncherPathExposed           bool            `json:"raw_launcher_path_exposed"`
	RawLauncherOutputExposed         bool            `json:"raw_launcher_output_exposed"`
	LauncherResult                   json.RawMessage `json:"launcher_result"`
}

type knownPortableBundleStageLaunchPayload struct {
	SchemaVersion                              string          `json:"schema_version"`
	RequestType                                string          `json:"request_type"`
	Status                                     string          `json:"status"`
	AppID                                      string          `json:"app_id"`
	DisplayName                                string          `json:"display_name"`
	KnownPortableBundleImportRequestType       string          `json:"known_portable_bundle_import_request_type"`
	KnownPortableBundleImportStatus            string          `json:"known_portable_bundle_import_status"`
	KnownPortableBundleImportRecorded          bool            `json:"known_portable_bundle_import_recorded"`
	KnownPortableBundleChecksumVerified        bool            `json:"known_portable_bundle_checksum_verified"`
	KnownPortableBundleArchiveVerified         bool            `json:"known_portable_bundle_archive_verified"`
	KnownPortableBundleExtracted               bool            `json:"known_portable_bundle_extracted"`
	KnownPortableBundleImportRecordRequestType string          `json:"known_portable_bundle_import_record_request_type"`
	KnownPortableBundleImportedArtifactKind    string          `json:"known_portable_bundle_imported_artifact_kind"`
	KnownPortableBundleManifestSHA256Present   bool            `json:"known_portable_bundle_manifest_sha256_present"`
	RecordFirstLaunchPath                      bool            `json:"record_first_launch_path"`
	StageLaunchRequestType                     string          `json:"stage_launch_request_type"`
	StageLaunchStatus                          string          `json:"stage_launch_status"`
	ExistingImportRecordConsumed               bool            `json:"existing_import_record_consumed"`
	ExternalAppImportRecordConsumed            bool            `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed                  bool            `json:"external_app_handle_consumed"`
	ImportedArtifactDigestVerified             bool            `json:"imported_artifact_digest_verified"`
	ArtifactKind                               string          `json:"artifact_kind"`
	ApplicationWorkspaceCopied                 bool            `json:"application_workspace_copied"`
	ApplicationWorkspaceMode                   string          `json:"application_workspace_mode"`
	ExternalFileBridgeReady                    bool            `json:"external_file_bridge_ready"`
	WindowsProcessFileArgumentWindowObserved   bool            `json:"windows_process_file_argument_window_observed"`
	RuntimeLaunchExecuted                      bool            `json:"runtime_launch_executed"`
	WindowObserved                             bool            `json:"window_observed"`
	XWindowObserved                            bool            `json:"x_window_observed"`
	RuntimeOwned                               bool            `json:"runtime_owned"`
	GoRuntimeBacked                            bool            `json:"go_runtime_backed"`
	KDEPolicyOwner                             bool            `json:"kde_policy_owner"`
	HostRootModified                           bool            `json:"host_root_modified"`
	PrivilegedContainerRequired                bool            `json:"privileged_container_required"`
	HostNetworkingRequired                     bool            `json:"host_networking_required"`
	DockerSocketMounted                        bool            `json:"docker_socket_mounted"`
	BroadHostMountRequired                     bool            `json:"broad_host_mount_required"`
	RawHostPathExposed                         bool            `json:"raw_host_path_exposed"`
	RawArchivePathExposed                      bool            `json:"raw_archive_path_exposed"`
	RawBundleRootPathExposed                   bool            `json:"raw_bundle_root_path_exposed"`
	RawImportRecordPathExposed                 bool            `json:"raw_import_record_path_exposed"`
	RawStateRootPathExposed                    bool            `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed                   bool            `json:"raw_executable_path_exposed"`
	RawLauncherPathExposed                     bool            `json:"raw_launcher_path_exposed"`
	RawLauncherOutputExposed                   bool            `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed                      bool            `json:"backend_details_exposed"`
	StageLaunchResult                          json.RawMessage `json:"stage_launch_result"`
}

func unwrapKnownPortableBundleStageLaunchPayload(content []byte) ([]byte, error) {
	var payload knownPortableBundleStageLaunchPayload
	if err := json.Unmarshal(content, &payload); err != nil {
		return nil, fmt.Errorf("parse known portable bundle staged launch payload: %w", err)
	}
	switch {
	case payload.SchemaVersion != "xnix.runtime.known_portable_bundle_stage_launch.v1":
		return nil, fmt.Errorf("known portable bundle staged launch payload has unsupported schema %q", payload.SchemaVersion)
	case payload.RequestType != "windows-known-app-bundle-stage-and-launch":
		return nil, errors.New("known portable bundle staged launch payload has invalid request type")
	case payload.Status != "passed":
		return nil, errors.New("known portable bundle staged launch payload requires passed status")
	case strings.TrimSpace(payload.AppID) == "" || strings.TrimSpace(payload.DisplayName) == "":
		return nil, errors.New("known portable bundle staged launch payload requires app identity")
	case payload.KnownPortableBundleImportRequestType != KnownPortableBundleImportRecordRequestType ||
		payload.KnownPortableBundleImportStatus != "passed" ||
		!payload.KnownPortableBundleImportRecorded ||
		!payload.KnownPortableBundleChecksumVerified ||
		!payload.KnownPortableBundleArchiveVerified ||
		!payload.KnownPortableBundleExtracted:
		return nil, errors.New("known portable bundle staged launch payload requires verified bundle import evidence")
	case payload.KnownPortableBundleImportRecordRequestType != ExternalWinAppBundleImportRecordRequestType ||
		payload.KnownPortableBundleImportedArtifactKind != "portable-directory" ||
		!payload.KnownPortableBundleManifestSHA256Present:
		return nil, errors.New("known portable bundle staged launch payload requires portable-directory import record evidence")
	case payload.StageLaunchRequestType != "external-winapp-import-stage-and-launch" || payload.StageLaunchStatus != "passed":
		return nil, errors.New("known portable bundle staged launch payload requires a passed staged launch result")
	case !payload.RecordFirstLaunchPath || !payload.ExistingImportRecordConsumed || !payload.ExternalAppImportRecordConsumed || !payload.ExternalAppHandleConsumed:
		return nil, errors.New("known portable bundle staged launch payload requires record-first import handle consumption")
	case !payload.ImportedArtifactDigestVerified || payload.ArtifactKind != "portable-directory" || !payload.ApplicationWorkspaceCopied || payload.ApplicationWorkspaceMode != "portable-directory":
		return nil, errors.New("known portable bundle staged launch payload requires copied portable-directory execution")
	case !payload.ExternalFileBridgeReady || !payload.WindowsProcessFileArgumentWindowObserved || !payload.RuntimeLaunchExecuted || !payload.WindowObserved || !payload.XWindowObserved:
		return nil, errors.New("known portable bundle staged launch payload requires observed Runtime GUI execution")
	case !payload.RuntimeOwned || !payload.GoRuntimeBacked || payload.KDEPolicyOwner:
		return nil, errors.New("known portable bundle staged launch payload has invalid ownership flags")
	case payload.HostRootModified || payload.PrivilegedContainerRequired || payload.HostNetworkingRequired || payload.DockerSocketMounted || payload.BroadHostMountRequired:
		return nil, errors.New("known portable bundle staged launch payload opens unsafe host or container gates")
	case payload.RawHostPathExposed || payload.RawArchivePathExposed || payload.RawBundleRootPathExposed || payload.RawImportRecordPathExposed || payload.RawStateRootPathExposed || payload.RawExecutablePathExposed || payload.RawLauncherPathExposed || payload.RawLauncherOutputExposed:
		return nil, errors.New("known portable bundle staged launch payload exposes unsafe raw paths or launcher output")
	case payload.BackendDetailsExposed:
		return nil, errors.New("known portable bundle staged launch payload exposes backend details")
	case len(payload.StageLaunchResult) == 0:
		return nil, errors.New("known portable bundle staged launch payload requires nested staged launch result")
	}
	return payload.StageLaunchResult, nil
}

func unwrapExternalWinAppImportStageLaunchPayload(content []byte) ([]byte, error) {
	var payload externalWinAppImportStageLaunchPayload
	if err := json.Unmarshal(content, &payload); err != nil {
		return nil, fmt.Errorf("parse one-shot external Windows app launch payload: %w", err)
	}
	switch {
	case payload.SchemaVersion != "xnix.runtime.external_winapp_import_stage_launch.v1":
		return nil, fmt.Errorf("one-shot external Windows app launch payload has unsupported schema %q", payload.SchemaVersion)
	case payload.RequestType != "external-winapp-import-stage-and-launch":
		return nil, errors.New("one-shot external Windows app launch payload has invalid request type")
	case payload.Status != "passed" || payload.LauncherStatus != "passed" || payload.LauncherRequestType != ExternalWinAppRunRequestType:
		return nil, errors.New("one-shot external Windows app launch payload requires a passed launcher result")
	case strings.TrimSpace(payload.ApplicationID) == "" || strings.TrimSpace(payload.ExternalAppHandle) == "":
		return nil, errors.New("one-shot external Windows app launch payload requires application and handle evidence")
	case !payload.ImportRecorded || !payload.DesktopActivationStaged || !payload.StagedLauncherInvoked || !payload.StagedLauncherFromActivationRoot:
		return nil, errors.New("one-shot external Windows app launch payload requires import, staging, and staged launcher evidence")
	case !payload.ManagedLauncherExecutableStaged || !payload.DesktopExecUsesExternalAppHandle || !payload.ExternalAppDesktopHandleReady || !payload.DesktopLaunchPacketWritten:
		return nil, errors.New("one-shot external Windows app launch payload requires safe desktop activation evidence")
	case !payload.ExternalAppImportRecordConsumed || !payload.ExternalAppHandleConsumed || !payload.ExternalFileBridgeReady || !payload.ImportedArtifactDigestVerified:
		return nil, errors.New("one-shot external Windows app launch payload requires imported app run evidence")
	case !payload.RuntimeLaunchExecuted || !payload.WindowObserved || !payload.XWindowObserved:
		return nil, errors.New("one-shot external Windows app launch payload requires observed Runtime launch evidence")
	case !payload.RuntimeOwned || !payload.GoRuntimeBacked || payload.KDEPolicyOwner:
		return nil, errors.New("one-shot external Windows app launch payload has invalid ownership flags")
	case payload.LaunchEnabled || payload.BackendLaunchEnabled || payload.HostRootModified:
		return nil, errors.New("one-shot external Windows app launch payload exposes unsafe launch or host mutation")
	case payload.PrivilegedContainerRequired || payload.HostNetworkingRequired || payload.DockerSocketMounted || payload.BroadHostMountRequired:
		return nil, errors.New("one-shot external Windows app launch payload requires unsafe container privileges")
	case payload.RawImportRecordPathExposed || payload.RawStateRootPathExposed || payload.RawExecutablePathExposed || payload.RawLauncherPathExposed || payload.RawLauncherOutputExposed:
		return nil, errors.New("one-shot external Windows app launch payload exposes unsafe paths or launcher output")
	case len(payload.LauncherResult) == 0:
		return nil, errors.New("one-shot external Windows app launch payload requires nested launcher result")
	}
	return payload.LauncherResult, nil
}

func wrapContainerXGUIRuntimePayload(content []byte) (containerXGUIRuntimePayload, guiSmokeReport, error) {
	var payload containerXGUIRuntimePayload
	if err := json.Unmarshal(content, &payload); err != nil {
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, fmt.Errorf("parse container X GUI Runtime payload: %w", err)
	}
	if payload.SchemaVersion != "xnix.runtime.windows_app_container_x_gui_smoke.v1" ||
		payload.RequestType != "windows-app-container-x-gui-smoke" {
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("real Windows app GUI evidence packet requires a container X GUI Runtime payload")
	}
	report := guiSmokeReport{
		Version:                     payload.AppVersion,
		SchemaVersion:               "xnix.runtime.winapp_smoke_report.v1",
		ReportType:                  "winapp-smoke",
		Status:                      payload.Status,
		Backend:                     "container-x-gui",
		ContainerGUIApp:             payload.ApplicationName,
		ContainerSmokeInvoked:       true,
		ContainerXGUISmokeInvoked:   true,
		SmokeInvoked:                true,
		XServerStarted:              payload.XServerStarted,
		StartupWindowObserved:       payload.XWindowObserved,
		ContainerImageAvailable:     payload.ImageAvailable,
		ContainerRecipeBacked:       payload.RecipeBacked,
		ContainerApplicationID:      payload.ApplicationID,
		ContainerDisplayName:        payload.DisplayName,
		ContainerAppVersion:         payload.AppVersion,
		HostRootModified:            payload.HostRootModified,
		PrivilegedContainerRequired: payload.PrivilegedContainerRequired,
		HostNetworkingRequired:      payload.HostNetworkingRequired,
		DockerSocketMounted:         payload.DockerSocketMounted,
		BroadHostMountRequired:      payload.BroadHostMountRequired,
		KDESafeOutputSummary:        payload.WindowEvidenceSummary,
	}
	report.ContainerPayload.SchemaVersion = payload.SchemaVersion
	report.ContainerPayload.RequestType = payload.RequestType
	report.ContainerPayload.Status = payload.Status
	report.ContainerPayload.ApplicationID = payload.ApplicationID
	report.ContainerPayload.DisplayName = payload.DisplayName
	report.ContainerPayload.AppVersion = payload.AppVersion
	report.ContainerPayload.RecipeBacked = payload.RecipeBacked
	report.ContainerPayload.ExternalAppRunRecordConsumed = payload.ExternalAppRunRecordConsumed
	report.ContainerPayload.ExternalAppHandleConsumed = payload.ExternalAppHandleConsumed
	report.ContainerPayload.ExternalAppImportRecordConsumed = payload.ExternalAppImportRecordConsumed
	report.ContainerPayload.ImportedArtifactDigestVerified = payload.ImportedArtifactDigestVerified
	report.ContainerPayload.ImportedArtifactSHA256 = payload.ImportedArtifactSHA256
	report.ContainerPayload.NetworkMode = payload.NetworkMode
	report.ContainerPayload.XServerStarted = payload.XServerStarted
	report.ContainerPayload.WineBootstrapAttempted = payload.WineBootstrapAttempted
	report.ContainerPayload.ImageAvailable = payload.ImageAvailable
	report.ContainerPayload.XWindowObserved = payload.XWindowObserved
	report.ContainerPayload.WindowEvidenceSummary = payload.WindowEvidenceSummary
	report.ContainerPayload.HostRootModified = payload.HostRootModified
	report.ContainerPayload.PrivilegedContainerRequired = payload.PrivilegedContainerRequired
	report.ContainerPayload.HostNetworkingRequired = payload.HostNetworkingRequired
	report.ContainerPayload.DockerSocketMounted = payload.DockerSocketMounted
	report.ContainerPayload.BroadHostMountRequired = payload.BroadHostMountRequired
	report.ContainerPayload.HostMountCount = payload.HostMountCount
	return payload, report, nil
}

func wrapExternalWinAppRunPayload(content []byte) (containerXGUIRuntimePayload, guiSmokeReport, error) {
	var run ExternalWinAppRunResult
	if err := json.Unmarshal(content, &run); err != nil {
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, fmt.Errorf("parse external Windows app run payload: %w", err)
	}
	switch {
	case run.SchemaVersion != ExternalWinAppRunSchemaVersion:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, fmt.Errorf("external Windows app run payload has unsupported schema %q", run.SchemaVersion)
	case run.RequestType != ExternalWinAppRunRequestType:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload has invalid request type")
	case run.RunType != "external-windows-app-container-gui-run":
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload has invalid run type")
	case !run.RuntimeRunRequested || !run.RuntimeRunExecuted || !run.ExecutionStarted || !run.BackendProcessStarted:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires an executed Runtime run")
	case !run.ContainerRuntimeUsed || run.ContainerNetworkMode != "none" || run.ContainerHostMountCount != 0:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires a network-isolated container without host mounts")
	case !run.XWindowObserved || !run.WindowObserved:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires generic and X-specific observed-window evidence")
	case !run.ExternalAppImportRecordConsumed || !run.ImportedArtifactDigestVerified || !validSHA256Hex(run.ImportedArtifactSHA256):
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires digest-verified import record evidence")
	case run.ExternalAppHandleConsumed && strings.TrimSpace(run.ExternalAppHandle) == "":
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload handle evidence requires an opaque handle")
	case run.KDEPolicyOwner || run.DesktopLaunchEnabled || run.ActionExecutionEnabled || run.BackendDetailsExposed:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload exposes unsafe desktop or backend authority")
	case run.RawImportRecordPathExposed || run.RawExternalAppHandlePathExposed || run.RawStateRootPathExposed || run.RawExecutablePathExposed || run.HostRootModified:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload exposes unsafe paths or host mutation")
	case run.PrivilegedContainerRequired || run.HostNetworkingRequired || run.DockerSocketMounted || run.BroadHostMountRequired:
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires unsafe container privileges")
	}
	payload := containerXGUIRuntimePayload{
		SchemaVersion:                   run.RuntimePayload.SchemaVersion,
		RequestType:                     run.RuntimePayload.RequestType,
		Status:                          run.RuntimePayload.Status,
		ApplicationID:                   run.ApplicationID,
		DisplayName:                     run.DisplayName,
		AppVersion:                      run.AppVersion,
		RecipeBacked:                    run.RuntimePayload.RecipeBacked,
		ExecutableName:                  run.ExecutableName,
		LocalExecutableCopied:           run.RuntimePayload.LocalExecutableCopied,
		ExternalAppRunRecordConsumed:    true,
		ExternalAppHandleConsumed:       run.ExternalAppHandleConsumed,
		ExternalAppImportRecordConsumed: run.ExternalAppImportRecordConsumed,
		ImportedArtifactDigestVerified:  run.ImportedArtifactDigestVerified,
		ImportedArtifactSHA256:          strings.TrimSpace(run.ImportedArtifactSHA256),
		ApplicationName:                 run.RuntimePayload.ApplicationName,
		WindowMatch:                     run.RuntimePayload.WindowMatch,
		ContainerImage:                  run.RuntimePayload.ContainerImage,
		ContainerPlatform:               run.RuntimePayload.ContainerPlatform,
		NetworkMode:                     run.RuntimePayload.NetworkMode,
		XServerStarted:                  run.RuntimePayload.XServerStarted,
		WineBootstrapAttempted:          run.RuntimePayload.WineBootstrapAttempted,
		ImageAvailable:                  run.RuntimePayload.ImageAvailable,
		XWindowObserved:                 run.RuntimePayload.XWindowObserved,
		WindowEvidenceSummary:           run.WindowEvidenceSummary,
		HostRootModified:                run.RuntimePayload.HostRootModified,
		PrivilegedContainerRequired:     run.RuntimePayload.PrivilegedContainerRequired,
		HostNetworkingRequired:          run.RuntimePayload.HostNetworkingRequired,
		DockerSocketMounted:             run.RuntimePayload.DockerSocketMounted,
		BroadHostMountRequired:          run.RuntimePayload.BroadHostMountRequired,
		HostMountCount:                  run.RuntimePayload.HostMountCount,
	}
	if strings.TrimSpace(payload.ApplicationName) == "" {
		payload.ApplicationName = "/" + run.ExecutableName
	}
	if strings.TrimSpace(payload.WindowEvidenceSummary) == "" {
		payload.WindowEvidenceSummary = run.RuntimePayload.WindowEvidenceSummary
	}
	if payload.SchemaVersion != "xnix.runtime.windows_app_container_x_gui_smoke.v1" ||
		payload.RequestType != "windows-app-container-x-gui-smoke" ||
		payload.Status != "passed" ||
		!payload.LocalExecutableCopied ||
		!payload.XWindowObserved {
		return containerXGUIRuntimePayload{}, guiSmokeReport{}, errors.New("external Windows app run payload requires a passed copied-executable container GUI runtime payload")
	}
	report := guiSmokeReport{
		Version:                     run.Version,
		SchemaVersion:               "xnix.runtime.winapp_smoke_report.v1",
		ReportType:                  "winapp-smoke",
		Status:                      payload.Status,
		Backend:                     "container-x-gui",
		ContainerGUIApp:             payload.ApplicationName,
		ContainerSmokeInvoked:       true,
		ContainerXGUISmokeInvoked:   true,
		SmokeInvoked:                true,
		XServerStarted:              payload.XServerStarted,
		StartupWindowObserved:       payload.XWindowObserved,
		ContainerImageAvailable:     payload.ImageAvailable,
		ContainerRecipeBacked:       payload.RecipeBacked,
		ContainerApplicationID:      payload.ApplicationID,
		ContainerDisplayName:        payload.DisplayName,
		ContainerAppVersion:         payload.AppVersion,
		HostRootModified:            payload.HostRootModified,
		PrivilegedContainerRequired: payload.PrivilegedContainerRequired,
		HostNetworkingRequired:      payload.HostNetworkingRequired,
		DockerSocketMounted:         payload.DockerSocketMounted,
		BroadHostMountRequired:      payload.BroadHostMountRequired,
		KDESafeOutputSummary:        payload.WindowEvidenceSummary,
	}
	report.ContainerPayload.SchemaVersion = payload.SchemaVersion
	report.ContainerPayload.RequestType = payload.RequestType
	report.ContainerPayload.Status = payload.Status
	report.ContainerPayload.ApplicationID = payload.ApplicationID
	report.ContainerPayload.DisplayName = payload.DisplayName
	report.ContainerPayload.AppVersion = payload.AppVersion
	report.ContainerPayload.RecipeBacked = payload.RecipeBacked
	report.ContainerPayload.ExternalAppRunRecordConsumed = payload.ExternalAppRunRecordConsumed
	report.ContainerPayload.ExternalAppHandleConsumed = payload.ExternalAppHandleConsumed
	report.ContainerPayload.ExternalAppImportRecordConsumed = payload.ExternalAppImportRecordConsumed
	report.ContainerPayload.ImportedArtifactDigestVerified = payload.ImportedArtifactDigestVerified
	report.ContainerPayload.ImportedArtifactSHA256 = payload.ImportedArtifactSHA256
	report.ContainerPayload.NetworkMode = payload.NetworkMode
	report.ContainerPayload.XServerStarted = payload.XServerStarted
	report.ContainerPayload.WineBootstrapAttempted = payload.WineBootstrapAttempted
	report.ContainerPayload.ImageAvailable = payload.ImageAvailable
	report.ContainerPayload.XWindowObserved = payload.XWindowObserved
	report.ContainerPayload.WindowEvidenceSummary = payload.WindowEvidenceSummary
	report.ContainerPayload.HostRootModified = payload.HostRootModified
	report.ContainerPayload.PrivilegedContainerRequired = payload.PrivilegedContainerRequired
	report.ContainerPayload.HostNetworkingRequired = payload.HostNetworkingRequired
	report.ContainerPayload.DockerSocketMounted = payload.DockerSocketMounted
	report.ContainerPayload.BroadHostMountRequired = payload.BroadHostMountRequired
	report.ContainerPayload.HostMountCount = payload.HostMountCount
	return payload, report, nil
}

func enhanceKnownAppEvidenceWithContainerRuntimePayload(evidence KnownAppSmokeEvidenceSummary, payload containerXGUIRuntimePayload) KnownAppSmokeEvidenceSummary {
	launchReceiptID := strings.TrimSpace(payload.LaunchAuthorizationReceiptID)
	sessionID := strings.TrimSpace(payload.ControlledExecutionSessionID)
	sessionRelativePath := strings.TrimSpace(payload.ControlledSessionRelativePath)
	postReviewDispatchState := strings.TrimSpace(payload.SessionGatedControlledDispatchState)
	reviewReceiptID := strings.TrimSpace(payload.SessionGatedReviewReceiptID)
	if payload.SessionGatedControlledDispatchConsumed &&
		payload.ControlledExecutionSessionConsumed &&
		payload.ControlledSessionDigestVerified &&
		payload.ControlledSessionWindowObserved &&
		!payload.ControlledSessionHostRootModified &&
		!payload.RawCommandExposed &&
		!payload.BackendDetailsExposed &&
		singleLine(launchReceiptID) &&
		singleLine(sessionID) &&
		singleLine(sessionRelativePath) &&
		singleLine(postReviewDispatchState) &&
		singleLine(reviewReceiptID) {
		evidence.StagedLauncherVerified = true
		evidence.LaunchAuthorizationReceiptRequired = true
		evidence.LaunchAuthorizationReceiptState = "recorded"
		evidence.LaunchAuthorizationReceiptID = launchReceiptID
		evidence.LaunchGateState = "controlled-dispatch-ready"
		evidence.LaunchGateConsumed = true
		evidence.LaunchGateReceiptAccepted = true
		evidence.LaunchGateGuestBoundaryAccepted = true
		evidence.ControlledDispatchReady = true
		evidence.ControlledExecutionSessionID = sessionID
		evidence.LauncherSessionGateConsumed = true
		evidence.LauncherSessionDigestVerified = true
		evidence.LauncherSessionRelativePath = sessionRelativePath
		evidence.LauncherSessionRuntimeOwnerConsumable = payload.RuntimeOwnerConsumableSession
		evidence.LauncherSessionKDEReadModelConsumable = payload.KDEReadModelConsumableSession
		evidence.PostReviewDispatchConsumed = true
		evidence.PostReviewDispatchState = postReviewDispatchState
		evidence.SessionGatedReviewReceiptID = reviewReceiptID
	}
	return evidence
}

func KnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket(payload []byte) (KnownAppSmokeEvidenceSummary, error) {
	var packet RealWinAppGUIEvidencePacket
	if err := json.Unmarshal(payload, &packet); err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("parse real Windows app GUI evidence packet: %w", err)
	}
	switch {
	case packet.SchemaVersion != RealWinAppGUIEvidencePacketSchemaVersion:
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("real Windows app GUI evidence packet has unsupported schema %q", packet.SchemaVersion)
	case packet.RequestType != RealWinAppGUIEvidencePacketRequestType:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet has invalid request type")
	case packet.PacketType != "real-windows-app-gui-evidence":
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet has invalid packet type")
	case packet.ReportStatus != "passed" || !packet.ReportConsumed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet has not consumed a passed report")
	case packet.ReportPathExposed || packet.BackendDetailsExposed || packet.RawOutputExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet exposes unsafe details")
	case packet.DesktopLaunchEnabled || packet.BackendLaunchEnabled || packet.ActionExecutionEnabled || packet.HostRootModified:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet enables unsafe execution")
	case packet.PrivilegedContainerRequired || packet.HostNetworkingRequired || packet.DockerSocketMounted || packet.BroadHostMountRequired:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet requires unsafe host or container privileges")
	case !packet.CompatibilityCenterProjectionReady || !packet.KDECenterProjectionReady:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet is not ready for center consumption")
	case !packet.RuntimeOwned || !packet.GoRuntimeBacked || packet.KDEPolicyOwner:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet must remain Runtime-owned and Go-backed")
	case !packet.XWindowObserved || !packet.WindowObserved || packet.XWindowChildCount <= 0:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet requires generic and X-specific observed-window evidence")
	case packet.KnownAppGUIEvidenceCount != 1 || packet.KnownAppGUIEvidenceVerifiedCount != 1:
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet requires exactly one verified GUI evidence item")
	}
	if packet.ContainerRuntimeUsed {
		switch {
		case packet.ContainerNetworkMode != "none":
			return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet requires a network-isolated container")
		case packet.ContainerHostMountCount != 0:
			return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet must not mount host paths")
		case packet.ExternalAppImportRecordConsumed && !packet.ImportedArtifactDigestVerified:
			return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet import record evidence requires digest verification")
		case packet.ExternalAppImportRecordConsumed && !validSHA256Hex(packet.ImportedArtifactSHA256):
			return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet import record evidence requires an artifact digest")
		case packet.ExternalAppHandleConsumed && (!packet.ExternalAppRunRecordConsumed || !packet.ExternalAppImportRecordConsumed):
			return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet handle evidence requires consumed run and import record evidence")
		}
	}
	evidence := packet.KnownAppSmokeEvidence
	if evidence.AppID != packet.AppID ||
		evidence.DisplayName != packet.DisplayName ||
		evidence.AppVersion != packet.AppVersion ||
		evidence.EvidenceSource != packet.EvidenceSource ||
		evidence.RecipeBacked != packet.RecipeBacked ||
		evidence.RecipeAppID != packet.RecipeAppID ||
		evidence.CompatibilityState != packet.CompatibilityState ||
		evidence.CenterCardState != packet.CenterCardState {
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet summary does not match nested evidence")
	}
	if evidence.XWindowObserved != packet.XWindowObserved ||
		evidence.WindowObserved != packet.WindowObserved {
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet window evidence summary does not match nested evidence")
	}
	if evidence.ExternalAppImportRecordConsumed != packet.ExternalAppImportRecordConsumed ||
		evidence.ExternalAppRunRecordConsumed != packet.ExternalAppRunRecordConsumed ||
		evidence.ExternalAppHandleConsumed != packet.ExternalAppHandleConsumed ||
		evidence.ImportedArtifactDigestVerified != packet.ImportedArtifactDigestVerified ||
		strings.TrimSpace(evidence.ImportedArtifactSHA256) != strings.TrimSpace(packet.ImportedArtifactSHA256) {
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet import record summary does not match nested evidence")
	}
	if !knownAppGUIEvidenceSource(evidence.EvidenceSource) || !evidence.ExecutionEvidenceRecorded || !evidence.RuntimeDispatchVerified {
		return KnownAppSmokeEvidenceSummary{}, errors.New("real Windows app GUI evidence packet requires recorded Runtime dispatch evidence")
	}
	normalized, err := normalizeKnownAppSmokeEvidenceItem(evidence)
	if err != nil {
		return KnownAppSmokeEvidenceSummary{}, err
	}
	return normalized, nil
}

func ExternalAppRecipeFromRealWinAppGUIEvidencePacket(payload []byte) (Recipe, Provenance, KnownAppSmokeEvidenceSummary, error) {
	var packet RealWinAppGUIEvidencePacket
	if err := json.Unmarshal(payload, &packet); err != nil {
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, fmt.Errorf("parse external Windows app GUI evidence packet: %w", err)
	}
	evidence, err := KnownAppSmokeEvidenceFromRealWinAppGUIEvidencePacket(payload)
	if err != nil {
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, err
	}
	switch {
	case packet.RecipeBacked || evidence.RecipeBacked:
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page requires non-recipe real GUI evidence")
	case !packet.ContainerRuntimeUsed:
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page requires container GUI Runtime evidence")
	case !packet.LocalExecutableCopied:
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page requires copied external executable evidence")
	case packet.ExternalAppImportRecordConsumed && !packet.ImportedArtifactDigestVerified:
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page import record evidence requires digest verification")
	case strings.TrimSpace(packet.ExecutableName) == "" || strings.ContainsAny(packet.ExecutableName, `/\`):
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page requires a safe executable name")
	case !strings.HasSuffix(strings.ToLower(packet.ExecutableName), ".exe"):
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, errors.New("external Windows app page requires a Windows executable name")
	}

	recipe := Recipe{
		ID:                  evidence.AppID,
		Name:                evidence.DisplayName,
		Version:             evidence.AppVersion,
		Icon:                "application-x-executable",
		Mode:                "automatic",
		SupportedExtensions: []string{},
	}
	if err := recipe.Validate(); err != nil {
		return Recipe{}, Provenance{}, KnownAppSmokeEvidenceSummary{}, err
	}
	provenance := Provenance{
		Source:          "real-gui-evidence-packet",
		RegistryName:    "runtime-observed-external-app",
		DigestVerified:  packet.ImportedArtifactDigestVerified,
		SignatureStatus: "observed-runtime-evidence",
	}
	if packet.ExternalAppImportRecordConsumed {
		provenance.SignatureStatus = "runtime-import-record-digest-verified"
	}
	return recipe, provenance, evidence, nil
}

func validSHA256Hex(value string) bool {
	value = strings.TrimSpace(value)
	if len(value) != 64 {
		return false
	}
	for _, char := range value {
		if (char < '0' || char > '9') && (char < 'a' || char > 'f') {
			return false
		}
	}
	return true
}
