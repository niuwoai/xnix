package appidentity

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/recipe"
)

const kdeFakeExecutionTestMode = "test-only"

type KDEFakeExecutionEvidenceOptions struct {
	StateRoot string
	Mode      string
}

type KDEFakeExecutionEvidenceRecord struct {
	SchemaVersion               string                         `json:"schema_version"`
	RecordType                  string                         `json:"record_type"`
	Source                      string                         `json:"source"`
	Mode                        string                         `json:"mode"`
	Application                 KDEFakeExecutionApplication    `json:"application"`
	RecipeTrust                 KDEFakeExecutionRecipeTrust    `json:"recipe_trust"`
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
	StateRootRecordCount        int                            `json:"state_root_record_count"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	FakeExecutionRecorded       bool                           `json:"fake_execution_recorded"`
	LaunchAllowed               bool                           `json:"launch_allowed"`
	LaunchEnabled               bool                           `json:"launch_enabled"`
	ExecutionStarted            bool                           `json:"execution_started"`
	BackendProcessStarted       bool                           `json:"backend_process_started"`
	RealPortalCallEnabled       bool                           `json:"real_portal_call_enabled"`
	ProductionBusOwnership      bool                           `json:"production_bus_ownership"`
	NetworkRequired             bool                           `json:"network_required"`
	HostRootModified            bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type KDEFakeExecutionApplication struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	DesktopFile    string `json:"desktop_file"`
	IdentityDigest string `json:"identity_digest"`
}

type KDEFakeExecutionRecipeTrust struct {
	DigestVerified  bool   `json:"digest_verified"`
	SignatureStatus string `json:"signature_status"`
	ProductionReady bool   `json:"production_ready"`
}

type KDEFakeExecutionLifecycle struct {
	ReceiptRelativePath string   `json:"receipt_relative_path"`
	State               string   `json:"state"`
	OverallStatus       string   `json:"overall_status"`
	SatisfiedGates      []string `json:"satisfied_gates"`
	PendingGates        []string `json:"pending_gates"`
	ReceiptPersisted    bool     `json:"receipt_persisted"`
	StateRootBacked     bool     `json:"state_root_backed"`
	LaunchEnabled       bool     `json:"launch_enabled"`
	ProcessStarted      bool     `json:"process_started"`
}

type KDEFakeExecutionTransaction struct {
	RequestID           string           `json:"request_id"`
	ReceiptRelativePath string           `json:"receipt_relative_path"`
	ReceiptSHA256       string           `json:"receipt_sha256"`
	State               string           `json:"state"`
	ReviewDecision      string           `json:"review_decision"`
	Gates               []execution.Gate `json:"gates"`
	GateCount           int              `json:"gate_count"`
	PassedGateCount     int              `json:"passed_gate_count"`
	BlockedReasons      []string         `json:"blocked_reasons"`
	ReceiptPersisted    bool             `json:"receipt_persisted"`
	LaunchAllowed       bool             `json:"launch_allowed"`
	LaunchEnabled       bool             `json:"launch_enabled"`
	ExecutionStarted    bool             `json:"execution_started"`
	ProcessStarted      bool             `json:"process_started"`
}

type KDEFakeExecutionSession struct {
	ReceiptRelativePath      string   `json:"receipt_relative_path"`
	ReceiptSHA256            string   `json:"receipt_sha256"`
	State                    string   `json:"state"`
	TaskManagerState         string   `json:"task_manager_state"`
	TrayState                string   `json:"tray_state"`
	KWinState                string   `json:"kwin_state"`
	CompatibilityCenterState string   `json:"compatibility_center_state"`
	BlockedReasons           []string `json:"blocked_reasons"`
	ReceiptPersisted         bool     `json:"receipt_persisted"`
	LiveStateObserved        bool     `json:"live_state_observed"`
	SessionActive            bool     `json:"session_active"`
}

type KDEFakeExecutionCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func NewKDEFakeExecutionEvidenceRecord(recipeRecord Recipe, provenance Provenance, options KDEFakeExecutionEvidenceOptions) (KDEFakeExecutionEvidenceRecord, error) {
	if options.Mode != kdeFakeExecutionTestMode {
		return KDEFakeExecutionEvidenceRecord{}, errors.New("fake execution evidence requires test-only mode")
	}
	if err := validateKDEFakeExecutionStateRoot(options.StateRoot); err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	if err := validateKDEFakeExecutionManagedPaths(options.StateRoot); err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	if !provenance.DigestVerified {
		return KDEFakeExecutionEvidenceRecord{}, errors.New("fake execution evidence requires a digest-verified registry recipe")
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}

	lifecycleRecord, err := prepareFakeExecutionLifecycle(plan, options.StateRoot)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	trust, err := fakeExecutionTrust(recipeRecord.ID, provenance)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	inputs := execution.Inputs{
		Trust:                   trust,
		Environment:             lifecycleRecord.EnvironmentRecord,
		SnapshotBaselinePresent: false,
		PortalRequiredOps:       []string{"file-open"},
		PortalGrantedOps:        []string{},
	}
	pipeline := execution.NewPipeline()
	tx, err := pipeline.Create(plan.ApplicationID, inputs)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	tx, err = tx.Review(execution.DecisionApproved)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	tx = tx.Preflight(inputs)
	ledger, err := execution.NewLedger(options.StateRoot)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	if _, err := ledger.Record(tx); err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	if _, err := ledger.RecordSession(tx.RequestID); err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}

	loadedTransaction, err := ledger.Load(tx.RequestID)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	sessionRecord, err := ledger.LoadSession(tx.RequestID)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, tx.RequestID)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	checks, err := fakeExecutionEvidenceChecks(options.StateRoot, plan, lifecycleRecord, loadedTransaction, sessionRecord, fanOut)
	if err != nil {
		return KDEFakeExecutionEvidenceRecord{}, err
	}
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}

	return KDEFakeExecutionEvidenceRecord{
		SchemaVersion: "xnix.runtime.kde_fake_execution_evidence.v1",
		RecordType:    "kde-fake-execution-evidence-record",
		Source:        "digest-verified-registry+environment-state-root+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application: KDEFakeExecutionApplication{
			ID:             plan.ApplicationID,
			Name:           plan.DisplayName,
			Icon:           plan.Icon,
			DesktopFile:    plan.DesktopFile,
			IdentityDigest: plan.StableIdentityDigest,
		},
		RecipeTrust: KDEFakeExecutionRecipeTrust{
			DigestVerified:  provenance.DigestVerified,
			SignatureStatus: provenance.SignatureStatus,
			ProductionReady: trust.ProductionTrusted,
		},
		Lifecycle: KDEFakeExecutionLifecycle{
			ReceiptRelativePath: lifecycleRecord.RelativePath,
			State:               lifecycleRecord.Preview.LifecycleState,
			OverallStatus:       lifecycleRecord.Preview.OverallStatus,
			SatisfiedGates:      append([]string{}, lifecycleRecord.Preview.SatisfiedGates...),
			PendingGates:        append([]string{}, lifecycleRecord.Preview.PendingGates...),
			ReceiptPersisted:    true,
			StateRootBacked:     lifecycleRecord.Preview.StateRootBacked,
			LaunchEnabled:       false,
			ProcessStarted:      false,
		},
		Execution: KDEFakeExecutionTransaction{
			RequestID:           loadedTransaction.RequestID,
			ReceiptRelativePath: loadedTransaction.RelativePath,
			ReceiptSHA256:       loadedTransaction.SHA256,
			State:               string(loadedTransaction.Transaction.State),
			ReviewDecision:      string(loadedTransaction.Transaction.ReviewDecision),
			Gates:               append([]execution.Gate{}, loadedTransaction.Transaction.Gates...),
			GateCount:           len(loadedTransaction.Transaction.Gates),
			PassedGateCount:     passedExecutionGateCount(loadedTransaction.Transaction.Gates),
			BlockedReasons:      append([]string{}, loadedTransaction.Transaction.BlockedReasons...),
			ReceiptPersisted:    true,
			LaunchAllowed:       false,
			LaunchEnabled:       false,
			ExecutionStarted:    false,
			ProcessStarted:      false,
		},
		Session: KDEFakeExecutionSession{
			ReceiptRelativePath:      sessionRecord.RelativePath,
			ReceiptSHA256:            sessionRecord.SHA256,
			State:                    sessionRecord.SessionState,
			TaskManagerState:         sessionRecord.TaskManagerState,
			TrayState:                sessionRecord.TrayState,
			KWinState:                sessionRecord.KWinState,
			CompatibilityCenterState: sessionRecord.CompatibilityCenterState,
			BlockedReasons:           append([]string{}, sessionRecord.BlockedReasons...),
			ReceiptPersisted:         true,
			LiveStateObserved:        false,
			SessionActive:            false,
		},
		FanOut:                      fanOut,
		Checks:                      checks,
		CheckCount:                  len(checks),
		PassedCheckCount:            passed,
		AllChecksPassed:             passed == len(checks),
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		StateRootRecordCount:        3,
		StateRootPathExposed:        false,
		FakeExecutionRecorded:       true,
		LaunchAllowed:               false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		BackendProcessStarted:       false,
		RealPortalCallEnabled:       false,
		ProductionBusOwnership:      false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "Runtime persisted and read back controlled lifecycle, blocked execution, and session evidence for one KDE application without starting a compatibility service or modifying the host root.",
	}, nil
}

func validateKDEFakeExecutionStateRoot(root string) error {
	root = strings.TrimSpace(root)
	if root == "" {
		return errors.New("fake execution evidence requires an explicit state root")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("resolve fake execution state root: %w", err)
	}
	if abs == string(os.PathSeparator) {
		return errors.New("refusing to use filesystem root for fake execution evidence")
	}
	info, err := os.Lstat(abs)
	if err != nil {
		return fmt.Errorf("fake execution state root must exist: %w", err)
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return errors.New("fake execution state root must be a real directory")
	}
	return nil
}

func validateKDEFakeExecutionManagedPaths(root string) error {
	for _, relativePath := range []string{
		"environments",
		"execution-ledger",
		filepath.Join("execution-ledger", "transactions"),
		filepath.Join("execution-ledger", "sessions"),
	} {
		path := filepath.Join(root, relativePath)
		info, err := os.Lstat(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("inspect fake execution managed path: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("fake execution managed path must be a real directory: %s", filepath.ToSlash(relativePath))
		}
	}
	return nil
}

func prepareFakeExecutionLifecycle(plan Plan, root string) (BackendLifecycleRecord, error) {
	lifecycle, err := environment.New(root)
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	profile := plan.backendLifecycleProfile()
	record, err := lifecycle.Get(plan.ApplicationID, profile)
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	if record.State == environment.StateBlocked || record.State == environment.StateRetired {
		record, err = lifecycle.Plan(plan.ApplicationID, profile)
	}
	if err == nil && record.State == environment.StateMissing {
		record, err = lifecycle.Plan(plan.ApplicationID, profile)
	}
	if err == nil && (record.State == environment.StatePlanned || record.State == environment.StateRepairRequired) {
		record, err = lifecycle.Stage(plan.ApplicationID, profile)
	}
	if err != nil {
		return BackendLifecycleRecord{}, err
	}
	if record.State == environment.StateStaged || record.State == environment.StateReady {
		for _, gate := range []string{"recipe-trust", "backend-binding"} {
			record, err = lifecycle.SatisfyGate(plan.ApplicationID, profile, gate)
			if err != nil {
				return BackendLifecycleRecord{}, err
			}
		}
	}
	return plan.RecordBackendLifecycleState(root, "inspect", "", "", "")
}

func fakeExecutionTrust(applicationID string, provenance Provenance) (recipe.TrustState, error) {
	trust := recipe.TrustState{
		ID:              applicationID,
		DigestVerified:  provenance.DigestVerified,
		SignatureStatus: recipe.SignatureStatus(provenance.SignatureStatus),
	}
	switch trust.SignatureStatus {
	case recipe.SignatureDevelopmentOnly, recipe.SignatureUnsigned:
		trust.DevelopmentOnly = true
		trust.Reason = "digest verified; test fixture is not production trusted"
	case recipe.SignatureSigned:
		trust.Reason = "digest verified; production signature still requires the production verifier"
	default:
		return recipe.TrustState{}, fmt.Errorf("unsupported fake execution signature status %q", provenance.SignatureStatus)
	}
	return trust, nil
}

func fakeExecutionEvidenceChecks(root string, plan Plan, lifecycle BackendLifecycleRecord, transaction execution.LedgerRecord, session execution.SessionRecord, fanOut ExecutionSessionFanOutEvidence) ([]KDEFakeExecutionCheck, error) {
	paths := []string{lifecycle.RelativePath, transaction.RelativePath, session.RelativePath}
	for _, relativePath := range paths {
		if err := verifyFakeExecutionReceipt(root, relativePath); err != nil {
			return nil, err
		}
	}
	identityMatches := lifecycle.EnvironmentRecord.ApplicationID == plan.ApplicationID && transaction.ApplicationID == plan.ApplicationID && session.ApplicationID == plan.ApplicationID
	safetyClosed := !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted && !session.ExecutionStarted && !session.BackendProcessStarted && !session.HostRootModified
	checks := []KDEFakeExecutionCheck{
		fakeExecutionCheck("recipe-identity", plan.RecipeDigestVerified && identityMatches, "The digest-verified recipe identity matches every persisted record."),
		fakeExecutionCheck("lifecycle-staged", lifecycle.EnvironmentRecord.State == environment.StateStaged && containsString(lifecycle.EnvironmentRecord.PendingGates(), "portal-policy-review") && containsString(lifecycle.EnvironmentRecord.PendingGates(), "snapshot-baseline"), "The controlled lifecycle receipt is staged with Portal and snapshot evidence still pending."),
		fakeExecutionCheck("transaction-readback", transaction.RecordType == "execution-transaction-ledger-record" && len(transaction.SHA256) == 64, "The execution transaction was read back from its relative receipt path."),
		fakeExecutionCheck("session-readback", session.RecordType == "execution-session-status-record" && session.StatusPersisted && len(session.SHA256) == 64, "The blocked session status was read back from its relative receipt path."),
		fakeExecutionCheck("kde-fan-out", fanOut.SafeForKDE && fanOut.SurfaceCount == 4, "One session receipt safely fans out to task manager, KWin, tray, and Compatibility Center."),
		fakeExecutionCheck("unsafe-gates-closed", safetyClosed, "Launch, execution, process start, and host-root mutation remain disabled."),
	}
	for _, check := range checks {
		if check.Status != "pass" {
			return nil, fmt.Errorf("fake execution evidence check failed: %s", check.ID)
		}
	}
	return checks, nil
}

func verifyFakeExecutionReceipt(root string, relativePath string) error {
	if relativePath == "" || filepath.IsAbs(relativePath) || strings.Contains(relativePath, "..") {
		return fmt.Errorf("unsafe fake execution receipt path %q", relativePath)
	}
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("read back fake execution receipt %s: %w", relativePath, err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("fake execution receipt is not a regular file: %s", relativePath)
	}
	return nil
}

func fakeExecutionCheck(id string, passed bool, summary string) KDEFakeExecutionCheck {
	status := "fail"
	if passed {
		status = "pass"
	}
	return KDEFakeExecutionCheck{ID: id, Status: status, Summary: summary}
}

func passedExecutionGateCount(gates []execution.Gate) int {
	count := 0
	for _, gate := range gates {
		if gate.Status == execution.GatePass {
			count++
		}
	}
	return count
}
