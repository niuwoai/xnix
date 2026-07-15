package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDECenterPagePreviewComposesSummaryDeckAndSettings(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	preview, err := NewKDECenterPagePreview(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("NewKDECenterPagePreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.kde_center_page.v1" ||
		preview.RequestType != "kde-center-page-preview" ||
		preview.PageType != "compatibility-center-application-page" ||
		preview.Source != "compatibility-center-preview+kde-action-card-deck-preview+settings-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDECenterPage" ||
		preview.ReadMethod != "GetKDECenterPagePreview" {
		t.Fatalf("unexpected KDE center page schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		len(preview.LauncherCommand) != 4 {
		t.Fatalf("unexpected KDE center page identity: %#v", preview)
	}
	if preview.Header.Title != "Example Ledger" ||
		preview.Header.Subtitle != "Compatibility Center application details" ||
		preview.Header.Badge != "Review ready" ||
		preview.Header.BadgeTone != "warning" ||
		preview.Header.PrimaryActionLabel != "Review required gates" ||
		preview.Header.PrimaryActionTarget != "compatibility-center-gates" ||
		!preview.Header.PrimaryActionEnabled ||
		preview.Header.BackendDetailsExposed {
		t.Fatalf("unexpected page header: %#v", preview.Header)
	}
	if preview.ApplicationSummary.ApplicationID != "org.example.ledger" ||
		preview.ApplicationSummary.CompatibilityState != "registered" ||
		preview.ApplicationSummary.CompatibilityLabel != "Registered" ||
		preview.ApplicationSummary.DiagnosticsState != "not-run" ||
		preview.ApplicationSummary.RuntimeMode != "Automatic" ||
		len(preview.ApplicationSummary.SupportedExtensions) != 1 ||
		preview.ApplicationSummary.SupportedExtensions[0] != ".xls" ||
		preview.ApplicationSummary.KnownIssueCount != 0 ||
		preview.ApplicationSummary.RepairRecordCount != 0 ||
		preview.ApplicationSummary.ActionExecutionEnabled ||
		preview.ApplicationSummary.RepairExecutionEnabled ||
		preview.ApplicationSummary.BackendLaunchEnabled ||
		preview.ApplicationSummary.SettingsPersistenceEnabled ||
		preview.ApplicationSummary.HostRootModified ||
		preview.ApplicationSummary.BackendDetailsExposed {
		t.Fatalf("unexpected application summary: %#v", preview.ApplicationSummary)
	}
	if preview.ActionDeck.RequestType != "kde-action-card-deck-preview" ||
		preview.ActionDeck.DeckType != "compatibility-center-kde-action-card-deck" ||
		preview.ActionDeck.CardCount != 7 ||
		preview.ActionDeck.WaitingCardCount != 7 ||
		preview.ActionDeck.DeferredCardCount != 0 ||
		preview.ActionDeck.RejectedCardCount != 0 ||
		preview.ActionDeck.AIAnalysisCardCount != 1 ||
		preview.ActionDeck.NavigationActionCount != 29 ||
		preview.ActionDeck.DisabledActionCount != 21 ||
		preview.ActionDeck.PrimaryCardID != "org.example.ledger:review-launcher-action:card" ||
		len(preview.ActionDeck.CardIDs) != 7 ||
		!preview.ActionDeck.ActionQueueCreated ||
		preview.ActionDeck.ActionQueuePersisted ||
		!preview.ActionDeck.DeckPreviewCreated ||
		preview.ActionDeck.DeckPersisted ||
		preview.ActionDeck.CardsPersisted ||
		preview.ActionDeck.CardActionsEnabled ||
		preview.ActionDeck.RuntimeLaunchApproval ||
		preview.ActionDeck.LaunchAllowed ||
		preview.ActionDeck.ExecutionStarted ||
		preview.ActionDeck.BackendDetailsExposed {
		t.Fatalf("unexpected action deck summary: %#v", preview.ActionDeck)
	}
	if preview.ActionDeck.AIAnalysis == nil ||
		preview.ActionDeck.AIAnalysis.Source != "dolphin-ai-analysis-preview" ||
		preview.ActionDeck.AIAnalysis.Disclosure != "count-and-extension-only" ||
		!preview.ActionDeck.AIAnalysis.SafeForAIDiagnostics ||
		preview.ActionDeck.AIAnalysis.AIProviderCallEnabled ||
		preview.ActionDeck.AIAnalysis.NetworkRequired ||
		preview.ActionDeck.AIAnalysis.FileContentRead ||
		preview.ActionDeck.AIAnalysis.FilePathsExposed ||
		preview.ActionDeck.AIAnalysis.RequestObjectCreated ||
		preview.ActionDeck.AIAnalysis.PermissionGranted ||
		preview.ActionDeck.AIAnalysis.BackendLaunchEnabled {
		t.Fatalf("unexpected center page AI analysis link: %#v", preview.ActionDeck.AIAnalysis)
	}
	if preview.SettingsSnapshot.RequestType != "settings-preview" ||
		preview.SettingsSnapshot.SettingsState != "planned" ||
		preview.SettingsSnapshot.SectionCount != 5 ||
		len(preview.SettingsSnapshot.Sections) != 5 ||
		preview.SettingsSnapshot.SettingsPersisted ||
		preview.SettingsSnapshot.SettingsPersistenceEnabled ||
		preview.SettingsSnapshot.HostRootModified ||
		preview.SettingsSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected settings snapshot: %#v", preview.SettingsSnapshot)
	}
	if preview.NavigationCount != 4 ||
		len(preview.Navigation) != 4 ||
		preview.PrimaryNavigationTarget != "compatibility-center-gates" ||
		preview.Navigation[0].ID != "overview" ||
		preview.Navigation[1].ID != "actions" ||
		preview.Navigation[2].ID != "settings" ||
		preview.Navigation[3].ID != "diagnostics" {
		t.Fatalf("unexpected navigation: %#v", preview.Navigation)
	}
	for _, item := range preview.Navigation {
		if !item.Enabled || !item.NavigationOnly || item.MutatesRuntime || item.StartsProgram {
			t.Fatalf("navigation item must stay navigation-only: %#v", item)
		}
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.UserVisible ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.PagePreviewCreated ||
		preview.PagePersisted || preview.DeckPersisted ||
		preview.CardsPersisted || preview.CardActionsEnabled ||
		preview.SettingsPersisted || preview.SettingsPersistenceEnabled ||
		preview.NotificationsSent || preview.ResourceGrantCreated ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected center page safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE center page from preview state") ||
		!containsString(preview.BlockedActions, "persist compatibility settings from center page preview") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from center page preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}
}

func TestKDECenterPagePreviewValidationAndDecisionStates(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}

	rejectedPreview, err := NewKDECenterPagePreview(recipe, Provenance{}, "rejected", nil)
	if err != nil {
		t.Fatalf("rejected NewKDECenterPagePreview returned error: %v", err)
	}
	if rejectedPreview.Header.Badge != "Needs guidance" ||
		rejectedPreview.Header.BadgeTone != "critical" ||
		rejectedPreview.ActionDeck.RejectedCardCount != 7 ||
		rejectedPreview.ActionDeck.WaitingCardCount != 0 ||
		rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected center page: %#v", rejectedPreview)
	}

	deferredPreview, err := NewKDECenterPagePreview(recipe, Provenance{}, "deferred", nil)
	if err != nil {
		t.Fatalf("deferred NewKDECenterPagePreview returned error: %v", err)
	}
	if deferredPreview.Header.Badge != "Deferred" ||
		deferredPreview.Header.BadgeTone != "neutral" ||
		deferredPreview.ActionDeck.DeferredCardCount != 7 ||
		deferredPreview.ActionDeck.WaitingCardCount != 0 ||
		deferredPreview.PagePersisted ||
		deferredPreview.SettingsPersisted {
		t.Fatalf("unexpected deferred center page: %#v", deferredPreview)
	}

	if _, err := NewKDECenterPagePreview(recipe, Provenance{}, "invalid", nil); err == nil {
		t.Fatalf("NewKDECenterPagePreview accepted an invalid decision")
	}
	if _, err := NewKDECenterPagePreview(recipe, Provenance{}, "approved\nbad", nil); err == nil {
		t.Fatalf("NewKDECenterPagePreview accepted a multiline decision")
	}
	if _, err := NewKDECenterPagePreview(recipe, Provenance{}, "approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("NewKDECenterPagePreview accepted a non-file URI")
	}

	preview, err := NewKDECenterPagePreview(recipe, Provenance{}, "reviewed", nil)
	if err != nil {
		t.Fatalf("NewKDECenterPagePreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKDECenterPageSectionsPreviewDefinesReadOnlyNavigation(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	preview, err := NewKDECenterPageSectionsPreview(recipe, Provenance{}, "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("NewKDECenterPageSectionsPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.kde_center_page_sections.v1" ||
		preview.RequestType != "kde-center-page-sections-preview" ||
		preview.PageType != "compatibility-center-application-page" ||
		preview.Source != "kde-center-page-preview" ||
		preview.RuntimeMethod != "GetKDECenterPageSections" ||
		preview.ReadMethod != "GetKDECenterPageSectionsPreview" {
		t.Fatalf("unexpected KDE center page sections schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.SectionCount != 4 ||
		preview.ReadOnlySectionCount != 4 ||
		preview.NavigationOnlySectionCount != 4 ||
		preview.ExecutableSectionCount != 0 ||
		preview.AIAnalysisSectionCount != 1 ||
		preview.PrimarySectionID != "overview" {
		t.Fatalf("unexpected KDE center page sections identity: %#v", preview)
	}
	if preview.AIAnalysis == nil ||
		preview.AIAnalysis.Source != "dolphin-ai-analysis-preview" ||
		preview.AIAnalysis.Disclosure != "count-and-extension-only" ||
		!preview.AIAnalysis.SafeForAIDiagnostics ||
		preview.AIAnalysis.AIProviderCallEnabled ||
		preview.AIAnalysis.NetworkRequired ||
		preview.AIAnalysis.FileContentRead ||
		preview.AIAnalysis.FilePathsExposed ||
		preview.AIAnalysis.RequestObjectCreated ||
		preview.AIAnalysis.PermissionGranted ||
		preview.AIAnalysis.BackendLaunchEnabled {
		t.Fatalf("unexpected sections AI analysis link: %#v", preview.AIAnalysis)
	}
	wantMethods := map[string]string{
		"overview":    "GetCompatibilityCenterSummary",
		"actions":     "GetCompatibilityActionQueue",
		"settings":    "GetCompatibilitySettings",
		"diagnostics": "GetAIDiagnosticInput",
	}
	wantModels := map[string]string{
		"overview":    "compatibility-center-summary",
		"actions":     "compatibility-center-action-queue",
		"settings":    "settings-model",
		"diagnostics": "ai-diagnostic-input",
	}
	for _, section := range preview.Sections {
		if section.RuntimeMethod != wantMethods[section.ID] ||
			section.ReadModel != wantModels[section.ID] ||
			!section.NavigationOnly ||
			!section.ReadOnly ||
			section.MutatesRuntime ||
			section.StartsProgram ||
			section.SettingsPersisted ||
			section.BackendDetailsExposed {
			t.Fatalf("unexpected section contract: %#v", section)
		}
		if section.ID == "diagnostics" && section.AIAnalysis == nil {
			t.Fatalf("diagnostics section should expose AI analysis link")
		}
		if section.ID != "diagnostics" && section.AIAnalysis != nil {
			t.Fatalf("%s section should not expose AI analysis link: %#v", section.ID, section.AIAnalysis)
		}
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.UserVisible ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.SectionsPreviewCreated ||
		preview.SectionsPersisted || preview.SectionActionsEnabled ||
		preview.SettingsPersisted || preview.SettingsPersistenceEnabled ||
		preview.NotificationsSent || preview.ResourceGrantCreated ||
		preview.RuntimeLaunchApproval || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected section safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist KDE center page sections from preview state") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from page sections") ||
		!containsString(preview.BlockedActions, "start compatibility profile from page sections") {
		t.Fatalf("unexpected section blocked actions: %#v", preview.BlockedActions)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page sections preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKDECenterPageSectionDetailPreviewRoutesSelectedReadModel(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	preview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "settings", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.kde_center_page_section_detail.v1" ||
		preview.RequestType != "kde-center-page-section-detail-preview" ||
		preview.PageType != "compatibility-center-application-page" ||
		preview.Source != "kde-center-page-sections-preview" ||
		preview.RuntimeMethod != "GetKDECenterPageSectionDetail" ||
		preview.ReadMethod != "GetKDECenterPageSectionDetailPreview" {
		t.Fatalf("unexpected KDE center page section detail schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.SectionID != "settings" ||
		preview.SectionLabel != "Settings" ||
		preview.SectionTarget != "compatibility-settings" ||
		preview.SectionState != "planned" ||
		preview.SectionRuntimeMethod != "GetCompatibilitySettings" ||
		preview.SectionReadModel != "settings-model" {
		t.Fatalf("unexpected section detail identity: %#v", preview)
	}
	if got := strings.Join(preview.AvailableSectionIDs, ","); got != "overview,actions,settings,diagnostics" {
		t.Fatalf("unexpected section ids: %#v", preview.AvailableSectionIDs)
	}
	if !preview.ReadOnlyNavigation || !preview.DetailPreviewCreated ||
		preview.DetailPersisted || preview.SectionActionsEnabled ||
		preview.SettingsPersisted || preview.SettingsPersistenceEnabled ||
		preview.NotificationsSent || preview.ResourceGrantCreated ||
		preview.RuntimeLaunchApproval || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.BackendDetailsExposed {
		t.Fatalf("unexpected section detail safety flags: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.UserVisible ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch {
		t.Fatalf("unexpected section detail ownership flags: %#v", preview)
	}
	diagnosticsPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "diagnostics", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("diagnostics NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if diagnosticsPreview.SectionID != "diagnostics" ||
		diagnosticsPreview.SectionRuntimeMethod != "GetAIDiagnosticInput" ||
		diagnosticsPreview.SectionReadModel != "ai-diagnostic-input" ||
		diagnosticsPreview.AIAnalysis == nil ||
		diagnosticsPreview.AIAnalysis.Source != "dolphin-ai-analysis-preview" ||
		diagnosticsPreview.AIAnalysis.Disclosure != "count-and-extension-only" ||
		diagnosticsPreview.AIAnalysis.AIProviderCallEnabled ||
		diagnosticsPreview.AIAnalysis.NetworkRequired ||
		diagnosticsPreview.AIAnalysis.FileContentRead ||
		diagnosticsPreview.AIAnalysis.FilePathsExposed ||
		diagnosticsPreview.AIAnalysis.RequestObjectCreated ||
		diagnosticsPreview.AIAnalysis.PermissionGranted ||
		diagnosticsPreview.AIAnalysis.BackendLaunchEnabled ||
		diagnosticsPreview.SectionActionsEnabled ||
		diagnosticsPreview.RequestObjectsCreated ||
		diagnosticsPreview.ExecutionStarted {
		t.Fatalf("unexpected diagnostics section detail: %#v", diagnosticsPreview)
	}
	if !containsString(preview.BlockedActions, "persist KDE center page section detail from preview state") ||
		!containsString(preview.BlockedActions, "create Runtime request objects from section detail") ||
		!containsString(preview.BlockedActions, "start compatibility profile from section detail") {
		t.Fatalf("unexpected section detail blocked actions: %#v", preview.BlockedActions)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page section detail preview exposes forbidden term %q: %s", forbidden, text)
		}
	}

	if _, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "unknown", "approved", nil); err == nil {
		t.Fatalf("NewKDECenterPageSectionDetailPreview accepted an unknown section")
	}
	if _, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "settings\nbad", "approved", nil); err == nil {
		t.Fatalf("NewKDECenterPageSectionDetailPreview accepted a multiline section")
	}
}
