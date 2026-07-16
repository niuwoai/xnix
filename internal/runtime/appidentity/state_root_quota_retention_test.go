package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStateRootQuotaRetentionPreviewMissingRootDoesNotCreateDirectories(t *testing.T) {
	missingRoot := filepath.Join(t.TempDir(), "missing-state-root")
	preview, err := NewStateRootQuotaRetentionPreview(StateRootQuotaRetentionOptions{
		ApplicationID: "org.example.storage",
		StateRoot:     missingRoot,
		QuotaBytes:    1024,
	})
	if err != nil {
		t.Fatalf("preview missing root: %v", err)
	}
	if preview.StateRootPresent {
		t.Fatalf("missing root must not be reported as present")
	}
	if preview.DirectoriesCreated || preview.FileDeletionEnabled || preview.HostRootModified {
		t.Fatalf("missing-root preview must not mutate state: %+v", preview)
	}
	if _, err := os.Stat(missingRoot); !os.IsNotExist(err) {
		t.Fatalf("preview must not create missing state root, stat err=%v", err)
	}
	if !containsString(preview.RetentionReasonCodes, "retain-missing-state-root") {
		t.Fatalf("missing root should retain missing-state-root reason: %+v", preview.RetentionReasonCodes)
	}
}

func TestStateRootQuotaRetentionPreviewUnderQuotaClassifiesKnownSections(t *testing.T) {
	root := t.TempDir()
	writeStateRootFixture(t, root, "snapshots/snap-001.json", `{"estimated_bytes":128}`)
	writeStateRootFixture(t, root, "diagnostics-ledger/runs/run-001.json", `{"estimated_bytes":64}`)
	writeStateRootFixture(t, root, "execution-ledger/transactions/request-001.json", `{"estimated_bytes":32}`)
	writeStateRootFixture(t, root, "portal-requests/handle-001.json", `{"estimated_bytes":16}`)
	writeStateRootFixture(t, root, "artifact-receipts/package-001.json", `{"estimated_bytes":8}`)
	writeStateRootFixture(t, root, "desktop-activation/receipt-001.json", `{"estimated_bytes":4}`)

	preview, err := NewStateRootQuotaRetentionPreview(StateRootQuotaRetentionOptions{
		ApplicationID: "org.example.storage",
		StateRoot:     root,
		QuotaBytes:    4096,
	})
	if err != nil {
		t.Fatalf("preview under quota: %v", err)
	}
	if !preview.StateRootPresent || preview.OverQuota {
		t.Fatalf("expected present under-quota state root: %+v", preview)
	}
	if preview.EstimatedBytes != 252 {
		t.Fatalf("expected estimated fixture bytes from metadata, got %d", preview.EstimatedBytes)
	}
	for _, id := range []string{"snapshots", "diagnostics", "execution-receipts", "portal-receipts", "artifact-receipts", "activation-receipts", "unknown-records"} {
		if !containsString(preview.SectionIDs, id) {
			t.Fatalf("missing section id %s in %+v", id, preview.SectionIDs)
		}
	}
	if !containsString(preview.RetentionReasonCodes, "retain-under-quota") {
		t.Fatalf("under-quota reason missing: %+v", preview.RetentionReasonCodes)
	}
	assertStateRootQuotaPreviewIsDryRun(t, preview, root)
}

func TestStateRootQuotaRetentionPreviewOverQuotaAndMalformedUnknownEvidence(t *testing.T) {
	root := t.TempDir()
	writeStateRootFixture(t, root, "diagnostics-ledger/runs/big-run.json", `{"estimated_bytes":900}`)
	writeStateRootFixture(t, root, "snapshots/bad.json", `{not-json`)
	writeStateRootFixture(t, root, "misc/unclassified.bin", `opaque`)

	preview, err := NewStateRootQuotaRetentionPreview(StateRootQuotaRetentionOptions{
		ApplicationID: "org.example.storage",
		StateRoot:     root,
		QuotaBytes:    128,
	})
	if err != nil {
		t.Fatalf("preview over quota: %v", err)
	}
	if !preview.OverQuota {
		t.Fatalf("expected over-quota preview")
	}
	if preview.Counts.MalformedRecords != 1 || preview.Counts.UnknownRecords != 1 {
		t.Fatalf("expected malformed and unknown counts, got %+v", preview.Counts)
	}
	for _, reason := range []string{"cleanup-candidate-over-quota", "retain-malformed-record", "retain-unknown-record-review"} {
		if !containsString(preview.RetentionReasonCodes, reason) {
			t.Fatalf("missing reason %s in %+v", reason, preview.RetentionReasonCodes)
		}
	}
	for _, reason := range []string{"malformed-record-review-required", "unknown-record-review-required"} {
		if !containsString(preview.BlockedCleanupReasons, reason) {
			t.Fatalf("missing blocked reason %s in %+v", reason, preview.BlockedCleanupReasons)
		}
	}
	assertStateRootQuotaPreviewIsDryRun(t, preview, root)
}

func TestStateRootQuotaRetentionPreviewActiveSessionBlocksExecutionReceipts(t *testing.T) {
	root := t.TempDir()
	writeStateRootFixture(t, root, "execution-ledger/transactions/request-001.json", `{"estimated_bytes":64}`)

	preview, err := NewStateRootQuotaRetentionPreview(StateRootQuotaRetentionOptions{
		ApplicationID: "org.example.storage",
		StateRoot:     root,
		QuotaBytes:    1024,
		ActiveSession: true,
	})
	if err != nil {
		t.Fatalf("preview active session: %v", err)
	}
	if !containsString(preview.RetentionReasonCodes, "retain-active-session") {
		t.Fatalf("missing active session retention reason: %+v", preview.RetentionReasonCodes)
	}
	if !containsString(preview.BlockedCleanupReasons, "active-session-blocked") {
		t.Fatalf("missing active session blocked reason: %+v", preview.BlockedCleanupReasons)
	}
	execution := findStateRootQuotaSection(t, preview, "execution-receipts")
	if execution.CleanupCandidate {
		t.Fatalf("execution receipts must not be cleanup candidates during active sessions: %+v", execution)
	}
}

func TestStateRootQuotaRetentionPreviewRetentionExemptEvidenceIsKept(t *testing.T) {
	root := t.TempDir()
	writeStateRootFixture(t, root, "snapshots/manual-keep.json", `{"estimated_bytes":512,"retention_exempt":true}`)

	preview, err := NewStateRootQuotaRetentionPreview(StateRootQuotaRetentionOptions{
		ApplicationID: "org.example.storage",
		StateRoot:     root,
		QuotaBytes:    64,
	})
	if err != nil {
		t.Fatalf("preview retention exempt: %v", err)
	}
	if !containsString(preview.RetentionReasonCodes, "retain-retention-exempt") {
		t.Fatalf("missing retention exempt reason: %+v", preview.RetentionReasonCodes)
	}
	if !containsString(preview.BlockedCleanupReasons, "retention-exempt") {
		t.Fatalf("missing retention exempt blocked reason: %+v", preview.BlockedCleanupReasons)
	}
	snapshots := findStateRootQuotaSection(t, preview, "snapshots")
	if snapshots.CleanupCandidate {
		t.Fatalf("retention-exempt snapshot must not be cleanup candidate: %+v", snapshots)
	}
}

func TestStateRootQuotaRetentionPreviewRejectsUnsafeInputs(t *testing.T) {
	root := t.TempDir()
	tests := []StateRootQuotaRetentionOptions{
		{ApplicationID: "not safe", StateRoot: root},
		{ApplicationID: "org.example.storage", StateRoot: ""},
		{ApplicationID: "org.example.storage", StateRoot: root, QuotaBytes: -1},
	}
	for _, tt := range tests {
		if _, err := NewStateRootQuotaRetentionPreview(tt); err == nil {
			t.Fatalf("expected unsafe options to fail: %+v", tt)
		}
	}
}

func writeStateRootFixture(t *testing.T, root string, relativePath string, data string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("create fixture directory: %v", err)
	}
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
}

func assertStateRootQuotaPreviewIsDryRun(t *testing.T, preview StateRootQuotaRetentionPreview, root string) {
	t.Helper()
	if preview.FileDeletionEnabled ||
		preview.DirectoriesCreated ||
		preview.LogTruncationEnabled ||
		preview.ReceiptsRewritten ||
		preview.SnapshotDeletionEnabled ||
		preview.StateRootPathExposed ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("preview must remain dry-run and local-only: %+v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("marshal preview: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("preview must not expose state root path: %s", encoded)
	}
	for _, section := range preview.Sections {
		if section.FileDeletionEnabled ||
			section.DirectoriesCreated ||
			section.LogTruncationEnabled ||
			section.ReceiptsRewritten ||
			section.SnapshotDeletionEnabled ||
			section.StateRootPathExposed ||
			section.HostRootModified {
			t.Fatalf("section must remain dry-run: %+v", section)
		}
	}
}

func findStateRootQuotaSection(t *testing.T, preview StateRootQuotaRetentionPreview, id string) StateRootQuotaRetentionSection {
	t.Helper()
	for _, section := range preview.Sections {
		if section.ID == id {
			return section
		}
	}
	t.Fatalf("missing section %s", id)
	return StateRootQuotaRetentionSection{}
}
