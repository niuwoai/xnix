package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionRequestPreviewBuildsBlockedRequestIntake(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	preview, err := plan.ExecutionRequestPreview(nil)
	if err != nil {
		t.Fatalf("ExecutionRequestPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.request_intake.v1" ||
		preview.RequestType != "execution-request-preview" ||
		preview.IntentType != "runtime-launch-intent" ||
		preview.RequestState != "blocked" ||
		preview.Source != "runtime-launch-intent" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionRequestPreview" {
		t.Fatalf("unexpected execution request schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected execution request identity: %#v", preview)
	}
	if preview.LaunchIntent.Source != "desktop-launcher" ||
		preview.LaunchIntent.IntentType != "runtime-launch-intent" ||
		preview.LaunchIntent.RuntimeMethod != "Launch" ||
		preview.LaunchIntent.ReadMethod != "GetLaunchIntent" ||
		preview.LaunchIntent.LaunchAllowed ||
		preview.LaunchIntent.LaunchEnabled ||
		preview.LaunchIntent.RequestObjectCreated ||
		preview.LaunchIntent.ExecutionStarted ||
		preview.LaunchIntent.BackendDetailsExposed {
		t.Fatalf("unexpected launch intent summary: %#v", preview.LaunchIntent)
	}
	if preview.CompatibilityProfile.ID != "local-compatibility" ||
		preview.CompatibilityProfile.Ready ||
		preview.CompatibilityProfile.LaunchEnabled ||
		preview.CompatibilityProfile.BackendDetailsExposed {
		t.Fatalf("unexpected compatibility profile: %#v", preview.CompatibilityProfile)
	}
	if preview.GateSummary.GateCount != 5 ||
		preview.GateSummary.RequiredGateCount != 2 ||
		preview.GateSummary.PendingGateCount != 1 ||
		preview.GateSummary.BlockedGateCount != 1 {
		t.Fatalf("unexpected gate summary: %#v", preview.GateSummary)
	}
	if preview.ExecutionState != "blocked" ||
		preview.OverallStatus != "not-ready" ||
		preview.WriteGateDecision != "blocked-until-production-backend" ||
		preview.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected request gate state: %#v", preview)
	}
	if preview.PortalRequired ||
		preview.FileCount != 0 ||
		len(preview.FileURIs) != 0 {
		t.Fatalf("plain execution request should not require files: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionRequestCreated || preview.ExecutionRequestPersisted ||
		preview.ExecutionStarted || preview.BackendBindingReady ||
		preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution request safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "persist execution request before Runtime gates pass") ||
		!containsString(preview.BlockedActions, "start compatibility profile from execution request preview") ||
		!containsString(preview.BlockedActions, "expose raw backend command to desktop shell") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	filePreview, err := plan.ExecutionRequestPreview([]string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("file ExecutionRequestPreview returned error: %v", err)
	}
	if !filePreview.PortalRequired ||
		filePreview.FileCount != 1 ||
		filePreview.FileURIs[0] != "file:///home/test/Documents/book.xls" ||
		!filePreview.LaunchIntent.PortalRequired ||
		filePreview.LaunchIntent.FileCount != 1 ||
		filePreview.LaunchIntent.FileURIs[0] != "file:///home/test/Documents/book.xls" ||
		filePreview.RequestObjectCreated ||
		filePreview.PermissionGranted {
		t.Fatalf("unexpected file execution request: %#v", filePreview)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "applications-games",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.ExecutionRequestPreview(nil)
	if err != nil {
		t.Fatalf("isolated ExecutionRequestPreview returned error: %v", err)
	}
	if isolatedPreview.CompatibilityProfile.ID != "isolated-compatibility" ||
		isolatedPreview.CompatibilityProfile.Kind != "isolated" {
		t.Fatalf("isolated execution request did not recommend isolated compatibility: %#v", isolatedPreview)
	}

	if _, err := plan.ExecutionRequestPreview([]string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionRequestPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(filePreview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution request preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
