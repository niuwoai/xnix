package appidentity

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestKDEActionDependencyGraphPreviewMapsActionEvidence(t *testing.T) {
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

	preview, err := plan.KDEActionDependencyGraphPreview("approved", []string{"file:///home/test/Documents/book.xls"})
	if err != nil {
		t.Fatalf("KDEActionDependencyGraphPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kde_action_dependency_graph.v1" ||
		preview.RequestType != "kde-action-dependency-graph-preview" ||
		preview.GraphType != "compatibility-center-action-dependency-graph" ||
		preview.Source != "kde-action-queue-preview+runtime-evidence-prerequisites+runtime-write-gate-preview" ||
		preview.Desktop != "KDE Plasma" ||
		preview.RuntimeMethod != "GetKDEActionDependencyGraph" ||
		preview.ReadMethod != "GetKDEActionDependencyGraphPreview" {
		t.Fatalf("unexpected dependency graph schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.ApplicationName != "Example Ledger" ||
		preview.Icon != "office-chart-area" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.FileCount != 1 ||
		preview.FileURIs[0] != "file:///home/test/Documents/book.xls" {
		t.Fatalf("unexpected dependency graph identity: %#v", preview)
	}
	if preview.Queue.RequestType != "kde-action-queue-preview" ||
		preview.Queue.ActionCount != 7 ||
		preview.Queue.PendingActionCount != 7 ||
		preview.Queue.UserReviewRequiredCount != 4 ||
		preview.Queue.PortalActionCount != 1 ||
		preview.Queue.RuntimeGateActionCount != 7 ||
		!preview.Queue.ActionQueueCreated ||
		preview.Queue.ActionQueuePersisted ||
		preview.Queue.RuntimeLaunchApproval ||
		preview.Queue.LaunchAllowed ||
		preview.Queue.ExecutionStarted ||
		preview.Queue.BackendDetailsExposed {
		t.Fatalf("unexpected queue summary: %#v", preview.Queue)
	}
	if preview.ActionNodeCount != 7 ||
		preview.EvidenceNodeCount != 35 ||
		preview.GateNodeCount != 7 ||
		preview.NodeCount != 49 ||
		preview.EdgeCount != 42 ||
		preview.MissingEvidenceCount != 35 ||
		preview.BlockedActionCount != 7 ||
		preview.ReadOnlyCheckCount < 10 {
		t.Fatalf("unexpected graph counts: %#v", preview)
	}
	if len(preview.NodeIDs) != preview.NodeCount ||
		len(preview.MissingEvidenceIDs) != preview.MissingEvidenceCount ||
		!containsString(preview.BlockedActions, "review-launcher-action") ||
		!containsString(preview.BlockedActions, "review-file-manager-action") ||
		!containsString(preview.BlockedActions, "review-settings-action") {
		t.Fatalf("unexpected graph identifiers: %#v", preview)
	}
	fileAction := findKDEActionDependencyNode(preview.Nodes, "action:review-file-manager-action")
	if fileAction.NodeType != "action" ||
		fileAction.EntryPointID != "file-manager" ||
		fileAction.State != "blocked" ||
		!containsString(fileAction.RequiredEvidence, "portal-file-access-receipt") ||
		!containsString(fileAction.MissingEvidence, "portal-file-access-receipt") ||
		!containsString(fileAction.BlockedReasons, "Portal permission receipt is missing") ||
		fileAction.NextSafeCheck != "kde-action-card-preview" ||
		fileAction.ExecutionEnabled ||
		fileAction.RequestObjectsCreated ||
		fileAction.PermissionGrantCreated ||
		fileAction.SettingsPersisted ||
		fileAction.RuntimeLaunchApproval ||
		fileAction.BackendProcessStarted ||
		fileAction.HostRootModified ||
		fileAction.BackendDetailsExposed {
		t.Fatalf("unexpected file action node: %#v", fileAction)
	}
	portalEvidence := findKDEActionDependencyNode(preview.Nodes, "evidence:review-file-manager-action:portal-file-access-receipt")
	if portalEvidence.NodeType != "evidence" ||
		portalEvidence.State != "missing" ||
		portalEvidence.Title != "Portal file access receipt" ||
		portalEvidence.NextSafeCheck != "portal-request-preview" ||
		portalEvidence.PermissionGrantCreated ||
		portalEvidence.HostRootModified {
		t.Fatalf("unexpected Portal evidence node: %#v", portalEvidence)
	}
	gate := findKDEActionDependencyNode(preview.Nodes, "gate:review-file-manager-action:portal-file-open-review")
	if gate.NodeType != "gate" ||
		gate.State != "blocked" ||
		gate.NextSafeCheck != "runtime-write-gate-preview" ||
		gate.ExecutionEnabled ||
		gate.RuntimeLaunchApproval ||
		gate.HostRootModified {
		t.Fatalf("unexpected gate node: %#v", gate)
	}
	if !containsKDEActionDependencyEdge(preview.Edges, "evidence:review-file-manager-action:portal-file-access-receipt", "action:review-file-manager-action", "required-before-action") ||
		!containsKDEActionDependencyEdge(preview.Edges, "gate:review-file-manager-action:portal-file-open-review", "action:review-file-manager-action", "blocked-by-runtime-gate") {
		t.Fatalf("missing dependency edges: %#v", preview.Edges)
	}
	if !preview.ReceiptValidation.RejectsMismatchedAppID ||
		!preview.ReceiptValidation.RejectsMalformedOperation ||
		!preview.ReceiptValidation.RejectsPathEscapeEvidence ||
		!preview.ReceiptValidation.RejectsUnsafeSideEffects ||
		preview.ReceiptValidation.ReceiptRecorded ||
		preview.ReceiptValidation.RequestObjectsCreated ||
		preview.ReceiptValidation.PermissionGrantCreated ||
		preview.ReceiptValidation.SettingsPersisted ||
		preview.ReceiptValidation.ExecutionStarted ||
		preview.ReceiptValidation.HostRootModified ||
		preview.ReceiptValidation.BackendDetailsExposed {
		t.Fatalf("unexpected receipt validation: %#v", preview.ReceiptValidation)
	}
	if !preview.RuntimeOwned || !preview.GoRuntimeBacked || preview.KDEPolicyOwner ||
		!preview.OfficialDesktopOnly || !preview.CompatibilityCenterGraph ||
		!preview.SafeForAIDiagnostics || !preview.UserDecisionCaptured ||
		!preview.UserDecisionAllowsLaunch || !preview.DependencyGraphCreated ||
		preview.DependencyGraphPersisted || preview.SettingsPersisted ||
		preview.PermissionGrantCreated || preview.RequestObjectsCreated ||
		preview.RuntimeLaunchApproval || preview.LaunchAllowed ||
		preview.LaunchEnabled || preview.ExecutionStarted ||
		preview.BackendProcessStarted || preview.NetworkRequired ||
		preview.HostRootModified || preview.StateRootPathExposed ||
		preview.RawCommandExposed || preview.FileContentRead ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected dependency graph safety flags: %#v", preview)
	}
}

func TestKDEActionDependencyGraphPreviewValidationAndSafeText(t *testing.T) {
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

	rejectedPreview, err := plan.KDEActionDependencyGraphPreview("rejected", nil)
	if err != nil {
		t.Fatalf("rejected KDEActionDependencyGraphPreview returned error: %v", err)
	}
	if rejectedPreview.UserDecisionAllowsLaunch ||
		rejectedPreview.BlockedActionCount != 7 ||
		rejectedPreview.DependencyGraphPersisted ||
		rejectedPreview.ExecutionStarted {
		t.Fatalf("unexpected rejected graph preview: %#v", rejectedPreview)
	}

	if _, err := plan.KDEActionDependencyGraphPreview("invalid", nil); err == nil {
		t.Fatalf("KDEActionDependencyGraphPreview accepted an invalid decision")
	}
	if _, err := plan.KDEActionDependencyGraphPreview("approved\nbad", nil); err == nil {
		t.Fatalf("KDEActionDependencyGraphPreview accepted a multiline decision")
	}
	if _, err := plan.KDEActionDependencyGraphPreview("approved", []string{"https://example.invalid/book.xls"}); err == nil {
		t.Fatalf("KDEActionDependencyGraphPreview accepted a non-file URI")
	}

	preview, err := plan.KDEActionDependencyGraphPreview("reviewed", nil)
	if err != nil {
		t.Fatalf("KDEActionDependencyGraphPreview returned error: %v", err)
	}
	if err := validateNoBackendTerms(preview, "KDE action dependency graph preview test"); err != nil {
		t.Fatalf("validateNoBackendTerms returned error: %v", err)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("KDE action dependency graph preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func findKDEActionDependencyNode(nodes []KDEActionDependencyNode, id string) KDEActionDependencyNode {
	for _, node := range nodes {
		if node.ID == id {
			return node
		}
	}
	return KDEActionDependencyNode{}
}

func containsKDEActionDependencyEdge(edges []KDEActionDependencyEdge, from string, to string, relation string) bool {
	for _, edge := range edges {
		if edge.From == from && edge.To == to && edge.Relation == relation && edge.Blocking && edge.RuntimeOwned {
			return true
		}
	}
	return false
}
