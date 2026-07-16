package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type repeatedAppIDs []string

func (ids *repeatedAppIDs) String() string {
	data, _ := json.Marshal([]string(*ids))
	return string(data)
}

func (ids *repeatedAppIDs) Set(value string) error {
	if value == "" {
		return errors.New("application id must not be empty")
	}
	*ids = append(*ids, value)
	return nil
}

func runMultiApplicationInstallQueuePreview(args []string, stdout io.Writer) error {
	plans, mode, err := parseMultiApplicationInstallQueueSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewMultiApplicationInstallQueuePreview(plans, mode)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseMultiApplicationInstallQueueSource(args []string) ([]appidentity.Plan, string, error) {
	flags := flag.NewFlagSet("multi-application-install-queue-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	mode := flags.String("mode", "development", "install planning mode: production or development")
	var appIDs repeatedAppIDs
	flags.Var(&appIDs, "app", "application id to include; may be repeated")
	if err := flags.Parse(args); err != nil {
		return nil, "", err
	}
	if *registryPath == "" {
		return nil, "", errors.New("multi-application-install-queue-preview requires --registry")
	}
	if flags.NArg() != 0 {
		return nil, "", errors.New("multi-application-install-queue-preview does not accept positional arguments")
	}

	ids, err := multiApplicationInstallQueueIDs(*registryPath, appIDs)
	if err != nil {
		return nil, "", err
	}
	plans := make([]appidentity.Plan, 0, len(ids))
	for _, id := range ids {
		recipe, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, id)
		if err != nil {
			return nil, "", err
		}
		plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
		if err != nil {
			return nil, "", err
		}
		plans = append(plans, plan)
	}
	return plans, *mode, nil
}

func multiApplicationInstallQueueIDs(registryPath string, selected repeatedAppIDs) ([]string, error) {
	if len(selected) > 0 {
		return append([]string(nil), selected...), nil
	}
	data, err := os.ReadFile(registryPath)
	if err != nil {
		return nil, err
	}
	registry, err := appidentity.ParseRegistry(data)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(registry.Recipes))
	for _, entry := range registry.Recipes {
		ids = append(ids, entry.ID)
	}
	return ids, nil
}
