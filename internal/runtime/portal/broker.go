package portal

import (
	"errors"
	"fmt"
)

// Outcome is the result reported back for a pending portal request.
type Outcome string

const (
	// OutcomeGranted means the user approved the request.
	OutcomeGranted Outcome = "granted"
	// OutcomeDenied means the user rejected the request.
	OutcomeDenied Outcome = "denied"
	// OutcomeCancelled means the request was withdrawn before a response.
	OutcomeCancelled Outcome = "cancelled"
	// OutcomeFailed means the Portal transport failed; recoverable.
	OutcomeFailed Outcome = "failed"
	// OutcomeExpired means the request timed out before a user response.
	OutcomeExpired Outcome = "expired"
)

// Broker owns portal request objects and their completion state. A real Portal
// transport can implement this interface later; FakeBroker serves tests.
type Broker interface {
	// CreateRequest explicitly creates a new portal request object. Creation
	// is never implicit — a caller must ask for it.
	CreateRequest(spec RequestSpec) (*Request, error)
	// Resolve records the Portal response for a pending request.
	Resolve(handleToken string, outcome Outcome) (*Request, error)
	// Complete marks a granted request as consumed by the caller.
	Complete(handleToken string) (*Request, error)
	// Cancel withdraws a pending request.
	Cancel(handleToken string) (*Request, error)
	// Expire times out a pending or failed request.
	Expire(handleToken string) (*Request, error)
	// Get returns a tracked request by handle token.
	Get(handleToken string) (*Request, bool)
	// List returns all tracked requests in creation order.
	List() []*Request
}

// FakeBroker is an in-memory, test-gated Broker. It performs no real Portal
// calls and is deterministic: request handle tokens use a monotonic sequence.
type FakeBroker struct {
	sequence int
	order    []string
	requests map[string]*Request
}

// NewFakeBroker returns an empty in-memory broker.
func NewFakeBroker() *FakeBroker {
	return &FakeBroker{requests: map[string]*Request{}}
}

// CreateRequest creates and tracks a new portal request object.
func (b *FakeBroker) CreateRequest(spec RequestSpec) (*Request, error) {
	b.sequence++
	req, err := newRequest(spec, b.sequence)
	if err != nil {
		b.sequence-- // do not consume a sequence number on invalid input
		return nil, err
	}
	if _, exists := b.requests[req.HandleToken]; exists {
		return nil, fmt.Errorf("portal request %q already exists", req.HandleToken)
	}
	b.requests[req.HandleToken] = req
	b.order = append(b.order, req.HandleToken)
	return req, nil
}

// Resolve records the Portal response for a pending request.
func (b *FakeBroker) Resolve(handleToken string, outcome Outcome) (*Request, error) {
	req, ok := b.requests[handleToken]
	if !ok {
		return nil, fmt.Errorf("unknown portal request %q", handleToken)
	}
	if req.Terminal() {
		return nil, fmt.Errorf("portal request %q is already %s", handleToken, req.State)
	}
	if req.State != StatePendingUserMediation && req.State != StateFailed {
		return nil, fmt.Errorf("portal request %q cannot be resolved from state %s", handleToken, req.State)
	}

	switch outcome {
	case OutcomeGranted:
		req.State = StateGranted
		req.PermissionState = PermissionGranted
		req.Recoverable = false
	case OutcomeDenied:
		req.State = StateDenied
		req.PermissionState = PermissionDenied
		req.Recoverable = false
	case OutcomeCancelled:
		req.State = StateCancelled
		req.PermissionState = PermissionNotGranted
		req.Recoverable = false
	case OutcomeFailed:
		// A transport failure is recoverable: the request stays open for retry
		// and the failure is visible in diagnostics.
		req.State = StateFailed
		req.PermissionState = PermissionPending
		req.Recoverable = true
		req.Diagnostics = append(req.Diagnostics, "Portal transport failed; the request can be retried.")
	case OutcomeExpired:
		// A timeout is terminal: no permission is granted and the request can
		// no longer be resolved.
		req.State = StateExpired
		req.PermissionState = PermissionNotGranted
		req.Recoverable = false
		req.Diagnostics = append(req.Diagnostics, "Portal request expired before a user response.")
	default:
		return nil, errors.New("outcome must be one of: granted, denied, cancelled, failed, expired")
	}
	return req, nil
}

// Complete marks a granted request as consumed.
func (b *FakeBroker) Complete(handleToken string) (*Request, error) {
	req, ok := b.requests[handleToken]
	if !ok {
		return nil, fmt.Errorf("unknown portal request %q", handleToken)
	}
	if req.State != StateGranted {
		return nil, fmt.Errorf("portal request %q must be granted before completion, not %s", handleToken, req.State)
	}
	req.State = StateCompleted
	return req, nil
}

// Cancel withdraws a pending or failed request.
func (b *FakeBroker) Cancel(handleToken string) (*Request, error) {
	req, ok := b.requests[handleToken]
	if !ok {
		return nil, fmt.Errorf("unknown portal request %q", handleToken)
	}
	if req.Terminal() || req.State == StateGranted {
		return nil, fmt.Errorf("portal request %q cannot be cancelled from state %s", handleToken, req.State)
	}
	req.State = StateCancelled
	req.PermissionState = PermissionNotGranted
	return req, nil
}

// Expire times out a pending or failed request. Terminal or granted requests
// cannot expire.
func (b *FakeBroker) Expire(handleToken string) (*Request, error) {
	req, ok := b.requests[handleToken]
	if !ok {
		return nil, fmt.Errorf("unknown portal request %q", handleToken)
	}
	if req.State != StatePendingUserMediation && req.State != StateFailed {
		return nil, fmt.Errorf("portal request %q cannot expire from state %s", handleToken, req.State)
	}
	req.State = StateExpired
	req.PermissionState = PermissionNotGranted
	req.Recoverable = false
	req.Diagnostics = append(req.Diagnostics, "Portal request expired before a user response.")
	return req, nil
}

// Get returns a tracked request by handle token.
func (b *FakeBroker) Get(handleToken string) (*Request, bool) {
	req, ok := b.requests[handleToken]
	return req, ok
}

// List returns all tracked requests in creation order.
func (b *FakeBroker) List() []*Request {
	out := make([]*Request, 0, len(b.order))
	for _, token := range b.order {
		out = append(out, b.requests[token])
	}
	return out
}

// compile-time assertion that FakeBroker satisfies Broker.
var _ Broker = (*FakeBroker)(nil)
