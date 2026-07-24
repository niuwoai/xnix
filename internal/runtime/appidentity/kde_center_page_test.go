package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
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
		preview.Source != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview" ||
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
	if preview.ApplicationReadinessEvidence.SchemaVersion != "xnix.runtime.application_readiness.v1" ||
		preview.ApplicationReadinessEvidence.RequestType != "application-readiness-preview" ||
		preview.ApplicationReadinessEvidence.GraphType != "runtime-application-readiness-evidence-graph" ||
		preview.ApplicationReadinessEvidence.RuntimeMethod != "GetApplicationReadiness" ||
		preview.ApplicationReadinessEvidence.Application.ID != "org.example.ledger" ||
		preview.ApplicationReadinessEvidence.NodeCount != 7 ||
		preview.ApplicationReadinessEvidence.Ready ||
		preview.ApplicationReadinessEvidence.LaunchAllowed ||
		preview.ApplicationReadinessEvidence.LaunchEnabled ||
		preview.ApplicationReadinessEvidence.ExecutionRequestCreated ||
		preview.ApplicationReadinessEvidence.ExecutionStarted ||
		preview.ApplicationReadinessEvidence.BackendLaunchEnabled ||
		preview.ApplicationReadinessEvidence.BackendProcessStarted ||
		preview.ApplicationReadinessEvidence.RealPortalTransportEnabled ||
		preview.ApplicationReadinessEvidence.RequestObjectCreated ||
		preview.ApplicationReadinessEvidence.PermissionGranted ||
		preview.ApplicationReadinessEvidence.SnapshotCreated ||
		preview.ApplicationReadinessEvidence.RestoreExecuted ||
		preview.ApplicationReadinessEvidence.HostRootModified ||
		preview.ApplicationReadinessEvidence.NetworkRequired ||
		preview.ApplicationReadinessEvidence.PrivilegedContainerRequired ||
		preview.ApplicationReadinessEvidence.StateRootPathExposed ||
		preview.ApplicationReadinessEvidence.BackendDetailsExposed ||
		preview.ApplicationReadinessEvidence.RawCommandExposed ||
		preview.ApplicationReadinessEvidence.RawExecutableExposed {
		t.Fatalf("unexpected application readiness evidence: %#v", preview.ApplicationReadinessEvidence)
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
	if preview.ActionDependencyGraph.RequestType != "kde-action-dependency-graph-preview" ||
		preview.ActionDependencyGraph.GraphType != "compatibility-center-action-dependency-graph" ||
		preview.ActionDependencyGraph.RuntimeMethod != "GetKDEActionDependencyGraph" ||
		preview.ActionDependencyGraph.ReadMethod != "GetKDEActionDependencyGraphPreview" ||
		preview.ActionDependencyGraph.ActionNodeCount != 7 ||
		preview.ActionDependencyGraph.EvidenceNodeCount != 35 ||
		preview.ActionDependencyGraph.GateNodeCount != 7 ||
		preview.ActionDependencyGraph.NodeCount != 49 ||
		preview.ActionDependencyGraph.EdgeCount != 42 ||
		preview.ActionDependencyGraph.MissingEvidenceCount != 35 ||
		preview.ActionDependencyGraph.BlockedActionCount != 7 ||
		len(preview.ActionDependencyGraph.MissingEvidenceIDs) != 35 ||
		!containsString(preview.ActionDependencyGraph.BlockedActions, "review-file-manager-action") ||
		!preview.ActionDependencyGraph.ReceiptValidation.RejectsMismatchedAppID ||
		!preview.ActionDependencyGraph.ReceiptValidation.RejectsMalformedOperation ||
		!preview.ActionDependencyGraph.ReceiptValidation.RejectsPathEscapeEvidence ||
		!preview.ActionDependencyGraph.ReceiptValidation.RejectsUnsafeSideEffects ||
		!preview.ActionDependencyGraph.DependencyGraphCreated ||
		preview.ActionDependencyGraph.DependencyGraphPersisted ||
		preview.ActionDependencyGraph.RequestObjectsCreated ||
		preview.ActionDependencyGraph.PermissionGrantCreated ||
		preview.ActionDependencyGraph.SettingsPersisted ||
		preview.ActionDependencyGraph.RuntimeLaunchApproval ||
		preview.ActionDependencyGraph.LaunchAllowed ||
		preview.ActionDependencyGraph.ExecutionStarted ||
		preview.ActionDependencyGraph.BackendProcessStarted ||
		preview.ActionDependencyGraph.HostRootModified ||
		preview.ActionDependencyGraph.BackendDetailsExposed {
		t.Fatalf("unexpected action dependency graph summary: %#v", preview.ActionDependencyGraph)
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
	if preview.BackendSelectionSnapshot.RequestType != "backend-selection-preview" ||
		preview.BackendSelectionSnapshot.RuntimeMethod != "GetBackendSelectionPlan" ||
		preview.BackendSelectionSnapshot.RecommendedProfileID != "local-compatibility" ||
		preview.BackendSelectionSnapshot.CandidateCount != 2 ||
		preview.BackendSelectionSnapshot.SelectionCommitted ||
		preview.BackendSelectionSnapshot.SelectionChangeEnabled ||
		preview.BackendSelectionSnapshot.BackendLaunchEnabled ||
		preview.BackendSelectionSnapshot.EnvironmentCreated ||
		preview.BackendSelectionSnapshot.HostRootModified ||
		preview.BackendSelectionSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected backend selection snapshot: %#v", preview.BackendSelectionSnapshot)
	}
	if preview.ActivationStatusSnapshot.RequestType != "desktop-activation-status-preview" ||
		preview.ActivationStatusSnapshot.RuntimeMethod != "GetDesktopActivationStatus" ||
		preview.ActivationStatusSnapshot.Renderer != "xnix-runtime-go desktop-activation-status-preview" ||
		preview.ActivationStatusSnapshot.ActivationState != "ready-for-runtime-commit" ||
		preview.ActivationStatusSnapshot.CommitEnabled ||
		preview.ActivationStatusSnapshot.LaunchEnabled ||
		preview.ActivationStatusSnapshot.HostRootModified ||
		preview.ActivationStatusSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected activation status snapshot: %#v", preview.ActivationStatusSnapshot)
	}
	if preview.ExecutionReadinessSnapshot.RequestType != "execution-readiness-preview" ||
		preview.ExecutionReadinessSnapshot.RuntimeMethod != "GetExecutionReadiness" ||
		preview.ExecutionReadinessSnapshot.ExecutionState != "blocked" ||
		preview.ExecutionReadinessSnapshot.OverallStatus != "not-ready" ||
		preview.ExecutionReadinessSnapshot.GateCount != 5 ||
		preview.ExecutionReadinessSnapshot.RequiredGateCount != 2 ||
		preview.ExecutionReadinessSnapshot.PendingGateCount != 1 ||
		preview.ExecutionReadinessSnapshot.BlockedGateCount != 1 ||
		!preview.ExecutionReadinessSnapshot.DesktopEntryLaunchVisible ||
		preview.ExecutionReadinessSnapshot.LaunchAllowed ||
		preview.ExecutionReadinessSnapshot.LaunchEnabled ||
		preview.ExecutionReadinessSnapshot.ExecutionRequestCreated ||
		preview.ExecutionReadinessSnapshot.BackendBindingReady ||
		!preview.ExecutionReadinessSnapshot.PortalPolicyRequired ||
		!preview.ExecutionReadinessSnapshot.SnapshotRequired ||
		!preview.ExecutionReadinessSnapshot.UserActionRequired ||
		preview.ExecutionReadinessSnapshot.HostRootModified ||
		preview.ExecutionReadinessSnapshot.NetworkRequired ||
		preview.ExecutionReadinessSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected execution readiness snapshot: %#v", preview.ExecutionReadinessSnapshot)
	}
	if preview.LaunchIntentSnapshot.RequestType != "launch-intent-preview" ||
		preview.LaunchIntentSnapshot.IntentType != "runtime-launch-intent" ||
		preview.LaunchIntentSnapshot.Source != "desktop-launcher" ||
		preview.LaunchIntentSnapshot.RuntimeMethod != "Launch" ||
		preview.LaunchIntentSnapshot.ReadMethod != "GetLaunchIntent" ||
		preview.LaunchIntentSnapshot.FileCount != 1 ||
		!preview.LaunchIntentSnapshot.PortalRequired ||
		!preview.LaunchIntentSnapshot.SnapshotRequired ||
		!preview.LaunchIntentSnapshot.StandardDesktopEntry ||
		!preview.LaunchIntentSnapshot.LaunchUsesRuntime ||
		preview.LaunchIntentSnapshot.LaunchAllowed ||
		preview.LaunchIntentSnapshot.LaunchEnabled ||
		preview.LaunchIntentSnapshot.ExecutionRequestCreated ||
		preview.LaunchIntentSnapshot.ExecutionStarted ||
		preview.LaunchIntentSnapshot.BackendBindingReady ||
		preview.LaunchIntentSnapshot.RequestObjectCreated ||
		preview.LaunchIntentSnapshot.PermissionGranted ||
		preview.LaunchIntentSnapshot.HostRootModified ||
		preview.LaunchIntentSnapshot.NetworkRequired ||
		preview.LaunchIntentSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected launch intent snapshot: %#v", preview.LaunchIntentSnapshot)
	}
	if preview.WindowIdentitySnapshot.SchemaVersion != "xnix.runtime.window_identity.v1" ||
		preview.WindowIdentitySnapshot.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.WindowIdentitySnapshot.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.WindowIdentitySnapshot.WindowKind != "compatibility-application" ||
		preview.WindowIdentitySnapshot.ClassGroup != "xnix-compatibility" ||
		preview.WindowIdentitySnapshot.ResourceName != "org.example.ledger" ||
		preview.WindowIdentitySnapshot.TitleHint != "Example Ledger" ||
		preview.WindowIdentitySnapshot.TaskManagerGroupingKey != "org.example.ledger" ||
		!preview.WindowIdentitySnapshot.TaskManagerPinningAllowed ||
		!preview.WindowIdentitySnapshot.TaskManagerRestoreAllowed ||
		preview.WindowIdentitySnapshot.TaskManagerSkipTaskbar ||
		!preview.WindowIdentitySnapshot.TaskManagerShowInSwitcher ||
		!preview.WindowIdentitySnapshot.PreferExistingWindow ||
		preview.WindowIdentitySnapshot.KWinScriptRole != "identity-and-layout" ||
		preview.WindowIdentitySnapshot.KWinPlacement != "normal-window" ||
		!preview.WindowIdentitySnapshot.WindowManagerPolicyOnly ||
		!preview.WindowIdentitySnapshot.RuntimeOwnsBackendPolicy ||
		preview.WindowIdentitySnapshot.TaskManagerEntryActive ||
		preview.WindowIdentitySnapshot.KWinRuleApplied ||
		preview.WindowIdentitySnapshot.HostRootModified ||
		preview.WindowIdentitySnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected window identity snapshot: %#v", preview.WindowIdentitySnapshot)
	}
	if preview.FileAssociationSnapshot.PlanType != "file-association-plan" ||
		preview.FileAssociationSnapshot.AssociationType != "desktop-file-association" ||
		preview.FileAssociationSnapshot.RuntimeMethod != "GetFileAssociationPlan" ||
		preview.FileAssociationSnapshot.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileAssociationSnapshot.MIMEAppsPath != "usr/share/applications/mimeapps.list" ||
		preview.FileAssociationSnapshot.MIMETypeCount != 1 ||
		len(preview.FileAssociationSnapshot.MIMETypes) != 1 ||
		preview.FileAssociationSnapshot.MIMETypes[0] != "application/x-xnix-xls" ||
		preview.FileAssociationSnapshot.FileOpenCommand != "xnix-compat-open" ||
		preview.FileAssociationSnapshot.FileOpenArgument != "%U" ||
		!preview.FileAssociationSnapshot.StandardMIMEAppsList ||
		!preview.FileAssociationSnapshot.StagedRootOnly ||
		preview.FileAssociationSnapshot.OverwriteExistingMIMEApps ||
		!preview.FileAssociationSnapshot.PortalRequiredForFileOpen ||
		!preview.FileAssociationSnapshot.FileAssociationReady ||
		!preview.FileAssociationSnapshot.FileOpenPreviewAvailable ||
		preview.FileAssociationSnapshot.DirectHostFileAccess ||
		preview.FileAssociationSnapshot.RequestObjectCreated ||
		preview.FileAssociationSnapshot.PermissionGranted ||
		preview.FileAssociationSnapshot.FilesWritten ||
		preview.FileAssociationSnapshot.MIMEAppsWritten ||
		preview.FileAssociationSnapshot.HostRootModified ||
		preview.FileAssociationSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected file association snapshot: %#v", preview.FileAssociationSnapshot)
	}
	if preview.TrayStatusSnapshot.StatusType != "tray-status-preview" ||
		preview.TrayStatusSnapshot.RuntimeMethod != "GetTrayStatus" ||
		preview.TrayStatusSnapshot.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.TrayStatusSnapshot.RegisteredApplicationCount != 1 ||
		preview.TrayStatusSnapshot.ActiveApplicationCount != 0 ||
		preview.TrayStatusSnapshot.AttentionRequiredCount != 0 ||
		preview.TrayStatusSnapshot.CompatibilityState != "ready" ||
		preview.TrayStatusSnapshot.CompatibilityLabel != "Ready" ||
		preview.TrayStatusSnapshot.TrayBridgeState != "planned" ||
		preview.TrayStatusSnapshot.TrayBridgeLabel != "Tray bridge is planned" ||
		preview.TrayStatusSnapshot.BridgedTrayApplicationCount != 0 ||
		preview.TrayStatusSnapshot.ActionCount != 2 ||
		!sameStrings(preview.TrayStatusSnapshot.Actions, []string{"open-compatibility-center", "open-settings"}) ||
		!preview.TrayStatusSnapshot.UserVisible ||
		!preview.TrayStatusSnapshot.TrayStatusReady ||
		preview.TrayStatusSnapshot.LiveBackendBridgeEnabled ||
		preview.TrayStatusSnapshot.BridgeConfigurationPersisted ||
		preview.TrayStatusSnapshot.HostRootModified ||
		preview.TrayStatusSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected tray status snapshot: %#v", preview.TrayStatusSnapshot)
	}
	if preview.NotificationSnapshot.RequestType != "desktop-notification-preview" ||
		preview.NotificationSnapshot.RuntimeMethod != "GetNotificationPlan" ||
		preview.NotificationSnapshot.Source != "runtime-event" ||
		preview.NotificationSnapshot.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.NotificationSnapshot.EventType != "approval-required" ||
		preview.NotificationSnapshot.NotificationID != "org.example.ledger.approval-required" ||
		preview.NotificationSnapshot.Urgency != "critical" ||
		preview.NotificationSnapshot.Category != "compatibility.approval" ||
		preview.NotificationSnapshot.Title != "Example Ledger needs approval" ||
		preview.NotificationSnapshot.ActionCount != 2 ||
		!sameStrings(preview.NotificationSnapshot.Actions, []string{"open-compatibility-center", "review-request"}) ||
		!preview.NotificationSnapshot.RequiresUserReview ||
		!preview.NotificationSnapshot.UserVisible ||
		!preview.NotificationSnapshot.NotificationReady ||
		preview.NotificationSnapshot.ActionExecutionEnabled ||
		preview.NotificationSnapshot.RepairExecutionEnabled ||
		preview.NotificationSnapshot.SettingsPersistenceEnabled ||
		preview.NotificationSnapshot.NotificationsSent ||
		preview.NotificationSnapshot.HostRootModified ||
		preview.NotificationSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected notification snapshot: %#v", preview.NotificationSnapshot)
	}
	if preview.NavigationCount != 12 ||
		len(preview.Navigation) != 12 ||
		preview.PrimaryNavigationTarget != "compatibility-center-gates" ||
		preview.Navigation[0].ID != "overview" ||
		preview.Navigation[1].ID != "backend" ||
		preview.Navigation[2].ID != "activation" ||
		preview.Navigation[3].ID != "execution" ||
		preview.Navigation[4].ID != "launch" ||
		preview.Navigation[5].ID != "window" ||
		preview.Navigation[6].ID != "files" ||
		preview.Navigation[7].ID != "tray" ||
		preview.Navigation[8].ID != "notifications" ||
		preview.Navigation[9].ID != "actions" ||
		preview.Navigation[10].ID != "settings" ||
		preview.Navigation[11].ID != "diagnostics" {
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

func TestKDECenterPagePreviewSupportsApplicationsWithoutFileAssociations(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.xnix.apps.mines",
		Name:                "Mines",
		Icon:                "applications-games",
		Mode:                "automatic",
		SupportedExtensions: []string{},
	}
	preview, err := NewKDECenterPagePreview(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "approved", nil)
	if err != nil {
		t.Fatalf("NewKDECenterPagePreview returned error for no-file-association app: %v", err)
	}
	if preview.ApplicationID != "org.xnix.apps.mines" ||
		preview.ApplicationName != "Mines" ||
		preview.FileAssociationSnapshot.MIMETypeCount != 0 ||
		len(preview.FileAssociationSnapshot.MIMETypes) != 0 ||
		preview.FileAssociationSnapshot.FileAssociationReady ||
		preview.FileAssociationSnapshot.FileOpenPreviewAvailable ||
		preview.FileAssociationSnapshot.PortalRequiredForFileOpen ||
		preview.FileAssociationSnapshot.MIMEAppsWritten ||
		preview.FileAssociationSnapshot.HostRootModified ||
		preview.FileAssociationSnapshot.BackendDetailsExposed {
		t.Fatalf("unexpected no-file-association KDE page: %#v", preview.FileAssociationSnapshot)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", ".wine", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("no-file-association KDE page exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKDECenterPagePreviewConsumesActivationReceipt(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	root := t.TempDir()
	writeKDECenterActivationReceipt(t, root, "org.example.ledger")

	preview, err := NewKDECenterPagePreviewWithOptions(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "approved", []string{"file:///home/test/Documents/book.xls"}, KDECenterPageOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("NewKDECenterPagePreviewWithOptions returned error: %v", err)
	}

	activation := preview.ActivationStatusSnapshot
	if activation.ActivationState != "receipt-backed-runtime-gated" ||
		activation.ReceiptEvidenceState != "receipt-backed" ||
		activation.ReceiptRelativePath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" ||
		!activation.ReceiptBacked ||
		!activation.RollbackAvailable ||
		activation.CommitEnabled ||
		activation.LaunchEnabled ||
		activation.HostRootModified ||
		activation.BackendDetailsExposed {
		t.Fatalf("unexpected receipt-backed activation snapshot: %#v", activation)
	}
	if activation.StatusSignalCount != 6 || activation.BlockedReasonCount != 3 {
		t.Fatalf("unexpected receipt-backed activation counts: %#v", activation)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(root)) {
		t.Fatalf("KDE center page exposed activation root: %s", text)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page receipt-backed preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKDECenterPagePreviewConsumesExecutionSessionRecord(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	root := t.TempDir()
	requestID := "xnix-exec-org-example-ledger-1"
	writeExecutionSessionRecord(t, root, requestID, "org.example.ledger")

	preview, err := NewKDECenterPagePreviewWithOptions(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "approved", []string{"file:///home/test/Documents/book.xls"}, KDECenterPageOptions{
		ExecutionSessionRoot:      root,
		ExecutionSessionRequestID: requestID,
	})
	if err != nil {
		t.Fatalf("NewKDECenterPagePreviewWithOptions returned error: %v", err)
	}

	if preview.Source != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview+execution-session-record" {
		t.Fatalf("unexpected session-backed center page source: %s", preview.Source)
	}
	window := preview.WindowIdentitySnapshot
	if !window.ExecutionSessionRoot ||
		!window.ExecutionSessionBacked ||
		window.ExecutionSessionPath != "execution-ledger/sessions/"+requestID+".json" ||
		window.TaskManagerSessionState != "blocked" ||
		window.KWinSessionState != "blocked" ||
		window.TaskManagerEntryActive ||
		window.KWinRuleApplied ||
		window.HostRootModified ||
		window.BackendDetailsExposed {
		t.Fatalf("unexpected session-backed window snapshot: %#v", window)
	}
	tray := preview.TrayStatusSnapshot
	if tray.CompatibilityState != "waiting-for-runtime-gates" ||
		tray.CompatibilityLabel != "Runtime gates required" ||
		!tray.ExecutionSessionRoot ||
		!tray.ExecutionSessionBacked ||
		tray.ExecutionSessionPath != "execution-ledger/sessions/"+requestID+".json" ||
		tray.ExecutionSessionState != "blocked" ||
		tray.LiveBackendBridgeEnabled ||
		tray.BridgeConfigurationPersisted ||
		tray.HostRootModified ||
		tray.BackendDetailsExposed {
		t.Fatalf("unexpected session-backed tray snapshot: %#v", tray)
	}
	if preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("session-backed center page must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(root)) {
		t.Fatalf("KDE center page exposed session root: %s", text)
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page session-backed preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestKDECenterPagePreviewSurfacesKnownAppLauncherSessionGateCards(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	preview, err := NewKDECenterPagePreviewWithOptions(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	}, "approved", []string{"file:///home/test/Documents/book.xls"}, KDECenterPageOptions{
		KnownAppSmokeEvidence: []KnownAppSmokeEvidenceSummary{{
			AppID:                                 "7zr",
			DisplayName:                           "7-Zip Console",
			AppVersion:                            "26.02",
			EvidenceSource:                        "staged-launcher-dispatch-smoke",
			SmokeStatus:                           "passed",
			MarkerObserved:                        true,
			ChecksumVerified:                      true,
			LaunchAuthorizationReceiptState:       "recorded",
			LaunchAuthorizationReceiptID:          KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
			LaunchGateConsumed:                    true,
			LaunchGateReceiptAccepted:             true,
			LaunchGateGuestBoundaryAccepted:       true,
			ControlledDispatchReady:               true,
			ControlledExecutionSessionID:          sessionID,
			LauncherSessionGateConsumed:           true,
			LauncherSessionDigestVerified:         true,
			LauncherSessionRelativePath:           "execution-ledger/sessions/" + sessionID + ".json",
			LauncherSessionRuntimeOwnerConsumable: true,
			LauncherSessionKDEReadModelConsumable: true,
			PostReviewDispatchConsumed:            true,
			PostReviewDispatchState:               "created-after-session-gated-review",
			SessionGatedReviewReceiptID:           reviewReceiptID,
		}},
	})
	if err != nil {
		t.Fatalf("NewKDECenterPagePreviewWithOptions returned error: %v", err)
	}

	if preview.Source != "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview+known-app-session-gate-evidence" {
		t.Fatalf("unexpected known app session-gated source: %s", preview.Source)
	}
	if preview.KnownAppSessionGateEvidenceCount != 1 ||
		preview.KnownAppLauncherSessionGateConsumedCount != 1 ||
		preview.KnownAppPostReviewDispatchConsumedCount != 1 ||
		len(preview.KnownAppSessionGateCards) != 1 {
		t.Fatalf("unexpected known app session gate counts: %#v", preview)
	}
	card := preview.KnownAppSessionGateCards[0]
	if card.AppID != "7zr" ||
		card.DisplayName != "7-Zip Console" ||
		card.AppVersion != "26.02" ||
		card.CompatibilityState != "validated" ||
		card.CenterCardState != "validated-post-review-dispatch" ||
		card.ControlledExecutionSessionID != sessionID ||
		card.LaunchAuthorizationReceiptID != KnownAppLaunchAuthorizationReceiptID("7zr", "26.02") ||
		!card.LauncherSessionGateConsumed ||
		!card.LauncherSessionDigestVerified ||
		card.LauncherSessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		!card.RuntimeOwnerConsumableSession ||
		!card.KDEReadModelConsumableSession ||
		!card.PostReviewDispatchConsumed ||
		card.PostReviewDispatchState != "created-after-session-gated-review" ||
		card.SessionGatedReviewReceiptID != reviewReceiptID ||
		!card.LaunchGateConsumed ||
		!card.ControlledDispatchReady ||
		card.PrimaryActionID != "show-runtime-controlled-launch" ||
		card.PrimaryActionKind != "runtime-status" ||
		!card.PrimaryActionEnabled ||
		card.RuntimeStatusLaunchRequestType != KnownAppKDERuntimeStatusLaunchRequestType ||
		card.RuntimeStatusLaunchRuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		card.RuntimeStatusLaunchReadMethod != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		card.RuntimeStatusLaunchRequiredIDCount != 3 ||
		card.RuntimeStatusLaunchCollectedIDCount != 3 ||
		strings.Join(card.RuntimeStatusLaunchManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --receipt-id "+KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		!card.RuntimeStatusLaunchRequestReady ||
		!card.RuntimeStatusLaunchStateRootRequired ||
		!card.RuntimeStatusLaunchStateRootOwnedByRuntime ||
		card.ReviewRouteRequestType != KnownAppSessionGatedLaunchReviewRequestType ||
		card.ReviewRouteRuntimeMethod != "PreviewKnownAppSessionGatedLaunchReview" ||
		card.ReviewRouteReadMethod != "GetKnownAppSessionGatedLaunchReview" ||
		!card.ReadBeforeWriteRequired ||
		!card.RuntimeReceiptRequired ||
		!card.UserVisible ||
		!card.RuntimeOwned ||
		!card.GoRuntimeBacked ||
		card.KDEPolicyOwner ||
		card.DesktopLaunchEnabled ||
		card.BackendLaunchEnabled ||
		card.HostRootModified ||
		card.BackendDetailsExposed {
		t.Fatalf("unexpected known app session gate card: %#v", card)
	}
	if preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("known app session-gated center page must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE center page known app session gate preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func writeKDECenterActivationReceipt(t *testing.T, root string, applicationID string) {
	t.Helper()
	relativePath := filepath.Join("usr/share/xnix/compatibility/activation-receipts", applicationID+".json")
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	receipt := `{
  "schema_version": "xnix.runtime.desktop_activation_receipt.v1",
  "receipt_type": "desktop-activation-receipt",
  "application_id": "` + applicationID + `",
  "installed": [
    {
      "id": "desktop-entry",
      "relative_path": "usr/share/applications/xnix-org.example.ledger.desktop",
      "sha256": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
      "written": true,
      "host_root_modified": false,
      "backend_details_exposed": false
    }
  ],
  "rollback": {
    "command": "xnix-rollback-desktop-integration",
    "requires_matching_sha256": true,
    "host_root_modified": false
  },
  "safety": {
    "runtime_owned": true,
    "host_root_modified": false,
    "backend_details_exposed": false
  }
}
`
	if err := os.WriteFile(path, []byte(receipt), 0o600); err != nil {
		t.Fatalf("WriteFile receipt returned error: %v", err)
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
		preview.SectionCount != 12 ||
		preview.ReadOnlySectionCount != 12 ||
		preview.NavigationOnlySectionCount != 12 ||
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
	if preview.ApplicationReadinessEvidence.Source != "application-readiness-preview" ||
		preview.ApplicationReadinessEvidence.GraphType != "runtime-application-readiness-evidence-graph" ||
		preview.ApplicationReadinessEvidence.RuntimeMethod != "GetApplicationReadiness" ||
		preview.ApplicationReadinessEvidence.NodeCount != 7 ||
		preview.ApplicationReadinessEvidence.BlockedNodeCount != 2 ||
		preview.ApplicationReadinessEvidence.LaunchAllowed ||
		preview.ApplicationReadinessEvidence.LaunchEnabled ||
		preview.ApplicationReadinessEvidence.ExecutionRequestCreated ||
		preview.ApplicationReadinessEvidence.ExecutionStarted ||
		preview.ApplicationReadinessEvidence.BackendProcessStarted ||
		preview.ApplicationReadinessEvidence.RealPortalTransportEnabled ||
		preview.ApplicationReadinessEvidence.RequestObjectCreated ||
		preview.ApplicationReadinessEvidence.PermissionGranted ||
		preview.ApplicationReadinessEvidence.SnapshotCreated ||
		preview.ApplicationReadinessEvidence.RestoreExecuted ||
		preview.ApplicationReadinessEvidence.HostRootModified ||
		preview.ApplicationReadinessEvidence.NetworkRequired ||
		preview.ApplicationReadinessEvidence.BackendDetailsExposed {
		t.Fatalf("unexpected sections readiness evidence: %#v", preview.ApplicationReadinessEvidence)
	}
	wantMethods := map[string]string{
		"overview":      "GetCompatibilityCenterSummary",
		"backend":       "GetBackendSelectionPlan",
		"activation":    "GetDesktopActivationStatus",
		"execution":     "GetExecutionReadiness",
		"launch":        "GetLaunchIntent",
		"window":        "GetTaskManagerIdentityPlan",
		"files":         "GetFileAssociationPlan",
		"tray":          "GetTrayStatus",
		"notifications": "GetNotificationPlan",
		"actions":       "GetCompatibilityActionQueue",
		"settings":      "GetCompatibilitySettings",
		"diagnostics":   "GetDiagnostics",
	}
	wantModels := map[string]string{
		"overview":      "compatibility-center-summary",
		"backend":       "backend-selection-preview",
		"activation":    "desktop-activation-status-preview",
		"execution":     "execution-readiness-preview",
		"launch":        "launch-intent-preview",
		"window":        "window-identity-preview",
		"files":         "file-association-plan",
		"tray":          "tray-status-preview",
		"notifications": "notification-preview",
		"actions":       "compatibility-center-action-queue",
		"settings":      "settings-model",
		"diagnostics":   "diagnostic-history-preview",
	}
	for _, section := range preview.Sections {
		if section.RuntimeMethod != wantMethods[section.ID] ||
			section.ReadModel != wantModels[section.ID] ||
			len(section.ReadinessNodeIDs) == 0 ||
			section.ReadinessStatus == "" ||
			section.ReadinessSummary == "" ||
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
		if (section.ID == "execution" || section.ID == "launch" || section.ID == "actions") &&
			(section.ReadinessStatus != "blocked" || !section.ReadinessBlocked) {
			t.Fatalf("section should reflect blocked readiness evidence: %#v", section)
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
	if strings.Join(preview.ReadinessNodeIDs, ",") != "portal-review,snapshot-baseline" ||
		preview.ReadinessStatus != "required" ||
		preview.ReadinessBlocked ||
		preview.ReadinessSummary == "" ||
		preview.ApplicationReadinessEvidence.Source != "application-readiness-preview" ||
		preview.ApplicationReadinessEvidence.NodeCount != 2 ||
		preview.ApplicationReadinessEvidence.BlockedNodeCount != 0 ||
		preview.ApplicationReadinessEvidence.LaunchEnabled ||
		preview.ApplicationReadinessEvidence.ExecutionStarted ||
		preview.ApplicationReadinessEvidence.BackendProcessStarted ||
		preview.ApplicationReadinessEvidence.RequestObjectCreated ||
		preview.ApplicationReadinessEvidence.PermissionGranted ||
		preview.ApplicationReadinessEvidence.SnapshotCreated ||
		preview.ApplicationReadinessEvidence.HostRootModified ||
		preview.ApplicationReadinessEvidence.NetworkRequired ||
		preview.ApplicationReadinessEvidence.BackendDetailsExposed {
		t.Fatalf("unexpected settings section readiness evidence: %#v", preview)
	}
	if got := strings.Join(preview.AvailableSectionIDs, ","); got != "overview,backend,activation,execution,launch,window,files,tray,notifications,actions,settings,diagnostics" {
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
	launchPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "launch", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("launch NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if launchPreview.SectionID != "launch" ||
		launchPreview.SectionLabel != "Launch" ||
		launchPreview.SectionTarget != "compatibility-launch-intent" ||
		launchPreview.SectionState != "blocked" ||
		launchPreview.SectionRuntimeMethod != "GetLaunchIntent" ||
		launchPreview.SectionReadModel != "launch-intent-preview" ||
		strings.Join(launchPreview.ReadinessNodeIDs, ",") != "execution-readiness,runtime-write-gate" ||
		launchPreview.ReadinessStatus != "blocked" ||
		!launchPreview.ReadinessBlocked ||
		launchPreview.ApplicationReadinessEvidence.NodeCount != 2 ||
		launchPreview.ApplicationReadinessEvidence.BlockedNodeCount != 2 ||
		!launchPreview.ReadOnlyNavigation ||
		launchPreview.SectionActionsEnabled ||
		launchPreview.RequestObjectsCreated ||
		launchPreview.PermissionGrantCreated ||
		launchPreview.LaunchEnabled ||
		launchPreview.ExecutionStarted ||
		launchPreview.HostRootModified ||
		launchPreview.BackendDetailsExposed {
		t.Fatalf("unexpected launch section detail: %#v", launchPreview)
	}
	windowPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "window", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("window NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if windowPreview.SectionID != "window" ||
		windowPreview.SectionLabel != "Window" ||
		windowPreview.SectionTarget != "compatibility-window-identity" ||
		windowPreview.SectionState != "planned" ||
		windowPreview.SectionRuntimeMethod != "GetTaskManagerIdentityPlan" ||
		windowPreview.SectionReadModel != "window-identity-preview" ||
		!windowPreview.ReadOnlyNavigation ||
		windowPreview.SectionActionsEnabled ||
		windowPreview.RequestObjectsCreated ||
		windowPreview.PermissionGrantCreated ||
		windowPreview.LaunchEnabled ||
		windowPreview.ExecutionStarted ||
		windowPreview.HostRootModified ||
		windowPreview.BackendDetailsExposed {
		t.Fatalf("unexpected window section detail: %#v", windowPreview)
	}
	filesPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "files", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("files NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if filesPreview.SectionID != "files" ||
		filesPreview.SectionLabel != "Files" ||
		filesPreview.SectionTarget != "compatibility-file-association" ||
		filesPreview.SectionState != "planned" ||
		filesPreview.SectionRuntimeMethod != "GetFileAssociationPlan" ||
		filesPreview.SectionReadModel != "file-association-plan" ||
		!filesPreview.ReadOnlyNavigation ||
		filesPreview.SectionActionsEnabled ||
		filesPreview.RequestObjectsCreated ||
		filesPreview.PermissionGrantCreated ||
		filesPreview.LaunchEnabled ||
		filesPreview.ExecutionStarted ||
		filesPreview.HostRootModified ||
		filesPreview.BackendDetailsExposed {
		t.Fatalf("unexpected files section detail: %#v", filesPreview)
	}
	trayPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "tray", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("tray NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if trayPreview.SectionID != "tray" ||
		trayPreview.SectionLabel != "Tray" ||
		trayPreview.SectionTarget != "compatibility-tray-status" ||
		trayPreview.SectionState != "planned" ||
		trayPreview.SectionRuntimeMethod != "GetTrayStatus" ||
		trayPreview.SectionReadModel != "tray-status-preview" ||
		!trayPreview.ReadOnlyNavigation ||
		trayPreview.SectionActionsEnabled ||
		trayPreview.RequestObjectsCreated ||
		trayPreview.PermissionGrantCreated ||
		trayPreview.LaunchEnabled ||
		trayPreview.ExecutionStarted ||
		trayPreview.HostRootModified ||
		trayPreview.BackendDetailsExposed {
		t.Fatalf("unexpected tray section detail: %#v", trayPreview)
	}
	notificationsPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "notifications", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("notifications NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if notificationsPreview.SectionID != "notifications" ||
		notificationsPreview.SectionLabel != "Notifications" ||
		notificationsPreview.SectionTarget != "compatibility-notification-plan" ||
		notificationsPreview.SectionState != "planned" ||
		notificationsPreview.SectionRuntimeMethod != "GetNotificationPlan" ||
		notificationsPreview.SectionReadModel != "notification-preview" ||
		notificationsPreview.DiagnosticHistoryRoute != nil ||
		!notificationsPreview.ReadOnlyNavigation ||
		notificationsPreview.SectionActionsEnabled ||
		notificationsPreview.RequestObjectsCreated ||
		notificationsPreview.PermissionGrantCreated ||
		notificationsPreview.LaunchEnabled ||
		notificationsPreview.NotificationsSent ||
		notificationsPreview.ExecutionStarted ||
		notificationsPreview.HostRootModified ||
		notificationsPreview.BackendDetailsExposed {
		t.Fatalf("unexpected notifications section detail: %#v", notificationsPreview)
	}
	diagnosticsPreview, err := NewKDECenterPageSectionDetailPreview(recipe, Provenance{}, "diagnostics", "approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("diagnostics NewKDECenterPageSectionDetailPreview returned error: %v", err)
	}
	if diagnosticsPreview.SectionID != "diagnostics" ||
		diagnosticsPreview.SectionRuntimeMethod != "GetDiagnostics" ||
		diagnosticsPreview.SectionReadModel != "diagnostic-history-preview" ||
		diagnosticsPreview.DiagnosticHistoryRoute == nil ||
		diagnosticsPreview.DiagnosticHistoryRoute.RequestType != "diagnostic-history-route" ||
		diagnosticsPreview.DiagnosticHistoryRoute.Source != "kde-center-page-section-detail-preview" ||
		diagnosticsPreview.DiagnosticHistoryRoute.RuntimeMethod != "GetDiagnostics" ||
		diagnosticsPreview.DiagnosticHistoryRoute.ReadMethod != "GetDiagnosticHistoryPreview" ||
		diagnosticsPreview.DiagnosticHistoryRoute.ReadModel != "diagnostic-history-preview" ||
		diagnosticsPreview.DiagnosticHistoryRoute.CLICommand != "diagnostic-history-preview" ||
		diagnosticsPreview.DiagnosticHistoryRoute.ApplicationID != "org.example.ledger" ||
		!diagnosticsPreview.DiagnosticHistoryRoute.UserVisible ||
		!diagnosticsPreview.DiagnosticHistoryRoute.RuntimeOwned ||
		!diagnosticsPreview.DiagnosticHistoryRoute.GoRuntimeBacked ||
		diagnosticsPreview.DiagnosticHistoryRoute.KDEPolicyOwner ||
		!diagnosticsPreview.DiagnosticHistoryRoute.StateRootRequired ||
		diagnosticsPreview.DiagnosticHistoryRoute.StateRootPathExposed ||
		diagnosticsPreview.DiagnosticHistoryRoute.HistoryPreviewCreated ||
		diagnosticsPreview.DiagnosticHistoryRoute.AIProviderCallEnabled ||
		diagnosticsPreview.DiagnosticHistoryRoute.FileContentRead ||
		diagnosticsPreview.DiagnosticHistoryRoute.FilePathsExposed ||
		diagnosticsPreview.DiagnosticHistoryRoute.RequestObjectCreated ||
		diagnosticsPreview.DiagnosticHistoryRoute.PermissionGranted ||
		diagnosticsPreview.DiagnosticHistoryRoute.LaunchEnabled ||
		diagnosticsPreview.DiagnosticHistoryRoute.ExecutionStarted ||
		diagnosticsPreview.DiagnosticHistoryRoute.RepairExecutionEnabled ||
		diagnosticsPreview.DiagnosticHistoryRoute.HostRootModified ||
		diagnosticsPreview.DiagnosticHistoryRoute.NetworkRequired ||
		diagnosticsPreview.DiagnosticHistoryRoute.PrivilegedContainerRequired ||
		diagnosticsPreview.DiagnosticHistoryRoute.BackendDetailsExposed ||
		diagnosticsPreview.AIAnalysis == nil ||
		diagnosticsPreview.AIAnalysis.Source != "dolphin-ai-analysis-preview" ||
		diagnosticsPreview.AIAnalysis.Disclosure != "count-and-extension-only" ||
		diagnosticsPreview.AIAnalysisInput == nil ||
		diagnosticsPreview.AIAnalysisInput.RequestType != "dolphin-ai-analysis-preview" ||
		diagnosticsPreview.AIAnalysisInput.RuntimeMethod != "GetAIDiagnosticInput" ||
		diagnosticsPreview.AIAnalysisInput.AnalysisTask != "compatibility-file-review" ||
		diagnosticsPreview.AIAnalysisInput.AnalysisSurface != "Dolphin" ||
		diagnosticsPreview.AIAnalysisInput.SelectionMode != "explicit-application" ||
		diagnosticsPreview.AIAnalysisInput.FileCount != 1 ||
		diagnosticsPreview.AIAnalysisInput.SelectedExtension != ".xls" ||
		diagnosticsPreview.AIAnalysisInput.SelectedFileDisclosure != "count-and-extension-only" ||
		!diagnosticsPreview.AIAnalysisInput.UserReviewRequired ||
		!diagnosticsPreview.AIAnalysisInput.SafeForAIDiagnostics ||
		diagnosticsPreview.AIAnalysisInput.AIProviderCallEnabled ||
		diagnosticsPreview.AIAnalysisInput.NetworkRequired ||
		diagnosticsPreview.AIAnalysisInput.FileContentRead ||
		diagnosticsPreview.AIAnalysisInput.FilePathsExposed ||
		diagnosticsPreview.AIAnalysisInput.RequestObjectCreated ||
		diagnosticsPreview.AIAnalysisInput.PermissionGranted ||
		diagnosticsPreview.AIAnalysisInput.BackendLaunchEnabled ||
		diagnosticsPreview.AIAnalysisInput.HostRootModified ||
		diagnosticsPreview.AIAnalysisInput.BackendDetailsExposed ||
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
	diagnosticsEncoded, err := json.Marshal(diagnosticsPreview)
	if err != nil {
		t.Fatalf("Marshal diagnosticsPreview returned error: %v", err)
	}
	diagnosticsText := strings.ToLower(string(diagnosticsEncoded))
	for _, forbidden := range []string{"file://", "/home/test", "documents/book.xls"} {
		if strings.Contains(diagnosticsText, forbidden) {
			t.Fatalf("diagnostics section detail exposes forbidden term %q: %s", forbidden, diagnosticsText)
		}
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
