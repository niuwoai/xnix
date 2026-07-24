package winapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	LauncherBundleSchemaVersion = "xnix.runtime.windows_app_launcher_bundle.v1"
	LauncherBundleRequestType   = "windows-app-launcher-bundle-record"
	defaultRuntimeBinary        = "xnix-runtime-go"
)

type LauncherBundleRequest struct {
	ProfilePath   string
	ApplicationID string
	DisplayName   string
	RuntimeBinary string
}

type LauncherBundleRecord struct {
	SchemaVersion                       string `json:"schema_version"`
	RequestType                         string `json:"request_type"`
	Status                              string `json:"status"`
	ApplicationID                       string `json:"application_id"`
	DisplayName                         string `json:"display_name"`
	DesktopFileName                     string `json:"desktop_file_name"`
	LauncherScriptName                  string `json:"launcher_script_name"`
	ReceiptFileName                     string `json:"receipt_file_name"`
	ProfileSupplied                     bool   `json:"profile_supplied"`
	ProfileSchemaVersion                string `json:"profile_schema_version"`
	StateRootConfigured                 bool   `json:"state_root_configured"`
	FilesWritten                        bool   `json:"files_written"`
	LauncherScriptWritten               bool   `json:"launcher_script_written"`
	DesktopEntryWritten                 bool   `json:"desktop_entry_written"`
	ReceiptWritten                      bool   `json:"receipt_written"`
	DesktopEntryExecUsesManagedLauncher bool   `json:"desktop_entry_exec_uses_managed_launcher"`
	DesktopEntryTerminalDisabled        bool   `json:"desktop_entry_terminal_disabled"`
	RawProfilePathExposed               bool   `json:"raw_profile_path_exposed"`
	RawStateRootPathExposed             bool   `json:"raw_state_root_path_exposed"`
	RawRunnerPathExposed                bool   `json:"raw_runner_path_exposed"`
	HostRootModified                    bool   `json:"host_root_modified"`
	PrivilegedContainerRequired         bool   `json:"privileged_container_required"`
	HostNetworkingRequired              bool   `json:"host_networking_required"`
	DockerSocketMounted                 bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired              bool   `json:"broad_host_mount_required"`
	FailureReason                       string `json:"failure_reason,omitempty"`
}

func RecordLauncherBundle(request LauncherBundleRequest) (LauncherBundleRecord, error) {
	record := baseLauncherBundleRecord(request)
	profilePath := strings.TrimSpace(request.ProfilePath)
	if profilePath == "" {
		record.FailureReason = "profile path is required"
		return record, nil
	}
	profileRequest, err := LoadSmokeProfile(profilePath)
	if err != nil {
		record.FailureReason = "profile is unreadable or invalid"
		return record, nil
	}
	record.ProfileSupplied = true
	record.ProfileSchemaVersion = SmokeProfileSchemaVersion
	if strings.TrimSpace(profileRequest.StateRoot) == "" {
		record.FailureReason = "profile state root is required"
		return record, nil
	}
	record.StateRootConfigured = true

	appID, err := sanitizeApplicationID(request.ApplicationID)
	if err != nil {
		record.FailureReason = err.Error()
		return record, nil
	}
	displayName, err := sanitizeDisplayName(request.DisplayName, appID)
	if err != nil {
		record.FailureReason = err.Error()
		return record, nil
	}
	runtimeBinary, err := sanitizeRuntimeBinary(request.RuntimeBinary)
	if err != nil {
		record.FailureReason = err.Error()
		return record, nil
	}

	stateRoot, err := filepath.Abs(profileRequest.StateRoot)
	if err != nil {
		return record, fmt.Errorf("resolve launcher bundle state root: %w", err)
	}
	profilePath, err = filepath.Abs(profilePath)
	if err != nil {
		return record, fmt.Errorf("resolve launcher bundle profile path: %w", err)
	}
	bundleRoot := filepath.Join(stateRoot, "launcher-bundle")
	launcherDir := filepath.Join(bundleRoot, "launchers")
	desktopDir := filepath.Join(bundleRoot, "applications")
	receiptDir := filepath.Join(bundleRoot, "receipts")
	if err := os.MkdirAll(launcherDir, 0o700); err != nil {
		return record, fmt.Errorf("create launcher directory: %w", err)
	}
	if err := os.MkdirAll(desktopDir, 0o700); err != nil {
		return record, fmt.Errorf("create desktop entry directory: %w", err)
	}
	if err := os.MkdirAll(receiptDir, 0o700); err != nil {
		return record, fmt.Errorf("create launcher receipt directory: %w", err)
	}

	launcherScriptName := appID + ".sh"
	desktopFileName := appID + ".desktop"
	receiptFileName := appID + ".launcher-bundle.json"
	launcherScriptPath := filepath.Join(launcherDir, launcherScriptName)
	desktopFilePath := filepath.Join(desktopDir, desktopFileName)
	receiptFilePath := filepath.Join(receiptDir, receiptFileName)

	launcherScript := renderLauncherScript(runtimeBinary, profilePath)
	if err := os.WriteFile(launcherScriptPath, []byte(launcherScript), 0o700); err != nil {
		return record, fmt.Errorf("write launcher script: %w", err)
	}
	record.LauncherScriptWritten = true
	desktopEntry := renderDesktopEntry(appID, displayName, launcherScriptPath)
	if err := os.WriteFile(desktopFilePath, []byte(desktopEntry), 0o600); err != nil {
		return record, fmt.Errorf("write desktop entry: %w", err)
	}
	record.DesktopEntryWritten = true

	record.ApplicationID = appID
	record.DisplayName = displayName
	record.DesktopFileName = desktopFileName
	record.LauncherScriptName = launcherScriptName
	record.ReceiptFileName = receiptFileName
	record.FilesWritten = true
	record.Status = PassedStatus
	record.DesktopEntryExecUsesManagedLauncher = true
	record.DesktopEntryTerminalDisabled = true

	receiptData, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return record, fmt.Errorf("encode launcher receipt: %w", err)
	}
	if err := os.WriteFile(receiptFilePath, append(receiptData, '\n'), 0o600); err != nil {
		return record, fmt.Errorf("write launcher receipt: %w", err)
	}
	record.ReceiptWritten = true
	receiptData, err = json.MarshalIndent(record, "", "  ")
	if err != nil {
		return record, fmt.Errorf("encode final launcher receipt: %w", err)
	}
	if err := os.WriteFile(receiptFilePath, append(receiptData, '\n'), 0o600); err != nil {
		return record, fmt.Errorf("write final launcher receipt: %w", err)
	}
	return record, nil
}

func baseLauncherBundleRecord(request LauncherBundleRequest) LauncherBundleRecord {
	return LauncherBundleRecord{
		SchemaVersion:                       LauncherBundleSchemaVersion,
		RequestType:                         LauncherBundleRequestType,
		Status:                              FailedStatus,
		ApplicationID:                       strings.TrimSpace(request.ApplicationID),
		DisplayName:                         strings.TrimSpace(request.DisplayName),
		ProfileSupplied:                     strings.TrimSpace(request.ProfilePath) != "",
		DesktopEntryExecUsesManagedLauncher: false,
		DesktopEntryTerminalDisabled:        false,
		RawProfilePathExposed:               false,
		RawStateRootPathExposed:             false,
		RawRunnerPathExposed:                false,
		HostRootModified:                    false,
		PrivilegedContainerRequired:         false,
		HostNetworkingRequired:              false,
		DockerSocketMounted:                 false,
		BroadHostMountRequired:              false,
	}
}

var applicationIDPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]{1,126}[A-Za-z0-9]$`)

func sanitizeApplicationID(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("application id is required")
	}
	if strings.Contains(value, "..") || !applicationIDPattern.MatchString(value) {
		return "", errors.New("application id must be stable and desktop-file safe")
	}
	return value, nil
}

func sanitizeDisplayName(value string, fallback string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		value = fallback
	}
	if strings.ContainsAny(value, "\r\n") {
		return "", errors.New("display name must be single-line")
	}
	return value, nil
}

func sanitizeRuntimeBinary(value string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultRuntimeBinary, nil
	}
	if strings.ContainsAny(value, "\r\n") {
		return "", errors.New("runtime binary must be single-line")
	}
	return value, nil
}

func renderLauncherScript(runtimeBinary string, profilePath string) string {
	return "#!/bin/sh\n" +
		"set -eu\n" +
		"exec " + shellQuote(runtimeBinary) + " windows-app-run-smoke --profile " + shellQuote(profilePath) + "\n"
}

func renderDesktopEntry(applicationID string, displayName string, launcherScriptPath string) string {
	return "[Desktop Entry]\n" +
		"Type=Application\n" +
		"Name=" + desktopEntryValue(displayName) + "\n" +
		"Exec=" + desktopEntryExec(launcherScriptPath) + "\n" +
		"Terminal=false\n" +
		"Categories=Utility;X-Xnix-WindowsApp;\n" +
		"StartupNotify=true\n" +
		"X-Xnix-ApplicationID=" + desktopEntryValue(applicationID) + "\n" +
		"X-Xnix-RuntimeOwned=true\n"
}

func desktopEntryValue(value string) string {
	value = strings.ReplaceAll(value, "\r", " ")
	value = strings.ReplaceAll(value, "\n", " ")
	return strings.ReplaceAll(value, "\\", "\\\\")
}

func desktopEntryExec(value string) string {
	escaped := strings.ReplaceAll(value, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
	return "\"" + escaped + "\""
}
