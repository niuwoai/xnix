package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func runCurrentMainlineOwnershipAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("current-mainline-ownership-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Xnix current mainline inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("current-mainline-ownership-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewCurrentMainlineOwnershipAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
