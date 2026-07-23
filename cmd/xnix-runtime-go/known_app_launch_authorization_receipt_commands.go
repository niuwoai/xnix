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
