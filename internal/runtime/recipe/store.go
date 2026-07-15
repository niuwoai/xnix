package recipe

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// applicationIDPattern matches reverse-DNS application identifiers.
var applicationIDPattern = regexp.MustCompile(`^[A-Za-z0-9]+(\.[A-Za-z0-9-]+)+$`)

// Registry is the signed manifest describing a recipe root's contents.
type Registry struct {
	SchemaVersion int     `json:"schema_version"`
	RegistryName  string  `json:"registry_name"`
	Recipes       []Entry `json:"recipes"`
}

// Entry is one registry record: the recipe id, its path relative to the recipe
// root, its expected content digest, and its declared signature status.
type Entry struct {
	ID              string          `json:"id"`
	Path            string          `json:"path"`
	SHA256          string          `json:"sha256"`
	SignatureStatus SignatureStatus `json:"signature_status"`
}

// Recipe is the loaded recipe together with its verified trust state and the
// registry it came from. The raw recipe fields are kept as generic data so this
// storage boundary does not couple to the full recipe schema.
type Recipe struct {
	ID           string          `json:"id"`
	RegistryName string          `json:"registry_name"`
	Trust        TrustState      `json:"trust"`
	Data         json.RawMessage `json:"data"`
}

// Store is the durable recipe storage boundary. Implementations are read-only:
// they never write recipes and never mutate the host root.
type Store interface {
	// Find loads and verifies a single recipe by application id.
	Find(applicationID string) (Recipe, error)
	// List loads and verifies every registry recipe, sorted by id.
	List() ([]Recipe, error)
	// TrustStates returns the trust state of every registry recipe without
	// requiring callers to hold recipe bodies, for trust diagnostics.
	TrustStates() ([]TrustState, error)
	// RegistryName is the name of the backing registry.
	RegistryName() (string, error)
	// Development reports whether the backing registry is a non-production
	// development registry.
	Development() (bool, error)
}

// LocalStore is a read-only Store backed by a local recipe root and registry
// file. It resolves recipe paths safely inside the root and verifies every
// recipe through the configured Verifier.
type LocalStore struct {
	root         string
	registryPath string
	verifier     Verifier
}

// developmentRegistryNames marks registries that are explicitly non-production.
func isDevelopmentRegistry(name string) bool {
	lower := strings.ToLower(name)
	return strings.Contains(lower, "development") || strings.Contains(lower, "local") || strings.Contains(lower, "test")
}

// NewLocalStore opens a read-only recipe store. The recipe root must contain a
// registry.json (or registryFile if named). A nil verifier defaults to a
// fail-closed DigestVerifier.
func NewLocalStore(root string, verifier Verifier) (*LocalStore, error) {
	if root == "" {
		return nil, errors.New("recipe store requires a recipe root")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolve recipe root: %w", err)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("recipe root must exist: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("recipe root %q is not a directory", abs)
	}
	if verifier == nil {
		verifier = DigestVerifier{}
	}
	return &LocalStore{root: abs, registryPath: filepath.Join(abs, "registry.json"), verifier: verifier}, nil
}

func (s *LocalStore) loadRegistry() (Registry, error) {
	data, err := os.ReadFile(s.registryPath)
	if err != nil {
		return Registry{}, fmt.Errorf("read registry: %w", err)
	}
	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return Registry{}, fmt.Errorf("parse registry: %w", err)
	}
	if registry.SchemaVersion != 1 {
		return Registry{}, fmt.Errorf("unsupported registry schema version %d", registry.SchemaVersion)
	}
	seen := map[string]bool{}
	for _, entry := range registry.Recipes {
		if !applicationIDPattern.MatchString(entry.ID) {
			return Registry{}, fmt.Errorf("registry entry id %q is not a reverse-DNS identifier", entry.ID)
		}
		if seen[entry.ID] {
			return Registry{}, fmt.Errorf("registry has duplicate entry for %q", entry.ID)
		}
		seen[entry.ID] = true
	}
	return registry, nil
}

// safeRecipePath resolves entryPath inside the recipe root and refuses escapes.
func (s *LocalStore) safeRecipePath(entryPath string) (string, error) {
	if entryPath == "" || filepath.IsAbs(entryPath) {
		return "", fmt.Errorf("registry recipe path must be relative: %q", entryPath)
	}
	joined := filepath.Join(s.root, filepath.FromSlash(entryPath))
	rel, err := filepath.Rel(s.root, joined)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("registry recipe path escapes recipe root: %q", entryPath)
	}
	return joined, nil
}

func (s *LocalStore) loadEntry(entry Entry) (Recipe, error) {
	recipePath, err := s.safeRecipePath(entry.Path)
	if err != nil {
		return Recipe{}, err
	}
	data, err := os.ReadFile(recipePath)
	if err != nil {
		return Recipe{}, fmt.Errorf("read recipe %s: %w", entry.ID, err)
	}
	trust := s.verifier.Verify(entry, data)
	if !trust.Usable() {
		return Recipe{}, fmt.Errorf("recipe %s failed trust verification: %s", entry.ID, trust.Reason)
	}
	registryName, _ := s.RegistryName()
	return Recipe{ID: entry.ID, RegistryName: registryName, Trust: trust, Data: json.RawMessage(append([]byte(nil), data...))}, nil
}

// Find loads and verifies a single recipe by application id.
func (s *LocalStore) Find(applicationID string) (Recipe, error) {
	if !applicationIDPattern.MatchString(applicationID) {
		return Recipe{}, errors.New("application id must be a reverse-DNS identifier")
	}
	registry, err := s.loadRegistry()
	if err != nil {
		return Recipe{}, err
	}
	for _, entry := range registry.Recipes {
		if entry.ID == applicationID {
			return s.loadEntry(entry)
		}
	}
	return Recipe{}, fmt.Errorf("unknown recipe %q", applicationID)
}

// List loads and verifies every registry recipe, sorted by id.
func (s *LocalStore) List() ([]Recipe, error) {
	registry, err := s.loadRegistry()
	if err != nil {
		return nil, err
	}
	out := make([]Recipe, 0, len(registry.Recipes))
	for _, entry := range registry.Recipes {
		recipe, loadErr := s.loadEntry(entry)
		if loadErr != nil {
			return nil, loadErr
		}
		out = append(out, recipe)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// TrustStates returns the trust state of every registry recipe, sorted by id.
// Unlike List, a fail-closed recipe is reported (not an error) so trust
// diagnostics can surface why a recipe is not usable.
func (s *LocalStore) TrustStates() ([]TrustState, error) {
	registry, err := s.loadRegistry()
	if err != nil {
		return nil, err
	}
	out := make([]TrustState, 0, len(registry.Recipes))
	for _, entry := range registry.Recipes {
		recipePath, pathErr := s.safeRecipePath(entry.Path)
		if pathErr != nil {
			out = append(out, TrustState{ID: entry.ID, FailedClosed: true, Reason: pathErr.Error()})
			continue
		}
		data, readErr := os.ReadFile(recipePath)
		if readErr != nil {
			out = append(out, TrustState{ID: entry.ID, FailedClosed: true, Reason: readErr.Error()})
			continue
		}
		out = append(out, s.verifier.Verify(entry, data))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// RegistryName returns the backing registry's declared name.
func (s *LocalStore) RegistryName() (string, error) {
	registry, err := s.loadRegistry()
	if err != nil {
		return "", err
	}
	return registry.RegistryName, nil
}

// Development reports whether the backing registry is non-production.
func (s *LocalStore) Development() (bool, error) {
	name, err := s.RegistryName()
	if err != nil {
		return false, err
	}
	return isDevelopmentRegistry(name), nil
}

// compile-time assertion that LocalStore satisfies Store.
var _ Store = (*LocalStore)(nil)
