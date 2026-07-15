package diagnostics

import (
	"errors"
	"testing"
)

func TestDisabledProviderIsDefault(t *testing.T) {
	res, _ := Run(app, failFixture())
	// Diagnose with nil provider must use the disabled default.
	if _, err := Diagnose(nil, res); !errors.Is(err, ErrAIProviderDisabled) {
		t.Fatalf("nil provider must be disabled, got %v", err)
	}
	if _, err := Diagnose(DisabledProvider{}, res); !errors.Is(err, ErrAIProviderDisabled) {
		t.Fatalf("disabled provider must fail closed, got %v", err)
	}
}

func TestFakeProviderIsDeterministicAndSafe(t *testing.T) {
	res, _ := Run(app, failFixture())
	rec, err := Diagnose(FakeProvider{}, res)
	if err != nil {
		t.Fatalf("Diagnose: %v", err)
	}
	if rec.Provider != "fake" || rec.SuggestIssue != "engine-binding-pending" || rec.Confidence != "medium" {
		t.Fatalf("fake recommendation wrong: %#v", rec)
	}
	// Determinism: same input, same advice.
	rec2, _ := Diagnose(FakeProvider{}, res)
	if rec2.Advice != rec.Advice {
		t.Fatalf("fake provider must be deterministic")
	}
}

func TestBuildDiagnosticInputRejectsUnsafeSummaries(t *testing.T) {
	res := TestResult{
		ApplicationID: app,
		TestType:      "preflight",
		Overall:       OutcomeFail,
		Signals:       []Signal{{ID: "leak", Outcome: OutcomeFail, Summary: "token=sk-abcdef123456 leaked"}},
	}
	if _, err := BuildDiagnosticInput(res); err == nil {
		t.Fatalf("secret-shaped summary must be rejected before reaching a provider")
	}
}

func TestDiagnosticInputCarriesNoFileContentsOrNetwork(t *testing.T) {
	res, _ := Run(app, failFixture())
	input, err := BuildDiagnosticInput(res)
	if err != nil {
		t.Fatalf("BuildDiagnosticInput: %v", err)
	}
	if input.FileContentsIncluded || input.NetworkRequired {
		t.Fatalf("diagnostic input must carry no file contents or network requirement: %#v", input)
	}
	if len(input.SignalSummaries) != 2 {
		t.Fatalf("unexpected signal summaries: %#v", input.SignalSummaries)
	}
}

func TestApprovalGateIsReviewFirstAndSnapshotGated(t *testing.T) {
	res, _ := Run(app, failFixture())
	rec, _ := Diagnose(FakeProvider{}, res)

	// Approved but no snapshot -> blocked.
	g1, err := Evaluate(rec, ApprovalApproved, false)
	if err != nil {
		t.Fatalf("Evaluate: %v", err)
	}
	if g1.RepairAuthorized || len(g1.BlockedReasons) != 1 {
		t.Fatalf("approved-without-snapshot must be blocked: %#v", g1)
	}

	// Snapshot present but not approved -> blocked.
	g2, _ := Evaluate(rec, ApprovalPending, true)
	if g2.RepairAuthorized {
		t.Fatalf("unapproved repair must be blocked: %#v", g2)
	}

	// Approved and snapshot present -> authorized, but never executed.
	g3, _ := Evaluate(rec, ApprovalApproved, true)
	if !g3.RepairAuthorized || g3.RepairExecuted || len(g3.BlockedReasons) != 0 {
		t.Fatalf("approved+snapshot must authorize but not execute: %#v", g3)
	}

	// A recommendation with no issue cannot be gated.
	if _, err := Evaluate(Recommendation{ApplicationID: app}, ApprovalApproved, true); err == nil {
		t.Fatalf("recommendation without an issue must not be gateable")
	}
}
