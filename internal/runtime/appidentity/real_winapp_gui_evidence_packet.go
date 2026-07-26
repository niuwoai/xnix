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
	XWindowChildCount                  int                          `json:"x_window_child_count"`
	CompatibilityCenterProjectionReady bool                         `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                         `json:"kde_center_projection_ready"`
	ContainerRuntimeUsed               bool                         `json:"container_runtime_used"`
	ContainerNetworkMode               string                       `json:"container_network_mode,omitempty"`
	ContainerHostMountCount            int                          `json:"container_host_mount_count"`
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
	projection, err := PreviewGUISmokeEvidenceJSON(content, GUISmokeEvidencePreviewRequest{
		AppID:       request.AppID,
		DisplayName: request.DisplayName,
		AppVersion:  request.AppVersion,
	})
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
	if containerRuntimeUsed {
		containerNetworkMode = report.ContainerPayload.NetworkMode
		containerHostMountCount = report.ContainerPayload.HostMountCount
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
		XWindowChildCount:                  projection.XWindowChildCount,
		CompatibilityCenterProjectionReady: projection.CompatibilityCenterProjectionReady,
		KDECenterProjectionReady:           projection.KDECenterProjectionReady,
		ContainerRuntimeUsed:               containerRuntimeUsed,
		ContainerNetworkMode:               containerNetworkMode,
		ContainerHostMountCount:            containerHostMountCount,
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
