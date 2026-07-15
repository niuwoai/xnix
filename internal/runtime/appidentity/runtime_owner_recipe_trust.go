package appidentity

import (
	"fmt"
	"os"
	"path/filepath"
)

const runtimeOwnerRecipeTrustDefaultRegistry = "runtime/recipes/registry.json"

type RuntimeOwnerRecipeTrustPreview struct {
	Version                     string                              `json:"version"`
	SchemaVersion               string                              `json:"schema_version"`
	RequestType                 string                              `json:"request_type"`
	TrustType                   string                              `json:"trust_type"`
	Source                      string                              `json:"source"`
	RuntimeMethod               string                              `json:"runtime_method"`
	ReadMethod                  string                              `json:"read_method"`
	RegistryPath                string                              `json:"registry_path"`
	RegistryName                string                              `json:"registry_name"`
	RecipeCount                 int                                 `json:"recipe_count"`
	SignatureStatuses           []RuntimeOwnerRecipeSignatureStatus `json:"signature_statuses"`
	Checks                      []RuntimeOwnerRecipeTrustCheck      `json:"checks"`
	CheckIDs                    []string                            `json:"check_ids"`
	Counts                      RuntimeOwnerRecipeTrustCounts       `json:"counts"`
	DigestVerified              bool                                `json:"digest_verified"`
	SignedRecipeValidation      bool                                `json:"signed_recipe_validation"`
	DevelopmentRegistry         bool                                `json:"development_registry"`
	UnsignedRecipesPresent      bool                                `json:"unsigned_recipes_present"`
	ProductionRecipeTrustReady  bool                                `json:"production_recipe_trust_ready"`
	RuntimeOwned                bool                                `json:"runtime_owned"`
	GoRuntimeBacked             bool                                `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                                `json:"kde_policy_owner"`
	NetworkRequired             bool                                `json:"network_required"`
	HostRootModified            bool                                `json:"host_root_modified"`
	PrivilegedContainerRequired bool                                `json:"privileged_container_required"`
	BackendDetailsExposed       bool                                `json:"backend_details_exposed"`
	BlockingReasons             []string                            `json:"blocking_reasons"`
	NextRequirements            []string                            `json:"next_requirements"`
	DesktopSafeSummary          string                              `json:"desktop_safe_summary"`
}

type RuntimeOwnerRecipeSignatureStatus struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type RuntimeOwnerRecipeTrustCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type RuntimeOwnerRecipeTrustCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewRuntimeOwnerRecipeTrustPreview(root string) (RuntimeOwnerRecipeTrustPreview, error) {
	if root == "" {
		root = "."
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return RuntimeOwnerRecipeTrustPreview{}, err
	}

	state := inspectRuntimeOwnerRecipeTrust(root)
	checks := runtimeOwnerRecipeTrustChecks(state)
	counts := countRuntimeOwnerRecipeTrustChecks(checks)
	ready := runtimeOwnerRecipeTrustReady(checks)

	preview := RuntimeOwnerRecipeTrustPreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.owner_recipe_trust.v1",
		RequestType:                 "runtime-owner-recipe-trust-preview",
		TrustType:                   "runtime-owner-recipe-trust",
		Source:                      "registry-digests+recipe-signature-status",
		RuntimeMethod:               "GetRuntimeOwnerRecipeTrust",
		ReadMethod:                  "GetRuntimeOwnerRecipeTrustPreview",
		RegistryPath:                runtimeOwnerRecipeTrustDefaultRegistry,
		RegistryName:                state.registryName,
		RecipeCount:                 state.recipeCount,
		SignatureStatuses:           runtimeOwnerRecipeSignatureStatuses(state.signatureCounts),
		Checks:                      checks,
		CheckIDs:                    runtimeOwnerRecipeTrustCheckIDs(checks),
		Counts:                      counts,
		DigestVerified:              state.digestVerified,
		SignedRecipeValidation:      state.signedRecipeValidation,
		DevelopmentRegistry:         state.developmentRegistry,
		UnsignedRecipesPresent:      state.unsignedRecipesPresent,
		ProductionRecipeTrustReady:  ready,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockingReasons:             runtimeOwnerRecipeTrustBlockingReasons(state),
		NextRequirements:            runtimeOwnerRecipeTrustNextRequirements(state),
		DesktopSafeSummary:          runtimeOwnerRecipeTrustSummary(ready, state),
	}
	if err := validateNoBackendTerms(preview, "Runtime owner recipe trust preview"); err != nil {
		return RuntimeOwnerRecipeTrustPreview{}, err
	}
	return preview, nil
}

type runtimeOwnerRecipeTrustState struct {
	registryPresent        bool
	registryReadable       bool
	registryName           string
	recipeCount            int
	digestVerified         bool
	signedRecipeValidation bool
	developmentRegistry    bool
	unsignedRecipesPresent bool
	signatureCounts        map[string]int
	blockingErrors         []string
}

func inspectRuntimeOwnerRecipeTrust(root string) runtimeOwnerRecipeTrustState {
	state := runtimeOwnerRecipeTrustState{
		signatureCounts: make(map[string]int),
	}
	registryPath := runtimeServicePath(root, runtimeOwnerRecipeTrustDefaultRegistry)
	registryData, err := os.ReadFile(registryPath)
	if err != nil {
		state.blockingErrors = append(state.blockingErrors, "recipe registry is not readable")
		return state
	}
	state.registryPresent = true
	state.registryReadable = true

	registry, err := ParseRegistry(registryData)
	if err != nil {
		state.blockingErrors = append(state.blockingErrors, "recipe registry metadata is invalid")
		return state
	}
	state.registryName = registry.RegistryName
	state.recipeCount = len(registry.Recipes)

	allDigestsVerified := true
	for _, entry := range registry.Recipes {
		state.signatureCounts[entry.SignatureStatus]++
		switch entry.SignatureStatus {
		case "development-only":
			state.developmentRegistry = true
		case "unsigned":
			state.unsignedRecipesPresent = true
		}

		recipePath, err := safeRecipePath(filepath.Dir(registryPath), entry.Path)
		if err != nil {
			allDigestsVerified = false
			state.blockingErrors = append(state.blockingErrors, fmt.Sprintf("recipe path is invalid for %s", entry.ID))
			continue
		}
		recipeData, err := os.ReadFile(recipePath)
		if err != nil {
			allDigestsVerified = false
			state.blockingErrors = append(state.blockingErrors, fmt.Sprintf("recipe file is not readable for %s", entry.ID))
			continue
		}
		if err := verifySHA256(recipeData, entry.SHA256); err != nil {
			allDigestsVerified = false
			state.blockingErrors = append(state.blockingErrors, fmt.Sprintf("recipe digest mismatch for %s", entry.ID))
			continue
		}
		recipe, err := ParseRecipe(recipeData)
		if err != nil {
			allDigestsVerified = false
			state.blockingErrors = append(state.blockingErrors, fmt.Sprintf("recipe metadata is invalid for %s", entry.ID))
			continue
		}
		if recipe.ID != entry.ID {
			allDigestsVerified = false
			state.blockingErrors = append(state.blockingErrors, fmt.Sprintf("recipe id mismatch for %s", entry.ID))
		}
	}
	state.digestVerified = state.registryPresent && state.registryReadable && allDigestsVerified
	state.signedRecipeValidation = state.digestVerified && state.recipeCount > 0 && state.signatureCounts["signed"] == state.recipeCount
	return state
}

func runtimeOwnerRecipeTrustChecks(state runtimeOwnerRecipeTrustState) []RuntimeOwnerRecipeTrustCheck {
	signedStatus := "pending"
	if state.signedRecipeValidation {
		signedStatus = "pass"
	} else if !state.digestVerified || state.unsignedRecipesPresent {
		signedStatus = "blocked"
	}
	developmentStatus := "pass"
	if state.developmentRegistry {
		developmentStatus = "pending"
	}
	unsignedStatus := "pass"
	if state.unsignedRecipesPresent {
		unsignedStatus = "blocked"
	}
	return []RuntimeOwnerRecipeTrustCheck{
		runtimeOwnerRecipeTrustCheck("registry-present", runtimeOwnerRecipeTrustPassBlocked(state.registryPresent && state.registryReadable), "A Runtime recipe registry must be available before production ownership."),
		runtimeOwnerRecipeTrustCheck("recipe-digests", runtimeOwnerRecipeTrustPassBlocked(state.digestVerified), "Every recipe digest must match registry metadata."),
		runtimeOwnerRecipeTrustCheck("signed-recipe-validation", signedStatus, "Every production recipe must carry a production-signed status before Runtime ownership is trusted."),
		runtimeOwnerRecipeTrustCheck("development-registry", developmentStatus, "Development-only recipe registries are allowed for local testing but not production ownership."),
		runtimeOwnerRecipeTrustCheck("unsigned-recipes", unsignedStatus, "Unsigned recipes must not be present in a production Runtime owner."),
	}
}

func runtimeOwnerRecipeTrustCheck(id string, status string, summary string) RuntimeOwnerRecipeTrustCheck {
	return RuntimeOwnerRecipeTrustCheck{
		ID:      id,
		Status:  status,
		Summary: summary,
	}
}

func runtimeOwnerRecipeTrustPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func runtimeOwnerRecipeTrustCheckIDs(checks []RuntimeOwnerRecipeTrustCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countRuntimeOwnerRecipeTrustChecks(checks []RuntimeOwnerRecipeTrustCheck) RuntimeOwnerRecipeTrustCounts {
	counts := RuntimeOwnerRecipeTrustCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		case "blocked":
			counts.Blocked++
		}
	}
	return counts
}

func runtimeOwnerRecipeTrustReady(checks []RuntimeOwnerRecipeTrustCheck) bool {
	for _, check := range checks {
		if check.Status != "pass" {
			return false
		}
	}
	return true
}

func runtimeOwnerRecipeSignatureStatuses(signatureCounts map[string]int) []RuntimeOwnerRecipeSignatureStatus {
	statuses := make([]RuntimeOwnerRecipeSignatureStatus, 0, 3)
	for _, status := range []string{"signed", "development-only", "unsigned"} {
		count := signatureCounts[status]
		if count > 0 {
			statuses = append(statuses, RuntimeOwnerRecipeSignatureStatus{Status: status, Count: count})
		}
	}
	return statuses
}

func runtimeOwnerRecipeTrustBlockingReasons(state runtimeOwnerRecipeTrustState) []string {
	reasons := append([]string{}, state.blockingErrors...)
	if state.registryPresent && state.registryReadable && !state.digestVerified {
		reasons = append(reasons, "recipe digests are not verified")
	}
	if !state.signedRecipeValidation {
		reasons = append(reasons, "production signed recipe validation is not enabled")
	}
	if state.developmentRegistry {
		reasons = append(reasons, "registry contains development-only recipes")
	}
	if state.unsignedRecipesPresent {
		reasons = append(reasons, "registry contains unsigned recipes")
	}
	return reasons
}

func runtimeOwnerRecipeTrustNextRequirements(state runtimeOwnerRecipeTrustState) []string {
	requirements := make([]string, 0)
	if !state.registryPresent || !state.registryReadable {
		requirements = append(requirements, "Provide a readable Runtime recipe registry.")
	}
	if !state.digestVerified {
		requirements = append(requirements, "Keep recipe SHA-256 digests aligned with registry metadata.")
	}
	if !state.signedRecipeValidation {
		requirements = append(requirements, "Require production-signed recipe status for every recipe before owner promotion.")
	}
	if state.developmentRegistry {
		requirements = append(requirements, "Replace development-only recipes with production-signed recipes before production ownership.")
	}
	if state.unsignedRecipesPresent {
		requirements = append(requirements, "Remove unsigned recipes from the production Runtime registry.")
	}
	return requirements
}

func runtimeOwnerRecipeTrustSummary(ready bool, state runtimeOwnerRecipeTrustState) string {
	if ready {
		return "Runtime recipes are digest-verified and production-signed for owner promotion."
	}
	if state.developmentRegistry {
		return "Runtime recipes are digest-verified for development, but production owner promotion still needs production-signed recipes."
	}
	return "Runtime recipe trust is not ready for production owner promotion."
}
