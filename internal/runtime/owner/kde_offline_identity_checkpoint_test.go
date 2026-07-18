package owner

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEOfflineIdentityCheckpointClosesNineSurfaceBand(t *testing.T) {
	checkpoint, err := NewKDEOfflineIdentityCheckpoint(projectRoot(t))
	if err != nil {
		t.Fatalf("NewKDEOfflineIdentityCheckpoint: %v", err)
	}
	if checkpoint.Version != currentProjectVersion(t) ||
		checkpoint.SchemaVersion != "xnix.runtime.kde_offline_identity_checkpoint.v1" ||
		checkpoint.RequestType != "kde-offline-identity-checkpoint" ||
		checkpoint.CheckpointType != "offline-kde-application-identity-band-checkpoint" ||
		checkpoint.ApplicationID != "org.xnix.sample.notepad" || checkpoint.DisplayName != "Sample Notepad" ||
		checkpoint.SurfaceCount != 9 || checkpoint.ExpectedSurfaceCount != 9 || len(checkpoint.SurfaceIDs) != 9 ||
		!checkpoint.RecipeDigestVerified || checkpoint.RecipeSignatureStatus != "development-only" || checkpoint.ProductionSignatureReady ||
		!checkpoint.CrossSurfaceIdentityConsistent || !checkpoint.OwnerRouteReady || !checkpoint.OwnerServiceCallReady ||
		checkpoint.FormalReadRouteCount != 61 || checkpoint.OwnerReadMethodCount != 71 ||
		checkpoint.OwnerLocalReadMethodCount != 10 || checkpoint.SmokeReadRecordCount != 71 ||
		checkpoint.SmokeWriteDenialCount != 4 || !checkpoint.MethodParityReady || !checkpoint.RouteBandReady ||
		!checkpoint.OfflineIdentityReady {
		t.Fatalf("unexpected offline KDE identity checkpoint: %#v", checkpoint)
	}
	if len(checkpoint.Checks) != 6 {
		t.Fatalf("checkpoint checks = %d, want 6", len(checkpoint.Checks))
	}
	for _, check := range checkpoint.Checks {
		if check.Status != "pass" {
			t.Fatalf("checkpoint check did not pass: %#v", check)
		}
	}
	if !checkpoint.RuntimeOwned || !checkpoint.GoRuntimeBacked || checkpoint.KDEPolicyOwner || !checkpoint.ReviewOnly ||
		checkpoint.DesktopFilesWritten || checkpoint.MIMEDefaultsWritten || checkpoint.KRunnerIndexPersisted ||
		checkpoint.TaskManagerEntryActive || checkpoint.KWinRuleApplied || checkpoint.LiveTrayBridgeEnabled ||
		checkpoint.NotificationSent || checkpoint.SettingsPersisted || checkpoint.CompatibilityCenterPersisted ||
		checkpoint.LaunchEnabled || checkpoint.ExecutionStarted || checkpoint.BackendProcessStarted ||
		checkpoint.ProductionBusClaimed || checkpoint.WriteMethodsEnabled || checkpoint.NetworkRequired ||
		checkpoint.HostRootModified || checkpoint.PrivilegedContainerRequired || checkpoint.RawCommandExposed ||
		checkpoint.BackendDetailsExposed {
		t.Fatalf("offline KDE identity checkpoint opened an unsafe gate: %#v", checkpoint)
	}
	if err := checkpoint.Validate(); err != nil {
		t.Fatalf("Validate: %v", err)
	}
	encoded, err := json.Marshal(checkpoint)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"/users", "/home", "/private", "/tmp", "file://", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("checkpoint exposed forbidden term %q: %s", forbidden, text)
		}
	}
}
