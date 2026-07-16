package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKDEFakePortalEvidenceChangesOnlyPortalGate(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()

	record, err := NewKDEFakePortalEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("NewKDEFakePortalEvidenceRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_fake_portal_evidence.v1" ||
		record.Application.ID != recipeRecord.ID ||
		record.BeforePortal.PortalGateStatus != "pending" ||
		record.Portal.CreatedState != "pending-user-mediation" ||
		record.Portal.ResolvedState != "granted" ||
		record.Portal.CompletedState != "completed" ||
		record.Portal.PermissionState != "granted" ||
		record.Lifecycle.State != "staged" ||
		!containsString(record.Lifecycle.SatisfiedGates, "portal-policy-review") ||
		!containsString(record.Lifecycle.PendingGates, "snapshot-baseline") ||
		record.Execution.State != "blocked" ||
		record.Execution.PassedGateCount != 2 ||
		fakeExecutionGateStatus(record.Execution.Gates, "portal-permission") != "pass" ||
		record.Session.State != "blocked" ||
		record.FanOut.SurfaceCount != 4 ||
		record.StateRootRecordCount != 4 ||
		!record.PortalEvidenceRecorded ||
		!record.PortalGateChangedOnly ||
		!record.AllChecksPassed ||
		record.CheckCount != 7 ||
		record.PassedCheckCount != 7 {
		t.Fatalf("unexpected fake Portal evidence: %+v", record)
	}
	if filepath.IsAbs(record.Portal.ReceiptRelativePath) || len(record.Portal.ReceiptSHA256) != 64 {
		t.Fatalf("fake Portal receipt must be relative and digest-backed: %+v", record.Portal)
	}
	if !record.StateRootWritesEnabled || record.StateRootWriteScope != "explicit-test-root-only" || record.StateRootPathExposed ||
		record.RealPortalCallEnabled || record.HostPermissionChanged || record.ExecutionApproved ||
		record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted ||
		record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified || record.PrivilegedContainerRequired || record.BackendDetailsExposed {
		t.Fatalf("unsafe fake Portal evidence gates: %+v", record)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(record.Portal.ReceiptRelativePath))); err != nil {
		t.Fatalf("expected fake Portal receipt: %v", err)
	}

	repeated, err := NewKDEFakePortalEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("repeated fake Portal evidence should reuse completed receipt: %v", err)
	}
	if !repeated.Portal.ReceiptReused || repeated.Portal.HandleToken != record.Portal.HandleToken || !repeated.AllChecksPassed {
		t.Fatalf("repeated fake Portal evidence did not reuse safe receipt: %+v", repeated)
	}
}

func TestKDEFakePortalEvidenceRejectsManagedPathSymlink(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "portal-requests")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	if _, err := NewKDEFakePortalEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"}); err == nil {
		t.Fatal("expected fake Portal managed path symlink to be rejected")
	}
	entries, err := os.ReadDir(outside)
	if err != nil {
		t.Fatalf("inspect outside directory: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("fake Portal evidence wrote through managed path symlink: %+v", entries)
	}
}
