package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout *os.File) error {
	if len(args) == 0 {
		return errors.New("usage: xnix-runtime-go desktop-identity-plan --recipe PATH")
	}

	switch args[0] {
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runDesktopIdentityPlan(args []string, stdout *os.File) error {
	flags := flag.NewFlagSet("desktop-identity-plan", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *recipePath == "" {
		return errors.New("desktop-identity-plan requires --recipe")
	}
	if flags.NArg() != 0 {
		return errors.New("desktop-identity-plan does not accept positional arguments")
	}

	data, err := os.ReadFile(*recipePath)
	if err != nil {
		return fmt.Errorf("read recipe: %w", err)
	}
	recipe, err := appidentity.ParseRecipe(data)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlan(recipe)
	if err != nil {
		return err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}
