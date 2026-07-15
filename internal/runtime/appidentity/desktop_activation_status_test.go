package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationStatusPreviewSummarizesActivationWithoutWriting(t *testing.T) {
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}

	preview, err := plan.DesktopActivationStatusPreview("development")
	if err != nil {
		t.Fatalf("DesktopActivationStatusPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_status.v1" ||
		preview.RequestType != "desktop-activation-status-preview" ||
		preview.StatusType != "kde-desktop-activation-status" ||
		preview.Source != "desktop-activation-transaction-preview" ||
		preview.PlannedRuntimeMethod != "GetDesktopActivationStatus" ||
		preview.CurrentRenderer != "xnix-runtime-go desktop-activation-status-preview" ||
		preview.WriteMethod != "ActivateDesktopIntegration" {
		t.Fatalf("unexpected status schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.InstallMode != "development" ||
		preview.PreflightDecision != "development-staging-ready" ||
		preview.TransactionState != "transaction-ready" ||
		preview.ActivationState != "ready-for-runtime-commit" {
		t.Fatalf("unexpected activation status identity or state: %#v", preview)
	}
	if preview.Transaction.RequestType != "desktop-activation-transaction-preview" ||
		preview.Transaction.TransactionState != "transaction-ready" ||
		preview.Transaction.TransactionStepCount != 9 ||
		preview.Transaction.ReadyStepCount != 9 ||
		preview.Transaction.BlockedStepCount != 0 ||
		preview.Transaction.RollbackStepCount != 7 ||
		!preview.Transaction.TransactionReady ||
		preview.Transaction.TransactionCommitted {
		t.Fatalf("unexpected transaction summary: %#v", preview.Transaction)
	}
	if preview.Staging.RequestType != "desktop-activation-staging-preview" ||
		preview.Staging.StagingState != "staging-plan-ready" ||
		preview.Staging.PlannedFileCount != 5 ||
		preview.Staging.ActivatedEntryPointCount != 7 ||
		preview.Staging.ReceiptFileID != "desktop-activation-receipt" ||
		!preview.Staging.StagingPlanReady ||
		!preview.Staging.InstallerMayProceed ||
		preview.Staging.HostRootAllowed {
		t.Fatalf("unexpected staging summary: %#v", preview.Staging)
	}
	if preview.WriteGate.MethodName != "ActivateDesktopIntegration" ||
		preview.WriteGate.GateDecision != "disabled-for-preview" ||
		preview.WriteGate.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		preview.WriteGate.WriteMethodEnabled ||
		preview.WriteGate.DispatchEnabled {
		t.Fatalf("unexpected write gate summary: %#v", preview.WriteGate)
	}
	if preview.CommitGate.CommitState != "commit-gated" ||
		preview.CommitGate.CommitEnabled ||
		!preview.CommitGate.RequiresDigestMatch ||
		!preview.CommitGate.RequiresRollbackReceipt ||
		!preview.CommitGate.RequiresRuntimeOwner ||
		!preview.CommitGate.RequiresUserReview {
		t.Fatalf("unexpected commit gate summary: %#v", preview.CommitGate)
	}
	if preview.KDESurface.EntryPointCount != 7 ||
		!preview.KDESurface.MenuEntryReady ||
		!preview.KDESurface.TaskManagerReady ||
		!preview.KDESurface.FileManagerReady ||
		!preview.KDESurface.TrayStatusReady ||
		!preview.KDESurface.NotificationReady ||
		!preview.KDESurface.CompatibilityCenterReady ||
		!preview.KDESurface.UnifiedSettingsReady {
		t.Fatalf("unexpected KDE surface summary: %#v", preview.KDESurface)
	}
	if !preview.UserVisible || !preview.RuntimeOwned || !preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ActivationReady ||
		preview.ActivationCommitted ||
		preview.CommitEnabled ||
		!preview.InstallerMayProceed ||
		!preview.StagingPlanReady ||
		!preview.RollbackPlanned ||
		preview.RollbackAvailable {
		t.Fatalf("unexpected activation readiness flags: %#v", preview)
	}
	if preview.HostRootModified || preview.FileWritesPerformed ||
		preview.DesktopFilesWritten || preview.MIMEAppsWritten ||
		preview.KDEServiceCacheRefreshed || preview.LaunchEnabled ||
		preview.BackendLaunchEnabled || preview.ExecutionStarted ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed || preview.RawWindowsExecutableExposed ||
		preview.CompatibilityStorageExposed {
		t.Fatalf("activation status opened an unsafe gate: %#v", preview)
	}
	expectedSignals := []string{"transaction-plan", "staging-plan", "write-gate", "rollback-plan", "kde-surface"}
	if got := preview.StatusSignalIDs; !sameStrings(got, expectedSignals) {
		t.Fatalf("status signal ids = %#v, want %#v", got, expectedSignals)
	}
	expectedReasons := []string{"write-method-disabled", "digest-verification-pending", "rollback-receipt-not-written", "production-owner-not-active"}
	if got := preview.BlockedReasonIDs; !sameStrings(got, expectedReasons) {
		t.Fatalf("blocked reason ids = %#v, want %#v", got, expectedReasons)
	}
	expectedActions := []string{"show-compatibility-center-status", "show-user-review-card", "commit-desktop-activation", "launch-application", "refresh-kde-service-cache"}
	if got := preview.NextSafeActionIDs; !sameStrings(got, expectedActions) {
		t.Fatalf("next safe action ids = %#v, want %#v", got, expectedActions)
	}
	for _, action := range preview.NextSafeActions {
		if action.ID == "commit-desktop-activation" && action.Enabled {
			t.Fatalf("activation commit action must remain disabled: %#v", action)
		}
	}

	productionPreview, err := plan.DesktopActivationStatusPreview("production")
	if err != nil {
		t.Fatalf("production DesktopActivationStatusPreview returned error: %v", err)
	}
	if productionPreview.ActivationState != "blocked-before-runtime-commit" ||
		productionPreview.ActivationReady ||
		productionPreview.InstallerMayProceed ||
		productionPreview.Transaction.BlockedStepCount != productionPreview.Transaction.TransactionStepCount ||
		!containsString(productionPreview.BlockedReasonIDs, "transaction-not-ready") {
		t.Fatalf("unexpected production activation status: %#v", productionPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("desktop activation status exposed forbidden detail %q in %s", forbidden, serialized)
		}
	}
}
