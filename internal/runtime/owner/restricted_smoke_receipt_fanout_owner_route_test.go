package owner

import "testing"

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreviewUsesOpaqueLookup(t *testing.T) {
	preview, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(projectRoot(t), RestrictedOwnerSmokeOpaqueReceiptID)
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1" ||
		preview.RequestType != "restricted-owner-smoke-receipt-fanout-owner-route-preview" ||
		preview.RouteType != "owner-local-restricted-smoke-receipt-fanout" ||
		preview.RuntimeMethod != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		preview.ReadMethod != "GetRestrictedOwnerSmokeReceiptFanOutPreview" ||
		preview.OpaqueReceiptID != RestrictedOwnerSmokeOpaqueReceiptID ||
		preview.SupportedOpaqueReceiptID != RestrictedOwnerSmokeOpaqueReceiptID {
		t.Fatalf("unexpected owner-route fan-out schema: %#v", preview)
	}
	if !preview.OwnerManagedLookup ||
		preview.CallerStateRootRequired ||
		!preview.OpaqueReceiptIDSupported ||
		!preview.ReadOnlyFanOut ||
		preview.FanOutResultState != "missing-receipt-fail-closed" ||
		!preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.SurfaceCount != 5 ||
		preview.ReceiptLookupState != "missing-receipt" ||
		preview.ReceiptAvailable ||
		preview.ReceiptConsumed ||
		!preview.MissingReceiptSafe {
		t.Fatalf("unexpected owner-route fan-out decision: %#v", preview)
	}
	if preview.ReadinessSurfacesSatisfied || preview.SupportSurfacesSatisfied ||
		preview.StateRootPathExposed ||
		preview.StateRootWritesEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.FanOutWritesEnabled ||
		preview.ProductionActivationReady ||
		preview.ProductionOwnerEnabled ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe owner-route fan-out flags: %#v", preview)
	}
	for _, surface := range preview.Surfaces {
		if surface.EvidenceState != "missing-receipt" ||
			!surface.ReadOnly ||
			!surface.UserVisible ||
			surface.ConsumesReceipt ||
			surface.MutatesRuntime ||
			surface.StartsService ||
			surface.ClaimsSessionBus ||
			surface.ClaimsProductionBus ||
			surface.EnablesWriteMethods ||
			surface.StartsBackend ||
			surface.ExportsSupportBundle ||
			surface.CreatesSupportCase ||
			surface.SendsNotification ||
			surface.ExposesStateRootPath ||
			surface.ExposesBackendDetails {
			t.Fatalf("unexpected owner-route fan-out surface: %#v", surface)
		}
	}
	if preview.CheckCount != 7 ||
		preview.PassedCheckCount != 7 ||
		!preview.AllChecksPassed ||
		!sameStrings(preview.CheckIDs, []string{"opaque-lookup-consumed", "caller-state-root-hidden", "missing-receipt-fails-closed", "surface-fanout-deferred", "support-side-effects-disabled", "owner-route-ready-production-dbus-blocked", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected owner-route fan-out checks: count=%d passed=%d ids=%#v", preview.CheckCount, preview.PassedCheckCount, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out owner route test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteRejectsUnknownOpaqueID(t *testing.T) {
	if _, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(projectRoot(t), "unknown-receipt"); err == nil {
		t.Fatal("owner-route fan-out accepted an unknown opaque receipt id")
	}
}
