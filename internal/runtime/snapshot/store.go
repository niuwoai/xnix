// Package snapshot implements a constrained, content-addressed snapshot store
// for Runtime-owned compatibility state. Every operation is confined to a
// caller-provided state root: the store creates, lists, verifies, and rolls
// back snapshots of that root and refuses any path that escapes it. It never
// touches the host root, uses only ordinary user-writable directories, and
// performs no privileged operations.
//
// Snapshots are content-addressed: each captured file is hashed with SHA-256
// and stored once under an objects directory, so identical content is not
// duplicated across snapshots. Snapshot identifiers are caller-provided to keep
// behaviour deterministic and testable.
package snapshot

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/record"
	"xnix.local/xnix/internal/runtime/rootfs"
)

// metaDirName is the store's private subdirectory inside the state root. It is
// excluded from captured state so snapshots never recurse into the store.
const metaDirName = ".xnix-snapshots"

// Store manages snapshots of a single state root directory.
type Store struct {
	root    string       // absolute path to the state root
	rootObj *rootfs.Root // controlled-root handle for the state root
	metaDir string       // absolute path to the store's private metadata dir
}

// FileEntry records one captured file: its path relative to the state root and
// the SHA-256 digest of its content.
type FileEntry struct {
	Path   string `json:"path"`
	Digest string `json:"digest"`
	Mode   uint32 `json:"mode"`
}

// Manifest is the content-addressed description of a snapshot.
type Manifest struct {
	ID          string      `json:"id"`
	Reason      string      `json:"reason"`
	FileCount   int         `json:"file_count"`
	Files       []FileEntry `json:"files"`
	ContentHash string      `json:"content_hash"`
}

// RollbackReceipt records the effect of a rollback for audit and reversal.
type RollbackReceipt struct {
	SnapshotID      string   `json:"snapshot_id"`
	Restored        []string `json:"restored"`
	Removed         []string `json:"removed"`
	HostRootTouched bool     `json:"host_root_touched"`
}

// New opens (creating if needed) a snapshot store rooted at the given state
// root. The root must be an existing directory; the store keeps its metadata in
// a private subdirectory of the root.
func New(stateRoot string) (*Store, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return nil, err
	}
	abs := root.Path()
	metaDir := filepath.Join(abs, metaDirName)
	if err := os.MkdirAll(filepath.Join(metaDir, "objects"), 0o700); err != nil {
		return nil, fmt.Errorf("initialize snapshot store: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(metaDir, "manifests"), 0o700); err != nil {
		return nil, fmt.Errorf("initialize snapshot store: %w", err)
	}
	return &Store{root: abs, rootObj: root, metaDir: metaDir}, nil
}

// Root returns the absolute state-root path the store is confined to.
func (s *Store) Root() string { return s.root }

// manifestRel is the state-root-relative path of a snapshot manifest.
func (s *Store) manifestRel(id string) string {
	return filepath.ToSlash(filepath.Join(metaDirName, "manifests", id+".json"))
}

// within reports whether abs is inside the state root (and not the store's own
// metadata directory). It is the guard that keeps every operation sandboxed.
func (s *Store) within(abs string) bool {
	return s.rootObj.Contains(abs)
}

func validSnapshotID(id string) bool {
	if id == "" || len(id) > 128 {
		return false
	}
	for _, r := range id {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_', r == '.':
		default:
			return false
		}
	}
	// Disallow ids that could traverse the manifest directory.
	return id != "." && id != ".." && !strings.Contains(id, string(os.PathSeparator))
}

// Create captures the current state root into a new snapshot with the given id.
// The store's own metadata directory is never captured.
func (s *Store) Create(id string, reason string) (Manifest, error) {
	if !validSnapshotID(id) {
		return Manifest{}, fmt.Errorf("invalid snapshot id %q", id)
	}
	if _, err := os.Stat(s.manifestPath(id)); err == nil {
		return Manifest{}, fmt.Errorf("snapshot %q already exists", id)
	}

	var files []FileEntry
	walkErr := filepath.WalkDir(s.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == s.metaDir {
			return filepath.SkipDir
		}
		if d.IsDir() {
			return nil
		}
		if !d.Type().IsRegular() {
			// Skip symlinks/devices: the store only snapshots regular files and
			// never follows links out of the state root.
			return nil
		}
		if !s.within(path) {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		digest := hashBytes(data)
		if writeErr := s.writeObject(digest, data); writeErr != nil {
			return writeErr
		}
		rel, relErr := filepath.Rel(s.root, path)
		if relErr != nil {
			return relErr
		}
		info, infoErr := d.Info()
		if infoErr != nil {
			return infoErr
		}
		files = append(files, FileEntry{Path: filepath.ToSlash(rel), Digest: digest, Mode: uint32(info.Mode().Perm())})
		return nil
	})
	if walkErr != nil {
		return Manifest{}, fmt.Errorf("capture state root: %w", walkErr)
	}

	sort.Slice(files, func(i, j int) bool { return files[i].Path < files[j].Path })
	manifest := Manifest{
		ID:          id,
		Reason:      reason,
		FileCount:   len(files),
		Files:       files,
		ContentHash: manifestContentHash(files),
	}
	if err := s.writeManifest(manifest); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

// List returns all snapshot manifests sorted by id.
func (s *Store) List() ([]Manifest, error) {
	entries, err := os.ReadDir(filepath.Join(s.metaDir, "manifests"))
	if err != nil {
		return nil, fmt.Errorf("list snapshots: %w", err)
	}
	var out []Manifest
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), ".json")
		manifest, loadErr := s.load(id)
		if loadErr != nil {
			return nil, loadErr
		}
		out = append(out, manifest)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// Verify checks that every object referenced by the snapshot exists and hashes
// to its recorded digest, and that the manifest content hash is intact.
func (s *Store) Verify(id string) error {
	manifest, err := s.load(id)
	if err != nil {
		return err
	}
	if manifest.ContentHash != manifestContentHash(manifest.Files) {
		return fmt.Errorf("snapshot %q manifest content hash mismatch", id)
	}
	for _, file := range manifest.Files {
		data, readErr := os.ReadFile(s.objectPath(file.Digest))
		if readErr != nil {
			return fmt.Errorf("snapshot %q missing object for %s: %w", id, file.Path, readErr)
		}
		if hashBytes(data) != file.Digest {
			return fmt.Errorf("snapshot %q object for %s is corrupt", id, file.Path)
		}
	}
	return nil
}

// Restore rolls the state root back to the snapshot: it rewrites every recorded
// file and removes captured files that are not part of the snapshot. Files the
// store never captured (e.g. the metadata dir) are left untouched. It returns a
// receipt and never writes outside the state root.
func (s *Store) Restore(id string) (RollbackReceipt, error) {
	if err := s.Verify(id); err != nil {
		return RollbackReceipt{}, err
	}
	manifest, err := s.load(id)
	if err != nil {
		return RollbackReceipt{}, err
	}

	want := make(map[string]FileEntry, len(manifest.Files))
	for _, file := range manifest.Files {
		want[file.Path] = file
	}

	receipt := RollbackReceipt{SnapshotID: id, HostRootTouched: false}

	// Remove currently-captured files that the snapshot does not contain.
	current, err := s.currentFiles()
	if err != nil {
		return RollbackReceipt{}, err
	}
	for _, rel := range current {
		if _, keep := want[rel]; keep {
			continue
		}
		abs := filepath.Join(s.root, filepath.FromSlash(rel))
		if !s.within(abs) {
			continue
		}
		if err := os.Remove(abs); err != nil {
			return RollbackReceipt{}, fmt.Errorf("remove %s: %w", rel, err)
		}
		receipt.Removed = append(receipt.Removed, rel)
	}

	// Restore every recorded file from content-addressed objects.
	for _, file := range manifest.Files {
		abs := filepath.Join(s.root, filepath.FromSlash(file.Path))
		if !s.within(abs) {
			return RollbackReceipt{}, fmt.Errorf("snapshot %q path escapes state root: %s", id, file.Path)
		}
		if err := os.MkdirAll(filepath.Dir(abs), 0o700); err != nil {
			return RollbackReceipt{}, err
		}
		data, readErr := os.ReadFile(s.objectPath(file.Digest))
		if readErr != nil {
			return RollbackReceipt{}, readErr
		}
		mode := fs.FileMode(file.Mode).Perm()
		if mode == 0 {
			mode = 0o600
		}
		if err := os.WriteFile(abs, data, mode); err != nil {
			return RollbackReceipt{}, fmt.Errorf("restore %s: %w", file.Path, err)
		}
		receipt.Restored = append(receipt.Restored, file.Path)
	}
	sort.Strings(receipt.Restored)
	sort.Strings(receipt.Removed)
	return receipt, nil
}

// BaselineStatus reports whether a verified restore point exists, for repair
// and execution preflight to require before a risky compatibility change. It
// performs no restore and exposes no host paths or backend details.
type BaselineStatus struct {
	Present               bool   `json:"present"`
	Verified              bool   `json:"verified"`
	SnapshotID            string `json:"snapshot_id,omitempty"`
	SnapshotCount         int    `json:"snapshot_count"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

// Baseline reports the restore-point baseline: whether the store holds at least
// one snapshot that verifies intact. It is what repair and execution preflight
// consult to require a restore point, without performing a restore. The most
// recent verifying snapshot (by id order) is reported as the baseline.
func (s *Store) Baseline() (BaselineStatus, error) {
	manifests, err := s.List()
	if err != nil {
		return BaselineStatus{}, err
	}
	status := BaselineStatus{SnapshotCount: len(manifests)}
	for i := len(manifests) - 1; i >= 0; i-- {
		if s.Verify(manifests[i].ID) == nil {
			status.Present = true
			status.Verified = true
			status.SnapshotID = manifests[i].ID
			break
		}
	}
	switch {
	case status.Verified:
		status.Summary = "A verified restore point is available before risky compatibility changes."
	case status.SnapshotCount > 0:
		status.Summary = "Restore points exist but none verify intact; create a new restore point."
	default:
		status.Summary = "No restore point exists yet; create one before risky compatibility changes."
	}
	return status, nil
}

// currentFiles lists regular files under the state root (excluding store meta),
// as slash-separated paths relative to the root.
func (s *Store) currentFiles() ([]string, error) {
	var out []string
	err := filepath.WalkDir(s.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == s.metaDir {
			return filepath.SkipDir
		}
		if d.IsDir() || !d.Type().IsRegular() {
			return nil
		}
		rel, relErr := filepath.Rel(s.root, path)
		if relErr != nil {
			return relErr
		}
		out = append(out, filepath.ToSlash(rel))
		return nil
	})
	return out, err
}

func (s *Store) manifestPath(id string) string {
	return filepath.Join(s.metaDir, "manifests", id+".json")
}

func (s *Store) objectPath(digest string) string {
	return filepath.Join(s.metaDir, "objects", digest)
}

func (s *Store) writeObject(digest string, data []byte) error {
	path := s.objectPath(digest)
	if _, err := os.Stat(path); err == nil {
		return nil // content-addressed: already stored
	}
	return os.WriteFile(path, data, 0o600)
}

func (s *Store) writeManifest(manifest Manifest) error {
	_, err := record.Save(s.rootObj, s.manifestRel(manifest.ID), manifest)
	return err
}

func (s *Store) load(id string) (Manifest, error) {
	if !validSnapshotID(id) {
		return Manifest{}, fmt.Errorf("invalid snapshot id %q", id)
	}
	var manifest Manifest
	if err := record.Load(s.rootObj, s.manifestRel(id), &manifest); err != nil {
		return Manifest{}, fmt.Errorf("snapshot %q: %w", id, err)
	}
	return manifest, nil
}

func hashBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func manifestContentHash(files []FileEntry) string {
	h := sha256.New()
	for _, file := range files {
		h.Write([]byte(file.Path))
		h.Write([]byte{0})
		h.Write([]byte(file.Digest))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))
}
