package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runSettingsProfileMigrationPreview(args []string, stdout io.Writer) error {
	plan, options, err := parseSettingsProfileMigrationPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsProfileMigrationPreview(options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseSettingsProfileMigrationPreviewSource(args []string) (appidentity.Plan, appidentity.SettingsProfileMigrationOptions, error) {
	flags := flag.NewFlagSet("settings-profile-migration-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	fromSchema := flags.String("from-schema", "xnix.runtime.settings.v0", "source settings schema version")
	toSchema := flags.String("to-schema", "xnix.runtime.settings.v1", "target settings schema version")
	blocked := flags.String("blocked-setting", "", "comma-separated settings ids blocked by migration policy")
	activationRoot := flags.String("activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	if err := flags.Parse(args); err != nil {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, errors.New("settings-profile-migration-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, errors.New("settings-profile-migration-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, errors.New("settings-profile-migration-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, errors.New("settings-profile-migration-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	if err != nil {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return appidentity.Plan{}, appidentity.SettingsProfileMigrationOptions{}, err
	}
	options := appidentity.SettingsProfileMigrationOptions{
		FromSchemaVersion: *fromSchema,
		ToSchemaVersion:   *toSchema,
		BlockedSettingIDs: settingsProfileMigrationCSVValues(*blocked),
		ActivationRoot:    *activationRoot,
	}
	return plan, options, nil
}

func settingsProfileMigrationCSVValues(value string) []string {
	out := []string{}
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		out = append(out, part)
	}
	return out
}
