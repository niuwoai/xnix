package owner

import (
	"path/filepath"
	"testing"
)

func TestResolveRestrictedOwnerSmokeReceiptReturnsOpaqueLookup(t *testing.T) {
	lookup, err := ResolveRestrictedOwnerSmokeReceipt(projectRoot(t), RestrictedOwnerSmokeOpaqueReceiptID)
	if err != nil {
		t.Fatalf("ResolveRestrictedOwnerSmokeReceipt returned error: %v", err)
	}

	if lookup.Version != currentProjectVersion(t) ||
		lookup.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt_lookup.v1" ||
		lookup.RequestType != "restricted-owner-smoke-receipt-lookup-preview" ||
		lookup.LookupType != "owner-managed-receipt-lookup" ||
		lookup.RuntimeMethod != "GetRestrictedOwnerSmokeReceiptLookup" ||
		lookup.ReadMethod != "GetRestrictedOwnerSmokeReceiptLookupPreview" ||
		lookup.OpaqueReceiptID != RestrictedOwnerSmokeOpaqueReceiptID ||
		lookup.SupportedOpaqueReceiptID != RestrictedOwnerSmokeOpaqueReceiptID ||
		lookup.ReceiptID != "restricted-owner-smoke-"+currentProjectVersion(t) ||
		lookup.ReceiptRelativePath != "owner-smoke/restricted-owner-smoke-receipt.json" ||
		lookup.ReceiptRecordType != "restricted-owner-smoke-execution-receipt" ||
		lookup.ReceiptLookupState != "missing-receipt" ||
		lookup.ReceiptAvailable ||
		lookup.ReceiptConsumed ||
		!lookup.MissingReceiptSafe ||
		!lookup.OwnerManagedLookup ||
		lookup.CallerStateRootRequired ||
		!lookup.OpaqueReceiptIDSupported ||
		!lookup.ReadOnlyLookup {
		t.Fatalf("unexpected restricted owner smoke receipt lookup: %#v", lookup)
	}
	if filepath.IsAbs(lookup.ReceiptRelativePath) {
		t.Fatalf("lookup must not expose an absolute receipt path: %s", lookup.ReceiptRelativePath)
	}
	if lookup.CheckCount != 6 ||
		lookup.PassedCheckCount != 6 ||
		!lookup.AllChecksPassed ||
		!sameStrings(lookup.CheckIDs, []string{"opaque-receipt-id-supported", "owner-managed-lookup", "caller-state-root-hidden", "receipt-slot-redacted", "missing-receipt-fails-closed", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected restricted owner smoke lookup checks: count=%d passed=%d ids=%#v", lookup.CheckCount, lookup.PassedCheckCount, lookup.CheckIDs)
	}
	if lookup.StateRootPathExposed ||
		lookup.StateRootWritesEnabled ||
		lookup.RuntimeWritesEnabled ||
		lookup.FanOutWritesEnabled ||
		lookup.ProductionOwnerEnabled ||
		lookup.SystemServiceStarted ||
		lookup.SessionBusClaimed ||
		lookup.ProductionBusClaimed ||
		lookup.WriteMethodsEnabled ||
		lookup.SupportBundleExported ||
		lookup.SupportCaseCreated ||
		lookup.NotificationSent ||
		lookup.BackendLaunchEnabled ||
		lookup.BackendProcessStarted ||
		lookup.NetworkRequired ||
		lookup.HostRootModified ||
		lookup.PrivilegedContainerRequired ||
		lookup.BackendDetailsExposed {
		t.Fatalf("unsafe restricted owner smoke lookup gates: %#v", lookup)
	}
	if err := validateNoBackendTerms(lookup, "restricted owner smoke receipt lookup test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestResolveRestrictedOwnerSmokeReceiptRejectsUnknownOpaqueID(t *testing.T) {
	if _, err := ResolveRestrictedOwnerSmokeReceipt(projectRoot(t), "owner-smoke/restricted-owner-smoke-receipt.json"); err == nil {
		t.Fatalf("lookup must reject path-like receipt identifiers")
	}
	if _, err := ResolveRestrictedOwnerSmokeReceipt(projectRoot(t), "unknown-receipt"); err == nil {
		t.Fatalf("lookup must reject unknown opaque receipt identifiers")
	}
}
