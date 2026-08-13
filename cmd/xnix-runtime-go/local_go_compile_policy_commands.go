package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func runLocalGoCompilePolicyPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("local-go-compile-policy-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Xnix q4 remote build policy inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("local-go-compile-policy-preview does not accept positional arguments")
	}
	preview, err := owner.NewLocalGoCompilePolicyPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
