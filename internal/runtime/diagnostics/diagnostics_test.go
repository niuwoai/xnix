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

func TestRunRecordStorePersistsSafeDiagnosticReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	store, err := NewRunRecordStore(stateRoot)
	if err != nil {
		t.Fatalf("NewRunRecordStore: %v", err)
	}

	record, err := store.Record(RunRecordRequest{
		ApplicationID: app,
		RunID:         "run-001",
		Fixture:       failFixture(),
	})
	if err != nil {
		t.Fatalf("Record: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.diagnostic_run_record.v1" ||
		record.RecordType != "diagnostic-run-record" ||
		record.Source != "go-runtime-state-root-diagnostic-run-record" ||
		record.RelativePath != "diagnostics-ledger/runs/run-001.json" ||
		record.SHA256 == "" ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		record.StateRootPathExposed ||
		record.FixturePathExposed ||
		record.BackendStarted ||
		record.AIProviderCalled ||
		record.RealAIProviderEnabled ||
		record.AutoRepairAllowed ||
		record.RepairExecuted ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.BackendDetailsExposed ||
		record.FileContentsIncluded {
		t.Fatalf("unsafe or unexpected diagnostic run record: %#v", record)
	}
	if record.Result.Overall != OutcomeFail ||
		record.DiagnosticInput.NetworkRequired ||
		record.DiagnosticInput.FileContentsIncluded ||
		record.RepairRecommendation == nil ||
		record.RepairRecommendation.Issue != "engine-binding-pending" ||
		!record.RepairRecommendation.SnapshotRequired ||
		record.RepairRecommendation.AutoApplyAllowed {
		t.Fatalf("unexpected diagnostic payload: %#v", record)
	}

	loaded, err := store.Load("run-001")
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.RunID != "run-001" || loaded.Result.ApplicationID != app {
		t.Fatalf("loaded record mismatch: %#v", loaded)
	}
	records, err := store.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(records) != 1 || records[0].RunID != "run-001" {
		t.Fatalf("unexpected record list: %#v", records)
	}

	resultStore, err := NewResultStore(stateRoot)
	if err != nil {
		t.Fatalf("NewResultStore: %v", err)
	}
	result, err := resultStore.Load(app, "smoke")
	if err != nil {
		t.Fatalf("Load persisted result: %v", err)
	}
	if result.Overall != OutcomeFail {
		t.Fatalf("persisted result mismatch: %#v", result)
	}
}

func TestRunRecordStoreRejectsUnsafeRootsAndIDs(t *testing.T) {
	if _, err := NewRunRecordStore(""); err == nil {
		t.Fatalf("empty state root must be rejected")
	}
	if _, err := NewRunRecordStore("/"); err == nil {
		t.Fatalf("filesystem root must be rejected")
	}
	store, err := NewRunRecordStore(t.TempDir())
	if err != nil {
		t.Fatalf("NewRunRecordStore: %v", err)
	}
	if _, err := store.Record(RunRecordRequest{ApplicationID: app, RunID: "../escape", Fixture: passFixture()}); err == nil {
		t.Fatalf("path traversal run id must be rejected")
	}
}

func TestRunRecordHistorySummarizesAndFiltersRecords(t *testing.T) {
	store, err := NewRunRecordStore(t.TempDir())
	if err != nil {
		t.Fatalf("NewRunRecordStore: %v", err)
	}
	if _, err := store.Record(RunRecordRequest{ApplicationID: app, RunID: "001-pass", Fixture: passFixture()}); err != nil {
		t.Fatalf("Record pass: %v", err)
	}
	if _, err := store.Record(RunRecordRequest{ApplicationID: app, RunID: "002-fail", Fixture: failFixture()}); err != nil {
		t.Fatalf("Record fail: %v", err)
	}
	otherFixture := Fixture{TestType: "repair-readiness", Signals: []Signal{
		{ID: "snapshot-baseline", Category: "snapshot", Outcome: OutcomeBlocked, Summary: "Awaiting restore point."},
	}}
	if _, err := store.Record(RunRecordRequest{ApplicationID: "org.example.viewer", RunID: "003-blocked", Fixture: otherFixture}); err != nil {
		t.Fatalf("Record other app: %v", err)
	}

	history, err := store.History(app)
	if err != nil {
		t.Fatalf("History: %v", err)
	}
	if history.SchemaVersion != "xnix.runtime.diagnostic_run_history.v1" ||
		history.RecordType != "diagnostic-run-history" ||
		history.Source != "go-runtime-state-root-diagnostic-run-history" ||
		history.ApplicationID != app ||
		!history.RuntimeOwned ||
		!history.GoRuntimeBacked ||
		history.KDEPolicyOwner ||
		history.StateRootPathExposed ||
		history.BackendStarted ||
		history.AIProviderCalled ||
		history.RealAIProviderEnabled ||
		history.AutoRepairAllowed ||
		history.RepairExecuted ||
		history.HostRootModified ||
		history.NetworkRequired ||
		history.PrivilegedContainerRequired ||
		history.BackendDetailsExposed ||
		history.FileContentsIncluded {
		t.Fatalf("unsafe or unexpected history: %#v", history)
	}
	if history.Counts.Total != 2 || history.Counts.Passed != 1 || history.Counts.Failed != 1 || len(history.Records) != 2 {
		t.Fatalf("unexpected filtered history counts: %#v", history)
	}
	if history.Latest == nil || history.Latest.RunID != "002-fail" || history.Latest.RepairIssue != "engine-binding-pending" || !history.Latest.SnapshotRequired {
		t.Fatalf("unexpected latest history record: %#v", history.Latest)
	}
	if history.Records[0].RelativePath != "diagnostics-ledger/runs/001-pass.json" ||
		history.Records[1].RelativePath != "diagnostics-ledger/runs/002-fail.json" ||
		history.Records[1].BackendStarted ||
		history.Records[1].AIProviderCalled ||
		history.Records[1].RepairExecuted ||
		history.Records[1].HostRootModified ||
		history.Records[1].BackendDetailsExposed ||
		history.Records[1].FileContentsIncluded {
		t.Fatalf("unexpected history records: %#v", history.Records)
	}

	all, err := store.History("")
	if err != nil {
		t.Fatalf("History all: %v", err)
	}
	if all.Counts.Total != 3 || all.Counts.Blocked != 1 {
		t.Fatalf("unexpected all-history counts: %#v", all.Counts)
	}
	if _, err := store.History("bad id"); err == nil {
		t.Fatalf("bad history application id must be rejected")
	}
}
