package portal

import "testing"

func TestNewRequestAskOperationStartsPending(t *testing.T) {
	broker := NewFakeBroker()
	req, err := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	if err != nil {
		t.Fatalf("CreateRequest returned error: %v", err)
	}
	if req.HandleToken != "xnix_org_example_ledger_file_open_1" {
		t.Fatalf("unexpected handle token: %q", req.HandleToken)
	}
	if req.State != StatePendingUserMediation || req.PermissionState != PermissionPending {
		t.Fatalf("ask operation must start pending: %#v", req)
	}
	if req.Interface != "org.freedesktop.portal.FileChooser" || req.Method != "OpenFile" {
		t.Fatalf("unexpected portal binding: %#v", req)
	}
	if req.Destination != PortalDestination || req.ObjectPath != PortalObjectPath {
		t.Fatalf("unexpected portal destination: %#v", req)
	}
	if !req.RequestObjectRequired || !req.UserMediationRequired || req.DirectAccessAllowed {
		t.Fatalf("ask operation must require a mediated request object: %#v", req)
	}
	if req.BackendDetailsExposed || req.HostPermissionChanged {
		t.Fatalf("request must not expose backend details or change host permission: %#v", req)
	}
	if req.Terminal() {
		t.Fatalf("pending request must not be terminal")
	}
}

func TestNewRequestDenyOperationIsTerminal(t *testing.T) {
	broker := NewFakeBroker()
	req, err := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "camera"})
	if err != nil {
		t.Fatalf("CreateRequest returned error: %v", err)
	}
	if req.State != StateDenied || req.PermissionState != PermissionDenied || !req.Terminal() {
		t.Fatalf("deny operation must be terminal denied: %#v", req)
	}
	if len(req.Diagnostics) == 0 {
		t.Fatalf("deny operation must record a diagnostic")
	}
	// Resolving a policy-denied request is invalid.
	if _, err := broker.Resolve(req.HandleToken, OutcomeGranted); err == nil {
		t.Fatalf("resolving a terminal request must error")
	}
}

func TestNewRequestRejectsBadInput(t *testing.T) {
	broker := NewFakeBroker()
	if _, err := broker.CreateRequest(RequestSpec{ApplicationID: "not-a-dns-id", Operation: "file-open"}); err == nil {
		t.Fatalf("invalid application id must error")
	}
	if _, err := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "unknown"}); err == nil {
		t.Fatalf("unknown operation must error")
	}
	// A rejected creation must not consume a sequence number.
	req, err := broker.CreateRequest(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	if err != nil {
		t.Fatalf("CreateRequest returned error: %v", err)
	}
	if req.HandleToken != "xnix_org_example_ledger_file_open_1" {
		t.Fatalf("sequence must not advance on invalid input: %q", req.HandleToken)
	}
}

func TestSupportedOperationsSorted(t *testing.T) {
	ops := SupportedOperations()
	want := []string{"camera", "clipboard", "file-open", "print", "remote-desktop", "screenshot", "uri-open"}
	if len(ops) != len(want) {
		t.Fatalf("unexpected operations: %#v", ops)
	}
	for i := range want {
		if ops[i] != want[i] {
			t.Fatalf("operations not sorted: %#v", ops)
		}
	}
}
