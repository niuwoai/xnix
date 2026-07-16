package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func runKDERestrictedProductSmokeCheckpointRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-restricted-product-smoke-checkpoint-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	stateRoot := flags.String("state-root", "", "existing controlled directory for test-only Runtime records")
	repositoryRoot := flags.String("repo-root", ".", "repository root containing fixed product-image evidence")
	manifestSource := flags.String("manifest", "image/kinoite/manifest.json", "product-image manifest relative to the repository root")
	mode := flags.String("mode", "", "required checkpoint mode; must be test-only")
	authorize := flags.String("authorize", "", "exact restricted test preparation authorization directive")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" || *applicationID == "" || *stateRoot == "" || *repositoryRoot == "" {
		return errors.New("kde-restricted-product-smoke-checkpoint-record requires --registry, --app, --state-root, and --repo-root")
	}
	if *mode != execution.RestrictedTestMode || *authorize != execution.RestrictedTestPreparationDirective {
		return errors.New("kde-restricted-product-smoke-checkpoint-record requires --mode test-only and --authorize authorize-restricted-test-preparation")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-restricted-product-smoke-checkpoint-record does not accept positional arguments")
	}
	recipeRecord, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	record, err := appidentity.NewKDERestrictedProductSmokeCheckpointRecord(recipeRecord, provenance, appidentity.KDERestrictedProductSmokeCheckpointOptions{
		KDERestrictedLaunchAuthorizationOptions: appidentity.KDERestrictedLaunchAuthorizationOptions{StateRoot: *stateRoot, Mode: *mode, Directive: *authorize},
		RepositoryRoot:                          *repositoryRoot,
		ManifestSource:                          *manifestSource,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}
