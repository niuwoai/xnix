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
)

type KnownAppVerifiedCatalogLaunchMaterializationRequest struct {
	StateRoot           string
	HandoffRelativePath string
	CacheRoot           string
	GuestBoundary       string
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
