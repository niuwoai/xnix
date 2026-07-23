package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationTransactionPreviewPlansCommitAndRollbackWithoutWriting(t *testing.T) {
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

	preview, err := plan.DesktopActivationTransactionPreview("development")
	if err != nil {
		t.Fatalf("DesktopActivationTransactionPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_transaction.v1" ||
		preview.RequestType != "desktop-activation-transaction-preview" ||
		preview.TransactionType != "kde-desktop-activation-transaction" ||
		preview.ReadMethod != "GetDesktopActivationTransactionPreview" ||
		preview.WriteMethod != "ActivateDesktopIntegration" {
		t.Fatalf("unexpected transaction schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.InstallMode != "development" ||
		preview.PreflightDecision != "development-staging-ready" ||
		preview.TransactionState != "transaction-ready" {
		t.Fatalf("unexpected transaction identity or state: %#v", preview)
	}
	if preview.Staging.RequestType != "desktop-activation-staging-preview" ||
		preview.Staging.StagingState != "staging-plan-ready" ||
		preview.Staging.PlannedFileCount != 6 ||
		preview.Staging.ReceiptFileID != "desktop-activation-receipt" ||
		preview.Staging.ActivatedEntryPointCount != 7 ||
		!preview.Staging.InstallerMayProceed ||
		!preview.Staging.StagingPlanReady ||
		preview.Staging.HostRootAllowed {
		t.Fatalf("unexpected staging summary: %#v", preview.Staging)
	}
	if got, want := preview.Staging.PlannedFileIDs, []string{"desktop-entry", "dolphin-service-menu", "mimeapps-list", "desktop-integration-manifest", "managed-launcher-artifact", "desktop-activation-receipt"}; !sameStrings(got, want) {
		t.Fatalf("staging planned file ids = %#v, want %#v", got, want)
	}
	if preview.WriteGate.MethodName != "ActivateDesktopIntegration" ||
		preview.WriteGate.GateDecision != "disabled-for-preview" ||
		preview.WriteGate.WriteMethodEnabled ||
		preview.WriteGate.DispatchEnabled ||
		preview.WriteGate.RequestObjectCreated {
		t.Fatalf("unexpected write gate: %#v", preview.WriteGate)
	}
	if preview.ReceiptEvidence.EvidenceType != "desktop-activation-transaction-receipt-evidence" ||
		preview.ReceiptEvidence.EvidenceState != "planned-runtime-gated" ||
		preview.ReceiptEvidence.ReceiptSchemaVersion != "xnix.runtime.desktop_activation_receipt.v1" ||
		preview.ReceiptEvidence.CommitReceiptType != "desktop-activation-receipt" ||
		preview.ReceiptEvidence.RollbackReceiptType != "desktop-activation-rollback-receipt" ||
		preview.ReceiptEvidence.ReceiptRelativePath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		preview.ReceiptEvidence.RequiredFileCount != 6 ||
		!sameStrings(preview.ReceiptEvidence.RequiredFileIDs, []string{"desktop-entry", "dolphin-service-menu", "mimeapps-list", "desktop-integration-manifest", "managed-launcher-artifact", "desktop-activation-receipt"}) ||
		!preview.ReceiptEvidence.InstalledFileDigestRequired ||
		!preview.ReceiptEvidence.RollbackDigestRequired ||
		!preview.ReceiptEvidence.CommitReceiptRequired ||
		!preview.ReceiptEvidence.CommitReceiptPlanned ||
		preview.ReceiptEvidence.CommitReceiptWritten ||
		!preview.ReceiptEvidence.RollbackReceiptRequired ||
		!preview.ReceiptEvidence.RollbackReceiptPlanned ||
		preview.ReceiptEvidence.RollbackReceiptWritten ||
		preview.ReceiptEvidence.DigestGateReady ||
		preview.ReceiptEvidence.CommitAvailable ||
		preview.ReceiptEvidence.RollbackAvailable ||
		!preview.ReceiptEvidence.RuntimeOwned ||
		preview.ReceiptEvidence.KDEPolicyOwner ||
		preview.ReceiptEvidence.RootPathExposed ||
		preview.ReceiptEvidence.TargetRootPathExposed ||
		preview.ReceiptEvidence.HostRootModified ||
		preview.ReceiptEvidence.FileWritesPerformed ||
		preview.ReceiptEvidence.BackendDetailsExposed {
		t.Fatalf("unexpected transaction receipt evidence: %#v", preview.ReceiptEvidence)
	}
	expectedStepIDs := []string{
		"validate-preflight",
		"prepare-staging-root",
		"verify-staged-file-digests",
		"install-desktop-entry",
		"install-dolphin-service-menu",
		"merge-mimeapps-associations",
		"install-desktop-integration-manifest",
		"write-rollback-receipt",
		"refresh-kde-service-cache",
	}
	if got := preview.TransactionStepIDs; !sameStrings(got, expectedStepIDs) {
		t.Fatalf("transaction step ids = %#v, want %#v", got, expectedStepIDs)
	}
	if preview.TransactionStepCount != len(expectedStepIDs) ||
		preview.ReadyStepCount != len(expectedStepIDs) ||
		preview.BlockedStepCount != 0 {
		t.Fatalf("unexpected transaction step counts: %#v", preview)
	}
	for _, step := range preview.TransactionSteps {
		if step.Status != "ready" {
			t.Fatalf("ready transaction contains non-ready step: %#v", step)
		}
	}
	expectedRollbackIDs := []string{
		"load-activation-receipt",
		"verify-installed-file-digests",
		"remove-desktop-entry",
		"remove-dolphin-service-menu",
		"restore-mimeapps-associations",
		"remove-desktop-integration-manifest",
		"mark-receipt-rolled-back",
	}
	if got := preview.RollbackStepIDs; !sameStrings(got, expectedRollbackIDs) {
		t.Fatalf("rollback step ids = %#v, want %#v", got, expectedRollbackIDs)
	}
	if preview.RollbackStepCount != len(expectedRollbackIDs) {
		t.Fatalf("unexpected rollback step count: %#v", preview)
	}
	if got, want := preview.ActivationCommandPreview, []string{"xnix-activate-desktop-integration", "--plan-source", "runtime-go", "--target", "staging-root"}; !sameStrings(got, want) {
		t.Fatalf("activation command preview = %#v, want %#v", got, want)
	}
	if got, want := preview.RollbackCommandPreview, []string{"xnix-rollback-desktop-integration", "--receipt-source", "runtime-go"}; !sameStrings(got, want) {
		t.Fatalf("rollback command preview = %#v, want %#v", got, want)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.InstallerMayProceed ||
		!preview.TransactionPlanCreated ||
		!preview.TransactionReady ||
		preview.TransactionCommitted ||
		!preview.StagingPlanReady ||
		!preview.StagedFileDigestsRequired ||
		preview.StagedFileDigestsVerified ||
		!preview.StagingRootRequired ||
		preview.StagingRootPathExposed ||
		preview.TargetRootPathExposed ||
		preview.HostRootAllowed {
		t.Fatalf("unexpected transaction readiness flags: %#v", preview)
	}
	if preview.FileWritesPerformed || preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten || preview.ManifestWritten ||
		preview.ReceiptWritten || !preview.RollbackReceiptRequired ||
		!preview.RollbackReceiptPlanned || preview.RollbackReceiptWritten ||
		preview.RollbackAvailable || preview.KDEServiceCacheRefreshed ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.TaskManagerEntryActive || preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled || preview.LaunchEnabled ||
		preview.BackendLaunchEnabled || preview.ExecutionStarted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.PrivilegedContainerRequired || preview.BackendDetailsExposed ||
		preview.RawWindowsExecutableExposed || preview.CompatibilityStorageExposed {
		t.Fatalf("desktop activation transaction opened an unsafe gate: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "commit desktop activation from transaction preview") ||
		!containsString(preview.BlockedActions, "write activation files before digest verification") ||
		!containsString(preview.BlockedActions, "refresh KDE service cache from transaction preview") {
		t.Fatalf("missing blocked actions: %#v", preview.BlockedActions)
	}

	productionPreview, err := plan.DesktopActivationTransactionPreview("production")
	if err != nil {
		t.Fatalf("production DesktopActivationTransactionPreview returned error: %v", err)
	}
	if productionPreview.TransactionState != "transaction-blocked" ||
		productionPreview.TransactionReady ||
		productionPreview.InstallerMayProceed ||
		productionPreview.ReadyStepCount != 0 ||
		productionPreview.BlockedStepCount != productionPreview.TransactionStepCount {
		t.Fatalf("unexpected production transaction state: %#v", productionPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("desktop activation transaction exposed forbidden detail %q in %s", forbidden, serialized)
		}
	}
}
