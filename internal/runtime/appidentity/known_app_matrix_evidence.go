package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppMatrixEvidencePreviewSchemaVersion = "xnix.runtime.known_app_matrix_evidence_preview.v1"
	KnownAppMatrixEvidencePreviewRequestType   = "known-app-matrix-evidence-preview"
)

type KnownAppMatrixEvidencePreviewRequest struct {
	MatrixReportPath string
}

type KnownAppMatrixEvidencePreview struct {
	SchemaVersion                      string                         `json:"schema_version"`
	RequestType                        string                         `json:"request_type"`
	Source                             string                         `json:"source"`
	RuntimeMethod                      string                         `json:"runtime_method"`
	ReadMethod                         string                         `json:"read_method"`
	MatrixStatus                       string                         `json:"matrix_status"`
	MatrixReportConsumed               bool                           `json:"matrix_report_consumed"`
	MatrixReportPathExposed            bool                           `json:"matrix_report_path_exposed"`
	MatrixReportOutputWritten          bool                           `json:"matrix_report_output_written"`
	AppCount                           int                            `json:"app_count"`
	PassedCount                        int                            `json:"passed_count"`
	FailedCount                        int                            `json:"failed_count"`
	EvidenceCount                      int                            `json:"evidence_count"`
	PassedEvidenceCount                int                            `json:"passed_evidence_count"`
	FailedEvidenceCount                int                            `json:"failed_evidence_count"`
	QEMUExecutedCount                  int                            `json:"qemu_executed_count"`
	WineExecutedCount                  int                            `json:"wine_executed_count"`
	MarkerObservedCount                int                            `json:"marker_observed_count"`
	ChecksumVerifiedCount              int                            `json:"checksum_verified_count"`
	RawOutputRedactedCount             int                            `json:"raw_output_redacted_count"`
	SerialLogEvidenceCount             int                            `json:"serial_log_evidence_count"`
	GuestStartedCount                  int                            `json:"guest_started_count"`
	GuestPortAutoCount                 int                            `json:"guest_port_auto_count"`
	CompatibilityCenterProjectionReady bool                           `json:"compatibility_center_projection_ready"`
	KDECenterProjectionReady           bool                           `json:"kde_center_projection_ready"`
	KnownAppSmokeEvidence              []KnownAppSmokeEvidenceSummary `json:"known_app_smoke_evidence"`
	Apps                               []KnownAppMatrixEvidenceApp    `json:"apps"`
	RuntimeOwned                       bool                           `json:"runtime_owned"`
	GoRuntimeBacked                    bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner                     bool                           `json:"kde_policy_owner"`
	DesktopLaunchEnabled               bool                           `json:"desktop_launch_enabled"`
	BackendLaunchEnabled               bool                           `json:"backend_launch_enabled"`
	ActionExecutionEnabled             bool                           `json:"action_execution_enabled"`
	BackendDetailsExposed              bool                           `json:"backend_details_exposed"`
	RawOutputExposed                   bool                           `json:"raw_output_exposed"`
	RemotePathExposed                  bool                           `json:"remote_path_exposed"`
	HostRootModified                   bool                           `json:"host_root_modified"`
	PrivilegedContainerRequired        bool                           `json:"privileged_container_required"`
	HostNetworkingRequired             bool                           `json:"host_networking_required"`
	DockerSocketMounted                bool                           `json:"docker_socket_mounted"`
	BroadHostMountRequired             bool                           `json:"broad_host_mount_required"`
	DesktopSafeSummary                 string                         `json:"desktop_safe_summary"`
}

type KnownAppMatrixEvidenceApp struct {
	AppID                 string `json:"app_id"`
	DisplayName           string `json:"display_name"`
	AppVersion            string `json:"app_version"`
	ExecutableName        string `json:"executable_name"`
	SmokeStatus           string `json:"smoke_status"`
	CompatibilityState    string `json:"compatibility_state"`
	MarkerObserved        bool   `json:"marker_observed"`
	ChecksumVerified      bool   `json:"checksum_verified"`
	QEMUExecuted          bool   `json:"qemu_executed"`
	WineExecuted          bool   `json:"wine_executed"`
	GuestStarted          bool   `json:"guest_started"`
	GuestPortAuto         bool   `json:"guest_port_auto"`
	RawOutputRedacted     bool   `json:"raw_output_redacted"`
	SerialLogEvidence     bool   `json:"serial_log_evidence"`
	ReportEvidence        bool   `json:"report_evidence"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	GoRuntimeBacked       bool   `json:"go_runtime_backed"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	DesktopLaunchEnabled  bool   `json:"desktop_launch_enabled"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	RawOutputExposed      bool   `json:"raw_output_exposed"`
	RemotePathExposed     bool   `json:"remote_path_exposed"`
	HostRootModified      bool   `json:"host_root_modified"`
	DesktopSafeSummary    string `json:"desktop_safe_summary"`
}

type knownAppMatrixReport struct {
	SchemaVersion               string                    `json:"schema_version"`
	RequestType                 string                    `json:"request_type"`
	Status                      string                    `json:"status"`
	Execute                     bool                      `json:"execute"`
	MatrixReportOutputWritten   bool                      `json:"matrix_report_output_written"`
	AppCount                    int                       `json:"app_count"`
	PassedCount                 int                       `json:"passed_count"`
	FailedCount                 int                       `json:"failed_count"`
	Backend                     string                    `json:"backend"`
	StartQEMU                   bool                      `json:"start_qemu"`
	GuestPort                   string                    `json:"guest_port"`
	RedactOutput                bool                      `json:"redact_output"`
	Apps                        []knownAppMatrixReportApp `json:"apps"`
	HostRootModified            bool                      `json:"host_root_modified"`
	PrivilegedContainerRequired bool                      `json:"privileged_container_required"`
	HostNetworkingRequired      bool                      `json:"host_networking_required"`
	DockerSocketMounted         bool                      `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool                      `json:"broad_host_mount_required"`
}

type knownAppMatrixReportApp struct {
	AppID                string `json:"app_id"`
	Status               string `json:"status"`
	AppVersion           string `json:"app_version"`
	ExecutableName       string `json:"executable_name"`
	Backend              string `json:"backend"`
	GuestStarted         bool   `json:"guest_started"`
	GuestPortAuto        bool   `json:"guest_port_auto"`
	ChecksumVerified     bool   `json:"checksum_verified"`
	RawOutputRedacted    bool   `json:"raw_output_redacted"`
	MarkerObserved       bool   `json:"marker_observed"`
	QEMUSerialLogWritten bool   `json:"qemu_serial_log_written"`
	QEMUExecuted         bool   `json:"qemu_executed"`
	WineExecuted         bool   `json:"wine_executed"`
}

func PreviewKnownAppMatrixEvidence(request KnownAppMatrixEvidencePreviewRequest) (KnownAppMatrixEvidencePreview, error) {
	path := strings.TrimSpace(request.MatrixReportPath)
	if path == "" {
		return KnownAppMatrixEvidencePreview{}, errors.New("known app matrix evidence requires --matrix-report")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return KnownAppMatrixEvidencePreview{}, fmt.Errorf("read known app matrix report: %w", err)
	}
	return PreviewKnownAppMatrixEvidenceJSON(content)
}

func PreviewKnownAppMatrixEvidenceJSON(content []byte) (KnownAppMatrixEvidencePreview, error) {
	var report knownAppMatrixReport
	if err := json.Unmarshal(content, &report); err != nil {
		return KnownAppMatrixEvidencePreview{}, fmt.Errorf("parse known app matrix report: %w", err)
	}
	if err := validateKnownAppMatrixReport(report); err != nil {
		return KnownAppMatrixEvidencePreview{}, err
	}

	apps := make([]KnownAppMatrixEvidenceApp, 0, len(report.Apps))
	evidence := make([]KnownAppSmokeEvidenceSummary, 0, len(report.Apps))
	passedEvidenceCount := 0
	qemuExecutedCount := 0
	wineExecutedCount := 0
	markerObservedCount := 0
	checksumVerifiedCount := 0
	rawOutputRedactedCount := 0
	serialLogEvidenceCount := 0
	guestStartedCount := 0
	guestPortAutoCount := 0

	for _, item := range report.Apps {
		app, err := knownAppMatrixEvidenceApp(item)
		if err != nil {
			return KnownAppMatrixEvidencePreview{}, err
		}
		apps = append(apps, app)
		evidence = append(evidence, knownAppMatrixSmokeEvidenceSummary(app))
		if app.SmokeStatus == "passed" {
			passedEvidenceCount++
		}
		if app.QEMUExecuted {
			qemuExecutedCount++
		}
		if app.WineExecuted {
			wineExecutedCount++
		}
		if app.MarkerObserved {
			markerObservedCount++
		}
		if app.ChecksumVerified {
			checksumVerifiedCount++
		}
		if app.RawOutputRedacted {
			rawOutputRedactedCount++
		}
		if app.SerialLogEvidence {
			serialLogEvidenceCount++
		}
		if app.GuestStarted {
			guestStartedCount++
		}
		if app.GuestPortAuto {
			guestPortAutoCount++
		}
	}

	projectionReady := report.Status == "passed" && passedEvidenceCount == len(report.Apps) && qemuExecutedCount == len(report.Apps) && wineExecutedCount == len(report.Apps) && rawOutputRedactedCount == len(report.Apps)
	return KnownAppMatrixEvidencePreview{
		SchemaVersion:                      KnownAppMatrixEvidencePreviewSchemaVersion,
		RequestType:                        KnownAppMatrixEvidencePreviewRequestType,
		Source:                             "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
		RuntimeMethod:                      "PreviewKnownAppMatrixEvidence",
		ReadMethod:                         "GetKnownAppMatrixEvidence",
		MatrixStatus:                       report.Status,
		MatrixReportConsumed:               true,
		MatrixReportPathExposed:            false,
		MatrixReportOutputWritten:          report.MatrixReportOutputWritten,
		AppCount:                           report.AppCount,
		PassedCount:                        report.PassedCount,
		FailedCount:                        report.FailedCount,
		EvidenceCount:                      len(evidence),
		PassedEvidenceCount:                passedEvidenceCount,
		FailedEvidenceCount:                len(evidence) - passedEvidenceCount,
		QEMUExecutedCount:                  qemuExecutedCount,
		WineExecutedCount:                  wineExecutedCount,
		MarkerObservedCount:                markerObservedCount,
		ChecksumVerifiedCount:              checksumVerifiedCount,
		RawOutputRedactedCount:             rawOutputRedactedCount,
		SerialLogEvidenceCount:             serialLogEvidenceCount,
		GuestStartedCount:                  guestStartedCount,
		GuestPortAutoCount:                 guestPortAutoCount,
		CompatibilityCenterProjectionReady: projectionReady,
		KDECenterProjectionReady:           projectionReady,
		KnownAppSmokeEvidence:              evidence,
		Apps:                               apps,
		RuntimeOwned:                       true,
		GoRuntimeBacked:                    true,
		KDEPolicyOwner:                     false,
		DesktopLaunchEnabled:               false,
		BackendLaunchEnabled:               false,
		ActionExecutionEnabled:             false,
		BackendDetailsExposed:              false,
		RawOutputExposed:                   false,
		RemotePathExposed:                  false,
		HostRootModified:                   false,
		PrivilegedContainerRequired:        false,
		HostNetworkingRequired:             false,
		DockerSocketMounted:                false,
		BroadHostMountRequired:             false,
		DesktopSafeSummary:                 fmt.Sprintf("%d known Windows apps were consumed from matrix evidence with %d passed Runtime-owned QEMU/Wine runs.", report.AppCount, report.PassedCount),
	}, nil
}

func validateKnownAppMatrixReport(report knownAppMatrixReport) error {
	switch {
	case report.SchemaVersion != "xnix.scripts.remote_known_windows_app_matrix_smoke.v1":
		return errors.New("known app matrix evidence requires the remote matrix smoke schema")
	case report.RequestType != "remote-known-winapp-matrix-smoke":
		return errors.New("known app matrix evidence requires a remote known-app matrix smoke report")
	case !report.Execute:
		return errors.New("known app matrix evidence requires an executed report")
	case report.Status == "":
		return errors.New("known app matrix evidence requires matrix status")
	case report.AppCount != len(report.Apps):
		return errors.New("known app matrix evidence app count does not match apps")
	case report.Backend != winapp.KnownRunBackendGuestWine:
		return errors.New("known app matrix evidence requires the guest-wine backend")
	case !report.StartQEMU:
		return errors.New("known app matrix evidence requires Runtime-started QEMU")
	case report.GuestPort != "auto":
		return errors.New("known app matrix evidence requires automatic loopback guest ports")
	case !report.RedactOutput:
		return errors.New("known app matrix evidence requires redacted output")
	case report.HostRootModified || report.PrivilegedContainerRequired || report.HostNetworkingRequired || report.DockerSocketMounted || report.BroadHostMountRequired:
		return errors.New("known app matrix evidence must keep host and container boundaries closed")
	}
	passed := 0
	failed := 0
	for _, app := range report.Apps {
		if err := validateKnownAppMatrixReportApp(app); err != nil {
			return err
		}
		if app.Status == "passed" {
			passed++
		} else {
			failed++
		}
	}
	if report.PassedCount != passed || report.FailedCount != failed {
		return errors.New("known app matrix evidence pass/fail counts do not match apps")
	}
	return nil
}

func validateKnownAppMatrixReportApp(app knownAppMatrixReportApp) error {
	if _, err := winapp.LookupKnownPortableApp(app.AppID); err != nil {
		return err
	}
	for _, value := range []string{app.AppID, app.Status, app.AppVersion, app.ExecutableName, app.Backend} {
		if value != "" && !singleLine(value) {
			return errors.New("known app matrix evidence requires single-line app fields")
		}
	}
	if app.Backend != winapp.KnownRunBackendGuestWine {
		return errors.New("known app matrix evidence app requires the guest-wine backend")
	}
	if app.Status == "passed" {
		switch {
		case !app.GuestStarted:
			return errors.New("passed known app matrix evidence requires a started guest")
		case !app.GuestPortAuto:
			return errors.New("passed known app matrix evidence requires an automatic guest port")
		case !app.ChecksumVerified:
			return errors.New("passed known app matrix evidence requires checksum verification")
		case !app.RawOutputRedacted:
			return errors.New("passed known app matrix evidence requires redacted output")
		case !app.MarkerObserved:
			return errors.New("passed known app matrix evidence requires marker observation")
		case !app.QEMUSerialLogWritten:
			return errors.New("passed known app matrix evidence requires serial log evidence")
		case !app.QEMUExecuted:
			return errors.New("passed known app matrix evidence requires QEMU execution")
		case !app.WineExecuted:
			return errors.New("passed known app matrix evidence requires Wine execution")
		}
	}
	return nil
}

func knownAppMatrixEvidenceApp(item knownAppMatrixReportApp) (KnownAppMatrixEvidenceApp, error) {
	catalogApp, err := winapp.LookupKnownPortableApp(item.AppID)
	if err != nil {
		return KnownAppMatrixEvidenceApp{}, err
	}
	state := "matrix-evidence-needs-review"
	if item.Status == "passed" {
		state = "real-qemu-wine-verified"
	}
	return KnownAppMatrixEvidenceApp{
		AppID:                 catalogApp.ID,
		DisplayName:           catalogApp.DisplayName,
		AppVersion:            item.AppVersion,
		ExecutableName:        item.ExecutableName,
		SmokeStatus:           item.Status,
		CompatibilityState:    state,
		MarkerObserved:        item.MarkerObserved,
		ChecksumVerified:      item.ChecksumVerified,
		QEMUExecuted:          item.QEMUExecuted,
		WineExecuted:          item.WineExecuted,
		GuestStarted:          item.GuestStarted,
		GuestPortAuto:         item.GuestPortAuto,
		RawOutputRedacted:     item.RawOutputRedacted,
		SerialLogEvidence:     item.QEMUSerialLogWritten,
		ReportEvidence:        true,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		DesktopLaunchEnabled:  false,
		BackendLaunchEnabled:  false,
		BackendDetailsExposed: false,
		RawOutputExposed:      false,
		RemotePathExposed:     false,
		HostRootModified:      false,
		DesktopSafeSummary:    catalogApp.DisplayName + " has redacted real QEMU/Wine matrix evidence.",
	}, nil
}

func knownAppMatrixSmokeEvidenceSummary(app KnownAppMatrixEvidenceApp) KnownAppSmokeEvidenceSummary {
	return KnownAppSmokeEvidenceSummary{
		AppID:                              app.AppID,
		DisplayName:                        app.DisplayName,
		AppVersion:                         app.AppVersion,
		EvidenceKind:                       "known-application-matrix-smoke",
		EvidenceSource:                     "remote-known-winapp-matrix-smoke",
		SmokeStatus:                        app.SmokeStatus,
		CompatibilityState:                 app.CompatibilityState,
		CenterCardState:                    "validated-real-runtime-run",
		LaunchAuthorizationState:           "not-recorded",
		PrimaryActionID:                    "review-known-app-matrix-evidence",
		PrimaryActionLabel:                 "Review real run evidence",
		PrimaryActionKind:                  "review",
		PrimaryActionEnabled:               true,
		DirectLaunchEnabled:                false,
		LaunchAuthorizationReceiptRequired: true,
		LaunchAuthorizationReceiptState:    "missing",
		MarkerObserved:                     app.MarkerObserved,
		ChecksumVerified:                   app.ChecksumVerified,
		ExecutionEvidenceRecorded:          true,
		StagedLauncherVerified:             false,
		RuntimeDispatchVerified:            true,
		LaunchAuthorizationRequired:        true,
		DesktopLaunchEnabled:               false,
		RuntimeOwned:                       true,
		KDEPolicyOwner:                     false,
		ActionExecutionEnabled:             false,
		BackendLaunchEnabled:               false,
		HostRootModified:                   false,
		BackendDetailsExposed:              false,
		RawArtifactPathExposed:             false,
		Summary:                            app.DisplayName + " real run matrix evidence is ready for Compatibility Center review.",
	}
}
