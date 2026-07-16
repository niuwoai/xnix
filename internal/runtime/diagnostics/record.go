package diagnostics

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

const runRecordSchemaVersion = "xnix.runtime.diagnostic_run_record.v1"

// RunRecordStore persists fixture-driven diagnostic run records under a
// caller-controlled Runtime state root. It never starts compatibility
// backends, calls real AI providers, executes repairs, or writes outside the
// configured root.
type RunRecordStore struct {
	root string
}

// RunRecordRequest describes one safe diagnostic run to record.
type RunRecordRequest struct {
	ApplicationID string
	RunID         string
	Fixture       Fixture
}

// RunRecord is the persisted, KDE-safe diagnostic run receipt.
type RunRecord struct {
	SchemaVersion               string                `json:"schema_version"`
	RecordType                  string                `json:"record_type"`
	Source                      string                `json:"source"`
	ApplicationID               string                `json:"application_id"`
	RunID                       string                `json:"run_id"`
	RelativePath                string                `json:"relative_path"`
	Result                      TestResult            `json:"result"`
	RepairRecommendation        *RepairRecommendation `json:"repair_recommendation,omitempty"`
	DiagnosticInput             DiagnosticInput       `json:"diagnostic_input"`
	SHA256                      string                `json:"sha256"`
	RuntimeOwned                bool                  `json:"runtime_owned"`
	GoRuntimeBacked             bool                  `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                  `json:"kde_policy_owner"`
	StateRootPathExposed        bool                  `json:"state_root_path_exposed"`
	FixturePathExposed          bool                  `json:"fixture_path_exposed"`
	BackendStarted              bool                  `json:"backend_started"`
	AIProviderCalled            bool                  `json:"ai_provider_called"`
	RealAIProviderEnabled       bool                  `json:"real_ai_provider_enabled"`
	AutoRepairAllowed           bool                  `json:"auto_repair_allowed"`
	RepairExecuted              bool                  `json:"repair_executed"`
	HostRootModified            bool                  `json:"host_root_modified"`
	NetworkRequired             bool                  `json:"network_required"`
	PrivilegedContainerRequired bool                  `json:"privileged_container_required"`
	BackendDetailsExposed       bool                  `json:"backend_details_exposed"`
	FileContentsIncluded        bool                  `json:"file_contents_included"`
	Summary                     string                `json:"summary"`
}

// NewRunRecordStore opens a diagnostics run ledger rooted under stateRoot.
func NewRunRecordStore(stateRoot string) (*RunRecordStore, error) {
	root, err := safeDiagnosticRecordRoot(stateRoot)
	if err != nil {
		return nil, err
	}
	return &RunRecordStore{root: root}, nil
}

// OpenRunRecordStoreReadOnly opens a diagnostics run ledger for read-only
// projections. It validates the supplied root but does not create directories
// or files, so preview-only commands can inspect existing evidence without
// mutating local state.
func OpenRunRecordStoreReadOnly(stateRoot string) (*RunRecordStore, error) {
	root, err := safeDiagnosticReadOnlyRoot(stateRoot)
	if err != nil {
		return nil, err
	}
	return &RunRecordStore{root: root}, nil
}

// Record runs the fixture, persists the result in the legacy result store, and
// writes a full diagnostic run receipt for audit and KDE consumption.
func (s *RunRecordStore) Record(request RunRecordRequest) (RunRecord, error) {
	if request.RunID == "" {
		return RunRecord{}, errors.New("diagnostic run record requires a run id")
	}
	if err := validateDiagnosticRecordID(request.RunID); err != nil {
		return RunRecord{}, err
	}

	result, err := Run(request.ApplicationID, request.Fixture)
	if err != nil {
		return RunRecord{}, err
	}
	diagnosticInput, err := BuildDiagnosticInput(result)
	if err != nil {
		return RunRecord{}, err
	}
	recommendation, ok := Recommend(result)
	var recommendationPtr *RepairRecommendation
	if ok {
		recommendationPtr = &recommendation
	}

	resultStore, err := NewResultStore(s.root)
	if err != nil {
		return RunRecord{}, err
	}
	if err := resultStore.Save(result); err != nil {
		return RunRecord{}, fmt.Errorf("persist diagnostic test result: %w", err)
	}

	relativePath, path, err := s.recordPath(request.RunID)
	if err != nil {
		return RunRecord{}, err
	}
	record := RunRecord{
		SchemaVersion:               runRecordSchemaVersion,
		RecordType:                  "diagnostic-run-record",
		Source:                      "go-runtime-state-root-diagnostic-run-record",
		ApplicationID:               result.ApplicationID,
		RunID:                       request.RunID,
		RelativePath:                relativePath,
		Result:                      result,
		RepairRecommendation:        recommendationPtr,
		DiagnosticInput:             diagnosticInput,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
		FixturePathExposed:          false,
		BackendStarted:              false,
		AIProviderCalled:            false,
		RealAIProviderEnabled:       false,
		AutoRepairAllowed:           false,
		RepairExecuted:              false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		FileContentsIncluded:        false,
		Summary:                     diagnosticRecordSummary(result, ok),
	}

	data, digest, err := marshalDiagnosticRunRecord(record)
	if err != nil {
		return RunRecord{}, err
	}
	record.SHA256 = digest
	data, _, err = marshalDiagnosticRunRecord(record)
	if err != nil {
		return RunRecord{}, err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return RunRecord{}, fmt.Errorf("prepare diagnostic run ledger directory: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return RunRecord{}, fmt.Errorf("write diagnostic run record: %w", err)
	}
	return record, nil
}

// Load returns one diagnostic run record by run id.
func (s *RunRecordStore) Load(runID string) (RunRecord, error) {
	relativePath, path, err := s.recordPath(runID)
	if err != nil {
		return RunRecord{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return RunRecord{}, fmt.Errorf("read diagnostic run record: %w", err)
	}
	var record RunRecord
	if err := json.Unmarshal(data, &record); err != nil {
		return RunRecord{}, fmt.Errorf("parse diagnostic run record: %w", err)
	}
	if record.RunID != runID {
		return RunRecord{}, fmt.Errorf("diagnostic run record id mismatch: %q", record.RunID)
	}
	if record.SchemaVersion != runRecordSchemaVersion || record.RecordType != "diagnostic-run-record" || record.Source != "go-runtime-state-root-diagnostic-run-record" {
		return RunRecord{}, errors.New("diagnostic run record has unsupported schema")
	}
	if record.RelativePath != relativePath || record.ApplicationID == "" || record.Result.ApplicationID != record.ApplicationID {
		return RunRecord{}, errors.New("diagnostic run record identity or path mismatch")
	}
	storedDigest := record.SHA256
	record.SHA256 = ""
	_, expectedDigest, err := marshalDiagnosticRunRecord(record)
	if err != nil {
		return RunRecord{}, err
	}
	if storedDigest == "" || storedDigest != expectedDigest {
		return RunRecord{}, errors.New("diagnostic run record digest mismatch")
	}
	record.SHA256 = storedDigest
	if record.StateRootPathExposed || record.FixturePathExposed || record.BackendStarted || record.AIProviderCalled || record.RealAIProviderEnabled || record.AutoRepairAllowed || record.RepairExecuted || record.HostRootModified || record.NetworkRequired || record.PrivilegedContainerRequired || record.BackendDetailsExposed || record.FileContentsIncluded || record.Result.BackendStarted || record.Result.HostRootModified || record.DiagnosticInput.NetworkRequired || record.DiagnosticInput.FileContentsIncluded {
		return RunRecord{}, errors.New("diagnostic run record has unsafe enabled gates")
	}
	return record, nil
}

// List returns every persisted diagnostic run record in deterministic order.
func (s *RunRecordStore) List() ([]RunRecord, error) {
	dir := filepath.Join(s.root, "diagnostics-ledger", "runs")
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("list diagnostic run records: %w", err)
	}
	var records []RunRecord
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		runID := strings.TrimSuffix(entry.Name(), ".json")
		record, err := s.Load(runID)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].RunID < records[j].RunID
	})
	return records, nil
}

// listLenient lists persisted diagnostic run records, skipping files that fail
// to parse and returning their run ids as malformed. Unlike List it never
// creates the ledger directory and never fails on a single corrupt receipt, so
// it is safe for read-only previews over possibly-corrupt evidence.
func (s *RunRecordStore) listLenient() ([]RunRecord, []string, error) {
	dir := filepath.Join(s.root, "diagnostics-ledger", "runs")
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("list diagnostic run records: %w", err)
	}
	var records []RunRecord
	var malformed []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		runID := strings.TrimSuffix(entry.Name(), ".json")
		record, err := s.Load(runID)
		if err != nil {
			malformed = append(malformed, runID)
			continue
		}
		records = append(records, record)
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].RunID < records[j].RunID
	})
	sort.Strings(malformed)
	return records, malformed, nil
}

func (s *RunRecordStore) recordPath(runID string) (string, string, error) {
	if err := validateDiagnosticRecordID(runID); err != nil {
		return "", "", err
	}
	relativePath := filepath.ToSlash(filepath.Join("diagnostics-ledger", "runs", runID+".json"))
	path := filepath.Join(s.root, filepath.FromSlash(relativePath))
	cleanRoot := filepath.Clean(s.root)
	rel, err := filepath.Rel(cleanRoot, path)
	if err != nil {
		return "", "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", "", fmt.Errorf("diagnostic run record path escapes state root: %s", relativePath)
	}
	return relativePath, path, nil
}

func safeDiagnosticRecordRoot(root string) (string, error) {
	if root == "" {
		return "", errors.New("diagnostic run record requires an explicit state root")
	}
	clean := filepath.Clean(root)
	if clean == string(os.PathSeparator) {
		return "", errors.New("refusing to use filesystem root as diagnostic run state root")
	}
	if err := os.MkdirAll(clean, 0o700); err != nil {
		return "", fmt.Errorf("prepare diagnostic run state root: %w", err)
	}
	return clean, nil
}

func safeDiagnosticReadOnlyRoot(root string) (string, error) {
	if root == "" {
		return "", errors.New("diagnostic run history read requires an explicit state root")
	}
	clean := filepath.Clean(root)
	if clean == string(os.PathSeparator) {
		return "", errors.New("refusing to use filesystem root as diagnostic run state root")
	}
	return clean, nil
}

func validateDiagnosticRecordID(value string) error {
	if value == "" {
		return errors.New("diagnostic run id must not be empty")
	}
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' || r == '.' {
			continue
		}
		return fmt.Errorf("diagnostic run id contains unsafe character: %q", r)
	}
	if strings.Contains(value, "..") {
		return fmt.Errorf("diagnostic run id must not contain path traversal: %q", value)
	}
	return nil
}

func diagnosticRecordSummary(result TestResult, hasRecommendation bool) string {
	if hasRecommendation {
		return "Runtime recorded fixture diagnostics and a review-first repair recommendation under the configured state root without calling an AI provider."
	}
	if result.Overall == OutcomePass {
		return "Runtime recorded healthy fixture diagnostics under the configured state root without calling an AI provider."
	}
	return "Runtime recorded fixture diagnostics under the configured state root without calling an AI provider."
}

func marshalDiagnosticRunRecord(record RunRecord) ([]byte, string, error) {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return nil, "", err
	}
	data = append(data, '\n')
	sum := sha256.Sum256(data)
	return data, hex.EncodeToString(sum[:]), nil
}
