// Package portal implements a Runtime-owned XDG Desktop Portal request broker
// for preview and container-test mode. It tracks portal request objects and
// their completion state behind a Broker interface so a real Portal transport
// can be plugged in later, while a fake in-memory broker serves tests.
//
// The broker never performs real Portal calls, never exposes host paths or
// backend details to callers, and tracks portal permission state separately
// from any downstream execution approval.
package portal

import (
	"errors"
	"fmt"
	"strings"

	"xnix.local/xnix/internal/runtime/appid"
)

// PortalDestination is the well-known XDG Desktop Portal bus name.
const PortalDestination = "org.freedesktop.portal.Desktop"

// PortalObjectPath is the well-known XDG Desktop Portal object path.
const PortalObjectPath = "/org/freedesktop/portal/desktop"

// RequestState is the lifecycle state of a portal request object.
type RequestState string

const (
	// StatePendingUserMediation means the request awaits a user-mediated
	// Portal response and no permission has been granted yet.
	StatePendingUserMediation RequestState = "pending-user-mediation"
	// StateGranted means the user approved the request through the Portal.
	StateGranted RequestState = "granted"
	// StateDenied means the request was denied, either by policy at creation
	// time or by the user through the Portal.
	StateDenied RequestState = "denied"
	// StateCancelled means the request was withdrawn before completion.
	StateCancelled RequestState = "cancelled"
	// StateFailed means the Portal transport failed; the request is
	// recoverable and can be retried.
	StateFailed RequestState = "failed"
	// StateExpired means the request timed out before a user response and can
	// no longer be resolved.
	StateExpired RequestState = "expired"
	// StateCompleted means a granted request was consumed by the caller.
	StateCompleted RequestState = "completed"
)

// PermissionState tracks the portal permission decision independently of any
// downstream execution or launch approval.
type PermissionState string

const (
	PermissionPending    PermissionState = "pending"
	PermissionGranted    PermissionState = "granted"
	PermissionDenied     PermissionState = "denied"
	PermissionNotGranted PermissionState = "not-granted"
)

// operation describes a sensitive desktop operation mediated by the Portal.
type operation struct {
	iface     string
	method    string
	decision  string
	resources []string
	summary   string
}

var operationCatalog = map[string]operation{
	"file-open": {
		iface: "org.freedesktop.portal.FileChooser", method: "OpenFile", decision: "ask",
		resources: []string{"documents", "downloads", "selected-files"},
		summary:   "File access requires a user-approved desktop portal request.",
	},
	"uri-open": {
		iface: "org.freedesktop.portal.OpenURI", method: "OpenURI", decision: "ask",
		resources: []string{"external-uri"},
		summary:   "URI handling requires a user-approved desktop portal request.",
	},
	"print": {
		iface: "org.freedesktop.portal.Print", method: "Print", decision: "ask",
		resources: []string{"printer"},
		summary:   "Printing requires a user-approved desktop portal request.",
	},
	"screenshot": {
		iface: "org.freedesktop.portal.Screenshot", method: "Screenshot", decision: "ask",
		resources: []string{"screen"},
		summary:   "Screenshots require a user-approved desktop portal request.",
	},
	"clipboard": {
		iface: "org.freedesktop.portal.Clipboard", method: "RequestClipboard", decision: "ask",
		resources: []string{"clipboard"},
		summary:   "Clipboard access requires a user-approved desktop portal request.",
	},
	"camera": {
		iface: "org.freedesktop.portal.Camera", method: "AccessCamera", decision: "deny",
		resources: []string{"camera"},
		summary:   "Camera access is denied until the user changes the application policy.",
	},
	"remote-desktop": {
		iface: "org.freedesktop.portal.RemoteDesktop", method: "CreateSession", decision: "deny",
		resources: []string{"screen", "input-devices"},
		summary:   "Remote desktop access is denied until the user changes the application policy.",
	},
}

// SupportedOperations returns the sorted set of broker-supported operations.
func SupportedOperations() []string {
	ids := make([]string, 0, len(operationCatalog))
	for id := range operationCatalog {
		ids = append(ids, id)
	}
	// Deterministic ordering without importing sort for one call site.
	for i := 1; i < len(ids); i++ {
		for j := i; j > 0 && ids[j-1] > ids[j]; j-- {
			ids[j-1], ids[j] = ids[j], ids[j-1]
		}
	}
	return ids
}

// RequestSpec is the caller-supplied intent for a new portal request.
type RequestSpec struct {
	ApplicationID string
	Operation     string
	Reason        string
}

// Request is a Runtime-owned portal request object. It carries no host paths
// and never exposes backend details.
type Request struct {
	HandleToken           string          `json:"handle_token"`
	ApplicationID         string          `json:"application_id"`
	Operation             string          `json:"operation"`
	Reason                string          `json:"reason"`
	Destination           string          `json:"destination"`
	ObjectPath            string          `json:"object_path"`
	Interface             string          `json:"interface"`
	Method                string          `json:"method"`
	Resources             []string        `json:"resources"`
	Decision              string          `json:"decision"`
	State                 RequestState    `json:"state"`
	PermissionState       PermissionState `json:"permission_state"`
	UserMediationRequired bool            `json:"user_mediation_required"`
	RequestObjectRequired bool            `json:"request_object_required"`
	DirectAccessAllowed   bool            `json:"direct_access_allowed"`
	Recoverable           bool            `json:"recoverable"`
	HostPermissionChanged bool            `json:"host_permission_changed"`
	BackendDetailsExposed bool            `json:"backend_details_exposed"`
	Diagnostics           []string        `json:"diagnostics"`
}

// Terminal reports whether the request has reached a final state.
func (r *Request) Terminal() bool {
	switch r.State {
	case StateDenied, StateCancelled, StateExpired, StateCompleted:
		return true
	default:
		return false
	}
}

func newRequest(spec RequestSpec, sequence int) (*Request, error) {
	if !appid.Valid(spec.ApplicationID) {
		return nil, errors.New("application id must be a reverse-DNS identifier")
	}
	op, ok := operationCatalog[spec.Operation]
	if !ok {
		return nil, fmt.Errorf("operation must be one of: %s", strings.Join(SupportedOperations(), ", "))
	}

	reason := strings.TrimSpace(spec.Reason)
	if reason == "" {
		reason = "Compatibility application requested " + spec.Operation + " access."
	}

	req := &Request{
		HandleToken:           handleToken(spec.ApplicationID, spec.Operation, sequence),
		ApplicationID:         spec.ApplicationID,
		Operation:             spec.Operation,
		Reason:                reason,
		Destination:           PortalDestination,
		ObjectPath:            PortalObjectPath,
		Interface:             op.iface,
		Method:                op.method,
		Resources:             append([]string(nil), op.resources...),
		Decision:              op.decision,
		UserMediationRequired: true,
		RequestObjectRequired: true,
		DirectAccessAllowed:   false,
		HostPermissionChanged: false,
		BackendDetailsExposed: false,
		Diagnostics:           []string{},
	}

	if op.decision == "deny" {
		req.State = StateDenied
		req.PermissionState = PermissionDenied
		req.Diagnostics = append(req.Diagnostics, op.summary)
	} else {
		req.State = StatePendingUserMediation
		req.PermissionState = PermissionPending
	}
	return req, nil
}

func handleToken(applicationID string, op string, sequence int) string {
	return fmt.Sprintf("xnix_%s_%s_%d", sanitizeToken(applicationID), sanitizeToken(op), sequence)
}

func sanitizeToken(value string) string {
	sanitized := make([]rune, 0, len(value))
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			sanitized = append(sanitized, r)
		} else {
			sanitized = append(sanitized, '_')
		}
	}
	return string(sanitized)
}
