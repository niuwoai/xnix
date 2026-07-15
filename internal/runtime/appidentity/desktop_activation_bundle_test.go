package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationBundlePreviewAggregatesNormalApplicationMaterials(t *testing.T) {
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

	preview, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		t.Fatalf("DesktopActivationBundlePreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_bundle.v1" ||
		preview.RequestType != "desktop-activation-bundle-preview" ||
		preview.PlanType != "normal-linux-application-activation" ||
		preview.RuntimeMethod != "GetDesktopActivationBundlePreview" {
		t.Fatalf("unexpected desktop activation bundle schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop activation bundle identity: %#v", preview)
	}
	if preview.DesktopEntryPreview == "" ||
		!strings.Contains(preview.DesktopEntryPreview, "Exec=xnix-compat-launch --app org.example.ledger %U\n") ||
		!strings.Contains(preview.DesktopEntryPreview, "MimeType=application/x-xnix-abc;application/x-xnix-xls;\n") {
		t.Fatalf("unexpected desktop entry preview: %q", preview.DesktopEntryPreview)
	}
	if preview.MIMEAppsPreview == "" ||
		!strings.Contains(preview.MIMEAppsPreview, "application/x-xnix-abc=xnix-org.example.ledger.desktop\n") ||
		!strings.Contains(preview.MIMEAppsPreview, "application/x-xnix-xls=xnix-org.example.ledger.desktop;\n") {
		t.Fatalf("unexpected MIME apps preview: %q", preview.MIMEAppsPreview)
	}
	if preview.DesktopIcon.RequestType != "desktop-icon-preview" ||
		preview.WindowIdentity.SchemaVersion != "xnix.runtime.window_identity.v1" ||
		preview.TrayStatus.StatusType != "tray-status-preview" ||
		preview.Notification.RequestType != "desktop-notification-preview" ||
		preview.Settings.RequestType != "settings-preview" {
		t.Fatalf("unexpected aggregated previews: %#v", preview)
	}
	if got, want := preview.MaterialIDs, []string{"launcher", "file-association", "desktop-icon", "task-manager", "kwin-window-rule", "system-tray", "notification-center", "unified-settings", "compatibility-center"}; !sameStrings(got, want) {
		t.Fatalf("MaterialIDs = %#v, want %#v", got, want)
	}
	if preview.MaterialCount != 9 || len(preview.Materials) != 9 {
		t.Fatalf("unexpected material count: %#v", preview)
	}
	for _, material := range preview.Materials {
		if !material.UserVisible || !material.RequiredForNormalApp ||
			material.WritesHost || material.StartsBackend || material.BackendDetailsExposed {
			t.Fatalf("material has unsafe flags: %#v", material)
		}
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.NormalApplicationSurface ||
		!preview.StandardDesktopEntry || !preview.FileAssociationReady ||
		!preview.TaskManagerIdentityReady || !preview.KWinIdentityReady ||
		!preview.TrayStatusReady || !preview.NotificationReady ||
		!preview.SettingsReady || !preview.CompatibilityCenterReady {
		t.Fatalf("unexpected readiness flags: %#v", preview)
	}
	if preview.DesktopFilesWritten || preview.MIMEAppsWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.LaunchEnabled || preview.BackendLaunchEnabled ||
		preview.ExecutionStarted || preview.HostRootModified ||
		preview.BackendDetailsExposed || preview.RawWindowsExecutableExposed ||
		preview.CompatibilityStorageExposed {
		t.Fatalf("desktop activation bundle gate unexpectedly open: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop activation bundle preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
