package main

import (
	"errors"
	"flag"
	"io"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownAppRuntimeStatusLaunchOwnerTriggerPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppRuntimeStatusLaunchOwnerTriggerRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root containing the Runtime-status launch evidence handoff")
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-runtime-status-launch-owner-trigger-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("known-app-runtime-status-launch-owner-trigger-preview requires --state-root")
	}
	preview, err := appidentity.PreviewKnownAppRuntimeStatusLaunchOwnerTrigger(appidentity.KnownAppRuntimeStatusLaunchOwnerTriggerRequest{
		StateRoot:            *stateRoot,
		EvidenceID:           *evidenceID,
		EvidenceRelativePath: *evidenceRelativePath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
