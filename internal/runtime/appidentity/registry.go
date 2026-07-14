package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Provenance struct {
	Source          string
	RegistryName    string
	DigestVerified  bool
	SignatureStatus string
}

type Registry struct {
	SchemaVersion int             `json:"schema_version"`
	RegistryName  string          `json:"registry_name"`
	Recipes       []RegistryEntry `json:"recipes"`
}

type RegistryEntry struct {
	ID              string `json:"id"`
	Path            string `json:"path"`
	SHA256          string `json:"sha256"`
	SignatureStatus string `json:"signature_status"`
}

func LoadRecipeFromRegistry(registryPath string, recipeRoot string, applicationID string) (Recipe, Provenance, error) {
	if registryPath == "" {
		return Recipe{}, Provenance{}, errors.New("registry path is required")
	}
	if recipeRoot == "" {
		recipeRoot = filepath.Dir(registryPath)
	}
	if !idPattern.MatchString(applicationID) {
		return Recipe{}, Provenance{}, errors.New("application id must be a reverse-DNS identifier")
	}

	registryData, err := os.ReadFile(registryPath)
	if err != nil {
		return Recipe{}, Provenance{}, fmt.Errorf("read registry: %w", err)
	}
	registry, err := ParseRegistry(registryData)
	if err != nil {
		return Recipe{}, Provenance{}, err
	}
	entry, err := registry.Entry(applicationID)
	if err != nil {
		return Recipe{}, Provenance{}, err
	}
	recipePath, err := safeRecipePath(recipeRoot, entry.Path)
	if err != nil {
		return Recipe{}, Provenance{}, err
	}

	recipeData, err := os.ReadFile(recipePath)
	if err != nil {
		return Recipe{}, Provenance{}, fmt.Errorf("read registry recipe: %w", err)
	}
	if err := verifySHA256(recipeData, entry.SHA256); err != nil {
		return Recipe{}, Provenance{}, err
	}
	recipe, err := ParseRecipe(recipeData)
	if err != nil {
		return Recipe{}, Provenance{}, err
	}
	if recipe.ID != entry.ID {
		return Recipe{}, Provenance{}, fmt.Errorf("registry recipe id mismatch: %s != %s", recipe.ID, entry.ID)
	}

	return recipe, Provenance{
		Source:          "registry",
		RegistryName:    registry.RegistryName,
		DigestVerified:  true,
		SignatureStatus: entry.SignatureStatus,
	}, nil
}

func LoadRecipesFromRegistry(registryPath string, recipeRoot string) ([]Recipe, Provenance, error) {
	if registryPath == "" {
		return nil, Provenance{}, errors.New("registry path is required")
	}
	if recipeRoot == "" {
		recipeRoot = filepath.Dir(registryPath)
	}

	registryData, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, Provenance{}, fmt.Errorf("read registry: %w", err)
	}
	registry, err := ParseRegistry(registryData)
	if err != nil {
		return nil, Provenance{}, err
	}

	recipes := make([]Recipe, 0, len(registry.Recipes))
	signatureStatuses := make(map[string]bool, len(registry.Recipes))
	for _, entry := range registry.Recipes {
		recipePath, err := safeRecipePath(recipeRoot, entry.Path)
		if err != nil {
			return nil, Provenance{}, err
		}
		recipeData, err := os.ReadFile(recipePath)
		if err != nil {
			return nil, Provenance{}, fmt.Errorf("read registry recipe: %w", err)
		}
		if err := verifySHA256(recipeData, entry.SHA256); err != nil {
			return nil, Provenance{}, err
		}
		recipe, err := ParseRecipe(recipeData)
		if err != nil {
			return nil, Provenance{}, err
		}
		if recipe.ID != entry.ID {
			return nil, Provenance{}, fmt.Errorf("registry recipe id mismatch: %s != %s", recipe.ID, entry.ID)
		}
		recipes = append(recipes, recipe)
		signatureStatuses[entry.SignatureStatus] = true
	}

	signatureStatus := "mixed"
	if len(signatureStatuses) == 1 {
		for status := range signatureStatuses {
			signatureStatus = status
		}
	}

	return recipes, Provenance{
		Source:          "registry",
		RegistryName:    registry.RegistryName,
		DigestVerified:  true,
		SignatureStatus: signatureStatus,
	}, nil
}

func ParseRegistry(data []byte) (Registry, error) {
	var registry Registry
	if err := json.Unmarshal(data, &registry); err != nil {
		return Registry{}, fmt.Errorf("parse registry JSON: %w", err)
	}
	if err := registry.Validate(); err != nil {
		return Registry{}, err
	}
	return registry, nil
}

func (registry Registry) Validate() error {
	if registry.SchemaVersion != 1 {
		return errors.New("registry schema_version must be 1")
	}
	if !singleLine(registry.RegistryName) {
		return errors.New("registry_name must be a non-empty single-line string")
	}
	if len(registry.Recipes) == 0 {
		return errors.New("registry must include at least one recipe")
	}
	seen := make(map[string]bool, len(registry.Recipes))
	for _, entry := range registry.Recipes {
		if err := entry.Validate(); err != nil {
			return err
		}
		if seen[entry.ID] {
			return fmt.Errorf("duplicate registry recipe id: %s", entry.ID)
		}
		seen[entry.ID] = true
	}
	return nil
}

func (registry Registry) Entry(applicationID string) (RegistryEntry, error) {
	for _, entry := range registry.Recipes {
		if entry.ID == applicationID {
			return entry, nil
		}
	}
	return RegistryEntry{}, fmt.Errorf("application is not registered: %s", applicationID)
}

func (entry RegistryEntry) Validate() error {
	if !idPattern.MatchString(entry.ID) {
		return errors.New("registry recipe id must be a reverse-DNS identifier")
	}
	if _, err := safeRecipePath(".", entry.Path); err != nil {
		return fmt.Errorf("invalid registry recipe path for %s: %w", entry.ID, err)
	}
	if len(entry.SHA256) != 64 {
		return fmt.Errorf("registry recipe sha256 must be 64 hex characters for %s", entry.ID)
	}
	if _, err := hex.DecodeString(entry.SHA256); err != nil {
		return fmt.Errorf("registry recipe sha256 must be hex for %s: %w", entry.ID, err)
	}
	switch entry.SignatureStatus {
	case "development-only", "signed", "unsigned":
	default:
		return fmt.Errorf("unsupported registry signature status for %s: %s", entry.ID, entry.SignatureStatus)
	}
	return nil
}

func safeRecipePath(root string, relativePath string) (string, error) {
	if relativePath == "" {
		return "", errors.New("path is required")
	}
	if filepath.IsAbs(relativePath) {
		return "", errors.New("absolute paths are not allowed")
	}
	cleaned := filepath.Clean(relativePath)
	if cleaned != relativePath {
		return "", errors.New("path must be normalized")
	}
	if cleaned == "." || cleaned == ".." || strings.HasPrefix(cleaned, ".."+string(filepath.Separator)) {
		return "", errors.New("path traversal is not allowed")
	}
	return filepath.Join(root, cleaned), nil
}

func verifySHA256(data []byte, expected string) error {
	sum := sha256.Sum256(data)
	actual := hex.EncodeToString(sum[:])
	if actual != expected {
		return fmt.Errorf("recipe digest mismatch: %s != %s", actual, expected)
	}
	return nil
}
