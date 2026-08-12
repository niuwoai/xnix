package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ExternalWinAppCompatibilityEvidenceBundleSchemaVersion = "xnix.runtime.external_winapp_compatibility_evidence_bundle.v1"
	ExternalWinAppCompatibilityEvidenceBundleRequestType   = "external-winapp-compatibility-evidence-bundle-preview"
)

type ExternalWinAppCompatibilityEvidenceBundleRequest struct {
	OneShotResultPath            string
	DesktopLaunchPacketPath      string
	RuntimeGUIEvidencePacketPath string
	KDEPagePath                  string
}

type ExternalWinAppCompatibilityEvidenceBundle struct {
	Version                          string                                               `json:"version"`
	SchemaVersion                    string                                               `json:"schema_version"`
	RequestType                      string                                               `json:"request_type"`
	BundleType                       string                                               `json:"bundle_type"`
	Source                           string                                               `json:"source"`
	RuntimeMethod                    string                                               `json:"runtime_method"`
	ReadMethod                       string                                               `json:"read_method"`
	Desktop                          string                                               `json:"desktop"`
	ApplicationID                    string                                               `json:"application_id"`
	DisplayName                      string                                               `json:"display_name"`
	AppVersion                       string                                               `json:"app_version,omitempty"`
	EvidenceArtifactCount            int                                                  `json:"evidence_artifact_count"`
	EvidenceArtifacts                []ExternalWinAppCompatibilityEvidenceArtifactSummary `json:"evidence_artifacts"`
	OneShotRuntimeLaunchVerified     bool                                                 `json:"one_shot_runtime_launch_verified"`
	DesktopLaunchPacketVerified      bool                                                 `json:"desktop_launch_packet_verified"`
	RuntimeGUIEvidencePacketVerified bool                                                 `json:"runtime_gui_evidence_packet_verified"`
	KDEExternalAppPageVerified       bool                                                 `json:"kde_external_app_page_verified"`
	RealWindowsAppRunVerified        bool                                                 `json:"real_windows_app_run_verified"`
	ExternalAppImportRecordConsumed  bool                                                 `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool                                                 `json:"external_app_handle_consumed"`
	ImportedArtifactDigestVerified   bool                                                 `json:"imported_artifact_digest_verified"`
	ExternalFileBridgeReady          bool                                                 `json:"external_file_bridge_ready"`
	ExternalFileOpenRequested        bool                                                 `json:"external_file_open_requested"`
	ExternalDesktopArgumentCount     int                                                  `json:"external_desktop_argument_count"`
	WindowObserved                   bool                                                 `json:"window_observed"`
	XWindowObserved                  bool                                                 `json:"x_window_observed"`
	ContainerRuntimeUsed             bool                                                 `json:"container_runtime_used"`
	ContainerNetworkMode             string                                               `json:"container_network_mode"`
	ContainerHostMountCount          int                                                  `json:"container_host_mount_count"`
	KnownAppGUIEvidenceCount         int                                                  `json:"known_app_gui_evidence_count"`
	KnownAppGUIEvidenceVerifiedCount int                                                  `json:"known_app_gui_evidence_verified_count"`
	KDEPageKnownAppGUIEvidenceCount  int                                                  `json:"kde_page_known_app_gui_evidence_count"`
	RuntimeOwned                     bool                                                 `json:"runtime_owned"`
	GoRuntimeBacked                  bool                                                 `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool                                                 `json:"kde_policy_owner"`
	SafeForKDE                       bool                                                 `json:"safe_for_kde"`
	SafeForAIDiagnostics             bool                                                 `json:"safe_for_ai_diagnostics"`
	DesktopLaunchEnabled             bool                                                 `json:"desktop_launch_enabled"`
	BackendLaunchEnabled             bool                                                 `json:"backend_launch_enabled"`
	BackendProcessStarted            bool                                                 `json:"backend_process_started"`
	ActionExecutionEnabled           bool                                                 `json:"action_execution_enabled"`
	BackendDetailsExposed            bool                                                 `json:"backend_details_exposed"`
	RawPathsExposed                  bool                                                 `json:"raw_paths_exposed"`
	RawLauncherOutputExposed         bool                                                 `json:"raw_launcher_output_exposed"`
	HostRootModified                 bool                                                 `json:"host_root_modified"`
	PrivilegedContainerRequired      bool                                                 `json:"privileged_container_required"`
	HostNetworkingRequired           bool                                                 `json:"host_networking_required"`
	DockerSocketMounted              bool                                                 `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool                                                 `json:"broad_host_mount_required"`
	DesktopSafeSummary               string                                               `json:"desktop_safe_summary"`
}

type ExternalWinAppCompatibilityEvidenceArtifactSummary struct {
	ArtifactKind string `json:"artifact_kind"`
	RequestType  string `json:"request_type"`
	Status       string `json:"status"`
	Consumed     bool   `json:"consumed"`
	SafeForKDE   bool   `json:"safe_for_kde"`
}

type externalWinAppCompatibilityOneShotResult struct {
	Version                          string `json:"version"`
	SchemaVersion                    string `json:"schema_version"`
	RequestType                      string `json:"request_type"`
	Status                           string `json:"status"`
	ApplicationID                    string `json:"application_id"`
	DisplayName                      string `json:"display_name"`
	ExternalAppHandle                string `json:"external_app_handle"`
	LauncherRequestType              string `json:"launcher_request_type"`
	LauncherStatus                   string `json:"launcher_status"`
	ImportRecorded                   bool   `json:"import_recorded"`
	DesktopActivationStaged          bool   `json:"desktop_activation_staged"`
	StagedLauncherInvoked            bool   `json:"staged_launcher_invoked"`
	ManagedLauncherExecutableStaged  bool   `json:"managed_launcher_executable_staged"`
	DesktopExecUsesExternalAppHandle bool   `json:"desktop_exec_uses_external_app_handle"`
	ExternalAppDesktopHandleReady    bool   `json:"external_app_desktop_handle_ready"`
	DesktopLaunchPacketWritten       bool   `json:"desktop_launch_packet_written"`
	ExternalAppImportRecordConsumed  bool   `json:"external_app_import_record_consumed"`
	ExternalAppHandleConsumed        bool   `json:"external_app_handle_consumed"`
	ExternalFileBridgeReady          bool   `json:"external_file_bridge_ready"`
	ImportedArtifactDigestVerified   bool   `json:"imported_artifact_digest_verified"`
	RuntimeLaunchExecuted            bool   `json:"runtime_launch_executed"`
	WindowObserved                   bool   `json:"window_observed"`
	XWindowObserved                  bool   `json:"x_window_observed"`
	RuntimeOwned                     bool   `json:"runtime_owned"`
	GoRuntimeBacked                  bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool   `json:"kde_policy_owner"`
	LaunchEnabled                    bool   `json:"launch_enabled"`
	BackendLaunchEnabled             bool   `json:"backend_launch_enabled"`
	HostRootModified                 bool   `json:"host_root_modified"`
	DockerSocketMounted              bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool   `json:"broad_host_mount_required"`
	RawImportRecordPathExposed       bool   `json:"raw_import_record_path_exposed"`
	RawStateRootPathExposed          bool   `json:"raw_state_root_path_exposed"`
	RawExecutablePathExposed         bool   `json:"raw_executable_path_exposed"`
	RawLauncherPathExposed           bool   `json:"raw_launcher_path_exposed"`
	RawLauncherOutputExposed         bool   `json:"raw_launcher_output_exposed"`
	PrivilegedContainerRequired      bool   `json:"privileged_container_required"`
	HostNetworkingRequired           bool   `json:"host_networking_required"`
}

func PreviewExternalWinAppCompatibilityEvidenceBundle(request ExternalWinAppCompatibilityEvidenceBundleRequest) (ExternalWinAppCompatibilityEvidenceBundle, error) {
	oneShot, err := loadExternalWinAppCompatibilityArtifact[externalWinAppCompatibilityOneShotResult](request.OneShotResultPath, "one-shot external Windows app launch result")
	if err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}
	launchPacket, err := loadExternalWinAppCompatibilityArtifact[DesktopExternalWinAppLaunchPacket](request.DesktopLaunchPacketPath, "desktop external Windows app launch packet")
	if err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}
	runtimePacket, err := loadExternalWinAppCompatibilityArtifact[RealWinAppGUIEvidencePacket](request.RuntimeGUIEvidencePacketPath, "real Windows app GUI evidence packet")
	if err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}
	kdePage, err := loadExternalWinAppCompatibilityArtifact[KDECenterPagePreview](request.KDEPagePath, "KDE external app page")
	if err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}

	appID := strings.TrimSpace(oneShot.ApplicationID)
	displayName := strings.TrimSpace(oneShot.DisplayName)
	if appID == "" || displayName == "" {
		return ExternalWinAppCompatibilityEvidenceBundle{}, errors.New("external Windows app compatibility evidence bundle requires one-shot application identity")
	}
	if err := validateExternalWinAppCompatibilityBundleInputs(appID, displayName, oneShot, launchPacket, runtimePacket, kdePage); err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}

	rawPathsExposed := oneShot.RawImportRecordPathExposed || oneShot.RawStateRootPathExposed || oneShot.RawExecutablePathExposed || oneShot.RawLauncherPathExposed ||
		launchPacket.RawImportRecordPathExposed || launchPacket.RawExternalAppHandlePathExposed || launchPacket.RawStateRootPathExposed || launchPacket.RawExecutablePathExposed
	safeForKDE := !rawPathsExposed &&
		!oneShot.RawLauncherOutputExposed &&
		launchPacket.SafeForKDE &&
		!runtimePacket.BackendDetailsExposed &&
		!kdePage.BackendDetailsExposed &&
		!oneShot.HostRootModified &&
		!launchPacket.HostRootModified &&
		!runtimePacket.HostRootModified &&
		!kdePage.HostRootModified &&
		!oneShot.DockerSocketMounted &&
		!launchPacket.DockerSocketMounted &&
		!runtimePacket.DockerSocketMounted &&
		!oneShot.BroadHostMountRequired &&
		!launchPacket.BroadHostMountRequired &&
		!runtimePacket.BroadHostMountRequired

	bundle := ExternalWinAppCompatibilityEvidenceBundle{
		Version:                          firstNonEmpty(oneShot.Version, launchPacket.Version, runtimePacket.Version),
		SchemaVersion:                    ExternalWinAppCompatibilityEvidenceBundleSchemaVersion,
		RequestType:                      ExternalWinAppCompatibilityEvidenceBundleRequestType,
		BundleType:                       "runtime-owned-external-windows-app-compatibility-evidence",
		Source:                           "external-winapp-import-stage-and-launch+desktop-launch-packet+real-winapp-gui-evidence-packet+kde-center-page",
		RuntimeMethod:                    "PreviewExternalWinAppCompatibilityEvidenceBundle",
		ReadMethod:                       "GetExternalWinAppCompatibilityEvidenceBundle",
		Desktop:                          "KDE Plasma",
		ApplicationID:                    appID,
		DisplayName:                      displayName,
		AppVersion:                       firstNonEmpty(launchPacket.AppVersion, runtimePacket.AppVersion),
		EvidenceArtifactCount:            4,
		EvidenceArtifacts:                externalWinAppCompatibilityArtifactSummaries(oneShot, launchPacket, runtimePacket, kdePage),
		OneShotRuntimeLaunchVerified:     true,
		DesktopLaunchPacketVerified:      true,
		RuntimeGUIEvidencePacketVerified: true,
		KDEExternalAppPageVerified:       true,
		RealWindowsAppRunVerified:        true,
		ExternalAppImportRecordConsumed:  oneShot.ExternalAppImportRecordConsumed && launchPacket.ExternalAppImportRecordConsumed && runtimePacket.ExternalAppImportRecordConsumed,
		ExternalAppHandleConsumed:        oneShot.ExternalAppHandleConsumed && launchPacket.ExternalAppHandleConsumed && runtimePacket.ExternalAppHandleConsumed,
		ImportedArtifactDigestVerified:   oneShot.ImportedArtifactDigestVerified && launchPacket.ImportedArtifactDigestVerified && runtimePacket.ImportedArtifactDigestVerified,
		ExternalFileBridgeReady:          oneShot.ExternalFileBridgeReady && launchPacket.ExternalFileBridgeReady,
		ExternalFileOpenRequested:        launchPacket.ExternalFileOpenRequested,
		ExternalDesktopArgumentCount:     launchPacket.ExternalDesktopArgumentCount,
		WindowObserved:                   oneShot.WindowObserved && launchPacket.WindowObserved && runtimePacket.WindowObserved,
		XWindowObserved:                  oneShot.XWindowObserved && launchPacket.XWindowObserved && runtimePacket.XWindowObserved,
		ContainerRuntimeUsed:             launchPacket.ContainerRuntimeUsed && runtimePacket.ContainerRuntimeUsed,
		ContainerNetworkMode:             launchPacket.ContainerNetworkMode,
		ContainerHostMountCount:          launchPacket.ContainerHostMountCount,
		KnownAppGUIEvidenceCount:         runtimePacket.KnownAppGUIEvidenceCount,
		KnownAppGUIEvidenceVerifiedCount: runtimePacket.KnownAppGUIEvidenceVerifiedCount,
		KDEPageKnownAppGUIEvidenceCount:  kdePage.KnownAppGUIEvidenceCount,
		RuntimeOwned:                     oneShot.RuntimeOwned && launchPacket.RuntimeOwned && runtimePacket.RuntimeOwned && kdePage.RuntimeOwned,
		GoRuntimeBacked:                  oneShot.GoRuntimeBacked && launchPacket.GoRuntimeBacked && runtimePacket.GoRuntimeBacked && kdePage.GoRuntimeBacked,
		KDEPolicyOwner:                   oneShot.KDEPolicyOwner || launchPacket.KDEPolicyOwner || runtimePacket.KDEPolicyOwner || kdePage.KDEPolicyOwner,
		SafeForKDE:                       safeForKDE,
		SafeForAIDiagnostics:             safeForKDE,
		DesktopLaunchEnabled:             oneShot.LaunchEnabled || runtimePacket.DesktopLaunchEnabled || kdePage.LaunchEnabled,
		BackendLaunchEnabled:             oneShot.BackendLaunchEnabled || runtimePacket.BackendLaunchEnabled,
		BackendProcessStarted:            launchPacket.BackendProcessStarted || kdePage.BackendProcessStarted,
		ActionExecutionEnabled:           runtimePacket.ActionExecutionEnabled,
		BackendDetailsExposed:            launchPacket.BackendDetailsExposed || runtimePacket.BackendDetailsExposed || kdePage.BackendDetailsExposed,
		RawPathsExposed:                  rawPathsExposed,
		RawLauncherOutputExposed:         oneShot.RawLauncherOutputExposed || runtimePacket.RawOutputExposed,
		HostRootModified:                 oneShot.HostRootModified || launchPacket.HostRootModified || runtimePacket.HostRootModified || kdePage.HostRootModified,
		PrivilegedContainerRequired:      oneShot.PrivilegedContainerRequired || launchPacket.PrivilegedContainerRequired || runtimePacket.PrivilegedContainerRequired,
		HostNetworkingRequired:           oneShot.HostNetworkingRequired || launchPacket.HostNetworkingRequired || runtimePacket.HostNetworkingRequired,
		DockerSocketMounted:              oneShot.DockerSocketMounted || launchPacket.DockerSocketMounted || runtimePacket.DockerSocketMounted,
		BroadHostMountRequired:           oneShot.BroadHostMountRequired || launchPacket.BroadHostMountRequired || runtimePacket.BroadHostMountRequired,
		DesktopSafeSummary:               "External Windows app compatibility evidence is Runtime-owned, verified through a real isolated GUI run, and ready for KDE display without exposing host paths or backend commands.",
	}
	if !bundle.SafeForKDE || bundle.KDEPolicyOwner || !bundle.RuntimeOwned || !bundle.GoRuntimeBacked {
		return ExternalWinAppCompatibilityEvidenceBundle{}, errors.New("external Windows app compatibility evidence bundle failed Runtime/KDE safety invariants")
	}
	if err := validateNoBackendTerms(bundle, "external Windows app compatibility evidence bundle"); err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, err
	}
	return bundle, nil
}

func loadExternalWinAppCompatibilityArtifact[T any](path string, label string) (T, error) {
	var payload T
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return payload, fmt.Errorf("external Windows app compatibility evidence bundle requires --%s", strings.ReplaceAll(label, " ", "-"))
	}
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return payload, fmt.Errorf("read %s: %w", label, err)
	}
	if err := json.Unmarshal(content, &payload); err != nil {
		return payload, fmt.Errorf("parse %s: %w", label, err)
	}
	return payload, nil
}

func validateExternalWinAppCompatibilityBundleInputs(appID string, displayName string, oneShot externalWinAppCompatibilityOneShotResult, launchPacket DesktopExternalWinAppLaunchPacket, runtimePacket RealWinAppGUIEvidencePacket, kdePage KDECenterPagePreview) error {
	switch {
	case oneShot.SchemaVersion != "xnix.runtime.external_winapp_import_stage_launch.v1" || oneShot.RequestType != "external-winapp-import-stage-and-launch":
		return errors.New("external Windows app compatibility evidence bundle requires a one-shot import-stage-and-launch result")
	case oneShot.Status != "passed" || oneShot.LauncherStatus != "passed" || oneShot.LauncherRequestType != ExternalWinAppRunRequestType:
		return errors.New("external Windows app compatibility evidence bundle requires a passed one-shot Runtime launch")
	case oneShot.ExternalAppHandle != appID || !oneShot.ImportRecorded || !oneShot.DesktopActivationStaged || !oneShot.StagedLauncherInvoked || !oneShot.ManagedLauncherExecutableStaged:
		return errors.New("external Windows app compatibility evidence bundle requires one-shot import and staged launcher evidence")
	case !oneShot.DesktopExecUsesExternalAppHandle || !oneShot.ExternalAppDesktopHandleReady || !oneShot.DesktopLaunchPacketWritten:
		return errors.New("external Windows app compatibility evidence bundle requires handle-only desktop activation")
	case !oneShot.RuntimeLaunchExecuted || !oneShot.WindowObserved || !oneShot.XWindowObserved || !oneShot.ExternalFileBridgeReady:
		return errors.New("external Windows app compatibility evidence bundle requires one-shot real GUI and file bridge evidence")
	case launchPacket.SchemaVersion != DesktopExternalWinAppLaunchPacketSchemaVersion || launchPacket.RequestType != DesktopExternalWinAppLaunchPacketRequestType:
		return errors.New("external Windows app compatibility evidence bundle requires a desktop external Windows app launch packet")
	case launchPacket.ApplicationID != appID || launchPacket.DisplayName != displayName || launchPacket.ExternalAppHandle != appID:
		return errors.New("external Windows app compatibility evidence bundle launch packet identity mismatch")
	case !launchPacket.DesktopLaunchPacketReady || !launchPacket.SafeForKDE || !launchPacket.ExternalAppHandleConsumed || !launchPacket.WindowObserved || !launchPacket.XWindowObserved:
		return errors.New("external Windows app compatibility evidence bundle requires a safe observed desktop launch packet")
	case runtimePacket.SchemaVersion != RealWinAppGUIEvidencePacketSchemaVersion || runtimePacket.RequestType != RealWinAppGUIEvidencePacketRequestType:
		return errors.New("external Windows app compatibility evidence bundle requires a real GUI evidence packet")
	case runtimePacket.AppID != appID || runtimePacket.DisplayName != displayName:
		return errors.New("external Windows app compatibility evidence bundle Runtime packet identity mismatch")
	case !runtimePacket.ReportConsumed || runtimePacket.KnownAppGUIEvidenceVerifiedCount != 1 || !runtimePacket.ExternalAppRunRecordConsumed || !runtimePacket.ExternalAppHandleConsumed:
		return errors.New("external Windows app compatibility evidence bundle requires verified Runtime GUI evidence")
	case kdePage.SchemaVersion != "xnix.runtime.kde_center_page.v1" || kdePage.RequestType != "kde-center-page-preview":
		return errors.New("external Windows app compatibility evidence bundle requires a KDE center page preview")
	case kdePage.ApplicationID != appID || kdePage.ApplicationName != displayName || kdePage.KnownAppGUIEvidenceCount != 1:
		return errors.New("external Windows app compatibility evidence bundle KDE page identity or evidence mismatch")
	case oneShot.KDEPolicyOwner || launchPacket.KDEPolicyOwner || runtimePacket.KDEPolicyOwner || kdePage.KDEPolicyOwner:
		return errors.New("external Windows app compatibility evidence bundle must keep KDE out of policy ownership")
	}
	return nil
}

func externalWinAppCompatibilityArtifactSummaries(oneShot externalWinAppCompatibilityOneShotResult, launchPacket DesktopExternalWinAppLaunchPacket, runtimePacket RealWinAppGUIEvidencePacket, kdePage KDECenterPagePreview) []ExternalWinAppCompatibilityEvidenceArtifactSummary {
	return []ExternalWinAppCompatibilityEvidenceArtifactSummary{
		{ArtifactKind: "one-shot-runtime-launch", RequestType: oneShot.RequestType, Status: oneShot.Status, Consumed: true, SafeForKDE: !oneShot.RawLauncherOutputExposed},
		{ArtifactKind: "desktop-launch-packet", RequestType: launchPacket.RequestType, Status: launchPacket.Status, Consumed: true, SafeForKDE: launchPacket.SafeForKDE},
		{ArtifactKind: "runtime-gui-evidence-packet", RequestType: runtimePacket.RequestType, Status: runtimePacket.ReportStatus, Consumed: runtimePacket.ReportConsumed, SafeForKDE: !runtimePacket.BackendDetailsExposed},
		{ArtifactKind: "kde-external-app-page", RequestType: kdePage.RequestType, Status: "rendered", Consumed: true, SafeForKDE: !kdePage.BackendDetailsExposed},
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
