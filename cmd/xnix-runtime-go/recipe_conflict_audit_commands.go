package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runRecipeConflictAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("recipe-conflict-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" {
		return errors.New("recipe-conflict-audit-preview requires --registry")
	}
	if flags.NArg() != 0 {
		return errors.New("recipe-conflict-audit-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRecipeConflictAuditPreview(appidentity.RecipeConflictAuditOptions{
		RegistryPath: *registryPath,
		RecipeRoot:   *recipeRoot,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
