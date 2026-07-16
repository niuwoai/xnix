package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEOfflineApplicationIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseKDEOfflineApplicationIdentityPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDEOfflineApplicationIdentityPreview(recipe, provenance)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseKDEOfflineApplicationIdentityPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet("kde-offline-application-identity-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	if *registryPath == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, errors.New("kde-offline-application-identity-preview requires --registry")
	}
	if *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, errors.New("kde-offline-application-identity-preview requires --app")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, errors.New("kde-offline-application-identity-preview does not accept positional arguments")
	}
	return appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
}
