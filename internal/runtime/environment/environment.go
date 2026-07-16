// Package environment implements a Runtime-owned lifecycle state machine for
// compatibility environments. It tracks each application/profile environment
// through missing → planned → staged → ready, with repair-required and blocked
// branches, and persists state in a caller-provided (test-controlled) state
// root.
//
// The lifecycle never starts a real backend, never spawns a process, and never
// exposes backend commands or host paths. It only records state and surfaces
// repair hints; executing repairs and launching backends are out of scope
// (they belong to later packages).
package environment

import (
	"errors"
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/appid"
)

// State is a lifecycle state of a compatibility environment.
type State string

const (
	// StateMissing means no environment has been planned yet.
	StateMissing State = "missing"
	// StatePlanned means a profile has been chosen but nothing staged.
	StatePlanned State = "planned"
	// StateStaged means Runtime artifacts/metadata are staged, pending gates.
	StateStaged State = "staged"
	// StateReady means every required gate is satisfied; the environment is
	// ready for a launch decision (launch itself remains disabled).
	StateReady State = "ready"
	// StateRepairRequired means a problem was detected; repair hints apply.
	StateRepairRequired State = "repair-required"
	// StateBlocked means the environment is held closed (policy or safety).
	StateBlocked State = "blocked"
	// StateRetired means the environment was decommissioned. It holds no gates
	// and can be reactivated by planning it again.
	StateRetired State = "retired"
)

func (s State) valid() bool {
	switch s {
	case StateMissing, StatePlanned, StateStaged, StateReady, StateRepairRequired, StateBlocked, StateRetired:
		return true
	default:
		return false
	}
}

// Profile is the compatibility profile kind for an environment.
type Profile string

const (
	ProfileLocal    Profile = "local-compatibility"
	ProfileIsolated Profile = "isolated-compatibility"
)

func (p Profile) valid() bool {
	return p == ProfileLocal || p == ProfileIsolated
}

// requiredGates are the preflight gates an environment must satisfy before it
// can become ready. They are Runtime-owned checks, never backend commands.
var requiredGates = []string{
	"recipe-trust",
	"backend-binding",
	"portal-policy-review",
	"snapshot-baseline",
}

// RequiredGates returns the preflight gates required for readiness.
func RequiredGates() []string {
	return append([]string(nil), requiredGates...)
}

func isRequiredGate(gate string) bool {
	for _, g := range requiredGates {
		if g == gate {
			return true
		}
	}
	return false
}

// Record is the persisted lifecycle state of one application/profile
// environment. It carries no backend commands and no host paths.
type Record struct {
	ApplicationID       string   `json:"application_id"`
	Profile             Profile  `json:"profile"`
	State               State    `json:"state"`
	SatisfiedGates      []string `json:"satisfied_gates"`
	RepairHints         []string `json:"repair_hints"`
	BlockReason         string   `json:"block_reason,omitempty"`
	LaunchEnabled       bool     `json:"launch_enabled"`
	BackendStarted      bool     `json:"backend_started"`
	HostRootModified    bool     `json:"host_root_modified"`
	BackendDetailsShown bool     `json:"backend_details_shown"`
}

// Ready reports whether the environment is in the ready state.
func (r Record) Ready() bool { return r.State == StateReady }

// PendingGates returns the required gates that are not yet satisfied, sorted.
func (r Record) PendingGates() []string {
	satisfied := map[string]bool{}
	for _, g := range r.SatisfiedGates {
		satisfied[g] = true
	}
	var pending []string
	for _, g := range requiredGates {
		if !satisfied[g] {
			pending = append(pending, g)
		}
	}
	sort.Strings(pending)
	return pending
}

func newRecord(applicationID string, profile Profile) (Record, error) {
	if !appid.Valid(applicationID) {
		return Record{}, errors.New("application id must be a reverse-DNS identifier")
	}
	if !profile.valid() {
		return Record{}, fmt.Errorf("unsupported profile %q", string(profile))
	}
	return Record{
		ApplicationID:  applicationID,
		Profile:        profile,
		State:          StateMissing,
		SatisfiedGates: []string{},
		RepairHints:    []string{},
	}, nil
}

// allowedTransitions maps each state to the states it may move to.
var allowedTransitions = map[State][]State{
	StateMissing:        {StatePlanned, StateBlocked},
	StatePlanned:        {StateStaged, StateRepairRequired, StateBlocked, StateRetired},
	StateStaged:         {StateReady, StateRepairRequired, StateBlocked, StateRetired},
	StateReady:          {StateRepairRequired, StateBlocked, StateRetired},
	StateRepairRequired: {StateStaged, StateBlocked, StateRetired},
	StateBlocked:        {StatePlanned, StateRetired},
	StateRetired:        {StatePlanned},
}

func canTransition(from, to State) bool {
	for _, allowed := range allowedTransitions[from] {
		if allowed == to {
			return true
		}
	}
	return false
}
