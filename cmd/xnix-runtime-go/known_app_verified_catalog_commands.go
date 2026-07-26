package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownAppVerifiedCatalogPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogPreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	matrixEvidence := flags.String("matrix-evidence", "", "Go-owned known app matrix evidence JSON path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogPreviewRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalog(appidentity.KnownAppVerifiedCatalogRequest{
		MatrixEvidencePath: *matrixEvidence,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKnownAppVerifiedCatalogRunPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogRunPlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	verifiedCatalog := flags.String("verified-catalog", "", "Go-owned known app verified catalog JSON path")
	appID := flags.String("app", "", "known Windows app id to plan for q4 execution")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogRunPlanRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalogRunPlan(appidentity.KnownAppVerifiedCatalogRunPlanRequest{
		VerifiedCatalogPath: *verifiedCatalog,
		AppID:               *appID,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKnownAppVerifiedCatalogRunAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogRunAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	runPlan := flags.String("run-plan", "", "Go-owned verified catalog run plan JSON path")
	runReport := flags.String("known-winapp-run", "", "passed known Windows app run JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogRunAcceptanceRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalogRunAcceptance(appidentity.KnownAppVerifiedCatalogRunAcceptanceRequest{
		RunPlanPath:   *runPlan,
		RunReportPath: *runReport,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
