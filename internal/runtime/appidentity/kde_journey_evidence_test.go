package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEJourneyEvidencePreviewStitchesSevenEntryPoints(t *testing.T) {
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

	preview, err := plan.KDEJourneyEvidencePreviewWithOptions("approved", []string{"file:///home/test/Documents/book.xls"}, KDEJourneyEvidenceOptions{RuntimeRoot: "../../.."})
	if err != nil {
		t.Fatalf("KDEJourneyEvidencePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_journey_evidence.v1" ||
		preview.RequestType != "kde-journey-evidence-preview" ||
		preview.JourneyType != "kde-seven-entrypoint-runtime-evidence" ||
		preview.Source != "application-readiness-preview+kde-action-dependency-graph-preview+task-manager-identity-preview+kwin-window-rule-preview+tray-status-preview+notification-preview+settings-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEJourneyEvidence" ||
		preview.ReadMethod != "GetKDEJourneyEvidencePreview" {
		t.Fatalf("unexpected journey schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected journey identity: %#v", preview)
	}
	if got, want := preview.EntryPointIDs, []string{"launcher", "task-manager", "file-manager", "system-tray", "notification-center", "ai-compatibility-center", "unified-settings"}; !sameStrings(got, want) {
		t.Fatalf("unexpected entry point ids: %#v", got)
	}
	if preview.EntryPointCount != 7 ||
		len(preview.EntryPoints) != 7 ||
		preview.CrossLinkedReadModelCount != 10 ||
		!containsString(preview.CrossLinkedReadModels, "application-readiness-preview") ||
		!containsString(preview.CrossLinkedReadModels, "kde-action-dependency-graph-preview") ||
		!containsString(preview.CrossLinkedReadModels, "kde-center-page-preview") ||
		preview.SharedReadinessStatus != "not-ready" ||
		preview.Ready ||
		preview.ReadinessNodeCount != 7 ||
		preview.BlockedReadinessNodeCount == 0 ||
		preview.MissingEvidenceCount != 35 ||
		preview.BlockedActionCount != 7 {
		t.Fatalf("unexpected journey counts: %#v", preview)
	}
	if !preview.Agreement.AppIDConsistent ||
		!preview.Agreement.DisplayNameConsistent ||
		!preview.Agreement.DesktopFileConsistent ||
		!preview.Agreement.ReadinessStateConsistent ||
		!preview.Agreement.DisabledActionStateShared ||
		!preview.Agreement.UnsafeSideEffectsDisabled ||
		!containsString(preview.Agreement.CheckedReadModels, "task-manager-identity-preview") ||
		!containsString(preview.Agreement.CheckedReadModels, "kwin-window-rule-preview") ||
		!containsString(preview.Agreement.CheckedReadModels, "desktop-notification-preview") {
		t.Fatalf("unexpected journey agreement: %#v", preview.Agreement)
	}

	fileManager := findKDEJourneyEntryPoint(preview.EntryPoints, "file-manager")
	if fileManager.ReadModel != "file-open-preview" ||
		fileManager.RuntimeMethod != "GetFileOpenPlan" ||
		fileManager.SharedReadinessStatus != preview.SharedReadinessStatus ||
		fileManager.Ready ||
		!containsString(fileManager.EvidenceIDs, "portal-file-access-receipt") ||
		fileManager.NextSafeReadOnlyCheck != "file-open-preview" ||
		fileManager.LaunchAllowed ||
		fileManager.ExecutionStarted ||
		fileManager.RequestObjectsCreated ||
		fileManager.PermissionGrantCreated ||
		fileManager.HostRootModified ||
		fileManager.BackendDetailsExposed {
		t.Fatalf("unexpected file manager journey entry: %#v", fileManager)
	}

	center := findKDEJourneyEntryPoint(preview.EntryPoints, "ai-compatibility-center")
	if center.ReadModel != "kde-center-page-preview" ||
		!containsString(center.EvidenceIDs, "action-card-deck") ||
		center.KDEPolicyOwner ||
		!center.RuntimeOwned ||
		!center.GoRuntimeBacked {
		t.Fatalf("unexpected center journey entry: %#v", center)
	}
	if preview.RuntimeWriteMethodsEnabled ||
		preview.JourneyEvidencePersisted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.SettingsPersisted ||
		preview.LaunchAllowed ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.KWinRuleApplied ||
		preview.TrayBridgeActivated ||
		preview.NotificationSent ||
		preview.HostRootModified ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.FileContentRead ||
		preview.BackendDetailsExposed {
		t.Fatalf("journey preview enabled unsafe behavior: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(strings.ToLower(string(encoded)), "program files") ||
		strings.Contains(strings.ToLower(string(encoded)), ".exe") ||
		strings.Contains(strings.ToLower(string(encoded)), ".wine") {
		t.Fatalf("journey preview exposed forbidden implementation detail: %s", string(encoded))
	}
}

func TestKDEJourneyEvidencePreviewRejectsMalformedInput(t *testing.T) {
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
	if _, err := plan.KDEJourneyEvidencePreviewWithOptions("approved\nbad", nil, KDEJourneyEvidenceOptions{RuntimeRoot: "../../.."}); err == nil {
		t.Fatalf("KDEJourneyEvidencePreview accepted a multiline decision")
	}
	if _, err := plan.KDEJourneyEvidencePreviewWithOptions("approved", []string{"https://example.invalid/book.xls"}, KDEJourneyEvidenceOptions{RuntimeRoot: "../../.."}); err == nil {
		t.Fatalf("KDEJourneyEvidencePreview accepted a non-file URI")
	}
}

func findKDEJourneyEntryPoint(entries []KDEJourneyEntryPoint, id string) KDEJourneyEntryPoint {
	for _, entry := range entries {
		if entry.ID == id {
			return entry
		}
	}
	return KDEJourneyEntryPoint{}
}
