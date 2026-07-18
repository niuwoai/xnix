package owner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview struct {
	Version                       string                                                  `json:"version"`
	SchemaVersion                 string                                                  `json:"schema_version"`
	RequestType                   string                                                  `json:"request_type"`
	AuditType                     string                                                  `json:"audit_type"`
	Source                        string                                                  `json:"source"`
	RuntimeMethod                 string                                                  `json:"runtime_method"`
	ReadMethod                    string                                                  `json:"read_method"`
	SubjectRequestType            string                                                  `json:"subject_request_type"`
	SubjectCommand                string                                                  `json:"subject_command"`
	ProposedOwnerMethod           string                                                  `json:"proposed_owner_method"`
	RouteDecision                 string                                                  `json:"route_decision"`
	RouteDecisionReason           string                                                  `json:"route_decision_reason"`
	CurrentRouteStatus            string                                                  `json:"current_route_status"`
	CLICommandRegistered          bool                                                    `json:"cli_command_registered"`
	GoReadModelPresent            bool                                                    `json:"go_read_model_present"`
	OwnerDispatchRoutePresent     bool                                                    `json:"owner_dispatch_route_present"`
	ProductionDBusMethodPresent   bool                                                    `json:"production_dbus_method_present"`
	ConsumesExistingReceipt       bool                                                    `json:"consumes_existing_receipt"`
	RequiresCallerStateRoot       bool                                                    `json:"requires_caller_state_root"`
	ReceiptLookupOwnerManaged     bool                                                    `json:"receipt_lookup_owner_managed"`
	OpaqueReceiptIDSupported      bool                                                    `json:"opaque_receipt_id_supported"`
	FanOutWritesEnabled           bool                                                    `json:"fan_out_writes_enabled"`
	StateRootWritesEnabled        bool                                                    `json:"state_root_writes_enabled"`
	SupportBundleExported         bool                                                    `json:"support_bundle_exported"`
	SupportCaseCreated            bool                                                    `json:"support_case_created"`
	NotificationSent              bool                                                    `json:"notification_sent"`
	OwnerLocalRouteCandidateReady bool                                                    `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady   bool                                                    `json:"production_dbus_exposure_ready"`
	RecommendedNextRoute          string                                                  `json:"recommended_next_route"`
	Checks                        []RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck `json:"checks"`
	CheckIDs                      []string                                                `json:"check_ids"`
	Counts                        RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCounts  `json:"counts"`
	RuntimeOwned                  bool                                                    `json:"runtime_owned"`
	GoRuntimeBacked               bool                                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                                                    `json:"kde_policy_owner"`
	SystemServiceStarted          bool                                                    `json:"system_service_started"`
	SessionBusClaimed             bool                                                    `json:"session_bus_claimed"`
	ProductionBusClaimed          bool                                                    `json:"production_bus_claimed"`
	WriteMethodsEnabled           bool                                                    `json:"write_methods_enabled"`
	RuntimeWritesEnabled          bool                                                    `json:"runtime_writes_enabled"`
	BackendLaunchEnabled          bool                                                    `json:"backend_launch_enabled"`
	BackendProcessStarted         bool                                                    `json:"backend_process_started"`
	NetworkRequired               bool                                                    `json:"network_required"`
	HostRootModified              bool                                                    `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                                                    `json:"privileged_container_required"`
	StateRootPathExposed          bool                                                    `json:"state_root_path_exposed"`
	BackendDetailsExposed         bool                                                    `json:"backend_details_exposed"`
	BlockedActions                []string                                                `json:"blocked_actions"`
	NextRequirements              []string                                                `json:"next_requirements"`
	DesktopSafeSummary            string                                                  `json:"desktop_safe_summary"`
}

type RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(root string) (RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview, error) {
	version, err := readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditVersion(root)
	if err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview{}, err
	}
	sources := restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root)
	cliCommandRegistered := restrictedOwnerSmokeReceiptFanOutAuditHasCLICommand(sources.GoCLI)
	goReadModelPresent := restrictedOwnerSmokeReceiptFanOutAuditHasReadModel(sources.GoReadModel)
	ownerDispatchRoutePresent := restrictedOwnerSmokeReceiptFanOutAuditHasOwnerRoute(sources.OwnerDispatch)
	productionDBusMethodPresent := restrictedOwnerSmokeReceiptFanOutAuditHasProductionDBusMethod(sources.DBusContract)
	consumesExistingReceipt := restrictedOwnerSmokeReceiptFanOutAuditConsumesReceipt(sources.GoReadModel)
	requiresCallerStateRoot := restrictedOwnerSmokeReceiptFanOutAuditHasCallerStateRoot(sources.OwnerDispatch)
	receiptLookupOwnerManaged := restrictedOwnerSmokeReceiptFanOutAuditHasOwnerManagedLookup(sources.OwnerDispatch)
	opaqueReceiptIDSupported := restrictedOwnerSmokeReceiptFanOutAuditHasOpaqueReceiptID(sources.OwnerDispatch)
	ownerLocalRouteCandidateReady := ownerDispatchRoutePresent && receiptLookupOwnerManaged && opaqueReceiptIDSupported && !requiresCallerStateRoot
	routeDecision := "owner-local-route-blocked"
	currentRouteStatus := "owner-local-read-route-missing"
	routeDecisionReason := "The restricted smoke receipt fan-out owner-local route is not ready until the owner dispatch route and opaque receipt lookup sources are present."
	recommendedNextRoute := "restricted-owner-smoke-fanout-owner-local-read-route"
	desktopSafeSummary := "The restricted smoke receipt fan-out owner-local route remains blocked because required owner dispatch or opaque receipt lookup evidence is missing; production D-Bus exposure, Runtime writes, support side effects, launch, and host mutation remain disabled."
	if ownerLocalRouteCandidateReady {
		routeDecision = "owner-local-route-ready"
		currentRouteStatus = "owner-local-read-route-ready-production-dbus-blocked"
		routeDecisionReason = "The restricted smoke receipt fan-out now routes through owner-managed opaque receipt lookup without caller-supplied state-root paths, while missing receipts stay fail-closed."
		recommendedNextRoute = "restricted-owner-smoke-fanout-owner-smoke-coverage"
		desktopSafeSummary = "The restricted smoke receipt fan-out has an owner-local read route through opaque receipt lookup; production D-Bus exposure, Runtime writes, support side effects, launch, and host mutation remain disabled."
	}
	audit := RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview{
		Version:                       version,
		SchemaVersion:                 "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1",
		RequestType:                   "restricted-owner-smoke-receipt-fanout-owner-route-audit-preview",
		AuditType:                     "restricted-owner-smoke-receipt-fanout-owner-route-audit",
		Source:                        "restricted-owner-smoke-receipt-fanout-preview+runtime-owner-dispatch+dbus-contract",
		RuntimeMethod:                 "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAudit",
		ReadMethod:                    "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview",
		SubjectRequestType:            "restricted-owner-smoke-receipt-fanout-preview",
		SubjectCommand:                "restricted-owner-smoke-receipt-fanout-preview",
		ProposedOwnerMethod:           "GetRestrictedOwnerSmokeReceiptFanOut",
		RouteDecision:                 routeDecision,
		RouteDecisionReason:           routeDecisionReason,
		CurrentRouteStatus:            currentRouteStatus,
		CLICommandRegistered:          cliCommandRegistered,
		GoReadModelPresent:            goReadModelPresent,
		OwnerDispatchRoutePresent:     ownerDispatchRoutePresent,
		ProductionDBusMethodPresent:   productionDBusMethodPresent,
		ConsumesExistingReceipt:       consumesExistingReceipt,
		RequiresCallerStateRoot:       requiresCallerStateRoot,
		ReceiptLookupOwnerManaged:     receiptLookupOwnerManaged,
		OpaqueReceiptIDSupported:      opaqueReceiptIDSupported,
		FanOutWritesEnabled:           false,
		StateRootWritesEnabled:        false,
		SupportBundleExported:         false,
		SupportCaseCreated:            false,
		NotificationSent:              false,
		OwnerLocalRouteCandidateReady: ownerLocalRouteCandidateReady,
		ProductionDBusExposureReady:   false,
		RecommendedNextRoute:          recommendedNextRoute,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		SystemServiceStarted:          false,
		SessionBusClaimed:             false,
		ProductionBusClaimed:          false,
		WriteMethodsEnabled:           false,
		RuntimeWritesEnabled:          false,
		BackendLaunchEnabled:          false,
		BackendProcessStarted:         false,
		NetworkRequired:               false,
		HostRootModified:              false,
		PrivilegedContainerRequired:   false,
		StateRootPathExposed:          false,
		BackendDetailsExposed:         false,
		BlockedActions:                restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditBlockedActions(),
		NextRequirements:              restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditNextRequirements(),
		DesktopSafeSummary:            desktopSafeSummary,
	}
	checks := restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditChecks(audit)
	audit.Checks = checks
	audit.CheckIDs = restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheckIDs(checks)
	audit.Counts = countRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditChecks(checks)
	if err := validateNoBackendTerms(audit, "restricted owner smoke receipt fan-out owner-route audit preview"); err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview{}, err
	}
	return audit, nil
}

type restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSourceSet struct {
	GoCLI         string
	CLICommand    string
	GoReadModel   string
	OwnerDispatch string
	DBusContract  string
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root string) restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSourceSet {
	return restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSourceSet{
		GoCLI:         readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root, []string{"cmd/xnix-runtime-go/main.go"}),
		CLICommand:    readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root, []string{"cmd/xnix-runtime-go/restricted_owner_smoke_receipt_commands.go"}),
		GoReadModel:   readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root, []string{"internal/runtime/owner/restricted_smoke_receipt_fanout.go"}),
		OwnerDispatch: readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root, []string{"internal/runtime/owner/dispatch.go", "internal/runtime/owner/restricted_smoke_receipt_lookup.go", "internal/runtime/owner/restricted_smoke_receipt_fanout.go"}),
		DBusContract:  readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root, []string{"runtime/dbus/org.xnix.Compatibility1.xml"}),
	}
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditChecks(audit RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview) []RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck {
	return []RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck{
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("cli-preview-registered", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(audit.CLICommandRegistered), "The restricted smoke receipt fan-out preview is registered as a Go CLI preview."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("go-read-model-present", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(audit.GoReadModelPresent), "The Go owner read model exists for the restricted smoke receipt fan-out."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("owner-route-present", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(audit.OwnerDispatchRoutePresent), "The fan-out is accepted by the owner dispatch table."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("production-dbus-absent", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(!audit.ProductionDBusMethodPresent), "The fan-out is not exposed as a production D-Bus method."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("receipt-consumption-present", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(audit.ConsumesExistingReceipt), "The fan-out consumes an existing restricted owner smoke receipt instead of creating one."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("caller-state-root-boundary", restrictedOwnerSmokeReceiptFanOutAuditPendingUnless(!audit.RequiresCallerStateRoot), "Owner-local read routes must not require caller-supplied state-root paths."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("owner-managed-receipt-lookup", restrictedOwnerSmokeReceiptFanOutAuditPendingUnless(audit.ReceiptLookupOwnerManaged && audit.OpaqueReceiptIDSupported), "Owner-local routing needs opaque receipt identifiers resolved inside the Runtime owner."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("route-decision", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(audit.RouteDecision == "owner-local-route-ready" && audit.ReceiptLookupOwnerManaged && audit.OpaqueReceiptIDSupported && audit.OwnerDispatchRoutePresent && audit.OwnerLocalRouteCandidateReady && !audit.ProductionDBusExposureReady), "The audit recognizes the owner-local fan-out route while keeping production D-Bus blocked."),
		restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck("unsafe-gates-closed", restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(!audit.FanOutWritesEnabled && !audit.StateRootWritesEnabled && !audit.SupportBundleExported && !audit.SupportCaseCreated && !audit.NotificationSent && !audit.SystemServiceStarted && !audit.SessionBusClaimed && !audit.ProductionBusClaimed && !audit.WriteMethodsEnabled && !audit.RuntimeWritesEnabled && !audit.BackendLaunchEnabled && !audit.BackendProcessStarted && !audit.NetworkRequired && !audit.HostRootModified && !audit.PrivilegedContainerRequired && !audit.StateRootPathExposed && !audit.BackendDetailsExposed), "The audit does not write, export support bundles, create cases, notify, start services, claim buses, launch, expose paths, or mutate the host."),
	}
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck(id string, status string, summary string) RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck {
	return RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck{ID: id, Status: status, Summary: summary}
}

func restrictedOwnerSmokeReceiptFanOutAuditPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func restrictedOwnerSmokeReceiptFanOutAuditPendingUnless(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheckIDs(checks []RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditChecks(checks []RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCheck) RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCounts {
	counts := RestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditCounts{Total: len(checks)}
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

func restrictedOwnerSmokeReceiptFanOutAuditHasCLICommand(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"restricted-owner-smoke-receipt-fanout-preview", "runRestrictedOwnerSmokeReceiptFanOutPreview"})
}

func restrictedOwnerSmokeReceiptFanOutAuditHasReadModel(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"RestrictedOwnerSmokeReceiptFanOutPreview", "NewRestrictedOwnerSmokeReceiptFanOutPreview", "xnix.runtime.restricted_owner_smoke_receipt_fanout.v1"})
}

func restrictedOwnerSmokeReceiptFanOutAuditHasOwnerRoute(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"GetRestrictedOwnerSmokeReceiptFanOut", "NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview", "ResolveRestrictedOwnerSmokeReceipt"})
}

func restrictedOwnerSmokeReceiptFanOutAuditHasProductionDBusMethod(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"GetRestrictedOwnerSmokeReceiptFanOut", "restricted-owner-smoke-receipt-fanout-preview"})
}

func restrictedOwnerSmokeReceiptFanOutAuditConsumesReceipt(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"LoadRestrictedOwnerSmokeReceipt", "ReceiptConsumed", "validateRestrictedOwnerSmokeReceiptFanOutReceipt"})
}

func restrictedOwnerSmokeReceiptFanOutAuditHasCallerStateRoot(source string) bool {
	return !restrictedOwnerSmokeReceiptFanOutAuditHasOwnerRoute(source)
}

func restrictedOwnerSmokeReceiptFanOutAuditHasOwnerManagedLookup(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"ResolveRestrictedOwnerSmokeReceipt", "owner-managed-receipt-lookup"})
}

func restrictedOwnerSmokeReceiptFanOutAuditHasOpaqueReceiptID(source string) bool {
	return restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source, []string{"opaque_receipt_id", "restricted-owner-smoke-receipt-id"})
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditBlockedActions() []string {
	return []string{
		"add restricted smoke receipt fan-out to production D-Bus before owner-smoke coverage is reviewed",
		"accept caller-supplied state-root paths through owner dispatch",
		"export support bundles from owner-route audit",
		"create support cases from owner-route audit",
		"send notifications from owner-route audit",
		"mutate host root during owner-route audit",
	}
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteAuditNextRequirements() []string {
	return []string{
		"Cover the owner-local fan-out route in restricted owner smoke evidence.",
		"Attach verified receipt evidence inside the Runtime owner without caller state-root paths.",
		"Keep missing receipts fail-closed until receipt availability is proven.",
		"Keep production D-Bus exposure blocked until route-manifest and owner-smoke coverage prove the path.",
	}
}

func readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditVersion(root string) (string, error) {
	content, err := os.ReadFile(filepath.Join(root, "VERSION"))
	if err != nil {
		return "", fmt.Errorf("read VERSION: %w", err)
	}
	return strings.TrimSpace(string(content)), nil
}

func readRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditSources(root string, relativePaths []string) string {
	var builder strings.Builder
	for _, relativePath := range relativePaths {
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(relativePath)))
		if err != nil {
			continue
		}
		builder.WriteString(string(content))
		builder.WriteString("\n")
	}
	return builder.String()
}

func restrictedOwnerSmokeReceiptFanOutAuditTextHasAll(source string, tokens []string) bool {
	for _, token := range tokens {
		if !strings.Contains(source, token) {
			return false
		}
	}
	return true
}
