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
	var reviewReceiptID string
	var sessionID string
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
	flags.StringVar(&reviewReceiptID, "review-receipt-id", "", "opaque Runtime session-gated launch review receipt id")
	flags.StringVar(&sessionID, "session-id", "", "optional opaque controlled execution session id")
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
	if stateRoot == "" || receiptID == "" || reviewReceiptID == "" {
		return fmt.Errorf("--state-root, --receipt-id, and --review-receipt-id are required when --guest-boundary requests dispatch")
	}
	controlledDispatch, err := consumeSessionGatedControlledDispatchForLaunch(appID, stateRoot, sessionID, reviewReceiptID, receiptID, cacheRoot, guestBoundary)
	if err != nil {
		return err
	}
	if !controlledDispatch.ControlledDispatchRequestCreated {
		return encode(stdout, controlledDispatch)
	}
	if !bridge.DispatchSmokeRequestMaterialized {
		return encode(stdout, bridge)
	}
	controlledSession, err := consumeControlledExecutionSessionForLaunch(appID, stateRoot, sessionID)
	if err != nil {
		return err
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
	return encode(stdout, launcherDispatchResult{
		KnownDispatchSmokeResult:               result,
		SessionGatedControlledDispatchConsumed: controlledDispatch.ControlledDispatchRequestCreated,
		SessionGatedControlledDispatchState:    controlledDispatch.ControlledDispatchRequestState,
		SessionGatedReviewReceiptID:            controlledDispatch.ReviewReceiptID,
		LaunchAuthorizationReceiptID:           controlledDispatch.LaunchAuthorizationReceiptID,
		ControlledExecutionSessionConsumed:     controlledSession.RecordConsumed,
		ControlledExecutionSessionID:           controlledSession.ExecutionSessionID,
		ControlledSessionDigestVerified:        controlledSession.SessionDigestVerified,
		ControlledSessionRelativePath:          controlledSession.SessionRelativePath,
		RuntimeOwnerConsumableSession:          controlledSession.RuntimeOwnerConsumable,
		KDEReadModelConsumableSession:          controlledSession.KDEReadModelConsumable,
		ControlledSessionLiveStateObserved:     controlledSession.LiveStateObserved,
		ControlledSessionRegistered:            controlledSession.SessionRegistered,
		ControlledSessionWindowObserved:        controlledSession.WindowObserved,
		ControlledSessionHostRootModified:      controlledSession.HostRootModified,
		ControlledSessionBackendProcessStart:   controlledSession.BackendProcessStarted,
	})
}

func consumeControlledExecutionSessionForLaunch(appID string, stateRoot string, sessionID string) (appidentity.KnownAppControlledExecutionSessionConsumePreview, error) {
	preview, err := appidentity.PreviewKnownAppControlledExecutionSessionConsumption(appidentity.KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     appID,
		StateRoot: stateRoot,
		SessionID: sessionID,
	})
	if err != nil {
		return appidentity.KnownAppControlledExecutionSessionConsumePreview{}, fmt.Errorf("controlled execution session gate rejected dispatch: %w", err)
	}
	return preview, nil
}

func consumeSessionGatedControlledDispatchForLaunch(appID string, stateRoot string, sessionID string, reviewReceiptID string, launchReceiptID string, cacheRoot string, guestBoundary string) (appidentity.KnownAppSessionGatedControlledDispatchPreview, error) {
	preview, err := appidentity.PreviewKnownAppSessionGatedControlledDispatchRequest(appidentity.KnownAppSessionGatedControlledDispatchRequest{
		AppID:           appID,
		StateRoot:       stateRoot,
		SessionID:       sessionID,
		ReviewReceiptID: reviewReceiptID,
		LaunchReceiptID: launchReceiptID,
		CacheRoot:       cacheRoot,
		GuestBoundary:   guestBoundary,
	})
	if err != nil {
		return appidentity.KnownAppSessionGatedControlledDispatchPreview{}, fmt.Errorf("session-gated controlled dispatch gate rejected dispatch: %w", err)
	}
	return preview, nil
}

type launcherDispatchResult struct {
	winapp.KnownDispatchSmokeResult
	SessionGatedControlledDispatchConsumed bool   `json:"session_gated_controlled_dispatch_consumed"`
	SessionGatedControlledDispatchState    string `json:"session_gated_controlled_dispatch_state"`
	SessionGatedReviewReceiptID            string `json:"session_gated_review_receipt_id"`
	LaunchAuthorizationReceiptID           string `json:"launch_authorization_receipt_id"`
	ControlledExecutionSessionConsumed     bool   `json:"controlled_execution_session_consumed"`
	ControlledExecutionSessionID           string `json:"controlled_execution_session_id"`
	ControlledSessionDigestVerified        bool   `json:"controlled_session_digest_verified"`
	ControlledSessionRelativePath          string `json:"controlled_session_relative_path"`
	RuntimeOwnerConsumableSession          bool   `json:"runtime_owner_consumable_session"`
	KDEReadModelConsumableSession          bool   `json:"kde_read_model_consumable_session"`
	ControlledSessionLiveStateObserved     bool   `json:"controlled_session_live_state_observed"`
	ControlledSessionRegistered            bool   `json:"controlled_session_registered"`
	ControlledSessionWindowObserved        bool   `json:"controlled_session_window_observed"`
	ControlledSessionHostRootModified      bool   `json:"controlled_session_host_root_modified"`
	ControlledSessionBackendProcessStart   bool   `json:"controlled_session_backend_process_start"`
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
