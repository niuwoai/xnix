package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationManifestPreviewAggregatesKDEContract(t *testing.T) {
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

	preview, err := plan.DesktopActivationManifestPreview()
	if err != nil {
		t.Fatalf("DesktopActivationManifestPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_manifest.v1" ||
		preview.RequestType != "desktop-activation-manifest-preview" ||
		preview.ManifestType != "kde-desktop-activation-manifest" ||
		preview.Source != "desktop-activation-bundle-preview+kde-entrypoints-preview" ||
		preview.RuntimeMethod != "GetDesktopActivationManifest" ||
		preview.ReadMethod != "GetDesktopActivationManifestPreview" {
		t.Fatalf("unexpected desktop activation manifest schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected desktop activation manifest identity: %#v", preview)
	}
	if preview.Bundle.RequestType != "desktop-activation-bundle-preview" ||
		preview.Bundle.MaterialCount != 9 ||
		!preview.Bundle.StandardDesktopEntry ||
		preview.Bundle.DesktopFilesWritten ||
		preview.Bundle.HostRootModified ||
		preview.Bundle.BackendDetailsExposed {
		t.Fatalf("unexpected manifest bundle summary: %#v", preview.Bundle)
	}
	if got, want := preview.EntryPointIDs, []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "settings"}; !sameStrings(got, want) {
		t.Fatalf("EntryPointIDs = %#v, want %#v", got, want)
	}
	if got, want := preview.ActivationMaterialIDs, []string{"launcher", "file-association", "desktop-icon", "task-manager", "kwin-window-rule", "system-tray", "notification-center", "unified-settings", "compatibility-center"}; !sameStrings(got, want) {
		t.Fatalf("ActivationMaterialIDs = %#v, want %#v", got, want)
	}
	if got, want := preview.ContractSectionIDs, []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "unified-settings"}; !sameStrings(got, want) {
		t.Fatalf("ContractSectionIDs = %#v, want %#v", got, want)
	}
	if preview.EntryPointCount != 7 ||
		preview.VisibleEntryPointCount != 7 ||
		preview.ActivationMaterialCount != 9 ||
		preview.ContractSectionCount != 7 {
		t.Fatalf("unexpected contract counts: %#v", preview)
	}
	for _, entrypoint := range preview.EntryPoints {
		if !entrypoint.Visible || !entrypoint.Planned ||
			entrypoint.WritesHost || entrypoint.StartsBackend || entrypoint.BackendDetailsExposed {
			t.Fatalf("entrypoint has unsafe flags: %#v", entrypoint)
		}
	}
	for _, material := range preview.ActivationMaterials {
		if !material.UserVisible || !material.RequiredForNormalApp ||
			material.WritesHost || material.StartsBackend || material.BackendDetailsExposed {
			t.Fatalf("activation material has unsafe flags: %#v", material)
		}
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.UserVisible || !preview.OfficialDesktopOnly ||
		!preview.StableDesktopContract || !preview.NormalApplicationSurface ||
		!preview.StandardDesktopEntry || !preview.SevenEntryPointContract ||
		!preview.PortalMediatedFileAccess {
		t.Fatalf("unexpected manifest readiness flags: %#v", preview)
	}
	if preview.DesktopFilesWritten || preview.MIMEAppsWritten ||
		preview.ManifestWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.LaunchEnabled || preview.BackendLaunchEnabled ||
		preview.ExecutionStarted || preview.HostRootModified ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed || preview.RawWindowsExecutableExposed ||
		preview.CompatibilityStorageExposed {
		t.Fatalf("desktop activation manifest gate unexpectedly open: %#v", preview)
	}
	if len(preview.BlockedActions) != 11 ||
		preview.BlockedActions[0] != "write desktop activation files from manifest preview" {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
	if err := validateNoBackendTerms(preview, "desktop activation manifest preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop activation manifest preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
