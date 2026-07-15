package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/appidentity"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: xnix-runtime-go {application-preview|applications-preview|backend-binding-preview|backend-environment-preview|backend-selection-preview|compatibility-center-preview|desktop-activation-bundle-preview|desktop-activation-preflight-preview|desktop-activation-staging-preview|desktop-activation-status-preview|desktop-activation-transaction-preview|desktop-entry-preview|desktop-icon-preview|desktop-identity-plan|desktop-resource-bridge-preview|diagnostics-preview|dolphin-ai-analysis-preview|dolphin-drop-preview|execution-decision-preview|execution-preflight-preview|execution-readiness-preview|execution-request-preview|execution-resource-grant-preview|execution-review-preview|execution-session-preview|execution-session-status-preview|execution-transaction-preview|file-open-preview|kde-action-card-deck-preview|kde-action-card-preview|kde-action-preflight-preview|kde-action-queue-preview|kde-action-receipt-preview|kde-action-review-preview|kde-action-status-preview|kde-center-page-preview|kde-center-page-section-detail-preview|kde-center-page-sections-preview|kde-entrypoint-action-preview|kde-entrypoints-preview|krunner-query-preview|launch-intent-preview|mimeapps-preview|mode-switch-preview|notification-preview|permission-review-preview|portal-request-preview|review-flow-preview|runtime-live-owner-gate-preview|runtime-method-parity-manifest-preview|runtime-owner-process-preview|runtime-owner-readiness-preview|runtime-owner-recipe-trust-preview|runtime-owner-route-manifest-preview|runtime-owner-smoke-plan-preview|runtime-service-binding-preview|settings-change-preview|settings-preview|tray-status-preview|window-identity-preview}")
	}

	switch args[0] {
	case "application-preview":
		return runApplicationPreview(args[1:], stdout)
	case "applications-preview":
		return runApplicationsPreview(args[1:], stdout)
	case "backend-binding-preview":
		return runBackendBindingPreview(args[1:], stdout)
	case "backend-environment-preview":
		return runBackendEnvironmentPreview(args[1:], stdout)
	case "backend-selection-preview":
		return runBackendSelectionPreview(args[1:], stdout)
	case "compatibility-center-preview":
		return runCompatibilityCenterPreview(args[1:], stdout)
	case "desktop-activation-bundle-preview":
		return runDesktopActivationBundlePreview(args[1:], stdout)
	case "desktop-activation-preflight-preview":
		return runDesktopActivationPreflightPreview(args[1:], stdout)
	case "desktop-activation-staging-preview":
		return runDesktopActivationStagingPreview(args[1:], stdout)
	case "desktop-activation-status-preview":
		return runDesktopActivationStatusPreview(args[1:], stdout)
	case "desktop-activation-transaction-preview":
		return runDesktopActivationTransactionPreview(args[1:], stdout)
	case "desktop-entry-preview":
		return runDesktopEntryPreview(args[1:], stdout)
	case "desktop-icon-preview":
		return runDesktopIconPreview(args[1:], stdout)
	case "desktop-identity-plan":
		return runDesktopIdentityPlan(args[1:], stdout)
	case "desktop-resource-bridge-preview":
		return runDesktopResourceBridgePreview(args[1:], stdout)
	case "diagnostics-preview":
		return runDiagnosticsPreview(args[1:], stdout)
	case "dolphin-ai-analysis-preview":
		return runDolphinAIAnalysisPreview(args[1:], stdout)
	case "dolphin-drop-preview":
		return runDolphinDropPreview(args[1:], stdout)
	case "execution-decision-preview":
		return runExecutionDecisionPreview(args[1:], stdout)
	case "execution-preflight-preview":
		return runExecutionPreflightPreview(args[1:], stdout)
	case "execution-readiness-preview":
		return runExecutionReadinessPreview(args[1:], stdout)
	case "execution-request-preview":
		return runExecutionRequestPreview(args[1:], stdout)
	case "execution-resource-grant-preview":
		return runExecutionResourceGrantPreview(args[1:], stdout)
	case "execution-review-preview":
		return runExecutionReviewPreview(args[1:], stdout)
	case "execution-session-preview":
		return runExecutionSessionPreview(args[1:], stdout)
	case "execution-session-status-preview":
		return runExecutionSessionStatusPreview(args[1:], stdout)
	case "execution-transaction-preview":
		return runExecutionTransactionPreview(args[1:], stdout)
	case "file-open-preview":
		return runFileOpenPreview(args[1:], stdout)
	case "kde-action-card-deck-preview":
		return runKDEActionCardDeckPreview(args[1:], stdout)
	case "kde-action-card-preview":
		return runKDEActionCardPreview(args[1:], stdout)
	case "kde-action-queue-preview":
		return runKDEActionQueuePreview(args[1:], stdout)
	case "kde-action-preflight-preview":
		return runKDEActionPreflightPreview(args[1:], stdout)
	case "kde-action-receipt-preview":
		return runKDEActionReceiptPreview(args[1:], stdout)
	case "kde-action-review-preview":
		return runKDEActionReviewPreview(args[1:], stdout)
	case "kde-action-status-preview":
		return runKDEActionStatusPreview(args[1:], stdout)
	case "kde-center-page-preview":
		return runKDECenterPagePreview(args[1:], stdout)
	case "kde-center-page-section-detail-preview":
		return runKDECenterPageSectionDetailPreview(args[1:], stdout)
	case "kde-center-page-sections-preview":
		return runKDECenterPageSectionsPreview(args[1:], stdout)
	case "kde-entrypoint-action-preview":
		return runKDEEntryPointActionPreview(args[1:], stdout)
	case "kde-entrypoints-preview":
		return runKDEEntryPointsPreview(args[1:], stdout)
	case "krunner-query-preview":
		return runKRunnerQueryPreview(args[1:], stdout)
	case "launch-intent-preview":
		return runLaunchIntentPreview(args[1:], stdout)
	case "mimeapps-preview":
		return runMIMEAppsPreview(args[1:], stdout)
	case "mode-switch-preview":
		return runModeSwitchPreview(args[1:], stdout)
	case "notification-preview":
		return runNotificationPreview(args[1:], stdout)
	case "permission-review-preview":
		return runPermissionReviewPreview(args[1:], stdout)
	case "portal-request-preview":
		return runPortalRequestPreview(args[1:], stdout)
	case "review-flow-preview":
		return runReviewFlowPreview(args[1:], stdout)
	case "runtime-live-owner-gate-preview":
		return runRuntimeLiveOwnerGatePreview(args[1:], stdout)
	case "runtime-method-parity-manifest-preview":
		return runRuntimeMethodParityManifestPreview(args[1:], stdout)
	case "runtime-owner-process-preview":
		return runRuntimeOwnerProcessPreview(args[1:], stdout)
	case "runtime-owner-readiness-preview":
		return runRuntimeOwnerReadinessPreview(args[1:], stdout)
	case "runtime-owner-recipe-trust-preview":
		return runRuntimeOwnerRecipeTrustPreview(args[1:], stdout)
	case "runtime-owner-route-manifest-preview":
		return runRuntimeOwnerRouteManifestPreview(args[1:], stdout)
	case "runtime-owner-smoke-plan-preview":
		return runRuntimeOwnerSmokePlanPreview(args[1:], stdout)
	case "runtime-service-binding-preview":
		return runRuntimeServiceBindingPreview(args[1:], stdout)
	case "settings-change-preview":
		return runSettingsChangePreview(args[1:], stdout)
	case "settings-preview":
		return runSettingsPreview(args[1:], stdout)
	case "tray-status-preview":
		return runTrayStatusPreview(args[1:], stdout)
	case "window-identity-preview":
		return runWindowIdentityPreview(args[1:], stdout)
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runApplicationsPreview(args []string, stdout io.Writer) error {
	recipes, provenance, err := parseRegistryPreviewSource("applications-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewApplicationsPreview(recipes, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runApplicationPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("application-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewApplicationPreview(recipe, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runCompatibilityCenterPreview(args []string, stdout io.Writer) error {
	recipes, provenance, err := parseRegistryPreviewSource("compatibility-center-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewCompatibilityCenterPreview(recipes, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendSelectionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-selection-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendSelectionPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendEnvironmentPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-environment-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendEnvironmentPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runBackendBindingPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("backend-binding-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.BackendBindingPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDiagnosticsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("diagnostics-preview", args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDiagnosticsPreview(recipe, provenance)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionReadinessPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("execution-readiness-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionReadinessPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runLaunchIntentPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseLaunchIntentPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.LaunchIntentPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionRequestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseExecutionRequestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionRequestPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, fileURIs, err := parseExecutionReviewPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionReviewPreview(fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionDecisionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionDecisionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionDecisionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionPreflightPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionResourceGrantPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionResourceGrantPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionResourceGrantPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionTransactionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionTransactionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionTransactionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionSessionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionSessionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionSessionPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runExecutionSessionStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseExecutionSessionStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ExecutionSessionStatusPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runFileOpenPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, err := parseFileOpenPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewFileOpenPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDolphinDropPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, err := parseDolphinDropPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDolphinDropPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDolphinAIAnalysisPreview(args []string, stdout io.Writer) error {
	recipes, provenance, applicationID, fileURIs, err := parseDolphinAIAnalysisPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewDolphinAIAnalysisPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEEntryPointsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEEntryPointsPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEEntryPointsPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEEntryPointActionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, entryPointID, decision, fileURIs, err := parseKDEEntryPointActionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEEntryPointActionPreview(entryPointID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionQueuePreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEActionQueuePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionQueuePreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionCardDeckPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDEActionCardDeckPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionCardDeckPreview(decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPagePreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDECenterPagePreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPagePreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPageSectionsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, decision, fileURIs, err := parseKDECenterPageSectionsPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPageSectionsPreview(recipe, provenance, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDECenterPageSectionDetailPreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, decision, fileURIs, err := parseKDECenterPageSectionDetailPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKDECenterPageSectionDetailPreview(recipe, provenance, sectionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionReviewPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionReviewPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionPreflightPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionReceiptPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionReceiptPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionReceiptPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionStatusPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKDEActionCardPreview(args []string, stdout io.Writer) error {
	recipe, provenance, actionID, decision, fileURIs, err := parseKDEActionCardPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.KDEActionCardPreview(actionID, decision, fileURIs)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runKRunnerQueryPreview(args []string, stdout io.Writer) error {
	recipes, provenance, query, err := parseKRunnerQueryPreviewSource(args)
	if err != nil {
		return err
	}
	preview, err := appidentity.NewKRunnerQueryPreview(recipes, provenance, query)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationBundlePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-activation-bundle-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationBundlePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationPreflightPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationPreflightPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationPreflightPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationStagingPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationStagingPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationStagingPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationTransactionPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationTransactionPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationTransactionPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopActivationStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, mode, err := parseDesktopActivationStatusPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopActivationStatusPreview(mode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runDesktopIdentityPlan(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-identity-plan", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(plan)
}

func runDesktopEntryPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-entry-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	entry, err := plan.RenderDesktopEntry()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, entry)
	return err
}

func runDesktopIconPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-icon-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopIconPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runMIMEAppsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("mimeapps-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	mimeapps, err := plan.RenderMIMEApps()
	if err != nil {
		return err
	}
	_, err = io.WriteString(stdout, mimeapps)
	return err
}

func runDesktopResourceBridgePreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("desktop-resource-bridge-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runNotificationPreview(args []string, stdout io.Writer) error {
	recipe, provenance, eventType, err := parseNotificationPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.NotificationPreview(eventType)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("settings-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runModeSwitchPreview(args []string, stdout io.Writer) error {
	recipe, provenance, requestedMode, err := parseModeSwitchPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ModeSwitchPreview(requestedMode)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runSettingsChangePreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, err := parseSettingsChangePreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.SettingsChangePreview(sectionID, fieldID, value)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runReviewFlowPreview(args []string, stdout io.Writer) error {
	recipe, provenance, sectionID, fieldID, value, operation, err := parseReviewFlowPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.ReviewFlowPreview(sectionID, fieldID, value, operation)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPermissionReviewPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("permission-review-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PermissionReviewPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runPortalRequestPreview(args []string, stdout io.Writer) error {
	recipe, provenance, operation, reason, err := parsePortalRequestPreviewSource(args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.PortalRequestPreview(operation, reason)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runTrayStatusPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("tray-status-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.TrayStatusPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func runWindowIdentityPreview(args []string, stdout io.Writer) error {
	recipe, provenance, err := parseRecipeSource("window-identity-preview", args)
	if err != nil {
		return err
	}
	plan, err := appidentity.NewPlanWithProvenance(recipe, provenance)
	if err != nil {
		return err
	}
	preview, err := plan.WindowIdentityPreview()
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(preview)
}

func parseNotificationPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("notification-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	eventType := flags.String("event", "", "notification event type")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *eventType == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires --event")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("notification-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *eventType, err
}

func parsePortalRequestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, error) {
	flags := flag.NewFlagSet("portal-request-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	operation := flags.String("operation", "", "sensitive desktop operation")
	reason := flags.String("reason", "", "user-facing Portal request reason")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", err
	}
	if *operation == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --operation")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", errors.New("portal-request-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *operation, *reason, err
}

func parseModeSwitchPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("mode-switch-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	requestedMode := flags.String("mode", "", "requested compatibility mode")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if *requestedMode == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --mode")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("mode-switch-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *requestedMode, err
}

func parseReviewFlowPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, string, error) {
	flags := flag.NewFlagSet("review-flow-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "resource-access", "settings section identifier")
	fieldID := flags.String("field", "documents", "settings field identifier")
	value := flags.String("value", "ask", "requested settings value")
	operation := flags.String("operation", "file-open", "sensitive desktop operation")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", "", errors.New("review-flow-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, *operation, err
}

func parseSettingsChangePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, string, error) {
	flags := flag.NewFlagSet("settings-change-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	sectionID := flags.String("section", "", "settings section identifier")
	fieldID := flags.String("field", "", "settings field identifier")
	value := flags.String("value", "", "requested settings value")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", err
	}
	if *sectionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --section")
	}
	if *fieldID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --field")
	}
	if *value == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --value")
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", "", errors.New("settings-change-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *fieldID, *value, err
}

func parseKRunnerQueryPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("krunner-query-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	query := flags.String("query", "", "KRunner query text")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", errors.New("krunner-query-preview requires --registry")
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, "", errors.New("krunner-query-preview does not accept positional arguments")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *query, err
}

func parseRegistryPreviewSource(commandName string, args []string) ([]appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s requires --registry", commandName)
	}
	if flags.NArg() != 0 {
		return nil, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
}

func parseFileOpenPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet("file-open-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the file-open preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, errors.New("file-open-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, errors.New("file-open-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, err
}

func parseDolphinDropPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet("dolphin-drop-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the Dolphin drop preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-drop-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-drop-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, err
}

func parseDolphinAIAnalysisPreviewSource(args []string) ([]appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet("dolphin-ai-analysis-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to use for the Dolphin AI analysis preview")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	if err := flags.Parse(args); err != nil {
		return nil, appidentity.Provenance{}, "", nil, err
	}
	if *registryPath == "" {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-ai-analysis-preview requires --registry")
	}
	fileURIs := flags.Args()
	if len(fileURIs) == 0 {
		return nil, appidentity.Provenance{}, "", nil, errors.New("dolphin-ai-analysis-preview requires at least one file URI")
	}

	recipes, provenance, err := appidentity.LoadRecipesFromRegistry(*registryPath, *recipeRoot)
	return recipes, provenance, *applicationID, fileURIs, err
}

func parseLaunchIntentPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("launch-intent-preview", args)
}

func parseExecutionRequestPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("execution-request-preview", args)
}

func parseExecutionReviewPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	return parseLaunchActionPreviewSource("execution-review-preview", args)
}

func parseExecutionDecisionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-decision-preview", args)
}

func parseExecutionPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-preflight-preview", args)
}

func parseExecutionResourceGrantPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-resource-grant-preview", args)
}

func parseExecutionTransactionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-transaction-preview", args)
}

func parseExecutionSessionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-session-preview", args)
}

func parseExecutionSessionStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("execution-session-status-preview", args)
}

func parseKDEEntryPointsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-entrypoints-preview", args)
}

func parseKDEActionQueuePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-action-queue-preview", args)
}

func parseKDEActionCardDeckPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-action-card-deck-preview", args)
}

func parseKDECenterPagePreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-center-page-preview", args)
}

func parseKDECenterPageSectionsPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	return parseLaunchDecisionPreviewSource("kde-center-page-sections-preview", args)
}

func parseKDECenterPageSectionDetailPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-center-page-section-detail-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	sectionID := flags.String("section", "", "KDE Compatibility Center section id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --app when --registry is used")
	}
	if *sectionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --section")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-center-page-section-detail-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *sectionID, *decision, flags.Args(), err
}

func parseKDEActionReviewPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-review-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-preflight-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionReceiptPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-receipt-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-status-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-status-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEActionCardPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-action-card-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	actionID := flags.String("action", "", "KDE action queue item id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --app when --registry is used")
	}
	if *actionID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --action")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-action-card-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *actionID, *decision, flags.Args(), err
}

func parseKDEEntryPointActionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, string, []string, error) {
	flags := flag.NewFlagSet("kde-entrypoint-action-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	entryPointID := flags.String("entrypoint", "", "KDE entrypoint id")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --app when --registry is used")
	}
	if *entryPointID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --entrypoint")
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview requires --decision")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", "", nil, errors.New("kde-entrypoint-action-preview --recipe cannot be combined with --app or --recipe-root")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *entryPointID, *decision, flags.Args(), err
}

func parseLaunchDecisionPreviewSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, string, []string, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	decision := flags.String("decision", "", "review decision: reviewed, approved, deferred, or rejected")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *decision == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s requires --decision", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", nil, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *decision, flags.Args(), err
}

func parseLaunchActionPreviewSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, []string, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, nil, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, flags.Args(), err
}

func parseRecipeSource(commandName string, args []string) (appidentity.Recipe, appidentity.Provenance, error) {
	flags := flag.NewFlagSet(commandName, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires exactly one source: --recipe or --registry", commandName)
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s requires --app when --registry is used", commandName)
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s --recipe cannot be combined with --app or --recipe-root", commandName)
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("%s does not accept positional arguments", commandName)
	}

	return loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
}

func parseDesktopActivationPreflightPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation preflight mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-preflight-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationStagingPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-staging-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation staging mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-staging-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationTransactionPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-transaction-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation transaction mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-transaction-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func parseDesktopActivationStatusPreviewSource(args []string) (appidentity.Recipe, appidentity.Provenance, string, error) {
	flags := flag.NewFlagSet("desktop-activation-status-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	applicationID := flags.String("app", "", "application id to load from the recipe registry")
	registryPath := flags.String("registry", "", "path to an Xnix recipe registry JSON file")
	recipeRoot := flags.String("recipe-root", "", "directory containing registered recipe files; defaults to the registry directory")
	recipePath := flags.String("recipe", "", "path to an Xnix application recipe JSON file")
	mode := flags.String("mode", "production", "activation status mode: production or development")
	if err := flags.Parse(args); err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", err
	}
	if (*recipePath == "") == (*registryPath == "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-status-preview requires exactly one source: --recipe or --registry")
	}
	if *registryPath != "" && *applicationID == "" {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-status-preview requires --app when --registry is used")
	}
	if *recipePath != "" && (*applicationID != "" || *recipeRoot != "") {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-status-preview --recipe cannot be combined with --app or --recipe-root")
	}
	if flags.NArg() != 0 {
		return appidentity.Recipe{}, appidentity.Provenance{}, "", errors.New("desktop-activation-status-preview does not accept positional arguments")
	}

	recipe, provenance, err := loadRecipe(*recipePath, *registryPath, *recipeRoot, *applicationID)
	return recipe, provenance, *mode, err
}

func loadRecipe(recipePath string, registryPath string, recipeRoot string, applicationID string) (appidentity.Recipe, appidentity.Provenance, error) {
	if registryPath != "" {
		return appidentity.LoadRecipeFromRegistry(registryPath, recipeRoot, applicationID)
	}

	data, err := os.ReadFile(recipePath)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, fmt.Errorf("read recipe: %w", err)
	}
	recipe, err := appidentity.ParseRecipe(data)
	if err != nil {
		return appidentity.Recipe{}, appidentity.Provenance{}, err
	}
	return recipe, appidentity.Provenance{Source: "direct-file"}, nil
}
