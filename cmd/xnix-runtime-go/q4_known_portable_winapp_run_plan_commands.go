package main

import (
	"errors"
	"flag"
	"io"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func runQ4KnownPortableWinAppRunPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.Q4KnownPortableWinAppRunPlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	appID := flags.String("app", "org.xnix.external.notepadplusplus", "known portable app id")
	remoteHost := flags.String("remote", "root@q4", "q4 SSH target")
	remoteMaterialsRoot := flags.String("remote-materials-root", "/home/xnix-run-materials", "scoped q4 materials root")
	remoteSourceRoot := flags.String("remote-source-root", "/home/xnix-build/xnix-q4-known-portable-winapp-run", "scoped q4 source root")
	remoteBuildRoot := flags.String("remote-build-root", "/home/xnix-build-cache", "scoped q4 build/cache root")
	output := flags.String("output", "", "optional delegated JSON output path")
	markdownOutput := flags.String("markdown-output", "", "optional delegated Markdown output path")
	execute := flags.Bool("execute", false, "include the delegated execute flag in the operator command")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("q4-known-portable-winapp-run-plan-preview does not accept positional arguments")
	}

	preview, err := appidentity.PreviewQ4KnownPortableWinAppRunPlan(appidentity.Q4KnownPortableWinAppRunPlanRequest{
		Version:             readRuntimeGoProjectVersion(),
		AppID:               *appID,
		RemoteHost:          *remoteHost,
		RemoteMaterialsRoot: *remoteMaterialsRoot,
		RemoteSourceRoot:    *remoteSourceRoot,
		RemoteBuildRoot:     *remoteBuildRoot,
		Output:              *output,
		MarkdownOutput:      *markdownOutput,
		Execute:             *execute,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
