package appidentity

import "testing"

func TestPreviewKnownAppRuntimeStatusLaunchOwnerTriggerConsumesVerifiedHandoff(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
		AppVersion:                             "26.02",
		RequestType:                            "windows-known-app-dispatch-smoke",
		Status:                                 "passed",
		GuestBoundary:                          "managed-known-app-guest-smoke",
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       true,
		MarkerObserved:                         true,
		SessionGatedControlledDispatchConsumed: true,
		SessionGatedControlledDispatchState:    "created-after-session-gated-review",
		SessionGatedReviewReceiptID:            reviewReceiptID,
		LaunchAuthorizationReceiptID:           launchReceiptID,
		LaunchAuthorizationReceiptState:        "recorded",
		LaunchGateState:                        "controlled-dispatch-ready",
		LaunchGateConsumed:                     true,
		LaunchGateReceiptAccepted:              true,
		LaunchGateGuestBoundaryAccepted:        true,
		ControlledDispatchReady:                true,
		ControlledExecutionSessionConsumed:     true,
		ControlledExecutionSessionID:           sessionID,
		ControlledSessionDigestVerified:        true,
		ControlledSessionRelativePath:          "execution-ledger/sessions/" + sessionID + ".json",
		RuntimeOwnerConsumableSession:          true,
		KDEReadModelConsumableSession:          true,
	})
	if err != nil {
		t.Fatalf("ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence returned error: %v", err)
	}
	record, err := RecordKnownAppKDERuntimeStatusLaunchEvidence(KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	trigger, err := PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppRuntimeStatusLaunchOwnerTrigger returned error: %v", err)
	}
	if trigger.SchemaVersion != KnownAppRuntimeStatusLaunchOwnerTriggerSchemaVersion ||
		trigger.RequestType != KnownAppRuntimeStatusLaunchOwnerTriggerRequestType ||
		trigger.RuntimeMethod != "PreviewKnownAppRuntimeStatusLaunchOwnerTrigger" ||
		trigger.ReadMethod != "GetKnownAppRuntimeStatusLaunchOwnerTrigger" ||
		trigger.EvidenceRelativePath != record.EvidenceRelativePath ||
		trigger.EvidenceSHA256 != record.EvidenceSHA256 ||
		!trigger.EvidenceHandoffConsumed ||
		!trigger.EvidenceDigestVerified ||
		!trigger.DesktopTriggerReady ||
		trigger.DesktopCallableRoute != "kde-dbus-runtime-status-action" ||
		trigger.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		trigger.DesktopCallableExecutionType != KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		trigger.DesktopDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		!trigger.OwnerServiceCallReady ||
		trigger.OwnerServiceBoundary != "go-runtime-owner-in-process-service" ||
		trigger.OwnerServiceMethod != "ShowRuntimeControlledLaunch" ||
		trigger.OwnerServiceCallType != "desktop-action-dispatch" ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(trigger.OwnerServiceCallArgs, []string{"ShowRuntimeControlledLaunch", "evidence-relative-path", record.EvidenceRelativePath}) ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(trigger.OwnerServiceCLIArgs, []string{"--service-call", "ShowRuntimeControlledLaunch", "evidence-relative-path", record.EvidenceRelativePath}) ||
		!trigger.RuntimeOwnerServiceSuppliesInputs ||
		!trigger.RuntimeOwned ||
		!trigger.RuntimeOwnedDispatch ||
		!trigger.GoRuntimeBacked ||
		trigger.KDEPolicyOwner ||
		!trigger.KDEForwardsOnlyEvidenceHandle ||
		trigger.DesktopReceiptFieldsReconstructed ||
		trigger.DesktopKDEStateRootAccess {
		t.Fatalf("unexpected owner trigger preview: %#v", trigger)
	}
	if trigger.StateRootPathExposed ||
		trigger.EvidencePathExposed ||
		trigger.ManagedLauncherPathExposed ||
		trigger.RawLauncherOutputExposed ||
		trigger.BackendDetailsExposed ||
		trigger.HostRootModified ||
		trigger.DockerSocketMounted ||
		trigger.BroadHostMountRequired ||
		trigger.DesktopLaunchEnabled ||
		trigger.BackendLaunchEnabled ||
		trigger.ExecutionStarted ||
		trigger.BackendProcessStarted {
		t.Fatalf("owner trigger preview opened unsafe gates: %#v", trigger)
	}
}
