package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestWindowsCompatibilityWorkstreamsPreviewDefinesKDEFirstGoOwnedPlan(t *testing.T) {
	preview, err := NewWindowsCompatibilityWorkstreamsPreview()
	if err != nil {
		t.Fatalf("NewWindowsCompatibilityWorkstreamsPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.windows_compatibility_workstreams.v1" ||
		preview.RequestType != "windows-compatibility-workstreams-preview" ||
		preview.PlanType != "kde-first-windows-compatibility-workstreams" ||
		preview.Source != "go-runtime-product-workstream-model" ||
		preview.ReadMethod != "GetWindowsCompatibilityWorkstreamsPreview" ||
		preview.OfficialDesktop != "KDE Plasma" ||
		preview.ProductTarget != "best Linux desktop for existing Windows applications" {
		t.Fatalf("unexpected workstream schema: %#v", preview)
	}
	if strings.Join(preview.FutureDesktopOrder, ",") != "KDE Plasma,GNOME,XFCE" {
		t.Fatalf("unexpected future desktop order: %#v", preview.FutureDesktopOrder)
	}
	if preview.EntryPointCount != 7 ||
		strings.Join(preview.EntryPointIDs, ",") != "start-menu,task-manager,file-manager,system-tray,notification-center,ai-compatibility-center,unified-settings" {
		t.Fatalf("unexpected entrypoints: %#v", preview.EntryPointIDs)
	}
	for _, entryPoint := range preview.EntryPoints {
		if !entryPoint.RequiresRuntimeGate ||
			entryPoint.KDEOwnsPolicy ||
			entryPoint.HostRootModified ||
			entryPoint.BackendLaunchEnabled ||
			entryPoint.BackendDetailsExposed ||
			!entryPoint.UserFacingOnly ||
			!entryPoint.NormalApplicationShape {
			t.Fatalf("entrypoint safety gate unexpectedly open for %s: %#v", entryPoint.ID, entryPoint)
		}
	}
	fileManager := findWindowsCompatibilityEntryPoint(preview.EntryPoints, "file-manager")
	if fileManager.KDEComponent != "Dolphin" ||
		fileManager.RuntimeSource != "file-open-preview" ||
		!fileManager.RequiresPortalReview {
		t.Fatalf("unexpected file-manager entrypoint: %#v", fileManager)
	}
	if preview.WorkstreamCount != 11 ||
		strings.Join(preview.WorkstreamIDs, ",") != "CW1,CW2,CW3,CW4,CW5,CW6,CW7,CW8,CW9,CW10,CW11" {
		t.Fatalf("unexpected workstream ids: %#v", preview.WorkstreamIDs)
	}
	if preview.FirstWaveCount != 4 ||
		strings.Join(preview.FirstWaveIDs, ",") != "CW1,CW2,CW3,CW10" {
		t.Fatalf("unexpected first wave: %#v", preview.FirstWaveIDs)
	}
	for _, entry := range preview.FirstWaveWorkstreams {
		if entry.HostRootModified || entry.NetworkRequired || entry.BackendLaunchEnabled {
			t.Fatalf("first-wave gate unexpectedly open for %s: %#v", entry.Workstream, entry)
		}
	}
	cw1 := findWindowsCompatibilityWorkstream(preview.Workstreams, "CW1")
	if cw1.MainlinePackage != "M1" ||
		!cw1.FirstWave ||
		cw1.ImplementationOwner != "Go Runtime" ||
		cw1.UnsafeBehaviorEnabled ||
		cw1.BackendLaunchEnabled ||
		cw1.HostRootModified {
		t.Fatalf("unexpected CW1 workstream: %#v", cw1)
	}
	cw10 := findWindowsCompatibilityWorkstream(preview.Workstreams, "CW10")
	if cw10.MainlinePackage != "M8" ||
		!cw10.FirstWave ||
		cw10.ImplementationOwner != "Ruby tooling plus Go fixtures" {
		t.Fatalf("unexpected CW10 workstream: %#v", cw10)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		!preview.CCoreAllowed ||
		preview.RubyCoreLogicAllowed ||
		!preview.RubyTestHarness ||
		preview.KDEPolicyOwner ||
		!preview.KDEPresentationOnly ||
		preview.DeepDesktopForkRequired ||
		preview.GNOMEFirstReleaseSupported ||
		preview.XFCEFirstReleaseSupported ||
		preview.ProductionDBusOwnershipEnabled ||
		preview.RuntimeWriteMethodsEnabled ||
		preview.BackendLaunchEnabled ||
		preview.PortalTransportCallsEnabled ||
		preview.NetworkFetchEnabled ||
		preview.HostPackageManagerEnabled ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.HostRootModified ||
		preview.RawExecutablePathExposed ||
		preview.CompatibilityStoragePathExposed ||
		preview.BackendCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected safety flags: %#v", preview)
	}
	if preview.ProtectedImplementationPackageDoc != "docs/claude-code-implementation-packages.md" {
		t.Fatalf("unexpected protected doc: %s", preview.ProtectedImplementationPackageDoc)
	}
	assertWindowsCompatibilityWorkstreamsSafe(t, preview)
}

func findWindowsCompatibilityEntryPoint(entryPoints []WindowsCompatibilityEntryPoint, id string) WindowsCompatibilityEntryPoint {
	for _, entryPoint := range entryPoints {
		if entryPoint.ID == id {
			return entryPoint
		}
	}
	return WindowsCompatibilityEntryPoint{}
}

func findWindowsCompatibilityWorkstream(workstreams []WindowsCompatibilityWorkstream, id string) WindowsCompatibilityWorkstream {
	for _, workstream := range workstreams {
		if workstream.ID == id {
			return workstream
		}
	}
	return WindowsCompatibilityWorkstream{}
}

func assertWindowsCompatibilityWorkstreamsSafe(t *testing.T, preview WindowsCompatibilityWorkstreamsPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("Windows compatibility workstreams preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
