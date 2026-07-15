package appidentity

type KDEAIAnalysisLink struct {
	Source                string `json:"source"`
	RuntimeMethod         string `json:"runtime_method"`
	Surface               string `json:"surface"`
	Task                  string `json:"task"`
	Disclosure            string `json:"disclosure"`
	Available             bool   `json:"available"`
	UserVisible           bool   `json:"user_visible"`
	RequiresPortal        bool   `json:"requires_portal"`
	SafeForAIDiagnostics  bool   `json:"safe_for_ai_diagnostics"`
	AIProviderCallEnabled bool   `json:"ai_provider_call_enabled"`
	NetworkRequired       bool   `json:"network_required"`
	FileContentRead       bool   `json:"file_content_read"`
	FilePathsExposed      bool   `json:"file_paths_exposed"`
	RequestObjectCreated  bool   `json:"request_object_created"`
	PermissionGranted     bool   `json:"permission_granted"`
	BackendLaunchEnabled  bool   `json:"backend_launch_enabled"`
}

func kdeAIAnalysisLinkForEntryPoint(entryPointID string) *KDEAIAnalysisLink {
	if entryPointID != "file-manager" {
		return nil
	}
	link := KDEAIAnalysisLink{
		Source:                "dolphin-ai-analysis-preview",
		RuntimeMethod:         "GetAIDiagnosticInput",
		Surface:               "Dolphin",
		Task:                  "compatibility-file-review",
		Disclosure:            "count-and-extension-only",
		Available:             true,
		UserVisible:           true,
		RequiresPortal:        true,
		SafeForAIDiagnostics:  true,
		AIProviderCallEnabled: false,
		NetworkRequired:       false,
		FileContentRead:       false,
		FilePathsExposed:      false,
		RequestObjectCreated:  false,
		PermissionGranted:     false,
		BackendLaunchEnabled:  false,
	}
	return &link
}
