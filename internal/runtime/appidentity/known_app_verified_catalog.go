package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

const (
	KnownAppVerifiedCatalogPreviewSchemaVersion       = "xnix.runtime.known_app_verified_catalog.v1"
	KnownAppVerifiedCatalogPreviewRequestType         = "known-app-verified-catalog-preview"
	KnownAppVerifiedCatalogRunPlanSchemaVersion       = "xnix.runtime.known_app_verified_catalog_run_plan.v1"
	KnownAppVerifiedCatalogRunPlanRequestType         = "known-app-verified-catalog-run-plan-preview"
	KnownAppVerifiedCatalogRunAcceptanceSchemaVersion = "xnix.runtime.known_app_verified_catalog_run_acceptance.v1"
	KnownAppVerifiedCatalogRunAcceptanceRequestType   = "known-app-verified-catalog-run-acceptance-preview"
)

var knownAppVerifiedCatalogRequiredAppIDs = []string{"7zr", "busybox-w32"}

type KnownAppVerifiedCatalogRequest struct {
	MatrixEvidencePath string
}

type KnownAppVerifiedCatalogRunPlanRequest struct {
	VerifiedCatalogPath string
	AppID               string
}

type KnownAppVerifiedCatalogRunAcceptanceRequest struct {
	RunPlanPath   string
	RunReportPath string
}

type KnownAppVerifiedCatalogPreview struct {
	SchemaVersion                      string                               `json:"schema_version"`
	RequestType                        string                               `json:"request_type"`
	Source                             string                               `json:"source"`
	Desktop                            string                               `json:"desktop"`
	RuntimeMethod                      string                               `json:"runtime_method"`
	ReadMethod                         string                               `json:"read_method"`
	VerificationSource                 string                               `json:"verification_source"`
	MatrixEvidenceConsumed             bool                                 `json:"matrix_evidence_consumed"`
	MatrixStatus                       string                               `json:"matrix_status"`
	RequiredAppIDs                     []string                             `json:"required_app_ids"`
	ApplicationIDs                     []string                             `json:"application_ids"`
	ApplicationCount                   int                                  `json:"application_count"`
	VerifiedApplicationCount           int                                  `json:"verified_application_count"`
	QEMUExecutedCount                  int                                  `json:"qemu_executed_count"`
	WineExecutedCount                  int                                  `json:"wine_executed_count"`
	ChecksumVerifiedCount              int                                  `json:"checksum_verified_count"`
	MarkerObservedCount                int                                  `json:"marker_observed_count"`
	RawOutputRedactedCount             int                                  `json:"raw_output_redacted_count"`
	CompatibilityCenterProjectionReady bool                                 `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                                 `json:"kde_center_projection_ready"`
	Applications                       []KnownAppVerifiedCatalogApplication `json:"applications"`
	RuntimeOwned                       bool                                 `json:"runtime_owned"`
	GoRuntimeBacked                    bool                                 `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                                 `json:"kde_policy_owner"`
	UserVisible                        bool                                 `json:"user_visible"`
	StandardDesktopEntries             bool                                 `json:"standard_desktop_entries"`
	ReviewOnly                         bool                                 `json:"review_only"`
	OperatorReviewRequired             bool                                 `json:"operator_review_required"`
	LaunchEnabled                      bool                                 `json:"launch_enabled"`
	ExecutionStarted                   bool                                 `json:"execution_started"`
	BackendLaunchEnabled               bool                                 `json:"backend_launch_enabled"`
	DesktopFilesWritten                bool                                 `json:"desktop_files_written"`
	HostRootModified                   bool                                 `json:"host_root_modified"`
	BackendDetailsExposed              bool                                 `json:"backend_details_exposed"`
	RawOutputExposed                   bool                                 `json:"raw_output_exposed"`
	RemotePathExposed                  bool                                 `json:"remote_path_exposed"`
	BlockedActions                     []string                             `json:"blocked_actions"`
	DesktopSafeSummary                 string                               `json:"desktop_safe_summary"`
}

type KnownAppVerifiedCatalogApplication struct {
	AppID                  string   `json:"app_id"`
	DisplayName            string   `json:"display_name"`
	AppVersion             string   `json:"app_version"`
	VerificationState      string   `json:"verification_state"`
	CompatibilityState     string   `json:"compatibility_state"`
	DesktopCatalogState    string   `json:"desktop_catalog_state"`
	LauncherSurface        string   `json:"launcher_surface"`
	LaunchRequestCommand   []string `json:"launch_request_command"`
	PrimaryActionID        string   `json:"primary_action_id"`
	PrimaryActionKind      string   `json:"primary_action_kind"`
	DirectLaunchEnabled    bool     `json:"direct_launch_enabled"`
	OperatorReviewRequired bool     `json:"operator_review_required"`
	RuntimeOwned           bool     `json:"runtime_owned"`
	GoRuntimeBacked        bool     `json:"go_runtime_backed"`
	KDEPolicyOwner         bool     `json:"kde_policy_owner"`
	QEMUExecuted           bool     `json:"qemu_executed"`
	WineExecuted           bool     `json:"wine_executed"`
	ChecksumVerified       bool     `json:"checksum_verified"`
	MarkerObserved         bool     `json:"marker_observed"`
	RawOutputRedacted      bool     `json:"raw_output_redacted"`
	SerialLogEvidence      bool     `json:"serial_log_evidence"`
	BackendLaunchEnabled   bool     `json:"backend_launch_enabled"`
	BackendDetailsExposed  bool     `json:"backend_details_exposed"`
	RawOutputExposed       bool     `json:"raw_output_exposed"`
	RemotePathExposed      bool     `json:"remote_path_exposed"`
	HostRootModified       bool     `json:"host_root_modified"`
	DesktopSafeSummary     string   `json:"desktop_safe_summary"`
}

type KnownAppVerifiedCatalogRunPlanPreview struct {
	SchemaVersion                   string   `json:"schema_version"`
	RequestType                     string   `json:"request_type"`
	Source                          string   `json:"source"`
	Desktop                         string   `json:"desktop"`
	RuntimeMethod                   string   `json:"runtime_method"`
	ReadMethod                      string   `json:"read_method"`
	VerifiedCatalogConsumed         bool     `json:"verified_catalog_consumed"`
	RequestedAppID                  string   `json:"requested_app_id"`
	AppID                           string   `json:"app_id"`
	DisplayName                     string   `json:"display_name"`
	AppVersion                      string   `json:"app_version"`
	VerificationState               string   `json:"verification_state"`
	CompatibilityState              string   `json:"compatibility_state"`
	DesktopCatalogState             string   `json:"desktop_catalog_state"`
	LauncherSurface                 string   `json:"launcher_surface"`
	LaunchRequestCommand            []string `json:"launch_request_command"`
	RemoteSmokeCommand              []string `json:"remote_smoke_command"`
	RemoteSmokeRequestType          string   `json:"remote_smoke_request_type"`
	Q4ExecutionRequired             bool     `json:"q4_execution_required"`
	Q4ExecutionPlanned              bool     `json:"q4_execution_planned"`
	Q4ExecutionStarted              bool     `json:"q4_execution_started"`
	ReviewOnly                      bool     `json:"review_only"`
	OperatorReviewRequired          bool     `json:"operator_review_required"`
	RuntimeOwned                    bool     `json:"runtime_owned"`
	GoRuntimeBacked                 bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool     `json:"kde_policy_owner"`
	DirectLaunchEnabled             bool     `json:"direct_launch_enabled"`
	LaunchEnabled                   bool     `json:"launch_enabled"`
	ExecutionStarted                bool     `json:"execution_started"`
	BackendLaunchEnabled            bool     `json:"backend_launch_enabled"`
	DesktopFilesWritten             bool     `json:"desktop_files_written"`
	HostRootModified                bool     `json:"host_root_modified"`
	BackendDetailsExposed           bool     `json:"backend_details_exposed"`
	RawOutputExposed                bool     `json:"raw_output_exposed"`
	RemotePathExposed               bool     `json:"remote_path_exposed"`
	PrivilegedContainerRequired     bool     `json:"privileged_container_required"`
	HostNetworkingRequired          bool     `json:"host_networking_required"`
	DockerSocketMounted             bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired          bool     `json:"broad_host_mount_required"`
	HostCompilationRequired         bool     `json:"host_compilation_required"`
	HostCompilationAvoided          bool     `json:"host_compilation_avoided"`
	TargetedRemoteVerificationReady bool     `json:"targeted_remote_verification_ready"`
	BlockedActions                  []string `json:"blocked_actions"`
	DesktopSafeSummary              string   `json:"desktop_safe_summary"`
}

type KnownAppVerifiedCatalogRunAcceptancePreview struct {
	Version                              string `json:"version"`
	SchemaVersion                        string `json:"schema_version"`
	RequestType                          string `json:"request_type"`
	Source                               string `json:"source"`
	RuntimeMethod                        string `json:"runtime_method"`
	ReadMethod                           string `json:"read_method"`
	AcceptanceType                       string `json:"acceptance_type"`
	RunPlanConsumed                      bool   `json:"run_plan_consumed"`
	RunReportConsumed                    bool   `json:"run_report_consumed"`
	RunPlanPathExposed                   bool   `json:"run_plan_path_exposed"`
	RunReportPathExposed                 bool   `json:"run_report_path_exposed"`
	RemoteHostExposed                    bool   `json:"remote_host_exposed"`
	RawPathExposed                       bool   `json:"raw_path_exposed"`
	RawOutputExposed                     bool   `json:"raw_output_exposed"`
	RuntimeArgvExposed                   bool   `json:"runtime_argv_exposed"`
	RunnerPathExposed                    bool   `json:"runner_path_exposed"`
	RequestedAppID                       string `json:"requested_app_id"`
	AppID                                string `json:"app_id"`
	DisplayName                          string `json:"display_name"`
	AppVersion                           string `json:"app_version"`
	VerificationState                    string `json:"verification_state"`
	CompatibilityState                   string `json:"compatibility_state"`
	DesktopCatalogState                  string `json:"desktop_catalog_state"`
	RunPlanMatched                       bool   `json:"run_plan_matched"`
	ExistingWindowsApp                   bool   `json:"existing_windows_app"`
	KnownPortableCatalogBacked           bool   `json:"known_portable_catalog_backed"`
	LaunchAttempted                      bool   `json:"launch_attempted"`
	ChecksumVerified                     bool   `json:"checksum_verified"`
	MarkerObserved                       bool   `json:"marker_observed"`
	RuntimeStartedIsolatedGuest          bool   `json:"runtime_started_isolated_guest"`
	IsolatedGuestExecutionObserved       bool   `json:"isolated_guest_execution_observed"`
	CompatibilityEngineExecutionObserved bool   `json:"compatibility_engine_execution_observed"`
	LoopbackOnlyNetworking               bool   `json:"loopback_only_networking"`
	SerialLogPersisted                   bool   `json:"serial_log_persisted"`
	OutputRedacted                       bool   `json:"output_redacted"`
	Q4ExecutionObserved                  bool   `json:"q4_execution_observed"`
	HostCompilationAvoided               bool   `json:"host_compilation_avoided"`
	NetworkRequired                      bool   `json:"network_required"`
	HostRootModified                     bool   `json:"host_root_modified"`
	PrivilegedContainerRequired          bool   `json:"privileged_container_required"`
	HostNetworkingRequired               bool   `json:"host_networking_required"`
	DockerSocketMounted                  bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool   `json:"broad_host_mount_required"`
	DockerExecuted                       bool   `json:"docker_executed"`
	ColimaExecuted                       bool   `json:"colima_executed"`
	NetworkChecksRun                     bool   `json:"network_checks_run"`
	PackageManagerInvoked                bool   `json:"package_manager_invoked"`
	AcceptanceReady                      bool   `json:"acceptance_ready"`
	DesktopSafeSummary                   string `json:"desktop_safe_summary"`
}

func PreviewKnownAppVerifiedCatalog(request KnownAppVerifiedCatalogRequest) (KnownAppVerifiedCatalogPreview, error) {
	if request.MatrixEvidencePath == "" {
		return KnownAppVerifiedCatalogPreview{}, errors.New("known app verified catalog requires --matrix-evidence")
	}
	content, err := os.ReadFile(request.MatrixEvidencePath)
	if err != nil {
		return KnownAppVerifiedCatalogPreview{}, fmt.Errorf("read known app matrix evidence: %w", err)
	}
	return PreviewKnownAppVerifiedCatalogJSON(content)
}

func PreviewKnownAppVerifiedCatalogJSON(content []byte) (KnownAppVerifiedCatalogPreview, error) {
	var evidence KnownAppMatrixEvidencePreview
	if err := json.Unmarshal(content, &evidence); err != nil {
		return KnownAppVerifiedCatalogPreview{}, fmt.Errorf("parse known app matrix evidence: %w", err)
	}
	if err := validateKnownAppVerifiedCatalogEvidence(evidence); err != nil {
		return KnownAppVerifiedCatalogPreview{}, err
	}

	applications := make([]KnownAppVerifiedCatalogApplication, 0, len(evidence.Apps))
	for _, app := range evidence.Apps {
		if !slices.Contains(knownAppVerifiedCatalogRequiredAppIDs, app.AppID) {
			continue
		}
		applications = append(applications, knownAppVerifiedCatalogApplication(app))
	}

	return KnownAppVerifiedCatalogPreview{
		SchemaVersion:                      KnownAppVerifiedCatalogPreviewSchemaVersion,
		RequestType:                        KnownAppVerifiedCatalogPreviewRequestType,
		Source:                             "known-app-matrix-evidence+runtime-verified-catalog",
		Desktop:                            "KDE Plasma",
		RuntimeMethod:                      "ListKnownVerifiedApplications",
		ReadMethod:                         "ListKnownVerifiedApplicationsPreview",
		VerificationSource:                 KnownAppMatrixEvidencePreviewRequestType,
		MatrixEvidenceConsumed:             true,
		MatrixStatus:                       evidence.MatrixStatus,
		RequiredAppIDs:                     append([]string(nil), knownAppVerifiedCatalogRequiredAppIDs...),
		ApplicationIDs:                     knownAppVerifiedCatalogApplicationIDs(applications),
		ApplicationCount:                   len(applications),
		VerifiedApplicationCount:           len(applications),
		QEMUExecutedCount:                  evidence.QEMUExecutedCount,
		WineExecutedCount:                  evidence.WineExecutedCount,
		ChecksumVerifiedCount:              evidence.ChecksumVerifiedCount,
		MarkerObservedCount:                evidence.MarkerObservedCount,
		RawOutputRedactedCount:             evidence.RawOutputRedactedCount,
		CompatibilityCenterProjectionReady: evidence.CompatibilityCenterProjectionReady,
		KDECenterProjectionReady:           evidence.KDECenterProjectionReady,
		Applications:                       applications,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		UserVisible:                        true,
		StandardDesktopEntries:             true,
		ReviewOnly:                         true,
		OperatorReviewRequired:             true,
		LaunchEnabled:                      false,
		ExecutionStarted:                   false,
		BackendLaunchEnabled:               false,
		DesktopFilesWritten:                false,
		HostRootModified:                   false,
		BackendDetailsExposed:              false,
		RawOutputExposed:                   false,
		RemotePathExposed:                  false,
		BlockedActions: []string{
			"launch verified app directly from catalog preview",
			"write desktop files from catalog preview",
			"start compatibility execution from catalog preview",
			"expose remote q4 paths from catalog preview",
			"expose raw app output from catalog preview",
		},
		DesktopSafeSummary: fmt.Sprintf("%d known Windows apps are verified by q4 matrix evidence and visible as review-only Runtime catalog entries.", len(applications)),
	}, nil
}

func PreviewKnownAppVerifiedCatalogRunPlan(request KnownAppVerifiedCatalogRunPlanRequest) (KnownAppVerifiedCatalogRunPlanPreview, error) {
	if request.VerifiedCatalogPath == "" {
		return KnownAppVerifiedCatalogRunPlanPreview{}, errors.New("known app verified catalog run plan requires --verified-catalog")
	}
	content, err := os.ReadFile(request.VerifiedCatalogPath)
	if err != nil {
		return KnownAppVerifiedCatalogRunPlanPreview{}, fmt.Errorf("read known app verified catalog: %w", err)
	}
	return PreviewKnownAppVerifiedCatalogRunPlanJSON(content, request.AppID)
}

func PreviewKnownAppVerifiedCatalogRunPlanJSON(content []byte, appID string) (KnownAppVerifiedCatalogRunPlanPreview, error) {
	if appID == "" {
		return KnownAppVerifiedCatalogRunPlanPreview{}, errors.New("known app verified catalog run plan requires --app")
	}
	var catalog KnownAppVerifiedCatalogPreview
	if err := json.Unmarshal(content, &catalog); err != nil {
		return KnownAppVerifiedCatalogRunPlanPreview{}, fmt.Errorf("parse known app verified catalog: %w", err)
	}
	normalized, err := normalizeKnownAppVerifiedCatalog(&catalog)
	if err != nil {
		return KnownAppVerifiedCatalogRunPlanPreview{}, err
	}
	var selected KnownAppVerifiedCatalogApplication
	found := false
	for _, app := range normalized.Applications {
		if app.AppID == appID {
			selected = app
			found = true
			break
		}
	}
	if !found {
		return KnownAppVerifiedCatalogRunPlanPreview{}, fmt.Errorf("known app verified catalog run plan cannot find app %s", appID)
	}

	remoteSmokeCommand := []string{"ruby", "scripts/remote_known_winapp_guest_wine_smoke.rb", "--execute", "--app", selected.AppID}
	preview := KnownAppVerifiedCatalogRunPlanPreview{
		SchemaVersion:                   KnownAppVerifiedCatalogRunPlanSchemaVersion,
		RequestType:                     KnownAppVerifiedCatalogRunPlanRequestType,
		Source:                          "known-app-verified-catalog+q4-run-plan",
		Desktop:                         "KDE Plasma",
		RuntimeMethod:                   "PlanKnownVerifiedApplicationRun",
		ReadMethod:                      "GetKnownVerifiedApplicationRunPlan",
		VerifiedCatalogConsumed:         true,
		RequestedAppID:                  appID,
		AppID:                           selected.AppID,
		DisplayName:                     selected.DisplayName,
		AppVersion:                      selected.AppVersion,
		VerificationState:               selected.VerificationState,
		CompatibilityState:              selected.CompatibilityState,
		DesktopCatalogState:             selected.DesktopCatalogState,
		LauncherSurface:                 selected.LauncherSurface,
		LaunchRequestCommand:            append([]string(nil), selected.LaunchRequestCommand...),
		RemoteSmokeCommand:              remoteSmokeCommand,
		RemoteSmokeRequestType:          "remote-known-winapp-guest-wine-smoke",
		Q4ExecutionRequired:             true,
		Q4ExecutionPlanned:              true,
		Q4ExecutionStarted:              false,
		ReviewOnly:                      true,
		OperatorReviewRequired:          true,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		DirectLaunchEnabled:             false,
		LaunchEnabled:                   false,
		ExecutionStarted:                false,
		BackendLaunchEnabled:            false,
		DesktopFilesWritten:             false,
		HostRootModified:                false,
		BackendDetailsExposed:           false,
		RawOutputExposed:                false,
		RemotePathExposed:               false,
		PrivilegedContainerRequired:     false,
		HostNetworkingRequired:          false,
		DockerSocketMounted:             false,
		BroadHostMountRequired:          false,
		HostCompilationRequired:         false,
		HostCompilationAvoided:          true,
		TargetedRemoteVerificationReady: true,
		BlockedActions: []string{
			"start q4 execution from run-plan preview",
			"launch verified app directly from KDE catalog card",
			"write desktop files from run-plan preview",
			"expose remote q4 paths from run-plan preview",
			"expose raw app output from run-plan preview",
			"compile known app runner on the host",
		},
		DesktopSafeSummary: selected.DisplayName + " is ready for an operator-triggered q4 known Windows app run using the verified catalog entry.",
	}
	if err := validateNoBackendTerms(preview, "known app verified catalog run plan"); err != nil {
		return KnownAppVerifiedCatalogRunPlanPreview{}, err
	}
	return preview, nil
}

func PreviewKnownAppVerifiedCatalogRunAcceptance(request KnownAppVerifiedCatalogRunAcceptanceRequest) (KnownAppVerifiedCatalogRunAcceptancePreview, error) {
	if request.RunPlanPath == "" {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, errors.New("known app verified catalog run acceptance requires --run-plan")
	}
	if request.RunReportPath == "" {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, errors.New("known app verified catalog run acceptance requires --known-winapp-run")
	}
	runPlanContent, err := os.ReadFile(request.RunPlanPath)
	if err != nil {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, fmt.Errorf("read known app verified catalog run plan: %w", err)
	}
	runReportContent, err := os.ReadFile(request.RunReportPath)
	if err != nil {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, fmt.Errorf("read known app run report: %w", err)
	}
	return PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent, runReportContent)
}

func PreviewKnownAppVerifiedCatalogRunAcceptanceJSON(runPlanContent []byte, runReportContent []byte) (KnownAppVerifiedCatalogRunAcceptancePreview, error) {
	var runPlan KnownAppVerifiedCatalogRunPlanPreview
	if err := json.Unmarshal(runPlanContent, &runPlan); err != nil {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, fmt.Errorf("parse known app verified catalog run plan: %w", err)
	}
	if runPlan.SchemaVersion != KnownAppVerifiedCatalogRunPlanSchemaVersion ||
		runPlan.RequestType != KnownAppVerifiedCatalogRunPlanRequestType ||
		!runPlan.VerifiedCatalogConsumed ||
		!runPlan.Q4ExecutionPlanned ||
		runPlan.Q4ExecutionStarted ||
		!runPlan.ReviewOnly ||
		runPlan.DirectLaunchEnabled ||
		runPlan.LaunchEnabled ||
		runPlan.ExecutionStarted ||
		runPlan.BackendLaunchEnabled ||
		runPlan.HostRootModified ||
		runPlan.BackendDetailsExposed ||
		runPlan.RawOutputExposed ||
		runPlan.RemotePathExposed ||
		!runPlan.HostCompilationAvoided {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, errors.New("known app verified catalog run acceptance requires a safe q4 run plan")
	}

	knownAcceptance, err := PreviewKnownExistingWinAppAcceptanceJSON(runReportContent)
	if err != nil {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, err
	}
	if runPlan.AppID != knownAcceptance.AppID {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, errors.New("known app verified catalog run acceptance requires run-plan and run-report app ids to match")
	}
	if !knownAcceptance.AcceptanceReady {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, errors.New("known app verified catalog run acceptance requires accepted known app run evidence")
	}

	preview := KnownAppVerifiedCatalogRunAcceptancePreview{
		Version:                              knownAcceptance.Version,
		SchemaVersion:                        KnownAppVerifiedCatalogRunAcceptanceSchemaVersion,
		RequestType:                          KnownAppVerifiedCatalogRunAcceptanceRequestType,
		Source:                               "known-app-verified-catalog-run-plan+known-existing-windows-app-acceptance",
		RuntimeMethod:                        "PreviewKnownVerifiedApplicationRunAcceptance",
		ReadMethod:                           "GetKnownVerifiedApplicationRunAcceptance",
		AcceptanceType:                       "verified-catalog-app-q4-real-run-acceptance",
		RunPlanConsumed:                      true,
		RunReportConsumed:                    true,
		RunPlanPathExposed:                   false,
		RunReportPathExposed:                 false,
		RemoteHostExposed:                    false,
		RawPathExposed:                       false,
		RawOutputExposed:                     false,
		RuntimeArgvExposed:                   false,
		RunnerPathExposed:                    false,
		RequestedAppID:                       runPlan.RequestedAppID,
		AppID:                                knownAcceptance.AppID,
		DisplayName:                          knownAcceptance.DisplayName,
		AppVersion:                           knownAcceptance.AppVersion,
		VerificationState:                    runPlan.VerificationState,
		CompatibilityState:                   runPlan.CompatibilityState,
		DesktopCatalogState:                  runPlan.DesktopCatalogState,
		RunPlanMatched:                       true,
		ExistingWindowsApp:                   knownAcceptance.ExistingWindowsApp,
		KnownPortableCatalogBacked:           knownAcceptance.KnownPortableCatalogBacked,
		LaunchAttempted:                      knownAcceptance.LaunchAttempted,
		ChecksumVerified:                     knownAcceptance.ChecksumVerified,
		MarkerObserved:                       knownAcceptance.MarkerObserved,
		RuntimeStartedIsolatedGuest:          knownAcceptance.RuntimeStartedIsolatedGuest,
		IsolatedGuestExecutionObserved:       knownAcceptance.IsolatedGuestExecutionObserved,
		CompatibilityEngineExecutionObserved: knownAcceptance.CompatibilityEngineExecutionObserved,
		LoopbackOnlyNetworking:               knownAcceptance.LoopbackOnlyNetworking,
		SerialLogPersisted:                   knownAcceptance.SerialLogPersisted,
		OutputRedacted:                       knownAcceptance.OutputRedacted,
		Q4ExecutionObserved:                  true,
		HostCompilationAvoided:               runPlan.HostCompilationAvoided,
		NetworkRequired:                      false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		DockerExecuted:                       false,
		ColimaExecuted:                       false,
		NetworkChecksRun:                     false,
		PackageManagerInvoked:                false,
		AcceptanceReady:                      true,
		DesktopSafeSummary:                   knownAcceptance.DisplayName + " matched the verified catalog run plan and completed the q4 known Windows app acceptance lane.",
	}
	if err := validateNoBackendTerms(preview, "known app verified catalog run acceptance"); err != nil {
		return KnownAppVerifiedCatalogRunAcceptancePreview{}, err
	}
	return preview, nil
}

func validateKnownAppVerifiedCatalogEvidence(evidence KnownAppMatrixEvidencePreview) error {
	if evidence.SchemaVersion != KnownAppMatrixEvidencePreviewSchemaVersion || evidence.RequestType != KnownAppMatrixEvidencePreviewRequestType {
		return errors.New("known app verified catalog requires known-app-matrix-evidence-preview input")
	}
	if evidence.MatrixStatus != "passed" {
		return errors.New("known app verified catalog requires passed matrix evidence")
	}
	if evidence.MatrixReportPathExposed || evidence.RawOutputExposed || evidence.RemotePathExposed || evidence.BackendDetailsExposed {
		return errors.New("known app verified catalog requires redacted matrix evidence")
	}
	if evidence.HostRootModified || evidence.PrivilegedContainerRequired || evidence.HostNetworkingRequired || evidence.DockerSocketMounted || evidence.BroadHostMountRequired {
		return errors.New("known app verified catalog requires closed host and container gates")
	}
	appsByID := map[string]KnownAppMatrixEvidenceApp{}
	for _, app := range evidence.Apps {
		appsByID[app.AppID] = app
	}
	for _, appID := range knownAppVerifiedCatalogRequiredAppIDs {
		app, ok := appsByID[appID]
		if !ok {
			return fmt.Errorf("known app verified catalog missing required app %s", appID)
		}
		if !knownAppVerifiedCatalogAppReady(app) {
			return fmt.Errorf("known app verified catalog app %s is not verified", appID)
		}
	}
	if !evidence.MatrixEvidenceReady() {
		return errors.New("known app verified catalog requires complete matrix evidence")
	}
	return nil
}

func (evidence KnownAppMatrixEvidencePreview) MatrixEvidenceReady() bool {
	required := len(knownAppVerifiedCatalogRequiredAppIDs)
	return evidence.MatrixReportConsumed &&
		evidence.MatrixReportOutputWritten &&
		evidence.AppCount >= required &&
		evidence.PassedCount >= required &&
		evidence.FailedCount == 0 &&
		evidence.EvidenceCount >= required &&
		evidence.PassedEvidenceCount >= required &&
		evidence.FailedEvidenceCount == 0 &&
		evidence.QEMUExecutedCount >= required &&
		evidence.WineExecutedCount >= required &&
		evidence.ChecksumVerifiedCount >= required &&
		evidence.MarkerObservedCount >= required &&
		evidence.RawOutputRedactedCount >= required &&
		evidence.SerialLogEvidenceCount >= required &&
		evidence.CompatibilityCenterProjectionReady &&
		evidence.KDECenterProjectionReady &&
		evidence.RuntimeOwned &&
		evidence.GoRuntimeBacked &&
		!evidence.KDEPolicyOwner &&
		!evidence.DesktopLaunchEnabled &&
		!evidence.BackendLaunchEnabled &&
		!evidence.ActionExecutionEnabled
}

func knownAppVerifiedCatalogAppReady(app KnownAppMatrixEvidenceApp) bool {
	return app.SmokeStatus == "passed" &&
		app.CompatibilityState == "real-qemu-wine-verified" &&
		app.MarkerObserved &&
		app.ChecksumVerified &&
		app.QEMUExecuted &&
		app.WineExecuted &&
		app.GuestStarted &&
		app.GuestPortAuto &&
		app.RawOutputRedacted &&
		app.SerialLogEvidence &&
		app.ReportEvidence &&
		app.RuntimeOwned &&
		app.GoRuntimeBacked &&
		!app.KDEPolicyOwner &&
		!app.DesktopLaunchEnabled &&
		!app.BackendLaunchEnabled &&
		!app.BackendDetailsExposed &&
		!app.RawOutputExposed &&
		!app.RemotePathExposed &&
		!app.HostRootModified
}

func knownAppVerifiedCatalogApplication(app KnownAppMatrixEvidenceApp) KnownAppVerifiedCatalogApplication {
	return KnownAppVerifiedCatalogApplication{
		AppID:                  app.AppID,
		DisplayName:            app.DisplayName,
		AppVersion:             app.AppVersion,
		VerificationState:      "verified-real-q4-matrix-run",
		CompatibilityState:     app.CompatibilityState,
		DesktopCatalogState:    "visible-review-only",
		LauncherSurface:        "xnix-compat-launch",
		LaunchRequestCommand:   []string{"xnix-compat-launch", "--app", app.AppID},
		PrimaryActionID:        "review-known-app-matrix-evidence",
		PrimaryActionKind:      "review",
		DirectLaunchEnabled:    false,
		OperatorReviewRequired: true,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		QEMUExecuted:           app.QEMUExecuted,
		WineExecuted:           app.WineExecuted,
		ChecksumVerified:       app.ChecksumVerified,
		MarkerObserved:         app.MarkerObserved,
		RawOutputRedacted:      app.RawOutputRedacted,
		SerialLogEvidence:      app.SerialLogEvidence,
		BackendLaunchEnabled:   false,
		BackendDetailsExposed:  false,
		RawOutputExposed:       false,
		RemotePathExposed:      false,
		HostRootModified:       false,
		DesktopSafeSummary:     app.DisplayName + " is verified by q4 matrix evidence and available as a review-only Runtime catalog entry.",
	}
}

func knownAppVerifiedCatalogApplicationIDs(applications []KnownAppVerifiedCatalogApplication) []string {
	ids := make([]string, 0, len(applications))
	for _, app := range applications {
		ids = append(ids, app.AppID)
	}
	slices.Sort(ids)
	return ids
}
