package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	ExternalWinAppBundleImportPlanSchemaVersion = "xnix.runtime.external_winapp_bundle_import_plan.v1"
	ExternalWinAppBundleImportPlanRequestType   = "external-winapp-bundle-import-plan-preview"
)

type ExternalWinAppBundleImportPlanRequest struct {
	Version                string
	BundleRoot             string
	ExecutableRelativePath string
	AppID                  string
	DisplayName            string
	AppVersion             string
}

type ExternalWinAppBundleManifestEntry struct {
	RelativePath string `json:"relative_path"`
	SizeBytes    int64  `json:"size_bytes"`
	SHA256       string `json:"sha256"`
}

type ExternalWinAppBundleImportPlan struct {
	Version                           string `json:"version"`
	SchemaVersion                     string `json:"schema_version"`
	RequestType                       string `json:"request_type"`
	PlanType                          string `json:"plan_type"`
	Source                            string `json:"source"`
	RuntimeMethod                     string `json:"runtime_method"`
	ReadMethod                        string `json:"read_method"`
	ApplicationID                     string `json:"application_id"`
	DisplayName                       string `json:"display_name"`
	AppVersion                        string `json:"app_version"`
	BundleKind                        string `json:"bundle_kind"`
	ExecutableName                    string `json:"executable_name"`
	ExecutableRelativePath            string `json:"executable_relative_path"`
	ExecutableSHA256                  string `json:"executable_sha256"`
	ExecutableSizeBytes               int64  `json:"executable_size_bytes"`
	BundleManifestSHA256              string `json:"bundle_manifest_sha256"`
	BundleFileCount                   int    `json:"bundle_file_count"`
	BundleDirectoryCount              int    `json:"bundle_directory_count"`
	SidecarFileCount                  int    `json:"sidecar_file_count"`
	SidecarDirectoryCount             int    `json:"sidecar_directory_count"`
	ManifestEntryCount                int    `json:"manifest_entry_count"`
	WindowsExecutableValidated        bool   `json:"windows_executable_validated"`
	WindowsExecutableMZHeaderVerified bool   `json:"windows_executable_mz_header_verified"`
	DirectoryLayoutValidated          bool   `json:"directory_layout_validated"`
	SidecarDirectorySupported         bool   `json:"sidecar_directory_supported"`
	SingleExecutableImportCompatible  bool   `json:"single_executable_import_compatible"`
	BundleRootConsumed                bool   `json:"bundle_root_consumed"`
	PortableBundleImportReady         bool   `json:"portable_bundle_import_ready"`
	RuntimeOwned                      bool   `json:"runtime_owned"`
	GoRuntimeBacked                   bool   `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool   `json:"kde_policy_owner"`
	DesktopPageRenderable             bool   `json:"desktop_page_renderable"`
	LaunchEnabled                     bool   `json:"launch_enabled"`
	BackendLaunchEnabled              bool   `json:"backend_launch_enabled"`
	ActionExecutionEnabled            bool   `json:"action_execution_enabled"`
	BundleRootPathExposed             bool   `json:"bundle_root_path_exposed"`
	RawExecutablePathExposed          bool   `json:"raw_executable_path_exposed"`
	BackendDetailsExposed             bool   `json:"backend_details_exposed"`
	HostRootModified                  bool   `json:"host_root_modified"`
	NetworkRequired                   bool   `json:"network_required"`
	PrivilegedContainerRequired       bool   `json:"privileged_container_required"`
	DockerSocketMounted               bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool   `json:"broad_host_mount_required"`
	PackageManagerInvoked             bool   `json:"package_manager_invoked"`
	DesktopSafeSummary                string `json:"desktop_safe_summary"`
}

type externalWinAppBundleFile struct {
	relativePath string
	sizeBytes    int64
	sha256       string
}

func PreviewExternalWinAppBundleImportPlan(request ExternalWinAppBundleImportPlanRequest) (ExternalWinAppBundleImportPlan, error) {
	version := strings.TrimSpace(request.Version)
	if version == "" {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires a version")
	}
	if !versionPattern.MatchString(version) {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan version must be semantic version format")
	}
	bundleRoot := strings.TrimSpace(request.BundleRoot)
	if bundleRoot == "" {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires --bundle-root")
	}
	bundleRoot, err := filepath.Abs(bundleRoot)
	if err != nil {
		return ExternalWinAppBundleImportPlan{}, fmt.Errorf("resolve external Windows app bundle root: %w", err)
	}
	rootInfo, err := os.Stat(bundleRoot)
	if err != nil {
		return ExternalWinAppBundleImportPlan{}, fmt.Errorf("stat external Windows app bundle root: %w", err)
	}
	if !rootInfo.IsDir() {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle root must be a directory")
	}
	executableRelativePath, err := cleanExternalWinAppBundleRelativePath(request.ExecutableRelativePath)
	if err != nil {
		return ExternalWinAppBundleImportPlan{}, err
	}
	executableName := filepath.Base(filepath.FromSlash(executableRelativePath))
	if !safeExternalExecutableName(executableName) {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires a safe .exe file name")
	}
	appID := strings.TrimSpace(request.AppID)
	if !idPattern.MatchString(appID) {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires a reverse-DNS app id")
	}
	displayName := strings.TrimSpace(request.DisplayName)
	if !singleLine(displayName) {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires a single-line display name")
	}
	appVersion := strings.TrimSpace(request.AppVersion)
	if appVersion == "" {
		appVersion = version
	}
	if !versionPattern.MatchString(appVersion) {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan app version must be semantic version format")
	}

	files, directoryCount, err := scanExternalWinAppBundle(bundleRoot)
	if err != nil {
		return ExternalWinAppBundleImportPlan{}, err
	}
	if len(files) == 0 {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires at least one file")
	}
	var executable externalWinAppBundleFile
	foundExecutable := false
	for _, file := range files {
		if file.relativePath == executableRelativePath {
			executable = file
			foundExecutable = true
			break
		}
	}
	if !foundExecutable {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan executable was not found inside bundle root")
	}
	executablePath := filepath.Join(bundleRoot, filepath.FromSlash(executableRelativePath))
	executableContent, err := os.ReadFile(executablePath)
	if err != nil {
		return ExternalWinAppBundleImportPlan{}, fmt.Errorf("read external Windows app bundle executable: %w", err)
	}
	if len(executableContent) < 2 || executableContent[0] != 'M' || executableContent[1] != 'Z' {
		return ExternalWinAppBundleImportPlan{}, errors.New("external Windows app bundle import plan requires an MZ executable")
	}
	manifestSHA256 := externalWinAppBundleManifestSHA256(files)
	sidecarFileCount := len(files) - 1
	sidecarDirectoryCount := directoryCount
	if sidecarDirectoryCount < 0 {
		sidecarDirectoryCount = 0
	}
	return ExternalWinAppBundleImportPlan{
		Version:                           version,
		SchemaVersion:                     ExternalWinAppBundleImportPlanSchemaVersion,
		RequestType:                       ExternalWinAppBundleImportPlanRequestType,
		PlanType:                          "external-windows-app-portable-bundle-import-plan",
		Source:                            "go-runtime-external-winapp-bundle-import",
		RuntimeMethod:                     "PreviewExternalWinAppBundleImportPlan",
		ReadMethod:                        "GetExternalWinAppBundleImportPlan",
		ApplicationID:                     appID,
		DisplayName:                       displayName,
		AppVersion:                        appVersion,
		BundleKind:                        "portable-directory",
		ExecutableName:                    executableName,
		ExecutableRelativePath:            executableRelativePath,
		ExecutableSHA256:                  executable.sha256,
		ExecutableSizeBytes:               executable.sizeBytes,
		BundleManifestSHA256:              manifestSHA256,
		BundleFileCount:                   len(files),
		BundleDirectoryCount:              directoryCount,
		SidecarFileCount:                  sidecarFileCount,
		SidecarDirectoryCount:             sidecarDirectoryCount,
		ManifestEntryCount:                len(files),
		WindowsExecutableValidated:        true,
		WindowsExecutableMZHeaderVerified: true,
		DirectoryLayoutValidated:          true,
		SidecarDirectorySupported:         true,
		SingleExecutableImportCompatible:  len(files) == 1,
		BundleRootConsumed:                true,
		PortableBundleImportReady:         true,
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		DesktopPageRenderable:             true,
		LaunchEnabled:                     false,
		BackendLaunchEnabled:              false,
		ActionExecutionEnabled:            false,
		BundleRootPathExposed:             false,
		RawExecutablePathExposed:          false,
		BackendDetailsExposed:             false,
		HostRootModified:                  false,
		NetworkRequired:                   false,
		PrivilegedContainerRequired:       false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		PackageManagerInvoked:             false,
		DesktopSafeSummary:                displayName + " portable directory can be imported as a Runtime-managed bundle with sidecar files accounted for by a deterministic manifest digest; desktop launch remains gated.",
	}, nil
}

func cleanExternalWinAppBundleRelativePath(value string) (string, error) {
	value = filepath.ToSlash(strings.TrimSpace(value))
	if value == "" {
		return "", errors.New("external Windows app bundle import plan requires --executable-relative-path")
	}
	if filepath.IsAbs(value) || strings.HasPrefix(value, "/") || strings.Contains(value, "\x00") {
		return "", errors.New("external Windows app bundle import plan executable relative path is unsafe")
	}
	clean := filepath.ToSlash(filepath.Clean(filepath.FromSlash(value)))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", errors.New("external Windows app bundle import plan executable relative path escapes bundle root")
	}
	return clean, nil
}

func scanExternalWinAppBundle(root string) ([]externalWinAppBundleFile, int, error) {
	var files []externalWinAppBundleFile
	directoryCount := 0
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("external Windows app bundle import plan rejects symlink: %s", entry.Name())
		}
		relativePath, err := externalWinAppBundleRelativePath(root, path)
		if err != nil {
			return err
		}
		if relativePath == "." {
			return nil
		}
		if entry.IsDir() {
			directoryCount++
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("external Windows app bundle import plan rejects non-regular file: %s", relativePath)
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read external Windows app bundle file %s: %w", relativePath, err)
		}
		sum := sha256.Sum256(content)
		files = append(files, externalWinAppBundleFile{
			relativePath: relativePath,
			sizeBytes:    int64(len(content)),
			sha256:       hex.EncodeToString(sum[:]),
		})
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].relativePath < files[j].relativePath
	})
	return files, directoryCount, nil
}

func externalWinAppBundleRelativePath(root string, path string) (string, error) {
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return "", fmt.Errorf("resolve external Windows app bundle relative path: %w", err)
	}
	relativePath = filepath.ToSlash(relativePath)
	if relativePath == "" || filepath.IsAbs(relativePath) || relativePath == ".." || strings.HasPrefix(relativePath, "../") || strings.Contains(relativePath, "\x00") {
		return "", errors.New("external Windows app bundle import plan relative path escapes bundle root")
	}
	return relativePath, nil
}

func externalWinAppBundleManifestSHA256(files []externalWinAppBundleFile) string {
	hash := sha256.New()
	for _, file := range files {
		hash.Write([]byte(file.relativePath))
		hash.Write([]byte{0})
		hash.Write([]byte(fmt.Sprintf("%d", file.sizeBytes)))
		hash.Write([]byte{0})
		hash.Write([]byte(file.sha256))
		hash.Write([]byte{'\n'})
	}
	return hex.EncodeToString(hash.Sum(nil))
}
