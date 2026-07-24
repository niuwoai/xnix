package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKDEControlledLaunchActionSurfaceAuditPreviewCommandAcceptsRepositoryMetadata(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{
		"kde-controlled-launch-action-surface-audit-preview",
		"--desktop-entry-file", filepath.Join(projectRootForRuntimeServiceBindingCommandTest(t), "kde/actions/xnix-runtime-status-controlled-launch.desktop"),
		"--evidence-handle", "runtime/kde-runtime-status-launch-evidence/fixture.json",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("surface audit output must be JSON: %v\n%s", err, output.String())
	}
	if payload["schema_version"] != "xnix.runtime.kde_controlled_launch_action_surface_audit.v1" ||
		payload["request_type"] != "kde-controlled-launch-action-surface-audit-preview" ||
		payload["audit_state"] != "safe" ||
		payload["surface_safe_for_human_smoke"] != true ||
		payload["kde_action_id"] != "xnix.runtime-status.controlled-launch" ||
		payload["public_dbus_method"] != "org.xnix.Compatibility1.ShowRuntimeControlledLaunch" ||
		payload["forwarded_argument_kind"] != "evidence-relative-path" ||
		payload["evidence_handle_shape_accepted"] != true ||
		payload["evidence_only_argument_shape"] != true ||
		payload["required_metadata_present"] != true ||
		payload["metadata_malformed"] != false ||
		payload["unsafe_owner_arguments_present"] != false ||
		payload["unsafe_backend_terms_present"] != false ||
		payload["kde_forwards_only_evidence_handle"] != true ||
		payload["kde_policy_owner"] != false ||
		payload["owner_service_args_exposed_to_kde"] != false ||
		payload["state_root_access"] != false ||
		payload["receipt_reconstruction"] != false ||
		payload["dbus_called"] != false ||
		payload["kde_configuration_written"] != false ||
		payload["runtime_state_written"] != false ||
		payload["desktop_launch_enabled"] != false ||
		payload["backend_launch_enabled"] != false ||
		payload["execution_started"] != false ||
		payload["host_root_modified"] != false ||
		payload["network_required"] != false {
		t.Fatalf("unexpected surface audit payload: %#v", payload)
	}
	if strings.Contains(output.String(), "kde/actions/xnix-runtime-status-controlled-launch.desktop") ||
		strings.Contains(output.String(), "--state-root") ||
		strings.Contains(output.String(), "owner_service_call_args") {
		t.Fatalf("surface audit output exposed unsafe or local details: %s", output.String())
	}
}

func TestKDEControlledLaunchActionSurfaceAuditPreviewCommandBlocksUnsafeMetadataWithoutLeakingValues(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "unsafe.desktop")
	contents := strings.Replace(safeKDEControlledLaunchActionSurfaceCLIMetadata(), "X-Xnix-State-Root-Access=false", "X-Xnix-State-Root-Access=true\nX-Xnix-Unsafe-Example=--state-root /private/runtime", 1)
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	var output bytes.Buffer
	err := run([]string{
		"kde-controlled-launch-action-surface-audit-preview",
		"--desktop-entry-file", path,
		"--evidence-handle", "runtime/kde-runtime-status-launch-evidence/fixture.json",
	}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(output.Bytes(), &payload); err != nil {
		t.Fatalf("surface audit output must be JSON: %v\n%s", err, output.String())
	}
	if payload["audit_state"] != "unsafe-owner-args" ||
		payload["surface_safe_for_human_smoke"] != false ||
		payload["unsafe_owner_arguments_present"] != true ||
		payload["state_root_access"] != true ||
		payload["dbus_called"] != false ||
		payload["runtime_state_written"] != false ||
		payload["host_root_modified"] != false {
		t.Fatalf("unexpected unsafe owner-args surface audit payload: %#v", payload)
	}
	for _, unsafe := range []string{path, "/private/runtime", "--state-root"} {
		if strings.Contains(output.String(), unsafe) {
			t.Fatalf("surface audit output leaked unsafe value %q: %s", unsafe, output.String())
		}
	}
}

func TestKDEControlledLaunchActionSurfaceAuditPreviewCommandRequiresDesktopEntryFile(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"kde-controlled-launch-action-surface-audit-preview"}, &output)
	if err == nil || !strings.Contains(err.Error(), "requires --desktop-entry-file") {
		t.Fatalf("missing desktop entry file must be rejected, got: %v", err)
	}
}

func safeKDEControlledLaunchActionSurfaceCLIMetadata() string {
	return `[Desktop Entry]
Type=Service
Name=Xnix Runtime controlled launch
Comment=Forward a Runtime-status evidence handle to the Xnix Runtime D-Bus controlled-launch action.
Icon=media-playback-start
X-Xnix-KDE-Action-ID=xnix.runtime-status.controlled-launch
X-Xnix-Runtime-Preview=xnix-runtime-go kde-controlled-launch-action-preview
X-Xnix-Restricted-Smoke-Plan=xnix-runtime-go kde-controlled-launch-session-bus-smoke-plan-preview
X-Xnix-DBus-Service=org.xnix.Compatibility1
X-Xnix-DBus-Object-Path=/org/xnix/Compatibility1
X-Xnix-DBus-Method=org.xnix.Compatibility1.ShowRuntimeControlledLaunch
X-Xnix-Forwarded-Argument=evidence-relative-path
X-Xnix-Forwards-Only-Evidence-Handle=true
X-Xnix-KDE-Policy-Owner=false
X-Xnix-Owner-Service-Args-Exposed-To-KDE=false
X-Xnix-State-Root-Access=false
X-Xnix-Receipt-Reconstruction=false
X-Xnix-Backend-Launch-Enabled=false
X-Xnix-Execution-Started=false
X-Xnix-Host-Root-Modified=false
X-Xnix-Docker-Socket-Mounted=false
X-Xnix-Privileged-Container-Required=false
X-Xnix-Host-Network-Required=false
`
}
