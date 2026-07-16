package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func newFallbackPlan(t *testing.T, mode string) Plan {
	t.Helper()
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                mode,
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan(%q): %v", mode, err)
	}
	return plan
}

func fallbackCandidate(t *testing.T, preview CompatibilityBackendFallbackPreview, id string) CompatibilityBackendFallbackCandidate {
	t.Helper()
	for _, candidate := range preview.Candidates {
		if candidate.ID == id {
			return candidate
		}
	}
	t.Fatalf("fallback candidate %q not present; ids=%v", id, preview.CandidateIDs)
	return CompatibilityBackendFallbackCandidate{}
}

func assertFallbackSafe(t *testing.T, preview CompatibilityBackendFallbackPreview) {
	t.Helper()
	if preview.SelectionPersisted || preview.EngineInstallEnabled || preview.BackendLaunchEnabled ||
		preview.VMStartEnabled || preview.BackendDetailsExposed || preview.NetworkRequired ||
		preview.PrivilegedContainerRequired || preview.StateRootPathExposed || preview.HostRootModified {
		t.Fatalf("fallback preview must keep every unsafe capability disabled: %+v", preview)
	}
	if !preview.ReviewOnly || !preview.RuntimeOwned || !preview.GoRuntimeBacked {
		t.Fatalf("fallback preview must be a runtime-owned review-only read model: %+v", preview)
	}
	if err := safety.ValidatePayload("compatibility backend fallback preview", preview); err != nil {
		t.Fatalf("fallback preview leaks forbidden content: %v", err)
	}
}

func TestCompatibilityBackendFallbackLocalPrimaryStaysBlocked(t *testing.T) {
	plan := newFallbackPlan(t, "automatic")
	preview, err := plan.CompatibilityBackendFallbackPreview(CompatibilityBackendFallbackOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.PrimaryProfileID != "local-compatibility" || preview.FallbackProfileID != "isolated-compatibility" {
		t.Fatalf("automatic recipe should prefer local with isolated fallback: %+v", preview)
	}
	if preview.IsolationRequired {
		t.Fatalf("automatic recipe should not require isolation")
	}
	if preview.SelectedProfileID != "" || preview.FallbackAvailable {
		t.Fatalf("no profile should be selectable without evidence: %+v", preview)
	}
	if preview.OverallState != "blocked-missing-evidence" {
		t.Fatalf("both candidates lacking evidence should stay blocked, got %q", preview.OverallState)
	}
	local := fallbackCandidate(t, preview, "local-compatibility")
	if local.FallbackState != "blocked-missing-evidence" || !containsFallbackReason(local.ReasonCodes, "local-selected-by-default") {
		t.Fatalf("local candidate should be the blocked default: %+v", local)
	}
	assertFallbackSafe(t, preview)
}

func TestCompatibilityBackendFallbackIsolationRequiredForbidsLocal(t *testing.T) {
	plan := newFallbackPlan(t, "vm")
	preview, err := plan.CompatibilityBackendFallbackPreview(CompatibilityBackendFallbackOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.PrimaryProfileID != "isolated-compatibility" || !preview.IsolationRequired {
		t.Fatalf("vm recipe should require isolation: %+v", preview)
	}
	if preview.OverallState != "blocked-isolation-required" {
		t.Fatalf("isolation-required audit should be blocked-isolation-required, got %q", preview.OverallState)
	}
	local := fallbackCandidate(t, preview, "local-compatibility")
	if local.FallbackState != "blocked-by-policy" || !containsFallbackReason(local.ReasonCodes, "recipe-requests-isolation") {
		t.Fatalf("local must be forbidden by policy under isolation: %+v", local)
	}
	assertFallbackSafe(t, preview)
}

func TestCompatibilityBackendFallbackDiagnosticsRiskIsSurfaced(t *testing.T) {
	plan := newFallbackPlan(t, "automatic")
	preview, err := plan.CompatibilityBackendFallbackPreview(CompatibilityBackendFallbackOptions{DiagnosticsFailingRuns: 2})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	local := fallbackCandidate(t, preview, "local-compatibility")
	if !containsFallbackReason(local.ReasonCodes, "diagnostics-regressions-present") {
		t.Fatalf("diagnostics regressions should be surfaced as a fallback risk: %+v", local.ReasonCodes)
	}
	if !containsFallbackReason(preview.ReasonCodes, "diagnostics-regressions-present") {
		t.Fatalf("aggregate reasons should include diagnostics risk: %+v", preview.ReasonCodes)
	}
	assertFallbackSafe(t, preview)
}

func TestCompatibilityBackendFallbackAutomaticExplainsBlocked(t *testing.T) {
	plan := newFallbackPlan(t, "automatic")
	preview, err := plan.CompatibilityBackendFallbackPreview(CompatibilityBackendFallbackOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	automatic := fallbackCandidate(t, preview, "automatic")
	if automatic.Role != "orchestrator" || automatic.FallbackState != "blocked" {
		t.Fatalf("automatic orchestrator should report blocked: %+v", automatic)
	}
	assertFallbackSafe(t, preview)
}

func containsFallbackReason(reasons []string, want string) bool {
	for _, reason := range reasons {
		if reason == want {
			return true
		}
	}
	return false
}
