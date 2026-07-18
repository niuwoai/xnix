package owner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview struct {
	Version                                 string                                        `json:"version"`
	SchemaVersion                           string                                        `json:"schema_version"`
	RequestType                             string                                        `json:"request_type"`
	CoverageType                            string                                        `json:"coverage_type"`
	Source                                  string                                        `json:"source"`
	RuntimeMethod                           string                                        `json:"runtime_method"`
	ReadMethod                              string                                        `json:"read_method"`
	OwnerRouteRequestType                   string                                        `json:"owner_route_request_type"`
	OwnerRouteCommand                       string                                        `json:"owner_route_command"`
	SmokeBatchRecordCount                   int                                           `json:"smoke_batch_record_count"`
	SmokeReadDispatchRecordCount            int                                           `json:"smoke_read_dispatch_record_count"`
	SmokeWriteDenialRecordCount             int                                           `json:"smoke_write_denial_record_count"`
	CoverageRecordFound                     bool                                          `json:"coverage_record_found"`
	CoverageRecordSequence                  int                                           `json:"coverage_record_sequence"`
	ServiceCallReadDispatch                 bool                                          `json:"service_call_read_dispatch"`
	ServiceCallDispatchReady                bool                                          `json:"service_call_dispatch_ready"`
	DispatchReadOnly                        bool                                          `json:"dispatch_read_only"`
	DispatchRouteReady                      bool                                          `json:"dispatch_route_ready"`
	DispatchRouteSource                     string                                        `json:"dispatch_route_source"`
	DispatchGoCommand                       string                                        `json:"dispatch_go_command"`
	OpaqueMaterializationReceiptID          string                                        `json:"opaque_materialization_receipt_id"`
	SupportedOpaqueMaterializationReceiptID string                                        `json:"supported_opaque_materialization_receipt_id"`
	ReceiptLookupState                      string                                        `json:"receipt_lookup_state"`
	ReceiptAvailable                        bool                                          `json:"receipt_available"`
	ReceiptConsumed                         bool                                          `json:"receipt_consumed"`
	MissingReceiptSafe                      bool                                          `json:"missing_receipt_safe"`
	FanOutResultState                       string                                        `json:"fan_out_result_state"`
	OwnerManagedOpaqueReceiptLookupReady    bool                                          `json:"owner_managed_opaque_receipt_lookup_ready"`
	OwnerLocalRouteCandidateReady           bool                                          `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady             bool                                          `json:"production_dbus_exposure_ready"`
	SurfaceCount                            int                                           `json:"surface_count"`
	CheckCount                              int                                           `json:"check_count"`
	PassedCheckCount                        int                                           `json:"passed_check_count"`
	AllChecksPassed                         bool                                          `json:"all_checks_passed"`
	Checks                                  []KDETestLaunchMaterializationOwnerSmokeCheck `json:"checks"`
	CheckIDs                                []string                                      `json:"check_ids"`
	RuntimeOwned                            bool                                          `json:"runtime_owned"`
	GoRuntimeBacked                         bool                                          `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                                          `json:"kde_policy_owner"`
	ReadOnlyCoverage                        bool                                          `json:"read_only_coverage"`
	SmokeCoverageReady                      bool                                          `json:"smoke_coverage_ready"`
	StateRootPathExposed                    bool                                          `json:"state_root_path_exposed"`
	StateRootWritesEnabled                  bool                                          `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled                    bool                                          `json:"runtime_writes_enabled"`
	FanOutWritesEnabled                     bool                                          `json:"fan_out_writes_enabled"`
	SystemServiceStarted                    bool                                          `json:"system_service_started"`
	SessionBusClaimed                       bool                                          `json:"session_bus_claimed"`
	ProductionBusClaimed                    bool                                          `json:"production_bus_claimed"`
	WriteMethodsEnabled                     bool                                          `json:"write_methods_enabled"`
	BackendLaunchEnabled                    bool                                          `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                          `json:"backend_process_started"`
	TaskManagerEntryActive                  bool                                          `json:"task_manager_entry_active"`
	LiveTrayBridgeEnabled                   bool                                          `json:"live_tray_bridge_enabled"`
	NotificationSent                        bool                                          `json:"notification_sent"`
	CompatibilityCenterActionsEnabled       bool                                          `json:"compatibility_center_actions_enabled"`
	NetworkRequired                         bool                                          `json:"network_required"`
	HostRootModified                        bool                                          `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                          `json:"privileged_container_required"`
	RawCommandExposed                       bool                                          `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                          `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                          `json:"backend_details_exposed"`
	BlockedActions                          []string                                      `json:"blocked_actions"`
	NextRequirements                        []string                                      `json:"next_requirements"`
	DesktopSafeSummary                      string                                        `json:"desktop_safe_summary"`
}

type KDETestLaunchMaterializationOwnerSmokeCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview(root string) (KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview, error) {
	records, err := NewSmokeBatchRecords(root)
	if err != nil {
		return KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview{}, err
	}
	return NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords(root, records)
}

func NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords(root string, records []SmokeBatchRecord) (KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview, error) {
	version, err := readKDETestLaunchMaterializationFanOutOwnerSmokeCoverageVersion(root)
	if err != nil {
		return KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview{}, err
	}
	preview := newKDETestLaunchMaterializationFanOutOwnerSmokeCoverageBase(version, records)
	for _, record := range records {
		switch record.RecordType {
		case "read-dispatch":
			preview.SmokeReadDispatchRecordCount++
		case "write-denial":
			preview.SmokeWriteDenialRecordCount++
		}
		if record.Method != "GetKDETestLaunchMaterializationFanOut" || record.RecordType != "read-dispatch" {
			continue
		}
		if err := applyKDETestLaunchMaterializationFanOutOwnerSmokeCoverageRecord(&preview, record); err != nil {
			return KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview{}, err
		}
	}
	checks := kdeTestLaunchMaterializationFanOutOwnerSmokeCoverageChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = kdeTestLaunchMaterializationFanOutOwnerSmokeCoverageCheckIDs(checks)
	preview.CheckCount = len(checks)
	preview.PassedCheckCount = countKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePasses(checks)
	preview.AllChecksPassed = preview.PassedCheckCount == preview.CheckCount
	preview.SmokeCoverageReady = preview.AllChecksPassed
	if err := validateNoBackendTerms(preview, "KDE test launch materialization fan-out owner smoke coverage preview"); err != nil {
		return KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview{}, err
	}
	return preview, nil
}

func newKDETestLaunchMaterializationFanOutOwnerSmokeCoverageBase(version string, records []SmokeBatchRecord) KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview {
	return KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview{
		Version:                           version,
		SchemaVersion:                     "xnix.runtime.kde_test_launch_materialization_fanout_owner_smoke_coverage.v1",
		RequestType:                       "kde-test-launch-materialization-fanout-owner-smoke-coverage-preview",
		CoverageType:                      "materialization-fanout-owner-smoke-coverage",
		Source:                            "owner-smoke-batch+runtime-owner-service-call+kde-test-launch-materialization-fanout-owner-route",
		RuntimeMethod:                     "GetKDETestLaunchMaterializationFanOut",
		ReadMethod:                        "GetKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview",
		OwnerRouteRequestType:             "kde-test-launch-materialization-fanout-owner-route-preview",
		OwnerRouteCommand:                 "kde-test-launch-materialization-fanout-owner-route-preview",
		SmokeBatchRecordCount:             len(records),
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		ReadOnlyCoverage:                  true,
		ProductionDBusExposureReady:       false,
		StateRootPathExposed:              false,
		StateRootWritesEnabled:            false,
		RuntimeWritesEnabled:              false,
		FanOutWritesEnabled:               false,
		SystemServiceStarted:              false,
		SessionBusClaimed:                 false,
		ProductionBusClaimed:              false,
		WriteMethodsEnabled:               false,
		BackendLaunchEnabled:              false,
		BackendProcessStarted:             false,
		TaskManagerEntryActive:            false,
		LiveTrayBridgeEnabled:             false,
		NotificationSent:                  false,
		CompatibilityCenterActionsEnabled: false,
		NetworkRequired:                   false,
		HostRootModified:                  false,
		PrivilegedContainerRequired:       false,
		RawCommandExposed:                 false,
		RawExecutableExposed:              false,
		BackendDetailsExposed:             false,
		BlockedActions: []string{
			"treat smoke batch coverage as permission to claim production D-Bus ownership",
			"enable Runtime writes, command materialization, executable resolution, or process launch from coverage evidence",
			"activate task-manager entries, tray bridges, notifications, or Compatibility Center actions from coverage evidence",
			"accept caller state-root paths or expose raw executable details through smoke coverage",
			"mutate host root, use network, or require privileged containers during coverage review",
		},
		NextRequirements: []string{
			"Keep the owner-local materialization fan-out route covered by restricted owner smoke evidence.",
			"Attach digest-verified materialization receipts before treating KDE fan-out as satisfied.",
			"Require a separate production D-Bus gate before exposing this route outside owner-local review.",
			"Continue proving all unsafe launch and desktop side-effect gates remain closed.",
		},
		DesktopSafeSummary: "Restricted owner smoke coverage proves the materialization fan-out owner-local read route is exercised through Service.Call while production ownership, writes, launch, desktop side effects, and host mutation remain disabled.",
	}
}

func applyKDETestLaunchMaterializationFanOutOwnerSmokeCoverageRecord(preview *KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview, record SmokeBatchRecord) error {
	var call ServiceCall
	if err := json.Unmarshal(record.Payload, &call); err != nil {
		return fmt.Errorf("decode materialization fan-out smoke service call: %w", err)
	}
	var dispatch ReadDispatch
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		return fmt.Errorf("decode materialization fan-out smoke dispatch: %w", err)
	}
	var route appidentity.KDETestLaunchMaterializationFanOutOwnerRoutePreview
	if err := json.Unmarshal(dispatch.Payload, &route); err != nil {
		return fmt.Errorf("decode materialization fan-out owner route payload: %w", err)
	}
	preview.CoverageRecordFound = true
	preview.CoverageRecordSequence = record.Sequence
	preview.ServiceCallReadDispatch = call.ReadOnlyDispatch
	preview.ServiceCallDispatchReady = call.DispatchReady
	preview.DispatchReadOnly = dispatch.ReadOnlyDispatch
	preview.DispatchRouteReady = dispatch.RouteReady
	preview.DispatchRouteSource = dispatch.RouteSource
	preview.DispatchGoCommand = dispatch.GoCommand
	preview.OpaqueMaterializationReceiptID = route.OpaqueMaterializationReceiptID
	preview.SupportedOpaqueMaterializationReceiptID = route.SupportedOpaqueMaterializationReceiptID
	preview.ReceiptLookupState = route.ReceiptLookupState
	preview.ReceiptAvailable = route.ReceiptAvailable
	preview.ReceiptConsumed = route.ReceiptConsumed
	preview.MissingReceiptSafe = route.MissingReceiptSafe
	preview.FanOutResultState = route.FanOutResultState
	preview.OwnerManagedOpaqueReceiptLookupReady = route.OwnerManagedOpaqueReceiptLookupReady
	preview.OwnerLocalRouteCandidateReady = route.OwnerLocalRouteCandidateReady
	preview.ProductionDBusExposureReady = route.ProductionDBusExposureReady
	preview.SurfaceCount = route.SurfaceCount
	preview.TaskManagerEntryActive = route.TaskManagerEntryActive
	preview.LiveTrayBridgeEnabled = route.LiveTrayBridgeEnabled
	preview.NotificationSent = route.NotificationSent
	preview.CompatibilityCenterActionsEnabled = route.CompatibilityCenterActionsEnabled
	preview.BackendLaunchEnabled = route.BackendLaunchEnabled
	preview.BackendProcessStarted = route.BackendProcessStarted
	preview.HostRootModified = route.HostRootModified
	preview.BackendDetailsExposed = route.BackendDetailsExposed
	return nil
}

func kdeTestLaunchMaterializationFanOutOwnerSmokeCoverageChecks(preview KDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview) []KDETestLaunchMaterializationOwnerSmokeCheck {
	return []KDETestLaunchMaterializationOwnerSmokeCheck{
		kdeTestLaunchMaterializationOwnerSmokeCheck("smoke-record-present", preview.CoverageRecordFound && preview.CoverageRecordSequence > 0, "Restricted owner smoke batch includes the materialization fan-out owner-local read route."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("service-call-read-dispatch", preview.ServiceCallReadDispatch && preview.ServiceCallDispatchReady && preview.DispatchReadOnly && preview.DispatchRouteReady, "Service.Call and owner dispatch both mark the coverage record as ready read-only evidence."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("owner-route-payload-present", preview.DispatchRouteSource == "go-owner-local-preview" && preview.DispatchGoCommand == preview.OwnerRouteCommand && preview.OwnerRouteRequestType == "kde-test-launch-materialization-fanout-owner-route-preview", "Coverage unwraps the owner-local route payload rather than relying on a method-name-only smoke row."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("opaque-receipt-preserved", preview.OpaqueMaterializationReceiptID == appidentity.KDETestLaunchMaterializationOpaqueReceiptID && preview.SupportedOpaqueMaterializationReceiptID == appidentity.KDETestLaunchMaterializationOpaqueReceiptID && preview.OwnerManagedOpaqueReceiptLookupReady, "Owner smoke coverage preserves the owner-managed opaque materialization receipt id."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("missing-receipt-fail-closed", !preview.ReceiptAvailable && !preview.ReceiptConsumed && preview.MissingReceiptSafe && preview.ReceiptLookupState == "missing-receipt" && preview.FanOutResultState == "missing-receipt-fail-closed", "Missing materialization receipts stay fail-closed under smoke coverage."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("desktop-side-effects-disabled", !preview.TaskManagerEntryActive && !preview.LiveTrayBridgeEnabled && !preview.NotificationSent && !preview.CompatibilityCenterActionsEnabled, "Smoke coverage does not activate task-manager entries, tray bridges, notifications, or Compatibility Center actions."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("production-dbus-blocked", preview.OwnerLocalRouteCandidateReady && !preview.ProductionDBusExposureReady && !preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed, "Owner-local coverage remains blocked from production D-Bus exposure."),
		kdeTestLaunchMaterializationOwnerSmokeCheck("unsafe-gates-closed", !preview.StateRootPathExposed && !preview.StateRootWritesEnabled && !preview.RuntimeWritesEnabled && !preview.FanOutWritesEnabled && !preview.WriteMethodsEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed, "Coverage keeps writes, launch, network, privilege, unsafe data, and host mutation disabled."),
	}
}

func kdeTestLaunchMaterializationOwnerSmokeCheck(id string, passed bool, summary string) KDETestLaunchMaterializationOwnerSmokeCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return KDETestLaunchMaterializationOwnerSmokeCheck{ID: id, Status: status, Summary: summary}
}

func kdeTestLaunchMaterializationFanOutOwnerSmokeCoverageCheckIDs(checks []KDETestLaunchMaterializationOwnerSmokeCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePasses(checks []KDETestLaunchMaterializationOwnerSmokeCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}

func readKDETestLaunchMaterializationFanOutOwnerSmokeCoverageVersion(root string) (string, error) {
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
