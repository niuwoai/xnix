package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewKnownAppKDERuntimeStatusLaunchActionTriggerConsumesHandoff(t *testing.T) {
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
	trigger, err := PreviewKnownAppKDERuntimeStatusLaunchActionTrigger(KnownAppKDERuntimeStatusLaunchActionTriggerRequest{
		StateRoot:            stateRoot,
		EvidenceRelativePath: record.EvidenceRelativePath,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppKDERuntimeStatusLaunchActionTrigger returned error: %v", err)
	}
	if trigger.SchemaVersion != KnownAppKDERuntimeStatusLaunchActionTriggerSchemaVersion ||
		trigger.RequestType != KnownAppKDERuntimeStatusLaunchActionTriggerRequestType ||
		trigger.RuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		trigger.ReadMethod != "GetKnownAppKDERuntimeStatusLaunchActionTrigger" ||
		trigger.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		trigger.ActionKind != "runtime-status" ||
		trigger.TriggerState != "runtime-launch-request-assembled" ||
		trigger.EvidenceReadState != "consumed" ||
		!trigger.EvidenceHandoffConsumed ||
		trigger.EvidenceRelativePath != record.EvidenceRelativePath ||
		trigger.EvidenceSHA256 != record.EvidenceSHA256 ||
		!trigger.EvidenceDigestVerified ||
		trigger.LaunchRequestType != KnownAppKDERuntimeStatusLaunchRequestType ||
		trigger.LaunchRequestRuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		trigger.LaunchRequestReadMethod != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		!trigger.LaunchRequestCreated ||
		!trigger.ManagedLauncherArgvReady ||
		strings.Join(trigger.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --receipt-id "+launchReceiptID+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		trigger.RequiredOpaqueIDCount != 3 ||
		trigger.CollectedOpaqueIDCount != 3 ||
		!trigger.RuntimeOwnedTrigger ||
		!trigger.RuntimeOwnedRequest ||
		!trigger.RuntimeOwnedLaunch ||
		!trigger.RuntimeOwnedDispatch ||
		!trigger.KDEPresentationOnly ||
		!trigger.KDEActionForwarded ||
		!trigger.StateRootRequired ||
		!trigger.StateRootSuppliedByRuntime ||
		trigger.KDEStateRootAccess ||
		trigger.DirectLaunchEnabled ||
		trigger.DesktopLaunchEnabled ||
		trigger.BackendLaunchEnabled ||
		trigger.ExecutionStarted ||
		trigger.BackendProcessStarted ||
		trigger.RequestObjectsCreated ||
		trigger.PermissionGrantCreated ||
		trigger.StateRootPathExposed ||
		trigger.EvidencePathExposed ||
		trigger.ReceiptPathExposed ||
		trigger.SessionPathExposed ||
		trigger.RawArtifactPathExposed ||
		trigger.RawCommandExposed ||
		trigger.RawLauncherOutputExposed ||
		trigger.BackendDetailsExposed ||
		trigger.HostRootModified ||
		trigger.PrivilegedContainerRequired ||
		trigger.DockerSocketMounted ||
		trigger.BroadHostMountRequired {
		t.Fatalf("unexpected Runtime-status launch action trigger: %#v", trigger)
	}
	if trigger.LaunchRequest.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		!trigger.LaunchRequest.LaunchRequestCreated ||
		!trigger.LaunchRequest.StateRootSuppliedByRuntime ||
		trigger.LaunchRequest.KDEStateRootAccess ||
		trigger.LaunchRequest.ExecutionStarted {
		t.Fatalf("unexpected nested launch request: %#v", trigger.LaunchRequest)
	}
	encoded, err := json.Marshal(trigger)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) {
		t.Fatalf("action trigger exposed state root: %s", string(encoded))
	}
}
