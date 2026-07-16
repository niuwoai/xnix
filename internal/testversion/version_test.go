package testversion

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFromFindsCanonicalVersionInParent(t *testing.T) {
	root := t.TempDir()
	nested := filepath.Join(root, "internal", "runtime")
	if err := os.MkdirAll(nested, 0o700); err != nil {
		t.Fatalf("create nested test directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("9.8.7\n"), 0o600); err != nil {
		t.Fatalf("write VERSION fixture: %v", err)
	}

	version, err := ReadFrom(nested)
	if err != nil {
		t.Fatalf("ReadFrom returned error: %v", err)
	}
	if version != "9.8.7" {
		t.Fatalf("ReadFrom returned %q, want 9.8.7", version)
	}
}

func TestReadFromRejectsEmptyVersion(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("\n"), 0o600); err != nil {
		t.Fatalf("write empty VERSION fixture: %v", err)
	}

	if _, err := ReadFrom(root); err == nil {
		t.Fatal("ReadFrom accepted an empty VERSION file")
	}
}
