package appidentity

import (
	"errors"
	"sort"
	"strings"
)

type RepairPlanPreview struct {
	SchemaVersion            string                 `json:"schema_version"`
	RequestType              string                 `json:"request_type"`
	PlanType                 string                 `json:"plan_type"`
	Source                   string                 `json:"source"`
	Desktop                  string                 `json:"desktop"`
	RuntimeMethod            string                 `json:"runtime_method"`
	ReadMethod               string                 `json:"read_method"`
	ApplicationID            string                 `json:"application_id"`
	Issue                    string                 `json:"issue"`
	RuntimeOwned             bool                   `json:"runtime_owned"`
	GoRuntimeBacked          bool                   `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                   `json:"kde_policy_owner"`
	Severity                 string                 `json:"severity"`
	AutomaticAllowed         bool                   `json:"automatic_allowed"`
	UserApprovalRequired     bool                   `json:"user_approval_required"`
	SnapshotRequired         bool                   `json:"snapshot_required"`
	SnapshotPlan             *RepairSnapshotSummary `json:"snapshot_plan"`
	RollbackAvailable        bool                   `json:"rollback_available"`
	Actions                  []RepairAction         `json:"actions"`
	ActionIDs                []string               `json:"action_ids"`
	NotificationEvent        string                 `json:"notification_event"`
	RepairExecutionRequested bool                   `json:"repair_execution_requested"`
	RepairExecuted           bool                   `json:"repair_executed"`
	BackendLaunchEnabled     bool                   `json:"backend_launch_enabled"`
	NetworkRequired          bool                   `json:"network_required"`
	HostRootModified         bool                   `json:"host_root_modified"`
	BackendDetailsExposed    bool                   `json:"backend_details_exposed"`
	DesktopSafeSummary       string                 `json:"desktop_safe_summary"`
}

type RepairAction struct {
	ID    string `json:"id"`
	Kind  string `json:"kind"`
	Label string `json:"label"`
}

type RepairSnapshotSummary struct {
	PlanType              string          `json:"plan_type"`
	Reason                string          `json:"reason"`
	EnabledByDefault      bool            `json:"enabled_by_default"`
	RestoreAvailable      bool            `json:"restore_available"`
	PreserveUserDocuments bool            `json:"preserve_user_documents"`
	Retention             RepairRetention `json:"retention"`
}

type RepairRetention struct {
	Policy             string `json:"policy"`
	KeepLatest         int    `json:"keep_latest"`
	PruneAutomatically bool   `json:"prune_automatically"`
}

type repairIssueRule struct {
	severity             string
	automaticAllowed     bool
	userApprovalRequired bool
	summary              string
	actions              []RepairAction
}

var repairIssueRules = map[string]repairIssueRule{
	"engine-binding-pending": {
		severity: "warning", automaticAllowed: false, userApprovalRequired: true,
		summary: "Compatibility engine setup is not ready yet.",
		actions: []RepairAction{
			{ID: "open-diagnostics", Kind: "user-visible", Label: "Open diagnostics"},
			{ID: "prepare-engine-binding", Kind: "runtime-task", Label: "Prepare compatibility engine binding"},
		},
	},
	"portal-approval-required": {
		severity: "info", automaticAllowed: false, userApprovalRequired: true,
		summary: "A desktop permission request needs user approval.",
		actions: []RepairAction{
			{ID: "request-portal-grant", Kind: "portal-request", Label: "Request desktop permission"},
		},
	},
	"recipe-trust-blocked": {
		severity: "critical", automaticAllowed: false, userApprovalRequired: true,
		summary: "Recipe trust requirements are not satisfied.",
		actions: []RepairAction{
			{ID: "review-recipe-source", Kind: "user-visible", Label: "Review recipe source"},
		},
	},
	"runtime-repair-applied": {
		severity: "info", automaticAllowed: true, userApprovalRequired: false,
		summary: "A safe compatibility repair was applied.",
		actions: []RepairAction{
			{ID: "show-repair-record", Kind: "user-visible", Label: "Show repair record"},
		},
	},
}

func repairIssueIDs() []string {
	ids := make([]string, 0, len(repairIssueRules))
	for id := range repairIssueRules {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func (plan Plan) RepairPlanPreview(issue string) (RepairPlanPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return RepairPlanPreview{}, err
	}
	rule, ok := repairIssueRules[issue]
	if !ok {
		return RepairPlanPreview{}, errors.New("issue must be one of: " + strings.Join(repairIssueIDs(), ", "))
	}

	snapshotRequired := issue == "engine-binding-pending" || issue == "runtime-repair-applied"
	var snapshotPlan *RepairSnapshotSummary
	if snapshotRequired {
		snapshotPlan = &RepairSnapshotSummary{
			PlanType:              "compatibility-snapshot",
			Reason:                "before-repair",
			EnabledByDefault:      true,
			RestoreAvailable:      true,
			PreserveUserDocuments: true,
			Retention:             RepairRetention{Policy: "bounded", KeepLatest: 5, PruneAutomatically: true},
		}
	}

	preview := RepairPlanPreview{
		SchemaVersion:            "xnix.runtime.repair_plan.v1",
		RequestType:              "repair-plan-preview",
		PlanType:                 "compatibility-repair",
		Source:                   "registry+go-runtime-repair-plan",
		Desktop:                  "KDE Plasma",
		RuntimeMethod:            "GetRepairPlan",
		ReadMethod:               "GetRepairPlanPreview",
		ApplicationID:            plan.ApplicationID,
		Issue:                    issue,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		Severity:                 rule.severity,
		AutomaticAllowed:         rule.automaticAllowed,
		UserApprovalRequired:     rule.userApprovalRequired,
		SnapshotRequired:         snapshotRequired,
		SnapshotPlan:             snapshotPlan,
		RollbackAvailable:        true,
		Actions:                  rule.actions,
		ActionIDs:                repairActionIDs(rule.actions),
		NotificationEvent:        repairNotificationEvent(issue, rule),
		RepairExecutionRequested: false,
		RepairExecuted:           false,
		BackendLaunchEnabled:     false,
		NetworkRequired:          false,
		HostRootModified:         false,
		BackendDetailsExposed:    false,
		DesktopSafeSummary:       rule.summary,
	}
	if err := validateNoBackendTerms(preview, "repair plan preview"); err != nil {
		return RepairPlanPreview{}, err
	}
	return preview, nil
}

func repairNotificationEvent(issue string, rule repairIssueRule) string {
	if issue == "runtime-repair-applied" {
		return "repair-applied"
	}
	if rule.userApprovalRequired {
		return "approval-required"
	}
	return "install-failed"
}

func repairActionIDs(actions []RepairAction) []string {
	ids := make([]string, 0, len(actions))
	for _, action := range actions {
		ids = append(ids, action.ID)
	}
	return ids
}
