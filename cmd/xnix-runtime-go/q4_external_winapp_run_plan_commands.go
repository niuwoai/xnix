package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4ExternalWinAppRunPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4ExternalWinAppRunPlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	executable := flags.String("executable", "", "local Windows .exe under this checkout or /tmp/xnix-*")
	appID := flags.String("app-id", "org.xnix.external.uploaded", "application id for the uploaded executable")
	displayName := flags.String("display-name", "Uploaded Windows App", "display name for the uploaded executable")
	windowMatch := flags.String("window-match", "", "observed-window text proving file-open behavior")
	remoteHost := flags.String("remote", "root@q4", "q4 SSH target")
	remoteMaterialsRoot := flags.String("remote-materials-root", "/home/xnix-run-materials", "scoped q4 materials root")
	projectRoot := flags.String("project-root", ".", "checkout root used for scoped local executable validation")
	output := flags.String("output", "", "optional delegated JSON output path")
	markdownOutput := flags.String("markdown-output", "", "optional delegated Markdown output path")
	execute := flags.Bool("execute", false, "include the delegated execute flag in the operator command")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-external-winapp-run-plan-preview does not accept positional arguments")
	}

	preview, err := appidentity.PreviewQ4ExternalWinAppRunPlan(appidentity.Q4ExternalWinAppRunPlanRequest{
		Version:             readRuntimeGoProjectVersion(),
		ProjectRoot:         *projectRoot,
		LocalExecutable:     *executable,
		AppID:               *appID,
		DisplayName:         *displayName,
		WindowMatch:         *windowMatch,
		RemoteHost:          *remoteHost,
		RemoteMaterialsRoot: *remoteMaterialsRoot,
		Output:              *output,
		MarkdownOutput:      *markdownOutput,
		Execute:             *execute,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
