package appidentity

import (
	"errors"
	"strings"
)

var compatibilityTestTypes = []string{"preflight", "smoke", "repair-readiness"}

type TestPlanPreview struct {
	SchemaVersion           string            `json:"schema_version"`
	RequestType             string            `json:"request_type"`
	PlanType                string            `json:"plan_type"`
	TestType                string            `json:"test_type"`
	Source                  string            `json:"source"`
	Desktop                 string            `json:"desktop"`
	RuntimeMethod           string            `json:"runtime_method"`
	ReadMethod              string            `json:"read_method"`
	ApplicationID           string            `json:"application_id"`
	Name                    string            `json:"name"`
	RuntimeOwned            bool              `json:"runtime_owned"`
	GoRuntimeBacked         bool              `json:"go_runtime_backed"`
	KDEPolicyOwner          bool              `json:"kde_policy_owner"`
	Steps                   []TestPlanStep    `json:"steps"`
	StepIDs                 []string          `json:"step_ids"`
	Blocked                 bool              `json:"blocked"`
	BlockingReasons         []string          `json:"blocking_reasons"`
	Artifacts               TestPlanArtifacts `json:"artifacts"`
	ExecutionRequestCreated bool              `json:"execution_request_created"`
	TestExecuted            bool              `json:"test_executed"`
	HostRootModified        bool              `json:"host_root_modified"`
	BackendDetailsExposed   bool              `json:"backend_details_exposed"`
	DesktopSafeSummary      string            `json:"desktop_safe_summary"`
}

type TestPlanStep struct {
	ID            string                `json:"id"`
	Status        string                `json:"status"`
	Title         string                `json:"title"`
	Summary       string                `json:"summary"`
	RequiredFor   []string              `json:"required_for,omitempty"`
	PortalRequest *TestPlanPortalStep   `json:"portal_request,omitempty"`
	Snapshot      *TestPlanSnapshotStep `json:"snapshot,omitempty"`
	RunPlan       *TestPlanRunStep      `json:"run_plan,omitempty"`
}

type TestPlanPortalStep struct {
	Destination string `json:"destination"`
	Interface   string `json:"interface"`
	Method      string `json:"method"`
	HandleToken string `json:"handle_token"`
}

type TestPlanSnapshotStep struct {
	Reason           string `json:"reason"`
	EnabledByDefault bool   `json:"enabled_by_default"`
	RestoreAvailable bool   `json:"restore_available"`
}

type TestPlanRunStep struct {
	Strategy              string `json:"strategy"`
	BackendReady          bool   `json:"backend_ready"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type TestPlanArtifacts struct {
	CompatibilityCenterCard bool   `json:"compatibility_center_card"`
	NotificationEvent       string `json:"notification_event"`
	RepairPlanIssue         string `json:"repair_plan_issue"`
}

func (plan Plan) TestPlanPreview(testType string) (TestPlanPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return TestPlanPreview{}, err
	}
	if !containsString(compatibilityTestTypes, testType) {
		return TestPlanPreview{}, errors.New("test type must be one of: " + strings.Join(compatibilityTestTypes, ", "))
	}

	steps := plan.testPlanSteps()
	blocked := false
	blockingReasons := make([]string, 0)
	stepIDs := make([]string, 0, len(steps))
	for _, step := range steps {
		stepIDs = append(stepIDs, step.ID)
		if step.Status == "blocked" {
			blocked = true
			blockingReasons = append(blockingReasons, step.Summary)
		}
	}

	notificationEvent := "mode-changed"
	repairIssue := "engine-binding-pending"
	summary := "Compatibility test plan is ready for Runtime-controlled preflight work."
	if blocked {
		notificationEvent = "approval-required"
		repairIssue = "portal-approval-required"
		summary = "Compatibility test plan is blocked by user-mediated desktop access."
	}

	preview := TestPlanPreview{
		SchemaVersion:   "xnix.runtime.test_plan.v1",
		RequestType:     "test-plan-preview",
		PlanType:        "compatibility-test",
		TestType:        testType,
		Source:          "registry+go-runtime-test-plan",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetTestPlan",
		ReadMethod:      "GetTestPlanPreview",
		ApplicationID:   plan.ApplicationID,
		Name:            plan.DisplayName,
		RuntimeOwned:    true,
		GoRuntimeBacked: true,
		KDEPolicyOwner:  false,
		Steps:           steps,
		StepIDs:         stepIDs,
		Blocked:         blocked,
		BlockingReasons: blockingReasons,
		Artifacts: TestPlanArtifacts{
			CompatibilityCenterCard: true,
			NotificationEvent:       notificationEvent,
			RepairPlanIssue:         repairIssue,
		},
		ExecutionRequestCreated: false,
		TestExecuted:            false,
		HostRootModified:        false,
		BackendDetailsExposed:   false,
		DesktopSafeSummary:      summary,
	}
	if err := validateNoBackendTerms(preview, "test plan preview"); err != nil {
		return TestPlanPreview{}, err
	}
	return preview, nil
}

func (plan Plan) testPlanSteps() []TestPlanStep {
	return []TestPlanStep{
		{
			ID:          "recipe-validation",
			Status:      "pass",
			Title:       "Validate application recipe",
			Summary:     "Recipe metadata can produce desktop integration artifacts.",
			RequiredFor: []string{"launcher", "file-association", "diagnostics"},
		},
		{
			ID:      "portal-preflight",
			Status:  "pending",
			Title:   "Prepare user-mediated file access",
			Summary: "File access must use an XDG Desktop Portal request before selected documents are opened.",
			PortalRequest: &TestPlanPortalStep{
				Destination: "org.freedesktop.portal.Desktop",
				Interface:   "org.freedesktop.portal.FileChooser",
				Method:      "OpenFile",
				HandleToken: testPlanHandleToken(plan.ApplicationID),
			},
		},
		{
			ID:      "snapshot-preflight",
			Status:  "pending",
			Title:   "Prepare restore point",
			Summary: "Create an application restore point before risky compatibility changes.",
			Snapshot: &TestPlanSnapshotStep{
				Reason:           "before-engine-change",
				EnabledByDefault: true,
				RestoreAvailable: true,
			},
		},
		{
			ID:      "runtime-launch-binding",
			Status:  "pending",
			Title:   "Bind managed launch backend",
			Summary: "Runtime launch backend binding is required before automated compatibility execution.",
			RunPlan: &TestPlanRunStep{
				Strategy:              plan.runPlanStrategy(),
				BackendReady:          false,
				BackendDetailsExposed: false,
			},
		},
	}
}

func testPlanHandleToken(applicationID string) string {
	sanitized := make([]rune, 0, len(applicationID))
	for _, r := range strings.ToLower(applicationID) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			sanitized = append(sanitized, r)
		} else {
			sanitized = append(sanitized, '_')
		}
	}
	return "xnix_" + string(sanitized) + "_file_open"
}
