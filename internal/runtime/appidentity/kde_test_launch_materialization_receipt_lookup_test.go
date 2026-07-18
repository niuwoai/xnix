package appidentity

import "testing"

func TestResolveKDETestLaunchMaterializationReceiptReturnsOpaqueLookup(t *testing.T) {
	lookup, err := ResolveKDETestLaunchMaterializationReceipt(projectRootForRuntimeServiceBindingTest(t), KDETestLaunchMaterializationOpaqueReceiptID)
	if err != nil {
		t.Fatalf("ResolveKDETestLaunchMaterializationReceipt returned error: %v", err)
	}

	if lookup.Version != currentProjectVersion(t) ||
		lookup.SchemaVersion != "xnix.runtime.kde_test_launch_materialization_receipt_lookup.v1" ||
		lookup.RequestType != "kde-test-launch-materialization-receipt-lookup-preview" ||
		lookup.LookupType != "owner-managed-materialization-receipt-lookup" ||
		lookup.RuntimeMethod != "GetKDETestLaunchMaterializationReceiptLookup" ||
		lookup.ReadMethod != "GetKDETestLaunchMaterializationReceiptLookupPreview" ||
		lookup.OpaqueMaterializationReceiptID != KDETestLaunchMaterializationOpaqueReceiptID ||
		lookup.SupportedOpaqueMaterializationReceiptID != "kde-test-launch-materialization-receipt-id" ||
		lookup.ReceiptRelativePath != "execution-ledger/materialization-plans/kde-test-launch-materialization-receipt.json" ||
		lookup.ReceiptRecordType != "restricted-launch-materialization-plan" {
		t.Fatalf("unexpected materialization receipt lookup schema: %#v", lookup)
	}
	if lookup.ReceiptLookupState != "missing-receipt" ||
		lookup.ReceiptAvailable ||
		lookup.ReceiptConsumed ||
		!lookup.MissingReceiptSafe ||
		!lookup.OwnerManagedOpaqueReceiptLookupReady ||
		!lookup.OpaqueMaterializationReceiptIDSupported ||
		lookup.RequiresCallerRegistryPath ||
		lookup.RequiresCallerApplicationID ||
		lookup.RequiresCallerStateRoot ||
		lookup.RequiresExplicitAuthorization ||
		!lookup.ReadOnlyLookup {
		t.Fatalf("unexpected materialization receipt lookup decision: %#v", lookup)
	}
	if !lookup.RuntimeOwned ||
		!lookup.GoRuntimeBacked ||
		lookup.KDEPolicyOwner ||
		lookup.StateRootPathExposed ||
		lookup.StateRootWritesEnabled ||
		lookup.RuntimeWritesEnabled ||
		lookup.MaterializationWritesEnabled ||
		lookup.FanOutWritesEnabled ||
		lookup.SystemServiceStarted ||
		lookup.SessionBusClaimed ||
		lookup.ProductionBusClaimed ||
		lookup.WriteMethodsEnabled ||
		lookup.LaunchAuthorized ||
		lookup.ExecutionApproved ||
		lookup.ProcessStartAuthorized ||
		lookup.CommandMaterialized ||
		lookup.ExecutablePathResolved ||
		lookup.BackendSelectedForLaunch ||
		lookup.BackendLaunchEnabled ||
		lookup.BackendProcessStarted ||
		lookup.NetworkRequired ||
		lookup.HostRootModified ||
		lookup.PrivilegedContainerRequired ||
		lookup.RawCommandExposed ||
		lookup.RawExecutableExposed ||
		lookup.BackendDetailsExposed {
		t.Fatalf("unsafe materialization receipt lookup gates: %#v", lookup)
	}
	if lookup.CheckCount != 6 ||
		lookup.PassedCheckCount != 6 ||
		!lookup.AllChecksPassed ||
		!sameStrings(lookup.CheckIDs, []string{"opaque-materialization-receipt-id-supported", "owner-managed-opaque-lookup", "caller-paths-hidden", "missing-receipt-fails-closed", "read-only-lookup", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected materialization receipt lookup checks: count=%d passed=%d ids=%#v", lookup.CheckCount, lookup.PassedCheckCount, lookup.CheckIDs)
	}
	if err := validateNoBackendTerms(lookup, "KDE test launch materialization receipt lookup test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestResolveKDETestLaunchMaterializationReceiptRejectsUnknownOpaqueID(t *testing.T) {
	if _, err := ResolveKDETestLaunchMaterializationReceipt(projectRootForRuntimeServiceBindingTest(t), "unknown-receipt"); err == nil {
		t.Fatal("materialization receipt lookup accepted an unknown opaque id")
	}
}
