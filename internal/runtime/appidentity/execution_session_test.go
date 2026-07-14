package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionSessionPreviewPlansDesktopIdentityWithoutStarting(t *testing.T) {
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

	preview, err := plan.ExecutionSessionPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionSessionPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.session_identity.v1" ||
		preview.RequestType != "execution-session-preview" ||
		preview.SessionType != "compatibility-execution-session" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-transaction-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionSessionPreview" {
		t.Fatalf("unexpected execution session schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution session identity: %#v", preview)
	}
	if preview.Transaction.RequestType != "execution-transaction-preview" ||
		preview.Transaction.TransactionType != "compatibility-launch-transaction" ||
		preview.Transaction.RequestState != "blocked" ||
		preview.Transaction.StepCount != 7 ||
		preview.Transaction.BlockedStepCount != 1 ||
		preview.Transaction.TransactionCommitted ||
		preview.Transaction.RuntimeLaunchApproval ||
		preview.Transaction.LaunchAllowed ||
		preview.Transaction.ExecutionStarted ||
		preview.Transaction.BackendBindingCommitted ||
		preview.Transaction.SnapshotBaselineCreated ||
		preview.Transaction.ResourceGrantsCommitted {
		t.Fatalf("unexpected transaction summary: %#v", preview.Transaction)
	}
	if preview.WindowIdentity.SchemaVersion != "xnix.runtime.window_identity.v1" ||
		preview.WindowIdentity.WindowKind != "compatibility-application" ||
		preview.WindowIdentity.ClassGroup != "xnix-compatibility" ||
		preview.WindowIdentity.ResourceName != "org.example.ledger" ||
		preview.WindowIdentity.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.WindowIdentity.TitleHint != "Example Ledger" ||
		preview.WindowIdentity.WindowObserved ||
		preview.WindowIdentity.WindowRegistration != "planned" ||
		preview.WindowIdentity.BackendDetailsExposed {
		t.Fatalf("unexpected window identity summary: %#v", preview.WindowIdentity)
	}
	if preview.TaskManager.GroupingKey != "org.example.ledger" ||
		!preview.TaskManager.PinningAllowed ||
		!preview.TaskManager.RestoreAllowed ||
		preview.TaskManager.SkipTaskbar ||
		!preview.TaskManager.ShowInSwitcher ||
		!preview.TaskManager.PreferExistingWindow ||
		!preview.TaskManager.EntryPlanned ||
		preview.TaskManager.EntryActive {
		t.Fatalf("unexpected task manager summary: %#v", preview.TaskManager)
	}
	if preview.KWin.ScriptRole != "identity-and-layout" ||
		preview.KWin.ResourceName != "org.example.ledger" ||
		preview.KWin.ClassGroup != "xnix-compatibility" ||
		preview.KWin.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.KWin.TaskManagerGroupingKey != "org.example.ledger" ||
		preview.KWin.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.KWin.Placement != "normal-window" ||
		!preview.KWin.WindowManagerPolicyOnly ||
		!preview.KWin.RulePlanned ||
		preview.KWin.RuleApplied {
		t.Fatalf("unexpected KWin summary: %#v", preview.KWin)
	}
	if preview.Tray.StatusType != "tray-status-preview" ||
		preview.Tray.RegisteredApplicationCount != 1 ||
		preview.Tray.ActiveApplicationCount != 0 ||
		preview.Tray.AttentionRequiredCount != 0 ||
		preview.Tray.CompatibilityState != "ready" ||
		preview.Tray.TrayBridgeState != "planned" ||
		!preview.Tray.EntryPlanned ||
		preview.Tray.LiveBridgeEnabled ||
		preview.Tray.BridgeConfigurationPersisted ||
		preview.Tray.BackendDetailsExposed {
		t.Fatalf("unexpected tray summary: %#v", preview.Tray)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		!preview.SessionPlanCreated || preview.SessionCreated ||
		preview.SessionRegistered || preview.WindowObserved ||
		!preview.TaskManagerEntryPlanned || preview.TaskManagerEntryActive ||
		!preview.KWinRulePlanned || preview.KWinRuleApplied ||
		!preview.TrayEntryPlanned || preview.LiveTrayBridgeEnabled ||
		preview.TransactionCommitted || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted ||
		preview.RequestObjectsCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution session safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "create execution session from preview") ||
		!containsString(preview.BlockedActions, "register live window before Runtime launch") ||
		!containsString(preview.BlockedActions, "activate task manager entry from session preview") ||
		!containsString(preview.BlockedActions, "enable live tray bridge before session exists") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionSessionPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionSessionPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.Transaction.BlockedStepCount != 2 ||
		rejectedPreview.SessionCreated ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected session preview: %#v", rejectedPreview)
	}

	if _, err := plan.ExecutionSessionPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionSessionPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionSessionPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionSessionPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution session preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
