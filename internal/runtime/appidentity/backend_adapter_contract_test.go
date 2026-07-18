package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendAdapterContractPreviewDefinesNoopBoundary(t *testing.T) {
	preview, err := NewBackendAdapterContractPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterContractPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.backend_adapter_contract.v1" ||
		preview.RequestType != "backend-adapter-contract-preview" ||
		preview.ContractType != "compatibility-backend-adapter-noop-contract" ||
		preview.Source != "go-runtime-backend-manager+adapter-noop-boundary" ||
		preview.RuntimeMethod != "GetBackendAdapterContract" ||
		preview.ReadMethod != "GetBackendAdapterContractPreview" {
		t.Fatalf("unexpected backend adapter contract schema: %#v", preview)
	}
	if !sameStrings(preview.AdapterIDs, []string{"wine", "proton", "windows-vm"}) ||
		preview.Counts.TotalAdapters != 3 ||
		preview.Counts.NoopAdapters != 3 ||
		preview.Counts.KDEFacingProfiles != 3 ||
		preview.Counts.EnabledInvocations != 0 ||
		preview.Counts.EnabledLaunches != 0 ||
		preview.Counts.EnabledInstallers != 0 ||
		preview.Counts.MaterializedCommand != 0 {
		t.Fatalf("unexpected backend adapter counts: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.NoopImplementation ||
		preview.AdapterInvocationEnabled ||
		preview.BackendInstallEnabled ||
		preview.BackendDownloadEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.VMProcessStarted ||
		preview.CommandMaterialized ||
		preview.ExecutablePathResolved ||
		preview.RawCommandExposed ||
		preview.ProfilePathExposed ||
		preview.StateRootPathExposed ||
		preview.BackendDetailsExposedToKDE ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.SecretsExposed {
		t.Fatalf("unexpected backend adapter contract safety flags: %#v", preview)
	}
	if !containsString(preview.RequiredRuntimeGates, "runtime-write-gate") ||
		!containsString(preview.RequiredRuntimeGates, "restricted-launch-preflight") ||
		!containsString(preview.RequiredRuntimeGates, "test-only-materialization") {
		t.Fatalf("backend adapter contract must require runtime gates: %#v", preview.RequiredRuntimeGates)
	}
}

func TestBackendAdapterContractAdaptersStayNoop(t *testing.T) {
	preview, err := NewBackendAdapterContractPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterContractPreview returned error: %v", err)
	}

	for _, adapter := range preview.Adapters {
		if adapter.ContractStatus != "noop-contract" ||
			adapter.Implementation != "noop" ||
			!sameStrings(adapter.AllowedOperations, []string{"inspect-contract", "report-disabled-gates"}) ||
			len(adapter.RequiredInputs) < 6 ||
			len(adapter.RequiredGates) < 10 ||
			adapter.AdapterInvocationEnabled ||
			adapter.InstallEnabled ||
			adapter.DownloadEnabled ||
			adapter.LaunchEnabled ||
			adapter.ProcessStarted ||
			adapter.VMProcessStarted ||
			adapter.CommandMaterialized ||
			adapter.ExecutablePathResolved ||
			adapter.RawCommandExposed ||
			adapter.ProfilePathExposed ||
			adapter.StateRootPathExposed ||
			adapter.BackendDetailsExposedToKDE ||
			adapter.NetworkRequired ||
			adapter.HostRootModified {
			t.Fatalf("unexpected adapter no-op contract: %#v", adapter)
		}
	}
}

func TestBackendAdapterContractKDEProjectionIsBackendSafe(t *testing.T) {
	preview, err := NewBackendAdapterContractPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterContractPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(struct {
		KDEFacingProfiles  []BackendAdapterProfile `json:"kde_facing_profiles"`
		DesktopSafeSummary string                  `json:"desktop_safe_summary"`
	}{
		KDEFacingProfiles:  preview.KDEFacingProfiles,
		DesktopSafeSummary: preview.DesktopSafeSummary,
	})
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "windows-vm"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend adapter KDE projection exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
}
