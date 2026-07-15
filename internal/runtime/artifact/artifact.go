// Package artifact implements a safe, reviewable compatibility-artifact
// acquisition pipeline. It parses an artifact manifest, plans a content-
// addressed cache, verifies digests, and stages artifacts into a controlled
// cache root — never the host root. Network fetching is disabled by default:
// only an explicitly supplied Source (e.g. a local fixture source) can provide
// bytes, and a digest mismatch always blocks staging.
package artifact

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
)

// applicationIDPattern matches reverse-DNS application identifiers.
var applicationIDPattern = regexp.MustCompile(`^[A-Za-z0-9]+(\.[A-Za-z0-9-]+)+$`)

// digestPattern matches a lowercase hex SHA-256 digest.
var digestPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

// Ref is a single required or optional artifact identified by content digest.
type Ref struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"`
	SHA256   string `json:"sha256"`
	Size     int64  `json:"size"`
	Required bool   `json:"required"`
}

// Group is a named collection of artifacts sharing a cache namespace.
type Group struct {
	ID   string `json:"id"`
	Refs []Ref  `json:"refs"`
}

// Manifest describes the artifacts an application needs.
type Manifest struct {
	ApplicationID string  `json:"application_id"`
	Groups        []Group `json:"groups"`
}

// ParseManifest parses and validates an artifact manifest.
func ParseManifest(data []byte) (Manifest, error) {
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, fmt.Errorf("parse artifact manifest: %w", err)
	}
	if err := manifest.validate(); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func (m Manifest) validate() error {
	if !applicationIDPattern.MatchString(m.ApplicationID) {
		return errors.New("artifact manifest application id must be a reverse-DNS identifier")
	}
	if len(m.Groups) == 0 {
		return errors.New("artifact manifest must declare at least one group")
	}
	seenGroups := map[string]bool{}
	for _, group := range m.Groups {
		if group.ID == "" {
			return errors.New("artifact group id must not be empty")
		}
		if seenGroups[group.ID] {
			return fmt.Errorf("duplicate artifact group %q", group.ID)
		}
		seenGroups[group.ID] = true
		seenRefs := map[string]bool{}
		for _, ref := range group.Refs {
			if ref.ID == "" {
				return fmt.Errorf("artifact ref id must not be empty in group %q", group.ID)
			}
			if seenRefs[ref.ID] {
				return fmt.Errorf("duplicate artifact ref %q in group %q", ref.ID, group.ID)
			}
			seenRefs[ref.ID] = true
			if !digestPattern.MatchString(ref.SHA256) {
				return fmt.Errorf("artifact ref %q must declare a lowercase hex sha256", ref.ID)
			}
			if ref.Size < 0 {
				return fmt.Errorf("artifact ref %q size must not be negative", ref.ID)
			}
		}
	}
	return nil
}

// CacheNamespace returns the sanitized per-application namespace for a group.
func (m Manifest) CacheNamespace(groupID string) string {
	return sanitizeNamespace(m.ApplicationID) + "." + sanitizeNamespace(groupID)
}

// CacheKey is the content-addressed key an artifact would occupy in the cache.
func CacheKey(ref Ref) string {
	return ref.SHA256
}

func sanitizeNamespace(value string) string {
	sanitized := make([]rune, 0, len(value))
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-':
			sanitized = append(sanitized, r)
		default:
			sanitized = append(sanitized, '-')
		}
	}
	return string(sanitized)
}

func digestOf(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// allRefs returns every ref across all groups, sorted by group then ref id.
func (m Manifest) allRefs() []groupedRef {
	var refs []groupedRef
	for _, group := range m.Groups {
		for _, ref := range group.Refs {
			refs = append(refs, groupedRef{group: group.ID, namespace: m.CacheNamespace(group.ID), ref: ref})
		}
	}
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].group != refs[j].group {
			return refs[i].group < refs[j].group
		}
		return refs[i].ref.ID < refs[j].ref.ID
	})
	return refs
}

type groupedRef struct {
	group     string
	namespace string
	ref       Ref
}
