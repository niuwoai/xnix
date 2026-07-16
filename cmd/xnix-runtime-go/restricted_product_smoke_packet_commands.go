package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/image"
)

func runRestrictedProductSmokePacketPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("restricted-product-smoke-packet-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	repoRoot := flags.String("repo-root", ".", "project root used for read-only evidence checks")
	manifest := flags.String("manifest", "image/kinoite/manifest.json", "manifest path relative to the project root")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("restricted-product-smoke-packet-preview does not accept positional arguments")
	}
	packet, err := image.PrepareRestrictedProductSmokePacket(*repoRoot, *manifest)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, packet)
}
