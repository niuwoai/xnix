package owner

import (
	"encoding/json"
	"fmt"
)

type Service struct {
	root      string
	candidate Candidate
}

type ServiceCall struct {
	Version                     string          `json:"version"`
	SchemaVersion               string          `json:"schema_version"`
	RequestType                 string          `json:"request_type"`
	ServiceType                 string          `json:"service_type"`
	Source                      string          `json:"source"`
	Method                      string          `json:"method"`
	Args                        []string        `json:"args"`
	CallType                    string          `json:"call_type"`
	BusName                     string          `json:"bus_name"`
	ObjectPath                  string          `json:"object_path"`
	Interface                   string          `json:"interface"`
	RouteCount                  int             `json:"route_count"`
	GoRouteCount                int             `json:"go_route_count"`
	WriteMethodCount            int             `json:"write_method_count"`
	ReadOnlyServeReady          bool            `json:"read_only_serve_ready"`
	ReadOnlyDispatch            bool            `json:"read_only_dispatch"`
	WriteMethod                 bool            `json:"write_method"`
	WriteMethodsEnabled         bool            `json:"write_methods_enabled"`
	DispatchReady               bool            `json:"dispatch_ready"`
	ErrorName                   string          `json:"error_name,omitempty"`
	RuntimeOwned                bool            `json:"runtime_owned"`
	GoRuntimeBacked             bool            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool            `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool            `json:"kde_may_claim_runtime_ownership"`
	InProcessServiceReady       bool            `json:"in_process_service_ready"`
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

func NewService(root string, mode CandidateMode) (Service, error) {
	candidate, err := NewCandidate(root, mode)
	if err != nil {
		return Service{}, err
	}
	return Service{root: root, candidate: candidate}, nil
}

func (service Service) Call(method string, args []string) (ServiceCall, error) {
	switch {
	case isSupportedReadMethod(method):
		dispatch, err := DispatchRead(service.root, method, args)
		if err != nil {
			return ServiceCall{}, err
		}
		payload, err := json.Marshal(dispatch)
		if err != nil {
			return ServiceCall{}, fmt.Errorf("encode service read call %s: %w", method, err)
		}
		return service.validatedCall(method, args, "read-dispatch", true, false, "", dispatch.RouteReady && dispatch.ReadOnlyDispatch && !dispatch.WriteMethodsEnabled, payload)
	case isReservedWriteMethod(method):
		response, err := DisabledWriteResponse(method)
		if err != nil {
			return ServiceCall{}, err
		}
		payload, err := json.Marshal(response)
		if err != nil {
			return ServiceCall{}, fmt.Errorf("encode service write denial %s: %w", method, err)
		}
		return service.validatedCall(method, args, "write-denial", false, true, response.ErrorName, !response.DispatchEnabled && !response.RequestCreated, payload)
	default:
		return ServiceCall{}, fmt.Errorf("unsupported Runtime owner service method: %s", method)
	}
}

func (service Service) validatedCall(method string, args []string, callType string, readOnlyDispatch bool, writeMethod bool, errorName string, dispatchReady bool, payload json.RawMessage) (ServiceCall, error) {
	call := service.newCall(method, args, callType, readOnlyDispatch, writeMethod, errorName, dispatchReady, payload)
	if err := validateNoBackendTerms(call, "Runtime owner service call"); err != nil {
		return ServiceCall{}, err
	}
	return call, nil
}

func (service Service) newCall(method string, args []string, callType string, readOnlyDispatch bool, writeMethod bool, errorName string, dispatchReady bool, payload json.RawMessage) ServiceCall {
	call := ServiceCall{
		Version:                     service.candidate.Version,
		SchemaVersion:               "xnix.runtime.owner_service_call.v1",
		RequestType:                 "runtime-owner-service-call",
		ServiceType:                 "go-runtime-owner-in-process-service",
		Source:                      "go-runtime-owner-service+read-dispatch+write-gate",
		Method:                      method,
		Args:                        append([]string(nil), args...),
		CallType:                    callType,
		BusName:                     service.candidate.BusName,
		ObjectPath:                  service.candidate.ObjectPath,
		Interface:                   service.candidate.Interface,
		RouteCount:                  service.candidate.RouteCount,
		GoRouteCount:                service.candidate.GoRouteCount,
		WriteMethodCount:            service.candidate.WriteMethodCount,
		ReadOnlyServeReady:          service.candidate.ReadOnlyServeReady,
		ReadOnlyDispatch:            readOnlyDispatch,
		WriteMethod:                 writeMethod,
		WriteMethodsEnabled:         false,
		DispatchReady:               dispatchReady,
		ErrorName:                   errorName,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		InProcessServiceReady:       service.candidate.ReadOnlyServeReady,
		EventLoopStarted:            false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		Payload:                     payload,
		BlockedActions:              serviceBlockedActions(),
		NextRequirements:            serviceNextRequirements(),
		DesktopSafeSummary:          serviceCallSummary(callType),
	}
	return call
}

func isSupportedReadMethod(method string) bool {
	for _, readMethod := range ownerReadDispatchMethodOrder {
		if method == readMethod {
			return true
		}
	}
	return false
}

func serviceBlockedActions() []string {
	return []string{
		"claim production D-Bus name from in-process owner service",
		"start system service from in-process owner service",
		"enable write methods from in-process owner service",
		"let KDE own Runtime policy",
		"mutate host root during owner service calls",
		"expose backend implementation details through owner service calls",
	}
}

func serviceNextRequirements() []string {
	return []string{
		"Bind this in-process service boundary to a restricted private session-bus loop.",
		"Route D-Bus read-only calls through Service.Call without changing payload semantics.",
		"Keep write methods mapped to deterministic disabled errors.",
		"Promote to production ownership only after bus-call parity and host-safety smoke pass.",
	}
}

func serviceCallSummary(callType string) string {
	if callType == "write-denial" {
		return "Runtime owner service boundary denies write calls before production ownership is enabled."
	}
	return "Runtime owner service boundary serves read calls in process while D-Bus ownership remains gated."
}
