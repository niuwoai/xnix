package appidentity

import (
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestKDETestLaunchMaterializationFanOutPreviewCoversKDESurfaces(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	preview, err := NewKDETestLaunchMaterializationFanOutPreview(recipeRecord, provenance, KDERestrictedLaunchAuthorizationOptions{StateRoot: t.TempDir(), Mode: "test-only", Directive: execution.RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationFanOutPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.kde_test_launch_materialization_fanout.v1" ||
		preview.RequestType != "kde-test-launch-materialization-fanout-preview" ||
		preview.Mode != "test-only" ||
		preview.MaterializationStatus != "blocked-plan-materialized" ||
		preview.MaterializationScope != "test-only-review-plan" ||
		preview.SurfaceCount != 4 ||
		preview.CheckCount != 9 ||
		preview.PassedCheckCount != 9 ||
		!preview.AllChecksPassed ||
		!preview.MaterializationReceiptConsumed ||
		!preview.ExecutionSessionFanOutConsumed ||
		!preview.TestOnly {
		t.Fatalf("unexpected materialization fan-out preview: %+v", preview)
	}
	if filepath.IsAbs(preview.MaterializationReceiptRelativePath) ||
		filepath.IsAbs(preview.ExecutionSessionReceiptPath) ||
		len(preview.MaterializationReceiptSHA256) != 64 {
		t.Fatalf("fan-out must expose relative receipts and materialization digest only: %+v", preview)
	}
	for _, surface := range preview.Surfaces {
		if !surface.ReadOnly ||
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
			t.Fatalf("surface must be a safe read-only KDE projection: %+v", surface)
		}
	}
	if preview.CompatibilityCenter.ID != "compatibility-center" ||
		preview.TaskManager.ID != "task-manager" ||
		preview.Tray.ID != "tray" ||
		preview.Notification.ID != "notification" ||
		preview.TaskManager.SessionState != "blocked" ||
		preview.Tray.SessionState != "blocked" ||
		preview.Notification.SessionState != "review-plan-available" {
		t.Fatalf("unexpected KDE surface projections: %+v", preview)
	}
	if preview.StateRootPathExposed ||
		!preview.StateRootWritesEnabled ||
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
		preview.RuntimeWritesEnabled ||
		preview.ProductionBusOwnership ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe materialization fan-out gates: %+v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE test launch materialization fan-out test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestKDETestLaunchMaterializationFanOutPreviewRequiresExactBoundary(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	for _, options := range []KDERestrictedLaunchAuthorizationOptions{
		{StateRoot: t.TempDir(), Mode: "", Directive: execution.RestrictedTestPreparationDirective},
		{StateRoot: t.TempDir(), Mode: "test-only", Directive: ""},
		{StateRoot: t.TempDir(), Mode: "test-only", Directive: "authorize"},
	} {
		if _, err := NewKDETestLaunchMaterializationFanOutPreview(recipeRecord, provenance, options); err == nil {
			t.Fatalf("expected exact restricted fan-out boundary to be required for options: %+v", options)
		}
	}
}
