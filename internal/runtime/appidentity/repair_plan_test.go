package appidentity

import "testing"

func newRepairTestPlan(t *testing.T) Plan {
	t.Helper()
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
	return plan
}

func TestRepairPlanPreviewMapsIssueRulesWithoutExecution(t *testing.T) {
	plan := newRepairTestPlan(t)

	preview, err := plan.RepairPlanPreview("engine-binding-pending")
	if err != nil {
		t.Fatalf("RepairPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.repair_plan.v1" ||
		preview.RequestType != "repair-plan-preview" ||
		preview.PlanType != "compatibility-repair" ||
		preview.RuntimeMethod != "GetRepairPlan" ||
		preview.ReadMethod != "GetRepairPlanPreview" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.Issue != "engine-binding-pending" {
		t.Fatalf("unexpected repair plan schema: %#v", preview)
	}
	if preview.Severity != "warning" ||
		preview.AutomaticAllowed ||
		!preview.UserApprovalRequired ||
		!preview.SnapshotRequired ||
		preview.SnapshotPlan == nil ||
		preview.SnapshotPlan.Reason != "before-repair" ||
		preview.SnapshotPlan.Retention.KeepLatest != 5 ||
		preview.NotificationEvent != "approval-required" {
		t.Fatalf("unexpected engine-binding-pending rule: %#v", preview)
	}
	if !sameStrings(preview.ActionIDs, []string{"open-diagnostics", "prepare-engine-binding"}) {
		t.Fatalf("unexpected repair actions: %#v", preview.ActionIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.RepairExecutionRequested ||
		preview.RepairExecuted ||
		preview.BackendLaunchEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected repair plan safety flags: %#v", preview)
	}

	// runtime-repair-applied requires a snapshot and is auto-allowed.
	applied, err := plan.RepairPlanPreview("runtime-repair-applied")
	if err != nil {
		t.Fatalf("RepairPlanPreview(runtime-repair-applied) returned error: %v", err)
	}
	if !applied.AutomaticAllowed || applied.UserApprovalRequired || !applied.SnapshotRequired || applied.NotificationEvent != "repair-applied" {
		t.Fatalf("unexpected runtime-repair-applied rule: %#v", applied)
	}

	// portal-approval-required does not require a snapshot.
	portal, err := plan.RepairPlanPreview("portal-approval-required")
	if err != nil {
		t.Fatalf("RepairPlanPreview(portal-approval-required) returned error: %v", err)
	}
	if portal.SnapshotRequired || portal.SnapshotPlan != nil || portal.Severity != "info" {
		t.Fatalf("unexpected portal-approval-required rule: %#v", portal)
	}

	if _, err := plan.RepairPlanPreview("unknown-issue"); err == nil {
		t.Fatalf("RepairPlanPreview must reject unknown issue")
	}

	if err := validateNoBackendTerms(preview, "repair plan preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
