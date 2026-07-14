package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestExecutionResourceGrantPreviewKeepsPermissionsUngrantable(t *testing.T) {
	plan, err := NewPlan(Recipe{
		ID:                  "org.example.ledger",
		Name:                "Example Ledger",
		Icon:                "office-chart-area",
		Mode:                "automatic",
		SupportedExtensions: []string{".xls"},
	})
	if err != nil {
		t.Fatalf("NewPlan returned error: %v", err)
	}

	preview, err := plan.ExecutionResourceGrantPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("ExecutionResourceGrantPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.resource_grant.v1" ||
		preview.RequestType != "execution-resource-grant-preview" ||
		preview.GrantType != "compatibility-launch-resource-grant" ||
		preview.RequestState != "blocked" ||
		preview.Source != "execution-preflight-preview" ||
		preview.RuntimeMethod != "Launch" ||
		preview.ReadMethod != "GetExecutionResourceGrantPreview" {
		t.Fatalf("unexpected execution resource grant schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected execution resource grant identity: %#v", preview)
	}
	if preview.ExecutionPreflight.SchemaVersion != "xnix.runtime.launch_preflight.v1" ||
		preview.ExecutionPreflight.RequestType != "execution-preflight-preview" ||
		preview.ExecutionPreflight.PreflightType != "compatibility-launch-preflight" ||
		!preview.ExecutionPreflight.UserDecisionAllowsLaunch ||
		preview.ExecutionPreflight.PreflightPassed ||
		!preview.ExecutionPreflight.PortalRequired ||
		!preview.ExecutionPreflight.SnapshotRequired ||
		preview.ExecutionPreflight.WriteGateOpen ||
		preview.ExecutionPreflight.RuntimeLaunchApproval ||
		preview.ExecutionPreflight.PermissionGranted {
		t.Fatalf("unexpected preflight summary: %#v", preview.ExecutionPreflight)
	}
	if preview.GrantCount != 7 ||
		preview.AllowCount != 1 ||
		preview.AskCount != 5 ||
		preview.DenyCount != 1 ||
		preview.PortalRequiredCount != 6 ||
		preview.PendingReviewCount != 5 {
		t.Fatalf("unexpected grant counts: %#v", preview)
	}
	if len(preview.ResourceGrants) != 7 {
		t.Fatalf("unexpected resource grants: %#v", preview.ResourceGrants)
	}
	grantsByID := map[string]ExecutionResourceGrant{}
	for _, grant := range preview.ResourceGrants {
		grantsByID[grant.ID] = grant
		if grant.RequestObjectCreated ||
			grant.PermissionGranted ||
			grant.DirectAccessAllowed ||
			grant.BackendDetailsExposed {
			t.Fatalf("grant enabled unsafe access: %#v", grant)
		}
	}
	if grantsByID["documents"].GrantState != "requires-review" ||
		!grantsByID["documents"].PortalRequired ||
		grantsByID["downloads"].GrantState != "requires-review" ||
		grantsByID["camera"].GrantState != "planned-deny" ||
		grantsByID["network"].GrantState != "planned-allow" ||
		grantsByID["network"].PortalRequired {
		t.Fatalf("unexpected resource grant states: %#v", grantsByID)
	}
	if preview.BridgeSummary.RequestType != "desktop-resource-bridge-preview" ||
		preview.BridgeSummary.PlanType != "desktop-resource-bridge-plan" ||
		preview.BridgeSummary.BridgeState != "planned" ||
		preview.BridgeSummary.ResourceCount != 5 ||
		!preview.BridgeSummary.PortalMediated ||
		preview.BridgeSummary.ResourceBridgesEnabled ||
		preview.BridgeSummary.RequestObjectsCreated ||
		preview.BridgeSummary.DirectHostFileAccess {
		t.Fatalf("unexpected bridge summary: %#v", preview.BridgeSummary)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.CompatibilityCenterCard || !preview.SafeForAIDiagnostics ||
		!preview.DesktopEntryLaunchVisible || !preview.LaunchIntentCaptured ||
		!preview.UserDecisionCaptured || !preview.UserDecisionAllowsLaunch ||
		preview.PreflightPassed || !preview.PortalReviewRequired ||
		!preview.SnapshotRequired || !preview.GrantPlanCreated ||
		preview.GrantObjectsCreated || preview.RequestObjectsCreated ||
		preview.PermissionGranted || preview.ResourceBridgesEnabled ||
		preview.SettingsPersisted || preview.RuntimeLaunchApproval ||
		preview.LaunchAllowed || preview.LaunchEnabled ||
		preview.ExecutionStarted || preview.HostPermissionChanged ||
		preview.HostRootModified || preview.NetworkRequired ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected execution resource grant safety flags: %#v", preview)
	}
	if !containsString(preview.BlockedActions, "create resource grant objects from preview") ||
		!containsString(preview.BlockedActions, "enable desktop resource bridges before Portal review") ||
		!containsString(preview.BlockedActions, "treat allowed network policy as launch approval") {
		t.Fatalf("unexpected blocked actions: %#v", preview.BlockedActions)
	}

	rejectedPreview, err := plan.ExecutionResourceGrantPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected ExecutionResourceGrantPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.ExecutionPreflight.UserDecisionAllowsLaunch ||
		rejectedPreview.GrantObjectsCreated ||
		rejectedPreview.PermissionGranted {
		t.Fatalf("unexpected rejected grant preview: %#v", rejectedPreview)
	}

	if _, err := plan.ExecutionResourceGrantPreview("invalid", nil); err == nil {
		t.Fatalf("ExecutionResourceGrantPreview accepted an invalid decision")
	}
	if _, err := plan.ExecutionResourceGrantPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("ExecutionResourceGrantPreview accepted a non-file URI")
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("execution resource grant preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}
