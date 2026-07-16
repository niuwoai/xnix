package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/portal"
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

func runPortalRequestRecord(args []string, stdout io.Writer) error {
	options, err := parsePortalRequestRecordSource(args)
	if err != nil {
		return err
	}
	ledger, err := portal.NewLedger(options.stateRoot)
	if err != nil {
		return err
	}
	var record portal.Record
	switch options.action {
	case "create":
		record, err = ledger.Create(portal.RequestSpec{ApplicationID: options.applicationID, Operation: options.operation, Reason: options.reason})
	case "inspect":
		record, err = ledger.Inspect(options.handleToken)
	case "resolve":
		record, err = ledger.Resolve(options.handleToken, portal.Outcome(options.outcome))
	case "complete":
		record, err = ledger.Complete(options.handleToken)
	case "cancel":
		record, err = ledger.Cancel(options.handleToken)
	case "expire":
		record, err = ledger.Expire(options.handleToken)
	default:
		return errors.New("portal-request-record --action must be one of: create, inspect, resolve, complete, cancel, expire")
	}
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}

type portalRequestRecordOptions struct {
	stateRoot     string
	applicationID string
	operation     string
	action        string
	outcome       string
	handleToken   string
	reason        string
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

func parsePortalRequestRecordSource(args []string) (portalRequestRecordOptions, error) {
	flags := flag.NewFlagSet("portal-request-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	stateRoot := flags.String("state-root", "", "Runtime state root for Portal request records")
	applicationID := flags.String("app", "", "application id")
	operation := flags.String("operation", "", "desktop operation")
	action := flags.String("action", "", "action: create, inspect, resolve, complete, cancel, expire")
	outcome := flags.String("outcome", "", "Portal outcome for resolve: granted, denied, cancelled, failed, expired")
	handleToken := flags.String("handle-token", "", "recorded Portal request handle token")
	reason := flags.String("reason", "", "user-safe reason for creating a Portal request")
	if err := flags.Parse(args); err != nil {
		return portalRequestRecordOptions{}, err
	}
	if *stateRoot == "" {
		return portalRequestRecordOptions{}, errors.New("portal-request-record requires --state-root")
	}
	if *action == "" {
		return portalRequestRecordOptions{}, errors.New("portal-request-record requires --action")
	}
	if *action == "create" {
		if *applicationID == "" {
			return portalRequestRecordOptions{}, errors.New("portal-request-record create requires --app")
		}
		if *operation == "" {
			return portalRequestRecordOptions{}, errors.New("portal-request-record create requires --operation")
		}
	} else if *handleToken == "" {
		return portalRequestRecordOptions{}, errors.New("portal-request-record non-create actions require --handle-token")
	}
	if *action == "resolve" && *outcome == "" {
		return portalRequestRecordOptions{}, errors.New("portal-request-record resolve requires --outcome")
	}
	if flags.NArg() != 0 {
		return portalRequestRecordOptions{}, errors.New("portal-request-record does not accept positional arguments")
	}
	return portalRequestRecordOptions{
		stateRoot:     *stateRoot,
		applicationID: *applicationID,
		operation:     *operation,
		action:        *action,
		outcome:       *outcome,
		handleToken:   *handleToken,
		reason:        *reason,
	}, nil
}
