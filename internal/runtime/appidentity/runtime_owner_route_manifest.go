package appidentity

import "strings"

type RuntimeOwnerRouteManifestPreview struct {
	Version                     string                           `json:"version"`
	SchemaVersion               string                           `json:"schema_version"`
	RequestType                 string                           `json:"request_type"`
	ManifestType                string                           `json:"manifest_type"`
	Source                      string                           `json:"source"`
	RuntimeMethod               string                           `json:"runtime_method"`
	ReadMethod                  string                           `json:"read_method"`
	BusName                     string                           `json:"bus_name"`
	ObjectPath                  string                           `json:"object_path"`
	Interface                   string                           `json:"interface"`
	Routes                      []RuntimeOwnerRoute              `json:"routes"`
	RouteMethodNames            []string                         `json:"route_method_names"`
	RouteCounts                 RuntimeOwnerRouteCounts          `json:"route_counts"`
	Checks                      []RuntimeOwnerRouteManifestCheck `json:"checks"`
	CheckIDs                    []string                         `json:"check_ids"`
	Counts                      RuntimeOwnerRouteManifestCounts  `json:"counts"`
	MethodParityReady           bool                             `json:"method_parity_ready"`
	GoOwnerRouteCoverageReady   bool                             `json:"go_owner_route_coverage_ready"`
	CCoreAdapterRequired        bool                             `json:"c_core_adapter_required"`
	LegacyRuntimeRoutesPresent  bool                             `json:"legacy_runtime_routes_present"`
	ProductionOwnerRoutesReady  bool                             `json:"production_owner_routes_ready"`
	RuntimeOwned                bool                             `json:"runtime_owned"`
	GoRuntimeBacked             bool                             `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                             `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool                             `json:"kde_may_claim_runtime_ownership"`
	SystemServiceStarted        bool                             `json:"system_service_started"`
	ProductionBusClaimed        bool                             `json:"production_bus_claimed"`
	WriteMethodsEnabled         bool                             `json:"write_methods_enabled"`
	NetworkRequired             bool                             `json:"network_required"`
	HostRootModified            bool                             `json:"host_root_modified"`
	PrivilegedContainerRequired bool                             `json:"privileged_container_required"`
	BackendDetailsExposed       bool                             `json:"backend_details_exposed"`
	BlockedActions              []string                         `json:"blocked_actions"`
	NextRequirements            []string                         `json:"next_requirements"`
	DesktopSafeSummary          string                           `json:"desktop_safe_summary"`
}

type RuntimeOwnerRoute struct {
	Method                 string `json:"method"`
	CurrentSource          string `json:"current_source"`
	TargetSource           string `json:"target_source"`
	GoCommand              string `json:"go_command,omitempty"`
	CCoreCommand           string `json:"c_core_command,omitempty"`
	LegacyDispatchMethod   string `json:"legacy_dispatch_method,omitempty"`
	RouteStatus            string `json:"route_status"`
	GoRouteReady           bool   `json:"go_route_ready"`
	CCoreBacked            bool   `json:"c_core_backed"`
	LegacyDispatchRequired bool   `json:"legacy_dispatch_required"`
	KDEVisibleReadOnly     bool   `json:"kde_visible_read_only"`
}

type RuntimeOwnerRouteCounts struct {
	Total       int `json:"total"`
	GoRouted    int `json:"go_routed"`
	CCoreBacked int `json:"c_core_backed"`
	RubyLegacy  int `json:"ruby_legacy"`
	Ready       int `json:"ready"`
	Pending     int `json:"pending"`
	Blocked     int `json:"blocked"`
}

type RuntimeOwnerRouteManifestCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeOwnerRouteManifestCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

var runtimeOwnerRouteGoCommands = map[string]string{
	"ListApplications":                       "applications-preview",
	"GetApplication":                         "application-preview",
	"GetEngineCatalog":                       "engine-catalog-preview",
	"GetRunPlan":                             "run-plan-preview",
	"GetCompatibilityPackageSource":          "package-source-preview",
	"GetCompatibilityAcquisitionPreflight":   "acquisition-preflight-preview",
	"GetCompatibilityArtifactManifest":       "artifact-manifest-preview",
	"GetDiagnostics":                         "diagnostics-preview",
	"GetDesktopActivationTransactionPreview": "desktop-activation-transaction-preview",
	"GetDesktopActivationStatus":             "desktop-activation-status-preview",
	"GetDesktopEntryPlan":                    "desktop-entry-preview",
	"GetDesktopIconPlan":                     "desktop-icon-preview",
	"GetDesktopResourceBridgePlan":           "desktop-resource-bridge-preview",
	"GetFileAssociationPlan":                 "mimeapps-preview",
	"GetNotificationPlan":                    "notification-preview",
	"GetTrayStatus":                          "tray-status-preview",
	"GetKRunnerQueryPlan":                    "krunner-query-preview",
	"GetPortalRequestPlan":                   "portal-request-preview",
	"GetBackendBinding":                      "backend-binding-preview",
	"GetBackendSelectionPlan":                "backend-selection-preview",
	"GetBackendEnvironmentPlan":              "backend-environment-preview",
	"GetExecutionReadiness":                  "execution-readiness-preview",
	"GetLaunchIntent":                        "launch-intent-preview",
	"GetRuntimeServiceBinding":               "runtime-service-binding-preview",
	"GetRuntimeLiveOwnerGate":                "runtime-live-owner-gate-preview",
	"GetRuntimeOwnerSmokePlan":               "runtime-owner-smoke-plan-preview",
	"GetRuntimeMethodParityManifest":         "runtime-method-parity-manifest-preview",
	"GetCompatibilitySettings":               "settings-preview",
	"GetCompatibilitySettingsChangePlan":     "settings-change-preview",
	"GetCompatibilityModeSwitchPlan":         "mode-switch-preview",
	"GetCompatibilityPermissionReviewPlan":   "permission-review-preview",
	"GetCompatibilityReviewFlowPlan":         "review-flow-preview",
	"GetCompatibilityActionQueue":            "kde-action-queue-preview",
	"GetCompatibilityActionReviewReceipt":    "kde-action-receipt-preview",
	"GetCompatibilityCenterSummary":          "compatibility-center-preview",
	"GetKDECenterPage":                       "kde-center-page-preview",
	"GetKDECenterPageSections":               "kde-center-page-sections-preview",
	"GetKDECenterPageSectionDetail":          "kde-center-page-section-detail-preview",
}

var runtimeOwnerRouteCCoreCommands = map[string]string{
	"GetDesktopActivationManifest":  "desktop-activation-manifest",
	"GetTaskManagerIdentityPlan":    "task-manager-identity-plan",
	"GetKDEIntegrationStatus":       "kde-integration-status",
	"GetKDEShellIntegrationPlan":    "kde-shell-integration-plan",
	"GetKDEApplicationSurfacePlan":  "kde-application-surface-plan",
	"GetKWinWindowRulePlan":         "kwin-window-rule-plan",
	"GetApplicationStateRoot":       "state-root-policy",
	"GetCompatibilityInstallPlan":   "compatibility-install-plan",
	"GetBackendCapabilityMatrix":    "backend-capability-matrix",
	"GetBackendLifecycle":           "backend-lifecycle",
	"GetRepairPlan":                 "compatibility-repair-plan",
	"GetTestPlan":                   "compatibility-test-plan",
	"GetTestResult":                 "compatibility-test-result",
	"GetAIDiagnosticInput":          "ai-diagnostic-input",
	"GetAIDiagnosticRecommendation": "ai-diagnostic-recommendation",
	"GetAIRepairApprovalGate":       "ai-repair-approval-gate",
	"GetSnapshotPlan":               "compatibility-snapshot-plan",
	"GetPortalAccessPolicy":         "portal-policy",
	"GetRuntimeWriteGate":           "write-gate",
}

func NewRuntimeOwnerRouteManifestPreview(root string) (RuntimeOwnerRouteManifestPreview, error) {
	methodParity, err := NewRuntimeMethodParityManifestPreview(root)
	if err != nil {
		return RuntimeOwnerRouteManifestPreview{}, err
	}

	routeSources := runtimeOwnerRouteSources(root)
	routes := runtimeOwnerRoutes(routeSources)
	routeCounts := countRuntimeOwnerRoutes(routes)
	checks := runtimeOwnerRouteManifestChecks(methodParity, routeCounts)
	counts := countRuntimeOwnerRouteManifestChecks(checks)

	preview := RuntimeOwnerRouteManifestPreview{
		Version:                     methodParity.Version,
		SchemaVersion:               "xnix.runtime.owner_route_manifest.v1",
		RequestType:                 "runtime-owner-route-manifest-preview",
		ManifestType:                "runtime-owner-route-manifest",
		Source:                      "runtime-method-parity-manifest-preview+go-runtime-cli+c-runtime-core+runtime-dispatch",
		RuntimeMethod:               "GetRuntimeOwnerRouteManifest",
		ReadMethod:                  "GetRuntimeOwnerRouteManifestPreview",
		BusName:                     methodParity.BusName,
		ObjectPath:                  methodParity.ObjectPath,
		Interface:                   methodParity.Interface,
		Routes:                      routes,
		RouteMethodNames:            runtimeOwnerRouteMethodNames(routes),
		RouteCounts:                 routeCounts,
		Checks:                      checks,
		CheckIDs:                    runtimeOwnerRouteManifestCheckIDs(checks),
		Counts:                      counts,
		MethodParityReady:           methodParity.ReadOnlyMethodParityReady,
		GoOwnerRouteCoverageReady:   routeCounts.Total > 0 && routeCounts.GoRouted == routeCounts.Total,
		CCoreAdapterRequired:        routeCounts.CCoreBacked > 0,
		LegacyRuntimeRoutesPresent:  routeCounts.RubyLegacy > 0,
		ProductionOwnerRoutesReady:  false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		SystemServiceStarted:        false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions:              runtimeOwnerRouteManifestBlockedActions(),
		NextRequirements:            runtimeOwnerRouteManifestNextRequirements(),
		DesktopSafeSummary:          runtimeOwnerRouteManifestSummary(routeCounts),
	}
	if err := validateNoBackendTerms(preview, "Runtime owner route manifest preview"); err != nil {
		return RuntimeOwnerRouteManifestPreview{}, err
	}
	return preview, nil
}

type runtimeOwnerRouteSourceSet struct {
	GoCLI           string
	CCoreDispatch   string
	RuntimeDispatch string
}

func runtimeOwnerRouteSources(root string) runtimeOwnerRouteSourceSet {
	return runtimeOwnerRouteSourceSet{
		GoCLI: readRuntimeMethodParitySources(root, []string{
			"cmd/xnix-runtime-go/main.go",
			"cmd/xnix-runtime-go/runtime_owner_commands.go",
		}),
		CCoreDispatch:   readRuntimeMethodParitySources(root, []string{"runtime/core/xnix_runtime_core_cli_dispatch.inc"}),
		RuntimeDispatch: readRuntimeMethodParitySources(root, []string{"lib/xnix/compatibility/runtime_daemon.rb"}),
	}
}

func runtimeOwnerRoutes(sources runtimeOwnerRouteSourceSet) []RuntimeOwnerRoute {
	routes := make([]RuntimeOwnerRoute, 0, len(runtimeMethodParityReadOnlyMethods))
	for _, method := range runtimeMethodParityReadOnlyMethods {
		routes = append(routes, runtimeOwnerRoute(method, sources))
	}
	return routes
}

func runtimeOwnerRoute(method string, sources runtimeOwnerRouteSourceSet) RuntimeOwnerRoute {
	if command, ok := runtimeOwnerRouteGoCommands[method]; ok {
		ready := strings.Contains(sources.GoCLI, "\""+command+"\"")
		status := runtimeOwnerRouteStatus(ready, "go-preview-ready", "go-preview-missing")
		return RuntimeOwnerRoute{
			Method:             method,
			CurrentSource:      "go-runtime-cli",
			TargetSource:       "go-owner-native",
			GoCommand:          command,
			RouteStatus:        status,
			GoRouteReady:       ready,
			KDEVisibleReadOnly: true,
		}
	}
	if command, ok := runtimeOwnerRouteCCoreCommands[method]; ok {
		ready := strings.Contains(sources.CCoreDispatch, "\""+command+"\"")
		status := runtimeOwnerRouteStatus(ready, "c-adapter-pending", "c-adapter-blocked")
		return RuntimeOwnerRoute{
			Method:             method,
			CurrentSource:      "c-runtime-core",
			TargetSource:       "go-owner-c-adapter",
			CCoreCommand:       command,
			RouteStatus:        status,
			CCoreBacked:        ready,
			KDEVisibleReadOnly: true,
		}
	}

	ready := strings.Contains(sources.RuntimeDispatch, "\""+method+"\"")
	status := runtimeOwnerRouteStatus(ready, "legacy-dispatch-pending", "legacy-dispatch-blocked")
	return RuntimeOwnerRoute{
		Method:                 method,
		CurrentSource:          "ruby-runtime-dispatch",
		TargetSource:           "go-owner-native",
		LegacyDispatchMethod:   method,
		RouteStatus:            status,
		LegacyDispatchRequired: ready,
		KDEVisibleReadOnly:     true,
	}
}

func runtimeOwnerRouteStatus(ready bool, readyStatus string, blockedStatus string) string {
	if ready {
		return readyStatus
	}
	return blockedStatus
}

func countRuntimeOwnerRoutes(routes []RuntimeOwnerRoute) RuntimeOwnerRouteCounts {
	counts := RuntimeOwnerRouteCounts{Total: len(routes)}
	for _, route := range routes {
		switch {
		case route.GoRouteReady:
			counts.GoRouted++
			counts.Ready++
		case route.CCoreBacked:
			counts.CCoreBacked++
			counts.Pending++
		case route.LegacyDispatchRequired:
			counts.RubyLegacy++
			counts.Pending++
		default:
			counts.Blocked++
		}
	}
	return counts
}

func runtimeOwnerRouteManifestChecks(methodParity RuntimeMethodParityManifestPreview, routeCounts RuntimeOwnerRouteCounts) []RuntimeOwnerRouteManifestCheck {
	return []RuntimeOwnerRouteManifestCheck{
		runtimeOwnerRouteManifestCheck("method-parity", runtimeOwnerRouteManifestPassBlocked(methodParity.ReadOnlyMethodParityReady), "Owner routes can only advance after read-only method parity is complete."),
		runtimeOwnerRouteManifestCheck("go-route-coverage", runtimeOwnerRouteManifestPendingUnless(routeCounts.GoRouted == routeCounts.Total), "Every read-only method should become a native Go owner route."),
		runtimeOwnerRouteManifestCheck("c-core-adapter-boundary", runtimeOwnerRouteManifestPendingUnless(routeCounts.CCoreBacked == 0), "C Runtime policy routes still need a Go owner adapter boundary."),
		runtimeOwnerRouteManifestCheck("ruby-legacy-dispatch", runtimeOwnerRouteManifestPendingUnless(routeCounts.RubyLegacy == 0), "Legacy Runtime dispatch routes must be replaced before production Go ownership."),
		runtimeOwnerRouteManifestCheck("write-route-gate", "pass", "Write methods must stay disabled until production owner readiness is proven."),
		runtimeOwnerRouteManifestCheck("host-safety-boundary", "pass", "Route manifest preview must not start services, claim bus names, require network, or mutate the host root."),
	}
}

func runtimeOwnerRouteManifestCheck(id string, status string, summary string) RuntimeOwnerRouteManifestCheck {
	return RuntimeOwnerRouteManifestCheck{ID: id, Status: status, Summary: summary}
}

func runtimeOwnerRouteManifestPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeOwnerRouteManifestPendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func countRuntimeOwnerRouteManifestChecks(checks []RuntimeOwnerRouteManifestCheck) RuntimeOwnerRouteManifestCounts {
	counts := RuntimeOwnerRouteManifestCounts{Total: len(checks)}
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

func runtimeOwnerRouteManifestCheckIDs(checks []RuntimeOwnerRouteManifestCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func runtimeOwnerRouteMethodNames(routes []RuntimeOwnerRoute) []string {
	methods := make([]string, 0, len(routes))
	for _, route := range routes {
		methods = append(methods, route.Method)
	}
	return methods
}

func runtimeOwnerRouteManifestBlockedActions() []string {
	return []string{
		"start production Runtime owner from route manifest preview",
		"claim production D-Bus name from route manifest preview",
		"serve C-backed read-only routes without Go owner adapters",
		"enable write methods before all read-only routes are owner-ready",
		"let KDE route Runtime policy directly",
		"mutate host root during route manifest preview",
	}
}

func runtimeOwnerRouteManifestNextRequirements() []string {
	return []string{
		"Complete native Go owner handlers or owner adapters for every non-Go read-only route.",
		"Add a Go owner adapter boundary for C Runtime policy routes.",
		"Serve the route table from a restricted owner smoke before production bus ownership.",
		"Keep KDE consumers on read-only Runtime methods while route migration continues.",
	}
}

func runtimeOwnerRouteManifestSummary(routeCounts RuntimeOwnerRouteCounts) string {
	if routeCounts.Blocked > 0 {
		return "Runtime owner routes are blocked by missing read-only route evidence."
	}
	if routeCounts.RubyLegacy > 0 {
		return "Runtime owner routes have Go coverage for current Go previews, with C adapter and legacy dispatch migration still pending."
	}
	if routeCounts.CCoreBacked > 0 {
		return "Runtime owner routes have Go coverage for migrated application catalog, engine catalog, run planning, package acquisition, diagnostics, and current Go previews, with C adapter migration still pending."
	}
	return "Runtime owner routes are fully native to the Go owner preview, but production bus ownership remains gated."
}
