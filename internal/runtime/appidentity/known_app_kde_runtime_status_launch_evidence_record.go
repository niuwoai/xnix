package appidentity

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion = "xnix.runtime.known_app_kde_runtime_status_launch_evidence_record.v1"
	KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType   = "known-app-kde-runtime-status-launch-evidence-record"
	knownAppKDERuntimeStatusLaunchEvidenceRecordDir           = "runtime/kde-runtime-status-launch-evidence"
)

type KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest struct {
	StateRoot     string
	Projection    KnownAppKDERuntimeStatusLaunchDelegatedEvidence
	RecordedAtUTC time.Time
}

type KnownAppKDERuntimeStatusLaunchEvidenceRecord struct {
	SchemaVersion                      string `json:"schema_version"`
	RequestType                        string `json:"request_type"`
	Source                             string `json:"source"`
	RuntimeMethod                      string `json:"runtime_method"`
	ReadMethod                         string `json:"read_method"`
	AppID                              string `json:"app_id"`
	DisplayName                        string `json:"display_name"`
	AppVersion                         string `json:"app_version"`
	EvidenceID                         string `json:"evidence_id"`
	EvidenceState                      string `json:"evidence_state"`
	EvidenceRelativePath               string `json:"evidence_relative_path"`
	EvidenceSHA256                     string `json:"evidence_sha256"`
	ProjectionType                     string `json:"projection_type"`
	CompatibilityCenterProjectionReady bool   `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool   `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidenceReady         bool   `json:"known_app_smoke_evidence_ready"`
	LaunchAuthorizationReceiptID       string `json:"launch_authorization_receipt_id"`
	SessionGatedReviewReceiptID        string `json:"session_gated_review_receipt_id"`
	ControlledExecutionSessionID       string `json:"controlled_execution_session_id"`
	ControlledSessionRelativePath      string `json:"controlled_session_relative_path"`
	RuntimeOwned                       bool   `json:"runtime_owned"`
	RuntimeOwnedDispatch               bool   `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                    bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool   `json:"kde_policy_owner"`
	StateRootPathExposed               bool   `json:"state_root_path_exposed"`
	EvidencePathExposed                bool   `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed         bool   `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed           bool   `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed              bool   `json:"backend_details_exposed"`
	HostRootModified                   bool   `json:"host_root_modified"`
	DockerSocketMounted                bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool   `json:"broad_host_mount_required"`
	DesktopLaunchEnabled               bool   `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool   `json:"backend_launch_enabled"`
	ExecutionStarted                   bool   `json:"execution_started"`
	BackendProcessStarted              bool   `json:"backend_process_started"`
	RecordedAtUTC                      string `json:"recorded_at_utc"`
	DesktopSafeSummary                 string `json:"desktop_safe_summary"`
}

func RecordKnownAppKDERuntimeStatusLaunchEvidence(request KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest) (KnownAppKDERuntimeStatusLaunchEvidenceRecord, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	projection, err := validateKnownAppKDERuntimeStatusLaunchDelegatedEvidence(request.Projection)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	if _, err := KnownAppSmokeEvidenceFromKDERuntimeStatusLaunchDelegatedEvidence(projection); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()
	evidenceID := KnownAppKDERuntimeStatusLaunchEvidenceID(projection.AppID, projection.AppVersion, projection.ControlledExecutionSessionID)
	relativePath := KnownAppKDERuntimeStatusLaunchEvidenceRelativePath(evidenceID)
	payload, err := json.MarshalIndent(projection, "", "  ")
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	payload = append(payload, '\n')
	recordPath, err := knownAppKDERuntimeStatusLaunchEvidencePath(stateRoot, evidenceID)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(recordPath), 0o700); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	if err := os.WriteFile(recordPath, payload, 0o600); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	record := KnownAppKDERuntimeStatusLaunchEvidenceRecord{
		SchemaVersion:                      KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion,
		RequestType:                        KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType,
		Source:                             projection.ProjectionType + "+runtime-state-root-handoff",
		RuntimeMethod:                      "RecordKnownAppKDERuntimeStatusLaunchEvidence",
		ReadMethod:                         "GetKnownAppKDERuntimeStatusLaunchEvidence",
		AppID:                              projection.AppID,
		DisplayName:                        projection.DisplayName,
		AppVersion:                         projection.AppVersion,
		EvidenceID:                         evidenceID,
		EvidenceState:                      "persisted",
		EvidenceRelativePath:               relativePath,
		EvidenceSHA256:                     sha256Hex(string(payload)),
		ProjectionType:                     projection.ProjectionType,
		CompatibilityCenterProjectionReady: projection.CompatibilityCenterProjectionReady,
		KDECenterProjectionReady:           projection.KDECenterProjectionReady,
		KnownAppSmokeEvidenceReady:         true,
		LaunchAuthorizationReceiptID:       projection.LaunchAuthorizationReceiptID,
		SessionGatedReviewReceiptID:        projection.SessionGatedReviewReceiptID,
		ControlledExecutionSessionID:       projection.ControlledExecutionSessionID,
		ControlledSessionRelativePath:      projection.ControlledSessionRelativePath,
		RuntimeOwned:                       true,
		RuntimeOwnedDispatch:               projection.RuntimeOwnedDispatch,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		StateRootPathExposed:               false,
		EvidencePathExposed:                false,
		ManagedLauncherPathExposed:         false,
		RawLauncherOutputExposed:           false,
		BackendDetailsExposed:              false,
		HostRootModified:                   false,
		DockerSocketMounted:                false,
		BroadHostMountRequired:             false,
		DesktopLaunchEnabled:               false,
		BackendLaunchEnabled:               false,
		ExecutionStarted:                   false,
		BackendProcessStarted:              false,
		RecordedAtUTC:                      recordedAt.Format(time.RFC3339),
		DesktopSafeSummary:                 projection.DisplayName + " Runtime-status launch evidence was recorded as a KDE-readable handoff without exposing Runtime paths or launcher output.",
	}
	return validateKnownAppKDERuntimeStatusLaunchEvidenceRecord(record)
}

func KnownAppKDERuntimeStatusLaunchEvidenceID(appID string, version string, sessionID string) string {
	appPart := stateRootNamespace(strings.TrimSpace(appID))
	if appPart == "" {
		appPart = "known-app"
	}
	versionPart := stateRootNamespace(strings.TrimSpace(version))
	if versionPart == "" {
		versionPart = "unknown"
	}
	sessionPart := stateRootNamespace(strings.TrimSpace(sessionID))
	if sessionPart == "" {
		sessionPart = "session"
	}
	return "known-app-kde-runtime-status-launch-evidence-" + appPart + "-" + versionPart + "-" + sessionPart
}

func KnownAppKDERuntimeStatusLaunchEvidenceRelativePath(evidenceID string) string {
	return knownAppKDERuntimeStatusLaunchEvidenceRecordDir + "/" + stateRootNamespace(evidenceID) + ".json"
}

func knownAppKDERuntimeStatusLaunchEvidencePath(stateRoot string, evidenceID string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app KDE Runtime-status launch evidence record requires a state root")
	}
	if strings.TrimSpace(evidenceID) == "" {
		return "", errors.New("known app KDE Runtime-status launch evidence record requires an evidence id")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", err
	}
	relative := filepath.FromSlash(KnownAppKDERuntimeStatusLaunchEvidenceRelativePath(evidenceID))
	fullPath := filepath.Join(cleanRoot, relative)
	cleanFull, err := filepath.Abs(filepath.Clean(fullPath))
	if err != nil {
		return "", err
	}
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		return "", errors.New("known app KDE Runtime-status launch evidence record escaped state root")
	}
	return cleanFull, nil
}

func validateKnownAppKDERuntimeStatusLaunchEvidenceRecord(record KnownAppKDERuntimeStatusLaunchEvidenceRecord) (KnownAppKDERuntimeStatusLaunchEvidenceRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion || record.RequestType != KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record has invalid schema")
	case record.AppID == "" || record.DisplayName == "" || record.AppVersion == "":
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires known app identity")
	case record.EvidenceID == "" || record.EvidenceState != "persisted" || record.EvidenceRelativePath == "" || record.EvidenceSHA256 == "":
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires persisted relative evidence")
	case filepath.IsAbs(record.EvidenceRelativePath) || strings.Contains(filepath.Clean(record.EvidenceRelativePath), ".."):
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires a safe relative evidence path")
	case record.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence":
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires delegated evidence projection")
	case !record.CompatibilityCenterProjectionReady || !record.KDECenterProjectionReady || !record.KnownAppSmokeEvidenceReady:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires Center-ready evidence")
	case record.LaunchAuthorizationReceiptID == "" || record.SessionGatedReviewReceiptID == "" || record.ControlledExecutionSessionID == "":
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires opaque Runtime ids")
	case !record.RuntimeOwned || !record.RuntimeOwnedDispatch || !record.GoRuntimeBacked || record.KDEPolicyOwner:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires Runtime-owned Go evidence")
	case record.StateRootPathExposed || record.EvidencePathExposed || record.ManagedLauncherPathExposed || record.RawLauncherOutputExposed || record.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record must not expose paths, launcher output, or backend details")
	case record.HostRootModified || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record must keep host and container boundaries closed")
	case record.DesktopLaunchEnabled || record.BackendLaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted:
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record must not start desktop, backend, or execution")
	}
	for _, value := range []string{record.SchemaVersion, record.RequestType, record.Source, record.RuntimeMethod, record.ReadMethod, record.AppID, record.DisplayName, record.AppVersion, record.EvidenceID, record.EvidenceState, record.EvidenceRelativePath, record.EvidenceSHA256, record.ProjectionType, record.LaunchAuthorizationReceiptID, record.SessionGatedReviewReceiptID, record.ControlledExecutionSessionID, record.ControlledSessionRelativePath, record.RecordedAtUTC} {
		if value != "" && !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, errors.New("known app KDE Runtime-status launch evidence record requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(record, "known app KDE Runtime-status launch evidence record"); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidenceRecord{}, err
	}
	return record, nil
}
