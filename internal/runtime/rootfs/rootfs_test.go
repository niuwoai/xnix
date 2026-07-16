package rootfs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenRequiresExistingDirectory(t *testing.T) {
	if _, err := Open(""); err == nil {
		t.Fatalf("empty root must fail")
	}
	if _, err := Open(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatalf("missing root must fail")
	}
	// A file is not a directory.
	file := filepath.Join(t.TempDir(), "f")
	if err := os.WriteFile(file, []byte("x"), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if _, err := Open(file); err == nil {
		t.Fatalf("file root must fail")
	}
	if _, err := Open(string(os.PathSeparator)); err == nil {
		t.Fatalf("filesystem root must be refused")
	}
	if _, err := Open(t.TempDir()); err != nil {
		t.Fatalf("valid dir should open: %v", err)
	}
}

func TestEnsureCreatesRoot(t *testing.T) {
	target := filepath.Join(t.TempDir(), "a", "b", "c")
	root, err := Ensure(target)
	if err != nil {
		t.Fatalf("Ensure: %v", err)
	}
	if info, statErr := os.Stat(root.Path()); statErr != nil || !info.IsDir() {
		t.Fatalf("Ensure did not create the root: %v", statErr)
	}
	if _, err := Ensure(string(os.PathSeparator)); err == nil {
		t.Fatalf("Ensure must refuse the filesystem root")
	}
}

func TestResolveConfinesToRoot(t *testing.T) {
	root, _ := Open(t.TempDir())

	ok, err := root.Resolve("a", "b.json")
	if err != nil {
		t.Fatalf("Resolve valid: %v", err)
	}
	if !root.Contains(ok) {
		t.Fatalf("resolved path must be contained: %q", ok)
	}

	escapes := [][]string{
		{".."},
		{"..", "escape"},
		{"a", "..", "..", "escape"},
		{"../escape"},
	}
	for _, elems := range escapes {
		if _, err := root.Resolve(elems...); err == nil {
			t.Fatalf("escape %v must be refused", elems)
		}
	}

	// Absolute and empty elements are refused.
	if _, err := root.Resolve(string(os.PathSeparator) + "etc"); err == nil {
		t.Fatalf("absolute element must be refused")
	}
	if _, err := root.Resolve(""); err == nil {
		t.Fatalf("empty element must be refused")
	}
	if _, err := root.Resolve(); err == nil {
		t.Fatalf("no elements must be refused")
	}
}

func TestContains(t *testing.T) {
	dir := t.TempDir()
	root, _ := Open(dir)
	if !root.Contains(dir) {
		t.Fatalf("root must contain itself")
	}
	if !root.Contains(filepath.Join(dir, "child", "x")) {
		t.Fatalf("root must contain descendants")
	}
	if root.Contains(filepath.Dir(dir)) {
		t.Fatalf("root must not contain its parent")
	}
	if root.Contains(filepath.Join(dir, "..", "sibling")) {
		t.Fatalf("root must not contain a parent-escape path")
	}
}

func TestReadWriteFileStayInRoot(t *testing.T) {
	root, _ := Open(t.TempDir())

	if err := root.WriteFile("nested/data.json", []byte("hello"), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	data, err := root.ReadFile("nested/data.json")
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("unexpected content: %q", data)
	}
	// The file really lives under the root.
	if _, err := os.Stat(filepath.Join(root.Path(), "nested", "data.json")); err != nil {
		t.Fatalf("file not written under root: %v", err)
	}
	// Escapes are refused for both read and write.
	if err := root.WriteFile("../escape.json", []byte("x"), 0o600); err == nil {
		t.Fatalf("write escape must be refused")
	}
	if _, err := root.ReadFile("../escape.json"); err == nil {
		t.Fatalf("read escape must be refused")
	}
}
