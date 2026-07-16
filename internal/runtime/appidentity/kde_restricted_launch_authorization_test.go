package appidentity

import (
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestKDERestrictedLaunchAuthorizationRecordsPreparationOnly(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	record, err := NewKDERestrictedLaunchAuthorizationRecord(recipeRecord, provenance, KDERestrictedLaunchAuthorizationOptions{StateRoot: t.TempDir(), Mode: "test-only", Directive: execution.RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("NewKDERestrictedLaunchAuthorizationRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_restricted_launch_authorization.v1" || !record.Prerequisite.BackendStateJoined || !record.Prerequisite.LifecycleReady || !record.Prerequisite.SnapshotVerified || !record.Prerequisite.DiagnosticVerified || !record.Prerequisite.PortalReceiptCompleted || !record.Prerequisite.ExecutionBlocked || !record.Prerequisite.BackendProcessesStopped || !record.Prerequisite.AllChecksPassed ||
		record.Authorization.State != "authorized-preparation-only" || record.Authorization.Scope != execution.RestrictedTestPreparationScope || !record.Authorization.ReceiptPersisted || !record.Authorization.ReceiptReadBack || !record.Authorization.PreparationAuthorized ||
		record.Execution.State != "blocked" || record.Execution.PassedGateCount != 4 || record.Session.State != "blocked" || record.FanOut.SurfaceCount != 4 || !record.AuthorizationBoundaryJoined || record.CoreReceiptCount != 9 || !record.AllChecksPassed || record.CheckCount != 8 || record.PassedCheckCount != 8 {
		t.Fatalf("unexpected restricted launch authorization evidence: %+v", record)
	}
	if filepath.IsAbs(record.Authorization.ReceiptRelativePath) || len(record.Authorization.ReceiptSHA256) != 64 {
		t.Fatalf("authorization receipt must be relative and digest-backed: %+v", record.Authorization)
	}
	if !record.StateRootWritesEnabled || record.StateRootWriteScope != "explicit-test-root-only" || record.StateRootPathExposed || record.Authorization.LaunchAuthorized || record.Authorization.ProcessStartAuthorized || record.Authorization.ExecutionApproved || record.ProductionTrustSatisfied || record.RuntimeWriteGateEnabled || record.ArtifactAcquisitionEnabled || record.BackendInstallEnabled || record.BackendLaunchEnabled || record.BackendProcessStarted || record.RealPortalCallEnabled || record.ExecutionApproved || record.LaunchAuthorized || record.LaunchAllowed || record.LaunchEnabled || record.ExecutionStarted || record.ProcessStartAuthorized || record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified || record.PrivilegedContainerRequired || record.RawCommandExposed || record.BackendDetailsExposed {
		t.Fatalf("unsafe restricted launch authorization gates: %+v", record)
	}
}

func TestKDERestrictedLaunchAuthorizationRequiresExactDirective(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	for _, options := range []KDERestrictedLaunchAuthorizationOptions{
		{StateRoot: t.TempDir(), Mode: "production", Directive: execution.RestrictedTestPreparationDirective},
		{StateRoot: t.TempDir(), Mode: "test-only", Directive: "yes"},
	} {
		if _, err := NewKDERestrictedLaunchAuthorizationRecord(recipeRecord, provenance, options); err == nil {
			t.Fatalf("expected unsafe authorization options to be rejected: %+v", options)
		}
	}
}
