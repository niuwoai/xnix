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
