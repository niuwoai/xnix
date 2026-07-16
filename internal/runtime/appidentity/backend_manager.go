package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BackendManagerPreview struct {
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	ManagerType                 string                        `json:"manager_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	ReadMethod                  string                        `json:"read_method"`
	Backends                    []ManagedCompatibilityBackend `json:"backends"`
	BackendIDs                  []string                      `json:"backend_ids"`
	UserFacingProfiles          []UserFacingBackendProfile    `json:"user_facing_profiles"`
	UserFacingProfileIDs        []string                      `json:"user_facing_profile_ids"`
	BackendCount                int                           `json:"backend_count"`
	UserFacingProfileCount      int                           `json:"user_facing_profile_count"`
	RuntimeOwned                bool                          `json:"runtime_owned"`
	GoRuntimeBacked             bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                          `json:"kde_policy_owner"`
	KDEVisible                  bool                          `json:"kde_visible"`
	WineManaged                 bool                          `json:"wine_managed"`
	ProtonManaged               bool                          `json:"proton_managed"`
	WindowsVMManaged            bool                          `json:"windows_vm_managed"`
	BackendInstallEnabled       bool                          `json:"backend_install_enabled"`
	BackendDownloadEnabled      bool                          `json:"backend_download_enabled"`
	BackendLaunchEnabled        bool                          `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                          `json:"backend_process_started"`
	VMProcessStarted            bool                          `json:"vm_process_started"`
	RawCommandExposed           bool                          `json:"raw_command_exposed"`
	ProfilePathExposed          bool                          `json:"profile_path_exposed"`
	BackendDetailsExposedToKDE  bool                          `json:"backend_details_exposed_to_kde"`
	HostRootModified            bool                          `json:"host_root_modified"`
	NetworkRequired             bool                          `json:"network_required"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	SecretsExposed              bool                          `json:"secrets_exposed"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type BackendManagerRecord struct {
	SchemaVersion               string                `json:"schema_version"`
	RecordType                  string                `json:"record_type"`
	Source                      string                `json:"source"`
	RelativePath                string                `json:"relative_path"`
	Preview                     BackendManagerPreview `json:"preview"`
	SHA256                      string                `json:"sha256"`
	RuntimeOwned                bool                  `json:"runtime_owned"`
	GoRuntimeBacked             bool                  `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                  `json:"kde_policy_owner"`
	StateRootPathExposed        bool                  `json:"state_root_path_exposed"`
	BackendInstallEnabled       bool                  `json:"backend_install_enabled"`
	BackendDownloadEnabled      bool                  `json:"backend_download_enabled"`
	BackendLaunchEnabled        bool                  `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                  `json:"backend_process_started"`
	VMProcessStarted            bool                  `json:"vm_process_started"`
	RawCommandExposed           bool                  `json:"raw_command_exposed"`
	ProfilePathExposed          bool                  `json:"profile_path_exposed"`
	BackendDetailsExposedToKDE  bool                  `json:"backend_details_exposed_to_kde"`
	HostRootModified            bool                  `json:"host_root_modified"`
	NetworkRequired             bool                  `json:"network_required"`
	PrivilegedContainerRequired bool                  `json:"privileged_container_required"`
	SecretsExposed              bool                  `json:"secrets_exposed"`
	Summary                     string                `json:"summary"`
}

type ManagedCompatibilityBackend struct {
	ID                         string   `json:"id"`
	Kind                       string   `json:"kind"`
	RuntimeRole                string   `json:"runtime_role"`
	Status                     string   `json:"status"`
	Readiness                  string   `json:"readiness"`
	UserFacingProfileID        string   `json:"user_facing_profile_id"`
	RequiredGates              []string `json:"required_gates"`
	InstallEnabled             bool     `json:"install_enabled"`
	DownloadEnabled            bool     `json:"download_enabled"`
	LaunchEnabled              bool     `json:"launch_enabled"`
	ProcessStarted             bool     `json:"process_started"`
	RawCommandExposed          bool     `json:"raw_command_exposed"`
	ProfilePathExposed         bool     `json:"profile_path_exposed"`
	BackendDetailsExposedToKDE bool     `json:"backend_details_exposed_to_kde"`
	HostRootModified           bool     `json:"host_root_modified"`
	NetworkRequired            bool     `json:"network_required"`
}

type UserFacingBackendProfile struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Summary               string `json:"summary"`
	DefaultForAutomatic   bool   `json:"default_for_automatic"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	LaunchEnabled         bool   `json:"launch_enabled"`
}

func NewBackendManagerPreview() BackendManagerPreview {
	backends := []ManagedCompatibilityBackend{
		managedCompatibilityBackend("wine", "local-windows-api", "local-compatibility", []string{"recipe-trust", "artifact-staging", "state-root", "portal-policy-review", "snapshot-baseline"}),
		managedCompatibilityBackend("proton", "gaming-compatibility", "local-compatibility", []string{"recipe-trust", "artifact-staging", "state-root", "portal-policy-review", "snapshot-baseline"}),
		managedCompatibilityBackend("windows-vm", "isolated-windows-runtime", "isolated-compatibility", []string{"recipe-trust", "artifact-staging", "state-root", "portal-policy-review", "snapshot-baseline", "vm-image-review"}),
	}
	profiles := []UserFacingBackendProfile{
		userFacingBackendProfile("automatic", "Automatic", "Runtime chooses the safest compatible mode after review gates pass.", true),
		userFacingBackendProfile("local-compatibility", "Local compatibility", "Runs close to the desktop session after Runtime approval.", false),
		userFacingBackendProfile("isolated-compatibility", "Isolated compatibility", "Uses an isolated Runtime environment after Runtime approval.", false),
	}

	return BackendManagerPreview{
		SchemaVersion:               "xnix.runtime.backend_manager.v1",
		RequestType:                 "backend-manager-preview",
		ManagerType:                 "compatibility-backend-manager",
		Source:                      "go-runtime-backend-manager",
		Desktop:                     "KDE Plasma",
		ReadMethod:                  "GetBackendManagerPreview",
		Backends:                    backends,
		BackendIDs:                  managedBackendIDs(backends),
		UserFacingProfiles:          profiles,
		UserFacingProfileIDs:        userFacingBackendProfileIDs(profiles),
		BackendCount:                len(backends),
		UserFacingProfileCount:      len(profiles),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEVisible:                  false,
		WineManaged:                 true,
		ProtonManaged:               true,
		WindowsVMManaged:            true,
		BackendInstallEnabled:       false,
		BackendDownloadEnabled:      false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		VMProcessStarted:            false,
		RawCommandExposed:           false,
		ProfilePathExposed:          false,
		BackendDetailsExposedToKDE:  false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		SecretsExposed:              false,
		BlockedActions: []string{
			"install compatibility backend from manager preview",
			"download compatibility backend from manager preview",
			"launch compatibility backend from manager preview",
			"start VM process from manager preview",
			"expose backend commands or storage paths to KDE",
			"mutate host root during backend manager preview",
		},
		DesktopSafeSummary: "Runtime owns the compatibility backend manager; KDE receives only user-facing compatibility modes while backend installation, launch, downloads, paths, and commands remain disabled.",
	}
}

func RecordBackendManagerPreview(stateRoot string) (BackendManagerRecord, error) {
	root, err := safeBackendManagerRoot(stateRoot)
	if err != nil {
		return BackendManagerRecord{}, err
	}
	relativePath, path, err := backendManagerRecordPath(root, true)
	if err != nil {
		return BackendManagerRecord{}, err
	}

	record := BackendManagerRecord{
		SchemaVersion:               "xnix.runtime.backend_manager_record.v1",
		RecordType:                  "backend-manager-inventory-record",
		Source:                      "go-runtime-state-root-backend-manager",
		RelativePath:                relativePath,
		Preview:                     NewBackendManagerPreview(),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		BackendInstallEnabled:       false,
		BackendDownloadEnabled:      false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		VMProcessStarted:            false,
		RawCommandExposed:           false,
		ProfilePathExposed:          false,
		BackendDetailsExposedToKDE:  false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		SecretsExposed:              false,
		Summary:                     "Runtime persisted compatibility backend inventory under the configured state root without installing, downloading, launching, exposing paths, or mutating the host root.",
	}
	data, digest, err := marshalBackendManagerRecord(record)
	if err != nil {
		return BackendManagerRecord{}, err
	}
	record.SHA256 = digest
	data, _, err = marshalBackendManagerRecord(record)
	if err != nil {
		return BackendManagerRecord{}, err
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return BackendManagerRecord{}, fmt.Errorf("write backend manager record: %w", err)
	}
	return record, nil
}

// LoadBackendManagerRecord reads and validates the persisted Runtime backend
// inventory without creating directories or exposing the configured state root.
func LoadBackendManagerRecord(stateRoot string) (BackendManagerRecord, error) {
	root, err := openBackendManagerRoot(stateRoot)
	if err != nil {
		return BackendManagerRecord{}, err
	}
	relativePath, path, err := backendManagerRecordPath(root, false)
	if err != nil {
		return BackendManagerRecord{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return BackendManagerRecord{}, fmt.Errorf("read backend manager record: %w", err)
	}
	var record BackendManagerRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return BackendManagerRecord{}, fmt.Errorf("parse backend manager record: %w", err)
	}
	if err := validateBackendManagerRecord(record, relativePath); err != nil {
		return BackendManagerRecord{}, err
	}
	return record, nil
}

func managedCompatibilityBackend(id string, kind string, userFacingProfileID string, gates []string) ManagedCompatibilityBackend {
	return ManagedCompatibilityBackend{
		ID:                         id,
		Kind:                       kind,
		RuntimeRole:                "managed-compatibility-backend",
		Status:                     "planned",
		Readiness:                  "blocked-until-runtime-gates-pass",
		UserFacingProfileID:        userFacingProfileID,
		RequiredGates:              gates,
		InstallEnabled:             false,
		DownloadEnabled:            false,
		LaunchEnabled:              false,
		ProcessStarted:             false,
		RawCommandExposed:          false,
		ProfilePathExposed:         false,
		BackendDetailsExposedToKDE: false,
		HostRootModified:           false,
		NetworkRequired:            false,
	}
}

func userFacingBackendProfile(id string, label string, summary string, defaultForAutomatic bool) UserFacingBackendProfile {
	return UserFacingBackendProfile{
		ID:                    id,
		Label:                 label,
		Summary:               summary,
		DefaultForAutomatic:   defaultForAutomatic,
		RuntimeOwned:          true,
		KDEPolicyOwner:        false,
		BackendDetailsExposed: false,
		LaunchEnabled:         false,
	}
}

func managedBackendIDs(backends []ManagedCompatibilityBackend) []string {
	ids := make([]string, 0, len(backends))
	for _, backend := range backends {
		ids = append(ids, backend.ID)
	}
	return ids
}

func userFacingBackendProfileIDs(profiles []UserFacingBackendProfile) []string {
	ids := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.ID)
	}
	return ids
}

func safeBackendManagerRoot(root string) (string, error) {
	clean, err := normalizeBackendManagerRoot(root)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(clean, 0o700); err != nil {
		return "", fmt.Errorf("prepare backend manager state root: %w", err)
	}
	return clean, nil
}

func openBackendManagerRoot(root string) (string, error) {
	clean, err := normalizeBackendManagerRoot(root)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(clean)
	if err != nil {
		return "", fmt.Errorf("backend manager state root must exist: %w", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("backend manager state root is not a directory: %s", clean)
	}
	return clean, nil
}

func normalizeBackendManagerRoot(root string) (string, error) {
	if root == "" {
		return "", errors.New("backend manager record requires an explicit state root")
	}
	clean, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve backend manager state root: %w", err)
	}
	clean = filepath.Clean(clean)
	if clean == string(os.PathSeparator) {
		return "", errors.New("refusing to use filesystem root as backend manager state root")
	}
	return clean, nil
}

func backendManagerRecordPath(root string, createDirectory bool) (string, string, error) {
	relativePath := filepath.ToSlash(filepath.Join("backend-manager", "inventory.json"))
	directory := filepath.Join(root, "backend-manager")
	info, err := os.Lstat(directory)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", "", fmt.Errorf("inspect backend manager record directory: %w", err)
		}
		if createDirectory {
			if err := os.Mkdir(directory, 0o700); err != nil {
				return "", "", fmt.Errorf("prepare backend manager record directory: %w", err)
			}
		} else {
			return "", "", fmt.Errorf("backend manager record directory does not exist: %s", filepath.Base(directory))
		}
	} else if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return "", "", errors.New("backend manager record directory must be a real directory")
	}
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return "", "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("backend manager record path escapes state root: %s", relativePath)
	}
	return relativePath, path, nil
}

func validateBackendManagerRecord(record BackendManagerRecord, relativePath string) error {
	if record.SchemaVersion != "xnix.runtime.backend_manager_record.v1" || record.RecordType != "backend-manager-inventory-record" || record.Source != "go-runtime-state-root-backend-manager" {
		return errors.New("backend manager record has unsupported schema")
	}
	if record.RelativePath != relativePath || record.Preview.SchemaVersion != "xnix.runtime.backend_manager.v1" || record.Preview.BackendCount != len(record.Preview.Backends) || record.Preview.UserFacingProfileCount != len(record.Preview.UserFacingProfiles) {
		return errors.New("backend manager record identity, path, or count mismatch")
	}
	storedDigest := record.SHA256
	record.SHA256 = ""
	_, expectedDigest, err := marshalBackendManagerRecord(record)
	if err != nil {
		return err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return errors.New("backend manager record digest mismatch")
	}
	if backendManagerRecordUnsafe(record) {
		return errors.New("backend manager record has unsafe enabled gates")
	}
	return nil
}

func backendManagerRecordUnsafe(record BackendManagerRecord) bool {
	if record.StateRootPathExposed || record.BackendInstallEnabled || record.BackendDownloadEnabled || record.BackendLaunchEnabled || record.BackendProcessStarted || record.VMProcessStarted || record.RawCommandExposed || record.ProfilePathExposed || record.BackendDetailsExposedToKDE || record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.SecretsExposed {
		return true
	}
	preview := record.Preview
	if preview.KDEVisible || preview.BackendInstallEnabled || preview.BackendDownloadEnabled || preview.BackendLaunchEnabled || preview.BackendProcessStarted || preview.VMProcessStarted || preview.RawCommandExposed || preview.ProfilePathExposed || preview.BackendDetailsExposedToKDE || preview.HostRootModified || preview.NetworkRequired || preview.PrivilegedContainerRequired || preview.SecretsExposed {
		return true
	}
	for _, backend := range preview.Backends {
		if backend.InstallEnabled || backend.DownloadEnabled || backend.LaunchEnabled || backend.ProcessStarted || backend.RawCommandExposed || backend.ProfilePathExposed || backend.BackendDetailsExposedToKDE || backend.HostRootModified || backend.NetworkRequired {
			return true
		}
	}
	for _, profile := range preview.UserFacingProfiles {
		if profile.KDEPolicyOwner || profile.BackendDetailsExposed || profile.LaunchEnabled {
			return true
		}
	}
	return false
}

func marshalBackendManagerRecord(record BackendManagerRecord) ([]byte, string, error) {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}
