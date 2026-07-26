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

func runKnownAppVerifiedCatalogLaunchHandoffRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogLaunchHandoffRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root where the verified catalog launch handoff is persisted")
	runAcceptance := flags.String("known-app-verified-catalog-run-acceptance", "", "Go-owned verified catalog app run acceptance JSON path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogLaunchHandoffRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogLaunchHandoff(appidentity.KnownAppVerifiedCatalogLaunchHandoffRequest{
		StateRoot:      *stateRoot,
		AcceptancePath: *runAcceptance,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runKnownAppVerifiedCatalogLaunchMaterializationRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root containing the verified catalog launch handoff")
	handoffRelativePath := flags.String("handoff-relative-path", "", "relative verified catalog launch handoff path forwarded by the desktop")
	cacheRoot := flags.String("cache-root", "", "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogLaunchMaterialization(appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequest{
		StateRoot:           *stateRoot,
		HandoffRelativePath: *handoffRelativePath,
		CacheRoot:           *cacheRoot,
		GuestBoundary:       *guestBoundary,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runKnownAppVerifiedCatalogDispatchRequestRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogDispatchRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root containing the verified catalog launch handoff")
	handoffRelativePath := flags.String("handoff-relative-path", "", "relative verified catalog launch handoff path forwarded by the desktop")
	cacheRoot := flags.String("cache-root", "", "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	launcherName := flags.String("launcher", "", "Runtime-owner managed launcher name or path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogDispatchRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogDispatchRequest(appidentity.KnownAppVerifiedCatalogDispatchRequestRecordRequest{
		StateRoot:           *stateRoot,
		HandoffRelativePath: *handoffRelativePath,
		CacheRoot:           *cacheRoot,
		GuestBoundary:       *guestBoundary,
		LauncherName:        *launcherName,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}
