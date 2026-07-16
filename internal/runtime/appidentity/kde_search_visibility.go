package appidentity

import (
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

// KDESearchVisibilityOptions carries the optional activation receipt root and
// the recipe-declared extensions used to derive searchable synonyms.
type KDESearchVisibilityOptions struct {
	ActivationRoot      string
	SupportedExtensions []string
}

type KDESearchVisibilityPlanPreview struct {
	SchemaVersion           string                    `json:"schema_version"`
	RequestType             string                    `json:"request_type"`
	PlanType                string                    `json:"plan_type"`
	Source                  string                    `json:"source"`
	Desktop                 string                    `json:"desktop"`
	RuntimeMethod           string                    `json:"runtime_method"`
	ReadMethod              string                    `json:"read_method"`
	ApplicationID           string                    `json:"application_id"`
	DisplayName             string                    `json:"display_name"`
	Icon                    string                    `json:"icon"`
	DesktopFile             string                    `json:"desktop_file"`
	OverallState            string                    `json:"overall_state"`
	SearchableLabels        []string                  `json:"searchable_labels"`
	Synonyms                []string                  `json:"synonyms"`
	MIMEEvidence            []string                  `json:"mime_evidence"`
	Rows                    []KDESearchVisibilityRow  `json:"rows"`
	RowIDs                  []string                  `json:"row_ids"`
	RowCount                int                       `json:"row_count"`
	Counts                  KDESearchVisibilityCounts `json:"counts"`
	ActivationReceiptRoot   bool                      `json:"activation_receipt_root"`
	ActivationReceiptBacked bool                      `json:"activation_receipt_backed"`
	ActivationReceiptPath   string                    `json:"activation_receipt_path,omitempty"`
	NextSafeReadOnlyChecks  []string                  `json:"next_safe_read_only_checks"`
	RuntimeOwned            bool                      `json:"runtime_owned"`
	GoRuntimeBacked         bool                      `json:"go_runtime_backed"`
	KDEPolicyOwner          bool                      `json:"kde_policy_owner"`
	UserVisible             bool                      `json:"user_visible"`
	ReviewOnly              bool                      `json:"review_only"`
	DesktopFilesWritten     bool                      `json:"desktop_files_written"`
	MIMEDefaultsWritten     bool                      `json:"mime_defaults_written"`
	KDECacheRefreshed       bool                      `json:"kde_cache_refreshed"`
	HostFilesIndexed        bool                      `json:"host_files_indexed"`
	SearchIndexPersisted    bool                      `json:"search_index_persisted"`
	BackendLaunchEnabled    bool                      `json:"backend_launch_enabled"`
	HostRootModified        bool                      `json:"host_root_modified"`
	BackendDetailsExposed   bool                      `json:"backend_details_exposed"`
	BlockedActions          []string                  `json:"blocked_actions"`
	DesktopSafeSummary      string                    `json:"desktop_safe_summary"`
}

type KDESearchVisibilityRow struct {
	ID                         string   `json:"id"`
	Surface                    string   `json:"surface"`
	Label                      string   `json:"label"`
	Synonyms                   []string `json:"synonyms"`
	MIMEEvidence               []string `json:"mime_evidence"`
	RuntimeMethod              string   `json:"runtime_method"`
	VisibilityState            string   `json:"visibility_state"`
	Searchable                 bool     `json:"searchable"`
	RequiresActivationReceipt  bool     `json:"requires_activation_receipt"`
	ActivationReceiptSatisfied bool     `json:"activation_receipt_satisfied"`
	RequiresMIMEAssociation    bool     `json:"requires_mime_association"`
	DisabledActionReason       string   `json:"disabled_action_reason,omitempty"`
	DesktopFilesWritten        bool     `json:"desktop_files_written"`
	MIMEDefaultsWritten        bool     `json:"mime_defaults_written"`
	HostRootModified           bool     `json:"host_root_modified"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
}

type KDESearchVisibilityCounts struct {
	TotalRows       int `json:"total_rows"`
	Visible         int `json:"visible"`
	Hidden          int `json:"hidden"`
	ReceiptRequired int `json:"receipt_required"`
	NoMIME          int `json:"no_mime"`
}

type kdeSearchVisibilitySurface struct {
	id              string
	surface         string
	runtimeMethod   string
	requiresReceipt bool
	requiresMIME    bool
}

func kdeSearchVisibilitySurfaces() []kdeSearchVisibilitySurface {
	return []kdeSearchVisibilitySurface{
		{"launcher", "KDE launcher", "GetDesktopEntryPlan", true, false},
		{"krunner", "KRunner search", "GetKRunnerQuery", true, false},
		{"file-association", "File association", "GetFileAssociationPlan", true, true},
		{"dolphin-action", "Dolphin action", "GetFileOpenPlan", true, true},
		{"compatibility-center", "Compatibility Center", "GetKDECenterPage", false, false},
		{"settings", "Settings search", "GetSettings", false, false},
		{"task-manager", "Task manager", "GetTaskManagerIdentityPlan", false, false},
	}
}

// KDESearchVisibilityPlanPreview explains how an application should appear in KDE
// search surfaces (launcher, KRunner, file association, Dolphin, Compatibility
// Center, settings, and task manager). It derives stable searchable labels and
// safe synonyms, records MIME/extension evidence, and reports disabled-action
// reasons and activation-receipt requirements without writing desktop files,
// refreshing caches, or indexing host files.
func (plan Plan) KDESearchVisibilityPlanPreview(options KDESearchVisibilityOptions) (KDESearchVisibilityPlanPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDESearchVisibilityPlanPreview{}, err
	}

	activationReceiptRoot := strings.TrimSpace(options.ActivationRoot) != ""
	activationReceiptBacked := false
	activationReceiptPath := ""
	if activationReceiptRoot {
		evidence, err := plan.DesktopActivationReceiptEvidence(options.ActivationRoot)
		if err != nil {
			return KDESearchVisibilityPlanPreview{}, err
		}
		activationReceiptBacked = evidence.SafeForKDE
		activationReceiptPath = evidence.ReceiptRelativePath
	}

	mimeEvidence := kdeSearchSafeSynonyms(plan.MIMETypes)
	synonyms := kdeSearchSynonyms(plan, options.SupportedExtensions)

	var rows []KDESearchVisibilityRow
	var counts KDESearchVisibilityCounts
	for _, surface := range kdeSearchVisibilitySurfaces() {
		row := KDESearchVisibilityRow{
			ID:                        surface.id,
			Surface:                   surface.surface,
			Label:                     plan.DisplayName,
			Synonyms:                  synonyms,
			RuntimeMethod:             surface.runtimeMethod,
			RequiresActivationReceipt: surface.requiresReceipt,
			RequiresMIMEAssociation:   surface.requiresMIME,
		}
		if surface.requiresMIME {
			row.MIMEEvidence = mimeEvidence
		}
		if surface.requiresReceipt {
			row.ActivationReceiptSatisfied = activationReceiptBacked
		}
		row.VisibilityState, row.DisabledActionReason = kdeSearchVisibilityState(plan.UserVisible, surface, mimeEvidence, activationReceiptBacked)
		row.Searchable = row.VisibilityState == "visible"
		counts.TotalRows++
		switch row.VisibilityState {
		case "visible":
			counts.Visible++
		case "hidden":
			counts.Hidden++
		case "receipt-required":
			counts.ReceiptRequired++
		case "no-mime-association":
			counts.NoMIME++
		}
		rows = append(rows, row)
	}

	preview := KDESearchVisibilityPlanPreview{
		SchemaVersion:           "xnix.runtime.kde_search_visibility.v1",
		RequestType:             "kde-search-visibility-plan-preview",
		PlanType:                "kde-search-visibility-plan",
		Source:                  "identity+krunner+file-association+kde-shell-surface+desktop-activation-manifest",
		Desktop:                 "KDE Plasma",
		RuntimeMethod:           "GetKDESearchVisibilityPlan",
		ReadMethod:              "GetKDESearchVisibilityPlanPreview",
		ApplicationID:           plan.ApplicationID,
		DisplayName:             plan.DisplayName,
		Icon:                    plan.Icon,
		DesktopFile:             plan.DesktopFile,
		SearchableLabels:        kdeSearchSafeSynonyms([]string{plan.DisplayName}),
		Synonyms:                synonyms,
		MIMEEvidence:            mimeEvidence,
		Rows:                    rows,
		RowIDs:                  kdeSearchRowIDs(rows),
		RowCount:                len(rows),
		Counts:                  counts,
		ActivationReceiptRoot:   activationReceiptRoot,
		ActivationReceiptBacked: activationReceiptBacked,
		ActivationReceiptPath:   activationReceiptPath,
		NextSafeReadOnlyChecks: []string{
			"review the launcher and KRunner read models",
			"review the file association and Dolphin action read models",
			"review the desktop activation manifest receipt requirement",
		},
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		UserVisible:           plan.UserVisible,
		ReviewOnly:            true,
		DesktopFilesWritten:   false,
		MIMEDefaultsWritten:   false,
		KDECacheRefreshed:     false,
		HostFilesIndexed:      false,
		SearchIndexPersisted:  false,
		BackendLaunchEnabled:  false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		BlockedActions: []string{
			"write a desktop file from the search visibility plan",
			"write a MIME default from the search visibility plan",
			"refresh the KDE service cache from the search visibility plan",
			"index host files from the search visibility plan",
			"persist a search index from the search visibility plan",
			"mutate host root during the search visibility plan",
		},
		OverallState:       kdeSearchOverallState(counts),
		DesktopSafeSummary: kdeSearchSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "KDE search visibility plan preview"); err != nil {
		return KDESearchVisibilityPlanPreview{}, err
	}
	if err := safety.ValidatePayload("KDE search visibility plan preview", preview); err != nil {
		return KDESearchVisibilityPlanPreview{}, err
	}
	return preview, nil
}

func kdeSearchVisibilityState(userVisible bool, surface kdeSearchVisibilitySurface, mimeEvidence []string, receiptBacked bool) (string, string) {
	if !userVisible {
		return "hidden", "application-hidden"
	}
	if surface.requiresMIME && len(mimeEvidence) == 0 {
		return "no-mime-association", "unsupported-mime"
	}
	if surface.requiresReceipt && !receiptBacked {
		return "receipt-required", "activation-receipt-required"
	}
	return "visible", ""
}

// kdeSearchSynonyms derives safe, deduplicated searchable synonyms from the
// display name tokens, launch verbs, and recipe extension names. Any token that
// fails the KDE-facing safety validator (for example a raw ".exe" extension) is
// dropped.
func kdeSearchSynonyms(plan Plan, extensions []string) []string {
	candidates := []string{}
	for _, token := range strings.Fields(strings.ToLower(plan.DisplayName)) {
		candidates = append(candidates, token)
	}
	for _, category := range plan.Categories {
		candidates = append(candidates, strings.ToLower(category))
	}
	candidates = append(candidates, "open", "launch", "run")
	for _, extension := range extensions {
		candidates = append(candidates, extension)
	}
	return kdeSearchSafeSynonyms(candidates)
}

// kdeSearchSafeSynonyms trims, lowercases where appropriate, drops empties and
// duplicates, and removes any value that is not safe for KDE-facing output.
func kdeSearchSafeSynonyms(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] || !safety.Safe(value) {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func kdeSearchRowIDs(rows []KDESearchVisibilityRow) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	return ids
}

func kdeSearchOverallState(counts KDESearchVisibilityCounts) string {
	switch {
	case counts.Hidden > 0:
		return "hidden"
	case counts.Visible == counts.TotalRows:
		return "fully-visible"
	case counts.ReceiptRequired > 0:
		return "receipt-required"
	default:
		return "partial-visibility"
	}
}

func kdeSearchSummary(counts KDESearchVisibilityCounts) string {
	switch kdeSearchOverallState(counts) {
	case "hidden":
		return "The application is hidden, so it would not appear in KDE search surfaces."
	case "fully-visible":
		return "The application would appear consistently across all KDE search surfaces once activation completes."
	case "receipt-required":
		return "Some KDE search surfaces need a verified activation receipt before the application appears; nothing was written."
	default:
		return "The application would appear in some KDE search surfaces; file surfaces need supported file types. Nothing was written."
	}
}
