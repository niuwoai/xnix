package appidentity

import (
	"errors"
	"sort"
	"strings"

	"xnix.local/xnix/internal/runtime/safety"
)

type PortalPermissionRenewalReceipt struct {
	Operation    string
	State        string
	ExpiringSoon bool
	EvidenceID   string
}

type PortalPermissionRenewalPreview struct {
	SchemaVersion               string                        `json:"schema_version"`
	RequestType                 string                        `json:"request_type"`
	PreviewType                 string                        `json:"preview_type"`
	Source                      string                        `json:"source"`
	Desktop                     string                        `json:"desktop"`
	RuntimeMethod               string                        `json:"runtime_method"`
	ReadMethod                  string                        `json:"read_method"`
	ApplicationID               string                        `json:"application_id"`
	Rows                        []PortalPermissionRenewalRow  `json:"rows"`
	RowIDs                      []string                      `json:"row_ids"`
	RowCount                    int                           `json:"row_count"`
	Counts                      PortalPermissionRenewalCounts `json:"counts"`
	MalformedReceiptIDs         []string                      `json:"malformed_receipt_ids"`
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
	FileContentRead             bool                          `json:"file_content_read"`
	FilePathsExposed            bool                          `json:"file_paths_exposed"`
	RawCommandExposed           bool                          `json:"raw_command_exposed"`
	BackendDetailsExposed       bool                          `json:"backend_details_exposed"`
	HostRootModified            bool                          `json:"host_root_modified"`
	NetworkRequired             bool                          `json:"network_required"`
	PrivilegedContainerRequired bool                          `json:"privileged_container_required"`
	BlockedActions              []string                      `json:"blocked_actions"`
	DesktopSafeSummary          string                        `json:"desktop_safe_summary"`
}

type PortalPermissionRenewalRow struct {
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	Operation               string   `json:"operation"`
	PortalInterface         string   `json:"portal_interface,omitempty"`
	PolicyDecision          string   `json:"policy_decision"`
	PortalRequired          bool     `json:"portal_required"`
	RenewalState            string   `json:"renewal_state"`
	ReceiptState            string   `json:"receipt_state"`
	ReceiptEvidenceIDs      []string `json:"receipt_evidence_ids"`
	KDESafeActionLabel      string   `json:"kde_safe_action_label"`
	UserSafeSummary         string   `json:"user_safe_summary"`
	NextSafeReadOnlyCheck   string   `json:"next_safe_read_only_check"`
	ReviewOnly              bool     `json:"review_only"`
	RealPortalCallEnabled   bool     `json:"real_portal_call_enabled"`
	PermissionGrantEnabled  bool     `json:"permission_grant_enabled"`
	PermissionRevokeEnabled bool     `json:"permission_revoke_enabled"`
	ReceiptWriteEnabled     bool     `json:"receipt_write_enabled"`
	SettingsPersisted       bool     `json:"settings_persisted"`
	ExecutionApproved       bool     `json:"execution_approved"`
	HostRootModified        bool     `json:"host_root_modified"`
	BackendDetailsExposed   bool     `json:"backend_details_exposed"`
	StateRootPathExposed    bool     `json:"state_root_path_exposed"`
}

type PortalPermissionRenewalCounts struct {
	TotalRows         int `json:"total_rows"`
	Current           int `json:"current"`
	NeedsReview       int `json:"needs_review"`
	ExpiringSoon      int `json:"expiring_soon"`
	Denied            int `json:"denied"`
	Revoked           int `json:"revoked"`
	MissingReceipt    int `json:"missing_receipt"`
	BlockedByPolicy   int `json:"blocked_by_policy"`
	MalformedReceipts int `json:"malformed_receipts"`
}

type portalPermissionRenewalResource struct {
	id              string
	title           string
	operation       string
	portalInterface string
	policyDecision  string
	portalRequired  bool
}

func portalPermissionRenewalResources() []portalPermissionRenewalResource {
	return []portalPermissionRenewalResource{
		{"files", "Files", "file-open", "org.freedesktop.portal.FileChooser", "ask", true},
		{"uris", "External links", "uri-open", "org.freedesktop.portal.OpenURI", "ask", true},
		{"print", "Printing", "print", "org.freedesktop.portal.Print", "ask", true},
		{"clipboard", "Clipboard", "clipboard", "org.freedesktop.portal.Clipboard", "ask", true},
		{"screen", "Screen capture", "screenshot", "org.freedesktop.portal.Screenshot", "ask", true},
		{"camera", "Camera", "camera", "org.freedesktop.portal.Camera", "deny", true},
		{"remote-desktop", "Remote desktop", "remote-desktop", "org.freedesktop.portal.RemoteDesktop", "deny", true},
		{"network", "Network policy", "network", "", "allow", false},
	}
}

func (plan Plan) PortalPermissionRenewalPreview(receipts []PortalPermissionRenewalReceipt, malformedReceipts []string) (PortalPermissionRenewalPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return PortalPermissionRenewalPreview{}, err
	}
	if plan.ApplicationID == "" || !idPattern.MatchString(plan.ApplicationID) {
		return PortalPermissionRenewalPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}

	latest := map[string]PortalPermissionRenewalReceipt{}
	for _, receipt := range receipts {
		operation := strings.TrimSpace(receipt.Operation)
		if operation == "" {
			continue
		}
		latest[operation] = receipt
	}

	rows := make([]PortalPermissionRenewalRow, 0, len(portalPermissionRenewalResources()))
	var counts PortalPermissionRenewalCounts
	for _, resource := range portalPermissionRenewalResources() {
		row := portalPermissionRenewalRow(resource, latest[resource.operation])
		rows = append(rows, row)
		tallyPortalPermissionRenewal(&counts, row.RenewalState)
	}
	counts.TotalRows = len(rows)
	counts.MalformedReceipts = len(malformedReceipts)

	preview := PortalPermissionRenewalPreview{
		SchemaVersion:       "xnix.runtime.portal_permission_renewal.v1",
		RequestType:         "portal-permission-renewal-preview",
		PreviewType:         "review-only-portal-permission-renewal",
		Source:              "portal-access-policy+portal-request-plan+portal-receipts+compatibility-permission-review+kde-center-page",
		Desktop:             "KDE Plasma",
		RuntimeMethod:       "GetPortalPermissionRenewal",
		ReadMethod:          "GetPortalPermissionRenewalPreview",
		ApplicationID:       plan.ApplicationID,
		Rows:                rows,
		RowIDs:              portalPermissionRenewalRowIDs(rows),
		RowCount:            len(rows),
		Counts:              counts,
		MalformedReceiptIDs: permissionAuditSafeIDs(malformedReceipts),
		NextSafeReadOnlyChecks: []string{
			"review Portal access policy",
			"review Portal request preview",
			"review compatibility permission review plan",
			"review KDE Compatibility Center permission state",
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
		FileContentRead:             false,
		FilePathsExposed:            false,
		RawCommandExposed:           false,
		BackendDetailsExposed:       false,
		HostRootModified:            false,
		NetworkRequired:             false,
		PrivilegedContainerRequired: false,
		BlockedActions: []string{
			"call real Portal transport from renewal preview",
			"grant a permission from renewal preview",
			"revoke a permission from renewal preview",
			"write a Portal receipt from renewal preview",
			"persist settings from renewal preview",
			"approve execution from renewal preview",
			"mutate host root during renewal preview",
		},
		DesktopSafeSummary: portalPermissionRenewalSummary(counts),
	}
	if err := validateNoBackendTerms(preview, "Portal permission renewal preview"); err != nil {
		return PortalPermissionRenewalPreview{}, err
	}
	if err := safety.ValidatePayload("Portal permission renewal preview", preview); err != nil {
		return PortalPermissionRenewalPreview{}, err
	}
	return preview, nil
}

func portalPermissionRenewalRow(resource portalPermissionRenewalResource, receipt PortalPermissionRenewalReceipt) PortalPermissionRenewalRow {
	receiptState := portalPermissionRenewalReceiptState(receipt.State)
	renewalState := portalPermissionRenewalState(resource, receiptState, receipt.ExpiringSoon)
	return PortalPermissionRenewalRow{
		ID:                      resource.id,
		Title:                   resource.title,
		Operation:               resource.operation,
		PortalInterface:         resource.portalInterface,
		PolicyDecision:          resource.policyDecision,
		PortalRequired:          resource.portalRequired,
		RenewalState:            renewalState,
		ReceiptState:            receiptState,
		ReceiptEvidenceIDs:      permissionAuditSafeIDs([]string{receipt.EvidenceID}),
		KDESafeActionLabel:      portalPermissionRenewalActionLabel(renewalState),
		UserSafeSummary:         portalPermissionRenewalRowSummary(resource, renewalState),
		NextSafeReadOnlyCheck:   portalPermissionRenewalNextCheck(renewalState),
		ReviewOnly:              true,
		RealPortalCallEnabled:   false,
		PermissionGrantEnabled:  false,
		PermissionRevokeEnabled: false,
		ReceiptWriteEnabled:     false,
		SettingsPersisted:       false,
		ExecutionApproved:       false,
		HostRootModified:        false,
		BackendDetailsExposed:   false,
		StateRootPathExposed:    false,
	}
}

func portalPermissionRenewalReceiptState(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "granted", "pending", "denied", "expired", "revoked", "absent":
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return "absent"
	}
}

func portalPermissionRenewalState(resource portalPermissionRenewalResource, receiptState string, expiringSoon bool) string {
	if !resource.portalRequired {
		return "current"
	}
	if resource.policyDecision == "deny" && receiptState == "absent" {
		return "blocked-by-policy"
	}
	switch receiptState {
	case "granted":
		if expiringSoon {
			return "expiring-soon"
		}
		return "current"
	case "pending", "expired":
		return "needs-review"
	case "denied":
		return "denied"
	case "revoked":
		return "revoked"
	default:
		return "missing-receipt"
	}
}

func portalPermissionRenewalActionLabel(state string) string {
	switch state {
	case "current":
		return "No permission action needed"
	case "expiring-soon":
		return "Review renewal"
	case "needs-review":
		return "Review permission"
	case "denied":
		return "Review denied permission"
	case "revoked":
		return "Review revoked permission"
	case "blocked-by-policy":
		return "Review application policy"
	default:
		return "Review missing permission receipt"
	}
}

func portalPermissionRenewalRowSummary(resource portalPermissionRenewalResource, state string) string {
	switch state {
	case "current":
		return resource.title + " permission evidence is current."
	case "expiring-soon":
		return resource.title + " permission evidence should be reviewed before it expires."
	case "needs-review":
		return resource.title + " permission evidence needs review before it is treated as settled."
	case "denied":
		return resource.title + " permission was denied and can only be requested again after review."
	case "revoked":
		return resource.title + " permission was revoked or cancelled and needs review before reuse."
	case "blocked-by-policy":
		return resource.title + " access is blocked by Runtime policy."
	default:
		return resource.title + " has no recorded Portal permission receipt yet."
	}
}

func portalPermissionRenewalNextCheck(state string) string {
	switch state {
	case "current":
		return "review permission evidence audit"
	case "expiring-soon":
		return "review expiry metadata and renewal prompt copy"
	case "needs-review":
		return "review Portal request preview"
	case "denied", "revoked":
		return "review Compatibility Center permission history"
	case "blocked-by-policy":
		return "review compatibility permission review plan"
	default:
		return "review Portal access policy and receipt ledger"
	}
}

func tallyPortalPermissionRenewal(counts *PortalPermissionRenewalCounts, state string) {
	switch state {
	case "current":
		counts.Current++
	case "needs-review":
		counts.NeedsReview++
	case "expiring-soon":
		counts.ExpiringSoon++
	case "denied":
		counts.Denied++
	case "revoked":
		counts.Revoked++
	case "missing-receipt":
		counts.MissingReceipt++
	case "blocked-by-policy":
		counts.BlockedByPolicy++
	}
}

func portalPermissionRenewalRowIDs(rows []PortalPermissionRenewalRow) []string {
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.ID)
	}
	sort.Strings(ids)
	return ids
}

func portalPermissionRenewalSummary(counts PortalPermissionRenewalCounts) string {
	if counts.MalformedReceipts > 0 {
		return "Some Portal receipts could not be parsed; renewal remains review-only and no permission was changed."
	}
	if counts.Denied+counts.Revoked+counts.MissingReceipt+counts.BlockedByPolicy+counts.NeedsReview+counts.ExpiringSoon > 0 {
		return "Portal permission renewal needs review for at least one capability; no Portal call, permission change, receipt write, settings change, or execution approval was performed."
	}
	return "Portal permission renewal evidence is current; no Portal call, permission change, receipt write, settings change, or execution approval was performed."
}
