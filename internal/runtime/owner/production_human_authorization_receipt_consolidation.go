package owner

import (
	"os"
	"path/filepath"
	"strings"
)

type ProductionHumanAuthorizationReceiptConsolidationPreview struct {
	Version                               string                                                  `json:"version"`
	SchemaVersion                         string                                                  `json:"schema_version"`
	RequestType                           string                                                  `json:"request_type"`
	ConsolidationType                     string                                                  `json:"consolidation_type"`
	Source                                string                                                  `json:"source"`
	ConsolidationDecision                 string                                                  `json:"consolidation_decision"`
	ReceiptSchema                         string                                                  `json:"receipt_schema"`
	OpaqueReceiptID                       string                                                  `json:"opaque_receipt_id"`
	ReceiptRequired                       bool                                                    `json:"receipt_required"`
	ReceiptPresent                        bool                                                    `json:"receipt_present"`
	ReceiptAccepted                       bool                                                    `json:"receipt_accepted"`
	ReceiptBoundaryConsolidated           bool                                                    `json:"receipt_boundary_consolidated"`
	OwnerManagedOpaqueReceiptLookupReady  bool                                                    `json:"owner_managed_opaque_receipt_lookup_ready"`
	CallerStateRootRequired               bool                                                    `json:"caller_state_root_required"`
	ExplicitOperatorActionRequired        bool                                                    `json:"explicit_operator_action_required"`
	AuthorizationGrantReady               bool                                                    `json:"authorization_grant_ready"`
	AuthorizationAccepted                 bool                                                    `json:"authorization_accepted"`
	ProductionReadiness                   bool                                                    `json:"production_readiness"`
	ProductionOwnershipReady              bool                                                    `json:"production_ownership_ready"`
	GateCount                             int                                                     `json:"gate_count"`
	RequiredGateCount                     int                                                     `json:"required_gate_count"`
	ConsumedGateCount                     int                                                     `json:"consumed_gate_count"`
	MissingGateCount                      int                                                     `json:"missing_gate_count"`
	AuthorizationAcceptedGateCount        int                                                     `json:"authorization_accepted_gate_count"`
	ProductionReadyGateCount              int                                                     `json:"production_ready_gate_count"`
	Gates                                 []ProductionHumanAuthorizationReceiptGate               `json:"gates"`
	GateIDs                               []string                                                `json:"gate_ids"`
	RequiredBeforeAuthorizationAcceptance []string                                                `json:"required_before_authorization_acceptance"`
	Checks                                []ProductionHumanAuthorizationReceiptConsolidationCheck `json:"checks"`
	CheckIDs                              []string                                                `json:"check_ids"`
	Counts                                ProductionHumanAuthorizationReceiptConsolidationCounts  `json:"counts"`
	RuntimeOwned                          bool                                                    `json:"runtime_owned"`
	GoRuntimeBacked                       bool                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool                                                    `json:"kde_policy_owner"`
	SystemServiceStarted                  bool                                                    `json:"system_service_started"`
	SessionBusClaimed                     bool                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed                  bool                                                    `json:"production_bus_claimed"`
	ProductionOwnerEnabled                bool                                                    `json:"production_owner_enabled"`
	ProductionActivationReady             bool                                                    `json:"production_activation_ready"`
	WriteMethodsEnabled                   bool                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled                  bool                                                    `json:"runtime_writes_enabled"`
	ReceiptWriterEnabled                  bool                                                    `json:"receipt_writer_enabled"`
	ReceiptPersistenceEnabled             bool                                                    `json:"receipt_persistence_enabled"`
	ReceiptLookupWritesEnabled            bool                                                    `json:"receipt_lookup_writes_enabled"`
	DesktopFilesWritten                   bool                                                    `json:"desktop_files_written"`
	MIMEAppsWritten                       bool                                                    `json:"mimeapps_written"`
	ShellConfigurationWritten             bool                                                    `json:"shell_configuration_written"`
	SettingsPersisted                     bool                                                    `json:"settings_persisted"`
	NotificationSent                      bool                                                    `json:"notification_sent"`
	NotificationDeliveryEnabled           bool                                                    `json:"notification_delivery_enabled"`
	PortalRequestCreated                  bool                                                    `json:"portal_request_created"`
	RequestObjectsCreated                 bool                                                    `json:"request_objects_created"`
	AdapterInvocationEnabled              bool                                                    `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled                  bool                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted                 bool                                                    `json:"backend_process_started"`
	SupportBundleExported                 bool                                                    `json:"support_bundle_exported"`
	SupportCaseCreated                    bool                                                    `json:"support_case_created"`
	SnapshotRestoreExecuted               bool                                                    `json:"snapshot_restore_executed"`
	StateCleanupExecuted                  bool                                                    `json:"state_cleanup_executed"`
	FileContentRead                       bool                                                    `json:"file_content_read"`
	FilePathsExposed                      bool                                                    `json:"file_paths_exposed"`
	NetworkRequired                       bool                                                    `json:"network_required"`
	HostRootModified                      bool                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired           bool                                                    `json:"privileged_container_required"`
	StateRootPathExposed                  bool                                                    `json:"state_root_path_exposed"`
	RawCommandExposed                     bool                                                    `json:"raw_command_exposed"`
	RawExecutableExposed                  bool                                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                 bool                                                    `json:"backend_details_exposed"`
	BlockedActions                        []string                                                `json:"blocked_actions"`
	NextRequirements                      []string                                                `json:"next_requirements"`
	DesktopSafeSummary                    string                                                  `json:"desktop_safe_summary"`
}

type ProductionHumanAuthorizationReceiptGate struct {
	ID                           string `json:"id"`
	RequestType                  string `json:"request_type"`
	GateType                     string `json:"gate_type"`
	EvidencePresent              bool   `json:"evidence_present"`
	ReceiptBoundaryReady         bool   `json:"receipt_boundary_ready"`
	HumanAuthorizationRequired   bool   `json:"human_authorization_required"`
	AuthorizationReceiptAccepted bool   `json:"authorization_receipt_accepted"`
	ProductionReadiness          bool   `json:"production_readiness"`
	ProductionOwnershipReady     bool   `json:"production_ownership_ready"`
	RuntimeOwned                 bool   `json:"runtime_owned"`
	GoRuntimeBacked              bool   `json:"go_runtime_backed"`
	KDEPolicyOwner               bool   `json:"kde_policy_owner"`
	ReviewOnly                   bool   `json:"review_only"`
	SideEffectsDisabled          bool   `json:"side_effects_disabled"`
	WriteMethodsEnabled          bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled         bool   `json:"runtime_writes_enabled"`
	BackendLaunchEnabled         bool   `json:"backend_launch_enabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	ConsolidationStatus          string `json:"consolidation_status"`
	NextRequirement              string `json:"next_requirement"`
}

type ProductionHumanAuthorizationReceiptConsolidationCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionHumanAuthorizationReceiptConsolidationCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionHumanAuthorizationReceiptConsolidationPreview(root string) (ProductionHumanAuthorizationReceiptConsolidationPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionHumanAuthorizationReceiptConsolidationPreview{}, err
	}
	sources := productionHumanAuthorizationReceiptConsolidationSources(root)
	gates := productionHumanAuthorizationReceiptGates(sources)
	preview := ProductionHumanAuthorizationReceiptConsolidationPreview{
		Version:                              version,
		SchemaVersion:                        "xnix.runtime.production_human_authorization_receipt_consolidation.v1",
		RequestType:                          "production-human-authorization-receipt-consolidation-preview",
		ConsolidationType:                    "owner-managed-opaque-human-authorization-receipt-boundary",
		Source:                               "production-dbus-human-authorization-preflight-preview+production-dbus-gate-review-preview+production-dbus-method-review-preview+runtime-service-activation-preflight-preview+runtime-write-gate-preview+production-rollback-diagnostics-review-preview+production-desktop-side-effect-review-preview",
		ConsolidationDecision:                "production-human-authorization-receipt-consolidation-blocked",
		ReceiptSchema:                        "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		OpaqueReceiptID:                      ProductionDBusHumanAuthorizationReceiptID,
		ReceiptRequired:                      true,
		ReceiptPresent:                       false,
		ReceiptAccepted:                      false,
		ReceiptBoundaryConsolidated:          productionHumanAuthorizationAllGatesReady(gates),
		OwnerManagedOpaqueReceiptLookupReady: true,
		CallerStateRootRequired:              false,
		ExplicitOperatorActionRequired:       true,
		AuthorizationGrantReady:              false,
		AuthorizationAccepted:                false,
		ProductionReadiness:                  false,
		ProductionOwnershipReady:             false,
		GateCount:                            len(gates),
		RequiredGateCount:                    7,
		ConsumedGateCount:                    productionHumanAuthorizationConsumedGateCount(gates),
		MissingGateCount:                     productionHumanAuthorizationMissingGateCount(gates),
		AuthorizationAcceptedGateCount:       0,
		ProductionReadyGateCount:             0,
		Gates:                                gates,
		GateIDs:                              productionHumanAuthorizationGateIDs(gates),
		RequiredBeforeAuthorizationAcceptance: []string{
			"production-dbus-human-authorization-preflight-preview",
			"production-dbus-gate-review-preview",
			"production-dbus-method-review-preview",
			"runtime-service-activation-preflight-preview",
			"runtime-write-gate-preview",
			"production-rollback-diagnostics-review-preview",
			"production-desktop-side-effect-review-preview",
			"separate operator authorization action outside this preview",
			"separate receipt writer authorization outside this preview",
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
			"treat receipt boundary consolidation as operator authorization",
			"write, persist, accept, replay, or look up authorization receipts from this preview",
			"claim production D-Bus ownership, start services, or enable Runtime writes from this preview",
			"enable desktop writes, notifications, Portal requests, request objects, adapter invocation, or compatibility engine launch",
			"export support bundles, create support cases, restore snapshots, clean state, read file contents, or expose paths",
			"require network, require privileged containers, expose raw commands, expose internal engine details, or mutate host root",
		},
		NextRequirements: []string{
			"Keep the consolidated receipt boundary read-only until an explicit receipt writer is separately authorized.",
			"Require every production gate to consume the same opaque receipt id before any production ownership decision.",
			"Keep the receipt writer outside preview commands and require a separate operator action for acceptance.",
			"Keep service start, bus claim, writes, desktop side effects, support side effects, restore, cleanup, launch, and host mutation disabled.",
		},
		DesktopSafeSummary: "The production human authorization receipt consolidation joins all production gates around one owner-managed opaque receipt boundary, but it does not grant authorization, write receipts, accept receipts, claim a bus, start services, enable writes, emit desktop side effects, launch engines, or mutate host state.",
	}
	checks := productionHumanAuthorizationReceiptConsolidationChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionHumanAuthorizationReceiptCheckIDs(checks)
	preview.Counts = countProductionHumanAuthorizationReceiptChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ConsolidationDecision = "production-human-authorization-receipt-consolidation-ready-authorization-disabled"
	}
	if err := validateNoBackendTerms(preview, "production human authorization receipt consolidation preview"); err != nil {
		return ProductionHumanAuthorizationReceiptConsolidationPreview{}, err
	}
	return preview, nil
}

type productionHumanAuthorizationReceiptConsolidationSourceSet struct {
	HumanAuthorizationPreflight string
	ProductionDBusGateReview    string
	ProductionDBusMethodReview  string
	ServiceActivationPreflight  string
	RuntimeWriteGate            string
	RollbackDiagnosticsReview   string
	DesktopSideEffectReview     string
}

func productionHumanAuthorizationReceiptConsolidationSources(root string) productionHumanAuthorizationReceiptConsolidationSourceSet {
	return productionHumanAuthorizationReceiptConsolidationSourceSet{
		HumanAuthorizationPreflight: productionHumanAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_human_authorization_preflight.go"}),
		ProductionDBusGateReview:    productionHumanAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
		ProductionDBusMethodReview:  productionHumanAuthorizationReadSources(root, []string{"internal/runtime/owner/production_dbus_method_review.go"}),
		ServiceActivationPreflight:  productionHumanAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_service_activation_preflight.go"}),
		RuntimeWriteGate:            productionHumanAuthorizationReadSources(root, []string{"internal/runtime/appidentity/runtime_write_gate.go"}),
		RollbackDiagnosticsReview:   productionHumanAuthorizationReadSources(root, []string{"internal/runtime/owner/production_rollback_diagnostics_review.go"}),
		DesktopSideEffectReview:     productionHumanAuthorizationReadSources(root, []string{"internal/runtime/owner/production_desktop_side_effect_review.go"}),
	}
}

func productionHumanAuthorizationReceiptGates(sources productionHumanAuthorizationReceiptConsolidationSourceSet) []ProductionHumanAuthorizationReceiptGate {
	return []ProductionHumanAuthorizationReceiptGate{
		productionHumanAuthorizationReceiptGate("human-authorization-preflight", "production-dbus-human-authorization-preflight-preview", "read-only-production-dbus-human-authorization-preflight", productionHumanAuthorizationHasAll(sources.HumanAuthorizationPreflight, []string{"production-dbus-human-authorization-preflight-preview", "xnix.runtime.production_dbus_human_authorization_receipt.v1", "AuthorizationReceiptRequired", "AuthorizationAccepted"}), "keep the authorization receipt shape read-only"),
		productionHumanAuthorizationReceiptGate("production-dbus-gate-review", "production-dbus-gate-review-preview", "owner-local-smoke-covered-production-dbus-gate-review", productionHumanAuthorizationHasAll(sources.ProductionDBusGateReview, []string{"production-dbus-gate-review-preview", "HumanAuthorizationPreflightReady", "AuthorizationReceiptAccepted", "desktop-side-effect-review", "rollback-and-diagnostics-review"}), "keep production D-Bus ownership disabled until receipt acceptance is separately authorized"),
		productionHumanAuthorizationReceiptGate("production-dbus-method-review", "production-dbus-method-review-preview", "route-by-route-production-dbus-method-review", productionHumanAuthorizationHasAll(sources.ProductionDBusMethodReview, []string{"production-dbus-method-review-preview", "human-authorization-receipt", "production-exposure-disabled", "NewProductionDBusMethodReviewPreview"}), "keep method exposure disabled until all production gates consume the receipt"),
		productionHumanAuthorizationReceiptGate("runtime-service-activation-preflight", "runtime-service-activation-preflight-preview", "production-runtime-service-activation-preflight", productionHumanAuthorizationHasAll(sources.ServiceActivationPreflight, []string{"runtime-service-activation-preflight-preview", "human-authorization-receipt", "AuthorizationReceiptAccepted", "ProductionActivationReady"}), "keep service start and stable bus claim disabled until receipt acceptance is separately authorized"),
		productionHumanAuthorizationReceiptGate("runtime-write-gate", "runtime-write-gate-preview", "runtime-write-gate", productionHumanAuthorizationHasAll(sources.RuntimeWriteGate, []string{"runtime-write-gate-preview", "human-authorization-receipt", "production-gates-consumed-write-gate-disabled", "WriteMethodEnabled"}), "keep write dispatch disabled after receipt boundary consolidation"),
		productionHumanAuthorizationReceiptGate("rollback-diagnostics-review", "production-rollback-diagnostics-review-preview", "production-dbus-rollback-diagnostics-review", productionHumanAuthorizationHasAll(sources.RollbackDiagnosticsReview, []string{"production-rollback-diagnostics-review-preview", "production-rollback-diagnostics-review-ready-side-effects-disabled", "support-side-effects-disabled", "restore-and-cleanup-disabled"}), "keep support export, restore, cleanup, and diagnostics side effects disabled"),
		productionHumanAuthorizationReceiptGate("desktop-side-effect-review", "production-desktop-side-effect-review-preview", "kde-production-desktop-side-effect-review", productionHumanAuthorizationHasAll(sources.DesktopSideEffectReview, []string{"production-desktop-side-effect-review-preview", "production-desktop-side-effect-review-ready-side-effects-disabled", "seven-kde-surfaces-reviewed", "desktop-writes-disabled"}), "keep all seven KDE production side effects disabled"),
	}
}

func productionHumanAuthorizationReceiptGate(id string, requestType string, gateType string, evidencePresent bool, nextRequirement string) ProductionHumanAuthorizationReceiptGate {
	status := "missing-source"
	if evidencePresent {
		status = "consumed-authorization-disabled"
	}
	return ProductionHumanAuthorizationReceiptGate{
		ID:                           id,
		RequestType:                  requestType,
		GateType:                     gateType,
		EvidencePresent:              evidencePresent,
		ReceiptBoundaryReady:         evidencePresent,
		HumanAuthorizationRequired:   true,
		AuthorizationReceiptAccepted: false,
		ProductionReadiness:          false,
		ProductionOwnershipReady:     false,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		ReviewOnly:                   true,
		SideEffectsDisabled:          true,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		BackendLaunchEnabled:         false,
		HostRootModified:             false,
		InternalDetailsExposed:       false,
		ConsolidationStatus:          status,
		NextRequirement:              nextRequirement,
	}
}

func productionHumanAuthorizationReceiptConsolidationChecks(preview ProductionHumanAuthorizationReceiptConsolidationPreview) []ProductionHumanAuthorizationReceiptConsolidationCheck {
	return []ProductionHumanAuthorizationReceiptConsolidationCheck{
		productionHumanAuthorizationReceiptCheck("preflight-shape-consumed", productionHumanAuthorizationReceiptPassBlocked(productionHumanAuthorizationGateReady(preview.Gates, "human-authorization-preflight") && preview.ReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.OpaqueReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The consolidation consumes the human authorization preflight receipt schema and opaque id."),
		productionHumanAuthorizationReceiptCheck("production-gate-consumed", productionHumanAuthorizationReceiptPassBlocked(productionHumanAuthorizationGateReady(preview.Gates, "production-dbus-gate-review")), "The consolidation consumes the production D-Bus gate review."),
		productionHumanAuthorizationReceiptCheck("method-review-consumed", productionHumanAuthorizationReceiptPassBlocked(productionHumanAuthorizationGateReady(preview.Gates, "production-dbus-method-review")), "The consolidation consumes the route-by-route production D-Bus method review."),
		productionHumanAuthorizationReceiptCheck("service-and-write-gates-consumed", productionHumanAuthorizationReceiptPassBlocked(productionHumanAuthorizationGateReady(preview.Gates, "runtime-service-activation-preflight") && productionHumanAuthorizationGateReady(preview.Gates, "runtime-write-gate")), "The consolidation consumes service activation and write-gate receipt requirements."),
		productionHumanAuthorizationReceiptCheck("rollback-and-desktop-reviews-consumed", productionHumanAuthorizationReceiptPassBlocked(productionHumanAuthorizationGateReady(preview.Gates, "rollback-diagnostics-review") && productionHumanAuthorizationGateReady(preview.Gates, "desktop-side-effect-review")), "The consolidation consumes rollback, diagnostics, and KDE desktop side-effect reviews."),
		productionHumanAuthorizationReceiptCheck("opaque-receipt-boundary", productionHumanAuthorizationReceiptPassBlocked(preview.ReceiptBoundaryConsolidated && preview.OwnerManagedOpaqueReceiptLookupReady && !preview.CallerStateRootRequired && preview.ReceiptRequired && !preview.ReceiptWriterEnabled && !preview.ReceiptPersistenceEnabled && !preview.ReceiptLookupWritesEnabled), "The receipt boundary is owner-managed, opaque, read-only, and not caller-state-root based."),
		productionHumanAuthorizationReceiptCheck("authorization-not-granted", productionHumanAuthorizationReceiptPassBlocked(!preview.ReceiptPresent && !preview.ReceiptAccepted && !preview.AuthorizationGrantReady && !preview.AuthorizationAccepted && !preview.ProductionReadiness && !preview.ProductionOwnershipReady && preview.AuthorizationAcceptedGateCount == 0 && preview.ProductionReadyGateCount == 0), "The consolidation does not grant or accept authorization."),
		productionHumanAuthorizationReceiptCheck("unsafe-gates-closed", productionHumanAuthorizationReceiptPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.DesktopFilesWritten && !preview.MIMEAppsWritten && !preview.ShellConfigurationWritten && !preview.SettingsPersisted && !preview.NotificationSent && !preview.NotificationDeliveryEnabled && !preview.PortalRequestCreated && !preview.RequestObjectsCreated && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.SnapshotRestoreExecuted && !preview.StateCleanupExecuted && !preview.FileContentRead && !preview.FilePathsExposed && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionHumanAuthorizationReceiptGatesSafe(preview.Gates)), "Service, bus, write, desktop, support, restore, cleanup, launch, data exposure, network, privileged, and host mutation gates remain closed."),
	}
}

func productionHumanAuthorizationReceiptCheck(id string, status string, summary string) ProductionHumanAuthorizationReceiptConsolidationCheck {
	return ProductionHumanAuthorizationReceiptConsolidationCheck{ID: id, Status: status, Summary: summary}
}

func productionHumanAuthorizationReceiptPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionHumanAuthorizationReceiptCheckIDs(checks []ProductionHumanAuthorizationReceiptConsolidationCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionHumanAuthorizationReceiptChecks(checks []ProductionHumanAuthorizationReceiptConsolidationCheck) ProductionHumanAuthorizationReceiptConsolidationCounts {
	counts := ProductionHumanAuthorizationReceiptConsolidationCounts{Total: len(checks)}
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

func productionHumanAuthorizationConsumedGateCount(gates []ProductionHumanAuthorizationReceiptGate) int {
	count := 0
	for _, gate := range gates {
		if gate.ReceiptBoundaryReady {
			count++
		}
	}
	return count
}

func productionHumanAuthorizationMissingGateCount(gates []ProductionHumanAuthorizationReceiptGate) int {
	count := 0
	for _, gate := range gates {
		if !gate.EvidencePresent {
			count++
		}
	}
	return count
}

func productionHumanAuthorizationAllGatesReady(gates []ProductionHumanAuthorizationReceiptGate) bool {
	if len(gates) != 7 {
		return false
	}
	for _, gate := range gates {
		if !gate.ReceiptBoundaryReady {
			return false
		}
	}
	return true
}

func productionHumanAuthorizationGateIDs(gates []ProductionHumanAuthorizationReceiptGate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		ids = append(ids, gate.ID)
	}
	return ids
}

func productionHumanAuthorizationGateReady(gates []ProductionHumanAuthorizationReceiptGate, id string) bool {
	for _, gate := range gates {
		if gate.ID == id {
			return gate.ReceiptBoundaryReady
		}
	}
	return false
}

func productionHumanAuthorizationReceiptGatesSafe(gates []ProductionHumanAuthorizationReceiptGate) bool {
	for _, gate := range gates {
		if gate.AuthorizationReceiptAccepted || gate.ProductionReadiness || gate.ProductionOwnershipReady || gate.KDEPolicyOwner || !gate.ReviewOnly || !gate.SideEffectsDisabled || gate.WriteMethodsEnabled || gate.RuntimeWritesEnabled || gate.BackendLaunchEnabled || gate.HostRootModified || gate.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionHumanAuthorizationHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

func productionHumanAuthorizationReadSources(root string, paths []string) string {
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
