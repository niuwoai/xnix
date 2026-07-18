package owner

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"xnix.local/xnix/internal/runtime/appidentity"
	"xnix.local/xnix/internal/runtime/image"
	"xnix.local/xnix/internal/runtime/recipe"
)

type ReadDispatch struct {
	Version                     string          `json:"version"`
	SchemaVersion               string          `json:"schema_version"`
	RequestType                 string          `json:"request_type"`
	DispatchType                string          `json:"dispatch_type"`
	Source                      string          `json:"source"`
	Method                      string          `json:"method"`
	Args                        []string        `json:"args"`
	RouteSource                 string          `json:"route_source"`
	GoCommand                   string          `json:"go_command"`
	RouteStatus                 string          `json:"route_status"`
	RouteReady                  bool            `json:"route_ready"`
	ReadOnlyDispatch            bool            `json:"read_only_dispatch"`
	WriteMethod                 bool            `json:"write_method"`
	WriteMethodsEnabled         bool            `json:"write_methods_enabled"`
	RuntimeOwned                bool            `json:"runtime_owned"`
	GoRuntimeBacked             bool            `json:"go_runtime_backed"`
	KDEPolicyOwner              bool            `json:"kde_policy_owner"`
	KDEMayClaimRuntimeOwnership bool            `json:"kde_may_claim_runtime_ownership"`
	EventLoopStarted            bool            `json:"event_loop_started"`
	SessionBusClaimed           bool            `json:"session_bus_claimed"`
	ProductionBusClaimed        bool            `json:"production_bus_claimed"`
	SystemServiceStarted        bool            `json:"system_service_started"`
	NetworkRequired             bool            `json:"network_required"`
	HostRootModified            bool            `json:"host_root_modified"`
	PrivilegedContainerRequired bool            `json:"privileged_container_required"`
	BackendDetailsExposed       bool            `json:"backend_details_exposed"`
	Payload                     json.RawMessage `json:"payload"`
	BlockedActions              []string        `json:"blocked_actions"`
	NextRequirements            []string        `json:"next_requirements"`
	DesktopSafeSummary          string          `json:"desktop_safe_summary"`
}

type ownerReadPayloadBuilder func(root string, args []string) (any, error)

const (
	defaultOwnerApplicationID = "org.xnix.sample.notepad"
	ownerDispatchReason       = "runtime-owner-dispatch"
	ownerDispatchDecision     = "reviewed"
)

var ownerReadDispatchMethodOrder = []string{
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
	"GetRuntimeOwnerProcess",
	"GetRuntimeOwnerRouteManifest",
	"GetRuntimeOwnerRecipeTrust",
	"GetRuntimeOwnerReadiness",
	"GetWindowsCompatibilityWorkstreamsPreview",
	"GetKDENotificationDigestPreview",
	"GetSignedRecipeVerificationPreview",
	"GetRestrictedProductSmokePacketPreview",
	"GetKDEOfflineApplicationIdentityPreview",
	"GetBackendAdapterProfileAudit",
	"GetRestrictedOwnerSmokeReceiptLookupPreview",
	"GetRestrictedOwnerSmokeReceiptFanOut",
	"GetKDETestLaunchMaterializationReceiptLookupPreview",
}

var ownerReadDispatchers = map[string]ownerReadPayloadBuilder{
	"ListApplications": func(root string, args []string) (any, error) {
		if err := requireArgCount("ListApplications", args, 0); err != nil {
			return nil, err
		}
		recipes, provenance, err := ownerRecipes(root)
		if err != nil {
			return nil, err
		}
		return appidentity.NewApplicationsPreview(recipes, provenance)
	},
	"GetApplication": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetApplication", args, 1); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewApplicationPreview(recipe, provenance)
	},
	"GetDiagnostics": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetDiagnostics", args, 1); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewDiagnosticsPreview(recipe, provenance)
	},
	"GetEngineCatalog": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetEngineCatalog", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewEngineCatalogPreview()
	},
	"GetRunPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetRunPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.RunPlanPreview()
	},
	"GetDesktopActivationManifest": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopActivationManifest", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.DesktopActivationManifestPreview()
	},
	"GetDesktopActivationTransactionPreview": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopActivationTransactionPreview", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.DesktopActivationTransactionPreview(args[1])
	},
	"GetDesktopActivationStatus": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopActivationStatus", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.DesktopActivationStatusPreview(args[1])
	},
	"GetDesktopEntryPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopEntryPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.RenderDesktopEntry()
	},
	"GetDesktopIconPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopIconPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.DesktopIconPreview()
	},
	"GetTaskManagerIdentityPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetTaskManagerIdentityPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.TaskManagerIdentityPlanPreview()
	},
	"GetKDEIntegrationStatus": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDEIntegrationStatus", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewKDEIntegrationStatusPreview()
	},
	"GetKDEShellIntegrationPlan": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDEShellIntegrationPlan", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewKDEShellIntegrationPlanPreview()
	},
	"GetKDEApplicationSurfacePlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetKDEApplicationSurfacePlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.KDEApplicationSurfacePlanPreview()
	},
	"GetDesktopResourceBridgePlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetDesktopResourceBridgePlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.DesktopResourceBridgePreview()
	},
	"GetKWinWindowRulePlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetKWinWindowRulePlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.KWinWindowRulePlanPreview()
	},
	"GetFileAssociationPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetFileAssociationPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.RenderMIMEApps()
	},
	"GetNotificationPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetNotificationPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.NotificationPreview(args[1])
	},
	"GetTrayStatus": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetTrayStatus", args, 0); err != nil {
			return nil, err
		}
		plan, err := ownerPlan(root, defaultOwnerApplicationID)
		if err != nil {
			return nil, err
		}
		return plan.TrayStatusPreview()
	},
	"GetKRunnerQueryPlan": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKRunnerQueryPlan", args, 1); err != nil {
			return nil, err
		}
		recipes, provenance, err := ownerRecipes(root)
		if err != nil {
			return nil, err
		}
		return appidentity.NewKRunnerQueryPreview(recipes, provenance, args[0])
	},
	"GetPortalRequestPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetPortalRequestPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.PortalRequestPreview(args[1], ownerDispatchReason)
	},
	"GetApplicationStateRoot": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetApplicationStateRoot", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.ApplicationStateRootPreview()
	},
	"GetCompatibilityPackageSource": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityPackageSource", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.PackageSourcePreview()
	},
	"GetCompatibilityAcquisitionPreflight": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityAcquisitionPreflight", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.AcquisitionPreflightPreview()
	},
	"GetCompatibilityArtifactManifest": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityArtifactManifest", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.ArtifactManifestPreview()
	},
	"GetCompatibilityInstallPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityInstallPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.CompatibilityInstallPlanPreview(args[1])
	},
	"GetBackendBinding": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetBackendBinding", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.BackendBindingPreview()
	},
	"GetBackendCapabilityMatrix": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetBackendCapabilityMatrix", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewBackendCapabilityMatrixPreview()
	},
	"GetBackendSelectionPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetBackendSelectionPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.BackendSelectionPreview()
	},
	"GetBackendLifecycle": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetBackendLifecycle", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.BackendLifecyclePreview()
	},
	"GetBackendEnvironmentPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetBackendEnvironmentPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.BackendEnvironmentPreview()
	},
	"GetRepairPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetRepairPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.RepairPlanPreview(args[1])
	},
	"GetTestPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetTestPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.TestPlanPreview(args[1])
	},
	"GetTestResult": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetTestResult", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.TestResultPreview(args[1])
	},
	"GetExecutionReadiness": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetExecutionReadiness", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.ExecutionReadinessPreview()
	},
	"GetLaunchIntent": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetLaunchIntent", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.LaunchIntentPreview(nil)
	},
	"GetAIDiagnosticInput": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetAIDiagnosticInput", args, 3)
		if err != nil {
			return nil, err
		}
		return plan.AIDiagnosticInputPreview(args[1], args[2])
	},
	"GetAIDiagnosticRecommendation": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetAIDiagnosticRecommendation", args, 3)
		if err != nil {
			return nil, err
		}
		return plan.AIDiagnosticRecommendationPreview(args[1], args[2])
	},
	"GetAIRepairApprovalGate": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetAIRepairApprovalGate", args, 3)
		if err != nil {
			return nil, err
		}
		return plan.AIRepairApprovalGatePreview(args[1], args[2])
	},
	"GetSnapshotPlan": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetSnapshotPlan", args, 2); err != nil {
			return nil, err
		}
		return appidentity.NewSnapshotPlanPreview(args[0], args[1])
	},
	"GetPortalAccessPolicy": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetPortalAccessPolicy", args, 2); err != nil {
			return nil, err
		}
		return appidentity.NewPortalAccessPolicyPreview(args[0], args[1])
	},
	"GetRuntimeServiceBinding": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeServiceBinding", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeServiceBindingPreview(root)
	},
	"GetRuntimeLiveOwnerGate": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeLiveOwnerGate", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeLiveOwnerGatePreview(root)
	},
	"GetRuntimeOwnerProcess": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerProcess", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerProcessPreview(root)
	},
	"GetRuntimeOwnerSmokePlan": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerSmokePlan", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerSmokePlanPreview(root)
	},
	"GetRuntimeMethodParityManifest": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeMethodParityManifest", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeMethodParityManifestPreview(root)
	},
	"GetRuntimeOwnerRouteManifest": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerRouteManifest", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	},
	"GetRuntimeOwnerRecipeTrust": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerRecipeTrust", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerRecipeTrustPreview(root)
	},
	"GetRuntimeOwnerReadiness": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeOwnerReadiness", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeOwnerReadinessPreview(root)
	},
	"GetWindowsCompatibilityWorkstreamsPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetWindowsCompatibilityWorkstreamsPreview", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewWindowsCompatibilityWorkstreamsPreview()
	},
	"GetKDENotificationDigestPreview": func(root string, args []string) (any, error) {
		if len(args) < 2 {
			return nil, fmt.Errorf("GetKDENotificationDigestPreview requires an application id and at least one event")
		}
		plan, err := ownerPlan(root, args[0])
		if err != nil {
			return nil, err
		}
		events, err := ownerNotificationDigestEvents(args[1:])
		if err != nil {
			return nil, err
		}
		return plan.KDENotificationDigestPreview(events)
	},
	"GetSignedRecipeVerificationPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetSignedRecipeVerificationPreview", args, 1); err != nil {
			return nil, err
		}
		store, err := recipe.NewLocalStore(filepath.Join(root, "runtime", "recipes"), nil)
		if err != nil {
			return nil, err
		}
		loaded, err := store.Find(args[0])
		if err != nil {
			return nil, err
		}
		return recipe.NewRegistrySignedRecipeVerificationPreview(loaded), nil
	},
	"GetRestrictedProductSmokePacketPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRestrictedProductSmokePacketPreview", args, 0); err != nil {
			return nil, err
		}
		return image.PrepareRestrictedProductSmokePacket(root, "image/kinoite/manifest.json")
	},
	"GetKDEOfflineApplicationIdentityPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDEOfflineApplicationIdentityPreview", args, 1); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewKDEOfflineApplicationIdentityPreview(recipe, provenance)
	},
	"GetBackendAdapterProfileAudit": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetBackendAdapterProfileAudit", args, 0); err != nil {
			return nil, err
		}
		return appidentity.NewBackendAdapterRedactedProfileAuditPreview(root)
	},
	"GetRestrictedOwnerSmokeReceiptLookupPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRestrictedOwnerSmokeReceiptLookupPreview", args, 1); err != nil {
			return nil, err
		}
		return ResolveRestrictedOwnerSmokeReceipt(root, args[0])
	},
	"GetRestrictedOwnerSmokeReceiptFanOut": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRestrictedOwnerSmokeReceiptFanOut", args, 1); err != nil {
			return nil, err
		}
		return NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(root, args[0])
	},
	"GetKDETestLaunchMaterializationReceiptLookupPreview": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDETestLaunchMaterializationReceiptLookupPreview", args, 1); err != nil {
			return nil, err
		}
		return appidentity.ResolveKDETestLaunchMaterializationReceipt(root, args[0])
	},
	"GetRuntimeWriteGate": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetRuntimeWriteGate", args, 1); err != nil {
			return nil, err
		}
		return appidentity.NewRuntimeWriteGatePreview(root, args[0])
	},
	"GetCompatibilitySettings": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilitySettings", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.SettingsPreview()
	},
	"GetCompatibilitySettingsChangePlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilitySettingsChangePlan", args, 4)
		if err != nil {
			return nil, err
		}
		return plan.SettingsChangePreview(args[1], args[2], args[3])
	},
	"GetCompatibilityModeSwitchPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityModeSwitchPlan", args, 2)
		if err != nil {
			return nil, err
		}
		return plan.ModeSwitchPreview(args[1])
	},
	"GetCompatibilityPermissionReviewPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityPermissionReviewPlan", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.PermissionReviewPreview()
	},
	"GetCompatibilityReviewFlowPlan": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityReviewFlowPlan", args, 5)
		if err != nil {
			return nil, err
		}
		return plan.ReviewFlowPreview(args[1], args[2], args[3], args[4])
	},
	"GetCompatibilityActionQueue": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityActionQueue", args, 1)
		if err != nil {
			return nil, err
		}
		return plan.KDEActionQueuePreview(ownerDispatchDecision, nil)
	},
	"GetCompatibilityActionReviewReceipt": func(root string, args []string) (any, error) {
		plan, err := ownerPlanFromArgs(root, "GetCompatibilityActionReviewReceipt", args, 3)
		if err != nil {
			return nil, err
		}
		return plan.KDEActionReceiptPreview(args[1], args[2], nil)
	},
	"GetCompatibilityCenterSummary": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetCompatibilityCenterSummary", args, 1); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewCompatibilityCenterPreview([]appidentity.Recipe{recipe}, provenance)
	},
	"GetKDECenterPage": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDECenterPage", args, 2); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewKDECenterPagePreview(recipe, provenance, args[1], nil)
	},
	"GetKDECenterPageSections": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDECenterPageSections", args, 2); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewKDECenterPageSectionsPreview(recipe, provenance, args[1], nil)
	},
	"GetKDECenterPageSectionDetail": func(root string, args []string) (any, error) {
		if err := requireArgCount("GetKDECenterPageSectionDetail", args, 3); err != nil {
			return nil, err
		}
		recipe, provenance, err := ownerRecipe(root, args[0])
		if err != nil {
			return nil, err
		}
		return appidentity.NewKDECenterPageSectionDetailPreview(recipe, provenance, args[1], args[2], nil)
	},
}

func DispatchRead(root string, method string, args []string) (ReadDispatch, error) {
	method = strings.TrimSpace(method)
	normalizedArgs := append([]string{}, args...)
	if method == "" {
		return ReadDispatch{}, fmt.Errorf("read method is required")
	}
	if isReservedWriteMethod(method) {
		return ReadDispatch{}, fmt.Errorf("%s is a write method; use disabled write-method response", method)
	}
	builder, ok := ownerReadDispatchers[method]
	if !ok {
		return ReadDispatch{}, fmt.Errorf("unsupported owner read dispatch method: %s", method)
	}

	routeManifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return ReadDispatch{}, err
	}
	route := ownerReadRoute(routeManifest.Routes, method)
	payloadValue, err := builder(root, normalizedArgs)
	if err != nil {
		return ReadDispatch{}, err
	}
	payload, err := json.Marshal(payloadValue)
	if err != nil {
		return ReadDispatch{}, fmt.Errorf("encode %s dispatch payload: %w", method, err)
	}

	dispatch := ReadDispatch{
		Version:                     routeManifest.Version,
		SchemaVersion:               "xnix.runtime.owner_read_dispatch.v1",
		RequestType:                 "runtime-owner-read-dispatch",
		DispatchType:                "go-owner-read-dispatch",
		Source:                      "go-runtime-owner-candidate+in-process-read-dispatch",
		Method:                      method,
		Args:                        normalizedArgs,
		RouteSource:                 route.CurrentSource,
		GoCommand:                   route.GoCommand,
		RouteStatus:                 route.RouteStatus,
		RouteReady:                  route.GoRouteReady,
		ReadOnlyDispatch:            true,
		WriteMethod:                 false,
		WriteMethodsEnabled:         false,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		KDEMayClaimRuntimeOwnership: false,
		EventLoopStarted:            false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		SystemServiceStarted:        false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		Payload:                     payload,
		BlockedActions:              readDispatchBlockedActions(),
		NextRequirements:            readDispatchNextRequirements(),
		DesktopSafeSummary:          "Runtime owner read dispatch can render selected read-only owner methods without claiming D-Bus ownership.",
	}
	if err := validateNoBackendTerms(dispatch, "Runtime owner read dispatch"); err != nil {
		return ReadDispatch{}, err
	}
	return dispatch, nil
}

func SupportedReadDispatchMethods() []string {
	return append([]string(nil), ownerReadDispatchMethodOrder...)
}

func requireArgCount(method string, args []string, expected int) error {
	if len(args) != expected {
		return fmt.Errorf("%s requires %d argument(s), got %d", method, expected, len(args))
	}
	return nil
}

func ownerDefaultRegistry(root string) string {
	return filepath.Join(root, "runtime", "recipes", "registry.json")
}

func ownerRecipes(root string) ([]appidentity.Recipe, appidentity.Provenance, error) {
	return appidentity.LoadRecipesFromRegistry(ownerDefaultRegistry(root), "")
}

func ownerRecipe(root string, applicationID string) (appidentity.Recipe, appidentity.Provenance, error) {
	return appidentity.LoadRecipeFromRegistry(ownerDefaultRegistry(root), "", applicationID)
}

func ownerPlan(root string, applicationID string) (appidentity.Plan, error) {
	recipe, provenance, err := ownerRecipe(root, applicationID)
	if err != nil {
		return appidentity.Plan{}, err
	}
	return appidentity.NewPlanWithProvenance(recipe, provenance)
}

func ownerPlanFromArgs(root string, method string, args []string, expected int) (appidentity.Plan, error) {
	if err := requireArgCount(method, args, expected); err != nil {
		return appidentity.Plan{}, err
	}
	return ownerPlan(root, args[0])
}

func ownerReadRoute(routes []appidentity.RuntimeOwnerRoute, method string) appidentity.RuntimeOwnerRoute {
	for _, route := range routes {
		if route.Method == method {
			return route
		}
	}
	return appidentity.RuntimeOwnerRoute{
		Method:        method,
		CurrentSource: "go-owner-local-preview",
		TargetSource:  "go-runtime-owner",
		GoCommand:     ownerReadDispatchCommand(method),
		RouteStatus:   "owner-local-preview-ready",
		GoRouteReady:  true,
	}
}

func ownerReadDispatchCommand(method string) string {
	switch method {
	case "GetRuntimeServiceBinding":
		return "runtime-service-binding-preview"
	case "GetRuntimeLiveOwnerGate":
		return "runtime-live-owner-gate-preview"
	case "GetRuntimeOwnerProcess":
		return "runtime-owner-process-preview"
	case "GetRuntimeOwnerSmokePlan":
		return "runtime-owner-smoke-plan-preview"
	case "GetRuntimeMethodParityManifest":
		return "runtime-method-parity-manifest-preview"
	case "GetRuntimeOwnerRouteManifest":
		return "runtime-owner-route-manifest-preview"
	case "GetRuntimeOwnerRecipeTrust":
		return "runtime-owner-recipe-trust-preview"
	case "GetRuntimeOwnerReadiness":
		return "runtime-owner-readiness-preview"
	case "GetWindowsCompatibilityWorkstreamsPreview":
		return "windows-compatibility-workstreams-preview"
	case "GetKDENotificationDigestPreview":
		return "kde-notification-digest-preview"
	case "GetSignedRecipeVerificationPreview":
		return "signed-recipe-verifier-preview"
	case "GetRestrictedProductSmokePacketPreview":
		return "restricted-product-smoke-packet-preview"
	case "GetKDEOfflineApplicationIdentityPreview":
		return "kde-offline-application-identity-preview"
	case "GetBackendAdapterProfileAudit":
		return "backend-adapter-redacted-profile-audit-preview"
	case "GetRestrictedOwnerSmokeReceiptLookupPreview":
		return "restricted-owner-smoke-receipt-lookup-preview"
	case "GetRestrictedOwnerSmokeReceiptFanOut":
		return "restricted-owner-smoke-receipt-fanout-owner-route-preview"
	case "GetKDETestLaunchMaterializationReceiptLookupPreview":
		return "kde-test-launch-materialization-receipt-lookup-preview"
	case "GetRuntimeWriteGate":
		return "runtime-write-gate-preview"
	default:
		return "unsupported-owner-read-preview"
	}
}

func ownerNotificationDigestEvents(values []string) ([]appidentity.KDENotificationDigestEvent, error) {
	events := make([]appidentity.KDENotificationDigestEvent, 0, len(values))
	for _, value := range values {
		group, eventID, ok := strings.Cut(value, ":")
		if !ok || strings.TrimSpace(group) == "" || strings.TrimSpace(eventID) == "" {
			return nil, fmt.Errorf("notification digest event must use <group>:<event-id>")
		}
		events = append(events, appidentity.KDENotificationDigestEvent{Group: group, EventID: eventID})
	}
	return events, nil
}

func isReservedWriteMethod(method string) bool {
	for _, writeMethod := range []string{"InstallRecipe", "Launch", "CreateSnapshot", "RestoreSnapshot"} {
		if method == writeMethod {
			return true
		}
	}
	return false
}

func readDispatchBlockedActions() []string {
	return []string{
		"claim a D-Bus name from read dispatch preview",
		"start a system service from read dispatch preview",
		"enable write methods from read dispatch preview",
		"let KDE route Runtime policy directly",
		"mutate host root during read dispatch preview",
		"expose backend implementation details through read dispatch output",
	}
}

func readDispatchNextRequirements() []string {
	return []string{
		"Bind read dispatch to the restricted owner event loop.",
		"Map every read-only Runtime method to an in-process owner handler.",
		"Keep write methods disabled until production owner readiness is proven.",
		"Promote the packaged entrypoint only after D-Bus read dispatch parity passes.",
	}
}
