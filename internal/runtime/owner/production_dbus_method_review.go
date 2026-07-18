package owner

import (
	"fmt"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type ProductionDBusMethodReviewPreview struct {
	Version                         string                            `json:"version"`
	SchemaVersion                   string                            `json:"schema_version"`
	RequestType                     string                            `json:"request_type"`
	ReviewType                      string                            `json:"review_type"`
	Source                          string                            `json:"source"`
	ReviewDecision                  string                            `json:"review_decision"`
	ReadOnlyContractMethodCount     int                               `json:"read_only_contract_method_count"`
	OwnerLocalCandidateCount        int                               `json:"owner_local_candidate_count"`
	WriteMethodCount                int                               `json:"write_method_count"`
	ReviewedMethodCount             int                               `json:"reviewed_method_count"`
	ProductionExposureReadyCount    int                               `json:"production_exposure_ready_count"`
	NewProductionMethodRequestCount int                               `json:"new_production_method_request_count"`
	ReadOnlyMethods                 []ProductionDBusMethodReviewRoute `json:"read_only_methods"`
	OwnerLocalCandidates            []ProductionDBusMethodReviewRoute `json:"owner_local_candidates"`
	WriteMethods                    []ProductionDBusMethodReviewRoute `json:"write_methods"`
	RequiredGates                   []string                          `json:"required_gates"`
	Checks                          []ProductionDBusMethodReviewCheck `json:"checks"`
	CheckIDs                        []string                          `json:"check_ids"`
	Counts                          ProductionDBusMethodReviewCounts  `json:"counts"`
	RuntimeOwned                    bool                              `json:"runtime_owned"`
	GoRuntimeBacked                 bool                              `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool                              `json:"kde_policy_owner"`
	SystemServiceStarted            bool                              `json:"system_service_started"`
	SessionBusClaimed               bool                              `json:"session_bus_claimed"`
	ProductionBusClaimed            bool                              `json:"production_bus_claimed"`
	ProductionOwnerEnabled          bool                              `json:"production_owner_enabled"`
	ProductionActivationReady       bool                              `json:"production_activation_ready"`
	WriteMethodsEnabled             bool                              `json:"write_methods_enabled"`
	RuntimeWritesEnabled            bool                              `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled        bool                              `json:"adapter_invocation_enabled"`
	BackendLaunchEnabled            bool                              `json:"backend_launch_enabled"`
	BackendProcessStarted           bool                              `json:"backend_process_started"`
	NotificationSent                bool                              `json:"notification_sent"`
	NetworkRequired                 bool                              `json:"network_required"`
	HostRootModified                bool                              `json:"host_root_modified"`
	PrivilegedContainerRequired     bool                              `json:"privileged_container_required"`
	StateRootPathExposed            bool                              `json:"state_root_path_exposed"`
	RawCommandExposed               bool                              `json:"raw_command_exposed"`
	RawExecutableExposed            bool                              `json:"raw_executable_exposed"`
	BackendDetailsExposed           bool                              `json:"backend_details_exposed"`
	BlockedActions                  []string                          `json:"blocked_actions"`
	NextRequirements                []string                          `json:"next_requirements"`
	DesktopSafeSummary              string                            `json:"desktop_safe_summary"`
}

type ProductionDBusMethodReviewRoute struct {
	ID                           string `json:"id"`
	Method                       string `json:"method"`
	RouteClass                   string `json:"route_class"`
	CurrentExposure              string `json:"current_exposure"`
	FutureExposureDecision       string `json:"future_exposure_decision"`
	ContractMethodPresent        bool   `json:"contract_method_present"`
	OwnerLocalCandidate          bool   `json:"owner_local_candidate"`
	GoRouteReady                 bool   `json:"go_route_ready"`
	SmokeCoverageReady           bool   `json:"smoke_coverage_ready"`
	ReadOnly                     bool   `json:"read_only"`
	WriteMethod                  bool   `json:"write_method"`
	NewProductionMethodRequested bool   `json:"new_production_method_requested"`
	ProductionExposureReady      bool   `json:"production_exposure_ready"`
	ProductionOwnerEnabled       bool   `json:"production_owner_enabled"`
	WriteMethodsEnabled          bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled         bool   `json:"runtime_writes_enabled"`
	BackendLaunchEnabled         bool   `json:"backend_launch_enabled"`
	HostRootModified             bool   `json:"host_root_modified"`
	InternalDetailsExposed       bool   `json:"internal_details_exposed"`
	ReviewStatus                 string `json:"review_status"`
}

type ProductionDBusMethodReviewCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type ProductionDBusMethodReviewCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewProductionDBusMethodReviewPreview(root string) (ProductionDBusMethodReviewPreview, error) {
	routeManifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return ProductionDBusMethodReviewPreview{}, err
	}
	methodParity, err := appidentity.NewRuntimeMethodParityManifestPreview(root)
	if err != nil {
		return ProductionDBusMethodReviewPreview{}, err
	}
	gateReview, err := NewProductionDBusGateReviewPreview(root)
	if err != nil {
		return ProductionDBusMethodReviewPreview{}, err
	}
	writeGate, err := appidentity.NewRuntimeWriteGatePreview(root, "Launch")
	if err != nil {
		return ProductionDBusMethodReviewPreview{}, err
	}

	readOnlyMethods := productionDBusMethodReviewReadOnlyRoutes(routeManifest.Routes)
	ownerLocalCandidates := productionDBusMethodReviewOwnerLocalRoutes(gateReview.Routes)
	writeMethods := productionDBusMethodReviewWriteRoutes(methodParity.WriteMethods)
	reviewedMethodCount := len(readOnlyMethods) + len(ownerLocalCandidates) + len(writeMethods)

	preview := ProductionDBusMethodReviewPreview{
		Version:                         routeManifest.Version,
		SchemaVersion:                   "xnix.runtime.production_dbus_method_review.v1",
		RequestType:                     "production-dbus-method-review-preview",
		ReviewType:                      "route-by-route-production-dbus-method-review",
		Source:                          "runtime-method-parity-manifest-preview+runtime-owner-route-manifest-preview+production-dbus-gate-review-preview+runtime-write-gate-preview+production-human-authorization-receipt-consolidation-preview",
		ReviewDecision:                  "production-dbus-method-review-blocked",
		ReadOnlyContractMethodCount:     len(readOnlyMethods),
		OwnerLocalCandidateCount:        len(ownerLocalCandidates),
		WriteMethodCount:                len(writeMethods),
		ReviewedMethodCount:             reviewedMethodCount,
		ProductionExposureReadyCount:    0,
		NewProductionMethodRequestCount: 0,
		ReadOnlyMethods:                 readOnlyMethods,
		OwnerLocalCandidates:            ownerLocalCandidates,
		WriteMethods:                    writeMethods,
		RequiredGates: []string{
			"production-dbus-gate-review",
			"production-service-activation-preflight",
			"runtime-write-gate-review",
			"human-authorization-receipt",
			"production-human-authorization-receipt-consolidation",
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
		NotificationSent:            false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"register new production D-Bus methods from this review",
			"claim the production bus name from this review",
			"promote owner-local candidates to production D-Bus exposure",
			"enable write-method dispatch, Runtime writes, request creation, or execution",
			"start services, invoke adapters, launch backends, send notifications, require network, or mutate the host root",
			"expose state-root paths, raw commands, raw executables, backend details, or internal adapter identifiers",
		},
		NextRequirements: []string{
			"Keep read-only contract methods routed through the Go owner and production-owner-disabled until the production service gate is authorized.",
			"Keep owner-local candidates owner-local until the consolidated opaque authorization receipt boundary is accepted by a separate operator action and rollback diagnostics review exists.",
			"Keep write methods disabled until production ownership, service activation, recipe trust, Portal, snapshot, and user review gates pass.",
			"Run the full product-image gate only after explicit operator authorization.",
		},
		DesktopSafeSummary: "The production D-Bus method review inventories read-only contract methods, owner-local candidates, and reserved write methods while registering no new production methods, claiming no bus, enabling no writes, launching no backend, and mutating no host state.",
	}
	checks := productionDBusMethodReviewChecks(preview, routeManifest, methodParity, gateReview, writeGate)
	preview.Checks = checks
	preview.CheckIDs = productionDBusMethodReviewCheckIDs(checks)
	preview.Counts = countProductionDBusMethodReviewChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.ReviewDecision = "production-dbus-method-review-ready-production-exposure-disabled"
	}
	if err := validateNoBackendTerms(preview, "production D-Bus method review preview"); err != nil {
		return ProductionDBusMethodReviewPreview{}, err
	}
	return preview, nil
}

func productionDBusMethodReviewReadOnlyRoutes(routes []appidentity.RuntimeOwnerRoute) []ProductionDBusMethodReviewRoute {
	reviewRoutes := make([]ProductionDBusMethodReviewRoute, 0, len(routes))
	for index, route := range routes {
		reviewRoutes = append(reviewRoutes, ProductionDBusMethodReviewRoute{
			ID:                           fmt.Sprintf("read-only-contract-%02d", index+1),
			Method:                       route.Method,
			RouteClass:                   "dbus-read-only-contract",
			CurrentExposure:              "read-only-contract-method-production-owner-disabled",
			FutureExposureDecision:       "reviewed-read-only-contract-production-owner-disabled",
			ContractMethodPresent:        true,
			OwnerLocalCandidate:          false,
			GoRouteReady:                 route.GoRouteReady,
			SmokeCoverageReady:           route.GoRouteReady,
			ReadOnly:                     true,
			WriteMethod:                  false,
			NewProductionMethodRequested: false,
			ProductionExposureReady:      false,
			ProductionOwnerEnabled:       false,
			WriteMethodsEnabled:          false,
			RuntimeWritesEnabled:         false,
			BackendLaunchEnabled:         false,
			HostRootModified:             false,
			InternalDetailsExposed:       false,
			ReviewStatus:                 productionDBusMethodReviewRouteStatus(route.GoRouteReady),
		})
	}
	return reviewRoutes
}

func productionDBusMethodReviewOwnerLocalRoutes(routes []ProductionDBusGateReviewRoute) []ProductionDBusMethodReviewRoute {
	reviewRoutes := make([]ProductionDBusMethodReviewRoute, 0, len(routes))
	for _, route := range routes {
		reviewRoutes = append(reviewRoutes, ProductionDBusMethodReviewRoute{
			ID:                           route.ID,
			Method:                       route.RuntimeMethod,
			RouteClass:                   "owner-local-candidate",
			CurrentExposure:              "owner-local-only",
			FutureExposureDecision:       "owner-local-only-no-production-dbus-method",
			ContractMethodPresent:        false,
			OwnerLocalCandidate:          true,
			GoRouteReady:                 route.OwnerLocalRouteReady,
			SmokeCoverageReady:           route.SmokeCoverageReady,
			ReadOnly:                     true,
			WriteMethod:                  false,
			NewProductionMethodRequested: false,
			ProductionExposureReady:      false,
			ProductionOwnerEnabled:       false,
			WriteMethodsEnabled:          route.WriteMethodsEnabled,
			RuntimeWritesEnabled:         route.RuntimeWritesEnabled,
			BackendLaunchEnabled:         route.BackendLaunchEnabled,
			HostRootModified:             route.HostRootModified,
			InternalDetailsExposed:       route.InternalDetailsExposed,
			ReviewStatus:                 productionDBusMethodReviewRouteStatus(route.OwnerLocalRouteReady && route.SmokeCoverageReady),
		})
	}
	return reviewRoutes
}

func productionDBusMethodReviewWriteRoutes(methods []string) []ProductionDBusMethodReviewRoute {
	reviewRoutes := make([]ProductionDBusMethodReviewRoute, 0, len(methods))
	for index, method := range methods {
		reviewRoutes = append(reviewRoutes, ProductionDBusMethodReviewRoute{
			ID:                           fmt.Sprintf("reserved-write-method-%02d", index+1),
			Method:                       method,
			RouteClass:                   "reserved-write-method",
			CurrentExposure:              "write-method-disabled",
			FutureExposureDecision:       "write-method-disabled-no-production-dispatch",
			ContractMethodPresent:        true,
			OwnerLocalCandidate:          false,
			GoRouteReady:                 false,
			SmokeCoverageReady:           false,
			ReadOnly:                     false,
			WriteMethod:                  true,
			NewProductionMethodRequested: false,
			ProductionExposureReady:      false,
			ProductionOwnerEnabled:       false,
			WriteMethodsEnabled:          false,
			RuntimeWritesEnabled:         false,
			BackendLaunchEnabled:         false,
			HostRootModified:             false,
			InternalDetailsExposed:       false,
			ReviewStatus:                 "reviewed-disabled",
		})
	}
	return reviewRoutes
}

func productionDBusMethodReviewChecks(preview ProductionDBusMethodReviewPreview, routeManifest appidentity.RuntimeOwnerRouteManifestPreview, methodParity appidentity.RuntimeMethodParityManifestPreview, gateReview ProductionDBusGateReviewPreview, writeGate appidentity.RuntimeWriteGatePreview) []ProductionDBusMethodReviewCheck {
	return []ProductionDBusMethodReviewCheck{
		productionDBusMethodReviewCheck("read-only-contract-methods-reviewed", productionDBusMethodReviewPassBlocked(preview.ReadOnlyContractMethodCount == 61 && len(preview.ReadOnlyMethods) == 61 && productionDBusMethodReviewRoutesReady(preview.ReadOnlyMethods)), "Every read-only D-Bus contract method has a route review entry."),
		productionDBusMethodReviewCheck("owner-local-candidates-reviewed", productionDBusMethodReviewPassBlocked(preview.OwnerLocalCandidateCount == 3 && productionDBusMethodReviewRoutesReady(preview.OwnerLocalCandidates)), "The three owner-local candidate routes remain owner-local and smoke-covered."),
		productionDBusMethodReviewCheck("write-methods-reviewed-disabled", productionDBusMethodReviewPassBlocked(preview.WriteMethodCount == 4 && !methodParity.WriteMethodsSupported && !methodParity.WriteMethodDispatchEnabled && !writeGate.WriteMethodEnabled && !writeGate.DispatchEnabled), "Reserved write methods are reviewed and remain disabled."),
		productionDBusMethodReviewCheck("no-new-production-methods-requested", productionDBusMethodReviewPassBlocked(preview.NewProductionMethodRequestCount == 0 && !productionDBusMethodReviewAnyNewMethodRequested(preview)), "The review registers no new production D-Bus methods."),
		productionDBusMethodReviewCheck("production-exposure-disabled", productionDBusMethodReviewPassBlocked(preview.ProductionExposureReadyCount == 0 && !gateReview.ProductionReadiness && !preview.ProductionOwnerEnabled && !preview.ProductionBusClaimed), "Production D-Bus exposure remains disabled after route review."),
		productionDBusMethodReviewCheck("route-manifest-consumed", productionDBusMethodReviewPassBlocked(routeManifest.GoOwnerRouteCoverageReady && routeManifest.MethodParityReady), "The review consumes the Go owner route manifest and method parity evidence."),
		productionDBusMethodReviewCheck("unsafe-gates-closed", productionDBusMethodReviewPassBlocked(!preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NotificationSent && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed), "The method review keeps service, bus, write, backend, desktop side-effect, unsafe-data, and host gates closed."),
	}
}

func productionDBusMethodReviewCheck(id string, status string, summary string) ProductionDBusMethodReviewCheck {
	return ProductionDBusMethodReviewCheck{ID: id, Status: status, Summary: summary}
}

func productionDBusMethodReviewRouteStatus(ready bool) string {
	if ready {
		return "reviewed"
	}
	return "blocked"
}

func productionDBusMethodReviewPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func productionDBusMethodReviewRoutesReady(routes []ProductionDBusMethodReviewRoute) bool {
	if len(routes) == 0 {
		return false
	}
	for _, route := range routes {
		if route.ReviewStatus != "reviewed" && route.ReviewStatus != "reviewed-disabled" {
			return false
		}
		if route.NewProductionMethodRequested || route.ProductionExposureReady || route.ProductionOwnerEnabled || route.WriteMethodsEnabled || route.RuntimeWritesEnabled || route.BackendLaunchEnabled || route.HostRootModified || route.InternalDetailsExposed {
			return false
		}
	}
	return true
}

func productionDBusMethodReviewAnyNewMethodRequested(preview ProductionDBusMethodReviewPreview) bool {
	for _, routes := range [][]ProductionDBusMethodReviewRoute{preview.ReadOnlyMethods, preview.OwnerLocalCandidates, preview.WriteMethods} {
		for _, route := range routes {
			if route.NewProductionMethodRequested {
				return true
			}
		}
	}
	return false
}

func productionDBusMethodReviewCheckIDs(checks []ProductionDBusMethodReviewCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countProductionDBusMethodReviewChecks(checks []ProductionDBusMethodReviewCheck) ProductionDBusMethodReviewCounts {
	counts := ProductionDBusMethodReviewCounts{Total: len(checks)}
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
