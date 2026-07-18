package appidentity

import (
	"errors"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/execution"
)

type KDETestLaunchMaterializationRecord struct {
	SchemaVersion               string                                  `json:"schema_version"`
	RecordType                  string                                  `json:"record_type"`
	Source                      string                                  `json:"source"`
	Mode                        string                                  `json:"mode"`
	Application                 KDEFakeExecutionApplication             `json:"application"`
	Prerequisite                KDERestrictedLaunchPrerequisiteEvidence `json:"prerequisite"`
	Authorization               KDERestrictedPreparationAuthorization   `json:"authorization"`
	Preflight                   KDERestrictedPreflightEvidence          `json:"preflight"`
	Materialization             KDETestLaunchMaterializationEvidence    `json:"materialization"`
	Execution                   KDEFakeExecutionTransaction             `json:"execution"`
	Session                     KDEFakeExecutionSession                 `json:"session"`
	Checks                      []KDEFakeExecutionCheck                 `json:"checks"`
	CheckCount                  int                                     `json:"check_count"`
	PassedCheckCount            int                                     `json:"passed_check_count"`
	AllChecksPassed             bool                                    `json:"all_checks_passed"`
	CoreReceiptCount            int                                     `json:"core_receipt_count"`
	StateRootWritesEnabled      bool                                    `json:"state_root_writes_enabled"`
	StateRootWriteScope         string                                  `json:"state_root_write_scope"`
	StateRootPathExposed        bool                                    `json:"state_root_path_exposed"`
	MaterializationBoundary     bool                                    `json:"materialization_boundary_joined"`
	PlanMaterialized            bool                                    `json:"plan_materialized"`
	TestOnly                    bool                                    `json:"test_only"`
	ProductImageReady           bool                                    `json:"product_image_ready"`
	ProductionTrustSatisfied    bool                                    `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled     bool                                    `json:"runtime_write_gate_enabled"`
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
	HostRootModified            bool                                    `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                    `json:"privileged_container_required"`
	RawCommandExposed           bool                                    `json:"raw_command_exposed"`
	RawExecutableExposed        bool                                    `json:"raw_executable_exposed"`
	BackendDetailsExposed       bool                                    `json:"backend_details_exposed"`
	DesktopSafeSummary          string                                  `json:"desktop_safe_summary"`
}

type KDETestLaunchMaterializationEvidence struct {
	PlanID                    string   `json:"plan_id"`
	ReceiptRelativePath       string   `json:"receipt_relative_path"`
	ReceiptSHA256             string   `json:"receipt_sha256"`
	Status                    string   `json:"status"`
	MaterializationScope      string   `json:"materialization_scope"`
	MaterializedArtifactIDs   []string `json:"materialized_artifact_ids"`
	MaterializedArtifactCount int      `json:"materialized_artifact_count"`
	BlockedByIDs              []string `json:"blocked_by_ids"`
	BlockedByCount            int      `json:"blocked_by_count"`
	ReceiptPersisted          bool     `json:"receipt_persisted"`
	ReceiptReadBack           bool     `json:"receipt_read_back"`
	SafeInputsReady           bool     `json:"safe_inputs_ready"`
	PreparationAuthorized     bool     `json:"preparation_authorized"`
	PreflightReadBack         bool     `json:"preflight_read_back"`
	PlanMaterialized          bool     `json:"plan_materialized"`
	TestOnly                  bool     `json:"test_only"`
	CommandMaterialized       bool     `json:"command_materialized"`
	ExecutablePathResolved    bool     `json:"executable_path_resolved"`
	BackendSelectedForLaunch  bool     `json:"backend_selected_for_launch"`
	BackendLaunchEnabled      bool     `json:"backend_launch_enabled"`
}

func NewKDETestLaunchMaterializationRecord(recipeRecord Recipe, provenance Provenance, options KDERestrictedLaunchAuthorizationOptions) (KDETestLaunchMaterializationRecord, error) {
	preflightRecord, err := NewKDERestrictedLaunchPreflightRecord(recipeRecord, provenance, options)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	authorizationStore, err := execution.NewRestrictedAuthorizationStore(options.StateRoot)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	authorization, err := authorizationStore.Load(preflightRecord.Authorization.AuthorizationID)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	preflightStore, err := execution.NewRestrictedPreflightStore(options.StateRoot)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	preflight, err := preflightStore.Load(preflightRecord.Preflight.PacketID)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	ledger, err := execution.NewLedger(options.StateRoot)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	transaction, err := ledger.Load(preflightRecord.Execution.RequestID)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	session, err := ledger.LoadSession(preflightRecord.Execution.RequestID)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	materializationStore, err := execution.NewRestrictedMaterializationStore(options.StateRoot)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	recorded, err := materializationStore.Record(execution.RestrictedMaterializationRequest{
		Authorization: authorization,
		Preflight:     preflight,
		Transaction:   transaction,
	})
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	materialization, err := materializationStore.Load(recorded.PlanID)
	if err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	checks := kdeTestLaunchMaterializationChecks(preflightRecord, materialization, transaction, session)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDETestLaunchMaterializationRecord{}, errors.New("test launch materialization checks did not all pass")
	}

	record := KDETestLaunchMaterializationRecord{
		SchemaVersion: "xnix.runtime.kde_test_launch_materialization.v1",
		RecordType:    "kde-test-launch-materialization-record",
		Source:        "restricted-launch-materialization-plan+restricted-launch-preflight-packet+restricted-launch-authorization-receipt+execution-ledger",
		Mode:          options.Mode,
		Application:   preflightRecord.Application,
		Prerequisite:  preflightRecord.Prerequisite,
		Authorization: preflightRecord.Authorization,
		Preflight:     preflightRecord.Preflight,
		Materialization: KDETestLaunchMaterializationEvidence{
			PlanID:                    materialization.PlanID,
			ReceiptRelativePath:       materialization.RelativePath,
			ReceiptSHA256:             materialization.SHA256,
			Status:                    materialization.Status,
			MaterializationScope:      materialization.MaterializationScope,
			MaterializedArtifactIDs:   append([]string{}, materialization.MaterializedArtifactIDs...),
			MaterializedArtifactCount: len(materialization.MaterializedArtifactIDs),
			BlockedByIDs:              append([]string{}, materialization.BlockedByIDs...),
			BlockedByCount:            materialization.BlockedByCount,
			ReceiptPersisted:          true,
			ReceiptReadBack:           true,
			SafeInputsReady:           materialization.SafeInputsReady,
			PreparationAuthorized:     materialization.PreparationAuthorized,
			PreflightReadBack:         materialization.PreflightReadBack,
			PlanMaterialized:          materialization.PlanMaterialized,
			TestOnly:                  materialization.TestOnly,
			CommandMaterialized:       materialization.CommandMaterialized,
			ExecutablePathResolved:    materialization.ExecutablePathResolved,
			BackendSelectedForLaunch:  materialization.BackendSelectedForLaunch,
			BackendLaunchEnabled:      materialization.BackendLaunchEnabled,
		},
		Execution:                   preflightRecord.Execution,
		Session:                     preflightRecord.Session,
		Checks:                      checks,
		CheckCount:                  len(checks),
		PassedCheckCount:            passed,
		AllChecksPassed:             true,
		CoreReceiptCount:            preflightRecord.CoreReceiptCount + 1,
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		StateRootPathExposed:        false,
		MaterializationBoundary:     true,
		PlanMaterialized:            true,
		TestOnly:                    true,
		ProductImageReady:           false,
		ProductionTrustSatisfied:    false,
		RuntimeWriteGateEnabled:     false,
		LaunchPreflightPassed:       false,
		LaunchAuthorized:            false,
		ExecutionApproved:           false,
		ProcessStartAuthorized:      false,
		CommandMaterialized:         false,
		ExecutablePathResolved:      false,
		BackendSelectedForLaunch:    false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		ProductionBusOwnership:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "Runtime materialized a test-only launch review plan while keeping command materialization, executable resolution, backend selection, launch, execution, and process start disabled.",
	}
	if err := validateNoBackendTerms(record, "KDE test launch materialization record"); err != nil {
		return KDETestLaunchMaterializationRecord{}, err
	}
	return record, nil
}

func kdeTestLaunchMaterializationChecks(preflight KDERestrictedLaunchPreflightRecord, materialization execution.RestrictedMaterializationPlan, transaction execution.LedgerRecord, session execution.SessionRecord) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("preflight-boundary", preflight.AllChecksPassed && preflight.Preflight.Status == "blocked" && preflight.Preflight.ReadyForPacketAssembly && !preflight.LaunchPreflightPassed, "Restricted launch preflight remains blocked and ready only for safe evidence materialization."),
		fakeExecutionCheck("materialization-readback", materialization.Status == "blocked-plan-materialized" && len(materialization.SHA256) == 64 && !filepath.IsAbs(materialization.RelativePath) && materialization.PlanMaterialized && materialization.TestOnly, "The test-only materialization plan passed digest-verified readback."),
		fakeExecutionCheck("materialized-plan-only", materialization.MaterializationScope == "test-only-review-plan" && len(materialization.MaterializedArtifactIDs) == 4 && containsString(materialization.MaterializedArtifactIDs, "launch-intent-reference") && containsString(materialization.MaterializedArtifactIDs, "preflight-reference"), "Only review-plan references were materialized."),
		fakeExecutionCheck("write-gate-remains-blocked", materialization.BlockedByCount == 2 && containsString(materialization.BlockedByIDs, "recipe-trust") && containsString(materialization.BlockedByIDs, "runtime-write-gate") && !materialization.RuntimeWriteGateEnabled, "Production trust and Runtime write gate remain the explicit blockers."),
		fakeExecutionCheck("command-boundary", !materialization.CommandMaterialized && !materialization.ExecutablePathResolved && !materialization.RawCommandExposed && !materialization.RawExecutableExposed, "No command line, executable path, or raw launch detail is materialized."),
		fakeExecutionCheck("launch-boundary", !materialization.LaunchAuthorized && !materialization.ExecutionApproved && !materialization.ProcessStartAuthorized && !materialization.BackendSelectedForLaunch && !materialization.BackendLaunchEnabled && !materialization.BackendProcessStarted, "Launch authority, execution approval, backend selection, and process start remain disabled."),
		fakeExecutionCheck("execution-unchanged", transaction.Transaction.State == execution.StateBlocked && fakePortalExecutionSummary(transaction).PassedGateCount == 4 && !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted, "The execution transaction remains blocked and unchanged."),
		fakeExecutionCheck("session-unchanged", session.SessionState == "blocked" && !session.SessionActive && !session.ExecutionStarted && !session.BackendProcessStarted, "The execution session remains blocked and inactive."),
		fakeExecutionCheck("unsafe-gates-closed", !materialization.ProductionTrustSatisfied && !materialization.RuntimeWriteGateEnabled && !materialization.NetworkRequired && !materialization.PrivilegedContainerRequired && !materialization.HostRootModified && !materialization.BackendDetailsExposed && !materialization.StateRootPathExposed, "Materialization does not enable network, privilege, host mutation, state-root exposure, or backend detail exposure."),
	}
}
