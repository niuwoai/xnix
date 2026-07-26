package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

const (
	KnownAppVerifiedCatalogPreviewSchemaVersion = "xnix.runtime.known_app_verified_catalog.v1"
	KnownAppVerifiedCatalogPreviewRequestType   = "known-app-verified-catalog-preview"
)

var knownAppVerifiedCatalogRequiredAppIDs = []string{"7zr", "busybox-w32"}

type KnownAppVerifiedCatalogRequest struct {
	MatrixEvidencePath string
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
