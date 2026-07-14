package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadRecipeFromRegistryVerifiesDigestAndProvenance(t *testing.T) {
	root := t.TempDir()
	recipeData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	sum := sha256.Sum256(recipeData)
	digest := hex.EncodeToString(sum[:])

	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), recipeData, 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + digest + `","signature_status":"development-only"}]}`)
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	recipe, provenance, err := LoadRecipeFromRegistry(registryPath, "", "org.example.ledger")
	if err != nil {
		t.Fatalf("LoadRecipeFromRegistry returned error: %v", err)
	}
	if recipe.ID != "org.example.ledger" {
		t.Fatalf("recipe ID = %q", recipe.ID)
	}
	if provenance.Source != "registry" || provenance.RegistryName != "test-registry" || !provenance.DigestVerified {
		t.Fatalf("unexpected provenance: %#v", provenance)
	}

	plan, err := NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	if plan.RecipeSource != "registry" || plan.RegistryName != "test-registry" || !plan.RecipeDigestVerified {
		t.Fatalf("unexpected plan provenance: %#v", plan)
	}
}

func TestLoadRecipeFromRegistryRejectsDigestMismatch(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic"}`), 0o600); err != nil {
		t.Fatalf("WriteFile recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"0000000000000000000000000000000000000000000000000000000000000000","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	if _, _, err := LoadRecipeFromRegistry(registryPath, "", "org.example.ledger"); err == nil {
		t.Fatal("LoadRecipeFromRegistry accepted a digest mismatch")
	}
}

func TestLoadRecipesFromRegistryLoadsAllRecipes(t *testing.T) {
	root := t.TempDir()
	ledgerData := []byte(`{"id":"org.example.ledger","name":"Example Ledger","icon":"office-chart-area","mode":"automatic","supported_extensions":[".abc"]}`)
	notesData := []byte(`{"id":"org.example.notes","name":"Example Notes","icon":"accessories-text-editor","mode":"automatic","supported_extensions":[".txt"]}`)
	ledgerSum := sha256.Sum256(ledgerData)
	notesSum := sha256.Sum256(notesData)
	ledgerDigest := hex.EncodeToString(ledgerSum[:])
	notesDigest := hex.EncodeToString(notesSum[:])
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), ledgerData, 0o600); err != nil {
		t.Fatalf("WriteFile ledger recipe returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "org.example.notes.json"), notesData, 0o600); err != nil {
		t.Fatalf("WriteFile notes recipe returned error: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"org.example.ledger.json","sha256":"` + ledgerDigest + `","signature_status":"development-only"},{"id":"org.example.notes","path":"org.example.notes.json","sha256":"` + notesDigest + `","signature_status":"development-only"}]}`)
	if err := os.WriteFile(registryPath, registryData, 0o600); err != nil {
		t.Fatalf("WriteFile registry returned error: %v", err)
	}

	recipes, provenance, err := LoadRecipesFromRegistry(registryPath, "")
	if err != nil {
		t.Fatalf("LoadRecipesFromRegistry returned error: %v", err)
	}
	if len(recipes) != 2 {
		t.Fatalf("recipe count = %d, want 2", len(recipes))
	}
	if recipes[0].ID != "org.example.ledger" || recipes[1].ID != "org.example.notes" {
		t.Fatalf("unexpected recipes: %#v", recipes)
	}
	if provenance.Source != "registry" || provenance.RegistryName != "test-registry" ||
		!provenance.DigestVerified || provenance.SignatureStatus != "development-only" {
		t.Fatalf("unexpected provenance: %#v", provenance)
	}
}

func TestParseRegistryRejectsUnsafeRecipePaths(t *testing.T) {
	registryData := []byte(`{"schema_version":1,"registry_name":"test-registry","recipes":[{"id":"org.example.ledger","path":"../org.example.ledger.json","sha256":"0000000000000000000000000000000000000000000000000000000000000000","signature_status":"development-only"}]}`)
	if _, err := ParseRegistry(registryData); err == nil {
		t.Fatal("ParseRegistry accepted a path traversal")
	}
}
