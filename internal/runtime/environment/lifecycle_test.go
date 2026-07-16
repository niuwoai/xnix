package environment

import (
	"testing"
)

const app = "org.example.ledger"

func TestFreshEnvironmentIsMissing(t *testing.T) {
	lc, err := New(t.TempDir())
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	rec, err := lc.Get(app, ProfileLocal)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if rec.State != StateMissing || rec.LaunchEnabled {
		t.Fatalf("fresh environment must be missing and non-launchable: %#v", rec)
	}
}

func TestHappyPathToReadySurvivesReload(t *testing.T) {
	root := t.TempDir()
	lc, _ := New(root)

	if _, err := lc.Plan(app, ProfileLocal); err != nil {
		t.Fatalf("Plan: %v", err)
	}
	if _, err := lc.Stage(app, ProfileLocal); err != nil {
		t.Fatalf("Stage: %v", err)
	}
	// Cannot become ready until all gates are satisfied.
	if _, err := lc.MarkReady(app, ProfileLocal); err == nil {
		t.Fatalf("MarkReady must fail with pending gates")
	}
	for _, gate := range RequiredGates() {
		if _, err := lc.SatisfyGate(app, ProfileLocal, gate); err != nil {
			t.Fatalf("SatisfyGate(%s): %v", gate, err)
		}
	}
	rec, err := lc.MarkReady(app, ProfileLocal)
	if err != nil {
		t.Fatalf("MarkReady: %v", err)
	}
	if !rec.Ready() || rec.LaunchEnabled || len(rec.PendingGates()) != 0 {
		t.Fatalf("ready environment wrong: %#v", rec)
	}

	// State must survive a fresh store over the same root.
	reloaded, _ := New(root)
	got, err := reloaded.Get(app, ProfileLocal)
	if err != nil {
		t.Fatalf("reload Get: %v", err)
	}
	if got.State != StateReady {
		t.Fatalf("state did not survive reload: %#v", got)
	}
}

func TestInvalidTransitionRejected(t *testing.T) {
	lc, _ := New(t.TempDir())
	// missing -> staged is not allowed (must plan first).
	if _, err := lc.Stage(app, ProfileLocal); err == nil {
		t.Fatalf("missing -> staged must be rejected")
	}
	if _, err := lc.Plan(app, ProfileLocal); err != nil {
		t.Fatalf("Plan: %v", err)
	}
	// planned -> ready is not allowed (must stage first).
	if _, err := lc.MarkReady(app, ProfileLocal); err == nil {
		t.Fatalf("planned -> ready must be rejected")
	}
}

func TestRepairFlowClearsGatesAndReturnsToStaged(t *testing.T) {
	lc, _ := New(t.TempDir())
	lc.Plan(app, ProfileIsolated)
	lc.Stage(app, ProfileIsolated)
	lc.SatisfyGate(app, ProfileIsolated, "recipe-trust")

	rec, err := lc.FlagRepair(app, ProfileIsolated, []string{"reinstall compatibility engine binding"})
	if err != nil {
		t.Fatalf("FlagRepair: %v", err)
	}
	if rec.State != StateRepairRequired || len(rec.RepairHints) != 1 || len(rec.SatisfiedGates) != 0 {
		t.Fatalf("repair must clear gates and set hints: %#v", rec)
	}
	// Repair with no hints is rejected.
	if _, err := lc.FlagRepair(app, ProfileIsolated, nil); err == nil {
		t.Fatalf("repair without hints must be rejected")
	}
	// repair-required -> staged is allowed.
	staged, err := lc.Stage(app, ProfileIsolated)
	if err != nil {
		t.Fatalf("Stage after repair: %v", err)
	}
	if staged.State != StateStaged || len(staged.RepairHints) != 0 {
		t.Fatalf("staged after repair wrong: %#v", staged)
	}
}

func TestBlockAndReset(t *testing.T) {
	lc, _ := New(t.TempDir())
	if _, err := lc.Block(app, ProfileLocal, "policy hold"); err != nil {
		t.Fatalf("Block: %v", err)
	}
	rec, _ := lc.Get(app, ProfileLocal)
	if rec.State != StateBlocked || rec.BlockReason != "policy hold" {
		t.Fatalf("block wrong: %#v", rec)
	}
	// Block with empty reason is rejected.
	if _, err := lc.Block(app, ProfileLocal, "  "); err == nil {
		t.Fatalf("block without reason must be rejected")
	}
	// blocked -> planned is allowed (reset).
	if _, err := lc.Plan(app, ProfileLocal); err != nil {
		t.Fatalf("reset via Plan: %v", err)
	}
}

func TestRetireDecommissionsAndReactivates(t *testing.T) {
	lc, _ := New(t.TempDir())
	lc.Plan(app, ProfileLocal)
	lc.Stage(app, ProfileLocal)
	lc.SatisfyGate(app, ProfileLocal, "recipe-trust")

	retired, err := lc.Retire(app, ProfileLocal)
	if err != nil {
		t.Fatalf("Retire: %v", err)
	}
	if retired.State != StateRetired || len(retired.SatisfiedGates) != 0 || retired.LaunchEnabled {
		t.Fatalf("retired environment wrong: %#v", retired)
	}

	// A retired environment survives reload.
	got, err := lc.Get(app, ProfileLocal)
	if err != nil || got.State != StateRetired {
		t.Fatalf("retired state must persist: %#v err=%v", got, err)
	}

	// retired -> planned reactivates.
	planned, err := lc.Plan(app, ProfileLocal)
	if err != nil {
		t.Fatalf("reactivate via Plan: %v", err)
	}
	if planned.State != StatePlanned {
		t.Fatalf("reactivation did not plan: %#v", planned)
	}
}

func TestRetireRejectedFromMissing(t *testing.T) {
	lc, _ := New(t.TempDir())
	// missing -> retired is not an allowed transition.
	if _, err := lc.Retire(app, ProfileLocal); err == nil {
		t.Fatalf("retiring a missing environment must be rejected")
	}
}

func TestGatesOnlyWhileStagedOrReadyAndUnknownRejected(t *testing.T) {
	lc, _ := New(t.TempDir())
	lc.Plan(app, ProfileLocal)
	// Satisfying a gate while merely planned is rejected.
	if _, err := lc.SatisfyGate(app, ProfileLocal, "recipe-trust"); err == nil {
		t.Fatalf("gate while planned must be rejected")
	}
	lc.Stage(app, ProfileLocal)
	if _, err := lc.SatisfyGate(app, ProfileLocal, "no-such-gate"); err == nil {
		t.Fatalf("unknown gate must be rejected")
	}
}

func TestListReturnsPersistedRecords(t *testing.T) {
	lc, _ := New(t.TempDir())
	lc.Plan(app, ProfileLocal)
	lc.Plan("org.example.other", ProfileIsolated)
	list, err := lc.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("expected 2 records: %#v", list)
	}
}

func TestInvalidInputsRejected(t *testing.T) {
	lc, _ := New(t.TempDir())
	if _, err := lc.Plan("bad id", ProfileLocal); err == nil {
		t.Fatalf("bad application id must be rejected")
	}
	if _, err := lc.Plan(app, Profile("bogus")); err == nil {
		t.Fatalf("bad profile must be rejected")
	}
	if _, err := New(""); err == nil {
		t.Fatalf("empty state root must be rejected")
	}
}
