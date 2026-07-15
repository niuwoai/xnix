package appidentity

type CompatibilityInstallPlanPreview struct {
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	PlanType                    string                        `json:"plan_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	RuntimeMethod               string                        `json:"runtime_method"`
	ReadMethod                  string                        `json:"read_method"`
	Application                 PackageSourceApplication      `json:"application"`
	Environment                 string                        `json:"environment"`
	SelectedStrategy            string                        `json:"selected_strategy"`
	InstallState                string                        `json:"install_state"`
	Readiness                   CompatibilityInstallReadiness `json:"readiness"`
	Phases                      []CompatibilityInstallPhase   `json:"phases"`
	PhaseIDs                    []string                      `json:"phase_ids"`
	RecipeTrust                 DesktopActivationRecipeTrust  `json:"recipe_trust"`
	InstallGate                 DesktopActivationInstallGate  `json:"install_gate"`
	RuntimeOwned                bool                          `json:"runtime_owned"`
	GoRuntimeBacked             bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                          `json:"kde_policy_owner"`
	UserVisible                 bool                          `json:"user_visible"`
	InstallReady                bool                          `json:"install_ready"`
	DesktopActivationReady      bool                          `json:"desktop_activation_ready"`
	DownloadEnabled             bool                          `json:"download_enabled"`
	InstallEnabled              bool                          `json:"install_enabled"`
	NetworkRequestCreated       bool                          `json:"network_request_created"`
	ArtifactsDownloaded         bool                          `json:"artifacts_downloaded"`
	HostRootModified            bool                          `json:"host_root_modified"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	DesktopShellCommandExposed  bool                          `json:"desktop_shell_command_exposed"`
	BackendLaunchEnabled        bool                          `json:"backend_launch_enabled"`
	ExecutionStarted            bool                          `json:"execution_started"`
	BackendDetailsExposed       bool                          `json:"backend_details_exposed"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type CompatibilityInstallReadiness struct {
	ArtifactManifestReady     bool   `json:"artifact_manifest_ready"`
	ArtifactSignatureVerified bool   `json:"artifact_signature_verified"`
	AcquisitionReady          bool   `json:"acquisition_ready"`
	PackageSourceReady        bool   `json:"package_source_ready"`
	StateRootAllocated        bool   `json:"state_root_allocated"`
	RecipeInstallAllowed      bool   `json:"recipe_install_allowed"`
	RecipeInstallDecision     string `json:"recipe_install_decision"`
}

type CompatibilityInstallPhase struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) CompatibilityInstallPlanPreview(environment string) (CompatibilityInstallPlanPreview, error) {
	mode, err := normalizeActivationPreflightMode(environment)
	if err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	artifactManifest, err := plan.ArtifactManifestPreview()
	if err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	acquisitionPreflight, err := plan.AcquisitionPreflightPreview()
	if err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	packageSource, err := plan.PackageSourcePreview()
	if err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	stateRoot, err := plan.ApplicationStateRootPreview()
	if err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	trust := desktopActivationRecipeTrust(plan)
	installGate := desktopActivationInstallGate(trust, mode)
	readiness := compatibilityInstallReadiness(artifactManifest, acquisitionPreflight, packageSource, stateRoot, installGate)
	phases := compatibilityInstallPhases(readiness, installGate)

	preview := CompatibilityInstallPlanPreview{
		SchemaVersion:               "xnix.runtime.compatibility_install_plan.v1",
		RequestType:                 "compatibility-install-preview",
		PlanType:                    "compatibility-install-plan",
		Source:                      "artifact-manifest-preview+acquisition-preflight-preview+package-source-preview+state-root-preview+recipe-install-gate",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetCompatibilityInstallPlan",
		ReadMethod:                  "GetCompatibilityInstallPlanPreview",
		Application:                 packageSource.Application,
		Environment:                 mode,
		SelectedStrategy:            artifactManifest.SelectedStrategy,
		InstallState:                "planned",
		Readiness:                   readiness,
		Phases:                      phases,
		PhaseIDs:                    compatibilityInstallPhaseIDs(phases),
		RecipeTrust:                 trust,
		InstallGate:                 installGate,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		InstallReady:                false,
		DesktopActivationReady:      false,
		DownloadEnabled:             false,
		InstallEnabled:              false,
		NetworkRequestCreated:       false,
		ArtifactsDownloaded:         false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		DesktopShellCommandExposed:  false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"download artifacts before signed manifest verification",
			"install packages before Runtime install plan readiness",
			"stage desktop integration before recipe install gate approval",
			"launch backend before managed binding readiness",
			"expose backend commands or storage paths to KDE",
			"mutate the host root during install planning",
		},
		DesktopSafeSummary: compatibilityInstallSummary(mode, installGate.Decision),
	}
	if err := validateNoBackendTerms(preview, "compatibility install plan preview"); err != nil {
		return CompatibilityInstallPlanPreview{}, err
	}
	return preview, nil
}

func compatibilityInstallReadiness(artifactManifest ArtifactManifestPreview, acquisitionPreflight AcquisitionPreflightPreview, packageSource PackageSourcePreview, stateRoot ApplicationStateRootPreview, installGate DesktopActivationInstallGate) CompatibilityInstallReadiness {
	return CompatibilityInstallReadiness{
		ArtifactManifestReady:     artifactManifest.ManifestReady,
		ArtifactSignatureVerified: artifactManifest.SignatureVerified,
		AcquisitionReady:          acquisitionPreflight.AcquisitionReady,
		PackageSourceReady:        packageSource.PackageSourceReady,
		StateRootAllocated:        stateRoot.AllocationState == "allocated",
		RecipeInstallAllowed:      installGate.Decision == "allow",
		RecipeInstallDecision:     installGate.Decision,
	}
}

func compatibilityInstallPhases(readiness CompatibilityInstallReadiness, installGate DesktopActivationInstallGate) []CompatibilityInstallPhase {
	return []CompatibilityInstallPhase{
		compatibilityInstallPhase("resolve-artifact-manifest", boolStatus(readiness.ArtifactManifestReady), "Wait for a signed compatibility artifact manifest."),
		compatibilityInstallPhase("verify-artifact-digests", boolStatus(readiness.ArtifactSignatureVerified), "Wait for digest verification before artifact activation."),
		compatibilityInstallPhase("prepare-package-source", boolStatus(readiness.PackageSourceReady), "Wait for Runtime-owned package source readiness."),
		compatibilityInstallPhase("allocate-application-state", boolStatus(readiness.StateRootAllocated), "Wait for Runtime-owned application state allocation."),
		compatibilityInstallPhase("review-recipe-install-gate", boolStatus(installGate.Decision == "allow"), "Runtime recipe install gate must allow this environment before staging."),
		compatibilityInstallPhase("stage-desktop-integration", "pending", "Stage launcher, file association, Dolphin, tray, notification, Compatibility Center, and settings artifacts only after install gates pass."),
		compatibilityInstallPhase("enable-launch-binding", "blocked", "Launch binding stays disabled until install preflight completes."),
	}
}

func compatibilityInstallPhase(id string, status string, summary string) CompatibilityInstallPhase {
	return CompatibilityInstallPhase{ID: id, Status: status, Summary: summary}
}

func compatibilityInstallPhaseIDs(phases []CompatibilityInstallPhase) []string {
	ids := make([]string, 0, len(phases))
	for _, phase := range phases {
		ids = append(ids, phase.ID)
	}
	return ids
}

func compatibilityInstallSummary(environment string, decision string) string {
	if decision == "allow" && environment == "development" {
		return "Compatibility install planning can continue toward development staging, but download, install, desktop activation, launch, and host writes remain disabled."
	}
	if decision == "allow" {
		return "Compatibility install planning can continue, but download, install, desktop activation, launch, and host writes remain disabled."
	}
	return "Compatibility install is planned and waiting for Runtime-owned readiness gates."
}
