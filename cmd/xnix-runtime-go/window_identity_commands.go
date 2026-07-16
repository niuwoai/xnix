package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runTaskManagerIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseTaskManagerIdentityPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TaskManagerIdentityPlanPreviewWithOptions(options)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseTaskManagerIdentityPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.TaskManagerIdentityOptions, error) {
	flags := flag.NewFlagSet("task-manager-identity-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TaskManagerIdentityOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TaskManagerIdentityOptions{}, errors.New("task-manager-identity-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TaskManagerIdentityOptions{}, errors.New("task-manager-identity-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TaskManagerIdentityOptions{}, errors.New("task-manager-identity-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.TaskManagerIdentityOptions{}, errors.New("task-manager-identity-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, appidentity.TaskManagerIdentityOptions{ActivationRoot: *activationRoot}, err
}

func runKWinWindowRulePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("kwin-window-rule-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KWinWindowRulePlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
