package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: xnix-runtime-go {desktop-entry-preview|desktop-identity-plan|mimeapps-preview|window-identity-preview} (--recipe PATH | --registry PATH --app ID)")
	}

	switch args[0] {
	case "desktop-entry-preview":
		return runDesktopEntryPreview(args[1:], stdout)
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	case "mimeapps-preview":
		return runMIMEAppsPreview(args[1:], stdout)
	case "window-identity-preview":
		return runWindowIdentityPreview(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runDesktopIdentityPlan(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-identity-plan", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}

func runDesktopEntryPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-entry-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	entry, err := plan.RenderDesktopEntry()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, entry)
	return err
}

func runMIMEAppsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("mimeapps-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	mimeapps, err := plan.RenderMIMEApps()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, mimeapps)
	return err
}

func runWindowIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("window-identity-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.WindowIdentityPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseRecipeSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
}

func loadRecipe(recipePath string, registryPath string, recipeRoot string, applicationID string) (appidentity.Recipe, appidentity.Provenance, error) {
	if registryPath != "" {
		return appidentity.LoadRecipeFromRegistry(registryPath, recipeRoot, applicationID)
	}

	data, err := os.ReadFile(recipePath)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("read recipe: %w", err)
	}
	recipe, err := appidentity.ParseRecipe(data)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	return recipe, appidentity.Provenance{Source: "direct-file"}, nil
}
