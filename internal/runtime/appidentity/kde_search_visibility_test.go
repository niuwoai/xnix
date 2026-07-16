package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func newSearchVisibilityPlan(t *testing.T, extensions []string) Plan {
	t.Helper()
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: extensions,
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance: %v", err)
	}
	return plan
}

func searchRow(t *testing.T, preview KDESearchVisibilityPlanPreview, id string) KDESearchVisibilityRow {
	t.Helper()
	for _, row := range preview.Rows {
		if row.ID == id {
			return row
		}
	}
	t.Fatalf("search visibility row %q not present; ids=%v", id, preview.RowIDs)
	return KDESearchVisibilityRow{}
}

func assertSearchVisibilitySafe(t *testing.T, preview KDESearchVisibilityPlanPreview) {
	t.Helper()
	if preview.DesktopFilesWritten || preview.MIMEDefaultsWritten || preview.KDECacheRefreshed ||
		preview.HostFilesIndexed || preview.SearchIndexPersisted || preview.BackendLaunchEnabled ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("search visibility preview must keep every unsafe capability disabled: %+v", preview)
	}
	if err := safety.ValidatePayload("KDE search visibility plan preview", preview); err != nil {
		t.Fatalf("search visibility preview leaks forbidden content: %v", err)
	}
}

func TestKDESearchVisibilityReceiptRequiredWithoutActivation(t *testing.T) {
	plan := newSearchVisibilityPlan(t, []string{".abc"})
	preview, err := plan.KDESearchVisibilityPlanPreview(KDESearchVisibilityOptions{SupportedExtensions: []string{".abc"}})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "receipt-required" || preview.RowCount != 7 {
		t.Fatalf("unactivated app should be receipt-required across 7 rows: %+v", preview.OverallState)
	}
	for _, id := range []string{"launcher", "krunner", "file-association", "dolphin-action"} {
		if row := searchRow(t, preview, id); row.VisibilityState != "receipt-required" || row.DisabledActionReason != "activation-receipt-required" {
			t.Fatalf("row %q should require an activation receipt: %+v", id, row)
		}
	}
	for _, id := range []string{"compatibility-center", "settings", "task-manager"} {
		if row := searchRow(t, preview, id); row.VisibilityState != "visible" {
			t.Fatalf("row %q should be visible without a receipt: %+v", id, row)
		}
	}
	assertSearchVisibilitySafe(t, preview)
}

func TestKDESearchVisibilityUnsupportedMIME(t *testing.T) {
	plan := newSearchVisibilityPlan(t, nil)
	preview, err := plan.KDESearchVisibilityPlanPreview(KDESearchVisibilityOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	for _, id := range []string{"file-association", "dolphin-action"} {
		if row := searchRow(t, preview, id); row.VisibilityState != "no-mime-association" || row.DisabledActionReason != "unsupported-mime" {
			t.Fatalf("row %q should report no MIME association: %+v", id, row)
		}
	}
	assertSearchVisibilitySafe(t, preview)
}

func TestKDESearchVisibilityHiddenApplication(t *testing.T) {
	plan := newSearchVisibilityPlan(t, []string{".abc"})
	plan.UserVisible = false
	preview, err := plan.KDESearchVisibilityPlanPreview(KDESearchVisibilityOptions{SupportedExtensions: []string{".abc"}})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "hidden" || preview.Counts.Hidden != preview.RowCount {
		t.Fatalf("hidden app should hide every row: %+v", preview.Counts)
	}
	for _, row := range preview.Rows {
		if row.VisibilityState != "hidden" || row.DisabledActionReason != "application-hidden" {
			t.Fatalf("row %q should be hidden: %+v", row.ID, row)
		}
	}
	assertSearchVisibilitySafe(t, preview)
}

func TestKDESearchSafeSynonymsDropUnsafeAndDuplicates(t *testing.T) {
	got := kdeSearchSafeSynonyms([]string{"notes", ".exe", "notes", "documents", "wineprefix", "  ", "proton"})
	for _, unsafe := range []string{".exe", "wineprefix", "proton"} {
		for _, value := range got {
			if value == unsafe {
				t.Fatalf("unsafe synonym %q must be dropped: %+v", unsafe, got)
			}
		}
	}
	seen := map[string]int{}
	for _, value := range got {
		seen[value]++
		if seen[value] > 1 {
			t.Fatalf("duplicate synonym %q must be deduplicated: %+v", value, got)
		}
	}
	if len(got) != 2 { // notes, documents
		t.Fatalf("expected only the two safe synonyms, got %+v", got)
	}
}

func TestKDESearchVisibilityRejectsUnsafePlan(t *testing.T) {
	plan := newSearchVisibilityPlan(t, []string{".abc"})
	plan.BackendDetailsExposed = true
	if _, err := plan.KDESearchVisibilityPlanPreview(KDESearchVisibilityOptions{}); err == nil {
		t.Fatalf("expected an unsafe plan to be rejected")
	}
}
