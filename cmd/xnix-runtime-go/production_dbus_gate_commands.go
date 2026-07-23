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

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementAuthorizationReceiptAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt writer authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptWriterAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt acceptance authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt consumption gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumptionGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt consumer authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement receipt consumer enablement gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementReceiptConsumerEnablementGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe status fan-out inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusFanOutAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe status persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeStatusPersistenceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence write-model inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceWriteModelAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer authorization receipt acceptance inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAuthorizationReceiptAcceptanceAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer accepted receipt gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterAcceptedReceiptGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer grant inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterGrantAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer enablement inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterEnablementAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed writer implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedWriterImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status writer persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusWriterPersistenceAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed persistence writer inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedPersistenceWriterPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status storage persistence gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusStoragePersistenceGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status persistence record contract inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusPersistenceRecordContractPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root policy inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootPolicyAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt consumption inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptConsumptionAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup evidence inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupEvidencePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result boundary inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultBoundaryPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionEvidenceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection evidence audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionEvidenceAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route authorization evidence gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteAuthorizationEvidenceGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement evidence gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementEvidenceGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement receipt gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement receipt acceptance audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementReceiptAcceptanceAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement accepted receipt gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementAcceptedReceiptGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route grant audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteGrantAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route enablement audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteEnablementAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunExecutionGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultCaptureAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redaction boundary audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactionBoundaryAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result redacted presentation eligibility audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultRedactedPresentationEligibilityAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result presentation consent audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultPresentationConsentAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification presentation audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationPresentationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery grant audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryGrantAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action enablement audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionEnablementAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result visibility audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultVisibilityAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAcceptanceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-acceptance-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt acceptance audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-acceptance-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptAcceptanceAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumptionGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumption-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumption gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumption-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumptionGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptWriterAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-writer-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt writer authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-writer-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptWriterAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptanceGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptedReceiptGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-accepted-receipt-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt accepted receipt gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-accepted-receipt-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptAcceptedReceiptGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptGrantAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-grant-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt grant audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-grant-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptGrantAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptEnablementAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-enablement-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt enablement audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-enablement-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptEnablementAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed persistence implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage persistence gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordContractPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedStorageRecordWriterImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed storage record writer implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptClosedStorageRecordWriterImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization receipt acceptance inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAuthorizationReceiptAcceptancePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer accepted receipt gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterAcceptedReceiptGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterGrantPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer grant inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterGrantPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterEnablementPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer enablement inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterEnablementPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt acceptance inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptancePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate grant inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate enablement inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptancePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateGrantPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateEnablementPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call-gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptAcceptancePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptAcceptancePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumptionGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumptionGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementGatePreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementGatePreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptPersistenceAuthorizationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt closed persistence implementation inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptClosedPersistenceImplementationPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt storage persistence gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStoragePersistenceGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordContractPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-contract-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt storage record contract inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-contract-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationActionRequestObjectDispatchDryRunResultPersistenceReceiptConsumerEnablementReceiptStorageRecordWriterCallAuthorizationReceiptAcceptedReceiptGateCallAuthorizationReceiptAcceptedReceiptGateCallReceiptConsumerEnablementReceiptStorageRecordContractPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsentAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery explicit consent audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExplicitConsentAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection receipt audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAcceptanceAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-acceptance-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection receipt acceptance audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-acceptance-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptAcceptanceAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumptionGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection receipt consumption gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumptionGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection receipt consumer authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerAuthorizationAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery consent collection receipt consumer enablement gate audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryConsentCollectionReceiptConsumerEnablementGateAuditPreview(*root)
	if err != nil {
		return err
	}
	return encodeIndentedJSON(stdout, preview)
}

func runProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExecutionAuthorizationAuditPreview(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-execution-authorization-audit-preview", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	root := flags.String("root", ".", "project root containing Runtime owner production receipt notification action dry-run result lookup consumer enablement KDE-safe redacted status closed record writer storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result notification delivery execution authorization audit inputs")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 {
		return errors.New("production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-execution-authorization-audit-preview does not accept positional arguments")
	}
	preview, err := owner.NewProductionReceiptNotificationActionDryRunResultLookupConsumerEnablementKDESafeRedactedStatusClosedRecordWriterStorageRootAuthorizationReceiptDryRunLookupResultConsumerProjectionRouteEnablementLookupRouteDispatchDryRunResultNotificationDeliveryExecutionAuthorizationAuditPreview(*root)
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
