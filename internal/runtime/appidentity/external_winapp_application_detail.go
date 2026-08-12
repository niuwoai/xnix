package appidentity

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ExternalWinAppApplicationDetailSchemaVersion = "xnix.runtime.external_winapp_application_detail.v1"
	ExternalWinAppApplicationDetailRequestType   = "external-winapp-application-detail-preview"
)

type ExternalWinAppApplicationDetailRequest struct {
	CompatibilityEvidenceBundlePath string
}

type ExternalWinAppApplicationDetail struct {
	Version                        string                                      `json:"version"`
	SchemaVersion                  string                                      `json:"schema_version"`
	RequestType                    string                                      `json:"request_type"`
	DetailType                     string                                      `json:"detail_type"`
	Source                         string                                      `json:"source"`
	RuntimeMethod                  string                                      `json:"runtime_method"`
	ReadMethod                     string                                      `json:"read_method"`
	Desktop                        string                                      `json:"desktop"`
	ApplicationID                  string                                      `json:"application_id"`
	DisplayName                    string                                      `json:"display_name"`
	AppVersion                     string                                      `json:"app_version,omitempty"`
	CompatibilityState             string                                      `json:"compatibility_state"`
	CompatibilityLabel             string                                      `json:"compatibility_label"`
	PrimaryStatusTone              string                                      `json:"primary_status_tone"`
	ExistingWindowsAppVerified     bool                                        `json:"existing_windows_app_verified"`
	RealWindowsAppRunVerified      bool                                        `json:"real_windows_app_run_verified"`
	FileOpenVerified               bool                                        `json:"file_open_verified"`
	RuntimeGUIEvidenceVerified     bool                                        `json:"runtime_gui_evidence_verified"`
	DesktopEvidenceVerified        bool                                        `json:"desktop_evidence_verified"`
	KDEPageEvidenceVerified        bool                                        `json:"kde_page_evidence_verified"`
	EvidenceBundleConsumed         bool                                        `json:"evidence_bundle_consumed"`
	EvidenceArtifactCount          int                                         `json:"evidence_artifact_count"`
	EvidenceSignals                []ExternalWinAppApplicationDetailSignal     `json:"evidence_signals"`
	ReviewCards                    []ExternalWinAppApplicationDetailReviewCard `json:"review_cards"`
	PrimaryAction                  ExternalWinAppApplicationDetailAction       `json:"primary_action"`
	SecondaryActions               []ExternalWinAppApplicationDetailAction     `json:"secondary_actions"`
	AICompatibilityRuntimeOwner    bool                                        `json:"ai_compatibility_runtime_owner"`
	RuntimeOwned                   bool                                        `json:"runtime_owned"`
	GoRuntimeBacked                bool                                        `json:"go_runtime_backed"`
	KDEPolicyOwner                 bool                                        `json:"kde_policy_owner"`
	SafeForKDE                     bool                                        `json:"safe_for_kde"`
	SafeForAIDiagnostics           bool                                        `json:"safe_for_ai_diagnostics"`
	LaunchEnabled                  bool                                        `json:"launch_enabled"`
	BackendLaunchEnabled           bool                                        `json:"backend_launch_enabled"`
	ActionExecutionEnabled         bool                                        `json:"action_execution_enabled"`
	BackendProcessStarted          bool                                        `json:"backend_process_started"`
	BackendProcessEvidenceRecorded bool                                        `json:"backend_process_evidence_recorded"`
	BackendDetailsExposed          bool                                        `json:"backend_details_exposed"`
	RawPathsExposed                bool                                        `json:"raw_paths_exposed"`
	RawLauncherOutputExposed       bool                                        `json:"raw_launcher_output_exposed"`
	HostRootModified               bool                                        `json:"host_root_modified"`
	PrivilegedContainerRequired    bool                                        `json:"privileged_container_required"`
	HostNetworkingRequired         bool                                        `json:"host_networking_required"`
	DockerSocketMounted            bool                                        `json:"docker_socket_mounted"`
	BroadHostMountRequired         bool                                        `json:"broad_host_mount_required"`
	DesktopSafeSummary             string                                      `json:"desktop_safe_summary"`
}

type ExternalWinAppApplicationDetailSignal struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Verified bool   `json:"verified"`
	Summary  string `json:"summary"`
}

type ExternalWinAppApplicationDetailReviewCard struct {
	ID                    string `json:"id"`
	Title                 string `json:"title"`
	State                 string `json:"state"`
	Tone                  string `json:"tone"`
	UserVisible           bool   `json:"user_visible"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
	Summary               string `json:"summary"`
}

type ExternalWinAppApplicationDetailAction struct {
	ID                    string `json:"id"`
	Label                 string `json:"label"`
	Kind                  string `json:"kind"`
	Enabled               bool   `json:"enabled"`
	RuntimeOwned          bool   `json:"runtime_owned"`
	KDEPolicyOwner        bool   `json:"kde_policy_owner"`
	BackendDetailsExposed bool   `json:"backend_details_exposed"`
}

func PreviewExternalWinAppApplicationDetail(request ExternalWinAppApplicationDetailRequest) (ExternalWinAppApplicationDetail, error) {
	bundle, err := loadExternalWinAppApplicationDetailBundle(request.CompatibilityEvidenceBundlePath)
	if err != nil {
		return ExternalWinAppApplicationDetail{}, err
	}
	if err := validateExternalWinAppApplicationDetailBundle(bundle); err != nil {
		return ExternalWinAppApplicationDetail{}, err
	}
	fileOpenVerified := bundle.ExternalFileOpenRequested && bundle.ExternalFileBridgeReady && bundle.ExternalDesktopArgumentCount > 0
	detail := ExternalWinAppApplicationDetail{
		Version:                        bundle.Version,
		SchemaVersion:                  ExternalWinAppApplicationDetailSchemaVersion,
		RequestType:                    ExternalWinAppApplicationDetailRequestType,
		DetailType:                     "runtime-owned-existing-windows-app-detail",
		Source:                         "external-winapp-compatibility-evidence-bundle",
		RuntimeMethod:                  "PreviewExternalWinAppApplicationDetail",
		ReadMethod:                     "GetExternalWinAppApplicationDetail",
		Desktop:                        "KDE Plasma",
		ApplicationID:                  bundle.ApplicationID,
		DisplayName:                    bundle.DisplayName,
		AppVersion:                     bundle.AppVersion,
		CompatibilityState:             "real-app-run-verified",
		CompatibilityLabel:             "Verified real app run",
		PrimaryStatusTone:              "success",
		ExistingWindowsAppVerified:     true,
		RealWindowsAppRunVerified:      bundle.RealWindowsAppRunVerified,
		FileOpenVerified:               fileOpenVerified,
		RuntimeGUIEvidenceVerified:     bundle.RuntimeGUIEvidencePacketVerified,
		DesktopEvidenceVerified:        bundle.DesktopLaunchPacketVerified,
		KDEPageEvidenceVerified:        bundle.KDEExternalAppPageVerified,
		EvidenceBundleConsumed:         true,
		EvidenceArtifactCount:          bundle.EvidenceArtifactCount,
		EvidenceSignals:                externalWinAppApplicationDetailSignals(bundle, fileOpenVerified),
		ReviewCards:                    externalWinAppApplicationDetailCards(bundle, fileOpenVerified),
		PrimaryAction:                  externalWinAppApplicationDetailPrimaryAction(),
		SecondaryActions:               externalWinAppApplicationDetailSecondaryActions(),
		AICompatibilityRuntimeOwner:    true,
		RuntimeOwned:                   bundle.RuntimeOwned,
		GoRuntimeBacked:                bundle.GoRuntimeBacked,
		KDEPolicyOwner:                 bundle.KDEPolicyOwner,
		SafeForKDE:                     bundle.SafeForKDE,
		SafeForAIDiagnostics:           bundle.SafeForAIDiagnostics,
		LaunchEnabled:                  false,
		BackendLaunchEnabled:           false,
		ActionExecutionEnabled:         false,
		BackendProcessStarted:          false,
		BackendProcessEvidenceRecorded: bundle.BackendProcessStarted,
		BackendDetailsExposed:          false,
		RawPathsExposed:                false,
		RawLauncherOutputExposed:       false,
		HostRootModified:               false,
		PrivilegedContainerRequired:    false,
		HostNetworkingRequired:         false,
		DockerSocketMounted:            false,
		BroadHostMountRequired:         false,
		DesktopSafeSummary:             "This imported Windows app has verified Runtime-owned real GUI evidence and can be shown in KDE as a reviewed compatibility application without exposing host paths or backend commands.",
	}
	if !detail.SafeForKDE || !detail.SafeForAIDiagnostics || !detail.RuntimeOwned || !detail.GoRuntimeBacked || detail.KDEPolicyOwner {
		return ExternalWinAppApplicationDetail{}, errors.New("external Windows app application detail failed Runtime/KDE safety invariants")
	}
	if err := validateNoBackendTerms(detail, "external Windows app application detail"); err != nil {
		return ExternalWinAppApplicationDetail{}, err
	}
	return detail, nil
}

func loadExternalWinAppApplicationDetailBundle(path string) (ExternalWinAppCompatibilityEvidenceBundle, error) {
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return ExternalWinAppCompatibilityEvidenceBundle{}, errors.New("external Windows app application detail requires --compatibility-evidence-bundle")
	}
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, fmt.Errorf("read external Windows app compatibility evidence bundle: %w", err)
	}
	var bundle ExternalWinAppCompatibilityEvidenceBundle
	if err := json.Unmarshal(content, &bundle); err != nil {
		return ExternalWinAppCompatibilityEvidenceBundle{}, fmt.Errorf("parse external Windows app compatibility evidence bundle: %w", err)
	}
	return bundle, nil
}

func LoadExternalWinAppApplicationDetail(path string) (ExternalWinAppApplicationDetail, error) {
	cleanPath := strings.TrimSpace(path)
	if cleanPath == "" {
		return ExternalWinAppApplicationDetail{}, errors.New("external Windows app application detail path is required")
	}
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return ExternalWinAppApplicationDetail{}, fmt.Errorf("read external Windows app application detail: %w", err)
	}
	detail, err := ExternalWinAppApplicationDetailFromJSON(content)
	if err != nil {
		return ExternalWinAppApplicationDetail{}, err
	}
	return detail, nil
}

func ExternalWinAppApplicationDetailFromJSON(content []byte) (ExternalWinAppApplicationDetail, error) {
	var detail ExternalWinAppApplicationDetail
	if err := json.Unmarshal(content, &detail); err != nil {
		return ExternalWinAppApplicationDetail{}, fmt.Errorf("parse external Windows app application detail: %w", err)
	}
	if err := validateExternalWinAppApplicationDetail(detail); err != nil {
		return ExternalWinAppApplicationDetail{}, err
	}
	return detail, nil
}

func ExternalAppRecipeFromApplicationDetail(detail ExternalWinAppApplicationDetail) (Recipe, Provenance, error) {
	if err := validateExternalWinAppApplicationDetail(detail); err != nil {
		return Recipe{}, Provenance{}, err
	}
	recipe := Recipe{
		ID:                  detail.ApplicationID,
		Name:                detail.DisplayName,
		Version:             detail.AppVersion,
		Icon:                "application-x-executable",
		Mode:                "automatic",
		SupportedExtensions: []string{},
	}
	if err := recipe.Validate(); err != nil {
		return Recipe{}, Provenance{}, err
	}
	return recipe, Provenance{
		Source:          "external-winapp-application-detail",
		RegistryName:    "runtime-observed-external-app",
		DigestVerified:  true,
		SignatureStatus: "observed-runtime-application-detail",
	}, nil
}

func validateExternalWinAppApplicationDetail(detail ExternalWinAppApplicationDetail) error {
	switch {
	case detail.SchemaVersion != ExternalWinAppApplicationDetailSchemaVersion || detail.RequestType != ExternalWinAppApplicationDetailRequestType:
		return errors.New("external Windows app application detail requires a Runtime application detail payload")
	case strings.TrimSpace(detail.ApplicationID) == "" || strings.TrimSpace(detail.DisplayName) == "":
		return errors.New("external Windows app application detail requires application identity")
	case !detail.EvidenceBundleConsumed || !detail.ExistingWindowsAppVerified || !detail.RealWindowsAppRunVerified || !detail.FileOpenVerified:
		return errors.New("external Windows app application detail requires verified existing app evidence")
	case !detail.RuntimeGUIEvidenceVerified || !detail.DesktopEvidenceVerified || !detail.KDEPageEvidenceVerified:
		return errors.New("external Windows app application detail requires complete Runtime and desktop evidence")
	case !detail.SafeForKDE || !detail.SafeForAIDiagnostics || !detail.RuntimeOwned || !detail.GoRuntimeBacked || detail.KDEPolicyOwner:
		return errors.New("external Windows app application detail requires Runtime-owned KDE-safe evidence")
	case detail.LaunchEnabled || detail.BackendLaunchEnabled || detail.ActionExecutionEnabled || detail.BackendProcessStarted:
		return errors.New("external Windows app application detail must remain read-only for KDE consumption")
	case detail.BackendDetailsExposed || detail.RawPathsExposed || detail.RawLauncherOutputExposed || detail.HostRootModified:
		return errors.New("external Windows app application detail requires redacted host-safe evidence")
	case detail.PrivilegedContainerRequired || detail.HostNetworkingRequired || detail.DockerSocketMounted || detail.BroadHostMountRequired:
		return errors.New("external Windows app application detail requires non-privileged isolated evidence")
	}
	return nil
}

func validateExternalWinAppApplicationDetailBundle(bundle ExternalWinAppCompatibilityEvidenceBundle) error {
	switch {
	case bundle.SchemaVersion != ExternalWinAppCompatibilityEvidenceBundleSchemaVersion || bundle.RequestType != ExternalWinAppCompatibilityEvidenceBundleRequestType:
		return errors.New("external Windows app application detail requires a compatibility evidence bundle")
	case strings.TrimSpace(bundle.ApplicationID) == "" || strings.TrimSpace(bundle.DisplayName) == "":
		return errors.New("external Windows app application detail requires bundle application identity")
	case !bundle.RealWindowsAppRunVerified || !bundle.RuntimeGUIEvidencePacketVerified || !bundle.DesktopLaunchPacketVerified || !bundle.KDEExternalAppPageVerified:
		return errors.New("external Windows app application detail requires complete verified real app evidence")
	case !bundle.SafeForKDE || !bundle.SafeForAIDiagnostics || !bundle.RuntimeOwned || !bundle.GoRuntimeBacked || bundle.KDEPolicyOwner:
		return errors.New("external Windows app application detail requires Runtime-owned KDE-safe evidence")
	case bundle.BackendDetailsExposed || bundle.RawPathsExposed || bundle.RawLauncherOutputExposed || bundle.HostRootModified:
		return errors.New("external Windows app application detail requires redacted host-safe evidence")
	case bundle.PrivilegedContainerRequired || bundle.HostNetworkingRequired || bundle.DockerSocketMounted || bundle.BroadHostMountRequired:
		return errors.New("external Windows app application detail requires non-privileged isolated evidence")
	}
	return nil
}

func externalWinAppApplicationDetailSignals(bundle ExternalWinAppCompatibilityEvidenceBundle, fileOpenVerified bool) []ExternalWinAppApplicationDetailSignal {
	return []ExternalWinAppApplicationDetailSignal{
		{ID: "real-gui-run", Label: "Real GUI run", Verified: bundle.RealWindowsAppRunVerified && bundle.WindowObserved && bundle.XWindowObserved, Summary: "Runtime observed a real GUI window for this imported Windows app."},
		{ID: "file-open", Label: "File open", Verified: fileOpenVerified, Summary: "Runtime verified a KDE-style file-open argument through the controlled file bridge."},
		{ID: "desktop-evidence", Label: "Desktop evidence", Verified: bundle.DesktopLaunchPacketVerified && bundle.KDEExternalAppPageVerified, Summary: "Runtime produced KDE-safe desktop and application page evidence."},
		{ID: "safety-boundary", Label: "Safety boundary", Verified: bundle.SafeForKDE && !bundle.HostRootModified && !bundle.RawPathsExposed, Summary: "Runtime evidence is redacted and safe for desktop display."},
	}
}

func externalWinAppApplicationDetailCards(bundle ExternalWinAppCompatibilityEvidenceBundle, fileOpenVerified bool) []ExternalWinAppApplicationDetailReviewCard {
	return []ExternalWinAppApplicationDetailReviewCard{
		{ID: "compatibility", Title: "Compatibility evidence", State: "verified", Tone: "success", UserVisible: true, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false, Summary: "Real GUI run evidence is available for review."},
		{ID: "file-open", Title: "File-open behavior", State: boolState(fileOpenVerified), Tone: boolTone(fileOpenVerified), UserVisible: true, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false, Summary: "File-open bridge evidence is recorded for this app."},
		{ID: "safety", Title: "Desktop safety", State: boolState(bundle.SafeForKDE), Tone: boolTone(bundle.SafeForKDE), UserVisible: true, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false, Summary: "KDE can display this app without owning compatibility policy."},
	}
}

func externalWinAppApplicationDetailPrimaryAction() ExternalWinAppApplicationDetailAction {
	return ExternalWinAppApplicationDetailAction{ID: "review-compatibility-evidence", Label: "Review compatibility evidence", Kind: "review", Enabled: true, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false}
}

func externalWinAppApplicationDetailSecondaryActions() []ExternalWinAppApplicationDetailAction {
	return []ExternalWinAppApplicationDetailAction{
		{ID: "open-ai-diagnostics", Label: "Open AI diagnostics", Kind: "diagnostics", Enabled: true, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false},
		{ID: "launch-after-review", Label: "Launch after Runtime review", Kind: "runtime-gated-launch", Enabled: false, RuntimeOwned: true, KDEPolicyOwner: false, BackendDetailsExposed: false},
	}
}

func boolState(value bool) string {
	if value {
		return "verified"
	}
	return "missing"
}

func boolTone(value bool) string {
	if value {
		return "success"
	}
	return "warning"
}
