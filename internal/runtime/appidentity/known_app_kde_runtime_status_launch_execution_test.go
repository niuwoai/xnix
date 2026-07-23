package appidentity

import (
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/winapp"
)

func TestPrepareKnownAppKDERuntimeStatusLaunchExecutionRevalidatesRuntimeOwnedState(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppRuntimeStatusLaunchExecutionFixture(t, stateRoot)
	launchReceiptID := KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	plan, err := PrepareKnownAppKDERuntimeStatusLaunchExecution(KnownAppKDERuntimeStatusLaunchExecutionRequest{
		AppID:                        "7zr",
		StateRoot:                    stateRoot,
		LaunchAuthorizationReceiptID: launchReceiptID,
		SessionGatedReviewReceiptID:  reviewReceiptID,
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err != nil {
		t.Fatalf("PrepareKnownAppKDERuntimeStatusLaunchExecution returned error: %v", err)
	}
	if plan.SchemaVersion != KnownAppKDERuntimeStatusLaunchExecutionSchemaVersion ||
		plan.RequestType != KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		plan.RuntimeMethod != "PrepareKnownAppKDERuntimeStatusLaunchExecution" ||
		plan.ExecutionMethod != "RunKnownAppKDERuntimeStatusLaunchExecution" ||
		plan.RequestPreviewType != KnownAppKDERuntimeStatusLaunchRequestType ||
		plan.ActionID != KnownAppKDERuntimeStatusLaunchAction ||
		plan.ActionKind != "runtime-status" ||
		plan.CenterCardState != "validated-post-review-dispatch" ||
		plan.PostReviewDispatchState != "created-after-session-gated-review" ||
		plan.LaunchAuthorizationReceiptID != launchReceiptID ||
		plan.SessionGatedReviewReceiptID != reviewReceiptID ||
		plan.ControlledExecutionSessionID != sessionID ||
		!plan.LaunchReceiptRevalidated ||
		!plan.GuestBoundaryRevalidated ||
		!plan.ReviewReceiptRevalidated ||
		!plan.ControlledSessionRevalidated ||
		!plan.ControlledSessionDigestVerified ||
		!plan.RuntimeManagedLauncherArgvReady ||
		strings.Join(plan.RuntimeManagedLauncherArgv, " ") != "xnix-compat-launch --app 7zr --guest-boundary managed-known-app-guest-smoke --state-root <runtime-owned-state-root> --receipt-id "+launchReceiptID+" --review-receipt-id "+reviewReceiptID+" --session-id "+sessionID ||
		plan.RequiredOpaqueIDCount != 3 ||
		plan.CollectedOpaqueIDCount != 3 ||
		!plan.StateRootRequired ||
		!plan.StateRootAccepted ||
		!plan.StateRootInjectedByRuntime ||
		!plan.StateRootSuppliedByRuntime ||
		plan.KDEStateRootAccess ||
		!plan.ManagedLauncherInvocationReady ||
		!plan.ExistingManagedLauncherPathUsed ||
		!plan.DelegatesArtifactGateToLauncher ||
		!plan.RuntimeOwnedRequest ||
		!plan.RuntimeOwnedLaunch ||
		!plan.RuntimeOwnedDispatch ||
		!plan.KDEPresentationOnly ||
		!plan.KDEActionForwarded ||
		plan.DirectLaunchEnabled ||
		plan.DesktopLaunchEnabled ||
		plan.BackendLaunchEnabled ||
		plan.ExecutionStarted ||
		plan.BackendProcessStarted ||
		plan.RequestObjectsCreated ||
		plan.PermissionGrantCreated ||
		plan.StateRootPathExposed ||
		plan.ReceiptPathExposed ||
		plan.SessionPathExposed ||
		plan.ManagedLauncherPathExposed ||
		plan.RawArtifactPathExposed ||
		plan.RawCommandExposed ||
		plan.RawLauncherOutputExposed ||
		plan.BackendDetailsExposed ||
		plan.HostRootModified ||
		plan.PrivilegedContainerRequired ||
		plan.DockerSocketMounted ||
		plan.BroadHostMountRequired {
		t.Fatalf("unexpected Runtime-status launch execution plan: %#v", plan)
	}
	argv, err := KnownAppKDERuntimeStatusLaunchExecutionArgv(plan, stateRoot, "", []string{"--timeout", "1s"})
	if err != nil {
		t.Fatalf("KnownAppKDERuntimeStatusLaunchExecutionArgv returned error: %v", err)
	}
	expected := []string{
		"--app", "7zr",
		"--cache-root", winapp.DefaultKnownAppCacheRoot,
		"--guest-boundary", winapp.KnownDispatchGuestBoundary,
		"--state-root", stateRoot,
		"--receipt-id", launchReceiptID,
		"--review-receipt-id", reviewReceiptID,
		"--session-id", sessionID,
		"--timeout", "1s",
	}
	if strings.Join(argv, "\x00") != strings.Join(expected, "\x00") {
		t.Fatalf("unexpected execution argv: %#v", argv)
	}
	encoded, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{strings.ToLower(stateRoot), ".exe", "program files", "qemu-system", "proton", "wine ", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Runtime-status launch execution plan exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPrepareKnownAppKDERuntimeStatusLaunchExecutionRejectsMissingReviewReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	_, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	_, err = PrepareKnownAppKDERuntimeStatusLaunchExecution(KnownAppKDERuntimeStatusLaunchExecutionRequest{
		AppID:                        "7zr",
		StateRoot:                    stateRoot,
		LaunchAuthorizationReceiptID: KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		SessionGatedReviewReceiptID:  KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID),
		ControlledExecutionSessionID: sessionID,
		CenterCardState:              "validated-post-review-dispatch",
		PrimaryActionID:              KnownAppKDERuntimeStatusLaunchAction,
		PostReviewDispatchState:      "created-after-session-gated-review",
	})
	if err == nil || !strings.Contains(err.Error(), "accepted session-gated review receipt") {
		t.Fatalf("expected missing review receipt to be rejected, got %v", err)
	}
}

func writeKnownAppRuntimeStatusLaunchExecutionFixture(t *testing.T, stateRoot string) string {
	t.Helper()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	if _, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	}); err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	if _, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	}); err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	return sessionID
}
