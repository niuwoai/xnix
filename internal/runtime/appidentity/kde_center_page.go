package appidentity

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/artifact"
	"xnix.local/xnix/internal/runtime/winapp"
)

type KDECenterPagePreview struct {
	SchemaVersion                            string                                 `json:"schema_version"`
	RequestType                              string                                 `json:"request_type"`
	PageType                                 string                                 `json:"page_type"`
	Source                                   string                                 `json:"source"`
	Desktop                                  string                                 `json:"desktop"`
	RuntimeMethod                            string                                 `json:"runtime_method"`
	ReadMethod                               string                                 `json:"read_method"`
	ApplicationID                            string                                 `json:"application_id"`
	ApplicationName                          string                                 `json:"application_name"`
	Icon                                     string                                 `json:"icon"`
	DesktopFile                              string                                 `json:"desktop_file"`
	LauncherCommand                          []string                               `json:"launcher_command"`
	Header                                   KDECenterPageHeader                    `json:"header"`
	ApplicationSummary                       KDECenterPageApplication               `json:"application_summary"`
	KnownAppSessionGateEvidenceCount         int                                    `json:"known_app_session_gate_evidence_count"`
	KnownAppLauncherSessionGateConsumedCount int                                    `json:"known_app_launcher_session_gate_consumed_count"`
	KnownAppPostReviewDispatchConsumedCount  int                                    `json:"known_app_post_review_dispatch_consumed_count"`
	KnownAppSessionGateCards                 []KDECenterPageKnownAppSessionGateCard `json:"known_app_session_gate_cards"`
	KnownAppMatrixEvidenceCount              int                                    `json:"known_app_matrix_evidence_count"`
	KnownAppMatrixEvidenceCards              []KDECenterPageKnownAppMatrixCard      `json:"known_app_matrix_evidence_cards"`
	KnownAppGUIEvidenceCount                 int                                    `json:"known_app_gui_evidence_count"`
	KnownAppOwnerControlledGUIEvidenceCount  int                                    `json:"known_app_owner_controlled_gui_evidence_count"`
	KnownAppOwnerManagedCopyVerifiedCount    int                                    `json:"known_app_owner_managed_copy_verified_count"`
	KnownAppGUIEvidenceCards                 []KDECenterPageKnownAppMatrixCard      `json:"known_app_gui_evidence_cards"`
	BackendSelectionSnapshot                 KDECenterPageBackend                   `json:"backend_selection_snapshot"`
	ActivationStatusSnapshot                 KDECenterPageActivation                `json:"activation_status_snapshot"`
	ExecutionReadinessSnapshot               KDECenterPageExecution                 `json:"execution_readiness_snapshot"`
	ApplicationReadinessEvidence             ApplicationReadinessPreview            `json:"application_readiness_evidence"`
	LaunchIntentSnapshot                     KDECenterPageLaunchIntent              `json:"launch_intent_snapshot"`
	WindowIdentitySnapshot                   KDECenterPageWindow                    `json:"window_identity_snapshot"`
	FileAssociationSnapshot                  KDECenterPageFiles                     `json:"file_association_snapshot"`
	TrayStatusSnapshot                       KDECenterPageTray                      `json:"tray_status_snapshot"`
	NotificationSnapshot                     KDECenterPageNotification              `json:"notification_snapshot"`
	ActionDeck                               KDECenterPageActionDeck                `json:"action_deck"`
	ActionDependencyGraph                    KDECenterPageActionGraph               `json:"action_dependency_graph"`
	SettingsSnapshot                         KDECenterPageSettings                  `json:"settings_snapshot"`
	Navigation                               []KDECenterPageNavigation              `json:"navigation"`
	NavigationCount                          int                                    `json:"navigation_count"`
	PrimaryNavigationTarget                  string                                 `json:"primary_navigation_target"`
	RuntimeOwned                             bool                                   `json:"runtime_owned"`
	GoRuntimeBacked                          bool                                   `json:"go_runtime_backed"`
	KDEPolicyOwner                           bool                                   `json:"kde_policy_owner"`
	OfficialDesktopOnly                      bool                                   `json:"official_desktop_only"`
	UserVisible                              bool                                   `json:"user_visible"`
	SafeForAIDiagnostics                     bool                                   `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured                     bool                                   `json:"user_decision_captured"`
	UserDecisionAllowsLaunch                 bool                                   `json:"user_decision_allows_launch"`
	PagePreviewCreated                       bool                                   `json:"page_preview_created"`
	PagePersisted                            bool                                   `json:"page_persisted"`
	DeckPersisted                            bool                                   `json:"deck_persisted"`
	CardsPersisted                           bool                                   `json:"cards_persisted"`
	CardActionsEnabled                       bool                                   `json:"card_actions_enabled"`
	SettingsPersisted                        bool                                   `json:"settings_persisted"`
	SettingsPersistenceEnabled               bool                                   `json:"settings_persistence_enabled"`
	NotificationsSent                        bool                                   `json:"notifications_sent"`
	ResourceGrantCreated                     bool                                   `json:"resource_grant_created"`
	RuntimeLaunchApproval                    bool                                   `json:"runtime_launch_approval"`
	LaunchAllowed                            bool                                   `json:"launch_allowed"`
	LaunchEnabled                            bool                                   `json:"launch_enabled"`
	ExecutionStarted                         bool                                   `json:"execution_started"`
	BackendProcessStarted                    bool                                   `json:"backend_process_started"`
	RequestObjectsCreated                    bool                                   `json:"request_objects_created"`
	PermissionGrantCreated                   bool                                   `json:"permission_grant_created"`
	HostRootModified                         bool                                   `json:"host_root_modified"`
	NetworkRequired                          bool                                   `json:"network_required"`
	BackendDetailsExposed                    bool                                   `json:"backend_details_exposed"`
	BlockedActions                           []string                               `json:"blocked_actions"`
	UserFacingSettings                       map[string]string                      `json:"user_facing_settings"`
	DesktopSafeSummary                       string                                 `json:"desktop_safe_summary"`
}

type KDECenterPageHeader struct {
	Title                 string `json:"title"`
	Subtitle              string `json:"subtitle"`
	Badge                 string `json:"badge"`
	BadgeTone             string `json:"badge_tone"`
	PrimaryActionLabel    string `json:"primary_action_label"`
	PrimaryActionTarget   string `json:"primary_action_target"`
	PrimaryActionEnabled  bool   `json:"primary_action_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type KDECenterPageApplication struct {
	ApplicationID              string   `json:"application_id"`
	DisplayName                string   `json:"display_name"`
	CompatibilityState         string   `json:"compatibility_state"`
	CompatibilityLabel         string   `json:"compatibility_label"`
	DiagnosticsState           string   `json:"diagnostics_state"`
	RuntimeMode                string   `json:"runtime_mode"`
	SupportedExtensions        []string `json:"supported_extensions"`
	KnownIssueCount            int      `json:"known_issue_count"`
	RepairRecordState          string   `json:"repair_record_state"`
	RepairRecordCount          int      `json:"repair_record_count"`
	ActionExecutionEnabled     bool     `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool     `json:"repair_execution_enabled"`
	BackendLaunchEnabled       bool     `json:"backend_launch_enabled"`
	SettingsPersistenceEnabled bool     `json:"settings_persistence_enabled"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	Summary                    string   `json:"summary"`
}

type KDECenterPageKnownAppSessionGateCard struct {
	AppID                                      string   `json:"app_id"`
	DisplayName                                string   `json:"display_name"`
	AppVersion                                 string   `json:"app_version"`
	CompatibilityState                         string   `json:"compatibility_state"`
	CenterCardState                            string   `json:"center_card_state"`
	ControlledExecutionSessionID               string   `json:"controlled_execution_session_id"`
	LaunchAuthorizationReceiptID               string   `json:"launch_authorization_receipt_id"`
	LauncherSessionGateConsumed                bool     `json:"launcher_session_gate_consumed"`
	LauncherSessionDigestVerified              bool     `json:"launcher_session_digest_verified"`
	LauncherSessionRelativePath                string   `json:"launcher_session_relative_path"`
	RuntimeOwnerConsumableSession              bool     `json:"runtime_owner_consumable_session"`
	KDEReadModelConsumableSession              bool     `json:"kde_read_model_consumable_session"`
	PostReviewDispatchConsumed                 bool     `json:"post_review_dispatch_consumed"`
	PostReviewDispatchState                    string   `json:"post_review_dispatch_state"`
	SessionGatedReviewReceiptID                string   `json:"session_gated_review_receipt_id"`
	LaunchGateConsumed                         bool     `json:"launch_gate_consumed"`
	ControlledDispatchReady                    bool     `json:"controlled_dispatch_ready"`
	PrimaryActionID                            string   `json:"primary_action_id"`
	PrimaryActionLabel                         string   `json:"primary_action_label"`
	PrimaryActionKind                          string   `json:"primary_action_kind"`
	PrimaryActionEnabled                       bool     `json:"primary_action_enabled"`
	RuntimeStatusLaunchRequestType             string   `json:"runtime_status_launch_request_type"`
	RuntimeStatusLaunchRuntimeMethod           string   `json:"runtime_status_launch_runtime_method"`
	RuntimeStatusLaunchReadMethod              string   `json:"runtime_status_launch_read_method"`
	RuntimeStatusLaunchRequiredIDCount         int      `json:"runtime_status_launch_required_id_count"`
	RuntimeStatusLaunchCollectedIDCount        int      `json:"runtime_status_launch_collected_id_count"`
	RuntimeStatusLaunchManagedLauncherArgv     []string `json:"runtime_status_launch_managed_launcher_argv"`
	RuntimeStatusLaunchRequestReady            bool     `json:"runtime_status_launch_request_ready"`
	RuntimeStatusLaunchStateRootRequired       bool     `json:"runtime_status_launch_state_root_required"`
	RuntimeStatusLaunchStateRootOwnedByRuntime bool     `json:"runtime_status_launch_state_root_owned_by_runtime"`
	ReviewRouteRequestType                     string   `json:"review_route_request_type"`
	ReviewRouteRuntimeMethod                   string   `json:"review_route_runtime_method"`
	ReviewRouteReadMethod                      string   `json:"review_route_read_method"`
	ReadBeforeWriteRequired                    bool     `json:"read_before_write_required"`
	RuntimeReceiptRequired                     bool     `json:"runtime_receipt_required"`
	UserVisible                                bool     `json:"user_visible"`
	RuntimeOwned                               bool     `json:"runtime_owned"`
	GoRuntimeBacked                            bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                             bool     `json:"kde_policy_owner"`
	DesktopLaunchEnabled                       bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                       bool     `json:"backend_launch_enabled"`
	HostRootModified                           bool     `json:"host_root_modified"`
	BackendDetailsExposed                      bool     `json:"backend_details_exposed"`
	Summary                                    string   `json:"summary"`
}

type KDECenterPageKnownAppMatrixCard struct {
	AppID                                string   `json:"app_id"`
	DisplayName                          string   `json:"display_name"`
	AppVersion                           string   `json:"app_version"`
	EvidenceKind                         string   `json:"evidence_kind"`
	EvidenceSource                       string   `json:"evidence_source"`
	SmokeStatus                          string   `json:"smoke_status"`
	CompatibilityState                   string   `json:"compatibility_state"`
	CenterCardState                      string   `json:"center_card_state"`
	PrimaryActionID                      string   `json:"primary_action_id"`
	PrimaryActionLabel                   string   `json:"primary_action_label"`
	PrimaryActionKind                    string   `json:"primary_action_kind"`
	PrimaryActionEnabled                 bool     `json:"primary_action_enabled"`
	DesktopCallableRoute                 string   `json:"desktop_callable_route,omitempty"`
	DesktopCallableRuntimeMethod         string   `json:"desktop_callable_runtime_method,omitempty"`
	DesktopCallableExecutionType         string   `json:"desktop_callable_execution_type,omitempty"`
	DesktopDBusMethod                    string   `json:"desktop_dbus_method,omitempty"`
	DesktopEvidenceHandleForwarded       bool     `json:"desktop_evidence_handle_forwarded"`
	KDEForwardedArgumentKind             string   `json:"kde_forwarded_argument_kind,omitempty"`
	KDEForwardedArguments                []string `json:"kde_forwarded_arguments,omitempty"`
	OwnerServiceArgsExposedToKDE         bool     `json:"owner_service_args_exposed_to_kde"`
	MarkerObserved                       bool     `json:"marker_observed"`
	ChecksumVerified                     bool     `json:"checksum_verified"`
	ExecutionEvidenceRecorded            bool     `json:"execution_evidence_recorded"`
	StagedLauncherVerified               bool     `json:"staged_launcher_verified"`
	OwnerControlledRuntimeLaunchVerified bool     `json:"owner_controlled_runtime_launch_verified"`
	OwnerManagedCopyVerified             bool     `json:"owner_managed_copy_verified"`
	OwnerServiceCallReady                bool     `json:"owner_service_call_ready"`
	OwnerEvidenceHandoffReady            bool     `json:"owner_evidence_handoff_ready"`
	OwnerEvidenceRelativePath            string   `json:"owner_evidence_relative_path,omitempty"`
	RuntimeDispatchVerified              bool     `json:"runtime_dispatch_verified"`
	LaunchAuthorizationRequired          bool     `json:"launch_authorization_required"`
	DesktopLaunchEnabled                 bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled                 bool     `json:"backend_launch_enabled"`
	RuntimeOwned                         bool     `json:"runtime_owned"`
	GoRuntimeBacked                      bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool     `json:"kde_policy_owner"`
	HostRootModified                     bool     `json:"host_root_modified"`
	BackendDetailsExposed                bool     `json:"backend_details_exposed"`
	RawArtifactPathExposed               bool     `json:"raw_artifact_path_exposed"`
	Summary                              string   `json:"summary"`
}

type KDECenterPageBackend struct {
	RequestType                 string `json:"request_type"`
	PlanType                    string `json:"plan_type"`
	RuntimeMethod               string `json:"runtime_method"`
	SelectedStrategy            string `json:"selected_strategy"`
	RecommendedProfileID        string `json:"recommended_profile_id"`
	CandidateCount              int    `json:"candidate_count"`
	ReadyCandidateCount         int    `json:"ready_candidate_count"`
	BlockedCandidateCount       int    `json:"blocked_candidate_count"`
	SelectionCommitted          bool   `json:"selection_committed"`
	SelectionChangeEnabled      bool   `json:"selection_change_enabled"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	CapabilityActivationEnabled bool   `json:"capability_activation_enabled"`
	EnvironmentCreated          bool   `json:"environment_created"`
	HostRootModified            bool   `json:"host_root_modified"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	Summary                     string `json:"summary"`
}

type KDECenterPageActivation struct {
	RequestType           string `json:"request_type"`
	StatusType            string `json:"status_type"`
	RuntimeMethod         string `json:"runtime_method"`
	Renderer              string `json:"renderer"`
	ActivationState       string `json:"activation_state"`
	TransactionState      string `json:"transaction_state"`
	PreflightDecision     string `json:"preflight_decision"`
	ReceiptEvidenceState  string `json:"receipt_evidence_state"`
	ReceiptRelativePath   string `json:"receipt_relative_path"`
	ReceiptBacked         bool   `json:"receipt_backed"`
	RollbackAvailable     bool   `json:"rollback_available"`
	StatusSignalCount     int    `json:"status_signal_count"`
	BlockedReasonCount    int    `json:"blocked_reason_count"`
	NextSafeActionCount   int    `json:"next_safe_action_count"`
	ActivationReady       bool   `json:"activation_ready"`
	ActivationCommitted   bool   `json:"activation_committed"`
	CommitEnabled         bool   `json:"commit_enabled"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

type KDECenterPageOptions struct {
	ActivationRoot                      string
	ExecutionSessionRoot                string
	ExecutionSessionRequestID           string
	ApplicationReadinessRoot            string
	ApplicationReadinessStateRoot       string
	ApplicationReadinessArtifactReceipt *artifact.StageReceipt
	ApplicationReadinessPortalOperation string
	ApplicationReadinessSnapshotReason  string
	KnownAppSmokeEvidence               []KnownAppSmokeEvidenceSummary
}

type KDECenterPageExecution struct {
	RequestType               string `json:"request_type"`
	ReadinessType             string `json:"readiness_type"`
	RuntimeMethod             string `json:"runtime_method"`
	ExecutionState            string `json:"execution_state"`
	OverallStatus             string `json:"overall_status"`
	RecommendedAction         string `json:"recommended_action"`
	GateCount                 int    `json:"gate_count"`
	RequiredGateCount         int    `json:"required_gate_count"`
	PendingGateCount          int    `json:"pending_gate_count"`
	BlockedGateCount          int    `json:"blocked_gate_count"`
	DesktopEntryLaunchVisible bool   `json:"desktop_entry_launch_visible"`
	LaunchAllowed             bool   `json:"launch_allowed"`
	LaunchEnabled             bool   `json:"launch_enabled"`
	ExecutionRequestCreated   bool   `json:"execution_request_created"`
	BackendBindingReady       bool   `json:"backend_binding_ready"`
	PortalPolicyRequired      bool   `json:"portal_policy_required"`
	SnapshotRequired          bool   `json:"snapshot_required"`
	UserActionRequired        bool   `json:"user_action_required"`
	HostRootModified          bool   `json:"host_root_modified"`
	NetworkRequired           bool   `json:"network_required"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
	Summary                   string `json:"summary"`
}

type KDECenterPageLaunchIntent struct {
	RequestType             string `json:"request_type"`
	IntentType              string `json:"intent_type"`
	Source                  string `json:"source"`
	RuntimeMethod           string `json:"runtime_method"`
	ReadMethod              string `json:"read_method"`
	FileCount               int    `json:"file_count"`
	PortalRequired          bool   `json:"portal_required"`
	SnapshotRequired        bool   `json:"snapshot_required"`
	StandardDesktopEntry    bool   `json:"standard_desktop_entry"`
	LaunchUsesRuntime       bool   `json:"launch_uses_runtime"`
	LaunchAllowed           bool   `json:"launch_allowed"`
	LaunchEnabled           bool   `json:"launch_enabled"`
	ExecutionRequestCreated bool   `json:"execution_request_created"`
	ExecutionStarted        bool   `json:"execution_started"`
	BackendBindingReady     bool   `json:"backend_binding_ready"`
	RequestObjectCreated    bool   `json:"request_object_created"`
	PermissionGranted       bool   `json:"permission_granted"`
	HostRootModified        bool   `json:"host_root_modified"`
	NetworkRequired         bool   `json:"network_required"`
	BackendDetailsExposed   bool   `json:"backend_details_exposed"`
	Summary                 string `json:"summary"`
}

type KDECenterPageWindow struct {
	SchemaVersion             string `json:"schema_version"`
	DesktopFile               string `json:"desktop_file"`
	LauncherURL               string `json:"launcher_url"`
	WindowKind                string `json:"window_kind"`
	ClassGroup                string `json:"class_group"`
	ResourceName              string `json:"resource_name"`
	TitleHint                 string `json:"title_hint"`
	TaskManagerGroupingKey    string `json:"task_manager_grouping_key"`
	TaskManagerPinningAllowed bool   `json:"task_manager_pinning_allowed"`
	TaskManagerRestoreAllowed bool   `json:"task_manager_restore_allowed"`
	TaskManagerSkipTaskbar    bool   `json:"task_manager_skip_taskbar"`
	TaskManagerShowInSwitcher bool   `json:"task_manager_show_in_switcher"`
	PreferExistingWindow      bool   `json:"prefer_existing_window"`
	KWinScriptRole            string `json:"kwin_script_role"`
	KWinPlacement             string `json:"kwin_placement"`
	WindowManagerPolicyOnly   bool   `json:"window_manager_policy_only"`
	RuntimeOwnsBackendPolicy  bool   `json:"runtime_owns_backend_policy"`
	ExecutionSessionRoot      bool   `json:"execution_session_root"`
	ExecutionSessionBacked    bool   `json:"execution_session_backed"`
	ExecutionSessionPath      string `json:"execution_session_path,omitempty"`
	TaskManagerSessionState   string `json:"task_manager_session_state,omitempty"`
	KWinSessionState          string `json:"kwin_session_state,omitempty"`
	TaskManagerEntryActive    bool   `json:"task_manager_entry_active"`
	KWinRuleApplied           bool   `json:"kwin_rule_applied"`
	HostRootModified          bool   `json:"host_root_modified"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
	Summary                   string `json:"summary"`
}

type KDECenterPageFiles struct {
	PlanType                  string   `json:"plan_type"`
	AssociationType           string   `json:"association_type"`
	RuntimeMethod             string   `json:"runtime_method"`
	DesktopFile               string   `json:"desktop_file"`
	MIMEAppsPath              string   `json:"mimeapps_path"`
	MIMETypes                 []string `json:"mime_types"`
	MIMETypeCount             int      `json:"mime_type_count"`
	FileOpenCommand           string   `json:"file_open_command"`
	FileOpenArgument          string   `json:"file_open_argument"`
	StandardMIMEAppsList      bool     `json:"standard_mimeapps_list"`
	StagedRootOnly            bool     `json:"staged_root_only"`
	OverwriteExistingMIMEApps bool     `json:"overwrite_existing_mimeapps"`
	PortalRequiredForFileOpen bool     `json:"portal_required_for_file_open"`
	FileAssociationReady      bool     `json:"file_association_ready"`
	FileOpenPreviewAvailable  bool     `json:"file_open_preview_available"`
	DirectHostFileAccess      bool     `json:"direct_host_file_access"`
	RequestObjectCreated      bool     `json:"request_object_created"`
	PermissionGranted         bool     `json:"permission_granted"`
	FilesWritten              bool     `json:"files_written"`
	MIMEAppsWritten           bool     `json:"mimeapps_written"`
	HostRootModified          bool     `json:"host_root_modified"`
	BackendDetailsExposed     bool     `json:"backend_details_exposed"`
	Summary                   string   `json:"summary"`
}

type KDECenterPageTray struct {
	StatusType                   string   `json:"status_type"`
	RuntimeMethod                string   `json:"runtime_method"`
	DesktopFile                  string   `json:"desktop_file"`
	RegisteredApplicationCount   int      `json:"registered_application_count"`
	ActiveApplicationCount       int      `json:"active_application_count"`
	AttentionRequiredCount       int      `json:"attention_required_count"`
	CompatibilityState           string   `json:"compatibility_state"`
	CompatibilityLabel           string   `json:"compatibility_label"`
	TrayBridgeState              string   `json:"tray_bridge_state"`
	TrayBridgeLabel              string   `json:"tray_bridge_label"`
	BridgedTrayApplicationCount  int      `json:"bridged_tray_application_count"`
	Actions                      []string `json:"actions"`
	ActionCount                  int      `json:"action_count"`
	UserVisible                  bool     `json:"user_visible"`
	TrayStatusReady              bool     `json:"tray_status_ready"`
	ExecutionSessionRoot         bool     `json:"execution_session_root"`
	ExecutionSessionBacked       bool     `json:"execution_session_backed"`
	ExecutionSessionPath         string   `json:"execution_session_path,omitempty"`
	ExecutionSessionState        string   `json:"execution_session_state,omitempty"`
	LiveBackendBridgeEnabled     bool     `json:"live_backend_bridge_enabled"`
	BridgeConfigurationPersisted bool     `json:"bridge_configuration_persisted"`
	HostRootModified             bool     `json:"host_root_modified"`
	BackendDetailsExposed        bool     `json:"backend_details_exposed"`
	Summary                      string   `json:"summary"`
}

type KDECenterPageNotification struct {
	RequestType                string   `json:"request_type"`
	RuntimeMethod              string   `json:"runtime_method"`
	Source                     string   `json:"source"`
	DesktopFile                string   `json:"desktop_file"`
	EventType                  string   `json:"event_type"`
	NotificationID             string   `json:"notification_id"`
	Urgency                    string   `json:"urgency"`
	Category                   string   `json:"category"`
	Title                      string   `json:"title"`
	ActionCount                int      `json:"action_count"`
	Actions                    []string `json:"actions"`
	RequiresUserReview         bool     `json:"requires_user_review"`
	UserVisible                bool     `json:"user_visible"`
	NotificationReady          bool     `json:"notification_ready"`
	ActionExecutionEnabled     bool     `json:"action_execution_enabled"`
	RepairExecutionEnabled     bool     `json:"repair_execution_enabled"`
	SettingsPersistenceEnabled bool     `json:"settings_persistence_enabled"`
	NotificationsSent          bool     `json:"notifications_sent"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	Summary                    string   `json:"summary"`
}

type KDECenterPageActionDeck struct {
	RequestType           string             `json:"request_type"`
	DeckType              string             `json:"deck_type"`
	CardCount             int                `json:"card_count"`
	WaitingCardCount      int                `json:"waiting_card_count"`
	DeferredCardCount     int                `json:"deferred_card_count"`
	RejectedCardCount     int                `json:"rejected_card_count"`
	AIAnalysis            *KDEAIAnalysisLink `json:"ai_analysis,omitempty"`
	AIAnalysisCardCount   int                `json:"ai_analysis_card_count"`
	NavigationActionCount int                `json:"navigation_action_count"`
	DisabledActionCount   int                `json:"disabled_action_count"`
	PrimaryCardID         string             `json:"primary_card_id"`
	CardIDs               []string           `json:"card_ids"`
	ActionQueueCreated    bool               `json:"action_queue_created"`
	ActionQueuePersisted  bool               `json:"action_queue_persisted"`
	DeckPreviewCreated    bool               `json:"deck_preview_created"`
	DeckPersisted         bool               `json:"deck_persisted"`
	CardsPersisted        bool               `json:"cards_persisted"`
	CardActionsEnabled    bool               `json:"card_actions_enabled"`
	RuntimeLaunchApproval bool               `json:"runtime_launch_approval"`
	LaunchAllowed         bool               `json:"launch_allowed"`
	ExecutionStarted      bool               `json:"execution_started"`
	BackendDetailsExposed bool               `json:"backend_details_exposed"`
}

type KDECenterPageActionGraph struct {
	RequestType              string                         `json:"request_type"`
	GraphType                string                         `json:"graph_type"`
	RuntimeMethod            string                         `json:"runtime_method"`
	ReadMethod               string                         `json:"read_method"`
	NodeCount                int                            `json:"node_count"`
	EdgeCount                int                            `json:"edge_count"`
	ActionNodeCount          int                            `json:"action_node_count"`
	EvidenceNodeCount        int                            `json:"evidence_node_count"`
	GateNodeCount            int                            `json:"gate_node_count"`
	MissingEvidenceCount     int                            `json:"missing_evidence_count"`
	BlockedActionCount       int                            `json:"blocked_action_count"`
	ReadOnlyCheckCount       int                            `json:"read_only_check_count"`
	MissingEvidenceIDs       []string                       `json:"missing_evidence_ids"`
	BlockedActions           []string                       `json:"blocked_actions"`
	ReceiptValidation        KDEActionReceiptValidation     `json:"receipt_validation"`
	NextReadOnlyChecks       []KDEActionDependencyNextCheck `json:"next_read_only_checks"`
	DependencyGraphCreated   bool                           `json:"dependency_graph_created"`
	DependencyGraphPersisted bool                           `json:"dependency_graph_persisted"`
	RequestObjectsCreated    bool                           `json:"request_objects_created"`
	PermissionGrantCreated   bool                           `json:"permission_grant_created"`
	SettingsPersisted        bool                           `json:"settings_persisted"`
	RuntimeLaunchApproval    bool                           `json:"runtime_launch_approval"`
	LaunchAllowed            bool                           `json:"launch_allowed"`
	ExecutionStarted         bool                           `json:"execution_started"`
	BackendProcessStarted    bool                           `json:"backend_process_started"`
	HostRootModified         bool                           `json:"host_root_modified"`
	BackendDetailsExposed    bool                           `json:"backend_details_exposed"`
}

type KDECenterPageSettings struct {
	RequestType                string            `json:"request_type"`
	SettingsState              string            `json:"settings_state"`
	SectionCount               int               `json:"section_count"`
	Sections                   []SettingsSection `json:"sections"`
	UserFacingSettings         map[string]string `json:"user_facing_settings"`
	SettingsPersisted          bool              `json:"settings_persisted"`
	SettingsPersistenceEnabled bool              `json:"settings_persistence_enabled"`
	HostRootModified           bool              `json:"host_root_modified"`
	BackendDetailsExposed      bool              `json:"backend_details_exposed"`
}

type KDECenterPageNavigation struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	Target         string `json:"target"`
	Enabled        bool   `json:"enabled"`
	NavigationOnly bool   `json:"navigation_only"`
	MutatesRuntime bool   `json:"mutates_runtime"`
	StartsProgram  bool   `json:"starts_program"`
}

type KDECenterPageSectionsPreview struct {
	SchemaVersion                string                         `json:"schema_version"`
	RequestType                  string                         `json:"request_type"`
	PageType                     string                         `json:"page_type"`
	Source                       string                         `json:"source"`
	Desktop                      string                         `json:"desktop"`
	RuntimeMethod                string                         `json:"runtime_method"`
	ReadMethod                   string                         `json:"read_method"`
	ApplicationID                string                         `json:"application_id"`
	ApplicationName              string                         `json:"application_name"`
	SectionCount                 int                            `json:"section_count"`
	ReadOnlySectionCount         int                            `json:"read_only_section_count"`
	NavigationOnlySectionCount   int                            `json:"navigation_only_section_count"`
	ExecutableSectionCount       int                            `json:"executable_section_count"`
	AIAnalysis                   *KDEAIAnalysisLink             `json:"ai_analysis,omitempty"`
	AIAnalysisSectionCount       int                            `json:"ai_analysis_section_count"`
	ApplicationReadinessEvidence KDECenterPageReadinessEvidence `json:"application_readiness_evidence"`
	PrimarySectionID             string                         `json:"primary_section_id"`
	Sections                     []KDECenterPageSection         `json:"sections"`
	RuntimeOwned                 bool                           `json:"runtime_owned"`
	GoRuntimeBacked              bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                           `json:"kde_policy_owner"`
	OfficialDesktopOnly          bool                           `json:"official_desktop_only"`
	UserVisible                  bool                           `json:"user_visible"`
	SafeForAIDiagnostics         bool                           `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured         bool                           `json:"user_decision_captured"`
	UserDecisionAllowsLaunch     bool                           `json:"user_decision_allows_launch"`
	SectionsPreviewCreated       bool                           `json:"sections_preview_created"`
	SectionsPersisted            bool                           `json:"sections_persisted"`
	SectionActionsEnabled        bool                           `json:"section_actions_enabled"`
	SettingsPersisted            bool                           `json:"settings_persisted"`
	SettingsPersistenceEnabled   bool                           `json:"settings_persistence_enabled"`
	NotificationsSent            bool                           `json:"notifications_sent"`
	ResourceGrantCreated         bool                           `json:"resource_grant_created"`
	RuntimeLaunchApproval        bool                           `json:"runtime_launch_approval"`
	LaunchEnabled                bool                           `json:"launch_enabled"`
	ExecutionStarted             bool                           `json:"execution_started"`
	RequestObjectsCreated        bool                           `json:"request_objects_created"`
	PermissionGrantCreated       bool                           `json:"permission_grant_created"`
	HostRootModified             bool                           `json:"host_root_modified"`
	NetworkRequired              bool                           `json:"network_required"`
	BackendDetailsExposed        bool                           `json:"backend_details_exposed"`
	BlockedActions               []string                       `json:"blocked_actions"`
	DesktopSafeSummary           string                         `json:"desktop_safe_summary"`
}

type KDECenterPageSection struct {
	ID                    string             `json:"id"`
	Label                 string             `json:"label"`
	Target                string             `json:"target"`
	RuntimeMethod         string             `json:"runtime_method"`
	ReadModel             string             `json:"read_model"`
	State                 string             `json:"state"`
	ReadinessNodeIDs      []string           `json:"readiness_node_ids"`
	ReadinessStatus       string             `json:"readiness_status"`
	ReadinessBlocked      bool               `json:"readiness_blocked"`
	ReadinessSummary      string             `json:"readiness_summary"`
	AIAnalysis            *KDEAIAnalysisLink `json:"ai_analysis,omitempty"`
	NavigationOnly        bool               `json:"navigation_only"`
	ReadOnly              bool               `json:"read_only"`
	MutatesRuntime        bool               `json:"mutates_runtime"`
	StartsProgram         bool               `json:"starts_program"`
	SettingsPersisted     bool               `json:"settings_persisted"`
	BackendDetailsExposed bool               `json:"backend_details_exposed"`
	Summary               string             `json:"summary"`
}

type KDECenterPageSectionDetailPreview struct {
	SchemaVersion                string                         `json:"schema_version"`
	RequestType                  string                         `json:"request_type"`
	PageType                     string                         `json:"page_type"`
	Source                       string                         `json:"source"`
	Desktop                      string                         `json:"desktop"`
	RuntimeMethod                string                         `json:"runtime_method"`
	ReadMethod                   string                         `json:"read_method"`
	ApplicationID                string                         `json:"application_id"`
	ApplicationName              string                         `json:"application_name"`
	SectionID                    string                         `json:"section_id"`
	SectionLabel                 string                         `json:"section_label"`
	SectionTarget                string                         `json:"section_target"`
	SectionState                 string                         `json:"section_state"`
	SectionRuntimeMethod         string                         `json:"section_runtime_method"`
	SectionReadModel             string                         `json:"section_read_model"`
	SectionSummary               string                         `json:"section_summary"`
	ReadinessNodeIDs             []string                       `json:"readiness_node_ids"`
	ReadinessStatus              string                         `json:"readiness_status"`
	ReadinessBlocked             bool                           `json:"readiness_blocked"`
	ReadinessSummary             string                         `json:"readiness_summary"`
	ApplicationReadinessEvidence KDECenterPageReadinessEvidence `json:"application_readiness_evidence"`
	DiagnosticHistoryRoute       *KDEDiagnosticHistoryRoute     `json:"diagnostic_history_route,omitempty"`
	AIAnalysis                   *KDEAIAnalysisLink             `json:"ai_analysis,omitempty"`
	AIAnalysisInput              *KDECenterPageAIAnalysisInput  `json:"ai_analysis_input,omitempty"`
	AvailableSectionIDs          []string                       `json:"available_section_ids"`
	ReadOnlyNavigation           bool                           `json:"read_only_navigation"`
	DetailPreviewCreated         bool                           `json:"detail_preview_created"`
	DetailPersisted              bool                           `json:"detail_persisted"`
	SectionActionsEnabled        bool                           `json:"section_actions_enabled"`
	SettingsPersisted            bool                           `json:"settings_persisted"`
	SettingsPersistenceEnabled   bool                           `json:"settings_persistence_enabled"`
	NotificationsSent            bool                           `json:"notifications_sent"`
	ResourceGrantCreated         bool                           `json:"resource_grant_created"`
	RuntimeLaunchApproval        bool                           `json:"runtime_launch_approval"`
	LaunchEnabled                bool                           `json:"launch_enabled"`
	ExecutionStarted             bool                           `json:"execution_started"`
	RequestObjectsCreated        bool                           `json:"request_objects_created"`
	PermissionGrantCreated       bool                           `json:"permission_grant_created"`
	HostRootModified             bool                           `json:"host_root_modified"`
	NetworkRequired              bool                           `json:"network_required"`
	BackendDetailsExposed        bool                           `json:"backend_details_exposed"`
	RuntimeOwned                 bool                           `json:"runtime_owned"`
	GoRuntimeBacked              bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                           `json:"kde_policy_owner"`
	OfficialDesktopOnly          bool                           `json:"official_desktop_only"`
	UserVisible                  bool                           `json:"user_visible"`
	SafeForAIDiagnostics         bool                           `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured         bool                           `json:"user_decision_captured"`
	UserDecisionAllowsLaunch     bool                           `json:"user_decision_allows_launch"`
	BlockedActions               []string                       `json:"blocked_actions"`
	DesktopSafeSummary           string                         `json:"desktop_safe_summary"`
}

type KDECenterPageReadinessEvidence struct {
	Source                     string            `json:"source"`
	GraphType                  string            `json:"graph_type"`
	RuntimeMethod              string            `json:"runtime_method"`
	ReadMethod                 string            `json:"read_method"`
	OverallStatus              string            `json:"overall_status"`
	NodeIDs                    []string          `json:"node_ids"`
	NodeStatuses               map[string]string `json:"node_statuses"`
	NodeCount                  int               `json:"node_count"`
	RequiredNodeCount          int               `json:"required_node_count"`
	BlockedNodeCount           int               `json:"blocked_node_count"`
	ReadyNodeCount             int               `json:"ready_node_count"`
	LaunchAllowed              bool              `json:"launch_allowed"`
	LaunchEnabled              bool              `json:"launch_enabled"`
	ExecutionRequestCreated    bool              `json:"execution_request_created"`
	ExecutionStarted           bool              `json:"execution_started"`
	BackendProcessStarted      bool              `json:"backend_process_started"`
	RealPortalTransportEnabled bool              `json:"real_portal_transport_enabled"`
	RequestObjectCreated       bool              `json:"request_object_created"`
	PermissionGranted          bool              `json:"permission_granted"`
	SnapshotCreated            bool              `json:"snapshot_created"`
	RestoreExecuted            bool              `json:"restore_executed"`
	HostRootModified           bool              `json:"host_root_modified"`
	NetworkRequired            bool              `json:"network_required"`
	BackendDetailsExposed      bool              `json:"backend_details_exposed"`
	Summary                    string            `json:"summary"`
}

type KDECenterPageAIAnalysisInput struct {
	RequestType            string `json:"request_type"`
	Source                 string `json:"source"`
	RuntimeMethod          string `json:"runtime_method"`
	AnalysisTask           string `json:"analysis_task"`
	AnalysisSurface        string `json:"analysis_surface"`
	SelectionMode          string `json:"selection_mode"`
	FileCount              int    `json:"file_count"`
	SelectedExtension      string `json:"selected_extension"`
	SelectedFileDisclosure string `json:"selected_file_disclosure"`
	UserReviewRequired     bool   `json:"user_review_required"`
	SafeForAIDiagnostics   bool   `json:"safe_for_ai_diagnostics"`
	AIProviderCallEnabled  bool   `json:"ai_provider_call_enabled"`
	NetworkRequired        bool   `json:"network_required"`
	FileContentRead        bool   `json:"file_content_read"`
	FilePathsExposed       bool   `json:"file_paths_exposed"`
	RequestObjectCreated   bool   `json:"request_object_created"`
	PermissionGranted      bool   `json:"permission_granted"`
	BackendLaunchEnabled   bool   `json:"backend_launch_enabled"`
	HostRootModified       bool   `json:"host_root_modified"`
	BackendDetailsExposed  bool   `json:"backend_details_exposed"`
}

type KDEDiagnosticHistoryRoute struct {
	RequestType                 string `json:"request_type"`
	Source                      string `json:"source"`
	RuntimeMethod               string `json:"runtime_method"`
	ReadMethod                  string `json:"read_method"`
	ReadModel                   string `json:"read_model"`
	CLICommand                  string `json:"cli_command"`
	ApplicationID               string `json:"application_id"`
	UserVisible                 bool   `json:"user_visible"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	StateRootRequired           bool   `json:"state_root_required"`
	StateRootPathExposed        bool   `json:"state_root_path_exposed"`
	HistoryPreviewCreated       bool   `json:"history_preview_created"`
	AIProviderCallEnabled       bool   `json:"ai_provider_call_enabled"`
	FileContentRead             bool   `json:"file_content_read"`
	FilePathsExposed            bool   `json:"file_paths_exposed"`
	RequestObjectCreated        bool   `json:"request_object_created"`
	PermissionGranted           bool   `json:"permission_granted"`
	LaunchEnabled               bool   `json:"launch_enabled"`
	ExecutionStarted            bool   `json:"execution_started"`
	RepairExecutionEnabled      bool   `json:"repair_execution_enabled"`
	HostRootModified            bool   `json:"host_root_modified"`
	NetworkRequired             bool   `json:"network_required"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
}

func NewKDECenterPagePreview(recipe Recipe, provenance Provenance, decision string, fileURIs []string) (KDECenterPagePreview, error) {
	return NewKDECenterPagePreviewWithOptions(recipe, provenance, decision, fileURIs, KDECenterPageOptions{})
}

func NewKDECenterPagePreviewWithOptions(recipe Recipe, provenance Provenance, decision string, fileURIs []string, options KDECenterPageOptions) (KDECenterPagePreview, error) {
	if !singleLine(decision) {
		return KDECenterPagePreview{}, errors.New("KDE center page preview requires a single-line decision")
	}

	center, err := NewCompatibilityCenterPreviewWithOptions([]Recipe{recipe}, provenance, CompatibilityCenterOptions{
		KnownAppSmokeEvidence: options.KnownAppSmokeEvidence,
	})
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	if len(center.Applications) != 1 {
		return KDECenterPagePreview{}, errors.New("KDE center page preview requires exactly one application summary")
	}

	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDECenterPagePreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDECenterPagePreview{}, errors.New("KDE center page preview requires single-line identity fields")
		}
	}

	deck, err := plan.KDEActionCardDeckPreview(decision, fileURIs)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	actionGraph, err := plan.KDEActionDependencyGraphPreview(decision, fileURIs)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	settings, err := plan.SettingsPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	backendSelection, err := plan.BackendSelectionPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	activationStatus, err := plan.DesktopActivationStatusPreviewWithReceipt(options.ActivationRoot, "development")
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	executionReadiness, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	applicationReadinessRoot := options.ApplicationReadinessRoot
	if applicationReadinessRoot == "" {
		applicationReadinessRoot = defaultApplicationReadinessRuntimeRoot()
	}
	applicationReadiness, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{
		Environment:     "development",
		RuntimeRoot:     applicationReadinessRoot,
		StateRoot:       options.ApplicationReadinessStateRoot,
		ArtifactReceipt: options.ApplicationReadinessArtifactReceipt,
		PortalOperation: options.ApplicationReadinessPortalOperation,
		SnapshotReason:  options.ApplicationReadinessSnapshotReason,
		WriteMethod:     "Launch",
	})
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	launchIntent, err := plan.LaunchIntentPreview(fileURIs)
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	windowIdentity, err := plan.WindowIdentityPreview()
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	taskManagerIdentity, err := plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{
		ActivationRoot:            options.ActivationRoot,
		ExecutionSessionRoot:      options.ExecutionSessionRoot,
		ExecutionSessionRequestID: options.ExecutionSessionRequestID,
	})
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	kwinRule, err := plan.KWinWindowRulePlanPreviewWithOptions(KWinWindowRuleOptions{
		ActivationRoot:            options.ActivationRoot,
		ExecutionSessionRoot:      options.ExecutionSessionRoot,
		ExecutionSessionRequestID: options.ExecutionSessionRequestID,
	})
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	fileAssociationReady := len(plan.MIMETypes) > 0
	if fileAssociationReady {
		if _, err := plan.RenderMIMEApps(); err != nil {
			return KDECenterPagePreview{}, err
		}
	}
	trayStatus, err := plan.TrayStatusPreviewWithOptions(TrayStatusOptions{
		ActivationRoot:            options.ActivationRoot,
		ExecutionSessionRoot:      options.ExecutionSessionRoot,
		ExecutionSessionRequestID: options.ExecutionSessionRequestID,
	})
	if err != nil {
		return KDECenterPagePreview{}, err
	}
	notification, err := plan.NotificationPreview("approval-required")
	if err != nil {
		return KDECenterPagePreview{}, err
	}

	application := center.Applications[0]
	navigation := kdeCenterPageNavigation()
	knownAppSessionGateCards := kdeCenterPageKnownAppSessionGateCards(center.KnownAppSmokeEvidence)
	knownAppMatrixCards := kdeCenterPageKnownAppMatrixCards(center.KnownAppSmokeEvidence)
	knownAppGUICards := kdeCenterPageKnownAppGUICards(center.KnownAppSmokeEvidence)
	knownAppOwnerControlledGUICount := countOwnerControlledKnownAppGUIEvidence(center.KnownAppSmokeEvidence)
	knownAppOwnerManagedCopyVerifiedCount := countOwnerManagedCopyVerifiedKnownAppGUIEvidence(center.KnownAppSmokeEvidence)
	source := "compatibility-center-preview+backend-selection-preview+desktop-activation-status-preview+execution-readiness-preview+application-readiness-preview+launch-intent-preview+window-identity-preview+file-association-plan+tray-status-preview+notification-preview+kde-action-card-deck-preview+kde-action-dependency-graph-preview+settings-preview"
	if options.ExecutionSessionRoot != "" {
		source += "+execution-session-record"
	}
	if len(knownAppSessionGateCards) > 0 {
		source += "+known-app-session-gate-evidence"
	}
	if len(knownAppMatrixCards) > 0 {
		source += "+known-app-matrix-evidence"
	}
	if len(knownAppGUICards) > 0 {
		source += "+known-app-gui-smoke-evidence"
	}
	if knownAppOwnerControlledGUICount > 0 {
		source += "+owner-controlled-gui-evidence"
	}
	preview := KDECenterPagePreview{
		SchemaVersion:   "xnix.runtime.kde_center_page.v1",
		RequestType:     "kde-center-page-preview",
		PageType:        "compatibility-center-application-page",
		Source:          source,
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDECenterPage",
		ReadMethod:      "GetKDECenterPagePreview",
		ApplicationID:   plan.ApplicationID,
		ApplicationName: plan.DisplayName,
		Icon:            plan.Icon,
		DesktopFile:     plan.DesktopFile,
		LauncherCommand: deck.LauncherCommand,
		Header: KDECenterPageHeader{
			Title:                 plan.DisplayName,
			Subtitle:              "Compatibility Center application details",
			Badge:                 kdeCenterPageBadge(deck),
			BadgeTone:             kdeCenterPageBadgeTone(deck),
			PrimaryActionLabel:    "Review required gates",
			PrimaryActionTarget:   "compatibility-center-gates",
			PrimaryActionEnabled:  true,
			BackendDetailsExposed: false,
		},
		ApplicationSummary: KDECenterPageApplication{
			ApplicationID:              application.ApplicationID,
			DisplayName:                application.DisplayName,
			CompatibilityState:         application.CompatibilityState,
			CompatibilityLabel:         application.CompatibilityLabel,
			DiagnosticsState:           application.DiagnosticsState,
			RuntimeMode:                application.RuntimeMode,
			SupportedExtensions:        application.SupportedExtensions,
			KnownIssueCount:            application.KnownIssueCount,
			RepairRecordState:          application.RepairRecordState,
			RepairRecordCount:          application.RepairRecordCount,
			ActionExecutionEnabled:     application.ActionExecutionEnabled,
			RepairExecutionEnabled:     application.RepairExecutionEnabled,
			BackendLaunchEnabled:       application.BackendLaunchEnabled,
			SettingsPersistenceEnabled: application.SettingsPersistenceEnabled,
			HostRootModified:           application.HostRootModified,
			BackendDetailsExposed:      application.BackendDetailsExposed,
			Summary:                    application.Summary,
		},
		KnownAppSessionGateEvidenceCount:         len(knownAppSessionGateCards),
		KnownAppLauncherSessionGateConsumedCount: center.KnownAppLauncherSessionGateConsumedCount,
		KnownAppPostReviewDispatchConsumedCount:  center.KnownAppPostReviewDispatchConsumedCount,
		KnownAppSessionGateCards:                 knownAppSessionGateCards,
		KnownAppMatrixEvidenceCount:              len(knownAppMatrixCards),
		KnownAppMatrixEvidenceCards:              knownAppMatrixCards,
		KnownAppGUIEvidenceCount:                 len(knownAppGUICards),
		KnownAppOwnerControlledGUIEvidenceCount:  knownAppOwnerControlledGUICount,
		KnownAppOwnerManagedCopyVerifiedCount:    knownAppOwnerManagedCopyVerifiedCount,
		KnownAppGUIEvidenceCards:                 knownAppGUICards,
		BackendSelectionSnapshot: KDECenterPageBackend{
			RequestType:                 backendSelection.RequestType,
			PlanType:                    backendSelection.PlanType,
			RuntimeMethod:               backendSelection.RuntimeMethod,
			SelectedStrategy:            backendSelection.SelectedStrategy,
			RecommendedProfileID:        backendSelection.RecommendedProfileID,
			CandidateCount:              backendSelection.CandidateCount,
			ReadyCandidateCount:         backendSelection.ReadyCandidateCount,
			BlockedCandidateCount:       backendSelection.BlockedCandidateCount,
			SelectionCommitted:          backendSelection.SelectionCommitted,
			SelectionChangeEnabled:      backendSelection.SelectionChangeEnabled,
			BackendLaunchEnabled:        backendSelection.BackendLaunchEnabled,
			CapabilityActivationEnabled: backendSelection.CapabilityActivationEnabled,
			EnvironmentCreated:          backendSelection.EnvironmentCreated,
			HostRootModified:            backendSelection.HostRootModified,
			BackendDetailsExposed:       backendSelection.BackendDetailsExposed,
			Summary:                     backendSelection.DesktopSafeSummary,
		},
		ActivationStatusSnapshot: KDECenterPageActivation{
			RequestType:           activationStatus.RequestType,
			StatusType:            activationStatus.StatusType,
			RuntimeMethod:         activationStatus.PlannedRuntimeMethod,
			Renderer:              activationStatus.CurrentRenderer,
			ActivationState:       activationStatus.ActivationState,
			TransactionState:      activationStatus.Transaction.TransactionState,
			PreflightDecision:     activationStatus.PreflightDecision,
			ReceiptEvidenceState:  activationStatus.ReceiptEvidence.EvidenceState,
			ReceiptRelativePath:   activationStatus.ReceiptEvidence.ReceiptRelativePath,
			ReceiptBacked:         activationStatus.ReceiptBacked,
			RollbackAvailable:     activationStatus.RollbackAvailable,
			StatusSignalCount:     len(activationStatus.StatusSignals),
			BlockedReasonCount:    len(activationStatus.BlockedReasons),
			NextSafeActionCount:   len(activationStatus.NextSafeActions),
			ActivationReady:       activationStatus.ActivationReady,
			ActivationCommitted:   activationStatus.ActivationCommitted,
			CommitEnabled:         activationStatus.CommitGate.CommitEnabled,
			LaunchEnabled:         activationStatus.LaunchEnabled,
			HostRootModified:      activationStatus.HostRootModified,
			BackendDetailsExposed: activationStatus.BackendDetailsExposed,
			Summary:               activationStatus.DesktopSafeSummary,
		},
		ExecutionReadinessSnapshot: KDECenterPageExecution{
			RequestType:               executionReadiness.RequestType,
			ReadinessType:             executionReadiness.ReadinessType,
			RuntimeMethod:             executionReadiness.RuntimeMethod,
			ExecutionState:            executionReadiness.ExecutionState,
			OverallStatus:             executionReadiness.OverallStatus,
			RecommendedAction:         executionReadiness.RecommendedAction,
			GateCount:                 executionReadiness.GateCount,
			RequiredGateCount:         executionReadiness.RequiredGateCount,
			PendingGateCount:          executionReadiness.PendingGateCount,
			BlockedGateCount:          executionReadiness.BlockedGateCount,
			DesktopEntryLaunchVisible: executionReadiness.DesktopEntryLaunchVisible,
			LaunchAllowed:             executionReadiness.LaunchAllowed,
			LaunchEnabled:             executionReadiness.LaunchEnabled,
			ExecutionRequestCreated:   executionReadiness.ExecutionRequestCreated,
			BackendBindingReady:       executionReadiness.BackendBindingReady,
			PortalPolicyRequired:      executionReadiness.PortalPolicyRequired,
			SnapshotRequired:          executionReadiness.SnapshotRequired,
			UserActionRequired:        executionReadiness.UserActionRequired,
			HostRootModified:          executionReadiness.HostRootModified,
			NetworkRequired:           executionReadiness.NetworkRequired,
			BackendDetailsExposed:     executionReadiness.BackendDetailsExposed,
			Summary:                   executionReadiness.DesktopSafeSummary,
		},
		ApplicationReadinessEvidence: applicationReadiness,
		LaunchIntentSnapshot: KDECenterPageLaunchIntent{
			RequestType:             launchIntent.RequestType,
			IntentType:              launchIntent.IntentType,
			Source:                  launchIntent.Source,
			RuntimeMethod:           launchIntent.RuntimeMethod,
			ReadMethod:              launchIntent.ReadMethod,
			FileCount:               launchIntent.FileCount,
			PortalRequired:          launchIntent.PortalRequired,
			SnapshotRequired:        launchIntent.SnapshotRequired,
			StandardDesktopEntry:    launchIntent.StandardDesktopEntry,
			LaunchUsesRuntime:       launchIntent.LaunchUsesRuntime,
			LaunchAllowed:           launchIntent.LaunchAllowed,
			LaunchEnabled:           launchIntent.LaunchEnabled,
			ExecutionRequestCreated: launchIntent.ExecutionRequestCreated,
			ExecutionStarted:        launchIntent.ExecutionStarted,
			BackendBindingReady:     launchIntent.BackendBindingReady,
			RequestObjectCreated:    launchIntent.RequestObjectCreated,
			PermissionGranted:       launchIntent.PermissionGranted,
			HostRootModified:        launchIntent.HostRootModified,
			NetworkRequired:         launchIntent.NetworkRequired,
			BackendDetailsExposed:   launchIntent.BackendDetailsExposed,
			Summary:                 launchIntent.DesktopSafeSummary,
		},
		WindowIdentitySnapshot: KDECenterPageWindow{
			SchemaVersion:             windowIdentity.SchemaVersion,
			DesktopFile:               windowIdentity.DesktopFile,
			LauncherURL:               windowIdentity.LauncherURL,
			WindowKind:                windowIdentity.WindowKind,
			ClassGroup:                windowIdentity.ClassGroup,
			ResourceName:              windowIdentity.ResourceName,
			TitleHint:                 windowIdentity.TitleHint,
			TaskManagerGroupingKey:    windowIdentity.TaskManager.GroupingKey,
			TaskManagerPinningAllowed: windowIdentity.TaskManager.PinningAllowed,
			TaskManagerRestoreAllowed: windowIdentity.TaskManager.RestoreAllowed,
			TaskManagerSkipTaskbar:    windowIdentity.TaskManager.SkipTaskbar,
			TaskManagerShowInSwitcher: windowIdentity.TaskManager.ShowInSwitcher,
			PreferExistingWindow:      windowIdentity.TaskManager.PreferExistingWindow,
			KWinScriptRole:            windowIdentity.KWin.ScriptRole,
			KWinPlacement:             windowIdentity.KWin.Placement,
			WindowManagerPolicyOnly:   windowIdentity.KWin.WindowManagerPolicyOnly,
			RuntimeOwnsBackendPolicy:  windowIdentity.KWin.RuntimeOwnsBackendPolicy,
			ExecutionSessionRoot:      taskManagerIdentity.ExecutionSessionRoot || kwinRule.ExecutionSessionRoot,
			ExecutionSessionBacked:    taskManagerIdentity.ExecutionSessionBacked && kwinRule.ExecutionSessionBacked,
			ExecutionSessionPath:      taskManagerIdentity.ExecutionSessionPath,
			TaskManagerSessionState:   taskManagerIdentity.ExecutionSessionState,
			KWinSessionState:          kwinRule.ExecutionSessionState,
			TaskManagerEntryActive:    taskManagerIdentity.TaskManagerEntryActive,
			KWinRuleApplied:           kwinRule.KWinRuleApplied,
			HostRootModified:          windowIdentity.HostRootModified,
			BackendDetailsExposed:     windowIdentity.BackendDetailsExposed,
			Summary:                   windowIdentity.Summary,
		},
		FileAssociationSnapshot: KDECenterPageFiles{
			PlanType:                  "file-association-plan",
			AssociationType:           "desktop-file-association",
			RuntimeMethod:             "GetFileAssociationPlan",
			DesktopFile:               plan.DesktopFile,
			MIMEAppsPath:              "usr/share/applications/mimeapps.list",
			MIMETypes:                 plan.MIMETypes,
			MIMETypeCount:             len(plan.MIMETypes),
			FileOpenCommand:           "xnix-compat-open",
			FileOpenArgument:          "%U",
			StandardMIMEAppsList:      true,
			StagedRootOnly:            true,
			OverwriteExistingMIMEApps: false,
			PortalRequiredForFileOpen: fileAssociationReady,
			FileAssociationReady:      fileAssociationReady,
			FileOpenPreviewAvailable:  fileAssociationReady,
			DirectHostFileAccess:      false,
			RequestObjectCreated:      false,
			PermissionGranted:         false,
			FilesWritten:              false,
			MIMEAppsWritten:           false,
			HostRootModified:          false,
			BackendDetailsExposed:     false,
			Summary:                   kdeCenterPageFileAssociationSummary(fileAssociationReady),
		},
		TrayStatusSnapshot: KDECenterPageTray{
			StatusType:                   trayStatus.StatusType,
			RuntimeMethod:                "GetTrayStatus",
			DesktopFile:                  trayStatus.DesktopFile,
			RegisteredApplicationCount:   trayStatus.RuntimeActivity.RegisteredApplicationCount,
			ActiveApplicationCount:       trayStatus.RuntimeActivity.ActiveApplicationCount,
			AttentionRequiredCount:       trayStatus.RuntimeActivity.AttentionRequiredCount,
			CompatibilityState:           trayStatus.CompatibilityStatus.State,
			CompatibilityLabel:           trayStatus.CompatibilityStatus.Label,
			TrayBridgeState:              trayStatus.TrayBridge.State,
			TrayBridgeLabel:              trayStatus.TrayBridge.Label,
			BridgedTrayApplicationCount:  trayStatus.TrayBridge.BridgedTrayApplicationCount,
			Actions:                      trayStatus.Actions,
			ActionCount:                  len(trayStatus.Actions),
			UserVisible:                  trayStatus.UserVisible,
			TrayStatusReady:              true,
			ExecutionSessionRoot:         trayStatus.ExecutionSessionRoot,
			ExecutionSessionBacked:       trayStatus.ExecutionSessionBacked,
			ExecutionSessionPath:         trayStatus.ExecutionSessionPath,
			ExecutionSessionState:        trayStatus.ExecutionSessionState,
			LiveBackendBridgeEnabled:     trayStatus.LiveBackendBridgeEnabled,
			BridgeConfigurationPersisted: trayStatus.BridgeConfigurationPersisted,
			HostRootModified:             trayStatus.HostRootModified,
			BackendDetailsExposed:        trayStatus.BackendDetailsExposed,
			Summary:                      trayStatus.Summary,
		},
		NotificationSnapshot: KDECenterPageNotification{
			RequestType:                notification.RequestType,
			RuntimeMethod:              "GetNotificationPlan",
			Source:                     notification.Source,
			DesktopFile:                notification.DesktopFile,
			EventType:                  notification.EventType,
			NotificationID:             notification.NotificationID,
			Urgency:                    notification.Urgency,
			Category:                   notification.Category,
			Title:                      notification.Title,
			ActionCount:                len(notification.Actions),
			Actions:                    notification.Actions,
			RequiresUserReview:         notification.RequiresUserReview,
			UserVisible:                notification.UserVisible,
			NotificationReady:          true,
			ActionExecutionEnabled:     notification.ActionExecutionEnabled,
			RepairExecutionEnabled:     notification.RepairExecutionEnabled,
			SettingsPersistenceEnabled: notification.SettingsPersistenceEnabled,
			NotificationsSent:          false,
			HostRootModified:           notification.HostRootModified,
			BackendDetailsExposed:      notification.BackendDetailsExposed,
			Summary:                    notification.Summary,
		},
		ActionDeck: KDECenterPageActionDeck{
			RequestType:           deck.RequestType,
			DeckType:              deck.DeckType,
			CardCount:             deck.CardCount,
			WaitingCardCount:      deck.WaitingCardCount,
			DeferredCardCount:     deck.DeferredCardCount,
			RejectedCardCount:     deck.RejectedCardCount,
			AIAnalysis:            deck.AIAnalysis,
			AIAnalysisCardCount:   deck.AIAnalysisCardCount,
			NavigationActionCount: deck.NavigationActionCount,
			DisabledActionCount:   deck.DisabledActionCount,
			PrimaryCardID:         deck.PrimaryCardID,
			CardIDs:               deck.CardIDs,
			ActionQueueCreated:    deck.ActionQueueCreated,
			ActionQueuePersisted:  deck.ActionQueuePersisted,
			DeckPreviewCreated:    deck.DeckPreviewCreated,
			DeckPersisted:         deck.DeckPersisted,
			CardsPersisted:        deck.CardsPersisted,
			CardActionsEnabled:    deck.CardActionsEnabled,
			RuntimeLaunchApproval: deck.RuntimeLaunchApproval,
			LaunchAllowed:         deck.LaunchAllowed,
			ExecutionStarted:      deck.ExecutionStarted,
			BackendDetailsExposed: deck.BackendDetailsExposed,
		},
		ActionDependencyGraph: KDECenterPageActionGraph{
			RequestType:              actionGraph.RequestType,
			GraphType:                actionGraph.GraphType,
			RuntimeMethod:            actionGraph.RuntimeMethod,
			ReadMethod:               actionGraph.ReadMethod,
			NodeCount:                actionGraph.NodeCount,
			EdgeCount:                actionGraph.EdgeCount,
			ActionNodeCount:          actionGraph.ActionNodeCount,
			EvidenceNodeCount:        actionGraph.EvidenceNodeCount,
			GateNodeCount:            actionGraph.GateNodeCount,
			MissingEvidenceCount:     actionGraph.MissingEvidenceCount,
			BlockedActionCount:       actionGraph.BlockedActionCount,
			ReadOnlyCheckCount:       actionGraph.ReadOnlyCheckCount,
			MissingEvidenceIDs:       actionGraph.MissingEvidenceIDs,
			BlockedActions:           actionGraph.BlockedActions,
			ReceiptValidation:        actionGraph.ReceiptValidation,
			NextReadOnlyChecks:       actionGraph.NextReadOnlyChecks,
			DependencyGraphCreated:   actionGraph.DependencyGraphCreated,
			DependencyGraphPersisted: actionGraph.DependencyGraphPersisted,
			RequestObjectsCreated:    actionGraph.RequestObjectsCreated,
			PermissionGrantCreated:   actionGraph.PermissionGrantCreated,
			SettingsPersisted:        actionGraph.SettingsPersisted,
			RuntimeLaunchApproval:    actionGraph.RuntimeLaunchApproval,
			LaunchAllowed:            actionGraph.LaunchAllowed,
			ExecutionStarted:         actionGraph.ExecutionStarted,
			BackendProcessStarted:    actionGraph.BackendProcessStarted,
			HostRootModified:         actionGraph.HostRootModified,
			BackendDetailsExposed:    actionGraph.BackendDetailsExposed,
		},
		SettingsSnapshot: KDECenterPageSettings{
			RequestType:                settings.RequestType,
			SettingsState:              settings.SettingsState,
			SectionCount:               settings.SectionCount,
			Sections:                   settings.Sections,
			UserFacingSettings:         settings.UserFacingSettings,
			SettingsPersisted:          settings.SettingsPersisted,
			SettingsPersistenceEnabled: settings.SettingsPersistenceEnabled,
			HostRootModified:           settings.HostRootModified,
			BackendDetailsExposed:      settings.BackendDetailsExposed,
		},
		Navigation:                 navigation,
		NavigationCount:            len(navigation),
		PrimaryNavigationTarget:    "compatibility-center-gates",
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		OfficialDesktopOnly:        true,
		UserVisible:                true,
		SafeForAIDiagnostics:       true,
		UserDecisionCaptured:       deck.UserDecisionCaptured,
		UserDecisionAllowsLaunch:   deck.UserDecisionAllowsLaunch,
		PagePreviewCreated:         true,
		PagePersisted:              false,
		DeckPersisted:              false,
		CardsPersisted:             false,
		CardActionsEnabled:         false,
		SettingsPersisted:          false,
		SettingsPersistenceEnabled: false,
		NotificationsSent:          false,
		ResourceGrantCreated:       false,
		RuntimeLaunchApproval:      false,
		LaunchAllowed:              false,
		LaunchEnabled:              false,
		ExecutionStarted:           false,
		BackendProcessStarted:      false,
		RequestObjectsCreated:      false,
		PermissionGrantCreated:     false,
		HostRootModified:           false,
		NetworkRequired:            false,
		BackendDetailsExposed:      false,
		BlockedActions:             []string{"persist KDE center page from preview state", "enable center page action buttons from preview state", "persist KDE action card deck from center page preview", "persist KDE action dependency graph from center page preview", "persist compatibility settings from center page preview", "record review receipts from center page preview", "create Runtime request objects from center page preview", "grant desktop resources from center page preview", "send desktop notifications from center page preview", "enable live tray bridge from center page preview", "start compatibility profile from center page preview", "create application readiness request objects from center page preview", "mutate host root during KDE center page preview", "expose raw backend command to desktop shell"},
		UserFacingSettings:         settings.UserFacingSettings,
		DesktopSafeSummary:         "KDE can render an application page from Runtime-owned summary, application readiness evidence, execution readiness, launch intent, window identity, file association, tray status, notification, action deck, action dependency graph, and settings previews, but the page remains read-only and cannot approve, persist, grant, bridge, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page preview"); err != nil {
		return KDECenterPagePreview{}, err
	}
	return preview, nil
}

func kdeCenterPageKnownAppSessionGateCards(evidence []KnownAppSmokeEvidenceSummary) []KDECenterPageKnownAppSessionGateCard {
	cards := make([]KDECenterPageKnownAppSessionGateCard, 0, len(evidence))
	for _, item := range evidence {
		if !item.LauncherSessionGateConsumed {
			continue
		}
		runtimeStatusArgv := kdeCenterPageRuntimeStatusLaunchArgv(item)
		cards = append(cards, KDECenterPageKnownAppSessionGateCard{
			AppID:                                      item.AppID,
			DisplayName:                                item.DisplayName,
			AppVersion:                                 item.AppVersion,
			CompatibilityState:                         item.CompatibilityState,
			CenterCardState:                            item.CenterCardState,
			ControlledExecutionSessionID:               item.ControlledExecutionSessionID,
			LaunchAuthorizationReceiptID:               item.LaunchAuthorizationReceiptID,
			LauncherSessionGateConsumed:                item.LauncherSessionGateConsumed,
			LauncherSessionDigestVerified:              item.LauncherSessionDigestVerified,
			LauncherSessionRelativePath:                item.LauncherSessionRelativePath,
			RuntimeOwnerConsumableSession:              item.LauncherSessionRuntimeOwnerConsumable,
			KDEReadModelConsumableSession:              item.LauncherSessionKDEReadModelConsumable,
			PostReviewDispatchConsumed:                 item.PostReviewDispatchConsumed,
			PostReviewDispatchState:                    item.PostReviewDispatchState,
			SessionGatedReviewReceiptID:                item.SessionGatedReviewReceiptID,
			LaunchGateConsumed:                         item.LaunchGateConsumed,
			ControlledDispatchReady:                    item.ControlledDispatchReady,
			PrimaryActionID:                            item.PrimaryActionID,
			PrimaryActionLabel:                         item.PrimaryActionLabel,
			PrimaryActionKind:                          item.PrimaryActionKind,
			PrimaryActionEnabled:                       item.PrimaryActionEnabled,
			RuntimeStatusLaunchRequestType:             KnownAppKDERuntimeStatusLaunchRequestType,
			RuntimeStatusLaunchRuntimeMethod:           "PreviewKnownAppKDERuntimeStatusLaunchRequest",
			RuntimeStatusLaunchReadMethod:              "GetKnownAppKDERuntimeStatusLaunchRequest",
			RuntimeStatusLaunchRequiredIDCount:         3,
			RuntimeStatusLaunchCollectedIDCount:        kdeCenterPageRuntimeStatusCollectedIDCount(item),
			RuntimeStatusLaunchManagedLauncherArgv:     runtimeStatusArgv,
			RuntimeStatusLaunchRequestReady:            item.PostReviewDispatchConsumed && len(runtimeStatusArgv) > 0,
			RuntimeStatusLaunchStateRootRequired:       true,
			RuntimeStatusLaunchStateRootOwnedByRuntime: true,
			ReviewRouteRequestType:                     KnownAppSessionGatedLaunchReviewRequestType,
			ReviewRouteRuntimeMethod:                   "PreviewKnownAppSessionGatedLaunchReview",
			ReviewRouteReadMethod:                      "GetKnownAppSessionGatedLaunchReview",
			ReadBeforeWriteRequired:                    true,
			RuntimeReceiptRequired:                     true,
			UserVisible:                                true,
			RuntimeOwned:                               true,
			GoRuntimeBacked:                            true,
			KDEPolicyOwner:                             false,
			DesktopLaunchEnabled:                       false,
			BackendLaunchEnabled:                       false,
			HostRootModified:                           false,
			BackendDetailsExposed:                      false,
			Summary:                                    item.Summary,
		})
	}
	return cards
}

func kdeCenterPageKnownAppMatrixCards(evidence []KnownAppSmokeEvidenceSummary) []KDECenterPageKnownAppMatrixCard {
	cards := make([]KDECenterPageKnownAppMatrixCard, 0, len(evidence))
	for _, item := range evidence {
		if item.EvidenceSource != "remote-known-winapp-matrix-smoke" {
			continue
		}
		cards = append(cards, KDECenterPageKnownAppMatrixCard{
			AppID:                                item.AppID,
			DisplayName:                          item.DisplayName,
			AppVersion:                           item.AppVersion,
			EvidenceKind:                         item.EvidenceKind,
			EvidenceSource:                       item.EvidenceSource,
			SmokeStatus:                          item.SmokeStatus,
			CompatibilityState:                   item.CompatibilityState,
			CenterCardState:                      item.CenterCardState,
			PrimaryActionID:                      item.PrimaryActionID,
			PrimaryActionLabel:                   item.PrimaryActionLabel,
			PrimaryActionKind:                    item.PrimaryActionKind,
			PrimaryActionEnabled:                 item.PrimaryActionEnabled,
			MarkerObserved:                       item.MarkerObserved,
			ChecksumVerified:                     item.ChecksumVerified,
			ExecutionEvidenceRecorded:            item.ExecutionEvidenceRecorded,
			StagedLauncherVerified:               item.StagedLauncherVerified,
			OwnerControlledRuntimeLaunchVerified: item.OwnerControlledRuntimeLaunchVerified,
			OwnerManagedCopyVerified:             item.OwnerManagedCopyVerified,
			OwnerServiceCallReady:                item.OwnerServiceCallReady,
			OwnerEvidenceHandoffReady:            item.OwnerEvidenceHandoffReady,
			OwnerEvidenceRelativePath:            item.OwnerEvidenceRelativePath,
			RuntimeDispatchVerified:              item.RuntimeDispatchVerified,
			LaunchAuthorizationRequired:          item.LaunchAuthorizationRequired,
			DesktopLaunchEnabled:                 false,
			BackendLaunchEnabled:                 false,
			RuntimeOwned:                         true,
			GoRuntimeBacked:                      true,
			KDEPolicyOwner:                       false,
			HostRootModified:                     false,
			BackendDetailsExposed:                false,
			RawArtifactPathExposed:               false,
			Summary:                              item.Summary,
		})
	}
	return cards
}

func kdeCenterPageKnownAppGUICards(evidence []KnownAppSmokeEvidenceSummary) []KDECenterPageKnownAppMatrixCard {
	cards := make([]KDECenterPageKnownAppMatrixCard, 0, len(evidence))
	for _, item := range evidence {
		if item.EvidenceSource != "wine-guest-gui-smoke" {
			continue
		}
		routeReady := kdeCenterPageOwnerGUIRouteReady(item)
		cards = append(cards, KDECenterPageKnownAppMatrixCard{
			AppID:                                item.AppID,
			DisplayName:                          item.DisplayName,
			AppVersion:                           item.AppVersion,
			EvidenceKind:                         item.EvidenceKind,
			EvidenceSource:                       item.EvidenceSource,
			SmokeStatus:                          item.SmokeStatus,
			CompatibilityState:                   item.CompatibilityState,
			CenterCardState:                      item.CenterCardState,
			PrimaryActionID:                      item.PrimaryActionID,
			PrimaryActionLabel:                   item.PrimaryActionLabel,
			PrimaryActionKind:                    item.PrimaryActionKind,
			PrimaryActionEnabled:                 item.PrimaryActionEnabled,
			DesktopCallableRoute:                 kdeCenterPageOwnerGUIDesktopRoute(routeReady),
			DesktopCallableRuntimeMethod:         kdeCenterPageOwnerGUIRuntimeMethod(routeReady),
			DesktopCallableExecutionType:         kdeCenterPageOwnerGUIExecutionType(routeReady),
			DesktopDBusMethod:                    kdeCenterPageOwnerGUIDBusMethod(routeReady),
			DesktopEvidenceHandleForwarded:       routeReady,
			KDEForwardedArgumentKind:             kdeCenterPageOwnerGUIForwardedArgumentKind(routeReady),
			KDEForwardedArguments:                kdeCenterPageOwnerGUIForwardedArguments(item, routeReady),
			OwnerServiceArgsExposedToKDE:         false,
			MarkerObserved:                       item.MarkerObserved,
			ChecksumVerified:                     item.ChecksumVerified,
			ExecutionEvidenceRecorded:            item.ExecutionEvidenceRecorded,
			StagedLauncherVerified:               item.StagedLauncherVerified,
			OwnerControlledRuntimeLaunchVerified: item.OwnerControlledRuntimeLaunchVerified,
			OwnerManagedCopyVerified:             item.OwnerManagedCopyVerified,
			OwnerServiceCallReady:                item.OwnerServiceCallReady,
			OwnerEvidenceHandoffReady:            item.OwnerEvidenceHandoffReady,
			OwnerEvidenceRelativePath:            item.OwnerEvidenceRelativePath,
			RuntimeDispatchVerified:              item.RuntimeDispatchVerified,
			LaunchAuthorizationRequired:          item.LaunchAuthorizationRequired,
			DesktopLaunchEnabled:                 false,
			BackendLaunchEnabled:                 false,
			RuntimeOwned:                         true,
			GoRuntimeBacked:                      true,
			KDEPolicyOwner:                       false,
			HostRootModified:                     false,
			BackendDetailsExposed:                false,
			RawArtifactPathExposed:               false,
			Summary:                              item.Summary,
		})
	}
	return cards
}

func kdeCenterPageOwnerGUIRouteReady(item KnownAppSmokeEvidenceSummary) bool {
	return item.OwnerControlledRuntimeLaunchVerified &&
		item.OwnerServiceCallReady &&
		item.OwnerEvidenceHandoffReady &&
		safeKnownAppOwnerEvidenceRelativePath(item.OwnerEvidenceRelativePath)
}

func kdeCenterPageOwnerGUIDesktopRoute(ready bool) string {
	if !ready {
		return ""
	}
	return "kde-dbus-runtime-status-action"
}

func kdeCenterPageOwnerGUIRuntimeMethod(ready bool) string {
	if !ready {
		return ""
	}
	return "ShowRuntimeControlledLaunch"
}

func kdeCenterPageOwnerGUIExecutionType(ready bool) string {
	if !ready {
		return ""
	}
	return KnownAppKDERuntimeStatusLaunchExecutionRequestType
}

func kdeCenterPageOwnerGUIDBusMethod(ready bool) string {
	if !ready {
		return ""
	}
	return "org.xnix.Compatibility1.ShowRuntimeControlledLaunch"
}

func kdeCenterPageOwnerGUIForwardedArgumentKind(ready bool) string {
	if !ready {
		return ""
	}
	return "evidence-relative-path"
}

func kdeCenterPageOwnerGUIForwardedArguments(item KnownAppSmokeEvidenceSummary, ready bool) []string {
	if !ready {
		return nil
	}
	return []string{strings.TrimSpace(item.OwnerEvidenceRelativePath)}
}

func countOwnerControlledKnownAppGUIEvidence(evidence []KnownAppSmokeEvidenceSummary) int {
	count := 0
	for _, item := range evidence {
		if item.EvidenceSource == "wine-guest-gui-smoke" &&
			item.SmokeStatus == "passed" &&
			item.ExecutionEvidenceRecorded &&
			item.RuntimeDispatchVerified &&
			item.OwnerControlledRuntimeLaunchVerified &&
			item.CompatibilityState == "owner-controlled-gui-qemu-wine-verified" &&
			item.CenterCardState == "validated-owner-controlled-gui-runtime-run" {
			count++
		}
	}
	return count
}

func countOwnerManagedCopyVerifiedKnownAppGUIEvidence(evidence []KnownAppSmokeEvidenceSummary) int {
	count := 0
	for _, item := range evidence {
		if item.EvidenceSource == "wine-guest-gui-smoke" &&
			item.SmokeStatus == "passed" &&
			item.OwnerControlledRuntimeLaunchVerified &&
			item.OwnerManagedCopyVerified {
			count++
		}
	}
	return count
}

func kdeCenterPageRuntimeStatusLaunchArgv(item KnownAppSmokeEvidenceSummary) []string {
	if !item.PostReviewDispatchConsumed ||
		strings.TrimSpace(item.LaunchAuthorizationReceiptID) == "" ||
		strings.TrimSpace(item.SessionGatedReviewReceiptID) == "" ||
		strings.TrimSpace(item.ControlledExecutionSessionID) == "" {
		return nil
	}
	return []string{
		"xnix-compat-launch",
		"--app", item.AppID,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--receipt-id", item.LaunchAuthorizationReceiptID,
		"--review-receipt-id", item.SessionGatedReviewReceiptID,
		"--session-id", item.ControlledExecutionSessionID,
	}
}

func kdeCenterPageRuntimeStatusCollectedIDCount(item KnownAppSmokeEvidenceSummary) int {
	count := 0
	for _, value := range []string{item.LaunchAuthorizationReceiptID, item.SessionGatedReviewReceiptID, item.ControlledExecutionSessionID} {
		if strings.TrimSpace(value) != "" {
			count++
		}
	}
	return count
}

func kdeCenterPageAIAnalysisInput(recipe Recipe, provenance Provenance, sectionID string, fileURIs []string) (*KDECenterPageAIAnalysisInput, error) {
	if sectionID != "diagnostics" || len(fileURIs) == 0 {
		return nil, nil
	}
	preview, err := NewDolphinAIAnalysisPreview([]Recipe{recipe}, provenance, fileURIs, recipe.ID)
	if err != nil {
		return nil, err
	}
	input := KDECenterPageAIAnalysisInput{
		RequestType:            preview.RequestType,
		Source:                 preview.Source,
		RuntimeMethod:          preview.RuntimeMethod,
		AnalysisTask:           preview.AnalysisTask,
		AnalysisSurface:        preview.AnalysisSurface,
		SelectionMode:          preview.SelectionMode,
		FileCount:              preview.FileCount,
		SelectedExtension:      preview.SelectedExtension,
		SelectedFileDisclosure: preview.SelectedFileDisclosure,
		UserReviewRequired:     preview.UserReviewRequired,
		SafeForAIDiagnostics:   preview.SafeForAIDiagnostics,
		AIProviderCallEnabled:  preview.AIProviderCallEnabled,
		NetworkRequired:        preview.NetworkRequired,
		FileContentRead:        preview.FileContentRead,
		FilePathsExposed:       preview.FilePathsExposed,
		RequestObjectCreated:   preview.RequestObjectCreated,
		PermissionGranted:      preview.PermissionGranted,
		BackendLaunchEnabled:   preview.BackendLaunchEnabled,
		HostRootModified:       preview.HostRootModified,
		BackendDetailsExposed:  preview.BackendDetailsExposed,
	}
	return &input, nil
}

func defaultApplicationReadinessRuntimeRoot() string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "."
	}
	for {
		if _, err := os.Stat(filepath.Join(workingDirectory, "VERSION")); err == nil {
			return workingDirectory
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory {
			return "."
		}
		workingDirectory = parent
	}
}

func kdeDiagnosticHistoryRoute(sectionID string, applicationID string) *KDEDiagnosticHistoryRoute {
	if sectionID != "diagnostics" {
		return nil
	}
	return &KDEDiagnosticHistoryRoute{
		RequestType:                 "diagnostic-history-route",
		Source:                      "kde-center-page-section-detail-preview",
		RuntimeMethod:               "GetDiagnostics",
		ReadMethod:                  "GetDiagnosticHistoryPreview",
		ReadModel:                   "diagnostic-history-preview",
		CLICommand:                  "diagnostic-history-preview",
		ApplicationID:               applicationID,
		UserVisible:                 true,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootRequired:           true,
		StateRootPathExposed:        false,
		HistoryPreviewCreated:       false,
		AIProviderCallEnabled:       false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		RequestObjectCreated:        false,
		PermissionGranted:           false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		RepairExecutionEnabled:      false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          "KDE can request the Runtime diagnostic history preview for this application, but the section detail does not expose state-root paths or enable actions.",
	}
}

func NewKDECenterPageSectionsPreview(recipe Recipe, provenance Provenance, decision string, fileURIs []string) (KDECenterPageSectionsPreview, error) {
	page, err := NewKDECenterPagePreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return KDECenterPageSectionsPreview{}, err
	}

	readinessEvidence := kdeCenterPageReadinessEvidence(page.ApplicationReadinessEvidence, nil)
	sections := kdeCenterPageSections(page.ApplicationReadinessEvidence)
	aiAnalysis := firstKDECenterPageSectionAIAnalysis(sections)
	preview := KDECenterPageSectionsPreview{
		SchemaVersion:                "xnix.runtime.kde_center_page_sections.v1",
		RequestType:                  "kde-center-page-sections-preview",
		PageType:                     page.PageType,
		Source:                       "kde-center-page-preview",
		Desktop:                      "KDE Plasma",
		RuntimeMethod:                "GetKDECenterPageSections",
		ReadMethod:                   "GetKDECenterPageSectionsPreview",
		ApplicationID:                page.ApplicationID,
		ApplicationName:              page.ApplicationName,
		SectionCount:                 len(sections),
		ReadOnlySectionCount:         len(sections),
		NavigationOnlySectionCount:   len(sections),
		ExecutableSectionCount:       0,
		AIAnalysis:                   aiAnalysis,
		AIAnalysisSectionCount:       countKDECenterPageSectionAIAnalysis(sections),
		ApplicationReadinessEvidence: readinessEvidence,
		PrimarySectionID:             "overview",
		Sections:                     sections,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OfficialDesktopOnly:          true,
		UserVisible:                  true,
		SafeForAIDiagnostics:         true,
		UserDecisionCaptured:         page.UserDecisionCaptured,
		UserDecisionAllowsLaunch:     page.UserDecisionAllowsLaunch,
		SectionsPreviewCreated:       true,
		SectionsPersisted:            false,
		SectionActionsEnabled:        false,
		SettingsPersisted:            false,
		SettingsPersistenceEnabled:   false,
		NotificationsSent:            false,
		ResourceGrantCreated:         false,
		RuntimeLaunchApproval:        false,
		LaunchEnabled:                false,
		ExecutionStarted:             false,
		RequestObjectsCreated:        false,
		PermissionGrantCreated:       false,
		HostRootModified:             false,
		NetworkRequired:              false,
		BackendDetailsExposed:        false,
		BlockedActions:               []string{"persist KDE center page sections from preview state", "enable section action buttons from preview state", "persist compatibility settings from section navigation", "record review receipts from page sections", "create Runtime request objects from page sections", "grant desktop resources from page sections", "send desktop notifications from page sections", "start compatibility profile from page sections", "mutate host root during KDE center page section preview", "expose raw backend command to desktop shell"},
		DesktopSafeSummary:           "KDE can navigate Compatibility Center page sections backed by Runtime read models and application readiness evidence, but sections remain read-only and cannot persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page sections preview"); err != nil {
		return KDECenterPageSectionsPreview{}, err
	}
	return preview, nil
}

func NewKDECenterPageSectionDetailPreview(recipe Recipe, provenance Provenance, sectionID string, decision string, fileURIs []string) (KDECenterPageSectionDetailPreview, error) {
	if !singleLine(sectionID) {
		return KDECenterPageSectionDetailPreview{}, errors.New("KDE center page section detail preview requires a single-line section id")
	}

	sections, err := NewKDECenterPageSectionsPreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}

	section, ok := findKDECenterPageSection(sections.Sections, sectionID)
	if !ok {
		return KDECenterPageSectionDetailPreview{}, errors.New("KDE center page section detail preview requires a known section id")
	}
	aiAnalysisInput, err := kdeCenterPageAIAnalysisInput(recipe, provenance, section.ID, fileURIs)
	if err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}
	diagnosticHistoryRoute := kdeDiagnosticHistoryRoute(section.ID, sections.ApplicationID)

	preview := KDECenterPageSectionDetailPreview{
		SchemaVersion:                "xnix.runtime.kde_center_page_section_detail.v1",
		RequestType:                  "kde-center-page-section-detail-preview",
		PageType:                     sections.PageType,
		Source:                       "kde-center-page-sections-preview",
		Desktop:                      "KDE Plasma",
		RuntimeMethod:                "GetKDECenterPageSectionDetail",
		ReadMethod:                   "GetKDECenterPageSectionDetailPreview",
		ApplicationID:                sections.ApplicationID,
		ApplicationName:              sections.ApplicationName,
		SectionID:                    section.ID,
		SectionLabel:                 section.Label,
		SectionTarget:                section.Target,
		SectionState:                 section.State,
		SectionRuntimeMethod:         section.RuntimeMethod,
		SectionReadModel:             section.ReadModel,
		SectionSummary:               section.Summary,
		ReadinessNodeIDs:             section.ReadinessNodeIDs,
		ReadinessStatus:              section.ReadinessStatus,
		ReadinessBlocked:             section.ReadinessBlocked,
		ReadinessSummary:             section.ReadinessSummary,
		ApplicationReadinessEvidence: kdeCenterPageReadinessEvidenceForSection(sections.ApplicationReadinessEvidence, section),
		DiagnosticHistoryRoute:       diagnosticHistoryRoute,
		AIAnalysis:                   section.AIAnalysis,
		AIAnalysisInput:              aiAnalysisInput,
		AvailableSectionIDs:          kdeCenterPageSectionIDs(sections.Sections),
		ReadOnlyNavigation:           true,
		DetailPreviewCreated:         true,
		DetailPersisted:              false,
		SectionActionsEnabled:        false,
		SettingsPersisted:            false,
		SettingsPersistenceEnabled:   false,
		NotificationsSent:            false,
		ResourceGrantCreated:         false,
		RuntimeLaunchApproval:        false,
		LaunchEnabled:                false,
		ExecutionStarted:             false,
		RequestObjectsCreated:        false,
		PermissionGrantCreated:       false,
		HostRootModified:             false,
		NetworkRequired:              false,
		BackendDetailsExposed:        false,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		OfficialDesktopOnly:          true,
		UserVisible:                  true,
		SafeForAIDiagnostics:         true,
		UserDecisionCaptured:         sections.UserDecisionCaptured,
		UserDecisionAllowsLaunch:     sections.UserDecisionAllowsLaunch,
		BlockedActions:               []string{"persist KDE center page section detail from preview state", "enable section action buttons from detail preview", "persist compatibility settings from section detail", "record review receipts from section detail", "create Runtime request objects from section detail", "grant desktop resources from section detail", "send desktop notifications from section detail", "start compatibility profile from section detail", "mutate host root during KDE center page section detail preview", "expose raw backend command to desktop shell"},
		DesktopSafeSummary:           "KDE can open a selected Compatibility Center section and route it to a Runtime read model plus application readiness evidence, but the detail remains read-only and cannot persist, grant, notify, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE center page section detail preview"); err != nil {
		return KDECenterPageSectionDetailPreview{}, err
	}
	return preview, nil
}

func kdeCenterPageNavigation() []KDECenterPageNavigation {
	return []KDECenterPageNavigation{
		kdeCenterPageNavigationItem("overview", "Overview", "compatibility-center-overview"),
		kdeCenterPageNavigationItem("backend", "Backend", "compatibility-backend-selection"),
		kdeCenterPageNavigationItem("activation", "Activation", "compatibility-activation-status"),
		kdeCenterPageNavigationItem("execution", "Execution", "compatibility-execution-readiness"),
		kdeCenterPageNavigationItem("launch", "Launch", "compatibility-launch-intent"),
		kdeCenterPageNavigationItem("window", "Window", "compatibility-window-identity"),
		kdeCenterPageNavigationItem("files", "Files", "compatibility-file-association"),
		kdeCenterPageNavigationItem("tray", "Tray", "compatibility-tray-status"),
		kdeCenterPageNavigationItem("notifications", "Notifications", "compatibility-notification-plan"),
		kdeCenterPageNavigationItem("actions", "Actions", "compatibility-center-gates"),
		kdeCenterPageNavigationItem("settings", "Settings", "compatibility-settings"),
		kdeCenterPageNavigationItem("diagnostics", "Diagnostics", "compatibility-diagnostics"),
	}
}

func kdeCenterPageNavigationItem(id string, label string, target string) KDECenterPageNavigation {
	return KDECenterPageNavigation{
		ID:             id,
		Label:          label,
		Target:         target,
		Enabled:        true,
		NavigationOnly: true,
		MutatesRuntime: false,
		StartsProgram:  false,
	}
}

func kdeCenterPageSections(readiness ApplicationReadinessPreview) []KDECenterPageSection {
	return []KDECenterPageSection{
		kdeCenterPageSection(readiness, "overview", "Overview", "compatibility-center-overview", "GetCompatibilityCenterSummary", "compatibility-center-summary", "ready", "Overview reads Runtime-owned application state, known issue counts, repair record state, and application readiness evidence."),
		kdeCenterPageSection(readiness, "backend", "Backend", "compatibility-backend-selection", "GetBackendSelectionPlan", "backend-selection-preview", "selection-pending", "Backend reads Runtime-owned recommended compatibility profile and backend lifecycle evidence while selection commit, environment creation, and launch remain closed."),
		kdeCenterPageSection(readiness, "activation", "Activation", "compatibility-activation-status", "GetDesktopActivationStatus", "desktop-activation-status-preview", "ready-for-runtime-commit", "Activation reads Runtime-owned KDE desktop activation status and artifact staging evidence while commit, launch, and host mutation gates remain closed."),
		kdeCenterPageSection(readiness, "execution", "Execution", "compatibility-execution-readiness", "GetExecutionReadiness", "execution-readiness-preview", "blocked", "Execution reads Runtime-owned launch readiness and write-gate evidence while request creation, launch, and backend process gates remain closed."),
		kdeCenterPageSection(readiness, "launch", "Launch", "compatibility-launch-intent", "GetLaunchIntent", "launch-intent-preview", "blocked", "Launch reads Runtime-owned desktop-launch intent, execution readiness, and Launch write-gate evidence while request creation, permission grants, execution, and backend process gates remain closed."),
		kdeCenterPageSection(readiness, "window", "Window", "compatibility-window-identity", "GetTaskManagerIdentityPlan", "window-identity-preview", "planned", "Window reads Runtime-owned task-manager, KWin identity hints, backend lifecycle evidence, and execution state while task-manager activation, KWin rule application, execution, and backend policy stay closed."),
		kdeCenterPageSection(readiness, "files", "Files", "compatibility-file-association", "GetFileAssociationPlan", "file-association-plan", "planned", "Files read Runtime-owned MIME association, Dolphin open-action plans, and Portal review evidence while MIME writes, direct host-file access, permission grants, and execution stay closed."),
		kdeCenterPageSection(readiness, "tray", "Tray", "compatibility-tray-status", "GetTrayStatus", "tray-status-preview", "planned", "Tray reads Runtime-owned compatibility status, tray bridge status, backend lifecycle evidence, and execution state while live backend tray bridging, persistence, notifications, and execution stay closed."),
		kdeCenterPageSection(readiness, "notifications", "Notifications", "compatibility-notification-plan", "GetNotificationPlan", "notification-preview", "planned", "Notifications read Runtime-owned event plans plus Portal and snapshot evidence while notification delivery, action execution, repair execution, settings persistence, and host mutation stay closed."),
		kdeCenterPageSection(readiness, "actions", "Actions", "compatibility-center-gates", "GetCompatibilityActionQueue", "compatibility-center-action-queue", "waiting-for-runtime-gates", "Actions read queued review cards and all required readiness gates while all execution and mutation gates remain closed."),
		kdeCenterPageSection(readiness, "settings", "Settings", "compatibility-settings", "GetCompatibilitySettings", "settings-model", "planned", "Settings read user-facing Runtime policy with Portal and snapshot evidence without persisting changes from KDE."),
		kdeCenterPageSection(readiness, "diagnostics", "Diagnostics", "compatibility-diagnostics", "GetDiagnostics", "diagnostic-history-preview", "planned", "Diagnostics read Runtime-owned diagnostic state, recipe trust, execution readiness, and history while Dolphin file analysis remains an AI-safe optional link."),
	}
}

func kdeCenterPageSection(readiness ApplicationReadinessPreview, id string, label string, target string, runtimeMethod string, readModel string, state string, summary string) KDECenterPageSection {
	nodeIDs := kdeCenterPageSectionReadinessNodeIDs(id, readiness.NodeIDs)
	status, blocked := kdeCenterPageReadinessStatus(readiness, nodeIDs)
	return KDECenterPageSection{
		ID:                    id,
		Label:                 label,
		Target:                target,
		RuntimeMethod:         runtimeMethod,
		ReadModel:             readModel,
		State:                 state,
		ReadinessNodeIDs:      nodeIDs,
		ReadinessStatus:       status,
		ReadinessBlocked:      blocked,
		ReadinessSummary:      kdeCenterPageSectionReadinessSummary(id, nodeIDs, status),
		AIAnalysis:            kdeCenterPageSectionAIAnalysis(id),
		NavigationOnly:        true,
		ReadOnly:              true,
		MutatesRuntime:        false,
		StartsProgram:         false,
		SettingsPersisted:     false,
		BackendDetailsExposed: false,
		Summary:               summary,
	}
}

func kdeCenterPageReadinessEvidence(readiness ApplicationReadinessPreview, nodeIDs []string) KDECenterPageReadinessEvidence {
	if len(nodeIDs) == 0 {
		nodeIDs = append([]string(nil), readiness.NodeIDs...)
	}
	statuses := make(map[string]string, len(nodeIDs))
	readyCount := 0
	blockedCount := 0
	for _, id := range nodeIDs {
		if node, ok := applicationReadinessNodeByID(readiness.Nodes, id); ok {
			statuses[id] = node.Status
			if node.Ready {
				readyCount++
			}
			if node.Status == "blocked" {
				blockedCount++
			}
		}
	}
	return KDECenterPageReadinessEvidence{
		Source:                     readiness.RequestType,
		GraphType:                  readiness.GraphType,
		RuntimeMethod:              readiness.RuntimeMethod,
		ReadMethod:                 readiness.ReadMethod,
		OverallStatus:              readiness.OverallStatus,
		NodeIDs:                    nodeIDs,
		NodeStatuses:               statuses,
		NodeCount:                  len(nodeIDs),
		RequiredNodeCount:          len(nodeIDs),
		BlockedNodeCount:           blockedCount,
		ReadyNodeCount:             readyCount,
		LaunchAllowed:              readiness.LaunchAllowed,
		LaunchEnabled:              readiness.LaunchEnabled,
		ExecutionRequestCreated:    readiness.ExecutionRequestCreated,
		ExecutionStarted:           readiness.ExecutionStarted,
		BackendProcessStarted:      readiness.BackendProcessStarted,
		RealPortalTransportEnabled: readiness.RealPortalTransportEnabled,
		RequestObjectCreated:       readiness.RequestObjectCreated,
		PermissionGranted:          readiness.PermissionGranted,
		SnapshotCreated:            readiness.SnapshotCreated,
		RestoreExecuted:            readiness.RestoreExecuted,
		HostRootModified:           readiness.HostRootModified,
		NetworkRequired:            readiness.NetworkRequired,
		BackendDetailsExposed:      readiness.BackendDetailsExposed,
		Summary:                    "KDE section navigation reads Runtime-owned application readiness evidence without creating requests, granting permissions, starting execution, or mutating the host root.",
	}
}

func kdeCenterPageReadinessEvidenceForSection(evidence KDECenterPageReadinessEvidence, section KDECenterPageSection) KDECenterPageReadinessEvidence {
	statuses := make(map[string]string, len(section.ReadinessNodeIDs))
	readyCount := 0
	blockedCount := 0
	for _, id := range section.ReadinessNodeIDs {
		status := evidence.NodeStatuses[id]
		statuses[id] = status
		if status == "pass" {
			readyCount++
		}
		if status == "blocked" {
			blockedCount++
		}
	}
	evidence.NodeIDs = append([]string(nil), section.ReadinessNodeIDs...)
	evidence.NodeStatuses = statuses
	evidence.NodeCount = len(section.ReadinessNodeIDs)
	evidence.RequiredNodeCount = len(section.ReadinessNodeIDs)
	evidence.ReadyNodeCount = readyCount
	evidence.BlockedNodeCount = blockedCount
	evidence.Summary = "KDE section detail reads the selected Runtime readiness nodes without creating requests, granting permissions, starting execution, or mutating the host root."
	return evidence
}

func kdeCenterPageSectionReadinessNodeIDs(sectionID string, allNodeIDs []string) []string {
	switch sectionID {
	case "overview", "actions":
		return append([]string(nil), allNodeIDs...)
	case "backend", "window", "tray":
		return []string{"backend-lifecycle", "execution-readiness"}
	case "activation":
		return []string{"artifact-stage-receipt", "snapshot-baseline"}
	case "execution", "launch":
		return []string{"execution-readiness", "runtime-write-gate"}
	case "files":
		return []string{"portal-review", "execution-readiness"}
	case "notifications", "settings":
		return []string{"portal-review", "snapshot-baseline"}
	case "diagnostics":
		return []string{"recipe-trust", "execution-readiness"}
	default:
		return []string{}
	}
}

func kdeCenterPageReadinessStatus(readiness ApplicationReadinessPreview, nodeIDs []string) (string, bool) {
	blocked := false
	required := false
	for _, id := range nodeIDs {
		node, ok := applicationReadinessNodeByID(readiness.Nodes, id)
		if !ok {
			continue
		}
		switch node.Status {
		case "blocked":
			blocked = true
		case "required":
			required = true
		}
	}
	if blocked {
		return "blocked", true
	}
	if required {
		return "required", false
	}
	return "pass", false
}

func applicationReadinessNodeByID(nodes []ApplicationReadinessNode, id string) (ApplicationReadinessNode, bool) {
	for _, node := range nodes {
		if node.ID == id {
			return node, true
		}
	}
	return ApplicationReadinessNode{}, false
}

func kdeCenterPageSectionReadinessSummary(sectionID string, nodeIDs []string, status string) string {
	if len(nodeIDs) == 0 {
		return "No Runtime readiness nodes are associated with this section."
	}
	return "This section is backed by Runtime application readiness nodes and currently reports " + status + " evidence."
}

func kdeCenterPageSectionAIAnalysis(sectionID string) *KDEAIAnalysisLink {
	if sectionID != "diagnostics" {
		return nil
	}
	link := kdeAIAnalysisLinkForEntryPoint("file-manager")
	return link
}

func firstKDECenterPageSectionAIAnalysis(sections []KDECenterPageSection) *KDEAIAnalysisLink {
	for _, section := range sections {
		if section.AIAnalysis != nil {
			return section.AIAnalysis
		}
	}
	return nil
}

func countKDECenterPageSectionAIAnalysis(sections []KDECenterPageSection) int {
	count := 0
	for _, section := range sections {
		if section.AIAnalysis != nil {
			count++
		}
	}
	return count
}

func findKDECenterPageSection(sections []KDECenterPageSection, sectionID string) (KDECenterPageSection, bool) {
	for _, section := range sections {
		if section.ID == sectionID {
			return section, true
		}
	}
	return KDECenterPageSection{}, false
}

func kdeCenterPageSectionIDs(sections []KDECenterPageSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func kdeCenterPageBadge(deck KDEActionCardDeckPreview) string {
	if deck.RejectedCardCount > 0 {
		return "Needs guidance"
	}
	if deck.DeferredCardCount > 0 {
		return "Deferred"
	}
	return "Review ready"
}

func kdeCenterPageBadgeTone(deck KDEActionCardDeckPreview) string {
	if deck.RejectedCardCount > 0 {
		return "critical"
	}
	if deck.DeferredCardCount > 0 {
		return "neutral"
	}
	return "warning"
}

func kdeCenterPageFileAssociationSummary(ready bool) string {
	if ready {
		return "Dolphin can show Runtime-owned file associations and portal-mediated open actions without writing MIME defaults or reading host files."
	}
	return "This application has no file associations yet; KDE can still show the Runtime-owned application page without writing MIME defaults or reading host files."
}
