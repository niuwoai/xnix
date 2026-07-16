package appidentity

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/diagnostics"
	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/portal"
	"xnix.local/xnix/internal/runtime/snapshot"
)

const (
	kdeStabilitySnapshotID = "stability-baseline-v0.2.315"
	kdeStabilityRunID      = "stability-diagnostics-v0.2.315"
)

type KDESnapshotDiagnosticsEvidenceRecord struct {
	SchemaVersion               string                         `json:"schema_version"`
	RecordType                  string                         `json:"record_type"`
	Source                      string                         `json:"source"`
	Mode                        string                         `json:"mode"`
	Application                 KDEFakeExecutionApplication    `json:"application"`
	Portal                      KDESnapshotPortalEvidence      `json:"portal"`
	Diagnostic                  KDESnapshotDiagnosticEvidence  `json:"diagnostic"`
	Snapshot                    KDESnapshotBaselineEvidence    `json:"snapshot"`
	Lifecycle                   KDEFakeExecutionLifecycle      `json:"lifecycle"`
	Execution                   KDEFakeExecutionTransaction    `json:"execution"`
	Session                     KDEFakeExecutionSession        `json:"session"`
	FanOut                      ExecutionSessionFanOutEvidence `json:"fan_out"`
	Checks                      []KDEFakeExecutionCheck        `json:"checks"`
	CheckCount                  int                            `json:"check_count"`
	PassedCheckCount            int                            `json:"passed_check_count"`
	AllChecksPassed             bool                           `json:"all_checks_passed"`
	StateRootWritesEnabled      bool                           `json:"state_root_writes_enabled"`
	StateRootWriteScope         string                         `json:"state_root_write_scope"`
	CoreReceiptCount            int                            `json:"core_receipt_count"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	SnapshotCreationEnabled     bool                           `json:"snapshot_creation_enabled"`
	SnapshotRestoreEnabled      bool                           `json:"snapshot_restore_enabled"`
	SnapshotDeletionEnabled     bool                           `json:"snapshot_deletion_enabled"`
	DiagnosticExecutionEnabled  bool                           `json:"diagnostic_execution_enabled"`
	AIProviderCallEnabled       bool                           `json:"ai_provider_call_enabled"`
	RepairExecutionEnabled      bool                           `json:"repair_execution_enabled"`
	RealPortalCallEnabled       bool                           `json:"real_portal_call_enabled"`
	HostPermissionChanged       bool                           `json:"host_permission_changed"`
	ExecutionApproved           bool                           `json:"execution_approved"`
	LaunchAllowed               bool                           `json:"launch_allowed"`
	LaunchEnabled               bool                           `json:"launch_enabled"`
	ExecutionStarted            bool                           `json:"execution_started"`
	BackendProcessStarted       bool                           `json:"backend_process_started"`
	ProductionBusOwnership      bool                           `json:"production_bus_ownership"`
	NetworkRequired             bool                           `json:"network_required"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	FileContentsExposed         bool                           `json:"file_contents_exposed"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type KDESnapshotPortalEvidence struct {
	ReceiptRelativePath   string `json:"receipt_relative_path"`
	ReceiptSHA256         string `json:"receipt_sha256"`
	State                 string `json:"state"`
	PermissionState       string `json:"permission_state"`
	ReceiptPersisted      bool   `json:"receipt_persisted"`
	RealPortalCallEnabled bool   `json:"real_portal_call_enabled"`
	HostPermissionChanged bool   `json:"host_permission_changed"`
	ExecutionApproved     bool   `json:"execution_approved"`
}

type KDESnapshotDiagnosticEvidence struct {
	RunID                 string   `json:"run_id"`
	ReceiptRelativePath   string   `json:"receipt_relative_path"`
	ReceiptSHA256         string   `json:"receipt_sha256"`
	TestType              string   `json:"test_type"`
	Overall               string   `json:"overall"`
	SignalCount           int      `json:"signal_count"`
	SignalIDs             []string `json:"signal_ids"`
	ReceiptPersisted      bool     `json:"receipt_persisted"`
	RedactionVerified     bool     `json:"redaction_verified"`
	FileContentsIncluded  bool     `json:"file_contents_included"`
	FilePathsExposed      bool     `json:"file_paths_exposed"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	RealAIProviderEnabled bool     `json:"real_ai_provider_enabled"`
	AutoRepairAllowed     bool     `json:"auto_repair_allowed"`
	RepairExecuted        bool     `json:"repair_executed"`
}

type KDESnapshotBaselineEvidence struct {
	SnapshotID           string `json:"snapshot_id"`
	ManifestRelativePath string `json:"manifest_relative_path"`
	ContentHash          string `json:"content_hash"`
	FileCount            int    `json:"file_count"`
	ObjectCount          int    `json:"object_count"`
	SnapshotCount        int    `json:"snapshot_count"`
	Created              bool   `json:"created"`
	Reused               bool   `json:"reused"`
	Verified             bool   `json:"verified"`
	BaselinePresent      bool   `json:"baseline_present"`
	RestoreEnabled       bool   `json:"restore_enabled"`
	DeletionEnabled      bool   `json:"deletion_enabled"`
	HostRootModified     bool   `json:"host_root_modified"`
}

func NewKDESnapshotDiagnosticsEvidenceRecord(recipeRecord Recipe, provenance Provenance, options KDEFakeExecutionEvidenceOptions) (KDESnapshotDiagnosticsEvidenceRecord, error) {
	if options.Mode != kdeFakeExecutionTestMode {
		return KDESnapshotDiagnosticsEvidenceRecord{}, errors.New("snapshot diagnostics evidence requires test-only mode")
	}
	if err := validateKDEFakeExecutionStateRoot(options.StateRoot); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if err := validateKDEFakeExecutionManagedPaths(options.StateRoot); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if err := validateKDEFakePortalManagedPath(options.StateRoot); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if err := validateKDESnapshotDiagnosticsManagedPaths(options.StateRoot); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if !provenance.DigestVerified {
		return KDESnapshotDiagnosticsEvidenceRecord{}, errors.New("snapshot diagnostics evidence requires a digest-verified registry recipe")
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if _, err := prepareFakeExecutionLifecycle(plan, options.StateRoot); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	portalRecord, _, err := ensureCompletedFakePortalReceipt(options.StateRoot, plan.ApplicationID)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if err := validateCompletedFakePortalRecord(portalRecord, plan.ApplicationID); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	portalDigest, err := fakePortalReceiptDigest(options.StateRoot, portalRecord.RelativePath)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	lifecycleStore, err := environment.New(options.StateRoot)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	profile := plan.backendLifecycleProfile()
	if _, err := lifecycleStore.SatisfyGate(plan.ApplicationID, profile, "portal-policy-review"); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	preSnapshotEnvironment, err := lifecycleStore.Get(plan.ApplicationID, profile)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	trust, err := fakeExecutionTrust(recipeRecord.ID, provenance)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	portalReceipt := fakePortalExecutionReceipt(portalRecord)
	preSnapshotTransaction, err := reviewedFakeExecutionTransaction(plan.ApplicationID, execution.Inputs{
		Trust:                    trust,
		Environment:              preSnapshotEnvironment,
		SnapshotBaselinePresent:  false,
		PortalRequiredOps:        []string{"file-open"},
		PortalPermissionReceipts: []execution.PortalPermissionReceipt{portalReceipt},
	})
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if _, _, err := persistFakeExecutionTransaction(options.StateRoot, preSnapshotTransaction); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}

	diagnosticRecord, err := recordStabilityDiagnostic(options.StateRoot, plan.ApplicationID)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	manifest, baseline, snapshotReused, err := ensureStabilitySnapshot(options.StateRoot)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if _, err := lifecycleStore.SatisfyGate(plan.ApplicationID, profile, "snapshot-baseline"); err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	currentEnvironment, err := lifecycleStore.Get(plan.ApplicationID, profile)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	if currentEnvironment.State == environment.StateStaged {
		if _, err := lifecycleStore.MarkReady(plan.ApplicationID, profile); err != nil {
			return KDESnapshotDiagnosticsEvidenceRecord{}, err
		}
	}
	lifecycleRecord, err := plan.RecordBackendLifecycleState(options.StateRoot, "inspect", "", "", "")
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	finalTransaction, err := reviewedFakeExecutionTransaction(plan.ApplicationID, execution.Inputs{
		Trust:                    trust,
		Environment:              lifecycleRecord.EnvironmentRecord,
		SnapshotBaselinePresent:  baseline.Present && baseline.Verified,
		PortalRequiredOps:        []string{"file-open"},
		PortalPermissionReceipts: []execution.PortalPermissionReceipt{portalReceipt},
	})
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	transactionRecord, sessionRecord, err := persistFakeExecutionTransaction(options.StateRoot, finalTransaction)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, finalTransaction.RequestID)
	if err != nil {
		return KDESnapshotDiagnosticsEvidenceRecord{}, err
	}

	checks := kdeSnapshotDiagnosticsChecks(portalRecord, diagnosticRecord, manifest, baseline, lifecycleRecord, transactionRecord, sessionRecord, fanOut)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDESnapshotDiagnosticsEvidenceRecord{}, errors.New("snapshot diagnostics evidence checks did not all pass")
	}
	manifestPath := filepath.ToSlash(filepath.Join(".xnix-snapshots", "manifests", manifest.ID+".json"))

	return KDESnapshotDiagnosticsEvidenceRecord{
		SchemaVersion: "xnix.runtime.kde_snapshot_diagnostics_evidence.v1",
		RecordType:    "kde-snapshot-diagnostics-evidence-record",
		Source:        "state-root-portal-request+diagnostic-run-record+content-addressed-snapshot+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application:   KDEFakeExecutionApplication{ID: plan.ApplicationID, Name: plan.DisplayName, Icon: plan.Icon, DesktopFile: plan.DesktopFile, IdentityDigest: plan.StableIdentityDigest},
		Portal: KDESnapshotPortalEvidence{
			ReceiptRelativePath:   portalRecord.RelativePath,
			ReceiptSHA256:         portalDigest,
			State:                 string(portalRecord.Request.State),
			PermissionState:       string(portalRecord.Request.PermissionState),
			ReceiptPersisted:      true,
			RealPortalCallEnabled: false,
			HostPermissionChanged: false,
			ExecutionApproved:     false,
		},
		Diagnostic: KDESnapshotDiagnosticEvidence{
			RunID:                 diagnosticRecord.RunID,
			ReceiptRelativePath:   diagnosticRecord.RelativePath,
			ReceiptSHA256:         diagnosticRecord.SHA256,
			TestType:              diagnosticRecord.Result.TestType,
			Overall:               string(diagnosticRecord.Result.Overall),
			SignalCount:           len(diagnosticRecord.Result.Signals),
			SignalIDs:             diagnosticSignalIDs(diagnosticRecord.Result.Signals),
			ReceiptPersisted:      true,
			RedactionVerified:     true,
			FileContentsIncluded:  false,
			FilePathsExposed:      false,
			AIProviderCalled:      false,
			RealAIProviderEnabled: false,
			AutoRepairAllowed:     false,
			RepairExecuted:        false,
		},
		Snapshot: KDESnapshotBaselineEvidence{
			SnapshotID:           manifest.ID,
			ManifestRelativePath: manifestPath,
			ContentHash:          manifest.ContentHash,
			FileCount:            manifest.FileCount,
			ObjectCount:          uniqueSnapshotObjectCount(manifest.Files),
			SnapshotCount:        baseline.SnapshotCount,
			Created:              !snapshotReused,
			Reused:               snapshotReused,
			Verified:             baseline.Verified,
			BaselinePresent:      baseline.Present,
			RestoreEnabled:       false,
			DeletionEnabled:      false,
			HostRootModified:     false,
		},
		Lifecycle: KDEFakeExecutionLifecycle{
			ReceiptRelativePath: lifecycleRecord.RelativePath,
			State:               lifecycleRecord.Preview.LifecycleState,
			OverallStatus:       lifecycleRecord.Preview.OverallStatus,
			SatisfiedGates:      append([]string{}, lifecycleRecord.Preview.SatisfiedGates...),
			PendingGates:        append([]string{}, lifecycleRecord.Preview.PendingGates...),
			ReceiptPersisted:    true,
			StateRootBacked:     true,
			LaunchEnabled:       false,
			ProcessStarted:      false,
		},
		Execution:                   fakePortalExecutionSummary(transactionRecord),
		Session:                     fakePortalSessionSummary(sessionRecord),
		FanOut:                      fanOut,
		Checks:                      checks,
		CheckCount:                  len(checks),
		PassedCheckCount:            passed,
		AllChecksPassed:             true,
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		CoreReceiptCount:            7,
		StateRootPathExposed:        false,
		SnapshotCreationEnabled:     true,
		SnapshotRestoreEnabled:      false,
		SnapshotDeletionEnabled:     false,
		DiagnosticExecutionEnabled:  false,
		AIProviderCallEnabled:       false,
		RepairExecutionEnabled:      false,
		RealPortalCallEnabled:       false,
		HostPermissionChanged:       false,
		ExecutionApproved:           false,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		BackendProcessStarted:       false,
		ProductionBusOwnership:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		FileContentsExposed:         false,
		DesktopSafeSummary:          "Runtime verified one redacted diagnostic receipt and one content-addressed restore-point baseline, then marked lifecycle readiness without enabling launch, restore, repair, or live diagnostics.",
	}, nil
}

func validateKDESnapshotDiagnosticsManagedPaths(root string) error {
	for _, relativePath := range []string{".xnix-snapshots", "diagnostics-ledger", "test-results"} {
		info, err := os.Lstat(filepath.Join(root, relativePath))
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("inspect snapshot diagnostics managed path: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("snapshot diagnostics managed path must be a real directory: %s", relativePath)
		}
	}
	return nil
}

func recordStabilityDiagnostic(root string, applicationID string) (diagnostics.RunRecord, error) {
	store, err := diagnostics.NewRunRecordStore(root)
	if err != nil {
		return diagnostics.RunRecord{}, err
	}
	fixture := diagnostics.Fixture{TestType: "preflight", Signals: []diagnostics.Signal{
		{ID: "recipe-identity", Category: "trust", Outcome: diagnostics.OutcomePass, Summary: "Recipe identity evidence is verified."},
		{ID: "portal-receipt", Category: "permission", Outcome: diagnostics.OutcomePass, Summary: "Desktop permission evidence is recorded."},
		{ID: "state-root-integrity", Category: "storage", Outcome: diagnostics.OutcomePass, Summary: "Controlled state evidence is internally consistent."},
	}}
	if _, err := store.Record(diagnostics.RunRecordRequest{ApplicationID: applicationID, RunID: kdeStabilityRunID, Fixture: fixture}); err != nil {
		return diagnostics.RunRecord{}, err
	}
	return store.Load(kdeStabilityRunID)
}

func ensureStabilitySnapshot(root string) (snapshot.Manifest, snapshot.BaselineStatus, bool, error) {
	store, err := snapshot.New(root)
	if err != nil {
		return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
	}
	manifests, err := store.List()
	if err != nil {
		return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
	}
	var manifest snapshot.Manifest
	reused := false
	for _, candidate := range manifests {
		if candidate.ID == kdeStabilitySnapshotID {
			manifest = candidate
			reused = true
			break
		}
	}
	if reused {
		if err := store.Verify(manifest.ID); err != nil {
			return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
		}
	} else {
		manifest, err = store.Create(kdeStabilitySnapshotID, "Controlled pre-execution stability baseline")
		if err != nil {
			return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
		}
		if err := store.Verify(manifest.ID); err != nil {
			return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
		}
	}
	baseline, err := store.Baseline()
	if err != nil {
		return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, err
	}
	if !baseline.Present || !baseline.Verified || baseline.SnapshotID != manifest.ID {
		return snapshot.Manifest{}, snapshot.BaselineStatus{}, false, errors.New("stability snapshot did not become the verified baseline")
	}
	return manifest, baseline, reused, nil
}

func kdeSnapshotDiagnosticsChecks(portalRecord portal.Record, diagnostic diagnostics.RunRecord, manifest snapshot.Manifest, baseline snapshot.BaselineStatus, lifecycle BackendLifecycleRecord, transaction execution.LedgerRecord, session execution.SessionRecord, fanOut ExecutionSessionFanOutEvidence) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("portal-evidence", portalRecord.Request.State == portal.StateCompleted && portalRecord.Request.PermissionState == portal.PermissionGranted && portalRecord.PermissionGranted && !portalRecord.RealPortalCallEnabled && !portalRecord.ExecutionApproved && !portalRecord.HostPermissionChanged && transaction.PortalPermissionReceiptCount == 1 && fakeExecutionGateStatus(transaction.Transaction.Gates, "portal-permission") == "pass", "The final execution receipt consumes one completed fake Portal permission receipt."),
		fakeExecutionCheck("diagnostic-readback", diagnostic.SchemaVersion == "xnix.runtime.diagnostic_run_record.v1" && len(diagnostic.SHA256) == 64 && diagnostic.Result.Overall == diagnostics.OutcomePass, "The redacted diagnostic receipt passed digest-verified readback."),
		fakeExecutionCheck("diagnostic-redaction", !diagnostic.StateRootPathExposed && !diagnostic.FixturePathExposed && !diagnostic.FileContentsIncluded && !diagnostic.AIProviderCalled && !diagnostic.RepairExecuted, "Diagnostic output contains no roots, fixture paths, file contents, provider calls, or repairs."),
		fakeExecutionCheck("snapshot-baseline", manifest.ID == kdeStabilitySnapshotID && len(manifest.ContentHash) == 64 && baseline.Present && baseline.Verified, "The content-addressed snapshot is the verified restore-point baseline."),
		fakeExecutionCheck("lifecycle-ready", lifecycle.EnvironmentRecord.Ready() && len(lifecycle.EnvironmentRecord.PendingGates()) == 0, "All lifecycle evidence gates are satisfied while launch remains disabled."),
		fakeExecutionCheck("execution-readiness", fakeExecutionGateStatus(transaction.Transaction.Gates, "environment-ready") == "pass" && fakeExecutionGateStatus(transaction.Transaction.Gates, "snapshot-baseline") == "pass" && fakeExecutionGateStatus(transaction.Transaction.Gates, "portal-permission") == "pass" && fakeExecutionGateStatus(transaction.Transaction.Gates, "recipe-trust") == "pending" && fakeExecutionGateStatus(transaction.Transaction.Gates, "runtime-write-gate") == "blocked", "Environment, snapshot, and Portal gates pass while trust and Runtime write gates remain closed."),
		fakeExecutionCheck("session-readback", session.SessionState == "blocked" && len(session.PortalPermissionReceiptPaths) == 1 && fanOut.SafeForKDE, "The blocked session and KDE fan-out consume the converged evidence."),
		fakeExecutionCheck("unsafe-gates-closed", !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted && !session.ExecutionStarted && !session.BackendProcessStarted && !diagnostic.RealAIProviderEnabled && !diagnostic.AutoRepairAllowed && !diagnostic.HostRootModified && !baseline.HostRootModified, "Launch, process start, real diagnostics, repair, restore, and host mutation remain disabled."),
	}
}

func diagnosticSignalIDs(signals []diagnostics.Signal) []string {
	ids := make([]string, 0, len(signals))
	for _, signal := range signals {
		ids = append(ids, signal.ID)
	}
	return ids
}

func uniqueSnapshotObjectCount(files []snapshot.FileEntry) int {
	digests := map[string]bool{}
	for _, file := range files {
		digests[file.Digest] = true
	}
	return len(digests)
}
