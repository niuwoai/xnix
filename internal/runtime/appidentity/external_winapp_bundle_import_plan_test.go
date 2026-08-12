package appidentity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPreviewExternalWinAppBundleImportPlanAccountsForPortableSidecars(t *testing.T) {
	root := t.TempDir()
	writeBundleFile(t, root, "notepad++.exe", []byte("MZportable-app"))
	writeBundleFile(t, root, "config.xml", []byte("<config/>"))
	writeBundleFile(t, root, filepath.Join("plugins", "compare", "compare.dll"), []byte("plugin-bytes"))
	version := currentProjectVersion(t)

	plan, err := PreviewExternalWinAppBundleImportPlan(ExternalWinAppBundleImportPlanRequest{
		Version:                version,
		BundleRoot:             root,
		ExecutableRelativePath: "notepad++.exe",
		AppID:                  "org.xnix.external.notepadplusplus",
		DisplayName:            "Notepad++ Portable",
		AppVersion:             "8.9.7",
	})
	if err != nil {
		t.Fatalf("PreviewExternalWinAppBundleImportPlan returned error: %v", err)
	}
	if plan.SchemaVersion != ExternalWinAppBundleImportPlanSchemaVersion ||
		plan.RequestType != ExternalWinAppBundleImportPlanRequestType ||
		plan.Source != "go-runtime-external-winapp-bundle-import" ||
		plan.RuntimeMethod != "PreviewExternalWinAppBundleImportPlan" {
		t.Fatalf("unexpected bundle import plan identity: %#v", plan)
	}
	if !plan.WindowsExecutableValidated ||
		!plan.WindowsExecutableMZHeaderVerified ||
		!plan.DirectoryLayoutValidated ||
		!plan.SidecarDirectorySupported ||
		!plan.BundleRootConsumed ||
		!plan.PortableBundleImportReady {
		t.Fatalf("portable bundle readiness flags were not closed over the inspected directory: %#v", plan)
	}
	if plan.BundleKind != "portable-directory" ||
		plan.ExecutableName != "notepad++.exe" ||
		plan.ExecutableRelativePath != "notepad++.exe" ||
		plan.BundleFileCount != 3 ||
		plan.SidecarFileCount != 2 ||
		plan.BundleDirectoryCount != 2 ||
		plan.SidecarDirectoryCount != 2 ||
		plan.ManifestEntryCount != 3 ||
		plan.SingleExecutableImportCompatible {
		t.Fatalf("portable bundle sidecar counts were not captured: %#v", plan)
	}
	if len(plan.ExecutableSHA256) != 64 || len(plan.BundleManifestSHA256) != 64 {
		t.Fatalf("portable bundle digests must be sha256 hex strings: %#v", plan)
	}
	if !plan.RuntimeOwned || !plan.GoRuntimeBacked || plan.KDEPolicyOwner ||
		plan.LaunchEnabled || plan.BackendLaunchEnabled || plan.ActionExecutionEnabled ||
		plan.BundleRootPathExposed || plan.RawExecutablePathExposed || plan.BackendDetailsExposed ||
		plan.HostRootModified || plan.NetworkRequired || plan.PrivilegedContainerRequired ||
		plan.DockerSocketMounted || plan.BroadHostMountRequired || plan.PackageManagerInvoked {
		t.Fatalf("portable bundle import plan exposed unsafe gates: %#v", plan)
	}
	if strings.Contains(plan.DesktopSafeSummary, root) {
		t.Fatalf("desktop safe summary must not expose the local bundle root: %q", plan.DesktopSafeSummary)
	}
}

func TestPreviewExternalWinAppBundleImportPlanRejectsSymlinkSidecar(t *testing.T) {
	root := t.TempDir()
	writeBundleFile(t, root, "app.exe", []byte("MZapp"))
	version := currentProjectVersion(t)
	target := filepath.Join(root, "app.exe")
	link := filepath.Join(root, "unsafe-link.exe")
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symlink creation is unavailable: %v", err)
	}

	_, err := PreviewExternalWinAppBundleImportPlan(ExternalWinAppBundleImportPlanRequest{
		Version:                version,
		BundleRoot:             root,
		ExecutableRelativePath: "app.exe",
		AppID:                  "org.xnix.external.symlinked",
		DisplayName:            "Symlinked App",
		AppVersion:             "1.0.0",
	})
	if err == nil || !strings.Contains(err.Error(), "rejects symlink") {
		t.Fatalf("expected symlink sidecar to be rejected, got %v", err)
	}
}

func writeBundleFile(t *testing.T, root string, relativePath string, content []byte) {
	t.Helper()
	path := filepath.Join(root, relativePath)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("create bundle fixture directory: %v", err)
	}
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("write bundle fixture file: %v", err)
	}
}
