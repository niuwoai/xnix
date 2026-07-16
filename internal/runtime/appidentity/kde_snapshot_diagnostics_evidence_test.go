package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKDESnapshotDiagnosticsEvidenceConvergesControlledState(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()

	record, err := NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("NewKDESnapshotDiagnosticsEvidenceRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_snapshot_diagnostics_evidence.v1" ||
		record.Application.ID != recipeRecord.ID ||
		record.Portal.State != "completed" ||
		record.Portal.PermissionState != "granted" ||
		record.Diagnostic.RunID != kdeStabilityRunID ||
		record.Diagnostic.Overall != "pass" ||
		record.Diagnostic.SignalCount != 3 ||
		record.Snapshot.SnapshotID != kdeStabilitySnapshotID ||
		!record.Snapshot.Created || record.Snapshot.Reused ||
		!record.Snapshot.Verified || !record.Snapshot.BaselinePresent ||
		record.Lifecycle.State != "ready" || len(record.Lifecycle.PendingGates) != 0 ||
		record.Execution.State != "blocked" || record.Execution.PassedGateCount != 4 ||
		fakeExecutionGateStatus(record.Execution.Gates, "environment-ready") != "pass" ||
		fakeExecutionGateStatus(record.Execution.Gates, "snapshot-baseline") != "pass" ||
		fakeExecutionGateStatus(record.Execution.Gates, "portal-permission") != "pass" ||
		fakeExecutionGateStatus(record.Execution.Gates, "recipe-trust") != "pending" ||
		fakeExecutionGateStatus(record.Execution.Gates, "runtime-write-gate") != "blocked" ||
		record.Session.State != "blocked" || record.FanOut.SurfaceCount != 4 ||
		record.CoreReceiptCount != 7 || !record.AllChecksPassed ||
		record.CheckCount != 8 || record.PassedCheckCount != 8 {
		t.Fatalf("unexpected snapshot diagnostics evidence: %+v", record)
	}
	if filepath.IsAbs(record.Portal.ReceiptRelativePath) || filepath.IsAbs(record.Diagnostic.ReceiptRelativePath) || filepath.IsAbs(record.Snapshot.ManifestRelativePath) ||
		len(record.Portal.ReceiptSHA256) != 64 || len(record.Diagnostic.ReceiptSHA256) != 64 || len(record.Snapshot.ContentHash) != 64 {
		t.Fatalf("evidence paths and digests must remain KDE-safe: %+v", record)
	}
	if !record.StateRootWritesEnabled || record.StateRootWriteScope != "explicit-test-root-only" || record.StateRootPathExposed ||
		!record.SnapshotCreationEnabled || record.SnapshotRestoreEnabled || record.SnapshotDeletionEnabled ||
		record.DiagnosticExecutionEnabled || record.AIProviderCallEnabled || record.RepairExecutionEnabled ||
		record.RealPortalCallEnabled || record.HostPermissionChanged || record.ExecutionApproved ||
		record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted ||
		record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified ||
		record.PrivilegedContainerRequired || record.BackendDetailsExposed || record.FileContentsExposed {
		t.Fatalf("unsafe snapshot diagnostics gates: %+v", record)
	}
	for _, relativePath := range []string{record.Portal.ReceiptRelativePath, record.Diagnostic.ReceiptRelativePath, record.Snapshot.ManifestRelativePath, record.Lifecycle.ReceiptRelativePath, record.Execution.ReceiptRelativePath, record.Session.ReceiptRelativePath} {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(relativePath))); err != nil {
			t.Fatalf("expected controlled receipt %s: %v", relativePath, err)
		}
	}

	repeated, err := NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("repeated snapshot diagnostics evidence should reuse baseline: %v", err)
	}
	if repeated.Snapshot.Created || !repeated.Snapshot.Reused || repeated.Snapshot.ContentHash != record.Snapshot.ContentHash || !repeated.AllChecksPassed {
		t.Fatalf("repeated evidence did not reuse verified baseline: %+v", repeated)
	}
}

func TestKDESnapshotDiagnosticsEvidenceRejectsUnsafeBoundaries(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	if _, err := NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: t.TempDir(), Mode: "production"}); err == nil {
		t.Fatal("expected production mode to be rejected")
	}

	for _, managedPath := range []string{".xnix-snapshots", "diagnostics-ledger", "test-results"} {
		t.Run(managedPath, func(t *testing.T) {
			root := t.TempDir()
			outside := t.TempDir()
			if err := os.Symlink(outside, filepath.Join(root, managedPath)); err != nil {
				t.Skipf("symlinks unavailable: %v", err)
			}
			if _, err := NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"}); err == nil {
				t.Fatalf("expected managed path symlink %s to be rejected", managedPath)
			}
			entries, err := os.ReadDir(outside)
			if err != nil {
				t.Fatalf("inspect outside directory: %v", err)
			}
			if len(entries) != 0 {
				t.Fatalf("evidence wrote through managed path symlink %s: %+v", managedPath, entries)
			}
		})
	}
}
