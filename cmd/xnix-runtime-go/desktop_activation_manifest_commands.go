package main

import (
	"encoding/json"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runDesktopActivationManifestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-activation-manifest-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationManifestPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
