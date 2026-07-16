package artifact

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/rootfs"
)

// Cache is a content-addressed artifact cache confined to a controlled root.
// It never writes outside that root and verifies every stored object's digest.
type Cache struct {
	root *rootfs.Root
}

// NewCache opens (creating if needed) a cache under the given controlled root.
// The root must be a project- or container-controlled directory, never the
// host root; the caller is responsible for choosing a sandboxed path.
func NewCache(rootDir string) (*Cache, error) {
	root, err := rootfs.Ensure(rootDir)
	if err != nil {
		return nil, err
	}
	return &Cache{root: root}, nil
}

// Root returns the absolute controlled cache root.
func (c *Cache) Root() string { return c.root.Path() }

// objectPath returns the content-addressed path for a namespace and digest,
// refusing any component that could escape the cache root.
func (c *Cache) objectPath(namespace, digest string) (string, error) {
	if !digestPattern.MatchString(digest) {
		return "", fmt.Errorf("cache digest must be a lowercase hex sha256: %q", digest)
	}
	if namespace == "" || strings.ContainsAny(namespace, `/\`) || namespace == "." || namespace == ".." {
		return "", fmt.Errorf("invalid cache namespace %q", namespace)
	}
	return c.root.Resolve(namespace, digest[:2], digest)
}

// Has reports whether an artifact is already cached and intact.
func (c *Cache) Has(namespace, digest string) bool {
	path, err := c.objectPath(namespace, digest)
	if err != nil {
		return false
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return digestOf(data) == digest
}

// Put stores artifact bytes under the namespace, verifying the digest first. A
// digest mismatch is refused, so corrupt or wrong content never lands in the
// cache. Put is idempotent for identical content.
func (c *Cache) Put(namespace, digest string, data []byte) error {
	if actual := digestOf(data); actual != digest {
		return fmt.Errorf("artifact digest mismatch: %s != %s", actual, digest)
	}
	path, err := c.objectPath(namespace, digest)
	if err != nil {
		return err
	}
	if _, statErr := os.Stat(path); statErr == nil {
		return nil // already stored
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

// Get returns cached artifact bytes, re-verifying the digest.
func (c *Cache) Get(namespace, digest string) ([]byte, error) {
	path, err := c.objectPath(namespace, digest)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("artifact %s/%s not cached: %w", namespace, digest, err)
	}
	if digestOf(data) != digest {
		return nil, fmt.Errorf("cached artifact %s/%s is corrupt", namespace, digest)
	}
	return data, nil
}
