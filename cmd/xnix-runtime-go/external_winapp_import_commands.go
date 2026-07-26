package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/testversion"
)

func runExternalWinAppImportRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.ExternalWinAppImportRecordRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for imported external Windows apps")
	executablePath := flags.String("executable", "", "local Windows .exe file to import")
	appID := flags.String("app-id", "", "application id for the imported external Windows app")
	displayName := flags.String("display-name", "", "display name for the imported external Windows app")
	appVersion := flags.String("app-version", "", "optional imported app version; defaults to the project version")
	version := flags.String("version", "", "optional record version; defaults to the project version")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.ExternalWinAppImportRecordRequestType)
	}
	recordVersion := *version
	if recordVersion == "" {
		current, err := testversion.Current()
		if err != nil {
			return err
		}
		recordVersion = current
	}
	record, err := appidentity.RecordExternalWinAppImport(appidentity.ExternalWinAppImportRequest{
		Version:        recordVersion,
		StateRoot:      *stateRoot,
		ExecutablePath: *executablePath,
		AppID:          *appID,
		DisplayName:    *displayName,
		AppVersion:     *appVersion,
	})
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}
