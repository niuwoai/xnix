package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMultiApplicationInstallQueuePreviewAggregatesInstallPlans(t *testing.T) {
	plans := []Plan{
		testMultiApplicationInstallPlan(t, "org.example.ledger", "Example Ledger", "automatic", "development-only"),
		testMultiApplicationInstallPlan(t, "org.example.viewer", "Example Viewer", "wine", "development-only"),
	}

	preview, err := NewMultiApplicationInstallQueuePreview(plans, "development")
	if err != nil {
		t.Fatalf("NewMultiApplicationInstallQueuePreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.multi_application_install_queue.v1" ||
		preview.RequestType != "multi-application-install-queue-preview" ||
		preview.QueueType != "review-only-multi-application-install-queue" ||
		preview.Source != "registry+compatibility-install-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetMultiApplicationInstallQueue" ||
		preview.ReadMethod != "GetMultiApplicationInstallQueuePreview" ||
		preview.Environment != "development" ||
		preview.ApplicationCount != 2 ||
		preview.Counts.Total != 2 ||
		preview.Counts.MissingEvidence != 2 ||
		preview.InstallReadyCount != 0 ||
		preview.BlockedApplicationCount != 0 ||
		preview.ReviewRequiredCount != 2 ||
		preview.QueueStatus != "missing-evidence" ||
		!preview.ReadyForReview {
		t.Fatalf("unexpected multi-application install queue schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.ReviewOnly ||
		preview.QueuePersisted ||
		preview.RequestObjectsCreated ||
		preview.ArtifactsStaged ||
		preview.ArtifactsDownloaded ||
		preview.NetworkRequestCreated ||
		preview.HostPackageManagerInvoked ||
		preview.DesktopActivationStarted ||
		preview.InstallStarted ||
		preview.BackendProcessStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.SettingsPersisted ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected multi-application install queue safety flags: %#v", preview)
	}
	if len(preview.Applications) != 2 ||
		preview.Applications[0].Position != 1 ||
		preview.Applications[0].ApplicationID != "org.example.ledger" ||
		preview.Applications[0].QueueState != "missing-evidence" ||
		preview.Applications[0].Priority != "medium" ||
		!preview.Applications[0].UserReviewRequired ||
		preview.Applications[0].RequestObjectCreated ||
		preview.Applications[0].ArtifactStaged ||
		preview.Applications[0].InstallStarted ||
		preview.Applications[0].BackendProcessStarted ||
		preview.Applications[0].LaunchEnabled ||
		preview.Applications[0].ExecutionStarted ||
		preview.Applications[0].HostRootModified ||
		preview.Applications[0].StateRootPathExposed ||
		preview.Applications[0].RawCommandExposed ||
		preview.Applications[0].BackendDetailsExposed {
		t.Fatalf("unexpected queue item: %#v", preview.Applications)
	}
	if !containsString(preview.SharedBlockingReasons, "artifact staging receipt is missing") ||
		!containsString(preview.SharedBlockingReasons, "required artifacts are not staged") {
		t.Fatalf("expected shared blocking reasons: %#v", preview.SharedBlockingReasons)
	}
	assertMultiApplicationInstallQueueSafe(t, preview)
}

func TestMultiApplicationInstallQueuePreviewRejectsEmptyQueue(t *testing.T) {
	if _, err := NewMultiApplicationInstallQueuePreview(nil, "development"); err == nil {
		t.Fatalf("NewMultiApplicationInstallQueuePreview must reject empty queues")
	}
}

func TestMultiApplicationInstallQueuePreviewRejectsBadMode(t *testing.T) {
	plans := []Plan{testMultiApplicationInstallPlan(t, "org.example.ledger", "Example Ledger", "automatic", "development-only")}
	if _, err := NewMultiApplicationInstallQueuePreview(plans, "unsafe"); err == nil {
		t.Fatalf("NewMultiApplicationInstallQueuePreview must reject invalid mode")
	}
}

func testMultiApplicationInstallPlan(t *testing.T, id string, name string, mode string, signature string) Plan {
	t.Helper()
	recipe := Recipe{
		ID:                  id,
		Name:                name,
		Icon:                "application-x-executable",
		Mode:                mode,
		SupportedExtensions: []string{".abc"},
	}
	plan, err := NewPlanWithProvenance(recipe, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: signature})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}

func assertMultiApplicationInstallQueueSafe(t *testing.T, preview MultiApplicationInstallQueuePreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, required := range []string{"request_objects_created", "artifacts_staged", "network_request_created", "host_package_manager_invoked", "backend_process_started", "host_root_modified", "raw_command_exposed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("multi-application install queue JSON must include %q: %s", required, string(encoded))
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("multi-application install queue exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
	if err := validateNoBackendTerms(preview, "multi-application install queue preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
