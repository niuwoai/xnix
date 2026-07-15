package appidentity

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var runtimeMethodParityReadOnlyMethods = []string{
	"ListApplications",
	"GetApplication",
	"GetDiagnostics",
	"GetEngineCatalog",
	"GetRunPlan",
	"GetDesktopActivationManifest",
	"GetDesktopActivationTransactionPreview",
	"GetDesktopActivationStatus",
	"GetDesktopEntryPlan",
	"GetDesktopIconPlan",
	"GetTaskManagerIdentityPlan",
	"GetKDEIntegrationStatus",
	"GetKDEShellIntegrationPlan",
	"GetKDEApplicationSurfacePlan",
	"GetDesktopResourceBridgePlan",
	"GetKWinWindowRulePlan",
	"GetFileAssociationPlan",
	"GetNotificationPlan",
	"GetTrayStatus",
	"GetKRunnerQueryPlan",
	"GetPortalRequestPlan",
	"GetApplicationStateRoot",
	"GetCompatibilityPackageSource",
	"GetCompatibilityAcquisitionPreflight",
	"GetCompatibilityArtifactManifest",
	"GetCompatibilityInstallPlan",
	"GetBackendBinding",
	"GetBackendCapabilityMatrix",
	"GetBackendSelectionPlan",
	"GetBackendLifecycle",
	"GetBackendEnvironmentPlan",
	"GetRepairPlan",
	"GetTestPlan",
	"GetTestResult",
	"GetExecutionReadiness",
	"GetLaunchIntent",
	"GetAIDiagnosticInput",
	"GetAIDiagnosticRecommendation",
	"GetAIRepairApprovalGate",
	"GetSnapshotPlan",
	"GetPortalAccessPolicy",
	"GetRuntimeServiceBinding",
	"GetRuntimeLiveOwnerGate",
	"GetRuntimeOwnerSmokePlan",
	"GetRuntimeMethodParityManifest",
	"GetRuntimeWriteGate",
	"GetCompatibilitySettings",
	"GetCompatibilitySettingsChangePlan",
	"GetCompatibilityModeSwitchPlan",
	"GetCompatibilityPermissionReviewPlan",
	"GetCompatibilityReviewFlowPlan",
	"GetCompatibilityActionQueue",
	"GetCompatibilityActionReviewReceipt",
	"GetCompatibilityCenterSummary",
	"GetKDECenterPage",
	"GetKDECenterPageSections",
	"GetKDECenterPageSectionDetail",
}

var runtimeMethodParityWriteMethods = []string{
	"InstallRecipe",
	"Launch",
	"CreateSnapshot",
	"RestoreSnapshot",
}

var runtimeMethodParityClientMethods = map[string]string{
	"ListApplications":                       "list_applications",
	"GetApplication":                         "get_application",
	"GetDiagnostics":                         "diagnostics",
	"GetEngineCatalog":                       "engine_catalog",
	"GetRunPlan":                             "run_plan",
	"GetDesktopActivationManifest":           "desktop_activation_manifest",
	"GetDesktopActivationTransactionPreview": "desktop_activation_transaction_preview",
	"GetDesktopActivationStatus":             "desktop_activation_status",
	"GetDesktopEntryPlan":                    "desktop_entry_plan",
	"GetDesktopIconPlan":                     "desktop_icon_plan",
	"GetTaskManagerIdentityPlan":             "task_manager_identity_plan",
	"GetKDEIntegrationStatus":                "kde_integration_status",
	"GetKDEShellIntegrationPlan":             "kde_shell_integration_plan",
	"GetKDEApplicationSurfacePlan":           "kde_application_surface_plan",
	"GetDesktopResourceBridgePlan":           "desktop_resource_bridge_plan",
	"GetKWinWindowRulePlan":                  "kwin_window_rule_plan",
	"GetFileAssociationPlan":                 "file_association_plan",
	"GetNotificationPlan":                    "notification_plan",
	"GetTrayStatus":                          "tray_status",
	"GetKRunnerQueryPlan":                    "krunner_query_plan",
	"GetPortalRequestPlan":                   "portal_request_plan",
	"GetApplicationStateRoot":                "state_root",
	"GetCompatibilityPackageSource":          "package_source",
	"GetCompatibilityAcquisitionPreflight":   "acquisition_preflight",
	"GetCompatibilityArtifactManifest":       "artifact_manifest",
	"GetCompatibilityInstallPlan":            "install_plan",
	"GetBackendBinding":                      "backend_binding",
	"GetBackendCapabilityMatrix":             "backend_capability_matrix",
	"GetBackendSelectionPlan":                "backend_selection_plan",
	"GetBackendLifecycle":                    "backend_lifecycle",
	"GetBackendEnvironmentPlan":              "backend_environment_plan",
	"GetRepairPlan":                          "repair_plan",
	"GetTestPlan":                            "test_plan",
	"GetTestResult":                          "test_result",
	"GetExecutionReadiness":                  "execution_readiness",
	"GetLaunchIntent":                        "launch_intent",
	"GetAIDiagnosticInput":                   "ai_diagnostic_input",
	"GetAIDiagnosticRecommendation":          "ai_diagnostic_recommendation",
	"GetAIRepairApprovalGate":                "ai_repair_approval_gate",
	"GetSnapshotPlan":                        "snapshot_plan",
	"GetPortalAccessPolicy":                  "portal_access_policy",
	"GetRuntimeServiceBinding":               "runtime_service_binding",
	"GetRuntimeLiveOwnerGate":                "runtime_live_owner_gate",
	"GetRuntimeOwnerSmokePlan":               "runtime_owner_smoke_plan",
	"GetRuntimeMethodParityManifest":         "runtime_method_parity_manifest",
	"GetRuntimeWriteGate":                    "runtime_write_gate",
	"GetCompatibilitySettings":               "settings",
	"GetCompatibilitySettingsChangePlan":     "settings_change_plan",
	"GetCompatibilityModeSwitchPlan":         "compatibility_mode_switch_plan",
	"GetCompatibilityPermissionReviewPlan":   "compatibility_permission_review_plan",
	"GetCompatibilityReviewFlowPlan":         "compatibility_review_flow_plan",
	"GetCompatibilityActionQueue":            "action_queue",
	"GetCompatibilityActionReviewReceipt":    "action_review_receipt",
	"GetCompatibilityCenterSummary":          "compatibility_center_summary",
	"GetKDECenterPage":                       "kde_center_page",
	"GetKDECenterPageSections":               "kde_center_page_sections",
	"GetKDECenterPageSectionDetail":          "kde_center_page_section_detail",
}

type RuntimeMethodParityManifestPreview struct {
	Version                     string                     `json:"version"`
	SchemaVersion               string                     `json:"schema_version"`
	RequestType                 string                     `json:"request_type"`
	ManifestType                string                     `json:"manifest_type"`
	Source                      string                     `json:"source"`
	RuntimeMethod               string                     `json:"runtime_method"`
	ReadMethod                  string                     `json:"read_method"`
	BusName                     string                     `json:"bus_name"`
	ObjectPath                  string                     `json:"object_path"`
	Interface                   string                     `json:"interface"`
	ReadOnlyMethods             []string                   `json:"read_only_methods"`
	MethodCount                 int                        `json:"method_count"`
	ParityChecks                []RuntimeMethodParityCheck `json:"parity_checks"`
	CheckIDs                    []string                   `json:"check_ids"`
	Counts                      RuntimeMethodParityCounts  `json:"counts"`
	ReadOnlyMethodParityReady   bool                       `json:"read_only_method_parity_ready"`
	WriteMethods                []string                   `json:"write_methods"`
	WriteMethodsSupported       bool                       `json:"write_methods_supported"`
	WriteMethodDispatchEnabled  bool                       `json:"write_method_dispatch_enabled"`
	RuntimeOwned                bool                       `json:"runtime_owned"`
	GoRuntimeBacked             bool                       `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                       `json:"kde_policy_owner"`
	NetworkRequired             bool                       `json:"network_required"`
	HostRootModified            bool                       `json:"host_root_modified"`
	PrivilegedContainerRequired bool                       `json:"privileged_container_required"`
	BackendDetailsExposed       bool                       `json:"backend_details_exposed"`
	DesktopSafeSummary          string                     `json:"desktop_safe_summary"`
}

type RuntimeMethodParityCheck struct {
	ID             string   `json:"id"`
	Status         string   `json:"status"`
	MethodCount    int      `json:"method_count"`
	MissingMethods []string `json:"missing_methods"`
	Summary        string   `json:"summary"`
}

type RuntimeMethodParityCounts struct {
	Total   int `json:"total"`
	Passed  int `json:"passed"`
	Blocked int `json:"blocked"`
	Pending int `json:"pending"`
}

func NewRuntimeMethodParityManifestPreview(root string) (RuntimeMethodParityManifestPreview, error) {
	if root == "" {
		root = "."
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return RuntimeMethodParityManifestPreview{}, err
	}

	checks := []RuntimeMethodParityCheck{
		runtimeMethodParitySourceCheck("dbus-contract", readRuntimeMethodParitySources(root, []string{"runtime/dbus/org.xnix.Compatibility1.xml"}), func(source string, method string) bool {
			return strings.Contains(source, "name=\""+method+"\"")
		}),
		runtimeMethodParitySourceCheck("runtime-dispatch", readRuntimeMethodParitySources(root, []string{"lib/xnix/compatibility/runtime_daemon.rb"}), func(source string, method string) bool {
			return strings.Contains(source, "\""+method+"\"")
		}),
		runtimeMethodParitySourceCheck("dbus-client", readRuntimeMethodParitySources(root, []string{"lib/xnix/compatibility/dbus_runtime_client.rb"}), func(source string, method string) bool {
			clientMethod, ok := runtimeMethodParityClientMethods[method]
			return ok && strings.Contains(source, "def "+clientMethod)
		}),
		runtimeMethodParitySourceCheck("smoke-adapter", readRuntimeMethodParitySources(root, []string{
			"runtime/dbus/xnix_compatd_smoke.c",
			"runtime/dbus/xnix_compatd_introspection.inc",
			"runtime/dbus/xnix_compatd_kde_center.inc",
			"runtime/dbus/xnix_compatd_runtime_models.inc",
		}), func(source string, method string) bool {
			return strings.Contains(source, method)
		}),
		runtimeMethodParitySourceCheck("session-smoke", readRuntimeMethodParitySources(root, []string{"scripts/dbus_session_smoke.rb"}), func(source string, method string) bool {
			return strings.Contains(source, method)
		}),
	}
	counts := countRuntimeMethodParityChecks(checks)
	ready := runtimeMethodParityReady(checks)

	preview := RuntimeMethodParityManifestPreview{
		Version:                     version,
		SchemaVersion:               "xnix.runtime.method_parity_manifest.v1",
		RequestType:                 "runtime-method-parity-manifest-preview",
		ManifestType:                "runtime-method-parity-manifest",
		Source:                      "dbus-contract+runtime-dispatch+dbus-client+smoke-adapter+session-smoke",
		RuntimeMethod:               "GetRuntimeMethodParityManifest",
		ReadMethod:                  "GetRuntimeMethodParityManifestPreview",
		BusName:                     runtimeServiceBindingBusName,
		ObjectPath:                  runtimeServiceBindingObjectPath,
		Interface:                   runtimeServiceBindingInterface,
		ReadOnlyMethods:             append([]string{}, runtimeMethodParityReadOnlyMethods...),
		MethodCount:                 len(runtimeMethodParityReadOnlyMethods),
		ParityChecks:                checks,
		CheckIDs:                    runtimeMethodParityCheckIDs(checks),
		Counts:                      counts,
		ReadOnlyMethodParityReady:   ready,
		WriteMethods:                append([]string{}, runtimeMethodParityWriteMethods...),
		WriteMethodsSupported:       false,
		WriteMethodDispatchEnabled:  false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		DesktopSafeSummary:          runtimeMethodParitySummary(ready),
	}
	if err := validateNoBackendTerms(preview, "Runtime method parity manifest preview"); err != nil {
		return RuntimeMethodParityManifestPreview{}, err
	}
	return preview, nil
}

func runtimeMethodParitySourceCheck(id string, source string, covers func(string, string) bool) RuntimeMethodParityCheck {
	missing := make([]string, 0)
	for _, method := range runtimeMethodParityReadOnlyMethods {
		if !covers(source, method) {
			missing = append(missing, method)
		}
	}
	status := "pass"
	summary := fmt.Sprintf("%s covers all Runtime read-only methods.", id)
	if len(missing) > 0 {
		status = "blocked"
		summary = fmt.Sprintf("%s is missing Runtime read-only methods.", id)
	}
	return RuntimeMethodParityCheck{
		ID:             id,
		Status:         status,
		MethodCount:    len(runtimeMethodParityReadOnlyMethods) - len(missing),
		MissingMethods: missing,
		Summary:        summary,
	}
}

func readRuntimeMethodParitySources(root string, relativePaths []string) string {
	parts := make([]string, 0, len(relativePaths))
	for _, relativePath := range relativePaths {
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(relativePath)))
		if err != nil {
			parts = append(parts, "")
			continue
		}
		parts = append(parts, string(content))
	}
	return strings.Join(parts, "\n")
}

func countRuntimeMethodParityChecks(checks []RuntimeMethodParityCheck) RuntimeMethodParityCounts {
	counts := RuntimeMethodParityCounts{Total: len(checks)}
	for _, check := range checks {
		switch check.Status {
		case "pass":
			counts.Passed++
		case "blocked":
			counts.Blocked++
		case "pending":
			counts.Pending++
		}
	}
	return counts
}

func runtimeMethodParityCheckIDs(checks []RuntimeMethodParityCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func runtimeMethodParityReady(checks []RuntimeMethodParityCheck) bool {
	for _, check := range checks {
		if check.Status != "pass" {
			return false
		}
	}
	return true
}

func runtimeMethodParitySummary(ready bool) string {
	if ready {
		return "Runtime read-only D-Bus method parity is ready for owner smoke."
	}
	return "Runtime read-only D-Bus method parity must be repaired before production owner smoke."
}
