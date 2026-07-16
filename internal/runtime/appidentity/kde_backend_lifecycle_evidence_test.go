package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKDEBackendLifecycleEvidenceJoinsInventoryWithoutStartingProcess(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()
	record, err := NewKDEBackendLifecycleEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("NewKDEBackendLifecycleEvidenceRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_backend_lifecycle_evidence.v1" || record.Application.ID != recipeRecord.ID ||
		!record.Prerequisite.PortalReceiptCompleted || !record.Prerequisite.DiagnosticVerified || !record.Prerequisite.DiagnosticRedacted || !record.Prerequisite.SnapshotVerified || !record.Prerequisite.LifecycleReady || !record.Prerequisite.AllChecksPassed ||
		!record.BackendManager.InventoryPersisted || !record.BackendManager.InventoryReadBack || record.BackendManager.ManagedBackendCount != 3 || record.BackendManager.UserFacingProfileCount != 3 || !record.BackendManager.AllBackendsPlanned || !record.BackendManager.AllBackendProcessesStopped ||
		record.Lifecycle.State != "ready" || record.Execution.State != "blocked" || record.Execution.PassedGateCount != 4 || record.Session.State != "blocked" || record.FanOut.SurfaceCount != 4 ||
		!record.BackendStateJoined || record.CoreReceiptCount != 8 || !record.AllChecksPassed || record.CheckCount != 8 || record.PassedCheckCount != 8 {
		t.Fatalf("unexpected backend lifecycle evidence: %+v", record)
	}
	if filepath.IsAbs(record.BackendManager.ReceiptRelativePath) || len(record.BackendManager.ReceiptSHA256) != 64 {
		t.Fatalf("backend manager evidence must use a relative digest-backed receipt: %+v", record.BackendManager)
	}
	if !record.StateRootWritesEnabled || record.StateRootWriteScope != "explicit-test-root-only" || record.StateRootPathExposed || record.BackendKindsExposedToKDE ||
		record.BackendManager.KDEPolicyOwner || record.BackendManager.KDEVisible || record.BackendManager.BackendKindsExposedToKDE || record.BackendManager.InstallEnabled || record.BackendManager.DownloadEnabled || record.BackendManager.LaunchEnabled || record.BackendManager.ProcessStarted || record.BackendManager.VMProcessStarted ||
		record.BackendInstallEnabled || record.BackendDownloadEnabled || record.BackendLaunchEnabled || record.BackendProcessStarted || record.VMProcessStarted || record.RawCommandExposed || record.ProfilePathExposed ||
		record.RealPortalCallEnabled || record.SnapshotRestoreEnabled || record.DiagnosticExecutionEnabled || record.AIProviderCallEnabled || record.RepairExecutionEnabled || record.ExecutionApproved || record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified || record.PrivilegedContainerRequired || record.BackendDetailsExposed || record.SecretsExposed {
		t.Fatalf("unsafe backend lifecycle evidence gates: %+v", record)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(record.BackendManager.ReceiptRelativePath))); err != nil {
		t.Fatalf("expected backend manager receipt: %v", err)
	}
}

func TestKDEBackendLifecycleEvidenceRejectsUnsafeBoundary(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	if _, err := NewKDEBackendLifecycleEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: t.TempDir(), Mode: "production"}); err == nil {
		t.Fatal("expected production mode to be rejected")
	}
	root := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "backend-manager")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	if _, err := NewKDEBackendLifecycleEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"}); err == nil {
		t.Fatal("expected backend manager managed path symlink to be rejected")
	}
	entries, err := os.ReadDir(outside)
	if err != nil {
		t.Fatalf("inspect outside directory: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("backend lifecycle evidence wrote through managed path symlink: %+v", entries)
	}
}
