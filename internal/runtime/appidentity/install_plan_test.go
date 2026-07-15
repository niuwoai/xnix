package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCompatibilityInstallPlanPreviewAggregatesReadinessWithoutInstalling(t *testing.T) {
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}

	preview, err := plan.CompatibilityInstallPlanPreview("development")
	if err != nil {
		t.Fatalf("CompatibilityInstallPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.compatibility_install_plan.v1" ||
		preview.RequestType != "compatibility-install-preview" ||
		preview.PlanType != "compatibility-install-plan" ||
		preview.RuntimeMethod != "GetCompatibilityInstallPlan" ||
		preview.ReadMethod != "GetCompatibilityInstallPlanPreview" ||
		preview.Environment != "development" ||
		preview.SelectedStrategy != "automatic-managed" ||
		preview.InstallState != "planned" {
		t.Fatalf("unexpected compatibility install schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.RequestedMode != "automatic" {
		t.Fatalf("unexpected application identity: %#v", preview.Application)
	}
	if preview.Readiness.ArtifactManifestReady ||
		preview.Readiness.ArtifactSignatureVerified ||
		preview.Readiness.AcquisitionReady ||
		preview.Readiness.PackageSourceReady ||
		preview.Readiness.StateRootAllocated ||
		!preview.Readiness.RecipeInstallAllowed ||
		preview.Readiness.RecipeInstallDecision != "allow" {
		t.Fatalf("unexpected install readiness: %#v", preview.Readiness)
	}
	if !sameStrings(preview.PhaseIDs, []string{
		"resolve-artifact-manifest",
		"verify-artifact-digests",
		"prepare-package-source",
		"allocate-application-state",
		"review-recipe-install-gate",
		"stage-desktop-integration",
		"enable-launch-binding",
	}) {
		t.Fatalf("unexpected install phases: %#v", preview.PhaseIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.InstallReady ||
		preview.DesktopActivationReady ||
		preview.DownloadEnabled ||
		preview.InstallEnabled ||
		preview.NetworkRequestCreated ||
		preview.ArtifactsDownloaded ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.DesktopShellCommandExposed ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected install safety flags: %#v", preview)
	}
	assertNoCompatibilityInstallBackendTerms(t, preview)
}

func TestCompatibilityInstallPlanPreviewBlocksProductionForDevelopmentRecipe(t *testing.T) {
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}

	preview, err := plan.CompatibilityInstallPlanPreview("production")
	if err != nil {
		t.Fatalf("CompatibilityInstallPlanPreview returned error: %v", err)
	}
	if preview.Environment != "production" ||
		preview.InstallGate.Decision != "block" ||
		preview.Readiness.RecipeInstallAllowed ||
		preview.Readiness.RecipeInstallDecision != "block" ||
		len(preview.InstallGate.BlockingReasons) != 2 {
		t.Fatalf("production install plan must block development recipes: %#v", preview.InstallGate)
	}
	assertNoCompatibilityInstallBackendTerms(t, preview)
}

func assertNoCompatibilityInstallBackendTerms(t *testing.T, value any) {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("compatibility install preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
