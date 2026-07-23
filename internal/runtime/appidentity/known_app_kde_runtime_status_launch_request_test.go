package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewKnownAppKDERuntimeStatusLaunchRequestCollectsOpaqueIDs(t *testing.T) {
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	preview, err := PreviewKnownAppKDERuntimeStatusLaunchRequest(KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        "7zr",
		LaunchAuthorizationReceiptID: launchReceiptID,
		SessionGatedReviewReceiptID:  reviewReceiptID,
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppKDERuntimeStatusLaunchRequest returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppKDERuntimeStatusLaunchRequestSchemaVersion ||
		preview.RequestType != KnownAppKDERuntimeStatusLaunchRequestType ||
		preview.RuntimeMethod != "PreviewKnownAppKDERuntimeStatusLaunchRequest" ||
		preview.ReadMethod != "GetKnownAppKDERuntimeStatusLaunchRequest" ||
		preview.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		preview.ActionKind != "runtime-status" ||
		preview.CenterCardState != "validated-post-review-dispatch" ||
		preview.PostReviewDispatchState != "created-after-session-gated-review" ||
		preview.LaunchAuthorizationReceiptID != launchReceiptID ||
		preview.SessionGatedReviewReceiptID != reviewReceiptID ||
		preview.ControlledExecutionSessionID != sessionID ||
		!preview.ManagedLauncherArgvReady ||
		strings.Join(preview.ManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --receipt-id "+launchReceiptID+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		preview.RequiredOpaqueIDCount != 3 ||
		preview.CollectedOpaqueIDCount != 3 ||
		!preview.LaunchRequestCreated ||
		!preview.RuntimeOwnedRequest ||
		!preview.RuntimeOwnedLaunch ||
		!preview.RuntimeOwnedDispatch ||
		!preview.KDEPresentationOnly ||
		!preview.KDEActionForwarded ||
		!preview.StateRootRequired ||
		!preview.StateRootSuppliedByRuntime ||
		preview.KDEStateRootAccess ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.StateRootPathExposed ||
		preview.ReceiptPathExposed ||
		preview.SessionPathExposed ||
		preview.RawArtifactPathExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected KDE Runtime-status launch request preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE Runtime-status launch request exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppKDERuntimeStatusLaunchRequestRejectsIncompleteCard(t *testing.T) {
	_, err := PreviewKnownAppKDERuntimeStatusLaunchRequest(KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        "7zr",
		LaunchAuthorizationReceiptID: KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		ControlledExecutionSessionID: KnownAppControlledExecutionSessionID("7zr", "26.02"),
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err == nil || !strings.Contains(err.Error(), "safe opaque ids") {
		t.Fatalf("expected safe opaque id error, got %v", err)
	}
}

func TestPreviewKnownAppKDERuntimeStatusLaunchRequestRejectsKDEOwnedLaunchAction(t *testing.T) {
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	_, err := PreviewKnownAppKDERuntimeStatusLaunchRequest(KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        "7zr",
		LaunchAuthorizationReceiptID: KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		SessionGatedReviewReceiptID:  KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID),
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              "launch-directly-from-kde",
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err == nil || !strings.Contains(err.Error(), KnownAppKDERuntimeStatusLaunchAction) {
		t.Fatalf("expected Runtime-status action error, got %v", err)
	}
}
