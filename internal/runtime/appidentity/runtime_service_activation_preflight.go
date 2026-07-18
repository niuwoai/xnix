package appidentity

import "strings"

type RuntimeServiceActivationPreflightPreview struct {
	Version                          string                                     `json:"version"`
	SchemaVersion                    string                                     `json:"schema_version"`
	RequestType                      string                                     `json:"request_type"`
	PreflightType                    string                                     `json:"preflight_type"`
	Source                           string                                     `json:"source"`
	RuntimeMethod                    string                                     `json:"runtime_method"`
	ReadMethod                       string                                     `json:"read_method"`
	BusName                          string                                     `json:"bus_name"`
	ObjectPath                       string                                     `json:"object_path"`
	Interface                        string                                     `json:"interface"`
	ServiceBinding                   RuntimeServiceActivationPreflightBinding   `json:"service_binding"`
	OwnerReadiness                   RuntimeServiceActivationPreflightReadiness `json:"owner_readiness"`
	OwnerSmokePlan                   RuntimeServiceActivationPreflightSmokePlan `json:"owner_smoke_plan"`
	ProductionDBusGate               RuntimeServiceActivationProductionDBusGate `json:"production_dbus_gate"`
	PreflightChecks                  []RuntimeServiceActivationPreflightCheck   `json:"preflight_checks"`
	CheckIDs                         []string                                   `json:"check_ids"`
	Counts                           RuntimeServiceActivationPreflightCounts    `json:"counts"`
	PreflightDecision                string                                     `json:"preflight_decision"`
	ProductionActivationReady        bool                                       `json:"production_activation_ready"`
	RestrictedSmokeReady             bool                                       `json:"restricted_smoke_ready"`
	ProductionDBusGateReady          bool                                       `json:"production_dbus_gate_ready"`
	HumanAuthorizationPreflightReady bool                                       `json:"human_authorization_preflight_ready"`
	HumanAuthorizationRequired       bool                                       `json:"human_authorization_required"`
	HumanAuthorizationGranted        bool                                       `json:"human_authorization_granted"`
	AuthorizationReceiptAccepted     bool                                       `json:"authorization_receipt_accepted"`
	RuntimeOwned                     bool                                       `json:"runtime_owned"`
	GoRuntimeBacked                  bool                                       `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool                                       `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership      bool                                       `json:"kde_may_claim_runtime_ownership"`
	SystemServiceStarted             bool                                       `json:"system_service_started"`
	ProductionBusClaimed             bool                                       `json:"production_bus_claimed"`
	WriteMethodsEnabled              bool                                       `json:"write_methods_enabled"`
	BackendLaunchEnabled             bool                                       `json:"backend_launch_enabled"`
	NetworkRequired                  bool                                       `json:"network_required"`
	HostRootModified                 bool                                       `json:"host_root_modified"`
	PrivilegedContainerRequired      bool                                       `json:"privileged_container_required"`
	BackendDetailsExposed            bool                                       `json:"backend_details_exposed"`
	BlockedReasons                   []string                                   `json:"blocked_reasons"`
	BlockedActions                   []string                                   `json:"blocked_actions"`
	NextRequirements                 []string                                   `json:"next_requirements"`
	DesktopSafeSummary               string                                     `json:"desktop_safe_summary"`
}

type RuntimeServiceActivationPreflightBinding struct {
	RequestType            string                      `json:"request_type"`
	ProductionStatus       string                      `json:"production_status"`
	ActivationBindingReady bool                        `json:"activation_binding_ready"`
	LiveDBusOwnerReady     bool                        `json:"live_dbus_owner_ready"`
	SmokeAdapterAvailable  bool                        `json:"smoke_adapter_available"`
	Counts                 RuntimeServiceBindingCounts `json:"counts"`
}

type RuntimeServiceActivationPreflightReadiness struct {
	RequestType                string                      `json:"request_type"`
	ReadinessType              string                      `json:"readiness_type"`
	ActivationBindingReady     bool                        `json:"activation_binding_ready"`
	ReadOnlyMethodParityReady  bool                        `json:"read_only_method_parity_ready"`
	OwnerSmokePlanned          bool                        `json:"owner_smoke_planned"`
	ProductionOwnerRoutesReady bool                        `json:"production_owner_routes_ready"`
	ProductionRecipeTrustReady bool                        `json:"production_recipe_trust_ready"`
	LiveDBusOwnerReady         bool                        `json:"live_dbus_owner_ready"`
	ProductionOwnerEnabled     bool                        `json:"production_owner_enabled"`
	OwnerTransitionReady       bool                        `json:"owner_transition_ready"`
	SystemServiceStarted       bool                        `json:"system_service_started"`
	ProductionBusClaimed       bool                        `json:"production_bus_claimed"`
	WriteMethodsEnabled        bool                        `json:"write_methods_enabled"`
	Counts                     RuntimeOwnerReadinessCounts `json:"counts"`
}

type RuntimeServiceActivationPreflightSmokePlan struct {
	RequestType            string                  `json:"request_type"`
	PlanType               string                  `json:"plan_type"`
	SmokeState             string                  `json:"smoke_state"`
	SmokeEnvironment       string                  `json:"smoke_environment"`
	ActivationBindingReady bool                    `json:"activation_binding_ready"`
	PendingStepCount       int                     `json:"pending_step_count"`
	Counts                 RuntimeOwnerSmokeCounts `json:"counts"`
}

type RuntimeServiceActivationProductionDBusGate struct {
	RequestType                        string `json:"request_type"`
	GateType                           string `json:"gate_type"`
	GateDecision                       string `json:"gate_decision"`
	GateReviewSourcePresent            bool   `json:"gate_review_source_present"`
	HumanAuthorizationPreflightPresent bool   `json:"human_authorization_preflight_present"`
	GateReviewReady                    bool   `json:"gate_review_ready"`
	HumanAuthorizationPreflightReady   bool   `json:"human_authorization_preflight_ready"`
	HumanAuthorizationRequired         bool   `json:"human_authorization_required"`
	HumanAuthorizationGranted          bool   `json:"human_authorization_granted"`
	AuthorizationReceiptAccepted       bool   `json:"authorization_receipt_accepted"`
	ProductionReadiness                bool   `json:"production_readiness"`
	ProductionOwnerEnabled             bool   `json:"production_owner_enabled"`
	ProductionActivationReady          bool   `json:"production_activation_ready"`
	SystemServiceStarted               bool   `json:"system_service_started"`
	SessionBusClaimed                  bool   `json:"session_bus_claimed"`
	ProductionBusClaimed               bool   `json:"production_bus_claimed"`
	WriteMethodsEnabled                bool   `json:"write_methods_enabled"`
	RuntimeWritesEnabled               bool   `json:"runtime_writes_enabled"`
	BackendLaunchEnabled               bool   `json:"backend_launch_enabled"`
	HostRootModified                   bool   `json:"host_root_modified"`
}

type RuntimeServiceActivationPreflightCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeServiceActivationPreflightCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeServiceActivationPreflightPreview(root string) (RuntimeServiceActivationPreflightPreview, error) {
	serviceBinding, err := NewRuntimeServiceBindingPreview(root)
	if err != nil {
		return RuntimeServiceActivationPreflightPreview{}, err
	}
	ownerReadiness, err := NewRuntimeOwnerReadinessPreview(root)
	if err != nil {
		return RuntimeServiceActivationPreflightPreview{}, err
	}
	ownerSmokePlan, err := NewRuntimeOwnerSmokePlanPreview(root)
	if err != nil {
		return RuntimeServiceActivationPreflightPreview{}, err
	}
	productionDBusGate := runtimeServiceActivationProductionDBusGate(root)

	checks := runtimeServiceActivationPreflightChecks(ownerReadiness, productionDBusGate)
	counts := countRuntimeServiceActivationPreflightChecks(checks)
	restrictedSmokeReady := runtimeServiceActivationRestrictedSmokeReady(ownerReadiness, ownerSmokePlan)
	productionActivationReady := runtimeServiceActivationProductionReady(ownerReadiness, productionDBusGate)

	preview := RuntimeServiceActivationPreflightPreview{
		Version:       serviceBinding.Version,
		SchemaVersion: "xnix.runtime.service_activation_preflight.v1",
		RequestType:   "runtime-service-activation-preflight-preview",
		PreflightType: "production-runtime-service-activation-preflight",
		Source:        "runtime-service-binding-preview+runtime-owner-readiness-preview+runtime-owner-smoke-plan-preview+production-dbus-gate-review-preview+production-dbus-human-authorization-preflight-preview",
		RuntimeMethod: "GetRuntimeServiceActivationPreflight",
		ReadMethod:    "GetRuntimeServiceActivationPreflightPreview",
		BusName:       serviceBinding.BusName,
		ObjectPath:    serviceBinding.ObjectPath,
		Interface:     serviceBinding.Interface,
		ServiceBinding: RuntimeServiceActivationPreflightBinding{
			RequestType:            serviceBinding.RequestType,
			ProductionStatus:       serviceBinding.ProductionStatus,
			ActivationBindingReady: serviceBinding.ActivationBindingReady,
			LiveDBusOwnerReady:     serviceBinding.LiveDBusOwnerReady,
			SmokeAdapterAvailable:  serviceBinding.SmokeAdapterAvailable,
			Counts:                 serviceBinding.Counts,
		},
		OwnerReadiness: RuntimeServiceActivationPreflightReadiness{
			RequestType:                ownerReadiness.RequestType,
			ReadinessType:              ownerReadiness.ReadinessType,
			ActivationBindingReady:     ownerReadiness.ActivationBindingReady,
			ReadOnlyMethodParityReady:  ownerReadiness.ReadOnlyMethodParityReady,
			OwnerSmokePlanned:          ownerReadiness.OwnerSmokePlanned,
			ProductionOwnerRoutesReady: ownerReadiness.ProductionOwnerRoutesReady,
			ProductionRecipeTrustReady: ownerReadiness.ProductionRecipeTrustReady,
			LiveDBusOwnerReady:         ownerReadiness.LiveDBusOwnerReady,
			ProductionOwnerEnabled:     ownerReadiness.ProductionOwnerEnabled,
			OwnerTransitionReady:       ownerReadiness.OwnerTransitionReady,
			SystemServiceStarted:       ownerReadiness.SystemServiceStarted,
			ProductionBusClaimed:       ownerReadiness.ProductionBusClaimed,
			WriteMethodsEnabled:        ownerReadiness.WriteMethodsEnabled,
			Counts:                     ownerReadiness.Counts,
		},
		OwnerSmokePlan: RuntimeServiceActivationPreflightSmokePlan{
			RequestType:            ownerSmokePlan.RequestType,
			PlanType:               ownerSmokePlan.PlanType,
			SmokeState:             ownerSmokePlan.SmokeState,
			SmokeEnvironment:       ownerSmokePlan.SmokeEnvironment,
			ActivationBindingReady: ownerSmokePlan.ActivationBindingReady,
			PendingStepCount:       ownerSmokePlan.PendingStepCount,
			Counts:                 ownerSmokePlan.Counts,
		},
		ProductionDBusGate:               productionDBusGate,
		PreflightChecks:                  checks,
		CheckIDs:                         runtimeServiceActivationPreflightCheckIDs(checks),
		Counts:                           counts,
		PreflightDecision:                runtimeServiceActivationPreflightDecision(productionActivationReady, restrictedSmokeReady, counts),
		ProductionActivationReady:        productionActivationReady,
		RestrictedSmokeReady:             restrictedSmokeReady,
		ProductionDBusGateReady:          productionDBusGate.GateReviewReady,
		HumanAuthorizationPreflightReady: productionDBusGate.HumanAuthorizationPreflightReady,
		HumanAuthorizationRequired:       true,
		HumanAuthorizationGranted:        productionDBusGate.HumanAuthorizationGranted,
		AuthorizationReceiptAccepted:     productionDBusGate.AuthorizationReceiptAccepted,
		RuntimeOwned:                     true,
		GoRuntimeBacked:                  true,
		KDEPolicyOwner:                   false,
		KDEMayClaimRuntimeOwnership:      false,
		SystemServiceStarted:             false,
		ProductionBusClaimed:             false,
		WriteMethodsEnabled:              false,
		BackendLaunchEnabled:             false,
		NetworkRequired:                  false,
		HostRootModified:                 false,
		PrivilegedContainerRequired:      false,
		BackendDetailsExposed:            false,
		BlockedReasons:                   runtimeServiceActivationPreflightBlockedReasons(ownerReadiness, productionDBusGate),
		BlockedActions:                   runtimeServiceActivationPreflightBlockedActions(),
		NextRequirements:                 runtimeServiceActivationPreflightNextRequirements(),
		DesktopSafeSummary:               runtimeServiceActivationPreflightSummary(productionActivationReady, restrictedSmokeReady, counts),
	}
	if err := validateNoBackendTerms(preview, "Runtime service activation preflight preview"); err != nil {
		return RuntimeServiceActivationPreflightPreview{}, err
	}
	return preview, nil
}

func runtimeServiceActivationPreflightChecks(ownerReadiness RuntimeOwnerReadinessPreview, productionDBusGate RuntimeServiceActivationProductionDBusGate) []RuntimeServiceActivationPreflightCheck {
	return []RuntimeServiceActivationPreflightCheck{
		runtimeServiceActivationPreflightCheck("activation-binding", runtimeServiceActivationPreflightPassBlocked(ownerReadiness.ActivationBindingReady), "D-Bus activation files, systemd unit, packaged entry point, and contract must be aligned."),
		runtimeServiceActivationPreflightCheck("read-only-method-parity", runtimeServiceActivationPreflightPassBlocked(ownerReadiness.ReadOnlyMethodParityReady), "The Runtime owner must cover every read-only planning method before production activation."),
		runtimeServiceActivationPreflightCheck("owner-smoke-plan", runtimeServiceActivationPreflightPassBlocked(ownerReadiness.OwnerSmokePlanned), "A deterministic restricted-session owner smoke must exist before production activation."),
		runtimeServiceActivationPreflightCheck("production-dbus-gate-review", runtimeServiceActivationPreflightPassBlocked(productionDBusGate.GateReviewReady), "Production service activation must consume the production D-Bus gate review before any production commit."),
		runtimeServiceActivationPreflightCheck("human-authorization-preflight", runtimeServiceActivationPreflightPassBlocked(productionDBusGate.HumanAuthorizationPreflightReady && productionDBusGate.HumanAuthorizationRequired), "Production service activation must consume the human authorization preflight shape."),
		runtimeServiceActivationPreflightCheck("human-authorization-receipt", runtimeServiceActivationPreflightPendingPass(productionDBusGate.AuthorizationReceiptAccepted), "Production service activation remains pending until an explicit authorization receipt is accepted."),
		runtimeServiceActivationPreflightCheck("write-method-gate", runtimeServiceActivationPreflightPassBlocked(!ownerReadiness.WriteMethodsEnabled), "Install, launch, snapshot, restore, repair, and settings persistence must remain disabled during preflight."),
		runtimeServiceActivationPreflightCheck("kde-ownership-boundary", runtimeServiceActivationPreflightPassBlocked(!ownerReadiness.KDEPolicyOwner && !ownerReadiness.KDEMayClaimRuntimeOwnership), "KDE must remain the presentation shell and must not own Runtime policy."),
		runtimeServiceActivationPreflightCheck("host-safety-boundary", "pass", "The preflight must not start services, claim bus names, require network, or mutate the host root."),
		runtimeServiceActivationPreflightCheck("long-running-runtime-owner", runtimeServiceActivationPreflightPendingPass(ownerReadiness.LiveDBusOwnerReady), "A packaged long-running Runtime owner must be proven outside this read-only preview."),
		runtimeServiceActivationPreflightCheck("restricted-owner-smoke", runtimeServiceActivationPreflightPendingPass(ownerReadiness.OwnerTransitionReady), "Restricted owner smoke must prove route coverage, disabled writes, and safe KDE summaries."),
		runtimeServiceActivationPreflightCheck("production-bus-claim", runtimeServiceActivationPreflightPendingPass(ownerReadiness.ProductionBusClaimed), "Production smoke must prove the packaged Runtime owner owns the stable bus name."),
		runtimeServiceActivationPreflightCheck("production-recipe-trust", runtimeServiceActivationPreflightPendingPass(ownerReadiness.ProductionRecipeTrustReady), "Production activation requires digest-verified and production-signed recipes."),
	}
}

func runtimeServiceActivationPreflightCheck(id string, status string, summary string) RuntimeServiceActivationPreflightCheck {
	return RuntimeServiceActivationPreflightCheck{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeServiceActivationPreflightPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeServiceActivationPreflightPendingPass(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func runtimeServiceActivationPreflightCheckIDs(checks []RuntimeServiceActivationPreflightCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRuntimeServiceActivationPreflightChecks(checks []RuntimeServiceActivationPreflightCheck) RuntimeServiceActivationPreflightCounts {
	counts := RuntimeServiceActivationPreflightCounts{Total: len(checks)}
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

func runtimeServiceActivationRestrictedSmokeReady(ownerReadiness RuntimeOwnerReadinessPreview, ownerSmokePlan RuntimeOwnerSmokePlanPreview) bool {
	return ownerReadiness.ActivationBindingReady &&
		ownerReadiness.ReadOnlyMethodParityReady &&
		ownerReadiness.OwnerSmokePlanned &&
		ownerSmokePlan.SmokeState == "planned" &&
		!ownerReadiness.WriteMethodsEnabled &&
		!ownerReadiness.KDEPolicyOwner &&
		!ownerReadiness.KDEMayClaimRuntimeOwnership &&
		!ownerReadiness.HostRootModified &&
		!ownerReadiness.NetworkRequired &&
		!ownerReadiness.PrivilegedContainerRequired
}

func runtimeServiceActivationProductionReady(ownerReadiness RuntimeOwnerReadinessPreview, productionDBusGate RuntimeServiceActivationProductionDBusGate) bool {
	return ownerReadiness.ActivationBindingReady &&
		ownerReadiness.ReadOnlyMethodParityReady &&
		ownerReadiness.OwnerSmokePlanned &&
		productionDBusGate.GateReviewReady &&
		productionDBusGate.HumanAuthorizationPreflightReady &&
		productionDBusGate.HumanAuthorizationGranted &&
		productionDBusGate.AuthorizationReceiptAccepted &&
		ownerReadiness.LiveDBusOwnerReady &&
		ownerReadiness.ProductionOwnerEnabled &&
		ownerReadiness.OwnerTransitionReady &&
		ownerReadiness.ProductionOwnerRoutesReady &&
		ownerReadiness.ProductionRecipeTrustReady &&
		!ownerReadiness.SystemServiceStarted &&
		!ownerReadiness.ProductionBusClaimed &&
		!ownerReadiness.WriteMethodsEnabled &&
		!ownerReadiness.HostRootModified &&
		!ownerReadiness.NetworkRequired &&
		!ownerReadiness.PrivilegedContainerRequired &&
		productionDBusGate.ProductionReadiness &&
		!productionDBusGate.ProductionOwnerEnabled &&
		!productionDBusGate.ProductionActivationReady &&
		!productionDBusGate.SystemServiceStarted &&
		!productionDBusGate.SessionBusClaimed &&
		!productionDBusGate.ProductionBusClaimed &&
		!productionDBusGate.WriteMethodsEnabled &&
		!productionDBusGate.RuntimeWritesEnabled &&
		!productionDBusGate.BackendLaunchEnabled &&
		!productionDBusGate.HostRootModified
}

func runtimeServiceActivationPreflightDecision(productionReady bool, restrictedSmokeReady bool, counts RuntimeServiceActivationPreflightCounts) string {
	if productionReady {
		return "production-activation-ready-for-human-commit"
	}
	if counts.Blocked > 0 {
		return "production-activation-blocked"
	}
	if restrictedSmokeReady {
		return "restricted-owner-smoke-ready"
	}
	return "production-activation-pending-evidence"
}

func runtimeServiceActivationPreflightBlockedReasons(ownerReadiness RuntimeOwnerReadinessPreview, productionDBusGate RuntimeServiceActivationProductionDBusGate) []string {
	reasons := []string{}
	if !ownerReadiness.ActivationBindingReady {
		reasons = append(reasons, "Runtime activation files are not aligned.")
	}
	if !ownerReadiness.ReadOnlyMethodParityReady {
		reasons = append(reasons, "Read-only Runtime method parity is incomplete.")
	}
	if !ownerReadiness.LiveDBusOwnerReady {
		reasons = append(reasons, "Live production Runtime ownership is not proven.")
	}
	if !ownerReadiness.OwnerTransitionReady {
		reasons = append(reasons, "Restricted owner transition smoke has not passed.")
	}
	if !ownerReadiness.ProductionRecipeTrustReady {
		reasons = append(reasons, "Production recipe trust is not ready.")
	}
	if !ownerReadiness.ProductionBusClaimed {
		reasons = append(reasons, "Production bus-name ownership is not proven.")
	}
	if !productionDBusGate.GateReviewReady {
		reasons = append(reasons, "Production D-Bus gate review is not ready.")
	}
	if !productionDBusGate.AuthorizationReceiptAccepted {
		reasons = append(reasons, "Production D-Bus human authorization receipt has not been accepted.")
	}
	return reasons
}

func runtimeServiceActivationPreflightBlockedActions() []string {
	return []string{
		"start production Runtime service from preflight",
		"claim production D-Bus name from preflight",
		"treat production D-Bus gate review as an activation commit",
		"accept production D-Bus ownership without a human authorization receipt",
		"enable launch, install, snapshot, restore, repair, or settings persistence",
		"let KDE claim Runtime ownership",
		"write KDE or system activation files from preflight",
		"mutate host root during service activation preflight",
	}
}

func runtimeServiceActivationPreflightNextRequirements() []string {
	return []string{
		"Run the restricted owner smoke against a packaged Runtime owner in an isolated session.",
		"Keep production service activation consuming production D-Bus gate review and human authorization preflight.",
		"Accept an explicit human authorization receipt before production service activation can be committed.",
		"Prove stable bus-name ownership with the packaged Runtime owner, not the smoke adapter.",
		"Require production-signed recipe trust before any production activation commit.",
		"Keep service activation commit behind explicit human authorization and package-managed writes.",
	}
}

func runtimeServiceActivationPreflightSummary(productionReady bool, restrictedSmokeReady bool, counts RuntimeServiceActivationPreflightCounts) string {
	if productionReady {
		return "Production Runtime service activation is ready for explicit human commit; the preview still performs no activation."
	}
	if counts.Blocked > 0 {
		return "Production Runtime service activation is blocked until activation, method parity, and safety defects are repaired."
	}
	if restrictedSmokeReady {
		return "Runtime service activation consumes production D-Bus gate preflights and is ready for restricted owner smoke; production activation, bus claim, and service start remain disabled."
	}
	return "Runtime service activation still needs owner evidence before production activation can be considered."
}

func runtimeServiceActivationProductionDBusGate(root string) RuntimeServiceActivationProductionDBusGate {
	gateSource := readRuntimeServiceBindingOptional(root, "internal/runtime/owner/production_dbus_gate_review.go")
	authorizationSource := readRuntimeServiceBindingOptional(root, "internal/runtime/owner/production_dbus_human_authorization_preflight.go")
	gateSourcePresent := runtimeServiceActivationSourceHasAll(gateSource, []string{
		"xnix.runtime.production_dbus_gate_review.v1",
		"production-dbus-gate-review-preview",
		"production-dbus-gate-review-ready",
		"human-authorization-preflight-present",
		"human-authorization-required",
		"production-dbus-disabled",
		"route-production-exposure-disabled",
		"writes-and-launch-disabled",
	})
	authorizationPreflightPresent := runtimeServiceActivationSourceHasAll(authorizationSource, []string{
		"xnix.runtime.production_dbus_human_authorization_preflight.v1",
		"production-dbus-human-authorization-preflight-preview",
		"read-only-production-dbus-human-authorization-preflight",
		"production-dbus-human-authorization-receipt-id",
		"AuthorizationReceiptRequired",
		"AuthorizationReceiptPresent",
		"AuthorizationGrantReady",
		"AuthorizationAccepted",
		"authorization-not-granted",
		"unsafe-gates-closed",
	})
	gateReady := gateSourcePresent && authorizationPreflightPresent

	return RuntimeServiceActivationProductionDBusGate{
		RequestType:                        "production-dbus-gate-review-preview",
		GateType:                           "owner-local-smoke-covered-production-dbus-gate-review",
		GateDecision:                       runtimeServiceActivationProductionDBusGateDecision(gateReady),
		GateReviewSourcePresent:            gateSourcePresent,
		HumanAuthorizationPreflightPresent: authorizationPreflightPresent,
		GateReviewReady:                    gateReady,
		HumanAuthorizationPreflightReady:   gateReady,
		HumanAuthorizationRequired:         true,
		HumanAuthorizationGranted:          false,
		AuthorizationReceiptAccepted:       false,
		ProductionReadiness:                false,
		ProductionOwnerEnabled:             false,
		ProductionActivationReady:          false,
		SystemServiceStarted:               false,
		SessionBusClaimed:                  false,
		ProductionBusClaimed:               false,
		WriteMethodsEnabled:                false,
		RuntimeWritesEnabled:               false,
		BackendLaunchEnabled:               false,
		HostRootModified:                   false,
	}
}

func runtimeServiceActivationProductionDBusGateDecision(gateReady bool) string {
	if gateReady {
		return "production-dbus-gate-review-consumed-activation-still-blocked"
	}
	return "production-dbus-gate-review-missing"
}

func runtimeServiceActivationSourceHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}
