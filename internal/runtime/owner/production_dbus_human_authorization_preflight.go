package owner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const ProductionDBusHumanAuthorizationReceiptID = "production-dbus-human-authorization-receipt-id"

type ProductionDBusHumanAuthorizationPreflightPreview struct {
	Version                            string                                           `json:"version"`
	SchemaVersion                      string                                           `json:"schema_version"`
	RequestType                        string                                           `json:"request_type"`
	PreflightType                      string                                           `json:"preflight_type"`
	Source                             string                                           `json:"source"`
	GateReviewPresent                  bool                                             `json:"gate_review_present"`
	RouteInventoryPresent              bool                                             `json:"route_inventory_present"`
	ExplicitHumanAuthorizationRequired bool                                             `json:"explicit_human_authorization_required"`
	AuthorizationReceiptSchema         string                                           `json:"authorization_receipt_schema"`
	AuthorizationReceiptID             string                                           `json:"authorization_receipt_id"`
	AuthorizationReceiptRequired       bool                                             `json:"authorization_receipt_required"`
	AuthorizationReceiptPresent        bool                                             `json:"authorization_receipt_present"`
	AuthorizationGrantReady            bool                                             `json:"authorization_grant_ready"`
	AuthorizationAccepted              bool                                             `json:"authorization_accepted"`
	PreflightReady                     bool                                             `json:"preflight_ready"`
	ProductionReadiness                bool                                             `json:"production_readiness"`
	ProductionDBusGateReviewRequired   bool                                             `json:"production_dbus_gate_review_required"`
	RequiredInputs                     []string                                         `json:"required_inputs"`
	Checks                             []ProductionDBusHumanAuthorizationPreflightCheck `json:"checks"`
	CheckIDs                           []string                                         `json:"check_ids"`
	Counts                             ProductionDBusHumanAuthorizationPreflightCounts  `json:"counts"`
	RuntimeOwned                       bool                                             `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                             `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                             `json:"kde_policy_owner"`
	SystemServiceStarted               bool                                             `json:"system_service_started"`
	SessionBusClaimed                  bool                                             `json:"session_bus_claimed"`
	ProductionBusClaimed               bool                                             `json:"production_bus_claimed"`
	ProductionOwnerEnabled             bool                                             `json:"production_owner_enabled"`
	ProductionActivationReady          bool                                             `json:"production_activation_ready"`
	WriteMethodsEnabled                bool                                             `json:"write_methods_enabled"`
	RuntimeWritesEnabled               bool                                             `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled           bool                                             `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled               bool                                             `json:"backend_launch_enabled"`
	BackendProcessStarted              bool                                             `json:"backend_process_started"`
	SupportBundleExported              bool                                             `json:"support_bundle_exported"`
	SupportCaseCreated                 bool                                             `json:"support_case_created"`
	NotificationSent                   bool                                             `json:"notification_sent"`
	NetworkRequired                    bool                                             `json:"network_required"`
	HostRootModified                   bool                                             `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                                             `json:"privileged_container_required"`
	StateRootPathExposed               bool                                             `json:"state_root_path_exposed"`
	RawCommandExposed                  bool                                             `json:"raw_command_exposed"`
	RawExecutableExposed               bool                                             `json:"raw_executable_exposed"`
	BackendDetailsExposed              bool                                             `json:"backend_details_exposed"`
	BlockedActions                     []string                                         `json:"blocked_actions"`
	NextRequirements                   []string                                         `json:"next_requirements"`
	DesktopSafeSummary                 string                                           `json:"desktop_safe_summary"`
}

type ProductionDBusHumanAuthorizationPreflightCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionDBusHumanAuthorizationPreflightCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionDBusHumanAuthorizationPreflightPreview(root string) (ProductionDBusHumanAuthorizationPreflightPreview, error) {
	version, err := readProductionDBusHumanAuthorizationPreflightVersion(root)
	if err != nil {
		return ProductionDBusHumanAuthorizationPreflightPreview{}, err
	}
	sources := productionDBusHumanAuthorizationPreflightSources(root)
	preview := ProductionDBusHumanAuthorizationPreflightPreview{
		Version:                            version,
		SchemaVersion:                      "xnix.runtime.production_dbus_human_authorization_preflight.v1",
		RequestType:                        "production-dbus-human-authorization-preflight-preview",
		PreflightType:                      "read-only-production-dbus-human-authorization-preflight",
		Source:                             "production-dbus-gate-review+explicit-human-authorization-boundary",
		GateReviewPresent:                  productionDBusHumanAuthorizationHasGateReview(sources.GateReview),
		RouteInventoryPresent:              productionDBusHumanAuthorizationHasRouteInventory(sources.GateReview),
		ExplicitHumanAuthorizationRequired: productionDBusHumanAuthorizationHasExplicitHumanGate(sources.GateReview),
		AuthorizationReceiptSchema:         "xnix.runtime.production_dbus_human_authorization_receipt.v1",
		AuthorizationReceiptID:             ProductionDBusHumanAuthorizationReceiptID,
		AuthorizationReceiptRequired:       true,
		AuthorizationReceiptPresent:        false,
		AuthorizationGrantReady:            false,
		AuthorizationAccepted:              false,
		PreflightReady:                     false,
		ProductionReadiness:                false,
		ProductionDBusGateReviewRequired:   true,
		RequiredInputs: []string{
			"production-dbus-gate-review-preview",
			"human-readable production ownership explanation",
			"explicit operator approval outside the preview command",
			"route-by-route production D-Bus method review",
			"Runtime write-gate review",
			"rollback and diagnostics review",
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
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat preflight presence as human authorization",
			"write or accept an authorization receipt from the preview command",
			"claim production D-Bus ownership from authorization preflight",
			"start the production Runtime service from authorization preflight",
			"enable Runtime writes, adapter invocation, backend launch, support side effects, or host mutation",
			"expose state-root paths, raw commands, raw executables, backend details, or internal adapter identifiers",
		},
		NextRequirements: []string{
			"Keep the authorization preflight read-only until a separate receipt writer is explicitly authorized.",
			"Require the production D-Bus gate review packet before any authorization receipt can be accepted.",
			"Require route-by-route production D-Bus method, Runtime write-gate, and rollback reviews before production ownership.",
			"Keep service start, bus claim, writes, adapter invocation, backend launch, desktop side effects, and host mutation disabled.",
		},
		DesktopSafeSummary: "The human authorization preflight defines the receipt shape required by the production D-Bus gate but does not grant authorization, write receipts, claim a bus, start services, enable writes, launch backends, emit desktop side effects, or mutate host state.",
	}
	checks := productionDBusHumanAuthorizationPreflightChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionDBusHumanAuthorizationPreflightCheckIDs(checks)
	preview.Counts = countProductionDBusHumanAuthorizationPreflightChecks(checks)
	preview.PreflightReady = preview.Counts.Blocked == 0 && preview.GateReviewPresent && preview.RouteInventoryPresent && preview.ExplicitHumanAuthorizationRequired
	if err := validateNoBackendTerms(preview, "production D-Bus human authorization preflight preview"); err != nil {
		return ProductionDBusHumanAuthorizationPreflightPreview{}, err
	}
	return preview, nil
}

type productionDBusHumanAuthorizationPreflightSourceSet struct {
	GateReview string
}

func productionDBusHumanAuthorizationPreflightSources(root string) productionDBusHumanAuthorizationPreflightSourceSet {
	return productionDBusHumanAuthorizationPreflightSourceSet{
		GateReview: readProductionDBusHumanAuthorizationPreflightSources(root, []string{"internal/runtime/owner/production_dbus_gate_review.go"}),
	}
}

func productionDBusHumanAuthorizationPreflightChecks(preview ProductionDBusHumanAuthorizationPreflightPreview) []ProductionDBusHumanAuthorizationPreflightCheck {
	return []ProductionDBusHumanAuthorizationPreflightCheck{
		productionDBusHumanAuthorizationPreflightCheck("gate-review-present", productionDBusHumanAuthorizationPreflightPassBlocked(preview.GateReviewPresent), "The production D-Bus gate review packet exists."),
		productionDBusHumanAuthorizationPreflightCheck("route-inventory-present", productionDBusHumanAuthorizationPreflightPassBlocked(preview.RouteInventoryPresent), "The gate review inventories smoke-covered owner-local routes."),
		productionDBusHumanAuthorizationPreflightCheck("human-authorization-required", productionDBusHumanAuthorizationPreflightPassBlocked(preview.ExplicitHumanAuthorizationRequired && preview.AuthorizationReceiptRequired), "The gate review explicitly requires human authorization before production ownership."),
		productionDBusHumanAuthorizationPreflightCheck("receipt-shape-declared", productionDBusHumanAuthorizationPreflightPassBlocked(preview.AuthorizationReceiptSchema == "xnix.runtime.production_dbus_human_authorization_receipt.v1" && preview.AuthorizationReceiptID == ProductionDBusHumanAuthorizationReceiptID), "The preflight declares a stable future authorization receipt shape."),
		productionDBusHumanAuthorizationPreflightCheck("authorization-not-granted", productionDBusHumanAuthorizationPreflightPassBlocked(!preview.AuthorizationReceiptPresent && !preview.AuthorizationGrantReady && !preview.AuthorizationAccepted && !preview.ProductionReadiness), "The preview does not grant authorization or mark production readiness."),
		productionDBusHumanAuthorizationPreflightCheck("production-ownership-disabled", productionDBusHumanAuthorizationPreflightPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady), "The preview does not start services or claim D-Bus ownership."),
		productionDBusHumanAuthorizationPreflightCheck("unsafe-gates-closed", productionDBusHumanAuthorizationPreflightPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.NotificationSent && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed), "The preview keeps writes, launch, side effects, host mutation, and internal detail exposure disabled."),
	}
}

func productionDBusHumanAuthorizationPreflightCheck(id string, status string, summary string) ProductionDBusHumanAuthorizationPreflightCheck {
	return ProductionDBusHumanAuthorizationPreflightCheck{ID: id, Status: status, Summary: summary}
}

func productionDBusHumanAuthorizationPreflightPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionDBusHumanAuthorizationPreflightCheckIDs(checks []ProductionDBusHumanAuthorizationPreflightCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionDBusHumanAuthorizationPreflightChecks(checks []ProductionDBusHumanAuthorizationPreflightCheck) ProductionDBusHumanAuthorizationPreflightCounts {
	counts := ProductionDBusHumanAuthorizationPreflightCounts{Total: len(checks)}
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

func productionDBusHumanAuthorizationHasGateReview(source string) bool {
	return productionDBusHumanAuthorizationTextHasAll(source, []string{"ProductionDBusGateReviewPreview", "production-dbus-gate-review-preview", "xnix.runtime.production_dbus_gate_review.v1"})
}

func productionDBusHumanAuthorizationHasRouteInventory(source string) bool {
	return productionDBusHumanAuthorizationTextHasAll(source, []string{"materialization-fanout-owner-route", "restricted-smoke-fanout-owner-route", "redacted-adapter-profile-owner-route", "SmokeCoveredRouteCount"})
}

func productionDBusHumanAuthorizationHasExplicitHumanGate(source string) bool {
	return productionDBusHumanAuthorizationTextHasAll(source, []string{"HumanAuthorizationRequired", "human-authorization-required", "human-authorization"})
}

func productionDBusHumanAuthorizationTextHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}

func readProductionDBusHumanAuthorizationPreflightSources(root string, paths []string) string {
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

func readProductionDBusHumanAuthorizationPreflightVersion(root string) (string, error) {
	content, err := os.ReadFile(filepath.Join(root, "VERSION"))
	if err != nil {
		return "", fmt.Errorf("read VERSION: %w", err)
	}
	version := strings.TrimSpace(string(content))
	if version == "" {
		return "", fmt.Errorf("read VERSION: empty version")
	}
	return version, nil
}
