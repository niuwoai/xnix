package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestRuntimeOwnerRecipeTrustPreviewReportsDevelopmentRegistryPending(t *testing.T) {
	preview, err := NewRuntimeOwnerRecipeTrustPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewRuntimeOwnerRecipeTrustPreview returned error: %v", err)
	}

	if preview.Version != "0.2.253" ||
		preview.SchemaVersion != "xnix.runtime.owner_recipe_trust.v1" ||
		preview.RequestType != "runtime-owner-recipe-trust-preview" ||
		preview.TrustType != "runtime-owner-recipe-trust" ||
		preview.Source != "registry-digests+recipe-signature-status" ||
		preview.RuntimeMethod != "GetRuntimeOwnerRecipeTrust" ||
		preview.ReadMethod != "GetRuntimeOwnerRecipeTrustPreview" ||
		preview.RegistryPath != "runtime/recipes/registry.json" {
		t.Fatalf("unexpected Runtime owner recipe trust schema: %#v", preview)
	}
	if preview.RegistryName != "xnix-local-development" ||
		preview.RecipeCount != 1 ||
		!preview.DigestVerified ||
		preview.SignedRecipeValidation ||
		!preview.DevelopmentRegistry ||
		preview.UnsignedRecipesPresent ||
		preview.ProductionRecipeTrustReady {
		t.Fatalf("unexpected development registry trust flags: %#v", preview)
	}
	if len(preview.SignatureStatuses) != 1 ||
		preview.SignatureStatuses[0].Status != "development-only" ||
		preview.SignatureStatuses[0].Count != 1 {
		t.Fatalf("unexpected signature status counts: %#v", preview.SignatureStatuses)
	}
	expectedIDs := []string{"registry-present", "recipe-digests", "signed-recipe-validation", "development-registry", "unsigned-recipes"}
	expectedStatuses := []string{"pass", "pass", "pending", "pending", "pass"}
	if len(preview.Checks) != len(expectedIDs) || len(preview.CheckIDs) != len(expectedIDs) {
		t.Fatalf("unexpected checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	for index, id := range expectedIDs {
		if preview.Checks[index].ID != id ||
			preview.CheckIDs[index] != id ||
			preview.Checks[index].Status != expectedStatuses[index] {
			t.Fatalf("unexpected check at %d: %#v ids=%#v", index, preview.Checks, preview.CheckIDs)
		}
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 3 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected trust counts: %#v", preview.Counts)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected recipe trust safety flags: %#v", preview)
	}
	if len(preview.BlockingReasons) != 2 ||
		preview.BlockingReasons[0] != "production signed recipe validation is not enabled" ||
		preview.BlockingReasons[1] != "registry contains development-only recipes" {
		t.Fatalf("unexpected blocking reasons: %#v", preview.BlockingReasons)
	}
	if preview.DesktopSafeSummary != "Runtime recipes are digest-verified for development, but production owner promotion still needs production-signed recipes." {
		t.Fatalf("unexpected summary: %q", preview.DesktopSafeSummary)
	}
	if err := validateNoBackendTerms(preview, "Runtime owner recipe trust preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRuntimeOwnerRecipeTrustPreviewPassesSignedRegistry(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.253\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	recipeDir := filepath.Join(root, "runtime", "recipes")
	if err := os.MkdirAll(recipeDir, 0o700); err != nil {
		t.Fatalf("MkdirAll recipeDir returned error: %v", err)
	}
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(filepath.Join(recipeDir, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryData := []byte(`{"schema_version":1,"registry_name":"signed-test","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"signed"}]}`)
	if err := os.WriteFile(filepath.Join(recipeDir, "registry.json"), registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerRecipeTrustPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerRecipeTrustPreview returned error: %v", err)
	}

	if preview.RegistryName != "signed-test" ||
		preview.RecipeCount != 1 ||
		!preview.DigestVerified ||
		!preview.SignedRecipeValidation ||
		preview.DevelopmentRegistry ||
		preview.UnsignedRecipesPresent ||
		!preview.ProductionRecipeTrustReady {
		t.Fatalf("signed registry must pass production trust: %#v", preview)
	}
	if len(preview.SignatureStatuses) != 1 ||
		preview.SignatureStatuses[0].Status != "signed" ||
		preview.SignatureStatuses[0].Count != 1 {
		t.Fatalf("unexpected signed status counts: %#v", preview.SignatureStatuses)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 5 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 {
		t.Fatalf("unexpected signed trust counts: %#v", preview.Counts)
	}
	if len(preview.BlockingReasons) != 0 ||
		len(preview.NextRequirements) != 0 {
		t.Fatalf("signed registry must not list blockers: reasons=%#v next=%#v", preview.BlockingReasons, preview.NextRequirements)
	}
	if preview.DesktopSafeSummary != "Runtime recipes are digest-verified and production-signed for owner promotion." {
		t.Fatalf("unexpected signed summary: %q", preview.DesktopSafeSummary)
	}
}

func TestRuntimeOwnerRecipeTrustPreviewBlocksMissingRegistry(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.253\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}

	preview, err := NewRuntimeOwnerRecipeTrustPreview(root)
	if err != nil {
		t.Fatalf("NewRuntimeOwnerRecipeTrustPreview returned error: %v", err)
	}

	if preview.DigestVerified ||
		preview.SignedRecipeValidation ||
		preview.DevelopmentRegistry ||
		preview.ProductionRecipeTrustReady ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("missing registry must not pass recipe trust: %#v", preview)
	}
	if preview.Checks[0].ID != "registry-present" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[1].ID != "recipe-digests" ||
		preview.Checks[1].Status != "blocked" ||
		preview.Checks[2].ID != "signed-recipe-validation" ||
		preview.Checks[2].Status != "blocked" {
		t.Fatalf("missing registry must block required checks: %#v", preview.Checks)
	}
	if preview.Counts.Total != 5 ||
		preview.Counts.Passed != 2 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 3 {
		t.Fatalf("unexpected missing registry counts: %#v", preview.Counts)
	}
	if len(preview.BlockingReasons) < 2 ||
		preview.BlockingReasons[0] != "recipe registry is not readable" {
		t.Fatalf("missing registry must explain blocker: %#v", preview.BlockingReasons)
	}
}
