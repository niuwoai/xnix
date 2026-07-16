package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runStateRootQuotaRetentionPreview(args []string, stdout io.Writer) error {
	options, err := parseStateRootQuotaRetentionPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewStateRootQuotaRetentionPreview(options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseStateRootQuotaRetentionPreviewSource(args []string) (appidentity.StateRootQuotaRetentionOptions, error) {
	flags := flag.NewFlagSet("state-root-quota-retention-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id")
	stateRoot := flags.String("state-root", "", "Runtime state root to inspect in read-only dry-run mode")
	quotaBytes := flags.Int64("quota-bytes", 0, "optional quota threshold in bytes")
	activeSession := flags.Bool("active-session", false, "treat execution receipts as blocked by an active session")
	if err := flags.Parse(args); err != nil {
		return appidentity.StateRootQuotaRetentionOptions{}, err
	}
	if *applicationID == "" {
		return appidentity.StateRootQuotaRetentionOptions{}, errors.New("state-root-quota-retention-preview requires --app")
	}
	if *stateRoot == "" {
		return appidentity.StateRootQuotaRetentionOptions{}, errors.New("state-root-quota-retention-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return appidentity.StateRootQuotaRetentionOptions{}, errors.New("state-root-quota-retention-preview does not accept positional arguments")
	}
	return appidentity.StateRootQuotaRetentionOptions{
		ApplicationID: *applicationID,
		StateRoot:     *stateRoot,
		QuotaBytes:    *quotaBytes,
		ActiveSession: *activeSession,
	}, nil
}
