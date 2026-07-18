package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBackendAdapterRedactedProfileAuditPreviewDefinesOwnerRouteSplit(t *testing.T) {
	preview, err := NewBackendAdapterRedactedProfileAuditPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterRedactedProfileAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.backend_adapter_redacted_profile_audit.v1" ||
		preview.RequestType != "backend-adapter-redacted-profile-audit-preview" ||
		preview.AuditType != "owner-local-redacted-adapter-profile-audit" ||
		preview.RuntimeMethod != "GetBackendAdapterProfileAudit" ||
		preview.ReadMethod != "GetBackendAdapterProfileAuditPreview" ||
		preview.SubjectRequestType != "backend-adapter-contract-preview" ||
		preview.SubjectReadMethod != "GetBackendAdapterContractPreview" ||
		preview.OwnerRouteMethod != "GetBackendAdapterProfileAudit" ||
		preview.RouteDecision != "redacted-profile-route-ready" {
		t.Fatalf("unexpected redacted adapter profile audit schema: %#v", preview)
	}
	if preview.ContractProfileCount != 3 ||
		len(preview.Profiles) != 3 ||
		!sameStrings(preview.ProfileIDs, []string{"automatic", "performance-priority", "compatibility-priority"}) ||
		preview.Counts.Total != 7 ||
		preview.Counts.Passed != 7 ||
		preview.Counts.Blocked != 0 ||
		preview.Counts.RedactedProfiles != 3 ||
		preview.Counts.UserVisibleProfiles != 3 ||
		preview.Counts.EnabledInvocations != 0 ||
		preview.Counts.EnabledLaunches != 0 ||
		preview.Counts.ExposedBackendDetails != 0 {
		t.Fatalf("unexpected redacted adapter profile audit counts: counts=%#v ids=%#v", preview.Counts, preview.ProfileIDs)
	}
	if !preview.FullContractConsumed ||
		!preview.KDEFacingProjectionConsumed ||
		!preview.InternalAdapterIDsRedacted ||
		!preview.InternalProfilePathsRedacted ||
		!preview.OwnerLocalRouteCandidateReady ||
		!preview.FullContractFixtureLocal ||
		preview.ProductionDBusExposureReady ||
		preview.CallerStateRootRequired ||
		!preview.CLIProjectRootOnly {
		t.Fatalf("unexpected redacted adapter profile audit route decision: %#v", preview)
	}
	if !sameStrings(preview.CheckIDs, []string{"contract-profile-projection-consumed", "internal-adapter-ids-redacted", "owner-local-route-candidate", "full-contract-fixture-local", "production-dbus-not-claimed", "no-caller-state-root", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected redacted adapter profile audit checks: %#v", preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.AdapterInvocationEnabled ||
		preview.BackendInstallEnabled ||
		preview.BackendDownloadEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.CommandMaterialized ||
		preview.ExecutablePathResolved ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe redacted adapter profile audit gates: %#v", preview)
	}
}

func TestBackendAdapterRedactedProfileAuditOutputIsRedacted(t *testing.T) {
	preview, err := NewBackendAdapterRedactedProfileAuditPreview(projectRootForRuntimeServiceBindingTest(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterRedactedProfileAuditPreview returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"wine", "proton", "windows-vm", "prefix", ".exe", "program files", "qemu-system", ".wine", "/tmp", "/users", "/private", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("redacted adapter profile audit exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
	for _, profile := range preview.Profiles {
		if !profile.SourceProfileMatched ||
			!profile.InternalAdapterIDRedacted ||
			!profile.InternalProfilePathRedacted ||
			profile.BackendDetailsExposed ||
			profile.AdapterInvocationEnabled ||
			profile.LaunchEnabled ||
			profile.HostRootModified {
			t.Fatalf("profile is not redacted or safe: %#v", profile)
		}
	}
	if err := validateNoBackendTerms(preview, "redacted adapter profile audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestBackendAdapterRedactedProfileAuditFailsClosedWithoutOwnerRouteSource(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewBackendAdapterRedactedProfileAuditPreview(root)
	if err != nil {
		t.Fatalf("NewBackendAdapterRedactedProfileAuditPreview returned error: %v", err)
	}
	if preview.OwnerLocalRouteCandidateReady ||
		preview.ProductionDBusExposureReady ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.HostRootModified {
		t.Fatalf("missing owner route sources must fail closed: %#v", preview)
	}
	if preview.Counts.Blocked != 1 ||
		preview.Counts.Passed != 6 ||
		!containsString(preview.CheckIDs, "owner-local-route-candidate") {
		t.Fatalf("missing owner route source must block only route readiness: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
}
