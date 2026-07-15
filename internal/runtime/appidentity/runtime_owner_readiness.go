package appidentity

type RuntimeOwnerReadinessPreview struct {
	Version                     string                             `json:"version"`
	SchemaVersion               string                             `json:"schema_version"`
	RequestType                 string                             `json:"request_type"`
	ReadinessType               string                             `json:"readiness_type"`
	Source                      string                             `json:"source"`
	RuntimeMethod               string                             `json:"runtime_method"`
	ReadMethod                  string                             `json:"read_method"`
	BusName                     string                             `json:"bus_name"`
	ObjectPath                  string                             `json:"object_path"`
	Interface                   string                             `json:"interface"`
	ServiceBinding              RuntimeOwnerReadinessBinding       `json:"service_binding"`
	LiveOwnerGate               RuntimeOwnerReadinessLiveGate      `json:"live_owner_gate"`
	OwnerProcess                RuntimeOwnerReadinessProcess       `json:"owner_process"`
	OwnerSmokePlan              RuntimeOwnerReadinessSmokePlan     `json:"owner_smoke_plan"`
	MethodParityManifest        RuntimeOwnerReadinessMethodParity  `json:"method_parity_manifest"`
	OwnerRouteManifest          RuntimeOwnerReadinessRouteManifest `json:"owner_route_manifest"`
	RecipeTrust                 RuntimeOwnerReadinessRecipeTrust   `json:"recipe_trust"`
	ReadinessChecks             []RuntimeOwnerReadinessCheck       `json:"readiness_checks"`
	CheckIDs                    []string                           `json:"check_ids"`
	Counts                      RuntimeOwnerReadinessCounts        `json:"counts"`
	ActivationBindingReady      bool                               `json:"activation_binding_ready"`
	ReadOnlyMethodParityReady   bool                               `json:"read_only_method_parity_ready"`
	OwnerSmokePlanned           bool                               `json:"owner_smoke_planned"`
	LiveDBusOwnerReady          bool                               `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled      bool                               `json:"production_owner_enabled"`
	OwnerTransitionReady        bool                               `json:"owner_transition_ready"`
	ProductionRecipeTrustReady  bool                               `json:"production_recipe_trust_ready"`
	ProductionOwnerRoutesReady  bool                               `json:"production_owner_routes_ready"`
	RuntimeOwned                bool                               `json:"runtime_owned"`
	GoRuntimeBacked             bool                               `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                               `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool                               `json:"kde_may_claim_runtime_ownership"`
	NetworkRequired             bool                               `json:"network_required"`
	HostRootModified            bool                               `json:"host_root_modified"`
	PrivilegedContainerRequired bool                               `json:"privileged_container_required"`
	SystemServiceStarted        bool                               `json:"system_service_started"`
	ProductionBusClaimed        bool                               `json:"production_bus_claimed"`
	WriteMethodsEnabled         bool                               `json:"write_methods_enabled"`
	BackendDetailsExposed       bool                               `json:"backend_details_exposed"`
	BlockedReasons              []string                           `json:"blocked_reasons"`
	BlockedActions              []string                           `json:"blocked_actions"`
	DesktopSafeSummary          string                             `json:"desktop_safe_summary"`
}

type RuntimeOwnerReadinessBinding struct {
	RequestType            string                      `json:"request_type"`
	ProductionStatus       string                      `json:"production_status"`
	ActivationBindingReady bool                        `json:"activation_binding_ready"`
	LiveDBusOwnerReady     bool                        `json:"live_dbus_owner_ready"`
	SmokeAdapterAvailable  bool                        `json:"smoke_adapter_available"`
	Counts                 RuntimeServiceBindingCounts `json:"counts"`
}

type RuntimeOwnerReadinessLiveGate struct {
	RequestType              string                     `json:"request_type"`
	GateType                 string                     `json:"gate_type"`
	ActivationBindingReady   bool                       `json:"activation_binding_ready"`
	LiveDBusOwnerReady       bool                       `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled   bool                       `json:"production_owner_enabled"`
	OwnerTransitionReady     bool                       `json:"owner_transition_ready"`
	SmokeAdapterIsProduction bool                       `json:"smoke_adapter_is_production_owner"`
	Counts                   RuntimeLiveOwnerGateCounts `json:"counts"`
}

type RuntimeOwnerReadinessProcess struct {
	RequestType                 string                    `json:"request_type"`
	ProcessType                 string                    `json:"process_type"`
	CurrentOwnerLanguage        string                    `json:"current_owner_language"`
	TargetOwnerLanguage         string                    `json:"target_owner_language"`
	ServiceActivationReady      bool                      `json:"service_activation_ready"`
	PackagedEntrypointReady     bool                      `json:"packaged_entrypoint_ready"`
	GoOwnerProcessReady         bool                      `json:"go_owner_process_ready"`
	ProductionOwnerProcessReady bool                      `json:"production_owner_process_ready"`
	Counts                      RuntimeOwnerProcessCounts `json:"counts"`
}

type RuntimeOwnerReadinessSmokePlan struct {
	RequestType            string                  `json:"request_type"`
	PlanType               string                  `json:"plan_type"`
	SmokeState             string                  `json:"smoke_state"`
	SmokeEnvironment       string                  `json:"smoke_environment"`
	ActivationBindingReady bool                    `json:"activation_binding_ready"`
	PendingStepCount       int                     `json:"pending_step_count"`
	Counts                 RuntimeOwnerSmokeCounts `json:"counts"`
}

type RuntimeOwnerReadinessMethodParity struct {
	RequestType                string                    `json:"request_type"`
	ManifestType               string                    `json:"manifest_type"`
	ReadOnlyMethodParityReady  bool                      `json:"read_only_method_parity_ready"`
	MethodCount                int                       `json:"method_count"`
	WriteMethodsSupported      bool                      `json:"write_methods_supported"`
	WriteMethodDispatchEnabled bool                      `json:"write_method_dispatch_enabled"`
	Counts                     RuntimeMethodParityCounts `json:"counts"`
}

type RuntimeOwnerReadinessRouteManifest struct {
	RequestType                string                          `json:"request_type"`
	ManifestType               string                          `json:"manifest_type"`
	RouteCount                 int                             `json:"route_count"`
	GoRouteCount               int                             `json:"go_route_count"`
	CCoreRouteCount            int                             `json:"c_core_route_count"`
	RubyLegacyRouteCount       int                             `json:"ruby_legacy_route_count"`
	GoOwnerRouteCoverageReady  bool                            `json:"go_owner_route_coverage_ready"`
	CCoreAdapterRequired       bool                            `json:"c_core_adapter_required"`
	LegacyRuntimeRoutesPresent bool                            `json:"legacy_runtime_routes_present"`
	ProductionOwnerRoutesReady bool                            `json:"production_owner_routes_ready"`
	Counts                     RuntimeOwnerRouteManifestCounts `json:"counts"`
}

type RuntimeOwnerReadinessRecipeTrust struct {
	RequestType                string                        `json:"request_type"`
	TrustType                  string                        `json:"trust_type"`
	RegistryName               string                        `json:"registry_name"`
	RecipeCount                int                           `json:"recipe_count"`
	DigestVerified             bool                          `json:"digest_verified"`
	SignedRecipeValidation     bool                          `json:"signed_recipe_validation"`
	DevelopmentRegistry        bool                          `json:"development_registry"`
	UnsignedRecipesPresent     bool                          `json:"unsigned_recipes_present"`
	ProductionRecipeTrustReady bool                          `json:"production_recipe_trust_ready"`
	Counts                     RuntimeOwnerRecipeTrustCounts `json:"counts"`
}

type RuntimeOwnerReadinessCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeOwnerReadinessCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeOwnerReadinessPreview(root string) (RuntimeOwnerReadinessPreview, error) {
	serviceBinding, err := NewRuntimeServiceBindingPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	liveOwnerGate, err := NewRuntimeLiveOwnerGatePreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	ownerProcess, err := NewRuntimeOwnerProcessPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	ownerSmokePlan, err := NewRuntimeOwnerSmokePlanPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	methodParityManifest, err := NewRuntimeMethodParityManifestPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	ownerRouteManifest, err := NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	recipeTrust, err := NewRuntimeOwnerRecipeTrustPreview(root)
	if err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}

	checks := runtimeOwnerReadinessChecks(serviceBinding, liveOwnerGate, ownerProcess, ownerSmokePlan, methodParityManifest, ownerRouteManifest, recipeTrust)
	counts := countRuntimeOwnerReadinessChecks(checks)

	preview := RuntimeOwnerReadinessPreview{
		Version:       serviceBinding.Version,
		SchemaVersion: "xnix.runtime.owner_readiness.v1",
		RequestType:   "runtime-owner-readiness-preview",
		ReadinessType: "runtime-owner-readiness",
		Source:        "runtime-service-binding-preview+runtime-live-owner-gate-preview+runtime-owner-process-preview+runtime-owner-smoke-plan-preview+runtime-method-parity-manifest-preview+runtime-owner-route-manifest-preview+runtime-owner-recipe-trust-preview",
		RuntimeMethod: "GetRuntimeOwnerReadiness",
		ReadMethod:    "GetRuntimeOwnerReadinessPreview",
		BusName:       serviceBinding.BusName,
		ObjectPath:    serviceBinding.ObjectPath,
		Interface:     serviceBinding.Interface,
		ServiceBinding: RuntimeOwnerReadinessBinding{
			RequestType:            serviceBinding.RequestType,
			ProductionStatus:       serviceBinding.ProductionStatus,
			ActivationBindingReady: serviceBinding.ActivationBindingReady,
			LiveDBusOwnerReady:     serviceBinding.LiveDBusOwnerReady,
			SmokeAdapterAvailable:  serviceBinding.SmokeAdapterAvailable,
			Counts:                 serviceBinding.Counts,
		},
		LiveOwnerGate: RuntimeOwnerReadinessLiveGate{
			RequestType:              liveOwnerGate.RequestType,
			GateType:                 liveOwnerGate.GateType,
			ActivationBindingReady:   liveOwnerGate.ActivationBindingReady,
			LiveDBusOwnerReady:       liveOwnerGate.LiveDBusOwnerReady,
			ProductionOwnerEnabled:   liveOwnerGate.ProductionOwnerEnabled,
			OwnerTransitionReady:     liveOwnerGate.OwnerTransitionReady,
			SmokeAdapterIsProduction: liveOwnerGate.SmokeAdapterIsProduction,
			Counts:                   liveOwnerGate.Counts,
		},
		OwnerProcess: RuntimeOwnerReadinessProcess{
			RequestType:                 ownerProcess.RequestType,
			ProcessType:                 ownerProcess.ProcessType,
			CurrentOwnerLanguage:        ownerProcess.CurrentOwnerLanguage,
			TargetOwnerLanguage:         ownerProcess.TargetOwnerLanguage,
			ServiceActivationReady:      ownerProcess.ServiceActivationReady,
			PackagedEntrypointReady:     ownerProcess.PackagedEntrypointReady,
			GoOwnerProcessReady:         ownerProcess.GoOwnerProcessReady,
			ProductionOwnerProcessReady: ownerProcess.ProductionOwnerProcessReady,
			Counts:                      ownerProcess.Counts,
		},
		OwnerSmokePlan: RuntimeOwnerReadinessSmokePlan{
			RequestType:            ownerSmokePlan.RequestType,
			PlanType:               ownerSmokePlan.PlanType,
			SmokeState:             ownerSmokePlan.SmokeState,
			SmokeEnvironment:       ownerSmokePlan.SmokeEnvironment,
			ActivationBindingReady: ownerSmokePlan.ActivationBindingReady,
			PendingStepCount:       ownerSmokePlan.PendingStepCount,
			Counts:                 ownerSmokePlan.Counts,
		},
		MethodParityManifest: RuntimeOwnerReadinessMethodParity{
			RequestType:                methodParityManifest.RequestType,
			ManifestType:               methodParityManifest.ManifestType,
			ReadOnlyMethodParityReady:  methodParityManifest.ReadOnlyMethodParityReady,
			MethodCount:                methodParityManifest.MethodCount,
			WriteMethodsSupported:      methodParityManifest.WriteMethodsSupported,
			WriteMethodDispatchEnabled: methodParityManifest.WriteMethodDispatchEnabled,
			Counts:                     methodParityManifest.Counts,
		},
		OwnerRouteManifest: RuntimeOwnerReadinessRouteManifest{
			RequestType:                ownerRouteManifest.RequestType,
			ManifestType:               ownerRouteManifest.ManifestType,
			RouteCount:                 ownerRouteManifest.RouteCounts.Total,
			GoRouteCount:               ownerRouteManifest.RouteCounts.GoRouted,
			CCoreRouteCount:            ownerRouteManifest.RouteCounts.CCoreBacked,
			RubyLegacyRouteCount:       ownerRouteManifest.RouteCounts.RubyLegacy,
			GoOwnerRouteCoverageReady:  ownerRouteManifest.GoOwnerRouteCoverageReady,
			CCoreAdapterRequired:       ownerRouteManifest.CCoreAdapterRequired,
			LegacyRuntimeRoutesPresent: ownerRouteManifest.LegacyRuntimeRoutesPresent,
			ProductionOwnerRoutesReady: ownerRouteManifest.ProductionOwnerRoutesReady,
			Counts:                     ownerRouteManifest.Counts,
		},
		RecipeTrust: RuntimeOwnerReadinessRecipeTrust{
			RequestType:                recipeTrust.RequestType,
			TrustType:                  recipeTrust.TrustType,
			RegistryName:               recipeTrust.RegistryName,
			RecipeCount:                recipeTrust.RecipeCount,
			DigestVerified:             recipeTrust.DigestVerified,
			SignedRecipeValidation:     recipeTrust.SignedRecipeValidation,
			DevelopmentRegistry:        recipeTrust.DevelopmentRegistry,
			UnsignedRecipesPresent:     recipeTrust.UnsignedRecipesPresent,
			ProductionRecipeTrustReady: recipeTrust.ProductionRecipeTrustReady,
			Counts:                     recipeTrust.Counts,
		},
		ReadinessChecks:             checks,
		CheckIDs:                    runtimeOwnerReadinessCheckIDs(checks),
		Counts:                      counts,
		ActivationBindingReady:      serviceBinding.ActivationBindingReady,
		ReadOnlyMethodParityReady:   methodParityManifest.ReadOnlyMethodParityReady,
		OwnerSmokePlanned:           ownerSmokePlan.SmokeState == "planned",
		LiveDBusOwnerReady:          false,
		ProductionOwnerEnabled:      false,
		OwnerTransitionReady:        false,
		ProductionRecipeTrustReady:  recipeTrust.ProductionRecipeTrustReady,
		ProductionOwnerRoutesReady:  ownerRouteManifest.ProductionOwnerRoutesReady,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		SystemServiceStarted:        false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		BackendDetailsExposed:       false,
		BlockedReasons:              runtimeOwnerReadinessBlockedReasons(),
		BlockedActions:              runtimeOwnerReadinessBlockedActions(),
		DesktopSafeSummary:          runtimeOwnerReadinessSummary(serviceBinding.ActivationBindingReady, methodParityManifest.ReadOnlyMethodParityReady, ownerRouteManifest.ProductionOwnerRoutesReady, recipeTrust.ProductionRecipeTrustReady),
	}
	if err := validateNoBackendTerms(preview, "Runtime owner readiness preview"); err != nil {
		return RuntimeOwnerReadinessPreview{}, err
	}
	return preview, nil
}

func runtimeOwnerReadinessChecks(serviceBinding RuntimeServiceBindingPreview, liveOwnerGate RuntimeLiveOwnerGatePreview, ownerProcess RuntimeOwnerProcessPreview, ownerSmokePlan RuntimeOwnerSmokePlanPreview, methodParityManifest RuntimeMethodParityManifestPreview, ownerRouteManifest RuntimeOwnerRouteManifestPreview, recipeTrust RuntimeOwnerRecipeTrustPreview) []RuntimeOwnerReadinessCheck {
	return []RuntimeOwnerReadinessCheck{
		runtimeOwnerReadinessCheck("activation-binding", runtimeOwnerReadinessPassBlocked(serviceBinding.ActivationBindingReady), "D-Bus activation files, systemd unit, libexec wrapper, and contract must be aligned."),
		runtimeOwnerReadinessCheck("read-only-method-parity", runtimeOwnerReadinessPassBlocked(methodParityManifest.ReadOnlyMethodParityReady), "The production owner must cover every read-only Runtime method required by KDE."),
		runtimeOwnerReadinessCheck("owner-smoke-plan", runtimeOwnerReadinessPassBlocked(ownerSmokePlan.SmokeState == "planned"), "Production owner smoke must have a deterministic restricted-session plan."),
		runtimeOwnerReadinessCheck("write-method-gate", runtimeOwnerReadinessPassBlocked(!methodParityManifest.WriteMethodsSupported && !methodParityManifest.WriteMethodDispatchEnabled), "Install, launch, snapshot, and restore writes must remain disabled until the Runtime owner is production-ready."),
		runtimeOwnerReadinessCheck("kde-ownership-boundary", runtimeOwnerReadinessPassBlocked(!serviceBinding.KDEMayClaimRuntimeOwnership && !liveOwnerGate.KDEMayClaimRuntimeOwnership && !liveOwnerGate.KDEPolicyOwner), "KDE must remain a presentation shell and must not own Runtime policy."),
		runtimeOwnerReadinessCheck("host-safety-boundary", "pass", "Readiness preview must not start services, claim bus names, require network, or mutate the host root."),
		runtimeOwnerReadinessCheck("long-running-runtime-owner", runtimeOwnerProcessReadinessStatus(ownerProcess), "A packaged Go long-running Runtime owner still needs implementation before production ownership."),
		runtimeOwnerReadinessCheck("read-only-owner-routes", runtimeOwnerRouteManifestReadinessStatus(ownerRouteManifest), "Read-only Runtime methods need native Go owner routes or explicit owner adapter boundaries before production ownership."),
		runtimeOwnerReadinessCheck("production-bus-claim", "pending", "A production smoke must prove the packaged Runtime owner owns org.xnix.Compatibility1."),
		runtimeOwnerReadinessCheck("production-recipe-trust", runtimeOwnerRecipeTrustReadinessStatus(recipeTrust), "Production owner readiness requires digest-verified and production-signed recipes."),
	}
}

func runtimeOwnerProcessReadinessStatus(ownerProcess RuntimeOwnerProcessPreview) string {
	if ownerProcess.ProductionOwnerProcessReady {
		return "pass"
	}
	if ownerProcess.Counts.Blocked > 0 {
		return "blocked"
	}
	return "pending"
}

func runtimeOwnerRecipeTrustReadinessStatus(recipeTrust RuntimeOwnerRecipeTrustPreview) string {
	if recipeTrust.ProductionRecipeTrustReady {
		return "pass"
	}
	if recipeTrust.Counts.Blocked > 0 {
		return "blocked"
	}
	return "pending"
}

func runtimeOwnerRouteManifestReadinessStatus(ownerRouteManifest RuntimeOwnerRouteManifestPreview) string {
	if ownerRouteManifest.ProductionOwnerRoutesReady {
		return "pass"
	}
	if ownerRouteManifest.Counts.Blocked > 0 || ownerRouteManifest.RouteCounts.Blocked > 0 {
		return "blocked"
	}
	return "pending"
}

func runtimeOwnerReadinessCheck(id string, status string, summary string) RuntimeOwnerReadinessCheck {
	return RuntimeOwnerReadinessCheck{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeOwnerReadinessPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeOwnerReadinessCheckIDs(checks []RuntimeOwnerReadinessCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRuntimeOwnerReadinessChecks(checks []RuntimeOwnerReadinessCheck) RuntimeOwnerReadinessCounts {
	counts := RuntimeOwnerReadinessCounts{Total: len(checks)}
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

func runtimeOwnerReadinessBlockedReasons() []string {
	return []string{
		"Live production Runtime ownership is still pending.",
		"Production bus-name acquisition has not been proven by a packaged owner smoke.",
		"Production recipe trust still requires signed registry validation.",
		"Go owner route coverage still needs C adapter boundaries for C-backed read-only routes.",
		"KDE remains a replaceable presentation shell and must not own Runtime policy.",
		"Write methods stay disabled until the production Runtime owner is proven.",
	}
}

func runtimeOwnerReadinessBlockedActions() []string {
	return []string{
		"start production Runtime owner from readiness preview",
		"claim production D-Bus name from readiness preview",
		"mark smoke adapter as production owner",
		"serve C-backed read-only routes without Go owner adapters",
		"enable launch, install, snapshot, restore, repair, or settings persistence",
		"let KDE claim Runtime ownership",
		"mutate host root during readiness evaluation",
	}
}

func runtimeOwnerReadinessSummary(activationBindingReady bool, methodParityReady bool, ownerRoutesReady bool, recipeTrustReady bool) string {
	if !activationBindingReady || !methodParityReady {
		return "Runtime owner readiness is blocked by activation or method parity defects."
	}
	if !ownerRoutesReady {
		return "Runtime owner readiness has aligned activation and method parity; C owner adapter migration, live production ownership, bus claim, and production-signed recipe trust remain pending."
	}
	if !recipeTrustReady {
		return "Runtime owner readiness has aligned activation and method parity; live production ownership, bus claim, and production-signed recipe trust remain pending."
	}
	return "Runtime owner readiness has aligned activation, method parity, and recipe trust; live production ownership and bus claim remain pending."
}
