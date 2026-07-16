package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/safety"
)

func restoreCandidate(t *testing.T, preview SnapshotRestoreCandidatesPreview, id string) SnapshotRestoreCandidate {
	t.Helper()
	for _, candidate := range preview.Candidates {
		if candidate.ID == id {
			return candidate
		}
	}
	t.Fatalf("restore candidate %q not present; ids=%v", id, preview.CandidateIDs)
	return SnapshotRestoreCandidate{}
}

func assertSnapshotRestoreSafe(t *testing.T, preview SnapshotRestoreCandidatesPreview) {
	t.Helper()
	if preview.RestoreExecuted || preview.SnapshotDeletionEnabled || preview.FileContentRead ||
		preview.SessionTerminated || preview.BackendLaunchEnabled || preview.StateRootPathExposed ||
		preview.HostRootModified || preview.BackendDetailsExposed {
		t.Fatalf("snapshot restore preview must keep every unsafe capability disabled: %+v", preview)
	}
	if err := safety.ValidatePayload("snapshot restore candidates preview", preview); err != nil {
		t.Fatalf("snapshot restore preview leaks forbidden content: %v", err)
	}
}

func TestSnapshotRestoreCandidatesRankAndClassify(t *testing.T) {
	preview, err := NewSnapshotRestoreCandidatesPreview(SnapshotRestoreCandidatesOptions{
		ApplicationID: "org.example.ledger",
		BaselineID:    "snap-004",
		Candidates: []SnapshotCandidateInput{
			{ID: "snap-001", Reason: "manual", FileCount: 2, ContentHash: "aa", Verified: true},
			{ID: "snap-002", Reason: "before-repair", FileCount: 2, ContentHash: "bb", Verified: true},
			{ID: "snap-003", Reason: "before-engine-change", FileCount: 2, ContentHash: "cc", Verified: true},
			{ID: "snap-004", Reason: "manual", FileCount: 2, ContentHash: "dd", Verified: true},
			{ID: "snap-005", Reason: "manual", FileCount: 0, ContentHash: "ee", Verified: false},
		},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "candidates-available" || preview.RecommendedCandidateID != "snap-004" {
		t.Fatalf("baseline should be the recommended candidate: %+v", preview.RecommendedCandidateID)
	}
	if latest := restoreCandidate(t, preview, "snap-004"); latest.Category != "latest-good" || latest.Rank != 1 || latest.CompatibilityRisk != "low" {
		t.Fatalf("baseline should rank first as latest-good: %+v", latest)
	}
	if tested := restoreCandidate(t, preview, "snap-003"); tested.Category != "latest-tested" {
		t.Fatalf("before-engine-change should be latest-tested: %+v", tested)
	}
	if running := restoreCandidate(t, preview, "snap-002"); running.Category != "last-known-running" || running.CompatibilityRisk != "medium" {
		t.Fatalf("before-repair should be last-known-running: %+v", running)
	}
	corrupt := restoreCandidate(t, preview, "snap-005")
	if corrupt.Category != "blocked" || corrupt.Selectable || corrupt.DigestStatus != "corrupt" {
		t.Fatalf("unverified snapshot should be blocked: %+v", corrupt)
	}
	if corrupt.Rank != preview.CandidateCount {
		t.Fatalf("blocked candidate should rank last, got %d of %d", corrupt.Rank, preview.CandidateCount)
	}
	if preview.Counts.DigestCorrupt != 1 || preview.Counts.Selectable != 4 {
		t.Fatalf("unexpected counts: %+v", preview.Counts)
	}
	assertSnapshotRestoreSafe(t, preview)
}

func TestSnapshotRestoreCandidatesActiveSessionBlocksAll(t *testing.T) {
	preview, err := NewSnapshotRestoreCandidatesPreview(SnapshotRestoreCandidatesOptions{
		ApplicationID: "org.example.ledger",
		BaselineID:    "snap-001",
		ActiveSession: true,
		Candidates: []SnapshotCandidateInput{
			{ID: "snap-001", Reason: "manual", FileCount: 1, ContentHash: "aa", Verified: true},
		},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "blocked-active-session" || preview.RecommendedCandidateID != "" {
		t.Fatalf("active session should block selection: %+v", preview.OverallState)
	}
	candidate := restoreCandidate(t, preview, "snap-001")
	if candidate.Selectable || !containsRestoreReason(candidate.BlockedReasons, "active-session") {
		t.Fatalf("candidate should be blocked by active session: %+v", candidate)
	}
	assertSnapshotRestoreSafe(t, preview)
}

func TestSnapshotRestoreCandidatesNoCandidates(t *testing.T) {
	preview, err := NewSnapshotRestoreCandidatesPreview(SnapshotRestoreCandidatesOptions{ApplicationID: "org.example.ledger"})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "no-candidates" || preview.CandidateCount != 0 {
		t.Fatalf("empty store should report no candidates: %+v", preview.OverallState)
	}
	assertSnapshotRestoreSafe(t, preview)
}

func TestSnapshotRestoreCandidatesRejectsUnsafeInputs(t *testing.T) {
	if _, err := NewSnapshotRestoreCandidatesPreview(SnapshotRestoreCandidatesOptions{ApplicationID: "Not Valid"}); err == nil {
		t.Fatalf("expected invalid application id to be rejected")
	}
}

func containsRestoreReason(reasons []string, want string) bool {
	for _, reason := range reasons {
		if reason == want {
			return true
		}
	}
	return false
}
