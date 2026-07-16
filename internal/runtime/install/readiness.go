// Package install joins the trust and artifact evidence that must all hold
// before a compatibility application may be installed: the recipe must be
// trusted, its artifacts must be staged with verified digests, and the Runtime
// owner must be ready. It never installs anything — install execution is a
// gated write that stays disabled — it only reports whether install is
// permitted and, if not, exactly why.
package install

import (
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/recipe"
)

// Environment is the trust environment an install is evaluated against.
type Environment string

const (
	// Development permits development-only recipes.
	Development Environment = "development"
	// Production requires a production-trusted recipe.
	Production Environment = "production"
)

func (e Environment) valid() bool { return e == Development || e == Production }

// GateStatus is the outcome of one install-readiness gate.
type GateStatus string

const (
	GatePass    GateStatus = "pass"
	GatePending GateStatus = "pending"
	GateBlocked GateStatus = "blocked"
)

// Gate is one install-readiness requirement.
type Gate struct {
	ID     string     `json:"id"`
	Status GateStatus `json:"status"`
	Reason string     `json:"reason"`
}

// Inputs are the subsystem facts an install readiness verdict joins.
type Inputs struct {
	Trust                   recipe.TrustState
	ArtifactsStaged         bool
	ArtifactDigestsVerified bool
	OwnerReady              bool
	Environment             Environment
}

// Readiness is the joined install-readiness verdict. It is KDE-safe: it exposes
// no host paths or backend details, and install is never enabled here.
type Readiness struct {
	Environment           Environment `json:"environment"`
	Ready                 bool        `json:"ready"`
	InstallEnabled        bool        `json:"install_enabled"`
	Gates                 []Gate      `json:"gates"`
	BlockedReasons        []string    `json:"blocked_reasons"`
	HostRootModified      bool        `json:"host_root_modified"`
	BackendDetailsExposed bool        `json:"backend_details_exposed"`
	Summary               string      `json:"summary"`
}

// Evaluate joins the inputs into an install-readiness verdict. Install
// execution remains disabled regardless of the verdict.
func Evaluate(inputs Inputs) (Readiness, error) {
	if !inputs.Environment.valid() {
		return Readiness{}, fmt.Errorf("install environment must be development or production, not %q", string(inputs.Environment))
	}

	gates := []Gate{
		recipeTrustGate(inputs.Trust, inputs.Environment),
		artifactDigestGate(inputs.ArtifactDigestsVerified),
		artifactStagingGate(inputs.ArtifactsStaged),
		ownerReadinessGate(inputs.OwnerReady),
	}

	var blocked []string
	ready := true
	for _, gate := range gates {
		if gate.Status != GatePass {
			ready = false
			blocked = append(blocked, fmt.Sprintf("%s: %s", gate.ID, gate.Reason))
		}
	}
	sort.Strings(blocked)

	summary := "Compatibility install prerequisites are satisfied; install remains a gated Runtime write."
	if !ready {
		summary = fmt.Sprintf("Compatibility install is not ready: %d prerequisite(s) not satisfied.", len(blocked))
	}

	return Readiness{
		Environment:           inputs.Environment,
		Ready:                 ready,
		InstallEnabled:        false,
		Gates:                 gates,
		BlockedReasons:        blocked,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		Summary:               summary,
	}, nil
}

func recipeTrustGate(trust recipe.TrustState, env Environment) Gate {
	switch {
	case !trust.Usable():
		return Gate{ID: "recipe-trust", Status: GateBlocked, Reason: "recipe failed trust verification"}
	case env == Production && !trust.ProductionTrusted:
		return Gate{ID: "recipe-trust", Status: GatePending, Reason: "recipe is not production trusted"}
	default:
		return Gate{ID: "recipe-trust", Status: GatePass, Reason: "recipe trust is sufficient for this environment"}
	}
}

func artifactDigestGate(verified bool) Gate {
	if verified {
		return Gate{ID: "artifact-digests", Status: GatePass, Reason: "artifact digests are verified"}
	}
	return Gate{ID: "artifact-digests", Status: GateBlocked, Reason: "artifact digests are not verified"}
}

func artifactStagingGate(staged bool) Gate {
	if staged {
		return Gate{ID: "artifact-staging", Status: GatePass, Reason: "artifacts are staged under the Runtime cache root"}
	}
	return Gate{ID: "artifact-staging", Status: GatePending, Reason: "artifacts are not staged yet"}
}

func ownerReadinessGate(ready bool) Gate {
	if ready {
		return Gate{ID: "owner-readiness", Status: GatePass, Reason: "the Runtime owner is ready"}
	}
	return Gate{ID: "owner-readiness", Status: GatePending, Reason: "the Runtime owner is not ready"}
}
