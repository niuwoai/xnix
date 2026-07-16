package appidentity

import (
	"errors"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

// SnapshotCandidateInput is one already-resolved snapshot record the command
// layer reads from the read-only snapshot store. It carries only safe derived
// fields: a stable id, the restore reason, a file count, a content digest, and
// whether the snapshot verified intact.
type SnapshotCandidateInput struct {
	ID          string
	Reason      string
	FileCount   int
	ContentHash string
	Verified    bool
}

// SnapshotRestoreCandidatesOptions carries the resolved snapshot evidence plus
// the baseline (latest verified) id and an active-session flag.
type SnapshotRestoreCandidatesOptions struct {
	ApplicationID string
	Candidates    []SnapshotCandidateInput
	BaselineID    string
	ActiveSession bool
}

type SnapshotRestoreCandidatesPreview struct {
	SchemaVersion           string                     `json:"schema_version"`
	RequestType             string                     `json:"request_type"`
	PreviewType             string                     `json:"preview_type"`
	Source                  string                     `json:"source"`
	Desktop                 string                     `json:"desktop"`
	RuntimeMethod           string                     `json:"runtime_method"`
	ReadMethod              string                     `json:"read_method"`
	ApplicationID           string                     `json:"application_id"`
	OverallState            string                     `json:"overall_state"`
	RecommendedCandidateID  string                     `json:"recommended_candidate_id,omitempty"`
	ActiveSession           bool                       `json:"active_session"`
	Candidates              []SnapshotRestoreCandidate `json:"candidates"`
	CandidateIDs            []string                   `json:"candidate_ids"`
	CandidateCount          int                        `json:"candidate_count"`
	Counts                  SnapshotRestoreCounts      `json:"counts"`
	BlockedReasons          []string                   `json:"blocked_reasons"`
	NextSafeReadOnlyChecks  []string                   `json:"next_safe_read_only_checks"`
	RuntimeOwned            bool                       `json:"runtime_owned"`
	GoRuntimeBacked         bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner          bool                       `json:"kde_policy_owner"`
	UserVisible             bool                       `json:"user_visible"`
	ReviewOnly              bool                       `json:"review_only"`
	RestoreExecuted         bool                       `json:"restore_executed"`
	SnapshotDeletionEnabled bool                       `json:"snapshot_deletion_enabled"`
	FileContentRead         bool                       `json:"file_content_read"`
	SessionTerminated       bool                       `json:"session_terminated"`
	BackendLaunchEnabled    bool                       `json:"backend_launch_enabled"`
	StateRootPathExposed    bool                       `json:"state_root_path_exposed"`
	HostRootModified        bool                       `json:"host_root_modified"`
	BackendDetailsExposed   bool                       `json:"backend_details_exposed"`
	BlockedActions          []string                   `json:"blocked_actions"`
	DesktopSafeSummary      string                     `json:"desktop_safe_summary"`
}

type SnapshotRestoreCandidate struct {
	ID                string   `json:"id"`
	Rank              int      `json:"rank"`
	Category          string   `json:"category"`
	DigestStatus      string   `json:"digest_status"`
	CompatibilityRisk string   `json:"compatibility_risk"`
	Selectable        bool     `json:"selectable"`
	FileCount         int      `json:"file_count"`
	ContentDigest     string   `json:"content_digest"`
	EvidenceIDs       []string `json:"evidence_ids"`
	ReasonCodes       []string `json:"reason_codes"`
	BlockedReasons    []string `json:"blocked_reasons"`
	UserSafeSummary   string   `json:"user_safe_summary"`
	RestoreExecuted   bool     `json:"restore_executed"`
	FileContentRead   bool     `json:"file_content_read"`
	HostRootModified  bool     `json:"host_root_modified"`
}

type SnapshotRestoreCounts struct {
	TotalCandidates int `json:"total_candidates"`
	Selectable      int `json:"selectable"`
	Blocked         int `json:"blocked"`
	DigestVerified  int `json:"digest_verified"`
	DigestCorrupt   int `json:"digest_corrupt"`
	ActiveSession   int `json:"active_session"`
}

// snapshotRestoreCategoryRank orders categories from best (lowest) to worst.
func snapshotRestoreCategoryRank(category string) int {
	switch category {
	case "latest-good":
		return 0
	case "latest-tested":
		return 1
	case "last-known-running":
		return 2
	case "manual":
		return 3
	case "unknown":
		return 4
	default: // blocked
		return 5
	}
}

// NewSnapshotRestoreCandidatesPreview ranks snapshot restore candidates and
// explains why each can or cannot be selected. It never restores data, deletes
// snapshots, reads private file contents, terminates sessions, starts a backend,
// or mutates the host.
func NewSnapshotRestoreCandidatesPreview(options SnapshotRestoreCandidatesOptions) (SnapshotRestoreCandidatesPreview, error) {
	options.ApplicationID = strings.TrimSpace(options.ApplicationID)
	if !idPattern.MatchString(options.ApplicationID) {
		return SnapshotRestoreCandidatesPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}

	candidates := make([]SnapshotRestoreCandidate, 0, len(options.Candidates))
	var counts SnapshotRestoreCounts
	for _, input := range options.Candidates {
		candidate := snapshotRestoreCandidate(input, options.BaselineID, options.ActiveSession)
		candidates = append(candidates, candidate)
		counts.TotalCandidates++
		if candidate.Selectable {
			counts.Selectable++
		} else {
			counts.Blocked++
		}
		if input.Verified {
			counts.DigestVerified++
		} else {
			counts.DigestCorrupt++
		}
	}
	if options.ActiveSession {
		counts.ActiveSession = counts.TotalCandidates
	}

	// Rank: best category first, then most recent id (snapshot ids sort
	// chronologically) so ties break toward the newest candidate.
	sort.SliceStable(candidates, func(i, j int) bool {
		ri, rj := snapshotRestoreCategoryRank(candidates[i].Category), snapshotRestoreCategoryRank(candidates[j].Category)
		if ri != rj {
			return ri < rj
		}
		return candidates[i].ID > candidates[j].ID
	})
	recommended := ""
	for index := range candidates {
		candidates[index].Rank = index + 1
		if recommended == "" && candidates[index].Selectable {
			recommended = candidates[index].ID
		}
	}

	overallState := snapshotRestoreOverallState(counts, options.ActiveSession)
	preview := SnapshotRestoreCandidatesPreview{
		SchemaVersion:          "xnix.runtime.snapshot_restore_candidates.v1",
		RequestType:            "snapshot-restore-candidates-preview",
		PreviewType:            "snapshot-restore-candidate-ranking",
		Source:                 "snapshot-store+baseline+execution-session-evidence",
		Desktop:                "KDE Plasma",
		RuntimeMethod:          "GetSnapshotRestoreCandidates",
		ReadMethod:             "GetSnapshotRestoreCandidatesPreview",
		ApplicationID:          options.ApplicationID,
		OverallState:           overallState,
		RecommendedCandidateID: recommended,
		ActiveSession:          options.ActiveSession,
		Candidates:             candidates,
		CandidateIDs:           snapshotRestoreCandidateIDs(candidates),
		CandidateCount:         len(candidates),
		Counts:                 counts,
		BlockedReasons:         snapshotRestoreAggregateReasons(candidates),
		NextSafeReadOnlyChecks: []string{
			"review the snapshot restore-point baseline",
			"review the application readiness snapshot node",
			"review the active session evidence before restoring",
		},
		RuntimeOwned:            true,
		GoRuntimeBacked:         true,
		KDEPolicyOwner:          false,
		UserVisible:             true,
		ReviewOnly:              true,
		RestoreExecuted:         false,
		SnapshotDeletionEnabled: false,
		FileContentRead:         false,
		SessionTerminated:       false,
		BackendLaunchEnabled:    false,
		StateRootPathExposed:    false,
		HostRootModified:        false,
		BackendDetailsExposed:   false,
		BlockedActions: []string{
			"restore a snapshot from the candidate preview",
			"delete a snapshot from the candidate preview",
			"read snapshot file contents from the candidate preview",
			"terminate an application session from the candidate preview",
			"start a compatibility backend from the candidate preview",
			"mutate host root during the candidate preview",
		},
		DesktopSafeSummary: snapshotRestoreSummary(overallState),
	}
	if err := validateNoBackendTerms(preview, "snapshot restore candidates preview"); err != nil {
		return SnapshotRestoreCandidatesPreview{}, err
	}
	if err := safety.ValidatePayload("snapshot restore candidates preview", preview); err != nil {
		return SnapshotRestoreCandidatesPreview{}, err
	}
	return preview, nil
}

func snapshotRestoreCandidate(input SnapshotCandidateInput, baselineID string, activeSession bool) SnapshotRestoreCandidate {
	candidate := SnapshotRestoreCandidate{
		ID:            input.ID,
		FileCount:     input.FileCount,
		ContentDigest: input.ContentHash,
		EvidenceIDs:   []string{"snapshot-manifest:" + input.ID},
	}
	var reasons []string
	var blocked []string

	switch {
	case !input.Verified:
		candidate.Category = "blocked"
		candidate.DigestStatus = "corrupt"
		blocked = append(blocked, "digest-mismatch")
	case input.ID == baselineID:
		candidate.Category = "latest-good"
		candidate.DigestStatus = "verified"
		reasons = append(reasons, "latest-verified-baseline")
	default:
		candidate.DigestStatus = "verified"
		candidate.Category = snapshotRestoreCategoryForReason(input.Reason)
		reasons = append(reasons, "restore-point-verified")
	}

	if activeSession && candidate.Category != "blocked" {
		blocked = append(blocked, "active-session")
	}
	candidate.Selectable = len(blocked) == 0
	candidate.CompatibilityRisk = snapshotRestoreRisk(candidate.Category, candidate.Selectable)
	candidate.ReasonCodes = snapshotRestoreUnique(reasons)
	candidate.BlockedReasons = snapshotRestoreUnique(blocked)
	candidate.UserSafeSummary = snapshotRestoreCandidateSummary(candidate)
	return candidate
}

func snapshotRestoreCategoryForReason(reason string) string {
	switch strings.TrimSpace(reason) {
	case "before-repair":
		return "last-known-running"
	case "before-engine-change":
		return "latest-tested"
	case "manual":
		return "manual"
	default:
		return "unknown"
	}
}

func snapshotRestoreRisk(category string, selectable bool) string {
	if !selectable {
		return "high"
	}
	switch category {
	case "latest-good", "latest-tested":
		return "low"
	case "last-known-running", "manual":
		return "medium"
	default:
		return "high"
	}
}

func snapshotRestoreCandidateSummary(candidate SnapshotRestoreCandidate) string {
	switch {
	case len(candidate.BlockedReasons) > 0:
		if candidate.DigestStatus == "corrupt" {
			return "This restore point cannot be selected because its recorded digest does not verify."
		}
		return "This restore point cannot be selected yet because an application session is active."
	case candidate.Category == "latest-good":
		return "This is the most recent verified restore point and is the recommended choice."
	case candidate.Category == "latest-tested":
		return "This restore point was taken before an engine change and verifies intact."
	case candidate.Category == "last-known-running":
		return "This restore point captures the last state before a repair and verifies intact."
	case candidate.Category == "manual":
		return "This is a manual restore point that verifies intact."
	default:
		return "This restore point verifies intact but its purpose is unknown; review before selecting."
	}
}

func snapshotRestoreOverallState(counts SnapshotRestoreCounts, activeSession bool) string {
	switch {
	case counts.TotalCandidates == 0:
		return "no-candidates"
	case activeSession:
		return "blocked-active-session"
	case counts.Selectable == 0:
		return "blocked"
	default:
		return "candidates-available"
	}
}

func snapshotRestoreSummary(overallState string) string {
	switch overallState {
	case "no-candidates":
		return "No restore points are recorded yet, so there is nothing to restore. Nothing was restored or deleted."
	case "blocked-active-session":
		return "Restore points exist but an active session blocks selection. Nothing was restored or deleted."
	case "blocked":
		return "Restore points exist but none can be selected because their digests do not verify. Nothing was restored or deleted."
	default:
		return "Ranked restore points are available for review. Nothing was restored, deleted, or read."
	}
}

func snapshotRestoreAggregateReasons(candidates []SnapshotRestoreCandidate) []string {
	var reasons []string
	for _, candidate := range candidates {
		reasons = append(reasons, candidate.BlockedReasons...)
	}
	return snapshotRestoreUnique(reasons)
}

func snapshotRestoreCandidateIDs(candidates []SnapshotRestoreCandidate) []string {
	ids := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		ids = append(ids, candidate.ID)
	}
	return ids
}

func snapshotRestoreUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
