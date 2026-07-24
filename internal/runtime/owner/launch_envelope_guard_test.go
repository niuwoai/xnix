package owner

import (
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestPreviewLaunchEnvelopeGuardAcceptsEvidenceOnlyEnvelope(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)

	preview, err := PreviewLaunchEnvelopeGuard(LaunchEnvelopeGuardRequest{
		StateRoot:              stateRoot,
		RouteID:                launchEnvelopeExpectedRoute,
		MethodID:               launchEnvelopeExpectedMethod,
		EvidenceHandleKind:     "evidence-relative-path",
		EvidenceHandle:         record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		ActionID:               launchEnvelopeExpectedActionID,
		CallerRole:             launchEnvelopeExpectedCallerRole,
		RequestFreshness:       "fresh",
		ReplayMarker:           "fresh",
	})
	if err != nil {
		t.Fatalf("PreviewLaunchEnvelopeGuard returned error: %v", err)
	}
	if preview.SchemaVersion != LaunchEnvelopeGuardSchemaVersion ||
		preview.RequestType != LaunchEnvelopeGuardRequestType ||
		preview.GuardState != "accepted-for-review" ||
		!preview.AcceptedForReview ||
		!preview.RouteMatched ||
		!preview.MethodMatched ||
		!preview.ActionMatched ||
		!preview.CallerRoleAccepted ||
		!preview.EvidenceHandleShapeAccepted ||
		!preview.EvidenceDigestVerified ||
		!preview.EvidenceDigestMatched ||
		preview.EvidenceSHA256 != record.EvidenceSHA256 ||
		preview.OwnerOnlyArgumentsPresent ||
		preview.OwnerOnlyArgumentCount != 0 ||
		!preview.KDECanOnlyForwardEvidenceHandle ||
		!preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		preview.DesktopReceiptFieldsReconstructed ||
		preview.DesktopKDEStateRootAccess ||
		preview.StateRootPathExposed ||
		preview.CacheRootPathExposed ||
		preview.LauncherPathExposed ||
		preview.RawExecutablePathExposed ||
		preview.RawCommandExposed ||
		preview.BackendDetailsExposed ||
		preview.ReceiptWriteEnabled ||
		preview.ReceiptAccepted ||
		preview.ProductionAuthorizationAccepted ||
		preview.ServiceCallDispatched ||
		preview.DBusOwnershipEnabled ||
		preview.DesktopLaunchEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionStarted ||
		preview.BackendProcessStarted ||
		preview.RuntimeStateWritten ||
		preview.NetworkRequired ||
		preview.HostRootModified ||
		preview.PrivilegedContainerRequired {
		t.Fatalf("unexpected accepted launch envelope guard: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), stateRoot) || strings.Contains(string(encoded), "owner_service_call_args") || strings.Contains(string(encoded), " --service-call ") {
		t.Fatalf("launch envelope guard exposed Runtime-owned details: %s", string(encoded))
	}
}

func TestPreviewLaunchEnvelopeGuardBlocksMismatchedRouteStaleReplayOwnerArgsAndMalformed(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	base := LaunchEnvelopeGuardRequest{
		StateRoot:              stateRoot,
		RouteID:                launchEnvelopeExpectedRoute,
		MethodID:               launchEnvelopeExpectedMethod,
		EvidenceHandleKind:     "evidence-relative-path",
		EvidenceHandle:         record.EvidenceRelativePath,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		ActionID:               launchEnvelopeExpectedActionID,
		CallerRole:             launchEnvelopeExpectedCallerRole,
		RequestFreshness:       "fresh",
		ReplayMarker:           "fresh",
	}

	cases := []struct {
		name string
		edit func(*LaunchEnvelopeGuardRequest)
		want string
	}{
		{name: "route", edit: func(request *LaunchEnvelopeGuardRequest) { request.RouteID = "unexpected-route" }, want: "blocked-mismatched-route"},
		{name: "method", edit: func(request *LaunchEnvelopeGuardRequest) { request.MethodID = "GetApplications" }, want: "blocked-mismatched-route"},
		{name: "action", edit: func(request *LaunchEnvelopeGuardRequest) { request.ActionID = "wrong-action" }, want: "blocked-mismatched-route"},
		{name: "caller", edit: func(request *LaunchEnvelopeGuardRequest) { request.CallerRole = "unknown-shell" }, want: "unsupported"},
		{name: "freshness", edit: func(request *LaunchEnvelopeGuardRequest) { request.RequestFreshness = "stale" }, want: "blocked-stale"},
		{name: "replay", edit: func(request *LaunchEnvelopeGuardRequest) { request.ReplayMarker = "seen" }, want: "blocked-replay"},
		{name: "owner-args", edit: func(request *LaunchEnvelopeGuardRequest) {
			request.KDEStateRoot = "/tmp/runtime"
			request.KDECacheRoot = "/tmp/cache"
			request.KDELauncherPath = "/tmp/launcher"
			request.KDETimeoutValue = "5m"
			request.KDERawExecutablePath = "C:/private/app"
			request.KDEBackendCommand = "private engine command"
			request.KDEReceiptID = "receipt-1"
			request.KDESessionID = "session-1"
			request.KDEDispatchID = "dispatch-1"
		}, want: "blocked-owner-args"},
		{name: "malformed-handle", edit: func(request *LaunchEnvelopeGuardRequest) { request.EvidenceHandle = "../escape.json" }, want: "malformed"},
		{name: "missing-evidence", edit: func(request *LaunchEnvelopeGuardRequest) {
			request.EvidenceHandle = "runtime/kde-runtime-status-launch-evidence/missing.json"
		}, want: "blocked-missing-evidence"},
		{name: "stale-digest", edit: func(request *LaunchEnvelopeGuardRequest) { request.ExpectedEvidenceSHA256 = strings.Repeat("0", 64) }, want: "blocked-stale"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := base
			tc.edit(&request)
			preview, err := PreviewLaunchEnvelopeGuard(request)
			if err != nil {
				t.Fatalf("PreviewLaunchEnvelopeGuard returned error: %v", err)
			}
			if preview.GuardState != tc.want || preview.AcceptedForReview || preview.ServiceCallDispatched || preview.RuntimeStateWritten || preview.HostRootModified {
				t.Fatalf("unexpected %s guard: %#v", tc.name, preview)
			}
			if tc.name == "owner-args" {
				if preview.OwnerOnlyArgumentCount != 9 ||
					!preview.KDEStateRootRejected ||
					!preview.KDECacheRootRejected ||
					!preview.KDELauncherPathRejected ||
					!preview.KDETimeoutValueRejected ||
					!preview.KDERawExecutablePathRejected ||
					!preview.KDEBackendCommandRejected ||
					!preview.KDEReceiptFieldsRejected ||
					!preview.KDESessionFieldsRejected ||
					!preview.KDEDispatchFieldsRejected {
					t.Fatalf("owner-only arguments were not rejected: %#v", preview)
				}
				encoded, err := json.Marshal(preview)
				if err != nil {
					t.Fatalf("Marshal returned error: %v", err)
				}
				for _, unsafe := range []string{"/tmp/runtime", "/tmp/cache", "/tmp/launcher", "C:/private/app", "private engine command", "receipt-1", "session-1", "dispatch-1"} {
					if strings.Contains(string(encoded), unsafe) {
						t.Fatalf("guard leaked owner-only value %q: %s", unsafe, string(encoded))
					}
				}
			}
		})
	}
}

func TestPreviewLaunchEnvelopeGuardAcceptsOpaqueEvidenceID(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordLaunchEnvelopeGuardEvidence(t, stateRoot)
	preview, err := PreviewLaunchEnvelopeGuard(LaunchEnvelopeGuardRequest{
		StateRoot:              stateRoot,
		RouteID:                launchEnvelopeExpectedRoute,
		MethodID:               launchEnvelopeExpectedMethod,
		EvidenceHandleKind:     "evidence-id",
		EvidenceHandle:         record.EvidenceID,
		ExpectedEvidenceSHA256: record.EvidenceSHA256,
		ActionID:               launchEnvelopeExpectedActionID,
		CallerRole:             launchEnvelopeExpectedCallerRole,
	})
	if err != nil {
		t.Fatalf("PreviewLaunchEnvelopeGuard returned error: %v", err)
	}
	if preview.GuardState != "accepted-for-review" || !preview.AcceptedForReview || preview.EvidenceSHA256 != record.EvidenceSHA256 {
		t.Fatalf("unexpected opaque-id envelope guard: %#v", preview)
	}
}

func recordLaunchEnvelopeGuardEvidence(t *testing.T, stateRoot string) appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

	sessionID := appidentity.KnownAppControlledExecutionSessionID("7zr", "26.02")
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("7zr", "26.02")
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("7zr", "26.02", sessionID)
	projection, err := appidentity.ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "7zr",
		DisplayName:                            "7-Zip standalone console executable",
		AppVersion:                             "26.02",
		RequestType:                            "windows-known-app-dispatch-smoke",
		Status:                                 "passed",
		GuestBoundary:                          "managed-known-app-guest-smoke",
		RuntimeOwnedDispatch:                   true,
		ArtifactVerified:                       true,
		MarkerObserved:                         true,
		SessionGatedControlledDispatchConsumed: true,
		SessionGatedControlledDispatchState:    "created-after-session-gated-review",
		SessionGatedReviewReceiptID:            reviewReceiptID,
		LaunchAuthorizationReceiptID:           launchReceiptID,
		LaunchAuthorizationReceiptState:        "recorded",
		LaunchGateState:                        "controlled-dispatch-ready",
		LaunchGateConsumed:                     true,
		LaunchGateReceiptAccepted:              true,
		LaunchGateGuestBoundaryAccepted:        true,
		ControlledDispatchReady:                true,
		ControlledExecutionSessionConsumed:     true,
		ControlledExecutionSessionID:           sessionID,
		ControlledSessionDigestVerified:        true,
		ControlledSessionRelativePath:          "execution-ledger/sessions/" + sessionID + ".json",
		RuntimeOwnerConsumableSession:          true,
		KDEReadModelConsumableSession:          true,
	})
	if err != nil {
		t.Fatalf("ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence returned error: %v", err)
	}
	record, err := appidentity.RecordKnownAppKDERuntimeStatusLaunchEvidence(appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecordRequest{
		StateRoot:  stateRoot,
		Projection: projection,
	})
	if err != nil {
		t.Fatalf("RecordKnownAppKDERuntimeStatusLaunchEvidence returned error: %v", err)
	}
	return record
}
