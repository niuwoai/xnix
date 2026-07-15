package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runStateRootPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("state-root-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ApplicationStateRootPreview()
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runSnapshotPlanPreview(args []string, stdout io.Writer) error {
	applicationID, reason, err := parseSnapshotPlanPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewSnapshotPlanPreview(applicationID, reason)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runPortalAccessPolicyPreview(args []string, stdout io.Writer) error {
	applicationID, operation, err := parsePortalAccessPolicyPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewPortalAccessPolicyPreview(applicationID, operation)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseSnapshotPlanPreviewSource(args []string) (string, string, error) {
	flags := flag.NewFlagSet("snapshot-plan-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id")
	reason := flags.String("reason", "manual", "snapshot reason")
	if err := flags.Parse(args); err != nil {
		return "", "", err
	}
	if *applicationID == "" {
		return "", "", errors.New("snapshot-plan-preview requires --app")
	}
	if flags.NArg() != 0 {
		return "", "", errors.New("snapshot-plan-preview does not accept positional arguments")
	}
	return *applicationID, *reason, nil
}

func parsePortalAccessPolicyPreviewSource(args []string) (string, string, error) {
	flags := flag.NewFlagSet("portal-access-policy-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id")
	operation := flags.String("operation", "", "desktop operation")
	if err := flags.Parse(args); err != nil {
		return "", "", err
	}
	if *applicationID == "" {
		return "", "", errors.New("portal-access-policy-preview requires --app")
	}
	if *operation == "" {
		return "", "", errors.New("portal-access-policy-preview requires --operation")
	}
	if flags.NArg() != 0 {
		return "", "", errors.New("portal-access-policy-preview does not accept positional arguments")
	}
	return *applicationID, *operation, nil
}
