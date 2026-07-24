package appidentity

import (
	"errors"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppRuntimeStatusLaunchOwnerFixtureSchemaVersion = "xnix.runtime.known_app_runtime_status_launch_owner_fixture.v1"
	KnownAppRuntimeStatusLaunchOwnerFixtureRequestType   = "known-app-runtime-status-launch-owner-fixture-record"
)

type KnownAppRuntimeStatusLaunchOwnerFixtureRequest struct {
	AppID         string
	StateRoot     string
	CacheRoot     string
	GuestBoundary string
	Status        string
	RecordedAtUTC time.Time
}

type KnownAppRuntimeStatusLaunchOwnerFixtureRecord struct {
	SchemaVersion                      string   `json:"schema_version"`
	RequestType                        string   `json:"request_type"`
	Source                             string   `json:"source"`
	RuntimeMethod                      string   `json:"runtime_method"`
	ReadMethod                         string   `json:"read_method"`
	AppID                              string   `json:"app_id"`
	DisplayName                        string   `json:"display_name"`
	AppVersion                         string   `json:"app_version"`
	GuestBoundary                      string   `json:"guest_boundary"`
	FixtureState                       string   `json:"fixture_state"`
	FixtureReady                       bool     `json:"fixture_ready"`
	SkipReason                         string   `json:"skip_reason,omitempty"`
	LaunchAuthorizationReceiptID       string   `json:"launch_authorization_receipt_id,omitempty"`
	LaunchAuthorizationReceiptRecorded bool     `json:"launch_authorization_receipt_recorded"`
	ControlledExecutionSessionID       string   `json:"controlled_execution_session_id,omitempty"`
	ControlledSessionRelativePath      string   `json:"controlled_session_relative_path,omitempty"`
	ControlledSessionDigestVerified    bool     `json:"controlled_session_digest_verified"`
	SessionGatedReviewReceiptID        string   `json:"session_gated_review_receipt_id,omitempty"`
	SessionGatedReviewReceiptRecorded  bool     `json:"session_gated_review_receipt_recorded"`
	EvidenceID                         string   `json:"evidence_id,omitempty"`
	EvidenceRelativePath               string   `json:"evidence_relative_path,omitempty"`
	EvidenceSHA256                     string   `json:"evidence_sha256,omitempty"`
	ProjectionType                     string   `json:"projection_type,omitempty"`
	CompatibilityCenterProjectionReady bool     `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool     `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidenceReady         bool     `json:"known_app_smoke_evidence_ready"`
	DesktopTriggerReady                bool     `json:"desktop_trigger_ready"`
	DesktopCallableRoute               string   `json:"desktop_callable_route,omitempty"`
	DesktopCallableRuntimeMethod       string   `json:"desktop_callable_runtime_method,omitempty"`
	DesktopCallableExecutionType       string   `json:"desktop_callable_execution_type,omitempty"`
	DesktopEvidenceHandleForwarded     bool     `json:"desktop_evidence_handle_forwarded"`
	DesktopDBusMethod                  string   `json:"desktop_dbus_method,omitempty"`
	OwnerServiceCallReady              bool     `json:"owner_service_call_ready"`
	OwnerServiceBoundary               string   `json:"owner_service_boundary,omitempty"`
	OwnerServiceMethod                 string   `json:"owner_service_method,omitempty"`
	OwnerServiceCallType               string   `json:"owner_service_call_type,omitempty"`
	OwnerServiceCallArgs               []string `json:"owner_service_call_args,omitempty"`
	RuntimeOwnerServiceSuppliesInputs  bool     `json:"runtime_owner_service_supplies_inputs"`
	RuntimeOwned                       bool     `json:"runtime_owned"`
	RuntimeOwnedDispatch               bool     `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                    bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool     `json:"kde_policy_owner"`
	KDEForwardsOnlyEvidenceHandle      bool     `json:"kde_forwards_only_evidence_handle"`
	DesktopReceiptFieldsReconstructed  bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess          bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed               bool     `json:"state_root_path_exposed"`
	EvidencePathExposed                bool     `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed         bool     `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed           bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed              bool     `json:"backend_details_exposed"`
	HostRootModified                   bool     `json:"host_root_modified"`
	NetworkRequired                    bool     `json:"network_required"`
	PrivilegedContainerRequired        bool     `json:"privileged_container_required"`
	DockerSocketMounted                bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool     `json:"broad_host_mount_required"`
	DesktopLaunchEnabled               bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool     `json:"backend_launch_enabled"`
	ExecutionStarted                   bool     `json:"execution_started"`
	BackendProcessStarted              bool     `json:"backend_process_started"`
	RecordedAtUTC                      string   `json:"recorded_at_utc"`
	DesktopSafeSummary                 string   `json:"desktop_safe_summary"`
}

func RecordKnownAppRuntimeStatusLaunchOwnerFixture(request KnownAppRuntimeStatusLaunchOwnerFixtureRequest) (KnownAppRuntimeStatusLaunchOwnerFixtureRecord, error) {
	appID := strings.TrimSpace(request.AppID)
	if appID == "" {
		appID = winapp.DefaultKnownAppID
	}
	app, err := winapp.LookupKnownPortableApp(appID)
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	cacheRoot := strings.TrimSpace(request.CacheRoot)
	if cacheRoot == "" {
		cacheRoot = winapp.DefaultKnownAppCacheRoot
	}
	guestBoundary := strings.TrimSpace(request.GuestBoundary)
	if guestBoundary == "" {
		guestBoundary = winapp.KnownDispatchGuestBoundary
	}
	status := strings.TrimSpace(request.Status)
	if status == "" {
		status = "passed"
	}
	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()

	launchReceipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:         app.ID,
		StateRoot:     stateRoot,
		Authorize:     KnownAppLaunchAuthorizationReceiptAction,
		RecordedAtUTC: recordedAt,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	result := baseKnownAppRuntimeStatusLaunchOwnerFixtureRecord(app, guestBoundary, recordedAt)
	result.LaunchAuthorizationReceiptID = launchReceipt.ReceiptID
	result.LaunchAuthorizationReceiptRecorded = launchReceipt.LaunchAuthorizationRecorded

	sessionRecord, err := RecordKnownAppControlledExecutionSession(KnownAppControlledExecutionSessionRecordRequest{
		AppID:         app.ID,
		StateRoot:     stateRoot,
		ReceiptID:     launchReceipt.ReceiptID,
		CacheRoot:     cacheRoot,
		GuestBoundary: guestBoundary,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	result.ControlledExecutionSessionID = sessionRecord.ExecutionSessionID
	result.ControlledSessionRelativePath = sessionRecord.SessionRelativePath
	result.ControlledSessionDigestVerified = sessionRecord.SessionSHA256 != "" && sessionRecord.SessionRecordWritten
	result.RuntimeOwnedDispatch = sessionRecord.RuntimeOwnedDispatchRequest
	if !sessionRecord.SessionHandoffReady || sessionRecord.RecordState != "persisted" {
		result.FixtureState = "blocked"
		result.SkipReason = "known Windows app artifact unavailable; run `ruby scripts/container.rb fetch-known-winapp` first"
		result.DesktopSafeSummary = app.DisplayName + " controlled launch owner fixture is blocked until the known Windows app artifact is verified in the managed cache."
		return validateKnownAppRuntimeStatusLaunchOwnerFixtureRecord(result)
	}

	reviewReceipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:         app.ID,
		StateRoot:     stateRoot,
		SessionID:     sessionRecord.ExecutionSessionID,
		ActionID:      KnownAppSessionGatedLaunchReviewAction,
		Decision:      "approved",
		RecordedAtUTC: recordedAt,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	result.SessionGatedReviewReceiptID = reviewReceipt.ReceiptID
	result.SessionGatedReviewReceiptRecorded = reviewReceipt.ReviewReceiptRecorded

	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  app.ID,
		DisplayName:                            app.DisplayName,
		AppVersion:                             app.Version,
		RequestType:                            winapp.KnownDispatchSmokeRequestType,
		Status:                                 status,
		GuestBoundary:                          guestBoundary,
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       true,
		MarkerObserved:                         true,
		SessionGatedControlledDispatchConsumed: true,
		SessionGatedControlledDispatchState:    "created-after-session-gated-review",
		SessionGatedReviewReceiptID:            reviewReceipt.ReceiptID,
		LaunchAuthorizationReceiptID:           launchReceipt.ReceiptID,
		LaunchAuthorizationReceiptState:        "recorded",
		LaunchGateState:                        "controlled-dispatch-ready",
		LaunchGateConsumed:                     sessionRecord.ReceiptAccepted,
		LaunchGateReceiptAccepted:              sessionRecord.ReceiptAccepted,
		LaunchGateGuestBoundaryAccepted:        sessionRecord.GuestBoundaryAccepted,
		ControlledDispatchReady:                sessionRecord.ControlledDispatchReady,
		ControlledExecutionSessionConsumed:     true,
		ControlledExecutionSessionID:           sessionRecord.ExecutionSessionID,
		ControlledSessionDigestVerified:        true,
		ControlledSessionRelativePath:          sessionRecord.SessionRelativePath,
		RuntimeOwnerConsumableSession:          true,
		KDEReadModelConsumableSession:          true,
		ControlledSessionLiveStateObserved:     false,
		ControlledSessionRegistered:            false,
		ControlledSessionWindowObserved:        false,
		ControlledSessionHostRootModified:      false,
		ControlledSessionBackendProcessStart:   false,
		HostRootModified:                       false,
		DockerSocketMounted:                    false,
		BroadHostMountRequired:                 false,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	evidenceRecord, err := RecordKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:     stateRoot,
		Projection:    projection,
		RecordedAtUTC: recordedAt,
	})
	if err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	result.FixtureState = "ready"
	result.FixtureReady = true
	result.EvidenceID = evidenceRecord.EvidenceID
	result.EvidenceRelativePath = evidenceRecord.EvidenceRelativePath
	result.EvidenceSHA256 = evidenceRecord.EvidenceSHA256
	result.ProjectionType = evidenceRecord.ProjectionType
	result.CompatibilityCenterProjectionReady = evidenceRecord.CompatibilityCenterProjectionReady
	result.KDECenterProjectionReady = evidenceRecord.KDECenterProjectionReady
	result.KnownAppSmokeEvidenceReady = evidenceRecord.KnownAppSmokeEvidenceReady
	result.DesktopTriggerReady = true
	result.DesktopCallableRoute = "kde-dbus-runtime-status-action"
	result.DesktopCallableRuntimeMethod = "ShowRuntimeControlledLaunch"
	result.DesktopCallableExecutionType = KnownAppKDERuntimeStatusLaunchExecutionRequestType
	result.DesktopEvidenceHandleForwarded = true
	result.DesktopDBusMethod = "org.xnix.Compatibility1.ShowRuntimeControlledLaunch"
	result.OwnerServiceCallReady = true
	result.OwnerServiceBoundary = "go-runtime-owner-in-process-service"
	result.OwnerServiceMethod = "ShowRuntimeControlledLaunch"
	result.OwnerServiceCallType = "desktop-action-dispatch"
	result.OwnerServiceCallArgs = []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", evidenceRecord.EvidenceRelativePath}
	result.RuntimeOwnerServiceSuppliesInputs = true
	result.DesktopSafeSummary = app.DisplayName + " controlled launch owner fixture prepared Runtime-owned evidence handoff state for D-Bus controlled launch without exposing Runtime paths."
	return validateKnownAppRuntimeStatusLaunchOwnerFixtureRecord(result)
}

func baseKnownAppRuntimeStatusLaunchOwnerFixtureRecord(app winapp.KnownPortableApp, guestBoundary string, recordedAt time.Time) KnownAppRuntimeStatusLaunchOwnerFixtureRecord {
	return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{
		SchemaVersion:                     KnownAppRuntimeStatusLaunchOwnerFixtureSchemaVersion,
		RequestType:                       KnownAppRuntimeStatusLaunchOwnerFixtureRequestType,
		Source:                            "known-app-launch-authorization+controlled-session+session-gated-review+runtime-status-evidence",
		RuntimeMethod:                     "RecordKnownAppRuntimeStatusLaunchOwnerFixture",
		ReadMethod:                        "GetKnownAppRuntimeStatusLaunchOwnerFixture",
		AppID:                             app.ID,
		DisplayName:                       app.DisplayName,
		AppVersion:                        app.Version,
		GuestBoundary:                     guestBoundary,
		FixtureState:                      "preparing",
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		KDEForwardsOnlyEvidenceHandle:     true,
		DesktopReceiptFieldsReconstructed: false,
		DesktopKDEStateRootAccess:         false,
		StateRootPathExposed:              false,
		EvidencePathExposed:               false,
		ManagedLauncherPathExposed:        false,
		RawLauncherOutputExposed:          false,
		BackendDetailsExposed:             false,
		HostRootModified:                  false,
		NetworkRequired:                   false,
		PrivilegedContainerRequired:       false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		RecordedAtUTC:                     recordedAt.Format(time.RFC3339),
		DesktopSafeSummary:                app.DisplayName + " controlled launch owner fixture is preparing Runtime-owned handoff state.",
	}
}

func validateKnownAppRuntimeStatusLaunchOwnerFixtureRecord(record KnownAppRuntimeStatusLaunchOwnerFixtureRecord) (KnownAppRuntimeStatusLaunchOwnerFixtureRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppRuntimeStatusLaunchOwnerFixtureSchemaVersion || record.RequestType != KnownAppRuntimeStatusLaunchOwnerFixtureRequestType:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture has invalid schema")
	case record.AppID == "" || record.DisplayName == "" || record.AppVersion == "" || record.GuestBoundary == "":
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture requires known app identity and guest boundary")
	case record.FixtureState == "":
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture requires a fixture state")
	case record.FixtureReady && (record.FixtureState != "ready" || record.EvidenceRelativePath == "" || record.EvidenceSHA256 == "" || record.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence"):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires evidence handoff fields")
	case record.FixtureReady && (!record.CompatibilityCenterProjectionReady || !record.KDECenterProjectionReady || !record.KnownAppSmokeEvidenceReady):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires Center-ready evidence")
	case record.FixtureReady && (record.LaunchAuthorizationReceiptID == "" || !record.LaunchAuthorizationReceiptRecorded || record.ControlledExecutionSessionID == "" || record.SessionGatedReviewReceiptID == "" || !record.SessionGatedReviewReceiptRecorded):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires Runtime-owned receipt and session evidence")
	case record.FixtureReady && (!record.DesktopTriggerReady || record.DesktopCallableRoute != "kde-dbus-runtime-status-action" || record.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" || record.DesktopCallableExecutionType != KnownAppKDERuntimeStatusLaunchExecutionRequestType || !record.DesktopEvidenceHandleForwarded || record.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch"):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires a desktop-callable D-Bus trigger")
	case record.FixtureReady && (!record.OwnerServiceCallReady || record.OwnerServiceBoundary != "go-runtime-owner-in-process-service" || record.OwnerServiceMethod != "ShowRuntimeControlledLaunch" || record.OwnerServiceCallType != "desktop-action-dispatch" || !record.RuntimeOwnerServiceSuppliesInputs):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires an owner service call bridge")
	case record.FixtureReady && !sameRuntimeStatusLaunchOwnerFixtureArgs(record.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", record.EvidenceRelativePath}):
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture readiness requires evidence-only owner service arguments")
	case !record.RuntimeOwned || !record.GoRuntimeBacked || record.KDEPolicyOwner:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture must remain Runtime-owned and Go-backed")
	case !record.KDEForwardsOnlyEvidenceHandle || record.DesktopReceiptFieldsReconstructed || record.DesktopKDEStateRootAccess:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture must keep KDE evidence-only")
	case record.StateRootPathExposed || record.EvidencePathExposed || record.ManagedLauncherPathExposed || record.RawLauncherOutputExposed || record.BackendDetailsExposed:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture must not expose paths, launcher output, or backend details")
	case record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture must keep host and container boundaries closed")
	case record.DesktopLaunchEnabled || record.BackendLaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted:
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture must not start launch or backend processes")
	}
	for _, value := range []string{
		record.SchemaVersion,
		record.RequestType,
		record.Source,
		record.RuntimeMethod,
		record.ReadMethod,
		record.AppID,
		record.DisplayName,
		record.AppVersion,
		record.GuestBoundary,
		record.FixtureState,
		record.SkipReason,
		record.LaunchAuthorizationReceiptID,
		record.ControlledExecutionSessionID,
		record.ControlledSessionRelativePath,
		record.SessionGatedReviewReceiptID,
		record.EvidenceID,
		record.EvidenceRelativePath,
		record.EvidenceSHA256,
		record.ProjectionType,
		record.DesktopCallableRoute,
		record.DesktopCallableRuntimeMethod,
		record.DesktopCallableExecutionType,
		record.DesktopDBusMethod,
		record.OwnerServiceBoundary,
		record.OwnerServiceMethod,
		record.OwnerServiceCallType,
		record.RecordedAtUTC,
		record.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture requires single-line fields")
		}
	}
	for _, value := range record.OwnerServiceCallArgs {
		if !singleLine(value) {
			return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, errors.New("known app Runtime-status launch owner fixture requires single-line owner service arguments")
		}
	}
	if err := validateNoBackendTerms(record, "known app Runtime-status launch owner fixture"); err != nil {
		return KnownAppRuntimeStatusLaunchOwnerFixtureRecord{}, err
	}
	return record, nil
}

func sameRuntimeStatusLaunchOwnerFixtureArgs(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
