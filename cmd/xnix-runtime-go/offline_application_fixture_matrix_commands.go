package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

type repeatedFixtureShapeIDs []string

func (ids *repeatedFixtureShapeIDs) String() string {
	data, _ := json.Marshal([]string(*ids))
	return string(data)
}

func (ids *repeatedFixtureShapeIDs) Set(value string) error {
	if value == "" {
		return errors.New("fixture shape id must not be empty")
	}
	*ids = append(*ids, value)
	return nil
}

func runOfflineApplicationFixtureMatrixPreview(args []string, stdout io.Writer) error {
	options, err := parseOfflineApplicationFixtureMatrixOptions(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewOfflineApplicationFixtureMatrixPreview(options)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func parseOfflineApplicationFixtureMatrixOptions(args []string) (appidentity.OfflineApplicationFixtureMatrixOptions, error) {
	flags := flag.NewFlagSet("offline-application-fixture-matrix-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	runtimeRoot := flags.String("runtime-root", ".", "project root used only for read-only Runtime evidence checks")
	artifactReceiptRoot := flags.String("artifact-receipt-root", "", "optional controlled root for read-only artifact stage receipt evidence")
	snapshotStateRoot := flags.String("snapshot-state-root", "", "optional controlled state root for read-only snapshot baseline evidence")
	var shapeIDs repeatedFixtureShapeIDs
	flags.Var(&shapeIDs, "shape", "fixture shape id to include; may be repeated")
	if err := flags.Parse(args); err != nil {
		return appidentity.OfflineApplicationFixtureMatrixOptions{}, err
	}
	if flags.NArg() != 0 {
		return appidentity.OfflineApplicationFixtureMatrixOptions{}, errors.New("offline-application-fixture-matrix-preview does not accept positional arguments")
	}
	return appidentity.OfflineApplicationFixtureMatrixOptions{
		ShapeIDs:            append([]string(nil), shapeIDs...),
		RuntimeRoot:         *runtimeRoot,
		ArtifactReceiptRoot: *artifactReceiptRoot,
		SnapshotStateRoot:   *snapshotStateRoot,
	}, nil
}
