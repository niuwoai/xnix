package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEEntryPointActionPreviewMapsFileManagerWithoutSideEffects(t *testing.T) {
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

	preview, err := plan.KDEEntryPointActionPreview("file-manager", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEEntryPointActionPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_entrypoint_action.v1" ||
		preview.RequestType != "kde-entrypoint-action-preview" ||
		preview.ActionType != "kde-entrypoint-action" ||
		preview.Source != "kde-entrypoints-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.ReadMethod != "GetKDEEntryPointActionPreview" {
		t.Fatalf("unexpected KDE entrypoint action schema: %#v", preview)
	}
	if preview.EntryPointID != "file-manager" ||
		preview.EntryPointLabel != "File manager" ||
		preview.KDEComponent != "Dolphin" ||
		preview.RuntimeSource != "file-open-preview" ||
		preview.RuntimeMethod != "Launch" {
		t.Fatalf("unexpected KDE entrypoint action routing: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE entrypoint action identity: %#v", preview)
	}
	if preview.Action.Intent != "open-files" ||
		preview.Action.UserAction != "Open selected files through Runtime" ||
		preview.Action.SafeResult != "show file access review" ||
		!preview.Action.RequiresPortal ||
		!preview.Action.RequiresRuntimeGate ||
		!preview.Action.BlockedByRuntimeGate ||
		!preview.Action.OpensCompatibilityCenter ||
		preview.Action.OpensSettings ||
		preview.Action.CreatesRequestObject ||
		preview.Action.StartsBackend ||
		preview.Action.MutatesHost ||
		preview.Action.BackendDetailsExposed {
		t.Fatalf("unexpected action summary: %#v", preview.Action)
	}
	if preview.EntryPoint.ID != "file-manager" ||
		preview.EntryPoint.Label != "File manager" ||
		preview.EntryPoint.KDEComponent != "Dolphin" ||
		!preview.EntryPoint.Visible ||
		!preview.EntryPoint.Planned ||
		preview.EntryPoint.Active ||
		!preview.EntryPoint.RequiresPortal ||
		!preview.EntryPoint.RequiresRuntimeGate ||
		!preview.EntryPoint.BlockedByRuntimeGate ||
		preview.EntryPoint.WritesHost ||
		preview.EntryPoint.StartsBackend ||
		preview.EntryPoint.BackendDetailsExposed {
		t.Fatalf("unexpected entrypoint summary: %#v", preview.EntryPoint)
	}
	if preview.SessionStatus.RequestType != "execution-session-status-preview" ||
		preview.SessionStatus.StatusType != "compatibility-session-status" ||
		preview.SessionStatus.SessionState != "planned-blocked" ||
		preview.SessionStatus.GateCount != 5 ||
		preview.SessionStatus.BlockedGateCount != 1 ||
		preview.SessionStatus.DesktopSurfaceState != "planned" ||
		preview.SessionStatus.UserVisibleState != "Runtime gates required" ||
		preview.SessionStatus.RuntimeLaunchApproval ||
		preview.SessionStatus.LaunchAllowed ||
		preview.SessionStatus.ExecutionStarted {
		t.Fatalf("unexpected session status summary: %#v", preview.SessionStatus)
	}
	if preview.GateSummary.GateCount != 5 ||
		preview.GateSummary.BlockedGateCount != 1 ||
		!preview.GateSummary.RequiresRuntimeGate ||
		!preview.GateSummary.BlockedByRuntimeGate ||
		preview.GateSummary.RuntimeLaunchApproval ||
		preview.GateSummary.LaunchAllowed ||
		preview.GateSummary.ExecutionStarted ||
		preview.GateSummary.RequestObjectCreated ||
		preview.GateSummary.BackendProcessStarted ||
		preview.GateSummary.BackendDetailsExposed {
		t.Fatalf("unexpected gate summary: %#v", preview.GateSummary)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.StableDesktopContract ||
		!preview.NormalApplicationSurface || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.EntryPointActionCaptured ||
		!preview.EntryPointPlanCreated || preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "commit KDE entrypoint action from preview") ||
		!containsString(preview.BlockedActions, "create Runtime request object from entrypoint action preview") ||
		!containsString(preview.BlockedActions, "start compatibility profile from entrypoint action preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDEEntryPointActionPreviewCoversEntryActionsAndValidation(t *testing.T) {
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

	expected := map[string]struct {
		intent        string
		safeResult    string
		opensSettings bool
	}{
		"launcher":             {"open-application", "show Compatibility Center launch review", false},
		"task-manager":         {"restore-application", "show planned session status", false},
		"file-manager":         {"open-files", "show file access review", false},
		"system-tray":          {"open-status", "show compatibility status", false},
		"notifications":        {"open-review-card", "show review details", false},
		"compatibility-center": {"open-compatibility-center", "show application compatibility card", false},
		"settings":             {"open-settings", "show unified settings", true},
	}
	for entryPointID, expectation := range expected {
		preview, err := plan.KDEEntryPointActionPreview(entryPointID, "approved", nil)
		if err != nil {
			t.Fatalf("KDEEntryPointActionPreview(%s) returned error: %v", entryPointID, err)
		}
		if preview.Action.Intent != expectation.intent ||
			preview.Action.SafeResult != expectation.safeResult ||
			preview.Action.OpensSettings != expectation.opensSettings ||
			!preview.Action.OpensCompatibilityCenter ||
			preview.Action.CreatesRequestObject ||
			preview.Action.StartsBackend ||
			preview.Action.MutatesHost {
			t.Fatalf("unexpected action for %s: %#v", entryPointID, preview.Action)
		}
	}

	rejectedPreview, err := plan.KDEEntryPointActionPreview("launcher", "rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEEntryPointActionPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.SessionStatus.SessionState != "review-declined" ||
		rejectedPreview.GateSummary.BlockedGateCount != 2 ||
		rejectedPreview.LaunchAllowed ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected action preview: %#v", rejectedPreview)
	}

	if _, err := plan.KDEEntryPointActionPreview("unsupported", "approved", nil); err == nil {
		t.Fatalf("KDEEntryPointActionPreview accepted an unsupported entrypoint")
	}
	if _, err := plan.KDEEntryPointActionPreview("launcher\nbad", "approved", nil); err == nil {
		t.Fatalf("KDEEntryPointActionPreview accepted a multiline entrypoint")
	}
	if _, err := plan.KDEEntryPointActionPreview("launcher", "invalid", nil); err == nil {
		t.Fatalf("KDEEntryPointActionPreview accepted an invalid decision")
	}
	if _, err := plan.KDEEntryPointActionPreview("file-manager", "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEEntryPointActionPreview accepted a non-file URI")
	}

	preview, err := plan.KDEEntryPointActionPreview("settings", "approved", nil)
	if err != nil {
		t.Fatalf("KDEEntryPointActionPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE entrypoint action preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
