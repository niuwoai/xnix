package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCompatibilityBackendFallbackPreviewCLI(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	var output bytes.Buffer
	if err := run([]string{"compatibility-backend-fallback-preview", "--registry", registryPath, "--app", app}, &output); err != nil {
		t.Fatalf("run fallback preview: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.compatibility_backend_fallback.v1" ||
		payload["request_type"] != "compatibility-backend-fallback-preview" ||
		payload["primary_profile_id"] != "local-compatibility" {
		t.Fatalf("unexpected payload identity: %+v", payload)
	}
	for _, key := range []string{"selection_persisted", "engine_install_enabled", "backend_launch_enabled", "vm_start_enabled", "backend_details_exposed", "host_root_modified", "state_root_path_exposed"} {
		if payload[key] != false {
			t.Fatalf("preview must keep %q disabled: %+v", key, payload)
		}
	}
	for _, term := range []string{"wine", "proton", "qemu"} {
		if bytes.Contains(bytes.ToLower(output.Bytes()), []byte(term)) {
			t.Fatalf("CLI output must not expose backend term %q: %s", term, output.String())
		}
	}
}

func TestCompatibilityBackendFallbackPreviewCLIRequiresSource(t *testing.T) {
	registryPath, app := writeTestRepairGroupRegistry(t)
	tests := [][]string{
		{"compatibility-backend-fallback-preview"},
		{"compatibility-backend-fallback-preview", "--registry", registryPath},
		{"compatibility-backend-fallback-preview", "--registry", registryPath, "--app", app, "extra"},
	}
	for _, args := range tests {
		var output bytes.Buffer
		if err := run(args, &output); err == nil {
			t.Fatalf("expected CLI args to fail: %+v", args)
		}
	}
}
