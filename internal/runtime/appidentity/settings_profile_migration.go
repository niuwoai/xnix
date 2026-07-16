package appidentity

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const currentSettingsProfileSchemaVersion = "xnix.runtime.settings.v1"

type SettingsProfileMigrationOptions struct {
	FromSchemaVersion string
	ToSchemaVersion   string
	BlockedSettingIDs []string
	ActivationRoot    string
}

type SettingsProfileMigrationPreview struct {
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	PreviewType                 string                         `json:"preview_type"`
	Source                      string                         `json:"source"`
	Desktop                     string                         `json:"desktop"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	ApplicationID               string                         `json:"application_id"`
	DisplayName                 string                         `json:"display_name"`
	Icon                        string                         `json:"icon"`
	DesktopFile                 string                         `json:"desktop_file"`
	FromSchemaVersion           string                         `json:"from_schema_version"`
	ToSchemaVersion             string                         `json:"to_schema_version"`
	SchemaState                 string                         `json:"schema_state"`
	CurrentSchemaSupported      bool                           `json:"current_schema_supported"`
	TargetSchemaSupported       bool                           `json:"target_schema_supported"`
	MigrationRequired           bool                           `json:"migration_required"`
	UserReviewRequired          bool                           `json:"user_review_required"`
	Rows                        []SettingsProfileMigrationRow  `json:"rows"`
	RowIDs                      []string                       `json:"row_ids"`
	RowCount                    int                            `json:"row_count"`
	Counts                      SettingsProfileMigrationCounts `json:"counts"`
	BlockedSettingIDs           []string                       `json:"blocked_setting_ids"`
	RollbackNotes               []string                       `json:"rollback_notes"`
	NextSafeReadOnlyChecks      []string                       `json:"next_safe_read_only_checks"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	UserVisible                 bool                           `json:"user_visible"`
	ReviewOnly                  bool                           `json:"review_only"`
	SettingsPersisted           bool                           `json:"settings_persisted"`
	SettingsPersistenceEnabled  bool                           `json:"settings_persistence_enabled"`
	ResourceGrantEnabled        bool                           `json:"resource_grant_enabled"`
	RealPortalCallEnabled       bool                           `json:"real_portal_call_enabled"`
	BackendLaunchEnabled        bool                           `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                           `json:"backend_process_started"`
	HostRootModified            bool                           `json:"host_root_modified"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	RawCommandExposed           bool                           `json:"raw_command_exposed"`
	RawExecutableExposed        bool                           `json:"raw_executable_exposed"`
	FileContentRead             bool                           `json:"file_content_read"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	NetworkRequired             bool                           `json:"network_required"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	BlockedActions              []string                       `json:"blocked_actions"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type SettingsProfileMigrationRow struct {
	ID                    string   `json:"id"`
	SectionID             string   `json:"section_id"`
	FieldID               string   `json:"field_id"`
	Title                 string   `json:"title"`
	CurrentValue          string   `json:"current_value"`
	TargetDefaultValue    string   `json:"target_default_value"`
	MigrationState        string   `json:"migration_state"`
	DefaultingRule        string   `json:"defaulting_rule"`
	UserReviewRequired    bool     `json:"user_review_required"`
	Blocked               bool     `json:"blocked"`
	BlockedReasons        []string `json:"blocked_reasons"`
	RollbackNote          string   `json:"rollback_note"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	SettingsPersisted     bool     `json:"settings_persisted"`
	ResourceGrantEnabled  bool     `json:"resource_grant_enabled"`
	RealPortalCallEnabled bool     `json:"real_portal_call_enabled"`
	BackendLaunchEnabled  bool     `json:"backend_launch_enabled"`
	HostRootModified      bool     `json:"host_root_modified"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type SettingsProfileMigrationCounts struct {
	TotalRows      int `json:"total_rows"`
	Unchanged      int `json:"unchanged"`
	Defaulted      int `json:"defaulted"`
	ReviewRequired int `json:"review_required"`
	Blocked        int `json:"blocked"`
	FutureSchema   int `json:"future_schema"`
}

type settingsProfileMigrationDefinition struct {
	id             string
	sectionID      string
	fieldID        string
	title          string
	targetDefault  string
	reviewRequired bool
	defaultRule    string
	nextCheck      string
}

func (plan Plan) SettingsProfileMigrationPreview(options SettingsProfileMigrationOptions) (SettingsProfileMigrationPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile, options.FromSchemaVersion, options.ToSchemaVersion} {
		if !singleLineOrBlank(value) {
			return SettingsProfileMigrationPreview{}, errors.New("settings profile migration preview requires single-line fields")
		}
	}
	if options.FromSchemaVersion == "" {
		options.FromSchemaVersion = "xnix.runtime.settings.v0"
	}
	if options.ToSchemaVersion == "" {
		options.ToSchemaVersion = currentSettingsProfileSchemaVersion
	}
	fromNumber, err := settingsProfileSchemaNumber(options.FromSchemaVersion)
	if err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	toNumber, err := settingsProfileSchemaNumber(options.ToSchemaVersion)
	if err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	if options.ToSchemaVersion != currentSettingsProfileSchemaVersion {
		return SettingsProfileMigrationPreview{}, fmt.Errorf("unsupported target settings schema: %s", options.ToSchemaVersion)
	}

	settings, err := plan.SettingsPreviewWithOptions(SettingsOptions{ActivationRoot: options.ActivationRoot})
	if err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	blockedIDs, err := settingsProfileMigrationBlockedIDs(options.BlockedSettingIDs)
	if err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	rows := settingsProfileMigrationRows(settings, fromNumber, toNumber, blockedIDs)
	counts := countSettingsProfileMigrationRows(rows)
	schemaState := settingsProfileMigrationSchemaState(fromNumber, toNumber)

	preview := SettingsProfileMigrationPreview{
		SchemaVersion:               "xnix.runtime.settings_profile_migration.v1",
		RequestType:                 "settings-profile-migration-preview",
		PreviewType:                 "compatibility-settings-profile-migration-preview",
		Source:                      "settings-preview+settings-change-preview+compatibility-onboarding-checklist+application-readiness-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetCompatibilitySettingsProfileMigration",
		ReadMethod:                  "GetCompatibilitySettingsProfileMigrationPreview",
		ApplicationID:               plan.ApplicationID,
		DisplayName:                 plan.DisplayName,
		Icon:                        plan.Icon,
		DesktopFile:                 plan.DesktopFile,
		FromSchemaVersion:           options.FromSchemaVersion,
		ToSchemaVersion:             options.ToSchemaVersion,
		SchemaState:                 schemaState,
		CurrentSchemaSupported:      fromNumber <= toNumber,
		TargetSchemaSupported:       true,
		MigrationRequired:           fromNumber != toNumber,
		UserReviewRequired:          counts.ReviewRequired > 0 || counts.Blocked > 0 || counts.FutureSchema > 0,
		Rows:                        rows,
		RowIDs:                      settingsProfileMigrationRowIDs(rows),
		RowCount:                    len(rows),
		Counts:                      counts,
		BlockedSettingIDs:           sortedSettingsProfileMigrationBlockedIDs(blockedIDs),
		RollbackNotes:               settingsProfileMigrationRollbackNotes(schemaState),
		NextSafeReadOnlyChecks:      []string{"settings-preview", "settings-change-preview", "mode-switch-preview", "portal-permission-renewal-preview", "runtime-policy-explanation-cards-preview"},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		SettingsPersisted:           false,
		SettingsPersistenceEnabled:  false,
		ResourceGrantEnabled:        false,
		RealPortalCallEnabled:       false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		HostRootModified:            false,
		StateRootPathExposed:        false,
		RawCommandExposed:           false,
		RawExecutableExposed:        false,
		FileContentRead:             false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BlockedActions: []string{
			"persist migrated settings from preview",
			"grant desktop resources during settings migration",
			"call real Portal transport during settings migration",
			"start compatibility backends during settings migration",
			"expose backend implementation settings to KDE",
			"mutate host root during settings migration inspection",
		},
		DesktopSafeSummary: "Settings profile migration preview explains schema defaults, review gates, blocked settings, and rollback notes without persisting settings or changing host state.",
	}
	if err := validateNoBackendTerms(preview, "settings profile migration preview"); err != nil {
		return SettingsProfileMigrationPreview{}, err
	}
	return preview, nil
}

func settingsProfileMigrationDefinitions() []settingsProfileMigrationDefinition {
	return []settingsProfileMigrationDefinition{
		{"run-mode.mode", "run-mode", "mode", "Run mode", "automatic", true, "default to Automatic so the Runtime can choose a safe compatibility route", "mode-switch-preview"},
		{"run-mode.preference", "run-mode", "preference", "Priority", "compatibility", true, "default to Compatibility to prefer correctness over speed", "settings-change-preview"},
		{"resource-access.documents", "resource-access", "documents", "Documents access", "ask", true, "default to Ask so KDE can request user review before document access", "portal-permission-renewal-preview"},
		{"resource-access.downloads", "resource-access", "downloads", "Downloads access", "ask", true, "default to Ask so KDE can request user review before downloads access", "portal-permission-renewal-preview"},
		{"devices.camera", "devices", "camera", "Camera access", "deny", true, "default to Deny because camera access is sensitive", "portal-permission-renewal-preview"},
		{"network.network", "network", "network", "Network access", "allow", false, "default to Allow for compatibility metadata checks while execution remains gated", "settings-change-preview"},
		{"snapshots.snapshots", "snapshots", "snapshots", "Snapshots", "enabled", false, "default to Enabled so risky compatibility changes can require restore-point planning", "snapshot-plan-preview"},
		{"diagnostics.privacy", "diagnostics", "privacy", "Diagnostics privacy", "metadata-only", true, "default to Metadata only so diagnostics avoid private file contents", "ai-diagnostic-input-preview"},
	}
}

func settingsProfileMigrationRows(settings SettingsPreview, fromNumber int, toNumber int, blockedIDs map[string]bool) []SettingsProfileMigrationRow {
	current := settingsProfileMigrationCurrentValues(settings)
	rows := make([]SettingsProfileMigrationRow, 0, len(settingsProfileMigrationDefinitions()))
	for _, definition := range settingsProfileMigrationDefinitions() {
		currentValue := current[definition.id]
		if currentValue == "" {
			currentValue = "not-present"
		}
		state := "unchanged"
		if fromNumber < toNumber {
			state = "defaulted"
			currentValue = "not-present"
		}
		if fromNumber > toNumber {
			state = "future-schema"
		}
		blockedReasons := []string{}
		if definition.reviewRequired && state == "defaulted" {
			state = "review-required"
		}
		if blockedIDs[definition.id] {
			state = "blocked"
			blockedReasons = append(blockedReasons, "setting is blocked by migration policy")
		}
		if fromNumber > toNumber {
			blockedReasons = append(blockedReasons, "source schema is newer than the supported target schema")
		}
		rows = append(rows, SettingsProfileMigrationRow{
			ID:                    definition.id,
			SectionID:             definition.sectionID,
			FieldID:               definition.fieldID,
			Title:                 definition.title,
			CurrentValue:          currentValue,
			TargetDefaultValue:    definition.targetDefault,
			MigrationState:        state,
			DefaultingRule:        definition.defaultRule,
			UserReviewRequired:    definition.reviewRequired || state == "future-schema" || state == "blocked",
			Blocked:               state == "blocked" || state == "future-schema",
			BlockedReasons:        blockedReasons,
			RollbackNote:          "Because this is a preview, rollback means keeping the existing settings profile unchanged.",
			NextSafeReadOnlyCheck: definition.nextCheck,
			SettingsPersisted:     false,
			ResourceGrantEnabled:  false,
			RealPortalCallEnabled: false,
			BackendLaunchEnabled:  false,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		})
	}
	return rows
}

func settingsProfileMigrationCurrentValues(settings SettingsPreview) map[string]string {
	values := map[string]string{}
	for _, section := range settings.Sections {
		for _, field := range section.Fields {
			values[section.ID+"."+field.ID] = field.Value
		}
	}
	return values
}

func settingsProfileMigrationBlockedIDs(values []string) (map[string]bool, error) {
	known := map[string]bool{}
	for _, definition := range settingsProfileMigrationDefinitions() {
		known[definition.id] = true
	}
	blocked := map[string]bool{}
	for _, value := range values {
		for _, part := range strings.Split(value, ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			if !known[part] {
				return nil, fmt.Errorf("unknown settings migration field: %s", part)
			}
			blocked[part] = true
		}
	}
	return blocked, nil
}

func settingsProfileSchemaNumber(value string) (int, error) {
	const prefix = "xnix.runtime.settings.v"
	if !strings.HasPrefix(value, prefix) {
		return 0, fmt.Errorf("unsupported settings schema format: %s", value)
	}
	number, err := strconv.Atoi(strings.TrimPrefix(value, prefix))
	if err != nil {
		return 0, fmt.Errorf("unsupported settings schema format: %s", value)
	}
	return number, nil
}

func settingsProfileMigrationSchemaState(fromNumber int, toNumber int) string {
	switch {
	case fromNumber < toNumber:
		return "old-schema"
	case fromNumber == toNumber:
		return "current-schema"
	default:
		return "future-schema"
	}
}

func countSettingsProfileMigrationRows(rows []SettingsProfileMigrationRow) SettingsProfileMigrationCounts {
	var counts SettingsProfileMigrationCounts
	counts.TotalRows = len(rows)
	for _, row := range rows {
		switch row.MigrationState {
		case "unchanged":
			counts.Unchanged++
		case "defaulted":
			counts.Defaulted++
		case "review-required":
			counts.ReviewRequired++
		case "blocked":
			counts.Blocked++
		case "future-schema":
			counts.FutureSchema++
		}
	}
	return counts
}

func settingsProfileMigrationRowIDs(rows []SettingsProfileMigrationRow) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	return ids
}

func sortedSettingsProfileMigrationBlockedIDs(blocked map[string]bool) []string {
	ids := make([]string, 0, len(blocked))
	for id := range blocked {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func settingsProfileMigrationRollbackNotes(schemaState string) []string {
	notes := []string{
		"Do not persist the migrated profile until the Runtime settings write gate exists.",
		"Keep the existing settings profile as the rollback source of truth.",
	}
	if schemaState == "future-schema" {
		notes = append(notes, "A newer source schema must be reviewed by a future Runtime before migration can proceed.")
	}
	return notes
}
