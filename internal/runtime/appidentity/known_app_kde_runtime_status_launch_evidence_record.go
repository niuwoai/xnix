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
	KnownAppKDERuntimeStatusLaunchEvidenceRecordSchemaVersion  = "xnix.runtime.known_app_kde_runtime_status_launch_evidence_record.v1"
	KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType    = "known-app-kde-runtime-status-launch-evidence-record"
	KnownAppKDERuntimeStatusLaunchEvidencePreviewSchemaVersion = "xnix.runtime.known_app_kde_runtime_status_launch_evidence_preview.v1"
	KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType   = "known-app-kde-runtime-status-launch-evidence-preview"
	knownAppKDERuntimeStatusLaunchEvidenceRecordDir            = "runtime/kde-runtime-status-launch-evidence"
)

type KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest struct {
	StateRoot     string
	Projection    KnownAppKDERuntimeStatusLaunchDelegatedEvidence
	RecordedAtUTC time.Time
}

type KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
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

type KnownAppKDERuntimeStatusLaunchEvidencePreview struct {
	SchemaVersion                      string                       `json:"schema_version"`
	RequestType                        string                       `json:"request_type"`
	Source                             string                       `json:"source"`
	RuntimeMethod                      string                       `json:"runtime_method"`
	ReadMethod                         string                       `json:"read_method"`
	AppID                              string                       `json:"app_id"`
	DisplayName                        string                       `json:"display_name"`
	AppVersion                         string                       `json:"app_version"`
	EvidenceID                         string                       `json:"evidence_id"`
	EvidenceReadState                  string                       `json:"evidence_read_state"`
	EvidenceHandoffConsumed            bool                         `json:"evidence_handoff_consumed"`
	EvidenceRelativePath               string                       `json:"evidence_relative_path"`
	EvidenceSHA256                     string                       `json:"evidence_sha256"`
	EvidenceDigestVerified             bool                         `json:"evidence_digest_verified"`
	ProjectionType                     string                       `json:"projection_type"`
	CompatibilityCenterProjectionReady bool                         `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                         `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidenceReady         bool                         `json:"known_app_smoke_evidence_ready"`
	KnownAppSmokeEvidence              KnownAppSmokeEvidenceSummary `json:"known_app_smoke_evidence"`
	LaunchAuthorizationReceiptID       string                       `json:"launch_authorization_receipt_id"`
	SessionGatedReviewReceiptID        string                       `json:"session_gated_review_receipt_id"`
	ControlledExecutionSessionID       string                       `json:"controlled_execution_session_id"`
	ControlledSessionRelativePath      string                       `json:"controlled_session_relative_path"`
	RuntimeOwned                       bool                         `json:"runtime_owned"`
	RuntimeOwnedDispatch               bool                         `json:"runtime_owned_dispatch"`
	GoRuntimeBacked                    bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                         `json:"kde_policy_owner"`
	StateRootPathExposed               bool                         `json:"state_root_path_exposed"`
	EvidencePathExposed                bool                         `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed         bool                         `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed           bool                         `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed              bool                         `json:"backend_details_exposed"`
	HostRootModified                   bool                         `json:"host_root_modified"`
	DockerSocketMounted                bool                         `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool                         `json:"broad_host_mount_required"`
	DesktopLaunchEnabled               bool                         `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool                         `json:"backend_launch_enabled"`
	ExecutionStarted                   bool                         `json:"execution_started"`
	BackendProcessStarted              bool                         `json:"backend_process_started"`
	DesktopSafeSummary                 string                       `json:"desktop_safe_summary"`
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

func PreviewKnownAppKDERuntimeStatusLaunchEvidence(request KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest) (KnownAppKDERuntimeStatusLaunchEvidencePreview, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	relativePath, err := knownAppKDERuntimeStatusLaunchEvidenceRequestRelativePath(request)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	evidencePath, err := knownAppKDERuntimeStatusLaunchEvidencePathFromRelativePath(stateRoot, relativePath)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	content, err := os.ReadFile(evidencePath)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	var projection KnownAppKDERuntimeStatusLaunchDelegatedEvidence
	if err := json.Unmarshal(content, &projection); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	projection, err = validateKnownAppKDERuntimeStatusLaunchDelegatedEvidence(projection)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	evidence, err := KnownAppSmokeEvidenceFromKDERuntimeStatusLaunchDelegatedEvidence(projection)
	if err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	evidenceID := strings.TrimSuffix(filepath.Base(relativePath), ".json")
	preview := KnownAppKDERuntimeStatusLaunchEvidencePreview{
		SchemaVersion:                      KnownAppKDERuntimeStatusLaunchEvidencePreviewSchemaVersion,
		RequestType:                        KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType,
		Source:                             KnownAppKDERuntimeStatusLaunchEvidenceRecordRequestType + "+runtime-state-root-handoff-read",
		RuntimeMethod:                      "PreviewKnownAppKDERuntimeStatusLaunchEvidence",
		ReadMethod:                         "GetKnownAppKDERuntimeStatusLaunchEvidence",
		AppID:                              projection.AppID,
		DisplayName:                        projection.DisplayName,
		AppVersion:                         projection.AppVersion,
		EvidenceID:                         evidenceID,
		EvidenceReadState:                  "consumed",
		EvidenceHandoffConsumed:            true,
		EvidenceRelativePath:               filepath.ToSlash(relativePath),
		EvidenceSHA256:                     sha256Hex(string(content)),
		EvidenceDigestVerified:             true,
		ProjectionType:                     projection.ProjectionType,
		CompatibilityCenterProjectionReady: projection.CompatibilityCenterProjectionReady,
		KDECenterProjectionReady:           projection.KDECenterProjectionReady,
		KnownAppSmokeEvidenceReady:         true,
		KnownAppSmokeEvidence:              evidence,
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
		DesktopSafeSummary:                 projection.DisplayName + " Runtime-status launch evidence handoff was consumed for the next KDE-triggered Runtime action without exposing Runtime paths or launcher output.",
	}
	return validateKnownAppKDERuntimeStatusLaunchEvidencePreview(preview)
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

func knownAppKDERuntimeStatusLaunchEvidenceRequestRelativePath(request KnownAppKDERuntimeStatusLaunchEvidencePreviewRequest) (string, error) {
	evidenceID := strings.TrimSpace(request.EvidenceID)
	relativePath := filepath.ToSlash(strings.TrimSpace(request.EvidenceRelativePath))
	if evidenceID == "" && relativePath == "" {
		return "", errors.New("known app KDE Runtime-status launch evidence preview requires an evidence id or relative path")
	}
	if evidenceID != "" && relativePath != "" {
		return "", errors.New("known app KDE Runtime-status launch evidence preview accepts either an evidence id or relative path")
	}
	if evidenceID != "" {
		return KnownAppKDERuntimeStatusLaunchEvidenceRelativePath(evidenceID), nil
	}
	if filepath.IsAbs(relativePath) || strings.Contains(filepath.Clean(relativePath), "..") {
		return "", errors.New("known app KDE Runtime-status launch evidence preview requires a safe relative evidence path")
	}
	if !strings.HasPrefix(relativePath, knownAppKDERuntimeStatusLaunchEvidenceRecordDir+"/") || !strings.HasSuffix(relativePath, ".json") {
		return "", errors.New("known app KDE Runtime-status launch evidence preview requires a Runtime-status launch evidence relative path")
	}
	return relativePath, nil
}

func knownAppKDERuntimeStatusLaunchEvidencePathFromRelativePath(stateRoot string, relativePath string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app KDE Runtime-status launch evidence preview requires a state root")
	}
	relativePath = filepath.ToSlash(strings.TrimSpace(relativePath))
	if relativePath == "" || filepath.IsAbs(relativePath) || strings.Contains(filepath.Clean(relativePath), "..") {
		return "", errors.New("known app KDE Runtime-status launch evidence preview requires a safe relative evidence path")
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
		return "", errors.New("known app KDE Runtime-status launch evidence preview escaped state root")
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

func validateKnownAppKDERuntimeStatusLaunchEvidencePreview(preview KnownAppKDERuntimeStatusLaunchEvidencePreview) (KnownAppKDERuntimeStatusLaunchEvidencePreview, error) {
	switch {
	case preview.SchemaVersion != KnownAppKDERuntimeStatusLaunchEvidencePreviewSchemaVersion || preview.RequestType != KnownAppKDERuntimeStatusLaunchEvidencePreviewRequestType:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview has invalid schema")
	case preview.AppID == "" || preview.DisplayName == "" || preview.AppVersion == "":
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires known app identity")
	case preview.EvidenceID == "" || preview.EvidenceReadState != "consumed" || !preview.EvidenceHandoffConsumed:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires consumed handoff evidence")
	case preview.EvidenceRelativePath == "" || filepath.IsAbs(preview.EvidenceRelativePath) || strings.Contains(filepath.Clean(preview.EvidenceRelativePath), ".."):
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires a safe relative evidence path")
	case preview.EvidenceSHA256 == "" || !preview.EvidenceDigestVerified:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires verified evidence digest")
	case preview.ProjectionType != "known-app-kde-runtime-status-launch-delegated-evidence":
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires delegated evidence projection")
	case !preview.CompatibilityCenterProjectionReady || !preview.KDECenterProjectionReady || !preview.KnownAppSmokeEvidenceReady:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires Center-ready evidence")
	case preview.LaunchAuthorizationReceiptID == "" || preview.SessionGatedReviewReceiptID == "" || preview.ControlledExecutionSessionID == "":
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires opaque Runtime ids")
	case !preview.RuntimeOwned || !preview.RuntimeOwnedDispatch || !preview.GoRuntimeBacked || preview.KDEPolicyOwner:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires Runtime-owned Go evidence")
	case preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ManagedLauncherPathExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview must not expose paths, launcher output, or backend details")
	case preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview must keep host and container boundaries closed")
	case preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted:
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview must not start desktop, backend, or execution")
	}
	for _, value := range []string{preview.SchemaVersion, preview.RequestType, preview.Source, preview.RuntimeMethod, preview.ReadMethod, preview.AppID, preview.DisplayName, preview.AppVersion, preview.EvidenceID, preview.EvidenceReadState, preview.EvidenceRelativePath, preview.EvidenceSHA256, preview.ProjectionType, preview.LaunchAuthorizationReceiptID, preview.SessionGatedReviewReceiptID, preview.ControlledExecutionSessionID, preview.ControlledSessionRelativePath} {
		if value != "" && !singleLine(value) {
			return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, errors.New("known app KDE Runtime-status launch evidence preview requires single-line fields")
		}
	}
	if err := validateNoBackendTerms(preview, "known app KDE Runtime-status launch evidence preview"); err != nil {
		return KnownAppKDERuntimeStatusLaunchEvidencePreview{}, err
	}
	return preview, nil
}
