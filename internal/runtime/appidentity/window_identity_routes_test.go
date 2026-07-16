package appidentity

import (
	"encoding/json"
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
