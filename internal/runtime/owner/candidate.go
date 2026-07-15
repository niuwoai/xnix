package owner

import (
	"encoding/json"
	"fmt"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	writeMethodDisabledError = "org.xnix.Compatibility1.Error.WriteMethodDisabled"
)

type CandidateMode string

const (
	ModePreview    CandidateMode = "preview"
	ModeSmokeOwner CandidateMode = "smoke-owner"
)

type Candidate struct {
	Version                     string           `json:"version"`
	SchemaVersion               string           `json:"schema_version"`
	RequestType                 string           `json:"request_type"`
	OwnerType                   string           `json:"owner_type"`
	Source                      string           `json:"source"`
	BusName                     string           `json:"bus_name"`
	ObjectPath                  string           `json:"object_path"`
	Interface                   string           `json:"interface"`
	Mode                        string           `json:"mode"`
	ServiceActivationReady      bool             `json:"service_activation_ready"`
	ReadOnlyRouteTableReady     bool             `json:"read_only_route_table_ready"`
	ReadOnlyServeReady          bool             `json:"read_only_serve_ready"`
	RouteCount                  int              `json:"route_count"`
	GoRouteCount                int              `json:"go_route_count"`
	CCoreRouteCount             int              `json:"c_core_route_count"`
	RubyLegacyRouteCount        int              `json:"ruby_legacy_route_count"`
	Routes                      []ReadRoute      `json:"routes"`
	WriteMethods                []DisabledWrite  `json:"write_methods"`
	WriteMethodCount            int              `json:"write_method_count"`
	WriteMethodsEnabled         bool             `json:"write_methods_enabled"`
	RuntimeOwned                bool             `json:"runtime_owned"`
	GoRuntimeBacked             bool             `json:"go_runtime_backed"`
	KDEPolicyOwner              bool             `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool             `json:"kde_may_claim_runtime_ownership"`
	SmokeOwnerMode              bool             `json:"smoke_owner_mode"`
	ProductionOwnerMode         bool             `json:"production_owner_mode"`
	EventLoopStarted            bool             `json:"event_loop_started"`
	SessionBusClaimed           bool             `json:"session_bus_claimed"`
	ProductionBusClaimed        bool             `json:"production_bus_claimed"`
	SystemServiceStarted        bool             `json:"system_service_started"`
	NetworkRequired             bool             `json:"network_required"`
	HostRootModified            bool             `json:"host_root_modified"`
	PrivilegedContainerRequired bool             `json:"privileged_container_required"`
	BackendDetailsExposed       bool             `json:"backend_details_exposed"`
	Checks                      []CandidateCheck `json:"checks"`
	CheckIDs                    []string         `json:"check_ids"`
	Counts                      CandidateCounts  `json:"counts"`
	BlockedActions              []string         `json:"blocked_actions"`
	NextRequirements            []string         `json:"next_requirements"`
	DesktopSafeSummary          string           `json:"desktop_safe_summary"`
}

type ReadRoute struct {
	Method      string `json:"method"`
	Source      string `json:"source"`
	GoCommand   string `json:"go_command"`
	RouteStatus string `json:"route_status"`
	Ready       bool   `json:"ready"`
}

type DisabledWrite struct {
	Method          string `json:"method"`
	ErrorName       string `json:"error_name"`
	DispatchEnabled bool   `json:"dispatch_enabled"`
	RequestCreated  bool   `json:"request_created"`
	Summary         string `json:"summary"`
}

type CandidateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type CandidateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewCandidate(root string, mode CandidateMode) (Candidate, error) {
	if mode == "" {
		mode = ModePreview
	}
	if mode != ModePreview && mode != ModeSmokeOwner {
		return Candidate{}, fmt.Errorf("unsupported Runtime owner mode: %s", mode)
	}

	serviceBinding, err := appidentity.NewRuntimeServiceBindingPreview(root)
	if err != nil {
		return Candidate{}, err
	}
	routeManifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return Candidate{}, err
	}
	writeGate, err := appidentity.NewRuntimeWriteGatePreview(root, "Launch")
	if err != nil {
		return Candidate{}, err
	}

	routes := candidateReadRoutes(routeManifest.Routes)
	writeMethods := candidateDisabledWrites(writeGate.SupportedWriteMethods)
	checks := candidateChecks(serviceBinding, routeManifest, mode)
	counts := countCandidateChecks(checks)
	readOnlyServeReady := serviceBinding.ActivationBindingReady && routeManifest.GoOwnerRouteCoverageReady

	candidate := Candidate{
		Version:                     serviceBinding.Version,
		SchemaVersion:               "xnix.runtime.owner_candidate.v1",
		RequestType:                 "runtime-owner-candidate",
		OwnerType:                   "go-runtime-owner-candidate",
		Source:                      "runtime-service-binding-preview+runtime-owner-route-manifest-preview+runtime-write-gate-preview",
		BusName:                     serviceBinding.BusName,
		ObjectPath:                  serviceBinding.ObjectPath,
		Interface:                   serviceBinding.Interface,
		Mode:                        string(mode),
		ServiceActivationReady:      serviceBinding.ActivationBindingReady,
		ReadOnlyRouteTableReady:     routeManifest.GoOwnerRouteCoverageReady,
		ReadOnlyServeReady:          readOnlyServeReady,
		RouteCount:                  routeManifest.RouteCounts.Total,
		GoRouteCount:                routeManifest.RouteCounts.GoRouted,
		CCoreRouteCount:             routeManifest.RouteCounts.CCoreBacked,
		RubyLegacyRouteCount:        routeManifest.RouteCounts.RubyLegacy,
		Routes:                      routes,
		WriteMethods:                writeMethods,
		WriteMethodCount:            len(writeMethods),
		WriteMethodsEnabled:         false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		SmokeOwnerMode:              mode == ModeSmokeOwner,
		ProductionOwnerMode:         false,
		EventLoopStarted:            false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		Checks:                      checks,
		CheckIDs:                    candidateCheckIDs(checks),
		Counts:                      counts,
		BlockedActions:              candidateBlockedActions(),
		NextRequirements:            candidateNextRequirements(),
		DesktopSafeSummary:          candidateSummary(readOnlyServeReady, mode),
	}
	if err := validateNoBackendTerms(candidate, "Runtime owner candidate"); err != nil {
		return Candidate{}, err
	}
	return candidate, nil
}

func DisabledWriteResponse(method string) (DisabledWrite, error) {
	if strings.TrimSpace(method) == "" {
		return DisabledWrite{}, fmt.Errorf("write method is required")
	}
	return DisabledWrite{
		Method:          method,
		ErrorName:       writeMethodDisabledError,
		DispatchEnabled: false,
		RequestCreated:  false,
		Summary:         "Runtime write methods remain disabled until production owner readiness is proven.",
	}, nil
}

func candidateReadRoutes(routes []appidentity.RuntimeOwnerRoute) []ReadRoute {
	result := make([]ReadRoute, 0, len(routes))
	for _, route := range routes {
		result = append(result, ReadRoute{
			Method:      route.Method,
			Source:      route.CurrentSource,
			GoCommand:   route.GoCommand,
			RouteStatus: route.RouteStatus,
			Ready:       route.GoRouteReady,
		})
	}
	return result
}

func candidateDisabledWrites(methods []string) []DisabledWrite {
	result := make([]DisabledWrite, 0, len(methods))
	for _, method := range methods {
		response, err := DisabledWriteResponse(method)
		if err == nil {
			result = append(result, response)
		}
	}
	return result
}

func candidateChecks(serviceBinding appidentity.RuntimeServiceBindingPreview, routeManifest appidentity.RuntimeOwnerRouteManifestPreview, mode CandidateMode) []CandidateCheck {
	return []CandidateCheck{
		candidateCheck("service-activation", passBlocked(serviceBinding.ActivationBindingReady), "D-Bus activation files and packaged Runtime wrapper must be aligned before owner smoke."),
		candidateCheck("read-only-route-table", passBlocked(routeManifest.GoOwnerRouteCoverageReady), "Every read-only Runtime method must have Go owner route evidence."),
		candidateCheck("write-method-gate", "pass", "Write methods stay disabled with deterministic Runtime errors."),
		candidateCheck("smoke-owner-mode", passPending(mode == ModeSmokeOwner), "Smoke owner mode can expose the owner candidate without claiming the production bus."),
		candidateCheck("production-bus-claim", "pending", "A later restricted D-Bus smoke must prove stable bus-name ownership."),
		candidateCheck("host-safety-boundary", "pass", "Owner candidate must not start services, require network, or mutate the host root."),
	}
}

func candidateCheck(id string, status string, summary string) CandidateCheck {
	return CandidateCheck{ID: id, Status: status, Summary: summary}
}

func passBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func passPending(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func candidateCheckIDs(checks []CandidateCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countCandidateChecks(checks []CandidateCheck) CandidateCounts {
	counts := CandidateCounts{Total: len(checks)}
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

func candidateBlockedActions() []string {
	return []string{
		"claim production D-Bus name from owner candidate",
		"start system service from owner candidate",
		"enable write methods from owner candidate",
		"let KDE claim Runtime ownership",
		"mutate host root during owner candidate evaluation",
		"expose backend implementation details through owner candidate output",
	}
}

func candidateNextRequirements() []string {
	return []string{
		"Bind this candidate to a restricted session-bus smoke.",
		"Implement the D-Bus event loop for read-only methods.",
		"Keep deterministic write-method denial until production readiness is proven.",
		"Promote the packaged entrypoint to the Go owner only after smoke parity passes.",
	}
}

func candidateSummary(readOnlyServeReady bool, mode CandidateMode) string {
	if !readOnlyServeReady {
		return "Runtime owner candidate is blocked until service activation and Go read-only route coverage are ready."
	}
	if mode == ModeSmokeOwner {
		return "Runtime owner candidate can serve as a restricted smoke target, but production D-Bus ownership remains gated."
	}
	return "Runtime owner candidate has read-only route coverage, but production D-Bus ownership remains gated."
}

func validateNoBackendTerms(value any, label string) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("encode %s JSON: %w", label, err)
	}
	text := strings.ToLower(string(encoded))
	forbidden := []string{"prefix", ".exe", "wine ", "wine/", "proton", "qemu-system", "program files", ".wine"}
	for _, term := range forbidden {
		if strings.Contains(text, term) {
			return fmt.Errorf("%s exposes forbidden backend term: %s", label, term)
		}
	}
	return nil
}
