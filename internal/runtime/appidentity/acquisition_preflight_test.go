package appidentity

import "testing"

func TestAcquisitionPreflightPreviewChainsPackageSourceWithoutDownload(t *testing.T) {
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

	preview, err := plan.AcquisitionPreflightPreview()
	if err != nil {
		t.Fatalf("AcquisitionPreflightPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.acquisition_preflight.v1" ||
		preview.RequestType != "acquisition-preflight-preview" ||
		preview.PreflightType != "compatibility-acquisition-preflight" ||
		preview.Source != "package-source-preview+go-runtime-acquisition-preflight" ||
		preview.RuntimeMethod != "GetCompatibilityAcquisitionPreflight" ||
		preview.ReadMethod != "GetCompatibilityAcquisitionPreflightPreview" {
		t.Fatalf("unexpected acquisition preflight schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.SelectedStrategy != "automatic-managed" ||
		preview.PreflightState != "planned" {
		t.Fatalf("unexpected acquisition preflight identity: %#v", preview)
	}
	if !sameStrings(preview.CheckIDs, []string{
		"package-source-ready",
		"signed-artifact-manifest",
		"runtime-cache-space",
		"network-policy-review",
		"rollback-marker",
	}) {
		t.Fatalf("unexpected acquisition checks: %#v", preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.AcquisitionReady ||
		preview.PackageSourceReady ||
		preview.DownloadEnabled ||
		preview.InstallEnabled ||
		preview.NetworkRequiredForPlanning ||
		preview.NetworkRequestCreated ||
		preview.ArtifactsDownloaded ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.DesktopShellCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected acquisition preflight safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 4 {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	if err := validateNoBackendTerms(preview, "acquisition preflight preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
