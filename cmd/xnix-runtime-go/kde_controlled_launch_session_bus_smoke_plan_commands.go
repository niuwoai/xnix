package main

import (
	"errors"
	"flag"
	"io"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKDEControlledLaunchSessionBusSmokePlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KDEControlledLaunchSessionBusSmokePlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root containing the Runtime-status launch evidence handoff")
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("kde-controlled-launch-session-bus-smoke-plan-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("kde-controlled-launch-session-bus-smoke-plan-preview requires --state-root")
	}
	preview, err := appidentity.PreviewKDEControlledLaunchSessionBusSmokePlan(appidentity.KDEControlledLaunchSessionBusSmokePlanRequest{
		StateRoot:            *stateRoot,
		EvidenceID:           *evidenceID,
		EvidenceRelativePath: *evidenceRelativePath,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
