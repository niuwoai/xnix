package appidentity

import (
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestKDETestLaunchMaterializationRecordMaterializesPlanOnly(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	record, err := NewKDETestLaunchMaterializationRecord(recipeRecord, provenance, KDERestrictedLaunchAuthorizationOptions{StateRoot: t.TempDir(), Mode: "test-only", Directive: execution.RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("NewKDETestLaunchMaterializationRecord returned error: %v", err)
	}

	if record.SchemaVersion != "xnix.runtime.kde_test_launch_materialization.v1" ||
		record.RecordType != "kde-test-launch-materialization-record" ||
		record.Mode != "test-only" ||
		record.StateRootWriteScope != "explicit-test-root-only" ||
		!record.MaterializationBoundary ||
		!record.PlanMaterialized ||
		!record.TestOnly ||
		record.CoreReceiptCount != 11 ||
		!record.AllChecksPassed ||
		record.CheckCount != 9 ||
		record.PassedCheckCount != 9 {
		t.Fatalf("unexpected test launch materialization evidence: %+v", record)
	}
	if record.Preflight.Status != "blocked" ||
		record.Preflight.BlockerCount != 2 ||
		!record.Preflight.ReadyForPacketAssembly ||
		record.Materialization.Status != "blocked-plan-materialized" ||
		record.Materialization.MaterializationScope != "test-only-review-plan" ||
		record.Materialization.MaterializedArtifactCount != 4 ||
		!containsString(record.Materialization.MaterializedArtifactIDs, "launch-intent-reference") ||
		!containsString(record.Materialization.MaterializedArtifactIDs, "preflight-reference") ||
		record.Materialization.BlockedByCount != 2 ||
		!containsString(record.Materialization.BlockedByIDs, "recipe-trust") ||
		!containsString(record.Materialization.BlockedByIDs, "runtime-write-gate") ||
		!record.Materialization.ReceiptPersisted ||
		!record.Materialization.ReceiptReadBack ||
		!record.Materialization.SafeInputsReady ||
		!record.Materialization.PreparationAuthorized ||
		!record.Materialization.PreflightReadBack ||
		!record.Materialization.PlanMaterialized ||
		!record.Materialization.TestOnly {
		t.Fatalf("unexpected materialization summary: %+v", record.Materialization)
	}
	if filepath.IsAbs(record.Materialization.ReceiptRelativePath) || len(record.Materialization.ReceiptSHA256) != 64 {
		t.Fatalf("materialization plan must be relative and digest-backed: %+v", record.Materialization)
	}
	if record.StateRootPathExposed ||
		record.ProductImageReady ||
		record.ProductionTrustSatisfied ||
		record.RuntimeWriteGateEnabled ||
		record.LaunchPreflightPassed ||
		record.LaunchAuthorized ||
		record.ExecutionApproved ||
		record.ProcessStartAuthorized ||
		record.CommandMaterialized ||
		record.ExecutablePathResolved ||
		record.BackendSelectedForLaunch ||
		record.BackendLaunchEnabled ||
		record.BackendProcessStarted ||
		record.ProductionBusOwnership ||
		record.NetworkRequired ||
		record.HostRootModified ||
		record.PrivilegedContainerRequired ||
		record.RawCommandExposed ||
		record.RawExecutableExposed ||
		record.BackendDetailsExposed ||
		record.Materialization.CommandMaterialized ||
		record.Materialization.ExecutablePathResolved ||
		record.Materialization.BackendSelectedForLaunch ||
		record.Materialization.BackendLaunchEnabled {
		t.Fatalf("unsafe test launch materialization gates: %+v", record)
	}
	if record.Execution.State != "blocked" || record.Session.State != "blocked" {
		t.Fatalf("execution and session must remain blocked: execution=%+v session=%+v", record.Execution, record.Session)
	}
	if err := validateNoBackendTerms(record, "KDE test launch materialization test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestKDETestLaunchMaterializationRecordRequiresExactBoundary(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	for _, options := range []KDERestrictedLaunchAuthorizationOptions{
		{StateRoot: t.TempDir(), Mode: "", Directive: execution.RestrictedTestPreparationDirective},
		{StateRoot: t.TempDir(), Mode: "test-only", Directive: ""},
		{StateRoot: t.TempDir(), Mode: "test-only", Directive: "approve"},
	} {
		if _, err := NewKDETestLaunchMaterializationRecord(recipeRecord, provenance, options); err == nil {
			t.Fatalf("expected exact restricted materialization boundary to be required for options: %+v", options)
		}
	}
}
