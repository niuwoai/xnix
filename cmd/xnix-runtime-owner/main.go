package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("xnix-runtime-owner", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner inputs")
	mode := flags.String("mode", string(owner.ModePreview), "owner mode: preview or smoke-owner")
	writeMethod := flags.String("deny-write", "", "render a deterministic disabled write-method response")
	readMethod := flags.String("dispatch-read", "", "render a read-only Runtime owner method dispatch response")
	lifecycleLog := flags.Bool("lifecycle-log", false, "render smoke-owner lifecycle events as JSON Lines")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *readMethod == "" && flags.NArg() != 0 {
		return errors.New("xnix-runtime-owner does not accept positional arguments")
	}

	var payload any
	selectedOperations := 0
	for _, selected := range []bool{*writeMethod != "", *readMethod != "", *lifecycleLog} {
		if selected {
			selectedOperations++
		}
	}
	if selectedOperations > 1 {
		return errors.New("xnix-runtime-owner accepts only one of --deny-write, --dispatch-read, or --lifecycle-log")
	}
	if *writeMethod != "" {
		response, err := owner.DisabledWriteResponse(*writeMethod)
		if err != nil {
			return err
		}
		payload = response
	} else if *readMethod != "" {
		dispatch, err := owner.DispatchRead(*root, *readMethod, flags.Args())
		if err != nil {
			return err
		}
		payload = dispatch
	} else if *lifecycleLog {
		events, err := owner.NewLifecycleEvents(*root, owner.CandidateMode(*mode))
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(stdout)
		for _, event := range events {
			if err := encoder.Encode(event); err != nil {
				return err
			}
		}
		return nil
	} else {
		candidate, err := owner.NewCandidate(*root, owner.CandidateMode(*mode))
		if err != nil {
			return err
		}
		payload = candidate
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(payload)
}
