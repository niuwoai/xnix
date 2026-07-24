package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandConsumesVerifiedHandoff(t *testing.T) {
	stateRoot := t.TempDir()
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
	var output bytes.Buffer
	err = run([]string{
		"known-app-runtime-status-launch-owner-trigger-preview",
		"--state-root", stateRoot,
		"--evidence-relative-path", record.EvidenceRelativePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("trigger output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.known_app_runtime_status_launch_owner_trigger.v1" ||
		payload["request_type"] != "known-app-runtime-status-launch-owner-trigger-preview" ||
		payload["desktop_trigger_ready"] != true ||
		payload["desktop_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["owner_service_call_ready"] != true ||
		payload["runtime_owner_service_supplies_inputs"] != true ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false {
		t.Fatalf("unexpected owner trigger payload: %#v", payload)
	}
	ownerServiceArgs, ok := payload["owner_service_call_args"].([]any)
	if !ok || len(ownerServiceArgs) != 3 ||
		ownerServiceArgs[0] != "ShowRuntimeControlledLaunch" ||
		ownerServiceArgs[1] != "evidence-relative-path" ||
		ownerServiceArgs[2] != record.EvidenceRelativePath {
		t.Fatalf("unexpected owner service args: %#v", payload)
	}
	ownerServiceCLIArgs, ok := payload["owner_service_cli_args"].([]any)
	if !ok || len(ownerServiceCLIArgs) != 4 ||
		ownerServiceCLIArgs[0] != "--service-call" ||
		ownerServiceCLIArgs[1] != "ShowRuntimeControlledLaunch" ||
		ownerServiceCLIArgs[2] != "evidence-relative-path" ||
		ownerServiceCLIArgs[3] != record.EvidenceRelativePath {
		t.Fatalf("unexpected owner service CLI args: %#v", payload)
	}
	if strings.Contains(output.String(), stateRoot) {
		t.Fatalf("trigger output exposed state root path: %s", output.String())
	}
}

func TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"known-app-runtime-status-launch-owner-trigger-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}
