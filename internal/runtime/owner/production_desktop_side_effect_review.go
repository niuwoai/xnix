package owner

import (
	"os"
	"path/filepath"
	"strings"
)

type ProductionDesktopSideEffectReviewPreview struct {
	Version                      string                                   `json:"version"`
	SchemaVersion                string                                   `json:"schema_version"`
	RequestType                  string                                   `json:"request_type"`
	ReviewType                   string                                   `json:"review_type"`
	Source                       string                                   `json:"source"`
	ReviewDecision               string                                   `json:"review_decision"`
	SurfaceCount                 int                                      `json:"surface_count"`
	RequiredSurfaceCount         int                                      `json:"required_surface_count"`
	ReviewedSurfaceCount         int                                      `json:"reviewed_surface_count"`
	ActiveSurfaceCount           int                                      `json:"active_surface_count"`
	SideEffectSurfaceCount       int                                      `json:"side_effect_surface_count"`
	ProductionOwnershipReady     bool                                     `json:"production_ownership_ready"`
	Surfaces                     []ProductionDesktopSideEffectSurface     `json:"surfaces"`
	SurfaceIDs                   []string                                 `json:"surface_ids"`
	RequiredBeforeProduction     []string                                 `json:"required_before_production"`
	Checks                       []ProductionDesktopSideEffectReviewCheck `json:"checks"`
	CheckIDs                     []string                                 `json:"check_ids"`
	Counts                       ProductionDesktopSideEffectReviewCounts  `json:"counts"`
	RuntimeOwned                 bool                                     `json:"runtime_owned"`
	GoRuntimeBacked              bool                                     `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                                     `json:"kde_policy_owner"`
	OfficialDesktopOnly          bool                                     `json:"official_desktop_only"`
	PlasmaForkRequired           bool                                     `json:"plasma_fork_required"`
	PlasmaSourceModified         bool                                     `json:"plasma_source_modified"`
	SystemServiceStarted         bool                                     `json:"system_service_started"`
	SessionBusClaimed            bool                                     `json:"session_bus_claimed"`
	ProductionBusClaimed         bool                                     `json:"production_bus_claimed"`
	ProductionOwnerEnabled       bool                                     `json:"production_owner_enabled"`
	ProductionActivationReady    bool                                     `json:"production_activation_ready"`
	WriteMethodsEnabled          bool                                     `json:"write_methods_enabled"`
	RuntimeWritesEnabled         bool                                     `json:"runtime_writes_enabled"`
	DesktopFilesWritten          bool                                     `json:"desktop_files_written"`
	MIMEAppsWritten              bool                                     `json:"mimeapps_written"`
	ShellConfigurationWritten    bool                                     `json:"shell_configuration_written"`
	SettingsPersisted            bool                                     `json:"settings_persisted"`
	KRunnerIndexPersisted        bool                                     `json:"krunner_index_persisted"`
	TaskManagerEntryActive       bool                                     `json:"task_manager_entry_active"`
	KWinRuleApplied              bool                                     `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled        bool                                     `json:"live_tray_bridge_enabled"`
	TrayBridgePersisted          bool                                     `json:"tray_bridge_persisted"`
	NotificationSent             bool                                     `json:"notification_sent"`
	NotificationDeliveryEnabled  bool                                     `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted bool                                     `json:"compatibility_center_persisted"`
	PortalRequestCreated         bool                                     `json:"portal_request_created"`
	RequestObjectsCreated        bool                                     `json:"request_objects_created"`
	AdapterInvocationEnabled     bool                                     `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled         bool                                     `json:"backend_launch_enabled"`
	BackendProcessStarted        bool                                     `json:"backend_process_started"`
	FileContentRead              bool                                     `json:"file_content_read"`
	FilePathsExposed             bool                                     `json:"file_paths_exposed"`
	NetworkRequired              bool                                     `json:"network_required"`
	HostRootModified             bool                                     `json:"host_root_modified"`
	PrivilegedContainerRequired  bool                                     `json:"privileged_container_required"`
	StateRootPathExposed         bool                                     `json:"state_root_path_exposed"`
	RawCommandExposed            bool                                     `json:"raw_command_exposed"`
	RawExecutableExposed         bool                                     `json:"raw_executable_exposed"`
	BackendDetailsExposed        bool                                     `json:"backend_details_exposed"`
	BlockedActions               []string                                 `json:"blocked_actions"`
	NextRequirements             []string                                 `json:"next_requirements"`
	DesktopSafeSummary           string                                   `json:"desktop_safe_summary"`
}

type ProductionDesktopSideEffectSurface struct {
	ID                           string `json:"id"`
	Label                        string `json:"label"`
	KDEComponent                 string `json:"kde_component"`
	SourcePreview                string `json:"source_preview"`
	RequiredBeforeProduction     bool   `json:"required_before_production"`
	EvidencePresent              bool   `json:"evidence_present"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	UserVisible                  bool   `json:"user_visible"`
	ReviewOnly                   bool   `json:"review_only"`
	Active                       bool   `json:"active"`
	SideEffectsEnabled           bool   `json:"side_effects_enabled"`
	DesktopFilesWritten          bool   `json:"desktop_files_written"`
	MIMEAppsWritten              bool   `json:"mimeapps_written"`
	ShellConfigurationWritten    bool   `json:"shell_configuration_written"`
	SettingsPersisted            bool   `json:"settings_persisted"`
	KRunnerIndexPersisted        bool   `json:"krunner_index_persisted"`
	TaskManagerEntryActive       bool   `json:"task_manager_entry_active"`
	KWinRuleApplied              bool   `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled        bool   `json:"live_tray_bridge_enabled"`
	NotificationSent             bool   `json:"notification_sent"`
	NotificationDeliveryEnabled  bool   `json:"notification_delivery_enabled"`
	CompatibilityCenterPersisted bool   `json:"compatibility_center_persisted"`
	PortalRequestCreated         bool   `json:"portal_request_created"`
	RequestObjectsCreated        bool   `json:"request_objects_created"`
	BackendLaunchEnabled         bool   `json:"backend_launch_enabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	BackendDetailsExposed        bool   `json:"backend_details_exposed"`
	ReviewStatus                 string `json:"review_status"`
	NextSafeReadOnlyCheck        string `json:"next_safe_read_only_check"`
}

type ProductionDesktopSideEffectReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionDesktopSideEffectReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionDesktopSideEffectReviewPreview(root string) (ProductionDesktopSideEffectReviewPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionDesktopSideEffectReviewPreview{}, err
	}
	sources := productionDesktopSideEffectReviewSources(root)
	surfaces := productionDesktopSideEffectReviewSurfaces(sources)
	preview := ProductionDesktopSideEffectReviewPreview{
		Version:                  version,
		SchemaVersion:            "xnix.runtime.production_desktop_side_effect_review.v1",
		RequestType:              "production-desktop-side-effect-review-preview",
		ReviewType:               "kde-production-desktop-side-effect-review",
		Source:                   "production-dbus-gate-review-preview+production-rollback-diagnostics-review-preview+production-human-authorization-receipt-consolidation-preview+kde-entrypoints-preview+kde-action-queue-preview+desktop-activation-manifest-preview+kde-shell-integration-preview+kde-application-surface-preview",
		ReviewDecision:           "production-desktop-side-effect-review-blocked",
		SurfaceCount:             len(surfaces),
		RequiredSurfaceCount:     7,
		ReviewedSurfaceCount:     productionDesktopSideEffectReviewedCount(surfaces),
		ActiveSurfaceCount:       productionDesktopSideEffectActiveCount(surfaces),
		SideEffectSurfaceCount:   productionDesktopSideEffectEnabledCount(surfaces),
		ProductionOwnershipReady: false,
		Surfaces:                 surfaces,
		SurfaceIDs:               productionDesktopSideEffectSurfaceIDs(surfaces),
		RequiredBeforeProduction: []string{
			"production-dbus-gate-review-preview",
			"production-dbus-method-review-preview",
			"production-rollback-diagnostics-review-preview",
			"production-human-authorization-receipt-consolidation-preview",
			"kde-entrypoints-preview",
			"kde-action-queue-preview",
			"desktop-activation-manifest-preview",
			"kde-shell-integration-preview",
			"kde-application-surface-preview",
			"explicit operator desktop side-effect plan",
		},
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OfficialDesktopOnly:          true,
		PlasmaForkRequired:           false,
		PlasmaSourceModified:         false,
		SystemServiceStarted:         false,
		SessionBusClaimed:            false,
		ProductionBusClaimed:         false,
		ProductionOwnerEnabled:       false,
		ProductionActivationReady:    false,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ShellConfigurationWritten:    false,
		SettingsPersisted:            false,
		KRunnerIndexPersisted:        false,
		TaskManagerEntryActive:       false,
		KWinRuleApplied:              false,
		LiveTrayBridgeEnabled:        false,
		TrayBridgePersisted:          false,
		NotificationSent:             false,
		NotificationDeliveryEnabled:  false,
		CompatibilityCenterPersisted: false,
		PortalRequestCreated:         false,
		RequestObjectsCreated:        false,
		AdapterInvocationEnabled:     false,
		BackendLaunchEnabled:         false,
		BackendProcessStarted:        false,
		FileContentRead:              false,
		FilePathsExposed:             false,
		NetworkRequired:              false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		StateRootPathExposed:         false,
		RawCommandExposed:            false,
		RawExecutableExposed:         false,
		BackendDetailsExposed:        false,
		BlockedActions: []string{
			"treat desktop side-effect review as production ownership approval",
			"write desktop files, MIME defaults, shell configuration, settings, or search indexes from this review",
			"activate task-manager entries, apply KWin rules, enable tray bridges, or persist Compatibility Center state from this review",
			"send or enable notifications from this review",
			"create Portal requests, Runtime request objects, write dispatch, execution, adapter invocation, or compatibility engine launch",
			"read file contents, expose paths, expose state-root paths, expose raw commands, or expose internal engine details",
			"claim production D-Bus ownership, require network, require privileged containers, or mutate host root",
		},
		NextRequirements: []string{
			"Keep all seven KDE first-release entry points review-only until a separate operator-approved side-effect plan exists.",
			"Require production gate, method review, rollback diagnostics review, consolidated human authorization receipt boundary, and service activation gates before desktop writes.",
			"Keep KDE as shell and presentation only; Runtime remains the owner of compatibility policy.",
			"Keep notifications, live tray, task manager activation, KWin rules, settings persistence, search index writes, Portal requests, and host mutation disabled.",
		},
		DesktopSafeSummary: "The desktop side-effect review inventories the seven KDE first-release entry points before production ownership while writing no desktop files, sending no notifications, activating no live shell component, creating no requests, launching no engine, and mutating no host state.",
	}
	checks := productionDesktopSideEffectReviewChecks(preview, sources)
	preview.Checks = checks
	preview.CheckIDs = productionDesktopSideEffectCheckIDs(checks)
	preview.Counts = countProductionDesktopSideEffectChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ReviewDecision = "production-desktop-side-effect-review-ready-side-effects-disabled"
	}
	if err := validateNoBackendTerms(preview, "production desktop side-effect review preview"); err != nil {
		return ProductionDesktopSideEffectReviewPreview{}, err
	}
	return preview, nil
}

type productionDesktopSideEffectReviewSourceSet struct {
	GateReview                string
	RollbackDiagnosticsReview string
	KDEEntryPoints            string
	KDEActionQueue            string
	DesktopActivationManifest string
	KDEShellIntegration       string
	KDEApplicationSurface     string
}

func productionDesktopSideEffectReviewSources(root string) productionDesktopSideEffectReviewSourceSet {
	return productionDesktopSideEffectReviewSourceSet{
		GateReview:                productionDesktopSideEffectReadSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
		RollbackDiagnosticsReview: productionDesktopSideEffectReadSources(root, []string{"internal/runtime/owner/production_rollback_diagnostics_review.go"}),
		KDEEntryPoints:            productionDesktopSideEffectReadSources(root, []string{"internal/runtime/appidentity/kde_entrypoints.go"}),
		KDEActionQueue:            productionDesktopSideEffectReadSources(root, []string{"internal/runtime/appidentity/kde_action_queue.go"}),
		DesktopActivationManifest: productionDesktopSideEffectReadSources(root, []string{"internal/runtime/appidentity/desktop_activation_manifest.go"}),
		KDEShellIntegration:       productionDesktopSideEffectReadSources(root, []string{"internal/runtime/appidentity/kde_shell_surface.go"}),
		KDEApplicationSurface:     productionDesktopSideEffectReadSources(root, []string{"internal/runtime/appidentity/kde_shell_surface.go"}),
	}
}

func productionDesktopSideEffectReviewSurfaces(sources productionDesktopSideEffectReviewSourceSet) []ProductionDesktopSideEffectSurface {
	return []ProductionDesktopSideEffectSurface{
		productionDesktopSideEffectSurface("launcher", "Start menu", "Plasma application launcher", "kde-entrypoints-preview+desktop-activation-manifest-preview", productionDesktopSideEffectHasAll(sources.KDEEntryPoints+sources.DesktopActivationManifest, []string{"launcher", "Start menu", "DesktopFilesWritten"}), "review launcher desktop-entry side effects"),
		productionDesktopSideEffectSurface("task-manager", "Task manager", "Plasma task manager", "kde-entrypoints-preview+kde-application-surface-preview", productionDesktopSideEffectHasAll(sources.KDEEntryPoints+sources.KDEApplicationSurface, []string{"task-manager", "Task manager", "TaskManagerEntryActive"}), "review task manager identity side effects"),
		productionDesktopSideEffectSurface("file-manager", "File manager", "Dolphin", "kde-entrypoints-preview+desktop-activation-manifest-preview", productionDesktopSideEffectHasAll(sources.KDEEntryPoints+sources.DesktopActivationManifest, []string{"file-manager", "File manager", "MIMEAppsWritten"}), "review Dolphin and file association side effects"),
		productionDesktopSideEffectSurface("system-tray", "System tray", "Plasma system tray", "kde-entrypoints-preview+kde-application-surface-preview", productionDesktopSideEffectHasAll(sources.KDEEntryPoints+sources.KDEApplicationSurface, []string{"system-tray", "System tray", "LiveTrayBridgeEnabled"}), "review tray bridge side effects"),
		productionDesktopSideEffectSurface("notifications", "Notification center", "Plasma notification center", "kde-action-queue-preview+desktop-activation-manifest-preview", productionDesktopSideEffectHasAll(sources.KDEActionQueue+sources.DesktopActivationManifest, []string{"notifications", "Notifications", "NotificationsSent"}), "review notification delivery side effects"),
		productionDesktopSideEffectSurface("compatibility-center", "AI Compatibility Center", "Plasma widget", "kde-action-queue-preview+kde-application-surface-preview", productionDesktopSideEffectHasAll(sources.KDEActionQueue+sources.KDEApplicationSurface, []string{"compatibility-center", "Compatibility Center", "KDEPolicyOwner"}), "review Compatibility Center persistence side effects"),
		productionDesktopSideEffectSurface("unified-settings", "Unified settings", "KDE system settings", "kde-entrypoints-preview+kde-action-queue-preview", productionDesktopSideEffectHasAll(sources.KDEEntryPoints+sources.KDEActionQueue, []string{"settings", "Unified settings", "SettingsPersisted"}), "review settings persistence side effects"),
	}
}

func productionDesktopSideEffectSurface(id string, label string, component string, sourcePreview string, evidencePresent bool, nextCheck string) ProductionDesktopSideEffectSurface {
	status := "blocked"
	if evidencePresent {
		status = "reviewed"
	}
	return ProductionDesktopSideEffectSurface{
		ID:                           id,
		Label:                        label,
		KDEComponent:                 component,
		SourcePreview:                sourcePreview,
		RequiredBeforeProduction:     true,
		EvidencePresent:              evidencePresent,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		UserVisible:                  true,
		ReviewOnly:                   true,
		Active:                       false,
		SideEffectsEnabled:           false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ShellConfigurationWritten:    false,
		SettingsPersisted:            false,
		KRunnerIndexPersisted:        false,
		TaskManagerEntryActive:       false,
		KWinRuleApplied:              false,
		LiveTrayBridgeEnabled:        false,
		NotificationSent:             false,
		NotificationDeliveryEnabled:  false,
		CompatibilityCenterPersisted: false,
		PortalRequestCreated:         false,
		RequestObjectsCreated:        false,
		BackendLaunchEnabled:         false,
		HostRootModified:             false,
		BackendDetailsExposed:        false,
		ReviewStatus:                 status,
		NextSafeReadOnlyCheck:        nextCheck,
	}
}

func productionDesktopSideEffectReviewChecks(preview ProductionDesktopSideEffectReviewPreview, sources productionDesktopSideEffectReviewSourceSet) []ProductionDesktopSideEffectReviewCheck {
	return []ProductionDesktopSideEffectReviewCheck{
		productionDesktopSideEffectCheck("production-gate-consumed", productionDesktopSideEffectPassBlocked(productionDesktopSideEffectHasAll(sources.GateReview, []string{"production-dbus-gate-review-preview", "desktop-side-effect-review"})), "The review consumes the production D-Bus gate packet."),
		productionDesktopSideEffectCheck("rollback-diagnostics-consumed", productionDesktopSideEffectPassBlocked(productionDesktopSideEffectHasAll(sources.RollbackDiagnosticsReview, []string{"production-rollback-diagnostics-review-preview", "support-side-effects-disabled"})), "The review consumes rollback and diagnostics evidence before any desktop side effect."),
		productionDesktopSideEffectCheck("seven-kde-surfaces-reviewed", productionDesktopSideEffectPassBlocked(preview.SurfaceCount == 7 && preview.RequiredSurfaceCount == 7 && preview.ReviewedSurfaceCount == 7 && productionDesktopSideEffectSurfacesReady(preview.Surfaces)), "All seven KDE first-release surfaces are reviewed."),
		productionDesktopSideEffectCheck("runtime-policy-owner", productionDesktopSideEffectPassBlocked(preview.RuntimeOwned && preview.GoRuntimeBacked && !preview.KDEPolicyOwner && productionDesktopSideEffectKDEPolicyDisabled(preview.Surfaces)), "KDE remains presentation-only and Runtime owns compatibility policy."),
		productionDesktopSideEffectCheck("desktop-writes-disabled", productionDesktopSideEffectPassBlocked(!preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.KRunnerIndexPersisted && !preview.CompatibilityCenterPersisted && productionDesktopSideEffectSurfacesKeepWritesDisabled(preview.Surfaces)), "Desktop file, MIME, shell, settings, search, and Compatibility Center persistence writes remain disabled."),
		productionDesktopSideEffectCheck("live-shell-activation-disabled", productionDesktopSideEffectPassBlocked(!preview.TaskManagerEntryActive && !preview.KWinRuleApplied && !preview.LiveTrayBridgeEnabled && !preview.TrayBridgePersisted && productionDesktopSideEffectSurfacesKeepLiveActivationDisabled(preview.Surfaces)), "Task manager activation, KWin rules, and live tray bridges remain disabled."),
		productionDesktopSideEffectCheck("notifications-and-requests-disabled", productionDesktopSideEffectPassBlocked(!preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && productionDesktopSideEffectSurfacesKeepNotificationsRequestsDisabled(preview.Surfaces)), "Notifications, Portal requests, and Runtime request objects remain disabled."),
		productionDesktopSideEffectCheck("host-boundary-closed", productionDesktopSideEffectPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed), "Service, bus, write, launch, data exposure, network, privileged, and host mutation gates remain closed."),
	}
}

func productionDesktopSideEffectCheck(id string, status string, summary string) ProductionDesktopSideEffectReviewCheck {
	return ProductionDesktopSideEffectReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionDesktopSideEffectPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionDesktopSideEffectReviewedCount(surfaces []ProductionDesktopSideEffectSurface) int {
	count := 0
	for _, surface := range surfaces {
		if surface.EvidencePresent && surface.ReviewStatus == "reviewed" {
			count++
		}
	}
	return count
}

func productionDesktopSideEffectActiveCount(surfaces []ProductionDesktopSideEffectSurface) int {
	count := 0
	for _, surface := range surfaces {
		if surface.Active {
			count++
		}
	}
	return count
}

func productionDesktopSideEffectEnabledCount(surfaces []ProductionDesktopSideEffectSurface) int {
	count := 0
	for _, surface := range surfaces {
		if surface.SideEffectsEnabled || surface.DesktopFilesWritten || surface.MIMEAppsWritten || surface.SettingsPersisted || surface.KRunnerIndexPersisted || surface.TaskManagerEntryActive || surface.KWinRuleApplied || surface.LiveTrayBridgeEnabled || surface.NotificationSent || surface.NotificationDeliveryEnabled || surface.CompatibilityCenterPersisted || surface.PortalRequestCreated || surface.RequestObjectsCreated || surface.HostRootModified {
			count++
		}
	}
	return count
}

func productionDesktopSideEffectSurfaceIDs(surfaces []ProductionDesktopSideEffectSurface) []string {
	ids := make([]string, 0, len(surfaces))
	for _, surface := range surfaces {
		ids = append(ids, surface.ID)
	}
	return ids
}

func productionDesktopSideEffectCheckIDs(checks []ProductionDesktopSideEffectReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionDesktopSideEffectChecks(checks []ProductionDesktopSideEffectReviewCheck) ProductionDesktopSideEffectReviewCounts {
	counts := ProductionDesktopSideEffectReviewCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}

func productionDesktopSideEffectSurfacesReady(surfaces []ProductionDesktopSideEffectSurface) bool {
	for _, surface := range surfaces {
		if !surface.RequiredBeforeProduction || !surface.EvidencePresent || surface.ReviewStatus != "reviewed" || !surface.UserVisible || !surface.ReviewOnly {
			return false
		}
		if surface.SideEffectsEnabled || surface.BackendLaunchEnabled || surface.HostRootModified || surface.BackendDetailsExposed {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectKDEPolicyDisabled(surfaces []ProductionDesktopSideEffectSurface) bool {
	for _, surface := range surfaces {
		if surface.KDEPolicyOwner || !surface.RuntimeOwned || !surface.GoRuntimeBacked {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectSurfacesKeepWritesDisabled(surfaces []ProductionDesktopSideEffectSurface) bool {
	for _, surface := range surfaces {
		if surface.DesktopFilesWritten || surface.MIMEAppsWritten || surface.ShellConfigurationWritten || surface.SettingsPersisted || surface.KRunnerIndexPersisted || surface.CompatibilityCenterPersisted {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectSurfacesKeepLiveActivationDisabled(surfaces []ProductionDesktopSideEffectSurface) bool {
	for _, surface := range surfaces {
		if surface.TaskManagerEntryActive || surface.KWinRuleApplied || surface.LiveTrayBridgeEnabled {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectSurfacesKeepNotificationsRequestsDisabled(surfaces []ProductionDesktopSideEffectSurface) bool {
	for _, surface := range surfaces {
		if surface.NotificationSent || surface.NotificationDeliveryEnabled || surface.PortalRequestCreated || surface.RequestObjectsCreated {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

func productionDesktopSideEffectReadSources(root string, paths []string) string {
	var builder strings.Builder
	for _, path := range paths {
		content, err := os.ReadFile(filepath.Join(root, path))
		if err != nil {
			continue
		}
		builder.Write(content)
		builder.WriteByte('\n')
	}
	return builder.String()
}
