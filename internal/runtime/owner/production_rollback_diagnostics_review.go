package owner

import (
	"os"
	"path/filepath"
	"strings"
)

type ProductionRollbackDiagnosticsReviewPreview struct {
	Version                     string                                     `json:"version"`
	SchemaVersion               string                                     `json:"schema_version"`
	RequestType                 string                                     `json:"request_type"`
	ReviewType                  string                                     `json:"review_type"`
	Source                      string                                     `json:"source"`
	ReviewDecision              string                                     `json:"review_decision"`
	ReviewItemCount             int                                        `json:"review_item_count"`
	RollbackControlCount        int                                        `json:"rollback_control_count"`
	DiagnosticsControlCount     int                                        `json:"diagnostics_control_count"`
	ReadyControlCount           int                                        `json:"ready_control_count"`
	SideEffectControlCount      int                                        `json:"side_effect_control_count"`
	ProductionOwnershipReady    bool                                       `json:"production_ownership_ready"`
	Items                       []ProductionRollbackDiagnosticsReviewItem  `json:"items"`
	ItemIDs                     []string                                   `json:"item_ids"`
	RequiredBeforeProduction    []string                                   `json:"required_before_production"`
	Checks                      []ProductionRollbackDiagnosticsReviewCheck `json:"checks"`
	CheckIDs                    []string                                   `json:"check_ids"`
	Counts                      ProductionRollbackDiagnosticsReviewCounts  `json:"counts"`
	RuntimeOwned                bool                                       `json:"runtime_owned"`
	GoRuntimeBacked             bool                                       `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                       `json:"kde_policy_owner"`
	SystemServiceStarted        bool                                       `json:"system_service_started"`
	SessionBusClaimed           bool                                       `json:"session_bus_claimed"`
	ProductionBusClaimed        bool                                       `json:"production_bus_claimed"`
	ProductionOwnerEnabled      bool                                       `json:"production_owner_enabled"`
	ProductionActivationReady   bool                                       `json:"production_activation_ready"`
	WriteMethodsEnabled         bool                                       `json:"write_methods_enabled"`
	RuntimeWritesEnabled        bool                                       `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled    bool                                       `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled        bool                                       `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                                       `json:"backend_process_started"`
	SupportBundleExported       bool                                       `json:"support_bundle_exported"`
	SupportCaseCreated          bool                                       `json:"support_case_created"`
	NotificationSent            bool                                       `json:"notification_sent"`
	SnapshotRestoreExecuted     bool                                       `json:"snapshot_restore_executed"`
	StateCleanupExecuted        bool                                       `json:"state_cleanup_executed"`
	FileContentRead             bool                                       `json:"file_content_read"`
	FilePathsExposed            bool                                       `json:"file_paths_exposed"`
	NetworkRequired             bool                                       `json:"network_required"`
	HostRootModified            bool                                       `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                       `json:"privileged_container_required"`
	StateRootPathExposed        bool                                       `json:"state_root_path_exposed"`
	RawCommandExposed           bool                                       `json:"raw_command_exposed"`
	RawExecutableExposed        bool                                       `json:"raw_executable_exposed"`
	BackendDetailsExposed       bool                                       `json:"backend_details_exposed"`
	BlockedActions              []string                                   `json:"blocked_actions"`
	NextRequirements            []string                                   `json:"next_requirements"`
	DesktopSafeSummary          string                                     `json:"desktop_safe_summary"`
}

type ProductionRollbackDiagnosticsReviewItem struct {
	ID                       string `json:"id"`
	Category                 string `json:"category"`
	SourcePreview            string `json:"source_preview"`
	RequiredBeforeProduction bool   `json:"required_before_production"`
	EvidencePresent          bool   `json:"evidence_present"`
	RollbackControl          bool   `json:"rollback_control"`
	DiagnosticsControl       bool   `json:"diagnostics_control"`
	UserVisible              bool   `json:"user_visible"`
	ReviewOnly               bool   `json:"review_only"`
	SideEffectsEnabled       bool   `json:"side_effects_enabled"`
	RestoreExecuted          bool   `json:"restore_executed"`
	CleanupExecuted          bool   `json:"cleanup_executed"`
	SupportBundleExported    bool   `json:"support_bundle_exported"`
	SupportCaseCreated       bool   `json:"support_case_created"`
	NotificationSent         bool   `json:"notification_sent"`
	FileContentRead          bool   `json:"file_content_read"`
	FilePathsExposed         bool   `json:"file_paths_exposed"`
	StateRootPathExposed     bool   `json:"state_root_path_exposed"`
	RawCommandExposed        bool   `json:"raw_command_exposed"`
	RawExecutableExposed     bool   `json:"raw_executable_exposed"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
	HostRootModified         bool   `json:"host_root_modified"`
	ProductionBusClaimed     bool   `json:"production_bus_claimed"`
	WriteMethodsEnabled      bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled     bool   `json:"runtime_writes_enabled"`
	BackendLaunchEnabled     bool   `json:"backend_launch_enabled"`
	ReviewStatus             string `json:"review_status"`
	NextSafeReadOnlyCheck    string `json:"next_safe_read_only_check"`
}

type ProductionRollbackDiagnosticsReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionRollbackDiagnosticsReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionRollbackDiagnosticsReviewPreview(root string) (ProductionRollbackDiagnosticsReviewPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionRollbackDiagnosticsReviewPreview{}, err
	}
	sources := productionRollbackDiagnosticsReviewSources(root)
	items := productionRollbackDiagnosticsReviewItems(sources)
	preview := ProductionRollbackDiagnosticsReviewPreview{
		Version:                  version,
		SchemaVersion:            "xnix.runtime.production_rollback_diagnostics_review.v1",
		RequestType:              "production-rollback-diagnostics-review-preview",
		ReviewType:               "production-dbus-rollback-diagnostics-review",
		Source:                   "production-dbus-gate-review-preview+production-dbus-method-review-preview+runtime-service-activation-preflight-preview+support-bundle-manifest-preview+support-case-timeline-preview+snapshot-restore-candidates-preview+state-root-quota-retention-preview",
		ReviewDecision:           "production-rollback-diagnostics-review-blocked",
		ReviewItemCount:          len(items),
		RollbackControlCount:     productionRollbackDiagnosticsReviewRollbackCount(items),
		DiagnosticsControlCount:  productionRollbackDiagnosticsReviewDiagnosticsCount(items),
		ReadyControlCount:        productionRollbackDiagnosticsReviewReadyCount(items),
		SideEffectControlCount:   productionRollbackDiagnosticsReviewSideEffectCount(items),
		ProductionOwnershipReady: false,
		Items:                    items,
		ItemIDs:                  productionRollbackDiagnosticsReviewItemIDs(items),
		RequiredBeforeProduction: []string{
			"production-dbus-gate-review-preview",
			"production-dbus-method-review-preview",
			"runtime-service-activation-preflight-preview",
			"support-bundle-manifest-preview",
			"support-case-timeline-preview",
			"snapshot-restore-candidates-preview",
			"state-root-quota-retention-preview",
			"explicit operator rollback plan",
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		ProductionOwnerEnabled:      false,
		ProductionActivationReady:   false,
		WriteMethodsEnabled:         false,
		RuntimeWritesEnabled:        false,
		AdapterInvocationEnabled:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
		NotificationSent:            false,
		SnapshotRestoreExecuted:     false,
		StateCleanupExecuted:        false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat rollback diagnostics review as production ownership approval",
			"claim production D-Bus ownership from rollback diagnostics review",
			"start the production Runtime service from rollback diagnostics review",
			"enable write dispatch, request creation, execution, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, send notifications, restore snapshots, or clean state",
			"read file contents, expose host paths, expose state-root paths, expose raw commands, or expose internal engine details",
			"require network, require privileged containers, or mutate host root",
		},
		NextRequirements: []string{
			"Keep rollback diagnostics review read-only until a separate operator-approved production rollback runbook exists.",
			"Require the production D-Bus gate review and method review before considering any production ownership claim.",
			"Require redacted diagnostics and support review surfaces before any service activation path can be treated as production-ready.",
			"Keep restore, cleanup, support export, support case creation, notifications, writes, engine launch, and host mutation disabled.",
		},
		DesktopSafeSummary: "The rollback diagnostics review proves the production gate has rollback and diagnostic evidence before ownership, while keeping service start, bus ownership, writes, support side effects, restore, cleanup, engine launch, unsafe data, and host mutation disabled.",
	}
	checks := productionRollbackDiagnosticsReviewChecks(preview, sources)
	preview.Checks = checks
	preview.CheckIDs = productionRollbackDiagnosticsReviewCheckIDs(checks)
	preview.Counts = countProductionRollbackDiagnosticsReviewChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ReviewDecision = "production-rollback-diagnostics-review-ready-side-effects-disabled"
	}
	if err := validateNoBackendTerms(preview, "production rollback diagnostics review preview"); err != nil {
		return ProductionRollbackDiagnosticsReviewPreview{}, err
	}
	return preview, nil
}

type productionRollbackDiagnosticsReviewSourceSet struct {
	GateReview                 string
	MethodReview               string
	ServiceActivationPreflight string
	SupportBundleManifest      string
	SupportCaseTimeline        string
	SnapshotRestoreCandidates  string
	StateRootQuotaRetention    string
}

func productionRollbackDiagnosticsReviewSources(root string) productionRollbackDiagnosticsReviewSourceSet {
	return productionRollbackDiagnosticsReviewSourceSet{
		GateReview:                 productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
		MethodReview:               productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/owner/production_dbus_method_review.go"}),
		ServiceActivationPreflight: productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/appidentity/runtime_service_activation_preflight.go"}),
		SupportBundleManifest:      productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/appidentity/support_bundle_manifest.go"}),
		SupportCaseTimeline:        productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/appidentity/support_case_timeline.go"}),
		SnapshotRestoreCandidates:  productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/appidentity/snapshot_restore_candidates.go"}),
		StateRootQuotaRetention:    productionRollbackDiagnosticsReviewReadSources(root, []string{"internal/runtime/appidentity/state_root_quota_retention.go"}),
	}
}

func productionRollbackDiagnosticsReviewItems(sources productionRollbackDiagnosticsReviewSourceSet) []ProductionRollbackDiagnosticsReviewItem {
	return []ProductionRollbackDiagnosticsReviewItem{
		productionRollbackDiagnosticsReviewItem("production-bus-rollback-boundary", "rollback", "production-dbus-gate-review-preview", productionRollbackDiagnosticsReviewHasAll(sources.GateReview, []string{"production-dbus-disabled", "ProductionBusClaimed", "HumanAuthorizationRequired"}), true, false, "review the production ownership release boundary"),
		productionRollbackDiagnosticsReviewItem("service-activation-fail-closed-boundary", "rollback", "runtime-service-activation-preflight-preview", productionRollbackDiagnosticsReviewHasAll(sources.ServiceActivationPreflight, []string{"RuntimeServiceActivationProductionDBusGate", "production-dbus-gate-review-consumed-activation-still-blocked", "ProductionActivationReady"}), true, false, "review service activation preflight before ownership"),
		productionRollbackDiagnosticsReviewItem("method-exposure-freeze-boundary", "rollback", "production-dbus-method-review-preview", productionRollbackDiagnosticsReviewHasAll(sources.MethodReview, []string{"production-dbus-method-review-ready-production-exposure-disabled", "ProductionExposureReadyCount", "NewProductionMethodRequestCount"}), true, false, "review route exposure before ownership"),
		productionRollbackDiagnosticsReviewItem("runtime-write-freeze-boundary", "rollback", "production-dbus-method-review-preview+runtime-write-gate-preview", productionRollbackDiagnosticsReviewHasAll(sources.MethodReview, []string{"write-methods-reviewed-disabled", "runtime-write-gate-preview", "WriteMethodsEnabled"}), true, false, "review write gate state before ownership"),
		productionRollbackDiagnosticsReviewItem("support-bundle-redaction-boundary", "diagnostics", "support-bundle-manifest-preview", productionRollbackDiagnosticsReviewHasAll(sources.SupportBundleManifest, []string{"support-bundle-manifest-preview", "redacted-offline-support-bundle-manifest", "ArchiveCreated"}), false, true, "review redacted support bundle manifest"),
		productionRollbackDiagnosticsReviewItem("support-case-timeline-boundary", "diagnostics", "support-case-timeline-preview", productionRollbackDiagnosticsReviewHasAll(sources.SupportCaseTimeline, []string{"support-case-timeline-preview", "redacted-runtime-support-case-timeline", "TicketCreated"}), false, true, "review redacted support case timeline"),
		productionRollbackDiagnosticsReviewItem("snapshot-restore-candidate-boundary", "rollback", "snapshot-restore-candidates-preview", productionRollbackDiagnosticsReviewHasAll(sources.SnapshotRestoreCandidates, []string{"snapshot-restore-candidates-preview", "RestoreExecuted", "SnapshotDeletionEnabled"}), true, true, "review restore candidates without restore"),
		productionRollbackDiagnosticsReviewItem("state-root-retention-dry-run-boundary", "diagnostics", "state-root-quota-retention-preview", productionRollbackDiagnosticsReviewHasAll(sources.StateRootQuotaRetention, []string{"state-root-quota-retention-preview", "FileDeletionEnabled", "ReceiptsRewritten"}), true, true, "review state retention dry run"),
	}
}

func productionRollbackDiagnosticsReviewItem(id string, category string, sourcePreview string, evidencePresent bool, rollbackControl bool, diagnosticsControl bool, nextCheck string) ProductionRollbackDiagnosticsReviewItem {
	status := "blocked"
	if evidencePresent {
		status = "reviewed"
	}
	return ProductionRollbackDiagnosticsReviewItem{
		ID:                       id,
		Category:                 category,
		SourcePreview:            sourcePreview,
		RequiredBeforeProduction: true,
		EvidencePresent:          evidencePresent,
		RollbackControl:          rollbackControl,
		DiagnosticsControl:       diagnosticsControl,
		UserVisible:              true,
		ReviewOnly:               true,
		SideEffectsEnabled:       false,
		RestoreExecuted:          false,
		CleanupExecuted:          false,
		SupportBundleExported:    false,
		SupportCaseCreated:       false,
		NotificationSent:         false,
		FileContentRead:          false,
		FilePathsExposed:         false,
		StateRootPathExposed:     false,
		RawCommandExposed:        false,
		RawExecutableExposed:     false,
		BackendDetailsExposed:    false,
		HostRootModified:         false,
		ProductionBusClaimed:     false,
		WriteMethodsEnabled:      false,
		RuntimeWritesEnabled:     false,
		BackendLaunchEnabled:     false,
		ReviewStatus:             status,
		NextSafeReadOnlyCheck:    nextCheck,
	}
}

func productionRollbackDiagnosticsReviewChecks(preview ProductionRollbackDiagnosticsReviewPreview, sources productionRollbackDiagnosticsReviewSourceSet) []ProductionRollbackDiagnosticsReviewCheck {
	return []ProductionRollbackDiagnosticsReviewCheck{
		productionRollbackDiagnosticsReviewCheck("production-gate-consumed", productionRollbackDiagnosticsReviewPassBlocked(productionRollbackDiagnosticsReviewHasAll(sources.GateReview, []string{"production-dbus-gate-review-preview", "rollback-and-diagnostics-review"})), "The review consumes the production D-Bus gate packet."),
		productionRollbackDiagnosticsReviewCheck("method-review-consumed", productionRollbackDiagnosticsReviewPassBlocked(productionRollbackDiagnosticsReviewHasAll(sources.MethodReview, []string{"production-dbus-method-review-preview", "rollback-and-diagnostics-review"})), "The review consumes the route-by-route production D-Bus method review."),
		productionRollbackDiagnosticsReviewCheck("rollback-controls-reviewed", productionRollbackDiagnosticsReviewPassBlocked(preview.RollbackControlCount == 6 && productionRollbackDiagnosticsReviewItemsReady(preview.Items, true, false)), "Every rollback control has review-only evidence before production ownership."),
		productionRollbackDiagnosticsReviewCheck("diagnostics-controls-reviewed", productionRollbackDiagnosticsReviewPassBlocked(preview.DiagnosticsControlCount == 4 && productionRollbackDiagnosticsReviewItemsReady(preview.Items, false, true)), "Every diagnostics control has redacted review-only evidence before production ownership."),
		productionRollbackDiagnosticsReviewCheck("support-side-effects-disabled", productionRollbackDiagnosticsReviewPassBlocked(!preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.NotificationSent && productionRollbackDiagnosticsReviewItemsKeepSideEffectsDisabled(preview.Items)), "Support bundle export, support case creation, and notifications remain disabled."),
		productionRollbackDiagnosticsReviewCheck("restore-and-cleanup-disabled", productionRollbackDiagnosticsReviewPassBlocked(!preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && productionRollbackDiagnosticsReviewItemsKeepRestoreCleanupDisabled(preview.Items)), "Snapshot restore and state cleanup remain disabled."),
		productionRollbackDiagnosticsReviewCheck("unsafe-data-hidden", productionRollbackDiagnosticsReviewPassBlocked(!preview.FileContentRead && !preview.FilePathsExposed && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionRollbackDiagnosticsReviewItemsHideUnsafeData(preview.Items)), "The review exposes no file contents, host paths, state-root paths, raw commands, executables, or internal engine details."),
		productionRollbackDiagnosticsReviewCheck("host-boundary-closed", productionRollbackDiagnosticsReviewPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired), "Service, bus, write, launch, network, privileged, and host mutation gates remain closed."),
	}
}

func productionRollbackDiagnosticsReviewCheck(id string, status string, summary string) ProductionRollbackDiagnosticsReviewCheck {
	return ProductionRollbackDiagnosticsReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionRollbackDiagnosticsReviewPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionRollbackDiagnosticsReviewRollbackCount(items []ProductionRollbackDiagnosticsReviewItem) int {
	count := 0
	for _, item := range items {
		if item.RollbackControl {
			count++
		}
	}
	return count
}

func productionRollbackDiagnosticsReviewDiagnosticsCount(items []ProductionRollbackDiagnosticsReviewItem) int {
	count := 0
	for _, item := range items {
		if item.DiagnosticsControl {
			count++
		}
	}
	return count
}

func productionRollbackDiagnosticsReviewReadyCount(items []ProductionRollbackDiagnosticsReviewItem) int {
	count := 0
	for _, item := range items {
		if item.EvidencePresent && item.ReviewStatus == "reviewed" {
			count++
		}
	}
	return count
}

func productionRollbackDiagnosticsReviewSideEffectCount(items []ProductionRollbackDiagnosticsReviewItem) int {
	count := 0
	for _, item := range items {
		if item.SideEffectsEnabled || item.RestoreExecuted || item.CleanupExecuted || item.SupportBundleExported || item.SupportCaseCreated || item.NotificationSent || item.HostRootModified {
			count++
		}
	}
	return count
}

func productionRollbackDiagnosticsReviewItemIDs(items []ProductionRollbackDiagnosticsReviewItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
}

func productionRollbackDiagnosticsReviewCheckIDs(checks []ProductionRollbackDiagnosticsReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionRollbackDiagnosticsReviewChecks(checks []ProductionRollbackDiagnosticsReviewCheck) ProductionRollbackDiagnosticsReviewCounts {
	counts := ProductionRollbackDiagnosticsReviewCounts{Total: len(checks)}
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

func productionRollbackDiagnosticsReviewItemsReady(items []ProductionRollbackDiagnosticsReviewItem, rollback bool, diagnostics bool) bool {
	for _, item := range items {
		if rollback && !item.RollbackControl {
			continue
		}
		if diagnostics && !item.DiagnosticsControl {
			continue
		}
		if !item.EvidencePresent || item.ReviewStatus != "reviewed" || !item.ReviewOnly || !item.RequiredBeforeProduction {
			return false
		}
		if item.ProductionBusClaimed || item.WriteMethodsEnabled || item.RuntimeWritesEnabled || item.BackendLaunchEnabled || item.HostRootModified {
			return false
		}
	}
	return true
}

func productionRollbackDiagnosticsReviewItemsKeepSideEffectsDisabled(items []ProductionRollbackDiagnosticsReviewItem) bool {
	for _, item := range items {
		if item.SideEffectsEnabled || item.SupportBundleExported || item.SupportCaseCreated || item.NotificationSent {
			return false
		}
	}
	return true
}

func productionRollbackDiagnosticsReviewItemsKeepRestoreCleanupDisabled(items []ProductionRollbackDiagnosticsReviewItem) bool {
	for _, item := range items {
		if item.RestoreExecuted || item.CleanupExecuted {
			return false
		}
	}
	return true
}

func productionRollbackDiagnosticsReviewItemsHideUnsafeData(items []ProductionRollbackDiagnosticsReviewItem) bool {
	for _, item := range items {
		if item.FileContentRead || item.FilePathsExposed || item.StateRootPathExposed || item.RawCommandExposed || item.RawExecutableExposed || item.BackendDetailsExposed {
			return false
		}
	}
	return true
}

func productionRollbackDiagnosticsReviewHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

func productionRollbackDiagnosticsReviewReadSources(root string, paths []string) string {
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
