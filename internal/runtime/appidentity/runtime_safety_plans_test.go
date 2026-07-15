package appidentity

import "testing"

func TestApplicationStateRootPreviewKeepsStoragePlanned(t *testing.T) {
	plan := testRuntimeSafetyPlan(t)

	preview, err := plan.ApplicationStateRootPreview()
	if err != nil {
		t.Fatalf("ApplicationStateRootPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.state_root.v1" ||
		preview.RequestType != "state-root-preview" ||
		preview.RootType != "compatibility-application-state-root" ||
		preview.RuntimeMethod != "GetApplicationStateRoot" ||
		preview.ReadMethod != "GetApplicationStateRootPreview" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.StateNamespace != "org.example.ledger" {
		t.Fatalf("unexpected state root schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DirectoriesCreated ||
		preview.HostRootModified ||
		preview.UserDocumentsIncluded ||
		!preview.PortalRequiredForUserFiles ||
		!preview.SnapshotEligible ||
		!preview.RestoreRequiresConfirm ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected state root safety flags: %#v", preview)
	}
	if len(preview.ManagedScopes) != 4 ||
		!sameStrings(preview.ManagedScopeIDs, []string{"application-data", "runtime-metadata", "diagnostic-cache", "desktop-activation-receipts"}) {
		t.Fatalf("unexpected state root scopes: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "application state root preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestSnapshotPlanPreviewKeepsSnapshotAndRestoreClosed(t *testing.T) {
	preview, err := NewSnapshotPlanPreview("org.example.ledger", "before-repair")
	if err != nil {
		t.Fatalf("NewSnapshotPlanPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.snapshot_plan.v1" ||
		preview.RequestType != "snapshot-plan-preview" ||
		preview.PlanType != "compatibility-snapshot" ||
		preview.RuntimeMethod != "GetSnapshotPlan" ||
		preview.ReadMethod != "GetSnapshotPlanPreview" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.Reason != "before-repair" {
		t.Fatalf("unexpected snapshot plan schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.EnabledByDefault ||
		!preview.SnapshotScope.ApplicationState ||
		preview.SnapshotScope.UserDocuments ||
		preview.SnapshotScope.HostSystem ||
		!preview.Restore.Available ||
		!preview.Restore.RequiresUserConfirmation ||
		!preview.Restore.PreserveUserDocuments ||
		preview.SnapshotRequestCreated ||
		preview.SnapshotCreated ||
		preview.RestoreRequested ||
		preview.RestoreExecuted ||
		preview.UserDocumentsIncluded ||
		preview.HostSystemIncluded ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected snapshot safety flags: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "snapshot plan preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestPortalAccessPolicyPreviewRequiresDesktopMediation(t *testing.T) {
	preview, err := NewPortalAccessPolicyPreview("org.example.ledger", "file-open")
	if err != nil {
		t.Fatalf("NewPortalAccessPolicyPreview returned error: %v", err)
	}

	if preview.SchemaVersion != "xnix.runtime.portal_access_policy.v1" ||
		preview.RequestType != "portal-access-policy-preview" ||
		preview.PolicyType != "portal-access" ||
		preview.RuntimeMethod != "GetPortalAccessPolicy" ||
		preview.ReadMethod != "GetPortalAccessPolicyPreview" ||
		preview.ApplicationID != "org.example.ledger" ||
		preview.Operation != "file-open" ||
		preview.Decision != "ask" {
		t.Fatalf("unexpected Portal policy schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.PortalRequired ||
		!preview.UserMediationRequired ||
		preview.DirectAccessAllowed ||
		!preview.RequestFlow.RequestObjectRequired ||
		!preview.RequestFlow.RuntimePolicyOwner ||
		preview.RequestFlow.DesktopShellPolicyOwner ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.HostPermissionChanged ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected Portal policy safety flags: %#v", preview)
	}
	if !containsString(preview.Resources, "selected-files") {
		t.Fatalf("Portal file-open policy must include selected-files scope: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "Portal access policy preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestPortalAccessPolicyPreviewDeniesSensitiveDevicesByDefault(t *testing.T) {
	camera, err := NewPortalAccessPolicyPreview("org.example.ledger", "camera")
	if err != nil {
		t.Fatalf("NewPortalAccessPolicyPreview camera returned error: %v", err)
	}
	remote, err := NewPortalAccessPolicyPreview("org.example.ledger", "remote-desktop")
	if err != nil {
		t.Fatalf("NewPortalAccessPolicyPreview remote-desktop returned error: %v", err)
	}

	if camera.Decision != "deny" || remote.Decision != "deny" {
		t.Fatalf("sensitive Portal operations must default to deny: camera=%#v remote=%#v", camera, remote)
	}
}

func testRuntimeSafetyPlan(t *testing.T) Plan {
	t.Helper()
	recipe := Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc"},
	}
	plan, err := NewPlanWithProvenance(recipe, Provenance{Source: "registry", RegistryName: "test-registry", DigestVerified: true, SignatureStatus: "development-only"})
	if err != nil {
		t.Fatalf("NewPlanWithProvenance returned error: %v", err)
	}
	return plan
}
