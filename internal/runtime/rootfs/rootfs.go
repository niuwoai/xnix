// Package rootfs is the canonical controlled-root path-confinement boundary for
// the Compatibility Runtime. Every Runtime store that persists state under a
// caller-provided root (snapshots, artifact caches, recipe roots, execution
// ledgers, diagnostic records, image checks) must confine all reads and writes
// to that root and refuse any path that escapes it.
//
// That security-critical check — resolve to an absolute path, reject the
// filesystem root, and reject any relative path that climbs out with "..", was
// previously reimplemented at many call sites with subtly different behavior.
// This package is the single reviewed, tested implementation.
package rootfs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Root is a validated controlled directory. All paths derived from it are
// guaranteed to stay inside it.
type Root struct {
	path string
}

// Open returns a Root for an existing directory. The directory must exist and
// must not be the filesystem root.
func Open(dir string) (*Root, error) {
	abs, err := normalize(dir)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("controlled root must exist: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("controlled root %q is not a directory", abs)
	}
	return &Root{path: abs}, nil
}

// Ensure returns a Root for a directory, creating it (and parents) if needed.
// It still refuses the filesystem root.
func Ensure(dir string) (*Root, error) {
	abs, err := normalize(dir)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, fmt.Errorf("initialize controlled root: %w", err)
	}
	return &Root{path: abs}, nil
}

func normalize(dir string) (string, error) {
	if dir == "" {
		return "", errors.New("controlled root must not be empty")
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("resolve controlled root: %w", err)
	}
	abs = filepath.Clean(abs)
	if abs == string(os.PathSeparator) {
		return "", errors.New("refusing to use the filesystem root as a controlled root")
	}
	return abs, nil
}

// Path returns the absolute controlled-root directory.
func (r *Root) Path() string { return r.path }

// Contains reports whether an absolute path is inside the root (the root itself
// counts as contained).
func (r *Root) Contains(abs string) bool {
	rel, err := filepath.Rel(r.path, abs)
	if err != nil {
		return false
	}
	return relInside(rel)
}

// relInside reports whether a path relative to the root stays inside it.
func relInside(rel string) bool {
	if rel == "." {
		return true
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(os.PathSeparator))
}

// Resolve joins one or more relative path elements to the root and returns the
// absolute path, refusing absolute elements and any result that escapes the
// root. It is the safe replacement for filepath.Join at a trust boundary.
func (r *Root) Resolve(elems ...string) (string, error) {
	if len(elems) == 0 {
		return "", errors.New("resolve requires at least one path element")
	}
	for _, elem := range elems {
		if elem == "" {
			return "", errors.New("path element must not be empty")
		}
		if filepath.IsAbs(elem) {
			return "", fmt.Errorf("path element must be relative: %q", elem)
		}
	}
	joined := filepath.Join(append([]string{r.path}, toNative(elems)...)...)
	rel, err := filepath.Rel(r.path, joined)
	if err != nil {
		return "", err
	}
	if !relInside(rel) {
		return "", fmt.Errorf("path escapes controlled root: %q", filepath.Join(elems...))
	}
	return joined, nil
}

func toNative(elems []string) []string {
	out := make([]string, len(elems))
	for i, e := range elems {
		out[i] = filepath.FromSlash(e)
	}
	return out
}

// ReadFile reads a file at a root-relative path, refusing escapes.
func (r *Root) ReadFile(rel string) ([]byte, error) {
	abs, err := r.Resolve(rel)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(abs)
}

// WriteFile writes data to a root-relative path, creating parent directories
// inside the root and refusing escapes.
func (r *Root) WriteFile(rel string, data []byte, perm os.FileMode) error {
	abs, err := r.Resolve(rel)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(abs), 0o700); err != nil {
		return err
	}
	return os.WriteFile(abs, data, perm)
}
