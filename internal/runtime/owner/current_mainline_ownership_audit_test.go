package owner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCurrentMainlineOwnershipAuditPreviewModelsAutonomousMainline(t *testing.T) {
	preview, err := NewCurrentMainlineOwnershipAuditPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewCurrentMainlineOwnershipAuditPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.current_mainline_ownership_audit.v1" ||
		preview.RequestType != "current-mainline-ownership-audit-preview" ||
		preview.AuditType != "current-mainline-ownership-audit" ||
		preview.AuditDecision != "current-mainline-ownership-audit-ready-autonomous-mainline" {
		t.Fatalf("unexpected current mainline ownership schema: %#v", preview)
	}
	if !preview.CurrentMainlineDocumentPresent ||
		!preview.CodexOwnedMainline ||
		preview.ExternalAgentDependencyRequired ||
		!preview.HistoricalExternalAgentDocumentsAllowed ||
		!preview.KDEFlagshipOnly ||
		preview.GNOMEFirstReleaseTarget ||
		preview.XFCEFirstReleaseTarget ||
		!preview.GoRuntimeOwned ||
		!preview.RubyTestHarnessOnly ||
		!preview.CLowLevelOnly ||
		!preview.KDEPluginsPresentationOnly ||
		!preview.RuntimeOwnsCompatibilityDecisions ||
		!preview.SevenKDEEntrypointsPresent ||
		!preview.SafeNextTaskPresent ||
		!preview.FullGateRequiresAuthorization {
		t.Fatalf("unexpected current mainline ownership decision: %#v", preview)
	}
	if preview.DockerExecuted ||
		preview.QEMUExecuted ||
		preview.NetworkFetchEnabled ||
		preview.PackageManagerInvoked ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.KDEConfigurationWritten ||
		preview.PortalCallExecuted ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.FilePathsExposed ||
		preview.FileContentRead ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unsafe current mainline ownership gates: %#v", preview)
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"current-mainline-document-present", "external-agent-dependency-removed", "kde-flagship-only", "runtime-ownership-language-split", "seven-kde-entrypoints-present", "safe-next-task-present", "full-gate-authorization-required", "unsafe-runtime-and-host-gates-closed"}) {
		t.Fatalf("unexpected current mainline ownership checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
	if err := validateNoBackendTerms(preview, "current mainline ownership audit test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestCurrentMainlineOwnershipAuditPreviewFailsClosedWithoutMainlineDocument(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte(currentProjectVersion(t)+"\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	preview, err := NewCurrentMainlineOwnershipAuditPreview(root)
	if err != nil {
		t.Fatalf("NewCurrentMainlineOwnershipAuditPreview returned error: %v", err)
	}

	if preview.CurrentMainlineDocumentPresent ||
		preview.CodexOwnedMainline ||
		preview.ExternalAgentDependencyRequired ||
		preview.KDEFlagshipOnly ||
		preview.GoRuntimeOwned ||
		preview.SevenKDEEntrypointsPresent ||
		preview.SafeNextTaskPresent ||
		preview.FullGateRequiresAuthorization ||
		preview.AuditDecision != "current-mainline-ownership-audit-blocked" ||
		preview.Counts.Total != 8 ||
		preview.Counts.Passed != 1 ||
		preview.Counts.Blocked != 7 {
		t.Fatalf("missing current mainline document must fail closed: %#v", preview)
	}
}
