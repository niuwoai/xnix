package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimePolicyExplanationCardsPreviewBuildsConsistentDeck(t *testing.T) {
	plan := testRuntimePolicyExplanationPlan(t)

	preview, err := plan.RuntimePolicyExplanationCardsPreview(RuntimePolicyExplanationCardsOptions{
		RuntimeRoot: "../../..",
		Issue:       "portal-approval-required",
	})
	if err != nil {
		t.Fatalf("RuntimePolicyExplanationCardsPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.policy_explanation_cards.v1" ||
		preview.RequestType != "runtime-policy-explanation-cards-preview" ||
		preview.CardDeckType != "kde-runtime-policy-explanation-card-deck" ||
		preview.Source != "application-readiness-preview+runtime-write-gate-preview+desktop-safety-policy-preview+repair-plan-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetRuntimePolicyExplanationCards" ||
		preview.ReadMethod != "GetRuntimePolicyExplanationCardsPreview" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.CardCount != 11 ||
		preview.Counts.Total != 11 ||
		preview.Counts.Blocked == 0 ||
		preview.Counts.MissingEvidence == 0 ||
		preview.Counts.ReviewOnly == 0 ||
		preview.Counts.NotYetImplemented == 0 {
		t.Fatalf("unexpected policy explanation preview: %#v", preview)
	}
	if !sameStrings(preview.CardIDs, []string{
		"install",
		"launch",
		"execution",
		"portal-permission",
		"snapshot",
		"diagnostics",
		"repair",
		"settings",
		"desktop-activation",
		"backend-readiness",
		"unsupported-production-route",
	}) {
		t.Fatalf("unexpected card ids: %#v", preview.CardIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.ReviewOnly ||
		preview.CardsPersisted ||
		preview.ActionEnablementChanged ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantsCreated ||
		preview.SettingsPersisted ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.BackendProcessStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected policy explanation safety flags: %#v", preview)
	}
	for _, card := range preview.Cards {
		if card.ActionEnabled ||
			card.CardPersisted ||
			card.RequestObjectCreated ||
			card.PermissionGranted ||
			card.SettingsPersisted ||
			card.AIProviderCalled ||
			card.BackendProcessStarted ||
			card.LaunchEnabled ||
			card.ExecutionStarted ||
			card.HostRootModified ||
			card.StateRootPathExposed ||
			card.RawCommandExposed ||
			card.BackendDetailsExposed {
			t.Fatalf("card enabled a side effect: %#v", card)
		}
		assertRuntimePolicyUserSummarySafe(t, card.UserFacingSummary)
	}
	assertRuntimePolicyExplanationSafe(t, preview)
}

func TestRuntimePolicyExplanationCardsPreviewDeduplicatesBlockers(t *testing.T) {
	card := runtimePolicyExplanationCard(
		"portal-permission",
		"portal-permission",
		"review-only",
		"warning",
		"user",
		[]string{"portal-review", "portal-review", "portal-access-policy-preview"},
		"Desktop resource access needs user-mediated review.",
		"Portal evidence is required before sensitive desktop resource access can be trusted.",
		"portal-access-policy-preview",
		[]string{"Portal permission receipt is required before sensitive desktop access.", "Portal permission receipt is required before sensitive desktop access."},
	)
	if !sameStrings(card.RelatedEvidenceIDs, []string{"portal-access-policy-preview", "portal-review"}) ||
		card.BlockerReasonCount != 1 ||
		len(card.BlockerReasons) != 1 {
		t.Fatalf("expected deduped evidence and blocker reasons: %#v", card)
	}
}

func TestRuntimePolicyExplanationCardsPreviewRejectsBadMode(t *testing.T) {
	plan := testRuntimePolicyExplanationPlan(t)
	if _, err := plan.RuntimePolicyExplanationCardsPreview(RuntimePolicyExplanationCardsOptions{Environment: "unsafe", RuntimeRoot: "../../.."}); err == nil {
		t.Fatalf("RuntimePolicyExplanationCardsPreview accepted an unsafe environment")
	}
}

func TestRuntimePolicyExplanationCardsPreviewRejectsBadRepairIssue(t *testing.T) {
	plan := testRuntimePolicyExplanationPlan(t)
	if _, err := plan.RuntimePolicyExplanationCardsPreview(RuntimePolicyExplanationCardsOptions{Issue: "unknown-issue", RuntimeRoot: "../../.."}); err == nil {
		t.Fatalf("RuntimePolicyExplanationCardsPreview accepted an unknown repair issue")
	}
}

func testRuntimePolicyExplanationPlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Version:             "1.0.0",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc"},
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}

func assertRuntimePolicyExplanationSafe(t *testing.T, preview RuntimePolicyExplanationCardsPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, required := range []string{"cards_persisted", "action_enablement_changed", "request_objects_created", "permission_grants_created", "settings_persisted", "ai_provider_called", "backend_process_started", "launch_enabled", "host_root_modified", "raw_command_exposed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("policy explanation JSON must include %q: %s", required, string(encoded))
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("policy explanation cards expose forbidden term %q: %s", forbidden, string(encoded))
		}
	}
	if err := validateNoBackendTerms(preview, "Runtime policy explanation cards preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func assertRuntimePolicyUserSummarySafe(t *testing.T, summary string) {
	t.Helper()
	text := strings.ToLower(summary)
	for _, forbidden := range []string{"backend", "prefix", "wine", "proton", ".exe", "virtual machine", "qemu", "program files"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("user-facing summary exposed %q: %s", forbidden, summary)
		}
	}
}
