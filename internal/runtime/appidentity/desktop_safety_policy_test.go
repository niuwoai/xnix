package appidentity

import (
	"strings"
	"testing"
)

func TestDesktopSafetyPolicyPreviewOwnsKDEFirstUserFacingSafetyPolicy(t *testing.T) {
	preview := NewDesktopSafetyPolicyPreview()

	if preview.SchemaVersion != "xnix.runtime.desktop_safety_policy.v1" ||
		preview.RequestType != "desktop-safety-policy-preview" ||
		preview.PolicyType != "kde-first-user-facing-safety-policy" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEIntegrationStatus" {
		t.Fatalf("unexpected desktop safety policy identity: %#v", preview)
	}
	if preview.EntryPointCount != 7 ||
		strings.Join(preview.EntryPoints, ",") != "launcher,task-manager,file-manager,system-tray,notifications,compatibility-center,settings" {
		t.Fatalf("unexpected KDE-first entrypoints: %#v", preview.EntryPoints)
	}
	if strings.Join(preview.SettingsFieldIDs, ",") != "mode,preference,documents,downloads,camera,network,snapshots" {
		t.Fatalf("unexpected user-facing settings fields: %#v", preview.SettingsFieldIDs)
	}
	for _, term := range []string{"prefix", "bottle", "wine", "proton", ".exe", "Program Files", "qemu-system", "/Users/", "/private/", "/var/", ".wine", "docker.sock"} {
		if !containsString(preview.ForbiddenUserTerms, term) {
			t.Fatalf("desktop safety policy missing forbidden user term %q: %#v", term, preview.ForbiddenUserTerms)
		}
	}
	for _, key := range []string{"backend_launch_enabled", "execution_started", "host_root_modified", "real_portal_transport_enabled", "request_object_created", "settings_persisted"} {
		if !containsString(preview.SafetyFalseKeys, key) {
			t.Fatalf("desktop safety policy missing false safety key %q: %#v", key, preview.SafetyFalseKeys)
		}
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		!preview.UserVisible ||
		!preview.BackendTerminologyHidden ||
		preview.KDEPolicyOwner ||
		preview.WriteMethodsEnabled ||
		preview.BackendLaunchEnabled ||
		preview.ExecutionEnabled ||
		preview.RealPortalTransport ||
		preview.AIProviderCallEnabled ||
		preview.HostRootModified ||
		preview.NetworkRequired ||
		preview.PrivilegedContainerNeeded {
		t.Fatalf("unexpected desktop safety policy flags: %#v", preview)
	}
}

func TestDesktopSafetyPolicyPreviewReturnsDefensiveCopies(t *testing.T) {
	preview := NewDesktopSafetyPolicyPreview()
	preview.EntryPoints[0] = "mutated"
	preview.SettingsFieldIDs[0] = "mutated"
	preview.ForbiddenUserTerms[0] = "mutated"
	preview.SafetyFalseKeys[0] = "mutated"

	next := NewDesktopSafetyPolicyPreview()
	if next.EntryPoints[0] != "launcher" ||
		next.SettingsFieldIDs[0] != "mode" ||
		next.ForbiddenUserTerms[0] != "prefix" ||
		next.SafetyFalseKeys[0] != "backend_details_exposed" {
		t.Fatalf("desktop safety policy did not return defensive copies: %#v", next)
	}
}
