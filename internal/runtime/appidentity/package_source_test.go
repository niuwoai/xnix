package appidentity

import (
	"strings"
	"testing"
)

func TestPackageSourcePreviewPlansRuntimeOwnedSourceWithoutInstall(t *testing.T) {
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

	preview, err := plan.PackageSourcePreview()
	if err != nil {
		t.Fatalf("PackageSourcePreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.package_source.v1" ||
		preview.RequestType != "package-source-preview" ||
		preview.SourceType != "compatibility-package-source" ||
		preview.Source != "registry+go-runtime-package-source" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetCompatibilityPackageSource" ||
		preview.ReadMethod != "GetCompatibilityPackageSourcePreview" {
		t.Fatalf("unexpected package source schema: %#v", preview)
	}
	if preview.Application.ID != "org.example.ledger" ||
		preview.Application.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Application.RequestedMode != "automatic" ||
		preview.SelectedStrategy != "automatic-managed" ||
		preview.SourceSelectionState != "planned" {
		t.Fatalf("unexpected package source identity: %#v", preview)
	}
	if !sameStrings(preview.SourceChannelIDs, []string{
		"os-managed-compatibility-packages",
		"runtime-managed-toolcache",
		"isolated-environment-template-catalog",
	}) {
		t.Fatalf("unexpected source channels: %#v", preview.SourceChannelIDs)
	}
	if !sameStrings(preview.RequiredPreflightIDs, []string{
		"signed-source-verification",
		"source-policy-review",
		"runtime-cache-quota",
		"offline-fallback",
	}) {
		t.Fatalf("unexpected required preflight: %#v", preview.RequiredPreflightIDs)
	}
	if !preview.SourcePolicy.SignedSourceRequired ||
		!preview.SourcePolicy.RuntimeCacheRequired ||
		preview.SourcePolicy.DirectDesktopInstallAllow ||
		preview.SourcePolicy.UserVisibleBackendNames ||
		preview.SourcePolicy.HostPackageManagerInvoked {
		t.Fatalf("unexpected source policy: %#v", preview.SourcePolicy)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.PackageSourceReady ||
		preview.InstallEnabled ||
		preview.NetworkRequiredForPlanning ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.DesktopShellCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected package source safety flags: %#v", preview)
	}
	if len(preview.BlockedActions) != 4 {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	if err := validateNoBackendTerms(preview, "package source preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}

	isolated, err := NewPlan(Recipe{ID: "org.example.iso", Name: "Iso", Icon: "application-x-executable", Mode: "vm", SupportedExtensions: []string{".abc"}})
	if err != nil {
		t.Fatalf("NewPlan isolated returned error: %v", err)
	}
	isolatedPreview, err := isolated.PackageSourcePreview()
	if err != nil {
		t.Fatalf("isolated PackageSourcePreview returned error: %v", err)
	}
	if isolatedPreview.SelectedStrategy != "isolated-compatible-managed" {
		t.Fatalf("isolated recipe did not map to isolated strategy: %#v", isolatedPreview.SelectedStrategy)
	}
	if strings.Contains(strings.ToLower(isolatedPreview.DesktopSafeSummary), "proton") {
		t.Fatalf("summary must stay backend-safe: %s", isolatedPreview.DesktopSafeSummary)
	}
}
