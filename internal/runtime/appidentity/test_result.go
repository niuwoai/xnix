package appidentity

type TestResultPreview struct {
	SchemaVersion         string              `json:"schema_version"`
	RequestType           string              `json:"request_type"`
	ResultType            string              `json:"result_type"`
	PlanType              string              `json:"plan_type"`
	TestType              string              `json:"test_type"`
	Source                string              `json:"source"`
	Desktop               string              `json:"desktop"`
	RuntimeMethod         string              `json:"runtime_method"`
	ReadMethod            string              `json:"read_method"`
	ApplicationID         string              `json:"application_id"`
	Name                  string              `json:"name"`
	RuntimeOwned          bool                `json:"runtime_owned"`
	GoRuntimeBacked       bool                `json:"go_runtime_backed"`
	KDEPolicyOwner        bool                `json:"kde_policy_owner"`
	ResultSource          string              `json:"result_source"`
	ExecutionState        string              `json:"execution_state"`
	OverallStatus         string              `json:"overall_status"`
	Counts                TestResultCounts    `json:"counts"`
	StepResults           []TestResultStep    `json:"step_results"`
	Artifacts             TestResultArtifacts `json:"artifacts"`
	SafeForAIDiagnostics  bool                `json:"safe_for_ai_diagnostics"`
	TestExecuted          bool                `json:"test_executed"`
	HostRootModified      bool                `json:"host_root_modified"`
	BackendDetailsExposed bool                `json:"backend_details_exposed"`
	DesktopSafeSummary    string              `json:"desktop_safe_summary"`
}

type TestResultCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

type TestResultStep struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Status     string   `json:"status"`
	Result     string   `json:"result"`
	Evidence   []string `json:"evidence"`
	NextAction string   `json:"next_action"`
}

type TestResultArtifacts struct {
	CompatibilityCenterCard bool   `json:"compatibility_center_card"`
	DiagnosticsRecord       bool   `json:"diagnostics_record"`
	NotificationEvent       string `json:"notification_event"`
	RepairPlanIssue         string `json:"repair_plan_issue"`
}

func (plan Plan) TestResultPreview(testType string) (TestResultPreview, error) {
	testPlan, err := plan.TestPlanPreview(testType)
	if err != nil {
		return TestResultPreview{}, err
	}

	stepResults := make([]TestResultStep, 0, len(testPlan.Steps))
	counts := TestResultCounts{}
	for _, step := range testPlan.Steps {
		counts.Total++
		switch step.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
		stepResults = append(stepResults, TestResultStep{
			ID:         step.ID,
			Title:      step.Title,
			Status:     step.Status,
			Result:     testResultForStatus(step.Status),
			Evidence:   testResultEvidence(step),
			NextAction: testResultNextAction(step),
		})
	}

	executionState := "complete"
	overallStatus := "pass"
	if counts.Pending > 0 {
		executionState = "waiting-for-runtime"
		overallStatus = "pending"
	}
	if counts.Blocked > 0 {
		executionState = "blocked"
		overallStatus = "blocked"
	}

	notificationEvent := "mode-changed"
	repairIssue := "engine-binding-pending"
	if overallStatus == "blocked" {
		notificationEvent = "approval-required"
	}
	if overallStatus == "pass" {
		repairIssue = ""
	}

	summary := "Compatibility test execution completed successfully."
	switch overallStatus {
	case "blocked":
		summary = "Compatibility test execution is blocked and needs user action."
	case "pending":
		summary = "Compatibility test execution is waiting for Runtime-controlled preflight work."
	}

	preview := TestResultPreview{
		SchemaVersion:   "xnix.runtime.test_result.v1",
		RequestType:     "test-result-preview",
		ResultType:      "compatibility-test-result",
		PlanType:        testPlan.PlanType,
		TestType:        testPlan.TestType,
		Source:          "test-plan-preview+go-runtime-test-result",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetTestResult",
		ReadMethod:      "GetTestResultPreview",
		ApplicationID:   testPlan.ApplicationID,
		Name:            testPlan.Name,
		RuntimeOwned:    true,
		GoRuntimeBacked: true,
		KDEPolicyOwner:  false,
		ResultSource:    "runtime-model",
		ExecutionState:  executionState,
		OverallStatus:   overallStatus,
		Counts:          counts,
		StepResults:     stepResults,
		Artifacts: TestResultArtifacts{
			CompatibilityCenterCard: true,
			DiagnosticsRecord:       true,
			NotificationEvent:       notificationEvent,
			RepairPlanIssue:         repairIssue,
		},
		SafeForAIDiagnostics:  true,
		TestExecuted:          false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		DesktopSafeSummary:    summary,
	}
	if err := validateNoBackendTerms(preview, "test result preview"); err != nil {
		return TestResultPreview{}, err
	}
	return preview, nil
}

func testResultForStatus(status string) string {
	switch status {
	case "pass":
		return "passed"
	case "blocked":
		return "blocked"
	default:
		return "not-run"
	}
}

func testResultEvidence(step TestPlanStep) []string {
	switch step.ID {
	case "recipe-validation":
		return []string{"Recipe metadata is loaded and valid."}
	case "portal-preflight":
		return []string{"XDG Desktop Portal request model is available."}
	case "snapshot-preflight":
		return []string{"Runtime snapshot plan is available and restore-capable."}
	case "runtime-launch-binding":
		return []string{"Runtime launch backend binding is still pending."}
	default:
		return []string{step.Summary}
	}
}

func testResultNextAction(step TestPlanStep) string {
	switch step.Status {
	case "pass":
		return "No action required."
	case "blocked":
		return "Request user approval before continuing."
	}
	switch step.ID {
	case "portal-preflight":
		return "Wait for a user-mediated Portal grant during execution."
	case "snapshot-preflight":
		return "Create a Runtime restore point before risky compatibility changes."
	case "runtime-launch-binding":
		return "Bind a managed compatibility launch backend before automated smoke execution."
	default:
		return "Run the pending compatibility test step."
	}
}
