package owner

type CurrentMainlineOwnershipAuditPreview struct {
	Version                                 string                               `json:"version"`
	SchemaVersion                           string                               `json:"schema_version"`
	RequestType                             string                               `json:"request_type"`
	AuditType                               string                               `json:"audit_type"`
	Source                                  string                               `json:"source"`
	AuditDecision                           string                               `json:"audit_decision"`
	CurrentMainlineDocumentPresent          bool                                 `json:"current_mainline_document_present"`
	CodexOwnedMainline                      bool                                 `json:"codex_owned_mainline"`
	ExternalAgentDependencyRequired         bool                                 `json:"external_agent_dependency_required"`
	HistoricalExternalAgentDocumentsAllowed bool                                 `json:"historical_external_agent_documents_allowed"`
	KDEFlagshipOnly                         bool                                 `json:"kde_flagship_only"`
	GNOMEFirstReleaseTarget                 bool                                 `json:"gnome_first_release_target"`
	XFCEFirstReleaseTarget                  bool                                 `json:"xfce_first_release_target"`
	GoRuntimeOwned                          bool                                 `json:"go_runtime_owned"`
	RubyTestHarnessOnly                     bool                                 `json:"ruby_test_harness_only"`
	CLowLevelOnly                           bool                                 `json:"c_low_level_only"`
	KDEPluginsPresentationOnly              bool                                 `json:"kde_plugins_presentation_only"`
	RuntimeOwnsCompatibilityDecisions       bool                                 `json:"runtime_owns_compatibility_decisions"`
	SevenKDEEntrypointsPresent              bool                                 `json:"seven_kde_entrypoints_present"`
	SafeNextTaskPresent                     bool                                 `json:"safe_next_task_present"`
	FullGateRequiresAuthorization           bool                                 `json:"full_gate_requires_authorization"`
	DockerExecuted                          bool                                 `json:"docker_executed"`
	QEMUExecuted                            bool                                 `json:"qemu_executed"`
	NetworkFetchEnabled                     bool                                 `json:"network_fetch_enabled"`
	PackageManagerInvoked                   bool                                 `json:"package_manager_invoked"`
	ProductionBusClaimed                    bool                                 `json:"production_bus_claimed"`
	WriteMethodsEnabled                     bool                                 `json:"write_methods_enabled"`
	RuntimeWritesEnabled                    bool                                 `json:"runtime_writes_enabled"`
	KDEConfigurationWritten                 bool                                 `json:"kde_configuration_written"`
	PortalCallExecuted                      bool                                 `json:"portal_call_executed"`
	BackendLaunchEnabled                    bool                                 `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                                 `json:"backend_process_started"`
	HostRootModified                        bool                                 `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                                 `json:"privileged_container_required"`
	StateRootPathExposed                    bool                                 `json:"state_root_path_exposed"`
	FilePathsExposed                        bool                                 `json:"file_paths_exposed"`
	FileContentRead                         bool                                 `json:"file_content_read"`
	RawCommandExposed                       bool                                 `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                                 `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                                 `json:"backend_details_exposed"`
	Checks                                  []CurrentMainlineOwnershipAuditCheck `json:"checks"`
	CheckIDs                                []string                             `json:"check_ids"`
	Counts                                  CurrentMainlineOwnershipAuditCounts  `json:"counts"`
	BlockedActions                          []string                             `json:"blocked_actions"`
	NextRequirements                        []string                             `json:"next_requirements"`
	DesktopSafeSummary                      string                               `json:"desktop_safe_summary"`
}

type CurrentMainlineOwnershipAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type CurrentMainlineOwnershipAuditCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
}

func NewCurrentMainlineOwnershipAuditPreview(root string) (CurrentMainlineOwnershipAuditPreview, error) {
	version, err := readProductionDBusGateReviewVersion(root)
	if err != nil {
		return CurrentMainlineOwnershipAuditPreview{}, err
	}
	source := productionAuthorizationReadSources(root, []string{"docs/xnix-current-mainline.md"})
	documentReady := productionAuthorizationHasAll(source, []string{"Xnix Current Mainline", "Codex-owned current mainline", "Product Goal", "Ownership Rules", "First-Release KDE Entrypoints", "Current Safe Next Task", "Safety Gates"})
	kdeReady := productionAuthorizationHasAll(source, []string{"KDE Plasma is the only flagship desktop", "GNOME and XFCE", "not first-release integration targets"})
	ownershipReady := productionAuthorizationHasAll(source, []string{"Runtime business logic is Go-first", "Ruby is for tests", "C remains for low-level", "Runtime owns recipes"})
	entrypointsReady := productionAuthorizationHasAll(source, []string{"Start menu integration", "Task manager identity", "Dolphin file-manager actions", "System tray status", "Notification Center events", "AI Compatibility Center pages", "Unified settings"})
	nextTaskReady := productionAuthorizationHasAll(source, []string{"redacted status writer authorization gate audit", "consume the v0.2.380 redacted status persistence write-model audit", "Xnix current mainline document"})
	safetyReady := productionAuthorizationHasAll(source, []string{"No production D-Bus ownership", "No Runtime write methods", "No KDE configuration writes", "No real Portal calls", "No backend launch", "No Docker, QEMU, or full build unless explicitly authorized"})
	ready := documentReady && kdeReady && ownershipReady && entrypointsReady && nextTaskReady && safetyReady
	preview := CurrentMainlineOwnershipAuditPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.current_mainline_ownership_audit.v1",
		RequestType:                             "current-mainline-ownership-audit-preview",
		AuditType:                               "current-mainline-ownership-audit",
		Source:                                  "xnix-current-mainline",
		AuditDecision:                           "current-mainline-ownership-audit-blocked",
		CurrentMainlineDocumentPresent:          documentReady,
		CodexOwnedMainline:                      documentReady,
		ExternalAgentDependencyRequired:         false,
		HistoricalExternalAgentDocumentsAllowed: true,
		KDEFlagshipOnly:                         kdeReady,
		GNOMEFirstReleaseTarget:                 false,
		XFCEFirstReleaseTarget:                  false,
		GoRuntimeOwned:                          ownershipReady,
		RubyTestHarnessOnly:                     ownershipReady,
		CLowLevelOnly:                           ownershipReady,
		KDEPluginsPresentationOnly:              ownershipReady,
		RuntimeOwnsCompatibilityDecisions:       ownershipReady,
		SevenKDEEntrypointsPresent:              entrypointsReady,
		SafeNextTaskPresent:                     nextTaskReady,
		FullGateRequiresAuthorization:           safetyReady,
		DockerExecuted:                          false,
		QEMUExecuted:                            false,
		NetworkFetchEnabled:                     false,
		PackageManagerInvoked:                   false,
		ProductionBusClaimed:                    false,
		WriteMethodsEnabled:                     false,
		RuntimeWritesEnabled:                    false,
		KDEConfigurationWritten:                 false,
		PortalCallExecuted:                      false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		StateRootPathExposed:                    false,
		FilePathsExposed:                        false,
		FileContentRead:                         false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions: []string{
			"treat historical external-agent handoff documents as the required source for new mainline work",
			"enable full build, container, or virtual-machine checks without explicit operator authorization",
			"claim production ownership, enable write methods, write KDE configuration, call desktop portals, launch backends, or mutate the host root",
		},
		NextRequirements: []string{
			"Route new implementation tasks through docs/xnix-current-mainline.md by default.",
			"Let future writer authorization gates consume the Xnix current mainline document instead of external-agent dispatch sheets.",
			"Keep every new small version targeted-test only unless a twentieth-version full gate is explicitly authorized.",
		},
		DesktopSafeSummary: "The current mainline ownership audit establishes docs/xnix-current-mainline.md as the Codex-owned planning source for new Xnix implementation work while preserving KDE as the only flagship first-release desktop, keeping Runtime product logic Go-first, and leaving full gates, writes, backend launch, and host mutation disabled.",
	}
	checks := currentMainlineOwnershipAuditChecks(preview, ready)
	preview.Checks = checks
	preview.CheckIDs = currentMainlineOwnershipAuditCheckIDs(checks)
	preview.Counts = countCurrentMainlineOwnershipAuditChecks(checks)
	if preview.Counts.Blocked == 0 {
		preview.AuditDecision = "current-mainline-ownership-audit-ready-autonomous-mainline"
	}
	if err := validateNoBackendTerms(preview, "current mainline ownership audit preview"); err != nil {
		return CurrentMainlineOwnershipAuditPreview{}, err
	}
	return preview, nil
}

func currentMainlineOwnershipAuditChecks(preview CurrentMainlineOwnershipAuditPreview, ready bool) []CurrentMainlineOwnershipAuditCheck {
	return []CurrentMainlineOwnershipAuditCheck{
		currentMainlineOwnershipAuditCheck("current-mainline-document-present", productionAuthorizationPassBlocked(preview.CurrentMainlineDocumentPresent && preview.CodexOwnedMainline), "The Xnix current mainline document is present and owned by this repository workflow."),
		currentMainlineOwnershipAuditCheck("external-agent-dependency-removed", productionAuthorizationPassBlocked(!preview.ExternalAgentDependencyRequired && preview.HistoricalExternalAgentDocumentsAllowed), "New mainline work does not require an external-agent dispatch source."),
		currentMainlineOwnershipAuditCheck("kde-flagship-only", productionAuthorizationPassBlocked(preview.KDEFlagshipOnly && !preview.GNOMEFirstReleaseTarget && !preview.XFCEFirstReleaseTarget), "KDE remains the only first-release flagship desktop."),
		currentMainlineOwnershipAuditCheck("runtime-ownership-language-split", productionAuthorizationPassBlocked(preview.GoRuntimeOwned && preview.RubyTestHarnessOnly && preview.CLowLevelOnly && preview.KDEPluginsPresentationOnly && preview.RuntimeOwnsCompatibilityDecisions), "Runtime ownership and language boundaries are explicit."),
		currentMainlineOwnershipAuditCheck("seven-kde-entrypoints-present", productionAuthorizationPassBlocked(preview.SevenKDEEntrypointsPresent), "The first-release KDE entrypoint list is present."),
		currentMainlineOwnershipAuditCheck("safe-next-task-present", productionAuthorizationPassBlocked(preview.SafeNextTaskPresent), "The next Codex-owned mainline task is present."),
		currentMainlineOwnershipAuditCheck("full-gate-authorization-required", productionAuthorizationPassBlocked(preview.FullGateRequiresAuthorization && !preview.DockerExecuted && !preview.QEMUExecuted && !preview.NetworkFetchEnabled && !preview.PackageManagerInvoked), "Full gates remain authorization-only."),
		currentMainlineOwnershipAuditCheck("unsafe-runtime-and-host-gates-closed", productionAuthorizationPassBlocked(ready && !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.KDEConfigurationWritten && !preview.PortalCallExecuted && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.StateRootPathExposed && !preview.FilePathsExposed && !preview.FileContentRead && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed), "Production ownership, writes, desktop side effects, backend launch, unsafe data exposure, and host mutation remain disabled."),
	}
}

func currentMainlineOwnershipAuditCheck(id string, status string, summary string) CurrentMainlineOwnershipAuditCheck {
	return CurrentMainlineOwnershipAuditCheck{ID: id, Status: status, Summary: summary}
}

func currentMainlineOwnershipAuditCheckIDs(checks []CurrentMainlineOwnershipAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countCurrentMainlineOwnershipAuditChecks(checks []CurrentMainlineOwnershipAuditCheck) CurrentMainlineOwnershipAuditCounts {
	counts := CurrentMainlineOwnershipAuditCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "pending":
			counts.Pending++
		default:
			counts.Blocked++
		}
	}
	return counts
}
