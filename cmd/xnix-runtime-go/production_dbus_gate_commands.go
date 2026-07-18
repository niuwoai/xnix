package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func runProductionDBusGateReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-gate-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner gate review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-gate-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusGateReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDBusHumanAuthorizationPreflightPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-human-authorization-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner human authorization preflight inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-human-authorization-preflight-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusHumanAuthorizationPreflightPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDBusMethodReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-method-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner method review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-method-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusMethodReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionRollbackDiagnosticsReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-rollback-diagnostics-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner rollback diagnostics review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-rollback-diagnostics-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionRollbackDiagnosticsReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDesktopSideEffectReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-desktop-side-effect-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner desktop side-effect review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-desktop-side-effect-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDesktopSideEffectReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
