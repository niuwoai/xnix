package artifact

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// ErrNetworkDisabled is returned by network-backed sources, which are disabled
// by default. Only an explicitly supplied fixture Source can provide bytes.
var ErrNetworkDisabled = errors.New("artifact network fetch is disabled")

// Source supplies artifact bytes for a ref. Implementations must return content
// whose digest matches the ref; the cache verifies before staging regardless.
type Source interface {
	Fetch(namespace string, ref Ref) ([]byte, error)
}

// LocalFixtureSource reads artifacts from a local fixture directory, addressed
// by content digest. It performs no network access.
type LocalFixtureSource struct {
	dir string
}

// NewLocalFixtureSource opens a fixture source rooted at dir.
func NewLocalFixtureSource(dir string) (*LocalFixtureSource, error) {
	if dir == "" {
		return nil, errors.New("fixture source requires a directory")
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(abs)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("fixture source directory must exist: %s", dir)
	}
	return &LocalFixtureSource{dir: abs}, nil
}

// Fetch reads the fixture file named by the ref digest.
func (s *LocalFixtureSource) Fetch(_ string, ref Ref) ([]byte, error) {
	if !digestPattern.MatchString(ref.SHA256) {
		return nil, fmt.Errorf("artifact ref %q has an invalid digest", ref.ID)
	}
	path := filepath.Join(s.dir, ref.SHA256)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fixture artifact %s not found: %w", ref.SHA256, err)
	}
	return data, nil
}

// networkSource is the disabled-by-default network transport.
type networkSource struct{}

// NewNetworkSource returns a network Source that is disabled: every Fetch
// returns ErrNetworkDisabled until a real, explicitly enabled transport exists.
func NewNetworkSource() Source { return networkSource{} }

func (networkSource) Fetch(string, Ref) ([]byte, error) { return nil, ErrNetworkDisabled }

// PlannedArtifact is one artifact's place in an acquisition plan.
type PlannedArtifact struct {
	Group          string `json:"group"`
	RefID          string `json:"ref_id"`
	CacheNamespace string `json:"cache_namespace"`
	CacheKey       string `json:"cache_key"`
	Required       bool   `json:"required"`
	AlreadyCached  bool   `json:"already_cached"`
}

// Plan is the reviewable outcome of acquisition: which artifacts would be
// fetched and where they would be cached. It performs no writes.
type Plan struct {
	ApplicationID   string            `json:"application_id"`
	DryRun          bool              `json:"dry_run"`
	NetworkEnabled  bool              `json:"network_enabled"`
	HostRootMutated bool              `json:"host_root_mutated"`
	Artifacts       []PlannedArtifact `json:"artifacts"`
	StagedKeys      []string          `json:"staged_keys"`
	RequiredCount   int               `json:"required_count"`
	OptionalCount   int               `json:"optional_count"`
}

// Acquirer plans and (optionally) stages artifacts into a cache from a Source.
type Acquirer struct {
	cache  *Cache
	source Source
}

// NewAcquirer builds an acquirer over a cache and source. A nil source yields a
// plan-only acquirer that can dry-run but never stage.
func NewAcquirer(cache *Cache, source Source) (*Acquirer, error) {
	if cache == nil {
		return nil, errors.New("acquirer requires a cache")
	}
	return &Acquirer{cache: cache, source: source}, nil
}

// Plan reports the acquisition plan for a manifest without fetching or writing.
func (a *Acquirer) Plan(manifest Manifest) Plan {
	plan := Plan{ApplicationID: manifest.ApplicationID, DryRun: true, NetworkEnabled: false, HostRootMutated: false}
	for _, gr := range manifest.allRefs() {
		planned := PlannedArtifact{
			Group:          gr.group,
			RefID:          gr.ref.ID,
			CacheNamespace: gr.namespace,
			CacheKey:       CacheKey(gr.ref),
			Required:       gr.ref.Required,
			AlreadyCached:  a.cache.Has(gr.namespace, gr.ref.SHA256),
		}
		if planned.Required {
			plan.RequiredCount++
		} else {
			plan.OptionalCount++
		}
		plan.Artifacts = append(plan.Artifacts, planned)
	}
	return plan
}

// Acquire stages every artifact into the cache from the source, verifying each
// digest before it lands. A digest mismatch blocks staging for that artifact
// and returns an error. Host root mutation never occurs: all writes are inside
// the cache root. Acquire requires a source.
func (a *Acquirer) Acquire(manifest Manifest) (Plan, error) {
	plan := a.Plan(manifest)
	plan.DryRun = false
	if a.source == nil {
		return plan, errors.New("acquire requires a source; use Plan for a dry run")
	}

	staged := map[string]bool{}
	for _, gr := range manifest.allRefs() {
		if a.cache.Has(gr.namespace, gr.ref.SHA256) {
			staged[gr.ref.SHA256] = true
			continue
		}
		data, err := a.source.Fetch(gr.namespace, gr.ref)
		if err != nil {
			return plan, fmt.Errorf("fetch artifact %s/%s: %w", gr.group, gr.ref.ID, err)
		}
		// Put verifies the digest and blocks staging on mismatch.
		if err := a.cache.Put(gr.namespace, gr.ref.SHA256, data); err != nil {
			return plan, fmt.Errorf("stage artifact %s/%s: %w", gr.group, gr.ref.ID, err)
		}
		staged[gr.ref.SHA256] = true
	}

	keys := make([]string, 0, len(staged))
	for key := range staged {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	plan.StagedKeys = keys
	return plan, nil
}
