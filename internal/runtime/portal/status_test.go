package portal

import "testing"

func TestUserStatusPendingAwaitsUser(t *testing.T) {
	broker := NewFakeBroker()
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	us := req.UserStatus()
	if us.Operation != "file-open" || us.Status != "Waiting for your approval" || !us.AwaitingUser {
		t.Fatalf("pending user status wrong: %#v", us)
	}
	if us.PermissionGranted || us.BackendDetailsExposed {
		t.Fatalf("pending request must not report permission or expose backend: %#v", us)
	}
	if len(us.Actions) != 2 || us.Actions[0] != "approve" || us.Actions[1] != "deny" {
		t.Fatalf("pending actions wrong: %#v", us.Actions)
	}
}

func TestUserStatusPerTerminalState(t *testing.T) {
	cases := []struct {
		operation string
		resolve   func(b *FakeBroker, token string)
		status    string
		granted   bool
	}{
		{"file-open", func(b *FakeBroker, tok string) { b.Resolve(tok, OutcomeGranted) }, "Allowed", true},
		{"print", func(b *FakeBroker, tok string) { b.Resolve(tok, OutcomeDenied) }, "Blocked", false},
		{"clipboard", func(b *FakeBroker, tok string) { b.Cancel(tok) }, "Cancelled", false},
		{"screenshot", func(b *FakeBroker, tok string) { b.Resolve(tok, OutcomeFailed) }, "Temporarily unavailable", false},
		{"uri-open", func(b *FakeBroker, tok string) { b.Expire(tok) }, "Request expired", false},
	}
	for _, c := range cases {
		broker := NewFakeBroker()
		req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: c.operation})
		c.resolve(broker, req.HandleToken)
		got, _ := broker.Get(req.HandleToken)
		us := got.UserStatus()
		if us.Status != c.status || us.PermissionGranted != c.granted {
			t.Fatalf("operation %s: unexpected status %#v", c.operation, us)
		}
		if us.AwaitingUser {
			t.Fatalf("terminal/resolved request must not await user: %#v", us)
		}
		if us.Actions == nil {
			t.Fatalf("actions must never be nil: %#v", us)
		}
	}
}

func TestUserStatusDenyByPolicyIsBlocked(t *testing.T) {
	broker := NewFakeBroker()
	// camera is deny-by-policy at creation.
	req, _ := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "camera"})
	us := req.UserStatus()
	if us.Status != "Blocked" || us.PermissionGranted || us.AwaitingUser {
		t.Fatalf("policy-denied request must read as blocked: %#v", us)
	}
}
