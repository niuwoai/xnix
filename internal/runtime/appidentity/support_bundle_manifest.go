package appidentity

import (
	"errors"
	"sort"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

type SupportBundleManifestOptions struct {
	RuntimeRoot string
	Issue       string
	TestType    string
}

type SupportBundleManifestPreview struct {
	Version                        string                              `json:"version"`
	SchemaVersion                  string                              `json:"schema_version"`
	RequestType                    string                              `json:"request_type"`
	ManifestType                   string                              `json:"manifest_type"`
	Source                         string                              `json:"source"`
	Desktop                        string                              `json:"desktop"`
	RuntimeMethod                  string                              `json:"runtime_method"`
	ReadMethod                     string                              `json:"read_method"`
	Application                    SupportBundleApplication            `json:"application"`
	Sections                       []SupportBundleManifestSection      `json:"sections"`
	SectionIDs                     []string                            `json:"section_ids"`
	SectionCount                   int                                 `json:"section_count"`
	DiagnosticRunSummaries         []SupportBundleDiagnosticRunSummary `json:"diagnostic_run_summaries"`
	DiagnosticRunCount             int                                 `json:"diagnostic_run_count"`
	FailingSignalIDs               []string                            `json:"failing_signal_ids"`
	RepairRecommendationCategories []string                            `json:"repair_recommendation_categories"`
	PrivacyRedaction               SupportBundlePrivacyRedaction       `json:"privacy_redaction"`
	OmittedEvidence                SupportBundleOmittedEvidenceCounts  `json:"omitted_evidence"`
	RuntimeOwned                   bool                                `json:"runtime_owned"`
	GoRuntimeBacked                bool                                `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                                `json:"kde_policy_owner"`
	UserVisible                    bool                                `json:"user_visible"`
	OfflineOnly                    bool                                `json:"offline_only"`
	ArchiveCreated                 bool                                `json:"archive_created"`
	FileContentRead                bool                                `json:"file_content_read"`
	FilePathsExposed               bool                                `json:"file_paths_exposed"`
	AIProviderCalled               bool                                `json:"ai_provider_called"`
	AIProviderCallEnabled          bool                                `json:"ai_provider_call_enabled"`
	AutoRepairRequested            bool                                `json:"auto_repair_requested"`
	AutoRepairExecuted             bool                                `json:"auto_repair_executed"`
	BackendProcessStarted          bool                                `json:"backend_process_started"`
	HostRootModified               bool                                `json:"host_root_modified"`
	StateRootPathExposed           bool                                `json:"state_root_path_exposed"`
	RawExecutableExposed           bool                                `json:"raw_executable_exposed"`
	RawCommandExposed              bool                                `json:"raw_command_exposed"`
	BackendDetailsExposed          bool                                `json:"backend_details_exposed"`
	NetworkRequired                bool                                `json:"network_required"`
	PrivilegedContainerRequired    bool                                `json:"privileged_container_required"`
	BlockedActions                 []string                            `json:"blocked_actions"`
	DesktopSafeSummary             string                              `json:"desktop_safe_summary"`
}

type SupportBundleApplication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DesktopFile string `json:"desktop_file"`
	RuntimeMode string `json:"runtime_mode"`
}

type SupportBundleManifestSection struct {
	ID                    string `json:"id"`
	Title                 string `json:"title"`
	EvidenceSource        string `json:"evidence_source"`
	ReadModel             string `json:"read_model"`
	Summary               string `json:"summary"`
	Included              bool   `json:"included"`
	Redacted              bool   `json:"redacted"`
	RequiresUserApproval  bool   `json:"requires_user_approval"`
	SideEffectsEnabled    bool   `json:"side_effects_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	HostRootModified      bool   `json:"host_root_modified"`
}

type SupportBundleDiagnosticRunSummary struct {
	RunID                 string   `json:"run_id"`
	ApplicationID         string   `json:"application_id"`
	TestType              string   `json:"test_type"`
	Overall               string   `json:"overall"`
	FailingIDs            []string `json:"failing_ids"`
	RepairIssue           string   `json:"repair_issue,omitempty"`
	SnapshotRequired      bool     `json:"snapshot_required"`
	ReceiptReference      string   `json:"receipt_reference"`
	AutoRepairAllowed     bool     `json:"auto_repair_allowed"`
	BackendStarted        bool     `json:"backend_started"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	RepairExecuted        bool     `json:"repair_executed"`
	HostRootModified      bool     `json:"host_root_modified"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
	FileContentsIncluded  bool     `json:"file_contents_included"`
}

type SupportBundlePrivacyRedaction struct {
	RedactionStatus                       string `json:"redaction_status"`
	UserDocumentsIncluded                 bool   `json:"user_documents_included"`
	HostPathsIncluded                     bool   `json:"host_paths_included"`
	StateRootPathsIncluded                bool   `json:"state_root_paths_included"`
	EnvironmentVariablesIncluded          bool   `json:"environment_variables_included"`
	TokenShapedValuesIncluded             bool   `json:"token_shaped_values_included"`
	CommandShapedValuesIncluded           bool   `json:"command_shaped_values_included"`
	UsernamesIncluded                     bool   `json:"usernames_included"`
	SecretsIncluded                       bool   `json:"secrets_included"`
	NetworkCallsAllowed                   bool   `json:"network_calls_allowed"`
	RequiresUserApprovalForSensitiveTasks bool   `json:"requires_user_approval_for_sensitive_actions"`
}

type SupportBundleOmittedEvidenceCounts struct {
	FileContents         int `json:"file_contents"`
	HostPaths            int `json:"host_paths"`
	EnvironmentVariables int `json:"environment_variables"`
	TokenShapedValues    int `json:"token_shaped_values"`
	CommandShapedValues  int `json:"command_shaped_values"`
	Usernames            int `json:"usernames"`
}

func (plan Plan) SupportBundleManifestPreview(history diagnostics.RunHistory, options SupportBundleManifestOptions) (SupportBundleManifestPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return SupportBundleManifestPreview{}, err
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
	if history.ApplicationID != "" && history.ApplicationID != plan.ApplicationID {
		return SupportBundleManifestPreview{}, errors.New("support bundle manifest history application id does not match recipe")
	}

	version, err := readRuntimeServiceBindingVersion(options.RuntimeRoot)
	if err != nil {
		return SupportBundleManifestPreview{}, err
	}
	diagnosticPreview, err := NewDiagnosticHistoryPreview(history)
	if err != nil {
		return SupportBundleManifestPreview{}, err
	}
	diagnosticInput, err := plan.AIDiagnosticInputPreview(options.Issue, options.TestType)
	if err != nil {
		return SupportBundleManifestPreview{}, err
	}
	recommendation, err := plan.AIDiagnosticRecommendationPreview(options.Issue, options.TestType)
	if err != nil {
		return SupportBundleManifestPreview{}, err
	}

	runSummaries := supportBundleRunSummaries(history.Records)
	failingIDs := supportBundleFailingSignalIDs(history.Records)
	categories := supportBundleRecommendationCategories(recommendation.Recommendations)
	omitted := supportBundleOmittedEvidence(history, diagnosticInput, recommendation)
	sections := supportBundleManifestSections(version, diagnosticPreview, diagnosticInput, recommendation)

	preview := SupportBundleManifestPreview{
		Version:       version,
		SchemaVersion: "xnix.runtime.support_bundle_manifest.v1",
		RequestType:   "support-bundle-manifest-preview",
		ManifestType:  "redacted-offline-support-bundle-manifest",
		Source:        "diagnostic-history-preview+ai-diagnostic-input-preview+ai-diagnostic-recommendation-preview",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetSupportBundleManifest",
		ReadMethod:    "GetSupportBundleManifestPreview",
		Application: SupportBundleApplication{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			RuntimeMode: plan.runtimeModeID(),
		},
		Sections:                       sections,
		SectionIDs:                     supportBundleSectionIDs(sections),
		SectionCount:                   len(sections),
		DiagnosticRunSummaries:         runSummaries,
		DiagnosticRunCount:             len(runSummaries),
		FailingSignalIDs:               failingIDs,
		RepairRecommendationCategories: categories,
		PrivacyRedaction: SupportBundlePrivacyRedaction{
			RedactionStatus:                       "redacted-preview-only",
			UserDocumentsIncluded:                 false,
			HostPathsIncluded:                     false,
			StateRootPathsIncluded:                false,
			EnvironmentVariablesIncluded:          false,
			TokenShapedValuesIncluded:             false,
			CommandShapedValuesIncluded:           false,
			UsernamesIncluded:                     false,
			SecretsIncluded:                       false,
			NetworkCallsAllowed:                   false,
			RequiresUserApprovalForSensitiveTasks: true,
		},
		OmittedEvidence:             omitted,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		OfflineOnly:                 true,
		ArchiveCreated:              false,
		FileContentRead:             false,
		FilePathsExposed:            false,
		AIProviderCalled:            false,
		AIProviderCallEnabled:       false,
		AutoRepairRequested:         false,
		AutoRepairExecuted:          false,
		BackendProcessStarted:       false,
		HostRootModified:            false,
		StateRootPathExposed:        false,
		RawExecutableExposed:        false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BlockedActions: []string{
			"create support archive from manifest preview",
			"read private file contents for support manifest",
			"call AI provider for support manifest",
			"request automatic repair from support manifest",
			"start compatibility service from support manifest",
			"mutate host root during support manifest preview",
		},
		DesktopSafeSummary: "The support bundle manifest lists redacted diagnostic evidence that could be shared later without exporting files, reading private content, calling an AI provider, starting services, or mutating the host root.",
	}
	if err := validateNoBackendTerms(preview, "support bundle manifest preview"); err != nil {
		return SupportBundleManifestPreview{}, err
	}
	return preview, nil
}

func supportBundleManifestSections(version string, history DiagnosticHistoryPreview, input AIDiagnosticInputPreview, recommendation AIDiagnosticRecommendationPreview) []SupportBundleManifestSection {
	return []SupportBundleManifestSection{
		supportBundleSection("runtime-version", "Runtime version", "VERSION", "runtime-service-binding-preview", "Runtime version "+version+" would be included as plain metadata.", true, false, false),
		supportBundleSection("application-identity", "Application identity", "registry", "application-preview", "Application id, display name, icon id, desktop file id, and user-facing run mode would be included.", true, false, false),
		supportBundleSection("diagnostic-summaries", "Diagnostic summaries", history.Source, history.RequestType, history.DesktopSafeSummary, true, false, false),
		supportBundleSection("failing-signals", "Failing signal ids", history.Source, history.RequestType, "Only stable failing signal ids are included, not logs or file contents.", true, true, false),
		supportBundleSection("repair-recommendations", "Repair recommendation categories", recommendation.Source, recommendation.RequestType, recommendation.DesktopSafeSummary, true, true, len(recommendation.ApprovalRequiredActions) > 0),
		supportBundleSection("ai-diagnostic-boundary", "AI diagnostic boundary", input.Source, input.RequestType, input.DesktopSafeSummary, true, true, true),
		supportBundleSection("omitted-evidence", "Omitted evidence", "support-bundle-manifest-preview", "support-bundle-manifest-preview", "Private content, host-local identifiers, sensitive values, and command-shaped values are counted as omitted evidence.", true, true, true),
	}
}

func supportBundleSection(id string, title string, source string, readModel string, summary string, included bool, redacted bool, approval bool) SupportBundleManifestSection {
	return SupportBundleManifestSection{
		ID:                    id,
		Title:                 title,
		EvidenceSource:        source,
		ReadModel:             readModel,
		Summary:               summary,
		Included:              included,
		Redacted:              redacted,
		RequiresUserApproval:  approval,
		SideEffectsEnabled:    false,
		BackendDetailsExposed: false,
		HostRootModified:      false,
	}
}

func supportBundleSectionIDs(sections []SupportBundleManifestSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func supportBundleRunSummaries(records []diagnostics.RunHistoryRecord) []SupportBundleDiagnosticRunSummary {
	summaries := make([]SupportBundleDiagnosticRunSummary, 0, len(records))
	for _, record := range records {
		summaries = append(summaries, SupportBundleDiagnosticRunSummary{
			RunID:                 record.RunID,
			ApplicationID:         record.ApplicationID,
			TestType:              record.TestType,
			Overall:               string(record.Overall),
			FailingIDs:            append([]string(nil), record.FailingIDs...),
			RepairIssue:           record.RepairIssue,
			SnapshotRequired:      record.SnapshotRequired,
			ReceiptReference:      "diagnostic-run:" + record.RunID,
			AutoRepairAllowed:     false,
			BackendStarted:        false,
			AIProviderCalled:      false,
			RepairExecuted:        false,
			HostRootModified:      false,
			BackendDetailsExposed: false,
			FileContentsIncluded:  false,
		})
	}
	return summaries
}

func supportBundleFailingSignalIDs(records []diagnostics.RunHistoryRecord) []string {
	seen := map[string]bool{}
	var ids []string
	for _, record := range records {
		for _, id := range record.FailingIDs {
			if id == "" || seen[id] {
				continue
			}
			seen[id] = true
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	return ids
}

func supportBundleRecommendationCategories(recommendations []AIDiagnosticRecommendation) []string {
	seen := map[string]bool{}
	var categories []string
	for _, recommendation := range recommendations {
		category := recommendation.SourceSignal
		if category == "" {
			category = recommendation.ID
		}
		if seen[category] {
			continue
		}
		seen[category] = true
		categories = append(categories, category)
	}
	sort.Strings(categories)
	return categories
}

func supportBundleOmittedEvidence(history diagnostics.RunHistory, input AIDiagnosticInputPreview, recommendation AIDiagnosticRecommendationPreview) SupportBundleOmittedEvidenceCounts {
	recordCount := history.Counts.Total
	if recordCount < len(history.Records) {
		recordCount = len(history.Records)
	}
	return SupportBundleOmittedEvidenceCounts{
		FileContents:         recordCount + 1,
		HostPaths:            recordCount + 1,
		EnvironmentVariables: len(input.ContextSections) + 1,
		TokenShapedValues:    len(recommendation.ApprovalRequiredActions) + 1,
		CommandShapedValues:  len(recommendation.BlockedActions) + 1,
		Usernames:            1,
	}
}
