package diagnostics

import "testing"

const app = "org.example.ledger"

func passFixture() Fixture {
	return Fixture{TestType: "preflight", Signals: []Signal{
		{ID: "recipe-validation", Category: "recipe", Outcome: OutcomePass, Summary: "Recipe metadata is valid."},
		{ID: "portal-preflight", Category: "portal", Outcome: OutcomePass, Summary: "Portal request model is available."},
	}}
}

func failFixture() Fixture {
	return Fixture{TestType: "smoke", Signals: []Signal{
		{ID: "recipe-validation", Category: "recipe", Outcome: OutcomePass, Summary: "Recipe metadata is valid."},
		{ID: "runtime-launch-binding", Category: "backend", Outcome: OutcomeFail, Summary: "Launch binding is not ready."},
	}}
}

func TestRunFoldsOutcomesWithoutBackend(t *testing.T) {
	res, err := Run(app, passFixture())
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if res.Overall != OutcomePass || res.BackendStarted || res.HostRootModified || len(res.FailingIDs) != 0 {
		t.Fatalf("healthy run wrong: %#v", res)
	}

	failRes, err := Run(app, failFixture())
	if err != nil {
		t.Fatalf("Run fail: %v", err)
	}
	if failRes.Overall != OutcomeFail || len(failRes.FailingIDs) != 1 || failRes.FailingIDs[0] != "runtime-launch-binding" {
		t.Fatalf("failing run wrong: %#v", failRes)
	}
}

func TestRunRejectsUnsafeSignalContent(t *testing.T) {
	unsafe := Fixture{TestType: "preflight", Signals: []Signal{
		{ID: "leak", Outcome: OutcomePass, Summary: "wrote to /home/user/secret.txt"},
	}}
	if _, err := Run(app, unsafe); err == nil {
		t.Fatalf("host-path leak in signal must be rejected")
	}
	backendLeak := Fixture{TestType: "preflight", Signals: []Signal{
		{ID: "leak", Outcome: OutcomePass, Summary: "started wine prefix"},
	}}
	if _, err := Run(app, backendLeak); err == nil {
		t.Fatalf("backend term in signal must be rejected")
	}
}

func TestRunValidatesInputs(t *testing.T) {
	if _, err := Run("bad id", passFixture()); err == nil {
		t.Fatalf("bad application id must be rejected")
	}
	if _, err := Run(app, Fixture{TestType: "nope", Signals: passFixture().Signals}); err == nil {
		t.Fatalf("bad test type must be rejected")
	}
	if _, err := Run(app, Fixture{TestType: "preflight"}); err == nil {
		t.Fatalf("empty signals must be rejected")
	}
}

func TestResultStorePersistence(t *testing.T) {
	store, err := NewResultStore(t.TempDir())
	if err != nil {
		t.Fatalf("NewResultStore: %v", err)
	}
	res, _ := Run(app, failFixture())
	if err := store.Save(res); err != nil {
		t.Fatalf("Save: %v", err)
	}
	loaded, err := store.Load(app, "smoke")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.Overall != OutcomeFail || loaded.ApplicationID != app {
		t.Fatalf("persisted result wrong: %#v", loaded)
	}
	if _, err := store.Load(app, "preflight"); err == nil {
		t.Fatalf("loading a missing result must error")
	}
}

func TestRecommendFromFailingSignals(t *testing.T) {
	res, _ := Run(app, failFixture())
	rec, ok := Recommend(res)
	if !ok {
		t.Fatalf("failing result must yield a recommendation")
	}
	if rec.Issue != "engine-binding-pending" || !rec.SnapshotRequired || rec.AutoApplyAllowed {
		t.Fatalf("recommendation wrong: %#v", rec)
	}
	// Healthy result yields no recommendation.
	pass, _ := Run(app, passFixture())
	if _, ok := Recommend(pass); ok {
		t.Fatalf("healthy result must not recommend a repair")
	}
	// Blocked-but-not-failed maps to a portal issue.
	blocked := Fixture{TestType: "preflight", Signals: []Signal{
		{ID: "portal", Outcome: OutcomeBlocked, Summary: "Awaiting desktop permission."},
	}}
	bres, _ := Run(app, blocked)
	brec, _ := Recommend(bres)
	if brec.Issue != "portal-approval-required" {
		t.Fatalf("blocked result should map to portal issue: %#v", brec)
	}
}
