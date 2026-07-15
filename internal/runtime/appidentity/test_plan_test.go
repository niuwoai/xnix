package appidentity

import "testing"

func TestTestPlanPreviewListsPreflightStepsWithoutExecution(t *testing.T) {
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

	preview, err := plan.TestPlanPreview("preflight")
	if err != nil {
		t.Fatalf("TestPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.test_plan.v1" ||
		preview.RequestType != "test-plan-preview" ||
		preview.PlanType != "compatibility-test" ||
		preview.TestType != "preflight" ||
		preview.RuntimeMethod != "GetTestPlan" ||
		preview.ReadMethod != "GetTestPlanPreview" {
		t.Fatalf("unexpected test plan schema: %#v", preview)
	}
	if !sameStrings(preview.StepIDs, []string{"recipe-validation", "portal-preflight", "snapshot-preflight", "runtime-launch-binding"}) {
		t.Fatalf("unexpected test plan steps: %#v", preview.StepIDs)
	}
	if preview.Steps[0].Status != "pass" ||
		preview.Steps[1].Status != "pending" ||
		preview.Steps[1].PortalRequest == nil ||
		preview.Steps[1].PortalRequest.Method != "OpenFile" ||
		preview.Steps[1].PortalRequest.HandleToken != "xnix_org_example_ledger_file_open" {
		t.Fatalf("unexpected portal step: %#v", preview.Steps[1])
	}
	if preview.Steps[3].RunPlan == nil ||
		preview.Steps[3].RunPlan.Strategy != "automatic-managed" ||
		preview.Steps[3].RunPlan.BackendReady {
		t.Fatalf("unexpected run plan step: %#v", preview.Steps[3])
	}
	if preview.Blocked ||
		len(preview.BlockingReasons) != 0 ||
		preview.Artifacts.NotificationEvent != "mode-changed" ||
		preview.Artifacts.RepairPlanIssue != "engine-binding-pending" {
		t.Fatalf("unexpected test plan artifacts: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.ExecutionRequestCreated ||
		preview.TestExecuted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected test plan safety flags: %#v", preview)
	}

	if _, err := plan.TestPlanPreview("unknown"); err == nil {
		t.Fatalf("TestPlanPreview must reject unknown test type")
	}

	if err := validateNoBackendTerms(preview, "test plan preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
