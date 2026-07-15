package main

import (
	"encoding/json"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runTaskManagerIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("task-manager-identity-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TaskManagerIdentityPlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKWinWindowRulePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("kwin-window-rule-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KWinWindowRulePlanPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
