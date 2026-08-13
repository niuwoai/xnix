package appidentity

import (
	"errors"
	"strings"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	Q4KnownPortableWinAppRunPlanSchemaVersion = "xnix.runtime.q4_known_portable_winapp_run_plan.v1"
	Q4KnownPortableWinAppRunPlanRequestType   = "q4-known-portable-winapp-run-plan-preview"
)

type Q4KnownPortableWinAppRunPlanRequest struct {
	Version             string
	AppID               string
	RemoteHost          string
	RemoteMaterialsRoot string
	RemoteSourceRoot    string
	RemoteBuildRoot     string
	Output              string
	MarkdownOutput      string
	Execute             bool
}

type Q4KnownPortableWinAppRunPlanPreview struct {
	Version                              string   `json:"version"`
	SchemaVersion                        string   `json:"schema_version"`
	RequestType                          string   `json:"request_type"`
	Source                               string   `json:"source"`
	RuntimeMethod                        string   `json:"runtime_method"`
	ReadMethod                           string   `json:"read_method"`
	AppID                                string   `json:"app_id"`
	DisplayName                          string   `json:"display_name"`
	AppVersion                           string   `json:"app_version"`
	CatalogArtifactKind                  string   `json:"catalog_artifact_kind"`
	DownloadArtifactName                 string   `json:"download_artifact_name"`
	ExecutableRelativePath               string   `json:"executable_relative_path"`
	SupportedKnownPortableAppIDs         []string `json:"supported_known_portable_app_ids"`
	KnownCatalogApp                      bool     `json:"known_catalog_app"`
	PortableDirectoryExternalApp         bool     `json:"portable_directory_external_app"`
	OfficialDownloadRequired             bool     `json:"official_download_required"`
	PinnedChecksumRequired               bool     `json:"pinned_checksum_required"`
	KnownBundleImportRequired            bool     `json:"known_bundle_import_required"`
	KnownBundleStageLaunchRequired       bool     `json:"known_bundle_stage_launch_required"`
	RuntimeAcceptanceRequired            bool     `json:"runtime_acceptance_required"`
	LaunchSourceRequestType              string   `json:"launch_source_request_type"`
	RemoteHostConfigured                 bool     `json:"remote_host_configured"`
	RemoteHostExposedToDesktop           bool     `json:"remote_host_exposed_to_desktop"`
	RemoteMaterialsRootClass             string   `json:"remote_materials_root_class"`
	RemoteSourceRootClass                string   `json:"remote_source_root_class"`
	RemoteBuildRootClass                 string   `json:"remote_build_root_class"`
	RemotePathsExposedToDesktop          bool     `json:"remote_paths_exposed_to_desktop"`
	Q4DownloadRequired                   bool     `json:"q4_download_required"`
	Q4ExtractRequired                    bool     `json:"q4_extract_required"`
	Q4CompileRequired                    bool     `json:"q4_compile_required"`
	Q4ExecutionRequired                  bool     `json:"q4_execution_required"`
	Q4BuildPreferred                     bool     `json:"q4_build_preferred"`
	HostCompilationRequired              bool     `json:"host_compilation_required"`
	HostCompilationAvoided               bool     `json:"host_compilation_avoided"`
	HostDownloadAvoided                  bool     `json:"host_download_avoided"`
	LocalHostRole                        string   `json:"local_host_role"`
	DelegatedScript                      string   `json:"delegated_script"`
	DelegatedRequestType                 string   `json:"delegated_request_type"`
	DelegatedAcceptanceRequestType       string   `json:"delegated_acceptance_request_type"`
	DelegatedCommand                     []string `json:"delegated_command"`
	DelegatedCommandExposedToOperator    bool     `json:"delegated_command_exposed_to_operator"`
	DelegatedCommandExposedToDesktop     bool     `json:"delegated_command_exposed_to_desktop"`
	AcceptedApplicationDetailStateNeeded string   `json:"accepted_application_detail_state_needed"`
	RuntimeOwned                         bool     `json:"runtime_owned"`
	GoRuntimeBacked                      bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                       bool     `json:"kde_policy_owner"`
	FullSmokeRequired                    bool     `json:"full_smoke_required"`
	HostRootModified                     bool     `json:"host_root_modified"`
	PrivilegedContainerRequired          bool     `json:"privileged_container_required"`
	HostNetworkingRequired               bool     `json:"host_networking_required"`
	DockerSocketMounted                  bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired               bool     `json:"broad_host_mount_required"`
	PlanReady                            bool     `json:"plan_ready"`
	BlockedActions                       []string `json:"blocked_actions"`
	DesktopSafeSummary                   string   `json:"desktop_safe_summary"`
}

func PreviewQ4KnownPortableWinAppRunPlan(request Q4KnownPortableWinAppRunPlanRequest) (Q4KnownPortableWinAppRunPlanPreview, error) {
	version := strings.TrimSpace(request.Version)
	if version == "" {
		return Q4KnownPortableWinAppRunPlanPreview{}, errors.New("q4 known portable Windows app run plan requires version metadata")
	}
	appID, err := requiredSingleLine("app id", request.AppID)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}
	app, err := winapp.LookupKnownPortableApp(appID)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}
	supportedAppIDs := supportedQ4KnownPortableWinAppRunAppIDs()
	if !stringSliceContains(supportedAppIDs, app.ID) {
		return Q4KnownPortableWinAppRunPlanPreview{}, errors.New("q4 known portable Windows app run plan requires a Runtime catalog portable bundle app")
	}
	remoteHost, err := requiredSingleLine("remote host", request.RemoteHost)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}
	remoteMaterialsClass, err := q4RemoteMaterialsRootClass(request.RemoteMaterialsRoot)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}
	remoteSourceClass, err := q4RemoteMaterialsRootClass(request.RemoteSourceRoot)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}
	remoteBuildClass, err := q4RemoteMaterialsRootClass(request.RemoteBuildRoot)
	if err != nil {
		return Q4KnownPortableWinAppRunPlanPreview{}, err
	}

	command := []string{
		"ruby",
		"scripts/q4_known_portable_winapp_run.rb",
		"--app", appID,
		"--remote", remoteHost,
		"--remote-materials-root", request.RemoteMaterialsRoot,
		"--remote-source-root", request.RemoteSourceRoot,
		"--remote-build-root", request.RemoteBuildRoot,
	}
	if output := strings.TrimSpace(request.Output); output != "" {
		command = append(command, "--output", output)
	}
	if markdownOutput := strings.TrimSpace(request.MarkdownOutput); markdownOutput != "" {
		command = append(command, "--markdown-output", markdownOutput)
	}
	if request.Execute {
		command = append(command, "--execute")
	}

	return Q4KnownPortableWinAppRunPlanPreview{
		Version:                              version,
		SchemaVersion:                        Q4KnownPortableWinAppRunPlanSchemaVersion,
		RequestType:                          Q4KnownPortableWinAppRunPlanRequestType,
		Source:                               "runtime-q4-known-portable-winapp-run-plan+catalog-backed-runner",
		RuntimeMethod:                        "PlanQ4KnownPortableWinAppRun",
		ReadMethod:                           "GetQ4KnownPortableWinAppRunPlanPreview",
		AppID:                                app.ID,
		DisplayName:                          app.DisplayName,
		AppVersion:                           app.Version,
		CatalogArtifactKind:                  app.ArtifactKind,
		DownloadArtifactName:                 app.DownloadArtifactName,
		ExecutableRelativePath:               app.ExecutableRelativePath,
		SupportedKnownPortableAppIDs:         supportedAppIDs,
		KnownCatalogApp:                      true,
		PortableDirectoryExternalApp:         true,
		OfficialDownloadRequired:             true,
		PinnedChecksumRequired:               true,
		KnownBundleImportRequired:            true,
		KnownBundleStageLaunchRequired:       true,
		RuntimeAcceptanceRequired:            true,
		LaunchSourceRequestType:              "windows-known-app-bundle-stage-and-launch",
		RemoteHostConfigured:                 true,
		RemoteHostExposedToDesktop:           false,
		RemoteMaterialsRootClass:             remoteMaterialsClass,
		RemoteSourceRootClass:                remoteSourceClass,
		RemoteBuildRootClass:                 remoteBuildClass,
		RemotePathsExposedToDesktop:          false,
		Q4DownloadRequired:                   true,
		Q4ExtractRequired:                    true,
		Q4CompileRequired:                    true,
		Q4ExecutionRequired:                  true,
		Q4BuildPreferred:                     true,
		HostCompilationRequired:              false,
		HostCompilationAvoided:               true,
		HostDownloadAvoided:                  true,
		LocalHostRole:                        "operator-plan-and-artifact-review-only",
		DelegatedScript:                      "scripts/q4_known_portable_winapp_run.rb",
		DelegatedRequestType:                 "q4-known-portable-winapp-run",
		DelegatedAcceptanceRequestType:       Q4KnownPortableBundleWinAppAcceptanceRequestType,
		DelegatedCommand:                     command,
		DelegatedCommandExposedToOperator:    true,
		DelegatedCommandExposedToDesktop:     false,
		AcceptedApplicationDetailStateNeeded: "runtime-accepted-real-app-run",
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		FullSmokeRequired:                    false,
		HostRootModified:                     false,
		PrivilegedContainerRequired:          false,
		HostNetworkingRequired:               false,
		DockerSocketMounted:                  false,
		BroadHostMountRequired:               false,
		PlanReady:                            true,
		BlockedActions: []string{
			"Do not download, extract, compile, or launch the Windows app on the local host.",
			"Do not expose q4 paths, remote hosts, backend commands, or raw launcher details to KDE-facing summaries.",
			"Do not bypass the known portable bundle import, staged launch, Runtime detail, KDE detail, and Go-owned acceptance chain.",
			"Do not run privileged containers, host networking, Docker socket mounts, or broad host mounts.",
		},
		DesktopSafeSummary: app.DisplayName + " is planned as a catalog-backed q4 known portable Windows app run through Runtime-owned import, staged launch, KDE detail, and Go-owned acceptance evidence.",
	}, nil
}

func supportedQ4KnownPortableWinAppRunAppIDs() []string {
	apps := winapp.KnownPortableCatalog()
	ids := make([]string, 0, len(apps))
	for _, app := range apps {
		if app.PortableBundleArchive && app.ArtifactKind == winapp.KnownPortableArtifactZipBundle && strings.TrimSpace(app.ExecutableRelativePath) != "" && strings.TrimSpace(app.DownloadArtifactName) != "" {
			ids = append(ids, app.ID)
		}
	}
	return ids
}

func stringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
