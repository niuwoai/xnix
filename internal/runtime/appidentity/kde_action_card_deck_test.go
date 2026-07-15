package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionCardDeckPreviewRendersSevenCards(t *testing.T) {
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

	preview, err := plan.KDEActionCardDeckPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionCardDeckPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_card_deck.v1" ||
		preview.RequestType != "kde-action-card-deck-preview" ||
		preview.DeckType != "compatibility-center-kde-action-card-deck" ||
		preview.Source != "kde-action-card-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEActionCardDeck" ||
		preview.ReadMethod != "GetKDEActionCardDeckPreview" {
		t.Fatalf("unexpected KDE action card deck schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action card deck identity: %#v", preview)
	}
	if preview.Queue.RequestType != "kde-action-queue-preview" ||
		preview.Queue.QueueType != "compatibility-center-kde-action-queue" ||
		preview.Queue.Source != "kde-entrypoint-action-preview" ||
		preview.Queue.ActionCount != 7 ||
		preview.Queue.PendingActionCount != 7 ||
		preview.Queue.UserReviewRequiredCount != 4 ||
		preview.Queue.PortalActionCount != 1 ||
		preview.Queue.RuntimeGateActionCount != 7 ||
		!preview.Queue.ActionQueueCreated ||
		preview.Queue.ActionQueuePersisted ||
		preview.Queue.RuntimeLaunchApproval ||
		preview.Queue.LaunchAllowed ||
		preview.Queue.ExecutionStarted ||
		preview.Queue.BackendDetailsExposed {
		t.Fatalf("unexpected queue summary: %#v", preview.Queue)
	}
	if preview.CardCount != 7 ||
		len(preview.Cards) != 7 ||
		len(preview.CardIDs) != 7 ||
		preview.WaitingCardCount != 7 ||
		preview.DeferredCardCount != 0 ||
		preview.RejectedCardCount != 0 ||
		preview.NavigationActionCount != 28 ||
		preview.DisabledActionCount != 21 ||
		preview.PrimaryCardID != "org.example.ledger:review-launcher-action:card" {
		t.Fatalf("unexpected deck counts: %#v", preview)
	}
	first := preview.Cards[0]
	if first.CardID != "org.example.ledger:review-launcher-action:card" ||
		first.ActionID != "review-launcher-action" ||
		first.EntryPointID != "launcher" ||
		first.KDEComponent != "Plasma application launcher" ||
		first.CardState != "waiting-for-runtime-gates" ||
		first.RequiredRuntimeGate != "runtime-launch-review" ||
		!first.UserReviewRequired ||
		first.RequiresPortal ||
		!first.RequiresRuntimeGate ||
		first.Card.Title != "Example Ledger" ||
		first.Card.Badge != "Waiting" ||
		first.Card.BadgeTone != "warning" ||
		first.PrimaryAction.ID != "review-required-gates" ||
		!first.PrimaryAction.NavigationOnly ||
		first.PrimaryAction.MutatesRuntime ||
		first.PrimaryAction.StartsProgram ||
		first.NavigationActionCount != 4 ||
		first.DisabledActionCount != 3 ||
		!first.CardPreviewCreated ||
		first.CardPersisted ||
		first.ExecutionEnabled ||
		first.RequestObjectsCreated ||
		first.ResourceGrantCreated ||
		first.NotificationsSent ||
		first.MutatesRuntime ||
		first.StartsProgram ||
		first.HostRootModified ||
		first.BackendDetailsExposed {
		t.Fatalf("unexpected first deck card: %#v", first)
	}
	fileCard := preview.Cards[2]
	if fileCard.ActionID != "review-file-manager-action" ||
		fileCard.EntryPointID != "file-manager" ||
		fileCard.KDEComponent != "Dolphin" ||
		fileCard.RequiredRuntimeGate != "portal-file-open-review" ||
		!fileCard.RequiresPortal ||
		!fileCard.UserReviewRequired {
		t.Fatalf("unexpected file-manager deck card: %#v", fileCard)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterDeck ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionQueueCreated ||
		preview.ActionQueuePersisted || !preview.DeckPreviewCreated ||
		preview.DeckPersisted || preview.CardsPersisted ||
		preview.CardActionsEnabled || preview.CardActionsPersisted ||
		preview.StatusPersisted || preview.DecisionRecorded ||
		preview.ReviewReceiptRecorded || preview.QueueStateChanged ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.ResourceGrantCreated || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.PermissionGrantCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action card deck safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE action card deck from preview state") ||
		!containsString(preview.BlockedActions, "enable card actions from deck preview") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from deck preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionCardDeckPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionCardDeckPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionCardDeckPreview returned error: %v", err)
	}
	if rejectedPreview.WaitingCardCount != 0 ||
		rejectedPreview.DeferredCardCount != 0 ||
		rejectedPreview.RejectedCardCount != 7 ||
		rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.Cards[0].Card.Badge != "Rejected" ||
		rejectedPreview.Cards[0].Card.BadgeTone != "critical" ||
		rejectedPreview.Cards[0].PrimaryAction.ID != "show-guidance" ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action card deck preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionCardDeckPreview("deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionCardDeckPreview returned error: %v", err)
	}
	if deferredPreview.WaitingCardCount != 0 ||
		deferredPreview.DeferredCardCount != 7 ||
		deferredPreview.RejectedCardCount != 0 ||
		deferredPreview.UserDecisionAllowsLaunch ||
		deferredPreview.Cards[0].Card.Badge != "Deferred" ||
		deferredPreview.Cards[0].Card.BadgeTone != "neutral" ||
		deferredPreview.Cards[0].PrimaryAction.ID != "resume-review" ||
		deferredPreview.DeckPersisted ||
		deferredPreview.CardsPersisted ||
		deferredPreview.QueueStateChanged {
		t.Fatalf("unexpected deferred action card deck preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionCardDeckPreview("invalid", nil); err == nil {
		t.Fatalf("KDEActionCardDeckPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionCardDeckPreview("approved\nbad", nil); err == nil {
		t.Fatalf("KDEActionCardDeckPreview accepted a multiline decision")
	}
	if _, err := plan.KDEActionCardDeckPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionCardDeckPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionCardDeckPreview("reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionCardDeckPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action card deck preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
