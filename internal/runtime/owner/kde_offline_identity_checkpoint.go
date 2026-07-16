package owner

import (
	"encoding/json"
	"fmt"
)

const offlineIdentityApplicationID = "org.xnix.sample.notepad"

type KDEOfflineIdentityCheckpointCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type KDEOfflineIdentityCheckpoint struct {
	Version                        string                              `json:"version"`
	SchemaVersion                  string                              `json:"schema_version"`
	RequestType                    string                              `json:"request_type"`
	CheckpointType                 string                              `json:"checkpoint_type"`
	Source                         string                              `json:"source"`
	ApplicationID                  string                              `json:"application_id"`
	DisplayName                    string                              `json:"display_name"`
	DesktopFile                    string                              `json:"desktop_file"`
	StableIdentityDigest           string                              `json:"stable_identity_digest"`
	RecipeDigestVerified           bool                                `json:"recipe_digest_verified"`
	RecipeSignatureStatus          string                              `json:"recipe_signature_status"`
	ProductionSignatureReady       bool                                `json:"production_signature_ready"`
	SurfaceIDs                     []string                            `json:"surface_ids"`
	SurfaceCount                   int                                 `json:"surface_count"`
	ExpectedSurfaceCount           int                                 `json:"expected_surface_count"`
	CrossSurfaceIdentityConsistent bool                                `json:"cross_surface_identity_consistent"`
	OwnerRouteReady                bool                                `json:"owner_route_ready"`
	OwnerServiceCallReady          bool                                `json:"owner_service_call_ready"`
	FormalReadRouteCount           int                                 `json:"formal_read_route_count"`
	OwnerReadMethodCount           int                                 `json:"owner_read_method_count"`
	OwnerLocalReadMethodCount      int                                 `json:"owner_local_read_method_count"`
	SmokeReadRecordCount           int                                 `json:"smoke_read_record_count"`
	SmokeWriteDenialCount          int                                 `json:"smoke_write_denial_count"`
	MethodParityReady              bool                                `json:"method_parity_ready"`
	RouteBandReady                 bool                                `json:"route_band_ready"`
	OfflineIdentityReady           bool                                `json:"offline_identity_ready"`
	Checks                         []KDEOfflineIdentityCheckpointCheck `json:"checks"`
	RuntimeOwned                   bool                                `json:"runtime_owned"`
	GoRuntimeBacked                bool                                `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                                `json:"kde_policy_owner"`
	ReviewOnly                     bool                                `json:"review_only"`
	DesktopFilesWritten            bool                                `json:"desktop_files_written"`
	MIMEDefaultsWritten            bool                                `json:"mime_defaults_written"`
	KRunnerIndexPersisted          bool                                `json:"krunner_index_persisted"`
	TaskManagerEntryActive         bool                                `json:"task_manager_entry_active"`
	KWinRuleApplied                bool                                `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled          bool                                `json:"live_tray_bridge_enabled"`
	NotificationSent               bool                                `json:"notification_sent"`
	SettingsPersisted              bool                                `json:"settings_persisted"`
	CompatibilityCenterPersisted   bool                                `json:"compatibility_center_persisted"`
	LaunchEnabled                  bool                                `json:"launch_enabled"`
	ExecutionStarted               bool                                `json:"execution_started"`
	BackendProcessStarted          bool                                `json:"backend_process_started"`
	ProductionBusClaimed           bool                                `json:"production_bus_claimed"`
	WriteMethodsEnabled            bool                                `json:"write_methods_enabled"`
	NetworkRequired                bool                                `json:"network_required"`
	HostRootModified               bool                                `json:"host_root_modified"`
	PrivilegedContainerRequired    bool                                `json:"privileged_container_required"`
	RawCommandExposed              bool                                `json:"raw_command_exposed"`
	BackendDetailsExposed          bool                                `json:"backend_details_exposed"`
	DesktopSafeSummary             string                              `json:"desktop_safe_summary"`
}

func NewKDEOfflineIdentityCheckpoint(root string) (KDEOfflineIdentityCheckpoint, error) {
	dispatch, err := DispatchRead(root, "GetKDEOfflineApplicationIdentityPreview", []string{offlineIdentityApplicationID})
	if err != nil {
		return KDEOfflineIdentityCheckpoint{}, err
	}
	var identity struct {
		ApplicationID                  string   `json:"application_id"`
		DisplayName                    string   `json:"display_name"`
		DesktopFile                    string   `json:"desktop_file"`
		StableIdentityDigest           string   `json:"stable_identity_digest"`
		RecipeDigestVerified           bool     `json:"recipe_digest_verified"`
		RecipeSignatureStatus          string   `json:"recipe_signature_status"`
		SurfaceIDs                     []string `json:"surface_ids"`
		SurfaceCount                   int      `json:"surface_count"`
		CrossSurfaceIdentityConsistent bool     `json:"cross_surface_identity_consistent"`
		DesktopFilesWritten            bool     `json:"desktop_files_written"`
		MIMEDefaultsWritten            bool     `json:"mime_defaults_written"`
		KRunnerIndexPersisted          bool     `json:"krunner_index_persisted"`
		TaskManagerEntryActive         bool     `json:"task_manager_entry_active"`
		KWinRuleApplied                bool     `json:"kwin_rule_applied"`
		LiveTrayBridgeEnabled          bool     `json:"live_tray_bridge_enabled"`
		NotificationSent               bool     `json:"notification_sent"`
		SettingsPersisted              bool     `json:"settings_persisted"`
		CompatibilityCenterPersisted   bool     `json:"compatibility_center_persisted"`
		LaunchEnabled                  bool     `json:"launch_enabled"`
		ExecutionStarted               bool     `json:"execution_started"`
		BackendProcessStarted          bool     `json:"backend_process_started"`
		NetworkRequired                bool     `json:"network_required"`
		HostRootModified               bool     `json:"host_root_modified"`
		RawCommandExposed              bool     `json:"raw_command_exposed"`
		BackendDetailsExposed          bool     `json:"backend_details_exposed"`
	}
	if err := json.Unmarshal(dispatch.Payload, &identity); err != nil {
		return KDEOfflineIdentityCheckpoint{}, fmt.Errorf("decode offline KDE identity payload: %w", err)
	}

	service, err := NewService(root, ModeSmokeOwner)
	if err != nil {
		return KDEOfflineIdentityCheckpoint{}, err
	}
	call, err := service.Call("GetKDEOfflineApplicationIdentityPreview", []string{offlineIdentityApplicationID})
	if err != nil {
		return KDEOfflineIdentityCheckpoint{}, err
	}
	routes, err := NewRouteCheckpoint(root)
	if err != nil {
		return KDEOfflineIdentityCheckpoint{}, err
	}

	expectedSurfaces := []string{"desktop-entry", "mime-associations", "krunner", "task-manager", "kwin", "system-tray", "notification-center", "unified-settings", "compatibility-center"}
	surfaceCoverageReady := sameOrderedStrings(identity.SurfaceIDs, expectedSurfaces) && identity.SurfaceCount == len(expectedSurfaces)
	unsafeSideEffect := identity.DesktopFilesWritten || identity.MIMEDefaultsWritten || identity.KRunnerIndexPersisted ||
		identity.TaskManagerEntryActive || identity.KWinRuleApplied || identity.LiveTrayBridgeEnabled || identity.NotificationSent ||
		identity.SettingsPersisted || identity.CompatibilityCenterPersisted || identity.LaunchEnabled || identity.ExecutionStarted ||
		identity.BackendProcessStarted || identity.NetworkRequired || identity.HostRootModified || identity.RawCommandExposed ||
		identity.BackendDetailsExposed
	ownerRouteReady := dispatch.RouteReady && dispatch.ReadOnlyDispatch && dispatch.RouteSource == "go-owner-local-preview" &&
		dispatch.GoCommand == "kde-offline-application-identity-preview"
	ownerServiceReady := call.ReadOnlyDispatch && call.DispatchReady && call.InProcessServiceReady && !call.WriteMethodsEnabled
	checks := []KDEOfflineIdentityCheckpointCheck{
		offlineIdentityCheckpointCheck("recipe-trust", identity.RecipeDigestVerified, "The Runtime-managed registry recipe is digest verified."),
		offlineIdentityCheckpointCheck("surface-coverage", surfaceCoverageReady, "All nine offline KDE identity surfaces are present."),
		offlineIdentityCheckpointCheck("identity-parity", identity.CrossSurfaceIdentityConsistent, "Every surface uses one canonical application identity."),
		offlineIdentityCheckpointCheck("owner-route", ownerRouteReady, "The owner-local read route is ready without caller paths."),
		offlineIdentityCheckpointCheck("owner-service", ownerServiceReady, "The in-process owner service serves the identity read."),
		offlineIdentityCheckpointCheck("side-effects-disabled", !unsafeSideEffect, "Desktop writes, persistence, delivery, launch, execution, and host mutation remain disabled."),
	}
	offlineReady := identity.RecipeDigestVerified && surfaceCoverageReady && identity.CrossSurfaceIdentityConsistent &&
		ownerRouteReady && ownerServiceReady && routes.RouteBandReady && !unsafeSideEffect
	checkpoint := KDEOfflineIdentityCheckpoint{
		Version: routes.Version, SchemaVersion: "xnix.runtime.kde_offline_identity_checkpoint.v1",
		RequestType: "kde-offline-identity-checkpoint", CheckpointType: "offline-kde-application-identity-band-checkpoint",
		Source:        "runtime-owner-read-dispatch+runtime-owner-service+runtime-owner-route-checkpoint",
		ApplicationID: identity.ApplicationID, DisplayName: identity.DisplayName, DesktopFile: identity.DesktopFile,
		StableIdentityDigest: identity.StableIdentityDigest, RecipeDigestVerified: identity.RecipeDigestVerified,
		RecipeSignatureStatus: identity.RecipeSignatureStatus, ProductionSignatureReady: identity.RecipeSignatureStatus == "verified",
		SurfaceIDs: append([]string(nil), identity.SurfaceIDs...), SurfaceCount: identity.SurfaceCount, ExpectedSurfaceCount: len(expectedSurfaces),
		CrossSurfaceIdentityConsistent: identity.CrossSurfaceIdentityConsistent,
		OwnerRouteReady:                ownerRouteReady, OwnerServiceCallReady: ownerServiceReady,
		FormalReadRouteCount: routes.FormalReadRouteCount, OwnerReadMethodCount: routes.OwnerReadMethodCount,
		OwnerLocalReadMethodCount: routes.OwnerLocalReadMethodCount, SmokeReadRecordCount: routes.SmokeReadRecordCount,
		SmokeWriteDenialCount: routes.SmokeWriteDenialCount, MethodParityReady: routes.MethodParityReady,
		RouteBandReady: routes.RouteBandReady, OfflineIdentityReady: offlineReady, Checks: checks,
		RuntimeOwned: true, GoRuntimeBacked: true, ReviewOnly: true,
		DesktopFilesWritten: identity.DesktopFilesWritten, MIMEDefaultsWritten: identity.MIMEDefaultsWritten,
		KRunnerIndexPersisted: identity.KRunnerIndexPersisted, TaskManagerEntryActive: identity.TaskManagerEntryActive,
		KWinRuleApplied: identity.KWinRuleApplied, LiveTrayBridgeEnabled: identity.LiveTrayBridgeEnabled,
		NotificationSent: identity.NotificationSent, SettingsPersisted: identity.SettingsPersisted,
		CompatibilityCenterPersisted: identity.CompatibilityCenterPersisted, LaunchEnabled: identity.LaunchEnabled,
		ExecutionStarted: identity.ExecutionStarted, BackendProcessStarted: identity.BackendProcessStarted,
		ProductionBusClaimed: routes.ProductionBusClaimed, WriteMethodsEnabled: routes.WriteMethodsEnabled,
		NetworkRequired: identity.NetworkRequired, HostRootModified: identity.HostRootModified,
		PrivilegedContainerRequired: routes.PrivilegedContainerRequired, RawCommandExposed: identity.RawCommandExposed,
		BackendDetailsExposed: identity.BackendDetailsExposed,
		DesktopSafeSummary:    "The offline KDE identity band passes trust, nine-surface parity, owner route, owner service, and disabled-side-effect checks without claiming production readiness.",
	}
	if err := validateNoBackendTerms(checkpoint, "offline KDE identity checkpoint"); err != nil {
		return KDEOfflineIdentityCheckpoint{}, err
	}
	return checkpoint, nil
}

func offlineIdentityCheckpointCheck(id string, ready bool, summary string) KDEOfflineIdentityCheckpointCheck {
	status := "blocked"
	if ready {
		status = "pass"
	}
	return KDEOfflineIdentityCheckpointCheck{ID: id, Status: status, Summary: summary}
}

func sameOrderedStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}

func (checkpoint KDEOfflineIdentityCheckpoint) Validate() error {
	if !checkpoint.OfflineIdentityReady {
		return fmt.Errorf("offline KDE identity checkpoint is not ready")
	}
	return nil
}
