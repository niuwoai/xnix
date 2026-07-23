package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xnix.local/xnix/internal/runtime/winapp"
)

const (
	KnownAppLaunchAuthorizationReceiptSchemaVersion = "xnix.runtime.known_app_launch_authorization_receipt.v1"
	KnownAppLaunchAuthorizationReceiptRequestType   = "known-app-launch-authorization-receipt-preview"
	KnownAppLaunchAuthorizationReceiptAction        = "review-launch-authorization"
	KnownAppLaunchGateSchemaVersion                 = "xnix.runtime.known_app_launch_gate.v1"
	KnownAppLaunchGateRequestType                   = "known-app-launch-gate-preview"
	knownAppLaunchAuthorizationReceiptDir           = "runtime/authorization-receipts"
)

type KnownAppLaunchAuthorizationReceiptRequest struct {
	AppID           string
	StateRoot       string
	Authorize       string
	EvidenceSource  string
	CenterCardState string
	RecordedAtUTC   time.Time
}

type KnownAppLaunchAuthorizationReceiptPreview struct {
	SchemaVersion               string   `json:"schema_version"`
	RequestType                 string   `json:"request_type"`
	Source                      string   `json:"source"`
	RuntimeMethod               string   `json:"runtime_method"`
	AppID                       string   `json:"app_id"`
	DisplayName                 string   `json:"display_name"`
	AppVersion                  string   `json:"app_version"`
	EvidenceSource              string   `json:"evidence_source"`
	CenterCardState             string   `json:"center_card_state"`
	RequestedActionID           string   `json:"requested_action_id"`
	ReceiptID                   string   `json:"receipt_id"`
	ReceiptType                 string   `json:"receipt_type"`
	ReceiptState                string   `json:"receipt_state"`
	ReceiptWritten              bool     `json:"receipt_written"`
	ReceiptPathExposed          bool     `json:"receipt_path_exposed"`
	StateRootPathExposed        bool     `json:"state_root_path_exposed"`
	ManagedStateRoot            bool     `json:"managed_state_root"`
	RuntimeOwned                bool     `json:"runtime_owned"`
	GoRuntimeBacked             bool     `json:"go_runtime_backed"`
	LaunchAuthorizationRequired bool     `json:"launch_authorization_required"`
	LaunchAuthorizationRecorded bool     `json:"launch_authorization_recorded"`
	LaunchGateReady             bool     `json:"launch_gate_ready"`
	DirectLaunchEnabled         bool     `json:"direct_launch_enabled"`
	DesktopLaunchEnabled        bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled        bool     `json:"backend_launch_enabled"`
	BackendProcessStarted       bool     `json:"backend_process_started"`
	HostRootModified            bool     `json:"host_root_modified"`
	StateRootModified           bool     `json:"state_root_modified"`
	DockerSocketMounted         bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool     `json:"broad_host_mount_required"`
	RawArtifactPathExposed      bool     `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	BlockedActions              []string `json:"blocked_actions"`
	DesktopSafeSummary          string   `json:"desktop_safe_summary"`
}

type KnownAppLaunchGateRequest struct {
	AppID         string
	StateRoot     string
	ReceiptID     string
	CacheRoot     string
	GuestBoundary string
}

type KnownAppLaunchGatePreview struct {
	SchemaVersion               string   `json:"schema_version"`
	RequestType                 string   `json:"request_type"`
	Source                      string   `json:"source"`
	RuntimeMethod               string   `json:"runtime_method"`
	AppID                       string   `json:"app_id"`
	DisplayName                 string   `json:"display_name"`
	AppVersion                  string   `json:"app_version"`
	ReceiptID                   string   `json:"receipt_id"`
	ReceiptLookupState          string   `json:"receipt_lookup_state"`
	ReceiptAccepted             bool     `json:"receipt_accepted"`
	ReceiptRejectedReason       string   `json:"receipt_rejected_reason,omitempty"`
	ReceiptPathExposed          bool     `json:"receipt_path_exposed"`
	StateRootPathExposed        bool     `json:"state_root_path_exposed"`
	GuestBoundaryRequired       bool     `json:"guest_boundary_required"`
	GuestBoundary               string   `json:"guest_boundary"`
	GuestBoundaryAccepted       bool     `json:"guest_boundary_accepted"`
	DispatchStatus              string   `json:"dispatch_status"`
	DispatchReady               bool     `json:"dispatch_ready"`
	DispatchRequestMaterialized bool     `json:"dispatch_request_materialized"`
	LaunchGateState             string   `json:"launch_gate_state"`
	LaunchGateBlockedReason     string   `json:"launch_gate_blocked_reason,omitempty"`
	ControlledDispatchReady     bool     `json:"controlled_dispatch_ready"`
	DirectLaunchEnabled         bool     `json:"direct_launch_enabled"`
	DesktopLaunchEnabled        bool     `json:"desktop_launch_enabled"`
	BackendLaunchEnabled        bool     `json:"backend_launch_enabled"`
	BackendProcessStarted       bool     `json:"backend_process_started"`
	ExecutionStarted            bool     `json:"execution_started"`
	HostRootModified            bool     `json:"host_root_modified"`
	DockerSocketMounted         bool     `json:"docker_socket_mounted"`
	BroadHostMountRequired      bool     `json:"broad_host_mount_required"`
	RawArtifactPathExposed      bool     `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed       bool     `json:"backend_details_exposed"`
	BlockedActions              []string `json:"blocked_actions"`
	DesktopSafeSummary          string   `json:"desktop_safe_summary"`
}

type knownAppLaunchAuthorizationReceiptFile struct {
	SchemaVersion            string `json:"schema_version"`
	ReceiptType              string `json:"receipt_type"`
	ReceiptID                string `json:"receipt_id"`
	AppID                    string `json:"app_id"`
	DisplayName              string `json:"display_name"`
	AppVersion               string `json:"app_version"`
	EvidenceSource           string `json:"evidence_source"`
	CenterCardState          string `json:"center_card_state"`
	AuthorizedActionID       string `json:"authorized_action_id"`
	LaunchAuthorizationState string `json:"launch_authorization_state"`
	LaunchGateState          string `json:"launch_gate_state"`
	RuntimeOwned             bool   `json:"runtime_owned"`
	GoRuntimeBacked          bool   `json:"go_runtime_backed"`
	DirectLaunchEnabled      bool   `json:"direct_launch_enabled"`
	DesktopLaunchEnabled     bool   `json:"desktop_launch_enabled"`
	BackendLaunchEnabled     bool   `json:"backend_launch_enabled"`
	BackendProcessStarted    bool   `json:"backend_process_started"`
	HostRootModified         bool   `json:"host_root_modified"`
	DockerSocketMounted      bool   `json:"docker_socket_mounted"`
	BroadHostMountRequired   bool   `json:"broad_host_mount_required"`
	RawArtifactPathExposed   bool   `json:"raw_artifact_path_exposed"`
	BackendDetailsExposed    bool   `json:"backend_details_exposed"`
	RecordedAtUTC            string `json:"recorded_at_utc"`
}

func RecordKnownAppLaunchAuthorizationReceipt(request KnownAppLaunchAuthorizationReceiptRequest) (KnownAppLaunchAuthorizationReceiptPreview, error) {
	if strings.TrimSpace(request.Authorize) != KnownAppLaunchAuthorizationReceiptAction {
		return KnownAppLaunchAuthorizationReceiptPreview{}, fmt.Errorf("known app launch authorization receipt requires --authorize %s", KnownAppLaunchAuthorizationReceiptAction)
	}
	if strings.TrimSpace(request.StateRoot) == "" {
		return KnownAppLaunchAuthorizationReceiptPreview{}, errors.New("known app launch authorization receipt requires --state-root")
	}
	app, err := winapp.LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, err
	}
	stateRoot, err := filepath.Abs(filepath.Clean(request.StateRoot))
	if err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, fmt.Errorf("resolve managed state root: %w", err)
	}
	receiptID := KnownAppLaunchAuthorizationReceiptID(app.ID, app.Version)
	receiptPath, err := KnownAppLaunchAuthorizationReceiptPath(stateRoot, receiptID)
	if err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, err
	}
	if err := os.MkdirAll(filepath.Dir(receiptPath), 0o700); err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, fmt.Errorf("create authorization receipt store: %w", err)
	}
	recordedAt := request.RecordedAtUTC.UTC()
	if recordedAt.IsZero() {
		recordedAt = time.Now().UTC()
	}
	evidenceSource := strings.TrimSpace(request.EvidenceSource)
	if evidenceSource == "" {
		evidenceSource = "staged-launcher-dispatch-smoke"
	}
	centerCardState := strings.TrimSpace(request.CenterCardState)
	if centerCardState == "" {
		centerCardState = "validated-launch-authorization-required"
	}
	receipt := knownAppLaunchAuthorizationReceiptFile{
		SchemaVersion:            KnownAppLaunchAuthorizationReceiptSchemaVersion,
		ReceiptType:              "opaque-known-app-launch-authorization",
		ReceiptID:                receiptID,
		AppID:                    app.ID,
		DisplayName:              app.DisplayName,
		AppVersion:               app.Version,
		EvidenceSource:           evidenceSource,
		CenterCardState:          centerCardState,
		AuthorizedActionID:       KnownAppLaunchAuthorizationReceiptAction,
		LaunchAuthorizationState: "recorded",
		LaunchGateState:          "receipt-recorded-launch-still-gated",
		RuntimeOwned:             true,
		GoRuntimeBacked:          true,
		DirectLaunchEnabled:      false,
		DesktopLaunchEnabled:     false,
		BackendLaunchEnabled:     false,
		BackendProcessStarted:    false,
		HostRootModified:         false,
		DockerSocketMounted:      false,
		BroadHostMountRequired:   false,
		RawArtifactPathExposed:   false,
		BackendDetailsExposed:    false,
		RecordedAtUTC:            recordedAt.Format(time.RFC3339),
	}
	content, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, fmt.Errorf("encode authorization receipt: %w", err)
	}
	content = append(content, '\n')
	if err := os.WriteFile(receiptPath, content, 0o600); err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, fmt.Errorf("write authorization receipt: %w", err)
	}
	preview := KnownAppLaunchAuthorizationReceiptPreview{
		SchemaVersion:               KnownAppLaunchAuthorizationReceiptSchemaVersion,
		RequestType:                 KnownAppLaunchAuthorizationReceiptRequestType,
		Source:                      "compatibility-center-preview+staged-launcher-dispatch-smoke+runtime-write-gate",
		RuntimeMethod:               "RecordKnownAppLaunchAuthorizationReceipt",
		AppID:                       app.ID,
		DisplayName:                 app.DisplayName,
		AppVersion:                  app.Version,
		EvidenceSource:              evidenceSource,
		CenterCardState:             centerCardState,
		RequestedActionID:           KnownAppLaunchAuthorizationReceiptAction,
		ReceiptID:                   receiptID,
		ReceiptType:                 "opaque-known-app-launch-authorization",
		ReceiptState:                "recorded",
		ReceiptWritten:              true,
		ReceiptPathExposed:          false,
		StateRootPathExposed:        false,
		ManagedStateRoot:            true,
		RuntimeOwned:                true,
		GoRuntimeBacked:             true,
		LaunchAuthorizationRequired: true,
		LaunchAuthorizationRecorded: true,
		LaunchGateReady:             true,
		DirectLaunchEnabled:         false,
		DesktopLaunchEnabled:        false,
		BackendLaunchEnabled:        false,
		BackendProcessStarted:       false,
		HostRootModified:            false,
		StateRootModified:           true,
		DockerSocketMounted:         false,
		BroadHostMountRequired:      false,
		RawArtifactPathExposed:      false,
		BackendDetailsExposed:       false,
		BlockedActions: []string{
			"start backend from authorization receipt",
			"expose receipt path to KDE",
			"enable direct desktop launch",
			"mutate host root",
		},
		DesktopSafeSummary: "Runtime recorded an opaque known-application launch authorization receipt; the desktop can show the receipt state, while actual launch remains gate-controlled.",
	}
	if err := validateNoBackendTerms(preview, "known app launch authorization receipt preview"); err != nil {
		return KnownAppLaunchAuthorizationReceiptPreview{}, err
	}
	return preview, nil
}

func PreviewKnownAppLaunchGate(request KnownAppLaunchGateRequest) (KnownAppLaunchGatePreview, error) {
	if strings.TrimSpace(request.StateRoot) == "" {
		return KnownAppLaunchGatePreview{}, errors.New("known app launch gate requires --state-root")
	}
	if strings.TrimSpace(request.ReceiptID) == "" {
		return KnownAppLaunchGatePreview{}, errors.New("known app launch gate requires --receipt-id")
	}
	app, err := winapp.LookupKnownPortableApp(request.AppID)
	if err != nil {
		return KnownAppLaunchGatePreview{}, err
	}
	receiptID := strings.TrimSpace(request.ReceiptID)
	preview := baseKnownAppLaunchGatePreview(app, receiptID, strings.TrimSpace(request.GuestBoundary))
	receiptPath, err := KnownAppLaunchAuthorizationReceiptPath(request.StateRoot, receiptID)
	if err != nil {
		preview.ReceiptLookupState = "receipt-path-invalid"
		preview.ReceiptRejectedReason = "authorization receipt path is outside the managed state root"
		preview.LaunchGateState = "receipt-path-invalid-fail-closed"
		preview.DesktopSafeSummary = app.DisplayName + " launch gate rejected an invalid authorization receipt path without exposing storage details."
		return validateKnownAppLaunchGatePreview(preview)
	}
	content, err := os.ReadFile(receiptPath)
	if err != nil {
		if os.IsNotExist(err) {
			preview.ReceiptLookupState = "missing-receipt"
			preview.ReceiptRejectedReason = "authorization receipt is missing"
			preview.LaunchGateState = "missing-receipt-fail-closed"
			preview.DesktopSafeSummary = app.DisplayName + " launch gate did not find an authorization receipt and remains closed."
			return validateKnownAppLaunchGatePreview(preview)
		}
		return KnownAppLaunchGatePreview{}, fmt.Errorf("read authorization receipt: %w", err)
	}
	var receipt knownAppLaunchAuthorizationReceiptFile
	if err := json.Unmarshal(content, &receipt); err != nil {
		preview.ReceiptLookupState = "malformed-receipt"
		preview.ReceiptRejectedReason = "authorization receipt could not be parsed"
		preview.LaunchGateState = "malformed-receipt-fail-closed"
		preview.DesktopSafeSummary = app.DisplayName + " launch gate rejected a malformed authorization receipt and remains closed."
		return validateKnownAppLaunchGatePreview(preview)
	}
	if reason := rejectKnownAppLaunchAuthorizationReceipt(receipt, app, receiptID); reason != "" {
		preview.ReceiptLookupState = "rejected-receipt"
		preview.ReceiptRejectedReason = reason
		preview.LaunchGateState = "rejected-receipt-fail-closed"
		preview.DesktopSafeSummary = app.DisplayName + " launch gate rejected the authorization receipt and remains closed."
		return validateKnownAppLaunchGatePreview(preview)
	}
	preview.ReceiptLookupState = "accepted-receipt"
	preview.ReceiptAccepted = true
	if strings.TrimSpace(request.GuestBoundary) != winapp.KnownDispatchGuestBoundary {
		preview.LaunchGateBlockedReason = "controlled managed guest boundary is required before dispatch"
		preview.LaunchGateState = "guest-boundary-missing-fail-closed"
		preview.DesktopSafeSummary = app.DisplayName + " launch gate accepted the receipt, but the controlled managed guest boundary is still required."
		return validateKnownAppLaunchGatePreview(preview)
	}
	preview.GuestBoundaryAccepted = true
	dispatchPreview, err := winapp.PreviewKnownPortableDispatch(winapp.KnownDispatchRequest{
		AppID:     app.ID,
		CacheRoot: request.CacheRoot,
	})
	if err != nil {
		return KnownAppLaunchGatePreview{}, err
	}
	preview.DispatchStatus = dispatchPreview.Status
	preview.DispatchReady = dispatchPreview.DispatchReady
	if !dispatchPreview.DispatchReady {
		preview.LaunchGateBlockedReason = "managed artifact preparation is required before dispatch"
		preview.LaunchGateState = "dispatch-preparation-required"
		preview.DesktopSafeSummary = app.DisplayName + " launch gate accepted the receipt and boundary, but managed artifact preparation is required before dispatch."
		return validateKnownAppLaunchGatePreview(preview)
	}
	preview.DispatchRequestMaterialized = true
	preview.ControlledDispatchReady = true
	preview.LaunchGateState = "controlled-dispatch-ready"
	preview.DesktopSafeSummary = app.DisplayName + " launch gate accepted the opaque receipt and controlled boundary; dispatch is ready but backend launch remains controlled by the Runtime."
	return validateKnownAppLaunchGatePreview(preview)
}

func KnownAppLaunchAuthorizationReceiptID(appID string, version string) string {
	id := stateRootNamespace(strings.TrimSpace(appID))
	if id == "" {
		id = "known-app"
	}
	return "known-app-launch-authorization-" + id + "-" + stateRootNamespace(strings.TrimSpace(version))
}

func KnownAppLaunchAuthorizationReceiptPath(stateRoot string, receiptID string) (string, error) {
	if strings.TrimSpace(stateRoot) == "" {
		return "", errors.New("known app launch authorization receipt path requires a state root")
	}
	if strings.TrimSpace(receiptID) == "" {
		return "", errors.New("known app launch authorization receipt path requires a receipt id")
	}
	cleanRoot, err := filepath.Abs(filepath.Clean(stateRoot))
	if err != nil {
		return "", fmt.Errorf("resolve managed state root: %w", err)
	}
	relative := filepath.Join(knownAppLaunchAuthorizationReceiptDir, receiptID+".json")
	candidate := filepath.Join(cleanRoot, relative)
	cleanCandidate := filepath.Clean(candidate)
	rel, err := filepath.Rel(cleanRoot, cleanCandidate)
	if err != nil {
		return "", fmt.Errorf("scope authorization receipt path: %w", err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel) {
		return "", errors.New("authorization receipt path escaped the managed state root")
	}
	return cleanCandidate, nil
}

func baseKnownAppLaunchGatePreview(app winapp.KnownPortableApp, receiptID string, guestBoundary string) KnownAppLaunchGatePreview {
	return KnownAppLaunchGatePreview{
		SchemaVersion:          KnownAppLaunchGateSchemaVersion,
		RequestType:            KnownAppLaunchGateRequestType,
		Source:                 KnownAppLaunchAuthorizationReceiptRequestType + "+windows-known-app-dispatch-preview",
		RuntimeMethod:          "PreviewKnownAppLaunchGate",
		AppID:                  app.ID,
		DisplayName:            app.DisplayName,
		AppVersion:             app.Version,
		ReceiptID:              receiptID,
		ReceiptLookupState:     "not-checked",
		ReceiptAccepted:        false,
		ReceiptPathExposed:     false,
		StateRootPathExposed:   false,
		GuestBoundaryRequired:  true,
		GuestBoundary:          guestBoundary,
		GuestBoundaryAccepted:  false,
		DispatchStatus:         "not-checked",
		DispatchReady:          false,
		LaunchGateState:        "closed",
		DirectLaunchEnabled:    false,
		DesktopLaunchEnabled:   false,
		BackendLaunchEnabled:   false,
		BackendProcessStarted:  false,
		ExecutionStarted:       false,
		HostRootModified:       false,
		DockerSocketMounted:    false,
		BroadHostMountRequired: false,
		RawArtifactPathExposed: false,
		BackendDetailsExposed:  false,
		BlockedActions: []string{
			"consume missing authorization receipt",
			"consume mismatched authorization receipt",
			"start backend from launch gate preview",
			"expose receipt path to KDE",
			"mutate host root",
		},
		DesktopSafeSummary: app.DisplayName + " launch gate is closed until the Runtime accepts an opaque authorization receipt.",
	}
}

func rejectKnownAppLaunchAuthorizationReceipt(receipt knownAppLaunchAuthorizationReceiptFile, app winapp.KnownPortableApp, receiptID string) string {
	switch {
	case receipt.SchemaVersion != KnownAppLaunchAuthorizationReceiptSchemaVersion:
		return "authorization receipt schema is not supported"
	case receipt.ReceiptType != "opaque-known-app-launch-authorization":
		return "authorization receipt type is not supported"
	case receipt.ReceiptID != receiptID:
		return "authorization receipt id does not match the requested receipt"
	case receipt.AppID != app.ID:
		return "authorization receipt app id does not match the requested app"
	case receipt.AppVersion != app.Version:
		return "authorization receipt app version does not match the requested app"
	case receipt.AuthorizedActionID != KnownAppLaunchAuthorizationReceiptAction:
		return "authorization receipt action is not accepted"
	case receipt.LaunchAuthorizationState != "recorded":
		return "authorization receipt is not recorded"
	case !receipt.RuntimeOwned || !receipt.GoRuntimeBacked:
		return "authorization receipt is not Runtime-owned Go evidence"
	case receipt.DirectLaunchEnabled || receipt.DesktopLaunchEnabled || receipt.BackendLaunchEnabled || receipt.BackendProcessStarted || receipt.HostRootModified:
		return "authorization receipt contains unsafe launch side effects"
	case receipt.DockerSocketMounted || receipt.BroadHostMountRequired || receipt.RawArtifactPathExposed || receipt.BackendDetailsExposed:
		return "authorization receipt contains unsafe exposure flags"
	default:
		return ""
	}
}

func validateKnownAppLaunchGatePreview(preview KnownAppLaunchGatePreview) (KnownAppLaunchGatePreview, error) {
	if err := validateNoBackendTerms(preview, "known app launch gate preview"); err != nil {
		return KnownAppLaunchGatePreview{}, err
	}
	return preview, nil
}
