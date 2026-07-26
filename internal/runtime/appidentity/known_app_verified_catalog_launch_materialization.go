package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppVerifiedCatalogLaunchMaterializationSchemaVersion = "xnix.runtime.known_app_verified_catalog_launch_materialization.v1"
	KnownAppVerifiedCatalogLaunchMaterializationRequestType   = "known-app-verified-catalog-launch-materialization-record"
	KnownAppVerifiedCatalogDispatchRequestSchemaVersion       = "xnix.runtime.known_app_verified_catalog_dispatch_request.v1"
	KnownAppVerifiedCatalogDispatchRequestType                = "known-app-verified-catalog-dispatch-request-record"
)

type KnownAppVerifiedCatalogLaunchMaterializationRequest struct {
	StateRoot           string
	HandoffRelativePath string
	CacheRoot           string
	GuestBoundary       string
	RecordedAtUTC       time.Time
}

type KnownAppVerifiedCatalogDispatchRequestRecordRequest struct {
	StateRoot           string
	HandoffRelativePath string
	CacheRoot           string
	GuestBoundary       string
	LauncherName        string
	RecordedAtUTC       time.Time
}

type KnownAppVerifiedCatalogLaunchMaterializationRecord struct {
	Version                              string `json:"version"`
	SchemaVersion                        string `json:"schema_version"`
	RequestType                          string `json:"request_type"`
	Source                               string `json:"source"`
	RuntimeMethod                        string `json:"runtime_method"`
	ReadMethod                           string `json:"read_method"`
	AppID                                string `json:"app_id"`
	DisplayName                          string `json:"display_name"`
	AppVersion                           string `json:"app_version"`
	HandoffConsumed                      bool   `json:"handoff_consumed"`
	HandoffRelativePath                  string `json:"handoff_relative_path"`
	HandoffDigestVerified                bool   `json:"handoff_digest_verified"`
	AcceptanceRequestType                string `json:"acceptance_request_type"`
	AcceptanceType                       string `json:"acceptance_type"`
	AcceptanceReady                      bool   `json:"acceptance_ready"`
	RunPlanMatched                       bool   `json:"run_plan_matched"`
	ExistingWindowsApp                   bool   `json:"existing_windows_app"`
	KnownPortableCatalogBacked           bool   `json:"known_portable_catalog_backed"`
	LaunchAttempted                      bool   `json:"launch_attempted"`
	ChecksumVerified                     bool   `json:"checksum_verified"`
	MarkerObserved                       bool   `json:"marker_observed"`
	RuntimeStartedIsolatedGuest          bool   `json:"runtime_started_isolated_guest"`
	IsolatedGuestExecutionObserved       bool   `json:"isolated_guest_execution_observed"`
	CompatibilityEngineExecutionObserved bool   `json:"compatibility_engine_execution_observed"`
	OutputRedacted                       bool   `json:"output_redacted"`
	Q4ExecutionObserved                  bool   `json:"q4_execution_observed"`
	HostCompilationAvoided               bool   `json:"host_compilation_avoided"`
	LaunchAuthorizationReceiptID         string `json:"launch_authorization_receipt_id"`
	LaunchAuthorizationReceiptRecorded   bool   `json:"launch_authorization_receipt_recorded"`
	ControlledExecutionSessionID         string `json:"controlled_execution_session_id,omitempty"`
	ControlledSessionRecordState         string `json:"controlled_session_record_state"`
	ControlledSessionRelativePath        string `json:"controlled_session_relative_path,omitempty"`
	ControlledSessionDigestVerified      bool   `json:"controlled_session_digest_verified"`
	SessionGatedReviewReceiptID          string `json:"session_gated_review_receipt_id,omitempty"`
	SessionGatedReviewReceiptRecorded    bool   `json:"session_gated_review_receipt_recorded"`
	MaterializationState                 string `json:"materialization_state"`
	MaterializationReady                 bool   `json:"materialization_ready"`
	RuntimeOwnerMaterialized             bool   `json:"runtime_owner_materialized"`
	OwnerMaterializationRequired         bool   `json:"owner_materialization_required"`
	RuntimeOwnerServiceSuppliesInputs    bool   `json:"runtime_owner_service_supplies_inputs"`
	DispatchRunnerRequired               bool   `json:"dispatch_runner_required"`
	DispatchAllowed                      bool   `json:"dispatch_allowed"`
	DispatchStarted                      bool   `json:"dispatch_started"`
	RuntimeOwned                         bool   `json:"runtime_owned"`
	RuntimeOwnedDispatch                 bool   `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                      bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool   `json:"kde_policy_owner"`
	KDEForwardsOnlyEvidenceHandle        bool   `json:"kde_forwards_only_evidence_handle"`
	DesktopReceiptFieldsReconstructed    bool   `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess            bool   `json:"desktop_kde_state_root_access"`
	StateRootPathExposed                 bool   `json:"state_root_path_exposed"`
	HandoffPathExposed                   bool   `json:"handoff_path_exposed"`
	ReceiptPathExposed                   bool   `json:"receipt_path_exposed"`
	SessionPathExposed                   bool   `json:"session_path_exposed"`
	ReviewReceiptPathExposed             bool   `json:"review_receipt_path_exposed"`
	RemoteHostExposed                    bool   `json:"remote_host_exposed"`
	RawOutputExposed                     bool   `json:"raw_output_exposed"`
	RuntimeArgvExposed                   bool   `json:"runtime_argv_exposed"`
	RunnerPathExposed                    bool   `json:"runner_path_exposed"`
	BackendDetailsExposed                bool   `json:"backend_details_exposed"`
	HostRootModified                     bool   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool   `json:"privileged_container_required"`
	HostNetworkingRequired               bool   `json:"host_networking_required"`
	DockerSocketMounted                  bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool   `json:"broad_host_mount_required"`
	DesktopLaunchEnabled                 bool   `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                 bool   `json:"backend_launch_enabled"`
	ExecutionStarted                     bool   `json:"execution_started"`
	BackendProcessStarted                bool   `json:"backend_process_started"`
	RecordedAtUTC                        string `json:"recorded_at_utc"`
	BlockedReason                        string `json:"blocked_reason,omitempty"`
	NextOwnerAction                      string `json:"next_owner_action"`
	DesktopSafeSummary                   string `json:"desktop_safe_summary"`
}

type knownAppVerifiedCatalogDispatchRequestPayload struct {
	SchemaVersion                 string   `json:"schema_version"`
	RequestType                   string   `json:"request_type"`
	AppID                         string   `json:"app_id"`
	DisplayName                   string   `json:"display_name"`
	AppVersion                    string   `json:"app_version"`
	LauncherName                  string   `json:"launcher_name"`
	RunnerRequestType             string   `json:"runner_request_type"`
	RunnerArgv                    []string `json:"runner_argv"`
	LaunchAuthorizationReceiptID  string   `json:"launch_authorization_receipt_id"`
	SessionGatedReviewReceiptID   string   `json:"session_gated_review_receipt_id"`
	ControlledExecutionSessionID  string   `json:"controlled_execution_session_id"`
	ControlledSessionRelativePath string   `json:"controlled_session_relative_path"`
	GuestBoundary                 string   `json:"guest_boundary"`
	RecordedAtUTC                 string   `json:"recorded_at_utc"`
}

type KnownAppVerifiedCatalogDispatchRequestRecord struct {
	Version                             string   `json:"version"`
	SchemaVersion                       string   `json:"schema_version"`
	RequestType                         string   `json:"request_type"`
	Source                              string   `json:"source"`
	RuntimeMethod                       string   `json:"runtime_method"`
	ReadMethod                          string   `json:"read_method"`
	AppID                               string   `json:"app_id"`
	DisplayName                         string   `json:"display_name"`
	AppVersion                          string   `json:"app_version"`
	HandoffConsumed                     bool     `json:"handoff_consumed"`
	HandoffRelativePath                 string   `json:"handoff_relative_path"`
	MaterializationState                string   `json:"materialization_state"`
	MaterializationReady                bool     `json:"materialization_ready"`
	RuntimeOwnerMaterialized            bool     `json:"runtime_owner_materialized"`
	LaunchAuthorizationReceiptID        string   `json:"launch_authorization_receipt_id,omitempty"`
	SessionGatedReviewReceiptID         string   `json:"session_gated_review_receipt_id,omitempty"`
	ControlledExecutionSessionID        string   `json:"controlled_execution_session_id,omitempty"`
	ControlledSessionRelativePath       string   `json:"controlled_session_relative_path,omitempty"`
	ControlledSessionDigestVerified     bool     `json:"controlled_session_digest_verified"`
	DispatchRequestRecordState          string   `json:"dispatch_request_record_state"`
	DispatchRequestWritten              bool     `json:"dispatch_request_written"`
	DispatchRequestID                   string   `json:"dispatch_request_id,omitempty"`
	DispatchRequestRelativePath         string   `json:"dispatch_request_relative_path,omitempty"`
	DispatchRequestSHA256               string   `json:"dispatch_request_sha256,omitempty"`
	DispatchRunnerRequestType           string   `json:"dispatch_runner_request_type"`
	DispatchRunnerName                  string   `json:"dispatch_runner_name"`
	DispatchRunnerArgumentCount         int      `json:"dispatch_runner_argument_count"`
	DispatchRunnerArgumentNames         []string `json:"dispatch_runner_argument_names,omitempty"`
	DispatchRunnerArgumentValuesExposed bool     `json:"dispatch_runner_argument_values_exposed"`
	RuntimeOwnerDispatchInputsReady     bool     `json:"runtime_owner_dispatch_inputs_ready"`
	RuntimeOwnerServiceSuppliesInputs   bool     `json:"runtime_owner_service_supplies_inputs"`
	RuntimeOwned                        bool     `json:"runtime_owned"`
	RuntimeOwnedDispatch                bool     `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                     bool     `json:"go_runtime_backed"`
	KDEForwardsOnlyEvidenceHandle       bool     `json:"kde_forwards_only_evidence_handle"`
	DesktopReceiptFieldsReconstructed   bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess           bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed                bool     `json:"state_root_path_exposed"`
	CacheRootPathExposed                bool     `json:"cache_root_path_exposed"`
	DispatchRequestPathExposed          bool     `json:"dispatch_request_path_exposed"`
	RunnerPathExposed                   bool     `json:"runner_path_exposed"`
	BackendDetailsExposed               bool     `json:"backend_details_exposed"`
	RawOutputExposed                    bool     `json:"raw_output_exposed"`
	RawCommandExposed                   bool     `json:"raw_command_exposed"`
	HostRootModified                    bool     `json:"host_root_modified"`
	PrivilegedContainerRequired         bool     `json:"privileged_container_required"`
	HostNetworkingRequired              bool     `json:"host_networking_required"`
	DockerSocketMounted                 bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired              bool     `json:"broad_host_mount_required"`
	DispatchAllowed                     bool     `json:"dispatch_allowed"`
	DispatchStarted                     bool     `json:"dispatch_started"`
	DesktopLaunchEnabled                bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                bool     `json:"backend_launch_enabled"`
	ExecutionStarted                    bool     `json:"execution_started"`
	BackendProcessStarted               bool     `json:"backend_process_started"`
	GuestTransportRequired              bool     `json:"guest_transport_required"`
	RunnerTransportArgsDeferred         bool     `json:"runner_transport_args_deferred"`
	RecordedAtUTC                       string   `json:"recorded_at_utc"`
	BlockedReason                       string   `json:"blocked_reason,omitempty"`
	NextOwnerAction                     string   `json:"next_owner_action"`
	DesktopSafeSummary                  string   `json:"desktop_safe_summary"`
}

func RecordKnownAppVerifiedCatalogLaunchMaterialization(request KnownAppVerifiedCatalogLaunchMaterializationRequest) (KnownAppVerifiedCatalogLaunchMaterializationRecord, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	relativePath := filepath.ToSlash(strings.TrimSpace(request.HandoffRelativePath))
	if relativePath == "" {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires --handoff-relative-path")
	}
	handoffPath, err := knownAppVerifiedCatalogLaunchHandoffPathFromRelativePath(stateRoot, relativePath)
	if err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	content, err := os.ReadFile(handoffPath)
	if err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, fmt.Errorf("read known app verified catalog launch handoff: %w", err)
	}
	var handoff knownAppVerifiedCatalogLaunchHandoffPayload
	if err := json.Unmarshal(content, &handoff); err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, fmt.Errorf("parse known app verified catalog launch handoff: %w", err)
	}
	if err := validateKnownAppVerifiedCatalogLaunchHandoffPayload(handoff); err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()
	recordedAtText := recordedAt.Format(time.RFC3339)
	guestBoundary := strings.TrimSpace(request.GuestBoundary)
	if guestBoundary == "" {
		guestBoundary = winapp.KnownDispatchGuestBoundary
	}

	launchReceipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:           handoff.AppID,
		StateRoot:       stateRoot,
		Authorize:       KnownAppLaunchAuthorizationReceiptAction,
		EvidenceSource:  handoff.AcceptanceType,
		CenterCardState: "validated-verified-catalog-real-runtime-run",
		RecordedAtUTC:   recordedAt,
	})
	if err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	sessionRecord, err := RecordKnownAppControlledExecutionSession(KnownAppControlledExecutionSessionRecordRequest{
		AppID:         handoff.AppID,
		StateRoot:     stateRoot,
		ReceiptID:     launchReceipt.ReceiptID,
		CacheRoot:     request.CacheRoot,
		GuestBoundary: guestBoundary,
	})
	if err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}

	record := baseKnownAppVerifiedCatalogLaunchMaterializationRecord(handoff, relativePath, sha256Hex(string(content)), launchReceipt, sessionRecord, recordedAtText)
	if !sessionRecord.SessionHandoffReady || sessionRecord.RecordState != "persisted" {
		record.MaterializationState = "blocked-managed-artifact-required"
		record.MaterializationReady = false
		record.RuntimeOwnerMaterialized = false
		record.DispatchRunnerRequired = false
		record.BlockedReason = "managed application artifact preparation is required before session materialization"
		record.NextOwnerAction = "prepare-managed-artifact"
		record.DesktopSafeSummary = handoff.DisplayName + " verified catalog launch materialization is blocked until the Runtime prepares the managed application artifact."
		return validateKnownAppVerifiedCatalogLaunchMaterializationRecord(record)
	}

	reviewReceipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:         handoff.AppID,
		StateRoot:     stateRoot,
		SessionID:     sessionRecord.ExecutionSessionID,
		ActionID:      KnownAppSessionGatedLaunchReviewAction,
		Decision:      "approved",
		RecordedAtUTC: recordedAt,
	})
	if err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	record.SessionGatedReviewReceiptID = reviewReceipt.ReceiptID
	record.SessionGatedReviewReceiptRecorded = reviewReceipt.ReviewReceiptRecorded
	record.MaterializationState = "persisted-runtime-owned-launch-session"
	record.MaterializationReady = true
	record.RuntimeOwnerMaterialized = true
	record.DispatchRunnerRequired = true
	record.NextOwnerAction = "dispatch-runner-consume-session"
	record.DesktopSafeSummary = handoff.DisplayName + " verified catalog launch handoff materialized Runtime-owned receipts and a controlled execution session; execution still waits for the dispatch runner."
	return validateKnownAppVerifiedCatalogLaunchMaterializationRecord(record)
}

func RecordKnownAppVerifiedCatalogDispatchRequest(request KnownAppVerifiedCatalogDispatchRequestRecordRequest) (KnownAppVerifiedCatalogDispatchRequestRecord, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()
	materialization, err := RecordKnownAppVerifiedCatalogLaunchMaterialization(KnownAppVerifiedCatalogLaunchMaterializationRequest{
		StateRoot:           stateRoot,
		HandoffRelativePath: request.HandoffRelativePath,
		CacheRoot:           request.CacheRoot,
		GuestBoundary:       request.GuestBoundary,
		RecordedAtUTC:       recordedAt,
	})
	if err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	record, err := RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization(stateRoot, request.CacheRoot, request.LauncherName, request.GuestBoundary, materialization, recordedAt)
	if err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	return record, nil
}

func RecordKnownAppVerifiedCatalogDispatchRequestFromMaterialization(stateRoot string, cacheRoot string, launcherName string, guestBoundary string, materialization KnownAppVerifiedCatalogLaunchMaterializationRecord, recordedAt time.Time) (KnownAppVerifiedCatalogDispatchRequestRecord, error) {
	stateRoot = strings.TrimSpace(stateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	if _, err := validateKnownAppVerifiedCatalogLaunchMaterializationRecord(materialization); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()
	recordedAtText := recordedAt.Format(time.RFC3339)
	if strings.TrimSpace(guestBoundary) == "" {
		guestBoundary = winapp.KnownDispatchGuestBoundary
	}
	if strings.TrimSpace(launcherName) == "" {
		launcherName = "xnix-compat-launch"
	}
	for _, value := range []string{cacheRoot, guestBoundary, launcherName} {
		if strings.ContainsAny(value, "\r\n") {
			return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires single-line owner inputs")
		}
	}
	base := baseKnownAppVerifiedCatalogDispatchRequestRecord(materialization, launcherName, recordedAtText)
	if !materialization.MaterializationReady {
		base.DispatchRequestRecordState = "blocked-materialization-required"
		base.BlockedReason = "verified catalog launch materialization must be ready before recording a dispatch request"
		base.NextOwnerAction = "materialize-launch-session"
		base.DesktopSafeSummary = materialization.DisplayName + " dispatch request is blocked until Runtime-owned materialization is ready."
		return validateKnownAppVerifiedCatalogDispatchRequestRecord(base)
	}
	if strings.ContainsAny(launcherName, "\r\n") {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires a single-line launcher name")
	}
	argv := []string{
		launcherName,
		"--app", materialization.AppID,
		"--cache-root", strings.TrimSpace(cacheRoot),
		"--guest-boundary", guestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", materialization.LaunchAuthorizationReceiptID,
		"--review-receipt-id", materialization.SessionGatedReviewReceiptID,
		"--session-id", materialization.ControlledExecutionSessionID,
	}
	payload := knownAppVerifiedCatalogDispatchRequestPayload{
		SchemaVersion:                 KnownAppVerifiedCatalogDispatchRequestSchemaVersion,
		RequestType:                   KnownAppVerifiedCatalogDispatchRequestType,
		AppID:                         materialization.AppID,
		DisplayName:                   materialization.DisplayName,
		AppVersion:                    materialization.AppVersion,
		LauncherName:                  launcherName,
		RunnerRequestType:             winapp.KnownDispatchSmokeRequestType,
		RunnerArgv:                    argv,
		LaunchAuthorizationReceiptID:  materialization.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:   materialization.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID:  materialization.ControlledExecutionSessionID,
		ControlledSessionRelativePath: materialization.ControlledSessionRelativePath,
		GuestBoundary:                 guestBoundary,
		RecordedAtUTC:                 recordedAtText,
	}
	encoded, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, fmt.Errorf("encode known app verified catalog dispatch request: %w", err)
	}
	encoded = append(encoded, '\n')
	relativePath := knownAppVerifiedCatalogDispatchRequestRelativePath(materialization.AppID, materialization.AppVersion)
	fullPath, err := knownAppVerifiedCatalogDispatchRequestPath(stateRoot, relativePath)
	if err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o700); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, fmt.Errorf("create known app verified catalog dispatch request directory: %w", err)
	}
	if err := os.WriteFile(fullPath, encoded, 0o600); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, fmt.Errorf("write known app verified catalog dispatch request: %w", err)
	}
	base.DispatchRequestRecordState = "persisted-runtime-owned-dispatch-request"
	base.DispatchRequestWritten = true
	base.DispatchRequestID = "known-app-verified-catalog-dispatch-request-" + stateRootNamespace(materialization.AppID) + "-" + stateRootNamespace(materialization.AppVersion)
	base.DispatchRequestRelativePath = relativePath
	base.DispatchRequestSHA256 = sha256Hex(string(encoded))
	base.DispatchRunnerArgumentCount = len(argv)
	base.DispatchRunnerArgumentNames = []string{"launcher", "--app", "--cache-root", "--guest-boundary", "--state-root", "--receipt-id", "--review-receipt-id", "--session-id"}
	base.RuntimeOwnerDispatchInputsReady = true
	base.RuntimeOwnedDispatch = true
	base.NextOwnerAction = "dispatch-runner-execute-request"
	base.DesktopSafeSummary = materialization.DisplayName + " verified catalog dispatch request is recorded for Runtime-owner runner execution."
	return validateKnownAppVerifiedCatalogDispatchRequestRecord(base)
}

func baseKnownAppVerifiedCatalogLaunchMaterializationRecord(handoff knownAppVerifiedCatalogLaunchHandoffPayload, relativePath string, digest string, launchReceipt KnownAppLaunchAuthorizationReceiptPreview, sessionRecord KnownAppControlledExecutionSessionRecord, recordedAtText string) KnownAppVerifiedCatalogLaunchMaterializationRecord {
	return KnownAppVerifiedCatalogLaunchMaterializationRecord{
		Version:                              handoff.Version,
		SchemaVersion:                        KnownAppVerifiedCatalogLaunchMaterializationSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogLaunchMaterializationRequestType,
		Source:                               KnownAppVerifiedCatalogLaunchHandoffRequestType + "+runtime-owner-materialization",
		RuntimeMethod:                        "RecordKnownAppVerifiedCatalogLaunchMaterialization",
		ReadMethod:                           "GetKnownAppVerifiedCatalogLaunchMaterialization",
		AppID:                                handoff.AppID,
		DisplayName:                          handoff.DisplayName,
		AppVersion:                           handoff.AppVersion,
		HandoffConsumed:                      true,
		HandoffRelativePath:                  relativePath,
		HandoffDigestVerified:                digest != "",
		AcceptanceRequestType:                handoff.AcceptanceRequestType,
		AcceptanceType:                       handoff.AcceptanceType,
		AcceptanceReady:                      handoff.AcceptanceReady,
		RunPlanMatched:                       handoff.RunPlanMatched,
		ExistingWindowsApp:                   handoff.ExistingWindowsApp,
		KnownPortableCatalogBacked:           handoff.KnownPortableCatalogBacked,
		LaunchAttempted:                      handoff.LaunchAttempted,
		ChecksumVerified:                     handoff.ChecksumVerified,
		MarkerObserved:                       handoff.MarkerObserved,
		RuntimeStartedIsolatedGuest:          handoff.RuntimeStartedIsolatedGuest,
		IsolatedGuestExecutionObserved:       handoff.IsolatedGuestExecutionObserved,
		CompatibilityEngineExecutionObserved: handoff.CompatibilityEngineExecutionObserved,
		OutputRedacted:                       handoff.OutputRedacted,
		Q4ExecutionObserved:                  handoff.Q4ExecutionObserved,
		HostCompilationAvoided:               handoff.HostCompilationAvoided,
		LaunchAuthorizationReceiptID:         launchReceipt.ReceiptID,
		LaunchAuthorizationReceiptRecorded:   launchReceipt.LaunchAuthorizationRecorded,
		ControlledExecutionSessionID:         sessionRecord.ExecutionSessionID,
		ControlledSessionRecordState:         sessionRecord.RecordState,
		ControlledSessionRelativePath:        sessionRecord.SessionRelativePath,
		ControlledSessionDigestVerified:      sessionRecord.SessionSHA256 != "" && sessionRecord.SessionRecordWritten,
		MaterializationState:                 "pending",
		MaterializationReady:                 false,
		RuntimeOwnerMaterialized:             false,
		OwnerMaterializationRequired:         true,
		RuntimeOwnerServiceSuppliesInputs:    true,
		DispatchRunnerRequired:               false,
		DispatchAllowed:                      false,
		DispatchStarted:                      false,
		RuntimeOwned:                         true,
		RuntimeOwnedDispatch:                 sessionRecord.RuntimeOwnedDispatchRequest,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		KDEForwardsOnlyEvidenceHandle:        true,
		DesktopReceiptFieldsReconstructed:    false,
		DesktopKDEStateRootAccess:            false,
		StateRootPathExposed:                 false,
		HandoffPathExposed:                   false,
		ReceiptPathExposed:                   false,
		SessionPathExposed:                   false,
		ReviewReceiptPathExposed:             false,
		RemoteHostExposed:                    false,
		RawOutputExposed:                     false,
		RuntimeArgvExposed:                   false,
		RunnerPathExposed:                    false,
		BackendDetailsExposed:                false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		DesktopLaunchEnabled:                 false,
		BackendLaunchEnabled:                 false,
		ExecutionStarted:                     false,
		BackendProcessStarted:                false,
		RecordedAtUTC:                        recordedAtText,
		NextOwnerAction:                      "prepare-managed-artifact",
		DesktopSafeSummary:                   handoff.DisplayName + " verified catalog launch materialization is pending Runtime owner processing.",
	}
}

func baseKnownAppVerifiedCatalogDispatchRequestRecord(materialization KnownAppVerifiedCatalogLaunchMaterializationRecord, launcherName string, recordedAtText string) KnownAppVerifiedCatalogDispatchRequestRecord {
	return KnownAppVerifiedCatalogDispatchRequestRecord{
		Version:                             materialization.Version,
		SchemaVersion:                       KnownAppVerifiedCatalogDispatchRequestSchemaVersion,
		RequestType:                         KnownAppVerifiedCatalogDispatchRequestType,
		Source:                              KnownAppVerifiedCatalogLaunchMaterializationRequestType + "+dispatch-runner-request",
		RuntimeMethod:                       "RecordKnownAppVerifiedCatalogDispatchRequest",
		ReadMethod:                          "GetKnownAppVerifiedCatalogDispatchRequest",
		AppID:                               materialization.AppID,
		DisplayName:                         materialization.DisplayName,
		AppVersion:                          materialization.AppVersion,
		HandoffConsumed:                     materialization.HandoffConsumed,
		HandoffRelativePath:                 materialization.HandoffRelativePath,
		MaterializationState:                materialization.MaterializationState,
		MaterializationReady:                materialization.MaterializationReady,
		RuntimeOwnerMaterialized:            materialization.RuntimeOwnerMaterialized,
		LaunchAuthorizationReceiptID:        materialization.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:         materialization.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID:        materialization.ControlledExecutionSessionID,
		ControlledSessionRelativePath:       materialization.ControlledSessionRelativePath,
		ControlledSessionDigestVerified:     materialization.ControlledSessionDigestVerified,
		DispatchRequestRecordState:          "blocked-materialization-required",
		DispatchRequestWritten:              false,
		DispatchRunnerRequestType:           winapp.KnownDispatchSmokeRequestType,
		DispatchRunnerName:                  filepath.Base(strings.TrimSpace(launcherName)),
		DispatchRunnerArgumentCount:         0,
		DispatchRunnerArgumentValuesExposed: false,
		RuntimeOwnerDispatchInputsReady:     false,
		RuntimeOwnerServiceSuppliesInputs:   true,
		RuntimeOwned:                        true,
		RuntimeOwnedDispatch:                false,
		GoRuntimeBacked:                     true,
		KDEForwardsOnlyEvidenceHandle:       true,
		DesktopReceiptFieldsReconstructed:   false,
		DesktopKDEStateRootAccess:           false,
		StateRootPathExposed:                false,
		CacheRootPathExposed:                false,
		DispatchRequestPathExposed:          false,
		RunnerPathExposed:                   false,
		BackendDetailsExposed:               false,
		RawOutputExposed:                    false,
		RawCommandExposed:                   false,
		HostRootModified:                    false,
		PrivilegedContainerRequired:         false,
		HostNetworkingRequired:              false,
		DockerSocketMounted:                 false,
		BroadHostMountRequired:              false,
		DispatchAllowed:                     false,
		DispatchStarted:                     false,
		DesktopLaunchEnabled:                false,
		BackendLaunchEnabled:                false,
		ExecutionStarted:                    false,
		BackendProcessStarted:               false,
		GuestTransportRequired:              true,
		RunnerTransportArgsDeferred:         true,
		RecordedAtUTC:                       recordedAtText,
		NextOwnerAction:                     "materialize-launch-session",
		DesktopSafeSummary:                  materialization.DisplayName + " dispatch request is waiting for Runtime-owned materialization.",
	}
}

func knownAppVerifiedCatalogDispatchRequestRelativePath(appID string, appVersion string) string {
	name := "known-app-verified-catalog-dispatch-request-" + stateRootNamespace(appID) + "-" + stateRootNamespace(appVersion) + ".json"
	return filepath.ToSlash(filepath.Join("runtime", "known-app-verified-catalog-dispatch-requests", name))
}

func knownAppVerifiedCatalogDispatchRequestPath(stateRoot string, relativePath string) (string, error) {
	relativePath = filepath.ToSlash(strings.TrimSpace(relativePath))
	if relativePath == "" || filepath.IsAbs(relativePath) || strings.Contains(filepath.Clean(relativePath), "..") {
		return "", errors.New("known app verified catalog dispatch request requires a safe relative path")
	}
	if !strings.HasPrefix(relativePath, "runtime/known-app-verified-catalog-dispatch-requests/") || !strings.HasSuffix(relativePath, ".json") {
		return "", errors.New("known app verified catalog dispatch request requires a dispatch request relative path")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(cleanRoot, filepath.FromSlash(relativePath))
	cleanFull, err := filepath.Abs(filepath.Clean(fullPath))
	if err != nil {
		return "", err
	}
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		return "", errors.New("known app verified catalog dispatch request escaped state root")
	}
	return cleanFull, nil
}

func knownAppVerifiedCatalogLaunchHandoffPathFromRelativePath(stateRoot string, relativePath string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app verified catalog launch handoff read requires a state root")
	}
	relativePath = filepath.ToSlash(strings.TrimSpace(relativePath))
	if relativePath == "" || filepath.IsAbs(relativePath) || strings.Contains(filepath.Clean(relativePath), "..") {
		return "", errors.New("known app verified catalog launch handoff read requires a safe relative path")
	}
	if !strings.HasPrefix(relativePath, knownAppVerifiedCatalogLaunchHandoffRecordDir+"/") || !strings.HasSuffix(relativePath, ".json") {
		return "", errors.New("known app verified catalog launch handoff read requires a launch handoff relative path")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(cleanRoot, filepath.FromSlash(relativePath))
	cleanFull, err := filepath.Abs(filepath.Clean(fullPath))
	if err != nil {
		return "", err
	}
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		return "", errors.New("known app verified catalog launch handoff read escaped state root")
	}
	return cleanFull, nil
}

func validateKnownAppVerifiedCatalogLaunchHandoffPayload(handoff knownAppVerifiedCatalogLaunchHandoffPayload) error {
	switch {
	case handoff.SchemaVersion != KnownAppVerifiedCatalogLaunchHandoffSchemaVersion || handoff.RequestType != KnownAppVerifiedCatalogLaunchHandoffRequestType:
		return errors.New("known app verified catalog launch materialization requires launch handoff input")
	case handoff.AppID == "" || handoff.DisplayName == "" || handoff.AppVersion == "":
		return errors.New("known app verified catalog launch materialization requires known app identity")
	case handoff.AcceptanceRequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType || handoff.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" || !handoff.AcceptanceReady || !handoff.RunPlanMatched:
		return errors.New("known app verified catalog launch materialization requires accepted verified catalog run evidence")
	case !handoff.ExistingWindowsApp || !handoff.KnownPortableCatalogBacked || !handoff.LaunchAttempted || !handoff.ChecksumVerified || !handoff.MarkerObserved || !handoff.RuntimeStartedIsolatedGuest || !handoff.IsolatedGuestExecutionObserved || !handoff.CompatibilityEngineExecutionObserved || !handoff.OutputRedacted || !handoff.Q4ExecutionObserved || !handoff.HostCompilationAvoided:
		return errors.New("known app verified catalog launch materialization requires complete safe handoff evidence")
	}
	for _, value := range []string{handoff.Version, handoff.SchemaVersion, handoff.RequestType, handoff.AppID, handoff.DisplayName, handoff.AppVersion, handoff.AcceptanceRequestType, handoff.AcceptanceType, handoff.RecordedAtUTC} {
		if value != "" && !singleLine(value) {
			return errors.New("known app verified catalog launch materialization requires single-line handoff fields")
		}
	}
	if err := validateNoBackendTerms(handoff, "known app verified catalog launch handoff payload"); err != nil {
		return err
	}
	return nil
}

func validateKnownAppVerifiedCatalogLaunchMaterializationRecord(record KnownAppVerifiedCatalogLaunchMaterializationRecord) (KnownAppVerifiedCatalogLaunchMaterializationRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppVerifiedCatalogLaunchMaterializationSchemaVersion || record.RequestType != KnownAppVerifiedCatalogLaunchMaterializationRequestType:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization has invalid schema")
	case record.AppID == "" || record.DisplayName == "" || record.AppVersion == "":
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires known app identity")
	case !record.HandoffConsumed || record.HandoffRelativePath == "" || !record.HandoffDigestVerified:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires consumed handoff evidence")
	case filepath.IsAbs(record.HandoffRelativePath) || strings.Contains(filepath.Clean(record.HandoffRelativePath), ".."):
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires a safe handoff path")
	case record.AcceptanceRequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType || record.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" || !record.AcceptanceReady || !record.RunPlanMatched:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires accepted run evidence")
	case !record.ExistingWindowsApp || !record.KnownPortableCatalogBacked || !record.LaunchAttempted || !record.ChecksumVerified || !record.MarkerObserved || !record.RuntimeStartedIsolatedGuest || !record.IsolatedGuestExecutionObserved || !record.CompatibilityEngineExecutionObserved || !record.OutputRedacted || !record.Q4ExecutionObserved || !record.HostCompilationAvoided:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires complete safe run evidence")
	case record.LaunchAuthorizationReceiptID == "" || !record.LaunchAuthorizationReceiptRecorded:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires launch authorization receipt")
	case record.MaterializationReady && (record.MaterializationState != "persisted-runtime-owned-launch-session" || !record.RuntimeOwnerMaterialized || !record.SessionGatedReviewReceiptRecorded || record.SessionGatedReviewReceiptID == "" || record.ControlledExecutionSessionID == "" || record.ControlledSessionRelativePath == "" || !record.ControlledSessionDigestVerified || record.ControlledSessionRecordState != "persisted"):
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization ready state requires persisted session and review receipts")
	case !record.MaterializationReady && (record.MaterializationState == "persisted-runtime-owned-launch-session" || record.RuntimeOwnerMaterialized || record.SessionGatedReviewReceiptRecorded):
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization blocked state must not claim persisted readiness")
	case !record.OwnerMaterializationRequired || !record.RuntimeOwnerServiceSuppliesInputs || !record.RuntimeOwned || !record.GoRuntimeBacked || record.KDEPolicyOwner || !record.KDEForwardsOnlyEvidenceHandle || record.DesktopReceiptFieldsReconstructed || record.DesktopKDEStateRootAccess:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization must remain Runtime-owned and KDE evidence-only")
	case record.StateRootPathExposed || record.HandoffPathExposed || record.ReceiptPathExposed || record.SessionPathExposed || record.ReviewReceiptPathExposed || record.RemoteHostExposed || record.RawOutputExposed || record.RuntimeArgvExposed || record.RunnerPathExposed || record.BackendDetailsExposed:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization must not expose paths, backend details, or raw output")
	case record.HostRootModified || record.PrivilegedContainerRequired || record.HostNetworkingRequired || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization must keep host and container boundaries closed")
	case record.DesktopLaunchEnabled || record.BackendLaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted || record.DispatchAllowed || record.DispatchStarted:
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization must not start dispatch, desktop, backend, or execution")
	}
	for _, value := range []string{record.Version, record.SchemaVersion, record.RequestType, record.Source, record.RuntimeMethod, record.ReadMethod, record.AppID, record.DisplayName, record.AppVersion, record.HandoffRelativePath, record.AcceptanceRequestType, record.AcceptanceType, record.LaunchAuthorizationReceiptID, record.ControlledExecutionSessionID, record.ControlledSessionRecordState, record.ControlledSessionRelativePath, record.SessionGatedReviewReceiptID, record.MaterializationState, record.RecordedAtUTC, record.BlockedReason, record.NextOwnerAction, record.DesktopSafeSummary} {
		if value != "" && !singleLine(value) {
			return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, errors.New("known app verified catalog launch materialization requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(record, "known app verified catalog launch materialization"); err != nil {
		return KnownAppVerifiedCatalogLaunchMaterializationRecord{}, err
	}
	return record, nil
}

func validateKnownAppVerifiedCatalogDispatchRequestRecord(record KnownAppVerifiedCatalogDispatchRequestRecord) (KnownAppVerifiedCatalogDispatchRequestRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppVerifiedCatalogDispatchRequestSchemaVersion || record.RequestType != KnownAppVerifiedCatalogDispatchRequestType:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request has invalid schema")
	case record.AppID == "" || record.DisplayName == "" || record.AppVersion == "":
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires known app identity")
	case !record.HandoffConsumed || record.HandoffRelativePath == "" || filepath.IsAbs(record.HandoffRelativePath) || strings.Contains(filepath.Clean(record.HandoffRelativePath), ".."):
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires safe consumed handoff evidence")
	case record.MaterializationReady && (record.MaterializationState != "persisted-runtime-owned-launch-session" || !record.RuntimeOwnerMaterialized || record.LaunchAuthorizationReceiptID == "" || record.SessionGatedReviewReceiptID == "" || record.ControlledExecutionSessionID == "" || record.ControlledSessionRelativePath == "" || !record.ControlledSessionDigestVerified):
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request ready state requires materialized receipts and session")
	case !record.MaterializationReady && (record.DispatchRequestWritten || record.RuntimeOwnerDispatchInputsReady || record.RuntimeOwnedDispatch):
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request blocked state must not claim dispatch readiness")
	case record.DispatchRequestWritten && (record.DispatchRequestRecordState != "persisted-runtime-owned-dispatch-request" || record.DispatchRequestID == "" || record.DispatchRequestRelativePath == "" || record.DispatchRequestSHA256 == "" || record.DispatchRunnerArgumentCount == 0 || !record.RuntimeOwnerDispatchInputsReady || !record.RuntimeOwnedDispatch):
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires persisted request evidence")
	case record.DispatchRequestRelativePath != "" && (filepath.IsAbs(record.DispatchRequestRelativePath) || strings.Contains(filepath.Clean(record.DispatchRequestRelativePath), "..")):
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires a safe relative request path")
	case record.DispatchRunnerRequestType != winapp.KnownDispatchSmokeRequestType || record.DispatchRunnerName == "":
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires the managed dispatch smoke runner")
	case !record.RuntimeOwnerServiceSuppliesInputs || !record.RuntimeOwned || !record.GoRuntimeBacked || !record.KDEForwardsOnlyEvidenceHandle || record.DesktopReceiptFieldsReconstructed || record.DesktopKDEStateRootAccess:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request must remain Runtime-owned and KDE evidence-only")
	case record.StateRootPathExposed || record.CacheRootPathExposed || record.DispatchRequestPathExposed || record.RunnerPathExposed || record.BackendDetailsExposed || record.RawOutputExposed || record.RawCommandExposed || record.DispatchRunnerArgumentValuesExposed:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request must not expose owner paths, command values, backend details, or raw output")
	case record.HostRootModified || record.PrivilegedContainerRequired || record.HostNetworkingRequired || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request must keep host and container boundaries closed")
	case record.DispatchAllowed || record.DispatchStarted || record.DesktopLaunchEnabled || record.BackendLaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request record must not start dispatch, desktop, backend, or execution")
	case !record.GuestTransportRequired || !record.RunnerTransportArgsDeferred:
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request must defer guest transport details to the Runtime owner")
	}
	for _, value := range []string{record.Version, record.SchemaVersion, record.RequestType, record.Source, record.RuntimeMethod, record.ReadMethod, record.AppID, record.DisplayName, record.AppVersion, record.HandoffRelativePath, record.MaterializationState, record.LaunchAuthorizationReceiptID, record.SessionGatedReviewReceiptID, record.ControlledExecutionSessionID, record.ControlledSessionRelativePath, record.DispatchRequestRecordState, record.DispatchRequestID, record.DispatchRequestRelativePath, record.DispatchRequestSHA256, record.DispatchRunnerRequestType, record.DispatchRunnerName, record.RecordedAtUTC, record.BlockedReason, record.NextOwnerAction, record.DesktopSafeSummary} {
		if value != "" && !singleLine(value) {
			return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires single-line fields")
		}
	}
	for _, value := range record.DispatchRunnerArgumentNames {
		if !singleLine(value) {
			return KnownAppVerifiedCatalogDispatchRequestRecord{}, errors.New("known app verified catalog dispatch request requires single-line argument names")
		}
	}
	if err := validateNoBackendTerms(record, "known app verified catalog dispatch request"); err != nil {
		return KnownAppVerifiedCatalogDispatchRequestRecord{}, err
	}
	return record, nil
}
