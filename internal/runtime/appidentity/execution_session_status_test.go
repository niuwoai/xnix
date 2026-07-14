package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionSessionStatusPreviewKeepsLiveStateReadOnly(t *testing.T) {
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

	preview, err := plan.ExecutionSessionStatusPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionSessionStatusPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.session_status.v1" ||
		preview.RequestType != "execution-session-status-preview" ||
		preview.StatusType != "compatibility-session-status" ||
		preview.SessionType != "compatibility-execution-session" ||
		preview.RequestState != "blocked" ||
		preview.SessionState != "planned-blocked" ||
		preview.Source != "execution-session-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionSessionStatusPreview" {
		t.Fatalf("unexpected execution session status schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution session status identity: %#v", preview)
	}
	if preview.Session.SessionType != "compatibility-execution-session" ||
		preview.Session.SessionState != "planned-blocked" ||
		preview.Session.TransactionState != "blocked" ||
		preview.Session.TransactionStepCount != 7 ||
		preview.Session.BlockedTransactionSteps != 1 ||
		preview.Session.WindowRegistration != "planned" ||
		preview.Session.DesktopSurfaceState != "planned" ||
		preview.Session.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		preview.Session.TaskManagerState != "planned" ||
		preview.Session.TrayState != "planned" ||
		!preview.Session.SessionPlanCreated ||
		preview.Session.SessionCreated ||
		preview.Session.SessionRegistered ||
		preview.Session.SessionActive ||
		preview.Session.LiveStateObserved ||
		preview.Session.StatusPersisted ||
		preview.Session.RuntimeLaunchApproval ||
		preview.Session.LaunchAllowed ||
		preview.Session.ExecutionStarted ||
		preview.Session.BackendProcessStarted ||
		preview.Session.BackendDetailsExposed {
		t.Fatalf("unexpected session status summary: %#v", preview.Session)
	}
	if preview.DesktopSurface.WindowKind != "compatibility-application" ||
		preview.DesktopSurface.ClassGroup != "xnix-compatibility" ||
		preview.DesktopSurface.ResourceName != "org.example.ledger" ||
		preview.DesktopSurface.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.DesktopSurface.WindowState != "not-observed" ||
		preview.DesktopSurface.TaskManagerGroupingKey != "org.example.ledger" ||
		preview.DesktopSurface.TaskManagerState != "planned" ||
		preview.DesktopSurface.KWinState != "planned" ||
		preview.DesktopSurface.TrayState != "planned" ||
		preview.DesktopSurface.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		preview.DesktopSurface.WindowObserved ||
		!preview.DesktopSurface.TaskManagerEntryPlanned ||
		preview.DesktopSurface.TaskManagerEntryActive ||
		!preview.DesktopSurface.KWinRulePlanned ||
		preview.DesktopSurface.KWinRuleApplied ||
		!preview.DesktopSurface.TrayEntryPlanned ||
		preview.DesktopSurface.LiveTrayBridgeEnabled {
		t.Fatalf("unexpected desktop surface status: %#v", preview.DesktopSurface)
	}
	if preview.UserVisibleState.PrimaryLabel != "Example Ledger" ||
		preview.UserVisibleState.SecondaryLabel != "Waiting for Runtime gates" ||
		preview.UserVisibleState.TaskbarBadge != "Planned" ||
		preview.UserVisibleState.TrayLabel != "Ready for review" ||
		preview.UserVisibleState.CompatibilityCenterStatus != "Runtime gates required" ||
		preview.UserVisibleState.NextUserAction != "Open Compatibility Center" ||
		preview.UserVisibleState.UserFacingMode != "Automatic" ||
		preview.UserVisibleState.UserFacingAccess != "Review required" ||
		preview.UserVisibleState.BackendDetailsExposed {
		t.Fatalf("unexpected user-visible status: %#v", preview.UserVisibleState)
	}
	if preview.GateCount != 5 ||
		preview.PassedGateCount != 2 ||
		preview.RequiredGateCount != 1 ||
		preview.PendingGateCount != 2 ||
		preview.BlockedGateCount != 1 {
		t.Fatalf("unexpected gate counts: %#v", preview)
	}
	if got := gateIDs(preview.Gates); strings.Join(got, ",") != "session-identity,user-decision,runtime-launch-approval,live-window-observation,desktop-surface-activation" {
		t.Fatalf("unexpected gate order: %#v", got)
	}
	if preview.Gates[2].Status != "blocked" ||
		!preview.Gates[2].Required ||
		!preview.Gates[2].BlocksLiveSession ||
		!preview.Gates[2].RuntimeGateRequired {
		t.Fatalf("unexpected runtime launch gate: %#v", preview.Gates[2])
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		!preview.StatusReadModelCreated || !preview.SessionPlanCreated ||
		preview.SessionCreated || preview.SessionRegistered || preview.SessionActive ||
		preview.LiveStateObserved || preview.StatusPersisted ||
		preview.WindowObserved || !preview.TaskManagerEntryPlanned ||
		preview.TaskManagerEntryActive || !preview.KWinRulePlanned ||
		preview.KWinRuleApplied || !preview.TrayEntryPlanned ||
		preview.LiveTrayBridgeEnabled || preview.TransactionCommitted ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution session status safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist live session status from preview") ||
		!containsString(preview.BlockedActions, "mark task manager entry active from status preview") ||
		!containsString(preview.BlockedActions, "observe windows before Runtime launch") ||
		!containsString(preview.BlockedActions, "enable live tray bridge before session exists") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionSessionStatusPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionSessionStatusPreview returned error: %v", err)
	}
	if rejectedPreview.SessionState != "review-declined" ||
		rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.RequiredGateCount != 2 ||
		rejectedPreview.BlockedGateCount != 2 ||
		rejectedPreview.SessionActive ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected session status preview: %#v", rejectedPreview)
	}

	if _, err := plan.ExecutionSessionStatusPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionSessionStatusPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionSessionStatusPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionSessionStatusPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution session status preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func gateIDs(gates []ExecutionSessionStatusGate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		ids = append(ids, gate.ID)
	}
	return ids
}
