package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionCardPreviewRendersCompatibilityCenterCard(t *testing.T) {
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

	preview, err := plan.KDEActionCardPreview("review-file-manager-action", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionCardPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_card.v1" ||
		preview.RequestType != "kde-action-card-preview" ||
		preview.CardType != "compatibility-center-kde-action-card" ||
		preview.CardState != "waiting-for-runtime-gates" ||
		preview.Source != "kde-action-status-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEActionCard" ||
		preview.ReadMethod != "GetKDEActionCardPreview" {
		t.Fatalf("unexpected KDE action card schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action card identity: %#v", preview)
	}
	if preview.Status.RequestType != "kde-action-status-preview" ||
		preview.Status.StatusType != "compatibility-center-kde-action-status" ||
		preview.Status.StatusState != "waiting-for-runtime-gates" ||
		preview.Status.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		preview.Status.NotificationIntent != "show-kde-action-status-waiting" ||
		!preview.Status.StatusPreviewCreated ||
		preview.Status.StatusPersisted ||
		preview.Status.RuntimeLaunchApproval ||
		preview.Status.LaunchAllowed ||
		preview.Status.ExecutionStarted ||
		preview.Status.BackendDetailsExposed {
		t.Fatalf("unexpected status summary: %#v", preview.Status)
	}
	if preview.Card.Title != "Example Ledger" ||
		preview.Card.Subtitle != "Waiting for desktop access review" ||
		preview.Card.Badge != "Waiting" ||
		preview.Card.BadgeTone != "warning" ||
		preview.Card.PrimaryAction.ID != "review-required-gates" ||
		preview.Card.PrimaryAction.Label != "Review required gates" ||
		!preview.Card.PrimaryAction.Enabled ||
		!preview.Card.PrimaryAction.NavigationOnly ||
		preview.Card.PrimaryAction.MutatesRuntime ||
		preview.Card.PrimaryAction.StartsProgram ||
		preview.Card.Footer != "This card is a read-only preview; execution remains blocked." ||
		preview.Card.BackendDetailsExposed {
		t.Fatalf("unexpected card visual state: %#v", preview.Card)
	}
	if len(preview.Card.SecondaryActions) != 3 ||
		preview.Card.SecondaryActions[0].ID != "open-settings" ||
		preview.Card.SecondaryActions[1].ID != "open-compatibility-center" ||
		preview.Card.SecondaryActions[2].ID != "show-status-details" {
		t.Fatalf("unexpected secondary actions: %#v", preview.Card.SecondaryActions)
	}
	for _, action := range preview.Card.SecondaryActions {
		if !action.Enabled || !action.NavigationOnly || action.MutatesRuntime || action.StartsProgram {
			t.Fatalf("secondary action must remain navigation-only: %#v", action)
		}
	}
	if len(preview.Card.DisabledActions) != 3 ||
		preview.Card.DisabledActions[0].ID != "start-application" ||
		preview.Card.DisabledActions[0].Enabled ||
		preview.Card.DisabledActions[0].MutatesRuntime ||
		preview.Card.DisabledActions[0].StartsProgram {
		t.Fatalf("unexpected disabled actions: %#v", preview.Card.DisabledActions)
	}
	if preview.Action.ID != "review-file-manager-action" ||
		preview.Action.KDEComponent != "Dolphin" ||
		preview.Preflight.RequestType != "kde-action-preflight-preview" ||
		preview.Receipt.RequestType != "kde-action-receipt-preview" ||
		preview.RequiredRuntimeGate != "portal-file-open-review" {
		t.Fatalf("unexpected action pipeline summaries: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionReviewCaptured ||
		!preview.ActionPreflightCreated || !preview.ReceiptPreviewCreated ||
		!preview.StatusPreviewCreated || !preview.CardPreviewCreated ||
		preview.CardPersisted || preview.StatusPersisted ||
		preview.DecisionRecorded || preview.ReviewReceiptRecorded ||
		preview.ActionQueuePersisted || preview.QueueStateChanged ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.ResourceGrantCreated || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.PermissionGrantCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action card safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE action card from preview state") ||
		!containsString(preview.BlockedActions, "treat KDE action card as Runtime execution approval") ||
		!containsString(preview.BlockedActions, "grant desktop resources from card preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionCardPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionCardPreview("review-launcher-action", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionCardPreview returned error: %v", err)
	}
	if rejectedPreview.CardState != "action-rejected" ||
		rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.Card.Badge != "Rejected" ||
		rejectedPreview.Card.BadgeTone != "critical" ||
		rejectedPreview.Card.PrimaryAction.ID != "show-guidance" ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action card preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionCardPreview("review-settings-action", "deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionCardPreview returned error: %v", err)
	}
	if deferredPreview.CardState != "action-deferred" ||
		deferredPreview.UserDecisionAllowsLaunch ||
		deferredPreview.Card.Badge != "Deferred" ||
		deferredPreview.Card.BadgeTone != "neutral" ||
		deferredPreview.Card.PrimaryAction.ID != "resume-review" ||
		deferredPreview.CardPersisted ||
		deferredPreview.ReviewReceiptRecorded ||
		deferredPreview.QueueStateChanged {
		t.Fatalf("unexpected deferred action card preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionCardPreview("missing-action", "approved", nil); err == nil {
		t.Fatalf("KDEActionCardPreview accepted an unknown action")
	}
	if _, err := plan.KDEActionCardPreview("review-launcher-action\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEActionCardPreview accepted a multiline action id")
	}
	if _, err := plan.KDEActionCardPreview("review-launcher-action", "invalid", nil); err == nil {
		t.Fatalf("KDEActionCardPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionCardPreview("review-file-manager-action", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionCardPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionCardPreview("review-settings-action", "reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionCardPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action card preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
