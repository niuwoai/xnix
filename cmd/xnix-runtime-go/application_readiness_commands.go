package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/artifact"
)

type applicationReadinessOptions struct {
	recipe          appidentity.Recipe
	provenance      appidentity.Provenance
	environment     string
	runtimeRoot     string
	stateRoot       string
	artifactReceipt string
	portalOperation string
	snapshotReason  string
}

func runApplicationReadinessPreview(args []string, stdout io.Writer) error {
	options, err := parseApplicationReadinessPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(options.recipe, options.provenance)
	if err != nil {
		return err
	}
	var receipt *artifact.StageReceipt
	if options.artifactReceipt != "" {
		loaded, err := loadArtifactStageReceipt(options.artifactReceipt)
		if err != nil {
			return err
		}
		receipt = &loaded
	}
	preview, err := plan.ApplicationReadinessPreview(appidentity.ApplicationReadinessOptions{
		Environment:     options.environment,
		RuntimeRoot:     options.runtimeRoot,
		StateRoot:       options.stateRoot,
		ArtifactReceipt: receipt,
		PortalOperation: options.portalOperation,
		SnapshotReason:  options.snapshotReason,
		WriteMethod:     "Launch",
	})
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseApplicationReadinessPreviewSource(args []string) (applicationReadinessOptions, error) {
	flags := flag.NewFlagSet("application-readiness-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	environment := flags.String("mode", "development", "readiness mode: production or development")
	runtimeRoot := flags.String("root", ".", "Runtime repository root for versioned gate reads")
	stateRoot := flags.String("state-root", "", "optional Runtime state root for backend lifecycle evidence")
	artifactReceipt := flags.String("artifact-receipt", "", "optional artifact stage receipt JSON file")
	portalOperation := flags.String("portal-operation", "file-open", "Portal operation to evaluate")
	snapshotReason := flags.String("snapshot-reason", "before-repair", "snapshot reason to evaluate")
	if err := flags.Parse(args); err != nil {
		return applicationReadinessOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return applicationReadinessOptions{}, errors.New("application-readiness-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return applicationReadinessOptions{}, errors.New("application-readiness-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return applicationReadinessOptions{}, errors.New("application-readiness-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return applicationReadinessOptions{}, errors.New("application-readiness-preview does not accept positional arguments")
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return applicationReadinessOptions{}, err
	}
	return applicationReadinessOptions{
		recipe:          recipe,
		provenance:      provenance,
		environment:     *environment,
		runtimeRoot:     *runtimeRoot,
		stateRoot:       *stateRoot,
		artifactReceipt: *artifactReceipt,
		portalOperation: *portalOperation,
		snapshotReason:  *snapshotReason,
	}, nil
}
