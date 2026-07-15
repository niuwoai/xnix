package appidentity

type DolphinAIAnalysisPreview struct {
	SchemaVersion          string            `json:"schema_version"`
	RequestType            string            `json:"request_type"`
	Source                 string            `json:"source"`
	Desktop                string            `json:"desktop"`
	ApplicationID          string            `json:"application_id"`
	DisplayName            string            `json:"display_name"`
	DesktopFile            string            `json:"desktop_file"`
	RuntimeMethod          string            `json:"runtime_method"`
	AnalysisTask           string            `json:"analysis_task"`
	AnalysisSurface        string            `json:"analysis_surface"`
	SelectionMode          string            `json:"selection_mode"`
	FileCount              int               `json:"file_count"`
	SelectedExtension      string            `json:"selected_extension"`
	SelectedFileDisclosure string            `json:"selected_file_disclosure"`
	PortalRequired         bool              `json:"portal_required"`
	PortalInterface        string            `json:"portal_interface"`
	PortalMethod           string            `json:"portal_method"`
	RuntimeOwned           bool              `json:"runtime_owned"`
	GoRuntimeBacked        bool              `json:"go_runtime_backed"`
	KDEPolicyOwner         bool              `json:"kde_policy_owner"`
	UserVisible            bool              `json:"user_visible"`
	SafeForAIDiagnostics   bool              `json:"safe_for_ai_diagnostics"`
	UserReviewRequired     bool              `json:"user_review_required"`
	AIProviderCallEnabled  bool              `json:"ai_provider_call_enabled"`
	NetworkRequired        bool              `json:"network_required"`
	FileContentRead        bool              `json:"file_content_read"`
	FilePathsExposed       bool              `json:"file_paths_exposed"`
	RequestObjectCreated   bool              `json:"request_object_created"`
	PermissionGranted      bool              `json:"permission_granted"`
	BackendLaunchEnabled   bool              `json:"backend_launch_enabled"`
	HostRootModified       bool              `json:"host_root_modified"`
	BackendDetailsExposed  bool              `json:"backend_details_exposed"`
	UserFacingSettings     map[string]string `json:"user_facing_settings"`
	AllowedAITasks         []string          `json:"allowed_ai_tasks"`
	BlockedActions         []string          `json:"blocked_actions"`
	Summary                FileOpenSummary   `json:"summary"`
}

func NewDolphinAIAnalysisPreview(recipes []Recipe, provenance Provenance, fileURIs []string, applicationID string) (DolphinAIAnalysisPreview, error) {
	fileOpen, err := NewFileOpenPreview(recipes, provenance, fileURIs, applicationID)
	if err != nil {
		return DolphinAIAnalysisPreview{}, err
	}

	preview := DolphinAIAnalysisPreview{
		SchemaVersion:          "xnix.runtime.dolphin_ai_analysis.v1",
		RequestType:            "dolphin-ai-analysis-preview",
		Source:                 "dolphin-ai-action",
		Desktop:                fileOpen.Desktop,
		ApplicationID:          fileOpen.ApplicationID,
		DisplayName:            fileOpen.DisplayName,
		DesktopFile:            fileOpen.DesktopFile,
		RuntimeMethod:          "GetAIDiagnosticInput",
		AnalysisTask:           "compatibility-file-review",
		AnalysisSurface:        "Dolphin",
		SelectionMode:          fileOpen.SelectionMode,
		FileCount:              fileOpen.FileCount,
		SelectedExtension:      fileOpen.SelectedExtension,
		SelectedFileDisclosure: "count-and-extension-only",
		PortalRequired:         true,
		PortalInterface:        fileOpen.PortalInterface,
		PortalMethod:           fileOpen.PortalMethod,
		RuntimeOwned:           true,
		GoRuntimeBacked:        true,
		KDEPolicyOwner:         false,
		UserVisible:            true,
		SafeForAIDiagnostics:   true,
		UserReviewRequired:     true,
		AIProviderCallEnabled:  false,
		NetworkRequired:        false,
		FileContentRead:        false,
		FilePathsExposed:       false,
		RequestObjectCreated:   false,
		PermissionGranted:      false,
		BackendLaunchEnabled:   false,
		HostRootModified:       false,
		BackendDetailsExposed:  false,
		UserFacingSettings:     fileOpen.UserFacingSettings,
		AllowedAITasks: []string{
			"explain compatibility risk",
			"summarize required user review",
			"suggest safe next steps",
		},
		BlockedActions: []string{
			"read selected file contents before Portal approval",
			"send user documents to an AI provider",
			"expose selected file paths to AI diagnostics",
			"create Portal request objects from an AI preview",
			"grant file permissions from an AI preview",
			"start compatibility backends from an AI preview",
			"mutate the host root from an AI preview",
			"expose backend implementation details in AI prompts",
		},
		Summary: FileOpenSummary{
			Headline: "Dolphin can ask the Runtime for an AI-safe file review.",
			Detail:   "The preview only exposes file count and extension; file contents, paths, network calls, permissions, and execution stay blocked.",
		},
	}
	if err := validateNoBackendTerms(preview, "Dolphin AI analysis preview"); err != nil {
		return DolphinAIAnalysisPreview{}, err
	}
	return preview, nil
}
