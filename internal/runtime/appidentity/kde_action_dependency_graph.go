package appidentity

import "errors"

type KDEActionDependencyGraphPreview struct {
	SchemaVersion            string                         `json:"schema_version"`
	RequestType              string                         `json:"request_type"`
	GraphType                string                         `json:"graph_type"`
	Source                   string                         `json:"source"`
	Desktop                  string                         `json:"desktop"`
	RuntimeMethod            string                         `json:"runtime_method"`
	ReadMethod               string                         `json:"read_method"`
	ApplicationID            string                         `json:"application_id"`
	ApplicationName          string                         `json:"application_name"`
	Icon                     string                         `json:"icon"`
	DesktopFile              string                         `json:"desktop_file"`
	LauncherCommand          []string                       `json:"launcher_command"`
	Queue                    KDEActionDependencyQueue       `json:"queue"`
	Nodes                    []KDEActionDependencyNode      `json:"nodes"`
	Edges                    []KDEActionDependencyEdge      `json:"edges"`
	NodeIDs                  []string                       `json:"node_ids"`
	MissingEvidenceIDs       []string                       `json:"missing_evidence_ids"`
	BlockedActions           []string                       `json:"blocked_actions"`
	NodeCount                int                            `json:"node_count"`
	EdgeCount                int                            `json:"edge_count"`
	ActionNodeCount          int                            `json:"action_node_count"`
	EvidenceNodeCount        int                            `json:"evidence_node_count"`
	GateNodeCount            int                            `json:"gate_node_count"`
	MissingEvidenceCount     int                            `json:"missing_evidence_count"`
	BlockedActionCount       int                            `json:"blocked_action_count"`
	ReadOnlyCheckCount       int                            `json:"read_only_check_count"`
	FileCount                int                            `json:"file_count"`
	FileURIs                 []string                       `json:"file_uris"`
	RuntimeOwned             bool                           `json:"runtime_owned"`
	GoRuntimeBacked          bool                           `json:"go_runtime_backed"`
	KDEPolicyOwner           bool                           `json:"kde_policy_owner"`
	OfficialDesktopOnly      bool                           `json:"official_desktop_only"`
	CompatibilityCenterGraph bool                           `json:"compatibility_center_graph"`
	SafeForAIDiagnostics     bool                           `json:"safe_for_ai_diagnostics"`
	UserDecisionCaptured     bool                           `json:"user_decision_captured"`
	UserDecisionAllowsLaunch bool                           `json:"user_decision_allows_launch"`
	DependencyGraphCreated   bool                           `json:"dependency_graph_created"`
	DependencyGraphPersisted bool                           `json:"dependency_graph_persisted"`
	SettingsPersisted        bool                           `json:"settings_persisted"`
	PermissionGrantCreated   bool                           `json:"permission_grant_created"`
	RequestObjectsCreated    bool                           `json:"request_objects_created"`
	RuntimeLaunchApproval    bool                           `json:"runtime_launch_approval"`
	LaunchAllowed            bool                           `json:"launch_allowed"`
	LaunchEnabled            bool                           `json:"launch_enabled"`
	ExecutionStarted         bool                           `json:"execution_started"`
	BackendProcessStarted    bool                           `json:"backend_process_started"`
	NetworkRequired          bool                           `json:"network_required"`
	HostRootModified         bool                           `json:"host_root_modified"`
	StateRootPathExposed     bool                           `json:"state_root_path_exposed"`
	RawCommandExposed        bool                           `json:"raw_command_exposed"`
	FileContentRead          bool                           `json:"file_content_read"`
	BackendDetailsExposed    bool                           `json:"backend_details_exposed"`
	UserFacingSettings       map[string]string              `json:"user_facing_settings"`
	ReceiptValidation        KDEActionReceiptValidation     `json:"receipt_validation"`
	NextReadOnlyChecks       []KDEActionDependencyNextCheck `json:"next_read_only_checks"`
	DesktopSafeSummary       string                         `json:"desktop_safe_summary"`
}

type KDEActionDependencyQueue struct {
	RequestType             string   `json:"request_type"`
	QueueType               string   `json:"queue_type"`
	Source                  string   `json:"source"`
	ActionIDs               []string `json:"action_ids"`
	ActionCount             int      `json:"action_count"`
	PendingActionCount      int      `json:"pending_action_count"`
	UserReviewRequiredCount int      `json:"user_review_required_count"`
	PortalActionCount       int      `json:"portal_action_count"`
	RuntimeGateActionCount  int      `json:"runtime_gate_action_count"`
	ActionQueueCreated      bool     `json:"action_queue_created"`
	ActionQueuePersisted    bool     `json:"action_queue_persisted"`
	RuntimeLaunchApproval   bool     `json:"runtime_launch_approval"`
	LaunchAllowed           bool     `json:"launch_allowed"`
	ExecutionStarted        bool     `json:"execution_started"`
	BackendDetailsExposed   bool     `json:"backend_details_exposed"`
}

type KDEActionDependencyNode struct {
	ID                     string   `json:"id"`
	NodeType               string   `json:"node_type"`
	ActionID               string   `json:"action_id,omitempty"`
	EntryPointID           string   `json:"entry_point_id,omitempty"`
	Title                  string   `json:"title"`
	State                  string   `json:"state"`
	RequiredEvidence       []string `json:"required_evidence,omitempty"`
	MissingEvidence        []string `json:"missing_evidence,omitempty"`
	BlockedReasons         []string `json:"blocked_reasons,omitempty"`
	NextSafeCheck          string   `json:"next_safe_check"`
	RuntimeOwned           bool     `json:"runtime_owned"`
	KDEPolicyOwner         bool     `json:"kde_policy_owner"`
	ExecutionEnabled       bool     `json:"execution_enabled"`
	RequestObjectsCreated  bool     `json:"request_objects_created"`
	PermissionGrantCreated bool     `json:"permission_grant_created"`
	SettingsPersisted      bool     `json:"settings_persisted"`
	RuntimeLaunchApproval  bool     `json:"runtime_launch_approval"`
	BackendProcessStarted  bool     `json:"backend_process_started"`
	HostRootModified       bool     `json:"host_root_modified"`
	BackendDetailsExposed  bool     `json:"backend_details_exposed"`
}

type KDEActionDependencyEdge struct {
	ID           string `json:"id"`
	From         string `json:"from"`
	To           string `json:"to"`
	Relation     string `json:"relation"`
	Blocking     bool   `json:"blocking"`
	RuntimeOwned bool   `json:"runtime_owned"`
}

type KDEActionReceiptValidation struct {
	ValidationType            string `json:"validation_type"`
	RejectsMismatchedAppID    bool   `json:"rejects_mismatched_app_id"`
	RejectsMalformedOperation bool   `json:"rejects_malformed_operation_id"`
	RejectsPathEscapeEvidence bool   `json:"rejects_path_escape_evidence"`
	RejectsUnsafeSideEffects  bool   `json:"rejects_unsafe_side_effects"`
	ReceiptRecorded           bool   `json:"receipt_recorded"`
	RequestObjectsCreated     bool   `json:"request_objects_created"`
	PermissionGrantCreated    bool   `json:"permission_grant_created"`
	SettingsPersisted         bool   `json:"settings_persisted"`
	ExecutionStarted          bool   `json:"execution_started"`
	HostRootModified          bool   `json:"host_root_modified"`
	BackendDetailsExposed     bool   `json:"backend_details_exposed"`
}

type KDEActionDependencyNextCheck struct {
	ID             string `json:"id"`
	ReadModel      string `json:"read_model"`
	Reason         string `json:"reason"`
	MutatesRuntime bool   `json:"mutates_runtime"`
	StartsProgram  bool   `json:"starts_program"`
}

func (plan Plan) KDEActionDependencyGraphPreview(decision string, fileURIs []string) (KDEActionDependencyGraphPreview, error) {
	if !singleLine(decision) {
		return KDEActionDependencyGraphPreview{}, errors.New("KDE action dependency graph preview requires a single-line decision")
	}
	for _, value := range []string{plan.ApplicationID, plan.DisplayName, plan.Icon, plan.DesktopFile} {
		if !singleLine(value) {
			return KDEActionDependencyGraphPreview{}, errors.New("KDE action dependency graph preview requires single-line identity fields")
		}
	}

	queue, err := plan.KDEActionQueuePreview(decision, fileURIs)
	if err != nil {
		return KDEActionDependencyGraphPreview{}, err
	}

	nodes, edges := kdeActionDependencyGraph(queue.Actions)
	nodeIDs, missingEvidenceIDs, blockedActionIDs, actionCount, evidenceCount, gateCount, readOnlyCount := summarizeKDEActionDependencyGraph(nodes)

	preview := KDEActionDependencyGraphPreview{
		SchemaVersion:   "xnix.runtime.kde_action_dependency_graph.v1",
		RequestType:     "kde-action-dependency-graph-preview",
		GraphType:       "compatibility-center-action-dependency-graph",
		Source:          "kde-action-queue-preview+runtime-evidence-prerequisites+runtime-write-gate-preview",
		Desktop:         "KDE Plasma",
		RuntimeMethod:   "GetKDEActionDependencyGraph",
		ReadMethod:      "GetKDEActionDependencyGraphPreview",
		ApplicationID:   queue.ApplicationID,
		ApplicationName: queue.ApplicationName,
		Icon:            queue.Icon,
		DesktopFile:     queue.DesktopFile,
		LauncherCommand: queue.LauncherCommand,
		Queue: KDEActionDependencyQueue{
			RequestType:             queue.RequestType,
			QueueType:               queue.QueueType,
			Source:                  queue.Source,
			ActionIDs:               queue.ActionIDs,
			ActionCount:             queue.ActionCount,
			PendingActionCount:      queue.PendingActionCount,
			UserReviewRequiredCount: queue.UserReviewRequiredCount,
			PortalActionCount:       queue.PortalActionCount,
			RuntimeGateActionCount:  queue.RuntimeGateActionCount,
			ActionQueueCreated:      queue.ActionQueueCreated,
			ActionQueuePersisted:    queue.ActionQueuePersisted,
			RuntimeLaunchApproval:   queue.RuntimeLaunchApproval,
			LaunchAllowed:           queue.LaunchAllowed,
			ExecutionStarted:        queue.ExecutionStarted,
			BackendDetailsExposed:   queue.BackendDetailsExposed,
		},
		Nodes:                    nodes,
		Edges:                    edges,
		NodeIDs:                  nodeIDs,
		MissingEvidenceIDs:       missingEvidenceIDs,
		BlockedActions:           blockedActionIDs,
		NodeCount:                len(nodes),
		EdgeCount:                len(edges),
		ActionNodeCount:          actionCount,
		EvidenceNodeCount:        evidenceCount,
		GateNodeCount:            gateCount,
		MissingEvidenceCount:     len(missingEvidenceIDs),
		BlockedActionCount:       len(blockedActionIDs),
		ReadOnlyCheckCount:       readOnlyCount,
		FileCount:                queue.FileCount,
		FileURIs:                 queue.FileURIs,
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		KDEPolicyOwner:           false,
		OfficialDesktopOnly:      true,
		CompatibilityCenterGraph: true,
		SafeForAIDiagnostics:     true,
		UserDecisionCaptured:     queue.UserDecisionCaptured,
		UserDecisionAllowsLaunch: queue.UserDecisionAllowsLaunch,
		DependencyGraphCreated:   true,
		DependencyGraphPersisted: false,
		SettingsPersisted:        false,
		PermissionGrantCreated:   false,
		RequestObjectsCreated:    false,
		RuntimeLaunchApproval:    false,
		LaunchAllowed:            false,
		LaunchEnabled:            false,
		ExecutionStarted:         false,
		BackendProcessStarted:    false,
		NetworkRequired:          false,
		HostRootModified:         false,
		StateRootPathExposed:     false,
		RawCommandExposed:        false,
		FileContentRead:          false,
		BackendDetailsExposed:    false,
		UserFacingSettings:       queue.UserFacingSettings,
		ReceiptValidation: KDEActionReceiptValidation{
			ValidationType:            "planned-review-receipt-validation",
			RejectsMismatchedAppID:    true,
			RejectsMalformedOperation: true,
			RejectsPathEscapeEvidence: true,
			RejectsUnsafeSideEffects:  true,
			ReceiptRecorded:           false,
			RequestObjectsCreated:     false,
			PermissionGrantCreated:    false,
			SettingsPersisted:         false,
			ExecutionStarted:          false,
			HostRootModified:          false,
			BackendDetailsExposed:     false,
		},
		NextReadOnlyChecks: []KDEActionDependencyNextCheck{
			{ID: "check-application-readiness", ReadModel: "application-readiness-preview", Reason: "refresh install, Portal, snapshot, execution, and write-gate evidence before enabling any action", MutatesRuntime: false, StartsProgram: false},
			{ID: "check-action-card-deck", ReadModel: "kde-action-card-deck-preview", Reason: "refresh renderable Compatibility Center card state from the dependency graph", MutatesRuntime: false, StartsProgram: false},
			{ID: "check-runtime-write-gate", ReadModel: "runtime-write-gate-preview", Reason: "confirm Runtime write dispatch is still disabled for action execution", MutatesRuntime: false, StartsProgram: false},
		},
		DesktopSafeSummary: "Compatibility Center can display why each KDE action is blocked, which Runtime evidence is missing, and which read-only checks are safe to run next; the graph does not approve, grant, persist, or start execution.",
	}
	if err := validateNoBackendTerms(preview, "KDE action dependency graph preview"); err != nil {
		return KDEActionDependencyGraphPreview{}, err
	}
	return preview, nil
}

func kdeActionDependencyGraph(actions []KDEActionQueueItem) ([]KDEActionDependencyNode, []KDEActionDependencyEdge) {
	nodes := make([]KDEActionDependencyNode, 0, len(actions)*4)
	edges := make([]KDEActionDependencyEdge, 0, len(actions)*3)
	for _, action := range actions {
		required := kdeActionRequiredEvidence(action)
		missing := kdeActionMissingEvidence(action)
		blocked := []string{"Runtime write gate is disabled", "review receipt is not recorded", "Compatibility Center action execution is not enabled"}
		if action.RequiresPortal {
			blocked = append(blocked, "Portal permission receipt is missing")
		}
		actionNode := KDEActionDependencyNode{
			ID:                     "action:" + action.ID,
			NodeType:               "action",
			ActionID:               action.ID,
			EntryPointID:           action.EntryPointID,
			Title:                  action.Title,
			State:                  "blocked",
			RequiredEvidence:       required,
			MissingEvidence:        missing,
			BlockedReasons:         blocked,
			NextSafeCheck:          "kde-action-card-preview",
			RuntimeOwned:           true,
			KDEPolicyOwner:         false,
			ExecutionEnabled:       false,
			RequestObjectsCreated:  false,
			PermissionGrantCreated: false,
			SettingsPersisted:      false,
			RuntimeLaunchApproval:  false,
			BackendProcessStarted:  false,
			HostRootModified:       false,
			BackendDetailsExposed:  false,
		}
		nodes = append(nodes, actionNode)
		for _, evidenceID := range required {
			evidenceNode := KDEActionDependencyNode{
				ID:                     "evidence:" + action.ID + ":" + evidenceID,
				NodeType:               "evidence",
				ActionID:               action.ID,
				EntryPointID:           action.EntryPointID,
				Title:                  kdeActionEvidenceTitle(evidenceID),
				State:                  "missing",
				BlockedReasons:         []string{"evidence is required before this action can become executable"},
				NextSafeCheck:          kdeActionEvidenceReadModel(evidenceID),
				RuntimeOwned:           true,
				KDEPolicyOwner:         false,
				ExecutionEnabled:       false,
				RequestObjectsCreated:  false,
				PermissionGrantCreated: false,
				SettingsPersisted:      false,
				RuntimeLaunchApproval:  false,
				BackendProcessStarted:  false,
				HostRootModified:       false,
				BackendDetailsExposed:  false,
			}
			nodes = append(nodes, evidenceNode)
			edges = append(edges, KDEActionDependencyEdge{
				ID:           action.ID + "->" + evidenceID,
				From:         evidenceNode.ID,
				To:           actionNode.ID,
				Relation:     "required-before-action",
				Blocking:     true,
				RuntimeOwned: true,
			})
		}
		gateID := "gate:" + action.ID + ":" + action.RuntimeGate
		nodes = append(nodes, KDEActionDependencyNode{
			ID:                     gateID,
			NodeType:               "gate",
			ActionID:               action.ID,
			EntryPointID:           action.EntryPointID,
			Title:                  "Runtime gate: " + action.RuntimeGate,
			State:                  "blocked",
			BlockedReasons:         []string{"Runtime gate is read-only and disabled for action execution"},
			NextSafeCheck:          "runtime-write-gate-preview",
			RuntimeOwned:           true,
			KDEPolicyOwner:         false,
			ExecutionEnabled:       false,
			RequestObjectsCreated:  false,
			PermissionGrantCreated: false,
			SettingsPersisted:      false,
			RuntimeLaunchApproval:  false,
			BackendProcessStarted:  false,
			HostRootModified:       false,
			BackendDetailsExposed:  false,
		})
		edges = append(edges, KDEActionDependencyEdge{
			ID:           action.ID + "->" + action.RuntimeGate,
			From:         gateID,
			To:           actionNode.ID,
			Relation:     "blocked-by-runtime-gate",
			Blocking:     true,
			RuntimeOwned: true,
		})
	}
	return nodes, edges
}

func summarizeKDEActionDependencyGraph(nodes []KDEActionDependencyNode) ([]string, []string, []string, int, int, int, int) {
	nodeIDs := make([]string, 0, len(nodes))
	missingEvidenceIDs := []string{}
	blockedActionIDs := []string{}
	actionCount := 0
	evidenceCount := 0
	gateCount := 0
	readOnlyChecks := map[string]bool{}
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.ID)
		if node.NextSafeCheck != "" {
			readOnlyChecks[node.NextSafeCheck] = true
		}
		switch node.NodeType {
		case "action":
			actionCount++
			if node.State == "blocked" {
				blockedActionIDs = append(blockedActionIDs, node.ActionID)
			}
		case "evidence":
			evidenceCount++
			if node.State == "missing" {
				missingEvidenceIDs = append(missingEvidenceIDs, node.ID)
			}
		case "gate":
			gateCount++
		}
	}
	return nodeIDs, missingEvidenceIDs, blockedActionIDs, actionCount, evidenceCount, gateCount, len(readOnlyChecks)
}

func kdeActionRequiredEvidence(action KDEActionQueueItem) []string {
	evidence := []string{"application-readiness", "action-review-receipt", "runtime-write-gate"}
	switch action.EntryPointID {
	case "launcher":
		evidence = append(evidence, "desktop-activation-status", "execution-preflight")
	case "task-manager":
		evidence = append(evidence, "execution-session-evidence", "task-manager-identity")
	case "file-manager":
		evidence = append(evidence, "portal-file-access-receipt", "file-association-plan")
	case "system-tray":
		evidence = append(evidence, "execution-session-evidence", "tray-status")
	case "notifications":
		evidence = append(evidence, "notification-plan", "action-status")
	case "compatibility-center":
		evidence = append(evidence, "center-page-readiness", "action-card-deck")
	case "settings":
		evidence = append(evidence, "settings-change-review", "review-flow")
	}
	return evidence
}

func kdeActionMissingEvidence(action KDEActionQueueItem) []string {
	missing := []string{"action-review-receipt", "runtime-write-gate"}
	if action.RequiresPortal {
		missing = append(missing, "portal-file-access-receipt")
	}
	return missing
}

func kdeActionEvidenceTitle(evidenceID string) string {
	switch evidenceID {
	case "application-readiness":
		return "Application readiness graph"
	case "action-review-receipt":
		return "Recorded action review receipt"
	case "runtime-write-gate":
		return "Runtime write gate"
	case "desktop-activation-status":
		return "Desktop activation status"
	case "execution-preflight":
		return "Execution preflight"
	case "execution-session-evidence":
		return "Execution session evidence"
	case "task-manager-identity":
		return "Task manager identity"
	case "portal-file-access-receipt":
		return "Portal file access receipt"
	case "file-association-plan":
		return "File association plan"
	case "tray-status":
		return "Tray status"
	case "notification-plan":
		return "Notification plan"
	case "action-status":
		return "Action status"
	case "center-page-readiness":
		return "Compatibility Center readiness"
	case "action-card-deck":
		return "Action card deck"
	case "settings-change-review":
		return "Settings change review"
	case "review-flow":
		return "Review flow"
	default:
		return "Runtime evidence"
	}
}

func kdeActionEvidenceReadModel(evidenceID string) string {
	switch evidenceID {
	case "application-readiness":
		return "application-readiness-preview"
	case "action-review-receipt":
		return "kde-action-receipt-preview"
	case "runtime-write-gate":
		return "runtime-write-gate-preview"
	case "desktop-activation-status":
		return "desktop-activation-status-preview"
	case "execution-preflight":
		return "execution-preflight-preview"
	case "execution-session-evidence":
		return "execution-session-status-preview"
	case "task-manager-identity":
		return "task-manager-identity-preview"
	case "portal-file-access-receipt":
		return "portal-request-preview"
	case "file-association-plan":
		return "mimeapps-preview"
	case "tray-status":
		return "tray-status-preview"
	case "notification-plan":
		return "notification-preview"
	case "action-status":
		return "kde-action-status-preview"
	case "center-page-readiness":
		return "kde-center-page-preview"
	case "action-card-deck":
		return "kde-action-card-deck-preview"
	case "settings-change-review":
		return "settings-change-preview"
	case "review-flow":
		return "review-flow-preview"
	default:
		return "application-readiness-preview"
	}
}
