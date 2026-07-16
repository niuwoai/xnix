package appidentity

import (
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

func TestSupportCaseTimelinePreviewNoHistory(t *testing.T) {
	plan := newSupportCaseTimelinePlan(t)
	preview, err := plan.SupportCaseTimelinePreview(diagnostics.RunHistory{ApplicationID: plan.ApplicationID}, SupportCaseTimelineOptions{RuntimeRoot: projectRootForRuntimeServiceBindingTest(t)})
	if err != nil {
		t.Fatalf("support case timeline preview: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.support_case_timeline.v1" ||
		preview.RequestType != "support-case-timeline-preview" ||
		preview.RuntimeMethod != "GetSupportCaseTimeline" ||
		preview.ReadMethod != "GetSupportCaseTimelinePreview" {
		t.Fatalf("unexpected timeline schema: %#v", preview)
	}
	if preview.Counts.DiagnosticRunEvents != 1 || preview.EventCount == 0 {
		t.Fatalf("no-history timeline should include a safe placeholder event: %#v", preview.Counts)
	}
	assertSupportCaseTimelineNoSideEffects(t, preview)
	assertSupportCaseTimelineSafe(t, preview)
}

func TestSupportCaseTimelinePreviewJoinsDiagnosticActionsRepairsOnboardingAndKDE(t *testing.T) {
	plan := newSupportCaseTimelinePlan(t)
	history := diagnostics.RunHistory{
		ApplicationID: plan.ApplicationID,
		Records: []diagnostics.RunHistoryRecord{
			{
				RunID:         "run-001",
				ApplicationID: plan.ApplicationID,
				TestType:      "smoke",
				Overall:       diagnostics.OutcomeFail,
				RelativePath:  "diagnostics-ledger/runs/run-001.json",
				FailingIDs:    []string{"crash-on-start", "portal-denied"},
				RepairIssue:   "portal-approval-required",
			},
			{
				RunID:         "run-002",
				ApplicationID: plan.ApplicationID,
				TestType:      "repair-readiness",
				Overall:       diagnostics.OutcomeBlocked,
				RelativePath:  "diagnostics-ledger/runs/run-002.json",
				FailingIDs:    []string{"timeout-startup"},
				RepairIssue:   "runtime-launch-binding",
			},
		},
	}
	preview, err := plan.SupportCaseTimelinePreview(history, SupportCaseTimelineOptions{RuntimeRoot: projectRootForRuntimeServiceBindingTest(t), Issue: "portal-approval-required", TestType: "smoke"})
	if err != nil {
		t.Fatalf("support case timeline preview: %v", err)
	}
	if preview.Counts.DiagnosticRunEvents != 2 ||
		preview.Counts.BlockedDiagnosticEvents != 1 ||
		preview.Counts.BlockedActionEvents == 0 ||
		preview.Counts.RepairRecommendationEvents == 0 ||
		preview.Counts.OnboardingGapEvents == 0 ||
		preview.Counts.KDEEntrypointEvents == 0 {
		t.Fatalf("timeline did not join expected evidence groups: %#v", preview.Counts)
	}
	for _, group := range []string{"diagnostic-runs", "blocked-actions", "repair-recommendations", "onboarding-gaps", "kde-entrypoint-state"} {
		if !containsString(preview.EventGroups, group) {
			t.Fatalf("timeline missing group %q in %#v", group, preview.EventGroups)
		}
	}
	assertSupportCaseTimelineNoSideEffects(t, preview)
	assertSupportCaseTimelineSafe(t, preview)
}

func TestSupportCaseTimelinePreviewMalformedMixedAndRedactedEvidence(t *testing.T) {
	plan := newSupportCaseTimelinePlan(t)
	history := diagnostics.RunHistory{
		Records: []diagnostics.RunHistoryRecord{
			{
				RunID:         "run-001",
				ApplicationID: plan.ApplicationID,
				Overall:       diagnostics.OutcomeFail,
				RelativePath:  "diagnostics-ledger/runs/run-001.json",
				FailingIDs:    []string{"safe-signal", "/Users/example/private.txt", "token=secret-value", "program.exe"},
			},
			{
				RunID:         "run-foreign",
				ApplicationID: "org.example.other",
				Overall:       diagnostics.OutcomeFail,
				RelativePath:  "diagnostics-ledger/runs/run-foreign.json",
			},
		},
	}
	preview, err := plan.SupportCaseTimelinePreview(history, SupportCaseTimelineOptions{
		RuntimeRoot:          projectRootForRuntimeServiceBindingTest(t),
		MalformedHistory:     true,
		MalformedEvidenceIDs: []string{"bad-receipt", "/Users/example/state.json", "token=abc"},
	})
	if err != nil {
		t.Fatalf("support case timeline preview: %v", err)
	}
	if !preview.MalformedHistory || preview.Counts.MalformedHistoryEvents != 1 {
		t.Fatalf("timeline must expose malformed history as blocked evidence: %#v", preview)
	}
	if preview.MixedApplicationRecordCount != 1 {
		t.Fatalf("timeline must count mixed application records: %#v", preview)
	}
	if preview.RedactionSummary.OmittedSensitiveEvidenceCount == 0 ||
		preview.Counts.RedactedEvidenceOmissions == 0 {
		t.Fatalf("timeline must count redacted evidence: %#v", preview.RedactionSummary)
	}
	encoded, _ := json.Marshal(preview)
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"/users/example", "token=", "secret-value", ".exe"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("timeline leaked forbidden evidence %q: %s", forbidden, text)
		}
	}
	assertSupportCaseTimelineNoSideEffects(t, preview)
	assertSupportCaseTimelineSafe(t, preview)
}

func TestSupportCaseTimelineRejectsMismatchedHistoryApplication(t *testing.T) {
	plan := newSupportCaseTimelinePlan(t)
	_, err := plan.SupportCaseTimelinePreview(diagnostics.RunHistory{ApplicationID: "org.example.other"}, SupportCaseTimelineOptions{RuntimeRoot: projectRootForRuntimeServiceBindingTest(t)})
	if err == nil {
		t.Fatalf("support case timeline accepted mismatched history application id")
	}
}

func newSupportCaseTimelinePlan(t *testing.T) Plan {
	t.Helper()
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".abc", ".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan: %v", err)
	}
	return plan
}

func assertSupportCaseTimelineNoSideEffects(t *testing.T, preview SupportCaseTimelinePreview) {
	t.Helper()
	if preview.TicketCreated ||
		preview.BundleExported ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.RepairApplied ||
		preview.RepairExecuted ||
		preview.ActionExecuted ||
		preview.BackendProcessStarted ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("timeline enabled unsafe side effects: %#v", preview)
	}
	for _, event := range preview.Events {
		if event.TicketCreated ||
			event.BundleExported ||
			event.AIProviderCalled ||
			event.RepairExecuted ||
			event.ActionExecuted ||
			event.BackendProcessStarted ||
			event.HostRootModified ||
			event.FileContentRead {
			t.Fatalf("timeline event enabled unsafe side effects: %#v", event)
		}
	}
}

func assertSupportCaseTimelineSafe(t *testing.T, preview SupportCaseTimelinePreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("marshal timeline: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "/home", "/users", "/private", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("support case timeline exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
