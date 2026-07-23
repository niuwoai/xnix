package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

const launcherName = "xnix-compat-launch"

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", launcherName, err)
		os.Exit(64)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(launcherName, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var appID string
	var cacheRoot string
	var guestBoundary string
	var stateRoot string
	var receiptID string
	var host string
	var port string
	var user string
	var keyPath string
	var remoteDir string
	var sshPath string
	var scpPath string
	var timeoutText string
	flags.StringVar(&appID, "app", "", "known Windows app id")
	flags.StringVar(&cacheRoot, "cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	flags.StringVar(&guestBoundary, "guest-boundary", "", "controlled managed guest boundary supplied by the Runtime owner or smoke harness")
	flags.StringVar(&stateRoot, "state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	flags.StringVar(&receiptID, "receipt-id", "", "opaque known Windows app launch authorization receipt id")
	flags.StringVar(&host, "host", winapp.DefaultGuestHost, "guest SSH host")
	flags.StringVar(&port, "port", winapp.DefaultGuestPort, "guest SSH port")
	flags.StringVar(&user, "user", winapp.DefaultGuestUser, "guest SSH user")
	flags.StringVar(&keyPath, "key", "", "guest SSH private key path")
	flags.StringVar(&remoteDir, "remote-dir", "/tmp/xnix-known-winapp-smoke", "guest remote smoke directory")
	flags.StringVar(&sshPath, "ssh", "", "explicit ssh client path")
	flags.StringVar(&scpPath, "scp", "", "explicit scp client path")
	flags.StringVar(&timeoutText, "timeout", winapp.DefaultKnownAppGuestTimeout.String(), "guest execution timeout")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app")

	if err := flags.Parse(args); err != nil {
		return err
	}
	if appID == "" {
		return fmt.Errorf("--app is required")
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", launcherName)
	}

	launcherArgv := []string{launcherName, "--app", appID}
	bridge, err := winapp.PreviewKnownPortableLaunchBridge(winapp.KnownLaunchBridgeRequest{
		AppID:               appID,
		CacheRoot:           cacheRoot,
		ManagedLauncherArgv: launcherArgv,
	})
	if err != nil {
		return err
	}
	if guestBoundary == "" {
		return encode(stdout, bridge)
	}
	if stateRoot == "" || receiptID == "" {
		return fmt.Errorf("--state-root and --receipt-id are required when --guest-boundary requests dispatch")
	}
	controlledDispatch, err := appidentity.PreviewKnownAppControlledDispatchRequest(appidentity.KnownAppControlledDispatchRequest{
		AppID:         appID,
		StateRoot:     stateRoot,
		ReceiptID:     receiptID,
		CacheRoot:     cacheRoot,
		GuestBoundary: guestBoundary,
	})
	if err != nil {
		return err
	}
	if !controlledDispatch.ControlledDispatchRequestCreated {
		return encode(stdout, controlledDispatch)
	}
	if !bridge.DispatchSmokeRequestMaterialized {
		return encode(stdout, bridge)
	}

	timeout, err := time.ParseDuration(timeoutText)
	if err != nil {
		return fmt.Errorf("parse timeout: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	result, err := winapp.RunKnownPortableDispatchSmoke(context.Background(), winapp.KnownDispatchSmokeRequest{
		AppID:         appID,
		CacheRoot:     cacheRoot,
		Arguments:     []string(appArgs),
		GuestBoundary: guestBoundary,
		Host:          host,
		Port:          parsedPort,
		User:          user,
		KeyPath:       keyPath,
		RemoteDir:     remoteDir,
		SSHPath:       sshPath,
		SCPPath:       scpPath,
		Timeout:       timeout,
	})
	if err != nil {
		return err
	}
	return encode(stdout, result)
}

func encode(stdout io.Writer, payload any) error {
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(payload)
}

type repeatedStringFlag []string

func (flag *repeatedStringFlag) String() string {
	return fmt.Sprint([]string(*flag))
}

func (flag *repeatedStringFlag) Set(value string) error {
	*flag = append(*flag, value)
	return nil
}
