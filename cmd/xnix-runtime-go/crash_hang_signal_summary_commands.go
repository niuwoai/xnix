package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/diagnostics"
)

func runCrashHangSignalSummaryPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("crash-hang-signal-summary-preview", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root used for diagnostic run records")
	applicationID := flags.String("app", "", "optional application id filter")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("crash-hang-signal-summary-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("crash-hang-signal-summary-preview does not accept positional arguments")
	}

	store, err := diagnostics.OpenRunRecordStoreReadOnly(*stateRoot)
	if err != nil {
		return err
	}
	// Read the full history and let the read model apply the optional
	// application filter, so its mixed-application accounting can see records
	// that belong to other applications.
	history, malformed, err := store.LenientHistory("")
	if err != nil {
		return err
	}
	preview, err := appidentity.NewCrashHangSignalSummaryPreview(history, appidentity.CrashHangSignalSummaryOptions{
		ApplicationID:        *applicationID,
		MalformedHistory:     len(malformed) > 0,
		MalformedEvidenceIDs: malformed,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
