package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewKnownAppVerifiedCatalogConsumesMatrixEvidence(t *testing.T) {
	content := knownAppVerifiedCatalogMatrixEvidenceFixture(t)

	preview, err := PreviewKnownAppVerifiedCatalogJSON(content)
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppVerifiedCatalogPreviewSchemaVersion ||
		preview.RequestType != KnownAppVerifiedCatalogPreviewRequestType ||
		preview.Source != "known-app-matrix-evidence+runtime-verified-catalog" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "ListKnownVerifiedApplications" ||
		preview.ReadMethod != "ListKnownVerifiedApplicationsPreview" ||
		preview.VerificationSource != KnownAppMatrixEvidencePreviewRequestType ||
		!preview.MatrixEvidenceConsumed ||
		preview.MatrixStatus != "passed" {
		t.Fatalf("unexpected verified catalog schema: %#v", preview)
	}
	if preview.ApplicationCount != 2 ||
		preview.VerifiedApplicationCount != 2 ||
		preview.QEMUExecutedCount != 2 ||
		preview.WineExecutedCount != 2 ||
		preview.ChecksumVerifiedCount != 2 ||
		preview.MarkerObservedCount != 2 ||
		preview.RawOutputRedactedCount != 2 {
		t.Fatalf("unexpected verified catalog counts: %#v", preview)
	}
	if !preview.CompatibilityCenterProjectionReady ||
		!preview.KDECenterProjectionReady ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.StandardDesktopEntries ||
		!preview.ReviewOnly ||
		!preview.OperatorReviewRequired ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.DesktopFilesWritten ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed {
		t.Fatalf("unexpected verified catalog safety flags: %#v", preview)
	}
	if strings.Join(preview.ApplicationIDs, ",") != "7zr,busybox-w32" {
		t.Fatalf("unexpected application ids: %#v", preview.ApplicationIDs)
	}
	busybox := preview.Applications[1]
	if busybox.AppID != "busybox-w32" ||
		busybox.DisplayName != "BusyBox-w32 standalone console executable" ||
		busybox.VerificationState != "verified-real-q4-matrix-run" ||
		busybox.DesktopCatalogState != "visible-review-only" ||
		busybox.LauncherSurface != "xnix-compat-launch" ||
		strings.Join(busybox.LaunchRequestCommand, " ") != "xnix-compat-launch --app busybox-w32" ||
		busybox.PrimaryActionKind != "review" ||
		busybox.DirectLaunchEnabled ||
		!busybox.OperatorReviewRequired ||
		!busybox.QEMUExecuted ||
		!busybox.WineExecuted ||
		!busybox.ChecksumVerified ||
		!busybox.MarkerObserved ||
		!busybox.RawOutputRedacted ||
		!busybox.SerialLogEvidence ||
		busybox.BackendLaunchEnabled ||
		busybox.BackendDetailsExposed ||
		busybox.RawOutputExposed ||
		busybox.RemotePathExposed ||
		busybox.HostRootModified {
		t.Fatalf("unexpected BusyBox verified catalog entry: %#v", busybox)
	}
}

func TestPreviewKnownAppVerifiedCatalogRejectsIncompleteMatrixEvidence(t *testing.T) {
	content := knownAppVerifiedCatalogMatrixEvidenceFixture(t)
	var evidence KnownAppMatrixEvidencePreview
	if err := json.Unmarshal(content, &evidence); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	evidence.Apps = evidence.Apps[:1]
	evidence.AppCount = 1
	malformed, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	_, err = PreviewKnownAppVerifiedCatalogJSON(malformed)
	if err == nil || !strings.Contains(err.Error(), "busybox-w32") {
		t.Fatalf("expected missing BusyBox evidence error, got %v", err)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunPlanSelectsQ4RunnableApp(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}

	preview, err := PreviewKnownAppVerifiedCatalogRunPlanJSON(content, "busybox-w32")
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogRunPlanJSON returned error: %v", err)
	}
	if preview.SchemaVersion != KnownAppVerifiedCatalogRunPlanSchemaVersion ||
		preview.RequestType != KnownAppVerifiedCatalogRunPlanRequestType ||
		preview.Source != "known-app-verified-catalog+q4-run-plan" ||
		preview.RuntimeMethod != "PlanKnownVerifiedApplicationRun" ||
		preview.ReadMethod != "GetKnownVerifiedApplicationRunPlan" ||
		!preview.VerifiedCatalogConsumed ||
		preview.RequestedAppID != "busybox-w32" ||
		preview.AppID != "busybox-w32" ||
		preview.VerificationState != "verified-real-q4-matrix-run" ||
		preview.DesktopCatalogState != "visible-review-only" {
		t.Fatalf("unexpected run plan schema: %#v", preview)
	}
	if strings.Join(preview.LaunchRequestCommand, " ") != "xnix-compat-launch --app busybox-w32" ||
		strings.Join(preview.RemoteSmokeCommand, " ") != "ruby scripts/remote_known_winapp_guest_wine_smoke.rb --execute --app busybox-w32" ||
		preview.RemoteSmokeRequestType != "remote-known-winapp-guest-wine-smoke" {
		t.Fatalf("unexpected run commands: %#v", preview)
	}
	if !preview.Q4ExecutionRequired ||
		!preview.Q4ExecutionPlanned ||
		preview.Q4ExecutionStarted ||
		!preview.ReviewOnly ||
		!preview.OperatorReviewRequired ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DirectLaunchEnabled ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendLaunchEnabled ||
		preview.DesktopFilesWritten ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed ||
		preview.RawOutputExposed ||
		preview.RemotePathExposed ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		preview.HostCompilationRequired ||
		!preview.HostCompilationAvoided ||
		!preview.TargetedRemoteVerificationReady {
		t.Fatalf("unexpected run plan safety flags: %#v", preview)
	}
}

func TestPreviewKnownAppVerifiedCatalogRunPlanRejectsUnknownApp(t *testing.T) {
	catalog, err := PreviewKnownAppVerifiedCatalogJSON(knownAppVerifiedCatalogMatrixEvidenceFixture(t))
	if err != nil {
		t.Fatalf("PreviewKnownAppVerifiedCatalogJSON returned error: %v", err)
	}
	content, err := json.Marshal(catalog)
	if err != nil {
		t.Fatalf("Marshal catalog returned error: %v", err)
	}

	_, err = PreviewKnownAppVerifiedCatalogRunPlanJSON(content, "missing-app")
	if err == nil || !strings.Contains(err.Error(), "missing-app") {
		t.Fatalf("expected unknown app error, got %v", err)
	}
}

func knownAppVerifiedCatalogMatrixEvidenceFixture(t *testing.T) []byte {
	t.Helper()
	apps := []KnownAppMatrixEvidenceApp{
		knownAppVerifiedCatalogMatrixEvidenceApp("7zr", "7-Zip standalone console executable", "26.02"),
		knownAppVerifiedCatalogMatrixEvidenceApp("busybox-w32", "BusyBox-w32 standalone console executable", "current-2026-07-24"),
	}
	evidence := KnownAppMatrixEvidencePreview{
		SchemaVersion:                      KnownAppMatrixEvidencePreviewSchemaVersion,
		RequestType:                        KnownAppMatrixEvidencePreviewRequestType,
		Source:                             "remote-known-winapp-matrix-smoke+runtime-evidence-consumer",
		RuntimeMethod:                      "PreviewKnownAppMatrixEvidence",
		ReadMethod:                         "GetKnownAppMatrixEvidence",
		MatrixStatus:                       "passed",
		MatrixReportConsumed:               true,
		MatrixReportPathExposed:            false,
		MatrixReportOutputWritten:          true,
		AppCount:                           2,
		PassedCount:                        2,
		FailedCount:                        0,
		EvidenceCount:                      2,
		PassedEvidenceCount:                2,
		FailedEvidenceCount:                0,
		QEMUExecutedCount:                  2,
		WineExecutedCount:                  2,
		MarkerObservedCount:                2,
		ChecksumVerifiedCount:              2,
		RawOutputRedactedCount:             2,
		SerialLogEvidenceCount:             2,
		GuestStartedCount:                  2,
		GuestPortAutoCount:                 2,
		CompatibilityCenterProjectionReady: true,
		KDECenterProjectionReady:           true,
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
	}
	content, err := json.Marshal(evidence)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	return content
}

func knownAppVerifiedCatalogMatrixEvidenceApp(appID string, displayName string, appVersion string) KnownAppMatrixEvidenceApp {
	return KnownAppMatrixEvidenceApp{
		AppID:                 appID,
		DisplayName:           displayName,
		AppVersion:            appVersion,
		SmokeStatus:           "passed",
		CompatibilityState:    "real-qemu-wine-verified",
		MarkerObserved:        true,
		ChecksumVerified:      true,
		QEMUExecuted:          true,
		WineExecuted:          true,
		GuestStarted:          true,
		GuestPortAuto:         true,
		RawOutputRedacted:     true,
		SerialLogEvidence:     true,
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
	}
}
