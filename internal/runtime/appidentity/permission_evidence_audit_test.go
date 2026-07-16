package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func permissionAuditRowByID(t *testing.T, preview PermissionEvidenceAuditPreview, id string) PermissionEvidenceAuditRow {
	t.Helper()
	for _, row := range preview.Rows {
		if row.ID == id {
			return row
		}
	}
	t.Fatalf("audit row %q not present; ids=%v", id, preview.RowIDs)
	return PermissionEvidenceAuditRow{}
}

func assertPermissionAuditSafe(t *testing.T, preview PermissionEvidenceAuditPreview) {
	t.Helper()
	if preview.RealPortalCallEnabled || preview.PermissionGrantEnabled || preview.PermissionRevokeEnabled ||
		preview.ReceiptWriteEnabled || preview.SettingsPersisted || preview.ExecutionApproved ||
		preview.StateRootPathExposed || preview.HostRootModified || preview.NetworkRequired ||
		preview.PrivilegedContainerRequired || preview.BackendDetailsExposed {
		t.Fatalf("audit preview must keep every unsafe capability disabled: %+v", preview)
	}
	if !preview.ReviewOnly || !preview.RuntimeOwned || !preview.GoRuntimeBacked {
		t.Fatalf("audit preview must be a runtime-owned review-only read model: %+v", preview)
	}
	if err := safety.ValidatePayload("permission evidence audit preview", preview); err != nil {
		t.Fatalf("audit preview leaks forbidden content: %v", err)
	}
}

func TestPermissionEvidenceAuditConsistencyStates(t *testing.T) {
	input := PermissionEvidenceAuditInput{
		ApplicationID: "org.example.ledger",
		Rows: []PermissionEvidenceRowInput{
			{Resource: "documents", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved", BridgeExposed: true},
			{Resource: "downloads", SettingState: "enabled", ReceiptState: "absent", ReviewState: "approved"},
			{Resource: "print", SettingState: "enabled", ReceiptState: "expired", ReviewState: "approved"},
			{Resource: "clipboard", SettingState: "enabled", ReceiptState: "denied", ReviewState: "approved"},
			{Resource: "screenshot", SettingState: "enabled", ReceiptState: "granted", ReviewState: "pending"},
			{Resource: "camera", SettingState: "disabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "remote-desktop", SettingState: "disabled", ReceiptState: "absent", ReviewState: "absent"},
			{Resource: "network", SettingState: "enabled", ReceiptState: "absent", ReviewState: "approved"},
		},
	}
	preview, err := NewPermissionEvidenceAuditPreview(input)
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	cases := map[string]string{
		"documents":      "consistent",
		"downloads":      "setting-only",
		"print":          "expired",
		"clipboard":      "denied",
		"screenshot":     "missing-review",
		"camera":         "receipt-only",
		"remote-desktop": "blocked-by-policy",
		"network":        "consistent",
	}
	for id, want := range cases {
		if got := permissionAuditRowByID(t, preview, id).ConsistencyState; got != want {
			t.Fatalf("row %q consistency = %q, want %q", id, got, want)
		}
	}
	if preview.OverallState != "needs-review" {
		t.Fatalf("mixed audit should need review, got %q", preview.OverallState)
	}
	if preview.Counts.Expired != 1 || preview.Counts.Denied != 1 || preview.Counts.MissingReview != 1 ||
		preview.Counts.ReceiptOnly != 1 || preview.Counts.SettingOnly != 1 {
		t.Fatalf("unexpected divergence counts: %+v", preview.Counts)
	}
	assertPermissionAuditSafe(t, preview)
}

func TestPermissionEvidenceAuditAllConsistentIsClean(t *testing.T) {
	input := PermissionEvidenceAuditInput{
		ApplicationID: "org.example.ledger",
		Rows: []PermissionEvidenceRowInput{
			{Resource: "documents", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "downloads", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "uris", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "print", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "clipboard", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
			{Resource: "screenshot", SettingState: "enabled", ReceiptState: "granted", ReviewState: "approved"},
		},
	}
	preview, err := NewPermissionEvidenceAuditPreview(input)
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	// camera, remote-desktop stay blocked-by-policy; network stays consistent.
	if preview.OverallState != "consistent" {
		t.Fatalf("granted+blocked-by-policy audit should be consistent, got %q", preview.OverallState)
	}
	if preview.Counts.Divergent != 0 {
		t.Fatalf("no divergence expected: %+v", preview.Counts)
	}
	assertPermissionAuditSafe(t, preview)
}

func TestPermissionEvidenceAuditMalformedReceiptsAreSurfaced(t *testing.T) {
	preview, err := NewPermissionEvidenceAuditPreview(PermissionEvidenceAuditInput{
		ApplicationID:     "org.example.ledger",
		MalformedReceipts: []string{"broken-1", "/home/user/secret-token"},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "needs-review" {
		t.Fatalf("malformed receipts should need review, got %q", preview.OverallState)
	}
	if preview.Counts.MalformedReceipts != 2 {
		t.Fatalf("malformed count should include both entries, got %d", preview.Counts.MalformedReceipts)
	}
	if len(preview.MalformedReceiptIDs) != 1 || preview.MalformedReceiptIDs[0] != "broken-1" {
		t.Fatalf("privacy-sensitive malformed id should be dropped from output: %+v", preview.MalformedReceiptIDs)
	}
	assertPermissionAuditSafe(t, preview)
}

func TestPermissionEvidenceAuditUnsupportedResourceIsFlagged(t *testing.T) {
	preview, err := NewPermissionEvidenceAuditPreview(PermissionEvidenceAuditInput{
		ApplicationID: "org.example.ledger",
		Rows: []PermissionEvidenceRowInput{
			{Resource: "location-tracking", SettingState: "enabled", ReceiptState: "granted"},
		},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.Counts.Unsupported != 1 {
		t.Fatalf("unsupported resource should be counted, got %d", preview.Counts.Unsupported)
	}
	row := permissionAuditRowByID(t, preview, "location-tracking")
	if row.ConsistencyState != "unsupported" {
		t.Fatalf("unsupported row should be flagged: %+v", row)
	}
	assertPermissionAuditSafe(t, preview)
}

func TestPermissionEvidenceAuditRejectsUnsafeInputs(t *testing.T) {
	if _, err := NewPermissionEvidenceAuditPreview(PermissionEvidenceAuditInput{ApplicationID: "Not Valid"}); err == nil {
		t.Fatalf("expected invalid application id to be rejected")
	}
}

func TestPlanPermissionEvidenceAuditPreviewJoinsSources(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan: %v", err)
	}
	receiptStates := map[string]string{"file-open": "granted", "camera": "denied"}
	receiptRefs := map[string][]string{"file-open": {"portal-operation:file-open"}}
	preview, err := plan.PermissionEvidenceAuditPreview(receiptStates, receiptRefs, nil)
	if err != nil {
		t.Fatalf("PermissionEvidenceAuditPreview: %v", err)
	}
	if preview.ApplicationID != "org.example.ledger" || preview.RowCount != 9 {
		t.Fatalf("unexpected audit shape: app=%q rows=%d", preview.ApplicationID, preview.RowCount)
	}
	if docs := permissionAuditRowByID(t, preview, "documents"); docs.ConsistencyState != "consistent" || docs.ReceiptState != "granted" || !docs.BridgeExposed {
		t.Fatalf("documents row should be consistent+granted+bridged: %+v", docs)
	}
	if cam := permissionAuditRowByID(t, preview, "camera"); cam.ConsistencyState != "blocked-by-policy" {
		t.Fatalf("camera row should stay blocked-by-policy: %+v", cam)
	}
	assertPermissionAuditSafe(t, preview)
}
