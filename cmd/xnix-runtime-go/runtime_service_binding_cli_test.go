package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRuntimeServiceBindingPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-service-binding-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.303" ||
		payload["schema_version"] != "xnix.runtime.service_binding.v1" ||
		payload["request_type"] != "runtime-service-binding-preview" ||
		payload["binding_type"] != "runtime-service-binding" ||
		payload["source"] != "activation-files+dbus-contract+runtime-owner-gate" ||
		payload["runtime_method"] != "GetRuntimeServiceBinding" ||
		payload["read_method"] != "GetRuntimeServiceBindingPreview" ||
		payload["production_status"] != "pending-live-owner" {
		t.Fatalf("unexpected Runtime service binding CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	activation := payload["activation"].(map[string]any)
	if activation["dbus_service_file"] != "runtime/dbus/org.xnix.Compatibility1.service" ||
		activation["systemd_unit"] != "runtime/systemd/xnix-compatd.service" ||
		activation["libexec_wrapper"] != "libexec/xnix/compatd" ||
		activation["dbus_contract"] != "runtime/dbus/org.xnix.Compatibility1.xml" ||
		activation["packaged_wrapper"] != "/usr/libexec/xnix/compatd" {
		t.Fatalf("unexpected activation payload: %#v", activation)
	}
	checks := payload["checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedIDs := []string{
		"dbus-service-activation",
		"systemd-service-hardening",
		"libexec-wrapper",
		"dbus-contract",
		"live-dbus-owner",
	}
	if len(checks) != len(expectedIDs) || len(checkIDs) != len(expectedIDs) {
		t.Fatalf("unexpected Runtime service binding checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id || checkIDs[index] != id {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
		expectedStatus := "pass"
		if id == "live-dbus-owner" {
			expectedStatus = "pending"
		}
		if check["status"] != expectedStatus {
			t.Fatalf("unexpected status for %s: %#v", id, check)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(4) ||
		counts["pending"] != float64(1) ||
		counts["blocked"] != float64(0) {
		t.Fatalf("unexpected Runtime service binding counts: %#v", counts)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["activation_binding_ready"] != true ||
		payload["live_dbus_owner_ready"] != false ||
		payload["smoke_adapter_available"] != true ||
		payload["service_started"] != false ||
		payload["production_bus_claimed"] != false ||
		payload["kde_may_claim_runtime_ownership"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected Runtime service binding safety flags: %#v", payload)
	}
	blockedActions := payload["blocked_actions"].([]any)
	if len(blockedActions) != 4 ||
		blockedActions[0] != "start Runtime service from preview" ||
		blockedActions[1] != "claim production D-Bus owner from preview" ||
		blockedActions[2] != "let KDE claim Runtime ownership" ||
		blockedActions[3] != "mutate host root during service binding planning" {
		t.Fatalf("unexpected blocked actions: %#v", blockedActions)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime service binding CLI exposed backend terms: %s", output.String())
	}
}

func projectRootForRuntimeServiceBindingCommandTest(t *testing.T) string {
	t.Helper()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd returned error: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(workingDirectory, "VERSION")); err == nil {
			return workingDirectory
		}
		parent := filepath.Dir(workingDirectory)
		if parent == workingDirectory || strings.TrimSpace(parent) == "" {
			t.Fatalf("could not find project root from %s", workingDirectory)
		}
		workingDirectory = parent
	}
}
