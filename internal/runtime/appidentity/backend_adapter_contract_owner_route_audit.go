package appidentity

type BackendAdapterContractOwnerRouteAuditPreview struct {
	Version                        string                                       `json:"version"`
	SchemaVersion                  string                                       `json:"schema_version"`
	RequestType                    string                                       `json:"request_type"`
	AuditType                      string                                       `json:"audit_type"`
	Source                         string                                       `json:"source"`
	RuntimeMethod                  string                                       `json:"runtime_method"`
	ReadMethod                     string                                       `json:"read_method"`
	SubjectRequestType             string                                       `json:"subject_request_type"`
	SubjectCommand                 string                                       `json:"subject_command"`
	ProposedOwnerMethod            string                                       `json:"proposed_owner_method"`
	RouteDecision                  string                                       `json:"route_decision"`
	RouteDecisionReason            string                                       `json:"route_decision_reason"`
	CurrentRouteStatus             string                                       `json:"current_route_status"`
	CLICommandRegistered           bool                                         `json:"cli_command_registered"`
	GoReadModelPresent             bool                                         `json:"go_read_model_present"`
	FixtureMatrixConsumesContract  bool                                         `json:"fixture_matrix_consumes_contract"`
	OwnerDispatchRoutePresent      bool                                         `json:"owner_dispatch_route_present"`
	ProductionDBusMethodPresent    bool                                         `json:"production_dbus_method_present"`
	FullContractContainsAdapterIDs bool                                         `json:"full_contract_contains_adapter_ids"`
	KDEFacingProjectionPresent     bool                                         `json:"kde_facing_projection_present"`
	RedactedProfileRoutePresent    bool                                         `json:"redacted_profile_route_present"`
	RequiresCallerRoot             bool                                         `json:"requires_caller_root"`
	OwnerSmokeCoverageReady        bool                                         `json:"owner_smoke_coverage_ready"`
	AdapterInvocationEnabled       bool                                         `json:"adapter_invocation_enabled"`
	BackendInstallEnabled          bool                                         `json:"backend_install_enabled"`
	BackendDownloadEnabled         bool                                         `json:"backend_download_enabled"`
	BackendLaunchEnabled           bool                                         `json:"backend_launch_enabled"`
	BackendProcessStarted          bool                                         `json:"backend_process_started"`
	CommandMaterialized            bool                                         `json:"command_materialized"`
	ExecutablePathResolved         bool                                         `json:"executable_path_resolved"`
	OwnerLocalRouteCandidateReady  bool                                         `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady    bool                                         `json:"production_dbus_exposure_ready"`
	RecommendedNextRoute           string                                       `json:"recommended_next_route"`
	Checks                         []BackendAdapterContractOwnerRouteAuditCheck `json:"checks"`
	CheckIDs                       []string                                     `json:"check_ids"`
	Counts                         BackendAdapterContractOwnerRouteAuditCounts  `json:"counts"`
	RuntimeOwned                   bool                                         `json:"runtime_owned"`
	GoRuntimeBacked                bool                                         `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                                         `json:"kde_policy_owner"`
	SystemServiceStarted           bool                                         `json:"system_service_started"`
	SessionBusClaimed              bool                                         `json:"session_bus_claimed"`
	ProductionBusClaimed           bool                                         `json:"production_bus_claimed"`
	WriteMethodsEnabled            bool                                         `json:"write_methods_enabled"`
	RuntimeWritesEnabled           bool                                         `json:"runtime_writes_enabled"`
	NetworkRequired                bool                                         `json:"network_required"`
	HostRootModified               bool                                         `json:"host_root_modified"`
	PrivilegedContainerRequired    bool                                         `json:"privileged_container_required"`
	StateRootPathExposed           bool                                         `json:"state_root_path_exposed"`
	RawCommandExposed              bool                                         `json:"raw_command_exposed"`
	RawExecutableExposed           bool                                         `json:"raw_executable_exposed"`
	BackendDetailsExposed          bool                                         `json:"backend_details_exposed"`
	BlockedActions                 []string                                     `json:"blocked_actions"`
	NextRequirements               []string                                     `json:"next_requirements"`
	DesktopSafeSummary             string                                       `json:"desktop_safe_summary"`
}

type BackendAdapterContractOwnerRouteAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type BackendAdapterContractOwnerRouteAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewBackendAdapterContractOwnerRouteAuditPreview(root string) (BackendAdapterContractOwnerRouteAuditPreview, error) {
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return BackendAdapterContractOwnerRouteAuditPreview{}, err
	}
	sources := backendAdapterContractOwnerRouteAuditSources(root)
	audit := BackendAdapterContractOwnerRouteAuditPreview{
		Version:                        version,
		SchemaVersion:                  "xnix.runtime.backend_adapter_contract_owner_route_audit.v1",
		RequestType:                    "backend-adapter-contract-owner-route-audit-preview",
		AuditType:                      "adapter-contract-owner-route-audit",
		Source:                         "backend-adapter-contract-preview+offline-application-fixture-matrix+runtime-owner-dispatch+owner-smoke-coverage+dbus-contract",
		RuntimeMethod:                  "GetBackendAdapterContractOwnerRouteAudit",
		ReadMethod:                     "GetBackendAdapterContractOwnerRouteAuditPreview",
		SubjectRequestType:             "backend-adapter-contract-preview",
		SubjectCommand:                 "backend-adapter-contract-preview",
		ProposedOwnerMethod:            "GetBackendAdapterProfileAudit",
		RouteDecision:                  "redacted-profile-route-ready",
		RouteDecisionReason:            "The no-op adapter contract remains fixture-local, while owner-local routing can expose only the redacted profile audit route.",
		CurrentRouteStatus:             "redacted-profile-route-ready-full-contract-fixture-local",
		CLICommandRegistered:           backendAdapterContractAuditHasCLICommand(sources.GoCLI),
		GoReadModelPresent:             backendAdapterContractAuditHasReadModel(sources.GoReadModel),
		FixtureMatrixConsumesContract:  backendAdapterContractAuditHasFixtureConsumption(sources.FixtureMatrix),
		OwnerDispatchRoutePresent:      backendAdapterContractAuditHasOwnerRoute(sources.OwnerDispatch),
		ProductionDBusMethodPresent:    backendAdapterContractAuditHasProductionDBusMethod(sources.DBusContract),
		FullContractContainsAdapterIDs: textHasAll(sources.GoReadModel, []string{"AdapterIDs", "backendAdapterContract("}),
		KDEFacingProjectionPresent:     textHasAll(sources.GoReadModel, []string{"KDEFacingProfiles", "backendAdapterProfile("}),
		RedactedProfileRoutePresent:    backendAdapterContractAuditHasRedactedRoute(sources.OwnerDispatch),
		RequiresCallerRoot:             backendAdapterContractAuditHasRootFlag(sources.CLICommand),
		OwnerSmokeCoverageReady:        backendAdapterContractAuditHasOwnerSmokeCoverage(sources.OwnerSmokeCoverage),
		AdapterInvocationEnabled:       false,
		BackendInstallEnabled:          false,
		BackendDownloadEnabled:         false,
		BackendLaunchEnabled:           false,
		BackendProcessStarted:          false,
		CommandMaterialized:            false,
		ExecutablePathResolved:         false,
		OwnerLocalRouteCandidateReady:  backendAdapterContractAuditHasRedactedRoute(sources.OwnerDispatch),
		ProductionDBusExposureReady:    false,
		RecommendedNextRoute:           "restricted-owner-smoke-redacted-adapter-profile-audit",
		RuntimeOwned:                   true,
		GoRuntimeBacked:                true,
		KDEPolicyOwner:                 false,
		SystemServiceStarted:           false,
		SessionBusClaimed:              false,
		ProductionBusClaimed:           false,
		WriteMethodsEnabled:            false,
		RuntimeWritesEnabled:           false,
		NetworkRequired:                false,
		HostRootModified:               false,
		PrivilegedContainerRequired:    false,
		StateRootPathExposed:           false,
		RawCommandExposed:              false,
		RawExecutableExposed:           false,
		BackendDetailsExposed:          false,
		BlockedActions:                 backendAdapterContractOwnerRouteAuditBlockedActions(),
		NextRequirements:               backendAdapterContractOwnerRouteAuditNextRequirements(),
		DesktopSafeSummary:             "The no-op adapter contract remains fixture-local while a redacted owner-local profile audit exposes user-safe compatibility modes without internal adapter identifiers, production D-Bus exposure, Runtime writes, launch, or host mutation.",
	}
	backendAdapterContractConfigureOwnerRouteDecision(&audit)
	checks := backendAdapterContractOwnerRouteAuditChecks(audit)
	audit.Checks = checks
	audit.CheckIDs = backendAdapterContractOwnerRouteAuditCheckIDs(checks)
	audit.Counts = countBackendAdapterContractOwnerRouteAuditChecks(checks)
	if err := validateNoBackendTerms(audit, "backend adapter contract owner-route audit preview"); err != nil {
		return BackendAdapterContractOwnerRouteAuditPreview{}, err
	}
	return audit, nil
}

type backendAdapterContractOwnerRouteAuditSourceSet struct {
	GoCLI              string
	CLICommand         string
	GoReadModel        string
	FixtureMatrix      string
	OwnerDispatch      string
	OwnerSmokeCoverage string
	DBusContract       string
}

func backendAdapterContractOwnerRouteAuditSources(root string) backendAdapterContractOwnerRouteAuditSourceSet {
	return backendAdapterContractOwnerRouteAuditSourceSet{
		GoCLI:              readRuntimeMethodParitySources(root, []string{"cmd/xnix-runtime-go/main.go"}),
		CLICommand:         readRuntimeMethodParitySources(root, []string{"cmd/xnix-runtime-go/backend_adapter_contract_commands.go"}),
		GoReadModel:        readRuntimeMethodParitySources(root, []string{"internal/runtime/appidentity/backend_adapter_contract.go"}),
		FixtureMatrix:      readRuntimeMethodParitySources(root, []string{"internal/runtime/appidentity/offline_application_fixture_matrix.go"}),
		OwnerDispatch:      readRuntimeMethodParitySources(root, []string{"internal/runtime/owner/dispatch.go"}),
		OwnerSmokeCoverage: readRuntimeMethodParitySources(root, []string{"internal/runtime/owner/backend_adapter_redacted_profile_owner_smoke_coverage.go"}),
		DBusContract:       readRuntimeMethodParitySources(root, []string{"runtime/dbus/org.xnix.Compatibility1.xml"}),
	}
}

func backendAdapterContractOwnerRouteAuditChecks(audit BackendAdapterContractOwnerRouteAuditPreview) []BackendAdapterContractOwnerRouteAuditCheck {
	return []BackendAdapterContractOwnerRouteAuditCheck{
		backendAdapterContractOwnerRouteAuditCheck("cli-preview-registered", backendAdapterContractAuditPassBlocked(audit.CLICommandRegistered), "The adapter contract preview is registered as a Go CLI preview."),
		backendAdapterContractOwnerRouteAuditCheck("go-read-model-present", backendAdapterContractAuditPassBlocked(audit.GoReadModelPresent), "The Go Runtime adapter contract read model exists."),
		backendAdapterContractOwnerRouteAuditCheck("fixture-consumption-present", backendAdapterContractAuditPassBlocked(audit.FixtureMatrixConsumesContract), "The offline fixture matrix consumes the no-op adapter contract as fixture evidence."),
		backendAdapterContractOwnerRouteAuditCheck("owner-route-absent", backendAdapterContractAuditPassBlocked(!audit.OwnerDispatchRoutePresent), "The full adapter contract is not accepted by the owner dispatch table."),
		backendAdapterContractOwnerRouteAuditCheck("production-dbus-absent", backendAdapterContractAuditPassBlocked(!audit.ProductionDBusMethodPresent), "The full adapter contract is not exposed as a production D-Bus method."),
		backendAdapterContractOwnerRouteAuditCheck("redacted-route-missing", backendAdapterContractAuditPendingUnless(audit.RedactedProfileRoutePresent), "Owner-local routing now uses a redacted profile audit route instead of the full internal contract."),
		backendAdapterContractOwnerRouteAuditCheck("internal-detail-boundary", backendAdapterContractAuditPassBlocked(audit.FullContractContainsAdapterIDs && audit.KDEFacingProjectionPresent && audit.RedactedProfileRoutePresent), "The full contract still contains internal adapter identifiers, so only the KDE-facing redacted projection is routed."),
		backendAdapterContractOwnerRouteAuditCheck("owner-smoke-coverage-present", backendAdapterContractAuditPendingUnless(audit.OwnerSmokeCoverageReady), "Restricted owner smoke coverage proves the redacted profile route is exercised through Service.Call."),
		backendAdapterContractOwnerRouteAuditCheck("route-decision", backendAdapterContractAuditPassBlocked(backendAdapterContractOwnerRouteDecisionClosed(audit)), "The audit recognizes smoke-covered owner-local redacted routing while keeping the full adapter contract fixture-local."),
		backendAdapterContractOwnerRouteAuditCheck("unsafe-gates-closed", backendAdapterContractAuditPassBlocked(!audit.AdapterInvocationEnabled && !audit.BackendInstallEnabled && !audit.BackendDownloadEnabled && !audit.BackendLaunchEnabled && !audit.BackendProcessStarted && !audit.CommandMaterialized && !audit.ExecutablePathResolved && !audit.SystemServiceStarted && !audit.ProductionBusClaimed && !audit.WriteMethodsEnabled && !audit.HostRootModified && !audit.StateRootPathExposed && !audit.BackendDetailsExposed), "The audit does not invoke, install, download, launch, materialize commands, start services, claim buses, expose paths, or mutate the host."),
	}
}

func backendAdapterContractConfigureOwnerRouteDecision(audit *BackendAdapterContractOwnerRouteAuditPreview) {
	if !audit.CLICommandRegistered || !audit.GoReadModelPresent || !audit.FixtureMatrixConsumesContract {
		audit.RouteDecision = "redacted-profile-route-sources-missing"
		audit.RouteDecisionReason = "The adapter contract owner-route audit remains blocked until required CLI, Go read-model, and fixture-consumption evidence is present."
		audit.CurrentRouteStatus = "fail-closed-owner-route-audit"
		audit.RecommendedNextRoute = "restore-redacted-profile-route-sources"
		audit.DesktopSafeSummary = "The adapter contract owner-route audit is fail-closed because required local source evidence is missing; production D-Bus exposure, Runtime writes, adapter invocation, launch, and host mutation remain disabled."
		return
	}
	if !audit.RedactedProfileRoutePresent || !audit.OwnerLocalRouteCandidateReady {
		audit.RouteDecision = "redacted-profile-route-missing"
		audit.RouteDecisionReason = "The full adapter contract remains fixture-local, but owner-local routing still needs the redacted profile audit route."
		audit.CurrentRouteStatus = "full-contract-fixture-local-redacted-route-missing"
		audit.RecommendedNextRoute = "redacted-adapter-profile-owner-route"
		audit.DesktopSafeSummary = "The no-op adapter contract remains fixture-local while the redacted profile route is missing; production D-Bus exposure, Runtime writes, launch, and host mutation remain disabled."
		return
	}
	if audit.OwnerSmokeCoverageReady {
		audit.RouteDecision = "redacted-profile-route-smoke-covered"
		audit.RouteDecisionReason = "Restricted owner smoke coverage proves the redacted adapter profile owner-local read route is exercised through Service.Call while the full adapter contract stays fixture-local."
		audit.CurrentRouteStatus = "redacted-profile-route-smoke-covered-full-contract-fixture-local-production-dbus-blocked"
		audit.RecommendedNextRoute = "redacted-adapter-profile-production-dbus-gate-review"
		audit.DesktopSafeSummary = "The redacted adapter profile owner-local read route is covered by restricted owner smoke evidence while the full adapter contract stays fixture-local and production D-Bus exposure, Runtime writes, adapter invocation, launch, raw-detail exposure, and host mutation remain disabled."
		return
	}
	audit.RouteDecision = "redacted-profile-route-ready"
	audit.RouteDecisionReason = "The no-op adapter contract remains fixture-local, while owner-local routing can expose only the redacted profile audit route; restricted owner smoke coverage is still required before production review."
	audit.CurrentRouteStatus = "redacted-profile-route-ready-full-contract-fixture-local"
	audit.RecommendedNextRoute = "redacted-adapter-profile-owner-smoke-coverage"
	audit.DesktopSafeSummary = "The no-op adapter contract remains fixture-local while a redacted owner-local profile audit exposes user-safe compatibility modes without internal adapter identifiers, production D-Bus exposure, Runtime writes, launch, or host mutation."
}

func backendAdapterContractOwnerRouteDecisionClosed(audit BackendAdapterContractOwnerRouteAuditPreview) bool {
	if audit.RouteDecision == "redacted-profile-route-ready" {
		return audit.OwnerLocalRouteCandidateReady &&
			audit.RedactedProfileRoutePresent &&
			!audit.OwnerSmokeCoverageReady &&
			!audit.ProductionDBusExposureReady
	}
	if audit.RouteDecision == "redacted-profile-route-smoke-covered" {
		return audit.OwnerLocalRouteCandidateReady &&
			audit.RedactedProfileRoutePresent &&
			audit.OwnerSmokeCoverageReady &&
			!audit.ProductionDBusExposureReady
	}
	return false
}

func backendAdapterContractOwnerRouteAuditCheck(id string, status string, summary string) BackendAdapterContractOwnerRouteAuditCheck {
	return BackendAdapterContractOwnerRouteAuditCheck{ID: id, Status: status, Summary: summary}
}

func backendAdapterContractAuditPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func backendAdapterContractAuditPendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func backendAdapterContractOwnerRouteAuditCheckIDs(checks []BackendAdapterContractOwnerRouteAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countBackendAdapterContractOwnerRouteAuditChecks(checks []BackendAdapterContractOwnerRouteAuditCheck) BackendAdapterContractOwnerRouteAuditCounts {
	counts := BackendAdapterContractOwnerRouteAuditCounts{Total: len(checks)}
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

func backendAdapterContractAuditHasCLICommand(source string) bool {
	return textHasAll(source, []string{"backend-adapter-contract-preview", "runBackendAdapterContractPreview"})
}

func backendAdapterContractAuditHasReadModel(source string) bool {
	return textHasAll(source, []string{"BackendAdapterContractPreview", "NewBackendAdapterContractPreview", "xnix.runtime.backend_adapter_contract.v1"})
}

func backendAdapterContractAuditHasFixtureConsumption(source string) bool {
	return textHasAll(source, []string{"backend-adapter-contract-preview", "NewBackendAdapterContractPreview", "BackendAdapterContractRead"})
}

func backendAdapterContractAuditHasOwnerRoute(source string) bool {
	return textHasAll(source, []string{"GetBackendAdapterContract", "NewBackendAdapterContractPreview"})
}

func backendAdapterContractAuditHasProductionDBusMethod(source string) bool {
	return textHasAll(source, []string{"GetBackendAdapterContract", "backend-adapter-contract-preview"})
}

func backendAdapterContractAuditHasRedactedRoute(source string) bool {
	return textHasAll(source, []string{"GetBackendAdapterProfileAudit", "NewBackendAdapterRedactedProfileAuditPreview"})
}

func backendAdapterContractAuditHasOwnerSmokeCoverage(source string) bool {
	return textHasAll(source, []string{"backend-adapter-redacted-profile-owner-smoke-coverage-preview", "redacted-adapter-profile-owner-smoke-coverage", "GetBackendAdapterProfileAudit", "owner-smoke-batch+runtime-owner-service-call+backend-adapter-redacted-profile-audit-owner-route", "SmokeCoverageReady"})
}

func backendAdapterContractAuditHasRootFlag(source string) bool {
	return textHasAll(source, []string{"backend-adapter-contract-owner-route-audit-preview", "flags.String(\"root\""})
}

func backendAdapterContractOwnerRouteAuditBlockedActions() []string {
	return []string{
		"add the full adapter contract to production D-Bus",
		"route internal adapter identifiers through owner dispatch",
		"expose adapter implementation details to KDE-facing output",
		"enable adapter invocation from owner-route audit",
		"launch compatibility backend from owner-route audit",
		"treat owner smoke coverage as production ownership approval",
		"mutate host root during owner-route audit",
	}
}

func backendAdapterContractOwnerRouteAuditNextRequirements() []string {
	return []string{
		"Keep the redacted adapter profile audit route covered by restricted owner smoke evidence.",
		"Keep routing only redacted user-safe profile states through owner-local dispatch.",
		"Keep the full adapter contract fixture-local until a separate production D-Bus gate review exists.",
		"Require production recipe trust and Runtime write-gate evidence before any adapter invocation.",
	}
}
