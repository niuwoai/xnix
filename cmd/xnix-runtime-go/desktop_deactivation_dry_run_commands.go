package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runDesktopDeactivationDryRunPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("desktop-deactivation-dry-run-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	activationRoot := flags.String("activation-root", "", "activation receipt root used to verify the recorded activation receipt")
	activeSession := flags.Bool("active-session", false, "treat the application as having an active session that blocks removal")
	sharedMIME := flags.String("shared-mime", "", "comma-separated MIME types also claimed by other applications")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return errors.New("desktop-deactivation-dry-run-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return errors.New("desktop-deactivation-dry-run-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return errors.New("desktop-deactivation-dry-run-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return errors.New("desktop-deactivation-dry-run-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopDeactivationDryRunPreview(appidentity.DesktopDeactivationDryRunOptions{
		ActivationRoot:         *activationRoot,
		ActiveSession:          *activeSession,
		SharedMIMEAssociations: splitSharedMIME(*sharedMIME),
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func splitSharedMIME(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	var out []string
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}
