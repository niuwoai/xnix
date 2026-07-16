package diagnostics

import (
	"errors"
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/appid"
)

// ErrAIProviderDisabled is returned by the disabled default provider. AI
// diagnostics never call a real provider unless one is explicitly supplied.
var ErrAIProviderDisabled = errors.New("AI diagnostics provider is disabled")

// DiagnosticInput is the privacy-filtered context handed to an AI provider. It
// is built from a test result and carries only safe signal summaries — never
// file contents, host paths, backend logs, commands, or secrets.
type DiagnosticInput struct {
	ApplicationID        string   `json:"application_id"`
	TestType             string   `json:"test_type"`
	Overall              Outcome  `json:"overall"`
	SignalSummaries      []string `json:"signal_summaries"`
	NetworkRequired      bool     `json:"network_required"`
	FileContentsIncluded bool     `json:"file_contents_included"`
}

// BuildDiagnosticInput constructs a privacy-filtered AI input from a result.
// It re-validates every summary so unsafe content cannot reach a provider.
func BuildDiagnosticInput(result TestResult) (DiagnosticInput, error) {
	if !appid.Valid(result.ApplicationID) {
		return DiagnosticInput{}, errors.New("application id must be a reverse-DNS identifier")
	}
	summaries := make([]string, 0, len(result.Signals))
	for _, s := range result.Signals {
		if err := validateSafe("signal summary", s.Summary); err != nil {
			return DiagnosticInput{}, err
		}
		summaries = append(summaries, s.Summary)
	}
	return DiagnosticInput{
		ApplicationID:        result.ApplicationID,
		TestType:             result.TestType,
		Overall:              result.Overall,
		SignalSummaries:      summaries,
		NetworkRequired:      false,
		FileContentsIncluded: false,
	}, nil
}

// Recommendation is an AI-produced diagnostic recommendation. Its advice is
// re-validated to remain privacy-safe.
type Recommendation struct {
	ApplicationID string `json:"application_id"`
	Confidence    string `json:"confidence"`
	Advice        string `json:"advice"`
	SuggestIssue  string `json:"suggest_issue"`
	Provider      string `json:"provider"`
}

// AIProvider is the replaceable AI boundary. The default provider is disabled;
// a fake provider serves tests and a real provider can be supplied later.
type AIProvider interface {
	Name() string
	Recommend(input DiagnosticInput) (Recommendation, error)
}

// DisabledProvider is the production default: it never calls out.
type DisabledProvider struct{}

func (DisabledProvider) Name() string { return "disabled" }

func (DisabledProvider) Recommend(DiagnosticInput) (Recommendation, error) {
	return Recommendation{}, ErrAIProviderDisabled
}

// FakeProvider is a deterministic in-process provider for tests. It makes no
// network calls and echoes only safe, derived advice.
type FakeProvider struct{}

func (FakeProvider) Name() string { return "fake" }

// Recommend returns deterministic advice derived from the input outcome.
func (FakeProvider) Recommend(input DiagnosticInput) (Recommendation, error) {
	if len(input.SignalSummaries) == 0 {
		return Recommendation{}, errors.New("diagnostic input has no signals")
	}
	advice := "Signals look healthy; no repair is recommended."
	issue := ""
	confidence := "low"
	switch input.Overall {
	case OutcomeFail:
		advice = "Prepare a Runtime-owned engine binding repair and take a restore point first."
		issue = "engine-binding-pending"
		confidence = "medium"
	case OutcomeBlocked:
		advice = "Ask the user to approve the pending desktop permission before retrying."
		issue = "portal-approval-required"
		confidence = "medium"
	}
	rec := Recommendation{
		ApplicationID: input.ApplicationID,
		Confidence:    confidence,
		Advice:        advice,
		SuggestIssue:  issue,
		Provider:      "fake",
	}
	if err := validateSafe("ai advice", rec.Advice); err != nil {
		return Recommendation{}, err
	}
	return rec, nil
}

// Diagnose runs a provider over a result, defaulting to the disabled provider.
func Diagnose(provider AIProvider, result TestResult) (Recommendation, error) {
	if provider == nil {
		provider = DisabledProvider{}
	}
	input, err := BuildDiagnosticInput(result)
	if err != nil {
		return Recommendation{}, err
	}
	return provider.Recommend(input)
}

// ApprovalDecision is the user's decision on an AI-suggested repair.
type ApprovalDecision string

const (
	ApprovalPending  ApprovalDecision = "pending"
	ApprovalApproved ApprovalDecision = "approved"
	ApprovalRejected ApprovalDecision = "rejected"
)

// ApprovalGate connects an AI recommendation to a repair plan. It is review-
// first and snapshot-gated: repair is authorized only when the user approved
// and a snapshot baseline exists. It never executes a repair itself.
type ApprovalGate struct {
	ApplicationID    string           `json:"application_id"`
	Issue            string           `json:"issue"`
	Decision         ApprovalDecision `json:"decision"`
	SnapshotPresent  bool             `json:"snapshot_present"`
	RepairAuthorized bool             `json:"repair_authorized"`
	RepairExecuted   bool             `json:"repair_executed"`
	BlockedReasons   []string         `json:"blocked_reasons"`
}

// Evaluate builds an approval gate from a recommendation, a user decision, and
// snapshot presence. Repair is authorized (but never executed) only when the
// user approved and a snapshot baseline exists.
func Evaluate(rec Recommendation, decision ApprovalDecision, snapshotPresent bool) (ApprovalGate, error) {
	if rec.SuggestIssue == "" {
		return ApprovalGate{}, fmt.Errorf("recommendation has no repair issue to gate")
	}
	gate := ApprovalGate{
		ApplicationID:   rec.ApplicationID,
		Issue:           rec.SuggestIssue,
		Decision:        decision,
		SnapshotPresent: snapshotPresent,
		RepairExecuted:  false,
	}
	var blocked []string
	if decision != ApprovalApproved {
		blocked = append(blocked, "user-review: repair not approved")
	}
	if !snapshotPresent {
		blocked = append(blocked, "snapshot-baseline: no restore point exists")
	}
	sort.Strings(blocked)
	gate.BlockedReasons = blocked
	gate.RepairAuthorized = len(blocked) == 0
	return gate, nil
}
