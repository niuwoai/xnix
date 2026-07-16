package appidentity

import (
	"errors"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/execution"
)

type KDERestrictedLaunchAuthorizationOptions struct {
	StateRoot string
	Mode      string
	Directive string
}

type KDERestrictedLaunchAuthorizationRecord struct {
	SchemaVersion               string                                  `json:"schema_version"`
	RecordType                  string                                  `json:"record_type"`
	Source                      string                                  `json:"source"`
	Mode                        string                                  `json:"mode"`
	Application                 KDEFakeExecutionApplication             `json:"application"`
	Prerequisite                KDERestrictedLaunchPrerequisiteEvidence `json:"prerequisite"`
	Authorization               KDERestrictedPreparationAuthorization   `json:"authorization"`
	Execution                   KDEFakeExecutionTransaction             `json:"execution"`
	Session                     KDEFakeExecutionSession                 `json:"session"`
	FanOut                      ExecutionSessionFanOutEvidence          `json:"fan_out"`
	Checks                      []KDEFakeExecutionCheck                 `json:"checks"`
	CheckCount                  int                                     `json:"check_count"`
	PassedCheckCount            int                                     `json:"passed_check_count"`
	AllChecksPassed             bool                                    `json:"all_checks_passed"`
	CoreReceiptCount            int                                     `json:"core_receipt_count"`
	StateRootWritesEnabled      bool                                    `json:"state_root_writes_enabled"`
	StateRootWriteScope         string                                  `json:"state_root_write_scope"`
	StateRootPathExposed        bool                                    `json:"state_root_path_exposed"`
	AuthorizationBoundaryJoined bool                                    `json:"authorization_boundary_joined"`
	ProductionTrustSatisfied    bool                                    `json:"production_trust_satisfied"`
	RuntimeWriteGateEnabled     bool                                    `json:"runtime_write_gate_enabled"`
	ArtifactAcquisitionEnabled  bool                                    `json:"artifact_acquisition_enabled"`
	BackendInstallEnabled       bool                                    `json:"backend_install_enabled"`
	BackendLaunchEnabled        bool                                    `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                                    `json:"backend_process_started"`
	RealPortalCallEnabled       bool                                    `json:"real_portal_call_enabled"`
	ExecutionApproved           bool                                    `json:"execution_approved"`
	LaunchAuthorized            bool                                    `json:"launch_authorized"`
	LaunchAllowed               bool                                    `json:"launch_allowed"`
	LaunchEnabled               bool                                    `json:"launch_enabled"`
	ExecutionStarted            bool                                    `json:"execution_started"`
	ProcessStartAuthorized      bool                                    `json:"process_start_authorized"`
	ProductionBusOwnership      bool                                    `json:"production_bus_ownership"`
	NetworkRequired             bool                                    `json:"network_required"`
	HostRootModified            bool                                    `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                    `json:"privileged_container_required"`
	RawCommandExposed           bool                                    `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                                    `json:"backend_details_exposed"`
	DesktopSafeSummary          string                                  `json:"desktop_safe_summary"`
}

type KDERestrictedLaunchPrerequisiteEvidence struct {
	BackendStateJoined      bool `json:"backend_state_joined"`
	LifecycleReady          bool `json:"lifecycle_ready"`
	SnapshotVerified        bool `json:"snapshot_verified"`
	DiagnosticVerified      bool `json:"diagnostic_verified"`
	PortalReceiptCompleted  bool `json:"portal_receipt_completed"`
	ExecutionBlocked        bool `json:"execution_blocked"`
	BackendProcessesStopped bool `json:"backend_processes_stopped"`
	AllChecksPassed         bool `json:"all_checks_passed"`
}

type KDERestrictedPreparationAuthorization struct {
	AuthorizationID        string `json:"authorization_id"`
	ReceiptRelativePath    string `json:"receipt_relative_path"`
	ReceiptSHA256          string `json:"receipt_sha256"`
	State                  string `json:"state"`
	Scope                  string `json:"scope"`
	Directive              string `json:"directive"`
	ReceiptPersisted       bool   `json:"receipt_persisted"`
	ReceiptReadBack        bool   `json:"receipt_read_back"`
	PreparationAuthorized  bool   `json:"preparation_authorized"`
	LaunchAuthorized       bool   `json:"launch_authorized"`
	ProcessStartAuthorized bool   `json:"process_start_authorized"`
	ExecutionApproved      bool   `json:"execution_approved"`
}

func NewKDERestrictedLaunchAuthorizationRecord(recipeRecord Recipe, provenance Provenance, options KDERestrictedLaunchAuthorizationOptions) (KDERestrictedLaunchAuthorizationRecord, error) {
	if options.Mode != execution.RestrictedTestMode || options.Directive != execution.RestrictedTestPreparationDirective {
		return KDERestrictedLaunchAuthorizationRecord{}, errors.New("restricted launch authorization evidence requires the exact test-only mode and authorization directive")
	}
	prerequisite, err := NewKDEBackendLifecycleEvidenceRecord(recipeRecord, provenance, KDEFakeExecutionEvidenceOptions{StateRoot: options.StateRoot, Mode: options.Mode})
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	store, err := execution.NewRestrictedAuthorizationStore(options.StateRoot)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	recorded, err := store.Record(execution.RestrictedAuthorizationRequest{
		ApplicationID: prerequisite.Application.ID,
		RequestID:     prerequisite.Execution.RequestID,
		Mode:          options.Mode,
		Scope:         execution.RestrictedTestPreparationScope,
		Directive:     options.Directive,
	})
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	authorization, err := store.Load(recorded.AuthorizationID)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	ledger, err := execution.NewLedger(options.StateRoot)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	transactionRecord, err := ledger.Load(prerequisite.Execution.RequestID)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	sessionRecord, err := ledger.LoadSession(prerequisite.Execution.RequestID)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, prerequisite.Execution.RequestID)
	if err != nil {
		return KDERestrictedLaunchAuthorizationRecord{}, err
	}
	checks := kdeRestrictedLaunchAuthorizationChecks(prerequisite, authorization, transactionRecord, sessionRecord, fanOut)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDERestrictedLaunchAuthorizationRecord{}, errors.New("restricted launch authorization evidence checks did not all pass")
	}

	return KDERestrictedLaunchAuthorizationRecord{
		SchemaVersion: "xnix.runtime.kde_restricted_launch_authorization.v1",
		RecordType:    "kde-restricted-launch-authorization-record",
		Source:        "restricted-launch-authorization-receipt+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application:   prerequisite.Application,
		Prerequisite: KDERestrictedLaunchPrerequisiteEvidence{
			BackendStateJoined:      prerequisite.BackendStateJoined,
			LifecycleReady:          prerequisite.Lifecycle.State == "ready",
			SnapshotVerified:        prerequisite.Prerequisite.SnapshotVerified,
			DiagnosticVerified:      prerequisite.Prerequisite.DiagnosticVerified,
			PortalReceiptCompleted:  prerequisite.Prerequisite.PortalReceiptCompleted,
			ExecutionBlocked:        prerequisite.Execution.State == "blocked",
			BackendProcessesStopped: prerequisite.BackendManager.AllBackendProcessesStopped,
			AllChecksPassed:         prerequisite.AllChecksPassed,
		},
		Authorization: KDERestrictedPreparationAuthorization{
			AuthorizationID:        authorization.AuthorizationID,
			ReceiptRelativePath:    authorization.RelativePath,
			ReceiptSHA256:          authorization.SHA256,
			State:                  authorization.State,
			Scope:                  authorization.Scope,
			Directive:              authorization.Directive,
			ReceiptPersisted:       true,
			ReceiptReadBack:        true,
			PreparationAuthorized:  authorization.PreparationAuthorized,
			LaunchAuthorized:       false,
			ProcessStartAuthorized: false,
			ExecutionApproved:      false,
		},
		Execution:                   fakePortalExecutionSummary(transactionRecord),
		Session:                     fakePortalSessionSummary(sessionRecord),
		FanOut:                      fanOut,
		Checks:                      checks,
		CheckCount:                  len(checks),
		PassedCheckCount:            passed,
		AllChecksPassed:             true,
		CoreReceiptCount:            prerequisite.CoreReceiptCount + 1,
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		StateRootPathExposed:        false,
		AuthorizationBoundaryJoined: true,
		ProductionTrustSatisfied:    false,
		RuntimeWriteGateEnabled:     false,
		ArtifactAcquisitionEnabled:  false,
		BackendInstallEnabled:       false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		RealPortalCallEnabled:       false,
		ExecutionApproved:           false,
		LaunchAuthorized:            false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		ProcessStartAuthorized:      false,
		ProductionBusOwnership:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "Runtime recorded explicit permission for restricted test preparation while production trust, Runtime writes, launch authorization, execution, and process start remain blocked.",
	}, nil
}

func kdeRestrictedLaunchAuthorizationChecks(prerequisite KDEBackendLifecycleEvidenceRecord, authorization execution.RestrictedAuthorizationReceipt, transaction execution.LedgerRecord, session execution.SessionRecord, fanOut ExecutionSessionFanOutEvidence) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("prerequisite-convergence", prerequisite.AllChecksPassed && prerequisite.BackendStateJoined && prerequisite.Lifecycle.State == "ready" && prerequisite.Execution.State == "blocked", "The backend lifecycle prerequisite remains ready and execution remains blocked."),
		fakeExecutionCheck("explicit-test-boundary", authorization.Mode == execution.RestrictedTestMode && authorization.Scope == execution.RestrictedTestPreparationScope && authorization.Directive == execution.RestrictedTestPreparationDirective, "The receipt records the exact test-only preparation scope and explicit directive."),
		fakeExecutionCheck("authorization-readback", authorization.State == "authorized-preparation-only" && authorization.PreparationAuthorized && len(authorization.SHA256) == 64 && !filepath.IsAbs(authorization.RelativePath), "The preparation-only authorization passed digest-verified readback."),
		fakeExecutionCheck("trust-independent", fakeExecutionGateStatus(transaction.Transaction.Gates, "recipe-trust") == "pending" && !authorization.ProductionTrustSatisfied, "Test preparation authorization does not satisfy production recipe trust."),
		fakeExecutionCheck("write-gate-independent", fakeExecutionGateStatus(transaction.Transaction.Gates, "runtime-write-gate") == "blocked" && !authorization.RuntimeWriteGateEnabled, "Test preparation authorization does not enable the Runtime write gate."),
		fakeExecutionCheck("execution-unchanged", transaction.Transaction.State == execution.StateBlocked && fakePortalExecutionSummary(transaction).PassedGateCount == 4 && !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted, "Execution remains the same blocked transaction after authorization is recorded."),
		fakeExecutionCheck("session-unchanged", session.SessionState == "blocked" && !session.SessionActive && !session.ExecutionStarted && !session.BackendProcessStarted && fanOut.SafeForKDE && fanOut.SurfaceCount == 4, "The blocked session remains unchanged for all four KDE consumers."),
		fakeExecutionCheck("unsafe-gates-closed", !authorization.LaunchAuthorized && !authorization.ProcessStartAuthorized && !authorization.ExecutionApproved && !authorization.ArtifactAcquisitionEnabled && !authorization.BackendInstallEnabled && !authorization.BackendLaunchEnabled && !authorization.BackendProcessStarted && !authorization.RealPortalCallEnabled && !authorization.NetworkRequired && !authorization.PrivilegedContainerRequired && !authorization.HostRootModified, "Authorization does not enable acquisition, install, launch, process start, network, privilege, or host mutation."),
	}
}
