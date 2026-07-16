package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const stateRootQuotaRetentionSchemaVersion = "xnix.runtime.state_root_quota_retention.v1"

type StateRootQuotaRetentionOptions struct {
	ApplicationID string
	StateRoot     string
	QuotaBytes    int64
	ActiveSession bool
}

type StateRootQuotaRetentionPreview struct {
	SchemaVersion               string                           `json:"schema_version"`
	RequestType                 string                           `json:"request_type"`
	PreviewType                 string                           `json:"preview_type"`
	Source                      string                           `json:"source"`
	Desktop                     string                           `json:"desktop"`
	RuntimeMethod               string                           `json:"runtime_method"`
	ReadMethod                  string                           `json:"read_method"`
	ApplicationID               string                           `json:"application_id"`
	StateRootPresent            bool                             `json:"state_root_present"`
	StateRootRelativeID         string                           `json:"state_root_relative_id"`
	QuotaBytes                  int64                            `json:"quota_bytes"`
	EstimatedBytes              int64                            `json:"estimated_bytes"`
	OverQuota                   bool                             `json:"over_quota"`
	ActiveSession               bool                             `json:"active_session"`
	Sections                    []StateRootQuotaRetentionSection `json:"sections"`
	SectionIDs                  []string                         `json:"section_ids"`
	SectionCount                int                              `json:"section_count"`
	Counts                      StateRootQuotaRetentionCounts    `json:"counts"`
	RetentionReasonCodes        []string                         `json:"retention_reason_codes"`
	BlockedCleanupReasons       []string                         `json:"blocked_cleanup_reasons"`
	UserSafeRecommendations     []string                         `json:"user_safe_recommendations"`
	RuntimeOwned                bool                             `json:"runtime_owned"`
	GoRuntimeBacked             bool                             `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                             `json:"kde_policy_owner"`
	ReviewOnly                  bool                             `json:"review_only"`
	FileDeletionEnabled         bool                             `json:"file_deletion_enabled"`
	DirectoriesCreated          bool                             `json:"directories_created"`
	LogTruncationEnabled        bool                             `json:"log_truncation_enabled"`
	ReceiptsRewritten           bool                             `json:"receipts_rewritten"`
	SnapshotDeletionEnabled     bool                             `json:"snapshot_deletion_enabled"`
	StateRootPathExposed        bool                             `json:"state_root_path_exposed"`
	HostRootModified            bool                             `json:"host_root_modified"`
	BackendDetailsExposed       bool                             `json:"backend_details_exposed"`
	NetworkRequired             bool                             `json:"network_required"`
	PrivilegedContainerRequired bool                             `json:"privileged_container_required"`
	DesktopSafeSummary          string                           `json:"desktop_safe_summary"`
}

type StateRootQuotaRetentionSection struct {
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	RecordCount             int      `json:"record_count"`
	EstimatedBytes          int64    `json:"estimated_bytes"`
	EvidenceIDs             []string `json:"evidence_ids"`
	RetentionReasonCodes    []string `json:"retention_reason_codes"`
	BlockedCleanupReasons   []string `json:"blocked_cleanup_reasons"`
	UserSafeRecommendation  string   `json:"user_safe_recommendation"`
	CleanupCandidate        bool     `json:"cleanup_candidate"`
	RetentionExempt         bool     `json:"retention_exempt"`
	MalformedRecordCount    int      `json:"malformed_record_count"`
	UnknownRecordCount      int      `json:"unknown_record_count"`
	FileDeletionEnabled     bool     `json:"file_deletion_enabled"`
	DirectoriesCreated      bool     `json:"directories_created"`
	LogTruncationEnabled    bool     `json:"log_truncation_enabled"`
	ReceiptsRewritten       bool     `json:"receipts_rewritten"`
	SnapshotDeletionEnabled bool     `json:"snapshot_deletion_enabled"`
	StateRootPathExposed    bool     `json:"state_root_path_exposed"`
	HostRootModified        bool     `json:"host_root_modified"`
}

type StateRootQuotaRetentionCounts struct {
	TotalRecords       int `json:"total_records"`
	CleanupCandidates  int `json:"cleanup_candidates"`
	RetentionExempt    int `json:"retention_exempt"`
	MalformedRecords   int `json:"malformed_records"`
	UnknownRecords     int `json:"unknown_records"`
	BlockedSections    int `json:"blocked_sections"`
	OverQuotaSections  int `json:"over_quota_sections"`
	UnderQuotaSections int `json:"under_quota_sections"`
}

type stateRootRetentionRecord struct {
	sectionID         string
	evidenceID        string
	estimatedBytes    int64
	malformed         bool
	unknown           bool
	retentionExempt   bool
	cleanupCandidate  bool
	blockedReasonCode string
}

func NewStateRootQuotaRetentionPreview(options StateRootQuotaRetentionOptions) (StateRootQuotaRetentionPreview, error) {
	normalized, err := normalizeStateRootQuotaRetentionOptions(options)
	if err != nil {
		return StateRootQuotaRetentionPreview{}, err
	}

	base := newStateRootQuotaRetentionPreviewBase(normalized)
	info, err := os.Stat(normalized.StateRoot)
	if err != nil {
		if os.IsNotExist(err) {
			base.Sections = missingStateRootQuotaRetentionSections()
			base.SectionIDs = stateRootQuotaRetentionSectionIDs(base.Sections)
			base.SectionCount = len(base.Sections)
			base.Counts = countStateRootQuotaRetentionSections(base.Sections, false)
			base.RetentionReasonCodes = quotaRetentionUniqueStrings([]string{"retain-missing-state-root"})
			base.BlockedCleanupReasons = quotaRetentionUniqueStrings([]string{"missing-state-root"})
			base.UserSafeRecommendations = []string{"No cleanup is available because the Runtime state record area is not present."}
			base.DesktopSafeSummary = "Runtime state records are not present yet, so there is nothing to clean up."
			return validateStateRootQuotaRetentionPreview(base)
		}
		return StateRootQuotaRetentionPreview{}, fmt.Errorf("inspect state root for quota retention preview: %w", err)
	}
	if !info.IsDir() {
		return StateRootQuotaRetentionPreview{}, errors.New("state-root-quota-retention-preview requires --state-root to point to a directory")
	}

	records, err := scanStateRootQuotaRetentionRecords(normalized)
	if err != nil {
		return StateRootQuotaRetentionPreview{}, err
	}
	base.StateRootPresent = true
	base.Sections = buildStateRootQuotaRetentionSections(records, normalized)
	base.SectionIDs = stateRootQuotaRetentionSectionIDs(base.Sections)
	base.SectionCount = len(base.Sections)
	base.EstimatedBytes = stateRootQuotaRetentionEstimatedBytes(base.Sections)
	base.OverQuota = normalized.QuotaBytes > 0 && base.EstimatedBytes > normalized.QuotaBytes
	base.Counts = countStateRootQuotaRetentionSections(base.Sections, base.OverQuota)
	base.RetentionReasonCodes = stateRootQuotaRetentionReasonCodes(base.Sections)
	base.BlockedCleanupReasons = stateRootQuotaRetentionBlockedReasons(base.Sections)
	base.UserSafeRecommendations = stateRootQuotaRetentionRecommendations(base.Sections, base.OverQuota)
	base.DesktopSafeSummary = stateRootQuotaRetentionSummary(base)
	return validateStateRootQuotaRetentionPreview(base)
}

func normalizeStateRootQuotaRetentionOptions(options StateRootQuotaRetentionOptions) (StateRootQuotaRetentionOptions, error) {
	options.ApplicationID = strings.TrimSpace(options.ApplicationID)
	if !idPattern.MatchString(options.ApplicationID) {
		return StateRootQuotaRetentionOptions{}, errors.New("application id must be a reverse-DNS identifier")
	}
	options.StateRoot = strings.TrimSpace(options.StateRoot)
	if options.StateRoot == "" {
		return StateRootQuotaRetentionOptions{}, errors.New("state-root-quota-retention-preview requires an explicit state root")
	}
	clean := filepath.Clean(options.StateRoot)
	if clean == string(os.PathSeparator) {
		return StateRootQuotaRetentionOptions{}, errors.New("refusing to inspect filesystem root as Runtime state root")
	}
	options.StateRoot = clean
	if options.QuotaBytes < 0 {
		return StateRootQuotaRetentionOptions{}, errors.New("quota bytes must not be negative")
	}
	return options, nil
}

func newStateRootQuotaRetentionPreviewBase(options StateRootQuotaRetentionOptions) StateRootQuotaRetentionPreview {
	return StateRootQuotaRetentionPreview{
		SchemaVersion:               stateRootQuotaRetentionSchemaVersion,
		RequestType:                 "state-root-quota-retention-preview",
		PreviewType:                 "dry-run-state-root-retention",
		Source:                      "go-runtime-state-root-quota-retention-preview",
		Desktop:                     "KDE Plasma",
		RuntimeMethod:               "GetStateRootQuotaRetention",
		ReadMethod:                  "GetStateRootQuotaRetentionPreview",
		ApplicationID:               options.ApplicationID,
		StateRootPresent:            false,
		StateRootRelativeID:         "runtime-state-root",
		QuotaBytes:                  options.QuotaBytes,
		ActiveSession:               options.ActiveSession,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		ReviewOnly:                  true,
		FileDeletionEnabled:         false,
		DirectoriesCreated:          false,
		LogTruncationEnabled:        false,
		ReceiptsRewritten:           false,
		SnapshotDeletionEnabled:     false,
		StateRootPathExposed:        false,
		HostRootModified:            false,
		BackendDetailsExposed:       false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
	}
}

func missingStateRootQuotaRetentionSections() []StateRootQuotaRetentionSection {
	section := newStateRootQuotaRetentionSection("missing-state-root")
	section.Title = "Missing state records"
	section.RetentionReasonCodes = []string{"retain-missing-state-root"}
	section.BlockedCleanupReasons = []string{"missing-state-root"}
	section.UserSafeRecommendation = "Wait until the Runtime records application state before offering cleanup."
	return []StateRootQuotaRetentionSection{section}
}

func scanStateRootQuotaRetentionRecords(options StateRootQuotaRetentionOptions) ([]stateRootRetentionRecord, error) {
	var records []stateRootRetentionRecord
	err := filepath.WalkDir(options.StateRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walk state root evidence: %w", walkErr)
		}
		if path == options.StateRoot {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			relativeID, err := stateRootRelativeEvidenceID(options.StateRoot, path)
			if err != nil {
				return err
			}
			sectionID := classifyStateRootEvidence(relativeID)
			records = append(records, stateRootRetentionRecord{
				sectionID:         sectionID,
				evidenceID:        evidenceIDForStateRootSection(sectionID, relativeID),
				estimatedBytes:    0,
				malformed:         false,
				unknown:           sectionID == "unknown-records",
				retentionExempt:   false,
				cleanupCandidate:  false,
				blockedReasonCode: "symlink-not-followed",
			})
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("inspect state root evidence: %w", err)
		}
		relativeID, err := stateRootRelativeEvidenceID(options.StateRoot, path)
		if err != nil {
			return err
		}
		sectionID := classifyStateRootEvidence(relativeID)
		record := stateRootRetentionRecord{
			sectionID:      sectionID,
			evidenceID:     evidenceIDForStateRootSection(sectionID, relativeID),
			estimatedBytes: info.Size(),
			unknown:        sectionID == "unknown-records",
		}
		if strings.HasSuffix(strings.ToLower(relativeID), ".json") {
			metadata, malformed, err := readStateRootRecordMetadata(path)
			if err != nil {
				return err
			}
			record.malformed = malformed
			record.retentionExempt = stateRootRetentionExempt(metadata)
			if stateRootRecordEstimatedBytes(metadata) > 0 {
				record.estimatedBytes = stateRootRecordEstimatedBytes(metadata)
			}
		}
		record.cleanupCandidate = stateRootRecordCleanupCandidate(record, options)
		record.blockedReasonCode = stateRootRecordBlockedReason(record, options)
		records = append(records, record)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].evidenceID < records[j].evidenceID
	})
	return records, nil
}

func stateRootRelativeEvidenceID(root string, path string) (string, error) {
	relative, err := filepath.Rel(root, path)
	if err != nil {
		return "", fmt.Errorf("derive state root relative evidence id: %w", err)
	}
	if relative == "." || relative == ".." || strings.HasPrefix(relative, ".."+string(os.PathSeparator)) || filepath.IsAbs(relative) {
		return "", errors.New("state root evidence path escapes the configured root")
	}
	return filepath.ToSlash(relative), nil
}

func classifyStateRootEvidence(relativeID string) string {
	lower := strings.ToLower(relativeID)
	switch {
	case strings.Contains(lower, "snapshot"):
		return "snapshots"
	case strings.Contains(lower, "diagnostic"):
		return "diagnostics"
	case strings.Contains(lower, "execution"):
		return "execution-receipts"
	case strings.Contains(lower, "portal"):
		return "portal-receipts"
	case strings.Contains(lower, "artifact"):
		return "artifact-receipts"
	case strings.Contains(lower, "activation") || strings.Contains(lower, "desktop"):
		return "activation-receipts"
	default:
		return "unknown-records"
	}
}

func evidenceIDForStateRootSection(sectionID string, relativeID string) string {
	return sectionID + ":" + relativeID
}

func readStateRootRecordMetadata(path string) (map[string]any, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false, fmt.Errorf("read state root evidence metadata: %w", err)
	}
	var metadata map[string]any
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, true, nil
	}
	return metadata, false, nil
}

func stateRootRetentionExempt(metadata map[string]any) bool {
	if metadata == nil {
		return false
	}
	for _, key := range []string{"retention_exempt", "retentionExempt"} {
		value, ok := metadata[key].(bool)
		if ok && value {
			return true
		}
	}
	if retention, ok := metadata["retention"].(map[string]any); ok {
		if value, ok := retention["exempt"].(bool); ok && value {
			return true
		}
	}
	return false
}

func stateRootRecordEstimatedBytes(metadata map[string]any) int64 {
	if metadata == nil {
		return 0
	}
	for _, key := range []string{"estimated_bytes", "estimatedBytes", "size_bytes", "sizeBytes"} {
		switch value := metadata[key].(type) {
		case float64:
			if value > 0 {
				return int64(value)
			}
		case int64:
			if value > 0 {
				return value
			}
		case json.Number:
			parsed, err := value.Int64()
			if err == nil && parsed > 0 {
				return parsed
			}
		}
	}
	return 0
}

func stateRootRecordCleanupCandidate(record stateRootRetentionRecord, options StateRootQuotaRetentionOptions) bool {
	if record.malformed || record.unknown || record.retentionExempt {
		return false
	}
	if options.ActiveSession && record.sectionID == "execution-receipts" {
		return false
	}
	return record.sectionID == "diagnostics" || record.sectionID == "snapshots" || record.sectionID == "execution-receipts" || record.sectionID == "portal-receipts"
}

func stateRootRecordBlockedReason(record stateRootRetentionRecord, options StateRootQuotaRetentionOptions) string {
	switch {
	case record.retentionExempt:
		return "retention-exempt"
	case record.malformed:
		return "malformed-record-review-required"
	case record.unknown:
		return "unknown-record-review-required"
	case options.ActiveSession && record.sectionID == "execution-receipts":
		return "active-session-blocked"
	case record.blockedReasonCode != "":
		return record.blockedReasonCode
	default:
		return ""
	}
}

func buildStateRootQuotaRetentionSections(records []stateRootRetentionRecord, options StateRootQuotaRetentionOptions) []StateRootQuotaRetentionSection {
	sections := map[string]StateRootQuotaRetentionSection{}
	for _, sectionID := range stateRootQuotaRetentionSectionOrder() {
		sections[sectionID] = newStateRootQuotaRetentionSection(sectionID)
	}
	for _, record := range records {
		section := sections[record.sectionID]
		section.RecordCount++
		section.EstimatedBytes += record.estimatedBytes
		section.EvidenceIDs = append(section.EvidenceIDs, record.evidenceID)
		if record.malformed {
			section.MalformedRecordCount++
		}
		if record.unknown {
			section.UnknownRecordCount++
		}
		if record.retentionExempt {
			section.RetentionExempt = true
		}
		if record.cleanupCandidate {
			section.CleanupCandidate = true
		}
		if record.blockedReasonCode != "" {
			section.BlockedCleanupReasons = append(section.BlockedCleanupReasons, record.blockedReasonCode)
		}
		sections[record.sectionID] = section
	}

	result := make([]StateRootQuotaRetentionSection, 0, len(sections))
	for _, sectionID := range stateRootQuotaRetentionSectionOrder() {
		section := sections[sectionID]
		section.EvidenceIDs = quotaRetentionUniqueStrings(section.EvidenceIDs)
		section.BlockedCleanupReasons = quotaRetentionUniqueStrings(section.BlockedCleanupReasons)
		section.RetentionReasonCodes = stateRootQuotaRetentionSectionReasons(section, options)
		section.UserSafeRecommendation = stateRootQuotaRetentionSectionRecommendation(section, options)
		result = append(result, section)
	}
	return result
}

func newStateRootQuotaRetentionSection(id string) StateRootQuotaRetentionSection {
	return StateRootQuotaRetentionSection{
		ID:                      id,
		Title:                   stateRootQuotaRetentionSectionTitle(id),
		FileDeletionEnabled:     false,
		DirectoriesCreated:      false,
		LogTruncationEnabled:    false,
		ReceiptsRewritten:       false,
		SnapshotDeletionEnabled: false,
		StateRootPathExposed:    false,
		HostRootModified:        false,
	}
}

func stateRootQuotaRetentionSectionOrder() []string {
	return []string{
		"snapshots",
		"diagnostics",
		"execution-receipts",
		"portal-receipts",
		"artifact-receipts",
		"activation-receipts",
		"unknown-records",
	}
}

func stateRootQuotaRetentionSectionTitle(id string) string {
	switch id {
	case "snapshots":
		return "Snapshots"
	case "diagnostics":
		return "Diagnostics"
	case "execution-receipts":
		return "Execution receipts"
	case "portal-receipts":
		return "Portal receipts"
	case "artifact-receipts":
		return "Artifact receipts"
	case "activation-receipts":
		return "Activation receipts"
	case "unknown-records":
		return "Unknown records"
	default:
		return id
	}
}

func stateRootQuotaRetentionSectionReasons(section StateRootQuotaRetentionSection, options StateRootQuotaRetentionOptions) []string {
	reasons := []string{}
	if section.RecordCount == 0 {
		reasons = append(reasons, "retain-empty-section")
	}
	if section.RetentionExempt {
		reasons = append(reasons, "retain-retention-exempt")
	}
	if section.MalformedRecordCount > 0 {
		reasons = append(reasons, "retain-malformed-record")
	}
	if section.UnknownRecordCount > 0 {
		reasons = append(reasons, "retain-unknown-record-review")
	}
	if options.ActiveSession && section.ID == "execution-receipts" && section.RecordCount > 0 {
		reasons = append(reasons, "retain-active-session")
	}
	if section.CleanupCandidate {
		reasons = append(reasons, "cleanup-candidate-dry-run")
	}
	if options.QuotaBytes > 0 && section.EstimatedBytes > options.QuotaBytes {
		reasons = append(reasons, "cleanup-candidate-over-quota")
	} else if options.QuotaBytes > 0 {
		reasons = append(reasons, "retain-under-quota")
	}
	return quotaRetentionUniqueStrings(reasons)
}

func stateRootQuotaRetentionSectionRecommendation(section StateRootQuotaRetentionSection, options StateRootQuotaRetentionOptions) string {
	switch {
	case section.ID == "unknown-records" && section.RecordCount > 0:
		return "Ask Runtime diagnostics to classify these records before offering cleanup."
	case section.MalformedRecordCount > 0:
		return "Keep these records until a Runtime repair or support review explains the malformed metadata."
	case section.RetentionExempt:
		return "Keep retention-exempt evidence unless the user explicitly changes the retention policy."
	case options.ActiveSession && section.ID == "execution-receipts" && section.RecordCount > 0:
		return "Wait for the active application session to end before offering receipt cleanup."
	case section.CleanupCandidate:
		return "Show this section as a cleanup candidate, but keep cleanup disabled until a later approved action."
	case section.RecordCount == 0:
		return "No records are present in this section."
	default:
		return "Keep this section as current Runtime evidence."
	}
}

func stateRootQuotaRetentionSectionIDs(sections []StateRootQuotaRetentionSection) []string {
	ids := make([]string, 0, len(sections))
	for _, section := range sections {
		ids = append(ids, section.ID)
	}
	return ids
}

func stateRootQuotaRetentionEstimatedBytes(sections []StateRootQuotaRetentionSection) int64 {
	var total int64
	for _, section := range sections {
		total += section.EstimatedBytes
	}
	return total
}

func countStateRootQuotaRetentionSections(sections []StateRootQuotaRetentionSection, overQuota bool) StateRootQuotaRetentionCounts {
	var counts StateRootQuotaRetentionCounts
	for _, section := range sections {
		counts.TotalRecords += section.RecordCount
		if section.CleanupCandidate {
			counts.CleanupCandidates++
		}
		if section.RetentionExempt {
			counts.RetentionExempt++
		}
		counts.MalformedRecords += section.MalformedRecordCount
		counts.UnknownRecords += section.UnknownRecordCount
		if len(section.BlockedCleanupReasons) > 0 {
			counts.BlockedSections++
		}
		if overQuota && section.RecordCount > 0 {
			counts.OverQuotaSections++
		} else if section.RecordCount > 0 {
			counts.UnderQuotaSections++
		}
	}
	return counts
}

func stateRootQuotaRetentionReasonCodes(sections []StateRootQuotaRetentionSection) []string {
	var reasons []string
	for _, section := range sections {
		reasons = append(reasons, section.RetentionReasonCodes...)
	}
	return quotaRetentionUniqueStrings(reasons)
}

func stateRootQuotaRetentionBlockedReasons(sections []StateRootQuotaRetentionSection) []string {
	var reasons []string
	for _, section := range sections {
		reasons = append(reasons, section.BlockedCleanupReasons...)
	}
	return quotaRetentionUniqueStrings(reasons)
}

func stateRootQuotaRetentionRecommendations(sections []StateRootQuotaRetentionSection, overQuota bool) []string {
	recommendations := []string{}
	if overQuota {
		recommendations = append(recommendations, "Storage is over quota; show cleanup candidates as review-only recommendations.")
	} else {
		recommendations = append(recommendations, "Storage is under quota; keep cleanup actions disabled.")
	}
	for _, section := range sections {
		if section.UserSafeRecommendation != "" {
			recommendations = append(recommendations, section.UserSafeRecommendation)
		}
	}
	return quotaRetentionUniqueStrings(recommendations)
}

func stateRootQuotaRetentionSummary(preview StateRootQuotaRetentionPreview) string {
	if !preview.StateRootPresent {
		return "Runtime state records are not present yet, so there is nothing to clean up."
	}
	if preview.OverQuota {
		return "Runtime found state records over the configured quota and produced review-only cleanup recommendations without deleting anything."
	}
	return "Runtime found state records within the configured quota and kept all cleanup actions disabled."
}

func validateStateRootQuotaRetentionPreview(preview StateRootQuotaRetentionPreview) (StateRootQuotaRetentionPreview, error) {
	if preview.FileDeletionEnabled || preview.DirectoriesCreated || preview.LogTruncationEnabled || preview.ReceiptsRewritten || preview.SnapshotDeletionEnabled || preview.StateRootPathExposed || preview.HostRootModified {
		return StateRootQuotaRetentionPreview{}, errors.New("state root quota retention preview must keep all mutating actions disabled")
	}
	if err := validateNoBackendTerms(preview, "state root quota retention preview"); err != nil {
		return StateRootQuotaRetentionPreview{}, err
	}
	return preview, nil
}

func quotaRetentionUniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
