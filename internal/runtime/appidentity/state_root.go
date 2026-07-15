package appidentity

import "errors"

type ApplicationStateRootPreview struct {
	SchemaVersion              string                   `json:"schema_version"`
	RequestType                string                   `json:"request_type"`
	RootType                   string                   `json:"root_type"`
	Source                     string                   `json:"source"`
	Desktop                    string                   `json:"desktop"`
	RuntimeMethod              string                   `json:"runtime_method"`
	ReadMethod                 string                   `json:"read_method"`
	Application                PackageSourceApplication `json:"application"`
	RuntimeOwned               bool                     `json:"runtime_owned"`
	GoRuntimeBacked            bool                     `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                     `json:"kde_policy_owner"`
	StateNamespace             string                   `json:"state_namespace"`
	StorageScope               string                   `json:"storage_scope"`
	AllocationState            string                   `json:"allocation_state"`
	DirectoriesCreated         bool                     `json:"directories_created"`
	HostRootModified           bool                     `json:"host_root_modified"`
	UserDocumentsIncluded      bool                     `json:"user_documents_included"`
	PortalRequiredForUserFiles bool                     `json:"portal_required_for_user_files"`
	SnapshotEligible           bool                     `json:"snapshot_eligible"`
	RestoreRequiresConfirm     bool                     `json:"restore_requires_confirmation"`
	RetentionPolicy            StateRootRetention       `json:"retention_policy"`
	ManagedScopes              []StateRootScope         `json:"managed_scopes"`
	ManagedScopeIDs            []string                 `json:"managed_scope_ids"`
	BlockedActions             []string                 `json:"blocked_actions"`
	BackendDetailsExposed      bool                     `json:"backend_details_exposed"`
	DesktopSafeSummary         string                   `json:"desktop_safe_summary"`
}

type StateRootRetention struct {
	AutomaticRestorePoints int  `json:"automatic_restore_points"`
	ManualRestorePoints    int  `json:"manual_restore_points"`
	UserDocumentsExcluded  bool `json:"user_documents_excluded"`
}

type StateRootScope struct {
	ID               string `json:"id"`
	RuntimeOwned     bool   `json:"runtime_owned"`
	SnapshotIncluded bool   `json:"snapshot_included"`
	Summary          string `json:"summary"`
}

func (plan Plan) ApplicationStateRootPreview() (ApplicationStateRootPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return ApplicationStateRootPreview{}, err
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return ApplicationStateRootPreview{}, errors.New("application state root preview requires single-line identity fields")
		}
	}

	scopes := stateRootScopes()
	preview := ApplicationStateRootPreview{
		SchemaVersion: "xnix.runtime.state_root.v1",
		RequestType:   "state-root-preview",
		RootType:      "compatibility-application-state-root",
		Source:        "registry+go-runtime-state-root",
		Desktop:       "KDE Plasma",
		RuntimeMethod: "GetApplicationStateRoot",
		ReadMethod:    "GetApplicationStateRootPreview",
		Application: PackageSourceApplication{
			ID:            plan.ApplicationID,
			Name:          plan.DisplayName,
			Icon:          plan.Icon,
			DesktopFile:   plan.DesktopFile,
			RequestedMode: plan.runtimeModeID(),
		},
		RuntimeOwned:               true,
		GoRuntimeBacked:            true,
		KDEPolicyOwner:             false,
		StateNamespace:             stateRootNamespace(plan.ApplicationID),
		StorageScope:               "per-application",
		AllocationState:            "planned",
		DirectoriesCreated:         false,
		HostRootModified:           false,
		UserDocumentsIncluded:      false,
		PortalRequiredForUserFiles: true,
		SnapshotEligible:           true,
		RestoreRequiresConfirm:     true,
		RetentionPolicy: StateRootRetention{
			AutomaticRestorePoints: 5,
			ManualRestorePoints:    10,
			UserDocumentsExcluded:  true,
		},
		ManagedScopes:   scopes,
		ManagedScopeIDs: stateRootScopeIDs(scopes),
		BlockedActions: []string{
			"write application state outside Runtime ownership",
			"include user documents in state snapshots",
			"expose host storage paths to KDE",
			"restore state without user confirmation",
		},
		BackendDetailsExposed: false,
		DesktopSafeSummary:    "Runtime application state root is planned and isolated from user documents.",
	}
	if err := validateNoBackendTerms(preview, "application state root preview"); err != nil {
		return ApplicationStateRootPreview{}, err
	}
	return preview, nil
}

func stateRootNamespace(applicationID string) string {
	sanitized := make([]rune, 0, len(applicationID))
	for _, r := range applicationID {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-':
			sanitized = append(sanitized, r)
		default:
			sanitized = append(sanitized, '-')
		}
	}
	return string(sanitized)
}

func stateRootScopes() []StateRootScope {
	return []StateRootScope{
		{ID: "application-data", RuntimeOwned: true, SnapshotIncluded: true, Summary: "Application-managed state is isolated under Runtime ownership."},
		{ID: "runtime-metadata", RuntimeOwned: true, SnapshotIncluded: true, Summary: "Runtime metadata tracks compatibility state without exposing host paths."},
		{ID: "diagnostic-cache", RuntimeOwned: true, SnapshotIncluded: false, Summary: "Diagnostic cache is Runtime-owned and can be regenerated."},
		{ID: "desktop-activation-receipts", RuntimeOwned: true, SnapshotIncluded: true, Summary: "Desktop activation receipts remain part of rollback planning."},
	}
}

func stateRootScopeIDs(scopes []StateRootScope) []string {
	ids := make([]string, 0, len(scopes))
	for _, scope := range scopes {
		ids = append(ids, scope.ID)
	}
	return ids
}
