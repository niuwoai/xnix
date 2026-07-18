package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditKeepsFanOutCLIOnly(t *testing.T) {
	preview, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(projectRoot(t))
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route_audit.v1" ||
		preview.RequestType != "restricted-owner-smoke-receipt-fanout-owner-route-audit-preview" ||
		preview.AuditType != "restricted-owner-smoke-receipt-fanout-owner-route-audit" ||
		preview.RuntimeMethod != "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAudit" ||
		preview.ReadMethod != "GetRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview" ||
		preview.SubjectRequestType != "restricted-owner-smoke-receipt-fanout-preview" ||
		preview.SubjectCommand != "restricted-owner-smoke-receipt-fanout-preview" ||
		preview.ProposedOwnerMethod != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		preview.RouteDecision != "remain-cli-only" ||
		preview.CurrentRouteStatus != "cli-preview-ready-owner-route-blocked" {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit schema: %#v", preview)
	}
	if !preview.CLICommandRegistered ||
		!preview.GoReadModelPresent ||
		preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		!preview.ConsumesExistingReceipt ||
		!preview.RequiresCallerStateRoot ||
		preview.ReceiptLookupOwnerManaged ||
		preview.OpaqueReceiptIDSupported ||
		preview.FanOutWritesEnabled ||
		preview.StateRootWritesEnabled ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.RecommendedNextRoute != "owner-local-read-route-after-opaque-receipt-lookup" {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit decision: %#v", preview)
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"cli-preview-registered", "go-read-model-present", "owner-route-absent", "production-dbus-absent", "receipt-consumption-present", "caller-state-root-boundary", "owner-managed-receipt-lookup", "route-decision", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected restricted owner smoke fan-out owner-route audit checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe restricted owner smoke fan-out owner-route audit gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out owner-route audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(root)
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview returned error: %v", err)
	}
	if preview.CLICommandRegistered ||
		preview.GoReadModelPresent ||
		preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		preview.ConsumesExistingReceipt ||
		preview.ReceiptLookupOwnerManaged ||
		preview.OpaqueReceiptIDSupported ||
		preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady {
		t.Fatalf("missing sources must not produce owner-route readiness: %#v", preview)
	}
	if preview.Counts.Blocked != 3 || preview.RouteDecision != "remain-cli-only" {
		t.Fatalf("missing sources must fail closed while keeping the decision CLI-only: %#v", preview.Counts)
	}
}
