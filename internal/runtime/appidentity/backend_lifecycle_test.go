package appidentity

import "testing"

func TestBackendLifecyclePreviewStaysBlockedBeforeGates(t *testing.T) {
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

	preview, err := plan.BackendLifecyclePreview()
	if err != nil {
		t.Fatalf("BackendLifecyclePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.backend_lifecycle.v1" ||
		preview.RequestType != "backend-lifecycle-preview" ||
		preview.LifecycleType != "compatibility-backend-lifecycle" ||
		preview.Source != "run-plan-preview+go-runtime-backend-lifecycle" ||
		preview.RuntimeMethod != "GetBackendLifecycle" ||
		preview.ReadMethod != "GetBackendLifecyclePreview" {
		t.Fatalf("unexpected backend lifecycle schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.RequestedMode != "automatic" ||
		preview.SelectedStrategy != "automatic-managed" ||
		preview.LifecycleState != "blocked" ||
		preview.OverallStatus != "not-ready" {
		t.Fatalf("unexpected backend lifecycle identity: %#v", preview)
	}
	if !sameStrings(preview.StageIDs, []string{
		"recipe-loaded",
		"state-root-ready",
		"backend-binding-ready",
		"portal-and-snapshot-review",
		"runtime-launch-write-gate",
	}) {
		t.Fatalf("unexpected backend lifecycle stages: %#v", preview.StageIDs)
	}
	if preview.Stages[0].Status != "pass" || preview.Stages[4].Status != "blocked" {
		t.Fatalf("unexpected stage statuses: %#v", preview.Stages)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.BackendBindingReady ||
		preview.LaunchEnabled ||
		preview.ExecutionRequestCreated ||
		preview.BackendProcessStarted ||
		preview.LocalBackendStarted ||
		preview.IsolatedBackendStarted ||
		preview.StateRootReady ||
		!preview.PortalReviewRequired ||
		!preview.SnapshotRequired ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected backend lifecycle safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 5 {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	if err := validateNoBackendTerms(preview, "backend lifecycle preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}

	isolated, err := NewPlan(Recipe{ID: "org.example.iso", Name: "Iso", Icon: "application-x-executable", Mode: "vm", SupportedExtensions: []string{".abc"}})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolated.BackendLifecyclePreview()
	if err != nil {
		t.Fatalf("isolated BackendLifecyclePreview returned error: %v", err)
	}
	if isolatedPreview.SelectedStrategy != "isolated-compatible-managed" {
		t.Fatalf("isolated recipe did not map to isolated strategy: %#v", isolatedPreview.SelectedStrategy)
	}
}
