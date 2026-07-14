package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionQueuePreviewAggregatesSevenEntryActions(t *testing.T) {
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

	preview, err := plan.KDEActionQueuePreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionQueuePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_queue.v1" ||
		preview.RequestType != "kde-action-queue-preview" ||
		preview.QueueType != "compatibility-center-kde-action-queue" ||
		preview.Source != "kde-entrypoint-action-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetKDEActionQueuePreview" {
		t.Fatalf("unexpected KDE action queue schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE action queue identity: %#v", preview)
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
	expectedIDs := []string{
		"review-launcher-action",
		"review-task-manager-action",
		"review-file-manager-action",
		"review-tray-status-action",
		"review-notification-action",
		"review-compatibility-center-action",
		"review-settings-action",
	}
	if strings.Join(preview.ActionIDs, ",") != strings.Join(expectedIDs, ",") ||
		preview.ActionCount != 7 ||
		preview.PendingActionCount != 7 ||
		preview.UserReviewRequiredCount != 4 ||
		preview.PortalActionCount != 1 ||
		preview.RuntimeGateActionCount != 7 {
		t.Fatalf("unexpected action queue counts: %#v", preview)
	}
	for _, action := range preview.Actions {
		if action.SourceType != "kde-entrypoint-action-preview" ||
			action.Status == "pass" ||
			!action.RequiresRuntimeGate ||
			!action.BlockedByRuntimeGate ||
			!action.OpensCompatibilityCenter ||
			action.ExecutionEnabled ||
			action.RequestObjectCreated ||
			action.BackendProcessStarted ||
			action.HostRootModified ||
			action.BackendDetailsExposed {
			t.Fatalf("unexpected queued action safety flags for %s: %#v", action.ID, action)
		}
	}
	fileManager := findKDEActionQueueItem(preview.Actions, "review-file-manager-action")
	if fileManager.EntryPointID != "file-manager" ||
		fileManager.KDEComponent != "Dolphin" ||
		fileManager.Intent != "open-files" ||
		fileManager.Status != "portal-review-required" ||
		fileManager.Priority != "high" ||
		!fileManager.RequiresPortal ||
		!fileManager.UserReviewRequired ||
		fileManager.RuntimeGate != "portal-file-open-review" {
		t.Fatalf("unexpected file manager queue item: %#v", fileManager)
	}
	settings := findKDEActionQueueItem(preview.Actions, "review-settings-action")
	if settings.EntryPointID != "settings" ||
		settings.Status != "settings-review-required" ||
		!settings.OpensSettings ||
		!settings.UserReviewRequired ||
		settings.RuntimeGate != "runtime-settings-review" {
		t.Fatalf("unexpected settings queue item: %#v", settings)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.ActionQueueCreated ||
		preview.ActionQueuePersisted || preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE action queue safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE action queue from preview") ||
		!containsString(preview.BlockedActions, "execute queued KDE actions without Runtime approval") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from queued KDE actions") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEActionQueuePreviewValidationAndSafeText(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionQueuePreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionQueuePreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.SessionStatus.SessionState != "review-declined" ||
		rejectedPreview.SessionStatus.BlockedGateCount != 2 ||
		rejectedPreview.LaunchAllowed ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action queue preview: %#v", rejectedPreview)
	}

	if _, err := plan.KDEActionQueuePreview("invalid", nil); err == nil {
		t.Fatalf("KDEActionQueuePreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionQueuePreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionQueuePreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionQueuePreview("approved", nil)
	if err != nil {
		t.Fatalf("KDEActionQueuePreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action queue preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func findKDEActionQueueItem(actions []KDEActionQueueItem, id string) KDEActionQueueItem {
	for _, action := range actions {
		if action.ID == id {
			return action
		}
	}
	return KDEActionQueueItem{}
}
