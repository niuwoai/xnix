package appidentity

import (
	"errors"
	"strings"
)

const (
	KDEControlledLaunchActionSchemaVersion = "xnix.runtime.kde_controlled_launch_action.v1"
	KDEControlledLaunchActionRequestType   = "kde-controlled-launch-action-preview"
)

type KDEControlledLaunchActionRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
	KDECenterGUICard     *KDECenterPageKnownAppMatrixCard
	KDECenterPage        *KDECenterPagePreview
	KDECenterPageAppID   string
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
	kdeCenterGUICard := request.KDECenterGUICard
	if kdeCenterGUICard == nil && request.KDECenterPage != nil {
		card, err := kdeControlledLaunchActionGUICardFromPage(*request.KDECenterPage, request.KDECenterPageAppID)
		if err != nil {
			return KDEControlledLaunchActionPreview{}, err
		}
		kdeCenterGUICard = &card
	}
	evidenceRelativePath := strings.TrimSpace(request.EvidenceRelativePath)
	if evidenceRelativePath == "" && kdeCenterGUICard != nil {
		evidenceRelativePath = kdeControlledLaunchActionEvidenceRelativePathFromCard(*kdeCenterGUICard)
	}
	trigger, err := PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: evidenceRelativePath,
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
	preview, err = validateKDEControlledLaunchActionPreview(preview)
	if err != nil {
		return KDEControlledLaunchActionPreview{}, err
	}
	if kdeCenterGUICard != nil {
		if err := validateKDEControlledLaunchActionGUICard(*kdeCenterGUICard, preview); err != nil {
			return KDEControlledLaunchActionPreview{}, err
		}
	}
	return preview, nil
}

func kdeControlledLaunchActionGUICardFromPage(page KDECenterPagePreview, appID string) (KDECenterPageKnownAppMatrixCard, error) {
	targetAppID := strings.TrimSpace(appID)
	if targetAppID == "" {
		targetAppID = strings.TrimSpace(page.ApplicationID)
	}
	switch {
	case page.SchemaVersion != "xnix.runtime.kde_center_page.v1" || page.RequestType != "kde-center-page-preview":
		return KDECenterPageKnownAppMatrixCard{}, errors.New("KDE controlled launch action center page has invalid schema")
	case !page.RuntimeOwned || !page.GoRuntimeBacked || page.KDEPolicyOwner:
		return KDECenterPageKnownAppMatrixCard{}, errors.New("KDE controlled launch action center page must remain Runtime-owned")
	case page.HostRootModified || page.BackendDetailsExposed || page.LaunchEnabled || page.ExecutionStarted || page.BackendProcessStarted:
		return KDECenterPageKnownAppMatrixCard{}, errors.New("KDE controlled launch action center page must keep launch and backend gates closed")
	case targetAppID == "":
		return KDECenterPageKnownAppMatrixCard{}, errors.New("KDE controlled launch action center page requires an application id")
	}

	var matches []KDECenterPageKnownAppMatrixCard
	for _, card := range page.KnownAppGUIEvidenceCards {
		if strings.TrimSpace(card.AppID) == targetAppID {
			matches = append(matches, card)
		}
	}
	if len(matches) != 1 {
		return KDECenterPageKnownAppMatrixCard{}, errors.New("KDE controlled launch action center page requires exactly one matching GUI evidence card")
	}
	return matches[0], nil
}

func kdeControlledLaunchActionEvidenceRelativePathFromCard(card KDECenterPageKnownAppMatrixCard) string {
	if len(card.KDEForwardedArguments) == 1 && strings.TrimSpace(card.KDEForwardedArguments[0]) != "" {
		return strings.TrimSpace(card.KDEForwardedArguments[0])
	}
	return strings.TrimSpace(card.OwnerEvidenceRelativePath)
}

func validateKDEControlledLaunchActionGUICard(card KDECenterPageKnownAppMatrixCard, preview KDEControlledLaunchActionPreview) error {
	switch {
	case card.EvidenceKind != "known-application-gui-smoke" || card.EvidenceSource != "wine-guest-gui-smoke":
		return errors.New("KDE controlled launch action GUI card requires GUI smoke evidence")
	case card.AppID != preview.ApplicationID || card.DisplayName != preview.ApplicationName || card.AppVersion != preview.ApplicationVersion:
		return errors.New("KDE controlled launch action GUI card must match Runtime evidence identity")
	case card.PrimaryActionID != KnownAppKDERuntimeStatusLaunchAction || card.PrimaryActionKind != "runtime-status" || !card.PrimaryActionEnabled:
		return errors.New("KDE controlled launch action GUI card requires Runtime-status primary action")
	case card.DesktopCallableRoute != preview.DesktopCallableRoute || card.DesktopCallableRuntimeMethod != preview.DesktopCallableRuntimeMethod || card.DesktopCallableExecutionType != preview.DesktopCallableExecutionType:
		return errors.New("KDE controlled launch action GUI card route does not match Runtime action")
	case card.DesktopDBusMethod != preview.PublicDBusMethod || !card.DesktopEvidenceHandleForwarded || card.KDEForwardedArgumentKind != "evidence-relative-path":
		return errors.New("KDE controlled launch action GUI card must forward the public Runtime D-Bus evidence handle")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(card.KDEForwardedArguments, []string{preview.EvidenceRelativePath}):
		return errors.New("KDE controlled launch action GUI card must forward only the Runtime evidence relative path")
	case card.OwnerServiceArgsExposedToKDE || card.DesktopLaunchEnabled || card.BackendLaunchEnabled:
		return errors.New("KDE controlled launch action GUI card must keep owner args and launch gates closed")
	case card.HostRootModified || card.BackendDetailsExposed || card.RawArtifactPathExposed || card.KDEPolicyOwner:
		return errors.New("KDE controlled launch action GUI card must keep unsafe desktop fields closed")
	case !card.RuntimeOwned || !card.GoRuntimeBacked || !card.ExecutionEvidenceRecorded || !card.RuntimeDispatchVerified:
		return errors.New("KDE controlled launch action GUI card requires Runtime-owned verified evidence")
	case !card.OwnerControlledRuntimeLaunchVerified || !card.OwnerManagedCopyVerified || !card.OwnerServiceCallReady || !card.OwnerEvidenceHandoffReady:
		return errors.New("KDE controlled launch action GUI card requires owner-controlled handoff readiness")
	case card.OwnerEvidenceRelativePath != preview.EvidenceRelativePath:
		return errors.New("KDE controlled launch action GUI card owner evidence path must match Runtime evidence")
	}
	for _, value := range []string{
		card.AppID,
		card.DisplayName,
		card.AppVersion,
		card.EvidenceKind,
		card.EvidenceSource,
		card.SmokeStatus,
		card.CompatibilityState,
		card.CenterCardState,
		card.PrimaryActionID,
		card.PrimaryActionLabel,
		card.PrimaryActionKind,
		card.DesktopCallableRoute,
		card.DesktopCallableRuntimeMethod,
		card.DesktopCallableExecutionType,
		card.DesktopDBusMethod,
		card.KDEForwardedArgumentKind,
		card.OwnerEvidenceRelativePath,
	} {
		if value != "" && !singleLine(value) {
			return errors.New("KDE controlled launch action GUI card requires single-line fields")
		}
	}
	for _, value := range card.KDEForwardedArguments {
		if value != "" && !singleLine(value) {
			return errors.New("KDE controlled launch action GUI card requires single-line forwarded arguments")
		}
	}
	return nil
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
