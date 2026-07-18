package appidentity

import "testing"

func TestKDETestLaunchMaterializationFanOutOwnerRoutePreviewUsesOpaqueLookup(t *testing.T) {
	preview, err := NewKDETestLaunchMaterializationFanOutOwnerRoutePreview(projectRootForRuntimeServiceBindingTest(t), KDETestLaunchMaterializationOpaqueReceiptID)
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationFanOutOwnerRoutePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.kde_test_launch_materialization_fanout_owner_route.v1" ||
		preview.RequestType != "kde-test-launch-materialization-fanout-owner-route-preview" ||
		preview.RouteType != "owner-local-kde-test-launch-materialization-fanout" ||
		preview.RuntimeMethod != "GetKDETestLaunchMaterializationFanOut" ||
		preview.ReadMethod != "GetKDETestLaunchMaterializationFanOutPreview" ||
		preview.OpaqueMaterializationReceiptID != "kde-test-launch-materialization-receipt-id" ||
		preview.SupportedOpaqueMaterializationReceiptID != KDETestLaunchMaterializationOpaqueReceiptID ||
		preview.MaterializationPlanID != "kde-test-launch-materialization-plan" {
		t.Fatalf("unexpected materialization fan-out owner-route schema: %#v", preview)
	}
	if !preview.OwnerManagedOpaqueReceiptLookupReady ||
		!preview.OpaqueMaterializationReceiptIDSupported ||
		preview.RequiresCallerRegistryPath ||
		preview.RequiresCallerApplicationID ||
		preview.RequiresCallerStateRoot ||
		!preview.ReadOnlyFanOut ||
		preview.FanOutResultState != "missing-receipt-fail-closed" ||
		!preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.SurfaceCount != 4 ||
		preview.ReceiptLookupState != "missing-receipt" ||
		preview.ReceiptAvailable ||
		preview.ReceiptConsumed ||
		preview.MaterializationReceiptConsumed ||
		preview.ExecutionSessionFanOutConsumed ||
		!preview.MissingReceiptSafe {
		t.Fatalf("unexpected materialization fan-out owner-route decision: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.StateRootPathExposed ||
		preview.StateRootWritesEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.FanOutWritesEnabled ||
		preview.ProductImageReady ||
		preview.ProductionTrustSatisfied ||
		preview.RuntimeWriteGateEnabled ||
		preview.LaunchPreflightPassed ||
		preview.LaunchAuthorized ||
		preview.ExecutionApproved ||
		preview.ProcessStartAuthorized ||
		preview.CommandMaterialized ||
		preview.ExecutablePathResolved ||
		preview.BackendSelectedForLaunch ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.TaskManagerEntryActive ||
		preview.LiveTrayBridgeEnabled ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.CompatibilityCenterActionsEnabled ||
		preview.RequestObjectsCreated ||
		preview.ProductionBusOwnership ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe materialization fan-out owner-route flags: %#v", preview)
	}
	for _, surface := range preview.Surfaces {
		if surface.MaterializationStatus != "missing-receipt" ||
			surface.SessionState != "missing-receipt" ||
			!surface.ReadOnly ||
			!surface.NavigationOnly ||
			!surface.UserVisible ||
			surface.MutatesRuntime ||
			surface.StartsProgram ||
			surface.DeliversNotification ||
			surface.ActivatesTaskManagerEntry ||
			surface.EnablesTrayBridge ||
			surface.EnablesCenterActions ||
			surface.ExposesStateRootPath ||
			surface.ExposesRawCommand ||
			surface.ExposesRawExecutable ||
			surface.ExposesBackendDetails {
			t.Fatalf("unexpected materialization fan-out owner-route surface: %#v", surface)
		}
	}
	if preview.CheckCount != 7 ||
		preview.PassedCheckCount != 7 ||
		!preview.AllChecksPassed ||
		!sameStrings(preview.CheckIDs, []string{"opaque-lookup-consumed", "caller-paths-hidden", "missing-receipt-fails-closed", "surface-fanout-deferred", "desktop-side-effects-disabled", "owner-route-ready-production-dbus-blocked", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected materialization fan-out owner-route checks: count=%d passed=%d ids=%#v", preview.CheckCount, preview.PassedCheckCount, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "KDE test launch materialization fan-out owner route test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestKDETestLaunchMaterializationFanOutOwnerRouteRejectsUnknownOpaqueID(t *testing.T) {
	if _, err := NewKDETestLaunchMaterializationFanOutOwnerRoutePreview(projectRootForRuntimeServiceBindingTest(t), "unknown-receipt"); err == nil {
		t.Fatal("materialization fan-out owner route accepted an unknown opaque id")
	}
}
