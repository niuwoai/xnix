package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/snapshot"
)

func runSnapshotRestoreCandidatesPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("snapshot-restore-candidates-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id")
	stateRoot := flags.String("state-root", "", "Runtime state root that holds the snapshot store")
	activeSession := flags.Bool("active-session", false, "treat the application as having an active session that blocks restore")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *applicationID == "" {
		return errors.New("snapshot-restore-candidates-preview requires --app")
	}
	if *stateRoot == "" {
		return errors.New("snapshot-restore-candidates-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("snapshot-restore-candidates-preview does not accept positional arguments")
	}

	store, err := snapshot.OpenReadOnly(*stateRoot)
	if err != nil {
		return err
	}
	manifests, err := store.List()
	if err != nil {
		return err
	}
	baseline, err := store.Baseline()
	if err != nil {
		return err
	}

	candidates := make([]appidentity.SnapshotCandidateInput, 0, len(manifests))
	for _, manifest := range manifests {
		candidates = append(candidates, appidentity.SnapshotCandidateInput{
			ID:          manifest.ID,
			Reason:      manifest.Reason,
			FileCount:   manifest.FileCount,
			ContentHash: manifest.ContentHash,
			Verified:    store.Verify(manifest.ID) == nil,
		})
	}

	preview, err := appidentity.NewSnapshotRestoreCandidatesPreview(appidentity.SnapshotRestoreCandidatesOptions{
		ApplicationID: *applicationID,
		Candidates:    candidates,
		BaselineID:    baseline.SnapshotID,
		ActiveSession: *activeSession,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
