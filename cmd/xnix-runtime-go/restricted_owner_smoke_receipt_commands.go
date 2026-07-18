package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func runRestrictedOwnerSmokeReceiptRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-owner-smoke-receipt-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime activation and owner smoke inputs")
	stateRoot := flags.String("state-root", "", "controlled state root for the restricted owner smoke receipt")
	mode := flags.String("mode", "", "required receipt mode; must be restricted-smoke")
	authorize := flags.String("authorize", "", "exact restricted owner smoke authorization directive")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("restricted-owner-smoke-receipt-record requires --state-root")
	}
	if *mode != owner.RestrictedOwnerSmokeMode || *authorize != owner.RestrictedOwnerSmokeDirective {
		return errors.New("restricted-owner-smoke-receipt-record requires --mode restricted-smoke and --authorize authorize-restricted-owner-smoke")
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-owner-smoke-receipt-record does not accept positional arguments")
	}
	receipt, err := owner.RecordRestrictedOwnerSmokeReceipt(*root, *stateRoot, *mode, *authorize)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, receipt)
}

func runRestrictedOwnerSmokeReceiptFanOutPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-owner-smoke-receipt-fanout-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	stateRoot := flags.String("state-root", "", "controlled state root containing the restricted owner smoke receipt")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("restricted-owner-smoke-receipt-fanout-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-owner-smoke-receipt-fanout-preview does not accept positional arguments")
	}
	preview, err := owner.NewRestrictedOwnerSmokeReceiptFanOutPreview(*stateRoot)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runRestrictedOwnerSmokeReceiptLookupPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-owner-smoke-receipt-lookup-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner route inputs")
	receiptID := flags.String("receipt-id", owner.RestrictedOwnerSmokeOpaqueReceiptID, "opaque restricted owner smoke receipt id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-owner-smoke-receipt-lookup-preview does not accept positional arguments")
	}
	preview, err := owner.ResolveRestrictedOwnerSmokeReceipt(*root, *receiptID)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-owner-smoke-receipt-fanout-owner-route-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner route inputs")
	receiptID := flags.String("receipt-id", owner.RestrictedOwnerSmokeOpaqueReceiptID, "opaque restricted owner smoke receipt id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-owner-smoke-receipt-fanout-owner-route-preview does not accept positional arguments")
	}
	preview, err := owner.NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(*root, *receiptID)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-owner-smoke-receipt-fanout-owner-route-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing restricted owner smoke receipt fan-out and owner route inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-owner-smoke-receipt-fanout-owner-route-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewRestrictedOwnerSmokeReceiptFanOutOwnerRouteAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
