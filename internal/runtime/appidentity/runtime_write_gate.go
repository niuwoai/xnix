package appidentity

import (
	"errors"
	"strings"
)

const runtimeWriteGateErrorName = "org.xnix.Compatibility1.Error.WriteMethodDisabled"

var runtimeWriteGateRequiredGateIDs = []string{
	"production-runtime-owner",
	"backend-binding-ready",
	"recipe-trust-production",
	"user-action-review",
	"portal-approval-if-sensitive",
	"snapshot-preflight-for-risky-change",
}

var runtimeWriteGateSupportedWriteMethods = []string{
	"InstallRecipe",
	"Launch",
	"CreateSnapshot",
	"RestoreSnapshot",
}

type RuntimeWriteGatePreview struct {
	Version                     string                  `json:"version"`
	SchemaVersion               string                  `json:"schema_version"`
	RequestType                 string                  `json:"request_type"`
	GateType                    string                  `json:"gate_type"`
	Source                      string                  `json:"source"`
	RuntimeMethod               string                  `json:"runtime_method"`
	ReadMethod                  string                  `json:"read_method"`
	MethodName                  string                  `json:"method_name"`
	RuntimeOwned                bool                    `json:"runtime_owned"`
	GoRuntimeBacked             bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                    `json:"kde_policy_owner"`
	GateDecision                string                  `json:"gate_decision"`
	WriteMethodEnabled          bool                    `json:"write_method_enabled"`
	DispatchEnabled             bool                    `json:"dispatch_enabled"`
	RequestObjectCreated        bool                    `json:"request_object_created"`
	ExecutionStarted            bool                    `json:"execution_started"`
	SupportedWriteMethods       []string                `json:"supported_write_methods"`
	RequiredGates               []RuntimeWriteGateCheck `json:"required_gates"`
	RequiredGateIDs             []string                `json:"required_gate_ids"`
	DenialErrorName             string                  `json:"denial_error_name"`
	NetworkRequired             bool                    `json:"network_required"`
	HostRootModified            bool                    `json:"host_root_modified"`
	PrivilegedContainerRequired bool                    `json:"privileged_container_required"`
	BackendDetailsExposed       bool                    `json:"backend_details_exposed"`
	DesktopSafeSummary          string                  `json:"desktop_safe_summary"`
}

type RuntimeWriteGateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewRuntimeWriteGatePreview(root string, methodName string) (RuntimeWriteGatePreview, error) {
	if root == "" {
		root = "."
	}
	if !containsString(runtimeWriteGateSupportedWriteMethods, methodName) {
		return RuntimeWriteGatePreview{}, errors.New("method must be one of: " + strings.Join(runtimeWriteGateSupportedWriteMethods, ", "))
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return RuntimeWriteGatePreview{}, err
	}

	preview := RuntimeWriteGatePreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.write_gate.v1",
		RequestType:                 "runtime-write-gate-preview",
		GateType:                    "runtime-write-gate",
		Source:                      "go-runtime-write-gate",
		RuntimeMethod:               "GetRuntimeWriteGate",
		ReadMethod:                  "GetRuntimeWriteGatePreview",
		MethodName:                  methodName,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		GateDecision:                "blocked-until-production-backend",
		WriteMethodEnabled:          false,
		DispatchEnabled:             false,
		RequestObjectCreated:        false,
		ExecutionStarted:            false,
		SupportedWriteMethods:       append([]string(nil), runtimeWriteGateSupportedWriteMethods...),
		RequiredGates:               runtimeWriteGateChecks(),
		RequiredGateIDs:             append([]string(nil), runtimeWriteGateRequiredGateIDs...),
		DenialErrorName:             runtimeWriteGateErrorName,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          methodName + " is visible to desktop integrations but remains blocked until production Runtime gates pass.",
	}
	if err := validateNoBackendTerms(preview, "Runtime write gate preview"); err != nil {
		return RuntimeWriteGatePreview{}, err
	}
	return preview, nil
}

func runtimeWriteGateChecks() []RuntimeWriteGateCheck {
	return []RuntimeWriteGateCheck{
		runtimeWriteGateCheck("production-runtime-owner", "pending", "The packaged Runtime service must own the stable D-Bus name."),
		runtimeWriteGateCheck("backend-binding-ready", "pending", "A managed compatibility backend must be selected and verified."),
		runtimeWriteGateCheck("recipe-trust-production", "pending", "The application recipe must satisfy production trust policy."),
		runtimeWriteGateCheck("user-action-review", "pending", "The Compatibility Center must record user approval for the operation."),
		runtimeWriteGateCheck("portal-approval-if-sensitive", "pending", "Sensitive desktop resources must be mediated through XDG Desktop Portal."),
		runtimeWriteGateCheck("snapshot-preflight-for-risky-change", "pending", "Risky state changes must have a Runtime restore point plan."),
	}
}

func runtimeWriteGateCheck(id string, status string, summary string) RuntimeWriteGateCheck {
	return RuntimeWriteGateCheck{ID: id, Status: status, Summary: summary}
}
