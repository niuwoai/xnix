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
	KnownAppVerifiedCatalogLaunchHandoffSchemaVersion = "xnix.runtime.known_app_verified_catalog_launch_handoff.v1"
	KnownAppVerifiedCatalogLaunchHandoffRequestType   = "known-app-verified-catalog-launch-handoff-record"
	knownAppVerifiedCatalogLaunchHandoffRecordDir     = "runtime/known-app-verified-catalog-launch-handoffs"
)

type KnownAppVerifiedCatalogLaunchHandoffRequest struct {
	StateRoot      string
	AcceptancePath string
	RecordedAtUTC  time.Time
}

type KnownAppVerifiedCatalogLaunchHandoffRecord struct {
	Version                              string   `json:"version"`
	SchemaVersion                        string   `json:"schema_version"`
	RequestType                          string   `json:"request_type"`
	Source                               string   `json:"source"`
	RuntimeMethod                        string   `json:"runtime_method"`
	ReadMethod                           string   `json:"read_method"`
	AppID                                string   `json:"app_id"`
	DisplayName                          string   `json:"display_name"`
	AppVersion                           string   `json:"app_version"`
	AcceptanceRequestType                string   `json:"acceptance_request_type"`
	AcceptanceType                       string   `json:"acceptance_type"`
	AcceptanceConsumed                   bool     `json:"acceptance_consumed"`
	AcceptanceReady                      bool     `json:"acceptance_ready"`
	RunPlanMatched                       bool     `json:"run_plan_matched"`
	ExistingWindowsApp                   bool     `json:"existing_windows_app"`
	KnownPortableCatalogBacked           bool     `json:"known_portable_catalog_backed"`
	LaunchAttempted                      bool     `json:"launch_attempted"`
	ChecksumVerified                     bool     `json:"checksum_verified"`
	MarkerObserved                       bool     `json:"marker_observed"`
	RuntimeStartedIsolatedGuest          bool     `json:"runtime_started_isolated_guest"`
	IsolatedGuestExecutionObserved       bool     `json:"isolated_guest_execution_observed"`
	CompatibilityEngineExecutionObserved bool     `json:"compatibility_engine_execution_observed"`
	OutputRedacted                       bool     `json:"output_redacted"`
	Q4ExecutionObserved                  bool     `json:"q4_execution_observed"`
	HostCompilationAvoided               bool     `json:"host_compilation_avoided"`
	HandoffID                            string   `json:"handoff_id"`
	HandoffState                         string   `json:"handoff_state"`
	HandoffRelativePath                  string   `json:"handoff_relative_path"`
	HandoffSHA256                        string   `json:"handoff_sha256"`
	HandoffPayloadDigestVerified         bool     `json:"handoff_payload_digest_verified"`
	DesktopTriggerReady                  bool     `json:"desktop_trigger_ready"`
	DesktopCallableRoute                 string   `json:"desktop_callable_route"`
	DesktopDBusMethod                    string   `json:"desktop_dbus_method"`
	DesktopEvidenceHandleForwarded       bool     `json:"desktop_evidence_handle_forwarded"`
	OwnerServiceCallReady                bool     `json:"owner_service_call_ready"`
	OwnerServiceBoundary                 string   `json:"owner_service_boundary"`
	OwnerServiceMethod                   string   `json:"owner_service_method"`
	OwnerServiceCallType                 string   `json:"owner_service_call_type"`
	OwnerServiceCallArgs                 []string `json:"owner_service_call_args"`
	OwnerServiceCLIArgs                  []string `json:"owner_service_cli_args"`
	OwnerMaterializationRequired         bool     `json:"owner_materialization_required"`
	RuntimeOwnerServiceSuppliesInputs    bool     `json:"runtime_owner_service_supplies_inputs"`
	RuntimeOwned                         bool     `json:"runtime_owned"`
	GoRuntimeBacked                      bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool     `json:"kde_policy_owner"`
	KDEForwardsOnlyEvidenceHandle        bool     `json:"kde_forwards_only_evidence_handle"`
	DesktopReceiptFieldsReconstructed    bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess            bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed                 bool     `json:"state_root_path_exposed"`
	AcceptancePathExposed                bool     `json:"acceptance_path_exposed"`
	HandoffPathExposed                   bool     `json:"handoff_path_exposed"`
	RemoteHostExposed                    bool     `json:"remote_host_exposed"`
	RawOutputExposed                     bool     `json:"raw_output_exposed"`
	RuntimeArgvExposed                   bool     `json:"runtime_argv_exposed"`
	RunnerPathExposed                    bool     `json:"runner_path_exposed"`
	BackendDetailsExposed                bool     `json:"backend_details_exposed"`
	HostRootModified                     bool     `json:"host_root_modified"`
	PrivilegedContainerRequired          bool     `json:"privileged_container_required"`
	HostNetworkingRequired               bool     `json:"host_networking_required"`
	DockerSocketMounted                  bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool     `json:"broad_host_mount_required"`
	DesktopLaunchEnabled                 bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                 bool     `json:"backend_launch_enabled"`
	ExecutionStarted                     bool     `json:"execution_started"`
	BackendProcessStarted                bool     `json:"backend_process_started"`
	RecordedAtUTC                        string   `json:"recorded_at_utc"`
	DesktopSafeSummary                   string   `json:"desktop_safe_summary"`
}

type knownAppVerifiedCatalogLaunchHandoffPayload struct {
	Version                              string `json:"version"`
	SchemaVersion                        string `json:"schema_version"`
	RequestType                          string `json:"request_type"`
	AppID                                string `json:"app_id"`
	DisplayName                          string `json:"display_name"`
	AppVersion                           string `json:"app_version"`
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
	RecordedAtUTC                        string `json:"recorded_at_utc"`
}

func RecordKnownAppVerifiedCatalogLaunchHandoff(request KnownAppVerifiedCatalogLaunchHandoffRequest) (KnownAppVerifiedCatalogLaunchHandoffRecord, error) {
	stateRoot := strings.TrimSpace(request.StateRoot)
	if err := validateKnownAppRuntimeOwnerStateRoot(stateRoot); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	acceptancePath := strings.TrimSpace(request.AcceptancePath)
	if acceptancePath == "" {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires --known-app-verified-catalog-run-acceptance")
	}
	content, err := os.ReadFile(acceptancePath)
	if err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, fmt.Errorf("read known app verified catalog run acceptance: %w", err)
	}
	var acceptance KnownAppVerifiedCatalogRunAcceptancePreview
	if err := json.Unmarshal(content, &acceptance); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, fmt.Errorf("parse known app verified catalog run acceptance: %w", err)
	}
	if _, err := KnownAppSmokeEvidenceFromKnownAppVerifiedCatalogRunAcceptance(content); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	recordedAt = recordedAt.UTC()
	recordedAtText := recordedAt.Format(time.RFC3339)
	handoffID := KnownAppVerifiedCatalogLaunchHandoffID(acceptance.AppID, acceptance.AppVersion, acceptance.AcceptanceType)
	relativePath := KnownAppVerifiedCatalogLaunchHandoffRelativePath(handoffID)
	payload := knownAppVerifiedCatalogLaunchHandoffPayload{
		Version:                              acceptance.Version,
		SchemaVersion:                        KnownAppVerifiedCatalogLaunchHandoffSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogLaunchHandoffRequestType,
		AppID:                                acceptance.AppID,
		DisplayName:                          acceptance.DisplayName,
		AppVersion:                           acceptance.AppVersion,
		AcceptanceRequestType:                acceptance.RequestType,
		AcceptanceType:                       acceptance.AcceptanceType,
		AcceptanceReady:                      acceptance.AcceptanceReady,
		RunPlanMatched:                       acceptance.RunPlanMatched,
		ExistingWindowsApp:                   acceptance.ExistingWindowsApp,
		KnownPortableCatalogBacked:           acceptance.KnownPortableCatalogBacked,
		LaunchAttempted:                      acceptance.LaunchAttempted,
		ChecksumVerified:                     acceptance.ChecksumVerified,
		MarkerObserved:                       acceptance.MarkerObserved,
		RuntimeStartedIsolatedGuest:          acceptance.RuntimeStartedIsolatedGuest,
		IsolatedGuestExecutionObserved:       acceptance.IsolatedGuestExecutionObserved,
		CompatibilityEngineExecutionObserved: acceptance.CompatibilityEngineExecutionObserved,
		OutputRedacted:                       acceptance.OutputRedacted,
		Q4ExecutionObserved:                  acceptance.Q4ExecutionObserved,
		HostCompilationAvoided:               acceptance.HostCompilationAvoided,
		RecordedAtUTC:                        recordedAtText,
	}
	payloadBytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	payloadBytes = append(payloadBytes, '\n')
	recordPath, err := knownAppVerifiedCatalogLaunchHandoffPath(stateRoot, handoffID)
	if err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(recordPath), 0o700); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	if err := os.WriteFile(recordPath, payloadBytes, 0o600); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	record := KnownAppVerifiedCatalogLaunchHandoffRecord{
		Version:                              acceptance.Version,
		SchemaVersion:                        KnownAppVerifiedCatalogLaunchHandoffSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogLaunchHandoffRequestType,
		Source:                               KnownAppVerifiedCatalogRunAcceptanceRequestType + "+runtime-owner-launch-handoff",
		RuntimeMethod:                        "RecordKnownAppVerifiedCatalogLaunchHandoff",
		ReadMethod:                           "GetKnownAppVerifiedCatalogLaunchHandoff",
		AppID:                                acceptance.AppID,
		DisplayName:                          acceptance.DisplayName,
		AppVersion:                           acceptance.AppVersion,
		AcceptanceRequestType:                acceptance.RequestType,
		AcceptanceType:                       acceptance.AcceptanceType,
		AcceptanceConsumed:                   true,
		AcceptanceReady:                      acceptance.AcceptanceReady,
		RunPlanMatched:                       acceptance.RunPlanMatched,
		ExistingWindowsApp:                   acceptance.ExistingWindowsApp,
		KnownPortableCatalogBacked:           acceptance.KnownPortableCatalogBacked,
		LaunchAttempted:                      acceptance.LaunchAttempted,
		ChecksumVerified:                     acceptance.ChecksumVerified,
		MarkerObserved:                       acceptance.MarkerObserved,
		RuntimeStartedIsolatedGuest:          acceptance.RuntimeStartedIsolatedGuest,
		IsolatedGuestExecutionObserved:       acceptance.IsolatedGuestExecutionObserved,
		CompatibilityEngineExecutionObserved: acceptance.CompatibilityEngineExecutionObserved,
		OutputRedacted:                       acceptance.OutputRedacted,
		Q4ExecutionObserved:                  acceptance.Q4ExecutionObserved,
		HostCompilationAvoided:               acceptance.HostCompilationAvoided,
		HandoffID:                            handoffID,
		HandoffState:                         "persisted-for-runtime-owner-materialization",
		HandoffRelativePath:                  relativePath,
		HandoffSHA256:                        sha256Hex(string(payloadBytes)),
		HandoffPayloadDigestVerified:         true,
		DesktopTriggerReady:                  true,
		DesktopCallableRoute:                 "kde-dbus-verified-catalog-launch-handoff",
		DesktopDBusMethod:                    "org.xnix.Compatibility1.RequestRuntimeOwnedLaunch",
		DesktopEvidenceHandleForwarded:       true,
		OwnerServiceCallReady:                true,
		OwnerServiceBoundary:                 "go-runtime-owner-in-process-service",
		OwnerServiceMethod:                   "RequestRuntimeOwnedLaunch",
		OwnerServiceCallType:                 "desktop-action-dispatch",
		OwnerServiceCallArgs:                 []string{"RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", relativePath},
		OwnerServiceCLIArgs:                  []string{"--service-call", "RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", relativePath},
		OwnerMaterializationRequired:         true,
		RuntimeOwnerServiceSuppliesInputs:    true,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		KDEForwardsOnlyEvidenceHandle:        true,
		DesktopReceiptFieldsReconstructed:    false,
		DesktopKDEStateRootAccess:            false,
		StateRootPathExposed:                 false,
		AcceptancePathExposed:                false,
		HandoffPathExposed:                   false,
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
		DesktopSafeSummary:                   acceptance.DisplayName + " verified catalog run acceptance is persisted as a Runtime-owner launch handoff; the desktop forwards only the handoff handle.",
	}
	return validateKnownAppVerifiedCatalogLaunchHandoffRecord(record)
}

func KnownAppVerifiedCatalogLaunchHandoffID(appID string, version string, acceptanceType string) string {
	appPart := stateRootNamespace(strings.TrimSpace(appID))
	if appPart == "" {
		appPart = "known-app"
	}
	versionPart := stateRootNamespace(strings.TrimSpace(version))
	if versionPart == "" {
		versionPart = "unknown"
	}
	acceptancePart := stateRootNamespace(strings.TrimSpace(acceptanceType))
	if acceptancePart == "" {
		acceptancePart = "acceptance"
	}
	return "known-app-verified-catalog-launch-handoff-" + appPart + "-" + versionPart + "-" + acceptancePart
}

func KnownAppVerifiedCatalogLaunchHandoffRelativePath(handoffID string) string {
	return knownAppVerifiedCatalogLaunchHandoffRecordDir + "/" + stateRootNamespace(handoffID) + ".json"
}

func knownAppVerifiedCatalogLaunchHandoffPath(stateRoot string, handoffID string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app verified catalog launch handoff requires a state root")
	}
	if strings.TrimSpace(handoffID) == "" {
		return "", errors.New("known app verified catalog launch handoff requires a handoff id")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", err
	}
	relative := filepath.FromSlash(KnownAppVerifiedCatalogLaunchHandoffRelativePath(handoffID))
	fullPath := filepath.Join(cleanRoot, relative)
	cleanFull, err := filepath.Abs(filepath.Clean(fullPath))
	if err != nil {
		return "", err
	}
	if cleanFull != cleanRoot && !strings.HasPrefix(cleanFull, cleanRoot+string(os.PathSeparator)) {
		return "", errors.New("known app verified catalog launch handoff escaped state root")
	}
	return cleanFull, nil
}

func validateKnownAppVerifiedCatalogLaunchHandoffRecord(record KnownAppVerifiedCatalogLaunchHandoffRecord) (KnownAppVerifiedCatalogLaunchHandoffRecord, error) {
	switch {
	case record.SchemaVersion != KnownAppVerifiedCatalogLaunchHandoffSchemaVersion || record.RequestType != KnownAppVerifiedCatalogLaunchHandoffRequestType:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff has invalid schema")
	case record.AppID == "" || record.DisplayName == "" || record.AppVersion == "":
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires known app identity")
	case record.AcceptanceRequestType != KnownAppVerifiedCatalogRunAcceptanceRequestType || record.AcceptanceType != "verified-catalog-app-q4-real-run-acceptance" || !record.AcceptanceConsumed || !record.AcceptanceReady || !record.RunPlanMatched:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires accepted verified catalog run evidence")
	case !record.ExistingWindowsApp || !record.KnownPortableCatalogBacked || !record.LaunchAttempted || !record.ChecksumVerified || !record.MarkerObserved || !record.RuntimeStartedIsolatedGuest || !record.IsolatedGuestExecutionObserved || !record.CompatibilityEngineExecutionObserved || !record.OutputRedacted || !record.Q4ExecutionObserved || !record.HostCompilationAvoided:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires complete safe run acceptance evidence")
	case record.HandoffID == "" || record.HandoffState != "persisted-for-runtime-owner-materialization" || record.HandoffRelativePath == "" || record.HandoffSHA256 == "" || !record.HandoffPayloadDigestVerified:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires persisted relative handoff evidence")
	case filepath.IsAbs(record.HandoffRelativePath) || strings.Contains(filepath.Clean(record.HandoffRelativePath), ".."):
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires safe relative handoff evidence")
	case !record.DesktopTriggerReady || record.DesktopCallableRoute != "kde-dbus-verified-catalog-launch-handoff" || record.DesktopDBusMethod != "org.xnix.Compatibility1.RequestRuntimeOwnedLaunch" || !record.DesktopEvidenceHandleForwarded:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires a desktop-callable Runtime owner route")
	case !record.OwnerServiceCallReady || record.OwnerServiceBoundary != "go-runtime-owner-in-process-service" || record.OwnerServiceMethod != "RequestRuntimeOwnedLaunch" || record.OwnerServiceCallType != "desktop-action-dispatch" || !record.OwnerMaterializationRequired || !record.RuntimeOwnerServiceSuppliesInputs:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires Runtime owner materialization")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(record.OwnerServiceCallArgs, []string{"RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", record.HandoffRelativePath}):
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires evidence-only owner service arguments")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(record.OwnerServiceCLIArgs, []string{"--service-call", "RequestRuntimeOwnedLaunch", "verified-catalog-launch-handoff", record.HandoffRelativePath}):
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires evidence-only owner service CLI arguments")
	case !record.RuntimeOwned || !record.GoRuntimeBacked || record.KDEPolicyOwner || !record.KDEForwardsOnlyEvidenceHandle || record.DesktopReceiptFieldsReconstructed || record.DesktopKDEStateRootAccess:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff must keep Runtime ownership and KDE evidence-only")
	case record.StateRootPathExposed || record.AcceptancePathExposed || record.HandoffPathExposed || record.RemoteHostExposed || record.RawOutputExposed || record.RuntimeArgvExposed || record.RunnerPathExposed || record.BackendDetailsExposed:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff must not expose paths, backend details, or raw output")
	case record.HostRootModified || record.PrivilegedContainerRequired || record.HostNetworkingRequired || record.DockerSocketMounted || record.BroadHostMountRequired:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff must keep host and container boundaries closed")
	case record.DesktopLaunchEnabled || record.BackendLaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted:
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff must not start desktop, backend, or execution")
	}
	for _, value := range []string{record.Version, record.SchemaVersion, record.RequestType, record.Source, record.RuntimeMethod, record.ReadMethod, record.AppID, record.DisplayName, record.AppVersion, record.AcceptanceRequestType, record.AcceptanceType, record.HandoffID, record.HandoffState, record.HandoffRelativePath, record.HandoffSHA256, record.DesktopCallableRoute, record.DesktopDBusMethod, record.OwnerServiceBoundary, record.OwnerServiceMethod, record.OwnerServiceCallType, record.RecordedAtUTC, record.DesktopSafeSummary} {
		if value != "" && !singleLine(value) {
			return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires single-line fields")
		}
	}
	for _, value := range append(append([]string{}, record.OwnerServiceCallArgs...), record.OwnerServiceCLIArgs...) {
		if strings.TrimSpace(value) == "" || !singleLine(value) {
			return KnownAppVerifiedCatalogLaunchHandoffRecord{}, errors.New("known app verified catalog launch handoff requires single-line owner service arguments")
		}
	}
	if err := validateNoBackendTerms(record, "known app verified catalog launch handoff"); err != nil {
		return KnownAppVerifiedCatalogLaunchHandoffRecord{}, err
	}
	return record, nil
}
