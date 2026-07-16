package owner

import (
	"encoding/json"
	"fmt"
)

const smokeBatchServiceCallRequestType = "runtime-owner-service-call"

type SmokeBatchRecord struct {
	Version                     string          `json:"version"`
	SchemaVersion               string          `json:"schema_version"`
	RequestType                 string          `json:"request_type"`
	BatchType                   string          `json:"batch_type"`
	Source                      string          `json:"source"`
	Sequence                    int             `json:"sequence"`
	RecordType                  string          `json:"record_type"`
	Method                      string          `json:"method"`
	Args                        []string        `json:"args"`
	ReadOnlyDispatch            bool            `json:"read_only_dispatch"`
	WriteMethod                 bool            `json:"write_method"`
	RouteReady                  bool            `json:"route_ready"`
	DispatchReady               bool            `json:"dispatch_ready"`
	ErrorName                   string          `json:"error_name,omitempty"`
	ReadDispatchMethodCount     int             `json:"read_dispatch_method_count"`
	WriteMethodCount            int             `json:"write_method_count"`
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
	DesktopSafeSummary          string          `json:"desktop_safe_summary"`
}

func NewSmokeBatchRecords(root string) ([]SmokeBatchRecord, error) {
	service, err := NewService(root, ModeSmokeOwner)
	if err != nil {
		return nil, err
	}
	methods := SupportedReadDispatchMethods()
	records := make([]SmokeBatchRecord, 0, len(methods)+service.candidate.WriteMethodCount)
	sequence := 1
	for _, method := range methods {
		args := smokeBatchArgs(method)
		call, err := service.Call(method, args)
		if err != nil {
			return nil, fmt.Errorf("smoke batch service read %s: %w", method, err)
		}
		payload, err := marshalSmokeBatchServiceCall(call, method)
		if err != nil {
			return nil, err
		}
		record := newSmokeBatchRecord(call.Version, sequence, "read-dispatch", method, args, payload, len(methods), service.candidate.WriteMethodCount)
		record.ReadOnlyDispatch = true
		record.RouteReady = call.DispatchReady
		record.DispatchReady = call.ReadOnlyDispatch && call.DispatchReady && !call.WriteMethodsEnabled
		records = append(records, record)
		sequence++
	}
	for _, writeMethod := range service.candidate.WriteMethods {
		call, err := service.Call(writeMethod.Method, nil)
		if err != nil {
			return nil, err
		}
		payload, err := marshalSmokeBatchServiceCall(call, writeMethod.Method)
		if err != nil {
			return nil, err
		}
		record := newSmokeBatchRecord(call.Version, sequence, "write-denial", writeMethod.Method, nil, payload, len(methods), service.candidate.WriteMethodCount)
		record.WriteMethod = true
		record.ErrorName = call.ErrorName
		record.DispatchReady = call.DispatchReady
		records = append(records, record)
		sequence++
	}
	for _, record := range records {
		if err := validateNoBackendTerms(record, "Runtime owner smoke batch record"); err != nil {
			return nil, err
		}
	}
	return records, nil
}

func marshalSmokeBatchServiceCall(call ServiceCall, method string) (json.RawMessage, error) {
	if call.RequestType != smokeBatchServiceCallRequestType {
		return nil, fmt.Errorf("smoke batch service call %s used request type %s", method, call.RequestType)
	}
	payload, err := json.Marshal(call)
	if err != nil {
		return nil, fmt.Errorf("encode smoke batch service call %s: %w", method, err)
	}
	return payload, nil
}

func newSmokeBatchRecord(version string, sequence int, recordType string, method string, args []string, payload json.RawMessage, readMethodCount int, writeMethodCount int) SmokeBatchRecord {
	return SmokeBatchRecord{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.owner_smoke_batch.v1",
		RequestType:                 "runtime-owner-smoke-batch-record",
		BatchType:                   "restricted-session-owner-call-batch",
		Source:                      "go-runtime-owner-service+service-call-batch",
		Sequence:                    sequence,
		RecordType:                  recordType,
		Method:                      method,
		Args:                        append([]string(nil), args...),
		ReadDispatchMethodCount:     readMethodCount,
		WriteMethodCount:            writeMethodCount,
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
		DesktopSafeSummary:          "Runtime owner smoke batch renders service-call read and write-gate evidence without claiming D-Bus ownership.",
	}
}

func smokeBatchArgs(method string) []string {
	const appID = defaultOwnerApplicationID
	switch method {
	case "ListApplications", "GetEngineCatalog", "GetKDEIntegrationStatus", "GetKDEShellIntegrationPlan",
		"GetTrayStatus", "GetBackendCapabilityMatrix", "GetRuntimeServiceBinding", "GetRuntimeLiveOwnerGate",
		"GetRuntimeOwnerSmokePlan", "GetRuntimeMethodParityManifest", "GetRuntimeOwnerProcess",
		"GetRuntimeOwnerRouteManifest", "GetRuntimeOwnerRecipeTrust", "GetRuntimeOwnerReadiness",
		"GetWindowsCompatibilityWorkstreamsPreview":
		return nil
	case "GetDesktopActivationTransactionPreview", "GetDesktopActivationStatus":
		return []string{appID, "development"}
	case "GetNotificationPlan":
		return []string{appID, "install-failed"}
	case "GetKRunnerQueryPlan":
		return []string{"notepad"}
	case "GetPortalRequestPlan", "GetPortalAccessPolicy":
		return []string{appID, "file-open"}
	case "GetCompatibilityInstallPlan":
		return []string{appID, "development"}
	case "GetRepairPlan":
		return []string{appID, "engine-binding-pending"}
	case "GetTestPlan", "GetTestResult":
		return []string{appID, "preflight"}
	case "GetAIDiagnosticInput", "GetAIDiagnosticRecommendation", "GetAIRepairApprovalGate":
		return []string{appID, "engine-binding-pending", "preflight"}
	case "GetSnapshotPlan":
		return []string{appID, "manual"}
	case "GetRuntimeWriteGate":
		return []string{"Launch"}
	case "GetCompatibilitySettingsChangePlan":
		return []string{appID, "run-mode", "mode", "automatic"}
	case "GetCompatibilityModeSwitchPlan":
		return []string{appID, "automatic"}
	case "GetCompatibilityReviewFlowPlan":
		return []string{appID, "run-mode", "mode", "automatic", "file-open"}
	case "GetCompatibilityActionReviewReceipt":
		return []string{appID, "review-file-manager-action", ownerDispatchDecision}
	case "GetKDECenterPage":
		return []string{appID, ownerDispatchDecision}
	case "GetKDECenterPageSections":
		return []string{appID, ownerDispatchDecision}
	case "GetKDECenterPageSectionDetail":
		return []string{appID, "overview", ownerDispatchDecision}
	default:
		return []string{appID}
	}
}
