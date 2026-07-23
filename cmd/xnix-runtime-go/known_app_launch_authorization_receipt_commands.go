package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func runKnownAppLaunchAuthorizationReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-launch-authorization-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root for the opaque launch authorization receipt")
	authorize := flags.String("authorize", "", "exact launch authorization action id")
	evidenceSource := flags.String("evidence-source", "staged-launcher-dispatch-smoke", "redacted known-app smoke evidence source")
	centerCardState := flags.String("center-card-state", "validated-launch-authorization-required", "Compatibility Center card state that requested authorization")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-launch-authorization-receipt-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-launch-authorization-receipt-preview does not accept positional arguments")
	}
	preview, err := appidentity.RecordKnownAppLaunchAuthorizationReceipt(appidentity.KnownAppLaunchAuthorizationReceiptRequest{
		AppID:           *appID,
		StateRoot:       *stateRoot,
		Authorize:       *authorize,
		EvidenceSource:  *evidenceSource,
		CenterCardState: *centerCardState,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppLaunchGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-launch-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-launch-gate-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-launch-gate-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppLaunchGate(appidentity.KnownAppLaunchGateRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledDispatchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-dispatch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-dispatch-request-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-dispatch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledDispatchRequest(appidentity.KnownAppControlledDispatchRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledExecutionSessionPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-execution-session-preview requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledExecutionSession(appidentity.KnownAppControlledExecutionSessionRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppControlledExecutionSessionRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the opaque launch authorization receipt")
	receiptID := flags.String("receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *receiptID == "" {
		return errors.New("known-app-controlled-execution-session-record requires --state-root and --receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-record does not accept positional arguments")
	}
	record, err := appidentity.RecordKnownAppControlledExecutionSession(appidentity.KnownAppControlledExecutionSessionRecordRequest{
		AppID:         *appID,
		StateRoot:     *stateRoot,
		ReceiptID:     *receiptID,
		CacheRoot:     *cacheRoot,
		GuestBoundary: *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}

func runKnownAppControlledExecutionSessionConsumePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-controlled-execution-session-consume-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-controlled-execution-session-consume-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-controlled-execution-session-consume-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppControlledExecutionSessionConsumption(appidentity.KnownAppControlledExecutionSessionConsumeRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedLaunchReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	actionID := flags.String("action", appidentity.KnownAppSessionGatedLaunchReviewAction, "exact session-gated launch review action id")
	decision := flags.String("decision", "reviewed", "review decision: reviewed, approved, deferred, or rejected")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedLaunchReview(appidentity.KnownAppSessionGatedLaunchReviewRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ActionID:  *actionID,
		Decision:  *decision,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedLaunchReviewReceiptRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-receipt-record", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted execution session record")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	actionID := flags.String("action", appidentity.KnownAppSessionGatedLaunchReviewAction, "exact session-gated launch review action id")
	decision := flags.String("decision", "reviewed", "review decision: reviewed, approved, deferred, or rejected")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-receipt-record requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-receipt-record does not accept positional arguments")
	}
	record, err := appidentity.RecordKnownAppSessionGatedLaunchReviewReceipt(appidentity.KnownAppSessionGatedLaunchReviewReceiptRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ActionID:  *actionID,
		Decision:  *decision,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, record)
}

func runKnownAppSessionGatedLaunchReviewGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-launch-review-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing the persisted session-gated launch review receipt")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	receiptID := flags.String("receipt-id", "", "optional opaque session-gated launch review receipt id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" {
		return errors.New("known-app-session-gated-launch-review-gate-preview requires --state-root")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-launch-review-gate-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedLaunchReviewGate(appidentity.KnownAppSessionGatedLaunchReviewGateRequest{
		AppID:     *appID,
		StateRoot: *stateRoot,
		SessionID: *sessionID,
		ReceiptID: *receiptID,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppSessionGatedControlledDispatchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-session-gated-controlled-dispatch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "controlled Runtime state root containing launch and review receipts")
	sessionID := flags.String("session-id", "", "optional opaque controlled execution session id")
	reviewReceiptID := flags.String("review-receipt-id", "", "opaque session-gated launch review receipt id")
	launchReceiptID := flags.String("launch-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *stateRoot == "" || *reviewReceiptID == "" || *launchReceiptID == "" {
		return errors.New("known-app-session-gated-controlled-dispatch-request-preview requires --state-root, --review-receipt-id, and --launch-receipt-id")
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-session-gated-controlled-dispatch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppSessionGatedControlledDispatchRequest(appidentity.KnownAppSessionGatedControlledDispatchRequest{
		AppID:           *appID,
		StateRoot:       *stateRoot,
		SessionID:       *sessionID,
		ReviewReceiptID: *reviewReceiptID,
		LaunchReceiptID: *launchReceiptID,
		CacheRoot:       *cacheRoot,
		GuestBoundary:   *guestBoundary,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppKDERuntimeStatusLaunchRequestPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("known-app-kde-runtime-status-launch-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	launchReceiptID := flags.String("launch-authorization-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	reviewReceiptID := flags.String("session-gated-review-receipt-id", "", "opaque known Windows app session-gated review receipt id")
	sessionID := flags.String("session-id", "", "opaque known Windows app controlled execution session id")
	centerCardState := flags.String("center-card-state", "", "Compatibility Center card state from the KDE read model")
	primaryActionID := flags.String("primary-action-id", "", "KDE primary action id from the Runtime-owned card")
	postReviewDispatchState := flags.String("post-review-dispatch-state", "", "post-review dispatch state from the Runtime-owned card")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-kde-runtime-status-launch-request-preview does not accept positional arguments")
	}
	preview, err := appidentity.PreviewKnownAppKDERuntimeStatusLaunchRequest(appidentity.KnownAppKDERuntimeStatusLaunchRequest{
		AppID:                        *appID,
		LaunchAuthorizationReceiptID: *launchReceiptID,
		SessionGatedReviewReceiptID:  *reviewReceiptID,
		ControlledExecutionSessionID: *sessionID,
		CenterCardState:              *centerCardState,
		PrimaryActionID:              *primaryActionID,
		PostReviewDispatchState:      *postReviewDispatchState,
	})
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runKnownAppKDERuntimeStatusLaunchExecution(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	appID := flags.String("app", "", "known Windows application id")
	stateRoot := flags.String("state-root", "", "Runtime-owner supplied state root; KDE must not provide this value")
	cacheRoot := flags.String("cache-root", winapp.DefaultKnownAppCacheRoot, "managed known Windows app cache root")
	launchReceiptID := flags.String("launch-authorization-receipt-id", "", "opaque known Windows app launch authorization receipt id")
	reviewReceiptID := flags.String("session-gated-review-receipt-id", "", "opaque known Windows app session-gated review receipt id")
	sessionID := flags.String("session-id", "", "opaque known Windows app controlled execution session id")
	centerCardState := flags.String("center-card-state", "", "Compatibility Center card state from the KDE read model")
	primaryActionID := flags.String("primary-action-id", "", "KDE primary action id from the Runtime-owned card")
	postReviewDispatchState := flags.String("post-review-dispatch-state", "", "post-review dispatch state from the Runtime-owned card")
	launcherPath := flags.String("launcher", "xnix-compat-launch", "Runtime-managed launcher executable")
	host := flags.String("host", "", "optional guest SSH host forwarded to the managed launcher")
	port := flags.String("port", "", "optional guest SSH port forwarded to the managed launcher")
	user := flags.String("user", "", "optional guest SSH user forwarded to the managed launcher")
	keyPath := flags.String("key", "", "optional guest SSH private key forwarded to the managed launcher")
	remoteDir := flags.String("remote-dir", "", "optional guest remote directory forwarded to the managed launcher")
	sshPath := flags.String("ssh", "", "optional ssh client path forwarded to the managed launcher")
	scpPath := flags.String("scp", "", "optional scp client path forwarded to the managed launcher")
	timeoutText := flags.String("timeout", "", "optional guest execution timeout forwarded to the managed launcher")
	ownerTimeoutText := flags.String("owner-timeout", "5m", "Runtime-owner launcher invocation timeout")
	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app through the managed launcher")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("known-app-kde-runtime-status-launch-execution does not accept positional arguments")
	}
	plan, err := appidentity.PrepareKnownAppKDERuntimeStatusLaunchExecution(appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequest{
		AppID:                        *appID,
		StateRoot:                    *stateRoot,
		CacheRoot:                    *cacheRoot,
		LaunchAuthorizationReceiptID: *launchReceiptID,
		SessionGatedReviewReceiptID:  *reviewReceiptID,
		ControlledExecutionSessionID: *sessionID,
		CenterCardState:              *centerCardState,
		PrimaryActionID:              *primaryActionID,
		PostReviewDispatchState:      *postReviewDispatchState,
	})
	if err != nil {
		return err
	}
	extra, err := knownAppKDERuntimeStatusLaunchExecutionExtraArgs(knownAppKDERuntimeStatusLaunchExecutionExtraOptions{
		Host:      *host,
		Port:      *port,
		User:      *user,
		KeyPath:   *keyPath,
		RemoteDir: *remoteDir,
		SSHPath:   *sshPath,
		SCPPath:   *scpPath,
		Timeout:   *timeoutText,
		AppArgs:   []string(appArgs),
	})
	if err != nil {
		return err
	}
	launcherArgs, err := appidentity.KnownAppKDERuntimeStatusLaunchExecutionArgv(plan, *stateRoot, *cacheRoot, extra)
	if err != nil {
		return err
	}
	ownerTimeout, err := time.ParseDuration(strings.TrimSpace(*ownerTimeoutText))
	if err != nil {
		return fmt.Errorf("parse Runtime-owner timeout: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), ownerTimeout)
	defer cancel()
	output, err := exec.CommandContext(ctx, strings.TrimSpace(*launcherPath), launcherArgs...).Output()
	if ctx.Err() != nil {
		return errors.New("managed launcher invocation timed out")
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("managed launcher invocation failed with exit code %d", exitErr.ExitCode())
		}
		return errors.New("managed launcher invocation failed")
	}
	result, err := knownAppKDERuntimeStatusLaunchExecutionResultFromOutput(plan, *launcherPath, output)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, result)
}

type knownAppKDERuntimeStatusLaunchExecutionResult struct {
	appidentity.KnownAppKDERuntimeStatusLaunchExecutionPlan
	ManagedLauncherName             string `json:"managed_launcher_name"`
	ManagedLauncherInvoked          bool   `json:"managed_launcher_invoked"`
	ExistingManagedLauncherInvoked  bool   `json:"existing_managed_launcher_invoked"`
	LauncherExitCode                int    `json:"launcher_exit_code"`
	LauncherOutputJSONObserved      bool   `json:"launcher_output_json_observed"`
	DelegatedRequestType            string `json:"delegated_request_type"`
	DelegatedStatus                 string `json:"delegated_status"`
	DelegatedSmokePassed            bool   `json:"delegated_smoke_passed"`
	DelegatedExecutionStarted       bool   `json:"delegated_execution_started"`
	DelegatedBackendProcessStarted  bool   `json:"delegated_backend_process_started"`
	DelegatedHostRootModified       bool   `json:"delegated_host_root_modified"`
	DelegatedDockerSocketMounted    bool   `json:"delegated_docker_socket_mounted"`
	DelegatedBroadHostMountRequired bool   `json:"delegated_broad_host_mount_required"`
	DelegatedRawCommandExposed      bool   `json:"delegated_raw_command_exposed"`
	DelegatedBackendDetailsExposed  bool   `json:"delegated_backend_details_exposed"`
}

type knownAppKDERuntimeStatusLaunchExecutionExtraOptions struct {
	Host      string
	Port      string
	User      string
	KeyPath   string
	RemoteDir string
	SSHPath   string
	SCPPath   string
	Timeout   string
	AppArgs   []string
}

func knownAppKDERuntimeStatusLaunchExecutionExtraArgs(options knownAppKDERuntimeStatusLaunchExecutionExtraOptions) ([]string, error) {
	var args []string
	for _, pair := range []struct {
		flag  string
		value string
	}{
		{"--host", options.Host},
		{"--port", options.Port},
		{"--user", options.User},
		{"--key", options.KeyPath},
		{"--remote-dir", options.RemoteDir},
		{"--ssh", options.SSHPath},
		{"--scp", options.SCPPath},
		{"--timeout", options.Timeout},
	} {
		value := strings.TrimSpace(pair.value)
		if value == "" {
			continue
		}
		if strings.ContainsAny(value, "\r\n") {
			return nil, fmt.Errorf("%s requires a single-line value", pair.flag)
		}
		args = append(args, pair.flag, value)
	}
	for _, value := range options.AppArgs {
		if strings.ContainsAny(value, "\r\n") {
			return nil, errors.New("--arg requires single-line values")
		}
		args = append(args, "--arg", value)
	}
	return args, nil
}

func knownAppKDERuntimeStatusLaunchExecutionResultFromOutput(plan appidentity.KnownAppKDERuntimeStatusLaunchExecutionPlan, launcherPath string, output []byte) (knownAppKDERuntimeStatusLaunchExecutionResult, error) {
	var delegated map[string]any
	if err := json.Unmarshal(output, &delegated); err != nil {
		return knownAppKDERuntimeStatusLaunchExecutionResult{}, errors.New("managed launcher must emit JSON")
	}
	result := knownAppKDERuntimeStatusLaunchExecutionResult{
		KnownAppKDERuntimeStatusLaunchExecutionPlan: plan,
		ManagedLauncherName:                         filepath.Base(strings.TrimSpace(launcherPath)),
		ManagedLauncherInvoked:                      true,
		ExistingManagedLauncherInvoked:              true,
		LauncherExitCode:                            0,
		LauncherOutputJSONObserved:                  true,
		DelegatedRequestType:                        stringJSONField(delegated, "request_type"),
		DelegatedStatus:                             stringJSONField(delegated, "status"),
		DelegatedSmokePassed:                        boolJSONField(delegated, "smoke_passed"),
		DelegatedExecutionStarted:                   boolJSONField(delegated, "execution_started"),
		DelegatedBackendProcessStarted:              boolJSONField(delegated, "backend_process_started"),
		DelegatedHostRootModified:                   boolJSONField(delegated, "host_root_modified"),
		DelegatedDockerSocketMounted:                boolJSONField(delegated, "docker_socket_mounted"),
		DelegatedBroadHostMountRequired:             boolJSONField(delegated, "broad_host_mount_required"),
		DelegatedRawCommandExposed:                  boolJSONField(delegated, "raw_command_exposed"),
		DelegatedBackendDetailsExposed:              boolJSONField(delegated, "backend_details_exposed"),
	}
	if result.ManagedLauncherName == "" || strings.ContainsAny(result.ManagedLauncherName, "\r\n") {
		return knownAppKDERuntimeStatusLaunchExecutionResult{}, errors.New("managed launcher name is unsafe")
	}
	if result.DelegatedHostRootModified || result.DelegatedDockerSocketMounted || result.DelegatedBroadHostMountRequired || result.DelegatedRawCommandExposed || result.DelegatedBackendDetailsExposed {
		return knownAppKDERuntimeStatusLaunchExecutionResult{}, errors.New("managed launcher reported an unsafe delegated result")
	}
	return result, nil
}

func stringJSONField(payload map[string]any, key string) string {
	value, _ := payload[key].(string)
	return value
}

func boolJSONField(payload map[string]any, key string) bool {
	value, _ := payload[key].(bool)
	return value
}
