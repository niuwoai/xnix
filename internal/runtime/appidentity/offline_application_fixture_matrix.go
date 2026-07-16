package appidentity

import (
	"errors"
	"fmt"
	"sort"
)

type OfflineApplicationFixtureMatrixOptions struct {
	ShapeIDs    []string
	RuntimeRoot string
}

type OfflineApplicationFixtureMatrixPreview struct {
	SchemaVersion               string                                `json:"schema_version"`
	RequestType                 string                                `json:"request_type"`
	MatrixType                  string                                `json:"matrix_type"`
	Source                      string                                `json:"source"`
	Desktop                     string                                `json:"desktop"`
	RuntimeMethod               string                                `json:"runtime_method"`
	ReadMethod                  string                                `json:"read_method"`
	Rows                        []OfflineApplicationFixtureMatrixRow  `json:"rows"`
	ShapeIDs                    []string                              `json:"shape_ids"`
	RequiredShapeIDs            []string                              `json:"required_shape_ids"`
	MissingShapeIDs             []string                              `json:"missing_shape_ids"`
	Counts                      OfflineApplicationFixtureMatrixCounts `json:"counts"`
	MatrixStatus                string                                `json:"matrix_status"`
	ReadyForReview              bool                                  `json:"ready_for_review"`
	RuntimeOwned                bool                                  `json:"runtime_owned"`
	GoRuntimeBacked             bool                                  `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                  `json:"kde_policy_owner"`
	UserVisible                 bool                                  `json:"user_visible"`
	ReviewOnly                  bool                                  `json:"review_only"`
	OfflineDefault              bool                                  `json:"offline_default"`
	MatrixPersisted             bool                                  `json:"matrix_persisted"`
	NetworkFetchEnabled         bool                                  `json:"network_fetch_enabled"`
	PackageManagerInvoked       bool                                  `json:"package_manager_invoked"`
	ArtifactStagingEnabled      bool                                  `json:"artifact_staging_enabled"`
	BackendLaunchEnabled        bool                                  `json:"backend_launch_enabled"`
	DockerRequired              bool                                  `json:"docker_required"`
	QEMURequired                bool                                  `json:"qemu_required"`
	RequestObjectsCreated       bool                                  `json:"request_objects_created"`
	SettingsPersisted           bool                                  `json:"settings_persisted"`
	FileContentRead             bool                                  `json:"file_content_read"`
	StateRootPathExposed        bool                                  `json:"state_root_path_exposed"`
	RawExecutableExposed        bool                                  `json:"raw_executable_exposed"`
	RawCommandExposed           bool                                  `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                                  `json:"backend_details_exposed"`
	HostRootModified            bool                                  `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                  `json:"privileged_container_required"`
	BlockedUnsafeActions        []string                              `json:"blocked_unsafe_actions"`
	NextSafeReadOnlyChecks      []string                              `json:"next_safe_read_only_checks"`
	DesktopSafeSummary          string                                `json:"desktop_safe_summary"`
}

type OfflineApplicationFixtureMatrixRow struct {
	Position                   int      `json:"position"`
	ShapeID                    string   `json:"shape_id"`
	ShapeLabel                 string   `json:"shape_label"`
	ApplicationID              string   `json:"application_id"`
	ApplicationName            string   `json:"application_name"`
	Icon                       string   `json:"icon"`
	DesktopFile                string   `json:"desktop_file"`
	UserSafeRunMode            string   `json:"user_safe_run_mode"`
	MatrixState                string   `json:"matrix_state"`
	RecipeTrustState           string   `json:"recipe_trust_state"`
	ArtifactReadiness          string   `json:"artifact_readiness"`
	BackendProfileMapping      string   `json:"backend_profile_mapping"`
	PortalNeeds                []string `json:"portal_needs"`
	SnapshotReadiness          string   `json:"snapshot_readiness"`
	DiagnosticReadiness        string   `json:"diagnostic_readiness"`
	KDEJourneyCoverageState    string   `json:"kde_journey_coverage_state"`
	KDEJourneyEntryPointCount  int      `json:"kde_journey_entry_point_count"`
	KDEJourneyBlockedNodeCount int      `json:"kde_journey_blocked_node_count"`
	RequiredEvidenceIDs        []string `json:"required_evidence_ids"`
	MissingEvidenceIDs         []string `json:"missing_evidence_ids"`
	BlockedReasons             []string `json:"blocked_reasons"`
	NextSafeReadOnlyCheck      string   `json:"next_safe_read_only_check"`
	UserReviewRequired         bool     `json:"user_review_required"`
	UnsupportedShape           bool     `json:"unsupported_shape"`
	RuntimeOwned               bool     `json:"runtime_owned"`
	GoRuntimeBacked            bool     `json:"go_runtime_backed"`
	KDEPolicyOwner             bool     `json:"kde_policy_owner"`
	NetworkFetchEnabled        bool     `json:"network_fetch_enabled"`
	PackageManagerInvoked      bool     `json:"package_manager_invoked"`
	ArtifactStagingEnabled     bool     `json:"artifact_staging_enabled"`
	BackendProcessStarted      bool     `json:"backend_process_started"`
	LaunchEnabled              bool     `json:"launch_enabled"`
	ExecutionStarted           bool     `json:"execution_started"`
	RequestObjectCreated       bool     `json:"request_object_created"`
	SettingsPersisted          bool     `json:"settings_persisted"`
	FileContentRead            bool     `json:"file_content_read"`
	StateRootPathExposed       bool     `json:"state_root_path_exposed"`
	RawExecutableExposed       bool     `json:"raw_executable_exposed"`
	RawCommandExposed          bool     `json:"raw_command_exposed"`
	BackendDetailsExposed      bool     `json:"backend_details_exposed"`
	HostRootModified           bool     `json:"host_root_modified"`
}

type OfflineApplicationFixtureMatrixCounts struct {
	Total           int `json:"total"`
	Covered         int `json:"covered"`
	NeedsReview     int `json:"needs_review"`
	MissingEvidence int `json:"missing_evidence"`
	Unsupported     int `json:"unsupported"`
	MissingFixture  int `json:"missing_fixture"`
	OfflineOnly     int `json:"offline_only"`
}

type offlineApplicationFixtureDefinition struct {
	ShapeID     string
	ShapeLabel  string
	Recipe      Recipe
	PortalNeeds []string
	Unsupported bool
}

func NewOfflineApplicationFixtureMatrixPreview(options OfflineApplicationFixtureMatrixOptions) (OfflineApplicationFixtureMatrixPreview, error) {
	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}
	definitions := offlineApplicationFixtureDefinitions()
	selected, missing, err := selectOfflineApplicationFixtureDefinitions(definitions, options.ShapeIDs)
	if err != nil {
		return OfflineApplicationFixtureMatrixPreview{}, err
	}

	rows := make([]OfflineApplicationFixtureMatrixRow, 0, len(selected))
	for index, definition := range selected {
		row, err := offlineApplicationFixtureMatrixRow(index+1, definition, options.RuntimeRoot)
		if err != nil {
			return OfflineApplicationFixtureMatrixPreview{}, err
		}
		rows = append(rows, row)
	}
	counts := countOfflineApplicationFixtureMatrixRows(rows, len(missing))
	preview := OfflineApplicationFixtureMatrixPreview{
		SchemaVersion:               "xnix.runtime.offline_application_fixture_matrix.v1",
		RequestType:                 "offline-application-fixture-matrix-preview",
		MatrixType:                  "offline-cross-application-fixture-matrix",
		Source:                      "built-in-fixtures+runtime-read-models",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetOfflineApplicationFixtureMatrix",
		ReadMethod:                  "GetOfflineApplicationFixtureMatrixPreview",
		Rows:                        rows,
		ShapeIDs:                    offlineApplicationFixtureMatrixShapeIDs(rows),
		RequiredShapeIDs:            offlineApplicationFixtureShapeIDs(definitions),
		MissingShapeIDs:             missing,
		Counts:                      counts,
		MatrixStatus:                offlineApplicationFixtureMatrixStatus(counts),
		ReadyForReview:              len(rows) > 0,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		OfflineDefault:              true,
		MatrixPersisted:             false,
		NetworkFetchEnabled:         false,
		PackageManagerInvoked:       false,
		ArtifactStagingEnabled:      false,
		BackendLaunchEnabled:        false,
		DockerRequired:              false,
		QEMURequired:                false,
		RequestObjectsCreated:       false,
		SettingsPersisted:           false,
		FileContentRead:             false,
		StateRootPathExposed:        false,
		RawExecutableExposed:        false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BlockedUnsafeActions: []string{
			"download fixture artifacts",
			"invoke host package managers",
			"stage artifacts from fixture matrix",
			"start compatibility services from fixture matrix",
			"launch fixture applications",
			"run container or machine smoke checks from fixture matrix",
			"read user file contents while building fixture evidence",
			"mutate the host root during fixture review",
		},
		NextSafeReadOnlyChecks: []string{
			"refresh offline fixture matrix after adding a recipe shape",
			"compare fixture matrix with KDE-first presence smoke output",
			"include matrix results in merge-readiness packet review",
		},
		DesktopSafeSummary: "Offline fixture matrix covers representative application shapes with Runtime evidence while downloads, staging, launch, container smoke, machine smoke, file-content reads, and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "offline application fixture matrix preview"); err != nil {
		return OfflineApplicationFixtureMatrixPreview{}, err
	}
	return preview, nil
}

func offlineApplicationFixtureDefinitions() []offlineApplicationFixtureDefinition {
	return []offlineApplicationFixtureDefinition{
		{
			ShapeID:    "document-editor",
			ShapeLabel: "Document editor",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.document",
				Name:                "Fixture Document Editor",
				Version:             "1.0.0",
				Icon:                "x-office-document",
				Mode:                "automatic",
				SupportedExtensions: []string{".docx", ".xlsx"},
			},
			PortalNeeds: []string{"documents", "clipboard", "printing"},
		},
		{
			ShapeID:    "game",
			ShapeLabel: "Game",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.game",
				Name:                "Fixture Game",
				Version:             "1.0.0",
				Icon:                "applications-games",
				Mode:                "vm",
				SupportedExtensions: []string{".sav"},
			},
			PortalNeeds: []string{"network", "gamepad", "audio"},
		},
		{
			ShapeID:    "installer",
			ShapeLabel: "Installer",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.installer",
				Name:                "Fixture Installer",
				Version:             "1.0.0",
				Icon:                "system-software-install",
				Mode:                "automatic",
				SupportedExtensions: []string{".msi"},
			},
			PortalNeeds: []string{"downloads", "documents"},
		},
		{
			ShapeID:    "launcher",
			ShapeLabel: "Launcher",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.launcher",
				Name:                "Fixture Launcher",
				Version:             "1.0.0",
				Icon:                "application-x-executable",
				Mode:                "automatic",
				SupportedExtensions: []string{".url"},
			},
			PortalNeeds: []string{"network", "desktop-links"},
		},
		{
			ShapeID:    "network-heavy",
			ShapeLabel: "Network-heavy app",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.network",
				Name:                "Fixture Network App",
				Version:             "1.0.0",
				Icon:                "network-workgroup",
				Mode:                "automatic",
				SupportedExtensions: []string{".dat"},
			},
			PortalNeeds: []string{"network", "certificates"},
		},
		{
			ShapeID:    "tray-heavy",
			ShapeLabel: "Tray-heavy app",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.tray",
				Name:                "Fixture Tray App",
				Version:             "1.0.0",
				Icon:                "preferences-system-notifications",
				Mode:                "automatic",
				SupportedExtensions: []string{".cfg"},
			},
			PortalNeeds: []string{"notifications", "clipboard"},
		},
		{
			ShapeID:    "unsupported",
			ShapeLabel: "Unsupported app shape",
			Recipe: Recipe{
				ID:                  "org.xnix.fixture.unsupported",
				Name:                "Fixture Unsupported App",
				Version:             "1.0.0",
				Icon:                "dialog-warning",
				Mode:                "automatic",
				SupportedExtensions: []string{".bin"},
			},
			PortalNeeds: []string{"manual-review"},
			Unsupported: true,
		},
	}
}

func selectOfflineApplicationFixtureDefinitions(definitions []offlineApplicationFixtureDefinition, requested []string) ([]offlineApplicationFixtureDefinition, []string, error) {
	if len(requested) == 0 {
		return append([]offlineApplicationFixtureDefinition(nil), definitions...), []string{}, nil
	}
	byID := make(map[string]offlineApplicationFixtureDefinition, len(definitions))
	for _, definition := range definitions {
		byID[definition.ShapeID] = definition
	}
	seen := make(map[string]bool, len(requested))
	selected := make([]offlineApplicationFixtureDefinition, 0, len(requested))
	for _, id := range requested {
		if id == "" {
			return nil, nil, errors.New("fixture shape id must not be empty")
		}
		definition, ok := byID[id]
		if !ok {
			return nil, nil, fmt.Errorf("unknown fixture shape id: %s", id)
		}
		if seen[id] {
			continue
		}
		seen[id] = true
		selected = append(selected, definition)
	}
	missing := make([]string, 0)
	for _, definition := range definitions {
		if !seen[definition.ShapeID] {
			missing = append(missing, definition.ShapeID)
		}
	}
	return selected, missing, nil
}

func offlineApplicationFixtureMatrixRow(position int, definition offlineApplicationFixtureDefinition, runtimeRoot string) (OfflineApplicationFixtureMatrixRow, error) {
	plan, err := NewPlanWithProvenance(definition.Recipe, Provenance{
		Source:          "offline-fixture-matrix",
		RegistryName:    "built-in-offline-fixtures",
		DigestVerified:  true,
		SignatureStatus: "development-fixture",
	})
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}
	install, err := plan.CompatibilityInstallPlanPreview("development")
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}
	selection, err := plan.BackendSelectionPreview()
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}
	snapshot, err := NewSnapshotPlanPreview(plan.ApplicationID, "before-repair")
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}
	diagnostics, err := plan.AIDiagnosticInputPreview("portal-approval-required", "preflight")
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}
	journey, err := plan.KDEJourneyEvidencePreviewWithOptions("approved", []string{"file:///home/xnix/Documents/fixture.dat"}, KDEJourneyEvidenceOptions{RuntimeRoot: runtimeRoot})
	if err != nil {
		return OfflineApplicationFixtureMatrixRow{}, err
	}

	missingEvidence := offlineApplicationFixtureMissingEvidence(install, snapshot)
	blockedReasons := offlineApplicationFixtureBlockedReasons(definition, install, snapshot, diagnostics, journey)
	state := offlineApplicationFixtureRowState(definition, missingEvidence)
	return OfflineApplicationFixtureMatrixRow{
		Position:                   position,
		ShapeID:                    definition.ShapeID,
		ShapeLabel:                 definition.ShapeLabel,
		ApplicationID:              plan.ApplicationID,
		ApplicationName:            plan.DisplayName,
		Icon:                       plan.Icon,
		DesktopFile:                plan.DesktopFile,
		UserSafeRunMode:            offlineApplicationFixtureRunMode(selection.RecommendedProfileID),
		MatrixState:                state,
		RecipeTrustState:           offlineApplicationFixtureRecipeTrustState(install),
		ArtifactReadiness:          offlineApplicationFixtureArtifactReadiness(install),
		BackendProfileMapping:      selection.RecommendedProfileID,
		PortalNeeds:                append([]string(nil), definition.PortalNeeds...),
		SnapshotReadiness:          offlineApplicationFixtureSnapshotReadiness(snapshot),
		DiagnosticReadiness:        offlineApplicationFixtureDiagnosticReadiness(diagnostics),
		KDEJourneyCoverageState:    offlineApplicationFixtureJourneyState(journey),
		KDEJourneyEntryPointCount:  journey.EntryPointCount,
		KDEJourneyBlockedNodeCount: journey.BlockedReadinessNodeCount,
		RequiredEvidenceIDs: []string{
			"recipe-trust",
			"artifact-readiness",
			"backend-profile-mapping",
			"portal-needs",
			"snapshot-readiness",
			"diagnostic-readiness",
			"kde-journey-coverage",
		},
		MissingEvidenceIDs:     missingEvidence,
		BlockedReasons:         blockedReasons,
		NextSafeReadOnlyCheck:  offlineApplicationFixtureNextCheck(state),
		UserReviewRequired:     state != "covered-review-only",
		UnsupportedShape:       definition.Unsupported,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		NetworkFetchEnabled:    false,
		PackageManagerInvoked:  false,
		ArtifactStagingEnabled: false,
		BackendProcessStarted:  false,
		LaunchEnabled:          false,
		ExecutionStarted:       false,
		RequestObjectCreated:   false,
		SettingsPersisted:      false,
		FileContentRead:        false,
		StateRootPathExposed:   false,
		RawExecutableExposed:   false,
		RawCommandExposed:      false,
		BackendDetailsExposed:  false,
		HostRootModified:       false,
	}, nil
}

func offlineApplicationFixtureRunMode(profileID string) string {
	if profileID == "isolated-compatibility" {
		return "isolated compatibility"
	}
	return "local compatibility"
}

func offlineApplicationFixtureRecipeTrustState(install CompatibilityInstallPlanPreview) string {
	if install.Readiness.RecipeInstallAllowed {
		return "development-fixture-trusted"
	}
	return "needs-review"
}

func offlineApplicationFixtureArtifactReadiness(install CompatibilityInstallPlanPreview) string {
	if install.Readiness.ArtifactStageReceiptReady && install.Readiness.RequiredArtifactsStaged {
		return "local-fixture-ready"
	}
	return "missing-local-stage-receipt"
}

func offlineApplicationFixtureSnapshotReadiness(snapshot SnapshotPlanPreview) string {
	if snapshot.EnabledByDefault && !snapshot.SnapshotCreated {
		return "planned-receipt-required"
	}
	return "needs-review"
}

func offlineApplicationFixtureDiagnosticReadiness(diagnostics AIDiagnosticInputPreview) string {
	if diagnostics.SafeForAIDiagnostics && !diagnostics.AIProviderCalled && !diagnostics.FileContentRead {
		return "metadata-only-ready"
	}
	return "blocked"
}

func offlineApplicationFixtureJourneyState(journey KDEJourneyEvidencePreview) string {
	if journey.EntryPointCount == 7 && journey.Agreement.AppIDConsistent && journey.Agreement.DisabledActionStateShared {
		return "seven-entrypoints-covered"
	}
	return "missing-journey-evidence"
}

func offlineApplicationFixtureMissingEvidence(install CompatibilityInstallPlanPreview, snapshot SnapshotPlanPreview) []string {
	var missing []string
	if !install.Readiness.ArtifactStageReceiptReady || !install.Readiness.RequiredArtifactsStaged {
		missing = append(missing, "artifact-stage-receipt")
	}
	if !snapshot.SnapshotCreated {
		missing = append(missing, "snapshot-receipt")
	}
	return uniqueSortedStrings(missing)
}

func offlineApplicationFixtureBlockedReasons(definition offlineApplicationFixtureDefinition, install CompatibilityInstallPlanPreview, snapshot SnapshotPlanPreview, diagnostics AIDiagnosticInputPreview, journey KDEJourneyEvidencePreview) []string {
	var reasons []string
	if definition.Unsupported {
		reasons = append(reasons, "application shape requires explicit unsupported-state handling")
	}
	if !install.Readiness.ArtifactStageReceiptReady || !install.Readiness.RequiredArtifactsStaged {
		reasons = append(reasons, "local artifact staging receipt is missing")
	}
	if !snapshot.SnapshotCreated {
		reasons = append(reasons, "restore-point receipt is missing")
	}
	if !diagnostics.SafeForAIDiagnostics || diagnostics.AIProviderCalled || diagnostics.FileContentRead {
		reasons = append(reasons, "diagnostic privacy boundary is not ready")
	}
	if journey.EntryPointCount != 7 {
		reasons = append(reasons, "KDE journey does not cover all seven entry points")
	}
	return uniqueSortedStrings(reasons)
}

func offlineApplicationFixtureRowState(definition offlineApplicationFixtureDefinition, missingEvidence []string) string {
	if definition.Unsupported {
		return "blocked-unsupported"
	}
	if len(missingEvidence) > 0 {
		return "missing-evidence"
	}
	return "covered-review-only"
}

func offlineApplicationFixtureNextCheck(state string) string {
	switch state {
	case "blocked-unsupported":
		return "review unsupported-state policy card"
	case "missing-evidence":
		return "review missing local receipts without staging artifacts"
	default:
		return "refresh KDE journey evidence"
	}
}

func countOfflineApplicationFixtureMatrixRows(rows []OfflineApplicationFixtureMatrixRow, missingFixtureCount int) OfflineApplicationFixtureMatrixCounts {
	counts := OfflineApplicationFixtureMatrixCounts{
		Total:          len(rows),
		MissingFixture: missingFixtureCount,
		OfflineOnly:    len(rows),
	}
	for _, row := range rows {
		switch row.MatrixState {
		case "covered-review-only":
			counts.Covered++
		case "blocked-unsupported":
			counts.Unsupported++
		case "missing-evidence":
			counts.MissingEvidence++
		default:
			counts.NeedsReview++
		}
		if row.UserReviewRequired {
			counts.NeedsReview++
		}
	}
	return counts
}

func offlineApplicationFixtureMatrixStatus(counts OfflineApplicationFixtureMatrixCounts) string {
	if counts.MissingFixture > 0 {
		return "missing-fixtures"
	}
	if counts.Unsupported > 0 {
		return "covered-with-unsupported-shape"
	}
	if counts.MissingEvidence > 0 {
		return "covered-with-missing-evidence"
	}
	return "covered-review-only"
}

func offlineApplicationFixtureShapeIDs(definitions []offlineApplicationFixtureDefinition) []string {
	ids := make([]string, 0, len(definitions))
	for _, definition := range definitions {
		ids = append(ids, definition.ShapeID)
	}
	sort.Strings(ids)
	return ids
}

func offlineApplicationFixtureMatrixShapeIDs(rows []OfflineApplicationFixtureMatrixRow) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ShapeID)
	}
	sort.Strings(ids)
	return ids
}
