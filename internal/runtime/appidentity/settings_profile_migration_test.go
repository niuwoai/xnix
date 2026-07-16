package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSettingsProfileMigrationOldSchemaRequiresReview(t *testing.T) {
	plan := newSettingsProfileMigrationPlan(t)
	preview, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "xnix.runtime.settings.v0",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
	})
	if err != nil {
		t.Fatalf("SettingsProfileMigrationPreview: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.settings_profile_migration.v1" ||
		preview.RequestType != "settings-profile-migration-preview" ||
		preview.RuntimeMethod != "GetCompatibilitySettingsProfileMigration" ||
		preview.ReadMethod != "GetCompatibilitySettingsProfileMigrationPreview" ||
		preview.SchemaState != "old-schema" ||
		!preview.MigrationRequired ||
		!preview.UserReviewRequired {
		t.Fatalf("unexpected migration schema: %#v", preview)
	}
	if preview.RowCount != 8 ||
		!sameStrings(preview.RowIDs, []string{"run-mode.mode", "run-mode.preference", "resource-access.documents", "resource-access.downloads", "devices.camera", "network.network", "snapshots.snapshots", "diagnostics.privacy"}) {
		t.Fatalf("unexpected migration rows: %#v", preview.RowIDs)
	}
	if preview.Counts.ReviewRequired != 6 || preview.Counts.Defaulted != 2 {
		t.Fatalf("unexpected counts: %#v", preview.Counts)
	}
	documents := settingsProfileMigrationRowByID(preview.Rows, "resource-access.documents")
	if documents.MigrationState != "review-required" ||
		documents.CurrentValue != "not-present" ||
		documents.TargetDefaultValue != "ask" ||
		!documents.UserReviewRequired ||
		documents.NextSafeReadOnlyCheck != "portal-permission-renewal-preview" {
		t.Fatalf("unexpected documents row: %#v", documents)
	}
	if len(preview.RollbackNotes) == 0 || !strings.Contains(preview.RollbackNotes[0], "Do not persist") {
		t.Fatalf("expected rollback notes: %#v", preview.RollbackNotes)
	}
	assertSettingsProfileMigrationNoSideEffects(t, preview)
	assertSettingsProfileMigrationSafe(t, preview)
}

func TestSettingsProfileMigrationCurrentSchemaIsUnchanged(t *testing.T) {
	plan := newSettingsProfileMigrationPlan(t)
	preview, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "xnix.runtime.settings.v1",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
	})
	if err != nil {
		t.Fatalf("SettingsProfileMigrationPreview: %v", err)
	}
	if preview.SchemaState != "current-schema" ||
		preview.MigrationRequired ||
		preview.Counts.Unchanged != 8 ||
		preview.Counts.ReviewRequired != 0 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("current schema should be unchanged: %#v", preview)
	}
	runMode := settingsProfileMigrationRowByID(preview.Rows, "run-mode.mode")
	if runMode.CurrentValue != "automatic" || runMode.TargetDefaultValue != "automatic" {
		t.Fatalf("current schema must keep current values: %#v", runMode)
	}
	assertSettingsProfileMigrationNoSideEffects(t, preview)
}

func TestSettingsProfileMigrationFutureSchemaBlocksMigration(t *testing.T) {
	plan := newSettingsProfileMigrationPlan(t)
	preview, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "xnix.runtime.settings.v2",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
	})
	if err != nil {
		t.Fatalf("SettingsProfileMigrationPreview: %v", err)
	}
	if preview.SchemaState != "future-schema" ||
		preview.CurrentSchemaSupported ||
		preview.Counts.FutureSchema != 8 ||
		!preview.UserReviewRequired {
		t.Fatalf("future schema must be blocked for review: %#v", preview)
	}
	if !settingsProfileMigrationRowByID(preview.Rows, "diagnostics.privacy").Blocked {
		t.Fatalf("future schema rows must be blocked: %#v", preview.Rows)
	}
	if len(preview.RollbackNotes) != 3 {
		t.Fatalf("future schema needs extra rollback note: %#v", preview.RollbackNotes)
	}
	assertSettingsProfileMigrationNoSideEffects(t, preview)
}

func TestSettingsProfileMigrationBlockedAndInvalidSetting(t *testing.T) {
	plan := newSettingsProfileMigrationPlan(t)
	preview, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "xnix.runtime.settings.v0",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
		BlockedSettingIDs: []string{"devices.camera"},
	})
	if err != nil {
		t.Fatalf("SettingsProfileMigrationPreview: %v", err)
	}
	camera := settingsProfileMigrationRowByID(preview.Rows, "devices.camera")
	if camera.MigrationState != "blocked" ||
		!camera.Blocked ||
		len(camera.BlockedReasons) != 1 ||
		preview.Counts.Blocked != 1 ||
		!sameStrings(preview.BlockedSettingIDs, []string{"devices.camera"}) {
		t.Fatalf("blocked setting not represented: %#v / %#v", camera, preview)
	}
	if _, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "xnix.runtime.settings.v0",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
		BlockedSettingIDs: []string{"unknown.setting"},
	}); err == nil {
		t.Fatalf("SettingsProfileMigrationPreview accepted an invalid setting id")
	}
	if _, err := plan.SettingsProfileMigrationPreview(SettingsProfileMigrationOptions{
		FromSchemaVersion: "bad-schema",
		ToSchemaVersion:   "xnix.runtime.settings.v1",
	}); err == nil {
		t.Fatalf("SettingsProfileMigrationPreview accepted an invalid schema")
	}
}

func newSettingsProfileMigrationPlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan: %v", err)
	}
	return plan
}

func settingsProfileMigrationRowByID(rows []SettingsProfileMigrationRow, id string) SettingsProfileMigrationRow {
	for _, row := range rows {
		if row.ID == id {
			return row
		}
	}
	return SettingsProfileMigrationRow{}
}

func assertSettingsProfileMigrationNoSideEffects(t *testing.T, preview SettingsProfileMigrationPreview) {
	t.Helper()
	if preview.SettingsPersisted ||
		preview.SettingsPersistenceEnabled ||
		preview.ResourceGrantEnabled ||
		preview.RealPortalCallEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.FileContentRead ||
		preview.BackendDetailsExposed ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("migration preview enabled unsafe side effects: %#v", preview)
	}
	for _, row := range preview.Rows {
		if row.SettingsPersisted ||
			row.ResourceGrantEnabled ||
			row.RealPortalCallEnabled ||
			row.BackendLaunchEnabled ||
			row.HostRootModified ||
			row.BackendDetailsExposed {
			t.Fatalf("migration row enabled unsafe side effects: %#v", row)
		}
	}
}

func assertSettingsProfileMigrationSafe(t *testing.T, preview SettingsProfileMigrationPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "file://", "token=", "secret"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("migration preview leaked forbidden term %q: %s", forbidden, text)
		}
	}
}
