package appidentity

import (
	"errors"
	"sort"
)

type MultiApplicationInstallQueuePreview struct {
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	QueueType                   string                        `json:"queue_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	RuntimeMethod               string                        `json:"runtime_method"`
	ReadMethod                  string                        `json:"read_method"`
	Environment                 string                        `json:"environment"`
	Applications                []MultiApplicationInstallItem `json:"applications"`
	ApplicationIDs              []string                      `json:"application_ids"`
	ApplicationCount            int                           `json:"application_count"`
	Counts                      MultiApplicationInstallCounts `json:"counts"`
	QueueStatus                 string                        `json:"queue_status"`
	ReadyForReview              bool                          `json:"ready_for_review"`
	InstallReadyCount           int                           `json:"install_ready_count"`
	BlockedApplicationCount     int                           `json:"blocked_application_count"`
	ReviewRequiredCount         int                           `json:"review_required_count"`
	SharedBlockingReasons       []string                      `json:"shared_blocking_reasons"`
	RuntimeOwned                bool                          `json:"runtime_owned"`
	GoRuntimeBacked             bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                          `json:"kde_policy_owner"`
	UserVisible                 bool                          `json:"user_visible"`
	ReviewOnly                  bool                          `json:"review_only"`
	QueuePersisted              bool                          `json:"queue_persisted"`
	RequestObjectsCreated       bool                          `json:"request_objects_created"`
	ArtifactsStaged             bool                          `json:"artifacts_staged"`
	ArtifactsDownloaded         bool                          `json:"artifacts_downloaded"`
	NetworkRequestCreated       bool                          `json:"network_request_created"`
	HostPackageManagerInvoked   bool                          `json:"host_package_manager_invoked"`
	DesktopActivationStarted    bool                          `json:"desktop_activation_started"`
	InstallStarted              bool                          `json:"install_started"`
	BackendProcessStarted       bool                          `json:"backend_process_started"`
	LaunchEnabled               bool                          `json:"launch_enabled"`
	ExecutionStarted            bool                          `json:"execution_started"`
	SettingsPersisted           bool                          `json:"settings_persisted"`
	HostRootModified            bool                          `json:"host_root_modified"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	StateRootPathExposed        bool                          `json:"state_root_path_exposed"`
	RawExecutableExposed        bool                          `json:"raw_executable_exposed"`
	RawCommandExposed           bool                          `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                          `json:"backend_details_exposed"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type MultiApplicationInstallItem struct {
	Position                     int      `json:"position"`
	ApplicationID                string   `json:"application_id"`
	ApplicationName              string   `json:"application_name"`
	Icon                         string   `json:"icon"`
	DesktopFile                  string   `json:"desktop_file"`
	RuntimeMode                  string   `json:"runtime_mode"`
	QueueState                   string   `json:"queue_state"`
	Priority                     string   `json:"priority"`
	InstallState                 string   `json:"install_state"`
	InstallReady                 bool     `json:"install_ready"`
	DesktopActivationReady       bool     `json:"desktop_activation_ready"`
	RecipeInstallDecision        string   `json:"recipe_install_decision"`
	RecipeTrustBlockingReasons   []string `json:"recipe_trust_blocking_reasons"`
	ArtifactStageBlockingReasons []string `json:"artifact_stage_blocking_reasons"`
	PhaseIDs                     []string `json:"phase_ids"`
	BlockedReasons               []string `json:"blocked_reasons"`
	NextSafeReadOnlyCheck        string   `json:"next_safe_read_only_check"`
	UserReviewRequired           bool     `json:"user_review_required"`
	RequestObjectCreated         bool     `json:"request_object_created"`
	ArtifactStaged               bool     `json:"artifact_staged"`
	InstallStarted               bool     `json:"install_started"`
	BackendProcessStarted        bool     `json:"backend_process_started"`
	LaunchEnabled                bool     `json:"launch_enabled"`
	ExecutionStarted             bool     `json:"execution_started"`
	HostRootModified             bool     `json:"host_root_modified"`
	StateRootPathExposed         bool     `json:"state_root_path_exposed"`
	RawCommandExposed            bool     `json:"raw_command_exposed"`
	BackendDetailsExposed        bool     `json:"backend_details_exposed"`
}

type MultiApplicationInstallCounts struct {
	Total           int `json:"total"`
	Ready           int `json:"ready"`
	NeedsReview     int `json:"needs_review"`
	MissingEvidence int `json:"missing_evidence"`
	Blocked         int `json:"blocked"`
}

func NewMultiApplicationInstallQueuePreview(plans []Plan, environment string) (MultiApplicationInstallQueuePreview, error) {
	mode, err := normalizeActivationPreflightMode(environment)
	if err != nil {
		return MultiApplicationInstallQueuePreview{}, err
	}
	if len(plans) == 0 {
		return MultiApplicationInstallQueuePreview{}, errors.New("multi-application install queue requires at least one application")
	}

	items := make([]MultiApplicationInstallItem, 0, len(plans))
	var sharedBlocking []string
	for index, plan := range plans {
		if err := plan.ValidateSafeForDesktop(); err != nil {
			return MultiApplicationInstallQueuePreview{}, err
		}
		installPlan, err := plan.CompatibilityInstallPlanPreview(mode)
		if err != nil {
			return MultiApplicationInstallQueuePreview{}, err
		}
		item := multiApplicationInstallItem(index+1, installPlan)
		items = append(items, item)
		sharedBlocking = append(sharedBlocking, item.BlockedReasons...)
	}

	counts := countMultiApplicationInstallItems(items)
	appIDs := multiApplicationInstallApplicationIDs(items)
	preview := MultiApplicationInstallQueuePreview{
		SchemaVersion:               "xnix.runtime.multi_application_install_queue.v1",
		RequestType:                 "multi-application-install-queue-preview",
		QueueType:                   "review-only-multi-application-install-queue",
		Source:                      "registry+compatibility-install-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetMultiApplicationInstallQueue",
		ReadMethod:                  "GetMultiApplicationInstallQueuePreview",
		Environment:                 mode,
		Applications:                items,
		ApplicationIDs:              appIDs,
		ApplicationCount:            len(items),
		Counts:                      counts,
		QueueStatus:                 multiApplicationInstallQueueStatus(counts),
		ReadyForReview:              true,
		InstallReadyCount:           counts.Ready,
		BlockedApplicationCount:     counts.Blocked,
		ReviewRequiredCount:         counts.NeedsReview + counts.MissingEvidence + counts.Blocked,
		SharedBlockingReasons:       uniqueSortedStrings(sharedBlocking),
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		QueuePersisted:              false,
		RequestObjectsCreated:       false,
		ArtifactsStaged:             false,
		ArtifactsDownloaded:         false,
		NetworkRequestCreated:       false,
		HostPackageManagerInvoked:   false,
		DesktopActivationStarted:    false,
		InstallStarted:              false,
		BackendProcessStarted:       false,
		LaunchEnabled:               false,
		ExecutionStarted:            false,
		SettingsPersisted:           false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		StateRootPathExposed:        false,
		RawExecutableExposed:        false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"persist multi-application install queue from preview",
			"create install request objects from queue preview",
			"stage artifacts from queue preview",
			"download artifacts from queue preview",
			"invoke host package managers from queue preview",
			"start desktop activation from queue preview",
			"start compatibility services from queue preview",
			"launch applications from queue preview",
			"mutate host root during queue preview",
		},
		DesktopSafeSummary: "KDE can review a Runtime-owned multi-application install queue, but the queue is not persisted and no install, download, activation, launch, or host mutation is performed.",
	}
	if err := validateNoBackendTerms(preview, "multi-application install queue preview"); err != nil {
		return MultiApplicationInstallQueuePreview{}, err
	}
	return preview, nil
}

func multiApplicationInstallItem(position int, installPlan CompatibilityInstallPlanPreview) MultiApplicationInstallItem {
	blockedReasons := multiApplicationInstallBlockedReasons(installPlan)
	queueState := multiApplicationInstallItemState(installPlan, blockedReasons)
	return MultiApplicationInstallItem{
		Position:                     position,
		ApplicationID:                installPlan.Application.ID,
		ApplicationName:              installPlan.Application.Name,
		Icon:                         installPlan.Application.Icon,
		DesktopFile:                  installPlan.Application.DesktopFile,
		RuntimeMode:                  installPlan.Application.RequestedMode,
		QueueState:                   queueState,
		Priority:                     multiApplicationInstallPriority(queueState),
		InstallState:                 installPlan.InstallState,
		InstallReady:                 installPlan.InstallReady,
		DesktopActivationReady:       installPlan.DesktopActivationReady,
		RecipeInstallDecision:        installPlan.Readiness.RecipeInstallDecision,
		RecipeTrustBlockingReasons:   append([]string(nil), installPlan.Readiness.RecipeTrustBlockingReasons...),
		ArtifactStageBlockingReasons: append([]string(nil), installPlan.Readiness.ArtifactStageBlockingReasons...),
		PhaseIDs:                     append([]string(nil), installPlan.PhaseIDs...),
		BlockedReasons:               blockedReasons,
		NextSafeReadOnlyCheck:        multiApplicationInstallNextCheck(queueState),
		UserReviewRequired:           queueState != "ready",
		RequestObjectCreated:         false,
		ArtifactStaged:               false,
		InstallStarted:               false,
		BackendProcessStarted:        false,
		LaunchEnabled:                false,
		ExecutionStarted:             false,
		HostRootModified:             false,
		StateRootPathExposed:         false,
		RawCommandExposed:            false,
		BackendDetailsExposed:        false,
	}
}

func multiApplicationInstallBlockedReasons(installPlan CompatibilityInstallPlanPreview) []string {
	var reasons []string
	reasons = append(reasons, installPlan.Readiness.RecipeTrustBlockingReasons...)
	reasons = append(reasons, installPlan.Readiness.ArtifactStageBlockingReasons...)
	if !installPlan.Readiness.ArtifactStageReceiptReady {
		reasons = append(reasons, "artifact staging receipt is missing")
	}
	if !installPlan.Readiness.RequiredArtifactsStaged {
		reasons = append(reasons, "required artifacts are not staged")
	}
	if !installPlan.Readiness.RecipeInstallAllowed {
		reasons = append(reasons, "recipe install gate is not allowing this environment")
	}
	if !installPlan.DesktopActivationReady {
		reasons = append(reasons, "desktop activation is not ready")
	}
	if !installPlan.InstallReady {
		reasons = append(reasons, "install readiness is incomplete")
	}
	return uniqueSortedStrings(reasons)
}

func multiApplicationInstallItemState(installPlan CompatibilityInstallPlanPreview, blockedReasons []string) string {
	if installPlan.InstallReady && installPlan.DesktopActivationReady && len(blockedReasons) == 0 {
		return "ready"
	}
	if len(installPlan.Readiness.ArtifactStageBlockingReasons) > 0 || !installPlan.Readiness.ArtifactStageReceiptReady || !installPlan.Readiness.RequiredArtifactsStaged {
		return "missing-evidence"
	}
	if !installPlan.Readiness.RecipeInstallAllowed {
		return "needs-review"
	}
	return "blocked"
}

func multiApplicationInstallPriority(state string) string {
	switch state {
	case "blocked":
		return "high"
	case "missing-evidence":
		return "medium"
	default:
		return "normal"
	}
}

func multiApplicationInstallNextCheck(state string) string {
	switch state {
	case "ready":
		return "review desktop activation status"
	case "needs-review":
		return "review recipe trust and install gate"
	case "missing-evidence":
		return "review artifact staging receipt"
	default:
		return "review blocked install readiness"
	}
}

func countMultiApplicationInstallItems(items []MultiApplicationInstallItem) MultiApplicationInstallCounts {
	counts := MultiApplicationInstallCounts{Total: len(items)}
	for _, item := range items {
		switch item.QueueState {
		case "ready":
			counts.Ready++
		case "needs-review":
			counts.NeedsReview++
		case "missing-evidence":
			counts.MissingEvidence++
		default:
			counts.Blocked++
		}
	}
	return counts
}

func multiApplicationInstallApplicationIDs(items []MultiApplicationInstallItem) []string {
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ApplicationID)
	}
	return ids
}

func multiApplicationInstallQueueStatus(counts MultiApplicationInstallCounts) string {
	if counts.Total == 0 {
		return "empty"
	}
	if counts.Blocked > 0 {
		return "blocked"
	}
	if counts.MissingEvidence > 0 {
		return "missing-evidence"
	}
	if counts.NeedsReview > 0 {
		return "needs-review"
	}
	return "ready"
}

func uniqueSortedStrings(values []string) []string {
	seen := map[string]bool{}
	var result []string
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
