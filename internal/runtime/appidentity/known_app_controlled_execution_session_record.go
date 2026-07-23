package appidentity

import (
	"errors"

	"xnix.local/xnix/internal/runtime/execution"
)

const (
	KnownAppControlledExecutionSessionRecordSchemaVersion = "xnix.runtime.known_app_controlled_execution_session_record.v1"
	KnownAppControlledExecutionSessionRecordRequestType   = "known-app-controlled-execution-session-record"
)

type KnownAppControlledExecutionSessionRecordRequest struct {
	AppID         string
	StateRoot     string
	ReceiptID     string
	CacheRoot     string
	GuestBoundary string
}

type KnownAppControlledExecutionSessionRecord struct {
	SchemaVersion                    string           `json:"schema_version"`
	RequestType                      string           `json:"request_type"`
	Source                           string           `json:"source"`
	RuntimeMethod                    string           `json:"runtime_method"`
	AppID                            string           `json:"app_id"`
	DisplayName                      string           `json:"display_name"`
	AppVersion                       string           `json:"app_version"`
	ReceiptID                        string           `json:"receipt_id"`
	ReceiptAccepted                  bool             `json:"receipt_accepted"`
	GuestBoundary                    string           `json:"guest_boundary"`
	GuestBoundaryAccepted            bool             `json:"guest_boundary_accepted"`
	ControlledDispatchReady          bool             `json:"controlled_dispatch_ready"`
	ControlledDispatchRequestCreated bool             `json:"controlled_dispatch_request_created"`
	RuntimeOwnedDispatchRequest      bool             `json:"runtime_owned_dispatch_request"`
	ExecutionSessionID               string           `json:"execution_session_id,omitempty"`
	ExecutionSessionState            string           `json:"execution_session_state"`
	ExecutionSessionHandoffCreated   bool             `json:"execution_session_handoff_created"`
	SessionHandoffReady              bool             `json:"session_handoff_ready"`
	RecordState                      string           `json:"record_state"`
	LedgerRecordWritten              bool             `json:"ledger_record_written"`
	SessionRecordWritten             bool             `json:"session_record_written"`
	TransactionRecordType            string           `json:"transaction_record_type,omitempty"`
	SessionRecordType                string           `json:"session_record_type,omitempty"`
	TransactionRelativePath          string           `json:"transaction_relative_path,omitempty"`
	SessionRelativePath              string           `json:"session_relative_path,omitempty"`
	SessionSHA256                    string           `json:"session_sha256,omitempty"`
	SessionState                     string           `json:"session_state"`
	CompatibilityCenterState         string           `json:"compatibility_center_state"`
	Gates                            []execution.Gate `json:"gates,omitempty"`
	BlockedReasons                   []string         `json:"blocked_reasons,omitempty"`
	RuntimeOwnedExecutionSession     bool             `json:"runtime_owned_execution_session"`
	GoRuntimeBacked                  bool             `json:"go_runtime_backed"`
	KDEPolicyOwner                   bool             `json:"kde_policy_owner"`
	StateRootPathExposed             bool             `json:"state_root_path_exposed"`
	ReceiptPathExposed               bool             `json:"receipt_path_exposed"`
	TransactionPathExposed           bool             `json:"transaction_path_exposed"`
	SessionPathExposed               bool             `json:"session_path_exposed"`
	SessionRegistered                bool             `json:"session_registered"`
	WindowObserved                   bool             `json:"window_observed"`
	TaskManagerEntryActive           bool             `json:"task_manager_entry_active"`
	KWinRuleApplied                  bool             `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled            bool             `json:"live_tray_bridge_enabled"`
	DispatchAllowed                  bool             `json:"dispatch_allowed"`
	DispatchStarted                  bool             `json:"dispatch_started"`
	ExecutionStarted                 bool             `json:"execution_started"`
	DirectLaunchEnabled              bool             `json:"direct_launch_enabled"`
	DesktopLaunchEnabled             bool             `json:"desktop_launch_enabled"`
	BackendLaunchEnabled             bool             `json:"backend_launch_enabled"`
	BackendProcessStarted            bool             `json:"backend_process_started"`
	HostRootModified                 bool             `json:"host_root_modified"`
	DockerSocketMounted              bool             `json:"docker_socket_mounted"`
	BroadHostMountRequired           bool             `json:"broad_host_mount_required"`
	RawArtifactPathExposed           bool             `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed            bool             `json:"backend_details_exposed"`
	DesktopSafeSummary               string           `json:"desktop_safe_summary"`
}

func RecordKnownAppControlledExecutionSession(request KnownAppControlledExecutionSessionRecordRequest) (KnownAppControlledExecutionSessionRecord, error) {
	handoff, err := PreviewKnownAppControlledExecutionSession(KnownAppControlledExecutionSessionRequest{
		AppID:         request.AppID,
		StateRoot:     request.StateRoot,
		ReceiptID:     request.ReceiptID,
		CacheRoot:     request.CacheRoot,
		GuestBoundary: request.GuestBoundary,
	})
	if err != nil {
		return KnownAppControlledExecutionSessionRecord{}, err
	}
	record := baseKnownAppControlledExecutionSessionRecord(handoff)
	if !handoff.SessionHandoffReady {
		record.RecordState = "blocked"
		record.DesktopSafeSummary = handoff.DisplayName + " controlled execution session record is blocked until the Runtime handoff is ready."
		return validateKnownAppControlledExecutionSessionRecord(record)
	}

	ledger, err := execution.NewLedger(request.StateRoot)
	if err != nil {
		return KnownAppControlledExecutionSessionRecord{}, err
	}
	tx := knownAppControlledExecutionSessionTransaction(handoff)
	transactionRecord, err := ledger.Record(tx)
	if err != nil {
		return KnownAppControlledExecutionSessionRecord{}, err
	}
	sessionRecord, err := ledger.RecordSession(tx.RequestID)
	if err != nil {
		return KnownAppControlledExecutionSessionRecord{}, err
	}

	record.RecordState = "persisted"
	record.LedgerRecordWritten = true
	record.SessionRecordWritten = true
	record.TransactionRecordType = transactionRecord.RecordType
	record.SessionRecordType = sessionRecord.RecordType
	record.TransactionRelativePath = transactionRecord.RelativePath
	record.SessionRelativePath = sessionRecord.RelativePath
	record.SessionSHA256 = sessionRecord.SHA256
	record.SessionState = sessionRecord.SessionState
	record.CompatibilityCenterState = sessionRecord.CompatibilityCenterState
	record.Gates = append([]execution.Gate{}, sessionRecord.Gates...)
	record.BlockedReasons = append([]string{}, sessionRecord.BlockedReasons...)
	record.RuntimeOwnedExecutionSession = true
	record.DesktopSafeSummary = handoff.DisplayName + " controlled execution session was recorded under the Runtime state root for later KDE-safe consumption."
	return validateKnownAppControlledExecutionSessionRecord(record)
}

func baseKnownAppControlledExecutionSessionRecord(handoff KnownAppControlledExecutionSessionPreview) KnownAppControlledExecutionSessionRecord {
	return KnownAppControlledExecutionSessionRecord{
		SchemaVersion:                    KnownAppControlledExecutionSessionRecordSchemaVersion,
		RequestType:                      KnownAppControlledExecutionSessionRecordRequestType,
		Source:                           KnownAppControlledExecutionSessionRequestType + "+execution-ledger+execution-session-record",
		RuntimeMethod:                    "RecordKnownAppControlledExecutionSession",
		AppID:                            handoff.AppID,
		DisplayName:                      handoff.DisplayName,
		AppVersion:                       handoff.AppVersion,
		ReceiptID:                        handoff.ReceiptID,
		ReceiptAccepted:                  handoff.ReceiptAccepted,
		GuestBoundary:                    handoff.GuestBoundary,
		GuestBoundaryAccepted:            handoff.GuestBoundaryAccepted,
		ControlledDispatchReady:          handoff.ControlledDispatchReady,
		ControlledDispatchRequestCreated: handoff.ControlledDispatchRequestCreated,
		RuntimeOwnedDispatchRequest:      handoff.RuntimeOwnedDispatchRequest,
		ExecutionSessionID:               handoff.ExecutionSessionID,
		ExecutionSessionState:            handoff.ExecutionSessionState,
		ExecutionSessionHandoffCreated:   handoff.ExecutionSessionHandoffCreated,
		SessionHandoffReady:              handoff.SessionHandoffReady,
		RecordState:                      "closed",
		SessionState:                     "not-recorded",
		CompatibilityCenterState:         "not-recorded",
		RuntimeOwnedExecutionSession:     false,
		GoRuntimeBacked:                  true,
		KDEPolicyOwner:                   false,
		StateRootPathExposed:             false,
		ReceiptPathExposed:               false,
		TransactionPathExposed:           false,
		SessionPathExposed:               false,
		SessionRegistered:                false,
		WindowObserved:                   false,
		TaskManagerEntryActive:           false,
		KWinRuleApplied:                  false,
		LiveTrayBridgeEnabled:            false,
		DispatchAllowed:                  false,
		DispatchStarted:                  false,
		ExecutionStarted:                 false,
		DirectLaunchEnabled:              false,
		DesktopLaunchEnabled:             false,
		BackendLaunchEnabled:             false,
		BackendProcessStarted:            false,
		HostRootModified:                 false,
		DockerSocketMounted:              false,
		BroadHostMountRequired:           false,
		RawArtifactPathExposed:           false,
		BackendDetailsExposed:            false,
		DesktopSafeSummary:               handoff.DisplayName + " controlled execution session record is closed until the Runtime handoff is ready.",
	}
}

func knownAppControlledExecutionSessionTransaction(handoff KnownAppControlledExecutionSessionPreview) execution.Transaction {
	return execution.Transaction{
		RequestID:      handoff.ExecutionSessionID,
		ApplicationID:  handoff.AppID,
		Profile:        "known-app-managed-guest",
		State:          execution.StateBlocked,
		ReviewDecision: execution.DecisionApproved,
		Gates: []execution.Gate{
			{ID: "known-app-launch-authorization", Status: execution.GatePass, Reason: "opaque authorization receipt accepted"},
			{ID: "managed-guest-boundary", Status: execution.GatePass, Reason: "controlled managed guest boundary accepted"},
			{ID: "controlled-dispatch-request", Status: execution.GatePass, Reason: "controlled dispatch request materialized"},
			{ID: "managed-artifact", Status: execution.GatePass, Reason: "managed application artifact verified"},
			{ID: "controlled-execution-session-handoff", Status: execution.GatePass, Reason: "Runtime execution session handoff ready"},
			{ID: "runtime-write-gate", Status: execution.GateBlocked, Reason: "dispatch runner must consume the recorded session before execution"},
		},
		BlockedReasons:   []string{"runtime-write-gate: dispatch runner must consume the recorded session before execution"},
		LaunchAllowed:    false,
		LaunchEnabled:    false,
		BackendStarted:   false,
		HostRootModified: false,
		NetworkRequired:  false,
		Summary:          "Runtime recorded a known application execution session handoff; execution remains disabled until the managed runner consumes the record.",
	}
}

func validateKnownAppControlledExecutionSessionRecord(record KnownAppControlledExecutionSessionRecord) (KnownAppControlledExecutionSessionRecord, error) {
	if record.LedgerRecordWritten || record.SessionRecordWritten {
		switch {
		case record.RecordState != "persisted":
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record writes require persisted state")
		case !record.SessionHandoffReady || !record.ExecutionSessionHandoffCreated:
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record requires a ready handoff")
		case !record.RuntimeOwnedDispatchRequest || !record.RuntimeOwnedExecutionSession || !record.GoRuntimeBacked:
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record requires Runtime-owned Go evidence")
		case record.TransactionRelativePath == "" || record.SessionRelativePath == "" || record.SessionSHA256 == "":
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record requires relative receipt evidence")
		case record.DispatchAllowed || record.DispatchStarted || record.ExecutionStarted || record.BackendLaunchEnabled || record.BackendProcessStarted:
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record must not start dispatch, execution, or backend processes")
		case record.SessionRegistered || record.WindowObserved || record.TaskManagerEntryActive || record.KWinRuleApplied || record.LiveTrayBridgeEnabled:
			return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record must not activate live desktop surfaces")
		}
	}
	if record.StateRootPathExposed || record.ReceiptPathExposed || record.TransactionPathExposed || record.SessionPathExposed || record.HostRootModified || record.DockerSocketMounted || record.BroadHostMountRequired || record.RawArtifactPathExposed || record.BackendDetailsExposed {
		return KnownAppControlledExecutionSessionRecord{}, errors.New("controlled execution session record exposes unsafe state or mutates host state")
	}
	if err := validateNoBackendTerms(record, "known app controlled execution session record"); err != nil {
		return KnownAppControlledExecutionSessionRecord{}, err
	}
	return record, nil
}
