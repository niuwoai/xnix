package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func TestRecordKnownAppSessionGatedLaunchReviewReceiptConsumesSessionBeforeWrite(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)

	record, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}

	expectedReceiptID := KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	expectedRelativePath := KnownAppSessionGatedLaunchReviewReceiptRelativePath(expectedReceiptID)
	if record.SchemaVersion != KnownAppSessionGatedLaunchReviewReceiptSchemaVersion ||
		record.RequestType != KnownAppSessionGatedLaunchReviewReceiptRequestType ||
		record.ReviewType != "known-app-session-gated-launch-review-receipt" ||
		record.Source != KnownAppSessionGatedLaunchReviewRequestType+"+runtime-review-receipt-store" ||
		record.RuntimeMethod != "RecordKnownAppSessionGatedLaunchReviewReceipt" ||
		record.ReadMethod != "GetKnownAppSessionGatedLaunchReviewReceipt" ||
		record.AppID != "7zr" ||
		record.DisplayName != "7-Zip standalone console executable" ||
		record.AppVersion != "26.02" ||
		record.ActionID != "review-session-gated-dispatch" ||
		record.ActionKind != "session-gate-review" ||
		record.Decision != "approved" ||
		!record.DecisionAccepted ||
		!record.ReadBeforeWriteConsumed ||
		record.ReviewRouteID != "runtime-owned-session-gated-launch-review" ||
		!record.ReviewRouteConsumed ||
		record.ReviewRouteRequestType != KnownAppSessionGatedLaunchReviewRequestType ||
		record.ReviewRouteRuntimeMethod != "PreviewKnownAppSessionGatedLaunchReview" ||
		record.ReviewRouteReadMethod != "GetKnownAppSessionGatedLaunchReview" ||
		!record.RuntimeReceiptRequired ||
		record.ExecutionSessionID != sessionID ||
		!record.SessionRecordConsumed ||
		!record.SessionDigestVerified ||
		record.SessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		record.SessionSHA256 == "" ||
		record.SessionState != "blocked" ||
		record.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		record.ReceiptID != expectedReceiptID ||
		record.ReceiptRelativePath != expectedRelativePath ||
		len(record.ReceiptSHA256) != 64 ||
		record.ReceiptState != "recorded-dispatch-still-gated" ||
		!record.ReviewReceiptRecorded ||
		!record.RuntimeOwnerConsumable ||
		!record.KDEReadModelConsumable ||
		!record.SafeForKDE ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		!record.UserReviewCaptured ||
		record.UserDecisionAllowsLaunch ||
		record.RuntimeLaunchApproval ||
		record.LaunchAllowed ||
		record.LaunchEnabled ||
		record.DesktopLaunchEnabled ||
		record.ExecutionStarted ||
		record.BackendLaunchEnabled ||
		record.BackendProcessStarted ||
		record.RequestObjectsCreated ||
		record.PermissionGrantCreated ||
		record.StateRootPathExposed ||
		record.SessionPathExposed ||
		record.ReceiptPathExposed ||
		record.RawArtifactPathExposed ||
		record.BackendDetailsExposed ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired {
		t.Fatalf("unexpected known app session-gated launch review receipt record: %#v", record)
	}
	receiptPath := filepath.Join(stateRoot, filepath.FromSlash(record.ReceiptRelativePath))
	data, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	if got := sha256Hex(string(data)); got != record.ReceiptSHA256 {
		t.Fatalf("receipt digest mismatch: got %s want %s", got, record.ReceiptSHA256)
	}
	var receipt map[string]any
	if err := json.Unmarshal(data, &receipt); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}
	if receipt["schema_version"] != KnownAppSessionGatedLaunchReviewReceiptSchemaVersion ||
		receipt["receipt_id"] != expectedReceiptID ||
		receipt["decision"] != "approved" ||
		receipt["session_digest_verified"] != true ||
		receipt["runtime_launch_approval"] != false ||
		receipt["launch_allowed"] != false ||
		receipt["desktop_launch_enabled"] != false ||
		receipt["backend_launch_enabled"] != false ||
		receipt["execution_started"] != false ||
		receipt["host_root_modified"] != false {
		t.Fatalf("unexpected persisted session-gated launch review receipt: %#v", receipt)
	}
	encoded, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("known app session-gated launch review receipt exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("known app session-gated launch review receipt exposes forbidden term %q: %s", forbidden, text)
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
