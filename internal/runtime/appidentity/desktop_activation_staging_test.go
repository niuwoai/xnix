package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDesktopActivationStagingPreviewPlansFilesWithoutWriting(t *testing.T) {
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

	preview, err := plan.DesktopActivationStagingPreview("development")
	if err != nil {
		t.Fatalf("DesktopActivationStagingPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.desktop_activation_staging.v1" ||
		preview.RequestType != "desktop-activation-staging-preview" ||
		preview.StagingType != "kde-activation-staging-plan" ||
		preview.RuntimeMethod != "GetDesktopActivationStaging" {
		t.Fatalf("unexpected staging schema: %#v", preview)
	}
	if preview.Desktop != "KDE Plasma" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" {
		t.Fatalf("unexpected staging identity: %#v", preview)
	}
	if preview.InstallMode != "development" ||
		preview.PreflightDecision != "development-staging-ready" ||
		preview.StagingState != "staging-plan-ready" ||
		preview.Preflight.RequestType != "desktop-activation-preflight-preview" ||
		preview.Preflight.InstallGateDecision != "allow" ||
		!preview.Preflight.DevelopmentStagingEligible ||
		preview.Preflight.HostRootAllowed {
		t.Fatalf("unexpected preflight summary: %#v", preview)
	}
	if got, want := preview.PlannedFileIDs, []string{"desktop-entry", "dolphin-service-menu", "mimeapps-list", "desktop-integration-manifest", "managed-launcher-artifact", "desktop-activation-receipt"}; !sameStrings(got, want) {
		t.Fatalf("PlannedFileIDs = %#v, want %#v", got, want)
	}
	if preview.PlannedFileCount != 6 || len(preview.PlannedFiles) != 6 ||
		preview.ReceiptFileID != "desktop-activation-receipt" {
		t.Fatalf("unexpected planned file counts: %#v", preview)
	}
	expectedPaths := []string{
		"usr/share/applications/xnix-org.example.ledger.desktop",
		"usr/share/kio/servicemenus/xnix-open-with-compatibility.desktop",
		"usr/share/applications/mimeapps.list",
		"usr/share/xnix/compatibility/manifests/org.example.ledger.json",
		"usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json",
		"usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json",
	}
	for index, file := range preview.PlannedFiles {
		if file.RelativePath != expectedPaths[index] ||
			file.Mode != "0644" ||
			len(file.SHA256) != 64 ||
			!file.PlannedForStaging ||
			file.Written ||
			file.HostRootModified ||
			file.BackendDetailsExposed {
			t.Fatalf("unexpected planned file at %d: %#v", index, file)
		}
	}
	if got, want := preview.ActivatedEntryPoints, []string{"launcher", "task-manager", "file-manager", "system-tray", "notifications", "compatibility-center", "unified-settings"}; !sameStrings(got, want) {
		t.Fatalf("ActivatedEntryPoints = %#v, want %#v", got, want)
	}
	if preview.ActivatedEntryPointCount != 7 {
		t.Fatalf("unexpected entry point count: %#v", preview)
	}
	if got, want := preview.InstallerCommandPreview, []string{"xnix-install-desktop-integration", "--desktop-entry-source", "runtime-go", "--file-association-source", "runtime-go"}; !sameStrings(got, want) {
		t.Fatalf("InstallerCommandPreview = %#v, want %#v", got, want)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.DevelopmentStagingEligible ||
		preview.ProductionActivationEligible ||
		!preview.InstallerMayProceed ||
		!preview.StagingPlanReady ||
		!preview.StagingRootRequired ||
		preview.StagingRootPathExposed ||
		preview.HostRootAllowed {
		t.Fatalf("unexpected staging readiness flags: %#v", preview)
	}
	if preview.FileWritesPerformed || preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten || preview.ManifestWritten ||
		preview.ReceiptWritten || !preview.RollbackReceiptPlanned ||
		preview.RollbackReceiptWritten || preview.SettingsPersisted ||
		preview.NotificationsSent || preview.TaskManagerEntryActive ||
		preview.KWinRuleApplied || preview.LiveTrayBridgeEnabled ||
		preview.LaunchEnabled || preview.BackendLaunchEnabled ||
		preview.ExecutionStarted || preview.HostRootModified ||
		preview.NetworkRequired || preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed || preview.RawWindowsExecutableExposed ||
		preview.CompatibilityStorageExposed {
		t.Fatalf("desktop activation staging opened an unsafe gate: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "write staged files from staging preview") ||
		!containsString(preview.BlockedActions, "install into host root from staging preview") ||
		!containsString(preview.BlockedActions, "mutate host root during activation staging preview") {
		t.Fatalf("missing blocked actions: %#v", preview.BlockedActions)
	}

	productionPreview, err := plan.DesktopActivationStagingPreview("production")
	if err != nil {
		t.Fatalf("production DesktopActivationStagingPreview returned error: %v", err)
	}
	if productionPreview.StagingState != "staging-blocked" ||
		productionPreview.InstallerMayProceed ||
		productionPreview.StagingPlanReady ||
		productionPreview.DevelopmentStagingEligible ||
		productionPreview.ProductionActivationEligible {
		t.Fatalf("production staging should remain blocked: %#v", productionPreview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine", "/home", "/users", "/var", "/opt", "/tmp"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("desktop activation staging preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
