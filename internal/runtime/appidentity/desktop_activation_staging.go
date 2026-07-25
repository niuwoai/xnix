package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

const (
	desktopActivationApplicationsDir = "usr/share/applications"
	desktopActivationServiceMenusDir = "usr/share/kio/servicemenus"
	desktopActivationManifestsDir    = "usr/share/xnix/compatibility/manifests"
	desktopActivationReceiptsDir     = "usr/share/xnix/compatibility/activation-receipts"
	desktopActivationLauncherDir     = "usr/share/xnix/compatibility/launcher-artifacts"
	desktopActivationRecipesDir      = "usr/share/xnix/compatibility/recipes"
	dolphinServiceMenuFileName       = "xnix-open-with-compatibility.desktop"
)

type DesktopActivationStagingPreview struct {
	SchemaVersion                string                            `json:"schema_version"`
	RequestType                  string                            `json:"request_type"`
	StagingType                  string                            `json:"staging_type"`
	Source                       string                            `json:"source"`
	Desktop                      string                            `json:"desktop"`
	RuntimeMethod                string                            `json:"runtime_method"`
	ApplicationID                string                            `json:"application_id"`
	DisplayName                  string                            `json:"display_name"`
	Icon                         string                            `json:"icon"`
	DesktopFile                  string                            `json:"desktop_file"`
	InstallMode                  string                            `json:"install_mode"`
	PreflightDecision            string                            `json:"preflight_decision"`
	StagingState                 string                            `json:"staging_state"`
	Preflight                    DesktopActivationStagingPreflight `json:"preflight"`
	PlannedFiles                 []DesktopActivationStagedFile     `json:"planned_files"`
	PlannedFileIDs               []string                          `json:"planned_file_ids"`
	PlannedFileCount             int                               `json:"planned_file_count"`
	ReceiptFileID                string                            `json:"receipt_file_id"`
	ActivatedEntryPoints         []string                          `json:"activated_entry_points"`
	ActivatedEntryPointCount     int                               `json:"activated_entry_point_count"`
	InstallerCommandPreview      []string                          `json:"installer_command_preview"`
	RuntimeOwned                 bool                              `json:"runtime_owned"`
	GoRuntimeBacked              bool                              `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                              `json:"kde_policy_owner"`
	UserVisible                  bool                              `json:"user_visible"`
	DevelopmentStagingEligible   bool                              `json:"development_staging_eligible"`
	ProductionActivationEligible bool                              `json:"production_activation_eligible"`
	InstallerMayProceed          bool                              `json:"installer_may_proceed"`
	StagingPlanReady             bool                              `json:"staging_plan_ready"`
	StagingRootRequired          bool                              `json:"staging_root_required"`
	StagingRootPathExposed       bool                              `json:"staging_root_path_exposed"`
	HostRootAllowed              bool                              `json:"host_root_allowed"`
	FileWritesPerformed          bool                              `json:"file_writes_performed"`
	DesktopFilesWritten          bool                              `json:"desktop_files_written"`
	MIMEAppsWritten              bool                              `json:"mimeapps_written"`
	ManifestWritten              bool                              `json:"manifest_written"`
	ReceiptWritten               bool                              `json:"receipt_written"`
	RollbackReceiptPlanned       bool                              `json:"rollback_receipt_planned"`
	RollbackReceiptWritten       bool                              `json:"rollback_receipt_written"`
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

type DesktopActivationStagingPreflight struct {
	RequestType                  string `json:"request_type"`
	PreflightDecision            string `json:"preflight_decision"`
	InstallMode                  string `json:"install_mode"`
	InstallGateDecision          string `json:"install_gate_decision"`
	DevelopmentStagingEligible   bool   `json:"development_staging_eligible"`
	ProductionActivationEligible bool   `json:"production_activation_eligible"`
	InstallerMayProceed          bool   `json:"installer_may_proceed"`
	HostRootAllowed              bool   `json:"host_root_allowed"`
}

type DesktopActivationStagedFile struct {
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	EntryPoint            string `json:"entry_point"`
	RelativePath          string `json:"relative_path"`
	Mode                  string `json:"mode"`
	SHA256                string `json:"sha256"`
	ContentSource         string `json:"content_source"`
	PlannedForStaging     bool   `json:"planned_for_staging"`
	Written               bool   `json:"written"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type desktopActivationManifestPreview struct {
	SchemaVersion string                              `json:"schema_version"`
	ManifestType  string                              `json:"manifest_type"`
	Desktop       string                              `json:"desktop"`
	Application   desktopActivationManifestAppPreview `json:"application"`
	EntryPoints   []string                            `json:"entry_points"`
	ArtifactIDs   []string                            `json:"artifact_ids"`
	Safety        desktopActivationManifestSafety     `json:"safety"`
}

type desktopActivationManifestAppPreview struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	DesktopFile string   `json:"desktop_file"`
	MIMETypes   []string `json:"mime_types"`
}

type desktopActivationManifestSafety struct {
	RuntimeOwned          bool `json:"runtime_owned"`
	HostRootModified      bool `json:"host_root_modified"`
	BackendDetailsExposed bool `json:"backend_details_exposed"`
}

type desktopActivationReceiptPreview struct {
	SchemaVersion string                           `json:"schema_version"`
	ReceiptType   string                           `json:"receipt_type"`
	ApplicationID string                           `json:"application_id"`
	Installed     []DesktopActivationStagedFile    `json:"installed"`
	Rollback      desktopActivationReceiptRollback `json:"rollback"`
}

type desktopActivationReceiptRollback struct {
	Command                string `json:"command"`
	RequiresMatchingSHA256 bool   `json:"requires_matching_sha256"`
}

type desktopActivationManagedLauncherArtifact struct {
	SchemaVersion          string `json:"schema_version"`
	ArtifactType           string `json:"artifact_type"`
	Command                string `json:"command"`
	SourcePackage          string `json:"source_package"`
	BuildOutput            string `json:"build_output"`
	StagedExecutable       string `json:"staged_executable"`
	DesktopExecUsesCommand bool   `json:"desktop_exec_uses_command"`
	RuntimeMethod          string `json:"runtime_method"`
	DispatchGate           string `json:"dispatch_gate"`
	BinaryCopied           bool   `json:"binary_copied"`
	ExecutableStaged       bool   `json:"executable_staged"`
	RuntimeOwned           bool   `json:"runtime_owned"`
	GoRuntimeBacked        bool   `json:"go_runtime_backed"`
	KDEPolicyOwner         bool   `json:"kde_policy_owner"`
	ExecutionStarted       bool   `json:"execution_started"`
	HostRootModified       bool   `json:"host_root_modified"`
	BackendDetailsExposed  bool   `json:"backend_details_exposed"`
	DesktopSafeSummary     string `json:"desktop_safe_summary"`
}

func (plan Plan) DesktopActivationStagingPreview(mode string) (DesktopActivationStagingPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopActivationStagingPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return DesktopActivationStagingPreview{}, errors.New("desktop activation staging requires single-line identity fields")
		}
	}
	preflight, err := plan.DesktopActivationPreflightPreview(mode)
	if err != nil {
		return DesktopActivationStagingPreview{}, err
	}
	bundle, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		return DesktopActivationStagingPreview{}, err
	}
	files, err := desktopActivationStagedFiles(plan, bundle)
	if err != nil {
		return DesktopActivationStagingPreview{}, err
	}
	activatedEntryPoints := []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "unified-settings"}
	stagingPlanReady := preflight.InstallerMayProceed

	preview := DesktopActivationStagingPreview{
		SchemaVersion:                "xnix.runtime.desktop_activation_staging.v1",
		RequestType:                  "desktop-activation-staging-preview",
		StagingType:                  "kde-activation-staging-plan",
		Source:                       "desktop-activation-preflight-preview+desktop-activation-bundle-preview",
		Desktop:                      "KDE Plasma",
		RuntimeMethod:                "GetDesktopActivationStaging",
		ApplicationID:                plan.ApplicationID,
		DisplayName:                  plan.DisplayName,
		Icon:                         plan.Icon,
		DesktopFile:                  plan.DesktopFile,
		InstallMode:                  preflight.InstallMode,
		PreflightDecision:            preflight.PreflightDecision,
		StagingState:                 desktopActivationStagingState(stagingPlanReady),
		Preflight:                    desktopActivationStagingPreflight(preflight),
		PlannedFiles:                 files,
		PlannedFileIDs:               desktopActivationStagedFileIDs(files),
		PlannedFileCount:             len(files),
		ReceiptFileID:                "desktop-activation-receipt",
		ActivatedEntryPoints:         activatedEntryPoints,
		ActivatedEntryPointCount:     len(activatedEntryPoints),
		InstallerCommandPreview:      []string{"xnix-install-desktop-integration", "--desktop-entry-source", "runtime-go", "--file-association-source", "runtime-go"},
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		UserVisible:                  true,
		DevelopmentStagingEligible:   preflight.DevelopmentStagingEligible,
		ProductionActivationEligible: preflight.ProductionActivationEligible,
		InstallerMayProceed:          preflight.InstallerMayProceed,
		StagingPlanReady:             stagingPlanReady,
		StagingRootRequired:          true,
		StagingRootPathExposed:       false,
		HostRootAllowed:              false,
		FileWritesPerformed:          false,
		DesktopFilesWritten:          false,
		MIMEAppsWritten:              false,
		ManifestWritten:              false,
		ReceiptWritten:               false,
		RollbackReceiptPlanned:       true,
		RollbackReceiptWritten:       false,
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
			"write staged files from staging preview",
			"write activation receipt from staging preview",
			"install into host root from staging preview",
			"overwrite existing MIME defaults from staging preview",
			"activate task manager entry before a live session",
			"apply KWin rule before Runtime launch approval",
			"enable live tray bridge before Runtime launch approval",
			"start compatibility backend from staging preview",
			"mutate host root during activation staging preview",
			"expose raw backend command to desktop shell",
		},
		DesktopSafeSummary: desktopActivationStagingSummary(stagingPlanReady),
	}
	if err := validateNoBackendTerms(preview, "desktop activation staging preview"); err != nil {
		return DesktopActivationStagingPreview{}, err
	}
	return preview, nil
}

func desktopActivationStagedFiles(plan Plan, bundle DesktopActivationBundlePreview) ([]DesktopActivationStagedFile, error) {
	manifestContent, err := desktopActivationManifestPreviewContent(plan, bundle)
	if err != nil {
		return nil, err
	}
	files := []DesktopActivationStagedFile{
		desktopActivationStagedFile("desktop-entry", "desktop-entry", "launcher", desktopActivationApplicationsDir+"/"+plan.DesktopFile, "0644", bundle.DesktopEntryPreview, "desktop-entry-preview"),
		desktopActivationStagedFile("dolphin-service-menu", "dolphin-service-menu", "file-manager", desktopActivationServiceMenusDir+"/"+dolphinServiceMenuFileName, "0644", renderDolphinServiceMenuPreview(), "dolphin-service-menu-preview"),
	}
	if bundle.MIMEAppsPreview != "" {
		files = append(files, desktopActivationStagedFile("mimeapps-list", "mimeapps-list", "file-manager", desktopActivationApplicationsDir+"/mimeapps.list", "0644", bundle.MIMEAppsPreview, "mimeapps-preview"))
	}
	if plan.ContainerGUISmoke != (ContainerGUISmokeHints{}) {
		recipeContent, registryContent, err := desktopActivationRecipeRegistryContents(plan)
		if err != nil {
			return nil, err
		}
		files = append(files,
			desktopActivationStagedFile("recipe-registry", "recipe-registry", "launcher", desktopActivationRecipesDir+"/registry.json", "0644", registryContent, "desktop-activation-staging-preview"),
			desktopActivationStagedFile("application-recipe", "application-recipe", "launcher", desktopActivationRecipesDir+"/"+plan.ApplicationID+".json", "0644", recipeContent, "desktop-activation-staging-preview"),
		)
	}
	files = append(files,
		desktopActivationStagedFile("desktop-integration-manifest", "desktop-integration-manifest", "all", desktopActivationManifestsDir+"/"+plan.ApplicationID+".json", "0644", manifestContent, "desktop-activation-staging-preview"),
		desktopActivationStagedFile("managed-launcher-artifact", "managed-launcher-artifact", "launcher", desktopActivationLauncherDir+"/xnix-compat-launch.json", "0644", renderManagedLauncherArtifactPreview(), "cmd/xnix-compat-launch"),
	)
	receiptContent, err := desktopActivationReceiptPreviewContent(plan, files)
	if err != nil {
		return nil, err
	}
	files = append(files, desktopActivationStagedFile("desktop-activation-receipt", "desktop-activation-receipt", "rollback", desktopActivationReceiptsDir+"/"+plan.ApplicationID+".json", "0644", receiptContent, "desktop-activation-staging-preview"))
	return files, nil
}

func desktopActivationRecipeRegistryContents(plan Plan) (string, string, error) {
	recipe := Recipe{
		ID:                  plan.ApplicationID,
		Name:                plan.DisplayName,
		Version:             plan.ApplicationVersion,
		Icon:                plan.Icon,
		Mode:                plan.RecipeMode,
		SupportedExtensions: desktopActivationExtensionsFromMIMETypes(plan.MIMETypes),
		ContainerGUISmoke:   plan.ContainerGUISmoke,
	}
	recipeData, err := json.MarshalIndent(recipe, "", "  ")
	if err != nil {
		return "", "", err
	}
	recipeContent := string(recipeData) + "\n"
	registry := Registry{
		SchemaVersion: 1,
		RegistryName:  "xnix-desktop-activation",
		Recipes: []RegistryEntry{
			{
				ID:              plan.ApplicationID,
				Path:            plan.ApplicationID + ".json",
				SHA256:          sha256Hex(recipeContent),
				SignatureStatus: plan.RecipeSignatureStatus,
			},
		},
	}
	if registry.Recipes[0].SignatureStatus == "" {
		registry.Recipes[0].SignatureStatus = "development-only"
	}
	registryData, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		return "", "", err
	}
	return recipeContent, string(registryData) + "\n", nil
}

func desktopActivationExtensionsFromMIMETypes(mimeTypes []string) []string {
	extensions := make([]string, 0, len(mimeTypes))
	for _, mimeType := range mimeTypes {
		const prefix = "application/x-xnix-"
		if len(mimeType) > len(prefix) && mimeType[:len(prefix)] == prefix {
			extensions = append(extensions, "."+mimeType[len(prefix):])
		}
	}
	return extensions
}

func desktopActivationStagedFile(id string, kind string, entryPoint string, relativePath string, mode string, content string, source string) DesktopActivationStagedFile {
	return DesktopActivationStagedFile{
		ID:                    id,
		Kind:                  kind,
		EntryPoint:            entryPoint,
		RelativePath:          relativePath,
		Mode:                  mode,
		SHA256:                sha256Hex(content),
		ContentSource:         source,
		PlannedForStaging:     true,
		Written:               false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
	}
}

func renderDolphinServiceMenuPreview() string {
	return "[Desktop Entry]\n" +
		"Type=Service\n" +
		"MimeType=application/octet-stream;text/plain;\n" +
		"Actions=openWithXnixCompatibility;\n" +
		"X-KDE-ServiceTypes=KonqPopupMenu/Plugin\n" +
		"X-KDE-Priority=TopLevel\n" +
		"\n" +
		"[Desktop Action openWithXnixCompatibility]\n" +
		"Name=Open with Xnix Compatibility\n" +
		"Icon=preferences-desktop\n" +
		"Exec=xnix-compat-open %U\n"
}

func renderManagedLauncherArtifactPreview() string {
	artifact := desktopActivationManagedLauncherArtifact{
		SchemaVersion:          "xnix.runtime.managed_launcher_artifact.v1",
		ArtifactType:           "managed-launcher-artifact",
		Command:                "xnix-compat-launch",
		SourcePackage:          "cmd/xnix-compat-launch",
		BuildOutput:            "usr/local/bin/xnix-compat-launch",
		StagedExecutable:       "usr/local/bin/xnix-compat-launch",
		DesktopExecUsesCommand: true,
		RuntimeMethod:          "PreviewKnownPortableLaunchBridge",
		DispatchGate:           "managed-known-app-guest-smoke",
		BinaryCopied:           false,
		ExecutableStaged:       false,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		ExecutionStarted:       false,
		HostRootModified:       false,
		BackendDetailsExposed:  false,
		DesktopSafeSummary:     "The managed launcher command is provided by the Go Runtime build and remains gated before execution.",
	}
	data, err := json.MarshalIndent(artifact, "", "  ")
	if err != nil {
		return "{}\n"
	}
	return string(data) + "\n"
}

func desktopActivationManifestPreviewContent(plan Plan, bundle DesktopActivationBundlePreview) (string, error) {
	manifest := desktopActivationManifestPreview{
		SchemaVersion: "xnix.runtime.desktop_activation_manifest_preview.v1",
		ManifestType:  "desktop-integration",
		Desktop:       "KDE Plasma",
		Application: desktopActivationManifestAppPreview{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			Icon:        plan.Icon,
			DesktopFile: plan.DesktopFile,
			MIMETypes:   plan.MIMETypes,
		},
		EntryPoints: []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "unified-settings"},
		ArtifactIDs: bundle.MaterialIDs,
		Safety: desktopActivationManifestSafety{
			RuntimeOwned:          true,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func desktopActivationReceiptPreviewContent(plan Plan, installed []DesktopActivationStagedFile) (string, error) {
	receipt := desktopActivationReceiptPreview{
		SchemaVersion: "xnix.runtime.desktop_activation_receipt_preview.v1",
		ReceiptType:   "desktop-activation-receipt",
		ApplicationID: plan.ApplicationID,
		Installed:     installed,
		Rollback: desktopActivationReceiptRollback{
			Command:                "xnix-rollback-desktop-integration",
			RequiresMatchingSHA256: true,
		},
	}
	data, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func desktopActivationStagingPreflight(preflight DesktopActivationPreflightPreview) DesktopActivationStagingPreflight {
	return DesktopActivationStagingPreflight{
		RequestType:                  preflight.RequestType,
		PreflightDecision:            preflight.PreflightDecision,
		InstallMode:                  preflight.InstallMode,
		InstallGateDecision:          preflight.InstallGate.Decision,
		DevelopmentStagingEligible:   preflight.DevelopmentStagingEligible,
		ProductionActivationEligible: preflight.ProductionActivationEligible,
		InstallerMayProceed:          preflight.InstallerMayProceed,
		HostRootAllowed:              preflight.HostRootAllowed,
	}
}

func desktopActivationStagedFileIDs(files []DesktopActivationStagedFile) []string {
	ids := make([]string, 0, len(files))
	for _, file := range files {
		ids = append(ids, file.ID)
	}
	return ids
}

func desktopActivationStagingState(ready bool) string {
	if ready {
		return "staging-plan-ready"
	}
	return "staging-blocked"
}

func desktopActivationStagingSummary(ready bool) string {
	if ready {
		return "Runtime can explain the KDE activation files that a staging installer may write, but this preview performs no file writes and exposes no host path."
	}
	return "Runtime can compute KDE activation files, but staging remains blocked until activation preflight allows the installer to proceed."
}

func sha256Hex(content string) string {
	sum := sha256.Sum256([]byte(content))
	return hex.EncodeToString(sum[:])
}
