package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ExternalWinAppImportRecordSchemaVersion = "xnix.runtime.external_winapp_import_record.v1"
	ExternalWinAppImportRecordRequestType   = "external-winapp-import-record"
)

var externalWinAppDigestPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

type ExternalWinAppImportRequest struct {
	Version        string
	StateRoot      string
	ExecutablePath string
	AppID          string
	DisplayName    string
	AppVersion     string
}

type ExternalWinAppImportRecord struct {
	Version                     string `json:"version"`
	SchemaVersion               string `json:"schema_version"`
	RequestType                 string `json:"request_type"`
	RecordType                  string `json:"record_type"`
	Source                      string `json:"source"`
	RuntimeMethod               string `json:"runtime_method"`
	ReadMethod                  string `json:"read_method"`
	ApplicationID               string `json:"application_id"`
	DisplayName                 string `json:"display_name"`
	AppVersion                  string `json:"app_version"`
	ExecutableName              string `json:"executable_name"`
	ArtifactSHA256              string `json:"artifact_sha256"`
	ArtifactSizeBytes           int64  `json:"artifact_size_bytes"`
	ArtifactRelativePath        string `json:"artifact_relative_path"`
	RecordRelativePath          string `json:"record_relative_path"`
	RecordSHA256                string `json:"record_sha256"`
	WindowsExecutableValidated  bool   `json:"windows_executable_validated"`
	ArtifactCopied              bool   `json:"artifact_copied"`
	ImportRecorded              bool   `json:"import_recorded"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	GoRuntimeBacked             bool   `json:"go_runtime_backed"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	DesktopPageRenderable       bool   `json:"desktop_page_renderable"`
	LaunchEnabled               bool   `json:"launch_enabled"`
	BackendLaunchEnabled        bool   `json:"backend_launch_enabled"`
	ActionExecutionEnabled      bool   `json:"action_execution_enabled"`
	StateRootPathExposed        bool   `json:"state_root_path_exposed"`
	RawExecutablePathExposed    bool   `json:"raw_executable_path_exposed"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	HostRootModified            bool   `json:"host_root_modified"`
	NetworkRequired             bool   `json:"network_required"`
	PrivilegedContainerRequired bool   `json:"privileged_container_required"`
	DockerSocketMounted         bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool   `json:"broad_host_mount_required"`
	PackageManagerInvoked       bool   `json:"package_manager_invoked"`
	DesktopSafeSummary          string `json:"desktop_safe_summary"`
}

func RecordExternalWinAppImport(request ExternalWinAppImportRequest) (ExternalWinAppImportRecord, error) {
	version := strings.TrimSpace(request.Version)
	if version == "" {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires a version")
	}
	if !versionPattern.MatchString(version) {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import version must be semantic version format")
	}
	stateRoot := strings.TrimSpace(request.StateRoot)
	if stateRoot == "" {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires --state-root")
	}
	stateRoot, err := filepath.Abs(stateRoot)
	if err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("resolve external Windows app state root: %w", err)
	}
	executablePath := strings.TrimSpace(request.ExecutablePath)
	if executablePath == "" {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires --executable")
	}
	appID := strings.TrimSpace(request.AppID)
	if !idPattern.MatchString(appID) {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires a reverse-DNS app id")
	}
	displayName := strings.TrimSpace(request.DisplayName)
	if !singleLine(displayName) {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires a single-line display name")
	}
	appVersion := strings.TrimSpace(request.AppVersion)
	if appVersion == "" {
		appVersion = version
	}
	if !versionPattern.MatchString(appVersion) {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import app version must be semantic version format")
	}
	executableName := filepath.Base(executablePath)
	if !safeExternalExecutableName(executableName) {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires a safe .exe file name")
	}

	content, err := os.ReadFile(executablePath)
	if err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("read external Windows executable: %w", err)
	}
	if len(content) < 2 || content[0] != 'M' || content[1] != 'Z' {
		return ExternalWinAppImportRecord{}, errors.New("external Windows app import requires an MZ executable")
	}
	sum := sha256.Sum256(content)
	artifactSHA256 := hex.EncodeToString(sum[:])
	artifactRelativePath := filepath.ToSlash(filepath.Join("external-apps", appID, "artifacts", artifactSHA256, executableName))
	recordRelativePath := filepath.ToSlash(filepath.Join("external-apps", appID, "import-record.json"))
	artifactPath, err := safeStateRootPath(stateRoot, artifactRelativePath)
	if err != nil {
		return ExternalWinAppImportRecord{}, err
	}
	recordPath, err := safeStateRootPath(stateRoot, recordRelativePath)
	if err != nil {
		return ExternalWinAppImportRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(artifactPath), 0o700); err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("prepare external Windows app artifact directory: %w", err)
	}
	if err := os.WriteFile(artifactPath, content, 0o644); err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("write external Windows app artifact: %w", err)
	}

	record := ExternalWinAppImportRecord{
		Version:                     version,
		SchemaVersion:               ExternalWinAppImportRecordSchemaVersion,
		RequestType:                 ExternalWinAppImportRecordRequestType,
		RecordType:                  "external-windows-app-import-record",
		Source:                      "go-runtime-external-winapp-import",
		RuntimeMethod:               "RecordExternalWinAppImport",
		ReadMethod:                  "GetExternalWinAppImportRecord",
		ApplicationID:               appID,
		DisplayName:                 displayName,
		AppVersion:                  appVersion,
		ExecutableName:              executableName,
		ArtifactSHA256:              artifactSHA256,
		ArtifactSizeBytes:           int64(len(content)),
		ArtifactRelativePath:        artifactRelativePath,
		RecordRelativePath:          recordRelativePath,
		WindowsExecutableValidated:  true,
		ArtifactCopied:              true,
		ImportRecorded:              true,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		DesktopPageRenderable:       true,
		LaunchEnabled:               false,
		BackendLaunchEnabled:        false,
		ActionExecutionEnabled:      false,
		StateRootPathExposed:        false,
		RawExecutablePathExposed:    false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		PackageManagerInvoked:       false,
		DesktopSafeSummary:          displayName + " was imported into the Runtime-managed external Windows app state root with a verified executable digest; desktop launch remains gated.",
	}
	data, digest, err := marshalExternalWinAppImportRecord(record)
	if err != nil {
		return ExternalWinAppImportRecord{}, err
	}
	record.RecordSHA256 = digest
	data, _, err = marshalExternalWinAppImportRecord(record)
	if err != nil {
		return ExternalWinAppImportRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(recordPath), 0o700); err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("prepare external Windows app import record directory: %w", err)
	}
	if err := os.WriteFile(recordPath, data, 0o600); err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("write external Windows app import record: %w", err)
	}
	return record, nil
}

func ValidateExternalWinAppImportRecord(record ExternalWinAppImportRecord) []string {
	var reasons []string
	switch {
	case record.SchemaVersion != ExternalWinAppImportRecordSchemaVersion:
		reasons = append(reasons, "external Windows app import record schema is not supported")
	case record.RequestType != ExternalWinAppImportRecordRequestType:
		reasons = append(reasons, "external Windows app import record request type is not supported")
	case record.RecordType != "external-windows-app-import-record":
		reasons = append(reasons, "external Windows app import record type is not supported")
	}
	if !versionPattern.MatchString(strings.TrimSpace(record.Version)) {
		reasons = append(reasons, "external Windows app import record version is invalid")
	}
	if !idPattern.MatchString(strings.TrimSpace(record.ApplicationID)) {
		reasons = append(reasons, "external Windows app import record application id is invalid")
	}
	if !singleLine(record.DisplayName) {
		reasons = append(reasons, "external Windows app import record display name is invalid")
	}
	if !versionPattern.MatchString(strings.TrimSpace(record.AppVersion)) {
		reasons = append(reasons, "external Windows app import record app version is invalid")
	}
	if !safeExternalExecutableName(record.ExecutableName) {
		reasons = append(reasons, "external Windows app import record executable name is invalid")
	}
	if !externalWinAppDigestPattern.MatchString(record.ArtifactSHA256) {
		reasons = append(reasons, "external Windows app import record artifact digest is invalid")
	}
	if record.ArtifactSizeBytes <= 0 {
		reasons = append(reasons, "external Windows app import record artifact size is invalid")
	}
	for _, relativePath := range []string{record.ArtifactRelativePath, record.RecordRelativePath} {
		if relativePath == "" || filepath.IsAbs(relativePath) ||
			strings.Contains(filepath.ToSlash(relativePath), "../") ||
			strings.HasPrefix(filepath.ToSlash(relativePath), "..") {
			reasons = append(reasons, "external Windows app import record relative path is invalid")
			break
		}
	}
	if record.RecordSHA256 == "" || !externalWinAppDigestPattern.MatchString(record.RecordSHA256) {
		reasons = append(reasons, "external Windows app import record digest is invalid")
	} else if digest, err := externalWinAppImportRecordDigest(record); err != nil {
		reasons = append(reasons, "external Windows app import record digest could not be verified")
	} else if digest != record.RecordSHA256 {
		reasons = append(reasons, "external Windows app import record digest mismatch")
	}
	if !record.WindowsExecutableValidated || !record.ArtifactCopied || !record.ImportRecorded || !record.DesktopPageRenderable {
		reasons = append(reasons, "external Windows app import record is incomplete")
	}
	if !record.RuntimeOwned || !record.GoRuntimeBacked || record.KDEPolicyOwner {
		reasons = append(reasons, "external Windows app import record owner flags are invalid")
	}
	if record.LaunchEnabled || record.BackendLaunchEnabled || record.ActionExecutionEnabled ||
		record.StateRootPathExposed || record.RawExecutablePathExposed || record.BackendDetailsExposed ||
		record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired ||
		record.DockerSocketMounted || record.BroadHostMountRequired || record.PackageManagerInvoked {
		reasons = append(reasons, "external Windows app import record contains unsafe enabled gates")
	}
	return reasons
}

func LoadExternalWinAppImportRecord(path string) (ExternalWinAppImportRecord, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("read external Windows app import record: %w", err)
	}
	var record ExternalWinAppImportRecord
	if err := json.Unmarshal(content, &record); err != nil {
		return ExternalWinAppImportRecord{}, fmt.Errorf("parse external Windows app import record: %w", err)
	}
	if reasons := ValidateExternalWinAppImportRecord(record); len(reasons) > 0 {
		return ExternalWinAppImportRecord{}, fmt.Errorf("external Windows app import record is invalid: %s", strings.Join(reasons, "; "))
	}
	return record, nil
}

func ResolveExternalWinAppImportedArtifact(recordPath string) (ExternalWinAppImportRecord, string, error) {
	record, err := LoadExternalWinAppImportRecord(recordPath)
	if err != nil {
		return ExternalWinAppImportRecord{}, "", err
	}
	absoluteRecordPath, err := filepath.Abs(recordPath)
	if err != nil {
		return ExternalWinAppImportRecord{}, "", fmt.Errorf("resolve external Windows app import record path: %w", err)
	}
	recordRelativePath := filepath.Clean(filepath.FromSlash(record.RecordRelativePath))
	if filepath.IsAbs(recordRelativePath) ||
		strings.Contains(filepath.ToSlash(recordRelativePath), "../") ||
		strings.HasPrefix(filepath.ToSlash(recordRelativePath), "..") {
		return ExternalWinAppImportRecord{}, "", errors.New("external Windows app import record has an unsafe record path")
	}
	stateRoot := filepath.Dir(absoluteRecordPath)
	recordRelativeDir := filepath.Dir(recordRelativePath)
	if recordRelativeDir != "." {
		for range strings.Split(filepath.ToSlash(recordRelativeDir), "/") {
			stateRoot = filepath.Dir(stateRoot)
		}
	}
	expectedRecordPath := filepath.Clean(filepath.Join(stateRoot, recordRelativePath))
	if filepath.Clean(absoluteRecordPath) != expectedRecordPath {
		return ExternalWinAppImportRecord{}, "", errors.New("external Windows app import record path does not match its state-root relative path")
	}
	artifactPath, err := safeStateRootPath(stateRoot, record.ArtifactRelativePath)
	if err != nil {
		return ExternalWinAppImportRecord{}, "", err
	}
	content, err := os.ReadFile(artifactPath)
	if err != nil {
		return ExternalWinAppImportRecord{}, "", fmt.Errorf("read imported external Windows app artifact: %w", err)
	}
	sum := sha256.Sum256(content)
	if hex.EncodeToString(sum[:]) != record.ArtifactSHA256 {
		return ExternalWinAppImportRecord{}, "", errors.New("external Windows app imported artifact digest mismatch")
	}
	if int64(len(content)) != record.ArtifactSizeBytes {
		return ExternalWinAppImportRecord{}, "", errors.New("external Windows app imported artifact size mismatch")
	}
	if len(content) < 2 || content[0] != 'M' || content[1] != 'Z' {
		return ExternalWinAppImportRecord{}, "", errors.New("external Windows app imported artifact is not an MZ executable")
	}
	return record, artifactPath, nil
}

func ExternalAppRecipeFromImportRecord(record ExternalWinAppImportRecord) (Recipe, Provenance, error) {
	if reasons := ValidateExternalWinAppImportRecord(record); len(reasons) > 0 {
		return Recipe{}, Provenance{}, fmt.Errorf("external Windows app import record is invalid: %s", strings.Join(reasons, "; "))
	}
	recipe := Recipe{
		ID:                  record.ApplicationID,
		Name:                record.DisplayName,
		Version:             record.AppVersion,
		Icon:                "application-x-executable",
		Mode:                "automatic",
		SupportedExtensions: []string{},
	}
	if err := recipe.Validate(); err != nil {
		return Recipe{}, Provenance{}, err
	}
	return recipe, Provenance{
		Source:          "external-winapp-import-record",
		RegistryName:    "runtime-managed-external-apps",
		DigestVerified:  true,
		SignatureStatus: "local-import-digest-verified",
	}, nil
}

func safeExternalExecutableName(name string) bool {
	name = strings.TrimSpace(name)
	return singleLine(name) &&
		!strings.ContainsAny(name, `/\`) &&
		strings.HasSuffix(strings.ToLower(name), ".exe")
}

func safeStateRootPath(root string, relativePath string) (string, error) {
	if filepath.IsAbs(relativePath) {
		return "", errors.New("state-root relative path must not be absolute")
	}
	cleanRelative := filepath.Clean(filepath.FromSlash(relativePath))
	path := filepath.Join(root, cleanRelative)
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", fmt.Errorf("state-root relative path escapes root: %s", relativePath)
	}
	return path, nil
}

func marshalExternalWinAppImportRecord(record ExternalWinAppImportRecord) ([]byte, string, error) {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}

func externalWinAppImportRecordDigest(record ExternalWinAppImportRecord) (string, error) {
	record.RecordSHA256 = ""
	_, digest, err := marshalExternalWinAppImportRecord(record)
	return digest, err
}
