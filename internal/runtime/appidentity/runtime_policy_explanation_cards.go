package appidentity

import "fmt"

type RuntimePolicyExplanationCardsOptions struct {
	Environment     string
	RuntimeRoot     string
	StateRoot       string
	PortalOperation string
	SnapshotReason  string
	Issue           string
	TestType        string
}

type RuntimePolicyExplanationCardsPreview struct {
	Version                     string                         `json:"version"`
	SchemaVersion               string                         `json:"schema_version"`
	RequestType                 string                         `json:"request_type"`
	CardDeckType                string                         `json:"card_deck_type"`
	Source                      string                         `json:"source"`
	Desktop                     string                         `json:"desktop"`
	RuntimeMethod               string                         `json:"runtime_method"`
	ReadMethod                  string                         `json:"read_method"`
	Application                 RuntimePolicyExplanationApp    `json:"application"`
	Cards                       []RuntimePolicyExplanationCard `json:"cards"`
	CardIDs                     []string                       `json:"card_ids"`
	CardCount                   int                            `json:"card_count"`
	Counts                      RuntimePolicyExplanationCounts `json:"counts"`
	RuntimeOwned                bool                           `json:"runtime_owned"`
	GoRuntimeBacked             bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                           `json:"kde_policy_owner"`
	UserVisible                 bool                           `json:"user_visible"`
	ReviewOnly                  bool                           `json:"review_only"`
	CardsPersisted              bool                           `json:"cards_persisted"`
	ActionEnablementChanged     bool                           `json:"action_enablement_changed"`
	RequestObjectsCreated       bool                           `json:"request_objects_created"`
	PermissionGrantsCreated     bool                           `json:"permission_grants_created"`
	SettingsPersisted           bool                           `json:"settings_persisted"`
	AIProviderCalled            bool                           `json:"ai_provider_called"`
	AIProviderCallEnabled       bool                           `json:"ai_provider_call_enabled"`
	BackendProcessStarted       bool                           `json:"backend_process_started"`
	LaunchEnabled               bool                           `json:"launch_enabled"`
	ExecutionStarted            bool                           `json:"execution_started"`
	HostRootModified            bool                           `json:"host_root_modified"`
	NetworkRequired             bool                           `json:"network_required"`
	PrivilegedContainerRequired bool                           `json:"privileged_container_required"`
	StateRootPathExposed        bool                           `json:"state_root_path_exposed"`
	RawExecutableExposed        bool                           `json:"raw_executable_exposed"`
	RawCommandExposed           bool                           `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                           `json:"backend_details_exposed"`
	BlockedActions              []string                       `json:"blocked_actions"`
	DesktopSafeSummary          string                         `json:"desktop_safe_summary"`
}

type RuntimePolicyExplanationApp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DesktopFile string `json:"desktop_file"`
	RuntimeMode string `json:"runtime_mode"`
}

type RuntimePolicyExplanationCard struct {
	ID                    string   `json:"id"`
	Topic                 string   `json:"topic"`
	State                 string   `json:"state"`
	Severity              string   `json:"severity"`
	Audience              string   `json:"audience"`
	RelatedEvidenceIDs    []string `json:"related_evidence_ids"`
	UserFacingSummary     string   `json:"user_facing_summary"`
	TechnicalSummary      string   `json:"technical_summary"`
	NextSafeReadOnlyCheck string   `json:"next_safe_read_only_check"`
	BlockerReasons        []string `json:"blocker_reasons"`
	BlockerReasonCount    int      `json:"blocker_reason_count"`
	ActionEnabled         bool     `json:"action_enabled"`
	CardPersisted         bool     `json:"card_persisted"`
	RequestObjectCreated  bool     `json:"request_object_created"`
	PermissionGranted     bool     `json:"permission_granted"`
	SettingsPersisted     bool     `json:"settings_persisted"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	BackendProcessStarted bool     `json:"backend_process_started"`
	LaunchEnabled         bool     `json:"launch_enabled"`
	ExecutionStarted      bool     `json:"execution_started"`
	HostRootModified      bool     `json:"host_root_modified"`
	StateRootPathExposed  bool     `json:"state_root_path_exposed"`
	RawCommandExposed     bool     `json:"raw_command_exposed"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type RuntimePolicyExplanationCounts struct {
	Total             int `json:"total"`
	Blocked           int `json:"blocked"`
	ReviewOnly        int `json:"review_only"`
	MissingEvidence   int `json:"missing_evidence"`
	NotYetImplemented int `json:"not_yet_implemented"`
	Ready             int `json:"ready"`
}

func (plan Plan) RuntimePolicyExplanationCardsPreview(options RuntimePolicyExplanationCardsOptions) (RuntimePolicyExplanationCardsPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return RuntimePolicyExplanationCardsPreview{}, err
	}
	if options.Environment == "" {
		options.Environment = "development"
	}
	if options.RuntimeRoot == "" {
		options.RuntimeRoot = "."
	}
	if options.PortalOperation == "" {
		options.PortalOperation = "file-open"
	}
	if options.SnapshotReason == "" {
		options.SnapshotReason = "before-repair"
	}
	if options.Issue == "" {
		options.Issue = "portal-approval-required"
	}
	if options.TestType == "" {
		options.TestType = "smoke"
	}

	readiness, err := plan.ApplicationReadinessPreview(ApplicationReadinessOptions{
		Environment:     options.Environment,
		RuntimeRoot:     options.RuntimeRoot,
		StateRoot:       options.StateRoot,
		PortalOperation: options.PortalOperation,
		SnapshotReason:  options.SnapshotReason,
		WriteMethod:     "Launch",
	})
	if err != nil {
		return RuntimePolicyExplanationCardsPreview{}, err
	}
	repair, err := plan.RepairPlanPreview(options.Issue)
	if err != nil {
		return RuntimePolicyExplanationCardsPreview{}, err
	}
	safety := NewDesktopSafetyPolicyPreview()
	cards := runtimePolicyExplanationCards(readiness, repair, safety)
	counts := countRuntimePolicyExplanationCards(cards)
	preview := RuntimePolicyExplanationCardsPreview{
		Version:       readiness.Version,
		SchemaVersion: "xnix.runtime.policy_explanation_cards.v1",
		RequestType:   "runtime-policy-explanation-cards-preview",
		CardDeckType:  "kde-runtime-policy-explanation-card-deck",
		Source:        "application-readiness-preview+runtime-write-gate-preview+desktop-safety-policy-preview+repair-plan-preview",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetRuntimePolicyExplanationCards",
		ReadMethod:    "GetRuntimePolicyExplanationCardsPreview",
		Application: RuntimePolicyExplanationApp{
			ID:          readiness.Application.ID,
			Name:        readiness.Application.Name,
			Icon:        readiness.Application.Icon,
			DesktopFile: readiness.Application.DesktopFile,
			RuntimeMode: readiness.Application.RuntimeMode,
		},
		Cards:                       cards,
		CardIDs:                     runtimePolicyExplanationCardIDs(cards),
		CardCount:                   len(cards),
		Counts:                      counts,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		CardsPersisted:              false,
		ActionEnablementChanged:     false,
		RequestObjectsCreated:       false,
		PermissionGrantsCreated:     false,
		SettingsPersisted:           false,
		AIProviderCalled:            false,
		AIProviderCallEnabled:       false,
		BackendProcessStarted:       false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawExecutableExposed:        false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"enable actions from policy explanation cards",
			"persist policy explanation cards",
			"create request objects from policy explanation cards",
			"grant desktop permissions from policy explanation cards",
			"persist settings from policy explanation cards",
			"call AI providers from policy explanation cards",
			"start compatibility services from policy explanation cards",
			"launch applications from policy explanation cards",
			"mutate host root during policy explanation",
		},
		DesktopSafeSummary: "KDE can show consistent Runtime policy explanations, but the card deck is review-only and cannot enable actions, create requests, grant permissions, call AI providers, start services, launch applications, or mutate the host root.",
	}
	if err := validateRuntimePolicyExplanationCards(preview); err != nil {
		return RuntimePolicyExplanationCardsPreview{}, err
	}
	if err := validateNoBackendTerms(preview, "Runtime policy explanation cards preview"); err != nil {
		return RuntimePolicyExplanationCardsPreview{}, err
	}
	return preview, nil
}

func runtimePolicyExplanationCards(readiness ApplicationReadinessPreview, repair RepairPlanPreview, safety DesktopSafetyPolicyPreview) []RuntimePolicyExplanationCard {
	return []RuntimePolicyExplanationCard{
		runtimePolicyExplanationCard("install", "install", runtimePolicyStateFromInstall(readiness), "warning", "user", []string{"compatibility-install-preview", "recipe-trust", "artifact-stage-receipt"}, "Installation needs verified recipe and local artifact evidence before it can continue.", "Install policy is derived from compatibility install readiness and recipe trust evidence.", "compatibility-install-preview", runtimePolicyInstallReasons(readiness)),
		runtimePolicyExplanationCard("launch", "launch", "blocked", "critical", "user", []string{"runtime-write-gate-preview", "runtime-write-gate"}, "Launch is visible in KDE but remains blocked until Runtime gates pass.", "Launch is blocked by the Runtime write gate and cannot be enabled by KDE.", "runtime-write-gate-preview", []string{readiness.WriteGateDecision}),
		runtimePolicyExplanationCard("execution", "execution", "blocked", "critical", "reviewer", []string{"execution-readiness-preview", "execution-readiness"}, "Execution is waiting for Runtime readiness evidence.", "Execution readiness must pass launch, resource, service, and write-gate checks.", "execution-readiness-preview", runtimePolicyNodeReasons(readiness, "execution-readiness")),
		runtimePolicyExplanationCard("portal-permission", "portal-permission", "review-only", "warning", "user", []string{"portal-access-policy-preview", "portal-review"}, "Desktop resource access needs user-mediated review.", "Portal evidence is required before sensitive desktop resource access can be trusted.", "portal-access-policy-preview", runtimePolicyNodeReasons(readiness, "portal-review")),
		runtimePolicyExplanationCard("snapshot", "snapshot", "missing-evidence", "warning", "user", []string{"snapshot-plan-preview", "snapshot-baseline"}, "A restore point is required before risky compatibility changes.", "Snapshot evidence is missing and this preview does not create restore points.", "snapshot-plan-preview", runtimePolicyNodeReasons(readiness, "snapshot-baseline")),
		runtimePolicyExplanationCard("diagnostics", "diagnostics", "review-only", "info", "user", []string{"diagnostics-preview", "ai-diagnostic-input-preview"}, "Diagnostics can explain safe metadata without reading private file contents.", "Diagnostics remain safe-for-AI metadata only and do not call an AI provider.", "diagnostics-preview", []string{"diagnostic history review is recommended"}),
		runtimePolicyExplanationCard("repair", "repair", runtimePolicyRepairState(repair), runtimePolicyRepairSeverity(repair), "user", []string{"repair-plan-preview", "ai-repair-approval-gate-preview"}, repair.DesktopSafeSummary, "Repair policy is review-first and execution remains disabled from explanation cards.", "repair-plan-preview", []string{repair.Issue}),
		runtimePolicyExplanationCard("settings", "settings", "review-only", "info", "user", []string{"settings-preview", "settings-change-preview", "desktop-safety-policy-preview"}, "Settings changes need Runtime review before they can be saved.", "Settings persistence is disabled while the desktop safety policy keeps user-facing controls backend-safe.", "settings-change-preview", []string{"settings persistence is disabled"}),
		runtimePolicyExplanationCard("desktop-activation", "desktop-activation", "missing-evidence", "warning", "reviewer", []string{"desktop-activation-transaction-preview", "desktop-activation-status-preview"}, "Desktop integration is preview-only until activation evidence is complete.", "Activation writes stay disabled until digest, transaction, and rollback evidence are present.", "desktop-activation-transaction-preview", []string{"desktop activation transaction is not committed"}),
		runtimePolicyExplanationCard("backend-readiness", "backend-readiness", "missing-evidence", "warning", "reviewer", []string{"backend-lifecycle-preview", "backend-lifecycle"}, "Compatibility service setup still needs Runtime evidence.", "Backend lifecycle evidence is required but implementation details remain hidden from KDE.", "backend-lifecycle-preview", runtimePolicyNodeReasons(readiness, "backend-lifecycle")),
		runtimePolicyExplanationCard("unsupported-production-route", "unsupported-production-route", "not-yet-implemented", "critical", "reviewer", []string{"runtime-owner-readiness-preview", "runtime-write-gate-preview"}, "Production write routes are not available yet.", "Production D-Bus ownership and write-method enablement remain deliberately unavailable.", "runtime-owner-readiness-preview", []string{"production write routes are not implemented", safety.DesktopSafeSummary}),
	}
}

func runtimePolicyExplanationCard(id string, topic string, state string, severity string, audience string, evidenceIDs []string, userSummary string, technicalSummary string, nextCheck string, reasons []string) RuntimePolicyExplanationCard {
	dedupedReasons := uniqueSortedStrings(reasons)
	return RuntimePolicyExplanationCard{
		ID:                    id,
		Topic:                 topic,
		State:                 state,
		Severity:              severity,
		Audience:              audience,
		RelatedEvidenceIDs:    uniqueSortedStrings(evidenceIDs),
		UserFacingSummary:     userSummary,
		TechnicalSummary:      technicalSummary,
		NextSafeReadOnlyCheck: nextCheck,
		BlockerReasons:        dedupedReasons,
		BlockerReasonCount:    len(dedupedReasons),
		ActionEnabled:         false,
		CardPersisted:         false,
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		SettingsPersisted:     false,
		AIProviderCalled:      false,
		BackendProcessStarted: false,
		LaunchEnabled:         false,
		ExecutionStarted:      false,
		HostRootModified:      false,
		StateRootPathExposed:  false,
		RawCommandExposed:     false,
		BackendDetailsExposed: false,
	}
}

func runtimePolicyStateFromInstall(readiness ApplicationReadinessPreview) string {
	if readiness.InstallReadiness.RecipeInstallDecision == "block" {
		return "blocked"
	}
	if !readiness.InstallReadiness.ArtifactStageReceiptReady || !readiness.InstallReadiness.RequiredArtifactsStaged {
		return "missing-evidence"
	}
	return "review-only"
}

func runtimePolicyRepairState(repair RepairPlanPreview) string {
	if repair.RepairExecuted {
		return "ready"
	}
	if repair.UserApprovalRequired {
		return "review-only"
	}
	return "not-yet-implemented"
}

func runtimePolicyRepairSeverity(repair RepairPlanPreview) string {
	if repair.Severity == "critical" {
		return "critical"
	}
	if repair.Severity == "warning" {
		return "warning"
	}
	return "info"
}

func runtimePolicyInstallReasons(readiness ApplicationReadinessPreview) []string {
	var reasons []string
	reasons = append(reasons, readiness.InstallReadiness.RecipeTrustBlockingReasons...)
	reasons = append(reasons, readiness.InstallReadiness.ArtifactStageBlockingReasons...)
	if !readiness.InstallReadiness.ArtifactStageReceiptReady {
		reasons = append(reasons, "artifact staging receipt is missing")
	}
	if !readiness.InstallReadiness.RequiredArtifactsStaged {
		reasons = append(reasons, "required artifacts are not staged")
	}
	if !readiness.InstallReadiness.RecipeInstallAllowed {
		reasons = append(reasons, "recipe install gate is not allowing this environment")
	}
	return reasons
}

func runtimePolicyNodeReasons(readiness ApplicationReadinessPreview, nodeID string) []string {
	for _, node := range readiness.Nodes {
		if node.ID == nodeID {
			return node.BlockingReasons
		}
	}
	return []string{"required Runtime evidence is missing"}
}

func countRuntimePolicyExplanationCards(cards []RuntimePolicyExplanationCard) RuntimePolicyExplanationCounts {
	counts := RuntimePolicyExplanationCounts{Total: len(cards)}
	for _, card := range cards {
		switch card.State {
		case "blocked":
			counts.Blocked++
		case "review-only":
			counts.ReviewOnly++
		case "missing-evidence":
			counts.MissingEvidence++
		case "not-yet-implemented":
			counts.NotYetImplemented++
		case "ready":
			counts.Ready++
		}
	}
	return counts
}

func runtimePolicyExplanationCardIDs(cards []RuntimePolicyExplanationCard) []string {
	ids := make([]string, 0, len(cards))
	for _, card := range cards {
		ids = append(ids, card.ID)
	}
	return ids
}

func validateRuntimePolicyExplanationCards(preview RuntimePolicyExplanationCardsPreview) error {
	for _, card := range preview.Cards {
		switch card.State {
		case "blocked", "review-only", "missing-evidence", "not-yet-implemented", "ready":
		default:
			return fmt.Errorf("unsupported policy explanation card state: %s", card.State)
		}
		if !singleLine(card.UserFacingSummary) || !singleLine(card.TechnicalSummary) || !singleLine(card.NextSafeReadOnlyCheck) {
			return fmt.Errorf("policy explanation card contains non-single-line text: %s", card.ID)
		}
	}
	return nil
}
