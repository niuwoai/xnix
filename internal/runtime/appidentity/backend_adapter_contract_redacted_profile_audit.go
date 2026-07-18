package appidentity

type BackendAdapterRedactedProfileAuditPreview struct {
	Version                       string                                    `json:"version"`
	SchemaVersion                 string                                    `json:"schema_version"`
	RequestType                   string                                    `json:"request_type"`
	AuditType                     string                                    `json:"audit_type"`
	Source                        string                                    `json:"source"`
	RuntimeMethod                 string                                    `json:"runtime_method"`
	ReadMethod                    string                                    `json:"read_method"`
	SubjectRequestType            string                                    `json:"subject_request_type"`
	SubjectReadMethod             string                                    `json:"subject_read_method"`
	OwnerRouteMethod              string                                    `json:"owner_route_method"`
	RouteDecision                 string                                    `json:"route_decision"`
	RouteDecisionReason           string                                    `json:"route_decision_reason"`
	Profiles                      []BackendAdapterRedactedProfile           `json:"profiles"`
	ProfileIDs                    []string                                  `json:"profile_ids"`
	Checks                        []BackendAdapterRedactedProfileAuditCheck `json:"checks"`
	CheckIDs                      []string                                  `json:"check_ids"`
	Counts                        BackendAdapterRedactedProfileAuditCounts  `json:"counts"`
	ContractProfileCount          int                                       `json:"contract_profile_count"`
	FullContractConsumed          bool                                      `json:"full_contract_consumed"`
	KDEFacingProjectionConsumed   bool                                      `json:"kde_facing_projection_consumed"`
	InternalAdapterIDsRedacted    bool                                      `json:"internal_adapter_ids_redacted"`
	InternalProfilePathsRedacted  bool                                      `json:"internal_profile_paths_redacted"`
	OwnerLocalRouteCandidateReady bool                                      `json:"owner_local_route_candidate_ready"`
	FullContractFixtureLocal      bool                                      `json:"full_contract_fixture_local"`
	ProductionDBusExposureReady   bool                                      `json:"production_dbus_exposure_ready"`
	CallerStateRootRequired       bool                                      `json:"caller_state_root_required"`
	CLIProjectRootOnly            bool                                      `json:"cli_project_root_only"`
	RuntimeOwned                  bool                                      `json:"runtime_owned"`
	GoRuntimeBacked               bool                                      `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                                      `json:"kde_policy_owner"`
	SystemServiceStarted          bool                                      `json:"system_service_started"`
	SessionBusClaimed             bool                                      `json:"session_bus_claimed"`
	ProductionBusClaimed          bool                                      `json:"production_bus_claimed"`
	WriteMethodsEnabled           bool                                      `json:"write_methods_enabled"`
	RuntimeWritesEnabled          bool                                      `json:"runtime_writes_enabled"`
	AdapterInvocationEnabled      bool                                      `json:"adapter_invocation_enabled"`
	BackendInstallEnabled         bool                                      `json:"backend_install_enabled"`
	BackendDownloadEnabled        bool                                      `json:"backend_download_enabled"`
	BackendLaunchEnabled          bool                                      `json:"backend_launch_enabled"`
	BackendProcessStarted         bool                                      `json:"backend_process_started"`
	CommandMaterialized           bool                                      `json:"command_materialized"`
	ExecutablePathResolved        bool                                      `json:"executable_path_resolved"`
	NetworkRequired               bool                                      `json:"network_required"`
	HostRootModified              bool                                      `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                                      `json:"privileged_container_required"`
	StateRootPathExposed          bool                                      `json:"state_root_path_exposed"`
	RawCommandExposed             bool                                      `json:"raw_command_exposed"`
	RawExecutableExposed          bool                                      `json:"raw_executable_exposed"`
	BackendDetailsExposed         bool                                      `json:"backend_details_exposed"`
	BlockedActions                []string                                  `json:"blocked_actions"`
	NextRequirements              []string                                  `json:"next_requirements"`
	DesktopSafeSummary            string                                    `json:"desktop_safe_summary"`
}

type BackendAdapterRedactedProfile struct {
	ID                          string `json:"id"`
	Label                       string `json:"label"`
	Summary                     string `json:"summary"`
	ContractStatus              string `json:"contract_status"`
	RecommendedUserMode         string `json:"recommended_user_mode"`
	UserVisibleState            string `json:"user_visible_state"`
	SourceProfileMatched        bool   `json:"source_profile_matched"`
	RuntimeOwned                bool   `json:"runtime_owned"`
	KDEPolicyOwner              bool   `json:"kde_policy_owner"`
	InternalAdapterIDRedacted   bool   `json:"internal_adapter_id_redacted"`
	InternalProfilePathRedacted bool   `json:"internal_profile_path_redacted"`
	BackendDetailsExposed       bool   `json:"backend_details_exposed"`
	AdapterInvocationEnabled    bool   `json:"adapter_invocation_enabled"`
	LaunchEnabled               bool   `json:"launch_enabled"`
	HostRootModified            bool   `json:"host_root_modified"`
}

type BackendAdapterRedactedProfileAuditCheck struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type BackendAdapterRedactedProfileAuditCounts struct {
	Total                 int `json:"total"`
	Passed                int `json:"passed"`
	Blocked               int `json:"blocked"`
	Pending               int `json:"pending"`
	RedactedProfiles      int `json:"redacted_profiles"`
	UserVisibleProfiles   int `json:"user_visible_profiles"`
	EnabledInvocations    int `json:"enabled_invocations"`
	EnabledLaunches       int `json:"enabled_launches"`
	ExposedBackendDetails int `json:"exposed_backend_details"`
}

func NewBackendAdapterRedactedProfileAuditPreview(root string) (BackendAdapterRedactedProfileAuditPreview, error) {
	contract, err := NewBackendAdapterContractPreview(root)
	if err != nil {
		return BackendAdapterRedactedProfileAuditPreview{}, err
	}
	sources := backendAdapterRedactedProfileAuditSources(root)
	profiles := backendAdapterRedactedProfiles(contract.KDEFacingProfiles)
	preview := BackendAdapterRedactedProfileAuditPreview{
		Version:                      contract.Version,
		SchemaVersion:                "xnix.runtime.backend_adapter_redacted_profile_audit.v1",
		RequestType:                  "backend-adapter-redacted-profile-audit-preview",
		AuditType:                    "owner-local-redacted-adapter-profile-audit",
		Source:                       "backend-adapter-contract-preview+kde-facing-profile-projection+owner-local-read-dispatch",
		RuntimeMethod:                "GetBackendAdapterProfileAudit",
		ReadMethod:                   "GetBackendAdapterProfileAuditPreview",
		SubjectRequestType:           contract.RequestType,
		SubjectReadMethod:            contract.ReadMethod,
		OwnerRouteMethod:             "GetBackendAdapterProfileAudit",
		RouteDecision:                "redacted-profile-route-ready",
		RouteDecisionReason:          "Owner-local routing can expose redacted user-safe profile audit states while the full adapter contract remains fixture-local.",
		Profiles:                     profiles,
		ProfileIDs:                   backendAdapterRedactedProfileIDs(profiles),
		ContractProfileCount:         contract.Counts.KDEFacingProfiles,
		FullContractConsumed:         contract.RequestType == "backend-adapter-contract-preview",
		KDEFacingProjectionConsumed:  len(contract.KDEFacingProfiles) == len(profiles) && len(profiles) > 0,
		InternalAdapterIDsRedacted:   true,
		InternalProfilePathsRedacted: true,
		FullContractFixtureLocal:     true,
		ProductionDBusExposureReady:  false,
		CallerStateRootRequired:      false,
		CLIProjectRootOnly:           true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		SystemServiceStarted:         false,
		SessionBusClaimed:            false,
		ProductionBusClaimed:         false,
		WriteMethodsEnabled:          false,
		RuntimeWritesEnabled:         false,
		AdapterInvocationEnabled:     false,
		BackendInstallEnabled:        false,
		BackendDownloadEnabled:       false,
		BackendLaunchEnabled:         false,
		BackendProcessStarted:        false,
		CommandMaterialized:          false,
		ExecutablePathResolved:       false,
		NetworkRequired:              false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		StateRootPathExposed:         false,
		RawCommandExposed:            false,
		RawExecutableExposed:         false,
		BackendDetailsExposed:        false,
		BlockedActions:               backendAdapterRedactedProfileAuditBlockedActions(),
		NextRequirements:             backendAdapterRedactedProfileAuditNextRequirements(),
		DesktopSafeSummary:           "Runtime exposes only redacted user-safe compatibility profile audit states through owner-local read dispatch while keeping the full adapter contract fixture-local and all launch or write paths disabled.",
	}
	preview.OwnerLocalRouteCandidateReady = backendAdapterRedactedProfileAuditHasOwnerRoute(sources.OwnerDispatch) &&
		backendAdapterRedactedProfileAuditHasCLICommand(sources.GoCLI) &&
		!backendAdapterRedactedProfileAuditHasProductionDBusMethod(sources.DBusContract)
	checks := backendAdapterRedactedProfileAuditChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = backendAdapterRedactedProfileAuditCheckIDs(checks)
	preview.Counts = countBackendAdapterRedactedProfileAuditChecks(checks, profiles)
	if err := validateNoBackendTerms(preview, "backend adapter redacted profile audit preview"); err != nil {
		return BackendAdapterRedactedProfileAuditPreview{}, err
	}
	return preview, nil
}

type backendAdapterRedactedProfileAuditSourceSet struct {
	GoCLI         string
	OwnerDispatch string
	DBusContract  string
}

func backendAdapterRedactedProfileAuditSources(root string) backendAdapterRedactedProfileAuditSourceSet {
	return backendAdapterRedactedProfileAuditSourceSet{
		GoCLI:         readRuntimeMethodParitySources(root, []string{"cmd/xnix-runtime-go/main.go", "cmd/xnix-runtime-go/backend_adapter_contract_commands.go"}),
		OwnerDispatch: readRuntimeMethodParitySources(root, []string{"internal/runtime/owner/dispatch.go"}),
		DBusContract:  readRuntimeMethodParitySources(root, []string{"runtime/dbus/org.xnix.Compatibility1.xml"}),
	}
}

func backendAdapterRedactedProfiles(source []BackendAdapterProfile) []BackendAdapterRedactedProfile {
	definitions := []struct {
		id          string
		label       string
		mode        string
		state       string
		summary     string
		sourceIndex int
	}{
		{"automatic", "Automatic", "automatic", "planned", "Runtime can present this as the default compatibility choice after required reviews pass.", 0},
		{"performance-priority", "Performance priority", "performance-priority", "planned", "Runtime can present this as a performance-oriented compatibility choice after required reviews pass.", 1},
		{"compatibility-priority", "Compatibility priority", "compatibility-priority", "planned", "Runtime can present this as a compatibility-oriented choice after required reviews pass.", 2},
	}
	profiles := make([]BackendAdapterRedactedProfile, 0, len(definitions))
	for _, definition := range definitions {
		matched := definition.sourceIndex < len(source) && source[definition.sourceIndex].ContractStatus == "noop-contract"
		profiles = append(profiles, BackendAdapterRedactedProfile{
			ID:                          definition.id,
			Label:                       definition.label,
			Summary:                     definition.summary,
			ContractStatus:              "redacted-noop-profile",
			RecommendedUserMode:         definition.mode,
			UserVisibleState:            definition.state,
			SourceProfileMatched:        matched,
			RuntimeOwned:                true,
			KDEPolicyOwner:              false,
			InternalAdapterIDRedacted:   true,
			InternalProfilePathRedacted: true,
			BackendDetailsExposed:       false,
			AdapterInvocationEnabled:    false,
			LaunchEnabled:               false,
			HostRootModified:            false,
		})
	}
	return profiles
}

func backendAdapterRedactedProfileIDs(profiles []BackendAdapterRedactedProfile) []string {
	ids := make([]string, 0, len(profiles))
	for _, profile := range profiles {
		ids = append(ids, profile.ID)
	}
	return ids
}

func backendAdapterRedactedProfileAuditChecks(preview BackendAdapterRedactedProfileAuditPreview) []BackendAdapterRedactedProfileAuditCheck {
	return []BackendAdapterRedactedProfileAuditCheck{
		backendAdapterRedactedProfileAuditCheck("contract-profile-projection-consumed", backendAdapterRedactedProfileAuditPassBlocked(preview.FullContractConsumed && preview.KDEFacingProjectionConsumed && preview.ContractProfileCount == len(preview.Profiles)), "The redacted audit consumes the existing KDE-facing adapter profile projection."),
		backendAdapterRedactedProfileAuditCheck("internal-adapter-ids-redacted", backendAdapterRedactedProfileAuditPassBlocked(preview.InternalAdapterIDsRedacted && preview.InternalProfilePathsRedacted && backendAdapterRedactedProfilesHideDetails(preview.Profiles)), "Internal adapter identifiers and profile paths are omitted from owner-facing output."),
		backendAdapterRedactedProfileAuditCheck("owner-local-route-candidate", backendAdapterRedactedProfileAuditPassBlocked(preview.OwnerLocalRouteCandidateReady), "Owner read dispatch exposes only the redacted profile audit route."),
		backendAdapterRedactedProfileAuditCheck("full-contract-fixture-local", backendAdapterRedactedProfileAuditPassBlocked(preview.FullContractFixtureLocal && preview.SubjectReadMethod == "GetBackendAdapterContractPreview"), "The full adapter contract remains fixture-local and is not routed as the owner method."),
		backendAdapterRedactedProfileAuditCheck("production-dbus-not-claimed", backendAdapterRedactedProfileAuditPassBlocked(!preview.ProductionDBusExposureReady && !preview.ProductionBusClaimed && !preview.SessionBusClaimed && !preview.SystemServiceStarted), "The redacted route is not exposed through production D-Bus or a started service."),
		backendAdapterRedactedProfileAuditCheck("no-caller-state-root", backendAdapterRedactedProfileAuditPassBlocked(!preview.CallerStateRootRequired && !preview.StateRootPathExposed), "Owner-local dispatch does not require caller-provided state roots or expose paths."),
		backendAdapterRedactedProfileAuditCheck("unsafe-gates-closed", backendAdapterRedactedProfileAuditPassBlocked(!preview.WriteMethodsEnabled && !preview.RuntimeWritesEnabled && !preview.AdapterInvocationEnabled && !preview.BackendInstallEnabled && !preview.BackendDownloadEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.CommandMaterialized && !preview.ExecutablePathResolved && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.RawCommandExposed && !preview.RawExecutableExposed && !preview.BackendDetailsExposed), "Adapter invocation, installation, download, launch, command materialization, executable resolution, networking, host mutation, and raw output remain disabled."),
	}
}

func backendAdapterRedactedProfileAuditCheck(id string, status string, summary string) BackendAdapterRedactedProfileAuditCheck {
	return BackendAdapterRedactedProfileAuditCheck{ID: id, Status: status, Summary: summary}
}

func backendAdapterRedactedProfileAuditPassBlocked(passed bool) string {
	if passed {
		return "pass"
	}
	return "blocked"
}

func backendAdapterRedactedProfileAuditCheckIDs(checks []BackendAdapterRedactedProfileAuditCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countBackendAdapterRedactedProfileAuditChecks(checks []BackendAdapterRedactedProfileAuditCheck, profiles []BackendAdapterRedactedProfile) BackendAdapterRedactedProfileAuditCounts {
	counts := BackendAdapterRedactedProfileAuditCounts{Total: len(checks), UserVisibleProfiles: len(profiles)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "blocked":
			counts.Blocked++
		case "pending":
			counts.Pending++
		}
	}
	for _, profile := range profiles {
		if profile.InternalAdapterIDRedacted && profile.InternalProfilePathRedacted {
			counts.RedactedProfiles++
		}
		if profile.AdapterInvocationEnabled {
			counts.EnabledInvocations++
		}
		if profile.LaunchEnabled {
			counts.EnabledLaunches++
		}
		if profile.BackendDetailsExposed {
			counts.ExposedBackendDetails++
		}
	}
	return counts
}

func backendAdapterRedactedProfilesHideDetails(profiles []BackendAdapterRedactedProfile) bool {
	if len(profiles) == 0 {
		return false
	}
	for _, profile := range profiles {
		if !profile.SourceProfileMatched || !profile.InternalAdapterIDRedacted || !profile.InternalProfilePathRedacted || profile.BackendDetailsExposed || profile.AdapterInvocationEnabled || profile.LaunchEnabled || profile.HostRootModified {
			return false
		}
	}
	return true
}

func backendAdapterRedactedProfileAuditHasCLICommand(source string) bool {
	return textHasAll(source, []string{"backend-adapter-redacted-profile-audit-preview", "runBackendAdapterRedactedProfileAuditPreview"})
}

func backendAdapterRedactedProfileAuditHasOwnerRoute(source string) bool {
	return textHasAll(source, []string{"GetBackendAdapterProfileAudit", "NewBackendAdapterRedactedProfileAuditPreview"})
}

func backendAdapterRedactedProfileAuditHasProductionDBusMethod(source string) bool {
	return textHasAll(source, []string{"GetBackendAdapterProfileAudit", "backend-adapter-redacted-profile-audit-preview"})
}

func backendAdapterRedactedProfileAuditBlockedActions() []string {
	return []string{
		"route the full internal adapter contract through owner dispatch",
		"expose internal adapter identifiers to KDE-facing output",
		"expose profile storage paths or raw commands",
		"enable adapter invocation from profile audit",
		"launch compatibility backend from profile audit",
		"mutate host root during profile audit",
	}
}

func backendAdapterRedactedProfileAuditNextRequirements() []string {
	return []string{
		"Cover the redacted profile audit route in restricted owner smoke evidence.",
		"Keep the full adapter contract fixture-local until production recipe trust and write gates mature.",
		"Require explicit launch preflight and review receipts before any adapter invocation.",
	}
}
