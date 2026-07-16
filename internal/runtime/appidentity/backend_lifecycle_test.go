package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/environment"
)

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
		preview.OverallStatus != "not-ready" ||
		preview.StateRootBacked ||
		preview.StateRootPathExposed ||
		preview.EnvironmentProfile != "local-compatibility" ||
		len(preview.SatisfiedGates) != 0 ||
		len(preview.PendingGates) != 4 {
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

func TestBackendLifecyclePreviewConsumesEnvironmentStateRoot(t *testing.T) {
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
	stateRoot := t.TempDir()
	lifecycle, err := environment.New(stateRoot)
	if err != nil {
		t.Fatalf("environment.New returned error: %v", err)
	}
	if _, err := lifecycle.Plan("org.example.ledger", environment.ProfileLocal); err != nil {
		t.Fatalf("Plan returned error: %v", err)
	}
	if _, err := lifecycle.Stage("org.example.ledger", environment.ProfileLocal); err != nil {
		t.Fatalf("Stage returned error: %v", err)
	}
	for _, gate := range environment.RequiredGates() {
		if _, err := lifecycle.SatisfyGate("org.example.ledger", environment.ProfileLocal, gate); err != nil {
			t.Fatalf("SatisfyGate(%s) returned error: %v", gate, err)
		}
	}
	if _, err := lifecycle.MarkReady("org.example.ledger", environment.ProfileLocal); err != nil {
		t.Fatalf("MarkReady returned error: %v", err)
	}

	preview, err := plan.BackendLifecyclePreviewWithStateRoot(stateRoot)
	if err != nil {
		t.Fatalf("BackendLifecyclePreviewWithStateRoot returned error: %v", err)
	}
	if preview.Source != "run-plan-preview+go-runtime-backend-lifecycle+environment-state-root" ||
		preview.LifecycleState != "ready" ||
		preview.OverallStatus != "ready-with-launch-disabled" ||
		!preview.StateRootBacked ||
		preview.StateRootPathExposed ||
		preview.EnvironmentProfile != "local-compatibility" ||
		!preview.StateRootReady ||
		!preview.BackendBindingReady ||
		preview.PortalReviewRequired ||
		preview.SnapshotRequired ||
		preview.LaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected state-backed lifecycle preview: %#v", preview)
	}
	if len(preview.SatisfiedGates) != 4 ||
		len(preview.PendingGates) != 0 ||
		len(preview.RepairHints) != 0 ||
		preview.BlockReason != "" {
		t.Fatalf("unexpected state-backed gate projection: %#v", preview)
	}
	if preview.Stages[1].Status != "pass" ||
		preview.Stages[2].Status != "pass" ||
		preview.Stages[3].Status != "pass" ||
		preview.Stages[4].Status != "blocked" {
		t.Fatalf("unexpected state-backed stage statuses: %#v", preview.Stages)
	}
	if err := validateNoBackendTerms(preview, "state-backed backend lifecycle preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRecordBackendLifecycleStatePersistsActionsAndPreview(t *testing.T) {
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
	stateRoot := t.TempDir()

	planned, err := plan.RecordBackendLifecycleState(stateRoot, "plan", "", "", "")
	if err != nil {
		t.Fatalf("RecordBackendLifecycleState plan returned error: %v", err)
	}
	if planned.SchemaVersion != "xnix.runtime.backend_lifecycle_record.v1" ||
		planned.RecordType != "backend-lifecycle-state-record" ||
		planned.Source != "go-runtime-state-root-backend-lifecycle" ||
		planned.Action != "plan" ||
		planned.RelativePath != "environments/org.example.ledger__local-compatibility.json" ||
		planned.EnvironmentRecord.State != environment.StatePlanned ||
		planned.Preview.LifecycleState != "planned" ||
		!planned.RuntimeOwned ||
		!planned.GoRuntimeBacked ||
		planned.KDEPolicyOwner ||
		planned.StateRootPathExposed ||
		planned.LaunchEnabled ||
		planned.BackendProcessStarted ||
		planned.HostRootModified ||
		planned.BackendDetailsExposed ||
		planned.NetworkRequired ||
		planned.PrivilegedContainerRequired {
		t.Fatalf("unexpected planned lifecycle record: %#v", planned)
	}

	staged, err := plan.RecordBackendLifecycleState(stateRoot, "stage", "", "", "")
	if err != nil {
		t.Fatalf("RecordBackendLifecycleState stage returned error: %v", err)
	}
	if staged.EnvironmentRecord.State != environment.StateStaged ||
		staged.Preview.LifecycleState != "staged" ||
		!staged.Preview.StateRootBacked ||
		staged.Preview.LaunchEnabled ||
		staged.Preview.BackendProcessStarted ||
		staged.Preview.HostRootModified ||
		staged.Preview.BackendDetailsExposed {
		t.Fatalf("unexpected staged lifecycle record: %#v", staged)
	}

	for _, gate := range environment.RequiredGates() {
		if _, err := plan.RecordBackendLifecycleState(stateRoot, "satisfy-gate", gate, "", ""); err != nil {
			t.Fatalf("RecordBackendLifecycleState satisfy-gate(%s) returned error: %v", gate, err)
		}
	}
	ready, err := plan.RecordBackendLifecycleState(stateRoot, "mark-ready", "", "", "")
	if err != nil {
		t.Fatalf("RecordBackendLifecycleState mark-ready returned error: %v", err)
	}
	if ready.EnvironmentRecord.State != environment.StateReady ||
		ready.Preview.OverallStatus != "ready-with-launch-disabled" ||
		ready.Preview.LaunchEnabled ||
		ready.Preview.BackendProcessStarted {
		t.Fatalf("unexpected ready lifecycle record: %#v", ready)
	}

	inspected, err := plan.RecordBackendLifecycleState(stateRoot, "inspect", "", "", "")
	if err != nil {
		t.Fatalf("RecordBackendLifecycleState inspect returned error: %v", err)
	}
	if inspected.EnvironmentRecord.State != environment.StateReady ||
		inspected.Preview.OverallStatus != "ready-with-launch-disabled" {
		t.Fatalf("inspect must preserve persisted ready state: %#v", inspected)
	}
}

func TestRecordBackendLifecycleStateRejectsUnsafeOrIncompleteActions(t *testing.T) {
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
	stateRoot := t.TempDir()

	if _, err := plan.RecordBackendLifecycleState("", "plan", "", "", ""); err == nil {
		t.Fatalf("missing state root must be rejected")
	}
	if _, err := plan.RecordBackendLifecycleState(stateRoot, "satisfy-gate", "", "", ""); err == nil {
		t.Fatalf("satisfy-gate without gate must be rejected")
	}
	if _, err := plan.RecordBackendLifecycleState(stateRoot, "flag-repair", "", "", ""); err == nil {
		t.Fatalf("flag-repair without hint must be rejected")
	}
	if _, err := plan.RecordBackendLifecycleState(stateRoot, "block", "", "", ""); err == nil {
		t.Fatalf("block without reason must be rejected")
	}
	if _, err := plan.RecordBackendLifecycleState(stateRoot, "unknown", "", "", ""); err == nil {
		t.Fatalf("unknown action must be rejected")
	}
}
