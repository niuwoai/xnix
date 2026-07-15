package owner

import "fmt"

type LifecycleEvent struct {
	Version                     string `json:"version"`
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	EventType                   string `json:"event_type"`
	Sequence                    int    `json:"sequence"`
	Mode                        string `json:"mode"`
	BusName                     string `json:"bus_name"`
	ObjectPath                  string `json:"object_path"`
	Interface                   string `json:"interface"`
	RouteTableVersion           string `json:"route_table_version"`
	RouteCount                  int    `json:"route_count"`
	GoRouteCount                int    `json:"go_route_count"`
	WriteMethodCount            int    `json:"write_method_count"`
	ReadOnlyServeReady          bool   `json:"read_only_serve_ready"`
	WriteMethodsEnabled         bool   `json:"write_methods_enabled"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool   `json:"kde_may_claim_runtime_ownership"`
	EventLoopStarted            bool   `json:"event_loop_started"`
	SessionBusClaimed           bool   `json:"session_bus_claimed"`
	ProductionBusClaimed        bool   `json:"production_bus_claimed"`
	SystemServiceStarted        bool   `json:"system_service_started"`
	NetworkRequired             bool   `json:"network_required"`
	HostRootModified            bool   `json:"host_root_modified"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	ShutdownReason              string `json:"shutdown_reason,omitempty"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
}

func NewLifecycleEvents(root string, mode CandidateMode) ([]LifecycleEvent, error) {
	candidate, err := NewCandidate(root, mode)
	if err != nil {
		return nil, err
	}
	events := []LifecycleEvent{
		newLifecycleEvent(candidate, 1, "startup", "Runtime owner smoke startup requested without claiming a bus name.", ""),
		newLifecycleEvent(candidate, 2, "route-table", fmt.Sprintf("Read-only route table loaded with %d Go owner routes.", candidate.GoRouteCount), ""),
		newLifecycleEvent(candidate, 3, "readiness", candidate.DesktopSafeSummary, ""),
		newLifecycleEvent(candidate, 4, "shutdown", "Runtime owner smoke lifecycle preview ended without starting an event loop.", "preview-complete"),
	}
	for _, event := range events {
		if err := validateNoBackendTerms(event, "Runtime owner lifecycle event"); err != nil {
			return nil, err
		}
	}
	return events, nil
}

func newLifecycleEvent(candidate Candidate, sequence int, eventType string, summary string, shutdownReason string) LifecycleEvent {
	return LifecycleEvent{
		Version:                     candidate.Version,
		SchemaVersion:               "xnix.runtime.owner_lifecycle_event.v1",
		RequestType:                 "runtime-owner-lifecycle-event",
		EventType:                   eventType,
		Sequence:                    sequence,
		Mode:                        candidate.Mode,
		BusName:                     candidate.BusName,
		ObjectPath:                  candidate.ObjectPath,
		Interface:                   candidate.Interface,
		RouteTableVersion:           candidate.Version,
		RouteCount:                  candidate.RouteCount,
		GoRouteCount:                candidate.GoRouteCount,
		WriteMethodCount:            candidate.WriteMethodCount,
		ReadOnlyServeReady:          candidate.ReadOnlyServeReady,
		WriteMethodsEnabled:         false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		EventLoopStarted:            false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		ShutdownReason:              shutdownReason,
		DesktopSafeSummary:          summary,
	}
}
