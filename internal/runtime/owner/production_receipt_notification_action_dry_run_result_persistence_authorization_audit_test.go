package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreviewModelsAuthorization(t *testing.T) {
	preview, err := NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.production_receipt_notification_action_dry_run_result_persistence_authorization_audit.v1" ||
		preview.RequestType != "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview" ||
		preview.AuditType != "receipt-notification-action-dry-run-result-persistence-authorization-audit" ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-ready-persistence-disabled" ||
		preview.ReceiptSchema != "xnix.runtime.production_dbus_human_authorization_receipt.v1" ||
		preview.OpaqueReceiptID != ProductionDBusHumanAuthorizationReceiptID {
		t.Fatalf("unexpected dry-run result persistence authorization schema: %#v", preview)
	}
	if !preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationAuditModeled ||
		!preview.ResultVisibilityAuditConsumed ||
		!preview.PersistenceAuthorizationGuidanceConsumed ||
		!preview.PersistenceAuthorizationReady ||
		preview.ResultPersistenceAuthorized ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.ResultVisibilityPersisted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorPersistenceApprovalRequired ||
		preview.OperatorPersistenceApprovalPresent ||
		preview.CallerStateRootRequired ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady {
		t.Fatalf("unexpected dry-run result persistence authorization decision: %#v", preview)
	}
	if preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 5 ||
		preview.MissingAuthorizationItemCount != 0 ||
		preview.GrantedPersistenceItemCount != 0 ||
		preview.PersistedResultItemCount != 0 ||
		preview.ExecutedDryRunItemCount != 0 ||
		preview.CreatedRequestObjectCount != 0 ||
		preview.PortalRequestCreatedCount != 0 ||
		preview.SideEffectAuthorizationItemCount != 0 ||
		!sameStrings(preview.AuthorizationItemIDs, []string{"review-receipt-dry-run-result-persistence-authorization", "renew-receipt-dry-run-result-persistence-authorization", "open-compatibility-center-dry-run-result-persistence-authorization", "dismiss-receipt-dry-run-result-persistence-authorization", "support-info-dry-run-result-persistence-authorization"}) {
		t.Fatalf("unexpected dry-run result persistence authorization inventory: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if !item.EvidencePresent ||
			!item.PersistenceAuthorizationModeled ||
			!item.RedactedForKDE ||
			!item.RuntimeDiagnosticsModeled ||
			!item.UserVisible ||
			!item.ReviewOnly ||
			!item.RuntimeOwned ||
			!item.GoRuntimeBacked ||
			item.KDEPolicyOwner ||
			!item.OperatorApprovalRequired ||
			item.OperatorApprovalPresent ||
			item.ResultPersistenceAuthorized ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.RequestObjectPersisted ||
			item.PortalRequestCreated ||
			item.NavigationRequested ||
			item.NotificationActionEnabled ||
			item.ActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.ReceiptWriterEnabled ||
			item.ReceiptPersistenceEnabled ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			!item.SideEffectsDisabled ||
			item.HostRootModified ||
			item.InternalDetailsExposed ||
			item.PersistenceAuthorizationStatus != "dry-run-result-persistence-authorization-modeled-persistence-disabled" {
			t.Fatalf("unsafe dry-run result persistence authorization item: %#v", item)
		}
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 9 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"result-visibility-audit-consumed", "persistence-authorization-guidance-consumed", "persistence-authorization-modeled-only", "five-persistence-authorizations-present", "persistence-authorizations-ready-persistence-disabled", "dry-run-execution-result-and-visibility-persistence-disabled", "request-creation-dispatch-actions-and-portal-disabled", "receipt-notification-navigation-and-support-writes-disabled", "production-and-host-boundary-closed"}) {
		t.Fatalf("unexpected dry-run result persistence authorization checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly ||
		preview.PlasmaForkRequired ||
		preview.PlasmaSourceModified ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.ProductionOwnerEnabled ||
		preview.ProductionActivationReady ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.RequestObjectCreationEnabled ||
		preview.RequestObjectDispatchEnabled ||
		preview.RequestObjectPersistenceEnabled ||
		preview.DispatchAuthorizationPersisted ||
		preview.DispatchDryRunExecutionEnabled ||
		preview.DryRunResultPersistenceEnabled ||
		preview.ResultVisibilityPersistenceEnabled ||
		preview.PortalRequestCreated ||
		preview.RequestObjectsCreated ||
		preview.RequestObjectsDispatched ||
		preview.NotificationSent ||
		preview.NotificationDeliveryEnabled ||
		preview.NotificationActionEnabled ||
		preview.ReviewActionEnabled ||
		preview.RenewActionEnabled ||
		preview.OpenCompatibilityCenterEnabled ||
		preview.DismissActionEnabled ||
		preview.SupportInfoActionEnabled ||
		preview.CompatibilityCenterOpened ||
		preview.CompatibilityCenterPersisted ||
		preview.RuntimeDiagnosticsPersisted ||
		preview.ReceiptWriterEnabled ||
		preview.ReceiptPersistenceEnabled ||
		preview.ReceiptLookupWritesEnabled ||
		preview.ReceiptReplayEnabled ||
		preview.ReceiptExpiryWriteEnabled ||
		preview.ReceiptRevocationWriteEnabled ||
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.ShellConfigurationWritten ||
		preview.SettingsPersisted ||
		preview.KRunnerIndexPersisted ||
		preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled ||
		preview.TrayBridgePersisted ||
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
		t.Fatalf("unsafe dry-run result persistence authorization gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "dry-run result persistence authorization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreviewFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(root)
	if err != nil {
		t.Fatalf("NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview returned error: %v", err)
	}
	if preview.ResultVisibilityAuditConsumed ||
		preview.PersistenceAuthorizationGuidanceConsumed ||
		!preview.PersistenceAuthorizationAuditRequired ||
		!preview.PersistenceAuthorizationAuditModeled ||
		preview.PersistenceAuthorizationReady ||
		preview.ResultPersistenceAuthorized ||
		preview.DispatchDryRunExecuted ||
		preview.DryRunResultPersisted ||
		preview.ResultVisibilityPersisted ||
		preview.DispatchAuthorizationGranted ||
		!preview.OperatorPersistenceApprovalRequired ||
		preview.OperatorPersistenceApprovalPresent ||
		preview.ReceiptPresent ||
		preview.ReceiptAccepted ||
		preview.AuthorizationAccepted ||
		preview.ProductionReadiness ||
		preview.ProductionOwnershipReady ||
		preview.AuthorizationItemCount != 5 ||
		preview.RequiredAuthorizationItemCount != 5 ||
		preview.ReadyAuthorizationItemCount != 0 ||
		preview.MissingAuthorizationItemCount != 5 ||
		preview.Counts.Total != 9 ||
		preview.Counts.Passed != 4 ||
		preview.Counts.Blocked != 5 ||
		preview.AuditDecision != "production-receipt-notification-action-dry-run-result-persistence-authorization-audit-blocked" {
		t.Fatalf("missing dry-run result persistence authorization sources must fail closed: %#v", preview)
	}
	for _, item := range preview.AuthorizationItems {
		if item.EvidencePresent ||
			item.PersistenceAuthorizationModeled ||
			item.RedactedForKDE ||
			item.RuntimeDiagnosticsModeled ||
			item.UserVisible ||
			item.OperatorApprovalPresent ||
			item.ResultPersistenceAuthorized ||
			item.DispatchAuthorizationGranted ||
			item.DispatchDryRunExecuted ||
			item.DryRunResultPersisted ||
			item.ResultVisibilityPersisted ||
			item.RequestObjectCreated ||
			item.RequestObjectDispatched ||
			item.RequestObjectPersisted ||
			item.PortalRequestCreated ||
			item.NavigationRequested ||
			item.NotificationActionEnabled ||
			item.ActionEnabled ||
			item.ReceiptAccepted ||
			item.AuthorizationAccepted ||
			item.CompatibilityCenterOpened ||
			item.CompatibilityCenterPersisted ||
			item.RuntimeDiagnosticsPersisted ||
			item.SupportBundleExported ||
			item.SupportCaseCreated ||
			item.ProductionReadiness ||
			item.ProductionOwnershipReady ||
			item.PersistenceAuthorizationStatus != "missing-dry-run-result-persistence-authorization-evidence" {
			t.Fatalf("missing source dry-run result persistence authorization item must remain closed: %#v", item)
		}
	}
}
