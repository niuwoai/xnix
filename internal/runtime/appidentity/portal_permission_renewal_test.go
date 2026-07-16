package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPortalPermissionRenewalStates(t *testing.T) {
	plan := newPortalPermissionRenewalPlan(t)
	preview, err := plan.PortalPermissionRenewalPreview([]PortalPermissionRenewalReceipt{
		{Operation: "file-open", State: "granted", EvidenceID: "portal-operation:file-open"},
		{Operation: "uri-open", State: "granted", ExpiringSoon: true, EvidenceID: "portal-operation:uri-open"},
		{Operation: "print", State: "pending", EvidenceID: "portal-operation:print"},
		{Operation: "clipboard", State: "denied", EvidenceID: "portal-operation:clipboard"},
		{Operation: "screenshot", State: "revoked", EvidenceID: "portal-operation:screenshot"},
	}, nil)
	if err != nil {
		t.Fatalf("PortalPermissionRenewalPreview: %v", err)
	}
	states := map[string]string{}
	for _, row := range preview.Rows {
		states[row.ID] = row.RenewalState
	}
	expected := map[string]string{
		"files":          "current",
		"uris":           "expiring-soon",
		"print":          "needs-review",
		"clipboard":      "denied",
		"screen":         "revoked",
		"camera":         "blocked-by-policy",
		"remote-desktop": "blocked-by-policy",
		"network":        "current",
	}
	for id, state := range expected {
		if states[id] != state {
			t.Fatalf("row %s state = %q, want %q in %#v", id, states[id], state, states)
		}
	}
	if preview.Counts.Current != 2 ||
		preview.Counts.ExpiringSoon != 1 ||
		preview.Counts.NeedsReview != 1 ||
		preview.Counts.Denied != 1 ||
		preview.Counts.Revoked != 1 ||
		preview.Counts.BlockedByPolicy != 2 {
		t.Fatalf("unexpected renewal counts: %#v", preview.Counts)
	}
	assertPortalPermissionRenewalNoSideEffects(t, preview)
	assertPortalPermissionRenewalSafe(t, preview)
}

func TestPortalPermissionRenewalMissingReceiptAndMalformedLedger(t *testing.T) {
	plan := newPortalPermissionRenewalPlan(t)
	preview, err := plan.PortalPermissionRenewalPreview(nil, []string{"bad-receipt", "/Users/example/state.json", "token=secret"})
	if err != nil {
		t.Fatalf("PortalPermissionRenewalPreview: %v", err)
	}
	states := map[string]string{}
	for _, row := range preview.Rows {
		states[row.ID] = row.RenewalState
	}
	if states["files"] != "missing-receipt" || states["camera"] != "blocked-by-policy" || states["network"] != "current" {
		t.Fatalf("missing receipt preview states wrong: %#v", states)
	}
	if preview.Counts.MalformedReceipts != 3 || len(preview.MalformedReceiptIDs) != 1 {
		t.Fatalf("malformed receipts must be counted and redacted: %#v / %#v", preview.Counts, preview.MalformedReceiptIDs)
	}
	assertPortalPermissionRenewalNoSideEffects(t, preview)
	assertPortalPermissionRenewalSafe(t, preview)
}

func TestPortalPermissionRenewalRejectsUnsafePlan(t *testing.T) {
	plan := newPortalPermissionRenewalPlan(t)
	plan.ApplicationID = "not valid"
	if _, err := plan.PortalPermissionRenewalPreview(nil, nil); err == nil {
		t.Fatalf("PortalPermissionRenewalPreview accepted invalid application id")
	}
}

func newPortalPermissionRenewalPlan(t *testing.T) Plan {
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

func assertPortalPermissionRenewalNoSideEffects(t *testing.T, preview PortalPermissionRenewalPreview) {
	t.Helper()
	if preview.RealPortalCallEnabled ||
		preview.PermissionGrantEnabled ||
		preview.PermissionRevokeEnabled ||
		preview.ReceiptWriteEnabled ||
		preview.SettingsPersisted ||
		preview.ExecutionApproved ||
		preview.StateRootPathExposed ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("renewal preview enabled unsafe side effects: %#v", preview)
	}
	for _, row := range preview.Rows {
		if row.RealPortalCallEnabled ||
			row.PermissionGrantEnabled ||
			row.PermissionRevokeEnabled ||
			row.ReceiptWriteEnabled ||
			row.SettingsPersisted ||
			row.ExecutionApproved ||
			row.HostRootModified ||
			row.BackendDetailsExposed ||
			row.StateRootPathExposed {
			t.Fatalf("renewal row enabled unsafe side effects: %#v", row)
		}
	}
}

func assertPortalPermissionRenewalSafe(t *testing.T, preview PortalPermissionRenewalPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("marshal renewal preview: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"/users/example", "token=", "secret", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/private", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("renewal preview leaked forbidden term %q: %s", forbidden, text)
		}
	}
}
