package owner

import "testing"

func TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute(t *testing.T) {
	preview, err := NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview(projectRoot(t))
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.kde_test_launch_materialization_fanout_owner_smoke_coverage.v1" ||
		preview.RequestType != "kde-test-launch-materialization-fanout-owner-smoke-coverage-preview" ||
		preview.CoverageType != "materialization-fanout-owner-smoke-coverage" ||
		preview.RuntimeMethod != "GetKDETestLaunchMaterializationFanOut" ||
		preview.ReadMethod != "GetKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreview" ||
		preview.OwnerRouteRequestType != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		preview.OwnerRouteCommand != "kde-test-launch-materialization-fanout-owner-route-preview" {
		t.Fatalf("unexpected smoke coverage schema: %#v", preview)
	}
	if preview.SmokeBatchRecordCount != 75 ||
		preview.SmokeReadDispatchRecordCount != 71 ||
		preview.SmokeWriteDenialRecordCount != 4 ||
		!preview.CoverageRecordFound ||
		preview.CoverageRecordSequence == 0 ||
		!preview.ServiceCallReadDispatch ||
		!preview.ServiceCallDispatchReady ||
		!preview.DispatchReadOnly ||
		!preview.DispatchRouteReady ||
		preview.DispatchRouteSource != "go-owner-local-preview" ||
		preview.DispatchGoCommand != "kde-test-launch-materialization-fanout-owner-route-preview" {
		t.Fatalf("unexpected smoke coverage route state: %#v", preview)
	}
	if preview.OpaqueMaterializationReceiptID != "kde-test-launch-materialization-receipt-id" ||
		preview.SupportedOpaqueMaterializationReceiptID != "kde-test-launch-materialization-receipt-id" ||
		preview.ReceiptLookupState != "missing-receipt" ||
		preview.ReceiptAvailable ||
		preview.ReceiptConsumed ||
		!preview.MissingReceiptSafe ||
		preview.FanOutResultState != "missing-receipt-fail-closed" ||
		!preview.OwnerManagedOpaqueReceiptLookupReady ||
		!preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.SurfaceCount != 4 {
		t.Fatalf("unexpected smoke coverage materialization state: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 8 ||
		!preview.AllChecksPassed ||
		!preview.SmokeCoverageReady ||
		!sameStrings(preview.CheckIDs, []string{"smoke-record-present", "service-call-read-dispatch", "owner-route-payload-present", "opaque-receipt-preserved", "missing-receipt-fail-closed", "desktop-side-effects-disabled", "production-dbus-blocked", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected smoke coverage checks: counts=%d/%d ids=%#v", preview.PassedCheckCount, preview.CheckCount, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ReadOnlyCoverage ||
		preview.StateRootPathExposed ||
		preview.StateRootWritesEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.FanOutWritesEnabled ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.TaskManagerEntryActive ||
		preview.LiveTrayBridgeEnabled ||
		preview.NotificationSent ||
		preview.CompatibilityCenterActionsEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected smoke coverage safety gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "materialization fan-out owner smoke coverage test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord(t *testing.T) {
	records, err := NewSmokeBatchRecords(projectRoot(t))
	if err != nil {
		t.Fatalf("NewSmokeBatchRecords returned error: %v", err)
	}
	filtered := make([]SmokeBatchRecord, 0, len(records))
	for _, record := range records {
		if record.Method == "GetKDETestLaunchMaterializationFanOut" {
			continue
		}
		filtered = append(filtered, record)
	}

	preview, err := NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords(projectRoot(t), filtered)
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationFanOutOwnerSmokeCoveragePreviewFromRecords returned error: %v", err)
	}
	if preview.CoverageRecordFound ||
		preview.CoverageRecordSequence != 0 ||
		preview.ServiceCallReadDispatch ||
		preview.DispatchRouteReady ||
		preview.SmokeCoverageReady ||
		preview.AllChecksPassed {
		t.Fatalf("missing smoke coverage record must fail closed: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 2 ||
		preview.Checks[0].ID != "smoke-record-present" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[7].ID != "unsafe-gates-closed" ||
		preview.Checks[7].Status != "pass" {
		t.Fatalf("unexpected fail-closed smoke coverage checks: counts=%d/%d checks=%#v", preview.PassedCheckCount, preview.CheckCount, preview.Checks)
	}
}
