package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func runWindowsCompatibilityWorkstreamsPreview(args []string, stdout io.Writer) error {
	if len(args) != 0 {
		return fmt.Errorf("%s does not accept arguments", "windows-compatibility-workstreams-preview")
	}
	preview, err := appidentity.NewWindowsCompatibilityWorkstreamsPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runWindowsAppRunSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-run-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var stateRoot string
	var runnerPath string
	var timeoutText string
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&runnerPath, "runner", "", "explicit compatibility runner path")
	flags.StringVar(&timeoutText, "timeout", "20s", "execution timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-run-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}

	result, err := winapp.RunSmoke(context.Background(), winapp.Request{
		ExecutablePath: exePath,
		Arguments:      []string(appArgs),
		StateRoot:      stateRoot,
		RunnerPath:     runnerPath,
		Timeout:        timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppContainerRunSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-container-run-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var stateRoot string
	var image string
	var platform string
	var dockerPath string
	var timeoutText string
	var bootstrapTimeoutText string
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&stateRoot, "state-root", "", "isolated Runtime state root")
	flags.StringVar(&image, "image", winapp.DefaultContainerImage, "local Wine container image")
	flags.StringVar(&platform, "platform", winapp.DefaultWinePlatform, "container platform")
	flags.StringVar(&dockerPath, "docker", "", "explicit docker runner path")
	flags.StringVar(&timeoutText, "timeout", "45s", "execution timeout")
	flags.StringVar(&bootstrapTimeoutText, "bootstrap-timeout", winapp.DefaultWineBootstrapTimeout.String(), "Wine prefix bootstrap timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-container-run-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	bootstrapTimeout, err := time.ParseDuration(bootstrapTimeoutText)
	if err != nil {
		return fmt.Errorf("parse bootstrap timeout: %w", err)
	}

	result, err := winapp.RunContainerSmoke(context.Background(), winapp.ContainerRequest{
		ExecutablePath:   exePath,
		Arguments:        []string(appArgs),
		StateRoot:        stateRoot,
		Image:            image,
		Platform:         platform,
		DockerPath:       dockerPath,
		Timeout:          timeout,
		BootstrapTimeout: bootstrapTimeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func runWindowsAppGuestWineSmoke(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("windows-app-guest-wine-smoke", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var exePath string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	flags.StringVar(&exePath, "exe", "", "Windows executable path")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", winapp.DefaultRemoteDir, "guest remote smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&timeoutText, "timeout", "60s", "guest execution timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the Windows executable")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", "windows-app-guest-wine-smoke")
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunGuestSmoke(context.Background(), winapp.GuestRequest{
		ExecutablePath: exePath,
		Arguments:      []string(appArgs),
		Host:           host,
		Port:           parsedPort,
		User:           user,
		KeyPath:        keyPath,
		RemoteDir:      remoteDir,
		SSHPath:        sshPath,
		SCPPath:        scpPath,
		Timeout:        timeout,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

type repeatedStringFlag []string

func (flag *repeatedStringFlag) String() string {
	return fmt.Sprint([]string(*flag))
}

func (flag *repeatedStringFlag) Set(value string) error {
	*flag = append(*flag, value)
	return nil
}
