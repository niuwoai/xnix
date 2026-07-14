package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionReviewPreviewBuildsCompatibilityCenterCard(t *testing.T) {
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

	preview, err := plan.ExecutionReviewPreview([]string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionReviewPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.request_review.v1" ||
		preview.RequestType != "execution-review-preview" ||
		preview.ReviewType != "compatibility-center-launch-review" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-request-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionReviewPreview" {
		t.Fatalf("unexpected execution review schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution review identity: %#v", preview)
	}
	if preview.ExecutionRequest.SchemaVersion != "xnix.runtime.request_intake.v1" ||
		preview.ExecutionRequest.RequestType != "execution-request-preview" ||
		preview.ExecutionRequest.RequestState != "blocked" ||
		preview.ExecutionRequest.Source != "runtime-launch-intent" ||
		preview.ExecutionRequest.ReadMethod != "GetExecutionRequestPreview" ||
		!preview.ExecutionRequest.PortalRequired ||
		!preview.ExecutionRequest.SnapshotRequired ||
		preview.ExecutionRequest.FileCount != 1 ||
		!preview.ExecutionRequest.LaunchIntentCaptured ||
		preview.ExecutionRequest.ExecutionRequestCreated ||
		preview.ExecutionRequest.ExecutionRequestPersisted ||
		preview.ExecutionRequest.ExecutionStarted ||
		preview.ExecutionRequest.BackendDetailsExposed {
		t.Fatalf("unexpected execution request summary: %#v", preview.ExecutionRequest)
	}
	if preview.ReviewCard.ID != "org.example.ledger:launch-review" ||
		preview.ReviewCard.CardType != "compatibility-center-launch-review" ||
		preview.ReviewCard.Status != "blocked" ||
		preview.ReviewCard.Severity != "requires-runtime-gates" ||
		!preview.ReviewCard.UserReviewRequired ||
		!preview.ReviewCard.RuntimeApprovalNeeded ||
		preview.ReviewCard.PrimaryAction != "Open Compatibility Center" ||
		!containsString(preview.ReviewCard.SecondaryActions, "Review resource access") {
		t.Fatalf("unexpected review card: %#v", preview.ReviewCard)
	}
	if preview.ActionQueue.QueueType != "compatibility-center-request-review-queue" ||
		preview.ActionQueue.ActionCount != 1 ||
		preview.ActionQueue.PendingActionCount != 1 ||
		preview.ActionQueue.UserReviewRequiredCount != 1 ||
		preview.ActionQueue.ExecutionEnabled ||
		preview.ActionQueue.QueuePersisted ||
		preview.ActionQueue.ReviewReceiptRecorded ||
		!containsString(preview.ActionQueue.Actions, "review-launch-request") {
		t.Fatalf("unexpected action queue: %#v", preview.ActionQueue)
	}
	if preview.CompatibilityProfile.ID != "local-compatibility" ||
		preview.CompatibilityProfile.Ready ||
		preview.CompatibilityProfile.LaunchEnabled ||
		preview.CompatibilityProfile.BackendDetailsExposed {
		t.Fatalf("unexpected compatibility profile: %#v", preview.CompatibilityProfile)
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
		t.Fatalf("unexpected execution review state: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.ActionQueueCandidate || !preview.ReviewReceiptRequired ||
		preview.ReviewReceiptRecorded || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.ExecutionRequestPersisted || preview.ActionQueuePersisted ||
		preview.ExecutionStarted || preview.BackendBindingReady ||
		preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution review safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist launch review before Runtime gates pass") ||
		!containsString(preview.BlockedActions, "record launch approval from preview state") ||
		!containsString(preview.BlockedActions, "start compatibility profile from review card") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "applications-games",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.ExecutionReviewPreview(nil)
	if err != nil {
		t.Fatalf("isolated ExecutionReviewPreview returned error: %v", err)
	}
	if isolatedPreview.CompatibilityProfile.ID != "isolated-compatibility" ||
		isolatedPreview.CompatibilityProfile.Kind != "isolated" {
		t.Fatalf("isolated execution review did not recommend isolated compatibility: %#v", isolatedPreview)
	}

	if _, err := plan.ExecutionReviewPreview([]string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionReviewPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution review preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
