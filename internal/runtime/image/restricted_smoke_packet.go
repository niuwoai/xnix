package image

import (
	"errors"
	"fmt"
	"path/filepath"
)

type RestrictedSmokeEvidence struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	EvidenceCount int    `json:"evidence_count"`
	Summary       string `json:"summary"`
}

type RestrictedProductSmokePacket struct {
	SchemaVersion                 string                    `json:"schema_version"`
	RequestType                   string                    `json:"request_type"`
	PacketType                    string                    `json:"packet_type"`
	ImageName                     string                    `json:"image_name"`
	Architecture                  string                    `json:"architecture"`
	Evidence                      []RestrictedSmokeEvidence `json:"evidence"`
	EvidenceCount                 int                       `json:"evidence_count"`
	ReadyEvidenceCount            int                       `json:"ready_evidence_count"`
	MissingEvidenceCount          int                       `json:"missing_evidence_count"`
	ImageManifestReady            bool                      `json:"image_manifest_ready"`
	KDEEntryPointsCovered         bool                      `json:"kde_entry_points_covered"`
	RuntimeOwnerEvidenceReady     bool                      `json:"runtime_owner_evidence_ready"`
	ArtifactTrustEvidenceReady    bool                      `json:"artifact_trust_evidence_ready"`
	BackendLifecycleEvidenceReady bool                      `json:"backend_lifecycle_evidence_ready"`
	PortalSafetyEvidenceReady     bool                      `json:"portal_safety_evidence_ready"`
	KDEEntryPointEvidenceReady    bool                      `json:"kde_entry_point_evidence_ready"`
	PacketPrepared                bool                      `json:"packet_prepared"`
	ReadyForAuthorizedSmoke       bool                      `json:"ready_for_authorized_smoke"`
	HumanAuthorizationRequired    bool                      `json:"human_authorization_required"`
	ExecutionAuthorized           bool                      `json:"execution_authorized"`
	DockerExecuted                bool                      `json:"docker_executed"`
	QEMUExecuted                  bool                      `json:"qemu_executed"`
	ProductSmokeExecuted          bool                      `json:"product_smoke_executed"`
	SerialLogPersistenceRequired  bool                      `json:"serial_log_persistence_required"`
	SerialLogPersisted            bool                      `json:"serial_log_persisted"`
	LoopbackOnlyNetworking        bool                      `json:"loopback_only_networking"`
	DockerSocketMounted           bool                      `json:"docker_socket_mounted"`
	HostNetworkEnabled            bool                      `json:"host_network_enabled"`
	BroadHostMountEnabled         bool                      `json:"broad_host_mount_enabled"`
	PrivilegedContainerRequired   bool                      `json:"privileged_container_required"`
	BackendLaunchEnabled          bool                      `json:"backend_launch_enabled"`
	HostRootModified              bool                      `json:"host_root_modified"`
	ReleaseReady                  bool                      `json:"release_ready"`
	BlockingReasons               []string                  `json:"blocking_reasons"`
	DesktopSafeSummary            string                    `json:"desktop_safe_summary"`
}

type restrictedSmokeEvidenceSpec struct {
	id      string
	summary string
	sources []string
}

var restrictedSmokeEvidenceSpecs = []restrictedSmokeEvidenceSpec{
	{id: "runtime-owner", summary: "Go Runtime owner and read-service evidence are present.", sources: []string{"cmd/xnix-runtime-owner/main.go", "internal/runtime/owner/service.go"}},
	{id: "artifact-trust", summary: "Digest, staging, and signed recipe verifier evidence are present.", sources: []string{"internal/runtime/artifact/stage.go", "internal/runtime/recipe/signature.go"}},
	{id: "backend-lifecycle", summary: "Runtime lifecycle and KDE-safe lifecycle evidence are present.", sources: []string{"internal/runtime/environment/lifecycle.go", "internal/runtime/appidentity/backend_lifecycle.go"}},
	{id: "portal-safety", summary: "Fake-mode Portal ledger and permission safety evidence are present.", sources: []string{"internal/runtime/portal/ledger.go", "internal/runtime/appidentity/portal_access_policy.go"}},
	{id: "kde-entrypoints", summary: "KDE entry-point planning and presence smoke evidence are present.", sources: []string{"internal/runtime/appidentity/kde_entrypoints.go", "scripts/kde_first_presence_smoke.rb"}},
}

func PrepareRestrictedProductSmokePacket(repoRoot, manifestSource string) (RestrictedProductSmokePacket, error) {
	if repoRoot == "" {
		return RestrictedProductSmokePacket{}, errors.New("restricted product smoke packet requires a repo root")
	}
	if manifestSource == "" {
		manifestSource = "image/kinoite/manifest.json"
	}
	manifestPath, err := safeSource(repoRoot, manifestSource)
	if err != nil {
		return RestrictedProductSmokePacket{}, err
	}
	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		return RestrictedProductSmokePacket{}, err
	}
	imageReport, err := Verify(repoRoot, manifest)
	if err != nil {
		return RestrictedProductSmokePacket{}, err
	}

	packet := RestrictedProductSmokePacket{
		SchemaVersion:                "xnix.runtime.restricted_product_smoke_packet.v1",
		RequestType:                  "restricted-product-smoke-packet-preview",
		PacketType:                   "dry-run-product-image-smoke-readiness",
		ImageName:                    manifest.ImageName,
		Architecture:                 manifest.Architecture,
		ImageManifestReady:           imageReport.RuntimeReady,
		KDEEntryPointsCovered:        imageReport.KDEEntryPointsCovered,
		HumanAuthorizationRequired:   true,
		SerialLogPersistenceRequired: true,
		LoopbackOnlyNetworking:       true,
		BlockingReasons:              []string{"Docker and QEMU execution require explicit human authorization.", "Persisted serial-log evidence is not available until the authorized smoke runs."},
		DesktopSafeSummary:           "The restricted product smoke packet is prepared for review while Docker, QEMU, backend launch, and host mutation remain disabled.",
	}

	for _, spec := range restrictedSmokeEvidenceSpecs {
		present := 0
		for _, source := range spec.sources {
			exists, sourceErr := sourceExists(repoRoot, filepath.ToSlash(source))
			if sourceErr != nil {
				return RestrictedProductSmokePacket{}, fmt.Errorf("check restricted smoke evidence %s: %w", spec.id, sourceErr)
			}
			if exists {
				present++
			}
		}
		status := "ready"
		if present != len(spec.sources) {
			status = "missing-evidence"
			packet.MissingEvidenceCount++
		} else {
			packet.ReadyEvidenceCount++
		}
		packet.Evidence = append(packet.Evidence, RestrictedSmokeEvidence{ID: spec.id, Status: status, EvidenceCount: present, Summary: spec.summary})
	}
	packet.EvidenceCount = len(packet.Evidence)
	packet.RuntimeOwnerEvidenceReady = restrictedSmokeEvidenceReady(packet.Evidence, "runtime-owner")
	packet.ArtifactTrustEvidenceReady = restrictedSmokeEvidenceReady(packet.Evidence, "artifact-trust")
	packet.BackendLifecycleEvidenceReady = restrictedSmokeEvidenceReady(packet.Evidence, "backend-lifecycle")
	packet.PortalSafetyEvidenceReady = restrictedSmokeEvidenceReady(packet.Evidence, "portal-safety")
	packet.KDEEntryPointEvidenceReady = restrictedSmokeEvidenceReady(packet.Evidence, "kde-entrypoints")
	packet.PacketPrepared = true
	packet.ReadyForAuthorizedSmoke = packet.ImageManifestReady && packet.KDEEntryPointsCovered && packet.MissingEvidenceCount == 0
	return packet, nil
}

func restrictedSmokeEvidenceReady(evidence []RestrictedSmokeEvidence, id string) bool {
	for _, item := range evidence {
		if item.ID == id {
			return item.Status == "ready"
		}
	}
	return false
}
