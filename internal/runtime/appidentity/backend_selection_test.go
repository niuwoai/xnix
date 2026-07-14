package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendSelectionPreviewExplainsProfilesWithoutCommittingSelection(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	preview, err := plan.BackendSelectionPreview()
	if err != nil {
		t.Fatalf("BackendSelectionPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.backend_selection.v1" ||
		preview.RequestType != "backend-selection-preview" ||
		preview.PlanType != "compatibility-backend-selection-plan" ||
		preview.Source != "compatibility-center" ||
		preview.RuntimeMethod != "GetBackendSelectionPlan" {
		t.Fatalf("unexpected backend selection schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected backend selection identity: %#v", preview)
	}
	if preview.SelectedStrategy != "local-compatibility-strategy" ||
		preview.RecommendedProfileID != "local-compatibility" ||
		preview.CandidateCount != 2 ||
		preview.ReadyCandidateCount != 0 ||
		preview.BlockedCandidateCount != 2 {
		t.Fatalf("unexpected backend selection recommendation: %#v", preview)
	}
	if got, want := backendSelectionProfileIDsForTest(preview.CandidateProfiles), []string{"local-compatibility", "isolated-compatibility"}; !sameStrings(got, want) {
		t.Fatalf("candidate ids = %#v, want %#v", got, want)
	}
	if !preview.CandidateProfiles[0].Recommended ||
		preview.CandidateProfiles[0].SelectionState != "recommended" ||
		preview.CandidateProfiles[0].Ready ||
		!preview.CandidateProfiles[0].Blocked ||
		preview.CandidateProfiles[0].SelectionCommitted ||
		preview.CandidateProfiles[0].EnvironmentCreated ||
		preview.CandidateProfiles[0].BackendProcessStarted ||
		preview.CandidateProfiles[0].BackendDetailsExposed {
		t.Fatalf("unexpected recommended candidate: %#v", preview.CandidateProfiles[0])
	}
	if !sameStrings(preview.RequiredReviews, []string{"backend-capability-review", "backend-binding-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review"}) {
		t.Fatalf("unexpected required reviews: %#v", preview.RequiredReviews)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.SelectionCommitted || preview.SelectionChangeEnabled ||
		preview.BackendLaunchEnabled || preview.CapabilityActivationEnabled ||
		preview.EnvironmentCreated || preview.RequestObjectCreated ||
		preview.StateRootCreated || preview.SnapshotCreated ||
		preview.HostRootModified || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected backend selection safety flags: %#v", preview)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "application-x-executable",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.BackendSelectionPreview()
	if err != nil {
		t.Fatalf("isolated BackendSelectionPreview returned error: %v", err)
	}
	if isolatedPreview.RecommendedProfileID != "isolated-compatibility" ||
		!isolatedPreview.CandidateProfiles[1].Recommended {
		t.Fatalf("isolated recipe did not recommend isolated compatibility: %#v", isolatedPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend selection preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func backendSelectionProfileIDsForTest(profiles []BackendSelectionProfile) []string {
	ids := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.ID)
	}
	return ids
}
