package appidentity

import (
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestPreviewKnownAppSessionGatedLaunchReviewConsumesSessionBeforeReviewRoute(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)

	preview, err := PreviewKnownAppSessionGatedLaunchReview(KnownAppSessionGatedLaunchReviewRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppSessionGatedLaunchReview returned error: %v", err)
	}

	if preview.SchemaVersion != KnownAppSessionGatedLaunchReviewSchemaVersion ||
		preview.RequestType != KnownAppSessionGatedLaunchReviewRequestType ||
		preview.ReviewType != "known-app-session-gated-launch-review" ||
		preview.Source != KnownAppControlledExecutionSessionConsumeRequestType+"+kde-center-page-session-gate-card" ||
		preview.RuntimeMethod != "PreviewKnownAppSessionGatedLaunchReview" ||
		preview.ReadMethod != "GetKnownAppSessionGatedLaunchReview" ||
		preview.AppID != "7zr" ||
		preview.DisplayName != "7-Zip standalone console executable" ||
		preview.AppVersion != "26.02" ||
		preview.ActionID != "review-session-gated-dispatch" ||
		preview.ActionKind != "session-gate-review" ||
		preview.Decision != "approved" ||
		!preview.DecisionAccepted ||
		preview.ReviewRouteID != "runtime-owned-session-gated-launch-review" ||
		!preview.ReviewRouteCreated ||
		preview.ReviewRouteRequestType != KnownAppSessionGatedLaunchReviewRequestType ||
		preview.ReviewRouteRuntimeMethod != "PreviewKnownAppSessionGatedLaunchReview" ||
		preview.ReviewRouteReadMethod != "GetKnownAppSessionGatedLaunchReview" ||
		!preview.ReadBeforeWriteRequired ||
		!preview.RuntimeReceiptRequired ||
		preview.ExecutionSessionID != sessionID ||
		!preview.SessionRecordConsumed ||
		!preview.SessionDigestVerified ||
		preview.SessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		preview.SessionSHA256 == "" ||
		preview.SessionState != "blocked" ||
		preview.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		!preview.RuntimeOwnerConsumable ||
		!preview.KDEReadModelConsumable ||
		!preview.SafeForKDE ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserReviewCaptured ||
		preview.UserDecisionAllowsLaunch ||
		preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed ||
		preview.LaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.ReviewReceiptRecorded ||
		preview.StateRootPathExposed ||
		preview.SessionPathExposed ||
		preview.RawArtifactPathExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected known app session-gated launch review preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("known app session-gated launch review exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("known app session-gated launch review exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppSessionGatedLaunchReviewRejectsInvalidAction(t *testing.T) {
	_, err := PreviewKnownAppSessionGatedLaunchReview(KnownAppSessionGatedLaunchReviewRequest{
		AppID:     "7zr",
		StateRoot: t.TempDir(),
		ActionID:  "run-directly",
		Decision:  "approved",
	})
	if err == nil {
		t.Fatal("expected invalid session-gated launch review action to be rejected")
	}
}

func writeKnownAppSessionGatedLaunchReviewFixture(t *testing.T, stateRoot string) string {
	t.Helper()
	sessionID := KnownAppControlledExecutionSessionID("7zr", "26.02")
	ledger, err := execution.NewLedger(stateRoot)
	if err != nil {
		t.Fatalf("NewLedger returned error: %v", err)
	}
	if _, err := ledger.Record(execution.Transaction{
		RequestID:      sessionID,
		ApplicationID:  "7zr",
		Profile:        "known-app-managed-guest",
		State:          execution.StateBlocked,
		ReviewDecision: execution.DecisionApproved,
		Gates: []execution.Gate{
			{ID: "controlled-execution-session-handoff", Status: execution.GatePass, Reason: "Runtime execution session handoff ready"},
			{ID: "runtime-write-gate", Status: execution.GateBlocked, Reason: "dispatch runner must consume the recorded session before execution"},
		},
		BlockedReasons:   []string{"runtime-write-gate: dispatch runner must consume the recorded session before execution"},
		LaunchAllowed:    false,
		LaunchEnabled:    false,
		BackendStarted:   false,
		HostRootModified: false,
		NetworkRequired:  false,
		Summary:          "Runtime recorded a known application execution session handoff for later consumption.",
	}); err != nil {
		t.Fatalf("Record returned error: %v", err)
	}
	if _, err := ledger.RecordSession(sessionID); err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}
	return sessionID
}
