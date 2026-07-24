package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	GUISmokeEvidencePreviewSchemaVersion = "xnix.runtime.gui_smoke_evidence_preview.v1"
	GUISmokeEvidencePreviewRequestType   = "gui-smoke-evidence-preview"
)

type GUISmokeEvidencePreviewRequest struct {
	ReportPath  string
	AppID       string
	DisplayName string
	AppVersion  string
}

type GUISmokeEvidencePreview struct {
	SchemaVersion                       string                       `json:"schema_version"`
	RequestType                         string                       `json:"request_type"`
	Source                              string                       `json:"source"`
	RuntimeMethod                       string                       `json:"runtime_method"`
	ReadMethod                          string                       `json:"read_method"`
	ReportStatus                        string                       `json:"report_status"`
	ReportConsumed                      bool                         `json:"report_consumed"`
	ReportPathExposed                   bool                         `json:"report_path_exposed"`
	AppID                               string                       `json:"app_id"`
	DisplayName                         string                       `json:"display_name"`
	AppVersion                          string                       `json:"app_version"`
	GUIAppName                          string                       `json:"gui_app_name"`
	LocalGUIExecutableConfigured        bool                         `json:"local_gui_executable_configured"`
	ExecutableCopied                    bool                         `json:"executable_copied"`
	OwnerControlledLaunchRequested      bool                         `json:"owner_controlled_launch_requested"`
	OwnerExternalGUIAppRequested        bool                         `json:"owner_external_gui_app_requested"`
	OwnerExternalGUIAppDelivery         string                       `json:"owner_external_gui_app_delivery,omitempty"`
	OwnerSeedEvidenceProjected          bool                         `json:"owner_seed_evidence_projected"`
	OwnerServiceCallReady               bool                         `json:"owner_service_call_ready"`
	OwnerManagedLauncherInvoked         bool                         `json:"owner_managed_launcher_invoked"`
	OwnerDelegatedManagedArtifactCopied bool                         `json:"owner_delegated_managed_artifact_copied"`
	OwnerDelegatedSmokePassed           bool                         `json:"owner_delegated_smoke_passed"`
	OwnerDelegatedEvidenceSource        string                       `json:"owner_delegated_evidence_source,omitempty"`
	OwnerDelegatedExecutionStarted      bool                         `json:"owner_delegated_execution_started"`
	OwnerDelegatedWindowObserved        bool                         `json:"owner_delegated_controlled_session_window_observed"`
	OwnerControlledLaunchVerified       bool                         `json:"owner_controlled_launch_verified"`
	OwnerManagedCopyVerified            bool                         `json:"owner_managed_copy_verified"`
	WinebootInvoked                     bool                         `json:"wineboot_invoked"`
	XWindowObserved                     bool                         `json:"x_window_observed"`
	XWindowChildCount                   int                          `json:"x_window_child_count"`
	XWinInfoBytes                       int                          `json:"xwininfo_bytes"`
	GuestStderrBytes                    int                          `json:"guest_stderr_bytes"`
	GuestGraphicsDriverErrorObserved    bool                         `json:"guest_graphics_driver_error_observed"`
	CompatibilityCenterProjectionReady  bool                         `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady            bool                         `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidence               KnownAppSmokeEvidenceSummary `json:"known_app_smoke_evidence"`
	RuntimeOwned                        bool                         `json:"runtime_owned"`
	GoRuntimeBacked                     bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner                      bool                         `json:"kde_policy_owner"`
	DesktopLaunchEnabled                bool                         `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                bool                         `json:"backend_launch_enabled"`
	ActionExecutionEnabled              bool                         `json:"action_execution_enabled"`
	BackendDetailsExposed               bool                         `json:"backend_details_exposed"`
	RawOutputExposed                    bool                         `json:"raw_output_exposed"`
	RemotePathExposed                   bool                         `json:"remote_path_exposed"`
	HostRootModified                    bool                         `json:"host_root_modified"`
	PrivilegedContainerRequired         bool                         `json:"privileged_container_required"`
	HostNetworkingRequired              bool                         `json:"host_networking_required"`
	DockerSocketMounted                 bool                         `json:"docker_socket_mounted"`
	BroadHostMountRequired              bool                         `json:"broad_host_mount_required"`
	DesktopSafeSummary                  string                       `json:"desktop_safe_summary"`
}

type guiSmokeReport struct {
	SchemaVersion                        string `json:"schema_version"`
	RequestType                          string `json:"request_type"`
	Status                               string `json:"status"`
	Execute                              bool   `json:"execute"`
	GUIAppName                           string `json:"gui_app_name"`
	LocalGUIExecutableConfigured         bool   `json:"local_gui_executable_configured"`
	ExecutableCopied                     bool   `json:"executable_copied"`
	OwnerControlledLaunchRequested       bool   `json:"owner_controlled_launch_requested"`
	OwnerExternalGUIAppRequested         bool   `json:"owner_external_gui_app_requested"`
	OwnerExternalGUIAppDelivery          string `json:"owner_external_gui_app_delivery"`
	OwnerSeedEvidenceProjected           bool   `json:"owner_seed_evidence_projected"`
	OwnerServiceCallReady                bool   `json:"owner_service_call_ready"`
	OwnerManagedLauncherInvoked          bool   `json:"owner_managed_launcher_invoked"`
	OwnerDelegatedManagedArtifactCopied  bool   `json:"owner_delegated_managed_artifact_copied"`
	OwnerDelegatedSmokePassed            bool   `json:"owner_delegated_smoke_passed"`
	OwnerDelegatedEvidenceSource         string `json:"owner_delegated_evidence_source"`
	OwnerDelegatedExecutionStarted       bool   `json:"owner_delegated_execution_started"`
	OwnerDelegatedWindowObserved         bool   `json:"owner_delegated_controlled_session_window_observed"`
	OwnerDelegatedHostRootModified       bool   `json:"owner_delegated_host_root_modified"`
	OwnerDelegatedDockerSocketMounted    bool   `json:"owner_delegated_docker_socket_mounted"`
	OwnerDelegatedBroadHostMountRequired bool   `json:"owner_delegated_broad_host_mount_required"`
	OwnerDelegatedRawCommandExposed      bool   `json:"owner_delegated_raw_command_exposed"`
	OwnerDelegatedBackendDetailsExposed  bool   `json:"owner_delegated_backend_details_exposed"`
	RuntimeGoOwnedGUISmoke               bool   `json:"runtime_go_owned_gui_smoke"`
	RuntimePayloadSchemaVersion          string `json:"runtime_payload_schema_version"`
	WinebootInvoked                      bool   `json:"wineboot_invoked"`
	XWindowObserved                      bool   `json:"x_window_observed"`
	XWindowChildCount                    int    `json:"x_window_child_count"`
	XWinInfoBytes                        int    `json:"xwininfo_bytes"`
	GuestStderrBytes                     int    `json:"guest_stderr_bytes"`
	GuestGraphicsDriverErrorObserved     bool   `json:"guest_graphics_driver_error_observed"`
	HostRootModified                     bool   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool   `json:"privileged_container_required"`
	HostNetworkingRequired               bool   `json:"host_networking_required"`
	DockerSocketMounted                  bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool   `json:"broad_host_mount_required"`
	KDESafeOutputSummary                 string `json:"kde_safe_output_summary"`
}

func PreviewGUISmokeEvidence(request GUISmokeEvidencePreviewRequest) (GUISmokeEvidencePreview, error) {
	path := strings.TrimSpace(request.ReportPath)
	if path == "" {
		return GUISmokeEvidencePreview{}, errors.New("GUI smoke evidence requires --gui-smoke-report")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return GUISmokeEvidencePreview{}, fmt.Errorf("read GUI smoke report: %w", err)
	}
	return PreviewGUISmokeEvidenceJSON(content, request)
}

func PreviewGUISmokeEvidenceJSON(content []byte, request GUISmokeEvidencePreviewRequest) (GUISmokeEvidencePreview, error) {
	var report guiSmokeReport
	if err := json.Unmarshal(content, &report); err != nil {
		return GUISmokeEvidencePreview{}, fmt.Errorf("parse GUI smoke report: %w", err)
	}
	if err := validateGUISmokeReport(report); err != nil {
		return GUISmokeEvidencePreview{}, err
	}

	appID := strings.TrimSpace(request.AppID)
	if appID == "" {
		appID = "org.xnix.fixture.messagebox"
	}
	displayName := strings.TrimSpace(request.DisplayName)
	if displayName == "" {
		displayName = "Xnix MessageBox GUI fixture"
	}
	appVersion := strings.TrimSpace(request.AppVersion)
	if appVersion == "" {
		appVersion = "local-fixture"
	}
	projectionReady := report.Status == "passed" && report.Execute && report.RuntimeGoOwnedGUISmoke && report.WinebootInvoked && report.XWindowObserved && report.XWindowChildCount > 0
	ownerControlledLaunchVerified := guiSmokeOwnerControlledLaunchVerified(report, projectionReady)
	ownerManagedCopyVerified := ownerControlledLaunchVerified && report.OwnerExternalGUIAppRequested && report.OwnerExternalGUIAppDelivery == "owner-managed-copy" && report.OwnerDelegatedManagedArtifactCopied
	evidence := guiSmokeKnownAppEvidence(appID, displayName, appVersion, report, projectionReady, ownerControlledLaunchVerified, ownerManagedCopyVerified)

	return GUISmokeEvidencePreview{
		SchemaVersion:                       GUISmokeEvidencePreviewSchemaVersion,
		RequestType:                         GUISmokeEvidencePreviewRequestType,
		Source:                              "wine-guest-gui-smoke+runtime-evidence-consumer",
		RuntimeMethod:                       "PreviewGUISmokeEvidence",
		ReadMethod:                          "GetGUISmokeEvidence",
		ReportStatus:                        report.Status,
		ReportConsumed:                      true,
		ReportPathExposed:                   false,
		AppID:                               appID,
		DisplayName:                         displayName,
		AppVersion:                          appVersion,
		GUIAppName:                          report.GUIAppName,
		LocalGUIExecutableConfigured:        report.LocalGUIExecutableConfigured,
		ExecutableCopied:                    report.ExecutableCopied,
		OwnerControlledLaunchRequested:      report.OwnerControlledLaunchRequested,
		OwnerExternalGUIAppRequested:        report.OwnerExternalGUIAppRequested,
		OwnerExternalGUIAppDelivery:         report.OwnerExternalGUIAppDelivery,
		OwnerSeedEvidenceProjected:          report.OwnerSeedEvidenceProjected,
		OwnerServiceCallReady:               report.OwnerServiceCallReady,
		OwnerManagedLauncherInvoked:         report.OwnerManagedLauncherInvoked,
		OwnerDelegatedManagedArtifactCopied: report.OwnerDelegatedManagedArtifactCopied,
		OwnerDelegatedSmokePassed:           report.OwnerDelegatedSmokePassed,
		OwnerDelegatedEvidenceSource:        report.OwnerDelegatedEvidenceSource,
		OwnerDelegatedExecutionStarted:      report.OwnerDelegatedExecutionStarted,
		OwnerDelegatedWindowObserved:        report.OwnerDelegatedWindowObserved,
		OwnerControlledLaunchVerified:       ownerControlledLaunchVerified,
		OwnerManagedCopyVerified:            ownerManagedCopyVerified,
		WinebootInvoked:                     report.WinebootInvoked,
		XWindowObserved:                     report.XWindowObserved,
		XWindowChildCount:                   report.XWindowChildCount,
		XWinInfoBytes:                       report.XWinInfoBytes,
		GuestStderrBytes:                    report.GuestStderrBytes,
		GuestGraphicsDriverErrorObserved:    report.GuestGraphicsDriverErrorObserved,
		CompatibilityCenterProjectionReady:  projectionReady,
		KDECenterProjectionReady:            projectionReady,
		KnownAppSmokeEvidence:               evidence,
		RuntimeOwned:                        true,
		GoRuntimeBacked:                     true,
		KDEPolicyOwner:                      false,
		DesktopLaunchEnabled:                false,
		BackendLaunchEnabled:                false,
		ActionExecutionEnabled:              false,
		BackendDetailsExposed:               false,
		RawOutputExposed:                    false,
		RemotePathExposed:                   false,
		HostRootModified:                    false,
		PrivilegedContainerRequired:         false,
		HostNetworkingRequired:              false,
		DockerSocketMounted:                 false,
		BroadHostMountRequired:              false,
		DesktopSafeSummary:                  fmt.Sprintf("%s GUI smoke evidence observed a real window with %d X child windows.", displayName, report.XWindowChildCount),
	}, nil
}

func KnownAppSmokeEvidenceFromGUISmokeProjection(payload []byte) (KnownAppSmokeEvidenceSummary, error) {
	var projection GUISmokeEvidencePreview
	if err := json.Unmarshal(payload, &projection); err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("parse GUI smoke Runtime evidence projection: %w", err)
	}
	switch {
	case projection.SchemaVersion != GUISmokeEvidencePreviewSchemaVersion:
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("GUI smoke Runtime evidence projection has unsupported schema %q", projection.SchemaVersion)
	case projection.RequestType != GUISmokeEvidencePreviewRequestType:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection has invalid request type")
	case projection.ReportStatus != "passed" || !projection.ReportConsumed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection has not consumed a passed report")
	case !projection.CompatibilityCenterProjectionReady || !projection.KDECenterProjectionReady:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection is not ready for center consumption")
	case projection.ReportPathExposed || projection.RemotePathExposed || projection.BackendDetailsExposed || projection.RawOutputExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection exposes unsafe details")
	case projection.DesktopLaunchEnabled || projection.BackendLaunchEnabled || projection.ActionExecutionEnabled || projection.HostRootModified:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection enables unsafe execution")
	case !projection.RuntimeOwned || !projection.GoRuntimeBacked || projection.KDEPolicyOwner:
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection must remain Runtime-owned and Go-backed")
	}
	evidence := projection.KnownAppSmokeEvidence
	if evidence.EvidenceSource != "wine-guest-gui-smoke" || !evidence.ExecutionEvidenceRecorded || !evidence.RuntimeDispatchVerified {
		return KnownAppSmokeEvidenceSummary{}, errors.New("GUI smoke Runtime evidence projection requires recorded Runtime dispatch evidence")
	}
	normalized, err := normalizeKnownAppSmokeEvidenceItem(evidence)
	if err != nil {
		return KnownAppSmokeEvidenceSummary{}, err
	}
	return normalized, nil
}

func validateGUISmokeReport(report guiSmokeReport) error {
	switch {
	case report.SchemaVersion != "xnix.scripts.wine_guest_gui_smoke.v1":
		return errors.New("GUI smoke evidence requires the Wine guest GUI smoke schema")
	case report.RequestType != "wine-guest-gui-smoke":
		return errors.New("GUI smoke evidence requires a Wine guest GUI smoke report")
	case !report.Execute:
		return errors.New("GUI smoke evidence requires an executed report")
	case report.Status == "":
		return errors.New("GUI smoke evidence requires report status")
	case report.GUIAppName == "" || !singleLine(report.GUIAppName):
		return errors.New("GUI smoke evidence requires a safe GUI app name")
	case !report.RuntimeGoOwnedGUISmoke:
		return errors.New("GUI smoke evidence requires Go Runtime-owned GUI smoke")
	case report.RuntimePayloadSchemaVersion != "xnix.runtime.windows_app_guest_wine_gui_smoke.v1":
		return errors.New("GUI smoke evidence requires the Runtime GUI smoke payload schema")
	case report.HostRootModified || report.PrivilegedContainerRequired || report.HostNetworkingRequired || report.DockerSocketMounted || report.BroadHostMountRequired:
		return errors.New("GUI smoke evidence must keep host and container boundaries closed")
	}
	if report.Status == "passed" {
		switch {
		case !report.WinebootInvoked:
			return errors.New("passed GUI smoke evidence requires wineboot")
		case !report.XWindowObserved:
			return errors.New("passed GUI smoke evidence requires an observed X window")
		case report.XWindowChildCount <= 0:
			return errors.New("passed GUI smoke evidence requires child window evidence")
		case report.XWinInfoBytes <= 0:
			return errors.New("passed GUI smoke evidence requires xwininfo bytes")
		case report.LocalGUIExecutableConfigured && !report.ExecutableCopied:
			return errors.New("local GUI executable smoke evidence requires copied executable evidence")
		}
	}
	if guiSmokeOwnerControlledEvidencePresent(report) {
		switch {
		case !report.OwnerSeedEvidenceProjected:
			return errors.New("owner-controlled GUI smoke evidence requires projected seed evidence")
		case !report.OwnerServiceCallReady:
			return errors.New("owner-controlled GUI smoke evidence requires owner service call readiness")
		case !report.OwnerManagedLauncherInvoked:
			return errors.New("owner-controlled GUI smoke evidence requires managed launcher invocation")
		case !report.OwnerDelegatedSmokePassed:
			return errors.New("owner-controlled GUI smoke evidence requires delegated smoke pass")
		case report.OwnerDelegatedEvidenceSource != "wine-guest-gui-smoke":
			return errors.New("owner-controlled GUI smoke evidence requires Wine GUI delegated evidence")
		case !report.OwnerDelegatedExecutionStarted:
			return errors.New("owner-controlled GUI smoke evidence requires delegated execution start")
		case !report.OwnerDelegatedWindowObserved:
			return errors.New("owner-controlled GUI smoke evidence requires delegated window observation")
		case report.OwnerDelegatedHostRootModified || report.OwnerDelegatedDockerSocketMounted || report.OwnerDelegatedBroadHostMountRequired || report.OwnerDelegatedRawCommandExposed || report.OwnerDelegatedBackendDetailsExposed:
			return errors.New("owner-controlled GUI smoke evidence exposes unsafe delegated state")
		case report.OwnerExternalGUIAppRequested && report.OwnerExternalGUIAppDelivery != "owner-managed-copy":
			return errors.New("owner-controlled external GUI smoke evidence requires owner-managed copy delivery")
		case report.OwnerExternalGUIAppRequested && !report.OwnerDelegatedManagedArtifactCopied:
			return errors.New("owner-controlled external GUI smoke evidence requires delegated managed artifact copy")
		}
	}
	return nil
}

func guiSmokeOwnerControlledLaunchVerified(report guiSmokeReport, projectionReady bool) bool {
	return projectionReady &&
		report.OwnerControlledLaunchRequested &&
		report.OwnerSeedEvidenceProjected &&
		report.OwnerServiceCallReady &&
		report.OwnerManagedLauncherInvoked &&
		report.OwnerDelegatedSmokePassed &&
		report.OwnerDelegatedEvidenceSource == "wine-guest-gui-smoke" &&
		report.OwnerDelegatedExecutionStarted &&
		report.OwnerDelegatedWindowObserved &&
		!report.OwnerDelegatedHostRootModified &&
		!report.OwnerDelegatedDockerSocketMounted &&
		!report.OwnerDelegatedBroadHostMountRequired &&
		!report.OwnerDelegatedRawCommandExposed &&
		!report.OwnerDelegatedBackendDetailsExposed
}

func guiSmokeOwnerControlledEvidencePresent(report guiSmokeReport) bool {
	return report.OwnerSeedEvidenceProjected ||
		report.OwnerServiceCallReady ||
		report.OwnerManagedLauncherInvoked ||
		report.OwnerDelegatedSmokePassed ||
		report.OwnerDelegatedExecutionStarted ||
		report.OwnerDelegatedWindowObserved ||
		report.OwnerDelegatedManagedArtifactCopied
}

func guiSmokeKnownAppEvidence(appID string, displayName string, appVersion string, report guiSmokeReport, projectionReady bool, ownerControlledLaunchVerified bool, ownerManagedCopyVerified bool) KnownAppSmokeEvidenceSummary {
	status := report.Status
	compatibilityState := "gui-smoke-review-required"
	centerCardState := "gui-smoke-evidence-review-required"
	summary := displayName + " GUI smoke evidence is available for review."
	if projectionReady {
		compatibilityState = "real-gui-qemu-wine-verified"
		centerCardState = "validated-real-gui-runtime-run"
		summary = displayName + " passed a real Runtime-owned QEMU/Wine GUI smoke with an observed X window."
	}
	if ownerControlledLaunchVerified {
		compatibilityState = "owner-controlled-gui-qemu-wine-verified"
		centerCardState = "validated-owner-controlled-gui-runtime-run"
		summary = displayName + " passed a Runtime-owner controlled QEMU/Wine GUI launch with an observed X window."
	}
	if ownerManagedCopyVerified {
		summary = displayName + " passed a Runtime-owner controlled QEMU/Wine GUI launch after the managed launcher copied the Windows executable into the guest."
	}
	return KnownAppSmokeEvidenceSummary{
		AppID:                                appID,
		DisplayName:                          displayName,
		AppVersion:                           appVersion,
		EvidenceKind:                         "known-application-gui-smoke",
		EvidenceSource:                       "wine-guest-gui-smoke",
		SmokeStatus:                          status,
		CompatibilityState:                   compatibilityState,
		CenterCardState:                      centerCardState,
		LaunchAuthorizationState:             "review-required",
		PrimaryActionID:                      "review-known-app-gui-evidence",
		PrimaryActionLabel:                   "Review GUI run evidence",
		PrimaryActionKind:                    "review",
		PrimaryActionEnabled:                 true,
		DirectLaunchEnabled:                  false,
		LaunchAuthorizationReceiptRequired:   true,
		LaunchAuthorizationReceiptState:      "missing",
		MarkerObserved:                       false,
		ChecksumVerified:                     false,
		ExecutionEvidenceRecorded:            projectionReady,
		StagedLauncherVerified:               ownerControlledLaunchVerified,
		OwnerControlledRuntimeLaunchVerified: ownerControlledLaunchVerified,
		OwnerManagedCopyVerified:             ownerManagedCopyVerified,
		RuntimeDispatchVerified:              projectionReady,
		LaunchAuthorizationRequired:          true,
		DesktopLaunchEnabled:                 false,
		RuntimeOwned:                         true,
		KDEPolicyOwner:                       false,
		ActionExecutionEnabled:               false,
		BackendLaunchEnabled:                 false,
		HostRootModified:                     false,
		BackendDetailsExposed:                false,
		RawArtifactPathExposed:               false,
		Summary:                              summary,
	}
}
