package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runBackendAdapterContractPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("backend-adapter-contract-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "repository root used to read the Xnix version")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("backend-adapter-contract-preview does not accept positional arguments")
	}
	preview, err := appidentity.NewBackendAdapterContractPreview(*root)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendAdapterContractOwnerRouteAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("backend-adapter-contract-owner-route-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing adapter contract and owner route inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("backend-adapter-contract-owner-route-audit-preview does not accept positional arguments")
	}
	preview, err := appidentity.NewBackendAdapterContractOwnerRouteAuditPreview(*root)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendAdapterRedactedProfileAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("backend-adapter-redacted-profile-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing adapter contract and owner route inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("backend-adapter-redacted-profile-audit-preview does not accept positional arguments")
	}
	preview, err := appidentity.NewBackendAdapterRedactedProfileAuditPreview(*root)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}
