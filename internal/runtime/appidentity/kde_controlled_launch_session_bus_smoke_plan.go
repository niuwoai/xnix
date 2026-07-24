package appidentity

import "errors"

const (
	KDEControlledLaunchSessionBusSmokePlanSchemaVersion = "xnix.runtime.kde_controlled_launch_session_bus_smoke_plan.v1"
	KDEControlledLaunchSessionBusSmokePlanRequestType   = "kde-controlled-launch-session-bus-smoke-plan-preview"
)

type KDEControlledLaunchSessionBusSmokePlanRequest struct {
	StateRoot            string
	EvidenceID           string
	EvidenceRelativePath string
}

type KDEControlledLaunchSessionBusSmokePlanPreview struct {
	SchemaVersion                         string   `json:"schema_version"`
	RequestType                           string   `json:"request_type"`
	Source                                string   `json:"source"`
	RuntimeMethod                         string   `json:"runtime_method"`
	ReadMethod                            string   `json:"read_method"`
	KDEActionID                           string   `json:"kde_action_id"`
	KDEActionLabel                        string   `json:"kde_action_label"`
	KDEActionPreviewRequestType           string   `json:"kde_action_preview_request_type"`
	PublicDBusService                     string   `json:"public_dbus_service"`
	PublicDBusObjectPath                  string   `json:"public_dbus_object_path"`
	PublicDBusInterface                   string   `json:"public_dbus_interface"`
	PublicDBusMethod                      string   `json:"public_dbus_method"`
	EvidenceRelativePath                  string   `json:"evidence_relative_path"`
	EvidenceSHA256                        string   `json:"evidence_sha256"`
	EvidenceHandoffConsumed               bool     `json:"evidence_handoff_consumed"`
	EvidenceDigestVerified                bool     `json:"evidence_digest_verified"`
	KDEForwardedArguments                 []string `json:"kde_forwarded_arguments"`
	KDEForwardedArgumentKind              string   `json:"kde_forwarded_argument_kind"`
	KDEForwardsOnlyEvidenceHandle         bool     `json:"kde_forwards_only_evidence_handle"`
	RestrictedSessionBusPlanReady         bool     `json:"restricted_session_bus_plan_ready"`
	PrivateSessionBusRequired             bool     `json:"private_session_bus_required"`
	DBusSessionBusAddressRequired         bool     `json:"dbus_session_bus_address_required"`
	OuterPrivateSessionBusCommand         []string `json:"outer_private_session_bus_command"`
	RuntimeStatusOwnerSessionSmokeCommand []string `json:"runtime_status_owner_session_smoke_command"`
	InnerStagedLauncherDispatchCommand    []string `json:"inner_staged_launcher_dispatch_command"`
	ContainerSmokeCommand                 []string `json:"container_smoke_command"`
	DBusControlledLaunchFixturePlanReady  bool     `json:"dbus_controlled_launch_fixture_plan_ready"`
	DBusControlledLaunchFixtureCommand    []string `json:"dbus_controlled_launch_fixture_command"`
	DBusControlledLaunchFixtureContainer  []string `json:"dbus_controlled_launch_fixture_container_command"`
	ExpectedPassMarker                    string   `json:"expected_pass_marker"`
	ExpectedSkipMarker                    string   `json:"expected_skip_marker"`
	ExpectedDBusFixturePassMarker         string   `json:"expected_dbus_fixture_pass_marker"`
	ExpectedDBusFixtureSkipMarker         string   `json:"expected_dbus_fixture_skip_marker"`
	RuntimeOwned                          bool     `json:"runtime_owned"`
	GoRuntimeBacked                       bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                        bool     `json:"kde_policy_owner"`
	OwnerServiceArgsExposedToKDE          bool     `json:"owner_service_args_exposed_to_kde"`
	DesktopKDEStateRootAccess             bool     `json:"desktop_kde_state_root_access"`
	DesktopReceiptFieldsReconstructed     bool     `json:"desktop_receipt_fields_reconstructed"`
	StateRootPathExposed                  bool     `json:"state_root_path_exposed"`
	RawLauncherOutputExposed              bool     `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed                 bool     `json:"backend_details_exposed"`
	HostRootModified                      bool     `json:"host_root_modified"`
	DockerSocketMounted                   bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired                bool     `json:"broad_host_mount_required"`
	PrivilegedContainerRequired           bool     `json:"privileged_container_required"`
	HostNetworkRequired                   bool     `json:"host_network_required"`
	ExecutionStarted                      bool     `json:"execution_started"`
	BackendProcessStarted                 bool     `json:"backend_process_started"`
	SmokeExecutedByPreview                bool     `json:"smoke_executed_by_preview"`
	DesktopSafeSummary                    string   `json:"desktop_safe_summary"`
}

func PreviewKDEControlledLaunchSessionBusSmokePlan(request KDEControlledLaunchSessionBusSmokePlanRequest) (KDEControlledLaunchSessionBusSmokePlanPreview, error) {
	action, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:            request.StateRoot,
		EvidenceID:           request.EvidenceID,
		EvidenceRelativePath: request.EvidenceRelativePath,
	})
	if err != nil {
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, err
	}
	preview := KDEControlledLaunchSessionBusSmokePlanPreview{
		SchemaVersion:                         KDEControlledLaunchSessionBusSmokePlanSchemaVersion,
		RequestType:                           KDEControlledLaunchSessionBusSmokePlanRequestType,
		Source:                                KDEControlledLaunchActionRequestType + "+restricted-session-bus-smoke-plan",
		RuntimeMethod:                         "PreviewKDEControlledLaunchSessionBusSmokePlan",
		ReadMethod:                            "GetKDEControlledLaunchSessionBusSmokePlan",
		KDEActionID:                           action.KDEActionID,
		KDEActionLabel:                        action.KDEActionLabel,
		KDEActionPreviewRequestType:           action.RequestType,
		PublicDBusService:                     action.PublicDBusService,
		PublicDBusObjectPath:                  action.PublicDBusObjectPath,
		PublicDBusInterface:                   action.PublicDBusInterface,
		PublicDBusMethod:                      action.PublicDBusMethod,
		EvidenceRelativePath:                  action.EvidenceRelativePath,
		EvidenceSHA256:                        action.EvidenceSHA256,
		EvidenceHandoffConsumed:               action.EvidenceHandoffConsumed,
		EvidenceDigestVerified:                action.EvidenceDigestVerified,
		KDEForwardedArguments:                 append([]string{}, action.KDEForwardedArguments...),
		KDEForwardedArgumentKind:              action.KDEForwardedArgumentKind,
		KDEForwardsOnlyEvidenceHandle:         action.KDEForwardsOnlyEvidenceHandle,
		RestrictedSessionBusPlanReady:         true,
		PrivateSessionBusRequired:             true,
		DBusSessionBusAddressRequired:         true,
		OuterPrivateSessionBusCommand:         []string{"dbus-run-session", "--", "ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"},
		RuntimeStatusOwnerSessionSmokeCommand: []string{"ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"},
		InnerStagedLauncherDispatchCommand:    []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"},
		ContainerSmokeCommand:                 []string{"ruby", "scripts/container.rb", "runtime-status-owner-service-session-bus-smoke"},
		DBusControlledLaunchFixturePlanReady:  true,
		DBusControlledLaunchFixtureCommand:    []string{"ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"},
		DBusControlledLaunchFixtureContainer:  []string{"ruby", "scripts/container.rb", "dbus-controlled-launch-owner-fixture-smoke"},
		ExpectedPassMarker:                    "PASS: Runtime-status owner service session-bus smoke",
		ExpectedSkipMarker:                    "SKIP: Runtime-status owner service session-bus smoke",
		ExpectedDBusFixturePassMarker:         "PASS: D-Bus controlled launch owner fixture smoke",
		ExpectedDBusFixtureSkipMarker:         "SKIP: D-Bus controlled launch owner fixture smoke",
		RuntimeOwned:                          true,
		GoRuntimeBacked:                       true,
		KDEPolicyOwner:                        false,
		OwnerServiceArgsExposedToKDE:          false,
		DesktopKDEStateRootAccess:             false,
		DesktopReceiptFieldsReconstructed:     false,
		StateRootPathExposed:                  false,
		RawLauncherOutputExposed:              false,
		BackendDetailsExposed:                 false,
		HostRootModified:                      false,
		DockerSocketMounted:                   false,
		BroadHostMountRequired:                false,
		PrivilegedContainerRequired:           false,
		HostNetworkRequired:                   false,
		ExecutionStarted:                      false,
		BackendProcessStarted:                 false,
		SmokeExecutedByPreview:                false,
		DesktopSafeSummary:                    action.KDEActionLabel + " has a restricted private session-bus smoke plan that forwards only Runtime-status evidence.",
	}
	return validateKDEControlledLaunchSessionBusSmokePlan(preview)
}

func validateKDEControlledLaunchSessionBusSmokePlan(preview KDEControlledLaunchSessionBusSmokePlanPreview) (KDEControlledLaunchSessionBusSmokePlanPreview, error) {
	switch {
	case preview.SchemaVersion != KDEControlledLaunchSessionBusSmokePlanSchemaVersion || preview.RequestType != KDEControlledLaunchSessionBusSmokePlanRequestType:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan has invalid schema")
	case preview.KDEActionID != "xnix.runtime-status.controlled-launch" || preview.KDEActionLabel == "" || preview.KDEActionPreviewRequestType != KDEControlledLaunchActionRequestType:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires KDE action identity")
	case preview.PublicDBusService != "org.xnix.Compatibility1" || preview.PublicDBusObjectPath != "/org/xnix/Compatibility1" || preview.PublicDBusInterface != "org.xnix.Compatibility1" || preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch":
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the public Runtime D-Bus method")
	case preview.EvidenceRelativePath == "" || preview.EvidenceSHA256 == "" || !preview.EvidenceHandoffConsumed || !preview.EvidenceDigestVerified:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires verified evidence")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{preview.EvidenceRelativePath}) || preview.KDEForwardedArgumentKind != "evidence-relative-path" || !preview.KDEForwardsOnlyEvidenceHandle:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan must forward only the evidence handle")
	case !preview.RestrictedSessionBusPlanReady || !preview.PrivateSessionBusRequired || !preview.DBusSessionBusAddressRequired:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires a private session bus")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.OuterPrivateSessionBusCommand, []string{"dbus-run-session", "--", "ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"}):
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the private session-bus runner")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.RuntimeStatusOwnerSessionSmokeCommand, []string{"ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"}) || !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.InnerStagedLauncherDispatchCommand, []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"}):
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the staged launcher smoke chain")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.ContainerSmokeCommand, []string{"ruby", "scripts/container.rb", "runtime-status-owner-service-session-bus-smoke"}):
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the restricted container smoke command")
	case !preview.DBusControlledLaunchFixturePlanReady || !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.DBusControlledLaunchFixtureCommand, []string{"ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"}):
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the D-Bus controlled-launch fixture command")
	case !sameRuntimeStatusLaunchOwnerFixtureArgs(preview.DBusControlledLaunchFixtureContainer, []string{"ruby", "scripts/container.rb", "dbus-controlled-launch-owner-fixture-smoke"}):
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires the restricted D-Bus fixture container command")
	case preview.ExpectedPassMarker != "PASS: Runtime-status owner service session-bus smoke" || preview.ExpectedSkipMarker != "SKIP: Runtime-status owner service session-bus smoke":
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires stable result markers")
	case preview.ExpectedDBusFixturePassMarker != "PASS: D-Bus controlled launch owner fixture smoke" || preview.ExpectedDBusFixtureSkipMarker != "SKIP: D-Bus controlled launch owner fixture smoke":
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires stable D-Bus fixture result markers")
	case !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || preview.OwnerServiceArgsExposedToKDE:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan must preserve Runtime ownership")
	case preview.DesktopKDEStateRootAccess || preview.DesktopReceiptFieldsReconstructed || preview.StateRootPathExposed || preview.RawLauncherOutputExposed || preview.BackendDetailsExposed:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan must not expose Runtime internals")
	case preview.HostRootModified || preview.DockerSocketMounted || preview.BroadHostMountRequired || preview.PrivilegedContainerRequired || preview.HostNetworkRequired:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan must keep host/container gates closed")
	case preview.ExecutionStarted || preview.BackendProcessStarted || preview.SmokeExecutedByPreview:
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan preview must not execute")
	}
	for _, value := range []string{
		preview.SchemaVersion,
		preview.RequestType,
		preview.Source,
		preview.RuntimeMethod,
		preview.ReadMethod,
		preview.KDEActionID,
		preview.KDEActionLabel,
		preview.KDEActionPreviewRequestType,
		preview.PublicDBusService,
		preview.PublicDBusObjectPath,
		preview.PublicDBusInterface,
		preview.PublicDBusMethod,
		preview.EvidenceRelativePath,
		preview.EvidenceSHA256,
		preview.KDEForwardedArgumentKind,
		preview.ExpectedPassMarker,
		preview.ExpectedSkipMarker,
		preview.ExpectedDBusFixturePassMarker,
		preview.ExpectedDBusFixtureSkipMarker,
		preview.DesktopSafeSummary,
	} {
		if value != "" && !singleLine(value) {
			return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires single-line fields")
		}
	}
	listValues := append([]string{}, preview.KDEForwardedArguments...)
	listValues = append(listValues, preview.OuterPrivateSessionBusCommand...)
	listValues = append(listValues, preview.RuntimeStatusOwnerSessionSmokeCommand...)
	listValues = append(listValues, preview.InnerStagedLauncherDispatchCommand...)
	listValues = append(listValues, preview.ContainerSmokeCommand...)
	listValues = append(listValues, preview.DBusControlledLaunchFixtureCommand...)
	listValues = append(listValues, preview.DBusControlledLaunchFixtureContainer...)
	for _, value := range listValues {
		if value != "" && !singleLine(value) {
			return KDEControlledLaunchSessionBusSmokePlanPreview{}, errors.New("KDE controlled launch session-bus smoke plan requires single-line list values")
		}
	}
	if err := validateNoBackendTerms(preview, "KDE controlled launch session-bus smoke plan"); err != nil {
		return KDEControlledLaunchSessionBusSmokePlanPreview{}, err
	}
	return preview, nil
}
