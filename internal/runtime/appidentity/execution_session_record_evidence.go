package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ExecutionSessionRecordEvidence struct {
	ReceiptRelativePath      string
	RequestID                string
	SessionState             string
	TaskManagerState         string
	TrayState                string
	KWinState                string
	CompatibilityCenterState string
	SafeForKDE               bool
}

type ExecutionSessionFanOutEvidence struct {
	SchemaVersion          string                          `json:"schema_version"`
	RequestType            string                          `json:"request_type"`
	Source                 string                          `json:"source"`
	RuntimeMethod          string                          `json:"runtime_method"`
	ReadMethod             string                          `json:"read_method"`
	ReceiptRelativePath    string                          `json:"receipt_relative_path"`
	RequestID              string                          `json:"request_id"`
	SessionState           string                          `json:"session_state"`
	SurfaceCount           int                             `json:"surface_count"`
	Surfaces               []ExecutionSessionFanOutSurface `json:"surfaces"`
	TaskManager            ExecutionSessionFanOutSurface   `json:"task_manager"`
	KWin                   ExecutionSessionFanOutSurface   `json:"kwin"`
	Tray                   ExecutionSessionFanOutSurface   `json:"tray"`
	CompatibilityCenter    ExecutionSessionFanOutSurface   `json:"compatibility_center"`
	SafeForKDE             bool                            `json:"safe_for_kde"`
	RuntimeOwned           bool                            `json:"runtime_owned"`
	GoRuntimeBacked        bool                            `json:"go_runtime_backed"`
	KDEPolicyOwner         bool                            `json:"kde_policy_owner"`
	StateRootPathExposed   bool                            `json:"state_root_path_exposed"`
	LiveStateObserved      bool                            `json:"live_state_observed"`
	WindowObserved         bool                            `json:"window_observed"`
	TaskManagerEntryActive bool                            `json:"task_manager_entry_active"`
	KWinRuleApplied        bool                            `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled  bool                            `json:"live_tray_bridge_enabled"`
	LaunchEnabled          bool                            `json:"launch_enabled"`
	ExecutionStarted       bool                            `json:"execution_started"`
	BackendProcessStarted  bool                            `json:"backend_process_started"`
	PermissionGranted      bool                            `json:"permission_granted"`
	HostRootModified       bool                            `json:"host_root_modified"`
	NetworkRequired        bool                            `json:"network_required"`
	BackendDetailsExposed  bool                            `json:"backend_details_exposed"`
	DesktopSafeSummary     string                          `json:"desktop_safe_summary"`
}

type ExecutionSessionFanOutSurface struct {
	ID                  string `json:"id"`
	Consumer            string `json:"consumer"`
	RuntimeMethod       string `json:"runtime_method"`
	State               string `json:"state"`
	ReceiptRelativePath string `json:"receipt_relative_path"`
	NavigationOnly      bool   `json:"navigation_only"`
	ReadOnly            bool   `json:"read_only"`
	MutatesRuntime      bool   `json:"mutates_runtime"`
	StartsProgram       bool   `json:"starts_program"`
	SafeForKDE          bool   `json:"safe_for_kde"`
	Summary             string `json:"summary"`
}

type executionSessionRecordFile struct {
	SchemaVersion            string `json:"schema_version"`
	RecordType               string `json:"record_type"`
	RequestID                string `json:"request_id"`
	ApplicationID            string `json:"application_id"`
	RelativePath             string `json:"relative_path"`
	SessionState             string `json:"session_state"`
	TaskManagerState         string `json:"task_manager_state"`
	TrayState                string `json:"tray_state"`
	KWinState                string `json:"kwin_state"`
	CompatibilityCenterState string `json:"compatibility_center_state"`
	StateRootPathExposed     bool   `json:"state_root_path_exposed"`
	StatusPersisted          bool   `json:"status_persisted"`
	SessionActive            bool   `json:"session_active"`
	LiveStateObserved        bool   `json:"live_state_observed"`
	WindowObserved           bool   `json:"window_observed"`
	TaskManagerEntryActive   bool   `json:"task_manager_entry_active"`
	KWinRuleApplied          bool   `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled    bool   `json:"live_tray_bridge_enabled"`
	LaunchEnabled            bool   `json:"launch_enabled"`
	ExecutionStarted         bool   `json:"execution_started"`
	BackendProcessStarted    bool   `json:"backend_process_started"`
	PermissionGranted        bool   `json:"permission_granted"`
	HostRootModified         bool   `json:"host_root_modified"`
	NetworkRequired          bool   `json:"network_required"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
}

func (plan Plan) ExecutionSessionRecordEvidence(root string, requestID string) (ExecutionSessionRecordEvidence, error) {
	root = strings.TrimSpace(root)
	requestID = strings.TrimSpace(requestID)
	if root == "" {
		return ExecutionSessionRecordEvidence{}, errors.New("execution session evidence requires a root")
	}
	if err := validateExecutionSessionRequestID(requestID); err != nil {
		return ExecutionSessionRecordEvidence{}, err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "sessions", requestID+".json"))
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	cleanRoot := filepath.Clean(root)
	rel, err := filepath.Rel(cleanRoot, path)
	if err != nil {
		return ExecutionSessionRecordEvidence{}, err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("execution session evidence path escapes root: %s", relativePath)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("read execution session evidence: %w", err)
	}
	var record executionSessionRecordFile
	if err := json.Unmarshal(data, &record); err != nil {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("parse execution session evidence: %w", err)
	}
	if record.SchemaVersion != "xnix.runtime.execution_session_record.v1" || record.RecordType != "execution-session-status-record" {
		return ExecutionSessionRecordEvidence{}, errors.New("execution session evidence has unsupported schema")
	}
	if record.RequestID != requestID {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("execution session evidence request mismatch: %s", record.RequestID)
	}
	if record.ApplicationID != plan.ApplicationID {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("execution session evidence application mismatch: %s", record.ApplicationID)
	}
	if record.RelativePath != relativePath || filepath.IsAbs(record.RelativePath) || strings.Contains(record.RelativePath, "..") {
		return ExecutionSessionRecordEvidence{}, fmt.Errorf("execution session evidence has unsafe relative path: %s", record.RelativePath)
	}
	if !record.StatusPersisted ||
		record.StateRootPathExposed ||
		record.SessionActive ||
		record.LiveStateObserved ||
		record.WindowObserved ||
		record.TaskManagerEntryActive ||
		record.KWinRuleApplied ||
		record.LiveTrayBridgeEnabled ||
		record.LaunchEnabled ||
		record.ExecutionStarted ||
		record.BackendProcessStarted ||
		record.PermissionGranted ||
		record.HostRootModified ||
		record.NetworkRequired ||
		record.BackendDetailsExposed {
		return ExecutionSessionRecordEvidence{}, errors.New("execution session evidence is not safe for KDE")
	}
	for _, value := range []string{record.SessionState, record.TaskManagerState, record.TrayState, record.KWinState, record.CompatibilityCenterState} {
		if !singleLine(value) {
			return ExecutionSessionRecordEvidence{}, errors.New("execution session evidence requires single-line states")
		}
	}
	return ExecutionSessionRecordEvidence{
		ReceiptRelativePath:      record.RelativePath,
		RequestID:                record.RequestID,
		SessionState:             record.SessionState,
		TaskManagerState:         record.TaskManagerState,
		TrayState:                record.TrayState,
		KWinState:                record.KWinState,
		CompatibilityCenterState: record.CompatibilityCenterState,
		SafeForKDE:               true,
	}, nil
}

func (plan Plan) ExecutionSessionFanOutEvidence(root string, requestID string) (ExecutionSessionFanOutEvidence, error) {
	record, err := plan.ExecutionSessionRecordEvidence(root, requestID)
	if err != nil {
		return ExecutionSessionFanOutEvidence{}, err
	}
	taskManager := executionSessionFanOutSurface("task-manager", "Plasma task manager", "GetTaskManagerIdentityPlan", record.TaskManagerState, record.ReceiptRelativePath, "KDE can show task-manager session state without activating task-manager entries or observing windows.")
	kwin := executionSessionFanOutSurface("kwin", "KWin rule planner", "GetKWinWindowRulePlan", record.KWinState, record.ReceiptRelativePath, "KDE can show KWin session state without applying rules or observing windows.")
	tray := executionSessionFanOutSurface("tray", "Plasma system tray", "GetTrayStatus", record.TrayState, record.ReceiptRelativePath, "KDE can show tray session state without enabling a live tray bridge.")
	center := executionSessionFanOutSurface("compatibility-center", "KDE Compatibility Center", "GetKDECenterPage", record.CompatibilityCenterState, record.ReceiptRelativePath, "KDE can show Compatibility Center session state without creating requests or starting execution.")
	surfaces := []ExecutionSessionFanOutSurface{taskManager, kwin, tray, center}
	fanOut := ExecutionSessionFanOutEvidence{
		SchemaVersion:          "xnix.runtime.session_fanout.v1",
		RequestType:            "execution-session-fanout-evidence",
		Source:                 "execution-session-record",
		RuntimeMethod:          "GetExecutionSessionFanOutEvidence",
		ReadMethod:             "GetExecutionSessionFanOutEvidence",
		ReceiptRelativePath:    record.ReceiptRelativePath,
		RequestID:              record.RequestID,
		SessionState:           record.SessionState,
		SurfaceCount:           len(surfaces),
		Surfaces:               surfaces,
		TaskManager:            taskManager,
		KWin:                   kwin,
		Tray:                   tray,
		CompatibilityCenter:    center,
		SafeForKDE:             record.SafeForKDE,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		StateRootPathExposed:   false,
		LiveStateObserved:      false,
		WindowObserved:         false,
		TaskManagerEntryActive: false,
		KWinRuleApplied:        false,
		LiveTrayBridgeEnabled:  false,
		LaunchEnabled:          false,
		ExecutionStarted:       false,
		BackendProcessStarted:  false,
		PermissionGranted:      false,
		HostRootModified:       false,
		NetworkRequired:        false,
		BackendDetailsExposed:  false,
		DesktopSafeSummary:     "Runtime fans out one safe execution session record to task manager, KWin, tray, and Compatibility Center consumers without observing live windows, launching backends, or mutating the host root.",
	}
	if err := validateNoBackendTerms(fanOut, "execution session fan-out evidence"); err != nil {
		return ExecutionSessionFanOutEvidence{}, err
	}
	return fanOut, nil
}

func executionSessionFanOutSurface(id string, consumer string, runtimeMethod string, state string, receiptPath string, summary string) ExecutionSessionFanOutSurface {
	return ExecutionSessionFanOutSurface{
		ID:                  id,
		Consumer:            consumer,
		RuntimeMethod:       runtimeMethod,
		State:               state,
		ReceiptRelativePath: receiptPath,
		NavigationOnly:      true,
		ReadOnly:            true,
		MutatesRuntime:      false,
		StartsProgram:       false,
		SafeForKDE:          true,
		Summary:             summary,
	}
}

func validateExecutionSessionRequestID(requestID string) error {
	if requestID == "" {
		return errors.New("execution session evidence requires a request id")
	}
	if strings.Contains(requestID, "..") {
		return fmt.Errorf("execution session request id must not contain path traversal: %q", requestID)
	}
	for _, r := range requestID {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			continue
		}
		return fmt.Errorf("execution session request id contains unsafe character: %q", r)
	}
	return nil
}
