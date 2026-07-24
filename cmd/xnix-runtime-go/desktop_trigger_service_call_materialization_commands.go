package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/owner"
)

func runDesktopTriggerServiceCallMaterializationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(owner.DesktopTriggerServiceCallMaterializationRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root containing Runtime-status launch evidence")
	desktopEntryFile := flags.String("desktop-entry-file", "", "KDE controlled-launch desktop entry metadata file")
	evidenceID := flags.String("evidence-id", "", "opaque Runtime-status launch evidence handoff id")
	evidenceRelativePath := flags.String("evidence-relative-path", "", "relative Runtime-status launch evidence handoff path")
	expectedEvidenceSHA256 := flags.String("expected-evidence-sha256", "", "expected Runtime-status launch evidence digest")
	fullCheckpointPromoted := flags.Bool("full-checkpoint-promoted", false, "treat the formal full checkpoint as promoted")
	humanAuthorizedSmoke := flags.Bool("human-authorized-smoke", false, "materialize evidence-only owner service arguments for an explicitly human-authorized smoke candidate without claiming checkpoint promotion")
	routeID := flags.String("route-id", "", "desktop route id")
	methodID := flags.String("method-id", "", "Runtime owner service method id")
	actionID := flags.String("action-id", "", "desktop action id")
	callerRole := flags.String("caller-role", "", "desktop caller role")
	requestFreshness := flags.String("request-freshness", "", "desktop envelope freshness state")
	replayMarker := flags.String("replay-marker", "", "desktop envelope replay marker state")
	kdeStateRoot := flags.String("kde-state-root", "", "forbidden KDE-supplied state root")
	kdeCacheRoot := flags.String("kde-cache-root", "", "forbidden KDE-supplied cache root")
	kdeLauncherPath := flags.String("kde-launcher-path", "", "forbidden KDE-supplied launcher path")
	kdeTimeoutValue := flags.String("kde-timeout", "", "forbidden KDE-supplied timeout value")
	kdeRawExecutablePath := flags.String("kde-raw-executable-path", "", "forbidden KDE-supplied raw executable path")
	kdeBackendCommand := flags.String("kde-backend-command", "", "forbidden KDE-supplied backend command")
	kdeReceiptID := flags.String("kde-receipt-id", "", "forbidden KDE-supplied receipt id")
	kdeSessionID := flags.String("kde-session-id", "", "forbidden KDE-supplied session id")
	kdeDispatchID := flags.String("kde-dispatch-id", "", "forbidden KDE-supplied dispatch id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("desktop-trigger-service-call-materialization-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("desktop-trigger-service-call-materialization-preview requires --state-root")
	}
	if strings.TrimSpace(*desktopEntryFile) == "" {
		return errors.New("desktop-trigger-service-call-materialization-preview requires --desktop-entry-file")
	}
	contents, err := os.ReadFile(*desktopEntryFile)
	if err != nil {
		return errors.New("desktop-trigger-service-call-materialization-preview could not read desktop entry metadata")
	}
	preview, err := owner.PreviewDesktopTriggerServiceCallMaterialization(owner.DesktopTriggerServiceCallMaterializationRequest{
		StateRoot:              *stateRoot,
		DesktopEntryContent:    string(contents),
		EvidenceID:             *evidenceID,
		EvidenceRelativePath:   *evidenceRelativePath,
		ExpectedEvidenceSHA256: *expectedEvidenceSHA256,
		FullCheckpointPromoted: *fullCheckpointPromoted,
		HumanAuthorizedSmoke:   *humanAuthorizedSmoke,
		RouteID:                *routeID,
		MethodID:               *methodID,
		ActionID:               *actionID,
		CallerRole:             *callerRole,
		RequestFreshness:       *requestFreshness,
		ReplayMarker:           *replayMarker,
		KDEStateRoot:           *kdeStateRoot,
		KDECacheRoot:           *kdeCacheRoot,
		KDELauncherPath:        *kdeLauncherPath,
		KDETimeoutValue:        *kdeTimeoutValue,
		KDERawExecutablePath:   *kdeRawExecutablePath,
		KDEBackendCommand:      *kdeBackendCommand,
		KDEReceiptID:           *kdeReceiptID,
		KDESessionID:           *kdeSessionID,
		KDEDispatchID:          *kdeDispatchID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
