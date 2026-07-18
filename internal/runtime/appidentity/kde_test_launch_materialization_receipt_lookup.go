package appidentity

import (
	"errors"
	"path/filepath"
)

const KDETestLaunchMaterializationOpaqueReceiptID = "kde-test-launch-materialization-receipt-id"

type KDETestLaunchMaterializationReceiptLookupPreview struct {
	Version                                 string                  `json:"version"`
	SchemaVersion                           string                  `json:"schema_version"`
	RequestType                             string                  `json:"request_type"`
	LookupType                              string                  `json:"lookup_type"`
	Source                                  string                  `json:"source"`
	RuntimeMethod                           string                  `json:"runtime_method"`
	ReadMethod                              string                  `json:"read_method"`
	OpaqueMaterializationReceiptID          string                  `json:"opaque_materialization_receipt_id"`
	SupportedOpaqueMaterializationReceiptID string                  `json:"supported_opaque_materialization_receipt_id"`
	MaterializationPlanID                   string                  `json:"materialization_plan_id"`
	ReceiptRelativePath                     string                  `json:"receipt_relative_path"`
	ReceiptRecordType                       string                  `json:"receipt_record_type"`
	ReceiptLookupState                      string                  `json:"receipt_lookup_state"`
	ReceiptAvailable                        bool                    `json:"receipt_available"`
	ReceiptConsumed                         bool                    `json:"receipt_consumed"`
	MissingReceiptSafe                      bool                    `json:"missing_receipt_safe"`
	OwnerManagedOpaqueReceiptLookupReady    bool                    `json:"owner_managed_opaque_receipt_lookup_ready"`
	OpaqueMaterializationReceiptIDSupported bool                    `json:"opaque_materialization_receipt_id_supported"`
	RequiresCallerRegistryPath              bool                    `json:"requires_caller_registry_path"`
	RequiresCallerApplicationID             bool                    `json:"requires_caller_application_id"`
	RequiresCallerStateRoot                 bool                    `json:"requires_caller_state_root"`
	RequiresExplicitAuthorization           bool                    `json:"requires_explicit_authorization"`
	ReadOnlyLookup                          bool                    `json:"read_only_lookup"`
	Checks                                  []KDEFakeExecutionCheck `json:"checks"`
	CheckIDs                                []string                `json:"check_ids"`
	CheckCount                              int                     `json:"check_count"`
	PassedCheckCount                        int                     `json:"passed_check_count"`
	AllChecksPassed                         bool                    `json:"all_checks_passed"`
	RuntimeOwned                            bool                    `json:"runtime_owned"`
	GoRuntimeBacked                         bool                    `json:"go_runtime_backed"`
	KDEPolicyOwner                          bool                    `json:"kde_policy_owner"`
	StateRootPathExposed                    bool                    `json:"state_root_path_exposed"`
	StateRootWritesEnabled                  bool                    `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled                    bool                    `json:"runtime_writes_enabled"`
	MaterializationWritesEnabled            bool                    `json:"materialization_writes_enabled"`
	FanOutWritesEnabled                     bool                    `json:"fan_out_writes_enabled"`
	SystemServiceStarted                    bool                    `json:"system_service_started"`
	SessionBusClaimed                       bool                    `json:"session_bus_claimed"`
	ProductionBusClaimed                    bool                    `json:"production_bus_claimed"`
	WriteMethodsEnabled                     bool                    `json:"write_methods_enabled"`
	LaunchAuthorized                        bool                    `json:"launch_authorized"`
	ExecutionApproved                       bool                    `json:"execution_approved"`
	ProcessStartAuthorized                  bool                    `json:"process_start_authorized"`
	CommandMaterialized                     bool                    `json:"command_materialized"`
	ExecutablePathResolved                  bool                    `json:"executable_path_resolved"`
	BackendSelectedForLaunch                bool                    `json:"backend_selected_for_launch"`
	BackendLaunchEnabled                    bool                    `json:"backend_launch_enabled"`
	BackendProcessStarted                   bool                    `json:"backend_process_started"`
	NetworkRequired                         bool                    `json:"network_required"`
	HostRootModified                        bool                    `json:"host_root_modified"`
	PrivilegedContainerRequired             bool                    `json:"privileged_container_required"`
	RawCommandExposed                       bool                    `json:"raw_command_exposed"`
	RawExecutableExposed                    bool                    `json:"raw_executable_exposed"`
	BackendDetailsExposed                   bool                    `json:"backend_details_exposed"`
	BlockedActions                          []string                `json:"blocked_actions"`
	NextRequirements                        []string                `json:"next_requirements"`
	DesktopSafeSummary                      string                  `json:"desktop_safe_summary"`
}

func ResolveKDETestLaunchMaterializationReceipt(root string, opaqueReceiptID string) (KDETestLaunchMaterializationReceiptLookupPreview, error) {
	if opaqueReceiptID != KDETestLaunchMaterializationOpaqueReceiptID {
		return KDETestLaunchMaterializationReceiptLookupPreview{}, errors.New("unsupported KDE test launch materialization opaque receipt id")
	}
	version, err := readRuntimeServiceBindingVersion(root)
	if err != nil {
		return KDETestLaunchMaterializationReceiptLookupPreview{}, err
	}
	lookup := KDETestLaunchMaterializationReceiptLookupPreview{
		Version:                                 version,
		SchemaVersion:                           "xnix.runtime.kde_test_launch_materialization_receipt_lookup.v1",
		RequestType:                             "kde-test-launch-materialization-receipt-lookup-preview",
		LookupType:                              "owner-managed-materialization-receipt-lookup",
		Source:                                  "runtime-owner-dispatch+opaque-materialization-receipt-id-registry",
		RuntimeMethod:                           "GetKDETestLaunchMaterializationReceiptLookup",
		ReadMethod:                              "GetKDETestLaunchMaterializationReceiptLookupPreview",
		OpaqueMaterializationReceiptID:          opaqueReceiptID,
		SupportedOpaqueMaterializationReceiptID: KDETestLaunchMaterializationOpaqueReceiptID,
		MaterializationPlanID:                   "kde-test-launch-materialization-plan",
		ReceiptRelativePath:                     kdeTestLaunchMaterializationReceiptLookupPath(),
		ReceiptRecordType:                       "restricted-launch-materialization-plan",
		ReceiptLookupState:                      "missing-receipt",
		ReceiptAvailable:                        false,
		ReceiptConsumed:                         false,
		MissingReceiptSafe:                      true,
		OwnerManagedOpaqueReceiptLookupReady:    true,
		OpaqueMaterializationReceiptIDSupported: true,
		RequiresCallerRegistryPath:              false,
		RequiresCallerApplicationID:             false,
		RequiresCallerStateRoot:                 false,
		RequiresExplicitAuthorization:           false,
		ReadOnlyLookup:                          true,
		RuntimeOwned:                            true,
		GoRuntimeBacked:                         true,
		KDEPolicyOwner:                          false,
		StateRootPathExposed:                    false,
		StateRootWritesEnabled:                  false,
		RuntimeWritesEnabled:                    false,
		MaterializationWritesEnabled:            false,
		FanOutWritesEnabled:                     false,
		SystemServiceStarted:                    false,
		SessionBusClaimed:                       false,
		ProductionBusClaimed:                    false,
		WriteMethodsEnabled:                     false,
		LaunchAuthorized:                        false,
		ExecutionApproved:                       false,
		ProcessStartAuthorized:                  false,
		CommandMaterialized:                     false,
		ExecutablePathResolved:                  false,
		BackendSelectedForLaunch:                false,
		BackendLaunchEnabled:                    false,
		BackendProcessStarted:                   false,
		NetworkRequired:                         false,
		HostRootModified:                        false,
		PrivilegedContainerRequired:             false,
		RawCommandExposed:                       false,
		RawExecutableExposed:                    false,
		BackendDetailsExposed:                   false,
		BlockedActions: []string{
			"resolve arbitrary caller-supplied materialization receipt paths",
			"accept caller registry, application id, state-root, or authorization inputs through owner lookup",
			"consume a missing materialization receipt as fan-out evidence",
			"materialize receipts or write state roots from the owner-managed lookup",
			"launch, materialize commands, expose raw executable data, or mutate host root from lookup",
		},
		NextRequirements: []string{
			"Attach the opaque materialization receipt id to a Runtime-owned receipt registry.",
			"Allow materialization fan-out to consume the owner lookup result instead of caller paths.",
			"Keep missing receipts fail-closed until digest-verified evidence is available.",
			"Promote only owner-local read coverage before any production D-Bus exposure.",
		},
		DesktopSafeSummary: "Runtime owner can resolve the materialization opaque receipt id to a safe receipt slot, but the receipt remains missing and cannot satisfy fan-out until digest-verified evidence exists.",
	}
	checks := kdeTestLaunchMaterializationReceiptLookupChecks(lookup)
	lookup.Checks = checks
	lookup.CheckIDs = kdeTestLaunchMaterializationReceiptLookupCheckIDs(checks)
	lookup.CheckCount = len(checks)
	lookup.PassedCheckCount = countKDETestLaunchMaterializationReceiptLookupPasses(checks)
	lookup.AllChecksPassed = lookup.PassedCheckCount == lookup.CheckCount
	if err := validateNoBackendTerms(lookup, "KDE test launch materialization receipt lookup preview"); err != nil {
		return KDETestLaunchMaterializationReceiptLookupPreview{}, err
	}
	return lookup, nil
}

func kdeTestLaunchMaterializationReceiptLookupPath() string {
	return filepath.ToSlash(filepath.Join("execution-ledger", "materialization-plans", "kde-test-launch-materialization-receipt.json"))
}

func kdeTestLaunchMaterializationReceiptLookupCheckIDs(checks []KDEFakeExecutionCheck) []string {
	ids := make([]string, 0, len(checks))
	for _, check := range checks {
		ids = append(ids, check.ID)
	}
	return ids
}

func countKDETestLaunchMaterializationReceiptLookupPasses(checks []KDEFakeExecutionCheck) int {
	passed := 0
	for _, check := range checks {
		if check.Status == "pass" {
			passed++
		}
	}
	return passed
}

func kdeTestLaunchMaterializationReceiptLookupChecks(lookup KDETestLaunchMaterializationReceiptLookupPreview) []KDEFakeExecutionCheck {
	return []KDEFakeExecutionCheck{
		fakeExecutionCheck("opaque-materialization-receipt-id-supported", lookup.OpaqueMaterializationReceiptID == KDETestLaunchMaterializationOpaqueReceiptID && lookup.OpaqueMaterializationReceiptIDSupported, "The owner recognizes the stable materialization opaque receipt id."),
		fakeExecutionCheck("owner-managed-opaque-lookup", lookup.OwnerManagedOpaqueReceiptLookupReady && lookup.LookupType == "owner-managed-materialization-receipt-lookup", "Materialization receipt lookup is owned by the Runtime owner."),
		fakeExecutionCheck("caller-paths-hidden", !lookup.RequiresCallerRegistryPath && !lookup.RequiresCallerApplicationID && !lookup.RequiresCallerStateRoot && !lookup.StateRootPathExposed && !filepath.IsAbs(lookup.ReceiptRelativePath), "Owner lookup does not accept or expose caller registry, application, or state-root paths."),
		fakeExecutionCheck("missing-receipt-fails-closed", !lookup.ReceiptAvailable && !lookup.ReceiptConsumed && lookup.MissingReceiptSafe && lookup.ReceiptLookupState == "missing-receipt", "Missing materialization receipt evidence stays blocked for downstream fan-out."),
		fakeExecutionCheck("read-only-lookup", lookup.ReadOnlyLookup && !lookup.StateRootWritesEnabled && !lookup.MaterializationWritesEnabled && !lookup.FanOutWritesEnabled && !lookup.RuntimeWritesEnabled, "Lookup is read-only and does not create or consume receipts."),
		fakeExecutionCheck("unsafe-gates-closed", !lookup.SystemServiceStarted && !lookup.SessionBusClaimed && !lookup.ProductionBusClaimed && !lookup.WriteMethodsEnabled && !lookup.LaunchAuthorized && !lookup.ExecutionApproved && !lookup.ProcessStartAuthorized && !lookup.CommandMaterialized && !lookup.ExecutablePathResolved && !lookup.BackendSelectedForLaunch && !lookup.BackendLaunchEnabled && !lookup.BackendProcessStarted && !lookup.NetworkRequired && !lookup.HostRootModified && !lookup.PrivilegedContainerRequired && !lookup.RawCommandExposed && !lookup.RawExecutableExposed && !lookup.BackendDetailsExposed, "Lookup keeps service ownership, writes, launch, command materialization, network, privilege, and host mutation disabled."),
	}
}
