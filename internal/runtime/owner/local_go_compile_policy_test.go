package owner

import "testing"

func TestLocalGoCompilePolicyPreviewRequiresQ4ByDefault(t *testing.T) {
	preview, err := NewLocalGoCompilePolicyPreview(projectRootForOwnerTest(t))
	if err != nil {
		t.Fatalf("NewLocalGoCompilePolicyPreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.local_go_compile_policy.v1" ||
		preview.RequestType != "local-go-compile-policy-preview" ||
		preview.PolicyType != "q4-first-go-compilation-policy" ||
		preview.PolicyDecision != "local-go-compilation-policy-ready-q4-first" {
		t.Fatalf("unexpected local Go compile policy schema: %#v", preview)
	}
	if preview.DefaultRemoteHost != "root@q4" ||
		preview.RemoteBuildCommand != "ruby scripts/remote_go_build.rb --execute" ||
		preview.RemoteTestCommand != "ruby scripts/remote_go_test.rb --execute" ||
		preview.LocalOverrideEnvironment != "XNIX_ALLOW_LOCAL_GO_COMPILE=1" {
		t.Fatalf("unexpected q4 compile policy commands: %#v", preview)
	}
	if preview.LocalGoCompilationAllowedByDefault ||
		preview.LocalGoRunAllowedByDefault ||
		!preview.Q4CompilationRequiredByDefault ||
		!preview.RemoteGoBuildRunnerPresent ||
		!preview.RemoteGoTestRunnerPresent ||
		!preview.RemoteGoCachesPinnedToQ4 ||
		!preview.TargetedSmallVersionTesting ||
		!preview.TwentiethVersionFullGateRequired ||
		!preview.ProtectedClaudePackageExcluded ||
		!preview.HostCompilationAvoidedObservable {
		t.Fatalf("unexpected q4 compile policy readiness: %#v", preview)
	}
	if preview.Q4HostRootModified ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired {
		t.Fatalf("unsafe q4 compile policy gates: %#v", preview)
	}
	if preview.Counts.Total != 8 ||
		preview.Counts.Passed != 8 ||
		preview.Counts.Pending != 0 ||
		preview.Counts.Blocked != 0 ||
		!sameStrings(preview.CheckIDs, []string{"q4-policy-document-present", "claude-guide-local-compile-guard-present", "remote-go-build-runner-ready", "remote-go-test-runner-ready", "remote-go-cache-roots-pinned", "protected-claude-package-excluded", "twentieth-version-full-gate-preserved", "unsafe-host-gates-closed"}) {
		t.Fatalf("unexpected q4 compile policy checks: counts=%#v ids=%#v", preview.Counts, preview.CheckIDs)
	}
}
