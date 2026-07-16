package appidentity

import (
	"errors"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/diagnostics"
	"xnix.local/xnix/internal/runtime/safety"
)

type CrashHangSignalSummaryOptions struct {
	ApplicationID         string
	MalformedHistory      bool
	MalformedEvidenceIDs  []string
	HistoryReadError      string
	PrivacySensitiveHints []string
}

type CrashHangSignalSummaryPreview struct {
	SchemaVersion               string                          `json:"schema_version"`
	RequestType                 string                          `json:"request_type"`
	SummaryType                 string                          `json:"summary_type"`
	Source                      string                          `json:"source"`
	Desktop                     string                          `json:"desktop"`
	RuntimeMethod               string                          `json:"runtime_method"`
	ReadMethod                  string                          `json:"read_method"`
	ApplicationID               string                          `json:"application_id,omitempty"`
	OverallState                string                          `json:"overall_state"`
	SignalGroups                []CrashHangSignalGroup          `json:"signal_groups"`
	SignalGroupIDs              []string                        `json:"signal_group_ids"`
	SignalGroupCount            int                             `json:"signal_group_count"`
	Counts                      CrashHangSignalSummaryCounts    `json:"counts"`
	LatestKnownState            string                          `json:"latest_known_state"`
	LatestRunID                 string                          `json:"latest_run_id,omitempty"`
	RecurrenceHints             []string                        `json:"recurrence_hints"`
	PrivacyRedaction            CrashHangSignalPrivacyRedaction `json:"privacy_redaction"`
	MalformedHistory            bool                            `json:"malformed_history"`
	MalformedEvidenceIDs        []string                        `json:"malformed_evidence_ids"`
	MixedApplicationRecordCount int                             `json:"mixed_application_record_count"`
	DuplicateSignalIDCount      int                             `json:"duplicate_signal_id_count"`
	BlockedRepairEvidenceCount  int                             `json:"blocked_repair_evidence_count"`
	RuntimeOwned                bool                            `json:"runtime_owned"`
	GoRuntimeBacked             bool                            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                            `json:"kde_policy_owner"`
	UserVisible                 bool                            `json:"user_visible"`
	ReviewOnly                  bool                            `json:"review_only"`
	PrivateLogRead              bool                            `json:"private_log_read"`
	FileContentRead             bool                            `json:"file_content_read"`
	FilePathsExposed            bool                            `json:"file_paths_exposed"`
	AIProviderCalled            bool                            `json:"ai_provider_called"`
	AIProviderCallEnabled       bool                            `json:"ai_provider_call_enabled"`
	RepairExecutionRequested    bool                            `json:"repair_execution_requested"`
	RepairExecuted              bool                            `json:"repair_executed"`
	RepairApplied               bool                            `json:"repair_applied"`
	BackendProcessStarted       bool                            `json:"backend_process_started"`
	NetworkRequired             bool                            `json:"network_required"`
	HostRootModified            bool                            `json:"host_root_modified"`
	StateRootPathExposed        bool                            `json:"state_root_path_exposed"`
	BackendDetailsExposed       bool                            `json:"backend_details_exposed"`
	PrivilegedContainerRequired bool                            `json:"privileged_container_required"`
	NextSafeReadOnlyChecks      []string                        `json:"next_safe_read_only_checks"`
	BlockedActions              []string                        `json:"blocked_actions"`
	DesktopSafeSummary          string                          `json:"desktop_safe_summary"`
}

type CrashHangSignalGroup struct {
	ID                    string   `json:"id"`
	Title                 string   `json:"title"`
	SignalClass           string   `json:"signal_class"`
	Count                 int      `json:"count"`
	LatestKnownState      string   `json:"latest_known_state"`
	RelatedRunIDs         []string `json:"related_diagnostic_run_ids"`
	RelatedSignalIDs      []string `json:"related_signal_ids"`
	RecurrenceHint        string   `json:"recurrence_hint"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	UserSafeSummary       string   `json:"user_safe_summary"`
	ReviewRequired        bool     `json:"review_required"`
	RepairBlocked         bool     `json:"repair_blocked"`
	PrivateLogRead        bool     `json:"private_log_read"`
	FileContentRead       bool     `json:"file_content_read"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	RepairExecuted        bool     `json:"repair_executed"`
	BackendProcessStarted bool     `json:"backend_process_started"`
	HostRootModified      bool     `json:"host_root_modified"`
}

type CrashHangSignalSummaryCounts struct {
	TotalRuns                  int `json:"total_runs"`
	MatchedRuns                int `json:"matched_runs"`
	FailedRuns                 int `json:"failed_runs"`
	BlockedRuns                int `json:"blocked_runs"`
	SignalOccurrences          int `json:"signal_occurrences"`
	GroupsWithSignals          int `json:"groups_with_signals"`
	RepeatedGroups             int `json:"repeated_groups"`
	MalformedRecords           int `json:"malformed_records"`
	MixedApplicationRecords    int `json:"mixed_application_records"`
	DuplicateSignalIDs         int `json:"duplicate_signal_ids"`
	PrivacySensitiveOmissions  int `json:"privacy_sensitive_omissions"`
	BlockedRepairEvidenceCount int `json:"blocked_repair_evidence_count"`
}

type CrashHangSignalPrivacyRedaction struct {
	RedactionStatus              string `json:"redaction_status"`
	PrivateLogsIncluded          bool   `json:"private_logs_included"`
	FileContentsIncluded         bool   `json:"file_contents_included"`
	HostPathsIncluded            bool   `json:"host_paths_included"`
	StateRootPathsIncluded       bool   `json:"state_root_paths_included"`
	RawCommandsIncluded          bool   `json:"raw_commands_included"`
	EnvironmentVariablesIncluded bool   `json:"environment_variables_included"`
	TokenShapedValuesIncluded    bool   `json:"token_shaped_values_included"`
	UsernamesIncluded            bool   `json:"usernames_included"`
	SecretsIncluded              bool   `json:"secrets_included"`
	BackendDetailsIncluded       bool   `json:"backend_details_included"`
	OmittedSensitiveHints        int    `json:"omitted_sensitive_hints"`
}

type crashHangSignalAccumulator struct {
	group                CrashHangSignalGroup
	relatedRunIDs        []string
	relatedSignalIDs     []string
	seenRunGroup         map[string]bool
	seenSignalIDs        map[string]bool
	duplicateSignalCount int
}

func NewCrashHangSignalSummaryPreview(history diagnostics.RunHistory, options CrashHangSignalSummaryOptions) (CrashHangSignalSummaryPreview, error) {
	options.ApplicationID = strings.TrimSpace(options.ApplicationID)
	if options.ApplicationID == "" {
		options.ApplicationID = history.ApplicationID
	}
	if options.ApplicationID != "" && !idPattern.MatchString(options.ApplicationID) {
		return CrashHangSignalSummaryPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}

	accumulators := crashHangSignalAccumulators()
	var counts CrashHangSignalSummaryCounts
	var latestRunID string
	var latestKnownState string
	var blockedRepairEvidence int
	mixedApplicationRecords := 0
	privacySensitiveOmissions := len(options.PrivacySensitiveHints)

	for _, record := range history.Records {
		if options.ApplicationID != "" && record.ApplicationID != "" && record.ApplicationID != options.ApplicationID {
			mixedApplicationRecords++
			continue
		}
		counts.TotalRuns++
		switch record.Overall {
		case diagnostics.OutcomeFail:
			counts.FailedRuns++
		case diagnostics.OutcomeBlocked:
			counts.BlockedRuns++
		}
		if record.RunID >= latestRunID {
			latestRunID = record.RunID
			latestKnownState = string(record.Overall)
		}
		if record.RepairIssue != "" && record.RepairIssue != "runtime-repair-applied" {
			blockedRepairEvidence++
		}
		for _, rawSignalID := range record.FailingIDs {
			signalID := strings.TrimSpace(rawSignalID)
			if signalID == "" {
				continue
			}
			if crashHangPrivacySensitive(signalID) {
				privacySensitiveOmissions++
				continue
			}
			groupID := crashHangGroupForSignal(signalID)
			acc := accumulators[groupID]
			runGroupKey := record.RunID + "\x00" + groupID
			if !acc.seenRunGroup[runGroupKey] {
				acc.group.Count++
				acc.relatedRunIDs = append(acc.relatedRunIDs, record.RunID)
				acc.seenRunGroup[runGroupKey] = true
				counts.SignalOccurrences++
			}
			if acc.seenSignalIDs[signalID] {
				acc.duplicateSignalCount++
			} else {
				acc.seenSignalIDs[signalID] = true
				acc.relatedSignalIDs = append(acc.relatedSignalIDs, signalID)
			}
			if record.Overall != "" {
				acc.group.LatestKnownState = string(record.Overall)
			}
			if record.RepairIssue != "" && record.RepairIssue != "runtime-repair-applied" {
				acc.group.RepairBlocked = true
			}
			accumulators[groupID] = acc
		}
		if record.RepairIssue == "runtime-repair-applied" && record.Overall == diagnostics.OutcomeFail {
			acc := accumulators["regression-after-repair"]
			if !acc.seenRunGroup[record.RunID+"\x00regression-after-repair"] {
				acc.group.Count++
				acc.relatedRunIDs = append(acc.relatedRunIDs, record.RunID)
				acc.seenRunGroup[record.RunID+"\x00regression-after-repair"] = true
				counts.SignalOccurrences++
			}
			acc.group.LatestKnownState = string(record.Overall)
			acc.relatedSignalIDs = append(acc.relatedSignalIDs, "repair-regression")
			accumulators["regression-after-repair"] = acc
		}
	}

	groups := finalizeCrashHangSignalGroups(accumulators)
	for _, group := range groups {
		if group.Count > 0 {
			counts.GroupsWithSignals++
			counts.MatchedRuns += len(group.RelatedRunIDs)
		}
		if group.Count > 1 {
			counts.RepeatedGroups++
		}
	}
	counts.MalformedRecords = len(options.MalformedEvidenceIDs)
	if options.MalformedHistory && counts.MalformedRecords == 0 {
		counts.MalformedRecords = 1
	}
	counts.MixedApplicationRecords = mixedApplicationRecords
	counts.DuplicateSignalIDs = crashHangDuplicateSignalCount(accumulators)
	counts.PrivacySensitiveOmissions = privacySensitiveOmissions
	counts.BlockedRepairEvidenceCount = blockedRepairEvidence

	preview := CrashHangSignalSummaryPreview{
		SchemaVersion:               "xnix.runtime.crash_hang_signal_summary.v1",
		RequestType:                 "crash-hang-signal-summary-preview",
		SummaryType:                 "privacy-safe-crash-hang-signal-summary",
		Source:                      "diagnostic-history+fixture-metadata+repair-evidence",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetCrashHangSignalSummary",
		ReadMethod:                  "GetCrashHangSignalSummaryPreview",
		ApplicationID:               options.ApplicationID,
		OverallState:                crashHangOverallState(counts),
		SignalGroups:                groups,
		SignalGroupIDs:              crashHangSignalGroupIDs(groups),
		SignalGroupCount:            len(groups),
		Counts:                      counts,
		LatestKnownState:            latestKnownStateOrDefault(latestKnownState, counts),
		LatestRunID:                 latestRunID,
		RecurrenceHints:             crashHangRecurrenceHints(groups, counts),
		PrivacyRedaction:            crashHangPrivacyRedaction(privacySensitiveOmissions),
		MalformedHistory:            options.MalformedHistory,
		MalformedEvidenceIDs:        safeEvidenceIDs(options.MalformedEvidenceIDs),
		MixedApplicationRecordCount: mixedApplicationRecords,
		DuplicateSignalIDCount:      counts.DuplicateSignalIDs,
		BlockedRepairEvidenceCount:  blockedRepairEvidence,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		PrivateLogRead:              false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		AIProviderCalled:            false,
		AIProviderCallEnabled:       false,
		RepairExecutionRequested:    false,
		RepairExecuted:              false,
		RepairApplied:               false,
		BackendProcessStarted:       false,
		NetworkRequired:             false,
		HostRootModified:            false,
		StateRootPathExposed:        false,
		BackendDetailsExposed:       false,
		PrivilegedContainerRequired: false,
		NextSafeReadOnlyChecks: []string{
			"review diagnostic run history",
			"review compatibility test result metadata",
			"review repair recommendation metadata",
			"review support bundle manifest preview",
		},
		BlockedActions: []string{
			"read private logs from signal summary",
			"read file contents from signal summary",
			"call AI provider from signal summary",
			"apply repair from signal summary",
			"start compatibility service from signal summary",
			"mutate host root during signal summary preview",
		},
		DesktopSafeSummary: crashHangDesktopSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "crash hang signal summary preview"); err != nil {
		return CrashHangSignalSummaryPreview{}, err
	}
	// Final privacy boundary: reject any host path, executable, URL, secret, or
	// token that slipped through into a user-facing field.
	if err := safety.ValidatePayload("crash hang signal summary preview", preview); err != nil {
		return CrashHangSignalSummaryPreview{}, err
	}
	return preview, nil
}

func crashHangSignalAccumulators() map[string]crashHangSignalAccumulator {
	accumulators := map[string]crashHangSignalAccumulator{}
	for _, id := range crashHangSignalGroupOrder() {
		accumulators[id] = crashHangSignalAccumulator{
			group: CrashHangSignalGroup{
				ID:                    id,
				Title:                 crashHangSignalGroupTitle(id),
				SignalClass:           id,
				LatestKnownState:      "not-observed",
				RecurrenceHint:        "not observed",
				NextSafeReadOnlyCheck: crashHangNextSafeReadOnlyCheck(id),
				UserSafeSummary:       crashHangGroupSummary(id, 0),
				PrivateLogRead:        false,
				FileContentRead:       false,
				AIProviderCalled:      false,
				RepairExecuted:        false,
				BackendProcessStarted: false,
				HostRootModified:      false,
			},
			seenRunGroup:  map[string]bool{},
			seenSignalIDs: map[string]bool{},
		}
	}
	return accumulators
}

func crashHangSignalGroupOrder() []string {
	return []string{
		"crash",
		"hang",
		"timeout",
		"missing-dependency",
		"permission-denial",
		"graphics-issue",
		"network-issue",
		"regression-after-repair",
	}
}

func crashHangSignalGroupTitle(id string) string {
	switch id {
	case "crash":
		return "Crash signals"
	case "hang":
		return "Hang signals"
	case "timeout":
		return "Timeout signals"
	case "missing-dependency":
		return "Missing dependency signals"
	case "permission-denial":
		return "Permission denial signals"
	case "graphics-issue":
		return "Graphics issue signals"
	case "network-issue":
		return "Network issue signals"
	case "regression-after-repair":
		return "Regression after repair signals"
	default:
		return id
	}
}

// crashHangGroupForSignal classifies a failing signal by its stable signal id
// only. It deliberately ignores the run's repair issue and test type: those
// carry Runtime-derived terms such as "engine-binding-pending" or
// "repair-readiness" that would otherwise shadow the real crash, hang, timeout,
// graphics, network, and permission groups.
func crashHangGroupForSignal(signalID string) string {
	text := strings.ToLower(signalID)
	switch {
	case strings.Contains(text, "hang") || strings.Contains(text, "freeze") || strings.Contains(text, "frozen") || strings.Contains(text, "stuck") || strings.Contains(text, "unresponsive") || strings.Contains(text, "deadlock"):
		return "hang"
	case strings.Contains(text, "timeout") || strings.Contains(text, "timed-out") || strings.Contains(text, "deadline"):
		return "timeout"
	case strings.Contains(text, "dependency") || strings.Contains(text, "missing") || strings.Contains(text, "library") || strings.Contains(text, "unresolved") || strings.Contains(text, "not-found") || strings.Contains(text, "engine-binding") || strings.Contains(text, "runtime-launch-binding"):
		return "missing-dependency"
	case strings.Contains(text, "permission") || strings.Contains(text, "portal") || strings.Contains(text, "denied") || strings.Contains(text, "unauthorized") || strings.Contains(text, "forbidden") || strings.Contains(text, "approval"):
		return "permission-denial"
	case strings.Contains(text, "graphics") || strings.Contains(text, "gpu") || strings.Contains(text, "display") || strings.Contains(text, "render") || strings.Contains(text, "vulkan") || strings.Contains(text, "opengl") || strings.Contains(text, "directx"):
		return "graphics-issue"
	case strings.Contains(text, "network") || strings.Contains(text, "offline") || strings.Contains(text, "connection") || strings.Contains(text, "socket") || strings.Contains(text, "dns") || strings.Contains(text, "unreachable"):
		return "network-issue"
	case strings.Contains(text, "regression"):
		return "regression-after-repair"
	case strings.Contains(text, "crash") || strings.Contains(text, "abort") || strings.Contains(text, "segfault") || strings.Contains(text, "fault") || strings.Contains(text, "exit"):
		return "crash"
	default:
		return "crash"
	}
}

func finalizeCrashHangSignalGroups(accumulators map[string]crashHangSignalAccumulator) []CrashHangSignalGroup {
	groups := make([]CrashHangSignalGroup, 0, len(accumulators))
	for _, id := range crashHangSignalGroupOrder() {
		acc := accumulators[id]
		group := acc.group
		group.RelatedRunIDs = uniqueStrings(acc.relatedRunIDs)
		group.RelatedSignalIDs = uniqueStrings(acc.relatedSignalIDs)
		group.ReviewRequired = group.Count > 0
		group.RecurrenceHint = crashHangRecurrenceHint(group.Count)
		group.UserSafeSummary = crashHangGroupSummary(group.ID, group.Count)
		groups = append(groups, group)
	}
	return groups
}

func crashHangDuplicateSignalCount(accumulators map[string]crashHangSignalAccumulator) int {
	total := 0
	for _, acc := range accumulators {
		total += acc.duplicateSignalCount
	}
	return total
}

func crashHangRecurrenceHint(count int) string {
	switch {
	case count == 0:
		return "not observed"
	case count == 1:
		return "observed once"
	default:
		return "repeated across diagnostic metadata"
	}
}

func crashHangGroupSummary(id string, count int) string {
	if count == 0 {
		return "No " + crashHangSignalGroupTitle(id) + " were found in diagnostic metadata."
	}
	return crashHangSignalGroupTitle(id) + " were found in diagnostic metadata and should be reviewed before any repair is offered."
}

func crashHangNextSafeReadOnlyCheck(id string) string {
	switch id {
	case "permission-denial":
		return "Review Portal permission receipt metadata."
	case "regression-after-repair":
		return "Review repair recommendation metadata and snapshot readiness."
	case "missing-dependency":
		return "Review compatibility test result metadata and install readiness."
	default:
		return "Review diagnostic run history metadata."
	}
}

func crashHangSignalGroupIDs(groups []CrashHangSignalGroup) []string {
	ids := make([]string, 0, len(groups))
	for _, group := range groups {
		ids = append(ids, group.ID)
	}
	return ids
}

func crashHangOverallState(counts CrashHangSignalSummaryCounts) string {
	switch {
	case counts.MalformedRecords > 0:
		return "blocked-malformed-history"
	case counts.TotalRuns == 0:
		return "no-history"
	case counts.GroupsWithSignals > 0 || counts.FailedRuns > 0:
		return "needs-review"
	case counts.BlockedRuns > 0:
		return "blocked"
	default:
		return "healthy"
	}
}

func latestKnownStateOrDefault(value string, counts CrashHangSignalSummaryCounts) string {
	if value != "" {
		return value
	}
	if counts.MalformedRecords > 0 {
		return "blocked"
	}
	return "not-run"
}

func crashHangRecurrenceHints(groups []CrashHangSignalGroup, counts CrashHangSignalSummaryCounts) []string {
	hints := []string{}
	if counts.MalformedRecords > 0 {
		hints = append(hints, "Some diagnostic receipts could not be parsed and were not used for signal grouping.")
	}
	if counts.PrivacySensitiveOmissions > 0 {
		hints = append(hints, "Sensitive-looking diagnostic metadata was omitted from this preview.")
	}
	for _, group := range groups {
		if group.Count > 1 {
			hints = append(hints, group.Title+" are repeated.")
		}
	}
	if len(hints) == 0 {
		hints = append(hints, "No repeated crash or hang pattern was found in available metadata.")
	}
	return uniqueStrings(hints)
}

func crashHangPrivacyRedaction(omitted int) CrashHangSignalPrivacyRedaction {
	return CrashHangSignalPrivacyRedaction{
		RedactionStatus:              "metadata-only-redacted-preview",
		PrivateLogsIncluded:          false,
		FileContentsIncluded:         false,
		HostPathsIncluded:            false,
		StateRootPathsIncluded:       false,
		RawCommandsIncluded:          false,
		EnvironmentVariablesIncluded: false,
		TokenShapedValuesIncluded:    false,
		UsernamesIncluded:            false,
		SecretsIncluded:              false,
		BackendDetailsIncluded:       false,
		OmittedSensitiveHints:        omitted,
	}
}

func crashHangDesktopSummary(counts CrashHangSignalSummaryCounts) string {
	switch crashHangOverallState(counts) {
	case "blocked-malformed-history":
		return "Some diagnostic history metadata is malformed, so the signal summary is blocked for review without reading private logs."
	case "no-history":
		return "No diagnostic history is available yet for crash or hang signal grouping."
	case "needs-review":
		return "Diagnostic metadata contains crash, hang, timeout, permission, dependency, graphics, network, or repair regression signals for review."
	case "blocked":
		return "Diagnostic metadata contains blocked runs that need review before any repair is offered."
	default:
		return "Diagnostic metadata does not show repeated crash or hang signals."
	}
}

func safeEvidenceIDs(values []string) []string {
	safe := []string{}
	for _, value := range values {
		if value == "" || crashHangPrivacySensitive(value) {
			continue
		}
		safe = append(safe, value)
	}
	sort.Strings(safe)
	return safe
}

func crashHangPrivacySensitive(value string) bool {
	lower := strings.ToLower(value)
	for _, marker := range []string{"/", "\\", "token", "secret", "sk-", ".exe", "program files", "prefix", "wine", "proton", "qemu", "user=", "home="} {
		if strings.Contains(lower, marker) {
			return true
		}
	}
	return false
}
