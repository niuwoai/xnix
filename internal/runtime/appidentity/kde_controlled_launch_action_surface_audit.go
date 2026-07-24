package appidentity

import (
	"errors"
	"path/filepath"
	"strings"
)

const (
	KDEControlledLaunchActionSurfaceAuditSchemaVersion = "xnix.runtime.kde_controlled_launch_action_surface_audit.v1"
	KDEControlledLaunchActionSurfaceAuditRequestType   = "kde-controlled-launch-action-surface-audit-preview"
)

type KDEControlledLaunchActionSurfaceAuditRequest struct {
	DesktopEntryContent string
	EvidenceHandle      string
}

type KDEControlledLaunchActionSurfaceAuditPreview struct {
	SchemaVersion                  string   `json:"schema_version"`
	RequestType                    string   `json:"request_type"`
	Source                         string   `json:"source"`
	RuntimeMethod                  string   `json:"runtime_method"`
	ReadMethod                     string   `json:"read_method"`
	AuditState                     string   `json:"audit_state"`
	KDEActionID                    string   `json:"kde_action_id"`
	KDEActionLabel                 string   `json:"kde_action_label"`
	PublicDBusMethod               string   `json:"public_dbus_method"`
	ForwardedArgumentKind          string   `json:"forwarded_argument_kind"`
	EvidenceHandleShapeAccepted    bool     `json:"evidence_handle_shape_accepted"`
	EvidenceOnlyArgumentShape      bool     `json:"evidence_only_argument_shape"`
	RequiredMetadataPresent        bool     `json:"required_metadata_present"`
	MetadataMalformed              bool     `json:"metadata_malformed"`
	UnsafeOwnerArgumentsPresent    bool     `json:"unsafe_owner_arguments_present"`
	UnsafeBackendTermsPresent      bool     `json:"unsafe_backend_terms_present"`
	KDEForwardsOnlyEvidenceHandle  bool     `json:"kde_forwards_only_evidence_handle"`
	KDEPolicyOwner                 bool     `json:"kde_policy_owner"`
	OwnerServiceArgsExposedToKDE   bool     `json:"owner_service_args_exposed_to_kde"`
	OwnerServiceArgumentsRecreated bool     `json:"owner_service_arguments_recreated"`
	StateRootAccess                bool     `json:"state_root_access"`
	ReceiptReconstruction          bool     `json:"receipt_reconstruction"`
	RawExecutablePathExposed       bool     `json:"raw_executable_path_exposed"`
	BackendCommandExposed          bool     `json:"backend_command_exposed"`
	BackendDetailsExposed          bool     `json:"backend_details_exposed"`
	DesktopLaunchEnabled           bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled           bool     `json:"backend_launch_enabled"`
	ExecutionStarted               bool     `json:"execution_started"`
	DBusCalled                     bool     `json:"dbus_called"`
	KDEConfigurationWritten        bool     `json:"kde_configuration_written"`
	RuntimeStateWritten            bool     `json:"runtime_state_written"`
	RuntimeOwned                   bool     `json:"runtime_owned"`
	GoRuntimeBacked                bool     `json:"go_runtime_backed"`
	SurfaceSafeForHumanSmoke       bool     `json:"surface_safe_for_human_smoke"`
	HostRootModified               bool     `json:"host_root_modified"`
	DockerSocketMounted            bool     `json:"docker_socket_mounted"`
	PrivilegedContainerRequired    bool     `json:"privileged_container_required"`
	HostNetworkRequired            bool     `json:"host_network_required"`
	NetworkRequired                bool     `json:"network_required"`
	UnsafeMetadataKeys             []string `json:"unsafe_metadata_keys,omitempty"`
	MissingMetadataKeys            []string `json:"missing_metadata_keys,omitempty"`
	BlockedActions                 []string `json:"blocked_actions"`
	BlockedReason                  string   `json:"blocked_reason,omitempty"`
	DesktopSafeSummary             string   `json:"desktop_safe_summary"`
}

func PreviewKDEControlledLaunchActionSurfaceAudit(request KDEControlledLaunchActionSurfaceAuditRequest) (KDEControlledLaunchActionSurfaceAuditPreview, error) {
	metadata, malformedLines := parseKDEControlledLaunchDesktopMetadata(request.DesktopEntryContent)
	preview := baseKDEControlledLaunchActionSurfaceAudit(metadata, request.EvidenceHandle)
	preview.MissingMetadataKeys = kdeControlledLaunchActionSurfaceMissingKeys(metadata)
	preview.UnsafeMetadataKeys = kdeControlledLaunchActionSurfaceUnsafeKeys(metadata, request.DesktopEntryContent)
	preview.RequiredMetadataPresent = len(preview.MissingMetadataKeys) == 0
	preview.MetadataMalformed = strings.TrimSpace(request.DesktopEntryContent) == "" || len(malformedLines) > 0
	preview.UnsafeOwnerArgumentsPresent = kdeControlledLaunchActionSurfaceOwnerArgsUnsafe(preview, request.DesktopEntryContent)
	preview.UnsafeBackendTermsPresent = kdeControlledLaunchActionSurfaceBackendTermsUnsafe(request.DesktopEntryContent)
	preview.EvidenceOnlyArgumentShape = preview.ForwardedArgumentKind == "evidence-relative-path" && preview.KDEForwardsOnlyEvidenceHandle
	preview.EvidenceHandleShapeAccepted = kdeControlledLaunchEvidenceHandleShapeAccepted(request.EvidenceHandle)

	switch {
	case preview.MetadataMalformed || !preview.RequiredMetadataPresent:
		preview.AuditState = "malformed"
		preview.BlockedReason = "KDE controlled-launch action metadata is incomplete or malformed"
	case preview.UnsafeOwnerArgumentsPresent:
		preview.AuditState = "unsafe-owner-args"
		preview.BlockedReason = "KDE controlled-launch action metadata includes owner-only launch inputs"
	case preview.UnsafeBackendTermsPresent:
		preview.AuditState = "unsafe-backend-terms"
		preview.BlockedReason = "KDE controlled-launch action metadata includes backend or raw executable terms"
	case !preview.EvidenceHandleShapeAccepted:
		preview.AuditState = "missing-evidence-handle"
		preview.BlockedReason = "KDE controlled-launch action audit requires a safe Runtime evidence handle"
	default:
		preview.AuditState = "safe"
	}
	preview.SurfaceSafeForHumanSmoke = preview.AuditState == "safe"
	preview.DesktopSafeSummary = kdeControlledLaunchActionSurfaceAuditSummary(preview)
	return validateKDEControlledLaunchActionSurfaceAuditPreview(preview)
}

func baseKDEControlledLaunchActionSurfaceAudit(metadata map[string]string, evidenceHandle string) KDEControlledLaunchActionSurfaceAuditPreview {
	return KDEControlledLaunchActionSurfaceAuditPreview{
		SchemaVersion:                 KDEControlledLaunchActionSurfaceAuditSchemaVersion,
		RequestType:                   KDEControlledLaunchActionSurfaceAuditRequestType,
		Source:                        "kde-controlled-launch-action.desktop+runtime-surface-audit",
		RuntimeMethod:                 "PreviewKDEControlledLaunchActionSurfaceAudit",
		ReadMethod:                    "GetKDEControlledLaunchActionSurfaceAuditPreview",
		AuditState:                    "missing-evidence-handle",
		KDEActionID:                   strings.TrimSpace(metadata["X-Xnix-KDE-Action-ID"]),
		KDEActionLabel:                strings.TrimSpace(metadata["Name"]),
		PublicDBusMethod:              strings.TrimSpace(metadata["X-Xnix-DBus-Method"]),
		ForwardedArgumentKind:         strings.TrimSpace(metadata["X-Xnix-Forwarded-Argument"]),
		KDEForwardsOnlyEvidenceHandle: truthyDesktopMetadata(metadata["X-Xnix-Forwards-Only-Evidence-Handle"]),
		KDEPolicyOwner:                truthyDesktopMetadata(metadata["X-Xnix-KDE-Policy-Owner"]),
		OwnerServiceArgsExposedToKDE:  truthyDesktopMetadata(metadata["X-Xnix-Owner-Service-Args-Exposed-To-KDE"]),
		StateRootAccess:               truthyDesktopMetadata(metadata["X-Xnix-State-Root-Access"]),
		ReceiptReconstruction:         truthyDesktopMetadata(metadata["X-Xnix-Receipt-Reconstruction"]),
		BackendLaunchEnabled:          truthyDesktopMetadata(metadata["X-Xnix-Backend-Launch-Enabled"]),
		ExecutionStarted:              truthyDesktopMetadata(metadata["X-Xnix-Execution-Started"]),
		HostRootModified:              truthyDesktopMetadata(metadata["X-Xnix-Host-Root-Modified"]),
		DockerSocketMounted:           truthyDesktopMetadata(metadata["X-Xnix-Docker-Socket-Mounted"]),
		PrivilegedContainerRequired:   truthyDesktopMetadata(metadata["X-Xnix-Privileged-Container-Required"]),
		HostNetworkRequired:           truthyDesktopMetadata(metadata["X-Xnix-Host-Network-Required"]),
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		SurfaceSafeForHumanSmoke:      false,
		DesktopLaunchEnabled:          false,
		DBusCalled:                    false,
		KDEConfigurationWritten:       false,
		RuntimeStateWritten:           false,
		NetworkRequired:               false,
		BlockedActions:                kdeControlledLaunchActionSurfaceBlockedActions(),
		DesktopSafeSummary:            "KDE controlled-launch action surface audit is blocked before D-Bus dispatch.",
	}
}

func parseKDEControlledLaunchDesktopMetadata(content string) (map[string]string, []string) {
	metadata := map[string]string{}
	var malformed []string
	for _, rawLine := range strings.Split(content, "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "[") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok || strings.TrimSpace(key) == "" {
			malformed = append(malformed, "unparseable-line")
			continue
		}
		metadata[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	return metadata, malformed
}

func truthyDesktopMetadata(value string) bool {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "true", "1", "yes", "on":
		return true
	default:
		return false
	}
}

func kdeControlledLaunchActionSurfaceMissingKeys(metadata map[string]string) []string {
	required := []string{
		"Name",
		"X-Xnix-KDE-Action-ID",
		"X-Xnix-Runtime-Preview",
		"X-Xnix-Restricted-Smoke-Plan",
		"X-Xnix-DBus-Service",
		"X-Xnix-DBus-Object-Path",
		"X-Xnix-DBus-Method",
		"X-Xnix-Forwarded-Argument",
		"X-Xnix-Forwards-Only-Evidence-Handle",
		"X-Xnix-KDE-Policy-Owner",
		"X-Xnix-Owner-Service-Args-Exposed-To-KDE",
		"X-Xnix-State-Root-Access",
		"X-Xnix-Receipt-Reconstruction",
		"X-Xnix-Backend-Launch-Enabled",
		"X-Xnix-Execution-Started",
		"X-Xnix-Host-Root-Modified",
		"X-Xnix-Docker-Socket-Mounted",
		"X-Xnix-Privileged-Container-Required",
		"X-Xnix-Host-Network-Required",
	}
	var missing []string
	for _, key := range required {
		if strings.TrimSpace(metadata[key]) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}

func kdeControlledLaunchActionSurfaceUnsafeKeys(metadata map[string]string, content string) []string {
	var unsafe []string
	expected := map[string]string{
		"X-Xnix-KDE-Action-ID":                     "xnix.runtime-status.controlled-launch",
		"X-Xnix-Runtime-Preview":                   "xnix-runtime-go kde-controlled-launch-action-preview",
		"X-Xnix-Restricted-Smoke-Plan":             "xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview",
		"X-Xnix-DBus-Service":                      "org.xnix.Compatibility1",
		"X-Xnix-DBus-Object-Path":                  "/org/xnix/Compatibility1",
		"X-Xnix-DBus-Method":                       "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		"X-Xnix-Forwarded-Argument":                "evidence-relative-path",
		"X-Xnix-Forwards-Only-Evidence-Handle":     "true",
		"X-Xnix-KDE-Policy-Owner":                  "false",
		"X-Xnix-Owner-Service-Args-Exposed-To-KDE": "false",
		"X-Xnix-State-Root-Access":                 "false",
		"X-Xnix-Receipt-Reconstruction":            "false",
		"X-Xnix-Backend-Launch-Enabled":            "false",
		"X-Xnix-Execution-Started":                 "false",
		"X-Xnix-Host-Root-Modified":                "false",
		"X-Xnix-Docker-Socket-Mounted":             "false",
		"X-Xnix-Privileged-Container-Required":     "false",
		"X-Xnix-Host-Network-Required":             "false",
	}
	for key, expectedValue := range expected {
		if value := strings.TrimSpace(metadata[key]); value != "" && value != expectedValue {
			unsafe = append(unsafe, key)
		}
	}
	lowerContent := strings.ToLower(content)
	for _, marker := range []string{"--state-root", "--cache-root", "--launcher", "--timeout", "--receipt-id", "--session-id", "--dispatch-id", "owner_service_call_args", "xnix_runtime_owner_"} {
		if strings.Contains(lowerContent, marker) {
			unsafe = append(unsafe, "owner-only-inline-argument")
			break
		}
	}
	return uniqueStrings(unsafe)
}

func kdeControlledLaunchActionSurfaceOwnerArgsUnsafe(preview KDEControlledLaunchActionSurfaceAuditPreview, content string) bool {
	if preview.KDEPolicyOwner || preview.OwnerServiceArgsExposedToKDE || preview.StateRootAccess || preview.ReceiptReconstruction {
		return true
	}
	lowerContent := strings.ToLower(content)
	for _, marker := range []string{"--state-root", "--cache-root", "--launcher", "--timeout", "--receipt-id", "--session-id", "--dispatch-id", "owner_service_call_args", "xnix_runtime_owner_"} {
		if strings.Contains(lowerContent, marker) {
			return true
		}
	}
	return false
}

func kdeControlledLaunchActionSurfaceBackendTermsUnsafe(content string) bool {
	lowerContent := strings.ToLower(content)
	for _, marker := range []string{".exe", "wine ", "wine/", "proton", "qemu-system", "program files", ".wine", "raw-executable", "backend-command"} {
		if strings.Contains(lowerContent, marker) {
			return true
		}
	}
	return false
}

func kdeControlledLaunchEvidenceHandleShapeAccepted(handle string) bool {
	trimmed := strings.TrimSpace(handle)
	if trimmed == "" {
		return false
	}
	clean := filepath.Clean(filepath.ToSlash(trimmed))
	return !filepath.IsAbs(trimmed) &&
		!strings.Contains(clean, "..") &&
		strings.HasPrefix(filepath.ToSlash(trimmed), "runtime/kde-runtime-status-launch-evidence/") &&
		strings.HasSuffix(trimmed, ".json")
}

func kdeControlledLaunchActionSurfaceAuditSummary(preview KDEControlledLaunchActionSurfaceAuditPreview) string {
	switch preview.AuditState {
	case "safe":
		return "KDE controlled-launch action surface is safe for review-only desktop-triggered smoke."
	case "unsafe-owner-args":
		return "KDE controlled-launch action surface is blocked because owner-only launch inputs are present."
	case "unsafe-backend-terms":
		return "KDE controlled-launch action surface is blocked because backend or raw executable terms are present."
	case "malformed":
		return "KDE controlled-launch action surface is blocked because required metadata is incomplete."
	default:
		return "KDE controlled-launch action surface is blocked until a safe Runtime evidence handle is available."
	}
}

func kdeControlledLaunchActionSurfaceBlockedActions() []string {
	return []string{
		"call D-Bus from surface audit",
		"write KDE configuration from surface audit",
		"write Runtime state from surface audit",
		"derive owner service arguments in KDE",
		"expose raw executable or backend command details",
		"start desktop or compatibility execution from surface audit",
		"mutate host root from surface audit",
	}
}

func validateKDEControlledLaunchActionSurfaceAuditPreview(preview KDEControlledLaunchActionSurfaceAuditPreview) (KDEControlledLaunchActionSurfaceAuditPreview, error) {
	validStates := map[string]bool{
		"safe":                    true,
		"unsafe-owner-args":       true,
		"unsafe-backend-terms":    true,
		"malformed":               true,
		"missing-evidence-handle": true,
	}
	switch {
	case preview.SchemaVersion != KDEControlledLaunchActionSurfaceAuditSchemaVersion || preview.RequestType != KDEControlledLaunchActionSurfaceAuditRequestType:
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit has invalid schema")
	case !validStates[preview.AuditState]:
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit has invalid state")
	case preview.SurfaceSafeForHumanSmoke != (preview.AuditState == "safe"):
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit safety flag is inconsistent")
	case preview.SurfaceSafeForHumanSmoke && (!preview.RequiredMetadataPresent || preview.MetadataMalformed || preview.UnsafeOwnerArgumentsPresent || preview.UnsafeBackendTermsPresent || !preview.EvidenceOnlyArgumentShape || !preview.EvidenceHandleShapeAccepted):
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit accepted incomplete metadata")
	case preview.SurfaceSafeForHumanSmoke && (preview.KDEActionID != "xnix.runtime-status.controlled-launch" || preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" || preview.ForwardedArgumentKind != "evidence-relative-path"):
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit accepted the wrong route")
	case preview.KDEPolicyOwner || preview.OwnerServiceArgsExposedToKDE || preview.OwnerServiceArgumentsRecreated || preview.StateRootAccess || preview.ReceiptReconstruction:
		if preview.AuditState == "safe" {
			return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit accepted owner-only metadata")
		}
	case preview.RawExecutablePathExposed || preview.BackendCommandExposed || preview.BackendDetailsExposed:
		if preview.AuditState == "safe" {
			return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit accepted backend metadata")
		}
	case preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.DBusCalled || preview.KDEConfigurationWritten || preview.RuntimeStateWritten:
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit must not dispatch, write, or launch")
	case preview.HostRootModified || preview.DockerSocketMounted || preview.PrivilegedContainerRequired || preview.HostNetworkRequired || preview.NetworkRequired:
		return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit must keep host boundaries closed")
	}
	for _, value := range []string{
		preview.SchemaVersion,
		preview.RequestType,
		preview.Source,
		preview.RuntimeMethod,
		preview.ReadMethod,
		preview.AuditState,
		preview.KDEActionID,
		preview.KDEActionLabel,
		preview.PublicDBusMethod,
		preview.ForwardedArgumentKind,
		preview.BlockedReason,
		preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit requires single-line fields")
		}
	}
	for _, value := range append(append(append([]string{}, preview.UnsafeMetadataKeys...), preview.MissingMetadataKeys...), preview.BlockedActions...) {
		if strings.TrimSpace(value) == "" || !singleLine(value) {
			return KDEControlledLaunchActionSurfaceAuditPreview{}, errors.New("KDE controlled-launch action surface audit requires single-line list values")
		}
	}
	if err := validateNoBackendTerms(preview, "KDE controlled-launch action surface audit"); err != nil {
		return KDEControlledLaunchActionSurfaceAuditPreview{}, err
	}
	return preview, nil
}
