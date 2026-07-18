package owner

import "testing"

func TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewCoversOwnerRoute(t *testing.T) {
	preview, err := NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview(projectRoot(t))
	if err != nil {
		t.Fatalf("NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreview returned error: %v", err)
	}

	if preview.Version != currentProjectVersion(t) ||
		preview.SchemaVersion != "xnix.runtime.backend_adapter_redacted_profile_owner_smoke_coverage.v1" ||
		preview.RequestType != "backend-adapter-redacted-profile-owner-smoke-coverage-preview" ||
		preview.CoverageType != "redacted-adapter-profile-owner-smoke-coverage" ||
		preview.RuntimeMethod != "GetBackendAdapterProfileAudit" ||
		preview.ReadMethod != "GetBackendAdapterRedactedProfileOwnerSmokeCoveragePreview" ||
		preview.OwnerRouteRequestType != "backend-adapter-redacted-profile-audit-preview" ||
		preview.OwnerRouteCommand != "backend-adapter-redacted-profile-audit-preview" {
		t.Fatalf("unexpected redacted adapter profile coverage schema: %#v", preview)
	}
	if preview.SmokeBatchRecordCount != 75 ||
		preview.SmokeReadDispatchRecordCount != 71 ||
		preview.SmokeWriteDenialRecordCount != 4 ||
		!preview.CoverageRecordFound ||
		preview.CoverageRecordSequence == 0 ||
		!preview.ServiceCallReadDispatch ||
		!preview.ServiceCallDispatchReady ||
		!preview.DispatchReadOnly ||
		!preview.DispatchRouteReady ||
		preview.DispatchRouteSource != "go-owner-local-preview" ||
		preview.DispatchGoCommand != "backend-adapter-redacted-profile-audit-preview" {
		t.Fatalf("unexpected redacted adapter profile coverage route state: %#v", preview)
	}
	if preview.AuditType != "owner-local-redacted-adapter-profile-audit" ||
		preview.RouteDecision != "redacted-profile-route-ready" ||
		preview.ProfileCount != 3 ||
		!sameStrings(preview.ProfileIDs, []string{"automatic", "performance-priority", "compatibility-priority"}) ||
		preview.ContractProfileCount != 3 ||
		!preview.FullContractConsumed ||
		!preview.KDEFacingProjectionConsumed ||
		!preview.InternalAdapterIDsRedacted ||
		!preview.InternalProfilePathsRedacted ||
		!preview.OwnerLocalRouteCandidateReady ||
		!preview.FullContractFixtureLocal ||
		preview.ProductionDBusExposureReady ||
		preview.CallerStateRootRequired {
		t.Fatalf("unexpected redacted adapter profile coverage audit state: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 8 ||
		!preview.AllChecksPassed ||
		!preview.SmokeCoverageReady ||
		!sameStrings(preview.CheckIDs, []string{"smoke-record-present", "service-call-read-dispatch", "owner-route-payload-present", "redacted-profiles-preserved", "full-contract-fixture-local", "production-dbus-blocked", "caller-paths-hidden", "unsafe-gates-closed"}) {
		t.Fatalf("unexpected redacted adapter profile coverage checks: counts=%d/%d ids=%#v", preview.PassedCheckCount, preview.CheckCount, preview.CheckIDs)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.ReadOnlyCoverage ||
		preview.SystemServiceStarted ||
		preview.SessionBusClaimed ||
		preview.ProductionBusClaimed ||
		preview.WriteMethodsEnabled ||
		preview.RuntimeWritesEnabled ||
		preview.AdapterInvocationEnabled ||
		preview.BackendInstallEnabled ||
		preview.BackendDownloadEnabled ||
		preview.BackendLaunchEnabled ||
		preview.BackendProcessStarted ||
		preview.CommandMaterialized ||
		preview.ExecutablePathResolved ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired ||
		preview.StateRootPathExposed ||
		preview.RawCommandExposed ||
		preview.RawExecutableExposed ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected redacted adapter profile coverage safety gates: %#v", preview)
	}
	if err := validateNoBackendTerms(preview, "redacted adapter profile owner smoke coverage test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
}

func TestBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFailsClosedWithoutRecord(t *testing.T) {
	records, err := NewSmokeBatchRecords(projectRoot(t))
	if err != nil {
		t.Fatalf("NewSmokeBatchRecords returned error: %v", err)
	}
	filtered := make([]SmokeBatchRecord, 0, len(records))
	for _, record := range records {
		if record.Method == "GetBackendAdapterProfileAudit" {
			continue
		}
		filtered = append(filtered, record)
	}

	preview, err := NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords(projectRoot(t), filtered)
	if err != nil {
		t.Fatalf("NewBackendAdapterRedactedProfileOwnerSmokeCoveragePreviewFromRecords returned error: %v", err)
	}
	if preview.CoverageRecordFound ||
		preview.CoverageRecordSequence != 0 ||
		preview.ServiceCallReadDispatch ||
		preview.DispatchRouteReady ||
		preview.SmokeCoverageReady ||
		preview.AllChecksPassed {
		t.Fatalf("missing redacted adapter profile coverage record must fail closed: %#v", preview)
	}
	if preview.CheckCount != 8 ||
		preview.PassedCheckCount != 2 ||
		preview.Checks[0].ID != "smoke-record-present" ||
		preview.Checks[0].Status != "blocked" ||
		preview.Checks[7].ID != "unsafe-gates-closed" ||
		preview.Checks[7].Status != "pass" {
		t.Fatalf("unexpected fail-closed redacted adapter profile coverage checks: counts=%d/%d checks=%#v", preview.PassedCheckCount, preview.CheckCount, preview.Checks)
	}
}
