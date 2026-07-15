package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionStatusPreviewShowsWaitingStateWithoutExecution(t *testing.T) {
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

	preview, err := plan.KDEActionStatusPreview("review-file-manager-action", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionStatusPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_status.v1" ||
		preview.RequestType != "kde-action-status-preview" ||
		preview.StatusType != "compatibility-center-kde-action-status" ||
		preview.StatusState != "waiting-for-runtime-gates" ||
		preview.Source != "kde-action-receipt-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEActionStatus" ||
		preview.ReadMethod != "GetKDEActionStatusPreview" {
		t.Fatalf("unexpected KDE action status schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action status identity: %#v", preview)
	}
	if preview.Receipt.RequestType != "kde-action-receipt-preview" ||
		preview.Receipt.ReceiptType != "compatibility-center-kde-action-review-receipt" ||
		preview.Receipt.ReceiptID != "compat-review-org.example.ledger-review-file-manager-action-approved" ||
		preview.Receipt.ReceiptState != "preview-only" ||
		preview.Receipt.Decision != "approved" ||
		!preview.Receipt.ReceiptPreviewCreated ||
		!preview.Receipt.ReceiptRecordable ||
		preview.Receipt.DecisionRecorded ||
		preview.Receipt.ReviewReceiptRecorded ||
		preview.Receipt.RuntimeLaunchApproval ||
		preview.Receipt.LaunchAllowed ||
		preview.Receipt.ExecutionStarted ||
		preview.Receipt.BackendDetailsExposed {
		t.Fatalf("unexpected receipt summary: %#v", preview.Receipt)
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
		preview.Preflight.ReceiptGateStatus != "pending" ||
		preview.Preflight.CheckCount != 5 ||
		preview.Preflight.PreflightComplete ||
		preview.Preflight.ReviewReceiptRecorded ||
		preview.Preflight.ExecutionStarted {
		t.Fatalf("unexpected preflight summary: %#v", preview.Preflight)
	}
	if preview.UserVisibleState.PrimaryLabel != "Example Ledger" ||
		preview.UserVisibleState.SecondaryLabel != "Waiting for portal-file-open-review" ||
		preview.UserVisibleState.Badge != "Waiting" ||
		preview.UserVisibleState.ActionStateLabel != "Runtime gates required" ||
		preview.UserVisibleState.CompatibilityCenterStatus != "waiting-for-runtime-gates" ||
		preview.UserVisibleState.NextUserAction != "Wait for Runtime gates" ||
		preview.UserVisibleState.BackendDetailsExposed {
		t.Fatalf("unexpected user-visible state: %#v", preview.UserVisibleState)
	}
	if preview.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		preview.NotificationIntent != "show-kde-action-status-waiting" ||
		preview.RequiredRuntimeGate != "portal-file-open-review" ||
		!strings.Contains(preview.NextStep, "portal-file-open-review") {
		t.Fatalf("unexpected status routing: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionReviewCaptured ||
		!preview.ActionPreflightCreated || !preview.ReceiptPreviewCreated ||
		!preview.StatusPreviewCreated || preview.StatusPersisted ||
		preview.DecisionRecorded || preview.ReviewReceiptRecorded ||
		preview.ActionQueuePersisted || preview.QueueStateChanged ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.ResourceGrantCreated || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.PermissionGrantCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action status safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE action status from preview state") ||
		!containsString(preview.BlockedActions, "treat KDE action status as Runtime execution approval") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from status preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionStatusPreviewValidationAndDecisionStates(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionStatusPreview("review-launcher-action", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionStatusPreview returned error: %v", err)
	}
	if rejectedPreview.StatusState != "review-rejected" ||
		rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.UserVisibleState.Badge != "Rejected" ||
		rejectedPreview.CompatibilityCenterState != "action-rejected" ||
		rejectedPreview.NotificationIntent != "show-kde-action-status-rejected" ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action status preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.KDEActionStatusPreview("review-settings-action", "deferred", nil)
	if err != nil {
		t.Fatalf("deferred KDEActionStatusPreview returned error: %v", err)
	}
	if deferredPreview.StatusState != "review-deferred" ||
		deferredPreview.UserDecisionAllowsLaunch ||
		deferredPreview.UserVisibleState.Badge != "Deferred" ||
		deferredPreview.CompatibilityCenterState != "action-deferred" ||
		deferredPreview.NotificationIntent != "show-kde-action-status-deferred" ||
		deferredPreview.StatusPersisted ||
		deferredPreview.ReviewReceiptRecorded ||
		deferredPreview.QueueStateChanged {
		t.Fatalf("unexpected deferred action status preview: %#v", deferredPreview)
	}

	if _, err := plan.KDEActionStatusPreview("missing-action", "approved", nil); err == nil {
		t.Fatalf("KDEActionStatusPreview accepted an unknown action")
	}
	if _, err := plan.KDEActionStatusPreview("review-launcher-action\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEActionStatusPreview accepted a multiline action id")
	}
	if _, err := plan.KDEActionStatusPreview("review-launcher-action", "invalid", nil); err == nil {
		t.Fatalf("KDEActionStatusPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionStatusPreview("review-file-manager-action", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionStatusPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionStatusPreview("review-settings-action", "reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionStatusPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action status preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
