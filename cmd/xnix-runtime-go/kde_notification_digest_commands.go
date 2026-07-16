package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type repeatedNotificationDigestEvents []appidentity.KDENotificationDigestEvent

func (events *repeatedNotificationDigestEvents) String() string {
	data, _ := json.Marshal([]appidentity.KDENotificationDigestEvent(*events))
	return string(data)
}

func (events *repeatedNotificationDigestEvents) Set(value string) error {
	parts := strings.Split(value, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return errors.New("notification digest event must use <group>:<event-id>")
	}
	*events = append(*events, appidentity.KDENotificationDigestEvent{Group: parts[0], EventID: parts[1]})
	return nil
}

func runKDENotificationDigestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, events, err := parseKDENotificationDigestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDENotificationDigestPreview(events)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseKDENotificationDigestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []appidentity.KDENotificationDigestEvent, error) {
	flags := flag.NewFlagSet("kde-notification-digest-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	var events repeatedNotificationDigestEvents
	flags.Var(&events, "event", "digest event as <group>:<event-id>; may be repeated")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, errors.New("kde-notification-digest-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, errors.New("kde-notification-digest-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, errors.New("kde-notification-digest-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, errors.New("kde-notification-digest-preview does not accept positional arguments")
	}
	if len(events) == 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, errors.New("kde-notification-digest-preview requires at least one --event")
	}
	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, append([]appidentity.KDENotificationDigestEvent(nil), events...), err
}
