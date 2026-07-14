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
		preview.KnownIssueCount != 0 || preview.RepairRecordCount != 0 || preview.PendingReviewCount != 0 {
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
