package owner

import (
	"encoding/json"
	"fmt"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type ReadDispatch struct {
	Version                     string          `json:"version"`
	SchemaVersion               string          `json:"schema_version"`
	RequestType                 string          `json:"request_type"`
	DispatchType                string          `json:"dispatch_type"`
	Source                      string          `json:"source"`
	Method                      string          `json:"method"`
	Args                        []string        `json:"args"`
	RouteSource                 string          `json:"route_source"`
	GoCommand                   string          `json:"go_command"`
	RouteStatus                 string          `json:"route_status"`
	RouteReady                  bool            `json:"route_ready"`
	ReadOnlyDispatch            bool            `json:"read_only_dispatch"`
	WriteMethod                 bool            `json:"write_method"`
	WriteMethodsEnabled         bool            `json:"write_methods_enabled"`
	RuntimeOwned                bool            `json:"runtime_owned"`
	GoRuntimeBacked             bool            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool            `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool            `json:"kde_may_claim_runtime_ownership"`
	EventLoopStarted            bool            `json:"event_loop_started"`
	SessionBusClaimed           bool            `json:"session_bus_claimed"`
	ProductionBusClaimed        bool            `json:"production_bus_claimed"`
	SystemServiceStarted        bool            `json:"system_service_started"`
	NetworkRequired             bool            `json:"network_required"`
	HostRootModified            bool            `json:"host_root_modified"`
	PrivilegedContainerRequired bool            `json:"privileged_container_required"`
	BackendDetailsExposed       bool            `json:"backend_details_exposed"`
	Payload                     json.RawMessage `json:"payload"`
	BlockedActions              []string        `json:"blocked_actions"`
	NextRequirements            []string        `json:"next_requirements"`
	DesktopSafeSummary          string          `json:"desktop_safe_summary"`
}

type ownerReadPayloadBuilder func(root string, args []string) (any, error)

var ownerReadDispatchers = map[string]ownerReadPayloadBuilder{
	"GetRuntimeServiceBinding": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeServiceBinding", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeServiceBindingPreview(root)
	},
	"GetRuntimeLiveOwnerGate": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeLiveOwnerGate", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeLiveOwnerGatePreview(root)
	},
	"GetRuntimeOwnerProcess": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerProcess", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerProcessPreview(root)
	},
	"GetRuntimeOwnerSmokePlan": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerSmokePlan", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerSmokePlanPreview(root)
	},
	"GetRuntimeMethodParityManifest": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeMethodParityManifest", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeMethodParityManifestPreview(root)
	},
	"GetRuntimeOwnerRouteManifest": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerRouteManifest", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	},
	"GetRuntimeOwnerRecipeTrust": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerRecipeTrust", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerRecipeTrustPreview(root)
	},
	"GetRuntimeOwnerReadiness": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerReadiness", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerReadinessPreview(root)
	},
	"GetRuntimeWriteGate": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeWriteGate", args, 1); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeWriteGatePreview(root, args[0])
	},
}

func DispatchRead(root string, method string, args []string) (ReadDispatch, error) {
	method = strings.TrimSpace(method)
	normalizedArgs := append([]string{}, args...)
	if method == "" {
		return ReadDispatch{}, fmt.Errorf("read method is required")
	}
	if isReservedWriteMethod(method) {
		return ReadDispatch{}, fmt.Errorf("%s is a write method; use disabled write-method response", method)
	}
	builder, ok := ownerReadDispatchers[method]
	if !ok {
		return ReadDispatch{}, fmt.Errorf("unsupported owner read dispatch method: %s", method)
	}

	routeManifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return ReadDispatch{}, err
	}
	route := ownerReadRoute(routeManifest.Routes, method)
	payloadValue, err := builder(root, normalizedArgs)
	if err != nil {
		return ReadDispatch{}, err
	}
	payload, err := json.Marshal(payloadValue)
	if err != nil {
		return ReadDispatch{}, fmt.Errorf("encode %s dispatch payload: %w", method, err)
	}

	dispatch := ReadDispatch{
		Version:                     routeManifest.Version,
		SchemaVersion:               "xnix.runtime.owner_read_dispatch.v1",
		RequestType:                 "runtime-owner-read-dispatch",
		DispatchType:                "go-owner-read-dispatch",
		Source:                      "go-runtime-owner-candidate+in-process-read-dispatch",
		Method:                      method,
		Args:                        normalizedArgs,
		RouteSource:                 route.CurrentSource,
		GoCommand:                   route.GoCommand,
		RouteStatus:                 route.RouteStatus,
		RouteReady:                  route.GoRouteReady,
		ReadOnlyDispatch:            true,
		WriteMethod:                 false,
		WriteMethodsEnabled:         false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		EventLoopStarted:            false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		Payload:                     payload,
		BlockedActions:              readDispatchBlockedActions(),
		NextRequirements:            readDispatchNextRequirements(),
		DesktopSafeSummary:          "Runtime owner read dispatch can render selected read-only owner methods without claiming D-Bus ownership.",
	}
	if err := validateNoBackendTerms(dispatch, "Runtime owner read dispatch"); err != nil {
		return ReadDispatch{}, err
	}
	return dispatch, nil
}

func SupportedReadDispatchMethods() []string {
	methods := []string{
		"GetRuntimeServiceBinding",
		"GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerProcess",
		"GetRuntimeOwnerSmokePlan",
		"GetRuntimeMethodParityManifest",
		"GetRuntimeOwnerRouteManifest",
		"GetRuntimeOwnerRecipeTrust",
		"GetRuntimeOwnerReadiness",
		"GetRuntimeWriteGate",
	}
	return append([]string(nil), methods...)
}

func requireArgCount(method string, args []string, expected int) error {
	if len(args) != expected {
		return fmt.Errorf("%s requires %d argument(s), got %d", method, expected, len(args))
	}
	return nil
}

func ownerReadRoute(routes []appidentity.RuntimeOwnerRoute, method string) appidentity.RuntimeOwnerRoute {
	for _, route := range routes {
		if route.Method == method {
			return route
		}
	}
	return appidentity.RuntimeOwnerRoute{
		Method:        method,
		CurrentSource: "go-owner-local-preview",
		TargetSource:  "go-runtime-owner",
		GoCommand:     ownerReadDispatchCommand(method),
		RouteStatus:   "owner-local-preview-ready",
		GoRouteReady:  true,
	}
}

func ownerReadDispatchCommand(method string) string {
	switch method {
	case "GetRuntimeServiceBinding":
		return "runtime-service-binding-preview"
	case "GetRuntimeLiveOwnerGate":
		return "runtime-live-owner-gate-preview"
	case "GetRuntimeOwnerProcess":
		return "runtime-owner-process-preview"
	case "GetRuntimeOwnerSmokePlan":
		return "runtime-owner-smoke-plan-preview"
	case "GetRuntimeMethodParityManifest":
		return "runtime-method-parity-manifest-preview"
	case "GetRuntimeOwnerRouteManifest":
		return "runtime-owner-route-manifest-preview"
	case "GetRuntimeOwnerRecipeTrust":
		return "runtime-owner-recipe-trust-preview"
	case "GetRuntimeOwnerReadiness":
		return "runtime-owner-readiness-preview"
	case "GetRuntimeWriteGate":
		return "runtime-write-gate-preview"
	default:
		return "unsupported-owner-read-preview"
	}
}

func isReservedWriteMethod(method string) bool {
	for _, writeMethod := range []string{"InstallRecipe", "Launch", "CreateSnapshot", "RestoreSnapshot"} {
		if method == writeMethod {
			return true
		}
	}
	return false
}

func readDispatchBlockedActions() []string {
	return []string{
		"claim a D-Bus name from read dispatch preview",
		"start a system service from read dispatch preview",
		"enable write methods from read dispatch preview",
		"let KDE route Runtime policy directly",
		"mutate host root during read dispatch preview",
		"expose backend implementation details through read dispatch output",
	}
}

func readDispatchNextRequirements() []string {
	return []string{
		"Bind read dispatch to the restricted owner event loop.",
		"Map every read-only Runtime method to an in-process owner handler.",
		"Keep write methods disabled until production owner readiness is proven.",
		"Promote the packaged entrypoint only after D-Bus read dispatch parity passes.",
	}
}
