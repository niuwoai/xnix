package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/diagnostics"
)

func runSupportBundleManifestPreview(args []string, stdout io.Writer) error {
	plan, history, options, err := parseSupportBundleManifestPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := plan.SupportBundleManifestPreview(history, options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseSupportBundleManifestPreviewSource(args []string) (appidentity.Plan, diagnostics.RunHistory, appidentity.SupportBundleManifestOptions, error) {
	flags := flag.NewFlagSet("support-bundle-manifest-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used for read-only diagnostic history")
	runtimeRoot := flags.String("runtime-root", ".", "project root used for Runtime version metadata")
	issue := flags.String("issue", "portal-approval-required", "compatibility issue to summarize")
	testType := flags.String("test-type", "smoke", "test type: preflight, smoke, or repair-readiness")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, errors.New("support-bundle-manifest-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, errors.New("support-bundle-manifest-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, errors.New("support-bundle-manifest-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, errors.New("support-bundle-manifest-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, err
	}
	history := diagnostics.RunHistory{
		SchemaVersion:               "xnix.runtime.diagnostic_run_history.v1",
		RecordType:                  "diagnostic-run-history",
		Source:                      "go-runtime-state-root-diagnostic-run-history",
		ApplicationID:               plan.ApplicationID,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		BackendStarted:              false,
		AIProviderCalled:            false,
		RealAIProviderEnabled:       false,
		AutoRepairAllowed:           false,
		RepairExecuted:              false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		FileContentsIncluded:        false,
		Summary:                     "Runtime has no persisted diagnostic runs under the configured state root.",
	}
	if *stateRoot != "" {
		store, err := diagnostics.OpenRunRecordStoreReadOnly(*stateRoot)
		if err != nil {
			return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, err
		}
		history, err = store.History(plan.ApplicationID)
		if err != nil {
			return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportBundleManifestOptions{}, err
		}
	}
	options := appidentity.SupportBundleManifestOptions{
		RuntimeRoot: *runtimeRoot,
		Issue:       *issue,
		TestType:    *testType,
	}
	return plan, history, options, nil
}
