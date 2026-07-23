package appidentity

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRecordKnownAppLaunchAuthorizationReceiptWritesOpaqueReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	preview, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:         "7zr",
		StateRoot:     stateRoot,
		Authorize:     KnownAppLaunchAuthorizationReceiptAction,
		RecordedAtUTC: time.Date(2026, 7, 24, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppLaunchAuthorizationReceiptSchemaVersion ||
		preview.RequestType != KnownAppLaunchAuthorizationReceiptRequestType ||
		preview.AppID != "7zr" ||
		preview.AppVersion != "26.02" ||
		preview.ReceiptID != "known-app-launch-authorization-7zr-26.02" ||
		preview.ReceiptState != "recorded" ||
		!preview.ReceiptWritten ||
		preview.ReceiptPathExposed ||
		preview.StateRootPathExposed ||
		!preview.ManagedStateRoot ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		!preview.LaunchAuthorizationRequired ||
		!preview.LaunchAuthorizationRecorded ||
		!preview.LaunchGateReady ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		!preview.StateRootModified ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.RawArtifactPathExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected receipt preview: %#v", preview)
	}
	receiptPath, err := KnownAppLaunchAuthorizationReceiptPath(stateRoot, preview.ReceiptID)
	if err != nil {
		t.Fatalf("KnownAppLaunchAuthorizationReceiptPath returned error: %v", err)
	}
	content, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	var receipt map[string]any
	if err := json.Unmarshal(content, &receipt); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}
	if receipt["receipt_id"] != preview.ReceiptID ||
		receipt["app_id"] != "7zr" ||
		receipt["authorized_action_id"] != KnownAppLaunchAuthorizationReceiptAction ||
		receipt["launch_authorization_state"] != "recorded" ||
		receipt["launch_gate_state"] != "receipt-recorded-launch-still-gated" ||
		receipt["backend_launch_enabled"] != false ||
		receipt["host_root_modified"] != false {
		t.Fatalf("unexpected receipt file: %#v", receipt)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("receipt preview exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("receipt preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestRecordKnownAppLaunchAuthorizationReceiptRequiresExplicitAuthorization(t *testing.T) {
	for _, request := range []KnownAppLaunchAuthorizationReceiptRequest{
		{AppID: "7zr", StateRoot: t.TempDir()},
		{AppID: "7zr", StateRoot: t.TempDir(), Authorize: "launch"},
		{AppID: "7zr", Authorize: KnownAppLaunchAuthorizationReceiptAction},
		{AppID: "missing", StateRoot: t.TempDir(), Authorize: KnownAppLaunchAuthorizationReceiptAction},
	} {
		if _, err := RecordKnownAppLaunchAuthorizationReceipt(request); err == nil {
			t.Fatalf("expected request to fail closed: %#v", request)
		}
	}
}

func TestPreviewKnownAppLaunchGateConsumesReceiptAndRequiresGuestBoundary(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	preview, err := PreviewKnownAppLaunchGate(KnownAppLaunchGateRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		ReceiptID: receipt.ReceiptID,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppLaunchGate returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppLaunchGateSchemaVersion ||
		preview.RequestType != KnownAppLaunchGateRequestType ||
		preview.ReceiptLookupState != "accepted-receipt" ||
		!preview.ReceiptAccepted ||
		!preview.GuestBoundaryRequired ||
		preview.GuestBoundaryAccepted ||
		preview.LaunchGateState != "guest-boundary-missing-fail-closed" ||
		preview.LaunchGateBlockedReason != "controlled managed guest boundary is required before dispatch" ||
		preview.ReceiptRejectedReason != "" ||
		preview.ControlledDispatchReady ||
		preview.DispatchRequestMaterialized ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.ReceiptPathExposed ||
		preview.StateRootPathExposed {
		t.Fatalf("unexpected missing-boundary gate preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("launch gate preview exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("launch gate preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppLaunchGateAcceptsBoundaryButKeepsDispatchPreparationGated(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	preview, err := PreviewKnownAppLaunchGate(KnownAppLaunchGateRequest{
		AppID:         "7zr",
		StateRoot:     stateRoot,
		ReceiptID:     receipt.ReceiptID,
		CacheRoot:     t.TempDir(),
		GuestBoundary: "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppLaunchGate returned error: %v", err)
	}
	if preview.ReceiptLookupState != "accepted-receipt" ||
		!preview.ReceiptAccepted ||
		!preview.GuestBoundaryAccepted ||
		preview.DispatchStatus != "dispatch-blocked" ||
		preview.DispatchReady ||
		preview.LaunchGateState != "dispatch-preparation-required" ||
		preview.LaunchGateBlockedReason != "managed artifact preparation is required before dispatch" ||
		preview.ReceiptRejectedReason != "" ||
		preview.ControlledDispatchReady ||
		preview.DispatchRequestMaterialized ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified {
		t.Fatalf("unexpected preparation-gated launch gate preview: %#v", preview)
	}
}

func TestPreviewKnownAppControlledDispatchRequestBlocksUntilArtifactVerified(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}

	preview, err := PreviewKnownAppControlledDispatchRequest(KnownAppControlledDispatchRequest{
		AppID:         "7zr",
		StateRoot:     stateRoot,
		ReceiptID:     receipt.ReceiptID,
		CacheRoot:     t.TempDir(),
		GuestBoundary: "managed-known-app-guest-smoke",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppControlledDispatchRequest returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppControlledDispatchSchemaVersion ||
		preview.RequestType != KnownAppControlledDispatchRequestType ||
		preview.ReceiptAccepted != true ||
		preview.GuestBoundaryAccepted != true ||
		preview.LaunchGateState != "dispatch-preparation-required" ||
		preview.ControlledDispatchReady ||
		preview.ControlledDispatchRequestCreated ||
		preview.ControlledDispatchRequestState != "blocked" ||
		preview.DispatchReady ||
		preview.DispatchAllowed ||
		preview.DispatchStarted ||
		preview.ExecutionStarted ||
		preview.DirectLaunchEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		preview.ReceiptPathExposed ||
		preview.StateRootPathExposed {
		t.Fatalf("unexpected controlled dispatch request preview: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	if strings.Contains(text, strings.ToLower(stateRoot)) {
		t.Fatalf("controlled dispatch request preview exposed state root path: %s", text)
	}
	for _, forbidden := range []string{".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("controlled dispatch request preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func TestPreviewKnownAppLaunchGateFailsClosedForMissingAndMismatchedReceipts(t *testing.T) {
	stateRoot := t.TempDir()
	missing, err := PreviewKnownAppLaunchGate(KnownAppLaunchGateRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		ReceiptID: "known-app-launch-authorization-7zr-26.02",
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppLaunchGate missing receipt returned error: %v", err)
	}
	if missing.ReceiptLookupState != "missing-receipt" ||
		missing.ReceiptAccepted ||
		missing.LaunchGateState != "missing-receipt-fail-closed" ||
		missing.ControlledDispatchReady ||
		missing.BackendLaunchEnabled {
		t.Fatalf("unexpected missing receipt preview: %#v", missing)
	}

	receipt, err := RecordKnownAppLaunchAuthorizationReceipt(KnownAppLaunchAuthorizationReceiptRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		Authorize: KnownAppLaunchAuthorizationReceiptAction,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppLaunchAuthorizationReceipt returned error: %v", err)
	}
	receiptPath, err := KnownAppLaunchAuthorizationReceiptPath(stateRoot, receipt.ReceiptID)
	if err != nil {
		t.Fatalf("KnownAppLaunchAuthorizationReceiptPath returned error: %v", err)
	}
	content, err := os.ReadFile(receiptPath)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	var tampered map[string]any
	if err := json.Unmarshal(content, &tampered); err != nil {
		t.Fatalf("Unmarshal receipt returned error: %v", err)
	}
	tampered["app_id"] = "other-app"
	tamperedContent, err := json.MarshalIndent(tampered, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent tampered receipt returned error: %v", err)
	}
	if err := os.WriteFile(receiptPath, append(tamperedContent, '\n'), 0o600); err != nil {
		t.Fatalf("WriteFile tampered receipt returned error: %v", err)
	}
	mismatched, err := PreviewKnownAppLaunchGate(KnownAppLaunchGateRequest{
		AppID:     "7zr",
		StateRoot: stateRoot,
		ReceiptID: receipt.ReceiptID,
	})
	if err != nil {
		t.Fatalf("PreviewKnownAppLaunchGate mismatched receipt returned error: %v", err)
	}
	if mismatched.ReceiptLookupState != "rejected-receipt" ||
		mismatched.ReceiptAccepted ||
		mismatched.LaunchGateState != "rejected-receipt-fail-closed" ||
		mismatched.ControlledDispatchReady ||
		mismatched.BackendLaunchEnabled {
		t.Fatalf("unexpected mismatched receipt preview: %#v", mismatched)
	}
}
