package appidentity

type RuntimeRouteConvergencePreview struct {
	Version                     string                           `json:"version"`
	SchemaVersion               string                           `json:"schema_version"`
	RequestType                 string                           `json:"request_type"`
	ReportType                  string                           `json:"report_type"`
	Source                      string                           `json:"source"`
	RuntimeMethod               string                           `json:"runtime_method"`
	ReadMethod                  string                           `json:"read_method"`
	Routes                      []RuntimeRouteConvergenceRoute   `json:"routes"`
	RouteMethodNames            []string                         `json:"route_method_names"`
	ClassificationCounts        RuntimeRouteClassificationCounts `json:"classification_counts"`
	MigrationGroups             []RuntimeRouteMigrationGroup     `json:"migration_groups"`
	Checks                      []RuntimeRouteConvergenceCheck   `json:"checks"`
	CheckIDs                    []string                         `json:"check_ids"`
	AllRoutesClassified         bool                             `json:"all_routes_classified"`
	NativeGoCoverageReady       bool                             `json:"native_go_coverage_ready"`
	ProductionOwnerReady        bool                             `json:"production_owner_ready"`
	RuntimeOwned                bool                             `json:"runtime_owned"`
	GoRuntimeBacked             bool                             `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                             `json:"kde_policy_owner"`
	WriteMethodsEnabled         bool                             `json:"write_methods_enabled"`
	NetworkRequired             bool                             `json:"network_required"`
	HostRootModified            bool                             `json:"host_root_modified"`
	PrivilegedContainerRequired bool                             `json:"privileged_container_required"`
	BackendLaunchEnabled        bool                             `json:"backend_launch_enabled"`
	BackendDetailsExposed       bool                             `json:"backend_details_exposed"`
	UnclassifiedRoutes          []string                         `json:"unclassified_routes"`
	BlockedActions              []string                         `json:"blocked_actions"`
	NextMigrationOrder          []string                         `json:"next_migration_order"`
	DesktopSafeSummary          string                           `json:"desktop_safe_summary"`
}

type RuntimeRouteConvergenceRoute struct {
	Method                string   `json:"method"`
	Domain                string   `json:"domain"`
	Classification        string   `json:"classification"`
	CurrentSource         string   `json:"current_source"`
	TargetSource          string   `json:"target_source"`
	GoCommand             string   `json:"go_command,omitempty"`
	RouteStatus           string   `json:"route_status"`
	MigrationPriority     int      `json:"migration_priority"`
	MigrationRisk         string   `json:"migration_risk"`
	BlockedReasons        []string `json:"blocked_reasons"`
	KDEVisibleReadOnly    bool     `json:"kde_visible_read_only"`
	SafeForProductionGate bool     `json:"safe_for_production_gate"`
}

type RuntimeRouteClassificationCounts struct {
	Total           int `json:"total"`
	GoProductLogic  int `json:"go_product_logic"`
	CPolicyBridge   int `json:"c_policy_bridge"`
	RubySmokeBridge int `json:"ruby_smoke_bridge"`
	FixtureOnly     int `json:"fixture_only"`
	ContractOnly    int `json:"contract_only"`
	Deprecated      int `json:"deprecated"`
	Unsupported     int `json:"unsupported"`
	Unclassified    int `json:"unclassified"`
}

type RuntimeRouteMigrationGroup struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Risk          string   `json:"risk"`
	Methods       []string `json:"methods"`
	NextAction    string   `json:"next_action"`
	Ready         bool     `json:"ready"`
	BlockedReason string   `json:"blocked_reason,omitempty"`
}

type RuntimeRouteConvergenceCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewRuntimeRouteConvergencePreview(root string) (RuntimeRouteConvergencePreview, error) {
	manifest, err := NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return RuntimeRouteConvergencePreview{}, err
	}
	routes := runtimeRouteConvergenceRoutes(manifest.Routes)
	counts := countRuntimeRouteClassifications(routes)
	unclassified := runtimeRouteUnclassifiedMethods(routes)
	groups := runtimeRouteMigrationGroups(routes)
	checks := runtimeRouteConvergenceChecks(manifest, counts)

	preview := RuntimeRouteConvergencePreview{
		Version:                     manifest.Version,
		SchemaVersion:               "xnix.runtime.route_convergence.v1",
		RequestType:                 "runtime-route-convergence-preview",
		ReportType:                  "runtime-route-convergence",
		Source:                      "runtime-owner-route-manifest-preview+runtime-method-parity-manifest-preview",
		RuntimeMethod:               "GetRuntimeOwnerRouteManifest",
		ReadMethod:                  "GetRuntimeRouteConvergencePreview",
		Routes:                      routes,
		RouteMethodNames:            runtimeRouteConvergenceMethodNames(routes),
		ClassificationCounts:        counts,
		MigrationGroups:             groups,
		Checks:                      checks,
		CheckIDs:                    runtimeRouteConvergenceCheckIDs(checks),
		AllRoutesClassified:         counts.Unclassified == 0,
		NativeGoCoverageReady:       counts.Total > 0 && counts.GoProductLogic == counts.Total,
		ProductionOwnerReady:        false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		WriteMethodsEnabled:         false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendLaunchEnabled:        false,
		BackendDetailsExposed:       false,
		UnclassifiedRoutes:          unclassified,
		BlockedActions:              runtimeRouteConvergenceBlockedActions(),
		NextMigrationOrder:          runtimeRouteNextMigrationOrder(groups),
		DesktopSafeSummary:          runtimeRouteConvergenceSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "Runtime route convergence preview"); err != nil {
		return RuntimeRouteConvergencePreview{}, err
	}
	return preview, nil
}

func runtimeRouteConvergenceRoutes(routes []RuntimeOwnerRoute) []RuntimeRouteConvergenceRoute {
	converged := make([]RuntimeRouteConvergenceRoute, 0, len(routes))
	for _, route := range routes {
		classification, blocked := runtimeRouteClassification(route)
		converged = append(converged, RuntimeRouteConvergenceRoute{
			Method:                route.Method,
			Domain:                runtimeRouteDomain(route.Method),
			Classification:        classification,
			CurrentSource:         route.CurrentSource,
			TargetSource:          route.TargetSource,
			GoCommand:             route.GoCommand,
			RouteStatus:           route.RouteStatus,
			MigrationPriority:     runtimeRouteMigrationPriority(classification),
			MigrationRisk:         runtimeRouteMigrationRisk(classification),
			BlockedReasons:        blocked,
			KDEVisibleReadOnly:    route.KDEVisibleReadOnly,
			SafeForProductionGate: classification == "go-product-logic",
		})
	}
	return converged
}

func runtimeRouteClassification(route RuntimeOwnerRoute) (string, []string) {
	switch {
	case route.GoRouteReady:
		return "go-product-logic", []string{}
	case route.CCoreBacked:
		return "c-policy-bridge", []string{"needs-go-owner-adapter"}
	case route.LegacyDispatchRequired:
		return "ruby-smoke-bridge", []string{"needs-go-product-route"}
	case route.CurrentSource == "fixture-only":
		return "fixture-only", []string{"needs-product-implementation"}
	case route.CurrentSource == "deprecated":
		return "deprecated", []string{"needs-removal-plan"}
	case route.CurrentSource == "unsupported":
		return "unsupported", []string{"must-fail-closed"}
	case route.Method != "":
		return "contract-only", []string{"needs-implementation-evidence"}
	default:
		return "unclassified", []string{"route-classification-missing"}
	}
}

func runtimeRouteDomain(method string) string {
	switch {
	case hasRuntimeRoutePrefix(method, "GetRuntime"):
		return "runtime-owner"
	case hasRuntimeRoutePrefix(method, "GetKDE") || hasRuntimeRoutePrefix(method, "GetDesktop") || method == "GetTaskManagerIdentityPlan" || method == "GetKWinWindowRulePlan" || method == "GetTrayStatus" || method == "GetNotificationPlan" || method == "GetKRunnerQueryPlan" || method == "GetFileAssociationPlan":
		return "kde-shell"
	case hasRuntimeRoutePrefix(method, "GetCompatibilityAction") || method == "GetCompatibilityCenterSummary":
		return "compatibility-center"
	case method == "GetPortalRequestPlan" || method == "GetPortalAccessPolicy":
		return "portal-safety"
	case method == "GetSnapshotPlan":
		return "snapshot-rollback"
	case method == "GetExecutionReadiness" || method == "GetLaunchIntent":
		return "execution-readiness"
	case hasRuntimeRoutePrefix(method, "GetAI"):
		return "diagnostics-ai"
	case method == "GetCompatibilityInstallPlan" || method == "GetCompatibilityArtifactManifest" || method == "GetCompatibilityAcquisitionPreflight" || method == "GetCompatibilityPackageSource":
		return "recipe-artifact"
	case method == "GetBackendBinding" || method == "GetBackendCapabilityMatrix" || method == "GetBackendSelectionPlan" || method == "GetBackendLifecycle" || method == "GetBackendEnvironmentPlan":
		return "environment-lifecycle"
	default:
		return "runtime-catalog"
	}
}

func hasRuntimeRoutePrefix(value string, prefix string) bool {
	if len(value) < len(prefix) {
		return false
	}
	return value[:len(prefix)] == prefix
}

func runtimeRouteMigrationPriority(classification string) int {
	switch classification {
	case "contract-only", "unclassified":
		return 1
	case "ruby-smoke-bridge", "fixture-only":
		return 2
	case "c-policy-bridge":
		return 3
	case "go-product-logic":
		return 4
	case "deprecated", "unsupported":
		return 5
	default:
		return 1
	}
}

func runtimeRouteMigrationRisk(classification string) string {
	switch classification {
	case "go-product-logic":
		return "low"
	case "c-policy-bridge":
		return "medium"
	case "ruby-smoke-bridge", "fixture-only", "contract-only":
		return "high"
	case "deprecated", "unsupported":
		return "review-required"
	default:
		return "blocked"
	}
}

func countRuntimeRouteClassifications(routes []RuntimeRouteConvergenceRoute) RuntimeRouteClassificationCounts {
	counts := RuntimeRouteClassificationCounts{Total: len(routes)}
	for _, route := range routes {
		switch route.Classification {
		case "go-product-logic":
			counts.GoProductLogic++
		case "c-policy-bridge":
			counts.CPolicyBridge++
		case "ruby-smoke-bridge":
			counts.RubySmokeBridge++
		case "fixture-only":
			counts.FixtureOnly++
		case "contract-only":
			counts.ContractOnly++
		case "deprecated":
			counts.Deprecated++
		case "unsupported":
			counts.Unsupported++
		default:
			counts.Unclassified++
		}
	}
	return counts
}

func runtimeRouteMigrationGroups(routes []RuntimeRouteConvergenceRoute) []RuntimeRouteMigrationGroup {
	groupOrder := []string{"runtime-owner", "recipe-artifact", "environment-lifecycle", "portal-safety", "snapshot-rollback", "kde-shell", "compatibility-center", "execution-readiness", "diagnostics-ai", "runtime-catalog"}
	methodsByGroup := map[string][]string{}
	readyByGroup := map[string]bool{}
	for _, group := range groupOrder {
		readyByGroup[group] = true
	}
	for _, route := range routes {
		methodsByGroup[route.Domain] = append(methodsByGroup[route.Domain], route.Method)
		if route.Classification != "go-product-logic" {
			readyByGroup[route.Domain] = false
		}
	}

	groups := []RuntimeRouteMigrationGroup{}
	for _, groupID := range groupOrder {
		methods := methodsByGroup[groupID]
		if len(methods) == 0 {
			continue
		}
		ready := readyByGroup[groupID]
		group := RuntimeRouteMigrationGroup{
			ID:         groupID,
			Title:      runtimeRouteGroupTitle(groupID),
			Risk:       runtimeRouteGroupRisk(ready),
			Methods:    methods,
			NextAction: runtimeRouteGroupNextAction(groupID, ready),
			Ready:      ready,
		}
		if !ready {
			group.BlockedReason = "one or more routes still need Go product ownership evidence"
		}
		groups = append(groups, group)
	}
	return groups
}

func runtimeRouteGroupTitle(groupID string) string {
	titles := map[string]string{
		"runtime-owner":         "Runtime owner boundary",
		"recipe-artifact":       "Recipe and artifact trust",
		"environment-lifecycle": "Environment lifecycle",
		"portal-safety":         "Portal safety",
		"snapshot-rollback":     "Snapshot and rollback",
		"kde-shell":             "KDE shell consumers",
		"compatibility-center":  "Compatibility Center actions",
		"execution-readiness":   "Execution readiness",
		"diagnostics-ai":        "Diagnostics and AI",
		"runtime-catalog":       "Runtime catalog",
	}
	if title, ok := titles[groupID]; ok {
		return title
	}
	return "Runtime route group"
}

func runtimeRouteGroupRisk(ready bool) string {
	if ready {
		return "low"
	}
	return "high"
}

func runtimeRouteGroupNextAction(groupID string, ready bool) string {
	if ready {
		return "keep covered by drift and layout verification"
	}
	switch groupID {
	case "runtime-owner":
		return "migrate owner boundary routes before production bus ownership"
	case "kde-shell":
		return "keep KDE as a read-only consumer while Runtime route ownership converges"
	case "execution-readiness":
		return "keep launch disabled until execution evidence routes converge"
	default:
		return "add Go product logic and blocked-state tests for pending routes"
	}
}

func runtimeRouteConvergenceChecks(manifest RuntimeOwnerRouteManifestPreview, counts RuntimeRouteClassificationCounts) []RuntimeRouteConvergenceCheck {
	return []RuntimeRouteConvergenceCheck{
		runtimeRouteConvergenceCheck("owner-route-manifest-ready", routeConvergencePassBlocked(manifest.GoOwnerRouteCoverageReady), "Owner route manifest must cover every read-only Runtime method."),
		runtimeRouteConvergenceCheck("all-routes-classified", routeConvergencePassBlocked(counts.Unclassified == 0), "Every route must have an implementation classification."),
		runtimeRouteConvergenceCheck("native-go-route-coverage", routeConvergencePendingUnless(counts.Total > 0 && counts.GoProductLogic == counts.Total), "Every read-only route should converge on Go product logic or an explicit low-level adapter boundary."),
		runtimeRouteConvergenceCheck("write-gate-disabled", "pass", "Route convergence cannot enable Runtime write methods."),
		runtimeRouteConvergenceCheck("host-safety-boundary", "pass", "Route convergence preview must not run Docker, QEMU, network, backends, or host-root writes."),
	}
}

func runtimeRouteConvergenceCheck(id string, status string, summary string) RuntimeRouteConvergenceCheck {
	return RuntimeRouteConvergenceCheck{ID: id, Status: status, Summary: summary}
}

func routeConvergencePassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func routeConvergencePendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func runtimeRouteConvergenceCheckIDs(checks []RuntimeRouteConvergenceCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func runtimeRouteConvergenceMethodNames(routes []RuntimeRouteConvergenceRoute) []string {
	methods := make([]string, 0, len(routes))
	for _, route := range routes {
		methods = append(methods, route.Method)
	}
	return methods
}

func runtimeRouteUnclassifiedMethods(routes []RuntimeRouteConvergenceRoute) []string {
	methods := []string{}
	for _, route := range routes {
		if route.Classification == "unclassified" {
			methods = append(methods, route.Method)
		}
	}
	return methods
}

func runtimeRouteNextMigrationOrder(groups []RuntimeRouteMigrationGroup) []string {
	order := make([]string, 0, len(groups))
	for _, group := range groups {
		if !group.Ready {
			order = append(order, group.ID)
		}
	}
	if len(order) == 0 {
		for _, group := range groups {
			order = append(order, group.ID)
		}
	}
	return order
}

func runtimeRouteConvergenceBlockedActions() []string {
	return []string{
		"enable Runtime write methods from route convergence preview",
		"claim production D-Bus ownership from route convergence preview",
		"start compatibility backends from route convergence preview",
		"run Docker or QEMU from route convergence preview",
		"mutate host root during route convergence preview",
		"let KDE own Runtime policy routes",
	}
}

func runtimeRouteConvergenceSummary(counts RuntimeRouteClassificationCounts) string {
	if counts.Unclassified > 0 {
		return "Runtime route convergence is blocked by unclassified read-only routes."
	}
	if counts.ContractOnly > 0 || counts.RubySmokeBridge > 0 || counts.FixtureOnly > 0 {
		return "Runtime route convergence still needs Go product ownership for one or more read-only route groups."
	}
	if counts.CPolicyBridge > 0 {
		return "Runtime route convergence still needs explicit Go owner adapter boundaries for low-level C policy routes."
	}
	return "Runtime read-only routes are classified and covered by Go product logic; production ownership remains gated."
}
