package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionDBusHumanAuthorizationPreflightPreviewDefinesReceiptShape(t *testing.T) {
	preview, err := NewProductionDBusHumanAuthorizationPreflightPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionDBusHumanAuthorizationPreflightPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_dbus_human_authorization_preflight.v1" ||
		preview.RequestType != "production-dbus-human-authorization-preflight-preview" ||
		preview.PreflightType != "read-only-production-dbus-human-authorization-preflight" ||
		!preview.GateReviewPresent ||
		!preview.RouteInventoryPresent ||
		!preview.ExplicitHumanAuthorizationRequired ||
		preview.AuthorizationReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.AuthorizationReceiptID != ProductionDBusHumanAuthorizationReceiptID ||
		!preview.AuthorizationReceiptRequired ||
		preview.AuthorizationReceiptPresent ||
		preview.AuthorizationGrantReady ||
		preview.AuthorizationAccepted ||
		!preview.PreflightReady ||
		preview.ProductionReadiness ||
		!preview.ProductionDBusGateReviewRequired {
		t.Fatalf("unexpected production D-Bus human authorization preflight schema: %#v", preview)
	}
	if preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"gate-review-present", "route-inventory-present", "human-authorization-required", "receipt-shape-declared", "authorization-not-granted", "production-ownership-disabled", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected production D-Bus human authorization preflight checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
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
		preview.AdapterInvocationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe production D-Bus human authorization preflight gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "production D-Bus human authorization preflight test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionDBusHumanAuthorizationPreflightPreviewFailsClosedWithoutGateReview(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionDBusHumanAuthorizationPreflightPreview(root)
	if err != nil {
		t.Fatalf("NewProductionDBusHumanAuthorizationPreflightPreview returned error: %v", err)
	}
	if preview.GateReviewPresent ||
		preview.RouteInventoryPresent ||
		preview.ExplicitHumanAuthorizationRequired ||
		preview.PreflightReady ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.Counts.Blocked != 3 {
		t.Fatalf("missing gate review sources must fail closed: %#v", preview)
	}
}
