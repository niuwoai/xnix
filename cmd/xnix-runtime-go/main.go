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
		return errors.New("usage: xnix-runtime-go desktop-identity-plan (--recipe PATH | --registry PATH --app ID)")
	}

	switch args[0] {
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runDesktopIdentityPlan(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("desktop-identity-plan", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return errors.New("desktop-identity-plan requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return errors.New("desktop-identity-plan requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return errors.New("desktop-identity-plan --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return errors.New("desktop-identity-plan does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
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
