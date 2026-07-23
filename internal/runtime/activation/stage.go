package activation

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const (
	applicationsDir      = "usr/share/applications"
	serviceMenusDir      = "usr/share/kio/servicemenus"
	manifestsDir         = "usr/share/xnix/compatibility/manifests"
	receiptsDir          = "usr/share/xnix/compatibility/activation-receipts"
	launcherArtifactsDir = "usr/share/xnix/compatibility/launcher-artifacts"
	launcherBinDir       = "usr/local/bin"
	managedLauncherName  = "xnix-compat-launch"
	dolphinServiceMenu   = "xnix-open-with-compatibility.desktop"
	stageSchemaVersion   = "xnix.runtime.desktop_activation_stage.v1"
	receiptSchemaVersion = "xnix.runtime.desktop_activation_receipt.v1"
)

type StageRequest struct {
	Root                  string
	Mode                  string
	Plan                  appidentity.Plan
	ManagedLauncherBinary string
}

type StageResult struct {
	SchemaVersion               string       `json:"schema_version"`
	RequestType                 string       `json:"request_type"`
	StageType                   string       `json:"stage_type"`
	Desktop                     string       `json:"desktop"`
	Source                      string       `json:"source"`
	ApplicationID               string       `json:"application_id"`
	DisplayName                 string       `json:"display_name"`
	DesktopFile                 string       `json:"desktop_file"`
	InstallMode                 string       `json:"install_mode"`
	PreflightDecision           string       `json:"preflight_decision"`
	WrittenFiles                []StagedFile `json:"written_files"`
	WrittenFileIDs              []string     `json:"written_file_ids"`
	WrittenFileCount            int          `json:"written_file_count"`
	ReceiptFileID               string       `json:"receipt_file_id"`
	RuntimeOwned                bool         `json:"runtime_owned"`
	GoRuntimeBacked             bool         `json:"go_runtime_backed"`
	KDEPolicyOwner              bool         `json:"kde_policy_owner"`
	StagingRootRequired         bool         `json:"staging_root_required"`
	StagingRootPathExposed      bool         `json:"staging_root_path_exposed"`
	HostRootAllowed             bool         `json:"host_root_allowed"`
	FileWritesPerformed         bool         `json:"file_writes_performed"`
	DesktopFilesWritten         bool         `json:"desktop_files_written"`
	MIMEAppsWritten             bool         `json:"mimeapps_written"`
	ManifestWritten             bool         `json:"manifest_written"`
	ReceiptWritten              bool         `json:"receipt_written"`
	RollbackReceiptWritten      bool         `json:"rollback_receipt_written"`
	SettingsPersisted           bool         `json:"settings_persisted"`
	NotificationsSent           bool         `json:"notifications_sent"`
	TaskManagerEntryActive      bool         `json:"task_manager_entry_active"`
	KWinRuleApplied             bool         `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled       bool         `json:"live_tray_bridge_enabled"`
	LaunchEnabled               bool         `json:"launch_enabled"`
	BackendLaunchEnabled        bool         `json:"backend_launch_enabled"`
	ExecutionStarted            bool         `json:"execution_started"`
	HostRootModified            bool         `json:"host_root_modified"`
	NetworkRequired             bool         `json:"network_required"`
	PrivilegedContainerRequired bool         `json:"privileged_container_required"`
	BackendDetailsExposed       bool         `json:"backend_details_exposed"`
	BlockedActions              []string     `json:"blocked_actions"`
	DesktopSafeSummary          string       `json:"desktop_safe_summary"`
}

type StagedFile struct {
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	EntryPoint            string `json:"entry_point"`
	RelativePath          string `json:"relative_path"`
	Mode                  string `json:"mode"`
	SHA256                string `json:"sha256"`
	ContentSource         string `json:"content_source"`
	Written               bool   `json:"written"`
	HostRootModified      bool   `json:"host_root_modified"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

type stageArtifact struct {
	file    StagedFile
	content string
}

type manifestFile struct {
	SchemaVersion string       `json:"schema_version"`
	ManifestType  string       `json:"manifest_type"`
	Desktop       string       `json:"desktop"`
	Application   manifestApp  `json:"application"`
	Files         []StagedFile `json:"files"`
	Safety        safety       `json:"safety"`
}

type manifestApp struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	DesktopFile string   `json:"desktop_file"`
	MIMETypes   []string `json:"mime_types"`
}

type receiptFile struct {
	SchemaVersion string       `json:"schema_version"`
	ReceiptType   string       `json:"receipt_type"`
	ApplicationID string       `json:"application_id"`
	Installed     []StagedFile `json:"installed"`
	Rollback      rollback     `json:"rollback"`
	Safety        safety       `json:"safety"`
}

type rollback struct {
	Command                string `json:"command"`
	RequiresMatchingSHA256 bool   `json:"requires_matching_sha256"`
	HostRootModified       bool   `json:"host_root_modified"`
}

type safety struct {
	RuntimeOwned          bool `json:"runtime_owned"`
	HostRootModified      bool `json:"host_root_modified"`
	BackendDetailsExposed bool `json:"backend_details_exposed"`
}

type managedLauncherArtifactFile struct {
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

func Stage(req StageRequest) (StageResult, error) {
	root, err := cleanStageRoot(req.Root)
	if err != nil {
		return StageResult{}, err
	}
	if err := req.Plan.ValidateSafeForDesktop(); err != nil {
		return StageResult{}, err
	}
	mode := req.Mode
	if mode == "" {
		mode = "development"
	}
	staging, err := req.Plan.DesktopActivationStagingPreview(mode)
	if err != nil {
		return StageResult{}, err
	}
	if !staging.InstallerMayProceed || !staging.StagingPlanReady {
		return StageResult{}, fmt.Errorf("desktop activation staging is blocked: %s", staging.PreflightDecision)
	}

	artifacts, err := stageArtifacts(req.Plan, req.ManagedLauncherBinary)
	if err != nil {
		return StageResult{}, err
	}
	if err := ensureNoConflicts(root, artifacts); err != nil {
		return StageResult{}, err
	}
	if err := writeArtifacts(root, artifacts); err != nil {
		return StageResult{}, err
	}

	files := make([]StagedFile, 0, len(artifacts))
	for _, artifact := range artifacts {
		files = append(files, artifact.file)
	}
	sort.Slice(files, func(i int, j int) bool {
		return files[i].RelativePath < files[j].RelativePath
	})

	return StageResult{
		SchemaVersion:               stageSchemaVersion,
		RequestType:                 "desktop-activation-stage",
		StageType:                   "kde-desktop-activation-test-root-stage",
		Desktop:                     "KDE Plasma",
		Source:                      "go-runtime-controlled-staging-writer",
		ApplicationID:               req.Plan.ApplicationID,
		DisplayName:                 req.Plan.DisplayName,
		DesktopFile:                 req.Plan.DesktopFile,
		InstallMode:                 staging.InstallMode,
		PreflightDecision:           staging.PreflightDecision,
		WrittenFiles:                files,
		WrittenFileIDs:              stagedFileIDs(files),
		WrittenFileCount:            len(files),
		ReceiptFileID:               "desktop-activation-receipt",
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StagingRootRequired:         true,
		StagingRootPathExposed:      false,
		HostRootAllowed:             false,
		FileWritesPerformed:         true,
		DesktopFilesWritten:         true,
		MIMEAppsWritten:             true,
		ManifestWritten:             true,
		ReceiptWritten:              true,
		RollbackReceiptWritten:      true,
		SettingsPersisted:           false,
		NotificationsSent:           false,
		TaskManagerEntryActive:      false,
		KWinRuleApplied:             false,
		LiveTrayBridgeEnabled:       false,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ExecutionStarted:            false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"write outside the provided staging root",
			"overwrite existing desktop activation files",
			"refresh KDE service cache from staging",
			"grant Runtime launch approval from staging",
			"start compatibility backend from staging",
			"mutate the host root from staging",
		},
		DesktopSafeSummary: "Runtime staged KDE desktop activation artifacts inside the provided test root without exposing the root path, mutating the host root, or enabling launch.",
	}, nil
}

func stageArtifacts(plan appidentity.Plan, managedLauncherBinary string) ([]stageArtifact, error) {
	desktopEntry, err := plan.RenderDesktopEntry()
	if err != nil {
		return nil, err
	}
	mimeapps, err := plan.RenderMIMEApps()
	if err != nil {
		return nil, err
	}

	initial := []stageArtifact{
		newArtifact("desktop-entry", "desktop-entry", "launcher", applicationsDir+"/"+plan.DesktopFile, desktopEntry, "desktop-entry-preview"),
		newArtifact("dolphin-service-menu", "dolphin-service-menu", "file-manager", serviceMenusDir+"/"+dolphinServiceMenu, renderDolphinServiceMenu(), "dolphin-service-menu-preview"),
		newArtifact("mimeapps-list", "mimeapps-list", "file-manager", applicationsDir+"/mimeapps.list", mimeapps, "mimeapps-preview"),
		newArtifact("managed-launcher-artifact", "managed-launcher-artifact", "launcher", launcherArtifactsDir+"/"+managedLauncherName+".json", renderManagedLauncherArtifact(managedLauncherBinary != ""), "cmd/xnix-compat-launch"),
	}
	if managedLauncherBinary != "" {
		launcherExecutable, err := newExecutableArtifact(
			"managed-launcher-executable",
			"managed-launcher-executable",
			"launcher",
			launcherBinDir+"/"+managedLauncherName,
			managedLauncherBinary,
			"cmd/xnix-compat-launch",
		)
		if err != nil {
			return nil, err
		}
		initial = append(initial, launcherExecutable)
	}
	manifestContent, err := renderManifest(plan, stagedFiles(initial))
	if err != nil {
		return nil, err
	}
	initial = append(initial, newArtifact("desktop-integration-manifest", "desktop-integration-manifest", "all", manifestsDir+"/"+plan.ApplicationID+".json", manifestContent, "desktop-activation-stage"))
	receiptContent, err := renderReceipt(plan, stagedFiles(initial))
	if err != nil {
		return nil, err
	}
	initial = append(initial, newArtifact("desktop-activation-receipt", "desktop-activation-receipt", "rollback", receiptsDir+"/"+plan.ApplicationID+".json", receiptContent, "desktop-activation-stage"))
	return initial, nil
}

func newArtifact(id string, kind string, entryPoint string, relativePath string, content string, source string) stageArtifact {
	return newArtifactWithMode(id, kind, entryPoint, relativePath, content, source, "0644")
}

func newExecutableArtifact(id string, kind string, entryPoint string, relativePath string, sourcePath string, source string) (stageArtifact, error) {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return stageArtifact{}, fmt.Errorf("read managed launcher binary: %w", err)
	}
	if len(data) == 0 {
		return stageArtifact{}, errors.New("managed launcher binary must not be empty")
	}
	return newArtifactWithMode(id, kind, entryPoint, relativePath, string(data), source, "0755"), nil
}

func newArtifactWithMode(id string, kind string, entryPoint string, relativePath string, content string, source string, mode string) stageArtifact {
	return stageArtifact{
		file: StagedFile{
			ID:                    id,
			Kind:                  kind,
			EntryPoint:            entryPoint,
			RelativePath:          relativePath,
			Mode:                  mode,
			SHA256:                sha256Hex(content),
			ContentSource:         source,
			Written:               true,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
		content: content,
	}
}

func renderDolphinServiceMenu() string {
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

func renderManagedLauncherArtifact(executableStaged bool) string {
	content, err := encodeJSON(managedLauncherArtifactFile{
		SchemaVersion:          "xnix.runtime.managed_launcher_artifact.v1",
		ArtifactType:           "managed-launcher-artifact",
		Command:                managedLauncherName,
		SourcePackage:          "cmd/xnix-compat-launch",
		BuildOutput:            launcherBinDir + "/" + managedLauncherName,
		StagedExecutable:       launcherBinDir + "/" + managedLauncherName,
		DesktopExecUsesCommand: true,
		RuntimeMethod:          "PreviewKnownPortableLaunchBridge",
		DispatchGate:           "managed-known-app-guest-smoke",
		BinaryCopied:           executableStaged,
		ExecutableStaged:       executableStaged,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		ExecutionStarted:       false,
		HostRootModified:       false,
		BackendDetailsExposed:  false,
		DesktopSafeSummary:     "The managed launcher command is provided by the Go Runtime build and remains gated before execution.",
	})
	if err != nil {
		return "{}\n"
	}
	return content
}

func renderManifest(plan appidentity.Plan, files []StagedFile) (string, error) {
	manifest := manifestFile{
		SchemaVersion: "xnix.runtime.desktop_activation_stage_manifest.v1",
		ManifestType:  "desktop-integration",
		Desktop:       "KDE Plasma",
		Application: manifestApp{
			ID:          plan.ApplicationID,
			Name:        plan.DisplayName,
			DesktopFile: plan.DesktopFile,
			MIMETypes:   append([]string{}, plan.MIMETypes...),
		},
		Files: files,
		Safety: safety{
			RuntimeOwned:          true,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
	}
	return encodeJSON(manifest)
}

func renderReceipt(plan appidentity.Plan, files []StagedFile) (string, error) {
	receipt := receiptFile{
		SchemaVersion: receiptSchemaVersion,
		ReceiptType:   "desktop-activation-receipt",
		ApplicationID: plan.ApplicationID,
		Installed:     files,
		Rollback: rollback{
			Command:                "xnix-rollback-desktop-integration",
			RequiresMatchingSHA256: true,
			HostRootModified:       false,
		},
		Safety: safety{
			RuntimeOwned:          true,
			HostRootModified:      false,
			BackendDetailsExposed: false,
		},
	}
	return encodeJSON(receipt)
}

func encodeJSON(value any) (string, error) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func cleanStageRoot(root string) (string, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return "", errors.New("desktop activation staging root is required")
	}
	cleaned := filepath.Clean(root)
	if cleaned == string(filepath.Separator) {
		return "", errors.New("refusing to stage desktop activation into filesystem root")
	}
	if err := os.MkdirAll(cleaned, 0o755); err != nil {
		return "", fmt.Errorf("prepare staging root: %w", err)
	}
	return cleaned, nil
}

func ensureNoConflicts(root string, artifacts []stageArtifact) error {
	for _, artifact := range artifacts {
		target, err := targetPath(root, artifact.file.RelativePath)
		if err != nil {
			return err
		}
		if _, err := os.Stat(target); err == nil {
			return fmt.Errorf("refusing to overwrite existing staged file: %s", artifact.file.RelativePath)
		} else if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("inspect staged file %s: %w", artifact.file.RelativePath, err)
		}
	}
	return nil
}

func writeArtifacts(root string, artifacts []stageArtifact) error {
	for _, artifact := range artifacts {
		target, err := targetPath(root, artifact.file.RelativePath)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return fmt.Errorf("prepare staged directory %s: %w", artifact.file.RelativePath, err)
		}
		mode, err := parseFileMode(artifact.file.Mode)
		if err != nil {
			return fmt.Errorf("parse staged file mode %s: %w", artifact.file.RelativePath, err)
		}
		file, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
		if err != nil {
			return fmt.Errorf("write staged file %s: %w", artifact.file.RelativePath, err)
		}
		if _, err := file.WriteString(artifact.content); err != nil {
			_ = file.Close()
			return fmt.Errorf("write staged file %s: %w", artifact.file.RelativePath, err)
		}
		if err := file.Close(); err != nil {
			return fmt.Errorf("close staged file %s: %w", artifact.file.RelativePath, err)
		}
		if err := os.Chmod(target, mode); err != nil {
			return fmt.Errorf("chmod staged file %s: %w", artifact.file.RelativePath, err)
		}
	}
	return nil
}

func parseFileMode(mode string) (fs.FileMode, error) {
	switch mode {
	case "0644":
		return 0o644, nil
	case "0755":
		return 0o755, nil
	default:
		return 0, fmt.Errorf("unsupported mode %q", mode)
	}
}

func targetPath(root string, relativePath string) (string, error) {
	if filepath.IsAbs(relativePath) {
		return "", fmt.Errorf("staged path must be relative: %s", relativePath)
	}
	cleaned := filepath.Clean(relativePath)
	if cleaned == "." || strings.HasPrefix(cleaned, ".."+string(filepath.Separator)) || cleaned == ".." {
		return "", fmt.Errorf("staged path escapes root: %s", relativePath)
	}
	target := filepath.Join(root, cleaned)
	rootWithSeparator := root + string(filepath.Separator)
	if target != root && !strings.HasPrefix(target, rootWithSeparator) {
		return "", fmt.Errorf("staged path escapes root: %s", relativePath)
	}
	return target, nil
}

func stagedFiles(artifacts []stageArtifact) []StagedFile {
	files := make([]StagedFile, 0, len(artifacts))
	for _, artifact := range artifacts {
		files = append(files, artifact.file)
	}
	return files
}

func stagedFileIDs(files []StagedFile) []string {
	ids := make([]string, 0, len(files))
	for _, file := range files {
		ids = append(ids, file.ID)
	}
	sort.Strings(ids)
	return ids
}

func sha256Hex(content string) string {
	sum := sha256.Sum256([]byte(content))
	return hex.EncodeToString(sum[:])
}
