package owner

import (
	"errors"
	"fmt"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/record"
	"xnix.local/xnix/internal/runtime/rootfs"
)

const (
	RestrictedOwnerSmokeMode      = "restricted-smoke"
	RestrictedOwnerSmokeDirective = "authorize-restricted-owner-smoke"
)

type RestrictedOwnerSmokeReceipt struct {
	Version                     string                           `json:"version"`
	SchemaVersion               string                           `json:"schema_version"`
	RecordType                  string                           `json:"record_type"`
	ReceiptID                   string                           `json:"receipt_id"`
	Source                      string                           `json:"source"`
	Mode                        string                           `json:"mode"`
	Directive                   string                           `json:"directive"`
	RelativePath                string                           `json:"relative_path"`
	SHA256                      string                           `json:"sha256"`
	ActivationPreflight         RestrictedOwnerSmokePreflight    `json:"activation_preflight"`
	SmokeBatch                  RestrictedOwnerSmokeBatchSummary `json:"smoke_batch"`
	Checks                      []RestrictedOwnerSmokeCheck      `json:"checks"`
	CheckIDs                    []string                         `json:"check_ids"`
	CheckCount                  int                              `json:"check_count"`
	PassedCheckCount            int                              `json:"passed_check_count"`
	AllChecksPassed             bool                             `json:"all_checks_passed"`
	ReceiptPersisted            bool                             `json:"receipt_persisted"`
	ReceiptReadBack             bool                             `json:"receipt_read_back"`
	StateRootWritesEnabled      bool                             `json:"state_root_writes_enabled"`
	StateRootWriteScope         string                           `json:"state_root_write_scope"`
	StateRootPathExposed        bool                             `json:"state_root_path_exposed"`
	RestrictedSmokeReady        bool                             `json:"restricted_smoke_ready"`
	RestrictedSmokeAuthorized   bool                             `json:"restricted_smoke_authorized"`
	ProductionActivationReady   bool                             `json:"production_activation_ready"`
	ProductionOwnerEnabled      bool                             `json:"production_owner_enabled"`
	SystemServiceStarted        bool                             `json:"system_service_started"`
	SessionBusClaimed           bool                             `json:"session_bus_claimed"`
	ProductionBusClaimed        bool                             `json:"production_bus_claimed"`
	WriteMethodsEnabled         bool                             `json:"write_methods_enabled"`
	BackendLaunchEnabled        bool                             `json:"backend_launch_enabled"`
	NetworkRequired             bool                             `json:"network_required"`
	HostRootModified            bool                             `json:"host_root_modified"`
	PrivilegedContainerRequired bool                             `json:"privileged_container_required"`
	BackendDetailsExposed       bool                             `json:"backend_details_exposed"`
	BlockedActions              []string                         `json:"blocked_actions"`
	NextRequirements            []string                         `json:"next_requirements"`
	DesktopSafeSummary          string                           `json:"desktop_safe_summary"`
}

type RestrictedOwnerSmokePreflight struct {
	RequestType               string `json:"request_type"`
	PreflightDecision         string `json:"preflight_decision"`
	RestrictedSmokeReady      bool   `json:"restricted_smoke_ready"`
	ProductionActivationReady bool   `json:"production_activation_ready"`
	PassedCheckCount          int    `json:"passed_check_count"`
	PendingCheckCount         int    `json:"pending_check_count"`
	BlockedCheckCount         int    `json:"blocked_check_count"`
}

type RestrictedOwnerSmokeBatchSummary struct {
	SchemaVersion           string `json:"schema_version"`
	RequestType             string `json:"request_type"`
	BatchType               string `json:"batch_type"`
	Source                  string `json:"source"`
	SHA256                  string `json:"sha256"`
	RecordCount             int    `json:"record_count"`
	ReadDispatchRecordCount int    `json:"read_dispatch_record_count"`
	WriteDenialRecordCount  int    `json:"write_denial_record_count"`
	ReadDispatchMethodCount int    `json:"read_dispatch_method_count"`
	WriteMethodCount        int    `json:"write_method_count"`
	AllReadDispatchReady    bool   `json:"all_read_dispatch_ready"`
	AllWriteDenialsReady    bool   `json:"all_write_denials_ready"`
	EventLoopStarted        bool   `json:"event_loop_started"`
	SessionBusClaimed       bool   `json:"session_bus_claimed"`
	ProductionBusClaimed    bool   `json:"production_bus_claimed"`
	SystemServiceStarted    bool   `json:"system_service_started"`
	WriteMethodsEnabled     bool   `json:"write_methods_enabled"`
	BackendDetailsExposed   bool   `json:"backend_details_exposed"`
	HostRootModified        bool   `json:"host_root_modified"`
}

type RestrictedOwnerSmokeCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func RecordRestrictedOwnerSmokeReceipt(runtimeRoot string, stateRoot string, mode string, directive string) (RestrictedOwnerSmokeReceipt, error) {
	if mode != RestrictedOwnerSmokeMode || directive != RestrictedOwnerSmokeDirective {
		return RestrictedOwnerSmokeReceipt{}, errors.New("restricted owner smoke receipt requires --mode restricted-smoke and --authorize authorize-restricted-owner-smoke")
	}
	if stateRoot == "" {
		return RestrictedOwnerSmokeReceipt{}, errors.New("restricted owner smoke receipt requires an explicit state root")
	}
	preflight, err := appidentity.NewRuntimeServiceActivationPreflightPreview(runtimeRoot)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	if !preflight.RestrictedSmokeReady || preflight.PreflightDecision != "restricted-owner-smoke-ready" {
		return RestrictedOwnerSmokeReceipt{}, fmt.Errorf("restricted owner smoke is not ready: %s", preflight.PreflightDecision)
	}
	batchRecords, err := NewSmokeBatchRecords(runtimeRoot)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	batch, err := summarizeRestrictedOwnerSmokeBatch(batchRecords)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	receipt := newRestrictedOwnerSmokeReceipt(preflight, batch, mode, directive)
	checks := restrictedOwnerSmokeChecks(receipt)
	passed := countRestrictedOwnerSmokePasses(checks)
	receipt.Checks = checks
	receipt.CheckIDs = restrictedOwnerSmokeCheckIDs(checks)
	receipt.CheckCount = len(checks)
	receipt.PassedCheckCount = passed
	receipt.AllChecksPassed = passed == len(checks)
	if !receipt.AllChecksPassed {
		return RestrictedOwnerSmokeReceipt{}, errors.New("restricted owner smoke receipt checks did not all pass")
	}
	receipt.SHA256 = ""
	digest, err := record.Digest(receipt)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	receipt.SHA256 = digest
	if err := validateNoBackendTerms(receipt, "restricted owner smoke receipt"); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	root, err := rootfs.Ensure(stateRoot)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	if _, err := record.Save(root, receipt.RelativePath, receipt); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	var loaded RestrictedOwnerSmokeReceipt
	if err := record.Load(root, receipt.RelativePath, &loaded); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	if err := validateRestrictedOwnerSmokeReceipt(loaded); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	return loaded, nil
}

func LoadRestrictedOwnerSmokeReceipt(stateRoot string) (RestrictedOwnerSmokeReceipt, error) {
	root, err := rootfs.Open(stateRoot)
	if err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	var loaded RestrictedOwnerSmokeReceipt
	if err := record.Load(root, restrictedOwnerSmokeReceiptPath(), &loaded); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	if err := validateRestrictedOwnerSmokeReceipt(loaded); err != nil {
		return RestrictedOwnerSmokeReceipt{}, err
	}
	return loaded, nil
}

func newRestrictedOwnerSmokeReceipt(preflight appidentity.RuntimeServiceActivationPreflightPreview, batch RestrictedOwnerSmokeBatchSummary, mode string, directive string) RestrictedOwnerSmokeReceipt {
	return RestrictedOwnerSmokeReceipt{
		Version:       preflight.Version,
		SchemaVersion: "xnix.runtime.restricted_owner_smoke_receipt.v1",
		RecordType:    "restricted-owner-smoke-execution-receipt",
		ReceiptID:     "restricted-owner-smoke-" + preflight.Version,
		Source:        "runtime-service-activation-preflight+runtime-owner-smoke-batch",
		Mode:          mode,
		Directive:     directive,
		RelativePath:  restrictedOwnerSmokeReceiptPath(),
		ActivationPreflight: RestrictedOwnerSmokePreflight{
			RequestType:               preflight.RequestType,
			PreflightDecision:         preflight.PreflightDecision,
			RestrictedSmokeReady:      preflight.RestrictedSmokeReady,
			ProductionActivationReady: preflight.ProductionActivationReady,
			PassedCheckCount:          preflight.Counts.Passed,
			PendingCheckCount:         preflight.Counts.Pending,
			BlockedCheckCount:         preflight.Counts.Blocked,
		},
		SmokeBatch:                  batch,
		ReceiptPersisted:            true,
		ReceiptReadBack:             true,
		StateRootWritesEnabled:      true,
		StateRootWriteScope:         "explicit-test-root-only",
		StateRootPathExposed:        false,
		RestrictedSmokeReady:        true,
		RestrictedSmokeAuthorized:   true,
		ProductionActivationReady:   false,
		ProductionOwnerEnabled:      false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		BackendLaunchEnabled:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"start production Runtime service from restricted owner smoke receipt",
			"claim production D-Bus name from restricted owner smoke receipt",
			"enable Runtime write methods from restricted owner smoke receipt",
			"launch compatibility backend from restricted owner smoke receipt",
			"expose host paths, state roots, or backend details from restricted owner smoke receipt",
			"mutate host root during restricted owner smoke receipt",
		},
		NextRequirements: []string{
			"Run a separately authorized private session-bus smoke before production ownership can advance.",
			"Prove packaged long-running Runtime owner readiness without accepting the smoke adapter as production.",
			"Keep production service activation behind explicit human authorization and package-managed writes.",
			"Require production recipe trust before enabling install or launch writes.",
		},
		DesktopSafeSummary: "Runtime recorded restricted owner smoke evidence while production service start, D-Bus ownership, Runtime writes, backend launch, and host mutation remain disabled.",
	}
}

func summarizeRestrictedOwnerSmokeBatch(records []SmokeBatchRecord) (RestrictedOwnerSmokeBatchSummary, error) {
	if len(records) == 0 {
		return RestrictedOwnerSmokeBatchSummary{}, errors.New("owner smoke batch is empty")
	}
	if records[0].RequestType != "runtime-owner-smoke-batch-record" {
		return RestrictedOwnerSmokeBatchSummary{}, fmt.Errorf("unexpected owner smoke batch request type: %s", records[0].RequestType)
	}
	digest, err := record.Digest(records)
	if err != nil {
		return RestrictedOwnerSmokeBatchSummary{}, err
	}
	summary := RestrictedOwnerSmokeBatchSummary{
		SchemaVersion:           records[0].SchemaVersion,
		RequestType:             records[0].RequestType,
		BatchType:               records[0].BatchType,
		Source:                  records[0].Source,
		SHA256:                  digest,
		RecordCount:             len(records),
		ReadDispatchMethodCount: records[0].ReadDispatchMethodCount,
		WriteMethodCount:        records[0].WriteMethodCount,
		AllReadDispatchReady:    true,
		AllWriteDenialsReady:    true,
	}
	for _, smokeRecord := range records {
		if smokeRecord.EventLoopStarted || smokeRecord.SessionBusClaimed || smokeRecord.ProductionBusClaimed || smokeRecord.SystemServiceStarted || smokeRecord.NetworkRequired || smokeRecord.HostRootModified || smokeRecord.PrivilegedContainerRequired || smokeRecord.BackendDetailsExposed {
			return RestrictedOwnerSmokeBatchSummary{}, fmt.Errorf("owner smoke batch opened unsafe gate at %s", smokeRecord.Method)
		}
		if smokeRecord.SchemaVersion != summary.SchemaVersion || smokeRecord.RequestType != summary.RequestType || smokeRecord.BatchType != summary.BatchType || smokeRecord.Source != summary.Source {
			return RestrictedOwnerSmokeBatchSummary{}, fmt.Errorf("owner smoke batch metadata mismatch at %s", smokeRecord.Method)
		}
		switch smokeRecord.RecordType {
		case "read-dispatch":
			summary.ReadDispatchRecordCount++
			if !smokeRecord.ReadOnlyDispatch || smokeRecord.WriteMethod || !smokeRecord.RouteReady || !smokeRecord.DispatchReady {
				summary.AllReadDispatchReady = false
			}
		case "write-denial":
			summary.WriteDenialRecordCount++
			if !smokeRecord.WriteMethod || smokeRecord.ReadOnlyDispatch || !smokeRecord.DispatchReady || smokeRecord.ErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
				summary.AllWriteDenialsReady = false
			}
		default:
			return RestrictedOwnerSmokeBatchSummary{}, fmt.Errorf("unknown owner smoke batch record type: %s", smokeRecord.RecordType)
		}
	}
	return summary, nil
}

func restrictedOwnerSmokeChecks(receipt RestrictedOwnerSmokeReceipt) []RestrictedOwnerSmokeCheck {
	return []RestrictedOwnerSmokeCheck{
		restrictedOwnerSmokeCheck("activation-preflight-ready", receipt.ActivationPreflight.RestrictedSmokeReady && receipt.ActivationPreflight.PreflightDecision == "restricted-owner-smoke-ready" && receipt.ActivationPreflight.BlockedCheckCount == 0, "Activation preflight allows only restricted owner smoke evidence."),
		restrictedOwnerSmokeCheck("smoke-batch-read-coverage", receipt.SmokeBatch.AllReadDispatchReady && receipt.SmokeBatch.ReadDispatchRecordCount == receipt.SmokeBatch.ReadDispatchMethodCount, "Owner smoke batch covers every read-only dispatch route."),
		restrictedOwnerSmokeCheck("write-denial-coverage", receipt.SmokeBatch.AllWriteDenialsReady && receipt.SmokeBatch.WriteDenialRecordCount == receipt.SmokeBatch.WriteMethodCount, "Owner smoke batch proves every write method is denied."),
		restrictedOwnerSmokeCheck("restricted-authorization", receipt.RestrictedSmokeAuthorized && receipt.Mode == RestrictedOwnerSmokeMode && receipt.Directive == RestrictedOwnerSmokeDirective, "Receipt creation requires the restricted owner smoke authorization directive."),
		restrictedOwnerSmokeCheck("ownership-boundary", !receipt.SystemServiceStarted && !receipt.SessionBusClaimed && !receipt.ProductionBusClaimed && !receipt.ProductionOwnerEnabled, "Receipt does not start a service or claim session or production D-Bus ownership."),
		restrictedOwnerSmokeCheck("unsafe-gates-closed", !receipt.WriteMethodsEnabled && !receipt.BackendLaunchEnabled && !receipt.NetworkRequired && !receipt.HostRootModified && !receipt.PrivilegedContainerRequired && !receipt.BackendDetailsExposed && !receipt.StateRootPathExposed, "Receipt keeps write methods, backend launch, network, privilege, host mutation, state-root paths, and backend details disabled."),
		restrictedOwnerSmokeCheck("receipt-boundary", receipt.ReceiptPersisted && receipt.ReceiptReadBack && receipt.StateRootWritesEnabled && receipt.StateRootWriteScope == "explicit-test-root-only" && !filepath.IsAbs(receipt.RelativePath), "Receipt is persisted only under the explicit test state root and read back by relative path."),
	}
}

func restrictedOwnerSmokeCheck(id string, passed bool, summary string) RestrictedOwnerSmokeCheck {
	status := "blocked"
	if passed {
		status = "pass"
	}
	return RestrictedOwnerSmokeCheck{ID: id, Status: status, Summary: summary}
}

func restrictedOwnerSmokeCheckIDs(checks []RestrictedOwnerSmokeCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRestrictedOwnerSmokePasses(checks []RestrictedOwnerSmokeCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}

func validateRestrictedOwnerSmokeReceipt(receipt RestrictedOwnerSmokeReceipt) error {
	if receipt.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt.v1" ||
		receipt.RecordType != "restricted-owner-smoke-execution-receipt" ||
		receipt.RelativePath != restrictedOwnerSmokeReceiptPath() ||
		receipt.SHA256 == "" ||
		receipt.CheckCount != len(receipt.Checks) ||
		receipt.PassedCheckCount != receipt.CheckCount ||
		!receipt.AllChecksPassed ||
		!receipt.ReceiptPersisted ||
		!receipt.ReceiptReadBack ||
		!receipt.RestrictedSmokeReady ||
		!receipt.RestrictedSmokeAuthorized ||
		receipt.ProductionActivationReady ||
		receipt.ProductionOwnerEnabled ||
		receipt.SystemServiceStarted ||
		receipt.SessionBusClaimed ||
		receipt.ProductionBusClaimed ||
		receipt.WriteMethodsEnabled ||
		receipt.BackendLaunchEnabled ||
		receipt.HostRootModified {
		return errors.New("restricted owner smoke receipt failed validation")
	}
	clone := receipt
	clone.SHA256 = ""
	digest, err := record.Digest(clone)
	if err != nil {
		return err
	}
	if digest != receipt.SHA256 {
		return errors.New("restricted owner smoke receipt digest mismatch")
	}
	return validateNoBackendTerms(receipt, "restricted owner smoke receipt")
}

func restrictedOwnerSmokeReceiptPath() string {
	return "owner-smoke/restricted-owner-smoke-receipt.json"
}
