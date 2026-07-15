package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runAIDiagnosticInputPreview(args []string, stdout io.Writer) error {
	plan, issue, testType, err := parseAIDiagnosticsPreviewSource("ai-diagnostic-input-preview", args)
	if err != nil {
		return err
	}
	preview, err := plan.AIDiagnosticInputPreview(issue, testType)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runAIDiagnosticRecommendationPreview(args []string, stdout io.Writer) error {
	plan, issue, testType, err := parseAIDiagnosticsPreviewSource("ai-diagnostic-recommendation-preview", args)
	if err != nil {
		return err
	}
	preview, err := plan.AIDiagnosticRecommendationPreview(issue, testType)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runAIRepairApprovalGatePreview(args []string, stdout io.Writer) error {
	plan, issue, testType, err := parseAIDiagnosticsPreviewSource("ai-repair-approval-gate-preview", args)
	if err != nil {
		return err
	}
	preview, err := plan.AIRepairApprovalGatePreview(issue, testType)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseAIDiagnosticsPreviewSource(command string, args []string) (appidentity.Plan, string, string, error) {
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	issue := flags.String("issue", "engine-binding-pending", "compatibility issue to diagnose")
	testType := flags.String("test-type", "preflight", "test type: preflight, smoke, or repair-readiness")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Plan{}, "", "", errors.New(command + " requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Plan{}, "", "", errors.New(command + " requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Plan{}, "", "", errors.New(command + " --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, "", "", errors.New(command + " does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Plan{}, "", "", err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, "", "", err
	}
	return plan, *issue, *testType, nil
}
