package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"xnix.local/xnix/internal/runtime/environment"
	"xnix.local/xnix/internal/runtime/execution"
	"xnix.local/xnix/internal/runtime/recipe"
)

func runExecutionLedgerRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("execution-ledger-record", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root used for the execution ledger")
	applicationID := flags.String("app", "", "application id")
	decision := flags.String("decision", "approved", "review decision: approved or rejected")
	trust := flags.String("recipe-trust", "production", "recipe trust: production, development, or failed")
	environmentState := flags.String("environment", "ready", "environment state: ready, planned, or repair-required")
	snapshotBaseline := flags.Bool("snapshot-baseline", true, "whether a snapshot baseline is present")
	portalRequired := flags.String("portal-required", "file-open", "comma-separated required Portal operations")
	portalGranted := flags.String("portal-granted", "file-open", "comma-separated granted Portal operations")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("execution-ledger-record requires --state-root")
	}
	if *applicationID == "" {
		return errors.New("execution-ledger-record requires --app")
	}

	inputs, err := executionLedgerInputs(*applicationID, *trust, *environmentState, *snapshotBaseline, csvValues(*portalRequired), csvValues(*portalGranted))
	if err != nil {
		return err
	}
	pipeline := execution.NewPipeline()
	tx, err := pipeline.Create(*applicationID, inputs)
	if err != nil {
		return err
	}
	reviewed, err := tx.Review(execution.Decision(*decision))
	if err != nil {
		return err
	}
	tx = reviewed.Preflight(inputs)

	ledger, err := execution.NewLedger(*stateRoot)
	if err != nil {
		return err
	}
	record, err := ledger.Record(tx)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func executionLedgerInputs(applicationID, trust, environmentState string, snapshotBaseline bool, portalRequired, portalGranted []string) (execution.Inputs, error) {
	trustState, err := executionLedgerTrust(trust)
	if err != nil {
		return execution.Inputs{}, err
	}
	record, err := executionLedgerEnvironment(applicationID, environmentState)
	if err != nil {
		return execution.Inputs{}, err
	}
	return execution.Inputs{
		Trust:                   trustState,
		Environment:             record,
		SnapshotBaselinePresent: snapshotBaseline,
		PortalRequiredOps:       portalRequired,
		PortalGrantedOps:        portalGranted,
	}, nil
}

func executionLedgerTrust(value string) (recipe.TrustState, error) {
	switch value {
	case "production":
		return recipe.TrustState{DigestVerified: true, ProductionTrusted: true}, nil
	case "development":
		return recipe.TrustState{DigestVerified: true, DevelopmentOnly: true}, nil
	case "failed":
		return recipe.TrustState{FailedClosed: true, Reason: "ledger fixture trust failure"}, nil
	default:
		return recipe.TrustState{}, fmt.Errorf("unsupported recipe trust %q", value)
	}
}

func executionLedgerEnvironment(applicationID, state string) (environment.Record, error) {
	record := environment.Record{
		ApplicationID: applicationID,
		Profile:       environment.ProfileLocal,
		State:         environment.State(state),
	}
	switch environment.State(state) {
	case environment.StateReady:
		record.SatisfiedGates = environment.RequiredGates()
	case environment.StatePlanned:
		record.SatisfiedGates = []string{}
	case environment.StateRepairRequired:
		record.RepairHints = []string{"review compatibility environment diagnostics"}
	default:
		return environment.Record{}, fmt.Errorf("unsupported environment state %q", state)
	}
	return record, nil
}

func csvValues(value string) []string {
	var values []string
	for _, part := range strings.Split(value, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			values = append(values, part)
		}
	}
	return values
}
