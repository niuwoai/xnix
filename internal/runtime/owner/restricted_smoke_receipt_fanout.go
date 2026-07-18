package owner

import (
	"errors"
	"path/filepath"
)

type RestrictedOwnerSmokeReceiptFanOutPreview struct {
	Version                      string                                  `json:"version"`
	SchemaVersion                string                                  `json:"schema_version"`
	RequestType                  string                                  `json:"request_type"`
	Source                       string                                  `json:"source"`
	RuntimeMethod                string                                  `json:"runtime_method"`
	ReadMethod                   string                                  `json:"read_method"`
	Mode                         string                                  `json:"mode"`
	ReceiptID                    string                                  `json:"receipt_id"`
	ReceiptRelativePath          string                                  `json:"receipt_relative_path"`
	ReceiptSHA256                string                                  `json:"receipt_sha256"`
	ReceiptRecordType            string                                  `json:"receipt_record_type"`
	ReceiptReadBack              bool                                    `json:"receipt_read_back"`
	ReceiptConsumed              bool                                    `json:"receipt_consumed"`
	ReceiptAllChecksPassed       bool                                    `json:"receipt_all_checks_passed"`
	ReceiptCheckCount            int                                     `json:"receipt_check_count"`
	SmokeReadDispatchRecordCount int                                     `json:"smoke_read_dispatch_record_count"`
	SmokeWriteDenialRecordCount  int                                     `json:"smoke_write_denial_record_count"`
	SurfaceCount                 int                                     `json:"surface_count"`
	Surfaces                     []RestrictedOwnerSmokeReceiptFanSurface `json:"surfaces"`
	RuntimeOwnerReadiness        RestrictedOwnerSmokeReceiptFanSurface   `json:"runtime_owner_readiness"`
	ServiceActivationPreflight   RestrictedOwnerSmokeReceiptFanSurface   `json:"service_activation_preflight"`
	CompatibilityOnboarding      RestrictedOwnerSmokeReceiptFanSurface   `json:"compatibility_onboarding"`
	SupportBundleManifest        RestrictedOwnerSmokeReceiptFanSurface   `json:"support_bundle_manifest"`
	SupportCaseTimeline          RestrictedOwnerSmokeReceiptFanSurface   `json:"support_case_timeline"`
	Checks                       []RestrictedOwnerSmokeCheck             `json:"checks"`
	CheckCount                   int                                     `json:"check_count"`
	PassedCheckCount             int                                     `json:"passed_check_count"`
	AllChecksPassed              bool                                    `json:"all_checks_passed"`
	RuntimeOwned                 bool                                    `json:"runtime_owned"`
	GoRuntimeBacked              bool                                    `json:"go_runtime_backed"`
	KDEPolicyOwner               bool                                    `json:"kde_policy_owner"`
	ReadOnlyFanOut               bool                                    `json:"read_only_fan_out"`
	ReadinessSurfacesSatisfied   bool                                    `json:"readiness_surfaces_satisfied"`
	SupportSurfacesSatisfied     bool                                    `json:"support_surfaces_satisfied"`
	StateRootPathExposed         bool                                    `json:"state_root_path_exposed"`
	StateRootWritesEnabled       bool                                    `json:"state_root_writes_enabled"`
	FanOutWritesEnabled          bool                                    `json:"fan_out_writes_enabled"`
	ProductionActivationReady    bool                                    `json:"production_activation_ready"`
	ProductionOwnerEnabled       bool                                    `json:"production_owner_enabled"`
	SystemServiceStarted         bool                                    `json:"system_service_started"`
	SessionBusClaimed            bool                                    `json:"session_bus_claimed"`
	ProductionBusClaimed         bool                                    `json:"production_bus_claimed"`
	WriteMethodsEnabled          bool                                    `json:"write_methods_enabled"`
	SupportBundleExported        bool                                    `json:"support_bundle_exported"`
	SupportCaseCreated           bool                                    `json:"support_case_created"`
	NotificationSent             bool                                    `json:"notification_sent"`
	BackendLaunchEnabled         bool                                    `json:"backend_launch_enabled"`
	NetworkRequired              bool                                    `json:"network_required"`
	HostRootModified             bool                                    `json:"host_root_modified"`
	PrivilegedContainerRequired  bool                                    `json:"privileged_container_required"`
	BackendDetailsExposed        bool                                    `json:"backend_details_exposed"`
	DesktopSafeSummary           string                                  `json:"desktop_safe_summary"`
}

type RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview struct {
	Version                       string                                  `json:"version"`
	SchemaVersion                 string                                  `json:"schema_version"`
	RequestType                   string                                  `json:"request_type"`
	RouteType                     string                                  `json:"route_type"`
	Source                        string                                  `json:"source"`
	RuntimeMethod                 string                                  `json:"runtime_method"`
	ReadMethod                    string                                  `json:"read_method"`
	OpaqueReceiptID               string                                  `json:"opaque_receipt_id"`
	SupportedOpaqueReceiptID      string                                  `json:"supported_opaque_receipt_id"`
	ReceiptID                     string                                  `json:"receipt_id"`
	ReceiptRelativePath           string                                  `json:"receipt_relative_path"`
	ReceiptRecordType             string                                  `json:"receipt_record_type"`
	ReceiptLookupState            string                                  `json:"receipt_lookup_state"`
	ReceiptAvailable              bool                                    `json:"receipt_available"`
	ReceiptConsumed               bool                                    `json:"receipt_consumed"`
	MissingReceiptSafe            bool                                    `json:"missing_receipt_safe"`
	OwnerManagedLookup            bool                                    `json:"owner_managed_lookup"`
	CallerStateRootRequired       bool                                    `json:"caller_state_root_required"`
	OpaqueReceiptIDSupported      bool                                    `json:"opaque_receipt_id_supported"`
	ReadOnlyFanOut                bool                                    `json:"read_only_fan_out"`
	FanOutResultState             string                                  `json:"fan_out_result_state"`
	OwnerLocalRouteCandidateReady bool                                    `json:"owner_local_route_candidate_ready"`
	ProductionDBusExposureReady   bool                                    `json:"production_dbus_exposure_ready"`
	SurfaceCount                  int                                     `json:"surface_count"`
	Surfaces                      []RestrictedOwnerSmokeReceiptFanSurface `json:"surfaces"`
	RuntimeOwnerReadiness         RestrictedOwnerSmokeReceiptFanSurface   `json:"runtime_owner_readiness"`
	ServiceActivationPreflight    RestrictedOwnerSmokeReceiptFanSurface   `json:"service_activation_preflight"`
	CompatibilityOnboarding       RestrictedOwnerSmokeReceiptFanSurface   `json:"compatibility_onboarding"`
	SupportBundleManifest         RestrictedOwnerSmokeReceiptFanSurface   `json:"support_bundle_manifest"`
	SupportCaseTimeline           RestrictedOwnerSmokeReceiptFanSurface   `json:"support_case_timeline"`
	Checks                        []RestrictedOwnerSmokeCheck             `json:"checks"`
	CheckIDs                      []string                                `json:"check_ids"`
	CheckCount                    int                                     `json:"check_count"`
	PassedCheckCount              int                                     `json:"passed_check_count"`
	AllChecksPassed               bool                                    `json:"all_checks_passed"`
	RuntimeOwned                  bool                                    `json:"runtime_owned"`
	GoRuntimeBacked               bool                                    `json:"go_runtime_backed"`
	KDEPolicyOwner                bool                                    `json:"kde_policy_owner"`
	ReadinessSurfacesSatisfied    bool                                    `json:"readiness_surfaces_satisfied"`
	SupportSurfacesSatisfied      bool                                    `json:"support_surfaces_satisfied"`
	StateRootPathExposed          bool                                    `json:"state_root_path_exposed"`
	StateRootWritesEnabled        bool                                    `json:"state_root_writes_enabled"`
	RuntimeWritesEnabled          bool                                    `json:"runtime_writes_enabled"`
	FanOutWritesEnabled           bool                                    `json:"fan_out_writes_enabled"`
	ProductionActivationReady     bool                                    `json:"production_activation_ready"`
	ProductionOwnerEnabled        bool                                    `json:"production_owner_enabled"`
	SystemServiceStarted          bool                                    `json:"system_service_started"`
	SessionBusClaimed             bool                                    `json:"session_bus_claimed"`
	ProductionBusClaimed          bool                                    `json:"production_bus_claimed"`
	WriteMethodsEnabled           bool                                    `json:"write_methods_enabled"`
	SupportBundleExported         bool                                    `json:"support_bundle_exported"`
	SupportCaseCreated            bool                                    `json:"support_case_created"`
	NotificationSent              bool                                    `json:"notification_sent"`
	BackendLaunchEnabled          bool                                    `json:"backend_launch_enabled"`
	BackendProcessStarted         bool                                    `json:"backend_process_started"`
	NetworkRequired               bool                                    `json:"network_required"`
	HostRootModified              bool                                    `json:"host_root_modified"`
	PrivilegedContainerRequired   bool                                    `json:"privileged_container_required"`
	BackendDetailsExposed         bool                                    `json:"backend_details_exposed"`
	BlockedActions                []string                                `json:"blocked_actions"`
	NextRequirements              []string                                `json:"next_requirements"`
	DesktopSafeSummary            string                                  `json:"desktop_safe_summary"`
}

type RestrictedOwnerSmokeReceiptFanSurface struct {
	ID                    string `json:"id"`
	Consumer              string `json:"consumer"`
	RuntimeMethod         string `json:"runtime_method"`
	ReadModel             string `json:"read_model"`
	EvidenceState         string `json:"evidence_state"`
	ReceiptRelativePath   string `json:"receipt_relative_path"`
	ReceiptSHA256         string `json:"receipt_sha256"`
	ReadOnly              bool   `json:"read_only"`
	UserVisible           bool   `json:"user_visible"`
	ConsumesReceipt       bool   `json:"consumes_receipt"`
	MutatesRuntime        bool   `json:"mutates_runtime"`
	StartsService         bool   `json:"starts_service"`
	ClaimsSessionBus      bool   `json:"claims_session_bus"`
	ClaimsProductionBus   bool   `json:"claims_production_bus"`
	EnablesWriteMethods   bool   `json:"enables_write_methods"`
	StartsBackend         bool   `json:"starts_backend"`
	ExportsSupportBundle  bool   `json:"exports_support_bundle"`
	CreatesSupportCase    bool   `json:"creates_support_case"`
	SendsNotification     bool   `json:"sends_notification"`
	ExposesStateRootPath  bool   `json:"exposes_state_root_path"`
	ExposesBackendDetails bool   `json:"exposes_backend_details"`
	Summary               string `json:"summary"`
}

func NewRestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview(root string, opaqueReceiptID string) (RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview, error) {
	lookup, err := ResolveRestrictedOwnerSmokeReceipt(root, opaqueReceiptID)
	if err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview{}, err
	}

	ownerReadiness := restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface("runtime-owner-readiness", "Runtime owner readiness", "GetRuntimeOwnerReadiness", "runtime-owner-readiness-preview", lookup, "Owner readiness can show where restricted smoke evidence would attach while missing receipts remain blocked.")
	servicePreflight := restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface("service-activation-preflight", "Runtime service activation preflight", "GetRuntimeServiceActivationPreflight", "runtime-service-activation-preflight-preview", lookup, "Service activation preflight can point to the owner-managed receipt slot without starting a service.")
	onboarding := restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface("compatibility-onboarding", "Compatibility onboarding checklist", "GetCompatibilityOnboardingChecklist", "compatibility-onboarding-checklist-preview", lookup, "Onboarding can explain that restricted smoke evidence is not yet available for readiness.")
	supportBundle := restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface("support-bundle-manifest", "Support bundle manifest", "GetSupportBundleManifest", "support-bundle-manifest-preview", lookup, "Support bundle previews can cite a missing receipt state without exporting a bundle.")
	supportCase := restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface("support-case-timeline", "Support case timeline", "GetSupportCaseTimeline", "support-case-timeline-preview", lookup, "Support timelines can keep the restricted smoke checkpoint absent without creating a case or notification.")
	surfaces := []RestrictedOwnerSmokeReceiptFanSurface{ownerReadiness, servicePreflight, onboarding, supportBundle, supportCase}

	preview := RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview{
		Version:                       lookup.Version,
		SchemaVersion:                 "xnix.runtime.restricted_owner_smoke_receipt_fanout_owner_route.v1",
		RequestType:                   "restricted-owner-smoke-receipt-fanout-owner-route-preview",
		RouteType:                     "owner-local-restricted-smoke-receipt-fanout",
		Source:                        "runtime-owner-dispatch+opaque-receipt-id-registry+restricted-owner-smoke-receipt-fanout",
		RuntimeMethod:                 "GetRestrictedOwnerSmokeReceiptFanOut",
		ReadMethod:                    "GetRestrictedOwnerSmokeReceiptFanOutPreview",
		OpaqueReceiptID:               lookup.OpaqueReceiptID,
		SupportedOpaqueReceiptID:      lookup.SupportedOpaqueReceiptID,
		ReceiptID:                     lookup.ReceiptID,
		ReceiptRelativePath:           lookup.ReceiptRelativePath,
		ReceiptRecordType:             lookup.ReceiptRecordType,
		ReceiptLookupState:            lookup.ReceiptLookupState,
		ReceiptAvailable:              lookup.ReceiptAvailable,
		ReceiptConsumed:               false,
		MissingReceiptSafe:            lookup.MissingReceiptSafe,
		OwnerManagedLookup:            lookup.OwnerManagedLookup,
		CallerStateRootRequired:       false,
		OpaqueReceiptIDSupported:      lookup.OpaqueReceiptIDSupported,
		ReadOnlyFanOut:                true,
		FanOutResultState:             "missing-receipt-fail-closed",
		OwnerLocalRouteCandidateReady: true,
		ProductionDBusExposureReady:   false,
		SurfaceCount:                  len(surfaces),
		Surfaces:                      surfaces,
		RuntimeOwnerReadiness:         ownerReadiness,
		ServiceActivationPreflight:    servicePreflight,
		CompatibilityOnboarding:       onboarding,
		SupportBundleManifest:         supportBundle,
		SupportCaseTimeline:           supportCase,
		RuntimeOwned:                  true,
		GoRuntimeBacked:               true,
		KDEPolicyOwner:                false,
		ReadinessSurfacesSatisfied:    false,
		SupportSurfacesSatisfied:      false,
		StateRootPathExposed:          false,
		StateRootWritesEnabled:        false,
		RuntimeWritesEnabled:          false,
		FanOutWritesEnabled:           false,
		ProductionActivationReady:     false,
		ProductionOwnerEnabled:        false,
		SystemServiceStarted:          false,
		SessionBusClaimed:             false,
		ProductionBusClaimed:          false,
		WriteMethodsEnabled:           false,
		SupportBundleExported:         false,
		SupportCaseCreated:            false,
		NotificationSent:              false,
		BackendLaunchEnabled:          false,
		BackendProcessStarted:         false,
		NetworkRequired:               false,
		HostRootModified:              false,
		PrivilegedContainerRequired:   false,
		BackendDetailsExposed:         false,
		BlockedActions: []string{
			"accept caller-supplied state-root paths through the owner fan-out route",
			"treat a missing restricted owner smoke receipt as readiness evidence",
			"export support bundles, create support cases, or send notifications from fan-out",
			"claim production D-Bus ownership or enable Runtime writes from owner-local fan-out",
			"start compatibility services, launch, use network, or mutate host root from fan-out",
		},
		NextRequirements: []string{
			"Attach a digest-verified restricted owner smoke receipt to the owner-managed receipt slot.",
			"Keep owner-local route coverage before considering production D-Bus exposure.",
			"Promote support and readiness surfaces only after the receipt is available and verified.",
			"Continue reporting missing receipts as fail-closed evidence.",
		},
		DesktopSafeSummary: "Runtime owner can route restricted smoke receipt fan-out through an opaque receipt id without caller state-root paths, but missing receipt evidence remains fail-closed and does not enable production ownership, writes, support side effects, launch, or host mutation.",
	}
	checks := restrictedOwnerSmokeReceiptFanOutOwnerRouteChecks(preview)
	preview.Checks = checks
	preview.CheckIDs = restrictedOwnerSmokeCheckIDs(checks)
	preview.CheckCount = len(checks)
	preview.PassedCheckCount = countRestrictedOwnerSmokePasses(checks)
	preview.AllChecksPassed = preview.PassedCheckCount == preview.CheckCount
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out owner route preview"); err != nil {
		return RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview{}, err
	}
	return preview, nil
}

func NewRestrictedOwnerSmokeReceiptFanOutPreview(stateRoot string) (RestrictedOwnerSmokeReceiptFanOutPreview, error) {
	if stateRoot == "" {
		return RestrictedOwnerSmokeReceiptFanOutPreview{}, errors.New("restricted owner smoke receipt fan-out requires an explicit state root")
	}
	receipt, err := LoadRestrictedOwnerSmokeReceipt(stateRoot)
	if err != nil {
		return RestrictedOwnerSmokeReceiptFanOutPreview{}, err
	}
	if err := validateRestrictedOwnerSmokeReceiptFanOutReceipt(receipt); err != nil {
		return RestrictedOwnerSmokeReceiptFanOutPreview{}, err
	}

	ownerReadiness := restrictedOwnerSmokeReceiptFanSurface("runtime-owner-readiness", "Runtime owner readiness", "GetRuntimeOwnerReadiness", "runtime-owner-readiness-preview", receipt, "Owner readiness can cite digest-verified restricted smoke evidence without accepting it as production ownership.")
	servicePreflight := restrictedOwnerSmokeReceiptFanSurface("service-activation-preflight", "Runtime service activation preflight", "GetRuntimeServiceActivationPreflight", "runtime-service-activation-preflight-preview", receipt, "Service activation preflight can show restricted smoke evidence while production service start remains blocked.")
	onboarding := restrictedOwnerSmokeReceiptFanSurface("compatibility-onboarding", "Compatibility onboarding checklist", "GetCompatibilityOnboardingChecklist", "compatibility-onboarding-checklist-preview", receipt, "Onboarding can explain that restricted owner smoke evidence exists while install and launch writes remain unavailable.")
	supportBundle := restrictedOwnerSmokeReceiptFanSurface("support-bundle-manifest", "Support bundle manifest", "GetSupportBundleManifest", "support-bundle-manifest-preview", receipt, "Support bundle previews can reference the receipt digest without exporting a bundle or exposing local paths.")
	supportCase := restrictedOwnerSmokeReceiptFanSurface("support-case-timeline", "Support case timeline", "GetSupportCaseTimeline", "support-case-timeline-preview", receipt, "Support timelines can include the restricted smoke checkpoint without creating a ticket or sending a notification.")
	surfaces := []RestrictedOwnerSmokeReceiptFanSurface{ownerReadiness, servicePreflight, onboarding, supportBundle, supportCase}
	checks := restrictedOwnerSmokeReceiptFanOutChecks(receipt, surfaces)
	passed := countRestrictedOwnerSmokePasses(checks)
	if passed != len(checks) {
		return RestrictedOwnerSmokeReceiptFanOutPreview{}, errors.New("restricted owner smoke receipt fan-out checks did not all pass")
	}

	preview := RestrictedOwnerSmokeReceiptFanOutPreview{
		Version:                      receipt.Version,
		SchemaVersion:                "xnix.runtime.restricted_owner_smoke_receipt_fanout.v1",
		RequestType:                  "restricted-owner-smoke-receipt-fanout-preview",
		Source:                       "restricted-owner-smoke-execution-receipt",
		RuntimeMethod:                "GetRestrictedOwnerSmokeReceiptFanOut",
		ReadMethod:                   "GetRestrictedOwnerSmokeReceiptFanOutPreview",
		Mode:                         receipt.Mode,
		ReceiptID:                    receipt.ReceiptID,
		ReceiptRelativePath:          receipt.RelativePath,
		ReceiptSHA256:                receipt.SHA256,
		ReceiptRecordType:            receipt.RecordType,
		ReceiptReadBack:              receipt.ReceiptReadBack,
		ReceiptConsumed:              true,
		ReceiptAllChecksPassed:       receipt.AllChecksPassed,
		ReceiptCheckCount:            receipt.CheckCount,
		SmokeReadDispatchRecordCount: receipt.SmokeBatch.ReadDispatchRecordCount,
		SmokeWriteDenialRecordCount:  receipt.SmokeBatch.WriteDenialRecordCount,
		SurfaceCount:                 len(surfaces),
		Surfaces:                     surfaces,
		RuntimeOwnerReadiness:        ownerReadiness,
		ServiceActivationPreflight:   servicePreflight,
		CompatibilityOnboarding:      onboarding,
		SupportBundleManifest:        supportBundle,
		SupportCaseTimeline:          supportCase,
		Checks:                       checks,
		CheckCount:                   len(checks),
		PassedCheckCount:             passed,
		AllChecksPassed:              true,
		RuntimeOwned:                 true,
		GoRuntimeBacked:              true,
		KDEPolicyOwner:               false,
		ReadOnlyFanOut:               true,
		ReadinessSurfacesSatisfied:   true,
		SupportSurfacesSatisfied:     true,
		StateRootPathExposed:         false,
		StateRootWritesEnabled:       false,
		FanOutWritesEnabled:          false,
		ProductionActivationReady:    false,
		ProductionOwnerEnabled:       false,
		SystemServiceStarted:         false,
		SessionBusClaimed:            false,
		ProductionBusClaimed:         false,
		WriteMethodsEnabled:          false,
		SupportBundleExported:        false,
		SupportCaseCreated:           false,
		NotificationSent:             false,
		BackendLaunchEnabled:         false,
		NetworkRequired:              false,
		HostRootModified:             false,
		PrivilegedContainerRequired:  false,
		BackendDetailsExposed:        false,
		DesktopSafeSummary:           "Runtime readiness and support previews can consume a digest-verified restricted owner smoke receipt while production service start, D-Bus ownership, Runtime writes, support export, ticket creation, backend launch, and host mutation remain disabled.",
	}
	if err := validateNoBackendTerms(preview, "restricted owner smoke receipt fan-out preview"); err != nil {
		return RestrictedOwnerSmokeReceiptFanOutPreview{}, err
	}
	return preview, nil
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteSurface(id string, consumer string, runtimeMethod string, readModel string, lookup RestrictedOwnerSmokeReceiptLookupPreview, summary string) RestrictedOwnerSmokeReceiptFanSurface {
	return RestrictedOwnerSmokeReceiptFanSurface{
		ID:                    id,
		Consumer:              consumer,
		RuntimeMethod:         runtimeMethod,
		ReadModel:             readModel,
		EvidenceState:         lookup.ReceiptLookupState,
		ReceiptRelativePath:   lookup.ReceiptRelativePath,
		ReceiptSHA256:         "",
		ReadOnly:              true,
		UserVisible:           true,
		ConsumesReceipt:       false,
		MutatesRuntime:        false,
		StartsService:         false,
		ClaimsSessionBus:      false,
		ClaimsProductionBus:   false,
		EnablesWriteMethods:   false,
		StartsBackend:         false,
		ExportsSupportBundle:  false,
		CreatesSupportCase:    false,
		SendsNotification:     false,
		ExposesStateRootPath:  false,
		ExposesBackendDetails: false,
		Summary:               summary,
	}
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteChecks(preview RestrictedOwnerSmokeReceiptFanOutOwnerRoutePreview) []RestrictedOwnerSmokeCheck {
	return []RestrictedOwnerSmokeCheck{
		restrictedOwnerSmokeCheck("opaque-lookup-consumed", preview.OwnerManagedLookup && preview.OpaqueReceiptIDSupported && preview.OpaqueReceiptID == RestrictedOwnerSmokeOpaqueReceiptID, "Owner-local fan-out consumes the opaque receipt lookup result instead of caller paths."),
		restrictedOwnerSmokeCheck("caller-state-root-hidden", !preview.CallerStateRootRequired && !preview.StateRootPathExposed && !filepath.IsAbs(preview.ReceiptRelativePath), "Owner-local fan-out does not accept or expose caller state-root paths."),
		restrictedOwnerSmokeCheck("missing-receipt-fails-closed", !preview.ReceiptAvailable && !preview.ReceiptConsumed && preview.MissingReceiptSafe && preview.FanOutResultState == "missing-receipt-fail-closed", "Missing receipt evidence stays blocked instead of satisfying readiness or support surfaces."),
		restrictedOwnerSmokeCheck("surface-fanout-deferred", restrictedOwnerSmokeReceiptFanOutOwnerRouteSurfacesDeferred(preview.Surfaces), "All readiness and support fan-out surfaces are present but do not consume a missing receipt."),
		restrictedOwnerSmokeCheck("support-side-effects-disabled", restrictedOwnerSmokeReceiptFanSupportSideEffectsDisabled(preview.Surfaces) && !preview.SupportBundleExported && !preview.SupportCaseCreated && !preview.NotificationSent, "Owner-local fan-out does not export bundles, create cases, or send notifications."),
		restrictedOwnerSmokeCheck("owner-route-ready-production-dbus-blocked", preview.OwnerLocalRouteCandidateReady && !preview.ProductionDBusExposureReady && !preview.ProductionBusClaimed, "The owner-local read route is ready while production D-Bus exposure remains blocked."),
		restrictedOwnerSmokeCheck("unsafe-gates-closed", !preview.StateRootWritesEnabled && !preview.RuntimeWritesEnabled && !preview.FanOutWritesEnabled && !preview.ProductionActivationReady && !preview.ProductionOwnerEnabled && !preview.SystemServiceStarted && !preview.SessionBusClaimed && !preview.ProductionBusClaimed && !preview.WriteMethodsEnabled && !preview.BackendLaunchEnabled && !preview.BackendProcessStarted && !preview.NetworkRequired && !preview.HostRootModified && !preview.PrivilegedContainerRequired && !preview.BackendDetailsExposed, "Owner-local fan-out keeps writes, ownership, launch, network, privilege, and host mutation disabled."),
	}
}

func restrictedOwnerSmokeReceiptFanOutOwnerRouteSurfacesDeferred(surfaces []RestrictedOwnerSmokeReceiptFanSurface) bool {
	if len(surfaces) != 5 {
		return false
	}
	for _, surface := range surfaces {
		if !surface.ReadOnly || !surface.UserVisible || surface.ConsumesReceipt || surface.MutatesRuntime || surface.StartsService || surface.ClaimsSessionBus || surface.ClaimsProductionBus || surface.EnablesWriteMethods || surface.StartsBackend || surface.ExposesStateRootPath || surface.ExposesBackendDetails || filepath.IsAbs(surface.ReceiptRelativePath) {
			return false
		}
		if surface.EvidenceState != "missing-receipt" {
			return false
		}
	}
	return true
}

func validateRestrictedOwnerSmokeReceiptFanOutReceipt(receipt RestrictedOwnerSmokeReceipt) error {
	if !receipt.ReceiptReadBack || !receipt.AllChecksPassed || !receipt.RestrictedSmokeReady || !receipt.RestrictedSmokeAuthorized {
		return errors.New("restricted owner smoke receipt fan-out requires a verified restricted smoke receipt")
	}
	if filepath.IsAbs(receipt.RelativePath) || receipt.SHA256 == "" {
		return errors.New("restricted owner smoke receipt fan-out requires relative receipt evidence and a digest")
	}
	if receipt.ProductionActivationReady || receipt.ProductionOwnerEnabled || receipt.SystemServiceStarted || receipt.SessionBusClaimed || receipt.ProductionBusClaimed || receipt.WriteMethodsEnabled || receipt.BackendLaunchEnabled || receipt.NetworkRequired || receipt.HostRootModified || receipt.PrivilegedContainerRequired || receipt.BackendDetailsExposed || receipt.StateRootPathExposed {
		return errors.New("restricted owner smoke receipt fan-out received an unsafe receipt")
	}
	return nil
}

func restrictedOwnerSmokeReceiptFanSurface(id string, consumer string, runtimeMethod string, readModel string, receipt RestrictedOwnerSmokeReceipt, summary string) RestrictedOwnerSmokeReceiptFanSurface {
	return RestrictedOwnerSmokeReceiptFanSurface{
		ID:                    id,
		Consumer:              consumer,
		RuntimeMethod:         runtimeMethod,
		ReadModel:             readModel,
		EvidenceState:         "restricted-smoke-receipt-ready",
		ReceiptRelativePath:   receipt.RelativePath,
		ReceiptSHA256:         receipt.SHA256,
		ReadOnly:              true,
		UserVisible:           true,
		ConsumesReceipt:       true,
		MutatesRuntime:        false,
		StartsService:         false,
		ClaimsSessionBus:      false,
		ClaimsProductionBus:   false,
		EnablesWriteMethods:   false,
		StartsBackend:         false,
		ExportsSupportBundle:  false,
		CreatesSupportCase:    false,
		SendsNotification:     false,
		ExposesStateRootPath:  false,
		ExposesBackendDetails: false,
		Summary:               summary,
	}
}

func restrictedOwnerSmokeReceiptFanOutChecks(receipt RestrictedOwnerSmokeReceipt, surfaces []RestrictedOwnerSmokeReceiptFanSurface) []RestrictedOwnerSmokeCheck {
	return []RestrictedOwnerSmokeCheck{
		restrictedOwnerSmokeCheck("receipt-consumed", receipt.ReceiptReadBack && receipt.AllChecksPassed && !filepath.IsAbs(receipt.RelativePath) && len(receipt.SHA256) == 64, "Fan-out consumed a digest-verified restricted owner smoke receipt."),
		restrictedOwnerSmokeCheck("readiness-surface-coverage", restrictedOwnerSmokeReceiptFanHasSurface(surfaces, "runtime-owner-readiness") && restrictedOwnerSmokeReceiptFanHasSurface(surfaces, "service-activation-preflight") && restrictedOwnerSmokeReceiptFanHasSurface(surfaces, "compatibility-onboarding"), "Owner readiness, service activation preflight, and onboarding surfaces consume the receipt."),
		restrictedOwnerSmokeCheck("support-surface-coverage", restrictedOwnerSmokeReceiptFanHasSurface(surfaces, "support-bundle-manifest") && restrictedOwnerSmokeReceiptFanHasSurface(surfaces, "support-case-timeline"), "Support bundle and support case surfaces consume the receipt."),
		restrictedOwnerSmokeCheck("surfaces-read-only", restrictedOwnerSmokeReceiptFanSurfacesReadOnly(surfaces), "All fan-out surfaces are read-only receipt projections."),
		restrictedOwnerSmokeCheck("ownership-boundary-closed", !receipt.ProductionActivationReady && !receipt.ProductionOwnerEnabled && !receipt.SystemServiceStarted && !receipt.SessionBusClaimed && !receipt.ProductionBusClaimed, "Restricted smoke evidence is not accepted as production service ownership."),
		restrictedOwnerSmokeCheck("support-side-effects-disabled", restrictedOwnerSmokeReceiptFanSupportSideEffectsDisabled(surfaces), "Support fan-out does not export bundles, create cases, or send notifications."),
		restrictedOwnerSmokeCheck("unsafe-data-hidden", !receipt.StateRootPathExposed && !receipt.BackendDetailsExposed && restrictedOwnerSmokeReceiptFanHidesUnsafeData(surfaces), "State-root paths and implementation details stay hidden."),
		restrictedOwnerSmokeCheck("host-boundary-closed", !receipt.BackendLaunchEnabled && !receipt.NetworkRequired && !receipt.HostRootModified && !receipt.PrivilegedContainerRequired, "Fan-out keeps launch, network, privileged-container, and host-root mutation gates closed."),
	}
}

func restrictedOwnerSmokeReceiptFanHasSurface(surfaces []RestrictedOwnerSmokeReceiptFanSurface, id string) bool {
	for _, surface := range surfaces {
		if surface.ID == id {
			return true
		}
	}
	return false
}

func restrictedOwnerSmokeReceiptFanSurfacesReadOnly(surfaces []RestrictedOwnerSmokeReceiptFanSurface) bool {
	if len(surfaces) != 5 {
		return false
	}
	for _, surface := range surfaces {
		if !surface.ReadOnly || !surface.UserVisible || !surface.ConsumesReceipt || surface.MutatesRuntime || surface.StartsService || surface.ClaimsSessionBus || surface.ClaimsProductionBus || surface.EnablesWriteMethods || surface.StartsBackend {
			return false
		}
	}
	return true
}

func restrictedOwnerSmokeReceiptFanSupportSideEffectsDisabled(surfaces []RestrictedOwnerSmokeReceiptFanSurface) bool {
	for _, surface := range surfaces {
		if surface.ExportsSupportBundle || surface.CreatesSupportCase || surface.SendsNotification {
			return false
		}
	}
	return true
}

func restrictedOwnerSmokeReceiptFanHidesUnsafeData(surfaces []RestrictedOwnerSmokeReceiptFanSurface) bool {
	for _, surface := range surfaces {
		if filepath.IsAbs(surface.ReceiptRelativePath) || surface.ExposesStateRootPath || surface.ExposesBackendDetails {
			return false
		}
	}
	return true
}
