package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func newKDENotificationDigestPlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlanWithProvenance(Recipe{
		ID: "org.example.ledger", Name: "Example Ledger", Icon: "office-chart-area", Mode: "automatic",
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance: %v", err)
	}
	return plan
}

func allKDENotificationDigestEvents() []KDENotificationDigestEvent {
	return []KDENotificationDigestEvent{
		{Group: "needs-review", EventID: "approval-required"},
		{Group: "blocked-action", EventID: "execution-blocked"},
		{Group: "permission-attention", EventID: "permission-review"},
		{Group: "diagnostic-issue", EventID: "diagnostic-failed"},
		{Group: "snapshot-warning", EventID: "snapshot-missing"},
		{Group: "readiness-change", EventID: "readiness-pending"},
	}
}

func assertKDENotificationDigestDisabled(t *testing.T, preview KDENotificationDigestPreview) {
	t.Helper()
	if preview.NotificationsSent || preview.LiveTrayBridgeEnabled || preview.RequestObjectsCreated ||
		preview.PermissionGrantEnabled || preview.BackendLaunchEnabled || preview.BackendProcessStarted ||
		preview.StateRootPathExposed || preview.RawCommandExposed || preview.BackendDetailsExposed || preview.HostRootModified {
		t.Fatalf("notification digest enabled an unsafe capability: %+v", preview)
	}
	if err := safety.ValidatePayload("KDE notification digest", preview); err != nil {
		t.Fatalf("notification digest exposed unsafe content: %v", err)
	}
}

func TestKDENotificationDigestCoversSixGroups(t *testing.T) {
	preview, err := newKDENotificationDigestPlan(t).KDENotificationDigestPreview(allKDENotificationDigestEvents())
	if err != nil {
		t.Fatalf("KDENotificationDigestPreview: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_notification_digest.v1" ||
		preview.RequestType != "kde-notification-digest-preview" || preview.Counts.DigestEntryCount != 6 ||
		preview.Counts.CriticalEntryCount != 3 || preview.Counts.ReviewEntryCount != 5 || preview.Counts.BlockedEntryCount != 3 {
		t.Fatalf("unexpected digest: %+v", preview)
	}
	wantGroups := []string{"needs-review", "blocked-action", "permission-attention", "diagnostic-issue", "snapshot-warning", "readiness-change"}
	for index, want := range wantGroups {
		if preview.Entries[index].Group != want || preview.Entries[index].DeduplicationKey == "" || preview.Entries[index].NextSafeReadRoute == "" {
			t.Fatalf("unexpected digest entry %d: %+v", index, preview.Entries[index])
		}
	}
	assertKDENotificationDigestDisabled(t, preview)
}

func TestKDENotificationDigestDeduplicatesEvents(t *testing.T) {
	events := allKDENotificationDigestEvents()[:2]
	events = append(events, events[0], events[1])
	preview, err := newKDENotificationDigestPlan(t).KDENotificationDigestPreview(events)
	if err != nil {
		t.Fatalf("KDENotificationDigestPreview: %v", err)
	}
	if preview.Counts.InputEventCount != 4 || preview.Counts.DigestEntryCount != 2 || preview.Counts.DuplicateEventCount != 2 {
		t.Fatalf("unexpected deduplication counts: %+v", preview.Counts)
	}
	assertKDENotificationDigestDisabled(t, preview)
}

func TestKDENotificationDigestRejectsUnsupportedAndUnsafeInput(t *testing.T) {
	plan := newKDENotificationDigestPlan(t)
	for _, events := range [][]KDENotificationDigestEvent{
		nil,
		{{Group: "blocked-action", EventID: "unknown"}},
		{{Group: "needs-review\n", EventID: "approval-required"}},
	} {
		if _, err := plan.KDENotificationDigestPreview(events); err == nil {
			t.Fatalf("expected events to fail: %+v", events)
		}
	}
}

func TestKDENotificationDigestRejectsUnsafeIdentity(t *testing.T) {
	plan := newKDENotificationDigestPlan(t)
	plan.DisplayName = "token=unsafe"
	if _, err := plan.KDENotificationDigestPreview(allKDENotificationDigestEvents()[:1]); err == nil {
		t.Fatal("expected unsafe identity to fail")
	}
}
