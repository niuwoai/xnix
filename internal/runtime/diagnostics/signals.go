// Package diagnostics implements a safe, fixture-driven compatibility test,
// repair, and AI-diagnostics pipeline. It runs no real backend, persists test
// results under a caller-controlled state root, derives repair recommendations
// from safe test signals, and routes optional AI diagnostics through a provider
// interface that is disabled by default (a fake provider serves tests).
//
// Every value that leaves this package is privacy-filtered: file contents, host
// paths, backend logs, raw commands, secrets, and tokens are rejected so they
// never reach user-facing or AI-bound output.
package diagnostics

import (
	"fmt"

	"xnix.local/xnix/internal/runtime/safety"
)

// Outcome is the result of a single test signal or an overall test run.
type Outcome string

const (
	OutcomePass    Outcome = "pass"
	OutcomePending Outcome = "pending"
	OutcomeBlocked Outcome = "blocked"
	OutcomeFail    Outcome = "fail"
)

func (o Outcome) valid() bool {
	switch o {
	case OutcomePass, OutcomePending, OutcomeBlocked, OutcomeFail:
		return true
	default:
		return false
	}
}

// Signal is one safe, user-presentable observation from a test run. It carries
// no host paths, commands, or file contents.
type Signal struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Outcome  Outcome `json:"outcome"`
	Summary  string  `json:"summary"`
}

// validateSafe rejects text that leaks host paths, backend terms, commands,
// secrets, or tokens. It is the privacy boundary for all diagnostics output and
// delegates to the canonical safety validator, which is a strict superset of the
// term list this package previously maintained on its own.
func validateSafe(label, text string) error {
	return safety.ValidateLine(label, text)
}

func (s Signal) validate() error {
	if s.ID == "" {
		return fmt.Errorf("signal id must not be empty")
	}
	if !s.Outcome.valid() {
		return fmt.Errorf("signal %q has invalid outcome %q", s.ID, string(s.Outcome))
	}
	for label, value := range map[string]string{
		"signal id":       s.ID,
		"signal category": s.Category,
		"signal summary":  s.Summary,
	} {
		if err := validateSafe(label, value); err != nil {
			return err
		}
	}
	return nil
}

// worstOutcome folds signal outcomes into an overall run outcome.
func worstOutcome(signals []Signal) Outcome {
	overall := OutcomePass
	rank := map[Outcome]int{OutcomePass: 0, OutcomePending: 1, OutcomeBlocked: 2, OutcomeFail: 3}
	for _, s := range signals {
		if rank[s.Outcome] > rank[overall] {
			overall = s.Outcome
		}
	}
	return overall
}
