package portal

import "testing"

func TestFakeBrokerGrantThenCompleteFlow(t *testing.T) {
	broker := NewFakeBroker()
	req, err := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open", Reason: "Open a spreadsheet"})
	if err != nil {
		t.Fatalf("CreateRequest returned error: %v", err)
	}
	if req.Reason != "Open a spreadsheet" {
		t.Fatalf("reason not preserved: %q", req.Reason)
	}

	granted, err := broker.Resolve(req.HandleToken, OutcomeGranted)
	if err != nil {
		t.Fatalf("Resolve granted returned error: %v", err)
	}
	if granted.State != StateGranted || granted.PermissionState != PermissionGranted {
		t.Fatalf("grant did not update state: %#v", granted)
	}

	// Cannot complete before... it is granted, so complete should work now.
	completed, err := broker.Complete(req.HandleToken)
	if err != nil {
		t.Fatalf("Complete returned error: %v", err)
	}
	if completed.State != StateCompleted || !completed.Terminal() {
		t.Fatalf("completion did not finalize: %#v", completed)
	}
	// Permission stays granted after completion (tracked separately).
	if completed.PermissionState != PermissionGranted {
		t.Fatalf("permission state must survive completion: %#v", completed)
	}
	// Re-resolving a completed request is invalid.
	if _, err := broker.Resolve(req.HandleToken, OutcomeDenied); err == nil {
		t.Fatalf("resolving a completed request must error")
	}
}

func TestFakeBrokerFailedIsRecoverableThenRetried(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "print"})

	failed, err := broker.Resolve(req.HandleToken, OutcomeFailed)
	if err != nil {
		t.Fatalf("Resolve failed returned error: %v", err)
	}
	if failed.State != StateFailed || !failed.Recoverable || failed.PermissionState != PermissionPending {
		t.Fatalf("failure must be recoverable and keep permission pending: %#v", failed)
	}
	if len(failed.Diagnostics) == 0 {
		t.Fatalf("failure must be visible in diagnostics")
	}
	// A failed (recoverable) request can be retried to a grant.
	granted, err := broker.Resolve(req.HandleToken, OutcomeGranted)
	if err != nil {
		t.Fatalf("retry after failure returned error: %v", err)
	}
	if granted.State != StateGranted {
		t.Fatalf("retry did not grant: %#v", granted)
	}
}

func TestFakeBrokerCompleteRequiresGrant(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	if _, err := broker.Complete(req.HandleToken); err == nil {
		t.Fatalf("completing a pending request must error")
	}
}

func TestFakeBrokerExpirePendingIsTerminal(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})

	expired, err := broker.Expire(req.HandleToken)
	if err != nil {
		t.Fatalf("Expire: %v", err)
	}
	if expired.State != StateExpired || expired.PermissionState != PermissionNotGranted || expired.Recoverable {
		t.Fatalf("expired request state wrong: %#v", expired)
	}
	if !expired.Terminal() {
		t.Fatalf("expired request must be terminal")
	}
	if len(expired.Diagnostics) == 0 {
		t.Fatalf("expiry must be visible in diagnostics")
	}
	// An expired request can no longer be resolved or expired again.
	if _, err := broker.Resolve(req.HandleToken, OutcomeGranted); err == nil {
		t.Fatalf("resolving an expired request must error")
	}
	if _, err := broker.Expire(req.HandleToken); err == nil {
		t.Fatalf("re-expiring must error")
	}
}

func TestFakeBrokerExpireViaResolveOutcome(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "print"})
	expired, err := broker.Resolve(req.HandleToken, OutcomeExpired)
	if err != nil {
		t.Fatalf("Resolve expired: %v", err)
	}
	if expired.State != StateExpired || expired.PermissionState != PermissionNotGranted {
		t.Fatalf("resolve-expired state wrong: %#v", expired)
	}
}

func TestFakeBrokerGrantedCannotExpire(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	if _, err := broker.Resolve(req.HandleToken, OutcomeGranted); err != nil {
		t.Fatalf("Resolve granted: %v", err)
	}
	if _, err := broker.Expire(req.HandleToken); err == nil {
		t.Fatalf("a granted request must not expire")
	}
}

func TestFakeBrokerCancelPending(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "clipboard"})
	cancelled, err := broker.Cancel(req.HandleToken)
	if err != nil {
		t.Fatalf("Cancel returned error: %v", err)
	}
	if cancelled.State != StateCancelled || cancelled.PermissionState != PermissionNotGranted {
		t.Fatalf("cancel did not finalize: %#v", cancelled)
	}
	if _, err := broker.Cancel(req.HandleToken); err == nil {
		t.Fatalf("cancelling a terminal request must error")
	}
}

func TestFakeBrokerListAndGetTrackInCreationOrder(t *testing.T) {
	broker := NewFakeBroker()
	a, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.a", Operation: "file-open"})
	b, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.b", Operation: "uri-open"})

	list := broker.List()
	if len(list) != 2 || list[0].HandleToken != a.HandleToken || list[1].HandleToken != b.HandleToken {
		t.Fatalf("List must return creation order: %#v", list)
	}
	got, ok := broker.Get(b.HandleToken)
	if !ok || got.ApplicationID != "org.example.b" {
		t.Fatalf("Get did not return tracked request: %#v ok=%v", got, ok)
	}
	if _, ok := broker.Get("missing"); ok {
		t.Fatalf("Get must report missing requests")
	}
}
