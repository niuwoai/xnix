package main

import (
	"encoding/json"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEIntegrationStatusPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "kde-integration-status-preview")
	}
	preview, err := appidentity.NewKDEIntegrationStatusPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEShellIntegrationPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "kde-shell-integration-preview")
	}
	preview, err := appidentity.NewKDEShellIntegrationPlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEApplicationSurfacePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("kde-application-surface-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEApplicationSurfacePlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
