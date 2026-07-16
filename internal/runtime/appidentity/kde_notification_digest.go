package appidentity

import (
	"fmt"
	"sort"

	"xnix.local/xnix/internal/runtime/safety"
)

type KDENotificationDigestEvent struct {
	Group   string `json:"group"`
	EventID string `json:"event_id"`
}

type KDENotificationDigestEntry struct {
	Group             string `json:"group"`
	EventID           string `json:"event_id"`
	DeduplicationKey  string `json:"deduplication_key"`
	Severity          string `json:"severity"`
	Label             string `json:"label"`
	Summary           string `json:"summary"`
	NextSafeReadRoute string `json:"next_safe_read_route"`
	Blocked           bool   `json:"blocked"`
	RequiresReview    bool   `json:"requires_review"`
}

type KDENotificationDigestCounts struct {
	InputEventCount     int `json:"input_event_count"`
	DigestEntryCount    int `json:"digest_entry_count"`
	DuplicateEventCount int `json:"duplicate_event_count"`
	CriticalEntryCount  int `json:"critical_entry_count"`
	ReviewEntryCount    int `json:"review_entry_count"`
	BlockedEntryCount   int `json:"blocked_entry_count"`
}

type KDENotificationDigestPreview struct {
	SchemaVersion          string                       `json:"schema_version"`
	RequestType            string                       `json:"request_type"`
	DigestType             string                       `json:"digest_type"`
	Source                 string                       `json:"source"`
	Desktop                string                       `json:"desktop"`
	RuntimeMethod          string                       `json:"runtime_method"`
	ReadMethod             string                       `json:"read_method"`
	ApplicationID          string                       `json:"application_id"`
	DisplayName            string                       `json:"display_name"`
	Entries                []KDENotificationDigestEntry `json:"entries"`
	Counts                 KDENotificationDigestCounts  `json:"counts"`
	DigestReady            bool                         `json:"digest_ready"`
	RuntimeOwned           bool                         `json:"runtime_owned"`
	KDEPolicyOwner         bool                         `json:"kde_policy_owner"`
	ReviewOnly             bool                         `json:"review_only"`
	NotificationsSent      bool                         `json:"notifications_sent"`
	LiveTrayBridgeEnabled  bool                         `json:"live_tray_bridge_enabled"`
	RequestObjectsCreated  bool                         `json:"request_objects_created"`
	PermissionGrantEnabled bool                         `json:"permission_grant_enabled"`
	BackendLaunchEnabled   bool                         `json:"backend_launch_enabled"`
	BackendProcessStarted  bool                         `json:"backend_process_started"`
	StateRootPathExposed   bool                         `json:"state_root_path_exposed"`
	RawCommandExposed      bool                         `json:"raw_command_exposed"`
	BackendDetailsExposed  bool                         `json:"backend_details_exposed"`
	HostRootModified       bool                         `json:"host_root_modified"`
	DesktopSafeSummary     string                       `json:"desktop_safe_summary"`
}

type kdeNotificationDigestRule struct {
	severity string
	label    string
	summary  string
	route    string
	blocked  bool
	review   bool
	order    int
}

var kdeNotificationDigestRules = map[string]kdeNotificationDigestRule{
	"needs-review:approval-required": {
		severity: "critical", label: "Approval review needed", summary: "A compatibility action is waiting for explicit review.", route: "kde-center-page-preview", review: true, order: 0,
	},
	"blocked-action:execution-blocked": {
		severity: "critical", label: "Action remains blocked", summary: "Runtime safety checks are keeping an action closed.", route: "application-readiness-preview", blocked: true, review: true, order: 1,
	},
	"permission-attention:permission-review": {
		severity: "normal", label: "Permission review needed", summary: "A requested desktop capability needs user review.", route: "portal-permission-renewal-preview", review: true, order: 2,
	},
	"diagnostic-issue:diagnostic-failed": {
		severity: "critical", label: "Diagnostic issue found", summary: "Runtime diagnostic evidence needs review before another action.", route: "support-case-timeline-preview", blocked: true, review: true, order: 3,
	},
	"snapshot-warning:snapshot-missing": {
		severity: "normal", label: "Restore point unavailable", summary: "A safe restore point has not been recorded for this application.", route: "snapshot-restore-candidates-preview", blocked: true, review: true, order: 4,
	},
	"readiness-change:readiness-pending": {
		severity: "low", label: "Readiness changed", summary: "Application readiness evidence is still being completed.", route: "application-readiness-preview", review: false, order: 5,
	},
}

func (plan Plan) KDENotificationDigestPreview(events []KDENotificationDigestEvent) (KDENotificationDigestPreview, error) {
	if err := plan.ValidateSafeForDesktop(); err != nil {
		return KDENotificationDigestPreview{}, err
	}
	if err := safety.ValidateLines(map[string]string{
		"application id": plan.ApplicationID,
		"display name":   plan.DisplayName,
		"desktop file":   plan.DesktopFile,
	}); err != nil {
		return KDENotificationDigestPreview{}, err
	}
	if len(events) == 0 {
		return KDENotificationDigestPreview{}, fmt.Errorf("KDE notification digest requires at least one event")
	}

	entries := make([]KDENotificationDigestEntry, 0, len(events))
	seen := make(map[string]bool, len(events))
	duplicates := 0
	for _, event := range events {
		kind := event.Group + ":" + event.EventID
		rule, ok := kdeNotificationDigestRules[kind]
		if !ok {
			return KDENotificationDigestPreview{}, fmt.Errorf("unsupported KDE notification digest event: %s", kind)
		}
		deduplicationKey := plan.ApplicationID + ":" + kind
		if seen[deduplicationKey] {
			duplicates++
			continue
		}
		seen[deduplicationKey] = true
		entries = append(entries, KDENotificationDigestEntry{
			Group: event.Group, EventID: event.EventID, DeduplicationKey: deduplicationKey,
			Severity: rule.severity, Label: rule.label, Summary: rule.summary,
			NextSafeReadRoute: rule.route, Blocked: rule.blocked, RequiresReview: rule.review,
		})
	}

	sort.SliceStable(entries, func(i, j int) bool {
		left := kdeNotificationDigestRules[entries[i].Group+":"+entries[i].EventID].order
		right := kdeNotificationDigestRules[entries[j].Group+":"+entries[j].EventID].order
		return left < right
	})

	counts := KDENotificationDigestCounts{InputEventCount: len(events), DigestEntryCount: len(entries), DuplicateEventCount: duplicates}
	for _, entry := range entries {
		if entry.Severity == "critical" {
			counts.CriticalEntryCount++
		}
		if entry.RequiresReview {
			counts.ReviewEntryCount++
		}
		if entry.Blocked {
			counts.BlockedEntryCount++
		}
	}

	return KDENotificationDigestPreview{
		SchemaVersion: "xnix.runtime.kde_notification_digest.v1", RequestType: "kde-notification-digest-preview",
		DigestType: "review-only-runtime-event-digest", Source: "runtime-review-events", Desktop: "KDE Plasma",
		RuntimeMethod: "GetKDENotificationDigest", ReadMethod: "GetKDENotificationDigestPreview",
		ApplicationID: plan.ApplicationID, DisplayName: plan.DisplayName, Entries: entries, Counts: counts,
		DigestReady: true, RuntimeOwned: true, KDEPolicyOwner: false, ReviewOnly: true,
		DesktopSafeSummary: "KDE can show one deduplicated review digest while notification delivery and every Runtime action stay disabled.",
	}, nil
}
