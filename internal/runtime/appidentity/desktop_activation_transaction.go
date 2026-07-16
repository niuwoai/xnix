package appidentity

import "errors"

type DesktopActivationTransactionPreview struct {
	SchemaVersion               string                                      `json:"schema_version"`
	RequestType                 string                                      `json:"request_type"`
	TransactionType             string                                      `json:"transaction_type"`
	TransactionState            string                                      `json:"transaction_state"`
	Source                      string                                      `json:"source"`
	Desktop                     string                                      `json:"desktop"`
	ReadMethod                  string                                      `json:"read_method"`
	WriteMethod                 string                                      `json:"write_method"`
	ApplicationID               string                                      `json:"application_id"`
	DisplayName                 string                                      `json:"display_name"`
	Icon                        string                                      `json:"icon"`
	DesktopFile                 string                                      `json:"desktop_file"`
	InstallMode                 string                                      `json:"install_mode"`
	PreflightDecision           string                                      `json:"preflight_decision"`
	Staging                     DesktopActivationTransactionStaging         `json:"staging"`
	WriteGate                   DesktopActivationTransactionWriteGate       `json:"write_gate"`
	ReceiptEvidence             DesktopActivationTransactionReceiptEvidence `json:"receipt_evidence"`
	TransactionSteps            []DesktopActivationTransactionStep          `json:"transaction_steps"`
	TransactionStepIDs          []string                                    `json:"transaction_step_ids"`
	TransactionStepCount        int                                         `json:"transaction_step_count"`
	ReadyStepCount              int                                         `json:"ready_step_count"`
	BlockedStepCount            int                                         `json:"blocked_step_count"`
	RollbackSteps               []DesktopActivationRollbackStep             `json:"rollback_steps"`
	RollbackStepIDs             []string                                    `json:"rollback_step_ids"`
	RollbackStepCount           int                                         `json:"rollback_step_count"`
	ActivationCommandPreview    []string                                    `json:"activation_command_preview"`
	RollbackCommandPreview      []string                                    `json:"rollback_command_preview"`
	RuntimeOwned                bool                                        `json:"runtime_owned"`
	GoRuntimeBacked             bool                                        `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                        `json:"kde_policy_owner"`
	UserVisible                 bool                                        `json:"user_visible"`
	InstallerMayProceed         bool                                        `json:"installer_may_proceed"`
	TransactionPlanCreated      bool                                        `json:"transaction_plan_created"`
	TransactionReady            bool                                        `json:"transaction_ready"`
	TransactionCommitted        bool                                        `json:"transaction_committed"`
	StagingPlanReady            bool                                        `json:"staging_plan_ready"`
	StagedFileDigestsRequired   bool                                        `json:"staged_file_digests_required"`
	StagedFileDigestsVerified   bool                                        `json:"staged_file_digests_verified"`
	StagingRootRequired         bool                                        `json:"staging_root_required"`
	StagingRootPathExposed      bool                                        `json:"staging_root_path_exposed"`
	TargetRootPathExposed       bool                                        `json:"target_root_path_exposed"`
	HostRootAllowed             bool                                        `json:"host_root_allowed"`
	FileWritesPerformed         bool                                        `json:"file_writes_performed"`
	DesktopFilesWritten         bool                                        `json:"desktop_files_written"`
	MIMEAppsWritten             bool                                        `json:"mimeapps_written"`
	ManifestWritten             bool                                        `json:"manifest_written"`
	ReceiptWritten              bool                                        `json:"receipt_written"`
	RollbackReceiptRequired     bool                                        `json:"rollback_receipt_required"`
	RollbackReceiptPlanned      bool                                        `json:"rollback_receipt_planned"`
	RollbackReceiptWritten      bool                                        `json:"rollback_receipt_written"`
	RollbackAvailable           bool                                        `json:"rollback_available"`
	KDEServiceCacheRefreshed    bool                                        `json:"kde_service_cache_refreshed"`
	SettingsPersisted           bool                                        `json:"settings_persisted"`
	NotificationsSent           bool                                        `json:"notifications_sent"`
	TaskManagerEntryActive      bool                                        `json:"task_manager_entry_active"`
	KWinRuleApplied             bool                                        `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled       bool                                        `json:"live_tray_bridge_enabled"`
	LaunchEnabled               bool                                        `json:"launch_enabled"`
	BackendLaunchEnabled        bool                                        `json:"backend_launch_enabled"`
	ExecutionStarted            bool                                        `json:"execution_started"`
	HostRootModified            bool                                        `json:"host_root_modified"`
	NetworkRequired             bool                                        `json:"network_required"`
	PrivilegedContainerRequired bool                                        `json:"privileged_container_required"`
	BackendDetailsExposed       bool                                        `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                                        `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed bool                                        `json:"compatibility_storage_exposed"`
	BlockedActions              []string                                    `json:"blocked_actions"`
	DesktopSafeSummary          string                                      `json:"desktop_safe_summary"`
}

type DesktopActivationTransactionStaging struct {
	RequestType              string   `json:"request_type"`
	StagingState             string   `json:"staging_state"`
	PlannedFileIDs           []string `json:"planned_file_ids"`
	PlannedFileCount         int      `json:"planned_file_count"`
	ReceiptFileID            string   `json:"receipt_file_id"`
	ActivatedEntryPointCount int      `json:"activated_entry_point_count"`
	InstallerMayProceed      bool     `json:"installer_may_proceed"`
	StagingPlanReady         bool     `json:"staging_plan_ready"`
	HostRootAllowed          bool     `json:"host_root_allowed"`
}

type DesktopActivationTransactionWriteGate struct {
	MethodName           string `json:"method_name"`
	GateDecision         string `json:"gate_decision"`
	DenialErrorName      string `json:"denial_error_name"`
	WriteMethodEnabled   bool   `json:"write_method_enabled"`
	DispatchEnabled      bool   `json:"dispatch_enabled"`
	RequestObjectCreated bool   `json:"request_object_created"`
}

type DesktopActivationTransactionReceiptEvidence struct {
	EvidenceType                string   `json:"evidence_type"`
	EvidenceState               string   `json:"evidence_state"`
	ReceiptSchemaVersion        string   `json:"receipt_schema_version"`
	CommitReceiptType           string   `json:"commit_receipt_type"`
	RollbackReceiptType         string   `json:"rollback_receipt_type"`
	ReceiptRelativePath         string   `json:"receipt_relative_path"`
	RequiredFileIDs             []string `json:"required_file_ids"`
	RequiredFileCount           int      `json:"required_file_count"`
	InstalledFileDigestRequired bool     `json:"installed_file_digest_required"`
	RollbackDigestRequired      bool     `json:"rollback_digest_required"`
	CommitReceiptRequired       bool     `json:"commit_receipt_required"`
	CommitReceiptPlanned        bool     `json:"commit_receipt_planned"`
	CommitReceiptWritten        bool     `json:"commit_receipt_written"`
	RollbackReceiptRequired     bool     `json:"rollback_receipt_required"`
	RollbackReceiptPlanned      bool     `json:"rollback_receipt_planned"`
	RollbackReceiptWritten      bool     `json:"rollback_receipt_written"`
	DigestGateReady             bool     `json:"digest_gate_ready"`
	CommitAvailable             bool     `json:"commit_available"`
	RollbackAvailable           bool     `json:"rollback_available"`
	RuntimeOwned                bool     `json:"runtime_owned"`
	KDEPolicyOwner              bool     `json:"kde_policy_owner"`
	RootPathExposed             bool     `json:"root_path_exposed"`
	TargetRootPathExposed       bool     `json:"target_root_path_exposed"`
	HostRootModified            bool     `json:"host_root_modified"`
	FileWritesPerformed         bool     `json:"file_writes_performed"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	Summary                     string   `json:"summary"`
}

type DesktopActivationTransactionStep struct {
	ID                     string `json:"id"`
	Status                 string `json:"status"`
	Required               bool   `json:"required"`
	RequiresWriteGate      bool   `json:"requires_write_gate"`
	RequiresDigestMatch    bool   `json:"requires_digest_match"`
	RequiresRollbackRecord bool   `json:"requires_rollback_record"`
	Summary                string `json:"summary"`
}

type DesktopActivationRollbackStep struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	Required            bool   `json:"required"`
	RequiresReceipt     bool   `json:"requires_receipt"`
	RequiresDigestMatch bool   `json:"requires_digest_match"`
	Summary             string `json:"summary"`
}

func (plan Plan) DesktopActivationTransactionPreview(mode string) (DesktopActivationTransactionPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationTransactionPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationTransactionPreview{}, errors.New("desktop activation transaction requires single-line identity fields")
		}
	}
	staging, err := plan.DesktopActivationStagingPreview(mode)
	if err != nil {
		return DesktopActivationTransactionPreview{}, err
	}

	ready := staging.StagingPlanReady && staging.InstallerMayProceed
	steps := desktopActivationTransactionSteps(ready)
	rollbackSteps := desktopActivationRollbackSteps()
	receiptEvidence := desktopActivationTransactionReceiptEvidence(plan.ApplicationID, staging.PlannedFileIDs, staging.RollbackReceiptPlanned)
	readyStepCount, blockedStepCount := countDesktopActivationTransactionSteps(steps)
	preview := DesktopActivationTransactionPreview{
		SchemaVersion:     "xnix.runtime.desktop_activation_transaction.v1",
		RequestType:       "desktop-activation-transaction-preview",
		TransactionType:   "kde-desktop-activation-transaction",
		TransactionState:  desktopActivationTransactionState(ready),
		Source:            "desktop-activation-staging-preview",
		Desktop:           "KDE Plasma",
		ReadMethod:        "GetDesktopActivationTransactionPreview",
		WriteMethod:       "ActivateDesktopIntegration",
		ApplicationID:     plan.ApplicationID,
		DisplayName:       plan.DisplayName,
		Icon:              plan.Icon,
		DesktopFile:       plan.DesktopFile,
		InstallMode:       staging.InstallMode,
		PreflightDecision: staging.PreflightDecision,
		Staging: DesktopActivationTransactionStaging{
			RequestType:              staging.RequestType,
			StagingState:             staging.StagingState,
			PlannedFileIDs:           staging.PlannedFileIDs,
			PlannedFileCount:         staging.PlannedFileCount,
			ReceiptFileID:            staging.ReceiptFileID,
			ActivatedEntryPointCount: staging.ActivatedEntryPointCount,
			InstallerMayProceed:      staging.InstallerMayProceed,
			StagingPlanReady:         staging.StagingPlanReady,
			HostRootAllowed:          staging.HostRootAllowed,
		},
		WriteGate: DesktopActivationTransactionWriteGate{
			MethodName:           "ActivateDesktopIntegration",
			GateDecision:         "disabled-for-preview",
			DenialErrorName:      "org.xnix.Compatibility1.Error.WriteMethodDisabled",
			WriteMethodEnabled:   false,
			DispatchEnabled:      false,
			RequestObjectCreated: false,
		},
		ReceiptEvidence:             receiptEvidence,
		TransactionSteps:            steps,
		TransactionStepIDs:          desktopActivationTransactionStepIDs(steps),
		TransactionStepCount:        len(steps),
		ReadyStepCount:              readyStepCount,
		BlockedStepCount:            blockedStepCount,
		RollbackSteps:               rollbackSteps,
		RollbackStepIDs:             desktopActivationRollbackStepIDs(rollbackSteps),
		RollbackStepCount:           len(rollbackSteps),
		ActivationCommandPreview:    []string{"xnix-activate-desktop-integration", "--plan-source", "runtime-go", "--target", "staging-root"},
		RollbackCommandPreview:      []string{"xnix-rollback-desktop-integration", "--receipt-source", "runtime-go"},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		InstallerMayProceed:         staging.InstallerMayProceed,
		TransactionPlanCreated:      true,
		TransactionReady:            ready,
		TransactionCommitted:        false,
		StagingPlanReady:            staging.StagingPlanReady,
		StagedFileDigestsRequired:   true,
		StagedFileDigestsVerified:   false,
		StagingRootRequired:         true,
		StagingRootPathExposed:      false,
		TargetRootPathExposed:       false,
		HostRootAllowed:             false,
		FileWritesPerformed:         false,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		ManifestWritten:             false,
		ReceiptWritten:              false,
		RollbackReceiptRequired:     true,
		RollbackReceiptPlanned:      staging.RollbackReceiptPlanned,
		RollbackReceiptWritten:      false,
		RollbackAvailable:           false,
		KDEServiceCacheRefreshed:    false,
		SettingsPersisted:           false,
		NotificationsSent:           false,
		TaskManagerEntryActive:      false,
		KWinRuleApplied:             false,
		LiveTrayBridgeEnabled:       false,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		CompatibilityStorageExposed: false,
		BlockedActions: []string{
			"commit desktop activation from transaction preview",
			"write activation files before digest verification",
			"write activation files before rollback receipt planning",
			"write into the host root from transaction preview",
			"refresh KDE service cache from transaction preview",
			"start compatibility backend from desktop activation transaction",
			"expose raw backend command to desktop shell",
		},
		DesktopSafeSummary: desktopActivationTransactionSummary(ready),
	}
	if err := validateNoBackendTerms(preview, "desktop activation transaction preview"); err != nil {
		return DesktopActivationTransactionPreview{}, err
	}
	return preview, nil
}

func desktopActivationTransactionSteps(ready bool) []DesktopActivationTransactionStep {
	status := "ready"
	if !ready {
		status = "blocked"
	}
	return []DesktopActivationTransactionStep{
		desktopActivationTransactionStep("validate-preflight", status, true, false, false, false, "Confirm Runtime preflight allows activation planning."),
		desktopActivationTransactionStep("prepare-staging-root", status, true, true, false, false, "Prepare the installer target without exposing target paths to KDE."),
		desktopActivationTransactionStep("verify-staged-file-digests", status, true, true, true, false, "Verify every staged desktop artifact digest before installation."),
		desktopActivationTransactionStep("install-desktop-entry", status, true, true, true, true, "Install the standard launcher entry for KDE menus and task manager."),
		desktopActivationTransactionStep("install-dolphin-service-menu", status, true, true, true, true, "Install the Dolphin action that routes files through Runtime review."),
		desktopActivationTransactionStep("merge-mimeapps-associations", status, true, true, true, true, "Merge file associations so supported files open through the managed launcher."),
		desktopActivationTransactionStep("install-desktop-integration-manifest", status, true, true, true, true, "Record Runtime-owned KDE integration metadata for audit."),
		desktopActivationTransactionStep("write-rollback-receipt", status, true, true, true, true, "Write the receipt needed to reverse desktop activation safely."),
		desktopActivationTransactionStep("refresh-kde-service-cache", status, false, true, false, true, "Refresh KDE desktop service metadata after files are installed."),
	}
}

func desktopActivationTransactionStep(id string, status string, required bool, requiresWriteGate bool, requiresDigestMatch bool, requiresRollbackRecord bool, summary string) DesktopActivationTransactionStep {
	return DesktopActivationTransactionStep{
		ID:                     id,
		Status:                 status,
		Required:               required,
		RequiresWriteGate:      requiresWriteGate,
		RequiresDigestMatch:    requiresDigestMatch,
		RequiresRollbackRecord: requiresRollbackRecord,
		Summary:                summary,
	}
}

func desktopActivationRollbackSteps() []DesktopActivationRollbackStep {
	return []DesktopActivationRollbackStep{
		desktopActivationRollbackStep("load-activation-receipt", true, true, false, "Load the Runtime activation receipt selected by application identity."),
		desktopActivationRollbackStep("verify-installed-file-digests", true, true, true, "Verify installed artifacts still match the receipt before removal."),
		desktopActivationRollbackStep("remove-desktop-entry", true, true, true, "Remove the generated launcher entry."),
		desktopActivationRollbackStep("remove-dolphin-service-menu", true, true, true, "Remove the generated Dolphin service menu."),
		desktopActivationRollbackStep("restore-mimeapps-associations", true, true, true, "Restore file associations captured by the receipt."),
		desktopActivationRollbackStep("remove-desktop-integration-manifest", true, true, true, "Remove Runtime-owned desktop integration metadata."),
		desktopActivationRollbackStep("mark-receipt-rolled-back", true, true, false, "Mark the receipt as rolled back for later audit."),
	}
}

func desktopActivationRollbackStep(id string, required bool, requiresReceipt bool, requiresDigestMatch bool, summary string) DesktopActivationRollbackStep {
	return DesktopActivationRollbackStep{
		ID:                  id,
		Status:              "planned",
		Required:            required,
		RequiresReceipt:     requiresReceipt,
		RequiresDigestMatch: requiresDigestMatch,
		Summary:             summary,
	}
}

func desktopActivationTransactionReceiptEvidence(applicationID string, plannedFileIDs []string, rollbackReceiptPlanned bool) DesktopActivationTransactionReceiptEvidence {
	requiredFileIDs := append([]string(nil), plannedFileIDs...)
	return DesktopActivationTransactionReceiptEvidence{
		EvidenceType:                "desktop-activation-transaction-receipt-evidence",
		EvidenceState:               "planned-runtime-gated",
		ReceiptSchemaVersion:        "xnix.runtime.desktop_activation_receipt.v1",
		CommitReceiptType:           "desktop-activation-receipt",
		RollbackReceiptType:         "desktop-activation-rollback-receipt",
		ReceiptRelativePath:         desktopActivationReceiptRelativePath(applicationID),
		RequiredFileIDs:             requiredFileIDs,
		RequiredFileCount:           len(requiredFileIDs),
		InstalledFileDigestRequired: true,
		RollbackDigestRequired:      true,
		CommitReceiptRequired:       true,
		CommitReceiptPlanned:        true,
		CommitReceiptWritten:        false,
		RollbackReceiptRequired:     true,
		RollbackReceiptPlanned:      rollbackReceiptPlanned,
		RollbackReceiptWritten:      false,
		DigestGateReady:             false,
		CommitAvailable:             false,
		RollbackAvailable:           false,
		RuntimeOwned:                true,
		KDEPolicyOwner:              false,
		RootPathExposed:             false,
		TargetRootPathExposed:       false,
		HostRootModified:            false,
		FileWritesPerformed:         false,
		BackendDetailsExposed:       false,
		Summary:                     "Runtime has planned commit and rollback receipt evidence, but no receipt is written and no host root is modified by the preview.",
	}
}

func countDesktopActivationTransactionSteps(steps []DesktopActivationTransactionStep) (int, int) {
	readyCount := 0
	blockedCount := 0
	for _, step := range steps {
		switch step.Status {
		case "ready":
			readyCount++
		case "blocked":
			blockedCount++
		}
	}
	return readyCount, blockedCount
}

func desktopActivationTransactionStepIDs(steps []DesktopActivationTransactionStep) []string {
	ids := make([]string, 0, len(steps))
	for _, step := range steps {
		ids = append(ids, step.ID)
	}
	return ids
}

func desktopActivationRollbackStepIDs(steps []DesktopActivationRollbackStep) []string {
	ids := make([]string, 0, len(steps))
	for _, step := range steps {
		ids = append(ids, step.ID)
	}
	return ids
}

func desktopActivationTransactionState(ready bool) string {
	if ready {
		return "transaction-ready"
	}
	return "transaction-blocked"
}

func desktopActivationTransactionSummary(ready bool) string {
	if ready {
		return "Runtime can describe the desktop activation transaction and rollback sequence, but this preview performs no writes and exposes no target paths."
	}
	return "Runtime can describe the desktop activation transaction, but commit remains blocked until preflight and staging both allow activation."
}
