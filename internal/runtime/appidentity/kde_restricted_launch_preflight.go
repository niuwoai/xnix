package appidentity

import (
	"errors"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/execution"
)

type KDERestrictedLaunchPreflightRecord struct {
	SchemaVersion               string                                  `json:"schema_version"`
	RecordType                  string                                  `json:"record_type"`
	Source                      string                                  `json:"source"`
	Mode                        string                                  `json:"mode"`
	Application                 KDEFakeExecutionApplication             `json:"application"`
	Prerequisite                KDERestrictedLaunchPrerequisiteEvidence `json:"prerequisite"`
	Authorization               KDERestrictedPreparationAuthorization   `json:"authorization"`
	Preflight                   KDERestrictedPreflightEvidence          `json:"preflight"`
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
	PreflightBoundaryJoined     bool                                    `json:"preflight_boundary_joined"`
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
	BackendDetailsExposed       bool                                    `json:"backend_details_exposed"`
	DesktopSafeSummary          string                                  `json:"desktop_safe_summary"`
}

type KDERestrictedPreflightEvidence struct {
	PacketID               string   `json:"packet_id"`
	ReceiptRelativePath    string   `json:"receipt_relative_path"`
	ReceiptSHA256          string   `json:"receipt_sha256"`
	Status                 string   `json:"status"`
	BlockerIDs             []string `json:"blocker_ids"`
	BlockerCount           int      `json:"blocker_count"`
	ReceiptPersisted       bool     `json:"receipt_persisted"`
	ReceiptReadBack        bool     `json:"receipt_read_back"`
	SafeInputsReady        bool     `json:"safe_inputs_ready"`
	PreparationAuthorized  bool     `json:"preparation_authorized"`
	ReadyForPacketAssembly bool     `json:"ready_for_packet_assembly"`
	ProductImageReady      bool     `json:"product_image_ready"`
	LaunchPreflightPassed  bool     `json:"launch_preflight_passed"`
}

func NewKDERestrictedLaunchPreflightRecord(recipeRecord Recipe, provenance Provenance, options KDERestrictedLaunchAuthorizationOptions) (KDERestrictedLaunchPreflightRecord, error) {
	authorizationEvidence, err := NewKDERestrictedLaunchAuthorizationRecord(recipeRecord, provenance, options)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	authorizationStore, err := execution.NewRestrictedAuthorizationStore(options.StateRoot)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	authorization, err := authorizationStore.Load(authorizationEvidence.Authorization.AuthorizationID)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	ledger, err := execution.NewLedger(options.StateRoot)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	transaction, err := ledger.Load(authorizationEvidence.Execution.RequestID)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	session, err := ledger.LoadSession(authorizationEvidence.Execution.RequestID)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	preflightStore, err := execution.NewRestrictedPreflightStore(options.StateRoot)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	recorded, err := preflightStore.Record(execution.RestrictedPreflightRequest{Authorization: authorization, Transaction: transaction, SafeInputsReady: authorizationEvidence.Prerequisite.AllChecksPassed, ProductImageReady: false})
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	preflight, err := preflightStore.Load(recorded.PacketID)
	if err != nil {
		return KDERestrictedLaunchPreflightRecord{}, err
	}
	checks := kdeRestrictedLaunchPreflightChecks(authorizationEvidence, preflight, transaction, session)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDERestrictedLaunchPreflightRecord{}, errors.New("restricted launch preflight evidence checks did not all pass")
	}

	return KDERestrictedLaunchPreflightRecord{
		SchemaVersion: "xnix.runtime.kde_restricted_launch_preflight.v1",
		RecordType:    "kde-restricted-launch-preflight-record",
		Source:        "restricted-launch-preflight-packet+restricted-launch-authorization-receipt+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application:   authorizationEvidence.Application,
		Prerequisite:  authorizationEvidence.Prerequisite,
		Authorization: authorizationEvidence.Authorization,
		Preflight: KDERestrictedPreflightEvidence{
			PacketID:               preflight.PacketID,
			ReceiptRelativePath:    preflight.RelativePath,
			ReceiptSHA256:          preflight.SHA256,
			Status:                 preflight.Status,
			BlockerIDs:             append([]string{}, preflight.BlockerIDs...),
			BlockerCount:           preflight.BlockerCount,
			ReceiptPersisted:       true,
			ReceiptReadBack:        true,
			SafeInputsReady:        preflight.SafeInputsReady,
			PreparationAuthorized:  preflight.PreparationAuthorized,
			ReadyForPacketAssembly: preflight.ReadyForPacketAssembly,
			ProductImageReady:      false,
			LaunchPreflightPassed:  false,
		},
		Execution:                   fakePortalExecutionSummary(transaction),
		Session:                     fakePortalSessionSummary(session),
		Checks:                      checks,
		CheckCount:                  len(checks),
		PassedCheckCount:            passed,
		AllChecksPassed:             true,
		CoreReceiptCount:            authorizationEvidence.CoreReceiptCount + 1,
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		StateRootPathExposed:        false,
		PreflightBoundaryJoined:     true,
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
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "Restricted launch preflight consumed explicit preparation authorization and converged safety evidence, then remained blocked by production trust and Runtime write gates without materializing a command.",
	}, nil
}

func kdeRestrictedLaunchPreflightChecks(authorization KDERestrictedLaunchAuthorizationRecord, preflight execution.RestrictedPreflightPacket, transaction execution.LedgerRecord, session execution.SessionRecord) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("authorization-boundary", authorization.AllChecksPassed && authorization.Authorization.PreparationAuthorized && !authorization.LaunchAuthorized && !authorization.ProcessStartAuthorized, "Preparation-only authorization remains valid and launch authority remains absent."),
		fakeExecutionCheck("safe-inputs", authorization.Prerequisite.AllChecksPassed && authorization.Prerequisite.LifecycleReady && authorization.Prerequisite.SnapshotVerified && authorization.Prerequisite.DiagnosticVerified && authorization.Prerequisite.PortalReceiptCompleted && authorization.Prerequisite.BackendProcessesStopped, "All non-launching preflight inputs are present and verified."),
		fakeExecutionCheck("preflight-readback", preflight.Status == "blocked" && len(preflight.SHA256) == 64 && !filepath.IsAbs(preflight.RelativePath) && preflight.SafeInputsReady && preflight.PreparationAuthorized, "The fail-closed preflight packet passed digest-verified readback."),
		fakeExecutionCheck("explicit-blockers", preflight.BlockerCount == 2 && containsString(preflight.BlockerIDs, "recipe-trust") && containsString(preflight.BlockerIDs, "runtime-write-gate"), "Preflight reports exactly production trust and Runtime write blockers."),
		fakeExecutionCheck("packet-assembly-boundary", preflight.ReadyForPacketAssembly && !preflight.ProductImageReady && !preflight.LaunchPreflightPassed, "Evidence may advance to product-image packet assembly but not launch preflight."),
		fakeExecutionCheck("execution-unchanged", transaction.Transaction.State == execution.StateBlocked && fakePortalExecutionSummary(transaction).PassedGateCount == 4 && !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted, "The execution transaction remains blocked and unchanged."),
		fakeExecutionCheck("session-unchanged", session.SessionState == "blocked" && !session.SessionActive && !session.ExecutionStarted && !session.BackendProcessStarted, "The execution session remains blocked and inactive."),
		fakeExecutionCheck("unsafe-gates-closed", !preflight.ProductionTrustSatisfied && !preflight.RuntimeWriteGateEnabled && !preflight.LaunchAuthorized && !preflight.ExecutionApproved && !preflight.ProcessStartAuthorized && !preflight.CommandMaterialized && !preflight.ExecutablePathResolved && !preflight.BackendSelectedForLaunch && !preflight.BackendLaunchEnabled && !preflight.BackendProcessStarted && !preflight.NetworkRequired && !preflight.PrivilegedContainerRequired && !preflight.HostRootModified, "Preflight does not materialize commands, resolve paths, select a launch backend, or enable side effects."),
	}
}
