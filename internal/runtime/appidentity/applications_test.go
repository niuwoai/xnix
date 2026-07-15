package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestApplicationsPreviewListsDesktopSafeApplications(t *testing.T) {
	recipes := []Recipe{
		{
			ID:                  "org.example.ledger",
			Name:                "Example Ledger",
			Icon:                "office-chart-area",
			Mode:                "automatic",
			SupportedExtensions: []string{".abc", ".xls"},
		},
	}
	preview, err := NewApplicationsPreview(recipes, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewApplicationsPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.applications.v1" ||
		preview.RequestType != "applications-preview" ||
		preview.CatalogType != "runtime-application-catalog" ||
		preview.Source != "registry+go-runtime-desktop-identity" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "ListApplications" ||
		preview.ReadMethod != "ListApplicationsPreview" {
		t.Fatalf("unexpected applications schema: %#v", preview)
	}
	if preview.RegistryName != "test-registry" ||
		!preview.RecipeDigestVerified ||
		preview.RecipeSignatureStatus != "development-only" ||
		preview.ApplicationCount != 1 ||
		!sameStrings(preview.ApplicationIDs, []string{"org.example.ledger"}) ||
		!sameStrings(preview.EntryPointIDs, []string{"start-menu", "task-manager", "file-manager", "system-tray", "notification-center", "compatibility-center", "unified-settings"}) {
		t.Fatalf("unexpected catalog summary: %#v", preview)
	}
	application := preview.Applications[0]
	if application.ID != "org.example.ledger" ||
		application.Name != "Example Ledger" ||
		application.RuntimeMode != "Automatic" ||
		application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		application.LauncherAction != "runtime-launch" ||
		!sameStrings(application.SupportedExtensions, []string{".abc", ".xls"}) ||
		!sameStrings(application.MIMETypes, []string{"application/x-xnix-abc", "application/x-xnix-xls"}) ||
		!sameStrings(application.LaunchCommand, []string{"xnix-compat-launch", "--app", "org.example.ledger", "%U"}) ||
		!application.UserVisible ||
		!application.StandardDesktopEntry ||
		!application.AcceptsFileURIs ||
		!application.RuntimeOwned ||
		!application.BackendTerminologyHidden ||
		application.BackendLaunchEnabled ||
		application.BackendDetailsExposed ||
		application.RawWindowsExecutableExposed ||
		application.HostRootMutationEnabled {
		t.Fatalf("unexpected application entry: %#v", application)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.StandardDesktopEntries ||
		!preview.RecipeRegistryVerified ||
		!preview.BackendTerminologyHidden ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.DesktopFilesWritten ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed ||
		preview.RawWindowsExecutableExposed {
		t.Fatalf("unexpected catalog safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 5 ||
		preview.BlockedActions[0] != "launch application from catalog preview" ||
		preview.DesktopSafeSummary != "Application catalog preview is Go-owned and presents registered compatibility applications as normal Linux applications while launch and host mutation remain disabled." {
		t.Fatalf("unexpected catalog metadata: actions=%#v summary=%q", preview.BlockedActions, preview.DesktopSafeSummary)
	}
	assertNoForbiddenApplicationTerms(t, preview)
}

func TestApplicationPreviewRendersOneDesktopSafeApplication(t *testing.T) {
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	}
	preview, err := NewApplicationPreview(recipe, Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewApplicationPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.application.v1" ||
		preview.RequestType != "application-preview" ||
		preview.CatalogType != "runtime-application" ||
		preview.Source != "registry+go-runtime-desktop-identity" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetApplication" ||
		preview.ReadMethod != "GetApplicationPreview" {
		t.Fatalf("unexpected application schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Application.RegistryName != "test-registry" ||
		!preview.Application.RecipeDigestVerified ||
		preview.Application.RecipeSignatureStatus != "development-only" {
		t.Fatalf("unexpected application payload: %#v", preview.Application)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.StandardDesktopEntry ||
		!preview.RecipeRegistryVerified ||
		!preview.BackendTerminologyHidden ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.DesktopFileWritten ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed ||
		preview.RawWindowsExecutableExposed {
		t.Fatalf("unexpected application safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 5 ||
		preview.BlockedActions[0] != "launch application from application preview" ||
		preview.DesktopSafeSummary != "Application preview is Go-owned and presents one registered compatibility application as a normal Linux application while launch and host mutation remain disabled." {
		t.Fatalf("unexpected application metadata: actions=%#v summary=%q", preview.BlockedActions, preview.DesktopSafeSummary)
	}
	assertNoForbiddenApplicationTerms(t, preview)
}

func assertNoForbiddenApplicationTerms(t *testing.T, value any) {
	t.Helper()
	if err := validateNoBackendTerms(value, "application preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	serialized := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine "} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("application preview exposes forbidden term %q: %s", forbidden, serialized)
		}
	}
}
