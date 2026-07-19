package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreviewModelsPolicy(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_lookup_consumer_enablement_kde_safe_redacted_status_closed_record_writer_storage_root_policy_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview" ||
		preview.PreviewType != "receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit" ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-ready-writes-disabled" {
		t.Fatalf("unexpected KDE-safe redacted status storage-root policy schema: %#v", preview)
	}
	if !preview.CurrentMainlineConsumed ||
		!preview.ClosedRecordWriterConsumed ||
		!preview.ClosedRecordWriterReady ||
		!preview.StorageRootPolicyRequired ||
		!preview.StorageRootPolicyModeled ||
		!preview.StorageRootPolicyReady ||
		!preview.StorageRootOwnershipModeled ||
		!preview.StorageRootNamespaceModeled ||
		!preview.StorageRootRetentionModeled ||
		!preview.StorageRootRedactionModeled ||
		!preview.RecordWriterCallGuardModeled ||
		!preview.KDESafeRedactedStatusOnly ||
		!preview.CompatibilityCenterPolicyModeled ||
		!preview.RuntimeDiagnosticsPolicyModeled ||
		preview.WriterCallable ||
		preview.ClosedRecordWriterEnabled ||
		preview.StorageRootResolved ||
		preview.StorageRootCreated ||
		preview.StorageRootMounted ||
		preview.StorageRootOwnershipGranted ||
		preview.StorageRootNamespaceGranted ||
		preview.StorageRetentionEnforced ||
		preview.StorageRedactionEnforced ||
		preview.StorageGatePassed ||
		preview.StoragePersistenceAuthorized ||
		preview.DurableRecordWriteEnabled ||
		preview.WriterAuthorizationGranted ||
		preview.StatusWriterEnabled ||
		preview.StatusPersistenceAuthorized ||
		preview.StatusPersistenceWriteEnabled ||
		preview.KDEStatusWriteEnabled ||
		preview.RuntimeDiagnosticsWriteEnabled ||
		preview.StorageWriteEnabled ||
		preview.ConsumerConsumptionAuthorized ||
		preview.ConsumerEnablementAuthorized ||
		preview.KDEConsumerEnabled ||
		preview.RuntimeConsumerEnabled ||
		preview.LookupRouteEnabled ||
		preview.OpaqueLookupEnabled ||
		preview.RedactedSummaryPersisted ||
		preview.KDEStatusPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.DryRunResultPersisted ||
		preview.RawResultExposed ||
		preview.DispatchDryRunExecuted ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.PortalRequestCreated ||
		preview.NotificationActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unsafe KDE-safe redacted status storage-root policy decision: %#v", preview)
	}
	if preview.PolicyItemCount != 10 ||
		preview.RequiredPolicyItemCount != 10 ||
		preview.ReadyPolicyItemCount != 10 ||
		preview.MissingPolicyItemCount != 0 ||
		preview.ResolvedStorageRootItemCount != 0 ||
		preview.GrantedStorageRootItemCount != 0 ||
		preview.EnforcedRetentionItemCount != 0 ||
		preview.EnforcedRedactionItemCount != 0 ||
		preview.CallableWriterItemCount != 0 ||
		preview.EnabledRecordWriterItemCount != 0 ||
		preview.DurableWrittenRecordItemCount != 0 ||
		preview.RawExposedPolicyItemCount != 0 ||
		preview.SideEffectPolicyItemCount != 0 ||
		preview.CompatibilityCenterItemCount != 5 ||
		preview.RuntimeDiagnosticsItemCount != 5 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root policy counts: %#v", preview)
	}
	if len(preview.PolicyItems) != 10 {
		t.Fatalf("unexpected KDE-safe redacted status storage-root policy items: %#v", preview.PolicyItemIDs)
	}
	for _, item := range preview.PolicyItems {
		if !item.EvidencePresent ||
			!item.CurrentMainlineConsumed ||
			!item.ClosedRecordWriterConsumed ||
			!item.ClosedRecordWriterReady ||
			!item.StorageRootPolicyModeled ||
			!item.StorageRootOwnershipModeled ||
			!item.StorageRootNamespaceModeled ||
			!item.StorageRootRetentionModeled ||
			!item.StorageRootRedactionModeled ||
			!item.RecordWriterCallGuardModeled ||
			!item.KDESafeRedactedStatusOnly ||
			item.NamespaceClass != "owner-managed-redacted-status-records" ||
			item.RetentionClass != "audit-preview-retention-window" ||
			item.RedactionClass != "kde-safe-redacted-status-only" ||
			item.WriterCallable ||
			item.ClosedRecordWriterEnabled ||
			item.StorageRootResolved ||
			item.StorageRootCreated ||
			item.StorageRootMounted ||
			item.StorageRootOwnershipGranted ||
			item.StorageRootNamespaceGranted ||
			item.StorageRetentionEnforced ||
			item.StorageRedactionEnforced ||
			item.StorageGatePassed ||
			item.StoragePersistenceAuthorized ||
			item.DurableRecordWriteEnabled ||
			item.StatusPersistenceAuthorized ||
			item.StatusPersistenceWriteEnabled ||
			item.StatusWriterEnabled ||
			item.KDEStatusWriteEnabled ||
			item.RuntimeDiagnosticsWriteEnabled ||
			item.StorageWriteEnabled ||
			item.RedactedSummaryPersisted ||
			item.KDEStatusPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.RawResultExposed ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.StorageRootPolicyStatus != "redacted-status-storage-root-policy-modeled-writes-disabled" {
			t.Fatalf("unsafe KDE-safe redacted status storage-root policy item: %#v", item)
		}
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-consumed", "closed-record-writer-consumed", "storage-root-policy-modeled", "compatibility-center-and-runtime-storage-root-policies-modeled", "ten-policy-items-ready-writes-disabled", "storage-root-and-writer-writes-disabled", "consumer-lookup-request-and-support-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected KDE-safe redacted status storage-root policy checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe KDE-safe redacted status storage-root policy gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "KDE-safe redacted status storage-root policy test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview returned error: %v", err)
	}
	if preview.CurrentMainlineConsumed ||
		preview.ClosedRecordWriterConsumed ||
		preview.ClosedRecordWriterReady ||
		!preview.StorageRootPolicyRequired ||
		preview.StorageRootPolicyModeled ||
		preview.StorageRootPolicyReady ||
		preview.StorageRootOwnershipModeled ||
		preview.StorageRootNamespaceModeled ||
		preview.StorageRootRetentionModeled ||
		preview.StorageRootRedactionModeled ||
		preview.RecordWriterCallGuardModeled ||
		preview.KDESafeRedactedStatusOnly ||
		preview.CompatibilityCenterPolicyModeled ||
		preview.RuntimeDiagnosticsPolicyModeled ||
		preview.WriterCallable ||
		preview.StorageRootResolved ||
		preview.StorageRootCreated ||
		preview.StorageRootOwnershipGranted ||
		preview.StorageRetentionEnforced ||
		preview.StorageRedactionEnforced ||
		preview.StorageWriteEnabled ||
		preview.RawResultExposed ||
		preview.PolicyItemCount != 10 ||
		preview.RequiredPolicyItemCount != 10 ||
		preview.ReadyPolicyItemCount != 0 ||
		preview.MissingPolicyItemCount != 10 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Blocked != 5 ||
		preview.PreviewDecision != "production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-blocked" {
		t.Fatalf("missing KDE-safe redacted status storage-root policy sources must fail closed: %#v", preview)
	}
	for _, item := range preview.PolicyItems {
		if item.EvidencePresent ||
			item.CurrentMainlineConsumed ||
			item.ClosedRecordWriterConsumed ||
			item.ClosedRecordWriterReady ||
			item.StorageRootPolicyModeled ||
			item.StorageRootOwnershipModeled ||
			item.StorageRootNamespaceModeled ||
			item.StorageRootRetentionModeled ||
			item.StorageRootRedactionModeled ||
			item.RecordWriterCallGuardModeled ||
			item.WriterCallable ||
			item.StorageRootResolved ||
			item.StorageRootCreated ||
			item.StorageRootOwnershipGranted ||
			item.StorageRetentionEnforced ||
			item.StorageRedactionEnforced ||
			item.StorageWriteEnabled ||
			item.RawResultExposed ||
			item.UserVisible ||
			item.StorageRootPolicyStatus != "missing-redacted-status-storage-root-policy-evidence" {
			t.Fatalf("missing KDE-safe redacted status storage-root policy item must remain closed: %#v", item)
		}
	}
}
