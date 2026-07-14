package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEEntryPointsPreviewCoversSevenEntryPointsWithoutActivating(t *testing.T) {
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

	preview, err := plan.KDEEntryPointsPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEEntryPointsPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_entrypoints.v1" ||
		preview.RequestType != "kde-entrypoints-preview" ||
		preview.SurfaceType != "kde-first-release-entrypoints" ||
		preview.Source != "execution-session-status-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetKDEEntryPointsPreview" {
		t.Fatalf("unexpected KDE entrypoints schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected KDE entrypoints identity: %#v", preview)
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
	if preview.EntryPointCount != 7 ||
		preview.VisibleEntryPointCount != 7 ||
		preview.PlannedEntryPointCount != 7 ||
		preview.ActiveEntryPointCount != 0 ||
		preview.PortalEntryPointCount != 1 ||
		preview.RuntimeGateEntryPointCount != 7 {
		t.Fatalf("unexpected entrypoint counts: %#v", preview)
	}
	if strings.Join(preview.EntryPointIDs, ",") != "launcher,task-manager,file-manager,system-tray,notifications,compatibility-center,settings" {
		t.Fatalf("unexpected entrypoint ids: %#v", preview.EntryPointIDs)
	}
	for _, entryPoint := range preview.EntryPoints {
		if !entryPoint.Visible ||
			!entryPoint.Planned ||
			entryPoint.Active ||
			!entryPoint.RequiresRuntimeGate ||
			!entryPoint.BlockedByRuntimeGate ||
			entryPoint.WritesHost ||
			entryPoint.StartsBackend ||
			entryPoint.BackendDetailsExposed {
			t.Fatalf("unexpected entrypoint safety flags for %s: %#v", entryPoint.ID, entryPoint)
		}
	}
	fileManager := findKDEEntryPoint(preview.EntryPoints, "file-manager")
	if fileManager.KDEComponent != "Dolphin" ||
		fileManager.RuntimeSource != "file-open-preview" ||
		!fileManager.RequiresPortal {
		t.Fatalf("unexpected file-manager entrypoint: %#v", fileManager)
	}
	taskManager := findKDEEntryPoint(preview.EntryPoints, "task-manager")
	if taskManager.KDEComponent != "Plasma task manager" ||
		taskManager.RuntimeSource != "execution-session-status-preview" ||
		taskManager.State != "planned" {
		t.Fatalf("unexpected task-manager entrypoint: %#v", taskManager)
	}
	settings := findKDEEntryPoint(preview.EntryPoints, "settings")
	if settings.RuntimeMethod != "GetCompatibilitySettings" ||
		settings.UserAction != "Open application settings" {
		t.Fatalf("unexpected settings entrypoint: %#v", settings)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.StableDesktopContract ||
		!preview.NormalApplicationSurface || !preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics || !preview.DesktopEntryLaunchVisible ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		!preview.EntryPointPlanCreated || preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE entrypoints safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "write KDE entrypoint files from preview") ||
		!containsString(preview.BlockedActions, "activate task manager entry from entrypoint preview") ||
		!containsString(preview.BlockedActions, "send notifications from entrypoint preview") ||
		!containsString(preview.BlockedActions, "start compatibility profile from entrypoint preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.KDEEntryPointsPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEEntryPointsPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.SessionStatus.SessionState != "review-declined" ||
		rejectedPreview.SessionStatus.BlockedGateCount != 2 ||
		rejectedPreview.LaunchAllowed ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected KDE entrypoints preview: %#v", rejectedPreview)
	}

	if _, err := plan.KDEEntryPointsPreview("invalid", nil); err == nil {
		t.Fatalf("KDEEntryPointsPreview accepted an invalid decision")
	}
	if _, err := plan.KDEEntryPointsPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEEntryPointsPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE entrypoints preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func findKDEEntryPoint(entryPoints []KDEEntryPointPreview, id string) KDEEntryPointPreview {
	for _, entryPoint := range entryPoints {
		if entryPoint.ID == id {
			return entryPoint
		}
	}
	return KDEEntryPointPreview{}
}
