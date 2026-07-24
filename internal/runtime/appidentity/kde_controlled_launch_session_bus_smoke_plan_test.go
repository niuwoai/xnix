package appidentity

import "testing"

func TestPreviewKDEControlledLaunchSessionBusSmokePlanLinksActionToRestrictedSmoke(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionEvidence(t, stateRoot)

	preview, err := PreviewKDEControlledLaunchSessionBusSmokePlan(KDEControlledLaunchSessionBusSmokePlanRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchSessionBusSmokePlan returned error: %v", err)
	}

	if preview.SchemaVersion != KDEControlledLaunchSessionBusSmokePlanSchemaVersion ||
		preview.RequestType != KDEControlledLaunchSessionBusSmokePlanRequestType ||
		preview.Source != KDEControlledLaunchActionRequestType+"+restricted-session-bus-smoke-plan" ||
		preview.RuntimeMethod != "PreviewKDEControlledLaunchSessionBusSmokePlan" ||
		preview.ReadMethod != "GetKDEControlledLaunchSessionBusSmokePlan" ||
		preview.KDEActionID != "xnix.runtime-status.controlled-launch" ||
		preview.KDEActionLabel != "Run with Xnix Runtime" ||
		preview.KDEActionPreviewRequestType != KDEControlledLaunchActionRequestType ||
		preview.PublicDBusService != "org.xnix.Compatibility1" ||
		preview.PublicDBusObjectPath != "/org/xnix/Compatibility1" ||
		preview.PublicDBusInterface != "org.xnix.Compatibility1" ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceHandoffConsumed ||
		!preview.EvidenceDigestVerified ||
		preview.KDEForwardedArgumentKind != "evidence-relative-path" ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{record.EvidenceRelativePath}) ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		!preview.RestrictedSessionBusPlanReady ||
		!preview.PrivateSessionBusRequired ||
		!preview.DBusSessionBusAddressRequired ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.OuterPrivateSessionBusCommand, []string{"dbus-run-session", "--", "ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.RuntimeStatusOwnerSessionSmokeCommand, []string{"ruby", "scripts/runtime_status_owner_service_session_bus_smoke.rb"}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.InnerStagedLauncherDispatchCommand, []string{"ruby", "scripts/staged_launcher_dispatch_smoke.rb"}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.ContainerSmokeCommand, []string{"ruby", "scripts/container.rb", "runtime-status-owner-service-session-bus-smoke"}) ||
		!preview.DBusControlledLaunchFixturePlanReady ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.DBusControlledLaunchFixtureCommand, []string{"ruby", "scripts/dbus_controlled_launch_owner_fixture_smoke.rb"}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.DBusControlledLaunchFixtureContainer, []string{"ruby", "scripts/container.rb", "dbus-controlled-launch-owner-fixture-smoke"}) ||
		preview.ExpectedPassMarker != "PASS: Runtime-status owner service session-bus smoke" ||
		preview.ExpectedSkipMarker != "SKIP: Runtime-status owner service session-bus smoke" ||
		preview.ExpectedDBusFixturePassMarker != "PASS: D-Bus controlled launch owner fixture smoke" ||
		preview.ExpectedDBusFixtureSkipMarker != "SKIP: D-Bus controlled launch owner fixture smoke" ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.OwnerServiceArgsExposedToKDE {
		t.Fatalf("unexpected KDE controlled launch session-bus smoke plan: %#v", preview)
	}

	if preview.DesktopKDEStateRootAccess ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.StateRootPathExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.SmokeExecutedByPreview {
		t.Fatalf("KDE controlled launch session-bus smoke plan opened unsafe gates: %#v", preview)
	}
}
