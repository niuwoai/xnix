package appidentity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

// RecipeConflictAuditOptions selects the registry to audit. The audit reads the
// registry and recipe files read-only and never writes, migrates, or fetches.
type RecipeConflictAuditOptions struct {
	RegistryPath string
	RecipeRoot   string
}

type RecipeConflictAuditPreview struct {
	SchemaVersion            string                  `json:"schema_version"`
	RequestType              string                  `json:"request_type"`
	AuditType                string                  `json:"audit_type"`
	Source                   string                  `json:"source"`
	Desktop                  string                  `json:"desktop"`
	RuntimeMethod            string                  `json:"runtime_method"`
	ReadMethod               string                  `json:"read_method"`
	RegistryName             string                  `json:"registry_name"`
	OverallState             string                  `json:"overall_state"`
	RecipeCount              int                     `json:"recipe_count"`
	Groups                   []RecipeConflictGroup   `json:"groups"`
	GroupIDs                 []string                `json:"group_ids"`
	Findings                 []RecipeConflictFinding `json:"findings"`
	Counts                   RecipeConflictCounts    `json:"counts"`
	ResolutionHints          []string                `json:"resolution_hints"`
	NextSafeReadOnlyChecks   []string                `json:"next_safe_read_only_checks"`
	RuntimeOwned             bool                    `json:"runtime_owned"`
	GoRuntimeBacked          bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                    `json:"kde_policy_owner"`
	UserVisible              bool                    `json:"user_visible"`
	ReviewOnly               bool                    `json:"review_only"`
	RecipeWritesEnabled      bool                    `json:"recipe_writes_enabled"`
	RegistryMigrationEnabled bool                    `json:"registry_migration_enabled"`
	ArtifactStagingEnabled   bool                    `json:"artifact_staging_enabled"`
	NetworkRequired          bool                    `json:"network_required"`
	PackageManagerInvoked    bool                    `json:"package_manager_invoked"`
	BackendLaunchEnabled     bool                    `json:"backend_launch_enabled"`
	HostRootModified         bool                    `json:"host_root_modified"`
	BackendDetailsExposed    bool                    `json:"backend_details_exposed"`
	BlockedActions           []string                `json:"blocked_actions"`
	DesktopSafeSummary       string                  `json:"desktop_safe_summary"`
}

type RecipeConflictGroup struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	FindingCount   int      `json:"finding_count"`
	ApplicationIDs []string `json:"application_ids"`
	ResolutionHint string   `json:"resolution_hint"`
	Blocking       bool     `json:"blocking"`
}

type RecipeConflictFinding struct {
	GroupID        string `json:"group_id"`
	ApplicationID  string `json:"application_id"`
	EntryIndex     int    `json:"entry_index"`
	DetailCode     string `json:"detail_code"`
	ResolutionHint string `json:"resolution_hint"`
	Blocking       bool   `json:"blocking"`
}

type RecipeConflictCounts struct {
	Recipes                    int `json:"recipes"`
	Conflicts                  int `json:"conflicts"`
	DuplicateAppIDs            int `json:"duplicate_app_ids"`
	StaleVersions              int `json:"stale_versions"`
	UnsupportedCapabilities    int `json:"unsupported_capabilities"`
	PackageSourcePinMismatches int `json:"package_source_pin_mismatches"`
	TrustBlockers              int `json:"trust_blockers"`
	DigestDrift                int `json:"digest_drift"`
	UnreadableRecipes          int `json:"unreadable_recipes"`
}

type recipeConflictEntry struct {
	index           int
	id              string
	signatureStatus string
	declaredDigest  string
	mode            string
	version         string
	readable        bool
	digestMatches   bool
}

const (
	conflictGroupDuplicateIDs     = "duplicate-app-ids"
	conflictGroupStaleVersions    = "stale-recipe-versions"
	conflictGroupUnsupportedCap   = "unsupported-capability-claims"
	conflictGroupPackageSourcePin = "mismatched-package-source-pins"
	conflictGroupTrustPolicy      = "trust-policy-blockers"
	conflictGroupDigestDrift      = "artifact-digest-drift"
)

// recipeConflictKnownModes mirrors the modes Recipe.Validate accepts. Any other
// declared mode is surfaced as an unsupported capability claim.
var recipeConflictKnownModes = map[string]bool{"automatic": true, "wine": true, "vm": true}

// NewRecipeConflictAuditPreview audits a recipe registry for version conflicts,
// package-source pins, capability constraints, trust blockers, and artifact
// digest drift before any update or install is attempted. It parses the registry
// leniently (so it can surface duplicate ids and digest drift that the strict
// loader would reject), and never edits recipes, stages artifacts, fetches
// packages, calls a package manager, starts a backend, or mutates the host.
func NewRecipeConflictAuditPreview(options RecipeConflictAuditOptions) (RecipeConflictAuditPreview, error) {
	registryPath := strings.TrimSpace(options.RegistryPath)
	if registryPath == "" {
		return RecipeConflictAuditPreview{}, errors.New("recipe-conflict-audit-preview requires a registry path")
	}
	recipeRoot := strings.TrimSpace(options.RecipeRoot)
	if recipeRoot == "" {
		recipeRoot = filepath.Dir(registryPath)
	}

	registryData, err := os.ReadFile(registryPath)
	if err != nil {
		return RecipeConflictAuditPreview{}, fmt.Errorf("read registry for conflict audit: %w", err)
	}
	var registry Registry
	if err := json.Unmarshal(registryData, &registry); err != nil {
		return RecipeConflictAuditPreview{}, fmt.Errorf("parse registry for conflict audit: %w", err)
	}
	if registry.SchemaVersion != 1 {
		return RecipeConflictAuditPreview{}, errors.New("recipe-conflict-audit-preview requires registry schema_version 1")
	}

	entries := make([]recipeConflictEntry, 0, len(registry.Recipes))
	for index, registryEntry := range registry.Recipes {
		entries = append(entries, recipeConflictAuditEntry(index, registryEntry, recipeRoot))
	}

	var findings []RecipeConflictFinding
	findings = append(findings, recipeConflictDuplicateFindings(entries)...)
	findings = append(findings, recipeConflictVersionFindings(entries)...)
	findings = append(findings, recipeConflictPerEntryFindings(entries)...)
	sort.SliceStable(findings, func(i, j int) bool {
		if findings[i].GroupID != findings[j].GroupID {
			return findings[i].GroupID < findings[j].GroupID
		}
		return findings[i].EntryIndex < findings[j].EntryIndex
	})

	counts := recipeConflictCounts(entries, findings)
	groups := recipeConflictGroups(findings)
	overallState := "clean"
	if len(findings) > 0 {
		overallState = "conflicts-found"
	}

	preview := RecipeConflictAuditPreview{
		SchemaVersion:   "xnix.runtime.recipe_conflict_audit.v1",
		RequestType:     "recipe-conflict-audit-preview",
		AuditType:       "recipe-conflict-and-pin-audit",
		Source:          "registry+recipe-trust+package-source+backend-capability+artifact-digest",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetRecipeConflictAudit",
		ReadMethod:      "GetRecipeConflictAuditPreview",
		RegistryName:    registry.RegistryName,
		OverallState:    overallState,
		RecipeCount:     len(entries),
		Groups:          groups,
		GroupIDs:        recipeConflictGroupIDs(groups),
		Findings:        findings,
		Counts:          counts,
		ResolutionHints: recipeConflictResolutionHints(groups),
		NextSafeReadOnlyChecks: []string{
			"review the recipe registry entries and declared digests",
			"review the recipe trust policy for each application",
			"review the compatibility install plan before any update",
		},
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		UserVisible:              true,
		ReviewOnly:               true,
		RecipeWritesEnabled:      false,
		RegistryMigrationEnabled: false,
		ArtifactStagingEnabled:   false,
		NetworkRequired:          false,
		PackageManagerInvoked:    false,
		BackendLaunchEnabled:     false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		BlockedActions: []string{
			"edit a recipe from the conflict audit",
			"migrate the registry from the conflict audit",
			"stage an artifact from the conflict audit",
			"fetch a package from the conflict audit",
			"call a host package manager from the conflict audit",
			"start a compatibility backend from the conflict audit",
			"mutate host root during the conflict audit",
		},
		DesktopSafeSummary: recipeConflictSummary(overallState, counts),
	}
	if err := validateNoBackendTerms(preview, "recipe conflict audit preview"); err != nil {
		return RecipeConflictAuditPreview{}, err
	}
	if err := safety.ValidatePayload("recipe conflict audit preview", preview); err != nil {
		return RecipeConflictAuditPreview{}, err
	}
	return preview, nil
}

func recipeConflictAuditEntry(index int, registryEntry RegistryEntry, recipeRoot string) recipeConflictEntry {
	entry := recipeConflictEntry{
		index:           index,
		id:              registryEntry.ID,
		signatureStatus: registryEntry.SignatureStatus,
		declaredDigest:  registryEntry.SHA256,
	}
	recipePath, err := safeRecipePath(recipeRoot, registryEntry.Path)
	if err != nil {
		return entry
	}
	recipeData, err := os.ReadFile(recipePath)
	if err != nil {
		return entry
	}
	entry.readable = true
	sum := sha256.Sum256(recipeData)
	entry.digestMatches = hex.EncodeToString(sum[:]) == registryEntry.SHA256
	recipe, err := ParseRecipe(recipeData)
	if err != nil {
		entry.readable = false
		return entry
	}
	entry.mode = recipe.Mode
	entry.version = strings.TrimSpace(recipe.Version)
	return entry
}

func recipeConflictDuplicateFindings(entries []recipeConflictEntry) []RecipeConflictFinding {
	counts := map[string]int{}
	for _, entry := range entries {
		counts[entry.id]++
	}
	var findings []RecipeConflictFinding
	for _, entry := range entries {
		if counts[entry.id] > 1 {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupDuplicateIDs,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "duplicate-registry-entry",
				ResolutionHint: "Remove or rename the duplicate registry entry before installing.",
				Blocking:       true,
			})
		}
	}
	return findings
}

func recipeConflictVersionFindings(entries []recipeConflictEntry) []RecipeConflictFinding {
	maxVersion := map[string]string{}
	for _, entry := range entries {
		if !entry.readable || entry.version == "" {
			continue
		}
		current, ok := maxVersion[entry.id]
		if !ok {
			maxVersion[entry.id] = entry.version
			continue
		}
		if cmp, err := compareSemanticVersions(entry.version, current); err == nil && cmp > 0 {
			maxVersion[entry.id] = entry.version
		}
	}

	var findings []RecipeConflictFinding
	for _, entry := range entries {
		if !entry.readable {
			continue
		}
		if entry.version == "" {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupStaleVersions,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "missing-version",
				ResolutionHint: "Add an explicit semantic version to the recipe before promoting it.",
				Blocking:       false,
			})
			continue
		}
		top, ok := maxVersion[entry.id]
		if !ok {
			continue
		}
		if cmp, err := compareSemanticVersions(entry.version, top); err == nil && cmp < 0 {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupStaleVersions,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "superseded-version",
				ResolutionHint: "Retire the superseded recipe version to avoid an ambiguous install target.",
				Blocking:       false,
			})
		} else if err != nil {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupStaleVersions,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "unparseable-version",
				ResolutionHint: "Use a semantic version so update ordering is deterministic.",
				Blocking:       false,
			})
		}
	}
	return findings
}

func recipeConflictPerEntryFindings(entries []recipeConflictEntry) []RecipeConflictFinding {
	var findings []RecipeConflictFinding
	for _, entry := range entries {
		if !entry.readable {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupDigestDrift,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "missing-or-unreadable-recipe",
				ResolutionHint: "Restore the recipe file so its digest can be verified before installing.",
				Blocking:       true,
			})
			continue
		}
		if !entry.digestMatches {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupDigestDrift,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "recipe-digest-drift",
				ResolutionHint: "Re-pin the registry digest to the reviewed recipe content before installing.",
				Blocking:       true,
			})
		}
		if !recipeConflictKnownModes[entry.mode] {
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupUnsupportedCap,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "unsupported-compatibility-mode",
				ResolutionHint: "Choose a supported compatibility mode before installing.",
				Blocking:       false,
			})
		}
		switch entry.signatureStatus {
		case "unsigned":
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupTrustPolicy,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "untrusted-unsigned-recipe",
				ResolutionHint: "Sign the recipe or lower the trust requirement before installing.",
				Blocking:       true,
			})
		case "development-only":
			findings = append(findings, RecipeConflictFinding{
				GroupID:        conflictGroupPackageSourcePin,
				ApplicationID:  entry.id,
				EntryIndex:     entry.index,
				DetailCode:     "development-pinned-source",
				ResolutionHint: "Promote the source pin to a signed channel before a production install.",
				Blocking:       false,
			})
		}
	}
	return findings
}

func recipeConflictCounts(entries []recipeConflictEntry, findings []RecipeConflictFinding) RecipeConflictCounts {
	counts := RecipeConflictCounts{Recipes: len(entries), Conflicts: len(findings)}
	for _, entry := range entries {
		if !entry.readable {
			counts.UnreadableRecipes++
		}
	}
	for _, finding := range findings {
		switch finding.GroupID {
		case conflictGroupDuplicateIDs:
			counts.DuplicateAppIDs++
		case conflictGroupStaleVersions:
			counts.StaleVersions++
		case conflictGroupUnsupportedCap:
			counts.UnsupportedCapabilities++
		case conflictGroupPackageSourcePin:
			counts.PackageSourcePinMismatches++
		case conflictGroupTrustPolicy:
			counts.TrustBlockers++
		case conflictGroupDigestDrift:
			counts.DigestDrift++
		}
	}
	return counts
}

func recipeConflictGroups(findings []RecipeConflictFinding) []RecipeConflictGroup {
	order := []string{
		conflictGroupDuplicateIDs,
		conflictGroupStaleVersions,
		conflictGroupUnsupportedCap,
		conflictGroupPackageSourcePin,
		conflictGroupTrustPolicy,
		conflictGroupDigestDrift,
	}
	byID := map[string]*RecipeConflictGroup{}
	for _, id := range order {
		byID[id] = &RecipeConflictGroup{ID: id, Title: recipeConflictGroupTitle(id), ResolutionHint: recipeConflictGroupHint(id)}
	}
	for _, finding := range findings {
		group := byID[finding.GroupID]
		if group == nil {
			continue
		}
		group.FindingCount++
		group.ApplicationIDs = append(group.ApplicationIDs, finding.ApplicationID)
		if finding.Blocking {
			group.Blocking = true
		}
	}
	var groups []RecipeConflictGroup
	for _, id := range order {
		group := byID[id]
		if group.FindingCount == 0 {
			continue
		}
		group.ApplicationIDs = recipeConflictUnique(group.ApplicationIDs)
		groups = append(groups, *group)
	}
	return groups
}

func recipeConflictGroupTitle(id string) string {
	switch id {
	case conflictGroupDuplicateIDs:
		return "Duplicate application ids"
	case conflictGroupStaleVersions:
		return "Stale recipe versions"
	case conflictGroupUnsupportedCap:
		return "Unsupported capability claims"
	case conflictGroupPackageSourcePin:
		return "Mismatched package-source pins"
	case conflictGroupTrustPolicy:
		return "Trust-policy blockers"
	case conflictGroupDigestDrift:
		return "Artifact digest drift"
	default:
		return id
	}
}

func recipeConflictGroupHint(id string) string {
	switch id {
	case conflictGroupDuplicateIDs:
		return "Keep exactly one registry entry per application id."
	case conflictGroupStaleVersions:
		return "Retire superseded versions and give every recipe a semantic version."
	case conflictGroupUnsupportedCap:
		return "Use a supported compatibility mode before installing."
	case conflictGroupPackageSourcePin:
		return "Promote development-pinned sources to signed channels for production."
	case conflictGroupTrustPolicy:
		return "Sign recipes or adjust the trust requirement before installing."
	case conflictGroupDigestDrift:
		return "Re-pin registry digests to the reviewed recipe content."
	default:
		return ""
	}
}

func recipeConflictResolutionHints(groups []RecipeConflictGroup) []string {
	hints := []string{}
	for _, group := range groups {
		if group.ResolutionHint != "" {
			hints = append(hints, group.ResolutionHint)
		}
	}
	if len(hints) == 0 {
		hints = append(hints, "No recipe conflicts were found; the registry is safe to review for install.")
	}
	return hints
}

func recipeConflictGroupIDs(groups []RecipeConflictGroup) []string {
	ids := make([]string, 0, len(groups))
	for _, group := range groups {
		ids = append(ids, group.ID)
	}
	return ids
}

func recipeConflictSummary(overallState string, counts RecipeConflictCounts) string {
	if overallState == "clean" {
		return "The recipe registry has no id, version, capability, pin, trust, or digest conflicts; nothing was edited or installed."
	}
	return "The recipe registry has conflicts to review before any update or install; nothing was edited, staged, fetched, or installed."
}

func recipeConflictUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
