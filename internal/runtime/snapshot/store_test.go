package snapshot

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, root, rel, content string) {
	t.Helper()
	abs := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(abs), 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(abs, []byte(content), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
}

func readFile(t *testing.T, root, rel string) (string, bool) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		if os.IsNotExist(err) {
			return "", false
		}
		t.Fatalf("ReadFile: %v", err)
	}
	return string(data), true
}

func TestCreateVerifyList(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "app/data.txt", "one")
	writeFile(t, root, "runtime/meta.json", "{}")

	store, err := New(root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	manifest, err := store.Create("before-repair-1", "before-repair")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if manifest.FileCount != 2 || manifest.Reason != "before-repair" {
		t.Fatalf("unexpected manifest: %#v", manifest)
	}
	// Files must be sorted and content-addressed with digests.
	if manifest.Files[0].Path != "app/data.txt" || manifest.Files[1].Path != "runtime/meta.json" {
		t.Fatalf("files not sorted: %#v", manifest.Files)
	}
	if err := store.Verify("before-repair-1"); err != nil {
		t.Fatalf("Verify: %v", err)
	}

	list, err := store.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 1 || list[0].ID != "before-repair-1" {
		t.Fatalf("unexpected list: %#v", list)
	}
}

func TestRestoreRollsBackFixtureState(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "app/keep.txt", "original")

	store, err := New(root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if _, err := store.Create("baseline", "manual"); err != nil {
		t.Fatalf("Create: %v", err)
	}

	// Mutate the state root: change a file and add a new one.
	writeFile(t, root, "app/keep.txt", "modified")
	writeFile(t, root, "app/added.txt", "new")

	receipt, err := store.Restore("baseline")
	if err != nil {
		t.Fatalf("Restore: %v", err)
	}
	if !contains(receipt.Restored, "app/keep.txt") {
		t.Fatalf("keep.txt should have been restored: %#v", receipt)
	}
	if !contains(receipt.Removed, "app/added.txt") {
		t.Fatalf("added.txt should have been removed: %#v", receipt)
	}
	if receipt.HostRootTouched {
		t.Fatalf("rollback must not touch the host root")
	}

	if got, ok := readFile(t, root, "app/keep.txt"); !ok || got != "original" {
		t.Fatalf("keep.txt not rolled back: %q ok=%v", got, ok)
	}
	if _, ok := readFile(t, root, "app/added.txt"); ok {
		t.Fatalf("added.txt should be gone after rollback")
	}
}

func TestVerifyDetectsCorruption(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "app/data.txt", "content")
	store, err := New(root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	manifest, err := store.Create("snap", "manual")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	// Corrupt the stored object.
	obj := store.objectPath(manifest.Files[0].Digest)
	if err := os.WriteFile(obj, []byte("tampered"), 0o600); err != nil {
		t.Fatalf("tamper: %v", err)
	}
	if err := store.Verify("snap"); err == nil {
		t.Fatalf("Verify must detect a corrupt object")
	}
}

func TestStoreRefusesPathsOutsideRoot(t *testing.T) {
	root := t.TempDir()
	store, err := New(root)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	// within() must reject escapes and accept in-root paths.
	if store.within(filepath.Join(root, "..", "escape")) {
		t.Fatalf("within must reject parent-escape paths")
	}
	if store.within(filepath.Dir(root)) {
		t.Fatalf("within must reject the parent directory")
	}
	if !store.within(filepath.Join(root, "app", "data.txt")) {
		t.Fatalf("within must accept in-root paths")
	}
	// Invalid ids (traversal) must be rejected.
	if !validSnapshotID("before-repair.1_ok") {
		t.Fatalf("valid id rejected")
	}
	for _, bad := range []string{"", "..", ".", "a/b", "a\\b"} {
		if validSnapshotID(bad) {
			t.Fatalf("invalid id %q accepted", bad)
		}
	}
}

func TestMissingStateRootFails(t *testing.T) {
	if _, err := New(filepath.Join(t.TempDir(), "does-not-exist")); err == nil {
		t.Fatalf("New must fail when the state root does not exist")
	}
	if _, err := New(""); err == nil {
		t.Fatalf("New must fail on an empty state root")
	}
}

func contains(values []string, target string) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}
	return false
}
