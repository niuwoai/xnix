package appidentity

import (
	"errors"

	"xnix.local/xnix/internal/runtime/image"
)

type KDERestrictedProductSmokeCheckpointOptions struct {
	KDERestrictedLaunchAuthorizationOptions
	RepositoryRoot string
	ManifestSource string
}

type KDERestrictedProductSmokeCheckpointRecord struct {
	SchemaVersion               string                                  `json:"schema_version"`
	RecordType                  string                                  `json:"record_type"`
	Source                      string                                  `json:"source"`
	Mode                        string                                  `json:"mode"`
	Application                 KDEFakeExecutionApplication             `json:"application"`
	Prerequisite                KDERestrictedLaunchPrerequisiteEvidence `json:"prerequisite"`
	Authorization               KDERestrictedPreparationAuthorization   `json:"authorization"`
	Preflight                   KDERestrictedPreflightEvidence          `json:"preflight"`
	ProductImage                KDERestrictedProductImageEvidence       `json:"product_image"`
	Execution                   KDEFakeExecutionTransaction             `json:"execution"`
	Session                     KDEFakeExecutionSession                 `json:"session"`
	Checks                      []KDEFakeExecutionCheck                 `json:"checks"`
	CheckCount                  int                                     `json:"check_count"`
	PassedCheckCount            int                                     `json:"passed_check_count"`
	AllChecksPassed             bool                                    `json:"all_checks_passed"`
	CoreReceiptCount            int                                     `json:"core_receipt_count"`
	CheckpointReady             bool                                    `json:"checkpoint_ready"`
	ReadyForTrainGate           bool                                    `json:"ready_for_train_gate"`
	HumanAuthorizationRequired  bool                                    `json:"human_authorization_required"`
	ExecutionAuthorized         bool                                    `json:"execution_authorized"`
	DockerExecuted              bool                                    `json:"docker_executed"`
	QEMUExecuted                bool                                    `json:"qemu_executed"`
	ProductSmokeExecuted        bool                                    `json:"product_smoke_executed"`
	SerialLogPersisted          bool                                    `json:"serial_log_persisted"`
	ReleaseReady                bool                                    `json:"release_ready"`
	LaunchPreflightPassed       bool                                    `json:"launch_preflight_passed"`
	LaunchAuthorized            bool                                    `json:"launch_authorized"`
	ExecutionApproved           bool                                    `json:"execution_approved"`
	ProcessStartAuthorized      bool                                    `json:"process_start_authorized"`
	CommandMaterialized         bool                                    `json:"command_materialized"`
	ExecutablePathResolved      bool                                    `json:"executable_path_resolved"`
	BackendSelectedForLaunch    bool                                    `json:"backend_selected_for_launch"`
	BackendLaunchEnabled        bool                                    `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                                    `json:"backend_process_started"`
	ProductionBusOwnership      bool                                    `json:"production_bus_ownership"`
	NetworkRequired             bool                                    `json:"network_required"`
	DockerSocketMounted         bool                                    `json:"docker_socket_mounted"`
	HostNetworkEnabled          bool                                    `json:"host_network_enabled"`
	BroadHostMountEnabled       bool                                    `json:"broad_host_mount_enabled"`
	PrivilegedContainerRequired bool                                    `json:"privileged_container_required"`
	HostRootModified            bool                                    `json:"host_root_modified"`
	StateRootPathExposed        bool                                    `json:"state_root_path_exposed"`
	RepositoryRootPathExposed   bool                                    `json:"repository_root_path_exposed"`
	ManifestSourcePathExposed   bool                                    `json:"manifest_source_path_exposed"`
	RawCommandExposed           bool                                    `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                                    `json:"backend_details_exposed"`
	DesktopSafeSummary          string                                  `json:"desktop_safe_summary"`
}

type KDERestrictedProductImageEvidence struct {
	ImageName                    string   `json:"image_name"`
	Architecture                 string   `json:"architecture"`
	ImageManifestReady           bool     `json:"image_manifest_ready"`
	KDEEntryPointsCovered        bool     `json:"kde_entry_points_covered"`
	EvidenceIDs                  []string `json:"evidence_ids"`
	EvidenceCount                int      `json:"evidence_count"`
	ReadyEvidenceCount           int      `json:"ready_evidence_count"`
	MissingEvidenceCount         int      `json:"missing_evidence_count"`
	PacketPrepared               bool     `json:"packet_prepared"`
	ProductImageMetadataReady    bool     `json:"product_image_metadata_ready"`
	ReadyForAuthorizedSmoke      bool     `json:"ready_for_authorized_smoke"`
	HumanAuthorizationRequired   bool     `json:"human_authorization_required"`
	SerialLogPersistenceRequired bool     `json:"serial_log_persistence_required"`
	LoopbackOnlyNetworking       bool     `json:"loopback_only_networking"`
	DockerExecuted               bool     `json:"docker_executed"`
	QEMUExecuted                 bool     `json:"qemu_executed"`
	ProductSmokeExecuted         bool     `json:"product_smoke_executed"`
	SerialLogPersisted           bool     `json:"serial_log_persisted"`
	ReleaseReady                 bool     `json:"release_ready"`
	BlockingReasons              []string `json:"blocking_reasons"`
}

func NewKDERestrictedProductSmokeCheckpointRecord(recipeRecord Recipe, provenance Provenance, options KDERestrictedProductSmokeCheckpointOptions) (KDERestrictedProductSmokeCheckpointRecord, error) {
	preflightRecord, err := NewKDERestrictedLaunchPreflightRecord(recipeRecord, provenance, options.KDERestrictedLaunchAuthorizationOptions)
	if err != nil {
		return KDERestrictedProductSmokeCheckpointRecord{}, err
	}
	smokePacket, err := image.PrepareRestrictedProductSmokePacket(options.RepositoryRoot, options.ManifestSource)
	if err != nil {
		return KDERestrictedProductSmokeCheckpointRecord{}, err
	}
	checks := kdeRestrictedProductSmokeCheckpointChecks(preflightRecord, smokePacket)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDERestrictedProductSmokeCheckpointRecord{}, errors.New("restricted product smoke checkpoint checks did not all pass")
	}
	evidenceIDs := make([]string, 0, len(smokePacket.Evidence))
	for _, item := range smokePacket.Evidence {
		evidenceIDs = append(evidenceIDs, item.ID)
	}
	checkpointReady := preflightRecord.Preflight.ReadyForPacketAssembly && smokePacket.ReadyForAuthorizedSmoke
	return KDERestrictedProductSmokeCheckpointRecord{
		SchemaVersion: "xnix.runtime.kde_restricted_product_smoke_checkpoint.v1",
		RecordType:    "kde-restricted-product-smoke-checkpoint-record",
		Source:        "restricted-launch-preflight-packet+restricted-product-smoke-packet+product-image-manifest",
		Mode:          options.Mode,
		Application:   preflightRecord.Application,
		Prerequisite:  preflightRecord.Prerequisite,
		Authorization: preflightRecord.Authorization,
		Preflight:     preflightRecord.Preflight,
		ProductImage: KDERestrictedProductImageEvidence{
			ImageName:                    smokePacket.ImageName,
			Architecture:                 smokePacket.Architecture,
			ImageManifestReady:           smokePacket.ImageManifestReady,
			KDEEntryPointsCovered:        smokePacket.KDEEntryPointsCovered,
			EvidenceIDs:                  evidenceIDs,
			EvidenceCount:                smokePacket.EvidenceCount,
			ReadyEvidenceCount:           smokePacket.ReadyEvidenceCount,
			MissingEvidenceCount:         smokePacket.MissingEvidenceCount,
			PacketPrepared:               smokePacket.PacketPrepared,
			ProductImageMetadataReady:    smokePacket.ImageManifestReady && smokePacket.KDEEntryPointsCovered,
			ReadyForAuthorizedSmoke:      smokePacket.ReadyForAuthorizedSmoke,
			HumanAuthorizationRequired:   smokePacket.HumanAuthorizationRequired,
			SerialLogPersistenceRequired: smokePacket.SerialLogPersistenceRequired,
			LoopbackOnlyNetworking:       smokePacket.LoopbackOnlyNetworking,
			BlockingReasons:              append([]string{}, smokePacket.BlockingReasons...),
		},
		Execution:                  preflightRecord.Execution,
		Session:                    preflightRecord.Session,
		Checks:                     checks,
		CheckCount:                 len(checks),
		PassedCheckCount:           passed,
		AllChecksPassed:            true,
		CoreReceiptCount:           preflightRecord.CoreReceiptCount,
		CheckpointReady:            checkpointReady,
		ReadyForTrainGate:          checkpointReady,
		HumanAuthorizationRequired: true,
		DesktopSafeSummary:         "Restricted product-image metadata is ready for the train gate; Docker, QEMU, product smoke, compatibility launch, and release readiness remain disabled pending authorized evidence.",
	}, nil
}

func kdeRestrictedProductSmokeCheckpointChecks(preflight KDERestrictedLaunchPreflightRecord, smoke image.RestrictedProductSmokePacket) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("preflight-boundary", preflight.AllChecksPassed && preflight.Preflight.ReadyForPacketAssembly && !preflight.LaunchPreflightPassed, "The restricted launch preflight packet is ready only for evidence assembly."),
		fakeExecutionCheck("product-image-manifest", smoke.ImageManifestReady && smoke.KDEEntryPointsCovered, "The fixed product-image manifest and KDE entry-point metadata are present."),
		fakeExecutionCheck("repository-evidence", smoke.EvidenceCount == 5 && smoke.ReadyEvidenceCount == 5 && smoke.MissingEvidenceCount == 0, "All five restricted product smoke evidence groups are present."),
		fakeExecutionCheck("authorization-boundary", smoke.HumanAuthorizationRequired && !smoke.ExecutionAuthorized, "Product smoke execution still requires separate human authorization."),
		fakeExecutionCheck("execution-unchanged", preflight.Execution.State == "blocked" && !preflight.Execution.LaunchAllowed && !preflight.Execution.LaunchEnabled && !preflight.Execution.ExecutionStarted && !preflight.Execution.ProcessStarted, "The compatibility execution transaction remains blocked."),
		fakeExecutionCheck("session-unchanged", preflight.Session.State == "blocked" && !preflight.Session.LiveStateObserved && !preflight.Session.SessionActive, "The compatibility session remains blocked and inactive."),
		fakeExecutionCheck("smoke-not-executed", !smoke.DockerExecuted && !smoke.QEMUExecuted && !smoke.ProductSmokeExecuted && !smoke.SerialLogPersisted && !smoke.ReleaseReady, "No product smoke or release evidence is claimed."),
		fakeExecutionCheck("host-boundary", smoke.LoopbackOnlyNetworking && !smoke.DockerSocketMounted && !smoke.HostNetworkEnabled && !smoke.BroadHostMountEnabled && !smoke.PrivilegedContainerRequired && !smoke.BackendLaunchEnabled && !smoke.HostRootModified, "The restricted host boundary remains closed."),
	}
}
