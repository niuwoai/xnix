// Package execution implements a blocked-by-default compatibility execution
// transaction pipeline. It composes the Runtime subsystems — recipe trust,
// environment lifecycle, snapshot baseline, and Portal permission — into a
// single deterministic transaction and reports exactly why launch is blocked.
//
// Launch is never enabled here: the runtime write gate stays closed, so a
// transaction can be created, reviewed, and preflighted, but never executes.
// The pipeline creates no real process, grants no real permission, and mutates
// no host state.
package execution

import (
	"errors"
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/recipe"
)

// GateStatus is the outcome of one preflight gate.
type GateStatus string

const (
	GatePass    GateStatus = "pass"
	GatePending GateStatus = "pending"
	GateBlocked GateStatus = "blocked"
)

// Gate is one requirement the transaction evaluates before launch.
type Gate struct {
	ID     string     `json:"id"`
	Status GateStatus `json:"status"`
	Reason string     `json:"reason"`
}

// Decision is the user review outcome.
type Decision string

const (
	DecisionPending  Decision = "pending"
	DecisionApproved Decision = "approved"
	DecisionRejected Decision = "rejected"
)

// State is the transaction lifecycle state.
type State string

const (
	StateCreated   State = "created"
	StateReviewed  State = "reviewed"
	StatePreflight State = "preflight"
	StateBlocked   State = "blocked"
)

// Inputs are the subsystem facts a transaction is evaluated against. They are
// produced by the recipe, environment, snapshot, and portal subsystems.
type Inputs struct {
	Trust                   recipe.TrustState
	Environment             environment.Record
	SnapshotBaselinePresent bool
	PortalRequiredOps       []string
	PortalGrantedOps        []string
}

// Transaction is a non-persistent, blocked-by-default execution request.
type Transaction struct {
	RequestID         string   `json:"request_id"`
	ApplicationID     string   `json:"application_id"`
	Profile           string   `json:"profile"`
	State             State    `json:"state"`
	ReviewDecision    Decision `json:"review_decision"`
	Gates             []Gate   `json:"gates"`
	BlockedReasons    []string `json:"blocked_reasons"`
	LaunchAllowed     bool     `json:"launch_allowed"`
	LaunchEnabled     bool     `json:"launch_enabled"`
	BackendStarted    bool     `json:"backend_started"`
	PermissionGranted bool     `json:"permission_granted"`
	HostRootModified  bool     `json:"host_root_modified"`
	NetworkRequired   bool     `json:"network_required"`
	Summary           string   `json:"summary"`
}

// Pipeline creates deterministic execution transactions. Request ids use a
// monotonic sequence so behaviour is reproducible.
type Pipeline struct {
	sequence int
}

// NewPipeline returns an empty execution pipeline.
func NewPipeline() *Pipeline { return &Pipeline{} }

// Create opens a new transaction for an application id from subsystem inputs.
func (p *Pipeline) Create(applicationID string, inputs Inputs) (Transaction, error) {
	if applicationID == "" {
		return Transaction{}, errors.New("execution transaction requires an application id")
	}
	if inputs.Environment.ApplicationID != "" && inputs.Environment.ApplicationID != applicationID {
		return Transaction{}, fmt.Errorf("environment record %q does not match application %q", inputs.Environment.ApplicationID, applicationID)
	}
	p.sequence++
	return Transaction{
		RequestID:      fmt.Sprintf("xnix-exec-%s-%d", sanitize(applicationID), p.sequence),
		ApplicationID:  applicationID,
		Profile:        string(inputs.Environment.Profile),
		State:          StateCreated,
		ReviewDecision: DecisionPending,
	}, nil
}

// Review records the user review decision, moving the transaction to reviewed.
func (t Transaction) Review(decision Decision) (Transaction, error) {
	if decision != DecisionApproved && decision != DecisionRejected {
		return Transaction{}, errors.New("review decision must be approved or rejected")
	}
	if t.State != StateCreated {
		return Transaction{}, fmt.Errorf("only a created transaction can be reviewed, not %s", t.State)
	}
	t.ReviewDecision = decision
	t.State = StateReviewed
	return t, nil
}

// Preflight evaluates every gate against the inputs and finalizes the
// transaction. Launch always remains blocked: the runtime write gate is closed.
func (t Transaction) Preflight(inputs Inputs) Transaction {
	gates := []Gate{
		recipeTrustGate(inputs.Trust),
		reviewGate(t.ReviewDecision),
		environmentGate(inputs.Environment),
		snapshotGate(inputs.SnapshotBaselinePresent),
		portalGate(inputs.PortalRequiredOps, inputs.PortalGrantedOps),
		runtimeWriteGate(),
	}

	var blocked []string
	for _, gate := range gates {
		if gate.Status != GatePass {
			blocked = append(blocked, fmt.Sprintf("%s: %s", gate.ID, gate.Reason))
		}
	}
	sort.Strings(blocked)

	t.Gates = gates
	t.BlockedReasons = blocked
	t.State = StatePreflight
	// Launch is unconditionally disabled in this pipeline.
	t.LaunchAllowed = false
	t.LaunchEnabled = false
	t.BackendStarted = false
	t.PermissionGranted = false
	t.HostRootModified = false
	t.NetworkRequired = false
	if len(blocked) == 0 {
		// Even with every subsystem gate passing, the write gate keeps launch
		// blocked; this state is unreachable while runtimeWriteGate is closed,
		// but the summary is defined for completeness.
		t.Summary = "All subsystem gates pass, but launch remains disabled by the Runtime write gate."
	} else {
		t.State = StateBlocked
		t.Summary = fmt.Sprintf("Execution is blocked by %d gate(s); launch is disabled.", len(blocked))
	}
	return t
}

func recipeTrustGate(trust recipe.TrustState) Gate {
	switch {
	case !trust.Usable():
		return Gate{ID: "recipe-trust", Status: GateBlocked, Reason: "recipe failed trust verification"}
	case trust.ProductionTrusted:
		return Gate{ID: "recipe-trust", Status: GatePass, Reason: "recipe is production trusted"}
	default:
		return Gate{ID: "recipe-trust", Status: GatePending, Reason: "recipe is development-only, not production trusted"}
	}
}

func reviewGate(decision Decision) Gate {
	switch decision {
	case DecisionApproved:
		return Gate{ID: "user-review", Status: GatePass, Reason: "user approved the operation"}
	case DecisionRejected:
		return Gate{ID: "user-review", Status: GateBlocked, Reason: "user rejected the operation"}
	default:
		return Gate{ID: "user-review", Status: GatePending, Reason: "user review is pending"}
	}
}

func environmentGate(record environment.Record) Gate {
	if record.Ready() {
		return Gate{ID: "environment-ready", Status: GatePass, Reason: "compatibility environment is ready"}
	}
	if record.State == environment.StateRepairRequired {
		return Gate{ID: "environment-ready", Status: GateBlocked, Reason: "compatibility environment needs repair"}
	}
	return Gate{ID: "environment-ready", Status: GatePending, Reason: fmt.Sprintf("environment is %s, not ready", string(record.State))}
}

func snapshotGate(present bool) Gate {
	if present {
		return Gate{ID: "snapshot-baseline", Status: GatePass, Reason: "a restore point exists"}
	}
	return Gate{ID: "snapshot-baseline", Status: GatePending, Reason: "no restore point exists yet"}
}

func portalGate(required, granted []string) Gate {
	grantedSet := map[string]bool{}
	for _, op := range granted {
		grantedSet[op] = true
	}
	var missing []string
	for _, op := range required {
		if !grantedSet[op] {
			missing = append(missing, op)
		}
	}
	if len(missing) == 0 {
		return Gate{ID: "portal-permission", Status: GatePass, Reason: "all required desktop permissions are granted"}
	}
	sort.Strings(missing)
	return Gate{ID: "portal-permission", Status: GatePending, Reason: fmt.Sprintf("awaiting desktop permission for: %v", missing)}
}

func runtimeWriteGate() Gate {
	return Gate{ID: "runtime-write-gate", Status: GateBlocked, Reason: "launch is disabled until production Runtime ownership is ready"}
}

func sanitize(value string) string {
	out := make([]rune, 0, len(value))
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			out = append(out, r)
		} else {
			out = append(out, '-')
		}
	}
	return string(out)
}
