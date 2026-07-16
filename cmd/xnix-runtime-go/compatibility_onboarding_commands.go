package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type compatibilityOnboardingOptions struct {
	recipe          appidentity.Recipe
	provenance      appidentity.Provenance
	runtimeRoot     string
	stateRoot       string
	artifactReceipt string
	portalOperation string
	snapshotReason  string
	issue           string
	testType        string
}

func runCompatibilityOnboardingChecklistPreview(args []string, stdout io.Writer) error {
	options, err := parseCompatibilityOnboardingChecklistPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(options.recipe, options.provenance)
	if err != nil {
		return err
	}
	artifactReceiptReady := false
	if options.artifactReceipt != "" {
		if _, err := loadArtifactStageReceipt(options.artifactReceipt); err != nil {
			return err
		}
		artifactReceiptReady = true
	}
	preview, err := plan.CompatibilityOnboardingChecklistPreview(appidentity.CompatibilityOnboardingChecklistOptions{
		RuntimeRoot:     options.runtimeRoot,
		StateRoot:       options.stateRoot,
		ArtifactReceipt: artifactReceiptReady,
		PortalOperation: options.portalOperation,
		SnapshotReason:  options.snapshotReason,
		Issue:           options.issue,
		TestType:        options.testType,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseCompatibilityOnboardingChecklistPreviewSource(args []string) (compatibilityOnboardingOptions, error) {
	flags := flag.NewFlagSet("compatibility-onboarding-checklist-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	runtimeRoot := flags.String("root", ".", "Runtime repository root for readiness evidence")
	stateRoot := flags.String("state-root", "", "optional Runtime state root for lifecycle evidence")
	artifactReceipt := flags.String("artifact-receipt", "", "optional artifact stage receipt JSON file to validate as ready evidence")
	portalOperation := flags.String("portal-operation", "file-open", "Portal operation to evaluate")
	snapshotReason := flags.String("snapshot-reason", "before-repair", "snapshot reason to evaluate")
	issue := flags.String("issue", "portal-approval-required", "diagnostic issue label for safe onboarding evidence")
	testType := flags.String("test-type", "smoke", "diagnostic test type for safe onboarding evidence")
	if err := flags.Parse(args); err != nil {
		return compatibilityOnboardingOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return compatibilityOnboardingOptions{}, errors.New("compatibility-onboarding-checklist-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return compatibilityOnboardingOptions{}, errors.New("compatibility-onboarding-checklist-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return compatibilityOnboardingOptions{}, errors.New("compatibility-onboarding-checklist-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return compatibilityOnboardingOptions{}, errors.New("compatibility-onboarding-checklist-preview does not accept positional arguments")
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return compatibilityOnboardingOptions{}, err
	}
	return compatibilityOnboardingOptions{
		recipe:          recipe,
		provenance:      provenance,
		runtimeRoot:     *runtimeRoot,
		stateRoot:       *stateRoot,
		artifactReceipt: *artifactReceipt,
		portalOperation: *portalOperation,
		snapshotReason:  *snapshotReason,
		issue:           *issue,
		testType:        *testType,
	}, nil
}
