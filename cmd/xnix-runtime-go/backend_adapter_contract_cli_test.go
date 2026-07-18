package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendAdapterContractPreviewCommandRendersNoopBoundary(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-contract-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if payload["version"] != currentProjectVersion(t) ||
		payload["schema_version"] != "xnix.runtime.backend_adapter_contract.v1" ||
		payload["request_type"] != "backend-adapter-contract-preview" ||
		payload["contract_type"] != "compatibility-backend-adapter-noop-contract" ||
		payload["runtime_method"] != "GetBackendAdapterContract" ||
		payload["go_runtime_backed"] != true ||
		payload["runtime_owned"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["noop_implementation"] != true ||
		payload["adapter_invocation_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["backend_install_enabled"] != false ||
		payload["backend_download_enabled"] != false ||
		payload["backend_process_started"] != false ||
		payload["vm_process_started"] != false ||
		payload["command_materialized"] != false ||
		payload["executable_path_resolved"] != false ||
		payload["backend_details_exposed_to_kde"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected backend adapter contract payload: %#v", payload)
	}
	counts := payload["counts"].(map[string]any)
	if counts["total_adapters"] != float64(3) ||
		counts["noop_adapters"] != float64(3) ||
		counts["enabled_invocations"] != float64(0) ||
		counts["enabled_launches"] != float64(0) ||
		counts["materialized_command"] != float64(0) {
		t.Fatalf("unexpected backend adapter contract counts: %#v", counts)
	}
	adapterIDs := payload["adapter_ids"].([]any)
	if len(adapterIDs) != 3 || adapterIDs[0] != "wine" || adapterIDs[1] != "proton" || adapterIDs[2] != "windows-vm" {
		t.Fatalf("unexpected backend adapter ids: %#v", adapterIDs)
	}
}

func TestBackendAdapterContractPreviewCommandKeepsKDEProjectionBackendSafe(t *testing.T) {
	var output bytes.Buffer
	if err := run([]string{"backend-adapter-contract-preview", "--root", projectRootForRuntimeServiceBindingCommandTest(t)}, &output); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	projection, err := json.Marshal(struct {
		KDEFacingProfiles  any `json:"kde_facing_profiles"`
		DesktopSafeSummary any `json:"desktop_safe_summary"`
	}{
		KDEFacingProfiles:  payload["kde_facing_profiles"],
		DesktopSafeSummary: payload["desktop_safe_summary"],
	})
	if err != nil {
		t.Fatalf("Marshal projection returned error: %v", err)
	}
	text := strings.ToLower(string(projection))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "windows-vm"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend adapter KDE projection exposed forbidden term %q: %s", forbidden, string(projection))
		}
	}
}

func TestBackendAdapterContractPreviewCommandRejectsArguments(t *testing.T) {
	var rejected bytes.Buffer
	if err := run([]string{"backend-adapter-contract-preview", "extra"}, &rejected); err == nil {
		t.Fatalf("backend-adapter-contract-preview must reject positional arguments")
	}
}
