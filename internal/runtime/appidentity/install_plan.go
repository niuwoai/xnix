package appidentity

import "xnix.local/xnix/internal/runtime/artifact"

type CompatibilityInstallPlanPreview struct {
	SchemaVersion               string                                    `json:"schema_version"`
	RequestType                 string                                    `json:"request_type"`
	PlanType                    string                                    `json:"plan_type"`
	Source                      string                                    `json:"source"`
	Desktop                     string                                    `json:"desktop"`
	RuntimeMethod               string                                    `json:"runtime_method"`
	ReadMethod                  string                                    `json:"read_method"`
	Application                 PackageSourceApplication                  `json:"application"`
	Environment                 string                                    `json:"environment"`
	SelectedStrategy            string                                    `json:"selected_strategy"`
	InstallState                string                                    `json:"install_state"`
	Readiness                   CompatibilityInstallReadiness             `json:"readiness"`
	Phases                      []CompatibilityInstallPhase               `json:"phases"`
	PhaseIDs                    []string                                  `json:"phase_ids"`
	RecipeTrust                 DesktopActivationRecipeTrust              `json:"recipe_trust"`
	InstallGate                 DesktopActivationInstallGate              `json:"install_gate"`
	ArtifactStageReceipt        *CompatibilityInstallArtifactStageReceipt `json:"artifact_stage_receipt,omitempty"`
	RuntimeOwned                bool                                      `json:"runtime_owned"`
	GoRuntimeBacked             bool                                      `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                      `json:"kde_policy_owner"`
	UserVisible                 bool                                      `json:"user_visible"`
	InstallReady                bool                                      `json:"install_ready"`
	DesktopActivationReady      bool                                      `json:"desktop_activation_ready"`
	DownloadEnabled             bool                                      `json:"download_enabled"`
	InstallEnabled              bool                                      `json:"install_enabled"`
	NetworkRequestCreated       bool                                      `json:"network_request_created"`
	ArtifactsDownloaded         bool                                      `json:"artifacts_downloaded"`
	HostRootModified            bool                                      `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                      `json:"privileged_container_required"`
	DesktopShellCommandExposed  bool                                      `json:"desktop_shell_command_exposed"`
	BackendLaunchEnabled        bool                                      `json:"backend_launch_enabled"`
	ExecutionStarted            bool                                      `json:"execution_started"`
	BackendDetailsExposed       bool                                      `json:"backend_details_exposed"`
	BlockedActions              []string                                  `json:"blocked_actions"`
	DesktopSafeSummary          string                                    `json:"desktop_safe_summary"`
}

type CompatibilityInstallReadiness struct {
	ArtifactManifestReady        bool     `json:"artifact_manifest_ready"`
	ArtifactSignatureVerified    bool     `json:"artifact_signature_verified"`
	ArtifactStageReceiptReady    bool     `json:"artifact_stage_receipt_ready"`
	ArtifactStageDigestVerified  bool     `json:"artifact_stage_digest_verified"`
	RequiredArtifactsStaged      bool     `json:"required_artifacts_staged"`
	ArtifactStageBlockingReasons []string `json:"artifact_stage_blocking_reasons"`
	AcquisitionReady             bool     `json:"acquisition_ready"`
	PackageSourceReady           bool     `json:"package_source_ready"`
	StateRootAllocated           bool     `json:"state_root_allocated"`
	RecipeTrustDiagnosticsReady  bool     `json:"recipe_trust_diagnostics_ready"`
	RecipeInstallAllowed         bool     `json:"recipe_install_allowed"`
	RecipeInstallDecision        string   `json:"recipe_install_decision"`
	RecipeTrustBlockingReasons   []string `json:"recipe_trust_blocking_reasons"`
}

type CompatibilityInstallArtifactStageReceipt struct {
	SchemaVersion           string   `json:"schema_version"`
	RecordType              string   `json:"record_type"`
	ApplicationID           string   `json:"application_id"`
	RelativePath            string   `json:"relative_path"`
	SHA256                  string   `json:"sha256"`
	RequiredArtifactsStaged bool     `json:"required_artifacts_staged"`
	ArtifactCount           int      `json:"artifact_count"`
	StagedKeyCount          int      `json:"staged_key_count"`
	BlockingReasons         []string `json:"blocking_reasons"`
}

type CompatibilityInstallPhase struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

func (plan Plan) CompatibilityInstallPlanPreview(environment string) (CompatibilityInstallPlanPreview, error) {
	return plan.CompatibilityInstallPlanPreviewWithArtifactReceipt(environment, nil)
}

func (plan Plan) CompatibilityInstallPlanPreviewWithArtifactReceipt(environment string, receipt *artifact.StageReceipt) (CompatibilityInstallPlanPreview, error) {
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
	receiptPreview := compatibilityInstallArtifactStageReceipt(receipt)
	readiness := compatibilityInstallReadiness(artifactManifest, acquisitionPreflight, packageSource, stateRoot, installGate, receiptPreview)
	phases := compatibilityInstallPhases(readiness, installGate)

	preview := CompatibilityInstallPlanPreview{
		SchemaVersion:               "xnix.runtime.compatibility_install_plan.v1",
		RequestType:                 "compatibility-install-preview",
		PlanType:                    "compatibility-install-plan",
		Source:                      compatibilityInstallSource(receiptPreview),
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
		ArtifactStageReceipt:        receiptPreview,
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

func compatibilityInstallReadiness(artifactManifest ArtifactManifestPreview, acquisitionPreflight AcquisitionPreflightPreview, packageSource PackageSourcePreview, stateRoot ApplicationStateRootPreview, installGate DesktopActivationInstallGate, receipt *CompatibilityInstallArtifactStageReceipt) CompatibilityInstallReadiness {
	blockingReasons := append([]string{}, installGate.BlockingReasons...)
	artifactStageReady := false
	artifactStageDigestVerified := false
	requiredArtifactsStaged := false
	var artifactStageBlockingReasons []string
	if receipt != nil {
		artifactStageBlockingReasons = append([]string{}, receipt.BlockingReasons...)
		artifactStageReady = len(receipt.BlockingReasons) == 0
		artifactStageDigestVerified = artifactStageReady && receipt.SHA256 != ""
		requiredArtifactsStaged = artifactStageReady && receipt.RequiredArtifactsStaged
	}
	return CompatibilityInstallReadiness{
		ArtifactManifestReady:        artifactManifest.ManifestReady,
		ArtifactSignatureVerified:    artifactManifest.SignatureVerified,
		ArtifactStageReceiptReady:    artifactStageReady,
		ArtifactStageDigestVerified:  artifactStageDigestVerified,
		RequiredArtifactsStaged:      requiredArtifactsStaged,
		ArtifactStageBlockingReasons: artifactStageBlockingReasons,
		AcquisitionReady:             acquisitionPreflight.AcquisitionReady,
		PackageSourceReady:           packageSource.PackageSourceReady,
		StateRootAllocated:           stateRoot.AllocationState == "allocated",
		RecipeTrustDiagnosticsReady:  installGate.Decision != "",
		RecipeInstallAllowed:         installGate.Decision == "allow",
		RecipeInstallDecision:        installGate.Decision,
		RecipeTrustBlockingReasons:   blockingReasons,
	}
}

func compatibilityInstallPhases(readiness CompatibilityInstallReadiness, installGate DesktopActivationInstallGate) []CompatibilityInstallPhase {
	return []CompatibilityInstallPhase{
		compatibilityInstallPhase("resolve-artifact-manifest", boolStatus(readiness.ArtifactManifestReady), "Wait for a signed compatibility artifact manifest."),
		compatibilityInstallPhase("verify-artifact-digests", boolStatus(readiness.ArtifactSignatureVerified), "Wait for digest verification before artifact activation."),
		compatibilityInstallPhase("consume-artifact-stage-receipt", boolStatus(readiness.ArtifactStageReceiptReady && readiness.RequiredArtifactsStaged), "Consume a Runtime-owned artifact staging receipt before install can depend on artifacts."),
		compatibilityInstallPhase("prepare-package-source", boolStatus(readiness.PackageSourceReady), "Wait for Runtime-owned package source readiness."),
		compatibilityInstallPhase("allocate-application-state", boolStatus(readiness.StateRootAllocated), "Wait for Runtime-owned application state allocation."),
		compatibilityInstallPhase("review-recipe-install-gate", boolStatus(installGate.Decision == "allow"), "Runtime recipe install gate must allow this environment before staging."),
		compatibilityInstallPhase("stage-desktop-integration", "pending", "Stage launcher, file association, Dolphin, tray, notification, Compatibility Center, and settings artifacts only after install gates pass."),
		compatibilityInstallPhase("enable-launch-binding", "blocked", "Launch binding stays disabled until install preflight completes."),
	}
}

func compatibilityInstallArtifactStageReceipt(receipt *artifact.StageReceipt) *CompatibilityInstallArtifactStageReceipt {
	if receipt == nil {
		return nil
	}
	blockingReasons := artifact.ValidateStageReceipt(*receipt)
	return &CompatibilityInstallArtifactStageReceipt{
		SchemaVersion:           receipt.SchemaVersion,
		RecordType:              receipt.RecordType,
		ApplicationID:           receipt.ApplicationID,
		RelativePath:            receipt.RelativePath,
		SHA256:                  receipt.SHA256,
		RequiredArtifactsStaged: receipt.Plan.RequiredStaged(),
		ArtifactCount:           len(receipt.Plan.Artifacts),
		StagedKeyCount:          len(receipt.Plan.StagedKeys),
		BlockingReasons:         blockingReasons,
	}
}

func compatibilityInstallSource(receipt *CompatibilityInstallArtifactStageReceipt) string {
	source := "artifact-manifest-preview+acquisition-preflight-preview+package-source-preview+state-root-preview+recipe-install-gate"
	if receipt != nil {
		return source + "+artifact-stage-receipt"
	}
	return source
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
