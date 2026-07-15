package appidentity

import "errors"

type SnapshotPlanPreview struct {
	SchemaVersion          string            `json:"schema_version"`
	RequestType            string            `json:"request_type"`
	PlanType               string            `json:"plan_type"`
	Source                 string            `json:"source"`
	Desktop                string            `json:"desktop"`
	RuntimeMethod          string            `json:"runtime_method"`
	ReadMethod             string            `json:"read_method"`
	ApplicationID          string            `json:"application_id"`
	Reason                 string            `json:"reason"`
	RuntimeOwned           bool              `json:"runtime_owned"`
	GoRuntimeBacked        bool              `json:"go_runtime_backed"`
	KDEPolicyOwner         bool              `json:"kde_policy_owner"`
	EnabledByDefault       bool              `json:"enabled_by_default"`
	SnapshotScope          SnapshotScope     `json:"snapshot_scope"`
	Restore                SnapshotRestore   `json:"restore"`
	Retention              SnapshotRetention `json:"retention"`
	SnapshotRequestCreated bool              `json:"snapshot_request_created"`
	SnapshotCreated        bool              `json:"snapshot_created"`
	RestoreRequested       bool              `json:"restore_requested"`
	RestoreExecuted        bool              `json:"restore_executed"`
	UserDocumentsIncluded  bool              `json:"user_documents_included"`
	HostSystemIncluded     bool              `json:"host_system_included"`
	HostRootModified       bool              `json:"host_root_modified"`
	BackendDetailsExposed  bool              `json:"backend_details_exposed"`
	DesktopSafeSummary     string            `json:"desktop_safe_summary"`
}

type SnapshotScope struct {
	ApplicationState          bool `json:"application_state"`
	RuntimeMetadata           bool `json:"runtime_metadata"`
	DesktopActivationReceipts bool `json:"desktop_activation_receipts"`
	UserDocuments             bool `json:"user_documents"`
	HostSystem                bool `json:"host_system"`
}

type SnapshotRestore struct {
	Available                bool   `json:"available"`
	Method                   string `json:"method"`
	RequiresUserConfirmation bool   `json:"requires_user_confirmation"`
	PreserveUserDocuments    bool   `json:"preserve_user_documents"`
}

type SnapshotRetention struct {
	Policy             string `json:"policy"`
	KeepLatest         int    `json:"keep_latest"`
	PruneAutomatically bool   `json:"prune_automatically"`
}

var snapshotPlanReasonSummaries = map[string]string{
	"before-repair":        "Create a restore point before a compatibility repair.",
	"before-engine-change": "Create a restore point before changing the compatibility engine.",
	"manual":               "Create a user-requested restore point.",
}

func NewSnapshotPlanPreview(applicationID string, reason string) (SnapshotPlanPreview, error) {
	if !idPattern.MatchString(applicationID) {
		return SnapshotPlanPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}
	summary, ok := snapshotPlanReasonSummaries[reason]
	if !ok {
		return SnapshotPlanPreview{}, errors.New("reason must be one of: before-repair, before-engine-change, manual")
	}

	preview := SnapshotPlanPreview{
		SchemaVersion:    "xnix.runtime.snapshot_plan.v1",
		RequestType:      "snapshot-plan-preview",
		PlanType:         "compatibility-snapshot",
		Source:           "go-runtime-snapshot-plan",
		Desktop:          "KDE Plasma",
		RuntimeMethod:    "GetSnapshotPlan",
		ReadMethod:       "GetSnapshotPlanPreview",
		ApplicationID:    applicationID,
		Reason:           reason,
		RuntimeOwned:     true,
		GoRuntimeBacked:  true,
		KDEPolicyOwner:   false,
		EnabledByDefault: true,
		SnapshotScope: SnapshotScope{
			ApplicationState:          true,
			RuntimeMetadata:           true,
			DesktopActivationReceipts: true,
			UserDocuments:             false,
			HostSystem:                false,
		},
		Restore: SnapshotRestore{
			Available:                true,
			Method:                   "RestoreSnapshot",
			RequiresUserConfirmation: true,
			PreserveUserDocuments:    true,
		},
		Retention: SnapshotRetention{
			Policy:             "bounded",
			KeepLatest:         5,
			PruneAutomatically: true,
		},
		SnapshotRequestCreated: false,
		SnapshotCreated:        false,
		RestoreRequested:       false,
		RestoreExecuted:        false,
		UserDocumentsIncluded:  false,
		HostSystemIncluded:     false,
		HostRootModified:       false,
		BackendDetailsExposed:  false,
		DesktopSafeSummary:     summary,
	}
	if err := validateNoBackendTerms(preview, "snapshot plan preview"); err != nil {
		return SnapshotPlanPreview{}, err
	}
	return preview, nil
}
