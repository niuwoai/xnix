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
	loaded, err := LoadExternalWinAppImportRecord(filepath.Join(stateRoot, filepath.FromSlash(record.RecordRelativePath)))
	if err != nil {
		t.Fatalf("LoadExternalWinAppImportRecord returned error: %v", err)
	}
	if loaded.RecordSHA256 != record.RecordSHA256 || loaded.ArtifactSHA256 != record.ArtifactSHA256 {
		t.Fatalf("loaded record mismatch: %#v != %#v", loaded, record)
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
