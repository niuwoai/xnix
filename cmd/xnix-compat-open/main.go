package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const commandName = "xnix-compat-open"
const packagedRecipeRegistryDir = "/usr/share/xnix/compatibility/recipes"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", commandName, err)
		os.Exit(64)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var applicationID string
	var registryPath string
	var recipeDir string
	var recipeRoot string
	var activationRoot string
	flags.StringVar(&applicationID, "app", "", "application id to use for the file-open preview")
	flags.StringVar(&registryPath, "registry", "", "path to an Xnix recipe registry JSON file")
	flags.StringVar(&recipeDir, "recipe-dir", packagedRecipeRegistryDir, "directory containing the default Xnix recipe registry")
	flags.StringVar(&recipeRoot, "recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	flags.StringVar(&activationRoot, "activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return err
	}

	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return fmt.Errorf("%s requires at least one file URI", commandName)
	}
	registry := strings.TrimSpace(registryPath)
	if registry == "" {
		registry = filepath.Join(strings.TrimSpace(recipeDir), "registry.json")
	}
	if strings.TrimSpace(registry) == "" {
		return fmt.Errorf("%s requires --registry or --recipe-dir", commandName)
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(registry, strings.TrimSpace(recipeRoot))
	if err != nil {
		return err
	}
	preview, err := appidentity.NewFileOpenPreviewWithOptions(
		recipes,
		provenance,
		fileURIs,
		strings.TrimSpace(applicationID),
		appidentity.FileOpenOptions{ActivationRoot: strings.TrimSpace(activationRoot)},
	)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
