package execution

import (
	"testing"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/recipe"
)

const app = "org.example.ledger"

func readyEnvironment(t *testing.T) environment.Record {
	t.Helper()
	lc, err := environment.New(t.TempDir())
	if err != nil {
		t.Fatalf("environment.New: %v", err)
	}
	lc.Plan(app, environment.ProfileLocal)
	lc.Stage(app, environment.ProfileLocal)
	for _, gate := range environment.RequiredGates() {
		lc.SatisfyGate(app, environment.ProfileLocal, gate)
	}
	rec, err := lc.MarkReady(app, environment.ProfileLocal)
	if err != nil {
		t.Fatalf("MarkReady: %v", err)
	}
	return rec
}

func gateStatus(tx Transaction, id string) GateStatus {
	for _, g := range tx.Gates {
		if g.ID == id {
			return g.Status
		}
	}
	return ""
}

func TestBlockedByWriteGateEvenWhenEverythingElsePasses(t *testing.T) {
	inputs := Inputs{
		Trust:                   recipe.TrustState{DigestVerified: true, ProductionTrusted: true},
		Environment:             readyEnvironment(t),
		SnapshotBaselinePresent: true,
		PortalRequiredOps:       []string{"file-open"},
		PortalGrantedOps:        []string{"file-open"},
	}

	p := NewPipeline()
	tx, err := p.Create(app, inputs)
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if tx.RequestID != "xnix-exec-org-example-ledger-1" || tx.State != StateCreated {
		t.Fatalf("unexpected created transaction: %#v", tx)
	}

	tx, err = tx.Review(DecisionApproved)
	if err != nil {
		t.Fatalf("Review: %v", err)
	}
	if tx.State != StateReviewed {
		t.Fatalf("review did not advance state: %#v", tx)
	}

	tx = tx.Preflight(inputs)
	// Every subsystem gate passes, but the write gate blocks launch.
	if gateStatus(tx, "recipe-trust") != GatePass ||
		gateStatus(tx, "user-review") != GatePass ||
		gateStatus(tx, "environment-ready") != GatePass ||
		gateStatus(tx, "snapshot-baseline") != GatePass ||
		gateStatus(tx, "portal-permission") != GatePass {
		t.Fatalf("subsystem gates should pass: %#v", tx.Gates)
	}
	if gateStatus(tx, "runtime-write-gate") != GateBlocked {
		t.Fatalf("write gate must be blocked: %#v", tx.Gates)
	}
	if tx.LaunchAllowed || tx.LaunchEnabled || tx.BackendStarted || tx.PermissionGranted || tx.HostRootModified || tx.NetworkRequired {
		t.Fatalf("launch and side effects must stay disabled: %#v", tx)
	}
	if tx.State != StateBlocked || len(tx.BlockedReasons) != 1 {
		t.Fatalf("transaction must be blocked with exactly the write-gate reason: %#v", tx)
	}
}

func TestPendingGatesReportedForDevelopmentAndUnready(t *testing.T) {
	// Development-only trust, environment not ready, no snapshot, portal ungranted.
	inputs := Inputs{
		Trust:                   recipe.TrustState{DigestVerified: true, DevelopmentOnly: true},
		Environment:             environment.Record{ApplicationID: app, Profile: environment.ProfileLocal, State: environment.StatePlanned},
		SnapshotBaselinePresent: false,
		PortalRequiredOps:       []string{"file-open", "print"},
		PortalGrantedOps:        []string{"file-open"},
	}
	p := NewPipeline()
	tx, _ := p.Create(app, inputs)
	tx, _ = tx.Review(DecisionApproved)
	tx = tx.Preflight(inputs)

	if gateStatus(tx, "recipe-trust") != GatePending {
		t.Fatalf("dev-only trust must be pending: %#v", tx.Gates)
	}
	if gateStatus(tx, "environment-ready") != GatePending {
		t.Fatalf("unready environment must be pending: %#v", tx.Gates)
	}
	if gateStatus(tx, "snapshot-baseline") != GatePending {
		t.Fatalf("absent snapshot must be pending")
	}
	if gateStatus(tx, "portal-permission") != GatePending {
		t.Fatalf("ungranted portal op must be pending")
	}
	if tx.State != StateBlocked {
		t.Fatalf("transaction must be blocked: %#v", tx)
	}
}

func TestFailClosedTrustBlocks(t *testing.T) {
	inputs := Inputs{
		Trust:       recipe.TrustState{FailedClosed: true, Reason: "digest mismatch"},
		Environment: readyEnvironment(t),
	}
	p := NewPipeline()
	tx, _ := p.Create(app, inputs)
	tx, _ = tx.Review(DecisionApproved)
	tx = tx.Preflight(inputs)
	if gateStatus(tx, "recipe-trust") != GateBlocked {
		t.Fatalf("fail-closed trust must block: %#v", tx.Gates)
	}
}

func TestRejectedReviewBlocks(t *testing.T) {
	inputs := Inputs{Trust: recipe.TrustState{DigestVerified: true, ProductionTrusted: true}, Environment: readyEnvironment(t)}
	p := NewPipeline()
	tx, _ := p.Create(app, inputs)
	tx, err := tx.Review(DecisionRejected)
	if err != nil {
		t.Fatalf("Review: %v", err)
	}
	tx = tx.Preflight(inputs)
	if gateStatus(tx, "user-review") != GateBlocked {
		t.Fatalf("rejected review must block: %#v", tx.Gates)
	}
}

func TestReviewGuards(t *testing.T) {
	p := NewPipeline()
	tx, _ := p.Create(app, Inputs{})
	if _, err := tx.Review(DecisionPending); err == nil {
		t.Fatalf("pending decision must be rejected")
	}
	reviewed, _ := tx.Review(DecisionApproved)
	if _, err := reviewed.Review(DecisionApproved); err == nil {
		t.Fatalf("double review must be rejected")
	}
}

func TestCreateGuards(t *testing.T) {
	p := NewPipeline()
	if _, err := p.Create("", Inputs{}); err == nil {
		t.Fatalf("empty application id must be rejected")
	}
	mismatch := Inputs{Environment: environment.Record{ApplicationID: "org.other.app", Profile: environment.ProfileLocal}}
	if _, err := p.Create(app, mismatch); err == nil {
		t.Fatalf("environment/application mismatch must be rejected")
	}
}

func TestRequestIDsAreDeterministicSequence(t *testing.T) {
	p := NewPipeline()
	a, _ := p.Create(app, Inputs{})
	b, _ := p.Create(app, Inputs{})
	if a.RequestID != "xnix-exec-org-example-ledger-1" || b.RequestID != "xnix-exec-org-example-ledger-2" {
		t.Fatalf("request ids must be a deterministic sequence: %q %q", a.RequestID, b.RequestID)
	}
}
