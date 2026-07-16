package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/portal"
)

type KDEFakePortalEvidenceRecord struct {
	SchemaVersion               string                         `json:"schema_version"`
	RecordType                  string                         `json:"record_type"`
	Source                      string                         `json:"source"`
	Mode                        string                         `json:"mode"`
	Application                 KDEFakeExecutionApplication    `json:"application"`
	BeforePortal                KDEFakePortalBefore            `json:"before_portal"`
	Portal                      KDEFakePortalReceipt           `json:"portal"`
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
	PortalEvidenceRecorded      bool                           `json:"portal_evidence_recorded"`
	PortalGateChangedOnly       bool                           `json:"portal_gate_changed_only"`
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
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type KDEFakePortalBefore struct {
	LifecycleState     string `json:"lifecycle_state"`
	PortalGateStatus   string `json:"portal_gate_status"`
	BlockedReasonCount int    `json:"blocked_reason_count"`
}

type KDEFakePortalReceipt struct {
	HandleToken           string `json:"handle_token"`
	Operation             string `json:"operation"`
	ReceiptRelativePath   string `json:"receipt_relative_path"`
	ReceiptSHA256         string `json:"receipt_sha256"`
	CreatedState          string `json:"created_state"`
	ResolvedState         string `json:"resolved_state"`
	CompletedState        string `json:"completed_state"`
	PermissionState       string `json:"permission_state"`
	ReceiptPersisted      bool   `json:"receipt_persisted"`
	ReceiptReused         bool   `json:"receipt_reused"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	PermissionEvidence    bool   `json:"permission_evidence"`
	RealPortalCallEnabled bool   `json:"real_portal_call_enabled"`
	HostPermissionChanged bool   `json:"host_permission_changed"`
	ExecutionApproved     bool   `json:"execution_approved"`
}

func NewKDEFakePortalEvidenceRecord(recipeRecord Recipe, provenance Provenance, options KDEFakeExecutionEvidenceOptions) (KDEFakePortalEvidenceRecord, error) {
	if options.Mode != kdeFakeExecutionTestMode {
		return KDEFakePortalEvidenceRecord{}, errors.New("fake Portal evidence requires test-only mode")
	}
	if err := validateKDEFakeExecutionStateRoot(options.StateRoot); err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	if err := validateKDEFakeExecutionManagedPaths(options.StateRoot); err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	if err := validateKDEFakePortalManagedPath(options.StateRoot); err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	if !provenance.DigestVerified {
		return KDEFakePortalEvidenceRecord{}, errors.New("fake Portal evidence requires a digest-verified registry recipe")
	}
	plan, err := NewPlanWithProvenance(recipeRecord, provenance)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	initialLifecycle, err := prepareFakeExecutionLifecycle(plan, options.StateRoot)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	trust, err := fakeExecutionTrust(recipeRecord.ID, provenance)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	baselineTransaction, err := reviewedFakeExecutionTransaction(plan.ApplicationID, execution.Inputs{
		Trust:                   trust,
		Environment:             initialLifecycle.EnvironmentRecord,
		SnapshotBaselinePresent: false,
		PortalRequiredOps:       []string{"file-open"},
		PortalGrantedOps:        []string{},
	})
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}

	portalRecord, reused, err := ensureCompletedFakePortalReceipt(options.StateRoot, plan.ApplicationID)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	if err := validateCompletedFakePortalRecord(portalRecord, plan.ApplicationID); err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	portalDigest, err := fakePortalReceiptDigest(options.StateRoot, portalRecord.RelativePath)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}

	lifecycleStore, err := environment.New(options.StateRoot)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	profile := plan.backendLifecycleProfile()
	if _, err := lifecycleStore.SatisfyGate(plan.ApplicationID, profile, "portal-policy-review"); err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	lifecycleRecord, err := plan.RecordBackendLifecycleState(options.StateRoot, "inspect", "", "", "")
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}

	portalReceipt := fakePortalExecutionReceipt(portalRecord)
	inputs := execution.Inputs{
		Trust:                    trust,
		Environment:              lifecycleRecord.EnvironmentRecord,
		SnapshotBaselinePresent:  false,
		PortalRequiredOps:        []string{"file-open"},
		PortalGrantedOps:         []string{},
		PortalPermissionReceipts: []execution.PortalPermissionReceipt{portalReceipt},
	}
	tx, err := reviewedFakeExecutionTransaction(plan.ApplicationID, inputs)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	transactionRecord, sessionRecord, err := persistFakeExecutionTransaction(options.StateRoot, tx)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}
	fanOut, err := plan.ExecutionSessionFanOutEvidence(options.StateRoot, tx.RequestID)
	if err != nil {
		return KDEFakePortalEvidenceRecord{}, err
	}

	portalGateChangedOnly := fakePortalGateDeltaOnly(baselineTransaction.Gates, transactionRecord.Transaction.Gates)
	checks := fakePortalEvidenceChecks(baselineTransaction, portalRecord, lifecycleRecord, transactionRecord, sessionRecord, fanOut, portalGateChangedOnly)
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	if passed != len(checks) {
		return KDEFakePortalEvidenceRecord{}, errors.New("fake Portal evidence checks did not all pass")
	}

	return KDEFakePortalEvidenceRecord{
		SchemaVersion: "xnix.runtime.kde_fake_portal_evidence.v1",
		RecordType:    "kde-fake-portal-evidence-record",
		Source:        "kde-fake-execution-evidence+state-root-portal-request+execution-ledger+execution-session-record",
		Mode:          options.Mode,
		Application: KDEFakeExecutionApplication{
			ID:             plan.ApplicationID,
			Name:           plan.DisplayName,
			Icon:           plan.Icon,
			DesktopFile:    plan.DesktopFile,
			IdentityDigest: plan.StableIdentityDigest,
		},
		BeforePortal: KDEFakePortalBefore{
			LifecycleState:     string(initialLifecycle.EnvironmentRecord.State),
			PortalGateStatus:   fakeExecutionGateStatus(baselineTransaction.Gates, "portal-permission"),
			BlockedReasonCount: len(baselineTransaction.BlockedReasons),
		},
		Portal: KDEFakePortalReceipt{
			HandleToken:           portalRecord.Request.HandleToken,
			Operation:             portalRecord.Request.Operation,
			ReceiptRelativePath:   portalRecord.RelativePath,
			ReceiptSHA256:         portalDigest,
			CreatedState:          string(portal.StatePendingUserMediation),
			ResolvedState:         string(portal.StateGranted),
			CompletedState:        string(portalRecord.Request.State),
			PermissionState:       string(portalRecord.Request.PermissionState),
			ReceiptPersisted:      true,
			ReceiptReused:         reused,
			RequestObjectCreated:  true,
			PermissionEvidence:    true,
			RealPortalCallEnabled: false,
			HostPermissionChanged: false,
			ExecutionApproved:     false,
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
		StateRootRecordCount:        4,
		StateRootPathExposed:        false,
		PortalEvidenceRecorded:      true,
		PortalGateChangedOnly:       portalGateChangedOnly,
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
		DesktopSafeSummary:          "Runtime recorded and consumed one fake Portal permission receipt for the KDE application while lifecycle readiness, snapshot, trust, and launch gates remain closed.",
	}, nil
}

func validateKDEFakePortalManagedPath(root string) error {
	path := filepath.Join(root, "portal-requests")
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("inspect fake Portal managed path: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return errors.New("fake Portal managed path must be a real directory: portal-requests")
	}
	return nil
}

func ensureCompletedFakePortalReceipt(root string, applicationID string) (portal.Record, bool, error) {
	requests, malformed, err := portal.ReadRequests(root)
	if err != nil {
		return portal.Record{}, false, err
	}
	if len(malformed) > 0 {
		return portal.Record{}, false, fmt.Errorf("fake Portal evidence found malformed receipts: %v", malformed)
	}
	ledger, err := portal.NewLedger(root)
	if err != nil {
		return portal.Record{}, false, err
	}
	for index := len(requests) - 1; index >= 0; index-- {
		request := requests[index]
		if request.ApplicationID != applicationID || request.Operation != "file-open" {
			continue
		}
		switch request.State {
		case portal.StateCompleted:
			record, inspectErr := ledger.Inspect(request.HandleToken)
			return record, true, inspectErr
		case portal.StateGranted:
			record, completeErr := ledger.Complete(request.HandleToken)
			return record, true, completeErr
		case portal.StatePendingUserMediation, portal.StateFailed:
			if _, resolveErr := ledger.Resolve(request.HandleToken, portal.OutcomeGranted); resolveErr != nil {
				return portal.Record{}, false, resolveErr
			}
			record, completeErr := ledger.Complete(request.HandleToken)
			return record, true, completeErr
		}
	}
	created, err := ledger.Create(portal.RequestSpec{ApplicationID: applicationID, Operation: "file-open", Reason: "Open a selected document in the compatibility application."})
	if err != nil {
		return portal.Record{}, false, err
	}
	if _, err := ledger.Resolve(created.Request.HandleToken, portal.OutcomeGranted); err != nil {
		return portal.Record{}, false, err
	}
	completed, err := ledger.Complete(created.Request.HandleToken)
	return completed, false, err
}

func validateCompletedFakePortalRecord(record portal.Record, applicationID string) error {
	if record.SchemaVersion != "xnix.runtime.portal_request_record.v1" || record.RecordType != "portal-permission-request-record" || record.Source != "go-runtime-state-root-portal-broker" {
		return errors.New("fake Portal receipt has unsupported schema")
	}
	if record.Request.ApplicationID != applicationID || record.Request.Operation != "file-open" || record.Request.State != portal.StateCompleted || record.Request.PermissionState != portal.PermissionGranted {
		return errors.New("fake Portal receipt is not a completed file-open permission record")
	}
	if filepath.IsAbs(record.RelativePath) || record.StateRootPathExposed || record.RealPortalCallEnabled || record.ExecutionApproved || record.HostPermissionChanged || record.HostRootModified || record.BackendDetailsExposed || record.NetworkRequired || record.PrivilegedContainerRequired || record.Request.DirectAccessAllowed || record.Request.HostPermissionChanged || record.Request.BackendDetailsExposed {
		return errors.New("fake Portal receipt has unsafe enabled gates")
	}
	return nil
}

func fakePortalReceiptDigest(root string, relativePath string) (string, error) {
	if err := verifyFakeExecutionReceipt(root, relativePath); err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(relativePath)))
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func fakePortalGateDeltaOnly(before []execution.Gate, after []execution.Gate) bool {
	if len(before) != len(after) {
		return false
	}
	beforeStates := make(map[string]execution.GateStatus, len(before))
	for _, gate := range before {
		beforeStates[gate.ID] = gate.Status
	}
	changed := []string{}
	for _, gate := range after {
		if beforeStates[gate.ID] != gate.Status {
			changed = append(changed, gate.ID)
		}
	}
	sort.Strings(changed)
	return len(changed) == 1 && changed[0] == "portal-permission" && fakeExecutionGateStatus(before, "portal-permission") == "pending" && fakeExecutionGateStatus(after, "portal-permission") == "pass"
}

func reviewedFakeExecutionTransaction(applicationID string, inputs execution.Inputs) (execution.Transaction, error) {
	pipeline := execution.NewPipeline()
	transaction, err := pipeline.Create(applicationID, inputs)
	if err != nil {
		return execution.Transaction{}, err
	}
	transaction, err = transaction.Review(execution.DecisionApproved)
	if err != nil {
		return execution.Transaction{}, err
	}
	return transaction.Preflight(inputs), nil
}

func fakePortalExecutionReceipt(record portal.Record) execution.PortalPermissionReceipt {
	return execution.PortalPermissionReceipt{
		HandleToken:           record.Request.HandleToken,
		Operation:             record.Request.Operation,
		RelativePath:          record.RelativePath,
		RequestState:          string(record.Request.State),
		PermissionState:       string(record.Request.PermissionState),
		PermissionGranted:     record.PermissionGranted,
		ExecutionApproved:     false,
		RealPortalCallEnabled: false,
		StateRootPathExposed:  false,
		HostPermissionChanged: false,
	}
}

func persistFakeExecutionTransaction(root string, transaction execution.Transaction) (execution.LedgerRecord, execution.SessionRecord, error) {
	ledger, err := execution.NewLedger(root)
	if err != nil {
		return execution.LedgerRecord{}, execution.SessionRecord{}, err
	}
	if _, err := ledger.Record(transaction); err != nil {
		return execution.LedgerRecord{}, execution.SessionRecord{}, err
	}
	if _, err := ledger.RecordSession(transaction.RequestID); err != nil {
		return execution.LedgerRecord{}, execution.SessionRecord{}, err
	}
	transactionRecord, err := ledger.Load(transaction.RequestID)
	if err != nil {
		return execution.LedgerRecord{}, execution.SessionRecord{}, err
	}
	sessionRecord, err := ledger.LoadSession(transaction.RequestID)
	if err != nil {
		return execution.LedgerRecord{}, execution.SessionRecord{}, err
	}
	return transactionRecord, sessionRecord, nil
}

func fakeExecutionGateStatus(gates []execution.Gate, gateID string) string {
	for _, gate := range gates {
		if gate.ID == gateID {
			return string(gate.Status)
		}
	}
	return "missing"
}

func fakePortalEvidenceChecks(baseline execution.Transaction, portalRecord portal.Record, lifecycle BackendLifecycleRecord, transaction execution.LedgerRecord, session execution.SessionRecord, fanOut ExecutionSessionFanOutEvidence, portalGateChangedOnly bool) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("base-evidence", baseline.State == execution.StateBlocked && fakeExecutionGateStatus(baseline.Gates, "portal-permission") == "pending", "The controlled no-receipt execution baseline is blocked before Portal evidence is added."),
		fakeExecutionCheck("portal-receipt-readback", portalRecord.Request.State == portal.StateCompleted && portalRecord.Request.PermissionState == portal.PermissionGranted, "The completed fake Portal receipt was read back from the controlled state root."),
		fakeExecutionCheck("portal-gate-delta", portalGateChangedOnly, "Only the execution Portal gate changed from pending to pass."),
		fakeExecutionCheck("lifecycle-portal-join", lifecycle.EnvironmentRecord.State == environment.StateStaged && containsString(lifecycle.EnvironmentRecord.SatisfiedGates, "portal-policy-review") && containsString(lifecycle.EnvironmentRecord.PendingGates(), "snapshot-baseline"), "The staged lifecycle consumed Portal review evidence while snapshot evidence remains pending."),
		fakeExecutionCheck("execution-receipt-join", transaction.PortalPermissionReceiptCount == 1 && transaction.PortalPermissionReceiptConsumed && fakeExecutionGateStatus(transaction.Transaction.Gates, "portal-permission") == "pass", "The blocked execution receipt consumed exactly one sanitized Portal receipt."),
		fakeExecutionCheck("session-receipt-join", len(session.PortalPermissionReceiptPaths) == 1 && len(session.PortalPermissionReceiptStates) == 1 && fanOut.SafeForKDE, "The blocked session and KDE fan-out retain the Portal receipt evidence."),
		fakeExecutionCheck("unsafe-gates-closed", !portalRecord.RealPortalCallEnabled && !portalRecord.ExecutionApproved && !portalRecord.HostPermissionChanged && !transaction.LaunchAllowed && !transaction.LaunchEnabled && !transaction.BackendStarted && !session.ExecutionStarted && !session.BackendProcessStarted && !session.HostRootModified, "Real Portal transport, execution approval, launch, process start, and host mutation remain disabled."),
	}
}

func fakePortalExecutionSummary(record execution.LedgerRecord) KDEFakeExecutionTransaction {
	return KDEFakeExecutionTransaction{
		RequestID:           record.RequestID,
		ReceiptRelativePath: record.RelativePath,
		ReceiptSHA256:       record.SHA256,
		State:               string(record.Transaction.State),
		ReviewDecision:      string(record.Transaction.ReviewDecision),
		Gates:               append([]execution.Gate{}, record.Transaction.Gates...),
		GateCount:           len(record.Transaction.Gates),
		PassedGateCount:     passedExecutionGateCount(record.Transaction.Gates),
		BlockedReasons:      append([]string{}, record.Transaction.BlockedReasons...),
		ReceiptPersisted:    true,
		LaunchAllowed:       false,
		LaunchEnabled:       false,
		ExecutionStarted:    false,
		ProcessStarted:      false,
	}
}

func fakePortalSessionSummary(record execution.SessionRecord) KDEFakeExecutionSession {
	return KDEFakeExecutionSession{
		ReceiptRelativePath:      record.RelativePath,
		ReceiptSHA256:            record.SHA256,
		State:                    record.SessionState,
		TaskManagerState:         record.TaskManagerState,
		TrayState:                record.TrayState,
		KWinState:                record.KWinState,
		CompatibilityCenterState: record.CompatibilityCenterState,
		BlockedReasons:           append([]string{}, record.BlockedReasons...),
		ReceiptPersisted:         true,
		LiveStateObserved:        false,
		SessionActive:            false,
	}
}
