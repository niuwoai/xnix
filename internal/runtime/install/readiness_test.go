package install

import (
	"testing"

	"xnix.local/xnix/internal/runtime/recipe"
)

func trusted(production bool) recipe.TrustState {
	return recipe.TrustState{DigestVerified: true, ProductionTrusted: production, DevelopmentOnly: !production}
}

func gateStatus(r Readiness, id string) GateStatus {
	for _, g := range r.Gates {
		if g.ID == id {
			return g.Status
		}
	}
	return ""
}

func TestReadyWhenAllPrerequisitesHold(t *testing.T) {
	r, err := Evaluate(Inputs{
		Trust:                   trusted(true),
		ArtifactsStaged:         true,
		ArtifactDigestsVerified: true,
		OwnerReady:              true,
		Environment:             Production,
	})
	if err != nil {
		t.Fatalf("Evaluate: %v", err)
	}
	if !r.Ready || r.InstallEnabled || len(r.BlockedReasons) != 0 {
		t.Fatalf("expected ready with install still disabled: %#v", r)
	}
	if r.HostRootModified || r.BackendDetailsExposed {
		t.Fatalf("readiness must not modify host or expose backend details: %#v", r)
	}
	for _, id := range []string{"recipe-trust", "artifact-digests", "artifact-staging", "owner-readiness"} {
		if gateStatus(r, id) != GatePass {
			t.Fatalf("gate %s should pass: %#v", id, r.Gates)
		}
	}
}

func TestDevelopmentTrustPassesInDevelopmentButNotProduction(t *testing.T) {
	dev := Inputs{Trust: trusted(false), ArtifactsStaged: true, ArtifactDigestsVerified: true, OwnerReady: true, Environment: Development}
	r, _ := Evaluate(dev)
	if !r.Ready || gateStatus(r, "recipe-trust") != GatePass {
		t.Fatalf("development-only recipe should be install-ready in development: %#v", r)
	}

	prod := dev
	prod.Environment = Production
	pr, _ := Evaluate(prod)
	if pr.Ready || gateStatus(pr, "recipe-trust") != GatePending {
		t.Fatalf("development-only recipe must not be production install-ready: %#v", pr)
	}
}

func TestFailedTrustBlocks(t *testing.T) {
	r, _ := Evaluate(Inputs{
		Trust:                   recipe.TrustState{FailedClosed: true},
		ArtifactsStaged:         true,
		ArtifactDigestsVerified: true,
		OwnerReady:              true,
		Environment:             Development,
	})
	if r.Ready || gateStatus(r, "recipe-trust") != GateBlocked {
		t.Fatalf("fail-closed trust must block install readiness: %#v", r)
	}
}

func TestMissingStagingAndDigestsAndOwnerBlock(t *testing.T) {
	r, _ := Evaluate(Inputs{
		Trust:                   trusted(true),
		ArtifactsStaged:         false,
		ArtifactDigestsVerified: false,
		OwnerReady:              false,
		Environment:             Production,
	})
	if r.Ready {
		t.Fatalf("missing prerequisites must not be ready")
	}
	if gateStatus(r, "artifact-digests") != GateBlocked ||
		gateStatus(r, "artifact-staging") != GatePending ||
		gateStatus(r, "owner-readiness") != GatePending {
		t.Fatalf("unexpected gate statuses: %#v", r.Gates)
	}
	if len(r.BlockedReasons) != 3 {
		t.Fatalf("expected 3 blocked reasons: %#v", r.BlockedReasons)
	}
}

func TestInvalidEnvironmentRejected(t *testing.T) {
	if _, err := Evaluate(Inputs{Trust: trusted(true), Environment: Environment("staging")}); err == nil {
		t.Fatalf("invalid environment must be rejected")
	}
}
