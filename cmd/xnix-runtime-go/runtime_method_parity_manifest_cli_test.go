package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRuntimeMethodParityManifestPreviewCommandRendersGoReadModel(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"runtime-method-parity-manifest-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != "0.2.294" ||
		payload["schema_version"] != "xnix.runtime.method_parity_manifest.v1" ||
		payload["request_type"] != "runtime-method-parity-manifest-preview" ||
		payload["manifest_type"] != "runtime-method-parity-manifest" ||
		payload["source"] != "dbus-contract+runtime-dispatch+dbus-client+smoke-adapter+session-smoke" ||
		payload["runtime_method"] != "GetRuntimeMethodParityManifest" ||
		payload["read_method"] != "GetRuntimeMethodParityManifestPreview" {
		t.Fatalf("unexpected Runtime method parity CLI schema: %#v", payload)
	}
	if payload["bus_name"] != "org.xnix.Compatibility1" ||
		payload["object_path"] != "/org/xnix/Compatibility1" ||
		payload["interface"] != "org.xnix.Compatibility1" {
		t.Fatalf("unexpected Runtime D-Bus identity: %#v", payload)
	}
	if payload["method_count"] != float64(61) {
		t.Fatalf("unexpected method count: %#v", payload["method_count"])
	}
	methods := payload["read_only_methods"].([]any)
	if len(methods) != 61 ||
		!containsAny(methods, "GetRuntimeOwnerSmokePlan") ||
		!containsAny(methods, "GetRuntimeMethodParityManifest") ||
		!containsAny(methods, "GetKDECenterPageSectionDetail") {
		t.Fatalf("unexpected read-only methods: %#v", methods)
	}
	checks := payload["parity_checks"].([]any)
	checkIDs := payload["check_ids"].([]any)
	expectedCheckIDs := []string{"dbus-contract", "runtime-dispatch", "dbus-client", "smoke-adapter", "session-smoke"}
	if len(checks) != len(expectedCheckIDs) || len(checkIDs) != len(expectedCheckIDs) {
		t.Fatalf("unexpected parity checks: %#v ids=%#v", checks, checkIDs)
	}
	for index, id := range expectedCheckIDs {
		check := checks[index].(map[string]any)
		if check["id"] != id ||
			checkIDs[index] != id ||
			check["status"] != "pass" ||
			check["method_count"] != float64(61) ||
			len(check["missing_methods"].([]any)) != 0 {
			t.Fatalf("unexpected parity check at %d: %#v ids=%#v", index, checks, checkIDs)
		}
	}
	counts := payload["counts"].(map[string]any)
	if counts["total"] != float64(5) ||
		counts["passed"] != float64(5) ||
		counts["blocked"] != float64(0) ||
		counts["pending"] != float64(0) ||
		payload["read_only_method_parity_ready"] != true {
		t.Fatalf("unexpected parity counts: counts=%#v ready=%#v", counts, payload["read_only_method_parity_ready"])
	}
	writeMethods := payload["write_methods"].([]any)
	if len(writeMethods) != 4 ||
		writeMethods[0] != "InstallRecipe" ||
		writeMethods[1] != "Launch" ||
		writeMethods[2] != "CreateSnapshot" ||
		writeMethods[3] != "RestoreSnapshot" ||
		payload["write_methods_supported"] != false ||
		payload["write_method_dispatch_enabled"] != false {
		t.Fatalf("unexpected write-method boundary: methods=%#v payload=%#v", writeMethods, payload)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["network_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["privileged_container_required"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected parity safety flags: %#v", payload)
	}
	if strings.Contains(strings.ToLower(output.String()), "prefix") ||
		strings.Contains(strings.ToLower(output.String()), ".exe") ||
		strings.Contains(strings.ToLower(output.String()), "program files") {
		t.Fatalf("Runtime method parity CLI exposed backend terms: %s", output.String())
	}
}

func containsAny(values []any, candidate string) bool {
	for _, value := range values {
		if value == candidate {
			return true
		}
	}
	return false
}
