package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runRuntimePolicyExplanationCardsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, options, err := parseRuntimePolicyExplanationCardsSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.RuntimePolicyExplanationCardsPreview(options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseRuntimePolicyExplanationCardsSource(args []string) (appidentity.Recipe, appidentity.Provenance, appidentity.RuntimePolicyExplanationCardsOptions, error) {
	flags := flag.NewFlagSet("runtime-policy-explanation-cards-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "development", "policy explanation mode: production or development")
	runtimeRoot := flags.String("runtime-root", ".", "Runtime root used for read-only version evidence")
	stateRoot := flags.String("state-root", "", "optional Runtime state root used only for read-only lifecycle evidence")
	portalOperation := flags.String("portal-operation", "file-open", "Portal operation to explain")
	snapshotReason := flags.String("snapshot-reason", "before-repair", "snapshot reason to explain")
	issue := flags.String("issue", "portal-approval-required", "repair issue to explain")
	testType := flags.String("test-type", "smoke", "diagnostic test type to explain")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, err
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, errors.New("runtime-policy-explanation-cards-preview does not accept positional arguments")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, errors.New("runtime-policy-explanation-cards-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, errors.New("runtime-policy-explanation-cards-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", "runtime-policy-explanation-cards-preview")
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, appidentity.RuntimePolicyExplanationCardsOptions{}, err
	}
	return recipe, provenance, appidentity.RuntimePolicyExplanationCardsOptions{
		Environment:     *mode,
		RuntimeRoot:     *runtimeRoot,
		StateRoot:       *stateRoot,
		PortalOperation: *portalOperation,
		SnapshotReason:  *snapshotReason,
		Issue:           *issue,
		TestType:        *testType,
	}, nil
}
