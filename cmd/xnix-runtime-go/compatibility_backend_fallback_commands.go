package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/diagnostics"
)

func runCompatibilityBackendFallbackPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("compatibility-backend-fallback-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used for read-only diagnostics risk evidence")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return errors.New("compatibility-backend-fallback-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return errors.New("compatibility-backend-fallback-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return errors.New("compatibility-backend-fallback-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return errors.New("compatibility-backend-fallback-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}

	options := appidentity.CompatibilityBackendFallbackOptions{}
	if *stateRoot != "" {
		store, err := diagnostics.OpenRunRecordStoreReadOnly(*stateRoot)
		if err != nil {
			return err
		}
		history, err := store.History(plan.ApplicationID)
		if err != nil {
			return err
		}
		options.DiagnosticsFailingRuns = history.Counts.Failed
		options.DiagnosticsBlockedRuns = history.Counts.Blocked
	}

	preview, err := plan.CompatibilityBackendFallbackPreview(options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
