package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRunPlanPreviewExplainsRunStrategyWithoutOpeningLaunchGates(t *testing.T) {
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

	preview, err := plan.RunPlanPreview()
	if err != nil {
		t.Fatalf("RunPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.run_plan.v1" ||
		preview.RequestType != "run-plan-preview" ||
		preview.PlanType != "compatibility-run" ||
		preview.Source != "registry+go-runtime-run-plan" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetRunPlan" ||
		preview.ReadMethod != "GetRunPlanPreview" {
		t.Fatalf("unexpected run plan schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Application.RuntimeMode != "automatic" {
		t.Fatalf("unexpected run plan application: %#v", preview.Application)
	}
	if preview.Execution.Strategy != "automatic-managed" ||
		preview.Execution.SelectionSource != "recipe" ||
		preview.Execution.SelectedEngine.ID != "automatic" ||
		preview.Execution.SelectedEngine.Label != "Automatic" ||
		preview.Execution.SelectedEngine.Ready ||
		preview.Execution.SelectedEngine.LaunchEnabled ||
		preview.Execution.SelectedEngine.BackendDetailsExposed {
		t.Fatalf("unexpected run plan execution engine: %#v", preview.Execution)
	}
	if preview.Execution.SelectedProfile.ID != "local-compatibility" ||
		preview.Execution.SelectedProfile.Ready ||
		preview.Execution.SelectedProfile.LaunchEnabled ||
		preview.Execution.SelectedProfile.BackendDetailsExposed {
		t.Fatalf("unexpected run plan profile: %#v", preview.Execution.SelectedProfile)
	}
	if preview.Execution.BackendBinding.Ready ||
		preview.Execution.BackendBinding.LaunchEnabled ||
		preview.Execution.BackendBinding.BackendDetailsExposed ||
		preview.Execution.BackendDetailsExposed ||
		preview.Execution.LaunchEnabled ||
		preview.Execution.ExecutionRequestCreated ||
		preview.Execution.ExecutionStarted {
		t.Fatalf("run plan execution gates unexpectedly open: %#v", preview.Execution)
	}
	if !preview.Preflight.PortalPolicyRequired ||
		!preview.Preflight.SnapshotBeforeRiskyChange ||
		!preview.Preflight.DiagnosticsRequired ||
		!preview.Preflight.BackendBindingRequired ||
		!preview.Preflight.RuntimeWriteGateRequired ||
		!sameStrings(preview.Preflight.RequiredGateIDs, []string{"portal-policy-review", "snapshot-baseline", "diagnostics", "backend-binding", "runtime-launch-write-gate"}) {
		t.Fatalf("unexpected run plan preflight: %#v", preview.Preflight)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		preview.LaunchEnabled ||
		preview.ExecutionRequestCreated ||
		preview.ExecutionStarted ||
		preview.BackendBindingReady ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed ||
		preview.RawCommandExposed {
		t.Fatalf("unexpected run plan safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 6 ||
		preview.BlockedActions[0] != "create execution request before Runtime gates pass" {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	localPlan, err := NewPlan(Recipe{
		ID:                  "org.example.local",
		Name:                "Example Local",
		Icon:                "application-x-executable",
		Mode:                "wine",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan local returned error: %v", err)
	}
	localPreview, err := localPlan.RunPlanPreview()
	if err != nil {
		t.Fatalf("local RunPlanPreview returned error: %v", err)
	}
	if localPreview.Execution.Strategy != "local-compatible-managed" ||
		localPreview.Execution.SelectedEngine.ID != "local-compatible" {
		t.Fatalf("local recipe did not map to local-compatible engine: %#v", localPreview.Execution)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "application-x-executable",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.RunPlanPreview()
	if err != nil {
		t.Fatalf("isolated RunPlanPreview returned error: %v", err)
	}
	if isolatedPreview.Execution.Strategy != "isolated-compatible-managed" ||
		isolatedPreview.Execution.SelectedEngine.ID != "isolated-compatible" ||
		isolatedPreview.Execution.SelectedProfile.ID != "isolated-compatibility" {
		t.Fatalf("isolated recipe did not map to isolated-compatible engine: %#v", isolatedPreview.Execution)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("run plan preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
	if err := validateNoBackendTerms(preview, "run plan preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
