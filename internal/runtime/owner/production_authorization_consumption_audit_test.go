package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionAuthorizationConsumptionAuditPreviewVerifiesConsumers(t *testing.T) {
	preview, err := NewProductionAuthorizationConsumptionAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionAuthorizationConsumptionAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_authorization_consumption_audit.v1" ||
		preview.RequestType != "production-authorization-consumption-audit-preview" ||
		preview.AuditType != "production-gate-consolidated-authorization-consumption-audit" ||
		preview.AuditDecision != "production-authorization-consumption-audit-ready-authorization-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected production authorization consumption audit schema: %#v", preview)
	}
	if !preview.ReceiptRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		!preview.ReceiptBoundaryConsolidated ||
		!preview.ConsolidationPreviewConsumed ||
		!preview.OwnerManagedOpaqueBoundaryReady ||
		preview.CallerStateRootRequired ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected production authorization consumption audit decision: %#v", preview)
	}
	if preview.ConsumerCount != 6 ||
		preview.RequiredConsumerCount != 6 ||
		preview.ConsumedConsumerCount != 6 ||
		preview.MissingConsumerCount != 0 ||
		preview.AuthorizationAcceptedConsumerCount != 0 ||
		preview.ProductionReadyConsumerCount != 0 ||
		preview.SideEffectConsumerCount != 0 ||
		!sameStrings(preview.ConsumerIDs, []string{"production-dbus-gate-review", "production-dbus-method-review", "runtime-service-activation-preflight", "runtime-write-gate", "rollback-diagnostics-review", "desktop-side-effect-review"}) {
		t.Fatalf("unexpected production authorization consumption inventory: counts=%#v ids=%#v", preview.ConsumerCount, preview.ConsumerIDs)
	}
	for _, consumer := range preview.Consumers {
		if !consumer.ConsumesConsolidatedBoundary ||
			!consumer.ReceiptBoundaryReady ||
			consumer.ReceiptAccepted ||
			consumer.AuthorizationAccepted ||
			consumer.ProductionReadiness ||
			consumer.ProductionOwnershipReady ||
			!consumer.RuntimeOwned ||
			!consumer.GoRuntimeBacked ||
			consumer.KDEPolicyOwner ||
			!consumer.ReviewOnly ||
			consumer.WriteMethodsEnabled ||
			consumer.RuntimeWritesEnabled ||
			consumer.DesktopSideEffectsEnabled ||
			consumer.SupportSideEffectsEnabled ||
			consumer.BackendLaunchEnabled ||
			consumer.HostRootModified ||
			consumer.InternalDetailsExposed ||
			consumer.AuditStatus != "consumed-authorization-disabled" {
			t.Fatalf("unsafe or missing production authorization consumer: %#v", consumer)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"consolidation-preview-consumed", "six-production-consumers-present", "consumers-use-consolidated-boundary", "authorization-not-accepted", "production-ownership-disabled", "write-and-launch-disabled", "desktop-and-support-side-effects-disabled", "host-boundary-closed"}) {
		t.Fatalf("unexpected production authorization consumption checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		t.Fatalf("unsafe production authorization consumption audit gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production authorization consumption audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionAuthorizationConsumptionAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionAuthorizationConsumptionAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionAuthorizationConsumptionAuditPreview returned error: %v", err)
	}
	if preview.ReceiptBoundaryConsolidated ||
		preview.ConsolidationPreviewConsumed ||
		preview.OwnerManagedOpaqueBoundaryReady ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.ConsumerCount != 6 ||
		preview.RequiredConsumerCount != 6 ||
		preview.ConsumedConsumerCount != 0 ||
		preview.MissingConsumerCount != 6 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Blocked != 3 ||
		preview.AuditDecision != "production-authorization-consumption-audit-blocked" {
		t.Fatalf("missing audit sources must fail closed: %#v", preview)
	}
	for _, consumer := range preview.Consumers {
		if consumer.ConsumesConsolidatedBoundary ||
			consumer.ReceiptBoundaryReady ||
			consumer.ReceiptAccepted ||
			consumer.AuthorizationAccepted ||
			consumer.ProductionReadiness ||
			consumer.ProductionOwnershipReady ||
			consumer.AuditStatus != "missing-consumption" {
			t.Fatalf("missing source consumer must remain closed: %#v", consumer)
		}
	}
}
