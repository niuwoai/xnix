package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

func runDiagnosticRunRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("diagnostic-run-record", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root used for diagnostic run records")
	applicationID := flags.String("app", "", "application id")
	runID := flags.String("run-id", "", "diagnostic run id")
	fixturePath := flags.String("fixture", "", "fixture JSON file containing diagnostic signals")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("diagnostic-run-record requires --state-root")
	}
	if *applicationID == "" {
		return errors.New("diagnostic-run-record requires --app")
	}
	if *runID == "" {
		return errors.New("diagnostic-run-record requires --run-id")
	}
	if *fixturePath == "" {
		return errors.New("diagnostic-run-record requires --fixture")
	}
	if flags.NArg() != 0 {
		return errors.New("diagnostic-run-record does not accept positional arguments")
	}

	fixture, err := readDiagnosticFixture(*fixturePath)
	if err != nil {
		return err
	}
	store, err := diagnostics.NewRunRecordStore(*stateRoot)
	if err != nil {
		return err
	}
	record, err := store.Record(diagnostics.RunRecordRequest{
		ApplicationID: *applicationID,
		RunID:         *runID,
		Fixture:       fixture,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runDiagnosticRunHistory(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("diagnostic-run-history", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root used for diagnostic run records")
	applicationID := flags.String("app", "", "optional application id filter")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("diagnostic-run-history requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("diagnostic-run-history does not accept positional arguments")
	}

	store, err := diagnostics.NewRunRecordStore(*stateRoot)
	if err != nil {
		return err
	}
	history, err := store.History(*applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(history)
}

func readDiagnosticFixture(path string) (diagnostics.Fixture, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return diagnostics.Fixture{}, fmt.Errorf("read diagnostic fixture: %w", err)
	}
	var fixture diagnostics.Fixture
	if err := json.Unmarshal(data, &fixture); err != nil {
		return diagnostics.Fixture{}, fmt.Errorf("parse diagnostic fixture: %w", err)
	}
	return fixture, nil
}
