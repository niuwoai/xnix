package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

const launcherName = "xnix-compat-launch"
const stagingRootEnv = "XNIX_STAGING_ROOT"
const packagedRecipeRegistryDir = "/usr/share/xnix/compatibility/recipes"

const (
	compatLaunchDockerEnv                       = "XNIX_DOCKER_BIN"
	compatLaunchContainerImageEnv               = "XNIX_WINE_IMAGE"
	compatLaunchContainerPlatformEnv            = "XNIX_CONTAINER_PLATFORM"
	compatLaunchTimeoutEnv                      = "XNIX_COMPAT_LAUNCH_TIMEOUT"
	externalDesktopActivationRootEnv            = "XNIX_EXTERNAL_APP_DESKTOP_ACTIVATION_ROOT"
	externalDesktopLaunchPacketOutputEnv        = "XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_OUTPUT"
	externalDesktopLaunchPacketModeEnv          = "XNIX_EXTERNAL_APP_DESKTOP_LAUNCH_PACKET_MODE"
	externalDesktopWindowMatchEnv               = "XNIX_EXTERNAL_APP_WINDOW_MATCH"
	defaultExternalDesktopLaunchPacketModeValue = "development"
)

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
	var xwininfoPath string
	var executablePath string
	var guiAppPath string
	var guestDisplay string
	var hostDisplay string
	var windowMatch string
	var registryPath string
	var externalAppImportRecord string
	var externalAppHandle string
	var externalDesktopActivationRoot string
	var externalDesktopLaunchPacketOutput string
	var externalDesktopLaunchPacketMode string
	var externalDesktopWindowMatch string
	var containerImage string
	var containerPlatform string
	var containerDockerPath string
	var timeoutText string
	var guiWaitText string
	var fileArguments repeatedStringFlag
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
	flags.StringVar(&xwininfoPath, "xwininfo", "", "explicit xwininfo client path for guest GUI dispatch")
	flags.StringVar(&executablePath, "executable", "", "Runtime-owner supplied local Windows GUI executable copied into the guest")
	flags.StringVar(&guiAppPath, "gui-app", "", "Runtime-owner supplied guest GUI app path for guest GUI dispatch")
	flags.StringVar(&guestDisplay, "guest-display", "", "guest DISPLAY value for guest GUI dispatch")
	flags.StringVar(&hostDisplay, "host-display", "", "host DISPLAY value for GUI window observation")
	flags.StringVar(&windowMatch, "window-match", "", "case-insensitive X window title/text required for guest GUI dispatch")
	flags.StringVar(&registryPath, "registry", "", "digest-verified recipe registry path for recipe-backed container GUI dispatch")
	flags.StringVar(&externalAppImportRecord, "external-app-import-record", "", "Runtime import record for an imported external Windows GUI app")
	flags.StringVar(&externalAppHandle, "external-app-handle", "", "opaque external Windows GUI app handle; currently the imported reverse-DNS app id")
	flags.StringVar(&externalDesktopActivationRoot, "activation-root", envOrDefault(externalDesktopActivationRootEnv, ""), "staged desktop activation root used only when writing an external Windows app desktop launch packet sidecar")
	flags.StringVar(&externalDesktopLaunchPacketOutput, "desktop-launch-packet-output", envOrDefault(externalDesktopLaunchPacketOutputEnv, ""), "optional JSON sidecar path for a KDE-safe external Windows app desktop launch packet")
	flags.StringVar(&externalDesktopLaunchPacketMode, "desktop-launch-packet-mode", envOrDefault(externalDesktopLaunchPacketModeEnv, defaultExternalDesktopLaunchPacketModeValue), "desktop launch packet activation mode: production or development")
	flags.StringVar(&externalDesktopWindowMatch, "external-window-match", envOrDefault(externalDesktopWindowMatchEnv, ""), "case-insensitive X window match text for external Windows app desktop launches")
	flags.StringVar(&containerImage, "image", envOrDefault(compatLaunchContainerImageEnv, winapp.DefaultContainerImage), "local Wine X GUI container image for recipe-backed container GUI dispatch")
	flags.StringVar(&containerPlatform, "platform", envOrDefault(compatLaunchContainerPlatformEnv, ""), "container platform for recipe-backed container GUI dispatch; empty uses the local image platform")
	flags.StringVar(&containerDockerPath, "docker", envOrDefault(compatLaunchDockerEnv, ""), "explicit docker runner path for recipe-backed container GUI dispatch")
	flags.StringVar(&timeoutText, "timeout", envOrDefault(compatLaunchTimeoutEnv, winapp.DefaultKnownAppGuestTimeout.String()), "guest execution timeout")
	flags.StringVar(&guiWaitText, "gui-wait", "10s", "guest GUI observation wait")

	var appArgs repeatedStringFlag
	flags.Var(&appArgs, "arg", "argument passed to the known Windows app")
	flags.Var(&fileArguments, "file-argument", "local file copied into the guest and passed to a known Windows GUI app; may be repeated")

	if err := flags.Parse(args); err != nil {
		return err
	}
	externalAppMode := strings.TrimSpace(externalAppImportRecord) != "" || strings.TrimSpace(externalAppHandle) != ""
	if flags.NArg() != 0 && !externalAppMode {
		return fmt.Errorf("%s does not accept positional arguments", launcherName)
	}
	if externalAppMode {
		if strings.TrimSpace(appID) != "" {
			return fmt.Errorf("--external-app-import-record or --external-app-handle cannot be combined with --app")
		}
		if strings.TrimSpace(externalAppImportRecord) != "" && strings.TrimSpace(externalAppHandle) != "" {
			return fmt.Errorf("--external-app-import-record cannot be combined with --external-app-handle")
		}
		if strings.TrimSpace(externalDesktopLaunchPacketOutput) != "" && strings.TrimSpace(externalDesktopActivationRoot) == "" {
			return fmt.Errorf("--desktop-launch-packet-output requires --activation-root")
		}
		if strings.TrimSpace(externalDesktopActivationRoot) != "" && strings.TrimSpace(externalDesktopLaunchPacketOutput) == "" {
			return fmt.Errorf("--activation-root requires --desktop-launch-packet-output")
		}
		timeout, err := time.ParseDuration(timeoutText)
		if err != nil {
			return fmt.Errorf("parse timeout: %w", err)
		}
		result, err := appidentity.RunExternalWinApp(context.Background(), appidentity.ExternalWinAppRunRequest{
			ImportRecordPath:         externalAppImportRecord,
			StateRoot:                stateRoot,
			ExternalAppHandle:        externalAppHandle,
			ExternalDesktopArguments: flags.Args(),
			WindowMatch:              externalDesktopWindowMatch,
			Image:                    containerImage,
			Platform:                 containerPlatform,
			DockerPath:               containerDockerPath,
			Timeout:                  timeout,
		})
		if err != nil {
			return err
		}
		if strings.TrimSpace(externalDesktopLaunchPacketOutput) != "" {
			if err := writeExternalDesktopLaunchPacketSidecar(result, externalDesktopActivationRoot, externalDesktopLaunchPacketOutput, externalDesktopLaunchPacketMode); err != nil {
				return err
			}
		}
		return encode(stdout, result)
	}
	if appID == "" {
		return fmt.Errorf("--app is required")
	}
	app, err := winapp.LookupKnownPortableApp(appID)
	if err != nil {
		return err
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
	guiWait, err := time.ParseDuration(guiWaitText)
	if err != nil {
		return fmt.Errorf("parse GUI wait: %w", err)
	}
	parsedPort, err := winapp.ParseGuestPort(port)
	if err != nil {
		return err
	}

	if app.RecipeBackedContainerGUI && strings.TrimSpace(registryPath) != "" {
		resolvedRegistryPath, err := resolveStagedRegistryPath(registryPath)
		if err != nil {
			return err
		}
		recipe, _, err := appidentity.LoadRecipeFromRegistry(resolvedRegistryPath, "", appID)
		if err != nil {
			return err
		}
		if err := recipe.Validate(); err != nil {
			return err
		}
		if recipe.ContainerGUISmoke == (appidentity.ContainerGUISmokeHints{}) {
			return fmt.Errorf("recipe %s does not define container_gui_smoke hints", recipe.ID)
		}
		result, err := winapp.RunContainerXGUISmoke(context.Background(), winapp.ContainerXGUIRequest{
			ApplicationName: recipe.ContainerGUISmoke.App,
			WindowMatch:     recipe.ContainerGUISmoke.WindowMatch,
			ApplicationID:   recipe.ID,
			DisplayName:     recipe.Name,
			AppVersion:      recipe.Version,
			RecipeBacked:    true,
			Image:           containerImage,
			Platform:        containerPlatform,
			DockerPath:      containerDockerPath,
			Timeout:         timeout,
		})
		if err != nil {
			return err
		}
		return encode(stdout, recipeBackedContainerGUILaunchResult{
			ContainerXGUIResult:                    result,
			EvidenceSource:                         "winapp-smoke-container-x-gui",
			DispatchGate:                           winapp.KnownDispatchGuestBoundary,
			DispatchStarted:                        result.RunnerAvailable,
			ExecutionStarted:                       result.RunnerAvailable && result.ImageAvailable,
			SmokePassed:                            result.Status == winapp.PassedStatus,
			RuntimeOwnedDispatch:                   true,
			KDEPresentationOnly:                    true,
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
			ControlledSessionLiveStateObserved:     result.Status == winapp.PassedStatus,
			ControlledSessionRegistered:            controlledSession.SessionRegistered,
			ControlledSessionWindowObserved:        result.XWindowObserved,
			ControlledSessionHostRootModified:      controlledSession.HostRootModified || result.HostRootModified,
			ControlledSessionContainerProcessStart: result.XServerStarted && result.WineBootstrapAttempted,
			RawCommandExposed:                      false,
			BackendDetailsExposed:                  false,
		})
	}

	if app.GuestBuiltinGUI {
		result, err := winapp.RunKnownPortableGuestGUIDispatchSmoke(context.Background(), winapp.KnownDispatchSmokeRequest{
			AppID:             appID,
			CacheRoot:         cacheRoot,
			Arguments:         []string(appArgs),
			GuestBoundary:     guestBoundary,
			Host:              host,
			Port:              parsedPort,
			User:              user,
			KeyPath:           keyPath,
			RemoteDir:         remoteDir,
			SSHPath:           sshPath,
			SCPPath:           scpPath,
			XWinInfoPath:      xwininfoPath,
			ExecutablePath:    executablePath,
			GUIAppPath:        guiAppPath,
			GuestDisplay:      guestDisplay,
			HostDisplay:       hostDisplay,
			FileArgumentPaths: []string(fileArguments),
			WindowMatch:       windowMatch,
			Timeout:           timeout,
			Wait:              guiWait,
		})
		if err != nil {
			return err
		}
		return encode(stdout, launcherDispatchResult{
			KnownDispatchSmokeResult:               result,
			EvidenceSource:                         "wine-guest-gui-smoke",
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
			ControlledSessionLiveStateObserved:     result.SmokePassed,
			ControlledSessionRegistered:            controlledSession.SessionRegistered,
			ControlledSessionWindowObserved:        result.SmokePassed,
			ControlledSessionHostRootModified:      controlledSession.HostRootModified || result.HostRootModified,
			ControlledSessionBackendProcessStart:   result.BackendProcessStarted,
		})
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

func writeExternalDesktopLaunchPacketSidecar(result appidentity.ExternalWinAppRunResult, activationRoot string, outputPath string, mode string) error {
	plan, err := appidentity.ExternalAppPlanFromRunResult(result)
	if err != nil {
		return err
	}
	packet, err := plan.DesktopExternalWinAppLaunchPacketPreviewFromRunResult(activationRoot, result, mode)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(packet); err != nil {
		return err
	}
	cleanedOutputPath := filepath.Clean(strings.TrimSpace(outputPath))
	if err := os.MkdirAll(filepath.Dir(cleanedOutputPath), 0o755); err != nil {
		return fmt.Errorf("create external desktop launch packet output directory: %w", err)
	}
	if err := os.WriteFile(cleanedOutputPath, buffer.Bytes(), 0o600); err != nil {
		return fmt.Errorf("write external desktop launch packet: %w", err)
	}
	return nil
}

func envOrDefault(name string, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); value != "" {
		return value
	}
	return fallback
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

func resolveStagedRegistryPath(registryPath string) (string, error) {
	trimmedPath := strings.TrimSpace(registryPath)
	info, err := os.Stat(trimmedPath)
	if err == nil {
		if info.IsDir() {
			return "", fmt.Errorf("registry path must reference a file")
		}
		return trimmedPath, nil
	}
	if !os.IsNotExist(err) {
		return "", fmt.Errorf("inspect registry path: %w", err)
	}

	stagingRoot := strings.TrimSpace(os.Getenv(stagingRootEnv))
	if stagingRoot == "" {
		return trimmedPath, nil
	}
	cleanRegistryPath := filepath.Clean(trimmedPath)
	if !filepath.IsAbs(cleanRegistryPath) {
		return trimmedPath, nil
	}
	cleanRegistryPathSlash := filepath.ToSlash(cleanRegistryPath)
	if cleanRegistryPathSlash != packagedRecipeRegistryDir+"/registry.json" &&
		!strings.HasPrefix(cleanRegistryPathSlash, packagedRecipeRegistryDir+"/") {
		return trimmedPath, nil
	}

	absoluteStagingRoot, err := filepath.Abs(stagingRoot)
	if err != nil {
		return "", fmt.Errorf("resolve staging root: %w", err)
	}
	candidatePath := filepath.Join(absoluteStagingRoot, strings.TrimPrefix(cleanRegistryPath, string(filepath.Separator)))
	relativeCandidate, err := filepath.Rel(absoluteStagingRoot, candidatePath)
	if err != nil {
		return "", fmt.Errorf("resolve staged registry path: %w", err)
	}
	if relativeCandidate == ".." || strings.HasPrefix(filepath.ToSlash(relativeCandidate), "../") {
		return "", fmt.Errorf("staged registry path escaped staging root")
	}
	info, err = os.Stat(candidatePath)
	if err != nil {
		if os.IsNotExist(err) {
			return trimmedPath, nil
		}
		return "", fmt.Errorf("inspect staged registry path: %w", err)
	}
	if info.IsDir() {
		return "", fmt.Errorf("staged registry path must reference a file")
	}
	return candidatePath, nil
}

type launcherDispatchResult struct {
	winapp.KnownDispatchSmokeResult
	EvidenceSource                         string `json:"evidence_source,omitempty"`
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

type recipeBackedContainerGUILaunchResult struct {
	winapp.ContainerXGUIResult
	EvidenceSource                         string `json:"evidence_source"`
	DispatchGate                           string `json:"dispatch_gate"`
	DispatchStarted                        bool   `json:"dispatch_started"`
	ExecutionStarted                       bool   `json:"execution_started"`
	SmokePassed                            bool   `json:"smoke_passed"`
	RuntimeOwnedDispatch                   bool   `json:"runtime_owned_dispatch"`
	KDEPresentationOnly                    bool   `json:"kde_presentation_only"`
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
	ControlledSessionContainerProcessStart bool   `json:"controlled_session_container_process_start"`
	RawCommandExposed                      bool   `json:"raw_command_exposed"`
	BackendDetailsExposed                  bool   `json:"backend_details_exposed"`
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
