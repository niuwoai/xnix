package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPreviewQ4ExternalWinAppRunPlanKeepsLocalHostAsUploaderOnly(t *testing.T) {
	preview, err := PreviewQ4ExternalWinAppRunPlan(Q4ExternalWinAppRunPlanRequest{
		Version:             "0.2.640-test",
		ProjectRoot:         "/workspace/xnix",
		LocalExecutable:     "/tmp/xnix-upload/app.exe",
		AppID:               "org.xnix.external.notepadplusplus",
		DisplayName:         "Notepad++",
		WindowMatch:         "Notepad++",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
		Output:              "/tmp/xnix-output/run.json",
		MarkdownOutput:      "/tmp/xnix-output/run.md",
		Execute:             true,
	})
	if err != nil {
		t.Fatalf("PreviewQ4ExternalWinAppRunPlan returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.q4_external_winapp_run_plan.v1" ||
		preview.RequestType != "q4-external-winapp-run-plan-preview" ||
		preview.RuntimeMethod != "PlanQ4ExternalWinAppRun" ||
		preview.ReadMethod != "GetQ4ExternalWinAppRunPlanPreview" ||
		preview.AppID != "org.xnix.external.notepadplusplus" ||
		preview.DisplayName != "Notepad++" ||
		preview.WindowMatchConfigured != true ||
		preview.LocalExecutablePathClass != "scoped-temporary" ||
		preview.RemoteMaterialsRootClass != "q4-home-scoped" ||
		preview.DelegatedScript != "scripts/q4_external_winapp_run.rb" ||
		preview.DelegatedRequestType != "q4-external-winapp-run" ||
		preview.DelegatedAcceptanceRequestType != "q4-staged-external-winapp-acceptance-preview" ||
		preview.AcceptedApplicationDetailStateRequired != "runtime-accepted-real-app-run" ||
		preview.LocalHostRole != "scoped-upload-and-operator-plan-only" {
		t.Fatalf("unexpected q4 external Windows app run plan: %#v", preview)
	}
	if !preview.Q4MaterializationRequired ||
		!preview.Q4BuildPreferred ||
		!preview.Q4ExecutionRequired ||
		preview.HostCompilationRequired ||
		!preview.HostCompilationAvoided ||
		!preview.LocalExecutableMZHeaderRequired ||
		!preview.LocalExecutableSHA256Required ||
		!preview.RemoteUploadRequired ||
		!preview.RealFileOpenEvidenceRequired ||
		!preview.WindowsProcessWindowObservationRequired ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.FullSmokeRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		!preview.PlanReady {
		t.Fatalf("unexpected q4 external Windows app run safety flags: %#v", preview)
	}
	if !containsQ4ExternalRunCommandArg(preview.DelegatedCommand, "--execute") ||
		!containsQ4ExternalRunCommandArg(preview.DelegatedCommand, "--executable") ||
		!containsQ4ExternalRunCommandArg(preview.DelegatedCommand, "/tmp/xnix-upload/app.exe") ||
		!containsQ4ExternalRunCommandArg(preview.DelegatedCommand, "--window-match") ||
		!containsQ4ExternalRunCommandArg(preview.DelegatedCommand, "Notepad++") {
		t.Fatalf("delegated command did not preserve operator inputs: %#v", preview.DelegatedCommand)
	}
	payload, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	for _, forbidden := range []string{"docker.sock", "--privileged", "--network host"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("q4 external Windows app run plan exposed forbidden term %q: %s", forbidden, string(payload))
		}
	}
	if strings.Contains(preview.DesktopSafeSummary, "/tmp/xnix-upload") ||
		strings.Contains(preview.DesktopSafeSummary, "root@q4") {
		t.Fatalf("desktop summary exposed operator paths or hosts: %s", preview.DesktopSafeSummary)
	}
}

func TestPreviewQ4ExternalWinAppRunPlanRejectsBroadLocalPaths(t *testing.T) {
	_, err := PreviewQ4ExternalWinAppRunPlan(Q4ExternalWinAppRunPlanRequest{
		Version:             "0.2.640-test",
		ProjectRoot:         "/workspace/xnix",
		LocalExecutable:     "/Users/rocky/Downloads/app.exe",
		AppID:               "org.xnix.external.app",
		DisplayName:         "Uploaded App",
		WindowMatch:         "Uploaded App",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
	})
	if err == nil || !strings.Contains(err.Error(), "local executable must stay under this checkout or /tmp/xnix-*") {
		t.Fatalf("expected broad local path rejection, got %v", err)
	}
}

func TestPreviewQ4ExternalWinAppRunPlanRejectsMissingWindowMatch(t *testing.T) {
	_, err := PreviewQ4ExternalWinAppRunPlan(Q4ExternalWinAppRunPlanRequest{
		Version:             "0.2.640-test",
		ProjectRoot:         "/workspace/xnix",
		LocalExecutable:     "/tmp/xnix-upload/app.exe",
		AppID:               "org.xnix.external.app",
		DisplayName:         "Uploaded App",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
	})
	if err == nil || !strings.Contains(err.Error(), "window match must be non-empty") {
		t.Fatalf("expected missing window match rejection, got %v", err)
	}
}

func containsQ4ExternalRunCommandArg(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}
