package appidentity

import (
	"errors"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppControlledExecutionSessionConsumeSchemaVersion = "xnix.runtime.known_app_controlled_execution_session_consume.v1"
	KnownAppControlledExecutionSessionConsumeRequestType   = "known-app-controlled-execution-session-consume-preview"
)

type KnownAppControlledExecutionSessionConsumeRequest struct {
	AppID     string
	StateRoot string
	SessionID string
}

type KnownAppControlledExecutionSessionConsumePreview struct {
	SchemaVersion               string                          `json:"schema_version"`
	RequestType                 string                          `json:"request_type"`
	Source                      string                          `json:"source"`
	RuntimeMethod               string                          `json:"runtime_method"`
	ReadMethod                  string                          `json:"read_method"`
	AppID                       string                          `json:"app_id"`
	DisplayName                 string                          `json:"display_name"`
	AppVersion                  string                          `json:"app_version"`
	ExecutionSessionID          string                          `json:"execution_session_id"`
	RecordConsumed              bool                            `json:"record_consumed"`
	LedgerRecordConsumed        bool                            `json:"ledger_record_consumed"`
	SessionRecordConsumed       bool                            `json:"session_record_consumed"`
	SessionDigestVerified       bool                            `json:"session_digest_verified"`
	TransactionRelativePath     string                          `json:"transaction_relative_path"`
	SessionRelativePath         string                          `json:"session_relative_path"`
	SessionSHA256               string                          `json:"session_sha256"`
	SessionState                string                          `json:"session_state"`
	TaskManagerState            string                          `json:"task_manager_state"`
	KWinState                   string                          `json:"kwin_state"`
	TrayState                   string                          `json:"tray_state"`
	CompatibilityCenterState    string                          `json:"compatibility_center_state"`
	FanOutRequestType           string                          `json:"fan_out_request_type"`
	SurfaceCount                int                             `json:"surface_count"`
	Surfaces                    []ExecutionSessionFanOutSurface `json:"surfaces"`
	TaskManager                 ExecutionSessionFanOutSurface   `json:"task_manager"`
	KWin                        ExecutionSessionFanOutSurface   `json:"kwin"`
	Tray                        ExecutionSessionFanOutSurface   `json:"tray"`
	CompatibilityCenter         ExecutionSessionFanOutSurface   `json:"compatibility_center"`
	RuntimeOwnerConsumable      bool                            `json:"runtime_owner_consumable"`
	KDEReadModelConsumable      bool                            `json:"kde_read_model_consumable"`
	SafeForKDE                  bool                            `json:"safe_for_kde"`
	RuntimeOwned                bool                            `json:"runtime_owned"`
	GoRuntimeBacked             bool                            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                            `json:"kde_policy_owner"`
	StateRootPathExposed        bool                            `json:"state_root_path_exposed"`
	TransactionPathExposed      bool                            `json:"transaction_path_exposed"`
	SessionPathExposed          bool                            `json:"session_path_exposed"`
	LiveStateObserved           bool                            `json:"live_state_observed"`
	SessionRegistered           bool                            `json:"session_registered"`
	WindowObserved              bool                            `json:"window_observed"`
	TaskManagerEntryActive      bool                            `json:"task_manager_entry_active"`
	KWinRuleApplied             bool                            `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled       bool                            `json:"live_tray_bridge_enabled"`
	DispatchStarted             bool                            `json:"dispatch_started"`
	ExecutionStarted            bool                            `json:"execution_started"`
	DirectLaunchEnabled         bool                            `json:"direct_launch_enabled"`
	DesktopLaunchEnabled        bool                            `json:"desktop_launch_enabled"`
	BackendLaunchEnabled        bool                            `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                            `json:"backend_process_started"`
	PermissionGranted           bool                            `json:"permission_granted"`
	HostRootModified            bool                            `json:"host_root_modified"`
	NetworkRequired             bool                            `json:"network_required"`
	PrivilegedContainerRequired bool                            `json:"privileged_container_required"`
	DockerSocketMounted         bool                            `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                            `json:"broad_host_mount_required"`
	RawArtifactPathExposed      bool                            `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed       bool                            `json:"backend_details_exposed"`
	DesktopSafeSummary          string                          `json:"desktop_safe_summary"`
}

func PreviewKnownAppControlledExecutionSessionConsumption(request KnownAppControlledExecutionSessionConsumeRequest) (KnownAppControlledExecutionSessionConsumePreview, error) {
	app, err := winapp.LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	stateRoot := strings.TrimSpace(request.StateRoot)
	if stateRoot == "" {
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires --state-root")
	}
	if info, err := os.Stat(stateRoot); err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	} else if !info.IsDir() {
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires a directory state root")
	}

	sessionID := strings.TrimSpace(request.SessionID)
	if sessionID == "" {
		sessionID = KnownAppControlledExecutionSessionID(app.ID, app.Version)
	}
	expectedSessionID := KnownAppControlledExecutionSessionID(app.ID, app.Version)
	if sessionID != expectedSessionID {
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session id does not match the known application catalog")
	}

	ledger, err := execution.NewLedger(stateRoot)
	if err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	transaction, err := ledger.Load(sessionID)
	if err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	session, err := ledger.LoadSession(sessionID)
	if err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	if transaction.ApplicationID != app.ID || session.ApplicationID != app.ID || session.TransactionRelativePath != transaction.RelativePath {
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session record identity mismatch")
	}

	fanOut, err := (Plan{ApplicationID: app.ID}).ExecutionSessionFanOutEvidence(stateRoot, sessionID)
	if err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	preview := KnownAppControlledExecutionSessionConsumePreview{
		SchemaVersion:               KnownAppControlledExecutionSessionConsumeSchemaVersion,
		RequestType:                 KnownAppControlledExecutionSessionConsumeRequestType,
		Source:                      KnownAppControlledExecutionSessionRecordRequestType + "+execution-session-fanout-evidence",
		RuntimeMethod:               "PreviewKnownAppControlledExecutionSessionConsumption",
		ReadMethod:                  "GetKnownAppControlledExecutionSessionConsumption",
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		ExecutionSessionID:          sessionID,
		RecordConsumed:              true,
		LedgerRecordConsumed:        true,
		SessionRecordConsumed:       true,
		SessionDigestVerified:       true,
		TransactionRelativePath:     transaction.RelativePath,
		SessionRelativePath:         session.RelativePath,
		SessionSHA256:               session.SHA256,
		SessionState:                session.SessionState,
		TaskManagerState:            fanOut.TaskManager.State,
		KWinState:                   fanOut.KWin.State,
		TrayState:                   fanOut.Tray.State,
		CompatibilityCenterState:    fanOut.CompatibilityCenter.State,
		FanOutRequestType:           fanOut.RequestType,
		SurfaceCount:                fanOut.SurfaceCount,
		Surfaces:                    append([]ExecutionSessionFanOutSurface{}, fanOut.Surfaces...),
		TaskManager:                 fanOut.TaskManager,
		KWin:                        fanOut.KWin,
		Tray:                        fanOut.Tray,
		CompatibilityCenter:         fanOut.CompatibilityCenter,
		RuntimeOwnerConsumable:      true,
		KDEReadModelConsumable:      true,
		SafeForKDE:                  fanOut.SafeForKDE,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		TransactionPathExposed:      false,
		SessionPathExposed:          false,
		LiveStateObserved:           false,
		SessionRegistered:           false,
		WindowObserved:              false,
		TaskManagerEntryActive:      false,
		KWinRuleApplied:             false,
		LiveTrayBridgeEnabled:       false,
		DispatchStarted:             false,
		ExecutionStarted:            false,
		DirectLaunchEnabled:         false,
		DesktopLaunchEnabled:        false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		PermissionGranted:           false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawArtifactPathExposed:      false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          app.DisplayName + " controlled execution session record is digest-verified and ready for Runtime owner and KDE read-model consumption.",
	}
	return validateKnownAppControlledExecutionSessionConsumePreview(preview)
}

func validateKnownAppControlledExecutionSessionConsumePreview(preview KnownAppControlledExecutionSessionConsumePreview) (KnownAppControlledExecutionSessionConsumePreview, error) {
	switch {
	case !preview.RecordConsumed || !preview.LedgerRecordConsumed || !preview.SessionRecordConsumed || !preview.SessionDigestVerified:
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires digest-verified ledger and session records")
	case !preview.RuntimeOwnerConsumable || !preview.KDEReadModelConsumable || !preview.SafeForKDE || !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner:
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires Runtime-owned KDE-safe read models")
	case preview.TransactionRelativePath == "" || preview.SessionRelativePath == "" || preview.SessionSHA256 == "":
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires relative record evidence")
	case preview.SurfaceCount != 4 || len(preview.Surfaces) != 4:
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires four KDE read-model consumers")
	case preview.StateRootPathExposed || preview.TransactionPathExposed || preview.SessionPathExposed || preview.LiveStateObserved || preview.SessionRegistered || preview.WindowObserved || preview.TaskManagerEntryActive || preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled:
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption must not expose paths or activate live desktop surfaces")
	case preview.DispatchStarted || preview.ExecutionStarted || preview.DirectLaunchEnabled || preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.BackendProcessStarted || preview.PermissionGranted || preview.HostRootModified || preview.NetworkRequired || preview.PrivilegedContainerRequired || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.RawArtifactPathExposed || preview.BackendDetailsExposed:
		return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption must not start execution or mutate host state")
	}
	for _, value := range []string{preview.ExecutionSessionID, preview.SessionState, preview.TaskManagerState, preview.KWinState, preview.TrayState, preview.CompatibilityCenterState} {
		if !singleLine(value) {
			return KnownAppControlledExecutionSessionConsumePreview{}, errors.New("known app controlled execution session consumption requires single-line state values")
		}
	}
	if err := validateNoBackendTerms(preview, "known app controlled execution session consumption"); err != nil {
		return KnownAppControlledExecutionSessionConsumePreview{}, err
	}
	return preview, nil
}
