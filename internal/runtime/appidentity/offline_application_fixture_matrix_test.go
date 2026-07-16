package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
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
