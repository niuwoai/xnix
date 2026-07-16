package portal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadRequestsMissingDirectoryIsEmpty(t *testing.T) {
	stateRoot := t.TempDir()
	requests, malformed, err := ReadRequests(stateRoot)
	if err != nil {
		t.Fatalf("ReadRequests returned error: %v", err)
	}
	if len(requests) != 0 || len(malformed) != 0 {
		t.Fatalf("missing ledger should be empty: %d requests, %d malformed", len(requests), len(malformed))
	}
	// The read must not create the ledger directory.
	if _, err := os.Stat(filepath.Join(stateRoot, "portal-requests")); !os.IsNotExist(err) {
		t.Fatalf("ReadRequests must not create the ledger directory: %v", err)
	}
}

func TestReadRequestsParsesAndFlagsMalformed(t *testing.T) {
	stateRoot := t.TempDir()
	ledger, err := NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger: %v", err)
	}
	record, err := ledger.Create(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open", Reason: "test"})
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	if _, err := ledger.Resolve(record.Request.HandleToken, OutcomeGranted); err != nil {
		t.Fatalf("resolve request: %v", err)
	}
	if err := os.WriteFile(filepath.Join(stateRoot, "portal-requests", "broken-1.json"), []byte("not-json"), 0o600); err != nil {
		t.Fatalf("write malformed record: %v", err)
	}

	requests, malformed, err := ReadRequests(stateRoot)
	if err != nil {
		t.Fatalf("ReadRequests returned error: %v", err)
	}
	if len(requests) != 1 || requests[0].ApplicationID != "org.example.ledger" {
		t.Fatalf("expected one parsed request, got %+v", requests)
	}
	if len(malformed) != 1 || malformed[0] != "broken-1" {
		t.Fatalf("expected broken-1 flagged malformed, got %+v", malformed)
	}

	summary := Summarize(requests, "org.example.ledger")
	if len(summary.Granted) != 1 || summary.Granted[0] != "file-open" {
		t.Fatalf("summary should report a granted file-open, got %+v", summary)
	}
}

func TestReadRequestsRequiresStateRoot(t *testing.T) {
	if _, _, err := ReadRequests(""); err == nil {
		t.Fatalf("expected empty state root to be rejected")
	}
}
