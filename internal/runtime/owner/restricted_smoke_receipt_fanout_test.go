package owner

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRestrictedOwnerSmokeReceiptFanOutPreviewCoversReadinessAndSupportSurfaces(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordRestrictedOwnerSmokeReceipt(projectRoot(t), stateRoot, RestrictedOwnerSmokeMode, RestrictedOwnerSmokeDirective)
	if err != nil {
		t.Fatalf("RecordRestrictedOwnerSmokeReceipt returned error: %v", err)
	}
	preview, err := NewRestrictedOwnerSmokeReceiptFanOutPreview(stateRoot)
	if err != nil {
		t.Fatalf("NewRestrictedOwnerSmokeReceiptFanOutPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.restricted_owner_smoke_receipt_fanout.v1" ||
		preview.RequestType != "restricted-owner-smoke-receipt-fanout-preview" ||
		preview.Source != "restricted-owner-smoke-execution-receipt" ||
		preview.RuntimeMethod != "GetRestrictedOwnerSmokeReceiptFanOut" ||
		preview.ReadMethod != "GetRestrictedOwnerSmokeReceiptFanOutPreview" ||
		preview.Mode != RestrictedOwnerSmokeMode ||
		preview.ReceiptID != receipt.ReceiptID ||
		preview.ReceiptRelativePath != receipt.RelativePath ||
		preview.ReceiptSHA256 != receipt.SHA256 ||
		preview.ReceiptRecordType != "restricted-owner-smoke-execution-receipt" ||
		!preview.ReceiptReadBack ||
		!preview.ReceiptConsumed ||
		!preview.ReceiptAllChecksPassed ||
		preview.ReceiptCheckCount != 7 ||
		preview.SurfaceCount != 5 ||
		preview.CheckCount != 8 ||
		preview.PassedCheckCount != 8 ||
		!preview.AllChecksPassed ||
		!preview.ReadOnlyFanOut ||
		!preview.ReadinessSurfacesSatisfied ||
		!preview.SupportSurfacesSatisfied {
		t.Fatalf("unexpected restricted owner smoke receipt fan-out preview: %#v", preview)
	}
	if filepath.IsAbs(preview.ReceiptRelativePath) || len(preview.ReceiptSHA256) != 64 {
		t.Fatalf("fan-out must expose only relative receipt path and digest: %#v", preview)
	}
	for _, surface := range preview.Surfaces {
		if surface.EvidenceState != "restricted-smoke-receipt-ready" ||
			surface.ReceiptRelativePath != receipt.RelativePath ||
			surface.ReceiptSHA256 != receipt.SHA256 ||
			!surface.ReadOnly ||
			!surface.UserVisible ||
			!surface.ConsumesReceipt ||
			surface.MutatesRuntime ||
			surface.StartsService ||
			surface.ClaimsSessionBus ||
			surface.ClaimsProductionBus ||
			surface.EnablesWriteMethods ||
			surface.StartsBackend ||
			surface.ExportsSupportBundle ||
			surface.CreatesSupportCase ||
			surface.SendsNotification ||
			surface.ExposesStateRootPath ||
			surface.ExposesBackendDetails {
			t.Fatalf("surface must be a safe read-only receipt projection: %#v", surface)
		}
	}
	if preview.RuntimeOwnerReadiness.ID != "runtime-owner-readiness" ||
		preview.ServiceActivationPreflight.ID != "service-activation-preflight" ||
		preview.CompatibilityOnboarding.ID != "compatibility-onboarding" ||
		preview.SupportBundleManifest.ID != "support-bundle-manifest" ||
		preview.SupportCaseTimeline.ID != "support-case-timeline" {
		t.Fatalf("unexpected fan-out surfaces: %#v", preview.Surfaces)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.StateRootPathExposed ||
		preview.StateRootWritesEnabled ||
		preview.FanOutWritesEnabled ||
		preview.ProductionActivationReady ||
		preview.ProductionOwnerEnabled ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.SupportBundleExported ||
		preview.SupportCaseCreated ||
		preview.NotificationSent ||
		preview.BackendLaunchEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe fan-out gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutPreviewRequiresExistingReceipt(t *testing.T) {
	if _, err := NewRestrictedOwnerSmokeReceiptFanOutPreview(""); err == nil {
		t.Fatalf("fan-out must require an explicit state root")
	}
	if _, err := NewRestrictedOwnerSmokeReceiptFanOutPreview(t.TempDir()); err == nil {
		t.Fatalf("fan-out must require an existing restricted owner smoke receipt")
	}
}

func TestRestrictedOwnerSmokeReceiptFanOutPreviewRejectsTamperedReceipt(t *testing.T) {
	stateRoot := t.TempDir()
	receipt, err := RecordRestrictedOwnerSmokeReceipt(projectRoot(t), stateRoot, RestrictedOwnerSmokeMode, RestrictedOwnerSmokeDirective)
	if err != nil {
		t.Fatalf("RecordRestrictedOwnerSmokeReceipt returned error: %v", err)
	}
	path := filepath.Join(stateRoot, filepath.FromSlash(receipt.RelativePath))
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile receipt returned error: %v", err)
	}
	tampered := bytes.Replace(data, []byte(`"receipt_read_back": true`), []byte(`"receipt_read_back": false`), 1)
	if err := os.WriteFile(path, tampered, 0o600); err != nil {
		t.Fatalf("WriteFile tampered receipt returned error: %v", err)
	}
	if _, err := NewRestrictedOwnerSmokeReceiptFanOutPreview(stateRoot); err == nil {
		t.Fatalf("fan-out must reject tampered restricted owner smoke receipt")
	}
}
