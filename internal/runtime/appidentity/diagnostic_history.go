package appidentity

import "xnix.local/xnix/internal/runtime/diagnostics"

type DiagnosticHistoryPreview struct {
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	Source                      string                         `json:"source"`
	Desktop                     string                         `json:"desktop"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	ApplicationID               string                         `json:"application_id,omitempty"`
	Records                     []diagnostics.RunHistoryRecord `json:"records"`
	Counts                      diagnostics.RunHistoryCounts   `json:"counts"`
	Latest                      *diagnostics.RunHistoryRecord  `json:"latest,omitempty"`
	CompatibilityCenter         DiagnosticHistoryCenterSummary `json:"compatibility_center"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	UserVisible                 bool                           `json:"user_visible"`
	CompatibilityCenterCard     bool                           `json:"compatibility_center_card"`
	SafeForAIDiagnostics        bool                           `json:"safe_for_ai_diagnostics"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	FileContentRead             bool                           `json:"file_content_read"`
	FilePathsExposed            bool                           `json:"file_paths_exposed"`
	AIProviderCallEnabled       bool                           `json:"ai_provider_call_enabled"`
	RequestObjectCreated        bool                           `json:"request_object_created"`
	PermissionGranted           bool                           `json:"permission_granted"`
	LaunchEnabled               bool                           `json:"launch_enabled"`
	ExecutionStarted            bool                           `json:"execution_started"`
	RepairExecutionEnabled      bool                           `json:"repair_execution_enabled"`
	SettingsPersisted           bool                           `json:"settings_persisted"`
	HostRootModified            bool                           `json:"host_root_modified"`
	NetworkRequired             bool                           `json:"network_required"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	BlockedActions              []string                       `json:"blocked_actions"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type DiagnosticHistoryCenterSummary struct {
	HistoryState           string `json:"history_state"`
	LatestRunID            string `json:"latest_run_id,omitempty"`
	LatestOverall          string `json:"latest_overall,omitempty"`
	LatestRepairIssue      string `json:"latest_repair_issue,omitempty"`
	TotalRunCount          int    `json:"total_run_count"`
	FailingRunCount        int    `json:"failing_run_count"`
	BlockedRunCount        int    `json:"blocked_run_count"`
	ActionExecutionEnabled bool   `json:"action_execution_enabled"`
	RepairExecutionEnabled bool   `json:"repair_execution_enabled"`
	BackendLaunchEnabled   bool   `json:"backend_launch_enabled"`
}

func NewDiagnosticHistoryPreview(history diagnostics.RunHistory) (DiagnosticHistoryPreview, error) {
	center := DiagnosticHistoryCenterSummary{
		HistoryState:           diagnosticHistoryState(history.Counts),
		TotalRunCount:          history.Counts.Total,
		FailingRunCount:        history.Counts.Failed,
		BlockedRunCount:        history.Counts.Blocked,
		ActionExecutionEnabled: false,
		RepairExecutionEnabled: false,
		BackendLaunchEnabled:   false,
	}
	if history.Latest != nil {
		center.LatestRunID = history.Latest.RunID
		center.LatestOverall = string(history.Latest.Overall)
		center.LatestRepairIssue = history.Latest.RepairIssue
	}
	preview := DiagnosticHistoryPreview{
		SchemaVersion:               "xnix.runtime.diagnostic_history_preview.v1",
		RequestType:                 "diagnostic-history-preview",
		Source:                      "go-runtime-state-root-diagnostic-run-history+kde-read-model",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetDiagnostics",
		ReadMethod:                  "GetDiagnosticHistoryPreview",
		ApplicationID:               history.ApplicationID,
		Records:                     append([]diagnostics.RunHistoryRecord(nil), history.Records...),
		Counts:                      history.Counts,
		Latest:                      history.Latest,
		CompatibilityCenter:         center,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		CompatibilityCenterCard:     true,
		SafeForAIDiagnostics:        true,
		StateRootPathExposed:        false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		AIProviderCallEnabled:       false,
		RequestObjectCreated:        false,
		PermissionGranted:           false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		RepairExecutionEnabled:      false,
		SettingsPersisted:           false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"start compatibility execution from diagnostic history",
			"create launch request from diagnostic history",
			"call AI provider from diagnostic history",
			"read selected file contents from diagnostic history",
			"execute repair from diagnostic history",
			"mutate host root during diagnostic history preview",
		},
		DesktopSafeSummary: diagnosticHistorySummary(history.Counts),
	}
	if err := validateNoBackendTerms(preview, "diagnostic history preview"); err != nil {
		return DiagnosticHistoryPreview{}, err
	}
	return preview, nil
}

func diagnosticHistoryState(counts diagnostics.RunHistoryCounts) string {
	switch {
	case counts.Total == 0:
		return "not-run"
	case counts.Failed > 0:
		return "needs-review"
	case counts.Blocked > 0:
		return "blocked"
	default:
		return "healthy"
	}
}

func diagnosticHistorySummary(counts diagnostics.RunHistoryCounts) string {
	switch diagnosticHistoryState(counts) {
	case "not-run":
		return "No diagnostic run history is available for this application."
	case "needs-review":
		return "Diagnostic history contains failing runs for Compatibility Center review."
	case "blocked":
		return "Diagnostic history contains blocked runs awaiting user or Runtime action."
	default:
		return "Diagnostic history contains no failing runs."
	}
}
