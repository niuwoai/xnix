package appidentity

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	KnownAppVerifiedCatalogRunPlanExecutionSchemaVersion = "xnix.runtime.known_app_verified_catalog_run_plan_execution.v1"
	KnownAppVerifiedCatalogRunPlanExecutionRequestType   = "known-app-verified-catalog-run-plan-execution"
	KnownAppVerifiedCatalogAppExecutionSchemaVersion     = "xnix.runtime.known_app_verified_catalog_app_execution.v1"
	KnownAppVerifiedCatalogAppExecutionRequestType       = "known-app-verified-catalog-app-execution"
)

type KnownAppVerifiedCatalogRunPlanExecutionRequest struct {
	RunPlanPath     string
	SmokeReportPath string
	Execute         bool
	WorkDir         string
	Timeout         time.Duration
	RecordedAtUTC   time.Time
}

type KnownAppVerifiedCatalogAppExecutionRequest struct {
	VerifiedCatalogPath string
	AppID               string
	SmokeReportPath     string
	Execute             bool
	WorkDir             string
	Timeout             time.Duration
	RecordedAtUTC       time.Time
}

type KnownAppVerifiedCatalogRunPlanExecutionResult struct {
	Version                              string                       `json:"version"`
	SchemaVersion                        string                       `json:"schema_version"`
	RequestType                          string                       `json:"request_type"`
	Source                               string                       `json:"source"`
	RuntimeMethod                        string                       `json:"runtime_method"`
	ReadMethod                           string                       `json:"read_method"`
	ExecutionMethod                      string                       `json:"execution_method"`
	AppID                                string                       `json:"app_id"`
	DisplayName                          string                       `json:"display_name"`
	AppVersion                           string                       `json:"app_version"`
	VerifiedCatalogConsumed              bool                         `json:"verified_catalog_consumed"`
	RequestedAppID                       string                       `json:"requested_app_id"`
	RunPlanGenerated                     bool                         `json:"run_plan_generated"`
	DirectRunPlanInput                   bool                         `json:"direct_run_plan_input"`
	RunPlanConsumed                      bool                         `json:"run_plan_consumed"`
	RunPlanRequestType                   string                       `json:"run_plan_request_type"`
	RuntimeOwnedActionReady              bool                         `json:"runtime_owned_action_ready"`
	DesktopCallableActionID              string                       `json:"desktop_callable_action_id"`
	DesktopCallableRoute                 string                       `json:"desktop_callable_route"`
	DesktopCallableRuntimeMethod         string                       `json:"desktop_callable_runtime_method"`
	DesktopCallableExecutionType         string                       `json:"desktop_callable_execution_type"`
	DesktopForwardedArgumentCount        int                          `json:"desktop_forwarded_argument_count"`
	DesktopForwardsOnlyAppID             bool                         `json:"desktop_forwards_only_app_id"`
	DesktopReceiptFieldsReconstructed    bool                         `json:"desktop_receipt_fields_reconstructed"`
	DesktopKDEStateRootAccess            bool                         `json:"desktop_kde_state_root_access"`
	DesktopOwnerInputsExposed            bool                         `json:"desktop_owner_inputs_exposed"`
	GUIEvidenceRequired                  bool                         `json:"gui_evidence_required"`
	GUIEvidenceConsumed                  bool                         `json:"gui_evidence_consumed"`
	WindowObservationRequired            bool                         `json:"window_observation_required"`
	WindowObserved                       bool                         `json:"window_observed"`
	WindowMatchObserved                  bool                         `json:"window_match_observed"`
	OwnerFileOpenRequired                bool                         `json:"owner_file_open_required"`
	OwnerFileOpenVerified                bool                         `json:"owner_file_open_verified"`
	OwnerFileOpenEntrypointInvoked       bool                         `json:"owner_file_open_entrypoint_invoked"`
	DocumentContentMarkerObserved        bool                         `json:"document_content_marker_observed"`
	SmokeCommandPlanned                  bool                         `json:"smoke_command_planned"`
	SmokeCommandName                     string                       `json:"smoke_command_name"`
	SmokeScript                          string                       `json:"smoke_script"`
	SmokeRequestType                     string                       `json:"smoke_request_type"`
	ExecutionRequested                   bool                         `json:"execution_requested"`
	ExecutionStarted                     bool                         `json:"execution_started"`
	ExecutionCompleted                   bool                         `json:"execution_completed"`
	ExecutionExitCode                    int                          `json:"execution_exit_code"`
	SmokeReportConsumed                  bool                         `json:"smoke_report_consumed"`
	SmokeStatus                          string                       `json:"smoke_status"`
	SmokePassed                          bool                         `json:"smoke_passed"`
	ActualWindowsAppRunObserved          bool                         `json:"actual_windows_app_run_observed"`
	CompatibilityCenterProjectionReady   bool                         `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady             bool                         `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidence                KnownAppSmokeEvidenceSummary `json:"known_app_smoke_evidence"`
	GoOwnedQ4WinAppAcceptanceReady       bool                         `json:"go_owned_q4_winapp_acceptance_ready"`
	GoOwnedQ4WinAppAcceptanceConsumed    bool                         `json:"go_owned_q4_winapp_acceptance_consumed"`
	GoOwnedQ4WinAppAcceptancePathExposed bool                         `json:"go_owned_q4_winapp_acceptance_path_exposed"`
	Q4CompileRequired                    bool                         `json:"q4_compile_required"`
	HostCompilationRequired              bool                         `json:"host_compilation_required"`
	HostCompilationAvoided               bool                         `json:"host_compilation_avoided"`
	TargetedRemoteVerificationReady      bool                         `json:"targeted_remote_verification_ready"`
	RuntimeOwned                         bool                         `json:"runtime_owned"`
	GoRuntimeBacked                      bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool                         `json:"kde_policy_owner"`
	ReviewOnly                           bool                         `json:"review_only"`
	DirectLaunchEnabled                  bool                         `json:"direct_launch_enabled"`
	LaunchEnabled                        bool                         `json:"launch_enabled"`
	DesktopFilesWritten                  bool                         `json:"desktop_files_written"`
	HostRootModified                     bool                         `json:"host_root_modified"`
	BackendLaunchEnabled                 bool                         `json:"backend_launch_enabled"`
	BackendDetailsExposed                bool                         `json:"backend_details_exposed"`
	RawOutputExposed                     bool                         `json:"raw_output_exposed"`
	RemotePathExposed                    bool                         `json:"remote_path_exposed"`
	SmokeCommandArgumentsExposed         bool                         `json:"smoke_command_arguments_exposed"`
	PrivilegedContainerRequired          bool                         `json:"privileged_container_required"`
	HostNetworkingRequired               bool                         `json:"host_networking_required"`
	DockerSocketMounted                  bool                         `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool                         `json:"broad_host_mount_required"`
	BlockedActions                       []string                     `json:"blocked_actions"`
	RecordedAtUTC                        string                       `json:"recorded_at_utc"`
	DesktopSafeSummary                   string                       `json:"desktop_safe_summary"`
}

func RunKnownAppVerifiedCatalogRunPlanExecution(request KnownAppVerifiedCatalogRunPlanExecutionRequest) (KnownAppVerifiedCatalogRunPlanExecutionResult, error) {
	if request.RunPlanPath == "" {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan execution requires --run-plan")
	}
	content, err := os.ReadFile(request.RunPlanPath)
	if err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, fmt.Errorf("read known app verified catalog run plan: %w", err)
	}
	var runPlan KnownAppVerifiedCatalogRunPlanPreview
	if err := json.Unmarshal(content, &runPlan); err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, fmt.Errorf("parse known app verified catalog run plan: %w", err)
	}
	return runKnownAppVerifiedCatalogRunPlanExecution(runPlan, request)
}

func RunKnownAppVerifiedCatalogAppExecution(request KnownAppVerifiedCatalogAppExecutionRequest) (KnownAppVerifiedCatalogRunPlanExecutionResult, error) {
	if request.VerifiedCatalogPath == "" {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog app execution requires --verified-catalog")
	}
	if request.AppID == "" {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog app execution requires --app")
	}
	runPlan, err := PreviewKnownAppVerifiedCatalogRunPlan(KnownAppVerifiedCatalogRunPlanRequest{
		VerifiedCatalogPath: request.VerifiedCatalogPath,
		AppID:               request.AppID,
	})
	if err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
	}
	result, err := runKnownAppVerifiedCatalogRunPlanExecution(runPlan, KnownAppVerifiedCatalogRunPlanExecutionRequest{
		SmokeReportPath: request.SmokeReportPath,
		Execute:         request.Execute,
		WorkDir:         request.WorkDir,
		Timeout:         request.Timeout,
		RecordedAtUTC:   request.RecordedAtUTC,
	})
	if err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
	}
	result.SchemaVersion = KnownAppVerifiedCatalogAppExecutionSchemaVersion
	result.RequestType = KnownAppVerifiedCatalogAppExecutionRequestType
	result.Source = "known-app-verified-catalog+runtime-owned-app-execution"
	result.RuntimeMethod = "RunKnownVerifiedApplicationFromCatalog"
	result.ReadMethod = "ReadKnownVerifiedApplicationCatalog"
	result.VerifiedCatalogConsumed = true
	result.RequestedAppID = request.AppID
	result.RunPlanGenerated = true
	result.DirectRunPlanInput = false
	result.DesktopSafeSummary = result.DisplayName + " can be executed by the Runtime owner from the verified catalog and app id while KDE forwards only the app id."
	if result.ActualWindowsAppRunObserved {
		evidence, err := knownAppVerifiedCatalogAppExecutionSmokeEvidence(result)
		if err != nil {
			return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
		}
		result.CompatibilityCenterProjectionReady = true
		result.KDECenterProjectionReady = true
		result.KnownAppSmokeEvidence = evidence
	}
	if err := validateNoBackendTerms(result, "known app verified catalog app execution"); err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
	}
	return result, nil
}

func KnownAppSmokeEvidenceFromKnownAppVerifiedCatalogAppExecution(payload []byte) (KnownAppSmokeEvidenceSummary, error) {
	var result KnownAppVerifiedCatalogRunPlanExecutionResult
	if err := json.Unmarshal(payload, &result); err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("parse known app verified catalog app execution: %w", err)
	}
	if result.SchemaVersion != KnownAppVerifiedCatalogAppExecutionSchemaVersion ||
		result.RequestType != KnownAppVerifiedCatalogAppExecutionRequestType {
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution has invalid schema or request type")
	}
	if !result.CompatibilityCenterProjectionReady || !result.KDECenterProjectionReady {
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution is not ready for center consumption")
	}
	if result.KnownAppSmokeEvidence.AppID == "" {
		return knownAppVerifiedCatalogAppExecutionSmokeEvidence(result)
	}
	return normalizeKnownAppSmokeEvidenceItem(result.KnownAppSmokeEvidence)
}

func knownAppVerifiedCatalogAppExecutionSmokeEvidence(result KnownAppVerifiedCatalogRunPlanExecutionResult) (KnownAppSmokeEvidenceSummary, error) {
	switch {
	case result.SchemaVersion != KnownAppVerifiedCatalogAppExecutionSchemaVersion || result.RequestType != KnownAppVerifiedCatalogAppExecutionRequestType:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires app-execution schema")
	case result.AppID == "" || result.DisplayName == "" || result.AppVersion == "":
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires app identity")
	case !result.VerifiedCatalogConsumed || !result.RunPlanGenerated || result.DirectRunPlanInput || !result.RunPlanConsumed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires Runtime-generated run-plan consumption")
	case result.SmokeRequestType != "q4-messagebox-smoke" || result.SmokeStatus != "passed" || !result.SmokePassed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires passed q4 GUI smoke")
	case !result.GUIEvidenceConsumed || !result.WindowObserved || !result.WindowMatchObserved || !result.ActualWindowsAppRunObserved:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires observed GUI execution")
	case !result.OwnerFileOpenVerified || !result.OwnerFileOpenEntrypointInvoked || !result.DocumentContentMarkerObserved:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires owner file-open and document marker evidence")
	case !result.GoOwnedQ4WinAppAcceptanceReady || !result.GoOwnedQ4WinAppAcceptanceConsumed || !result.HostCompilationAvoided:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence requires Go-owned q4 acceptance and host compilation avoidance")
	case result.GoOwnedQ4WinAppAcceptancePathExposed || result.RawOutputExposed || result.RemotePathExposed || result.SmokeCommandArgumentsExposed:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence exposes unsafe runtime details")
	case result.HostRootModified || result.PrivilegedContainerRequired || result.HostNetworkingRequired || result.DockerSocketMounted || result.BroadHostMountRequired:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence opens unsafe host or container gates")
	case result.BackendLaunchEnabled || result.LaunchEnabled || result.DirectLaunchEnabled || result.DesktopFilesWritten:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence must remain review-only")
	case !result.RuntimeOwned || !result.GoRuntimeBacked || result.KDEPolicyOwner:
		return KnownAppSmokeEvidenceSummary{}, errors.New("known app verified catalog app execution evidence must remain Runtime-owned and Go-backed")
	}
	evidence, err := normalizeKnownAppSmokeEvidenceItem(KnownAppSmokeEvidenceSummary{
		AppID:                                result.AppID,
		DisplayName:                          result.DisplayName,
		AppVersion:                           result.AppVersion,
		EvidenceKind:                         "known-application-gui-smoke",
		EvidenceSource:                       GUISmokeEvidenceSourceWineGuest,
		SmokeStatus:                          "passed",
		XWindowObserved:                      true,
		WindowObserved:                       true,
		CompatibilityState:                   "owner-controlled-gui-qemu-wine-verified",
		CenterCardState:                      "validated-owner-controlled-gui-runtime-run",
		LaunchAuthorizationState:             "review-required",
		MarkerObserved:                       result.DocumentContentMarkerObserved,
		ChecksumVerified:                     result.GoOwnedQ4WinAppAcceptanceReady,
		ExecutionEvidenceRecorded:            true,
		OwnerControlledRuntimeLaunchVerified: true,
		RuntimeDispatchVerified:              true,
		LaunchAuthorizationRequired:          true,
		DesktopLaunchEnabled:                 false,
		RuntimeOwned:                         true,
		KDEPolicyOwner:                       false,
		ActionExecutionEnabled:               false,
		BackendLaunchEnabled:                 false,
		HostRootModified:                     false,
		BackendDetailsExposed:                false,
		RawArtifactPathExposed:               false,
		Summary:                              result.DisplayName + " completed a Runtime-owned verified-catalog q4 GUI execution with owner file-open and document marker evidence.",
	})
	if err != nil {
		return KnownAppSmokeEvidenceSummary{}, fmt.Errorf("consume known app verified catalog app execution: %w", err)
	}
	return evidence, nil
}

func runKnownAppVerifiedCatalogRunPlanExecution(runPlan KnownAppVerifiedCatalogRunPlanPreview, request KnownAppVerifiedCatalogRunPlanExecutionRequest) (KnownAppVerifiedCatalogRunPlanExecutionResult, error) {
	if err := validateKnownAppVerifiedCatalogRunPlanExecutionInput(runPlan); err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
	}

	recordedAt := request.RecordedAtUTC
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	result := knownAppVerifiedCatalogRunPlanExecutionBaseResult(runPlan, request.Execute, recordedAt)

	var smokeReport map[string]any
	if request.SmokeReportPath != "" {
		reportContent, err := os.ReadFile(request.SmokeReportPath)
		if err != nil {
			return KnownAppVerifiedCatalogRunPlanExecutionResult{}, fmt.Errorf("read known app verified catalog smoke report: %w", err)
		}
		if err := json.Unmarshal(reportContent, &smokeReport); err != nil {
			return KnownAppVerifiedCatalogRunPlanExecutionResult{}, fmt.Errorf("parse known app verified catalog smoke report: %w", err)
		}
		result.SmokeReportConsumed = true
	}
	if request.Execute {
		executedReport, err := runKnownAppVerifiedCatalogSmokeCommand(request, runPlan)
		if err != nil {
			return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
		}
		smokeReport = executedReport
		result.ExecutionStarted = true
		result.ExecutionCompleted = true
		result.ExecutionExitCode = 0
		result.SmokeReportConsumed = true
	}
	if smokeReport != nil {
		updated, err := consumeKnownAppVerifiedCatalogRunPlanSmokeReport(result, runPlan, smokeReport)
		if err != nil {
			return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
		}
		result = updated
	}
	if err := validateNoBackendTerms(result, "known app verified catalog run plan execution"); err != nil {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, err
	}
	return result, nil
}

func validateKnownAppVerifiedCatalogRunPlanExecutionInput(runPlan KnownAppVerifiedCatalogRunPlanPreview) error {
	if runPlan.SchemaVersion != KnownAppVerifiedCatalogRunPlanSchemaVersion ||
		runPlan.RequestType != KnownAppVerifiedCatalogRunPlanRequestType ||
		!runPlan.VerifiedCatalogConsumed ||
		!runPlan.RuntimeOwnedActionReady ||
		runPlan.DesktopCallableRoute != "runtime-owner://known-app-verified-catalog/run-plan" ||
		runPlan.DesktopCallableRuntimeMethod != "PlanKnownVerifiedApplicationRun" ||
		len(runPlan.DesktopForwardedArguments) != 1 ||
		runPlan.DesktopForwardedArguments[0] != runPlan.AppID ||
		!runPlan.DesktopForwardsOnlyAppID ||
		runPlan.DesktopReceiptFieldsReconstructed ||
		runPlan.DesktopKDEStateRootAccess ||
		runPlan.DesktopOwnerInputsExposed {
		return errors.New("known app verified catalog run plan execution requires a Runtime-owned app-id-only handoff")
	}
	if runPlan.AppID == "" || strings.ContainsAny(runPlan.AppID, "\r\n") {
		return errors.New("known app verified catalog run plan execution requires a safe app id")
	}
	if !runPlan.Q4ExecutionRequired ||
		!runPlan.Q4ExecutionPlanned ||
		runPlan.Q4ExecutionStarted ||
		!runPlan.ReviewOnly ||
		!runPlan.OperatorReviewRequired ||
		!runPlan.RuntimeOwned ||
		!runPlan.GoRuntimeBacked ||
		runPlan.KDEPolicyOwner ||
		runPlan.DirectLaunchEnabled ||
		runPlan.LaunchEnabled ||
		runPlan.ExecutionStarted ||
		runPlan.BackendLaunchEnabled ||
		runPlan.DesktopFilesWritten ||
		runPlan.HostRootModified ||
		runPlan.BackendDetailsExposed ||
		runPlan.RawOutputExposed ||
		runPlan.RemotePathExposed ||
		runPlan.PrivilegedContainerRequired ||
		runPlan.HostNetworkingRequired ||
		runPlan.DockerSocketMounted ||
		runPlan.BroadHostMountRequired ||
		runPlan.HostCompilationRequired ||
		!runPlan.HostCompilationAvoided ||
		!runPlan.TargetedRemoteVerificationReady {
		return errors.New("known app verified catalog run plan execution requires a closed review-only q4 plan")
	}
	if runPlan.GUIEvidenceRequired || runPlan.GUIEvidenceConsumed || runPlan.WindowObservationRequired || runPlan.OwnerFileOpenRequired {
		if !runPlan.GUIEvidenceRequired ||
			!runPlan.GUIEvidenceConsumed ||
			!runPlan.WindowObservationRequired ||
			!runPlan.OwnerFileOpenRequired ||
			!runPlan.OwnerFileOpenVerified ||
			runPlan.DesktopCallableExecutionType != "review-only-q4-gui-smoke" {
			return errors.New("known app verified catalog GUI run plan execution requires consumed GUI evidence")
		}
	}
	return validateKnownAppVerifiedCatalogSmokeCommand(runPlan)
}

func validateKnownAppVerifiedCatalogSmokeCommand(runPlan KnownAppVerifiedCatalogRunPlanPreview) error {
	command := runPlan.RemoteSmokeCommand
	if len(command) < 3 || command[0] != "ruby" {
		return errors.New("known app verified catalog run plan execution requires a Ruby smoke harness command")
	}
	for _, arg := range command {
		if arg == "" || strings.ContainsAny(arg, "\r\n") {
			return errors.New("known app verified catalog run plan execution requires single-line smoke command arguments")
		}
	}
	switch runPlan.RemoteSmokeRequestType {
	case "q4-messagebox-smoke":
		if strings.Join(command, "\x00") != strings.Join([]string{"ruby", "scripts/q4_messagebox_smoke.rb", "--execute", "--owner-file-open"}, "\x00") {
			return errors.New("known app verified catalog MessageBox run plan execution requires the owner-file-open q4 smoke")
		}
	case "q4-winapp-smoke":
		if len(command) != 5 ||
			command[1] != "scripts/q4_winapp_smoke.rb" ||
			command[2] != "--execute" ||
			command[3] != "--known-app-id" ||
			command[4] != runPlan.AppID {
			return errors.New("known app verified catalog GUI run plan execution requires the generic q4 Windows app smoke")
		}
	case "remote-known-winapp-guest-wine-smoke":
		if len(command) != 5 ||
			command[1] != "scripts/remote_known_winapp_guest_wine_smoke.rb" ||
			command[2] != "--execute" ||
			command[3] != "--app" ||
			command[4] != runPlan.AppID {
			return errors.New("known app verified catalog run plan execution requires the remote known Windows app smoke")
		}
	default:
		return fmt.Errorf("known app verified catalog run plan execution does not support smoke request type %q", runPlan.RemoteSmokeRequestType)
	}
	return nil
}

func knownAppVerifiedCatalogRunPlanExecutionBaseResult(runPlan KnownAppVerifiedCatalogRunPlanPreview, execute bool, recordedAt time.Time) KnownAppVerifiedCatalogRunPlanExecutionResult {
	script := ""
	if len(runPlan.RemoteSmokeCommand) > 1 {
		script = runPlan.RemoteSmokeCommand[1]
	}
	return KnownAppVerifiedCatalogRunPlanExecutionResult{
		Version:                           knownAppVerifiedCatalogRunPlanExecutionVersion(runPlan),
		SchemaVersion:                     KnownAppVerifiedCatalogRunPlanExecutionSchemaVersion,
		RequestType:                       KnownAppVerifiedCatalogRunPlanExecutionRequestType,
		Source:                            "known-app-verified-catalog-run-plan+runtime-owned-smoke-execution",
		RuntimeMethod:                     "RunKnownVerifiedApplicationRunPlanExecution",
		ReadMethod:                        "ReadKnownVerifiedApplicationRunPlan",
		ExecutionMethod:                   "InvokeKnownVerifiedApplicationSmokeHarness",
		AppID:                             runPlan.AppID,
		DisplayName:                       runPlan.DisplayName,
		AppVersion:                        runPlan.AppVersion,
		VerifiedCatalogConsumed:           false,
		RequestedAppID:                    runPlan.AppID,
		RunPlanGenerated:                  false,
		DirectRunPlanInput:                true,
		RunPlanConsumed:                   true,
		RunPlanRequestType:                runPlan.RequestType,
		RuntimeOwnedActionReady:           runPlan.RuntimeOwnedActionReady,
		DesktopCallableActionID:           runPlan.DesktopCallableActionID,
		DesktopCallableRoute:              runPlan.DesktopCallableRoute,
		DesktopCallableRuntimeMethod:      runPlan.DesktopCallableRuntimeMethod,
		DesktopCallableExecutionType:      runPlan.DesktopCallableExecutionType,
		DesktopForwardedArgumentCount:     len(runPlan.DesktopForwardedArguments),
		DesktopForwardsOnlyAppID:          runPlan.DesktopForwardsOnlyAppID,
		DesktopReceiptFieldsReconstructed: runPlan.DesktopReceiptFieldsReconstructed,
		DesktopKDEStateRootAccess:         runPlan.DesktopKDEStateRootAccess,
		DesktopOwnerInputsExposed:         runPlan.DesktopOwnerInputsExposed,
		GUIEvidenceRequired:               runPlan.GUIEvidenceRequired,
		GUIEvidenceConsumed:               runPlan.GUIEvidenceConsumed,
		WindowObservationRequired:         runPlan.WindowObservationRequired,
		OwnerFileOpenRequired:             runPlan.OwnerFileOpenRequired,
		OwnerFileOpenVerified:             runPlan.OwnerFileOpenVerified,
		SmokeCommandPlanned:               true,
		SmokeCommandName:                  filepath.Base(script),
		SmokeScript:                       script,
		SmokeRequestType:                  runPlan.RemoteSmokeRequestType,
		ExecutionRequested:                execute,
		Q4CompileRequired:                 runPlan.GUIEvidenceRequired,
		HostCompilationRequired:           false,
		HostCompilationAvoided:            true,
		TargetedRemoteVerificationReady:   true,
		RuntimeOwned:                      true,
		GoRuntimeBacked:                   true,
		KDEPolicyOwner:                    false,
		ReviewOnly:                        true,
		DirectLaunchEnabled:               false,
		LaunchEnabled:                     false,
		DesktopFilesWritten:               false,
		HostRootModified:                  false,
		BackendLaunchEnabled:              false,
		BackendDetailsExposed:             false,
		RawOutputExposed:                  false,
		RemotePathExposed:                 false,
		SmokeCommandArgumentsExposed:      false,
		PrivilegedContainerRequired:       false,
		HostNetworkingRequired:            false,
		DockerSocketMounted:               false,
		BroadHostMountRequired:            false,
		BlockedActions: []string{
			"launch verified app directly from KDE",
			"write desktop files during run-plan execution",
			"expose smoke command argument values to desktop",
			"expose raw q4 smoke output to desktop",
			"compile known app runner on the host",
		},
		RecordedAtUTC:      recordedAt.Format(time.RFC3339),
		DesktopSafeSummary: runPlan.DisplayName + " can be executed by the Runtime owner from a verified catalog run plan while KDE forwards only the app id.",
	}
}

func knownAppVerifiedCatalogRunPlanExecutionVersion(runPlan KnownAppVerifiedCatalogRunPlanPreview) string {
	version, err := readRuntimeServiceBindingVersion(".")
	if err == nil {
		return version
	}
	return runPlan.AppVersion
}

func runKnownAppVerifiedCatalogSmokeCommand(request KnownAppVerifiedCatalogRunPlanExecutionRequest, runPlan KnownAppVerifiedCatalogRunPlanPreview) (map[string]any, error) {
	timeout := request.Timeout
	if timeout <= 0 {
		timeout = 10 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	command := runPlan.RemoteSmokeCommand
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	if request.WorkDir != "" {
		cmd.Dir = request.WorkDir
	}
	output, err := cmd.Output()
	if ctx.Err() != nil {
		return nil, errors.New("known app verified catalog run plan smoke execution timed out")
	}
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("known app verified catalog run plan smoke execution failed with exit code %d", exitErr.ExitCode())
		}
		return nil, errors.New("known app verified catalog run plan smoke execution failed")
	}
	var report map[string]any
	if err := json.Unmarshal(output, &report); err != nil {
		return nil, errors.New("known app verified catalog run plan smoke execution must emit JSON")
	}
	return report, nil
}

func consumeKnownAppVerifiedCatalogRunPlanSmokeReport(result KnownAppVerifiedCatalogRunPlanExecutionResult, runPlan KnownAppVerifiedCatalogRunPlanPreview, report map[string]any) (KnownAppVerifiedCatalogRunPlanExecutionResult, error) {
	if remoteString(report, "request_type") != runPlan.RemoteSmokeRequestType {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan smoke report request type does not match the run plan")
	}
	appID := remoteString(report, "app_id")
	if appID == "" {
		appID = remoteString(report, "known_app_id")
	}
	if appID != "" && appID != runPlan.AppID {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan smoke report app id does not match the run plan")
	}
	result.SmokeStatus = remoteString(report, "status")
	result.SmokePassed = result.SmokeStatus == "passed"
	result.WindowObserved = remoteBool(report, "window_observed") || remoteBool(report, "x_window_observed")
	result.WindowMatchObserved = remoteBool(report, "window_match_observed")
	result.OwnerFileOpenEntrypointInvoked = remoteBool(report, "owner_file_open_entrypoint_invoked") ||
		remoteBool(report, "runtime_evidence_owner_file_open_entrypoint_invoked")
	result.DocumentContentMarkerObserved = remoteBool(report, "document_content_marker_observed") ||
		remoteBool(report, "real_run_acceptance_document_content_marker_observed") ||
		remoteBool(report, "go_owned_q4_winapp_acceptance_document_content_marker_observed")
	result.GoOwnedQ4WinAppAcceptanceReady = remoteBool(report, "go_owned_q4_winapp_acceptance_ready") ||
		remoteBool(report, "real_run_acceptance_ready")
	result.GoOwnedQ4WinAppAcceptanceConsumed = remoteBool(report, "go_owned_q4_winapp_acceptance_consumed") ||
		remoteBool(report, "real_run_acceptance_center_projection_consumed")
	result.GoOwnedQ4WinAppAcceptancePathExposed = remoteBool(report, "go_owned_q4_winapp_acceptance_path_exposed")
	result.Q4CompileRequired = remoteBool(report, "q4_compile_required")
	if !remoteBool(report, "host_compilation_avoided") {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan smoke report must avoid host compilation")
	}
	result.HostCompilationAvoided = true
	result.HostRootModified = remoteBool(report, "host_root_modified")
	result.PrivilegedContainerRequired = remoteBool(report, "privileged_container_required")
	result.HostNetworkingRequired = remoteBool(report, "host_networking_required")
	result.DockerSocketMounted = remoteBool(report, "docker_socket_mounted")
	result.BroadHostMountRequired = remoteBool(report, "broad_host_mount_required")
	if result.GoOwnedQ4WinAppAcceptancePathExposed ||
		result.HostRootModified ||
		result.PrivilegedContainerRequired ||
		result.HostNetworkingRequired ||
		result.DockerSocketMounted ||
		result.BroadHostMountRequired {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan smoke report opens unsafe host/container gates")
	}
	result.ActualWindowsAppRunObserved = result.SmokePassed
	if runPlan.GUIEvidenceRequired {
		result.ActualWindowsAppRunObserved = result.SmokePassed &&
			(result.WindowObserved || result.WindowMatchObserved) &&
			(!runPlan.OwnerFileOpenRequired || (result.OwnerFileOpenVerified && result.OwnerFileOpenEntrypointInvoked && result.DocumentContentMarkerObserved)) &&
			result.GoOwnedQ4WinAppAcceptanceReady
	} else if result.SmokePassed {
		result.ActualWindowsAppRunObserved = true
	}
	if !result.ActualWindowsAppRunObserved {
		return KnownAppVerifiedCatalogRunPlanExecutionResult{}, errors.New("known app verified catalog run plan smoke report did not prove actual Windows app execution")
	}
	return result, nil
}
