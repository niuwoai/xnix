package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewPlanPresentsCompatibilityAppAsNormalDesktopApp(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc", ".XLS"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	if plan.ApplicationID != "org.example.ledger" {
		t.Fatalf("ApplicationID = %q", plan.ApplicationID)
	}
	if plan.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("DesktopFile = %q", plan.DesktopFile)
	}
	if got, want := plan.LaunchCommand, []string{"xnix-compat-launch", "--app", "org.example.ledger", "%U"}; !sameStrings(got, want) {
		t.Fatalf("LaunchCommand = %#v, want %#v", got, want)
	}
	if got, want := plan.MIMETypes, []string{"application/x-xnix-abc", "application/x-xnix-xls"}; !sameStrings(got, want) {
		t.Fatalf("MIMETypes = %#v, want %#v", got, want)
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		t.Fatalf("ValidateSafeForDesktop returned error: %v", err)
	}
}

func TestPlanJSONHidesBackendTerminology(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.notepad",
		Name:                "Example Notepad",
		Icon:                "accessories-text-editor",
		Mode:                "wine",
		SupportedExtensions: []string{".txt"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("plan JSON exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestRenderDesktopEntryUsesManagedRuntimeLauncher(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	entry, err := plan.RenderDesktopEntry()
	if err != nil {
		t.Fatalf("RenderDesktopEntry returned error: %v", err)
	}
	required := []string{
		"[Desktop Entry]\n",
		"Type=Application\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"Icon=office-chart-area\n",
		"StartupWMClass=xnix-org.example.ledger\n",
		"X-Xnix-ApplicationId=org.example.ledger\n",
		"X-Xnix-RuntimeOwned=true\n",
		"MimeType=application/x-xnix-abc;application/x-xnix-xls;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(entry, fragment) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(strings.ToLower(entry), forbidden) {
			t.Fatalf("desktop entry exposes forbidden term %q: %s", forbidden, entry)
		}
	}
}

func TestRenderDesktopEntryConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	entry, err := plan.RenderDesktopEntryWithOptions(DesktopEntryOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("RenderDesktopEntryWithOptions returned error: %v", err)
	}
	required := []string{
		"[Desktop Entry]\n",
		"Type=Application\n",
		"Name=Example Ledger\n",
		"Exec=xnix-compat-launch --app org.example.ledger %U\n",
		"X-Xnix-ApplicationId=org.example.ledger\n",
		"X-Xnix-RuntimeOwned=true\n",
		"MimeType=application/x-xnix-abc;application/x-xnix-xls;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(entry, fragment) {
			t.Fatalf("desktop entry missing %q in:\n%s", fragment, entry)
		}
	}
	if strings.Contains(entry, root) {
		t.Fatalf("desktop entry exposes activation root: %s", entry)
	}
	if _, err := plan.RenderDesktopEntryWithOptions(DesktopEntryOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("RenderDesktopEntryWithOptions accepted a missing activation receipt")
	}
}

func TestRenderMIMEAppsUsesGeneratedDesktopFile(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	mimeapps, err := plan.RenderMIMEApps()
	if err != nil {
		t.Fatalf("RenderMIMEApps returned error: %v", err)
	}
	required := []string{
		"[Default Applications]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop\n",
		"application/x-xnix-xls=xnix-org.example.ledger.desktop\n",
		"[Added Associations]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop;\n",
		"application/x-xnix-xls=xnix-org.example.ledger.desktop;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(mimeapps, fragment) {
			t.Fatalf("MIME apps preview missing %q in:\n%s", fragment, mimeapps)
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(strings.ToLower(mimeapps), forbidden) {
			t.Fatalf("MIME apps preview exposes forbidden term %q: %s", forbidden, mimeapps)
		}
	}
}

func TestRenderMIMEAppsConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	mimeapps, err := plan.RenderMIMEAppsWithOptions(MIMEAppsOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("RenderMIMEAppsWithOptions returned error: %v", err)
	}
	required := []string{
		"[Default Applications]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop\n",
		"application/x-xnix-xls=xnix-org.example.ledger.desktop\n",
		"[Added Associations]\n",
		"application/x-xnix-abc=xnix-org.example.ledger.desktop;\n",
		"application/x-xnix-xls=xnix-org.example.ledger.desktop;\n",
	}
	for _, fragment := range required {
		if !strings.Contains(mimeapps, fragment) {
			t.Fatalf("MIME apps preview missing %q in:\n%s", fragment, mimeapps)
		}
	}
	if strings.Contains(mimeapps, root) {
		t.Fatalf("MIME apps preview exposes activation root: %s", mimeapps)
	}
	if _, err := plan.RenderMIMEAppsWithOptions(MIMEAppsOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("RenderMIMEAppsWithOptions accepted a missing activation receipt")
	}
}

func TestWindowIdentityPreviewUsesNormalDesktopWindowIdentity(t *testing.T) {
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

	preview, err := plan.WindowIdentityPreview()
	if err != nil {
		t.Fatalf("WindowIdentityPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.window_identity.v1" {
		t.Fatalf("SchemaVersion = %q", preview.SchemaVersion)
	}
	if preview.Desktop != "KDE Plasma" || preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop identity: %#v", preview)
	}
	if preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("LauncherURL = %q", preview.LauncherURL)
	}
	if preview.TaskManager.GroupingKey != "org.example.ledger" || !preview.TaskManager.PinningAllowed ||
		!preview.TaskManager.RestoreAllowed || preview.TaskManager.SkipTaskbar || !preview.TaskManager.ShowInSwitcher {
		t.Fatalf("unexpected task manager hints: %#v", preview.TaskManager)
	}
	if preview.KWin.ScriptRole != "identity-and-layout" || preview.KWin.DesktopFile != "xnix-org.example.ledger.desktop" ||
		!preview.KWin.WindowManagerPolicyOnly || !preview.KWin.RuntimeOwnsBackendPolicy {
		t.Fatalf("unexpected KWin hints: %#v", preview.KWin)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("window identity preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestDesktopIconPreviewUsesStandardDesktopEntry(t *testing.T) {
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

	preview, err := plan.DesktopIconPreview()
	if err != nil {
		t.Fatalf("DesktopIconPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.desktop_icon.v1" ||
		preview.RequestType != "desktop-icon-preview" ||
		preview.PlanType != "desktop-icon-plan" ||
		preview.Source != "desktop-entry-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetDesktopIconPlan" {
		t.Fatalf("unexpected desktop icon schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.TargetDirectory != "xdg-desktop-dir" ||
		preview.Placement != "user-desktop" {
		t.Fatalf("unexpected desktop icon identity: %#v", preview)
	}
	if got, want := preview.LaunchCommand, []string{"xnix-compat-launch", "--app", "org.example.ledger", "%U"}; !sameStrings(got, want) {
		t.Fatalf("LaunchCommand = %#v, want %#v", got, want)
	}
	if !preview.StandardDesktopEntry || !preview.UserVisible || !preview.DesktopIconVisible ||
		preview.DesktopFileCopyEnabled || preview.DesktopFileWriteEnabled ||
		preview.IconPlacementPersisted || preview.LaunchEnabled ||
		preview.BackendLaunchEnabled || preview.HostRootModified ||
		preview.BackendDetailsExposed || preview.RawExecutableExposed ||
		!preview.RuntimeOwned || !preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner || !preview.OfficialDesktopOnly {
		t.Fatalf("unexpected desktop icon safety flags: %#v", preview)
	}
	if !sameStrings(preview.BlockedActions, []string{
		"copy desktop entry into user desktop from preview state",
		"persist desktop icon placement from preview state",
		"launch application from desktop icon preview",
		"expose raw backend command in desktop icon",
		"mutate host root during desktop icon preview",
	}) {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop icon preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestDesktopIconPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")
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

	preview, err := plan.DesktopIconPreviewWithOptions(DesktopIconOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("DesktopIconPreviewWithOptions returned error: %v", err)
	}
	if preview.Source != "desktop-entry-preview+desktop-activation-receipt" ||
		!preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("desktop icon preview did not consume activation receipt: %#v", preview)
	}
	if preview.DesktopFileCopyEnabled ||
		preview.DesktopFileWriteEnabled ||
		preview.IconPlacementPersisted ||
		preview.LaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.RawExecutableExposed {
		t.Fatalf("receipt-backed desktop icon preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("desktop icon preview exposes activation root: %s", string(encoded))
	}
	if _, err := plan.DesktopIconPreviewWithOptions(DesktopIconOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("DesktopIconPreviewWithOptions accepted a missing activation receipt")
	}
}

func TestTrayStatusPreviewKeepsLiveBridgeGated(t *testing.T) {
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

	preview, err := plan.TrayStatusPreview()
	if err != nil {
		t.Fatalf("TrayStatusPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.tray_status.v1" || preview.StatusType != "tray-status-preview" {
		t.Fatalf("unexpected tray schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" || preview.ApplicationID != "org.example.ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" || preview.Icon != "office-chart-area" {
		t.Fatalf("unexpected tray identity: %#v", preview)
	}
	if preview.RuntimeActivity.RegisteredApplicationCount != 1 || preview.RuntimeActivity.ActiveApplicationCount != 0 ||
		preview.RuntimeActivity.AttentionRequiredCount != 0 {
		t.Fatalf("unexpected tray runtime activity: %#v", preview.RuntimeActivity)
	}
	if preview.CompatibilityStatus.State != "ready" || preview.TrayBridge.State != "planned" ||
		preview.TrayBridge.BridgedTrayApplicationCount != 0 {
		t.Fatalf("unexpected tray status: %#v %#v", preview.CompatibilityStatus, preview.TrayBridge)
	}
	if !sameStrings(preview.Actions, []string{"open-compatibility-center", "open-settings"}) {
		t.Fatalf("Actions = %#v", preview.Actions)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.LiveBackendBridgeEnabled || preview.BridgeConfigurationPersisted ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected tray safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("tray status preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestTrayStatusPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

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

	preview, err := plan.TrayStatusPreviewWithOptions(TrayStatusOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("TrayStatusPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("unexpected receipt-backed tray fields: %#v", preview)
	}
	if preview.LiveBackendBridgeEnabled ||
		preview.BridgeConfigurationPersisted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed tray preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("tray preview exposed activation root: %s", encoded)
	}

	if _, err := plan.TrayStatusPreviewWithOptions(TrayStatusOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("tray preview accepted a missing activation receipt")
	}
}

func TestTrayStatusPreviewConsumesExecutionSessionRecord(t *testing.T) {
	root := t.TempDir()
	writeExecutionSessionRecord(t, root, "xnix-exec-org-example-ledger-1", "org.example.ledger")

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

	preview, err := plan.TrayStatusPreviewWithOptions(TrayStatusOptions{
		ExecutionSessionRoot:      root,
		ExecutionSessionRequestID: "xnix-exec-org-example-ledger-1",
	})
	if err != nil {
		t.Fatalf("TrayStatusPreviewWithOptions returned error: %v", err)
	}
	if !preview.ExecutionSessionRoot ||
		!preview.ExecutionSessionBacked ||
		preview.ExecutionSessionPath != "execution-ledger/sessions/xnix-exec-org-example-ledger-1.json" ||
		preview.ExecutionSessionState != "blocked" ||
		preview.CompatibilityStatus.State != "waiting-for-runtime-gates" ||
		preview.CompatibilityStatus.Label != "Runtime gates required" {
		t.Fatalf("tray did not consume execution session record: %#v", preview)
	}
	if preview.LiveBackendBridgeEnabled ||
		preview.BridgeConfigurationPersisted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("session-backed tray preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("tray preview exposed session root: %s", encoded)
	}
	if _, err := plan.TrayStatusPreviewWithOptions(TrayStatusOptions{ExecutionSessionRoot: root, ExecutionSessionRequestID: "missing"}); err == nil {
		t.Fatalf("tray preview accepted a missing execution session record")
	}
}

func TestNotificationPreviewKeepsExecutionGatesClosed(t *testing.T) {
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

	preview, err := plan.NotificationPreview("install-failed")
	if err != nil {
		t.Fatalf("NotificationPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.notification.v1" || preview.RequestType != "desktop-notification-preview" {
		t.Fatalf("unexpected notification schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" || preview.Source != "runtime-event" ||
		preview.ApplicationID != "org.example.ledger" || preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected notification identity: %#v", preview)
	}
	if preview.EventType != "install-failed" || preview.Urgency != "critical" ||
		preview.Category != "compatibility.install" || !preview.RequiresUserReview {
		t.Fatalf("unexpected install-failed event metadata: %#v", preview)
	}
	if !sameStrings(preview.Actions, []string{"open-compatibility-center", "show-diagnostics"}) {
		t.Fatalf("Actions = %#v", preview.Actions)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.ActionExecutionEnabled || preview.RepairExecutionEnabled ||
		preview.SettingsPersistenceEnabled || preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected notification safety flags: %#v", preview)
	}

	approval, err := plan.NotificationPreview("approval-required")
	if err != nil {
		t.Fatalf("NotificationPreview approval returned error: %v", err)
	}
	if approval.Urgency != "critical" || approval.Category != "compatibility.approval" || !approval.RequiresUserReview {
		t.Fatalf("unexpected approval metadata: %#v", approval)
	}
	if !sameStrings(approval.Actions, []string{"open-compatibility-center", "review-request"}) {
		t.Fatalf("approval actions = %#v", approval.Actions)
	}
	if _, err := plan.NotificationPreview("unknown-event"); err == nil {
		t.Fatalf("NotificationPreview accepted an unknown event")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("notification preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestNotificationPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

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

	preview, err := plan.NotificationPreviewWithOptions("approval-required", NotificationOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("NotificationPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("unexpected receipt-backed notification fields: %#v", preview)
	}
	if preview.ActionExecutionEnabled ||
		preview.RepairExecutionEnabled ||
		preview.SettingsPersistenceEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed notification preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("notification preview exposed activation root: %s", encoded)
	}

	if _, err := plan.NotificationPreviewWithOptions("approval-required", NotificationOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("notification preview accepted a missing activation receipt")
	}
}

func TestSettingsPreviewExposesUserFacingControls(t *testing.T) {
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

	preview, err := plan.SettingsPreview()
	if err != nil {
		t.Fatalf("SettingsPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.settings.v1" || preview.RequestType != "settings-preview" {
		t.Fatalf("unexpected settings schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" || preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" || preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected settings identity: %#v", preview)
	}
	if preview.SectionCount != 5 || len(preview.Sections) != 5 {
		t.Fatalf("unexpected section count: %#v", preview.Sections)
	}
	if got, want := settingSectionIDs(preview.Sections), []string{"run-mode", "resource-access", "devices", "network", "snapshots"}; !sameStrings(got, want) {
		t.Fatalf("section ids = %#v, want %#v", got, want)
	}
	if settingFieldValue(preview.Sections, "run-mode", "mode") != "automatic" {
		t.Fatalf("run mode did not default to automatic: %#v", preview.Sections)
	}
	if settingFieldValue(preview.Sections, "resource-access", "documents") != "ask" ||
		settingFieldValue(preview.Sections, "resource-access", "downloads") != "ask" {
		t.Fatalf("file access did not default to review: %#v", preview.Sections)
	}
	if settingFieldValue(preview.Sections, "devices", "camera") != "deny" ||
		settingFieldValue(preview.Sections, "network", "network") != "allow" ||
		settingFieldValue(preview.Sections, "snapshots", "snapshots") != "enabled" {
		t.Fatalf("unexpected device, network, or snapshot defaults: %#v", preview.Sections)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.SettingsPersisted || preview.SettingsPersistenceEnabled ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected settings safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("settings preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestSettingsPreviewConsumesActivationReceipt(t *testing.T) {
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

	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := plan.SettingsPreviewWithOptions(SettingsOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("SettingsPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("settings preview did not consume activation receipt: %#v", preview)
	}
	if preview.SettingsPersisted ||
		preview.SettingsPersistenceEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed settings preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("settings preview exposed activation root: %s", encoded)
	}

	if _, err := plan.SettingsPreviewWithOptions(SettingsOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("settings preview accepted a missing activation receipt")
	}
}

func TestModeSwitchPreviewKeepsRuntimeStateGated(t *testing.T) {
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

	preview, err := plan.ModeSwitchPreview("prefer-compatibility")
	if err != nil {
		t.Fatalf("ModeSwitchPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.mode_switch.v1" ||
		preview.RequestType != "mode-switch-preview" ||
		preview.PlanType != "compatibility-mode-switch-plan" ||
		preview.Source != "unified-settings" ||
		preview.RuntimeMethod != "GetCompatibilityModeSwitchPlan" {
		t.Fatalf("unexpected mode switch schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected mode switch identity: %#v", preview)
	}
	if preview.CurrentMode != "automatic" ||
		preview.RequestedMode != "prefer-compatibility" ||
		preview.ModeState != "planned" ||
		preview.ModeCount != 4 {
		t.Fatalf("unexpected mode switch request: %#v", preview)
	}
	if got, want := modeSwitchOptionIDs(preview.Modes), []string{"automatic", "prefer-performance", "prefer-compatibility", "isolated-execution"}; !sameStrings(got, want) {
		t.Fatalf("mode ids = %#v, want %#v", got, want)
	}
	if countSelectedModes(preview.Modes) != 1 || countRequestedModes(preview.Modes) != 1 {
		t.Fatalf("mode switch preview must mark one selected and one requested mode: %#v", preview.Modes)
	}
	if !sameStrings(preview.RequiredRuntimeGates, []string{"settings-review", "portal-policy-review", "snapshot-baseline", "backend-environment-plan", "runtime-write-gate"}) {
		t.Fatalf("unexpected required gates: %#v", preview.RequiredRuntimeGates)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || !preview.UserVisible ||
		!preview.ValidMode || !preview.RequiresUserConfirmation || !preview.PortalReviewRequired ||
		!preview.SnapshotRequired || preview.SettingsPersistenceEnabled ||
		preview.BackendReconfigurationEnabled || preview.BackendProcessStarted ||
		preview.LaunchEnabled || preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected mode switch safety flags: %#v", preview)
	}

	if _, err := plan.ModeSwitchPreview("unsupported-mode"); err == nil {
		t.Fatalf("ModeSwitchPreview accepted an unsupported mode")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("mode switch preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestSettingsChangePreviewPlansReviewBeforePersistence(t *testing.T) {
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

	preview, err := plan.SettingsChangePreview("resource-access", "documents", "allow")
	if err != nil {
		t.Fatalf("SettingsChangePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.settings_change.v1" ||
		preview.RequestType != "settings-change-preview" ||
		preview.PlanType != "settings-change-plan" ||
		preview.Source != "unified-settings" ||
		preview.RuntimeMethod != "GetCompatibilitySettingsChangePlan" {
		t.Fatalf("unexpected settings change schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected settings change identity: %#v", preview)
	}
	if preview.SectionID != "resource-access" ||
		preview.FieldID != "documents" ||
		preview.RequestedValue != "allow" ||
		preview.ChangeState != "planned" {
		t.Fatalf("unexpected settings change request: %#v", preview)
	}
	if !preview.UserConfirmationRequired || !preview.PortalPolicyReviewRequired ||
		preview.SnapshotRecommended || preview.RuntimeRestartRequired {
		t.Fatalf("unexpected review requirements: %#v", preview)
	}
	if preview.AffectedPolicy.Section != "resource-access" ||
		preview.AffectedPolicy.Field != "documents" ||
		preview.AffectedPolicy.Value != "allow" ||
		!sameStrings(preview.AffectedPolicy.Options, []string{"allow", "ask", "deny"}) {
		t.Fatalf("unexpected affected policy: %#v", preview.AffectedPolicy)
	}
	if len(preview.Steps) != 5 ||
		preview.Steps[0].Status != "pass" ||
		preview.Steps[1].Status != "required" ||
		preview.Steps[2].Status != "required" ||
		preview.Steps[3].Status != "pass" ||
		preview.Steps[4].Status != "pending" {
		t.Fatalf("unexpected settings change steps: %#v", preview.Steps)
	}
	if !sameStrings(preview.BlockedActions, []string{
		"persist compatibility settings before Runtime confirmation",
		"grant desktop resources without Portal policy review",
		"modify host root while planning settings changes",
		"expose backend implementation settings to KDE",
	}) {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.ApplyEnabled || preview.SettingsPersisted || preview.SettingsPersistenceEnabled ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected settings change safety flags: %#v", preview)
	}

	mode, err := plan.SettingsChangePreview("run-mode", "mode", "performance")
	if err != nil {
		t.Fatalf("run-mode SettingsChangePreview returned error: %v", err)
	}
	if !mode.SnapshotRecommended || mode.PortalPolicyReviewRequired ||
		mode.Steps[3].Status != "recommended" {
		t.Fatalf("unexpected run-mode review requirements: %#v", mode)
	}

	if _, err := plan.SettingsChangePreview("resource-access", "documents", "always"); err == nil {
		t.Fatalf("SettingsChangePreview accepted an unsupported value")
	}
	if _, err := plan.SettingsChangePreview("unknown", "mode", "automatic"); err == nil {
		t.Fatalf("SettingsChangePreview accepted an unknown field")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("settings change preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestReviewFlowPreviewConnectsSettingsPermissionsAndPortalReviews(t *testing.T) {
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

	preview, err := plan.ReviewFlowPreview("resource-access", "documents", "ask", "file-open")
	if err != nil {
		t.Fatalf("ReviewFlowPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.review_flow.v1" ||
		preview.RequestType != "review-flow-preview" ||
		preview.PlanType != "compatibility-review-flow-plan" ||
		preview.Source != "compatibility-center-review" ||
		preview.RuntimeMethod != "GetCompatibilityReviewFlowPlan" {
		t.Fatalf("unexpected review flow schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected review flow identity: %#v", preview)
	}
	if preview.SectionID != "resource-access" ||
		preview.FieldID != "documents" ||
		preview.RequestedValue != "ask" ||
		preview.Operation != "file-open" ||
		preview.ReviewState != "planned" {
		t.Fatalf("unexpected review flow request: %#v", preview)
	}
	if preview.StepCount != 5 ||
		preview.RequiredReviewCount != 3 ||
		preview.BlockedStepCount != 1 ||
		preview.PendingStepCount != 1 {
		t.Fatalf("unexpected review flow counts: %#v", preview)
	}
	if got, want := reviewFlowStepIDs(preview.Steps), []string{"settings-change-review", "permission-review", "portal-request-review", "runtime-write-gate", "review-receipt"}; !sameStrings(got, want) {
		t.Fatalf("review flow step ids = %#v, want %#v", got, want)
	}
	if preview.SettingsChangePlan.PlanType != "settings-change-plan" ||
		preview.SettingsChangePlan.ChangeState != "planned" ||
		preview.SettingsChangePlan.ApplyEnabled ||
		preview.SettingsChangePlan.SettingsPersisted {
		t.Fatalf("unexpected settings change summary: %#v", preview.SettingsChangePlan)
	}
	if preview.PermissionReviewPlan.PlanType != "compatibility-permission-review-plan" ||
		preview.PermissionReviewPlan.PermissionCount != 7 ||
		!preview.PermissionReviewPlan.UserReviewRequired ||
		!preview.PermissionReviewPlan.PortalReviewRequired ||
		preview.PermissionReviewPlan.PermissionChangesApplied ||
		preview.PermissionReviewPlan.PermissionsGranted {
		t.Fatalf("unexpected permission review summary: %#v", preview.PermissionReviewPlan)
	}
	if preview.PortalRequestPlan.RequestType != "portal-request-preview" ||
		preview.PortalRequestPlan.Operation != "file-open" ||
		preview.PortalRequestPlan.Decision != "ask" ||
		!preview.PortalRequestPlan.RequestAllowed ||
		!preview.PortalRequestPlan.PortalRequired ||
		preview.PortalRequestPlan.RequestObjectCreated ||
		preview.PortalRequestPlan.PermissionGranted ||
		preview.PortalRequestPlan.HostPermissionChanged ||
		preview.PortalRequestPlan.BackendDetailsExposed {
		t.Fatalf("unexpected Portal request summary: %#v", preview.PortalRequestPlan)
	}
	if preview.RuntimeWriteGate.GateType != "runtime-write-gate" ||
		preview.RuntimeWriteGate.MethodName != "Launch" ||
		preview.RuntimeWriteGate.WriteMethodEnabled ||
		preview.RuntimeWriteGate.DispatchEnabled ||
		preview.RuntimeWriteGate.RequestObjectCreated ||
		!sameStrings(preview.RuntimeWriteGate.RequiredGates, []string{
			"production-runtime-owner",
			"backend-binding-ready",
			"recipe-trust-production",
			"user-action-review",
			"portal-approval-if-sensitive",
			"snapshot-preflight-for-risky-change",
		}) {
		t.Fatalf("unexpected write gate summary: %#v", preview.RuntimeWriteGate)
	}
	if preview.ReviewReceipt.ReceiptType != "compatibility-center-action-review-receipt" ||
		preview.ReviewReceipt.DecisionRecorded ||
		preview.ReviewReceipt.ExecutionEnabled ||
		preview.ReviewReceipt.SettingsPersistenceEnabled ||
		preview.ReviewReceipt.ResourceGrantCreated {
		t.Fatalf("unexpected review receipt summary: %#v", preview.ReviewReceipt)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || !preview.UserVisible ||
		!preview.UserConfirmationRequired || !preview.PortalPolicyReviewRequired ||
		!preview.SettingsChangePlanned || !preview.PermissionReviewPlanned ||
		!preview.PortalRequestPlanned || !preview.ReviewReceiptRequired ||
		preview.ApplyEnabled || preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.SettingsPersisted || preview.ExecutionStarted || preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected review flow safety flags: %#v", preview)
	}

	if _, err := plan.ReviewFlowPreview("resource-access", "documents", "always", "file-open"); err == nil {
		t.Fatalf("ReviewFlowPreview accepted an unsupported settings value")
	}
	if _, err := plan.ReviewFlowPreview("resource-access", "documents", "ask", "unknown"); err == nil {
		t.Fatalf("ReviewFlowPreview accepted an unsupported Portal operation")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("review flow preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPermissionReviewPreviewKeepsPermissionGatesClosed(t *testing.T) {
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

	preview, err := plan.PermissionReviewPreview()
	if err != nil {
		t.Fatalf("PermissionReviewPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.permission_review.v1" ||
		preview.RequestType != "permission-review-preview" ||
		preview.PlanType != "compatibility-permission-review-plan" {
		t.Fatalf("unexpected permission review schema: %#v", preview)
	}
	if preview.Source != "unified-settings" || preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetCompatibilityPermissionReviewPlan" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected permission review identity: %#v", preview)
	}
	if preview.ReviewState != "planned" || preview.PermissionCount != 7 ||
		preview.AllowCount != 1 || preview.AskCount != 5 || preview.DenyCount != 1 {
		t.Fatalf("unexpected permission counts: %#v", preview)
	}
	if got, want := permissionReviewIDs(preview.Permissions), []string{"documents", "downloads", "camera", "network", "clipboard", "print", "screenshot"}; !sameStrings(got, want) {
		t.Fatalf("permission ids = %#v, want %#v", got, want)
	}
	if permissionDecision(preview.Permissions, "network") != "allow" ||
		permissionDecision(preview.Permissions, "camera") != "deny" ||
		permissionDecision(preview.Permissions, "documents") != "ask" {
		t.Fatalf("unexpected permission decisions: %#v", preview.Permissions)
	}
	for _, permission := range preview.Permissions {
		if permission.ChangePending || permission.RequestObjectCreated ||
			permission.PermissionGranted || permission.DirectAccessAllowed ||
			permission.BackendDetailsExposed {
			t.Fatalf("permission gate unexpectedly open: %#v", permission)
		}
	}
	if !sameStrings(preview.RequiredRuntimeGates, []string{"user-review", "portal-policy-review", "runtime-write-gate", "settings-persistence", "audit-log"}) {
		t.Fatalf("required gates = %#v", preview.RequiredRuntimeGates)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		!preview.UserReviewRequired || !preview.PortalReviewRequired ||
		preview.PermissionChangesApplied || preview.RequestObjectsCreated ||
		preview.PermissionsGranted || preview.SettingsPersisted ||
		preview.SettingsPersistenceEnabled || preview.HostPermissionChanged ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected permission safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("permission review preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestDesktopResourceBridgePreviewKeepsBridgesDisabled(t *testing.T) {
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

	preview, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		t.Fatalf("DesktopResourceBridgePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.desktop_resource_bridge.v1" ||
		preview.RequestType != "desktop-resource-bridge-preview" ||
		preview.PlanType != "desktop-resource-bridge-plan" {
		t.Fatalf("unexpected desktop resource bridge schema: %#v", preview)
	}
	if preview.Source != "runtime-resource-boundary" || preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetDesktopResourceBridgePlan" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop resource bridge identity: %#v", preview)
	}
	if preview.BridgeState != "planned" || preview.ResourceCount != 5 {
		t.Fatalf("unexpected bridge state or count: %#v", preview)
	}
	if got, want := desktopResourceBridgeIDs(preview.Resources), []string{"file-open", "uri-open", "print", "clipboard", "screenshot"}; !sameStrings(got, want) {
		t.Fatalf("resource ids = %#v, want %#v", got, want)
	}
	for _, resource := range preview.Resources {
		if resource.RuntimeMethod != "GetPortalRequestPlan" ||
			resource.State != "planned" ||
			!resource.PortalRequired ||
			!resource.UserApprovalRequired ||
			resource.BridgeEnabled ||
			resource.RequestCreated ||
			resource.DirectBackendAccessAllowed ||
			resource.BackendDetailsExposed {
			t.Fatalf("unexpected resource bridge gate: %#v", resource)
		}
	}
	if !sameStrings(preview.RequiredRuntimeGates, []string{"portal-policy-review", "portal-request-plan", "snapshot-baseline", "backend-environment-plan", "runtime-write-gate"}) {
		t.Fatalf("required gates = %#v", preview.RequiredRuntimeGates)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		!preview.PortalMediated || !preview.FileBridgePlanned ||
		!preview.URIBridgePlanned || !preview.PrintBridgePlanned ||
		!preview.ClipboardBridgePlanned || !preview.ScreenshotBridgePlanned ||
		preview.BridgesEnabled || preview.RequestsCreated ||
		preview.BackendProcessStarted || preview.DirectHostFileAccess ||
		preview.DirectClipboardAccess || preview.DirectPrintAccess ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected desktop resource bridge safety flags: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop resource bridge preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
	for _, hostPath := range []string{"/users", "/home", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, hostPath) {
			t.Fatalf("desktop resource bridge preview exposes host path %q: %s", hostPath, text)
		}
	}
}

func TestPortalRequestPreviewPlansPortalFlowWithoutGrantingAccess(t *testing.T) {
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

	preview, err := plan.PortalRequestPreview("file-open", "Open a selected document.")
	if err != nil {
		t.Fatalf("PortalRequestPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.portal_request.v1" ||
		preview.RequestType != "portal-request-preview" ||
		preview.Source != "runtime-portal-request-plan" {
		t.Fatalf("unexpected Portal request schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetPortalRequestPlan" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected Portal request identity: %#v", preview)
	}
	if preview.Operation != "file-open" ||
		preview.Reason != "Open a selected document." ||
		preview.Decision != "ask" ||
		!preview.RequestAllowed {
		t.Fatalf("unexpected file-open decision: %#v", preview)
	}
	if preview.Portal.Destination != "org.freedesktop.portal.Desktop" ||
		preview.Portal.Interface != "org.freedesktop.portal.FileChooser" ||
		preview.Portal.Method != "OpenFile" ||
		preview.Portal.ObjectPath != "/org/freedesktop/portal/desktop" ||
		preview.Portal.DBusAPI != "XDG Desktop Portal" {
		t.Fatalf("unexpected Portal endpoint: %#v", preview.Portal)
	}
	if !preview.Request.ObjectPathRequired ||
		preview.Request.RequestObjectCreated ||
		preview.Request.HandleToken != "xnix_org_example_ledger_file_open" ||
		!preview.Request.UserMediationRequired ||
		!sameStrings(preview.Request.Resources, []string{"documents", "downloads", "selected-files"}) ||
		!preview.Request.RuntimePolicyOwner ||
		preview.Request.DesktopShellPolicyOwner {
		t.Fatalf("unexpected Portal request object: %#v", preview.Request)
	}
	if preview.Completion.Signal != "Response" ||
		preview.Completion.ResponseField != "response" ||
		preview.Completion.SuccessCode != 0 ||
		preview.Completion.CancelledCode != 1 ||
		preview.Completion.DeniedCode != 2 ||
		preview.Completion.ResultOwner != "Runtime" {
		t.Fatalf("unexpected Portal completion: %#v", preview.Completion)
	}
	if preview.Denied != nil {
		t.Fatalf("allowed Portal request included denial guidance: %#v", preview.Denied)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.Safety.DirectAccessAllowed || !preview.Safety.PortalRequired ||
		preview.Safety.PermissionGranted || preview.Safety.HostPermissionChanged ||
		preview.Safety.HostRootModified || preview.Safety.BackendDetailsExposed {
		t.Fatalf("unexpected Portal safety flags: %#v", preview)
	}

	camera, err := plan.PortalRequestPreview("camera", "")
	if err != nil {
		t.Fatalf("camera PortalRequestPreview returned error: %v", err)
	}
	if camera.Decision != "deny" || camera.RequestAllowed ||
		camera.Portal.Interface != "org.freedesktop.portal.Camera" ||
		camera.Portal.Method != "AccessCamera" ||
		camera.Denied == nil ||
		camera.Denied.NextAction != "open-compatibility-settings" ||
		camera.Denied.NotificationEvent != "approval-required" ||
		camera.Safety.PermissionGranted {
		t.Fatalf("unexpected camera denial: %#v", camera)
	}
	if _, err := plan.PortalRequestPreview("unknown", ""); err == nil {
		t.Fatalf("PortalRequestPreview accepted an unknown operation")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Portal request preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
	for _, hostPath := range []string{"/users", "/home", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, hostPath) {
			t.Fatalf("Portal request preview exposes host path %q: %s", hostPath, text)
		}
	}
}

func TestKRunnerQueryPreviewReturnsSafeLauncherMatches(t *testing.T) {
	preview, err := NewKRunnerQueryPreview([]Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "wine",
			SupportedExtensions: []string{".xls", ".abc"},
		},
		{
			ID:                  "org.example.notes",
			Name:                "Example Notes",
			Icon:                "accessories-text-editor",
			Mode:                "automatic",
			SupportedExtensions: []string{".txt"},
		},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "open xls")
	if err != nil {
		t.Fatalf("NewKRunnerQueryPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.krunner_query.v1" || preview.QueryType != "krunner-query-plan" {
		t.Fatalf("unexpected KRunner schema: %#v", preview)
	}
	if preview.EntryPoint != "krunner" || preview.Desktop != "KDE Plasma" || preview.Query != "open xls" {
		t.Fatalf("unexpected KRunner identity: %#v", preview)
	}
	if preview.Source.Kind != "runtime-go-registry" || preview.Source.RegistryName != "test-registry" ||
		!preview.Source.RecipeDigestVerified || preview.Source.RecipeSignatureStatus != "development-only" {
		t.Fatalf("unexpected KRunner source: %#v", preview.Source)
	}
	if len(preview.Matches) != 1 {
		t.Fatalf("match count = %d, want 1: %#v", len(preview.Matches), preview.Matches)
	}
	match := preview.Matches[0]
	if match.ApplicationID != "org.example.ledger" || match.Name != "Example Ledger" ||
		match.Action.DesktopEntryID != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected KRunner match identity: %#v", match)
	}
	if match.ModeLabel != "Managed compatibility" || match.RelevancePercent != 55 ||
		match.Action.Type != "runtime-launch" || !sameStrings(match.Action.Argv, []string{"xnix-compat-launch", "--app", "org.example.ledger"}) {
		t.Fatalf("unexpected KRunner match details: %#v", match)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || preview.HostRootModified ||
		preview.BackendDetailsExposed || preview.Summary.QueryExecutionEnabled ||
		preview.Summary.BackendLaunchEnabled || preview.Summary.BackendDetailsExposed ||
		!preview.Summary.RuntimeOwnedLaunch {
		t.Fatalf("unexpected KRunner safety flags: %#v", preview)
	}

	blank, err := NewKRunnerQueryPreview([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}}, Provenance{Source: "registry"}, " ")
	if err != nil {
		t.Fatalf("blank NewKRunnerQueryPreview returned error: %v", err)
	}
	if len(blank.Matches) != 0 || blank.Summary.MatchCount != 0 {
		t.Fatalf("blank query returned matches: %#v", blank)
	}
	if _, err := NewKRunnerQueryPreview(nil, Provenance{Source: "registry"}, "line\nbreak"); err == nil {
		t.Fatalf("KRunner query preview accepted a multiline query")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KRunner query preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKRunnerQueryPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := NewKRunnerQueryPreviewWithOptions([]Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "automatic",
			SupportedExtensions: []string{".xls", ".abc"},
		},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "ledger", KRunnerQueryOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("NewKRunnerQueryPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		preview.Summary.ReceiptBackedMatchCount != 1 ||
		preview.Summary.QueryExecutionEnabled ||
		preview.Summary.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected receipt-backed KRunner summary: %#v", preview)
	}
	if len(preview.Matches) != 1 {
		t.Fatalf("match count = %d, want 1", len(preview.Matches))
	}
	match := preview.Matches[0]
	if !match.ActivationReceiptBacked ||
		match.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		!match.RuntimeOwnedLaunch ||
		match.BackendDetailsExposed ||
		match.Action.Type != "runtime-launch" {
		t.Fatalf("unexpected receipt-backed KRunner match: %#v", match)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(root)) {
		t.Fatalf("KRunner receipt-backed preview exposed activation root: %s", text)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KRunner receipt-backed preview exposes forbidden term %q: %s", forbidden, text)
		}
	}

	if _, err := NewKRunnerQueryPreviewWithOptions([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}}, Provenance{Source: "registry"}, "ledger", KRunnerQueryOptions{ActivationRoot: t.TempDir()}); err == nil ||
		!strings.Contains(err.Error(), "read desktop activation receipt") {
		t.Fatalf("missing activation receipt must fail closed, got %v", err)
	}
}

func TestCompatibilityCenterPreviewSummarizesApplicationsSafely(t *testing.T) {
	preview, err := NewCompatibilityCenterPreview([]Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "wine",
			SupportedExtensions: []string{".xls", ".abc"},
		},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewCompatibilityCenterPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.compatibility_center.v1" ||
		preview.SummaryType != "compatibility-center-preview" ||
		preview.Desktop != "KDE Plasma" || preview.Title != "Xnix Compatibility Center" {
		t.Fatalf("unexpected center identity: %#v", preview)
	}
	if preview.Source.Kind != "runtime-go-registry" || preview.Source.RegistryName != "test-registry" ||
		!preview.Source.RecipeDigestVerified || preview.Source.RecipeSignatureStatus != "development-only" {
		t.Fatalf("unexpected center source: %#v", preview.Source)
	}
	if preview.ApplicationCount != 1 || len(preview.Applications) != 1 ||
		preview.KnownIssueCount != 0 || preview.RepairRecordCount != 0 || preview.PendingReviewCount != 0 ||
		preview.KnownAppSmokeEvidenceCount != 0 || preview.KnownAppSmokePassedCount != 0 ||
		len(preview.KnownAppSmokeEvidence) != 0 {
		t.Fatalf("unexpected center counts: %#v", preview)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner ||
		preview.ActionExecutionEnabled || preview.RepairExecutionEnabled ||
		preview.BackendLaunchEnabled || preview.SettingsPersistenceEnabled ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected center safety flags: %#v", preview)
	}

	application := preview.Applications[0]
	if application.ApplicationID != "org.example.ledger" || application.DisplayName != "Example Ledger" ||
		application.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected center application identity: %#v", application)
	}
	if application.CompatibilityState != "registered" || application.DiagnosticsState != "not-run" ||
		application.RuntimeMode != "Managed compatibility" || application.KnownIssueCount != 0 ||
		application.RepairRecordState != "none" || application.RepairRecordCount != 0 ||
		application.LastRepairEvent != "none" {
		t.Fatalf("unexpected center application status: %#v", application)
	}
	if !sameStrings(application.Actions, []string{"open-settings", "show-diagnostics", "review-application"}) {
		t.Fatalf("Actions = %#v", application.Actions)
	}
	if !application.RuntimeOwned || application.KDEPolicyOwner || !application.UserVisible ||
		application.ActionExecutionEnabled || application.RepairExecutionEnabled ||
		application.BackendLaunchEnabled || application.SettingsPersistenceEnabled ||
		application.HostRootModified || application.BackendDetailsExposed {
		t.Fatalf("unexpected center application safety flags: %#v", application)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Compatibility Center preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestCompatibilityCenterPreviewSummarizesKnownAppSmokeEvidenceSafely(t *testing.T) {
	preview, err := NewCompatibilityCenterPreviewWithOptions([]Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "automatic",
			SupportedExtensions: []string{".abc"},
		},
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true}, CompatibilityCenterOptions{
		KnownAppSmokeEvidence: []KnownAppSmokeEvidenceSummary{{
			AppID:            "7zr",
			DisplayName:      "7-Zip Console",
			AppVersion:       "26.02",
			EvidenceSource:   "staged-launcher-dispatch-smoke",
			SmokeStatus:      "passed",
			MarkerObserved:   true,
			ChecksumVerified: true,
		}},
	})
	if err != nil {
		t.Fatalf("NewCompatibilityCenterPreviewWithOptions returned error: %v", err)
	}
	if preview.KnownAppSmokeEvidenceCount != 1 ||
		preview.KnownAppSmokePassedCount != 1 ||
		preview.KnownAppStagedLauncherPassedCount != 1 ||
		preview.KnownAppLaunchAuthorizationRequiredCount != 1 ||
		len(preview.KnownAppSmokeEvidence) != 1 ||
		preview.Summary.Headline != "A known Windows application has passed the staged launcher Runtime dispatch smoke." {
		t.Fatalf("unexpected known app smoke counts: %#v", preview)
	}
	evidence := preview.KnownAppSmokeEvidence[0]
	if evidence.AppID != "7zr" ||
		evidence.DisplayName != "7-Zip Console" ||
		evidence.AppVersion != "26.02" ||
		evidence.EvidenceKind != "known-application-managed-smoke" ||
		evidence.EvidenceSource != "staged-launcher-dispatch-smoke" ||
		evidence.SmokeStatus != "passed" ||
		evidence.CompatibilityState != "validated" ||
		evidence.CenterCardState != "validated-launch-authorization-required" ||
		evidence.LaunchAuthorizationState != "review-required" ||
		evidence.PrimaryActionID != "review-launch-authorization" ||
		evidence.PrimaryActionLabel != "Review launch authorization" ||
		evidence.PrimaryActionKind != "authorization-review" ||
		!evidence.PrimaryActionEnabled ||
		evidence.DirectLaunchEnabled ||
		!evidence.MarkerObserved ||
		!evidence.ChecksumVerified ||
		!evidence.ExecutionEvidenceRecorded ||
		!evidence.StagedLauncherVerified ||
		!evidence.RuntimeDispatchVerified ||
		!evidence.LaunchAuthorizationRequired ||
		evidence.DesktopLaunchEnabled ||
		!evidence.RuntimeOwned ||
		evidence.KDEPolicyOwner ||
		evidence.ActionExecutionEnabled ||
		evidence.BackendLaunchEnabled ||
		evidence.HostRootModified ||
		evidence.BackendDetailsExposed ||
		evidence.RawArtifactPathExposed {
		t.Fatalf("unexpected known app smoke evidence: %#v", evidence)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Compatibility Center known app evidence exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestFileOpenPreviewRequiresPortalAndSelectsByExtension(t *testing.T) {
	preview, err := NewFileOpenPreview([]Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "automatic",
			SupportedExtensions: []string{".xls", ".abc"},
		},
	}, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true}, []string{"file:///home/test/Documents/book.xls"}, "")
	if err != nil {
		t.Fatalf("NewFileOpenPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.file_open.v1" || preview.RequestType != "file-open-preview" {
		t.Fatalf("unexpected file-open schema: %#v", preview)
	}
	if preview.Source != "dolphin-service-menu" || preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" || preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected file-open identity: %#v", preview)
	}
	if preview.RuntimeMethod != "Launch" || !preview.PortalRequired ||
		preview.PortalInterface != "org.freedesktop.portal.FileChooser" ||
		preview.PortalMethod != "OpenFile" {
		t.Fatalf("unexpected Portal metadata: %#v", preview)
	}
	if preview.FileCount != 1 || !sameStrings(preview.FileURIs, []string{"file:///home/test/Documents/book.xls"}) ||
		preview.SelectedExtension != ".xls" || preview.SelectionMode != "extension-match" {
		t.Fatalf("unexpected file selection: %#v", preview)
	}
	if preview.Action.Type != "runtime-file-open" ||
		!sameStrings(preview.Action.Argv, []string{"xnix-compat-open", "--app", "org.example.ledger", "%U"}) {
		t.Fatalf("unexpected file-open action: %#v", preview.Action)
	}
	if !preview.RuntimeOwned || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.BackendLaunchEnabled || preview.DirectHostFileAccess ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("unexpected file-open safety flags: %#v", preview)
	}

	explicit, err := NewFileOpenPreview([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}}, Provenance{Source: "registry"}, []string{"file:///home/test/Documents/other.txt"}, "org.example.ledger")
	if err != nil {
		t.Fatalf("explicit NewFileOpenPreview returned error: %v", err)
	}
	if explicit.SelectionMode != "explicit-application" || explicit.SelectedExtension != ".txt" {
		t.Fatalf("unexpected explicit selection: %#v", explicit)
	}
	if _, err := NewFileOpenPreview(nil, Provenance{Source: "registry"}, []string{"https://example.invalid/file.xls"}, ""); err == nil {
		t.Fatalf("file-open preview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("file-open preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestFileOpenPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := NewFileOpenPreviewWithOptions([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	}}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, []string{"file:///home/test/Documents/book.xls"}, "", FileOpenOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("NewFileOpenPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("unexpected receipt-backed file-open fields: %#v", preview)
	}
	if preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.BackendLaunchEnabled ||
		preview.DirectHostFileAccess ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed file-open preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("file-open preview exposed activation root: %s", encoded)
	}

	if _, err := NewFileOpenPreviewWithOptions([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}}, Provenance{Source: "registry"}, []string{"file:///home/test/Documents/book.xls"}, "", FileOpenOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("file-open preview accepted a missing activation receipt")
	}
}

func TestDolphinDropPreviewWrapsPortalMediatedFileOpen(t *testing.T) {
	preview, err := NewDolphinDropPreview([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	}}, Provenance{Source: "registry", RegistryName: "test-registry"}, []string{"file:///home/test/Documents/book.xls", "file:///home/test/Documents/tax.xls"}, "")
	if err != nil {
		t.Fatalf("NewDolphinDropPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.dolphin_drop.v1" ||
		preview.RequestType != "dolphin-drop-preview" ||
		preview.Source != "dolphin-drag-and-drop" ||
		preview.DropSurface != "Dolphin" {
		t.Fatalf("unexpected Dolphin drop schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.SelectedExtension != ".xls" ||
		preview.SelectionMode != "extension-match" ||
		preview.FileCount != 2 {
		t.Fatalf("unexpected Dolphin drop identity: %#v", preview)
	}
	if preview.DropOperation != "open-selected-files" ||
		!preview.DropAccepted ||
		preview.Action.Type != "runtime-file-open" ||
		!sameStrings(preview.Action.Argv, []string{"xnix-compat-open", "--app", "org.example.ledger", "%U"}) {
		t.Fatalf("unexpected Dolphin drop action: %#v", preview)
	}
	if !preview.PortalRequired ||
		preview.PortalInterface != "org.freedesktop.portal.FileChooser" ||
		preview.PortalMethod != "OpenFile" {
		t.Fatalf("unexpected Dolphin drop Portal metadata: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.BackendLaunchEnabled ||
		preview.DirectHostFileAccess ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Dolphin drop safety flags: %#v", preview)
	}
	if !sameStrings(preview.BlockedActions, []string{
		"create Portal request objects from a drag preview",
		"grant file permissions from a drag preview",
		"read selected files directly from Dolphin",
		"start compatibility backends from a drag preview",
		"mutate the host root from a drag preview",
		"expose backend implementation details in drag targets",
	}) {
		t.Fatalf("Dolphin drop preview must block direct file reads: %#v", preview.BlockedActions)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Dolphin drop preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestDolphinDropPreviewConsumesActivationReceipt(t *testing.T) {
	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := NewDolphinDropPreviewWithOptions([]Recipe{{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	}}, Provenance{Source: "registry", RegistryName: "test-registry"}, []string{"file:///home/test/Documents/book.xls"}, "", FileOpenOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("NewDolphinDropPreviewWithOptions returned error: %v", err)
	}
	if !preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("unexpected receipt-backed Dolphin drop fields: %#v", preview)
	}
	if preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.BackendLaunchEnabled ||
		preview.DirectHostFileAccess ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed Dolphin drop preview must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("Dolphin drop preview exposed activation root: %s", encoded)
	}
}

func TestRecipeValidationRejectsUnsafeIdentityInput(t *testing.T) {
	cases := []Recipe{
		{ID: "not-reverse-dns", Name: "Example", Icon: "icon", Mode: "automatic"},
		{ID: "org.example.app", Name: "Line\nBreak", Icon: "icon", Mode: "automatic"},
		{ID: "org.example.app", Name: "Example", Icon: "icon", Mode: "unknown"},
		{ID: "org.example.app", Name: "Example", Icon: "icon", Mode: "automatic", SupportedExtensions: []string{"bad"}},
	}
	for _, recipe := range cases {
		if _, err := NewPlan(recipe); err == nil {
			t.Fatalf("NewPlan accepted invalid recipe: %#v", recipe)
		}
	}
}

func settingSectionIDs(sections []SettingsSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func settingFieldValue(sections []SettingsSection, sectionID string, fieldID string) string {
	for _, section := range sections {
		if section.ID != sectionID {
			continue
		}
		for _, field := range section.Fields {
			if field.ID == fieldID {
				return field.Value
			}
		}
	}
	return ""
}

func modeSwitchOptionIDs(options []ModeSwitchOption) []string {
	ids := make([]string, 0, len(options))
	for _, option := range options {
		ids = append(ids, option.ID)
	}
	return ids
}

func countSelectedModes(options []ModeSwitchOption) int {
	count := 0
	for _, option := range options {
		if option.Selected {
			count++
		}
	}
	return count
}

func countRequestedModes(options []ModeSwitchOption) int {
	count := 0
	for _, option := range options {
		if option.Requested {
			count++
		}
	}
	return count
}

func permissionReviewIDs(permissions []PermissionReviewEntry) []string {
	ids := make([]string, 0, len(permissions))
	for _, permission := range permissions {
		ids = append(ids, permission.ID)
	}
	return ids
}

func permissionDecision(permissions []PermissionReviewEntry, permissionID string) string {
	for _, permission := range permissions {
		if permission.ID == permissionID {
			return permission.Decision
		}
	}
	return ""
}

func desktopResourceBridgeIDs(resources []DesktopResourceBridgeResource) []string {
	ids := make([]string, 0, len(resources))
	for _, resource := range resources {
		ids = append(ids, resource.ID)
	}
	return ids
}

func reviewFlowStepIDs(steps []ReviewFlowStep) []string {
	ids := make([]string, 0, len(steps))
	for _, step := range steps {
		ids = append(ids, step.ID)
	}
	return ids
}

func sameStrings(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}
	return true
}
