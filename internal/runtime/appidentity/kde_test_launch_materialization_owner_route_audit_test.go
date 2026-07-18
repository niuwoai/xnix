package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKDETestLaunchMaterializationOwnerRouteAuditSeesConsumeSplit(t *testing.T) {
	preview, err := NewKDETestLaunchMaterializationOwnerRouteAuditPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationOwnerRouteAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.kde_test_launch_materialization_owner_route_audit.v1" ||
		preview.RequestType != "kde-test-launch-materialization-owner-route-audit-preview" ||
		preview.AuditType != "materialization-fanout-owner-route-audit" ||
		preview.RuntimeMethod != "GetKDETestLaunchMaterializationOwnerRouteAudit" ||
		preview.ReadMethod != "GetKDETestLaunchMaterializationOwnerRouteAuditPreview" ||
		preview.SubjectRequestType != "kde-test-launch-materialization-fanout-preview" ||
		preview.SubjectCommand != "kde-test-launch-materialization-fanout-preview" ||
		preview.ProposedOwnerMethod != "GetKDETestLaunchMaterializationFanOut" ||
		preview.RouteDecision != "owner-local-route-smoke-covered" ||
		preview.CurrentRouteStatus != "owner-local-read-route-smoke-covered-production-dbus-blocked" {
		t.Fatalf("unexpected materialization owner-route audit schema: %#v", preview)
	}
	if !preview.CLICommandRegistered ||
		!preview.GoReadModelPresent ||
		!preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		!preview.RequiresCallerRegistryPath ||
		!preview.RequiresCallerApplicationID ||
		!preview.RequiresCallerStateRoot ||
		!preview.RequiresExplicitAuthorization ||
		!preview.MaterializationWritesStateRoot ||
		!preview.ReadOnlyConsumeCommandRegistered ||
		!preview.ReadOnlyReceiptConsumptionReady ||
		!preview.ReadOnlyConsumeRequiresCallerStateRoot ||
		!preview.OwnerManagedOpaqueReceiptLookupReady ||
		!preview.OpaqueMaterializationReceiptIDSupported ||
		!preview.OwnerSmokeCoverageReady ||
		preview.FanOutWritesEnabled ||
		!preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.RecommendedNextRoute != "materialization-fanout-production-dbus-gate-review" {
		t.Fatalf("unexpected materialization owner-route audit decision: %#v", preview)
	}
	if preview.Counts.Total != 11 ||
		preview.Counts.Passed != 11 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"cli-preview-registered", "go-read-model-present", "read-only-consume-registered", "owner-route-present", "production-dbus-absent", "caller-path-boundary", "receipt-creation-split", "owner-managed-opaque-lookup", "owner-smoke-coverage-present", "route-decision", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected materialization owner-route audit checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe materialization owner-route audit gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE test launch materialization owner-route audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestKDETestLaunchMaterializationOwnerRouteAuditFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewKDETestLaunchMaterializationOwnerRouteAuditPreview(root)
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationOwnerRouteAuditPreview returned error: %v", err)
	}
	if preview.CLICommandRegistered ||
		preview.GoReadModelPresent ||
		preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		preview.OwnerSmokeCoverageReady ||
		preview.ReadOnlyConsumeCommandRegistered ||
		preview.ReadOnlyReceiptConsumptionReady ||
		preview.OwnerManagedOpaqueReceiptLookupReady ||
		preview.OpaqueMaterializationReceiptIDSupported ||
		preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady {
		t.Fatalf("missing sources must not produce owner-route readiness: %#v", preview)
	}
	if preview.Counts.Total != 11 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Pending != 3 ||
		preview.Counts.Blocked != 6 ||
		preview.RouteDecision != "materialization-owner-route-sources-missing" {
		t.Fatalf("missing sources must fail closed: %#v", preview.Counts)
	}
}
