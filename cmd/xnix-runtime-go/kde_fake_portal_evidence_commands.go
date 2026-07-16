package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEFakePortalEvidenceRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-fake-portal-evidence-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	stateRoot := flags.String("state-root", "", "existing controlled directory for test-only Runtime records")
	mode := flags.String("mode", "", "required Portal evidence mode; must be test-only")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" || *applicationID == "" || *stateRoot == "" {
		return errors.New("kde-fake-portal-evidence-record requires --registry, --app, and --state-root")
	}
	if *mode != "test-only" {
		return errors.New("kde-fake-portal-evidence-record requires --mode test-only")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-fake-portal-evidence-record does not accept positional arguments")
	}
	recipeRecord, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	record, err := appidentity.NewKDEFakePortalEvidenceRecord(recipeRecord, provenance, appidentity.KDEFakeExecutionEvidenceOptions{StateRoot: *stateRoot, Mode: *mode})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}
