package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
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
