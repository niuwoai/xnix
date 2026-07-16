package appidentity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ApplicationUpgradeImpactPreview struct {
	SchemaVersion                 string                            `json:"schema_version"`
	RequestType                   string                            `json:"request_type"`
	ImpactType                    string                            `json:"impact_type"`
	Source                        string                            `json:"source"`
	Desktop                       string                            `json:"desktop"`
	RuntimeMethod                 string                            `json:"runtime_method"`
	ReadMethod                    string                            `json:"read_method"`
	Environment                   string                            `json:"environment"`
	ApplicationID                 string                            `json:"application_id"`
	CandidateApplicationID        string                            `json:"candidate_application_id"`
	DisplayName                   string                            `json:"display_name"`
	CurrentVersion                string                            `json:"current_version,omitempty"`
	CandidateVersion              string                            `json:"candidate_version,omitempty"`
	CandidateRelation             string                            `json:"candidate_relation"`
	Sections                      []ApplicationUpgradeImpactSection `json:"sections"`
	SectionIDs                    []string                          `json:"section_ids"`
	SectionCount                  int                               `json:"section_count"`
	Counts                        ApplicationUpgradeImpactCounts    `json:"counts"`
	OverallState                  string                            `json:"overall_state"`
	ReadyForReview                bool                              `json:"ready_for_review"`
	RuntimeOwned                  bool                              `json:"runtime_owned"`
	GoRuntimeBacked               bool                              `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                              `json:"kde_policy_owner"`
	UserVisible                   bool                              `json:"user_visible"`
	ReviewOnly                    bool                              `json:"review_only"`
	RecipeWritten                 bool                              `json:"recipe_written"`
	RegistryMigrated              bool                              `json:"registry_migrated"`
	RequestObjectsCreated         bool                              `json:"request_objects_created"`
	ArtifactsStaged               bool                              `json:"artifacts_staged"`
	ArtifactsDownloaded           bool                              `json:"artifacts_downloaded"`
	NetworkRequestCreated         bool                              `json:"network_request_created"`
	HostPackageManagerInvoked     bool                              `json:"host_package_manager_invoked"`
	SettingsPersisted             bool                              `json:"settings_persisted"`
	DesktopActivationStarted      bool                              `json:"desktop_activation_started"`
	DesktopFilesWritten           bool                              `json:"desktop_files_written"`
	BackendProcessStarted         bool                              `json:"backend_process_started"`
	LaunchEnabled                 bool                              `json:"launch_enabled"`
	ExecutionStarted              bool                              `json:"execution_started"`
	HostRootModified              bool                              `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                              `json:"privileged_container_required"`
	StateRootPathExposed          bool                              `json:"state_root_path_exposed"`
	RawExecutableExposed          bool                              `json:"raw_executable_exposed"`
	RawCommandExposed             bool                              `json:"raw_command_exposed"`
	BackendDetailsExposed         bool                              `json:"backend_details_exposed"`
	RawBackendImplementationShown bool                              `json:"raw_backend_implementation_shown"`
	BlockedActions                []string                          `json:"blocked_actions"`
	DesktopSafeSummary            string                            `json:"desktop_safe_summary"`
}

type ApplicationUpgradeImpactSection struct {
	ID                    string   `json:"id"`
	Title                 string   `json:"title"`
	State                 string   `json:"state"`
	CurrentEvidence       string   `json:"current_evidence"`
	CandidateEvidence     string   `json:"candidate_evidence"`
	Summary               string   `json:"summary"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	Changed               bool     `json:"changed"`
	UserReviewRequired    bool     `json:"user_review_required"`
	BlockingReasons       []string `json:"blocking_reasons"`
	RecipeWritten         bool     `json:"recipe_written"`
	RegistryMigrated      bool     `json:"registry_migrated"`
	ArtifactsStaged       bool     `json:"artifacts_staged"`
	SettingsPersisted     bool     `json:"settings_persisted"`
	DesktopFilesWritten   bool     `json:"desktop_files_written"`
	BackendProcessStarted bool     `json:"backend_process_started"`
	LaunchEnabled         bool     `json:"launch_enabled"`
	ExecutionStarted      bool     `json:"execution_started"`
	HostRootModified      bool     `json:"host_root_modified"`
	StateRootPathExposed  bool     `json:"state_root_path_exposed"`
	RawCommandExposed     bool     `json:"raw_command_exposed"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type ApplicationUpgradeImpactCounts struct {
	Total           int `json:"total"`
	Unchanged       int `json:"unchanged"`
	Changed         int `json:"changed"`
	NeedsReview     int `json:"needs_review"`
	MissingEvidence int `json:"missing_evidence"`
	Blocked         int `json:"blocked"`
	Unsupported     int `json:"unsupported"`
}

func NewApplicationUpgradeImpactPreview(currentRecipe Recipe, currentProvenance Provenance, candidateRecipe Recipe, candidateProvenance Provenance, environment string) (ApplicationUpgradeImpactPreview, error) {
	currentPlan, err := NewPlanWithProvenance(currentRecipe, currentProvenance)
	if err != nil {
		return ApplicationUpgradeImpactPreview{}, fmt.Errorf("current recipe: %w", err)
	}
	candidatePlan, err := NewPlanWithProvenance(candidateRecipe, candidateProvenance)
	if err != nil {
		return ApplicationUpgradeImpactPreview{}, fmt.Errorf("candidate recipe: %w", err)
	}
	return NewApplicationUpgradeImpactPreviewFromPlans(currentRecipe, currentPlan, candidateRecipe, candidatePlan, environment)
}

func NewApplicationUpgradeImpactPreviewFromPlans(currentRecipe Recipe, currentPlan Plan, candidateRecipe Recipe, candidatePlan Plan, environment string) (ApplicationUpgradeImpactPreview, error) {
	mode, err := normalizeActivationPreflightMode(environment)
	if err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	if err := currentPlan.ValidateSafeForDesktop(); err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	if err := candidatePlan.ValidateSafeForDesktop(); err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	currentInstall, err := currentPlan.CompatibilityInstallPlanPreview(mode)
	if err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	candidateInstall, err := candidatePlan.CompatibilityInstallPlanPreview(mode)
	if err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}

	relation := recipeVersionRelation(currentRecipe.Version, candidateRecipe.Version, currentPlan.StableIdentityDigest, candidatePlan.StableIdentityDigest)
	sections := applicationUpgradeImpactSections(currentRecipe, currentPlan, currentInstall, candidateRecipe, candidatePlan, candidateInstall, relation)
	counts := countApplicationUpgradeImpactSections(sections)
	sectionIDs := applicationUpgradeImpactSectionIDs(sections)
	preview := ApplicationUpgradeImpactPreview{
		SchemaVersion:                 "xnix.runtime.application_upgrade_impact.v1",
		RequestType:                   "application-upgrade-impact-preview",
		ImpactType:                    "review-only-application-upgrade-impact",
		Source:                        "registry+compatibility-install-preview+desktop-activation-evidence",
		Desktop:                       "KDE Plasma",
		RuntimeMethod:                 "GetApplicationUpgradeImpact",
		ReadMethod:                    "GetApplicationUpgradeImpactPreview",
		Environment:                   mode,
		ApplicationID:                 currentPlan.ApplicationID,
		CandidateApplicationID:        candidatePlan.ApplicationID,
		DisplayName:                   currentPlan.DisplayName,
		CurrentVersion:                currentRecipe.Version,
		CandidateVersion:              candidateRecipe.Version,
		CandidateRelation:             relation,
		Sections:                      sections,
		SectionIDs:                    sectionIDs,
		SectionCount:                  len(sections),
		Counts:                        counts,
		OverallState:                  applicationUpgradeImpactOverallState(counts),
		ReadyForReview:                true,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		UserVisible:                   true,
		ReviewOnly:                    true,
		RecipeWritten:                 false,
		RegistryMigrated:              false,
		RequestObjectsCreated:         false,
		ArtifactsStaged:               false,
		ArtifactsDownloaded:           false,
		NetworkRequestCreated:         false,
		HostPackageManagerInvoked:     false,
		SettingsPersisted:             false,
		DesktopActivationStarted:      false,
		DesktopFilesWritten:           false,
		BackendProcessStarted:         false,
		LaunchEnabled:                 false,
		ExecutionStarted:              false,
		HostRootModified:              false,
		PrivilegedContainerRequired:   false,
		StateRootPathExposed:          false,
		RawExecutableExposed:          false,
		RawCommandExposed:             false,
		BackendDetailsExposed:         false,
		RawBackendImplementationShown: false,
		BlockedActions: []string{
			"write candidate recipe from upgrade preview",
			"migrate recipe registry from upgrade preview",
			"stage artifacts from upgrade preview",
			"download artifacts from upgrade preview",
			"invoke host package managers from upgrade preview",
			"persist compatibility settings from upgrade preview",
			"start desktop activation from upgrade preview",
			"launch application from upgrade preview",
			"start compatibility services from upgrade preview",
			"mutate host root during upgrade preview",
		},
		DesktopSafeSummary: applicationUpgradeImpactSummary(counts, relation),
	}
	if err := validateApplicationUpgradeImpactPreview(preview); err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	if err := validateNoBackendTerms(preview, "application upgrade impact preview"); err != nil {
		return ApplicationUpgradeImpactPreview{}, err
	}
	return preview, nil
}

func applicationUpgradeImpactSections(currentRecipe Recipe, currentPlan Plan, currentInstall CompatibilityInstallPlanPreview, candidateRecipe Recipe, candidatePlan Plan, candidateInstall CompatibilityInstallPlanPreview, relation string) []ApplicationUpgradeImpactSection {
	recipeState := applicationUpgradeRecipeMetadataState(currentPlan, candidatePlan, relation)
	trustReasons := append([]string{}, candidateInstall.Readiness.RecipeTrustBlockingReasons...)
	if currentPlan.ApplicationID != candidatePlan.ApplicationID {
		trustReasons = append(trustReasons, "candidate application identity does not match the current application")
	}
	return []ApplicationUpgradeImpactSection{
		applicationUpgradeImpactSection(
			"recipe-metadata",
			"Recipe metadata",
			recipeState,
			recipeEvidence(currentRecipe, currentPlan),
			recipeEvidence(candidateRecipe, candidatePlan),
			"Runtime compares application identity, display metadata, supported document types, and recipe trust before any registry migration.",
			"Review the candidate recipe metadata and trust decision.",
			currentPlan.StableIdentityDigest != candidatePlan.StableIdentityDigest || currentRecipe.Version != candidateRecipe.Version,
			trustReasons,
		),
		applicationUpgradeImpactSection(
			"package-source",
			"Package source",
			applicationUpgradeMissingIfFalse(candidateInstall.Readiness.PackageSourceReady, currentInstall.SelectedStrategy != candidateInstall.SelectedStrategy),
			"current strategy "+currentInstall.SelectedStrategy,
			"candidate strategy "+candidateInstall.SelectedStrategy,
			"Runtime source selection remains planned; no package manager is invoked by the preview.",
			"Run package-source-preview for the candidate recipe.",
			currentInstall.SelectedStrategy != candidateInstall.SelectedStrategy,
			nil,
		),
		applicationUpgradeImpactSection(
			"artifact-digests",
			"Artifact digests",
			applicationUpgradeArtifactState(candidateInstall),
			applicationUpgradeArtifactEvidence(currentInstall),
			applicationUpgradeArtifactEvidence(candidateInstall),
			"Candidate artifacts must have signed manifests and staging receipts before upgrade activation can be trusted.",
			"Run artifact-manifest-preview and artifact-stage-record with local fixtures.",
			currentInstall.SelectedStrategy != candidateInstall.SelectedStrategy,
			candidateInstall.Readiness.ArtifactStageBlockingReasons,
		),
		applicationUpgradeImpactSection(
			"backend-profile",
			"Compatibility profile",
			applicationUpgradeChangedState(currentInstall.SelectedStrategy != candidateInstall.SelectedStrategy),
			"current profile "+currentInstall.SelectedStrategy,
			"candidate profile "+candidateInstall.SelectedStrategy,
			"Runtime compares only the user-facing compatibility profile and keeps implementation details hidden from KDE.",
			"Run backend-selection-preview for the candidate recipe.",
			currentInstall.SelectedStrategy != candidateInstall.SelectedStrategy,
			nil,
		),
		applicationUpgradeImpactSection(
			"portal-permissions",
			"Portal permissions",
			"needs-review",
			"current permissions require user review",
			"candidate permissions require user review",
			"File, document, device, and desktop resource access must remain mediated through Portal review.",
			"Run portal-request-preview or permission-review-preview for changed resource access.",
			false,
			nil,
		),
		applicationUpgradeImpactSection(
			"snapshot-requirements",
			"Snapshot requirements",
			"needs-review",
			"current rollback point is required before activation",
			"candidate rollback point is required before activation",
			"Upgrade activation must be paired with a Runtime-owned rollback point before writes are enabled.",
			"Run snapshot-plan-preview for the candidate upgrade reason.",
			false,
			nil,
		),
		applicationUpgradeImpactSection(
			"desktop-activation",
			"Desktop activation",
			applicationUpgradeMissingIfFalse(candidateInstall.DesktopActivationReady, currentPlan.DesktopFile != candidatePlan.DesktopFile),
			"current desktop entry "+currentPlan.DesktopFile,
			"candidate desktop entry "+candidatePlan.DesktopFile,
			"KDE launchers, file associations, tray state, notifications, Compatibility Center cards, and settings stay preview-only.",
			"Run desktop-activation-transaction-preview after upgrade evidence is complete.",
			currentPlan.DesktopFile != candidatePlan.DesktopFile,
			nil,
		),
		applicationUpgradeImpactSection(
			"diagnostics",
			"Diagnostics",
			"missing-evidence",
			"current diagnostic baseline is not attached",
			"candidate diagnostic baseline is not attached",
			"Diagnostics should compare recent failing signals and repair recommendations before accepting an upgrade.",
			"Run diagnostic-history-preview and support-bundle-manifest-preview for the candidate application.",
			false,
			nil,
		),
	}
}

func applicationUpgradeImpactSection(id string, title string, state string, currentEvidence string, candidateEvidence string, summary string, nextCheck string, changed bool, blockingReasons []string) ApplicationUpgradeImpactSection {
	blocking := uniqueSortedStrings(blockingReasons)
	if len(blocking) > 0 && state != "missing-evidence" {
		state = "blocked"
	}
	return ApplicationUpgradeImpactSection{
		ID:                    id,
		Title:                 title,
		State:                 state,
		CurrentEvidence:       currentEvidence,
		CandidateEvidence:     candidateEvidence,
		Summary:               summary,
		NextSafeReadOnlyCheck: nextCheck,
		Changed:               changed,
		UserReviewRequired:    state != "unchanged",
		BlockingReasons:       blocking,
		RecipeWritten:         false,
		RegistryMigrated:      false,
		ArtifactsStaged:       false,
		SettingsPersisted:     false,
		DesktopFilesWritten:   false,
		BackendProcessStarted: false,
		LaunchEnabled:         false,
		ExecutionStarted:      false,
		HostRootModified:      false,
		StateRootPathExposed:  false,
		RawCommandExposed:     false,
		BackendDetailsExposed: false,
	}
}

func applicationUpgradeRecipeMetadataState(currentPlan Plan, candidatePlan Plan, relation string) string {
	if currentPlan.ApplicationID != candidatePlan.ApplicationID {
		return "blocked"
	}
	switch relation {
	case "same-version":
		return "unchanged"
	case "candidate-older":
		return "needs-review"
	case "candidate-newer":
		return "changed"
	default:
		return "needs-review"
	}
}

func applicationUpgradeChangedState(changed bool) string {
	if changed {
		return "changed"
	}
	return "unchanged"
}

func applicationUpgradeMissingIfFalse(ready bool, changed bool) string {
	if !ready {
		return "missing-evidence"
	}
	return applicationUpgradeChangedState(changed)
}

func applicationUpgradeArtifactState(candidateInstall CompatibilityInstallPlanPreview) string {
	if len(candidateInstall.Readiness.ArtifactStageBlockingReasons) > 0 || !candidateInstall.Readiness.ArtifactStageReceiptReady || !candidateInstall.Readiness.RequiredArtifactsStaged {
		return "missing-evidence"
	}
	if !candidateInstall.Readiness.ArtifactManifestReady || !candidateInstall.Readiness.ArtifactSignatureVerified {
		return "needs-review"
	}
	return "unchanged"
}

func recipeEvidence(recipe Recipe, plan Plan) string {
	version := recipe.Version
	if version == "" {
		version = "unversioned"
	}
	return strings.Join([]string{
		"version " + version,
		"desktop digest " + shortDigest(plan.StableIdentityDigest),
		"trust " + recipeTrustLabel(plan),
		"document types " + strconv.Itoa(len(plan.MIMETypes)),
	}, "; ")
}

func applicationUpgradeArtifactEvidence(install CompatibilityInstallPlanPreview) string {
	return strings.Join([]string{
		"manifest ready " + boolWord(install.Readiness.ArtifactManifestReady),
		"signature verified " + boolWord(install.Readiness.ArtifactSignatureVerified),
		"stage receipt " + boolWord(install.Readiness.ArtifactStageReceiptReady),
		"required staged " + boolWord(install.Readiness.RequiredArtifactsStaged),
	}, "; ")
}

func recipeTrustLabel(plan Plan) string {
	if plan.RecipeSignatureStatus == "" {
		return "direct-file"
	}
	return plan.RecipeSignatureStatus
}

func boolWord(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

func shortDigest(value string) string {
	if len(value) <= 12 {
		return value
	}
	return value[:12]
}

func recipeVersionRelation(currentVersion string, candidateVersion string, currentDigest string, candidateDigest string) string {
	if currentVersion == "" || candidateVersion == "" {
		if currentDigest == candidateDigest {
			return "same-version"
		}
		return "changed-without-version"
	}
	comparison, err := compareSemanticVersions(currentVersion, candidateVersion)
	if err != nil {
		return "unknown"
	}
	switch {
	case comparison < 0:
		return "candidate-newer"
	case comparison > 0:
		return "candidate-older"
	default:
		if currentDigest == candidateDigest {
			return "same-version"
		}
		return "same-version-metadata-changed"
	}
}

func compareSemanticVersions(left string, right string) (int, error) {
	leftParts, err := semanticVersionParts(left)
	if err != nil {
		return 0, err
	}
	rightParts, err := semanticVersionParts(right)
	if err != nil {
		return 0, err
	}
	for index := range leftParts {
		if leftParts[index] < rightParts[index] {
			return -1, nil
		}
		if leftParts[index] > rightParts[index] {
			return 1, nil
		}
	}
	return 0, nil
}

func semanticVersionParts(version string) ([3]int, error) {
	if !versionPattern.MatchString(version) {
		return [3]int{}, errors.New("version must be semantic version format")
	}
	base := strings.SplitN(version, "-", 2)[0]
	rawParts := strings.Split(base, ".")
	if len(rawParts) != 3 {
		return [3]int{}, errors.New("version must include major, minor, and patch")
	}
	var parts [3]int
	for index, raw := range rawParts {
		value, err := strconv.Atoi(raw)
		if err != nil {
			return [3]int{}, err
		}
		parts[index] = value
	}
	return parts, nil
}

func countApplicationUpgradeImpactSections(sections []ApplicationUpgradeImpactSection) ApplicationUpgradeImpactCounts {
	counts := ApplicationUpgradeImpactCounts{Total: len(sections)}
	for _, section := range sections {
		switch section.State {
		case "unchanged":
			counts.Unchanged++
		case "changed":
			counts.Changed++
		case "needs-review":
			counts.NeedsReview++
		case "missing-evidence":
			counts.MissingEvidence++
		case "blocked":
			counts.Blocked++
		case "unsupported":
			counts.Unsupported++
		}
	}
	return counts
}

func applicationUpgradeImpactSectionIDs(sections []ApplicationUpgradeImpactSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func applicationUpgradeImpactOverallState(counts ApplicationUpgradeImpactCounts) string {
	if counts.Blocked > 0 {
		return "blocked"
	}
	if counts.Unsupported > 0 {
		return "unsupported"
	}
	if counts.MissingEvidence > 0 {
		return "missing-evidence"
	}
	if counts.NeedsReview > 0 {
		return "needs-review"
	}
	if counts.Changed > 0 {
		return "changed"
	}
	return "unchanged"
}

func applicationUpgradeImpactSummary(counts ApplicationUpgradeImpactCounts, relation string) string {
	if counts.Blocked > 0 {
		return "Runtime can show the candidate upgrade impact, but trust or identity blockers must be resolved before any upgrade action is enabled."
	}
	if counts.MissingEvidence > 0 {
		return "Runtime can review the candidate upgrade impact, but local artifact, activation, and diagnostic evidence is still missing."
	}
	if relation == "same-version" && counts.Changed == 0 {
		return "Candidate recipe appears unchanged; Runtime keeps all upgrade actions disabled until review explicitly approves a later transaction."
	}
	return "Runtime can review the candidate upgrade impact while recipe writes, downloads, activation, launch, and host mutation remain disabled."
}

func validateApplicationUpgradeImpactPreview(preview ApplicationUpgradeImpactPreview) error {
	for _, section := range preview.Sections {
		switch section.State {
		case "unchanged", "changed", "needs-review", "missing-evidence", "blocked", "unsupported":
		default:
			return fmt.Errorf("unsupported upgrade impact section state: %s", section.State)
		}
	}
	return nil
}
