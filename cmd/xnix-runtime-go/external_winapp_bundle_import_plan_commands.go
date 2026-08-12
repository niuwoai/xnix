package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/testversion"
)

func runExternalWinAppBundleImportPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppBundleImportPlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	bundleRoot := flags.String("bundle-root", "", "local portable Windows app directory to inspect")
	executableRelativePath := flags.String("executable-relative-path", "", "portable bundle relative path to the Windows .exe")
	appID := flags.String("app-id", "", "application id for the imported external Windows app bundle")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app bundle")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	version := flags.String("version", "", "optional plan version; defaults to the project version")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppBundleImportPlanRequestType)
	}
	planVersion := *version
	if planVersion == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		planVersion = current
	}
	plan, err := appidentity.PreviewExternalWinAppBundleImportPlan(appidentity.ExternalWinAppBundleImportPlanRequest{
		Version:                planVersion,
		BundleRoot:             *bundleRoot,
		ExecutableRelativePath: *executableRelativePath,
		AppID:                  *appID,
		DisplayName:            *displayName,
		AppVersion:             *appVersion,
	})
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}
