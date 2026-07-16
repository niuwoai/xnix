package execution

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRestrictedAuthorizationStorePersistsPreparationOnlyReceipt(t *testing.T) {
	root := t.TempDir()
	store, err := NewRestrictedAuthorizationStore(root)
	if err != nil {
		t.Fatalf("NewRestrictedAuthorizationStore returned error: %v", err)
	}
	receipt, err := store.Record(RestrictedAuthorizationRequest{
		ApplicationID: "org.xnix.sample.notepad",
		RequestID:     "xnix-exec-org-xnix-sample-notepad-1",
		Mode:          RestrictedTestMode,
		Scope:         RestrictedTestPreparationScope,
		Directive:     RestrictedTestPreparationDirective,
	})
	if err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	loaded, err := store.Load(receipt.AuthorizationID)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if loaded.SHA256 != receipt.SHA256 || len(loaded.SHA256) != 64 || !loaded.PreparationAuthorized || loaded.State != "authorized-preparation-only" || filepath.IsAbs(loaded.RelativePath) {
		t.Fatalf("unexpected restricted authorization receipt: %+v", loaded)
	}
	if restrictedAuthorizationUnsafe(loaded) {
		t.Fatalf("restricted authorization enabled unsafe gate: %+v", loaded)
	}
}

func TestRestrictedAuthorizationStoreRejectsInvalidDirectiveTamperingAndSymlink(t *testing.T) {
	t.Run("invalid directive", func(t *testing.T) {
		store, err := NewRestrictedAuthorizationStore(t.TempDir())
		if err != nil {
			t.Fatalf("NewRestrictedAuthorizationStore returned error: %v", err)
		}
		if _, err := store.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: "yes"}); err == nil {
			t.Fatal("expected imprecise authorization directive to be rejected")
		}
	})

	t.Run("tampered receipt", func(t *testing.T) {
		root := t.TempDir()
		store, err := NewRestrictedAuthorizationStore(root)
		if err != nil {
			t.Fatalf("NewRestrictedAuthorizationStore returned error: %v", err)
		}
		receipt, err := store.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: RestrictedTestPreparationDirective})
		if err != nil {
			t.Fatalf("Record returned error: %v", err)
		}
		path := filepath.Join(root, filepath.FromSlash(receipt.RelativePath))
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read receipt: %v", err)
		}
		tampered := bytes.Replace(data, []byte(`"launch_authorized": false`), []byte(`"launch_authorized": true`), 1)
		if err := os.WriteFile(path, tampered, 0o600); err != nil {
			t.Fatalf("tamper receipt: %v", err)
		}
		if _, err := store.Load(receipt.AuthorizationID); err == nil {
			t.Fatal("expected tampered authorization receipt to be rejected")
		}
	})

	t.Run("managed path symlink", func(t *testing.T) {
		root := t.TempDir()
		outside := t.TempDir()
		if err := os.Symlink(outside, filepath.Join(root, "execution-ledger")); err != nil {
			t.Skipf("symlinks unavailable: %v", err)
		}
		store, err := NewRestrictedAuthorizationStore(root)
		if err != nil {
			t.Fatalf("NewRestrictedAuthorizationStore returned error: %v", err)
		}
		if _, err := store.Record(RestrictedAuthorizationRequest{ApplicationID: "org.xnix.sample.notepad", RequestID: "request-1", Mode: RestrictedTestMode, Scope: RestrictedTestPreparationScope, Directive: RestrictedTestPreparationDirective}); err == nil {
			t.Fatal("expected managed path symlink to be rejected")
		}
		entries, err := os.ReadDir(outside)
		if err != nil {
			t.Fatalf("inspect outside directory: %v", err)
		}
		if len(entries) != 0 {
			t.Fatalf("authorization store wrote through managed path symlink: %+v", entries)
		}
	})
}
