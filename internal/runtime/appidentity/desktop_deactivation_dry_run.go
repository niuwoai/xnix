package appidentity

import (
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

// DesktopDeactivationDryRunOptions carries the activation receipt root plus the
// caller-resolved conditions a dry run cannot cheaply derive on its own (an
// active session, or MIME types shared with other applications).
type DesktopDeactivationDryRunOptions struct {
	ActivationRoot         string
	ActiveSession          bool
	SharedMIMEAssociations []string
}

type DesktopDeactivationDryRunPreview struct {
	SchemaVersion              string                       `json:"schema_version"`
	RequestType                string                       `json:"request_type"`
	PlanType                   string                       `json:"plan_type"`
	Source                     string                       `json:"source"`
	Desktop                    string                       `json:"desktop"`
	RuntimeMethod              string                       `json:"runtime_method"`
	ReadMethod                 string                       `json:"read_method"`
	ApplicationID              string                       `json:"application_id"`
	DisplayName                string                       `json:"display_name"`
	OverallState               string                       `json:"overall_state"`
	RollbackSafetyStatus       string                       `json:"rollback_safety_status"`
	ReceiptState               string                       `json:"receipt_state"`
	ReceiptRelativePath        string                       `json:"receipt_relative_path,omitempty"`
	DigestVerificationRequired bool                         `json:"digest_verification_required"`
	DigestGateReady            bool                         `json:"digest_gate_ready"`
	Surfaces                   []DesktopDeactivationSurface `json:"surfaces"`
	SurfaceIDs                 []string                     `json:"surface_ids"`
	SurfaceCount               int                          `json:"surface_count"`
	RemovalSteps               []DesktopDeactivationStep    `json:"removal_steps"`
	Counts                     DesktopDeactivationCounts    `json:"counts"`
	BlockedReasons             []string                     `json:"blocked_reasons"`
	EvidenceIDs                []string                     `json:"evidence_ids"`
	NextSafeReadOnlyChecks     []string                     `json:"next_safe_read_only_checks"`
	RuntimeOwned               bool                         `json:"runtime_owned"`
	GoRuntimeBacked            bool                         `json:"go_runtime_backed"`
	KDEPolicyOwner             bool                         `json:"kde_policy_owner"`
	UserVisible                bool                         `json:"user_visible"`
	ReviewOnly                 bool                         `json:"review_only"`
	FileDeletionEnabled        bool                         `json:"file_deletion_enabled"`
	MIMEDefaultsWritten        bool                         `json:"mime_defaults_written"`
	KDECacheRefreshed          bool                         `json:"kde_cache_refreshed"`
	ReceiptsRewritten          bool                         `json:"receipts_rewritten"`
	SessionTerminated          bool                         `json:"session_terminated"`
	BackendLaunchEnabled       bool                         `json:"backend_launch_enabled"`
	TargetPathExposed          bool                         `json:"target_path_exposed"`
	HostRootModified           bool                         `json:"host_root_modified"`
	BackendDetailsExposed      bool                         `json:"backend_details_exposed"`
	BlockedActions             []string                     `json:"blocked_actions"`
	DesktopSafeSummary         string                       `json:"desktop_safe_summary"`
}

type DesktopDeactivationSurface struct {
	ID                    string   `json:"id"`
	Label                 string   `json:"label"`
	HasRemovableArtifact  bool     `json:"has_removable_artifact"`
	DeactivationState     string   `json:"deactivation_state"`
	RequiresReceipt       bool     `json:"requires_receipt"`
	RequiresDigestMatch   bool     `json:"requires_digest_match"`
	RollbackSafe          bool     `json:"rollback_safe"`
	BlockedReasons        []string `json:"blocked_reasons"`
	EvidenceIDs           []string `json:"evidence_ids"`
	UserSafeSummary       string   `json:"user_safe_summary"`
	FileDeletionEnabled   bool     `json:"file_deletion_enabled"`
	TargetPathExposed     bool     `json:"target_path_exposed"`
	HostRootModified      bool     `json:"host_root_modified"`
	BackendDetailsExposed bool     `json:"backend_details_exposed"`
}

type DesktopDeactivationStep struct {
	ID                  string `json:"id"`
	Status              string `json:"status"`
	Required            bool   `json:"required"`
	RequiresReceipt     bool   `json:"requires_receipt"`
	RequiresDigestMatch bool   `json:"requires_digest_match"`
	Summary             string `json:"summary"`
}

type DesktopDeactivationCounts struct {
	TotalSurfaces  int `json:"total_surfaces"`
	Removable      int `json:"removable"`
	Blocked        int `json:"blocked"`
	NoHostArtifact int `json:"no_host_artifact"`
}

type desktopDeactivationSurfaceSpec struct {
	id          string
	label       string
	hasArtifact bool
}

func desktopDeactivationSurfaceSpecs() []desktopDeactivationSurfaceSpec {
	return []desktopDeactivationSurfaceSpec{
		{"launcher", "KDE launcher", true},
		{"desktop-icon", "Desktop icon", true},
		{"mime-association", "MIME association", true},
		{"dolphin-service-menu", "Dolphin service menu", true},
		{"kwin", "KWin window rule", false},
		{"system-tray", "System tray", false},
		{"notification", "Notification route", false},
		{"compatibility-center", "Compatibility Center", false},
		{"settings", "Settings surface", false},
	}
}

// DesktopDeactivationDryRunPreview explains how a managed application would be
// removed from KDE entry points without deleting files, changing host MIME
// associations, refreshing caches, or ending sessions. It reports per-surface
// removability, receipt and digest requirements, rollback safety, and blocked
// reasons for missing receipt, digest mismatch, unknown file owner, shared MIME
// association, and active session evidence.
func (plan Plan) DesktopDeactivationDryRunPreview(options DesktopDeactivationDryRunOptions) (DesktopDeactivationDryRunPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return DesktopDeactivationDryRunPreview{}, err
	}

	receiptState := "missing"
	receiptRelativePath := ""
	digestGateReady := false
	receiptBlockedReason := "missing-activation-receipt"
	if strings.TrimSpace(options.ActivationRoot) != "" {
		evidence, err := plan.DesktopActivationReceiptEvidence(options.ActivationRoot)
		if err != nil {
			receiptState = "unverified"
			receiptBlockedReason = classifyDeactivationReceiptError(err)
		} else {
			receiptState = "receipt-backed"
			receiptRelativePath = evidence.ReceiptRelativePath
			digestGateReady = evidence.DigestGateReady
			receiptBlockedReason = ""
		}
	}

	transaction, err := plan.DesktopActivationTransactionPreview("development")
	if err != nil {
		return DesktopDeactivationDryRunPreview{}, err
	}

	var surfaces []DesktopDeactivationSurface
	var counts DesktopDeactivationCounts
	for _, spec := range desktopDeactivationSurfaceSpecs() {
		surface := plan.desktopDeactivationSurface(spec, options, receiptState, receiptBlockedReason, digestGateReady, receiptRelativePath)
		counts.TotalSurfaces++
		switch surface.DeactivationState {
		case "removable":
			counts.Removable++
		case "no-host-artifact":
			counts.NoHostArtifact++
		default:
			counts.Blocked++
		}
		surfaces = append(surfaces, surface)
	}

	overallState := "removable"
	rollbackSafety := "safe"
	if counts.Blocked > 0 {
		overallState = "blocked"
		rollbackSafety = "blocked"
	}

	preview := DesktopDeactivationDryRunPreview{
		SchemaVersion:              "xnix.runtime.desktop_deactivation_dry_run.v1",
		RequestType:                "desktop-deactivation-dry-run-preview",
		PlanType:                   "desktop-deactivation-dry-run",
		Source:                     "desktop-activation-transaction+desktop-activation-receipt+kde-journey-evidence",
		Desktop:                    "KDE Plasma",
		RuntimeMethod:              "GetDesktopDeactivationDryRun",
		ReadMethod:                 "GetDesktopDeactivationDryRunPreview",
		ApplicationID:              plan.ApplicationID,
		DisplayName:                plan.DisplayName,
		OverallState:               overallState,
		RollbackSafetyStatus:       rollbackSafety,
		ReceiptState:               receiptState,
		ReceiptRelativePath:        receiptRelativePath,
		DigestVerificationRequired: true,
		DigestGateReady:            digestGateReady,
		Surfaces:                   surfaces,
		SurfaceIDs:                 desktopDeactivationSurfaceIDs(surfaces),
		SurfaceCount:               len(surfaces),
		RemovalSteps:               desktopDeactivationRemovalSteps(transaction.RollbackSteps),
		Counts:                     counts,
		BlockedReasons:             desktopDeactivationAggregateReasons(surfaces),
		EvidenceIDs:                desktopDeactivationEvidenceIDs(surfaces, receiptRelativePath),
		NextSafeReadOnlyChecks: []string{
			"review the desktop activation receipt evidence",
			"review the activation transaction rollback steps",
			"review shared MIME association ownership before removal",
		},
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		UserVisible:           true,
		ReviewOnly:            true,
		FileDeletionEnabled:   false,
		MIMEDefaultsWritten:   false,
		KDECacheRefreshed:     false,
		ReceiptsRewritten:     false,
		SessionTerminated:     false,
		BackendLaunchEnabled:  false,
		TargetPathExposed:     false,
		HostRootModified:      false,
		BackendDetailsExposed: false,
		BlockedActions: []string{
			"delete a staged desktop file from the deactivation dry run",
			"write a MIME default from the deactivation dry run",
			"refresh the KDE service cache from the deactivation dry run",
			"rewrite an activation receipt from the deactivation dry run",
			"terminate an application session from the deactivation dry run",
			"mutate host root during the deactivation dry run",
		},
		DesktopSafeSummary: desktopDeactivationSummary(overallState, receiptState),
	}
	if err := validateNoBackendTerms(preview, "desktop deactivation dry run preview"); err != nil {
		return DesktopDeactivationDryRunPreview{}, err
	}
	if err := safety.ValidatePayload("desktop deactivation dry run preview", preview); err != nil {
		return DesktopDeactivationDryRunPreview{}, err
	}
	return preview, nil
}

func (plan Plan) desktopDeactivationSurface(spec desktopDeactivationSurfaceSpec, options DesktopDeactivationDryRunOptions, receiptState, receiptBlockedReason string, digestGateReady bool, receiptRelativePath string) DesktopDeactivationSurface {
	surface := DesktopDeactivationSurface{
		ID:                   spec.id,
		Label:                spec.label,
		HasRemovableArtifact: spec.hasArtifact,
		RequiresReceipt:      spec.hasArtifact,
		RequiresDigestMatch:  spec.hasArtifact,
	}
	if !spec.hasArtifact {
		surface.DeactivationState = "no-host-artifact"
		surface.RollbackSafe = true
		surface.UserSafeSummary = spec.label + " has no host artifact to remove; it clears when the Runtime read model stops reporting the application."
		return surface
	}

	surface.EvidenceIDs = []string{"activation-surface:" + spec.id}
	if receiptRelativePath != "" {
		surface.EvidenceIDs = append(surface.EvidenceIDs, "activation-receipt:"+receiptRelativePath)
	}

	var reasons []string
	switch {
	case receiptState != "receipt-backed":
		surface.DeactivationState = "blocked"
		reasons = append(reasons, receiptBlockedReason)
	case !digestGateReady:
		surface.DeactivationState = "blocked"
		reasons = append(reasons, "installed-file-digest-mismatch")
	case options.ActiveSession:
		surface.DeactivationState = "blocked"
		reasons = append(reasons, "active-session-evidence")
	case spec.id == "mime-association" && len(options.SharedMIMEAssociations) > 0:
		surface.DeactivationState = "blocked"
		reasons = append(reasons, "shared-mime-association")
	default:
		surface.DeactivationState = "removable"
		surface.RollbackSafe = true
	}
	surface.BlockedReasons = desktopDeactivationUnique(reasons)
	surface.UserSafeSummary = desktopDeactivationSurfaceSummary(surface.DeactivationState, spec.label)
	return surface
}

// classifyDeactivationReceiptError maps the activation receipt reader's errors
// to stable, user-safe blocked-reason codes without echoing the error text.
func classifyDeactivationReceiptError(err error) string {
	text := strings.ToLower(err.Error())
	switch {
	case strings.Contains(text, "application mismatch"):
		return "unknown-file-owner"
	case strings.Contains(text, "sha256") || strings.Contains(text, "digest"):
		return "installed-file-digest-mismatch"
	case strings.Contains(text, "read desktop activation receipt") || strings.Contains(text, "no such file"):
		return "missing-activation-receipt"
	case strings.Contains(text, "not safe"):
		return "activation-receipt-unsafe"
	default:
		return "activation-receipt-unverified"
	}
}

func desktopDeactivationRemovalSteps(steps []DesktopActivationRollbackStep) []DesktopDeactivationStep {
	out := make([]DesktopDeactivationStep, 0, len(steps))
	for _, step := range steps {
		out = append(out, DesktopDeactivationStep{
			ID:                  step.ID,
			Status:              step.Status,
			Required:            step.Required,
			RequiresReceipt:     step.RequiresReceipt,
			RequiresDigestMatch: step.RequiresDigestMatch,
			Summary:             step.Summary,
		})
	}
	return out
}

func desktopDeactivationSurfaceSummary(state, label string) string {
	switch state {
	case "removable":
		return label + " can be removed after verifying the recorded file digest; nothing is deleted in this dry run."
	case "no-host-artifact":
		return label + " has no host artifact to remove."
	default:
		return label + " cannot be removed yet; the recorded evidence must be resolved first."
	}
}

func desktopDeactivationSummary(overallState, receiptState string) string {
	if overallState == "removable" {
		return "Every managed KDE surface can be removed after digest verification; this dry run deletes nothing and changes no host association."
	}
	if receiptState != "receipt-backed" {
		return "Deactivation is blocked because the activation receipt evidence is missing or unverified; this dry run deletes nothing."
	}
	return "Deactivation is blocked for some KDE surfaces pending review; this dry run deletes nothing and changes no host association."
}

func desktopDeactivationAggregateReasons(surfaces []DesktopDeactivationSurface) []string {
	var reasons []string
	for _, surface := range surfaces {
		reasons = append(reasons, surface.BlockedReasons...)
	}
	return desktopDeactivationUnique(reasons)
}

func desktopDeactivationEvidenceIDs(surfaces []DesktopDeactivationSurface, receiptRelativePath string) []string {
	var ids []string
	if receiptRelativePath != "" {
		ids = append(ids, "activation-receipt:"+receiptRelativePath)
	}
	for _, surface := range surfaces {
		ids = append(ids, surface.EvidenceIDs...)
	}
	return desktopDeactivationUnique(ids)
}

func desktopDeactivationSurfaceIDs(surfaces []DesktopDeactivationSurface) []string {
	ids := make([]string, 0, len(surfaces))
	for _, surface := range surfaces {
		ids = append(ids, surface.ID)
	}
	return ids
}

func desktopDeactivationUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}
