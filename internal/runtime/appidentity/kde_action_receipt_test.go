package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionReceiptPreviewShapesReceiptWithoutRecording(t *testing.T) {
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

	preview, err := plan.KDEActionReceiptPreview("review-file-manager-action", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionReceiptPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_receipt.v1" ||
		preview.RequestType != "kde-action-receipt-preview" ||
		preview.ReceiptType != "compatibility-center-kde-action-review-receipt" ||
		preview.ReceiptID != "compat-review-org.example.ledger-review-file-manager-action-approved" ||
		preview.ReceiptState != "preview-only" ||
		preview.Source != "kde-action-preflight-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "RecordKDEActionReviewReceipt" ||
		preview.ReadMethod != "GetKDEActionReceiptPreview" {
		t.Fatalf("unexpected KDE action receipt schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action receipt identity: %#v", preview)
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
		preview.Action.ExecutionEnabled ||
		preview.Action.RequestObjectCreated ||
		preview.Action.BackendProcessStarted ||
		preview.Action.HostRootModified ||
		preview.Action.BackendDetailsExposed {
		t.Fatalf("unexpected action summary: %#v", preview.Action)
	}
	if preview.Preflight.RequestType != "kde-action-preflight-preview" ||
		preview.Preflight.PreflightType != "compatibility-center-kde-action-preflight" ||
		preview.Preflight.ReceiptGateStatus != "pending" ||
		preview.Preflight.CheckCount != 5 ||
		preview.Preflight.PassedCheckCount != 1 ||
		preview.Preflight.RequiredCheckCount != 1 ||
		preview.Preflight.PendingCheckCount != 1 ||
		preview.Preflight.BlockedCheckCount != 2 ||
		preview.Preflight.PreflightComplete ||
		preview.Preflight.PreflightPassed ||
		!preview.Preflight.ReviewReceiptRequired ||
		preview.Preflight.ReviewReceiptRecorded ||
		preview.Preflight.RuntimeLaunchApproval ||
		preview.Preflight.LaunchAllowed ||
		preview.Preflight.ExecutionStarted ||
		preview.Preflight.BackendDetailsExposed {
		t.Fatalf("unexpected preflight summary: %#v", preview.Preflight)
	}
	if preview.FieldCount != 5 || len(preview.ReceiptFields) != 5 {
		t.Fatalf("unexpected receipt fields: %#v", preview.ReceiptFields)
	}
	expectedFieldIDs := []string{"application-id", "action-id", "decision", "required-runtime-gate", "execution-disabled"}
	for index, expected := range expectedFieldIDs {
		if preview.ReceiptFields[index].ID != expected ||
			!preview.ReceiptFields[index].Required ||
			preview.ReceiptFields[index].Recorded {
			t.Fatalf("unexpected receipt field at %d: %#v", index, preview.ReceiptFields[index])
		}
	}
	if preview.RequiredRuntimeGate != "portal-file-open-review" ||
		!strings.Contains(preview.NextStep, "portal-file-open-review") {
		t.Fatalf("unexpected receipt next step: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionReviewCaptured ||
		!preview.ActionPreflightCreated || !preview.ReceiptPreviewCreated ||
		!preview.ReceiptRecordable || preview.DecisionRecorded ||
		preview.ReviewReceiptCreated || preview.ReviewReceiptRecorded ||
		preview.ActionQueuePersisted || preview.QueueStateChanged ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.ResourceGrantCreated || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.PermissionGrantCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action receipt safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "record KDE action receipt from preview state") ||
		!containsString(preview.BlockedActions, "treat KDE action receipt as Runtime execution approval") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from receipt preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionReceiptPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionReceiptPreview("review-launcher-action", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionReceiptPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		!strings.Contains(rejectedPreview.NextStep, "rejection") ||
		rejectedPreview.ReviewReceiptRecorded ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action receipt preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionReceiptPreview("review-settings-action", "deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionReceiptPreview returned error: %v", err)
	}
	if deferredPreview.UserDecisionAllowsLaunch ||
		!strings.Contains(deferredPreview.NextStep, "deferral") ||
		deferredPreview.Review.DecisionRecorded ||
		deferredPreview.ReviewReceiptRecorded ||
		deferredPreview.QueueStateChanged {
		t.Fatalf("unexpected deferred action receipt preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionReceiptPreview("missing-action", "approved", nil); err == nil {
		t.Fatalf("KDEActionReceiptPreview accepted an unknown action")
	}
	if _, err := plan.KDEActionReceiptPreview("review-launcher-action\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEActionReceiptPreview accepted a multiline action id")
	}
	if _, err := plan.KDEActionReceiptPreview("review-launcher-action", "invalid", nil); err == nil {
		t.Fatalf("KDEActionReceiptPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionReceiptPreview("review-file-manager-action", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionReceiptPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionReceiptPreview("review-settings-action", "reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionReceiptPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action receipt preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
