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
	if len(preview.SupportedKnownPortableAppIDs) != 1 || preview.SupportedKnownPortableAppIDs[0] != "org.xnix.external.notepadplusplus" {
		t.Fatalf("unexpected supported q4 known portable app ids: %#v", preview.SupportedKnownPortableAppIDs)
	}
	if !preview.KnownCatalogApp ||
		!preview.PortableDirectoryExternalApp ||
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

func TestPreviewQ4KnownPortableWinAppRunPlanRejectsKnownNonBundleCatalogApp(t *testing.T) {
	_, err := PreviewQ4KnownPortableWinAppRunPlan(Q4KnownPortableWinAppRunPlanRequest{
		Version:             currentProjectVersion(t),
		AppID:               "7zr",
		RemoteHost:          "root@q4",
		RemoteMaterialsRoot: "/home/xnix-run-materials",
		RemoteSourceRoot:    "/home/xnix-build/xnix-known-portable",
		RemoteBuildRoot:     "/home/xnix-build-cache",
	})
	if err == nil {
		t.Fatalf("expected known non-bundle app rejection")
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
