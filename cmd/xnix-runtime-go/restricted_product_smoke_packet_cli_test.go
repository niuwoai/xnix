package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRestrictedProductSmokePacketPreviewCLI(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-product-smoke-packet-preview", "--repo-root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output); err != nil {
		t.Fatalf("run restricted smoke packet: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("parse output: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.restricted_product_smoke_packet.v1" ||
		payload["request_type"] != "restricted-product-smoke-packet-preview" || payload["packet_prepared"] != true || payload["ready_for_authorized_smoke"] != true {
		t.Fatalf("unexpected packet identity: %+v", payload)
	}
	for _, key := range []string{"execution_authorized", "docker_executed", "qemu_executed", "product_smoke_executed", "serial_log_persisted", "docker_socket_mounted", "host_network_enabled", "broad_host_mount_enabled", "privileged_container_required", "backend_launch_enabled", "host_root_modified", "release_ready"} {
		if payload[key] != false {
			t.Fatalf("packet must keep %q disabled: %+v", key, payload)
		}
	}
	if payload["human_authorization_required"] != true || payload["loopback_only_networking"] != true || payload["serial_log_persistence_required"] != true {
		t.Fatalf("packet did not preserve restricted smoke requirements: %+v", payload)
	}
}

func TestRestrictedProductSmokePacketPreviewCLIRejectsManifestEscape(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"restricted-product-smoke-packet-preview", "--repo-root", t.TempDir(), "--manifest", "../manifest.json"}, &output); err == nil {
		t.Fatal("expected manifest escape to fail")
	}
}
