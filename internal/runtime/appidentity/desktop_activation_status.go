package appidentity

import "errors"

type DesktopActivationStatusPreview struct {
	SchemaVersion               string                                    `json:"schema_version"`
	RequestType                 string                                    `json:"request_type"`
	StatusType                  string                                    `json:"status_type"`
	ActivationState             string                                    `json:"activation_state"`
	Source                      string                                    `json:"source"`
	Desktop                     string                                    `json:"desktop"`
	PlannedRuntimeMethod        string                                    `json:"planned_runtime_method"`
	CurrentRenderer             string                                    `json:"current_renderer"`
	WriteMethod                 string                                    `json:"write_method"`
	ApplicationID               string                                    `json:"application_id"`
	DisplayName                 string                                    `json:"display_name"`
	Icon                        string                                    `json:"icon"`
	DesktopFile                 string                                    `json:"desktop_file"`
	InstallMode                 string                                    `json:"install_mode"`
	PreflightDecision           string                                    `json:"preflight_decision"`
	TransactionState            string                                    `json:"transaction_state"`
	Transaction                 DesktopActivationStatusTransactionSummary `json:"transaction"`
	Staging                     DesktopActivationStatusStagingSummary     `json:"staging"`
	WriteGate                   DesktopActivationStatusWriteGateSummary   `json:"write_gate"`
	CommitGate                  DesktopActivationStatusCommitGateSummary  `json:"commit_gate"`
	KDESurface                  DesktopActivationStatusKDESurfaceSummary  `json:"kde_surface"`
	UserVisible                 bool                                      `json:"user_visible"`
	RuntimeOwned                bool                                      `json:"runtime_owned"`
	GoRuntimeBacked             bool                                      `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                      `json:"kde_policy_owner"`
	ActivationReady             bool                                      `json:"activation_ready"`
	ActivationCommitted         bool                                      `json:"activation_committed"`
	CommitEnabled               bool                                      `json:"commit_enabled"`
	InstallerMayProceed         bool                                      `json:"installer_may_proceed"`
	StagingPlanReady            bool                                      `json:"staging_plan_ready"`
	RollbackPlanned             bool                                      `json:"rollback_planned"`
	RollbackAvailable           bool                                      `json:"rollback_available"`
	HostRootModified            bool                                      `json:"host_root_modified"`
	FileWritesPerformed         bool                                      `json:"file_writes_performed"`
	DesktopFilesWritten         bool                                      `json:"desktop_files_written"`
	MIMEAppsWritten             bool                                      `json:"mimeapps_written"`
	KDEServiceCacheRefreshed    bool                                      `json:"kde_service_cache_refreshed"`
	LaunchEnabled               bool                                      `json:"launch_enabled"`
	BackendLaunchEnabled        bool                                      `json:"backend_launch_enabled"`
	ExecutionStarted            bool                                      `json:"execution_started"`
	NetworkRequired             bool                                      `json:"network_required"`
	PrivilegedContainerRequired bool                                      `json:"privileged_container_required"`
	BackendDetailsExposed       bool                                      `json:"backend_details_exposed"`
	RawWindowsExecutableExposed bool                                      `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed bool                                      `json:"compatibility_storage_exposed"`
	StatusSignals               []DesktopActivationStatusSignal           `json:"status_signals"`
	StatusSignalIDs             []string                                  `json:"status_signal_ids"`
	BlockedReasons              []DesktopActivationStatusBlockedReason    `json:"blocked_reasons"`
	BlockedReasonIDs            []string                                  `json:"blocked_reason_ids"`
	NextSafeActions             []DesktopActivationStatusNextAction       `json:"next_safe_actions"`
	NextSafeActionIDs           []string                                  `json:"next_safe_action_ids"`
	DesktopSafeSummary          string                                    `json:"desktop_safe_summary"`
}

type DesktopActivationStatusTransactionSummary struct {
	RequestType          string `json:"request_type"`
	TransactionState     string `json:"transaction_state"`
	TransactionStepCount int    `json:"transaction_step_count"`
	ReadyStepCount       int    `json:"ready_step_count"`
	BlockedStepCount     int    `json:"blocked_step_count"`
	RollbackStepCount    int    `json:"rollback_step_count"`
	TransactionReady     bool   `json:"transaction_ready"`
	TransactionCommitted bool   `json:"transaction_committed"`
}

type DesktopActivationStatusStagingSummary struct {
	RequestType              string `json:"request_type"`
	StagingState             string `json:"staging_state"`
	PlannedFileCount         int    `json:"planned_file_count"`
	ActivatedEntryPointCount int    `json:"activated_entry_point_count"`
	ReceiptFileID            string `json:"receipt_file_id"`
	StagingPlanReady         bool   `json:"staging_plan_ready"`
	InstallerMayProceed      bool   `json:"installer_may_proceed"`
	HostRootAllowed          bool   `json:"host_root_allowed"`
}

type DesktopActivationStatusWriteGateSummary struct {
	MethodName         string `json:"method_name"`
	GateDecision       string `json:"gate_decision"`
	DenialErrorName    string `json:"denial_error_name"`
	WriteMethodEnabled bool   `json:"write_method_enabled"`
	DispatchEnabled    bool   `json:"dispatch_enabled"`
}

type DesktopActivationStatusCommitGateSummary struct {
	CommitState             string `json:"commit_state"`
	CommitEnabled           bool   `json:"commit_enabled"`
	RequiresDigestMatch     bool   `json:"requires_digest_match"`
	RequiresRollbackReceipt bool   `json:"requires_rollback_receipt"`
	RequiresRuntimeOwner    bool   `json:"requires_runtime_owner"`
	RequiresUserReview      bool   `json:"requires_user_review"`
	Summary                 string `json:"summary"`
}

type DesktopActivationStatusKDESurfaceSummary struct {
	Desktop                  string `json:"desktop"`
	EntryPointCount          int    `json:"entry_point_count"`
	MenuEntryReady           bool   `json:"menu_entry_ready"`
	TaskManagerReady         bool   `json:"task_manager_ready"`
	FileManagerReady         bool   `json:"file_manager_ready"`
	TrayStatusReady          bool   `json:"tray_status_ready"`
	NotificationReady        bool   `json:"notification_ready"`
	CompatibilityCenterReady bool   `json:"compatibility_center_ready"`
	UnifiedSettingsReady     bool   `json:"unified_settings_ready"`
}

type DesktopActivationStatusSignal struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type DesktopActivationStatusBlockedReason struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Summary  string `json:"summary"`
}

type DesktopActivationStatusNextAction struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Summary string `json:"summary"`
}

func (plan Plan) DesktopActivationStatusPreview(mode string) (DesktopActivationStatusPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationStatusPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationStatusPreview{}, errors.New("desktop activation status requires single-line identity fields")
		}
	}

	transaction, err := plan.DesktopActivationTransactionPreview(mode)
	if err != nil {
		return DesktopActivationStatusPreview{}, err
	}

	activationReady := transaction.TransactionReady && transaction.StagingPlanReady && transaction.InstallerMayProceed
	signals := desktopActivationStatusSignals(transaction)
	blockedReasons := desktopActivationBlockedReasons(transaction)
	nextActions := desktopActivationNextActions(activationReady)
	preview := DesktopActivationStatusPreview{
		SchemaVersion:        "xnix.runtime.desktop_activation_status.v1",
		RequestType:          "desktop-activation-status-preview",
		StatusType:           "kde-desktop-activation-status",
		ActivationState:      desktopActivationStatusState(activationReady),
		Source:               "desktop-activation-transaction-preview",
		Desktop:              "KDE Plasma",
		PlannedRuntimeMethod: "GetDesktopActivationStatus",
		CurrentRenderer:      "xnix-runtime-go desktop-activation-status-preview",
		WriteMethod:          transaction.WriteMethod,
		ApplicationID:        plan.ApplicationID,
		DisplayName:          plan.DisplayName,
		Icon:                 plan.Icon,
		DesktopFile:          plan.DesktopFile,
		InstallMode:          transaction.InstallMode,
		PreflightDecision:    transaction.PreflightDecision,
		TransactionState:     transaction.TransactionState,
		Transaction: DesktopActivationStatusTransactionSummary{
			RequestType:          transaction.RequestType,
			TransactionState:     transaction.TransactionState,
			TransactionStepCount: transaction.TransactionStepCount,
			ReadyStepCount:       transaction.ReadyStepCount,
			BlockedStepCount:     transaction.BlockedStepCount,
			RollbackStepCount:    transaction.RollbackStepCount,
			TransactionReady:     transaction.TransactionReady,
			TransactionCommitted: transaction.TransactionCommitted,
		},
		Staging: DesktopActivationStatusStagingSummary{
			RequestType:              transaction.Staging.RequestType,
			StagingState:             transaction.Staging.StagingState,
			PlannedFileCount:         transaction.Staging.PlannedFileCount,
			ActivatedEntryPointCount: transaction.Staging.ActivatedEntryPointCount,
			ReceiptFileID:            transaction.Staging.ReceiptFileID,
			StagingPlanReady:         transaction.Staging.StagingPlanReady,
			InstallerMayProceed:      transaction.Staging.InstallerMayProceed,
			HostRootAllowed:          transaction.Staging.HostRootAllowed,
		},
		WriteGate: DesktopActivationStatusWriteGateSummary{
			MethodName:         transaction.WriteGate.MethodName,
			GateDecision:       transaction.WriteGate.GateDecision,
			DenialErrorName:    transaction.WriteGate.DenialErrorName,
			WriteMethodEnabled: transaction.WriteGate.WriteMethodEnabled,
			DispatchEnabled:    transaction.WriteGate.DispatchEnabled,
		},
		CommitGate: DesktopActivationStatusCommitGateSummary{
			CommitState:             desktopActivationCommitState(activationReady),
			CommitEnabled:           false,
			RequiresDigestMatch:     true,
			RequiresRollbackReceipt: true,
			RequiresRuntimeOwner:    true,
			RequiresUserReview:      true,
			Summary:                 "Activation status is visible to KDE, but commit remains owned by the Runtime write gate.",
		},
		KDESurface: DesktopActivationStatusKDESurfaceSummary{
			Desktop:                  "KDE Plasma",
			EntryPointCount:          transaction.Staging.ActivatedEntryPointCount,
			MenuEntryReady:           activationReady,
			TaskManagerReady:         activationReady,
			FileManagerReady:         activationReady,
			TrayStatusReady:          activationReady,
			NotificationReady:        activationReady,
			CompatibilityCenterReady: true,
			UnifiedSettingsReady:     true,
		},
		UserVisible:                 true,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ActivationReady:             activationReady,
		ActivationCommitted:         false,
		CommitEnabled:               false,
		InstallerMayProceed:         transaction.InstallerMayProceed,
		StagingPlanReady:            transaction.StagingPlanReady,
		RollbackPlanned:             transaction.RollbackReceiptPlanned,
		RollbackAvailable:           false,
		HostRootModified:            false,
		FileWritesPerformed:         false,
		DesktopFilesWritten:         false,
		MIMEAppsWritten:             false,
		KDEServiceCacheRefreshed:    false,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		RawWindowsExecutableExposed: false,
		CompatibilityStorageExposed: false,
		StatusSignals:               signals,
		StatusSignalIDs:             desktopActivationStatusSignalIDs(signals),
		BlockedReasons:              blockedReasons,
		BlockedReasonIDs:            desktopActivationBlockedReasonIDs(blockedReasons),
		NextSafeActions:             nextActions,
		NextSafeActionIDs:           desktopActivationNextActionIDs(nextActions),
		DesktopSafeSummary:          desktopActivationStatusSummary(activationReady),
	}
	if err := validateNoBackendTerms(preview, "desktop activation status preview"); err != nil {
		return DesktopActivationStatusPreview{}, err
	}
	return preview, nil
}

func desktopActivationStatusSignals(transaction DesktopActivationTransactionPreview) []DesktopActivationStatusSignal {
	return []DesktopActivationStatusSignal{
		desktopActivationStatusSignal("transaction-plan", transaction.TransactionState, "Runtime has built an auditable desktop activation transaction preview."),
		desktopActivationStatusSignal("staging-plan", transaction.Staging.StagingState, "Runtime has planned the desktop activation file set without writing it."),
		desktopActivationStatusSignal("write-gate", transaction.WriteGate.GateDecision, "Runtime write methods remain closed for preview-only inspection."),
		desktopActivationStatusSignal("rollback-plan", "planned", "Runtime has planned the rollback receipt and rollback sequence."),
		desktopActivationStatusSignal("kde-surface", "visible", "KDE may display activation readiness without owning activation policy."),
	}
}

func desktopActivationStatusSignal(id string, status string, summary string) DesktopActivationStatusSignal {
	return DesktopActivationStatusSignal{ID: id, Status: status, Summary: summary}
}

func desktopActivationBlockedReasons(transaction DesktopActivationTransactionPreview) []DesktopActivationStatusBlockedReason {
	reasons := []DesktopActivationStatusBlockedReason{
		desktopActivationBlockedReason("write-method-disabled", "required", "Activation commit requires the Runtime write gate to open."),
		desktopActivationBlockedReason("digest-verification-pending", "required", "Staged desktop artifact digests must be verified before commit."),
		desktopActivationBlockedReason("rollback-receipt-not-written", "required", "Rollback receipt must be written during the commit transaction."),
		desktopActivationBlockedReason("production-owner-not-active", "required", "A production Runtime owner must perform activation writes."),
	}
	if !transaction.TransactionReady {
		reasons = append(reasons, desktopActivationBlockedReason("transaction-not-ready", "blocking", "Activation transaction is blocked by preflight or staging state."))
	}
	return reasons
}

func desktopActivationBlockedReason(id string, severity string, summary string) DesktopActivationStatusBlockedReason {
	return DesktopActivationStatusBlockedReason{ID: id, Severity: severity, Summary: summary}
}

func desktopActivationNextActions(activationReady bool) []DesktopActivationStatusNextAction {
	return []DesktopActivationStatusNextAction{
		desktopActivationNextAction("show-compatibility-center-status", true, "Show activation readiness in the KDE Compatibility Center."),
		desktopActivationNextAction("show-user-review-card", true, "Explain that activation commit is still gated by Runtime policy."),
		desktopActivationNextAction("commit-desktop-activation", false, "Keep desktop activation commit disabled until production Runtime gates pass."),
		desktopActivationNextAction("launch-application", false, "Keep application launch disabled until activation and backend launch gates pass."),
		desktopActivationNextAction("refresh-kde-service-cache", false, "Refresh KDE service metadata only inside a committed activation transaction."),
	}
}

func desktopActivationNextAction(id string, enabled bool, summary string) DesktopActivationStatusNextAction {
	return DesktopActivationStatusNextAction{ID: id, Enabled: enabled, Summary: summary}
}

func desktopActivationStatusSignalIDs(signals []DesktopActivationStatusSignal) []string {
	ids := make([]string, 0, len(signals))
	for _, signal := range signals {
		ids = append(ids, signal.ID)
	}
	return ids
}

func desktopActivationBlockedReasonIDs(reasons []DesktopActivationStatusBlockedReason) []string {
	ids := make([]string, 0, len(reasons))
	for _, reason := range reasons {
		ids = append(ids, reason.ID)
	}
	return ids
}

func desktopActivationNextActionIDs(actions []DesktopActivationStatusNextAction) []string {
	ids := make([]string, 0, len(actions))
	for _, action := range actions {
		ids = append(ids, action.ID)
	}
	return ids
}

func desktopActivationStatusState(ready bool) string {
	if ready {
		return "ready-for-runtime-commit"
	}
	return "blocked-before-runtime-commit"
}

func desktopActivationCommitState(ready bool) string {
	if ready {
		return "commit-gated"
	}
	return "commit-blocked"
}

func desktopActivationStatusSummary(ready bool) string {
	if ready {
		return "KDE can show that desktop activation is planned and ready for a future Runtime-owned commit, while writes and launch remain disabled."
	}
	return "KDE can show desktop activation blockers without exposing backend details or enabling writes."
}
