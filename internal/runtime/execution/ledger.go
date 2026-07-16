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

const ledgerSchemaVersion = "xnix.runtime.execution_ledger.v1"

// Ledger stores blocked-by-default execution transaction records under a
// caller-provided Runtime state root. It never starts backends, grants
// permissions, or writes outside that root.
type Ledger struct {
	root string
}

// LedgerRecord is the persisted, KDE-safe transaction record.
type LedgerRecord struct {
	SchemaVersion                   string      `json:"schema_version"`
	RecordType                      string      `json:"record_type"`
	Source                          string      `json:"source"`
	RequestID                       string      `json:"request_id"`
	ApplicationID                   string      `json:"application_id"`
	RelativePath                    string      `json:"relative_path"`
	Transaction                     Transaction `json:"transaction"`
	SHA256                          string      `json:"sha256"`
	PortalPermissionReceiptCount    int         `json:"portal_permission_receipt_count"`
	PortalPermissionReceiptPaths    []string    `json:"portal_permission_receipt_relative_paths"`
	PortalPermissionReceiptStates   []string    `json:"portal_permission_receipt_states"`
	PortalPermissionReceiptConsumed bool        `json:"portal_permission_receipt_consumed"`
	RuntimeOwned                    bool        `json:"runtime_owned"`
	GoRuntimeBacked                 bool        `json:"go_runtime_backed"`
	KDEPolicyOwner                  bool        `json:"kde_policy_owner"`
	StateRootPathExposed            bool        `json:"state_root_path_exposed"`
	LaunchAllowed                   bool        `json:"launch_allowed"`
	LaunchEnabled                   bool        `json:"launch_enabled"`
	BackendStarted                  bool        `json:"backend_started"`
	PermissionGranted               bool        `json:"permission_granted"`
	HostRootModified                bool        `json:"host_root_modified"`
	NetworkRequired                 bool        `json:"network_required"`
	PrivilegedContainerRequired     bool        `json:"privileged_container_required"`
	BackendDetailsExposed           bool        `json:"backend_details_exposed"`
	Summary                         string      `json:"summary"`
}

// NewLedger opens a ledger rooted under stateRoot.
func NewLedger(stateRoot string) (*Ledger, error) {
	root, err := safeLedgerRoot(stateRoot)
	if err != nil {
		return nil, err
	}
	return &Ledger{root: root}, nil
}

// Record persists the transaction snapshot for later KDE/Runtime inspection.
func (l *Ledger) Record(tx Transaction) (LedgerRecord, error) {
	if tx.RequestID == "" {
		return LedgerRecord{}, errors.New("execution ledger requires a transaction request id")
	}
	if tx.ApplicationID == "" {
		return LedgerRecord{}, errors.New("execution ledger requires an application id")
	}
	relativePath, path, err := l.recordPath(tx.RequestID)
	if err != nil {
		return LedgerRecord{}, err
	}
	record := LedgerRecord{
		SchemaVersion:                   ledgerSchemaVersion,
		RecordType:                      "execution-transaction-ledger-record",
		Source:                          "go-runtime-state-root-execution-ledger",
		RequestID:                       tx.RequestID,
		ApplicationID:                   tx.ApplicationID,
		RelativePath:                    relativePath,
		Transaction:                     tx,
		PortalPermissionReceiptCount:    len(tx.PortalPermissionReceipts),
		PortalPermissionReceiptPaths:    portalPermissionReceiptPaths(tx.PortalPermissionReceipts),
		PortalPermissionReceiptStates:   portalPermissionReceiptStates(tx.PortalPermissionReceipts),
		PortalPermissionReceiptConsumed: len(tx.PortalPermissionReceipts) > 0,
		RuntimeOwned:                    true,
		GoRuntimeBacked:                 true,
		KDEPolicyOwner:                  false,
		StateRootPathExposed:            false,
		LaunchAllowed:                   false,
		LaunchEnabled:                   false,
		BackendStarted:                  false,
		PermissionGranted:               false,
		HostRootModified:                false,
		NetworkRequired:                 false,
		PrivilegedContainerRequired:     false,
		BackendDetailsExposed:           false,
		Summary:                         "Runtime recorded a blocked-by-default execution transaction under the configured state root without launching a backend.",
	}
	data, digest, err := marshalRecord(record)
	if err != nil {
		return LedgerRecord{}, err
	}
	record.SHA256 = digest
	data, _, err = marshalRecord(record)
	if err != nil {
		return LedgerRecord{}, err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return LedgerRecord{}, fmt.Errorf("prepare execution ledger directory: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return LedgerRecord{}, fmt.Errorf("write execution ledger record: %w", err)
	}
	return record, nil
}

func portalPermissionReceiptPaths(receipts []PortalPermissionReceipt) []string {
	paths := make([]string, 0, len(receipts))
	for _, receipt := range receipts {
		if receipt.RelativePath != "" && !filepath.IsAbs(receipt.RelativePath) && !strings.Contains(receipt.RelativePath, "..") {
			paths = append(paths, filepath.ToSlash(receipt.RelativePath))
		}
	}
	sort.Strings(paths)
	return paths
}

func portalPermissionReceiptStates(receipts []PortalPermissionReceipt) []string {
	states := make([]string, 0, len(receipts))
	for _, receipt := range receipts {
		if receipt.PermissionState == "" {
			continue
		}
		state := receipt.Operation + ":" + receipt.PermissionState
		if receipt.RequestState != "" {
			state += "/" + receipt.RequestState
		}
		states = append(states, state)
	}
	sort.Strings(states)
	return states
}

// Load returns one persisted transaction record by request id.
func (l *Ledger) Load(requestID string) (LedgerRecord, error) {
	relativePath, path, err := l.recordPath(requestID)
	if err != nil {
		return LedgerRecord{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return LedgerRecord{}, fmt.Errorf("read execution ledger record: %w", err)
	}
	var record LedgerRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return LedgerRecord{}, fmt.Errorf("parse execution ledger record: %w", err)
	}
	if record.RequestID != requestID {
		return LedgerRecord{}, fmt.Errorf("execution ledger record id mismatch: %q", record.RequestID)
	}
	if record.SchemaVersion != ledgerSchemaVersion || record.RecordType != "execution-transaction-ledger-record" || record.Source != "go-runtime-state-root-execution-ledger" {
		return LedgerRecord{}, errors.New("execution ledger record has unsupported schema")
	}
	if record.RelativePath != relativePath || record.ApplicationID == "" || record.Transaction.ApplicationID != record.ApplicationID || record.Transaction.RequestID != requestID {
		return LedgerRecord{}, errors.New("execution ledger record identity or path mismatch")
	}
	storedDigest := record.SHA256
	record.SHA256 = ""
	_, expectedDigest, err := marshalRecord(record)
	if err != nil {
		return LedgerRecord{}, err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return LedgerRecord{}, errors.New("execution ledger record digest mismatch")
	}
	record.SHA256 = storedDigest
	if record.StateRootPathExposed || record.LaunchAllowed || record.LaunchEnabled || record.BackendStarted || record.PermissionGranted || record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.BackendDetailsExposed ||
		record.Transaction.LaunchAllowed || record.Transaction.LaunchEnabled || record.Transaction.BackendStarted || record.Transaction.PermissionGranted || record.Transaction.HostRootModified || record.Transaction.NetworkRequired {
		return LedgerRecord{}, errors.New("execution ledger record has unsafe enabled gates")
	}
	return record, nil
}

// List returns every persisted transaction record in deterministic order.
func (l *Ledger) List() ([]LedgerRecord, error) {
	dir := filepath.Join(l.root, "execution-ledger", "transactions")
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("list execution ledger records: %w", err)
	}
	var records []LedgerRecord
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		requestID := strings.TrimSuffix(entry.Name(), ".json")
		record, err := l.Load(requestID)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].RequestID < records[j].RequestID
	})
	return records, nil
}

func (l *Ledger) recordPath(requestID string) (string, string, error) {
	if err := validateLedgerID(requestID); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("execution-ledger", "transactions", requestID+".json"))
	path := filepath.Join(l.root, filepath.FromSlash(relativePath))
	cleanRoot := filepath.Clean(l.root)
	rel, err := filepath.Rel(cleanRoot, path)
	if err != nil {
		return "", "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("execution ledger path escapes state root: %s", relativePath)
	}
	return relativePath, path, nil
}

func safeLedgerRoot(root string) (string, error) {
	if root == "" {
		return "", errors.New("execution ledger requires an explicit state root")
	}
	clean := filepath.Clean(root)
	if clean == string(os.PathSeparator) {
		return "", errors.New("refusing to use filesystem root as execution ledger state root")
	}
	if err := os.MkdirAll(clean, 0o700); err != nil {
		return "", fmt.Errorf("prepare execution ledger state root: %w", err)
	}
	return clean, nil
}

func validateLedgerID(value string) error {
	if value == "" {
		return errors.New("execution ledger id must not be empty")
	}
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			continue
		}
		return fmt.Errorf("execution ledger id contains unsafe character: %q", r)
	}
	if strings.Contains(value, "..") {
		return fmt.Errorf("execution ledger id must not contain path traversal: %q", value)
	}
	return nil
}

func marshalRecord(record LedgerRecord) ([]byte, string, error) {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}
