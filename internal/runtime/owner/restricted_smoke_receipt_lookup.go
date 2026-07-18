package owner

import (
	"errors"
	"path/filepath"

	"xnix.local/xnix/internal/runtime/appidentity"
)

const RestrictedOwnerSmokeOpaqueReceiptID = "restricted-owner-smoke-receipt-id"

type RestrictedOwnerSmokeReceiptLookupPreview struct {
	Version                     string                      `json:"version"`
	SchemaVersion               string                      `json:"schema_version"`
	RequestType                 string                      `json:"request_type"`
	LookupType                  string                      `json:"lookup_type"`
	Source                      string                      `json:"source"`
	RuntimeMethod               string                      `json:"runtime_method"`
	ReadMethod                  string                      `json:"read_method"`
	OpaqueReceiptID             string                      `json:"opaque_receipt_id"`
	SupportedOpaqueReceiptID    string                      `json:"supported_opaque_receipt_id"`
	ReceiptID                   string                      `json:"receipt_id"`
	ReceiptRelativePath         string                      `json:"receipt_relative_path"`
	ReceiptRecordType           string                      `json:"receipt_record_type"`
	ReceiptLookupState          string                      `json:"receipt_lookup_state"`
	ReceiptAvailable            bool                        `json:"receipt_available"`
	ReceiptConsumed             bool                        `json:"receipt_consumed"`
	MissingReceiptSafe          bool                        `json:"missing_receipt_safe"`
	OwnerManagedLookup          bool                        `json:"owner_managed_lookup"`
	CallerStateRootRequired     bool                        `json:"caller_state_root_required"`
	OpaqueReceiptIDSupported    bool                        `json:"opaque_receipt_id_supported"`
	ReadOnlyLookup              bool                        `json:"read_only_lookup"`
	Checks                      []RestrictedOwnerSmokeCheck `json:"checks"`
	CheckIDs                    []string                    `json:"check_ids"`
	CheckCount                  int                         `json:"check_count"`
	PassedCheckCount            int                         `json:"passed_check_count"`
	AllChecksPassed             bool                        `json:"all_checks_passed"`
	StateRootPathExposed        bool                        `json:"state_root_path_exposed"`
	StateRootWritesEnabled      bool                        `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled        bool                        `json:"runtime_writes_enabled"`
	FanOutWritesEnabled         bool                        `json:"fan_out_writes_enabled"`
	ProductionOwnerEnabled      bool                        `json:"production_owner_enabled"`
	SystemServiceStarted        bool                        `json:"system_service_started"`
	SessionBusClaimed           bool                        `json:"session_bus_claimed"`
	ProductionBusClaimed        bool                        `json:"production_bus_claimed"`
	WriteMethodsEnabled         bool                        `json:"write_methods_enabled"`
	SupportBundleExported       bool                        `json:"support_bundle_exported"`
	SupportCaseCreated          bool                        `json:"support_case_created"`
	NotificationSent            bool                        `json:"notification_sent"`
	BackendLaunchEnabled        bool                        `json:"backend_launch_enabled"`
	BackendProcessStarted       bool                        `json:"backend_process_started"`
	NetworkRequired             bool                        `json:"network_required"`
	HostRootModified            bool                        `json:"host_root_modified"`
	PrivilegedContainerRequired bool                        `json:"privileged_container_required"`
	BackendDetailsExposed       bool                        `json:"backend_details_exposed"`
	BlockedActions              []string                    `json:"blocked_actions"`
	NextRequirements            []string                    `json:"next_requirements"`
	DesktopSafeSummary          string                      `json:"desktop_safe_summary"`
}

func ResolveRestrictedOwnerSmokeReceipt(root string, opaqueReceiptID string) (RestrictedOwnerSmokeReceiptLookupPreview, error) {
	if opaqueReceiptID != RestrictedOwnerSmokeOpaqueReceiptID {
		return RestrictedOwnerSmokeReceiptLookupPreview{}, errors.New("unsupported restricted owner smoke opaque receipt id")
	}
	manifest, err := appidentity.NewRuntimeOwnerRouteManifestPreview(root)
	if err != nil {
		return RestrictedOwnerSmokeReceiptLookupPreview{}, err
	}
	lookup := RestrictedOwnerSmokeReceiptLookupPreview{
		Version:                     manifest.Version,
		SchemaVersion:               "xnix.runtime.restricted_owner_smoke_receipt_lookup.v1",
		RequestType:                 "restricted-owner-smoke-receipt-lookup-preview",
		LookupType:                  "owner-managed-receipt-lookup",
		Source:                      "runtime-owner-dispatch+opaque-receipt-id-registry",
		RuntimeMethod:               "GetRestrictedOwnerSmokeReceiptLookup",
		ReadMethod:                  "GetRestrictedOwnerSmokeReceiptLookupPreview",
		OpaqueReceiptID:             opaqueReceiptID,
		SupportedOpaqueReceiptID:    RestrictedOwnerSmokeOpaqueReceiptID,
		ReceiptID:                   "restricted-owner-smoke-" + manifest.Version,
		ReceiptRelativePath:         restrictedOwnerSmokeReceiptPath(),
		ReceiptRecordType:           "restricted-owner-smoke-execution-receipt",
		ReceiptLookupState:          "missing-receipt",
		ReceiptAvailable:            false,
		ReceiptConsumed:             false,
		MissingReceiptSafe:          true,
		OwnerManagedLookup:          true,
		CallerStateRootRequired:     false,
		OpaqueReceiptIDSupported:    true,
		ReadOnlyLookup:              true,
		StateRootPathExposed:        false,
		StateRootWritesEnabled:      false,
		RuntimeWritesEnabled:        false,
		FanOutWritesEnabled:         false,
		ProductionOwnerEnabled:      false,
		SystemServiceStarted:        false,
		SessionBusClaimed:           false,
		ProductionBusClaimed:        false,
		WriteMethodsEnabled:         false,
		SupportBundleExported:       false,
		SupportCaseCreated:          false,
		NotificationSent:            false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		NetworkRequired:             false,
		HostRootModified:            false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"resolve arbitrary caller-supplied receipt paths",
			"read a caller-supplied state root from owner dispatch",
			"consume a missing restricted owner smoke receipt as readiness evidence",
			"fan out restricted owner smoke evidence before receipt lookup is attached to a verified receipt",
			"start services, claim D-Bus ownership, enable writes, launch, or mutate host root from receipt lookup",
		},
		NextRequirements: []string{
			"Attach the opaque receipt id to a Runtime-owned receipt registry.",
			"Allow fan-out to consume the owner lookup result instead of a caller state-root path.",
			"Keep missing receipts fail-closed until a digest-verified record is available.",
			"Promote only owner-local read coverage before any production D-Bus exposure.",
		},
		DesktopSafeSummary: "Runtime owner can resolve the restricted smoke opaque receipt id to a safe receipt slot, but the receipt is not consumed until digest-verified evidence exists.",
	}
	checks := restrictedOwnerSmokeReceiptLookupChecks(lookup)
	lookup.Checks = checks
	lookup.CheckIDs = restrictedOwnerSmokeCheckIDs(checks)
	lookup.CheckCount = len(checks)
	lookup.PassedCheckCount = countRestrictedOwnerSmokePasses(checks)
	lookup.AllChecksPassed = lookup.PassedCheckCount == lookup.CheckCount
	if err := validateNoBackendTerms(lookup, "restricted owner smoke receipt lookup preview"); err != nil {
		return RestrictedOwnerSmokeReceiptLookupPreview{}, err
	}
	return lookup, nil
}

func restrictedOwnerSmokeReceiptLookupChecks(lookup RestrictedOwnerSmokeReceiptLookupPreview) []RestrictedOwnerSmokeCheck {
	return []RestrictedOwnerSmokeCheck{
		restrictedOwnerSmokeCheck("opaque-receipt-id-supported", lookup.OpaqueReceiptID == RestrictedOwnerSmokeOpaqueReceiptID && lookup.OpaqueReceiptIDSupported, "The owner recognizes the stable restricted smoke opaque receipt id."),
		restrictedOwnerSmokeCheck("owner-managed-lookup", lookup.OwnerManagedLookup && lookup.LookupType == "owner-managed-receipt-lookup", "Receipt lookup is owned by the Runtime owner, not by KDE or caller paths."),
		restrictedOwnerSmokeCheck("caller-state-root-hidden", !lookup.CallerStateRootRequired && !lookup.StateRootPathExposed, "Owner-local lookup does not accept or expose caller state-root paths."),
		restrictedOwnerSmokeCheck("receipt-slot-redacted", !filepath.IsAbs(lookup.ReceiptRelativePath) && lookup.ReceiptRelativePath == restrictedOwnerSmokeReceiptPath(), "Lookup exposes only a stable relative receipt slot."),
		restrictedOwnerSmokeCheck("missing-receipt-fails-closed", !lookup.ReceiptAvailable && !lookup.ReceiptConsumed && lookup.MissingReceiptSafe && lookup.ReceiptLookupState == "missing-receipt", "Missing receipt evidence stays safe and blocked for downstream fan-out."),
		restrictedOwnerSmokeCheck("unsafe-gates-closed", !lookup.StateRootWritesEnabled && !lookup.RuntimeWritesEnabled && !lookup.FanOutWritesEnabled && !lookup.ProductionOwnerEnabled && !lookup.SystemServiceStarted && !lookup.SessionBusClaimed && !lookup.ProductionBusClaimed && !lookup.WriteMethodsEnabled && !lookup.SupportBundleExported && !lookup.SupportCaseCreated && !lookup.NotificationSent && !lookup.BackendLaunchEnabled && !lookup.BackendProcessStarted && !lookup.NetworkRequired && !lookup.HostRootModified && !lookup.PrivilegedContainerRequired && !lookup.BackendDetailsExposed, "Lookup keeps writes, service ownership, support side effects, launch, network, privilege, and host mutation disabled."),
	}
}
