package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/artifact"
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
		!preview.Readiness.RecipeTrustDiagnosticsReady ||
		!preview.Readiness.RecipeInstallAllowed ||
		preview.Readiness.RecipeInstallDecision != "allow" ||
		len(preview.Readiness.RecipeTrustBlockingReasons) != 0 {
		t.Fatalf("unexpected install readiness: %#v", preview.Readiness)
	}
	if !sameStrings(preview.PhaseIDs, []string{
		"resolve-artifact-manifest",
		"verify-artifact-digests",
		"consume-artifact-stage-receipt",
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

func TestCompatibilityInstallPlanPreviewConsumesArtifactStageReceipt(t *testing.T) {
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
	receipt := artifact.StageReceipt{
		SchemaVersion: "xnix.runtime.artifact_stage_receipt.v1",
		RecordType:    "compatibility-artifact-stage-receipt",
		Source:        "go-runtime-local-fixture-artifact-staging",
		ApplicationID: "org.example.ledger",
		RelativePath:  "artifact-ledger/receipts/org.example.ledger.json",
		Plan: artifact.Plan{
			ApplicationID: "org.example.ledger",
			Artifacts: []artifact.PlannedArtifact{
				{Group: "runtime-launch-metadata", RefID: "launch.json", CacheNamespace: "org.example.ledger.runtime-launch-metadata", CacheKey: strings.Repeat("a", 64), Required: true},
			},
			StagedKeys: []string{strings.Repeat("a", 64)},
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		CacheRootPathExposed:        false,
		FixtureRootPathExposed:      false,
		NetworkRequired:             false,
		NetworkFetchEnabled:         false,
		PackageManagerInvoked:       false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendLaunchEnabled:        false,
		BackendDetailsExposed:       false,
	}
	receipt.SHA256 = artifactStageReceiptTestDigest(t, receipt)

	preview, err := plan.CompatibilityInstallPlanPreviewWithArtifactReceipt("development", &receipt)
	if err != nil {
		t.Fatalf("CompatibilityInstallPlanPreviewWithArtifactReceipt returned error: %v", err)
	}
	if preview.ArtifactStageReceipt == nil ||
		preview.ArtifactStageReceipt.RelativePath != "artifact-ledger/receipts/org.example.ledger.json" ||
		!preview.ArtifactStageReceipt.RequiredArtifactsStaged ||
		len(preview.ArtifactStageReceipt.BlockingReasons) != 0 {
		t.Fatalf("unexpected artifact stage receipt preview: %#v", preview.ArtifactStageReceipt)
	}
	if !preview.Readiness.ArtifactStageReceiptReady ||
		!preview.Readiness.ArtifactStageDigestVerified ||
		!preview.Readiness.RequiredArtifactsStaged ||
		len(preview.Readiness.ArtifactStageBlockingReasons) != 0 ||
		preview.Source != "artifact-manifest-preview+acquisition-preflight-preview+package-source-preview+state-root-preview+recipe-install-gate+artifact-stage-receipt" {
		t.Fatalf("install preview must consume valid artifact receipt: %#v", preview.Readiness)
	}
	assertNoCompatibilityInstallBackendTerms(t, preview)
}

func TestCompatibilityInstallPlanPreviewFailsClosedForBadArtifactReceipt(t *testing.T) {
	plan, err := NewPlanWithProvenance(Recipe{ID: "org.example.ledger", Name: "Example Ledger", Icon: "office-chart-area", Mode: "automatic"}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	receipt := artifact.StageReceipt{
		SchemaVersion:          "xnix.runtime.artifact_stage_receipt.v1",
		RecordType:             "compatibility-artifact-stage-receipt",
		ApplicationID:          "org.example.ledger",
		RelativePath:           "../escape.json",
		SHA256:                 strings.Repeat("b", 64),
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		HostRootModified:       true,
		BackendDetailsExposed:  true,
		CacheRootPathExposed:   true,
		FixtureRootPathExposed: true,
	}

	preview, err := plan.CompatibilityInstallPlanPreviewWithArtifactReceipt("development", &receipt)
	if err != nil {
		t.Fatalf("CompatibilityInstallPlanPreviewWithArtifactReceipt returned error: %v", err)
	}
	if preview.Readiness.ArtifactStageReceiptReady ||
		preview.Readiness.ArtifactStageDigestVerified ||
		preview.Readiness.RequiredArtifactsStaged ||
		len(preview.Readiness.ArtifactStageBlockingReasons) == 0 {
		t.Fatalf("bad artifact receipt must fail closed: %#v", preview.Readiness)
	}
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
		!preview.Readiness.RecipeTrustDiagnosticsReady ||
		preview.Readiness.RecipeInstallAllowed ||
		preview.Readiness.RecipeInstallDecision != "block" ||
		!sameStrings(preview.Readiness.RecipeTrustBlockingReasons, []string{
			"production signed recipe validation is not enabled",
			"registry contains development-only recipes",
		}) ||
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

func artifactStageReceiptTestDigest(t *testing.T, receipt artifact.StageReceipt) string {
	t.Helper()
	receipt.SHA256 = ""
	data, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent returned error: %v", err)
	}
	data = append(data, '\n')
	return artifactTestSHA256(data)
}

func artifactTestSHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}
