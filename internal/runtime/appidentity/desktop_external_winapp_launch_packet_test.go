package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/winapp"
)

func TestDesktopExternalWinAppLaunchPacketPreviewConsumesReceiptAndRunRecord(t *testing.T) {
	plan := testExternalDesktopLaunchPlan(t)
	root := t.TempDir()
	writeExternalDesktopLaunchReceipt(t, root, plan.ApplicationID)
	runRecordPath := writeExternalDesktopLaunchRunRecord(t, root, ExternalWinAppRunResult{
		Version:                         "0.2.640-test",
		SchemaVersion:                   ExternalWinAppRunSchemaVersion,
		RequestType:                     ExternalWinAppRunRequestType,
		Status:                          winapp.PassedStatus,
		ApplicationID:                   plan.ApplicationID,
		DisplayName:                     plan.DisplayName,
		AppVersion:                      plan.ApplicationVersion,
		ExternalAppImportRecordConsumed: true,
		ExternalAppHandleConsumed:       true,
		ExternalAppHandle:               plan.ApplicationID,
		ImportedArtifactDigestVerified:  true,
		RuntimeRunRequested:             true,
		RuntimeRunExecuted:              true,
		ExecutionStarted:                true,
		BackendProcessStarted:           true,
		ContainerRuntimeUsed:            true,
		ContainerNetworkMode:            "none",
		ContainerHostMountCount:         0,
		WindowObserved:                  true,
		XWindowObserved:                 true,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
	})

	packet, err := plan.DesktopExternalWinAppLaunchPacketPreview(root, runRecordPath, "development")
	if err != nil {
		t.Fatalf("DesktopExternalWinAppLaunchPacketPreview returned error: %v", err)
	}
	if packet.SchemaVersion != DesktopExternalWinAppLaunchPacketSchemaVersion ||
		packet.RequestType != DesktopExternalWinAppLaunchPacketRequestType ||
		packet.PacketType != "kde-desktop-external-winapp-launch-evidence" ||
		packet.Status != "passed" ||
		packet.Source != "desktop-activation-status-preview+external-app-run-record" ||
		packet.RuntimeMethod != "GetDesktopExternalWinAppLaunchPacket" ||
		packet.ApplicationID != plan.ApplicationID ||
		packet.DisplayName != plan.DisplayName ||
		packet.ExternalAppHandle != plan.ApplicationID ||
		packet.ActivationStatusRequestType != "desktop-activation-status-preview" ||
		!packet.ActivationReceiptBacked ||
		!packet.ActivationReceiptSafeForKDE ||
		!packet.DesktopExecUsesExternalAppHandle ||
		!packet.ExternalAppDesktopHandleReady ||
		packet.DesktopExecUsesRawImportRecord ||
		packet.DesktopExecUsesStateRoot ||
		!packet.RunRecordConsumed ||
		packet.RunRecordRequestType != ExternalWinAppRunRequestType ||
		!packet.ExternalAppRunRecordConsumed ||
		!packet.ExternalAppImportRecordConsumed ||
		!packet.ExternalAppHandleConsumed ||
		!packet.ImportedArtifactDigestVerified ||
		!packet.RuntimeRunRequested ||
		!packet.RuntimeLaunchExecuted ||
		!packet.ExecutionStarted ||
		!packet.BackendProcessStarted ||
		!packet.WindowObserved ||
		!packet.XWindowObserved ||
		!packet.ContainerRuntimeUsed ||
		packet.ContainerNetworkMode != "none" ||
		packet.ContainerHostMountCount != 0 ||
		!packet.RuntimeOwned ||
		!packet.GoRuntimeBacked ||
		!packet.RuntimeLaunchAuthority ||
		packet.KDEPolicyOwner ||
		packet.KDELaunchAuthority ||
		!packet.DesktopLaunchPacketReady ||
		!packet.SafeForKDE ||
		len(packet.UnsafeReasonIDs) != 0 ||
		packet.BackendDetailsExposed ||
		packet.RawImportRecordPathExposed ||
		packet.RawExternalAppHandlePathExposed ||
		packet.RawStateRootPathExposed ||
		packet.RawExecutablePathExposed ||
		packet.HostRootModified ||
		packet.PrivilegedContainerRequired ||
		packet.HostNetworkingRequired ||
		packet.DockerSocketMounted ||
		packet.BroadHostMountRequired {
		t.Fatalf("unexpected desktop external app launch packet: %#v", packet)
	}
	encoded, err := json.Marshal(packet)
	if err != nil {
		t.Fatalf("Marshal packet returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{".exe", "wine ", "wine/", ".wine", "docker run", "/var/run/docker.sock", root, runRecordPath} {
		if strings.Contains(text, strings.ToLower(forbidden)) {
			t.Fatalf("desktop external app launch packet exposed forbidden term %q: %s", forbidden, string(encoded))
		}
	}
}

func TestDesktopExternalWinAppLaunchPacketPreviewMarksMissingWindowUnsafe(t *testing.T) {
	plan := testExternalDesktopLaunchPlan(t)
	root := t.TempDir()
	writeExternalDesktopLaunchReceipt(t, root, plan.ApplicationID)
	runRecordPath := writeExternalDesktopLaunchRunRecord(t, root, ExternalWinAppRunResult{
		Version:                         "0.2.640-test",
		SchemaVersion:                   ExternalWinAppRunSchemaVersion,
		RequestType:                     ExternalWinAppRunRequestType,
		Status:                          winapp.PassedStatus,
		ApplicationID:                   plan.ApplicationID,
		DisplayName:                     plan.DisplayName,
		AppVersion:                      plan.ApplicationVersion,
		ExternalAppImportRecordConsumed: true,
		ExternalAppHandleConsumed:       true,
		ExternalAppHandle:               plan.ApplicationID,
		ImportedArtifactDigestVerified:  true,
		RuntimeRunRequested:             true,
		RuntimeRunExecuted:              true,
		ExecutionStarted:                true,
		BackendProcessStarted:           true,
		ContainerRuntimeUsed:            true,
		ContainerNetworkMode:            "none",
	})

	packet, err := plan.DesktopExternalWinAppLaunchPacketPreview(root, runRecordPath, "development")
	if err != nil {
		t.Fatalf("DesktopExternalWinAppLaunchPacketPreview returned error: %v", err)
	}
	if packet.Status != "blocked" ||
		packet.DesktopLaunchPacketReady ||
		packet.SafeForKDE ||
		!containsString(packet.UnsafeReasonIDs, "window-not-observed") ||
		!containsString(packet.UnsafeReasonIDs, "x-window-not-observed") {
		t.Fatalf("missing window evidence should make packet unsafe: %#v", packet)
	}
}

func testExternalDesktopLaunchPlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlanWithProvenance(Recipe{
		ID:      "org.xnix.external.gui",
		Name:    "External GUI",
		Version: "0.2.640-test",
		Icon:    "application-x-executable",
		Mode:    "automatic",
		ContainerGUISmoke: ContainerGUISmokeHints{
			App:         "ExternalGui.exe",
			WindowMatch: "External GUI",
		},
	}, Provenance{
		Source:          "external-winapp-import-record",
		RegistryName:    "runtime-managed-external-apps",
		DigestVerified:  true,
		SignatureStatus: "local-import-digest-verified",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}

func writeExternalDesktopLaunchReceipt(t *testing.T, root string, applicationID string) {
	t.Helper()
	receiptPath := filepath.Join(root, filepath.FromSlash(desktopActivationReceiptRelativePath(applicationID)))
	if err := os.MkdirAll(filepath.Dir(receiptPath), 0o755); err != nil {
		t.Fatalf("MkdirAll receipt returned error: %v", err)
	}
	receipt := map[string]any{
		"schema_version": "xnix.runtime.desktop_activation_receipt.v1",
		"receipt_type":   "desktop-activation-receipt",
		"application_id": applicationID,
		"desktop_launch": map[string]any{
			"external_app_handle":                   applicationID,
			"desktop_exec_uses_external_app_handle": true,
			"external_app_desktop_handle_ready":     true,
			"desktop_exec_uses_raw_import_record":   false,
			"desktop_exec_uses_state_root":          false,
		},
		"installed": []map[string]any{
			{
				"id":                      "desktop-entry",
				"relative_path":           "usr/share/applications/xnix-org.xnix.external.gui.desktop",
				"sha256":                  strings.Repeat("a", 64),
				"written":                 true,
				"host_root_modified":      false,
				"backend_details_exposed": false,
			},
		},
		"rollback": map[string]any{
			"requires_matching_sha256": true,
			"host_root_modified":       false,
		},
		"safety": map[string]any{
			"runtime_owned":           true,
			"host_root_modified":      false,
			"backend_details_exposed": false,
		},
	}
	data, err := json.Marshal(receipt)
	if err != nil {
		t.Fatalf("Marshal receipt returned error: %v", err)
	}
	if err := os.WriteFile(receiptPath, data, 0o600); err != nil {
		t.Fatalf("WriteFile receipt returned error: %v", err)
	}
}

func writeExternalDesktopLaunchRunRecord(t *testing.T, root string, record ExternalWinAppRunResult) string {
	t.Helper()
	path := filepath.Join(root, "run-record.json")
	data, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("Marshal run record returned error: %v", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile run record returned error: %v", err)
	}
	return path
}
