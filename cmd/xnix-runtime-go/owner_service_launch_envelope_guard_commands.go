package main

import (
	"errors"
	"flag"
	"io"
	"strings"

	"xnix.local/xnix/internal/runtime/owner"
)

func runOwnerServiceLaunchEnvelopeGuardPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(owner.LaunchEnvelopeGuardRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime owner state root used internally for evidence binding")
	routeID := flags.String("route-id", "", "desktop route id")
	methodID := flags.String("method-id", "", "Runtime owner service method id")
	evidenceHandleKind := flags.String("evidence-handle-kind", "", "evidence handle kind: evidence-id or evidence-relative-path")
	evidenceHandle := flags.String("evidence-handle", "", "opaque or relative Runtime evidence handle")
	expectedEvidenceSHA256 := flags.String("expected-evidence-sha256", "", "expected Runtime evidence digest")
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
		return errors.New("owner-service-launch-envelope-guard-preview does not accept positional arguments")
	}
	if strings.TrimSpace(*stateRoot) == "" {
		return errors.New("owner-service-launch-envelope-guard-preview requires --state-root")
	}
	preview, err := owner.PreviewLaunchEnvelopeGuard(owner.LaunchEnvelopeGuardRequest{
		StateRoot:              *stateRoot,
		RouteID:                *routeID,
		MethodID:               *methodID,
		EvidenceHandleKind:     *evidenceHandleKind,
		EvidenceHandle:         *evidenceHandle,
		ExpectedEvidenceSHA256: *expectedEvidenceSHA256,
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
