package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEIntegrationStatusPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "kde-integration-status-preview")
	}
	preview, err := appidentity.NewKDEIntegrationStatusPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEShellIntegrationPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "kde-shell-integration-preview")
	}
	preview, err := appidentity.NewKDEShellIntegrationPlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEApplicationSurfacePreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseKDEApplicationSurfacePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEApplicationSurfacePlanPreviewWithOptions(options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseKDEApplicationSurfacePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.KDEApplicationSurfaceOptions, error) {
	flags := flag.NewFlagSet("kde-application-surface-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.KDEApplicationSurfaceOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.KDEApplicationSurfaceOptions{}, errors.New("kde-application-surface-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.KDEApplicationSurfaceOptions{}, errors.New("kde-application-surface-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.KDEApplicationSurfaceOptions{}, errors.New("kde-application-surface-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.KDEApplicationSurfaceOptions{}, errors.New("kde-application-surface-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.KDEApplicationSurfaceOptions{ActivationRoot: *activationRoot}, err
}
