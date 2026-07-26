package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/winapp"
)

func TestPreviewKDEControlledLaunchActionForwardsOnlyEvidenceHandle(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionEvidence(t, stateRoot)

	preview, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchAction returned error: %v", err)
	}

	if preview.SchemaVersion != KDEControlledLaunchActionSchemaVersion ||
		preview.RequestType != KDEControlledLaunchActionRequestType ||
		preview.Source != KnownAppRuntimeStatusLaunchOwnerTriggerRequestType+"+kde-controlled-launch-action-stub" ||
		preview.Desktop != "KDE Plasma" ||
		preview.KDEComponent != "Compatibility Center" ||
		preview.KDEActionID != "xnix.runtime-status.controlled-launch" ||
		preview.KDEActionLabel != "Run with Xnix Runtime" ||
		preview.KDEActionState != "ready-to-forward-evidence" ||
		preview.RuntimeMethod != "PreviewKDEControlledLaunchAction" ||
		preview.ReadMethod != "GetKDEControlledLaunchActionPreview" ||
		preview.ApplicationID != "7zr" ||
		preview.ApplicationName != "7-Zip standalone console executable" ||
		preview.ApplicationVersion != "26.02" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		!preview.EvidenceHandoffConsumed ||
		!preview.EvidenceDigestVerified ||
		preview.PublicDBusService != "org.xnix.Compatibility1" ||
		preview.PublicDBusObjectPath != "/org/xnix/Compatibility1" ||
		preview.PublicDBusInterface != "org.xnix.Compatibility1" ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.KDEForwardedArgumentKind != "evidence-relative-path" ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{record.EvidenceRelativePath}) ||
		!preview.KDEForwardsOnlyEvidenceHandle ||
		!preview.DesktopEvidenceHandleForwarded ||
		!preview.DesktopTriggerReady ||
		preview.DesktopCallableRoute != "kde-dbus-runtime-status-action" ||
		preview.DesktopCallableRuntimeMethod != "ShowRuntimeControlledLaunch" ||
		preview.DesktopCallableExecutionType != KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEOwnsOwnerServiceArgs ||
		preview.OwnerServiceArgsExposedToKDE ||
		!preview.OwnerServiceBoundaryHiddenFromKDE ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.RuntimePreviewCommand, []string{"xnix-runtime-go", KDEControlledLaunchActionRequestType, "--evidence-relative-path", record.EvidenceRelativePath}) {
		t.Fatalf("unexpected KDE controlled launch action preview: %#v", preview)
	}

	if preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.StateRootPathExposed ||
		preview.EvidencePathExposed ||
		preview.ManagedLauncherPathExposed ||
		preview.RawLauncherOutputExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkRequired ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RequestObjectCreatedByKDE {
		t.Fatalf("KDE controlled launch action preview opened unsafe gates: %#v", preview)
	}
}

func TestPreviewKDEControlledLaunchActionConsumesKDEGUICardRoute(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionGUICard(record)

	preview, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:        stateRoot,
		KDECenterGUICard: &card,
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchAction returned error: %v", err)
	}

	if preview.ApplicationID != "org.xnix.apps.messagebox" ||
		preview.ApplicationName != "Xnix MessageBox" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		preview.PublicDBusMethod != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		preview.DesktopCallableRoute != card.DesktopCallableRoute ||
		preview.DesktopCallableRuntimeMethod != card.DesktopCallableRuntimeMethod ||
		preview.DesktopCallableExecutionType != card.DesktopCallableExecutionType ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{record.EvidenceRelativePath}) ||
		preview.OwnerServiceArgsExposedToKDE ||
		preview.StateRootPathExposed ||
		preview.BackendDetailsExposed ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted {
		t.Fatalf("unexpected GUI card controlled launch action preview: %#v", preview)
	}
}

func TestPreviewKDEControlledLaunchActionAcceptsBuiltinGUIWithoutManagedCopy(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionGUICard(record)
	card.OwnerManagedCopyVerified = false

	preview, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:        stateRoot,
		KDECenterGUICard: &card,
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchAction returned error for builtin GUI evidence without managed copy: %v", err)
	}

	if preview.ApplicationID != "org.xnix.apps.messagebox" ||
		preview.ApplicationName != "Xnix MessageBox" ||
		preview.DesktopCallableRoute != "kde-dbus-runtime-status-action" ||
		!preview.DesktopEvidenceHandleForwarded ||
		preview.OwnerServiceArgsExposedToKDE ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled {
		t.Fatalf("unexpected builtin GUI action preview: %#v", preview)
	}
}

func TestPreviewKDEControlledLaunchActionRejectsUnsafeKDEGUICardRoute(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionGUICard(record)
	card.KDEForwardedArguments = []string{record.EvidenceRelativePath, "--state-root", stateRoot}
	card.OwnerServiceArgsExposedToKDE = true

	_, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:        stateRoot,
		KDECenterGUICard: &card,
	})
	if err == nil {
		t.Fatalf("unsafe GUI card route must be rejected")
	}
}

func TestPreviewKDEControlledLaunchActionSelectsKDECenterPageGUICard(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionGUICard(record)
	page := kdeControlledLaunchActionCenterPage(card)

	preview, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:          stateRoot,
		KDECenterPage:      &page,
		KDECenterPageAppID: "org.xnix.apps.messagebox",
	})
	if err != nil {
		t.Fatalf("PreviewKDEControlledLaunchAction returned error: %v", err)
	}

	if preview.ApplicationID != "org.xnix.apps.messagebox" ||
		preview.EvidenceRelativePath != record.EvidenceRelativePath ||
		!sameRuntimeStatusLaunchOwnerFixtureArgs(preview.KDEForwardedArguments, []string{record.EvidenceRelativePath}) ||
		preview.OwnerServiceArgsExposedToKDE ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected center page controlled launch action preview: %#v", preview)
	}
}

func TestPreviewKDEControlledLaunchActionRejectsUnsafeKDECenterPage(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionGUICard(record)
	page := kdeControlledLaunchActionCenterPage(card)
	page.LaunchEnabled = true

	_, err := PreviewKDEControlledLaunchAction(KDEControlledLaunchActionRequest{
		StateRoot:          stateRoot,
		KDECenterPage:      &page,
		KDECenterPageAppID: "org.xnix.apps.messagebox",
	})
	if err == nil {
		t.Fatalf("unsafe KDE center page must be rejected")
	}
}

func recordKDEControlledLaunchActionEvidence(t *testing.T, stateRoot string) KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

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
	return record
}

func recordKDEControlledLaunchActionGUIEvidence(t *testing.T, stateRoot string) KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

	appVersion := currentProjectVersion(t)
	sessionID := KnownAppControlledExecutionSessionID("org.xnix.apps.messagebox", appVersion)
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("org.xnix.apps.messagebox", appVersion)
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("org.xnix.apps.messagebox", appVersion, sessionID)
	projection, err := ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "org.xnix.apps.messagebox",
		DisplayName:                            "Xnix MessageBox",
		AppVersion:                             appVersion,
		RequestType:                            winapp.KnownDispatchSmokeRequestType,
		Status:                                 "passed",
		EvidenceSource:                         "wine-guest-gui-smoke",
		GuestBoundary:                          winapp.KnownDispatchGuestBoundary,
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
		ControlledSessionWindowObserved:        true,
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
	return record
}

func kdeControlledLaunchActionGUICard(record KnownAppKDERuntimeStatusLaunchEvidenceRecord) KDECenterPageKnownAppMatrixCard {
	return KDECenterPageKnownAppMatrixCard{
		AppID:                                "org.xnix.apps.messagebox",
		DisplayName:                          "Xnix MessageBox",
		AppVersion:                           record.AppVersion,
		EvidenceKind:                         "known-application-gui-smoke",
		EvidenceSource:                       "wine-guest-gui-smoke",
		SmokeStatus:                          "passed",
		XWindowObserved:                      true,
		WindowObserved:                       true,
		CompatibilityState:                   "owner-controlled-gui-qemu-wine-verified",
		CenterCardState:                      "validated-owner-controlled-gui-runtime-run",
		PrimaryActionID:                      KnownAppKDERuntimeStatusLaunchAction,
		PrimaryActionLabel:                   "Show Runtime-controlled launch",
		PrimaryActionKind:                    "runtime-status",
		PrimaryActionEnabled:                 true,
		DesktopCallableRoute:                 "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:         "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:         KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopDBusMethod:                    "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		DesktopEvidenceHandleForwarded:       true,
		KDEForwardedArgumentKind:             "evidence-relative-path",
		KDEForwardedArguments:                []string{record.EvidenceRelativePath},
		OwnerServiceArgsExposedToKDE:         false,
		ExecutionEvidenceRecorded:            true,
		StagedLauncherVerified:               true,
		OwnerControlledRuntimeLaunchVerified: true,
		OwnerManagedCopyVerified:             true,
		OwnerServiceCallReady:                true,
		OwnerEvidenceHandoffReady:            true,
		OwnerEvidenceRelativePath:            record.EvidenceRelativePath,
		RuntimeDispatchVerified:              true,
		LaunchAuthorizationRequired:          true,
		DesktopLaunchEnabled:                 false,
		BackendLaunchEnabled:                 false,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		HostRootModified:                     false,
		BackendDetailsExposed:                false,
		RawArtifactPathExposed:               false,
		Summary:                              "Xnix MessageBox GUI card can forward only a safe Runtime evidence handle.",
	}
}

func kdeControlledLaunchActionCenterPage(card KDECenterPageKnownAppMatrixCard) KDECenterPagePreview {
	return KDECenterPagePreview{
		SchemaVersion:                           "xnix.runtime.kde_center_page.v1",
		RequestType:                             "kde-center-page-preview",
		PageType:                                "compatibility-center-application-page",
		Desktop:                                 "KDE Plasma",
		RuntimeMethod:                           "GetKDECenterPage",
		ReadMethod:                              "GetKDECenterPagePreview",
		ApplicationID:                           card.AppID,
		ApplicationName:                         card.DisplayName,
		KnownAppGUIEvidenceCount:                1,
		KnownAppOwnerControlledGUIEvidenceCount: 1,
		KnownAppOwnerManagedCopyVerifiedCount:   1,
		KnownAppGUIEvidenceCards:                []KDECenterPageKnownAppMatrixCard{card},
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		UserVisible:                             true,
		PagePreviewCreated:                      true,
		LaunchEnabled:                           false,
		ExecutionStarted:                        false,
		BackendProcessStarted:                   false,
		HostRootModified:                        false,
		BackendDetailsExposed:                   false,
		DesktopSafeSummary:                      "KDE center page carries a safe owner-controlled GUI card.",
	}
}
