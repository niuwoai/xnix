package portal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLedgerPersistsGrantAndCompletionWithoutRealPortal(t *testing.T) {
	root := t.TempDir()
	ledger, err := NewLedger(root)
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	created, err := ledger.Create(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open", Reason: "Open a spreadsheet"})
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if created.SchemaVersion != "xnix.runtime.portal_request_record.v1" ||
		created.RecordType != "portal-permission-request-record" ||
		created.Source != "go-runtime-state-root-portal-broker" ||
		created.Action != "create" ||
		created.RelativePath != "portal-requests/xnix_org_example_ledger_file_open_1.json" ||
		created.Request.State != StatePendingUserMediation ||
		created.Request.PermissionState != PermissionPending ||
		!created.RuntimeOwned ||
		!created.GoRuntimeBacked ||
		created.KDEPolicyOwner ||
		created.StateRootPathExposed ||
		created.RealPortalCallEnabled ||
		!created.RequestObjectCreated ||
		created.PermissionGranted ||
		created.ExecutionApproved ||
		created.HostPermissionChanged ||
		created.HostRootModified ||
		created.BackendDetailsExposed ||
		created.NetworkRequired ||
		created.PrivilegedContainerRequired {
		t.Fatalf("unexpected created Portal record: %#v", created)
	}
	if _, err := os.Stat(filepath.Join(root, "portal-requests", "xnix_org_example_ledger_file_open_1.json")); err != nil {
		t.Fatalf("Portal request record must persist under state root: %v", err)
	}

	granted, err := ledger.Resolve(created.Request.HandleToken, OutcomeGranted)
	if err != nil {
		t.Fatalf("Resolve granted returned error: %v", err)
	}
	if granted.Action != "resolve" ||
		granted.Request.State != StateGranted ||
		granted.Request.PermissionState != PermissionGranted ||
		!granted.PermissionGranted ||
		granted.ExecutionApproved ||
		granted.RealPortalCallEnabled ||
		granted.HostRootModified {
		t.Fatalf("unexpected granted Portal record: %#v", granted)
	}

	completed, err := ledger.Complete(created.Request.HandleToken)
	if err != nil {
		t.Fatalf("Complete returned error: %v", err)
	}
	if completed.Action != "complete" ||
		completed.Request.State != StateCompleted ||
		completed.Request.PermissionState != PermissionGranted ||
		!completed.PermissionGranted ||
		completed.ExecutionApproved ||
		completed.HostPermissionChanged {
		t.Fatalf("unexpected completed Portal record: %#v", completed)
	}

	inspected, err := ledger.Inspect(created.Request.HandleToken)
	if err != nil {
		t.Fatalf("Inspect returned error: %v", err)
	}
	if inspected.Request.State != StateCompleted ||
		inspected.Request.PermissionState != PermissionGranted ||
		!inspected.PermissionGranted {
		t.Fatalf("inspect must preserve completed Portal state: %#v", inspected)
	}
}

func TestLedgerRecordsDeniedAndExpiredRequests(t *testing.T) {
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	denied, err := ledger.Create(RequestSpec{ApplicationID: "org.example.ledger", Operation: "camera"})
	if err != nil {
		t.Fatalf("Create camera returned error: %v", err)
	}
	if denied.Request.State != StateDenied ||
		denied.Request.PermissionState != PermissionDenied ||
		denied.PermissionGranted ||
		len(denied.Diagnostics) == 0 {
		t.Fatalf("camera policy denial must persist diagnostics: %#v", denied)
	}

	pending, err := ledger.Create(RequestSpec{ApplicationID: "org.example.ledger", Operation: "clipboard"})
	if err != nil {
		t.Fatalf("Create clipboard returned error: %v", err)
	}
	expired, err := ledger.Expire(pending.Request.HandleToken)
	if err != nil {
		t.Fatalf("Expire returned error: %v", err)
	}
	if expired.Request.State != StateExpired ||
		expired.Request.PermissionState != PermissionNotGranted ||
		expired.PermissionGranted ||
		len(expired.Diagnostics) == 0 {
		t.Fatalf("expired request must persist not-granted diagnostics: %#v", expired)
	}
}

func TestLedgerRejectsInvalidStateRootAndTransitions(t *testing.T) {
	if _, err := NewLedger(""); err == nil {
		t.Fatalf("empty state root must be rejected")
	}
	ledger, err := NewLedger(t.TempDir())
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	if _, err := ledger.Create(RequestSpec{ApplicationID: "bad id", Operation: "file-open"}); err == nil {
		t.Fatalf("bad application id must be rejected")
	}
	if _, err := ledger.Inspect("missing"); err == nil {
		t.Fatalf("missing request must be rejected")
	}
	pending, err := ledger.Create(RequestSpec{ApplicationID: "org.example.ledger", Operation: "file-open"})
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if _, err := ledger.Complete(pending.Request.HandleToken); err == nil {
		t.Fatalf("completion before grant must be rejected")
	}
}

func TestRequestRelativePathIsScoped(t *testing.T) {
	relativePath, err := RequestRelativePath("xnix_org_example_ledger_file_open_1")
	if err != nil {
		t.Fatalf("RequestRelativePath returned error: %v", err)
	}
	if relativePath != "portal-requests/xnix_org_example_ledger_file_open_1.json" {
		t.Fatalf("unexpected relative path: %s", relativePath)
	}
	if _, err := RequestRelativePath(""); err == nil {
		t.Fatalf("empty handle token must be rejected")
	}
}
