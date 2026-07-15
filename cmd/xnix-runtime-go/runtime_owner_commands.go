package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runRuntimeServiceBindingPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-service-binding-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime activation files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-service-binding-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeServiceBindingPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeLiveOwnerGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-live-owner-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime activation files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-live-owner-gate-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeLiveOwnerGatePreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeOwnerProcessPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-owner-process-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner process inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-owner-process-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeOwnerProcessPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeOwnerSmokePlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-owner-smoke-plan-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime activation files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-owner-smoke-plan-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeOwnerSmokePlanPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeMethodParityManifestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-method-parity-manifest-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime D-Bus contract and smoke files")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-method-parity-manifest-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeMethodParityManifestPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeOwnerReadinessPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-owner-readiness-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner readiness inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-owner-readiness-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeOwnerReadinessPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeOwnerRouteManifestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-owner-route-manifest-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner route inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-owner-route-manifest-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeOwnerRouteManifestPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeWriteGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-write-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime write gate inputs")
	methodName := flags.String("method", "", "reserved Runtime write method to evaluate")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *methodName == "" {
		return errors.New("runtime-write-gate-preview requires --method")
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-write-gate-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeWriteGatePreview(*root, *methodName)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func runRuntimeOwnerRecipeTrustPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("runtime-owner-recipe-trust-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime recipe registry")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("runtime-owner-recipe-trust-preview does not accept positional arguments")
	}

	preview, err := appidentity.NewRuntimeOwnerRecipeTrustPreview(*root)
	if err != nil {
		return err
	}

	return encodeIndentedJSON(stdout, preview)
}

func encodeIndentedJSON(stdout io.Writer, value any) error {
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}
