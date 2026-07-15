package appidentity

import "testing"

func TestBackendCapabilityMatrixPreviewPlansCapabilitiesWithoutSelection(t *testing.T) {
	preview, err := NewBackendCapabilityMatrixPreview()
	if err != nil {
		t.Fatalf("NewBackendCapabilityMatrixPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.backend_capability_matrix.v1" ||
		preview.RequestType != "backend-capability-matrix-preview" ||
		preview.MatrixType != "compatibility-backend-capability-matrix" ||
		preview.RuntimeMethod != "GetBackendCapabilityMatrix" ||
		preview.ReadMethod != "GetBackendCapabilityMatrixPreview" {
		t.Fatalf("unexpected backend capability matrix schema: %#v", preview)
	}
	if !sameStrings(preview.ProfileIDs, []string{"local-compatibility", "isolated-compatibility"}) ||
		preview.ProfileCount != 2 ||
		preview.CapabilityCount != 7 {
		t.Fatalf("unexpected backend capability matrix profiles: %#v", preview)
	}
	// Diagnostics is ready in each of the two profiles; everything else is pending.
	if preview.ReadyCapabilityCount != 2 ||
		preview.PendingCapabilityCount != 12 ||
		preview.BlockedCapabilityCount != 0 {
		t.Fatalf("unexpected capability counts: ready=%d pending=%d blocked=%d", preview.ReadyCapabilityCount, preview.PendingCapabilityCount, preview.BlockedCapabilityCount)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.SelectionEnabled ||
		preview.BackendLaunchEnabled ||
		preview.CapabilityActivationEnabled ||
		preview.RequestObjectsCreated ||
		preview.StateRootCreated ||
		preview.SnapshotsCreated ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected backend capability matrix safety flags: %#v", preview)
	}

	local := preview.Profiles[0]
	if local.ID != "local-compatibility" ||
		local.SelectionState != "planned" ||
		local.CapabilityCount != 7 ||
		local.ReadyCount != 1 ||
		local.PendingCount != 6 ||
		local.ProfileReady ||
		local.SelectionEnabled ||
		local.BackendDetailsExposed {
		t.Fatalf("unexpected local profile: %#v", local)
	}
	// The Portal-gated bridges must be flagged as requiring review.
	var fileBridge BackendCapability
	for _, capability := range local.Capabilities {
		if capability.ID == "file-bridge" {
			fileBridge = capability
		}
	}
	if !fileBridge.RequiresPortalReview || !fileBridge.RequiresStateRoot || fileBridge.RequiresSnapshot || fileBridge.ActivationEnabled {
		t.Fatalf("unexpected file-bridge capability: %#v", fileBridge)
	}

	if err := validateNoBackendTerms(preview, "backend capability matrix preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
