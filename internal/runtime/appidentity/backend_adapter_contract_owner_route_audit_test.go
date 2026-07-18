package appidentity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBackendAdapterContractOwnerRouteAuditKeepsContractFixtureLocal(t *testing.T) {
	preview, err := NewBackendAdapterContractOwnerRouteAuditPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterContractOwnerRouteAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.backend_adapter_contract_owner_route_audit.v1" ||
		preview.RequestType != "backend-adapter-contract-owner-route-audit-preview" ||
		preview.AuditType != "adapter-contract-owner-route-audit" ||
		preview.RuntimeMethod != "GetBackendAdapterContractOwnerRouteAudit" ||
		preview.ReadMethod != "GetBackendAdapterContractOwnerRouteAuditPreview" ||
		preview.SubjectRequestType != "backend-adapter-contract-preview" ||
		preview.SubjectCommand != "backend-adapter-contract-preview" ||
		preview.ProposedOwnerMethod != "GetBackendAdapterContract" ||
		preview.RouteDecision != "remain-fixture-local" ||
		preview.CurrentRouteStatus != "fixture-local-audit-ready-owner-route-blocked" {
		t.Fatalf("unexpected adapter contract owner-route audit schema: %#v", preview)
	}
	if !preview.CLICommandRegistered ||
		!preview.GoReadModelPresent ||
		!preview.FixtureMatrixConsumesContract ||
		preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		!preview.FullContractContainsAdapterIDs ||
		!preview.KDEFacingProjectionPresent ||
		preview.RedactedProfileRoutePresent ||
		!preview.RequiresCallerRoot ||
		preview.AdapterInvocationEnabled ||
		preview.BackendInstallEnabled ||
		preview.BackendDownloadEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.CommandMaterialized ||
		preview.ExecutablePathResolved ||
		preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.RecommendedNextRoute != "owner-local-redacted-adapter-profile-audit" {
		t.Fatalf("unexpected adapter contract owner-route audit decision: %#v", preview)
	}
	if preview.Counts.Total != 9 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Pending != 2 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"cli-preview-registered", "go-read-model-present", "fixture-consumption-present", "owner-route-absent", "production-dbus-absent", "redacted-route-missing", "internal-detail-boundary", "route-decision", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected adapter contract owner-route audit checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe adapter contract owner-route audit gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "backend adapter contract owner-route audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestBackendAdapterContractOwnerRouteAuditFailsClosedWithoutSources(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewBackendAdapterContractOwnerRouteAuditPreview(root)
	if err != nil {
		t.Fatalf("NewBackendAdapterContractOwnerRouteAuditPreview returned error: %v", err)
	}
	if preview.CLICommandRegistered ||
		preview.GoReadModelPresent ||
		preview.FixtureMatrixConsumesContract ||
		preview.OwnerDispatchRoutePresent ||
		preview.ProductionDBusMethodPresent ||
		preview.RedactedProfileRoutePresent ||
		preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady {
		t.Fatalf("missing sources must not produce owner-route readiness: %#v", preview)
	}
	if preview.Counts.Blocked != 3 || preview.RouteDecision != "remain-fixture-local" {
		t.Fatalf("missing sources must fail closed while keeping the decision fixture-local: %#v", preview.Counts)
	}
}
