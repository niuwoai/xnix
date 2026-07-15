package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/artifact"
)

func runArtifactStageRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("artifact-stage-record", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	manifestPath := flags.String("manifest", "", "artifact manifest JSON file")
	cacheRoot := flags.String("cache-root", "", "controlled artifact cache root")
	fixtureRoot := flags.String("fixture-root", "", "local fixture artifact source root")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *manifestPath == "" {
		return errors.New("artifact-stage-record requires --manifest")
	}
	if *cacheRoot == "" {
		return errors.New("artifact-stage-record requires --cache-root")
	}
	if *fixtureRoot == "" {
		return errors.New("artifact-stage-record requires --fixture-root")
	}
	data, err := os.ReadFile(*manifestPath)
	if err != nil {
		return err
	}
	manifest, err := artifact.ParseManifest(data)
	if err != nil {
		return err
	}
	receipt, err := artifact.StageFromFixture(artifact.StageRequest{
		Manifest:   manifest,
		CacheRoot:  *cacheRoot,
		FixtureDir: *fixtureRoot,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(receipt)
}
