package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestRecordKnownAppControlledExecutionSessionBlocksUntilHandoffReady(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	record, err := RecordKnownAppControlledExecutionSession(KnownAppControlledExecutionSessionRecordRequest{
		AppID:         "7zr",
		StateRoot:     stateRoot,
		ReceiptID:     receipt.ReceiptID,
		CacheRoot:     t.TempDir(),
		GuestBoundary: "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("RecordKnownAppControlledExecutionSession returned error: %v", err)
	}
	if record.SchemaVersion != KnownAppControlledExecutionSessionRecordSchemaVersion ||
		record.RequestType != KnownAppControlledExecutionSessionRecordRequestType ||
		record.Source != KnownAppControlledExecutionSessionRequestType+"+execution-ledger+execution-session-record" ||
		record.AppID != "7zr" ||
		record.ReceiptAccepted != true ||
		record.GuestBoundaryAccepted != true ||
		record.ControlledDispatchRequestCreated ||
		record.ExecutionSessionHandoffCreated ||
		record.SessionHandoffReady ||
		record.RecordState != "blocked" ||
		record.LedgerRecordWritten ||
		record.SessionRecordWritten ||
		record.TransactionRelativePath != "" ||
		record.SessionRelativePath != "" ||
		record.SessionSHA256 != "" ||
		record.SessionState != "not-recorded" ||
		record.CompatibilityCenterState != "not-recorded" ||
		record.RuntimeOwnedExecutionSession ||
		record.StateRootPathExposed ||
		record.TransactionPathExposed ||
		record.SessionPathExposed ||
		record.DispatchStarted ||
		record.ExecutionStarted ||
		record.BackendProcessStarted ||
		record.HostRootModified {
		t.Fatalf("unexpected blocked controlled execution session record: %#v", record)
	}
	if _, err := os.Stat(filepath.Join(stateRoot, "execution-ledger")); !os.IsNotExist(err) {
		t.Fatalf("blocked controlled execution session record must not create execution ledger: %v", err)
	}
	encoded, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session record exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("controlled execution session record exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppControlledExecutionSessionConsumptionReadsDigestVerifiedRecord(t *testing.T) {
	stateRoot := t.TempDir()
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
	session, err := ledger.RecordSession(sessionID)
	if err != nil {
		t.Fatalf("RecordSession returned error: %v", err)
	}

	preview, err := PreviewKnownAppControlledExecutionSessionConsumption(KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppControlledExecutionSessionConsumption returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppControlledExecutionSessionConsumeSchemaVersion ||
		preview.RequestType != KnownAppControlledExecutionSessionConsumeRequestType ||
		preview.Source != KnownAppControlledExecutionSessionRecordRequestType+"+execution-session-fanout-evidence" ||
		preview.AppID != "7zr" ||
		preview.AppVersion != "26.02" ||
		preview.ExecutionSessionID != sessionID ||
		!preview.RecordConsumed ||
		!preview.LedgerRecordConsumed ||
		!preview.SessionRecordConsumed ||
		!preview.SessionDigestVerified ||
		preview.SessionRelativePath != session.RelativePath ||
		preview.SessionSHA256 != session.SHA256 ||
		preview.SessionState != "blocked" ||
		preview.TaskManagerState != "blocked" ||
		preview.KWinState != "blocked" ||
		preview.TrayState != "blocked" ||
		preview.CompatibilityCenterState != "waiting-for-runtime-gates" ||
		preview.FanOutRequestType != "execution-session-fanout-evidence" ||
		preview.SurfaceCount != 4 ||
		!preview.RuntimeOwnerConsumable ||
		!preview.KDEReadModelConsumable ||
		!preview.SafeForKDE ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.StateRootPathExposed ||
		preview.TransactionPathExposed ||
		preview.SessionPathExposed ||
		preview.LiveStateObserved ||
		preview.SessionRegistered ||
		preview.WindowObserved ||
		preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled ||
		preview.DispatchStarted ||
		preview.ExecutionStarted ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified {
		t.Fatalf("unexpected controlled execution session consume preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled execution session consume preview exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("controlled execution session consume preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppControlledExecutionSessionConsumptionRejectsMismatchedSessionID(t *testing.T) {
	_, err := PreviewKnownAppControlledExecutionSessionConsumption(KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     "7zr",
		StateRoot: t.TempDir(),
		SessionID: "known-app-controlled-execution-session-7zr-99.99",
	})
	if err == nil {
		t.Fatal("expected mismatched controlled execution session id to be rejected")
	}
}
