package appidentity

import (
	"errors"
	"strings"
)

const runtimeWriteGateErrorName = "org.xnix.Compatibility1.Error.WriteMethodDisabled"

var runtimeWriteGateRequiredGateIDs = []string{
	"production-dbus-gate-review",
	"production-service-activation-preflight",
	"human-authorization-receipt",
	"production-human-authorization-receipt-consolidation",
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
	Version                     string                         `json:"version"`
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	GateType                    string                         `json:"gate_type"`
	Source                      string                         `json:"source"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	MethodName                  string                         `json:"method_name"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	GateDecision                string                         `json:"gate_decision"`
	ProductionGateDecision      string                         `json:"production_gate_decision"`
	WriteMethodEnabled          bool                           `json:"write_method_enabled"`
	DispatchEnabled             bool                           `json:"dispatch_enabled"`
	RequestObjectCreated        bool                           `json:"request_object_created"`
	ExecutionStarted            bool                           `json:"execution_started"`
	ProductionGate              RuntimeWriteGateProductionGate `json:"production_gate"`
	SupportedWriteMethods       []string                       `json:"supported_write_methods"`
	RequiredGates               []RuntimeWriteGateCheck        `json:"required_gates"`
	RequiredGateIDs             []string                       `json:"required_gate_ids"`
	Counts                      RuntimeWriteGateCounts         `json:"counts"`
	DenialErrorName             string                         `json:"denial_error_name"`
	NetworkRequired             bool                           `json:"network_required"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type RuntimeWriteGateProductionGate struct {
	RequestType                      string `json:"request_type"`
	PreflightType                    string `json:"preflight_type"`
	PreflightDecision                string `json:"preflight_decision"`
	ServiceActivationPreflightReady  bool   `json:"service_activation_preflight_ready"`
	ProductionDBusGateReady          bool   `json:"production_dbus_gate_ready"`
	HumanAuthorizationPreflightReady bool   `json:"human_authorization_preflight_ready"`
	HumanAuthorizationRequired       bool   `json:"human_authorization_required"`
	HumanAuthorizationGranted        bool   `json:"human_authorization_granted"`
	AuthorizationReceiptAccepted     bool   `json:"authorization_receipt_accepted"`
	ProductionActivationReady        bool   `json:"production_activation_ready"`
	RestrictedSmokeReady             bool   `json:"restricted_smoke_ready"`
	SystemServiceStarted             bool   `json:"system_service_started"`
	ProductionBusClaimed             bool   `json:"production_bus_claimed"`
	WriteMethodsEnabled              bool   `json:"write_methods_enabled"`
	BackendLaunchEnabled             bool   `json:"backend_launch_enabled"`
	NetworkRequired                  bool   `json:"network_required"`
	HostRootModified                 bool   `json:"host_root_modified"`
	PrivilegedContainerRequired      bool   `json:"privileged_container_required"`
	BackendDetailsExposed            bool   `json:"backend_details_exposed"`
}

type RuntimeWriteGateCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeWriteGateCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
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
	serviceActivation, err := NewRuntimeServiceActivationPreflightPreview(root)
	if err != nil {
		return RuntimeWriteGatePreview{}, err
	}
	productionGate := runtimeWriteGateProductionGate(serviceActivation)
	requiredGates := runtimeWriteGateChecks(productionGate)

	preview := RuntimeWriteGatePreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.write_gate.v1",
		RequestType:                 "runtime-write-gate-preview",
		GateType:                    "runtime-write-gate",
		Source:                      "go-runtime-write-gate+runtime-service-activation-preflight-preview+production-dbus-gate-review-preview+production-dbus-human-authorization-preflight-preview+production-human-authorization-receipt-consolidation-preview",
		RuntimeMethod:               "GetRuntimeWriteGate",
		ReadMethod:                  "GetRuntimeWriteGatePreview",
		MethodName:                  methodName,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		GateDecision:                "blocked-until-production-backend",
		ProductionGateDecision:      runtimeWriteGateProductionDecision(productionGate),
		WriteMethodEnabled:          false,
		DispatchEnabled:             false,
		RequestObjectCreated:        false,
		ExecutionStarted:            false,
		ProductionGate:              productionGate,
		SupportedWriteMethods:       append([]string(nil), runtimeWriteGateSupportedWriteMethods...),
		RequiredGates:               requiredGates,
		RequiredGateIDs:             append([]string(nil), runtimeWriteGateRequiredGateIDs...),
		Counts:                      countRuntimeWriteGateChecks(requiredGates),
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

func runtimeWriteGateChecks(productionGate RuntimeWriteGateProductionGate) []RuntimeWriteGateCheck {
	return []RuntimeWriteGateCheck{
		runtimeWriteGateCheck("production-dbus-gate-review", runtimeWriteGatePassBlocked(productionGate.ProductionDBusGateReady), "The write gate must consume the production D-Bus gate review before any write dispatch can be considered."),
		runtimeWriteGateCheck("production-service-activation-preflight", runtimeWriteGatePassBlocked(productionGate.ServiceActivationPreflightReady), "The write gate must consume the production service activation preflight while service start remains disabled."),
		runtimeWriteGateCheck("human-authorization-receipt", runtimeWriteGatePendingPass(productionGate.AuthorizationReceiptAccepted), "A human authorization receipt remains required before production write dispatch can be considered."),
		runtimeWriteGateCheck("production-human-authorization-receipt-consolidation", runtimeWriteGatePassBlocked(productionGate.HumanAuthorizationPreflightReady && productionGate.HumanAuthorizationRequired), "The write gate must consume the consolidated opaque authorization receipt boundary while receipt acceptance remains disabled."),
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

func runtimeWriteGateProductionGate(serviceActivation RuntimeServiceActivationPreflightPreview) RuntimeWriteGateProductionGate {
	return RuntimeWriteGateProductionGate{
		RequestType:                      serviceActivation.RequestType,
		PreflightType:                    serviceActivation.PreflightType,
		PreflightDecision:                serviceActivation.PreflightDecision,
		ServiceActivationPreflightReady:  serviceActivation.Counts.Blocked == 0 && serviceActivation.ProductionDBusGateReady && serviceActivation.HumanAuthorizationPreflightReady,
		ProductionDBusGateReady:          serviceActivation.ProductionDBusGateReady,
		HumanAuthorizationPreflightReady: serviceActivation.HumanAuthorizationPreflightReady,
		HumanAuthorizationRequired:       serviceActivation.HumanAuthorizationRequired,
		HumanAuthorizationGranted:        serviceActivation.HumanAuthorizationGranted,
		AuthorizationReceiptAccepted:     serviceActivation.AuthorizationReceiptAccepted,
		ProductionActivationReady:        serviceActivation.ProductionActivationReady,
		RestrictedSmokeReady:             serviceActivation.RestrictedSmokeReady,
		SystemServiceStarted:             serviceActivation.SystemServiceStarted,
		ProductionBusClaimed:             serviceActivation.ProductionBusClaimed,
		WriteMethodsEnabled:              serviceActivation.WriteMethodsEnabled,
		BackendLaunchEnabled:             serviceActivation.BackendLaunchEnabled,
		NetworkRequired:                  serviceActivation.NetworkRequired,
		HostRootModified:                 serviceActivation.HostRootModified,
		PrivilegedContainerRequired:      serviceActivation.PrivilegedContainerRequired,
		BackendDetailsExposed:            serviceActivation.BackendDetailsExposed,
	}
}

func runtimeWriteGateProductionDecision(productionGate RuntimeWriteGateProductionGate) string {
	if !productionGate.ProductionDBusGateReady || !productionGate.ServiceActivationPreflightReady {
		return "production-gate-consumption-blocked"
	}
	if !productionGate.AuthorizationReceiptAccepted {
		return "production-gates-consumed-write-gate-disabled"
	}
	return "production-authorization-present-write-gate-still-disabled"
}

func runtimeWriteGatePassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeWriteGatePendingPass(passed bool) string {
	if passed {
		return "pass"
	}
	return "pending"
}

func countRuntimeWriteGateChecks(checks []RuntimeWriteGateCheck) RuntimeWriteGateCounts {
	counts := RuntimeWriteGateCounts{Total: len(checks)}
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
