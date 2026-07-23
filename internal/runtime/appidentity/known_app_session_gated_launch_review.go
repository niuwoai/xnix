package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	KnownAppSessionGatedLaunchReviewSchemaVersion        = "xnix.runtime.known_app_session_gated_launch_review.v1"
	KnownAppSessionGatedLaunchReviewRequestType          = "known-app-session-gated-launch-review-preview"
	KnownAppSessionGatedLaunchReviewReceiptSchemaVersion = "xnix.runtime.known_app_session_gated_launch_review_receipt.v1"
	KnownAppSessionGatedLaunchReviewReceiptRequestType   = "known-app-session-gated-launch-review-receipt-record"
	KnownAppSessionGatedLaunchReviewAction               = "review-session-gated-dispatch"
	knownAppSessionGatedLaunchReviewReceiptDir           = "runtime/session-gated-launch-review-receipts"
)

type KnownAppSessionGatedLaunchReviewRequest struct {
	AppID     string
	StateRoot string
	SessionID string
	ActionID  string
	Decision  string
}

type KnownAppSessionGatedLaunchReviewPreview struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	ReviewType                  string `json:"review_type"`
	Source                      string `json:"source"`
	Desktop                     string `json:"desktop"`
	RuntimeMethod               string `json:"runtime_method"`
	ReadMethod                  string `json:"read_method"`
	AppID                       string `json:"app_id"`
	DisplayName                 string `json:"display_name"`
	AppVersion                  string `json:"app_version"`
	ActionID                    string `json:"action_id"`
	ActionKind                  string `json:"action_kind"`
	Decision                    string `json:"decision"`
	DecisionAccepted            bool   `json:"decision_accepted"`
	ReviewRouteID               string `json:"review_route_id"`
	ReviewRouteCreated          bool   `json:"review_route_created"`
	ReviewRouteRequestType      string `json:"review_route_request_type"`
	ReviewRouteRuntimeMethod    string `json:"review_route_runtime_method"`
	ReviewRouteReadMethod       string `json:"review_route_read_method"`
	ReadBeforeWriteRequired     bool   `json:"read_before_write_required"`
	RuntimeReceiptRequired      bool   `json:"runtime_receipt_required"`
	ExecutionSessionID          string `json:"execution_session_id"`
	SessionRecordConsumed       bool   `json:"session_record_consumed"`
	SessionDigestVerified       bool   `json:"session_digest_verified"`
	SessionRelativePath         string `json:"session_relative_path"`
	SessionSHA256               string `json:"session_sha256"`
	SessionState                string `json:"session_state"`
	CompatibilityCenterState    string `json:"compatibility_center_state"`
	RuntimeOwnerConsumable      bool   `json:"runtime_owner_consumable"`
	KDEReadModelConsumable      bool   `json:"kde_read_model_consumable"`
	SafeForKDE                  bool   `json:"safe_for_kde"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	UserReviewCaptured          bool   `json:"user_review_captured"`
	UserDecisionAllowsLaunch    bool   `json:"user_decision_allows_launch"`
	RuntimeLaunchApproval       bool   `json:"runtime_launch_approval"`
	LaunchAllowed               bool   `json:"launch_allowed"`
	LaunchEnabled               bool   `json:"launch_enabled"`
	DesktopLaunchEnabled        bool   `json:"desktop_launch_enabled"`
	ExecutionStarted            bool   `json:"execution_started"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	BackendProcessStarted       bool   `json:"backend_process_started"`
	RequestObjectsCreated       bool   `json:"request_objects_created"`
	PermissionGrantCreated      bool   `json:"permission_grant_created"`
	ReviewReceiptRecorded       bool   `json:"review_receipt_recorded"`
	StateRootPathExposed        bool   `json:"state_root_path_exposed"`
	SessionPathExposed          bool   `json:"session_path_exposed"`
	RawArtifactPathExposed      bool   `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	HostRootModified            bool   `json:"host_root_modified"`
	NetworkRequired             bool   `json:"network_required"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	NextStep                    string `json:"next_step"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
}

type KnownAppSessionGatedLaunchReviewReceiptRequest struct {
	AppID         string
	StateRoot     string
	SessionID     string
	ActionID      string
	Decision      string
	RecordedAtUTC time.Time
}

type KnownAppSessionGatedLaunchReviewReceiptRecord struct {
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	ReviewType                  string `json:"review_type"`
	Source                      string `json:"source"`
	RuntimeMethod               string `json:"runtime_method"`
	ReadMethod                  string `json:"read_method"`
	AppID                       string `json:"app_id"`
	DisplayName                 string `json:"display_name"`
	AppVersion                  string `json:"app_version"`
	ActionID                    string `json:"action_id"`
	ActionKind                  string `json:"action_kind"`
	Decision                    string `json:"decision"`
	DecisionAccepted            bool   `json:"decision_accepted"`
	ReadBeforeWriteConsumed     bool   `json:"read_before_write_consumed"`
	ReviewRouteID               string `json:"review_route_id"`
	ReviewRouteConsumed         bool   `json:"review_route_consumed"`
	ReviewRouteRequestType      string `json:"review_route_request_type"`
	ReviewRouteRuntimeMethod    string `json:"review_route_runtime_method"`
	ReviewRouteReadMethod       string `json:"review_route_read_method"`
	RuntimeReceiptRequired      bool   `json:"runtime_receipt_required"`
	ExecutionSessionID          string `json:"execution_session_id"`
	SessionRecordConsumed       bool   `json:"session_record_consumed"`
	SessionDigestVerified       bool   `json:"session_digest_verified"`
	SessionRelativePath         string `json:"session_relative_path"`
	SessionSHA256               string `json:"session_sha256"`
	SessionState                string `json:"session_state"`
	CompatibilityCenterState    string `json:"compatibility_center_state"`
	ReceiptID                   string `json:"receipt_id"`
	ReceiptRelativePath         string `json:"receipt_relative_path"`
	ReceiptSHA256               string `json:"receipt_sha256"`
	ReceiptState                string `json:"receipt_state"`
	ReviewReceiptRecorded       bool   `json:"review_receipt_recorded"`
	RuntimeOwnerConsumable      bool   `json:"runtime_owner_consumable"`
	KDEReadModelConsumable      bool   `json:"kde_read_model_consumable"`
	SafeForKDE                  bool   `json:"safe_for_kde"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	UserReviewCaptured          bool   `json:"user_review_captured"`
	UserDecisionAllowsLaunch    bool   `json:"user_decision_allows_launch"`
	RuntimeLaunchApproval       bool   `json:"runtime_launch_approval"`
	LaunchAllowed               bool   `json:"launch_allowed"`
	LaunchEnabled               bool   `json:"launch_enabled"`
	DesktopLaunchEnabled        bool   `json:"desktop_launch_enabled"`
	ExecutionStarted            bool   `json:"execution_started"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	BackendProcessStarted       bool   `json:"backend_process_started"`
	RequestObjectsCreated       bool   `json:"request_objects_created"`
	PermissionGrantCreated      bool   `json:"permission_grant_created"`
	StateRootPathExposed        bool   `json:"state_root_path_exposed"`
	SessionPathExposed          bool   `json:"session_path_exposed"`
	ReceiptPathExposed          bool   `json:"receipt_path_exposed"`
	RawArtifactPathExposed      bool   `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	HostRootModified            bool   `json:"host_root_modified"`
	NetworkRequired             bool   `json:"network_required"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	RecordedAtUTC               string `json:"recorded_at_utc"`
	NextStep                    string `json:"next_step"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
}

type knownAppSessionGatedLaunchReviewReceiptFile struct {
	SchemaVersion           string `json:"schema_version"`
	ReceiptType             string `json:"receipt_type"`
	ReceiptID               string `json:"receipt_id"`
	AppID                   string `json:"app_id"`
	DisplayName             string `json:"display_name"`
	AppVersion              string `json:"app_version"`
	ActionID                string `json:"action_id"`
	ActionKind              string `json:"action_kind"`
	Decision                string `json:"decision"`
	ReviewState             string `json:"review_state"`
	ExecutionSessionID      string `json:"execution_session_id"`
	SessionRelativePath     string `json:"session_relative_path"`
	SessionSHA256           string `json:"session_sha256"`
	SessionDigestVerified   bool   `json:"session_digest_verified"`
	ReadBeforeWriteConsumed bool   `json:"read_before_write_consumed"`
	RuntimeOwned            bool   `json:"runtime_owned"`
	GoRuntimeBacked         bool   `json:"go_runtime_backed"`
	KDEPolicyOwner          bool   `json:"kde_policy_owner"`
	RuntimeLaunchApproval   bool   `json:"runtime_launch_approval"`
	LaunchAllowed           bool   `json:"launch_allowed"`
	DesktopLaunchEnabled    bool   `json:"desktop_launch_enabled"`
	BackendLaunchEnabled    bool   `json:"backend_launch_enabled"`
	ExecutionStarted        bool   `json:"execution_started"`
	HostRootModified        bool   `json:"host_root_modified"`
	StateRootPathExposed    bool   `json:"state_root_path_exposed"`
	SessionPathExposed      bool   `json:"session_path_exposed"`
	ReceiptPathExposed      bool   `json:"receipt_path_exposed"`
	RawArtifactPathExposed  bool   `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed   bool   `json:"backend_details_exposed"`
	RecordedAtUTC           string `json:"recorded_at_utc"`
}

func PreviewKnownAppSessionGatedLaunchReview(request KnownAppSessionGatedLaunchReviewRequest) (KnownAppSessionGatedLaunchReviewPreview, error) {
	actionID := strings.TrimSpace(request.ActionID)
	if actionID == "" {
		actionID = KnownAppSessionGatedLaunchReviewAction
	}
	if actionID != KnownAppSessionGatedLaunchReviewAction {
		return KnownAppSessionGatedLaunchReviewPreview{}, fmt.Errorf("known app session-gated launch review requires --action %s", KnownAppSessionGatedLaunchReviewAction)
	}
	decision := strings.TrimSpace(request.Decision)
	if decision == "" {
		decision = "reviewed"
	}
	if !singleLine(decision) {
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires a single-line decision")
	}
	if !supportedExecutionReviewDecision(decision) {
		return KnownAppSessionGatedLaunchReviewPreview{}, fmt.Errorf("unsupported known app session-gated launch review decision: %s", decision)
	}

	session, err := PreviewKnownAppControlledExecutionSessionConsumption(KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     request.AppID,
		StateRoot: request.StateRoot,
		SessionID: request.SessionID,
	})
	if err != nil {
		return KnownAppSessionGatedLaunchReviewPreview{}, err
	}

	preview := KnownAppSessionGatedLaunchReviewPreview{
		SchemaVersion:               KnownAppSessionGatedLaunchReviewSchemaVersion,
		RequestType:                 KnownAppSessionGatedLaunchReviewRequestType,
		ReviewType:                  "known-app-session-gated-launch-review",
		Source:                      KnownAppControlledExecutionSessionConsumeRequestType + "+kde-center-page-session-gate-card",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "PreviewKnownAppSessionGatedLaunchReview",
		ReadMethod:                  "GetKnownAppSessionGatedLaunchReview",
		AppID:                       session.AppID,
		DisplayName:                 session.DisplayName,
		AppVersion:                  session.AppVersion,
		ActionID:                    actionID,
		ActionKind:                  "session-gate-review",
		Decision:                    decision,
		DecisionAccepted:            true,
		ReviewRouteID:               "runtime-owned-session-gated-launch-review",
		ReviewRouteCreated:          true,
		ReviewRouteRequestType:      KnownAppSessionGatedLaunchReviewRequestType,
		ReviewRouteRuntimeMethod:    "PreviewKnownAppSessionGatedLaunchReview",
		ReviewRouteReadMethod:       "GetKnownAppSessionGatedLaunchReview",
		ReadBeforeWriteRequired:     true,
		RuntimeReceiptRequired:      true,
		ExecutionSessionID:          session.ExecutionSessionID,
		SessionRecordConsumed:       session.SessionRecordConsumed,
		SessionDigestVerified:       session.SessionDigestVerified,
		SessionRelativePath:         session.SessionRelativePath,
		SessionSHA256:               session.SessionSHA256,
		SessionState:                session.SessionState,
		CompatibilityCenterState:    session.CompatibilityCenterState,
		RuntimeOwnerConsumable:      session.RuntimeOwnerConsumable,
		KDEReadModelConsumable:      session.KDEReadModelConsumable,
		SafeForKDE:                  session.SafeForKDE,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserReviewCaptured:          true,
		UserDecisionAllowsLaunch:    false,
		RuntimeLaunchApproval:       false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		DesktopLaunchEnabled:        false,
		ExecutionStarted:            false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		RequestObjectsCreated:       false,
		PermissionGrantCreated:      false,
		ReviewReceiptRecorded:       false,
		StateRootPathExposed:        false,
		SessionPathExposed:          false,
		RawArtifactPathExposed:      false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		NextStep:                    knownAppSessionGatedLaunchReviewNextStep(decision),
		DesktopSafeSummary:          session.DisplayName + " session-gated launch review route is ready for a Runtime-owned read-before-write review; no approval receipt is recorded and no execution is started.",
	}
	return validateKnownAppSessionGatedLaunchReviewPreview(preview)
}

func RecordKnownAppSessionGatedLaunchReviewReceipt(request KnownAppSessionGatedLaunchReviewReceiptRequest) (KnownAppSessionGatedLaunchReviewReceiptRecord, error) {
	preview, err := PreviewKnownAppSessionGatedLaunchReview(KnownAppSessionGatedLaunchReviewRequest{
		AppID:     request.AppID,
		StateRoot: request.StateRoot,
		SessionID: request.SessionID,
		ActionID:  request.ActionID,
		Decision:  request.Decision,
	})
	if err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, err
	}
	if !preview.SessionRecordConsumed || !preview.SessionDigestVerified || !preview.ReadBeforeWriteRequired {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires digest-verified read-before-write evidence")
	}

	receiptID := KnownAppSessionGatedLaunchReviewReceiptID(preview.AppID, preview.AppVersion, preview.ExecutionSessionID)
	receiptRelativePath := KnownAppSessionGatedLaunchReviewReceiptRelativePath(receiptID)
	receiptPath, err := KnownAppSessionGatedLaunchReviewReceiptPath(request.StateRoot, receiptID)
	if err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(receiptPath), 0o700); err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, fmt.Errorf("create session-gated launch review receipt store: %w", err)
	}
	recordedAt := request.RecordedAtUTC.UTC()
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}

	receipt := knownAppSessionGatedLaunchReviewReceiptFile{
		SchemaVersion:           KnownAppSessionGatedLaunchReviewReceiptSchemaVersion,
		ReceiptType:             "runtime-owned-session-gated-launch-review",
		ReceiptID:               receiptID,
		AppID:                   preview.AppID,
		DisplayName:             preview.DisplayName,
		AppVersion:              preview.AppVersion,
		ActionID:                preview.ActionID,
		ActionKind:              preview.ActionKind,
		Decision:                preview.Decision,
		ReviewState:             "recorded-dispatch-still-gated",
		ExecutionSessionID:      preview.ExecutionSessionID,
		SessionRelativePath:     preview.SessionRelativePath,
		SessionSHA256:           preview.SessionSHA256,
		SessionDigestVerified:   true,
		ReadBeforeWriteConsumed: true,
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		RuntimeLaunchApproval:   false,
		LaunchAllowed:           false,
		DesktopLaunchEnabled:    false,
		BackendLaunchEnabled:    false,
		ExecutionStarted:        false,
		HostRootModified:        false,
		StateRootPathExposed:    false,
		SessionPathExposed:      false,
		ReceiptPathExposed:      false,
		RawArtifactPathExposed:  false,
		BackendDetailsExposed:   false,
		RecordedAtUTC:           recordedAt.Format(time.RFC3339),
	}
	content, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, fmt.Errorf("encode session-gated launch review receipt: %w", err)
	}
	content = append(content, '\n')
	if err := os.WriteFile(receiptPath, content, 0o600); err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, fmt.Errorf("write session-gated launch review receipt: %w", err)
	}

	record := KnownAppSessionGatedLaunchReviewReceiptRecord{
		SchemaVersion:               KnownAppSessionGatedLaunchReviewReceiptSchemaVersion,
		RequestType:                 KnownAppSessionGatedLaunchReviewReceiptRequestType,
		ReviewType:                  "known-app-session-gated-launch-review-receipt",
		Source:                      KnownAppSessionGatedLaunchReviewRequestType + "+runtime-review-receipt-store",
		RuntimeMethod:               "RecordKnownAppSessionGatedLaunchReviewReceipt",
		ReadMethod:                  "GetKnownAppSessionGatedLaunchReviewReceipt",
		AppID:                       preview.AppID,
		DisplayName:                 preview.DisplayName,
		AppVersion:                  preview.AppVersion,
		ActionID:                    preview.ActionID,
		ActionKind:                  preview.ActionKind,
		Decision:                    preview.Decision,
		DecisionAccepted:            true,
		ReadBeforeWriteConsumed:     true,
		ReviewRouteID:               preview.ReviewRouteID,
		ReviewRouteConsumed:         true,
		ReviewRouteRequestType:      preview.ReviewRouteRequestType,
		ReviewRouteRuntimeMethod:    preview.ReviewRouteRuntimeMethod,
		ReviewRouteReadMethod:       preview.ReviewRouteReadMethod,
		RuntimeReceiptRequired:      true,
		ExecutionSessionID:          preview.ExecutionSessionID,
		SessionRecordConsumed:       true,
		SessionDigestVerified:       true,
		SessionRelativePath:         preview.SessionRelativePath,
		SessionSHA256:               preview.SessionSHA256,
		SessionState:                preview.SessionState,
		CompatibilityCenterState:    preview.CompatibilityCenterState,
		ReceiptID:                   receiptID,
		ReceiptRelativePath:         receiptRelativePath,
		ReceiptSHA256:               sha256Hex(string(content)),
		ReceiptState:                "recorded-dispatch-still-gated",
		ReviewReceiptRecorded:       true,
		RuntimeOwnerConsumable:      true,
		KDEReadModelConsumable:      true,
		SafeForKDE:                  true,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserReviewCaptured:          true,
		UserDecisionAllowsLaunch:    false,
		RuntimeLaunchApproval:       false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		DesktopLaunchEnabled:        false,
		ExecutionStarted:            false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		RequestObjectsCreated:       false,
		PermissionGrantCreated:      false,
		StateRootPathExposed:        false,
		SessionPathExposed:          false,
		ReceiptPathExposed:          false,
		RawArtifactPathExposed:      false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RecordedAtUTC:               recordedAt.Format(time.RFC3339),
		NextStep:                    knownAppSessionGatedLaunchReviewReceiptNextStep(preview.Decision),
		DesktopSafeSummary:          preview.DisplayName + " session-gated launch review receipt was recorded after digest-verified session evidence; dispatch remains blocked until a later Runtime launch gate consumes the receipt.",
	}
	return validateKnownAppSessionGatedLaunchReviewReceiptRecord(record)
}

func knownAppSessionGatedLaunchReviewNextStep(decision string) string {
	switch decision {
	case "approved":
		return "Record a Runtime-owned launch approval receipt in a later write step before any dispatch."
	case "rejected":
		return "Keep the session-gated dispatch blocked and show rejection guidance."
	case "deferred":
		return "Keep the session-gated dispatch pending for later Runtime review."
	default:
		return "Show the digest-verified session evidence and keep dispatch blocked until an explicit Runtime receipt exists."
	}
}

func knownAppSessionGatedLaunchReviewReceiptNextStep(decision string) string {
	switch decision {
	case "approved":
		return "Consume the Runtime-owned review receipt in a later launch gate before dispatch."
	case "rejected":
		return "Keep the session-gated dispatch blocked and surface rejection guidance."
	case "deferred":
		return "Keep the session-gated dispatch pending until another Runtime review receipt replaces this state."
	default:
		return "Keep dispatch blocked until a later Runtime launch gate consumes the recorded review receipt."
	}
}

func KnownAppSessionGatedLaunchReviewReceiptID(appID string, version string, sessionID string) string {
	id := stateRootNamespace(strings.TrimSpace(appID))
	if id == "" {
		id = "known-app"
	}
	versionID := stateRootNamespace(strings.TrimSpace(version))
	if versionID == "" {
		versionID = "unknown"
	}
	sessionPart := stateRootNamespace(strings.TrimSpace(sessionID))
	if sessionPart == "" {
		sessionPart = "unknown-session"
	}
	return "known-app-session-gated-launch-review-" + id + "-" + versionID + "-" + sessionPart
}

func KnownAppSessionGatedLaunchReviewReceiptRelativePath(receiptID string) string {
	return filepath.ToSlash(filepath.Join(knownAppSessionGatedLaunchReviewReceiptDir, receiptID+".json"))
}

func KnownAppSessionGatedLaunchReviewReceiptPath(stateRoot string, receiptID string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app session-gated launch review receipt path requires a state root")
	}
	if strings.TrimSpace(receiptID) == "" {
		return "", errors.New("known app session-gated launch review receipt path requires a receipt id")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", fmt.Errorf("resolve managed state root: %w", err)
	}
	relative := filepath.FromSlash(KnownAppSessionGatedLaunchReviewReceiptRelativePath(receiptID))
	candidate := filepath.Join(cleanRoot, relative)
	cleanCandidate := filepath.Clean(candidate)
	rel, err := filepath.Rel(cleanRoot, cleanCandidate)
	if err != nil {
		return "", fmt.Errorf("scope session-gated launch review receipt path: %w", err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel) {
		return "", errors.New("session-gated launch review receipt path escaped the managed state root")
	}
	return cleanCandidate, nil
}

func validateKnownAppSessionGatedLaunchReviewPreview(preview KnownAppSessionGatedLaunchReviewPreview) (KnownAppSessionGatedLaunchReviewPreview, error) {
	switch {
	case preview.SchemaVersion != KnownAppSessionGatedLaunchReviewSchemaVersion || preview.RequestType != KnownAppSessionGatedLaunchReviewRequestType:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review has invalid schema")
	case preview.ActionID != KnownAppSessionGatedLaunchReviewAction || preview.ActionKind != "session-gate-review":
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires the session gate review action")
	case !preview.DecisionAccepted || !preview.ReviewRouteCreated || !preview.ReadBeforeWriteRequired || !preview.RuntimeReceiptRequired:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires a read-before-write review route")
	case !preview.SessionRecordConsumed || !preview.SessionDigestVerified || preview.SessionRelativePath == "" || preview.SessionSHA256 == "":
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires digest-verified session evidence")
	case !preview.RuntimeOwnerConsumable || !preview.KDEReadModelConsumable || !preview.SafeForKDE || !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires Runtime-owned KDE-safe evidence")
	case preview.UserDecisionAllowsLaunch || preview.RuntimeLaunchApproval || preview.LaunchAllowed || preview.LaunchEnabled || preview.DesktopLaunchEnabled || preview.ExecutionStarted:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review must not approve or start launch")
	case preview.BackendLaunchEnabled || preview.BackendProcessStarted || preview.RequestObjectsCreated || preview.PermissionGrantCreated || preview.ReviewReceiptRecorded:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review must not write Runtime objects")
	case preview.StateRootPathExposed || preview.SessionPathExposed || preview.RawArtifactPathExposed || preview.BackendDetailsExposed:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review must not expose raw paths or backend details")
	case preview.HostRootModified || preview.NetworkRequired || preview.PrivilegedContainerRequired || preview.DockerSocketMounted || preview.BroadHostMountRequired:
		return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review must keep host and container boundaries closed")
	}
	for _, value := range []string{preview.AppID, preview.DisplayName, preview.AppVersion, preview.ExecutionSessionID, preview.SessionState, preview.CompatibilityCenterState, preview.Decision} {
		if !singleLine(value) {
			return KnownAppSessionGatedLaunchReviewPreview{}, errors.New("known app session-gated launch review requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(preview, "known app session-gated launch review"); err != nil {
		return KnownAppSessionGatedLaunchReviewPreview{}, err
	}
	return preview, nil
}

func validateKnownAppSessionGatedLaunchReviewReceiptRecord(record KnownAppSessionGatedLaunchReviewReceiptRecord) (KnownAppSessionGatedLaunchReviewReceiptRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppSessionGatedLaunchReviewReceiptSchemaVersion || record.RequestType != KnownAppSessionGatedLaunchReviewReceiptRequestType:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt has invalid schema")
	case record.ActionID != KnownAppSessionGatedLaunchReviewAction || record.ActionKind != "session-gate-review":
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires the session gate review action")
	case !record.DecisionAccepted || !record.ReadBeforeWriteConsumed || !record.ReviewRouteConsumed || !record.RuntimeReceiptRequired:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires consumed read-before-write evidence")
	case !record.SessionRecordConsumed || !record.SessionDigestVerified || record.SessionRelativePath == "" || record.SessionSHA256 == "":
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires digest-verified session evidence")
	case record.ReceiptID == "" || record.ReceiptRelativePath == "" || record.ReceiptSHA256 == "" || filepath.IsAbs(record.ReceiptRelativePath) || strings.Contains(filepath.Clean(record.ReceiptRelativePath), ".."):
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires relative receipt evidence")
	case record.ReceiptState != "recorded-dispatch-still-gated" || !record.ReviewReceiptRecorded:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires a recorded gated state")
	case !record.RuntimeOwnerConsumable || !record.KDEReadModelConsumable || !record.SafeForKDE || !record.RuntimeOwned || !record.GoRuntimeBacked || record.KDEPolicyOwner:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires Runtime-owned KDE-safe evidence")
	case record.UserDecisionAllowsLaunch || record.RuntimeLaunchApproval || record.LaunchAllowed || record.LaunchEnabled || record.DesktopLaunchEnabled || record.ExecutionStarted:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt must not approve or start launch")
	case record.BackendLaunchEnabled || record.BackendProcessStarted || record.RequestObjectsCreated || record.PermissionGrantCreated:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt must not create launch request objects")
	case record.StateRootPathExposed || record.SessionPathExposed || record.ReceiptPathExposed || record.RawArtifactPathExposed || record.BackendDetailsExposed:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt must not expose raw paths or backend details")
	case record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt must keep host and container boundaries closed")
	}
	for _, value := range []string{record.AppID, record.DisplayName, record.AppVersion, record.ExecutionSessionID, record.SessionState, record.CompatibilityCenterState, record.Decision, record.RecordedAtUTC} {
		if !singleLine(value) {
			return KnownAppSessionGatedLaunchReviewReceiptRecord{}, errors.New("known app session-gated launch review receipt requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(record, "known app session-gated launch review receipt"); err != nil {
		return KnownAppSessionGatedLaunchReviewReceiptRecord{}, err
	}
	return record, nil
}
