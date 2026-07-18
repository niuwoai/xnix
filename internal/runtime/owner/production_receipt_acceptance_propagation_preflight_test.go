package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptAcceptancePropagationPreflightPreviewModelsFutureAcceptance(t *testing.T) {
	preview, err := NewProductionReceiptAcceptancePropagationPreflightPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptAcceptancePropagationPreflightPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_acceptance_propagation_preflight.v1" ||
		preview.RequestType != "production-receipt-acceptance-propagation-preflight-preview" ||
		preview.PreflightType != "future-authorization-receipt-acceptance-propagation-preflight" ||
		preview.PreflightDecision != "production-receipt-acceptance-propagation-ready-acceptance-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production receipt acceptance propagation preflight schema: %#v", preview)
	}
	if !preview.ReceiptRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		!preview.FutureAcceptanceModeled ||
		!preview.AcceptanceSimulationOnly ||
		!preview.ConsumptionAuditConsumed ||
		!preview.OwnerManagedOpaqueBoundaryReady ||
		preview.CallerStateRootRequired ||
		preview.AuthorizationAccepted ||
		!preview.PropagationPreflightReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production receipt acceptance propagation decision: %#v", preview)
	}
	if preview.TargetCount != 6 ||
		preview.RequiredTargetCount != 6 ||
		preview.PropagationReadyTargetCount != 6 ||
		preview.MissingTargetCount != 0 ||
		preview.AcceptanceEnabledTargetCount != 0 ||
		preview.ProductionReadyTargetCount != 0 ||
		preview.SideEffectTargetCount != 0 ||
		!sameStrings(preview.TargetIDs, []string{"production-dbus-gate-review", "production-dbus-method-review", "runtime-service-activation-preflight", "runtime-write-gate", "rollback-diagnostics-review", "desktop-side-effect-review"}) {
		t.Fatalf("unexpected production receipt acceptance target inventory: counts=%#v ids=%#v", preview.TargetCount, preview.TargetIDs)
	}
	for _, target := range preview.Targets {
		if !target.ConsumesAuditBoundary ||
			!target.PropagatesFutureAcceptance ||
			!target.ReceiptBoundaryReady ||
			!target.FutureAcceptanceModeled ||
			target.ReceiptAccepted ||
			target.AuthorizationAccepted ||
			target.ProductionReadiness ||
			target.ProductionOwnershipReady ||
			!target.RuntimeOwned ||
			!target.GoRuntimeBacked ||
			target.KDEPolicyOwner ||
			!target.ReviewOnly ||
			target.WriteMethodsEnabled ||
			target.RuntimeWritesEnabled ||
			target.DesktopSideEffectsEnabled ||
			target.SupportSideEffectsEnabled ||
			target.BackendLaunchEnabled ||
			target.HostRootModified ||
			target.InternalDetailsExposed ||
			target.PropagationStatus != "future-acceptance-modeled-side-effects-disabled" {
			t.Fatalf("unsafe or missing production receipt acceptance target: %#v", target)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"consumption-audit-consumed", "future-acceptance-modeled-only", "six-propagation-targets-present", "targets-propagate-future-acceptance", "acceptance-and-production-disabled", "production-ownership-disabled", "write-and-launch-disabled", "desktop-support-and-host-boundary-closed"}) {
		t.Fatalf("unexpected production receipt acceptance propagation checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.ShellConfigurationWritten ||
		preview.SettingsPersisted ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.PortalRequestCreated ||
		preview.RequestObjectsCreated ||
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.SnapshotRestoreExecuted ||
		preview.StateCleanupExecuted ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production receipt acceptance propagation preflight gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production receipt acceptance propagation preflight test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptAcceptancePropagationPreflightPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptAcceptancePropagationPreflightPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptAcceptancePropagationPreflightPreview returned error: %v", err)
	}
	if preview.ConsumptionAuditConsumed ||
		preview.OwnerManagedOpaqueBoundaryReady ||
		!preview.FutureAcceptanceModeled ||
		!preview.AcceptanceSimulationOnly ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.PropagationPreflightReady ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.TargetCount != 6 ||
		preview.RequiredTargetCount != 6 ||
		preview.PropagationReadyTargetCount != 0 ||
		preview.MissingTargetCount != 6 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Blocked != 3 ||
		preview.PreflightDecision != "production-receipt-acceptance-propagation-preflight-blocked" {
		t.Fatalf("missing propagation sources must fail closed: %#v", preview)
	}
	for _, target := range preview.Targets {
		if target.ConsumesAuditBoundary ||
			target.PropagatesFutureAcceptance ||
			target.ReceiptBoundaryReady ||
			target.FutureAcceptanceModeled ||
			target.ReceiptAccepted ||
			target.AuthorizationAccepted ||
			target.ProductionReadiness ||
			target.ProductionOwnershipReady ||
			target.PropagationStatus != "missing-propagation-source" {
			t.Fatalf("missing source target must remain closed: %#v", target)
		}
	}
}
