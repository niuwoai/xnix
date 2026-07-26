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
	launch := launchOptions{LauncherBin: envOrDefault(launcherBinEnv, "xnix-compat-launch")}
	flags.StringVar(&applicationID, "app", "", "application id to use for the file-open preview")
	flags.StringVar(&registryPath, "registry", "", "path to an Xnix recipe registry JSON file")
	flags.StringVar(&recipeDir, "recipe-dir", packagedRecipeRegistryDir, "directory containing the default Xnix recipe registry")
	flags.StringVar(&recipeRoot, "recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	flags.StringVar(&activationRoot, "activation-root", "", "read staged desktop activation receipt evidence from this explicit root")
	flags.BoolVar(&launch.Execute, "execute", false, "delegate the selected file-open request to the managed Runtime launcher")
	flags.StringVar(&launch.LauncherBin, "launcher-bin", launch.LauncherBin, "managed Runtime launcher command")
	flags.StringVar(&launch.LauncherRegistry, "launcher-registry", "", "optional recipe registry path passed through to the managed launcher")
	flags.StringVar(&launch.CacheRoot, "cache-root", "", "managed known Windows app cache root passed to the launcher")
	flags.StringVar(&launch.GuestBoundary, "guest-boundary", "", "controlled managed guest boundary passed to the launcher")
	flags.StringVar(&launch.StateRoot, "state-root", "", "controlled Runtime state root passed to the launcher")
	flags.StringVar(&launch.ReceiptID, "receipt-id", "", "opaque launch authorization receipt id passed to the launcher")
	flags.StringVar(&launch.ReviewReceiptID, "review-receipt-id", "", "opaque session-gated review receipt id passed to the launcher")
	flags.StringVar(&launch.SessionID, "session-id", "", "optional opaque controlled execution session id passed to the launcher")
	flags.StringVar(&launch.Host, "host", "", "guest SSH host passed to the launcher")
	flags.StringVar(&launch.Port, "port", "", "guest SSH port passed to the launcher")
	flags.StringVar(&launch.User, "user", "", "guest SSH user passed to the launcher")
	flags.StringVar(&launch.KeyPath, "key", "", "guest SSH private key path passed to the launcher")
	flags.StringVar(&launch.RemoteDir, "remote-dir", "", "guest remote work directory passed to the launcher")
	flags.StringVar(&launch.SSHPath, "ssh", "", "explicit ssh client passed to the launcher")
	flags.StringVar(&launch.SCPPath, "scp", "", "explicit scp client passed to the launcher")
	flags.StringVar(&launch.XWinInfoPath, "xwininfo", "", "explicit xwininfo client passed to the launcher")
	flags.StringVar(&launch.GuestDisplay, "guest-display", "", "guest DISPLAY value passed to the launcher")
	flags.StringVar(&launch.HostDisplay, "host-display", "", "host DISPLAY value passed to the launcher")
	flags.StringVar(&launch.WindowMatch, "window-match", "", "case-insensitive X window title/text required by the launcher")
	flags.StringVar(&launch.Timeout, "timeout", "", "launcher execution timeout")
	flags.StringVar(&launch.GUIWait, "gui-wait", "", "launcher GUI observation wait")
	if err := flags.Parse(args); err != nil {
		return err
	}

	fileURIs := flags.Args()
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
		return runManagedLauncher(stdout, preview, fileURIs, launch)
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runManagedLauncher(stdout io.Writer, preview appidentity.FileOpenPreview, fileURIs []string, options launchOptions) error {
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

func envOrDefault(name string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}
