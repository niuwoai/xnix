package appidentity

import "errors"

const (
	KDEControlledLaunchActionSchemaVersion = "xnix.runtime.kde_controlled_launch_action.v1"
	KDEControlledLaunchActionRequestType   = "kde-controlled-launch-action-preview"
)

type KDEControlledLaunchActionRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
}

type KDEControlledLaunchActionPreview struct {
	SchemaVersion                     string   `json:"schema_version"`
	RequestType                       string   `json:"request_type"`
	Source                            string   `json:"source"`
	Desktop                           string   `json:"desktop"`
	KDEComponent                      string   `json:"kde_component"`
	KDEActionID                       string   `json:"kde_action_id"`
	KDEActionLabel                    string   `json:"kde_action_label"`
	KDEActionState                    string   `json:"kde_action_state"`
	RuntimeMethod                     string   `json:"runtime_method"`
	ReadMethod                        string   `json:"read_method"`
	ApplicationID                     string   `json:"application_id"`
	ApplicationName                   string   `json:"application_name"`
	ApplicationVersion                string   `json:"application_version"`
	EvidenceID                        string   `json:"evidence_id"`
	EvidenceRelativePath              string   `json:"evidence_relative_path"`
	EvidenceSHA256                    string   `json:"evidence_sha256"`
	EvidenceHandoffConsumed           bool     `json:"evidence_handoff_consumed"`
	EvidenceDigestVerified            bool     `json:"evidence_digest_verified"`
	PublicDBusService                 string   `json:"public_dbus_service"`
	PublicDBusObjectPath              string   `json:"public_dbus_object_path"`
	PublicDBusInterface               string   `json:"public_dbus_interface"`
	PublicDBusMethod                  string   `json:"public_dbus_method"`
	KDEForwardedArguments             []string `json:"kde_forwarded_arguments"`
	KDEForwardedArgumentKind          string   `json:"kde_forwarded_argument_kind"`
	KDEForwardsOnlyEvidenceHandle     bool     `json:"kde_forwards_only_evidence_handle"`
	DesktopEvidenceHandleForwarded    bool     `json:"desktop_evidence_handle_forwarded"`
	DesktopTriggerReady               bool     `json:"desktop_trigger_ready"`
	DesktopCallableRoute              string   `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod      string   `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType      string   `json:"desktop_callable_execution_type"`
	RuntimeOwned                      bool     `json:"runtime_owned"`
	GoRuntimeBacked                   bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                    bool     `json:"kde_policy_owner"`
	KDEOwnsOwnerServiceArgs           bool     `json:"kde_owns_owner_service_args"`
	OwnerServiceArgsExposedToKDE      bool     `json:"owner_service_args_exposed_to_kde"`
	OwnerServiceBoundaryHiddenFromKDE bool     `json:"owner_service_boundary_hidden_from_kde"`
	DesktopReceiptFieldsReconstructed bool     `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess         bool     `json:"desktop_kde_state_root_access"`
	StateRootPathExposed              bool     `json:"state_root_path_exposed"`
	EvidencePathExposed               bool     `json:"evidence_path_exposed"`
	ManagedLauncherPathExposed        bool     `json:"managed_launcher_path_exposed"`
	RawLauncherOutputExposed          bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed             bool     `json:"backend_details_exposed"`
	HostRootModified                  bool     `json:"host_root_modified"`
	DockerSocketMounted               bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired            bool     `json:"broad_host_mount_required"`
	PrivilegedContainerRequired       bool     `json:"privileged_container_required"`
	HostNetworkRequired               bool     `json:"host_network_required"`
	DesktopLaunchEnabled              bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled              bool     `json:"backend_launch_enabled"`
	ExecutionStarted                  bool     `json:"execution_started"`
	BackendProcessStarted             bool     `json:"backend_process_started"`
	RequestObjectCreatedByKDE         bool     `json:"request_object_created_by_kde"`
	RuntimePreviewCommand             []string `json:"runtime_preview_command"`
	BlockedActions                    []string `json:"blocked_actions"`
	DesktopSafeSummary                string   `json:"desktop_safe_summary"`
}

func PreviewKDEControlledLaunchAction(request KDEControlledLaunchActionRequest) (KDEControlledLaunchActionPreview, error) {
	trigger, err := PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		return KDEControlledLaunchActionPreview{}, err
	}
	preview := KDEControlledLaunchActionPreview{
		SchemaVersion:                     KDEControlledLaunchActionSchemaVersion,
		RequestType:                       KDEControlledLaunchActionRequestType,
		Source:                            KnownAppRuntimeStatusLaunchOwnerTriggerRequestType + "+kde-controlled-launch-action-stub",
		Desktop:                           "KDE Plasma",
		KDEComponent:                      "Compatibility Center",
		KDEActionID:                       "xnix.runtime-status.controlled-launch",
		KDEActionLabel:                    "Run with Xnix Runtime",
		KDEActionState:                    "ready-to-forward-evidence",
		RuntimeMethod:                     "PreviewKDEControlledLaunchAction",
		ReadMethod:                        "GetKDEControlledLaunchActionPreview",
		ApplicationID:                     trigger.AppID,
		ApplicationName:                   trigger.DisplayName,
		ApplicationVersion:                trigger.AppVersion,
		EvidenceID:                        trigger.EvidenceID,
		EvidenceRelativePath:              trigger.EvidenceRelativePath,
		EvidenceSHA256:                    trigger.EvidenceSHA256,
		EvidenceHandoffConsumed:           trigger.EvidenceHandoffConsumed,
		EvidenceDigestVerified:            trigger.EvidenceDigestVerified,
		PublicDBusService:                 "org.xnix.Compatibility1",
		PublicDBusObjectPath:              "/org/xnix/Compatibility1",
		PublicDBusInterface:               "org.xnix.Compatibility1",
		PublicDBusMethod:                  trigger.DesktopDBusMethod,
		KDEForwardedArguments:             []string{trigger.EvidenceRelativePath},
		KDEForwardedArgumentKind:          "evidence-relative-path",
		KDEForwardsOnlyEvidenceHandle:     true,
		DesktopEvidenceHandleForwarded:    trigger.DesktopEvidenceHandleForwarded,
		DesktopTriggerReady:               trigger.DesktopTriggerReady,
		DesktopCallableRoute:              trigger.DesktopCallableRoute,
		DesktopCallableRuntimeMethod:      trigger.DesktopCallableRuntimeMethod,
		DesktopCallableExecutionType:      trigger.DesktopCallableExecutionType,
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		KDEOwnsOwnerServiceArgs:           false,
		OwnerServiceArgsExposedToKDE:      false,
		OwnerServiceBoundaryHiddenFromKDE: true,
		DesktopReceiptFieldsReconstructed: false,
		DesktopKDEStateRootAccess:         false,
		StateRootPathExposed:              false,
		EvidencePathExposed:               false,
		ManagedLauncherPathExposed:        false,
		RawLauncherOutputExposed:          false,
		BackendDetailsExposed:             false,
		HostRootModified:                  false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		PrivilegedContainerRequired:       false,
		HostNetworkRequired:               false,
		DesktopLaunchEnabled:              false,
		BackendLaunchEnabled:              false,
		ExecutionStarted:                  false,
		BackendProcessStarted:             false,
		RequestObjectCreatedByKDE:         false,
		RuntimePreviewCommand:             []string{"xnix-runtime-go", KDEControlledLaunchActionRequestType, "--evidence-relative-path", trigger.EvidenceRelativePath},
		BlockedActions:                    []string{"derive owner service arguments in KDE", "read Runtime state root from KDE", "reconstruct Runtime receipts in KDE", "start compatibility engine from KDE", "expose raw launcher output to KDE", "mutate host root from KDE action stub"},
		DesktopSafeSummary:                trigger.DisplayName + " can be presented as a KDE controlled-launch action that forwards only the Runtime-status evidence handle to D-Bus.",
	}
	return validateKDEControlledLaunchActionPreview(preview)
}

func validateKDEControlledLaunchActionPreview(preview KDEControlledLaunchActionPreview) (KDEControlledLaunchActionPreview, error) {
	switch {
	case preview.SchemaVersion != KDEControlledLaunchActionSchemaVersion || preview.RequestType != KDEControlledLaunchActionRequestType:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action has invalid schema")
	case preview.Desktop != "KDE Plasma" || preview.KDEComponent == "" || preview.KDEActionID == "" || preview.KDEActionLabel == "":
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires KDE action identity")
	case preview.ApplicationID == "" || preview.ApplicationName == "" || preview.ApplicationVersion == "":
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires application identity")
	case preview.EvidenceRelativePath == "" || preview.EvidenceSHA256 == "" || !preview.EvidenceHandoffConsumed || !preview.EvidenceDigestVerified:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires verified evidence")
	case preview.PublicDBusService != "org.xnix.Compatibility1" || preview.PublicDBusObjectPath != "/org/xnix/Compatibility1" || preview.PublicDBusInterface != "org.xnix.Compatibility1" || preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch":
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires the public Runtime D-Bus method")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{preview.EvidenceRelativePath}) || preview.KDEForwardedArgumentKind != "evidence-relative-path":
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must forward only the evidence handle")
	case !preview.KDEForwardsOnlyEvidenceHandle || !preview.DesktopEvidenceHandleForwarded || !preview.DesktopTriggerReady:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must be an evidence-only desktop trigger")
	case preview.DesktopCallableRoute != "kde-dbus-runtime-status-action" || preview.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" || preview.DesktopCallableExecutionType != KnownAppKDERuntimeStatusLaunchExecutionRequestType:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must preserve Runtime trigger metadata")
	case !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || preview.KDEOwnsOwnerServiceArgs || preview.OwnerServiceArgsExposedToKDE || !preview.OwnerServiceBoundaryHiddenFromKDE:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must keep Runtime ownership")
	case preview.DesktopReceiptFieldsReconstructed || preview.DesktopKDEStateRootAccess || preview.StateRootPathExposed || preview.EvidencePathExposed || preview.ManagedLauncherPathExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must not expose Runtime internals")
	case preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action must keep host/container gates closed")
	case preview.DesktopLaunchEnabled || preview.BackendLaunchEnabled || preview.ExecutionStarted || preview.BackendProcessStarted || preview.RequestObjectCreatedByKDE:
		return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action stub must not start execution")
	}
	for _, value := range []string{
		preview.SchemaVersion,
		preview.RequestType,
		preview.Source,
		preview.Desktop,
		preview.KDEComponent,
		preview.KDEActionID,
		preview.KDEActionLabel,
		preview.KDEActionState,
		preview.RuntimeMethod,
		preview.ReadMethod,
		preview.ApplicationID,
		preview.ApplicationName,
		preview.ApplicationVersion,
		preview.EvidenceID,
		preview.EvidenceRelativePath,
		preview.EvidenceSHA256,
		preview.PublicDBusService,
		preview.PublicDBusObjectPath,
		preview.PublicDBusInterface,
		preview.PublicDBusMethod,
		preview.KDEForwardedArgumentKind,
		preview.DesktopCallableRoute,
		preview.DesktopCallableRuntimeMethod,
		preview.DesktopCallableExecutionType,
		preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires single-line fields")
		}
	}
	for _, value := range append(append([]string{}, preview.KDEForwardedArguments...), append(preview.RuntimePreviewCommand, preview.BlockedActions...)...) {
		if value != "" && !singleLine(value) {
			return KDEControlledLaunchActionPreview{}, errors.New("KDE controlled launch action requires single-line list values")
		}
	}
	if err := validateNoBackendTerms(preview, "KDE controlled launch action"); err != nil {
		return KDEControlledLaunchActionPreview{}, err
	}
	return preview, nil
}
