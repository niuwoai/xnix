package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/artifact"
	"xnix.local/xnix/internal/runtime/snapshot"
)

func TestOfflineApplicationFixtureMatrixPreviewCoversRepresentativeShapes(t *testing.T) {
	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}

	expectedShapes := []string{
		"document-editor",
		"game",
		"installer",
		"launcher",
		"network-heavy",
		"tray-heavy",
		"unsupported",
	}
	if preview.SchemaVersion != "xnix.runtime.offline_application_fixture_matrix.v1" ||
		preview.RequestType != "offline-application-fixture-matrix-preview" ||
		preview.RuntimeMethod != "GetOfflineApplicationFixtureMatrix" ||
		preview.ReadMethod != "GetOfflineApplicationFixtureMatrixPreview" {
		t.Fatalf("unexpected matrix metadata: %#v", preview)
	}
	if preview.Counts.Total != len(expectedShapes) ||
		preview.Counts.MissingFixture != 0 ||
		preview.Counts.Unsupported != 1 ||
		len(preview.Rows) != len(expectedShapes) {
		t.Fatalf("unexpected matrix counts: %#v", preview.Counts)
	}
	for _, shape := range expectedShapes {
		if !containsString(preview.ShapeIDs, shape) || !containsString(preview.RequiredShapeIDs, shape) {
			t.Fatalf("matrix missing expected shape %q: shape=%#v required=%#v", shape, preview.ShapeIDs, preview.RequiredShapeIDs)
		}
	}

	for _, row := range preview.Rows {
		if row.RecipeTrustState != "development-fixture-trusted" ||
			row.ArtifactReadiness != "missing-local-stage-receipt" ||
			row.SnapshotReadiness != "planned-receipt-required" ||
			row.DiagnosticReadiness != "metadata-only-ready" ||
			row.KDEJourneyCoverageState != "seven-entrypoints-covered" ||
			row.KDEJourneyEntryPointCount != 7 {
			t.Fatalf("unexpected row evidence for %s: %#v", row.ShapeID, row)
		}
		if row.NetworkFetchEnabled ||
			row.PackageManagerInvoked ||
			row.ArtifactStagingEnabled ||
			row.BackendProcessStarted ||
			row.LaunchEnabled ||
			row.ExecutionStarted ||
			row.RequestObjectCreated ||
			row.SettingsPersisted ||
			row.FileContentRead ||
			row.StateRootPathExposed ||
			row.RawExecutableExposed ||
			row.RawCommandExposed ||
			row.BackendDetailsExposed ||
			row.HostRootModified {
			t.Fatalf("row enabled unsafe side effect: %#v", row)
		}
	}
}

func TestOfflineApplicationFixtureMatrixPreviewReportsMissingFixtures(t *testing.T) {
	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:    []string{"document-editor", "game"},
		RuntimeRoot: "../../..",
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	if preview.MatrixStatus != "missing-fixtures" ||
		preview.Counts.Total != 2 ||
		preview.Counts.MissingFixture != 5 ||
		len(preview.MissingShapeIDs) != 5 {
		t.Fatalf("partial matrix did not report missing fixtures: %#v", preview)
	}
	if !containsString(preview.MissingShapeIDs, "unsupported") {
		t.Fatalf("missing shapes should include unsupported fixture: %#v", preview.MissingShapeIDs)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewRejectsUnknownShape(t *testing.T) {
	if _, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{ShapeIDs: []string{"unknown-shape"}}); err == nil {
		t.Fatalf("unknown fixture shape was accepted")
	}
}

func TestOfflineApplicationFixtureMatrixPreviewUnsupportedShapeIsBlocked(t *testing.T) {
	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:    []string{"unsupported"},
		RuntimeRoot: "../../..",
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if !row.UnsupportedShape ||
		row.MatrixState != "blocked-unsupported" ||
		!row.UserReviewRequired ||
		!containsString(row.BlockedReasons, "application shape requires explicit unsupported-state handling") {
		t.Fatalf("unsupported fixture shape not blocked: %#v", row)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewConsumesArtifactReceiptReadOnly(t *testing.T) {
	receiptRoot := t.TempDir()
	writeOfflineFixtureArtifactReceipt(t, receiptRoot, "org.xnix.fixture.document")

	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:            []string{"document-editor"},
		RuntimeRoot:         "../../..",
		ArtifactReceiptRoot: receiptRoot,
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if row.ArtifactReadiness != "local-fixture-ready" ||
		row.ArtifactStageReceipt == nil ||
		row.ArtifactStageReceipt.State != "ready" ||
		row.ArtifactStageReceipt.RelativePath != "artifact-ledger/receipts/org.xnix.fixture.document.json" ||
		row.ArtifactStageReceipt.RequiredArtifactCount != 1 ||
		!row.ArtifactStageReceipt.RequiredArtifactsStaged ||
		containsString(row.MissingEvidenceIDs, "artifact-stage-receipt") ||
		containsString(row.BlockedReasons, "local artifact staging receipt is missing") ||
		row.ArtifactStageReceipt.RootPathExposed ||
		row.ArtifactStageReceipt.NetworkFetchEnabled ||
		row.ArtifactStageReceipt.PackageManagerInvoked ||
		row.ArtifactStageReceipt.BackendLaunchEnabled ||
		row.ArtifactStageReceipt.HostRootModified {
		t.Fatalf("artifact receipt evidence was not consumed safely: %#v", row)
	}
	if row.MatrixState != "missing-evidence" ||
		!containsString(row.MissingEvidenceIDs, "snapshot-receipt") {
		t.Fatalf("artifact receipt should not bypass snapshot evidence: %#v", row)
	}
	encoded, err := json.Marshal(row)
	if err != nil {
		t.Fatalf("marshal row: %v", err)
	}
	if strings.Contains(string(encoded), receiptRoot) {
		t.Fatalf("artifact receipt evidence exposed receipt root: %s", string(encoded))
	}
}

func TestOfflineApplicationFixtureMatrixPreviewReportsInvalidArtifactReceipt(t *testing.T) {
	receiptRoot := t.TempDir()
	receiptDir := filepath.Join(receiptRoot, "artifact-ledger", "receipts")
	if err := os.MkdirAll(receiptDir, 0o700); err != nil {
		t.Fatalf("MkdirAll receipt dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(receiptDir, "org.xnix.fixture.document.json"), []byte("{"), 0o600); err != nil {
		t.Fatalf("WriteFile invalid receipt: %v", err)
	}

	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:            []string{"document-editor"},
		RuntimeRoot:         "../../..",
		ArtifactReceiptRoot: receiptRoot,
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if row.ArtifactReadiness != "invalid-local-stage-receipt" ||
		row.ArtifactStageReceipt == nil ||
		row.ArtifactStageReceipt.State != "invalid" ||
		!containsString(row.MissingEvidenceIDs, "artifact-stage-receipt") ||
		!containsString(row.BlockedReasons, "local artifact staging receipt is invalid") {
		t.Fatalf("invalid artifact receipt was not blocked: %#v", row)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewConsumesSnapshotBaselineReadOnly(t *testing.T) {
	snapshotRoot := writeOfflineFixtureSnapshotBaseline(t, "baseline-1")

	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:          []string{"document-editor"},
		RuntimeRoot:       "../../..",
		SnapshotStateRoot: snapshotRoot,
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if row.SnapshotReadiness != "baseline-receipt-ready" ||
		row.SnapshotBaselineReceipt == nil ||
		row.SnapshotBaselineReceipt.State != "ready" ||
		row.SnapshotBaselineReceipt.SnapshotID != "baseline-1" ||
		row.SnapshotBaselineReceipt.Reason != "before-repair" ||
		row.SnapshotBaselineReceipt.FileCount != 1 ||
		row.SnapshotBaselineReceipt.SnapshotCount != 1 ||
		!row.SnapshotBaselineReceipt.Verified ||
		containsString(row.MissingEvidenceIDs, "snapshot-receipt") ||
		containsString(row.BlockedReasons, "restore-point receipt is missing") ||
		row.SnapshotBaselineReceipt.StateRootPathExposed ||
		row.SnapshotBaselineReceipt.RestoreExecuted ||
		row.SnapshotBaselineReceipt.SnapshotCreated ||
		row.SnapshotBaselineReceipt.SnapshotDeleted ||
		row.SnapshotBaselineReceipt.FileContentRead ||
		row.SnapshotBaselineReceipt.BackendLaunchEnabled ||
		row.SnapshotBaselineReceipt.HostRootModified ||
		row.SnapshotBaselineReceipt.BackendDetailsExposed {
		t.Fatalf("snapshot baseline evidence was not consumed safely: %#v", row)
	}
	if row.MatrixState != "missing-evidence" ||
		!containsString(row.MissingEvidenceIDs, "artifact-stage-receipt") {
		t.Fatalf("snapshot baseline should not bypass artifact evidence: %#v", row)
	}
	encoded, err := json.Marshal(row)
	if err != nil {
		t.Fatalf("marshal row: %v", err)
	}
	if strings.Contains(string(encoded), snapshotRoot) {
		t.Fatalf("snapshot baseline evidence exposed state root: %s", string(encoded))
	}
}

func TestOfflineApplicationFixtureMatrixPreviewCoversReadyFixtureWithArtifactAndSnapshotEvidence(t *testing.T) {
	artifactRoot := t.TempDir()
	writeOfflineFixtureArtifactReceipt(t, artifactRoot, "org.xnix.fixture.document")
	snapshotRoot := writeOfflineFixtureSnapshotBaseline(t, "baseline-2")

	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:            []string{"document-editor"},
		RuntimeRoot:         "../../..",
		ArtifactReceiptRoot: artifactRoot,
		SnapshotStateRoot:   snapshotRoot,
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if row.MatrixState != "covered-review-only" ||
		row.UserReviewRequired ||
		len(row.MissingEvidenceIDs) != 0 ||
		len(row.BlockedReasons) != 0 ||
		row.ArtifactReadiness != "local-fixture-ready" ||
		row.SnapshotReadiness != "baseline-receipt-ready" ||
		preview.Counts.Covered != 1 ||
		preview.Counts.MissingEvidence != 0 ||
		preview.Counts.NeedsReview != 0 {
		t.Fatalf("ready fixture was not covered review-only: row=%#v counts=%#v status=%s", row, preview.Counts, preview.MatrixStatus)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewReportsInvalidSnapshotBaseline(t *testing.T) {
	snapshotRoot := writeOfflineFixtureSnapshotBaseline(t, "baseline-corrupt")
	if err := os.WriteFile(filepath.Join(snapshotRoot, ".xnix-snapshots", "objects", offlineFixtureSnapshotObjectDigest(t, snapshotRoot, "baseline-corrupt")), []byte("tampered"), 0o600); err != nil {
		t.Fatalf("tamper snapshot object: %v", err)
	}

	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:          []string{"document-editor"},
		RuntimeRoot:       "../../..",
		SnapshotStateRoot: snapshotRoot,
	})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	row := preview.Rows[0]
	if row.SnapshotReadiness != "invalid-baseline-receipt" ||
		row.SnapshotBaselineReceipt == nil ||
		row.SnapshotBaselineReceipt.State != "invalid" ||
		!containsString(row.MissingEvidenceIDs, "snapshot-receipt") ||
		!containsString(row.BlockedReasons, "restore-point receipt is invalid") {
		t.Fatalf("invalid snapshot baseline was not blocked: %#v", row)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewNeverEnablesSideEffects(t *testing.T) {
	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	if preview.NetworkFetchEnabled ||
		preview.PackageManagerInvoked ||
		preview.ArtifactStagingEnabled ||
		preview.BackendLaunchEnabled ||
		preview.DockerRequired ||
		preview.QEMURequired ||
		preview.RequestObjectsCreated ||
		preview.SettingsPersisted ||
		preview.FileContentRead ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("matrix enabled unsafe side effect: %#v", preview)
	}
}

func writeOfflineFixtureSnapshotBaseline(t *testing.T, snapshotID string) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "runtime-state.json"), []byte(`{"state":"ready"}`), 0o600); err != nil {
		t.Fatalf("WriteFile snapshot source: %v", err)
	}
	store, err := snapshot.New(root)
	if err != nil {
		t.Fatalf("snapshot.New returned error: %v", err)
	}
	if _, err := store.Create(snapshotID, "before-repair"); err != nil {
		t.Fatalf("snapshot Create returned error: %v", err)
	}
	return root
}

func offlineFixtureSnapshotObjectDigest(t *testing.T, root string, snapshotID string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, ".xnix-snapshots", "manifests", snapshotID+".json"))
	if err != nil {
		t.Fatalf("ReadFile snapshot manifest: %v", err)
	}
	var manifest snapshot.Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatalf("Unmarshal snapshot manifest: %v", err)
	}
	if len(manifest.Files) == 0 {
		t.Fatalf("snapshot manifest has no files: %#v", manifest)
	}
	return manifest.Files[0].Digest
}

func writeOfflineFixtureArtifactReceipt(t *testing.T, receiptRoot string, applicationID string) {
	t.Helper()
	fixtureRoot := filepath.Join(t.TempDir(), "fixtures")
	if err := os.MkdirAll(fixtureRoot, 0o700); err != nil {
		t.Fatalf("MkdirAll fixture root: %v", err)
	}
	payload := []byte("offline fixture artifact for " + applicationID)
	sum := sha256.Sum256(payload)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(fixtureRoot, digest), payload, 0o600); err != nil {
		t.Fatalf("WriteFile fixture artifact: %v", err)
	}
	_, err := artifact.StageFromFixture(artifact.StageRequest{
		CacheRoot:  receiptRoot,
		FixtureDir: fixtureRoot,
		Manifest: artifact.Manifest{
			ApplicationID: applicationID,
			Groups: []artifact.Group{
				{
					ID: "runtime-launch-metadata",
					Refs: []artifact.Ref{
						{ID: "launch.json", Kind: "metadata", SHA256: digest, Size: int64(len(payload)), Required: true},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("StageFromFixture returned error: %v", err)
	}
}

func TestOfflineApplicationFixtureMatrixPreviewIsDesktopSafe(t *testing.T) {
	preview, err := NewOfflineApplicationFixtureMatrixPreview(OfflineApplicationFixtureMatrixOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("NewOfflineApplicationFixtureMatrixPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("marshal preview: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, key := range []string{
		"network_fetch_enabled",
		"package_manager_invoked",
		"artifact_staging_enabled",
		"backend_process_started",
		"host_root_modified",
		"raw_command_exposed",
	} {
		if !strings.Contains(text, `"`+key+`":false`) {
			t.Fatalf("offline fixture matrix JSON missing disabled key %q: %s", key, string(encoded))
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("offline fixture matrix exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
}
