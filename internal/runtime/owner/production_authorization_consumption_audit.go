package owner

import (
	"os"
	"path/filepath"
	"strings"
)

type ProductionAuthorizationConsumptionAuditPreview struct {
	Version                            string                                         `json:"version"`
	SchemaVersion                      string                                         `json:"schema_version"`
	RequestType                        string                                         `json:"request_type"`
	AuditType                          string                                         `json:"audit_type"`
	Source                             string                                         `json:"source"`
	AuditDecision                      string                                         `json:"audit_decision"`
	ReceiptSchema                      string                                         `json:"receipt_schema"`
	OpaqueReceiptID                    string                                         `json:"opaque_receipt_id"`
	ReceiptRequired                    bool                                           `json:"receipt_required"`
	ReceiptPresent                     bool                                           `json:"receipt_present"`
	ReceiptAccepted                    bool                                           `json:"receipt_accepted"`
	ReceiptBoundaryConsolidated        bool                                           `json:"receipt_boundary_consolidated"`
	ConsolidationPreviewConsumed       bool                                           `json:"consolidation_preview_consumed"`
	OwnerManagedOpaqueBoundaryReady    bool                                           `json:"owner_managed_opaque_boundary_ready"`
	CallerStateRootRequired            bool                                           `json:"caller_state_root_required"`
	AuthorizationAccepted              bool                                           `json:"authorization_accepted"`
	ProductionReadiness                bool                                           `json:"production_readiness"`
	ProductionOwnershipReady           bool                                           `json:"production_ownership_ready"`
	ConsumerCount                      int                                            `json:"consumer_count"`
	RequiredConsumerCount              int                                            `json:"required_consumer_count"`
	ConsumedConsumerCount              int                                            `json:"consumed_consumer_count"`
	MissingConsumerCount               int                                            `json:"missing_consumer_count"`
	AuthorizationAcceptedConsumerCount int                                            `json:"authorization_accepted_consumer_count"`
	ProductionReadyConsumerCount       int                                            `json:"production_ready_consumer_count"`
	SideEffectConsumerCount            int                                            `json:"side_effect_consumer_count"`
	Consumers                          []ProductionAuthorizationConsumptionConsumer   `json:"consumers"`
	ConsumerIDs                        []string                                       `json:"consumer_ids"`
	RequiredBeforeProductionOwnership  []string                                       `json:"required_before_production_ownership"`
	Checks                             []ProductionAuthorizationConsumptionAuditCheck `json:"checks"`
	CheckIDs                           []string                                       `json:"check_ids"`
	Counts                             ProductionAuthorizationConsumptionAuditCounts  `json:"counts"`
	RuntimeOwned                       bool                                           `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                           `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                           `json:"kde_policy_owner"`
	SystemServiceStarted               bool                                           `json:"system_service_started"`
	SessionBusClaimed                  bool                                           `json:"session_bus_claimed"`
	ProductionBusClaimed               bool                                           `json:"production_bus_claimed"`
	ProductionOwnerEnabled             bool                                           `json:"production_owner_enabled"`
	ProductionActivationReady          bool                                           `json:"production_activation_ready"`
	WriteMethodsEnabled                bool                                           `json:"write_methods_enabled"`
	RuntimeWritesEnabled               bool                                           `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled               bool                                           `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled          bool                                           `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled         bool                                           `json:"receipt_lookup_writes_enabled"`
	DesktopFilesWritten                bool                                           `json:"desktop_files_written"`
	MIMEAppsWritten                    bool                                           `json:"mimeapps_written"`
	ShellConfigurationWritten          bool                                           `json:"shell_configuration_written"`
	SettingsPersisted                  bool                                           `json:"settings_persisted"`
	NotificationSent                   bool                                           `json:"notification_sent"`
	NotificationDeliveryEnabled        bool                                           `json:"notification_delivery_enabled"`
	PortalRequestCreated               bool                                           `json:"portal_request_created"`
	RequestObjectsCreated              bool                                           `json:"request_objects_created"`
	AdapterInvocationEnabled           bool                                           `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled               bool                                           `json:"backend_launch_enabled"`
	BackendProcessStarted              bool                                           `json:"backend_process_started"`
	SupportBundleExported              bool                                           `json:"support_bundle_exported"`
	SupportCaseCreated                 bool                                           `json:"support_case_created"`
	SnapshotRestoreExecuted            bool                                           `json:"snapshot_restore_executed"`
	StateCleanupExecuted               bool                                           `json:"state_cleanup_executed"`
	FileContentRead                    bool                                           `json:"file_content_read"`
	FilePathsExposed                   bool                                           `json:"file_paths_exposed"`
	NetworkRequired                    bool                                           `json:"network_required"`
	HostRootModified                   bool                                           `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                                           `json:"privileged_container_required"`
	StateRootPathExposed               bool                                           `json:"state_root_path_exposed"`
	RawCommandExposed                  bool                                           `json:"raw_command_exposed"`
	RawExecutableExposed               bool                                           `json:"raw_executable_exposed"`
	BackendDetailsExposed              bool                                           `json:"backend_details_exposed"`
	BlockedActions                     []string                                       `json:"blocked_actions"`
	NextRequirements                   []string                                       `json:"next_requirements"`
	DesktopSafeSummary                 string                                         `json:"desktop_safe_summary"`
}

type ProductionAuthorizationConsumptionConsumer struct {
	ID                           string `json:"id"`
	RequestType                  string `json:"request_type"`
	SourceFile                   string `json:"source_file"`
	ConsumesConsolidatedBoundary bool   `json:"consumes_consolidated_boundary"`
	ReceiptBoundaryReady         bool   `json:"receipt_boundary_ready"`
	ReceiptAccepted              bool   `json:"receipt_accepted"`
	AuthorizationAccepted        bool   `json:"authorization_accepted"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	ReviewOnly                   bool   `json:"review_only"`
	WriteMethodsEnabled          bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled         bool   `json:"runtime_writes_enabled"`
	DesktopSideEffectsEnabled    bool   `json:"desktop_side_effects_enabled"`
	SupportSideEffectsEnabled    bool   `json:"support_side_effects_enabled"`
	BackendLaunchEnabled         bool   `json:"backend_launch_enabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	AuditStatus                  string `json:"audit_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionAuthorizationConsumptionAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionAuthorizationConsumptionAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionAuthorizationConsumptionAuditPreview(root string) (ProductionAuthorizationConsumptionAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionAuthorizationConsumptionAuditPreview{}, err
	}
	sources := productionAuthorizationConsumptionSources(root)
	consumers := productionAuthorizationConsumptionConsumers(sources)
	preview := ProductionAuthorizationConsumptionAuditPreview{
		Version:                            version,
		SchemaVersion:                      "xnix.runtime.production_authorization_consumption_audit.v1",
		RequestType:                        "production-authorization-consumption-audit-preview",
		AuditType:                          "production-gate-consolidated-authorization-consumption-audit",
		Source:                             "production-human-authorization-receipt-consolidation-preview+production-dbus-gate-review-preview+production-dbus-method-review-preview+runtime-service-activation-preflight-preview+runtime-write-gate-preview+production-rollback-diagnostics-review-preview+production-desktop-side-effect-review-preview",
		AuditDecision:                      "production-authorization-consumption-audit-blocked",
		ReceiptSchema:                      "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                    ProductionDBusHumanAuthorizationReceiptID,
		ReceiptRequired:                    true,
		ReceiptPresent:                     false,
		ReceiptAccepted:                    false,
		ReceiptBoundaryConsolidated:        productionAuthorizationConsolidationReady(sources.Consolidation),
		ConsolidationPreviewConsumed:       productionAuthorizationConsolidationReady(sources.Consolidation),
		OwnerManagedOpaqueBoundaryReady:    productionAuthorizationConsolidationReady(sources.Consolidation),
		CallerStateRootRequired:            false,
		AuthorizationAccepted:              false,
		ProductionReadiness:                false,
		ProductionOwnershipReady:           false,
		ConsumerCount:                      len(consumers),
		RequiredConsumerCount:              6,
		ConsumedConsumerCount:              productionAuthorizationConsumedConsumerCount(consumers),
		MissingConsumerCount:               productionAuthorizationMissingConsumerCount(consumers),
		AuthorizationAcceptedConsumerCount: 0,
		ProductionReadyConsumerCount:       0,
		SideEffectConsumerCount:            0,
		Consumers:                          consumers,
		ConsumerIDs:                        productionAuthorizationConsumerIDs(consumers),
		RequiredBeforeProductionOwnership: []string{
			"production-human-authorization-receipt-consolidation-preview",
			"production-dbus-gate-review-preview",
			"production-dbus-method-review-preview",
			"runtime-service-activation-preflight-preview",
			"runtime-write-gate-preview",
			"production-rollback-diagnostics-review-preview",
			"production-desktop-side-effect-review-preview",
			"separate accepted operator authorization receipt",
			"separate production service ownership proof",
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
		ReceiptWriterEnabled:        false,
		ReceiptPersistenceEnabled:   false,
		ReceiptLookupWritesEnabled:  false,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		ShellConfigurationWritten:   false,
		SettingsPersisted:           false,
		NotificationSent:            false,
		NotificationDeliveryEnabled: false,
		PortalRequestCreated:        false,
		RequestObjectsCreated:       false,
		AdapterInvocationEnabled:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
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
			"treat authorization consumption audit as authorization acceptance",
			"write, persist, accept, replay, or look up authorization receipts from this audit",
			"claim production D-Bus ownership, start services, or enable Runtime writes from this audit",
			"enable desktop writes, notifications, Portal requests, request objects, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, restore snapshots, clean state, read file contents, or expose paths",
			"require network, require privileged containers, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep every production gate consuming the same owner-managed opaque authorization receipt boundary.",
			"Add a separate receipt acceptance path only after explicit operator authorization.",
			"Keep production service start, bus claim, write dispatch, desktop side effects, support side effects, restore, cleanup, launch, and host mutation disabled.",
			"Audit receipt acceptance propagation before any production ownership commit.",
		},
		DesktopSafeSummary: "The production authorization consumption audit proves every production gate points at the consolidated opaque receipt boundary, but it does not accept authorization, write receipts, claim a bus, start services, enable writes, emit desktop side effects, launch engines, or mutate host state.",
	}
	checks := productionAuthorizationConsumptionChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionAuthorizationConsumptionCheckIDs(checks)
	preview.Counts = countProductionAuthorizationConsumptionChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "production-authorization-consumption-audit-ready-authorization-disabled"
	}
	if err := validateNoBackendTerms(preview, "production authorization consumption audit preview"); err != nil {
		return ProductionAuthorizationConsumptionAuditPreview{}, err
	}
	return preview, nil
}

type productionAuthorizationConsumptionSourceSet struct {
	Consolidation       string
	ProductionDBusGate  string
	MethodReview        string
	ServiceActivation   string
	RuntimeWriteGate    string
	RollbackDiagnostics string
	DesktopSideEffects  string
}

func productionAuthorizationConsumptionSources(root string) productionAuthorizationConsumptionSourceSet {
	return productionAuthorizationConsumptionSourceSet{
		Consolidation:       productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_human_authorization_receipt_consolidation.go"}),
		ProductionDBusGate:  productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
		MethodReview:        productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_method_review.go"}),
		ServiceActivation:   productionAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_service_activation_preflight.go"}),
		RuntimeWriteGate:    productionAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_write_gate.go"}),
		RollbackDiagnostics: productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_rollback_diagnostics_review.go"}),
		DesktopSideEffects:  productionAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
	}
}

func productionAuthorizationConsumptionConsumers(sources productionAuthorizationConsumptionSourceSet) []ProductionAuthorizationConsumptionConsumer {
	return []ProductionAuthorizationConsumptionConsumer{
		productionAuthorizationConsumer("production-dbus-gate-review", "production-dbus-gate-review-preview", "internal/runtime/owner/production_dbus_gate_review.go", sources.ProductionDBusGate, []string{"production-human-authorization-receipt-consolidation-preview", "production-human-authorization-receipt-consolidation", "AuthorizationReceiptAccepted"}, "keep production bus ownership disabled after consumption"),
		productionAuthorizationConsumer("production-dbus-method-review", "production-dbus-method-review-preview", "internal/runtime/owner/production_dbus_method_review.go", sources.MethodReview, []string{"production-human-authorization-receipt-consolidation-preview", "production-human-authorization-receipt-consolidation", "human-authorization-receipt"}, "keep D-Bus method exposure disabled after consumption"),
		productionAuthorizationConsumer("runtime-service-activation-preflight", "runtime-service-activation-preflight-preview", "internal/runtime/appidentity/runtime_service_activation_preflight.go", sources.ServiceActivation, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated opaque authorization receipt boundary", "AuthorizationReceiptAccepted"}, "keep service activation disabled after consumption"),
		productionAuthorizationConsumer("runtime-write-gate", "runtime-write-gate-preview", "internal/runtime/appidentity/runtime_write_gate.go", sources.RuntimeWriteGate, []string{"production-human-authorization-receipt-consolidation-preview", "production-human-authorization-receipt-consolidation", "WriteMethodEnabled"}, "keep write dispatch disabled after consumption"),
		productionAuthorizationConsumer("rollback-diagnostics-review", "production-rollback-diagnostics-review-preview", "internal/runtime/owner/production_rollback_diagnostics_review.go", sources.RollbackDiagnostics, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated opaque authorization receipt boundary", "support-side-effects-disabled"}, "keep rollback and support side effects disabled after consumption"),
		productionAuthorizationConsumer("desktop-side-effect-review", "production-desktop-side-effect-review-preview", "internal/runtime/owner/production_desktop_side_effect_review.go", sources.DesktopSideEffects, []string{"production-human-authorization-receipt-consolidation-preview", "consolidated human authorization receipt boundary", "desktop-writes-disabled"}, "keep KDE desktop side effects disabled after consumption"),
	}
}

func productionAuthorizationConsumer(id string, requestType string, sourceFile string, source string, tokens []string, nextRequirement string) ProductionAuthorizationConsumptionConsumer {
	consumes := productionAuthorizationHasAll(source, tokens)
	status := "missing-consumption"
	if consumes {
		status = "consumed-authorization-disabled"
	}
	return ProductionAuthorizationConsumptionConsumer{
		ID:                           id,
		RequestType:                  requestType,
		SourceFile:                   sourceFile,
		ConsumesConsolidatedBoundary: consumes,
		ReceiptBoundaryReady:         consumes,
		ReceiptAccepted:              false,
		AuthorizationAccepted:        false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		ReviewOnly:                   true,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		DesktopSideEffectsEnabled:    false,
		SupportSideEffectsEnabled:    false,
		BackendLaunchEnabled:         false,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		AuditStatus:                  status,
		NextRequirement:              nextRequirement,
	}
}

func productionAuthorizationConsumptionChecks(preview ProductionAuthorizationConsumptionAuditPreview) []ProductionAuthorizationConsumptionAuditCheck {
	return []ProductionAuthorizationConsumptionAuditCheck{
		productionAuthorizationCheck("consolidation-preview-consumed", productionAuthorizationPassBlocked(preview.ConsolidationPreviewConsumed && preview.ReceiptBoundaryConsolidated && preview.OwnerManagedOpaqueBoundaryReady && preview.ReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.OpaqueReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The audit consumes the consolidated opaque human authorization receipt boundary."),
		productionAuthorizationCheck("six-production-consumers-present", productionAuthorizationPassBlocked(preview.ConsumerCount == 6 && preview.RequiredConsumerCount == 6 && preview.MissingConsumerCount == 0), "The audit tracks all six production gate consumers."),
		productionAuthorizationCheck("consumers-use-consolidated-boundary", productionAuthorizationPassBlocked(preview.ConsumedConsumerCount == 6 && productionAuthorizationConsumersReady(preview.Consumers)), "Every production consumer references the consolidated authorization receipt boundary."),
		productionAuthorizationCheck("authorization-not-accepted", productionAuthorizationPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationAccepted && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && preview.AuthorizationAcceptedConsumerCount == 0 && preview.ProductionReadyConsumerCount == 0), "The audit does not accept authorization or mark production readiness."),
		productionAuthorizationCheck("production-ownership-disabled", productionAuthorizationPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady), "Service start, session bus claim, production bus claim, and production ownership remain disabled."),
		productionAuthorizationCheck("write-and-launch-disabled", productionAuthorizationPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.RequestObjectsCreated && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && productionAuthorizationConsumersKeepWritesLaunchDisabled(preview.Consumers)), "Runtime writes, request creation, adapter invocation, and launch remain disabled."),
		productionAuthorizationCheck("desktop-and-support-side-effects-disabled", productionAuthorizationPassBlocked(!preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && preview.SideEffectConsumerCount == 0 && productionAuthorizationConsumersKeepSideEffectsDisabled(preview.Consumers)), "KDE desktop, Portal, support, restore, and cleanup side effects remain disabled."),
		productionAuthorizationCheck("host-boundary-closed", productionAuthorizationPassBlocked(!preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionAuthorizationConsumersKeepHostBoundaryClosed(preview.Consumers)), "Data exposure, network, privilege, internal detail exposure, and host mutation gates remain closed."),
	}
}

func productionAuthorizationCheck(id string, status string, summary string) ProductionAuthorizationConsumptionAuditCheck {
	return ProductionAuthorizationConsumptionAuditCheck{ID: id, Status: status, Summary: summary}
}

func productionAuthorizationPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionAuthorizationConsumptionCheckIDs(checks []ProductionAuthorizationConsumptionAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionAuthorizationConsumptionChecks(checks []ProductionAuthorizationConsumptionAuditCheck) ProductionAuthorizationConsumptionAuditCounts {
	counts := ProductionAuthorizationConsumptionAuditCounts{Total: len(checks)}
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

func productionAuthorizationConsolidationReady(source string) bool {
	return productionAuthorizationHasAll(source, []string{
		"production-human-authorization-receipt-consolidation-preview",
		"owner-managed-opaque-human-authorization-receipt-boundary",
		"OwnerManagedOpaqueReceiptLookupReady",
		"ReceiptBoundaryConsolidated",
		"authorization-not-granted",
		"unsafe-gates-closed",
	})
}

func productionAuthorizationConsumedConsumerCount(consumers []ProductionAuthorizationConsumptionConsumer) int {
	count := 0
	for _, consumer := range consumers {
		if consumer.ConsumesConsolidatedBoundary {
			count++
		}
	}
	return count
}

func productionAuthorizationMissingConsumerCount(consumers []ProductionAuthorizationConsumptionConsumer) int {
	count := 0
	for _, consumer := range consumers {
		if !consumer.ConsumesConsolidatedBoundary {
			count++
		}
	}
	return count
}

func productionAuthorizationConsumerIDs(consumers []ProductionAuthorizationConsumptionConsumer) []string {
	ids := make([]string, 0, len(consumers))
	for _, consumer := range consumers {
		ids = append(ids, consumer.ID)
	}
	return ids
}

func productionAuthorizationConsumersReady(consumers []ProductionAuthorizationConsumptionConsumer) bool {
	if len(consumers) != 6 {
		return false
	}
	for _, consumer := range consumers {
		if !consumer.ConsumesConsolidatedBoundary || !consumer.ReceiptBoundaryReady || consumer.AuditStatus != "consumed-authorization-disabled" {
			return false
		}
		if consumer.ReceiptAccepted || consumer.AuthorizationAccepted || consumer.ProductionReadiness || consumer.ProductionOwnershipReady || consumer.KDEPolicyOwner || !consumer.ReviewOnly {
			return false
		}
	}
	return true
}

func productionAuthorizationConsumersKeepWritesLaunchDisabled(consumers []ProductionAuthorizationConsumptionConsumer) bool {
	for _, consumer := range consumers {
		if consumer.WriteMethodsEnabled || consumer.RuntimeWritesEnabled || consumer.BackendLaunchEnabled {
			return false
		}
	}
	return true
}

func productionAuthorizationConsumersKeepSideEffectsDisabled(consumers []ProductionAuthorizationConsumptionConsumer) bool {
	for _, consumer := range consumers {
		if consumer.DesktopSideEffectsEnabled || consumer.SupportSideEffectsEnabled {
			return false
		}
	}
	return true
}

func productionAuthorizationConsumersKeepHostBoundaryClosed(consumers []ProductionAuthorizationConsumptionConsumer) bool {
	for _, consumer := range consumers {
		if consumer.HostRootModified || consumer.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionAuthorizationHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

func productionAuthorizationReadSources(root string, paths []string) string {
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
