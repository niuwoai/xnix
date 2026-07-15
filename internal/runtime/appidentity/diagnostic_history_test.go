package appidentity

import (
	"testing"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

func TestDiagnosticHistoryPreviewProjectsKDESafeHistory(t *testing.T) {
	latest := diagnostics.RunHistoryRecord{
		RunID:            "smoke-002",
		ApplicationID:    "org.example.ledger",
		TestType:         "smoke",
		Overall:          diagnostics.OutcomeFail,
		RelativePath:     "diagnostics-ledger/runs/smoke-002.json",
		FailingIDs:       []string{"runtime-launch-binding"},
		RepairIssue:      "engine-binding-pending",
		SnapshotRequired: true,
	}
	history := diagnostics.RunHistory{
		SchemaVersion: "xnix.runtime.diagnostic_run_history.v1",
		RecordType:    "diagnostic-run-history",
		Source:        "go-runtime-state-root-diagnostic-run-history",
		ApplicationID: "org.example.ledger",
		Records: []diagnostics.RunHistoryRecord{
			{
				RunID:         "smoke-001",
				ApplicationID: "org.example.ledger",
				TestType:      "smoke",
				Overall:       diagnostics.OutcomePass,
				RelativePath:  "diagnostics-ledger/runs/smoke-001.json",
			},
			latest,
		},
		Counts: diagnostics.RunHistoryCounts{
			Total:  2,
			Passed: 1,
			Failed: 1,
		},
		Latest:                      &latest,
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
	}
	preview, err := NewDiagnosticHistoryPreview(history)
	if err != nil {
		t.Fatalf("NewDiagnosticHistoryPreview: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.diagnostic_history_preview.v1" ||
		preview.RequestType != "diagnostic-history-preview" ||
		preview.Source != "go-runtime-state-root-diagnostic-run-history+kde-read-model" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetDiagnostics" ||
		preview.ReadMethod != "GetDiagnosticHistoryPreview" ||
		preview.ApplicationID != "org.example.ledger" ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.CompatibilityCenterCard ||
		!preview.SafeForAIDiagnostics ||
		preview.StateRootPathExposed ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.AIProviderCallEnabled ||
		preview.RequestObjectCreated ||
		preview.PermissionGranted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.RepairExecutionEnabled ||
		preview.SettingsPersisted ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected diagnostic history preview: %#v", preview)
	}
	if preview.Counts.Total != 2 || preview.Counts.Failed != 1 || len(preview.Records) != 2 {
		t.Fatalf("unexpected preview counts: %#v", preview)
	}
	if preview.Latest == nil || preview.Latest.RunID != "smoke-002" || preview.Latest.RepairIssue != "engine-binding-pending" {
		t.Fatalf("unexpected latest run: %#v", preview.Latest)
	}
	if preview.CompatibilityCenter.HistoryState != "needs-review" ||
		preview.CompatibilityCenter.LatestRunID != "smoke-002" ||
		preview.CompatibilityCenter.LatestOverall != "fail" ||
		preview.CompatibilityCenter.LatestRepairIssue != "engine-binding-pending" ||
		preview.CompatibilityCenter.ActionExecutionEnabled ||
		preview.CompatibilityCenter.RepairExecutionEnabled ||
		preview.CompatibilityCenter.BackendLaunchEnabled {
		t.Fatalf("unexpected Compatibility Center projection: %#v", preview.CompatibilityCenter)
	}
	if err := validateNoBackendTerms(preview, "diagnostic history preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms: %v", err)
	}
}

func TestDiagnosticHistoryPreviewEmptyHistory(t *testing.T) {
	preview, err := NewDiagnosticHistoryPreview(diagnostics.RunHistory{})
	if err != nil {
		t.Fatalf("NewDiagnosticHistoryPreview: %v", err)
	}
	if preview.CompatibilityCenter.HistoryState != "not-run" ||
		preview.CompatibilityCenter.TotalRunCount != 0 ||
		preview.Latest != nil ||
		len(preview.Records) != 0 {
		t.Fatalf("unexpected empty diagnostic history preview: %#v", preview)
	}
}
