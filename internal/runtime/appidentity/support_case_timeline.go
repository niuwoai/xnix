package appidentity

import (
	"errors"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/diagnostics"
	"xnix.local/xnix/internal/runtime/safety"
)

type SupportCaseTimelineOptions struct {
	RuntimeRoot           string
	Issue                 string
	TestType              string
	Decision              string
	FileURIs              []string
	MalformedHistory      bool
	MalformedEvidenceIDs  []string
	PrivacySensitiveHints []string
}

type SupportCaseTimelinePreview struct {
	Version                        string                       `json:"version"`
	SchemaVersion                  string                       `json:"schema_version"`
	RequestType                    string                       `json:"request_type"`
	TimelineType                   string                       `json:"timeline_type"`
	Source                         string                       `json:"source"`
	Desktop                        string                       `json:"desktop"`
	RuntimeMethod                  string                       `json:"runtime_method"`
	ReadMethod                     string                       `json:"read_method"`
	Application                    SupportBundleApplication     `json:"application"`
	Events                         []SupportCaseTimelineEvent   `json:"events"`
	EventIDs                       []string                     `json:"event_ids"`
	EventGroups                    []string                     `json:"event_groups"`
	EventCount                     int                          `json:"event_count"`
	Counts                         SupportCaseTimelineCounts    `json:"counts"`
	RedactionSummary               SupportCaseTimelineRedaction `json:"redaction_summary"`
	MalformedHistory               bool                         `json:"malformed_history"`
	MalformedEvidenceIDs           []string                     `json:"malformed_evidence_ids"`
	MixedApplicationRecordCount    int                          `json:"mixed_application_record_count"`
	BlockedDiagnosticEvidenceCount int                          `json:"blocked_diagnostic_evidence_count"`
	RuntimeOwned                   bool                         `json:"runtime_owned"`
	GoRuntimeBacked                bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                         `json:"kde_policy_owner"`
	UserVisible                    bool                         `json:"user_visible"`
	ReviewOnly                     bool                         `json:"review_only"`
	TicketCreated                  bool                         `json:"ticket_created"`
	BundleExported                 bool                         `json:"bundle_exported"`
	AIProviderCalled               bool                         `json:"ai_provider_called"`
	AIProviderCallEnabled          bool                         `json:"ai_provider_call_enabled"`
	RepairApplied                  bool                         `json:"repair_applied"`
	RepairExecuted                 bool                         `json:"repair_executed"`
	ActionExecuted                 bool                         `json:"action_executed"`
	BackendProcessStarted          bool                         `json:"backend_process_started"`
	FileContentRead                bool                         `json:"file_content_read"`
	FilePathsExposed               bool                         `json:"file_paths_exposed"`
	StateRootPathExposed           bool                         `json:"state_root_path_exposed"`
	RawCommandExposed              bool                         `json:"raw_command_exposed"`
	RawExecutableExposed           bool                         `json:"raw_executable_exposed"`
	BackendDetailsExposed          bool                         `json:"backend_details_exposed"`
	HostRootModified               bool                         `json:"host_root_modified"`
	NetworkRequired                bool                         `json:"network_required"`
	PrivilegedContainerRequired    bool                         `json:"privileged_container_required"`
	NextSafeReadOnlyChecks         []string                     `json:"next_safe_read_only_checks"`
	BlockedActions                 []string                     `json:"blocked_actions"`
	DesktopSafeSummary             string                       `json:"desktop_safe_summary"`
}

type SupportCaseTimelineEvent struct {
	ID                    string   `json:"id"`
	Group                 string   `json:"group"`
	Severity              string   `json:"severity"`
	SourceReadModel       string   `json:"source_read_model"`
	RelatedEvidenceIDs    []string `json:"related_evidence_ids"`
	RelatedRunIDs         []string `json:"related_diagnostic_run_ids,omitempty"`
	UserSafeSummary       string   `json:"user_safe_summary"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	ReviewRequired        bool     `json:"review_required"`
	TicketCreated         bool     `json:"ticket_created"`
	BundleExported        bool     `json:"bundle_exported"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	RepairExecuted        bool     `json:"repair_executed"`
	ActionExecuted        bool     `json:"action_executed"`
	BackendProcessStarted bool     `json:"backend_process_started"`
	HostRootModified      bool     `json:"host_root_modified"`
	FileContentRead       bool     `json:"file_content_read"`
}

type SupportCaseTimelineCounts struct {
	DiagnosticRunEvents        int `json:"diagnostic_run_events"`
	BlockedActionEvents        int `json:"blocked_action_events"`
	RepairRecommendationEvents int `json:"repair_recommendation_events"`
	OnboardingGapEvents        int `json:"onboarding_gap_events"`
	KDEEntrypointEvents        int `json:"kde_entrypoint_events"`
	MalformedHistoryEvents     int `json:"malformed_history_events"`
	BlockedDiagnosticEvents    int `json:"blocked_diagnostic_events"`
	ReviewRequiredEvents       int `json:"review_required_events"`
	RedactedEvidenceOmissions  int `json:"redacted_evidence_omissions"`
}

type SupportCaseTimelineRedaction struct {
	RedactionStatus               string `json:"redaction_status"`
	FileContentsIncluded          bool   `json:"file_contents_included"`
	HostPathsIncluded             bool   `json:"host_paths_included"`
	StateRootPathsIncluded        bool   `json:"state_root_paths_included"`
	EnvironmentVariablesIncluded  bool   `json:"environment_variables_included"`
	UsernamesIncluded             bool   `json:"usernames_included"`
	TokenShapedValuesIncluded     bool   `json:"token_shaped_values_included"`
	CredentialsIncluded           bool   `json:"credentials_included"`
	RawCommandsIncluded           bool   `json:"raw_commands_included"`
	RawExecutablePathsIncluded    bool   `json:"raw_executable_paths_included"`
	BackendDetailsIncluded        bool   `json:"backend_details_included"`
	OmittedSensitiveEvidenceCount int    `json:"omitted_sensitive_evidence_count"`
}

func (plan Plan) SupportCaseTimelinePreview(history diagnostics.RunHistory, options SupportCaseTimelineOptions) (SupportCaseTimelinePreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}
	if options.Issue == "" {
		options.Issue = "portal-approval-required"
	}
	if options.TestType == "" {
		options.TestType = "smoke"
	}
	if options.Decision == "" {
		options.Decision = "deferred"
	}
	if history.ApplicationID != "" && history.ApplicationID != plan.ApplicationID {
		return SupportCaseTimelinePreview{}, errors.New("support case timeline history application id does not match recipe")
	}

	version, err := readRuntimeServiceBindingVersion(options.RuntimeRoot)
	if err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	recommendation, err := plan.AIDiagnosticRecommendationPreview(options.Issue, options.TestType)
	if err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	onboarding, err := plan.CompatibilityOnboardingChecklistPreview(CompatibilityOnboardingChecklistOptions{
		RuntimeRoot: options.RuntimeRoot,
		Issue:       options.Issue,
		TestType:    options.TestType,
	})
	if err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	graph, err := plan.KDEActionDependencyGraphPreview(options.Decision, options.FileURIs)
	if err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	journey, err := plan.KDEJourneyEvidencePreviewWithOptions(options.Decision, options.FileURIs, KDEJourneyEvidenceOptions{RuntimeRoot: options.RuntimeRoot})
	if err != nil {
		return SupportCaseTimelinePreview{}, err
	}

	events, mixedCount, blockedDiagnosticCount, redactedCount := supportCaseTimelineEvents(plan.ApplicationID, history, recommendation, onboarding, graph, journey, options)
	counts := supportCaseTimelineCounts(events, blockedDiagnosticCount, redactedCount)
	preview := SupportCaseTimelinePreview{
		Version:       version,
		SchemaVersion: "xnix.runtime.support_case_timeline.v1",
		RequestType:   "support-case-timeline-preview",
		TimelineType:  "redacted-runtime-support-case-timeline",
		Source:        "diagnostic-history+ai-diagnostic-recommendation-preview+compatibility-onboarding-checklist-preview+kde-action-dependency-graph-preview+kde-journey-evidence-preview",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetSupportCaseTimeline",
		ReadMethod:    "GetSupportCaseTimelinePreview",
		Application: SupportBundleApplication{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			RuntimeMode: plan.runtimeModeID(),
		},
		Events:                         events,
		EventIDs:                       supportCaseTimelineEventIDs(events),
		EventGroups:                    supportCaseTimelineGroups(events),
		EventCount:                     len(events),
		Counts:                         counts,
		RedactionSummary:               supportCaseTimelineRedaction(redactedCount),
		MalformedHistory:               options.MalformedHistory,
		MalformedEvidenceIDs:           safeEvidenceIDs(options.MalformedEvidenceIDs),
		MixedApplicationRecordCount:    mixedCount,
		BlockedDiagnosticEvidenceCount: blockedDiagnosticCount,
		RuntimeOwned:                   true,
		GoRuntimeBacked:                true,
		KDEPolicyOwner:                 false,
		UserVisible:                    true,
		ReviewOnly:                     true,
		TicketCreated:                  false,
		BundleExported:                 false,
		AIProviderCalled:               false,
		AIProviderCallEnabled:          false,
		RepairApplied:                  false,
		RepairExecuted:                 false,
		ActionExecuted:                 false,
		BackendProcessStarted:          false,
		FileContentRead:                false,
		FilePathsExposed:               false,
		StateRootPathExposed:           false,
		RawCommandExposed:              false,
		RawExecutableExposed:           false,
		BackendDetailsExposed:          false,
		HostRootModified:               false,
		NetworkRequired:                false,
		PrivilegedContainerRequired:    false,
		NextSafeReadOnlyChecks: []string{
			"review diagnostic run history metadata",
			"review repair recommendation categories",
			"review onboarding checklist gaps",
			"review KDE action dependency evidence",
			"review support bundle manifest preview",
		},
		BlockedActions: []string{
			"create support ticket from timeline preview",
			"export support bundle from timeline preview",
			"call AI provider from timeline preview",
			"apply repair from timeline preview",
			"execute KDE or Runtime action from timeline preview",
			"start compatibility service from timeline preview",
			"mutate host root during timeline preview",
		},
		DesktopSafeSummary: supportCaseTimelineSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "support case timeline preview"); err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	if err := safety.ValidatePayload("support case timeline preview", preview); err != nil {
		return SupportCaseTimelinePreview{}, err
	}
	return preview, nil
}

func supportCaseTimelineEvents(applicationID string, history diagnostics.RunHistory, recommendation AIDiagnosticRecommendationPreview, onboarding CompatibilityOnboardingChecklistPreview, graph KDEActionDependencyGraphPreview, journey KDEJourneyEvidencePreview, options SupportCaseTimelineOptions) ([]SupportCaseTimelineEvent, int, int, int) {
	events := []SupportCaseTimelineEvent{}
	mixedCount := 0
	blockedDiagnosticCount := 0
	redactedCount := len(options.PrivacySensitiveHints)

	for _, record := range history.Records {
		if record.ApplicationID != "" && record.ApplicationID != applicationID {
			mixedCount++
			continue
		}
		omittedSignals := 0
		for _, signalID := range record.FailingIDs {
			if crashHangPrivacySensitive(signalID) {
				omittedSignals++
			}
		}
		redactedCount += omittedSignals
		relatedEvidence := []string{}
		if record.RelativePath != "" && !crashHangPrivacySensitive(record.RelativePath) {
			relatedEvidence = append(relatedEvidence, record.RelativePath)
		}
		severity := "info"
		reviewRequired := false
		if record.Overall == diagnostics.OutcomeFail {
			severity = "warning"
			reviewRequired = true
		}
		if record.Overall == diagnostics.OutcomeBlocked {
			severity = "blocked"
			reviewRequired = true
			blockedDiagnosticCount++
		}
		events = append(events, supportCaseTimelineEvent(
			"diagnostic-run-"+record.RunID,
			"diagnostic-runs",
			severity,
			"diagnostic-run-history",
			relatedEvidence,
			[]string{record.RunID},
			supportCaseDiagnosticSummary(record, omittedSignals),
			"Review diagnostic run receipt metadata.",
			reviewRequired,
		))
	}

	if options.MalformedHistory {
		events = append(events, supportCaseTimelineEvent(
			"malformed-history",
			"malformed-history",
			"blocked",
			"diagnostic-run-history",
			safeEvidenceIDs(options.MalformedEvidenceIDs),
			nil,
			"Some diagnostic history metadata could not be parsed and was omitted from this timeline.",
			"Review diagnostic receipt integrity metadata.",
			true,
		))
	}
	if len(history.Records) == 0 && !options.MalformedHistory {
		events = append(events, supportCaseTimelineEvent(
			"diagnostic-run-none",
			"diagnostic-runs",
			"info",
			"diagnostic-run-history",
			nil,
			nil,
			"No diagnostic run history is available for this application.",
			"Run a fixture diagnostic preview before opening a support case.",
			false,
		))
	}

	for _, actionID := range graph.BlockedActions {
		if crashHangPrivacySensitive(actionID) {
			redactedCount++
			continue
		}
		events = append(events, supportCaseTimelineEvent(
			"blocked-action-"+actionID,
			"blocked-actions",
			"blocked",
			graph.RequestType,
			[]string{actionID},
			nil,
			"A KDE action remains blocked by Runtime evidence gates.",
			"Review the KDE action dependency graph.",
			true,
		))
	}

	for _, category := range supportCaseRecommendationCategories(recommendation) {
		if crashHangPrivacySensitive(category) {
			redactedCount++
			continue
		}
		events = append(events, supportCaseTimelineEvent(
			"repair-recommendation-"+category,
			"repair-recommendations",
			"review",
			recommendation.RequestType,
			[]string{category},
			nil,
			"Repair recommendation metadata is available for review.",
			"Review recommendation category metadata and approval gates.",
			true,
		))
	}

	for _, section := range onboarding.Sections {
		if section.State == "ready" {
			continue
		}
		events = append(events, supportCaseTimelineEvent(
			"onboarding-gap-"+section.ID,
			"onboarding-gaps",
			supportCaseSeverityForState(section.State),
			onboarding.RequestType,
			supportCaseSafeStrings(section.MissingEvidence, &redactedCount),
			nil,
			section.Summary,
			section.NextSafeReadOnlyCheck,
			true,
		))
	}

	for _, entry := range journey.EntryPoints {
		if entry.Ready {
			continue
		}
		events = append(events, supportCaseTimelineEvent(
			"kde-entrypoint-"+entry.ID,
			"kde-entrypoint-state",
			"review",
			journey.RequestType,
			supportCaseSafeStrings(entry.EvidenceIDs, &redactedCount),
			nil,
			entry.UserFacingSummary,
			entry.NextSafeReadOnlyCheck,
			true,
		))
	}

	sort.SliceStable(events, func(i, j int) bool { return events[i].ID < events[j].ID })
	return events, mixedCount, blockedDiagnosticCount, redactedCount
}

func supportCaseTimelineEvent(id string, group string, severity string, source string, evidence []string, runIDs []string, summary string, nextCheck string, reviewRequired bool) SupportCaseTimelineEvent {
	return SupportCaseTimelineEvent{
		ID:                    id,
		Group:                 group,
		Severity:              severity,
		SourceReadModel:       source,
		RelatedEvidenceIDs:    uniqueStrings(evidence),
		RelatedRunIDs:         uniqueStrings(runIDs),
		UserSafeSummary:       summary,
		NextSafeReadOnlyCheck: nextCheck,
		ReviewRequired:        reviewRequired,
		TicketCreated:         false,
		BundleExported:        false,
		AIProviderCalled:      false,
		RepairExecuted:        false,
		ActionExecuted:        false,
		BackendProcessStarted: false,
		HostRootModified:      false,
		FileContentRead:       false,
	}
}

func supportCaseTimelineCounts(events []SupportCaseTimelineEvent, blockedDiagnosticCount int, redactedCount int) SupportCaseTimelineCounts {
	var counts SupportCaseTimelineCounts
	counts.BlockedDiagnosticEvents = blockedDiagnosticCount
	counts.RedactedEvidenceOmissions = redactedCount
	for _, event := range events {
		switch event.Group {
		case "diagnostic-runs":
			counts.DiagnosticRunEvents++
		case "blocked-actions":
			counts.BlockedActionEvents++
		case "repair-recommendations":
			counts.RepairRecommendationEvents++
		case "onboarding-gaps":
			counts.OnboardingGapEvents++
		case "kde-entrypoint-state":
			counts.KDEEntrypointEvents++
		case "malformed-history":
			counts.MalformedHistoryEvents++
		}
		if event.ReviewRequired {
			counts.ReviewRequiredEvents++
		}
	}
	return counts
}

func supportCaseTimelineEventIDs(events []SupportCaseTimelineEvent) []string {
	ids := make([]string, 0, len(events))
	for _, event := range events {
		ids = append(ids, event.ID)
	}
	return ids
}

func supportCaseTimelineGroups(events []SupportCaseTimelineEvent) []string {
	groups := make([]string, 0, len(events))
	for _, event := range events {
		groups = append(groups, event.Group)
	}
	return uniqueStrings(groups)
}

func supportCaseTimelineRedaction(omitted int) SupportCaseTimelineRedaction {
	return SupportCaseTimelineRedaction{
		RedactionStatus:               "metadata-only-redacted-timeline",
		FileContentsIncluded:          false,
		HostPathsIncluded:             false,
		StateRootPathsIncluded:        false,
		EnvironmentVariablesIncluded:  false,
		UsernamesIncluded:             false,
		TokenShapedValuesIncluded:     false,
		CredentialsIncluded:           false,
		RawCommandsIncluded:           false,
		RawExecutablePathsIncluded:    false,
		BackendDetailsIncluded:        false,
		OmittedSensitiveEvidenceCount: omitted,
	}
}

func supportCaseTimelineSummary(counts SupportCaseTimelineCounts) string {
	if counts.MalformedHistoryEvents > 0 {
		return "Support timeline is blocked for diagnostic receipt review because some metadata is malformed."
	}
	if counts.DiagnosticRunEvents == 1 && counts.ReviewRequiredEvents == 0 {
		return "Support timeline has no diagnostic history yet and lists only safe next checks."
	}
	return "Support timeline joins diagnostic metadata, repair recommendations, onboarding gaps, and KDE state changes for review without creating tickets, exporting bundles, calling providers, applying repairs, starting services, or mutating the host root."
}

func supportCaseDiagnosticSummary(record diagnostics.RunHistoryRecord, omittedSignals int) string {
	state := string(record.Overall)
	if state == "" {
		state = "unknown"
	}
	if omittedSignals > 0 {
		return "Diagnostic run " + record.RunID + " ended as " + state + " with sensitive-looking signal metadata omitted."
	}
	return "Diagnostic run " + record.RunID + " ended as " + state + " and is available as metadata-only support evidence."
}

func supportCaseSeverityForState(state string) string {
	switch state {
	case "blocked", "missing-evidence", "not-yet-implemented":
		return "blocked"
	case "needs-review":
		return "review"
	default:
		return "info"
	}
}

func supportCaseRecommendationCategories(recommendation AIDiagnosticRecommendationPreview) []string {
	categories := []string{}
	for _, item := range recommendation.Recommendations {
		if item.ID != "" {
			categories = append(categories, item.ID)
		}
	}
	sort.Strings(categories)
	return uniqueStrings(categories)
}

func supportCaseSafeStrings(values []string, redactedCount *int) []string {
	safe := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if crashHangPrivacySensitive(value) {
			if redactedCount != nil {
				(*redactedCount)++
			}
			continue
		}
		safe = append(safe, value)
	}
	return uniqueStrings(safe)
}
