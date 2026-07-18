package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/execution"
)

func runKDETestLaunchMaterializationFanOutPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-test-launch-materialization-fanout-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	stateRoot := flags.String("state-root", "", "existing controlled directory for test-only Runtime records")
	mode := flags.String("mode", "", "required fan-out mode; must be test-only")
	authorize := flags.String("authorize", "", "exact restricted test preparation authorization directive")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" || *applicationID == "" || *stateRoot == "" {
		return errors.New("kde-test-launch-materialization-fanout-preview requires --registry, --app, and --state-root")
	}
	if *mode != execution.RestrictedTestMode || *authorize != execution.RestrictedTestPreparationDirective {
		return errors.New("kde-test-launch-materialization-fanout-preview requires --mode test-only and --authorize authorize-restricted-test-preparation")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-test-launch-materialization-fanout-preview does not accept positional arguments")
	}
	recipeRecord, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDETestLaunchMaterializationFanOutPreview(recipeRecord, provenance, appidentity.KDERestrictedLaunchAuthorizationOptions{StateRoot: *stateRoot, Mode: *mode, Directive: *authorize})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKDETestLaunchMaterializationFanOutConsumePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-test-launch-materialization-fanout-consume-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	stateRoot := flags.String("state-root", "", "existing controlled directory containing test-only Runtime records")
	materializationPlanID := flags.String("materialization-plan-id", "", "existing restricted materialization plan id to consume")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *registryPath == "" || *applicationID == "" || *stateRoot == "" || *materializationPlanID == "" {
		return errors.New("kde-test-launch-materialization-fanout-consume-preview requires --registry, --app, --state-root, and --materialization-plan-id")
	}
	if flags.NArg() != 0 {
		return errors.New("kde-test-launch-materialization-fanout-consume-preview does not accept positional arguments")
	}
	recipeRecord, provenance, err := appidentity.LoadRecipeFromRegistry(*registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDETestLaunchMaterializationFanOutPreviewFromReceipt(recipeRecord, provenance, appidentity.KDETestLaunchMaterializationFanOutOptions{StateRoot: *stateRoot, MaterializationPlanID: *materializationPlanID})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKDETestLaunchMaterializationReceiptLookupPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-test-launch-materialization-receipt-lookup-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner route inputs")
	receiptID := flags.String("receipt-id", appidentity.KDETestLaunchMaterializationOpaqueReceiptID, "opaque KDE test launch materialization receipt id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("kde-test-launch-materialization-receipt-lookup-preview does not accept positional arguments")
	}
	preview, err := appidentity.ResolveKDETestLaunchMaterializationReceipt(*root, *receiptID)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKDETestLaunchMaterializationOwnerRouteAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("kde-test-launch-materialization-owner-route-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing materialization fan-out and owner route inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("kde-test-launch-materialization-owner-route-audit-preview does not accept positional arguments")
	}
	preview, err := appidentity.NewKDETestLaunchMaterializationOwnerRouteAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
