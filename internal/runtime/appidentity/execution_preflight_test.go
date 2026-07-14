package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionPreflightPreviewBlocksLaunchUntilGatesPass(t *testing.T) {
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

	preview, err := plan.ExecutionPreflightPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionPreflightPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.launch_preflight.v1" ||
		preview.RequestType != "execution-preflight-preview" ||
		preview.PreflightType != "compatibility-launch-preflight" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-decision-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionPreflightPreview" {
		t.Fatalf("unexpected execution preflight schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution preflight identity: %#v", preview)
	}
	if preview.ExecutionDecision.SchemaVersion != "xnix.runtime.request_decision.v1" ||
		preview.ExecutionDecision.RequestType != "execution-decision-preview" ||
		preview.ExecutionDecision.Decision != "approved" ||
		!preview.ExecutionDecision.DecisionAccepted ||
		!preview.ExecutionDecision.UserIntentCaptured ||
		preview.ExecutionDecision.DecisionRecorded ||
		preview.ExecutionDecision.RuntimeApprovalGranted ||
		preview.ExecutionDecision.ExecutionAllowed ||
		preview.ExecutionDecision.ReviewReceiptCreated ||
		preview.ExecutionDecision.QueueStateChanged {
		t.Fatalf("unexpected execution decision summary: %#v", preview.ExecutionDecision)
	}
	if preview.CheckCount != 5 ||
		preview.PassedCheckCount != 1 ||
		preview.RequiredCheckCount != 2 ||
		preview.PendingCheckCount != 1 ||
		preview.BlockedCheckCount != 1 {
		t.Fatalf("unexpected preflight counts: %#v", preview)
	}
	if len(preview.PreflightChecks) != 5 ||
		preview.PreflightChecks[0].ID != "user-decision" ||
		preview.PreflightChecks[0].Status != "pass" ||
		preview.PreflightChecks[1].ID != "portal-policy-review" ||
		preview.PreflightChecks[1].Status != "required" ||
		preview.PreflightChecks[2].ID != "snapshot-baseline" ||
		preview.PreflightChecks[3].ID != "backend-binding" ||
		preview.PreflightChecks[4].ID != "runtime-launch-write-gate" ||
		preview.PreflightChecks[4].Status != "blocked" {
		t.Fatalf("unexpected preflight checks: %#v", preview.PreflightChecks)
	}
	if preview.ExecutionState != "blocked" ||
		preview.OverallStatus != "not-ready" ||
		preview.WriteGateDecision != "blocked-until-production-backend" ||
		preview.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		!preview.PortalRequired ||
		!preview.SnapshotRequired ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution preflight state: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		preview.PreflightComplete || preview.PreflightPassed ||
		preview.PortalPreflightReady || preview.SnapshotPreflightReady ||
		preview.BackendPreflightReady || preview.WriteGateOpen ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.ExecutionRequestPersisted || preview.ActionQueuePersisted ||
		preview.ReviewReceiptRecorded || preview.ExecutionStarted ||
		preview.BackendBindingReady || preview.RequestObjectCreated ||
		preview.PermissionGranted || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution preflight safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "start compatibility profile before preflight passes") ||
		!containsString(preview.BlockedActions, "grant desktop resources before Portal review") ||
		!containsString(preview.BlockedActions, "bind compatibility profile from preflight preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionPreflightPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionPreflightPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.PreflightChecks[0].Status != "blocked" ||
		rejectedPreview.BlockedCheckCount != 2 {
		t.Fatalf("unexpected rejected preflight: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.ExecutionPreflightPreview("deferred", nil)
	if err != nil {
		t.Fatalf("deferred ExecutionPreflightPreview returned error: %v", err)
	}
	if deferredPreview.UserDecisionAllowsLaunch ||
		deferredPreview.PreflightChecks[0].Status != "pending" ||
		deferredPreview.PendingCheckCount != 2 {
		t.Fatalf("unexpected deferred preflight: %#v", deferredPreview)
	}

	if _, err := plan.ExecutionPreflightPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionPreflightPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionPreflightPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionPreflightPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution preflight preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
