package appidentity

import "strings"

type KDETestLaunchMaterializationOwnerRouteAuditPreview struct {
	Version                                 string                                             `json:"version"`
	SchemaVersion                           string                                             `json:"schema_version"`
	RequestType                             string                                             `json:"request_type"`
	AuditType                               string                                             `json:"audit_type"`
	Source                                  string                                             `json:"source"`
	RuntimeMethod                           string                                             `json:"runtime_method"`
	ReadMethod                              string                                             `json:"read_method"`
	SubjectRequestType                      string                                             `json:"subject_request_type"`
	SubjectCommand                          string                                             `json:"subject_command"`
	ProposedOwnerMethod                     string                                             `json:"proposed_owner_method"`
	RouteDecision                           string                                             `json:"route_decision"`
	RouteDecisionReason                     string                                             `json:"route_decision_reason"`
	CurrentRouteStatus                      string                                             `json:"current_route_status"`
	CLICommandRegistered                    bool                                               `json:"cli_command_registered"`
	GoReadModelPresent                      bool                                               `json:"go_read_model_present"`
	OwnerDispatchRoutePresent               bool                                               `json:"owner_dispatch_route_present"`
	ProductionDBusMethodPresent             bool                                               `json:"production_dbus_method_present"`
	RequiresCallerRegistryPath              bool                                               `json:"requires_caller_registry_path"`
	RequiresCallerApplicationID             bool                                               `json:"requires_caller_application_id"`
	RequiresCallerStateRoot                 bool                                               `json:"requires_caller_state_root"`
	RequiresExplicitAuthorization           bool                                               `json:"requires_explicit_authorization"`
	MaterializationWritesStateRoot          bool                                               `json:"materialization_writes_state_root"`
	ReadOnlyConsumeCommandRegistered        bool                                               `json:"read_only_consume_command_registered"`
	ReadOnlyReceiptConsumptionReady         bool                                               `json:"read_only_receipt_consumption_ready"`
	ReadOnlyConsumeRequiresCallerStateRoot  bool                                               `json:"read_only_consume_requires_caller_state_root"`
	OwnerManagedOpaqueReceiptLookupReady    bool                                               `json:"owner_managed_opaque_receipt_lookup_ready"`
	OpaqueMaterializationReceiptIDSupported bool                                               `json:"opaque_materialization_receipt_id_supported"`
	FanOutWritesEnabled                     bool                                               `json:"fan_out_writes_enabled"`
	OwnerLocalRouteCandidateReady           bool                                               `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady             bool                                               `json:"production_dbus_exposure_ready"`
	RecommendedNextRoute                    string                                             `json:"recommended_next_route"`
	Checks                                  []KDETestLaunchMaterializationOwnerRouteAuditCheck `json:"checks"`
	CheckIDs                                []string                                           `json:"check_ids"`
	Counts                                  KDETestLaunchMaterializationOwnerRouteAuditCounts  `json:"counts"`
	RuntimeOwned                            bool                                               `json:"runtime_owned"`
	GoRuntimeBacked                         bool                                               `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                                               `json:"kde_policy_owner"`
	SystemServiceStarted                    bool                                               `json:"system_service_started"`
	SessionBusClaimed                       bool                                               `json:"session_bus_claimed"`
	ProductionBusClaimed                    bool                                               `json:"production_bus_claimed"`
	WriteMethodsEnabled                     bool                                               `json:"write_methods_enabled"`
	RuntimeWritesEnabled                    bool                                               `json:"runtime_writes_enabled"`
	BackendLaunchEnabled                    bool                                               `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                               `json:"backend_process_started"`
	NetworkRequired                         bool                                               `json:"network_required"`
	HostRootModified                        bool                                               `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                               `json:"privileged_container_required"`
	StateRootPathExposed                    bool                                               `json:"state_root_path_exposed"`
	RawCommandExposed                       bool                                               `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                               `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                               `json:"backend_details_exposed"`
	BlockedActions                          []string                                           `json:"blocked_actions"`
	NextRequirements                        []string                                           `json:"next_requirements"`
	DesktopSafeSummary                      string                                             `json:"desktop_safe_summary"`
}

type KDETestLaunchMaterializationOwnerRouteAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type KDETestLaunchMaterializationOwnerRouteAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewKDETestLaunchMaterializationOwnerRouteAuditPreview(root string) (KDETestLaunchMaterializationOwnerRouteAuditPreview, error) {
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return KDETestLaunchMaterializationOwnerRouteAuditPreview{}, err
	}
	sources := kdeTestLaunchMaterializationOwnerRouteAuditSources(root)
	audit := KDETestLaunchMaterializationOwnerRouteAuditPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.kde_test_launch_materialization_owner_route_audit.v1",
		RequestType:                             "kde-test-launch-materialization-owner-route-audit-preview",
		AuditType:                               "materialization-fanout-owner-route-audit",
		Source:                                  "kde-test-launch-materialization-fanout-preview+runtime-owner-dispatch+dbus-contract",
		RuntimeMethod:                           "GetKDETestLaunchMaterializationOwnerRouteAudit",
		ReadMethod:                              "GetKDETestLaunchMaterializationOwnerRouteAuditPreview",
		SubjectRequestType:                      "kde-test-launch-materialization-fanout-preview",
		SubjectCommand:                          "kde-test-launch-materialization-fanout-preview",
		ProposedOwnerMethod:                     "GetKDETestLaunchMaterializationFanOut",
		RouteDecision:                           "materialization-owner-route-sources-missing",
		RouteDecisionReason:                     "The materialization fan-out owner route remains blocked until required read-only fan-out, owner lookup, and owner-route evidence is present.",
		CurrentRouteStatus:                      "fail-closed-owner-route-audit",
		CLICommandRegistered:                    kdeTestLaunchMaterializationAuditHasCLICommand(sources.GoCLI),
		GoReadModelPresent:                      kdeTestLaunchMaterializationAuditHasReadModel(sources.GoReadModel),
		OwnerDispatchRoutePresent:               kdeTestLaunchMaterializationAuditHasOwnerRoute(sources.OwnerDispatch),
		ProductionDBusMethodPresent:             kdeTestLaunchMaterializationAuditHasProductionDBusMethod(sources.DBusContract),
		RequiresCallerRegistryPath:              textHasAll(sources.CLICommand, []string{"--registry", "registryPath"}),
		RequiresCallerApplicationID:             textHasAll(sources.CLICommand, []string{"--app", "applicationID"}),
		RequiresCallerStateRoot:                 textHasAll(sources.CLICommand, []string{"--state-root", "stateRoot"}),
		RequiresExplicitAuthorization:           textHasAll(sources.CLICommand, []string{"--mode", "--authorize", "authorize-restricted-test-preparation"}),
		MaterializationWritesStateRoot:          textHasAll(sources.GoReadModel, []string{"NewKDETestLaunchMaterializationRecord", "StateRootWritesEnabled"}),
		ReadOnlyConsumeCommandRegistered:        textHasAll(sources.CLICommand, []string{"kde-test-launch-materialization-fanout-consume-preview", "runKDETestLaunchMaterializationFanOutConsumePreview", "--materialization-plan-id"}),
		ReadOnlyReceiptConsumptionReady:         textHasAll(sources.GoReadModel, []string{"NewKDETestLaunchMaterializationFanOutPreviewFromReceipt", "LoadKDETestLaunchMaterializationRecordForFanOut", "read-only-existing-materialization-receipt", "read-only-receipt-consumption"}),
		ReadOnlyConsumeRequiresCallerStateRoot:  textHasAll(sources.CLICommand, []string{"kde-test-launch-materialization-fanout-consume-preview", "--state-root", "materializationPlanID"}),
		OwnerManagedOpaqueReceiptLookupReady:    textHasAll(sources.OwnerDispatch, []string{"ResolveKDETestLaunchMaterializationReceipt", "owner-managed-materialization-receipt-lookup"}),
		OpaqueMaterializationReceiptIDSupported: textHasAll(sources.OwnerDispatch, []string{"opaque_materialization_receipt_id", "kde-test-launch-materialization-receipt-id"}),
		FanOutWritesEnabled:                     false,
		OwnerLocalRouteCandidateReady:           false,
		ProductionDBusExposureReady:             false,
		RecommendedNextRoute:                    "materialization-fanout-owner-local-read-route",
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		SystemServiceStarted:                    false,
		SessionBusClaimed:                       false,
		ProductionBusClaimed:                    false,
		WriteMethodsEnabled:                     false,
		RuntimeWritesEnabled:                    false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		NetworkRequired:                         false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		StateRootPathExposed:                    false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions:                          kdeTestLaunchMaterializationOwnerRouteAuditBlockedActions(),
		NextRequirements:                        kdeTestLaunchMaterializationOwnerRouteAuditNextRequirements(),
		DesktopSafeSummary:                      "The materialization fan-out owner-route audit is fail-closed until the required read-only fan-out, owner lookup, and owner-route evidence exists without production D-Bus exposure, Runtime writes, launch, or host mutation.",
	}
	kdeTestLaunchMaterializationConfigureOwnerRouteDecision(&audit)
	checks := kdeTestLaunchMaterializationOwnerRouteAuditChecks(audit)
	audit.Checks = checks
	audit.CheckIDs = kdeTestLaunchMaterializationOwnerRouteAuditCheckIDs(checks)
	audit.Counts = countKDETestLaunchMaterializationOwnerRouteAuditChecks(checks)
	if err := validateNoBackendTerms(audit, "KDE test launch materialization owner-route audit preview"); err != nil {
		return KDETestLaunchMaterializationOwnerRouteAuditPreview{}, err
	}
	return audit, nil
}

type kdeTestLaunchMaterializationOwnerRouteAuditSourceSet struct {
	GoCLI         string
	CLICommand    string
	GoReadModel   string
	OwnerDispatch string
	DBusContract  string
}

func kdeTestLaunchMaterializationOwnerRouteAuditSources(root string) kdeTestLaunchMaterializationOwnerRouteAuditSourceSet {
	return kdeTestLaunchMaterializationOwnerRouteAuditSourceSet{
		GoCLI:         readRuntimeMethodParitySources(root, []string{"cmd/xnix-runtime-go/main.go"}),
		CLICommand:    readRuntimeMethodParitySources(root, []string{"cmd/xnix-runtime-go/kde_test_launch_materialization_fanout_commands.go"}),
		GoReadModel:   readRuntimeMethodParitySources(root, []string{"internal/runtime/appidentity/kde_test_launch_materialization_fanout.go"}),
		OwnerDispatch: readRuntimeMethodParitySources(root, []string{"internal/runtime/owner/dispatch.go", "internal/runtime/appidentity/kde_test_launch_materialization_receipt_lookup.go"}),
		DBusContract:  readRuntimeMethodParitySources(root, []string{"runtime/dbus/org.xnix.Compatibility1.xml"}),
	}
}

func kdeTestLaunchMaterializationOwnerRouteAuditChecks(audit KDETestLaunchMaterializationOwnerRouteAuditPreview) []KDETestLaunchMaterializationOwnerRouteAuditCheck {
	return []KDETestLaunchMaterializationOwnerRouteAuditCheck{
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("cli-preview-registered", kdeTestLaunchMaterializationAuditPassBlocked(audit.CLICommandRegistered), "The materialization fan-out preview is registered as a Go CLI preview."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("go-read-model-present", kdeTestLaunchMaterializationAuditPassBlocked(audit.GoReadModelPresent), "The Go Runtime materialization fan-out read model exists."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("read-only-consume-registered", kdeTestLaunchMaterializationAuditPassBlocked(audit.ReadOnlyConsumeCommandRegistered && audit.ReadOnlyReceiptConsumptionReady), "The read-only materialization receipt consumption split is registered."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("owner-route-absent", kdeTestLaunchMaterializationAuditPassBlocked(!audit.OwnerDispatchRoutePresent), "The fan-out is not yet accepted by the owner dispatch table."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("production-dbus-absent", kdeTestLaunchMaterializationAuditPassBlocked(!audit.ProductionDBusMethodPresent), "The fan-out is not exposed as a production D-Bus method."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("caller-path-boundary", kdeTestLaunchMaterializationAuditPendingUnless(!audit.RequiresCallerRegistryPath && !audit.RequiresCallerStateRoot && !audit.ReadOnlyConsumeRequiresCallerStateRoot), "Owner-local routes must not depend on caller-supplied registry or state-root paths."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("receipt-creation-split", kdeTestLaunchMaterializationAuditPassBlocked(audit.ReadOnlyReceiptConsumptionReady && audit.MaterializationWritesStateRoot), "The old test receipt creation path is split from the read-only consumption path."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("owner-managed-opaque-lookup", kdeTestLaunchMaterializationAuditPendingUnless(audit.OwnerManagedOpaqueReceiptLookupReady && audit.OpaqueMaterializationReceiptIDSupported), "Owner-local routing needs opaque materialization receipt identifiers resolved inside the Runtime owner."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("route-decision", kdeTestLaunchMaterializationAuditPassBlocked(kdeTestLaunchMaterializationOwnerRouteDecisionClosed(audit)), "The audit keeps the fan-out CLI-only until owner-local fan-out routing can consume owner-managed opaque receipt lookup."),
		kdeTestLaunchMaterializationOwnerRouteAuditCheck("unsafe-gates-closed", kdeTestLaunchMaterializationAuditPassBlocked(!audit.SystemServiceStarted && !audit.SessionBusClaimed && !audit.ProductionBusClaimed && !audit.WriteMethodsEnabled && !audit.BackendLaunchEnabled && !audit.HostRootModified && !audit.StateRootPathExposed && !audit.BackendDetailsExposed), "The audit does not start services, claim buses, enable writes, launch, expose paths, or mutate the host."),
	}
}

func kdeTestLaunchMaterializationConfigureOwnerRouteDecision(audit *KDETestLaunchMaterializationOwnerRouteAuditPreview) {
	if !audit.CLICommandRegistered || !audit.GoReadModelPresent || !audit.ReadOnlyReceiptConsumptionReady {
		audit.RouteDecision = "materialization-owner-route-sources-missing"
		audit.RouteDecisionReason = "The materialization fan-out owner route remains blocked until required CLI, Go read-model, and read-only receipt-consumption evidence is present."
		audit.CurrentRouteStatus = "fail-closed-owner-route-audit"
		audit.RecommendedNextRoute = "restore-materialization-fanout-route-sources"
		audit.DesktopSafeSummary = "The materialization fan-out owner-route audit is fail-closed because required local source evidence is missing; production D-Bus exposure, Runtime writes, launch, and host mutation remain disabled."
		return
	}
	if audit.OwnerManagedOpaqueReceiptLookupReady && audit.OpaqueMaterializationReceiptIDSupported {
		audit.RouteDecision = "opaque-lookup-ready-owner-route-blocked"
		audit.RouteDecisionReason = "Owner-managed opaque materialization receipt lookup exists, so the remaining blocker is an owner-local read route that consumes lookup results without caller registry or state-root paths."
		audit.CurrentRouteStatus = "owner-managed-lookup-ready-fanout-cli-only"
		audit.RecommendedNextRoute = "materialization-fanout-owner-local-read-route"
		audit.DesktopSafeSummary = "The test-only materialization fan-out has read-only receipt consumption and owner-managed opaque receipt lookup, but it remains CLI-only until an owner-local fan-out route consumes that lookup without production D-Bus exposure, Runtime writes, launch, or host mutation."
		return
	}
	audit.RouteDecision = "consume-ready-opaque-lookup-missing"
	audit.RouteDecisionReason = "The read-only materialization fan-out consumption split exists, but owner-local routing still needs owner-managed opaque materialization receipt lookup without caller registry or state-root paths."
	audit.CurrentRouteStatus = "read-only-consume-ready-owner-route-blocked"
	audit.RecommendedNextRoute = "owner-managed-materialization-receipt-lookup"
	audit.DesktopSafeSummary = "The test-only materialization fan-out has a read-only receipt-consumption path, but it remains CLI-only until the Runtime owner can resolve opaque materialization receipt evidence without caller paths, production D-Bus exposure, Runtime writes, launch, or host mutation."
}

func kdeTestLaunchMaterializationOwnerRouteDecisionClosed(audit KDETestLaunchMaterializationOwnerRouteAuditPreview) bool {
	if audit.RouteDecision == "materialization-owner-route-sources-missing" {
		return false
	}
	if audit.RouteDecision == "consume-ready-opaque-lookup-missing" {
		return audit.ReadOnlyReceiptConsumptionReady &&
			!audit.OwnerManagedOpaqueReceiptLookupReady &&
			!audit.OwnerLocalRouteCandidateReady &&
			!audit.ProductionDBusExposureReady
	}
	if audit.RouteDecision == "opaque-lookup-ready-owner-route-blocked" {
		return audit.ReadOnlyReceiptConsumptionReady &&
			audit.OwnerManagedOpaqueReceiptLookupReady &&
			audit.OpaqueMaterializationReceiptIDSupported &&
			!audit.OwnerLocalRouteCandidateReady &&
			!audit.ProductionDBusExposureReady
	}
	return false
}

func kdeTestLaunchMaterializationOwnerRouteAuditCheck(id string, status string, summary string) KDETestLaunchMaterializationOwnerRouteAuditCheck {
	return KDETestLaunchMaterializationOwnerRouteAuditCheck{ID: id, Status: status, Summary: summary}
}

func kdeTestLaunchMaterializationAuditPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func kdeTestLaunchMaterializationAuditPendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func kdeTestLaunchMaterializationOwnerRouteAuditCheckIDs(checks []KDETestLaunchMaterializationOwnerRouteAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countKDETestLaunchMaterializationOwnerRouteAuditChecks(checks []KDETestLaunchMaterializationOwnerRouteAuditCheck) KDETestLaunchMaterializationOwnerRouteAuditCounts {
	counts := KDETestLaunchMaterializationOwnerRouteAuditCounts{Total: len(checks)}
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

func kdeTestLaunchMaterializationAuditHasCLICommand(source string) bool {
	return textHasAll(source, []string{"kde-test-launch-materialization-fanout-preview", "runKDETestLaunchMaterializationFanOutPreview"})
}

func kdeTestLaunchMaterializationAuditHasReadModel(source string) bool {
	return textHasAll(source, []string{"KDETestLaunchMaterializationFanOutPreview", "NewKDETestLaunchMaterializationFanOutPreview", "xnix.runtime.kde_test_launch_materialization_fanout.v1"})
}

func kdeTestLaunchMaterializationAuditHasOwnerRoute(source string) bool {
	return textHasAll(source, []string{"GetKDETestLaunchMaterializationFanOut", "NewKDETestLaunchMaterializationFanOutPreview"})
}

func kdeTestLaunchMaterializationAuditHasProductionDBusMethod(source string) bool {
	return textHasAll(source, []string{"GetKDETestLaunchMaterializationFanOut", "kde-test-launch-materialization-fanout-preview"})
}

func kdeTestLaunchMaterializationOwnerRouteAuditBlockedActions() []string {
	return []string{
		"add materialization fan-out to production D-Bus before owner-local evidence consumption exists",
		"accept caller-supplied state-root paths through owner dispatch",
		"materialize test receipts from an owner-local read route",
		"accept materialization plan ids as caller-owned state-root lookups",
		"enable Runtime writes from materialization fan-out audit",
		"launch compatibility backend from materialization fan-out audit",
		"mutate host root during materialization fan-out audit",
	}
}

func kdeTestLaunchMaterializationOwnerRouteAuditNextRequirements() []string {
	return []string{
		"Introduce owner-managed opaque materialization receipt identifiers before owner-local routing.",
		"Resolve materialization receipts inside the Runtime owner without caller registry or state-root paths.",
		"Keep production D-Bus exposure blocked until owner smoke covers the read-only route.",
		"Require route-manifest and service-call evidence before promoting beyond CLI preview.",
	}
}

func textHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}
