package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestApplicationUpgradeImpactPreviewClassifiesNewerCandidate(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.0.0", "automatic", ".abc")
	candidate := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.1.0", "automatic", ".abc")

	preview, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), candidate, testApplicationUpgradeProvenance("signed"), "development")
	if err != nil {
		t.Fatalf("NewApplicationUpgradeImpactPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.application_upgrade_impact.v1" ||
		preview.RequestType != "application-upgrade-impact-preview" ||
		preview.ImpactType != "review-only-application-upgrade-impact" ||
		preview.Source != "registry+compatibility-install-preview+desktop-activation-evidence" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetApplicationUpgradeImpact" ||
		preview.ReadMethod != "GetApplicationUpgradeImpactPreview" ||
		preview.Environment != "development" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.CandidateApplicationID != "org.example.ledger" ||
		preview.CurrentVersion != "1.0.0" ||
		preview.CandidateVersion != "1.1.0" ||
		preview.CandidateRelation != "candidate-newer" ||
		preview.SectionCount != 8 ||
		preview.Counts.Total != 8 ||
		preview.Counts.Changed == 0 ||
		preview.Counts.MissingEvidence == 0 ||
		preview.OverallState != "missing-evidence" ||
		!preview.ReadyForReview {
		t.Fatalf("unexpected application upgrade impact preview: %#v", preview)
	}
	if !sameStrings(preview.SectionIDs, []string{
		"recipe-metadata",
		"package-source",
		"artifact-digests",
		"backend-profile",
		"portal-permissions",
		"snapshot-requirements",
		"desktop-activation",
		"diagnostics",
	}) {
		t.Fatalf("unexpected section ids: %#v", preview.SectionIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.ReviewOnly ||
		preview.RecipeWritten ||
		preview.RegistryMigrated ||
		preview.RequestObjectsCreated ||
		preview.ArtifactsStaged ||
		preview.ArtifactsDownloaded ||
		preview.NetworkRequestCreated ||
		preview.HostPackageManagerInvoked ||
		preview.SettingsPersisted ||
		preview.DesktopActivationStarted ||
		preview.DesktopFilesWritten ||
		preview.BackendProcessStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.RawBackendImplementationShown {
		t.Fatalf("unexpected safety flags: %#v", preview)
	}
	metadata := preview.Sections[0]
	if metadata.ID != "recipe-metadata" ||
		metadata.State != "changed" ||
		!metadata.Changed ||
		!metadata.UserReviewRequired ||
		metadata.RecipeWritten ||
		metadata.RegistryMigrated ||
		metadata.ArtifactsStaged ||
		metadata.SettingsPersisted ||
		metadata.DesktopFilesWritten ||
		metadata.BackendProcessStarted ||
		metadata.LaunchEnabled ||
		metadata.ExecutionStarted ||
		metadata.HostRootModified ||
		metadata.StateRootPathExposed ||
		metadata.RawCommandExposed ||
		metadata.BackendDetailsExposed {
		t.Fatalf("unexpected recipe metadata section: %#v", metadata)
	}
	if !containsString(preview.BlockedActions, "write candidate recipe from upgrade preview") ||
		!containsString(preview.BlockedActions, "migrate recipe registry from upgrade preview") ||
		!containsString(preview.BlockedActions, "launch application from upgrade preview") {
		t.Fatalf("expected blocked actions: %#v", preview.BlockedActions)
	}
	assertApplicationUpgradeImpactSafe(t, preview)
}

func TestApplicationUpgradeImpactPreviewClassifiesSameAndOlderCandidates(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "2.0.0", "automatic", ".abc")
	same := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "2.0.0", "automatic", ".abc")
	older := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.9.0", "automatic", ".abc")

	samePreview, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), same, testApplicationUpgradeProvenance("signed"), "development")
	if err != nil {
		t.Fatalf("same candidate preview returned error: %v", err)
	}
	if samePreview.CandidateRelation != "same-version" || samePreview.Sections[0].State != "unchanged" {
		t.Fatalf("unexpected same-version preview: %#v", samePreview)
	}

	olderPreview, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), older, testApplicationUpgradeProvenance("signed"), "development")
	if err != nil {
		t.Fatalf("older candidate preview returned error: %v", err)
	}
	if olderPreview.CandidateRelation != "candidate-older" ||
		olderPreview.Sections[0].State != "needs-review" ||
		olderPreview.Counts.NeedsReview == 0 {
		t.Fatalf("unexpected older candidate preview: %#v", olderPreview)
	}
	assertApplicationUpgradeImpactSafe(t, olderPreview)
}

func TestApplicationUpgradeImpactPreviewBlocksUntrustedProductionCandidate(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.0.0", "automatic", ".abc")
	candidate := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.1.0", "automatic", ".abc")

	preview, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), candidate, testApplicationUpgradeProvenance("development-only"), "production")
	if err != nil {
		t.Fatalf("NewApplicationUpgradeImpactPreview returned error: %v", err)
	}

	if preview.OverallState != "blocked" ||
		preview.Counts.Blocked == 0 ||
		preview.Sections[0].State != "blocked" ||
		!containsString(preview.Sections[0].BlockingReasons, "production signed recipe validation is not enabled") ||
		!containsString(preview.Sections[0].BlockingReasons, "registry contains development-only recipes") {
		t.Fatalf("expected trust blocker: %#v", preview)
	}
	assertApplicationUpgradeImpactSafe(t, preview)
}

func TestApplicationUpgradeImpactPreviewRejectsMalformedCandidate(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.0.0", "automatic", ".abc")
	candidate := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "bad-version", "automatic", ".abc")

	if _, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), candidate, testApplicationUpgradeProvenance("signed"), "development"); err == nil {
		t.Fatalf("NewApplicationUpgradeImpactPreview accepted a malformed candidate recipe")
	}
}

func TestApplicationUpgradeImpactPreviewBlocksCandidateIdentityDrift(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.0.0", "automatic", ".abc")
	candidate := testApplicationUpgradeRecipe("org.example.viewer", "Example Viewer", "1.1.0", "automatic", ".abc")

	preview, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), candidate, testApplicationUpgradeProvenance("signed"), "development")
	if err != nil {
		t.Fatalf("NewApplicationUpgradeImpactPreview returned error: %v", err)
	}
	if preview.OverallState != "blocked" ||
		preview.Sections[0].State != "blocked" ||
		!containsString(preview.Sections[0].BlockingReasons, "candidate application identity does not match the current application") {
		t.Fatalf("expected identity drift blocker: %#v", preview)
	}
	assertApplicationUpgradeImpactSafe(t, preview)
}

func TestApplicationUpgradeImpactPreviewRejectsBadMode(t *testing.T) {
	current := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.0.0", "automatic", ".abc")
	candidate := testApplicationUpgradeRecipe("org.example.ledger", "Example Ledger", "1.1.0", "automatic", ".abc")

	if _, err := NewApplicationUpgradeImpactPreview(current, testApplicationUpgradeProvenance("signed"), candidate, testApplicationUpgradeProvenance("signed"), "unsafe"); err == nil {
		t.Fatalf("NewApplicationUpgradeImpactPreview accepted an unsafe environment")
	}
}

func testApplicationUpgradeRecipe(id string, name string, version string, mode string, extension string) Recipe {
	return Recipe{
		ID:                  id,
		Name:                name,
		Version:             version,
		Icon:                "application-x-executable",
		Mode:                mode,
		SupportedExtensions: []string{extension},
	}
}

func testApplicationUpgradeProvenance(signature string) Provenance {
	return Provenance{
		Source:          "registry",
		RegistryName:    "test-registry",
		DigestVerified:  true,
		SignatureStatus: signature,
	}
}

func assertApplicationUpgradeImpactSafe(t *testing.T, preview ApplicationUpgradeImpactPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, required := range []string{"recipe_written", "registry_migrated", "artifacts_staged", "network_request_created", "host_package_manager_invoked", "desktop_activation_started", "backend_process_started", "host_root_modified", "raw_command_exposed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("application upgrade impact JSON must include %q: %s", required, string(encoded))
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("application upgrade impact exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
	if err := validateNoBackendTerms(preview, "application upgrade impact preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
