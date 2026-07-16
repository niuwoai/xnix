package appidentity

import (
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/diagnostics"
	"xnix.local/xnix/internal/runtime/safety"
)

func crashHangHistory(records ...diagnostics.RunHistoryRecord) diagnostics.RunHistory {
	return diagnostics.RunHistory{
		SchemaVersion: "xnix.runtime.diagnostic_run_history.v1",
		RecordType:    "diagnostic-run-history",
		Records:       records,
	}
}

func crashHangGroup(t *testing.T, preview CrashHangSignalSummaryPreview, id string) CrashHangSignalGroup {
	t.Helper()
	for _, group := range preview.SignalGroups {
		if group.ID == id {
			return group
		}
	}
	t.Fatalf("signal group %q not present; ids=%v", id, preview.SignalGroupIDs)
	return CrashHangSignalGroup{}
}

func assertCrashHangSafe(t *testing.T, preview CrashHangSignalSummaryPreview) {
	t.Helper()
	if preview.PrivateLogRead || preview.FileContentRead || preview.FilePathsExposed ||
		preview.AIProviderCalled || preview.AIProviderCallEnabled || preview.RepairExecuted ||
		preview.RepairApplied || preview.BackendProcessStarted || preview.NetworkRequired ||
		preview.HostRootModified || preview.StateRootPathExposed || preview.BackendDetailsExposed ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("preview must keep every unsafe capability disabled: %+v", preview)
	}
	if !preview.ReviewOnly || !preview.RuntimeOwned || !preview.GoRuntimeBacked {
		t.Fatalf("preview must be a runtime-owned review-only read model: %+v", preview)
	}
	if err := safety.ValidatePayload("crash hang signal summary preview", preview); err != nil {
		t.Fatalf("preview leaks forbidden content: %v", err)
	}
}

func TestCrashHangSignalSummaryNoHistoryIsSafe(t *testing.T) {
	preview, err := NewCrashHangSignalSummaryPreview(crashHangHistory(), CrashHangSignalSummaryOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "no-history" {
		t.Fatalf("empty history should be no-history, got %q", preview.OverallState)
	}
	if len(preview.SignalGroups) != 8 {
		t.Fatalf("expected 8 stable signal groups, got %d", len(preview.SignalGroups))
	}
	for _, group := range preview.SignalGroups {
		if group.Count != 0 || group.RecurrenceHint != "not observed" {
			t.Fatalf("empty history group %q should be unobserved: %+v", group.ID, group)
		}
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryGroupsAndRecurrence(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", TestType: "smoke", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"app-crash-on-start"}, RepairIssue: "engine-binding-pending"},
		diagnostics.RunHistoryRecord{RunID: "run-002", ApplicationID: "org.example.editor", TestType: "smoke", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"app-crash-on-start"}, RepairIssue: "engine-binding-pending"},
		diagnostics.RunHistoryRecord{RunID: "run-003", ApplicationID: "org.example.editor", TestType: "smoke", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"window-hang-detected"}, RepairIssue: "engine-binding-pending"},
		diagnostics.RunHistoryRecord{RunID: "run-004", ApplicationID: "org.example.editor", TestType: "preflight", Overall: diagnostics.OutcomeBlocked, FailingIDs: []string{"portal-approval-required"}, RepairIssue: "portal-approval-required"},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{ApplicationID: "org.example.editor"})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "needs-review" {
		t.Fatalf("failing history should need review, got %q", preview.OverallState)
	}
	crash := crashHangGroup(t, preview, "crash")
	if crash.Count != 2 || crash.RecurrenceHint != "repeated across diagnostic metadata" {
		t.Fatalf("crash group should be recurring with two runs: %+v", crash)
	}
	if len(crash.RelatedRunIDs) != 2 || crash.RelatedRunIDs[0] != "run-001" {
		t.Fatalf("crash group should relate run-001 and run-002: %+v", crash.RelatedRunIDs)
	}
	if hang := crashHangGroup(t, preview, "hang"); hang.Count != 1 || hang.RecurrenceHint != "observed once" {
		t.Fatalf("hang group should be observed once: %+v", hang)
	}
	if perm := crashHangGroup(t, preview, "permission-denial"); perm.Count != 1 {
		t.Fatalf("permission group should classify a blocked portal signal: %+v", perm)
	}
	if preview.LatestRunID != "run-004" || preview.LatestKnownState != "blocked" {
		t.Fatalf("latest known state should track run-004 blocked: %q/%q", preview.LatestRunID, preview.LatestKnownState)
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryMalformedHistoryIsBlocked(t *testing.T) {
	preview, err := NewCrashHangSignalSummaryPreview(crashHangHistory(), CrashHangSignalSummaryOptions{
		MalformedHistory:     true,
		MalformedEvidenceIDs: []string{"run-broken"},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.OverallState != "blocked-malformed-history" {
		t.Fatalf("malformed history should block, got %q", preview.OverallState)
	}
	if preview.Counts.MalformedRecords != 1 || len(preview.MalformedEvidenceIDs) != 1 {
		t.Fatalf("malformed record should be surfaced: %+v", preview.Counts)
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryMixedApplicationIDsAreFiltered(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"app-crash-on-start"}},
		diagnostics.RunHistoryRecord{RunID: "run-002", ApplicationID: "org.example.viewer", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"window-hang-detected"}},
		diagnostics.RunHistoryRecord{RunID: "run-003", ApplicationID: "org.example.viewer", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"network-connection-lost"}},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{ApplicationID: "org.example.viewer"})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.Counts.TotalRuns != 2 {
		t.Fatalf("only the two viewer runs should be counted, got %d", preview.Counts.TotalRuns)
	}
	if preview.MixedApplicationRecordCount != 1 {
		t.Fatalf("the editor run should be recorded as a mixed-application skip, got %d", preview.MixedApplicationRecordCount)
	}
	if crash := crashHangGroup(t, preview, "crash"); crash.Count != 0 {
		t.Fatalf("the editor crash must not leak into the viewer summary: %+v", crash)
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryDuplicateSignalIDsAreCounted(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"app-crash-on-start", "app-crash-on-start"}},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.DuplicateSignalIDCount != 1 || preview.Counts.DuplicateSignalIDs != 1 {
		t.Fatalf("duplicate signal id should be counted once, got %d", preview.DuplicateSignalIDCount)
	}
	crash := crashHangGroup(t, preview, "crash")
	if len(crash.RelatedSignalIDs) != 1 {
		t.Fatalf("duplicate signal id should be de-duplicated in output: %+v", crash.RelatedSignalIDs)
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryPrivacySensitiveFieldsAreOmitted(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"/home/user/crash-token", "app-crash-on-start"}},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{
		MalformedEvidenceIDs: []string{"/root/secret-run", "run-broken"},
	})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.Counts.PrivacySensitiveOmissions == 0 {
		t.Fatalf("privacy-sensitive signal id should be omitted and counted")
	}
	if len(preview.MalformedEvidenceIDs) != 1 || preview.MalformedEvidenceIDs[0] != "run-broken" {
		t.Fatalf("sensitive malformed evidence id should be dropped: %+v", preview.MalformedEvidenceIDs)
	}
	crash := crashHangGroup(t, preview, "crash")
	for _, id := range crash.RelatedSignalIDs {
		if strings.Contains(id, "/") || strings.Contains(id, "token") {
			t.Fatalf("privacy-sensitive signal id leaked into output: %q", id)
		}
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryBlockedRepairEvidenceIsCounted(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"app-crash-on-start"}, RepairIssue: "engine-binding-pending"},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if preview.BlockedRepairEvidenceCount != 1 {
		t.Fatalf("blocked repair evidence should be counted, got %d", preview.BlockedRepairEvidenceCount)
	}
	if crash := crashHangGroup(t, preview, "crash"); !crash.RepairBlocked {
		t.Fatalf("crash group should mark repair as blocked: %+v", crash)
	}
	if preview.RepairExecuted || preview.RepairApplied {
		t.Fatalf("blocked repair evidence must never imply execution")
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryRegressionAfterRepair(t *testing.T) {
	history := crashHangHistory(
		diagnostics.RunHistoryRecord{RunID: "run-001", ApplicationID: "org.example.editor", Overall: diagnostics.OutcomeFail, FailingIDs: []string{"post-repair-regression"}},
	)
	preview, err := NewCrashHangSignalSummaryPreview(history, CrashHangSignalSummaryOptions{})
	if err != nil {
		t.Fatalf("build preview: %v", err)
	}
	if regression := crashHangGroup(t, preview, "regression-after-repair"); regression.Count != 1 {
		t.Fatalf("regression signal should be grouped: %+v", regression)
	}
	assertCrashHangSafe(t, preview)
}

func TestCrashHangSignalSummaryRejectsUnsafeInputs(t *testing.T) {
	if _, err := NewCrashHangSignalSummaryPreview(crashHangHistory(), CrashHangSignalSummaryOptions{ApplicationID: "Not A Valid ID"}); err == nil {
		t.Fatalf("expected invalid application id to be rejected")
	}
}
