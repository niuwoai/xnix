package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDiagnosticsPreviewAggregatesGoRuntimeStatus(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	}
	preview, err := NewDiagnosticsPreview(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewDiagnosticsPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.diagnostics.v1" ||
		preview.RequestType != "diagnostics-preview" ||
		preview.DiagnosticsType != "runtime-diagnostics" ||
		preview.Source != "registry+go-runtime-previews" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetDiagnostics" ||
		preview.ReadMethod != "GetDiagnosticsPreview" {
		t.Fatalf("unexpected diagnostics schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Application.RegistryName != "test-registry" ||
		!preview.Application.DigestVerified ||
		preview.Application.SignatureStatus != "development-only" ||
		preview.Status != "known" ||
		preview.RuntimeMode != "Automatic" {
		t.Fatalf("unexpected diagnostics application: %#v", preview.Application)
	}
	if !sameStrings(preview.Application.SupportedExtensions, []string{".abc", ".xls"}) {
		t.Fatalf("unexpected supported extensions: %#v", preview.Application.SupportedExtensions)
	}
	expectedCheckIDs := []string{"recipe-validation", "desktop-identity", "execution-readiness", "launch-write-gate", "action-queue", "ai-diagnostics-safety"}
	expectedStatuses := []string{"pass", "pass", "blocked", "blocked", "pending", "pass"}
	if len(preview.Checks) != len(expectedCheckIDs) || len(preview.CheckIDs) != len(expectedCheckIDs) {
		t.Fatalf("unexpected diagnostics checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	for index, id := range expectedCheckIDs {
		if preview.Checks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.Checks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected diagnostics check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 6 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Pending != 1 ||
		preview.Counts.Blocked != 2 {
		t.Fatalf("unexpected diagnostics counts: %#v", preview.Counts)
	}
	if preview.ExecutionReadiness.RequestType != "execution-readiness-preview" ||
		preview.ExecutionReadiness.OverallStatus != "not-ready" ||
		preview.ExecutionReadiness.ExecutionState != "blocked" ||
		preview.ExecutionReadiness.GateCount != 5 ||
		preview.ExecutionReadiness.RequiredGateCount != 2 ||
		preview.ExecutionReadiness.PendingGateCount != 1 ||
		preview.ExecutionReadiness.BlockedGateCount != 1 ||
		preview.ExecutionReadiness.LaunchEnabled ||
		preview.ExecutionReadiness.BackendBindingReady ||
		!preview.ExecutionReadiness.SafeForAIDiagnostics ||
		preview.ExecutionReadiness.BackendDetailsExposed {
		t.Fatalf("unexpected execution diagnostics: %#v", preview.ExecutionReadiness)
	}
	if preview.LaunchIntent.RequestType != "launch-intent-preview" ||
		preview.LaunchIntent.WriteGateDecision != "blocked-until-production-backend" ||
		preview.LaunchIntent.PortalRequired ||
		!preview.LaunchIntent.SnapshotRequired ||
		preview.LaunchIntent.LaunchEnabled ||
		preview.LaunchIntent.ExecutionStarted ||
		preview.LaunchIntent.BackendDetailsExposed {
		t.Fatalf("unexpected launch diagnostics: %#v", preview.LaunchIntent)
	}
	if preview.ActionQueue.RequestType != "kde-action-queue-preview" ||
		preview.ActionQueue.ActionCount != 7 ||
		preview.ActionQueue.PendingActionCount != 7 ||
		preview.ActionQueue.UserReviewRequiredCount != 4 ||
		preview.ActionQueue.RuntimeGateActionCount != 7 ||
		!preview.ActionQueue.ActionQueueCreated ||
		preview.ActionQueue.ActionQueuePersisted ||
		preview.ActionQueue.ExecutionStarted ||
		preview.ActionQueue.BackendDetailsExposed {
		t.Fatalf("unexpected action diagnostics: %#v", preview.ActionQueue)
	}
	if preview.CompatibilityCenterSummary.RequestType != "compatibility-center-preview" ||
		preview.CompatibilityCenterSummary.CompatibilityState != "registered" ||
		preview.CompatibilityCenterSummary.DiagnosticsState != "not-run" ||
		preview.CompatibilityCenterSummary.KnownIssueCount != 0 ||
		preview.CompatibilityCenterSummary.ActionExecutionEnabled ||
		preview.CompatibilityCenterSummary.RepairExecutionEnabled ||
		preview.CompatibilityCenterSummary.BackendLaunchEnabled ||
		preview.CompatibilityCenterSummary.SettingsPersistenceEnabled ||
		preview.CompatibilityCenterSummary.BackendDetailsExposed {
		t.Fatalf("unexpected center diagnostics: %#v", preview.CompatibilityCenterSummary)
	}
	if preview.KDECenterSections.RequestType != "kde-center-page-sections-preview" ||
		preview.KDECenterSections.SectionCount != 12 ||
		preview.KDECenterSections.ReadOnlySectionCount != 12 ||
		!containsString(preview.KDECenterSections.RuntimeMethods, "GetAIDiagnosticInput") ||
		preview.KDECenterSections.SectionActionsEnabled ||
		preview.KDECenterSections.RequestObjectsCreated ||
		preview.KDECenterSections.ExecutionStarted ||
		preview.KDECenterSections.BackendDetailsExposed {
		t.Fatalf("unexpected KDE section diagnostics: %#v", preview.KDECenterSections)
	}
	if preview.AI.RuntimeMethod != "GetAIDiagnosticInput" ||
		preview.AI.AnalysisTask != "compatibility-status-review" ||
		preview.AI.Disclosure != "runtime-metadata-only" ||
		!preview.AI.SafeForAIDiagnostics ||
		preview.AI.AIProviderCallEnabled ||
		preview.AI.NetworkRequired ||
		preview.AI.FileContentRead ||
		preview.AI.FilePathsExposed ||
		preview.AI.RequestObjectCreated ||
		preview.AI.PermissionGranted {
		t.Fatalf("unexpected AI diagnostics: %#v", preview.AI)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics ||
		preview.AIProviderCallEnabled ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.RepairExecutionEnabled ||
		preview.SettingsPersisted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected diagnostics safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "start compatibility execution from diagnostics preview" ||
		preview.DesktopSafeSummary != "Diagnostics preview is Go-owned and summarizes Runtime-safe application status while execution, repair, AI calls, and host mutation remain disabled." {
		t.Fatalf("unexpected diagnostics metadata: actions=%#v summary=%q", preview.BlockedActions, preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Diagnostics preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("diagnostics preview exposes forbidden term %q: %s", forbidden, serialized)
		}
	}
}
