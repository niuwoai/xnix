package appidentity

import (
	"errors"
	"fmt"
	"strings"
)

const (
	KnownAppSessionGatedLaunchReviewSchemaVersion = "xnix.runtime.known_app_session_gated_launch_review.v1"
	KnownAppSessionGatedLaunchReviewRequestType   = "known-app-session-gated-launch-review-preview"
	KnownAppSessionGatedLaunchReviewAction        = "review-session-gated-dispatch"
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
