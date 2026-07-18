package owner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type ProductionDBusGateReviewPreview struct {
	Version                          string                          `json:"version"`
	SchemaVersion                    string                          `json:"schema_version"`
	RequestType                      string                          `json:"request_type"`
	GateType                         string                          `json:"gate_type"`
	Source                           string                          `json:"source"`
	RouteCount                       int                             `json:"route_count"`
	SmokeCoveredRouteCount           int                             `json:"smoke_covered_route_count"`
	GateDecision                     string                          `json:"gate_decision"`
	GateDecisionReason               string                          `json:"gate_decision_reason"`
	CurrentGateStatus                string                          `json:"current_gate_status"`
	ProductionReadiness              bool                            `json:"production_readiness"`
	HumanAuthorizationRequired       bool                            `json:"human_authorization_required"`
	HumanAuthorizationPreflightReady bool                            `json:"human_authorization_preflight_ready"`
	HumanAuthorizationGranted        bool                            `json:"human_authorization_granted"`
	AuthorizationReceiptAccepted     bool                            `json:"authorization_receipt_accepted"`
	Routes                           []ProductionDBusGateReviewRoute `json:"routes"`
	RouteIDs                         []string                        `json:"route_ids"`
	RequiredGates                    []string                        `json:"required_gates"`
	Checks                           []ProductionDBusGateReviewCheck `json:"checks"`
	CheckIDs                         []string                        `json:"check_ids"`
	Counts                           ProductionDBusGateReviewCounts  `json:"counts"`
	RuntimeOwned                     bool                            `json:"runtime_owned"`
	GoRuntimeBacked                  bool                            `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool                            `json:"kde_policy_owner"`
	SystemServiceStarted             bool                            `json:"system_service_started"`
	SessionBusClaimed                bool                            `json:"session_bus_claimed"`
	ProductionBusClaimed             bool                            `json:"production_bus_claimed"`
	ProductionOwnerEnabled           bool                            `json:"production_owner_enabled"`
	ProductionActivationReady        bool                            `json:"production_activation_ready"`
	WriteMethodsEnabled              bool                            `json:"write_methods_enabled"`
	RuntimeWritesEnabled             bool                            `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled         bool                            `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled             bool                            `json:"backend_launch_enabled"`
	BackendProcessStarted            bool                            `json:"backend_process_started"`
	SupportBundleExported            bool                            `json:"support_bundle_exported"`
	SupportCaseCreated               bool                            `json:"support_case_created"`
	NotificationSent                 bool                            `json:"notification_sent"`
	NetworkRequired                  bool                            `json:"network_required"`
	HostRootModified                 bool                            `json:"host_root_modified"`
	PrivilegedContainerRequired      bool                            `json:"privileged_container_required"`
	StateRootPathExposed             bool                            `json:"state_root_path_exposed"`
	RawCommandExposed                bool                            `json:"raw_command_exposed"`
	RawExecutableExposed             bool                            `json:"raw_executable_exposed"`
	BackendDetailsExposed            bool                            `json:"backend_details_exposed"`
	BlockedActions                   []string                        `json:"blocked_actions"`
	NextRequirements                 []string                        `json:"next_requirements"`
	DesktopSafeSummary               string                          `json:"desktop_safe_summary"`
}

type ProductionDBusGateReviewRoute struct {
	ID                          string `json:"id"`
	RuntimeMethod               string `json:"runtime_method"`
	RequestType                 string `json:"request_type"`
	AuditRequestType            string `json:"audit_request_type"`
	RouteDecision               string `json:"route_decision"`
	CurrentRouteStatus          string `json:"current_route_status"`
	SmokeCoverageReady          bool   `json:"smoke_coverage_ready"`
	OwnerLocalRouteReady        bool   `json:"owner_local_route_ready"`
	ProductionDBusMethodPresent bool   `json:"production_dbus_method_present"`
	ProductionDBusExposureReady bool   `json:"production_dbus_exposure_ready"`
	WriteMethodsEnabled         bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled        bool   `json:"runtime_writes_enabled"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	HostRootModified            bool   `json:"host_root_modified"`
	InternalDetailsExposed      bool   `json:"internal_details_exposed"`
	RecommendedNextRoute        string `json:"recommended_next_route"`
}

type ProductionDBusGateReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionDBusGateReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionDBusGateReviewPreview(root string) (ProductionDBusGateReviewPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return ProductionDBusGateReviewPreview{}, err
	}
	routes, err := productionDBusGateReviewRoutes(root)
	if err != nil {
		return ProductionDBusGateReviewPreview{}, err
	}
	humanAuthorization, err := NewProductionDBusHumanAuthorizationPreflightPreview(root)
	if err != nil {
		return ProductionDBusGateReviewPreview{}, err
	}
	preview := ProductionDBusGateReviewPreview{
		Version:                          version,
		SchemaVersion:                    "xnix.runtime.production_dbus_gate_review.v1",
		RequestType:                      "production-dbus-gate-review-preview",
		GateType:                         "owner-local-smoke-covered-production-dbus-gate-review",
		Source:                           "materialization-audit+restricted-owner-smoke-fanout-audit+redacted-adapter-profile-audit+production-dbus-human-authorization-preflight-preview",
		RouteCount:                       len(routes),
		SmokeCoveredRouteCount:           countProductionDBusGateSmokeCoveredRoutes(routes),
		GateDecision:                     "production-dbus-gate-review-blocked",
		GateDecisionReason:               "Production D-Bus ownership remains disabled until every owner-local route is smoke-covered and a separate human-authorized production gate exists.",
		CurrentGateStatus:                "owner-local-smoke-covered-production-dbus-disabled",
		ProductionReadiness:              false,
		HumanAuthorizationRequired:       true,
		HumanAuthorizationPreflightReady: humanAuthorization.PreflightReady,
		HumanAuthorizationGranted:        humanAuthorization.AuthorizationAccepted,
		AuthorizationReceiptAccepted:     humanAuthorization.AuthorizationReceiptPresent && humanAuthorization.AuthorizationAccepted,
		Routes:                           routes,
		RouteIDs:                         productionDBusGateReviewRouteIDs(routes),
		RequiredGates: []string{
			"human-authorization",
			"production-service-activation-preflight",
			"runtime-write-gate-review",
			"route-by-route-production-dbus-method-review",
			"desktop-side-effect-review",
			"rollback-and-diagnostics-review",
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
			"claim production D-Bus ownership from smoke coverage",
			"start the production Runtime service from this review packet",
			"enable Runtime writes or write methods from this review packet",
			"invoke adapters, launch backends, export support bundles, create support cases, or send notifications",
			"expose state-root paths, raw commands, raw executables, backend details, or internal adapter identifiers",
			"mutate host root, require network, or require privileged containers",
		},
		NextRequirements: []string{
			"Keep all owner-local routes smoke-covered and non-launching.",
			"Add a separate human-authorized production service activation gate before claiming any production bus name.",
			"Add a route-by-route production D-Bus method review before exposing owner-local routes externally.",
			"Keep Runtime writes, adapter invocation, backend launch, support side effects, and host mutation disabled until their own gates exist.",
		},
		DesktopSafeSummary: "The production D-Bus gate review inventories smoke-covered owner-local routes and keeps production ownership disabled; it starts no service, claims no bus, enables no writes, launches no backend, emits no desktop side effects, and mutates no host state.",
	}
	checks := productionDBusGateReviewChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = productionDBusGateReviewCheckIDs(checks)
	preview.Counts = countProductionDBusGateReviewChecks(checks)
	if preview.SmokeCoveredRouteCount == preview.RouteCount && preview.HumanAuthorizationPreflightReady && preview.Counts.Blocked == 0 {
		preview.GateDecision = "production-dbus-gate-review-ready"
		preview.GateDecisionReason = "All tracked owner-local routes are smoke-covered and the human authorization preflight shape exists, but production D-Bus ownership still requires a separate accepted authorization receipt and production gate."
	}
	if err := validateNoBackendTerms(preview, "production D-Bus gate review preview"); err != nil {
		return ProductionDBusGateReviewPreview{}, err
	}
	return preview, nil
}

func productionDBusGateReviewRoutes(root string) ([]ProductionDBusGateReviewRoute, error) {
	materialization, err := appidentity.NewKDETestLaunchMaterializationOwnerRouteAuditPreview(root)
	if err != nil {
		return nil, err
	}
	restrictedFanOut, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(root)
	if err != nil {
		return nil, err
	}
	redactedProfile, err := appidentity.NewBackendAdapterContractOwnerRouteAuditPreview(root)
	if err != nil {
		return nil, err
	}
	return []ProductionDBusGateReviewRoute{
		{
			ID:                          "materialization-fanout-owner-route",
			RuntimeMethod:               materialization.ProposedOwnerMethod,
			RequestType:                 materialization.SubjectRequestType,
			AuditRequestType:            materialization.RequestType,
			RouteDecision:               materialization.RouteDecision,
			CurrentRouteStatus:          materialization.CurrentRouteStatus,
			SmokeCoverageReady:          materialization.OwnerSmokeCoverageReady,
			OwnerLocalRouteReady:        materialization.OwnerLocalRouteCandidateReady,
			ProductionDBusMethodPresent: materialization.ProductionDBusMethodPresent,
			ProductionDBusExposureReady: materialization.ProductionDBusExposureReady,
			WriteMethodsEnabled:         materialization.WriteMethodsEnabled,
			RuntimeWritesEnabled:        materialization.RuntimeWritesEnabled,
			BackendLaunchEnabled:        materialization.BackendLaunchEnabled,
			HostRootModified:            materialization.HostRootModified,
			InternalDetailsExposed:      materialization.StateRootPathExposed || materialization.RawCommandExposed || materialization.RawExecutableExposed || materialization.BackendDetailsExposed,
			RecommendedNextRoute:        materialization.RecommendedNextRoute,
		},
		{
			ID:                          "restricted-smoke-fanout-owner-route",
			RuntimeMethod:               restrictedFanOut.ProposedOwnerMethod,
			RequestType:                 restrictedFanOut.SubjectRequestType,
			AuditRequestType:            restrictedFanOut.RequestType,
			RouteDecision:               restrictedFanOut.RouteDecision,
			CurrentRouteStatus:          restrictedFanOut.CurrentRouteStatus,
			SmokeCoverageReady:          restrictedFanOut.OwnerSmokeCoverageReady,
			OwnerLocalRouteReady:        restrictedFanOut.OwnerLocalRouteCandidateReady,
			ProductionDBusMethodPresent: restrictedFanOut.ProductionDBusMethodPresent,
			ProductionDBusExposureReady: restrictedFanOut.ProductionDBusExposureReady,
			WriteMethodsEnabled:         restrictedFanOut.WriteMethodsEnabled,
			RuntimeWritesEnabled:        restrictedFanOut.RuntimeWritesEnabled,
			BackendLaunchEnabled:        restrictedFanOut.BackendLaunchEnabled,
			HostRootModified:            restrictedFanOut.HostRootModified,
			InternalDetailsExposed:      restrictedFanOut.StateRootPathExposed || restrictedFanOut.BackendDetailsExposed,
			RecommendedNextRoute:        restrictedFanOut.RecommendedNextRoute,
		},
		{
			ID:                          "redacted-adapter-profile-owner-route",
			RuntimeMethod:               redactedProfile.ProposedOwnerMethod,
			RequestType:                 redactedProfile.SubjectRequestType,
			AuditRequestType:            redactedProfile.RequestType,
			RouteDecision:               redactedProfile.RouteDecision,
			CurrentRouteStatus:          redactedProfile.CurrentRouteStatus,
			SmokeCoverageReady:          redactedProfile.OwnerSmokeCoverageReady,
			OwnerLocalRouteReady:        redactedProfile.OwnerLocalRouteCandidateReady,
			ProductionDBusMethodPresent: redactedProfile.ProductionDBusMethodPresent,
			ProductionDBusExposureReady: redactedProfile.ProductionDBusExposureReady,
			WriteMethodsEnabled:         redactedProfile.WriteMethodsEnabled,
			RuntimeWritesEnabled:        redactedProfile.RuntimeWritesEnabled,
			BackendLaunchEnabled:        redactedProfile.BackendLaunchEnabled,
			HostRootModified:            redactedProfile.HostRootModified,
			InternalDetailsExposed:      redactedProfile.StateRootPathExposed || redactedProfile.RawCommandExposed || redactedProfile.RawExecutableExposed || redactedProfile.BackendDetailsExposed,
			RecommendedNextRoute:        redactedProfile.RecommendedNextRoute,
		},
	}, nil
}

func productionDBusGateReviewChecks(preview ProductionDBusGateReviewPreview) []ProductionDBusGateReviewCheck {
	allRoutesSmokeCovered := preview.RouteCount == 3 && preview.SmokeCoveredRouteCount == preview.RouteCount
	return []ProductionDBusGateReviewCheck{
		productionDBusGateReviewCheck("tracked-routes-present", productionDBusGateReviewPassBlocked(preview.RouteCount == 3 && productionDBusGateReviewSameStrings(preview.RouteIDs, []string{"materialization-fanout-owner-route", "restricted-smoke-fanout-owner-route", "redacted-adapter-profile-owner-route"})), "The review packet inventories the three smoke-covered owner-local routes."),
		productionDBusGateReviewCheck("smoke-coverage-present", productionDBusGateReviewPendingUnless(allRoutesSmokeCovered), "Every tracked owner-local route must already be covered by restricted owner smoke evidence."),
		productionDBusGateReviewCheck("production-dbus-disabled", productionDBusGateReviewPassBlocked(!preview.ProductionReadiness && !preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionOwnerEnabled && !preview.ProductionActivationReady), "The packet keeps production D-Bus ownership and service activation disabled."),
		productionDBusGateReviewCheck("route-production-exposure-disabled", productionDBusGateReviewPassBlocked(productionDBusGateReviewRoutesKeepProductionDisabled(preview.Routes)), "No tracked route is exposed as a production D-Bus method or production D-Bus route."),
		productionDBusGateReviewCheck("writes-and-launch-disabled", productionDBusGateReviewPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && productionDBusGateReviewRoutesKeepWritesAndLaunchDisabled(preview.Routes)), "The packet keeps writes, adapter invocation, and backend launch disabled."),
		productionDBusGateReviewCheck("desktop-side-effects-disabled", productionDBusGateReviewPassBlocked(!preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.NotificationSent), "The packet does not export support bundles, create support cases, or send notifications."),
		productionDBusGateReviewCheck("host-boundary-closed", productionDBusGateReviewPassBlocked(!preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed && productionDBusGateReviewRoutesHideInternals(preview.Routes)), "The packet does not require network, privileged containers, host mutation, or internal detail exposure."),
		productionDBusGateReviewCheck("human-authorization-preflight-present", productionDBusGateReviewPassBlocked(preview.HumanAuthorizationPreflightReady && preview.HumanAuthorizationRequired), "The gate consumes a read-only human authorization preflight shape."),
		productionDBusGateReviewCheck("human-authorization-required", productionDBusGateReviewPassBlocked(preview.HumanAuthorizationRequired && !preview.HumanAuthorizationGranted && !preview.AuthorizationReceiptAccepted && !preview.ProductionReadiness), "Human authorization remains required before any production ownership claim."),
	}
}

func productionDBusGateReviewCheck(id string, status string, summary string) ProductionDBusGateReviewCheck {
	return ProductionDBusGateReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionDBusGateReviewPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionDBusGateReviewPendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func productionDBusGateReviewCheckIDs(checks []ProductionDBusGateReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionDBusGateReviewChecks(checks []ProductionDBusGateReviewCheck) ProductionDBusGateReviewCounts {
	counts := ProductionDBusGateReviewCounts{Total: len(checks)}
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

func countProductionDBusGateSmokeCoveredRoutes(routes []ProductionDBusGateReviewRoute) int {
	count := 0
	for _, route := range routes {
		if route.SmokeCoverageReady && route.OwnerLocalRouteReady {
			count++
		}
	}
	return count
}

func productionDBusGateReviewRouteIDs(routes []ProductionDBusGateReviewRoute) []string {
	ids := make([]string, 0, len(routes))
	for _, route := range routes {
		ids = append(ids, route.ID)
	}
	return ids
}

func productionDBusGateReviewRoutesKeepProductionDisabled(routes []ProductionDBusGateReviewRoute) bool {
	for _, route := range routes {
		if route.ProductionDBusMethodPresent || route.ProductionDBusExposureReady {
			return false
		}
	}
	return true
}

func productionDBusGateReviewRoutesKeepWritesAndLaunchDisabled(routes []ProductionDBusGateReviewRoute) bool {
	for _, route := range routes {
		if route.WriteMethodsEnabled || route.RuntimeWritesEnabled || route.BackendLaunchEnabled || route.HostRootModified {
			return false
		}
	}
	return true
}

func productionDBusGateReviewRoutesHideInternals(routes []ProductionDBusGateReviewRoute) bool {
	for _, route := range routes {
		if route.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionDBusGateReviewSameStrings(left []string, right []string) bool {
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

func readProductionDBusGateReviewVersion(root string) (string, error) {
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
