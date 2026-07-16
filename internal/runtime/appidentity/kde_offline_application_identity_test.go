package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEOfflineApplicationIdentityPreviewJoinsCanonicalKDESurfaces(t *testing.T) {
	preview, err := NewKDEOfflineApplicationIdentityPreview(Recipe{
		ID:                  "org.xnix.sample.notepad",
		Name:                "Sample Notepad",
		Icon:                "accessories-text-editor",
		Mode:                "automatic",
		SupportedExtensions: []string{".txt", ".log"},
	}, Provenance{
		Source:          "registry",
		RegistryName:    "xnix-local-development",
		DigestVerified:  true,
		SignatureStatus: "development-only",
	})
	if err != nil {
		t.Fatalf("NewKDEOfflineApplicationIdentityPreview: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.kde_offline_application_identity.v1" ||
		preview.RequestType != "kde-offline-application-identity-preview" ||
		preview.IdentityType != "digest-verified-offline-kde-application" ||
		preview.RuntimeMethod != "GetKDEOfflineApplicationIdentity" ||
		preview.ReadMethod != "GetKDEOfflineApplicationIdentityPreview" ||
		preview.ApplicationID != "org.xnix.sample.notepad" ||
		preview.DisplayName != "Sample Notepad" ||
		preview.Icon != "accessories-text-editor" ||
		preview.DesktopFile != "xnix-org.xnix.sample.notepad.desktop" ||
		preview.LauncherURL != "applications:xnix-org.xnix.sample.notepad.desktop" ||
		preview.SurfaceCount != 5 || len(preview.SurfaceIDs) != 5 ||
		!preview.RecipeDigestVerified || !preview.CrossSurfaceIdentityConsistent ||
		!preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.ReviewOnly || !preview.Offline {
		t.Fatalf("unexpected offline KDE identity: %#v", preview)
	}
	if len(preview.DesktopEntry.ContentSHA256) != 64 || !preview.DesktopEntry.IdentityFieldsMatch ||
		preview.DesktopEntry.MIMETypeCount != 2 || preview.DesktopEntry.WriteEnabled ||
		len(preview.MIMEAssociations.ContentSHA256) != 64 || !preview.MIMEAssociations.IdentityFieldsMatch ||
		preview.MIMEAssociations.DefaultCount != 2 || preview.MIMEAssociations.AssociationCount != 2 ||
		preview.MIMEAssociations.WriteEnabled {
		t.Fatalf("unexpected rendered identity evidence: %#v %#v", preview.DesktopEntry, preview.MIMEAssociations)
	}
	if !preview.KRunner.IdentityFieldsMatch || preview.KRunner.MatchCount != 1 ||
		preview.KRunner.DesktopEntryID != preview.DesktopFile || preview.KRunner.QueryExecutionEnabled ||
		!preview.TaskManager.IdentityFieldsMatch || preview.TaskManager.GroupingKey != preview.ApplicationID || preview.TaskManager.EntryActive ||
		!preview.KWin.IdentityFieldsMatch || preview.KWin.ResourceName != preview.ApplicationID || preview.KWin.RuleApplied {
		t.Fatalf("unexpected KDE surface evidence: %#v %#v %#v", preview.KRunner, preview.TaskManager, preview.KWin)
	}
	if preview.DesktopFilesWritten || preview.MIMEDefaultsWritten || preview.KRunnerIndexPersisted ||
		preview.TaskManagerEntryActive || preview.KWinRuleApplied || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.BackendProcessStarted || preview.NetworkRequired ||
		preview.HostRootModified || preview.RawCommandExposed || preview.BackendDetailsExposed {
		t.Fatalf("offline KDE identity opened an unsafe gate: %#v", preview)
	}
	assertKDEOfflineApplicationIdentitySafe(t, preview)
}

func TestKDEOfflineApplicationIdentityPreviewRequiresVerifiedRegistryRecipe(t *testing.T) {
	recipe := Recipe{ID: "org.example.notes", Name: "Notes", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	for _, provenance := range []Provenance{
		{Source: "direct-file"},
		{Source: "registry", RegistryName: "test-registry", DigestVerified: false},
	} {
		if _, err := NewKDEOfflineApplicationIdentityPreview(recipe, provenance); err == nil {
			t.Fatalf("accepted unverified provenance: %#v", provenance)
		}
	}
}

func assertKDEOfflineApplicationIdentitySafe(t *testing.T, value any) {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("offline KDE identity exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
