package owner

import (
	"fmt"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type RouteCheckpointCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RouteCheckpoint struct {
	Version                        string                 `json:"version"`
	SchemaVersion                  string                 `json:"schema_version"`
	RequestType                    string                 `json:"request_type"`
	CheckpointType                 string                 `json:"checkpoint_type"`
	Source                         string                 `json:"source"`
	FormalReadRouteCount           int                    `json:"formal_read_route_count"`
	GoFormalReadRouteCount         int                    `json:"go_formal_read_route_count"`
	OwnerReadMethodCount           int                    `json:"owner_read_method_count"`
	OwnerLocalReadMethodCount      int                    `json:"owner_local_read_method_count"`
	SmokeReadRecordCount           int                    `json:"smoke_read_record_count"`
	SmokeWriteDenialCount          int                    `json:"smoke_write_denial_count"`
	MethodParityReady              bool                   `json:"method_parity_ready"`
	FormalRouteCoverageReady       bool                   `json:"formal_route_coverage_ready"`
	OwnerLocalRouteCoverageReady   bool                   `json:"owner_local_route_coverage_ready"`
	SmokeBatchCoverageReady        bool                   `json:"smoke_batch_coverage_ready"`
	DeterministicWriteDenialsReady bool                   `json:"deterministic_write_denials_ready"`
	RouteBandReady                 bool                   `json:"route_band_ready"`
	Checks                         []RouteCheckpointCheck `json:"checks"`
	RuntimeOwned                   bool                   `json:"runtime_owned"`
	GoRuntimeBacked                bool                   `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                   `json:"kde_policy_owner"`
	ProductionBusClaimed           bool                   `json:"production_bus_claimed"`
	SystemServiceStarted           bool                   `json:"system_service_started"`
	WriteMethodsEnabled            bool                   `json:"write_methods_enabled"`
	NetworkRequired                bool                   `json:"network_required"`
	HostRootModified               bool                   `json:"host_root_modified"`
	PrivilegedContainerRequired    bool                   `json:"privileged_container_required"`
	BackendDetailsExposed          bool                   `json:"backend_details_exposed"`
	DesktopSafeSummary             string                 `json:"desktop_safe_summary"`
}

func NewRouteCheckpoint(root string) (RouteCheckpoint, error) {
	manifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return RouteCheckpoint{}, err
	}
	records, err := NewSmokeBatchRecords(root)
	if err != nil {
		return RouteCheckpoint{}, err
	}

	readRecords := 0
	writeDenials := 0
	writesReady := true
	for _, record := range records {
		switch record.RecordType {
		case "read-dispatch":
			readRecords++
		case "write-denial":
			writeDenials++
			writesReady = writesReady && record.DispatchReady && record.ErrorName == writeMethodDisabledError
		}
	}

	ownerReadCount := len(SupportedReadDispatchMethods())
	ownerLocalCount := ownerReadCount - manifest.RouteCounts.Total
	methodParityReady := manifest.MethodParityReady
	formalCoverageReady := manifest.GoOwnerRouteCoverageReady && manifest.RouteCounts.GoRouted == manifest.RouteCounts.Total
	ownerLocalCoverageReady := ownerLocalCount == 9
	smokeCoverageReady := readRecords == ownerReadCount
	writesReady = writesReady && writeDenials == 4
	checks := []RouteCheckpointCheck{
		checkpointCheck("method-parity", methodParityReady, "The formal Runtime read contract has method parity."),
		checkpointCheck("formal-go-routes", formalCoverageReady, "Every formal Runtime read method has a Go owner route."),
		checkpointCheck("owner-local-routes", ownerLocalCoverageReady, "Nine review-only owner-local routes are present outside the production ABI."),
		checkpointCheck("smoke-read-coverage", smokeCoverageReady, "The smoke batch covers every owner read method through Service.Call."),
		checkpointCheck("write-denials", writesReady, "Every reserved Runtime write method has a deterministic disabled response."),
	}
	routeBandReady := methodParityReady && formalCoverageReady && ownerLocalCoverageReady && smokeCoverageReady && writesReady
	checkpoint := RouteCheckpoint{
		Version: manifest.Version, SchemaVersion: "xnix.runtime.owner_route_checkpoint.v1",
		RequestType: "runtime-owner-route-checkpoint", CheckpointType: "go-owner-read-route-band-checkpoint",
		Source:               "runtime-owner-route-manifest+runtime-method-parity+owner-service-smoke-batch+write-gate",
		FormalReadRouteCount: manifest.RouteCounts.Total, GoFormalReadRouteCount: manifest.RouteCounts.GoRouted,
		OwnerReadMethodCount: ownerReadCount, OwnerLocalReadMethodCount: ownerLocalCount,
		SmokeReadRecordCount: readRecords, SmokeWriteDenialCount: writeDenials,
		MethodParityReady: methodParityReady, FormalRouteCoverageReady: formalCoverageReady,
		OwnerLocalRouteCoverageReady: ownerLocalCoverageReady, SmokeBatchCoverageReady: smokeCoverageReady,
		DeterministicWriteDenialsReady: writesReady, RouteBandReady: routeBandReady, Checks: checks,
		RuntimeOwned: true, GoRuntimeBacked: true,
		DesktopSafeSummary: "Go owner read routes and deterministic write denials pass the route-band checkpoint while production ownership remains disabled.",
	}
	if err := validateNoBackendTerms(checkpoint, "Runtime owner route checkpoint"); err != nil {
		return RouteCheckpoint{}, err
	}
	return checkpoint, nil
}

func checkpointCheck(id string, ready bool, summary string) RouteCheckpointCheck {
	status := "blocked"
	if ready {
		status = "pass"
	}
	return RouteCheckpointCheck{ID: id, Status: status, Summary: summary}
}

func (checkpoint RouteCheckpoint) Validate() error {
	if !checkpoint.RouteBandReady {
		return fmt.Errorf("Runtime owner route checkpoint is not ready")
	}
	return nil
}
