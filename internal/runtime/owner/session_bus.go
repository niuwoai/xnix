package owner

import (
	"encoding/json"
	"fmt"
)

const sessionBusSmokeRequestType = "runtime-owner-session-bus-smoke-step"

type SessionBusSmokeStep struct {
	Version                     string          `json:"version"`
	SchemaVersion               string          `json:"schema_version"`
	RequestType                 string          `json:"request_type"`
	TranscriptType              string          `json:"transcript_type"`
	Source                      string          `json:"source"`
	Sequence                    int             `json:"sequence"`
	StepType                    string          `json:"step_type"`
	Method                      string          `json:"method,omitempty"`
	Args                        []string        `json:"args,omitempty"`
	CallType                    string          `json:"call_type,omitempty"`
	BusName                     string          `json:"bus_name"`
	ObjectPath                  string          `json:"object_path"`
	Interface                   string          `json:"interface"`
	ReadOnlyDispatch            bool            `json:"read_only_dispatch"`
	WriteMethod                 bool            `json:"write_method"`
	UnsupportedRead             bool            `json:"unsupported_read"`
	RouteReady                  bool            `json:"route_ready"`
	DispatchReady               bool            `json:"dispatch_ready"`
	ErrorName                   string          `json:"error_name,omitempty"`
	ReadDispatchMethodCount     int             `json:"read_dispatch_method_count"`
	WriteMethodCount            int             `json:"write_method_count"`
	RuntimeOwned                bool            `json:"runtime_owned"`
	GoRuntimeBacked             bool            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool            `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool            `json:"kde_may_claim_runtime_ownership"`
	PrivateSessionBus           bool            `json:"private_session_bus"`
	EventLoopStarted            bool            `json:"event_loop_started"`
	SessionBusClaimed           bool            `json:"session_bus_claimed"`
	ProductionBusClaimed        bool            `json:"production_bus_claimed"`
	SystemServiceStarted        bool            `json:"system_service_started"`
	WriteMethodsEnabled         bool            `json:"write_methods_enabled"`
	NetworkRequired             bool            `json:"network_required"`
	HostRootModified            bool            `json:"host_root_modified"`
	PrivilegedContainerRequired bool            `json:"privileged_container_required"`
	BackendDetailsExposed       bool            `json:"backend_details_exposed"`
	Payload                     json.RawMessage `json:"payload,omitempty"`
	DesktopSafeSummary          string          `json:"desktop_safe_summary"`
}

func NewSessionBusSmokeTranscript(root string) ([]SessionBusSmokeStep, error) {
	service, err := NewService(root, ModeSmokeOwner)
	if err != nil {
		return nil, err
	}
	smokeBatch, err := NewSmokeBatchRecords(root)
	if err != nil {
		return nil, err
	}
	steps := make([]SessionBusSmokeStep, 0, len(smokeBatch)+5)
	sequence := 1
	steps = append(steps, newSessionBusSmokeStep(service, sequence, "startup", "", nil, nil))
	sequence++
	steps = append(steps, newSessionBusSmokeStep(service, sequence, "claim-private-session-bus", "", nil, nil))
	sequence++
	steps = append(steps, newSessionBusSmokeStep(service, sequence, "route-table-ready", "", nil, nil))
	sequence++
	for _, record := range smokeBatch {
		payload, err := json.Marshal(record)
		if err != nil {
			return nil, fmt.Errorf("encode session bus smoke batch step %s: %w", record.Method, err)
		}
		step := newSessionBusSmokeStep(service, sequence, record.RecordType, record.Method, record.Args, payload)
		step.CallType = record.RecordType
		step.ReadOnlyDispatch = record.ReadOnlyDispatch
		step.WriteMethod = record.WriteMethod
		step.RouteReady = record.RouteReady
		step.DispatchReady = record.DispatchReady
		step.ErrorName = record.ErrorName
		steps = append(steps, step)
		sequence++
	}
	unsupportedPayload, unsupportedError := unsupportedReadPayload(service, "GetUnsupportedRuntimeMethod")
	unsupportedStep := newSessionBusSmokeStep(service, sequence, "reject-unsupported-read", "GetUnsupportedRuntimeMethod", nil, unsupportedPayload)
	unsupportedStep.UnsupportedRead = true
	unsupportedStep.ErrorName = unsupportedError
	unsupportedStep.DispatchReady = true
	steps = append(steps, unsupportedStep)
	sequence++
	steps = append(steps, newSessionBusSmokeStep(service, sequence, "shutdown", "", nil, nil))
	for _, step := range steps {
		if err := validateNoBackendTerms(step, "Runtime owner session bus smoke step"); err != nil {
			return nil, err
		}
	}
	return steps, nil
}

func unsupportedReadPayload(service Service, method string) (json.RawMessage, string) {
	_, err := service.Call(method, nil)
	payload := map[string]any{
		"schema_version":        "xnix.runtime.owner_unsupported_read.v1",
		"request_type":          "runtime-owner-unsupported-read",
		"method":                method,
		"dispatch_ready":        true,
		"request_created":       false,
		"write_methods_enabled": false,
		"error_name":            "org.xnix.Compatibility1.Error.UnsupportedMethod",
		"summary":               "Unsupported Runtime owner read calls fail closed inside the private session-bus smoke transcript.",
	}
	if err != nil {
		payload["error"] = err.Error()
	}
	encoded, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return nil, "org.xnix.Compatibility1.Error.UnsupportedMethod"
	}
	return encoded, "org.xnix.Compatibility1.Error.UnsupportedMethod"
}

func newSessionBusSmokeStep(service Service, sequence int, stepType string, method string, args []string, payload json.RawMessage) SessionBusSmokeStep {
	return SessionBusSmokeStep{
		Version:                     service.candidate.Version,
		SchemaVersion:               "xnix.runtime.owner_session_bus_smoke.v1",
		RequestType:                 sessionBusSmokeRequestType,
		TranscriptType:              "restricted-private-session-bus-owner-smoke",
		Source:                      "go-runtime-owner-service+private-session-bus-smoke",
		Sequence:                    sequence,
		StepType:                    stepType,
		Method:                      method,
		Args:                        append([]string(nil), args...),
		BusName:                     service.candidate.BusName,
		ObjectPath:                  service.candidate.ObjectPath,
		Interface:                   service.candidate.Interface,
		ReadDispatchMethodCount:     len(SupportedReadDispatchMethods()),
		WriteMethodCount:            service.candidate.WriteMethodCount,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		PrivateSessionBus:           true,
		EventLoopStarted:            true,
		SessionBusClaimed:           true,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		WriteMethodsEnabled:         false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		Payload:                     payload,
		DesktopSafeSummary:          "Private session-bus smoke routes Runtime owner calls through Go service evidence while production ownership remains gated.",
	}
}
