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
