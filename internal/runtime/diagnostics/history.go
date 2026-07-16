package diagnostics

import (
	"errors"
	"sort"

	"xnix.local/xnix/internal/runtime/appid"
)

const runHistorySchemaVersion = "xnix.runtime.diagnostic_run_history.v1"

// RunHistory summarizes persisted diagnostic runs for KDE and Compatibility
// Center consumers. It contains only relative receipt paths and safe derived
// fields, never local roots, file contents, backend commands, or provider data.
type RunHistory struct {
	SchemaVersion               string             `json:"schema_version"`
	RecordType                  string             `json:"record_type"`
	Source                      string             `json:"source"`
	ApplicationID               string             `json:"application_id,omitempty"`
	Records                     []RunHistoryRecord `json:"records"`
	Counts                      RunHistoryCounts   `json:"counts"`
	Latest                      *RunHistoryRecord  `json:"latest,omitempty"`
	RuntimeOwned                bool               `json:"runtime_owned"`
	GoRuntimeBacked             bool               `json:"go_runtime_backed"`
	KDEPolicyOwner              bool               `json:"kde_policy_owner"`
	StateRootPathExposed        bool               `json:"state_root_path_exposed"`
	BackendStarted              bool               `json:"backend_started"`
	AIProviderCalled            bool               `json:"ai_provider_called"`
	RealAIProviderEnabled       bool               `json:"real_ai_provider_enabled"`
	AutoRepairAllowed           bool               `json:"auto_repair_allowed"`
	RepairExecuted              bool               `json:"repair_executed"`
	HostRootModified            bool               `json:"host_root_modified"`
	NetworkRequired             bool               `json:"network_required"`
	PrivilegedContainerRequired bool               `json:"privileged_container_required"`
	BackendDetailsExposed       bool               `json:"backend_details_exposed"`
	FileContentsIncluded        bool               `json:"file_contents_included"`
	Summary                     string             `json:"summary"`
}

// RunHistoryRecord is the compact, desktop-safe projection of one run receipt.
type RunHistoryRecord struct {
	RunID                 string   `json:"run_id"`
	ApplicationID         string   `json:"application_id"`
	TestType              string   `json:"test_type"`
	Overall               Outcome  `json:"overall"`
	RelativePath          string   `json:"relative_path"`
	FailingIDs            []string `json:"failing_ids"`
	RepairIssue           string   `json:"repair_issue,omitempty"`
	SnapshotRequired      bool     `json:"snapshot_required"`
	AutoRepairAllowed     bool     `json:"auto_repair_allowed"`
	BackendStarted        bool     `json:"backend_started"`
	AIProviderCalled      bool     `json:"ai_provider_called"`
	RepairExecuted        bool     `json:"repair_executed"`
	HostRootModified      bool     `json:"host_root_modified"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
	FileContentsIncluded  bool     `json:"file_contents_included"`
}

// RunHistoryCounts groups diagnostic run outcomes.
type RunHistoryCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Pending int `json:"pending"`
	Blocked int `json:"blocked"`
	Failed  int `json:"failed"`
}

// History returns a KDE-safe summary of diagnostic run records. If
// applicationID is empty, it summarizes every stored run.
func (s *RunRecordStore) History(applicationID string) (RunHistory, error) {
	if applicationID != "" && !appid.Valid(applicationID) {
		return RunHistory{}, errBadHistoryApplicationID
	}
	records, err := s.List()
	if err != nil {
		return RunHistory{}, err
	}
	return s.buildHistory(applicationID, records), nil
}

// LenientHistory is History for read-only previews over possibly-corrupt
// evidence: records that fail to parse are skipped and their run ids are
// returned as malformed instead of failing the whole read. It never creates the
// ledger directory or mutates state.
func (s *RunRecordStore) LenientHistory(applicationID string) (RunHistory, []string, error) {
	if applicationID != "" && !appid.Valid(applicationID) {
		return RunHistory{}, nil, errBadHistoryApplicationID
	}
	records, malformed, err := s.listLenient()
	if err != nil {
		return RunHistory{}, nil, err
	}
	return s.buildHistory(applicationID, records), malformed, nil
}

func (s *RunRecordStore) buildHistory(applicationID string, records []RunRecord) RunHistory {
	historyRecords := make([]RunHistoryRecord, 0, len(records))
	var counts RunHistoryCounts
	for _, record := range records {
		if applicationID != "" && record.ApplicationID != applicationID {
			continue
		}
		item := historyRecord(record)
		historyRecords = append(historyRecords, item)
		counts.Total++
		switch item.Overall {
		case OutcomePass:
			counts.Passed++
		case OutcomePending:
			counts.Pending++
		case OutcomeBlocked:
			counts.Blocked++
		case OutcomeFail:
			counts.Failed++
		}
	}
	sort.Slice(historyRecords, func(i, j int) bool {
		return historyRecords[i].RunID < historyRecords[j].RunID
	})

	var latest *RunHistoryRecord
	if len(historyRecords) > 0 {
		copyOfLatest := historyRecords[len(historyRecords)-1]
		latest = &copyOfLatest
	}
	return RunHistory{
		SchemaVersion:               runHistorySchemaVersion,
		RecordType:                  "diagnostic-run-history",
		Source:                      "go-runtime-state-root-diagnostic-run-history",
		ApplicationID:               applicationID,
		Records:                     historyRecords,
		Counts:                      counts,
		Latest:                      latest,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		StateRootPathExposed:        false,
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
		Summary:                     runHistorySummary(counts),
	}
}

func historyRecord(record RunRecord) RunHistoryRecord {
	item := RunHistoryRecord{
		RunID:                 record.RunID,
		ApplicationID:         record.ApplicationID,
		TestType:              record.Result.TestType,
		Overall:               record.Result.Overall,
		RelativePath:          record.RelativePath,
		FailingIDs:            append([]string(nil), record.Result.FailingIDs...),
		AutoRepairAllowed:     false,
		BackendStarted:        false,
		AIProviderCalled:      false,
		RepairExecuted:        false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		FileContentsIncluded:  false,
	}
	if record.RepairRecommendation != nil {
		item.RepairIssue = record.RepairRecommendation.Issue
		item.SnapshotRequired = record.RepairRecommendation.SnapshotRequired
	}
	return item
}

func runHistorySummary(counts RunHistoryCounts) string {
	if counts.Total == 0 {
		return "Runtime has no persisted diagnostic runs under the configured state root."
	}
	if counts.Failed > 0 {
		return "Runtime diagnostic history contains failing runs for Compatibility Center review."
	}
	if counts.Blocked > 0 {
		return "Runtime diagnostic history contains blocked runs awaiting user or Runtime action."
	}
	return "Runtime diagnostic history contains no failing runs."
}

var errBadHistoryApplicationID = errors.New("diagnostic history application id must be a reverse-DNS identifier")
