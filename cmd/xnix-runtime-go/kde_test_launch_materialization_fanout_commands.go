package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func runKDETestLaunchMaterializationFanOutPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-test-launch-materialization-fanout-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	stateRoot := flags.String("state-root", "", "existing controlled directory for test-only Runtime records")
	mode := flags.String("mode", "", "required fan-out mode; must be test-only")
	authorize := flags.String("authorize", "", "exact restricted test preparation authorization directive")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" || *applicationID == "" || *stateRoot == "" {
		return errors.New("kde-test-launch-materialization-fanout-preview requires --registry, --app, and --state-root")
	}
	if *mode != execution.RestrictedTestMode || *authorize != execution.RestrictedTestPreparationDirective {
		return errors.New("kde-test-launch-materialization-fanout-preview requires --mode test-only and --authorize authorize-restricted-test-preparation")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-test-launch-materialization-fanout-preview does not accept positional arguments")
	}
	recipeRecord, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDETestLaunchMaterializationFanOutPreview(recipeRecord, provenance, appidentity.KDERestrictedLaunchAuthorizationOptions{StateRoot: *stateRoot, Mode: *mode, Directive: *authorize})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
