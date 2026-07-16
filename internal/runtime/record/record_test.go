package record

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/rootfs"
)

type sample struct {
	ID    string `json:"id"`
	Count int    `json:"count"`
}

func TestSaveLoadRoundTrip(t *testing.T) {
	root, err := rootfs.Ensure(t.TempDir())
	if err != nil {
		t.Fatalf("Ensure: %v", err)
	}
	in := sample{ID: "org.example.ledger", Count: 3}

	digest, err := Save(root, "records/ledger.json", in)
	if err != nil {
		t.Fatalf("Save: %v", err)
	}
	if len(digest) != 64 {
		t.Fatalf("unexpected digest: %q", digest)
	}

	var out sample
	if err := Load(root, "records/ledger.json", &out); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if out != in {
		t.Fatalf("round trip mismatch: %#v != %#v", out, in)
	}

	// The digest must match a re-encode of the same value.
	want, _ := Digest(in)
	if digest != want {
		t.Fatalf("save digest %q != value digest %q", digest, want)
	}
}

func TestSaveIsAtomicAndLeavesNoTempFile(t *testing.T) {
	dir := t.TempDir()
	root, _ := rootfs.Ensure(dir)
	if _, err := Save(root, "a/b.json", sample{ID: "x", Count: 1}); err != nil {
		t.Fatalf("Save: %v", err)
	}
	// No temporary write file may remain.
	entries, _ := os.ReadDir(filepath.Join(dir, "a"))
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), tempSuffix) {
			t.Fatalf("atomic write left a temp file: %s", e.Name())
		}
	}
	if len(entries) != 1 || entries[0].Name() != "b.json" {
		t.Fatalf("unexpected directory contents: %#v", entries)
	}
}

func TestSaveOverwritesAtomically(t *testing.T) {
	root, _ := rootfs.Ensure(t.TempDir())
	if _, err := Save(root, "r.json", sample{ID: "first", Count: 1}); err != nil {
		t.Fatalf("Save first: %v", err)
	}
	if _, err := Save(root, "r.json", sample{ID: "second", Count: 2}); err != nil {
		t.Fatalf("Save second: %v", err)
	}
	var out sample
	if err := Load(root, "r.json", &out); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if out.ID != "second" || out.Count != 2 {
		t.Fatalf("overwrite did not replace content: %#v", out)
	}
}

func TestSaveRefusesEscapeAndNilRoot(t *testing.T) {
	root, _ := rootfs.Ensure(t.TempDir())
	if _, err := Save(root, "../escape.json", sample{ID: "x"}); err == nil {
		t.Fatalf("escape path must be refused")
	}
	if _, err := Save(nil, "r.json", sample{}); err == nil {
		t.Fatalf("nil root must be refused")
	}
}

func TestLoadMissingAndCorrupt(t *testing.T) {
	dir := t.TempDir()
	root, _ := rootfs.Ensure(dir)
	var out sample
	if err := Load(root, "missing.json", &out); err == nil {
		t.Fatalf("missing record must error")
	}
	// Write invalid JSON directly, then Load must report corruption.
	if err := os.WriteFile(filepath.Join(dir, "bad.json"), []byte("{not json"), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := Load(root, "bad.json", &out); err == nil {
		t.Fatalf("corrupt record must error")
	}
}

func TestLoadAllOrdersSkipsAndToleratesMissingDir(t *testing.T) {
	dir := t.TempDir()
	root, _ := rootfs.Ensure(dir)

	// A missing directory yields an empty slice, not an error.
	empty, err := LoadAll[sample](root, "records")
	if err != nil {
		t.Fatalf("LoadAll on missing dir: %v", err)
	}
	if len(empty) != 0 {
		t.Fatalf("missing dir must yield no records: %#v", empty)
	}

	// Filenames are returned in sorted order regardless of write order.
	Save(root, "records/b.json", sample{ID: "second", Count: 2})
	Save(root, "records/a.json", sample{ID: "first", Count: 1})
	Save(root, "records/c.json", sample{ID: "third", Count: 3})

	// A subdirectory, a non-.json file, and a lingering temp file must be
	// skipped rather than decoded.
	if err := os.MkdirAll(filepath.Join(dir, "records", "nested"), 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	os.WriteFile(filepath.Join(dir, "records", "notes.txt"), []byte("ignore"), 0o600)
	os.WriteFile(filepath.Join(dir, "records", "d.json"+tempSuffix), []byte("{partial"), 0o600)

	all, err := LoadAll[sample](root, "records")
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	if len(all) != 3 {
		t.Fatalf("expected three records, got %#v", all)
	}
	if all[0].ID != "first" || all[1].ID != "second" || all[2].ID != "third" {
		t.Fatalf("records not ordered by filename: %#v", all)
	}
}

func TestLoadAllReportsCorruptRecordAndNilRoot(t *testing.T) {
	dir := t.TempDir()
	root, _ := rootfs.Ensure(dir)
	if err := os.MkdirAll(filepath.Join(dir, "records"), 0o700); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	os.WriteFile(filepath.Join(dir, "records", "bad.json"), []byte("{not json"), 0o600)
	if _, err := LoadAll[sample](root, "records"); err == nil {
		t.Fatalf("a corrupt record must surface an error")
	}
	if _, err := LoadAll[sample](nil, "records"); err == nil {
		t.Fatalf("nil root must be refused")
	}
}

func TestExists(t *testing.T) {
	root, _ := rootfs.Ensure(t.TempDir())
	ok, err := Exists(root, "r.json")
	if err != nil || ok {
		t.Fatalf("missing record must not exist: ok=%v err=%v", ok, err)
	}
	if _, err := Save(root, "r.json", sample{ID: "x"}); err != nil {
		t.Fatalf("Save: %v", err)
	}
	ok, err = Exists(root, "r.json")
	if err != nil || !ok {
		t.Fatalf("saved record must exist: ok=%v err=%v", ok, err)
	}
}

func TestDigestIsDeterministic(t *testing.T) {
	a, _ := Digest(sample{ID: "x", Count: 1})
	b, _ := Digest(sample{ID: "x", Count: 1})
	if a != b || len(a) != 64 {
		t.Fatalf("digest not deterministic: %q %q", a, b)
	}
	c, _ := Digest(sample{ID: "x", Count: 2})
	if c == a {
		t.Fatalf("different values must have different digests")
	}
}
