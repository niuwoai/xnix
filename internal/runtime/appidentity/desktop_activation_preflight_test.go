package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationPreflightPreviewBlocksProductionDevelopmentRegistry(t *testing.T) {
	plan, err := NewPlanWithProvenance(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls", ".abc"},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}

	preview, err := plan.DesktopActivationPreflightPreview("production")
	if err != nil {
		t.Fatalf("DesktopActivationPreflightPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_preflight.v1" ||
		preview.RequestType != "desktop-activation-preflight-preview" ||
		preview.PreflightType != "normal-linux-application-activation-preflight" ||
		preview.RuntimeMethod != "GetDesktopActivationPreflight" {
		t.Fatalf("unexpected preflight schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected preflight identity: %#v", preview)
	}
	if preview.InstallMode != "production" ||
		preview.PreflightDecision != "production-blocked" ||
		preview.ActivationState != "preflight-only" {
		t.Fatalf("unexpected production decision: %#v", preview)
	}
	if preview.Bundle.RequestType != "desktop-activation-bundle-preview" ||
		preview.Bundle.MaterialCount != 9 ||
		!preview.Bundle.StandardDesktopEntry ||
		!preview.Bundle.FileAssociationReady ||
		preview.Bundle.DesktopFilesWritten ||
		preview.Bundle.HostRootModified ||
		preview.Bundle.BackendDetailsExposed {
		t.Fatalf("unexpected bundle summary: %#v", preview.Bundle)
	}
	if preview.RecipeTrust.Source != "registry" ||
		preview.RecipeTrust.RegistryName != "test-registry" ||
		!preview.RecipeTrust.DigestVerified ||
		preview.RecipeTrust.SignatureStatus != "development-only" ||
		preview.RecipeTrust.TrustDecision != "development-only" ||
		preview.RecipeTrust.ProductionTrusted ||
		!preview.RecipeTrust.DevelopmentOnly ||
		preview.RecipeTrust.ExternalRecipeAccepted {
		t.Fatalf("unexpected recipe trust: %#v", preview.RecipeTrust)
	}
	if preview.InstallGate.GateType != "recipe-install" ||
		preview.InstallGate.Mode != "production" ||
		preview.InstallGate.Decision != "block" ||
		!containsString(preview.InstallGate.BlockingReasons, "production signed recipe validation is not enabled") ||
		!containsString(preview.InstallGate.BlockingReasons, "registry contains development-only recipes") {
		t.Fatalf("unexpected install gate: %#v", preview.InstallGate)
	}
	if preview.BackendBinding.RequestType != "backend-binding-preview" ||
		preview.BackendBinding.BindingState != "planned-blocked" ||
		preview.BackendBinding.RecommendedProfileID != "local-compatibility" ||
		preview.BackendBinding.RequiredReviewCount != 5 ||
		preview.BackendBinding.BindingCommitted ||
		preview.BackendBinding.BindingPersisted ||
		preview.BackendBinding.LaunchEnabled ||
		preview.BackendBinding.BackendDetailsExposed {
		t.Fatalf("unexpected backend binding summary: %#v", preview.BackendBinding)
	}
	if got, want := preview.CheckIDs, []string{"recipe-digest", "recipe-signature", "recipe-install-gate", "activation-materials", "backend-binding", "staging-root", "host-root-write-gate"}; !sameStrings(got, want) {
		t.Fatalf("CheckIDs = %#v, want %#v", got, want)
	}
	if preview.CheckCount != 7 ||
		preview.PassedCheckCount != 3 ||
		preview.PendingCheckCount != 2 ||
		preview.BlockedCheckCount != 2 {
		t.Fatalf("unexpected check counts: %#v", preview)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.UserVisible || !preview.NormalApplicationSurface ||
		!preview.DesktopActivationReady ||
		preview.DevelopmentStagingEligible ||
		preview.ProductionActivationEligible ||
		preview.InstallerMayProceed ||
		!preview.StagingRootRequired ||
		preview.HostRootAllowed {
		t.Fatalf("unexpected readiness flags: %#v", preview)
	}
	if preview.DesktopFilesWritten || preview.MIMEAppsWritten ||
		preview.ManifestWritten || preview.ReceiptWritten ||
		preview.SettingsPersisted || preview.NotificationsSent ||
		preview.TaskManagerEntryActive || preview.KWinRuleApplied ||
		preview.LiveTrayBridgeEnabled || preview.LaunchEnabled ||
		preview.BackendLaunchEnabled || preview.ExecutionStarted ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.PrivilegedContainerRequired || preview.BackendDetailsExposed ||
		preview.RawWindowsExecutableExposed || preview.CompatibilityStorageExposed {
		t.Fatalf("desktop activation preflight opened an unsafe gate: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "write desktop files from activation preflight preview") ||
		!containsString(preview.BlockedActions, "promote development registry to production activation") ||
		!containsString(preview.BlockedActions, "mutate host root during activation preflight") {
		t.Fatalf("missing blocked actions: %#v", preview.BlockedActions)
	}

	developmentPreview, err := plan.DesktopActivationPreflightPreview("development")
	if err != nil {
		t.Fatalf("development DesktopActivationPreflightPreview returned error: %v", err)
	}
	if developmentPreview.InstallMode != "development" ||
		developmentPreview.PreflightDecision != "development-staging-ready" ||
		developmentPreview.InstallGate.Decision != "allow" ||
		!developmentPreview.DevelopmentStagingEligible ||
		developmentPreview.ProductionActivationEligible ||
		!developmentPreview.InstallerMayProceed ||
		developmentPreview.DesktopFilesWritten ||
		developmentPreview.HostRootModified ||
		developmentPreview.LaunchEnabled {
		t.Fatalf("unexpected development staging preview: %#v", developmentPreview)
	}
	if developmentPreview.PassedCheckCount != 4 ||
		developmentPreview.PendingCheckCount != 2 ||
		developmentPreview.BlockedCheckCount != 1 {
		t.Fatalf("unexpected development check counts: %#v", developmentPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop activation preflight preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
