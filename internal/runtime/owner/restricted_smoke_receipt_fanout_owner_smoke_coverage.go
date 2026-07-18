package owner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview struct {
	Version                       string                                    `json:"version"`
	SchemaVersion                 string                                    `json:"schema_version"`
	RequestType                   string                                    `json:"request_type"`
	CoverageType                  string                                    `json:"coverage_type"`
	Source                        string                                    `json:"source"`
	RuntimeMethod                 string                                    `json:"runtime_method"`
	ReadMethod                    string                                    `json:"read_method"`
	OwnerRouteRequestType         string                                    `json:"owner_route_request_type"`
	OwnerRouteCommand             string                                    `json:"owner_route_command"`
	SmokeBatchRecordCount         int                                       `json:"smoke_batch_record_count"`
	SmokeReadDispatchRecordCount  int                                       `json:"smoke_read_dispatch_record_count"`
	SmokeWriteDenialRecordCount   int                                       `json:"smoke_write_denial_record_count"`
	CoverageRecordFound           bool                                      `json:"coverage_record_found"`
	CoverageRecordSequence        int                                       `json:"coverage_record_sequence"`
	ServiceCallReadDispatch       bool                                      `json:"service_call_read_dispatch"`
	ServiceCallDispatchReady      bool                                      `json:"service_call_dispatch_ready"`
	DispatchReadOnly              bool                                      `json:"dispatch_read_only"`
	DispatchRouteReady            bool                                      `json:"dispatch_route_ready"`
	DispatchRouteSource           string                                    `json:"dispatch_route_source"`
	DispatchGoCommand             string                                    `json:"dispatch_go_command"`
	OpaqueReceiptID               string                                    `json:"opaque_receipt_id"`
	SupportedOpaqueReceiptID      string                                    `json:"supported_opaque_receipt_id"`
	ReceiptLookupState            string                                    `json:"receipt_lookup_state"`
	ReceiptAvailable              bool                                      `json:"receipt_available"`
	ReceiptConsumed               bool                                      `json:"receipt_consumed"`
	MissingReceiptSafe            bool                                      `json:"missing_receipt_safe"`
	FanOutResultState             string                                    `json:"fan_out_result_state"`
	OwnerManagedLookup            bool                                      `json:"owner_managed_lookup"`
	OwnerLocalRouteCandidateReady bool                                      `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady   bool                                      `json:"production_dbus_exposure_ready"`
	SurfaceCount                  int                                       `json:"surface_count"`
	CheckCount                    int                                       `json:"check_count"`
	PassedCheckCount              int                                       `json:"passed_check_count"`
	AllChecksPassed               bool                                      `json:"all_checks_passed"`
	Checks                        []RestrictedOwnerSmokeFanOutCoverageCheck `json:"checks"`
	CheckIDs                      []string                                  `json:"check_ids"`
	RuntimeOwned                  bool                                      `json:"runtime_owned"`
	GoRuntimeBacked               bool                                      `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                                      `json:"kde_policy_owner"`
	ReadOnlyCoverage              bool                                      `json:"read_only_coverage"`
	SmokeCoverageReady            bool                                      `json:"smoke_coverage_ready"`
	ReadinessSurfacesSatisfied    bool                                      `json:"readiness_surfaces_satisfied"`
	SupportSurfacesSatisfied      bool                                      `json:"support_surfaces_satisfied"`
	StateRootPathExposed          bool                                      `json:"state_root_path_exposed"`
	StateRootWritesEnabled        bool                                      `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled          bool                                      `json:"runtime_writes_enabled"`
	FanOutWritesEnabled           bool                                      `json:"fan_out_writes_enabled"`
	SystemServiceStarted          bool                                      `json:"system_service_started"`
	SessionBusClaimed             bool                                      `json:"session_bus_claimed"`
	ProductionBusClaimed          bool                                      `json:"production_bus_claimed"`
	WriteMethodsEnabled           bool                                      `json:"write_methods_enabled"`
	ProductionActivationReady     bool                                      `json:"production_activation_ready"`
	ProductionOwnerEnabled        bool                                      `json:"production_owner_enabled"`
	SupportBundleExported         bool                                      `json:"support_bundle_exported"`
	SupportCaseCreated            bool                                      `json:"support_case_created"`
	NotificationSent              bool                                      `json:"notification_sent"`
	BackendLaunchEnabled          bool                                      `json:"backend_launch_enabled"`
	BackendProcessStarted         bool                                      `json:"backend_process_started"`
	NetworkRequired               bool                                      `json:"network_required"`
	HostRootModified              bool                                      `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                                      `json:"privileged_container_required"`
	BackendDetailsExposed         bool                                      `json:"backend_details_exposed"`
	BlockedActions                []string                                  `json:"blocked_actions"`
	NextRequirements              []string                                  `json:"next_requirements"`
	DesktopSafeSummary            string                                    `json:"desktop_safe_summary"`
}

type RestrictedOwnerSmokeFanOutCoverageCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview(root string) (RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview, error) {
	records, err := NewSmokeBatchRecords(root)
	if err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview{}, err
	}
	return NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords(root, records)
}

func NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords(root string, records []SmokeBatchRecord) (RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview, error) {
	version, err := readRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageVersion(root)
	if err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview{}, err
	}
	preview := newRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageBase(version, records)
	for _, record := range records {
		switch record.RecordType {
		case "read-dispatch":
			preview.SmokeReadDispatchRecordCount++
		case "write-denial":
			preview.SmokeWriteDenialRecordCount++
		}
		if record.Method != "GetRestrictedOwnerSmokeReceiptFanOut" || record.RecordType != "read-dispatch" {
			continue
		}
		if err := applyRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageRecord(&preview, record); err != nil {
			return RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview{}, err
		}
	}
	checks := restrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = restrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageCheckIDs(checks)
	preview.CheckCount = len(checks)
	preview.PassedCheckCount = countRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePasses(checks)
	preview.AllChecksPassed = preview.PassedCheckCount == preview.CheckCount
	preview.SmokeCoverageReady = preview.AllChecksPassed
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out owner smoke coverage preview"); err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview{}, err
	}
	return preview, nil
}

func newRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageBase(version string, records []SmokeBatchRecord) RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview {
	return RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_smoke_coverage.v1",
		RequestType:                 "restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview",
		CoverageType:                "restricted-owner-smoke-fanout-owner-smoke-coverage",
		Source:                      "owner-smoke-batch+runtime-owner-service-call+restricted-owner-smoke-receipt-fanout-owner-route",
		RuntimeMethod:               "GetRestrictedOwnerSmokeReceiptFanOut",
		ReadMethod:                  "GetRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview",
		OwnerRouteRequestType:       "restricted-owner-smoke-receipt-fanout-owner-route-preview",
		OwnerRouteCommand:           "restricted-owner-smoke-receipt-fanout-owner-route-preview",
		SmokeBatchRecordCount:       len(records),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ReadOnlyCoverage:            true,
		ProductionDBusExposureReady: false,
		StateRootPathExposed:        false,
		StateRootWritesEnabled:      false,
		RuntimeWritesEnabled:        false,
		FanOutWritesEnabled:         false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		ProductionActivationReady:   false,
		ProductionOwnerEnabled:      false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
		NotificationSent:            false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"treat restricted owner smoke fan-out coverage as production ownership approval",
			"enable Runtime writes, support bundle export, support case creation, notifications, or launch from coverage evidence",
			"accept caller state-root paths or expose implementation details through smoke coverage",
			"claim session or production D-Bus ownership from coverage review",
			"mutate host root, use network, or require privileged containers during coverage review",
		},
		NextRequirements: []string{
			"Keep the restricted owner smoke fan-out route covered by restricted owner smoke evidence.",
			"Attach digest-verified restricted owner smoke receipts before treating readiness or support surfaces as satisfied.",
			"Require a separate production D-Bus gate before exposing this route outside owner-local review.",
			"Continue proving all write, support-side-effect, launch, and host-mutation gates remain closed.",
		},
		DesktopSafeSummary: "Restricted owner smoke coverage proves the restricted smoke receipt fan-out owner-local read route is exercised through Service.Call while production ownership, writes, support side effects, launch, and host mutation remain disabled.",
	}
}

func applyRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageRecord(preview *RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview, record SmokeBatchRecord) error {
	var call ServiceCall
	if err := json.Unmarshal(record.Payload, &call); err != nil {
		return fmt.Errorf("decode restricted owner smoke fan-out service call: %w", err)
	}
	var dispatch ReadDispatch
	if err := json.Unmarshal(call.Payload, &dispatch); err != nil {
		return fmt.Errorf("decode restricted owner smoke fan-out dispatch: %w", err)
	}
	var route RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview
	if err := json.Unmarshal(dispatch.Payload, &route); err != nil {
		return fmt.Errorf("decode restricted owner smoke fan-out owner route payload: %w", err)
	}
	preview.CoverageRecordFound = true
	preview.CoverageRecordSequence = record.Sequence
	preview.ServiceCallReadDispatch = call.ReadOnlyDispatch
	preview.ServiceCallDispatchReady = call.DispatchReady
	preview.DispatchReadOnly = dispatch.ReadOnlyDispatch
	preview.DispatchRouteReady = dispatch.RouteReady
	preview.DispatchRouteSource = dispatch.RouteSource
	preview.DispatchGoCommand = dispatch.GoCommand
	preview.OpaqueReceiptID = route.OpaqueReceiptID
	preview.SupportedOpaqueReceiptID = route.SupportedOpaqueReceiptID
	preview.ReceiptLookupState = route.ReceiptLookupState
	preview.ReceiptAvailable = route.ReceiptAvailable
	preview.ReceiptConsumed = route.ReceiptConsumed
	preview.MissingReceiptSafe = route.MissingReceiptSafe
	preview.FanOutResultState = route.FanOutResultState
	preview.OwnerManagedLookup = route.OwnerManagedLookup
	preview.OwnerLocalRouteCandidateReady = route.OwnerLocalRouteCandidateReady
	preview.ProductionDBusExposureReady = route.ProductionDBusExposureReady
	preview.SurfaceCount = route.SurfaceCount
	preview.ReadinessSurfacesSatisfied = route.ReadinessSurfacesSatisfied
	preview.SupportSurfacesSatisfied = route.SupportSurfacesSatisfied
	preview.ProductionActivationReady = route.ProductionActivationReady
	preview.ProductionOwnerEnabled = route.ProductionOwnerEnabled
	preview.SupportBundleExported = route.SupportBundleExported
	preview.SupportCaseCreated = route.SupportCaseCreated
	preview.NotificationSent = route.NotificationSent
	preview.BackendLaunchEnabled = route.BackendLaunchEnabled
	preview.BackendProcessStarted = route.BackendProcessStarted
	preview.HostRootModified = route.HostRootModified
	preview.BackendDetailsExposed = route.BackendDetailsExposed
	return nil
}

func restrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageChecks(preview RestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview) []RestrictedOwnerSmokeFanOutCoverageCheck {
	return []RestrictedOwnerSmokeFanOutCoverageCheck{
		restrictedOwnerSmokeFanOutCoverageCheck("smoke-record-present", preview.CoverageRecordFound && preview.CoverageRecordSequence > 0, "Restricted owner smoke batch includes the restricted smoke receipt fan-out owner-local read route."),
		restrictedOwnerSmokeFanOutCoverageCheck("service-call-read-dispatch", preview.ServiceCallReadDispatch && preview.ServiceCallDispatchReady && preview.DispatchReadOnly && preview.DispatchRouteReady, "Service.Call and owner dispatch both mark the coverage record as ready read-only evidence."),
		restrictedOwnerSmokeFanOutCoverageCheck("owner-route-payload-present", preview.DispatchRouteSource == "go-owner-local-preview" && preview.DispatchGoCommand == preview.OwnerRouteCommand && preview.OwnerRouteRequestType == "restricted-owner-smoke-receipt-fanout-owner-route-preview", "Coverage unwraps the owner-local route payload rather than relying on a method-name-only smoke row."),
		restrictedOwnerSmokeFanOutCoverageCheck("opaque-receipt-preserved", preview.OpaqueReceiptID == RestrictedOwnerSmokeOpaqueReceiptID && preview.SupportedOpaqueReceiptID == RestrictedOwnerSmokeOpaqueReceiptID && preview.OwnerManagedLookup, "Owner smoke coverage preserves the owner-managed opaque restricted smoke receipt id."),
		restrictedOwnerSmokeFanOutCoverageCheck("missing-receipt-fail-closed", !preview.ReceiptAvailable && !preview.ReceiptConsumed && preview.MissingReceiptSafe && preview.ReceiptLookupState == "missing-receipt" && preview.FanOutResultState == "missing-receipt-fail-closed" && !preview.ReadinessSurfacesSatisfied && !preview.SupportSurfacesSatisfied, "Missing restricted owner smoke receipts stay fail-closed under smoke coverage."),
		restrictedOwnerSmokeFanOutCoverageCheck("support-side-effects-disabled", !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.NotificationSent, "Smoke coverage does not export support bundles, create support cases, or send notifications."),
		restrictedOwnerSmokeFanOutCoverageCheck("production-dbus-blocked", preview.OwnerLocalRouteCandidateReady && !preview.ProductionDBusExposureReady && !preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.ProductionActivationReady && !preview.ProductionOwnerEnabled, "Owner-local coverage remains blocked from production activation and D-Bus exposure."),
		restrictedOwnerSmokeFanOutCoverageCheck("unsafe-gates-closed", !preview.StateRootPathExposed && !preview.StateRootWritesEnabled && !preview.RuntimeWritesEnabled && !preview.FanOutWritesEnabled && !preview.WriteMethodsEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.BackendDetailsExposed, "Coverage keeps writes, support side effects, launch, network, privilege, implementation details, and host mutation disabled."),
	}
}

func restrictedOwnerSmokeFanOutCoverageCheck(id string, passed bool, summary string) RestrictedOwnerSmokeFanOutCoverageCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return RestrictedOwnerSmokeFanOutCoverageCheck{ID: id, Status: status, Summary: summary}
}

func restrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageCheckIDs(checks []RestrictedOwnerSmokeFanOutCoverageCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePasses(checks []RestrictedOwnerSmokeFanOutCoverageCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}

func readRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoverageVersion(root string) (string, error) {
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
