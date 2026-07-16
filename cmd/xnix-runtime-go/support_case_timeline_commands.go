package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/diagnostics"
)

func runSupportCaseTimelinePreview(args []string, stdout io.Writer) error {
	plan, history, options, err := parseSupportCaseTimelinePreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := plan.SupportCaseTimelinePreview(history, options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseSupportCaseTimelinePreviewSource(args []string) (appidentity.Plan, diagnostics.RunHistory, appidentity.SupportCaseTimelineOptions, error) {
	flags := flag.NewFlagSet("support-case-timeline-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used for read-only diagnostic history")
	runtimeRoot := flags.String("runtime-root", ".", "project root used for Runtime version metadata")
	issue := flags.String("issue", "portal-approval-required", "compatibility issue to summarize")
	testType := flags.String("test-type", "smoke", "test type: preflight, smoke, or repair-readiness")
	decision := flags.String("decision", "deferred", "KDE review decision metadata")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, errors.New("support-case-timeline-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, errors.New("support-case-timeline-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, errors.New("support-case-timeline-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, errors.New("support-case-timeline-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, err
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
	malformed := []string{}
	if *stateRoot != "" {
		store, err := diagnostics.OpenRunRecordStoreReadOnly(*stateRoot)
		if err != nil {
			return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, err
		}
		// Read all records and let the Runtime read model account for
		// mixed-application evidence while rendering only the target app.
		history, malformed, err = store.LenientHistory("")
		if err != nil {
			return appidentity.Plan{}, diagnostics.RunHistory{}, appidentity.SupportCaseTimelineOptions{}, err
		}
	}
	options := appidentity.SupportCaseTimelineOptions{
		RuntimeRoot:          *runtimeRoot,
		Issue:                *issue,
		TestType:             *testType,
		Decision:             *decision,
		MalformedHistory:     len(malformed) > 0,
		MalformedEvidenceIDs: malformed,
	}
	return plan, history, options, nil
}
