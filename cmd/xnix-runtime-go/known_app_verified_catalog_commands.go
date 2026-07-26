package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func runKnownAppVerifiedCatalogPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogPreviewRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	matrixEvidence := flags.String("matrix-evidence", "", "Go-owned known app matrix evidence JSON path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogPreviewRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalog(appidentity.KnownAppVerifiedCatalogRequest{
		MatrixEvidencePath: *matrixEvidence,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKnownAppVerifiedCatalogRunPlanPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogRunPlanRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	verifiedCatalog := flags.String("verified-catalog", "", "Go-owned known app verified catalog JSON path")
	appID := flags.String("app", "", "known Windows app id to plan for q4 execution")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogRunPlanRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalogRunPlan(appidentity.KnownAppVerifiedCatalogRunPlanRequest{
		VerifiedCatalogPath: *verifiedCatalog,
		AppID:               *appID,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKnownAppVerifiedCatalogRunAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogRunAcceptanceRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	runPlan := flags.String("run-plan", "", "Go-owned verified catalog run plan JSON path")
	runReport := flags.String("known-winapp-run", "", "passed known Windows app run JSON")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogRunAcceptanceRequestType)
	}

	preview, err := appidentity.PreviewKnownAppVerifiedCatalogRunAcceptance(appidentity.KnownAppVerifiedCatalogRunAcceptanceRequest{
		RunPlanPath:   *runPlan,
		RunReportPath: *runReport,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKnownAppVerifiedCatalogLaunchHandoffRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogLaunchHandoffRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root where the verified catalog launch handoff is persisted")
	runAcceptance := flags.String("known-app-verified-catalog-run-acceptance", "", "Go-owned verified catalog app run acceptance JSON path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogLaunchHandoffRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogLaunchHandoff(appidentity.KnownAppVerifiedCatalogLaunchHandoffRequest{
		StateRoot:      *stateRoot,
		AcceptancePath: *runAcceptance,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runKnownAppVerifiedCatalogLaunchMaterializationRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root containing the verified catalog launch handoff")
	handoffRelativePath := flags.String("handoff-relative-path", "", "relative verified catalog launch handoff path forwarded by the desktop")
	cacheRoot := flags.String("cache-root", "", "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogLaunchMaterialization(appidentity.KnownAppVerifiedCatalogLaunchMaterializationRequest{
		StateRoot:           *stateRoot,
		HandoffRelativePath: *handoffRelativePath,
		CacheRoot:           *cacheRoot,
		GuestBoundary:       *guestBoundary,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runKnownAppVerifiedCatalogDispatchRequestRecord(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogDispatchRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root containing the verified catalog launch handoff")
	handoffRelativePath := flags.String("handoff-relative-path", "", "relative verified catalog launch handoff path forwarded by the desktop")
	cacheRoot := flags.String("cache-root", "", "managed known Windows app cache root")
	guestBoundary := flags.String("guest-boundary", "", "controlled managed guest boundary")
	launcherName := flags.String("launcher", "", "Runtime-owner managed launcher name or path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogDispatchRequestType)
	}

	record, err := appidentity.RecordKnownAppVerifiedCatalogDispatchRequest(appidentity.KnownAppVerifiedCatalogDispatchRequestRecordRequest{
		StateRoot:           *stateRoot,
		HandoffRelativePath: *handoffRelativePath,
		CacheRoot:           *cacheRoot,
		GuestBoundary:       *guestBoundary,
		LauncherName:        *launcherName,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(record)
}

func runKnownAppVerifiedCatalogDispatchRunnerExecution(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(appidentity.KnownAppVerifiedCatalogDispatchExecutionRequestType, flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	stateRoot := flags.String("state-root", "", "Runtime owner state root containing the verified catalog dispatch request")
	dispatchRequestRelativePath := flags.String("dispatch-request-relative-path", "", "relative verified catalog dispatch request path")
	expectedSHA256 := flags.String("expected-sha256", "", "optional expected dispatch request SHA256")
	launcherPath := flags.String("launcher", "", "Runtime-owner managed launcher path; defaults to the request launcher name")
	ownerTimeoutText := flags.String("owner-timeout", "5m", "Runtime-owner launcher invocation timeout")
	host := flags.String("host", winapp.DefaultGuestHost, "guest SSH host")
	port := flags.String("port", winapp.DefaultGuestPort, "guest SSH port")
	user := flags.String("user", winapp.DefaultGuestUser, "guest SSH user")
	keyPath := flags.String("key", "", "guest SSH private key path")
	remoteDir := flags.String("remote-dir", "/tmp/xnix-known-winapp-smoke", "guest remote smoke directory")
	sshPath := flags.String("ssh", "", "explicit ssh client path")
	scpPath := flags.String("scp", "", "explicit scp client path")
	startQEMU := flags.Bool("start-qemu", false, "start a controlled QEMU guest before invoking the dispatch runner")
	qemuBinary := flags.String("qemu-binary", winapp.DefaultQEMUBinary, "QEMU binary for controlled guest startup")
	qemuKernel := flags.String("qemu-kernel", "", "QEMU kernel image for controlled guest startup")
	qemuMemory := flags.String("qemu-memory", winapp.DefaultQEMUMemory, "QEMU guest memory")
	qemuCPUs := flags.String("qemu-cpus", winapp.DefaultQEMUCPUCount, "QEMU guest CPU count")
	qemuCPU := flags.String("qemu-cpu", winapp.DefaultQEMUCPUModel, "QEMU guest CPU model")
	qemuBootTimeoutText := flags.String("qemu-boot-timeout", winapp.DefaultQEMUBootTimeout.String(), "QEMU SSH boot timeout")
	qemuSerialLog := flags.String("qemu-serial-log", "", "QEMU serial log path")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("%s does not accept positional arguments", appidentity.KnownAppVerifiedCatalogDispatchExecutionRequestType)
	}

	plan, err := appidentity.ReadKnownAppVerifiedCatalogDispatchRequest(appidentity.KnownAppVerifiedCatalogDispatchRequestReadRequest{
		StateRoot:                   *stateRoot,
		DispatchRequestRelativePath: *dispatchRequestRelativePath,
		ExpectedSHA256:              *expectedSHA256,
	})
	if err != nil {
		return err
	}
	ownerTimeout, err := time.ParseDuration(strings.TrimSpace(*ownerTimeoutText))
	if err != nil {
		return fmt.Errorf("parse Runtime-owner timeout: %w", err)
	}
	qemuBootTimeout, err := time.ParseDuration(strings.TrimSpace(*qemuBootTimeoutText))
	if err != nil {
		return fmt.Errorf("parse QEMU boot timeout: %w", err)
	}
	execution, err := executeKnownAppVerifiedCatalogDispatchRequest(knownAppVerifiedCatalogDispatchExecutionRequest{
		Version:         readRuntimeGoProjectVersion(),
		Plan:            plan,
		LauncherPath:    *launcherPath,
		OwnerTimeout:    ownerTimeout,
		Host:            *host,
		Port:            *port,
		User:            *user,
		KeyPath:         *keyPath,
		RemoteDir:       *remoteDir,
		SSHPath:         *sshPath,
		SCPPath:         *scpPath,
		StartQEMU:       *startQEMU,
		QEMUBinary:      *qemuBinary,
		QEMUKernel:      *qemuKernel,
		QEMUMemory:      *qemuMemory,
		QEMUCPUCount:    *qemuCPUs,
		QEMUCPUModel:    *qemuCPU,
		QEMUBootTimeout: qemuBootTimeout,
		QEMUSerialLog:   *qemuSerialLog,
	})
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(execution)
}

type knownAppVerifiedCatalogDispatchExecutionRequest struct {
	Version         string
	Plan            appidentity.KnownAppVerifiedCatalogDispatchRequestPlan
	LauncherPath    string
	OwnerTimeout    time.Duration
	Host            string
	Port            string
	User            string
	KeyPath         string
	RemoteDir       string
	SSHPath         string
	SCPPath         string
	StartQEMU       bool
	QEMUBinary      string
	QEMUKernel      string
	QEMUMemory      string
	QEMUCPUCount    string
	QEMUCPUModel    string
	QEMUBootTimeout time.Duration
	QEMUSerialLog   string
}

type knownAppVerifiedCatalogDispatchExecutionResult struct {
	Version                             string `json:"version"`
	SchemaVersion                       string `json:"schema_version"`
	RequestType                         string `json:"request_type"`
	Source                              string `json:"source"`
	RuntimeMethod                       string `json:"runtime_method"`
	ExecutionMethod                     string `json:"execution_method"`
	AppID                               string `json:"app_id"`
	DisplayName                         string `json:"display_name"`
	AppVersion                          string `json:"app_version"`
	DispatchRequestConsumed             bool   `json:"dispatch_request_consumed"`
	DispatchRequestRelativePath         string `json:"dispatch_request_relative_path"`
	DispatchRequestDigestVerified       bool   `json:"dispatch_request_digest_verified"`
	DispatchRunnerRequestType           string `json:"dispatch_runner_request_type"`
	DispatchRunnerName                  string `json:"dispatch_runner_name"`
	DispatchRunnerInvoked               bool   `json:"dispatch_runner_invoked"`
	DispatchRunnerExitCode              int    `json:"dispatch_runner_exit_code"`
	DispatchRunnerOutputJSONObserved    bool   `json:"dispatch_runner_output_json_observed"`
	DispatchRunnerArgumentValuesExposed bool   `json:"dispatch_runner_argument_values_exposed"`
	RuntimeOwnerTransportSupplied       bool   `json:"runtime_owner_transport_supplied"`
	GuestStartRequested                 bool   `json:"guest_start_requested"`
	GuestStartAttempted                 bool   `json:"guest_start_attempted"`
	GuestStarted                        bool   `json:"guest_started"`
	GuestStartMode                      string `json:"guest_start_mode"`
	QEMUExecuted                        bool   `json:"qemu_executed"`
	QEMUSerialLogWritten                bool   `json:"qemu_serial_log_written"`
	DelegatedRequestType                string `json:"delegated_request_type"`
	DelegatedEvidenceSource             string `json:"delegated_evidence_source"`
	DelegatedStatus                     string `json:"delegated_status"`
	DelegatedGuestBoundary              string `json:"delegated_guest_boundary"`
	DelegatedCacheStatus                string `json:"delegated_cache_status"`
	DelegatedRuntimeOwnedDispatch       bool   `json:"delegated_runtime_owned_dispatch"`
	DelegatedArtifactVerified           bool   `json:"delegated_artifact_verified"`
	DelegatedManagedArtifactCopied      bool   `json:"delegated_managed_artifact_copied"`
	DelegatedMarkerObserved             bool   `json:"delegated_marker_observed"`
	DelegatedSmokePassed                bool   `json:"delegated_smoke_passed"`
	DelegatedExecutionStarted           bool   `json:"delegated_execution_started"`
	DelegatedBackendProcessStarted      bool   `json:"delegated_backend_process_started"`
	DelegatedManagedGuestRunnerInvoked  bool   `json:"delegated_managed_guest_runner_invoked"`
	DelegatedManagedGuestReachable      bool   `json:"delegated_managed_guest_reachable"`
	DelegatedManagedGuestRuntimeReady   bool   `json:"delegated_managed_guest_runtime_ready"`
	DelegatedFileArgumentCount          int    `json:"delegated_file_argument_count"`
	DelegatedFileArgumentCopiedCount    int    `json:"delegated_file_argument_copied_count"`
	DelegatedFileArgumentsPassed        bool   `json:"delegated_file_arguments_passed"`
	DelegatedFileArgumentWinePathReady  bool   `json:"delegated_file_argument_winepath_translated"`
	DelegatedFileArgumentWinePathCount  int    `json:"delegated_file_argument_winepath_translated_count"`
	DelegatedWindowMatchObserved        bool   `json:"delegated_window_match_observed"`
	DelegatedWindowEvidenceObserved     bool   `json:"delegated_window_evidence_observed"`
	DelegatedGUIEvidenceReady           bool   `json:"delegated_gui_evidence_ready"`
	DelegatedSessionConsumed            bool   `json:"delegated_controlled_execution_session_consumed"`
	DelegatedSessionDigestVerified      bool   `json:"delegated_controlled_session_digest_verified"`
	DelegatedControlledSessionWindow    bool   `json:"delegated_controlled_session_window_observed"`
	HostRootModified                    bool   `json:"host_root_modified"`
	PrivilegedContainerRequired         bool   `json:"privileged_container_required"`
	HostNetworkingRequired              bool   `json:"host_networking_required"`
	DockerSocketMounted                 bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired              bool   `json:"broad_host_mount_required"`
	RawHostPathExposed                  bool   `json:"raw_host_path_exposed"`
	RawExecutablePathExposed            bool   `json:"raw_executable_path_exposed"`
	RawFileArgumentPathExposed          bool   `json:"raw_file_argument_path_exposed"`
	StateRootPathExposed                bool   `json:"state_root_path_exposed"`
	CacheRootPathExposed                bool   `json:"cache_root_path_exposed"`
	RunnerPathExposed                   bool   `json:"runner_path_exposed"`
	WindowEvidenceSummaryExposed        bool   `json:"window_evidence_summary_exposed"`
	RawCommandExposed                   bool   `json:"raw_command_exposed"`
	RawLauncherOutputExposed            bool   `json:"raw_launcher_output_exposed"`
	BackendDetailsExposed               bool   `json:"backend_details_exposed"`
	DesktopSafeSummary                  string `json:"desktop_safe_summary"`
}

func executeKnownAppVerifiedCatalogDispatchRequest(request knownAppVerifiedCatalogDispatchExecutionRequest) (knownAppVerifiedCatalogDispatchExecutionResult, error) {
	plan := request.Plan
	if len(plan.RunnerArgv) == 0 {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch execution requires runner argv")
	}
	launcherPath := strings.TrimSpace(request.LauncherPath)
	if launcherPath == "" {
		launcherPath = plan.RunnerArgv[0]
	}
	if strings.ContainsAny(launcherPath, "\r\n") {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch execution requires a single-line launcher path")
	}
	ownerTimeout := request.OwnerTimeout
	if ownerTimeout <= 0 {
		ownerTimeout = 5 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), ownerTimeout)
	defer cancel()

	var qemuGuest *winapp.QEMUStartedGuest
	guestStarted := false
	qemuExecuted := false
	host := strings.TrimSpace(request.Host)
	port := strings.TrimSpace(request.Port)
	if request.StartQEMU {
		guest, err := winapp.StartQEMUStartedGuest(ctx, winapp.QEMUStartedGuestRequest{
			Binary:        request.QEMUBinary,
			KernelImage:   request.QEMUKernel,
			Memory:        request.QEMUMemory,
			CPUCount:      request.QEMUCPUCount,
			CPUModel:      request.QEMUCPUModel,
			Host:          host,
			Port:          port,
			User:          request.User,
			KeyPath:       request.KeyPath,
			SSHPath:       request.SSHPath,
			BootTimeout:   request.QEMUBootTimeout,
			SerialLogPath: request.QEMUSerialLog,
		})
		if err != nil {
			return knownAppVerifiedCatalogDispatchExecutionResult{}, err
		}
		qemuGuest = guest
		defer qemuGuest.Stop()
		guestStarted = true
		qemuExecuted = true
		host = guest.Host
		port = guest.Port
	}
	transportArgs, err := knownAppVerifiedCatalogDispatchExecutionTransportArgs(knownAppVerifiedCatalogDispatchExecutionTransport{
		Host:      host,
		Port:      port,
		User:      request.User,
		KeyPath:   request.KeyPath,
		RemoteDir: request.RemoteDir,
		SSHPath:   request.SSHPath,
		SCPPath:   request.SCPPath,
	})
	if err != nil {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, err
	}
	argv := append([]string{}, plan.RunnerArgv[1:]...)
	argv = append(argv, transportArgs...)
	output, err := exec.CommandContext(ctx, launcherPath, argv...).Output()
	if ctx.Err() != nil {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch runner invocation timed out")
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return knownAppVerifiedCatalogDispatchExecutionResult{}, fmt.Errorf("known app verified catalog dispatch runner failed with exit code %d", exitErr.ExitCode())
		}
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch runner invocation failed")
	}
	var delegated map[string]any
	if err := json.Unmarshal(output, &delegated); err != nil {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch runner must emit JSON")
	}
	result := knownAppVerifiedCatalogDispatchExecutionResult{
		Version:                             request.Version,
		SchemaVersion:                       appidentity.KnownAppVerifiedCatalogDispatchExecutionSchemaVersion,
		RequestType:                         appidentity.KnownAppVerifiedCatalogDispatchExecutionRequestType,
		Source:                              appidentity.KnownAppVerifiedCatalogDispatchRequestType + "+runner-consumption",
		RuntimeMethod:                       "ReadKnownAppVerifiedCatalogDispatchRequest",
		ExecutionMethod:                     "RunKnownAppVerifiedCatalogDispatchRunnerExecution",
		AppID:                               plan.AppID,
		DisplayName:                         plan.DisplayName,
		AppVersion:                          plan.AppVersion,
		DispatchRequestConsumed:             true,
		DispatchRequestRelativePath:         plan.DispatchRequestRelativePath,
		DispatchRequestDigestVerified:       plan.DigestMatched,
		DispatchRunnerRequestType:           plan.RunnerRequestType,
		DispatchRunnerName:                  filepath.Base(launcherPath),
		DispatchRunnerInvoked:               true,
		DispatchRunnerExitCode:              0,
		DispatchRunnerOutputJSONObserved:    true,
		DispatchRunnerArgumentValuesExposed: false,
		RuntimeOwnerTransportSupplied:       len(transportArgs) > 0,
		GuestStartRequested:                 request.StartQEMU,
		GuestStartAttempted:                 request.StartQEMU,
		GuestStarted:                        guestStarted,
		GuestStartMode:                      map[bool]string{true: "go-qemu", false: "external"}[request.StartQEMU],
		QEMUExecuted:                        qemuExecuted,
		QEMUSerialLogWritten:                strings.TrimSpace(request.QEMUSerialLog) != "",
		DelegatedRequestType:                stringJSONField(delegated, "request_type"),
		DelegatedEvidenceSource:             stringJSONField(delegated, "evidence_source"),
		DelegatedStatus:                     stringJSONField(delegated, "status"),
		DelegatedGuestBoundary:              stringJSONField(delegated, "guest_boundary"),
		DelegatedCacheStatus:                stringJSONField(delegated, "cache_status"),
		DelegatedRuntimeOwnedDispatch:       boolJSONField(delegated, "runtime_owned_dispatch"),
		DelegatedArtifactVerified:           boolJSONField(delegated, "artifact_verified"),
		DelegatedManagedArtifactCopied:      boolJSONField(delegated, "managed_artifact_copied"),
		DelegatedMarkerObserved:             boolJSONField(delegated, "marker_observed"),
		DelegatedSmokePassed:                boolJSONField(delegated, "smoke_passed"),
		DelegatedExecutionStarted:           boolJSONField(delegated, "execution_started"),
		DelegatedBackendProcessStarted:      boolJSONField(delegated, "backend_process_started"),
		DelegatedManagedGuestRunnerInvoked:  boolJSONField(delegated, "managed_guest_runner_invoked"),
		DelegatedManagedGuestReachable:      boolJSONField(delegated, "managed_guest_reachable"),
		DelegatedManagedGuestRuntimeReady:   boolJSONField(delegated, "managed_guest_runtime_ready"),
		DelegatedFileArgumentCount:          intJSONField(delegated, "file_argument_count"),
		DelegatedFileArgumentCopiedCount:    intJSONField(delegated, "file_argument_copied_count"),
		DelegatedFileArgumentsPassed:        boolJSONField(delegated, "file_arguments_passed"),
		DelegatedFileArgumentWinePathReady:  boolJSONField(delegated, "file_argument_winepath_translated"),
		DelegatedFileArgumentWinePathCount:  intJSONField(delegated, "file_argument_winepath_translated_count"),
		DelegatedWindowMatchObserved:        boolJSONField(delegated, "window_match_observed"),
		DelegatedSessionConsumed:            boolJSONField(delegated, "controlled_execution_session_consumed"),
		DelegatedSessionDigestVerified:      boolJSONField(delegated, "controlled_session_digest_verified"),
		DelegatedControlledSessionWindow:    boolJSONField(delegated, "controlled_session_window_observed"),
		HostRootModified:                    boolJSONField(delegated, "host_root_modified"),
		PrivilegedContainerRequired:         boolJSONField(delegated, "privileged_container_required"),
		HostNetworkingRequired:              boolJSONField(delegated, "host_networking_required"),
		DockerSocketMounted:                 boolJSONField(delegated, "docker_socket_mounted"),
		BroadHostMountRequired:              boolJSONField(delegated, "broad_host_mount_required"),
		RawHostPathExposed:                  boolJSONField(delegated, "raw_host_path_exposed"),
		RawExecutablePathExposed:            boolJSONField(delegated, "raw_executable_path_exposed"),
		RawFileArgumentPathExposed:          boolJSONField(delegated, "raw_file_argument_path_exposed"),
		StateRootPathExposed:                false,
		CacheRootPathExposed:                false,
		RunnerPathExposed:                   false,
		WindowEvidenceSummaryExposed:        false,
		RawCommandExposed:                   boolJSONField(delegated, "raw_command_exposed"),
		RawLauncherOutputExposed:            false,
		BackendDetailsExposed:               boolJSONField(delegated, "backend_details_exposed"),
		DesktopSafeSummary:                  plan.DisplayName + " dispatch request was consumed by the Runtime owner runner bridge.",
	}
	result.DelegatedWindowEvidenceObserved = result.DelegatedWindowMatchObserved ||
		result.DelegatedControlledSessionWindow ||
		boolJSONField(delegated, "x_window_observed") ||
		boolJSONField(delegated, "window_observed")
	result.DelegatedGUIEvidenceReady = result.DelegatedSmokePassed &&
		result.DelegatedExecutionStarted &&
		result.DelegatedWindowEvidenceObserved
	if result.DelegatedRequestType != winapp.KnownDispatchSmokeRequestType {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch runner returned an unexpected request type")
	}
	if result.HostRootModified || result.PrivilegedContainerRequired || result.HostNetworkingRequired ||
		result.DockerSocketMounted || result.BroadHostMountRequired || result.RawHostPathExposed ||
		result.RawExecutablePathExposed || result.RawFileArgumentPathExposed || result.RawCommandExposed ||
		result.BackendDetailsExposed {
		return knownAppVerifiedCatalogDispatchExecutionResult{}, errors.New("known app verified catalog dispatch runner reported unsafe host or backend exposure")
	}
	return result, nil
}

type knownAppVerifiedCatalogDispatchExecutionTransport struct {
	Host      string
	Port      string
	User      string
	KeyPath   string
	RemoteDir string
	SSHPath   string
	SCPPath   string
}

func knownAppVerifiedCatalogDispatchExecutionTransportArgs(transport knownAppVerifiedCatalogDispatchExecutionTransport) ([]string, error) {
	var args []string
	for _, pair := range []struct {
		flag  string
		value string
	}{
		{"--host", transport.Host},
		{"--port", transport.Port},
		{"--user", transport.User},
		{"--key", transport.KeyPath},
		{"--remote-dir", transport.RemoteDir},
		{"--ssh", transport.SSHPath},
		{"--scp", transport.SCPPath},
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
	return args, nil
}
