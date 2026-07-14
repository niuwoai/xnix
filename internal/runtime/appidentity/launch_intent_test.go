package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLaunchIntentPreviewCapturesIntentWithoutExecution(t *testing.T) {
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

	preview, err := plan.LaunchIntentPreview(nil)
	if err != nil {
		t.Fatalf("LaunchIntentPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.launch_intent.v1" ||
		preview.RequestType != "launch-intent-preview" ||
		preview.IntentType != "runtime-launch-intent" ||
		preview.Source != "desktop-launcher" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetLaunchIntent" {
		t.Fatalf("unexpected launch intent schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected launch intent identity: %#v", preview)
	}
	if preview.ExecutionState != "blocked" ||
		preview.OverallStatus != "not-ready" ||
		preview.WriteGateDecision != "blocked-until-production-backend" ||
		preview.DenialErrorName != "org.xnix.Compatibility1.Error.WriteMethodDisabled" {
		t.Fatalf("unexpected launch gate state: %#v", preview)
	}
	if preview.PortalRequired ||
		preview.FileCount != 0 ||
		len(preview.FileURIs) != 0 {
		t.Fatalf("plain launch intent should not require files: %#v", preview)
	}
	if preview.CompatibilityProfile.ID != "local-compatibility" ||
		preview.CompatibilityProfile.Ready ||
		preview.CompatibilityProfile.LaunchEnabled {
		t.Fatalf("unexpected launch profile: %#v", preview.CompatibilityProfile)
	}
	if preview.RunPlan.PlanType != "compatibility-run" ||
		preview.RunPlan.Strategy != "automatic-managed" ||
		preview.RunPlan.BackendDetailsExposed ||
		preview.RunPlan.BackendReady ||
		!preview.RunPlan.PortalPolicyRequired ||
		!preview.RunPlan.SnapshotBeforeRiskyChange ||
		!preview.RunPlan.RuntimeWriteGateRequired ||
		preview.RunPlan.ExecutionRequestCreated {
		t.Fatalf("unexpected launch run plan: %#v", preview.RunPlan)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.StandardDesktopEntry || !preview.LaunchUsesRuntime ||
		!preview.DesktopEntryLaunchVisible || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.ExecutionStarted || preview.BackendBindingReady ||
		preview.RequestObjectCreated || preview.PermissionGranted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected launch intent safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "start compatibility profile from KDE") ||
		!containsString(preview.BlockedActions, "expose raw backend command to desktop shell") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	filePreview, err := plan.LaunchIntentPreview([]string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("file LaunchIntentPreview returned error: %v", err)
	}
	if !filePreview.PortalRequired ||
		filePreview.FileCount != 1 ||
		filePreview.FileURIs[0] != "file:///home/test/Documents/book.xls" ||
		filePreview.RequestObjectCreated ||
		filePreview.PermissionGranted {
		t.Fatalf("unexpected file launch intent: %#v", filePreview)
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
	isolatedPreview, err := isolatedPlan.LaunchIntentPreview(nil)
	if err != nil {
		t.Fatalf("isolated LaunchIntentPreview returned error: %v", err)
	}
	if isolatedPreview.CompatibilityProfile.ID != "isolated-compatibility" ||
		isolatedPreview.CompatibilityProfile.Kind != "isolated" {
		t.Fatalf("isolated launch intent did not recommend isolated compatibility: %#v", isolatedPreview)
	}

	if _, err := plan.LaunchIntentPreview([]string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("LaunchIntentPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(filePreview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("launch intent preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
