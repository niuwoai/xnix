package diagnostics

import (
	"errors"
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/appid"
	"xnix.local/xnix/internal/runtime/record"
	"xnix.local/xnix/internal/runtime/rootfs"
)

// supportedTestTypes enumerates the fixture-runnable test types.
var supportedTestTypes = map[string]bool{"preflight": true, "smoke": true, "repair-readiness": true}

// Fixture is a deterministic, backend-free description of a test run's signals.
// It is the only input to the runner: no real backend is ever started.
type Fixture struct {
	TestType string   `json:"test_type"`
	Signals  []Signal `json:"signals"`
}

// TestResult is the persisted outcome of a fixture run. It is safe for KDE and
// Compatibility Center summaries.
type TestResult struct {
	ApplicationID    string   `json:"application_id"`
	TestType         string   `json:"test_type"`
	Overall          Outcome  `json:"overall"`
	Signals          []Signal `json:"signals"`
	FailingIDs       []string `json:"failing_ids"`
	BackendStarted   bool     `json:"backend_started"`
	HostRootModified bool     `json:"host_root_modified"`
	Summary          string   `json:"summary"`
}

// Run evaluates a fixture into a TestResult without starting any backend.
func Run(applicationID string, fixture Fixture) (TestResult, error) {
	if !appid.Valid(applicationID) {
		return TestResult{}, errors.New("application id must be a reverse-DNS identifier")
	}
	if !supportedTestTypes[fixture.TestType] {
		return TestResult{}, fmt.Errorf("unsupported test type %q", fixture.TestType)
	}
	if len(fixture.Signals) == 0 {
		return TestResult{}, errors.New("fixture must declare at least one signal")
	}
	for _, s := range fixture.Signals {
		if err := s.validate(); err != nil {
			return TestResult{}, err
		}
	}

	overall := worstOutcome(fixture.Signals)
	var failing []string
	for _, s := range fixture.Signals {
		if s.Outcome == OutcomeFail || s.Outcome == OutcomeBlocked {
			failing = append(failing, s.ID)
		}
	}
	sort.Strings(failing)

	summary := "Compatibility test signals are healthy."
	switch overall {
	case OutcomeFail:
		summary = "Compatibility test found failing signals; review recommended."
	case OutcomeBlocked:
		summary = "Compatibility test is blocked pending user or Runtime action."
	case OutcomePending:
		summary = "Compatibility test is waiting for Runtime-controlled preflight."
	}

	return TestResult{
		ApplicationID:    applicationID,
		TestType:         fixture.TestType,
		Overall:          overall,
		Signals:          append([]Signal(nil), fixture.Signals...),
		FailingIDs:       failing,
		BackendStarted:   false,
		HostRootModified: false,
		Summary:          summary,
	}, nil
}

// ResultStore persists test results under a controlled state root.
type ResultStore struct {
	root *rootfs.Root
}

// NewResultStore opens a result store under the state root.
func NewResultStore(stateRoot string) (*ResultStore, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return nil, err
	}
	return &ResultStore{root: root}, nil
}

func resultKey(applicationID, testType string) string {
	return sanitizeKey(applicationID) + "__" + sanitizeKey(testType)
}

// resultRel is the state-root-relative path of a persisted test result.
func resultRel(applicationID, testType string) string {
	return "test-results/" + resultKey(applicationID, testType) + ".json"
}

func sanitizeKey(value string) string {
	out := make([]rune, 0, len(value))
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-':
			out = append(out, r)
		default:
			out = append(out, '_')
		}
	}
	return string(out)
}

// Save atomically persists a test result (latest wins per application/test-type).
func (s *ResultStore) Save(result TestResult) error {
	_, err := record.Save(s.root, resultRel(result.ApplicationID, result.TestType), result)
	return err
}

// Load returns the persisted result for an application/test-type.
func (s *ResultStore) Load(applicationID, testType string) (TestResult, error) {
	var result TestResult
	if err := record.Load(s.root, resultRel(applicationID, testType), &result); err != nil {
		return TestResult{}, err
	}
	return result, nil
}

// RepairRecommendation is a safe, review-first repair suggestion derived from
// failing test signals. It never executes a repair.
type RepairRecommendation struct {
	ApplicationID    string   `json:"application_id"`
	Issue            string   `json:"issue"`
	FromSignals      []string `json:"from_signals"`
	SnapshotRequired bool     `json:"snapshot_required"`
	AutoApplyAllowed bool     `json:"auto_apply_allowed"`
	Summary          string   `json:"summary"`
}

// Recommend derives a repair recommendation from a test result. A healthy
// result yields no recommendation (ok=false).
func Recommend(result TestResult) (RepairRecommendation, bool) {
	if result.Overall == OutcomePass || len(result.FailingIDs) == 0 {
		return RepairRecommendation{}, false
	}
	issue := "engine-binding-pending"
	if containsSignal(result.Signals, OutcomeBlocked) && !containsSignal(result.Signals, OutcomeFail) {
		issue = "portal-approval-required"
	}
	return RepairRecommendation{
		ApplicationID:    result.ApplicationID,
		Issue:            issue,
		FromSignals:      append([]string(nil), result.FailingIDs...),
		SnapshotRequired: true,
		AutoApplyAllowed: false,
		Summary:          "Runtime recommends a review-first, snapshot-gated repair.",
	}, true
}

func containsSignal(signals []Signal, outcome Outcome) bool {
	for _, s := range signals {
		if s.Outcome == outcome {
			return true
		}
	}
	return false
}
