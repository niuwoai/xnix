package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCompatibilityOnboardingChecklistPreviewAggregatesFirstRunEvidence(t *testing.T) {
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

	preview, err := plan.CompatibilityOnboardingChecklistPreview(CompatibilityOnboardingChecklistOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("CompatibilityOnboardingChecklistPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.compatibility_onboarding_checklist.v1" ||
		preview.RequestType != "compatibility-onboarding-checklist-preview" ||
		preview.ChecklistType != "first-run-compatibility-onboarding" ||
		preview.RuntimeMethod != "GetCompatibilityOnboardingChecklist" ||
		preview.ReadMethod != "GetCompatibilityOnboardingChecklistPreview" ||
		preview.Desktop != "KDE Plasma" {
		t.Fatalf("unexpected checklist schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Application.RuntimeMode != "automatic" {
		t.Fatalf("unexpected checklist identity: %#v", preview.Application)
	}
	if got, want := preview.SectionIDs, []string{
		"runtime-owner-readiness",
		"recipe-trust",
		"artifact-staging",
		"backend-lifecycle",
		"portal-review",
		"snapshot-baseline",
		"diagnostics-privacy",
		"kde-entry-points",
		"production-activation",
	}; !sameStrings(got, want) {
		t.Fatalf("unexpected checklist section ids: %#v", got)
	}
	if preview.SectionCount != 9 ||
		preview.States.Ready == 0 ||
		preview.States.NeedsReview == 0 ||
		preview.States.MissingEvidence == 0 ||
		preview.States.Blocked == 0 ||
		preview.States.NotYetImplemented == 0 ||
		preview.Ready {
		t.Fatalf("unexpected checklist state counts: %#v", preview)
	}

	assertOnboardingSection(t, preview, "runtime-owner-readiness", "needs-review", "runtime-owner-readiness-preview")
	assertOnboardingSection(t, preview, "recipe-trust", "needs-review", "runtime-owner-recipe-trust-preview")
	assertOnboardingSection(t, preview, "artifact-staging", "missing-evidence", "application-readiness-preview")
	assertOnboardingSection(t, preview, "backend-lifecycle", "blocked", "backend-lifecycle-preview")
	assertOnboardingSection(t, preview, "portal-review", "needs-review", "portal-access-policy-preview")
	assertOnboardingSection(t, preview, "snapshot-baseline", "missing-evidence", "snapshot-plan-preview")
	assertOnboardingSection(t, preview, "diagnostics-privacy", "ready", "ai-diagnostic-input-preview")
	assertOnboardingSection(t, preview, "kde-entry-points", "ready", "kde-journey-evidence-preview")
	assertOnboardingSection(t, preview, "production-activation", "not-yet-implemented", "runtime-write-gate-preview")

	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.RuntimeWriteMethodsEnabled ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantsCreated ||
		preview.ArtifactStaged ||
		preview.SettingsPersisted ||
		preview.BackendProcessStarted ||
		preview.LaunchEnabled ||
		preview.NetworkRequired ||
		preview.HostPackageManagerInvoked ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.AIProviderCalled ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.FileContentRead ||
		preview.BackendDetailsExposed {
		t.Fatalf("onboarding checklist enabled unsafe behavior: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	assertCompatibilityOnboardingSafe(t, string(encoded))
}

func TestCompatibilityOnboardingChecklistPreviewRejectsMalformedInput(t *testing.T) {
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
	if _, err := plan.CompatibilityOnboardingChecklistPreview(CompatibilityOnboardingChecklistOptions{RuntimeRoot: "../../..", PortalOperation: "full-disk"}); err == nil {
		t.Fatalf("CompatibilityOnboardingChecklistPreview accepted an unsupported Portal operation")
	}
	if _, err := plan.CompatibilityOnboardingChecklistPreview(CompatibilityOnboardingChecklistOptions{RuntimeRoot: "../../..", SnapshotReason: "unsafe"}); err == nil {
		t.Fatalf("CompatibilityOnboardingChecklistPreview accepted an unsupported snapshot reason")
	}
}

func assertOnboardingSection(t *testing.T, preview CompatibilityOnboardingChecklistPreview, id string, state string, nextCheck string) {
	t.Helper()
	section := findOnboardingSection(preview.Sections, id)
	if section.ID == "" {
		t.Fatalf("missing onboarding section %q", id)
	}
	if section.State != state ||
		section.NextSafeReadOnlyCheck != nextCheck ||
		!section.RuntimeOwned ||
		!section.GoRuntimeBacked ||
		section.KDEPolicyOwner ||
		section.SideEffectsEnabled ||
		section.HostRootModified ||
		section.BackendDetailsExposed {
		t.Fatalf("unexpected onboarding section %q: %#v", id, section)
	}
}

func findOnboardingSection(sections []CompatibilityOnboardingSection, id string) CompatibilityOnboardingSection {
	for _, section := range sections {
		if section.ID == id {
			return section
		}
	}
	return CompatibilityOnboardingSection{}
}

func assertCompatibilityOnboardingSafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("compatibility onboarding checklist exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
