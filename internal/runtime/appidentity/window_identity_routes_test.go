package appidentity

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTaskManagerIdentityPlanPreviewUsesWindowIdentityWithoutActivating(t *testing.T) {
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

	preview, err := plan.TaskManagerIdentityPlanPreview()
	if err != nil {
		t.Fatalf("TaskManagerIdentityPlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.task_manager_identity.v1" ||
		preview.RequestType != "task-manager-identity-preview" ||
		preview.PlanType != "task-manager-identity-plan" ||
		preview.Source != "window-identity-preview" ||
		preview.RuntimeMethod != "GetTaskManagerIdentityPlan" ||
		preview.ReadMethod != "GetTaskManagerIdentityPlanPreview" ||
		preview.Desktop != "KDE Plasma" {
		t.Fatalf("unexpected task manager identity schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.WindowKind != "compatibility-application" ||
		preview.GroupingKey != "org.example.ledger" {
		t.Fatalf("unexpected task manager identity fields: %#v", preview)
	}
	if !preview.TaskManager.PinningAllowed ||
		!preview.TaskManager.RestoreAllowed ||
		preview.TaskManager.SkipTaskbar ||
		!preview.TaskManager.ShowInSwitcher ||
		!preview.TaskManager.PreferExistingWindow ||
		preview.Restore.RestoreKey != "org.example.ledger" {
		t.Fatalf("unexpected task manager hints: %#v restore=%#v", preview.TaskManager, preview.Restore)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.PinningAllowed ||
		!preview.RestoreAllowed ||
		!preview.PreferExistingWindow ||
		preview.SkipTaskbar ||
		!preview.ShowInSwitcher ||
		preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected task manager safety flags: %#v", preview)
	}
	assertNoWindowRouteBackendTerms(t, preview)
}

func TestTaskManagerIdentityPlanPreviewConsumesActivationReceipt(t *testing.T) {
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

	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("TaskManagerIdentityPlanPreviewWithOptions returned error: %v", err)
	}
	if preview.Source != "window-identity-preview+desktop-activation-receipt" ||
		!preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("task manager identity did not consume activation receipt: %#v", preview)
	}
	if preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed task manager identity must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("task manager identity exposed activation root: %s", encoded)
	}

	if _, err := plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("task manager identity accepted a missing activation receipt")
	}
}

func TestTaskManagerIdentityPlanPreviewConsumesExecutionSessionRecord(t *testing.T) {
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
	root := t.TempDir()
	writeExecutionSessionRecord(t, root, "xnix-exec-org-example-ledger-1", "org.example.ledger")

	preview, err := plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{
		ExecutionSessionRoot:      root,
		ExecutionSessionRequestID: "xnix-exec-org-example-ledger-1",
	})
	if err != nil {
		t.Fatalf("TaskManagerIdentityPlanPreviewWithOptions returned error: %v", err)
	}
	if preview.Source != "window-identity-preview+execution-session-record" ||
		!preview.ExecutionSessionRoot ||
		!preview.ExecutionSessionBacked ||
		preview.ExecutionSessionPath != "execution-ledger/sessions/xnix-exec-org-example-ledger-1.json" ||
		preview.ExecutionSessionState != "blocked" {
		t.Fatalf("task manager identity did not consume execution session record: %#v", preview)
	}
	if preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("session-backed task manager identity must remain gated: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("task manager identity exposed session root: %s", encoded)
	}
	if _, err := plan.TaskManagerIdentityPlanPreviewWithOptions(TaskManagerIdentityOptions{ExecutionSessionRoot: root, ExecutionSessionRequestID: "missing"}); err == nil {
		t.Fatalf("task manager identity accepted missing execution session record")
	}
}

func TestExecutionSessionFanOutEvidenceSharesOneRecordAcrossKDESurfaces(t *testing.T) {
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
	root := t.TempDir()
	requestID := "xnix-exec-org-example-ledger-1"
	writeExecutionSessionRecord(t, root, requestID, "org.example.ledger")

	fanOut, err := plan.ExecutionSessionFanOutEvidence(root, requestID)
	if err != nil {
		t.Fatalf("ExecutionSessionFanOutEvidence returned error: %v", err)
	}

	if fanOut.SchemaVersion != "xnix.runtime.session_fanout.v1" ||
		fanOut.RequestType != "execution-session-fanout-evidence" ||
		fanOut.Source != "execution-session-record" ||
		fanOut.RuntimeMethod != "GetExecutionSessionFanOutEvidence" ||
		fanOut.ReadMethod != "GetExecutionSessionFanOutEvidence" ||
		fanOut.ReceiptRelativePath != "execution-ledger/sessions/"+requestID+".json" ||
		fanOut.RequestID != requestID ||
		fanOut.SessionState != "blocked" ||
		fanOut.SurfaceCount != 4 ||
		len(fanOut.Surfaces) != 4 {
		t.Fatalf("unexpected fan-out identity: %#v", fanOut)
	}
	if fanOut.TaskManager.ID != "task-manager" ||
		fanOut.TaskManager.State != "blocked" ||
		fanOut.KWin.ID != "kwin" ||
		fanOut.KWin.State != "blocked" ||
		fanOut.Tray.ID != "tray" ||
		fanOut.Tray.State != "blocked" ||
		fanOut.CompatibilityCenter.ID != "compatibility-center" ||
		fanOut.CompatibilityCenter.State != "waiting-for-runtime-gates" {
		t.Fatalf("unexpected fan-out surface states: %#v", fanOut)
	}
	for _, surface := range fanOut.Surfaces {
		if surface.ReceiptRelativePath != fanOut.ReceiptRelativePath ||
			!surface.NavigationOnly ||
			!surface.ReadOnly ||
			surface.MutatesRuntime ||
			surface.StartsProgram ||
			!surface.SafeForKDE {
			t.Fatalf("unexpected fan-out surface safety: %#v", surface)
		}
	}
	if !fanOut.SafeForKDE ||
		!fanOut.RuntimeOwned ||
		!fanOut.GoRuntimeBacked ||
		fanOut.KDEPolicyOwner ||
		fanOut.StateRootPathExposed ||
		fanOut.LiveStateObserved ||
		fanOut.WindowObserved ||
		fanOut.TaskManagerEntryActive ||
		fanOut.KWinRuleApplied ||
		fanOut.LiveTrayBridgeEnabled ||
		fanOut.LaunchEnabled ||
		fanOut.ExecutionStarted ||
		fanOut.BackendProcessStarted ||
		fanOut.PermissionGranted ||
		fanOut.HostRootModified ||
		fanOut.NetworkRequired ||
		fanOut.BackendDetailsExposed {
		t.Fatalf("fan-out evidence opened an unsafe gate: %#v", fanOut)
	}
	assertNoWindowRouteBackendTerms(t, fanOut)
}

func TestKWinWindowRulePlanPreviewUsesWindowIdentityWithoutApplyingRule(t *testing.T) {
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

	preview, err := plan.KWinWindowRulePlanPreview()
	if err != nil {
		t.Fatalf("KWinWindowRulePlanPreview returned error: %v", err)
	}
	if preview.SchemaVersion != "xnix.runtime.kwin_window_rule.v1" ||
		preview.RequestType != "kwin-window-rule-preview" ||
		preview.PlanType != "kwin-window-rule-plan" ||
		preview.Source != "window-identity-preview" ||
		preview.RuntimeMethod != "GetKWinWindowRulePlan" ||
		preview.ReadMethod != "GetKWinWindowRulePlanPreview" ||
		preview.Desktop != "KDE Plasma" {
		t.Fatalf("unexpected KWin rule schema: %#v", preview)
	}
	if preview.ApplicationID != "org.example.ledger" ||
		preview.DisplayName != "Example Ledger" ||
		preview.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.WindowKind != "compatibility-application" {
		t.Fatalf("unexpected KWin rule identity: %#v", preview)
	}
	if preview.Match.ResourceName != "org.example.ledger" ||
		preview.Match.ClassGroup != "xnix-compatibility" ||
		preview.Match.TitleHint != "Example Ledger" ||
		preview.Set.DesktopFile != "xnix-org.example.ledger.desktop" ||
		preview.Set.ApplicationID != "org.example.ledger" ||
		preview.Set.TaskManagerGroupingKey != "org.example.ledger" ||
		preview.Set.LauncherURL != "applications:xnix-org.example.ledger.desktop" ||
		preview.Set.SkipTaskbar ||
		!preview.Set.ShowInSwitcher ||
		preview.Set.Placement != "normal-window" {
		t.Fatalf("unexpected KWin match/set hints: match=%#v set=%#v", preview.Match, preview.Set)
	}
	if !preview.RuntimeOwned ||
		!preview.GoRuntimeBacked ||
		preview.KDEPolicyOwner ||
		!preview.WindowManagerPolicyOnly ||
		!preview.RuntimeOwnsBackendPolicy ||
		preview.KWinRuleApplied ||
		preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("unexpected KWin safety flags: %#v", preview)
	}
	assertNoWindowRouteBackendTerms(t, preview)
}

func TestKWinWindowRulePlanPreviewConsumesActivationReceipt(t *testing.T) {
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

	root := t.TempDir()
	writeActivationReceipt(t, root, "org.example.ledger")

	preview, err := plan.KWinWindowRulePlanPreviewWithOptions(KWinWindowRuleOptions{ActivationRoot: root})
	if err != nil {
		t.Fatalf("KWinWindowRulePlanPreviewWithOptions returned error: %v", err)
	}
	if preview.Source != "window-identity-preview+desktop-activation-receipt" ||
		!preview.ActivationReceiptRoot ||
		!preview.ActivationReceiptBacked ||
		preview.ActivationReceiptPath != "usr/share/xnix/compatibility/activation-receipts/org.example.ledger.json" {
		t.Fatalf("KWin window rule did not consume activation receipt: %#v", preview)
	}
	if preview.KWinRuleApplied ||
		preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("receipt-backed KWin window rule must remain gated: %#v", preview)
	}

	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("KWin window rule exposed activation root: %s", encoded)
	}

	if _, err := plan.KWinWindowRulePlanPreviewWithOptions(KWinWindowRuleOptions{ActivationRoot: t.TempDir()}); err == nil {
		t.Fatalf("KWin window rule accepted a missing activation receipt")
	}
}

func TestKWinWindowRulePlanPreviewConsumesExecutionSessionRecord(t *testing.T) {
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
	root := t.TempDir()
	writeExecutionSessionRecord(t, root, "xnix-exec-org-example-ledger-1", "org.example.ledger")

	preview, err := plan.KWinWindowRulePlanPreviewWithOptions(KWinWindowRuleOptions{
		ExecutionSessionRoot:      root,
		ExecutionSessionRequestID: "xnix-exec-org-example-ledger-1",
	})
	if err != nil {
		t.Fatalf("KWinWindowRulePlanPreviewWithOptions returned error: %v", err)
	}
	if preview.Source != "window-identity-preview+execution-session-record" ||
		!preview.ExecutionSessionRoot ||
		!preview.ExecutionSessionBacked ||
		preview.ExecutionSessionPath != "execution-ledger/sessions/xnix-exec-org-example-ledger-1.json" ||
		preview.ExecutionSessionState != "blocked" {
		t.Fatalf("KWin rule did not consume execution session record: %#v", preview)
	}
	if preview.KWinRuleApplied ||
		preview.TaskManagerEntryActive ||
		preview.WindowObservationStarted ||
		preview.LaunchEnabled ||
		preview.ExecutionStarted ||
		preview.HostRootModified ||
		preview.BackendDetailsExposed {
		t.Fatalf("session-backed KWin rule must remain gated: %#v", preview)
	}
	encoded, err := json.Marshal(preview)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	if strings.Contains(string(encoded), root) {
		t.Fatalf("KWin rule exposed session root: %s", encoded)
	}
}

func assertNoWindowRouteBackendTerms(t *testing.T, value any) {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	text := strings.ToLower(string(encoded))
	for _, forbidden := range []string{"prefix", ".exe", "program files", "qemu-system", "proton", "wine ", "virtual machine"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("window route preview exposes forbidden term %q: %s", forbidden, text)
		}
	}
}

func writeExecutionSessionRecord(t *testing.T, root string, requestID string, applicationID string) {
	t.Helper()
	dir := filepath.Join(root, "execution-ledger", "sessions")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatalf("MkdirAll returned error: %v", err)
	}
	relativePath := "execution-ledger/sessions/" + requestID + ".json"
	data := []byte(`{
  "schema_version": "xnix.runtime.execution_session_record.v1",
  "record_type": "execution-session-status-record",
  "request_id": "` + requestID + `",
  "application_id": "` + applicationID + `",
  "relative_path": "` + relativePath + `",
  "session_state": "blocked",
  "task_manager_state": "blocked",
  "tray_state": "blocked",
  "kwin_state": "blocked",
  "compatibility_center_state": "waiting-for-runtime-gates",
  "state_root_path_exposed": false,
  "status_persisted": true,
  "session_active": false,
  "live_state_observed": false,
  "window_observed": false,
  "task_manager_entry_active": false,
  "kwin_rule_applied": false,
  "live_tray_bridge_enabled": false,
  "launch_enabled": false,
  "execution_started": false,
  "backend_process_started": false,
  "permission_granted": false,
  "host_root_modified": false,
  "network_required": false,
  "backend_details_exposed": false
}`)
	if err := os.WriteFile(filepath.Join(dir, requestID+".json"), data, 0o600); err != nil {
		t.Fatalf("WriteFile session record returned error: %v", err)
	}
}
