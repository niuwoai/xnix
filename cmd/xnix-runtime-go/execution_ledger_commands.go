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
	"xnix.local/xnix/internal/runtime/portal"
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
	portalRequest := flags.String("portal-request", "", "Portal request handle token recorded in the same state root")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("execution-ledger-record requires --state-root")
	}
	if *applicationID == "" {
		return errors.New("execution-ledger-record requires --app")
	}

	portalReceipts, err := executionLedgerPortalReceipts(*stateRoot, csvValues(*portalRequest))
	if err != nil {
		return err
	}
	inputs, err := executionLedgerInputs(*applicationID, *trust, *environmentState, *snapshotBaseline, csvValues(*portalRequired), csvValues(*portalGranted), portalReceipts)
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

func runExecutionSessionRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("execution-session-record", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "Runtime state root used for the execution session record")
	requestID := flags.String("request-id", "", "execution transaction request id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("execution-session-record requires --state-root")
	}
	if *requestID == "" {
		return errors.New("execution-session-record requires --request-id")
	}
	if flags.NArg() != 0 {
		return errors.New("execution-session-record does not accept positional arguments")
	}

	ledger, err := execution.NewLedger(*stateRoot)
	if err != nil {
		return err
	}
	record, err := ledger.RecordSession(*requestID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func executionLedgerInputs(applicationID, trust, environmentState string, snapshotBaseline bool, portalRequired, portalGranted []string, portalReceipts []execution.PortalPermissionReceipt) (execution.Inputs, error) {
	trustState, err := executionLedgerTrust(trust)
	if err != nil {
		return execution.Inputs{}, err
	}
	record, err := executionLedgerEnvironment(applicationID, environmentState)
	if err != nil {
		return execution.Inputs{}, err
	}
	return execution.Inputs{
		Trust:                    trustState,
		Environment:              record,
		SnapshotBaselinePresent:  snapshotBaseline,
		PortalRequiredOps:        portalRequired,
		PortalGrantedOps:         portalGranted,
		PortalPermissionReceipts: portalReceipts,
	}, nil
}

func executionLedgerPortalReceipts(stateRoot string, handleTokens []string) ([]execution.PortalPermissionReceipt, error) {
	if len(handleTokens) == 0 {
		return nil, nil
	}
	ledger, err := portal.NewLedger(stateRoot)
	if err != nil {
		return nil, err
	}
	receipts := make([]execution.PortalPermissionReceipt, 0, len(handleTokens))
	for _, handleToken := range handleTokens {
		record, err := ledger.Inspect(handleToken)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, execution.PortalPermissionReceipt{
			HandleToken:           record.Request.HandleToken,
			Operation:             record.Request.Operation,
			RelativePath:          record.RelativePath,
			RequestState:          string(record.Request.State),
			PermissionState:       string(record.Request.PermissionState),
			PermissionGranted:     record.PermissionGranted,
			ExecutionApproved:     record.ExecutionApproved,
			RealPortalCallEnabled: record.RealPortalCallEnabled,
			StateRootPathExposed:  record.StateRootPathExposed,
			HostPermissionChanged: record.HostPermissionChanged,
		})
	}
	return receipts, nil
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
