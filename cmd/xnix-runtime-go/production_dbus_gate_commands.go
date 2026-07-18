package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"xnix.local/xnix/internal/runtime/owner"
)

func runProductionDBusGateReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-gate-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner gate review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-gate-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusGateReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDBusHumanAuthorizationPreflightPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-human-authorization-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner human authorization preflight inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-human-authorization-preflight-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusHumanAuthorizationPreflightPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionHumanAuthorizationReceiptConsolidationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-human-authorization-receipt-consolidation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-human-authorization-receipt-consolidation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionHumanAuthorizationReceiptConsolidationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionAuthorizationConsumptionAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-authorization-consumption-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production authorization consumption inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-authorization-consumption-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionAuthorizationConsumptionAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptAcceptancePropagationPreflightPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-acceptance-propagation-preflight-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt acceptance propagation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-acceptance-propagation-preflight-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptAcceptancePropagationPreflightPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptWriterAuthorizationReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-writer-authorization-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt writer authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-writer-authorization-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptWriterAuthorizationReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptPersistenceThreatReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-persistence-threat-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt persistence threat inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-persistence-threat-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptPersistenceThreatReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptRevocationVisibilityAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-revocation-visibility-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt revocation visibility inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-revocation-visibility-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptRevocationVisibilityAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationDeliveryGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-delivery-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification delivery inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-delivery-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationDeliveryGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionSafetyAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-safety-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-safety-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionSafetyAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionRequestObjectAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-request-object-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action request-object inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-request-object-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionRequestObjectAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dispatch-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dispatch authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dispatch-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDispatchAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDispatchDryRunAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dispatch-dry-run-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dispatch dry-run inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dispatch-dry-run-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDispatchDryRunAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-visibility-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result visibility inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-visibility-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultVisibilityAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultPersistenceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result retention redaction policy inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultRetentionRedactionPolicyAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result opaque lookup inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultOpaqueLookupAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup route authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupRouteAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer redaction inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerRedactionAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDBusMethodReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-dbus-method-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner method review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-dbus-method-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDBusMethodReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionRollbackDiagnosticsReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-rollback-diagnostics-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner rollback diagnostics review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-rollback-diagnostics-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionRollbackDiagnosticsReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionDesktopSideEffectReviewPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-desktop-side-effect-review-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner desktop side-effect review inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-desktop-side-effect-review-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionDesktopSideEffectReviewPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}
