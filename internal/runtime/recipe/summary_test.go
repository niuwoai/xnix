package recipe

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// writeMultiRegistry builds a recipe root with several recipes and a registry
// describing them. Each spec is {id, body, sha, status}.
func writeMultiRegistry(t *testing.T, registryName string, specs [][4]string) string {
	t.Helper()
	root := t.TempDir()
	recipes := make([]map[string]any, 0, len(specs))
	for _, spec := range specs {
		id, body, sha, status := spec[0], spec[1], spec[2], spec[3]
		path := id + ".json"
		if err := os.WriteFile(filepath.Join(root, path), []byte(body), 0o600); err != nil {
			t.Fatalf("write recipe: %v", err)
		}
		recipes = append(recipes, map[string]any{"id": id, "path": path, "sha256": sha, "signature_status": status})
	}
	registry := map[string]any{"schema_version": 1, "registry_name": registryName, "recipes": recipes}
	data, _ := json.Marshal(registry)
	if err := os.WriteFile(filepath.Join(root, "registry.json"), data, 0o600); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	return root
}

func TestTrustSummaryBucketsRecipes(t *testing.T) {
	prodBody := `{"id":"org.example.prod"}`
	devBody := `{"id":"org.example.dev"}`
	badBody := `{"id":"org.example.bad"}`
	root := writeMultiRegistry(t, "xnix-production", [][4]string{
		{"org.example.prod", prodBody, digestOf(prodBody), "signed"},
		{"org.example.dev", devBody, digestOf(devBody), "development-only"},
		{"org.example.bad", badBody, digestOf("different"), "development-only"}, // digest mismatch
	})

	// A verifier that trusts declared signatures, so "signed" is production.
	store, err := NewLocalStore(root, DigestVerifier{TreatSignedAsProduction: true})
	if err != nil {
		t.Fatalf("NewLocalStore: %v", err)
	}
	summary, err := store.TrustSummary()
	if err != nil {
		t.Fatalf("TrustSummary: %v", err)
	}
	if summary.RegistryName != "xnix-production" || summary.Development {
		t.Fatalf("unexpected registry identity: %#v", summary)
	}
	if summary.Total != 3 ||
		!equalStrings(summary.ProductionTrusted, []string{"org.example.prod"}) ||
		!equalStrings(summary.DevelopmentOnly, []string{"org.example.dev"}) ||
		!equalStrings(summary.FailedClosed, []string{"org.example.bad"}) {
		t.Fatalf("unexpected trust buckets: %#v", summary)
	}
	if summary.AllUsable {
		t.Fatalf("a fail-closed recipe must make the summary not all-usable")
	}
}

func TestTrustSummaryDevelopmentRegistry(t *testing.T) {
	body := `{"id":"org.example.ledger"}`
	root := writeMultiRegistry(t, "xnix-local-development", [][4]string{
		{"org.example.ledger", body, digestOf(body), "development-only"},
	})
	store, _ := NewLocalStore(root, nil)
	summary, err := store.TrustSummary()
	if err != nil {
		t.Fatalf("TrustSummary: %v", err)
	}
	if !summary.Development || !summary.AllUsable || len(summary.ProductionTrusted) != 0 || len(summary.DevelopmentOnly) != 1 {
		t.Fatalf("development registry summary wrong: %#v", summary)
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
