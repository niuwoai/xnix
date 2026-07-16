package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

type conflictFixtureEntry struct {
	id             string
	file           string
	body           string
	signature      string
	digestOverride string
}

func writeConflictRegistry(t *testing.T, name string, entries []conflictFixtureEntry) string {
	t.Helper()
	root := t.TempDir()
	recipesJSON := ""
	for i, entry := range entries {
		path := filepath.Join(root, entry.file)
		if err := os.WriteFile(path, []byte(entry.body), 0o600); err != nil {
			t.Fatalf("write recipe %s: %v", entry.file, err)
		}
		digest := entry.digestOverride
		if digest == "" {
			sum := sha256.Sum256([]byte(entry.body))
			digest = hex.EncodeToString(sum[:])
		}
		if i > 0 {
			recipesJSON += ","
		}
		recipesJSON += fmt.Sprintf(`{"id":%q,"path":%q,"sha256":%q,"signature_status":%q}`, entry.id, entry.file, digest, entry.signature)
	}
	registryPath := filepath.Join(root, "registry.json")
	registry := fmt.Sprintf(`{"schema_version":1,"registry_name":%q,"recipes":[%s]}`, name, recipesJSON)
	if err := os.WriteFile(registryPath, []byte(registry), 0o600); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	return registryPath
}

func recipeBody(id, version, mode string) string {
	return fmt.Sprintf(`{"id":%q,"name":"Example","version":%q,"icon":"x","mode":%q,"supported_extensions":[".abc"]}`, id, version, mode)
}

func auditGroup(t *testing.T, preview RecipeConflictAuditPreview, id string) RecipeConflictGroup {
	t.Helper()
	for _, group := range preview.Groups {
		if group.ID == id {
			return group
		}
	}
	t.Fatalf("conflict group %q not present; ids=%v", id, preview.GroupIDs)
	return RecipeConflictGroup{}
}

func assertConflictSafe(t *testing.T, preview RecipeConflictAuditPreview) {
	t.Helper()
	if preview.RecipeWritesEnabled || preview.RegistryMigrationEnabled || preview.ArtifactStagingEnabled ||
		preview.NetworkRequired || preview.PackageManagerInvoked || preview.BackendLaunchEnabled ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("conflict audit must keep every unsafe capability disabled: %+v", preview)
	}
	if err := safety.ValidatePayload("recipe conflict audit preview", preview); err != nil {
		t.Fatalf("conflict audit leaks forbidden content: %v", err)
	}
}

func TestRecipeConflictAuditCleanRegistry(t *testing.T) {
	registryPath := writeConflictRegistry(t, "clean", []conflictFixtureEntry{
		{id: "org.example.clean", file: "clean.json", body: recipeBody("org.example.clean", "1.2.0", "automatic"), signature: "signed"},
	})
	preview, err := NewRecipeConflictAuditPreview(RecipeConflictAuditOptions{RegistryPath: registryPath})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "clean" || preview.Counts.Conflicts != 0 || len(preview.Groups) != 0 {
		t.Fatalf("clean registry should have no conflicts: %+v", preview.Counts)
	}
	assertConflictSafe(t, preview)
}

func TestRecipeConflictAuditDetectsEveryConflictClass(t *testing.T) {
	registryPath := writeConflictRegistry(t, "conflict", []conflictFixtureEntry{
		{id: "org.example.app", file: "app1.json", body: recipeBody("org.example.app", "1.0.0", "automatic"), signature: "signed"},
		{id: "org.example.app", file: "app2.json", body: recipeBody("org.example.app", "2.0.0", "automatic"), signature: "signed"},
		{id: "org.example.badmode", file: "bad.json", body: recipeBody("org.example.badmode", "1.0.0", "experimental"), signature: "unsigned"},
		{id: "org.example.drift", file: "drift.json", body: recipeBody("org.example.drift", "1.0.0", "wine"), signature: "development-only", digestOverride: "0000000000000000000000000000000000000000000000000000000000000000"},
	})
	preview, err := NewRecipeConflictAuditPreview(RecipeConflictAuditOptions{RegistryPath: registryPath})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "conflicts-found" {
		t.Fatalf("expected conflicts, got %q", preview.OverallState)
	}
	if dup := auditGroup(t, preview, "duplicate-app-ids"); dup.FindingCount != 2 || !dup.Blocking {
		t.Fatalf("duplicate ids should be blocking with two findings: %+v", dup)
	}
	if stale := auditGroup(t, preview, "stale-recipe-versions"); stale.FindingCount != 1 {
		t.Fatalf("older duplicate version should be stale: %+v", stale)
	}
	if cap := auditGroup(t, preview, "unsupported-capability-claims"); cap.FindingCount != 1 {
		t.Fatalf("experimental mode should be unsupported: %+v", cap)
	}
	if pin := auditGroup(t, preview, "mismatched-package-source-pins"); pin.FindingCount != 1 {
		t.Fatalf("development-only entry should be a pin mismatch: %+v", pin)
	}
	if trust := auditGroup(t, preview, "trust-policy-blockers"); trust.FindingCount != 1 || !trust.Blocking {
		t.Fatalf("unsigned entry should be a blocking trust blocker: %+v", trust)
	}
	if drift := auditGroup(t, preview, "artifact-digest-drift"); drift.FindingCount != 1 || !drift.Blocking {
		t.Fatalf("wrong digest should be blocking digest drift: %+v", drift)
	}
	assertConflictSafe(t, preview)
}

func TestRecipeConflictAuditFlagsMissingRecipeFile(t *testing.T) {
	registryPath := writeConflictRegistry(t, "missing", []conflictFixtureEntry{
		{id: "org.example.present", file: "present.json", body: recipeBody("org.example.present", "1.0.0", "automatic"), signature: "signed"},
	})
	// Point a second entry at a recipe file that does not exist.
	root := filepath.Dir(registryPath)
	registry := fmt.Sprintf(`{"schema_version":1,"registry_name":"missing","recipes":[{"id":"org.example.present","path":"present.json","sha256":%q,"signature_status":"signed"},{"id":"org.example.gone","path":"gone.json","sha256":"%s","signature_status":"signed"}]}`,
		conflictDigest(recipeBody("org.example.present", "1.0.0", "automatic")), "1111111111111111111111111111111111111111111111111111111111111111")
	if err := os.WriteFile(registryPath, []byte(registry), 0o600); err != nil {
		t.Fatalf("rewrite registry: %v", err)
	}
	_ = root
	preview, err := NewRecipeConflictAuditPreview(RecipeConflictAuditOptions{RegistryPath: registryPath})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.Counts.UnreadableRecipes != 1 {
		t.Fatalf("missing recipe file should be counted unreadable: %+v", preview.Counts)
	}
	if drift := auditGroup(t, preview, "artifact-digest-drift"); drift.FindingCount == 0 {
		t.Fatalf("missing recipe file should surface under digest drift: %+v", drift)
	}
	assertConflictSafe(t, preview)
}

func TestRecipeConflictAuditRejectsMissingRegistry(t *testing.T) {
	if _, err := NewRecipeConflictAuditPreview(RecipeConflictAuditOptions{}); err == nil {
		t.Fatalf("expected a missing registry path to be rejected")
	}
	missing := filepath.Join(t.TempDir(), "nope.json")
	if _, err := NewRecipeConflictAuditPreview(RecipeConflictAuditOptions{RegistryPath: missing}); err == nil {
		t.Fatalf("expected an unreadable registry to be rejected")
	}
}

func conflictDigest(body string) string {
	sum := sha256.Sum256([]byte(body))
	return hex.EncodeToString(sum[:])
}
