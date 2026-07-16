package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runApplicationUpgradeImpactPreview(args []string, stdout io.Writer) error {
	currentRecipe, currentProvenance, candidateRecipe, candidateProvenance, mode, err := parseApplicationUpgradeImpactPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewApplicationUpgradeImpactPreview(currentRecipe, currentProvenance, candidateRecipe, candidateProvenance, mode)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseApplicationUpgradeImpactPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("application-upgrade-impact-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to the current Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing current recipe files; defaults to the current registry directory")
	applicationID := flags.String("app", "", "current application id to load from the current recipe registry")
	candidateRegistryPath := flags.String("candidate-registry", "", "path to the candidate Xnix recipe registry JSON file")
	candidateRecipeRoot := flags.String("candidate-recipe-root", "", "directory containing candidate recipe files; defaults to the candidate registry directory")
	candidateApplicationID := flags.String("candidate-app", "", "candidate application id; defaults to --app")
	mode := flags.String("mode", "development", "upgrade impact review mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *registryPath == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("application-upgrade-impact-preview requires --registry")
	}
	if *candidateRegistryPath == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("application-upgrade-impact-preview requires --candidate-registry")
	}
	if *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("application-upgrade-impact-preview requires --app")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("application-upgrade-impact-preview does not accept positional arguments")
	}
	candidateID := *candidateApplicationID
	if candidateID == "" {
		candidateID = *applicationID
	}

	currentRecipe, currentProvenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	candidateRecipe, candidateProvenance, err := appidentity.LoadRecipeFromRegistry(*candidateRegistryPath, *candidateRecipeRoot, candidateID)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	return currentRecipe, currentProvenance, candidateRecipe, candidateProvenance, *mode, nil
}
