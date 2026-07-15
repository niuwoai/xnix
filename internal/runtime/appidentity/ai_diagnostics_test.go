package appidentity

import "testing"

func TestAIDiagnosticInputPreviewUsesRuntimeSafeMetadata(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)

	preview, err := plan.AIDiagnosticInputPreview("engine-binding-pending", "preflight")
	if err != nil {
		t.Fatalf("AIDiagnosticInputPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.ai_diagnostic_input.v1" ||
		preview.RequestType != "ai-diagnostic-input-preview" ||
		preview.InputType != "ai-diagnostic-input" ||
		preview.RuntimeMethod != "GetAIDiagnosticInput" ||
		preview.ReadMethod != "GetAIDiagnosticInputPreview" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.Issue != "engine-binding-pending" ||
		preview.TestType != "preflight" {
		t.Fatalf("unexpected AI diagnostic input schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.NetworkRequired ||
		!preview.SafeForAIDiagnostics ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected AI diagnostic input safety flags: %#v", preview)
	}
	if len(preview.ContextSections) != 4 ||
		preview.ContextSections[0].ID != "recipe" ||
		len(preview.DiagnosticSignals) != 3 ||
		preview.DiagnosticSignals[0].ID != "pending-runtime-launch-binding" ||
		preview.PrivacyBoundaries.UserDocumentsIncluded ||
		preview.PrivacyBoundaries.HostPathsIncluded ||
		preview.PrivacyBoundaries.NetworkCallsAllowed ||
		!preview.PrivacyBoundaries.RequiresUserApprovalForSensitiveTasks {
		t.Fatalf("unexpected AI diagnostic input content: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "AI diagnostic input preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestAIDiagnosticRecommendationPreviewRequiresReviewOnly(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)

	preview, err := plan.AIDiagnosticRecommendationPreview("engine-binding-pending", "smoke")
	if err != nil {
		t.Fatalf("AIDiagnosticRecommendationPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.ai_diagnostic_recommendation.v1" ||
		preview.RequestType != "ai-diagnostic-recommendation-preview" ||
		preview.RecommendationType != "ai-diagnostic-recommendation" ||
		preview.RuntimeMethod != "GetAIDiagnosticRecommendation" ||
		preview.ReadMethod != "GetAIDiagnosticRecommendationPreview" {
		t.Fatalf("unexpected AI diagnostic recommendation schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.NetworkRequired ||
		!preview.SafeForAIDiagnostics ||
		preview.AutoExecutionAllowed ||
		preview.RepairExecutionRequested ||
		preview.RepairExecuted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected AI diagnostic recommendation safety flags: %#v", preview)
	}
	if len(preview.Recommendations) != 3 ||
		len(preview.ApprovalRequiredActions) != 1 ||
		preview.ApprovalRequiredActions[0].ID != "prepare-safe-restore-point" {
		t.Fatalf("unexpected AI diagnostic recommendations: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "AI diagnostic recommendation preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestAIRepairApprovalGatePreviewBlocksExecution(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)

	preview, err := plan.AIRepairApprovalGatePreview("engine-binding-pending", "repair-readiness")
	if err != nil {
		t.Fatalf("AIRepairApprovalGatePreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.ai_repair_approval_gate.v1" ||
		preview.RequestType != "ai-repair-approval-gate-preview" ||
		preview.GateType != "ai-repair-approval-gate" ||
		preview.RuntimeMethod != "GetAIRepairApprovalGate" ||
		preview.ReadMethod != "GetAIRepairApprovalGatePreview" ||
		preview.GateDecision != "blocked-until-approval" ||
		preview.ApprovalSurface != "Compatibility Center" {
		t.Fatalf("unexpected AI repair approval gate schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.NetworkRequired ||
		!preview.SafeForAIDiagnostics ||
		preview.RepairExecutionRequested ||
		preview.RepairExecuted ||
		preview.AutoExecutionAllowed ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected AI repair approval gate safety flags: %#v", preview)
	}
	if len(preview.RequiredGates) != 3 ||
		preview.RequiredGates[0].ID != "compatibility-center-review" ||
		len(preview.ApprovalRequiredActions) != 1 ||
		!containsString(preview.BlockedActions, "execute repair without approval") {
		t.Fatalf("unexpected AI repair approval gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "AI repair approval gate preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func testAIDiagnosticsPlan(t *testing.T) Plan {
	t.Helper()
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc"},
	}
	plan, err := NewPlanWithProvenance(recipe, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}
