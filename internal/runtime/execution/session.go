package execution

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const sessionRecordSchemaVersion = "xnix.runtime.execution_session_record.v1"

// SessionRecord is a durable, KDE-safe read model derived from an execution
// transaction. It records why a compatibility application session is not live
// yet without observing windows, launching backends, or mutating host state.
type SessionRecord struct {
	SchemaVersion                 string   `json:"schema_version"`
	RecordType                    string   `json:"record_type"`
	Source                        string   `json:"source"`
	RequestID                     string   `json:"request_id"`
	ApplicationID                 string   `json:"application_id"`
	RelativePath                  string   `json:"relative_path"`
	TransactionRelativePath       string   `json:"transaction_relative_path"`
	TransactionState              State    `json:"transaction_state"`
	SessionState                  string   `json:"session_state"`
	TaskManagerState              string   `json:"task_manager_state"`
	TrayState                     string   `json:"tray_state"`
	KWinState                     string   `json:"kwin_state"`
	CompatibilityCenterState      string   `json:"compatibility_center_state"`
	Gates                         []Gate   `json:"gates"`
	BlockedReasons                []string `json:"blocked_reasons"`
	PortalPermissionReceiptPaths  []string `json:"portal_permission_receipt_relative_paths"`
	PortalPermissionReceiptStates []string `json:"portal_permission_receipt_states"`
	SHA256                        string   `json:"sha256"`
	RuntimeOwned                  bool     `json:"runtime_owned"`
	GoRuntimeBacked               bool     `json:"go_runtime_backed"`
	KDEPolicyOwner                bool     `json:"kde_policy_owner"`
	StateRootPathExposed          bool     `json:"state_root_path_exposed"`
	StatusPersisted               bool     `json:"status_persisted"`
	SessionCreated                bool     `json:"session_created"`
	SessionRegistered             bool     `json:"session_registered"`
	SessionActive                 bool     `json:"session_active"`
	LiveStateObserved             bool     `json:"live_state_observed"`
	WindowObserved                bool     `json:"window_observed"`
	TaskManagerEntryActive        bool     `json:"task_manager_entry_active"`
	KWinRuleApplied               bool     `json:"kwin_rule_applied"`
	LiveTrayBridgeEnabled         bool     `json:"live_tray_bridge_enabled"`
	LaunchAllowed                 bool     `json:"launch_allowed"`
	LaunchEnabled                 bool     `json:"launch_enabled"`
	ExecutionStarted              bool     `json:"execution_started"`
	BackendProcessStarted         bool     `json:"backend_process_started"`
	PermissionGranted             bool     `json:"permission_granted"`
	HostRootModified              bool     `json:"host_root_modified"`
	NetworkRequired               bool     `json:"network_required"`
	PrivilegedContainerRequired   bool     `json:"privileged_container_required"`
	BackendDetailsExposed         bool     `json:"backend_details_exposed"`
	DesktopSafeSummary            string   `json:"desktop_safe_summary"`
}

// RecordSession derives and persists a session status record from an existing
// execution transaction record under the same state root.
func (l *Ledger) RecordSession(requestID string) (SessionRecord, error) {
	transactionRecord, err := l.Load(requestID)
	if err != nil {
		return SessionRecord{}, err
	}
	relativePath, path, err := l.sessionRecordPath(requestID)
	if err != nil {
		return SessionRecord{}, err
	}

	tx := transactionRecord.Transaction
	record := SessionRecord{
		SchemaVersion:                 sessionRecordSchemaVersion,
		RecordType:                    "execution-session-status-record",
		Source:                        "go-runtime-state-root-execution-session",
		RequestID:                     requestID,
		ApplicationID:                 transactionRecord.ApplicationID,
		RelativePath:                  relativePath,
		TransactionRelativePath:       transactionRecord.RelativePath,
		TransactionState:              tx.State,
		SessionState:                  sessionStateForTransaction(tx),
		TaskManagerState:              "blocked",
		TrayState:                     "blocked",
		KWinState:                     "blocked",
		CompatibilityCenterState:      "waiting-for-runtime-gates",
		Gates:                         append([]Gate{}, tx.Gates...),
		BlockedReasons:                append([]string{}, tx.BlockedReasons...),
		PortalPermissionReceiptPaths:  append([]string{}, transactionRecord.PortalPermissionReceiptPaths...),
		PortalPermissionReceiptStates: append([]string{}, transactionRecord.PortalPermissionReceiptStates...),
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		StateRootPathExposed:          false,
		StatusPersisted:               true,
		SessionCreated:                false,
		SessionRegistered:             false,
		SessionActive:                 false,
		LiveStateObserved:             false,
		WindowObserved:                false,
		TaskManagerEntryActive:        false,
		KWinRuleApplied:               false,
		LiveTrayBridgeEnabled:         false,
		LaunchAllowed:                 false,
		LaunchEnabled:                 false,
		ExecutionStarted:              false,
		BackendProcessStarted:         false,
		PermissionGranted:             false,
		HostRootModified:              false,
		NetworkRequired:               false,
		PrivilegedContainerRequired:   false,
		BackendDetailsExposed:         false,
		DesktopSafeSummary:            "Runtime persisted a blocked compatibility session status from the execution transaction without launching a backend or observing live windows.",
	}
	sort.Strings(record.BlockedReasons)
	sort.Slice(record.Gates, func(i, j int) bool { return record.Gates[i].ID < record.Gates[j].ID })

	data, digest, err := marshalSessionRecord(record)
	if err != nil {
		return SessionRecord{}, err
	}
	record.SHA256 = digest
	data, _, err = marshalSessionRecord(record)
	if err != nil {
		return SessionRecord{}, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return SessionRecord{}, fmt.Errorf("prepare execution session directory: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return SessionRecord{}, fmt.Errorf("write execution session record: %w", err)
	}
	return record, nil
}

// LoadSession reads and verifies one persisted session status record.
func (l *Ledger) LoadSession(requestID string) (SessionRecord, error) {
	relativePath, path, err := l.sessionRecordPath(requestID)
	if err != nil {
		return SessionRecord{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return SessionRecord{}, fmt.Errorf("read execution session record: %w", err)
	}
	var record SessionRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return SessionRecord{}, fmt.Errorf("parse execution session record: %w", err)
	}
	if record.SchemaVersion != sessionRecordSchemaVersion || record.RecordType != "execution-session-status-record" || record.Source != "go-runtime-state-root-execution-session" {
		return SessionRecord{}, errors.New("execution session record has unsupported schema")
	}
	if record.RequestID != requestID || record.ApplicationID == "" || record.RelativePath != relativePath {
		return SessionRecord{}, errors.New("execution session record identity or path mismatch")
	}
	storedDigest := record.SHA256
	record.SHA256 = ""
	_, expectedDigest, err := marshalSessionRecord(record)
	if err != nil {
		return SessionRecord{}, err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return SessionRecord{}, errors.New("execution session record digest mismatch")
	}
	record.SHA256 = storedDigest
	if record.StateRootPathExposed || record.SessionCreated || record.SessionRegistered || record.SessionActive || record.LiveStateObserved || record.WindowObserved || record.TaskManagerEntryActive || record.KWinRuleApplied || record.LiveTrayBridgeEnabled || record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.BackendProcessStarted || record.PermissionGranted || record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.BackendDetailsExposed {
		return SessionRecord{}, errors.New("execution session record has unsafe enabled gates")
	}
	return record, nil
}

func sessionStateForTransaction(tx Transaction) string {
	if tx.ReviewDecision == DecisionRejected {
		return "review-declined"
	}
	if tx.State == StateBlocked {
		return "blocked"
	}
	if tx.State == StatePreflight {
		return "preflight"
	}
	return "planned"
}

func (l *Ledger) sessionRecordPath(requestID string) (string, string, error) {
	if err := validateLedgerID(requestID); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "sessions", requestID+".json"))
	path := filepath.Join(l.root, filepath.FromSlash(relativePath))
	cleanRoot := filepath.Clean(l.root)
	rel, err := filepath.Rel(cleanRoot, path)
	if err != nil {
		return "", "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("execution session path escapes state root: %s", relativePath)
	}
	return relativePath, path, nil
}

func marshalSessionRecord(record SessionRecord) ([]byte, string, error) {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}
