package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/activation"
	"xnix.local/xnix/internal/runtime/appidentity"
)

func runDesktopActivationStage(args []string, stdout io.Writer) error {
	recipe, provenance, mode, stagingRoot, managedLauncherBinary, err := parseDesktopActivationStageSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	result, err := activation.Stage(activation.StageRequest{
		Root:                  stagingRoot,
		Mode:                  mode,
		Plan:                  plan,
		ManagedLauncherBinary: managedLauncherBinary,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func parseDesktopActivationStageSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, error) {
	flags := flag.NewFlagSet("desktop-activation-stage", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	externalAppImportRecord := flags.String("external-app-import-record", "", "Runtime external Windows app import record")
	mode := flags.String("mode", "development", "activation staging mode: production or development")
	stagingRoot := flags.String("staging-root", "", "test root where desktop activation files may be staged")
	managedLauncherBinary := flags.String("managed-launcher-bin", "", "optional path to a prebuilt xnix-compat-launch binary to copy into the staging root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", err
	}
	sourceCount := 0
	for _, source := range []string{*recipePath, *registryPath, *externalAppImportRecord} {
		if source != "" {
			sourceCount++
		}
	}
	if sourceCount != 1 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage requires exactly one source: --recipe, --registry, or --external-app-import-record")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage --recipe cannot be combined with --app or --recipe-root")
	}
	if *externalAppImportRecord != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage --external-app-import-record cannot be combined with --app or --recipe-root")
	}
	if *stagingRoot == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage requires --staging-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("desktop-activation-stage does not accept positional arguments")
	}

	if *externalAppImportRecord != "" {
		record, err := appidentity.LoadExternalWinAppImportRecord(*externalAppImportRecord)
		if err != nil {
			return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", err
		}
		recipe, provenance, err := appidentity.ExternalAppRecipeFromImportRecord(record)
		return recipe, provenance, *mode, *stagingRoot, *managedLauncherBinary, err
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, *stagingRoot, *managedLauncherBinary, err
}
