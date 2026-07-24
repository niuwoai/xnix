package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/winapp"
)

func TestKDEControlledLaunchActionPreviewCommandForwardsOnlyEvidenceHandle(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIEvidence(t, stateRoot)

	var output bytes.Buffer
	err := run([]string{
		"kde-controlled-launch-action-preview",
		"--state-root", stateRoot,
		"--evidence-relative-path", record.EvidenceRelativePath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("KDE controlled launch action output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.kde_controlled_launch_action.v1" ||
		payload["request_type"] != "kde-controlled-launch-action-preview" ||
		payload["kde_action_id"] != "xnix.runtime-status.controlled-launch" ||
		payload["kde_action_label"] != "Run with Xnix Runtime" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["kde_forwarded_argument_kind"] != "evidence-relative-path" ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["desktop_evidence_handle_forwarded"] != true ||
		payload["desktop_trigger_ready"] != true ||
		payload["runtime_owned"] != true ||
		payload["go_runtime_backed"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["kde_owns_owner_service_args"] != false ||
		payload["owner_service_args_exposed_to_kde"] != false ||
		payload["owner_service_boundary_hidden_from_kde"] != true ||
		payload["desktop_receipt_fields_reconstructed"] != false ||
		payload["desktop_kde_state_root_access"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["raw_launcher_output_exposed"] != false ||
		payload["host_root_modified"] != false ||
		payload["docker_socket_mounted"] != false ||
		payload["privileged_container_required"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["execution_started"] != false {
		t.Fatalf("unexpected KDE controlled launch action payload: %#v", payload)
	}

	forwardedArgs, ok := payload["kde_forwarded_arguments"].([]any)
	if !ok || len(forwardedArgs) != 1 || forwardedArgs[0] != record.EvidenceRelativePath {
		t.Fatalf("KDE action must forward only the evidence handle: %#v", payload)
	}
	if strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), stateRoot) {
		t.Fatalf("KDE action output exposed Runtime-owned details: %s", output.String())
	}
}

func TestKDEControlledLaunchActionPreviewCommandConsumesKDEGUICardFile(t *testing.T) {
	stateRoot := t.TempDir()
	record := recordKDEControlledLaunchActionCLIGUIEvidence(t, stateRoot)
	card := kdeControlledLaunchActionCLIGUICard(record)
	cardPath := filepath.Join(t.TempDir(), "messagebox-gui-card.json")
	encodedCard, err := json.Marshal(card)
	if err != nil {
		t.Fatalf("Marshal card returned error: %v", err)
	}
	if err := os.WriteFile(cardPath, encodedCard, 0o600); err != nil {
		t.Fatalf("WriteFile card returned error: %v", err)
	}

	var output bytes.Buffer
	err = run([]string{
		"kde-controlled-launch-action-preview",
		"--state-root", stateRoot,
		"--kde-gui-card-file", cardPath,
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("KDE controlled launch action output must be JSON: %v\n%s", err, output.String())
	}
	if payload["application_id"] != "org.xnix.apps.messagebox" ||
		payload["application_name"] != "Xnix MessageBox" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["desktop_callable_route"] != "kde-dbus-runtime-status-action" ||
		payload["desktop_callable_runtime_method"] != "ShowRuntimeControlledLaunch" ||
		payload["desktop_callable_execution_type"] != appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["owner_service_args_exposed_to_kde"] != false ||
		payload["state_root_path_exposed"] != false ||
		payload["backend_details_exposed"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["execution_started"] != false {
		t.Fatalf("unexpected KDE GUI card action payload: %#v", payload)
	}
	forwardedArgs, ok := payload["kde_forwarded_arguments"].([]any)
	if !ok || len(forwardedArgs) != 1 || forwardedArgs[0] != record.EvidenceRelativePath {
		t.Fatalf("KDE GUI card action must forward only the evidence handle: %#v", payload)
	}
	if strings.Contains(output.String(), "owner_service_call_args") ||
		strings.Contains(output.String(), stateRoot) ||
		strings.Contains(output.String(), cardPath) {
		t.Fatalf("KDE GUI card action output exposed Runtime-owned details: %s", output.String())
	}
}

func TestKDEControlledLaunchActionPreviewCommandRejectsMissingStateRoot(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"kde-controlled-launch-action-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --state-root") {
		t.Fatalf("missing state root must be rejected, got: %v", err)
	}
}

func recordKDEControlledLaunchActionCLIEvidence(t *testing.T, stateRoot string) appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord {
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

func recordKDEControlledLaunchActionCLIGUIEvidence(t *testing.T, stateRoot string) appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord {
	t.Helper()

	appVersion := currentProjectVersion(t)
	sessionID := appidentity.KnownAppControlledExecutionSessionID("org.xnix.apps.messagebox", appVersion)
	launchReceiptID := appidentity.KnownAppLaunchAuthorizationReceiptID("org.xnix.apps.messagebox", appVersion)
	reviewReceiptID := appidentity.KnownAppSessionGatedLaunchReviewReceiptID("org.xnix.apps.messagebox", appVersion, sessionID)
	projection, err := appidentity.ProjectKnownAppKDERuntimeStatusLaunchDelegatedEvidence(appidentity.KnownAppKDERuntimeStatusLaunchDelegatedEvidenceRequest{
		AppID:                                  "org.xnix.apps.messagebox",
		DisplayName:                            "Xnix MessageBox",
		AppVersion:                             appVersion,
		RequestType:                            winapp.KnownDispatchSmokeRequestType,
		Status:                                 "passed",
		EvidenceSource:                         "wine-guest-gui-smoke",
		GuestBoundary:                          winapp.KnownDispatchGuestBoundary,
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
		ControlledSessionWindowObserved:        true,
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

func kdeControlledLaunchActionCLIGUICard(record appidentity.KnownAppKDERuntimeStatusLaunchEvidenceRecord) appidentity.KDECenterPageKnownAppMatrixCard {
	return appidentity.KDECenterPageKnownAppMatrixCard{
		AppID:                                "org.xnix.apps.messagebox",
		DisplayName:                          "Xnix MessageBox",
		AppVersion:                           record.AppVersion,
		EvidenceKind:                         "known-application-gui-smoke",
		EvidenceSource:                       "wine-guest-gui-smoke",
		SmokeStatus:                          "passed",
		CompatibilityState:                   "owner-controlled-gui-qemu-wine-verified",
		CenterCardState:                      "validated-owner-controlled-gui-runtime-run",
		PrimaryActionID:                      appidentity.KnownAppKDERuntimeStatusLaunchAction,
		PrimaryActionLabel:                   "Show Runtime-controlled launch",
		PrimaryActionKind:                    "runtime-status",
		PrimaryActionEnabled:                 true,
		DesktopCallableRoute:                 "kde-dbus-runtime-status-action",
		DesktopCallableRuntimeMethod:         "ShowRuntimeControlledLaunch",
		DesktopCallableExecutionType:         appidentity.KnownAppKDERuntimeStatusLaunchExecutionRequestType,
		DesktopDBusMethod:                    "org.xnix.Compatibility1.ShowRuntimeControlledLaunch",
		DesktopEvidenceHandleForwarded:       true,
		KDEForwardedArgumentKind:             "evidence-relative-path",
		KDEForwardedArguments:                []string{record.EvidenceRelativePath},
		OwnerServiceArgsExposedToKDE:         false,
		ExecutionEvidenceRecorded:            true,
		StagedLauncherVerified:               true,
		OwnerControlledRuntimeLaunchVerified: true,
		OwnerManagedCopyVerified:             true,
		OwnerServiceCallReady:                true,
		OwnerEvidenceHandoffReady:            true,
		OwnerEvidenceRelativePath:            record.EvidenceRelativePath,
		RuntimeDispatchVerified:              true,
		LaunchAuthorizationRequired:          true,
		DesktopLaunchEnabled:                 false,
		BackendLaunchEnabled:                 false,
		RuntimeOwned:                         true,
		GoRuntimeBacked:                      true,
		KDEPolicyOwner:                       false,
		HostRootModified:                     false,
		BackendDetailsExposed:                false,
		RawArtifactPathExposed:               false,
		Summary:                              "Xnix MessageBox GUI card can forward only a safe Runtime evidence handle.",
	}
}
