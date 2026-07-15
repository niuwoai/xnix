package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendEnvironmentPreviewPlansProfilesAndBridgesWithoutCreatingEnvironment(t *testing.T) {
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

	preview, err := plan.BackendEnvironmentPreview()
	if err != nil {
		t.Fatalf("BackendEnvironmentPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.backend_environment.v1" ||
		preview.RequestType != "backend-environment-preview" ||
		preview.PlanType != "compatibility-backend-environment-plan" ||
		preview.Source != "backend-selection-preview+desktop-resource-bridge-preview" ||
		preview.RuntimeMethod != "GetBackendEnvironmentPlan" {
		t.Fatalf("unexpected backend environment schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected backend environment identity: %#v", preview)
	}
	if preview.SelectedStrategy != "local-compatibility-strategy" ||
		preview.RecommendedProfileID != "local-compatibility" ||
		preview.EnvironmentState != "planned-blocked" ||
		preview.ProfileCount != 2 ||
		preview.ReadyProfileCount != 0 ||
		preview.BlockedProfileCount != 2 {
		t.Fatalf("unexpected backend environment recommendation: %#v", preview)
	}
	if got, want := preview.ProfileIDs, []string{"local-compatibility-environment", "isolated-compatibility-environment"}; !sameStrings(got, want) {
		t.Fatalf("ProfileIDs = %#v, want %#v", got, want)
	}
	localProfile := preview.Profiles[0]
	if localProfile.ID != "local-compatibility-environment" ||
		localProfile.Kind != "local" ||
		localProfile.Status != "blocked" ||
		!localProfile.Recommended ||
		localProfile.Ready ||
		!localProfile.Blocked ||
		localProfile.EnvironmentCreated ||
		localProfile.BackendProcessStarted ||
		localProfile.StoragePathExposed ||
		localProfile.BackendDetailsExposed ||
		localProfile.PrivilegedContainerUsed ||
		!localProfile.CompatibilityProfileOnly {
		t.Fatalf("unexpected local environment profile: %#v", localProfile)
	}
	if got, want := preview.RequiredReviews, []string{"package-source-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review"}; !sameStrings(got, want) {
		t.Fatalf("RequiredReviews = %#v, want %#v", got, want)
	}
	if got, want := preview.BridgeCapabilityIDs, []string{"file-open", "uri-open", "clipboard", "print", "screenshot"}; !sameStrings(got, want) {
		t.Fatalf("BridgeCapabilityIDs = %#v, want %#v", got, want)
	}
	if preview.BridgeCapabilityCount != 5 {
		t.Fatalf("unexpected bridge capability count: %#v", preview)
	}
	for _, bridge := range preview.BridgeCapabilities {
		if bridge.RuntimeMethod != "GetPortalRequestPlan" ||
			bridge.State != "planned" ||
			!bridge.PortalRequired ||
			!bridge.UserReviewRequired ||
			bridge.BridgeEnabled ||
			bridge.RequestObjectCreated ||
			bridge.DirectBackendAccessAllowed ||
			bridge.BackendDetailsExposed {
			t.Fatalf("bridge gate unexpectedly open: %#v", bridge)
		}
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.LocalEnvironmentReady || preview.IsolatedEnvironmentReady ||
		preview.EnvironmentCreated || preview.BackendProcessStarted ||
		preview.BackendBindingReady || preview.LaunchEnabled ||
		preview.RequestObjectCreated || preview.HostStorageExposed ||
		preview.ClipboardBridgeEnabled || preview.PrintBridgeEnabled ||
		preview.FileBridgeEnabled || !preview.PortalReviewRequired ||
		!preview.SnapshotRequired || preview.HostRootModified ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed || preview.CompatibilityStorageExposed ||
		preview.RawBackendCommandExposed {
		t.Fatalf("unexpected backend environment safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "create local compatibility environment from KDE") ||
		!containsString(preview.BlockedActions, "create isolated compatibility environment from KDE") ||
		!containsString(preview.BlockedActions, "start compatibility backend from environment preview") {
		t.Fatalf("missing blocked actions: %#v", preview.BlockedActions)
	}

	isolatedPlan, err := NewPlan(Recipe{
		ID:                  "org.example.isolated",
		Name:                "Example Isolated",
		Icon:                "application-x-executable",
		Mode:                "vm",
		SupportedExtensions: []string{".abc"},
	})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolatedPlan.BackendEnvironmentPreview()
	if err != nil {
		t.Fatalf("isolated BackendEnvironmentPreview returned error: %v", err)
	}
	if isolatedPreview.RecommendedProfileID != "isolated-compatibility" ||
		!isolatedPreview.Profiles[1].Recommended {
		t.Fatalf("isolated recipe did not recommend isolated environment: %#v", isolatedPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend environment preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
