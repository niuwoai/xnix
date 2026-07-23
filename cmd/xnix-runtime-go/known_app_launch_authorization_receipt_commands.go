package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func runKnownAppLaunchAuthorizationReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-launch-authorization-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for the opaque launch authorization receipt")
	authorize := flags.String("authorize", "", "exact launch authorization action id")
	evidenceSource := flags.String("evidence-source", "staged-launcher-dispatch-smoke", "redacted known-app smoke evidence source")
	centerCardState := flags.String("center-card-state", "validated-launch-authorization-required", "Compatibility Center card state that requested authorization")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-launch-authorization-receipt-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-launch-authorization-receipt-preview does not accept positional arguments")
	}
	preview, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:           *appID,
		StateRoot:       *stateRoot,
		Authorize:       *authorize,
		EvidenceSource:  *evidenceSource,
		CenterCardState: *centerCardState,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppLaunchGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-launch-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-launch-gate-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-launch-gate-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppLaunchGate(appidentity.KnownAppLaunchGateRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledDispatchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-dispatch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-dispatch-request-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-dispatch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledDispatchRequest(appidentity.KnownAppControlledDispatchRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledExecutionSessionPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-execution-session-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledExecutionSession(appidentity.KnownAppControlledExecutionSessionRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledExecutionSessionRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-execution-session-record requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-record does not accept positional arguments")
	}
	record, err := appidentity.RecordKnownAppControlledExecutionSession(appidentity.KnownAppControlledExecutionSessionRecordRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}

func runKnownAppControlledExecutionSessionConsumePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-consume-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-controlled-execution-session-consume-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-consume-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledExecutionSessionConsumption(appidentity.KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedLaunchReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	actionID := flags.String("action", appidentity.KnownAppSessionGatedLaunchReviewAction, "exact session-gated launch review action id")
	decision := flags.String("decision", "reviewed", "review decision: reviewed, approved, deferred, or rejected")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedLaunchReview(appidentity.KnownAppSessionGatedLaunchReviewRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ActionID:  *actionID,
		Decision:  *decision,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedLaunchReviewReceiptRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-receipt-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	actionID := flags.String("action", appidentity.KnownAppSessionGatedLaunchReviewAction, "exact session-gated launch review action id")
	decision := flags.String("decision", "reviewed", "review decision: reviewed, approved, deferred, or rejected")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-receipt-record requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-receipt-record does not accept positional arguments")
	}
	record, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ActionID:  *actionID,
		Decision:  *decision,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}

func runKnownAppSessionGatedLaunchReviewGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted session-gated launch review receipt")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	receiptID := flags.String("receipt-id", "", "optional opaque session-gated launch review receipt id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-gate-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-gate-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedLaunchReviewGate(appidentity.KnownAppSessionGatedLaunchReviewGateRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ReceiptID: *receiptID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedControlledDispatchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-controlled-dispatch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing launch and review receipts")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	reviewReceiptID := flags.String("review-receipt-id", "", "opaque session-gated launch review receipt id")
	launchReceiptID := flags.String("launch-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *reviewReceiptID == "" || *launchReceiptID == "" {
		return errors.New("known-app-session-gated-controlled-dispatch-request-preview requires --state-root, --review-receipt-id, and --launch-receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-controlled-dispatch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedControlledDispatchRequest(appidentity.KnownAppSessionGatedControlledDispatchRequest{
		AppID:           *appID,
		StateRoot:       *stateRoot,
		SessionID:       *sessionID,
		ReviewReceiptID: *reviewReceiptID,
		LaunchReceiptID: *launchReceiptID,
		CacheRoot:       *cacheRoot,
		GuestBoundary:   *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppKDERuntimeStatusLaunchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-kde-runtime-status-launch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	launchReceiptID := flags.String("launch-authorization-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	reviewReceiptID := flags.String("session-gated-review-receipt-id", "", "opaque known Windows app session-gated review receipt id")
	sessionID := flags.String("session-id", "", "opaque known Windows app controlled execution session id")
	centerCardState := flags.String("center-card-state", "", "Compatibility Center card state from the KDE read model")
	primaryActionID := flags.String("primary-action-id", "", "KDE primary action id from the Runtime-owned card")
	postReviewDispatchState := flags.String("post-review-dispatch-state", "", "post-review dispatch state from the Runtime-owned card")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-kde-runtime-status-launch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppKDERuntimeStatusLaunchRequest(appidentity.KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        *appID,
		LaunchAuthorizationReceiptID: *launchReceiptID,
		SessionGatedReviewReceiptID:  *reviewReceiptID,
		ControlledExecutionSessionID: *sessionID,
		CenterCardState:              *centerCardState,
		PrimaryActionID:              *primaryActionID,
		PostReviewDispatchState:      *postReviewDispatchState,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
