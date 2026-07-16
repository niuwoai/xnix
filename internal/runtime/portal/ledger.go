package portal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Record captures one persisted Portal request transition under a caller-owned
// state root. It never performs real Portal transport calls and never exposes
// the absolute state root to KDE-facing readers.
type Record struct {
	SchemaVersion               string   `json:"schema_version"`
	RecordType                  string   `json:"record_type"`
	Source                      string   `json:"source"`
	Action                      string   `json:"action"`
	RelativePath                string   `json:"relative_path"`
	Request                     Request  `json:"request"`
	RuntimeOwned                bool     `json:"runtime_owned"`
	GoRuntimeBacked             bool     `json:"go_runtime_backed"`
	KDEPolicyOwner              bool     `json:"kde_policy_owner"`
	StateRootPathExposed        bool     `json:"state_root_path_exposed"`
	RealPortalCallEnabled       bool     `json:"real_portal_call_enabled"`
	RequestObjectCreated        bool     `json:"request_object_created"`
	PermissionGranted           bool     `json:"permission_granted"`
	ExecutionApproved           bool     `json:"execution_approved"`
	HostPermissionChanged       bool     `json:"host_permission_changed"`
	HostRootModified            bool     `json:"host_root_modified"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	NetworkRequired             bool     `json:"network_required"`
	PrivilegedContainerRequired bool     `json:"privileged_container_required"`
	Diagnostics                 []string `json:"diagnostics"`
	Summary                     string   `json:"summary"`
}

// Ledger persists fake-mode Portal requests under an explicit state root.
type Ledger struct {
	root string
	dir  string
}

// NewLedger opens a state-root scoped Portal request ledger.
func NewLedger(stateRoot string) (*Ledger, error) {
	if stateRoot == "" {
		return nil, errors.New("portal request ledger requires a state root")
	}
	abs, err := filepath.Abs(stateRoot)
	if err != nil {
		return nil, fmt.Errorf("resolve state root: %w", err)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("state root must exist: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("state root %q is not a directory", abs)
	}
	dir := filepath.Join(abs, "portal-requests")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("initialize portal request ledger: %w", err)
	}
	return &Ledger{root: abs, dir: dir}, nil
}

func RequestRelativePath(handleToken string) (string, error) {
	if strings.TrimSpace(handleToken) == "" {
		return "", errors.New("portal request handle token is required")
	}
	return filepath.ToSlash(filepath.Join("portal-requests", sanitizeToken(handleToken)+".json")), nil
}

func (l *Ledger) Create(spec RequestSpec) (Record, error) {
	sequence, err := l.nextSequence()
	if err != nil {
		return Record{}, err
	}
	request, err := newRequest(spec, sequence)
	if err != nil {
		return Record{}, err
	}
	if err := l.save(*request); err != nil {
		return Record{}, err
	}
	return newLedgerRecord("create", *request)
}

func (l *Ledger) Inspect(handleToken string) (Record, error) {
	request, err := l.load(handleToken)
	if err != nil {
		return Record{}, err
	}
	return newLedgerRecord("inspect", request)
}

func (l *Ledger) Resolve(handleToken string, outcome Outcome) (Record, error) {
	request, err := l.load(handleToken)
	if err != nil {
		return Record{}, err
	}
	broker := NewFakeBroker()
	broker.requests[request.HandleToken] = &request
	resolved, err := broker.Resolve(request.HandleToken, outcome)
	if err != nil {
		return Record{}, err
	}
	if err := l.save(*resolved); err != nil {
		return Record{}, err
	}
	return newLedgerRecord("resolve", *resolved)
}

func (l *Ledger) Complete(handleToken string) (Record, error) {
	request, err := l.load(handleToken)
	if err != nil {
		return Record{}, err
	}
	broker := NewFakeBroker()
	broker.requests[request.HandleToken] = &request
	completed, err := broker.Complete(request.HandleToken)
	if err != nil {
		return Record{}, err
	}
	if err := l.save(*completed); err != nil {
		return Record{}, err
	}
	return newLedgerRecord("complete", *completed)
}

func (l *Ledger) Cancel(handleToken string) (Record, error) {
	request, err := l.load(handleToken)
	if err != nil {
		return Record{}, err
	}
	broker := NewFakeBroker()
	broker.requests[request.HandleToken] = &request
	cancelled, err := broker.Cancel(request.HandleToken)
	if err != nil {
		return Record{}, err
	}
	if err := l.save(*cancelled); err != nil {
		return Record{}, err
	}
	return newLedgerRecord("cancel", *cancelled)
}

func (l *Ledger) Expire(handleToken string) (Record, error) {
	request, err := l.load(handleToken)
	if err != nil {
		return Record{}, err
	}
	broker := NewFakeBroker()
	broker.requests[request.HandleToken] = &request
	expired, err := broker.Expire(request.HandleToken)
	if err != nil {
		return Record{}, err
	}
	if err := l.save(*expired); err != nil {
		return Record{}, err
	}
	return newLedgerRecord("expire", *expired)
}

func (l *Ledger) nextSequence() (int, error) {
	entries, err := os.ReadDir(l.dir)
	if err != nil {
		return 0, fmt.Errorf("list portal request ledger: %w", err)
	}
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			count++
		}
	}
	return count + 1, nil
}

func (l *Ledger) path(handleToken string) (string, error) {
	relativePath, err := RequestRelativePath(handleToken)
	if err != nil {
		return "", err
	}
	return filepath.Join(l.root, filepath.FromSlash(relativePath)), nil
}

func (l *Ledger) save(request Request) error {
	request.DirectAccessAllowed = false
	request.HostPermissionChanged = false
	request.BackendDetailsExposed = false
	path, err := l.path(request.HandleToken)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

func (l *Ledger) load(handleToken string) (Request, error) {
	path, err := l.path(handleToken)
	if err != nil {
		return Request{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Request{}, fmt.Errorf("portal request %q is not recorded", handleToken)
		}
		return Request{}, fmt.Errorf("read portal request record: %w", err)
	}
	var request Request
	if err := json.Unmarshal(data, &request); err != nil {
		return Request{}, fmt.Errorf("portal request record is corrupt: %w", err)
	}
	if request.HandleToken != handleToken {
		return Request{}, fmt.Errorf("portal request record handle mismatch: %s", request.HandleToken)
	}
	return request, nil
}

func newLedgerRecord(action string, request Request) (Record, error) {
	relativePath, err := RequestRelativePath(request.HandleToken)
	if err != nil {
		return Record{}, err
	}
	return Record{
		SchemaVersion:               "xnix.runtime.portal_request_record.v1",
		RecordType:                  "portal-permission-request-record",
		Source:                      "go-runtime-state-root-portal-broker",
		Action:                      action,
		RelativePath:                relativePath,
		Request:                     request,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		RealPortalCallEnabled:       false,
		RequestObjectCreated:        true,
		PermissionGranted:           request.PermissionState == PermissionGranted,
		ExecutionApproved:           false,
		HostPermissionChanged:       false,
		HostRootModified:            false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		Diagnostics:                 append([]string{}, request.Diagnostics...),
		Summary:                     "Runtime persisted a fake-mode Portal permission request under the configured state root without making real Portal calls, granting execution, exposing paths, or mutating the host root.",
	}, nil
}
