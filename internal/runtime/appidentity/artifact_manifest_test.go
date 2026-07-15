package appidentity

import "testing"

func TestArtifactManifestPreviewPlansGroupsWithoutDownload(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	preview, err := plan.ArtifactManifestPreview()
	if err != nil {
		t.Fatalf("ArtifactManifestPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.artifact_manifest.v1" ||
		preview.RequestType != "artifact-manifest-preview" ||
		preview.ManifestType != "compatibility-artifact-manifest" ||
		preview.Source != "acquisition-preflight-preview+go-runtime-artifact-manifest" ||
		preview.RuntimeMethod != "GetCompatibilityArtifactManifest" ||
		preview.ReadMethod != "GetCompatibilityArtifactManifestPreview" {
		t.Fatalf("unexpected artifact manifest schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.SelectedStrategy != "automatic-managed" ||
		preview.ManifestState != "planned" {
		t.Fatalf("unexpected artifact manifest identity: %#v", preview)
	}
	if !sameStrings(preview.ArtifactGroupIDs, []string{
		"runtime-launch-metadata",
		"local-execution-artifacts",
		"isolated-environment-artifacts",
	}) {
		t.Fatalf("unexpected artifact groups: %#v", preview.ArtifactGroupIDs)
	}
	if preview.ArtifactGroups[0].CacheNamespace != "org.example.ledger.launch-metadata" {
		t.Fatalf("unexpected cache namespace: %#v", preview.ArtifactGroups[0].CacheNamespace)
	}
	if !sameStrings(preview.RequiredPreflightIDs, []string{
		"acquisition-preflight-ready",
		"manifest-signature-verification",
		"artifact-digest-verification",
		"cache-namespace-allocation",
		"rollback-reference",
	}) {
		t.Fatalf("unexpected required preflight: %#v", preview.RequiredPreflightIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.ManifestReady ||
		preview.SignatureVerified ||
		preview.AcquisitionPreflightReady ||
		preview.DownloadEnabled ||
		preview.InstallEnabled ||
		preview.NetworkRequestCreated ||
		preview.ArtifactsDownloaded ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.DesktopShellCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected artifact manifest safety flags: %#v", preview)
	}

	if err := validateNoBackendTerms(preview, "artifact manifest preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}

	// Cache namespace must sanitize unusual identifier characters.
	if got := artifactCacheNamespace("org.example/app id", "local-execution"); got != "org.example-app-id.local-execution" {
		t.Fatalf("unexpected sanitized namespace: %q", got)
	}
}
