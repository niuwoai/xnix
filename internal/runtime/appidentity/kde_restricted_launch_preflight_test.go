package appidentity

import (
	"path/filepath"
	"testing"

	"xnix.local/xnix/internal/runtime/execution"
)

func TestKDERestrictedLaunchPreflightRemainsBlocked(t *testing.T) {
	recipeRecord := Recipe{ID: "org.xnix.sample.notepad", Name: "Sample Notepad", Version: "1.0.0", Icon: "accessories-text-editor", Mode: "automatic", SupportedExtensions: []string{".txt"}}
	provenance := Provenance{Source: "registry", RegistryName: "xnix-offline-samples", DigestVerified: true, SignatureStatus: "development-only"}
	record, err := NewKDERestrictedLaunchPreflightRecord(recipeRecord, provenance, KDERestrictedLaunchAuthorizationOptions{StateRoot: t.TempDir(), Mode: "test-only", Directive: execution.RestrictedTestPreparationDirective})
	if err != nil {
		t.Fatalf("NewKDERestrictedLaunchPreflightRecord returned error: %v", err)
	}
	if record.SchemaVersion != "xnix.runtime.kde_restricted_launch_preflight.v1" || record.StateRootWriteScope != "explicit-test-root-only" || record.Preflight.Status != "blocked" || record.Preflight.BlockerCount != 2 || !containsString(record.Preflight.BlockerIDs, "recipe-trust") || !containsString(record.Preflight.BlockerIDs, "runtime-write-gate") || !record.Preflight.ReceiptPersisted || !record.Preflight.ReceiptReadBack || !record.Preflight.SafeInputsReady || !record.Preflight.PreparationAuthorized || !record.Preflight.ReadyForPacketAssembly ||
		record.Execution.State != "blocked" || record.Session.State != "blocked" || !record.PreflightBoundaryJoined || record.CoreReceiptCount != 10 || !record.AllChecksPassed || record.CheckCount != 8 || record.PassedCheckCount != 8 {
		t.Fatalf("unexpected restricted launch preflight evidence: %+v", record)
	}
	if filepath.IsAbs(record.Preflight.ReceiptRelativePath) || len(record.Preflight.ReceiptSHA256) != 64 {
		t.Fatalf("preflight packet must be relative and digest-backed: %+v", record.Preflight)
	}
	if record.StateRootPathExposed || record.Preflight.ProductImageReady || record.Preflight.LaunchPreflightPassed || record.ProductImageReady || record.ProductionTrustSatisfied || record.RuntimeWriteGateEnabled || record.LaunchPreflightPassed || record.LaunchAuthorized || record.ExecutionApproved || record.ProcessStartAuthorized || record.CommandMaterialized || record.ExecutablePathResolved || record.BackendSelectedForLaunch || record.BackendLaunchEnabled || record.BackendProcessStarted || record.ProductionBusOwnership || record.NetworkRequired || record.HostRootModified || record.PrivilegedContainerRequired || record.RawCommandExposed || record.BackendDetailsExposed {
		t.Fatalf("unsafe restricted launch preflight gates: %+v", record)
	}
}
