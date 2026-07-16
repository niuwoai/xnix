// Package preview provides the canonical, embeddable safety flags and header
// that Runtime preview and record payloads share. The same cluster of safety
// flags (runtime_owned, kde_policy_owner, host_root_modified,
// backend_details_exposed, ...) is declared in ~70 structs across the codebase;
// this package lets a payload embed them once and, more importantly, assert the
// Runtime safety invariant with a single reusable check.
//
// It pairs with internal/runtime/safety: safety validates payload *content*
// (no host paths, backend terms, or secrets in text), while this package
// validates payload *flags* (the payload never claims an unsafe capability).
package preview

import (
	"fmt"
	"regexp"
	"strings"
)

// SafetyFlags is the near-universal cluster of Runtime safety flags. Embed it in
// a preview or record struct to declare them once; JSON marshaling flattens the
// embedded fields to the top level with their canonical tags.
type SafetyFlags struct {
	RuntimeOwned                bool `json:"runtime_owned"`
	GoRuntimeBacked             bool `json:"go_runtime_backed"`
	KDEPolicyOwner              bool `json:"kde_policy_owner"`
	HostRootModified            bool `json:"host_root_modified"`
	NetworkRequired             bool `json:"network_required"`
	PrivilegedContainerRequired bool `json:"privileged_container_required"`
	BackendLaunchEnabled        bool `json:"backend_launch_enabled"`
	BackendDetailsExposed       bool `json:"backend_details_exposed"`
}

// Safe returns SafetyFlags with the Runtime-owned Go defaults: the Runtime owns
// the payload and it is Go-backed, KDE is not the policy owner, and every unsafe
// capability flag is disabled.
func Safe() SafetyFlags {
	return SafetyFlags{RuntimeOwned: true, GoRuntimeBacked: true}
}

// unsafeInvariants are the flags that must always be false in preview/gated
// mode: KDE never owns Runtime policy, and no preview may claim to have mutated
// the host, required the network, needed a privileged container, launched a
// backend, or exposed backend details.
func (f SafetyFlags) unsafeViolations() []string {
	var bad []string
	if f.KDEPolicyOwner {
		bad = append(bad, "kde_policy_owner")
	}
	if f.HostRootModified {
		bad = append(bad, "host_root_modified")
	}
	if f.NetworkRequired {
		bad = append(bad, "network_required")
	}
	if f.PrivilegedContainerRequired {
		bad = append(bad, "privileged_container_required")
	}
	if f.BackendLaunchEnabled {
		bad = append(bad, "backend_launch_enabled")
	}
	if f.BackendDetailsExposed {
		bad = append(bad, "backend_details_exposed")
	}
	return bad
}

// Validate asserts the Runtime safety invariant: the Runtime must own the
// payload and every unsafe capability flag must be false. It returns an error
// naming each violated flag.
func (f SafetyFlags) Validate() error {
	if !f.RuntimeOwned {
		return fmt.Errorf("safety invariant violated: runtime_owned must be true")
	}
	if bad := f.unsafeViolations(); len(bad) > 0 {
		return fmt.Errorf("safety invariant violated: these flags must be false: %s", strings.Join(bad, ", "))
	}
	return nil
}

// Safe reports whether the flags satisfy the Runtime safety invariant.
func (f SafetyFlags) Safe() bool { return f.Validate() == nil }

// schemaVersionPattern matches a Runtime schema version like
// "xnix.runtime.launch_readiness.v1".
var schemaVersionPattern = regexp.MustCompile(`^xnix\.runtime\.[a-z0-9_]+\.v[0-9]+$`)

// Header is the common identity header shared by Runtime preview payloads.
type Header struct {
	SchemaVersion string `json:"schema_version"`
	RequestType   string `json:"request_type"`
	RuntimeMethod string `json:"runtime_method"`
	ReadMethod    string `json:"read_method,omitempty"`
	Desktop       string `json:"desktop,omitempty"`
}

// Validate checks that the header identity fields are well formed: the schema
// version follows the Runtime scheme, request type and runtime method are set,
// and the desktop (when present) is the supported KDE Plasma target.
func (h Header) Validate() error {
	if !schemaVersionPattern.MatchString(h.SchemaVersion) {
		return fmt.Errorf("schema_version %q must match xnix.runtime.<domain>.v<N>", h.SchemaVersion)
	}
	if h.RequestType == "" {
		return fmt.Errorf("request_type must not be empty")
	}
	if h.RuntimeMethod == "" {
		return fmt.Errorf("runtime_method must not be empty")
	}
	if h.Desktop != "" && h.Desktop != "KDE Plasma" {
		return fmt.Errorf("desktop %q is not the supported KDE Plasma target", h.Desktop)
	}
	return nil
}
