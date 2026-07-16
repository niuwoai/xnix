package appidentity

import (
	"errors"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

// PermissionEvidenceRowInput is the normalized, already-safe evidence the
// command layer resolves for one permission resource before the audit read
// model reconciles it. Every field is a stable identifier, never a host path,
// receipt path, secret, or backend detail.
type PermissionEvidenceRowInput struct {
	Resource      string
	SettingState  string // enabled | disabled | absent
	ReceiptState  string // granted | denied | expired | pending | absent
	ReviewState   string // approved | pending | blocked | absent
	BridgeExposed bool
	ReceiptRefs   []string
}

// PermissionEvidenceAuditInput carries the resolved evidence rows and any
// top-level context (malformed receipts, execution gating) for the audit.
type PermissionEvidenceAuditInput struct {
	ApplicationID     string
	Rows              []PermissionEvidenceRowInput
	MalformedReceipts []string
	ExecutionGated    bool
	ExecutionReason   string
}

type PermissionEvidenceAuditPreview struct {
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	AuditType                   string                        `json:"audit_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	RuntimeMethod               string                        `json:"runtime_method"`
	ReadMethod                  string                        `json:"read_method"`
	ApplicationID               string                        `json:"application_id"`
	OverallState                string                        `json:"overall_state"`
	Rows                        []PermissionEvidenceAuditRow  `json:"rows"`
	RowIDs                      []string                      `json:"row_ids"`
	RowCount                    int                           `json:"row_count"`
	Counts                      PermissionEvidenceAuditCounts `json:"counts"`
	MalformedReceiptIDs         []string                      `json:"malformed_receipt_ids"`
	ExecutionGated              bool                          `json:"execution_gated"`
	ExecutionGateReason         string                        `json:"execution_gate_reason,omitempty"`
	RemediationHints            []string                      `json:"remediation_hints"`
	NextSafeReadOnlyChecks      []string                      `json:"next_safe_read_only_checks"`
	RuntimeOwned                bool                          `json:"runtime_owned"`
	GoRuntimeBacked             bool                          `json:"go_runtime_backed"`
	KDEPolicyOwner              bool                          `json:"kde_policy_owner"`
	UserVisible                 bool                          `json:"user_visible"`
	ReviewOnly                  bool                          `json:"review_only"`
	RealPortalCallEnabled       bool                          `json:"real_portal_call_enabled"`
	PermissionGrantEnabled      bool                          `json:"permission_grant_enabled"`
	PermissionRevokeEnabled     bool                          `json:"permission_revoke_enabled"`
	ReceiptWriteEnabled         bool                          `json:"receipt_write_enabled"`
	SettingsPersisted           bool                          `json:"settings_persisted"`
	ExecutionApproved           bool                          `json:"execution_approved"`
	StateRootPathExposed        bool                          `json:"state_root_path_exposed"`
	HostRootModified            bool                          `json:"host_root_modified"`
	NetworkRequired             bool                          `json:"network_required"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	BackendDetailsExposed       bool                          `json:"backend_details_exposed"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type PermissionEvidenceAuditRow struct {
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	PortalInterface         string   `json:"portal_interface,omitempty"`
	PolicyDecision          string   `json:"policy_decision"`
	ConsistencyState        string   `json:"consistency_state"`
	SettingState            string   `json:"setting_state"`
	ReceiptState            string   `json:"receipt_state"`
	ReviewState             string   `json:"review_state"`
	BridgeExposed           bool     `json:"bridge_exposed"`
	ReceiptRefs             []string `json:"receipt_refs"`
	RemediationHint         string   `json:"remediation_hint"`
	ReviewOnly              bool     `json:"review_only"`
	PermissionGrantEnabled  bool     `json:"permission_grant_enabled"`
	PermissionRevokeEnabled bool     `json:"permission_revoke_enabled"`
	HostRootModified        bool     `json:"host_root_modified"`
}

type PermissionEvidenceAuditCounts struct {
	TotalRows         int `json:"total_rows"`
	Consistent        int `json:"consistent"`
	SettingOnly       int `json:"setting_only"`
	ReceiptOnly       int `json:"receipt_only"`
	Expired           int `json:"expired"`
	Denied            int `json:"denied"`
	MissingReview     int `json:"missing_review"`
	BlockedByPolicy   int `json:"blocked_by_policy"`
	Unsupported       int `json:"unsupported"`
	Divergent         int `json:"divergent"`
	MalformedReceipts int `json:"malformed_receipts"`
}

type permissionAuditResource struct {
	id              string
	title           string
	portalInterface string
	policyDecision  string // ask | deny | allow
	portalRequired  bool
}

// permissionAuditResourceOrder is the canonical, stable audit row order. The
// policy decisions and portal-required flags mirror the Runtime permission
// review plan and portal access policy: file, uri, print, clipboard, and
// screenshot access ask for user consent through a Portal request; camera and
// remote desktop stay denied until the user changes the application policy; and
// network is allowed without a Portal request.
func permissionAuditResourceOrder() []permissionAuditResource {
	return []permissionAuditResource{
		{"documents", "Documents", "org.freedesktop.portal.FileChooser", "ask", true},
		{"downloads", "Downloads", "org.freedesktop.portal.FileChooser", "ask", true},
		{"uris", "External links", "org.freedesktop.portal.OpenURI", "ask", true},
		{"print", "Printing", "org.freedesktop.portal.Print", "ask", true},
		{"clipboard", "Clipboard", "org.freedesktop.portal.Clipboard", "ask", true},
		{"screenshot", "Screenshot", "org.freedesktop.portal.Screenshot", "ask", true},
		{"camera", "Camera", "org.freedesktop.portal.Camera", "deny", true},
		{"remote-desktop", "Remote desktop", "org.freedesktop.portal.RemoteDesktop", "deny", true},
		{"network", "Network policy", "", "allow", false},
	}
}

// permissionAuditRowToOperation maps each canonical audit row to the Portal
// operation whose receipts and resource bridge describe it. The network row has
// no Portal operation.
func permissionAuditRowToOperation() map[string]string {
	return map[string]string{
		"documents":      "file-open",
		"downloads":      "file-open",
		"uris":           "uri-open",
		"print":          "print",
		"clipboard":      "clipboard",
		"screenshot":     "screenshot",
		"camera":         "camera",
		"remote-desktop": "remote-desktop",
		"network":        "network",
	}
}

// permissionAuditRowToReviewID maps each canonical audit row to the permission
// review entry id that carries its setting and review state. The uris and
// remote-desktop rows have no review entry.
func permissionAuditRowToReviewID() map[string]string {
	return map[string]string{
		"documents":  "documents",
		"downloads":  "downloads",
		"print":      "print",
		"clipboard":  "clipboard",
		"screenshot": "screenshot",
		"camera":     "camera",
		"network":    "network",
	}
}

// PermissionEvidenceAuditPreview assembles the offline permission audit for a
// recipe by joining the Runtime permission review plan, the KDE resource bridge
// plan, recorded Portal receipt states (already rolled up per operation by the
// caller), and execution preflight evidence. It changes no permission, writes no
// receipt, and starts no backend.
func (plan Plan) PermissionEvidenceAuditPreview(receiptStates map[string]string, receiptRefs map[string][]string, malformedReceipts []string) (PermissionEvidenceAuditPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return PermissionEvidenceAuditPreview{}, err
	}
	review, err := plan.PermissionReviewPreview()
	if err != nil {
		return PermissionEvidenceAuditPreview{}, err
	}
	bridge, err := plan.DesktopResourceBridgePreview()
	if err != nil {
		return PermissionEvidenceAuditPreview{}, err
	}

	reviewByID := map[string]PermissionReviewEntry{}
	for _, entry := range review.Permissions {
		reviewByID[entry.ID] = entry
	}
	bridgeByOperation := map[string]bool{}
	for _, resource := range bridge.Resources {
		bridgeByOperation[resource.Operation] = true
	}

	rowToOperation := permissionAuditRowToOperation()
	rowToReviewID := permissionAuditRowToReviewID()

	rows := make([]PermissionEvidenceRowInput, 0, len(permissionAuditResourceOrder()))
	for _, resource := range permissionAuditResourceOrder() {
		operation := rowToOperation[resource.id]
		row := PermissionEvidenceRowInput{Resource: resource.id}
		if reviewID := rowToReviewID[resource.id]; reviewID != "" {
			if entry, ok := reviewByID[reviewID]; ok {
				row.SettingState = permissionAuditSettingFromDecision(entry.CurrentValue)
				if entry.ChangePending {
					row.ReviewState = "pending"
				} else {
					row.ReviewState = "approved"
				}
			}
		}
		if row.SettingState == "" {
			row.SettingState = permissionAuditSettingFromDecision(resource.policyDecision)
			row.ReviewState = "absent"
		}
		if operation != "" {
			if state, ok := receiptStates[operation]; ok {
				row.ReceiptState = state
			}
			row.ReceiptRefs = receiptRefs[operation]
		}
		if operation != "" && bridgeByOperation[operation] {
			row.BridgeExposed = true
		}
		rows = append(rows, row)
	}

	input := PermissionEvidenceAuditInput{
		ApplicationID:     plan.ApplicationID,
		Rows:              rows,
		MalformedReceipts: malformedReceipts,
	}
	if preflight, err := plan.ExecutionPreflightPreview("reviewed", nil); err == nil {
		input.ExecutionGated = !preflight.LaunchAllowed
		if !preflight.LaunchAllowed {
			input.ExecutionReason = "Execution stays gated until Runtime preflight checks pass."
		}
	}
	return NewPermissionEvidenceAuditPreview(input)
}

func permissionAuditSettingFromDecision(decision string) string {
	switch strings.ToLower(strings.TrimSpace(decision)) {
	case "allow", "ask":
		return "enabled"
	case "deny":
		return "disabled"
	default:
		return "absent"
	}
}

func NewPermissionEvidenceAuditPreview(input PermissionEvidenceAuditInput) (PermissionEvidenceAuditPreview, error) {
	input.ApplicationID = strings.TrimSpace(input.ApplicationID)
	if !idPattern.MatchString(input.ApplicationID) {
		return PermissionEvidenceAuditPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}

	provided := map[string]PermissionEvidenceRowInput{}
	canonical := map[string]bool{}
	for _, resource := range permissionAuditResourceOrder() {
		canonical[resource.id] = true
	}
	var unsupported []PermissionEvidenceRowInput
	for _, row := range input.Rows {
		id := strings.TrimSpace(row.Resource)
		if id == "" {
			continue
		}
		if !canonical[id] {
			unsupported = append(unsupported, row)
			continue
		}
		provided[id] = row
	}

	var rows []PermissionEvidenceAuditRow
	var counts PermissionEvidenceAuditCounts
	remediation := []string{}

	for _, resource := range permissionAuditResourceOrder() {
		row := buildPermissionAuditRow(resource, provided[resource.id])
		rows = append(rows, row)
		tallyPermissionAuditRow(&counts, row.ConsistencyState)
		if hint := permissionAuditGlobalHint(row); hint != "" {
			remediation = append(remediation, hint)
		}
	}
	for _, extra := range unsupported {
		row := buildUnsupportedPermissionAuditRow(extra)
		rows = append(rows, row)
		tallyPermissionAuditRow(&counts, row.ConsistencyState)
		remediation = append(remediation, "An unrecognized permission was reported and needs Runtime review before it is trusted.")
	}

	counts.TotalRows = len(rows)
	counts.MalformedReceipts = len(input.MalformedReceipts)
	counts.Divergent = counts.SettingOnly + counts.ReceiptOnly + counts.Expired + counts.Denied + counts.MissingReview + counts.Unsupported

	preview := PermissionEvidenceAuditPreview{
		SchemaVersion:       "xnix.runtime.permission_evidence_audit.v1",
		RequestType:         "permission-evidence-audit-preview",
		AuditType:           "offline-permission-evidence-audit",
		Source:              "portal-access-policy+portal-receipts+settings+permission-review+resource-bridge",
		Desktop:             "KDE Plasma",
		RuntimeMethod:       "GetPermissionEvidenceAudit",
		ReadMethod:          "GetPermissionEvidenceAuditPreview",
		ApplicationID:       input.ApplicationID,
		Rows:                rows,
		RowIDs:              permissionAuditRowIDs(rows),
		RowCount:            len(rows),
		Counts:              counts,
		MalformedReceiptIDs: permissionAuditSafeIDs(input.MalformedReceipts),
		ExecutionGated:      input.ExecutionGated,
		ExecutionGateReason: permissionAuditSafeText(input.ExecutionReason),
		RemediationHints:    permissionAuditUniqueStrings(remediation),
		NextSafeReadOnlyChecks: []string{
			"review Portal access policy for each capability",
			"review recorded Portal permission receipts",
			"review compatibility permission review plan",
			"review Runtime-backed compatibility settings",
		},
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		KDEPolicyOwner:              false,
		UserVisible:                 true,
		ReviewOnly:                  true,
		RealPortalCallEnabled:       false,
		PermissionGrantEnabled:      false,
		PermissionRevokeEnabled:     false,
		ReceiptWriteEnabled:         false,
		SettingsPersisted:           false,
		ExecutionApproved:           false,
		StateRootPathExposed:        false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"grant a permission from the audit preview",
			"revoke a permission from the audit preview",
			"call real Portal transport from the audit preview",
			"write a Portal receipt from the audit preview",
			"persist settings from the audit preview",
			"approve execution from the audit preview",
			"mutate host root during the audit preview",
		},
		OverallState:       permissionAuditOverallState(counts),
		DesktopSafeSummary: permissionAuditSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "permission evidence audit preview"); err != nil {
		return PermissionEvidenceAuditPreview{}, err
	}
	if err := safety.ValidatePayload("permission evidence audit preview", preview); err != nil {
		return PermissionEvidenceAuditPreview{}, err
	}
	return preview, nil
}

func buildPermissionAuditRow(resource permissionAuditResource, input PermissionEvidenceRowInput) PermissionEvidenceAuditRow {
	setting := normalizePermissionState(input.SettingState, "enabled", "disabled", "absent")
	receipt := normalizePermissionState(input.ReceiptState, "granted", "denied", "expired", "pending", "absent")
	review := normalizePermissionState(input.ReviewState, "approved", "pending", "blocked", "absent")

	state := permissionAuditConsistencyState(resource.policyDecision, resource.portalRequired, setting, receipt, review)
	row := PermissionEvidenceAuditRow{
		ID:               resource.id,
		Title:            resource.title,
		PortalInterface:  resource.portalInterface,
		PolicyDecision:   resource.policyDecision,
		ConsistencyState: state,
		SettingState:     setting,
		ReceiptState:     receipt,
		ReviewState:      review,
		BridgeExposed:    input.BridgeExposed,
		ReceiptRefs:      permissionAuditSafeIDs(input.ReceiptRefs),
		RemediationHint:  permissionAuditRemediation(state),
		ReviewOnly:       true,
	}
	return row
}

func buildUnsupportedPermissionAuditRow(input PermissionEvidenceRowInput) PermissionEvidenceAuditRow {
	return PermissionEvidenceAuditRow{
		ID:               permissionAuditSafeText(strings.TrimSpace(input.Resource)),
		Title:            "Unrecognized permission",
		PolicyDecision:   "unsupported",
		ConsistencyState: "unsupported",
		SettingState:     normalizePermissionState(input.SettingState, "enabled", "disabled", "absent"),
		ReceiptState:     normalizePermissionState(input.ReceiptState, "granted", "denied", "expired", "pending", "absent"),
		ReviewState:      normalizePermissionState(input.ReviewState, "approved", "pending", "blocked", "absent"),
		BridgeExposed:    input.BridgeExposed,
		ReceiptRefs:      permissionAuditSafeIDs(input.ReceiptRefs),
		RemediationHint:  permissionAuditRemediation("unsupported"),
		ReviewOnly:       true,
	}
}

// permissionAuditConsistencyState reconciles the policy decision, whether a
// Portal request is required, the user setting, the recorded Portal receipt, and
// the review state into one deterministic state.
func permissionAuditConsistencyState(policyDecision string, portalRequired bool, setting, receipt, review string) string {
	if policyDecision == "deny" {
		if receipt == "granted" {
			// A grant exists for a capability policy says is denied.
			return "receipt-only"
		}
		return "blocked-by-policy"
	}
	if !portalRequired {
		// Allowed without a Portal request (for example network); there is no
		// receipt to reconcile.
		return "consistent"
	}
	switch receipt {
	case "expired":
		return "expired"
	case "denied":
		return "denied"
	case "pending":
		return "missing-review"
	}
	reviewSettled := review == "approved" || review == "absent" || review == ""
	declared := setting == "enabled"
	if declared && !reviewSettled {
		return "missing-review"
	}
	switch receipt {
	case "granted":
		if !reviewSettled {
			return "missing-review"
		}
		return "consistent"
	default: // absent
		if declared {
			return "setting-only"
		}
		return "consistent"
	}
}

func permissionAuditRemediation(state string) string {
	switch state {
	case "setting-only":
		return "The setting is enabled but no Portal permission is recorded yet; it will be requested when the application next needs it."
	case "receipt-only":
		return "A Portal permission is recorded without a matching setting; review whether the setting should be enabled."
	case "expired":
		return "The Portal permission expired and will be requested again the next time it is needed."
	case "denied":
		return "The Portal permission was denied; the user can re-request it from the Compatibility Center."
	case "missing-review":
		return "This permission needs a review before it is treated as settled."
	case "blocked-by-policy":
		return "This capability stays blocked by Runtime policy until the user changes the application policy."
	case "unsupported":
		return "This permission is not recognized by the Runtime audit and needs review."
	default:
		return "Settings and Portal receipts agree; no action is needed."
	}
}

func permissionAuditGlobalHint(row PermissionEvidenceAuditRow) string {
	switch row.ConsistencyState {
	case "consistent", "blocked-by-policy":
		return ""
	default:
		return row.Title + ": " + row.RemediationHint
	}
}

func tallyPermissionAuditRow(counts *PermissionEvidenceAuditCounts, state string) {
	switch state {
	case "consistent":
		counts.Consistent++
	case "setting-only":
		counts.SettingOnly++
	case "receipt-only":
		counts.ReceiptOnly++
	case "expired":
		counts.Expired++
	case "denied":
		counts.Denied++
	case "missing-review":
		counts.MissingReview++
	case "blocked-by-policy":
		counts.BlockedByPolicy++
	case "unsupported":
		counts.Unsupported++
	}
}

func permissionAuditOverallState(counts PermissionEvidenceAuditCounts) string {
	if counts.MalformedReceipts > 0 || counts.Unsupported > 0 {
		return "needs-review"
	}
	if counts.Divergent > 0 {
		return "needs-review"
	}
	return "consistent"
}

func permissionAuditSummary(counts PermissionEvidenceAuditCounts) string {
	if counts.MalformedReceipts > 0 {
		return "Some Portal receipts could not be parsed, so the permission audit is marked for review without changing any permission."
	}
	if counts.Divergent > 0 || counts.Unsupported > 0 {
		return "The permission audit found settings and Portal receipts that disagree and needs review; it changed no permissions."
	}
	return "The permission audit found settings and Portal receipts in agreement and changed no permissions."
}

func permissionAuditRowIDs(rows []PermissionEvidenceAuditRow) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	return ids
}

func normalizePermissionState(value string, allowed ...string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	for _, candidate := range allowed {
		if value == candidate {
			return value
		}
	}
	return allowed[len(allowed)-1]
}

func permissionAuditSafeIDs(values []string) []string {
	safeValues := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || !safety.Safe(value) {
			continue
		}
		safeValues = append(safeValues, value)
	}
	sort.Strings(safeValues)
	return permissionAuditUniqueStrings(safeValues)
}

func permissionAuditSafeText(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || !safety.Safe(value) {
		return ""
	}
	return value
}

func permissionAuditUniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}
