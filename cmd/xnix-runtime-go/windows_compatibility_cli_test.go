package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestWindowsCompatibilityWorkstreamsPreviewCommand(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"windows-compatibility-workstreams-preview"}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["schema_version"] != "xnix.runtime.windows_compatibility_workstreams.v1" ||
		payload["request_type"] != "windows-compatibility-workstreams-preview" ||
		payload["plan_type"] != "kde-first-windows-compatibility-workstreams" ||
		payload["source"] != "go-runtime-product-workstream-model" ||
		payload["read_method"] != "GetWindowsCompatibilityWorkstreamsPreview" ||
		payload["official_desktop"] != "KDE Plasma" ||
		payload["product_target"] != "best Linux desktop for existing Windows applications" {
		t.Fatalf("unexpected workstream payload schema: %#v", payload)
	}
	if payload["entry_point_count"] != float64(7) ||
		payload["workstream_count"] != float64(11) ||
		payload["first_wave_count"] != float64(4) {
		t.Fatalf("unexpected workstream counts: %#v", payload)
	}
	entryIDs := payload["entry_point_ids"].([]any)
	if strings.Join(anyStrings(entryIDs), ",") != "start-menu,task-manager,file-manager,system-tray,notification-center,ai-compatibility-center,unified-settings" {
		t.Fatalf("unexpected entrypoint ids: %#v", entryIDs)
	}
	workstreamIDs := payload["workstream_ids"].([]any)
	if strings.Join(anyStrings(workstreamIDs), ",") != "CW1,CW2,CW3,CW4,CW5,CW6,CW7,CW8,CW9,CW10,CW11" {
		t.Fatalf("unexpected workstream ids: %#v", workstreamIDs)
	}
	firstWaveIDs := payload["first_wave_ids"].([]any)
	if strings.Join(anyStrings(firstWaveIDs), ",") != "CW1,CW2,CW3,CW10" {
		t.Fatalf("unexpected first-wave ids: %#v", firstWaveIDs)
	}
	if payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["c_core_allowed"] != true ||
		payload["ruby_core_logic_allowed"] != false ||
		payload["ruby_test_harness"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_presentation_only"] != true ||
		payload["deep_desktop_fork_required"] != false ||
		payload["gnome_first_release_supported"] != false ||
		payload["xfce_first_release_supported"] != false ||
		payload["production_dbus_ownership_enabled"] != false ||
		payload["runtime_write_methods_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["portal_transport_calls_enabled"] != false ||
		payload["network_fetch_enabled"] != false ||
		payload["host_package_manager_enabled"] != false ||
		payload["privileged_container_required"] != false ||
		payload["host_networking_required"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["broad_host_mount_required"] != false ||
		payload["host_root_modified"] != false ||
		payload["raw_executable_path_exposed"] != false ||
		payload["compatibility_storage_path_exposed"] != false ||
		payload["backend_command_exposed"] != false ||
		payload["backend_details_exposed"] != false {
		t.Fatalf("unexpected safety flags: %#v", payload)
	}
	if payload["protected_implementation_package_doc"] != "docs/claude-code-implementation-packages.md" {
		t.Fatalf("unexpected protected implementation package doc: %#v", payload["protected_implementation_package_doc"])
	}
	if !strings.Contains(payload["desktop_safe_summary"].(string), "CW1, CW2, CW3, and CW10") {
		t.Fatalf("desktop summary does not name the first wave: %#v", payload["desktop_safe_summary"])
	}
	assertWindowsCompatibilityCLISafe(t, output.String())
}

func TestWindowsCompatibilityWorkstreamsPreviewCommandRejectsArguments(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"windows-compatibility-workstreams-preview", "--registry", "ignored"}, &output)
	if err == nil || !strings.Contains(err.Error(), "does not accept arguments") {
		t.Fatalf("expected argument rejection, got %v", err)
	}
}

func anyStrings(values []any) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.(string))
	}
	return result
}

func assertWindowsCompatibilityCLISafe(t *testing.T, text string) {
	t.Helper()
	serialized := strings.ToLower(text)
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("Windows compatibility CLI exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
