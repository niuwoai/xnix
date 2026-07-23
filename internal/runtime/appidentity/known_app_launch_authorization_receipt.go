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
