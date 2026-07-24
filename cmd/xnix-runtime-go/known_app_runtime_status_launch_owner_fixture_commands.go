package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runKnownAppRuntimeStatusLaunchOwnerFixtureRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppRuntimeStatusLaunchOwnerFixtureRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "Runtime owner state root where launch handoff state is prepared")
	cacheRoot := flags.String("cache-root", "", "managed known Windows application cache root")
	guestBoundary := flags.String("guest-boundary", "", "managed known Windows application guest boundary")
	status := flags.String("status", "", "delegated smoke status to project into Runtime-status launch evidence")
	guiSmokeEvidenceFile := flags.String("gui-smoke-evidence-file", "", "projected GUI smoke evidence file to consume before preparing Runtime-status launch handoff")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-runtime-status-launch-owner-fixture-record does not accept positional arguments")
	}
	if *stateRoot == "" {
		return errors.New("known-app-runtime-status-launch-owner-fixture-record requires --state-root")
	}
	record, err := appidentity.RecordKnownAppRuntimeStatusLaunchOwnerFixture(appidentity.KnownAppRuntimeStatusLaunchOwnerFixtureRequest{
		AppID:                *appID,
		StateRoot:            *stateRoot,
		CacheRoot:            *cacheRoot,
		GuestBoundary:        *guestBoundary,
		Status:               *status,
		GUISmokeEvidencePath: *guiSmokeEvidenceFile,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}
