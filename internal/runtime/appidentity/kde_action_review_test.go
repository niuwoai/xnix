package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionReviewPreviewCapturesQueueDecisionWithoutMutating(t *testing.T) {
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

	preview, err := plan.KDEActionReviewPreview("review-file-manager-action", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionReviewPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_review.v1" ||
		preview.RequestType != "kde-action-review-preview" ||
		preview.ReviewType != "compatibility-center-kde-action-review" ||
		preview.Source != "kde-action-queue-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "ReviewKDEAction" ||
		preview.ReadMethod != "GetKDEActionReviewPreview" {
		t.Fatalf("unexpected KDE action review schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action review identity: %#v", preview)
	}
	if preview.Queue.QueueType != "compatibility-center-kde-action-queue" ||
		preview.Queue.SourceRequestType != "kde-action-queue-preview" ||
		preview.Queue.ActionCount != 7 ||
		preview.Queue.PendingActionCount != 7 ||
		preview.Queue.QueuePersisted ||
		preview.Queue.ExecutionEnabled ||
		preview.Queue.BackendDetailsExposed {
		t.Fatalf("unexpected queue summary: %#v", preview.Queue)
	}
	if preview.Action.ID != "review-file-manager-action" ||
		preview.Action.SourceType != "kde-entrypoint-action-preview" ||
		preview.Action.EntryPointID != "file-manager" ||
		preview.Action.KDEComponent != "Dolphin" ||
		preview.Action.Intent != "open-files" ||
		preview.Action.Status != "portal-review-required" ||
		preview.Action.RuntimeGate != "portal-file-open-review" ||
		!preview.Action.UserReviewRequired ||
		!preview.Action.RequiresPortal ||
		!preview.Action.RequiresRuntimeGate ||
		preview.Action.ExecutionEnabled ||
		preview.Action.RequestObjectCreated ||
		preview.Action.BackendProcessStarted ||
		preview.Action.HostRootModified ||
		preview.Action.BackendDetailsExposed {
		t.Fatalf("unexpected action summary: %#v", preview.Action)
	}
	if preview.Decision.Decision != "approved" ||
		!preview.Decision.DecisionAccepted ||
		!preview.Decision.UserIntentCaptured ||
		preview.Decision.DecisionRecorded ||
		preview.Decision.RuntimeApprovalGranted ||
		preview.Decision.ExecutionAllowed ||
		preview.Decision.ReviewReceiptCreated ||
		preview.Decision.QueueStateChanged ||
		preview.Decision.PermissionGrantCreated ||
		preview.Decision.SettingsPersisted ||
		preview.Decision.DesktopNotificationIntent != "show-kde-action-approval-pending" {
		t.Fatalf("unexpected decision summary: %#v", preview.Decision)
	}
	if preview.SessionStatus.RequestType != "execution-session-status-preview" ||
		preview.SessionStatus.SessionState != "planned-blocked" ||
		preview.SessionStatus.GateCount != 5 ||
		preview.SessionStatus.BlockedGateCount != 1 ||
		preview.SessionStatus.RuntimeLaunchApproval ||
		preview.SessionStatus.LaunchAllowed ||
		preview.SessionStatus.ExecutionStarted {
		t.Fatalf("unexpected session status summary: %#v", preview.SessionStatus)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionReviewCaptured ||
		!preview.ReviewReceiptRequired || preview.ReviewReceiptRecorded ||
		preview.ActionQueuePersisted || preview.QueueStateChanged ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action review safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "record KDE action review from preview state") ||
		!containsString(preview.BlockedActions, "treat KDE review intent as Runtime execution approval") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from review preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionReviewPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionReviewPreview("review-launcher-action", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionReviewPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.SessionStatus.SessionState != "review-declined" ||
		rejectedPreview.SessionStatus.BlockedGateCount != 2 ||
		rejectedPreview.Decision.DesktopNotificationIntent != "show-kde-action-rejected" ||
		rejectedPreview.Decision.ExecutionAllowed ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action review preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionReviewPreview("review-settings-action", "deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionReviewPreview returned error: %v", err)
	}
	if deferredPreview.Decision.DesktopNotificationIntent != "show-kde-action-deferred" ||
		deferredPreview.Decision.DecisionRecorded ||
		deferredPreview.QueueStateChanged ||
		deferredPreview.SettingsPersisted {
		t.Fatalf("unexpected deferred action review preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionReviewPreview("missing-action", "approved", nil); err == nil {
		t.Fatalf("KDEActionReviewPreview accepted an unknown action")
	}
	if _, err := plan.KDEActionReviewPreview("review-launcher-action\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEActionReviewPreview accepted a multiline action id")
	}
	if _, err := plan.KDEActionReviewPreview("review-launcher-action", "invalid", nil); err == nil {
		t.Fatalf("KDEActionReviewPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionReviewPreview("review-file-manager-action", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionReviewPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionReviewPreview("review-settings-action", "reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionReviewPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action review preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
