package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const commandName = "xnix-compat-open"
const packagedRecipeRegistryDir = "/usr/share/xnix/compatibility/recipes"
const launcherBinEnv = "XNIX_COMPAT_LAUNCH"
const registryPathEnv = "XNIX_COMPAT_OPEN_REGISTRY"
const executeEnv = "XNIX_COMPAT_OPEN_EXECUTE"

const (
	launcherRegistryEnv = "XNIX_COMPAT_OPEN_LAUNCHER_REGISTRY"
	cacheRootEnv        = "XNIX_COMPAT_OPEN_CACHE_ROOT"
	guestBoundaryEnv    = "XNIX_COMPAT_OPEN_GUEST_BOUNDARY"
	stateRootEnv        = "XNIX_COMPAT_OPEN_STATE_ROOT"
	receiptIDEnv        = "XNIX_COMPAT_OPEN_RECEIPT_ID"
	reviewReceiptIDEnv  = "XNIX_COMPAT_OPEN_REVIEW_RECEIPT_ID"
	sessionIDEnv        = "XNIX_COMPAT_OPEN_SESSION_ID"
	hostEnv             = "XNIX_COMPAT_OPEN_GUEST_HOST"
	portEnv             = "XNIX_COMPAT_OPEN_GUEST_PORT"
	userEnv             = "XNIX_COMPAT_OPEN_GUEST_USER"
	keyPathEnv          = "XNIX_COMPAT_OPEN_GUEST_KEY"
	remoteDirEnv        = "XNIX_COMPAT_OPEN_GUEST_REMOTE_DIR"
	sshPathEnv          = "XNIX_COMPAT_OPEN_GUEST_SSH"
	scpPathEnv          = "XNIX_COMPAT_OPEN_GUEST_SCP"
	xwininfoPathEnv     = "XNIX_COMPAT_OPEN_GUEST_XWININFO"
	guestDisplayEnv     = "XNIX_COMPAT_OPEN_GUEST_DISPLAY"
	hostDisplayEnv      = "XNIX_COMPAT_OPEN_HOST_DISPLAY"
	windowMatchEnv      = "XNIX_COMPAT_OPEN_WINDOW_MATCH"
	timeoutEnv          = "XNIX_COMPAT_OPEN_TIMEOUT"
	guiWaitEnv          = "XNIX_COMPAT_OPEN_GUI_WAIT"
	executableEnv       = "XNIX_COMPAT_OPEN_EXECUTABLE"
)

type launchOptions struct {
	Execute                bool
	LauncherBin            string
	LauncherRegistry       string
	CacheRoot              string
	GuestBoundary          string
	StateRoot              string
	ReceiptID              string
	ReviewReceiptID        string
	SessionID              string
	Host                   string
	Port                   string
	User                   string
	KeyPath                string
	RemoteDir              string
	SSHPath                string
	SCPPath                string
	XWinInfoPath           string
	GuestDisplay           string
	HostDisplay            string
	WindowMatch            string
	Timeout                string
	GUIWait                string
	ExecutablePath         string
	FileArguments          repeatedStringFlag
	ExternalDesktopAppMode bool
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", commandName, err)
		os.Exit(64)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	var applicationID string
	var registryPath string
	var recipeDir string
	var recipeRoot string
	var activationRoot string
	launch := launchOptions{
		Execute:          os.Getenv(executeEnv) == "1",
		LauncherBin:      envOrDefault(launcherBinEnv, "xnix-compat-launch"),
		LauncherRegistry: envOrDefault(launcherRegistryEnv, ""),
		CacheRoot:        envOrDefault(cacheRootEnv, ""),
		GuestBoundary:    envOrDefault(guestBoundaryEnv, ""),
		StateRoot:        envOrDefault(stateRootEnv, ""),
		ReceiptID:        envOrDefault(receiptIDEnv, ""),
		ReviewReceiptID:  envOrDefault(reviewReceiptIDEnv, ""),
		SessionID:        envOrDefault(sessionIDEnv, ""),
		Host:             envOrDefault(hostEnv, ""),
		Port:             envOrDefault(portEnv, ""),
		User:             envOrDefault(userEnv, ""),
		KeyPath:          envOrDefault(keyPathEnv, ""),
		RemoteDir:        envOrDefault(remoteDirEnv, ""),
		SSHPath:          envOrDefault(sshPathEnv, ""),
		SCPPath:          envOrDefault(scpPathEnv, ""),
		XWinInfoPath:     envOrDefault(xwininfoPathEnv, ""),
		GuestDisplay:     envOrDefault(guestDisplayEnv, ""),
		HostDisplay:      envOrDefault(hostDisplayEnv, ""),
		WindowMatch:      envOrDefault(windowMatchEnv, ""),
		Timeout:          envOrDefault(timeoutEnv, ""),
		GUIWait:          envOrDefault(guiWaitEnv, ""),
		ExecutablePath:   envOrDefault(executableEnv, ""),
	}
	flags.StringVar(&applicationID, "app", "", "application id to use for the file-open preview")
	flags.StringVar(&registryPath, "registry", envOrDefault(registryPathEnv, ""), "path to an Xnix recipe registry JSON file")
	flags.StringVar(&recipeDir, "recipe-dir", packagedRecipeRegistryDir, "directory containing the default Xnix recipe registry")
	flags.StringVar(&recipeRoot, "recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	flags.StringVar(&activationRoot, "activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	flags.BoolVar(&launch.Execute, "execute", launch.Execute, "delegate the selected file-open request to the managed Runtime launcher")
	flags.StringVar(&launch.LauncherBin, "launcher-bin", launch.LauncherBin, "managed Runtime launcher command")
	flags.StringVar(&launch.LauncherRegistry, "launcher-registry", launch.LauncherRegistry, "optional recipe registry path passed through to the managed launcher")
	flags.StringVar(&launch.CacheRoot, "cache-root", launch.CacheRoot, "managed known Windows app cache root passed to the launcher")
	flags.StringVar(&launch.GuestBoundary, "guest-boundary", launch.GuestBoundary, "controlled managed guest boundary passed to the launcher")
	flags.StringVar(&launch.StateRoot, "state-root", launch.StateRoot, "controlled Runtime state root passed to the launcher")
	flags.StringVar(&launch.ReceiptID, "receipt-id", launch.ReceiptID, "opaque launch authorization receipt id passed to the launcher")
	flags.StringVar(&launch.ReviewReceiptID, "review-receipt-id", launch.ReviewReceiptID, "opaque session-gated review receipt id passed to the launcher")
	flags.StringVar(&launch.SessionID, "session-id", launch.SessionID, "optional opaque controlled execution session id passed to the launcher")
	flags.StringVar(&launch.Host, "host", launch.Host, "guest SSH host passed to the launcher")
	flags.StringVar(&launch.Port, "port", launch.Port, "guest SSH port passed to the launcher")
	flags.StringVar(&launch.User, "user", launch.User, "guest SSH user passed to the launcher")
	flags.StringVar(&launch.KeyPath, "key", launch.KeyPath, "guest SSH private key path passed to the launcher")
	flags.StringVar(&launch.RemoteDir, "remote-dir", launch.RemoteDir, "guest remote work directory passed to the launcher")
	flags.StringVar(&launch.SSHPath, "ssh", launch.SSHPath, "explicit ssh client passed to the launcher")
	flags.StringVar(&launch.SCPPath, "scp", launch.SCPPath, "explicit scp client passed to the launcher")
	flags.StringVar(&launch.XWinInfoPath, "xwininfo", launch.XWinInfoPath, "explicit xwininfo client passed to the launcher")
	flags.StringVar(&launch.GuestDisplay, "guest-display", launch.GuestDisplay, "guest DISPLAY value passed to the launcher")
	flags.StringVar(&launch.HostDisplay, "host-display", launch.HostDisplay, "host DISPLAY value passed to the launcher")
	flags.StringVar(&launch.WindowMatch, "window-match", launch.WindowMatch, "case-insensitive X window title/text required by the launcher")
	flags.StringVar(&launch.Timeout, "timeout", launch.Timeout, "launcher execution timeout")
	flags.StringVar(&launch.GUIWait, "gui-wait", launch.GUIWait, "launcher GUI observation wait")
	flags.StringVar(&launch.ExecutablePath, "executable", launch.ExecutablePath, "Runtime-owner supplied Windows GUI executable copied into the guest")
	flags.Var(&launch.FileArguments, "file-argument", "local file path delegated by the Runtime owner and passed to the managed launcher; may be repeated")
	if err := flags.Parse(args); err != nil {
		return err
	}

	positionalFileURIs := flags.Args()
	fileURIs, err := fileOpenURIs(positionalFileURIs, []string(launch.FileArguments), launch.Execute)
	if err != nil {
		return err
	}
	if len(fileURIs) == 0 {
		return fmt.Errorf("%s requires at least one file URI", commandName)
	}
	registry := strings.TrimSpace(registryPath)
	if registry == "" {
		registry = filepath.Join(strings.TrimSpace(recipeDir), "registry.json")
	}
	if strings.TrimSpace(registry) == "" {
		return fmt.Errorf("%s requires --registry or --recipe-dir", commandName)
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(registry, strings.TrimSpace(recipeRoot))
	if err != nil {
		return err
	}
	preview, err := appidentity.NewFileOpenPreviewWithOptions(
		recipes,
		provenance,
		fileURIs,
		strings.TrimSpace(applicationID),
		appidentity.FileOpenOptions{ActivationRoot: strings.TrimSpace(activationRoot)},
	)
	if err != nil {
		return err
	}

	if launch.Execute {
		return runManagedLauncher(stdout, preview, fileURIs, nil, launch)
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runManagedLauncher(stdout io.Writer, preview appidentity.FileOpenPreview, fileURIs []string, fileArgumentPaths []string, options launchOptions) error {
	launcher := strings.TrimSpace(options.LauncherBin)
	if launcher == "" {
		return fmt.Errorf("%s --execute requires --launcher-bin", commandName)
	}
	argv := []string{"--app", preview.ApplicationID}
	appendOptionalFlag := func(name string, value string) {
		if strings.TrimSpace(value) != "" {
			argv = append(argv, name, strings.TrimSpace(value))
		}
	}
	appendOptionalFlag("--registry", options.LauncherRegistry)
	appendOptionalFlag("--cache-root", options.CacheRoot)
	appendOptionalFlag("--guest-boundary", options.GuestBoundary)
	appendOptionalFlag("--state-root", options.StateRoot)
	appendOptionalFlag("--receipt-id", options.ReceiptID)
	appendOptionalFlag("--review-receipt-id", options.ReviewReceiptID)
	appendOptionalFlag("--session-id", options.SessionID)
	appendOptionalFlag("--host", options.Host)
	appendOptionalFlag("--port", options.Port)
	appendOptionalFlag("--user", options.User)
	appendOptionalFlag("--key", options.KeyPath)
	appendOptionalFlag("--remote-dir", options.RemoteDir)
	appendOptionalFlag("--ssh", options.SSHPath)
	appendOptionalFlag("--scp", options.SCPPath)
	appendOptionalFlag("--xwininfo", options.XWinInfoPath)
	appendOptionalFlag("--guest-display", options.GuestDisplay)
	appendOptionalFlag("--host-display", options.HostDisplay)
	appendOptionalFlag("--window-match", options.WindowMatch)
	appendOptionalFlag("--timeout", options.Timeout)
	appendOptionalFlag("--gui-wait", options.GUIWait)
	appendOptionalFlag("--executable", options.ExecutablePath)
	for _, fileURI := range fileURIs {
		path, err := localFilePath(fileURI)
		if err != nil {
			return err
		}
		argv = append(argv, "--file-argument", path)
	}

	command := exec.Command(launcher, argv...)
	var stderr bytes.Buffer
	command.Stdout = stdout
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		detail := strings.TrimSpace(stderr.String())
		if detail == "" {
			return fmt.Errorf("managed launcher failed: %w", err)
		}
		return fmt.Errorf("managed launcher failed: %w: %s", err, detail)
	}
	return nil
}

func fileOpenURIs(positional []string, fileArgumentPaths []string, execute bool) ([]string, error) {
	if len(fileArgumentPaths) > 0 && !execute {
		return nil, fmt.Errorf("%s --file-argument requires --execute", commandName)
	}
	uris := append([]string{}, positional...)
	for _, path := range fileArgumentPaths {
		cleaned, err := localFileArgumentPath(path)
		if err != nil {
			return nil, err
		}
		uris = append(uris, filePathURI(cleaned))
	}
	return uris, nil
}

func localFilePath(fileURI string) (string, error) {
	parsed, err := url.Parse(fileURI)
	if err != nil {
		return "", fmt.Errorf("invalid file URI: %q", fileURI)
	}
	if parsed.Scheme != "file" {
		return "", fmt.Errorf("only file URIs are accepted for launcher execution")
	}
	if parsed.Host != "" && parsed.Host != "localhost" {
		return "", fmt.Errorf("file URI host must be empty or localhost for launcher execution")
	}
	path, err := url.PathUnescape(parsed.EscapedPath())
	if err != nil {
		return "", fmt.Errorf("invalid file URI path escape: %w", err)
	}
	if path == "" || !filepath.IsAbs(path) {
		return "", fmt.Errorf("file URI must include an absolute local path for launcher execution")
	}
	return path, nil
}

func localFileArgumentPath(path string) (string, error) {
	cleaned := strings.TrimSpace(path)
	if cleaned == "" {
		return "", fmt.Errorf("%s --file-argument requires a non-empty path", commandName)
	}
	if strings.ContainsAny(cleaned, "\r\n") {
		return "", fmt.Errorf("%s --file-argument requires a single-line path", commandName)
	}
	if !filepath.IsAbs(cleaned) {
		return "", fmt.Errorf("%s --file-argument requires an absolute local path", commandName)
	}
	return cleaned, nil
}

func filePathURI(path string) string {
	return (&url.URL{Scheme: "file", Path: filepath.ToSlash(path)}).String()
}

func envOrDefault(name string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

type repeatedStringFlag []string

func (flag *repeatedStringFlag) String() string {
	return strings.Join(*flag, ",")
}

func (flag *repeatedStringFlag) Set(value string) error {
	*flag = append(*flag, value)
	return nil
}
