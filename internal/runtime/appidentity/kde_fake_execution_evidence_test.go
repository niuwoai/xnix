package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestKDEFakeExecutionEvidencePersistsAndReadsBackControlledRecords(t *testing.T) {
	recipeRecord := Recipe{
		ID:                  "org.xnix.sample.notepad",
		Name:                "Sample Notepad",
		Version:             "1.0.0",
		Icon:                "accessories-text-editor",
		Mode:                "automatic",
		SupportedExtensions: []string{".txt"},
	}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()

	record, err := NewKDEFakeExecutionEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("NewKDEFakeExecutionEvidenceRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_fake_execution_evidence.v1" ||
		record.Application.ID != recipeRecord.ID ||
		record.Lifecycle.State != "staged" ||
		record.Lifecycle.OverallStatus != "not-ready" ||
		record.Execution.State != "blocked" ||
		record.Execution.ReviewDecision != "approved" ||
		record.Execution.GateCount != 6 ||
		record.Execution.PassedGateCount != 1 ||
		record.Session.State != "blocked" ||
		record.FanOut.SurfaceCount != 4 ||
		record.StateRootRecordCount != 3 ||
		!record.FakeExecutionRecorded ||
		!record.AllChecksPassed ||
		record.CheckCount != 6 ||
		record.PassedCheckCount != 6 {
		t.Fatalf("unexpected fake execution evidence: %+v", record)
	}
	if !record.StateRootWritesEnabled || record.StateRootWriteScope != "explicit-test-root-only" || record.StateRootPathExposed ||
		record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted ||
		record.RealPortalCallEnabled || record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified ||
		record.PrivilegedContainerRequired || record.BackendDetailsExposed {
		t.Fatalf("unsafe fake execution evidence gates: %+v", record)
	}
	for _, relativePath := range []string{record.Lifecycle.ReceiptRelativePath, record.Execution.ReceiptRelativePath, record.Session.ReceiptRelativePath} {
		if filepath.IsAbs(relativePath) {
			t.Fatalf("receipt path must be relative: %s", relativePath)
		}
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(relativePath))); err != nil {
			t.Fatalf("expected receipt %s: %v", relativePath, err)
		}
	}
	if record.Execution.ReceiptRelativePath == record.Session.ReceiptRelativePath || record.Execution.ReceiptSHA256 == "" || record.Session.ReceiptSHA256 == "" {
		t.Fatalf("execution and session receipts must be distinct and digest-backed: %+v", record)
	}

	repeated, err := NewKDEFakeExecutionEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"})
	if err != nil {
		t.Fatalf("repeated fake execution evidence should be safe: %v", err)
	}
	if repeated.Lifecycle.State != "staged" || !repeated.AllChecksPassed {
		t.Fatalf("repeated fake execution evidence lost readiness: %+v", repeated)
	}
}

func TestKDEFakeExecutionEvidenceRejectsUnsafeInputs(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	cases := []struct {
		name       string
		provenance Provenance
		options    KDEFakeExecutionEvidenceOptions
	}{
		{name: "missing mode", provenance: provenance, options: KDEFakeExecutionEvidenceOptions{StateRoot: t.TempDir()}},
		{name: "production mode", provenance: provenance, options: KDEFakeExecutionEvidenceOptions{StateRoot: t.TempDir(), Mode: "production"}},
		{name: "missing root", provenance: provenance, options: KDEFakeExecutionEvidenceOptions{Mode: "test-only"}},
		{name: "filesystem root", provenance: provenance, options: KDEFakeExecutionEvidenceOptions{StateRoot: string(os.PathSeparator), Mode: "test-only"}},
		{name: "unverified recipe", provenance: Provenance{Source: "registry", SignatureStatus: "development-only"}, options: KDEFakeExecutionEvidenceOptions{StateRoot: t.TempDir(), Mode: "test-only"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := NewKDEFakeExecutionEvidenceRecord(recipeRecord, tc.provenance, tc.options); err == nil {
				t.Fatal("expected unsafe fake execution evidence input to fail")
			}
		})
	}
}

func TestKDEFakeExecutionEvidenceRejectsManagedPathSymlink(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	root := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "environments")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	if _, err := NewKDEFakeExecutionEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: root, Mode: "test-only"}); err == nil {
		t.Fatal("expected managed path symlink to be rejected")
	}
	entries, err := os.ReadDir(outside)
	if err != nil {
		t.Fatalf("inspect outside directory: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("fake execution evidence wrote through managed path symlink: %+v", entries)
	}
}
