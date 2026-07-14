package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionDecisionPreviewCapturesUserIntentWithoutApproval(t *testing.T) {
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

	preview, err := plan.ExecutionDecisionPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionDecisionPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.request_decision.v1" ||
		preview.RequestType != "execution-decision-preview" ||
		preview.DecisionType != "compatibility-center-launch-decision" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-review-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionDecisionPreview" {
		t.Fatalf("unexpected execution decision schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution decision identity: %#v", preview)
	}
	if preview.ExecutionReview.SchemaVersion != "xnix.runtime.request_review.v1" ||
		preview.ExecutionReview.RequestType != "execution-review-preview" ||
		preview.ExecutionReview.ReviewType != "compatibility-center-launch-review" ||
		preview.ExecutionReview.Source != "execution-request-preview" ||
		preview.ExecutionReview.ReadMethod != "GetExecutionReviewPreview" ||
		!preview.ExecutionReview.ActionQueueCandidate ||
		!preview.ExecutionReview.ReviewReceiptRequired ||
		preview.ExecutionReview.ReviewReceiptRecorded ||
		preview.ExecutionReview.ActionQueuePersisted ||
		preview.ExecutionReview.ExecutionStarted ||
		preview.ExecutionReview.BackendDetailsExposed {
		t.Fatalf("unexpected execution review summary: %#v", preview.ExecutionReview)
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
		preview.Decision.NextStep != "Require Runtime gates before recording approval or launching." ||
		preview.Decision.DesktopNotificationIntent != "show-launch-approval-pending" {
		t.Fatalf("unexpected decision preview: %#v", preview.Decision)
	}
	if preview.ReviewCard.CardType != "compatibility-center-launch-review" ||
		preview.ReviewCard.Status != "blocked" ||
		!preview.ReviewCard.UserReviewRequired ||
		!preview.ReviewCard.RuntimeApprovalNeeded {
		t.Fatalf("unexpected review card: %#v", preview.ReviewCard)
	}
	if preview.ActionQueue.QueueType != "compatibility-center-request-review-queue" ||
		preview.ActionQueue.ExecutionEnabled ||
		preview.ActionQueue.QueuePersisted ||
		preview.ActionQueue.ReviewReceiptRecorded {
		t.Fatalf("unexpected action queue: %#v", preview.ActionQueue)
	}
	if preview.GateSummary.GateCount != 5 ||
		preview.GateSummary.RequiredGateCount != 2 ||
		preview.GateSummary.PendingGateCount != 1 ||
		preview.GateSummary.BlockedGateCount != 1 {
		t.Fatalf("unexpected gate summary: %#v", preview.GateSummary)
	}
	if preview.ExecutionState != "blocked" ||
		preview.OverallStatus != "not-ready" ||
		preview.WriteGateDecision != "blocked-until-production-backend" ||
		preview.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		!preview.PortalRequired ||
		!preview.SnapshotRequired ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution decision state: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.ActionQueueCandidate || !preview.UserDecisionCaptured ||
		!preview.ReviewReceiptRequired || preview.ReviewReceiptRecorded ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.ExecutionRequestPersisted || preview.ActionQueuePersisted ||
		preview.ExecutionStarted || preview.BackendBindingReady ||
		preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution decision safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "record launch decision from preview state") ||
		!containsString(preview.BlockedActions, "treat user decision as Runtime launch approval") ||
		!containsString(preview.BlockedActions, "start compatibility profile from decision preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionDecisionPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionDecisionPreview returned error: %v", err)
	}
	if rejectedPreview.Decision.NextStep != "Keep the launch request blocked and show rejection guidance." ||
		rejectedPreview.Decision.DesktopNotificationIntent != "show-launch-rejected" {
		t.Fatalf("unexpected rejected decision: %#v", rejectedPreview.Decision)
	}

	deferredPreview, err := plan.ExecutionDecisionPreview("deferred", nil)
	if err != nil {
		t.Fatalf("deferred ExecutionDecisionPreview returned error: %v", err)
	}
	if deferredPreview.Decision.NextStep != "Keep the launch request pending for later review." ||
		deferredPreview.Decision.DesktopNotificationIntent != "show-launch-deferred" {
		t.Fatalf("unexpected deferred decision: %#v", deferredPreview.Decision)
	}

	if _, err := plan.ExecutionDecisionPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionDecisionPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionDecisionPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionDecisionPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution decision preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
