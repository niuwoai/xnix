// Package record is the canonical persistence primitive for Runtime JSON
// records. Every Runtime store previously encoded a value with
// json.MarshalIndent and wrote it with os.WriteFile — a non-atomic write that
// leaves a truncated, corrupt file if the process is interrupted mid-write.
// None of the stores used a temp-and-rename, so this package adds crash-safe
// atomic writes while consolidating the encode + digest + write idiom.
//
// Records are confined to a rootfs.Root, so a record path can never escape its
// controlled state root.
package record

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/rootfs"
)

// tempSuffix marks the in-flight file of an atomic write.
const tempSuffix = ".writing.tmp"

// Encode returns the deterministic indented JSON encoding of v. The same value
// always encodes to the same bytes, so digests are stable.
func Encode(v any) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encode record: %w", err)
	}
	return data, nil
}

// Digest returns the lowercase hex SHA-256 of a value's canonical encoding.
func Digest(v any) (string, error) {
	data, err := Encode(v)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// DigestBytes returns the lowercase hex SHA-256 of raw bytes.
func DigestBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// Save atomically writes v as JSON to a root-relative path. It encodes v, writes
// the bytes to a temporary file in the same directory, then renames it over the
// target so a reader never observes a partial record. Parent directories are
// created inside the root. The returned digest is the SHA-256 of the written
// bytes.
func Save(root *rootfs.Root, rel string, v any) (string, error) {
	if root == nil {
		return "", fmt.Errorf("record save requires a controlled root")
	}
	data, err := Encode(v)
	if err != nil {
		return "", err
	}
	target, err := root.Resolve(rel)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o700); err != nil {
		return "", err
	}
	temp := target + tempSuffix
	if err := os.WriteFile(temp, data, 0o600); err != nil {
		return "", fmt.Errorf("write record: %w", err)
	}
	if err := os.Rename(temp, target); err != nil {
		// Best-effort cleanup so a failed write does not leave a temp file.
		_ = os.Remove(temp)
		return "", fmt.Errorf("commit record: %w", err)
	}
	return DigestBytes(data), nil
}

// Load reads and decodes a JSON record from a root-relative path into v.
func Load(root *rootfs.Root, rel string, v any) error {
	if root == nil {
		return fmt.Errorf("record load requires a controlled root")
	}
	data, err := root.ReadFile(rel)
	if err != nil {
		return fmt.Errorf("read record %q: %w", rel, err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("record %q is corrupt: %w", rel, err)
	}
	return nil
}

// Exists reports whether a record exists at a root-relative path.
func Exists(root *rootfs.Root, rel string) (bool, error) {
	if root == nil {
		return false, fmt.Errorf("record exists check requires a controlled root")
	}
	abs, err := root.Resolve(rel)
	if err != nil {
		return false, err
	}
	if _, statErr := os.Stat(abs); statErr != nil {
		if os.IsNotExist(statErr) {
			return false, nil
		}
		return false, statErr
	}
	return true, nil
}
