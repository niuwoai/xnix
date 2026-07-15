package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBackendBindingPreviewPlansProfileBindingWithoutCommitting(t *testing.T) {
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

	preview, err := plan.BackendBindingPreview()
	if err != nil {
		t.Fatalf("BackendBindingPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.backend_binding.v1" ||
		preview.RequestType != "backend-binding-preview" ||
		preview.BindingType != "compatibility-backend-binding" ||
		preview.Source != "backend-selection-preview+backend-environment-preview+execution-readiness-preview" ||
		preview.RuntimeMethod != "GetBackendBinding" {
		t.Fatalf("unexpected backend binding schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected backend binding identity: %#v", preview)
	}
	if preview.SelectedStrategy != "local-compatibility-strategy" ||
		preview.RecommendedProfileID != "local-compatibility" ||
		preview.ProfileKind != "local" ||
		preview.BindingState != "planned-blocked" ||
		preview.BindingKey != "org.example.ledger:local-compatibility" {
		t.Fatalf("unexpected backend binding recommendation: %#v", preview)
	}
	if preview.Selection.RequestType != "backend-selection-preview" ||
		preview.Selection.RecommendedProfileID != "local-compatibility" ||
		preview.Selection.CandidateCount != 2 ||
		preview.Selection.SelectionCommitted ||
		preview.Selection.BackendLaunchEnabled ||
		preview.Selection.BackendDetailsExposed {
		t.Fatalf("unexpected selection summary: %#v", preview.Selection)
	}
	if preview.Environment.RequestType != "backend-environment-preview" ||
		preview.Environment.EnvironmentState != "planned-blocked" ||
		preview.Environment.ProfileCount != 2 ||
		preview.Environment.BridgeCapabilityCount != 5 ||
		preview.Environment.EnvironmentCreated ||
		preview.Environment.BackendProcessStarted ||
		preview.Environment.BackendBindingReady ||
		preview.Environment.BackendDetailsExposed {
		t.Fatalf("unexpected environment summary: %#v", preview.Environment)
	}
	if got, want := preview.PreflightGateIDs, []string{"recipe-validation", "package-source-review", "application-state-root", "portal-policy-review", "snapshot-baseline", "runtime-launch-write-gate"}; !sameStrings(got, want) {
		t.Fatalf("PreflightGateIDs = %#v, want %#v", got, want)
	}
	if preview.PreflightGateCount != 6 ||
		preview.PassedGateCount != 1 ||
		preview.PendingGateCount != 4 ||
		preview.BlockedGateCount != 1 {
		t.Fatalf("unexpected gate counts: %#v", preview)
	}
	if got, want := preview.RequiredReviews, []string{"package-source-review", "application-state-root-review", "portal-policy-review", "snapshot-baseline-review", "runtime-launch-write-gate"}; !sameStrings(got, want) {
		t.Fatalf("RequiredReviews = %#v, want %#v", got, want)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner || !preview.UserVisible ||
		preview.ManagedBindingReady || preview.BindingCommitted || preview.BindingPersisted ||
		preview.LaunchEnabled || preview.ExecutionRequestCreated ||
		preview.EnvironmentCreated || preview.BackendProcessStarted ||
		preview.StateRootCreated || preview.PortalRequestCreated ||
		preview.SnapshotCreated || preview.HostRootModified ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed || preview.CompatibilityStorageExposed ||
		preview.RawBackendCommandExposed {
		t.Fatalf("unexpected backend binding safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "commit compatibility profile binding from preview") ||
		!containsString(preview.BlockedActions, "start compatibility backend from backend binding preview") ||
		!containsString(preview.BlockedActions, "create execution request from backend binding preview") {
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
	isolatedPreview, err := isolatedPlan.BackendBindingPreview()
	if err != nil {
		t.Fatalf("isolated BackendBindingPreview returned error: %v", err)
	}
	if isolatedPreview.RecommendedProfileID != "isolated-compatibility" ||
		isolatedPreview.ProfileKind != "isolated" ||
		isolatedPreview.BindingKey != "org.example.isolated:isolated-compatibility" {
		t.Fatalf("isolated recipe did not recommend isolated binding: %#v", isolatedPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("backend binding preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
