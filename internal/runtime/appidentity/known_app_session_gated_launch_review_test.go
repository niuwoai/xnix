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

func TestPreviewKnownAppSessionGatedLaunchReviewGateAcceptsApprovedReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	receipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}

	gate, err := PreviewKnownAppSessionGatedLaunchReviewGate(KnownAppSessionGatedLaunchReviewGateRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ReceiptID: receipt.ReceiptID,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppSessionGatedLaunchReviewGate returned error: %v", err)
	}
	if gate.SchemaVersion != KnownAppSessionGatedLaunchReviewGateSchemaVersion ||
		gate.RequestType != KnownAppSessionGatedLaunchReviewGateRequestType ||
		gate.Source != KnownAppSessionGatedLaunchReviewReceiptRequestType+"+runtime-review-receipt-gate" ||
		gate.RuntimeMethod != "PreviewKnownAppSessionGatedLaunchReviewGate" ||
		gate.ReadMethod != "GetKnownAppSessionGatedLaunchReviewGate" ||
		gate.AppID != "7zr" ||
		gate.DisplayName != "7-Zip standalone console executable" ||
		gate.AppVersion != "26.02" ||
		gate.ActionID != "review-session-gated-dispatch" ||
		gate.ActionKind != "session-gate-review" ||
		gate.ExecutionSessionID != sessionID ||
		!gate.SessionRecordConsumed ||
		!gate.SessionDigestVerified ||
		gate.SessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		gate.SessionSHA256 == "" ||
		!gate.ReadBeforeWriteRevalidated ||
		gate.ReceiptID != receipt.ReceiptID ||
		gate.ReceiptRelativePath != receipt.ReceiptRelativePath ||
		gate.ReceiptSHA256 != receipt.ReceiptSHA256 ||
		gate.ReceiptLookupState != "accepted-receipt" ||
		gate.ReceiptDecision != "approved" ||
		!gate.ReviewReceiptConsumed ||
		!gate.ReviewReceiptAccepted ||
		gate.ReviewGateState != "review-receipt-accepted-dispatch-still-gated" ||
		!gate.ReviewGateReady ||
		!gate.ControlledDispatchGateReady ||
		!gate.DispatchStateAdvanceReady ||
		gate.DispatchStateAdvanced ||
		!gate.RuntimeOwnerConsumable ||
		!gate.KDEReadModelConsumable ||
		!gate.SafeForKDE ||
		!gate.RuntimeOwned ||
		!gate.GoRuntimeBacked ||
		gate.KDEPolicyOwner ||
		!gate.UserDecisionAllowsLaunch ||
		gate.RuntimeLaunchApproval ||
		gate.LaunchAllowed ||
		gate.LaunchEnabled ||
		gate.DesktopLaunchEnabled ||
		gate.ExecutionStarted ||
		gate.BackendLaunchEnabled ||
		gate.BackendProcessStarted ||
		gate.ControlledDispatchRequestMade ||
		gate.PermissionGrantCreated ||
		gate.StateRootPathExposed ||
		gate.SessionPathExposed ||
		gate.ReceiptPathExposed ||
		gate.RawArtifactPathExposed ||
		gate.BackendDetailsExposed ||
		gate.HostRootModified ||
		gate.NetworkRequired ||
		gate.PrivilegedContainerRequired ||
		gate.DockerSocketMounted ||
		gate.BroadHostMountRequired {
		t.Fatalf("unexpected known app session-gated launch review gate preview: %#v", gate)
	}
	encoded, err := json.Marshal(gate)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("known app session-gated launch review gate exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("known app session-gated launch review gate exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppSessionGatedLaunchReviewGateRejectsDeferredReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	receipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "deferred",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	gate, err := PreviewKnownAppSessionGatedLaunchReviewGate(KnownAppSessionGatedLaunchReviewGateRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ReceiptID: receipt.ReceiptID,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppSessionGatedLaunchReviewGate returned error: %v", err)
	}
	if gate.ReceiptLookupState != "rejected-receipt" ||
		gate.ReceiptRejectedReason != "session-gated launch review receipt decision is not approved" ||
		gate.ReceiptDecision != "deferred" ||
		!gate.ReviewReceiptConsumed ||
		gate.ReviewReceiptAccepted ||
		gate.ReviewGateReady ||
		gate.ControlledDispatchGateReady ||
		gate.DispatchStateAdvanceReady ||
		gate.UserDecisionAllowsLaunch ||
		gate.RuntimeLaunchApproval ||
		gate.LaunchAllowed ||
		gate.DesktopLaunchEnabled ||
		gate.ExecutionStarted ||
		gate.ControlledDispatchRequestMade ||
		gate.HostRootModified {
		t.Fatalf("unexpected deferred receipt gate preview: %#v", gate)
	}
}

func TestPreviewKnownAppSessionGatedControlledDispatchRequestRequiresAcceptedReviewGate(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	reviewReceipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "approved",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	launchReceipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	preview, err := PreviewKnownAppSessionGatedControlledDispatchRequest(KnownAppSessionGatedControlledDispatchRequest{
		AppID:           "7zr",
		StateRoot:       stateRoot,
		SessionID:       sessionID,
		ReviewReceiptID: reviewReceipt.ReceiptID,
		LaunchReceiptID: launchReceipt.ReceiptID,
		CacheRoot:       t.TempDir(),
		GuestBoundary:   "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppSessionGatedControlledDispatchRequest returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppSessionGatedControlledDispatchSchemaVersion ||
		preview.RequestType != KnownAppSessionGatedControlledDispatchRequestType ||
		preview.Source != KnownAppSessionGatedLaunchReviewGateRequestType+"+"+KnownAppControlledDispatchRequestType ||
		preview.RuntimeMethod != "PreviewKnownAppSessionGatedControlledDispatchRequest" ||
		preview.ReadMethod != "GetKnownAppSessionGatedControlledDispatchRequest" ||
		preview.AppID != "7zr" ||
		preview.ExecutionSessionID != sessionID ||
		!preview.SessionRecordConsumed ||
		!preview.SessionDigestVerified ||
		preview.SessionRelativePath != "execution-ledger/sessions/"+sessionID+".json" ||
		preview.SessionSHA256 == "" ||
		preview.ReviewReceiptID != reviewReceipt.ReceiptID ||
		preview.ReviewReceiptRelativePath != reviewReceipt.ReceiptRelativePath ||
		preview.ReviewReceiptSHA256 != reviewReceipt.ReceiptSHA256 ||
		!preview.ReviewReceiptConsumed ||
		!preview.ReviewReceiptAccepted ||
		preview.ReviewGateState != "review-receipt-accepted-dispatch-still-gated" ||
		!preview.ReviewGateReady ||
		!preview.DispatchStateAdvanceReady ||
		preview.LaunchAuthorizationReceiptID != launchReceipt.ReceiptID ||
		preview.LaunchGateState != "dispatch-preparation-required" ||
		!preview.LaunchGateReceiptAccepted ||
		!preview.LaunchGateGuestBoundaryAccepted ||
		preview.ControlledDispatchGateReady ||
		preview.ControlledDispatchRequestCreated ||
		preview.ControlledDispatchRequestState != "blocked" ||
		preview.RuntimeOwnedDispatchRequest ||
		preview.ArtifactVerified ||
		preview.DispatchReady ||
		preview.DispatchAllowed ||
		preview.DispatchStarted ||
		preview.ExecutionStarted ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.RequestObjectsCreated ||
		preview.PermissionGrantCreated ||
		preview.ReviewReceiptPathExposed ||
		preview.ReceiptPathExposed ||
		preview.StateRootPathExposed ||
		preview.SessionPathExposed ||
		preview.RawArtifactPathExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unexpected known app session-gated controlled dispatch preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("known app session-gated controlled dispatch exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("known app session-gated controlled dispatch exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppSessionGatedControlledDispatchRequestRejectsDeferredReviewReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	sessionID := writeKnownAppSessionGatedLaunchReviewFixture(t, stateRoot)
	reviewReceipt, err := RecordKnownAppSessionGatedLaunchReviewReceipt(KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		SessionID: sessionID,
		ActionID:  KnownAppSessionGatedLaunchReviewAction,
		Decision:  "deferred",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppSessionGatedLaunchReviewReceipt returned error: %v", err)
	}
	preview, err := PreviewKnownAppSessionGatedControlledDispatchRequest(KnownAppSessionGatedControlledDispatchRequest{
		AppID:           "7zr",
		StateRoot:       stateRoot,
		SessionID:       sessionID,
		ReviewReceiptID: reviewReceipt.ReceiptID,
		LaunchReceiptID: KnownAppLaunchAuthorizationReceiptID("7zr", "26.02"),
		CacheRoot:       t.TempDir(),
		GuestBoundary:   "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppSessionGatedControlledDispatchRequest returned error: %v", err)
	}
	if preview.ReviewReceiptAccepted ||
		preview.ReviewGateReady ||
		preview.DispatchStateAdvanceReady ||
		preview.ControlledDispatchGateReady ||
		preview.ControlledDispatchRequestCreated ||
		preview.RequestObjectsCreated ||
		preview.DispatchReady ||
		preview.DispatchAllowed ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.ControlledDispatchRequestState != "blocked" ||
		preview.LaunchGateBlockedReason != "session-gated launch review receipt decision is not approved" {
		t.Fatalf("unexpected rejected review receipt dispatch preview: %#v", preview)
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
