package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionReadinessPreviewKeepsLaunchGated(t *testing.T) {
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

	preview, err := plan.ExecutionReadinessPreview()
	if err != nil {
		t.Fatalf("ExecutionReadinessPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.launch_readiness.v1" ||
		preview.RequestType != "execution-readiness-preview" ||
		preview.ReadinessType != "compatibility-execution-readiness" ||
		preview.Source != "compatibility-center" ||
		preview.RuntimeMethod != "GetExecutionReadiness" {
		t.Fatalf("unexpected execution readiness schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution readiness identity: %#v", preview)
	}
	if preview.CompatibilityProfile.ID != "local-compatibility" ||
		preview.CompatibilityProfile.Label != "Local compatibility" ||
		preview.CompatibilityProfile.Kind != "local" ||
		preview.CompatibilityProfile.Ready ||
		preview.CompatibilityProfile.LaunchEnabled ||
		preview.CompatibilityProfile.BackendDetailsExposed {
		t.Fatalf("unexpected readiness profile: %#v", preview.CompatibilityProfile)
	}
	if preview.ExecutionState != "blocked" ||
		preview.OverallStatus != "not-ready" ||
		preview.GateCount != 5 ||
		preview.RequiredGateCount != 2 ||
		preview.PendingGateCount != 1 ||
		preview.BlockedGateCount != 1 {
		t.Fatalf("unexpected readiness state: %#v", preview)
	}
	if got, want := executionReadinessGateIDsForTest(preview.Gates), []string{"recipe-validation", "portal-policy-review", "snapshot-baseline", "backend-binding", "runtime-launch-write-gate"}; !sameStrings(got, want) {
		t.Fatalf("gate ids = %#v, want %#v", got, want)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.BackendBindingReady || !preview.PortalPolicyRequired ||
		!preview.SnapshotRequired || !preview.UserActionRequired ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution readiness safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "launch compatibility profile") ||
		!containsString(preview.BlockedActions, "expose raw backend command to desktop shell") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "applications-games",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.ExecutionReadinessPreview()
	if err != nil {
		t.Fatalf("isolated ExecutionReadinessPreview returned error: %v", err)
	}
	if isolatedPreview.CompatibilityProfile.ID != "isolated-compatibility" ||
		isolatedPreview.CompatibilityProfile.Kind != "isolated" {
		t.Fatalf("isolated recipe did not recommend isolated compatibility: %#v", isolatedPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution readiness preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func executionReadinessGateIDsForTest(gates []ExecutionReadinessGate) []string {
	ids := make([]string, 0, len(gates))
	for _, gate := range gates {
		ids = append(ids, gate.ID)
	}
	return ids
}
