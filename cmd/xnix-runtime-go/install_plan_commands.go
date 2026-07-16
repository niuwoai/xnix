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

func runCompatibilityInstallPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, artifactReceiptPath, err := parseCompatibilityInstallPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	var receipt *artifact.StageReceipt
	if artifactReceiptPath != "" {
		loaded, err := loadArtifactStageReceipt(artifactReceiptPath)
		if err != nil {
			return err
		}
		receipt = &loaded
	}
	preview, err := plan.CompatibilityInstallPlanPreviewWithArtifactReceipt(mode, receipt)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseCompatibilityInstallPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, error) {
	flags := flag.NewFlagSet("compatibility-install-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "development", "install planning mode: production or development")
	artifactReceiptPath := flags.String("artifact-receipt", "", "artifact stage receipt JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("compatibility-install-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("compatibility-install-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("compatibility-install-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("compatibility-install-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, *artifactReceiptPath, err
}

func loadArtifactStageReceipt(path string) (artifact.StageReceipt, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return artifact.StageReceipt{}, err
	}
	var receipt artifact.StageReceipt
	if err := json.Unmarshal(data, &receipt); err != nil {
		return artifact.StageReceipt{}, err
	}
	return receipt, nil
}
