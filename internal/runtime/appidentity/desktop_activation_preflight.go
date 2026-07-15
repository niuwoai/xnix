package appidentity

import "errors"

type DesktopActivationPreflightPreview struct {
	SchemaVersion                string                            `json:"schema_version"`
	RequestType                  string                            `json:"request_type"`
	PreflightType                string                            `json:"preflight_type"`
	Source                       string                            `json:"source"`
	Desktop                      string                            `json:"desktop"`
	RuntimeMethod                string                            `json:"runtime_method"`
	ApplicationID                string                            `json:"application_id"`
	DisplayName                  string                            `json:"display_name"`
	Icon                         string                            `json:"icon"`
	DesktopFile                  string                            `json:"desktop_file"`
	InstallMode                  string                            `json:"install_mode"`
	PreflightDecision            string                            `json:"preflight_decision"`
	ActivationState              string                            `json:"activation_state"`
	Bundle                       DesktopActivationPreflightBundle  `json:"bundle"`
	RecipeTrust                  DesktopActivationRecipeTrust      `json:"recipe_trust"`
	InstallGate                  DesktopActivationInstallGate      `json:"install_gate"`
	BackendBinding               DesktopActivationBindingSummary   `json:"backend_binding"`
	Checks                       []DesktopActivationPreflightCheck `json:"checks"`
	CheckIDs                     []string                          `json:"check_ids"`
	CheckCount                   int                               `json:"check_count"`
	PassedCheckCount             int                               `json:"passed_check_count"`
	PendingCheckCount            int                               `json:"pending_check_count"`
	BlockedCheckCount            int                               `json:"blocked_check_count"`
	RequiredReviews              []string                          `json:"required_reviews"`
	RuntimeOwned                 bool                              `json:"runtime_owned"`
	GoRuntimeBacked              bool                              `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                              `json:"kde_policy_owner"`
	UserVisible                  bool                              `json:"user_visible"`
	NormalApplicationSurface     bool                              `json:"normal_application_surface"`
	DesktopActivationReady       bool                              `json:"desktop_activation_ready"`
	DevelopmentStagingEligible   bool                              `json:"development_staging_eligible"`
	ProductionActivationEligible bool                              `json:"production_activation_eligible"`
	InstallerMayProceed          bool                              `json:"installer_may_proceed"`
	StagingRootRequired          bool                              `json:"staging_root_required"`
	HostRootAllowed              bool                              `json:"host_root_allowed"`
	DesktopFilesWritten          bool                              `json:"desktop_files_written"`
	MIMEAppsWritten              bool                              `json:"mimeapps_written"`
	ManifestWritten              bool                              `json:"manifest_written"`
	ReceiptWritten               bool                              `json:"receipt_written"`
	SettingsPersisted            bool                              `json:"settings_persisted"`
	NotificationsSent            bool                              `json:"notifications_sent"`
	TaskManagerEntryActive       bool                              `json:"task_manager_entry_active"`
	KWinRuleApplied              bool                              `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled        bool                              `json:"live_tray_bridge_enabled"`
	LaunchEnabled                bool                              `json:"launch_enabled"`
	BackendLaunchEnabled         bool                              `json:"backend_launch_enabled"`
	ExecutionStarted             bool                              `json:"execution_started"`
	HostRootModified             bool                              `json:"host_root_modified"`
	NetworkRequired              bool                              `json:"network_required"`
	PrivilegedContainerRequired  bool                              `json:"privileged_container_required"`
	BackendDetailsExposed        bool                              `json:"backend_details_exposed"`
	RawWindowsExecutableExposed  bool                              `json:"raw_windows_executable_exposed"`
	CompatibilityStorageExposed  bool                              `json:"compatibility_storage_exposed"`
	BlockedActions               []string                          `json:"blocked_actions"`
	DesktopSafeSummary           string                            `json:"desktop_safe_summary"`
}

type DesktopActivationPreflightBundle struct {
	RequestType              string   `json:"request_type"`
	MaterialCount            int      `json:"material_count"`
	MaterialIDs              []string `json:"material_ids"`
	StandardDesktopEntry     bool     `json:"standard_desktop_entry"`
	FileAssociationReady     bool     `json:"file_association_ready"`
	TaskManagerIdentityReady bool     `json:"task_manager_identity_ready"`
	KWinIdentityReady        bool     `json:"kwin_identity_ready"`
	TrayStatusReady          bool     `json:"tray_status_ready"`
	NotificationReady        bool     `json:"notification_ready"`
	SettingsReady            bool     `json:"settings_ready"`
	CompatibilityCenterReady bool     `json:"compatibility_center_ready"`
	DesktopFilesWritten      bool     `json:"desktop_files_written"`
	HostRootModified         bool     `json:"host_root_modified"`
	BackendDetailsExposed    bool     `json:"backend_details_exposed"`
}

type DesktopActivationRecipeTrust struct {
	Source                 string `json:"source"`
	RegistryName           string `json:"registry_name,omitempty"`
	DigestVerified         bool   `json:"digest_verified"`
	SignatureStatus        string `json:"signature_status"`
	TrustDecision          string `json:"trust_decision"`
	ProductionTrusted      bool   `json:"production_trusted"`
	DevelopmentOnly        bool   `json:"development_only"`
	ExternalRecipeAccepted bool   `json:"external_recipe_accepted"`
}

type DesktopActivationInstallGate struct {
	GateType        string   `json:"gate_type"`
	Mode            string   `json:"mode"`
	Decision        string   `json:"decision"`
	BlockingReasons []string `json:"blocking_reasons"`
	Requirements    []string `json:"requirements"`
}

type DesktopActivationBindingSummary struct {
	RequestType           string `json:"request_type"`
	BindingState          string `json:"binding_state"`
	RecommendedProfileID  string `json:"recommended_profile_id"`
	RequiredReviewCount   int    `json:"required_review_count"`
	BindingCommitted      bool   `json:"binding_committed"`
	BindingPersisted      bool   `json:"binding_persisted"`
	LaunchEnabled         bool   `json:"launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type DesktopActivationPreflightCheck struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Summary  string `json:"summary"`
}

func (plan Plan) DesktopActivationPreflightPreview(mode string) (DesktopActivationPreflightPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationPreflightPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationPreflightPreview{}, errors.New("desktop activation preflight requires single-line identity fields")
		}
	}
	normalizedMode, err := normalizeActivationPreflightMode(mode)
	if err != nil {
		return DesktopActivationPreflightPreview{}, err
	}

	bundle, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		return DesktopActivationPreflightPreview{}, err
	}
	binding, err := plan.BackendBindingPreview()
	if err != nil {
		return DesktopActivationPreflightPreview{}, err
	}
	trust := desktopActivationRecipeTrust(plan)
	installGate := desktopActivationInstallGate(trust, normalizedMode)
	checks := desktopActivationPreflightChecks(bundle, trust, installGate, binding, normalizedMode)
	passedCheckCount, pendingCheckCount, blockedCheckCount := countDesktopActivationPreflightChecks(checks)
	developmentStagingEligible := normalizedMode == "development" && installGate.Decision == "allow"
	productionActivationEligible := normalizedMode == "production" && installGate.Decision == "allow"
	installerMayProceed := developmentStagingEligible || productionActivationEligible

	preview := DesktopActivationPreflightPreview{
		SchemaVersion:                "xnix.runtime.desktop_activation_preflight.v1",
		RequestType:                  "desktop-activation-preflight-preview",
		PreflightType:                "normal-linux-application-activation-preflight",
		Source:                       "desktop-activation-bundle-preview+recipe-trust-policy+recipe-install-gate+backend-binding-preview",
		Desktop:                      "KDE Plasma",
		RuntimeMethod:                "GetDesktopActivationPreflight",
		ApplicationID:                plan.ApplicationID,
		DisplayName:                  plan.DisplayName,
		Icon:                         plan.Icon,
		DesktopFile:                  plan.DesktopFile,
		InstallMode:                  normalizedMode,
		PreflightDecision:            desktopActivationPreflightDecision(normalizedMode, installGate.Decision),
		ActivationState:              "preflight-only",
		Bundle:                       desktopActivationPreflightBundle(bundle),
		RecipeTrust:                  trust,
		InstallGate:                  installGate,
		BackendBinding:               desktopActivationBindingSummary(binding),
		Checks:                       checks,
		CheckIDs:                     desktopActivationPreflightCheckIDs(checks),
		CheckCount:                   len(checks),
		PassedCheckCount:             passedCheckCount,
		PendingCheckCount:            pendingCheckCount,
		BlockedCheckCount:            blockedCheckCount,
		RequiredReviews:              []string{"recipe-signature-review", "recipe-install-gate", "backend-binding-review", "staging-root-review", "host-root-write-gate"},
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		UserVisible:                  true,
		NormalApplicationSurface:     true,
		DesktopActivationReady:       bundle.MaterialCount == 9,
		DevelopmentStagingEligible:   developmentStagingEligible,
		ProductionActivationEligible: productionActivationEligible,
		InstallerMayProceed:          installerMayProceed,
		StagingRootRequired:          true,
		HostRootAllowed:              false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ManifestWritten:              false,
		ReceiptWritten:               false,
		SettingsPersisted:            false,
		NotificationsSent:            false,
		TaskManagerEntryActive:       false,
		KWinRuleApplied:              false,
		LiveTrayBridgeEnabled:        false,
		LaunchEnabled:                false,
		BackendLaunchEnabled:         false,
		ExecutionStarted:             false,
		HostRootModified:             false,
		NetworkRequired:              false,
		PrivilegedContainerRequired:  false,
		BackendDetailsExposed:        false,
		RawWindowsExecutableExposed:  false,
		CompatibilityStorageExposed:  false,
		BlockedActions: []string{
			"write desktop files from activation preflight preview",
			"write MIME defaults from activation preflight preview",
			"write activation manifest from activation preflight preview",
			"write activation receipt from activation preflight preview",
			"promote development registry to production activation",
			"activate task manager entry before a live session",
			"apply KWin rule before Runtime launch approval",
			"enable live tray bridge before Runtime launch approval",
			"start compatibility backend from activation preflight preview",
			"mutate host root during activation preflight",
			"expose raw backend command to desktop shell",
		},
		DesktopSafeSummary: desktopActivationPreflightSummary(normalizedMode, installGate.Decision),
	}
	if err := validateNoBackendTerms(preview, "desktop activation preflight preview"); err != nil {
		return DesktopActivationPreflightPreview{}, err
	}
	return preview, nil
}

func normalizeActivationPreflightMode(mode string) (string, error) {
	if mode == "" {
		return "production", nil
	}
	switch mode {
	case "production", "development":
		return mode, nil
	default:
		return "", errors.New("desktop activation preflight mode must be production or development")
	}
}

func desktopActivationRecipeTrust(plan Plan) DesktopActivationRecipeTrust {
	signatureStatus := plan.RecipeSignatureStatus
	if signatureStatus == "" {
		signatureStatus = "unverified"
	}
	productionTrusted := plan.RecipeDigestVerified && signatureStatus == "signed"
	developmentOnly := plan.RecipeDigestVerified && signatureStatus == "development-only"
	trustDecision := "untrusted"
	if productionTrusted {
		trustDecision = "production-trusted"
	} else if developmentOnly {
		trustDecision = "development-only"
	}
	return DesktopActivationRecipeTrust{
		Source:                 plan.RecipeSource,
		RegistryName:           plan.RegistryName,
		DigestVerified:         plan.RecipeDigestVerified,
		SignatureStatus:        signatureStatus,
		TrustDecision:          trustDecision,
		ProductionTrusted:      productionTrusted,
		DevelopmentOnly:        developmentOnly,
		ExternalRecipeAccepted: productionTrusted,
	}
}

func desktopActivationInstallGate(trust DesktopActivationRecipeTrust, mode string) DesktopActivationInstallGate {
	reasons := desktopActivationInstallBlockingReasons(trust, mode)
	decision := "allow"
	if len(reasons) > 0 {
		decision = "block"
	}
	return DesktopActivationInstallGate{
		GateType:        "recipe-install",
		Mode:            mode,
		Decision:        decision,
		BlockingReasons: reasons,
		Requirements:    desktopActivationInstallRequirements(decision, mode),
	}
}

func desktopActivationInstallBlockingReasons(trust DesktopActivationRecipeTrust, mode string) []string {
	reasons := []string{}
	if !trust.DigestVerified {
		reasons = append(reasons, "recipe digest is not verified")
	}
	if mode == "production" && !trust.ProductionTrusted {
		reasons = append(reasons, "production signed recipe validation is not enabled")
		if trust.DevelopmentOnly {
			reasons = append(reasons, "registry contains development-only recipes")
		}
	}
	if mode == "development" && !trust.DigestVerified {
		reasons = append(reasons, "development staging requires digest verification")
	}
	return uniqueStrings(reasons)
}

func desktopActivationInstallRequirements(decision string, mode string) []string {
	if decision == "allow" {
		return []string{}
	}
	if mode == "production" {
		return []string{
			"Use a registry with a production signed source and verified recipe signatures.",
			"Keep SHA-256 digest verification enabled before activation.",
			"Do not promote development-only recipes into production installation.",
		}
	}
	return []string{
		"Register the application recipe in the selected registry.",
		"Keep SHA-256 digest verification enabled for development staging.",
	}
}

func desktopActivationPreflightChecks(bundle DesktopActivationBundlePreview, trust DesktopActivationRecipeTrust, installGate DesktopActivationInstallGate, binding BackendBindingPreview, mode string) []DesktopActivationPreflightCheck {
	signatureStatus := "pending"
	if trust.ProductionTrusted {
		signatureStatus = "pass"
	} else if !trust.DigestVerified || trust.SignatureStatus == "unsigned" || trust.SignatureStatus == "unverified" {
		signatureStatus = "blocked"
	}
	installStatus := "blocked"
	if installGate.Decision == "allow" {
		installStatus = "pass"
	}
	return []DesktopActivationPreflightCheck{
		desktopActivationPreflightCheck("recipe-digest", boolStatus(trust.DigestVerified), true, "Recipe digest verification must pass before activation."),
		desktopActivationPreflightCheck("recipe-signature", signatureStatus, true, "Production activation requires signed recipe validation; development staging may continue with digest-verified development recipes."),
		desktopActivationPreflightCheck("recipe-install-gate", installStatus, true, "Runtime recipe install gate controls whether activation may continue."),
		desktopActivationPreflightCheck("activation-materials", boolStatus(bundle.MaterialCount == 9), true, "Runtime has generated the normal application activation materials for KDE."),
		desktopActivationPreflightCheck("backend-binding", binding.BindingStateStatus(), true, "Runtime profile binding remains separate from desktop activation."),
		desktopActivationPreflightCheck("staging-root", "pass", true, "Activation writes must target a caller-supplied staging root."),
		desktopActivationPreflightCheck("host-root-write-gate", "blocked", true, "Activation preflight must not authorize host-root writes."),
	}
}

func (preview BackendBindingPreview) BindingStateStatus() string {
	if preview.BindingCommitted && preview.BindingPersisted && preview.LaunchEnabled {
		return "pass"
	}
	if preview.BindingState == "planned-blocked" {
		return "pending"
	}
	return "blocked"
}

func desktopActivationPreflightCheck(id string, status string, required bool, summary string) DesktopActivationPreflightCheck {
	return DesktopActivationPreflightCheck{
		ID:       id,
		Status:   status,
		Required: required,
		Summary:  summary,
	}
}

func boolStatus(ok bool) string {
	if ok {
		return "pass"
	}
	return "blocked"
}

func desktopActivationPreflightBundle(bundle DesktopActivationBundlePreview) DesktopActivationPreflightBundle {
	return DesktopActivationPreflightBundle{
		RequestType:              bundle.RequestType,
		MaterialCount:            bundle.MaterialCount,
		MaterialIDs:              bundle.MaterialIDs,
		StandardDesktopEntry:     bundle.StandardDesktopEntry,
		FileAssociationReady:     bundle.FileAssociationReady,
		TaskManagerIdentityReady: bundle.TaskManagerIdentityReady,
		KWinIdentityReady:        bundle.KWinIdentityReady,
		TrayStatusReady:          bundle.TrayStatusReady,
		NotificationReady:        bundle.NotificationReady,
		SettingsReady:            bundle.SettingsReady,
		CompatibilityCenterReady: bundle.CompatibilityCenterReady,
		DesktopFilesWritten:      bundle.DesktopFilesWritten,
		HostRootModified:         bundle.HostRootModified,
		BackendDetailsExposed:    bundle.BackendDetailsExposed,
	}
}

func desktopActivationBindingSummary(binding BackendBindingPreview) DesktopActivationBindingSummary {
	return DesktopActivationBindingSummary{
		RequestType:           binding.RequestType,
		BindingState:          binding.BindingState,
		RecommendedProfileID:  binding.RecommendedProfileID,
		RequiredReviewCount:   len(binding.RequiredReviews),
		BindingCommitted:      binding.BindingCommitted,
		BindingPersisted:      binding.BindingPersisted,
		LaunchEnabled:         binding.LaunchEnabled,
		BackendDetailsExposed: binding.BackendDetailsExposed,
	}
}

func countDesktopActivationPreflightChecks(checks []DesktopActivationPreflightCheck) (int, int, int) {
	passed := 0
	pending := 0
	blocked := 0
	for _, check := range checks {
		switch check.Status {
		case "pass":
			passed++
		case "pending":
			pending++
		case "blocked":
			blocked++
		}
	}
	return passed, pending, blocked
}

func desktopActivationPreflightCheckIDs(checks []DesktopActivationPreflightCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func desktopActivationPreflightDecision(mode string, installDecision string) string {
	if mode == "development" && installDecision == "allow" {
		return "development-staging-ready"
	}
	if mode == "production" && installDecision == "allow" {
		return "production-activation-ready"
	}
	return mode + "-blocked"
}

func desktopActivationPreflightSummary(mode string, installDecision string) string {
	if mode == "development" && installDecision == "allow" {
		return "Runtime can hand the activation bundle to a staging installer, but this preview does not write files or enable launch."
	}
	if mode == "production" && installDecision == "allow" {
		return "Runtime production activation preflight passed, but this preview still does not write files or enable launch."
	}
	return "Runtime can render KDE activation materials, but activation is blocked until recipe trust and install gates pass."
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		if seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}
