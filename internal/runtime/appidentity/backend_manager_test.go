package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBackendManagerPreviewTracksRuntimeBackendsWithoutStartingThem(t *testing.T) {
	preview := NewBackendManagerPreview()

	if preview.SchemaVersion != "xnix.runtime.backend_manager.v1" ||
		preview.RequestType != "backend-manager-preview" ||
		preview.ManagerType != "compatibility-backend-manager" ||
		preview.Source != "go-runtime-backend-manager" ||
		preview.ReadMethod != "GetBackendManagerPreview" {
		t.Fatalf("unexpected backend manager schema: %#v", preview)
	}
	if !sameStrings(preview.BackendIDs, []string{"wine", "proton", "windows-vm"}) ||
		preview.BackendCount != 3 ||
		!preview.WineManaged ||
		!preview.ProtonManaged ||
		!preview.WindowsVMManaged {
		t.Fatalf("unexpected managed backend inventory: %#v", preview)
	}
	if !sameStrings(preview.UserFacingProfileIDs, []string{"automatic", "local-compatibility", "isolated-compatibility"}) ||
		preview.UserFacingProfileCount != 3 {
		t.Fatalf("unexpected user-facing backend profiles: %#v", preview.UserFacingProfiles)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.KDEVisible ||
		preview.BackendInstallEnabled ||
		preview.BackendDownloadEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.VMProcessStarted ||
		preview.RawCommandExposed ||
		preview.ProfilePathExposed ||
		preview.BackendDetailsExposedToKDE ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.SecretsExposed {
		t.Fatalf("unexpected backend manager safety flags: %#v", preview)
	}
	for _, backend := range preview.Backends {
		if backend.Status != "planned" ||
			backend.Readiness != "blocked-until-runtime-gates-pass" ||
			len(backend.RequiredGates) < 5 ||
			backend.InstallEnabled ||
			backend.DownloadEnabled ||
			backend.LaunchEnabled ||
			backend.ProcessStarted ||
			backend.RawCommandExposed ||
			backend.ProfilePathExposed ||
			backend.BackendDetailsExposedToKDE ||
			backend.HostRootModified ||
			backend.NetworkRequired {
			t.Fatalf("unexpected managed backend safety flags: %#v", backend)
		}
	}
}

func TestBackendManagerPreviewKeepsUserFacingProfilesBackendSafe(t *testing.T) {
	preview := NewBackendManagerPreview()
	encoded, err := json.Marshal(struct {
		UserFacingProfiles []UserFacingBackendProfile `json:"user_facing_profiles"`
		DesktopSafeSummary string                     `json:"desktop_safe_summary"`
	}{
		UserFacingProfiles: preview.UserFacingProfiles,
		DesktopSafeSummary: preview.DesktopSafeSummary,
	})
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "windows-vm"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend manager user-facing projection exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
}

func TestRecordBackendManagerPreviewPersistsStateRootInventory(t *testing.T) {
	stateRoot := t.TempDir()
	record, err := RecordBackendManagerPreview(stateRoot)
	if err != nil {
		t.Fatalf("RecordBackendManagerPreview returned error: %v", err)
	}

	if record.SchemaVersion != "xnix.runtime.backend_manager_record.v1" ||
		record.RecordType != "backend-manager-inventory-record" ||
		record.Source != "go-runtime-state-root-backend-manager" ||
		record.RelativePath != "backend-manager/inventory.json" ||
		record.SHA256 == "" {
		t.Fatalf("unexpected backend manager record schema: %#v", record)
	}
	if !record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		record.StateRootPathExposed ||
		record.BackendInstallEnabled ||
		record.BackendDownloadEnabled ||
		record.BackendLaunchEnabled ||
		record.BackendProcessStarted ||
		record.VMProcessStarted ||
		record.RawCommandExposed ||
		record.ProfilePathExposed ||
		record.BackendDetailsExposedToKDE ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.SecretsExposed {
		t.Fatalf("unexpected backend manager record safety flags: %#v", record)
	}
	if !sameStrings(record.Preview.BackendIDs, []string{"wine", "proton", "windows-vm"}) {
		t.Fatalf("unexpected backend manager record inventory: %#v", record.Preview.BackendIDs)
	}
	data, err := os.ReadFile(filepath.Join(stateRoot, filepath.FromSlash(record.RelativePath)))
	if err != nil {
		t.Fatalf("ReadFile backend manager record returned error: %v", err)
	}
	var persisted BackendManagerRecord
	if err := json.Unmarshal(data, &persisted); err != nil {
		t.Fatalf("Unmarshal backend manager record returned error: %v", err)
	}
	if persisted.SHA256 != record.SHA256 ||
		persisted.StateRootPathExposed ||
		persisted.HostRootModified ||
		persisted.BackendLaunchEnabled {
		t.Fatalf("unexpected persisted backend manager record: %#v", persisted)
	}
}

func TestRecordBackendManagerPreviewRejectsFilesystemRoot(t *testing.T) {
	if _, err := RecordBackendManagerPreview(string(os.PathSeparator)); err == nil {
		t.Fatalf("RecordBackendManagerPreview must reject filesystem root")
	}
}
