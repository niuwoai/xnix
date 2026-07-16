package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDESearchVisibilityPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-search-visibility-plan-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "optional activation receipt root used to gate receipt-backed visibility")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return errors.New("kde-search-visibility-plan-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return errors.New("kde-search-visibility-plan-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return errors.New("kde-search-visibility-plan-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-search-visibility-plan-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDESearchVisibilityPlanPreview(appidentity.KDESearchVisibilityOptions{
		ActivationRoot:      *activationRoot,
		SupportedExtensions: recipe.SupportedExtensions,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
