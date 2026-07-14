package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionTransactionPreviewBlocksCommitAndLaunch(t *testing.T) {
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

	preview, err := plan.ExecutionTransactionPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionTransactionPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.launch_transaction.v1" ||
		preview.RequestType != "execution-transaction-preview" ||
		preview.TransactionType != "compatibility-launch-transaction" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-resource-grant-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionTransactionPreview" {
		t.Fatalf("unexpected execution transaction schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution transaction identity: %#v", preview)
	}
	if preview.ResourceGrant.RequestType != "execution-resource-grant-preview" ||
		preview.ResourceGrant.GrantType != "compatibility-launch-resource-grant" ||
		preview.ResourceGrant.GrantCount != 7 ||
		preview.ResourceGrant.PortalRequiredCount != 6 ||
		preview.ResourceGrant.PendingReviewCount != 5 ||
		!preview.ResourceGrant.GrantPlanCreated ||
		preview.ResourceGrant.GrantObjectsCreated ||
		preview.ResourceGrant.PermissionGranted ||
		preview.ResourceGrant.ResourceBridgesEnabled ||
		!preview.ResourceGrant.PortalReviewRequired {
		t.Fatalf("unexpected resource grant summary: %#v", preview.ResourceGrant)
	}
	if preview.Readiness.RequestType != "execution-readiness-preview" ||
		preview.Readiness.ReadinessType != "compatibility-execution-readiness" ||
		preview.Readiness.ExecutionState != "blocked" ||
		preview.Readiness.OverallStatus != "not-ready" ||
		preview.Readiness.GateCount != 5 ||
		preview.Readiness.RequiredGateCount != 2 ||
		preview.Readiness.PendingGateCount != 1 ||
		preview.Readiness.BlockedGateCount != 1 ||
		preview.Readiness.LaunchAllowed ||
		preview.Readiness.LaunchEnabled ||
		preview.Readiness.BackendBindingReady ||
		!preview.Readiness.SnapshotRequired ||
		!preview.Readiness.PortalPolicyRequired {
		t.Fatalf("unexpected readiness summary: %#v", preview.Readiness)
	}
	if preview.BackendBinding.RecommendedProfileID != "local-compatibility" ||
		preview.BackendBinding.CandidateCount != 2 ||
		preview.BackendBinding.ReadyCandidateCount != 0 ||
		preview.BackendBinding.BlockedCandidateCount != 2 ||
		!preview.BackendBinding.BindingRequired ||
		preview.BackendBinding.BindingCommitted ||
		preview.BackendBinding.EnvironmentCreated ||
		preview.BackendBinding.BackendLaunchEnabled ||
		preview.BackendBinding.BackendDetailsExposed {
		t.Fatalf("unexpected backend binding summary: %#v", preview.BackendBinding)
	}
	if !preview.SnapshotBaseline.Required ||
		preview.SnapshotBaseline.State != "required" ||
		preview.SnapshotBaseline.Scope != "application-state-and-runtime-metadata" ||
		preview.SnapshotBaseline.BaselineCreated ||
		preview.SnapshotBaseline.RestorePointAvailable ||
		preview.SnapshotBaseline.UserDocumentsIncluded ||
		preview.SnapshotBaseline.HostSystemIncluded {
		t.Fatalf("unexpected snapshot baseline: %#v", preview.SnapshotBaseline)
	}
	if preview.WriteGate.MethodName != "Launch" ||
		preview.WriteGate.GateDecision != "blocked-until-production-backend" ||
		preview.WriteGate.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" ||
		preview.WriteGate.WriteMethodEnabled ||
		preview.WriteGate.DispatchEnabled ||
		preview.WriteGate.RequestObjectCreated {
		t.Fatalf("unexpected write gate: %#v", preview.WriteGate)
	}
	if preview.StepCount != 7 ||
		preview.PassedStepCount != 2 ||
		preview.RequiredStepCount != 2 ||
		preview.PendingStepCount != 2 ||
		preview.BlockedStepCount != 1 {
		t.Fatalf("unexpected transaction step counts: %#v", preview)
	}
	if len(preview.TransactionSteps) != 7 ||
		preview.TransactionSteps[0].ID != "identity-validation" ||
		preview.TransactionSteps[0].Status != "pass" ||
		preview.TransactionSteps[1].ID != "user-decision" ||
		preview.TransactionSteps[1].Status != "pass" ||
		preview.TransactionSteps[2].ID != "portal-resource-review" ||
		preview.TransactionSteps[2].Status != "required" ||
		preview.TransactionSteps[3].ID != "resource-grant-objects" ||
		preview.TransactionSteps[3].Status != "pending" ||
		preview.TransactionSteps[4].ID != "snapshot-baseline" ||
		preview.TransactionSteps[5].ID != "backend-binding" ||
		preview.TransactionSteps[6].ID != "runtime-launch-write-gate" ||
		preview.TransactionSteps[6].Status != "blocked" {
		t.Fatalf("unexpected transaction steps: %#v", preview.TransactionSteps)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		!preview.TransactionPlanCreated || preview.TransactionCommitted ||
		preview.ReadinessPassed || preview.ResourceGrantsCommitted ||
		preview.PortalApprovalRecorded || preview.SnapshotBaselineCreated ||
		preview.BackendBindingCommitted || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.RequestObjectsCreated ||
		preview.HostPermissionChanged || preview.HostRootModified ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution transaction safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "commit launch transaction from preview") ||
		!containsString(preview.BlockedActions, "create restore point from transaction preview") ||
		!containsString(preview.BlockedActions, "commit compatibility profile binding from transaction preview") ||
		!containsString(preview.BlockedActions, "start compatibility profile from transaction preview") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionTransactionPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionTransactionPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.TransactionSteps[1].Status != "blocked" ||
		rejectedPreview.BlockedStepCount != 2 ||
		rejectedPreview.TransactionCommitted ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected transaction preview: %#v", rejectedPreview)
	}

	deferredPreview, err := plan.ExecutionTransactionPreview("deferred", nil)
	if err != nil {
		t.Fatalf("deferred ExecutionTransactionPreview returned error: %v", err)
	}
	if deferredPreview.UserDecisionAllowsLaunch ||
		deferredPreview.TransactionSteps[1].Status != "pending" ||
		deferredPreview.PendingStepCount != 3 ||
		deferredPreview.TransactionCommitted ||
		deferredPreview.ExecutionStarted {
		t.Fatalf("unexpected deferred transaction preview: %#v", deferredPreview)
	}

	if _, err := plan.ExecutionTransactionPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionTransactionPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionTransactionPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionTransactionPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution transaction preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
