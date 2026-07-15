package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func TestDesktopActivationStatusPreviewConsumesActivationReceipt(t *testing.T) {
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
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := plan.DesktopActivationStatusPreviewWithReceipt(root, "development")
	if err != nil {
		t.Fatalf("DesktopActivationStatusPreviewWithReceipt returned error: %v", err)
	}

	if preview.Source != "desktop-activation-transaction-preview+desktop-activation-receipt" ||
		preview.ActivationState != "receipt-backed-runtime-gated" ||
		!preview.ReceiptBacked ||
		!preview.RollbackAvailable ||
		preview.ActivationCommitted ||
		preview.CommitEnabled ||
		preview.LaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected receipt-backed activation status: %#v", preview)
	}
	evidence := preview.ReceiptEvidence
	if evidence.EvidenceState != "receipt-backed" ||
		evidence.SchemaVersion != "xnix.runtime.desktop_activation_receipt.v1" ||
		evidence.ReceiptType != "desktop-activation-receipt" ||
		evidence.ReceiptRelativePath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		evidence.ApplicationID != "org.example.ledger" ||
		evidence.InstalledFileCount != 2 ||
		!sameStrings(evidence.InstalledFileIDs, []string{"desktop-entry", "desktop-activation-receipt"}) ||
		!evidence.DigestGateReady ||
		!evidence.RollbackReceiptReady ||
		!evidence.RuntimeOwned ||
		evidence.RootPathExposed ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		!evidence.SafeForKDE {
		t.Fatalf("unexpected receipt evidence: %#v", evidence)
	}
	if !containsString(preview.StatusSignalIDs, "activation-receipt") {
		t.Fatalf("receipt-backed status must expose activation-receipt signal: %#v", preview.StatusSignalIDs)
	}
	if containsString(preview.BlockedReasonIDs, "rollback-receipt-not-written") {
		t.Fatalf("receipt-backed status must not claim the receipt is unwritten: %#v", preview.BlockedReasonIDs)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	if strings.Contains(serialized, strings.ToLower(root)) {
		t.Fatalf("receipt-backed status exposed the staging root: %s", serialized)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("receipt-backed activation status exposed forbidden detail %q in %s", forbidden, serialized)
		}
	}

	if _, err := plan.DesktopActivationStatusPreviewWithReceipt(t.TempDir(), "development"); err == nil ||
		!strings.Contains(err.Error(), "read desktop activation receipt") {
		t.Fatalf("missing receipt must fail closed, got %v", err)
	}
	mismatchRoot := t.TempDir()
	writeMismatchedActivationReceipt(t, mismatchRoot)
	if _, err := plan.DesktopActivationStatusPreviewWithReceipt(mismatchRoot, "development"); err == nil ||
		!strings.Contains(err.Error(), "application mismatch") {
		t.Fatalf("mismatched receipt must fail closed, got %v", err)
	}
}

func writeMismatchedActivationReceipt(t *testing.T, root string) {
	t.Helper()
	writeActivationReceipt(t, root, "org.example.ledger")
	path := filepath.Join(root, "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	data = []byte(strings.Replace(string(data), `"application_id": "org.example.ledger"`, `"application_id": "org.example.other"`, 1))
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile mismatched receipt returned error: %v", err)
	}
}

func writeActivationReceipt(t *testing.T, root string, applicationID string) {
	t.Helper()
	relativePath := filepath.Join("usr/share/xnix/compatibility/activation-receipts", applicationID+".json")
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	receipt := `{
  "schema_version": "xnix.runtime.desktop_activation_receipt.v1",
  "receipt_type": "desktop-activation-receipt",
  "application_id": "` + applicationID + `",
  "installed": [
    {
      "id": "desktop-entry",
      "relative_path": "usr/share/applications/xnix-org.example.ledger.desktop",
      "sha256": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
      "written": true,
      "host_root_modified": false,
      "backend_details_exposed": false
    },
    {
      "id": "desktop-activation-receipt",
      "relative_path": "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json",
      "sha256": "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
      "written": true,
      "host_root_modified": false,
      "backend_details_exposed": false
    }
  ],
  "rollback": {
    "command": "xnix-rollback-desktop-integration",
    "requires_matching_sha256": true,
    "host_root_modified": false
  },
  "safety": {
    "runtime_owned": true,
    "host_root_modified": false,
    "backend_details_exposed": false
  }
}
`
	if err := os.WriteFile(path, []byte(receipt), 0o600); err != nil {
		t.Fatalf("WriteFile receipt returned error: %v", err)
	}
}
