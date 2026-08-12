package appidentity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecordExternalWinAppImportCopiesExecutableIntoStateRoot(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalTool.exe")
	executable := []byte{'M', 'Z', 0x90, 0x00, 'x', 'n', 'i', 'x'}
	if err := os.WriteFile(executablePath, executable, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")

	record, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.tool",
		DisplayName:    "External Tool",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	if record.SchemaVersion != ExternalWinAppImportRecordSchemaVersion ||
		record.RequestType != ExternalWinAppImportRecordRequestType ||
		record.RecordType != "external-windows-app-import-record" ||
		record.Version != "0.2.640-test" ||
		record.ApplicationID != "org.xnix.external.tool" ||
		record.DisplayName != "External Tool" ||
		record.AppVersion != "0.2.640-test" ||
		record.ExecutableName != "ExternalTool.exe" ||
		record.ArtifactSizeBytes != int64(len(executable)) ||
		!strings.HasPrefix(record.ArtifactRelativePath, "external-apps/org.xnix.external.tool/artifacts/") ||
		record.RecordRelativePath != "external-apps/org.xnix.external.tool/import-record.json" ||
		record.RecordSHA256 == "" ||
		!record.WindowsExecutableValidated ||
		!record.ArtifactCopied ||
		!record.ImportRecorded ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		!record.DesktopPageRenderable ||
		record.LaunchEnabled ||
		record.BackendLaunchEnabled ||
		record.ActionExecutionEnabled ||
		record.StateRootPathExposed ||
		record.RawExecutablePathExposed ||
		record.BackendDetailsExposed ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired ||
		record.PackageManagerInvoked {
		t.Fatalf("unexpected external import record: %#v", record)
	}
	if reasons := ValidateExternalWinAppImportRecord(record); len(reasons) != 0 {
		t.Fatalf("ValidateExternalWinAppImportRecord returned reasons: %v", reasons)
	}
	artifactBytes, err := os.ReadFile(filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath)))
	if err != nil {
		t.Fatalf("ReadFile artifact returned error: %v", err)
	}
	if string(artifactBytes) != string(executable) {
		t.Fatalf("artifact bytes mismatch: %q != %q", string(artifactBytes), string(executable))
	}
	artifactInfo, err := os.Stat(filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath)))
	if err != nil {
		t.Fatalf("Stat artifact returned error: %v", err)
	}
	if artifactInfo.Mode().Perm() != 0o644 {
		t.Fatalf("imported artifact must be container-readable, got mode %o", artifactInfo.Mode().Perm())
	}
	loaded, err := LoadExternalWinAppImportRecord(filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath)))
	if err != nil {
		t.Fatalf("LoadExternalWinAppImportRecord returned error: %v", err)
	}
	if loaded.RecordSHA256 != record.RecordSHA256 || loaded.ArtifactSHA256 != record.ArtifactSHA256 {
		t.Fatalf("loaded record mismatch: %#v != %#v", loaded, record)
	}
	resolvedRecord, resolvedArtifactPath, err := ResolveExternalWinAppImportedArtifact(filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath)))
	if err != nil {
		t.Fatalf("ResolveExternalWinAppImportedArtifact returned error: %v", err)
	}
	if resolvedRecord.RecordSHA256 != record.RecordSHA256 ||
		resolvedArtifactPath != filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath)) {
		t.Fatalf("unexpected resolved import artifact: record=%#v artifact=%s", resolvedRecord, resolvedArtifactPath)
	}
	recipe, provenance, err := ExternalAppRecipeFromImportRecord(loaded)
	if err != nil {
		t.Fatalf("ExternalAppRecipeFromImportRecord returned error: %v", err)
	}
	if recipe.ID != record.ApplicationID ||
		recipe.Name != record.DisplayName ||
		recipe.Version != record.AppVersion ||
		recipe.Icon != "application-x-executable" ||
		recipe.Mode != "automatic" ||
		len(recipe.SupportedExtensions) != 0 ||
		provenance.Source != "external-winapp-import-record" ||
		provenance.RegistryName != "runtime-managed-external-apps" ||
		!provenance.DigestVerified ||
		provenance.SignatureStatus != "local-import-digest-verified" {
		t.Fatalf("unexpected import recipe projection: recipe=%#v provenance=%#v", recipe, provenance)
	}
	page, err := NewKDECenterPagePreview(recipe, provenance, "approved", nil)
	if err != nil {
		t.Fatalf("NewKDECenterPagePreview returned error: %v", err)
	}
	if page.ApplicationID != record.ApplicationID ||
		page.ApplicationName != record.DisplayName ||
		page.Icon != "application-x-executable" ||
		page.LaunchEnabled ||
		page.BackendProcessStarted ||
		page.HostRootModified ||
		page.BackendDetailsExposed {
		t.Fatalf("unexpected KDE page from import record: %#v", page)
	}
}

func TestRecordExternalWinAppBundleImportCopiesPortableDirectoryIntoStateRoot(t *testing.T) {
	tempDir := t.TempDir()
	bundleRoot := filepath.Join(tempDir, "NotepadPlusPlusPortable")
	writeBundleFile(t, bundleRoot, "notepad++.exe", []byte{'M', 'Z', 0x90, 0x00, 'n', 'p', 'p'})
	writeBundleFile(t, bundleRoot, "config.xml", []byte("<config/>"))
	writeBundleFile(t, bundleRoot, filepath.Join("plugins", "mimeTools.dll"), []byte("plugin"))
	stateRoot := filepath.Join(tempDir, "state")

	record, err := RecordExternalWinAppBundleImport(ExternalWinAppBundleImportRequest{
		Version:                "0.2.640-test",
		StateRoot:              stateRoot,
		BundleRoot:             bundleRoot,
		ExecutableRelativePath: "notepad++.exe",
		AppID:                  "org.xnix.external.notepadplusplus",
		DisplayName:            "Notepad++ Portable",
		AppVersion:             "8.9.7",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppBundleImport returned error: %v", err)
	}
	if record.SchemaVersion != ExternalWinAppImportRecordSchemaVersion ||
		record.RequestType != ExternalWinAppBundleImportRecordRequestType ||
		record.RecordType != "external-windows-app-import-record" ||
		record.Source != "go-runtime-external-winapp-bundle-import" ||
		record.RuntimeMethod != "RecordExternalWinAppBundleImport" ||
		record.ArtifactKind != ExternalWinAppArtifactKindPortableDirectory ||
		record.ExecutableName != "notepad++.exe" ||
		record.ExecutableRelativePath != "notepad++.exe" ||
		record.BundleManifestSHA256 == "" ||
		record.BundleFileCount != 3 ||
		record.SidecarFileCount != 2 ||
		record.BundleDirectoryCount != 1 ||
		record.SidecarDirectoryCount != 1 ||
		!strings.HasPrefix(record.BundleRelativePath, "external-apps/org.xnix.external.notepadplusplus/bundles/") ||
		record.ArtifactRelativePath != record.BundleRelativePath+"/notepad++.exe" ||
		!record.WindowsExecutableValidated ||
		!record.ArtifactCopied ||
		!record.ImportRecorded ||
		!record.RuntimeOwned ||
		!record.GoRuntimeBacked ||
		record.KDEPolicyOwner ||
		record.LaunchEnabled ||
		record.BackendLaunchEnabled ||
		record.ActionExecutionEnabled ||
		record.StateRootPathExposed ||
		record.RawExecutablePathExposed ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.PrivilegedContainerRequired ||
		record.DockerSocketMounted ||
		record.BroadHostMountRequired {
		t.Fatalf("unexpected bundle import record: %#v", record)
	}
	if reasons := ValidateExternalWinAppImportRecord(record); len(reasons) != 0 {
		t.Fatalf("ValidateExternalWinAppImportRecord returned reasons: %v", reasons)
	}
	recordPath := filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath))
	resolvedRecord, artifactPath, workspacePath, executableRelativePath, err := ResolveExternalWinAppImportedArtifactLocation(recordPath)
	if err != nil {
		t.Fatalf("ResolveExternalWinAppImportedArtifactLocation returned error: %v", err)
	}
	if resolvedRecord.RecordSHA256 != record.RecordSHA256 ||
		artifactPath != filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath)) ||
		workspacePath != filepath.Join(stateRoot, filepath.FromSlash(record.BundleRelativePath)) ||
		executableRelativePath != "notepad++.exe" {
		t.Fatalf("unexpected resolved bundle import location: record=%#v artifact=%s workspace=%s exe=%s", resolvedRecord, artifactPath, workspacePath, executableRelativePath)
	}
	for _, relativePath := range []string{"notepad++.exe", "config.xml", filepath.ToSlash(filepath.Join("plugins", "mimeTools.dll"))} {
		if _, err := os.Stat(filepath.Join(workspacePath, filepath.FromSlash(relativePath))); err != nil {
			t.Fatalf("imported bundle missing %s: %v", relativePath, err)
		}
	}
	if err := os.WriteFile(filepath.Join(workspacePath, "config.xml"), []byte("<tampered/>"), 0o600); err != nil {
		t.Fatalf("WriteFile tampered bundle sidecar returned error: %v", err)
	}
	if _, _, _, _, err := ResolveExternalWinAppImportedArtifactLocation(recordPath); err == nil ||
		!strings.Contains(err.Error(), "bundle manifest digest mismatch") {
		t.Fatalf("expected bundle manifest mismatch after tampering, got %v", err)
	}
}

func TestResolveExternalWinAppImportedArtifactRejectsTamperedArtifact(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalTool.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.tool",
		DisplayName:    "External Tool",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}
	artifactPath := filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath))
	if err := os.WriteFile(artifactPath, []byte{'M', 'Z', 'x'}, 0o600); err != nil {
		t.Fatalf("WriteFile tampered artifact returned error: %v", err)
	}
	if _, _, err := ResolveExternalWinAppImportedArtifact(filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath))); err == nil ||
		!strings.Contains(err.Error(), "digest mismatch") {
		t.Fatalf("expected imported artifact digest mismatch, got %v", err)
	}
}

func TestResolveExternalWinAppImportedArtifactByHandle(t *testing.T) {
	tempDir := t.TempDir()
	executablePath := filepath.Join(tempDir, "ExternalTool.exe")
	if err := os.WriteFile(executablePath, []byte{'M', 'Z', 0x90, 0x00}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	stateRoot := filepath.Join(tempDir, "state")
	record, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      stateRoot,
		ExecutablePath: executablePath,
		AppID:          "org.xnix.external.tool",
		DisplayName:    "External Tool",
		AppVersion:     "0.2.640-test",
	})
	if err != nil {
		t.Fatalf("RecordExternalWinAppImport returned error: %v", err)
	}

	recordPath, err := ExternalWinAppImportRecordPathFromHandle(stateRoot, "org.xnix.external.tool")
	if err != nil {
		t.Fatalf("ExternalWinAppImportRecordPathFromHandle returned error: %v", err)
	}
	if recordPath != filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath)) {
		t.Fatalf("unexpected import record path from handle: %s", recordPath)
	}
	resolvedRecord, resolvedArtifactPath, err := ResolveExternalWinAppImportedArtifactByHandle(stateRoot, "org.xnix.external.tool")
	if err != nil {
		t.Fatalf("ResolveExternalWinAppImportedArtifactByHandle returned error: %v", err)
	}
	if resolvedRecord.ApplicationID != record.ApplicationID ||
		resolvedRecord.RecordSHA256 != record.RecordSHA256 ||
		resolvedArtifactPath != filepath.Join(stateRoot, filepath.FromSlash(record.ArtifactRelativePath)) {
		t.Fatalf("unexpected handle resolution result: record=%#v artifact=%s", resolvedRecord, resolvedArtifactPath)
	}
	for _, unsafeHandle := range []string{"", "../org.xnix.external.tool", "org/xnix/external/tool", "/org.xnix.external.tool"} {
		if _, err := ExternalWinAppImportRecordPathFromHandle(stateRoot, unsafeHandle); err == nil {
			t.Fatalf("expected unsafe handle %q to be rejected", unsafeHandle)
		}
	}
	if _, _, err := ResolveExternalWinAppImportedArtifactByHandle(stateRoot, "org.xnix.external.missing"); err == nil {
		t.Fatalf("expected missing external app handle to fail")
	}
}

func TestRecordExternalWinAppImportRejectsUnsafeInputs(t *testing.T) {
	tempDir := t.TempDir()
	notExecutable := filepath.Join(tempDir, "tool.exe")
	if err := os.WriteFile(notExecutable, []byte("not a PE"), 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}

	if _, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      filepath.Join(tempDir, "state"),
		ExecutablePath: notExecutable,
		AppID:          "org.xnix.external.tool",
		DisplayName:    "External Tool",
		AppVersion:     "0.2.640-test",
	}); err == nil || !strings.Contains(err.Error(), "MZ executable") {
		t.Fatalf("expected MZ validation failure, got %v", err)
	}

	validExecutable := filepath.Join(tempDir, "tool.bat")
	if err := os.WriteFile(validExecutable, []byte{'M', 'Z'}, 0o600); err != nil {
		t.Fatalf("WriteFile executable returned error: %v", err)
	}
	if _, err := RecordExternalWinAppImport(ExternalWinAppImportRequest{
		Version:        "0.2.640-test",
		StateRoot:      filepath.Join(tempDir, "state"),
		ExecutablePath: validExecutable,
		AppID:          "org.xnix.external.tool",
		DisplayName:    "External Tool",
		AppVersion:     "0.2.640-test",
	}); err == nil || !strings.Contains(err.Error(), "safe .exe") {
		t.Fatalf("expected executable name validation failure, got %v", err)
	}
}
