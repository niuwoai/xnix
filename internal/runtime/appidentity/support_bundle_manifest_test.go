package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/diagnostics"
)

func TestSupportBundleManifestPreviewSummarizesRedactedEvidence(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)
	root := writeSupportBundleRuntimeRoot(t)
	history := supportBundleTestHistory()

	preview, err := plan.SupportBundleManifestPreview(history, SupportBundleManifestOptions{
		RuntimeRoot: root,
		Issue:       "engine-binding-pending",
		TestType:    "smoke",
	})
	if err != nil {
		t.Fatalf("SupportBundleManifestPreview returned error: %v", err)
	}

	if preview.Version != "0.2.301" ||
		preview.SchemaVersion != "xnix.runtime.support_bundle_manifest.v1" ||
		preview.RequestType != "support-bundle-manifest-preview" ||
		preview.ManifestType != "redacted-offline-support-bundle-manifest" ||
		preview.Source != "diagnostic-history-preview+ai-diagnostic-input-preview+ai-diagnostic-recommendation-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetSupportBundleManifest" ||
		preview.ReadMethod != "GetSupportBundleManifestPreview" ||
		preview.Application.ID != "org.example.ledger" ||
		preview.Application.Name != "Example Ledger" ||
		preview.Application.RuntimeMode != "automatic" {
		t.Fatalf("unexpected support bundle manifest schema: %#v", preview)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.UserVisible ||
		!preview.OfflineOnly ||
		preview.ArchiveCreated ||
		preview.FileContentRead ||
		preview.FilePathsExposed ||
		preview.AIProviderCalled ||
		preview.AIProviderCallEnabled ||
		preview.AutoRepairRequested ||
		preview.AutoRepairExecuted ||
		preview.BackendProcessStarted ||
		preview.HostRootModified ||
		preview.StateRootPathExposed ||
		preview.RawExecutableExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("unexpected support bundle safety flags: %#v", preview)
	}
	if preview.SectionCount != 7 ||
		len(preview.SectionIDs) != 7 ||
		preview.DiagnosticRunCount != 2 ||
		len(preview.DiagnosticRunSummaries) != 2 ||
		len(preview.FailingSignalIDs) != 2 ||
		preview.FailingSignalIDs[0] != "portal-review" ||
		preview.FailingSignalIDs[1] != "runtime-launch-binding" ||
		len(preview.RepairRecommendationCategories) != 3 {
		t.Fatalf("unexpected support bundle evidence content: %#v", preview)
	}
	if preview.OmittedEvidence.FileContents == 0 ||
		preview.OmittedEvidence.HostPaths == 0 ||
		preview.OmittedEvidence.EnvironmentVariables == 0 ||
		preview.OmittedEvidence.TokenShapedValues == 0 ||
		preview.OmittedEvidence.CommandShapedValues == 0 ||
		preview.OmittedEvidence.Usernames == 0 {
		t.Fatalf("support bundle manifest must count omitted evidence: %#v", preview.OmittedEvidence)
	}
	if preview.PrivacyRedaction.RedactionStatus != "redacted-preview-only" ||
		preview.PrivacyRedaction.UserDocumentsIncluded ||
		preview.PrivacyRedaction.HostPathsIncluded ||
		preview.PrivacyRedaction.StateRootPathsIncluded ||
		preview.PrivacyRedaction.EnvironmentVariablesIncluded ||
		preview.PrivacyRedaction.TokenShapedValuesIncluded ||
		preview.PrivacyRedaction.CommandShapedValuesIncluded ||
		preview.PrivacyRedaction.UsernamesIncluded ||
		preview.PrivacyRedaction.SecretsIncluded ||
		preview.PrivacyRedaction.NetworkCallsAllowed ||
		!preview.PrivacyRedaction.RequiresUserApprovalForSensitiveTasks {
		t.Fatalf("unexpected privacy redaction summary: %#v", preview.PrivacyRedaction)
	}
	assertSupportBundleManifestSafe(t, preview)
}

func TestSupportBundleManifestPreviewHandlesEmptyHistory(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)
	root := writeSupportBundleRuntimeRoot(t)
	history := diagnostics.RunHistory{ApplicationID: "org.example.ledger"}

	preview, err := plan.SupportBundleManifestPreview(history, SupportBundleManifestOptions{RuntimeRoot: root})
	if err != nil {
		t.Fatalf("SupportBundleManifestPreview returned error: %v", err)
	}
	if preview.DiagnosticRunCount != 0 ||
		len(preview.DiagnosticRunSummaries) != 0 ||
		len(preview.FailingSignalIDs) != 0 ||
		preview.ArchiveCreated ||
		preview.FileContentRead ||
		preview.AIProviderCalled ||
		preview.HostRootModified {
		t.Fatalf("unexpected empty support bundle manifest: %#v", preview)
	}
	assertSupportBundleManifestSafe(t, preview)
}

func TestSupportBundleManifestPreviewRejectsMismatchedHistory(t *testing.T) {
	plan := testAIDiagnosticsPlan(t)
	root := writeSupportBundleRuntimeRoot(t)
	history := supportBundleTestHistory()
	history.ApplicationID = "org.example.other"

	if _, err := plan.SupportBundleManifestPreview(history, SupportBundleManifestOptions{RuntimeRoot: root}); err == nil {
		t.Fatalf("SupportBundleManifestPreview must reject mismatched history application id")
	}
}

func supportBundleTestHistory() diagnostics.RunHistory {
	records := []diagnostics.RunHistoryRecord{
		{
			RunID:         "smoke-001",
			ApplicationID: "org.example.ledger",
			TestType:      "smoke",
			Overall:       diagnostics.OutcomeFail,
			FailingIDs:    []string{"runtime-launch-binding"},
			RepairIssue:   "engine-binding-pending",
		},
		{
			RunID:            "smoke-002",
			ApplicationID:    "org.example.ledger",
			TestType:         "smoke",
			Overall:          diagnostics.OutcomeBlocked,
			FailingIDs:       []string{"portal-review"},
			RepairIssue:      "portal-approval-required",
			SnapshotRequired: true,
		},
	}
	latest := records[1]
	return diagnostics.RunHistory{
		SchemaVersion: "xnix.runtime.diagnostic_run_history.v1",
		RecordType:    "diagnostic-run-history",
		Source:        "go-runtime-state-root-diagnostic-run-history",
		ApplicationID: "org.example.ledger",
		Records:       records,
		Counts: diagnostics.RunHistoryCounts{
			Total:   2,
			Failed:  1,
			Blocked: 1,
		},
		Latest: &latest,
	}
}

func writeSupportBundleRuntimeRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "VERSION"), []byte("0.2.301\n"), 0o600); err != nil {
		t.Fatalf("WriteFile VERSION returned error: %v", err)
	}
	return root
}

func assertSupportBundleManifestSafe(t *testing.T, preview SupportBundleManifestPreview) {
	t.Helper()
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, required := range []string{"archive_created", "file_content_read", "ai_provider_called", "host_root_modified", "state_root_path_exposed", "raw_command_exposed"} {
		if !strings.Contains(text, required) {
			t.Fatalf("support bundle manifest JSON must include %q: %s", required, string(encoded))
		}
	}
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "wine/", ".wine", "virtual machine", "/home", "/users", "/private", "/var", "/opt", "/tmp", "file://"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("support bundle manifest exposes forbidden term %q: %s", forbidden, string(encoded))
		}
	}
	if err := validateNoBackendTerms(preview, "support bundle manifest preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}
