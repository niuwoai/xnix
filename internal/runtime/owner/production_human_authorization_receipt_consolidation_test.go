package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionHumanAuthorizationReceiptConsolidationPreviewConsumesProductionGates(t *testing.T) {
	preview, err := NewProductionHumanAuthorizationReceiptConsolidationPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionHumanAuthorizationReceiptConsolidationPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_human_authorization_receipt_consolidation.v1" ||
		preview.RequestType != "production-human-authorization-receipt-consolidation-preview" ||
		preview.ConsolidationType != "owner-managed-opaque-human-authorization-receipt-boundary" ||
		preview.ConsolidationDecision != "production-human-authorization-receipt-consolidation-ready-authorization-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production human authorization receipt consolidation schema: %#v", preview)
	}
	if !preview.ReceiptRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		!preview.ReceiptBoundaryConsolidated ||
		!preview.OwnerManagedOpaqueReceiptLookupReady ||
		preview.CallerStateRootRequired ||
		!preview.ExplicitOperatorActionRequired ||
		preview.AuthorizationGrantReady ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production human authorization receipt consolidation decision: %#v", preview)
	}
	if preview.GateCount != 7 ||
		preview.RequiredGateCount != 7 ||
		preview.ConsumedGateCount != 7 ||
		preview.MissingGateCount != 0 ||
		preview.AuthorizationAcceptedGateCount != 0 ||
		preview.ProductionReadyGateCount != 0 ||
		!sameStrings(preview.GateIDs, []string{"human-authorization-preflight", "production-dbus-gate-review", "production-dbus-method-review", "runtime-service-activation-preflight", "runtime-write-gate", "rollback-diagnostics-review", "desktop-side-effect-review"}) {
		t.Fatalf("unexpected production human authorization receipt gate inventory: counts=%#v ids=%#v", preview.GateCount, preview.GateIDs)
	}
	for _, gate := range preview.Gates {
		if !gate.EvidencePresent ||
			!gate.ReceiptBoundaryReady ||
			!gate.HumanAuthorizationRequired ||
			gate.AuthorizationReceiptAccepted ||
			gate.ProductionReadiness ||
			gate.ProductionOwnershipReady ||
			!gate.RuntimeOwned ||
			!gate.GoRuntimeBacked ||
			gate.KDEPolicyOwner ||
			!gate.ReviewOnly ||
			!gate.SideEffectsDisabled ||
			gate.WriteMethodsEnabled ||
			gate.RuntimeWritesEnabled ||
			gate.BackendLaunchEnabled ||
			gate.HostRootModified ||
			gate.InternalDetailsExposed ||
			gate.ConsolidationStatus != "consumed-authorization-disabled" {
			t.Fatalf("unsafe or unconsolidated production human authorization gate: %#v", gate)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"preflight-shape-consumed", "production-gate-consumed", "method-review-consumed", "service-and-write-gates-consumed", "rollback-and-desktop-reviews-consumed", "opaque-receipt-boundary", "authorization-not-granted", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected production human authorization receipt checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production human authorization receipt consolidation gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production human authorization receipt consolidation test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionHumanAuthorizationReceiptConsolidationPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionHumanAuthorizationReceiptConsolidationPreview(root)
	if err != nil {
		t.Fatalf("NewProductionHumanAuthorizationReceiptConsolidationPreview returned error: %v", err)
	}
	if preview.ReceiptBoundaryConsolidated ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.GateCount != 7 ||
		preview.RequiredGateCount != 7 ||
		preview.ConsumedGateCount != 0 ||
		preview.MissingGateCount != 7 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Blocked != 6 ||
		preview.ConsolidationDecision != "production-human-authorization-receipt-consolidation-blocked" {
		t.Fatalf("missing consolidation sources must fail closed: %#v", preview)
	}
	for _, gate := range preview.Gates {
		if gate.EvidencePresent ||
			gate.ReceiptBoundaryReady ||
			gate.AuthorizationReceiptAccepted ||
			gate.ProductionReadiness ||
			gate.ProductionOwnershipReady ||
			gate.ConsolidationStatus != "missing-source" {
			t.Fatalf("missing source gate must remain closed: %#v", gate)
		}
	}
}
