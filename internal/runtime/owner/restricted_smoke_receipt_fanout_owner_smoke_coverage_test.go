package owner

import "testing"

func TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewCoversOwnerRoute(t *testing.T) {
	preview, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview(projectRoot(t))
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_smoke_coverage.v1" ||
		preview.RequestType != "restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview" ||
		preview.CoverageType != "restricted-owner-smoke-fanout-owner-smoke-coverage" ||
		preview.RuntimeMethod != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		preview.ReadMethod != "GetRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreview" ||
		preview.OwnerRouteRequestType != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		preview.OwnerRouteCommand != "restricted-owner-smoke-receipt-fanout-owner-route-preview" {
		t.Fatalf("unexpected restricted smoke fan-out coverage schema: %#v", preview)
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
		preview.DispatchGoCommand != "restricted-owner-smoke-receipt-fanout-owner-route-preview" {
		t.Fatalf("unexpected restricted smoke fan-out coverage route state: %#v", preview)
	}
	if preview.OpaqueReceiptID != "restricted-owner-smoke-receipt-id" ||
		preview.SupportedOpaqueReceiptID != "restricted-owner-smoke-receipt-id" ||
		preview.ReceiptLookupState != "missing-receipt" ||
		preview.ReceiptAvailable ||
		preview.ReceiptConsumed ||
		!preview.MissingReceiptSafe ||
		preview.FanOutResultState != "missing-receipt-fail-closed" ||
		!preview.OwnerManagedLookup ||
		!preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.SurfaceCount != 5 ||
		preview.ReadinessSurfacesSatisfied ||
		preview.SupportSurfacesSatisfied {
		t.Fatalf("unexpected restricted smoke fan-out coverage receipt state: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 8 ||
		!preview.AllChecksPassed ||
		!preview.SmokeCoverageReady ||
		!sameStrings(preview.CheckIDs, []string{"smoke-record-present", "service-call-read-dispatch", "owner-route-payload-present", "opaque-receipt-preserved", "missing-receipt-fail-closed", "support-side-effects-disabled", "production-dbus-blocked", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected restricted smoke fan-out coverage checks: counts=%d/%d ids=%#v", preview.PassedCheckCount, preview.CheckCount, preview.CheckIDs)
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
		preview.ProductionActivationReady ||
		preview.ProductionOwnerEnabled ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected restricted smoke fan-out coverage safety gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "restricted smoke fan-out owner smoke coverage test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFailsClosedWithoutRecord(t *testing.T) {
	records, err := NewSmokeBatchRecords(projectRoot(t))
	if err != nil {
		t.Fatalf("NewSmokeBatchRecords returned error: %v", err)
	}
	filtered := make([]SmokeBatchRecord, 0, len(records))
	for _, record := range records {
		if record.Method == "GetRestrictedOwnerSmokeReceiptFanOut" {
			continue
		}
		filtered = append(filtered, record)
	}

	preview, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords(projectRoot(t), filtered)
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutOwnerSmokeCoveragePreviewFromRecords returned error: %v", err)
	}
	if preview.CoverageRecordFound ||
		preview.CoverageRecordSequence != 0 ||
		preview.ServiceCallReadDispatch ||
		preview.DispatchRouteReady ||
		preview.SmokeCoverageReady ||
		preview.AllChecksPassed {
		t.Fatalf("missing restricted smoke fan-out coverage record must fail closed: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 2 ||
		preview.Checks[0].ID != "smoke-record-present" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[7].ID != "unsafe-gates-closed" ||
		preview.Checks[7].Status != "pass" {
		t.Fatalf("unexpected fail-closed restricted smoke fan-out coverage checks: counts=%d/%d checks=%#v", preview.PassedCheckCount, preview.CheckCount, preview.Checks)
	}
}
