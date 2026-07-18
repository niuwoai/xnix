package owner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type BackendAdapterRedactedProfileOwnerSmokeCoveragePreview struct {
	Version                       string                                            `json:"version"`
	SchemaVersion                 string                                            `json:"schema_version"`
	RequestType                   string                                            `json:"request_type"`
	CoverageType                  string                                            `json:"coverage_type"`
	Source                        string                                            `json:"source"`
	RuntimeMethod                 string                                            `json:"runtime_method"`
	ReadMethod                    string                                            `json:"read_method"`
	OwnerRouteRequestType         string                                            `json:"owner_route_request_type"`
	OwnerRouteCommand             string                                            `json:"owner_route_command"`
	SmokeBatchRecordCount         int                                               `json:"smoke_batch_record_count"`
	SmokeReadDispatchRecordCount  int                                               `json:"smoke_read_dispatch_record_count"`
	SmokeWriteDenialRecordCount   int                                               `json:"smoke_write_denial_record_count"`
	CoverageRecordFound           bool                                              `json:"coverage_record_found"`
	CoverageRecordSequence        int                                               `json:"coverage_record_sequence"`
	ServiceCallReadDispatch       bool                                              `json:"service_call_read_dispatch"`
	ServiceCallDispatchReady      bool                                              `json:"service_call_dispatch_ready"`
	DispatchReadOnly              bool                                              `json:"dispatch_read_only"`
	DispatchRouteReady            bool                                              `json:"dispatch_route_ready"`
	DispatchRouteSource           string                                            `json:"dispatch_route_source"`
	DispatchGoCommand             string                                            `json:"dispatch_go_command"`
	AuditType                     string                                            `json:"audit_type"`
	RouteDecision                 string                                            `json:"route_decision"`
	ProfileCount                  int                                               `json:"profile_count"`
	ProfileIDs                    []string                                          `json:"profile_ids"`
	ContractProfileCount          int                                               `json:"contract_profile_count"`
	FullContractConsumed          bool                                              `json:"full_contract_consumed"`
	KDEFacingProjectionConsumed   bool                                              `json:"kde_facing_projection_consumed"`
	InternalAdapterIDsRedacted    bool                                              `json:"internal_adapter_ids_redacted"`
	InternalProfilePathsRedacted  bool                                              `json:"internal_profile_paths_redacted"`
	OwnerLocalRouteCandidateReady bool                                              `json:"owner_local_route_candidate_ready"`
	FullContractFixtureLocal      bool                                              `json:"full_contract_fixture_local"`
	ProductionDBusExposureReady   bool                                              `json:"production_dbus_exposure_ready"`
	CallerStateRootRequired       bool                                              `json:"caller_state_root_required"`
	CheckCount                    int                                               `json:"check_count"`
	PassedCheckCount              int                                               `json:"passed_check_count"`
	AllChecksPassed               bool                                              `json:"all_checks_passed"`
	Checks                        []BackendAdapterRedactedProfileSmokeCoverageCheck `json:"checks"`
	CheckIDs                      []string                                          `json:"check_ids"`
	RuntimeOwned                  bool                                              `json:"runtime_owned"`
	GoRuntimeBacked               bool                                              `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                                              `json:"kde_policy_owner"`
	ReadOnlyCoverage              bool                                              `json:"read_only_coverage"`
	SmokeCoverageReady            bool                                              `json:"smoke_coverage_ready"`
	SystemServiceStarted          bool                                              `json:"system_service_started"`
	SessionBusClaimed             bool                                              `json:"session_bus_claimed"`
	ProductionBusClaimed          bool                                              `json:"production_bus_claimed"`
	WriteMethodsEnabled           bool                                              `json:"write_methods_enabled"`
	StateRootWritesEnabled        bool                                              `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled          bool                                              `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled      bool                                              `json:"adapter_invocation_enabled"`
	BackendInstallEnabled         bool                                              `json:"backend_install_enabled"`
	BackendDownloadEnabled        bool                                              `json:"backend_download_enabled"`
	BackendLaunchEnabled          bool                                              `json:"backend_launch_enabled"`
	BackendProcessStarted         bool                                              `json:"backend_process_started"`
	CommandMaterialized           bool                                              `json:"command_materialized"`
	ExecutablePathResolved        bool                                              `json:"executable_path_resolved"`
	NetworkRequired               bool                                              `json:"network_required"`
	HostRootModified              bool                                              `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                                              `json:"privileged_container_required"`
	StateRootPathExposed          bool                                              `json:"state_root_path_exposed"`
	RawCommandExposed             bool                                              `json:"raw_command_exposed"`
	RawExecutableExposed          bool                                              `json:"raw_executable_exposed"`
	BackendDetailsExposed         bool                                              `json:"backend_details_exposed"`
	BlockedActions                []string                                          `json:"blocked_actions"`
	NextRequirements              []string                                          `json:"next_requirements"`
	DesktopSafeSummary            string                                            `json:"desktop_safe_summary"`
}

type BackendAdapterRedactedProfileSmokeCoverageCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview(root string) (BackendAdapterRedactedProfileOwnerSmokeCoveragePreview, error) {
	records, err := NewSmokeBatchRecords(root)
	if err != nil {
		return BackendAdapterRedactedProfileOwnerSmokeCoveragePreview{}, err
	}
	return NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords(root, records)
}

func NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords(root string, records []SmokeBatchRecord) (BackendAdapterRedactedProfileOwnerSmokeCoveragePreview, error) {
	version, err := readBackendAdapterRedactedProfileOwnerSmokeCoverageVersion(root)
	if err != nil {
		return BackendAdapterRedactedProfileOwnerSmokeCoveragePreview{}, err
	}
	preview := newBackendAdapterRedactedProfileOwnerSmokeCoverageBase(version, records)
	for _, record := range records {
		switch record.RecordType {
		case "read-dispatch":
			preview.SmokeReadDispatchRecordCount++
		case "write-denial":
			preview.SmokeWriteDenialRecordCount++
		}
		if record.Method != "GetBackendAdapterProfileAudit" || record.RecordType != "read-dispatch" {
			continue
		}
		if err := applyBackendAdapterRedactedProfileOwnerSmokeCoverageRecord(&preview, record); err != nil {
			return BackendAdapterRedactedProfileOwnerSmokeCoveragePreview{}, err
		}
	}
	checks := backendAdapterRedactedProfileOwnerSmokeCoverageChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = backendAdapterRedactedProfileOwnerSmokeCoverageCheckIDs(checks)
	preview.CheckCount = len(checks)
	preview.PassedCheckCount = countBackendAdapterRedactedProfileOwnerSmokeCoveragePasses(checks)
	preview.AllChecksPassed = preview.PassedCheckCount == preview.CheckCount
	preview.SmokeCoverageReady = preview.AllChecksPassed
	if err := validateNoBackendTerms(preview, "backend adapter redacted profile owner smoke coverage preview"); err != nil {
		return BackendAdapterRedactedProfileOwnerSmokeCoveragePreview{}, err
	}
	return preview, nil
}

func newBackendAdapterRedactedProfileOwnerSmokeCoverageBase(version string, records []SmokeBatchRecord) BackendAdapterRedactedProfileOwnerSmokeCoveragePreview {
	return BackendAdapterRedactedProfileOwnerSmokeCoveragePreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.backend_adapter_redacted_profile_owner_smoke_coverage.v1",
		RequestType:                 "backend-adapter-redacted-profile-owner-smoke-coverage-preview",
		CoverageType:                "redacted-adapter-profile-owner-smoke-coverage",
		Source:                      "owner-smoke-batch+runtime-owner-service-call+backend-adapter-redacted-profile-audit-owner-route",
		RuntimeMethod:               "GetBackendAdapterProfileAudit",
		ReadMethod:                  "GetBackendAdapterRedactedProfileOwnerSmokeCoveragePreview",
		OwnerRouteRequestType:       "backend-adapter-redacted-profile-audit-preview",
		OwnerRouteCommand:           "backend-adapter-redacted-profile-audit-preview",
		SmokeBatchRecordCount:       len(records),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ReadOnlyCoverage:            true,
		ProductionDBusExposureReady: false,
		CallerStateRootRequired:     false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		StateRootWritesEnabled:      false,
		RuntimeWritesEnabled:        false,
		AdapterInvocationEnabled:    false,
		BackendInstallEnabled:       false,
		BackendDownloadEnabled:      false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		CommandMaterialized:         false,
		ExecutablePathResolved:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat redacted adapter profile smoke coverage as production ownership approval",
			"enable Runtime writes, adapter invocation, installation, download, command materialization, or launch from coverage evidence",
			"expose internal adapter identifiers, profile paths, raw commands, raw executables, or backend details through smoke coverage",
			"claim session or production D-Bus ownership from coverage review",
			"mutate host root, use network, or require privileged containers during coverage review",
		},
		NextRequirements: []string{
			"Keep the redacted adapter profile route covered by restricted owner smoke evidence.",
			"Keep the full adapter contract fixture-local until production trust gates exist.",
			"Require a separate production D-Bus gate before exposing this route outside owner-local review.",
			"Continue proving all write, adapter invocation, launch, raw-detail, and host-mutation gates remain closed.",
		},
		DesktopSafeSummary: "Restricted owner smoke coverage proves the redacted adapter profile owner-local read route is exercised through Service.Call while the full adapter contract stays fixture-local and all write, launch, raw-detail, and host mutation gates remain disabled.",
	}
}

func applyBackendAdapterRedactedProfileOwnerSmokeCoverageRecord(preview *BackendAdapterRedactedProfileOwnerSmokeCoveragePreview, record SmokeBatchRecord) error {
	var call ServiceCall
	if err := json.Unmarshal(record.Payload, &call); err != nil {
		return fmt.Errorf("decode redacted adapter profile service call: %w", err)
	}
	var dispatch ReadDispatch
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		return fmt.Errorf("decode redacted adapter profile dispatch: %w", err)
	}
	var route appidentity.BackendAdapterRedactedProfileAuditPreview
	if err := json.Unmarshal(dispatch.Payload, &route); err != nil {
		return fmt.Errorf("decode redacted adapter profile owner route payload: %w", err)
	}
	preview.CoverageRecordFound = true
	preview.CoverageRecordSequence = record.Sequence
	preview.ServiceCallReadDispatch = call.ReadOnlyDispatch
	preview.ServiceCallDispatchReady = call.DispatchReady
	preview.DispatchReadOnly = dispatch.ReadOnlyDispatch
	preview.DispatchRouteReady = dispatch.RouteReady
	preview.DispatchRouteSource = dispatch.RouteSource
	preview.DispatchGoCommand = dispatch.GoCommand
	preview.AuditType = route.AuditType
	preview.RouteDecision = route.RouteDecision
	preview.ProfileCount = len(route.Profiles)
	preview.ProfileIDs = append([]string(nil), route.ProfileIDs...)
	preview.ContractProfileCount = route.ContractProfileCount
	preview.FullContractConsumed = route.FullContractConsumed
	preview.KDEFacingProjectionConsumed = route.KDEFacingProjectionConsumed
	preview.InternalAdapterIDsRedacted = route.InternalAdapterIDsRedacted
	preview.InternalProfilePathsRedacted = route.InternalProfilePathsRedacted
	preview.OwnerLocalRouteCandidateReady = route.OwnerLocalRouteCandidateReady
	preview.FullContractFixtureLocal = route.FullContractFixtureLocal
	preview.ProductionDBusExposureReady = route.ProductionDBusExposureReady
	preview.CallerStateRootRequired = route.CallerStateRootRequired
	preview.SystemServiceStarted = route.SystemServiceStarted
	preview.SessionBusClaimed = route.SessionBusClaimed
	preview.ProductionBusClaimed = route.ProductionBusClaimed
	preview.WriteMethodsEnabled = route.WriteMethodsEnabled
	preview.StateRootWritesEnabled = false
	preview.RuntimeWritesEnabled = route.RuntimeWritesEnabled
	preview.AdapterInvocationEnabled = route.AdapterInvocationEnabled
	preview.BackendInstallEnabled = route.BackendInstallEnabled
	preview.BackendDownloadEnabled = route.BackendDownloadEnabled
	preview.BackendLaunchEnabled = route.BackendLaunchEnabled
	preview.BackendProcessStarted = route.BackendProcessStarted
	preview.CommandMaterialized = route.CommandMaterialized
	preview.ExecutablePathResolved = route.ExecutablePathResolved
	preview.NetworkRequired = route.NetworkRequired
	preview.HostRootModified = route.HostRootModified
	preview.PrivilegedContainerRequired = route.PrivilegedContainerRequired
	preview.StateRootPathExposed = route.StateRootPathExposed
	preview.RawCommandExposed = route.RawCommandExposed
	preview.RawExecutableExposed = route.RawExecutableExposed
	preview.BackendDetailsExposed = route.BackendDetailsExposed
	return nil
}

func backendAdapterRedactedProfileOwnerSmokeCoverageChecks(preview BackendAdapterRedactedProfileOwnerSmokeCoveragePreview) []BackendAdapterRedactedProfileSmokeCoverageCheck {
	return []BackendAdapterRedactedProfileSmokeCoverageCheck{
		backendAdapterRedactedProfileSmokeCoverageCheck("smoke-record-present", preview.CoverageRecordFound && preview.CoverageRecordSequence > 0, "Restricted owner smoke batch includes the redacted adapter profile owner-local read route."),
		backendAdapterRedactedProfileSmokeCoverageCheck("service-call-read-dispatch", preview.ServiceCallReadDispatch && preview.ServiceCallDispatchReady && preview.DispatchReadOnly && preview.DispatchRouteReady, "Service.Call and owner dispatch both mark the coverage record as ready read-only evidence."),
		backendAdapterRedactedProfileSmokeCoverageCheck("owner-route-payload-present", preview.DispatchRouteSource == "go-owner-local-preview" && preview.DispatchGoCommand == preview.OwnerRouteCommand && preview.OwnerRouteRequestType == "backend-adapter-redacted-profile-audit-preview", "Coverage unwraps the owner-local route payload rather than relying on a method-name-only smoke row."),
		backendAdapterRedactedProfileSmokeCoverageCheck("redacted-profiles-preserved", preview.AuditType == "owner-local-redacted-adapter-profile-audit" && preview.RouteDecision == "redacted-profile-route-ready" && preview.ProfileCount == 3 && len(preview.ProfileIDs) == 3 && preview.InternalAdapterIDsRedacted && preview.InternalProfilePathsRedacted, "Coverage preserves user-safe redacted profile output without internal identifiers or paths."),
		backendAdapterRedactedProfileSmokeCoverageCheck("full-contract-fixture-local", preview.FullContractConsumed && preview.KDEFacingProjectionConsumed && preview.ContractProfileCount == preview.ProfileCount && preview.FullContractFixtureLocal, "Coverage keeps the full adapter contract fixture-local while consuming only the KDE-facing projection."),
		backendAdapterRedactedProfileSmokeCoverageCheck("production-dbus-blocked", preview.OwnerLocalRouteCandidateReady && !preview.ProductionDBusExposureReady && !preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed, "Owner-local coverage remains blocked from production D-Bus exposure."),
		backendAdapterRedactedProfileSmokeCoverageCheck("caller-paths-hidden", !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed, "Smoke coverage does not require caller state roots or expose raw commands, executables, or backend details."),
		backendAdapterRedactedProfileSmokeCoverageCheck("unsafe-gates-closed", !preview.WriteMethodsEnabled && !preview.StateRootWritesEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendInstallEnabled && !preview.BackendDownloadEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.CommandMaterialized && !preview.ExecutablePathResolved && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired, "Coverage keeps writes, state-root writes, adapter invocation, install, download, command materialization, launch, network, privilege, and host mutation disabled."),
	}
}

func backendAdapterRedactedProfileSmokeCoverageCheck(id string, passed bool, summary string) BackendAdapterRedactedProfileSmokeCoverageCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return BackendAdapterRedactedProfileSmokeCoverageCheck{ID: id, Status: status, Summary: summary}
}

func backendAdapterRedactedProfileOwnerSmokeCoverageCheckIDs(checks []BackendAdapterRedactedProfileSmokeCoverageCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countBackendAdapterRedactedProfileOwnerSmokeCoveragePasses(checks []BackendAdapterRedactedProfileSmokeCoverageCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}

func readBackendAdapterRedactedProfileOwnerSmokeCoverageVersion(root string) (string, error) {
	content, err := os.ReadFile(filepath.Join(root, "VERSION"))
	if err != nil {
		return "", fmt.Errorf("read VERSION: %w", err)
	}
	version := strings.TrimSpace(string(content))
	if version == "" {
		return "", fmt.Errorf("read VERSION: empty version")
	}
	return version, nil
}
