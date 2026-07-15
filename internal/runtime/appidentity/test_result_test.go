package appidentity

import "testing"

func TestTestResultPreviewDerivesPendingResultFromPlan(t *testing.T) {
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

	preview, err := plan.TestResultPreview("preflight")
	if err != nil {
		t.Fatalf("TestResultPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.test_result.v1" ||
		preview.RequestType != "test-result-preview" ||
		preview.ResultType != "compatibility-test-result" ||
		preview.PlanType != "compatibility-test" ||
		preview.RuntimeMethod != "GetTestResult" ||
		preview.ReadMethod != "GetTestResultPreview" ||
		preview.ResultSource != "runtime-model" {
		t.Fatalf("unexpected test result schema: %#v", preview)
	}
	if preview.ExecutionState != "waiting-for-runtime" ||
		preview.OverallStatus != "pending" ||
		preview.Counts.Total != 4 ||
		preview.Counts.Passed != 1 ||
		preview.Counts.Pending != 3 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected test result counts: %#v", preview)
	}
	if len(preview.StepResults) != 4 ||
		preview.StepResults[0].Result != "passed" ||
		preview.StepResults[0].NextAction != "No action required." ||
		preview.StepResults[1].Result != "not-run" {
		t.Fatalf("unexpected step results: %#v", preview.StepResults)
	}
	if !preview.SafeForAIDiagnostics ||
		preview.TestExecuted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.Artifacts.RepairPlanIssue != "engine-binding-pending" {
		t.Fatalf("unexpected test result safety flags: %#v", preview)
	}

	if err := validateNoBackendTerms(preview, "test result preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
