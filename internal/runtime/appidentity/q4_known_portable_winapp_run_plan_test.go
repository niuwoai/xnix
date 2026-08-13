package appidentity

import "testing"

func TestPreviewQ4KnownPortableWinAppRunPlanConsumesRuntimeCatalog(t *testing.T) {
	preview, err := PreviewQ4KnownPortableWinAppRunPlan(Q4KnownPortableWinAppRunPlanRequest{
		Version:             currentProjectVersion(t),
		AppID:               "org.xnix.external.notepadplusplus",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
		RemoteSourceRoot:    "/home/xnix-build/xnix-known-portable",
		RemoteBuildRoot:     "/home/xnix-build-cache",
		Execute:             true,
	})
	if err != nil {
		t.Fatalf("PreviewQ4KnownPortableWinAppRunPlan returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != Q4KnownPortableWinAppRunPlanSchemaVersion ||
		preview.RequestType != Q4KnownPortableWinAppRunPlanRequestType ||
		preview.AppID != "org.xnix.external.notepadplusplus" ||
		preview.DisplayName != "Notepad++ Portable" ||
		preview.AppVersion != "8.9.7" ||
		preview.CatalogArtifactKind != "portable-zip-bundle" ||
		preview.DownloadArtifactName != "npp.8.9.7.portable.zip" ||
		preview.ExecutableRelativePath != "notepad++.exe" {
		t.Fatalf("unexpected known portable catalog-backed plan: %#v", preview)
	}
	if len(preview.SupportedKnownPortableAppIDs) != 2 ||
		!sliceContainsString(preview.SupportedKnownPortableAppIDs, "org.xnix.external.notepadplusplus") ||
		!sliceContainsString(preview.SupportedKnownPortableAppIDs, "org.xnix.external.putty") {
		t.Fatalf("unexpected supported q4 known portable app ids: %#v", preview.SupportedKnownPortableAppIDs)
	}
	if !preview.KnownCatalogApp ||
		!preview.PortableDirectoryExternalApp ||
		preview.SingleFileExternalApp ||
		!preview.OfficialDownloadRequired ||
		!preview.PinnedChecksumRequired ||
		!preview.KnownBundleImportRequired ||
		!preview.KnownBundleStageLaunchRequired ||
		!preview.RuntimeAcceptanceRequired ||
		!preview.Q4DownloadRequired ||
		!preview.Q4ExtractRequired ||
		!preview.Q4CompileRequired ||
		!preview.Q4ExecutionRequired ||
		!preview.Q4BuildPreferred ||
		preview.HostCompilationRequired ||
		!preview.HostCompilationAvoided ||
		!preview.HostDownloadAvoided ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.FullSmokeRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		!preview.PlanReady {
		t.Fatalf("unexpected known portable plan flags: %#v", preview)
	}
	if !sliceContainsString(preview.DelegatedCommand, "--execute") ||
		!sliceContainsString(preview.DelegatedCommand, "org.xnix.external.notepadplusplus") {
		t.Fatalf("unexpected delegated command: %#v", preview.DelegatedCommand)
	}
}

func TestPreviewQ4KnownPortableWinAppRunPlanSelectsPuttySingleExecutableLane(t *testing.T) {
	preview, err := PreviewQ4KnownPortableWinAppRunPlan(Q4KnownPortableWinAppRunPlanRequest{
		Version:             currentProjectVersion(t),
		AppID:               "org.xnix.external.putty",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
		RemoteSourceRoot:    "/home/xnix-build/xnix-known-portable",
		RemoteBuildRoot:     "/home/xnix-build-cache",
		Output:              "/tmp/xnix-output/putty-known-run.json",
		MarkdownOutput:      "/tmp/xnix-output/putty-known-run.md",
		Execute:             true,
	})
	if err != nil {
		t.Fatalf("PreviewQ4KnownPortableWinAppRunPlan returned error: %v", err)
	}

	if preview.AppID != "org.xnix.external.putty" ||
		preview.DisplayName != "PuTTY" ||
		preview.AppVersion != "0.84" ||
		preview.CatalogArtifactKind != "single-executable" ||
		preview.DownloadArtifactName != "putty.exe" ||
		preview.ExecutableRelativePath != "putty.exe" ||
		preview.LaunchSourceRequestType != "external-winapp-import-stage-and-launch" ||
		preview.DelegatedScript != "scripts/q4_putty_external_winapp_smoke.rb" ||
		preview.DelegatedRequestType != "q4-putty-external-winapp-smoke" ||
		preview.DelegatedAcceptanceRequestType != Q4StagedExternalWinAppAcceptanceRequestType {
		t.Fatalf("unexpected PuTTY single-executable plan: %#v", preview)
	}
	if !preview.KnownCatalogApp ||
		preview.PortableDirectoryExternalApp ||
		!preview.SingleFileExternalApp ||
		!preview.OfficialDownloadRequired ||
		!preview.PinnedChecksumRequired ||
		preview.KnownBundleImportRequired ||
		preview.KnownBundleStageLaunchRequired ||
		!preview.RuntimeAcceptanceRequired ||
		!preview.Q4DownloadRequired ||
		preview.Q4ExtractRequired ||
		!preview.Q4CompileRequired ||
		!preview.Q4ExecutionRequired ||
		!preview.HostCompilationAvoided ||
		!preview.HostDownloadAvoided ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.HostNetworkingRequired ||
		preview.DockerSocketMounted ||
		preview.BroadHostMountRequired ||
		!preview.PlanReady {
		t.Fatalf("unexpected PuTTY single-executable plan flags: %#v", preview)
	}
	if !sliceContainsString(preview.DelegatedCommand, "scripts/q4_putty_external_winapp_smoke.rb") ||
		!sliceContainsString(preview.DelegatedCommand, "--execute") ||
		!sliceContainsString(preview.DelegatedCommand, "/tmp/xnix-output/putty-known-run.json") {
		t.Fatalf("unexpected PuTTY delegated command: %#v", preview.DelegatedCommand)
	}
}

func TestPreviewQ4KnownPortableWinAppRunPlanRejectsKnownNonRunnableCatalogApp(t *testing.T) {
	_, err := PreviewQ4KnownPortableWinAppRunPlan(Q4KnownPortableWinAppRunPlanRequest{
		Version:             currentProjectVersion(t),
		AppID:               "7zr",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
		RemoteSourceRoot:    "/home/xnix-build/xnix-known-portable",
		RemoteBuildRoot:     "/home/xnix-build-cache",
	})
	if err == nil {
		t.Fatalf("expected known non-runnable app rejection")
	}
}

func sliceContainsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
