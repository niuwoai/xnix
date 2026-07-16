package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDesktopIconPreviewCommandConsumesActivationRoot(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	stageRoot := t.TempDir()

	var stageOutput bytes.Buffer
	if err := run([]string{"desktop-activation-stage", "--registry", registryPath, "--app", app, "--mode", "development", "--staging-root", stageRoot}, &stageOutput); err != nil {
		t.Fatalf("stage run returned error: %v", err)
	}

	var output bytes.Buffer
	if err := run([]string{"desktop-icon-preview", "--registry", registryPath, "--app", app, "--activation-root", stageRoot}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	payload := decodeWindowIdentityPayload(t, output.Bytes())
	if payload["source"] != "desktop-entry-preview+desktop-activation-receipt" ||
		payload["activation_receipt_root"] != true ||
		payload["activation_receipt_backed"] != true ||
		payload["activation_receipt_path"] != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("desktop icon preview did not consume activation receipt: %#v", payload)
	}
	if payload["desktop_file_copy_enabled"] != false ||
		payload["desktop_file_write_enabled"] != false ||
		payload["icon_placement_persisted"] != false ||
		payload["launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["host_root_modified"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["raw_executable_exposed"] != false {
		t.Fatalf("receipt-backed desktop icon preview must remain gated: %#v", payload)
	}
	if strings.Contains(output.String(), stageRoot) {
		t.Fatalf("desktop icon preview exposed activation root: %s", output.String())
	}
	assertWindowIdentityPayloadSafe(t, output.String())
}
