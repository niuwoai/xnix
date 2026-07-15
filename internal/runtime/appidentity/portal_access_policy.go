package appidentity

import "errors"

type PortalAccessPolicyPreview struct {
	SchemaVersion          string              `json:"schema_version"`
	RequestType            string              `json:"request_type"`
	PolicyType             string              `json:"policy_type"`
	Source                 string              `json:"source"`
	Desktop                string              `json:"desktop"`
	RuntimeMethod          string              `json:"runtime_method"`
	ReadMethod             string              `json:"read_method"`
	ApplicationID          string              `json:"application_id"`
	Operation              string              `json:"operation"`
	Decision               string              `json:"decision"`
	RuntimeOwned           bool                `json:"runtime_owned"`
	GoRuntimeBacked        bool                `json:"go_runtime_backed"`
	KDEPolicyOwner         bool                `json:"kde_policy_owner"`
	PortalRequired         bool                `json:"portal_required"`
	PortalInterface        string              `json:"portal_interface"`
	UserMediationRequired  bool                `json:"user_mediation_required"`
	DirectAccessAllowed    bool                `json:"direct_access_allowed"`
	Resources              []string            `json:"resources"`
	RequestFlow            PortalRequestFlow   `json:"request_flow"`
	RequestObjectCreated   bool                `json:"request_object_created"`
	PermissionGranted      bool                `json:"permission_granted"`
	HostPermissionChanged  bool                `json:"host_permission_changed"`
	HostRootModified       bool                `json:"host_root_modified"`
	NetworkRequired        bool                `json:"network_required"`
	BackendDetailsExposed  bool                `json:"backend_details_exposed"`
	DesktopSafeSummary     string              `json:"desktop_safe_summary"`
}

type PortalRequestFlow struct {
	DBusAPI                 string `json:"dbus_api"`
	RequestObjectRequired   bool   `json:"request_object_required"`
	CompletionSignal        string `json:"completion_signal"`
	RuntimePolicyOwner      bool   `json:"runtime_policy_owner"`
	DesktopShellPolicyOwner bool   `json:"desktop_shell_policy_owner"`
}

type portalAccessRule struct {
	portalInterface string
	decision        string
	resources       []string
	summary         string
}

var portalAccessRules = map[string]portalAccessRule{
	"file-open": {
		portalInterface: "org.freedesktop.portal.FileChooser",
		decision:        "ask",
		resources:       []string{"documents", "downloads", "selected-files"},
		summary:         "File access requires a user-approved desktop portal request.",
	},
	"uri-open": {
		portalInterface: "org.freedesktop.portal.OpenURI",
		decision:        "ask",
		resources:       []string{"external-uri"},
		summary:         "URI handling requires a user-approved desktop portal request.",
	},
	"print": {
		portalInterface: "org.freedesktop.portal.Print",
		decision:        "ask",
		resources:       []string{"printer"},
		summary:         "Printing requires a user-approved desktop portal request.",
	},
	"screenshot": {
		portalInterface: "org.freedesktop.portal.Screenshot",
		decision:        "ask",
		resources:       []string{"screen"},
		summary:         "Screenshots require a user-approved desktop portal request.",
	},
	"clipboard": {
		portalInterface: "org.freedesktop.portal.Clipboard",
		decision:        "ask",
		resources:       []string{"clipboard"},
		summary:         "Clipboard access requires a user-approved desktop portal request.",
	},
	"camera": {
		portalInterface: "org.freedesktop.portal.Camera",
		decision:        "deny",
		resources:       []string{"camera"},
		summary:         "Camera access is denied until the user changes the application policy.",
	},
	"remote-desktop": {
		portalInterface: "org.freedesktop.portal.RemoteDesktop",
		decision:        "deny",
		resources:       []string{"screen", "input-devices"},
		summary:         "Remote desktop access is denied until the user changes the application policy.",
	},
}

func NewPortalAccessPolicyPreview(applicationID string, operation string) (PortalAccessPolicyPreview, error) {
	if !idPattern.MatchString(applicationID) {
		return PortalAccessPolicyPreview{}, errors.New("application id must be a reverse-DNS identifier")
	}
	rule, ok := portalAccessRules[operation]
	if !ok {
		return PortalAccessPolicyPreview{}, errors.New("operation must be one of: file-open, uri-open, print, screenshot, clipboard, camera, remote-desktop")
	}

	preview := PortalAccessPolicyPreview{
		SchemaVersion:         "xnix.runtime.portal_access_policy.v1",
		RequestType:           "portal-access-policy-preview",
		PolicyType:            "portal-access",
		Source:                "go-runtime-portal-access-policy",
		Desktop:               "KDE Plasma",
		RuntimeMethod:         "GetPortalAccessPolicy",
		ReadMethod:            "GetPortalAccessPolicyPreview",
		ApplicationID:         applicationID,
		Operation:             operation,
		Decision:              rule.decision,
		RuntimeOwned:          true,
		GoRuntimeBacked:       true,
		KDEPolicyOwner:        false,
		PortalRequired:        true,
		PortalInterface:       rule.portalInterface,
		UserMediationRequired: true,
		DirectAccessAllowed:   false,
		Resources:             append([]string(nil), rule.resources...),
		RequestFlow: PortalRequestFlow{
			DBusAPI:                 "XDG Desktop Portal",
			RequestObjectRequired:   true,
			CompletionSignal:        "portal-response",
			RuntimePolicyOwner:      true,
			DesktopShellPolicyOwner: false,
		},
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		HostPermissionChanged: false,
		HostRootModified:      false,
		NetworkRequired:       false,
		BackendDetailsExposed: false,
		DesktopSafeSummary:    rule.summary,
	}
	if err := validateNoBackendTerms(preview, "Portal access policy preview"); err != nil {
		return PortalAccessPolicyPreview{}, err
	}
	return preview, nil
}
