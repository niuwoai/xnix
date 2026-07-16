package appidentity

type CompatibilityOnboardingChecklistOptions struct {
	RuntimeRoot     string
	StateRoot       string
	ArtifactReceipt bool
	PortalOperation string
	SnapshotReason  string
	Issue           string
	TestType        string
}

type CompatibilityOnboardingChecklistPreview struct {
	Version                     string                            `json:"version"`
	SchemaVersion               string                            `json:"schema_version"`
	RequestType                 string                            `json:"request_type"`
	ChecklistType               string                            `json:"checklist_type"`
	Source                      string                            `json:"source"`
	Desktop                     string                            `json:"desktop"`
	RuntimeMethod               string                            `json:"runtime_method"`
	ReadMethod                  string                            `json:"read_method"`
	Application                 CompatibilityOnboardingIdentity   `json:"application"`
	Sections                    []CompatibilityOnboardingSection  `json:"sections"`
	SectionIDs                  []string                          `json:"section_ids"`
	SectionCount                int                               `json:"section_count"`
	States                      CompatibilityOnboardingStateCount `json:"states"`
	RuntimeOwned                bool                              `json:"runtime_owned"`
	GoRuntimeBacked             bool                              `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                              `json:"kde_policy_owner"`
	UserVisible                 bool                              `json:"user_visible"`
	Ready                       bool                              `json:"ready"`
	RecommendedNextStep         string                            `json:"recommended_next_step"`
	RuntimeWriteMethodsEnabled  bool                              `json:"runtime_write_methods_enabled"`
	RequestObjectsCreated       bool                              `json:"request_objects_created"`
	PermissionGrantsCreated     bool                              `json:"permission_grants_created"`
	ArtifactStaged              bool                              `json:"artifact_staged"`
	SettingsPersisted           bool                              `json:"settings_persisted"`
	BackendProcessStarted       bool                              `json:"backend_process_started"`
	LaunchEnabled               bool                              `json:"launch_enabled"`
	NetworkRequired             bool                              `json:"network_required"`
	HostPackageManagerInvoked   bool                              `json:"host_package_manager_invoked"`
	HostRootModified            bool                              `json:"host_root_modified"`
	PrivilegedContainerRequired bool                              `json:"privileged_container_required"`
	AIProviderCalled            bool                              `json:"ai_provider_called"`
	StateRootPathExposed        bool                              `json:"state_root_path_exposed"`
	RawExecutableExposed        bool                              `json:"raw_executable_exposed"`
	RawCommandExposed           bool                              `json:"raw_command_exposed"`
	FileContentRead             bool                              `json:"file_content_read"`
	BackendDetailsExposed       bool                              `json:"backend_details_exposed"`
	BlockedActions              []string                          `json:"blocked_actions"`
	DesktopSafeSummary          string                            `json:"desktop_safe_summary"`
}

type CompatibilityOnboardingIdentity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DesktopFile string `json:"desktop_file"`
	RuntimeMode string `json:"runtime_mode"`
}

type CompatibilityOnboardingSection struct {
	ID                    string   `json:"id"`
	Title                 string   `json:"title"`
	State                 string   `json:"state"`
	Required              bool     `json:"required"`
	EvidenceSource        string   `json:"evidence_source"`
	ReadModel             string   `json:"read_model"`
	RuntimeMethod         string   `json:"runtime_method"`
	Summary               string   `json:"summary"`
	MissingEvidence       []string `json:"missing_evidence"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	RuntimeOwned          bool     `json:"runtime_owned"`
	GoRuntimeBacked       bool     `json:"go_runtime_backed"`
	KDEPolicyOwner        bool     `json:"kde_policy_owner"`
	SideEffectsEnabled    bool     `json:"side_effects_enabled"`
	HostRootModified      bool     `json:"host_root_modified"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type CompatibilityOnboardingStateCount struct {
	Ready             int `json:"ready"`
	NeedsReview       int `json:"needs_review"`
	MissingEvidence   int `json:"missing_evidence"`
	Blocked           int `json:"blocked"`
	NotYetImplemented int `json:"not_yet_implemented"`
}

func (plan Plan) CompatibilityOnboardingChecklistPreview(options CompatibilityOnboardingChecklistOptions) (CompatibilityOnboardingChecklistPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}
	if options.PortalOperation == "" {
		options.PortalOperation = "file-open"
	}
	if options.SnapshotReason == "" {
		options.SnapshotReason = "before-repair"
	}
	if options.Issue == "" {
		options.Issue = "portal-approval-required"
	}
	if options.TestType == "" {
		options.TestType = "smoke"
	}

	ownerReadiness, err := NewRuntimeOwnerReadinessPreview(options.RuntimeRoot)
	if err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	applicationReadiness, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{
		RuntimeRoot:     options.RuntimeRoot,
		StateRoot:       options.StateRoot,
		PortalOperation: options.PortalOperation,
		SnapshotReason:  options.SnapshotReason,
		WriteMethod:     "Launch",
	})
	if err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	portalPolicy, err := NewPortalAccessPolicyPreview(plan.ApplicationID, options.PortalOperation)
	if err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	snapshot, err := NewSnapshotPlanPreview(plan.ApplicationID, options.SnapshotReason)
	if err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	diagnostics, err := plan.AIDiagnosticInputPreview(options.Issue, options.TestType)
	if err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	desktopSafety := NewDesktopSafetyPolicyPreview()

	sections := compatibilityOnboardingSections(ownerReadiness, applicationReadiness, portalPolicy, snapshot, diagnostics, desktopSafety, options.ArtifactReceipt)
	stateCounts := countCompatibilityOnboardingStates(sections)
	preview := CompatibilityOnboardingChecklistPreview{
		Version:       ownerReadiness.Version,
		SchemaVersion: "xnix.runtime.compatibility_onboarding_checklist.v1",
		RequestType:   "compatibility-onboarding-checklist-preview",
		ChecklistType: "first-run-compatibility-onboarding",
		Source:        "runtime-owner-readiness-preview+application-readiness-preview+portal-access-policy-preview+snapshot-plan-preview+ai-diagnostic-input-preview+desktop-safety-policy-preview",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetCompatibilityOnboardingChecklist",
		ReadMethod:    "GetCompatibilityOnboardingChecklistPreview",
		Application: CompatibilityOnboardingIdentity{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			RuntimeMode: plan.runtimeModeID(),
		},
		Sections:                    sections,
		SectionIDs:                  compatibilityOnboardingSectionIDs(sections),
		SectionCount:                len(sections),
		States:                      stateCounts,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		Ready:                       false,
		RecommendedNextStep:         "Review the missing evidence sections before enabling any compatibility action.",
		RuntimeWriteMethodsEnabled:  false,
		RequestObjectsCreated:       false,
		PermissionGrantsCreated:     false,
		ArtifactStaged:              false,
		SettingsPersisted:           false,
		BackendProcessStarted:       false,
		LaunchEnabled:               false,
		NetworkRequired:             false,
		HostPackageManagerInvoked:   false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		AIProviderCalled:            false,
		StateRootPathExposed:        false,
		RawExecutableExposed:        false,
		RawCommandExposed:           false,
		FileContentRead:             false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"create desktop permission requests from onboarding",
			"stage artifacts from onboarding",
			"persist settings from onboarding",
			"start compatibility services from onboarding",
			"fetch network artifacts from onboarding",
			"invoke host package managers from onboarding",
			"mutate host root during onboarding inspection",
		},
		DesktopSafeSummary: "First-run onboarding explains Runtime readiness for KDE without creating requests, staging artifacts, changing settings, starting services, fetching network artifacts, or mutating the host root.",
	}
	if err := validateNoBackendTerms(preview, "compatibility onboarding checklist preview"); err != nil {
		return CompatibilityOnboardingChecklistPreview{}, err
	}
	return preview, nil
}

func compatibilityOnboardingSections(owner RuntimeOwnerReadinessPreview, readiness ApplicationReadinessPreview, portal PortalAccessPolicyPreview, snapshot SnapshotPlanPreview, diagnostics AIDiagnosticInputPreview, desktopSafety DesktopSafetyPolicyPreview, artifactReceipt bool) []CompatibilityOnboardingSection {
	return []CompatibilityOnboardingSection{
		compatibilityOnboardingSection("runtime-owner-readiness", "Runtime owner readiness", compatibilityOwnerReadinessState(owner), true, owner.Source, owner.RequestType, owner.RuntimeMethod, "Runtime read ownership must be explainable before KDE depends on it.", compatibilityOwnerReadinessMissing(owner), owner.RequestType),
		compatibilityOnboardingSection("recipe-trust", "Recipe trust", compatibilityRecipeTrustState(owner), true, "runtime-owner-recipe-trust-preview", owner.RecipeTrust.RequestType, "GetRuntimeOwnerRecipeTrust", "Local recipe evidence must be digest-verified and production trust still needs review.", compatibilityRecipeTrustMissing(owner), owner.RecipeTrust.RequestType),
		compatibilityOnboardingSection("artifact-staging", "Artifact staging", compatibilityArtifactStagingState(readiness, artifactReceipt), true, readiness.Source, "artifact-stage-receipt", "GetCompatibilityArtifactManifest", "Required local artifacts need a verified staging receipt before install can proceed.", compatibilityArtifactStagingMissing(readiness, artifactReceipt), "application-readiness-preview"),
		compatibilityOnboardingSection("backend-lifecycle", "Compatibility service lifecycle", compatibilityBackendLifecycleState(readiness), true, "backend-lifecycle-preview", "backend-lifecycle-preview", "GetBackendLifecycle", "Compatibility service state must be recorded before execution can become safe.", compatibilityBackendLifecycleMissing(readiness), "backend-lifecycle-preview"),
		compatibilityOnboardingSection("portal-review", "Desktop resource review", compatibilityPortalReviewState(portal), true, portal.Source, portal.RequestType, portal.RuntimeMethod, "Sensitive desktop resources require user-mediated review.", compatibilityPortalReviewMissing(portal), portal.RequestType),
		compatibilityOnboardingSection("snapshot-baseline", "Restore-point baseline", compatibilitySnapshotState(snapshot), true, snapshot.Source, snapshot.RequestType, snapshot.RuntimeMethod, "A restore-point plan exists, but no restore-point receipt has been created by this preview.", compatibilitySnapshotMissing(snapshot), snapshot.RequestType),
		compatibilityOnboardingSection("diagnostics-privacy", "Diagnostics privacy", compatibilityDiagnosticsPrivacyState(diagnostics), true, diagnostics.Source, diagnostics.RequestType, diagnostics.RuntimeMethod, "AI diagnostics are limited to safe Runtime metadata and do not call a provider.", compatibilityDiagnosticsMissing(diagnostics), diagnostics.RequestType),
		compatibilityOnboardingSection("kde-entry-points", "KDE entry points", compatibilityKDEEntryPointState(desktopSafety), true, "desktop-safety-policy-preview", desktopSafety.RequestType, desktopSafety.RuntimeMethod, "KDE can show the seven first-release entry points while Runtime owns policy.", compatibilityKDEEntryPointMissing(desktopSafety), "kde-journey-evidence-preview"),
		compatibilityOnboardingSection("production-activation", "Production activation", "not-yet-implemented", true, "runtime-write-gate-preview", "runtime-write-gate-preview", "GetRuntimeWriteGate", "Production install, launch, snapshot, restore, and settings writes remain deliberately unavailable.", []string{"production write enablement evidence", "production owner bus-claim evidence"}, "runtime-write-gate-preview"),
	}
}

func compatibilityOnboardingSection(id string, title string, state string, required bool, evidenceSource string, readModel string, runtimeMethod string, summary string, missingEvidence []string, nextCheck string) CompatibilityOnboardingSection {
	return CompatibilityOnboardingSection{
		ID:                    id,
		Title:                 title,
		State:                 state,
		Required:              required,
		EvidenceSource:        evidenceSource,
		ReadModel:             readModel,
		RuntimeMethod:         runtimeMethod,
		Summary:               summary,
		MissingEvidence:       append([]string(nil), missingEvidence...),
		NextSafeReadOnlyCheck: nextCheck,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		SideEffectsEnabled:    false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
	}
}

func compatibilityOwnerReadinessState(owner RuntimeOwnerReadinessPreview) string {
	if owner.Counts.Blocked > 0 {
		return "blocked"
	}
	if owner.Counts.Pending > 0 {
		return "needs-review"
	}
	return "ready"
}

func compatibilityRecipeTrustState(owner RuntimeOwnerReadinessPreview) string {
	if owner.RecipeTrust.ProductionRecipeTrustReady {
		return "ready"
	}
	if owner.RecipeTrust.Counts.Blocked > 0 {
		return "blocked"
	}
	return "needs-review"
}

func compatibilityArtifactStagingState(readiness ApplicationReadinessPreview, artifactReceipt bool) string {
	if artifactReceipt {
		return "ready"
	}
	if readiness.InstallReadiness.ArtifactStageReceiptReady && readiness.InstallReadiness.RequiredArtifactsStaged {
		return "ready"
	}
	return "missing-evidence"
}

func compatibilityBackendLifecycleState(readiness ApplicationReadinessPreview) string {
	if readiness.BackendLifecycleState == "ready" {
		return "ready"
	}
	if readiness.BackendLifecycleState == "blocked" {
		return "blocked"
	}
	return "missing-evidence"
}

func compatibilityPortalReviewState(portal PortalAccessPolicyPreview) string {
	if portal.PermissionGranted {
		return "ready"
	}
	if portal.Decision == "deny" {
		return "blocked"
	}
	return "needs-review"
}

func compatibilitySnapshotState(snapshot SnapshotPlanPreview) string {
	if snapshot.SnapshotCreated {
		return "ready"
	}
	return "missing-evidence"
}

func compatibilityDiagnosticsPrivacyState(diagnostics AIDiagnosticInputPreview) string {
	if diagnostics.SafeForAIDiagnostics && !diagnostics.AIProviderCalled && !diagnostics.FileContentRead {
		return "ready"
	}
	return "blocked"
}

func compatibilityKDEEntryPointState(desktopSafety DesktopSafetyPolicyPreview) string {
	if desktopSafety.EntryPointCount == 7 && desktopSafety.RuntimeOwned && !desktopSafety.KDEPolicyOwner {
		return "ready"
	}
	return "blocked"
}

func compatibilityOwnerReadinessMissing(owner RuntimeOwnerReadinessPreview) []string {
	if owner.Counts.Blocked == 0 && owner.Counts.Pending == 0 {
		return nil
	}
	return append([]string(nil), owner.BlockedReasons...)
}

func compatibilityRecipeTrustMissing(owner RuntimeOwnerReadinessPreview) []string {
	if owner.RecipeTrust.ProductionRecipeTrustReady {
		return nil
	}
	missing := []string{}
	if !owner.RecipeTrust.SignedRecipeValidation {
		missing = append(missing, "production recipe signature validation")
	}
	if owner.RecipeTrust.DevelopmentRegistry {
		missing = append(missing, "production recipe registry")
	}
	if owner.RecipeTrust.UnsignedRecipesPresent {
		missing = append(missing, "signed recipe records")
	}
	return missing
}

func compatibilityArtifactStagingMissing(readiness ApplicationReadinessPreview, artifactReceipt bool) []string {
	if artifactReceipt || (readiness.InstallReadiness.ArtifactStageReceiptReady && readiness.InstallReadiness.RequiredArtifactsStaged) {
		return nil
	}
	return []string{"verified local artifact staging receipt", "required artifact digest evidence"}
}

func compatibilityBackendLifecycleMissing(readiness ApplicationReadinessPreview) []string {
	if readiness.BackendLifecycleState == "ready" {
		return nil
	}
	return []string{"recorded compatibility service lifecycle state", "service binding readiness evidence", "state-root lifecycle receipt"}
}

func compatibilityPortalReviewMissing(portal PortalAccessPolicyPreview) []string {
	if portal.PermissionGranted {
		return nil
	}
	if portal.Decision == "deny" {
		return []string{"user policy change for denied desktop resource"}
	}
	return []string{"user-mediated desktop resource approval receipt"}
}

func compatibilitySnapshotMissing(snapshot SnapshotPlanPreview) []string {
	if snapshot.SnapshotCreated {
		return nil
	}
	return []string{"restore-point creation receipt"}
}

func compatibilityDiagnosticsMissing(diagnostics AIDiagnosticInputPreview) []string {
	if diagnostics.SafeForAIDiagnostics && !diagnostics.AIProviderCalled && !diagnostics.FileContentRead {
		return nil
	}
	return []string{"safe diagnostic metadata boundary"}
}

func compatibilityKDEEntryPointMissing(desktopSafety DesktopSafetyPolicyPreview) []string {
	if desktopSafety.EntryPointCount == 7 && desktopSafety.RuntimeOwned && !desktopSafety.KDEPolicyOwner {
		return nil
	}
	return []string{"seven KDE entry-point safety evidence"}
}

func compatibilityOnboardingSectionIDs(sections []CompatibilityOnboardingSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func countCompatibilityOnboardingStates(sections []CompatibilityOnboardingSection) CompatibilityOnboardingStateCount {
	var counts CompatibilityOnboardingStateCount
	for _, section := range sections {
		switch section.State {
		case "ready":
			counts.Ready++
		case "needs-review":
			counts.NeedsReview++
		case "missing-evidence":
			counts.MissingEvidence++
		case "blocked":
			counts.Blocked++
		case "not-yet-implemented":
			counts.NotYetImplemented++
		}
	}
	return counts
}
