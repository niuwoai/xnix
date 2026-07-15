package recipe

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func digestOf(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

// writeRegistry builds a recipe root with one recipe and a registry entry.
func writeRegistry(t *testing.T, registryName, recipeBody, sha, status string) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "org.example.ledger.json"), []byte(recipeBody), 0o600); err != nil {
		t.Fatalf("write recipe: %v", err)
	}
	registry := map[string]any{
		"schema_version": 1,
		"registry_name":  registryName,
		"recipes": []map[string]any{
			{"id": "org.example.ledger", "path": "org.example.ledger.json", "sha256": sha, "signature_status": status},
		},
	}
	data, _ := json.Marshal(registry)
	if err := os.WriteFile(filepath.Join(root, "registry.json"), data, 0o600); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	return root
}

func TestFindDevelopmentOnlyRecipe(t *testing.T) {
	body := `{"id":"org.example.ledger","name":"Ledger"}`
	root := writeRegistry(t, "xnix-local-development", body, digestOf(body), "development-only")

	store, err := NewLocalStore(root, nil)
	if err != nil {
		t.Fatalf("NewLocalStore: %v", err)
	}
	recipe, err := store.Find("org.example.ledger")
	if err != nil {
		t.Fatalf("Find: %v", err)
	}
	if !recipe.Trust.DigestVerified || !recipe.Trust.DevelopmentOnly || recipe.Trust.ProductionTrusted {
		t.Fatalf("dev-only recipe must be verified but never production trusted: %#v", recipe.Trust)
	}
	if recipe.RegistryName != "xnix-local-development" {
		t.Fatalf("registry name not propagated: %q", recipe.RegistryName)
	}
	dev, _ := store.Development()
	if !dev {
		t.Fatalf("development registry must be detected as non-production")
	}
}

func TestDigestMismatchFailsClosed(t *testing.T) {
	body := `{"id":"org.example.ledger","name":"Ledger"}`
	root := writeRegistry(t, "xnix-local-development", body, digestOf("different"), "development-only")

	store, _ := NewLocalStore(root, nil)
	if _, err := store.Find("org.example.ledger"); err == nil {
		t.Fatalf("digest mismatch must fail closed on Find")
	}
	states, err := store.TrustStates()
	if err != nil {
		t.Fatalf("TrustStates: %v", err)
	}
	if len(states) != 1 || !states[0].FailedClosed || states[0].DigestVerified {
		t.Fatalf("digest mismatch must be reported as fail-closed: %#v", states)
	}
	if states[0].Usable() {
		t.Fatalf("fail-closed recipe must not be usable")
	}
}

func TestSignedIsNotProductionTrustedByDefault(t *testing.T) {
	body := `{"id":"org.example.ledger"}`
	root := writeRegistry(t, "xnix-production", body, digestOf(body), "signed")

	// Default verifier does not confirm production signatures.
	store, _ := NewLocalStore(root, nil)
	recipe, err := store.Find("org.example.ledger")
	if err != nil {
		t.Fatalf("Find: %v", err)
	}
	if recipe.Trust.ProductionTrusted {
		t.Fatalf("default verifier must not grant production trust to a declared signature")
	}
	if recipe.Trust.DevelopmentOnly {
		t.Fatalf("signed status must not be classified development-only")
	}

	// A verifier that opts into trusting declared signatures grants production.
	prodStore, _ := NewLocalStore(root, DigestVerifier{TreatSignedAsProduction: true})
	prodRecipe, err := prodStore.Find("org.example.ledger")
	if err != nil {
		t.Fatalf("Find (prod): %v", err)
	}
	if !prodRecipe.Trust.ProductionTrusted {
		t.Fatalf("opt-in verifier must grant production trust: %#v", prodRecipe.Trust)
	}
	dev, _ := prodStore.Development()
	if dev {
		t.Fatalf("production registry must not be flagged development")
	}
}

func TestUnknownSignatureStatusFailsClosed(t *testing.T) {
	body := `{"id":"org.example.ledger"}`
	root := writeRegistry(t, "xnix-local-development", body, digestOf(body), "mystery")

	store, _ := NewLocalStore(root, nil)
	if _, err := store.Find("org.example.ledger"); err == nil {
		t.Fatalf("unknown signature status must fail closed")
	}
}

func TestPathEscapeRefused(t *testing.T) {
	root := t.TempDir()
	registry := `{"schema_version":1,"registry_name":"x-development","recipes":[{"id":"org.example.ledger","path":"../escape.json","sha256":"` + digestOf("x") + `","signature_status":"development-only"}]}`
	if err := os.WriteFile(filepath.Join(root, "registry.json"), []byte(registry), 0o600); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	store, _ := NewLocalStore(root, nil)
	if _, err := store.Find("org.example.ledger"); err == nil {
		t.Fatalf("path escape must be refused")
	}
}

func TestListSortsAndVerifies(t *testing.T) {
	body := `{"id":"org.example.ledger"}`
	root := writeRegistry(t, "xnix-local-development", body, digestOf(body), "development-only")
	store, _ := NewLocalStore(root, nil)
	list, err := store.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 1 || list[0].ID != "org.example.ledger" {
		t.Fatalf("unexpected list: %#v", list)
	}
}

func TestMissingRootFails(t *testing.T) {
	if _, err := NewLocalStore(filepath.Join(t.TempDir(), "nope"), nil); err == nil {
		t.Fatalf("missing recipe root must fail")
	}
	if _, err := NewLocalStore("", nil); err == nil {
		t.Fatalf("empty recipe root must fail")
	}
}
