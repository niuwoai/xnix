package appidentity

import (
	"errors"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/execution"
)

// KDEBackendLifecycleEvidenceRecord joins the ready lifecycle, persisted
// backend inventory, and blocked execution state without exposing backend
// implementation details to KDE-facing consumers.
type KDEBackendLifecycleEvidenceRecord struct {
	SchemaVersion               string                         `json:"schema_version"`
	RecordType                  string                         `json:"record_type"`
	Source                      string                         `json:"source"`
	Mode                        string                         `json:"mode"`
	Application                 KDEFakeExecutionApplication    `json:"application"`
	Prerequisite                KDEBackendPrerequisiteEvidence `json:"prerequisite"`
	BackendManager              KDEBackendManagerEvidence      `json:"backend_manager"`
	Lifecycle                   KDEFakeExecutionLifecycle      `json:"lifecycle"`
	Execution                   KDEFakeExecutionTransaction    `json:"execution"`
	Session                     KDEFakeExecutionSession        `json:"session"`
	FanOut                      ExecutionSessionFanOutEvidence `json:"fan_out"`
	Checks                      []KDEFakeExecutionCheck        `json:"checks"`
	CheckCount                  int                            `json:"check_count"`
	PassedCheckCount            int                            `json:"passed_check_count"`
	AllChecksPassed             bool                           `json:"all_checks_passed"`
	CoreReceiptCount            int                            `json:"core_receipt_count"`
	StateRootWritesEnabled      bool                           `json:"state_root_writes_enabled"`
	StateRootWriteScope         string                         `json:"state_root_write_scope"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	BackendStateJoined          bool                           `json:"backend_state_joined"`
	BackendKindsExposedToKDE    bool                           `json:"backend_kinds_exposed_to_kde"`
	BackendInstallEnabled       bool                           `json:"backend_install_enabled"`
	BackendDownloadEnabled      bool                           `json:"backend_download_enabled"`
	BackendLaunchEnabled        bool                           `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                           `json:"backend_process_started"`
	VMProcessStarted            bool                           `json:"vm_process_started"`
	RawCommandExposed           bool                           `json:"raw_command_exposed"`
	ProfilePathExposed          bool                           `json:"profile_path_exposed"`
	RealPortalCallEnabled       bool                           `json:"real_portal_call_enabled"`
	SnapshotRestoreEnabled      bool                           `json:"snapshot_restore_enabled"`
	DiagnosticExecutionEnabled  bool                           `json:"diagnostic_execution_enabled"`
	AIProviderCallEnabled       bool                           `json:"ai_provider_call_enabled"`
	RepairExecutionEnabled      bool                           `json:"repair_execution_enabled"`
	ExecutionApproved           bool                           `json:"execution_approved"`
	LaunchAllowed               bool                           `json:"launch_allowed"`
	LaunchEnabled               bool                           `json:"launch_enabled"`
	ExecutionStarted            bool                           `json:"execution_started"`
	ProductionBusOwnership      bool                           `json:"production_bus_ownership"`
	NetworkRequired             bool                           `json:"network_required"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	SecretsExposed              bool                           `json:"secrets_exposed"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type KDEBackendPrerequisiteEvidence struct {
	PortalReceiptCompleted bool `json:"portal_receipt_completed"`
	DiagnosticVerified     bool `json:"diagnostic_verified"`
	DiagnosticRedacted     bool `json:"diagnostic_redacted"`
	SnapshotVerified       bool `json:"snapshot_verified"`
	LifecycleReady         bool `json:"lifecycle_ready"`
	AllChecksPassed        bool `json:"all_checks_passed"`
}

type KDEBackendManagerEvidence struct {
	ReceiptRelativePath        string `json:"receipt_relative_path"`
	ReceiptSHA256              string `json:"receipt_sha256"`
	InventoryPersisted         bool   `json:"inventory_persisted"`
	InventoryReadBack          bool   `json:"inventory_read_back"`
	ManagedBackendCount        int    `json:"managed_backend_count"`
	UserFacingProfileCount     int    `json:"user_facing_profile_count"`
	AllBackendsPlanned         bool   `json:"all_backends_planned"`
	AllBackendProcessesStopped bool   `json:"all_backend_processes_stopped"`
	RuntimeOwned               bool   `json:"runtime_owned"`
	KDEPolicyOwner             bool   `json:"kde_policy_owner"`
	KDEVisible                 bool   `json:"kde_visible"`
	BackendKindsExposedToKDE   bool   `json:"backend_kinds_exposed_to_kde"`
	InstallEnabled             bool   `json:"install_enabled"`
	DownloadEnabled            bool   `json:"download_enabled"`
	LaunchEnabled              bool   `json:"launch_enabled"`
	ProcessStarted             bool   `json:"process_started"`
	VMProcessStarted           bool   `json:"vm_process_started"`
}

func NewKDEBackendLifecycleEvidenceRecord(recipeRecord Recipe, provenance Provenance, options KDEFakeExecutionEvidenceOptions) (KDEBackendLifecycleEvidenceRecord, error) {
	if options.Mode != kdeFakeExecutionTestMode {
		return KDEBackendLifecycleEvidenceRecord{}, errors.New("backend lifecycle evidence requires test-only mode")
	}
	prerequisite, err := NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord, provenance, options)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	if _, err := RecordBackendManagerPreview(options.StateRoot); err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	managerRecord, err := LoadBackendManagerRecord(options.StateRoot)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	lifecycleRecord, err := plan.RecordBackendLifecycleState(options.StateRoot, "inspect", "", "", "")
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	ledger, err := execution.NewLedger(options.StateRoot)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	transactionRecord, err := ledger.Load(prerequisite.Execution.RequestID)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	sessionRecord, err := ledger.LoadSession(prerequisite.Execution.RequestID)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, prerequisite.Execution.RequestID)
	if err != nil {
		return KDEBackendLifecycleEvidenceRecord{}, err
	}

	checks := kdeBackendLifecycleEvidenceChecks(prerequisite, managerRecord, lifecycleRecord, transactionRecord, sessionRecord, fanOut)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDEBackendLifecycleEvidenceRecord{}, errors.New("backend lifecycle evidence checks did not all pass")
	}
	allPlanned, allStopped := managedBackendInventoryState(managerRecord.Preview.Backends)

	return KDEBackendLifecycleEvidenceRecord{
		SchemaVersion: "xnix.runtime.kde_backend_lifecycle_evidence.v1",
		RecordType:    "kde-backend-lifecycle-evidence-record",
		Source:        "backend-manager-inventory-record+backend-lifecycle-record+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application:   prerequisite.Application,
		Prerequisite: KDEBackendPrerequisiteEvidence{
			PortalReceiptCompleted: prerequisite.Portal.State == "completed" && prerequisite.Portal.PermissionState == "granted",
			DiagnosticVerified:     len(prerequisite.Diagnostic.ReceiptSHA256) == 64 && prerequisite.Diagnostic.Overall == "pass",
			DiagnosticRedacted:     prerequisite.Diagnostic.RedactionVerified && !prerequisite.Diagnostic.FileContentsIncluded && !prerequisite.Diagnostic.FilePathsExposed,
			SnapshotVerified:       prerequisite.Snapshot.Verified && prerequisite.Snapshot.BaselinePresent,
			LifecycleReady:         prerequisite.Lifecycle.State == "ready" && len(prerequisite.Lifecycle.PendingGates) == 0,
			AllChecksPassed:        prerequisite.AllChecksPassed,
		},
		BackendManager: KDEBackendManagerEvidence{
			ReceiptRelativePath:        managerRecord.RelativePath,
			ReceiptSHA256:              managerRecord.SHA256,
			InventoryPersisted:         true,
			InventoryReadBack:          true,
			ManagedBackendCount:        managerRecord.Preview.BackendCount,
			UserFacingProfileCount:     managerRecord.Preview.UserFacingProfileCount,
			AllBackendsPlanned:         allPlanned,
			AllBackendProcessesStopped: allStopped,
			RuntimeOwned:               managerRecord.RuntimeOwned,
			KDEPolicyOwner:             false,
			KDEVisible:                 false,
			BackendKindsExposedToKDE:   false,
			InstallEnabled:             false,
			DownloadEnabled:            false,
			LaunchEnabled:              false,
			ProcessStarted:             false,
			VMProcessStarted:           false,
		},
		Lifecycle:                   backendLifecycleEvidenceSummary(lifecycleRecord),
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
		BackendStateJoined:          true,
		BackendKindsExposedToKDE:    false,
		BackendInstallEnabled:       false,
		BackendDownloadEnabled:      false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		VMProcessStarted:            false,
		RawCommandExposed:           false,
		ProfilePathExposed:          false,
		RealPortalCallEnabled:       false,
		SnapshotRestoreEnabled:      false,
		DiagnosticExecutionEnabled:  false,
		AIProviderCallEnabled:       false,
		RepairExecutionEnabled:      false,
		ExecutionApproved:           false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		ProductionBusOwnership:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		SecretsExposed:              false,
		DesktopSafeSummary:          "Runtime joined one verified compatibility inventory to the ready environment and blocked execution records without exposing backend kinds to KDE or starting a compatibility process.",
	}, nil
}

func backendLifecycleEvidenceSummary(record BackendLifecycleRecord) KDEFakeExecutionLifecycle {
	return KDEFakeExecutionLifecycle{
		ReceiptRelativePath: record.RelativePath,
		State:               record.Preview.LifecycleState,
		OverallStatus:       record.Preview.OverallStatus,
		SatisfiedGates:      append([]string{}, record.Preview.SatisfiedGates...),
		PendingGates:        append([]string{}, record.Preview.PendingGates...),
		ReceiptPersisted:    true,
		StateRootBacked:     true,
		LaunchEnabled:       false,
		ProcessStarted:      false,
	}
}

func kdeBackendLifecycleEvidenceChecks(prerequisite KDESnapshotDiagnosticsEvidenceRecord, manager BackendManagerRecord, lifecycle BackendLifecycleRecord, transaction execution.LedgerRecord, session execution.SessionRecord, fanOut ExecutionSessionFanOutEvidence) []KDEFakeExecutionCheck {
	allPlanned, allStopped := managedBackendInventoryState(manager.Preview.Backends)
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("prerequisite-convergence", prerequisite.AllChecksPassed && prerequisite.Snapshot.Verified && prerequisite.Diagnostic.RedactionVerified && prerequisite.Lifecycle.State == "ready", "Portal, diagnostic, snapshot, and lifecycle prerequisite evidence remains converged."),
		fakeExecutionCheck("inventory-readback", manager.RelativePath == filepath.ToSlash(filepath.Join("backend-manager", "inventory.json")) && len(manager.SHA256) == 64 && manager.Preview.BackendCount == 3 && manager.Preview.UserFacingProfileCount == 3, "The fixed backend inventory receipt passed digest-verified readback."),
		fakeExecutionCheck("backend-state", allPlanned && allStopped, "Every managed backend remains planned and every compatibility process remains stopped."),
		fakeExecutionCheck("lifecycle-join", lifecycle.EnvironmentRecord.Ready() && lifecycle.Preview.LifecycleState == "ready" && len(lifecycle.Preview.PendingGates) == 0 && !lifecycle.BackendProcessStarted, "The persisted inventory joins a ready lifecycle without implying process readiness."),
		fakeExecutionCheck("execution-readback", transaction.Transaction.State == execution.StateBlocked && fakePortalExecutionSummary(transaction).PassedGateCount == 4 && !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.Transaction.BackendStarted, "The digest-verified execution receipt remains blocked by trust and Runtime write gates."),
		fakeExecutionCheck("session-readback", session.SessionState == "blocked" && !session.SessionActive && !session.ExecutionStarted && !session.BackendProcessStarted && fanOut.SafeForKDE && fanOut.SurfaceCount == 4, "The blocked session remains safe for four KDE consumers."),
		fakeExecutionCheck("backend-boundary", !manager.Preview.KDEVisible && !manager.BackendInstallEnabled && !manager.BackendDownloadEnabled && !manager.BackendLaunchEnabled && !manager.BackendProcessStarted && !manager.VMProcessStarted && !manager.BackendDetailsExposedToKDE, "Backend inventory stays Runtime-owned and hidden from KDE policy surfaces."),
		fakeExecutionCheck("unsafe-gates-closed", !manager.HostRootModified && !manager.NetworkRequired && !manager.PrivilegedContainerRequired && !manager.SecretsExposed && !lifecycle.LaunchEnabled && !transaction.BackendStarted && !session.BackendProcessStarted, "Install, download, launch, process start, network, privilege, secrets, and host mutation remain disabled."),
	}
}

func managedBackendInventoryState(backends []ManagedCompatibilityBackend) (bool, bool) {
	if len(backends) == 0 {
		return false, false
	}
	allPlanned := true
	allStopped := true
	for _, backend := range backends {
		if backend.Status != "planned" || backend.Readiness != "blocked-until-runtime-gates-pass" {
			allPlanned = false
		}
		if backend.InstallEnabled || backend.DownloadEnabled || backend.LaunchEnabled || backend.ProcessStarted {
			allStopped = false
		}
	}
	return allPlanned, allStopped
}
