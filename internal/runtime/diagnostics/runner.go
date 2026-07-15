package diagnostics

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

// applicationIDPattern matches reverse-DNS application identifiers.
var applicationIDPattern = regexp.MustCompile(`^[A-Za-z0-9]+(\.[A-Za-z0-9-]+)+$`)

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
	if !applicationIDPattern.MatchString(applicationID) {
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
	dir string
}

// NewResultStore opens (creating if needed) a result store under the state root.
func NewResultStore(stateRoot string) (*ResultStore, error) {
	if stateRoot == "" {
		return nil, errors.New("result store requires a state root")
	}
	abs, err := filepath.Abs(stateRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve state root: %w", err)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("state root must exist: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("state root %q is not a directory", abs)
	}
	dir := filepath.Join(abs, "test-results")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("initialize result store: %w", err)
	}
	return &ResultStore{dir: dir}, nil
}

func resultKey(applicationID, testType string) string {
	return sanitizeKey(applicationID) + "__" + sanitizeKey(testType)
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

// Save persists a test result (latest wins per application/test-type).
func (s *ResultStore) Save(result TestResult) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(s.dir, resultKey(result.ApplicationID, result.TestType)+".json"), data, 0o600)
}

// Load returns the persisted result for an application/test-type.
func (s *ResultStore) Load(applicationID, testType string) (TestResult, error) {
	data, err := os.ReadFile(filepath.Join(s.dir, resultKey(applicationID, testType)+".json"))
	if err != nil {
		return TestResult{}, fmt.Errorf("no persisted result for %s/%s: %w", applicationID, testType, err)
	}
	var result TestResult
	if err := json.Unmarshal(data, &result); err != nil {
		return TestResult{}, fmt.Errorf("persisted result is corrupt: %w", err)
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
