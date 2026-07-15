package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEIntegrationStatusPreviewCoversSevenRuntimeBackedEntryPoints(t *testing.T) {
	preview, err := NewKDEIntegrationStatusPreview()
	if err != nil {
		t.Fatalf("NewKDEIntegrationStatusPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_integration_status.v1" ||
		preview.RequestType != "kde-integration-status-preview" ||
		preview.StatusType != "kde-integration-status" ||
		preview.Source != "go-runtime-kde-integration-status" ||
		preview.RuntimeMethod != "GetKDEIntegrationStatus" ||
		preview.ReadMethod != "GetKDEIntegrationStatusPreview" ||
		preview.Desktop != "KDE Plasma" {
		t.Fatalf("unexpected KDE integration status schema: %#v", preview)
	}
	if preview.EntryPointCount != 7 ||
		preview.InitialCount != 6 ||
		preview.PlannedCount != 1 ||
		preview.CompleteCount != 0 ||
		strings.Join(preview.EntryPointIDs, ",") != "launcher,task-manager,file-manager,system-tray,notifications,compatibility-center,settings" {
		t.Fatalf("unexpected KDE integration counts: %#v ids=%#v", preview, preview.EntryPointIDs)
	}
	for _, entry := range preview.EntryPoints {
		if !entry.RuntimeBacked ||
			!entry.GoRuntimeBacked ||
			!entry.DBusReadAvailable ||
			entry.KDEPolicyOwner ||
			entry.RuntimeMethod == "" ||
			entry.AdapterRole == "" {
			t.Fatalf("unexpected entry point safety flags for %s: %#v", entry.ID, entry)
		}
	}
	if len(preview.Checks) != 3 ||
		strings.Join(preview.CheckIDs, ",") != "official-desktop,runtime-policy-owner,desktop-safety" {
		t.Fatalf("unexpected checks: %#v ids=%#v", preview.Checks, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly ||
		!preview.StableDesktopContract ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE integration safety flags: %#v", preview)
	}
	assertNoKDEPreviewBackendTerms(t, preview)
}

func TestKDEShellIntegrationPlanPreviewKeepsShellReplaceable(t *testing.T) {
	preview, err := NewKDEShellIntegrationPlanPreview()
	if err != nil {
		t.Fatalf("NewKDEShellIntegrationPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_shell_integration.v1" ||
		preview.RequestType != "kde-shell-integration-preview" ||
		preview.PlanType != "kde-shell-integration-plan" ||
		preview.Source != "go-runtime-kde-shell-integration" ||
		preview.RuntimeMethod != "GetKDEShellIntegrationPlan" ||
		preview.ReadMethod != "GetKDEShellIntegrationPlanPreview" ||
		preview.DesktopShell != "KDE Plasma" {
		t.Fatalf("unexpected KDE shell integration schema: %#v", preview)
	}
	if preview.ComponentCount != 9 ||
		preview.InitialComponentCount != 8 ||
		preview.PlannedComponentCount != 1 ||
		!containsString(preview.ComponentIDs, "krunner-search") ||
		!containsString(preview.ComponentIDs, "kwin-window-management") {
		t.Fatalf("unexpected KDE shell components: %#v ids=%#v", preview.Components, preview.ComponentIDs)
	}
	for _, component := range preview.Components {
		if !component.RuntimeOwned ||
			!component.GoRuntimeBacked ||
			component.KDEPolicyOwner ||
			component.ShellWritesEnabled ||
			component.BackendLaunchEnabled ||
			component.BackendDetailsExposed ||
			component.RuntimeMethod == "" {
			t.Fatalf("unexpected component safety flags for %s: %#v", component.ID, component)
		}
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly ||
		preview.FallbackDesktopsSupported ||
		preview.PlasmaForkRequired ||
		preview.PlasmaSourceModified ||
		preview.ShellConfigurationWritten ||
		preview.ComponentActivationEnabled ||
		preview.BackendLaunchEnabled ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE shell integration safety flags: %#v", preview)
	}
	assertNoKDEPreviewBackendTerms(t, preview)
}

func TestKDEApplicationSurfacePlanPreviewKeepsApplicationNormalAndBlocked(t *testing.T) {
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

	preview, err := plan.KDEApplicationSurfacePlanPreview()
	if err != nil {
		t.Fatalf("KDEApplicationSurfacePlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_application_surface.v1" ||
		preview.RequestType != "kde-application-surface-preview" ||
		preview.PlanType != "kde-application-surface-plan" ||
		preview.Source != "go-runtime-kde-application-surface" ||
		preview.RuntimeMethod != "GetKDEApplicationSurfacePlan" ||
		preview.ReadMethod != "GetKDEApplicationSurfacePlanPreview" ||
		preview.SurfaceState != "planned" ||
		preview.DesktopShell != "KDE Plasma" {
		t.Fatalf("unexpected KDE application surface schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.Icon != "office-chart-area" ||
		preview.Application.RequestedMode != "automatic" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		len(preview.Application.SupportedExtensions) != 1 ||
		preview.Application.SupportedExtensions[0] != "application/x-xnix-xls" {
		t.Fatalf("unexpected application identity: %#v", preview.Application)
	}
	if preview.EntryPointCount != 7 ||
		strings.Join(preview.EntryPointIDs, ",") != "launcher,task-manager,file-manager,system-tray,notifications,compatibility-center,settings" ||
		len(preview.RequiredRuntimeGates) != 6 ||
		!containsString(preview.RequiredRuntimeGates, "runtime-write-gate") {
		t.Fatalf("unexpected application surface entry points or gates: %#v gates=%#v", preview.EntryPointIDs, preview.RequiredRuntimeGates)
	}
	for _, entry := range preview.EntryPoints {
		if !entry.RuntimeBacked ||
			!entry.GoRuntimeBacked ||
			entry.KDEWritesPolicy ||
			entry.BackendDetailsExposed ||
			entry.State != "planned" {
			t.Fatalf("unexpected application surface entry for %s: %#v", entry.ID, entry)
		}
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly ||
		!preview.NormalLinuxApplicationSurface ||
		!preview.StandardLauncherVisible ||
		!preview.TaskManagerIdentityReady ||
		!preview.FileAssociationsPlanned ||
		!preview.DolphinActionPlanned ||
		!preview.KRunnerQueryPlanned ||
		!preview.TrayStatusPlanned ||
		!preview.NotificationRoutePlanned ||
		!preview.SettingsSurfacePlanned ||
		!preview.PortalReviewRequired ||
		preview.ExecutionReady ||
		preview.LaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.DesktopFilesWritten ||
		preview.MIMEAppsWritten ||
		preview.HostRootModified ||
		preview.BackendCommandExposed ||
		preview.RawWindowsExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KDE application surface safety flags: %#v", preview)
	}
	assertNoKDEPreviewBackendTerms(t, preview)
}

func assertNoKDEPreviewBackendTerms(t *testing.T, value any) {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE shell preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
