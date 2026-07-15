package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runCompatibilityInstallPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseCompatibilityInstallPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.CompatibilityInstallPlanPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseCompatibilityInstallPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("compatibility-install-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "development", "install planning mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("compatibility-install-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("compatibility-install-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("compatibility-install-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("compatibility-install-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}
