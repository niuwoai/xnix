package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: xnix-runtime-go {backend-selection-preview|compatibility-center-preview|desktop-entry-preview|desktop-identity-plan|desktop-resource-bridge-preview|file-open-preview|krunner-query-preview|mimeapps-preview|mode-switch-preview|notification-preview|permission-review-preview|portal-request-preview|review-flow-preview|settings-change-preview|settings-preview|tray-status-preview|window-identity-preview}")
	}

	switch args[0] {
	case "backend-selection-preview":
		return runBackendSelectionPreview(args[1:], stdout)
	case "compatibility-center-preview":
		return runCompatibilityCenterPreview(args[1:], stdout)
	case "desktop-entry-preview":
		return runDesktopEntryPreview(args[1:], stdout)
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	case "desktop-resource-bridge-preview":
		return runDesktopResourceBridgePreview(args[1:], stdout)
	case "file-open-preview":
		return runFileOpenPreview(args[1:], stdout)
	case "krunner-query-preview":
		return runKRunnerQueryPreview(args[1:], stdout)
	case "mimeapps-preview":
		return runMIMEAppsPreview(args[1:], stdout)
	case "mode-switch-preview":
		return runModeSwitchPreview(args[1:], stdout)
	case "notification-preview":
		return runNotificationPreview(args[1:], stdout)
	case "permission-review-preview":
		return runPermissionReviewPreview(args[1:], stdout)
	case "portal-request-preview":
		return runPortalRequestPreview(args[1:], stdout)
	case "review-flow-preview":
		return runReviewFlowPreview(args[1:], stdout)
	case "settings-change-preview":
		return runSettingsChangePreview(args[1:], stdout)
	case "settings-preview":
		return runSettingsPreview(args[1:], stdout)
	case "tray-status-preview":
		return runTrayStatusPreview(args[1:], stdout)
	case "window-identity-preview":
		return runWindowIdentityPreview(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runCompatibilityCenterPreview(args []string, stdout io.Writer) error {
	recipes, provenance, err := parseRegistryPreviewSource("compatibility-center-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewCompatibilityCenterPreview(recipes, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendSelectionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-selection-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendSelectionPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runFileOpenPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, err := parseFileOpenPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewFileOpenPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKRunnerQueryPreview(args []string, stdout io.Writer) error {
	recipes, provenance, query, err := parseKRunnerQueryPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKRunnerQueryPreview(recipes, provenance, query)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopIdentityPlan(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-identity-plan", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}

func runDesktopEntryPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-entry-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	entry, err := plan.RenderDesktopEntry()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, entry)
	return err
}

func runMIMEAppsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("mimeapps-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	mimeapps, err := plan.RenderMIMEApps()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, mimeapps)
	return err
}

func runDesktopResourceBridgePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-resource-bridge-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runNotificationPreview(args []string, stdout io.Writer) error {
	recipe, provenance, eventType, err := parseNotificationPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.NotificationPreview(eventType)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("settings-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runModeSwitchPreview(args []string, stdout io.Writer) error {
	recipe, provenance, requestedMode, err := parseModeSwitchPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ModeSwitchPreview(requestedMode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsChangePreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, err := parseSettingsChangePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsChangePreview(sectionID, fieldID, value)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runReviewFlowPreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, operation, err := parseReviewFlowPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ReviewFlowPreview(sectionID, fieldID, value, operation)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPermissionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("permission-review-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PermissionReviewPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPortalRequestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, operation, reason, err := parsePortalRequestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PortalRequestPreview(operation, reason)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runTrayStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("tray-status-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TrayStatusPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runWindowIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("window-identity-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.WindowIdentityPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseNotificationPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("notification-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	eventType := flags.String("event", "", "notification event type")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *eventType == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires --event")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *eventType, err
}

func parsePortalRequestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, error) {
	flags := flag.NewFlagSet("portal-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	operation := flags.String("operation", "", "sensitive desktop operation")
	reason := flags.String("reason", "", "user-facing Portal request reason")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
	}
	if *operation == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --operation")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *operation, *reason, err
}

func parseModeSwitchPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("mode-switch-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	requestedMode := flags.String("mode", "", "requested compatibility mode")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *requestedMode == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --mode")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *requestedMode, err
}

func parseReviewFlowPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, string, error) {
	flags := flag.NewFlagSet("review-flow-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "resource-access", "settings section identifier")
	fieldID := flags.String("field", "documents", "settings field identifier")
	value := flags.String("value", "ask", "requested settings value")
	operation := flags.String("operation", "file-open", "sensitive desktop operation")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, *operation, err
}

func parseSettingsChangePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, error) {
	flags := flag.NewFlagSet("settings-change-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "", "settings section identifier")
	fieldID := flags.String("field", "", "settings field identifier")
	value := flags.String("value", "", "requested settings value")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", err
	}
	if *sectionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --section")
	}
	if *fieldID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --field")
	}
	if *value == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --value")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, err
}

func parseKRunnerQueryPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("krunner-query-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	query := flags.String("query", "", "KRunner query text")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", errors.New("krunner-query-preview requires --registry")
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, "", errors.New("krunner-query-preview does not accept positional arguments")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *query, err
}

func parseRegistryPreviewSource(commandName string, args []string) ([]appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s requires --registry", commandName)
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
}

func parseFileOpenPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet("file-open-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the file-open preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, errors.New("file-open-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, errors.New("file-open-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, err
}

func parseRecipeSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
}

func loadRecipe(recipePath string, registryPath string, recipeRoot string, applicationID string) (appidentity.Recipe, appidentity.Provenance, error) {
	if registryPath != "" {
		return appidentity.LoadRecipeFromRegistry(registryPath, recipeRoot, applicationID)
	}

	data, err := os.ReadFile(recipePath)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("read recipe: %w", err)
	}
	recipe, err := appidentity.ParseRecipe(data)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	return recipe, appidentity.Provenance{Source: "direct-file"}, nil
}
