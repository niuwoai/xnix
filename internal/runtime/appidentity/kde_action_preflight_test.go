package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionPreflightPreviewKeepsApprovedActionGated(t *testing.T) {
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

	preview, err := plan.KDEActionPreflightPreview("review-file-manager-action", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionPreflightPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_preflight.v1" ||
		preview.RequestType != "kde-action-preflight-preview" ||
		preview.PreflightType != "compatibility-center-kde-action-preflight" ||
		preview.Source != "kde-action-review-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "PreflightKDEAction" ||
		preview.ReadMethod != "GetKDEActionPreflightPreview" {
		t.Fatalf("unexpected KDE action preflight schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action preflight identity: %#v", preview)
	}
	if preview.Review.RequestType != "kde-action-review-preview" ||
		preview.Review.ReviewType != "compatibility-center-kde-action-review" ||
		preview.Review.Decision != "approved" ||
		!preview.Review.DecisionAccepted ||
		!preview.Review.UserIntentCaptured ||
		preview.Review.DecisionRecorded ||
		preview.Review.RuntimeApprovalGranted ||
		preview.Review.ExecutionAllowed ||
		preview.Review.ReviewReceiptCreated ||
		preview.Review.QueueStateChanged ||
		preview.Review.PermissionGrantCreated {
		t.Fatalf("unexpected review summary: %#v", preview.Review)
	}
	if preview.Action.ID != "review-file-manager-action" ||
		preview.Action.EntryPointID != "file-manager" ||
		preview.Action.KDEComponent != "Dolphin" ||
		preview.Action.RuntimeGate != "portal-file-open-review" ||
		!preview.Action.RequiresPortal ||
		!preview.Action.RequiresRuntimeGate ||
		preview.Action.ExecutionEnabled ||
		preview.Action.RequestObjectCreated ||
		preview.Action.BackendProcessStarted ||
		preview.Action.HostRootModified ||
		preview.Action.BackendDetailsExposed {
		t.Fatalf("unexpected action summary: %#v", preview.Action)
	}
	if preview.ExecutionPreflight.RequestType != "execution-preflight-preview" ||
		preview.ExecutionPreflight.PreflightType != "compatibility-launch-preflight" ||
		preview.ExecutionPreflight.CheckCount != 5 ||
		preview.ExecutionPreflight.PassedCheckCount != 1 ||
		preview.ExecutionPreflight.RequiredCheckCount != 2 ||
		preview.ExecutionPreflight.PendingCheckCount != 1 ||
		preview.ExecutionPreflight.BlockedCheckCount != 1 ||
		preview.ExecutionPreflight.PreflightComplete ||
		preview.ExecutionPreflight.PreflightPassed ||
		preview.ExecutionPreflight.RuntimeLaunchApproval ||
		preview.ExecutionPreflight.LaunchAllowed ||
		preview.ExecutionPreflight.ExecutionStarted ||
		preview.ExecutionPreflight.BackendDetailsExposed {
		t.Fatalf("unexpected execution preflight summary: %#v", preview.ExecutionPreflight)
	}
	if preview.CheckCount != 5 ||
		preview.PassedCheckCount != 1 ||
		preview.RequiredCheckCount != 1 ||
		preview.PendingCheckCount != 1 ||
		preview.BlockedCheckCount != 2 {
		t.Fatalf("unexpected action preflight counts: %#v", preview)
	}
	gateIDs := []string{}
	for _, gate := range preview.PreflightChecks {
		gateIDs = append(gateIDs, gate.ID)
	}
	expectedGateIDs := []string{"kde-action-review", "portal-policy-review", "review-receipt", "request-object", "runtime-launch-write-gate"}
	for index, expected := range expectedGateIDs {
		if gateIDs[index] != expected {
			t.Fatalf("unexpected action preflight gates: %#v", preview.PreflightChecks)
		}
	}
	if preview.PreflightChecks[1].Status != "required" ||
		preview.PreflightChecks[3].Status != "blocked" ||
		preview.PreflightChecks[4].Status != "blocked" {
		t.Fatalf("unexpected action preflight gate states: %#v", preview.PreflightChecks)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionReviewCaptured ||
		!preview.ActionPreflightCreated || !preview.ReviewAllowsPreflight ||
		preview.PreflightComplete || preview.PreflightPassed ||
		preview.PortalPreflightReady || !preview.ReviewReceiptRequired ||
		preview.ReviewReceiptRecorded || preview.ActionQueuePersisted ||
		preview.QueueStateChanged || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.PermissionGrantCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action preflight safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "treat KDE action preflight as Runtime execution approval") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from action preflight preview") ||
		!containsString(preview.BlockedActions, "grant desktop resources from action preflight preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionPreflightPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionPreflightPreview("review-launcher-action", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionPreflightPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.ReviewAllowsPreflight ||
		rejectedPreview.PreflightChecks[0].Status != "blocked" ||
		rejectedPreview.SessionStatus.SessionState != "review-declined" ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action preflight preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionPreflightPreview("review-settings-action", "deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionPreflightPreview returned error: %v", err)
	}
	if deferredPreview.ReviewAllowsPreflight ||
		deferredPreview.PreflightChecks[0].Status != "pending" ||
		deferredPreview.Review.DecisionRecorded ||
		deferredPreview.ReviewReceiptRecorded ||
		deferredPreview.QueueStateChanged {
		t.Fatalf("unexpected deferred action preflight preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionPreflightPreview("missing-action", "approved", nil); err == nil {
		t.Fatalf("KDEActionPreflightPreview accepted an unknown action")
	}
	if _, err := plan.KDEActionPreflightPreview("review-launcher-action\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEActionPreflightPreview accepted a multiline action id")
	}
	if _, err := plan.KDEActionPreflightPreview("review-launcher-action", "invalid", nil); err == nil {
		t.Fatalf("KDEActionPreflightPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionPreflightPreview("review-file-manager-action", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionPreflightPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionPreflightPreview("review-settings-action", "reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionPreflightPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action preflight preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
